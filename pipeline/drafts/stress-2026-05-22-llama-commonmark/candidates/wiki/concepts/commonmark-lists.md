---
title: Commonmark Lists
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Lists

CommonMark lists are container blocks that organize content into items, governed by strict indentation rules to distinguish between list continuations, paragraph continuations, and indented code blocks. The specification defines a "four-space rule" where content indented four spaces from the margin belongs to the list item, while three spaces or fewer continue the preceding paragraph.

## Structure and Tightness

A list is classified as "loose" if items are separated by blank lines or if an item contains two block-level elements separated by a blank line. Otherwise, the list is considered "tight." The specification also addresses the "laziness" of markers in block quotes and lists to ensure consistent parsing.

## Indentation Rules

- **List Continuation:** Content indented four spaces from the margin is treated as part of the list item.
- **Paragraph Continuation:** Indentation of three spaces or fewer continues the paragraph within the list item.
- **Code Block:** Four spaces after a blank line create a code block rather than a continuation of the list item.

## Parsing Strategy

The CommonMark specification employs a two-phase parsing strategy. First, it constructs the block structure, including paragraphs, headings, lists, and code blocks. Second, it parses the raw text contents of these blocks into inline elements such as emphasis, links, and code spans. This separation ensures that block-level logic, like list tightness, is resolved before inline processing occurs.

## Related Concepts

- [[commonmark-block-structure]]
- [[commonmark-escaping-and-entities]]
- [[commonmark-inline-elements]]
- [[commonmark-line-breaks]]
