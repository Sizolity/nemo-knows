---
title: Go Error Handling Panic Recovery
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Error Handling Panic Recovery

In Go, distinguishing between recoverable errors and unrecoverable conditions is a fundamental aspect of robust application design. The language provides distinct mechanisms for handling both: returning `error` interfaces for standard failures and using `panic` with `recover` for exceptional states that require immediate interruption or cleanup.

## Standard Error Handling

The idiomatic approach to handling failures in Go is to return values of type `error`. Functions typically return a second value of this type, allowing the caller to check for success before proceeding. This pattern encourages programs to continue running even when minor issues occur, masking problems rather than crashing immediately.

When dealing with errors returned from functions:
- The `else` clause is often omitted after an `if` statement checking for an error to allow successful flow to continue naturally.
- Short declarations (`:=`) are frequently used to reassign variables like errors within loops or conditionals.
- The blank identifier `_` is used to ignore return values when they are not needed.

For cases where a map lookup might fail, the "comma ok" idiom (`val, ok := map[key]`) distinguishes between zero values and missing entries without panicking.

## Panic Recovery Mechanisms

While standard errors should be handled gracefully, `panic` is reserved for unrecoverable errors or impossible conditions. When a function calls `panic`, execution stops immediately unless intercepted by a `recover` call within a deferred function.

### The `defer-recover` Pattern

To catch a panic and allow the program to continue (or exit gracefully), developers often use the `defer` statement in combination with `recover`:

```go
func recover() {
    defer func() {
        if r := recover(); r != nil {
            // Handle the panic here
        }
    }()
}
```

This pattern ensures that cleanup logic runs even when a panic occurs, preventing resource leaks or leaving the application in an inconsistent state. However, `panic` should generally be avoided in library functions unless during initialization, as it can disrupt the caller's control flow unexpectedly.

## Guidelines for Usage

- **Prefer error returns:** Use `error` interfaces to signal conditions that callers can handle and recover from.
- **Use panic sparingly:** Reserve `panic` for truly exceptional situations like programming errors or impossible states.
- **Avoid shared mutable state in concurrent code:** Go encourages avoiding shared mutable state in favor of passing data through channels ("share by communicating").
- **Follow formatting rules:** Programs should be formatted by `gofmt`; manual formatting is discouraged.

## Related Concepts

- [[go-control-flow-idioms]]
- [[go-concurrency-model]]
