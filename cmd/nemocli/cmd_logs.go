package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/huic/nemo-knows/internal/runslog"
)

// cmdLogs reads (and optionally tails / GCs) the runs log under
// ~/.local/state/nemo/runs/. The default action concatenates the last
// seven days of daily files to stdout; --follow tails today's file;
// --gc removes daily files older than --retention-days (default
// $NEMO_RUNS_RETENTION_DAYS, falling back to 30).
func cmdLogs(args []string, g *globalConfig) int {
	fs := flag.NewFlagSet("nemocli logs", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	follow := fs.Bool("follow", false, "tail today's runs log")
	fs.BoolVar(follow, "f", false, "shorthand for -follow")
	doGC := fs.Bool("gc", false, "remove daily run files older than --retention-days")
	dryRun := fs.Bool("dry-run", false, "with --gc, only print what would be removed")
	retentionDays := fs.Int("retention-days", envIntDefault("NEMO_RUNS_RETENTION_DAYS", 30),
		"days to keep when --gc is set")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	if *doGC {
		return runLogsGC(*retentionDays, *dryRun)
	}
	if *follow {
		return runLogsFollow()
	}
	return runLogsCat()
}

func runLogsGC(retentionDays int, dryRun bool) int {
	res, err := runslog.GC(runslog.GCOptions{RetentionDays: retentionDays, DryRun: dryRun})
	if err != nil {
		fmt.Fprintln(os.Stderr, "logs: gc:", err)
		return 1
	}
	if dryRun {
		fmt.Fprintf(os.Stderr, "would remove %d file(s):\n", len(res.Removed))
	} else {
		fmt.Fprintf(os.Stderr, "removed %d file(s):\n", len(res.Removed))
	}
	for _, p := range res.Removed {
		fmt.Println(p)
	}
	return 0
}

// runLogsCat prints the most recent 7 daily files (oldest first).
func runLogsCat() int {
	dir, err := runslog.RunsDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "logs:", err)
		return 1
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "logs: no runs yet (", dir, "does not exist)")
			return 0
		}
		fmt.Fprintln(os.Stderr, "logs:", err)
		return 1
	}
	names := []string{}
	for _, e := range entries {
		if !runslog.IsDailyLogName(e.Name()) {
			continue
		}
		names = append(names, e.Name())
	}
	sort.Strings(names)
	if len(names) > 7 {
		names = names[len(names)-7:]
	}
	for _, n := range names {
		body, err := os.ReadFile(filepath.Join(dir, n))
		if err != nil {
			fmt.Fprintln(os.Stderr, "logs:", err)
			return 1
		}
		fmt.Fprintf(os.Stdout, "# %s\n\n", n)
		_, _ = os.Stdout.Write(body)
		fmt.Println()
	}
	return 0
}

// runLogsFollow tails today's daily file. When the date rolls past
// midnight the function continues reading the new file. It is a thin
// loop and not designed for daemon-grade efficiency.
func runLogsFollow() int {
	dir, err := runslog.RunsDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "logs:", err)
		return 1
	}
	current := todayPath(dir)
	f, err := openFollow(current)
	if err != nil {
		fmt.Fprintln(os.Stderr, "logs:", err)
		return 1
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		// Drain whatever new content is available.
		if _, err := io.Copy(os.Stdout, br); err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "logs:", err)
			return 1
		}
		time.Sleep(500 * time.Millisecond)

		// Detect a midnight roll-over by recomputing today's filename.
		next := todayPath(dir)
		if next != current {
			f.Close()
			current = next
			var openErr error
			f, openErr = openFollow(current)
			if openErr != nil {
				fmt.Fprintln(os.Stderr, "logs:", openErr)
				return 1
			}
			br = bufio.NewReader(f)
		}
	}
}

func todayPath(dir string) string {
	return filepath.Join(dir, time.Now().Format("2006-01-02")+".md")
}

func openFollow(path string) (*os.File, error) {
	// Create the file if it doesn't exist yet so tail can attach
	// immediately; nemocli creates daily files on demand via AppendRun.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return nil, err
		}
		f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0o644)
		if err != nil {
			return nil, err
		}
		f.Close()
	}
	return os.Open(path)
}
