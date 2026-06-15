---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
Lines 416–859 of raw/web/corpus-2026-05-18/032-go-modules-reference.md cover the `go.mod` file structure, lexical elements, directives (module, go, toolchain, godebug, require, tool, ignore, exclude, replace), versioning rules, and compatibility notes for non-module repositories.

Local Summary
This section explains how `go.mod` files are structured, including allowed keywords (`module`, `go`, `require`, `replace`, `exclude`, `retract`), lexical tokens (whitespace, comments, punctuation, identifiers, strings), and the grammar using EBNF. It details each directive’s purpose, syntax, and behavior across Go versions, especially changes introduced in Go 1.14–1.24.

Key Claims
- A `go.mod` file must contain exactly one `module` directive.
- The `go` directive sets the minimum required Go version; since Go 1.21 it is mandatory.
- Since Go 1.17, indirect dependencies are recorded in a separate block and enable module graph pruning.
- Since Go 1.24, the `tool` directive adds packages as module-level tool dependencies.
- `ignore` directives prevent directories from being matched by package patterns.
- `exclude` directives prevent specific module versions from being loaded; behavior changed in Go 1.16 to avoid non-deterministic version selection.
- Non-canonical versions (without leading `v`) are allowed only in the main module’s `go.mod` and are automatically replaced with canonical versions when possible.

Entities And Concepts
- **Directives**: `module`, `go`, `toolchain`, `godebug`, `require`, `tool`, `ignore`, `exclude`, `replace`, `retract`.
- **Tokens**: whitespace, comments (`//`), punctuation (`(`, `)`, `=>`), keywords, identifiers, strings (interpreted and raw).
- **Versioning**: canonical (`v` + semantic version), non-canonical (restricted to main module).
- **Module paths**: domain-like leading element, slash-separated elements, restrictions on Windows reserved names and tilde-suffixes.
- **Indirect dependencies**: tracked separately since Go 1.17 for pruning and lazy loading.

Procedures And API Details
- `go get` upgrades/downgrades specific dependencies and updates `go.mod`.
- `go mod edit` performs low-level edits to `go.mod`.
- `golang.org/x/mod/modfile` package allows programmatic changes to `go.mod`.
- `go list -m -u` checks deprecated modules since Go 1.17.
- `go mod vendor` behavior varies with the `go` directive version (omits vendored dependency `go.mod`/`go.sum` files).

Nuance Or Contradictions
- Before Go 1.21, the `go` directive was advisory; now it is mandatory and affects toolchain selection.
- Indirect requirement handling differs between Go <1.16 and ≥1.17.
- `exclude` directives only apply in the main module’s `go.mod`; prior to Go 1.16 they caused non-deterministic version fallbacks.
- Deprecation applies to all minor versions of a major release; individual minor/patch versions should use `retract`.

Candidate Wiki Hints
- **Page**: `go-mod-directives` — summarize each directive with syntax, purpose, and version-specific behavior.
- **Page**: `go-mod-versioning` — canonical vs non-canonical versions, path restrictions, and toolchain integration.
- **Page**: `indirect-dependencies-in-go-mod` — explain the separate block for indirect dependencies since Go 1.17.
- **Page**: `deprecation-and-retraction-in-go-mod` — distinguish deprecation messages from version retractions.
