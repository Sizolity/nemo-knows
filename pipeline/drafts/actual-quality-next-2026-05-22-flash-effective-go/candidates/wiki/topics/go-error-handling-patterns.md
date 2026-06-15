---
title: Go Error Handling Patterns
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Error Handling Patterns

Go error handling is built around the `error` interface and multiple return values. Functions that can fail return a result together with an error; callers check the error before using the result. This convention eliminates in‑band error codes and out‑parameters, making the control flow explicit.

Several recurring patterns emerge from this design:

- **Multiple return values** – The signature `(result, error)` is idiomatic for any operation that may fail. Examples include `Write`, which returns bytes written and an error when the buffer could not be fully written. This pattern replaces C‑style negative counts and volatile error locations.

- **Initialisation statement in `if`** – A short variable declaration inside an `if` statement scopes the error variable to the condition. For instance, `if err := file.Chmod(0664); err != nil { … }` sets up a local error and immediately handles it. This keeps the error‑handling logic close to the operation that produced it.

- **Omitting `else` after a terminal body** – When an `if` body ends in `return`, `break`, `continue`, or `goto`, the unnecessary `else` is dropped. The happy path continues at the next line, avoiding nested blocks. The result is a linear sequence of guard clauses that read top to bottom, with each error case returning early.

- **Guard clause sequences** – Successive operations each check their own error and return immediately on failure. An example pattern: `f, err := os.Open(…)`; if error, return; then `d, err := f.Stat()`; if error, close and return. No `else` is needed, and the successful flow runs straight down the page.

- **`defer` for cleanup** – Idiomatic Go uses `defer` to schedule resource release immediately before the function returns, regardless of which error path is taken. Closing files, unlocking mutexes, and similar cleanup are commonly deferred right after acquisition. This centralises release code and avoids repetition.

- **Named result parameters** – Naming result parameters can clarify which value is which and permit a bare `return` that automatically returns the current values of those parameters. In error‑handling code, functions like `ReadFull` use named results to accumulate a count and an error, then rely on a bare return to deliver the final state without extra variables.

- **`panic` and `recover`** – These mechanisms exist for truly unrecoverable conditions and should be used sparingly. They are occasionally employed within a package to unwind deep call stacks—converting an internal panic into a returned error—but the preferred approach for expected failures is always a normal error value.

Because errors are plain values, callers are free to inspect, wrap, or propagate them. The conventions described here keep error handling explicit, composable, and in step with the principle that the successful path is the most readable.
