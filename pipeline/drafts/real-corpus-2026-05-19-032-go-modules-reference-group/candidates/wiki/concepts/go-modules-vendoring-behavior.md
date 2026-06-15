---
title: Go Modules Vendoring Behavior
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Go Modules Vendoring Behavior

## Definition

**Go Modules Vendoring Behavior** describes the conditions under which the Go build tool uses local copies of dependencies stored in a vendor directory rather than fetching them from the network or cache. This behavior is specifically enabled for `go build` and `go test` commands when the `-mod=vendor` flag is used or a vendor directory exists at the module root.

## Mechanism

When vendoring is active, the build process prioritizes packages found in the local vendor directory. However, this behavior is scoped strictly to the execution of specific commands:

- **Affected Commands:** `go build` and `go test`.
- **Unaffected Commands:** Operations such as `go mod tidy`, `go get`, and others continue to interact with the network and cache regardless of whether a vendor directory exists.

The Go tool respects only vendor directories located at the root of the main module. It does not follow vendor paths defined deep within subdirectories or indirect dependencies in the same way.

## Interaction with Directives

Vendoring is closely related to the `require` directive found in the `go.mod` file, which lists direct and indirect dependencies. While these directives define the dependency graph, the actual usage of local copies depends on the build command flags and the presence of the vendor directory structure. The behavior ensures that builds are reproducible and do not depend on external network connectivity during the compilation phase when vendoring is enforced.

## See Also

- [[go-mod-directives]]
- [[go-workspace-files]]
- [[module-environment-variables]]
