---
kind: topic
sources: [raw/web/corpus-2026-05-18/031-effective-go.md]
status: draft
---

# Ingest Plan

## Source Summary
- Official Go guide (originally 2009) covering idiomatic style, formatting, naming, control structures, functions, data allocation, methods, interfaces, concurrency, and error handling.
- Not a specification; a tutorial-style resource for writing clear, performant, and idiomatic Go code.
- Remains a foundational guide despite not covering generics or modules added since.

## Candidate Wiki Pages
- wiki/sources/effective-go.md — Source page with metadata and a summary of the guide.
- wiki/topics/go-style-formatting.md — Covers formatting (gofmt), comment conventions, and naming rules (packages, getters, interfaces, MixedCaps).
- wiki/topics/go-control-structures.md — Covers if, for (including range), switch, type switch, redeclaration/reassignment, and semicolon insertion rules.
- wiki/topics/go-data-allocation.md — Covers new vs make, arrays, slices, maps, composite literals, append, and zero-value usefulness.
- wiki/topics/go-concurrency-patterns.md — Covers goroutines, channels (buffered/unbuffered), select, semaphore pattern, leaky buffer, parallelization, and “share memory by communicating”.

## Suggested Links
- Links explicitly present in the source: Go Language Specification, Tour of Go, How to Write Go Code, `fmt` package godoc, `html/template` package documentation. (These could be captured as related references.)

## Review Checklist
- [ ] Confirm all candidate page slugs are direct children of `wiki/sources/` or `wiki/topics/` with no nested directories.
- [ ] Verify the source page includes a concise summary and correct metadata.
- [ ] Ensure the selected pages align with the source’s tutorial/guide nature (topic pages) and avoid speculative concept pages unless justified by API/technical spec content.
- [ ] Check for overlap with existing wiki pages on Go idioms or style.
