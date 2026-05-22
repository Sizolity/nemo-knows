package evalharness

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEvaluateBundleScoresReviewArtifacts(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\n---\n# Source\n")
	writeFile(t, filepath.Join(dir, "ingest-plan.md"), "---\nkind: topic\n---\n# Ingest Plan\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "# Reviewed Ingest Apply Plan\n\n"+
		"This is a review artifact. Do not apply this plan automatically.\n\n"+
		"## Validation\n\n"+
		"- [x] `source.md` has YAML frontmatter\n"+
		"- [x] `ingest-plan.md` has YAML frontmatter\n"+
		"- [x] `source.md` frontmatter `kind` is `source`\n"+
		"- [x] `ingest-plan.md` frontmatter `kind` is `topic`\n"+
		"- [x] `source.md` includes required section `What It Is`\n"+
		"- [x] `source.md` includes required section `Summary`\n"+
		"- [x] `source.md` includes required section `Key Claims`\n"+
		"- [x] `source.md` includes required section `Suggested Links`\n"+
		"- [x] `ingest-plan.md` includes required section `Source Summary`\n"+
		"- [x] `ingest-plan.md` includes required section `Candidate Wiki Pages`\n"+
		"- [x] `ingest-plan.md` includes required section `Suggested Links`\n"+
		"- [x] `ingest-plan.md` includes required section `Review Checklist`\n\n"+
		"## Candidate Changes\n\n"+
		"- `wiki/sources/llm-wiki-pattern.md` — create new page; possible duplicate of `wiki/sources/llm-wiki.md`.\n"+
		"- `wiki/concepts/persistent-wiki.md` — update existing page.\n")
	writeFile(t, filepath.Join(dir, "candidates", "wiki", "concepts", "persistent-wiki.md"), "---\ntitle: Persistent Wiki\nkind: concept\nsources:\n  - source.md\n  - raw/source.md\nconfidence: medium\n---\n\n# Persistent Wiki\n\nA persistent wiki stores reusable source-backed notes.\n")

	result, err := EvaluateBundle(dir)
	if err != nil {
		t.Fatalf("EvaluateBundle returned error: %v", err)
	}

	if result.Scores.Schema != "pass" {
		t.Fatalf("schema score = %q, want pass", result.Scores.Schema)
	}
	if result.Scores.WikiSafety != "pass" {
		t.Fatalf("wiki safety score = %q, want pass", result.Scores.WikiSafety)
	}
	if result.Scores.CandidatePaths != "pass" {
		t.Fatalf("candidate paths score = %q, want pass", result.Scores.CandidatePaths)
	}
	if result.Scores.DuplicateDetection != "pass" {
		t.Fatalf("duplicate detection score = %q, want pass", result.Scores.DuplicateDetection)
	}
	if result.Scores.SourceCompleteness != "pass" {
		t.Fatalf("source completeness score = %q, want pass", result.Scores.SourceCompleteness)
	}
	if result.Scores.ApplyPlanCoverage != "pass" {
		t.Fatalf("apply plan coverage score = %q, want pass", result.Scores.ApplyPlanCoverage)
	}
	if result.Scores.ApplyReadiness != "pass" {
		t.Fatalf("apply readiness score = %q, want pass", result.Scores.ApplyReadiness)
	}
	if result.Scores.Overall != "pass" {
		t.Fatalf("overall score = %q, want pass", result.Scores.Overall)
	}

	encoded, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !strings.Contains(string(encoded), `"overall":"pass"`) {
		t.Fatalf("encoded scores missing overall pass: %s", encoded)
	}
}

func TestEvaluateBundleMarksUnsafeApplyPlanAsFail(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "apply-plan.md"), "# Reviewed Ingest Apply Plan\n\n"+
		"## Validation\n\n"+
		"- [x] `source.md` has YAML frontmatter\n\n"+
		"## Candidate Changes\n\n"+
		"- `wiki/index.md` — create new page.\n")

	result, err := EvaluateBundle(dir)
	if err != nil {
		t.Fatalf("EvaluateBundle returned error: %v", err)
	}
	if result.Scores.WikiSafety != "fail" {
		t.Fatalf("wiki safety score = %q, want fail", result.Scores.WikiSafety)
	}
	if result.Scores.CandidatePaths != "fail" {
		t.Fatalf("candidate paths score = %q, want fail", result.Scores.CandidatePaths)
	}
	if result.Scores.Overall != "fail" {
		t.Fatalf("overall score = %q, want fail", result.Scores.Overall)
	}
}

func TestEvaluateBundleFailsWhenPlannedTopicCandidateIsMissing(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/source.md\n---\n\n# Source\n")
	writeFile(t, filepath.Join(dir, "ingest-plan.md"), "---\nkind: topic\n---\n# Ingest Plan\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), validApplyPlanWithCandidates("wiki/sources/source.md", "wiki/topics/missing-topic.md"))

	result, err := EvaluateBundle(dir)
	if err != nil {
		t.Fatalf("EvaluateBundle returned error: %v", err)
	}
	if result.Scores.ApplyPlanCoverage != "fail" {
		t.Fatalf("apply plan coverage score = %q, want fail", result.Scores.ApplyPlanCoverage)
	}
	if result.Scores.Overall != "fail" {
		t.Fatalf("overall score = %q, want fail", result.Scores.Overall)
	}
}

func TestEvaluateBundleTreatsSourceCandidateAsRepresentedBySourceDraft(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/source.md\n---\n\n# Source\n")
	writeFile(t, filepath.Join(dir, "ingest-plan.md"), "---\nkind: topic\n---\n# Ingest Plan\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), validApplyPlan("wiki/sources/source.md"))

	result, err := EvaluateBundle(dir)
	if err != nil {
		t.Fatalf("EvaluateBundle returned error: %v", err)
	}
	if result.Scores.ApplyPlanCoverage != "pass" {
		t.Fatalf("apply plan coverage score = %q, want pass", result.Scores.ApplyPlanCoverage)
	}
	if result.Scores.Overall != "pass" {
		t.Fatalf("overall score = %q, want pass", result.Scores.Overall)
	}
}

func TestEvaluateBundleMarksMultipleSourceCandidatesAsBorderline(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), "---\nkind: source\nsources:\n  - raw/source.md\n---\n\n# Source\n")
	writeFile(t, filepath.Join(dir, "ingest-plan.md"), "---\nkind: topic\n---\n# Ingest Plan\n")
	writeFile(t, filepath.Join(dir, "apply-plan.md"), validApplyPlanWithCandidates("wiki/sources/source-a.md", "wiki/sources/source-b.md"))

	result, err := EvaluateBundle(dir)
	if err != nil {
		t.Fatalf("EvaluateBundle returned error: %v", err)
	}
	if result.Scores.ApplyPlanCoverage != "borderline" {
		t.Fatalf("apply plan coverage score = %q, want borderline", result.Scores.ApplyPlanCoverage)
	}
	if result.Scores.Overall != "borderline" {
		t.Fatalf("overall score = %q, want borderline", result.Scores.Overall)
	}
}

func TestEvaluateBundleFlagsCompletenessClaimsForTruncatedRawSource(t *testing.T) {
	dir := t.TempDir()
	rawPath := filepath.Join(dir, "raw", "moby-dick.md")
	writeFile(t, rawPath, "Chapter 93\n\nmiserably.\n\n[truncated at 900000 characters]\n")
	writeFile(t, filepath.Join(dir, "drafts", "bundle", "source.md"), "---\nkind: source\nsources:\n  - raw/moby-dick.md\n---\n\n## What It Is\n\nThe entire novel through the final chapters.\n")
	writeFile(t, filepath.Join(dir, "drafts", "bundle", "ingest-plan.md"), "---\nkind: topic\n---\n\n## Source Summary\n\nComplete plain-text retrieval with all 135 chapters.\n")
	writeFile(t, filepath.Join(dir, "drafts", "bundle", "apply-plan.md"), validApplyPlan("wiki/sources/moby-dick.md"))

	result, err := EvaluateBundle(filepath.Join(dir, "drafts", "bundle"))
	if err != nil {
		t.Fatalf("EvaluateBundle returned error: %v", err)
	}
	if result.Scores.SourceCompleteness != "fail" {
		t.Fatalf("source completeness score = %q, want fail", result.Scores.SourceCompleteness)
	}
	if result.Scores.Overall != "fail" {
		t.Fatalf("overall score = %q, want fail", result.Scores.Overall)
	}
}

func TestEvaluateBundlePassesTruncatedRawSourceWhenDraftAcknowledgesBoundary(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "raw", "moby-dick.md"), "Chapter 93\n\n[truncated at 900000 characters]\n")
	writeFile(t, filepath.Join(dir, "drafts", "bundle", "source.md"), "---\nkind: source\nsources:\n  - raw/moby-dick.md\n---\n\n## What It Is\n\nA partial corpus item truncated at 900000 characters near Chapter 93.\n")
	writeFile(t, filepath.Join(dir, "drafts", "bundle", "ingest-plan.md"), "---\nkind: topic\n---\n\n## Source Summary\n\nThe source is incomplete and ends around Chapter 93.\n")
	writeFile(t, filepath.Join(dir, "drafts", "bundle", "apply-plan.md"), validApplyPlan("wiki/sources/moby-dick.md"))

	result, err := EvaluateBundle(filepath.Join(dir, "drafts", "bundle"))
	if err != nil {
		t.Fatalf("EvaluateBundle returned error: %v", err)
	}
	if result.Scores.SourceCompleteness != "pass" {
		t.Fatalf("source completeness score = %q, want pass", result.Scores.SourceCompleteness)
	}
	if result.Scores.Overall != "pass" {
		t.Fatalf("overall score = %q, want pass", result.Scores.Overall)
	}
}

func TestEvaluateBundleMarksUnmentionedTruncationAsBorderline(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "raw", "moby-dick.md"), "Chapter 93\n\n[truncated at 900000 characters]\n")
	writeFile(t, filepath.Join(dir, "drafts", "bundle", "source.md"), "---\nkind: source\nsources:\n  - raw/moby-dick.md\n---\n\n## What It Is\n\nA Project Gutenberg corpus item for Moby-Dick.\n")
	writeFile(t, filepath.Join(dir, "drafts", "bundle", "ingest-plan.md"), "---\nkind: topic\n---\n\n## Source Summary\n\nA Project Gutenberg corpus item for Moby-Dick.\n")
	writeFile(t, filepath.Join(dir, "drafts", "bundle", "apply-plan.md"), validApplyPlan("wiki/sources/moby-dick.md"))

	result, err := EvaluateBundle(filepath.Join(dir, "drafts", "bundle"))
	if err != nil {
		t.Fatalf("EvaluateBundle returned error: %v", err)
	}
	if result.Scores.SourceCompleteness != "borderline" {
		t.Fatalf("source completeness score = %q, want borderline", result.Scores.SourceCompleteness)
	}
	if result.Scores.Overall != "borderline" {
		t.Fatalf("overall score = %q, want borderline", result.Scores.Overall)
	}
}

func validApplyPlan(candidate string) string {
	return validApplyPlanWithCandidates(candidate)
}

func validApplyPlanWithCandidates(candidates ...string) string {
	candidateLines := ""
	for _, candidate := range candidates {
		candidateLines += "- `" + candidate + "` — create new page.\n"
	}
	return "# Reviewed Ingest Apply Plan\n\n" +
		"This is a review artifact. Do not apply this plan automatically.\n\n" +
		"## Validation\n\n" +
		"- [x] `source.md` has YAML frontmatter\n" +
		"- [x] `ingest-plan.md` has YAML frontmatter\n" +
		"- [x] `source.md` frontmatter `kind` is `source`\n" +
		"- [x] `ingest-plan.md` frontmatter `kind` is `topic`\n" +
		"- [x] `source.md` includes required section `What It Is`\n" +
		"- [x] `source.md` includes required section `Summary`\n" +
		"- [x] `source.md` includes required section `Key Claims`\n" +
		"- [x] `source.md` includes required section `Suggested Links`\n" +
		"- [x] `ingest-plan.md` includes required section `Source Summary`\n" +
		"- [x] `ingest-plan.md` includes required section `Candidate Wiki Pages`\n" +
		"- [x] `ingest-plan.md` includes required section `Suggested Links`\n" +
		"- [x] `ingest-plan.md` includes required section `Review Checklist`\n\n" +
		"## Candidate Changes\n\n" +
		candidateLines
}

func writeFile(t *testing.T, path string, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
