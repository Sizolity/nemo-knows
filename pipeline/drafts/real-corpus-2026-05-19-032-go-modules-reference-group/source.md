---
title: Go Modules Reference
kind: source
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## What It Is

The **Go Modules Reference** is the official documentation for Go's dependency management system. It details the lifecycle of modules from discovery and resolution to installation and security verification. The reference covers the core unit of a module (a collection of versioned packages identified by a `go.mod` file), the `go` command's capabilities for managing dependencies, build procedures, workspace management (`go.work`), and diagnostic tools for inspecting builds.

## Summary

The document establishes the workflow of fetching metadata and resolving dependencies via proxies (`GOPROXY`). It defines a module as a collection of versioned packages identified by a path declared in a `go.mod` file located in the root directory. The reference elaborates on path construction, semantic versioning rules (including major version suffixes for v2+), and pseudo-versions for commit-specific references.

A significant portion is dedicated to the `go.mod` file itself, covering its lexical grammar, mandatory directives (e.g., `module`, `go`, `toolchain`), and optional ones (e.g., `require`, `replace`, `retract`, `exclude`). The build tool uses these directives to construct a dependency graph, utilizing algorithms like Minimal Version Selection (MVS) to determine the exact set of versions needed.

The text addresses legacy compatibility, detailing how Go handles pre-module repositories and GOPATH mode alongside modern modules. It further explores workspace management via `go.work` files, which define toolchain versions and include multiple modules. Finally, it covers build-time behaviors such as vendoring, where local copies of dependencies are used instead of the network, and diagnostic commands for inspecting the final build state and executable metadata.

## Key Claims

- **Module Definition:** A module is a collection of versioned packages released together, identified by a path declared in a `go.mod` file located in the module's root directory.
- **Dependency Resolution:** The `go` command resolves package paths by searching the build list for matching module prefixes, consulting proxies (defined via `GOPROXY`), and preferring the longest matching path.
- **Versioning Strategy:** Semantic versioning (`vX.Y.Z`) is standard; pseudo-versions encode specific revisions for testing. Starting with major version 2, a suffix (e.g., `/v2`) is required to distinguish incompatible packages unless using legacy paths.
- **Directives in `go.mod`:** The file uses specific directives: `module` (defines root), `go` (sets minimum Go version, mandatory since 1.21), `toolchain`, `godebug`, `require` (direct deps), `replace` (substitution), `exclude`/`retract` (lifecycle management), and `ignore`.
- **Indirect Dependencies:** Since Go 1.17, indirect dependencies are recorded in a separate block to enable module graph pruning and lazy loading.
- **Minimal Version Selection (MVS):** An algorithm that traverses the module graph to compute the minimal set of versions required for a build, ensuring deterministic selection.
- **Vendoring Behavior:** `go build` and `go test` use vendored packages when enabled; other commands (`go mod tidy`, `go get`) continue to interact with the network/cache regardless of vendoring status. Only vendor directories at the main module's root are respected.
- **Security and Verification:** Module integrity is verified by comparing SHA-256 hashes in the local `go.sum` file against a global checksum database (`sum.golang.org`). This prevents tampering by untrusted proxies or origin servers.
- **Environment Variables Drive Behavior:** A central theme is that `go` command behavior (module mode, proxy usage, security checks) is largely controlled by environment variables with defined defaults (e.g., `GOPROXY`, `GOSUMDB`, `GOVCS`).

## Suggested Links

- [Go Modules Reference](https://go.dev/ref/mod)
