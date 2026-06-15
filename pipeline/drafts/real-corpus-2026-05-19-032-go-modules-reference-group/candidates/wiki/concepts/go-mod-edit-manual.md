---
title: Go Mod Edit Manual
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Go Mod Edit Manual

This manual provides guidance on editing and maintaining the `go.mod` file, which serves as the primary declaration of a Go module's identity and dependencies. It outlines the syntax, mandatory directives, optional lifecycle directives, and interaction with build tools like `go mod tidy`.

## Module Identity

A module is defined by a path declared in the `module` directive within a `go.mod` file located at the root of the project. This path acts as the unique identifier for the module in the dependency graph.

## Directives

The `go.mod` file uses specific directives to configure behavior and manage dependencies:

### Mandatory Directives
- **`module`**: Declares the module path (required).
- **`go`**: Sets the minimum Go version required for the module (mandatory since Go 1.21).
- **`toolchain`**: Specifies the toolchain version for building and testing.

### Optional Directives
- **`require`**: Adds direct dependencies to the build graph.
- **`replace`**: Substitutes a specific module path with another local or remote path (useful for development overrides).
- **`exclude`**: Prevents specific versions of modules from being used, even if they appear in indirect dependencies.
- **`retract`**: Marks specific versions as deprecated or unavailable to be selected by the resolver.
- **`ignore`**: Disables a module path entirely (advanced usage).

## Dependency Resolution and Selection

When resolving dependencies, the build tool uses algorithms like **[[minimal-version-selection]]** (MVS) to traverse the module graph. This ensures the minimal set of versions required for a build is selected, promoting deterministic builds and reproducible environments.

Since Go 1.17, indirect dependencies are recorded in a separate block within `go.mod` to enable efficient pruning and lazy loading. The **[[indirect-dependencies-in-go-mod]]** mechanism helps manage the complexity of large dependency trees.

## Vendoring Behavior

The manual covers **[[go-modules-vendoring-behavior]]**, where local copies of dependencies are used instead of fetching from the network. Note that `go build` and `go test` respect vendored packages, while commands like `go mod tidy` continue to interact with the cache or network regardless of vendoring status.

## Environment Variables

The behavior of the `go` command, including module mode activation, proxy usage, and security checks, is largely controlled by environment variables such as **[[module-environment-variables]]**. Key variables include:
- `GOPROXY`: Defines the proxy server for fetching modules.
- `GOSUMDB`: Specifies the checksum database for verifying module integrity.
- `GOVCS`: Sets the version control system to use for fetching source code.

## Security and Verification

Module integrity is verified by comparing SHA-256 hashes in the local `go.sum` file against a global checksum database (**[[module-security-and-cache]]**). This prevents tampering by untrusted proxies or origin servers, ensuring that only verified modules are installed.

## Legacy Compatibility

The manual addresses **[[legacy-module-compatibility]]**, detailing how Go handles pre-module repositories and GOPATH mode alongside modern module paths. It also explains the requirements for major version 2+ paths (**[[go-mod-versioning]]**), where a suffix (e.g., `/v2`) is required to distinguish incompatible packages unless using legacy paths.

## Related Concepts

- **[[go-mod-directives]]**
- **[[go-workspace-files]]**
- **[[proxy-protocol-specification]]**
- **[[module-upgrade-strategies]]**
