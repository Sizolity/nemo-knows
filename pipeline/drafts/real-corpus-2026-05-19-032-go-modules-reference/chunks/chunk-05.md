---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## Chunk Context
This chunk details the syntax, directives, and behavior of `go.work` files used to define Go workspaces. It covers lexical elements, grammar, specific directives (`go`, `toolchain`, `godebug`, `use`, `replace`), compatibility handling for non-module repositories (specifically the `+incompatible` suffix), minimal module compatibility rules for GOPATH vs module mode, and the behavior of module-aware commands including vendoring.

## Local Summary
The section explains how to define a workspace using a `go.work` file, which is line-oriented with directives like `use`, `replace`, and `go`. It details the grammar (EBNF) for these directives. The text also addresses legacy compatibility issues where Go commands handle repositories without `go.mod` files by synthesizing them or adding an `+incompatible` suffix to major versions >= 2 in the root directory. Finally, it describes how the `go` command behaves in module-aware mode versus GOPATH mode, including the `-mod` flags and vendoring mechanisms.

## Key Claims
- A workspace is defined by a UTF-8 encoded text file named `go.work`.
- The `go.work` file must contain exactly one `go` directive specifying the toolchain version.
- Directives in `go.work` include `use`, `replace`, `toolchain`, and `godebug`.
- `use` directives add module directories to the workspace; they do not automatically include subdirectories.
- A wildcard replace in `go.work` overrides a version-specific replace in `go.mod`.
- Repositories with major version 2+ releases but no `go.mod` file get an `+incompatible` suffix appended by the Go command.
- Minimal module compatibility rules allow importing packages from a major version subdirectory (e.g., `/v2`) even if the GOPATH directory structure does not reflect that subdirectory, provided specific conditions are met.
- As of Go 1.16, module-aware mode is enabled by default regardless of the presence of a `go.mod` file.

## Entities And Concepts
- **go.work**: A text file defining a Go workspace containing multiple modules.
- **go.work.sum**: A file maintaining hashes for dependencies not present in collective workspace modules' `go.sum` files.
- **Directive**: A line in `go.work` starting with a keyword (e.g., `use`, `replace`).
- **+incompatible**: A suffix added by the Go command to versions >= 2.0.0 found in repositories lacking a `go.mod` file to indicate they are part of the same legacy module as lower major versions.
- **Minimal Module Compatibility**: Rules introduced in Go 1.11 allowing mixed usage of modules and GOPATH imports across major version subdirectories.
- **Module-aware mode**: The default state (Go 1.16+) where the `go` command uses `go.mod` files for dependency resolution.
- **Vendoring**: Storing dependencies in a local `vendor` directory instead of downloading from the module cache.

## Procedures And API Details
### Creating and Editing go.work Files
- **go work init**: Creates a new `go.work` file.
- **go work use**: Adds module directories to the `go.work` file.
- **go work edit**: Performs low-level edits on the `go.work` file.
- **golang.org/x/mod/modfile**: A package for making programmatic changes to `go.work` files.

### go.work Directives
- **go <version>**: Specifies the toolchain version (e.g., `go 1.23.0`). Must be a valid release version.
- **toolchain <name>**: Declares a suggested Go toolchain (effective only if the default is older).
- **godebug <setting>**: Applies a GODEBUG setting for the workspace; ignores `godebug` directives in individual `go.mod` files within the workspace.
- **use <path>**: Adds a module directory to the workspace. Supports grouping:
  ```go
  use (
    ./my/first/thing
    ./my/second/thing
  )
  ```
- **replace <old> => <new>**: Replaces module contents. Supports wildcards and version ranges. Overrides replaces in `go.mod`.
  ```go
  replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
  ```

### Compatibility Handling
- **Synthetic go.mod**: If a module path equals the repository root and lacks a `go.mod`, the command synthesizes one in the cache with only a module directive.
- **+incompatible suffix**: Applied to versions >= 2.0.0 in repositories without `go.mod`. Tags like `v4.1.2+incompatible` are ignored by the repository; the suffix appears only in Go command usage.

### Module-Aware Commands and Flags
- **Module-aware commands**: `go build`, `go test`, `go install`, etc., use `go.mod` files to interpret import paths.
- **-mod flag**: Controls dependency resolution:
  - `-mod=mod`: Ignore vendor, auto-update `go.mod`.
  - `-mod=readonly`: Ignore vendor, error if `go.mod` needs update.
  - `-mod=vendor`: Use vendor directory; no network/cache access.
- **-modcacherw**: Creates module cache directories with read-write permissions.
- **-modfile=file.mod**: Reads/writes an alternate `.mod` file instead of `go.mod`.

### Vendoring
- **go mod vendor**: Constructs a `vendor` directory and `vendor/modules.txt`.
- Vendor directory is used automatically if present and `go.mod` version >= 1.14, or explicitly via `-mod=vendor`.

## Nuance Or Contradictions
- **Committing go.work files**: Generally inadvisable due to potential conflicts with parent directory workspaces or CI systems testing wrong dependency versions. However, it is acceptable if modules are developed exclusively together without external dependencies.
- **GOPATH vs Module Path Mismatch**: In GOPATH mode, a module at `example.com/repo/v2` might be found at `$GOPATH/src/example.com/repo/sub` (without the `/v2`) if not developed in the subdirectory, causing import path mismatches unless minimal compatibility rules apply.
- **godebug precedence**: `godebug` directives in `go.mod` files are ignored when a workspace (`go.work`) is active; only `go.work` `godebug` directives apply.

## Candidate Wiki Hints
- **Topic: go.work Syntax** – Document the structure, directives, and EBNF grammar for workspace definition files.
- **Topic: Workspace Directives (use/replace)** – Detail how to add modules and override dependencies in a workspace context.
- **Topic: Legacy Module Compatibility** – Explain the `+incompatible` suffix and handling of pre-module repositories.
- **Topic: Minimal Module Compatibility** – Describe the rules allowing cross-mode imports between major version subdirectories.
- **Topic: Vendoring in Go Modules** – Cover the creation and usage of `vendor` directories with `go mod vendor`.
