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
