---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
## Chunk Context
This chunk details the `go mod edit` command for editing `go.mod` files, provides its JSON output schema via `-json`, and describes `go mod graph`, `go mod init`, `go mod tidy`, `go mod vendor`, `go mod verify`, and `go mod why`. It concludes with the `go version -m` command for inspecting executable build versions.

## Local Summary
The section begins by explaining that `go mod edit -json` outputs the `go.mod` file as a JSON structure representing the module, Go version, requirements, exclusions, replacements, retractions, and tools. The text defines the associated Go types (`Module`, `GoMod`, `ModPath`, `Require`, `Replace`, `Retract`, `Tool`). It notes that this schema describes only the `go.mod` file itself, not indirect modules, directing users to `go list -m -json all` for the full set. Subsequent subsections cover `go mod graph` (printing the module requirement graph), `go mod init` (creating a new module), `go mod tidy` (syncing go.mod with imports), `go mod vendor` (copying dependencies to a vendor directory), `go mod verify` (checking integrity of cached modules), and `go mod why` (showing import paths). The chunk ends with `go version -m`, which prints the Go version used to build an executable, optionally including module versions.

## Key Claims
- `go mod edit` reads only one `go.mod` file and does not look up information about other modules.
- `go mod edit -json` outputs a JSON structure corresponding to specific Go types that describe the `go.mod` file itself, excluding indirect modules.
- For the full set of modules available to a build (including indirect ones), use `go list -m -json all`.
- `go mod graph` prints the module requirement graph with replacements applied in text form.
- `go mod init` creates a new `go.mod` file in the current directory if it does not already exist.
- `go mod tidy` ensures `go.mod` matches source code, adding missing requirements and removing unused ones.
- `go mod vendor` constructs a `vendor` directory containing copies of needed packages and generates `vendor/modules.txt`.
- `go mod verify` checks that dependencies in the module cache have not been modified since download.
- `go mod why` shows a shortest path in the import graph from the main module to listed packages or modules.
- `go version -m` prints the Go version and module versions used to build a specific executable.

## Entities And Concepts
- **Commands**: `go mod edit`, `go list`, `go mod graph`, `go mod init`, `go mod tidy`, `go mod vendor`, `go mod verify`, `go mod why`, `go version`.
- **Files**: `go.mod`, `vendor/modules.txt`.
- **Directives/Flags**: `-json`, `-fmt`, `-print`, `-require`, `-droprequire`, `-exclude`, `-dropexclude`, `-replace`, `-dropreplace`, `-retract`, `-dropretract`, `-tool`, `-droptool`, `-e`, `-v`, `-x`, `-diff`, `-go`, `-compat`, `-o`, `-m`, `-vendor`.
- **Types**: `Module`, `GoMod`, `ModPath`, `Require`, `Replace`, `Retract`, `Tool`.
- **Build Tags**: `ignore`.

## Procedures And API Details
- **Printing JSON Representation of go.mod**:
  - Command: `$ go mod edit -json`
  - Output corresponds to Go types:
    ```go
    type Module struct { Path string; Version string }
    type GoMod struct { Module ModPath; Go string; Require []Require; Exclude []Module; Replace []Replace; Retract []Retract }
    type ModPath struct { Path string; Deprecated string }
    type Require struct { Path string; Version string; Indirect bool }
    type Replace struct { Old Module; New Module }
    type Retract struct { Low string; High string; Rationale string }
    type Tool struct { Path string }
    ```
- **Printing Go version used to build an executable**:
  - Command: `$ go version ~/go/bin/gopls`
- **Printing Go version and module versions used to build an executable**:
  - Command: `$ go version -m ~/go/bin/gopls`

## Nuance Or Contradictions
- `go mod edit` does not look up information about other modules; it only reads and writes the specified target file.
- The JSON output from `go mod edit -json` describes only the `go.mod` file itself, not other modules referred to indirectly.
- `go mod tidy` acts as if all build tags are enabled except for the `ignore` tag.
- `go mod vendor` removes the existing `vendor` directory before re-constructing it.
- After Go 1.16 in modules declaring `go 1.16` or higher, the `-vendor` flag of `go mod why` has no effect because the meaning of `all` changed to match the set of packages matched by `go mod vendor`.

## Candidate Wiki Hints
- **Page**: Go Module Editing Commands (`go mod edit`)
  - Covers editing flags (`-module`, `-require`, `-replace`, etc.) and output control (`-json`, `-fmt`).
- **Page**: Inspecting Module Graphs and Versions (`go mod graph`, `go version -m`)
  - Explains `go mod graph` output format and `go version -m` usage for build inspection.
