package wiki

import "testing"

// testResolver builds a LinkResolver from a slug -> repo path map, deriving a
// title from the slug so the default-label path is exercised.
func testResolver(bySlug map[string]string) LinkResolver {
	return func(slug string) (LinkTarget, bool) {
		repoPath, ok := bySlug[slug]
		if !ok {
			return LinkTarget{}, false
		}
		return LinkTarget{RepoPath: repoPath, Title: titleCaseForTest(slug)}, true
	}
}

func titleCaseForTest(slug string) string {
	switch slug {
	case "llama-cpp-hardware-backends":
		return "Hardware Backends"
	case "data-preservation-formats":
		return "Data Preservation Formats"
	default:
		return slug
	}
}

func TestConvertWikilinksCrossDirectoryUsesRelativePath(t *testing.T) {
	resolve := testResolver(map[string]string{
		"llama-cpp-hardware-backends": "wiki/concepts/llama-cpp-hardware-backends.md",
	})
	got := ConvertWikilinks("wiki/entities/llama-cpp.md",
		"Backends are described in [[llama-cpp-hardware-backends]].", resolve)
	want := "Backends are described in [Hardware Backends](../concepts/llama-cpp-hardware-backends.md)."
	if got != want {
		t.Fatalf("cross-directory convert = %q, want %q", got, want)
	}
}

func TestConvertWikilinksSameDirectoryDropsDotPrefix(t *testing.T) {
	resolve := testResolver(map[string]string{
		"gguf": "wiki/entities/gguf.md",
	})
	got := ConvertWikilinks("wiki/entities/llama-cpp.md", "Uses the [[gguf]] format.", resolve)
	want := "Uses the [gguf](gguf.md) format."
	if got != want {
		t.Fatalf("same-directory convert = %q, want %q", got, want)
	}
}

func TestConvertWikilinksExplicitLabelWins(t *testing.T) {
	resolve := testResolver(map[string]string{
		"sqlite-wal": "wiki/sources/sqlite-wal.md",
	})
	got := ConvertWikilinks("wiki/concepts/wal.md", "See [[sqlite-wal|SQLite WAL]] for details.", resolve)
	want := "See [SQLite WAL](../sources/sqlite-wal.md) for details."
	if got != want {
		t.Fatalf("labelled convert = %q, want %q", got, want)
	}
}

func TestConvertWikilinksMissingTargetDegradesToPlainText(t *testing.T) {
	resolve := testResolver(map[string]string{
		"llama-cpp-hardware-backends": "wiki/concepts/llama-cpp-hardware-backends.md",
	})
	got := ConvertWikilinks("wiki/entities/llama-cpp.md", "Uses the [[gguf]] format.", resolve)
	want := "Uses the gguf format."
	if got != want {
		t.Fatalf("missing-target convert = %q, want %q", got, want)
	}
}

func TestConvertWikilinksMissingTargetKeepsExplicitLabel(t *testing.T) {
	got := ConvertWikilinks("wiki/entities/llama-cpp.md", "Uses [[gguf|the GGUF format]].", testResolver(nil))
	want := "Uses the GGUF format."
	if got != want {
		t.Fatalf("missing-target labelled convert = %q, want %q", got, want)
	}
}

func TestConvertWikilinksPreservesAnchor(t *testing.T) {
	resolve := testResolver(map[string]string{
		"sqlite-wal": "wiki/sources/sqlite-wal.md",
	})
	got := ConvertWikilinks("wiki/sources/other.md", "Jump to [[sqlite-wal#checkpoints]].", resolve)
	want := "Jump to [sqlite-wal](sqlite-wal.md#checkpoints)."
	if got != want {
		t.Fatalf("anchor convert = %q, want %q", got, want)
	}
}

func TestConvertWikilinksLeavesCodeUntouched(t *testing.T) {
	resolve := testResolver(map[string]string{"gguf": "wiki/entities/gguf.md"})
	body := "Inline `[[gguf]]` stays.\n\n```\n[[gguf]] in a fence stays\n```\n\nProse [[gguf]] converts."
	got := ConvertWikilinks("wiki/entities/llama-cpp.md", body, resolve)
	want := "Inline `[[gguf]]` stays.\n\n```\n[[gguf]] in a fence stays\n```\n\nProse [gguf](gguf.md) converts."
	if got != want {
		t.Fatalf("code-aware convert = %q, want %q", got, want)
	}
}

func TestRelativeLinkPath(t *testing.T) {
	cases := []struct {
		from string
		to   string
		want string
	}{
		{"wiki/entities/llama-cpp.md", "wiki/concepts/llama-cpp-hardware-backends.md", "../concepts/llama-cpp-hardware-backends.md"},
		{"wiki/entities/llama-cpp.md", "wiki/entities/gguf.md", "gguf.md"},
		{"wiki/entities/sqlite.md", "wiki/topics/data-preservation-formats.md", "../topics/data-preservation-formats.md"},
		{"wiki/index.md", "wiki/sources/sqlite-wal.md", "sources/sqlite-wal.md"},
	}
	for _, tc := range cases {
		if got := RelativeLinkPath(tc.from, tc.to); got != tc.want {
			t.Fatalf("RelativeLinkPath(%q, %q) = %q, want %q", tc.from, tc.to, got, tc.want)
		}
	}
}

func TestRepoPathFromHref(t *testing.T) {
	cases := []struct {
		from   string
		href   string
		want   string
		wantOK bool
	}{
		{"wiki/entities/llama-cpp.md", "../concepts/x.md", "wiki/concepts/x.md", true},
		{"wiki/entities/llama-cpp.md", "gguf.md", "wiki/entities/gguf.md", true},
		{"wiki/index.md", "sources/sqlite-wal.md", "wiki/sources/sqlite-wal.md", true},
		{"wiki/sources/x.md", "../sources/y.md#anchor", "wiki/sources/y.md", true},
		{"wiki/sources/x.md", "https://example.com/page", "", false},
		{"wiki/sources/x.md", "#section", "", false},
		{"wiki/sources/x.md", "../assets/diagram.png", "", false},
		{"wiki/sources/x.md", "mailto:a@b.com", "", false},
	}
	for _, tc := range cases {
		got, ok := RepoPathFromHref(tc.from, tc.href)
		if ok != tc.wantOK || got != tc.want {
			t.Fatalf("RepoPathFromHref(%q, %q) = (%q, %v), want (%q, %v)", tc.from, tc.href, got, ok, tc.want, tc.wantOK)
		}
	}
}

func TestMarkdownLinkTargetsSkipsImagesAndCode(t *testing.T) {
	content := "A [link](../concepts/x.md) and an image ![alt](../assets/y.png).\n\n`[code](z.md)`\n\n```\n[fenced](w.md)\n```\n"
	targets := MarkdownLinkTargets(content)
	if len(targets) != 1 || targets[0] != "../concepts/x.md" {
		t.Fatalf("MarkdownLinkTargets = %v, want [../concepts/x.md]", targets)
	}
}

func TestWikilinkReferencesIgnoresCode(t *testing.T) {
	content := "Prose [[real-ref]] here.\n\n`[[code-ref]]`\n\n```\n[[fenced-ref]]\n```\n"
	refs := WikilinkReferences(content)
	if len(refs) != 1 || refs[0] != "real-ref" {
		t.Fatalf("WikilinkReferences = %v, want [real-ref]", refs)
	}
}
