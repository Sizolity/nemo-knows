---
title: Commonmark Parsing Phases
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Parsing Phases

The CommonMark specification establishes a rigorous model for parsing Markdown, distinguishing between block-level structure and inline-level syntax. This approach ensures that the document is readable as plain text while resolving ambiguities present in earlier Markdown descriptions.

The parsing process follows a strict two-phase strategy. First, the parser constructs the block structure by consuming lines to identify headings, code blocks, lists, and other containers. Second, the parser processes the raw text contained within those blocks to identify inline elements such as emphasis, links, and code spans.

## Block Structure

The document is modeled as a tree of blocks, including the root document, block quotes, lists, and paragraphs. The parser handles various block-level elements, including ATX and Setext headings, indented and fenced code blocks, thematic breaks, and HTML blocks.

Container blocks, such as block quotes and lists, follow specific rules regarding indentation and "laziness." These rules allow markers to be omitted on continuation lines of paragraphs to improve readability. Indentation is relative to the start of a block or container, where four spaces typically trigger a code block, while fewer spaces may be used for container markers or skipped.

## Inline Elements

Once the block structure is established, the parser interprets the textual content within each block. This phase identifies inline elements like links, emphasis, and code spans. The resolution of these elements often relies on a delimiter stack algorithm to track active and inactive delimiters.

Characters not explicitly interpreted by previous rules are treated as plain textual content, preserving internal spaces verbatim. Special handling applies to tabs, which are treated as four spaces in specific contexts defining block structure, and backslashes, which escape ASCII punctuation characters to treat them as literals outside of code contexts.

## Global Scope

Certain definitions, such as link reference definitions, operate within the global scope of the entire document rather than being confined to the container in which they appear. HTML blocks are treated as raw text and can interrupt paragraphs, adhering to specific rules for start and end conditions.
