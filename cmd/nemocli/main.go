// nemocli is the single-instance front-door for the nemo-knows wiki
// pipeline. It owns the run-directory lifecycle, the cross-process lock,
// the runs log, and the subcommand vocabulary; the actual draft / review
// / eval / apply stages reuse cmd/nemo via subprocess (heavy generation)
// and internal/* directly (deterministic gates).
//
// See docs/design and the v3 design subagent transcript for the contract
// that drives this binary. Anchor sections:
//
//   - §3.1 / §3.3 subcommand matrix
//   - §6 runDir lifecycle
//   - §7.4 secrets guard
//   - §8 runs log
//   - §9 config & env vars
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/huic/nemo-knows/internal/cliroot"
	"github.com/huic/nemo-knows/internal/runslog"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

// globalConfig is the post-parse view of nemocli-level flags / env.
type globalConfig struct {
	WikiRoot string // resolved absolute path; never has a trailing "wiki/"
	Provider string // empty means "use NEMO_MODEL_PROVIDER or DeepSeek"
	Profile  string // "stable" by default; passed through to cmd/nemo
	Verbose  bool
}

const defaultProfile = "stable"

// run is the testable inner entry point. It returns the desired process
// exit code; main() forwards it to os.Exit.
func run(args []string) int {
	if len(args) == 0 {
		printUsage(os.Stdout)
		return 0
	}
	switch args[0] {
	case "-h", "--help", "help":
		printUsage(os.Stdout)
		return 0
	case "-v", "--version", "version":
		fmt.Println("nemocli v0 (single-binary, design v3)")
		return 0
	}

	sub, subArgs, g, err := parseGlobalFlags(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "nemocli:", err)
		fmt.Fprintln(os.Stderr)
		printUsage(os.Stderr)
		return 2
	}

	// $WIKI_ROOT/.env is loaded for every subcommand because nearly all
	// of them need provider secrets to do anything useful. Errors are
	// non-fatal: missing .env is expected for fresh installs (init runs
	// before .env exists) and for non-ingest commands.
	if err := loadWikiEnv(g.WikiRoot); err != nil && g.Verbose {
		fmt.Fprintln(os.Stderr, "warn: load wiki env:", err)
	}

	// Schedule the runs-log auto-GC immediately. It runs in a goroutine
	// and is throttled to once per 24h via runs/.last-gc; if the
	// subcommand finishes before the goroutine wakes up, the goroutine
	// may simply die without effect (acceptable for V0).
	runslog.StartGCAsync(envIntDefault("NEMO_RUNS_RETENTION_DAYS", 30))

	switch sub {
	case "init":
		return cmdInit(subArgs, g)
	case "ingest":
		return cmdIngest(subArgs, g)
	case "lint":
		return cmdLint(subArgs, g)
	case "maintain":
		return cmdMaintain(subArgs, g)
	case "once":
		return cmdOnce(subArgs, g)
	case "logs":
		return cmdLogs(subArgs, g)
	case "gc":
		return cmdGC(subArgs, g)
	case "query":
		return cmdQueryStub(subArgs, g)
	case "serve":
		return cmdServeStub(subArgs, g)
	default:
		fmt.Fprintf(os.Stderr, "nemocli: unknown subcommand %q\n\n", sub)
		printUsage(os.Stderr)
		return 2
	}
}

// parseGlobalFlags consumes the global flag block that precedes the
// subcommand. It deliberately does NOT consume flags that appear after
// the subcommand — those belong to the subcommand-local FlagSet.
func parseGlobalFlags(args []string) (string, []string, *globalConfig, error) {
	fs := flag.NewFlagSet("nemocli", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	wikiRoot := fs.String("wiki-root", "", "wiki root path (default $NEMO_WIKI_ROOT or ~/.wiki)")
	provider := fs.String("provider", "", "model provider override: deepseek | llama")
	profile := fs.String("profile", defaultProfile, "generation profile: fast | stable | deep | fallback")
	verbose := fs.Bool("verbose", false, "extra stderr logging")
	fs.BoolVar(verbose, "v", false, "shorthand for -verbose")

	if err := fs.Parse(args); err != nil {
		return "", nil, nil, err
	}
	rest := fs.Args()
	if len(rest) == 0 {
		return "", nil, nil, fmt.Errorf("missing subcommand")
	}
	resolved, err := cliroot.ResolveWikiRoot(*wikiRoot)
	if err != nil {
		return "", nil, nil, fmt.Errorf("resolve --wiki-root: %w", err)
	}
	return rest[0], rest[1:], &globalConfig{
		WikiRoot: resolved,
		Provider: *provider,
		Profile:  *profile,
		Verbose:  *verbose,
	}, nil
}

func envIntDefault(key string, fallback int) int {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
