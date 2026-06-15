## group-01

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

## group-02

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

## group-03

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

This group of notes aggregates content from the **CommonMark Specification**, covering the document structure, metadata retrieval, and the core parsing logic for lists, text content, and inline elements. The collection spans from the initial document header through detailed algorithmic descriptions of the delimiter stack, soft line breaks, and the two-phase parsing strategy (block structure followed by inline structure). The content ranges from line 1 to line 8133 of the source file, encompassing sections on list continuation, nested items, textual content preservation, and the specific mechanics of resolving emphasis and links.

# Cross-Chunk Summary

The document begins with high-level metadata and structural definitions, including fetch metadata and retrieved text handling. It transitions into specific examples involving list items, nested lists, and continuation rules (Chunks 01–02). A significant portion of the text (Chunks 03–04) focuses on "baz" and "Heading" paths, likely representing specific test cases or edge cases in the specification. The middle sections (Chunks 05–07) cover bracketed text `[Foo]` and generic "aaa" paths, suggesting coverage of link/image syntax and general text nodes. The latter half of the document (Chunks 08–13) delves deeply into list item enumeration (2. b, 3. c) and concludes with the sophisticated logic of soft line breaks, textual content preservation, and the delimiter stack algorithm used to parse inline elements like emphasis and links.

# Repeated Or Central Claims

- **Two-Phase Parsing**: The specification consistently emphasizes a two-phase approach: first constructing the block structure (paragraphs, block quotes, lists) and then parsing raw text contents into inline elements (strings, code spans, links, emphasis).
- **Textual Content Preservation**: Any characters not given an interpretation by previous rules are parsed as plain textual content, and internal spaces are preserved verbatim.
- **Delimiter Stack Mechanism**: The resolution of nested emphasis and links relies on a doubly linked list (delimiter stack) that tracks potential openers and closers for delimiters such as `*`, `_`, `[`, and `![]`.
- **Rendering Flexibility**: Soft line breaks (regular line endings not preceded by spaces or backslashes) allow renderers to choose between rendering them as a line ending or a space, as browser behavior renders both identically.
- **List Continuation**: The specification details how list items can be continued with open blocks or nested without requiring explicit item continuation markers in certain contexts.

# Important Local Details

- **Soft Break Definition**: A line ending is parsed as a soft break if it is not preceded by two or more spaces or a backslash.
- **Delimiter Stack Operations**:
  - Insertion: When encountering `*`, `_`, `[`, or `![`, a text node is inserted, and a pointer is added to the delimiter stack.
  - Link/Image Resolution: When encountering `]`, the system calls `look for link or image`, searching backwards for an active opening delimiter. If a link is found, preceding `[` delimiters are set to inactive.
  - Emphasis Processing: The `process emphasis` function iterates through the stack to find matching openers and closers, determining emphasis levels based on delimiter run length and inserting `emph` or `strong emph` nodes.
- **Openers_bottom**: This concept defines a lower bound for searching delimiters of specific types during emphasis processing.
- **Lazy Continuation**: Describes a line added to an open block without closing it, relevant to list item handling.
- **Chunk Line Ranges**: The notes cover a continuous range from line 1 to 8133, with specific focus on the final chunk (7802-8133) detailing the delimiter stack and soft breaks.

# Candidate Wiki Hints

- **Soft Line Breaks**: Create a page explaining the conditions for soft breaks and the rendering variations (space vs. line ending).
- **Parsing Strategy**: Develop a technical guide detailing the two-phase parsing model and the tree construction process.
- **Delimiter Stack Algorithm**: Write a deep dive into the logic for resolving nested emphasis and links, including the `look for link or image` and `process emphasis` procedures.
- **List Continuation Rules**: Document the rules for list item continuation, nested lists, and open blocks.
- **Textual Content**: A page on how plain text and internal spaces are handled when no other interpretation applies.

# Gaps Or Cautions

- **Missing Raw Content**: The provided chunk notes and index do not contain the actual text of the specification, only metadata about headings and line ranges. Specific examples or test cases referenced in the "baz" or "aaa" sections cannot be fully verified without the raw text.
- **Ambiguous Heading Paths**: Several chunks (e.g., Chunk 02, 03) have heading paths that appear to be test case identifiers (e.g., "foo > foo > foo") rather than semantic section titles, which may limit the ability to create intuitive navigation links without the full context of the test cases.
- **Incomplete Algorithmic Steps**: While the delimiter stack algorithm is described, the specific implementation details of the `look for link or image` search logic (e.g., exact distance limits, handling of inactive delimiters) are implied but not fully detailed in the provided notes.
- **Rendering Options**: The notes mention that renderers may offer options to treat soft line breaks as hard line breaks, but the specific API or configuration options for this are not detailed.

