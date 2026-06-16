package wikilint

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLintWikiReportsStructuralIssues(t *testing.T) {
	root := t.TempDir()
	writeWikiFile(t, root, "wiki/index.md", `---
title: Index
kind: index
---

## Concepts
- [known](concepts/known.md) — Known concept.
- [[known]] — Duplicate concept.
- [missing-stub](concepts/missing-stub.md) — Missing stub.
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
		"forbidden-wikilink",
		"missing-link-target",
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

func TestLintWikiAcceptsMarkdownRelativeLinks(t *testing.T) {
	root := t.TempDir()
	writeWikiFile(t, root, "wiki/index.md", `---
title: Index
kind: index
---

## Concepts
- [serverless-database](concepts/serverless-database.md) — Serverless engine.

## Topics
- [data-preservation-formats](topics/data-preservation-formats.md) — Formats.
`)
	writeWikiFile(t, root, "wiki/log.md", `---
title: Log
kind: log
---

## [2026-06-16] note | ok
`)
	writeWikiFile(t, root, "wiki/concepts/serverless-database.md", `---
title: Serverless Database
kind: concept
sources:
  - raw/source.md
confidence: medium
---

# Serverless Database

Recommended for [data preservation](../topics/data-preservation-formats.md).
`)
	writeWikiFile(t, root, "wiki/topics/data-preservation-formats.md", `---
title: Data Preservation Formats
kind: topic
sources:
  - raw/source.md
confidence: medium
---

# Data Preservation Formats

Used by the [serverless database](../concepts/serverless-database.md).
`)

	result, err := LintWiki(root)
	if err != nil {
		t.Fatalf("LintWiki returned error: %v", err)
	}
	for _, code := range []string{"forbidden-wikilink", "missing-link-target", "orphan-page"} {
		if hasIssue(result, code) {
			t.Fatalf("did not expect issue code %q for valid relative links in %#v", code, result.Issues)
		}
	}
}

func TestLintWikiFlagsMissingRelativeLinkTarget(t *testing.T) {
	root := t.TempDir()
	writeWikiFile(t, root, "wiki/index.md", `---
title: Index
kind: index
---

## Concepts
- [present](concepts/present.md) — Present.
`)
	writeWikiFile(t, root, "wiki/log.md", `---
title: Log
kind: log
---

## [2026-06-16] note | ok
`)
	writeWikiFile(t, root, "wiki/concepts/present.md", `---
title: Present
kind: concept
sources:
  - raw/source.md
confidence: medium
---

# Present

Points at a [moved page](../concepts/gone.md).
`)

	result, err := LintWiki(root)
	if err != nil {
		t.Fatalf("LintWiki returned error: %v", err)
	}
	var msg string
	for _, issue := range result.Issues {
		if issue.Code == "missing-link-target" {
			msg = issue.Message
		}
	}
	if msg == "" {
		t.Fatalf("expected missing-link-target issue in %#v", result.Issues)
	}
	if !strings.Contains(msg, "gone.md") {
		t.Fatalf("missing-link-target message should name the broken target, got %q", msg)
	}
}

func TestLintWikiFlagsCrossCategorySlugCollision(t *testing.T) {
	root := t.TempDir()
	writeWikiFile(t, root, "wiki/index.md", `---
title: Index
kind: index
---

## Sources
- [sqlite](sources/sqlite.md) — Source page.

## Entities
- [sqlite](entities/sqlite.md) — Entity page.
`)
	writeWikiFile(t, root, "wiki/log.md", `---
title: Log
kind: log
---

## [2026-05-16] note | ok
`)
	writeWikiFile(t, root, "wiki/sources/sqlite.md", `---
title: SQLite Source
kind: source
sources:
  - raw/sqlite.md
confidence: medium
---

# SQLite Source
`)
	writeWikiFile(t, root, "wiki/entities/sqlite.md", `---
title: SQLite
kind: entity
sources:
  - raw/sqlite.md
confidence: medium
---

# SQLite
`)

	result, err := LintWiki(root)
	if err != nil {
		t.Fatalf("LintWiki returned error: %v", err)
	}
	if !hasIssue(result, "duplicate-slug") {
		t.Fatalf("expected duplicate-slug issue for cross-category collision in %#v", result.Issues)
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
	for _, code := range []string{"forbidden-wikilink", "invalid-log-action"} {
		if hasIssue(result, code) {
			t.Fatalf("did not expect issue code %q in %#v", code, result.Issues)
		}
	}
}

func TestLintWikiChecksImageExistence(t *testing.T) {
	root := t.TempDir()
	writeWikiFile(t, root, "wiki/index.md", `---
title: Index
kind: index
---

## Sources
- [slides](sources/slides.md) — Slide deck.
`)
	writeWikiFile(t, root, "wiki/log.md", `---
title: Log
kind: log
---

## [2026-06-16] note | ok
`)
	writeWikiFile(t, root, "wiki/sources/slides.md", `---
title: Slides
kind: source
sources:
  - https://example.com/deck
confidence: medium
---

# Slides

![present](../assets/deck/present.png)
![missing](../assets/deck/missing.png)
![remote](https://example.com/remote.png)
`)
	// Only the "present" asset exists on disk.
	writeWikiFile(t, root, "wiki/assets/deck/present.png", "\x89PNG\r\n\x1a\n")

	result, err := LintWiki(root)
	if err != nil {
		t.Fatalf("LintWiki returned error: %v", err)
	}
	if !hasIssue(result, "missing-image") {
		t.Fatalf("expected missing-image issue in %#v", result.Issues)
	}
	count := 0
	var msg string
	for _, issue := range result.Issues {
		if issue.Code == "missing-image" {
			count++
			msg = issue.Message
		}
	}
	if count != 1 {
		t.Fatalf("expected exactly 1 missing-image (present + remote excluded), got %d in %#v", count, result.Issues)
	}
	if !strings.Contains(msg, "missing.png") {
		t.Fatalf("missing-image message should name the missing file, got %q", msg)
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
