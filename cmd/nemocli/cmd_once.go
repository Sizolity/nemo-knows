package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/huic/nemo-knows/internal/wikimaint"
)

// cmdOnce is the systemd-timer entry point. It composes lint and
// maintain in that order; the maintain phase always runs even when lint
// reports errors so the operator gets a complete picture in one pass.
//
// The mode defaults to NEMO_ONCE_MODE (env), then to wikimaint.ModeReport
// when the env is unset. Both stages share the global config.
func cmdOnce(args []string, g *globalConfig) int {
	fs := flag.NewFlagSet("nemocli once", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	mode := fs.String("mode", "", "override NEMO_ONCE_MODE (report|safe|propose|auto)")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	chosen := strings.TrimSpace(*mode)
	if chosen == "" {
		chosen = strings.TrimSpace(os.Getenv("NEMO_ONCE_MODE"))
	}
	if chosen == "" {
		chosen = wikimaint.ModeReport
	}
	if g.Verbose {
		fmt.Fprintf(os.Stderr, "once: lint then maintain --mode=%s\n", chosen)
	}

	lintExit := cmdLint(nil, g)
	maintainExit := cmdMaintain([]string{"--mode", chosen}, g)

	if lintExit != 0 || maintainExit != 0 {
		// Surface the worst non-zero exit; lint precedes maintain so we
		// keep lint's exit code when it is non-zero, otherwise return
		// maintain's.
		if lintExit != 0 {
			return lintExit
		}
		return maintainExit
	}
	return 0
}
