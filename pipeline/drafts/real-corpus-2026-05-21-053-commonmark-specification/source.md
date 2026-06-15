---
title: CommonMark Specification
kind: source
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## What It Is

The **CommonMark Specification** (Version 0.31.2) is a formal standard authored by John MacFarlane designed to resolve the ambiguity of original Markdown syntax. The specification aims to ensure consistent rendering across different platforms by providing unambiguous rules for parsing block-level and inline elements, utilizing side-by-side Markdown/HTML examples as conformance tests.

## Summary

The document defines a rigorous two-phase parsing strategy:
1.  **Block Structure Phase**: Identifies container blocks (lists, block quotes) and leaf blocks (paragraphs, code blocks, thematic breaks). It handles structural boundaries, indentation rules (relative vs. four-space), and "lazy continuation" for markers like `>`.
2.  **Inline Structure Phase**: Processes content within identified blocks to resolve emphasis (`*`, `_`), strong emphasis (`**`, `__`), code spans, links, and raw HTML tags.

The specification systematically covers:
-   **Metadata & Introduction**: Rationale, conformance testing tools (`spec_tests.py`), and character definitions.
-   **Block-Level Elements**: Thematic breaks, ATX/Setext headings, fenced/indented code blocks, and seven types of HTML blocks.
-   **Container Blocks**: Block quotes and lists (bullets/ordered), including rules for tightness, laziness, and interruption of paragraphs.
-   **Inline Elements**: Links (full, collapsed, shortcut, autolink), image syntax, entities, and backslash escaping limits.

## Key Claims

-   **Readability Goal**: Markdown is designed to be readable as plain text without appearing marked up with tags or instructions.
-   **Resolution of Ambiguity**: The original description lacked rules for sublist indentation, blank lines before block quotes/headings, code block requirements, and list item wrapping. This spec resolves these to prevent implementation divergence.
-   **Block Structure Precedence**: Indicators of block structure (e.g., `>` for block quotes) always take precedence over inline indicators (e.g., text looking like emphasis).
-   **Escaping Limits**: Backslashes escape specific characters (like `*`, `#`) in certain contexts but do not work in code blocks, code spans, autolinks, or raw HTML. Escapes also fail inside HTML attributes.
-   **Indentation Rules**: Tabs are generally passed through literally but behave as four spaces in contexts defining block structure (indented code blocks, list continuation). Internal tabs within content are preserved.
-   **Blank Line Semantics**: Multiple blank lines between paragraphs do not affect paragraph boundaries; they define list tightness/looseness or separate distinct block quotes. Blank lines inside HTML blocks are disallowed to avoid expensive balanced tag parsing.
-   **Parsing Phases**: The parser constructs a block-level tree via lazy continuations first, then parses inline content using a pre-computed link map and delimiter stacks.
-   **Soft Line Breaks**: Line endings not preceded by two spaces or a backslash are treated as soft breaks. Renderers have flexibility in outputting these visually, though browsers typically treat them identically to hard breaks.

## Suggested Links

none
