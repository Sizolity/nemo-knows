// Command nemo-web is the dedicated launcher binary for the nemo-knows web
// console. It binds the internal/web package to a local address and shells
// out to the existing nemo CLI for the heavy ingest pipeline so this binary
// stays small and decoupled from the pipeline source tree.
//
// Usage:
//
//	nemo-web                                  # default 127.0.0.1:8787
//	nemo-web -addr 127.0.0.1:9876
//	nemo-web -nemo-binary .bin/nemo
//	NEMO_BINARY=.bin/nemo nemo-web            # same as -nemo-binary
//
// The ingest pipeline endpoint (/run) only works when the nemo binary
// referenced by -nemo-binary exists and is executable. Otherwise the rest of
// the console (browsing, graph, /build save) still works in degraded mode.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/huic/nemo-knows/internal/config"
	"github.com/huic/nemo-knows/internal/web"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8787", "address to listen on")
	nemoBinary := flag.String("nemo-binary", defaultNemoBinary(), "path to the nemo CLI used to run the ingest pipeline (set NEMO_BINARY to override)")
	flag.Parse()

	cfg, err := config.ForProfile("")
	if err != nil {
		fmt.Fprintln(os.Stderr, "config:", err)
		os.Exit(1)
	}

	pipeline := newExecPipeline(*nemoBinary)
	if err := web.Run(*addr, cfg, pipeline); err != nil {
		fmt.Fprintln(os.Stderr, "server:", err)
		os.Exit(1)
	}
}

func defaultNemoBinary() string {
	if v := os.Getenv("NEMO_BINARY"); v != "" {
		return v
	}
	// Prefer the local build artifact under .bin/ if available; otherwise
	// fall back to whatever `nemo` resolves to in PATH at exec time.
	if _, err := os.Stat(".bin/nemo"); err == nil {
		return ".bin/nemo"
	}
	return "nemo"
}

// execPipeline implements web.Pipeline by invoking the nemo CLI as a
// subprocess. This keeps cmd/nemo-web independent of the pipeline source code
// (which lives in package main of cmd/nemo) while still letting /run work in
// the standalone binary.
type execPipeline struct {
	binary string
}

func newExecPipeline(binary string) *execPipeline {
	return &execPipeline{binary: binary}
}

func (p *execPipeline) RunBundle(source string, bundleDir string, cfg config.Config) error {
	args := []string{
		"-provider", cfg.Provider,
		"-profile", cfg.Profile,
		"-source", source,
		"-bundle-dir", bundleDir,
	}
	return p.run(args)
}

func (p *execPipeline) RunReviewBundle(bundleDir string, out string) error {
	args := []string{
		"-review-bundle", bundleDir,
		"-out", out,
	}
	return p.run(args)
}

func (p *execPipeline) run(args []string) error {
	cmd := exec.Command(p.binary, args...)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %v: %w", p.binary, args, err)
	}
	return nil
}
