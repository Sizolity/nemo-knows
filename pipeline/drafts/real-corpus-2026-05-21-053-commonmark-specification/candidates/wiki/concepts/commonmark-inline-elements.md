---
title: Commonmark Inline Elements
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Inline Elements

The **CommonMark Specification** defines a rigorous standard for parsing and rendering text content within block structures. Unlike the original Markdown description, which lacked specific rules for many edge cases, CommonMark provides unambiguous definitions to ensure consistent behavior across different platforms.

## Structure

The specification employs a two-phase parsing strategy:
1.  **Block Structure Phase**: Identifies container blocks (such as lists and block quotes) and leaf blocks (like paragraphs and code blocks). It handles structural boundaries, indentation rules, and "lazy continuation" for markers like `>`.
2.  **Inline Structure Phase**: Processes content within identified blocks to resolve emphasis, strong emphasis, code spans, links, and raw HTML tags.

## Inline Elements

The inline structure phase resolves the following elements:
-   **Emphasis**: Uses asterisks (`*`) or underscores (`_`).
-   **Strong Emphasis**: Uses double asterisks (`**`) or double underscores (`__`).
-   **Code Spans**: Encloses code fragments in backticks.
-   **Links**: Supports full, collapsed, shortcut, and autolink syntaxes.
-   **Images**: Defines image syntax distinct from links.
-   **Entities**: Handles character entities.
-   **Escaping**: Uses backslashes to escape specific characters like `*` or `#`.

## Parsing Rules

CommonMark adheres to several key rules regarding inline content:
-   **Block Precedence**: Indicators of block structure always take precedence over inline indicators. For example, text inside a code block is treated literally, ignoring emphasis markers.
-   **Escaping Limits**: Backslashes escape specific characters in certain contexts but do not work in code blocks, code spans, autolinks, or raw HTML. Escapes also fail inside HTML attributes.
-   **Soft Line Breaks**: Line endings not preceded by two spaces or a backslash are treated as soft breaks. Renderers have flexibility in outputting these visually, though browsers typically treat them identically to hard breaks.

## Metadata & Introduction

The specification includes rationale for the design choices, conformance testing tools (`spec_tests.py`), and character definitions. It aims to resolve ambiguity regarding sublist indentation, blank lines before block quotes or headings, code block requirements, and list item wrapping to prevent implementation divergence.
