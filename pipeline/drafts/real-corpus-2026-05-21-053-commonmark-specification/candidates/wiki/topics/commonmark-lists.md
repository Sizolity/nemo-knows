---
title: Commonmark Lists
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Lists

In the **CommonMark Specification**, lists are defined as a specific type of container block. The specification distinguishes between two primary forms: unordered lists (bullets) and ordered lists. These structures are governed by strict rules regarding tightness, laziness, and how they interact with surrounding paragraphs.

## Structure and Parsing

The parsing of list items occurs during the **Block Structure Phase**. This phase identifies container blocks such as lists and block quotes before processing the internal content of those blocks. List items can contain complex nested structures, including other lists, code blocks, thematic breaks, and paragraph blocks.

A key feature of CommonMark lists is the handling of "lazy continuation." This mechanism allows list items to continue across multiple lines or even blank lines without requiring explicit markers at every line break, provided the indentation rules are met. The specification defines specific indentation requirements relative to the bullet point or number to determine if a line belongs to the same list item or starts a new one.

## Interactions with Other Elements

Lists interact closely with other block-level elements:
-   **Block Quotes**: A list can be nested inside a block quote, and vice versa, though specific rules apply to how they are combined.
-   **Paragraphs**: A paragraph immediately following a list item is considered part of that item unless separated by a blank line or specific structural boundaries.
-   **HTML Blocks**: Lists generally do not interrupt HTML blocks in the same way they interrupt paragraphs, maintaining strict boundaries for raw HTML tags within the list context.

## Tightness and Looseness

The specification introduces concepts of tightness and looseness to handle multiple blank lines. Multiple blank lines between list items or between a list item and a following block quote define the structural separation (looseness) rather than creating new paragraphs within the list itself. This ensures consistent rendering where extra vertical space does not inadvertently alter the logical structure of the list content.

## See Also

- [[commonmark-inline-elements]]
- [[commonmark-parsing-phases]]
