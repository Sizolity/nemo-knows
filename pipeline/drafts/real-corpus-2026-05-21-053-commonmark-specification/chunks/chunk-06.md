---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 5: Container blocks**, specifically defining the syntax for block quotes and list items. It details how these structures are built recursively from their contents, including rules for markers, indentation, laziness, consecutiveness, and specific behaviors regarding code blocks, thematic breaks, and nested structures.

# Local Summary

The specification defines container blocks as those containing other blocks (block quotes and list items). Block quotes are introduced via a `>` marker with specific rules for indentation, spacing, and "laziness" (omitting markers on continuation lines). List items are defined by bullet (`-`, `+`, `*`) or ordered markers (1–9 digits), with complex rules governing indentation requirements relative to the list marker width. The chunk includes numerous examples illustrating how blank lines separate block quotes, how lazy continuation lines work within nested structures, and the specific indentation math required to keep content inside a list item versus outside it.

# Key Claims

- Container blocks are defined recursively by transforming a sequence of blocks into a container of type Y.
- A block quote marker is `>` followed by a space, or just `>`. Up to three spaces of indentation may precede the marker; four spaces creates a code block instead.
- List markers for bullets are `-`, `+`, or `*`. Ordered list markers use 1–9 digits followed by `.` or `)`.
- The "Laziness" rule allows omitting the `>` marker on lines containing paragraph continuation text, provided the indentation requirements relative to containing blocks are met.
- List item indentation is relative: content must be indented sufficiently past the list marker and any preceding containers (like blockquotes).
- Ordered list start numbers are limited to 9 digits due to browser integer overflow concerns; leading zeros are allowed.
- A list item cannot interrupt a paragraph unless it starts with specific conditions, and thematic breaks terminate list items.

# Entities And Concepts

- **Container blocks**: Blocks that hold other blocks as contents.
- **Block quotes**: Defined by `>` markers, supporting laziness and nesting.
- **List items**: Defined by bullet or ordered markers, containing blocks separated by blank lines.
- **Lazy continuation lines**: Lines within a block quote or list item where the opening marker (`>`) is omitted if it would be redundant for paragraph text.
- **Indented code blocks**: Code blocks inside list items requiring specific indentation relative to the list marker.

# Procedures And API Details

**Block Quote Construction:**
1. Prepend `>` (optionally with 1–3 spaces of indentation) to lines.
2. Apply "Laziness": Remove initial `>` from continuation lines if they are paragraph text.
3. Ensure blank lines separate distinct block quotes ("Consecutiveness").

**List Item Construction:**
1. **Basic Case**: Prepend list marker `M` (width `W`) and `N` spaces (`1 ≤ N ≤ 4`) to the first line. Indent subsequent lines by `W + N`.
2. **Indented Code Start**: If starting with indented code, prepend `M` and exactly one space; indent subsequent lines by `W + 1`.
3. **Blank Line Start**: Prepend `M` to a blank line; indent subsequent lines by `W + 1`.
4. **Indentation Rule**: Lines may be uniformly indented by up to three spaces without changing the list item status.

**Validation Rules:**
- Ordered markers must use digits 0–9 (max 9 digits).
- Negative numbers (e.g., `-1.`) are invalid markers.
- Thematic breaks (`---`) end a list item.

# Nuance Or Contradictions

- **Laziness Limitation**: The `>` marker cannot be omitted on lines that start a new block type (like code blocks or thematic breaks) within the quote, even if indentation suggests it might continue.
- **Indentation Columns vs. Relative Indentation**: While one might assume content must align in a specific column, the rule is strictly about relative indentation past the last containing block marker. Content can appear far to the right and still be inside the list item if indented enough relative to the container edge, or far to the left and outside if not.
- **Empty List Items**: An empty list item (marker followed by blank line) is valid but cannot interrupt a paragraph; it must be part of a list context.

# Candidate Wiki Hints

- **Page: Commonmark Block Quotes** – Covers syntax, laziness, nesting, and separation rules.
- **Page: Commonmark List Items** – Covers marker types, indentation math, code block inclusion, and empty items.
