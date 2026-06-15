---
title: Commonmark Indentation Rules
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Indentation Rules

In the CommonMark specification, indentation serves as a primary mechanism for defining block-level structure, particularly for code blocks and container elements. Unlike the original Markdown description, which allowed for significant implementation divergence, this formal standard establishes a rigorous parsing model to ensure consistent behavior across different renderers.

## Code Blocks

Indentation is the defining characteristic for indented code blocks. When a line begins with four or more spaces, or a tab character, the parser treats the content as code. Tabs are not globally expanded but are specifically interpreted as equivalent to four spaces within the context of determining block structure. This rule allows authors to distinguish code from prose without relying on backticks or tildes, provided the indentation is maintained consistently.

## Container Blocks

Container blocks, such as block quotes and lists, rely on indentation relative to the start of the container to determine their boundaries. The specification employs a concept known as "laziness" to improve readability. This allows the omission of markers, such as the `>` character for block quotes or the `-` for list items, on continuation lines of paragraphs. Consequently, a paragraph inside a block quote can be indented relative to the quote marker without requiring the marker itself to be repeated on every line.

## Relative Indentation

Indentation is always calculated relative to the beginning of the current block or container. For example, inside a list item, the indentation required to start a nested list or a code block is measured from the start of the list item's text, not the document margin. This relative approach prevents accidental nesting of blocks and ensures that content remains within its intended structural hierarchy.

## Textual Content

Any characters that do not trigger a specific block-level rule are parsed as plain textual content. Within this content, internal spaces are preserved verbatim. However, a line ending without a preceding backslash or two spaces is treated as a soft line break, which may render as a space or a line break in HTML output, depending on the rendering engine.

## Escaping and Indentation

Backslashes are used to escape ASCII punctuation characters, treating them as literals rather than markup delimiters. This mechanism is valid in most contexts but is explicitly invalid within code blocks, code spans, autolinks, and raw HTML blocks. Therefore, indentation rules apply strictly to the block structure definition, while escaping rules govern the inline content within those blocks.
