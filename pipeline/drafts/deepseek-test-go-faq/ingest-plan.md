---
kind: topic
sources: [raw/web/corpus-2026-05-18/034-go-faq.md]
status: draft
---

# Ingest Plan

## Source Summary
- Official Go FAQ covering language origins, design principles, type system, concurrency, memory management, tooling, and idioms.
- Structured as question–answer pairs; explains the rationale behind language choices.
- Published on go.dev; maintained by the Go team; content is current and authoritative.
- Serves as a primary reference for understanding Go’s philosophy and mechanics.

## Candidate Wiki Pages
- wiki/sources/go-faq.md — raw FAQ page as a source reference
- wiki/concepts/go-design-philosophy.md — origins, goals, principles, and design trade-offs (coverage: “Origins”, “Design”, “Changes from C”)
- wiki/concepts/go-type-system.md — interfaces, methods, type embedding, generics, underlying types, zero values (coverage: “Types”, “Type Parameters”)
- wiki/concepts/go-concurrency.md — goroutines, channels, CSP model, synchronization primitives (coverage: “Concurrency”, “Why goroutines instead of threads?”)
- wiki/concepts/go-error-handling.md — error values, panics, defer, error patterns (coverage: “Why does Go not have exceptions?”, “Defer, Panic, and Recover”)
- wiki/concepts/go-memory-management.md — garbage collection, stack vs heap, pointers, allocation (coverage: “Pointers and Allocation”, “Performance”, “Garbage collection”)
- wiki/concepts/go-tooling-and-modules.md — go command, modules, testing, documentation, IDEs (coverage: “Writing Code”, “Packages and Testing”, “How do I manage package versions?”)

## Suggested Links
- go.dev/doc/faq (original URL)
- “Go at Google: Language Design in the Service of Software Engineering” article
- “Effective Go” document
- “Go Code Review Comments” document
- “Constants” blog post
- “Defer, Panic, and Recover” article
- “Errors are values” blog post
- “Share Memory By Communicating” code walk and article
- Tutorial: Create a module (Go documentation)
- “Developing modules” guide (Go documentation)
- Go 1 compatibility guidelines

## Review Checklist
- [ ] Verify that all extracted text is correctly attributed and matches the source.
- [ ] Ensure concept pages are scoped to the immediate child slugs and accurately represent the FAQ content.
- [ ] Confirm suggested links point to existing external references; decide whether to capture them as separate source pages.
- [ ] Check that no wiki/index.md, wiki/log.md, AGENTS.md, or schema files are accidentally proposed.
- [ ] Validate that the source kind (FAQ) is appropriate for generating concept pages.
