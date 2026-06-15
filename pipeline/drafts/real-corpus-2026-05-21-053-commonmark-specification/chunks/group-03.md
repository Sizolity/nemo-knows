---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

This document group constitutes a substantial portion of the **CommonMark Specification**, specifically focusing on the transition from initial metadata fetching and introductory examples into the core parsing algorithms. The content spans from the top-level document structure through specific edge cases involving list continuations, nested emphasis, soft line breaks, and textual content preservation. The sequence covers the implementation strategy for a two-phase parser (block structure followed by inline structure) and details the stack-based algorithm for resolving nested delimiters.

# Cross-Chunk Summary

The notes progress from high-level document identification to specific technical deep-dives:
1.  **Initial Structure (Chunks 01–04):** Establishes the document root, fetch metadata, retrieved text handling, and list item continuations. It covers basic structural elements like nested lists and block transitions.
2.  **Edge Cases & Examples (Chunks 05–07):** Explores specific formatting scenarios including bracketed items (`[Foo]`), deeply nested "foo" headers illustrating structural complexity, and textual content handling.
3.  **Parsing Logic (Chunks 08–13):** Shifts focus to the mechanical implementation of the parser. This includes the definition of soft line breaks, the two-phase parsing model (building block trees via lazy continuations), and the detailed stack-based algorithm for processing inline elements like links and emphasis.

# Repeated Or Central Claims

-   **Two-Phase Parsing Model:** The document consistently asserts that CommonMark parsing occurs in two distinct phases: Phase 1 constructs a block-level tree structure (handling indentation, markers, and lazy continuations) without fully parsing inline content; Phase 2 then parses the raw text of paragraphs and headings into inline elements using a pre-computed link map.
-   **Soft Line Breaks:** It is repeatedly noted that line endings not preceded by two or more spaces or a backslash are treated as soft breaks. While renderers have flexibility in outputting these as spaces or newlines, browsers typically treat them identically to hard breaks visually.
-   **Lazy Continuation:** The concept of "lazy continuation" is central to the block structure phase, where lines that do not start new blocks (e.g., block quotes marked by `>`) are added to the current open block even if the block's structural condition changes mid-line, ensuring complex nesting logic is handled correctly.
-   **Delimiter Stack:** The use of a doubly linked list (delimiter stack) to track potential openers (`[`, `*`, `_`) and closers for resolving nested emphasis and links is a persistent technical detail across the inline parsing sections.

# Important Local Details

-   **Block Closure Timing:** Blocks are not strictly closed immediately upon encountering a new block start on the same logical line. Instead, unmatched blocks from the previous step are closed *before* creating the new block, allowing for complex nesting logic within a single line.
-   **Delimiter Stack Algorithm Mechanics:**
    -   **Insertion:** When hitting `*`, `_`, `[`, or `![`, a text node with literal content is inserted, and a pointer is added to the delimiter stack.
    -   **Stack Element Data:** Contains pointers to text nodes, delimiter type, count of delimiters, active status, and potential role (opener/closer).
    -   **Link/Image Resolution:** The algorithm starts at the top of the stack, looking back for opening brackets. If inactive, it returns a literal `]`. If active, it parses ahead for link/image types.
    -   **Emphasis Processing:** Iterates until potential closers are exhausted, moves forward to find the first potential closer (`*` or `_`), looks back above `stack_bottom` for matching openers, determines strength (length >= 2), inserts emph/strong nodes, and updates text nodes.
-   **Textual Content Preservation:** Characters not matched by specific rules (links, emphasis, code spans) are parsed as plain text, preserving internal spaces verbatim.

# Candidate Wiki Hints

-   **Page: CommonMark Parsing Strategy** (Concept: Two-phase parsing model of block vs. inline processing).
-   **Page: Delimiter Stack Algorithm** (Concept: Handling nested emphasis and links using a stack-based approach).
-   **Page: Soft Line Breaks** (Concept: Rules for soft breaks and rendering variance between browsers and renderers).
-   **Page: Block Tree Construction** (Concept: Building hierarchical document representations via lazy continuations).

# Gaps Or Cautions

-   **Rendering Variance:** The specification explicitly notes that while it mandates how soft breaks are parsed, the visual output in browsers is identical for soft and hard breaks, implying significant renderer flexibility regarding the final CSS/HTML generation of these line endings.
-   **Incomplete Header Paths:** Several chunks (specifically Chunks 02–13) show heading paths that are either "not a heading" or generic placeholders like "foo", "[Foo]", or "aaa". This suggests these sections may be testing edge cases, examples, or specific internal states rather than defining the primary structural headers of the document itself.
-   **Line Continuation Ambiguity:** The handling of lines where a block quote marker (`>`) appears but the content continues with text that might suggest a new block requires careful attention to the "lazy continuation" rules to avoid misinterpreting block boundaries.
