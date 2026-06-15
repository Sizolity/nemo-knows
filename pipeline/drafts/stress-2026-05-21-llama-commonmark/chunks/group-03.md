---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

This group of notes covers the **CommonMark Specification**, specifically focusing on the transition from high-level document structure to low-level parsing algorithms. The content spans from the initial document metadata and list structures (Chunks 01-02) through various edge cases involving nested headings, code spans, and raw text (Chunks 03-07), culminating in the detailed parsing strategies for block and inline structures, soft line breaks, and the delimiter stack algorithm (Chunks 08-13).

The source material defines a two-phase parsing strategy:
1.  **Phase 1**: Constructing the block structure (tree of blocks) by consuming lines and managing open/closed states.
2.  **Phase 2**: Parsing the raw text within those blocks into inline elements (strings, code spans, emphasis, links).

# Cross-Chunk Summary

The document outlines the evolution of a Markdown parser's logic. Early chunks establish the document root, metadata fetching, and basic list handling. Subsequent chunks introduce complexity through nested headings, code spans (`foo *bar* \*baz\*`), and raw text preservation. The later chunks (08-13) provide the rigorous algorithmic definitions for how these structures are interpreted, specifically detailing:
-   The handling of soft line breaks versus hard breaks.
-   The definition of textual content as any character not interpreted by syntax rules.
-   The specific mechanics of the **delimiter stack**, a doubly linked list used to track opening and closing delimiters (`*`, `_`, `[`, `!`, `]`) to resolve nested emphasis and links.
-   The distinction between active delimiters and potential openers/closers to prevent invalid nesting (e.g., links within links).

# Repeated Or Central Claims

-   **Two-Phase Parsing**: The specification consistently emphasizes that parsing occurs in two distinct phases: first constructing the block structure, then parsing raw text contents into inline elements.
-   **Textual Content**: Any characters not given an interpretation by previous rules are parsed as plain textual content, with internal spaces preserved verbatim.
-   **Delimiter Stack**: The resolution of emphasis and links relies on a specific algorithm using a delimiter stack to track state, distinguishing between active and inactive delimiters.
-   **Soft Line Breaks**: A regular line ending not preceded by two or more spaces or a backslash is parsed as a soft break, which may render as a line ending or a space in HTML.
-   **Block Structure**: The document is modeled as a tree of blocks (document, block quotes, lists, paragraphs).

# Important Local Details

-   **Chunk 01**: Covers the document root, fetch metadata, retrieved text, and basic list item continuations (e.g., "List item two continued with an open block").
-   **Chunk 02**: Introduces complex nested heading paths (e.g., `foo > foo > foo > foo > foo > foo`) and raw text variations including code spans and HTML-like tags (`bar</p>`).
-   **Chunk 03-04**: Focuses on the "baz" heading path and the transition to "Heading" and "baz" structures.
-   **Chunk 05-07**: Covers specific bracketed headings (`[Foo]`) and the "aaa" heading path, representing sections of the spec dealing with specific syntax elements.
-   **Chunk 08-12**: Deep dive into list items (`2. b`, `3. c`) and their continuations across multiple line ranges (4756-7801).
-   **Chunk 13**: Details the "3. c > foo\" path, covering soft line breaks, textual content, and the full algorithm for the delimiter stack and lazy continuation.

# Candidate Wiki Hints

-   **Soft Line Breaks**: A dedicated page explaining the rules for line breaks, their HTML rendering variations, and the concept of hard line breaks.
-   **Parsing Strategy**: A technical page detailing the two-phase parsing approach (block vs. inline) and the tree-based document model.
-   **Delimiter Stack Algorithm**: A deep-dive page on the specific algorithm for resolving nested emphasis and links, including the logic for the delimiter stack.
-   **Textual Content**: A page covering the handling of special characters, internal spaces, and plain text preservation.
-   **Block Structure**: A page defining the tree of blocks (document, block quotes, lists, paragraphs) and how open/closed states are managed.

# Gaps Or Cautions

-   **Rendering Flexibility**: While the spec defines a soft line break, the rendering in HTML is flexible; it can be a line ending or a space without changing the result in browsers. Renderers may provide an option to render soft line breaks as hard line breaks.
-   **Parser Options**: The specification notes that renderers may offer options regarding how soft line breaks are handled, implying implementation variance.
-   **Delimiter Logic Nuance**: The algorithm distinguishes between "active" delimiters and those that are potential openers/closers based on preceding and following characters. Inactive delimiters are removed from the stack to prevent invalid nesting.
-   **Line Range Continuity**: The chunk notes indicate specific line ranges (e.g., 7802-8133 for Chunk 13), suggesting that the full document is segmented. Care must be taken to ensure synthesis does not assume content exists outside the provided line ranges or chunk notes.
