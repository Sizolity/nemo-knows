package wikilint

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLintWikiReportsStructuralIssues(t *testing.T) {
	root := t.TempDir()
	writeWikiFile(t, root, "wiki/index.md", `---
title: Index
kind: index
---

## Concepts
- [[known]] — Known concept.
- [[known]] — Duplicate concept.
- [[missing-stub]] — Missing stub.
`)
	writeWikiFile(t, root, "wiki/log.md", `---
title: Log
kind: log
---

## [2026-05-16] apply-approved | bad action
Touched:
- wiki/concepts/known.md
`)
	writeWikiFile(t, root, "wiki/concepts/known.md", `---
title: Known
kind: concept
sources:
  - raw/source.md
confidence: medium
---

# Known

Links to [[missing-stub]].
`)
	writeWikiFile(t, root, "wiki/concepts/no-frontmatter.md", "# No Frontmatter\n")
	writeWikiFile(t, root, "wiki/topics/orphan-topic.md", `---
title: Orphan Topic
kind: topic
sources:
  - raw/source.md
confidence: medium
---

# Orphan Topic
`)

	result, err := LintWiki(root)
	if err != nil {
		t.Fatalf("LintWiki returned error: %v", err)
	}
	for _, code := range []string{
		"duplicate-index-entry",
		"missing-wikilink-target",
		"missing-frontmatter",
		"invalid-log-action",
		"orphan-page",
	} {
		if !hasIssue(result, code) {
			t.Fatalf("expected issue code %q in %#v", code, result.Issues)
		}
	}
	if result.Summary.Total == 0 {
		t.Fatal("expected non-empty lint summary")
	}
}

func TestLintWikiIgnoresExamplesInCode(t *testing.T) {
	root := t.TempDir()
	writeWikiFile(t, root, "wiki/index.md", "```text\n[[example-stub]]\n```\n\n`[[inline-example]]`\n")
	writeWikiFile(t, root, "wiki/log.md", `---
title: Log
kind: log
---

`+"```"+`
## [YYYY-MM-DD] <action> | <subject>
`+"```"+`

## [2026-05-16] note | ok
`)
	writeWikiFile(t, root, "wiki/concepts/known.md", `---
title: Known
kind: concept
sources:
  - raw/source.md
confidence: medium
---

# Known
`)

	result, err := LintWiki(root)
	if err != nil {
		t.Fatalf("LintWiki returned error: %v", err)
	}
	for _, code := range []string{"missing-wikilink-target", "invalid-log-action"} {
		if hasIssue(result, code) {
			t.Fatalf("did not expect issue code %q in %#v", code, result.Issues)
		}
	}
}

func writeWikiFile(t *testing.T, root string, rel string, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func hasIssue(result Result, code string) bool {
	for _, issue := range result.Issues {
		if issue.Code == code {
			return true
		}
	}
	return false
}
