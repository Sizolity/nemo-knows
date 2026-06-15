---
title: Commonmark Code Blocks
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Code Blocks

In the CommonMark specification, code blocks are a primary form of block-level syntax designed to represent literal text content. They are categorized into two distinct types: indented code blocks and fenced code blocks. This distinction allows users to preserve whitespace and special characters exactly as written, distinguishing them from the surrounding prose.

## Indented Code Blocks

Indented code blocks are defined by lines that begin with at least four spaces of indentation relative to the start of the block or container. This indentation rule ensures that the content is treated as a code block rather than a continuation of a paragraph or list item. The specification treats tabs in these contexts specifically as equivalent to four spaces, ensuring consistent behavior across different input formats.

## Fenced Code Blocks

Fenced code blocks are delimited by a "fence" consisting of three or more backticks (```) or tildes (~~~). These fences may optionally include an information string, such as a language identifier, immediately following the opening fence. The content of the block extends from the line immediately following the opening fence to the line containing the closing fence.

## Escaping and Raw Text

Within the context of code blocks, the standard backslash escape mechanism is disabled. Unlike inline code spans or other contexts where a backslash escapes ASCII punctuation characters to treat them as literals, backslashes inside fenced or indented code blocks are treated as ordinary characters. This ensures that the raw text inside the block is preserved verbatim without interpretation of escape sequences.

## Parsing Context

The specification models the document as a tree of blocks, where code blocks occupy specific nodes within this hierarchy. The parsing process follows a two-phase strategy: first, the block structure is constructed by consuming lines and identifying block-level elements like code blocks; second, the raw text within those blocks is processed. For code blocks, the second phase largely involves preserving the text as-is, distinguishing them from paragraphs where inline elements like emphasis and links are parsed.
