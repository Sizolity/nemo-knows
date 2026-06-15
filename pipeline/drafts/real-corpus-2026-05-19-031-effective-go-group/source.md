---
title: Effective Go
kind: source
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## What It Is

"Effective Go" is a canonical style and language idiom guide for the Go programming language, originally written for the 2009 release. It serves as a reference for writing clear, idiomatic code that prioritizes simplicity, reliability, and efficiency. The document covers everything from basic formatting conventions to advanced concurrency patterns and error handling strategies, explicitly advising against direct translations of C++ or Java.

## Summary

The guide flows logically from introductory metadata to detailed discussions on formatting (using `gofmt`), naming rules, control flow structures (`if`, `for`, `switch`), memory management (`new` vs `make`, slices, maps), type system nuances (interfaces, embedding), and finally concurrency models ("sharing memory by communicating") and robust error handling. It concludes with practical examples of safe initialization, panic recovery patterns, and a complete web server implementation using HTML templates.

## Key Claims

### Idiomatic Style & Formatting
- Go is designed for building software at scale; code should be simple and avoid complex idioms from other languages.
- Formatting issues are handled automatically by `gofmt` (or `go fmt`), which aligns columns and manages indentation using tabs. Comments should be aligned vertically if part of a block.

### Naming Conventions
- **Package Names**: Should be short, concise, evocative, lower-case, and single-word (e.g., `bufio.Reader`, not `BufReader`).
- **Getters/Methods**: Use PascalCase methods (e.g., `Owner()`) to expose fields; avoid `Get` prefixes. Getters should be methods rather than prefixed fields.
- **Interface Names**: Follow the `<method>-er` pattern for agent nouns (e.g., `Reader`, `Writer`).
- **Multiword Identifiers**: Use MixedCaps (or mixedCaps), not underscores.

### Control Structures & Logic
- Semicolons are implicitly inserted by the lexer after specific tokens before newlines, except before closing braces.
- `if` and `switch` statements support optional initialization; bodies must always be brace-delimited.
- The `for` loop unifies C-style `for`, `while`, and infinite loops; range clauses manage iteration over arrays, slices, strings, and maps.
- `switch` is more flexible than C's, supporting non-constant expressions and omitting fall-through logic (cases run top-to-bottom).

### Memory & Collections
- Use `make` for slices/maps/channels to initialize underlying data structures.
- Use `new(T)` for pointers to zero values; since Go 1.26, it accepts an initial value expression.
- Slices wrap arrays and pass references; use `_` to discard unwanted range values.
- Map lookups return `(value, ok)`. Use the "comma ok" idiom to check for existence. `delete(map, key)` removes entries safely even if absent.

### Interfaces & Types
- Interfaces are sets of methods. A type implements an interface if it has the required methods (duck typing).
- The blank identifier `_` is used for compile-time checks (`var _ Type = value`) and discarding values/errors, though discarding errors in practice is discouraged.
- Embedding promotes methods but retains the inner receiver type unless changed.

### Concurrency Models
- Avoid shared mutable state; use channels to synchronize goroutines ("sharing memory by communicating").
- Channels provide synchronization primitives (blocking send/receive) that replace locks.
- Goroutines are lightweight threads. Channels can be buffered or unbuffered. Semaphore patterns use buffered channels to limit concurrency.

### Error Handling & Panic
- Errors follow the convention `type error interface { Error() string }`.
- Library functions should prioritize masking problems or working around them rather than causing the entire program to halt.
- `panic` halts execution for unrecoverable states. `recover` is only effective within deferred functions during stack unwinding.
- Internal panics should be caught, converted into error values, and returned to callers.

### Web & Printing
- `fmt.Printf` uses `%v`, `%+v` (struct fields), and `%#v` (Go syntax). Custom types should define a `String()` method to control output.
- The `html/template` package automatically escapes data, ensuring it is safe for display in HTML contexts without additional manual escaping.

## Suggested Links

- [Effective Go](https://go.dev/doc/effective_go)
