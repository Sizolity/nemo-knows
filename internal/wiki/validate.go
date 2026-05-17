package wiki

import (
	"strings"
	"unicode"
)

// IsRawPath reports whether path points at a repository raw source path.
//
// The path argument is expected to be repository-relative.
//
// The returned bool is true only when the path is valid for immutable source
// material under raw/.
//
// The error is returned if the path cannot be validated.
func IsRawPath(path string) (bool, error) {
	return hasPathPrefix(path, RawDir), nil
}

// IsWikiPath reports whether path points at a repository wiki path.
//
// The path argument is expected to be repository-relative.
//
// The returned bool is true only when the path is valid for maintained
// Markdown content under wiki/.
//
// The error is returned if the path cannot be validated.
func IsWikiPath(path string) (bool, error) {
	return hasPathPrefix(path, WikiDir), nil
}

// HasFrontmatter reports whether markdown contains YAML frontmatter.
//
// The markdown argument is the full Markdown document content.
//
// The returned bool is true when the content has a frontmatter block.
//
// The error is returned if the frontmatter cannot be parsed or validated.
func HasFrontmatter(markdown string) (bool, error) {
	lines := strings.Split(strings.TrimLeftFunc(markdown, unicode.IsSpace), "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return false, nil
	}

	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "---" {
			return true, nil
		}
	}

	return false, nil
}

func hasPathPrefix(path string, prefix string) bool {
	clean := strings.TrimPrefix(path, "./")
	return clean == prefix || strings.HasPrefix(clean, prefix+"/")
}
