## chunk-01

---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Chunk Context

**Source:** `raw/web/corpus-2026-05-18/032-go-modules-reference.md`
**Line Range:** 1–26
**Heading Path:** Document > Go Modules Reference > Fetch Metadata
**Original Title:** Go Modules Reference - The Go Programming Language

# Local Summary

This initial chunk establishes the metadata for a reference document regarding Go Modules. It confirms the retrieval of content from `https://go.dev/ref/mod` on 2026-05-18. The specific section covered introduces the concept of fetching module metadata, which is fundamental to dependency management in the Go language ecosystem.

# Key Claims

- The document serves as a reference for Go Modules.
- Fetching metadata is identified as the primary initial step in the module workflow described.
- The source content is verified as HTML text retrieved successfully.

# Entities And Concepts

- **Go Modules:** The dependency management system for Go.
- **Fetch Metadata:** The process of retrieving information about available modules and their versions.
- **Dependency Management:** The broader category of tasks facilitated by fetching metadata.

# Procedures And API Details

No specific API calls or procedural steps are detailed in this chunk; it serves primarily as a structural introduction to the reference material.

# Nuance Or Contradictions

None identified within this short excerpt. The text is purely declarative regarding document metadata and retrieval status.

# Candidate Wiki Hints

- **Page:** `go-modules-fetch-metadata`
  - **Focus:** Explain the initial phase of module resolution where the build tool queries remote registries for version tags and information.

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## Chunk Context

The chunk covers the core concepts of Go Modules, including their definition, identification via `go.mod` files, path structures, versioning schemes (semantic and pseudo-versions), major version suffixes for v2+, and the resolution logic used by the `go` command to locate packages within modules.

## Local Summary

This section introduces the Go module system as the mechanism for managing dependencies. It defines a module as a collection of versioned packages, identified by a path declared in a `go.mod` file. The text details how module paths are constructed (repository root, subdirectory, major version suffix), explains semantic versioning rules, and describes pseudo-versions used to reference specific commits or tags. It also outlines the rules for major version suffixes starting at v2 to ensure import compatibility between incompatible versions and concludes with an overview of how the `go` command resolves package paths to specific modules using proxies and build lists.

## Key Claims

- Modules are the standard way Go manages dependencies, consisting of collections of packages released and distributed together.
- A module is identified by a module path declared in a `go.mod` file located in the module root directory.
- Module paths must describe both functionality and location, typically comprising a repository root path, an optional subdirectory, and a major version suffix for versions 2+.
- Semantic versioning requires incrementing the major version after backwards-incompatible changes, minor versions after compatible changes, and patch versions for bug fixes or optimizations.
- Starting with major version 2, modules must use a major version suffix (e.g., `/v2`) to distinguish incompatible packages; this is not required for v0 or v1 unless using `gopkg.in/` paths.
- Pseudo-versions encode specific revision information (timestamp and commit hash) and sort between the base version and the next tagged version to facilitate testing and dependency resolution without manual typing.
- The `go` command resolves packages by searching the build list for matching module prefixes, consulting proxies defined in `GOPROXY`, and preferring the module with the longest matching path.

## Entities And Concepts

- **Module**: A collection of packages released, versioned, and distributed together.
- **Module Path**: The canonical name declared in `go.mod`; serves as a prefix for package paths within the module.
- **Package Path**: The module path joined with the subdirectory containing the package.
- **Semantic Versioning**: A system where versions are formatted as `vX.Y.Z`, with rules for major, minor, and patch increments to indicate compatibility changes.
- **Pseudo-version**: A pre-release version encoding a specific revision (timestamp + hash) used to reference commits or tags without a stable semantic tag.
- **Major Version Suffix**: A `/vN` suffix in the module path required from major version 2 onwards to handle incompatibility and diamond dependencies.
- **Build Metadata**: Optional suffixes like `+meta`, `+incompatible`, or `+dirty` appended to versions, largely ignored for comparison purposes except for specific legacy cases.
- **GOPROXY**: An environment variable controlling the list of module proxy servers the `go` command contacts to resolve dependencies.

## Procedures And API Details

- **Module Identification**: The `go` command looks for a `go.mod` file in the directory where the `go` command is invoked; this defines the main module and its root.
- **Version Resolution**: Commands like `go get`, `go mod tidy`, or `go list -m` can accept commit hashes (e.g., `@daa7c041`) or branch names, which the tool automatically converts into pseudo-versions or tagged versions.
- **Dependency Lookup**: When loading a package path, the `go` command searches the build list for modules whose paths are prefixes of the package path, checking directories for `.go` files to confirm package existence.
- **Proxy Usage**: The `go` command queries proxies in order (e.g., `https://corp.example.com`, then `https://proxy.golang.org`) requesting the latest version of module path prefixes until a match is found or an error occurs.

## Nuance Or Contradictions

- **Legacy Incompatibility**: Modules released at v2+ before adopting the module system are marked with `+incompatible` (e.g., `v2.0.0+incompatible`) to distinguish them from modern module paths without implying a specific version constraint.
- **Pseudo-version Sorting**: Pseudo-versions must sort higher than their base version but lower than subsequent tagged versions; the `go` command enforces this by verifying that the timestamp and revision match actual repository data, preventing version flooding or manipulation.
- **Path Restrictions**: Module paths cannot contain dots in the first element to avoid confusion with package paths, and specific names like `example` and `test` are reserved for user-defined modules to prevent collisions with standard library expectations.

## Candidate Wiki Hints

- Go Modules Reference
- Semantic Versioning in Go
- Pseudo-versions and Commit Hashes
- Major Version Suffixes (v2+)
- Module Path Resolution Logic

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
Lines 416–859 of raw/web/corpus-2026-05-18/032-go-modules-reference.md cover the `go.mod` file structure, lexical elements, directives (module, go, toolchain, godebug, require, tool, ignore, exclude, replace), versioning rules, and compatibility notes for non-module repositories.

Local Summary
This section explains how `go.mod` files are structured, including allowed keywords (`module`, `go`, `require`, `replace`, `exclude`, `retract`), lexical tokens (whitespace, comments, punctuation, identifiers, strings), and the grammar using EBNF. It details each directive’s purpose, syntax, and behavior across Go versions, especially changes introduced in Go 1.14–1.24.

Key Claims
- A `go.mod` file must contain exactly one `module` directive.
- The `go` directive sets the minimum required Go version; since Go 1.21 it is mandatory.
- Since Go 1.17, indirect dependencies are recorded in a separate block and enable module graph pruning.
- Since Go 1.24, the `tool` directive adds packages as module-level tool dependencies.
- `ignore` directives prevent directories from being matched by package patterns.
- `exclude` directives prevent specific module versions from being loaded; behavior changed in Go 1.16 to avoid non-deterministic version selection.
- Non-canonical versions (without leading `v`) are allowed only in the main module’s `go.mod` and are automatically replaced with canonical versions when possible.

Entities And Concepts
- **Directives**: `module`, `go`, `toolchain`, `godebug`, `require`, `tool`, `ignore`, `exclude`, `replace`, `retract`.
- **Tokens**: whitespace, comments (`//`), punctuation (`(`, `)`, `=>`), keywords, identifiers, strings (interpreted and raw).
- **Versioning**: canonical (`v` + semantic version), non-canonical (restricted to main module).
- **Module paths**: domain-like leading element, slash-separated elements, restrictions on Windows reserved names and tilde-suffixes.
- **Indirect dependencies**: tracked separately since Go 1.17 for pruning and lazy loading.

Procedures And API Details
- `go get` upgrades/downgrades specific dependencies and updates `go.mod`.
- `go mod edit` performs low-level edits to `go.mod`.
- `golang.org/x/mod/modfile` package allows programmatic changes to `go.mod`.
- `go list -m -u` checks deprecated modules since Go 1.17.
- `go mod vendor` behavior varies with the `go` directive version (omits vendored dependency `go.mod`/`go.sum` files).

Nuance Or Contradictions
- Before Go 1.21, the `go` directive was advisory; now it is mandatory and affects toolchain selection.
- Indirect requirement handling differs between Go <1.16 and ≥1.17.
- `exclude` directives only apply in the main module’s `go.mod`; prior to Go 1.16 they caused non-deterministic version fallbacks.
- Deprecation applies to all minor versions of a major release; individual minor/patch versions should use `retract`.

Candidate Wiki Hints
- **Page**: `go-mod-directives` — summarize each directive with syntax, purpose, and version-specific behavior.
- **Page**: `go-mod-versioning` — canonical vs non-canonical versions, path restrictions, and toolchain integration.
- **Page**: `indirect-dependencies-in-go-mod` — explain the separate block for indirect dependencies since Go 1.17.
- **Page**: `deprecation-and-retraction-in-go-mod` — distinguish deprecation messages from version retractions.

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
This chunk covers the `replace` directive syntax and semantics, the `retract` directive for version management, automatic updates via `-mod=mod`, Minimal Version Selection (MVS), and optimizations in Go 1.17+ such as module graph pruning and lazy loading.

Local Summary
The document explains how to manipulate the module graph using directives: replacing local or remote modules, retracting versions to prevent upgrades, and updating `go.mod` files automatically. It details MVS, which deterministically selects the minimum required set of versions. Finally, it describes Go 1.17+ features that optimize the build process by pruning unnecessary transitive dependencies and loading the full graph only when needed.

Key Claims
- A `replace` directive substitutes a specific module version with another (local or remote) path; if no left-side version is specified, all versions are replaced.
- `retract` directives mark versions as problematic, preventing automatic upgrades but keeping them accessible for existing builds.
- The `-mod=mod` flag enables the Go command to automatically rewrite non-canonical versions, respect exclusions, remove redundant requirements, and reformat `go.mod`.
- Minimal Version Selection (MVS) computes a build list by traversing the module graph from main modules, tracking the highest required version for each dependency.
- In Go 1.17+, the module graph is pruned to exclude transitive dependencies of modules specifying `go 1.17+`, unless those dependencies are also required by older Go versions.

Entities And Concepts
- **Replace Directive**: Substitutes a module path/version with a local file path or another module path/version.
- **Retract Directive**: Marks specific version ranges as unavailable for new upgrades.
- **Minimal Version Selection (MVS)**: Algorithm to determine the set of module versions used for building.
- **Module Graph Pruning**: Optimization in Go 1.17+ to exclude transitive dependencies of modern modules from the graph unless needed by older modules.
- **Lazy Module Loading**: Strategy to load the full module graph only when required packages are not found among immediate requirements.

Procedures And API Details
- **Syntax for Replace**: `replace <module> [<version>] => <path|module> [version]`
  - Example: `replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5`
  - Grouping multiple directives is allowed within parentheses.
- **Syntax for Retract**: `retract <version>` or `retract [<lower>, <upper>]`
  - Example: `retract [v1.0.0, v1.9.9]`
  - Comments explaining the rationale can be added above or below the directive.
- **Automatic Update Command**: `go get -mod=mod` or similar commands with `-mod=mod` flag to rewrite `go.mod`.
- **Inspecting Build List**: `go list -m all` displays the selected versions determined by MVS.

Nuance Or Contradictions
- **Replace vs Require**: A `replace` directive alone does not add a module to the dependency graph; a corresponding `require` directive is still necessary to include the replaced module in the build list.
- **Retraction Visibility**: Retracted versions are hidden from `go list -m -versions` unless the `-retracted` flag is used.
- **Pruning Behavior**: Modules with pruned requirements still appear in `go list -m all`, but their packages cannot be directly imported into `go build` or `go test` without promoting them to explicit dependencies via `go get`.

Candidate Wiki Hints
- **Page: Go Modules Replace Directive** (Concept: Modifying the module graph)
- **Page: Go Modules Retract Directive** (Concept: Version lifecycle management)
- **Page: Minimal Version Selection in Go** (Concept: Dependency resolution algorithm)
- **Page: Module Graph Pruning and Lazy Loading** (Concept: Go 1.17+ optimizations)

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## Chunk Context
This chunk covers the `go.work` file format, its directives, and compatibility handling for pre-module repositories. It details lexical elements, grammar rules (EBNF), and specific behaviors regarding major version suffixes (`+incompatible`) and minimal module compatibility in GOPATH mode.

## Local Summary
The document explains how Go workspaces are defined using `go.work` files, which include directives for toolchain versioning, module inclusion (`use`), and dependency replacement (`replace`). It describes the syntax rules, including the mandatory `go` directive specifying the toolchain version. The text also addresses legacy compatibility, specifically how the Go command handles repositories that lack a `go.mod` file but have tags with major versions 2 or higher (appending `+incompatible` suffix). Finally, it outlines "minimal module compatibility" logic introduced in Go 1.11 to allow seamless imports between modules and GOPATH directories even when major version subdirectories are not strictly used.

## Key Claims
- A valid `go.work` file must contain exactly one `go` directive specifying the toolchain version.
- The `use` directive adds a module directory to the workspace; it does not recursively include submodules.
- If a repository root lacks a `go.mod` file but has a major version 2+ tag, the Go command synthesizes a synthetic `go.mod` and may append a `+incompatible` suffix to versions in that range.
- Minimal module compatibility allows packages imported via `$modpath/$vn/$dir` (where `$vn` is a major version suffix) to resolve to GOPATH locations without the suffix, provided specific conditions are met.

## Entities And Concepts
- **go.work**: UTF-8 text file defining a workspace.
- **Directive**: A line in `go.work` consisting of a keyword and arguments (e.g., `use`, `replace`).
- **Module-aware mode**: Build mode using `go.mod` files to resolve dependencies.
- **GOPATH mode**: Legacy build mode ignoring modules, looking in `$GOPATH/src`.
- **+incompatible**: Suffix added by the Go command for versions with major version 2+ lacking a `go.mod` file.
- **Minimal module compatibility**: Logic allowing mixed usage of module paths and GOPATH directories during imports.

## Procedures And API Details
- **Syntax Rule**: A `go.work` file follows Extended Backus-Naur Form (EBNF) where `GoWork = { Directive }`.
- **Directive Types**: `GoDirective`, `ToolchainDirective`, `UseDirective`, `ReplaceDirective`.
- **Example Command**: `go work init` creates a new file; `go work use` adds modules.
- **Compatibility Check**: When resolving an import `$modpath/$vn/$dir` in GOPATH mode:
  - If `$GOPATH[d]/src/$modpath/go.mod` exists and declares `$modpath/$vn`.
  - And the directory `$GOPATH[d]/src/$modpath/$vn/$dir` does not exist.
  - Then resolve to `$GOPATH[d]/src/$modpath/$dir`.

## Nuance Or Contradictions
- **Committing `go.work`**: Generally inadvisable because it can override parent workspace definitions or cause CI systems to test incorrect dependency versions, though exceptions exist for repositories developed exclusively together.
- **Version Tags vs. Versions**: The `+incompatible` suffix appears in Go command usage but should not appear on repository tags (e.g., a tag named `v4.1.2+incompatible` is ignored).
- **GOPATH Directory Structure**: A module with a major version suffix does not necessarily need to be developed in a subdirectory matching that suffix; the Go command supports resolving imports regardless of this structure if minimal compatibility rules apply.

## Candidate Wiki Hints
- **Topic: Go Workspace Files** – Overview of `go.work` syntax and directives.
- **Topic: Legacy Module Compatibility** – Handling pre-module repositories with major version 2+ tags.
- **Topic: Minimal Module Compatibility Rules** – Conditions for resolving imports across GOPATH and module paths.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
- Source File: `raw/web/corpus-2026-05-18/032-go-modules-reference.md`
- Chunk Index: 6 of 13
- Line Range: 1666–1685
- Heading Path: Go Modules Reference > Retrieved Text

Local Summary
This section explains the behavior of build commands and module management tools when vendor directories are enabled. It clarifies that while `go build` and `go test` use vendored packages, other commands like `go mod download`, `go mod tidy`, and `go get` continue to interact with the network and local cache regardless of vendoring status. It also notes restrictions on using vendor directories outside the main module's root and their exclusion from zip files.

Key Claims
- When vendoring is enabled, `go build` and `go test` load packages from the vendor directory instead of accessing the network or local module cache.
- The `go list -m` command only prints information about modules listed in `go.mod`.
- Commands such as `go mod download`, `go mod tidy`, and `go get` do not behave differently when vendoring is enabled; they still download modules and access the module cache.
- Unlike GOPATH mode, the `go` command ignores vendor directories in locations other than the main module’s root directory.
- The `go` command does not include vendor directories from other modules when building module zip files (with references to known bugs #31562 and #37397).

Entities And Concepts
- Vendoring: A mechanism to include dependency code directly in the project source tree.
- Vendor Directory: The directory (`vendor/`) containing copied dependencies used by build tools when vendoring is enabled.
- Go Modules Reference: Documentation section covering module management behaviors.
- `go list -m`: Command to list modules, restricted to those in `go.mod`.
- Module Zip Files: Bundled archives of modules, excluding vendor directories from sub-modules.
- Known Bugs #31562 and #37397: Issues related to vendor directory handling in zip files.

Procedures And API Details
- Vendoring Enabled Build Flow:
  1. Execute `go build` or `go test`.
  2. If vendoring is enabled, the tool loads packages from the `vendor/` directory at the project root.
  3. Network access and local module cache are bypassed for these packages.
- Module Management Under Vendoring:
  - Running `go mod download`, `go mod tidy`, or `go get` does not change behavior; they still interact with the network and cache.
- Vendor Directory Usage Restrictions:
  - Only vendor directories at the main module’s root are respected by the `go` command.
  - Vendor directories in other modules are ignored during builds and excluded from zip files.

Nuance Or Contradictions
- The documentation highlights a discrepancy between GOPATH mode vendoring (where any vendor directory might be used) and Go Modules mode (which restricts usage to the main module’s root).
- There is an acknowledged inconsistency regarding zip file generation, where vendor directories are omitted despite potential relevance, noted via bug reports #31562 and #37397.

Candidate Wiki Hints
- Page: `Go Modules/Vendoring Behavior`
  - Summary: Explains how vendoring affects build commands versus module management commands in Go Modules mode.
  - Content Focus: Differences between GOPATH and Modules vendoring, scope of vendor directory usage, and known limitations with zip files.

## chunk-07

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

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
## Chunk Context
This chunk details the `go mod edit` command for editing `go.mod` files, provides its JSON output schema via `-json`, and describes `go mod graph`, `go mod init`, `go mod tidy`, `go mod vendor`, `go mod verify`, and `go mod why`. It concludes with the `go version -m` command for inspecting executable build versions.

## Local Summary
The section begins by explaining that `go mod edit -json` outputs the `go.mod` file as a JSON structure representing the module, Go version, requirements, exclusions, replacements, retractions, and tools. The text defines the associated Go types (`Module`, `GoMod`, `ModPath`, `Require`, `Replace`, `Retract`, `Tool`). It notes that this schema describes only the `go.mod` file itself, not indirect modules, directing users to `go list -m -json all` for the full set. Subsequent subsections cover `go mod graph` (printing the module requirement graph), `go mod init` (creating a new module), `go mod tidy` (syncing go.mod with imports), `go mod vendor` (copying dependencies to a vendor directory), `go mod verify` (checking integrity of cached modules), and `go mod why` (showing import paths). The chunk ends with `go version -m`, which prints the Go version used to build an executable, optionally including module versions.

## Key Claims
- `go mod edit` reads only one `go.mod` file and does not look up information about other modules.
- `go mod edit -json` outputs a JSON structure corresponding to specific Go types that describe the `go.mod` file itself, excluding indirect modules.
- For the full set of modules available to a build (including indirect ones), use `go list -m -json all`.
- `go mod graph` prints the module requirement graph with replacements applied in text form.
- `go mod init` creates a new `go.mod` file in the current directory if it does not already exist.
- `go mod tidy` ensures `go.mod` matches source code, adding missing requirements and removing unused ones.
- `go mod vendor` constructs a `vendor` directory containing copies of needed packages and generates `vendor/modules.txt`.
- `go mod verify` checks that dependencies in the module cache have not been modified since download.
- `go mod why` shows a shortest path in the import graph from the main module to listed packages or modules.
- `go version -m` prints the Go version and module versions used to build a specific executable.

## Entities And Concepts
- **Commands**: `go mod edit`, `go list`, `go mod graph`, `go mod init`, `go mod tidy`, `go mod vendor`, `go mod verify`, `go mod why`, `go version`.
- **Files**: `go.mod`, `vendor/modules.txt`.
- **Directives/Flags**: `-json`, `-fmt`, `-print`, `-require`, `-droprequire`, `-exclude`, `-dropexclude`, `-replace`, `-dropreplace`, `-retract`, `-dropretract`, `-tool`, `-droptool`, `-e`, `-v`, `-x`, `-diff`, `-go`, `-compat`, `-o`, `-m`, `-vendor`.
- **Types**: `Module`, `GoMod`, `ModPath`, `Require`, `Replace`, `Retract`, `Tool`.
- **Build Tags**: `ignore`.

## Procedures And API Details
- **Printing JSON Representation of go.mod**:
  - Command: `$ go mod edit -json`
  - Output corresponds to Go types:
    ```go
    type Module struct { Path string; Version string }
    type GoMod struct { Module ModPath; Go string; Require []Require; Exclude []Module; Replace []Replace; Retract []Retract }
    type ModPath struct { Path string; Deprecated string }
    type Require struct { Path string; Version string; Indirect bool }
    type Replace struct { Old Module; New Module }
    type Retract struct { Low string; High string; Rationale string }
    type Tool struct { Path string }
    ```
- **Printing Go version used to build an executable**:
  - Command: `$ go version ~/go/bin/gopls`
- **Printing Go version and module versions used to build an executable**:
  - Command: `$ go version -m ~/go/bin/gopls`

## Nuance Or Contradictions
- `go mod edit` does not look up information about other modules; it only reads and writes the specified target file.
- The JSON output from `go mod edit -json` describes only the `go.mod` file itself, not other modules referred to indirectly.
- `go mod tidy` acts as if all build tags are enabled except for the `ignore` tag.
- `go mod vendor` removes the existing `vendor` directory before re-constructing it.
- After Go 1.16 in modules declaring `go 1.16` or higher, the `-vendor` flag of `go mod why` has no effect because the meaning of `all` changed to match the set of packages matched by `go mod vendor`.

## Candidate Wiki Hints
- **Page**: Go Module Editing Commands (`go mod edit`)
  - Covers editing flags (`-module`, `-require`, `-replace`, etc.) and output control (`-json`, `-fmt`).
- **Page**: Inspecting Module Graphs and Versions (`go mod graph`, `go version -m`)
  - Explains `go mod graph` output format and `go version -m` usage for build inspection.

## chunk-09

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

## chunk-10

---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
## Chunk Context
This chunk details the internal procedures of the `go` command when building executables, specifically focusing on module resolution, downloading `.mod` and `.zip` files, verifying checksums, and handling modules served directly from version control repositories (Git, SVN, etc.). It covers the protocol for finding repository URLs via `?go-get=1`, mapping semantic and pseudo-versions to commits, and locating `go.mod` files within complex repository structures.

## Local Summary
The document explains that `go build` first computes a build list using Minimal Version Selection (MVS), then loads required packages by downloading `.mod` and `.zip` files from proxies or version control systems. It describes the specific HTTP requests made to proxies (e.g., `$module/@v/$version.info`) and how the command verifies file integrity using hashes against `go.sum`. The text further elaborates on "direct mode," where modules are fetched directly from Git/SVN repositories, requiring a `<meta name="go-import">` tag to resolve the repository URL. Finally, it details how version tags, pseudo-versions, and branch names are mapped to specific commits and how the command locates the correct `go.mod` directory within a repository root or subdirectory.

## Key Claims
- The `go build` procedure involves computing a build list via MVS, reading packages, finding missing modules, and building.
- `.mod` files are downloaded using `$module/@v/$version.mod` requests; `.zip` files use `$module/@v/$version.zip`.
- Checksums for downloaded files are verified against the main module's `go.sum`; mismatches trigger security errors unless `GOSUMDB` is set to off.
- Direct mode allows downloading modules from version control repositories (Git, Mercurial, etc.) when a proxy is unavailable or for private repos.
- Repository resolution relies on an HTML response containing `<meta name="go-import" content="root-path vcs repo-url [subdirectory]">`.
- Version tags must match the module path's major version suffix; tags for modules in subdirectories include the subdirectory prefix (e.g., `gopls/v0.4.0`).
- Pseudo-versions (e.g., `v1.3.2-0.20191109021931-daa7c04131f5`) encode a commit hash and timestamp to ensure reproducible builds.

## Entities And Concepts
- **GOPROXY protocol**: Requests sent to proxy servers for module metadata and source code.
- **Minimal Version Selection (MVS)**: Algorithm used to select the latest compatible version of modules in the build list.
- **Direct mode**: Fetching modules directly from a VCS repository instead of a proxy.
- **go-import meta tag**: HTML tag used to signal a repository's root path, VCS type, and URL.
- **Pseudo-version**: A specific revision encoded with a timestamp and commit hash prefix (e.g., `v1.3.2-0.20191109021931-daa7c04131f5`).
- **Semantic version tags**: Tags like `v1.2.3` indicating specific commits for a module.
- **GOPRIVATE / GONOPROXY**: Environment variables to configure the go command to download from source repositories directly.

## Procedures And API Details
- **Downloading `.mod` files**: The command sends `$module/@v/$version.mod`. Example: `curl https://proxy.golang.org/golang.org/x/mod/@v/v0.2.0.mod`.
- **Downloading `.zip` files**: The command sends `$module/@v/$version.zip`. Example: `curl -O https://proxy.golang.org/golang.org/x/mod/@v/v0.2.0.zip`.
- **Fetching version list**: Request `$module/@v/list` returns available versions (e.g., `v0.1.0`, `v0.2.0`).
- **Fetching version info**: Request `$module/@v/$version.info` returns JSON metadata (`{"Version":"...", "Time":"..."}`).
- **Resolving repository URL**: Send `GET https://<module-path>?go-get=1`. Parse `<meta name="go-import" content="root-path vcs repo-url [subdirectory]">`.
- **Mapping branch to version**: Use `go get <path>@<branch>` (e.g., `go get example.com/mod@master`). The command converts the branch/tag into a canonical version for MVS.

## Nuance Or Contradictions
- Synthetic `go.mod` files: If a project lacks a `go.mod`, the proxy serves a synthetic file containing only a module directive.
- Version list authentication: Unlike `.mod` and `.zip` files, version lists (`.info`) and metadata are not authenticated and may change over time.
- Subdirectory support: `<meta>` tags providing a subdirectory are only recognized by Go 1.25 and later; earlier versions ignore them and fail resolution if the module isn't in the root.
- GOPATH mode limitations: Modules served directly from a proxy cannot be downloaded with `go get` in GOPATH mode.

## Candidate Wiki Hints
- **Module Resolution Lifecycle**: A step-by-step guide on how `go build` resolves dependencies, handles missing modules, and downloads artifacts.
- **GOPROXY Protocol Specification**: Details on the HTTP requests (`@v`, `@latest`) used to query module proxies.
- **Direct Mode Configuration**: How to set up a proxy or repository to serve modules directly using `?go-get=1` meta tags.
- **Version Tagging Best Practices**: Guidelines for naming semantic version tags and pseudo-versions in repositories compatible with Go modules.

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
This chunk covers the internal mechanics of Go modules, specifically focusing on how `go` creates and validates `.zip` files for module distribution. It details constraints on file paths, sizes, and content (e.g., ignoring vendor directories or symlinks), special handling for LICENSE files, and security controls via the `GOVCS` environment variable to restrict which version control systems are used when downloading from public vs. private servers.

Local Summary
The `go` command packages module contents into `.zip` files after locating the module root. These archives are authenticated before extraction. The chunk outlines strict rules for what goes into a module zip (no vendor dirs, no nested go.mod dirs, specific path prefixes) and how to handle license files across subdirectories. It also explains the `GOVCS` variable, which balances the need for decentralized package hosting with security by restricting default VCS usage to Git and Mercurial for public modules while allowing others for private ones via proxies.

Key Claims
- Module `.zip` files are authenticated before extraction into the cache, similar to proxy downloads.
- Vendor directories and nested modules (subdirectories with `go.mod`) are excluded from module zip files.
- If a subdirectory module lacks a `LICENSE` file, `go` copies one from the repository root if present.
- By default, `go` uses Git and Mercurial for public servers; other VCS tools are reserved for private modules or proxied downloads.
- The `GOVCS` environment variable allows explicit control over allowed version control systems per module path pattern.
- Module zip files are limited to 500 MiB total size (compressed and uncompressed).

Entities And Concepts
- `go mod download`: Command to fetch and extract modules.
- `GOPROXY`: Environment variable listing proxy URLs for module downloads.
- `GOVCS`: Environment variable controlling allowed version control systems per path pattern.
- `GONOPROXY`, `GONOSUMDB`, `GOPRIVATE`: Variables configuring private module access and security checks.
- `.zip` file format constraints (path prefixes, size limits, ignored metadata).
- Version Control Systems: Git, Mercurial, Bazaar, Fossil, Subversion.

Procedures And API Details
- **Setting GOVCS**: Use a comma-separated list of `pattern:vcslist` rules.
  - Example: `GOVCS=github.com:git,evil.com:off,*:git|hg`
  - Default logic: `public:git|hg,private:all`.
- **Disabling VCS usage**: Set `GOVCS=*:off`.
- **Enabling all VCS**: Set `GOVCS=*:all`.
- **Configuring private proxy**:
  - Set `GOPROXY=https://proxy.corp.example.com` and `GONOSUMDB=corp.example.com`.
- **Direct access to private modules**:
  - Set `GOPRIVATE=corp.example.com`.
  - Ensure repository URLs use HTTPS or have a VCS suffix (e.g., `.git`).

Nuance Or Contradictions
- While `go` ignores symbolic links in zip files, authors can manually copy license files into subdirectory modules if the root lacks one.
- The default restriction to Git and Mercurial for public servers is a security measure; other VCS tools are allowed for private modules or when using the proxy mirror (proxy.golang.org).
- `GOVCS` patterns match leading elements of module/import paths; the earliest matching pattern applies, even if later ones also match.

Candidate Wiki Hints
- Page: Go Modules Security and Version Control Configuration
  - Focus on `GOVCS`, `GOPRIVATE`, and managing private module access.
- Page: Module Zip File Constraints
  - Document size limits, path prefixes, and ignored file types (vendor, symlinks).

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

Chunk Context
The chunk details the privacy and security architecture of Go module downloads. It covers HTTP authentication configuration, proxy settings (GOPROXY, GOPRIVATE), checksum database interactions (GOSUMDB, GONOSUMDB), module cache structure, and the cryptographic verification process involving go.sum files and the global checksum database.

Local Summary
This section explains how `go` handles network traffic regarding private modules and security verification. It distinguishes between public proxies (Google's) and private proxies, explaining fallback mechanisms when 404/410 errors occur. The text describes the module cache layout, file permissions, and the role of `go mod verify`. Finally, it details the `go.sum` format and the Merkle tree-based checksum database protocol used to ensure module integrity without trusting individual origin servers.

Key Claims
- The default `GOPROXY` setting is `https://proxy.golang.org,direct`, prioritizing Google's public proxy before falling back to direct version control system access.
- `GOPRIVATE` and `GONOPROXY` use glob patterns to exclude private modules from any proxy requests, forcing direct fetches from version control repositories.
- The module cache resides at `$GOPATH/pkg/mod` by default but can be moved via the `GOMODCACHE` environment variable.
- Module files in the cache are stored with read-only permissions to prevent accidental modification; deletion requires `go clean -modcache`.
- Hashes for downloaded modules are verified against the main module's `go.sum` file before caching. If `go.sum` is missing, the global checksum database (sum.golang.org) is consulted.
- The checksum database uses a Transparent Log (Merkle Tree) structure backed by Trillian to allow independent auditors to verify data integrity.

Entities And Concepts
- **GOPROXY**: Environment variable controlling which module proxy servers are used. Default: `https://proxy.golang.org,direct`.
- **GOPRIVATE / GONOPROXY**: Variables setting glob patterns for private modules that bypass proxies and fetch directly from version control.
- **GOSUMDB**: Variable setting the checksum database name/URL (default: `sum.golang.org`). Can be set to `off` to disable verification entirely.
- **GONOSUMDB**: Variable specifying module prefixes that should not request hashes from the checksum database.
- **Module Cache**: Directory (`$GOPATH/pkg/mod`) storing downloaded modules. Contains extracted contents, proxy caches, and VCS clones.
- **go.sum**: File containing cryptographic hashes (SHA-256) for dependencies to ensure integrity. Format: `module path version hash`.
- **Checksum Database**: Global service (`sum.golang.org`) providing signed logs of module hashes to prevent tampering by untrusted proxies.

Procedures And API Details
- **Disabling Proxy Access**: Set `GOPRIVATE=*.corp.example.com` or use `GONOPROXY` for specific patterns.
- **Configuring Private Proxy**: Use a trusted proxy with fallback: `GOPROXY=https://proxy.corp.example.com,https://proxy.golang.org` combined with `GONOSUMDB`.
- **Verifying Cache Integrity**: Run `go mod verify` to ensure extracted module contents match hashes in `go.sum`.
- **Managing Cache Size/Permissions**: Use `go clean -modcache` to clear. Use `-modcacherw` flag if writable permissions are required (increases risk).
- **Checksum Database Lookup**: Query `/lookup/$module@$version` to get record data and tree description for inclusion proofs.

Nuance Or Contradictions
- **Hash Verification Scope**: The checksum database cannot compute checksums for non-public modules; verification relies on `go.sum` or is skipped if `GOSUMDB=off`.
- **Fallback Behavior**: If a private proxy returns 404/410, the command falls back to the public proxy. If it returns any other error code, no fallback occurs.
- **Case Sensitivity Handling**: Module paths are case-encoded (e.g., `example.com/M` becomes `example.com/!m`) in cache paths to handle case-insensitive file systems correctly.

Candidate Wiki Hints
- How GOPROXY and GOPRIVATE interact with private module repositories.
- Understanding the security implications of the `go.sum` file and checksum database.
- Structure of the Go module cache directory (`cache/download/`, `cache/vcs/`).
- Protocol details for interacting with the global checksum database (Merkle Tree).

## chunk-13

---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
# Chunk Context
The chunk details the module-related environment variables controlling `go` command behavior, including cache paths, proxy configurations, and security settings. It concludes with a glossary defining core module terminology such as "main module," "direct dependency," and "minimal version selection."

# Local Summary
This section lists environment variables (GO111MODULE, GOMODCACHE, GOPROXY, etc.) that configure how the `go` command handles modules, proxies, and checksums. It also provides a glossary of terms used throughout the Go module system documentation.

# Key Claims
- The `go` command's module behavior is configurable via environment variables.
- `GO111MODULE` controls whether to run in module-aware mode or GOPATH mode (`off`, `on`, `auto`).
- `GOPROXY` defaults to `https://proxy.golang.org,direct`.
- `GOSUMDB` defaults to `sum.golang.org`.
- `GOVCS` defaults to using `git` and `hg` for public modules.
- Minimal version selection (MVS) determines the build list of module versions.

# Entities And Concepts
- **Environment Variables**: GO111MODULE, GOMODCACHE, GOINSECURE, GONOPROXY, GONOSUMDB, GOPATH, GOPRIVATE, GOPROXY, GOSUMDB, GOVCS, GOWORK.
- **Module Terms**: main module, direct dependency, indirect dependency, build list, canonical version, selected version, pseudo-version, release version, pre-release version.
- **Modes**: module-aware mode, GOPATH mode, workspace mode, single-module mode.
- **Proxies/Security**: module proxy, insecure download, checksum database validation.

# Procedures And API Details
- **Setting GOPROXY**: Use commas (`,`) for fallback on 404/410 errors; use pipes (`|`) to fall back on any error (including timeouts).
- **Configuring GOSUMDB**: Format is `database-name` or `database-name+<publickey> <url>`. Set to `off` to disable checksum verification.
- **Disabling Module Proxy**: Set `GOPROXY=direct` or use `GONOPROXY`/`GOPRIVATE` patterns.
- **Enabling Single-Module Mode**: Use `GOWORK=off go build .` when no `go.work` file is desired.

# Nuance Or Contradictions
- `GO111MODULE=auto` was the default in Go 1.15 and lower, whereas it defaults to `on` (or unset) in newer versions unless explicitly set to `off`.
- `GOPRIVATE` serves as a default value for both `GONOPROXY` and `GONOSUMDB`, meaning private modules are excluded from proxy checks by default if not overridden.
- Setting `GOSUMDB=off` or using `-insecure` bypasses checksum validation entirely, accepting all unrecognized modules without security guarantees.

# Candidate Wiki Hints
- **Topic: Module Environment Variables**: A reference page listing all module-related env vars with their defaults and usage examples.
- **Topic: Minimal Version Selection (MVS)**: Explanation of how MVS determines the build list and its impact on transitive dependencies.
- **Topic: Proxy Configuration**: Guide on configuring `GOPROXY`, handling errors, and using local file proxies.

