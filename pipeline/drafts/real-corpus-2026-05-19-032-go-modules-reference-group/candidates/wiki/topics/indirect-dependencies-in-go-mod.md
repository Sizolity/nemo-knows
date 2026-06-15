---
title: Indirect Dependencies In Go Mod
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Indirect Dependencies In Go Mod

## Overview

In Go's module system, **indirect dependencies** are packages required by a direct dependency but not explicitly listed in the `require` block of the current module. Since Go 1.17, these dependencies are managed separately from direct ones to support more efficient build workflows.

## Management and Storage

Direct dependencies are declared in the `require` block of the [`go.mod`](go-mod-directives) file. Indirect dependencies are resolved during the build process but are not automatically added to this list. Instead, they are recorded in a separate section within `go.mod`. This separation enables the Go toolchain to prune unused indirect dependencies and supports lazy loading mechanisms.

## Resolution Process

When the `go` command resolves dependencies:
1. It consults the `require` block for direct dependencies.
2. It analyzes the dependency graph of those packages to identify any indirect requirements.
3. Indirect dependencies are resolved using the **Minimal Version Selection (MVS)** algorithm to ensure deterministic versioning across builds.

## Build Behavior

The presence of indirect dependencies does not affect standard build commands like `go build` or `go test` as long as the necessary packages are reachable. However, commands such as `go mod tidy` and `go get` continue to interact with the network or cache regardless of vendoring status. When vendoring is enabled, only vendor directories at the main module's root are respected; indirect dependencies outside this scope may be ignored during local builds.

## Maintenance

To keep the dependency graph clean:
- Use `go mod tidy` to update both direct and indirect dependencies in `go.mod`.
- Manually remove unused indirect dependencies if desired, though this is rarely necessary for standard projects.
- Monitor [`go.sum`](module-security-and-cache) files to ensure integrity checks pass for all resolved versions, including indirect ones.

## Security Considerations

All dependencies, whether direct or indirect, are verified against checksums in the global database (`sum.golang.org`) to prevent tampering by untrusted proxies or origin servers. The [`GOPROXY`](module-environment-variables) environment variable controls how metadata is fetched during resolution.
