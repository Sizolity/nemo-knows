---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
This chunk covers the CommonMark specification section detailing list indentation rules. It specifically addresses how list items are interpreted based on the number of leading spaces (distinguishing between paragraph continuations and code blocks) and introduces examples involving ordered lists with blank lines.

Local Summary
The text clarifies that list items cannot be preceded by more than three spaces of indentation; exceeding this limit changes the interpretation of subsequent text. It provides an example where a fifth item 'e' is treated as a paragraph continuation within the previous item due to excessive indentation. Conversely, it shows that four spaces of indentation following a blank line creates an indented code block.

Key Claims
- List items may not be preceded by more than three spaces of indentation.
- Indenting more than three spaces causes the text to be treated as a paragraph continuation line within the list item context (in the specific example provided).
- Indenting four spaces and preceding with a blank line creates an indented code block.

Entities And Concepts
- List items
- Paragraph continuation line
- Indented code block
- CommonMark specification

Procedures And API Details
None.

Nuance Or Contradictions
The text presents specific examples where indentation thresholds (3 spaces vs 4 spaces) drastically alter the parsing of list structures, shifting content from being part of a list item to either a continuation or a code block.

Candidate Wiki Hints
- CommonMark List Indentation Rules
- Parsing Indented Code Blocks in Lists
