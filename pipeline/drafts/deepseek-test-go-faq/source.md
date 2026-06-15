---
title: Go FAQ
kind: source
sources:
  - raw/web/corpus-2026-05-18/034-go-faq.md
confidence: medium
---

## What It Is

The official Frequently Asked Questions (FAQ) document for the Go programming language, maintained at `https://go.dev/doc/faq`. It covers the language's origins, design rationale, usage patterns, type system, concurrency model, and common practical questions from the Go community.

## Summary

Go was created at Google by Robert Griesemer, Rob Pike, and Ken Thompson, starting with whiteboard sketches in September 2007. Frustrated by the complexity of C++ and Java, and the efficiency/safety tradeoffs of dynamic languages like Python, they aimed for a language combining efficient compilation, efficient execution, and ease of programming—with first-class concurrency, garbage collection, and fast build times. The language became a public open-source project in November 2009.

The FAQ explains Go's design philosophy: simplicity, orthogonality of concepts, and reducing both kinds of typing (keystrokes and type hierarchies). It addresses why Go omits features such as exceptions, assertions, type inheritance, operator overloading, implicit numeric conversions, and pointer arithmetic—each omission driven by a desire for clarity, safety, and maintainability.

Concurrency is built on CSP-inspired goroutines and channels rather than threads. Interfaces are satisfied implicitly, enabling lightweight structural typing without `implements` declarations. The FAQ also covers practical topics: maps and slices as reference types, zero-size type behavior, nil interface gotchas, value vs. pointer receivers, generics (added in Go 1.18 with square-bracket syntax), and the module system for dependency management.

## Key Claims

- Go aims to combine the ease of dynamically typed languages with the efficiency and safety of statically typed, compiled languages.
- Implicit interface satisfaction eliminates type hierarchies and reduces boilerplate while enabling flexible, composable designs.
- Goroutines are lightweight, multiplexed onto OS threads, with resizable stacks, making it practical to run hundreds of thousands concurrently.
- Concurrency follows the CSP model: "Do not communicate by sharing memory. Instead, share memory by communicating."
- Go rejects exceptions and assertions in favor of explicit, multi-value error returns to keep error handling visible and straightforward.
- Map operations are not atomic by default; safe concurrent access requires external synchronization or `sync.Map`.
- Go 1.18 introduced generics using square brackets, with mandatory constraints expressed as interface types.
- The standard Go compiler (`gc`) is self-hosting (written in Go since 1.5) and produces statically linked binaries that include the full runtime.
- Garbage collection is a mark-and-sweep collector with sub-millisecond pause times, running concurrently on multiprocessor systems.
- The Go 1 compatibility promise prevents breaking changes to the language and standard library at the source level.

## Suggested Links

- Official Go FAQ: [https://go.dev/doc/faq](https://go.dev/doc/faq)
- Go at Google: Language Design in the Service of Software Engineering (article referenced in the FAQ; no explicit URL provided in the raw text beyond mention)
- Effective Go (guidance mentioned; no explicit URL in raw text)
- Go Code Review Comments (collection of idiom notes; no explicit URL)
- Defer, Panic, and Recover (blog post; no explicit URL)
- Errors are values (blog post; no explicit URL)
- Share Memory By Communicating (code walk; no explicit URL)
- Concurrency is not Parallelism (talk; no explicit URL)
- Why Generics? (blog post; no explicit URL)
- Protocol Buffers Go support: github.com/golang/protobuf/ (explicitly mentioned)
- `pkg.go.dev/pkg/` – global package discovery (explicitly mentioned)
- `goimports` – tool for automatic import management (mentioned, no URL)
- `gopls` – Go language server for LSP (mentioned, no URL)
