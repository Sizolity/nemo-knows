---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## Chunk Context
This chunk (lines 1686–2071) details the `go get` command for managing module dependencies, including upgrading/downgrading specific modules, handling transitive dependencies, and managing version queries. It also introduces `go install` for building executables, `go tool` for running tools, `go list -m` for listing modules, `go mod download` for caching, and `go mod edit` for modifying `go.mod`.

## Local Summary
The text explains how to use `go get` to update, downgrade, or remove specific module versions, including handling transitive dependencies. It highlights the deprecation of building/installing packages with `go get` in favor of `go install`. The chunk also covers version query syntax (`@v1.0.0`, `@master`, `@latest`, `@none`), flags for `go get` (like `-u`, `-d`, `-tool`), and related commands like `go mod download` and `go mod edit`.

## Key Claims
- `go get` updates module dependencies in the main module’s `go.mod` file and builds/installing packages is deprecated since Go 1.17.
- Version query suffixes (`@v`, `@master`, `@latest`, `@none`) control version selection for modules.
- Transitive dependencies are upgraded or downgraded automatically when their direct dependencies change versions.
- `go install` is recommended for installing programs since Go 1.16, ignoring the local `go.mod` if version suffixes are provided.
- `go mod download` pre-fills the module cache; `go mod edit` modifies `go.mod` (e.g., adding replace directives).

## Entities And Concepts
- **Commands**: `go get`, `go install`, `go tool`, `go list -m`, `go mod download`, `go mod edit`.
- **Version Queries**: `@v1.0.0`, `@master`, `@latest`, `@upgrade`, `@patch`, `@none`.
- **Flags**: `-u` (upgrade), `-d` (no build, deprecated), `-tool`, `-insecure`, `-t` (test dependencies).
- **Structs**: `Module`, `ModuleError`.
- **Environment Variables**: `GOBIN`, `GOPATH`, `GOROOT`, `GOINSECURE`, `GO111MODULE`.

## Procedures And API Details
- **Upgrading a specific module**:
  ```bash
  go get golang.org/x/net
  ```
- **Downgrading to `@none`**:
  ```bash
  go get golang.org/x/text@none
  ```
- **Installing a program (ignoring local `go.mod`)**:
  ```bash
  go install golang.org/x/tools/gopls@latest
  ```
- **Listing modules with upgrades**:
  ```bash
  go list -m -u all
  ```
- **Downloading modules**:
  ```bash
  go mod download golang.org/x/mod@v0.2.0
  ```
- **Adding a replace directive**:
  ```bash
  go mod edit -replace example.com/a@v1.0.0=./a
  ```

## Nuance Or Contradictions
- `go get` without `-d` is deprecated since Go 1.17; in Go 1.18, `-d` is always enabled.
- `go install` ignores the local `go.mod` if version suffixes are provided (module-aware mode), whereas without suffixes it may run in GOPATH or module-aware mode depending on `GO111MODULE`.
- Retracted or deprecated modules are flagged by `go get` and `go list -m -u`, but `go list -m -retracted` explicitly includes retracted versions in version lists.

## Candidate Wiki Hints
- **Page**: `go-get-command` – Covers usage, flags, version queries, and transitive dependency handling.
- **Page**: `go-install-command` – Focuses on installing executables with module-aware behavior.
- **Page**: `module-version-queries` – Details syntax for `@v`, `@master`, `@latest`, etc.
