---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
# Chunk Context

This chunk covers the introduction to Go modules, defining them as collections of versioned packages. It details module paths, semantic versions (including pre-release and build metadata), pseudo-versions encoding revision history, major version suffixes for v2+ compatibility, and the resolution process where `go` searches the build list or GOPROXY to find the providing module.

# Local Summary

The text introduces Go modules as the dependency management system, explaining that a module is a collection of packages released and versioned together. It outlines the structure of module paths (repository root, subdirectory, major version suffix), semantic versioning rules, and the mechanics of pseudo-versions used for unreleased commits. The chunk further explains the introduction of major version suffixes starting with v2 to handle incompatibilities and resolves how `go` identifies which module provides a specific package path by checking prefixes in the build list or proxy servers.

# Key Claims

- Modules are collections of packages released, versioned, and distributed together.
- A module is identified by a module path declared in a `go.mod` file.
- Module paths should describe functionality and location, typically including a repository root, optional subdirectory, and major version suffix (v2+).
- Semantic versions consist of major.minor.patch; major increments require incompatible changes, minor for compatible additions, patch for internal fixes.
- Pseudo-versions encode specific revision identifiers (commit hashes) and timestamps to ensure canonical ordering without manual typing.
- Starting with v2, module paths must include a major version suffix (e.g., `/v2`) to maintain import compatibility rules between incompatible versions.
- The `go` command resolves packages by searching modules in the build list for matching path prefixes; if none are found locally, it queries GOPROXY entries.

# Entities And Concepts

- **Module**: A collection of packages released and versioned together.
- **Module Path**: The canonical name declared in `go.mod`, acting as a prefix for package paths within the module.
- **Semantic Versioning**: Format `vX.Y.Z` with optional `-pre-release` and `+build-metadata`.
- **Pseudo-version**: A pre-release version encoding a specific revision identifier (e.g., `v0.0.0-20191109...`).
- **Major Version Suffix**: Required for v2+ to distinguish incompatible module paths (e.g., `/v2`).
- **GOPROXY**: Environment variable controlling the list of proxy URLs or keywords (`direct`, `off`) for downloading modules.
- **Build List**: The set of modules currently available to the build process, checked first during resolution.

# Procedures And API Details

- **Resolving a Package**:
  1. `go` searches the build list for modules whose paths are prefixes of the package path.
  2. If exactly one module provides the package, it is used.
  3. If none or multiple match, an error is reported unless `-mod=mod` is used to fetch new modules.
- **GOPROXY Requests**: For each entry in `GOPROXY`, `go` requests the latest version of each potential module path prefix (e.g., for package `golang.org/x/net/html`, it requests `golang.org/x/net/html`, `golang.org/x/net`, etc.).
- **Version Conversion**: Commands like `go get` or `go list -m` can accept branch names or commit hashes, automatically translating them into pseudo-versions or tagged versions.

# Nuance Or Contradictions

- **v0/v1 vs v2 Suffixes**: Major version suffixes are not allowed at v0 or v1 because v0 is unstable and v1 implies compatibility with the previous v0 release.
- **gopkg.in Exception**: Modules starting with `gopkg.in/` must always have a major version suffix, even at v0/v1, but using a dot separator (e.g., `gopkg.in/yaml.v2`) instead of a slash.
- **Pseudo-version Ordering**: Pseudo-versions sort higher than their base version but lower than the next tagged version, ensuring they sit correctly in version ordering without manual intervention.

# Candidate Wiki Hints

- Create a page on **Go Modules Basics** explaining what modules are and how `go.mod` identifies them.
- Draft a guide on **Semantic Versioning in Go**, detailing major/minor/patch rules and the meaning of `-pre` and `+incompatible`.
- Write an article on **Pseudo-versions**, explaining how they encode commit hashes and timestamps for unreleased code.
- Develop a section on **Major Version Suffixes**, covering when `/v2` is required and the `gopkg.in` special case.
- Create a troubleshooting page for **Module Resolution Errors**, discussing build list checks and GOPROXY configuration.
