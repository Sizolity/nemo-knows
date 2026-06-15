---
kind: topic
sources: [raw/web/corpus-2026-05-18/053-commonmark-specification.md]
status: draft
---

# Ingest Plan

## Source Summary
- The document is the **CommonMark Specification** (v0.31.2) by John MacFarlane, designed to resolve ambiguity in Markdown syntax through strict conformance tests and side-by-side Markdown/HTML examples.
- Content spans from metadata and introduction through block-level elements (headings, code blocks, lists), container blocks (quotes), and inline structure (links, emphasis, entities).
- The text emphasizes a two-phase parsing strategy: **Block Structure** followed by **Inline Structure**, with specific rules for indentation, lazy continuation, and delimiter stacks.

## Candidate Wiki Pages
- wiki/sources/commonmark-specification.md — Primary source file containing the full CommonMark specification text and conformance test definitions.
- wiki/concepts/commonmark-parsing-phases.md — Explains the distinction between Block Structure (identifying containers/leaves) and Inline Structure (processing links/emphasis).
- wiki/topics/commonmark-lists.md — Covers list syntax, tight vs. loose lists, indentation math, and interruption of paragraphs.
- wiki/concepts/commonmark-inline-elements.md — Details parsing precedence for code spans, autolinks, HTML tags, links, and emphasis delimiters.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that `wiki/sources/commonmark-specification.md` accurately reflects the version 0.31.2 metadata found in Chunk 01.
- [ ] Ensure `wiki/concepts/commonmark-parsing-phases.md` captures the "two-phase" model mentioned repeatedly in Group Notes.
- [ ] Confirm `wiki/topics/commonmark-lists.md` includes definitions for both tight and loose lists as per Section 3.c.
