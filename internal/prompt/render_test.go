package prompt

import (
	"strings"
	"testing"
)

func TestRenderReplacesKnownVariables(t *testing.T) {
	rendered, err := Render(
		"Path: {{RAW_SOURCE_PATH}}\nBody:\n{{RAW_SOURCE_CONTENT}}\nConcept: {{CONCEPT_NAME}}\nSources: {{SOURCE_LIST}}\nSource: {{SOURCE_CONTENT}}\nTitle: {{PAGE_TITLE}}\nKind: {{PAGE_KIND}}\nTarget: {{TARGET_PATH}}\nAllowed: {{ALLOWED_LINKS}}\nChunk: {{CHUNK_CONTENT}}\nNotes: {{CHUNK_NOTES}}\nGroup notes: {{CHUNK_GROUP_NOTES}}\nOutline: {{CHUNK_OUTLINE}}\nIndex: {{CHUNK_INDEX}}",
		Variables{
			RawSourcePath:    "raw/llm-wiki.md",
			RawSourceContent: "# LLM Wiki",
			ConceptName:      "Persistent Wiki",
			SourceList:       "wiki/sources/llm-wiki.md",
			SourceContent:    "Source summary",
			PageTitle:        "Persistent Wiki",
			PageKind:         "concept",
			TargetPath:       "wiki/concepts/persistent-wiki.md",
			AllowedLinks:     "- [[query]]",
			ChunkContent:     "chunk text",
			ChunkNotes:       "chunk notes",
			ChunkGroupNotes:  "chunk group notes",
			ChunkOutline:     "chunk outline",
			ChunkIndex:       "chunk index",
		},
	)
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}

	for _, want := range []string{
		"raw/llm-wiki.md",
		"# LLM Wiki",
		"Persistent Wiki",
		"wiki/sources/llm-wiki.md",
		"Source summary",
		"concept",
		"wiki/concepts/persistent-wiki.md",
		"- [[query]]",
		"chunk text",
		"chunk notes",
		"chunk group notes",
		"chunk outline",
		"chunk index",
	} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("rendered prompt missing %q:\n%s", want, rendered)
		}
	}

	for _, placeholder := range []string{
		"{{RAW_SOURCE_PATH}}",
		"{{RAW_SOURCE_CONTENT}}",
		"{{CONCEPT_NAME}}",
		"{{SOURCE_LIST}}",
		"{{SOURCE_CONTENT}}",
	} {
		if strings.Contains(rendered, placeholder) {
			t.Fatalf("rendered prompt still contains placeholder %q:\n%s", placeholder, rendered)
		}
	}
}

func TestRenderDoesNotPartiallyReplaceOverlappingLegacyNames(t *testing.T) {
	rendered, err := Render("Raw: RAW_SOURCE_CONTENT\nSource: SOURCE_CONTENT", Variables{
		RawSourceContent: "full raw source",
		SourceContent:    "summary source",
	})
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}

	if strings.Contains(rendered, "RAW_") {
		t.Fatalf("rendered prompt has partial replacement artifact:\n%s", rendered)
	}
	if !strings.Contains(rendered, "full raw source") {
		t.Fatalf("rendered prompt missing raw source content:\n%s", rendered)
	}
	if !strings.Contains(rendered, "summary source") {
		t.Fatalf("rendered prompt missing source content:\n%s", rendered)
	}
}
