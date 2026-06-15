---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
This chunk covers the `replace` directive syntax and semantics, the `retract` directive for version management, automatic updates via `-mod=mod`, Minimal Version Selection (MVS), and optimizations in Go 1.17+ such as module graph pruning and lazy loading.

Local Summary
The document explains how to manipulate the module graph using directives: replacing local or remote modules, retracting versions to prevent upgrades, and updating `go.mod` files automatically. It details MVS, which deterministically selects the minimum required set of versions. Finally, it describes Go 1.17+ features that optimize the build process by pruning unnecessary transitive dependencies and loading the full graph only when needed.

Key Claims
- A `replace` directive substitutes a specific module version with another (local or remote) path; if no left-side version is specified, all versions are replaced.
- `retract` directives mark versions as problematic, preventing automatic upgrades but keeping them accessible for existing builds.
- The `-mod=mod` flag enables the Go command to automatically rewrite non-canonical versions, respect exclusions, remove redundant requirements, and reformat `go.mod`.
- Minimal Version Selection (MVS) computes a build list by traversing the module graph from main modules, tracking the highest required version for each dependency.
- In Go 1.17+, the module graph is pruned to exclude transitive dependencies of modules specifying `go 1.17+`, unless those dependencies are also required by older Go versions.

Entities And Concepts
- **Replace Directive**: Substitutes a module path/version with a local file path or another module path/version.
- **Retract Directive**: Marks specific version ranges as unavailable for new upgrades.
- **Minimal Version Selection (MVS)**: Algorithm to determine the set of module versions used for building.
- **Module Graph Pruning**: Optimization in Go 1.17+ to exclude transitive dependencies of modern modules from the graph unless needed by older modules.
- **Lazy Module Loading**: Strategy to load the full module graph only when required packages are not found among immediate requirements.

Procedures And API Details
- **Syntax for Replace**: `replace <module> [<version>] => <path|module> [version]`
  - Example: `replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5`
  - Grouping multiple directives is allowed within parentheses.
- **Syntax for Retract**: `retract <version>` or `retract [<lower>, <upper>]`
  - Example: `retract [v1.0.0, v1.9.9]`
  - Comments explaining the rationale can be added above or below the directive.
- **Automatic Update Command**: `go get -mod=mod` or similar commands with `-mod=mod` flag to rewrite `go.mod`.
- **Inspecting Build List**: `go list -m all` displays the selected versions determined by MVS.

Nuance Or Contradictions
- **Replace vs Require**: A `replace` directive alone does not add a module to the dependency graph; a corresponding `require` directive is still necessary to include the replaced module in the build list.
- **Retraction Visibility**: Retracted versions are hidden from `go list -m -versions` unless the `-retracted` flag is used.
- **Pruning Behavior**: Modules with pruned requirements still appear in `go list -m all`, but their packages cannot be directly imported into `go build` or `go test` without promoting them to explicit dependencies via `go get`.

Candidate Wiki Hints
- **Page: Go Modules Replace Directive** (Concept: Modifying the module graph)
- **Page: Go Modules Retract Directive** (Concept: Version lifecycle management)
- **Page: Minimal Version Selection in Go** (Concept: Dependency resolution algorithm)
- **Page: Module Graph Pruning and Lazy Loading** (Concept: Go 1.17+ optimizations)
