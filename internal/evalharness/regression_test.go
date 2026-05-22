package evalharness

import (
	"path/filepath"
	"testing"
)

func TestEvaluateRegressionCasesSummarizesMultipleCases(t *testing.T) {
	dir := t.TempDir()
	makeRegressionCase(t, dir, "article", "wiki/concepts/article-pattern.md")
	makeRegressionCase(t, dir, "meeting-notes", "wiki/topics/meeting-decisions.md")

	result, err := EvaluateRegressionCases(dir)
	if err != nil {
		t.Fatalf("EvaluateRegressionCases returned error: %v", err)
	}
	if len(result.Cases) != 2 {
		t.Fatalf("case count = %d, want 2", len(result.Cases))
	}
	if result.Summary.Total != 2 || result.Summary.Passed != 2 || result.Summary.Failed != 0 {
		t.Fatalf("unexpected summary: %#v", result.Summary)
	}
	if result.Overall != "pass" {
		t.Fatalf("overall = %q, want pass", result.Overall)
	}
}

func makeRegressionCase(t *testing.T, casesDir string, name string, candidate string) {
	t.Helper()
	caseDir := filepath.Join(casesDir, name)
	writeFile(t, filepath.Join(caseDir, "expected.json"), `{
  "case": "`+name+`",
  "source": "fixture:`+name+`",
  "bundle": "`+filepath.ToSlash(filepath.Join(caseDir, "bundle"))+`",
  "minimum_scores": {
    "schema": "pass",
    "wiki_safety": "pass",
    "candidate_paths": "pass",
    "duplicate_detection": "pass",
    "source_completeness": "pass",
    "apply_plan_coverage": "pass",
    "apply_readiness": "pass",
    "overall": "pass"
  }
}
`)
	writeFile(t, filepath.Join(caseDir, "bundle", "source.md"), "---\nkind: source\n---\n\n# Source\n\n## What It Is\nOk.\n\n## Summary\nOk.\n\n## Key Claims\nOk.\n\n## Suggested Links\nOk.\n")
	writeFile(t, filepath.Join(caseDir, "bundle", "ingest-plan.md"), "---\nkind: topic\n---\n\n# Ingest Plan\n")
	writeFile(t, filepath.Join(caseDir, "bundle", "apply-plan.md"), "# Reviewed Ingest Apply Plan\n\n"+
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
		"- `"+candidate+"` — create new page.\n")
	writeFile(t, filepath.Join(caseDir, "bundle", "candidates", filepath.FromSlash(candidate)), "---\ntitle: Regression Candidate\nkind: "+candidateKind(candidate)+"\nsources:\n  - source.md\n  - raw/source.md\nconfidence: medium\n---\n\n# Regression Candidate\n\nThis regression candidate represents the planned page.\n")
}

func candidateKind(candidate string) string {
	if filepath.Dir(candidate) == "wiki/topics" {
		return "topic"
	}
	return "concept"
}
