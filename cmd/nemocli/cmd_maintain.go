package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/huic/nemo-knows/internal/cliroot"
	"github.com/huic/nemo-knows/internal/config"
	"github.com/huic/nemo-knows/internal/runslog"
	"github.com/huic/nemo-knows/internal/wikimaint"
)

// cmdMaintain is the V0 maintenance pass. The mode parameter is passed
// straight through to internal/wikimaint, whose accepted values are
// listed in wikimaint.ModeReport / ModeSafe / ModePropose / ModeAuto.
//
// Propose / auto modes need a model generator; we build one from the
// resolved config (and that's the only path that touches the LLM).
func cmdMaintain(args []string, g *globalConfig) int {
	fs := flag.NewFlagSet("nemocli maintain", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	mode := fs.String("mode", wikimaint.ModeSafe, "report | safe | propose | auto")
	outDir := fs.String("out-dir", "", "directory to write wiki-maintain.json/.md (default ~/.local/state/nemo/maintain/<date>/)")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if err := cliroot.RequireInitialized(g.WikiRoot); err != nil {
		fmt.Fprintln(os.Stderr, "maintain:", err)
		return 2
	}

	resolvedOut, err := resolveMaintainOutDir(*outDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "maintain:", err)
		return 1
	}
	if err := os.MkdirAll(resolvedOut, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "maintain: create out dir:", err)
		return 1
	}

	start := time.Now()
	opts := wikimaint.Options{
		Mode:   *mode,
		OutDir: resolvedOut,
	}
	if *mode == wikimaint.ModePropose || *mode == wikimaint.ModeAuto {
		cfg, err := config.ForProfileWithProvider(g.Profile, g.Provider)
		if err != nil {
			fmt.Fprintln(os.Stderr, "maintain:", err)
			logMaintainResult(g, start, *mode, "fail", err.Error())
			return 1
		}
		opts.Generator = generatorFromCfg(cfg)
	}

	result, err := wikimaint.Maintain(g.WikiRoot, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "maintain:", err)
		logMaintainResult(g, start, *mode, "fail", err.Error())
		return 1
	}
	fmt.Fprintf(os.Stderr, "wrote %s and %s\n",
		filepath.Join(resolvedOut, "wiki-maintain.json"),
		filepath.Join(resolvedOut, "wiki-maintain.md"))
	if result.Changed {
		fmt.Fprintln(os.Stderr, "wiki maintenance applied safe changes")
	}
	logMaintainResult(g, start, *mode, "ok", "")
	return 0
}

// resolveMaintainOutDir returns the user-supplied --out-dir or the
// XDG-friendly default ~/.local/state/nemo/maintain/<date>/.
func resolveMaintainOutDir(flagVal string) (string, error) {
	if flagVal != "" {
		return flagVal, nil
	}
	base, err := nemoStateDir()
	if err != nil {
		return "", err
	}
	day := time.Now().Format("2006-01-02")
	return filepath.Join(base, "maintain", day), nil
}

func nemoStateDir() (string, error) {
	if v := os.Getenv("XDG_STATE_HOME"); v != "" {
		return filepath.Join(v, "nemo"), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "state", "nemo"), nil
}

func logMaintainResult(g *globalConfig, start time.Time, mode, status, errMsg string) {
	entry := runslog.RunEntry{
		RunID:      newRunIDForCmd(start),
		Subcommand: "maintain --mode " + mode,
		Start:      start,
		End:        time.Now(),
		WikiRoot:   g.WikiRoot,
		Provider:   g.Provider,
		Profile:    g.Profile,
		Status:     status,
		Error:      errMsg,
	}
	if err := runslog.AppendRun(entry); err != nil && g.Verbose {
		fmt.Fprintln(os.Stderr, "warn: write runs log:", err)
	}
}
