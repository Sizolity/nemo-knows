---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Section: 5.3 Lists
- Lines: 4230–4754
- Heading: aaa
- Scope: Rules for sublists, indentation logic, loose vs tight lists, interrupting paragraphs, and list separation.

Local Summary
- Defines how sublists are indented relative to list markers.
- Explains the four-space rule vs Markdown.pl’s two-space behavior.
- Introduces loose vs tight lists and HTML output differences.
- Details how lists interrupt paragraphs and how to separate lists.

Key Claims
- Sublists must be indented enough to fit the list marker plus any initial indentation.
- The four-space rule is preferred over fixed indentation from the margin.
- A list is loose if items are separated by blank lines or contain internal blank lines.
- Lists can interrupt paragraphs in CommonMark, unlike Markdown.pl.
- Only lists starting with "1" are allowed to interrupt paragraphs to avoid spurious captures.
- Blank HTML comments can separate consecutive lists or prevent unintended code blocks.

Entities And Concepts
- List marker: Bullet (-, +, *) or ordered (., ))
- Four-space rule: Indentation requirement for blocks under list items.
- Loose list: Items separated by blank lines; paragraphs wrapped in `<p>`.
- Tight list: No internal blank lines; paragraphs not wrapped in `<p>`.
- Principle of uniformity: Text meaning remains consistent inside containers.
- Spurious list capture: Unintended list creation from hard-wrapped numerals.

Procedures And API Details
- Indentation calculation: Measure from the start of the list marker, not the margin.
- Code block indentation: Eight spaces from the margin (or six from the marker in some proposals).
- Separating lists: Insert `<!-- -->` between lists to reset parsing.
- List interruption: No blank line needed between paragraph and list.

Nuance Or Contradictions
- Markdown.pl allowed two-space indentation for sublists inconsistently.
- Different implementations (Pandoc, discount, redcarpet) handled indentation differently.
- Fixed four-space rule feels unnatural for some layouts.
- Two-space rule risks including unintended text in list items.
- Indented code inside lists requires special handling to avoid breaking existing patterns.

Candidate Wiki Hints
- List indentation rules in CommonMark
- Loose vs tight lists
- Lists interrupting paragraphs
- Separating lists with HTML comments
- Spurious list capture prevention
