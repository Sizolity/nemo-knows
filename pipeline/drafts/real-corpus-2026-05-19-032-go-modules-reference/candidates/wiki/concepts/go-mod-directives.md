---
title: Go Mod Directives
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Go Mod Directives

In the context of **[[go-work-syntax]]**, a `go.mod` file defines a module using UTF-8 text with specific directives that control dependency resolution, tooling, and version constraints. These directives allow developers to manage the lifecycle of dependencies from defining paths to verifying integrity.

## Directives in `go.mod`

The following directives are recognized within the main module's `go.mod` file:

*   **module**: Declares the module path, which identifies the collection of versioned packages released and distributed together.
*   **go**: Sets a mandatory minimum Go version; toolchains refuse to use modules requiring newer versions than specified.
*   **require**: Specifies direct dependencies. Indirect dependencies are typically marked with a `// indirect` comment.
*   **replace**: Substitutes a specific module version or all versions with a local file path or another remote module path. This is useful for overriding dependencies during development or testing.
*   **exclude**: Prevents loading specific module versions in the main module since Go 1.16.
*   **retract**: Marks versions as unusable for automatic upgrades while keeping them accessible in repositories.
*   **tool**: Specifies a toolchain to be used, distinct from the main Go version directive.
*   **ignore**: Allows ignoring specific requirements or constraints during resolution.

## Directives in `go.work`

When using workspaces to run **Minimal Version Selection (MVS)** across multiple main modules defined by relative paths, the `go.work` file supports additional directives:

*   **use**: Declares a directory containing a main module.
*   **replace**: Substitutes modules within the workspace context.
*   **toolchain**: Specifies the toolchain for the workspace.
*   **godebug**: Controls debugging behavior for the workspace.

## Related Concepts

These directives interact closely with the overall **[[module-resolution]]** process, which searches the local build list first and then queries configured proxies or direct Version Control System repositories if no match is found locally. Security mechanisms verify downloaded modules against checksums, while environment variables like `GOPRIVATE` allow handling private modules bypassing external checks.
