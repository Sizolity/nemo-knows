package evalharness

import (
	"strings"
	"testing"
)

func TestReviewCandidatesSuggestsRepairsForWeakSemanticLinks(t *testing.T) {
	result := CandidateResult{
		Bundle: "drafts/web-e2e-sqlite",
		Scores: CandidateAggregateScore{
			Overall:   "borderline",
			Wikilinks: "borderline",
		},
		Candidates: []CandidateFileResult{
			{
				Path: "wiki/concepts/database-journal-modes.md",
				Scores: CandidateAggregateScore{
					Overall:   "borderline",
					Wikilinks: "borderline",
				},
				Trace: []string{
					"frontmatter: required fields present",
					"wikilinks: weak semantic targets: index, query",
				},
			},
		},
	}

	review := ReviewCandidates(result)
	if review.Overall != "needs-review" {
		t.Fatalf("overall = %q, want needs-review", review.Overall)
	}
	if len(review.Items) != 1 {
		t.Fatalf("review item count = %d, want 1", len(review.Items))
	}
	item := review.Items[0]
	if item.Path != "wiki/concepts/database-journal-modes.md" {
		t.Fatalf("path = %q", item.Path)
	}
	if !strings.Contains(item.Problem, "weak semantic targets: index, query") {
		t.Fatalf("problem should describe weak links, got %q", item.Problem)
	}
	if !strings.Contains(item.Recommendation, "Convert unsupported wikilinks to plain text") {
		t.Fatalf("recommendation should suggest plain text repair, got %q", item.Recommendation)
	}
}

func TestReviewCandidatesSuggestsRepairsForTitleAlignment(t *testing.T) {
	result := CandidateResult{
		Bundle: "drafts/dom",
		Scores: CandidateAggregateScore{
			Overall:        "borderline",
			TitleAlignment: "borderline",
		},
		Candidates: []CandidateFileResult{
			{
				Path: "wiki/topics/browser-compatibility-matrix.md",
				Scores: CandidateAggregateScore{
					Overall:        "borderline",
					TitleAlignment: "borderline",
				},
				Trace: []string{
					"title_alignment: title terms weakly supported in body: compatibility, matrix",
				},
			},
		},
	}

	review := ReviewCandidates(result)
	if review.Overall != "needs-review" {
		t.Fatalf("overall = %q, want needs-review", review.Overall)
	}
	if len(review.Items) != 1 {
		t.Fatalf("review item count = %d, want 1", len(review.Items))
	}
	if !strings.Contains(review.Items[0].Recommendation, "body directly supports the page title") {
		t.Fatalf("recommendation should suggest title/body repair, got %q", review.Items[0].Recommendation)
	}
}

func TestRenderCandidateReviewIncludesManualBoundary(t *testing.T) {
	review := CandidateReview{
		Bundle:  "drafts/web-e2e-sqlite",
		Overall: "needs-review",
		Items: []CandidateReviewItem{
			{
				Path:           "wiki/concepts/database-journal-modes.md",
				Severity:       "borderline",
				Problem:        "wikilinks: weak semantic targets: index, query",
				Recommendation: "Convert unsupported wikilinks to plain text unless source support is added.",
			},
		},
	}

	rendered := RenderCandidateReview(review)
	for _, want := range []string{
		"# Candidate Review",
		"Bundle: `drafts/web-e2e-sqlite`",
		"Do not apply these suggestions automatically.",
		"`wiki/concepts/database-journal-modes.md`",
		"Convert unsupported wikilinks to plain text",
	} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("rendered review missing %q:\n%s", want, rendered)
		}
	}
}
