---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Chunk Context

This chunk covers the `go version` command's module reporting capabilities, followed by a detailed guide on version queries (semver, pseudo-versions, revision identifiers), behavior of module commands outside a main module context, and an introduction to workspace management via `go work`. It concludes with a comprehensive specification for module proxies, including protocol requirements, error handling, caching strategies, and the specific HTTP endpoints a proxy must implement.

# Local Summary

The section details how to use `go version -m` to print embedded module information for executables, explaining the table format (path, mod, dep, =>). It defines various version query syntaxes (`@latest`, `@upgrade`, `@patch`, etc.) and their selection logic. The text outlines which commands require a `go.mod` file versus those that operate in "module-aware mode" without one. Finally, it introduces the workspace mechanism (`go work init/edit/sync`) for managing multi-module projects and defines the strict contract module proxies must adhere to regarding version lists, metadata, source files, and zip archives.

# Key Claims

- `go version -m` prints a tab-separated table of an executable's embedded module information if available.
- Version queries support semantic versions, prefixes, comparisons, revision identifiers (commit hashes/branches), and keywords like `latest`, `upgrade`, and `patch`.
- The `latest` keyword prefers release versions over pre-release versions; pseudo-versions are selected only if no tagged versions exist.
- Commands like `go build`, `go test`, and `go run` behave differently when a `go.mod` file is absent (packages from other modules cannot be built).
- Module proxies must respond to specific paths (`@v/list`, `@v/$version.info`, `@v/$version.mod`, `@v/$version.zip`) with specific content types.
- The `go work sync` command uses the Minimal Version Selection (MVS) algorithm to upgrade workspace module dependencies to match a consistent build list.

# Entities And Concepts

- **Command**: `go version -m` (prints executable module info), `go clean -modcache`, `go get @version`, `go mod download @version`, `go work init`, `go work edit`, `go work sync`.
- **Version Query Types**: Semantic version, semantic version prefix, semantic version comparison, revision identifier, `latest`, `upgrade`, `patch`.
- **Module Proxy Endpoints**: `$base/$module/@v/list` (list versions), `$base/$module/@v/$version.info` (metadata JSON), `$base/$module/@v/$version.mod` (go.mod file), `$base/$module/@v/$version.zip` (source archive).
- **Workspace Tooling**: `go.work`, `go work init`, `go work edit`, `go work sync`, `go work use`.
- **Environment Variables**: `GOPROXY`, `GOFLAGS`.

# Procedures And API Details

**1. Printing Module Versions:**
Run `$ go version -m <path/to/bin>`. If no files are named, it prints its own version. If a directory is given, it walks recursively for binaries. Use `-v` to report unrecognized files. Use `-m` to show the table format:
- `path`: Main package path.
- `mod`: Module containing main (path, version, sum).
- `dep`: Linked module (path, version, sum).
- `=>`: Replacement for a module (local dir or module path/version/sum).

**2. Version Queries:**
Append `@query` to a module path.
- `@latest`: Highest available release; if none, highest pre-release; if none, pseudo-version of default branch tip.
- `@upgrade`: Like `latest` but preserves current higher version if applicable.
- `@patch`: Latest version with same major.minor as currently required.
- `@revision`: Commit hash, tag, or branch name.

**3. Module Commands Without `go.mod`:**
Commands like `go build`, `go test`, `go run` load only standard library and `.go` files provided on the command line. They cannot import packages from other modules because there is no place to record requirements. `go get` works without a main module but ignores `replace`/`exclude` directives.

**4. Workspace Management:**
- `go work init [moddirs]`: Creates `go.work`. Arguments are added as `use` directives.
- `go work edit [-fmt] [-use=path] [-dropuse=path] [-replace=old[@v]=new[@v]] ...`: Edits `go.work` via CLI flags. Supports `-print` and `-json` output modes.
- `go work sync`: Runs MVS to compute a consistent build list for all workspace modules and updates their `go.mod` files accordingly.

**5. Module Proxy Protocol:**
A proxy serves at `$base/$module/@v/$version.*`.
- **List**: `$base/$module/@v/list` returns plain text versions (no pseudo-versions).
- **Info**: `$base/$module/@v/$version.info` returns JSON `{Version: string, Time: time.Time}`.
- **Mod**: `$base/$module/@v/$version.mod` returns the `go.mod` file for that version.
- **Zip**: `$base/$module/@v/$version.zip` returns the zip archive.
- **Fallback**: Use comma (`,`) to fallback after 404/410; use pipe (`|`) to fallback after any error.

# Nuance Or Contradictions

- **Version Selection Priority**: `latest` prefers release versions over pre-releases. Even if a pre-release is higher than the highest release, `latest` ignores it.
- **Pseudo-versions in Proxies**: Pseudo-versions are excluded from `$base/$module/@v/list`.
- **Case Encoding**: Module paths and versions in proxy requests are case-encoded (uppercase replaced by `!lower`) to handle case-insensitive file systems (e.g., `example.com/M` becomes `example.com/!m`).
- **Workspace Sync Behavior**: `go work sync` guarantees the workspace build list version is always the same or higher than what is currently in each module's `go.mod`.

# Candidate Wiki Hints

- Page: "Module Version Queries" (covers syntax for `@latest`, `@patch`, etc.)
- Page: "Workspace Management with go work" (explains `go work init`, `sync`, and MVS)
- Page: "Module Proxy Protocol Specification" (defines endpoints and response formats)
