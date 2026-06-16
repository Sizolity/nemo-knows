package wikilint

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	wikischema "github.com/huic/nemo-knows/internal/wiki"
)

var (
	frontmatterBlockRE = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*`)
	logHeadingRE       = regexp.MustCompile(`^## \[[0-9]{4}-[0-9]{2}-[0-9]{2}\] ([^ |]+) \| .+`)
	// imageRE matches a Markdown image ![alt](src); group 1 captures the src.
	imageRE = regexp.MustCompile(`!\[[^\]]*\]\(([^)\s]+)\)`)
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
	// LinkRefs holds every local Markdown relative link (e.g. [x](../sub/x.md))
	// resolved to a repo-relative path, used for target-existence and inbound
	// (orphan) checks.
	LinkRefs []linkRef
	// Wikilinks holds any residual Obsidian [[...]] cross-references found
	// outside code. They are forbidden after the migration to Markdown links.
	Wikilinks []string
}

// linkRef is a single local Markdown relative link on a page.
type linkRef struct {
	Href     string // raw href as written, e.g. "../topics/x.md"
	RepoPath string // resolved repo-relative slash path, e.g. "wiki/topics/x.md"
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
		// Slugs must be unique across the whole wiki so a [[slug]] wikilink
		// resolves to a single page. Flag cross-category collisions (for example
		// a source and an entity both named "sqlite") instead of silently
		// keeping whichever page was walked last.
		if existing, ok := slugToPath[page.Slug]; ok {
			addIssue(&result, "duplicate-slug", "error", page.Path, "slug also used by "+existing+"; slugs must be unique across the wiki")
			continue
		}
		slugToPath[page.Slug] = page.Path
	}

	for _, page := range pages {
		lintFrontmatter(page, &result)
		if refs := dedupeStrings(page.Wikilinks); len(refs) > 0 {
			addIssue(&result, "forbidden-wikilink", "error", page.Path,
				"Obsidian [[wikilink]] syntax is no longer allowed; use Markdown relative links instead: [["+strings.Join(refs, "]], [[")+"]]")
		}
		for _, ref := range page.LinkRefs {
			if ref.RepoPath == "" || strings.HasPrefix(ref.RepoPath, "../") || ref.RepoPath == ".." {
				addIssue(&result, "missing-link-target", "error", page.Path, "relative link escapes the repository: "+ref.Href)
				continue
			}
			if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(ref.RepoPath))); err != nil {
				addIssue(&result, "missing-link-target", "error", page.Path, "relative link target does not exist: "+ref.Href+" (resolved to "+ref.RepoPath+")")
				continue
			}
			if strings.HasPrefix(ref.RepoPath, "wiki/") {
				inbound[wikischema.SlugFromWikiPath(ref.RepoPath)]++
			}
		}
	}
	lintImages(root, pages, &result)
	lintIndex(root, &result)
	lintLog(root, &result)
	for _, page := range pages {
		if isContractDoc(page.Path) {
			continue
		}
		if inbound[page.Slug] == 0 {
			addIssue(&result, "orphan-page", "info", page.Path, "page has no inbound Markdown links")
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
		pagePath := filepath.ToSlash(rel)
		fm, _ := splitFrontmatter(string(content))
		refs := []linkRef{}
		for _, href := range wikischema.MarkdownLinkTargets(string(content)) {
			if resolved, ok := wikischema.RepoPathFromHref(pagePath, href); ok {
				refs = append(refs, linkRef{Href: href, RepoPath: resolved})
			}
		}
		pages = append(pages, page{
			Path:        pagePath,
			Slug:        strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
			Content:     string(content),
			Frontmatter: fm,
			LinkRefs:    refs,
			Wikilinks:   wikischema.WikilinkReferences(string(content)),
		})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk wiki: %w", err)
	}
	return pages, nil
}

func lintFrontmatter(page page, result *Result) {
	// Contract/skeleton documents are not knowledge pages and are exempt from
	// the generic frontmatter contract (see isContractDoc).
	if isContractDoc(page.Path) {
		return
	}
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
	// Index entries are deduped by slug: with wiki-wide slug uniqueness enforced
	// (see the duplicate-slug check in LintWiki), one slug maps to exactly one
	// page, so a repeated slug here means the same page is listed twice. Genuine
	// cross-category slug collisions are reported as duplicate-slug instead.
	seen := map[string]bool{}
	for _, line := range strings.Split(stripMarkdownCode(string(content)), "\n") {
		slug, ok := wikischema.IndexEntrySlug(line)
		if !ok {
			continue
		}
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

// lintImages flags Markdown image references whose local target file is
// missing. Remote http/https images are skipped. Local targets are resolved
// relative to the referencing page (or taken as repo-relative when rooted at
// wiki/) and checked for existence on disk under root; a target that escapes
// the repository or does not exist is reported as a warning so the broken
// reference is visible without blocking the wiki.
func lintImages(root string, pages []page, result *Result) {
	for _, page := range pages {
		body := stripMarkdownCode(page.Content)
		for _, match := range imageRE.FindAllStringSubmatch(body, -1) {
			target := strings.TrimSpace(match[1])
			repoPath, local := imageRepoPath(page.Path, target)
			if !local {
				continue
			}
			if repoPath == "" || repoPath == ".." || strings.HasPrefix(repoPath, "../") {
				addIssue(result, "missing-image", "warn", page.Path, "image reference escapes the wiki: "+target)
				continue
			}
			full := filepath.Join(root, filepath.FromSlash(repoPath))
			if _, err := os.Stat(full); err != nil {
				addIssue(result, "missing-image", "warn", page.Path, "image file does not exist: "+target+" (resolved to "+repoPath+")")
			}
		}
	}
}

// imageRepoPath resolves a Markdown image src to a repo-relative path. The
// boolean result reports whether the target is a local reference that should be
// existence-checked; http/https and other-scheme/protocol-relative targets
// return false so they are skipped.
func imageRepoPath(pagePath string, target string) (string, bool) {
	trimmed := strings.TrimSpace(target)
	if trimmed == "" {
		return "", false
	}
	lower := strings.ToLower(trimmed)
	if strings.HasPrefix(lower, "http://") || strings.HasPrefix(lower, "https://") {
		return "", false
	}
	if strings.Contains(trimmed, "://") || strings.HasPrefix(trimmed, "//") {
		return "", false
	}
	if cut := strings.IndexAny(trimmed, "#?"); cut != -1 {
		trimmed = trimmed[:cut]
	}
	if trimmed == "" {
		return "", false
	}
	repoPath := filepath.ToSlash(trimmed)
	if !strings.HasPrefix(repoPath, "wiki/") {
		repoPath = path.Clean(path.Join(path.Dir(filepath.ToSlash(pagePath)), repoPath))
	} else {
		repoPath = path.Clean(repoPath)
	}
	return repoPath, true
}

func splitFrontmatter(content string) (string, string) {
	match := frontmatterBlockRE.FindStringSubmatch(content)
	if len(match) != 2 {
		return "", content
	}
	return match[1], frontmatterBlockRE.ReplaceAllString(content, "")
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

// isContractDoc reports whether path is a structural/contract document rather
// than a knowledge page. These are the wiki index, the append-only audit log,
// and the agent contract (wiki/AGENTS.md). They are exempt from the generic
// frontmatter and orphan-page checks: index.md and log.md have their own
// dedicated linters (lintIndex/lintLog) and structural roles, while AGENTS.md
// is prose contract documentation with no YAML frontmatter and no inbound
// links. The exemption is intentionally limited to these exact paths so
// ordinary knowledge pages (sources/entities/concepts/topics) still require
// frontmatter and inbound links.
func isContractDoc(path string) bool {
	switch path {
	case "wiki/index.md", "wiki/log.md", "wiki/AGENTS.md":
		return true
	default:
		return false
	}
}

func addIssue(result *Result, code string, level string, path string, message string) {
	result.Issues = append(result.Issues, Issue{Code: code, Level: level, Path: path, Message: message})
}

// dedupeStrings returns the input with duplicates removed, preserving first-seen
// order, so a page that repeats the same forbidden reference is reported once.
func dedupeStrings(in []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(in))
	for _, s := range in {
		if seen[s] {
			continue
		}
		seen[s] = true
		out = append(out, s)
	}
	return out
}
