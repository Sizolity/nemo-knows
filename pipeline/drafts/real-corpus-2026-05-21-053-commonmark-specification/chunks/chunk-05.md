---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
## Chunk Context
This chunk covers link reference definitions (Examples 215–218), paragraph formation rules (Section 4.8, Examples 219–227), and the handling of blank lines (Section 4.9).

## Local Summary
Link reference definitions can follow each other without blank lines and may appear inside block containers like lists or quotations; they apply globally to the document. Paragraphs are sequences of non-blank lines interpreted as inlines, where initial/final spaces are removed, multiple blank lines between paragraphs are ignored, and indentation rules distinguish between inline text and indented code blocks. Blank lines generally do not affect paragraph boundaries but determine list tightness.

## Key Claims
- Link reference definitions can occur consecutively without intervening blank lines.
- Definitions inside block containers (e.g., quotations) affect the entire document, not just the container.
- A paragraph is a sequence of non-blank lines that cannot be interpreted as other block types.
- Paragraph raw content is formed by concatenating lines and removing initial/final spaces or tabs.
- Multiple blank lines between paragraphs have no effect on paragraph boundaries.
- Lines after the first in a paragraph may be indented any amount; only the first line allows up to three spaces of indentation before being treated as code.
- Final spaces or tabs are stripped before inline parsing, so trailing spaces do not create hard line breaks within a paragraph.

## Entities And Concepts
- Link reference definitions
- Paragraphs
- Block containers (lists, quotations)
- Indented code blocks
- Hard line breaks (via two or more final spaces—though stripped in paragraphs)

## Procedures And API Details
1. **Link Reference Definitions**: Place `[id]: url` lines; multiple can follow consecutively. They may appear inside block containers.
2. **Paragraph Formation**:
   - Collect consecutive non-blank lines.
   - Concatenate lines and remove initial/final spaces or tabs.
   - Interpret the result as inlines.
3. **Indentation Rules for Paragraphs**:
   - First line: up to 3 leading spaces allowed; 4+ spaces start an indented code block.
   - Subsequent lines: any indentation allowed.
4. **Blank Lines**:
   - Ignored between block-level elements except for list tightness/looseness.
   - Ignored at document start/end.

## Nuance Or Contradictions
- Trailing spaces in a paragraph are stripped before inline parsing, so they do not produce hard line breaks within paragraphs (unlike in some other contexts).
- Blank lines between paragraphs are ignored; only the sequence of non-blank lines defines paragraph boundaries.

## Candidate Wiki Hints
- **Link Reference Definitions**: Global definitions can be placed anywhere; order and placement matter for resolution.
- **Paragraph Structure**: Spaces/tabs handling and indentation rules are critical for distinguishing text from code.
- **Blank Line Semantics**: Blank lines are mostly ignored except in list formatting contexts.
