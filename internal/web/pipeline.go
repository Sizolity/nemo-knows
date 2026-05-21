package web

import "github.com/huic/nemo-knows/internal/config"

// Pipeline is the subset of nemo's ingest pipeline that the web console
// invokes on behalf of a user. It is intentionally narrow: the web layer only
// needs to start a bundle and review it; deeper stages (apply, eval, etc.)
// remain in the dedicated nemo CLI.
//
// Implementations are provided by the caller. cmd/nemo wires up an in-process
// adapter; cmd/nemo-web uses a subprocess adapter that shells out to the
// existing nemo binary.
type Pipeline interface {
	// RunBundle generates source.md + ingest-plan.md drafts under bundleDir
	// for the given raw source, using cfg for model/profile selection.
	RunBundle(source string, bundleDir string, cfg config.Config) error

	// RunReviewBundle produces a deterministic apply-plan.md review at out
	// for the bundle at bundleDir.
	RunReviewBundle(bundleDir string, out string) error
}
