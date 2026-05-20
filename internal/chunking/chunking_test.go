package chunking

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPlanSourceUsesHeadingBoundaries(t *testing.T) {
	source := strings.Join([]string{
		"# Title",
		"",
		"Intro.",
		"",
		"1. Overview",
		"",
		strings.Repeat("overview ", 20),
		"",
		"1.1 Detail",
		"",
		strings.Repeat("detail ", 20),
		"",
		"2. API",
		"",
		strings.Repeat("api ", 20),
	}, "\n")

	plan := PlanSource("raw/source.md", source, 180)
	if len(plan.Chunks) < 2 {
		t.Fatalf("expected multiple chunks, got %#v", plan.Chunks)
	}
	if plan.Chunks[0].StartLine != 1 {
		t.Fatalf("first chunk starts at line %d, want 1", plan.Chunks[0].StartLine)
	}
	foundDetail := false
	for _, chunk := range plan.Chunks {
		if strings.Contains(strings.Join(chunk.HeadingPath, " > "), "1.1 Detail") {
			foundDetail = true
		}
		if chunk.Index == 0 {
			t.Fatalf("chunk index was not set: %#v", chunk)
		}
	}
	if !foundDetail {
		t.Fatalf("expected detail heading path in chunks: %#v", plan.Chunks)
	}
}

func TestPlanSourcePreservesHeadingCoverageForPackedChunks(t *testing.T) {
	source := strings.Join([]string{
		"# Title",
		"",
		"1. First",
		"",
		"alpha",
		"",
		"2. Second",
		"",
		"beta",
		"",
		"3. Third",
		"",
		"gamma",
	}, "\n")

	plan := PlanSource("raw/source.md", source, 1000)
	if len(plan.Chunks) != 1 {
		t.Fatalf("expected short sections to pack into one chunk, got %d", len(plan.Chunks))
	}
	chunk := plan.Chunks[0]
	for _, want := range []string{"1. First", "2. Second", "3. Third"} {
		found := false
		for _, path := range chunk.SectionHeadingPaths {
			if strings.Contains(strings.Join(path, " > "), want) {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing heading %q in coverage: %#v", want, chunk.SectionHeadingPaths)
		}
	}
	formatted := FormatChunkForPrompt("raw/source.md", chunk, len(plan.Chunks))
	if !strings.Contains(formatted, "Heading coverage:") || !strings.Contains(formatted, "2. Second") {
		t.Fatalf("formatted chunk did not include coverage:\n%s", formatted)
	}
	outline := plan.OutlineMarkdown()
	if !strings.Contains(outline, "Heading coverage:") || !strings.Contains(outline, "3. Third") {
		t.Fatalf("outline did not include coverage:\n%s", outline)
	}
}

func TestPlanSourceDoesNotTreatNumberedProseAsHeading(t *testing.T) {
	source := strings.Join([]string{
		"# Title",
		"",
		"3.7. FTS5 Boolean Operators",
		"",
		"Body.",
		"2. for which the number of tokens between phrases is limited",
		"continues the paragraph.",
		"",
		"4. FTS5 Table Creation and Initialization",
		"",
		"Next body.",
	}, "\n")

	plan := PlanSource("raw/source.md", source, 1000)
	for _, chunk := range plan.Chunks {
		path := strings.Join(chunk.HeadingPath, " > ")
		if strings.Contains(path, "for which the number") {
			t.Fatalf("numbered prose was treated as heading: %#v", plan.Chunks)
		}
	}
}

func TestPlanSourceSplitsOversizedSectionByParagraph(t *testing.T) {
	source := "# Title\n\n1. Big Section\n\n" +
		strings.Repeat("alpha ", 80) + "\n\n" +
		strings.Repeat("beta ", 80) + "\n\n" +
		strings.Repeat("gamma ", 80)

	plan := PlanSource("raw/source.md", source, 250)
	if len(plan.Chunks) < 3 {
		t.Fatalf("expected oversized section to split, got %d chunks", len(plan.Chunks))
	}
	for _, chunk := range plan.Chunks {
		if !strings.Contains(chunk.Text, "Context:") && chunk.Index > 1 {
			t.Fatalf("split chunk missing heading context:\n%s", chunk.Text)
		}
	}
}

func TestPlanSourceSplitsOversizedParagraphBySoftBoundaries(t *testing.T) {
	source := "# Title\n\n1. Big Section\n\n" +
		strings.Repeat("alpha words stay together. ", 8) +
		strings.Repeat("beta words stay together. ", 8) +
		strings.Repeat("gamma words stay together. ", 8)

	plan := PlanSource("raw/source.md", source, 160)
	if len(plan.Chunks) < 3 {
		t.Fatalf("expected oversized paragraph to split by sentence groups, got %d chunks", len(plan.Chunks))
	}
	combined := ""
	for _, chunk := range plan.Chunks {
		if strings.Contains(chunk.Text, "Oversized paragraph segment:") {
			if !strings.Contains(chunk.Text, "Context: 1. Big Section") {
				t.Fatalf("split chunk missing heading context:\n%s", chunk.Text)
			}
			combined += chunk.Text
		}
	}
	for _, want := range []string{"alpha words stay together", "beta words stay together", "gamma words stay together"} {
		if !strings.Contains(combined, want) {
			t.Fatalf("split chunks lost content containing %q:\n%s", want, combined)
		}
	}
}

func TestPlanSourceSplitsBoundarylessParagraphWithoutDroppingContent(t *testing.T) {
	longToken := strings.Repeat("x", 420)
	source := "# Title\n\n1. Big Section\n\n" + longToken

	plan := PlanSource("raw/source.md", source, 100)
	if len(plan.Chunks) < 4 {
		t.Fatalf("expected boundaryless paragraph to split, got %d chunks", len(plan.Chunks))
	}
	combined := ""
	for _, chunk := range plan.Chunks {
		if strings.Contains(chunk.Text, "Oversized paragraph segment:") {
			parts := strings.SplitN(chunk.Text, "\n\n", 2)
			if len(parts) == 2 {
				combined += parts[1]
			}
		}
	}
	if strings.Count(combined, "x") != len(longToken) {
		t.Fatalf("split chunks dropped content: got %d x chars, want %d", strings.Count(combined, "x"), len(longToken))
	}
}

func TestPlanRendersOutlineAndIndex(t *testing.T) {
	plan := PlanSource("raw/source.md", "# Title\n\n1. Overview\n\nBody", 1000)
	outline := plan.OutlineMarkdown()
	if !strings.Contains(outline, "Chunk 01") || !strings.Contains(outline, "raw/source.md") {
		t.Fatalf("outline missing expected content:\n%s", outline)
	}
	index, err := plan.IndexJSON()
	if err != nil {
		t.Fatalf("IndexJSON returned error: %v", err)
	}
	var decoded struct {
		SourcePath string `json:"source_path"`
		Chunks     []struct {
			Chars               int        `json:"chars"`
			SectionHeadingPaths [][]string `json:"heading_paths"`
		} `json:"chunks"`
	}
	if err := json.Unmarshal(index, &decoded); err != nil {
		t.Fatalf("index is not JSON: %v\n%s", err, index)
	}
	if decoded.SourcePath != "raw/source.md" || len(decoded.Chunks) == 0 || decoded.Chunks[0].Chars == 0 || len(decoded.Chunks[0].SectionHeadingPaths) == 0 {
		t.Fatalf("unexpected index: %#v", decoded)
	}
}
