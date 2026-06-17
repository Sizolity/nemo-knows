package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/huic/nemo-knows/internal/rundir"
)

// cmdGC manually removes failed-run checkpoint directories older than
// --days. There is no automatic checkpoint GC in V0; failed checkpoints
// are kept forever until the operator runs this command.
//
// --days default cascade: --days flag > $NEMO_CHECKPOINT_RETENTION_DAYS > 7
func cmdGC(args []string, g *globalConfig) int {
	defaultDays := envIntDefault("NEMO_CHECKPOINT_RETENTION_DAYS", 7)
	fs := flag.NewFlagSet("nemocli gc", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	days := fs.Int("days", defaultDays, "remove checkpoints older than N days (0 = keep forever)")
	dryRun := fs.Bool("dry-run", false, "list what would be removed, don't remove")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *days <= 0 {
		fmt.Fprintln(os.Stderr, "gc: days=0 (keep forever); nothing to do")
		return 0
	}

	base := rundir.CheckpointBase
	entries, err := os.ReadDir(base)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "gc: no checkpoint dir at %s; nothing to do\n", base)
			return 0
		}
		fmt.Fprintln(os.Stderr, "gc:", err)
		return 1
	}
	cutoff := time.Now().Add(-time.Duration(*days) * 24 * time.Hour)
	removed := 0
	for _, ent := range entries {
		if !ent.IsDir() {
			continue
		}
		full := filepath.Join(base, ent.Name())
		info, err := ent.Info()
		if err != nil {
			continue
		}
		if !info.ModTime().Before(cutoff) {
			continue
		}
		if *dryRun {
			fmt.Println(full)
		} else {
			if err := os.RemoveAll(full); err != nil {
				fmt.Fprintln(os.Stderr, "gc: remove", full, ":", err)
				continue
			}
			fmt.Println(full)
		}
		removed++
	}
	if *dryRun {
		fmt.Fprintf(os.Stderr, "gc: %d checkpoint(s) would be removed (dry run)\n", removed)
	} else {
		fmt.Fprintf(os.Stderr, "gc: %d checkpoint(s) removed\n", removed)
	}
	if g.Verbose {
		fmt.Fprintf(os.Stderr, "gc: cutoff=%s base=%s\n", cutoff.Format(time.RFC3339), base)
	}
	return 0
}
