---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Group Context

**Source:** `raw/web/corpus-2026-05-18/032-go-modules-reference.md`
**Line Range:** 1–4266
**Primary Heading Path:** Document > Go Modules Reference
**Scope:** This group covers the comprehensive reference for Go Modules, spanning from initial metadata retrieval and fundamental concepts (modules, paths, versioning) to detailed `go.mod` file structure, directives, upgrade procedures, workspace management (`go.work`), vendoring behavior, and diagnostic commands for inspecting builds.

# Cross-Chunk Summary

The document systematically details the Go module system's lifecycle and tooling. It begins by establishing the workflow of fetching metadata and resolving dependencies via proxies (`GOPROXY`). It then defines the core unit of a module: a collection of versioned packages identified by a `go.mod` file in the root directory. The reference elaborates on path construction, semantic versioning rules, and pseudo-versions for commit-specific references.

A significant portion is dedicated to the `go.mod` file itself, covering its lexical grammar, mandatory directives (e.g., `module`, `go`, `toolchain`), and optional ones (e.g., `require`, `replace`, `retract`, `exclude`). The document explains how the build tool uses these directives to construct a dependency graph, utilizing algorithms like Minimal Version Selection (MVS) to determine the exact set of versions needed.

The text addresses legacy compatibility, detailing how Go handles pre-module repositories and GOPATH mode alongside modern modules (including minimal module compatibility rules). It further explores workspace management via `go.work` files, which define toolchain versions and include multiple modules. Finally, it covers build-time behaviors such as vendoring, where local copies of dependencies are used instead of the network, and diagnostic commands for inspecting the final build state and executable metadata.

# Repeated Or Central Claims

- **Module Definition:** A module is a collection of versioned packages released together, identified by a path declared in a `go.mod` file located in the module's root directory.
- **Dependency Resolution:** The `go` command resolves package paths by searching the build list for matching module prefixes, consulting proxies (defined via `GOPROXY`), and preferring the longest matching path.
- **Versioning Strategy:** Semantic versioning (`vX.Y.Z`) is standard; pseudo-versions encode specific revisions for testing. Starting with major version 2, a suffix (e.g., `/v2`) is required to distinguish incompatible packages unless using legacy paths.
- **Directives in `go.mod`:** The file uses specific directives: `module` (defines root), `go` (sets minimum Go version, mandatory since 1.21), `toolchain`, `godebug`, `require` (direct deps), `replace` (substitution), `exclude`/`retract` (lifecycle management), and `ignore`.
- **Indirect Dependencies:** Since Go 1.17, indirect dependencies are recorded in a separate block to enable module graph pruning and lazy loading.
- **Minimal Version Selection (MVS):** An algorithm that traverses the module graph to compute the minimal set of versions required for a build, ensuring deterministic selection.
- **Vendoring Behavior:** `go build` and `go test` use vendored packages when enabled; other commands (`go mod tidy`, `go get`) continue to interact with the network/cache regardless of vendoring status. Only vendor directories at the main module's root are respected.

# Important Local Details

- **Directives Syntax & Behavior:**
  - `replace <module> [<version>] => <path|module> [version]`: Substitutes a specific version with another path or version. Without a left-side version, all versions are replaced.
  - `retract <version>`: Marks a version as problematic; prevents upgrades but keeps it accessible for existing builds. Ranges can be specified (e.g., `[v1.0.0, v1.9.9]`).
  - `exclude`: Prevents specific module versions from being loaded; behavior changed in Go 1.16 to ensure deterministic selection.
  - `tool` (Go 1.24+): Adds packages as module-level tool dependencies.
- **Version Constraints:** Non-canonical versions (without leading `v`) are allowed only in the main module's `go.mod`. They are automatically replaced with canonical versions when possible.
- **Workspace Files (`go.work`):** Must contain exactly one `go` directive specifying the toolchain version. The `use` directive adds a module directory but does not recursively include submodules.
- **Legacy Compatibility:** Repositories lacking `go.mod` but having major version 2+ tags may have synthetic `go.mod` files created, potentially appending a `+incompatible` suffix. Minimal module compatibility allows imports from GOPATH directories to resolve against modules with major version suffixes under specific conditions.
- **Build Diagnostics:**
  - `go list -m all`: Displays the selected versions determined by MVS.
  - `go get -mod=mod`: Automatically rewrites non-canonical versions, respects exclusions, removes redundant requirements, and reformats `go.mod`.
  - `go work init` / `go work use`: Commands to initialize and populate workspace files.

# Candidate Wiki Hints

- **Page: `go-modules-reference`** – The overarching reference document covering the module system.
- **Page: `go-mod-directives`** – Detailed syntax and version-specific behavior for `module`, `go`, `require`, `replace`, `retract`, `exclude`, etc.
- **Page: `go-mod-versioning`** – Canonical vs. non-canonical versions, path restrictions, and toolchain integration.
- **Page: `indirect-dependencies-in-go-mod`** – Explanation of the separate block for indirect dependencies introduced in Go 1.17.
- **Page: `go-workspace-files`** – Overview of `go.work` syntax, directives (`use`, `replace`), and toolchain management.
- **Page: `go-modules-vendoring-behavior`** – How vendoring affects build commands versus module management tools, including scope restrictions and known bugs regarding zip files.
- **Page: `minimal-version-selection`** – The MVS algorithm for dependency resolution.
- **Page: `legacy-module-compatibility`** – Handling pre-module repositories and GOPATH mode alongside modern modules.

# Gaps Or Cautions

- **Committing `go.work`:** Generally inadvisable as it can override parent workspace definitions or cause CI systems to test incorrect versions, unless the repository is developed exclusively together.
- **Zip File Limitations:** Vendor directories from sub-modules are excluded from module zip files (referencing bugs #31562 and #37397).
- **Retraction Visibility:** Retracted versions are hidden from `go list -m -versions` unless the `-retracted` flag is used.
- **Replace vs Require:** A `replace` directive alone does not add a module to the dependency graph; a corresponding `require` directive is still necessary.
- **Tag Naming:** The `+incompatible` suffix appears in Go command usage but should not appear on repository tags (e.g., `v4.1.2+incompatible` is ignored by the tag system).
