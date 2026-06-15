---
title: Effective Go
kind: source
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## What It Is

**Effective Go** is an authoritative guide detailing idiomatic usage, style conventions, and language patterns for the Go programming language. Hosted at [https://go.dev/doc/effective_go](https://go.dev/doc/effective_go), it serves as a reference for writing clear, consistent, and efficient Go code. The document covers everything from basic formatting rules to advanced topics like concurrency models, interface embedding, and panic recovery patterns.

## Summary

The guide is structured to introduce readers to the "Go way" of doing things. It begins with foundational rules on formatting (relying on `gofmt`), naming conventions (package names, getters, interfaces), and syntax specifics (control structures, semicolons). As it progresses, it delves into data structures like maps and slices, explaining their semantics (value vs. reference) and manipulation techniques (`append`, composite literals). The middle sections explore advanced type system features, including multiple return values with named results, `defer` mechanics, interface polymorphism, and compile-time contract enforcement via `var _ Interface = Value`.

The latter parts of the document focus on concurrency ("share by communicating") using goroutines and channels, error handling conventions (returning `error` interfaces), and the specific usage of `panic` and `recover`. It concludes with practical examples, such as building a QR code web server using `html/template`, illustrating how these idioms combine in real-world applications.

## Key Claims

## Style and Formatting
- Programs should be formatted by `gofmt`; manual formatting is discouraged. Use tabs for indentation and avoid arbitrary line length limits.
- Package names should be lowercase, single words, and concise (e.g., `bytes`).
- Getters should return capitalized field names (e.g., `Owner`) rather than using the `Get` prefix.
- Interface names often follow the `<Method>-er` suffix pattern (e.g., `Reader`, `Writer`).

## Syntax and Control Flow
- Semicolons are automatically inserted by the lexer; they rarely appear in source code but are required in specific contexts like `for` loop clauses.
- Control structures (`if`, `switch`) do not require parentheses around their expressions or conditions.
- The `else` clause is often omitted after an `if` statement that returns an error to allow successful flow to continue naturally.
- Short declarations (`:=`) are used for declaring and initializing variables in a single statement, particularly for reassigning existing variables like errors.
- The blank identifier `_` is used to ignore values in range loops or discard unused return values.

## Data Structures and Memory
- Maps associate keys (any equality-defined type) with values; slices cannot be keys. Missing map keys return the zero value of the entry type.
- The "comma ok" idiom (`val, ok := map[key]`) distinguishes between zero values and missing entries.
- `new(T)` allocates zeroed storage returning a pointer; `make(T, args)` creates initialized slices, maps, or channels.
- Arrays are values (copy on assignment), whereas slices are references to underlying arrays.
- Two-dimensional slices are implemented as slices of slices, allowing independent inner lengths or shared underlying arrays.

## Interfaces and Types
- Interface implementation is implicit; a type satisfies an interface if it implements the required methods.
- Compile-time contract enforcement can be achieved via `var _ Interface = Value`.
- Embedding structs or interfaces automatically promotes their methods without manual forwarding.
- Type switches (`switch t := t.(type)`) are used to discover dynamic types, while type assertions extract specific types from interfaces.

## Concurrency and Errors
- Go encourages avoiding shared mutable state in favor of passing data through channels ("share by communicating").
- Goroutines are lightweight functions managed by the runtime; unbuffered channels act as rendezvous points, while buffered ones limit concurrency.
- Errors are typically returned as second return values of type `error`, which implements a single `Error() string` method.
- `panic` stops execution immediately unless recovered from inside a deferred function; it is intended for unrecoverable errors or impossible conditions.
- Library functions should avoid panicking unless during initialization, preferring to mask problems to keep programs running.

## Suggested Links

[https://go.dev/doc/effective_go](https://go.dev/doc/effective_go)
