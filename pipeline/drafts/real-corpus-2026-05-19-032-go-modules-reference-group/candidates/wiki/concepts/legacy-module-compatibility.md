---
title: Legacy Module Compatibility
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Legacy Module Compatibility

## Definition

**Legacy Module Compatibility** refers to the mechanisms within Go's build system that allow projects to interact with pre-module repositories and the traditional GOPATH directory layout while transitioning toward or coexisting with the modern module system. This ensures deterministic builds for code that predates `go.mod` usage or relies on legacy import paths.

## How It Works

When a build is performed, Go determines whether to use the module system or fallback to legacy behavior based on the presence of a `go.mod` file and the configuration of environment variables like `GOPROXY`. If a project lacks a `go.mod` file in its root directory, the tool may enter **GOPATH mode**, treating packages as local files within the GOPATH hierarchy.

In this mode, import paths are resolved relative to the GOPATH structure rather than via semantic versioning or proxy lookups. This fallback is critical for maintaining compatibility with older codebases that have not yet migrated to the module system.

## Key Directives and Behaviors

While modern modules use directives like `require`, `replace`, and `exclude` within `go.mod`, legacy compatibility relies on implicit resolution rules:

- **Path Resolution:** Without a `go.mod` file, import paths are matched against local GOPATH directories.
- **Proxy Usage:** Legacy paths generally bypass the proxy protocol specification unless explicitly configured to fetch from a remote source.
- **Versioning:** Unlike modern modules which enforce semantic versioning (`vX.Y.Z`) and require suffixes (e.g., `/v2`) for major versions, legacy paths rely on directory structures within GOPATH.

## Related Concepts

- [[go-mod-directives]]
- [[go-mod-versioning]]
- [[go-workspace-files]]
- [[module-environment-variables]]
- [[proxy-protocol-specification]]
