---
title: Effective Go
kind: source
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## What It Is
A classic style guide for writing clear, idiomatic Go code. Originally published for Go’s 2009 release and not actively updated since, it still serves as a foundational reference on core language conventions—despite not covering later additions like generics or modules. It complements the language specification, the Tour of Go, and How to Write Go Code.

## Summary
The document walks through the key idioms that distinguish effective Go programs from direct translations of C++ or Java. It emphasises letting `gofmt` handle formatting, using naming rules (lowercase package names, exported names begin with upper case, no `Get` prefix), and understanding semicolon insertion so that braces stay on the same line as control statements. Control structures (if, for, switch, type switch) are presented with idiomatic patterns such as omitting `else` when the body exits, and using a blank identifier in `for range` loops.

Functions benefit from multiple return values—especially for error handling—and `defer` for cleanup. Named result parameters can serve as documentation and simplify returns. The data section clarifies `new` (zeroed memory, returning a pointer) and `make` (only for slices, maps, channels; returns an initialized value). The zero-value‑is‑useful property is encouraged; composite literals and constructors are preferred when zero values aren’t enough. Slices are the primary sequence type, maps are built‑in, and the `append` function handles growth. Formatted printing is done with `fmt` and custom `String` methods.

In methods, pointer vs. value receivers have clear rules. Interfaces are satisfied implicitly; the guide illustrates how to use type assertions, type switches, and interface conversions, and recommends returning interface types from constructors to hide implementation details. Embedding promotes methods of inner types, avoiding manual forwarding.

The concurrency section introduces goroutines and channels with the motto “share memory by communicating.” Unbuffered channels synchronise, buffered channels can act as semaphores, and channels of channels enable request‑response patterns. Examples show parallel computation and a leaky‑buffer free list.

Error handling uses the built‑in `error` interface, with `panic` only for truly unrecoverable situations; `recover` inside a deferred function can catch panics and restore control. A full working web server that generates QR codes closes the guide, demonstrating templates and the `net/http` package.

## Key Claims
- Go programs are most successful when written in Go’s own idiom, not as a translation from other languages.
- `gofmt` eliminates formatting debates; all standard library code is formatted with it.
- Names have semantic effect: initial letter case controls export. Short, concise package names, no `Get` prefix for getters, and `MixedCaps` naming are conventions.
- Control structures avoid parentheses, require braces, and `switch` does not fall through. Type switch is a powerful idiom for dispatching on dynamic types.
- Multiple return values, especially for errors, keep error handling linear and clear. `defer` guarantees cleanup of resources like files and mutexes.
- `new` allocates zeroed memory and returns a pointer; `make` initialises slices, maps, and channels. Types whose zero value is usable simplify code.
- Slices are the default sequence type. `append` adds elements efficiently. Maps are reference types; test presence with the “comma ok” idiom.
- Implicit interface satisfaction decouples types. Embedding promotes methods without forwarder boilerplate.
- Concurrent code should prefer communicating over channels to sharing memory, enabling clear synchronisation.
- Errors are values; `panic`/`recover` is for package‑internal error handling shortcuts, not for exposing panics to clients.

## Suggested Links
- [Effective Go (original)](https://go.dev/doc/effective_go)
- Language Specification (the Go Spec)
- Tour of Go
- How to Write Go Code
- Go Release Notes
