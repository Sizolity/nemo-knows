## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## Chunk Context
- **Source Path:** `raw/web/corpus-2026-05-18/032-go-modules-reference.md`
- **Line Range:** 1–26
- **Document Title:** Go Modules Reference
- **Origin URL:** https://go.dev/ref/mod
- **Retrieval Date:** 2026-05-18

## Local Summary
This initial chunk introduces the "Go Modules Reference" document, which serves as a source for dependency management concepts in the Go programming language. It includes metadata indicating the fetch status is okay and categorizes the content under "Go" within a web corpus.

## Key Claims
- The document is titled "Go Modules Reference".
- The reference covers dependency-management concepts.
- The content was successfully fetched from `https://go.dev/ref/mod`.

## Entities And Concepts
- **Go:** Programming language.
- **Modules:** Dependency management system for Go.
- **Corpus Item 32:** Specific identifier within the curated web corpus.

## Procedures And API Details
- No specific procedures or API calls are detailed in this chunk; it serves as an introduction and metadata header.

## Nuance Or Contradictions
- None observed in this introductory segment. The text confirms a successful retrieval with a standard content type (`text/html`).

## Candidate Wiki Hints
- **Page Suggestion:** Create a page for "Go Modules Reference" to document the source material regarding Go's dependency management system.

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
# Chunk Context

This chunk covers the introduction to Go modules, defining them as collections of versioned packages. It details module paths, semantic versions (including pre-release and build metadata), pseudo-versions encoding revision history, major version suffixes for v2+ compatibility, and the resolution process where `go` searches the build list or GOPROXY to find the providing module.

# Local Summary

The text introduces Go modules as the dependency management system, explaining that a module is a collection of packages released and versioned together. It outlines the structure of module paths (repository root, subdirectory, major version suffix), semantic versioning rules, and the mechanics of pseudo-versions used for unreleased commits. The chunk further explains the introduction of major version suffixes starting with v2 to handle incompatibilities and resolves how `go` identifies which module provides a specific package path by checking prefixes in the build list or proxy servers.

# Key Claims

- Modules are collections of packages released, versioned, and distributed together.
- A module is identified by a module path declared in a `go.mod` file.
- Module paths should describe functionality and location, typically including a repository root, optional subdirectory, and major version suffix (v2+).
- Semantic versions consist of major.minor.patch; major increments require incompatible changes, minor for compatible additions, patch for internal fixes.
- Pseudo-versions encode specific revision identifiers (commit hashes) and timestamps to ensure canonical ordering without manual typing.
- Starting with v2, module paths must include a major version suffix (e.g., `/v2`) to maintain import compatibility rules between incompatible versions.
- The `go` command resolves packages by searching modules in the build list for matching path prefixes; if none are found locally, it queries GOPROXY entries.

# Entities And Concepts

- **Module**: A collection of packages released and versioned together.
- **Module Path**: The canonical name declared in `go.mod`, acting as a prefix for package paths within the module.
- **Semantic Versioning**: Format `vX.Y.Z` with optional `-pre-release` and `+build-metadata`.
- **Pseudo-version**: A pre-release version encoding a specific revision identifier (e.g., `v0.0.0-20191109...`).
- **Major Version Suffix**: Required for v2+ to distinguish incompatible module paths (e.g., `/v2`).
- **GOPROXY**: Environment variable controlling the list of proxy URLs or keywords (`direct`, `off`) for downloading modules.
- **Build List**: The set of modules currently available to the build process, checked first during resolution.

# Procedures And API Details

- **Resolving a Package**:
  1. `go` searches the build list for modules whose paths are prefixes of the package path.
  2. If exactly one module provides the package, it is used.
  3. If none or multiple match, an error is reported unless `-mod=mod` is used to fetch new modules.
- **GOPROXY Requests**: For each entry in `GOPROXY`, `go` requests the latest version of each potential module path prefix (e.g., for package `golang.org/x/net/html`, it requests `golang.org/x/net/html`, `golang.org/x/net`, etc.).
- **Version Conversion**: Commands like `go get` or `go list -m` can accept branch names or commit hashes, automatically translating them into pseudo-versions or tagged versions.

# Nuance Or Contradictions

- **v0/v1 vs v2 Suffixes**: Major version suffixes are not allowed at v0 or v1 because v0 is unstable and v1 implies compatibility with the previous v0 release.
- **gopkg.in Exception**: Modules starting with `gopkg.in/` must always have a major version suffix, even at v0/v1, but using a dot separator (e.g., `gopkg.in/yaml.v2`) instead of a slash.
- **Pseudo-version Ordering**: Pseudo-versions sort higher than their base version but lower than the next tagged version, ensuring they sit correctly in version ordering without manual intervention.

# Candidate Wiki Hints

- Create a page on **Go Modules Basics** explaining what modules are and how `go.mod` identifies them.
- Draft a guide on **Semantic Versioning in Go**, detailing major/minor/patch rules and the meaning of `-pre` and `+incompatible`.
- Write an article on **Pseudo-versions**, explaining how they encode commit hashes and timestamps for unreleased code.
- Develop a section on **Major Version Suffixes**, covering when `/v2` is required and the `gopkg.in` special case.
- Create a troubleshooting page for **Module Resolution Errors**, discussing build list checks and GOPROXY configuration.

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

Chunk Context
This chunk details the syntax and behavior of `go.mod` files, including directives (`module`, `go`, `require`, `replace`, etc.), lexical elements (tokens, strings), version constraints, deprecation handling, and specific behaviors introduced in Go 1.16, 1.17, and later regarding indirect dependencies and toolchains.

Local Summary
The `go.mod` file defines a module using UTF-8 text with line-oriented directives. It supports keywords like `module`, `go`, `require`, `replace`, `exclude`, `retract`, `tool`, `ignore`, and `godebug`. The grammar allows block syntax for grouping similar directives (e.g., multiple `require` statements). Lexical analysis distinguishes between interpreted strings (with escapes) and raw strings. Module paths must follow specific ASCII rules, and versions can be canonical (`v1.2.3`) or non-canonical (restricted to the main module). The chunk also explains how `go` directives affect toolchain selection and language feature usage, and how indirect dependencies are managed starting from Go 1.17.

Key Claims
- A `go.mod` file is required for the main module and any local replacement modules.
- Indirect dependencies are marked with a `// indirect` comment; automatic addition of these comments was refined in Go 1.16 (for explicit upgrades/removals) and Go 1.17 (comprehensive tracking).
- Since Go 1.21, the `go` directive declares a mandatory minimum Go version; toolchains refuse to use modules requiring newer versions.
- Deprecation comments (`// Deprecated:`) apply to all minor versions of a module but not major versions higher than v2.
- The `tool` directive (Go 1.24+) adds packages as dependencies runnable via `go tool`.
- `ignore` directives exclude directory paths from package pattern matching.
- `exclude` directives prevent loading specific module versions in the main module since Go 1.16, avoiding non-deterministic version selection seen in earlier versions.

Entities And Concepts
- **Directives**: `module`, `go`, `require`, `replace`, `exclude`, `retract`, `tool`, `ignore`, `godebug`.
- **Tokens**: Whitespace, comments (`//`), punctuation, keywords, identifiers, strings (interpreted vs. raw).
- **Module Path Rules**: ASCII letters/digits, limited punctuation; cannot start/end with slash or dot; Windows reserved names avoided.
- **Version Types**: Canonical (`vN.M.P`) vs. non-canonical (allowed only in main module).
- **Indirect Dependencies**: Automatically added comments indicating no direct import from the main module.
- **Toolchain Directives**: Suggest specific Go toolchains for reproducible builds.

Procedures And API Details
- **Adding Requirements**: Use `go get` to upgrade/downgrade dependencies; `go mod edit` for low-level edits; `golang.org/x/mod/modfile` for programmatic changes.
- **Syntax Examples**:
  - Block syntax: `require ( example.com/new/thing/v2 v2.3.4 )`
  - Deprecation comment placement: Immediately before or after the `module` directive.
  - Toolchain declaration: `toolchain go1.21.0`.
- **Grammar EBNF**: `GoMod = { Directive } .`, `Directive` variants for each keyword type.

Nuance Or Contradictions
- **Non-canonical versions**: Allowed only in the main module; the `go` command attempts to replace them with canonical versions during updates.
- **Indirect dependencies behavior change**: Before Go 1.16, indirect requirements were added only under specific upgrade/removal scenarios. From Go 1.17, they are added for all modules providing transitively imported packages.
- **Deprecation scope**: Applies to minor versions; major versions >v2 are treated as separate modules. Use `retract` for specific version removals rather than deprecation.
- **Tool directive requirement**: If the tool package is external, a corresponding `require` directive must exist.

Candidate Wiki Hints
- Create page: **Go Modules Reference** (comprehensive syntax and command guide).
- Create page: **go.mod Directives** (summary of available directives and their purposes).
- Create page: **Indirect Dependencies in Go** (explanation of `// indirect` comments and MVS interaction).
- Create page: **Go Version Compatibility** (role of the `go` directive and toolchain management).

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
Lines 861–1255 of `raw/web/corpus-2026-05-18/032-go-modules-reference.md`, covering the "Go Modules Reference" section. Topics include `replace` directives, `retract` directives, automatic updates via `-mod=mod`, Minimal Version Selection (MVS), module graph modifications (replacement, exclusion, upgrades/downgrades), module graph pruning for Go 1.17+, and workspaces.

Local Summary
This chunk details how Go manages dependency resolution through the module graph. It explains syntax and behavior for replacing modules locally or remotely, retracting versions to prevent usage, and using `-mod=mod` to auto-correct `go.mod`. It describes Minimal Version Selection (MVS), which deterministically computes a build list from the main module's requirements, traversing the graph while respecting exclusions, replacements, and upgrades. Later sections cover optimizations like module graph pruning in Go 1.17+ and lazy loading, followed by workspace definitions for managing multiple main modules.

Key Claims
- A `replace` directive substitutes a specific module version or all versions with a local file path or another remote module path.
- A `retract` directive marks versions as unusable for automatic upgrades while keeping them accessible in repositories.
- The `-mod=mod` flag instructs the Go command to automatically rewrite non-canonical versions, respect exclusions, remove redundant requirements, and reformat `go.mod`.
- Minimal Version Selection (MVS) is deterministic and recalculated at every module-aware command execution.
- Module graph pruning in Go 1.17+ limits the loaded graph to immediate dependencies for high-GO-version modules unless transitively required by lower versions.
- Workspaces (`go.work`) allow running MVS across multiple main modules defined by relative paths in a `go.work` file.

Entities And Concepts
- `replace` directive: Swaps a module version/path with another.
- `retract` directive: Marks versions as deprecated/unusable for upgrades.
- `-mod=mod`: Flag enabling automatic correction of `go.mod`.
- Minimal Version Selection (MVS): Algorithm to compute the build list.
- Module graph: Directed graph of module versions and dependencies.
- Exclusion (`exclude`): Removes specific versions from the graph.
- Upgrade/Downgrade: Modifying the graph to prefer higher/lower versions.
- Module graph pruning: Optimization in Go 1.17+ to limit loaded graph size.
- Lazy module loading: On-demand loading of the full graph only when necessary.
- Workspace (`go.work`): Collection of modules for multi-root builds.

Procedures And API Details
- **Replace Directive Syntax**: `replace <ModulePath> [Version] => <FilePath | ModulePath Version>`
  - Example: `replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5`
  - Multi-line block syntax allowed using parentheses.
- **Retract Directive Syntax**: `retract (Version | [Version, Version])`
  - Comments may precede or follow the directive to explain rationale.
  - Example: `retract v1.0.0 // Published accidentally.`
- **Automatic Update Flow**: Running commands with `-mod=mod` triggers rewriting of non-canonical versions (e.g., `v1` → `v1.0.0`) and resolving exclusions.
- **MVS Execution**: Traverses the graph starting from main modules, tracking highest required versions per module to produce the build list.
- **Pruning Logic**: For Go 1.17+, only immediate dependencies are loaded unless a lower-GO-version module transitively requires them.

Nuance Or Contradictions
- `replace` directives alone do not add a module to the graph; a corresponding `require` directive is still necessary.
- In Go 1.15 and earlier, `-mod=mod` was default; since Go 1.16, the command defaults to readonly mode unless explicitly invoked with `-mod=mod`.
- Retracted versions remain available but are excluded from version lists (`go list -m -versions`) and resolution queries like `@latest` unless the `-retracted` flag is used.
- Module graph pruning affects which checksums are recorded in `go.sum` for a given Go version, potentially requiring the `-compat` flag to adjust behavior.

Candidate Wiki Hints
- **Page**: `go-modules-replace-directive` (Concept: Local/Remote module substitution)
- **Page**: `go-modules-retract-directive` (Concept: Version retraction and lifecycle management)
- **Page**: `go-modules-minimal-version-selection` (Concept: MVS algorithm and build list computation)
- **Page**: `go-modules-graph-pruning` (Concept: Optimization in Go 1.17+ for dependency loading)

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## Chunk Context
This chunk details the syntax, directives, and behavior of `go.work` files used to define Go workspaces. It covers lexical elements, grammar, specific directives (`go`, `toolchain`, `godebug`, `use`, `replace`), compatibility handling for non-module repositories (specifically the `+incompatible` suffix), minimal module compatibility rules for GOPATH vs module mode, and the behavior of module-aware commands including vendoring.

## Local Summary
The section explains how to define a workspace using a `go.work` file, which is line-oriented with directives like `use`, `replace`, and `go`. It details the grammar (EBNF) for these directives. The text also addresses legacy compatibility issues where Go commands handle repositories without `go.mod` files by synthesizing them or adding an `+incompatible` suffix to major versions >= 2 in the root directory. Finally, it describes how the `go` command behaves in module-aware mode versus GOPATH mode, including the `-mod` flags and vendoring mechanisms.

## Key Claims
- A workspace is defined by a UTF-8 encoded text file named `go.work`.
- The `go.work` file must contain exactly one `go` directive specifying the toolchain version.
- Directives in `go.work` include `use`, `replace`, `toolchain`, and `godebug`.
- `use` directives add module directories to the workspace; they do not automatically include subdirectories.
- A wildcard replace in `go.work` overrides a version-specific replace in `go.mod`.
- Repositories with major version 2+ releases but no `go.mod` file get an `+incompatible` suffix appended by the Go command.
- Minimal module compatibility rules allow importing packages from a major version subdirectory (e.g., `/v2`) even if the GOPATH directory structure does not reflect that subdirectory, provided specific conditions are met.
- As of Go 1.16, module-aware mode is enabled by default regardless of the presence of a `go.mod` file.

## Entities And Concepts
- **go.work**: A text file defining a Go workspace containing multiple modules.
- **go.work.sum**: A file maintaining hashes for dependencies not present in collective workspace modules' `go.sum` files.
- **Directive**: A line in `go.work` starting with a keyword (e.g., `use`, `replace`).
- **+incompatible**: A suffix added by the Go command to versions >= 2.0.0 found in repositories lacking a `go.mod` file to indicate they are part of the same legacy module as lower major versions.
- **Minimal Module Compatibility**: Rules introduced in Go 1.11 allowing mixed usage of modules and GOPATH imports across major version subdirectories.
- **Module-aware mode**: The default state (Go 1.16+) where the `go` command uses `go.mod` files for dependency resolution.
- **Vendoring**: Storing dependencies in a local `vendor` directory instead of downloading from the module cache.

## Procedures And API Details
### Creating and Editing go.work Files
- **go work init**: Creates a new `go.work` file.
- **go work use**: Adds module directories to the `go.work` file.
- **go work edit**: Performs low-level edits on the `go.work` file.
- **golang.org/x/mod/modfile**: A package for making programmatic changes to `go.work` files.

### go.work Directives
- **go <version>**: Specifies the toolchain version (e.g., `go 1.23.0`). Must be a valid release version.
- **toolchain <name>**: Declares a suggested Go toolchain (effective only if the default is older).
- **godebug <setting>**: Applies a GODEBUG setting for the workspace; ignores `godebug` directives in individual `go.mod` files within the workspace.
- **use <path>**: Adds a module directory to the workspace. Supports grouping:
  ```go
  use (
    ./my/first/thing
    ./my/second/thing
  )
  ```
- **replace <old> => <new>**: Replaces module contents. Supports wildcards and version ranges. Overrides replaces in `go.mod`.
  ```go
  replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
  ```

### Compatibility Handling
- **Synthetic go.mod**: If a module path equals the repository root and lacks a `go.mod`, the command synthesizes one in the cache with only a module directive.
- **+incompatible suffix**: Applied to versions >= 2.0.0 in repositories without `go.mod`. Tags like `v4.1.2+incompatible` are ignored by the repository; the suffix appears only in Go command usage.

### Module-Aware Commands and Flags
- **Module-aware commands**: `go build`, `go test`, `go install`, etc., use `go.mod` files to interpret import paths.
- **-mod flag**: Controls dependency resolution:
  - `-mod=mod`: Ignore vendor, auto-update `go.mod`.
  - `-mod=readonly`: Ignore vendor, error if `go.mod` needs update.
  - `-mod=vendor`: Use vendor directory; no network/cache access.
- **-modcacherw**: Creates module cache directories with read-write permissions.
- **-modfile=file.mod**: Reads/writes an alternate `.mod` file instead of `go.mod`.

### Vendoring
- **go mod vendor**: Constructs a `vendor` directory and `vendor/modules.txt`.
- Vendor directory is used automatically if present and `go.mod` version >= 1.14, or explicitly via `-mod=vendor`.

## Nuance Or Contradictions
- **Committing go.work files**: Generally inadvisable due to potential conflicts with parent directory workspaces or CI systems testing wrong dependency versions. However, it is acceptable if modules are developed exclusively together without external dependencies.
- **GOPATH vs Module Path Mismatch**: In GOPATH mode, a module at `example.com/repo/v2` might be found at `$GOPATH/src/example.com/repo/sub` (without the `/v2`) if not developed in the subdirectory, causing import path mismatches unless minimal compatibility rules apply.
- **godebug precedence**: `godebug` directives in `go.mod` files are ignored when a workspace (`go.work`) is active; only `go.work` `godebug` directives apply.

## Candidate Wiki Hints
- **Topic: go.work Syntax** – Document the structure, directives, and EBNF grammar for workspace definition files.
- **Topic: Workspace Directives (use/replace)** – Detail how to add modules and override dependencies in a workspace context.
- **Topic: Legacy Module Compatibility** – Explain the `+incompatible` suffix and handling of pre-module repositories.
- **Topic: Minimal Module Compatibility** – Describe the rules allowing cross-mode imports between major version subdirectories.
- **Topic: Vendoring in Go Modules** – Cover the creation and usage of `vendor` directories with `go mod vendor`.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
Heading path: Go Modules Reference > Retrieved Text
Lines: 1666-1685

Local Summary
This section explains how the `go` command behaves when vendoring is enabled versus GOPATH mode. It clarifies that build and test commands use the vendor directory at the main module's root, while module management commands like `go mod download`, `go mod tidy`, and `go get` continue to interact with the network and module cache regardless of vendoring status. A key distinction is made regarding where vendor directories are ignored outside the main module's root.

Key Claims
- When vendoring is enabled, build commands (`go build`, `go test`) load packages from the vendor directory instead of accessing the network or local module cache.
- The `go list -m` command only prints information about modules listed in `go.mod`.
- Module management commands (`go mod download`, `go mod tidy`) function identically when vendoring is enabled; they still download modules and access the module cache.
- The `go get` command does not behave differently when vendoring is enabled.
- Unlike GOPATH mode, the `go` command ignores vendor directories located outside the main module’s root directory.
- Because vendor directories in other modules are not used during builds, they are excluded from generated module zip files (with references to known bugs #31562 and #37397).

Entities And Concepts
- Vendoring: A practice of copying dependencies into a local directory to avoid network access.
- Vendor Directory: The directory containing vendored packages; only the one in the main module's root is respected by build commands.
- Module Cache: The local storage for downloaded modules, accessed even when vendoring is active for management tasks.
- Main Module’s Root Directory: The specific location where vendor directories are recognized and utilized.
- Module Zip Files: Archive files generated during builds that exclude unused vendor directories from submodules.

Procedures And API Details
Command Usage:
- `go build`: Loads packages from the vendor directory if enabled; otherwise, uses network/cache.
- `go test`: Same behavior as `go build` regarding vendoring.
- `go list -m`: Prints module info strictly for those listed in `go.mod`.
- `go mod download`: Downloads modules and accesses cache regardless of vendoring status.
- `go mod tidy`: Cleans up `go.mod` and `go.sum` without being affected by vendoring status.
- `go get`: Does not change behavior when vendoring is enabled; usage includes flags like `-d`, `-t`, `-u`, `-tool`.

Nuance Or Contradictions
- **Vendoring Scope**: While `go build` uses the vendor directory at the main module's root, it explicitly ignores vendor directories in other modules. This means sub-dependencies with their own vendor folders are not utilized for building the main module.
- **Module Zip Files**: Building a module zip file excludes vendor directories found in non-main modules due to the logic described above, though this behavior is noted as having known bugs (#31562, #37397).
- **Management Commands vs Build Commands**: There is a functional divergence; build commands respect the main vendor directory, but management commands (`go mod`, `go get`) bypass vendoring to ensure dependencies are up-to-date in the cache.

Candidate Wiki Hints
- Create a page explaining the scope of vendor directories in Go modules, specifically distinguishing between the main module's root and submodules.
- Document the behavior differences between build-time loading and module-management operations regarding vendoring.

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## Chunk Context
This chunk (lines 1686–2071) details the `go get` command for managing module dependencies, including upgrading/downgrading specific modules, handling transitive dependencies, and managing version queries. It also introduces `go install` for building executables, `go tool` for running tools, `go list -m` for listing modules, `go mod download` for caching, and `go mod edit` for modifying `go.mod`.

## Local Summary
The text explains how to use `go get` to update, downgrade, or remove specific module versions, including handling transitive dependencies. It highlights the deprecation of building/installing packages with `go get` in favor of `go install`. The chunk also covers version query syntax (`@v1.0.0`, `@master`, `@latest`, `@none`), flags for `go get` (like `-u`, `-d`, `-tool`), and related commands like `go mod download` and `go mod edit`.

## Key Claims
- `go get` updates module dependencies in the main module’s `go.mod` file and builds/installing packages is deprecated since Go 1.17.
- Version query suffixes (`@v`, `@master`, `@latest`, `@none`) control version selection for modules.
- Transitive dependencies are upgraded or downgraded automatically when their direct dependencies change versions.
- `go install` is recommended for installing programs since Go 1.16, ignoring the local `go.mod` if version suffixes are provided.
- `go mod download` pre-fills the module cache; `go mod edit` modifies `go.mod` (e.g., adding replace directives).

## Entities And Concepts
- **Commands**: `go get`, `go install`, `go tool`, `go list -m`, `go mod download`, `go mod edit`.
- **Version Queries**: `@v1.0.0`, `@master`, `@latest`, `@upgrade`, `@patch`, `@none`.
- **Flags**: `-u` (upgrade), `-d` (no build, deprecated), `-tool`, `-insecure`, `-t` (test dependencies).
- **Structs**: `Module`, `ModuleError`.
- **Environment Variables**: `GOBIN`, `GOPATH`, `GOROOT`, `GOINSECURE`, `GO111MODULE`.

## Procedures And API Details
- **Upgrading a specific module**:
  ```bash
  go get golang.org/x/net
  ```
- **Downgrading to `@none`**:
  ```bash
  go get golang.org/x/text@none
  ```
- **Installing a program (ignoring local `go.mod`)**:
  ```bash
  go install golang.org/x/tools/gopls@latest
  ```
- **Listing modules with upgrades**:
  ```bash
  go list -m -u all
  ```
- **Downloading modules**:
  ```bash
  go mod download golang.org/x/mod@v0.2.0
  ```
- **Adding a replace directive**:
  ```bash
  go mod edit -replace example.com/a@v1.0.0=./a
  ```

## Nuance Or Contradictions
- `go get` without `-d` is deprecated since Go 1.17; in Go 1.18, `-d` is always enabled.
- `go install` ignores the local `go.mod` if version suffixes are provided (module-aware mode), whereas without suffixes it may run in GOPATH or module-aware mode depending on `GO111MODULE`.
- Retracted or deprecated modules are flagged by `go get` and `go list -m -u`, but `go list -m -retracted` explicitly includes retracted versions in version lists.

## Candidate Wiki Hints
- **Page**: `go-get-command` – Covers usage, flags, version queries, and transitive dependency handling.
- **Page**: `go-install-command` – Focuses on installing executables with module-aware behavior.
- **Page**: `module-version-queries` – Details syntax for `@v`, `@master`, `@latest`, etc.

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
## Chunk Context
Lines 2072–2430 cover commands for managing `go.mod` files and module graphs, specifically focusing on printing JSON representations of `go.mod`, editing operations (`go mod edit`), and subsequent tools like `go mod graph`, `go mod init`, `go mod tidy`, `go mod vendor`, `go mod verify`, `go mod why`, and `go version -m`.

## Local Summary
This chunk details how to inspect, manipulate, and maintain Go module metadata. It explains the JSON output of `go mod edit -json` (with associated Go types), describes editing flags for modifying requirements, replacements, retractions, and tools. It then moves into commands for generating dependency graphs (`go mod graph`), initializing modules (`go mod init`), tidying dependencies (`go mod tidy`), vendoring packages (`go mod vendor`), verifying integrity (`go mod verify`), tracing import paths (`go mod why`), and reporting Go versions (`go version -m`).

## Key Claims
- `go mod edit -json` prints the `go.mod` file in JSON format without writing to disk.
- Editing flags like `-require`, `-exclude`, `-replace`, `-retract`, and `-tool` modify the module graph or directives.
- `go mod tidy` aligns `go.mod` with imported packages, adding missing requirements and removing unused ones.
- `go mod vendor` creates a `vendor` directory containing copies of necessary packages for builds/tests.
- `go mod verify` ensures downloaded modules haven't been tampered with by comparing hashes in the module cache.
- `go mod why` displays shortest import paths from the main module to specified packages/modules.

## Entities And Concepts
- **Commands**: `go mod edit`, `go mod graph`, `go mod init`, `go mod tidy`, `go mod vendor`, `go mod verify`, `go mod why`, `go version`.
- **Flags**: `-json`, `-fmt`, `-print`, `-require`, `-droprequire`, `-exclude`, `-replace`, `-dropreplace`, `-retract`, `-tool`, `-droptool`, `-e`, `-v`, `-x`, `-diff`, `-go`, `-compat`, `-o`, `-m`, `-vendor`.
- **Types (from JSON output)**: `Module`, `GoMod`, `ModPath`, `Require`, `Replace`, `Retract`, `Tool`.
- **Files**: `go.mod`, `go.sum`, `vendor/modules.txt`.

## Procedures And API Details
1. **Printing JSON representation of `go.mod`**:
   ```bash
   go mod edit -json
   ```
   Output corresponds to Go types:
   ```go
   type Module struct { Path string; Version string }
   type GoMod struct { Module ModPath; Go string; Require []Require; Exclude []Module; Replace []Replace; Retract []Retract }
   type ModPath struct { Path string; Deprecated string }
   type Require struct { Path string; Version string; Indirect bool }
   type Replace struct { Old Module; New Module }
   type Retract struct { Low string; High string; Rationale string }
   type Tool struct { Path string }
   ```

2. **Editing `go.mod`**:
   - Use `-module`, `-go=version`, `-require=path@version`, `-droprequire=path`, `-exclude=path@version`, `-dropexclude=path@version`, `-replace=old[@v]=new[@v]`, `-dropreplace=old[@v]`, `-retract=version`, `-dropretract=version`, `-tool=path`, `-droptool=path`.
   - Repeat flags; changes apply in order.

3. **Formatting `go.mod`**:
   ```bash
   go mod edit -fmt
   ```
   (Implied by other modification flags.)

4. **Generating module graph**:
   ```bash
   go mod graph [-go=version]
   ```
   Outputs edges as `module@version dependency`.

5. **Initializing a module**:
   ```bash
   go mod init [module-path]
   ```
   Infers path if omitted (uses import comments and GOPATH).

6. **Tidying dependencies**:
   ```bash
   go mod tidy [-e] [-v] [-x] [-diff] [-go=version] [-compat=version]
   ```
   Adds missing requirements, removes unused ones, updates `go.sum`.

7. **Vendoring packages**:
   ```bash
   go mod vendor [-e] [-v] [-o]
   ```
   Creates `vendor` directory; generates `vendor/modules.txt`.

8. **Verifying module integrity**:
   ```bash
   go mod verify
   ```
   Compares hashes of downloaded modules with those in the module cache.

9. **Tracing import paths**:
   ```bash
   go mod why [-m] [-vendor] packages...
   ```
   Shows shortest path from main module to specified packages/modules.

10. **Reporting Go version**:
    ```bash
    go version [-m] [-v] [file ...]
    ```

## Nuance Or Contradictions
- `-require` overrides existing requirements on the same path, unlike `go get` which adjusts constraints automatically.
- `-replace` without `@v` applies to all versions of the old module; with `@v`, it targets a specific version.
- `go mod tidy` considers all packages imported by tests but excludes those in `.go` files tagged with `// +build ignore`.
- `go mod vendor` removes existing `vendor` directory before recreating it; local changes to vendored packages are not checked for integrity.
- `go mod verify` does not download missing modules; it only checks cached ones and may add entries to `go.sum` if needed.

## Candidate Wiki Hints
- **Page**: `go-mod-edit-json` – Documenting the JSON output schema of `go mod edit`.
- **Page**: `go-mod-tidy-behavior` – Explaining how `go mod tidy` handles imports, tests, and build tags.
- **Page**: `go-mod-vendor-workflow` – Covering vendoring setup, manifest usage, and integrity checks.

## chunk-09

---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
The chunk covers three main areas within the Go module system:
1.  **Inspecting Executables**: Using `go version` to print Go and module versions for binaries in a directory.
2.  **Module Version Queries**: Specifying versions using queries (e.g., `@latest`, `@master`) and understanding how they resolve against available tags and pseudo-versions.
3.  **Module Proxies**: The protocol, path structures (`$base/$module/@v/list`), and behavior of HTTP module proxies used by the `go` command to fetch source code and metadata.

Local Summary
This section details how to inspect build information for Go executables using `go version -m`, which outputs a table containing the main module path, version, sum, and dependency trees. It explains that this output format may change in the future and points to `runtime/debug.ReadBuildInfo` as an alternative source. The text then defines various version query syntaxes (e.g., semantic versions, revision hashes, `latest`, `upgrade`, `patch`) and describes their resolution logic, noting that release versions are preferred over pre-releases unless specific flags are used. Finally, it outlines the module proxy protocol, specifying required endpoints like `/@v/list` for version lists and `/@v/$version.mod` for source code, including details on case encoding for path safety and fallback behaviors using comma or pipe separators in the `GOPROXY` environment variable.

Key Claims
-   `go version -m` prints a tab-separated table with columns: `path`, `mod`, `dep`, and `=>`.
-   The `-v` flag of `go version` reports unrecognized files found during a directory scan.
-   Version queries can be specific semantic versions, prefixes, comparisons, revision identifiers (hashes/tags), or keywords like `latest` and `upgrade`.
-   Except for specific named versions or revisions, queries consider only tagged versions reported by `go list -m -versions`, excluding pseudo-versions unless necessary.
-   Release versions are preferred over pre-release versions; if no release is available, `latest`, `upgrade`, and `patch` select the highest pre-release or a pseudo-version for the tip of the default branch.
-   Module proxies must respond to specific paths (e.g., `$base/$module/@v/list`) with plain text lists of versions or JSON metadata.
-   Proxies must serve consistent content for `.mod` and `.zip` files, which are authenticated via `go.sum`.
-   The `GOPROXY` environment variable accepts comma-separated URLs (fallback on 404/410) or pipe-separated URLs (fallback on any error).

Entities And Concepts
-   **`go version -m`**: Command to print Go and module versions for executables.
-   **Table Columns**: `path` (main package path), `mod` (main module info), `dep` (dependency modules), `=>` (module replacements).
-   **Version Queries**: Syntax following the `@` character (e.g., `@v1.2.3`, `@master`, `@latest`).
-   **Pseudo-version**: A version string representing a commit at the tip of a branch when no tagged version exists.
-   **Module Proxy**: An HTTP server responding to module metadata and source requests.
-   **GOPROXY**: Environment variable configuring proxy URLs and fallback strategies.
-   **Case Encoding**: Replacing uppercase letters with `!lower-case` in paths to handle case-insensitive filesystems (e.g., `example.com/M` becomes `example.com/!m`).

Procedures And API Details
-   **Inspecting Binaries**:
    -   Command: `$ go version -m <directory>`
    -   Output Format: Tab-separated table per executable.
    -   Example columns: `path`, `mod`, `dep`, `=>`.
-   **Version Queries**:
    -   Syntax: `<module>@<query>` (e.g., `example.com/m@latest`).
    -   Query Types:
        -   Fully-specified semantic version (`v1.2.3`).
        -   Semantic version prefix (`v1`, `v1.2`).
        -   Comparison (`<v1.2.3`, `>=v1.5.6`).
        -   Revision identifier (commit hash, tag, branch).
        -   Keywords: `latest`, `upgrade`, `patch`.
-   **Module Proxy Endpoints**:
    -   List versions: `$base/$module/@v/list` (Plain text list).
    -   Get metadata: `$base/$module/@v/$version.info` (JSON with `Version` and optional `Time`).
    -   Get go.mod: `$base/$module/@v/$version.mod`.
    -   Get zip: `$base/$module/@v/$version.zip`.
    -   Get latest info: `$base/$module/@latest`.
-   **Configuring Proxies**:
    -   Set `GOPROXY` to a list of URLs.
    -   Use `,` for fallback on 404/410.
    -   Use `|` for fallback on any error.

Nuance Or Contradictions
-   **Version Query Resolution**: While most queries consider only tagged versions, specific revision identifiers (like commit hashes) select pseudo-versions if the revision isn't tagged with a semantic version. The `latest`, `upgrade`, and `patch` queries behave differently when no release/pre-release exists compared to other queries which report an error in that scenario.
-   **Case Encoding**: Module paths are case-encoded (uppercase -> `!lowercase`) specifically to avoid ambiguity on case-insensitive filesystems, allowing storage of both `Example.com/Module` and `example.com/module`.
-   **Proxy Content Consistency**: Proxies must always serve the same content for `.mod` and `.zip` queries to ensure cryptographic authentication via `go.sum` works correctly.

Candidate Wiki Hints
-   Page: **Go Version Queries** (Explaining syntax, resolution order, and behavior of `latest`, `upgrade`, `patch`).
-   Page: **Module Proxy Protocol** (Defining required endpoints, JSON structures, and error handling).
-   Page: **GOPROXY Configuration** (Explaining comma vs. pipe separators and fallback logic).

## chunk-10

---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
## Chunk Context
This chunk details the internal mechanics of how the `go` command resolves, downloads, and builds modules. It covers the build list computation procedure, the specific HTTP requests made to proxies (`.mod`, `.zip`, `.info`, `.list`), verification via cryptographic hashes, direct mode serving from version control systems, repository discovery via `<meta name="go-import">` tags, mapping semantic/pseudo-versions to commits, and handling module subdirectories within repositories.

## Local Summary
The document explains that `go build` computes a build list using Minimal Version Selection (MVS), loads modules by downloading `.mod` and `.zip` files via proxy requests, and verifies integrity using checksums against `go.sum`. It describes "direct mode" for downloading from version control systems (Git, Subversion, etc.) when proxies are unavailable or private. The text details how the `go` command discovers repositories via `?go-get=1` queries and parses `<meta name="go-import">` tags to determine the root path, VCS type, and repository URL. It further explains how version tags map to commits (semantic versions vs. pseudo-versions) and how modules are located within repository subdirectories, including handling of major version suffixes for v2+ compatibility.

## Key Claims
*   The `go build` procedure involves computing a build list via MVS, loading packages, finding missing modules, and building.
*   Module source code is distributed in `.zip` files extracted into the module cache; `.mod` files are downloaded separately if not included.
*   If a package is not provided by any module in the build list, the command requests information about the latest version of potential module paths to find a provider.
*   When requesting a module version, the sequence is `$module/@v/list` -> `$module/@latest` (if needed) -> `$module/@v/$version.info` -> `.mod`/`.zip`.
*   Downloaded files are verified via cryptographic hash comparison against `go.sum`; mismatches trigger security errors unless `GOPRIVATE`, `GONOSUMDB`, or `GOSUMDB=off` is set.
*   Direct mode allows downloading from VCS repositories; this requires a tool (like `git`) in PATH and often uses `GOPROXY=direct`.
*   Repository discovery relies on an HTTP GET with `?go-get=1` looking for `<meta name="go-import">` containing `root-path vcs repo-url [subdirectory]`.
*   Pseudo-versions (e.g., `v1.3.2-0.20191109021931-daa7c04131f5`) encode a timestamp and commit hash prefix to ensure reproducible builds from specific revisions.
*   Modules defined in subdirectories must have their `go.mod` file located within that subdirectory, which may or may not match the major version suffix depending on the release strategy.

## Entities And Concepts
*   **Build List**: A list of modules and versions selected for a build, computed via Minimal Version Selection (MVS).
*   **GOPROXY Protocol**: The protocol used to request `.mod`, `.zip`, and metadata files from a proxy server.
*   **Direct Mode**: Downloading modules directly from a VCS repository instead of a proxy.
*   **go-import Meta Tag**: An HTML meta tag (`<meta name="go-import">`) used by servers to advertise their version control repository details to the `go` command.
*   **Pseudo-version**: A specific revision encoding (e.g., `vX.Y.Z-0.TIMESTAMP-HASH`) used when a module is not tagged with a semantic version or needs precise commit selection.
*   **Module Subdirectory**: The portion of the module path that corresponds to a directory within the repository root (e.g., `foo/bar` for `example.com/foo/bar`).
*   **Major Version Subdirectory**: A subdirectory matching a major version suffix (e.g., `/v2`) used to host multiple major versions on a single branch.

## Procedures And API Details
### Module Download Sequence
1.  **Request List**: `$module/@v/list` to get available versions.
2.  **Select Version**: If list is empty or unusable, request `$module/@latest`.
3.  **Get Metadata**: Request `$module/@v/$version.info`.
4.  **Download Files**: Request `$module/@v/$version.mod` and `$module/@v/$version.zip`.
5.  **Verify**: Compute hash of downloaded files and check against `go.sum`.

### Repository Discovery (Direct Mode)
1.  Construct URL: `https://<root-path>?go-get=1`.
2.  Parse HTML Response: Look for `<meta name="go-import" content="<root-path> <vcs> <repo-url> [<subdirectory]>">`.
3.  Clone/Fetch: Use the specified VCS tool (`git`, `hg`, etc.) with the `repo-url` to clone into the module cache.

### Version Mapping Logic
*   **Semantic Tag**: If a tag exists (e.g., `v1.2.3`), use it directly.
*   **Pseudo-Tag Generation**: If no valid semantic tag, generate pseudo-version based on commit hash and timestamp.
*   **Branch/Commit Query**: Use `go get example.com/mod@master`; the command converts branch/revision names to canonical versions.

## Nuance Or Contradictions
*   `.mod` files are usually inside `.zip` files but can be requested separately because `.mod` requests are smaller and faster; the text emphasizes they are "separate" in terms of request handling.
*   Modules served directly from a proxy cannot be downloaded with `go get` in GOPATH mode.
*   Tags for modules in subdirectories (e.g., `gopls/v0.4.0`) must include the subdirectory prefix, unlike root-level modules where tag names match versions exactly.
*   Go 1.25+ supports subdirectories in `<meta name="go-import">` tags; earlier versions ignore these and may fail resolution if the module isn't in the repository root.

## Candidate Wiki Hints
*   **Topic**: Go Module Proxy Protocol (`GOPROXY`)
*   **Topic**: Minimal Version Selection (MVS) and Build List Computation
*   **Topic**: Direct Mode Downloading from VCS
*   **Topic**: Repository Discovery via `?go-get=1`
*   **Topic**: Pseudo-versions and Reproducible Builds

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
# Chunk Context

This chunk details the mechanics of module distribution via `.zip` files, security constraints on version control usage (`GOVCS`), and configuration for accessing private modules (proxies, direct VCS access, authentication). It covers file size limits, path constraints within zip archives, and environment variables like `GOPRIVATE`, `GOPROXY`, and `GONOSUMDB`.

# Local Summary

The Go toolchain distributes modules as authenticated `.zip` files containing only the module root contents (excluding vendor directories or nested modules). The `go` command enforces strict file constraints to ensure cross-platform compatibility and security, such as ignoring symbolic links and limiting zip sizes to 500 MiB. To manage security risks associated with untrusted version control servers, the `GOVCS` environment variable restricts allowed VCS tools (defaulting to `git` and `hg` for public paths) unless overridden. For private modules, administrators can configure `GOPROXY` and `GONOSUMDB` to use internal proxies or bypass checksum verification, while `GOPRIVATE` marks module paths as private to avoid external proxy lookups. Authentication for private resources is handled via `.netrc` files or embedded credentials in URLs, with support for non-interactive Git configurations.

# Key Claims

- Module zip files are authenticated before extraction; they do not include vendor directories or nested modules (subdirectories with `go.mod`).
- If a module lacks a `LICENSE` file in its root, the `go` command copies it from the repository root if present in the same revision.
- Symbolic links are excluded from module zip files to ensure portability across operating systems.
- The default behavior for public modules is to download via proxy (`proxy.golang.org`) rather than directly cloning version control repositories.
- `GOVCS` allows fine-grained control over which VCS tools (e.g., `git`, `hg`) are permitted for specific module paths.
- Private modules can be accessed without a private proxy by setting `GOPRIVATE` and configuring direct VCS access, potentially requiring non-interactive authentication.

# Entities And Concepts

- **Module Zip Files**: Distributed artifacts containing module contents; created, downloaded, and extracted automatically by the `go` command.
- **GOVCS**: Environment variable controlling allowed version control systems for downloading code.
- **GOPRIVATE**: Glob patterns marking module paths as private to bypass public proxies and checksum databases.
- **GOPROXY**: List of module proxy URLs; used to download modules sequentially, falling back to direct VCS access if unavailable.
- **GONOSUMDB**: Patterns for modules not to be checked against the public checksum database (`sum.golang.org`).
- **.netrc**: File used to store credentials for HTTP basic authentication with proxy servers.
- **Direct Access**: Configuration mode where `go` downloads private modules directly from VCS repositories without a proxy intermediary.

# Procedures And API Details

- **Setting GOVCS**: Define a comma-separated list of rules (e.g., `github.com:git,evil.com:off,*:git|hg`) to restrict VCS usage per module path pattern.
- **Configuring Private Proxy (All Modules)**: Set `GOPROXY` to the internal proxy URL and `GONOSUMDB` to the module prefix to disable public checksum checks.
- **Configuring Direct Access**: Set `GOPRIVATE=corp.example.com` to bypass proxies for private modules; ensure VCS tools have non-interactive authentication configured (e.g., `.gitconfig`).
- **Authenticating with .netrc**: Place credentials in `$HOME/.netrc` (or `%USERPROFILE%\_netrc`) using the `machine`, `login`, and `password` fields for proxy servers.
- **Embedding Credentials in URLs**: Append username:password to the `GOPROXY` URL (e.g., `https://user:pass@proxy.example.com`), though this may expose secrets in shell history or logs.

# Nuance Or Contradictions

- While `GOVCS` defaults to restricting public paths to `git` and `hg`, public modules served via the module mirror (`proxy.golang.org`) are still accessible even if they originate from unsupported VCS systems like Bazaar, because the proxy handles the download securely.
- The special case for copying `LICENSE` files applies only to files named exactly `LICENSE` (no extension); extending this behavior would break cryptographic sums used for authentication.
- Empty directories can be included in zip files but are not extracted by the `go` command.
- File names up to the first dot cannot be reserved Windows filenames (e.g., `CON`, `COM1`) regardless of case, even if the full path is valid.

# Candidate Wiki Hints

- **Module Distribution Format**: Explain the structure and security constraints of Go module zip files.
- **GOVCS Configuration Guide**: Document how to configure version control restrictions for mixed public/private ecosystems.
- **Private Module Setup**: Step-by-step guide for configuring `GOPRIVATE`, `GOPROXY`, and authentication for internal repositories.
- **License Handling in Modules**: Clarify the automatic `LICENSE` file copying behavior and its limitations.

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
This chunk details the privacy and security mechanisms of the Go toolchain, focusing on module proxy configuration (GOPROXY), checksum database usage (GOSUMDB), private module handling (GOPRIVATE/GONOPROXY/GONOSUMDB), and the structure of the module cache. It explains how `go` verifies cryptographic hashes to ensure download integrity and outlines the fallback behaviors when accessing non-existent or private modules.

Local Summary
The `go` command manages privacy and security by routing requests through configured proxies while protecting private module paths. It utilizes a global checksum database (default: `sum.golang.org`) to verify downloaded modules against cryptographic hashes. The tool distinguishes between public and private modules using environment variables, falling back to direct version control access if a proxy returns 404/410 for private paths. The module cache stores extracted sources and metadata with read-only permissions by default, requiring specific commands or flags for cleanup or modification.

Key Claims
- The default `GOPROXY` setting is `https://proxy.golang.org,direct`, prioritizing Google's public proxy before falling back to direct access.
- `GOPRIVATE` acts as a default for both `GONOPROXY` and `GONOSUMDB`, meaning explicit settings are only needed if the proxy or checksum database behaviors differ between them.
- If a private proxy responds with 404 (Not Found) or 410 (Gone), the `go` command falls back to the public proxy, transmitting the full module path; other error codes halt the process and print an error.
- The checksum database (`sum.golang.org`) uses a Transparent Log (Merkle Tree) structure backed by Trillian to ensure untrusted proxies cannot serve wrong code without detection.
- Module cache files are created with read-only permissions by default to prevent accidental modification, making manual deletion difficult without `go clean -modcache` or using the `-modcacherw` flag.

Entities And Concepts
- **GOPROXY**: Environment variable controlling module proxy servers.
- **GONOPROXY**: Environment variable for modules that should not be requested from any proxy (defaulted by GOPRIVATE).
- **GONOSUMDB**: Environment variable for modules that should not be requested from the checksum database (defaulted by GOPRIVATE).
- **GOPRIVATE**: Pattern list for private module prefixes; defaults `GONOPROXY` and `GONOSUMDB`.
- **GOSUMDB**: Environment variable setting the name, URL, and public key of the checksum database.
- **Module Cache**: Directory (`$GOPATH/pkg/mod`) storing downloaded module files.
- **Checksum Database**: Global source of hashes (default: `sum.golang.org`).
- **go.sum**: File containing cryptographic hashes of direct and indirect dependencies.
- **Inclusion/Consistency Proofs**: Cryptographic proofs performed by the `go` command to verify data integrity against the checksum database log.

Procedures And API Details
- **Configuring Private Modules**: Set `GOPRIVATE=*.corp.example.com,*.research.example.com` to prevent proxy requests for specific module prefixes.
- **Disabling Checksum Verification**: Use `GOSUMDB=off` or invoke `go get -insecure` to bypass the checksum database (not recommended).
- **Clearing Module Cache**: Run `go clean -modcache`. Alternatively, use `go env -w GOMODCACHE=<path>` and delete contents manually if using the `-modcacherw` flag.
- **Verifying Dependencies**: Use `go mod verify` to scan extracted module contents and confirm they match expected hashes in `go.sum`.
- **Checksum Database Lookup**: The client sends a GET request for `$base/lookup/$module@$version`. If not found, it fetches from the origin server before replying with log record data and signed tree descriptions.

Nuance Or Contradictions
- There is no direct contradiction, but there is a trade-off: while read-only permissions protect against accidental edits, they complicate manual cache maintenance. The `-modcacherw` flag increases security risk by allowing editors to modify files.
- A typo in a module path (e.g., `corp.example.com/secret-product/typo`) causes the private proxy to return 404/410, triggering a fallback to the public proxy which leaks the path; however, other error codes prevent this fallback and result in an immediate error.

Candidate Wiki Hints
- **Proxy Configuration**: How to configure `GOPROXY`, `GONOPROXY`, and `GOPRIVATE` for corporate environments with trusted internal proxies.
- **Module Security**: Understanding the role of `go.sum`, `GOSUMDB`, and checksum proofs in securing module downloads.
- **Private Modules**: Strategies for handling private modules using environment variables and proxy fallback logic.
- **Cache Management**: Best practices for managing the module cache, including read-only constraints and cleanup procedures.

## chunk-13

---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Chunk Context

The document discusses the behavior of the `go` command regarding module-related environment variables. It lists specific variables like `GO111MODULE`, `GOMODCACHE`, `GOINSECURE`, and others, detailing their functions and default behaviors. Additionally, it includes a glossary defining key terms related to Go modules, such as "build constraint," "canonical version," "direct dependency," and "minimal version selection (MVS)." The chunk concludes with navigation links and footer information from the go.dev website.

# Local Summary

This section outlines various environment variables that control the `go` command's behavior in module-aware or GOPATH modes. Key variables include `GO111MODULE` for toggling module mode, `GOPROXY` for managing proxy URLs, and `GOSUMDB` for checksum verification. The glossary provides definitions for essential concepts like "main module," "pseudo-version," and "workspace."

# Key Claims

- The `go` command supports three modes for `GO111MODULE`: off, on, and auto.
- `GOPROXY` can be set to URLs or keywords like "off" and "direct".
- `GOSUMDB` defaults to `sum.golang.org`, the Go checksum database run by Google.
- Minimal version selection (MVS) determines the versions of all modules used in a build.
- Major version suffixes are required at v2.0.0 and later.

# Entities And Concepts

- **GO111MODULE**: Controls module-aware mode vs. GOPATH mode.
- **GOMODCACHE**: Directory for storing downloaded modules.
- **GOPROXY**: List of module proxy URLs or keywords ("off", "direct").
- **GOSUMDB**: Checksum database identifier and URL.
- **GOVCS**: Controls version control tools for downloading modules.
- **GOWORK**: Enables workspace mode using a `go.work` file.
- **Build constraint**: Condition determining if a Go source file is used.
- **Minimal version selection (MVS)**: Algorithm for selecting module versions.

# Procedures And API Details

- To disable module checksum verification, set `GOPRIVATE` or `GONOSUMDB`.
- Use `GOPROXY=direct` to download modules directly from version control systems.
- Set `GO111MODULE=off` to ignore `go.mod` files and run in GOPATH mode.

# Nuance Or Contradictions

- In Go 1.15 and lower, `auto` was the default for `GO111MODULE`.
- `GOPROXY` defaults to `https://proxy.golang.org,direct`, contacting Google's mirror first.
- Pseudo-versions encode revision identifiers and timestamps for compatibility with non-module repositories.

# Candidate Wiki Hints

- **Environment Variables in Go Modules**: A comprehensive guide to configuring `go` command behavior using environment variables like `GO111MODULE`, `GOPROXY`, and `GOSUMDB`.
- **Minimal Version Selection (MVS)**: Explanation of how MVS determines module versions for builds.
- **Major Version Suffixes**: Guidelines on when and how to use major version suffixes in module paths.

