// Package cliroot resolves the $WIKI_ROOT directory that nemocli operates on.
//
// nemocli accepts the wiki root from three layers (highest priority first):
//  1. --wiki-root <path> CLI flag
//  2. NEMO_WIKI_ROOT environment variable
//  3. built-in default ~/.wiki
//
// Resolution is best-effort: ResolveWikiRoot expands a leading ~ and returns
// an absolute path, but does not require the directory to exist on disk. The
// init subcommand is the only entry point that should create it; other
// subcommands call RequireInitialized to refuse to operate on an uninitialized
// root.
//
// The "wiki root" is a workspace directory whose wiki/ subdirectory holds the
// content that internal/wikilint, internal/wikimaint, internal/apply, and
// internal/query consume. Operational siblings such as .env and AGENTS.md
// live alongside that wiki/ subdirectory, never inside it.
package cliroot

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// EnvVar names the environment variable that overrides the default wiki root.
const EnvVar = "NEMO_WIKI_ROOT"

// DefaultRelative is the home-relative fallback when neither flag nor env is set.
const DefaultRelative = ".wiki"

// ResolveWikiRoot picks the effective wiki root from flag > env > default.
//
// The returned path is always absolute and has a leading ~ expanded.
// The directory itself is not required to exist; subcommands other than
// init should call RequireInitialized to verify the layout.
func ResolveWikiRoot(flagValue string) (string, error) {
	raw := strings.TrimSpace(flagValue)
	if raw == "" {
		raw = strings.TrimSpace(os.Getenv(EnvVar))
	}
	if raw == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve default wiki root: %w", err)
		}
		raw = filepath.Join(home, DefaultRelative)
	}
	return expandAndAbs(raw)
}

func expandAndAbs(raw string) (string, error) {
	if raw == "~" || strings.HasPrefix(raw, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("expand ~ in %q: %w", raw, err)
		}
		if raw == "~" {
			raw = home
		} else {
			raw = filepath.Join(home, raw[2:])
		}
	}
	abs, err := filepath.Abs(raw)
	if err != nil {
		return "", fmt.Errorf("absolute path for %q: %w", raw, err)
	}
	return abs, nil
}

// RequireInitialized returns nil iff wikiRoot looks like it has been
// scaffolded by `nemocli init`. The check is intentionally cheap: it
// looks for wiki/index.md or wiki/log.md as sentinel markers.
func RequireInitialized(wikiRoot string) error {
	markers := []string{
		filepath.Join(wikiRoot, "wiki", "index.md"),
		filepath.Join(wikiRoot, "wiki", "log.md"),
	}
	for _, m := range markers {
		if _, err := os.Stat(m); err == nil {
			return nil
		}
	}
	return fmt.Errorf(
		"wiki root %s does not look initialized (no wiki/index.md);\n"+
			"  run `nemocli init --wiki-root %s` first",
		wikiRoot, wikiRoot,
	)
}

// WikiContentDir returns the path of the content directory that
// internal/wikilint and friends walk, relative to wikiRoot.
//
// The wiki content lives at <wikiRoot>/wiki/, matching the layout that
// internal/* expects (filepath.Join(root, "wiki")). Keeping the actual
// pages in a fixed-name subdirectory lets the wiki root also hold
// operational siblings such as .env and AGENTS.md.
func WikiContentDir(wikiRoot string) string {
	return filepath.Join(wikiRoot, "wiki")
}
