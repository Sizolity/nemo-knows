package main

import (
	"os"
	"path/filepath"
	"testing"
)

// TestRunInitThenLint exercises the dispatcher end-to-end for the two
// subcommands that do not touch the model: init creates the scaffold
// under a temp wiki root, lint then runs cleanly against it. Coverage
// here is intentionally narrow — it verifies argv plumbing and the
// init→lint contract without spinning up any LLM.
func TestRunInitThenLint(t *testing.T) {
	tmp := t.TempDir()
	root := filepath.Join(tmp, "wikiroot")

	// `go test` already provides a non-tty os.Stdin so init's tty
	// detection naturally falls into the non-interactive branch.

	// Redirect runs-log writes to the temp dir so the test does not
	// touch the developer's real ~/.local/state/nemo.
	t.Setenv("XDG_STATE_HOME", filepath.Join(tmp, "state"))
	t.Setenv("HOME", tmp)
	t.Setenv("NEMO_WIKI_ROOT", root)

	if got := run([]string{"--wiki-root", root, "init"}); got != 0 {
		t.Fatalf("init exit = %d, want 0", got)
	}

	// init should have created the wiki layout.
	for _, sub := range []string{"sources", "entities", "concepts", "topics", "assets"} {
		if _, err := os.Stat(filepath.Join(root, "wiki", sub)); err != nil {
			t.Errorf("init did not create wiki/%s: %v", sub, err)
		}
	}
	for _, name := range []string{"index.md", "log.md", "AGENTS.md"} {
		if _, err := os.Stat(filepath.Join(root, "wiki", name)); err != nil {
			t.Errorf("init did not create wiki/%s: %v", name, err)
		}
	}
	if _, err := os.Stat(filepath.Join(root, ".env.example")); err != nil {
		t.Errorf("init did not create .env.example: %v", err)
	}

	// lint should pass cleanly on the freshly-scaffolded root.
	if got := run([]string{"--wiki-root", root, "lint"}); got != 0 {
		t.Fatalf("lint exit = %d, want 0 on a freshly initialized wiki", got)
	}
}

func TestRunHelpAndStubs(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("XDG_STATE_HOME", t.TempDir())
	cases := []struct {
		name string
		args []string
		want int
	}{
		{"no-args", []string{}, 0},
		{"-h", []string{"-h"}, 0},
		{"--help", []string{"--help"}, 0},
		{"version", []string{"version"}, 0},
		{"query-stub", []string{"query", "what"}, 0},
		{"serve-stub", []string{"serve"}, 0},
		{"unknown", []string{"frobnicate"}, 2},
		{"missing-sub", []string{"--profile", "stable"}, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := run(c.args); got != c.want {
				t.Fatalf("run(%v) = %d, want %d", c.args, got, c.want)
			}
		})
	}
}

func TestRequireInitializedBlocksLintBeforeInit(t *testing.T) {
	tmp := t.TempDir()
	root := filepath.Join(tmp, "empty")
	t.Setenv("HOME", tmp)
	t.Setenv("XDG_STATE_HOME", filepath.Join(tmp, "state"))
	if got := run([]string{"--wiki-root", root, "lint"}); got == 0 {
		t.Fatalf("lint should refuse to run before init, got exit 0")
	}
}
