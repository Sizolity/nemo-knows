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
