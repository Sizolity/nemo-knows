---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
Lines 861–1255 of `raw/web/corpus-2026-05-18/032-go-modules-reference.md`, covering the "Go Modules Reference" section. Topics include `replace` directives, `retract` directives, automatic updates via `-mod=mod`, Minimal Version Selection (MVS), module graph modifications (replacement, exclusion, upgrades/downgrades), module graph pruning for Go 1.17+, and workspaces.

Local Summary
This chunk details how Go manages dependency resolution through the module graph. It explains syntax and behavior for replacing modules locally or remotely, retracting versions to prevent usage, and using `-mod=mod` to auto-correct `go.mod`. It describes Minimal Version Selection (MVS), which deterministically computes a build list from the main module's requirements, traversing the graph while respecting exclusions, replacements, and upgrades. Later sections cover optimizations like module graph pruning in Go 1.17+ and lazy loading, followed by workspace definitions for managing multiple main modules.

Key Claims
- A `replace` directive substitutes a specific module version or all versions with a local file path or another remote module path.
- A `retract` directive marks versions as unusable for automatic upgrades while keeping them accessible in repositories.
- The `-mod=mod` flag instructs the Go command to automatically rewrite non-canonical versions, respect exclusions, remove redundant requirements, and reformat `go.mod`.
- Minimal Version Selection (MVS) is deterministic and recalculated at every module-aware command execution.
- Module graph pruning in Go 1.17+ limits the loaded graph to immediate dependencies for high-GO-version modules unless transitively required by lower versions.
- Workspaces (`go.work`) allow running MVS across multiple main modules defined by relative paths in a `go.work` file.

Entities And Concepts
- `replace` directive: Swaps a module version/path with another.
- `retract` directive: Marks versions as deprecated/unusable for upgrades.
- `-mod=mod`: Flag enabling automatic correction of `go.mod`.
- Minimal Version Selection (MVS): Algorithm to compute the build list.
- Module graph: Directed graph of module versions and dependencies.
- Exclusion (`exclude`): Removes specific versions from the graph.
- Upgrade/Downgrade: Modifying the graph to prefer higher/lower versions.
- Module graph pruning: Optimization in Go 1.17+ to limit loaded graph size.
- Lazy module loading: On-demand loading of the full graph only when necessary.
- Workspace (`go.work`): Collection of modules for multi-root builds.

Procedures And API Details
- **Replace Directive Syntax**: `replace <ModulePath> [Version] => <FilePath | ModulePath Version>`
  - Example: `replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5`
  - Multi-line block syntax allowed using parentheses.
- **Retract Directive Syntax**: `retract (Version | [Version, Version])`
  - Comments may precede or follow the directive to explain rationale.
  - Example: `retract v1.0.0 // Published accidentally.`
- **Automatic Update Flow**: Running commands with `-mod=mod` triggers rewriting of non-canonical versions (e.g., `v1` → `v1.0.0`) and resolving exclusions.
- **MVS Execution**: Traverses the graph starting from main modules, tracking highest required versions per module to produce the build list.
- **Pruning Logic**: For Go 1.17+, only immediate dependencies are loaded unless a lower-GO-version module transitively requires them.

Nuance Or Contradictions
- `replace` directives alone do not add a module to the graph; a corresponding `require` directive is still necessary.
- In Go 1.15 and earlier, `-mod=mod` was default; since Go 1.16, the command defaults to readonly mode unless explicitly invoked with `-mod=mod`.
- Retracted versions remain available but are excluded from version lists (`go list -m -versions`) and resolution queries like `@latest` unless the `-retracted` flag is used.
- Module graph pruning affects which checksums are recorded in `go.sum` for a given Go version, potentially requiring the `-compat` flag to adjust behavior.

Candidate Wiki Hints
- **Page**: `go-modules-replace-directive` (Concept: Local/Remote module substitution)
- **Page**: `go-modules-retract-directive` (Concept: Version retraction and lifecycle management)
- **Page**: `go-modules-minimal-version-selection` (Concept: MVS algorithm and build list computation)
- **Page**: `go-modules-graph-pruning` (Concept: Optimization in Go 1.17+ for dependency loading)
