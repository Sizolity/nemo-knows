---
title: Go Error Handling
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Error Handling

Go treats errors as ordinary values rather than exceptions. A function that can fail returns its result alongside an `error`; by convention the error is the last return value. This design keeps error paths explicit and linear.

Callers check the error immediately with a compact `if` statement that often includes a short assignment. The body of the block typically ends with a `return` or a cleanup action, allowing the “happy path” to continue unindented and avoiding redundant `else` clauses.

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

The pattern scales to sequences of fallible operations: each step guards against an error, and when an error occurs the function returns or defers resource release so that later code only runs after all checks pass. The `defer` statement schedules cleanup—such as closing a file or unlocking a mutex—that runs regardless of how a function exits, making error-handling code both safe and concise.

Named result parameters can further clarify an error-bearing signature. Because named results are initialised to their zero values and returned automatically by an unadorned `return`, functions like `io.ReadFull` can accumulate partial results and an error without repeating values.

The built-in `error` interface (a single `Error() string` method) is the universal contract. Custom error types satisfy it implicitly, and callers can inspect the concrete error with a type switch or by extracting additional context. The standard library’s `Write` method on `*os.File`, for instance, returns an `(int, error)` pair where a non‑nil error indicates that fewer bytes were written than requested.

Panic and recover are reserved for truly unrecoverable situations, not for routine error flow. A deferred function can call `recover` to intercept a panic and turn it into a normal error return, a technique used sparingly inside packages to avoid exposing panics to clients.

The insistence on explicit, value‑based error checking is a defining [[go-idioms|Go idiom]], shaping the structure of nearly every function that can fail.
