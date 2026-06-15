---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Section 3. c (Inline Elements)
- Lines 5732–6635
- Covers Rules 8–17 (emphasis/strong emphasis delimiters), nesting logic, and Section 6.3 (Links).

Local Summary
- Defines how `*`, `**`, `_`, and `__` delimiters parse emphasis and strong emphasis, including nesting, mismatched delimiters, and whitespace/punctuation constraints.
- Clarifies that empty emphasis/strong emphasis is forbidden.
- Introduces Rules 13–17 for delimiter matching, nesting, and interactions with links/code spans.
- Section 6.3 defines link structure (text, destination, title), inline vs reference links, destination/title syntax, and parsing precedence (links bind looser than code spans but tighter than emphasis).

Key Claims
- Emphasis and strong emphasis can be nested indefinitely; nesting requires matching delimiter lengths unless both are multiples of 3.
- Empty emphasis/strong emphasis is invalid (`**`, `****`, `__`, `____` are not empty spans).
- Delimiters must not be preceded by punctuation or followed by alphanumeric characters to form emphasis.
- Links cannot contain nested links; the innermost definition wins if multiple nested definitions exist.
- Link text may contain inline elements but not nested links; brackets in link text must be balanced or escaped.
- Link destinations may contain spaces only if enclosed in `<...>`; line endings are forbidden in destinations.
- Link titles may use `""`, `''`, or `()` delimiters; nested quotes require escaping or alternate quote types.

Entities And Concepts
- Emphasis (`*`, `_`)
- Strong emphasis (`**`, `__`)
- Nested emphasis/strong emphasis
- Delimiter runs and matching logic
- Inline elements (links, code spans, images)
- Link text, destination, title
- Inline links vs reference links
- Pointy brackets `<...>`
- Backslash escapes
- Entity/numeric character references

Procedures And API Details
- Parsing emphasis:
  - Identify delimiter runs; sum lengths must not be a multiple of 3 unless both are multiples of 3.
  - Delimiters must not be preceded by punctuation or followed by alphanumeric characters.
  - Whitespace before closing delimiter invalidates strong emphasis with `__`.
- Parsing links:
  - Match `[...]` for link text; `(...)` for destination/title.
  - Destination can be `<...>` or plain text (no spaces unless in `<...>`).
  - Title uses `""`, `''`, or `()`.
  - Escapes (`\`) and entity references apply within destination/title.
- Precedence:
  - Code spans, autolinks, raw HTML bind tighter than link brackets.
  - Link brackets bind tighter than emphasis/strong emphasis markers.

Nuance Or Contradictions
- Rule 11/12: Excess literal `*` or `_` characters appear outside emphasis when delimiters mismatch.
- Rule 13: Nested emphasis inside emphasis requires different delimiters (`*_foo_*`), but nested strong emphasis within strong emphasis can reuse `****`.
- Markdown.pl allowed double quotes inside double-quoted titles; this spec rejects that for simplicity.
- Non-breaking space in titles breaks parsing; only standard spaces/tabs/one line ending allowed.

Candidate Wiki Hints
- Page: **CommonMark Inline Emphasis Rules** (Rules 8–17, nesting, delimiter matching)
- Page: **CommonMark Link Syntax** (inline links, destinations, titles, escaping)
- Page: **CommonMark Parsing Precedence** (code spans, links, emphasis binding order)
