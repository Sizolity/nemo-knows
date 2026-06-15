---
title: Go Formatting Rules
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Formatting Rules

The style and formatting of Go programs are governed by the **Effective Go** guide, which emphasizes consistency and automation. Programs should be formatted automatically using `gofmt`; manual formatting is discouraged. Indentation must use tabs rather than spaces, and arbitrary line length limits are avoided to keep code clean.

## Naming Conventions

Package names should be lowercase, single words, and concise (e.g., `bytes`). Getters for struct fields typically return capitalized field names (e.g., `Owner`) instead of using a `Get` prefix. Interface names often follow the `<Method>-er` suffix pattern, such as `Reader` or `Writer`.

## Syntax and Control Flow

Go's lexer automatically inserts semicolons, so they rarely appear explicitly in source code except in specific contexts like `for` loop clauses. Control structures like `if` and `switch` do not require parentheses around their expressions or conditions. The `else` clause is frequently omitted after an `if` statement that returns an error to allow successful flow to continue naturally. Short declarations (`:=`) are used for declaring and initializing variables in a single statement, particularly for reassigning existing variables like errors. The blank identifier `_` is used to ignore values in range loops or discard unused return values.

## Data Structures and Memory

Maps associate keys (any equality-defined type) with values; slices cannot be keys. Missing map keys return the zero value of the entry type. The "comma ok" idiom (`val, ok := map[key]`) distinguishes between zero values and missing entries. `new(T)` allocates zeroed storage returning a pointer, while `make(T, args)` creates initialized slices, maps, or channels. Arrays are values (copy on assignment), whereas slices are references to underlying arrays. Two-dimensional slices are implemented as slices of slices, allowing independent inner lengths or shared underlying arrays.

## Interfaces and Types

Interface implementation is implicit; a type satisfies an interface if it implements the required methods. Compile-time contract enforcement can be achieved via `var _ Interface = Value`. Embedding structs or interfaces automatically promotes their methods without manual forwarding. Type switches (`switch t := t.(type)`) are used to discover dynamic types, while type assertions extract specific types from interfaces.

## Concurrency and Errors

Go encourages avoiding shared mutable state in favor of passing data through channels ("share by communicating"). Goroutines are lightweight functions managed by the runtime; unbuffered channels act as rendezvous points, while buffered ones limit concurrency. Errors are typically returned as second return values of type `error`, which implements a single `Error() string` method. `panic` stops execution immediately unless recovered from inside a deferred function; it is intended for unrecoverable errors or impossible conditions. Library functions should avoid panicking unless during initialization, preferring to mask problems to keep programs running.
