---
kind: topic
sources: [raw/web/corpus-2026-05-18/032-go-modules-reference.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source is the "Go Modules Reference" documentation from `go.dev`, covering dependency management, module syntax, and tooling.
- Content spans metadata headers, core concepts (modules, versions, resolution), `go.mod`/`go.work` directives, command usage (`go get`, `go install`, `go mod`), security (checksums, private modules), and environment configuration.
- Multiple chunks detail specific behaviors like Minimal Version Selection (MVS), pseudo-versions, proxy protocols, vendoring scope, and workspace management.

## Candidate Wiki Pages
- wiki/sources/go-modules-reference.md — Comprehensive reference for Go's dependency management system derived from the official documentation.
- wiki/concepts/module-resolution.md — Covers resolution mechanics, build lists, GOPROXY fallbacks, and version query logic.
- wiki/concepts/semantic-versioning.md — Explains major/minor/patch rules, pseudo-versions, and v2+ compatibility suffixes.
- wiki/concepts/go-mod-directives.md — Details `go.mod` syntax, directives (module, go, require, replace, retract), and block grouping.
- wiki/concepts/go-work-syntax.md — Documents `go.work` file structure, directives (use, replace, toolchain), and workspace behavior.
- wiki/topics/module-vendor-scope.md — Clarifies how vendoring interacts with build commands versus module management commands.
- wiki/topics/private-module-setup.md — Guides configuration of `GOPRIVATE`, `GONOPROXY`, `GOSUMDB`, and direct VCS access for private code.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that all candidate pages map to distinct concepts without significant overlap (e.g., separating MVS into its own concept page).
- [ ] Ensure no nested directories are created; keep slugs simple and descriptive.
- [ ] Confirm that tool/API pages (like `go mod edit`) are placed under `wiki/concepts/` or `wiki/topics/`.
