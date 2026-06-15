---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Chunk Context

**Chunk:** 7 of 13
**Lines:** 1686-2071
**Heading path:** Upgrade a specific module.

The chunk covers the `go get` command's ability to upgrade, downgrade, or remove modules and their transitive dependencies. It details version query suffixes (e.g., `@v1.0.0`, `@master`, `@none`) and explains how `go get` updates the main module's `go.mod` file while managing dependency chains. The text transitions into a comparison with `go install` (for installing executables ignoring local `go.mod`) and concludes with usage details for `go mod edit`, `go mod download`, and `go list -m`.

# Local Summary

This section documents the `go get` command's behavior when managing module dependencies, including specific version targeting, transitive upgrades/downgrades, and removal of dependencies. It contrasts `go get` (dependency management) with `go install` (executable installation). The chunk also introduces `go mod edit` for manual file manipulation (adding replacements, formatting) and `go mod download` for pre-filling the cache.

# Key Claims

- **Dependency Management:** `go get` updates the main module's `go.mod` file to reflect new dependencies or version changes, automatically upgrading/downgrading transitive requirements as needed.
- **Version Queries:** Suffixes like `@v0.3.2`, `@master`, `@latest`, `@upgrade`, `@patch`, and `@none` allow precise control over module versions; `@none` removes a requirement entirely.
- **Transitive Changes:** Upgrading a module may upgrade its dependencies if the new version requires them. Downgrading a module may downgrade dependencies to satisfy compatibility.
- **Deprecation/Retraction:** `go get` checks for retracted or deprecated modules and prints warnings; `go list -m -u` can inspect all dependencies for these states.
- **Executable Installation:** Since Go 1.16, `go install` is preferred for installing programs. It builds in module-aware mode by default if version suffixes are provided, ignoring the current directory's `go.mod`.
- **Tool Directives:** `go tool` can build and run tools declared in a module's `go.mod` via a `tool` directive.
- **Module Editing:** `go mod edit` allows adding replacements (`-replace`), dropping them (`-dropreplace`), setting Go versions, and formatting files without saving (`-print`).

# Entities And Concepts

- **Command:** `go get`, `go install`, `go tool`, `go list -m`, `go mod download`, `go mod edit`.
- **File:** `go.mod`, `go.sum`.
- **Environment Variables:** `GOBIN`, `GOPATH`, `GOROOT`, `GOTOOLDIR`, `GOINSECURE`, `GO111MODULE`.
- **Concepts:** Module-aware mode, GOPATH mode, transitive dependencies, version queries (semantic versioning), retracted versions, deprecated modules, minimal version selection (MVS).

# Procedures And API Details

### Using `go get`
- **Syntax:** `go get [flags] module/path[@version]`
- **Examples:**
  - Upgrade: `$ go get golang.org/x/net`
  - Specific version: `$ go get golang.org/x/text@v0.3.2`
  - Master branch: `$ go get golang.org/x/text@master`
  - Remove dependency: `$ go get golang.org/x/text@none`
  - Upgrade minimum Go version: `$ go get go`
- **Flags:**
  - `-u`: Upgrade imported packages' modules to latest.
  - `-u=patch`: Upgrade to latest patch version only.
  - `-t`: Include test dependencies.
  - `-d`: Do not build/install (deprecated in Go 1.18).

### Using `go install`
- **Syntax:** `go install [build flags] [packages]`
- **Behavior:** Installs executables to `$GOBIN`. Ignores local `go.mod` if version suffixes are used.
- **Constraints:** Arguments must be package paths/patterns, not standard packages or file paths. All arguments must share the same version suffix.

### Using `go mod edit`
- **Syntax:** `go mod edit [editing flags] [-fmt|-print] [go.mod]`
- **Examples:**
  - Add replace: `$ go mod edit -replace example.com/a@v1.0.0=./a`
  - Drop replace: `$ go mod edit -dropreplace example.com/a@v1.0.0`
  - Set Go version & print: `$ go mod edit -go=1.14 -require=example.com/m@v1.0.0 -print`
  - Format: `$ go mod edit -fmt`

### Using `go list -m`
- **Syntax:** `go list -m [-u] [-retracted] [-versions] [modules]`
- **Output:** Lists modules with version info, available upgrades (`-u`), retracted versions (`-retracted`), or all known versions (`-versions`).

### Using `go mod download`
- **Syntax:** `go mod download [-x] [-json] [-reuse=old.json] [modules]`
- **Purpose:** Downloads modules into the cache; useful for pre-filling caches or module proxies.

# Nuance Or Contradictions

- **`go get` vs. `go install`:** `go get` manages dependencies in `go.mod`, while `go install` focuses on installing executables and ignores local `go.mod` when version suffixes are present (since Go 1.16). The `-d` flag for `go get` is deprecated as of Go 1.17 and always enabled in Go 1.18.
- **Version Query Ambiguity:** All arguments to `go install` must have the same version suffix; mixing queries (e.g., `@latest` and `@v1.0.0`) causes errors.
- **Module Context:** In module-aware mode, `go install` runs in the context of the main module, which may differ from the module containing the package being installed.
- **Retracted Versions:** By default, retracted versions are omitted from `go list -m -versions` unless `-retracted` is specified.

# Candidate Wiki Hints

- **Page:** Managing Module Dependencies with `go get`
  - Covers version queries (`@v`, `@master`, `@none`), transitive dependency updates, and deprecation warnings.

- **Page:** Installing Executables with `go install`
  - Explains module-aware mode, ignoring local `go.mod`, and constraints on arguments (same version suffix, main packages only).

- **Page:** Editing `go.mod` Manually with `go mod edit`
  - Details replace directives, dropping replacements, setting Go versions, and formatting files.

- **Page:** Inspecting Modules with `go list -m`
  - Describes output formats for upgrades, retracted versions, and available version lists.
