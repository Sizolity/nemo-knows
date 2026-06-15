---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Section 3.c defines loose and tight lists.
- Section 6 introduces inlines, code spans, and emphasis rules.
- Lines 4799-5730 cover list examples (314–327) and inline syntax (6.1–6.2).

Local Summary
This chunk details how Commonmark determines if a list is "loose" or "tight" based on blank lines and block content within items. It then transitions to inline parsing, specifically code spans (backticks) and the complex rules for emphasis (asterisks and underscores), including delimiter runs and flanking conditions.

Key Claims
- A list is loose if there is a blank line between items, or if an item contains two block-level elements with a blank line between them.
- A list is tight if blank lines occur only within code blocks, block quotes, or between paragraphs of a sublist.
- Code spans are delimited by backticks; the delimiter strings must be equal in length.
- Leading and trailing spaces are stripped from code spans only if present on both sides.
- Emphasis and strong emphasis rely on "delimiter runs" which must be left-flanking (to open) and right-flanking (to close).
- Intraword emphasis with underscores is disallowed, while asterisks allow it.
- Inline code spans, links, images, and HTML tags have higher precedence than emphasis.

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
- Intraword emphasis
- Unicode whitespace
- Unicode punctuation

Procedures And API Details
- **Code Span Parsing**:
  1. Identify a backtick string (one or more backticks) not preceded/followed by a backtick.
  2. Match opening and closing backtick strings of equal length.
  3. Normalize contents: convert line endings to spaces.
  4. Strip a single leading and trailing space if the string both begins and ends with a space (and is not all spaces).
  5. Treat line endings as spaces for normalization.
- **Emphasis Parsing**:
  1. Identify delimiter runs (sequences of `*` or `_` not preceded/followed by non-escaped versions).
  2. Classify runs as left-flanking, right-flanking, or both based on surrounding whitespace and punctuation.
  3. Apply rules 1–12 to determine opening and closing delimiters.
  4. Resolve ambiguities using principles 13–17 (minimize nesting, prefer outer emphasis, prefer shorter span, respect higher precedence constructs).

Nuance Or Contradictions
- **Space Stripping**: Only ASCII spaces are stripped from code spans, not general Unicode whitespace.
- **Underscore Restrictions**: Intraword emphasis is forbidden for `_` but allowed for `*`.
- **Delimiter Flanking**: A delimiter run can be both left- and right-flanking, but its ability to open/close depends on the specific character before/after (whitespace vs punctuation).
- **Precedence**: Code spans and HTML tags interrupt emphasis patterns; e.g., `*[foo*](bar)` is parsed as `*<a href="bar">foo*</a>`, not nested emphasis.

Candidate Wiki Hints
- **Loose vs Tight Lists**: Explain the structural difference and provide examples of when blank lines inside items affect list tightness.
- **Code Span Normalization**: Document the specific rules for stripping spaces and handling line endings in code spans.
- **Emphasis Delimiter Runs**: Create a guide on determining left/right-flanking status and how punctuation affects emphasis validity.
- **Inline Precedence**: Highlight the hierarchy where code, links, and HTML tags override emphasis markers.
