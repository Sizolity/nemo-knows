---
kind: topic
sources: [raw/web/corpus-2026-05-18/053-commonmark-specification.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source is the CommonMark Specification (Version 0.31.2), a formal grammar designed to resolve implementation divergences found in the original Markdown description.
- The document defines a two-phase parsing strategy: first constructing a tree of block-level elements (headings, lists, code blocks), then parsing raw text within those blocks into inline elements (links, emphasis, code spans).
- Key technical rules cover character handling (tabs, backslash escapes), block structure (ATX/Setext headings, fenced/indented code blocks), container blocks (block quotes, lists with "laziness"), and inline syntax (delimiter runs, link reference definitions, raw HTML).

## Candidate Wiki Pages
- wiki/sources/commonmark-specification.md — Comprehensive reference for the CommonMark grammar, parsing rules, and syntax definitions.
- wiki/concepts/commonmark-parsing-phases.md — Technical explanation of the two-phase parsing model (block structure vs. inline content) and the delimiter stack algorithm.
- wiki/concepts/commonmark-indentation-rules.md — Rules for indentation, the four-space code block rule, list item indentation, and handling of tabs.
- wiki/concepts/commonmark-inline-elements.md — Syntax and precedence for emphasis, strong emphasis, links, images, and code spans.
- wiki/topics/commonmark-heading-syntax.md — Overview of ATX and Setext heading definitions, including multiline support and indentation limits.
- wiki/topics/commonmark-code-blocks.md — Comparison and syntax for indented and fenced code blocks, including info strings and closing fence matching.
- wiki/topics/commonmark-list-structure.md — Classification of loose vs. tight lists, sublist indentation, and marker rules.
- wiki/topics/commonmark-html-blocks.md — Rules for parsing raw HTML blocks, start/end conditions, and interaction with Markdown syntax.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that all candidate pages map to distinct sections of the specification without overlap.
- [ ] Ensure the "Parsing Phases" concept page accurately reflects the delimiter stack algorithm details from chunks 08–13.
- [ ] Confirm that the "Indentation Rules" page distinguishes between margin-relative and marker-relative indentation.
- [ ] Check that the "HTML Blocks" topic covers the seven types and the restriction on blank lines.
- [ ] Validate that the "Inline Elements" page addresses the minimization of nesting depth for emphasis and links.
