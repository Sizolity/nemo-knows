---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Heading: aaa
- Lines: 4230–4754
- Source: raw/web/corpus-2026-05-18/053-commonmark-specification.md

Local Summary
This chunk details the CommonMark specification for list items, sublists, and list structure. It contrasts the "four-space rule" (indented blocks under a list item) with the historical "Markdown.pl" behavior (two-space indentation). It explains how list markers determine the required indentation for subsequent content, how lists interrupt paragraphs, and how to separate consecutive lists using blank HTML comments.

Key Claims
- A sublist must be indented the same number of spaces a paragraph would need to be included in the list item.
- The "four-space rule" requires block-level content (paragraphs, sublists, code) under a list item to be indented four spaces from the margin.
- CommonMark allows lists to interrupt paragraphs (e.g., `Foo\n- bar`), unlike Markdown.pl.
- Only ordered lists starting with `1` are allowed to interrupt paragraphs to avoid spurious list captures (e.g., `14.`).
- Changing the list marker character (e.g., `-` to `+`) or number format starts a new list.
- A list is "loose" if items are separated by blank lines or contain internal blank lines; otherwise, it is "tight."
- Blank HTML comments (`<!-- -->`) can separate consecutive lists or prevent indented code from being parsed as a subparagraph.

Entities And Concepts
- List Item: A block containing list markers and content.
- Sublist: A nested list within a list item.
- Four-Space Rule: The principle that content under a list item must be indented four spaces.
- Tight List: A list where items are not separated by blank lines.
- Loose List: A list where items are separated by blank lines.
- List Marker: The symbol (`-`, `+`, `*`, `1.`, etc.) starting a list item.
- HTML Comment: Used to separate lists (`<!-- -->`).

Procedures And API Details
- Indentation for Sublists: Match the indentation required for a paragraph to be included in the list item (determined by list marker width).
- Indentation for Code: Eight spaces from the margin (or four spaces from the list marker).
- Indentation for Blockquotes: Indented (typically four spaces).
- Separating Lists: Insert a blank HTML comment between lists to ensure they are treated as distinct.

Nuance Or Contradictions
- Markdown.pl allowed two-space indentation for sublists but was inconsistent (requiring three spaces for sub-sublists).
- The four-space rule is arbitrary and unintuitive for beginners but provides a principled standard.
- A two-space rule from the list marker would allow text indented less than the marker to be included, which is unintuitive.
- Indented code in a list item must be indented eight spaces from the margin to avoid breaking existing Markdown patterns.

Candidate Wiki Hints
- List Item Indentation Rules
- Four-Space Rule vs. Markdown.pl
- Loose vs. Tight Lists
- Separating Lists with HTML Comments
- Lists Interrupting Paragraphs
