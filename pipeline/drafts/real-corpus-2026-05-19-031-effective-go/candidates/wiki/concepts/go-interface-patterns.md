---
title: Go Interface Patterns
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Interface Patterns

In the Go programming language, interface implementation is implicit. A type satisfies an interface automatically if it implements all the required methods; there are no explicit declaration statements needed for this relationship. This design encourages developers to focus on behavior rather than naming contracts in the source code.

To enforce compile-time contract verification without relying solely on implicit checks, developers can use the blank identifier syntax: `var _ Interface = Value`. This forces the compiler to verify that the specific type satisfies the interface at build time.

## Method Promotion and Embedding

Structs and interfaces support embedding, which automatically promotes their methods to the outer type without requiring manual forwarding. When a struct embeds an interface or another struct, it inherits the methods defined in those embedded types. This pattern allows for flexible composition of behavior.

## Type Discovery and Assertions

Go provides mechanisms for working with dynamic types through the type system:
- **Type Switches**: Using `switch t := t.(type)` allows developers to discover the dynamic concrete type stored within an interface variable.
- **Type Assertions**: These are used to extract specific concrete types from an interface, allowing access to underlying data or methods associated with that specific type.

## Error Handling and Panic Recovery

Error handling in Go typically involves returning errors as the second return value of a function, where the `error` type implements a single `Error() string` method. While `panic` is intended for unrecoverable errors or impossible conditions, it can be combined with `recover` inside deferred functions to handle runtime panics gracefully.

## Style and Formatting

When writing interface definitions, adherence to style conventions improves code readability:
- Interface names often follow the `<Method>-er` suffix pattern (e.g., `Reader`, `Writer`).
- Programs should be formatted using `gofmt`.
- Package names should be lowercase, single words, and concise.
