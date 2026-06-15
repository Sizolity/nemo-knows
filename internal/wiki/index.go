package wiki

import (
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	legacyIndexEntryRE   = regexp.MustCompile(`^\s*-\s*\[\[([^\]|#]+)(?:[|#][^\]]*)?\]\].*$`)
	markdownIndexEntryRE = regexp.MustCompile(`^\s*-\s*\[[^\]]+\]\(([^)\s]+)(?:\s+["'][^)]*["'])?\).*$`)
)

// FormatIndexEntry returns the canonical navigation entry for wiki/index.md.
func FormatIndexEntry(repoPath string, title string) string {
	return fmt.Sprintf("- %s — %s.", FormatIndexLink(repoPath), strings.TrimSpace(title))
}

// FormatIndexLink returns a Markdown link relative to wiki/index.md.
func FormatIndexLink(repoPath string) string {
	return fmt.Sprintf("[%s](%s)", SlugFromWikiPath(repoPath), IndexRelativePath(repoPath))
}

// IndexRelativePath returns the target path for links written in wiki/index.md.
func IndexRelativePath(repoPath string) string {
	clean := path.Clean(strings.TrimPrefix(filepath.ToSlash(repoPath), "./"))
	return strings.TrimPrefix(clean, "wiki/")
}

// IndexEntrySlug extracts a page slug from canonical Markdown index entries
// and legacy [[slug]] entries so maintenance can migrate older indexes.
func IndexEntrySlug(line string) (string, bool) {
	if match := markdownIndexEntryRE.FindStringSubmatch(line); len(match) == 2 {
		return SlugFromReference(match[1]), true
	}
	if match := legacyIndexEntryRE.FindStringSubmatch(line); len(match) == 2 {
		return SlugFromReference(match[1]), true
	}
	return "", false
}

// SlugFromWikiPath returns the filename slug for a maintained wiki page.
func SlugFromWikiPath(repoPath string) string {
	base := path.Base(filepath.ToSlash(repoPath))
	return strings.TrimSuffix(base, path.Ext(base))
}

// SlugFromReference normalizes a wikilink or Markdown target into a page slug.
func SlugFromReference(ref string) string {
	slug := strings.TrimSpace(ref)
	if cut := strings.IndexAny(slug, "#?"); cut != -1 {
		slug = slug[:cut]
	}
	slug = strings.Trim(slug, "<>")
	slug = path.Base(filepath.ToSlash(slug))
	slug = strings.TrimSuffix(slug, path.Ext(slug))
	slug = strings.ToLower(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
