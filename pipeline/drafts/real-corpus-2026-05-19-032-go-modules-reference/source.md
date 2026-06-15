---
title: Go Modules Reference
kind: source
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## What It Is

The **Go Modules Reference** is the official documentation for Go's dependency management system. Originally introduced as a replacement for GOPATH-based imports, modules define a module as a collection of versioned packages released and distributed together. The reference covers the entire lifecycle of managing dependencies, from defining module paths and semantic versions to resolving dependencies, handling replacements, and verifying integrity via checksums.

## Summary

Go Modules allows developers to manage dependencies using `go.mod` files for the main module and optional `go.work` files for workspaces containing multiple modules. The system relies on **Minimal Version Selection (MVS)** to deterministically compute a build list from requirements, respecting exclusions and replacements. Modules are identified by paths that typically include a repository root, subdirectory, and major version suffix (v2+) for incompatible versions. The `go` command resolves packages by searching the local build list first, then querying configured proxies (`GOPROXY`) or direct Version Control System (VCS) repositories.

Key capabilities include:
- **Versioning**: Support for semantic versions (`vX.Y.Z`), pseudo-versions (encoding commit hashes), and keywords like `@latest` or `@master`.
- **Management Commands**: Tools such as `go get`, `go install`, `go mod tidy`, `go mod vendor`, and `go mod edit` handle dependency updates, installation, cleaning, and file manipulation.
- **Security & Privacy**: The system verifies downloaded modules against a checksum database (`sum.golang.org`) to prevent tampering. Environment variables like `GOPRIVATE` and `GOVCS` allow handling of private modules and controlling VCS access.
- **Compatibility**: Mechanisms exist for legacy GOPATH repositories (via `+incompatible` suffix) and ensuring cross-platform compatibility in distributed zip files.

## Key Claims

### Module Definition and Paths
- A module is identified by a path declared in a `go.mod` file.
- Module paths should describe functionality and location, often including a repository root and optional subdirectory (e.g., `/v2` for incompatible versions).
- Starting with v2, major version suffixes are required to maintain import compatibility rules between incompatible versions.

### Versioning and Resolution
- **Semantic Versioning**: Versions follow `major.minor.patch`; major increments imply incompatibility.
- **Pseudo-versions**: Encode specific revision identifiers (commit hashes) and timestamps for unreleased code, ensuring canonical ordering.
- **Resolution Process**: The `go` command searches the build list for matching path prefixes. If none are found locally, it queries GOPROXY entries.

### go.mod Directives and Syntax
- A `go.mod` file defines a module using UTF-8 text with directives like `module`, `go`, `require`, `replace`, `exclude`, `retract`, `tool`, and `ignore`.
- Indirect dependencies are marked with a `// indirect` comment.
- The `go` directive declares a mandatory minimum Go version; toolchains refuse to use modules requiring newer versions.

### Module Graph Manipulation
- **Replace**: Substitutes a specific module version or all versions with a local file path or another remote module path.
- **Retract**: Marks versions as unusable for automatic upgrades while keeping them accessible in repositories.
- **Exclude**: Prevents loading specific module versions in the main module since Go 1.16.

### Workspaces and Vendoring
- **Workspaces (`go.work`)**: Allow running MVS across multiple main modules defined by relative paths in a `go.work` file. Directives include `use`, `replace`, `toolchain`, and `godebug`.
- **Vendoring**: Copies dependencies into a local `vendor` directory to avoid network access during builds. Build commands use the vendor directory at the main module's root, while management commands continue to interact with the cache.

### Commands and Utilities
- **go get**: Updates module dependencies (deprecated for building packages since Go 1.17).
- **go install**: Recommended for installing programs since Go 1.16, ignoring the local `go.mod` if version suffixes are provided.
- **go mod edit**: Modifies `go.mod` and can output JSON representations of the file structure.
- **go mod tidy**: Aligns `go.mod` with imported packages, adding missing requirements and removing unused ones.
- **go mod vendor**: Constructs a `vendor` directory and `vendor/modules.txt`.

### Security and Environment Variables
- **Checksum Verification**: The `go` command verifies downloaded modules against hashes in `go.sum`.
- **Environment Variables**:
  - `GOPROXY`: Controls proxy URLs or keywords (`direct`, `off`).
  - `GONOSUMDB`: Specifies patterns for modules not to be checked against the public checksum database.
  - `GOPRIVATE`: Marks module paths as private, bypassing external proxies and checksum databases.
  - `GOVCS`: Controls allowed version control tools (e.g., `git`, `hg`) for downloading code.

## Suggested Links

- https://go.dev/ref/mod
