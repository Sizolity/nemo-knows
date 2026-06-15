---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
The document discusses the specification of list items, specifically focusing on indentation rules for sublists, block elements within lists, and the distinction between "tight" and "loose" lists. It contrasts the proposed CommonMark rules with John Gruber's original Markdown spec and Markdown.pl behavior. The text includes examples demonstrating how lists interrupt paragraphs, handle nested structures, and manage list markers of varying widths.

## Local Summary
This section defines how to determine if content belongs inside a list item versus starting a new one based on indentation relative to the list marker. It establishes the "four-space rule" as a principle for indented block elements (paragraphs, code blocks) within list items but introduces a more flexible strategy where indentation is measured from the end of the list marker. The text details how sublists must be indented sufficiently to nest, how lists can interrupt paragraphs (with restrictions on ordered numerals starting with 1), and how to separate consecutive lists using HTML comments or blank lines.

## Key Claims
- Sublists must be indented by a number of spaces equal to what a paragraph would need to be included in the parent list item.
- The "four-space rule" suggests block-level content under a list item requires four spaces of indentation, though this is arbitrary and potentially unintuitive for beginners.
- A common strategy allows the width and indentation of the list marker to determine the necessary indentation for blocks falling under the list item, rather than a fixed margin count.
- Lists may interrupt paragraphs in CommonMark, provided specific conditions regarding ordered list markers (starting with 1) are met to avoid spurious captures from hard-wrapped text.
- Changing the bullet or ordered list delimiter (e.g., from `-` to `+`, or `.` to `)` ) starts a new, separate list.

## Entities And Concepts
- **List Item**: A sequence of content marked by a list marker (`-`, `*`, `+`, `1.`, etc.).
- **Sublist**: A nested list within a list item; requires specific indentation relative to the parent list marker.
- **Tight vs. Loose List**: Tight lists have no blank lines between items or internal block elements; loose lists contain blank lines, resulting in wrapped `<p>` tags in HTML output.
- **Four-Space Rule**: An inference that all block elements under a list item must be indented four spaces from the margin.
- **Indentation Strategy**: Measuring indentation relative to the end of the list marker to accommodate variable marker widths (e.g., `10)` vs `-`).

## Procedures And API Details
- **Sublist Indentation**: Calculate required indentation by adding the width of the parent list marker and any initial indentation to the standard block indentation amount.
  - *Example*: A bullet list item with no extra indent requires two spaces for a sublist; an ordered list like `10)` requires four spaces total.
- **Separating Lists**: Insert a blank HTML comment (`<!-- -->`) between consecutive lists of the same type or to separate a list from an indented code block that would otherwise be parsed as a subparagraph.
- **List Interruption**: A paragraph can be followed immediately by a list without a blank line. Ordered lists interrupting paragraphs are restricted to those starting with `1` to prevent parsing text like "is 14." as a list.

## Nuance Or Contradictions
- **Markdown.pl vs. CommonMark**: Markdown.pl allowed only two spaces of indentation for sublists and exhibited inconsistent behavior (requiring three spaces for nested sublists). CommonMark adopts a more forgiving approach to handle both the strict four-space rule and Markdown.pl's legacy formatting, provided the layout is natural for humans.
- **Fixed vs. Relative Indentation**: A fixed indent from the margin (like the four-space rule) can lead to unintuitive results where text indented less than the list marker is excluded. A relative indent from the marker itself is proposed as superior but requires handling cases where the list item starts with indented code.
- **Hard-Wrapped Numerals**: While allowing lists to interrupt paragraphs solves natural usage patterns, it risks capturing unintended lists (e.g., "The number of windows in my house is 14."). The spec mitigates this by restricting interruption to ordered lists starting with `1`.

## Candidate Wiki Hints
- List Indentation Rules in CommonMark
- Tight vs. Loose Lists
- Handling Nested Sublists
