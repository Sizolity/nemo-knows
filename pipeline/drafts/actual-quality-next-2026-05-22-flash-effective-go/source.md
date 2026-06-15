---
title: Effective Go
kind: source
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## What It Is

Effective Go is a foundational (though aging) guide to writing idiomatic, clear, and performant Go programs. It supplements the Go specification, the Tour of Go, and How to Write Go Code, offering a concentrated tour of the language’s conventions and best practices.

## Summary

The document walks through the major areas that shape idiomatic Go: automated formatting with `gofmt`, the naming conventions that make packages, types, and methods expressive, automatic semicolon insertion and its effect on brace placement, the “no-parentheses, always braces” control structures (`if`, `for`, `switch`), and the use of multiple return values instead of in-band error codes or out-parameters. It details how `new` and `make` differ, explains arrays as value types versus slices as references, and shows maps, composite literals, and the `append` built-in. Printing, `iota`-based constants, and `init` functions are covered. Methods on named types, the rules for pointer vs. value receivers, and implicit interface satisfaction lead into embedding, type assertions, and the use of the blank identifier for compile-time interface checks, silencing unused imports/variables, and import‑for‑side‑effect. Concurrency is introduced through the “share by communicating” principle, goroutines, channels (unbuffered/buffered), and `select`. Error handling is described as a core convention using the `error` interface, with `panic`/`recover` reserved for truly unrecoverable conditions and occasionally used inside packages to unwind parsing errors. The document closes with a complete QR‑generating web server that exemplifies server setup, templates, and the expressiveness achievable in a few lines of Go.

## Key Claims

- `gofmt` (or `go fmt`) should be trusted to handle formatting; manual layout debates are unnecessary.
- Exported identifiers begin with an uppercase letter; package names should be short, lowercase, single‑word, and mirror the source directory name.
- Getters should be named like the unexported field (e.g., `Owner()` not `GetOwner()`); one‑method interfaces typically end with `‑er`.
- Semicolons are elided by the lexer; braces cannot appear on a new line in control structures.
- Control structures require no parentheses and mandate braces; `if`/`switch` accept an initialisation statement.
- Multiple return values are idiomatic for results and errors; named result parameters can simplify returns.
- `defer` ensures cleanup on function exit and is function‑scoped, not block‑scoped.
- `new` returns a pointer to zeroed memory, while `make` is only for slices, maps, and channels and returns an initialised (not pointer) value.
- Arrays are values; slices are descriptors with pointers to underlying arrays and are the primary mechanism for working with sequences.
- Implicit interface satisfaction means any type that has the required methods implements the interface; embedding structs promotes their methods and fields.
- Concurrency is modelled around goroutines and channels; “do not communicate by sharing memory; instead, share memory by communicating.”
- Errors are values; callers should always check error returns. `panic`/`recover` should be used sparingly, often within a package to translate panics into errors.

## Suggested Links

- [Effective Go](https://go.dev/doc/effective_go)
