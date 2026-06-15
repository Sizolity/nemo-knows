---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
### Chunk Context
- **Heading path**: 3. c > foo\
- **Coverage**: Sections 6.8 (Soft line breaks), 6.9 (Textual content), and Appendix A (Parsing strategy).
- **Line range**: 7802-8133.

### Local Summary
This chunk details how CommonMark handles soft line breaks, preserves textual content verbatim, and outlines a two-phase parsing strategy (block structure then inline structure). It includes a detailed walkthrough of the delimiter stack algorithm used to resolve nested emphasis and links.

### Key Claims
- A regular line ending not preceded by two or more spaces or a backslash is parsed as a **softbreak**.
- Soft line breaks may be rendered in HTML as either a line ending or a space; the result is the same in browsers.
- Renderers may provide an option to render soft line breaks as hard line breaks.
- Any characters not given an interpretation by previous rules are parsed as plain textual content.
- Internal spaces are preserved verbatim.
- Parsing occurs in two phases: first constructing the block structure, then parsing raw text contents into inline elements.
- The delimiter stack is a doubly linked list used to track potential openers and closers for emphasis and links.

### Entities And Concepts
- **Softbreak**: A line ending parsed when not preceded by spaces or a backslash.
- **Textual content**: Characters parsed as plain text when no other interpretation applies.
- **Delimiter stack**: A doubly linked list tracking delimiters (`*`, `_`, `[`, `![]`) with metadata on activity and potential to open/close.
- **Openers_bottom**: A lower bound for searching delimiters of specific types during emphasis processing.
- **Lazy continuation**: A line added to an open block without closing it.

### Procedures And API Details
- **Parsing Strategy**:
  1. **Phase 1 (Block Structure)**: Lines are consumed to construct the document tree (paragraphs, block quotes, lists). Blocks are closed or new ones created based on line content.
  2. **Phase 2 (Inline Structure)**: Raw text in paragraphs and headings is parsed into inlines (strings, code spans, links, emphasis) using the link reference map from Phase 1.
- **Delimiter Stack Algorithm**:
  - When hitting `*`, `_`, `[`, or `![`, insert a text node and add a pointer to the delimiter stack.
  - When hitting `]`, call `look for link or image`.
  - When hitting end of input, call `process emphasis` with `stack_bottom = NULL`.
  - **look for link or image**: Search backwards for an active opening delimiter. If found and active, parse ahead for link/image types. If a link is found, set preceding `[` delimiters to inactive.
  - **process emphasis**: Iterate through the stack to find matching openers and closers. Determine if emphasis or strong emphasis exists based on delimiter run length. Insert `emph` or `strong emph` nodes and remove intermediate delimiters.

### Nuance Or Contradictions
- **Rendering Flexibility**: While the spec defines a soft line break, it explicitly allows renderers to choose between rendering it as a line ending or a space, noting that browser behavior makes the result identical.
- **Parser Implementation**: Renderers may offer an option to treat soft line breaks as hard line breaks, implying implementation flexibility beyond the strict parsing rules.

### Candidate Wiki Hints
- **Soft Line Breaks**: A dedicated page explaining the conditions for soft breaks and rendering variations.
- **Parsing Strategy**: A technical guide to the two-phase parsing model and tree construction.
- **Delimiter Stack Algorithm**: A deep dive into the logic for resolving nested emphasis and links.
