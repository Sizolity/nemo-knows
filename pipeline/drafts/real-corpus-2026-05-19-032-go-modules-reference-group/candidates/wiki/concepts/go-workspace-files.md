---
title: Go Workspace Files
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Go Workspace Files

**Go workspace files** (`go.work`) define a multi-module environment for the Go toolchain. They allow a single `go` command to manage dependencies across multiple modules within a directory tree, specifying which module roots are included and their respective toolchain versions.

## Structure and Directives

The file is identified by its name, `go.work`, located in the project root. Its lexical grammar supports specific directives:

- **use**: Declares the list of module paths to include in the workspace. These correspond to directories containing a `go.mod` file.
- **toolchain**: Specifies the Go version for the workspace. This ensures consistent builds across different environments when multiple modules are involved.

## Workflow Integration

The tool uses these directives to construct a unified dependency graph, effectively merging the requirements of all included modules. When resolving dependencies, the command consults proxies (defined via `GOPROXY`) and prefers the longest matching path within the workspace context.

## Legacy Compatibility

Go handles pre-module repositories alongside modern modules. In a workspace context, this ensures that legacy module paths or GOPATH mode dependencies can coexist with standard module definitions without breaking existing workflows.

## Security Considerations

Module integrity is verified by comparing SHA-256 hashes in the local `go.sum` file against a global checksum database (`sum.golang.org`). This applies to all modules included in the workspace, preventing tampering by untrusted proxies or origin servers.

## Environment Variables

The behavior of the workspace setup and dependency resolution is largely controlled by environment variables with defined defaults, such as:

- `GOPROXY`: Defines the proxy for fetching module metadata.
- `GOSUMDB`: Specifies the checksum database URL.
- `GOVCS`: Sets the version control system to use for pseudo-versioning.

## Related Concepts

[[go-mod-directives]]
[[go-mod-edit-manual]]
[[go-mod-versioning]]
[[indirect-dependencies-in-go-mod]]
[[module-environment-variables]]
