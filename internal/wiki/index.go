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

// IndexEmptyPlaceholder is the literal a wiki index shows under a category
// heading that has no entries yet.
const IndexEmptyPlaceholder = "(none yet)"

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

// IndexEntryTarget returns the link target of a canonical Markdown index entry.
// It reports false for legacy [[slug]] entries and non-entry lines.
func IndexEntryTarget(line string) (string, bool) {
	if match := markdownIndexEntryRE.FindStringSubmatch(line); len(match) == 2 {
		return match[1], true
	}
	return "", false
}

// AppendIndexEntry inserts entry under the given "## Section" heading of a wiki
// index document. The first real entry of a section replaces its
// IndexEmptyPlaceholder; when the heading is absent it is appended along with
// the entry. This is the single shared index-append path for the ingest and
// query workflows. The returned document always ends with a trailing newline.
func AppendIndexEntry(index string, section string, entry string) string {
	entryLine := strings.TrimSuffix(entry, "\n")
	lines := strings.Split(index, "\n")

	sectionLine := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == section {
			sectionLine = i
			break
		}
	}
	if sectionLine == -1 {
		trimmed := strings.TrimRight(index, "\n")
		if trimmed == "" {
			return section + "\n\n" + entryLine + "\n"
		}
		return trimmed + "\n\n" + section + "\n\n" + entryLine + "\n"
	}

	end := len(lines)
	for i := sectionLine + 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "## ") {
			end = i
			break
		}
	}

	// Drop the "(none yet)" placeholder so the first real entry replaces it.
	body := make([]string, 0, end-sectionLine)
	for _, line := range lines[sectionLine+1 : end] {
		if strings.TrimSpace(line) == IndexEmptyPlaceholder {
			continue
		}
		body = append(body, line)
	}

	// Insert after the last non-blank body line; if no entries remain, keep a
	// single blank line between the heading and the new entry.
	insertAt := len(body)
	for insertAt > 0 && strings.TrimSpace(body[insertAt-1]) == "" {
		insertAt--
	}
	var newBody []string
	if insertAt == 0 {
		newBody = []string{"", entryLine}
	} else {
		newBody = append(newBody, body[:insertAt]...)
		newBody = append(newBody, entryLine)
	}

	result := make([]string, 0, len(lines)+2)
	result = append(result, lines[:sectionLine+1]...)
	result = append(result, newBody...)
	if end < len(lines) {
		result = append(result, "")
	}
	result = append(result, lines[end:]...)

	out := strings.Join(result, "\n")
	if !strings.HasSuffix(out, "\n") {
		out += "\n"
	}
	return out
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
