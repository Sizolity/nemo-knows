---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
This chunk covers the CommonMark specification's handling of backslash escapes, entity and numeric character references, and the structural parsing of blocks (thematic breaks, ATX headings) and their precedence rules.

Local Summary
The text details how backslashes function as escape characters in CommonMark, noting their limitations in code contexts. It defines rules for HTML entity and numeric character references, specifying where they are valid and where they are treated as literal text. Finally, it outlines the syntax for thematic breaks and ATX headings, including indentation limits, character requirements, and precedence rules when structural elements overlap.

Key Claims
- A backslash escapes the following character unless the backslash itself is escaped.
- Backslash escapes are invalid within code blocks, code spans, autolinks, and raw HTML.
- Entity and numeric character references are valid in most contexts but cannot replace structural symbols (e.g., `*` for emphasis, `#` for headings).
- ATX headings require 1–6 unescaped `#` characters at the start, followed by spaces or tabs.
- Thematic breaks consist of 3+ matching `-`, `_`, or `*` characters with optional indentation (up to 3 spaces).

Entities And Concepts
- Backslash Escape: Mechanism to neutralize special characters like `*` or `\n`.
- Entity Reference: HTML5 named entities (e.g., `&copy;`) used to represent Unicode characters.
- Numeric Character Reference: Decimal (`&#123;`) or hexadecimal (`&#x7F;`) representations of Unicode code points.
- ATX Heading: Headings defined by `#` characters (e.g., `# foo`).
- Thematic Break: Horizontal rule defined by sequences of `-`, `_`, or `*`.
- Leaf Block: A block element that cannot contain other blocks (e.g., paragraphs, headings).
- Container Block: A block element that can contain other blocks (e.g., block quotes, lists).

Procedures And API Details
- **Backslash Escaping**:
  - `\*emphasis*` renders as `<p>\<em>emphasis</em></p>`.
  - `foo\` followed by a newline renders as `<p>foo<br />\nbar</p>`.
  - Escaping fails in code spans: `` `\[\` `` renders as `<code>\[\`</code>`.
- **Entity Parsing**:
  - Valid: `&amp;`, `&#65;`, `&#x41;`.
  - Invalid (treated as literal): `&copy` (missing semicolon), `&MadeUpEntity;`.
  - Restricted: `&#42;` in `&#42;foo&#42;` renders as `*foo*`, not `&#42;foo&#42;`.
- **Heading Parsing**:
  - Opening: 1–6 `#` chars, optional spaces/tabs.
  - Closing: Optional `#` chars, must be preceded by space/tab.
  - Indentation: Max 3 spaces; 4 spaces forces code block interpretation.
- **Thematic Break Parsing**:
  - Pattern: `[0-3 spaces][3+ matching chars][0-3 spaces]`.
  - Precedence: If a line matches both a thematic break and a setext heading underline, the heading wins.

Nuance Or Contradictions
- **Backslash Scope**: While backslashes escape characters in text, they have no effect in code blocks or code spans.
- **Entity Ambiguity**: HTML5 allows entities without semicolons (e.g., `&copy`), but CommonMark requires the semicolon to avoid grammar ambiguity.
- **Structural Override**: Entity references cannot define structural elements; `&#42;` inside `&#42;foo&#42;` does not create emphasis delimiters.
- **Heading Indentation**: Indenting a heading line by 4 spaces converts it from a heading to a code block, regardless of content.

Candidate Wiki Hints
- **Backslash Escaping Rules**: A guide on when and how to use backslashes to escape special characters in Markdown.
- **Entity Reference Usage**: Best practices for using HTML entities in URLs, titles, and text, excluding code blocks.
- **ATX Heading Syntax**: A reference for creating headings with `#` characters, including closing sequences and indentation limits.
- **Thematic Break Requirements**: Guidelines for creating horizontal rules using `-`, `_`, or `*`.
