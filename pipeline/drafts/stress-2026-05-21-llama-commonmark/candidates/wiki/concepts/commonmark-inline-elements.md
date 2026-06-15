---
title: Commonmark Inline Elements
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Inline Elements

CommonMark defines a rigorous parsing model that separates document structure into block-level and inline-level components. While block elements handle headings, lists, and code blocks, inline elements manage the syntax within paragraphs, such as emphasis, links, and code spans.

The specification treats the document as a tree of blocks. Once the block structure is established, the parser processes the raw text contained within those blocks to identify inline elements. This two-phase approach ensures that block-level formatting does not interfere with the interpretation of text-level markup.

## Syntax Rules

The specification establishes specific rules for character handling and escaping to ensure consistent rendering:

- **Escaping**: Backslashes allow ASCII punctuation characters to be treated as literals. However, this mechanism is invalid within code blocks, code spans, autolinks, and raw HTML.
- **Line Breaks**: A standard line ending without preceding spaces or a backslash creates a soft break. This may render as a line ending or a space in HTML output.
- **Textual Content**: Characters not matched by specific inline rules are interpreted as plain text, preserving internal spaces verbatim.

## Parsing Mechanics

Resolving complex inline structures relies on a delimiter stack algorithm. This mechanism tracks the state of delimiters to correctly identify nested emphasis and links. The parser distinguishes between active and inactive delimiters to determine the scope of formatting elements.

## Related Concepts

- [[commonmark-parsing-phases]]
- [[commonmark-code-blocks]]
- [[commonmark-html-blocks]]
