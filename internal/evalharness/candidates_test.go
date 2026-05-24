package evalharness

import (
	"fmt"
	"path/filepath"
	"strings"
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

This concept is useful because the human keeps control over sources while the model handles structure, links, summaries, and bookkeeping. The agent reads raw inputs, produces structured Markdown pages, and logs every operation it performs.

Each ingest cycle creates or updates entity, concept, and topic pages so that the wiki grows as a compounding reference rather than a collection of one-off chat answers. Frontmatter metadata tracks provenance, confidence, and freshness.

The lint workflow finds orphan pages, broken links, and contradictions so that quality improves over time without requiring the human to inspect every page manually.
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

The architecture stays auditable because raw sources remain immutable while accepted wiki edits are explicit and logged. Every change traces back to a raw document through the sources list in each page's frontmatter.

Index and log files act as structural anchors: the index catalogues every page, and the append-only log provides a greppable audit trail of what changed and when.

Compared to a flat document store, this layered design lets multiple sources contribute to the same concept page, making it possible to synthesise knowledge across ingests.
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

The LLM maintenance pattern turns raw documents into durable wiki pages through a structured ingest, query, and lint cycle that a model runs on behalf of a human maintainer.

Each ingest produces source, concept, and topic pages with provenance metadata so that every claim traces back to a specific raw input. The human reviews proposed changes before they are applied to the wiki.

The lint workflow periodically scans for broken links, orphan pages, and contradictions, surfacing repair suggestions that the human approves or rejects.
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

A checkpoint transfers pages from the WAL file into the original database file by reading each committed frame and writing it to the corresponding page offset. The operation can proceed concurrently with readers because they continue using the WAL snapshot they started with.

Three checkpoint modes exist: passive, full, and restart. Passive checkpoints only transfer frames that no reader currently needs, while full checkpoints block new readers until all frames are transferred.

The truncate mode additionally resets the WAL file to zero length after a successful full checkpoint, reclaiming disk space.
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

func TestCandidateDepthFailsForStubWithOnlyHeadersAndLists(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/stub.md\n---\n\n# Source\n\nA reference source.\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/concepts/stub-concept.md` — create new page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "concepts", "stub-concept.md"), `---
title: Stub Concept
kind: concept
sources:
  - source.md
  - raw/stub.md
confidence: medium
---

# Stub Concept

- Point one
- Point two
- Point three
- Point four
- Point five
- Point six
- Point seven
- Point eight about the concept
- Point nine about the concept
- Point ten about the concept
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	if got := result.Candidates[0].Scores.Depth; got != "fail" {
		t.Fatalf("depth score = %q, want fail; trace=%v", got, result.Candidates[0].Trace)
	}
}

func TestCandidateDepthBorderlineForThinProse(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/thin.md\n---\n\n# Source\n\nA reference about thin topics.\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/topics/thin-topic.md` — create new page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "topics", "thin-topic.md"), `---
title: Thin Topic
kind: topic
sources:
  - source.md
  - raw/thin.md
confidence: medium
---

# Thin Topic

This topic covers one specific aspect of the subject matter in enough detail to be considered a reviewable reference page.

A second sentence provides additional context about why this topic matters in the broader context of the wiki.
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	if got := result.Candidates[0].Scores.Depth; got != "borderline" {
		t.Fatalf("depth score = %q, want borderline; trace=%v", got, result.Candidates[0].Trace)
	}
}

func TestCandidateScopeBorderlineForSourceMirror(t *testing.T) {
	dir := t.TempDir()
	sourceDraft := `---
kind: source
sources:
  - raw/wal.md
---

# Source

SQLite WAL mode allows concurrent readers by appending changes to a separate write-ahead log file rather than modifying the database directly. Readers see a consistent snapshot while writers append new frames. Checkpointing transfers committed frames back into the main database file. The WAL file grows as transactions commit new frames and shrinks when a checkpoint completes successfully. Three checkpoint modes exist: passive, full, and restart, each with different concurrency trade-offs.`

	writeFile(t, filepath.Join(dir, "source.md"), sourceDraft)
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n- `wiki/concepts/wal-mode.md` — create new page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "concepts", "wal-mode.md"), `---
title: WAL Mode
kind: concept
sources:
  - source.md
  - raw/wal.md
confidence: medium
---

# WAL Mode

SQLite WAL mode allows concurrent readers by appending changes to a separate write-ahead log file rather than modifying the database directly. Readers see a consistent snapshot while writers append new frames to the log file.

Checkpointing transfers committed frames back into the main database file when no readers need them. The WAL file grows as transactions commit new frames and shrinks when checkpoints complete.

Three checkpoint modes exist: passive, full, and restart, each with different concurrency trade-offs for read and write operations against the database.

The WAL file grows as transactions commit and shrinks when a checkpoint transfers the committed frames back into the main database successfully.
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	if got := result.Candidates[0].Scores.Scope; got != "borderline" {
		t.Fatalf("scope score = %q, want borderline; trace=%v", got, result.Candidates[0].Trace)
	}
}

func TestCandidateRedundancyFlagsPairWithHighOverlap(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/wal.md\n---\n\n# Source\n\nSQLite WAL is a journaling mode.\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n"+
		"- `wiki/concepts/wal-overview.md` — create new page.\n"+
		"- `wiki/topics/wal-details.md` — create new page.\n")

	sharedProse := `SQLite WAL mode uses a write-ahead log to record changes before they reach the main database file. The WAL file grows as transactions commit and shrinks when checkpoints transfer frames back. Readers see consistent snapshots without blocking writers because each reader uses the WAL index to find its starting point. Checkpoint modes include passive, full, restart, and truncate, each offering different trade-offs between concurrency and disk reclamation.`

	writeFile(t, filepath.Join(dir, "candidates", "wiki", "concepts", "wal-overview.md"), fmt.Sprintf(`---
title: WAL Overview
kind: concept
sources:
  - source.md
  - raw/wal.md
confidence: medium
---

# WAL Overview

%s
`, sharedProse))
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "topics", "wal-details.md"), fmt.Sprintf(`---
title: WAL Details
kind: topic
sources:
  - source.md
  - raw/wal.md
confidence: medium
---

# WAL Details

%s
`, sharedProse))

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	for _, candidate := range result.Candidates {
		if candidate.Scores.Redundancy == "pass" {
			t.Fatalf("redundancy score for %s = pass, want borderline or fail; trace=%v", candidate.Path, candidate.Trace)
		}
	}
}

func TestCandidateRedundancyPassesForDistinctCandidates(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/moby.md\n---\n\n# Source\n\nMoby-Dick tells the story of whaling.\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "## Candidate Changes\n\n"+
		"- `wiki/topics/ahab.md` — create new page.\n"+
		"- `wiki/topics/whiteness.md` — create new page.\n")

	writeFile(t, filepath.Join(dir, "candidates", "wiki", "topics", "ahab.md"), `---
title: Ahab
kind: topic
sources:
  - source.md
  - raw/moby.md
confidence: medium
---

# Ahab

Captain Ahab commands the Pequod on a monomaniacal quest to kill the white whale that took his leg. His obsession drives the narrative forward and shapes the fates of every crew member aboard the ship.

Ahab's character draws from Shakespearean tragedy: he is eloquent, self-aware, and doomed. His soliloquies reveal a mind that understands its own destruction yet cannot turn from its course.

The crew follows Ahab partly from contractual obligation and partly from the force of his personality, which overwhelms the pragmatic objections raised by Starbuck and the quiet despair of Pip.

Ahab's ivory leg, forged from whale bone, serves as a physical reminder of his unfinished conflict with Moby Dick and the broader natural world he seeks to dominate.
`)
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "topics", "whiteness.md"), `---
title: Whiteness
kind: topic
sources:
  - source.md
  - raw/moby.md
confidence: medium
---

# Whiteness

The whiteness of Moby Dick is the subject of an entire chapter in which Ishmael meditates on why the color white inspires both reverence and terror across cultures and natural phenomena.

Ishmael catalogues white objects that terrify despite their color: the polar bear, the albatross, the White Mountains in mist, and albino humans shunned by their communities.

The chapter argues that whiteness amplifies whatever quality an object already possesses, making the sacred more sacred and the horrifying more horrifying through a blankness that the mind fills with its own projections.

Applied to Moby Dick, whiteness becomes the screen onto which Ahab projects his revenge, Starbuck his dread, and Ishmael his philosophical uncertainty about the universe.
`)

	result, err := EvaluateCandidates(dir)
	if err != nil {
		t.Fatalf("EvaluateCandidates returned error: %v", err)
	}
	for _, candidate := range result.Candidates {
		if candidate.Scores.Redundancy != "pass" {
			t.Fatalf("redundancy score for %s = %q, want pass; trace=%v", candidate.Path, candidate.Scores.Redundancy, candidate.Trace)
		}
	}
}

func TestReviewCrosslinksProducesRecommendationForZeroInbound(t *testing.T) {
	crosslinks := CrosslinkResult{
		Bundle: "drafts/test-bundle",
		Issues: []CrosslinkIssue{
			{Path: "wiki/topics/orphan-page.md", Code: "zero-inbound", Message: "candidate has no inbound links from sibling candidates"},
			{Path: "wiki/topics/broken-link.md", Code: "missing-target", Message: "wikilink target does not exist: missing-ref"},
		},
	}
	items := ReviewCrosslinks(crosslinks)
	if len(items) != 2 {
		t.Fatalf("expected 2 crosslink review items, got %d", len(items))
	}
	if items[0].Severity != "borderline" {
		t.Fatalf("zero-inbound severity = %q, want borderline", items[0].Severity)
	}
	if items[1].Severity != "borderline" {
		t.Fatalf("missing-target severity = %q, want borderline", items[1].Severity)
	}
}

func TestRenderCandidateReviewIncludesCrosslinkItems(t *testing.T) {
	review := CandidateReview{
		Bundle:  "drafts/test-bundle",
		Overall: "pass",
	}
	crosslinkItems := []CandidateReviewItem{
		{
			Path:           "wiki/topics/orphan.md",
			Severity:       "borderline",
			Problem:        "crosslink: candidate has no inbound links",
			Recommendation: "Add a wikilink from at least one sibling candidate.",
		},
	}
	rendered := RenderCandidateReview(review, crosslinkItems...)
	if !strings.Contains(rendered, "needs-review") {
		t.Fatal("expected overall to be upgraded to needs-review when crosslink items exist")
	}
	if !strings.Contains(rendered, "orphan.md") {
		t.Fatal("expected crosslink item to appear in rendered review")
	}
	if !strings.Contains(rendered, "items: `1`") {
		t.Fatalf("expected item count 1 in rendered review, got: %s", rendered)
	}
}
