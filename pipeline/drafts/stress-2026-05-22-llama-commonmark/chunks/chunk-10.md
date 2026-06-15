---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Section 3.c of the CommonMark specification, covering inline emphasis and strong emphasis rules.
- Lines 5732–6635.
- Heading path: 3. c.

Local Summary
This chunk details the parsing rules for emphasis (`*`, `_`) and strong emphasis (`**`, `__`) in Markdown. It covers delimiter matching, nesting, whitespace restrictions, and interactions with other inline elements like links and code spans.

Key Claims
- Emphasis and strong emphasis require matching delimiter runs; unmatched delimiters result in literal characters.
- Delimiters cannot be preceded by punctuation or followed by alphanumeric characters unless specific conditions are met.
- Intraword emphasis (e.g., `**foo**bar`) is forbidden.
- Nested emphasis is allowed with indefinite levels, provided delimiter lengths match specific conditions.
- Links bind more tightly than brackets in link text, which in turn bind more tightly than emphasis markers.

Entities And Concepts
- Emphasis: Text wrapped in single delimiters (`*` or `_`).
- Strong Emphasis: Text wrapped in double delimiters (`**` or `__`).
- Delimiter Runs: Sequences of characters used to denote emphasis.
- Link Text: Content within square brackets `[]`.
- Link Destination: URI following the link text.
- Inline Elements: Textual components like links, code spans, and images.

Procedures And API Details
- Rule 8: Closing delimiter preceded by whitespace is not valid for strong emphasis.
- Rule 9: Nonempty sequences of inline elements can be contents of an emphasized span.
- Rule 10: Nonempty sequences of inline elements can be contents of a strongly emphasized span.
- Rule 11: Determines handling of excess literal characters when delimiters do not match evenly.
- Rule 12: Similar to Rule 11 but for underscores.
- Rule 13: Requires different delimiters for emphasis nested directly inside emphasis.
- Rule 14: Allows strong emphasis within strong emphasis without switching delimiters.
- Rule 15: Handles cases where delimiters do not match evenly in nested structures.
- Rule 16: Deals with cases where delimiters are not properly matched in nested structures.
- Rule 17: Addresses interactions between emphasis markers and links.

Nuance Or Contradictions
- Emphasis and strong emphasis can be nested, but the rules for nesting differ based on the type of delimiter used.
- Links within link text must be handled carefully to avoid unintended nesting issues.
- Backslash escapes and entity references can be used in titles but not in link destinations without proper handling.

Candidate Wiki Hints
- Emphasis and Strong Emphasis Rules
- Link Text and Destination Parsing
- Nested Inline Elements
- Delimiter Matching Logic
