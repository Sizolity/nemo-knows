package query

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	wikischema "github.com/huic/nemo-knows/internal/wiki"
)

var (
	ErrQuestionRequired = errors.New("query requires a question")
	ErrApprovalRequired = errors.New("filing a query answer requires explicit approval")

	frontmatterRE = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*`)
	titleRE       = regexp.MustCompile(`(?m)^title:\s*(.+?)\s*$`)
	kindRE        = regexp.MustCompile(`(?m)^kind:\s*(.+?)\s*$`)
	headingRE     = regexp.MustCompile(`(?m)^#\s+(.+?)\s*$`)
	tokenRE       = regexp.MustCompile(`[A-Za-z0-9]+`)
	// inlineLinkRE matches inline Markdown links and images so the topic draft
	// path can flatten them to their visible label: [label](href) and
	// ![alt](src). Group 1 captures the label/alt text.
	inlineLinkRE = regexp.MustCompile(`!?\[([^\]]*)\]\([^)]*\)`)
)

type Options struct {
	Root     string
	Question string
	File     bool
	Approve  bool
	Out      string
	Now      time.Time
}

type Result struct {
	Answer   string
	Draft    string
	DraftOut string
	Written  []string
	Relevant []Page
}

type Page struct {
	Path    string
	Title   string
	Kind    string
	Content string
	Score   int
}

// Execute answers a question from the maintained wiki and optionally files the
// answer back as an approved topic page.
func Execute(opts Options) (Result, error) {
	if strings.TrimSpace(opts.Question) == "" {
		return Result{}, ErrQuestionRequired
	}
	root := opts.Root
	if root == "" {
		root = "."
	}
	now := opts.Now
	if now.IsZero() {
		now = time.Now()
	}

	pages, err := relevantPages(root, opts.Question)
	if err != nil {
		return Result{}, err
	}
	result := Result{Relevant: pages}
	result.Answer = renderAnswer(opts.Question, pages)
	if !opts.File {
		return result, nil
	}

	result.Draft = renderTopicDraft(opts.Question, pages, now)
	if !opts.Approve {
		if strings.HasPrefix(filepath.ToSlash(opts.Out), "wiki/") {
			return result, ErrApprovalRequired
		}
		draftPath := opts.Out
		if draftPath == "" {
			draftPath = filepath.Join("tmp", "query-drafts", slugForQuestion(opts.Question)+".md")
		}
		writePath := draftPath
		if !filepath.IsAbs(writePath) {
			writePath = filepath.Join(root, filepath.FromSlash(writePath))
		}
		if err := writeReviewDraft(writePath, result.Draft); err != nil {
			return Result{}, err
		}
		result.DraftOut = displayPath(root, writePath)
		result.Written = append(result.Written, result.DraftOut)
		return result, nil
	}

	target := opts.Out
	if target == "" {
		target = filepath.ToSlash(filepath.Join("wiki", "topics", slugForQuestion(opts.Question)+".md"))
	}
	cleanTarget, ok := cleanTopicPath(target)
	if !ok {
		return Result{}, fmt.Errorf("refuse to file query outside wiki/topics: %s", target)
	}
	if err := writeApprovedTopic(root, cleanTarget, result.Draft); err != nil {
		return Result{}, err
	}
	result.DraftOut = cleanTarget
	result.Written = append(result.Written, cleanTarget)
	if written, err := updateIndex(root, cleanTarget, topicTitle(opts.Question)); err != nil {
		return Result{}, err
	} else if written {
		result.Written = append(result.Written, "wiki/index.md")
	}
	if err := appendQueryLog(root, opts.Question, cleanTarget, pages, now); err != nil {
		return Result{}, err
	}
	result.Written = append(result.Written, "wiki/log.md")
	return result, nil
}

func relevantPages(root string, question string) ([]Page, error) {
	if _, err := os.ReadFile(filepath.Join(root, "wiki", "index.md")); err != nil {
		return nil, fmt.Errorf("read wiki index: %w", err)
	}
	pages, err := scanWikiPages(root)
	if err != nil {
		return nil, err
	}
	queryTokens := tokenSet(question)
	for i := range pages {
		pages[i].Score = scorePage(pages[i], queryTokens)
	}
	sort.Slice(pages, func(i int, j int) bool {
		if pages[i].Score == pages[j].Score {
			return pages[i].Path < pages[j].Path
		}
		return pages[i].Score > pages[j].Score
	})
	filtered := pages[:0]
	for _, page := range pages {
		if page.Score > 0 {
			filtered = append(filtered, page)
		}
	}
	if len(filtered) > 5 {
		filtered = filtered[:5]
	}
	return filtered, nil
}

func scanWikiPages(root string) ([]Page, error) {
	var pages []Page
	for _, dir := range []string{"sources", "entities", "concepts", "topics"} {
		base := filepath.Join(root, "wiki", dir)
		if _, err := os.Stat(base); errors.Is(err, os.ErrNotExist) {
			continue
		}
		err := filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || filepath.Ext(path) != ".md" {
				return err
			}
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			pages = append(pages, pageFromContent(filepath.ToSlash(rel), string(content)))
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("scan wiki %s: %w", dir, err)
		}
	}
	return pages, nil
}

func pageFromContent(path string, content string) Page {
	frontmatter := ""
	if match := frontmatterRE.FindStringSubmatch(content); len(match) == 2 {
		frontmatter = match[1]
	}
	title := frontmatterValue(frontmatter, titleRE)
	if title == "" {
		title = firstHeading(content)
	}
	if title == "" {
		title = titleFromSlug(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
	}
	kind := frontmatterValue(frontmatter, kindRE)
	if kind == "" {
		kind = strings.TrimSuffix(strings.TrimPrefix(filepath.Dir(path), "wiki/"), "s")
	}
	return Page{Path: path, Title: title, Kind: kind, Content: content}
}

func scorePage(page Page, queryTokens map[string]bool) int {
	if len(queryTokens) == 0 {
		return 0
	}
	text := strings.ToLower(page.Path + " " + page.Title + " " + page.Content)
	score := 0
	for token := range queryTokens {
		if strings.Contains(text, token) {
			score++
		}
	}
	return score
}

func renderAnswer(question string, pages []Page) string {
	var b strings.Builder
	b.WriteString("# Query Answer\n\n")
	b.WriteString("Question: ")
	b.WriteString(strings.TrimSpace(question))
	b.WriteString("\n\n")
	if len(pages) == 0 {
		b.WriteString("The current wiki does not contain enough source-backed material to answer this question. Ingest more relevant sources, then query again.\n")
		return b.String()
	}
	b.WriteString("Relevant wiki pages:\n")
	for _, page := range pages {
		b.WriteString("- `")
		b.WriteString(page.Path)
		b.WriteString("` — ")
		b.WriteString(page.Title)
		b.WriteString(". ")
		b.WriteString(snippet(page.Content))
		b.WriteString(" (source: ")
		b.WriteString(page.Path)
		b.WriteString(")\n")
	}
	return b.String()
}

func renderTopicDraft(question string, pages []Page, now time.Time) string {
	fromRepoPath := topicDraftRepoPath(question)
	// The question is untrusted free text that flows into the filed topic. The
	// title (frontmatter + H1) is sanitized here, and the assembled body is run
	// through the same shared cleaner below, so a question carrying
	// [[wikilink]], {{placeholder}}, or a dangerous/escaping link target cannot
	// land verbatim in wiki/topics/. Code spans are preserved.
	title := wikischema.SanitizeBodyLinks(fromRepoPath, topicTitle(question))
	date := now.Format("2006-01-02")
	var fm strings.Builder
	fm.WriteString("---\n")
	fm.WriteString("title: ")
	fm.WriteString(title)
	fm.WriteString("\nkind: topic\ncreated: ")
	fm.WriteString(date)
	fm.WriteString("\nupdated: ")
	fm.WriteString(date)
	fm.WriteString("\nsources:\n")
	if len(pages) == 0 {
		fm.WriteString("  - wiki/index.md\n")
	} else {
		for _, page := range pages {
			fm.WriteString("  - ")
			fm.WriteString(page.Path)
			fm.WriteByte('\n')
		}
	}
	fm.WriteString("confidence: medium\n---\n\n")

	var b strings.Builder
	b.WriteString("# ")
	b.WriteString(title)
	b.WriteString("\n\n")
	b.WriteString("Question: ")
	b.WriteString(strings.TrimSpace(question))
	b.WriteString("\n\n")
	if len(pages) == 0 {
		b.WriteString("The current wiki does not contain enough source-backed material to answer this question. This draft should not be filed as a durable answer until relevant sources are ingested.\n")
		return fm.String() + wikischema.SanitizeBodyLinks(fromRepoPath, b.String())
	}
	b.WriteString("The current answer is grounded in the maintained wiki pages listed below. It should be reviewed before filing because this deterministic query path only performs keyword matching and excerpt extraction.\n\n")
	b.WriteString("## Answer Notes\n\n")
	for _, page := range pages {
		b.WriteString("- ")
		b.WriteString(page.Title)
		b.WriteString(": ")
		b.WriteString(topicExcerpt(page.Content))
		b.WriteString(" (source: ")
		b.WriteString(page.Path)
		b.WriteString(")\n")
	}
	b.WriteString("\n## References\n\n")
	for _, page := range pages {
		b.WriteString("- ")
		b.WriteString(page.Path)
		b.WriteByte('\n')
	}
	return fm.String() + wikischema.SanitizeBodyLinks(fromRepoPath, b.String())
}

// topicDraftRepoPath is the wiki path a filed query topic will occupy. Only the
// wiki/topics/ directory depth matters: it anchors the link sanitizer's
// repository-escape (../) check at the right level; the exact slug is
// irrelevant to that resolution.
func topicDraftRepoPath(question string) string {
	return "wiki/topics/" + slugForQuestion(question) + ".md"
}

func writeReviewDraft(path string, draft string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create query draft directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(draft), 0o644); err != nil {
		return fmt.Errorf("write query draft: %w", err)
	}
	return nil
}

func displayPath(root string, path string) string {
	rel, err := filepath.Rel(root, path)
	if err == nil && !strings.HasPrefix(rel, "..") {
		return filepath.ToSlash(rel)
	}
	return filepath.ToSlash(path)
}

func writeApprovedTopic(root string, repoPath string, draft string) error {
	path := filepath.Join(root, filepath.FromSlash(repoPath))
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("query topic already exists: %s", repoPath)
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("stat query topic: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create query topic directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(draft), 0o644); err != nil {
		return fmt.Errorf("write query topic: %w", err)
	}
	return nil
}

func cleanTopicPath(repoPath string) (string, bool) {
	if filepath.IsAbs(repoPath) {
		return "", false
	}
	clean := filepath.ToSlash(filepath.Clean(filepath.FromSlash(repoPath)))
	if clean != filepath.ToSlash(repoPath) || filepath.Ext(clean) != ".md" {
		return "", false
	}
	if !strings.HasPrefix(clean, "wiki/topics/") {
		return "", false
	}
	return clean, true
}

func updateIndex(root string, target string, title string) (bool, error) {
	path := filepath.Join(root, "wiki", "index.md")
	content, err := os.ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("read wiki index: %w", err)
	}
	// Catalog the index by section + relative path, not by slug, so a filed
	// topic is never silently dropped just because another category already has
	// a page with the same slug.
	wantRel := wikischema.IndexRelativePath(target)
	for _, line := range strings.Split(string(content), "\n") {
		if existing, ok := wikischema.IndexEntryTarget(line); ok && wikischema.IndexRelativePath(existing) == wantRel {
			return false, nil
		}
	}
	entry := wikischema.FormatIndexEntry(target, title) + "\n"
	updated := wikischema.AppendIndexEntry(string(content), "## Topics", entry)
	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		return false, fmt.Errorf("write wiki index: %w", err)
	}
	return true, nil
}

func appendQueryLog(root string, question string, target string, pages []Page, now time.Time) error {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("\n## [%s] query-filed | %s\n", now.Format("2006-01-02"), topicTitle(question)))
	b.WriteString("Question: ")
	b.WriteString(strings.TrimSpace(question))
	b.WriteString("\nTouched:\n")
	b.WriteString("- ")
	b.WriteString(strings.TrimPrefix(target, "wiki/"))
	b.WriteString(" (created)\n")
	b.WriteString("- index.md (updated)\n")
	if len(pages) > 0 {
		b.WriteString("References:\n")
		for _, page := range pages {
			b.WriteString("- ")
			b.WriteString(page.Path)
			b.WriteByte('\n')
		}
	}
	b.WriteString("Open: review whether this filed answer should be expanded after future ingests.\n")

	path := filepath.Join(root, "wiki", "log.md")
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open wiki log: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(b.String()); err != nil {
		return fmt.Errorf("append wiki log: %w", err)
	}
	return nil
}

func tokenSet(text string) map[string]bool {
	tokens := map[string]bool{}
	for _, token := range tokenRE.FindAllString(strings.ToLower(text), -1) {
		if len(token) < 3 || stopWord(token) {
			continue
		}
		tokens[token] = true
	}
	return tokens
}

func stopWord(token string) bool {
	switch token {
	case "the", "and", "for", "with", "from", "this", "that", "what", "how", "why", "are", "can", "does", "into":
		return true
	default:
		return false
	}
}

func snippet(content string) string {
	return truncateExcerpt(firstBodyLine(content))
}

// firstBodyLine returns the first non-empty, non-heading body line of a page,
// or "" when the page has no usable prose line.
func firstBodyLine(content string) string {
	_, body := splitFrontmatter(content)
	for _, line := range strings.Split(body, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "---") {
			continue
		}
		return line
	}
	return ""
}

// truncateExcerpt caps an excerpt at 220 characters and supplies the fallback
// text used when a page has no usable body line.
func truncateExcerpt(line string) string {
	if line == "" {
		return "No short summary is available."
	}
	if len(line) > 220 {
		return strings.TrimSpace(line[:220]) + "..."
	}
	return line
}

// topicExcerpt is the excerpt used for FILED topic drafts. It flattens inline
// Markdown links/images to their visible label BEFORE truncation, so relative
// links (written relative to the source page's directory) are never copied
// verbatim into wiki/topics/, where they would dangle and trip the
// missing-link-target lint. The authoritative source path is preserved
// separately in the draft's "(source: …)" note and "## References" list.
func topicExcerpt(content string) string {
	return truncateExcerpt(stripInlineMarkdownLinks(firstBodyLine(content)))
}

// stripInlineMarkdownLinks rewrites [label](href) -> label and ![alt](src) ->
// alt, dropping the link/image target entirely.
func stripInlineMarkdownLinks(s string) string {
	return inlineLinkRE.ReplaceAllString(s, "$1")
}

func splitFrontmatter(content string) (string, string) {
	match := frontmatterRE.FindStringSubmatch(content)
	if len(match) != 2 {
		return "", strings.TrimSpace(content)
	}
	return match[1], strings.TrimSpace(frontmatterRE.ReplaceAllString(content, ""))
}

func frontmatterValue(frontmatter string, re *regexp.Regexp) string {
	match := re.FindStringSubmatch(frontmatter)
	if len(match) != 2 {
		return ""
	}
	return strings.Trim(strings.TrimSpace(match[1]), `"'`)
}

func firstHeading(content string) string {
	match := headingRE.FindStringSubmatch(content)
	if len(match) != 2 {
		return ""
	}
	return strings.TrimSpace(match[1])
}

func topicTitle(question string) string {
	clean := strings.TrimSpace(question)
	clean = strings.TrimRight(clean, "?.! ")
	if clean == "" {
		return "Filed Query Answer"
	}
	words := strings.Fields(clean)
	if len(words) > 8 {
		words = words[:8]
	}
	return strings.Join(words, " ")
}

func slugForQuestion(question string) string {
	tokens := []string{}
	for _, token := range tokenRE.FindAllString(strings.ToLower(question), -1) {
		if stopWord(token) {
			continue
		}
		tokens = append(tokens, token)
		if len(tokens) == 6 {
			break
		}
	}
	if len(tokens) == 0 {
		return "filed-query-answer"
	}
	return strings.Join(tokens, "-")
}

func titleFromSlug(slug string) string {
	parts := strings.Split(slug, "-")
	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}
