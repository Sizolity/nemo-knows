package main

import (
	"fmt"
	"os"
)

// cmdQueryStub keeps the V1 query verb wired into the dispatcher so the
// design's §3.1 lock — "never add new top-level verbs after V0" — is
// respected. Calling it today simply explains where the implementation
// is parked and exits 0 so scripts don't accidentally fail.
func cmdQueryStub(_ []string, _ *globalConfig) int {
	fmt.Fprintln(os.Stderr,
		"nemocli query: deferred to V1 (see v3 design §3.1 / §12)")
	return 0
}

// cmdServeStub does the same for the V1 daemon entry point. V1 will
// flesh this out into the watcher + ticker described in §4.3.
func cmdServeStub(_ []string, _ *globalConfig) int {
	fmt.Fprintln(os.Stderr,
		"nemocli serve: deferred to V1 (see v3 design §3.1 / §4.3)")
	return 0
}
