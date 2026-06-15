---
kind: topic
sources: [raw/web/corpus-2026-05-18/053-commonmark-specification.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source is the CommonMark Specification (version 0.31.2) by John MacFarlane, covering the complete parsing grammar from document metadata to inline element resolution.
- The document defines a two-phase parsing strategy: first constructing block structure (paragraphs, lists, code blocks) and then parsing inline text (emphasis, links, code spans).
- Key structural rules include the "four-space rule" for indentation, strict handling of list tightness/looseness, and specific syntax for ATX/Setext headings, fenced code blocks, and HTML blocks.
- Inline parsing details cover delimiter stack algorithms for nested emphasis, link reference normalization, autolink handling, and hard line break conversion.

## Candidate Wiki Pages
- wiki/sources/commonmark-specification.md — Comprehensive reference for the CommonMark spec, including block structure, list rules, and inline parsing logic.
- wiki/concepts/commonmark-block-structure.md — Technical guide covering thematic breaks, ATX/Setext headings, code blocks, HTML blocks, and link reference definitions.
- wiki/concepts/commonmark-inline-elements.md — Deep dive into emphasis delimiters, code spans, links, images, autolinks, and the delimiter stack algorithm.
- wiki/concepts/commonmark-lists.md — Rules for list markers, indentation hierarchy, loose vs. tight lists, and sublist handling.
- wiki/topics/commonmark-escaping-and-entities.md — Guide on backslash escapes, HTML entity references, and character replacement rules.
- wiki/topics/commonmark-line-breaks.md — Explanation of soft vs. hard line breaks and rendering variations.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that all candidate pages adhere to the `wiki/sources/`, `wiki/concepts/`, or `wiki/topics/` directory constraints.
- [ ] Ensure no nested directories are created within the candidate pages.
- [ ] Confirm that the "Candidate Wiki Hints" from group notes are consolidated into the proposed pages.
- [ ] Check that the source summary accurately reflects the scope of the CommonMark specification chunks.
