package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/huic/nemo-knows/internal/apply"
	"github.com/huic/nemo-knows/internal/cliroot"
	"github.com/huic/nemo-knows/internal/config"
	"github.com/huic/nemo-knows/internal/evalharness"
	"github.com/huic/nemo-knows/internal/review"
	"github.com/huic/nemo-knows/internal/rundir"
	"github.com/huic/nemo-knows/internal/runslog"
)

// cmdIngest drives the end-to-end pipeline for a single source file.
//
// Architecture (v3 design §6.2):
//   - cmd/nemocli owns the runDir lifecycle, the flock, the runs log,
//     and stage sequencing.
//   - cmd/nemo (subprocess) generates the heavy LLM artifacts (bundle
//     source/ingest-plan, candidate drafts). nemocli cannot duplicate
//     that logic without modifying cmd/nemo, which the task forbids.
//   - internal/* (in-process) does the deterministic gates (review,
//     eval-bundle, eval-candidates, crosslink lint, apply).
//
// On success the runDir is removed; on failure it is renamed into
// /tmp/nemo-checkpoints/<run-id>/. Either way the runs log gets a
// section. Resume mode skips stages whose products already exist in
// the checkpoint and removes the checkpoint when the run completes.
func cmdIngest(args []string, g *globalConfig) int {
	fs := flag.NewFlagSet("nemocli ingest", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	force := fs.Bool("force", false, "allow re-apply even when wiki/log.md already has this ingest")
	resume := fs.String("resume", "", "resume from /tmp/nemo-checkpoints/<run-id>/")
	dryRun := fs.Bool("dry-run", false, "(V1) run pipeline but skip apply; currently a no-op stub")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *dryRun {
		fmt.Fprintln(os.Stderr, "nemocli ingest --dry-run: deferred to V1 (see v3 design §3.4)")
		return 0
	}
	if *resume == "" && fs.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: nemocli ingest <file>   or   nemocli ingest --resume <path>")
		return 2
	}

	if err := cliroot.RequireInitialized(g.WikiRoot); err != nil {
		fmt.Fprintln(os.Stderr, "ingest:", err)
		return 2
	}

	cfg, err := config.ForProfileWithProvider(g.Profile, g.Provider)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ingest:", err)
		return 2
	}
	if err := requireSecrets(cfg, g.WikiRoot); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	unlock, err := acquireFlock(IngestLockPath)
	if err != nil {
		// fail-fast: another writer already holds the lock
		fmt.Fprintln(os.Stderr, "ingest:", err)
		_ = runslog.AppendRun(runslog.RunEntry{
			RunID:      newRunIDForCmd(time.Now()),
			Subcommand: "ingest",
			Start:      time.Now(),
			End:        time.Now(),
			WikiRoot:   g.WikiRoot,
			Provider:   cfg.Provider,
			Profile:    cfg.Profile,
			Status:     "skipped",
			Error:      err.Error(),
		})
		return 1
	}
	defer unlock()

	if *resume != "" {
		return runIngestResume(*resume, cfg, g, *force)
	}
	return runIngestFresh(fs.Arg(0), cfg, g, *force)
}

// runIngestFresh implements the happy path for a brand-new source.
// Stage list mirrors the v3 design §6.2 sequence. Each stage failure
// short-circuits to the cleanup path which renames runDir into a
// checkpoint and records a "fail" runs entry.
func runIngestFresh(source string, cfg config.Config, g *globalConfig, force bool) int {
	if _, err := os.Stat(source); err != nil {
		fmt.Fprintf(os.Stderr, "ingest: source %s: %v\n", source, err)
		return 2
	}

	runDir, runID, err := rundir.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ingest:", err)
		return 1
	}
	bundleDir := filepath.Join(runDir, "bundle")
	if err := os.MkdirAll(bundleDir, 0o755); err != nil {
		_, _ = rundir.Close(false, runDir, runID)
		fmt.Fprintln(os.Stderr, "ingest:", err)
		return 1
	}

	start := time.Now()
	pipeErr := runPipeline(source, bundleDir, cfg, g, force)
	end := time.Now()

	checkpoint, closeErr := rundir.Close(pipeErr == nil, runDir, runID)
	entry := runslog.RunEntry{
		RunID:      runID,
		Subcommand: "ingest",
		Source:     source,
		Start:      start,
		End:        end,
		WikiRoot:   g.WikiRoot,
		Provider:   cfg.Provider,
		Profile:    cfg.Profile,
	}
	if pipeErr != nil {
		entry.Status = "fail"
		entry.Checkpoint = checkpoint
		entry.Error = pipeErr.Error()
	} else {
		entry.Status = "ok"
	}
	if logErr := runslog.AppendRun(entry); logErr != nil && g.Verbose {
		fmt.Fprintln(os.Stderr, "warn: write runs log:", logErr)
	}

	if pipeErr != nil {
		fmt.Fprintln(os.Stderr, "ingest:", pipeErr)
		if checkpoint != "" {
			fmt.Fprintf(os.Stderr, "checkpoint: %s\n", checkpoint)
			fmt.Fprintf(os.Stderr, "resume:     nemocli ingest --resume %s\n", checkpoint)
		}
		return 1
	}
	if closeErr != nil && g.Verbose {
		fmt.Fprintln(os.Stderr, "warn: close run dir:", closeErr)
	}
	fmt.Fprintf(os.Stderr, "ingest: ok (run-id %s)\n", runID)
	return 0
}

// runIngestResume picks up where a previous failed run left off. The
// resume path is treated as the bundleDir directly — that is the
// directory shape v3 §6.5 documents and what cmd/nemo's resume mode
// already understands.
func runIngestResume(path string, cfg config.Config, g *globalConfig, force bool) int {
	if _, err := os.Stat(filepath.Join(path, "bundle")); err == nil {
		path = filepath.Join(path, "bundle")
	}
	if _, err := os.Stat(filepath.Join(path, "source.md")); err != nil {
		fmt.Fprintf(os.Stderr, "ingest: resume %s: bundle/source.md not found (%v)\n", path, err)
		return 2
	}
	bundleDir := path
	runID := rundir.NewRunID(time.Now())
	start := time.Now()

	pipeErr := runPipelineFromBundle(bundleDir, cfg, g, force)
	end := time.Now()

	entry := runslog.RunEntry{
		RunID:      runID,
		Subcommand: "ingest (resume)",
		Source:     bundleDir,
		Start:      start,
		End:        end,
		WikiRoot:   g.WikiRoot,
		Provider:   cfg.Provider,
		Profile:    cfg.Profile,
	}
	if pipeErr != nil {
		entry.Status = "fail"
		entry.Checkpoint = bundleDir
		entry.Error = pipeErr.Error()
	} else {
		entry.Status = "ok"
	}
	if logErr := runslog.AppendRun(entry); logErr != nil && g.Verbose {
		fmt.Fprintln(os.Stderr, "warn: write runs log:", logErr)
	}

	if pipeErr != nil {
		fmt.Fprintln(os.Stderr, "ingest (resume):", pipeErr)
		return 1
	}

	// Resume success → clean the checkpoint (§6.5: "resume self-cleanup
	// is not GC; it just means the run finally finished").
	resumeRoot := bundleDir
	if filepath.Base(resumeRoot) == "bundle" {
		resumeRoot = filepath.Dir(resumeRoot)
	}
	if err := os.RemoveAll(resumeRoot); err != nil && g.Verbose {
		fmt.Fprintf(os.Stderr, "warn: remove resumed checkpoint %s: %v\n", resumeRoot, err)
	}
	fmt.Fprintf(os.Stderr, "ingest (resume): ok (run-id %s)\n", runID)
	return 0
}

// runPipeline executes the full source → apply flow for a fresh ingest.
//
// force is forwarded all the way to apply.ApplyApproved so a user can
// re-run an ingest that previously logged a wiki/log.md entry.
func runPipeline(source, bundleDir string, cfg config.Config, g *globalConfig, force bool) error {
	if err := stageBundleViaNemo(source, bundleDir, cfg, g); err != nil {
		return fmt.Errorf("stage=bundle: %w", err)
	}
	return runPipelineFromBundle(bundleDir, cfg, g, force)
}

// runPipelineFromBundle covers stages 2-7. Resume mode reuses this
// directly because the bundle directory layout is identical.
func runPipelineFromBundle(bundleDir string, cfg config.Config, g *globalConfig, force bool) error {
	if !fileExists(filepath.Join(bundleDir, "apply-plan.md")) {
		if err := stageReview(bundleDir); err != nil {
			return fmt.Errorf("stage=review: %w", err)
		}
	}
	if !fileExists(filepath.Join(bundleDir, "scores.json")) {
		if err := stageEvalBundle(bundleDir); err != nil {
			return fmt.Errorf("stage=eval-bundle: %w", err)
		}
	}
	if !hasCandidateDrafts(bundleDir) {
		if err := stageGenerateCandidatesViaNemo(bundleDir, cfg, g); err != nil {
			return fmt.Errorf("stage=generate-candidates: %w", err)
		}
	}
	if err := stageEvalCandidates(g.WikiRoot, bundleDir); err != nil {
		return fmt.Errorf("stage=eval-candidates: %w", err)
	}
	if err := stageCrosslinkLint(g.WikiRoot, bundleDir); err != nil {
		return fmt.Errorf("stage=crosslink: %w", err)
	}
	if err := stageApply(g.WikiRoot, bundleDir, force); err != nil {
		return fmt.Errorf("stage=apply: %w", err)
	}
	return nil
}

// --- stage implementations ---------------------------------------------------

// stageBundleViaNemo runs cmd/nemo to produce bundle/source.md and
// bundle/ingest-plan.md. We subprocess instead of duplicating
// cmd/nemo's runBundle / runChunkedBundle (which together exceed 500
// lines) — the task explicitly forbids modifying cmd/nemo, and
// re-implementing the same logic would violate the "do not rewrite the
// pipeline" rule.
func stageBundleViaNemo(source, bundleDir string, cfg config.Config, g *globalConfig) error {
	bin, err := findNemoBinary()
	if err != nil {
		return err
	}
	args := []string{
		"-provider", cfg.Provider,
		"-profile", cfg.Profile,
		"-source", source,
		"-bundle-dir", bundleDir,
	}
	if g.Verbose {
		fmt.Fprintf(os.Stderr, "ingest: bundle via %s %v\n", bin, args)
	}
	cmd := exec.Command(bin, args...)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func stageReview(bundleDir string) error {
	plan, err := review.ReviewBundle(bundleDir)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(bundleDir, "apply-plan.md"), []byte(plan), 0o644)
}

func stageEvalBundle(bundleDir string) error {
	result, err := evalharness.EvaluateBundle(bundleDir)
	if err != nil {
		return err
	}
	body, err := encodeJSON(result)
	if err != nil {
		return err
	}
	// apply.requirePassingEval reads <bundle>/scores.json as the
	// approval gate; write it where apply expects to find it.
	return os.WriteFile(filepath.Join(bundleDir, "scores.json"), body, 0o644)
}

func stageGenerateCandidatesViaNemo(bundleDir string, cfg config.Config, g *globalConfig) error {
	bin, err := findNemoBinary()
	if err != nil {
		return err
	}
	args := []string{
		"-provider", cfg.Provider,
		"-profile", cfg.Profile,
		"-generate-candidates", bundleDir,
	}
	if g.Verbose {
		fmt.Fprintf(os.Stderr, "ingest: generate-candidates via %s %v\n", bin, args)
	}
	cmd := exec.Command(bin, args...)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func stageEvalCandidates(wikiRoot, bundleDir string) error {
	_, err := evalharness.EvaluateCandidatesWithRoot(wikiRoot, bundleDir)
	return err
}

func stageCrosslinkLint(wikiRoot, bundleDir string) error {
	result, err := evalharness.EvaluateBundleCrosslinks(wikiRoot, bundleDir)
	if err != nil {
		return err
	}
	// crosslink lint is best-effort context for review; we do not block
	// apply on it unless apply's own safety gate catches the same issue.
	// The result is still serialized for forensic value when failures
	// land in a checkpoint.
	body, err := encodeJSON(result)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(bundleDir, "bundle-crosslinks.json"), body, 0o644)
}

func stageApply(wikiRoot, bundleDir string, force bool) error {
	_, err := apply.ApplyApproved(wikiRoot, bundleDir, apply.Options{Approve: true, Force: force})
	return err
}

// --- helpers ----------------------------------------------------------------

func fileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func hasCandidateDrafts(bundleDir string) bool {
	root := filepath.Join(bundleDir, "candidates", "wiki")
	found := false
	_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			found = true
			return filepath.SkipAll
		}
		return nil
	})
	return found
}
