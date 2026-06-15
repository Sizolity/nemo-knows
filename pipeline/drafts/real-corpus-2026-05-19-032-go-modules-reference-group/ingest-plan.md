---
kind: topic
sources: [raw/web/corpus-2026-05-18/032-go-modules-reference.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source is a comprehensive reference for Go Modules, covering the full lifecycle from metadata fetching and path resolution to dependency upgrades, workspace management, and security verification.
- Key topics include `go.mod` directives (`module`, `go`, `require`, `replace`, `retract`, `exclude`, `tool`), versioning strategies (semantic, pseudo, canonical), Minimal Version Selection (MVS), and vendoring behavior.
- The document details the `go` command's capabilities for managing dependencies (`go get`, `go mod tidy`, `go mod edit`), installing executables (`go install`), and introspecting builds (`go list -m`, `go version -m`).
- It provides extensive documentation on environment variables controlling proxy logic (`GOPROXY`, `GOPRIVATE`), security checks (`GOSUMDB`, `GOINSECURE`), and workspace configuration (`go work`).

## Candidate Wiki Pages
- wiki/sources/go-modules-reference.md — The primary source document covering the entire Go Modules reference system.
- wiki/concepts/go-mod-directives.md — Details syntax, purpose, and version-specific behavior for all `go.mod` directives including `module`, `go`, `require`, `replace`, `retract`, and `tool`.
- wiki/concepts/go-mod-versioning.md — Covers canonical vs non-canonical versions, major version suffixes (v2+), pseudo-versions, and path restrictions.
- wiki/topics/indirect-dependencies-in-go-mod.md — Explains the separate block for indirect dependencies introduced in Go 1.17 and its impact on pruning.
- wiki/concepts/go-workspace-files.md — Overview of `go.work` syntax, directives (`use`, `replace`), toolchain management, and workspace synchronization.
- wiki/concepts/go-modules-vendoring-behavior.md — Explains how vendoring affects build commands versus module management tools, including scope restrictions and known zip file limitations.
- wiki/concepts/minimal-version-selection.md — Explanation of the MVS algorithm used for dependency resolution in `go build`, `go get`, and `go work sync`.
- wiki/concepts/legacy-module-compatibility.md — Handling pre-module repositories, GOPATH mode, synthetic `go.mod` files, and minimal compatibility rules.
- wiki/topics/module-upgrade-strategies.md — Guide covering upgrades/downgrades for specific modules, transitive dependencies, Go versions, and toolchains using version queries (`@latest`, `@master`).
- wiki/concepts/go-mod-edit-manual.md — Details flags for manual `go.mod` manipulation including replacements, exclusions, Go version setting, and JSON printing.
- wiki/topics/proxy-protocol-specification.md — Defines HTTP endpoints (`@v/list`, `@v/$version.info`), response formats, fallback logic, and case encoding for module proxies.
- wiki/concepts/module-security-and-cache.md — Details the checksum database protocol, `go.sum` verification, cache structure, and environment variables for private modules.
- wiki/topics/module-environment-variables.md — A reference page listing all module-related env vars (`GO111MODULE`, `GOPROXY`, `GOSUMDB`, etc.) with defaults and usage examples.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that candidate pages do not duplicate content from existing wiki/index.md or schema files.
- [ ] Ensure all candidate paths are immediate children of `wiki/sources/`, `wiki/concepts/`, or `wiki/topics/`.
- [ ] Confirm that no nested directories are created in the file paths.
- [ ] Check that `go-mod-edit-manual` is placed under `wiki/concepts/` as it relates to tool usage rather than a generic API page.
- [ ] Validate that "Module Environment Variables" and "Proxy Configuration" are consolidated into the single `module-environment-variables.md` topic where applicable.
