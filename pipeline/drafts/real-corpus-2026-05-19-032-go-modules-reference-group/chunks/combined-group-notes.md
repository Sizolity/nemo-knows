## group-01

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

## group-02

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Group Context

This group of notes covers the `go` command's comprehensive module management capabilities, focusing on dependency resolution, version control interactions, and build verification. The content spans from basic commands like `go get`, `go install`, and `go mod edit` to advanced topics such as workspace management (`go work`), proxy protocols, direct repository access, and security mechanisms involving checksum databases. It details the lifecycle of a module from discovery via proxies or version control systems, through resolution using Minimal Version Selection (MVS), caching, verification against `go.sum`, and installation into the module cache.

# Cross-Chunk Summary

The documentation is structured around three primary pillars: dependency management commands, build/installation mechanics, and security/proxy configuration.

**Dependency Management:** The group details how to use `go get` for managing transitive dependencies (upgrading, downgrading, removing) and `go mod edit` for manual file manipulation (replacements, exclusions, formatting). It distinguishes between updating the module graph (`go get`) and installing executables (`go install`), noting that `go install` ignores local `go.mod` when version suffixes are provided.

**Build and Resolution:** The notes explain the internal procedure of building executables: computing a build list via MVS, downloading `.mod` and `.zip` files from proxies or VCS, verifying checksums against `go.sum`, and handling pseudo-versions for reproducible builds. It covers how to query versions using semantic prefixes (`@latest`, `@upgrade`) and revision identifiers (`@master`).

**Proxies and Direct Access:** A significant portion is dedicated to the module proxy protocol (endpoints like `@v/list`, `@v/$version.mod`), fallback strategies for private modules, and the use of environment variables (`GOPROXY`, `GOPRIVATE`, `GONOPROXY`) to control network traffic. It also covers direct access from Git/SVN repositories using `?go-get=1` meta tags.

**Workspace and Security:** The group introduces workspace management (`go work init/edit/sync`) for multi-module projects, ensuring consistent builds across a tree of modules. Security is addressed through the checksum database (Merkle Tree/Trillian), `go.sum` verification, and controls on version control systems via `GOVCS`.

# Repeated Or Central Claims

- **Minimal Version Selection (MVS):** This algorithm is central to Go's module system, used by `go build`, `go install`, and `go work sync` to compute a consistent set of versions that satisfy all requirements while preferring the latest compatible version.
- **Version Query Suffixes:** The syntax `<module>@<query>` allows precise control over which version is selected. Common queries include `@latest` (prefer release), `@upgrade` (preserve higher if applicable), `@patch`, and `@revision` (commit hash/branch). `@none` removes the dependency.
- **Transitive Dependency Updates:** Commands like `go get -u` or `go work sync` trigger cascading updates to transitive dependencies, ensuring that if a direct dependency upgrades, its own requirements are also resolved to compatible versions.
- **Checksum Verification:** Module integrity is verified by comparing SHA-256 hashes in the local `go.sum` file against a global checksum database (`sum.golang.org`). This prevents tampering by untrusted proxies or origin servers.
- **Proxy Fallback Logic:** The default behavior prioritizes public proxies (e.g., `proxy.golang.org`) but falls back to direct access from version control systems if the proxy returns 404/410 errors, ensuring modules remain accessible even when proxies are down or misconfigured.

# Important Local Details

- **Command Syntax Nuances:**
  - `go get` updates `go.mod` and builds; `-d` flag is deprecated.
  - `go install` ignores local `go.mod` if version suffixes are present, building in module-aware mode.
  - `go mod edit` does not look up external modules; it only reads/writes the target file. It supports `-print` to output changes without saving and `-json` for schema inspection.
  - `go list -m` can list dependencies with flags like `-u` (show upgrades), `-retracted`, and `-versions`.

- **File and Directory Structures:**
  - **Module Cache:** Located at `$GOPATH/pkg/mod` (or `$GOMODCACHE`). Contains extracted module contents, proxy caches, and VCS clones. Files are read-only by default.
  - **Vendor Directory:** Created by `go mod vendor`, containing copies of dependencies and a `modules.txt` file. It is removed before reconstruction by subsequent runs of `go mod vendor`.
  - **Zip Constraints:** Module zip files cannot contain vendor directories, nested `go.mod` directories, or symbolic links. Size is limited to 500 MiB.

- **Environment Variables:**
  - `GOPROXY`: Comma-separated list of proxy URLs (default: `https://proxy.golang.org,direct`).
  - `GOPRIVATE`: Glob patterns for modules that should not be proxied (fetch directly from VCS).
  - `GONOPROXY`: Patterns for modules to skip proxy requests entirely.
  - `GOSUMDB`: Name/URL of the checksum database (default: `sum.golang.org`). Set to `off` to disable verification.
  - `GOVCS`: Controls which VCS tools are allowed for downloading specific module paths.
  - `GOINSECURE`: Allows fetching from insecure HTTP servers (not recommended).

- **Proxy Protocol Details:**
  - Proxies must serve specific endpoints: `/@v/list` (versions), `/@v/$version.info` (JSON metadata), `/@v/$version.mod` (go.mod file), and `/@v/$version.zip` (source archive).
  - Case encoding is applied to paths (e.g., `example.com/M` -> `example.com/!m`) to handle case-insensitive filesystems.
  - Pseudo-versions are excluded from version lists but included in metadata/info files.

# Candidate Wiki Hints

- **Page: Managing Dependencies with `go get`**
  - Covers upgrading/downgrading, version suffixes (`@v`, `@master`, `@none`), and transitive updates.

- **Page: Installing Executables with `go install`**
  - Explains module-aware mode, ignoring local `go.mod`, and argument constraints (same version suffix).

- **Page: Editing `go.mod` with `go mod edit`**
  - Details flags for replacements, exclusions, Go version setting, and JSON printing.

- **Page: Workspace Management (`go work`)**
  - Explains `go work init`, `edit`, and `sync` for managing multi-module projects using MVS.

- **Page: Module Proxy Protocol Specification**
  - Defines HTTP endpoints, response formats, and fallback logic for module proxies.

- **Page: Security and Verification (`go.sum`, Checksum DB)**
  - Details the Merkle Tree protocol, `go mod verify`, and configuration of private modules.

- **Page: Direct Repository Access**
  - Covers using `?go-get=1` meta tags and direct VCS fetching when proxies are unavailable.

# Gaps Or Cautions

- **Retracted Versions:** By default, retracted versions are omitted from version lists (`go list -m -versions`). Users must explicitly request them with `-retracted` to see warnings or details about deprecated modules.
- **Case Sensitivity:** Module paths are case-encoded in requests and cache paths. If a repository uses case-insensitive filesystems (like Windows), the path `example.com/M` might be requested as `example.com/!m`. This can lead to confusion if the actual code is named differently in different environments.
- **Direct Mode Limitations:** Direct mode (fetching from Git/SVN) only works if the proxy returns a 404/410 error or if explicitly configured via `GOPRIVATE`. It does not work with standard proxies that serve modules correctly.
- **Pseudo-versions vs. Releases:** The `latest` query prefers release versions over pre-releases. Even if a pre-release is newer, it will be ignored unless explicitly requested via a revision identifier (e.g., `@master`).
- **Cache Permissions:** Module files in the cache are stored with read-only permissions. Deleting or modifying them requires `go clean -modcache` or setting writable permissions via `-modcacherw`, which increases security risks.
- **Subdirectory Support:** Recognition of `<meta>` tags specifying a subdirectory within a repository is only supported in Go 1.25 and later. Earlier versions will fail to resolve modules located in subdirectories without proper root-path configuration.

## group-03

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
# Group Context

This group of notes synthesizes documentation regarding the Go Modules Reference, specifically focusing on module management operations, environment configuration, and tooling capabilities. The content spans from initial metadata fetching to complex upgrade strategies, JSON introspection of `go.mod`, and comprehensive details on environment variables controlling module behavior (such as proxy settings, security flags, and version selection). The documentation covers procedures for upgrading/downgrading modules, managing dependencies, modifying Go versions, and formatting files, alongside a glossary of core terminology.

# Cross-Chunk Summary

The document progresses through distinct phases of module lifecycle management:
1.  **Initialization & Metadata**: Early chunks cover fetching metadata and general retrieval text.
2.  **Module Upgrades & Modifications (Chunks 07)**: Detailed strategies for upgrading specific modules, transitive dependencies, Go versions, and toolchains. Includes operations like adding/removing replace directives, ignoring `go.mod`, and formatting files.
3.  **Introspection & Reporting (Chunks 08-13)**: A significant portion is dedicated to printing information about the build environment and module versions. This includes JSON representations of `go.mod`, checking Go versions for specific executables or directories, and listing all programs in a directory.
4.  **Environment Configuration (Chunk 13)**: The final section details environment variables (`GO111MODULE`, `GOPROXY`, etc.) and security settings, concluding with a glossary of module terms.

# Repeated Or Central Claims

-   **Module Management is Extensive**: The documentation heavily emphasizes the ability to manipulate the module ecosystem via command-line flags (e.g., upgrade specific modules, downgrade dependencies, remove replace directives).
-   **Introspection Capabilities**: There is a strong emphasis on being able to inspect the build environment, specifically printing Go versions and module versions used to build executables or directories, often in JSON format.
-   **Environment Variables Drive Behavior**: A central theme is that `go` command behavior (module mode, proxy usage, security checks) is largely controlled by a specific set of environment variables with defined defaults.
-   **Version Selection Logic**: The concept of Minimal Version Selection (MVS) and canonical/selected versions is treated as a foundational mechanism for determining build lists and transitive dependencies.

# Important Local Details

-   **Environment Variable Defaults**:
    -   `GOPROXY` defaults to `https://proxy.golang.org,direct`.
    -   `GOSUMDB` defaults to `sum.golang.org`.
    -   `GOVCS` defaults to using `git` and `hg`.
-   **Proxy Configuration Syntax**:
    -   Use commas (`,`) for fallback on 404/410 errors.
    -   Use pipes (`|`) to fall back on any error, including timeouts.
-   **Security Bypasses**: Setting `GOSUMDB=off` or using `-insecure` bypasses checksum validation entirely. `GOINSECURE` allows insecure downloads.
-   **Version Defaults**: `GO111MODULE` defaults to `on` (or unset) in newer versions, whereas it was `auto` in Go 1.15 and lower.
-   **Private Module Handling**: `GOPRIVATE` acts as a default value for both `GONOPROXY` and `GONOSUMDB`, excluding private modules from proxy checks by default.
-   **Single-Module Mode**: Can be forced using `GOWORK=off go build .`.

# Candidate Wiki Hints

-   **Topic: Module Environment Variables**: A reference page listing all module-related env vars (`GO111MODULE`, `GOPROXY`, `GOSUMDB`, etc.) with their defaults and usage examples.
-   **Topic: Proxy Configuration Guide**: Instructions on configuring `GOPROXY`, handling specific HTTP errors, and setting up local file proxies or disabling the proxy entirely.
-   **Topic: Minimal Version Selection (MVS)**: An explanation of how MVS determines the build list and its impact on transitive dependencies.
-   **Topic: Module Upgrade Strategies**: A guide covering upgrades/downgrades for specific modules, transitive dependencies, Go versions, and toolchains.
-   **Topic: Build Introspection**: Methods for printing Go/module versions used to build executables or directories, including JSON output formats.

# Gaps Or Cautions

-   **Security Risks**: Disabling checksum verification (`GOSUMDB=off`) or using insecure flags accepts all unrecognized modules without security guarantees.
-   **Version Compatibility**: The default behavior of `GO111MODULE` changed between Go 1.15 and newer versions; explicit configuration may be necessary to maintain expected behavior across upgrades.
-   **Incomplete Heading Coverage**: While the chunk index lists many headings, some are nested deeply or appear as fragments (e.g., "to versions that don't require it"), suggesting potential fragmentation in the source text provided for synthesis.
-   **Limited Content in Retrieved Text Sections**: Chunks 02 through 06 and 09 through 13 are labeled with generic "Retrieved Text" or repeated introspection headings, implying the bulk of specific procedural content resides in Chunk 07 (Upgrades) and Chunk 13 (Environment), while other chunks may contain repetitive examples or less distinct structural information.

