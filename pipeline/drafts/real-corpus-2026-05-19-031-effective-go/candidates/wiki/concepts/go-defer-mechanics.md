---
title: Go Defer Mechanics
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Defer Mechanics

The `defer` statement in Go allows a function to schedule the execution of a function call, which occurs after the surrounding function returns. This mechanism is particularly useful for resource cleanup and ensuring that certain actions occur regardless of how a function exits, whether through normal completion or early return statements.

## Execution Order

Deferred calls are executed in the **Last In, First Out (LIFO)** order. When a function returns, all deferred calls scheduled during its execution are processed before the actual return happens. This ensures that nested defer statements execute in reverse order of their scheduling.

## Interaction with Return Values

When a `defer` statement captures a variable that is modified after the defer call but before the function returns, the value captured at the time of deferral remains unchanged. This behavior allows for predictable cleanup operations even when local variables are reassigned.

## Usage in Error Handling

Defer statements are commonly used with error handling patterns, often paired with `panic` and `recover`. They ensure that resources are released or errors are logged before a function panics or returns an error value.

## Common Patterns

- **Resource Cleanup**: Closing file handles, network connections, or mutex locks.
- **Stack Traces**: Capturing stack traces for debugging purposes.
- **Logging**: Ensuring log messages are written even if the function exits early.

For more on error handling and panic recovery patterns, see [[go-error-handling-panic-recovery]].
