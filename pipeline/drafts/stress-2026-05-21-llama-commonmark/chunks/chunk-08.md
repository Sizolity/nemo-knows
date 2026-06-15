---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
This chunk covers sections 2.b and 3.c of the CommonMark specification, focusing on the indentation rules for list items. It demonstrates how list items are parsed when preceded by varying amounts of indentation (0, 1-3 spaces, and 4+ spaces) and how blank lines affect the interpretation of subsequent text.

## Local Summary
The text explains that list items cannot be preceded by more than three spaces of indentation. If a line is indented by more than three spaces, it is treated as a paragraph continuation rather than a new list item. If a line is indented by four or more spaces and preceded by a blank line, it is treated as an indented code block.

## Key Claims
- List items may not be preceded by more than three spaces of indentation.
- Indentation of more than three spaces causes a line to be treated as a paragraph continuation.
- Indentation of four or more spaces, when preceded by a blank line, creates an indented code block.

## Entities And Concepts
- List items
- Indentation
- Paragraph continuation
- Indented code block
- Blank line

## Procedures And API Details
- **Indentation Rule**: Check the number of spaces preceding a list item.
  - 0-3 spaces: Valid list item continuation or new item.
  - >3 spaces: Treated as paragraph continuation.
  - >=4 spaces (with blank line): Treated as indented code block.

## Nuance Or Contradictions
The specification clarifies that the distinction between a list item continuation and a code block depends on both the amount of indentation and the presence of a preceding blank line.

## Candidate Wiki Hints
- **CommonMark Indentation Rules**: A page detailing how indentation affects list parsing in CommonMark.
- **Indented Code Blocks**: A guide on when text is interpreted as code versus paragraph text.
