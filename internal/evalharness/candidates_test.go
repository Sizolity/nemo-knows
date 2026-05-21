package evalharness

import (
	"path/filepath"
	"testing"
)

func TestEvaluateCandidatesScoresGeneratedDrafts(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source\n\nThe LLM Wiki pattern uses an LLM agent to maintain a persistent, structured wiki from raw documents.\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n"+
		"- `wiki/concepts/llm-maintenance-pattern.md` — create new page.\n"+
		"- `wiki/topics/persistent-wiki-architecture.md` — create new page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "concepts", "llm-maintenance-pattern.md"), `---
title: LLM Maintenance Pattern
kind: concept
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# LLM Maintenance Pattern

The [[LLM Wiki]] maintenance pattern describes how an LLM keeps a durable wiki current across ingest, query, and lint operations. It turns source summaries into maintained pages rather than treating every answer as a temporary response.

This concept is useful because the human keeps control over sources while the model handles structure, links, summaries, and bookkeeping.
`)
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "topics", "persistent-wiki-architecture.md"), `---
title: Persistent Wiki Architecture
kind: topic
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# Persistent Wiki Architecture

Persistent wiki architecture uses [[Markdown]] pages as the durable knowledge layer around raw sources. Source pages capture individual inputs, concept pages define reusable ideas, and topic pages connect those ideas into cross-cutting explanations.

The architecture stays auditable because raw sources remain immutable while accepted wiki edits are explicit and logged.
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	if result.Scores.Overall != "pass" {
		t.Fatalf("overall score = %q, want pass; trace=%v", result.Scores.Overall, result.Trace)
	}
	if len(result.Candidates) != 2 {
		t.Fatalf("candidate count = %d, want 2", len(result.Candidates))
	}
}

func TestEvaluateCandidatesFlagsMissingSourcesAndShortDraft(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/concepts/llm-maintenance-pattern.md` — create new page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "concepts", "llm-maintenance-pattern.md"), `---
title: LLM Maintenance Pattern
kind: concept
confidence: medium
---

# LLM Maintenance Pattern

Too short.
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	if result.Scores.Overall != "fail" {
		t.Fatalf("overall score = %q, want fail", result.Scores.Overall)
	}
	if got := result.Candidates[0].Scores.Sources; got != "fail" {
		t.Fatalf("sources score = %q, want fail", got)
	}
	if got := result.Candidates[0].Scores.Length; got != "fail" {
		t.Fatalf("length score = %q, want fail", got)
	}
}

func TestEvaluateCandidatesAllowsMissingWikilinksWhenDraftIsOtherwiseReviewable(t *testing.T) {
	dir := t.TempDir()
	sourceLine := "The LLM Wiki pattern uses an LLM agent to maintain a persistent, structured wiki from raw documents."
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source\n\n"+sourceLine+"\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/topics/persistent-wiki-architecture.md` — create new page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "topics", "persistent-wiki-architecture.md"), `---
title: Persistent Wiki Architecture
kind: topic
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# Persistent Wiki Architecture

The LLM Wiki pattern uses an LLM agent to maintain a persistent, structured wiki from raw documents.

The LLM Wiki pattern uses an LLM agent to maintain a persistent, structured wiki from raw documents.

The LLM Wiki pattern uses an LLM agent to maintain a persistent, structured wiki from raw documents.
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	if result.Scores.Overall != "borderline" {
		t.Fatalf("overall score = %q, want borderline; trace=%v", result.Scores.Overall, result.Trace)
	}
	if got := result.Candidates[0].Scores.Wikilinks; got != "pass" {
		t.Fatalf("wikilinks score = %q, want pass", got)
	}
	if got := result.Candidates[0].Scores.Originality; got != "borderline" {
		t.Fatalf("originality score = %q, want borderline", got)
	}
}

func TestEvaluateCandidatesFlagsNearCopiedSourceProse(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/dom.md\n---\n\n# Source\n\nThe event dispatch algorithm collects an ordered path, invokes capturing listeners before target listeners, and then invokes bubbling listeners while retargeting nodes across shadow boundaries.\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/concepts/event-dispatch.md` — create new page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "concepts", "event-dispatch.md"), `---
title: Event Dispatch
kind: concept
sources:
  - source.md
  - raw/dom.md
confidence: medium
---

# Event Dispatch

The event dispatch algorithm collects an ordered path, invokes capturing listeners before target listeners, and then invokes bubbling listeners while retargeting nodes across shadow roots.

This event dispatch description preserves the same structure and vocabulary as the source while swapping only a few verbs, which should still be treated as too close.
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	if got := result.Candidates[0].Scores.Originality; got != "borderline" {
		t.Fatalf("originality score = %q, want borderline; trace=%v", got, result.Candidates[0].Trace)
	}
}

func TestEvaluateCandidatesFlagsWeakTitleBodyAlignment(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/dom.md\n---\n\n# Source\n\nThe DOM standard defines trees, events, mutation observers, and legacy browser interfaces.\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/topics/browser-compatibility-matrix.md` — create new page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "topics", "browser-compatibility-matrix.md"), `---
title: Browser Compatibility Matrix
kind: topic
sources:
  - source.md
  - raw/dom.md
confidence: medium
---

# Browser Compatibility Matrix

The DOM standard defines a tree-shaped model for documents, node relationships, event dispatch, mutation observers, and legacy interfaces preserved for old content.

This draft explains event propagation, shadow-boundary retargeting, mutation callbacks, range traversal, and namespace limitations without giving any implementation support table.
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	if got := result.Candidates[0].Scores.TitleAlignment; got != "borderline" {
		t.Fatalf("title alignment score = %q, want borderline; trace=%v", got, result.Candidates[0].Trace)
	}
	if result.Scores.Overall != "borderline" {
		t.Fatalf("overall score = %q, want borderline", result.Scores.Overall)
	}
}

func TestEvaluateCandidatesFlagsUnclosedCodeFence(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/moby-dick.md\n---\n\n# Source\n\nMoby-Dick includes long descriptions of whaling practices aboard the Pequod.\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/topics/whaling-practices.md` — create new page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "topics", "whaling-practices.md"), `---
title: Whaling Practices
kind: topic
sources:
  - source.md
  - raw/moby-dick.md
confidence: medium
---

# Whaling Practices

Moby-Dick describes the whaling voyage as a sequence of practical work aboard the Pequod, including boat lowering, pursuit, harpooning, towing, cutting in, and trying out. The details present whaling as coordinated industrial labor rather than only symbolic adventure.
`+"```"+`
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	if got := result.Candidates[0].Scores.Markdown; got != "fail" {
		t.Fatalf("markdown score = %q, want fail; trace=%v", got, result.Candidates[0].Trace)
	}
	if result.Scores.Overall != "fail" {
		t.Fatalf("overall score = %q, want fail", result.Scores.Overall)
	}
}

func TestEvaluateCandidatesMarksMissingWikilinkTargetsBorderline(t *testing.T) {
	root := t.TempDir()
	bundle := filepath.Join(root, "drafts", "bundle")
	writeFile(t, filepath.Join(root, "wiki", "concepts", "known.md"), `---
title: Known
kind: concept
sources:
  - raw/source.md
confidence: medium
---

# Known
`)
	writeFile(t, filepath.Join(bundle, "source.md"), "---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source\n\nA source about wiki maintenance.\n")
	writeFile(t, filepath.Join(bundle, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/concepts/llm-maintenance-pattern.md` — create new page.\n")
	writeFile(t, filepath.Join(bundle, "candidates", "wiki", "concepts", "llm-maintenance-pattern.md"), `---
title: LLM Maintenance Pattern
kind: concept
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# LLM Maintenance Pattern

The [[Known]] page is allowed, but [[Missing Target]] is not. This draft has enough words to pass the basic length check while still proving that missing wikilink targets are surfaced before apply.
`)

	result, err := EvaluateCandidatesWithRoot(root, bundle)
	if err != nil {
		t.Fatalf("EvaluateCandidatesWithRoot returned error: %v", err)
	}
	if got := result.Candidates[0].Scores.Wikilinks; got != "borderline" {
		t.Fatalf("wikilinks score = %q, want borderline; trace=%v", got, result.Candidates[0].Trace)
	}
	if result.Scores.Overall != "borderline" {
		t.Fatalf("overall score = %q, want borderline", result.Scores.Overall)
	}
}

func TestEvaluateCandidatesMarksWeakExistingWikilinkTargetsBorderline(t *testing.T) {
	root := t.TempDir()
	bundle := filepath.Join(root, "drafts", "bundle")
	writeFile(t, filepath.Join(root, "wiki", "concepts", "unrelated.md"), `---
title: Unrelated
kind: concept
sources:
  - raw/other.md
confidence: medium
---

# Unrelated
`)
	writeFile(t, filepath.Join(bundle, "source.md"), "---\nkind: source\nsources:\n  - raw/sqlite-wal.md\n---\n\n# Source\n\nSQLite WAL uses checkpointing to copy committed frames back into the main database.\n")
	writeFile(t, filepath.Join(bundle, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/concepts/checkpointing.md` — create new page.\n")
	writeFile(t, filepath.Join(bundle, "candidates", "wiki", "concepts", "checkpointing.md"), `---
title: Checkpointing
kind: concept
sources:
  - source.md
  - raw/sqlite-wal.md
confidence: medium
---

# Checkpointing

SQLite WAL checkpointing copies committed frames back into the main database while readers keep using a stable snapshot. This draft has enough words to pass the length gate while linking to [[Unrelated]], which exists but is not supported by the source or reviewed candidates.
`)

	result, err := EvaluateCandidatesWithRoot(root, bundle)
	if err != nil {
		t.Fatalf("EvaluateCandidatesWithRoot returned error: %v", err)
	}
	if got := result.Candidates[0].Scores.Wikilinks; got != "borderline" {
		t.Fatalf("wikilinks score = %q, want borderline; trace=%v", got, result.Candidates[0].Trace)
	}
	if result.Scores.Overall != "borderline" {
		t.Fatalf("overall score = %q, want borderline", result.Scores.Overall)
	}
}

func TestEvaluateBundleCrosslinksReportsMissingAndZeroInbound(t *testing.T) {
	root := t.TempDir()
	bundle := filepath.Join(root, "drafts", "bundle")
	writeFile(t, filepath.Join(root, "wiki", "concepts", "known.md"), "# Known\n")
	writeFile(t, filepath.Join(bundle, "apply-plan.md"), "## Candidate Changes\n\n"+
		"- `wiki/concepts/first.md` — create new page.\n"+
		"- `wiki/topics/second.md` — create new page.\n")
	writeFile(t, filepath.Join(bundle, "candidates", "wiki", "concepts", "first.md"), `# First

This page links to [[second]] and [[Missing Target]].
`)
	writeFile(t, filepath.Join(bundle, "candidates", "wiki", "topics", "second.md"), `# Second

This page links to [[known]].
`)

	result, err := EvaluateBundleCrosslinks(root, bundle)
	if err != nil {
		t.Fatalf("EvaluateBundleCrosslinks returned error: %v", err)
	}
	if len(result.Graph) != 2 {
		t.Fatalf("graph edge count = %d, want 2: %#v", len(result.Graph), result.Graph)
	}
	hasMissing := false
	hasZeroInbound := false
	for _, issue := range result.Issues {
		if issue.Code == "missing-target" {
			hasMissing = true
		}
		if issue.Code == "zero-inbound" && issue.Path == "wiki/concepts/first.md" {
			hasZeroInbound = true
		}
	}
	if !hasMissing || !hasZeroInbound {
		t.Fatalf("expected missing-target and zero-inbound issues, got %#v", result.Issues)
	}
}
