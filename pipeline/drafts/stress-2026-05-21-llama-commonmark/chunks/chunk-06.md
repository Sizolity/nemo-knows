---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 5: Container blocks**, specifically detailing the syntax and parsing rules for **Block quotes (5.1)** and **List items (5.2)**. It defines how container blocks are constructed recursively from their contents, the specific markers required for block quotes (the `>` character), and the complex indentation and laziness rules governing list items.

# Local Summary

The text defines container blocks as blocks containing other blocks, specifically block quotes and list items. It establishes a recursive definition where a container is formed by transforming a sequence of blocks. Section 5.1 details block quote markers (`>`), allowing up to three spaces of indentation before the marker, and introduces the "Laziness" rule which permits omitting the marker on continuation lines of paragraphs. Section 5.2 defines list items via bullet (`-`, `+`, `*`) or ordered (1–9 digits) markers, explaining how indentation depth relative to the marker determines content inclusion, handling of indented code blocks within lists, and the "Laziness" rule for lists.

# Key Claims

- Container blocks are defined recursively by transforming a sequence of blocks.
- Block quotes require a `>` marker followed by a space or no space, preceded by up to three spaces of indentation.
- Block quotes exhibit "Laziness," allowing the `>` marker to be omitted on lines containing paragraph continuation text.
- List items are defined by markers (`-`, `+`, `*`, or `1`–`9` digits) followed by 1–4 spaces of indentation.
- Indentation depth is relative; content must be indented sufficiently to fall under the list item's edge, not just a fixed column.
- List items can contain any block type, including indented code blocks, block quotes, and thematic breaks (which terminate the list item).
- Ordered list start numbers are limited to 9 digits to avoid integer overflows in some browsers.
- List items can start with a blank line, but only one.
- Indented code blocks within list items require specific indentation relative to the list marker and the code block edge.

# Entities And Concepts

- **Container blocks**: Blocks that have other blocks as contents (block quotes, list items).
- **Block quote marker**: The `>` character.
- **Laziness**: A rule allowing omission of block quote markers on continuation lines of paragraphs.
- **List marker**: Bullet (`-`, `+`, `*`) or ordered (`1`–`9` digits + `.` or `)`) markers.
- **Indentation**: Spaces or tabs used to nest content within containers; relative indentation determines inclusion.
- **Indented code block**: A code block preceded by four spaces of indentation.
- **Thematic break**: A line that terminates a list item.
- **Paragraph continuation text**: Text parsed as part of a paragraph but not at the beginning.

# Procedures And API Details

- **Block Quote Construction**: Prepend `> ` (or just `>`) to lines. Omit `>` on continuation lines of paragraphs.
- **List Item Construction**: Prepend marker (e.g., `- ` or `1. `) to the first line. Indent subsequent lines by the width of the marker plus the spaces following it.
- **Indentation Calculation**: Calculate required indentation as `width of marker + spaces after marker`. Content must be indented at least this much to be included.
- **Ordered List Numbering**: Use 1–9 digits. Numbers starting with 0 are valid (e.g., `0.`). Negative numbers are invalid.
- **Code Block Handling**: Indented code blocks within lists require 4 spaces beyond the list item edge.

# Nuance Or Contradictions

- **Laziness vs. Structure**: The "Laziness" rule applies only to lines that would be paragraph continuations. It does not apply to lines starting new blocks (like lists or code blocks) or thematic breaks.
- **Indentation Columns**: One might assume content must align in a specific column, but the spec relies on relative indentation from the marker. Content can be in the same column as the marker but still be included if indented enough past the containing block's edge.
- **Blank Lines**: Blank lines separate block quotes but are not strictly required between a block quote and a following paragraph unless laziness rules apply. List items can contain multiple blank lines.
- **Empty List Items**: A list item can be empty (just a marker and a blank line), but an empty list item cannot interrupt a paragraph.

# Candidate Wiki Hints

- **Page: CommonMark Container Blocks**
  - Summary: Overview of block quotes and list items as container blocks.
- **Page: CommonMark Block Quote Syntax**
  - Summary: Rules for `>` markers, indentation limits, and the Laziness rule.
- **Page: CommonMark List Item Indentation**
  - Summary: How indentation depth determines content inclusion in list items.
- **Page: CommonMark Ordered List Numbers**
  - Summary: Constraints on ordered list start numbers (1–9 digits, no negatives).
