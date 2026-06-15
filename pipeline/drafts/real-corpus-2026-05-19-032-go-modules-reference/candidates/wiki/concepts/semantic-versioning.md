---
title: Semantic Versioning
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Semantic Versioning

In the context of Go modules, **Semantic Versioning** is a standard for identifying module versions using the format `major.minor.patch` (e.g., `vX.Y.Z`). This system ensures that major increments imply incompatibility between API or ABI changes.

Versions are encoded in the module path suffix starting with v2 (e.g., `github.com/example/module/v2`). The Go toolchain relies on this format to determine upgrade safety and dependency resolution order. Alongside standard versions, the system supports **pseudo-versions** for unreleased code, which encode specific revision identifiers and timestamps to ensure canonical ordering during development.

## Resolution Behavior

When resolving dependencies:
1. The `go` command searches the local build list first.
2. If a matching version is not found locally, it queries configured proxies or Version Control System (VCS) repositories.
3. **Minimal Version Selection (MVS)** computes a deterministic build list from requirements while respecting exclusions and replacements.

## Directives and Management

Module versions are managed through specific `go.mod` directives:
- **require**: Specifies dependencies with their version constraints.
- **replace**: Substitutes a specific module version or all versions with a local file path or another remote module path.
- **exclude**: Prevents loading specific module versions in the main module since Go 1.16.
- **retract**: Marks versions as unusable for automatic upgrades while keeping them accessible in repositories.

Commands such as `go mod tidy` and `go mod vendor` interact with these versioned dependencies to maintain consistency between imported packages and the declared requirements.
