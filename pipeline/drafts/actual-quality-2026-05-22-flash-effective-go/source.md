---
title: Effective Go
kind: source
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## What It Is
A classic style and idiom guide for writing clear, performant Go code.
Written for Go’s 2009 release and not actively updated; still valuable for core language usage, but does not cover generics, modules, or later library additions.

Source: [https://go.dev/doc/effective_go](https://go.dev/doc/effective_go).

## Summary
- **Formatting** – Standardised by `gofmt` (tabs, no parentheses around control structures, line‑length flexibility).
- **Naming** – Exported names begin with uppercase; packages are short, lower‑case, single‑word; getters omit `Get`; one‑method interfaces use `‑er` suffix; mixedCaps, not underscores.
- **Semicolons** – Inserted by the lexer; opening braces must stay on the same line.
- **Control structures** – `if`/`switch` accept init statements; `for` unifies `while`; `switch` without expression acts like `if‑else`; type switches test interface dynamic type.
- **Functions** – Multiple return values (commonly for results and errors); named result parameters; `defer` schedules cleanup at function exit.
- **Data** – `new` returns zeroed memory (Go 1.26+ allows initial value); `make` creates slices/maps/channels and initialises them; composite literals simplify construction; slices wrap arrays and are passed by value (descriptor); maps use “comma ok” idiom for missing keys.
- **Methods** – Receiver can be value or pointer; pointer methods can modify the receiver and satisfy interfaces.
- **Interfaces** – Implicit satisfaction; type assertions and type switches for dynamic checks; constructor often returns interface, not concrete type.
- **Embedding** – Structs embed other types, gaining their methods like composition; no subclassing.
- **Concurrency** – “Share memory by communicating”; goroutines are lightweight threads; channels provide synchronisation; buffered channels as semaphores; `select` for multiplexing.
- **Errors** – Idiomatic error handling: return `error` (built‑in interface); specific error types (e.g. `os.PathError`) enable type assertion for detail; use `panic`/`recover` only for truly unrecoverable situations or to simplify error propagation within a package.
- **Example** – A complete QR‑code web server demonstrates `net/http`, `html/template`, and flag parsing.

## Key Claims
- `gofmt` solves formatting style debates by enforcing a single canonical format.
- Visibility is determined by name casing: upper‑case first letter means exported.
- Go’s design prefers explicit error returns over exceptions; `defer` replaces manual resource‑cleanup blocks.
- Zero values should be usable whenever possible (e.g. `sync.Mutex`, `bytes.Buffer`).
- `new` vs `make`: `new` allocates zeroed memory and returns a pointer; `make` initialises slices, maps, and channels, returning the value (not a pointer).
- Slices, maps, and channels are reference types; assign them without copying underlying data.
- Concurrency is built around goroutines and channels; avoid sharing memory by communicating.
- Interface satisfaction is structural and implicit; compile‑time checks can be forced with blank identifier assignments.

## Suggested Links
- [Effective Go](https://go.dev/doc/effective_go)
- Related resources (not linked in source): Go language specification, Tour of Go, How to Write Go Code, `gofmt`, Go release notes, issue 28782
