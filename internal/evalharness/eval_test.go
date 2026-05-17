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

func writeFile(t *testing.T, path string, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
