---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## Chunk Context
This chunk covers the `go.work` file format, its directives, and compatibility handling for pre-module repositories. It details lexical elements, grammar rules (EBNF), and specific behaviors regarding major version suffixes (`+incompatible`) and minimal module compatibility in GOPATH mode.

## Local Summary
The document explains how Go workspaces are defined using `go.work` files, which include directives for toolchain versioning, module inclusion (`use`), and dependency replacement (`replace`). It describes the syntax rules, including the mandatory `go` directive specifying the toolchain version. The text also addresses legacy compatibility, specifically how the Go command handles repositories that lack a `go.mod` file but have tags with major versions 2 or higher (appending `+incompatible` suffix). Finally, it outlines "minimal module compatibility" logic introduced in Go 1.11 to allow seamless imports between modules and GOPATH directories even when major version subdirectories are not strictly used.

## Key Claims
- A valid `go.work` file must contain exactly one `go` directive specifying the toolchain version.
- The `use` directive adds a module directory to the workspace; it does not recursively include submodules.
- If a repository root lacks a `go.mod` file but has a major version 2+ tag, the Go command synthesizes a synthetic `go.mod` and may append a `+incompatible` suffix to versions in that range.
- Minimal module compatibility allows packages imported via `$modpath/$vn/$dir` (where `$vn` is a major version suffix) to resolve to GOPATH locations without the suffix, provided specific conditions are met.

## Entities And Concepts
- **go.work**: UTF-8 text file defining a workspace.
- **Directive**: A line in `go.work` consisting of a keyword and arguments (e.g., `use`, `replace`).
- **Module-aware mode**: Build mode using `go.mod` files to resolve dependencies.
- **GOPATH mode**: Legacy build mode ignoring modules, looking in `$GOPATH/src`.
- **+incompatible**: Suffix added by the Go command for versions with major version 2+ lacking a `go.mod` file.
- **Minimal module compatibility**: Logic allowing mixed usage of module paths and GOPATH directories during imports.

## Procedures And API Details
- **Syntax Rule**: A `go.work` file follows Extended Backus-Naur Form (EBNF) where `GoWork = { Directive }`.
- **Directive Types**: `GoDirective`, `ToolchainDirective`, `UseDirective`, `ReplaceDirective`.
- **Example Command**: `go work init` creates a new file; `go work use` adds modules.
- **Compatibility Check**: When resolving an import `$modpath/$vn/$dir` in GOPATH mode:
  - If `$GOPATH[d]/src/$modpath/go.mod` exists and declares `$modpath/$vn`.
  - And the directory `$GOPATH[d]/src/$modpath/$vn/$dir` does not exist.
  - Then resolve to `$GOPATH[d]/src/$modpath/$dir`.

## Nuance Or Contradictions
- **Committing `go.work`**: Generally inadvisable because it can override parent workspace definitions or cause CI systems to test incorrect dependency versions, though exceptions exist for repositories developed exclusively together.
- **Version Tags vs. Versions**: The `+incompatible` suffix appears in Go command usage but should not appear on repository tags (e.g., a tag named `v4.1.2+incompatible` is ignored).
- **GOPATH Directory Structure**: A module with a major version suffix does not necessarily need to be developed in a subdirectory matching that suffix; the Go command supports resolving imports regardless of this structure if minimal compatibility rules apply.

## Candidate Wiki Hints
- **Topic: Go Workspace Files** – Overview of `go.work` syntax and directives.
- **Topic: Legacy Module Compatibility** – Handling pre-module repositories with major version 2+ tags.
- **Topic: Minimal Module Compatibility Rules** – Conditions for resolving imports across GOPATH and module paths.
