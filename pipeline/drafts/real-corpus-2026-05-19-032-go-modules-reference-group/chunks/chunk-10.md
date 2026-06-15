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
