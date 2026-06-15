---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Lines 517–1160 of the CommonMark specification.
- Covers backslash escaping rules, entity and numeric character references, and the structure of blocks and inlines.
- Includes detailed rules for thematic breaks, ATX headings, and precedence between block and inline structures.

Local Summary
This chunk defines how backslashes interact with special characters, clarifying that backslash escapes function in most contexts but not within code blocks, code spans, autolinks, or raw HTML. It details the parsing of HTML entity and numeric character references, noting their exclusion from code contexts and structural symbols. The text then introduces the block structure of documents, distinguishing between container and leaf blocks, and provides specific rules for thematic breaks and ATX headings, including indentation limits and character matching requirements.

Key Claims
- Backslash escapes are invalid in code blocks, code spans, autolinks, and raw HTML.
- Entity and numeric character references are valid in most contexts except code spans, code blocks, and structural symbols (e.g., emphasis delimiters, list markers).
- Thematic breaks require 3+ matching characters (-, _, or *) with optional spaces/tabs, and must not contain other characters.
- ATX headings consist of 1–6 unescaped # characters with optional closing #s, requiring at least one space/tab after the opening #s unless the heading is empty.
- Indentation of up to three spaces is allowed for thematic breaks and ATX headings; four spaces or more forces a code block.

Entities And Concepts
- Backslash escape: A mechanism to escape special characters, valid in most contexts but not in code blocks, code spans, autolinks, or raw HTML.
- Entity reference: HTML5 entity names (e.g., &amp;) used to represent Unicode characters.
- Numeric character reference: Decimal (e.g., &#35;) or hexadecimal (e.g., &#x22;) representations of Unicode characters.
- Thematic break: A horizontal rule formed by 3+ matching characters (-, _, or *) with optional spaces/tabs.
- ATX heading: A heading level 1–6 formed by # characters with optional closing #s.
- Container block: A block that can contain other blocks (e.g., block quotes, list items).
- Leaf block: A block that cannot contain other blocks (e.g., paragraphs, headings, thematic breaks).

Procedures And API Details
- Backslash escape procedure:
  - If a backslash is itself escaped, the following character is not escaped.
  - Backslash escapes do not work in code blocks, code spans, autolinks, or raw HTML.
  - Backslash escapes work in URLs, link titles, link references, and info strings in fenced code blocks.
- Entity reference parsing:
  - Valid HTML5 entity names are recognized.
  - Decimal numeric character references: &# + 1–7 arabic digits + ;.
  - Hexadecimal numeric character references: &# + X/x + 1–6 hexadecimal digits + ;.
  - Invalid Unicode code points are replaced by U+FFFD.
- Thematic break formation:
  - 3+ matching characters (-, _, or *) with optional spaces/tabs.
  - No other characters allowed in the line.
  - All characters other than spaces/tabs must be the same.
- ATX heading formation:
  - Opening sequence: 1–6 unescaped # characters followed by spaces/tabs or end of line.
  - Closing sequence: Optional, preceded by spaces/tabs, followed by spaces/tabs only.
  - Indentation: Up to 3 spaces allowed; 4+ spaces forces a code block.

Nuance Or Contradictions
- Backslash escapes are not recognized in code blocks, code spans, autolinks, or raw HTML, but are valid in URLs and link titles.
- Entity and numeric character references cannot replace structural symbols (e.g., * for emphasis, - for list markers).
- Thematic breaks take precedence over setext headings if a line could be interpreted as both.
- ATX headings require at least one space/tab after the opening # characters, unlike many implementations that do not require it.

Candidate Wiki Hints
- Backslash Escaping Rules
- Entity and Numeric Character References
- Thematic Breaks
- ATX Headings
- Block and Inline Precedence
