---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
- **Source**: `raw/web/corpus-2026-05-18/053-commonmark-specification.md`
- **Chunk**: 3 of 13
- **Lines**: 1161-1978
- **Heading**: `baz > ` (covers sections 4.3 Setext headings, 4.4 Indented code blocks, and 4.5 Fenced code blocks)

## Local Summary
This chunk details the syntax and parsing rules for three distinct block-level Markdown constructs: **Setext headings**, **Indented code blocks**, and **Fenced code blocks**. It defines the structural requirements for each (e.g., underline characters for setext, four-space indentation for indented code, backtick/tilde sequences for fenced code) and explains how they interact with surrounding text, including rules for blank lines, indentation limits, and precedence over other block types.

## Key Claims
- **Setext Headings**: Defined by text lines followed by an underline of `=` (level 1) or `-` (level 2). They cannot interrupt paragraphs without a preceding blank line. Multiline content is supported, but the text must not be interpretable as other block constructs (like lists or thematic breaks).
- **Indented Code Blocks**: Composed of lines indented by at least four spaces. They are literal text (no Markdown parsing). They cannot interrupt paragraphs.
- **Fenced Code Blocks**: Defined by opening and closing fences of three or more backticks or tildes. They can interrupt paragraphs. The content is literal text, and the info string (language hint) is optional and not strictly mandated for rendering.
- **Precedence**: Block structure indicators (like lists or thematic breaks) take precedence over inline or heading indicators where ambiguity exists.

## Entities And Concepts
- **Setext Heading**: A heading style using underlines (`=` or `-`).
- **Indented Code Block**: A code block defined by 4+ space indentation.
- **Fenced Code Block**: A code block defined by backticks or tildes.
- **Info String**: Optional text following the opening fence in a fenced code block.
- **Code Fence**: The sequence of backticks or tildes delimiting a fenced code block.
- **Thematic Break**: A horizontal rule (`---` or `***`) that can conflict with setext heading underlines.

## Procedures And API Details
- **Setext Heading Parsing**:
  1. Check for text lines followed by an underline (`=` or `-`).
  2. Ensure the underline has no more than 3 spaces of indentation.
  3. Verify the text is not interpretable as a code fence, ATX heading, block quote, thematic break, list item, or HTML block.
  4. If followed by a paragraph, insert a blank line to separate them.
- **Indented Code Block Parsing**:
  1. Identify lines with 4+ spaces of indentation.
  2. Remove exactly 4 spaces of indentation from each line.
  3. Treat content as literal text.
  4. Stop the block if a line has fewer than 4 spaces of indentation.
- **Fenced Code Block Parsing**:
  1. Locate a code fence (3+ backticks or tildes) with up to 3 spaces of indentation.
  2. Extract the info string (trimmed of spaces/tabs).
  3. Collect subsequent lines until a matching closing fence is found.
  4. Remove indentation from content lines if the opening fence was indented.
  5. If no closing fence is found, the block extends to the end of the document/block.

## Nuance Or Contradictions
- **Multiline Setext Headings**: While the spec supports multiline headings (e.g., `Foo\nbar\n---`), most existing Markdown implementations do not. This creates a compatibility gap where users might expect a paragraph followed by a thematic break instead of a multiline heading.
- **Indentation Ambiguity**: In indented code blocks, if a line has more than 4 spaces of indentation, the extra spaces are preserved in the content. However, if a line has fewer than 4 spaces, the code block ends immediately.
- **Fence Matching**: The closing fence must use the same character (backticks or tildes) and be at least as long as the opening fence. Mixing characters or using fewer characters for the closing fence is invalid.
- **Internal Spaces in Fences**: Code fences cannot contain internal spaces or tabs (e.g., `` ``` ``` `` is invalid).

## Candidate Wiki Hints
- **Page: Setext Headings**
  - Summary: Rules for defining level 1 and 2 headings using underlines.
  - Key points: Underline characters (`=` vs `-`), indentation limits, multiline support, and interaction with paragraphs.
- **Page: Indented Code Blocks**
  - Summary: Syntax and behavior of code blocks defined by indentation.
  - Key points: 4-space indentation rule, literal text content, precedence over list items.
- **Page: Fenced Code Blocks**
  - Summary: Syntax and behavior of code blocks defined by backticks or tildes.
  - Key points: Fence length matching, info string handling, interruption of paragraphs.
- **Page: Block Precedence**
  - Summary: How CommonMark resolves conflicts between different block types (e.g., list vs. code block, thematic break vs. heading).
