---
kind: topic
sources: [raw/web/corpus-2026-05-18/031-effective-go.md]
status: draft
---

# Ingest Plan

## Source Summary
- Original Go style and idiom guide (written for Go’s 2009 release, not actively updated); covers formatting, naming, control structures, functions, data allocation, methods, interfaces, embedding, concurrency, and error handling.
- Includes practical patterns: `gofmt` usage, named returns, `defer`, embedded structs, channel-based concurrency, `panic`/`recover`, and a complete QR-code web server example.

## Candidate Wiki Pages
- `wiki/sources/effective-go.md` — capture of the full raw source document for reference.
- `wiki/topics/go-idioms.md` — core style and idiom advice: formatting, naming, control structures, functions, and data.
- `wiki/topics/go-concurrency.md` — goroutines, channels, and patterns like semaphores, parallelization, and the leaky buffer.
- `wiki/topics/go-error-handling.md` — error reporting, `panic`/`recover`, and the re-panic pattern for internal error handling.
- `wiki/topics/go-web-server-example.md` — the QR code web server built with `net/http` and `html/template`.

## Suggested Links
- https://go.dev/doc/effective_go
- https://go.dev/ref/spec
- https://go.dev/tour/
- https://go.dev/doc/code
- https://go.dev/doc/devel/release

## Review Checklist
- [ ] Confirm source text is completely preserved in `wiki/sources/effective-go.md`.
- [ ] Verify extracted topic pages align with major sections and do not overlap excessively.
- [ ] Ensure all candidate page slugs are unique, lowercase, hyphenated, and placed directly under the specified directories.
- [ ] Check that suggested links resolve correctly and are the canonical URLs.
