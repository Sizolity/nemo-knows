---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
**Heading path:** Document
**Line range:** 1–516
**Scope:** Introduction to the CommonMark Specification (Version 0.31.2), covering metadata, rationale for a formal spec, document structure, and preliminary character definitions.

## Local Summary
This chunk introduces the **CommonMark Specification**, version 0.31.2, authored by John MacFarlane under a Creative Commons BY-SA license. It addresses the lack of unambiguous syntax in John Gruber’s original Markdown description, which led to divergent implementations across platforms like GitHub and Reddit. The spec aims to resolve ambiguities regarding lists, code blocks, headings, and inline structure through side-by-side Markdown/HTML examples that serve as conformance tests.

## Key Claims
- **Readability Goal:** Markdown is designed so that a formatted document is readable as plain text without appearing marked up with tags or instructions.
- **Ambiguity in Original Syntax:** Gruber’s original description does not specify rules for sublist indentation, blank lines before block quotes/headings, code block requirements, list item wrapping, right-aligned markers, thematic breaks within lists, marker precedence, section headings inside lists, empty list items, link reference scope, or definition precedence.
- **Implementation Divergence:** Without a formal spec, implementations consulted buggy scripts (Markdown.pl) or made arbitrary choices, leading to documents rendering differently on different systems without triggering syntax errors.
- **HTML as Test Representation:** The spec uses HTML for side-by-side examples because it represents structural distinctions well; however, not every HTML feature in the examples is mandated by the spec (e.g., percent-encoding of non-ASCII URLs).

## Entities And Concepts
- **CommonMark**: A specification for Markdown syntax to ensure interoperability.
- **Markdown.pl**: The original Perl script by John Gruber used as a reference but deemed buggy and insufficient for defining unambiguous syntax.
- **Conformance Tests**: Examples in the spec that double as tests against implementations using `spec_tests.py`.
- **Abstract Syntax Tree**: The internal representation Markdown is parsed into; HTML samples approximate this structure.
- **Unicode Code Points**: Defined as characters for the purpose of the spec, including combining accents.
- **Line Endings**: Defined as line feed (U+000A), carriage return not followed by line feed, or both.

## Procedures And API Details
**Parsing Strategy Overview:**
The document outlines a two-phase parsing strategy:
1.  **Phase 1: Block Structure**: Identifying container blocks and leaf blocks.
2.  **Phase 2: Inline Structure**: Processing inlines within identified blocks.

**Test Execution Command:**
```bash
python test/spec_tests.py --spec spec.txt --program PROGRAM
```

**Tooling:**
- `tools/makespec.py`: Converts the source text file (`spec.txt`) into HTML or CommonMark.

## Nuance Or Contradictions
- **Tab Handling**: Tabs are not expanded to spaces generally but behave as if replaced by four spaces in contexts defining block structure (e.g., indented code blocks, list item continuation). Internal tabs within content are passed through literally.
- **HTML Mandates vs. Examples**: The spec provides HTML examples where the destination URL might contain non-ASCII characters not percent-encoded. While the spec defines what counts as a link destination, it does not mandate encoding; implementers using automatic tests must provide a renderer that conforms to these expectations (encoding non-ASCII), but conforming implementations may choose different renderers.
- **Original vs. Spec**: The spec explicitly contrasts its unambiguous rules with Gruber’s original description, noting that the latter allows interpretations (like sublist indentation) that contradict common assumptions or specific implementations like Markdown.pl.

## Candidate Wiki Hints
- **CommonMark Specification**: A canonical reference for Markdown parsers to ensure cross-platform consistency.
- **Markdown Conformance Testing**: Methodology for validating Markdown implementations against a standard spec using `spec_tests.py`.
- **Tab Behavior in Markdown**: Rules defining how tabs function in block structure contexts versus inline content.
