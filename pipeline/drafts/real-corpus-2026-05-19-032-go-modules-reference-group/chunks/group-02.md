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
