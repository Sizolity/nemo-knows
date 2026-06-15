---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
Heading path: Go Modules Reference > Retrieved Text
Lines: 1666-1685

Local Summary
This section explains how the `go` command behaves when vendoring is enabled versus GOPATH mode. It clarifies that build and test commands use the vendor directory at the main module's root, while module management commands like `go mod download`, `go mod tidy`, and `go get` continue to interact with the network and module cache regardless of vendoring status. A key distinction is made regarding where vendor directories are ignored outside the main module's root.

Key Claims
- When vendoring is enabled, build commands (`go build`, `go test`) load packages from the vendor directory instead of accessing the network or local module cache.
- The `go list -m` command only prints information about modules listed in `go.mod`.
- Module management commands (`go mod download`, `go mod tidy`) function identically when vendoring is enabled; they still download modules and access the module cache.
- The `go get` command does not behave differently when vendoring is enabled.
- Unlike GOPATH mode, the `go` command ignores vendor directories located outside the main module’s root directory.
- Because vendor directories in other modules are not used during builds, they are excluded from generated module zip files (with references to known bugs #31562 and #37397).

Entities And Concepts
- Vendoring: A practice of copying dependencies into a local directory to avoid network access.
- Vendor Directory: The directory containing vendored packages; only the one in the main module's root is respected by build commands.
- Module Cache: The local storage for downloaded modules, accessed even when vendoring is active for management tasks.
- Main Module’s Root Directory: The specific location where vendor directories are recognized and utilized.
- Module Zip Files: Archive files generated during builds that exclude unused vendor directories from submodules.

Procedures And API Details
Command Usage:
- `go build`: Loads packages from the vendor directory if enabled; otherwise, uses network/cache.
- `go test`: Same behavior as `go build` regarding vendoring.
- `go list -m`: Prints module info strictly for those listed in `go.mod`.
- `go mod download`: Downloads modules and accesses cache regardless of vendoring status.
- `go mod tidy`: Cleans up `go.mod` and `go.sum` without being affected by vendoring status.
- `go get`: Does not change behavior when vendoring is enabled; usage includes flags like `-d`, `-t`, `-u`, `-tool`.

Nuance Or Contradictions
- **Vendoring Scope**: While `go build` uses the vendor directory at the main module's root, it explicitly ignores vendor directories in other modules. This means sub-dependencies with their own vendor folders are not utilized for building the main module.
- **Module Zip Files**: Building a module zip file excludes vendor directories found in non-main modules due to the logic described above, though this behavior is noted as having known bugs (#31562, #37397).
- **Management Commands vs Build Commands**: There is a functional divergence; build commands respect the main vendor directory, but management commands (`go mod`, `go get`) bypass vendoring to ensure dependencies are up-to-date in the cache.

Candidate Wiki Hints
- Create a page explaining the scope of vendor directories in Go modules, specifically distinguishing between the main module's root and submodules.
- Document the behavior differences between build-time loading and module-management operations regarding vendoring.
