---
title: Commonmark Parsing Phases
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Parsing Phases

The CommonMark specification defines a rigorous two-phase parsing strategy to resolve ambiguity and ensure consistent rendering across platforms. This approach prioritizes structural integrity before processing internal content details.

## Block Structure Phase

The first phase constructs the **block-level tree**. During this stage, the parser identifies container blocks (such as lists and block quotes) and leaf blocks (such as paragraphs and code blocks). Key activities include:

-   Handling structural boundaries and indentation rules (relative vs. four-space).
-   Processing "lazy continuation" for markers like `>`.
-   Determining list tightness and looseness based on blank lines.
-   Resolving interruptions between block elements.

## Inline Structure Phase

Once the block structure is established, the second phase processes content within identified blocks to resolve **inline elements**. This includes:

-   Resolving emphasis (`*`, `_`) and strong emphasis (`**`, `__`).
-   Identifying code spans, links, and raw HTML tags.
-   Processing entities and backslash escaping limits.

This phase utilizes a pre-computed link map and delimiter stacks to efficiently handle complex inline syntax.

## Key Principles

-   **Block Precedence**: Indicators of block structure always take precedence over inline indicators. For example, text that looks like emphasis is not parsed as such if it appears inside a code block or similar structural context.
-   **Escaping Limits**: Backslashes escape specific characters in certain contexts but do not work in code blocks, code spans, autolinks, or raw HTML. Escapes also fail inside HTML attributes.
-   **Indentation Rules**: Tabs are generally passed through literally but behave as four spaces in contexts defining block structure. Internal tabs within content are preserved.
-   **Soft Line Breaks**: Line endings not preceded by two spaces or a backslash are treated as soft breaks, giving renderers flexibility in output while typically matching browser behavior for hard breaks.
