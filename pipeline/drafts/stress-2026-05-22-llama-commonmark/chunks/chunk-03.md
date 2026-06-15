---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context

This chunk covers sections 4.3 (Setext headings), 4.4 (Indented code blocks), and the beginning of 4.5 (Fenced code blocks) from the CommonMark specification. It details the syntax, parsing rules, and examples for these block-level elements, including interactions with paragraphs, lists, and block quotes.

## Local Summary

The text defines **setext headings** (using `=` or `-` underlines), **indented code blocks** (using 4+ spaces of indentation), and **fenced code blocks** (using backticks or tildes). It explains how to distinguish headings from thematic breaks or paragraphs, how indentation affects code block content, and the specific rules for opening and closing code fences.

## Key Claims

- **Setext Headings**: Consist of text lines followed by an underline (`=` for level 1, `-` for level 2). They cannot interrupt a paragraph; a blank line is required before a following heading if it follows a paragraph.
- **Indented Code Blocks**: Composed of non-blank lines preceded by four or more spaces. Content is literal text (no Markdown parsing). Ambiguity with list items favors the list item interpretation.
- **Fenced Code Blocks**: Begin with a code fence (3+ backticks or tildes). Content is literal. Closing fences must match the opening character and have at least as many characters. They can interrupt paragraphs.
- **Compatibility Note**: Most existing Markdown implementations do not support multi-line setext headings, though CommonMark does.

## Entities And Concepts

- **Setext Heading**: A heading defined by an underline (`=` or `-`) rather than `#` symbols.
- **Indented Code Block**: A code block defined by indentation (4+ spaces).
- **Fenced Code Block**: A code block defined by backticks or tildes.
- **Code Fence**: The line starting a fenced code block (3+ backticks/tildes).
- **Info String**: Optional text following the opening code fence, typically indicating the language.
- **Thematic Break**: A horizontal rule (`---` or `***`), which can be confused with setext headings if not separated by a blank line.

## Procedures And API Details

- **Setext Heading Parsing**:
  1. Check for text lines followed by an underline (`=` or `-`).
  2. Ensure the underline has no more than 3 spaces of indentation.
  3. Ensure the text lines are not interpretable as other block constructs (code fence, ATX heading, block quote, etc.).
  4. If the underline follows a paragraph, a blank line is required.
- **Indented Code Block Parsing**:
  1. Identify lines with 4+ spaces of indentation.
  2. Treat content as literal text (remove the 4 spaces).
  3. If a line has fewer than 4 spaces, the code block ends.
- **Fenced Code Block Parsing**:
  1. Identify a code fence (3+ backticks/tildes) with up to 3 spaces of indentation.
  2. Parse content until a closing fence of the same type and length (or longer) is found.
  3. Remove indentation from content lines if the opening fence was indented.

## Nuance Or Contradictions

- **Multi-line Headings**: CommonMark allows multi-line setext headings, but most existing Markdown implementations do not. This creates a compatibility gap.
- **Indentation Ambiguity**: If a line with 4+ spaces could be part of a list item, the list item interpretation takes precedence over a code block.
- **Closing Fence Indentation**: Closing fences can be indented up to 3 spaces and do not need to match the opening fence's indentation.
- **Internal Spaces in Fences**: Code fences cannot contain internal spaces or tabs; otherwise, they are not treated as fences.

## Candidate Wiki Hints

- **Setext Headings**: A reusable concept for defining headings with underlines.
- **Indented Code Blocks**: A distinct method for defining code blocks via indentation.
- **Fenced Code Blocks**: A flexible method for defining code blocks with fences.
- **Code Fence Syntax**: Rules for opening and closing code fences.
- **Info String**: Optional metadata for code blocks.
