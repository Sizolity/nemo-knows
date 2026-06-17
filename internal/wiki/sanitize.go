package wiki

import (
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// The sanitizer neutralizes link/cross-reference syntax that violates the
// wiki's body invariants (0 [[wikilink]], 0 {{placeholder}}, no path-traversal
// or dangerous-scheme links) while leaving every legitimate construct — plain
// text, http(s) URLs, in-repo relative .md links, and #anchors — untouched.
//
// It exists because source/ingest drafts copy links verbatim out of the raw
// document (see prompts/source-page.md). On well-formed sources the links are
// already clean, so this is adversarial-only data cleaning: it makes malformed
// model output conform to the Markdown-link format the wiki requires instead of
// rejecting the draft.
//
// CRITICAL: sanitization runs strictly outside inline code spans and fenced
// code blocks (via rewriteOutsideCode / StripCode). Real technical sources
// legitimately contain tokens such as the Web IDL internal slot
// `[[ArrayBufferData]]` or a literal `javascript:` inside code samples, and
// those must survive byte-for-byte.

var (
	// placeholderRE matches a {{...}} template placeholder token. The body of a
	// rendered wiki page must never contain template artifacts, so any residual
	// placeholder is flattened to its inner text.
	placeholderRE = regexp.MustCompile(`\{\{([^{}]*)\}\}`)

	// linkTargetPattern matches a Markdown link/image target, allowing a single
	// level of balanced parentheses so adversarial targets like
	// javascript:alert(1) or base64 data: URIs are captured in full and fully
	// neutralized rather than partially matched (which would leave dangling
	// fragments behind).
	linkTargetPattern = `(?:[^()\s]|\([^()]*\))*`

	// sanitizeImageRE matches a Markdown image ![alt](target); group 1 is the
	// alt text and group 2 is the target.
	sanitizeImageRE = regexp.MustCompile(`!\[([^\]]*)\]\((` + linkTargetPattern + `)\)`)

	// sanitizeLinkRE matches an inline Markdown link [label](target). The leading
	// (^|[^!]) keeps it from matching image syntax; group 1 is that guard char,
	// group 2 is the label, and group 3 is the target.
	sanitizeLinkRE = regexp.MustCompile(`(^|[^!])\[([^\]]+)\]\((` + linkTargetPattern + `)\)`)

	// schemeRE matches a leading URI scheme, e.g. "https:" or "javascript:".
	schemeRE = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+.\-]*:`)
)

// safeSchemes are the URI schemes the sanitizer preserves on a link target.
// Anything else carrying a scheme (javascript:, data:, file:, vbscript:, and
// other unexpected schemes) is stripped at generation time because a
// well-formed source only ever yields http(s)/mailto links or in-repo relative
// references; preserving unknown schemes would risk leaking unsafe targets.
var safeSchemes = map[string]bool{
	"http":   true,
	"https":  true,
	"mailto": true,
}

// dangerousSchemes are the URI schemes the scanner explicitly flags as unsafe
// link targets. This is intentionally a precise, well-known set (rather than
// "any non-allowlisted scheme") so detection over human-curated wiki pages does
// not false-positive on legitimate but unusual schemes.
var dangerousSchemes = map[string]bool{
	"javascript": true,
	"data":       true,
	"file":       true,
	"vbscript":   true,
}

// UnsafeLink is a single forbidden link/image target found by UnsafeLinkTargets.
type UnsafeLink struct {
	// Target is the raw href/src as written in the page.
	Target string
	// Reason classifies the finding: "dangerous-scheme" or "path-traversal".
	Reason string
	// Scheme is the offending URI scheme when Reason is "dangerous-scheme".
	Scheme string
}

const (
	// ReasonDangerousScheme marks a link target using a dangerous URI scheme.
	ReasonDangerousScheme = "dangerous-scheme"
	// ReasonPathTraversal marks a relative link target that escapes the repo.
	ReasonPathTraversal = "path-traversal"
)

// SanitizeBodyLinks rewrites a wiki/draft body so it conforms to the wiki's
// Markdown-link format invariants, leaving code spans and fenced blocks intact.
// It is the single shared cleaner reused by source/ingest generation today and
// is exported so a future LLM-based maintainer can apply the identical rules.
//
// Transformations (applied only outside code):
//  1. [[wikilink]] / [[target|label]] -> plain visible text. A draft has no
//     slug map, so a nil resolver degrades every cross-reference to its label.
//  2. {{placeholder}} -> its inner text (the {{}} braces are removed).
//  3. [label](target) and ![alt](target) with an UNSAFE target -> plain
//     label/alt, dropping the target. UNSAFE = a relative path escaping the
//     repository root via ../, or a scheme outside the http/https/mailto
//     allowlist (e.g. javascript:, data:, file:, vbscript:).
//
// Everything safe is preserved: plain text, http(s)/mailto URLs, in-repo
// relative .md links, protocol-relative URLs, and #anchors.
//
// fromRepoPath is the repo-relative slash path the body will live at; it sets
// the directory depth used to decide whether a relative target escapes the
// repository (mirroring RepoPathFromHref/ConvertWikilinks).
func SanitizeBodyLinks(fromRepoPath string, body string) string {
	// Flatten cross-references first using the shared migration primitive with a
	// nil resolver so all [[...]] degrade to plain text.
	body = ConvertWikilinks(fromRepoPath, body, nil)
	return SanitizeConvertedBody(fromRepoPath, body)
}

// SanitizeConvertedBody enforces the body-link invariants that do not involve
// wikilink resolution: it flattens {{placeholder}} tokens to their inner text
// and drops unsafe link/image targets (dangerous scheme or repo-escaping
// relative path), always operating strictly outside code spans and fenced
// blocks.
//
// It is split out of SanitizeBodyLinks for callers that must resolve [[...]]
// cross-references with their own LinkResolver first — candidate generation, for
// example, turns [[slug]] into real Markdown relative links via ConvertWikilinks
// and then calls this to apply the remaining invariants without re-running
// wikilink conversion with a nil resolver (which would discard those links).
func SanitizeConvertedBody(fromRepoPath string, body string) string {
	return rewriteOutsideCode(body, func(segment string) string {
		segment = stripPlaceholders(segment)
		segment = flattenUnsafeImages(fromRepoPath, segment)
		segment = flattenUnsafeLinks(fromRepoPath, segment)
		return segment
	})
}

// PlaceholderTokens returns the {{...}} template placeholders that remain in
// content outside code spans and fenced blocks, in document order. A non-empty
// result means generation-time sanitization (SanitizeBodyLinks, which flattens
// these) was skipped or bypassed; the apply write-gate uses it to refuse such
// content. Placeholders inside inline code or fenced blocks are ignored so
// legitimate examples survive verbatim.
func PlaceholderTokens(content string) []string {
	tokens := []string{}
	tokens = append(tokens, placeholderRE.FindAllString(StripCode(content), -1)...)
	return tokens
}

// stripPlaceholders flattens {{...}} tokens in segment to their trimmed inner
// text. Segment is assumed to already be outside any code region.
func stripPlaceholders(segment string) string {
	return placeholderRE.ReplaceAllStringFunc(segment, func(match string) string {
		return strings.TrimSpace(match[2 : len(match)-2])
	})
}

// flattenUnsafeImages replaces ![alt](target) with alt when target is unsafe.
func flattenUnsafeImages(fromRepoPath string, segment string) string {
	return sanitizeImageRE.ReplaceAllStringFunc(segment, func(match string) string {
		sub := sanitizeImageRE.FindStringSubmatch(match)
		if len(sub) != 3 {
			return match
		}
		if unsafeForSanitize(fromRepoPath, sub[2]) {
			return sub[1]
		}
		return match
	})
}

// flattenUnsafeLinks replaces [label](target) with label when target is unsafe,
// preserving the non-image guard character captured before the link.
func flattenUnsafeLinks(fromRepoPath string, segment string) string {
	return sanitizeLinkRE.ReplaceAllStringFunc(segment, func(match string) string {
		sub := sanitizeLinkRE.FindStringSubmatch(match)
		if len(sub) != 4 {
			return match
		}
		if unsafeForSanitize(fromRepoPath, sub[3]) {
			return sub[1] + sub[2]
		}
		return match
	})
}

// unsafeForSanitize reports whether a link/image target should be stripped by
// the generation-time cleaner. It is conservative: it preserves only the
// http/https/mailto schemes, protocol-relative and #anchor targets, and in-repo
// relative paths; anything else (including any other scheme) is unsafe.
func unsafeForSanitize(fromRepoPath string, target string) bool {
	t := strings.TrimSpace(target)
	if t == "" {
		return true
	}
	if strings.HasPrefix(t, "#") || strings.HasPrefix(t, "//") {
		return false
	}
	if scheme := uriScheme(t); scheme != "" {
		return !safeSchemes[scheme]
	}
	return relativeTargetEscapesRepo(fromRepoPath, t)
}

// UnsafeLinkTargets scans content for forbidden link/image targets outside code
// and returns them classified. It powers the lint/scan backstop: dangerous URI
// schemes are flagged on any target, and path-traversal is flagged for relative
// targets that escape the repo. In-repo relative .md links (including escaping
// ones) are deliberately skipped here because the existing missing-link-target
// check already covers them, avoiding duplicate findings.
func UnsafeLinkTargets(fromRepoPath string, content string) []UnsafeLink {
	stripped := StripCode(content)
	findings := []UnsafeLink{}
	for _, target := range markdownTargetsForScan(stripped) {
		if link, ok := classifyTargetForScan(fromRepoPath, target); ok {
			findings = append(findings, link)
		}
	}
	return findings
}

// markdownTargetsForScan returns the targets of every Markdown link and image in
// already-code-stripped content, preserving document order.
func markdownTargetsForScan(stripped string) []string {
	targets := []string{}
	for _, m := range sanitizeImageRE.FindAllStringSubmatch(stripped, -1) {
		targets = append(targets, m[2])
	}
	for _, m := range sanitizeLinkRE.FindAllStringSubmatch(stripped, -1) {
		targets = append(targets, m[3])
	}
	return targets
}

// classifyTargetForScan reports whether target is an unsafe link for the scan
// and, if so, how. It mirrors the detection contract described on
// UnsafeLinkTargets.
func classifyTargetForScan(fromRepoPath string, target string) (UnsafeLink, bool) {
	t := strings.TrimSpace(target)
	if t == "" {
		return UnsafeLink{}, false
	}
	if scheme := uriScheme(t); scheme != "" {
		if dangerousSchemes[scheme] {
			return UnsafeLink{Target: t, Reason: ReasonDangerousScheme, Scheme: scheme}, true
		}
		return UnsafeLink{}, false
	}
	if strings.HasPrefix(t, "//") || strings.HasPrefix(t, "#") {
		return UnsafeLink{}, false
	}
	// In-repo relative .md links are handled by the existing missing-link-target
	// check; don't double-report them as traversal here.
	if _, ok := RepoPathFromHref(fromRepoPath, t); ok {
		return UnsafeLink{}, false
	}
	if relativeTargetEscapesRepo(fromRepoPath, t) {
		return UnsafeLink{Target: t, Reason: ReasonPathTraversal}, true
	}
	return UnsafeLink{}, false
}

// uriScheme returns the lowercased URI scheme of target (without the trailing
// colon), or "" when target carries no scheme. A bare relative path such as
// "dir/page.md" or "sibling.md#sec" has no scheme.
func uriScheme(target string) string {
	match := schemeRE.FindString(target)
	if match == "" {
		return ""
	}
	return strings.ToLower(strings.TrimSuffix(match, ":"))
}

// relativeTargetEscapesRepo reports whether a scheme-less relative target,
// resolved against the directory of fromRepoPath, escapes the repository root
// (resolves to ".." or a "../"-prefixed path). Anchors and query strings are
// stripped before resolution. Targets already rooted at "wiki/" are cleaned in
// place. This mirrors RepoPathFromHref's resolution but works for any target,
// not just ".md" references.
func relativeTargetEscapesRepo(fromRepoPath string, target string) bool {
	trimmed := strings.TrimSpace(target)
	if cut := strings.IndexAny(trimmed, "#?"); cut != -1 {
		trimmed = trimmed[:cut]
	}
	trimmed = strings.TrimSpace(trimmed)
	if trimmed == "" {
		return false
	}
	repoPath := filepath.ToSlash(trimmed)
	if strings.HasPrefix(repoPath, "wiki/") {
		repoPath = path.Clean(repoPath)
	} else {
		repoPath = path.Clean(path.Join(path.Dir(filepath.ToSlash(fromRepoPath)), repoPath))
	}
	return repoPath == ".." || strings.HasPrefix(repoPath, "../")
}
