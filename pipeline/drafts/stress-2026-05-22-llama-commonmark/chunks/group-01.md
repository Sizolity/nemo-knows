---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

**Source Document:** CommonMark Specification (version 0.31.2) by John MacFarlane.
**Line Range:** 1–2982 (Chapters 1 through 4.7).
**Coverage:**
- Document introduction, metadata, and design goals.
- Preliminaries: characters, lines, tabs, and insecure characters.
- Backslash escapes and entity/numeric character references.
- Block structure: thematic breaks, ATX headings, setext headings, indented/fenced code blocks.
- HTML blocks and link reference definitions.
- Paragraphs and blank lines.
- Container blocks: block quotes and list items.

# Cross-Chunk Summary

The document begins by establishing the necessity of an unambiguous specification for Markdown to resolve historical ambiguities in John Gruber's original description. It defines the fundamental units of the language (characters, lines, tabs) and the mechanism for escaping special characters.

The specification then moves to block-level structures. It distinguishes between container blocks (which can contain other blocks, like block quotes and lists) and leaf blocks (like paragraphs and headings). Specific rules are provided for:
- **Thematic breaks** (horizontal rules).
- **ATX headings** (using `#` symbols).
- **Setext headings** (using underlines).
- **Code blocks** (indented or fenced).
- **HTML blocks** (raw HTML handling).
- **Link reference definitions** (defining labels for links).
- **Paragraphs** (the default block type).

Finally, the text introduces container blocks in detail, defining block quotes via `>` markers and list items via bullet or ordered markers, including rules for "laziness" (omitting markers on continuation lines) and indentation handling.

# Repeated Or Central Claims

- **Readability:** Markdown is designed so that source text is readable and publishable as plain text.
- **Unambiguity:** The specification aims to resolve ambiguities in the original Markdown syntax to ensure consistent rendering across implementations.
- **Escaping:** Backslashes escape ASCII punctuation characters in most contexts but are invalid within code blocks, code spans, autolinks, and raw HTML.
- **Tabs:** Tabs are not globally expanded but act as four spaces in contexts defining block structure (e.g., indented code blocks, list continuations).
- **Blank Lines:** Blank lines generally separate blocks but have specific nuances regarding list tightness/looseness and HTML block termination.
- **Indentation Thresholds:** Four spaces of indentation typically trigger a code block, while three spaces or fewer allow for paragraph continuation or list item content.
- **HTML Blocks:** HTML blocks are treated as raw text; types 1–6 end at matching end tags, while type 7 ends at a blank line.

# Important Local Details

- **Insecure Characters:** U+0000 is replaced with U+FFFD.
- **Entity References:** Valid HTML5 entity names (e.g., `&amp;`) and numeric character references (decimal/hex) are parsed, except in code contexts.
- **ATX Heading Syntax:** Requires 1–6 unescaped `#` characters. At least one space/tab is required after the opening `#` unless the heading is empty.
- **Setext Heading Syntax:** Level 1 uses `=`, Level 2 uses `-`. Multi-line setext headings are allowed in CommonMark but not in most existing implementations.
- **Fenced Code Blocks:** Opened with 3+ backticks or tildes. Closing fences must match the opening character count. Info strings (language tags) are allowed.
- **Link Reference Definitions:** Consist of a label, colon, destination, and optional title. They can appear consecutively and affect the entire document scope.
- **Block Quote Laziness:** Markers (`>`) can be omitted on lines where the content is clearly a continuation of the previous paragraph.
- **List Marker Laziness:** Similar to block quotes, markers can be omitted on continuation lines within list items.
- **Ordered List Limits:** Ordered list markers are limited to 9 digits to prevent browser integer overflow issues.

# Candidate Wiki Hints

- **Page: CommonMark Specification Overview** – Introduction, design goals, and comparison with AsciiDoc.
- **Page: Escaping and Entities** – Rules for backslashes, HTML entities, and numeric character references.
- **Page: Block Structure** – Thematic breaks, ATX headings, and setext headings.
- **Page: Code Blocks** – Indented and fenced code block syntax and rules.
- **Page: HTML Blocks** – Handling raw HTML, block types, and termination conditions.
- **Page: Link Reference Definitions** – Syntax and scoping of link definitions.
- **Page: Paragraphs and Blank Lines** – Formation, indentation rules, and whitespace handling.
- **Page: Container Blocks** – Block quotes and list items, including laziness and consecutiveness.

# Gaps Or Cautions

- **Compatibility Gaps:** CommonMark allows multi-line setext headings, which creates a divergence from most existing Markdown implementations.
- **Indentation Ambiguity:** Lines with 4+ spaces are code blocks, but if they could be part of a list item, the list item interpretation takes precedence.
- **HTML Block Termination:** Unlike Gruber's original spec, CommonMark HTML blocks (types 1–6) do not require blank lines to terminate; they end at the matching end tag.
- **Link Definition Scope:** Link reference definitions inside block containers (like lists) affect the entire document, not just the container, which may differ from intuitive expectations.
- **Hard Line Breaks:** Two or more spaces at the end of a line are stripped before inline parsing, meaning hard line breaks are not preserved as expected in some other Markdown flavors.
