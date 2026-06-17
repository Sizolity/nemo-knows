package main

import (
	"time"

	"github.com/huic/nemo-knows/internal/rundir"
)

// newRunIDForCmd is a tiny shim so subcommands that don't open a runDir
// still get a stable run-id for their runs-log entries. It defers to
// rundir.NewRunID so the format matches ingest's anchor IDs.
func newRunIDForCmd(at time.Time) string {
	return rundir.NewRunID(at)
}
