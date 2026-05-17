package review

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReviewBundleCreatesApplyPlan(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), `---
kind: source
sources: [raw/llm-wiki.md]
---

# LLM Wiki

## What It Is
Source overview.

## Summary
Source summary.

## Key Claims
- Claim.

## Suggested Links
- https://example.com
`)
	writeFile(t, filepath.Join(dir, "ingest-plan.md"), `---
kind: topic
sources: [raw/llm-wiki.md]
status: draft
---

# Ingest Plan

## Source Summary
- Summary.

## Candidate Wiki Pages
- wiki/sources/llm-wiki.md — source page
- wiki/concepts/persistent-wiki.md — concept page
- wiki/topics/rag-vs-wiki.md — topic page

## Suggested Links
- https://example.com

## Review Checklist
- [ ] Check facts.
`)

	plan, err := ReviewBundle(dir)
	if err != nil {
		t.Fatalf("ReviewBundle returned error: %v", err)
	}

	for _, want := range []string{
		"# Reviewed Ingest Apply Plan",
		"Bundle: `" + dir + "`",
		"- [x] `source.md` has YAML frontmatter",
		"- [x] `ingest-plan.md` includes required section `Candidate Wiki Pages`",
		"- `wiki/sources/llm-wiki.md`",
		"- `wiki/concepts/persistent-wiki.md`",
		"- `wiki/topics/rag-vs-wiki.md`",
		"Do not apply this plan automatically.",
	} {
		if !strings.Contains(plan, want) {
			t.Fatalf("apply plan missing %q:\n%s", want, plan)
		}
	}
}

func TestReviewBundleRejectsInvalidCandidateWikiPath(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), `---
kind: source
---

# Source

## What It Is
Ok.

## Summary
Ok.

## Key Claims
Ok.

## Suggested Links
Ok.
`)
	writeFile(t, filepath.Join(dir, "ingest-plan.md"), `---
kind: topic
---

# Ingest Plan

## Source Summary
Ok.

## Candidate Wiki Pages
- wiki/index.md — invalid candidate

## Suggested Links
Ok.

## Review Checklist
Ok.
`)

	if _, err := ReviewBundle(dir); err == nil {
		t.Fatal("expected invalid candidate path error")
	}
}

func TestReviewBundleRejectsWrongFrontmatterKind(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), validSourceDraft("topic"))
	writeFile(t, filepath.Join(dir, "ingest-plan.md"), validIngestPlanDraft("- wiki/sources/llm-wiki.md — source page"))

	if _, err := ReviewBundle(dir); err == nil {
		t.Fatal("expected wrong source kind error")
	}
}

func TestReviewBundleOnlyParsesCandidateWikiPagesSection(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "source.md"), validSourceDraft("source"))
	writeFile(t, filepath.Join(dir, "ingest-plan.md"), `---
kind: topic
sources: [raw/llm-wiki.md]
---

# Ingest Plan

## Source Summary
Ok.

## Candidate Wiki Pages
- wiki/sources/llm-wiki.md — source page

## Suggested Links
Ok.

## Review Checklist
- [ ] Do not edit wiki/index.md automatically.
`)

	plan, err := ReviewBundle(dir)
	if err != nil {
		t.Fatalf("ReviewBundle returned error: %v", err)
	}
	if strings.Contains(plan, "- `wiki/index.md`") {
		t.Fatalf("apply plan included checklist-only path:\n%s", plan)
	}
}

func TestReviewBundleLabelsExistingAndDuplicateCandidates(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "wiki", "sources"), 0o755); err != nil {
		t.Fatalf("mkdir source wiki: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(root, "wiki", "concepts"), 0o755); err != nil {
		t.Fatalf("mkdir concept wiki: %v", err)
	}
	writeFile(t, filepath.Join(root, "wiki", "sources", "llm-wiki.md"), "---\nkind: source\n---\n")
	writeFile(t, filepath.Join(root, "wiki", "concepts", "persistent-wiki.md"), "---\nkind: concept\n---\n")

	bundle := filepath.Join(root, "drafts", "bundle")
	if err := os.MkdirAll(bundle, 0o755); err != nil {
		t.Fatalf("mkdir bundle: %v", err)
	}
	writeFile(t, filepath.Join(bundle, "source.md"), validSourceDraft("source"))
	writeFile(t, filepath.Join(bundle, "ingest-plan.md"), validIngestPlanDraft(strings.Join([]string{
		"- wiki/sources/llm-wiki.md — source page",
		"- wiki/concepts/persistent-wiki-architecture.md — similar concept page",
		"- wiki/topics/rag-vs-wiki.md — topic page",
	}, "\n")))

	plan, err := ReviewBundleWithRoot(root, bundle)
	if err != nil {
		t.Fatalf("ReviewBundleWithRoot returned error: %v", err)
	}
	for _, want := range []string{
		"- `wiki/sources/llm-wiki.md` — update existing page.",
		"- `wiki/concepts/persistent-wiki-architecture.md` — create new page; possible duplicate of `wiki/concepts/persistent-wiki.md`.",
		"- `wiki/topics/rag-vs-wiki.md` — create new page.",
	} {
		if !strings.Contains(plan, want) {
			t.Fatalf("apply plan missing %q:\n%s", want, plan)
		}
	}
}

func validSourceDraft(kind string) string {
	return `---
kind: ` + kind + `
sources: [raw/llm-wiki.md]
---

# Source

## What It Is
Ok.

## Summary
Ok.

## Key Claims
Ok.

## Suggested Links
Ok.
`
}

func validIngestPlanDraft(candidateLines string) string {
	return `---
kind: topic
sources: [raw/llm-wiki.md]
---

# Ingest Plan

## Source Summary
Ok.

## Candidate Wiki Pages
` + candidateLines + `

## Suggested Links
Ok.

## Review Checklist
Ok.
`
}

func writeFile(t *testing.T, path string, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
