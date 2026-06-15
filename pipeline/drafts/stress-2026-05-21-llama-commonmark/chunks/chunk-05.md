---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
## Chunk Context
This chunk covers link reference definitions (Examples 215–227) and the start of the Paragraphs section (4.8) and Blank lines (4.9) within the CommonMark specification.

## Local Summary
The text details how link reference definitions interact with inline content, specifically regarding placement inside block containers like blockquotes. It establishes that definitions affect the entire document scope, not just the container. The section then transitions to defining paragraphs as sequences of non-blank lines that cannot be interpreted as other blocks, detailing rules for line breaks, indentation, and blank line handling.

## Key Claims
- Link reference definitions can occur consecutively without intervening blank lines.
- Definitions inside block containers (e.g., blockquotes) affect the entire document, not just the container.
- A paragraph is a sequence of non-blank lines that cannot be interpreted as other block types.
- Multiple blank lines between paragraphs have no effect on the output.
- Leading spaces or tabs are skipped when determining paragraph content.
- Indented code blocks cannot interrupt paragraphs, but four spaces of indentation on the first line creates a code block instead.
- Final spaces or tabs are stripped before inline parsing, preventing hard line breaks if two or more spaces are present.

## Entities And Concepts
- Link Reference Definition
- Block Container
- Blockquote
- Paragraph
- Indented Code Block
- Hard Line Break

## Procedures And API Details
- **Paragraph Formation**: Concatenate lines, remove initial/final spaces/tabs.
- **Link Definition Scope**: Definitions are global; placement inside a blockquote does not limit their scope.
- **Indentation Rules**:
  - 0–3 spaces: Skipped (part of paragraph).
  - 4+ spaces: Creates an indented code block (interrupts paragraph).
- **Line Break Handling**:
  - Two or more spaces at end of line: Stripped (no hard break).
  - Blank lines: Ignored except for list tightness/looseness.

## Nuance Or Contradictions
- **Indentation Ambiguity**: While lines after the first in a paragraph may be indented any amount, the first line is restricted to up to three spaces of indentation to remain part of the paragraph.
- **Blank Line Role**: Blank lines are generally ignored between block elements but are critical for determining if a list is "tight" or "loose."

## Candidate Wiki Hints
- **Link Reference Definitions**: A dedicated page explaining global vs. local link definitions and their placement rules.
- **Paragraph Parsing**: A guide on how CommonMark handles line breaks, indentation, and blank lines within paragraphs.
