---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Heading path: 2. b
- Heading coverage: 2. b, 3. c
- Line range: 4756-4798
- Source file: raw/web/corpus-2026-05-18/053-commonmark-specification.md

Local Summary
This chunk illustrates the indentation rules for ordered and unordered list items in CommonMark. It demonstrates that list items cannot be preceded by more than three spaces of indentation. If indented by more than three spaces, a line is treated as a paragraph continuation rather than part of the list item. Additionally, it shows that four spaces of indentation following a blank line creates an indented code block.

Key Claims
- List items may not be preceded by more than three spaces of indentation.
- Indentation exceeding three spaces causes the line to be treated as a paragraph continuation.
- Indentation of four spaces, when preceded by a blank line, creates an indented code block.

Entities And Concepts
- Ordered lists (`<ol>`)
- Unordered lists (`<ul>`)
- List items
- Indented code blocks
- Paragraph continuation lines

Procedures And API Details
- Indentation Rule: Keep indentation at three spaces or less for list item content.
- Code Block Trigger: Use four spaces of indentation after a blank line to start an indented code block.
- Example 312: Shows a list where item 'd' has a continuation line 'e' indented more than three spaces, resulting in 'e' being part of the paragraph within item 'd'.
- Example 313: Shows a list where a subsequent item is indented four spaces after a blank line, forming an indented code block.

Nuance Or Contradictions
- The distinction between three spaces (paragraph continuation) and four spaces (code block) is critical for correct parsing.
- Blank lines preceding indented content determine whether the indentation creates a code block or is ignored/interpreted differently.

Candidate Wiki Hints
- CommonMark Indentation Rules for Lists
- Indented Code Blocks vs. List Continuations
- Parsing List Item Continuation Lines
