package evalharness

import (
	"fmt"
	"strings"
)

type CandidateReview struct {
	Bundle  string
	Overall string
	Items   []CandidateReviewItem
}

type CandidateReviewItem struct {
	Path           string
	Severity       string
	Problem        string
	Recommendation string
}

func ReviewCandidates(result CandidateResult) CandidateReview {
	review := CandidateReview{
		Bundle:  result.Bundle,
		Overall: "pass",
	}
	for _, candidate := range result.Candidates {
		if candidate.Scores.Overall == "pass" {
			continue
		}
		for _, entry := range candidate.Trace {
			recommendation := candidateReviewRecommendation(entry)
			if recommendation == "" {
				continue
			}
			review.Items = append(review.Items, CandidateReviewItem{
				Path:           candidate.Path,
				Severity:       candidate.Scores.Overall,
				Problem:        entry,
				Recommendation: recommendation,
			})
		}
	}
	if len(review.Items) > 0 {
		review.Overall = "needs-review"
	}
	return review
}

func candidateReviewRecommendation(problem string) string {
	switch {
	case strings.Contains(problem, "wikilinks: weak semantic targets:"):
		return "Convert unsupported wikilinks to plain text unless source support is added or the reviewed candidate set explicitly justifies the link."
	case strings.Contains(problem, "wikilinks: missing targets:"):
		return "Remove missing wikilinks or add explicit reviewed candidate pages before approved apply."
	case strings.Contains(problem, "wikilinks: no wikilinks found"):
		return "Add one source-supported wikilink to an existing page or reviewed candidate; otherwise keep the term as plain text and accept the borderline score."
	case strings.Contains(problem, "markdown: unclosed fenced code block"):
		return "Fix malformed Markdown before approved apply; unclosed fences usually mean the model wrapped the page in a code block."
	case strings.Contains(problem, "sources: missing"):
		return "Regenerate after adding a durable raw or wiki/sources reference; candidate pages must not rely only on local source.md."
	case strings.Contains(problem, "title_alignment: title terms weakly supported"):
		return "Rewrite the draft so the body directly supports the page title, or rename the candidate to match what the source-backed body actually covers."
	case strings.Contains(problem, "title: missing") || strings.Contains(problem, "title: heading does not match"):
		return "Normalize the frontmatter title and top-level heading to match the reviewed target path."
	case strings.Contains(problem, "frontmatter: missing") || strings.Contains(problem, "frontmatter: kind"):
		return "Restore canonical candidate frontmatter: title, kind, sources, and confidence."
	case strings.Contains(problem, "length: too short") || strings.Contains(problem, "length: short"):
		return "Expand the draft with concise source-backed prose before approved apply."
	case strings.Contains(problem, "originality:") && !strings.Contains(problem, "not mostly copied"):
		return "Rewrite copied passages in original reference prose while preserving source-backed claims."
	case strings.Contains(problem, "candidate: missing draft file"):
		return "Regenerate candidates for the reviewed apply plan before approved apply."
	default:
		return ""
	}
}

func RenderCandidateReview(review CandidateReview) string {
	var b strings.Builder
	b.WriteString("# Candidate Review\n\n")
	b.WriteString(fmt.Sprintf("Bundle: `%s`\n\n", review.Bundle))
	b.WriteString("This is a review artifact. Do not apply these suggestions automatically.\n\n")
	b.WriteString(fmt.Sprintf("- overall: `%s`\n", review.Overall))
	b.WriteString(fmt.Sprintf("- items: `%d`\n\n", len(review.Items)))
	b.WriteString("## Suggested Repairs\n\n")
	if len(review.Items) == 0 {
		b.WriteString("(none)\n")
		return b.String()
	}
	for _, item := range review.Items {
		b.WriteString(fmt.Sprintf("### `%s`\n\n", item.Path))
		b.WriteString(fmt.Sprintf("- severity: `%s`\n", item.Severity))
		b.WriteString("- problem: ")
		b.WriteString(item.Problem)
		b.WriteByte('\n')
		b.WriteString("- recommendation: ")
		b.WriteString(item.Recommendation)
		b.WriteString("\n\n")
	}
	return b.String()
}
