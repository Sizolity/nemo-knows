package cliroot

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResolveWikiRoot_FlagWins(t *testing.T) {
	t.Setenv(EnvVar, "/should/be/ignored")
	t.Setenv("HOME", t.TempDir())
	got, err := ResolveWikiRoot("/tmp/explicit-path")
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if got != "/tmp/explicit-path" {
		t.Fatalf("flag should win, got %q", got)
	}
}

func TestResolveWikiRoot_EnvFallback(t *testing.T) {
	dir := t.TempDir()
	envVal := filepath.Join(dir, "envroot")
	t.Setenv(EnvVar, envVal)
	t.Setenv("HOME", dir)
	got, err := ResolveWikiRoot("")
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if got != envVal {
		t.Fatalf("env should win when flag empty, got %q", got)
	}
}

func TestResolveWikiRoot_DefaultFallback(t *testing.T) {
	home := t.TempDir()
	t.Setenv(EnvVar, "")
	t.Setenv("HOME", home)
	got, err := ResolveWikiRoot("")
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	want := filepath.Join(home, DefaultRelative)
	if got != want {
		t.Fatalf("default should be %q, got %q", want, got)
	}
}

func TestResolveWikiRoot_ExpandsTilde(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv(EnvVar, "")
	got, err := ResolveWikiRoot("~/some/sub")
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	want := filepath.Join(home, "some", "sub")
	if got != want {
		t.Fatalf("expand ~, want %q got %q", want, got)
	}
}

func TestResolveWikiRoot_RelativeBecomesAbsolute(t *testing.T) {
	t.Setenv(EnvVar, "")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	got, err := ResolveWikiRoot("./relative")
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if !filepath.IsAbs(got) {
		t.Fatalf("want absolute path, got %q", got)
	}
	if !strings.HasPrefix(got, cwd) {
		t.Fatalf("relative path should resolve under cwd %q, got %q", cwd, got)
	}
}

func TestRequireInitialized_Missing(t *testing.T) {
	root := t.TempDir()
	if err := RequireInitialized(root); err == nil {
		t.Fatal("expected error for empty root, got nil")
	}
}

func TestRequireInitialized_HasIndex(t *testing.T) {
	root := t.TempDir()
	wikiDir := filepath.Join(root, "wiki")
	if err := os.MkdirAll(wikiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(wikiDir, "index.md"), []byte("---\n---\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := RequireInitialized(root); err != nil {
		t.Fatalf("expected ok, got %v", err)
	}
}

func TestWikiContentDir(t *testing.T) {
	got := WikiContentDir("/some/root")
	want := filepath.Join("/some/root", "wiki")
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}
