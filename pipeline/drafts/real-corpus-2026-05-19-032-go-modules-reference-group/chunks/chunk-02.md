---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

## Chunk Context

The chunk covers the core concepts of Go Modules, including their definition, identification via `go.mod` files, path structures, versioning schemes (semantic and pseudo-versions), major version suffixes for v2+, and the resolution logic used by the `go` command to locate packages within modules.

## Local Summary

This section introduces the Go module system as the mechanism for managing dependencies. It defines a module as a collection of versioned packages, identified by a path declared in a `go.mod` file. The text details how module paths are constructed (repository root, subdirectory, major version suffix), explains semantic versioning rules, and describes pseudo-versions used to reference specific commits or tags. It also outlines the rules for major version suffixes starting at v2 to ensure import compatibility between incompatible versions and concludes with an overview of how the `go` command resolves package paths to specific modules using proxies and build lists.

## Key Claims

- Modules are the standard way Go manages dependencies, consisting of collections of packages released and distributed together.
- A module is identified by a module path declared in a `go.mod` file located in the module root directory.
- Module paths must describe both functionality and location, typically comprising a repository root path, an optional subdirectory, and a major version suffix for versions 2+.
- Semantic versioning requires incrementing the major version after backwards-incompatible changes, minor versions after compatible changes, and patch versions for bug fixes or optimizations.
- Starting with major version 2, modules must use a major version suffix (e.g., `/v2`) to distinguish incompatible packages; this is not required for v0 or v1 unless using `gopkg.in/` paths.
- Pseudo-versions encode specific revision information (timestamp and commit hash) and sort between the base version and the next tagged version to facilitate testing and dependency resolution without manual typing.
- The `go` command resolves packages by searching the build list for matching module prefixes, consulting proxies defined in `GOPROXY`, and preferring the module with the longest matching path.

## Entities And Concepts

- **Module**: A collection of packages released, versioned, and distributed together.
- **Module Path**: The canonical name declared in `go.mod`; serves as a prefix for package paths within the module.
- **Package Path**: The module path joined with the subdirectory containing the package.
- **Semantic Versioning**: A system where versions are formatted as `vX.Y.Z`, with rules for major, minor, and patch increments to indicate compatibility changes.
- **Pseudo-version**: A pre-release version encoding a specific revision (timestamp + hash) used to reference commits or tags without a stable semantic tag.
- **Major Version Suffix**: A `/vN` suffix in the module path required from major version 2 onwards to handle incompatibility and diamond dependencies.
- **Build Metadata**: Optional suffixes like `+meta`, `+incompatible`, or `+dirty` appended to versions, largely ignored for comparison purposes except for specific legacy cases.
- **GOPROXY**: An environment variable controlling the list of module proxy servers the `go` command contacts to resolve dependencies.

## Procedures And API Details

- **Module Identification**: The `go` command looks for a `go.mod` file in the directory where the `go` command is invoked; this defines the main module and its root.
- **Version Resolution**: Commands like `go get`, `go mod tidy`, or `go list -m` can accept commit hashes (e.g., `@daa7c041`) or branch names, which the tool automatically converts into pseudo-versions or tagged versions.
- **Dependency Lookup**: When loading a package path, the `go` command searches the build list for modules whose paths are prefixes of the package path, checking directories for `.go` files to confirm package existence.
- **Proxy Usage**: The `go` command queries proxies in order (e.g., `https://corp.example.com`, then `https://proxy.golang.org`) requesting the latest version of module path prefixes until a match is found or an error occurs.

## Nuance Or Contradictions

- **Legacy Incompatibility**: Modules released at v2+ before adopting the module system are marked with `+incompatible` (e.g., `v2.0.0+incompatible`) to distinguish them from modern module paths without implying a specific version constraint.
- **Pseudo-version Sorting**: Pseudo-versions must sort higher than their base version but lower than subsequent tagged versions; the `go` command enforces this by verifying that the timestamp and revision match actual repository data, preventing version flooding or manipulation.
- **Path Restrictions**: Module paths cannot contain dots in the first element to avoid confusion with package paths, and specific names like `example` and `test` are reserved for user-defined modules to prevent collisions with standard library expectations.

## Candidate Wiki Hints

- Go Modules Reference
- Semantic Versioning in Go
- Pseudo-versions and Commit Hashes
- Major Version Suffixes (v2+)
- Module Path Resolution Logic
