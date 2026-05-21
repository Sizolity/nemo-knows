package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/huic/nemo-knows/internal/apply"
	"github.com/huic/nemo-knows/internal/chunking"
	"github.com/huic/nemo-knows/internal/config"
	"github.com/huic/nemo-knows/internal/deepseek"
	"github.com/huic/nemo-knows/internal/draft"
	"github.com/huic/nemo-knows/internal/evalharness"
	"github.com/huic/nemo-knows/internal/llama"
	"github.com/huic/nemo-knows/internal/prompt"
	"github.com/huic/nemo-knows/internal/review"
	"github.com/huic/nemo-knows/internal/web"
	"github.com/huic/nemo-knows/internal/wikilint"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	fs := flag.NewFlagSet("nemo", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	source := fs.String("source", "", "path to a raw source file")
	promptPath := fs.String("prompt", "", "path to a prompt template")
	out := fs.String("out", "", "path to the cleaned Markdown draft")
	outDir := fs.String("out-dir", "", "directory for command output artifacts")
	bundleDir := fs.String("bundle-dir", "", "directory for a local ingest draft bundle")
	reviewBundle := fs.String("review-bundle", "", "directory for a local ingest draft bundle to review")
	generateCandidates := fs.String("generate-candidates", "", "directory for a reviewed bundle whose concept/topic candidate drafts should be generated")
	evalCandidates := fs.String("eval-candidates", "", "directory for a reviewed bundle whose concept/topic candidate drafts should be evaluated")
	reviewCandidates := fs.String("review-candidates", "", "directory for a reviewed bundle whose concept/topic candidate drafts should be reviewed")
	llmReviewCandidates := fs.String("llm-review-candidates", "", "directory for candidate drafts to review with the configured model")
	lintBundle := fs.String("lint-bundle", "", "directory for reviewed candidate drafts to crosslink-lint")
	evalBundle := fs.String("eval-bundle", "", "directory for a reviewed ingest bundle to evaluate")
	evalRegression := fs.String("eval-regression", "", "directory containing eval regression cases")
	resumeBundle := fs.String("resume", "", "resume a reviewed bundle pipeline by running missing post-bundle stages")
	lintWiki := fs.Bool("lint-wiki", false, "run deterministic read-only lint checks over wiki/")
	applyApproved := fs.String("apply-approved", "", "directory for an approved reviewed bundle to apply")
	approve := fs.Bool("approve", false, "explicitly approve wiki writes for apply mode")
	forceApply := fs.Bool("force-apply", false, "allow re-applying a bundle that already has an ingest log entry")
	persistRawWeb := fs.Bool("persist-raw-web", false, "copy the input source to raw/web/<slug>.md before bundle generation")
	profile := fs.String("profile", "stable", "generation profile: fast, stable, deep, or fallback")
	provider := fs.String("provider", "", "generation backend override: llama or deepseek (wins over .env)")
	serve := fs.Bool("serve", false, "start the local nemo-knows web console")
	addr := fs.String("addr", "127.0.0.1:8787", "address for -serve mode")

	if err := fs.Parse(args); err != nil {
		return 2
	}

	cfg, err := config.ForProfileWithProvider(*profile, *provider)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if *serve {
		if err := web.Run(*addr, cfg, inProcessPipeline{}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *applyApproved != "" {
		if err := runApplyApproved(*applyApproved, *approve, *forceApply); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *resumeBundle != "" {
		if *outDir == "" {
			fmt.Fprintln(os.Stderr, "required flags for resume mode: -resume, -out-dir")
			return 2
		}
		if err := runResumeBundle(*resumeBundle, *outDir, cfg); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *generateCandidates != "" {
		if err := runGenerateCandidates(*generateCandidates, cfg); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *evalCandidates != "" {
		if *outDir == "" {
			fmt.Fprintln(os.Stderr, "required flags for candidate eval mode: -eval-candidates, -out-dir")
			return 2
		}
		if err := runEvalCandidates(*evalCandidates, *outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *reviewCandidates != "" {
		if *outDir == "" {
			fmt.Fprintln(os.Stderr, "required flags for candidate review mode: -review-candidates, -out-dir")
			return 2
		}
		if err := runReviewCandidates(*reviewCandidates, *outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *llmReviewCandidates != "" {
		if *outDir == "" {
			fmt.Fprintln(os.Stderr, "required flags for LLM candidate review mode: -llm-review-candidates, -out-dir")
			return 2
		}
		if err := runLLMReviewCandidates(*llmReviewCandidates, *outDir, cfg); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *lintBundle != "" {
		if *outDir == "" {
			fmt.Fprintln(os.Stderr, "required flags for bundle lint mode: -lint-bundle, -out-dir")
			return 2
		}
		if err := runLintBundle(*lintBundle, *outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *lintWiki {
		if *outDir == "" {
			fmt.Fprintln(os.Stderr, "required flags for wiki lint mode: -lint-wiki, -out-dir")
			return 2
		}
		if err := runLintWiki(*outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *evalRegression != "" {
		if *outDir == "" {
			fmt.Fprintln(os.Stderr, "required flags for regression eval mode: -eval-regression, -out-dir")
			return 2
		}
		if err := runEvalRegression(*evalRegression, *outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *evalBundle != "" {
		if *outDir == "" {
			fmt.Fprintln(os.Stderr, "required flags for eval mode: -eval-bundle, -out-dir")
			return 2
		}
		if err := runEvalBundle(*evalBundle, *outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *reviewBundle != "" {
		if *out == "" {
			fmt.Fprintln(os.Stderr, "required flags for review mode: -review-bundle, -out")
			return 2
		}
		if err := runReviewBundle(*reviewBundle, *out); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}
	if *bundleDir != "" {
		if *source == "" {
			fmt.Fprintln(os.Stderr, "required flags for bundle mode: -source, -bundle-dir")
			return 2
		}
		bundleSource := *source
		if *persistRawWeb {
			persisted, err := persistRawWebSource(*source)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return 1
			}
			bundleSource = persisted
		}
		if err := runBundle(bundleSource, *bundleDir, cfg); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	}

	if *source == "" || *promptPath == "" || *out == "" {
		fmt.Fprintln(os.Stderr, "required flags: -source, -prompt, -out")
		return 2
	}
	if err := runDraft(*source, *promptPath, *out, cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

var candidatePlanLineRE = regexp.MustCompile("(?m)^- `([^`]+)` — .+$")
var markdownFrontmatterRE = regexp.MustCompile(`(?s)^---\s*\n.*?\n---\s*\n?`)
var nestedFrontmatterPreludeRE = regexp.MustCompile(`(?s)^\s*(?:title|kind|sources|confidence):.*?\n---\s*\n?`)
var sourceReferenceLineRE = regexp.MustCompile(`(?m)^\s*-\s*(raw/[^ \n]+|wiki/sources/[^ \n]+)\s*$`)
var sourceReferenceRE = regexp.MustCompile(`(?:raw|wiki/sources)/[A-Za-z0-9._/-]*[A-Za-z0-9_-]\.md`)
var candidateWikilinkRE = regexp.MustCompile(`\[\[([^\]|#]+)(?:[|#][^\]]*)?\]\]`)
var markdownHeadingRE = regexp.MustCompile(`(?m)^#\s+.+$`)
var sourceDraftFrontmatterRE = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*`)

// chunkNotesPerGroup is the minimum chunk count that triggers the middle
// group-notes layer. The chunked-bundle character threshold and per-chunk size
// cap are configurable per provider via config.Config; see cfg.ChunkedBundleCharThreshold
// and cfg.MaxChunkChars.
const chunkNotesPerGroup = 6

type candidateDraftTarget struct {
	Path  string
	Kind  string
	Title string
}

func runGenerateCandidates(bundleDir string, cfg config.Config) error {
	applyPlan, err := os.ReadFile(filepath.Join(bundleDir, "apply-plan.md"))
	if err != nil {
		return fmt.Errorf("read apply plan: %w", err)
	}
	sourceDraft, err := os.ReadFile(filepath.Join(bundleDir, "source.md"))
	if err != nil {
		return fmt.Errorf("read source draft: %w", err)
	}
	sourceRefs := sourceRefsForCandidate(sourceDraft)

	targets := candidateDraftTargets(string(applyPlan))
	allowedLinks := allowedLinkSlugs(string(sourceDraft), targets)
	for _, target := range targets {
		promptPath := filepath.Join("prompts", target.Kind+"-page.md")
		out := filepath.Join(bundleDir, "candidates", filepath.FromSlash(target.Path))
		if err := runCandidateDraft(promptPath, out, target, string(sourceDraft), sourceRefs, allowedLinks, cfg); err != nil {
			return err
		}
	}

	fmt.Fprintf(os.Stderr, "generated %d candidate drafts in %s\n", len(targets), filepath.Join(bundleDir, "candidates"))
	return nil
}

func candidateDraftTargets(applyPlan string) []candidateDraftTarget {
	matches := candidatePlanLineRE.FindAllStringSubmatch(applyPlan, -1)
	targets := make([]candidateDraftTarget, 0, len(matches))
	for _, match := range matches {
		path := match[1]
		kind := ""
		switch {
		case strings.HasPrefix(path, "wiki/concepts/"):
			kind = "concept"
		case strings.HasPrefix(path, "wiki/topics/"):
			kind = "topic"
		default:
			continue
		}
		targets = append(targets, candidateDraftTarget{
			Path:  path,
			Kind:  kind,
			Title: titleFromWikiPath(path),
		})
	}

	return targets
}

func titleFromWikiPath(path string) string {
	slug := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	parts := strings.Split(slug, "-")
	for i, part := range parts {
		if part == "" {
			continue
		}
		if acronym := knownTitleAcronym(part); acronym != "" {
			parts[i] = acronym
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}

	return strings.Join(parts, " ")
}

func knownTitleAcronym(part string) string {
	switch strings.ToLower(part) {
	case "llm":
		return "LLM"
	case "rag":
		return "RAG"
	case "mvp":
		return "MVP"
	default:
		return ""
	}
}

func runCandidateDraft(promptPath string, out string, target candidateDraftTarget, sourceContent string, sourceRefs []string, allowedLinks map[string]bool, cfg config.Config) error {
	templateContent, err := os.ReadFile(promptPath)
	if err != nil {
		return fmt.Errorf("read candidate prompt: %w", err)
	}
	targetEvidence := targetEvidenceForCandidate(target, sourceRefs)
	rendered, err := prompt.Render(string(templateContent), prompt.Variables{
		ConceptName:    target.Title,
		SourceList:     markdownSourceList(sourceRefs),
		SourceContent:  sourceContent,
		TargetEvidence: targetEvidence,
		PageTitle:      target.Title,
		PageKind:       target.Kind,
		TargetPath:     target.Path,
		AllowedLinks:   markdownAllowedLinks(allowedLinks),
	})
	if err != nil {
		return fmt.Errorf("render candidate prompt: %w", err)
	}

	generator := generatorFromConfig(cfg)

	rawOutput, err := generator.Generate(context.Background(), rendered)
	if err != nil {
		return fmt.Errorf("generate candidate draft: %w", err)
	}

	paths := draft.PathsFor(out)
	if err := os.MkdirAll(filepath.Dir(paths.Cleaned), 0o755); err != nil {
		return fmt.Errorf("create candidate draft directory: %w", err)
	}
	if err := os.WriteFile(paths.Raw, []byte(rawOutput), 0o644); err != nil {
		return fmt.Errorf("write raw candidate draft: %w", err)
	}
	cleaned, err := draft.Clean(rawOutput)
	if err != nil {
		cleaned, err = runFallbackCandidateDraft(rendered, paths.Cleaned, cfg)
		if err != nil {
			return fmt.Errorf("clean candidate draft and fallback for %s: %w", target.Path, err)
		}
	} else if err := os.Remove(fallbackRawPath(paths.Cleaned)); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("remove stale fallback raw candidate draft: %w", err)
	}
	cleaned = normalizeCandidateDraft(cleaned, target, sourceRefs, allowedLinks)
	if err := os.WriteFile(paths.Cleaned, []byte(cleaned), 0o644); err != nil {
		return fmt.Errorf("write candidate draft: %w", err)
	}

	return nil
}

func runFallbackCandidateDraft(renderedPrompt string, cleanedPath string, cfg config.Config) (string, error) {
	if cfg.Profile == "fallback" {
		return "", fmt.Errorf("clean candidate draft: %w", draft.ErrNoFrontmatter)
	}
	fallback, err := config.ForProfile("fallback")
	if err != nil {
		return "", fmt.Errorf("load candidate fallback profile: %w", err)
	}
	copyBackendConfig(&fallback, cfg)

	generator := generatorFromConfig(fallback)
	rawOutput, err := generator.Generate(context.Background(), renderedPrompt)
	if err != nil {
		return "", fmt.Errorf("generate fallback candidate draft: %w", err)
	}
	if err := os.WriteFile(fallbackRawPath(cleanedPath), []byte(rawOutput), 0o644); err != nil {
		return "", fmt.Errorf("write fallback raw candidate draft: %w", err)
	}
	cleaned, err := draft.Clean(rawOutput)
	if err != nil {
		return "", fmt.Errorf("clean fallback candidate draft: %w", err)
	}

	return cleaned, nil
}

func sourceRefsForCandidate(sourceDraft []byte) []string {
	refs := []string{"source.md"}
	for _, match := range sourceReferenceLineRE.FindAllSubmatch(sourceDraft, -1) {
		ref := strings.TrimSpace(string(match[1]))
		if !containsString(refs, ref) {
			refs = append(refs, ref)
		}
	}
	for _, match := range sourceReferenceRE.FindAll(sourceDraft, -1) {
		ref := strings.TrimSpace(string(match))
		if !containsString(refs, ref) {
			refs = append(refs, ref)
		}
	}

	return refs
}

func markdownSourceList(refs []string) string {
	var b strings.Builder
	for _, ref := range refs {
		b.WriteString("- " + ref + "\n")
	}

	return strings.TrimRight(b.String(), "\n")
}

func targetEvidenceForCandidate(target candidateDraftTarget, sourceRefs []string) string {
	rawPath := ""
	for _, ref := range sourceRefs {
		if strings.HasPrefix(ref, "raw/") {
			rawPath = ref
			break
		}
	}
	if rawPath == "" {
		return "(none)"
	}
	content, err := os.ReadFile(filepath.Clean(rawPath))
	if err != nil {
		return "(raw source unavailable for targeted evidence)"
	}
	excerpts := targetEvidenceExcerpts(string(content), target.Title, 4, 1200)
	if len(excerpts) == 0 {
		return "(no direct target-title matches found in raw source)"
	}
	var b strings.Builder
	for i, excerpt := range excerpts {
		b.WriteString(fmt.Sprintf("### Excerpt %d from `%s`\n\n", i+1, filepath.ToSlash(rawPath)))
		b.WriteString(excerpt)
		b.WriteString("\n\n")
	}
	return strings.TrimSpace(b.String())
}

func targetEvidenceExcerpts(content string, title string, maxExcerpts int, windowChars int) []string {
	terms := evidenceTerms(title)
	if len(terms) == 0 || maxExcerpts <= 0 || windowChars <= 0 {
		return nil
	}
	lower := strings.ToLower(content)
	excerpts := []string{}
	seenRanges := [][2]int{}
	for _, term := range terms {
		searchFrom := 0
		for len(excerpts) < maxExcerpts {
			idx := strings.Index(lower[searchFrom:], term)
			if idx == -1 {
				break
			}
			center := searchFrom + idx
			start := center - windowChars/2
			if start < 0 {
				start = 0
			}
			end := center + windowChars/2
			if end > len(content) {
				end = len(content)
			}
			start = adjustExcerptStart(content, start)
			end = adjustExcerptEnd(content, end)
			if !rangeOverlaps(seenRanges, start, end) {
				seenRanges = append(seenRanges, [2]int{start, end})
				excerpts = append(excerpts, strings.TrimSpace(content[start:end]))
			}
			searchFrom = center + len(term)
			if searchFrom >= len(lower) {
				break
			}
		}
		if len(excerpts) >= maxExcerpts {
			break
		}
	}
	return excerpts
}

func evidenceTerms(title string) []string {
	stop := map[string]bool{
		"a": true, "an": true, "and": true, "as": true, "for": true,
		"in": true, "of": true, "on": true, "or": true, "the": true,
		"to": true, "with": true,
	}
	parts := strings.FieldsFunc(strings.ToLower(title), func(r rune) bool {
		return !(r >= 'a' && r <= 'z' || r >= '0' && r <= '9')
	})
	terms := []string{}
	seen := map[string]bool{}
	for _, part := range parts {
		if stop[part] || len(part) < 4 || seen[part] {
			continue
		}
		seen[part] = true
		terms = append(terms, part)
	}
	return terms
}

func rangeOverlaps(ranges [][2]int, start int, end int) bool {
	for _, existing := range ranges {
		if start < existing[1] && end > existing[0] {
			return true
		}
	}
	return false
}

func adjustExcerptStart(content string, start int) int {
	for start > 0 && content[start-1] != '\n' {
		start--
	}
	return start
}

func adjustExcerptEnd(content string, end int) int {
	for end < len(content) && content[end-1] != '\n' {
		end++
	}
	return end
}

func normalizeCandidateDraft(cleaned string, target candidateDraftTarget, sourceRefs []string, allowedLinks map[string]bool) string {
	body := markdownFrontmatterRE.ReplaceAllString(cleaned, "")
	body = nestedFrontmatterPreludeRE.ReplaceAllString(body, "")
	body = strings.TrimSpace(body)
	body = normalizeCandidateWikilinks(body, allowedLinks)
	body = normalizeCandidateHeading(body, target.Title)

	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("title: " + target.Title + "\n")
	b.WriteString("kind: " + target.Kind + "\n")
	b.WriteString("sources:\n")
	for _, ref := range sourceRefs {
		b.WriteString("  - " + ref + "\n")
	}
	b.WriteString("confidence: medium\n")
	b.WriteString("---\n\n")
	b.WriteString(body)
	b.WriteString("\n")

	return b.String()
}

func allowedLinkSlugs(sourceContent string, targets []candidateDraftTarget) map[string]bool {
	allowed := map[string]bool{}
	for _, target := range targets {
		allowed[slugFromWikiPath(target.Path)] = true
	}
	_ = filepath.WalkDir("wiki", func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		slug := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		if sourceSupportsLinkSlug(sourceContent, slug) {
			allowed[slug] = true
		}
		return nil
	})
	return allowed
}

func sourceSupportsLinkSlug(sourceContent string, slug string) bool {
	normalized := strings.ToLower(sourceContent)
	return strings.Contains(normalized, slug) || strings.Contains(normalized, strings.ToLower(titleFromSlug(slug)))
}

func markdownAllowedLinks(allowed map[string]bool) string {
	if len(allowed) == 0 {
		return "(none)"
	}
	slugs := make([]string, 0, len(allowed))
	for slug := range allowed {
		slugs = append(slugs, slug)
	}
	sort.Strings(slugs)
	var b strings.Builder
	for _, slug := range slugs {
		b.WriteString("- [[" + slug + "]]\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

func normalizeCandidateWikilinks(body string, allowed map[string]bool) string {
	return candidateWikilinkRE.ReplaceAllStringFunc(body, func(link string) string {
		match := candidateWikilinkRE.FindStringSubmatch(link)
		if len(match) != 2 {
			return link
		}
		target := match[1]
		if allowed[slugFromLink(target)] {
			return link
		}
		return target
	})
}

func normalizeCandidateHeading(body string, title string) string {
	heading := "# " + title
	if markdownHeadingRE.MatchString(body) {
		return markdownHeadingRE.ReplaceAllString(body, heading)
	}
	if body == "" {
		return heading
	}
	return heading + "\n\n" + body
}

func ensureCandidateWikilink(body string, allowed map[string]bool) string {
	if candidateWikilinkRE.MatchString(body) || len(allowed) == 0 {
		return body
	}
	slugs := make([]string, 0, len(allowed))
	for slug := range allowed {
		slugs = append(slugs, slug)
	}
	sort.Slice(slugs, func(i int, j int) bool {
		if len(slugs[i]) == len(slugs[j]) {
			return slugs[i] < slugs[j]
		}
		return len(slugs[i]) > len(slugs[j])
	})
	for _, slug := range slugs {
		phrase := titleFromSlug(slug)
		re := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(phrase) + `\b`)
		lines := strings.Split(body, "\n")
		for i, line := range lines {
			if strings.HasPrefix(strings.TrimSpace(line), "#") || !re.MatchString(line) {
				continue
			}
			lines[i] = re.ReplaceAllStringFunc(line, func(match string) string {
				return "[[" + slug + "|" + match + "]]"
			})
			return strings.Join(lines, "\n")
		}
	}
	return body
}

func titleFromSlug(slug string) string {
	parts := strings.Split(slug, "-")
	for i, part := range parts {
		if acronym := knownTitleAcronym(part); acronym != "" {
			parts[i] = acronym
			continue
		}
		if part != "" {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, " ")
}

func slugFromWikiPath(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

func slugFromLink(link string) string {
	slug := strings.TrimSpace(link)
	slug = strings.TrimSuffix(slug, filepath.Ext(slug))
	slug = strings.ToLower(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

func containsString(items []string, want string) bool {
	for _, item := range items {
		if item == want {
			return true
		}
	}

	return false
}

func runApplyApproved(bundleDir string, approve bool, force bool) error {
	_, err := apply.ApplyApproved(".", bundleDir, apply.Options{Approve: approve, Force: force})
	if err != nil {
		return fmt.Errorf("apply approved bundle: %w", err)
	}

	fmt.Fprintf(os.Stderr, "applied %s\n", bundleDir)
	return nil
}

func runResumeBundle(bundleDir string, outDir string, cfg config.Config) error {
	if !pathExists(filepath.Join(bundleDir, "source.md")) || !pathExists(filepath.Join(bundleDir, "ingest-plan.md")) {
		return fmt.Errorf("resume currently requires existing source.md and ingest-plan.md in %s", bundleDir)
	}
	applyPlan := filepath.Join(bundleDir, "apply-plan.md")
	if !pathExists(applyPlan) {
		if err := runReviewBundle(bundleDir, applyPlan); err != nil {
			return err
		}
	}
	if !pathExists(filepath.Join(outDir, "bundle", "scores.json")) {
		if err := runEvalBundle(bundleDir, filepath.Join(outDir, "bundle")); err != nil {
			return err
		}
	}
	if !hasCandidateDrafts(bundleDir) {
		if err := runGenerateCandidates(bundleDir, cfg); err != nil {
			return err
		}
	}
	if !pathExists(filepath.Join(outDir, "candidates", "candidate-scores.json")) {
		if err := runEvalCandidates(bundleDir, filepath.Join(outDir, "candidates")); err != nil {
			return err
		}
	}
	if !pathExists(filepath.Join(outDir, "review", "candidate-review.md")) {
		if err := runReviewCandidates(bundleDir, filepath.Join(outDir, "review")); err != nil {
			return err
		}
	}
	fmt.Fprintf(os.Stderr, "resumed %s through %s\n", bundleDir, outDir)
	return nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func hasCandidateDrafts(bundleDir string) bool {
	root := filepath.Join(bundleDir, "candidates", "wiki")
	found := false
	_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			found = true
			return filepath.SkipAll
		}
		return nil
	})
	return found
}

func runEvalCandidates(bundleDir string, outDir string) error {
	result, err := evalharness.EvaluateCandidatesWithRoot(".", bundleDir)
	if err != nil {
		return fmt.Errorf("evaluate candidate drafts: %w", err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create candidate eval output directory: %w", err)
	}
	scores, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("encode candidate scores: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "candidate-scores.json"), append(scores, '\n'), 0o644); err != nil {
		return fmt.Errorf("write candidate scores: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "candidate-trace.md"), []byte(renderCandidateEvalTrace(result)), 0o644); err != nil {
		return fmt.Errorf("write candidate trace: %w", err)
	}

	fmt.Fprintf(os.Stderr, "wrote %s and %s\n", filepath.Join(outDir, "candidate-scores.json"), filepath.Join(outDir, "candidate-trace.md"))
	return nil
}

func runReviewCandidates(bundleDir string, outDir string) error {
	result, err := evalharness.EvaluateCandidatesWithRoot(".", bundleDir)
	if err != nil {
		return fmt.Errorf("review candidate drafts: %w", err)
	}
	review := evalharness.ReviewCandidates(result)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create candidate review output directory: %w", err)
	}
	path := filepath.Join(outDir, "candidate-review.md")
	if err := os.WriteFile(path, []byte(evalharness.RenderCandidateReview(review)), 0o644); err != nil {
		return fmt.Errorf("write candidate review: %w", err)
	}

	fmt.Fprintf(os.Stderr, "wrote %s\n", path)
	return nil
}

type llmCandidateReviewResult struct {
	Bundle  string               `json:"bundle"`
	Reviews []llmCandidateReview `json:"reviews"`
}

type llmCandidateReview struct {
	Path   string `json:"path"`
	Review string `json:"review"`
}

func runLLMReviewCandidates(bundleDir string, outDir string, cfg config.Config) error {
	applyPlan, err := os.ReadFile(filepath.Join(bundleDir, "apply-plan.md"))
	if err != nil {
		return fmt.Errorf("read apply plan: %w", err)
	}
	sourceDraft, err := os.ReadFile(filepath.Join(bundleDir, "source.md"))
	if err != nil {
		return fmt.Errorf("read source draft: %w", err)
	}
	result := llmCandidateReviewResult{Bundle: bundleDir}
	generator := generatorFromConfig(cfg)
	for _, target := range candidateDraftTargets(string(applyPlan)) {
		path := filepath.Join(bundleDir, "candidates", filepath.FromSlash(target.Path))
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read candidate draft %s: %w", target.Path, err)
		}
		prompt := renderLLMCandidateReviewPrompt(target.Path, string(sourceDraft), string(content))
		review, err := generator.Generate(context.Background(), prompt)
		if err != nil {
			return fmt.Errorf("review candidate %s: %w", target.Path, err)
		}
		result.Reviews = append(result.Reviews, llmCandidateReview{Path: target.Path, Review: strings.TrimSpace(review)})
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create LLM review output directory: %w", err)
	}
	encoded, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("encode LLM candidate review: %w", err)
	}
	jsonPath := filepath.Join(outDir, "candidate-llm-review.json")
	mdPath := filepath.Join(outDir, "candidate-llm-review.md")
	if err := os.WriteFile(jsonPath, append(encoded, '\n'), 0o644); err != nil {
		return fmt.Errorf("write LLM candidate review json: %w", err)
	}
	if err := os.WriteFile(mdPath, []byte(renderLLMCandidateReview(result)), 0o644); err != nil {
		return fmt.Errorf("write LLM candidate review report: %w", err)
	}
	fmt.Fprintf(os.Stderr, "wrote %s and %s\n", jsonPath, mdPath)
	return nil
}

func renderLLMCandidateReviewPrompt(path string, sourceDraft string, candidateDraft string) string {
	return strings.Join([]string{
		"You are reviewing a candidate Markdown wiki page.",
		"Return concise Markdown only. Do not rewrite the page.",
		"",
		"Assess three things:",
		"- title fidelity: does the body support the page title?",
		"- information density: does each paragraph add source-backed value?",
		"- originality: is the candidate rewritten rather than copied from source.md?",
		"",
		"Return this shape:",
		"- title_fidelity: pass|borderline|fail — reason",
		"- density: pass|borderline|fail — reason",
		"- originality: pass|borderline|fail — reason",
		"- recommendation: one sentence",
		"",
		"Candidate path:",
		path,
		"",
		"source.md:",
		sourceDraft,
		"",
		"Candidate draft:",
		candidateDraft,
	}, "\n")
}

func renderLLMCandidateReview(result llmCandidateReviewResult) string {
	var b strings.Builder
	b.WriteString("# LLM Candidate Review\n\n")
	b.WriteString(fmt.Sprintf("Bundle: `%s`\n\n", result.Bundle))
	if len(result.Reviews) == 0 {
		b.WriteString("(none)\n")
		return b.String()
	}
	for _, review := range result.Reviews {
		b.WriteString(fmt.Sprintf("## `%s`\n\n%s\n\n", review.Path, review.Review))
	}
	return b.String()
}

func runLintBundle(bundleDir string, outDir string) error {
	result, err := evalharness.EvaluateBundleCrosslinks(".", bundleDir)
	if err != nil {
		return fmt.Errorf("lint bundle crosslinks: %w", err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create bundle lint output directory: %w", err)
	}
	encoded, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("encode bundle crosslink lint: %w", err)
	}
	jsonPath := filepath.Join(outDir, "bundle-crosslinks.json")
	mdPath := filepath.Join(outDir, "bundle-crosslinks.md")
	if err := os.WriteFile(jsonPath, append(encoded, '\n'), 0o644); err != nil {
		return fmt.Errorf("write bundle crosslink json: %w", err)
	}
	if err := os.WriteFile(mdPath, []byte(renderBundleCrosslinks(result)), 0o644); err != nil {
		return fmt.Errorf("write bundle crosslink report: %w", err)
	}
	fmt.Fprintf(os.Stderr, "wrote %s and %s\n", jsonPath, mdPath)
	return nil
}

func renderBundleCrosslinks(result evalharness.CrosslinkResult) string {
	var b strings.Builder
	b.WriteString("# Bundle Crosslink Lint\n\n")
	b.WriteString(fmt.Sprintf("Bundle: `%s`\n\n", result.Bundle))
	b.WriteString("## Issues\n\n")
	if len(result.Issues) == 0 {
		b.WriteString("(none)\n\n")
	} else {
		for _, issue := range result.Issues {
			b.WriteString(fmt.Sprintf("- `%s` `%s`: %s\n", issue.Code, issue.Path, issue.Message))
		}
		b.WriteString("\n")
	}
	b.WriteString("## Graph\n\n")
	if len(result.Graph) == 0 {
		b.WriteString("(none)\n")
		return b.String()
	}
	for _, edge := range result.Graph {
		b.WriteString(fmt.Sprintf("- `%s` -> `%s`\n", edge.From, edge.To))
	}
	return b.String()
}

func runLintWiki(outDir string) error {
	result, err := wikilint.LintWiki(".")
	if err != nil {
		return fmt.Errorf("lint wiki: %w", err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create wiki lint output directory: %w", err)
	}
	encoded, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("encode wiki lint result: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "wiki-lint.json"), append(encoded, '\n'), 0o644); err != nil {
		return fmt.Errorf("write wiki lint json: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "wiki-lint.md"), []byte(renderWikiLint(result)), 0o644); err != nil {
		return fmt.Errorf("write wiki lint report: %w", err)
	}

	fmt.Fprintf(os.Stderr, "wrote %s and %s\n", filepath.Join(outDir, "wiki-lint.json"), filepath.Join(outDir, "wiki-lint.md"))
	return nil
}

func runEvalRegression(casesDir string, outDir string) error {
	result, err := evalharness.EvaluateRegressionCases(casesDir)
	if err != nil {
		return fmt.Errorf("evaluate regression cases: %w", err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create regression output directory: %w", err)
	}
	encoded, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("encode regression result: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "regression-summary.json"), append(encoded, '\n'), 0o644); err != nil {
		return fmt.Errorf("write regression summary json: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "regression-summary.md"), []byte(renderRegressionSummary(result)), 0o644); err != nil {
		return fmt.Errorf("write regression summary report: %w", err)
	}
	fmt.Fprintf(os.Stderr, "wrote %s and %s\n", filepath.Join(outDir, "regression-summary.json"), filepath.Join(outDir, "regression-summary.md"))
	return nil
}

func renderRegressionSummary(result evalharness.RegressionResult) string {
	var b strings.Builder
	b.WriteString("# Regression Eval Summary\n\n")
	b.WriteString(fmt.Sprintf("- overall: `%s`\n", result.Overall))
	b.WriteString(fmt.Sprintf("- total: `%d`\n", result.Summary.Total))
	b.WriteString(fmt.Sprintf("- passed: `%d`\n", result.Summary.Passed))
	b.WriteString(fmt.Sprintf("- failed: `%d`\n\n", result.Summary.Failed))
	b.WriteString("## Cases\n\n")
	for _, item := range result.Cases {
		b.WriteString(fmt.Sprintf("### `%s`\n\n", item.Case))
		b.WriteString(fmt.Sprintf("- status: `%s`\n", item.Status))
		b.WriteString(fmt.Sprintf("- bundle: `%s`\n", item.Bundle))
		b.WriteString(fmt.Sprintf("- overall score: `%s`\n", item.ActualScores.Overall))
		for _, failure := range item.Failures {
			b.WriteString("- failure: " + failure + "\n")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func renderWikiLint(result wikilint.Result) string {
	var b strings.Builder
	b.WriteString("# Wiki Lint Report\n\n")
	b.WriteString(fmt.Sprintf("- total issues: `%d`\n", result.Summary.Total))
	b.WriteString(fmt.Sprintf("- pages checked: `%d`\n\n", result.Summary.PageCount))
	b.WriteString("## Issues\n\n")
	if len(result.Issues) == 0 {
		b.WriteString("(none)\n")
		return b.String()
	}
	for _, issue := range result.Issues {
		b.WriteString(fmt.Sprintf("- `%s` `%s` %s: %s\n", issue.Level, issue.Code, issue.Path, issue.Message))
	}
	return b.String()
}

func renderCandidateEvalTrace(result evalharness.CandidateResult) string {
	var b strings.Builder
	b.WriteString("# Candidate Draft Eval Trace\n\n")
	b.WriteString(fmt.Sprintf("Bundle: `%s`\n\n", result.Bundle))
	b.WriteString("## Scores\n\n")
	b.WriteString(fmt.Sprintf("- frontmatter: `%s`\n", result.Scores.Frontmatter))
	b.WriteString(fmt.Sprintf("- sources: `%s`\n", result.Scores.Sources))
	b.WriteString(fmt.Sprintf("- title: `%s`\n", result.Scores.Title))
	b.WriteString(fmt.Sprintf("- title_alignment: `%s`\n", result.Scores.TitleAlignment))
	b.WriteString(fmt.Sprintf("- wikilinks: `%s`\n", result.Scores.Wikilinks))
	b.WriteString(fmt.Sprintf("- length: `%s`\n", result.Scores.Length))
	b.WriteString(fmt.Sprintf("- originality: `%s`\n", result.Scores.Originality))
	b.WriteString(fmt.Sprintf("- overall: `%s`\n\n", result.Scores.Overall))
	b.WriteString("## Candidates\n\n")
	for _, candidate := range result.Candidates {
		b.WriteString(fmt.Sprintf("### `%s`\n\n", candidate.Path))
		b.WriteString(fmt.Sprintf("- overall: `%s`\n", candidate.Scores.Overall))
		for _, entry := range candidate.Trace {
			b.WriteString("- " + entry + "\n")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func runEvalBundle(bundleDir string, outDir string) error {
	result, err := evalharness.EvaluateBundle(bundleDir)
	if err != nil {
		return fmt.Errorf("evaluate bundle: %w", err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create eval output directory: %w", err)
	}
	scores, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("encode scores: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "scores.json"), append(scores, '\n'), 0o644); err != nil {
		return fmt.Errorf("write scores: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "trace.md"), []byte(renderEvalTrace(result)), 0o644); err != nil {
		return fmt.Errorf("write trace: %w", err)
	}

	fmt.Fprintf(os.Stderr, "wrote %s and %s\n", filepath.Join(outDir, "scores.json"), filepath.Join(outDir, "trace.md"))
	return nil
}

func renderEvalTrace(result evalharness.Result) string {
	var b strings.Builder
	b.WriteString("# Ingest Eval Trace\n\n")
	b.WriteString(fmt.Sprintf("Bundle: `%s`\n\n", result.Bundle))
	b.WriteString("## Scores\n\n")
	b.WriteString(fmt.Sprintf("- schema: `%s`\n", result.Scores.Schema))
	b.WriteString(fmt.Sprintf("- wiki_safety: `%s`\n", result.Scores.WikiSafety))
	b.WriteString(fmt.Sprintf("- candidate_paths: `%s`\n", result.Scores.CandidatePaths))
	b.WriteString(fmt.Sprintf("- duplicate_detection: `%s`\n", result.Scores.DuplicateDetection))
	b.WriteString(fmt.Sprintf("- apply_readiness: `%s`\n", result.Scores.ApplyReadiness))
	b.WriteString(fmt.Sprintf("- overall: `%s`\n\n", result.Scores.Overall))
	b.WriteString("## Trace\n\n")
	for _, entry := range result.Trace {
		b.WriteString("- " + entry + "\n")
	}
	return b.String()
}

func runReviewBundle(bundleDir string, out string) error {
	plan, err := review.ReviewBundle(bundleDir)
	if err != nil {
		return fmt.Errorf("review bundle: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		return fmt.Errorf("create review output directory: %w", err)
	}
	if err := os.WriteFile(out, []byte(plan), 0o644); err != nil {
		return fmt.Errorf("write review apply plan: %w", err)
	}

	fmt.Fprintf(os.Stderr, "wrote %s\n", out)
	return nil
}

// inProcessPipeline adapts cmd/nemo's local pipeline functions to the
// web.Pipeline interface. It exists so the -serve flag can hand the same
// in-memory functions to internal/web without re-exporting them.
type inProcessPipeline struct{}

func (inProcessPipeline) RunBundle(source string, bundleDir string, cfg config.Config) error {
	return runBundle(source, bundleDir, cfg)
}

func (inProcessPipeline) RunReviewBundle(bundleDir string, out string) error {
	return runReviewBundle(bundleDir, out)
}

func runBundle(source string, bundleDir string, cfg config.Config) error {
	sourceContent, err := os.ReadFile(source)
	if err != nil {
		return fmt.Errorf("read source: %w", err)
	}
	if cfg.ChunkedBundleCharThreshold > 0 && len(sourceContent) > cfg.ChunkedBundleCharThreshold {
		return runChunkedBundle(source, string(sourceContent), bundleDir, cfg)
	}

	jobs := []struct {
		promptPath string
		outputPath string
	}{
		{
			promptPath: filepath.Join("prompts", "source-page.md"),
			outputPath: filepath.Join(bundleDir, "source.md"),
		},
		{
			promptPath: filepath.Join("prompts", "ingest-plan.md"),
			outputPath: filepath.Join(bundleDir, "ingest-plan.md"),
		},
	}

	for _, job := range jobs {
		if err := runDraft(source, job.promptPath, job.outputPath, cfg); err != nil {
			return err
		}
		if filepath.Base(job.outputPath) == "source.md" {
			if err := normalizeSourceDraftFile(job.outputPath, source); err != nil {
				return err
			}
		}
	}

	return nil
}

func runChunkedBundle(source string, sourceContent string, bundleDir string, cfg config.Config) error {
	maxChunkChars := cfg.MaxChunkChars
	if maxChunkChars <= 0 {
		maxChunkChars = chunking.DefaultMaxChunkChars
	}
	plan := chunking.PlanSource(filepath.ToSlash(source), sourceContent, maxChunkChars)
	chunksDir := filepath.Join(bundleDir, "chunks")
	if err := os.RemoveAll(chunksDir); err != nil {
		return fmt.Errorf("reset chunks directory: %w", err)
	}
	if err := os.MkdirAll(chunksDir, 0o755); err != nil {
		return fmt.Errorf("create chunks directory: %w", err)
	}
	outline := plan.OutlineMarkdown()
	if err := os.WriteFile(filepath.Join(chunksDir, "outline.md"), []byte(outline), 0o644); err != nil {
		return fmt.Errorf("write chunk outline: %w", err)
	}
	indexJSON, err := plan.IndexJSON()
	if err != nil {
		return fmt.Errorf("encode chunk index: %w", err)
	}
	if err := os.WriteFile(filepath.Join(chunksDir, "chunk-index.json"), append(indexJSON, '\n'), 0o644); err != nil {
		return fmt.Errorf("write chunk index: %w", err)
	}

	chunkTemplate, err := os.ReadFile(filepath.Join("prompts", "chunk-notes.md"))
	if err != nil {
		return fmt.Errorf("read chunk notes prompt: %w", err)
	}
	notePaths := make([]string, 0, len(plan.Chunks))
	for _, chunk := range plan.Chunks {
		rendered, err := prompt.Render(string(chunkTemplate), prompt.Variables{
			RawSourcePath: filepath.ToSlash(source),
			ChunkContent:  chunking.FormatChunkForPrompt(filepath.ToSlash(source), chunk, len(plan.Chunks)),
		})
		if err != nil {
			return fmt.Errorf("render chunk notes prompt: %w", err)
		}
		out := filepath.Join(chunksDir, fmt.Sprintf("chunk-%02d.md", chunk.Index))
		if err := runRenderedDraft(rendered, out, cfg); err != nil {
			return fmt.Errorf("generate chunk %02d notes: %w", chunk.Index, err)
		}
		notePaths = append(notePaths, out)
	}

	notes, err := readChunkNotes(notePaths)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(chunksDir, "combined-notes.md"), []byte(notes), 0o644); err != nil {
		return fmt.Errorf("write combined chunk notes: %w", err)
	}

	groupNotes, err := runChunkGroupNotes(chunksDir, notePaths, source, outline, string(indexJSON), cfg)
	if err != nil {
		return err
	}
	if groupNotes != "" {
		if err := os.WriteFile(filepath.Join(chunksDir, "combined-group-notes.md"), []byte(groupNotes), 0o644); err != nil {
			return fmt.Errorf("write combined group notes: %w", err)
		}
	}

	// When group notes are present they are the whole-document summary; the
	// final source/ingest synthesis should rely on them plus the outline and
	// index, and skip the raw combined chunk notes. Passing both layers into a
	// single prompt easily exceeds the local model's context window for very
	// long sources (observed 36k-43k tokens for 16-24 chunk standards).
	finalChunkNotes := notes
	if groupNotes != "" {
		finalChunkNotes = ""
	}
	if err := runChunkSynthesis(filepath.Join("prompts", "chunk-source-page.md"), filepath.Join(bundleDir, "source.md"), source, outline, string(indexJSON), finalChunkNotes, groupNotes, cfg); err != nil {
		return err
	}
	if err := normalizeSourceDraftFile(filepath.Join(bundleDir, "source.md"), source); err != nil {
		return err
	}
	if err := runChunkSynthesis(filepath.Join("prompts", "chunk-ingest-plan.md"), filepath.Join(bundleDir, "ingest-plan.md"), source, outline, string(indexJSON), finalChunkNotes, groupNotes, cfg); err != nil {
		return err
	}
	return nil
}

func runChunkGroupNotes(chunksDir string, notePaths []string, source string, outline string, index string, cfg config.Config) (string, error) {
	if len(notePaths) <= chunkNotesPerGroup {
		return "", nil
	}
	// The group layer is for whole-document understanding, not truncation.
	// When a long source produces many short-but-coherent chunk notes, final
	// synthesis should not have to infer cross-document themes from a flat
	// list. Group notes compress adjacent chunks into regional summaries while
	// keeping the original chunk notes available for source-backed detail.
	templateContent, err := os.ReadFile(filepath.Join("prompts", "chunk-group-notes.md"))
	if err != nil {
		return "", fmt.Errorf("read chunk group notes prompt: %w", err)
	}
	groupPaths := make([]string, 0, (len(notePaths)+chunkNotesPerGroup-1)/chunkNotesPerGroup)
	for start := 0; start < len(notePaths); start += chunkNotesPerGroup {
		end := start + chunkNotesPerGroup
		if end > len(notePaths) {
			end = len(notePaths)
		}
		notes, err := readChunkNotes(notePaths[start:end])
		if err != nil {
			return "", err
		}
		rendered, err := prompt.Render(string(templateContent), prompt.Variables{
			RawSourcePath: filepath.ToSlash(source),
			ChunkOutline:  outline,
			ChunkIndex:    index,
			ChunkNotes:    notes,
		})
		if err != nil {
			return "", fmt.Errorf("render chunk group notes prompt: %w", err)
		}
		out := filepath.Join(chunksDir, fmt.Sprintf("group-%02d.md", len(groupPaths)+1))
		if err := runRenderedDraft(rendered, out, cfg); err != nil {
			return "", fmt.Errorf("generate chunk group %02d notes: %w", len(groupPaths)+1, err)
		}
		groupPaths = append(groupPaths, out)
	}
	return readChunkNotes(groupPaths)
}

func readChunkNotes(paths []string) (string, error) {
	var b strings.Builder
	for _, path := range paths {
		content, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("read chunk note: %w", err)
		}
		b.WriteString("## ")
		b.WriteString(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
		b.WriteString("\n\n")
		b.Write(content)
		b.WriteString("\n")
	}
	return b.String(), nil
}

func runChunkSynthesis(promptPath string, out string, source string, outline string, index string, notes string, groupNotes string, cfg config.Config) error {
	templateContent, err := os.ReadFile(promptPath)
	if err != nil {
		return fmt.Errorf("read chunk synthesis prompt: %w", err)
	}
	rendered, err := prompt.Render(string(templateContent), prompt.Variables{
		RawSourcePath:   filepath.ToSlash(source),
		ChunkOutline:    outline,
		ChunkIndex:      index,
		ChunkNotes:      notes,
		ChunkGroupNotes: groupNotes,
	})
	if err != nil {
		return fmt.Errorf("render chunk synthesis prompt: %w", err)
	}
	return runRenderedDraft(rendered, out, cfg)
}

func normalizeSourceDraftFile(path string, source string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read source draft for normalization: %w", err)
	}
	normalized := normalizeSourceDraft(string(content), filepath.ToSlash(source))
	if err := os.WriteFile(path, []byte(normalized), 0o644); err != nil {
		return fmt.Errorf("write normalized source draft: %w", err)
	}
	return nil
}

func normalizeSourceDraft(content string, source string) string {
	frontmatter, body := splitMarkdownFrontmatter(content)
	title := frontmatterField(frontmatter, "title")
	if title == "" {
		title = firstMarkdownHeading(body)
	}
	if title == "" {
		title = titleFromSlug(strings.TrimSuffix(filepath.Base(source), filepath.Ext(source)))
	}
	body = strings.TrimSpace(body)
	body = normalizeSourceSectionHeadings(body)

	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("title: " + title + "\n")
	b.WriteString("kind: source\n")
	b.WriteString("sources:\n")
	b.WriteString("  - " + source + "\n")
	b.WriteString("confidence: medium\n")
	b.WriteString("---\n\n")
	b.WriteString(body)
	b.WriteString("\n")
	return b.String()
}

func normalizeSourceSectionHeadings(body string) string {
	for _, section := range []string{"What It Is", "Summary", "Key Claims", "Suggested Links"} {
		re := regexp.MustCompile(`(?m)^#{1,6}[ \t]+` + regexp.QuoteMeta(section) + `[ \t#]*$`)
		body = re.ReplaceAllString(body, "## "+section)
	}
	return body
}

func splitMarkdownFrontmatter(content string) (string, string) {
	match := sourceDraftFrontmatterRE.FindStringSubmatch(content)
	if len(match) != 2 {
		return "", strings.TrimSpace(content)
	}
	return match[1], strings.TrimSpace(sourceDraftFrontmatterRE.ReplaceAllString(content, ""))
}

func frontmatterField(frontmatter string, field string) string {
	re := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(field) + `:\s*(.+?)\s*$`)
	match := re.FindStringSubmatch(frontmatter)
	if len(match) != 2 {
		return ""
	}
	return strings.Trim(strings.TrimSpace(match[1]), `"'`)
}

func firstMarkdownHeading(content string) string {
	match := regexp.MustCompile(`(?m)^#\s+(.+?)\s*$`).FindStringSubmatch(content)
	if len(match) != 2 {
		return ""
	}
	return strings.TrimSpace(match[1])
}

func persistRawWebSource(source string) (string, error) {
	cleanSource := filepath.Clean(source)
	if strings.HasPrefix(filepath.ToSlash(cleanSource), "raw/") {
		return filepath.ToSlash(cleanSource), nil
	}
	content, err := os.ReadFile(source)
	if err != nil {
		return "", fmt.Errorf("read source for raw persistence: %w", err)
	}
	slug := slugFromFilename(filepath.Base(source))
	if slug == "" {
		return "", fmt.Errorf("persist raw web source: cannot derive slug from %s", source)
	}
	target := filepath.Join("raw", "web", slug+".md")
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return "", fmt.Errorf("create raw web directory: %w", err)
	}
	file, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return "", fmt.Errorf("persist raw web source: %s already exists", filepath.ToSlash(target))
		}
		return "", fmt.Errorf("create raw web source: %w", err)
	}
	defer file.Close()
	if _, err := file.Write(content); err != nil {
		return "", fmt.Errorf("write raw web source: %w", err)
	}
	return filepath.ToSlash(target), nil
}

func slugFromFilename(name string) string {
	base := strings.TrimSuffix(name, filepath.Ext(name))
	base = strings.ToLower(base)
	var b strings.Builder
	lastHyphen := false
	for _, r := range base {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
			lastHyphen = false
		default:
			if !lastHyphen && b.Len() > 0 {
				b.WriteByte('-')
				lastHyphen = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}

func runDraft(source string, promptPath string, out string, cfg config.Config) error {
	sourceContent, err := os.ReadFile(source)
	if err != nil {
		return fmt.Errorf("read source: %w", err)
	}

	templateContent, err := os.ReadFile(promptPath)
	if err != nil {
		return fmt.Errorf("read prompt: %w", err)
	}

	rendered, err := prompt.Render(string(templateContent), prompt.Variables{
		RawSourcePath:    source,
		RawSourceContent: string(sourceContent),
	})
	if err != nil {
		return fmt.Errorf("render prompt: %w", err)
	}

	return runRenderedDraft(rendered, out, cfg)
}

func runRenderedDraft(rendered string, out string, cfg config.Config) error {
	generator := generatorFromConfig(cfg)

	rawOutput, err := generator.Generate(context.Background(), rendered)
	if err != nil {
		return fmt.Errorf("generate draft: %w", err)
	}

	paths := draft.PathsFor(out)
	if err := os.MkdirAll(filepath.Dir(paths.Cleaned), 0o755); err != nil {
		return fmt.Errorf("create draft directory: %w", err)
	}
	if err := os.WriteFile(paths.Raw, []byte(rawOutput), 0o644); err != nil {
		return fmt.Errorf("write raw draft: %w", err)
	}

	cleaned, err := draft.Clean(rawOutput)
	if err != nil {
		if cfg.Profile == "fallback" {
			return fmt.Errorf("clean draft: %w", err)
		}
		cleaned, err = runFallbackDraft(rendered, paths.Cleaned, cfg)
		if err != nil {
			return err
		}
	}
	if err := os.WriteFile(paths.Cleaned, []byte(cleaned), 0o644); err != nil {
		return fmt.Errorf("write cleaned draft: %w", err)
	}

	fmt.Fprintf(os.Stderr, "wrote %s and %s\n", paths.Raw, paths.Cleaned)
	return nil
}

func runFallbackDraft(renderedPrompt string, cleanedPath string, cfg config.Config) (string, error) {
	fallback, err := config.ForProfile("fallback")
	if err != nil {
		return "", fmt.Errorf("load fallback profile: %w", err)
	}
	copyBackendConfig(&fallback, cfg)

	generator := generatorFromConfig(fallback)

	rawOutput, err := generator.Generate(context.Background(), renderedPrompt)
	if err != nil {
		return "", fmt.Errorf("generate fallback draft: %w", err)
	}
	if err := os.WriteFile(fallbackRawPath(cleanedPath), []byte(rawOutput), 0o644); err != nil {
		return "", fmt.Errorf("write fallback raw draft: %w", err)
	}

	cleaned, err := draft.Clean(rawOutput)
	if err != nil {
		return "", fmt.Errorf("clean fallback draft: %w", err)
	}

	return cleaned, nil
}

func generatorFromConfig(cfg config.Config) llama.Generator {
	if cfg.Provider == "deepseek" {
		return deepseek.Client{
			BaseURL:         cfg.DeepSeek.BaseURL,
			APIKey:          cfg.DeepSeek.APIKey,
			Model:           cfg.DeepSeek.Model,
			MaxTokens:       cfg.DeepSeek.MaxTokens,
			Temperature:     cfg.DeepSeek.Temperature,
			TopP:            cfg.DeepSeek.TopP,
			Thinking:        cfg.DeepSeek.Thinking,
			ReasoningEffort: cfg.DeepSeek.ReasoningEffort,
			ResponseFormat:  cfg.DeepSeek.ResponseFormat,
			UserID:          cfg.DeepSeek.UserID,
			SystemPrompt:    cfg.DeepSeek.SystemPrompt,
			RetryMax:        cfg.DeepSeek.RetryMax,
			RetryBaseDelay:  time.Duration(cfg.DeepSeek.RetryBaseDelayMS) * time.Millisecond,
		}
	}
	return llamaCLIFromConfig(cfg)
}

func copyBackendConfig(dst *config.Config, src config.Config) {
	dst.Provider = src.Provider
	dst.LlamaCLI = src.LlamaCLI
	dst.LlamaModel = src.LlamaModel
	dst.GPULayers = src.GPULayers
	dst.DeepSeek.BaseURL = src.DeepSeek.BaseURL
	dst.DeepSeek.APIKey = src.DeepSeek.APIKey
	dst.DeepSeek.ResponseFormat = src.DeepSeek.ResponseFormat
	dst.DeepSeek.UserID = src.DeepSeek.UserID
	dst.DeepSeek.SystemPrompt = src.DeepSeek.SystemPrompt
	dst.DeepSeek.RetryMax = src.DeepSeek.RetryMax
	dst.DeepSeek.RetryBaseDelayMS = src.DeepSeek.RetryBaseDelayMS
}

func llamaCLIFromConfig(cfg config.Config) llama.CLI {
	return llama.CLI{
		Binary:                 cfg.LlamaCLI,
		Model:                  cfg.LlamaModel,
		GPULayers:              cfg.GPULayers,
		MaxTokens:              cfg.MaxTokens,
		CtxSize:                cfg.CtxSize,
		Temp:                   cfg.Temp,
		TopP:                   cfg.TopP,
		TopK:                   cfg.TopK,
		MinP:                   cfg.MinP,
		PresencePenalty:        cfg.PresencePenalty,
		RepeatPenalty:          cfg.RepeatPenalty,
		Reasoning:              cfg.Reasoning,
		ReasoningBudget:        cfg.ReasoningBudget,
		ReasoningBudgetMessage: cfg.ReasoningBudgetMessage,
		ChatTemplateKwargs:     cfg.ChatTemplateKwargs,
		Jinja:                  cfg.Jinja,
		NoContextShift:         cfg.NoContextShift,
	}
}

func fallbackRawPath(cleanedPath string) string {
	ext := filepath.Ext(cleanedPath)
	if ext == "" {
		return cleanedPath + ".fallback.raw.txt"
	}

	return cleanedPath[:len(cleanedPath)-len(ext)] + ".fallback.raw.txt"
}
