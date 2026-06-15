---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
## Chunk Context
This chunk covers link reference definitions (Examples 215–218) and the definition of paragraphs (Section 4.8) and blank lines (Section 4.9) within the CommonMark specification.

## Local Summary
The text explains that link reference definitions can appear consecutively without blank lines and can be nested inside block containers like quotations, affecting the entire document scope. It then defines a paragraph as a sequence of non-blank lines that cannot be interpreted as other blocks, detailing how raw content is processed (concatenation, space removal) and how indentation rules apply to distinguish paragraphs from code blocks. Finally, it notes that blank lines are generally ignored except for determining list tightness.

## Key Claims
- Link reference definitions can occur one after another without intervening blank lines.
- Link reference definitions inside block containers (e.g., lists, block quotations) affect the entire document, not just the container.
- A paragraph is formed by a sequence of non-blank lines that cannot be interpreted as other kinds of blocks.
- Paragraph raw content is formed by concatenating lines and removing initial and final spaces or tabs.
- Multiple blank lines between paragraphs have no effect on the output.
- Leading spaces or tabs are skipped; lines after the first may be indented any amount.
- Four spaces of indentation on the first line of a paragraph context creates a code block instead of a paragraph.
- Final spaces or tabs are stripped before inline parsing, preventing hard line breaks if two or more spaces are present.
- Blank lines between block-level elements are ignored except for determining list tightness.

## Entities And Concepts
- Link reference definitions
- Block containers (lists, block quotations)
- Paragraphs
- Raw content
- Indented code blocks
- Blank lines
- List tightness/looseness

## Procedures And API Details
- **Paragraph Formation**: Concatenate lines, remove initial/final spaces/tabs.
- **Indentation Rule**: First line may have up to three spaces; four spaces triggers a code block.
- **Hard Line Breaks**: Two or more spaces at the end of a line are stripped before inline parsing.

## Nuance Or Contradictions
- While blank lines are ignored for paragraph separation, they are critical for determining whether a list is tight or loose.
- Indentation rules allow subsequent lines in a paragraph to be indented arbitrarily, but the first line is restricted to three spaces of indentation.

## Candidate Wiki Hints
- **Link Reference Definitions**: Scope and placement rules.
- **Paragraph Parsing**: Indentation thresholds and whitespace handling.
- **Blank Line Semantics**: Role in list structure vs. general block separation.
