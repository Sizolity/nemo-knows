---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
Section 3. c of the CommonMark specification, detailing rules for emphasis and strong emphasis delimiters (`*`, `_`, `**`, `__`), including nesting, mismatched delimiters, and interactions with links and code spans.

## Local Summary
This chunk outlines how emphasis and strong emphasis are parsed, focusing on delimiter matching, nesting constraints, and interactions with other inline elements like links, images, and HTML tags. It also introduces rules for link text parsing, including restrictions on nested links and bracket handling.

## Key Claims
- Emphasis and strong emphasis can be nested indefinitely, but empty emphasis is forbidden.
- Mismatched delimiters result in literal characters appearing outside the emphasized span.
- Links cannot contain other links at any nesting level; innermost link definitions take precedence if nested.
- Brackets in link text are allowed only if escaped or balanced with unescaped inline content.

## Entities And Concepts
- Emphasis: Span enclosed by single delimiters (`*` or `_`).
- Strong Emphasis: Span enclosed by double delimiters (`**` or `__`).
- Inline Link: A link with destination and title immediately following the link text.
- Reference Link: A link where destination and title are defined elsewhere.
- Delimiter Runs: Sequences of delimiters used to open/close emphasis spans.

## Procedures And API Details
- **Delimiter Matching**: Emphasis requires an odd number of delimiters; strong emphasis requires an even number. Mismatches result in literal characters outside the span.
- **Nesting Rules**: Different delimiter types (`*` vs `_`) must be used for nested emphasis within emphasis. Strong emphasis within strong emphasis can use the same delimiter type.
- **Link Text Parsing**: Brackets bind more tightly than emphasis markers; backtick code spans and autolinks also take precedence over brackets.

## Nuance Or Contradictions
- While links cannot contain other links, emphasis can nest indefinitely within links if delimiters are balanced correctly.
- Pointy brackets (`<...>`) in link destinations allow spaces but not line endings, unlike plain text destinations.
- Titles in links may use single quotes, double quotes, or parentheses, with nested balanced quotes requiring escaping unless using a different quote type.

## Candidate Wiki Hints
- **Nesting Rules for Emphasis**: A reusable guide on how to nest emphasis and strong emphasis correctly in Markdown.
- **Link Text Constraints**: Documentation on parsing link text, including bracket handling and precedence rules.
