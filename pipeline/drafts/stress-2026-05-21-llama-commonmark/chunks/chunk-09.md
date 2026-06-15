---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Section 3.c defines loose and tight lists.
- Section 6 introduces inline parsing.
- Section 6.2 details emphasis and strong emphasis rules.
- Lines 4799-5730 cover list examples and inline syntax definitions.

Local Summary
This chunk explains how Commonmark determines whether a list is "loose" or "tight" based on blank lines and block-level content within items. It then transitions to inline parsing, specifically focusing on code spans and the complex rules for emphasis and strong emphasis using asterisks and underscores.

Key Claims
- A list is loose if there is a blank line between items or if an item contains multiple block-level elements separated by blank lines.
- A list is tight if blank lines appear only within code blocks, block quotes, or between paragraphs of a sublist.
- Inline parsing proceeds sequentially from left to right.
- Code spans are delimited by backticks and normalize content by converting line endings to spaces and stripping leading/trailing spaces if present on both sides.
- Emphasis and strong emphasis rely on "delimiter runs" that are classified as left-flanking, right-flanking, or both, based on surrounding whitespace and punctuation.
- Ambiguities in emphasis parsing are resolved by minimizing nesting depth and preferring shorter spans when closing delimiters match.

Entities And Concepts
- Loose list
- Tight list
- Code span
- Backtick string
- Delimiter run
- Left-flanking delimiter run
- Right-flanking delimiter run
- Emphasis
- Strong emphasis
- Unicode whitespace
- Unicode punctuation

Procedures And API Details
- **Code Span Normalization**:
  1. Convert line endings to spaces.
  2. If the string begins and ends with a space (but is not all spaces), remove one space from each end.
- **Emphasis Opening Rules**:
  - `*` opens emphasis if part of a left-flanking delimiter run.
  - `_` opens emphasis if part of a left-flanking delimiter run and either not right-flanking or right-flanking preceded by punctuation.
- **Emphasis Closing Rules**:
  - `*` closes emphasis if part of a right-flanking delimiter run.
  - `_` closes emphasis if part of a right-flanking delimiter run and either not left-flanking or left-flanking followed by punctuation.
- **Ambiguity Resolution**:
  - Minimize nestings (prefer `<strong>` over `<em><em>`).
  - Prefer `<em><strong>` over `<strong><em>`.
  - When spans overlap, the first one takes precedence.
  - When spans share a closing delimiter, the shorter one (opening later) takes precedence.
  - Inline code, links, images, and HTML tags group more tightly than emphasis.

Nuance Or Contradictions
- Intraword emphasis with `_` is generally disallowed to avoid unwanted emphasis in words containing internal underscores, unlike `*`.
- Backslash escapes do not work in code spans; all backslashes are treated literally.
- The stripping of leading/trailing spaces in code spans only occurs if spaces exist on both sides; single spaces are preserved if only one side has them.
- Browsers typically collapse consecutive spaces in `<code>` elements, so CSS `white-space: pre-wrap` is recommended.

Candidate Wiki Hints
- Page: Commonmark List Types (Loose vs Tight)
- Page: Commonmark Inline Syntax (Code Spans)
- Page: Commonmark Emphasis Rules (Delimiter Runs)
