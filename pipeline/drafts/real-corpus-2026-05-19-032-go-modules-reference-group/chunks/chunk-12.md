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
