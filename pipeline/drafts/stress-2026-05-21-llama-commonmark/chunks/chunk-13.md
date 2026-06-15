---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
### Chunk Context
- **Heading path**: 3. c > foo\
- **Line range**: 7802-8133
- **Coverage**: Sections 6.8 (Soft line breaks), 6.9 (Textual content), and Appendix A (Parsing strategy: block and inline structure).

### Local Summary
This chunk details how CommonMark handles soft line breaks, preserves raw textual content (including special characters and spaces), and defines a two-phase parsing strategy. Phase 1 constructs the block structure (tree of blocks) by consuming lines and managing open/closed states. Phase 2 parses the raw text within those blocks into inline elements (strings, code spans, emphasis). The text includes a detailed algorithm for handling nested emphasis and links using a "delimiter stack."

### Key Claims
- A regular line ending not preceded by two or more spaces or a backslash is parsed as a **softbreak**.
- A conforming parser may render a soft line break in HTML either as a line ending or as a space.
- Any characters not given an interpretation by previous rules are parsed as **plain textual content**.
- Internal spaces are preserved verbatim.
- Parsing occurs in two phases: first constructing the block structure, then parsing raw text contents into inline elements.
- The delimiter stack is a doubly linked list used to track opening and closing delimiters for emphasis and links.

### Entities And Concepts
- **Soft line break**: A line ending parsed as a break or space.
- **Textual content**: Plain text characters not interpreted as Markdown syntax.
- **Block structure**: The tree of blocks (document, block quotes, lists, paragraphs) constructed in Phase 1.
- **Inline structure**: The sequence of inline elements (strings, code spans, links, emphasis) constructed in Phase 2.
- **Delimiter stack**: A data structure tracking delimiters (`*`, `_`, `[`, `!`, `]`) to resolve emphasis and links.
- **Lazy continuation**: A line that continues an open block without closing it.
- **Setext headings**: Headings formed by an underline line.
- **Reference link definitions**: Detected when a paragraph is closed.

### Procedures And API Details
- **Phase 1 Procedure**:
  1. Iterate through open blocks to check conditions for remaining open.
  2. Consume continuation markers and look for new block starts (e.g., `>`). Close unmatched blocks before creating new ones.
  3. Incorporate the remainder of the line into the last open block.
- **Phase 2 Procedure**:
  - Walk the tree and parse raw string contents of paragraphs and headings as inlines.
  - Resolve reference links using the map constructed in Phase 1.
- **Inline Parsing Algorithm**:
  - Insert text nodes for runs of `*`, `_`, `[`, or `![` and add pointers to the delimiter stack.
  - On hitting `]`, call `look for link or image`.
  - On end of input, call `process emphasis`.
- **look for link or image**:
  - Search backwards from the top of the delimiter stack for an active opening `[` or `![`.
  - If found and active, parse ahead to determine link/image type (inline, reference, collapsed, shortcut).
  - If a link is found, set all `[` delimiters before the opening delimiter to inactive.
- **process emphasis**:
  - Use `stack_bottom` to limit descent.
  - Track `openers_bottom` for each delimiter type.
  - Repeat until potential closers run out:
    - Move `current_position` to the first potential closer (`*` or `_`).
    - Look back for a matching potential opener.
    - If found, determine emphasis/strong emphasis, insert nodes, and remove delimiters.
    - If not found, update `openers_bottom` and remove inactive closers.

### Nuance Or Contradictions
- **Rendering Flexibility**: While the spec defines a soft line break, the rendering in HTML is flexible; it can be a line ending or a space without changing the result in browsers.
- **Parser Options**: Renderers may provide an option to render soft line breaks as hard line breaks.
- **Delimiter Stack Logic**: The algorithm distinguishes between "active" delimiters and those that are potential openers/closers based on preceding and following characters. Inactive delimiters are removed from the stack to prevent invalid nesting (e.g., links within links).

### Candidate Wiki Hints
- **Soft Line Breaks**: A dedicated page explaining the rules for line breaks, their HTML rendering variations, and the concept of hard line breaks.
- **Parsing Strategy**: A technical page detailing the two-phase parsing approach (block vs. inline) and the tree-based document model.
- **Delimiter Stack Algorithm**: A deep-dive page on the specific algorithm for resolving nested emphasis and links, including the logic for the delimiter stack.
- **Textual Content**: A page covering the handling of special characters, internal spaces, and plain text preservation.
