---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
- Source File: `raw/web/corpus-2026-05-18/032-go-modules-reference.md`
- Chunk Index: 6 of 13
- Line Range: 1666–1685
- Heading Path: Go Modules Reference > Retrieved Text

Local Summary
This section explains the behavior of build commands and module management tools when vendor directories are enabled. It clarifies that while `go build` and `go test` use vendored packages, other commands like `go mod download`, `go mod tidy`, and `go get` continue to interact with the network and local cache regardless of vendoring status. It also notes restrictions on using vendor directories outside the main module's root and their exclusion from zip files.

Key Claims
- When vendoring is enabled, `go build` and `go test` load packages from the vendor directory instead of accessing the network or local module cache.
- The `go list -m` command only prints information about modules listed in `go.mod`.
- Commands such as `go mod download`, `go mod tidy`, and `go get` do not behave differently when vendoring is enabled; they still download modules and access the module cache.
- Unlike GOPATH mode, the `go` command ignores vendor directories in locations other than the main module’s root directory.
- The `go` command does not include vendor directories from other modules when building module zip files (with references to known bugs #31562 and #37397).

Entities And Concepts
- Vendoring: A mechanism to include dependency code directly in the project source tree.
- Vendor Directory: The directory (`vendor/`) containing copied dependencies used by build tools when vendoring is enabled.
- Go Modules Reference: Documentation section covering module management behaviors.
- `go list -m`: Command to list modules, restricted to those in `go.mod`.
- Module Zip Files: Bundled archives of modules, excluding vendor directories from sub-modules.
- Known Bugs #31562 and #37397: Issues related to vendor directory handling in zip files.

Procedures And API Details
- Vendoring Enabled Build Flow:
  1. Execute `go build` or `go test`.
  2. If vendoring is enabled, the tool loads packages from the `vendor/` directory at the project root.
  3. Network access and local module cache are bypassed for these packages.
- Module Management Under Vendoring:
  - Running `go mod download`, `go mod tidy`, or `go get` does not change behavior; they still interact with the network and cache.
- Vendor Directory Usage Restrictions:
  - Only vendor directories at the main module’s root are respected by the `go` command.
  - Vendor directories in other modules are ignored during builds and excluded from zip files.

Nuance Or Contradictions
- The documentation highlights a discrepancy between GOPATH mode vendoring (where any vendor directory might be used) and Go Modules mode (which restricts usage to the main module’s root).
- There is an acknowledged inconsistency regarding zip file generation, where vendor directories are omitted despite potential relevance, noted via bug reports #31562 and #37397.

Candidate Wiki Hints
- Page: `Go Modules/Vendoring Behavior`
  - Summary: Explains how vendoring affects build commands versus module management commands in Go Modules mode.
  - Content Focus: Differences between GOPATH and Modules vendoring, scope of vendor directory usage, and known limitations with zip files.
