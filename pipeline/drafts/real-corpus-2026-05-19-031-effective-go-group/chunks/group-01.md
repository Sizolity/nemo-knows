---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

Group Context
- **Source Document**: Effective Go (from https://go.dev/doc/effective_go).
- **Scope**: Covers the complete style and language idiom guide, ranging from initial metadata retrieval to detailed discussions on formatting, control structures, memory allocation, maps, interfaces, concurrency models, and error handling.
- **Chunk Range**: Lines 1–2730 (Chunks 01–07).
- **Primary Heading Path**: Effective Go > Retrieved Text (with introductory sections in Chunk 01).

Cross-Chunk Summary
The document serves as a canonical guide for writing idiomatic Go code, emphasizing simplicity, reliability, and efficiency. It explicitly advises against direct translations of C++ or Java. The content flows logically from basic formatting conventions (`gofmt`) and naming rules to the mechanics of control flow (`if`, `for`, `switch`), memory management (`new` vs `make`, slices, maps), type system nuances (interfaces, embedding, type assertions), and finally concurrency patterns ("sharing memory by communicating") and robust error handling.

Repeated Or Central Claims
- **Idiomatic Style**: Go is designed for building software at scale; code should be simple and avoid complex idioms from other languages. Formatting is automatic via `gofmt`.
- **Naming Conventions**: Package names are short, lower-case, and single-word. Interface names follow the `<method>-er` pattern (e.g., `Reader`). Multiword identifiers use MixedCaps. Getters should be methods (e.g., `Owner()`) rather than prefixed fields (`GetOwner()`).
- **Control Structures**: Semicolons are implicit; braces must always delimit bodies. `switch` supports non-constant expressions and omitting fall-through logic. `for` loops unify C-style and range iterations.
- **Memory & Collections**: Use `make` for slices/maps/channels and `new` for pointers to zero values. Slices wrap arrays and pass references; use `_` to discard unwanted range values.
- **Interfaces**: Interfaces are sets of methods. A type implements an interface if it has the required methods (duck typing). The blank identifier `_` is used for compile-time checks (`var _ Type = value`) and discarding values/errors, though discarding errors in practice is discouraged.
- **Concurrency**: Avoid shared mutable state; use channels to synchronize goroutines ("sharing memory by communicating"). Channels provide synchronization primitives (blocking send/receive) that replace locks.

Important Local Details
- **Formatting**: `gofmt` aligns columns and uses tabs for indentation. Comments should be aligned vertically if they are part of a block.
- **Allocation**: `new(T)` returns `*T` with zeroed storage. `make(T, args)` initializes the underlying array/map/channel. Since Go 1.26, `new` accepts an initial value expression.
- **Maps**: Map lookups return `(value, ok)`. Use the "comma ok" idiom to check for existence. `delete(map, key)` removes entries safely even if absent. Keys must be comparable types; slices cannot be keys.
- **Printing**: `fmt.Printf` uses `%v`, `%+v` (struct fields), and `%#v` (Go syntax). Custom types should define a `String()` method to control output, avoiding recursion that prints the receiver as a string.
- **Methods & Receivers**: Value receivers work on both values and pointers; pointer receivers modify the original. Addressable values allow calling pointer methods without explicit `&`. Embedding promotes methods but retains the inner receiver type unless changed.
- **Errors & Panic**: Errors follow the convention `type error interface { Error() string }`. `panic` halts execution for unrecoverable states. Do not recover from panics in production code typically.
- **Concurrency Primitives**: Goroutines are lightweight threads. Channels can be buffered or unbuffered. Semaphore patterns use buffered channels to limit concurrency. Parallelism is achieved by breaking work into independent pieces executed across cores.

Candidate Wiki Hints
- **Style Guide**: `Effective Go` (Main Page)
- **Formatting**: `gofmt`, Indentation, Line Length, Comments
- **Naming**: Package Names, Getters/Setters, Interface Naming (`-er`), MixedCaps
- **Control Flow**: Semicolon Insertion, `if`/`switch`/`for` Logic, Blank Identifier Usage
- **Memory**: `new` vs `make`, Slices (Capacity, Length, Appending), Maps (Zero Value, Deletion)
- **Types & Interfaces**: Type Assertions, Embedding, Method Sets, Compile-Time Checks (`var _`)
- **Concurrency**: Goroutines, Channels (Buffered/Unbuffered), Semaphore Pattern, Parallelism vs Concurrency
- **Error Handling**: Error Interface Convention, `panic` Usage, Recovering from Errors

Gaps Or Cautions
- **Metadata Only Chunk**: Chunk 01 contains only retrieval metadata and navigation headers; it does not contain substantive code or style rules.
- **Implicit Semicolons**: Be cautious of implicit semicolon insertion rules (e.g., after `if`, `return`, `break`), especially when writing multi-statement lines or placing braces on new lines.
- **Error Discarding**: Using `_` to discard errors (e.g., in file operations) is a dangerous practice that can lead to runtime crashes; always check error conditions unless the context guarantees safety.
- **Interface Satisfaction**: While most interface satisfaction is checked at compile time, some packages (like `encoding/json`) may require run-time checks or specific type assertions due to how they define their interfaces.
- **Concurrency Bugs**: Historically, loop variables in goroutines (`range`) could cause data races; ensure variables are captured correctly (e.g., using an index variable) to avoid shared state issues.
