---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Group Context
- **Source Document**: CommonMark Specification (`raw/web/corpus-2026-05-18/053-commonmark-specification.md`).
- **Coverage Range**: Lines 4230–8133 (Chunks 7 through 13).
- **Primary Topics**: List structure and indentation rules, loose vs. tight lists, inline emphasis and strong emphasis parsing, link and image syntax (including reference links and autolinks), and raw HTML tag parsing.
- **Chunk Breakdown**:
  - **Chunk 7**: List item indentation, sublists, four-space rule, and loose/tight list definitions.
  - **Chunk 8**: Indentation limits for list items (3 spaces) vs. code blocks (4 spaces).
  - **Chunk 9**: Loose/tight list criteria, code span parsing, and emphasis delimiter runs.
  - **Chunk 10**: Detailed emphasis rules (nesting, matching, precedence over links).
  - **Chunk 11**: Reference links, image syntax, autolinks, and normalization rules.
  - **Chunk 12**: HTML tag syntax, comments, CDATA, and hard line break rules.
  - **Chunk 13**: Nested list items and continuation lines within complex structures.

Cross-Chunk Summary
This section of the CommonMark specification defines the rigorous parsing rules for block-level lists and inline-level text formatting. It establishes a strict "four-space rule" for indentation within lists to distinguish between list continuations, paragraph continuations, and indented code blocks. The specification distinguishes between "loose" lists (separated by blank lines or containing internal block breaks) and "tight" lists. The inline parsing logic prioritizes code spans, links, and HTML tags over emphasis markers, ensuring that structural elements are not misinterpreted as formatting. Specific rules govern the matching of emphasis delimiters (`*`, `_`, `**`, `__`), the normalization of link labels, and the strict syntax requirements for raw HTML tags, including handling of illegal characters and hard line breaks.

Repeated Or Central Claims
- **Indentation Hierarchy**: Content under a list item must be indented four spaces from the margin (or four spaces from the list marker). Indentation of three spaces or less continues the paragraph; four spaces after a blank line creates a code block.
- **List Tightness**: A list is "loose" if items are separated by blank lines or if an item contains two block-level elements separated by a blank line. Otherwise, it is "tight."
- **Emphasis Precedence**: Inline elements like code spans, links, and HTML tags interrupt or take precedence over emphasis markers. Links bind more tightly than brackets, which bind more tightly than emphasis.
- **Delimiter Matching**: Emphasis requires matching delimiter runs. Intraword emphasis (e.g., `**foo**bar`) is forbidden for underscores but allowed for asterisks under specific conditions.
- **Link Normalization**: Link labels are normalized via Unicode case folding and whitespace collapsing before matching.
- **HTML Tag Integrity**: Raw HTML tags are parsed as-is, but illegal tag names, attributes, or values result in escaped output rather than parsed tags. Hard line breaks are converted to `<br />` tags unless inside code spans or HTML tags.

Important Local Details
- **Four-Space Rule**: Block-level content (paragraphs, sublists, code) under a list item must be indented four spaces from the margin. This prevents ambiguity between list items and indented code.
- **Sublist Indentation**: A sublist must be indented the same number of spaces a paragraph would need to be included in the parent list item.
- **Code Span Normalization**: Leading and trailing spaces are stripped from code spans only if the string begins and ends with a space. Line endings are converted to spaces.
- **Emphasis Rules**:
  - Delimiter runs must be left-flanking (to open) and right-flanking (to close).
  - Punctuation preceding a delimiter affects its flanking status.
  - Nested emphasis requires different delimiters (e.g., `*foo**bar*`) or specific length matching.
- **Reference Links**: Spaces, tabs, or line endings are not allowed between the link text and the link label.
- **Image Alt Text**: Only the plain string content of the image description is used for the `alt` attribute; inline formatting is ignored.
- **Hard Line Breaks**: Two or more spaces or a backslash before a line ending creates a `<br />` tag, except inside code spans, HTML tags, or at the end of a block.
- **Illegal HTML**: Tag names starting with digits or underscores, and attribute names starting with digits or underscores, are illegal and escaped.

Candidate Wiki Hints
- **List Indentation Rules**: Guide on the four-space rule and distinguishing list continuations from code blocks.
- **Loose vs. Tight Lists**: Visual examples and definitions for determining list tightness.
- **Inline Element Precedence**: Hierarchy of parsing for emphasis, links, code, and HTML.
- **Emphasis Delimiter Logic**: How to determine left/right-flanking status and handle nesting.
- **Link Label Normalization**: Steps for case folding and whitespace collapsing in reference links.
- **HTML Tag Parsing**: Rules for valid tags, comments, CDATA, and handling illegal syntax.
- **Hard Line Breaks**: When and where line breaks are converted to `<br />` tags.

Gaps Or Cautions
- **Autolink Validity**: The spec accepts strings like `a+b+c:d` as autolinks, which may not be valid URIs under strict RFC standards.
- **Escaping in Autolinks**: Backslash escapes do not function inside autolinks; they are treated as literal characters.
- **Whitespace Sensitivity**: The spec is highly sensitive to whitespace placement, particularly between list markers and content, and between link text and labels.
- **Unicode Whitespace**: Only ASCII spaces are stripped from code spans; general Unicode whitespace is not handled the same way.
- **Escaped Characters**: Backslash escapes do not work in HTML attributes or inside autolinks.
