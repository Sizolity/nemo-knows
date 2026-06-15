---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context

This chunk covers Section 5 of the CommonMark specification, focusing on **Container Blocks**. It details the syntax and parsing rules for **Block Quotes** (5.1) and **List Items** (5.2). The text defines how these containers are constructed recursively from their contents, including specific rules for markers, indentation, laziness, and consecutiveness.

## Local Summary

The document defines container blocks as blocks containing other blocks, specifically block quotes and list items. It establishes recursive definitions for syntax rather than a direct parsing recipe. Section 5.1 outlines block quote markers (`>`), rules for basic cases, laziness (omitting markers on continuation lines), and consecutiveness (requiring blank lines between separate quotes). Section 5.2 defines list items via bullet (`-`, `+`, `*`) and ordered (digits followed by `.` or `)`) markers. It details four rules for list items: basic case, starting with indented code, starting with a blank line, and indentation handling. Laziness for list items is also described.

## Key Claims

- Container blocks are defined recursively based on their contents.
- Block quotes require a `>` marker followed by a space or no space, preceded by up to three spaces of indentation.
- List items are defined by markers (`-`, `+`, `*`, `1-9.`) followed by 1–4 spaces of indentation.
- Ordered list markers are limited to 9 digits to avoid integer overflows in browsers.
- Laziness allows omitting block quote markers on lines where the content would be paragraph continuation text.
- List items can contain any kind of block, including indented code blocks and block quotes.
- Indented code blocks within list items require specific indentation relative to the list marker.

## Entities And Concepts

- **Container Blocks**: Blocks that contain other blocks (block quotes, list items).
- **Block Quotes**: Defined by `>` markers; support laziness and nesting.
- **List Items**: Defined by bullet or ordered markers; support various content types.
- **Laziness**: Rule allowing omission of markers on continuation lines.
- **Consecutiveness**: Requirement for blank lines between separate block quotes.
- **Indented Code Blocks**: Code blocks indented by four spaces; treated specially within lists.
- **Thematic Breaks**: Lines that terminate list items.

## Procedures And API Details

- **Block Quote Marker**: `>` followed by a space or no space, with up to three spaces of indentation.
- **List Marker**: `-`, `+`, `*` for bullets; `1-9.` or `1-9)` for ordered lists.
- **Indentation Rules**:
  - List items: Marker width + 1–4 spaces.
  - Indented code blocks: Four spaces beyond the list item edge.
- **Laziness Application**: Remove markers from lines where the next non-space/tab character is paragraph continuation text.

## Nuance Or Contradictions

- **Blank Line Separation**: Block quotes must be separated by blank lines, unlike some Markdown implementations that merge them.
- **Lazy Continuation**: Markers can be omitted only on lines that would be paragraph continuations; they cannot be omitted before thematic breaks, code blocks, or list items.
- **Indentation Columns**: Indentation is relative to the list marker, not strictly column-based; content can be in the same column as the marker but still be inside the list if indented sufficiently past the last containing block.
- **Empty List Items**: Allowed at the start or end of a list but cannot interrupt a paragraph.

## Candidate Wiki Hints

- **Page: CommonMark Container Blocks**
  - Summary: Overview of block quotes and list items as container blocks.
- **Page: CommonMark Block Quotes**
  - Summary: Syntax, laziness, and nesting rules for block quotes.
- **Page: CommonMark List Items**
  - Summary: Marker types, indentation, and content rules for lists.
