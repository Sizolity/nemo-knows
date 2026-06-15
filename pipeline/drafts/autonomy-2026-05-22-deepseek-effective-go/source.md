---
title: Effective Go
kind: source
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## What It Is

A classic guide to writing clear, idiomatic Go code.  Written for Go’s initial release in 2009, it remains a useful resource for the core language, though it does not cover generics, modules, or standard library additions made after that time.  The document augments the Go specification, the Tour of Go, and How to Write Go Code.

## Summary

Effective Go walks through the language’s philosophy, formatting, naming, control structures, functions, data types, methods, interfaces, concurrency, and error handling.  It emphasises Go-specific conventions (e.g., `gofmt` for formatting, MixedCaps naming, `defer` for cleanup, goroutines and channels for concurrency) and illustrates them with concise examples drawn from the standard library.  A complete web‑server example at the end demonstrates how the idioms combine into a real program.

## Key Claims

- **Formatting is mechanical, not personal.**  The `gofmt` tool (or `go fmt`) canonicalises layout; let the machine decide.
- **Names matter.**  Package names should be short, lower‑case, single‑word.  Exported names use `MixedCaps`; getters omit “Get” (e.g., `Owner()` not `GetOwner()`).  One‑method interfaces often end in `-er` (e.g., `Reader`, `Writer`).
- **Semicolons are inserted by the lexer**, so opening braces must stay on the same line as the control structure (`if`, `for`, etc.).
- **Control structures are streamlined.**  `if`, `switch`, and `for` can include an initialisation statement; `switch` evaluates cases top‑to‑bottom with no automatic fall‑through; `for` is the only loop.
- **Functions return multiple values** (typically a result and an `error`); named result parameters can clarify intent and simplify bare `return`.
- **`defer`** schedules a function to run just before the enclosing function returns, useful for paired operations like open/close and lock/unlock.
- **`new` vs. `make`:** `new(T)` allocates zeroed memory and returns `*T`; `make` creates slices, maps, and channels and returns an initialised (not zeroed) `T`.
- **Slices are the idiomatic sequence type.**  Arrays are values (copied on assignment); passing a slice gives reference behaviour.
- **Interfaces are satisfied implicitly.**  A type need not declare its intent; a constructor often returns an interface value to hide the concrete type.
- **Embedding** (in structs or interfaces) promotes methods and provides a form of composition instead of subclassing.
- **Concurrency mantra:** “Do not communicate by sharing memory; instead, share memory by communicating.”  Goroutines are lightweight; channels coordinate them safely.
- **Errors are values**, returned as the last value; the `error` interface is simple.  `panic` and `recover` handle truly exceptional situations but should be used sparingly, often only inside libraries to translate panics into errors.
- **The blank identifier (`_`)** discards unused values, silences compiler complaints during development, enables side‑effect‑only imports, and provides compile‑time interface satisfaction checks.

## Suggested Links

- [Effective Go](https://go.dev/doc/effective_go)
- Go specification (referred to in the text)
- Tour of Go (referred to in the text)
- How to Write Go Code (referred to in the text)
- Go release notes (referred to in the text)
