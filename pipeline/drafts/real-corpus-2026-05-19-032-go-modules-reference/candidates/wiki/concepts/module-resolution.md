---
title: Module Resolution
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Module Resolution

Module resolution is the process by which the Go toolchain determines where to fetch packages for a given import path. It relies on a **build list** (the local cache) and configured remote sources, typically via proxies or direct Version Control System (VCS) repositories.

## How It Works

The `go` command resolves packages by searching the local build list first. If no matching module is found locally, it queries configured proxy URLs (`GOPROXY`) or accesses VCS repositories directly. This process respects exclusions and replacements defined in the project's configuration files.

### Module Identification
A module is identified by a path declared in a `go.mod` file. These paths typically include:
- A repository root.
- An optional subdirectory.
- A major version suffix (e.g., `/v2`) for incompatible versions, introduced starting with Go 1.5.

### Versioning Strategies
Resolution supports multiple versioning schemes to handle both released and unreleased code:
- **Semantic Versions**: Follows `major.minor.patch` rules where major increments imply incompatibility.
- **Pseudo-versions**: Encode specific revision identifiers (commit hashes) and timestamps for unreleased code, ensuring canonical ordering without breaking import paths.
- **Keywords**: Terms like `@latest` or `@master` can be used to reference specific branches or tags.

## Configuration and Directives

Resolution behavior is controlled by directives in configuration files:

### go.mod Directives
The primary file for a main module uses UTF-8 text with specific directives:
- **module**: Declares the module path.
- **go**: Sets a mandatory minimum Go version.
- **require**: Lists direct dependencies (indirect ones are marked with `// indirect`).
- **replace**: Substitutes a specific module version or all versions with a local file path or another remote module path.
- **exclude**: Prevents loading specific module versions in the main module since Go 1.16.
- **retract**: Marks versions as unusable for automatic upgrades while keeping them accessible in repositories.

### go.work Directives
For multi-module workspaces, a `go.work` file allows running resolution across multiple main modules defined by relative paths. This file includes directives such as:
- **use**: Points to the root of a module.
- **replace**: Applies replacement rules within the workspace context.
- **toolchain**: Specifies the Go toolchain version for the workspace.

## Environment Variables

Resolution and access control are influenced by environment variables:
- **GOPROXY**: Controls proxy URLs or keywords (`direct`, `off`) to manage where modules are fetched from.
- **GONOSUMDB**: Specifies patterns for modules not to be checked against the public checksum database.
- **GOPRIVATE**: Marks module paths as private, bypassing external proxies and checksum databases.
- **GOVCS**: Controls allowed version control tools (e.g., `git`, `hg`) for downloading code.

## Security and Integrity

The system verifies downloaded modules against a checksum database (`sum.golang.org`) to prevent tampering. Modules fetched from private sources or those matching `GOPRIVATE` patterns bypass public verification checks.

## Commands

Several commands interact with the resolution state:
- **go get**: Updates module dependencies (deprecated for building packages since Go 1.17).
- **go install**: Recommended for installing programs since Go 1.16, ignoring the local `go.mod` if version suffixes are provided.
- **go mod tidy**: Aligns `go.mod` with imported packages, adding missing requirements and removing unused ones.
- **go mod vendor**: Constructs a `vendor` directory and `vendor/modules.txt` to avoid network access during builds.
