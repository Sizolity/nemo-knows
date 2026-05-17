package apply

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestApplyApprovedRequiresApproval(t *testing.T) {
	root, bundle := makeApplyFixture(t, "pass")

	_, err := ApplyApproved(root, bundle, Options{Approve: false})
	if err == nil {
		t.Fatal("expected approval error")
	}
}

func TestApplyApprovedRejectsFailingEval(t *testing.T) {
	root, bundle := makeApplyFixture(t, "fail")

	_, err := ApplyApproved(root, bundle, Options{Approve: true})
	if err == nil {
		t.Fatal("expected failing eval error")
	}
}

func TestApplyApprovedUpdatesDuplicateSourceTargetAndWritesReport(t *testing.T) {
	root, bundle := makeApplyFixture(t, "pass")

	result, err := ApplyApproved(root, bundle, Options{Approve: true})
	if err != nil {
		t.Fatalf("ApplyApproved returned error: %v", err)
	}

	sourcePath := filepath.Join(root, "wiki", "sources", "llm-wiki.md")
	source, err := os.ReadFile(sourcePath)
	if err != nil {
		t.Fatalf("read applied source: %v", err)
	}
	if !strings.Contains(string(source), "Updated source draft") {
		t.Fatalf("source target was not updated:\n%s", source)
	}
	if _, err := os.Stat(filepath.Join(root, "wiki", "sources", "llm-wiki-pattern.md")); err == nil {
		t.Fatal("duplicate source page should not be created")
	}
	if len(result.Written) == 0 {
		t.Fatal("expected written files in result")
	}

	reportPath := filepath.Join(bundle, "apply-report.md")
	report, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("read apply report: %v", err)
	}
	if !strings.Contains(string(report), "wiki/sources/llm-wiki.md") {
		t.Fatalf("report missing applied source target:\n%s", report)
	}
	if !strings.Contains(string(report), "Skipped") {
		t.Fatalf("report missing skipped candidates:\n%s", report)
	}

	log, err := os.ReadFile(filepath.Join(root, "wiki", "log.md"))
	if err != nil {
		t.Fatalf("read log: %v", err)
	}
	if !strings.Contains(string(log), "ingest | drafts/bundle") {
		t.Fatalf("log missing apply entry:\n%s", log)
	}
}

func TestApplyApprovedRejectsDuplicateApplyWithoutForce(t *testing.T) {
	root, bundle := makeApplyFixture(t, "pass")
	writeFile(t, filepath.Join(root, "wiki", "log.md"), "# Log\n\n## [2026-05-16] ingest | drafts/bundle\nTouched:\n- wiki/sources/llm-wiki.md\n")

	_, err := ApplyApproved(root, bundle, Options{Approve: true})
	if err == nil {
		t.Fatal("expected duplicate apply error")
	}
}

func TestApplyApprovedAllowsDuplicateApplyWithForce(t *testing.T) {
	root, bundle := makeApplyFixture(t, "pass")
	writeFile(t, filepath.Join(root, "wiki", "log.md"), "# Log\n\n## [2026-05-16] ingest | drafts/bundle\nTouched:\n- wiki/sources/llm-wiki.md\n")

	_, err := ApplyApproved(root, bundle, Options{Approve: true, Force: true})
	if err != nil {
		t.Fatalf("ApplyApproved returned error: %v", err)
	}
}

func TestApplyApprovedAppliesReviewedConceptDraftAndUpdatesIndex(t *testing.T) {
	root, bundle := makeApplyFixture(t, "pass")
	conceptDraft := filepath.Join(bundle, "candidates", "wiki", "concepts", "llm-maintenance-pattern.md")
	writeFile(t, conceptDraft, `---
title: LLM Maintenance Pattern
kind: concept
sources:
  - raw/llm-wiki.md
---

# LLM Maintenance Pattern

A reviewed concept draft.
`)

	result, err := ApplyApproved(root, bundle, Options{Approve: true})
	if err != nil {
		t.Fatalf("ApplyApproved returned error: %v", err)
	}
	if !contains(result.Written, "wiki/concepts/llm-maintenance-pattern.md") {
		t.Fatalf("expected concept draft to be written, got %#v", result.Written)
	}

	applied, err := os.ReadFile(filepath.Join(root, "wiki", "concepts", "llm-maintenance-pattern.md"))
	if err != nil {
		t.Fatalf("read applied concept: %v", err)
	}
	if !strings.Contains(string(applied), "A reviewed concept draft.") {
		t.Fatalf("concept content not applied:\n%s", applied)
	}

	index, err := os.ReadFile(filepath.Join(root, "wiki", "index.md"))
	if err != nil {
		t.Fatalf("read index: %v", err)
	}
	if !strings.Contains(string(index), "- [[llm-maintenance-pattern]] — LLM Maintenance Pattern.") {
		t.Fatalf("index missing concept entry:\n%s", index)
	}
}

func TestApplyApprovedAddsNewSourcePageToIndex(t *testing.T) {
	root, bundle := makeApplyFixture(t, "pass")
	writeFile(t, filepath.Join(bundle, "apply-plan.md"), "# Reviewed Ingest Apply Plan\n\n"+
		"## Candidate Changes\n\n"+
		"- `wiki/sources/new-source.md` — create new page.\n")
	writeFile(t, filepath.Join(bundle, "source.md"), `---
title: New Source
kind: source
sources:
  - raw/new-source.md
confidence: medium
---

# New Source

Reviewed source summary.
`)

	result, err := ApplyApproved(root, bundle, Options{Approve: true})
	if err != nil {
		t.Fatalf("ApplyApproved returned error: %v", err)
	}
	if !contains(result.Written, "wiki/sources/new-source.md") {
		t.Fatalf("expected source page to be written, got %#v", result.Written)
	}
	if !contains(result.Written, "wiki/index.md") {
		t.Fatalf("expected index to be written, got %#v", result.Written)
	}

	index, err := os.ReadFile(filepath.Join(root, "wiki", "index.md"))
	if err != nil {
		t.Fatalf("read index: %v", err)
	}
	if !strings.Contains(string(index), "- [[new-source]] — New Source.") {
		t.Fatalf("index missing source entry:\n%s", index)
	}
}

func TestApplyApprovedRecordsIndexOnceForMultipleNewCandidatePages(t *testing.T) {
	root, bundle := makeApplyFixture(t, "pass")
	writeFile(t, filepath.Join(bundle, "apply-plan.md"), "# Reviewed Ingest Apply Plan\n\n"+
		"## Candidate Changes\n\n"+
		"- `wiki/concepts/llm-maintenance-pattern.md` — create new page.\n"+
		"- `wiki/topics/persistent-wiki-architecture.md` — create new page.\n")
	writeFile(t, filepath.Join(bundle, "candidates", "wiki", "concepts", "llm-maintenance-pattern.md"), `---
title: LLM Maintenance Pattern
kind: concept
sources:
  - raw/llm-wiki.md
---

# LLM Maintenance Pattern
`)
	writeFile(t, filepath.Join(bundle, "candidates", "wiki", "topics", "persistent-wiki-architecture.md"), `---
title: Persistent Wiki Architecture
kind: topic
sources:
  - raw/llm-wiki.md
---

# Persistent Wiki Architecture
`)

	result, err := ApplyApproved(root, bundle, Options{Approve: true})
	if err != nil {
		t.Fatalf("ApplyApproved returned error: %v", err)
	}
	if got := countItems(result.Written, "wiki/index.md"); got != 1 {
		t.Fatalf("expected wiki/index.md to be recorded once, got %d in %#v", got, result.Written)
	}

	report, err := os.ReadFile(filepath.Join(bundle, "apply-report.md"))
	if err != nil {
		t.Fatalf("read apply report: %v", err)
	}
	if got := strings.Count(string(report), "- wiki/index.md"); got != 1 {
		t.Fatalf("expected apply report to record index once, got %d:\n%s", got, report)
	}
}

func TestApplyApprovedRejectsCandidateDraftWithWrongKind(t *testing.T) {
	root, bundle := makeApplyFixture(t, "pass")
	conceptDraft := filepath.Join(bundle, "candidates", "wiki", "concepts", "llm-maintenance-pattern.md")
	writeFile(t, conceptDraft, `---
title: LLM Maintenance Pattern
kind: topic
sources:
  - raw/llm-wiki.md
---

# LLM Maintenance Pattern
`)

	_, err := ApplyApproved(root, bundle, Options{Approve: true})
	if err == nil {
		t.Fatal("expected wrong candidate draft kind error")
	}
}

func TestApplyApprovedSkipsCandidateWithoutReviewedDraft(t *testing.T) {
	root, bundle := makeApplyFixture(t, "pass")

	_, err := ApplyApproved(root, bundle, Options{Approve: true})
	if err != nil {
		t.Fatalf("ApplyApproved returned error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "wiki", "concepts", "llm-maintenance-pattern.md")); err == nil {
		t.Fatal("concept page should not be created without reviewed draft")
	}

	report, err := os.ReadFile(filepath.Join(bundle, "apply-report.md"))
	if err != nil {
		t.Fatalf("read apply report: %v", err)
	}
	if !strings.Contains(string(report), "wiki/concepts/llm-maintenance-pattern.md — missing reviewed candidate draft") {
		t.Fatalf("report missing missing-draft skip reason:\n%s", report)
	}
}

func makeApplyFixture(t *testing.T, overall string) (string, string) {
	t.Helper()
	root := t.TempDir()
	for _, dir := range []string{
		filepath.Join(root, "wiki", "sources"),
		filepath.Join(root, "wiki", "concepts"),
		filepath.Join(root, "wiki", "topics"),
		filepath.Join(root, "drafts", "bundle"),
	} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", dir, err)
		}
	}

	writeFile(t, filepath.Join(root, "wiki", "sources", "llm-wiki.md"), "---\nkind: source\n---\n# Old\n")
	writeFile(t, filepath.Join(root, "wiki", "index.md"), "---\ntitle: Index\nkind: index\n---\n\n## Sources\n- [[llm-wiki]] — Existing source.\n")
	writeFile(t, filepath.Join(root, "wiki", "log.md"), "# Log\n")

	bundle := filepath.Join(root, "drafts", "bundle")
	writeFile(t, filepath.Join(bundle, "source.md"), "---\nkind: source\nsources: [raw/llm-wiki.md]\n---\n\n# LLM Wiki\n\nUpdated source draft.\n")
	writeFile(t, filepath.Join(bundle, "apply-plan.md"), "# Reviewed Ingest Apply Plan\n\n"+
		"This is a review artifact. Do not apply this plan automatically.\n\n"+
		"## Candidate Changes\n\n"+
		"- `wiki/sources/llm-wiki-pattern.md` — create new page; possible duplicate of `wiki/sources/llm-wiki.md`.\n"+
		"- `wiki/concepts/llm-maintenance-pattern.md` — create new page.\n")
	writeFile(t, filepath.Join(bundle, "scores.json"), "{\n  \"scores\": {\n    \"overall\": \""+overall+"\"\n  }\n}\n")

	return root, bundle
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

func contains(items []string, want string) bool {
	for _, item := range items {
		if item == want {
			return true
		}
	}

	return false
}

func countItems(items []string, want string) int {
	count := 0
	for _, item := range items {
		if item == want {
			count++
		}
	}

	return count
}
