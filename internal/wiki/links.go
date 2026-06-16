package wiki

import (
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	// wikilinkConvertRE matches an Obsidian-style cross-reference written as
	// [[target]] or [[target|label]]. Group 1 captures the target (a page slug,
	// optionally carrying a ".md" suffix or a "#anchor"); group 2, when present,
	// captures the explicit display label.
	wikilinkConvertRE = regexp.MustCompile(`\[\[([^\]|]+)(?:\|([^\]]*))?\]\]`)

	// wikilinkDetectRE matches any residual [[...]] cross-reference. It backs the
	// migration checks that flag the deprecated Obsidian syntax once the wiki has
	// moved to standard Markdown relative links.
	wikilinkDetectRE = regexp.MustCompile(`\[\[[^\]]*\]\]`)

	// markdownInlineLinkRE matches an inline Markdown link [label](target). The
	// leading (^|[^!]) keeps it from matching image syntax (![alt](src)); group 2
	// is the label and group 3 is the link target.
	markdownInlineLinkRE = regexp.MustCompile(`(^|[^!])\[([^\]]+)\]\(([^)\s]+)\)`)
)

// LinkTarget is a resolved cross-reference destination.
type LinkTarget struct {
	// RepoPath is the repo-relative slash path of the destination page, for
	// example "wiki/concepts/serverless-database.md".
	RepoPath string
	// Title is the human-friendly label used when a reference omits an explicit
	// "|label". It is typically the destination page's frontmatter title.
	Title string
}

// LinkResolver maps a normalized page slug to its destination page. An ok value
// of false means the slug has no destination page in the wiki, so the reference
// must degrade to plain text rather than link to a file that does not exist.
type LinkResolver func(slug string) (LinkTarget, bool)

// ConvertWikilinks rewrites every [[slug]] and [[slug|label]] cross-reference in
// body into a standard Markdown relative link from fromRepoPath to the resolved
// destination page. References whose slug does not resolve degrade to plain-text
// labels (no link) so a link never points at a non-existent file. Text inside
// fenced code blocks and inline code spans is left untouched.
//
// This is the single shared migration primitive reused by the draft cleaner, the
// approved-apply path, and the one-time wiki content migration so they all emit
// identical relative links.
func ConvertWikilinks(fromRepoPath string, body string, resolve LinkResolver) string {
	if resolve == nil {
		resolve = func(string) (LinkTarget, bool) { return LinkTarget{}, false }
	}
	return rewriteOutsideCode(body, func(segment string) string {
		return wikilinkConvertRE.ReplaceAllStringFunc(segment, func(match string) string {
			sub := wikilinkConvertRE.FindStringSubmatch(match)
			if len(sub) != 3 {
				return match
			}
			rawTarget := sub[1]
			anchor := ""
			if i := strings.IndexByte(rawTarget, '#'); i != -1 {
				anchor = rawTarget[i:]
				rawTarget = rawTarget[:i]
			}
			label := strings.TrimSpace(sub[2])
			target, ok := resolve(SlugFromReference(rawTarget))
			if !ok {
				// No destination page: degrade to plain text so the rendered
				// output never carries a dangling link.
				if label != "" {
					return label
				}
				return strings.TrimSpace(rawTarget)
			}
			if label == "" {
				label = strings.TrimSpace(target.Title)
			}
			if label == "" {
				label = strings.TrimSpace(rawTarget)
			}
			return "[" + label + "](" + RelativeLinkPath(fromRepoPath, target.RepoPath) + anchor + ")"
		})
	})
}

// RelativeLinkPath returns the Markdown link target needed to reach toRepoPath
// from the page at fromRepoPath. Both inputs are repo-relative slash paths. The
// result is relative to the directory that contains fromRepoPath, so a sibling
// page is "name.md" and a page in another wiki subdirectory is "../sub/name.md".
func RelativeLinkPath(fromRepoPath string, toRepoPath string) string {
	fromDir := path.Dir(path.Clean(filepath.ToSlash(fromRepoPath)))
	to := path.Clean(filepath.ToSlash(toRepoPath))
	return relativeSlashPath(fromDir, to)
}

func relativeSlashPath(baseDir string, target string) string {
	baseParts := splitSlashPath(baseDir)
	targetParts := splitSlashPath(target)
	common := 0
	for common < len(baseParts) && common < len(targetParts) && baseParts[common] == targetParts[common] {
		common++
	}
	out := make([]string, 0, (len(baseParts)-common)+(len(targetParts)-common))
	for i := common; i < len(baseParts); i++ {
		out = append(out, "..")
	}
	out = append(out, targetParts[common:]...)
	if len(out) == 0 {
		return "."
	}
	return strings.Join(out, "/")
}

func splitSlashPath(p string) []string {
	p = strings.Trim(strings.TrimSpace(p), "/")
	if p == "" || p == "." {
		return nil
	}
	return strings.Split(p, "/")
}

// RepoPathFromHref resolves an inline Markdown link href to a repo-relative
// slash path and reports whether it is a local, in-repo Markdown reference. It
// returns ok=false for empty hrefs, external URLs (http/https/mailto or any
// scheme), protocol-relative URLs, pure "#anchor" fragments, and non-".md"
// targets. Anchors and query strings are stripped before resolution. A target
// already rooted at "wiki/" is cleaned in place; otherwise it is resolved
// against the directory of fromRepoPath. The returned path may escape the repo
// (for example "../foo.md"); callers decide whether such a path is acceptable.
func RepoPathFromHref(fromRepoPath string, href string) (string, bool) {
	trimmed := strings.TrimSpace(href)
	if trimmed == "" {
		return "", false
	}
	lower := strings.ToLower(trimmed)
	if strings.HasPrefix(lower, "http://") || strings.HasPrefix(lower, "https://") || strings.HasPrefix(lower, "mailto:") {
		return "", false
	}
	if strings.Contains(trimmed, "://") || strings.HasPrefix(trimmed, "//") {
		return "", false
	}
	if strings.HasPrefix(trimmed, "#") {
		return "", false
	}
	if cut := strings.IndexAny(trimmed, "#?"); cut != -1 {
		trimmed = trimmed[:cut]
	}
	if trimmed == "" {
		return "", false
	}
	if !strings.EqualFold(path.Ext(trimmed), ".md") {
		return "", false
	}
	repoPath := filepath.ToSlash(trimmed)
	if !strings.HasPrefix(repoPath, "wiki/") {
		repoPath = path.Clean(path.Join(path.Dir(filepath.ToSlash(fromRepoPath)), repoPath))
	} else {
		repoPath = path.Clean(repoPath)
	}
	return repoPath, true
}

// MarkdownLinkTargets returns the raw hrefs of inline Markdown links in content,
// excluding image syntax and anything inside code spans or fenced code blocks.
func MarkdownLinkTargets(content string) []string {
	targets := []string{}
	for _, m := range markdownInlineLinkRE.FindAllStringSubmatch(StripCode(content), -1) {
		targets = append(targets, m[3])
	}
	return targets
}

// WikilinkReferences returns the raw inside-bracket targets of any [[...]]
// cross-references that remain in content, ignoring code spans and fenced code
// blocks. A non-empty result means the deprecated Obsidian syntax is still
// present and should be migrated to Markdown relative links.
func WikilinkReferences(content string) []string {
	refs := []string{}
	for _, m := range wikilinkDetectRE.FindAllString(StripCode(content), -1) {
		refs = append(refs, strings.Trim(m, "[]"))
	}
	return refs
}

// StripCode removes fenced code blocks and inline code spans from content while
// keeping the remaining text on its original lines. It is the shared scanner
// pre-pass so link and cross-reference checks never match inside code samples.
func StripCode(content string) string {
	var b strings.Builder
	inFence := false
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		switch {
		case strings.HasPrefix(strings.TrimSpace(line), "```"):
			inFence = !inFence
		case inFence:
			// drop fenced code content
		default:
			b.WriteString(stripInlineCodeSpans(line))
		}
		if i < len(lines)-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}

// rewriteOutsideCode applies fn to the portions of content that lie outside
// fenced code blocks and inline code spans, leaving every code region (and the
// fence/backtick delimiters themselves) byte-for-byte intact.
func rewriteOutsideCode(content string, fn func(string) string) string {
	var b strings.Builder
	inFence := false
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		switch {
		case strings.HasPrefix(strings.TrimSpace(line), "```"):
			inFence = !inFence
			b.WriteString(line)
		case inFence:
			b.WriteString(line)
		default:
			b.WriteString(rewriteOutsideInlineCode(line, fn))
		}
		if i < len(lines)-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}

func rewriteOutsideInlineCode(line string, fn func(string) string) string {
	var b strings.Builder
	var buf strings.Builder
	inCode := false
	for _, r := range line {
		if r == '`' {
			if inCode {
				b.WriteString(buf.String())
				b.WriteByte('`')
				inCode = false
			} else {
				b.WriteString(fn(buf.String()))
				b.WriteByte('`')
				inCode = true
			}
			buf.Reset()
			continue
		}
		buf.WriteRune(r)
	}
	if inCode {
		// Unterminated inline code span: emit the buffered remainder verbatim so
		// no content is lost and links after a stray backtick stay untouched.
		b.WriteString(buf.String())
	} else {
		b.WriteString(fn(buf.String()))
	}
	return b.String()
}

func stripInlineCodeSpans(line string) string {
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
