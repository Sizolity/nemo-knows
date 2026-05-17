package wikilint

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	frontmatterBlockRE = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*`)
	wikilinkRE         = regexp.MustCompile(`\[\[([^\]|#]+)(?:[|#][^\]]*)?\]\]`)
	logHeadingRE       = regexp.MustCompile(`^## \[[0-9]{4}-[0-9]{2}-[0-9]{2}\] ([^ |]+) \| .+`)
	indexEntryRE       = regexp.MustCompile(`- \[\[([^\]|#]+)(?:[|#][^\]]*)?\]\]`)
)

type Result struct {
	Summary Summary `json:"summary"`
	Issues  []Issue `json:"issues"`
}

type Summary struct {
	Total     int            `json:"total"`
	ByCode    map[string]int `json:"by_code"`
	ByLevel   map[string]int `json:"by_level"`
	PageCount int            `json:"page_count"`
}

type Issue struct {
	Code    string `json:"code"`
	Level   string `json:"level"`
	Path    string `json:"path"`
	Message string `json:"message"`
}

type page struct {
	Path        string
	Slug        string
	Content     string
	Frontmatter string
	Links       []string
}

func LintWiki(root string) (Result, error) {
	pages, err := readWikiPages(root)
	if err != nil {
		return Result{}, err
	}

	result := Result{}
	slugToPath := map[string]string{}
	inbound := map[string]int{}
	for _, page := range pages {
		slugToPath[page.Slug] = page.Path
	}

	for _, page := range pages {
		lintFrontmatter(page, &result)
		for _, link := range page.Links {
			if _, ok := slugToPath[link]; !ok {
				addIssue(&result, "missing-wikilink-target", "warn", page.Path, "wikilink target does not exist: "+link)
				continue
			}
			inbound[link]++
		}
	}
	lintIndex(root, &result)
	lintLog(root, &result)
	for _, page := range pages {
		if page.Path == "wiki/index.md" || page.Path == "wiki/log.md" {
			continue
		}
		if inbound[page.Slug] == 0 {
			addIssue(&result, "orphan-page", "info", page.Path, "page has no inbound wikilinks")
		}
	}

	result.Summary.PageCount = len(pages)
	result.Summary.Total = len(result.Issues)
	result.Summary.ByCode = map[string]int{}
	result.Summary.ByLevel = map[string]int{}
	for _, issue := range result.Issues {
		result.Summary.ByCode[issue.Code]++
		result.Summary.ByLevel[issue.Level]++
	}
	sort.Slice(result.Issues, func(i, j int) bool {
		if result.Issues[i].Path == result.Issues[j].Path {
			return result.Issues[i].Code < result.Issues[j].Code
		}
		return result.Issues[i].Path < result.Issues[j].Path
	})

	return result, nil
}

func readWikiPages(root string) ([]page, error) {
	wikiRoot := filepath.Join(root, "wiki")
	pages := []page{}
	err := filepath.WalkDir(wikiRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		repoPath := filepath.ToSlash(rel)
		fm, _ := splitFrontmatter(string(content))
		pages = append(pages, page{
			Path:        repoPath,
			Slug:        strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
			Content:     string(content),
			Frontmatter: fm,
			Links:       wikilinks(string(content)),
		})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk wiki: %w", err)
	}
	return pages, nil
}

func lintFrontmatter(page page, result *Result) {
	if page.Frontmatter == "" {
		addIssue(result, "missing-frontmatter", "error", page.Path, "page is missing YAML frontmatter")
		return
	}
	kind := fmValue(page.Frontmatter, "kind")
	if kind == "" {
		addIssue(result, "missing-kind", "error", page.Path, "frontmatter is missing kind")
	}
	if !validKindForPath(kind, page.Path) {
		addIssue(result, "invalid-kind", "error", page.Path, "frontmatter kind does not match path")
	}
	if requiresSources(page.Path) && !strings.Contains(page.Frontmatter, "sources:") {
		addIssue(result, "missing-sources", "error", page.Path, "frontmatter is missing sources")
	}
	if requiresSources(page.Path) {
		confidence := fmValue(page.Frontmatter, "confidence")
		if confidence != "high" && confidence != "medium" && confidence != "low" {
			addIssue(result, "invalid-confidence", "error", page.Path, "confidence must be high, medium, or low")
		}
	}
}

func lintIndex(root string, result *Result) {
	content, err := os.ReadFile(filepath.Join(root, "wiki", "index.md"))
	if err != nil {
		addIssue(result, "missing-index", "error", "wiki/index.md", "wiki index could not be read")
		return
	}
	seen := map[string]bool{}
	for _, match := range indexEntryRE.FindAllStringSubmatch(string(content), -1) {
		slug := match[1]
		if seen[slug] {
			addIssue(result, "duplicate-index-entry", "warn", "wiki/index.md", "duplicate index entry: "+slug)
		}
		seen[slug] = true
	}
}

func lintLog(root string, result *Result) {
	content, err := os.ReadFile(filepath.Join(root, "wiki", "log.md"))
	if err != nil {
		addIssue(result, "missing-log", "error", "wiki/log.md", "wiki log could not be read")
		return
	}
	validActions := map[string]bool{"ingest": true, "query-filed": true, "lint": true, "schema-change": true, "note": true}
	body := stripMarkdownCode(string(content))
	for _, line := range strings.Split(body, "\n") {
		if !strings.HasPrefix(line, "## [") {
			continue
		}
		match := logHeadingRE.FindStringSubmatch(line)
		if len(match) != 2 || !validActions[match[1]] {
			addIssue(result, "invalid-log-action", "error", "wiki/log.md", "invalid log heading action: "+line)
		}
	}
}

func splitFrontmatter(content string) (string, string) {
	match := frontmatterBlockRE.FindStringSubmatch(content)
	if len(match) != 2 {
		return "", content
	}
	return match[1], frontmatterBlockRE.ReplaceAllString(content, "")
}

func wikilinks(content string) []string {
	links := []string{}
	body := stripMarkdownCode(content)
	for _, match := range wikilinkRE.FindAllStringSubmatch(body, -1) {
		links = append(links, match[1])
	}
	return links
}

func stripMarkdownCode(content string) string {
	var b strings.Builder
	inFence := false
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inFence = !inFence
			continue
		}
		if inFence {
			continue
		}
		b.WriteString(stripInlineCode(line))
		b.WriteString("\n")
	}
	return b.String()
}

func stripInlineCode(line string) string {
	var b strings.Builder
	inCode := false
	for _, r := range line {
		if r == '`' {
			inCode = !inCode
			continue
		}
		if !inCode {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func fmValue(frontmatter string, key string) string {
	re := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(key) + `:\s*(.+?)\s*$`)
	match := re.FindStringSubmatch(frontmatter)
	if len(match) != 2 {
		return ""
	}
	return strings.Trim(strings.TrimSpace(match[1]), `"'`)
}

func validKindForPath(kind string, path string) bool {
	switch {
	case path == "wiki/index.md":
		return kind == "index"
	case path == "wiki/log.md":
		return kind == "log"
	case strings.HasPrefix(path, "wiki/sources/"):
		return kind == "source"
	case strings.HasPrefix(path, "wiki/entities/"):
		return kind == "entity"
	case strings.HasPrefix(path, "wiki/concepts/"):
		return kind == "concept"
	case strings.HasPrefix(path, "wiki/topics/"):
		return kind == "topic"
	default:
		return true
	}
}

func requiresSources(path string) bool {
	return strings.HasPrefix(path, "wiki/sources/") ||
		strings.HasPrefix(path, "wiki/entities/") ||
		strings.HasPrefix(path, "wiki/concepts/") ||
		strings.HasPrefix(path, "wiki/topics/")
}

func addIssue(result *Result, code string, level string, path string, message string) {
	result.Issues = append(result.Issues, Issue{Code: code, Level: level, Path: path, Message: message})
}
