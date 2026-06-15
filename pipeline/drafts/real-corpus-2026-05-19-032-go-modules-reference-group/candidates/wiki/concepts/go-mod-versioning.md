---
title: Go Mod Versioning
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Go Mod Versioning

Go Mod Versioning defines the lifecycle of modules within the Go language, managing dependencies from discovery and resolution to installation and security verification. A module is identified as a collection of versioned packages located in a root directory containing a `go.mod` file.

## Module Definition and Structure

The core unit of a module is defined by the path declared in its `go.mod` file. This file serves as the identifier for the module's root directory and dictates how dependencies are managed. The reference documentation establishes specific lexical grammar and mandatory directives, such as the `module` declaration and the `go` directive (setting the minimum Go version), which became mandatory starting with Go 1.21.

## Versioning Strategies

Go employs Semantic Versioning (`vX.Y.Z`) as its standard for version identification. For major versions 2 and above, a suffix (e.g., `/v2`) is required in the import path to distinguish incompatible packages from legacy paths unless specific legacy handling is applied. Additionally, pseudo-versions are utilized to encode specific revisions, allowing for commit-specific references during testing or development.

## Dependency Resolution

The `go` command resolves package paths by searching the build list for matching module prefixes. This process consults proxies defined via the `GOPROXY` environment variable and prefers the longest matching path. Since Go 1.17, indirect dependencies are recorded in a separate block within the dependency graph. This separation enables module graph pruning and lazy loading, optimizing build performance by only including necessary packages.

## Minimal Version Selection (MVS)

To ensure deterministic builds, Go utilizes an algorithm known as Minimal Version Selection (MVS). MVS traverses the module graph to compute the minimal set of versions required for a specific build. This algorithm ensures that the exact set of versions needed is selected consistently across different environments.

## Directives and Lifecycle Management

The `go.mod` file contains various directives used to manage dependencies:
- **Mandatory:** `module`, `go`.
- **Optional:** `require` (for direct dependencies), `replace` (for substitution), `exclude`, `retract`, and `ignore`.

These directives allow developers to construct a dependency graph and manage the lifecycle of packages, including retracting or excluding specific versions.

## Vendoring and Environment

Build behaviors are heavily influenced by environment variables such as `GOPROXY`, `GOSUMDB`, and `GOVCS`. The build tool respects local vendor directories at the main module's root for vendoring behavior; when enabled, `go build` and `go test` use local copies of dependencies instead of fetching from the network. Other commands like `go mod tidy` continue to interact with the network or cache regardless of the vendoring status.

## Security and Verification

Module integrity is a critical aspect of Go Mod Versioning. The system verifies module integrity by comparing SHA-256 hashes found in the local `go.sum` file against a global checksum database (`sum.golang.org`). This verification process prevents tampering by untrusted proxies or origin servers, ensuring that the code downloaded matches the expected version.
