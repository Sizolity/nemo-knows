---
title: Go Mod Directives
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Go Mod Directives

**Go Mod Directives** are the specific lexical statements within a `go.mod` file that instruct the Go build tool on how to manage dependencies, define module identity, and control lifecycle behaviors. These directives enable the construction of a dependency graph, facilitate version resolution via algorithms like Minimal Version Selection (MVS), and ensure security verification against checksum databases.

## Core Directives

### Mandatory Directives
Certain directives are required for a valid module definition in modern Go versions:

- `module`: Defines the root path of the module. This is the primary identifier used by the dependency resolver to locate packages.
- `go`: Sets the minimum supported Go version for the module. This directive became mandatory starting with Go 1.21 to ensure compatibility and toolchain consistency.
- `toolchain`: Specifies the specific toolchain version required for building the module, distinct from the minimum Go version constraint.

### Dependency Management Directives
These directives directly influence how packages are fetched and substituted:

- `require`: Declares direct dependencies on other modules. The build tool uses these to construct the initial dependency graph. Since Go 1.17, indirect dependencies discovered during resolution are recorded in a separate block to enable pruning and lazy loading.
- `replace`: Provides a substitution mechanism, allowing one module path to be replaced by another local or remote path. This is often used for testing or fixing upstream bugs.
- `exclude`: Prevents the use of specific versions of a module, effectively removing them from consideration during resolution.
- `retract`: Marks a version as deprecated and signals that it should not be selected by the resolver, aiding in lifecycle management.
- `ignore`: Allows skipping verification of specific modules or paths, typically used for untrusted sources or internal tools.

## Lifecycle Directives

Beyond immediate dependency resolution, directives manage the health and state of the module graph:

- **Versioning Strategy**: The `go.mod` file supports Semantic Versioning (`vX.Y.Z`). For major versions 2 and above, a suffix (e.g., `/v2`) is required in the path to distinguish incompatible packages unless using legacy paths. Pseudo-versions can be used to encode specific revisions for testing scenarios.
- **Legacy Compatibility**: Directives and behaviors exist to handle pre-module repositories and GOPATH mode alongside modern modules, ensuring smooth transitions during migration periods.

## Behavior and Verification

The effectiveness of these directives relies on the surrounding environment and security protocols:

- **Environment Variables**: The interpretation and behavior of many directives (such as proxy usage `GOPROXY` and security checks `GOSUMDB`) are driven by environment variables.
- **Security Checks**: Integrity is verified by comparing SHA-256 hashes in the local `go.sum` file against a global checksum database. This verification process ensures that directives pointing to external sources have not been tampered with by untrusted proxies or origin servers.
- **Vendoring**: When vendoring is enabled, build commands utilize local copies of dependencies defined in vendor directories at the main module's root, respecting the directives within `go.mod` only for the primary module structure.
