---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
### Chunk Context
This chunk covers **Section 3.c** (specifically `foo\` backslash escapes) and transitions into **Section 6.8** (Soft line breaks), **6.9** (Textual content), and the **Appendix: A parsing strategy**. It details the two-phase parsing model of CommonMark, the construction of block trees via lazy continuations, and the specific algorithm for handling nested emphasis and links using a delimiter stack.

### Local Summary
The document defines how trailing backslashes create soft line breaks or escape characters, followed by rules for plain text preservation. The core focus shifts to the implementation strategy: a two-phase process where Phase 1 builds a block-level tree structure (paragraphs, lists, quotes) handling indentation and markers, while Phase 2 parses inline content (links, emphasis). A specific stack-based algorithm is detailed for resolving nested delimiters like `*` and `[`.

### Key Claims
- **Soft Line Breaks**: A regular line ending not preceded by two or more spaces or a backslash is parsed as a soft break. Renderers may output this as a space or a line ending; browsers treat them identically.
- **Textual Content**: Characters not matched by specific rules (links, emphasis, code spans) are parsed as plain text, preserving internal spaces verbatim.
- **Two-Phase Parsing**:
  - *Phase 1*: Consumes lines to build block structure (document -> blocks -> children). Text is assigned but not fully parsed. Link definitions are mapped here.
  - *Phase 2*: Parses raw text of paragraphs/headings into inline elements using the link map from Phase 1.
- **Lazy Continuation**: Lines that do not start new blocks (like `>`) are added to the current open block if they satisfy its condition, even if the block is technically "closed" by a subsequent structural change on the same logical line.

### Entities And Concepts
- **Soft Break**: A line ending rendered as a space or newline.
- **Textual Content**: Uninterpreted characters preserved verbatim.
- **Block Tree**: The hierarchical representation of the document (document -> block_quote -> paragraph).
- **Lazy Continuation**: Text added to an open block that precedes a new block start on the same input line.
- **Delimiter Stack**: A doubly linked list used in inline parsing to track potential link/emphasis openers (`[`, `*`, `_`) and closers.
- **Inline Elements**: Strings, code spans, links, emphasis.

### Procedures And API Details
**Phase 1: Block Structure Procedure**
For each line processed:
1. Iterate through open blocks (root down to deepest) to check if the line satisfies their conditions (e.g., block quote needs `>`).
2. Consume continuation markers; look for new block starts. If found, close unmatched blocks and create the new block as a child of the last matched container.
3. Incorporate remaining text into the last open block.

**Phase 2: Inline Structure Procedure**
1. Close all open blocks.
2. Walk the tree visiting every node.
3. Parse raw string contents of paragraphs/headings using the link reference map.

**Delimiter Stack Algorithm (for nested emphasis/links)**
- **Insertion**: When hitting `*`, `_`, `[`, or `![`, insert a text node with literal content and add a pointer to the delimiter stack.
- **Stack Element Data**: Contains pointer to text node, delimiter type (`[`, `![, *, _`), count of delimiters, active status, and potential role (opener/closer).
- **Look for Link/Image**:
  - Start at top of stack, look back for opening `[` or `![`.
  - If inactive, remove and return literal `]`.
  - If active, parse ahead for link/image types. If found, create node, run process emphasis on children, and remove opening delimiter.
- **Process Emphasis**:
  - Iterate until potential closers are exhausted.
  - Move forward to find first potential closer (`*` or `_`).
  - Look back above `stack_bottom` for matching opener.
  - If found: Determine strength (length >= 2), insert emph/strong node, remove intermediate delimiters, update text nodes.
  - If not found: Update bounds and advance current position.

### Nuance Or Contradictions
- **Rendering Variance**: While the spec mandates a soft break is parsed as a line ending or space, it explicitly notes that in browsers these render the same, implying renderer flexibility for the visual output of soft breaks versus hard breaks.
- **Block Closure Timing**: Blocks are not strictly closed immediately upon encountering a new block start; instead, unmatched blocks from the previous step are closed *before* creating the new block, allowing for complex nesting logic within a single line.

### Candidate Wiki Hints
- **Page: CommonMark Parsing Strategy** (Concept: Two-phase parsing model)
- **Page: Delimiter Stack Algorithm** (Concept: Handling nested emphasis and links)
- **Page: Soft Line Breaks** (Concept: Rendering differences between soft and hard breaks)
