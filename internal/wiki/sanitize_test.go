package wiki

import (
	"strings"
	"testing"
)

// fromPage is a representative source-page repo path (depth 2 under the repo
// root) used to anchor relative-target resolution in the sanitizer tests.
const fromPage = "wiki/sources/example.md"

func TestSanitizeBodyLinksTable(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		// Positive cases: forbidden syntax must be neutralized.
		{"wikilink bare", "See [[serverless-database]] here.", "See serverless-database here."},
		{"wikilink labelled", "See [[serverless-database|the engine]] here.", "See the engine here."},
		{"placeholder", "Path is {{RAW_SOURCE_PATH}} now.", "Path is RAW_SOURCE_PATH now."},
		{"placeholder spaced", "Value {{ X }} done.", "Value X done."},
		{"dangerous scheme javascript", "Click [run](javascript:alert(1)) now.", "Click run now."},
		{"dangerous scheme data image", "Logo ![brand](data:image/png;base64,AAAA) shown.", "Logo brand shown."},
		{"dangerous scheme file", "Open [cfg](file:///etc/hosts) please.", "Open cfg please."},
		{"dangerous scheme vbscript", "Run [x](vbscript:msgbox) now.", "Run x now."},
		{"path traversal", "Read [secret](../../../etc/passwd) now.", "Read secret now."},
		{"path traversal image", "See ![p](../../../../etc/shadow) here.", "See p here."},

		// Negative/preserve cases: legitimate content is left intact.
		{"https link", "Visit [site](https://ok.com) today.", "Visit [site](https://ok.com) today."},
		{"mailto link", "Mail [us](mailto:a@b.com) anytime.", "Mail [us](mailto:a@b.com) anytime."},
		{"relative sibling md", "See [sibling](sibling.md) page.", "See [sibling](sibling.md) page."},
		{"relative dir md with anchor", "See [page](dir/page.md#sec) section.", "See [page](dir/page.md#sec) section."},
		{"cross dir md", "See [topic](../topics/x.md) page.", "See [topic](../topics/x.md) page."},
		{"anchor only", "Jump [down](#section) now.", "Jump [down](#section) now."},
		{"remote image", "Logo ![brand](https://cdn.example.com/l.png) shown.", "Logo ![brand](https://cdn.example.com/l.png) shown."},
		{"plain text", "Just prose with no links at all.", "Just prose with no links at all."},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SanitizeBodyLinks(fromPage, tc.in); got != tc.want {
				t.Fatalf("SanitizeBodyLinks(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

// TestSanitizeBodyLinksPreservesCode is the critical code-span-safety guarantee:
// forbidden-looking tokens inside inline code and fenced code blocks (e.g. the
// Web IDL internal slot [[ArrayBufferData]]) MUST survive verbatim, while the
// same tokens in prose are sanitized.
func TestSanitizeBodyLinksPreservesCode(t *testing.T) {
	in := strings.Join([]string{
		"Prose [[wikilink]] and {{TOKEN}} and [x](javascript:alert(1)).",
		"",
		"Inline code `[[ArrayBufferData]]` and `{{NotASlot}}` and `javascript:noop` survive.",
		"",
		"```js",
		"const slot = obj[[ArrayBufferData]]; // {{TEMPLATE}}",
		"const u = \"javascript:alert(1)\";",
		"const evil = \"[t](../../../etc/passwd)\";",
		"```",
		"",
		"Trailing prose [[again]] cleaned.",
	}, "\n")

	want := strings.Join([]string{
		"Prose wikilink and TOKEN and x.",
		"",
		"Inline code `[[ArrayBufferData]]` and `{{NotASlot}}` and `javascript:noop` survive.",
		"",
		"```js",
		"const slot = obj[[ArrayBufferData]]; // {{TEMPLATE}}",
		"const u = \"javascript:alert(1)\";",
		"const evil = \"[t](../../../etc/passwd)\";",
		"```",
		"",
		"Trailing prose again cleaned.",
	}, "\n")

	if got := SanitizeBodyLinks(fromPage, in); got != want {
		t.Fatalf("SanitizeBodyLinks code-span safety failed:\n got=%q\nwant=%q", got, want)
	}
}

// TestSanitizeBodyLinksIsIdempotent ensures a second pass over already-clean
// output is a no-op, so applying the cleaner repeatedly never degrades content.
func TestSanitizeBodyLinksIsIdempotent(t *testing.T) {
	in := "Mix [[ref]] {{T}} [bad](javascript:x) [ok](https://ok.com) `[[keep]]`."
	once := SanitizeBodyLinks(fromPage, in)
	twice := SanitizeBodyLinks(fromPage, once)
	if once != twice {
		t.Fatalf("sanitizer not idempotent:\nonce=%q\ntwice=%q", once, twice)
	}
	if strings.Contains(once, "javascript:") || strings.Contains(once, "{{") {
		t.Fatalf("sanitizer left forbidden syntax outside code: %q", once)
	}
	if !strings.Contains(once, "`[[keep]]`") {
		t.Fatalf("sanitizer mangled inline code: %q", once)
	}
}

func TestUnsafeLinkTargets(t *testing.T) {
	content := strings.Join([]string{
		"---",
		"kind: source",
		"---",
		"",
		"Prose [bad-js](javascript:alert(1)) and [bad-data](data:text/html,x).",
		"Traversal [esc](../../../etc/passwd) and image ![e](../../../../etc/shadow).",
		"Safe [site](https://ok.com), [mail](mailto:a@b.com), [sib](sibling.md), [anc](#top).",
		"In-repo md escape [md](../../../other.md) is left to missing-link-target.",
		"",
		"`[ignore](javascript:inline)` and:",
		"```",
		"[fenced](javascript:fenced) and [t](../../../etc/passwd)",
		"```",
	}, "\n")

	got := UnsafeLinkTargets("wiki/sources/page.md", content)

	wantSchemes := map[string]bool{"javascript:alert(1)": true, "data:text/html,x": true}
	wantTraversal := map[string]bool{"../../../etc/passwd": true, "../../../../etc/shadow": true}

	gotSchemes := map[string]bool{}
	gotTraversal := map[string]bool{}
	for _, f := range got {
		switch f.Reason {
		case ReasonDangerousScheme:
			gotSchemes[f.Target] = true
		case ReasonPathTraversal:
			gotTraversal[f.Target] = true
		default:
			t.Fatalf("unexpected reason %q for %q", f.Reason, f.Target)
		}
	}

	if !mapsEqual(gotSchemes, wantSchemes) {
		t.Fatalf("dangerous-scheme findings = %v, want %v (all: %#v)", gotSchemes, wantSchemes, got)
	}
	if !mapsEqual(gotTraversal, wantTraversal) {
		t.Fatalf("path-traversal findings = %v, want %v (all: %#v)", gotTraversal, wantTraversal, got)
	}

	// The dangerous scheme finding must carry the scheme name.
	for _, f := range got {
		if f.Reason == ReasonDangerousScheme && f.Scheme == "" {
			t.Fatalf("dangerous-scheme finding missing scheme: %#v", f)
		}
	}
}

func TestUnsafeLinkTargetsCleanContentReturnsNone(t *testing.T) {
	content := strings.Join([]string{
		"# Clean Page",
		"",
		"Visit [site](https://ok.com) and [sibling](sibling.md#sec).",
		"Image ![ok](../assets/logo.png) and anchor [top](#intro).",
		"`[[code-only]]` stays and so does:",
		"```",
		"javascript:not-flagged",
		"```",
	}, "\n")
	if got := UnsafeLinkTargets("wiki/sources/page.md", content); len(got) != 0 {
		t.Fatalf("expected no unsafe targets on clean content, got %#v", got)
	}
}

func mapsEqual(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}

// TestSanitizeConvertedBody covers the post-wikilink-resolution stage used by
// candidate generation: it must flatten {{placeholder}} tokens and drop unsafe
// link/image targets while leaving already-resolved [[...]] (the caller's job)
// and safe relative links alone, and never touch code spans.
func TestSanitizeConvertedBody(t *testing.T) {
	in := "Keep [[raw]] as-is, drop {{TOK}}, strip [x](javascript:bad), keep [ok](ok.md), code `{{KEEP}}`."
	want := "Keep [[raw]] as-is, drop TOK, strip x, keep [ok](ok.md), code `{{KEEP}}`."
	if got := SanitizeConvertedBody(fromPage, in); got != want {
		t.Fatalf("SanitizeConvertedBody = %q, want %q", got, want)
	}
}

// TestPlaceholderTokens verifies the apply-gate detector finds {{...}} tokens in
// prose but ignores those confined to inline code or fenced blocks.
func TestPlaceholderTokens(t *testing.T) {
	content := strings.Join([]string{
		"Prose {{ALPHA}} and {{ BETA }} remain visible.",
		"Inline `{{GAMMA}}` and fenced:",
		"```",
		"{{DELTA}}",
		"```",
	}, "\n")
	if got := PlaceholderTokens(content); len(got) != 2 {
		t.Fatalf("PlaceholderTokens = %#v, want 2 tokens outside code", got)
	}
	if got := PlaceholderTokens("No tokens, just `{{code}}` and a fenced block:\n```\n{{x}}\n```\n"); len(got) != 0 {
		t.Fatalf("PlaceholderTokens on code-only content = %#v, want none", got)
	}
}
