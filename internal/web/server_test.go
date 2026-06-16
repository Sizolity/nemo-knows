package web

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestValidateMarkdownUploadAcceptsBasicMarkdown(t *testing.T) {
	name, content, err := validateMarkdownUpload("my-note.md", "# My Note\n\nBody", time.Time{})
	if err != nil {
		t.Fatalf("validateMarkdownUpload returned error: %v", err)
	}
	if name != "my-note.md" {
		t.Fatalf("name = %q, want my-note.md", name)
	}
	if !strings.HasSuffix(content, "\n") {
		t.Fatalf("content should be newline-terminated: %q", content)
	}
}

func TestValidateMarkdownUploadDerivesNameFromTitleOrTimestamp(t *testing.T) {
	name, _, err := validateMarkdownUpload("", "# My New Note\n\nBody", time.Time{})
	if err != nil {
		t.Fatalf("validateMarkdownUpload returned error: %v", err)
	}
	if name != "my-new-note.md" {
		t.Fatalf("name = %q, want my-new-note.md", name)
	}

	now := time.Date(2026, 5, 24, 19, 58, 0, 0, time.UTC)
	name, _, err = validateMarkdownUpload("", "plain Markdown text without heading", now)
	if err != nil {
		t.Fatalf("validateMarkdownUpload returned error: %v", err)
	}
	if name != "source-20260524-195800.md" {
		t.Fatalf("name = %q, want timestamp source name", name)
	}
}

func TestValidateMarkdownUploadRejectsUnsafeInput(t *testing.T) {
	cases := []struct {
		name    string
		content string
	}{
		{"../evil.md", "# Title\n"},
		{"!!!.md", "# Title\n"},
		{"empty.md", ""},
		{"not-md.txt", "# Title\n"},
		{"script.md", "# Title\n\n<script>alert(1)</script>"},
	}
	for _, tc := range cases {
		if _, _, err := validateMarkdownUpload(tc.name, tc.content, time.Time{}); err == nil {
			t.Fatalf("expected %q to be rejected", tc.name)
		}
	}
}

func TestRenderMarkdownResolvesRelativeMarkdownLinks(t *testing.T) {
	html := string(renderMarkdown("wiki/index.md", []byte("- [sqlite-wal](sources/sqlite-wal.md) — notes.\n"), nil))
	if !strings.Contains(html, `<a href="/view?path=wiki%2Fsources%2Fsqlite-wal.md">sqlite-wal</a>`) {
		t.Fatalf("expected rendered index link, got:\n%s", html)
	}
}

func TestRenderMarkdownStillResolvesSemanticWikilinks(t *testing.T) {
	html := string(renderMarkdown("wiki/concepts/wal.md", []byte("See [[sqlite-wal|SQLite WAL]].\n"), map[string]string{
		"sqlite-wal": "wiki/sources/sqlite-wal.md",
	}))
	if !strings.Contains(html, `<a href="/view?path=wiki/sources/sqlite-wal.md" class="wikilink">SQLite WAL</a>`) {
		t.Fatalf("expected rendered wikilink, got:\n%s", html)
	}
}

func TestRenderMarkdownRendersRelativeImage(t *testing.T) {
	html := string(renderMarkdown("wiki/sources/slides.md",
		[]byte("![Beta ISA Summary](../assets/c10s1-slides/Slide02.png)\n"), nil))
	want := `<img src="/assets/c10s1-slides/Slide02.png" alt="Beta ISA Summary" loading="lazy">`
	if !strings.Contains(html, want) {
		t.Fatalf("expected rendered image, got:\n%s", html)
	}
}

func TestRenderMarkdownResolvesIndexRelativeImage(t *testing.T) {
	html := string(renderMarkdown("wiki/index.md",
		[]byte("![demo](assets/demo/sample.png)\n"), nil))
	if !strings.Contains(html, `<img src="/assets/demo/sample.png" alt="demo" loading="lazy">`) {
		t.Fatalf("expected index-relative image, got:\n%s", html)
	}
}

func TestRenderMarkdownDegradesUnsafeOrNonImage(t *testing.T) {
	// Escapes wiki/assets/: must not emit an <img>, falls back to alt text.
	escape := string(renderMarkdown("wiki/sources/x.md",
		[]byte("![secret](../../etc/passwd.png)\n"), nil))
	if strings.Contains(escape, "<img") {
		t.Fatalf("path traversal should not render an image, got:\n%s", escape)
	}
	if !strings.Contains(escape, "secret") {
		t.Fatalf("expected alt fallback text, got:\n%s", escape)
	}
	// Non-image extension also degrades.
	nonImage := string(renderMarkdown("wiki/sources/x.md",
		[]byte("![notes](../assets/x.txt)\n"), nil))
	if strings.Contains(nonImage, "<img") {
		t.Fatalf("non-image target should not render an image, got:\n%s", nonImage)
	}
}

func TestResolveImageSrc(t *testing.T) {
	cases := []struct {
		name   string
		page   string
		target string
		want   string
		wantOK bool
	}{
		{"relative-from-source", "wiki/sources/s.md", "../assets/g/a.png", "/assets/g/a.png", true},
		{"repo-rooted", "wiki/sources/s.md", "wiki/assets/g/a.png", "/assets/g/a.png", true},
		{"index-relative", "wiki/index.md", "assets/g/a.png", "/assets/g/a.png", true},
		{"https-passthrough", "wiki/sources/s.md", "https://example.com/a.png", "https://example.com/a.png", true},
		{"space-escaped", "wiki/sources/s.md", "../assets/g/My Slide.png", "/assets/g/My%20Slide.png", true},
		{"escape-rejected", "wiki/sources/s.md", "../../etc/passwd.png", "", false},
		{"outside-assets", "wiki/sources/s.md", "../sources/other.png", "", false},
		{"non-image", "wiki/sources/s.md", "../assets/g/a.txt", "", false},
		{"other-scheme", "wiki/sources/s.md", "file:///etc/passwd.png", "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := resolveImageSrc(tc.page, tc.target)
			if ok != tc.wantOK || got != tc.want {
				t.Fatalf("resolveImageSrc(%q, %q) = (%q, %v), want (%q, %v)",
					tc.page, tc.target, got, ok, tc.want, tc.wantOK)
			}
		})
	}
}

func TestCleanAssetPathRejectsTraversal(t *testing.T) {
	ok := []struct{ in, want string }{
		{"demo/sample.png", "wiki/assets/demo/sample.png"},
		{"c10s1-slides/Slide02.png", "wiki/assets/c10s1-slides/Slide02.png"},
	}
	for _, tc := range ok {
		got, valid := cleanAssetPath(tc.in)
		if !valid || got != tc.want {
			t.Fatalf("cleanAssetPath(%q) = (%q, %v), want (%q, true)", tc.in, got, valid, tc.want)
		}
	}
	bad := []string{"", "/etc/passwd", "../../etc/passwd", "..", "demo/../../../etc/passwd", "demo/./sample.png", "a//b.png"}
	for _, in := range bad {
		if got, valid := cleanAssetPath(in); valid {
			t.Fatalf("cleanAssetPath(%q) = (%q, true), want rejected", in, got)
		}
	}
}

func TestHandleAssetServesPNGWithMIME(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	assetDir := filepath.Join("wiki", "assets", "demo")
	if err := os.MkdirAll(assetDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	want := []byte("\x89PNG\r\n\x1a\n\x00demo-bytes")
	if err := os.WriteFile(filepath.Join(assetDir, "sample.png"), want, 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	srv := &Server{}
	rec := httptest.NewRecorder()
	srv.handleAsset(rec, httptest.NewRequest(http.MethodGet, "/assets/demo/sample.png", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "image/png" {
		t.Fatalf("content-type = %q, want image/png", ct)
	}
	if rec.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Fatalf("missing nosniff header")
	}
	if rec.Body.String() != string(want) {
		t.Fatalf("body mismatch")
	}
}

func TestBuildWikiGraphDerivesEdgesFromMarkdownLinks(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	mustWrite := func(rel, content string) {
		p := filepath.Join(dir, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatalf("write: %v", err)
		}
	}
	mustWrite("wiki/entities/sqlite.md", "---\ntitle: SQLite\nkind: entity\n---\n\n# SQLite\n\nRecommended for [data preservation](../topics/data-preservation-formats.md).\n")
	mustWrite("wiki/topics/data-preservation-formats.md", "---\ntitle: Data Preservation Formats\nkind: topic\n---\n\n# Data Preservation Formats\n")

	found := false
	for _, e := range wikiGraphEdges(0) {
		if e.FromPath == "wiki/entities/sqlite.md" && e.ToPath == "wiki/topics/data-preservation-formats.md" {
			found = true
			if !e.Known {
				t.Fatalf("edge to existing page should be Known: %#v", e)
			}
		}
	}
	if !found {
		t.Fatalf("expected Markdown-link edge sqlite -> data-preservation-formats from graph edges")
	}
}

func TestHandleAssetBlocksTraversalAndMissing(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	// A secret file outside wiki/assets/ that must never be reachable.
	if err := os.WriteFile(filepath.Join(dir, "secret.txt"), []byte("top secret"), 0o644); err != nil {
		t.Fatalf("write secret: %v", err)
	}
	for _, p := range []string{
		"/assets/../../secret.txt",
		"/assets/../secret.txt",
		"/assets/",
		"/assets/missing/none.png",
	} {
		rec := httptest.NewRecorder()
		srv := &Server{}
		srv.handleAsset(rec, httptest.NewRequest(http.MethodGet, p, nil))
		if rec.Code != http.StatusNotFound {
			t.Fatalf("GET %s = %d, want 404 (body=%q)", p, rec.Code, rec.Body.String())
		}
	}
}
