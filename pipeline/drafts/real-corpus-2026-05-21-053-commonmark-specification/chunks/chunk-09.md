---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
Section 3. c (Loose Lists) and Section 6 (Inlines, Code Spans, Emphasis). The chunk details the parsing logic for loose versus tight lists based on blank lines and block-level content within items, followed by the syntax rules for code spans and emphasis/strong emphasis delimiters.

Local Summary
This section defines how Markdown distinguishes between "loose" and "tight" list items based on internal structure (blank lines, block elements) rather than just line breaks. It then transitions to inline parsing, specifically defining code spans using backticks and the complex algorithm for interpreting asterisks (*) and underscores (_) as emphasis or strong emphasis markers.

Key Claims
- A list is loose if a blank line exists between items, or if an item contains two block-level elements separated by a blank line (even without a physical blank line in the raw text).
- Code spans are delimited by backtick strings; contents are normalized by converting line endings to spaces and stripping single leading/trailing spaces if present on both sides.
- Emphasis rules rely on "delimiter runs" that must be left-flanking or right-flanking depending on surrounding whitespace and punctuation.
- Backslash escapes do not work inside code spans; backslashes are treated literally.
- Inline code spans have higher precedence than emphasis markers, links, images, and HTML tags (except autolinks).

Entities And Concepts
- Loose List: A list where items are separated by blank lines or contain internal block-level separation.
- Tight List: A list where items do not contain internal blank lines separating block elements.
- Delimiter Run: A sequence of `*` or `_` characters not preceded/followed by a non-backslash-escaped version of itself.
- Left-flanking delimiter run: Not followed by whitespace, and either not followed by punctuation OR followed by punctuation and preceded by whitespace/punctuation.
- Right-flanking delimiter run: Not preceded by whitespace, and either not preceded by punctuation OR preceded by punctuation and followed by whitespace/punctuation.
- Code Span: Text enclosed in matching backtick strings with normalized content.

Procedures And API Details
- Code Span Parsing:
  1. Identify a backtick string (one or more backticks) not adjacent to another backtick.
  2. Match the opening and closing backtick strings of equal length.
  3. Normalize contents: convert line endings to spaces.
  4. Strip single leading/trailing space if the normalized string starts and ends with a space but is not all spaces.
- Emphasis Parsing Rules (Simplified):
  - Opening `*` requires a left-flanking delimiter run.
  - Closing `*` requires a right-flanking delimiter run.
  - Opening `_` requires left-flanking, closing `_` requires right-flanking (with specific punctuation exceptions).
  - Strong emphasis (`**`, `__`) follows similar flanking rules but with stricter spacing requirements for `__`.

Nuance Or Contradictions
- Intraword emphasis: Asterisks allow emphasis inside words (e.g., `foo*bar*`), while underscores do not (e.g., `foo_bar_`).
- Flanking logic is counter-intuitive: a delimiter followed by punctuation can act as left-flanking if preceded by whitespace, whereas standard Markdown might expect whitespace to be the primary separator.
- Ambiguity resolution prefers minimizing nesting depth (`<strong>` over `<em><em>`).

Candidate Wiki Hints
- CommonMark Specification: Loose vs Tight Lists
- CommonMark Syntax: Code Spans and Backticks
- Emphasis Algorithms: Delimiter Runs and Flanking Rules
