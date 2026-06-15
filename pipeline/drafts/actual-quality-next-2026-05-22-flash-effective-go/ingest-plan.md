---
kind: topic
sources: [raw/web/corpus-2026-05-18/031-effective-go.md]
status: draft
---

# Ingest Plan

## Source Summary
- The official “Effective Go” guide (2009, not actively updated for generics/modules) provides advice for writing idiomatic, clear, and performant Go code.
- It covers formatting with `gofmt`, idiomatic naming, control structures, functions, data handling, methods, interfaces, embedding, concurrency, error handling, and a complete web-server example.
- The source is a tutorial/guide; therefore candidate pages are primarily topic pages that extract reusable practices and idioms.

## Candidate Wiki Pages
- wiki/sources/effective-go.md — Raw import of the full guide for reference.
- wiki/topics/go-formatting-conventions.md — Idiosyncratic formatting rules: `gofmt` as sole layout tool, tab indentation, no line-length limits, minimal parentheses.
- wiki/topics/go-naming-idioms.md — Conventions for package names, getters (no `Get` prefix), one-method interfaces (`-er` suffix), and `MixedCaps` naming.
- wiki/topics/go-error-handling-patterns.md — Patterns for returning errors as values, `panic`/`recover` usage, custom error types like `PathError`, and best practices for error messages.
- wiki/topics/go-concurrency-patterns.md — “Share memory by communicating”: goroutines, channels, `select`, rate-limiting with channel semaphores, and parallelization patterns.

## Suggested Links
- Original source URL: [https://go.dev/doc/effective_go](https://go.dev/doc/effective_go)
- Referenced guides (to be linked or created later): Go language specification, Tour of Go, How to Write Go Code

## Review Checklist
- [ ] All major idiom categories from Effective Go (formatting, naming, control flow, errors, concurrency) are covered by candidate topic pages.
- [ ] Candidate pages are immediate children of `wiki/sources/` or `wiki/topics/` and do not include forbidden pages (index, log, AGENTS, schema).
- [ ] Suggested links include the original source and any essential upstream guides mentioned in the text.
- [ ] Source summary accurately reflects the 2009 vintage note and the core content of the guide.
