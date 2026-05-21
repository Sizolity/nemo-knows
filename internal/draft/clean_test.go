package draft

import (
	"strings"
	"testing"
)

func TestCleanRemovesThinkingAndRuntimeNoise(t *testing.T) {
	raw := `Prompt echo
[Start thinking]
internal reasoning that must not be in the draft
[End thinking]

---
kind: source
sources:
  - raw/llm-wiki.md
title: LLM Wiki
---

# What It Is

A concise source page.

[ Prompt: 100 t/s | Generation: 50 t/s ]

Exiting...
common_memory_breakdown_print: noisy gpu log
`

	cleaned, err := Clean(raw)
	if err != nil {
		t.Fatalf("Clean returned error: %v", err)
	}

	if !strings.HasPrefix(cleaned, "---\nkind: source") {
		t.Fatalf("expected cleaned draft to start at frontmatter, got:\n%s", cleaned)
	}
	for _, unwanted := range []string{
		"[Start thinking]",
		"internal reasoning",
		"[ Prompt:",
		"Exiting...",
		"common_memory_breakdown_print",
	} {
		if strings.Contains(cleaned, unwanted) {
			t.Fatalf("cleaned draft still contains %q:\n%s", unwanted, cleaned)
		}
	}
}

func TestCleanRecoversFencedMarkdownWhenFinalAnswerIsIncomplete(t *testing.T) {
	raw := `[Start thinking]
The model drafted the useful page here:

` + "```markdown" + `
---
kind: source
sources: [raw/llm-wiki.md]
title: LLM Wiki
---

# LLM Wiki

## Summary
The useful draft is inside the thinking block.
` + "```" + `
[End thinking]

---
kind:

[ Prompt: 2239.0 t/s | Generation: 56.5 t/s ]

Exiting...
`

	cleaned, err := Clean(raw)
	if err != nil {
		t.Fatalf("Clean returned error: %v", err)
	}

	if !strings.HasPrefix(cleaned, "---\nkind: source") {
		t.Fatalf("expected fenced markdown fallback, got:\n%s", cleaned)
	}
	if strings.Contains(cleaned, "```") {
		t.Fatalf("expected cleaner to remove markdown fences, got:\n%s", cleaned)
	}
	if !strings.Contains(cleaned, "The useful draft is inside the thinking block.") {
		t.Fatalf("expected fallback draft content, got:\n%s", cleaned)
	}
}

func TestCleanErrorsWhenNoFrontmatterCanBeRecovered(t *testing.T) {
	_, err := Clean("plain output without a page")
	if err == nil {
		t.Fatal("expected error for output without recoverable frontmatter")
	}
}

func TestCleanRemovesInlineMemoryBreakdownLog(t *testing.T) {
	raw := `---
kind: source
---
# Draft
- [Dataview](https://example.test)common_memory_breakdown_print: | memory breakdown [MiB] | total
common_memory_breakdown_print: |   - Host | noisy
`

	cleaned, err := Clean(raw)
	if err != nil {
		t.Fatalf("Clean returned error: %v", err)
	}
	if strings.Contains(cleaned, "common_memory_breakdown_print") {
		t.Fatalf("cleaned draft still contains memory log:\n%s", cleaned)
	}
	if !strings.Contains(cleaned, "- [Dataview](https://example.test)") {
		t.Fatalf("cleaned draft lost content before memory log:\n%s", cleaned)
	}
}

func TestCleanExtractsFrontmatterOnlyFromDelimiterLine(t *testing.T) {
	raw := "Begin with exactly `---`.\n\n---\nkind: topic\n---\n# Draft\n"

	cleaned, err := Clean(raw)
	if err != nil {
		t.Fatalf("Clean returned error: %v", err)
	}
	if strings.Contains(cleaned, "Begin with exactly") {
		t.Fatalf("cleaned draft included prompt text:\n%s", cleaned)
	}
	if !strings.HasPrefix(cleaned, "---\nkind: topic") {
		t.Fatalf("expected frontmatter delimiter line, got:\n%s", cleaned)
	}
}

func TestCleanRecoversYamlFencedMarkdown(t *testing.T) {
	raw := "```yaml\n---\ntitle: Whaling Practices\nkind: topic\n---\n# Whaling Practices\n\nBody.\n```"

	cleaned, err := Clean(raw)
	if err != nil {
		t.Fatalf("Clean returned error: %v", err)
	}
	if strings.Contains(cleaned, "```") {
		t.Fatalf("cleaned draft still contains fence:\n%s", cleaned)
	}
	if !strings.Contains(cleaned, "# Whaling Practices") {
		t.Fatalf("cleaned draft lost body:\n%s", cleaned)
	}
}

func TestCleanRecoversFencedPseudoFrontmatter(t *testing.T) {
	raw := "```yaml\ntitle: Whiteness Of The Whale\nkind: topic\npath: wiki/topics/whiteness-of-the-whale.md\n---\n\n# Whiteness Of The Whale\n\nBody.\n```"

	cleaned, err := Clean(raw)
	if err != nil {
		t.Fatalf("Clean returned error: %v", err)
	}
	if !strings.HasPrefix(cleaned, "---\ntitle: Whiteness Of The Whale") {
		t.Fatalf("expected pseudo-frontmatter to be wrapped, got:\n%s", cleaned)
	}
	if strings.Contains(cleaned, "```") {
		t.Fatalf("cleaned draft still contains fence:\n%s", cleaned)
	}
}
