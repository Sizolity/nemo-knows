## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Chunk Context

**Heading path:** Document
**Line range:** 1–516
**Coverage:**
- Document overview and metadata
- Introduction to Markdown and the need for a specification
- Preliminaries: characters, lines, tabs, and insecure characters
- Backslash escapes and entity references

# Local Summary

This chunk introduces the CommonMark Specification (version 0.31.2) by John MacFarlane. It explains Markdown’s design goal of readability, contrasts Markdown with AsciiDoc, and justifies the need for an unambiguous spec due to historical ambiguities in John Gruber’s original description. The text covers character definitions, line endings, tab handling, and backslash escape rules.

# Key Claims

- Markdown is a plain text format for structured documents, prioritizing readability so that source text is publishable as-is.
- The original Markdown syntax description by John Gruber is ambiguous in several cases (e.g., list indentation, blank line requirements, code block rules).
- Without an unambiguous spec, implementations diverge, leading to inconsistent rendering across platforms.
- Tabs are not expanded to spaces by default but behave as four spaces in contexts defining block structure.
- Backslashes before ASCII punctuation characters escape their Markdown meaning; backslashes before other characters are literal.

# Entities And Concepts

- **CommonMark**: A specification for a Markdown-compatible markup language.
- **Markdown**: A lightweight markup language developed by John Gruber and Aaron Swartz.
- **AsciiDoc**: A markup language used for comparison to illustrate Markdown’s readability.
- **Backslash escapes**: Mechanism to treat punctuation characters literally.
- **Unicode code point**: The unit of character used in the spec.
- **Line ending**: LF, CR, or CRLF.
- **Blank line**: A line with no characters or only spaces/tabs.
- **Insecure characters**: Specifically U+0000, replaced with U+FFFD.

# Procedures And API Details

- **Tab handling**: Tabs are treated as four spaces in block structure contexts (e.g., indented code blocks, list item continuation).
- **Escape sequence**: `\` + ASCII punctuation character → literal character.
- **Test runner**: `python test/spec_tests.py --spec spec.txt --program PROGRAM` can run conformance tests.
- **Spec generation**: `tools/makespec.py` converts `spec.txt` (Markdown with test extensions) to HTML or CommonMark.

# Nuance Or Contradictions

- **Tabs vs. spaces**: Tabs are not expanded globally but act as four spaces in structural contexts. Internal tabs in code blocks remain literal.
- **Blank lines**: Some implementations require blank lines before block quotes or headings; the spec aims to clarify this.
- **List markers**: Ambiguities exist regarding indentation of sublists, right-aligned markers, and marker changes (numbers to bullets).
- **HTML rendering**: The spec uses HTML in examples but does not mandate all HTML features (e.g., percent-encoding of non-ASCII URLs).

# Candidate Wiki Hints

- **Page: CommonMark Specification** – Overview of the spec, its history, and design goals.
- **Page: Markdown vs. AsciiDoc** – Comparison of readability and syntax.
- **Page: Backslash Escapes** – Rules for escaping punctuation and literal backslashes.
- **Page: Tab Handling in Markdown** – How tabs interact with block structure.
- **Page: CommonMark Conformance Tests** – Using `spec_tests.py` for validation.

## chunk-02

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

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context

This chunk covers sections 4.3 (Setext headings), 4.4 (Indented code blocks), and the beginning of 4.5 (Fenced code blocks) from the CommonMark specification. It details the syntax, parsing rules, and examples for these block-level elements, including interactions with paragraphs, lists, and block quotes.

## Local Summary

The text defines **setext headings** (using `=` or `-` underlines), **indented code blocks** (using 4+ spaces of indentation), and **fenced code blocks** (using backticks or tildes). It explains how to distinguish headings from thematic breaks or paragraphs, how indentation affects code block content, and the specific rules for opening and closing code fences.

## Key Claims

- **Setext Headings**: Consist of text lines followed by an underline (`=` for level 1, `-` for level 2). They cannot interrupt a paragraph; a blank line is required before a following heading if it follows a paragraph.
- **Indented Code Blocks**: Composed of non-blank lines preceded by four or more spaces. Content is literal text (no Markdown parsing). Ambiguity with list items favors the list item interpretation.
- **Fenced Code Blocks**: Begin with a code fence (3+ backticks or tildes). Content is literal. Closing fences must match the opening character and have at least as many characters. They can interrupt paragraphs.
- **Compatibility Note**: Most existing Markdown implementations do not support multi-line setext headings, though CommonMark does.

## Entities And Concepts

- **Setext Heading**: A heading defined by an underline (`=` or `-`) rather than `#` symbols.
- **Indented Code Block**: A code block defined by indentation (4+ spaces).
- **Fenced Code Block**: A code block defined by backticks or tildes.
- **Code Fence**: The line starting a fenced code block (3+ backticks/tildes).
- **Info String**: Optional text following the opening code fence, typically indicating the language.
- **Thematic Break**: A horizontal rule (`---` or `***`), which can be confused with setext headings if not separated by a blank line.

## Procedures And API Details

- **Setext Heading Parsing**:
  1. Check for text lines followed by an underline (`=` or `-`).
  2. Ensure the underline has no more than 3 spaces of indentation.
  3. Ensure the text lines are not interpretable as other block constructs (code fence, ATX heading, block quote, etc.).
  4. If the underline follows a paragraph, a blank line is required.
- **Indented Code Block Parsing**:
  1. Identify lines with 4+ spaces of indentation.
  2. Treat content as literal text (remove the 4 spaces).
  3. If a line has fewer than 4 spaces, the code block ends.
- **Fenced Code Block Parsing**:
  1. Identify a code fence (3+ backticks/tildes) with up to 3 spaces of indentation.
  2. Parse content until a closing fence of the same type and length (or longer) is found.
  3. Remove indentation from content lines if the opening fence was indented.

## Nuance Or Contradictions

- **Multi-line Headings**: CommonMark allows multi-line setext headings, but most existing Markdown implementations do not. This creates a compatibility gap.
- **Indentation Ambiguity**: If a line with 4+ spaces could be part of a list item, the list item interpretation takes precedence over a code block.
- **Closing Fence Indentation**: Closing fences can be indented up to 3 spaces and do not need to match the opening fence's indentation.
- **Internal Spaces in Fences**: Code fences cannot contain internal spaces or tabs; otherwise, they are not treated as fences.

## Candidate Wiki Hints

- **Setext Headings**: A reusable concept for defining headings with underlines.
- **Indented Code Blocks**: A distinct method for defining code blocks via indentation.
- **Fenced Code Blocks**: A flexible method for defining code blocks with fences.
- **Code Fence Syntax**: Rules for opening and closing code fences.
- **Info String**: Optional metadata for code blocks.

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
- **Heading**: `baz`
- **Line Range**: 1979–2982
- **Coverage**: Section 4.6 (HTML blocks) and the beginning of Section 4.7 (Link reference definitions).

Local Summary
This chunk details the CommonMark specification for handling HTML blocks. It defines seven types of HTML blocks based on start and end conditions, explaining how they interrupt or continue paragraphs. It covers specific behaviors for tags like `<pre>`, `<script>`, and comments, noting that certain blocks end at matching end tags rather than blank lines. The text also contrasts these rules with John Gruber's original Markdown syntax, highlighting differences regarding indentation and blank lines. The section concludes by introducing link reference definitions, outlining their structure (label, colon, destination, optional title) and parsing rules.

Key Claims
- HTML blocks are treated as raw HTML and are not escaped in output.
- There are seven kinds of HTML blocks defined by specific start and end conditions.
- Blocks of type 1–6 (e.g., `<pre>`, `<script>`, comments) end at the first line containing a corresponding end tag, allowing blank lines inside them.
- Blocks of type 7 (generic tags) end at the first blank line following the block.
- HTML blocks of types 1–6 can interrupt a paragraph; type 7 blocks cannot.
- Link reference definitions consist of a label, a colon, a destination, and an optional title.
- Link reference definitions do not correspond to structural elements but define labels for reference links.
- Matching of link labels is case-insensitive.
- A link reference definition cannot interrupt a paragraph.

Entities And Concepts
- **HTML Block**: A group of lines treated as raw HTML.
- **Start/End Conditions**: Rules determining when an HTML block begins and ends.
- **Type 1–6 Blocks**: Specific HTML constructs (pre, script, style, textarea, comments, processing instructions, declarations) ending at matching tags.
- **Type 7 Blocks**: Generic HTML blocks ending at a blank line.
- **Link Reference Definition**: A structural element defining a label for reference links.
- **Link Label**: The text used to identify a link.
- **Link Destination**: The URL or anchor target.
- **Link Title**: Optional text describing the link.

Procedures And API Details
- **HTML Block Parsing**:
  1. Check if a line meets a start condition (e.g., begins with `<pre`, `<!--`, `<?`).
  2. If matched, treat subsequent lines as raw HTML until the matching end condition is met.
  3. For types 1–5, look for the corresponding end tag (e.g., `</pre>`, `-->`).
  4. For type 6, look for a blank line.
  5. For type 7, look for a blank line.
- **Link Reference Definition Parsing**:
  1. Identify a line starting with a label followed by a colon (`:`).
  2. Parse the destination and optional title.
  3. Ensure no further characters occur after the title.
  4. Store the definition for use in reference links later in the document.

Nuance Or Contradictions
- **Indentation**: HTML blocks of types 1–6 can be preceded by up to three spaces of indentation, but not four. Type 7 blocks also allow indentation up to three spaces.
- **Blank Lines**: Unlike Gruber's original specification, CommonMark does not require blank lines before HTML blocks of types 1–6, nor does it allow blank lines inside type 7 blocks.
- **Tag Matching**: End tags need not match the start tag exactly (case-insensitive), but the content between them is treated as raw HTML.
- **Garbage In, Garbage Out**: Partial or invalid tags are passed through as-is if they start like a valid tag.

Candidate Wiki Hints
- **HTML Blocks in Markdown**: Explaining how to embed raw HTML safely and the differences between block types.
- **Link Reference Definitions**: A guide to defining custom links using labels and destinations.
- **CommonMark vs. Original Markdown**: Comparing the handling of HTML blocks and indentation rules.

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
## Chunk Context
This chunk covers link reference definitions (Examples 215–218) and the definition of paragraphs (Section 4.8) and blank lines (Section 4.9) within the CommonMark specification.

## Local Summary
The text explains that link reference definitions can appear consecutively without blank lines and can be nested inside block containers like quotations, affecting the entire document scope. It then defines a paragraph as a sequence of non-blank lines that cannot be interpreted as other blocks, detailing how raw content is processed (concatenation, space removal) and how indentation rules apply to distinguish paragraphs from code blocks. Finally, it notes that blank lines are generally ignored except for determining list tightness.

## Key Claims
- Link reference definitions can occur one after another without intervening blank lines.
- Link reference definitions inside block containers (e.g., lists, block quotations) affect the entire document, not just the container.
- A paragraph is formed by a sequence of non-blank lines that cannot be interpreted as other kinds of blocks.
- Paragraph raw content is formed by concatenating lines and removing initial and final spaces or tabs.
- Multiple blank lines between paragraphs have no effect on the output.
- Leading spaces or tabs are skipped; lines after the first may be indented any amount.
- Four spaces of indentation on the first line of a paragraph context creates a code block instead of a paragraph.
- Final spaces or tabs are stripped before inline parsing, preventing hard line breaks if two or more spaces are present.
- Blank lines between block-level elements are ignored except for determining list tightness.

## Entities And Concepts
- Link reference definitions
- Block containers (lists, block quotations)
- Paragraphs
- Raw content
- Indented code blocks
- Blank lines
- List tightness/looseness

## Procedures And API Details
- **Paragraph Formation**: Concatenate lines, remove initial/final spaces/tabs.
- **Indentation Rule**: First line may have up to three spaces; four spaces triggers a code block.
- **Hard Line Breaks**: Two or more spaces at the end of a line are stripped before inline parsing.

## Nuance Or Contradictions
- While blank lines are ignored for paragraph separation, they are critical for determining whether a list is tight or loose.
- Indentation rules allow subsequent lines in a paragraph to be indented arbitrarily, but the first line is restricted to three spaces of indentation.

## Candidate Wiki Hints
- **Link Reference Definitions**: Scope and placement rules.
- **Paragraph Parsing**: Indentation thresholds and whitespace handling.
- **Blank Line Semantics**: Role in list structure vs. general block separation.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context

This chunk covers Section 5 of the CommonMark specification, focusing on **Container Blocks**. It details the syntax and parsing rules for **Block Quotes** (5.1) and **List Items** (5.2). The text defines how these containers are constructed recursively from their contents, including specific rules for markers, indentation, laziness, and consecutiveness.

## Local Summary

The document defines container blocks as blocks containing other blocks, specifically block quotes and list items. It establishes recursive definitions for syntax rather than a direct parsing recipe. Section 5.1 outlines block quote markers (`>`), rules for basic cases, laziness (omitting markers on continuation lines), and consecutiveness (requiring blank lines between separate quotes). Section 5.2 defines list items via bullet (`-`, `+`, `*`) and ordered (digits followed by `.` or `)`) markers. It details four rules for list items: basic case, starting with indented code, starting with a blank line, and indentation handling. Laziness for list items is also described.

## Key Claims

- Container blocks are defined recursively based on their contents.
- Block quotes require a `>` marker followed by a space or no space, preceded by up to three spaces of indentation.
- List items are defined by markers (`-`, `+`, `*`, `1-9.`) followed by 1–4 spaces of indentation.
- Ordered list markers are limited to 9 digits to avoid integer overflows in browsers.
- Laziness allows omitting block quote markers on lines where the content would be paragraph continuation text.
- List items can contain any kind of block, including indented code blocks and block quotes.
- Indented code blocks within list items require specific indentation relative to the list marker.

## Entities And Concepts

- **Container Blocks**: Blocks that contain other blocks (block quotes, list items).
- **Block Quotes**: Defined by `>` markers; support laziness and nesting.
- **List Items**: Defined by bullet or ordered markers; support various content types.
- **Laziness**: Rule allowing omission of markers on continuation lines.
- **Consecutiveness**: Requirement for blank lines between separate block quotes.
- **Indented Code Blocks**: Code blocks indented by four spaces; treated specially within lists.
- **Thematic Breaks**: Lines that terminate list items.

## Procedures And API Details

- **Block Quote Marker**: `>` followed by a space or no space, with up to three spaces of indentation.
- **List Marker**: `-`, `+`, `*` for bullets; `1-9.` or `1-9)` for ordered lists.
- **Indentation Rules**:
  - List items: Marker width + 1–4 spaces.
  - Indented code blocks: Four spaces beyond the list item edge.
- **Laziness Application**: Remove markers from lines where the next non-space/tab character is paragraph continuation text.

## Nuance Or Contradictions

- **Blank Line Separation**: Block quotes must be separated by blank lines, unlike some Markdown implementations that merge them.
- **Lazy Continuation**: Markers can be omitted only on lines that would be paragraph continuations; they cannot be omitted before thematic breaks, code blocks, or list items.
- **Indentation Columns**: Indentation is relative to the list marker, not strictly column-based; content can be in the same column as the marker but still be inside the list if indented sufficiently past the last containing block.
- **Empty List Items**: Allowed at the start or end of a list but cannot interrupt a paragraph.

## Candidate Wiki Hints

- **Page: CommonMark Container Blocks**
  - Summary: Overview of block quotes and list items as container blocks.
- **Page: CommonMark Block Quotes**
  - Summary: Syntax, laziness, and nesting rules for block quotes.
- **Page: CommonMark List Items**
  - Summary: Marker types, indentation, and content rules for lists.

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Heading: aaa
- Lines: 4230–4754
- Source: raw/web/corpus-2026-05-18/053-commonmark-specification.md

Local Summary
This chunk details the CommonMark specification for list items, sublists, and list structure. It contrasts the "four-space rule" (indented blocks under a list item) with the historical "Markdown.pl" behavior (two-space indentation). It explains how list markers determine the required indentation for subsequent content, how lists interrupt paragraphs, and how to separate consecutive lists using blank HTML comments.

Key Claims
- A sublist must be indented the same number of spaces a paragraph would need to be included in the list item.
- The "four-space rule" requires block-level content (paragraphs, sublists, code) under a list item to be indented four spaces from the margin.
- CommonMark allows lists to interrupt paragraphs (e.g., `Foo\n- bar`), unlike Markdown.pl.
- Only ordered lists starting with `1` are allowed to interrupt paragraphs to avoid spurious list captures (e.g., `14.`).
- Changing the list marker character (e.g., `-` to `+`) or number format starts a new list.
- A list is "loose" if items are separated by blank lines or contain internal blank lines; otherwise, it is "tight."
- Blank HTML comments (`<!-- -->`) can separate consecutive lists or prevent indented code from being parsed as a subparagraph.

Entities And Concepts
- List Item: A block containing list markers and content.
- Sublist: A nested list within a list item.
- Four-Space Rule: The principle that content under a list item must be indented four spaces.
- Tight List: A list where items are not separated by blank lines.
- Loose List: A list where items are separated by blank lines.
- List Marker: The symbol (`-`, `+`, `*`, `1.`, etc.) starting a list item.
- HTML Comment: Used to separate lists (`<!-- -->`).

Procedures And API Details
- Indentation for Sublists: Match the indentation required for a paragraph to be included in the list item (determined by list marker width).
- Indentation for Code: Eight spaces from the margin (or four spaces from the list marker).
- Indentation for Blockquotes: Indented (typically four spaces).
- Separating Lists: Insert a blank HTML comment between lists to ensure they are treated as distinct.

Nuance Or Contradictions
- Markdown.pl allowed two-space indentation for sublists but was inconsistent (requiring three spaces for sub-sublists).
- The four-space rule is arbitrary and unintuitive for beginners but provides a principled standard.
- A two-space rule from the list marker would allow text indented less than the marker to be included, which is unintuitive.
- Indented code in a list item must be indented eight spaces from the margin to avoid breaking existing Markdown patterns.

Candidate Wiki Hints
- List Item Indentation Rules
- Four-Space Rule vs. Markdown.pl
- Loose vs. Tight Lists
- Separating Lists with HTML Comments
- Lists Interrupting Paragraphs

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Heading path: 2. b
- Heading coverage: 2. b, 3. c
- Line range: 4756-4798
- Source file: raw/web/corpus-2026-05-18/053-commonmark-specification.md

Local Summary
This chunk illustrates the indentation rules for ordered and unordered list items in CommonMark. It demonstrates that list items cannot be preceded by more than three spaces of indentation. If indented by more than three spaces, a line is treated as a paragraph continuation rather than part of the list item. Additionally, it shows that four spaces of indentation following a blank line creates an indented code block.

Key Claims
- List items may not be preceded by more than three spaces of indentation.
- Indentation exceeding three spaces causes the line to be treated as a paragraph continuation.
- Indentation of four spaces, when preceded by a blank line, creates an indented code block.

Entities And Concepts
- Ordered lists (`<ol>`)
- Unordered lists (`<ul>`)
- List items
- Indented code blocks
- Paragraph continuation lines

Procedures And API Details
- Indentation Rule: Keep indentation at three spaces or less for list item content.
- Code Block Trigger: Use four spaces of indentation after a blank line to start an indented code block.
- Example 312: Shows a list where item 'd' has a continuation line 'e' indented more than three spaces, resulting in 'e' being part of the paragraph within item 'd'.
- Example 313: Shows a list where a subsequent item is indented four spaces after a blank line, forming an indented code block.

Nuance Or Contradictions
- The distinction between three spaces (paragraph continuation) and four spaces (code block) is critical for correct parsing.
- Blank lines preceding indented content determine whether the indentation creates a code block or is ignored/interpreted differently.

Candidate Wiki Hints
- CommonMark Indentation Rules for Lists
- Indented Code Blocks vs. List Continuations
- Parsing List Item Continuation Lines

## chunk-09

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

## chunk-10

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

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Heading path: 3. c
- Line range: 6637-7513
- Section: Reference links, image syntax, and autolinks.

Local Summary
This chunk details the parsing rules for reference links (full, collapsed, shortcut), image syntax, and autolinks. It establishes precedence rules where link text grouping overrides emphasis, HTML tags, code spans, and autolinks. It defines the normalization process for link labels (case folding, whitespace collapsing) and specifies constraints such as the 999-character limit for link labels and the prohibition of unescaped brackets within labels. The section concludes with the grammar for raw HTML tags.

Key Claims
- Links cannot contain other links at any level of nesting.
- Link text grouping takes precedence over emphasis grouping.
- HTML tags, code spans, and autolinks take precedence over link grouping.
- Link labels must be normalized (case fold, strip whitespace) for matching.
- Matching is case-insensitive.
- Reference links are case-insensitive.
- Spaces, tabs, or line endings are not allowed between the link text and the link label.
- Image descriptions may contain links, but only the plain string content is used for the `alt` attribute.
- Autolinks consist of absolute URIs or email addresses enclosed in `< >`.
- Raw HTML text between `<` and `>` is parsed as tags without escaping.

Entities And Concepts
- Reference links (full, collapsed, shortcut)
- Link label
- Link text
- Image description
- Autolink (URI autolink, email autolink)
- Raw HTML tag
- Normalization (Unicode case fold, whitespace collapsing)
- Scheme (for URIs)

Procedures And API Details
- **Link Label Normalization**: Strip opening/closing brackets, perform Unicode case fold, strip leading/trailing whitespace, collapse internal whitespace to a single space.
- **Link Matching**: Compare normalized forms of the link text and label. Use the first matching definition if multiple exist.
- **Image Alt Attribute**: Extract only the plain string content of the image description; ignore inline formatting.
- **Autolink Parsing**: Match `<` followed by an absolute URI or email address followed by `>`.
- **Raw HTML Grammar**:
  - Tag name: ASCII letter followed by letters, digits, or hyphens.
  - Attribute name: ASCII letter, `_`, or `:` followed by letters, digits, `_`, `.`, `:`, or `-`.
  - Attribute value: Unquoted, single-quoted, or double-quoted.

Nuance Or Contradictions
- **Whitespace in Reference Links**: This spec forbids whitespace between link text and label, departing from John Gruber’s original Markdown syntax which allowed it. This change prevents unintended capture of consecutive shortcut reference links.
- **Image Formatting**: While image descriptions can contain links, the rendered `alt` attribute uses only the plain text, stripping all formatting.
- **Autolink Validity**: Many strings parsed as autolinks (e.g., `a+b+c:d`) are not valid URIs according to strict standards but are accepted by this spec.
- **Escaping in Autolinks**: Backslash-escapes do not work inside autolinks; they are treated as literal characters.

Candidate Wiki Hints
- **Reference Link Precedence**: How full and collapsed references take precedence over shortcut references.
- **Image Alt Text Rules**: Why only plain text is used for `alt` attributes in Markdown images.
- **Autolink Syntax**: Rules for valid URI and email autolinks, including scheme requirements.
- **Raw HTML Parsing**: Grammar for tags and attributes in the CommonMark spec.

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
This chunk (lines 7515-7801) covers section 3.c of the CommonMark specification, detailing the parsing rules for HTML tags, attributes, comments, processing instructions, declarations, and CDATA sections. It includes numerous examples (613–646) demonstrating valid and invalid syntax for open tags, closing tags, whitespace handling, illegal characters, and the specific behavior of hard line breaks within HTML contexts.

Local Summary
The specification defines the structural components of HTML tags, including open tags, closing tags, and various embedded content types like comments and CDATA sections. It provides concrete examples of valid tag syntax, including custom tag names and attributes with various quote styles and values. The text explicitly lists illegal tag names, illegal attribute names, and illegal attribute values, showing how the parser escapes these in the output. A significant portion is dedicated to "Hard line breaks" (section 6.7), explaining how line endings preceded by spaces or backslashes are rendered as `<br />` tags, with specific constraints regarding where these breaks can occur (e.g., not inside code spans or at the end of a block).

Key Claims
- A double-quoted attribute value consists of a starting `"`, zero or more characters not including `"`, and a final `"`.
- An open tag consists of `<`, a tag name, optional attributes/spaces, an optional `/`, and `>`.
- An HTML comment consists of `<!--`, a string not including `-->`, and `-->`.
- A processing instruction consists of `<?`, a string not including `?>`, and `?>`.
- A declaration consists of `<!`, an ASCII letter, zero or more characters not including `>`, and `>`.
- A CDATA section consists of `<![CDATA[`, a string not including `]]>`, and `]]>`.
- An HTML tag is defined as an open tag, a closing tag, an HTML comment, a processing instruction, a declaration, or a CDATA section.
- Hard line breaks (two or more spaces or a backslash before a line ending) are parsed as `<br />` tags, except inside code spans or at the end of a block.
- Backslash escapes do not work in HTML attributes.
- Entity and numeric character references are preserved in HTML attributes.

Entities And Concepts
- Open tag
- Closing tag
- HTML comment
- Processing instruction
- Declaration
- CDATA section
- HTML tag
- Hard line break
- Entity reference
- Numeric character reference
- Attribute value
- Tag name

Procedures And API Details
- To form a valid open tag: Start with `<`, followed by a tag name, then zero or more attributes, optional spaces/tabs/line endings, an optional `/`, and end with `>`.
- To form a valid closing tag: Start with `</`, followed by a tag name, optional spaces/tabs/line endings, and end with `>`.
- To form a hard line break: Place two or more spaces or a backslash before a line ending (not in a code span or HTML tag, and not at the end of a block).
- To form an HTML comment: Use `<!--`, insert content not containing `-->`, and close with `-->`.
- To form a processing instruction: Use `<?`, insert content not containing `?>`, and close with `?>`.
- To form a declaration: Use `<!`, an ASCII letter, insert content not containing `>`, and close with `>`.
- To form a CDATA section: Use `<![CDATA[`, insert content not containing `]]>`, and close with `]]>`.

Nuance Or Contradictions
- The specification notes that illegal tag names (e.g., `<33>`, `<__>`) are not parsed as HTML tags but are escaped in the output.
- Illegal attribute names (e.g., `h*#ref`) and illegal attribute values (e.g., missing closing quote) result in escaped output.
- Illegal whitespace (e.g., space after `<`, space before `>`) causes the tag to be treated as text.
- Missing whitespace between attributes (e.g., `href='bar'title=title`) is considered illegal and results in escaped output.
- Hard line breaks do not occur inside code spans or HTML tags, nor at the end of a paragraph or other block element.
- Backslash escapes are explicitly stated to not work in HTML attributes.

Candidate Wiki Hints
- HTML Tag Syntax in CommonMark
- Hard Line Breaks in Markdown
- HTML Comments and CDATA in CommonMark
- Illegal Characters in HTML Tags

## chunk-13

---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
### Chunk Context
- **Heading path**: 3. c > foo\
- **Coverage**: Sections 6.8 (Soft line breaks), 6.9 (Textual content), and Appendix A (Parsing strategy).
- **Line range**: 7802-8133.

### Local Summary
This chunk details how CommonMark handles soft line breaks, preserves textual content verbatim, and outlines a two-phase parsing strategy (block structure then inline structure). It includes a detailed walkthrough of the delimiter stack algorithm used to resolve nested emphasis and links.

### Key Claims
- A regular line ending not preceded by two or more spaces or a backslash is parsed as a **softbreak**.
- Soft line breaks may be rendered in HTML as either a line ending or a space; the result is the same in browsers.
- Renderers may provide an option to render soft line breaks as hard line breaks.
- Any characters not given an interpretation by previous rules are parsed as plain textual content.
- Internal spaces are preserved verbatim.
- Parsing occurs in two phases: first constructing the block structure, then parsing raw text contents into inline elements.
- The delimiter stack is a doubly linked list used to track potential openers and closers for emphasis and links.

### Entities And Concepts
- **Softbreak**: A line ending parsed when not preceded by spaces or a backslash.
- **Textual content**: Characters parsed as plain text when no other interpretation applies.
- **Delimiter stack**: A doubly linked list tracking delimiters (`*`, `_`, `[`, `![]`) with metadata on activity and potential to open/close.
- **Openers_bottom**: A lower bound for searching delimiters of specific types during emphasis processing.
- **Lazy continuation**: A line added to an open block without closing it.

### Procedures And API Details
- **Parsing Strategy**:
  1. **Phase 1 (Block Structure)**: Lines are consumed to construct the document tree (paragraphs, block quotes, lists). Blocks are closed or new ones created based on line content.
  2. **Phase 2 (Inline Structure)**: Raw text in paragraphs and headings is parsed into inlines (strings, code spans, links, emphasis) using the link reference map from Phase 1.
- **Delimiter Stack Algorithm**:
  - When hitting `*`, `_`, `[`, or `![`, insert a text node and add a pointer to the delimiter stack.
  - When hitting `]`, call `look for link or image`.
  - When hitting end of input, call `process emphasis` with `stack_bottom = NULL`.
  - **look for link or image**: Search backwards for an active opening delimiter. If found and active, parse ahead for link/image types. If a link is found, set preceding `[` delimiters to inactive.
  - **process emphasis**: Iterate through the stack to find matching openers and closers. Determine if emphasis or strong emphasis exists based on delimiter run length. Insert `emph` or `strong emph` nodes and remove intermediate delimiters.

### Nuance Or Contradictions
- **Rendering Flexibility**: While the spec defines a soft line break, it explicitly allows renderers to choose between rendering it as a line ending or a space, noting that browser behavior makes the result identical.
- **Parser Implementation**: Renderers may offer an option to treat soft line breaks as hard line breaks, implying implementation flexibility beyond the strict parsing rules.

### Candidate Wiki Hints
- **Soft Line Breaks**: A dedicated page explaining the conditions for soft breaks and rendering variations.
- **Parsing Strategy**: A technical guide to the two-phase parsing model and tree construction.
- **Delimiter Stack Algorithm**: A deep dive into the logic for resolving nested emphasis and links.

