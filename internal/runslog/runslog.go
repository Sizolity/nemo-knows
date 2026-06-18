// Package runslog writes the nemocli runtime audit log.
//
// Runs are recorded as append-only Markdown sections under
// $XDG_STATE_HOME/nemo/runs/ (or ~/.local/state/nemo/runs/), with one
// daily file per local date. Each section is small enough to fit in a
// single write(2) syscall when possible, preserving append atomicity on
// POSIX filesystems.
//
// This package never writes to $WIKI_ROOT/log.md; that file is reserved
// for wiki-content actions enforced by the wikilint allowlist
// (ingest/query-filed/lint/schema-change/note). See the wiki-log-boundary
// rule in .cursor/rules/ for the boundary contract.
package runslog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"
	"time"
)

// RunEntry holds the data nemocli records for one run.
//
// Fields are intentionally flat: rendering happens in renderEntry; the
// caller fills in whichever fields apply to its subcommand (e.g., Source
// and Checkpoint are ingest-specific).
type RunEntry struct {
	RunID      string    // e.g. "20260617-173042-a1b2c3d4"
	Subcommand string    // "ingest" | "lint" | "maintain" | "once" | ...
	Source     string    // ingest: source file path; "" otherwise
	Status     string    // "ok" | "fail" | "skipped"
	Start      time.Time // wall clock at run start
	End        time.Time // wall clock at run end (== Start if pre-run failure)
	WikiRoot   string    // resolved $WIKI_ROOT
	Provider   string    // "deepseek" | "llama" | "" if subcommand doesn't need it
	Profile    string    // "stable" | "fast" | ...; "" if not applicable
	Checkpoint string    // path under /tmp/nemo-checkpoints/ on failure
	Error      string    // one-line error summary on failure
}

// RunsDir returns the directory that holds daily run log files.
//
// Honors XDG_STATE_HOME if set; otherwise falls back to
// $HOME/.local/state/nemo/runs.
func RunsDir() (string, error) {
	if v := strings.TrimSpace(os.Getenv("XDG_STATE_HOME")); v != "" {
		return filepath.Join(v, "nemo", "runs"), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home for runs dir: %w", err)
	}
	return filepath.Join(home, ".local", "state", "nemo", "runs"), nil
}

// AppendRun appends entry to today's daily Markdown log. The file is
// opened with O_APPEND|O_CREATE and 0644; missing parent directories are
// created with 0755.
func AppendRun(entry RunEntry) error {
	if entry.Start.IsZero() {
		entry.Start = time.Now()
	}
	if entry.End.IsZero() {
		entry.End = entry.Start
	}
	dir, err := RunsDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create runs dir: %w", err)
	}
	day := entry.Start.Format("2006-01-02")
	path := filepath.Join(dir, day+".md")
	body := renderEntry(entry)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open runs file %s: %w", path, err)
	}
	defer f.Close()
	if _, err := f.WriteString(body); err != nil {
		return fmt.Errorf("write runs entry: %w", err)
	}
	return nil
}

func renderEntry(e RunEntry) string {
	var b strings.Builder
	title := fmt.Sprintf("## %s %s", e.RunID, e.Subcommand)
	if e.Source != "" {
		title += " " + e.Source
	}
	b.WriteString(title)
	b.WriteByte('\n')
	b.WriteString(fmt.Sprintf("- start: %s\n", e.Start.Format(time.RFC3339)))
	b.WriteString(fmt.Sprintf("- end: %s\n", e.End.Format(time.RFC3339)))
	b.WriteString(fmt.Sprintf("- status: %s\n", emptyDash(e.Status)))
	if dur := e.End.Sub(e.Start); dur >= 0 {
		b.WriteString(fmt.Sprintf("- duration: %s\n", dur.Round(time.Second)))
	}
	if e.Source != "" {
		b.WriteString(fmt.Sprintf("- source: %s\n", e.Source))
	}
	if e.Provider != "" {
		b.WriteString(fmt.Sprintf("- provider: %s\n", e.Provider))
	}
	if e.Profile != "" {
		b.WriteString(fmt.Sprintf("- profile: %s\n", e.Profile))
	}
	if e.WikiRoot != "" {
		b.WriteString(fmt.Sprintf("- wiki-root: %s\n", e.WikiRoot))
	}
	if e.Checkpoint != "" {
		b.WriteString(fmt.Sprintf("- checkpoint: %s\n", e.Checkpoint))
	}
	if e.Error != "" {
		b.WriteString(fmt.Sprintf("- error: %s\n", oneLine(e.Error)))
	}
	b.WriteByte('\n')
	return b.String()
}

func emptyDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

func oneLine(s string) string {
	s = strings.ReplaceAll(s, "\r\n", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.TrimSpace(s)
	if len(s) > 240 {
		s = s[:240] + "..."
	}
	return s
}

// dailyLogRE matches YYYY-MM-DD.md so GC never touches latest.md /
// .last-gc / non-daily files in the same directory.
var dailyLogRE = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}\.md$`)

// IsDailyLogName reports whether name looks like a daily-rolled file.
// Exposed for testing GC paths.
func IsDailyLogName(name string) bool {
	return dailyLogRE.MatchString(name)
}

// GCResult summarizes a synchronous GC pass.
type GCResult struct {
	Examined int      // files inspected
	Removed  []string // file paths removed (or that would be removed under DryRun)
}

// GCOptions configures GC behavior.
//
// RetentionDays<=0 disables GC entirely (treat as "keep forever"). Now
// is settable so tests can inject a fixed clock; production callers
// usually pass time.Now().
type GCOptions struct {
	RetentionDays int
	DryRun        bool
	Now           time.Time
}

// GC removes daily log files older than RetentionDays. It always
// inspects every daily log; .last-gc and latest.md are skipped.
//
// Touches .last-gc on success when DryRun is false so MaybeAutoGC's 24h
// throttle can read it back.
func GC(opts GCOptions) (GCResult, error) {
	res := GCResult{}
	if opts.RetentionDays <= 0 {
		return res, nil
	}
	if opts.Now.IsZero() {
		opts.Now = time.Now()
	}
	dir, err := RunsDir()
	if err != nil {
		return res, err
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return res, nil
	} else if err != nil {
		return res, err
	}
	cutoff := opts.Now.Add(-time.Duration(opts.RetentionDays) * 24 * time.Hour)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return res, err
	}
	for _, ent := range entries {
		if ent.IsDir() {
			continue
		}
		name := ent.Name()
		if !IsDailyLogName(name) {
			continue
		}
		res.Examined++
		info, err := ent.Info()
		if err != nil {
			continue
		}
		if !info.ModTime().Before(cutoff) {
			continue
		}
		full := filepath.Join(dir, name)
		res.Removed = append(res.Removed, full)
		if !opts.DryRun {
			_ = os.Remove(full)
		}
	}
	if !opts.DryRun {
		_ = touch(filepath.Join(dir, ".last-gc"), opts.Now)
	}
	return res, nil
}

// MaybeAutoGC runs GC only when the .last-gc stamp is missing or older
// than 24h. Designed for the per-process startup goroutine.
//
// Returns (true, result, err) when GC actually ran; (false, _, nil)
// when throttled.
func MaybeAutoGC(opts GCOptions) (bool, GCResult, error) {
	if opts.RetentionDays <= 0 {
		return false, GCResult{}, nil
	}
	if opts.Now.IsZero() {
		opts.Now = time.Now()
	}
	dir, err := RunsDir()
	if err != nil {
		return false, GCResult{}, err
	}
	stampPath := filepath.Join(dir, ".last-gc")
	if info, err := os.Stat(stampPath); err == nil {
		if opts.Now.Sub(info.ModTime()) < 24*time.Hour {
			return false, GCResult{}, nil
		}
	}
	res, err := GC(opts)
	return true, res, err
}

// StartGCAsync launches MaybeAutoGC in a goroutine, swallowing panics.
//
// All nemocli subcommands call this from main before dispatch so the
// background sweep never blocks the user-visible work.
func StartGCAsync(retentionDays int) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("runslog auto-GC panic: %v\n%s", r, debug.Stack())
			}
		}()
		if _, _, err := MaybeAutoGC(GCOptions{RetentionDays: retentionDays}); err != nil {
			log.Printf("runslog auto-GC: %v", err)
		}
	}()
}

// touch sets the mtime/atime of path to t, creating the file if missing.
func touch(path string, t time.Time) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	_ = f.Close()
	return os.Chtimes(path, t, t)
}
