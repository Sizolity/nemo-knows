package runslog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func setupRunsDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("XDG_STATE_HOME", dir)
	want := filepath.Join(dir, "nemo", "runs")
	if err := os.MkdirAll(want, 0o755); err != nil {
		t.Fatal(err)
	}
	return want
}

func TestAppendRun_WritesExpectedFields(t *testing.T) {
	dir := setupRunsDir(t)
	start := time.Date(2026, 6, 17, 17, 30, 0, 0, time.UTC)
	end := start.Add(4*time.Minute + 12*time.Second)
	if err := AppendRun(RunEntry{
		RunID:      "20260617-173000-aabbccdd",
		Subcommand: "ingest",
		Source:     "wiki/sources/example.md",
		Status:     "ok",
		Start:      start,
		End:        end,
		WikiRoot:   "/home/u/.wiki",
		Provider:   "deepseek",
		Profile:    "stable",
	}); err != nil {
		t.Fatalf("AppendRun: %v", err)
	}
	path := filepath.Join(dir, "2026-06-17.md")
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read log: %v", err)
	}
	s := string(got)
	required := []string{
		"## 20260617-173000-aabbccdd ingest wiki/sources/example.md",
		"- start: 2026-06-17T17:30:00Z",
		"- end: 2026-06-17T17:34:12Z",
		"- status: ok",
		"- duration: 4m12s",
		"- source: wiki/sources/example.md",
		"- provider: deepseek",
		"- profile: stable",
		"- wiki-root: /home/u/.wiki",
	}
	for _, want := range required {
		if !strings.Contains(s, want) {
			t.Errorf("log missing %q\n---\n%s", want, s)
		}
	}
}

func TestAppendRun_AppendsAcrossCalls(t *testing.T) {
	setupRunsDir(t)
	at := time.Date(2026, 6, 17, 12, 0, 0, 0, time.UTC)
	if err := AppendRun(RunEntry{RunID: "a", Subcommand: "lint", Status: "ok", Start: at, End: at}); err != nil {
		t.Fatal(err)
	}
	if err := AppendRun(RunEntry{RunID: "b", Subcommand: "lint", Status: "ok", Start: at, End: at}); err != nil {
		t.Fatal(err)
	}
	day := at.Format("2006-01-02")
	dir, _ := RunsDir()
	body, err := os.ReadFile(filepath.Join(dir, day+".md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), "## a lint") || !strings.Contains(string(body), "## b lint") {
		t.Fatalf("expected both run sections, got:\n%s", body)
	}
}

func TestIsDailyLogName(t *testing.T) {
	good := []string{"2026-06-17.md", "1999-12-31.md"}
	bad := []string{"latest.md", ".last-gc", "2026-06.md", "lint/foo.md", "2026-06-17.md.bak"}
	for _, g := range good {
		if !IsDailyLogName(g) {
			t.Errorf("want true for %q", g)
		}
	}
	for _, b := range bad {
		if IsDailyLogName(b) {
			t.Errorf("want false for %q", b)
		}
	}
}

func TestGC_RetentionAndDryRun(t *testing.T) {
	dir := setupRunsDir(t)
	now := time.Date(2026, 6, 17, 0, 0, 0, 0, time.UTC)
	// 5 daily files at varying ages.
	makeAged := func(day int, ageDays int) string {
		when := now.AddDate(0, 0, -ageDays)
		name := when.Format("2006-01-02") + ".md"
		p := filepath.Join(dir, name)
		if err := os.WriteFile(p, []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
		if err := os.Chtimes(p, when, when); err != nil {
			t.Fatal(err)
		}
		return p
	}
	old1 := makeAged(0, 40)
	old2 := makeAged(0, 35)
	fresh1 := makeAged(0, 5)
	fresh2 := makeAged(0, 1)
	// keep file that's exactly at the cutoff (not strictly before)
	keep := makeAged(0, 30)

	// dry-run reports old files but removes nothing
	res, err := GC(GCOptions{RetentionDays: 30, DryRun: true, Now: now})
	if err != nil {
		t.Fatalf("GC dry-run: %v", err)
	}
	if len(res.Removed) != 2 {
		t.Fatalf("dry-run want 2 removed, got %d: %v", len(res.Removed), res.Removed)
	}
	for _, p := range res.Removed {
		if _, err := os.Stat(p); err != nil {
			t.Fatalf("dry-run should not delete %q (%v)", p, err)
		}
	}

	// real GC removes the old ones, keeps the fresh + boundary.
	res, err = GC(GCOptions{RetentionDays: 30, Now: now})
	if err != nil {
		t.Fatalf("GC: %v", err)
	}
	if len(res.Removed) != 2 {
		t.Fatalf("want 2 removed, got %d", len(res.Removed))
	}
	for _, p := range []string{old1, old2} {
		if _, err := os.Stat(p); err == nil {
			t.Errorf("old file should be removed: %q", p)
		}
	}
	for _, p := range []string{fresh1, fresh2, keep} {
		if _, err := os.Stat(p); err != nil {
			t.Errorf("fresh file should stay: %q (%v)", p, err)
		}
	}
	// .last-gc was touched
	if _, err := os.Stat(filepath.Join(dir, ".last-gc")); err != nil {
		t.Errorf(".last-gc not touched: %v", err)
	}
}

func TestMaybeAutoGC_24hThrottle(t *testing.T) {
	dir := setupRunsDir(t)
	now := time.Date(2026, 6, 17, 12, 0, 0, 0, time.UTC)

	// First call: no .last-gc → should run
	ran, _, err := MaybeAutoGC(GCOptions{RetentionDays: 30, Now: now})
	if err != nil {
		t.Fatalf("first MaybeAutoGC: %v", err)
	}
	if !ran {
		t.Fatal("first call should run GC (no stamp)")
	}
	info, err := os.Stat(filepath.Join(dir, ".last-gc"))
	if err != nil {
		t.Fatalf(".last-gc should exist: %v", err)
	}
	stamped := info.ModTime()
	// .last-gc should be roughly = now (touch wrote it).
	if delta := now.Sub(stamped); delta < 0 || delta > time.Second {
		t.Fatalf(".last-gc time should equal now (got %v, now %v)", stamped, now)
	}

	// Second call 1 hour later: should be throttled
	later := now.Add(time.Hour)
	ran, _, err = MaybeAutoGC(GCOptions{RetentionDays: 30, Now: later})
	if err != nil {
		t.Fatalf("throttled MaybeAutoGC: %v", err)
	}
	if ran {
		t.Fatal("should be throttled within 24h")
	}

	// Third call 25 hours later: should run again
	stale := now.Add(25 * time.Hour)
	ran, _, err = MaybeAutoGC(GCOptions{RetentionDays: 30, Now: stale})
	if err != nil {
		t.Fatalf("post-24h MaybeAutoGC: %v", err)
	}
	if !ran {
		t.Fatal("should run after 24h")
	}
}

func TestGC_ZeroRetentionMeansForever(t *testing.T) {
	setupRunsDir(t)
	res, err := GC(GCOptions{RetentionDays: 0, Now: time.Now()})
	if err != nil {
		t.Fatalf("GC zero retention: %v", err)
	}
	if len(res.Removed) != 0 || res.Examined != 0 {
		t.Fatalf("retention=0 should be a no-op, got %#v", res)
	}
}
