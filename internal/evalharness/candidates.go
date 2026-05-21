package evalharness

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	candidateFrontmatterRE = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*`)
	candidateTitleRE       = regexp.MustCompile(`(?m)^title:\s*(.+?)\s*$`)
	candidateKindRE        = regexp.MustCompile(`(?m)^kind:\s*(.+?)\s*$`)
	candidateConfidenceRE  = regexp.MustCompile(`(?m)^confidence:\s*(.+?)\s*$`)
	candidateHeadingRE     = regexp.MustCompile(`(?m)^#\s+(.+?)\s*$`)
	candidateTokenRE       = regexp.MustCompile(`[a-z0-9]+`)
	wikilinkRE             = regexp.MustCompile(`\[\[([^\]|#]+)(?:[|#][^\]]*)?\]\]`)
)

type CandidateResult struct {
	Bundle     string                  `json:"bundle"`
	Scores     CandidateAggregateScore `json:"scores"`
	Candidates []CandidateFileResult   `json:"candidates"`
	Trace      []string                `json:"trace"`
}

type CandidateAggregateScore struct {
	Frontmatter    string `json:"frontmatter"`
	Sources        string `json:"sources"`
	Title          string `json:"title"`
	TitleAlignment string `json:"title_alignment"`
	Wikilinks      string `json:"wikilinks"`
	Markdown       string `json:"markdown"`
	Length         string `json:"length"`
	Originality    string `json:"originality"`
	Overall        string `json:"overall"`
}

type CandidateFileResult struct {
	Path   string                  `json:"path"`
	Scores CandidateAggregateScore `json:"scores"`
	Trace  []string                `json:"trace"`
}

// EvaluateCandidates scores generated concept/topic candidate drafts.
func EvaluateCandidates(bundleDir string) (CandidateResult, error) {
	return evaluateCandidates("", bundleDir, false)
}

// EvaluateCandidatesWithRoot scores candidates and verifies wikilinks against
// the current wiki plus concept/topic candidates named in the apply plan.
func EvaluateCandidatesWithRoot(root string, bundleDir string) (CandidateResult, error) {
	return evaluateCandidates(root, bundleDir, true)
}

func evaluateCandidates(root string, bundleDir string, validateLinks bool) (CandidateResult, error) {
	applyPlan, err := os.ReadFile(filepath.Join(bundleDir, "apply-plan.md"))
	if err != nil {
		return CandidateResult{}, fmt.Errorf("read apply plan: %w", err)
	}
	sourceDraft, err := os.ReadFile(filepath.Join(bundleDir, "source.md"))
	if err != nil {
		return CandidateResult{}, fmt.Errorf("read source draft: %w", err)
	}

	result := CandidateResult{Bundle: bundleDir}
	targets := candidateDraftPaths(string(applyPlan))
	allowedLinks := map[string]bool{}
	supportedLinks := map[string]bool{}
	if validateLinks {
		allowedLinks = candidateAllowedLinkSlugs(root, targets)
		supportedLinks = candidateSupportedLinkSlugs(string(sourceDraft), targets, allowedLinks)
	}
	for _, target := range targets {
		path := filepath.Join(bundleDir, "candidates", filepath.FromSlash(target))
		content, err := os.ReadFile(path)
		if err != nil {
			file := CandidateFileResult{Path: target}
			file.Scores.Overall = "fail"
			file.Trace = append(file.Trace, "candidate: missing draft file")
			result.Candidates = append(result.Candidates, file)
			continue
		}
		result.Candidates = append(result.Candidates, scoreCandidateFile(target, string(content), string(sourceDraft), allowedLinks, supportedLinks, validateLinks))
	}
	if len(result.Candidates) == 0 {
		result.Trace = append(result.Trace, "candidates: no concept/topic candidates found")
		result.Scores.Overall = "fail"
		return result, nil
	}

	result.Scores = aggregateCandidateScores(result.Candidates)
	for _, candidate := range result.Candidates {
		for _, entry := range candidate.Trace {
			result.Trace = append(result.Trace, candidate.Path+": "+entry)
		}
	}

	return result, nil
}

func candidateDraftPaths(applyPlan string) []string {
	paths := []string{}
	for _, path := range candidatePaths(applyPlan) {
		if strings.HasPrefix(path, "wiki/concepts/") || strings.HasPrefix(path, "wiki/topics/") {
			paths = append(paths, path)
		}
	}

	return paths
}

func scoreCandidateFile(target string, content string, sourceDraft string, allowedLinks map[string]bool, supportedLinks map[string]bool, validateLinks bool) CandidateFileResult {
	result := CandidateFileResult{Path: target}
	frontmatter, body := splitCandidateFrontmatter(content)
	result.Scores.Frontmatter = scoreCandidateFrontmatter(target, frontmatter, &result.Trace)
	result.Scores.Sources = scoreCandidateSources(frontmatter, &result.Trace)
	result.Scores.Title = scoreCandidateTitle(frontmatter, body, &result.Trace)
	result.Scores.TitleAlignment = scoreCandidateTitleAlignment(frontmatterValue(frontmatter, candidateTitleRE), body, &result.Trace)
	result.Scores.Wikilinks = scoreCandidateWikilinks(body, allowedLinks, supportedLinks, validateLinks, &result.Trace)
	result.Scores.Markdown = scoreCandidateMarkdown(body, &result.Trace)
	result.Scores.Length = scoreCandidateLength(body, &result.Trace)
	result.Scores.Originality = scoreCandidateOriginality(body, sourceDraft, &result.Trace)
	result.Scores.Overall = aggregateOneCandidate(result.Scores)

	return result
}

func splitCandidateFrontmatter(content string) (string, string) {
	match := candidateFrontmatterRE.FindStringSubmatch(content)
	if len(match) != 2 {
		return "", content
	}

	body := candidateFrontmatterRE.ReplaceAllString(content, "")
	return match[1], strings.TrimSpace(body)
}

func scoreCandidateFrontmatter(target string, frontmatter string, trace *[]string) string {
	if frontmatter == "" {
		*trace = append(*trace, "frontmatter: missing YAML frontmatter")
		return "fail"
	}
	title := frontmatterValue(frontmatter, candidateTitleRE)
	kind := frontmatterValue(frontmatter, candidateKindRE)
	confidence := frontmatterValue(frontmatter, candidateConfidenceRE)
	if title == "" || kind == "" || confidence == "" || !strings.Contains(frontmatter, "sources:") {
		*trace = append(*trace, "frontmatter: missing title, kind, sources, or confidence")
		return "fail"
	}
	wantKind := "concept"
	if strings.HasPrefix(target, "wiki/topics/") {
		wantKind = "topic"
	}
	if kind != wantKind {
		*trace = append(*trace, fmt.Sprintf("frontmatter: kind %q does not match target kind %q", kind, wantKind))
		return "fail"
	}
	*trace = append(*trace, "frontmatter: required fields present")
	return "pass"
}

func scoreCandidateSources(frontmatter string, trace *[]string) string {
	if !strings.Contains(frontmatter, "- source.md") {
		*trace = append(*trace, "sources: missing source.md reference")
		return "fail"
	}
	if !strings.Contains(frontmatter, "raw/") && !strings.Contains(frontmatter, "wiki/sources/") {
		*trace = append(*trace, "sources: missing durable raw or source-summary reference")
		return "fail"
	}
	*trace = append(*trace, "sources: source.md and durable source reference present")
	return "pass"
}

func scoreCandidateTitle(frontmatter string, body string, trace *[]string) string {
	title := frontmatterValue(frontmatter, candidateTitleRE)
	if title == "" {
		*trace = append(*trace, "title: missing frontmatter title")
		return "fail"
	}
	match := candidateHeadingRE.FindStringSubmatch(body)
	if len(match) != 2 {
		*trace = append(*trace, "title: missing top-level heading")
		return "fail"
	}
	if strings.TrimSpace(match[1]) != title {
		*trace = append(*trace, "title: heading does not match frontmatter title")
		return "borderline"
	}
	*trace = append(*trace, "title: frontmatter and heading match")
	return "pass"
}

func scoreCandidateTitleAlignment(title string, body string, trace *[]string) string {
	words := significantTitleWords(title)
	if len(words) == 0 {
		*trace = append(*trace, "title_alignment: no significant title words to check")
		return "pass"
	}
	counts := tokenCounts(body)
	hits := 0
	missing := []string{}
	for _, word := range words {
		if counts[word] >= 2 {
			hits++
			continue
		}
		missing = append(missing, word)
	}
	if hits*2 < len(words) {
		*trace = append(*trace, fmt.Sprintf("title_alignment: title terms weakly supported in body: %s", strings.Join(missing, ", ")))
		return "borderline"
	}
	*trace = append(*trace, "title_alignment: title terms are supported by body content")
	return "pass"
}

func scoreCandidateWikilinks(body string, allowedLinks map[string]bool, supportedLinks map[string]bool, validateLinks bool, trace *[]string) string {
	matches := wikilinkRE.FindAllStringSubmatch(body, -1)
	if len(matches) == 0 {
		*trace = append(*trace, "wikilinks: none present; acceptable when no source-supported link is available")
		return "pass"
	}
	if validateLinks {
		missing := []string{}
		weak := []string{}
		for _, match := range matches {
			if len(match) != 2 {
				continue
			}
			slug := slugFromCandidateLink(match[1])
			if !allowedLinks[slug] {
				missing = append(missing, match[1])
				continue
			}
			if !supportedLinks[slug] {
				weak = append(weak, match[1])
			}
		}
		if len(missing) > 0 {
			*trace = append(*trace, "wikilinks: missing targets: "+strings.Join(missing, ", "))
			return "borderline"
		}
		if len(weak) > 0 {
			*trace = append(*trace, "wikilinks: weak semantic targets: "+strings.Join(weak, ", "))
			return "borderline"
		}
	}
	*trace = append(*trace, "wikilinks: at least one wikilink present")
	return "pass"
}

func scoreCandidateMarkdown(body string, trace *[]string) string {
	if strings.Count(body, "```")%2 != 0 {
		*trace = append(*trace, "markdown: unclosed fenced code block")
		return "fail"
	}
	*trace = append(*trace, "markdown: no unclosed fenced code block")
	return "pass"
}

func candidateAllowedLinkSlugs(root string, targets []string) map[string]bool {
	allowed := map[string]bool{}
	for _, target := range targets {
		allowed[strings.TrimSuffix(filepath.Base(target), filepath.Ext(target))] = true
	}
	_ = filepath.WalkDir(filepath.Join(root, "wiki"), func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		allowed[strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))] = true
		return nil
	})
	return allowed
}

func candidateSupportedLinkSlugs(sourceDraft string, targets []string, allowedLinks map[string]bool) map[string]bool {
	supported := map[string]bool{}
	normalizedSource := strings.ToLower(sourceDraft)
	for _, target := range targets {
		slug := strings.TrimSuffix(filepath.Base(target), filepath.Ext(target))
		supported[slug] = true
	}
	for slug := range allowedLinks {
		if strings.Contains(normalizedSource, strings.ReplaceAll(slug, "-", " ")) || strings.Contains(normalizedSource, slug) {
			supported[slug] = true
		}
	}
	return supported
}

func slugFromCandidateLink(link string) string {
	slug := strings.TrimSpace(link)
	slug = strings.TrimSuffix(slug, filepath.Ext(slug))
	slug = strings.ToLower(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

func scoreCandidateLength(body string, trace *[]string) string {
	words := strings.Fields(body)
	switch {
	case len(words) < 30:
		*trace = append(*trace, fmt.Sprintf("length: too short (%d words)", len(words)))
		return "fail"
	case len(words) < 50:
		*trace = append(*trace, fmt.Sprintf("length: short but reviewable (%d words)", len(words)))
		return "borderline"
	default:
		*trace = append(*trace, fmt.Sprintf("length: sufficient (%d words)", len(words)))
		return "pass"
	}
}

func scoreCandidateOriginality(body string, sourceDraft string, trace *[]string) string {
	bodyLines := meaningfulLines(body)
	if len(bodyLines) == 0 {
		*trace = append(*trace, "originality: no meaningful body lines")
		return "fail"
	}
	sourceShingles := tokenShingles(tokenizeForSimilarity(sourceDraft), 5)
	tooClose := 0
	for _, line := range bodyLines {
		if strings.Contains(sourceDraft, line) || shingleOverlap(line, sourceShingles, 5) >= 0.75 {
			tooClose++
		}
	}
	if tooClose >= 2 || tooClose*2 >= len(bodyLines) {
		*trace = append(*trace, fmt.Sprintf("originality: %d of %d meaningful lines are too close to source", tooClose, len(bodyLines)))
		return "borderline"
	}
	*trace = append(*trace, "originality: not mostly copied from source")
	return "pass"
}

func meaningfulLines(content string) []string {
	lines := []string{}
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "-") || strings.HasPrefix(line, "*") {
			continue
		}
		if len(line) < 50 {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

func significantTitleWords(title string) []string {
	stop := map[string]bool{
		"a": true, "an": true, "and": true, "as": true, "for": true,
		"in": true, "of": true, "on": true, "or": true, "the": true,
		"to": true, "with": true,
	}
	seen := map[string]bool{}
	words := []string{}
	for _, token := range tokenizeForSimilarity(title) {
		if stop[token] || (len(token) < 3 && token != "go") || seen[token] {
			continue
		}
		seen[token] = true
		words = append(words, token)
	}
	return words
}

func tokenCounts(content string) map[string]int {
	counts := map[string]int{}
	for _, token := range tokenizeForSimilarity(content) {
		counts[token]++
	}
	return counts
}

func tokenizeForSimilarity(content string) []string {
	return candidateTokenRE.FindAllString(strings.ToLower(content), -1)
}

func tokenShingles(tokens []string, size int) map[string]bool {
	shingles := map[string]bool{}
	if size <= 0 || len(tokens) < size {
		return shingles
	}
	for i := 0; i <= len(tokens)-size; i++ {
		shingles[strings.Join(tokens[i:i+size], " ")] = true
	}
	return shingles
}

func shingleOverlap(line string, sourceShingles map[string]bool, size int) float64 {
	lineShingles := tokenShingles(tokenizeForSimilarity(line), size)
	if len(lineShingles) == 0 || len(sourceShingles) == 0 {
		return 0
	}
	matches := 0
	for shingle := range lineShingles {
		if sourceShingles[shingle] {
			matches++
		}
	}
	return float64(matches) / float64(len(lineShingles))
}

func aggregateCandidateScores(candidates []CandidateFileResult) CandidateAggregateScore {
	score := CandidateAggregateScore{
		Frontmatter:    "pass",
		Sources:        "pass",
		Title:          "pass",
		TitleAlignment: "pass",
		Wikilinks:      "pass",
		Markdown:       "pass",
		Length:         "pass",
		Originality:    "pass",
		Overall:        "pass",
	}
	for _, candidate := range candidates {
		score.Frontmatter = worstScore(score.Frontmatter, candidate.Scores.Frontmatter)
		score.Sources = worstScore(score.Sources, candidate.Scores.Sources)
		score.Title = worstScore(score.Title, candidate.Scores.Title)
		score.TitleAlignment = worstScore(score.TitleAlignment, candidate.Scores.TitleAlignment)
		score.Wikilinks = worstScore(score.Wikilinks, candidate.Scores.Wikilinks)
		score.Markdown = worstScore(score.Markdown, candidate.Scores.Markdown)
		score.Length = worstScore(score.Length, candidate.Scores.Length)
		score.Originality = worstScore(score.Originality, candidate.Scores.Originality)
		score.Overall = worstScore(score.Overall, candidate.Scores.Overall)
	}
	return score
}

func aggregateOneCandidate(scores CandidateAggregateScore) string {
	overall := "pass"
	for _, score := range []string{scores.Frontmatter, scores.Sources, scores.Title, scores.TitleAlignment, scores.Wikilinks, scores.Markdown, scores.Length, scores.Originality} {
		overall = worstScore(overall, score)
	}
	return overall
}

func worstScore(a string, b string) string {
	if a == "fail" || b == "fail" {
		return "fail"
	}
	if a == "borderline" || b == "borderline" {
		return "borderline"
	}
	return "pass"
}
