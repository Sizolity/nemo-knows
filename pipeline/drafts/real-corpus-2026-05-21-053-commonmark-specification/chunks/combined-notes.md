## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
**Heading path:** Document
**Line range:** 1–516
**Scope:** Introduction to the CommonMark Specification (Version 0.31.2), covering metadata, rationale for a formal spec, document structure, and preliminary character definitions.

## Local Summary
This chunk introduces the **CommonMark Specification**, version 0.31.2, authored by John MacFarlane under a Creative Commons BY-SA license. It addresses the lack of unambiguous syntax in John Gruber’s original Markdown description, which led to divergent implementations across platforms like GitHub and Reddit. The spec aims to resolve ambiguities regarding lists, code blocks, headings, and inline structure through side-by-side Markdown/HTML examples that serve as conformance tests.

## Key Claims
- **Readability Goal:** Markdown is designed so that a formatted document is readable as plain text without appearing marked up with tags or instructions.
- **Ambiguity in Original Syntax:** Gruber’s original description does not specify rules for sublist indentation, blank lines before block quotes/headings, code block requirements, list item wrapping, right-aligned markers, thematic breaks within lists, marker precedence, section headings inside lists, empty list items, link reference scope, or definition precedence.
- **Implementation Divergence:** Without a formal spec, implementations consulted buggy scripts (Markdown.pl) or made arbitrary choices, leading to documents rendering differently on different systems without triggering syntax errors.
- **HTML as Test Representation:** The spec uses HTML for side-by-side examples because it represents structural distinctions well; however, not every HTML feature in the examples is mandated by the spec (e.g., percent-encoding of non-ASCII URLs).

## Entities And Concepts
- **CommonMark**: A specification for Markdown syntax to ensure interoperability.
- **Markdown.pl**: The original Perl script by John Gruber used as a reference but deemed buggy and insufficient for defining unambiguous syntax.
- **Conformance Tests**: Examples in the spec that double as tests against implementations using `spec_tests.py`.
- **Abstract Syntax Tree**: The internal representation Markdown is parsed into; HTML samples approximate this structure.
- **Unicode Code Points**: Defined as characters for the purpose of the spec, including combining accents.
- **Line Endings**: Defined as line feed (U+000A), carriage return not followed by line feed, or both.

## Procedures And API Details
**Parsing Strategy Overview:**
The document outlines a two-phase parsing strategy:
1.  **Phase 1: Block Structure**: Identifying container blocks and leaf blocks.
2.  **Phase 2: Inline Structure**: Processing inlines within identified blocks.

**Test Execution Command:**
```bash
python test/spec_tests.py --spec spec.txt --program PROGRAM
```

**Tooling:**
- `tools/makespec.py`: Converts the source text file (`spec.txt`) into HTML or CommonMark.

## Nuance Or Contradictions
- **Tab Handling**: Tabs are not expanded to spaces generally but behave as if replaced by four spaces in contexts defining block structure (e.g., indented code blocks, list item continuation). Internal tabs within content are passed through literally.
- **HTML Mandates vs. Examples**: The spec provides HTML examples where the destination URL might contain non-ASCII characters not percent-encoded. While the spec defines what counts as a link destination, it does not mandate encoding; implementers using automatic tests must provide a renderer that conforms to these expectations (encoding non-ASCII), but conforming implementations may choose different renderers.
- **Original vs. Spec**: The spec explicitly contrasts its unambiguous rules with Gruber’s original description, noting that the latter allows interpretations (like sublist indentation) that contradict common assumptions or specific implementations like Markdown.pl.

## Candidate Wiki Hints
- **CommonMark Specification**: A canonical reference for Markdown parsers to ensure cross-platform consistency.
- **Markdown Conformance Testing**: Methodology for validating Markdown implementations against a standard spec using `spec_tests.py`.
- **Tab Behavior in Markdown**: Rules defining how tabs function in block structure contexts versus inline content.

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
This chunk details the CommonMark specification rules for backslash escaping, entity and numeric character references, and the structure of block-level elements. It covers thematic breaks (horizontal rules), ATX headings, and precedence rules between block and inline structures.

Local Summary
The document explains that backslashes escape only specific characters (like `*`, `#`) in certain contexts but not others (like code blocks). It defines valid HTML entity references and numeric character references (`&#` or `&#x`), noting their restrictions in structural elements like lists and headings. The chunk also introduces the concept of block structure (paragraphs, lists, headings) versus inline content, establishing that block indicators take precedence over inline ones. Specific rules for thematic breaks (`***`, `---`) and ATX headings (`# foo`) are provided, including requirements for spacing, indentation limits (3 spaces), and character matching.

Key Claims
- Backslash escapes do not work in code blocks, code spans, autolinks, or raw HTML.
- Valid HTML entity references and numeric character references can be used in place of Unicode characters, except in code contexts and structural elements like emphasis delimiters or list markers.
- Entity/numeric references cannot replace special characters defining structure (e.g., `&#42;` cannot create an emphasis bullet).
- Indicators of block structure always take precedence over indicators of inline structure.
- Thematic breaks require 3+ matching `-`, `_`, or `*` with optional indentation (up to 3 spaces).
- ATX headings use 1–6 unescaped `#` characters; more than six is not a heading.
- ATX headings must have at least one space between the opening `#` and content unless empty.

Entities And Concepts
- **Backslash Escaping**: Context-dependent escaping mechanism.
- **Entity References**: HTML5 named entities (e.g., `&copy;`).
- **Numeric Character References**: Decimal (`&#35;`) or hexadecimal (`&#x22;`).
- **Thematic Breaks**: Horizontal rules defined by sequences of `-`, `_`, or `*`.
- **ATX Headings**: Headings defined by leading/trailing `#` characters.
- **Block Structure**: Structural elements like paragraphs, lists, and headings.
- **Inline Content**: Text, links, emphasized text within blocks.

Procedures And API Details
- **Thematic Break Validation**: Line must contain 3+ matching characters (`-`, `_`, `*`) with optional spaces/tabs; no other characters allowed.
- **ATX Heading Parsing**:
  - Opening sequence: 1–6 unescaped `#`.
  - Closing sequence: Optional, any number of unescaped `#`.
  - Spacing: Spaces or tabs required after opening and before closing `#`.
  - Indentation: Up to 3 spaces allowed; 4+ spaces treats it as a code block.
- **Entity Parsing**: Check against HTML5 entity list; invalid codes replaced by U+FFFD.

Nuance Or Contradictions
- Original ATX implementation required a space after `#`, but the spec allows headings without it if strictly following the grammar, though many implementations still require it to avoid ambiguity (e.g., `#5 bolt`).
- HTML5 accepts entities without semicolons (e.g., `&copy`), but CommonMark requires the semicolon (`&copy;`) to avoid grammar ambiguity.

Candidate Wiki Hints
- **Thematic Breaks**: Rules for horizontal rules in Markdown.
- **ATX Headings**: Syntax and parsing rules for `#` based headings.
- **Backslash Escaping**: Limitations and valid contexts for backslashes.
- **Entity References**: Usage of HTML entities in Markdown text.

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
- Source Section: **4.3 Setext headings** and **4.4 Indented code blocks**.
- Line Range: 1161–1978.
- Scope: Defines syntax for Setext-style headers (underline-based), rules for multiline content, indentation limits, precedence over other block types, compatibility notes regarding existing implementations, and the transition to indented code blocks.

Local Summary
This chunk details the mechanics of **Setext headings**, which rely on underlines (`=` or `-`) rather than the `#` syntax used in ATX headings. It clarifies that heading content can span multiple lines (unlike many legacy Markdown parsers), provided a blank line separates preceding paragraphs. The text also introduces **Indented code blocks**, defined as sequences of non-blank lines preceded by at least four spaces, and notes their precedence over list items when ambiguity exists.

Key Claims
- A Setext heading consists of one or more lines of text followed by an underline (`=` for level 1, `-` for level 2).
- The underline must not exceed three spaces of indentation; four spaces terminates the block as code or paragraph content.
- Multiline headings are supported if the text lines do not form other block constructs (e.g., lists, thematic breaks) before the underline.
- A blank line is required between a preceding paragraph and a following Setext heading to prevent the paragraph from becoming part of the heading's content.
- Indented code blocks require four or more spaces of indentation per line and treat content as literal text (no Markdown parsing).
- List item interpretations take precedence over indented code block interpretations if ambiguity exists regarding indentation.

Entities And Concepts
- **Setext Heading**: A header style using underlines (`=`, `-`).
- **Indented Code Block**: A code block formed by lines indented four or more spaces.
- **ATX Heading**: Implicitly referenced as the alternative heading syntax (using `#`).
- **Thematic Break**: A horizontal rule defined by three or more dashes, which can interrupt a potential Setext heading if not properly delimited.

Procedures And API Details
- **Heading Level Determination**: Check underline character (`=` -> Level 1, `-` -> Level 2).
- **Indentation Rule**:
    - Content lines: Up to 3 spaces indentation allowed.
    - Underlines: Up to 3 spaces indentation allowed; trailing spaces/tabs ignored for the underline itself but not for content.
    - Code blocks: Minimum 4 spaces indentation required.
- **Closing Condition**: A line with fewer than 4 spaces of indentation ends an indented code block immediately.

Nuance Or Contradictions
- **Multiline Compatibility**: The specification allows multiline headings (e.g., "Foo\nbar\n---\nbaz"), but notes that most existing Markdown implementations do not support this. These legacy tools often interpret such text as separate paragraphs or thematic breaks.
- **Lazy Continuation**: Setext underlines cannot function as lazy continuation lines within list items or block quotes; they must appear on their own structural level.
- **Internal Spaces in Underlines**: An underline cannot contain internal spaces or tabs (e.g., `= =` is invalid for a heading underline).

Candidate Wiki Hints
- **Topic: Setext Headings** – A page explaining the specific syntax of underlined headers, including multiline examples and indentation constraints.
- **Topic: Indented Code Blocks** – Documentation on how four-space indentation defines code blocks and their interaction with list items.

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
This chunk covers the `baz` heading and details CommonMark rules for code blocks (info strings, tilde vs backtick fences) and HTML blocks. It defines seven types of HTML block start/end conditions, illustrates how they interrupt or contain paragraphs, and contrasts these rules with John Gruber’s original Markdown syntax regarding blank lines and indentation. The section concludes by introducing link reference definitions.

Local Summary
The specification clarifies that info strings in backtick code fences cannot contain backticks, whereas tilde fences can. HTML blocks are defined by seven distinct start/end patterns (e.g., `<pre>...</pre>`, comments `<!-- -->`, or generic tags). Unlike Gruber’s original syntax which required blank lines around block-level HTML and forbade indentation, this spec allows indented HTML blocks to interrupt paragraphs (except for type 7) and treats content inside tags as raw HTML until a matching end tag or blank line terminates the block.

Key Claims
- Info strings for backtick code fences cannot contain backticks; tilde fences allow them.
- There are seven kinds of HTML block defined by specific start/end conditions.
- HTML blocks continue until an appropriate end condition, the document end, or a container boundary.
- HTML blocks of types 1–6 may interrupt a paragraph without preceding blank lines (except at document end).
- Blocks of type 7 cannot interrupt a paragraph.
- This spec disallows blank lines inside HTML blocks to avoid expensive balanced tag parsing and allow Markdown content insertion via blank lines.
- Link reference definitions consist of a label, colon, destination, and optional title; they do not correspond to structural elements.

Entities And Concepts
- Code block info strings (language specification)
- Backtick code fences (` ``` `)
- Tilde code fences (` ~~~ `)
- HTML blocks (types 1–7)
- Block-level HTML tags (e.g., `<div>`, `<table>`)
- Inline HTML tags (e.g., `<del>`)
- Link reference definitions
- Raw HTML vs. escaped HTML

Procedures And API Details
- **HTML Block Detection**:
  - Type 1: Starts with `<pre`, `<script`, `<style`, or `<textarea` followed by space/tab/`>`/EOL; ends at matching `</pre>`, etc.
  - Type 2: Starts with `<!--`; ends at `-->`.
  - Type 3: Starts with `<?`; ends at `?>`.
  - Type 4: Starts with `<!` followed by ASCII letter; ends at `>`.
  - Type 5: Starts with `<![CDATA[`; ends at `]]>`.
  - Type 6: Starts with block-level tags (e.g., `<div>`, `<table>`); ends at blank line.
  - Type 7: Starts with any complete open/closing tag; ends at blank line.
- **Link Reference Definition Syntax**:
  - Label + `:` + Destination + Optional Title.
  - Matching is case-insensitive.
  - Definitions can precede or follow usage.

Nuance Or Contradictions
- Gruber’s original Markdown requires blank lines around block-level HTML and forbids indentation; this spec relaxes indentation rules but removes blank line flexibility inside blocks to simplify parsing.
- In type 1–6 HTML blocks, a partial tag on the first line is valid if split where whitespace would occur (e.g., `<div id="foo"`).

Candidate Wiki Hints
- Page: CommonMark HTML Blocks
  - Explain the seven types of HTML block start/end conditions.
  - Discuss how HTML blocks interact with paragraphs and code blocks.
  - Compare with Gruber’s original Markdown restrictions.

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
## Chunk Context
This chunk covers link reference definitions (Examples 215–218), paragraph formation rules (Section 4.8, Examples 219–227), and the handling of blank lines (Section 4.9).

## Local Summary
Link reference definitions can follow each other without blank lines and may appear inside block containers like lists or quotations; they apply globally to the document. Paragraphs are sequences of non-blank lines interpreted as inlines, where initial/final spaces are removed, multiple blank lines between paragraphs are ignored, and indentation rules distinguish between inline text and indented code blocks. Blank lines generally do not affect paragraph boundaries but determine list tightness.

## Key Claims
- Link reference definitions can occur consecutively without intervening blank lines.
- Definitions inside block containers (e.g., quotations) affect the entire document, not just the container.
- A paragraph is a sequence of non-blank lines that cannot be interpreted as other block types.
- Paragraph raw content is formed by concatenating lines and removing initial/final spaces or tabs.
- Multiple blank lines between paragraphs have no effect on paragraph boundaries.
- Lines after the first in a paragraph may be indented any amount; only the first line allows up to three spaces of indentation before being treated as code.
- Final spaces or tabs are stripped before inline parsing, so trailing spaces do not create hard line breaks within a paragraph.

## Entities And Concepts
- Link reference definitions
- Paragraphs
- Block containers (lists, quotations)
- Indented code blocks
- Hard line breaks (via two or more final spaces—though stripped in paragraphs)

## Procedures And API Details
1. **Link Reference Definitions**: Place `[id]: url` lines; multiple can follow consecutively. They may appear inside block containers.
2. **Paragraph Formation**:
   - Collect consecutive non-blank lines.
   - Concatenate lines and remove initial/final spaces or tabs.
   - Interpret the result as inlines.
3. **Indentation Rules for Paragraphs**:
   - First line: up to 3 leading spaces allowed; 4+ spaces start an indented code block.
   - Subsequent lines: any indentation allowed.
4. **Blank Lines**:
   - Ignored between block-level elements except for list tightness/looseness.
   - Ignored at document start/end.

## Nuance Or Contradictions
- Trailing spaces in a paragraph are stripped before inline parsing, so they do not produce hard line breaks within paragraphs (unlike in some other contexts).
- Blank lines between paragraphs are ignored; only the sequence of non-blank lines defines paragraph boundaries.

## Candidate Wiki Hints
- **Link Reference Definitions**: Global definitions can be placed anywhere; order and placement matter for resolution.
- **Paragraph Structure**: Spaces/tabs handling and indentation rules are critical for distinguishing text from code.
- **Blank Line Semantics**: Blank lines are mostly ignored except in list formatting contexts.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 5: Container blocks**, specifically defining the syntax for block quotes and list items. It details how these structures are built recursively from their contents, including rules for markers, indentation, laziness, consecutiveness, and specific behaviors regarding code blocks, thematic breaks, and nested structures.

# Local Summary

The specification defines container blocks as those containing other blocks (block quotes and list items). Block quotes are introduced via a `>` marker with specific rules for indentation, spacing, and "laziness" (omitting markers on continuation lines). List items are defined by bullet (`-`, `+`, `*`) or ordered markers (1–9 digits), with complex rules governing indentation requirements relative to the list marker width. The chunk includes numerous examples illustrating how blank lines separate block quotes, how lazy continuation lines work within nested structures, and the specific indentation math required to keep content inside a list item versus outside it.

# Key Claims

- Container blocks are defined recursively by transforming a sequence of blocks into a container of type Y.
- A block quote marker is `>` followed by a space, or just `>`. Up to three spaces of indentation may precede the marker; four spaces creates a code block instead.
- List markers for bullets are `-`, `+`, or `*`. Ordered list markers use 1–9 digits followed by `.` or `)`.
- The "Laziness" rule allows omitting the `>` marker on lines containing paragraph continuation text, provided the indentation requirements relative to containing blocks are met.
- List item indentation is relative: content must be indented sufficiently past the list marker and any preceding containers (like blockquotes).
- Ordered list start numbers are limited to 9 digits due to browser integer overflow concerns; leading zeros are allowed.
- A list item cannot interrupt a paragraph unless it starts with specific conditions, and thematic breaks terminate list items.

# Entities And Concepts

- **Container blocks**: Blocks that hold other blocks as contents.
- **Block quotes**: Defined by `>` markers, supporting laziness and nesting.
- **List items**: Defined by bullet or ordered markers, containing blocks separated by blank lines.
- **Lazy continuation lines**: Lines within a block quote or list item where the opening marker (`>`) is omitted if it would be redundant for paragraph text.
- **Indented code blocks**: Code blocks inside list items requiring specific indentation relative to the list marker.

# Procedures And API Details

**Block Quote Construction:**
1. Prepend `>` (optionally with 1–3 spaces of indentation) to lines.
2. Apply "Laziness": Remove initial `>` from continuation lines if they are paragraph text.
3. Ensure blank lines separate distinct block quotes ("Consecutiveness").

**List Item Construction:**
1. **Basic Case**: Prepend list marker `M` (width `W`) and `N` spaces (`1 ≤ N ≤ 4`) to the first line. Indent subsequent lines by `W + N`.
2. **Indented Code Start**: If starting with indented code, prepend `M` and exactly one space; indent subsequent lines by `W + 1`.
3. **Blank Line Start**: Prepend `M` to a blank line; indent subsequent lines by `W + 1`.
4. **Indentation Rule**: Lines may be uniformly indented by up to three spaces without changing the list item status.

**Validation Rules:**
- Ordered markers must use digits 0–9 (max 9 digits).
- Negative numbers (e.g., `-1.`) are invalid markers.
- Thematic breaks (`---`) end a list item.

# Nuance Or Contradictions

- **Laziness Limitation**: The `>` marker cannot be omitted on lines that start a new block type (like code blocks or thematic breaks) within the quote, even if indentation suggests it might continue.
- **Indentation Columns vs. Relative Indentation**: While one might assume content must align in a specific column, the rule is strictly about relative indentation past the last containing block marker. Content can appear far to the right and still be inside the list item if indented enough relative to the container edge, or far to the left and outside if not.
- **Empty List Items**: An empty list item (marker followed by blank line) is valid but cannot interrupt a paragraph; it must be part of a list context.

# Candidate Wiki Hints

- **Page: Commonmark Block Quotes** – Covers syntax, laziness, nesting, and separation rules.
- **Page: Commonmark List Items** – Covers marker types, indentation math, code block inclusion, and empty items.

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
The document discusses the specification of list items, specifically focusing on indentation rules for sublists, block elements within lists, and the distinction between "tight" and "loose" lists. It contrasts the proposed CommonMark rules with John Gruber's original Markdown spec and Markdown.pl behavior. The text includes examples demonstrating how lists interrupt paragraphs, handle nested structures, and manage list markers of varying widths.

## Local Summary
This section defines how to determine if content belongs inside a list item versus starting a new one based on indentation relative to the list marker. It establishes the "four-space rule" as a principle for indented block elements (paragraphs, code blocks) within list items but introduces a more flexible strategy where indentation is measured from the end of the list marker. The text details how sublists must be indented sufficiently to nest, how lists can interrupt paragraphs (with restrictions on ordered numerals starting with 1), and how to separate consecutive lists using HTML comments or blank lines.

## Key Claims
- Sublists must be indented by a number of spaces equal to what a paragraph would need to be included in the parent list item.
- The "four-space rule" suggests block-level content under a list item requires four spaces of indentation, though this is arbitrary and potentially unintuitive for beginners.
- A common strategy allows the width and indentation of the list marker to determine the necessary indentation for blocks falling under the list item, rather than a fixed margin count.
- Lists may interrupt paragraphs in CommonMark, provided specific conditions regarding ordered list markers (starting with 1) are met to avoid spurious captures from hard-wrapped text.
- Changing the bullet or ordered list delimiter (e.g., from `-` to `+`, or `.` to `)` ) starts a new, separate list.

## Entities And Concepts
- **List Item**: A sequence of content marked by a list marker (`-`, `*`, `+`, `1.`, etc.).
- **Sublist**: A nested list within a list item; requires specific indentation relative to the parent list marker.
- **Tight vs. Loose List**: Tight lists have no blank lines between items or internal block elements; loose lists contain blank lines, resulting in wrapped `<p>` tags in HTML output.
- **Four-Space Rule**: An inference that all block elements under a list item must be indented four spaces from the margin.
- **Indentation Strategy**: Measuring indentation relative to the end of the list marker to accommodate variable marker widths (e.g., `10)` vs `-`).

## Procedures And API Details
- **Sublist Indentation**: Calculate required indentation by adding the width of the parent list marker and any initial indentation to the standard block indentation amount.
  - *Example*: A bullet list item with no extra indent requires two spaces for a sublist; an ordered list like `10)` requires four spaces total.
- **Separating Lists**: Insert a blank HTML comment (`<!-- -->`) between consecutive lists of the same type or to separate a list from an indented code block that would otherwise be parsed as a subparagraph.
- **List Interruption**: A paragraph can be followed immediately by a list without a blank line. Ordered lists interrupting paragraphs are restricted to those starting with `1` to prevent parsing text like "is 14." as a list.

## Nuance Or Contradictions
- **Markdown.pl vs. CommonMark**: Markdown.pl allowed only two spaces of indentation for sublists and exhibited inconsistent behavior (requiring three spaces for nested sublists). CommonMark adopts a more forgiving approach to handle both the strict four-space rule and Markdown.pl's legacy formatting, provided the layout is natural for humans.
- **Fixed vs. Relative Indentation**: A fixed indent from the margin (like the four-space rule) can lead to unintuitive results where text indented less than the list marker is excluded. A relative indent from the marker itself is proposed as superior but requires handling cases where the list item starts with indented code.
- **Hard-Wrapped Numerals**: While allowing lists to interrupt paragraphs solves natural usage patterns, it risks capturing unintended lists (e.g., "The number of windows in my house is 14."). The spec mitigates this by restricting interruption to ordered lists starting with `1`.

## Candidate Wiki Hints
- List Indentation Rules in CommonMark
- Tight vs. Loose Lists
- Handling Nested Sublists

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
This chunk covers the CommonMark specification section detailing list indentation rules. It specifically addresses how list items are interpreted based on the number of leading spaces (distinguishing between paragraph continuations and code blocks) and introduces examples involving ordered lists with blank lines.

Local Summary
The text clarifies that list items cannot be preceded by more than three spaces of indentation; exceeding this limit changes the interpretation of subsequent text. It provides an example where a fifth item 'e' is treated as a paragraph continuation within the previous item due to excessive indentation. Conversely, it shows that four spaces of indentation following a blank line creates an indented code block.

Key Claims
- List items may not be preceded by more than three spaces of indentation.
- Indenting more than three spaces causes the text to be treated as a paragraph continuation line within the list item context (in the specific example provided).
- Indenting four spaces and preceding with a blank line creates an indented code block.

Entities And Concepts
- List items
- Paragraph continuation line
- Indented code block
- CommonMark specification

Procedures And API Details
None.

Nuance Or Contradictions
The text presents specific examples where indentation thresholds (3 spaces vs 4 spaces) drastically alter the parsing of list structures, shifting content from being part of a list item to either a continuation or a code block.

Candidate Wiki Hints
- CommonMark List Indentation Rules
- Parsing Indented Code Blocks in Lists

## chunk-09

---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
Section 3. c (Loose Lists) and Section 6 (Inlines, Code Spans, Emphasis). The chunk details the parsing logic for loose versus tight lists based on blank lines and block-level content within items, followed by the syntax rules for code spans and emphasis/strong emphasis delimiters.

Local Summary
This section defines how Markdown distinguishes between "loose" and "tight" list items based on internal structure (blank lines, block elements) rather than just line breaks. It then transitions to inline parsing, specifically defining code spans using backticks and the complex algorithm for interpreting asterisks (*) and underscores (_) as emphasis or strong emphasis markers.

Key Claims
- A list is loose if a blank line exists between items, or if an item contains two block-level elements separated by a blank line (even without a physical blank line in the raw text).
- Code spans are delimited by backtick strings; contents are normalized by converting line endings to spaces and stripping single leading/trailing spaces if present on both sides.
- Emphasis rules rely on "delimiter runs" that must be left-flanking or right-flanking depending on surrounding whitespace and punctuation.
- Backslash escapes do not work inside code spans; backslashes are treated literally.
- Inline code spans have higher precedence than emphasis markers, links, images, and HTML tags (except autolinks).

Entities And Concepts
- Loose List: A list where items are separated by blank lines or contain internal block-level separation.
- Tight List: A list where items do not contain internal blank lines separating block elements.
- Delimiter Run: A sequence of `*` or `_` characters not preceded/followed by a non-backslash-escaped version of itself.
- Left-flanking delimiter run: Not followed by whitespace, and either not followed by punctuation OR followed by punctuation and preceded by whitespace/punctuation.
- Right-flanking delimiter run: Not preceded by whitespace, and either not preceded by punctuation OR preceded by punctuation and followed by whitespace/punctuation.
- Code Span: Text enclosed in matching backtick strings with normalized content.

Procedures And API Details
- Code Span Parsing:
  1. Identify a backtick string (one or more backticks) not adjacent to another backtick.
  2. Match the opening and closing backtick strings of equal length.
  3. Normalize contents: convert line endings to spaces.
  4. Strip single leading/trailing space if the normalized string starts and ends with a space but is not all spaces.
- Emphasis Parsing Rules (Simplified):
  - Opening `*` requires a left-flanking delimiter run.
  - Closing `*` requires a right-flanking delimiter run.
  - Opening `_` requires left-flanking, closing `_` requires right-flanking (with specific punctuation exceptions).
  - Strong emphasis (`**`, `__`) follows similar flanking rules but with stricter spacing requirements for `__`.

Nuance Or Contradictions
- Intraword emphasis: Asterisks allow emphasis inside words (e.g., `foo*bar*`), while underscores do not (e.g., `foo_bar_`).
- Flanking logic is counter-intuitive: a delimiter followed by punctuation can act as left-flanking if preceded by whitespace, whereas standard Markdown might expect whitespace to be the primary separator.
- Ambiguity resolution prefers minimizing nesting depth (`<strong>` over `<em><em>`).

Candidate Wiki Hints
- CommonMark Specification: Loose vs Tight Lists
- CommonMark Syntax: Code Spans and Backticks
- Emphasis Algorithms: Delimiter Runs and Flanking Rules

## chunk-10

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

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
Section 3.c details the rules for parsing reference links, including full, collapsed, and shortcut variants. It covers link text grouping precedence over emphasis, HTML/code span/autolink precedence, label normalization (case folding, whitespace collapse), matching logic, and image syntax (6.4). Section 6.5 defines autolinks (URI and email) and Section 6.6 describes raw HTML tag parsing.

Local Summary
The chunk explains that links cannot nest other links at any level. It defines three types of reference links (full, collapsed, shortcut) with specific rules for label matching, normalization, and whitespace handling. Precedence rules determine whether text is treated as link text or inline content. Image syntax mirrors link syntax but allows embedded links in descriptions (though only plain text is used for `alt`). Autolinks are defined by `< >` wrappers around absolute URIs or email addresses. Raw HTML tags between `< >` are passed through without escaping.

Key Claims
- Links may not contain other links, at any level of nesting.
- Reference link labels must be normalized (case fold, strip/normalize whitespace) before matching.
- Matching is case-insensitive; consecutive internal spaces/tabs/line endings count as one space.
- Spaces/tabs/line endings are not allowed between the link text and link label for full/collapsed references.
- Shortcut reference links cannot be followed by `[]` or another link label.
- Full and collapsed references take precedence over shortcut references; inline links also take precedence.
- Image descriptions start with `!` and may contain links, but only plain string content is used for the `alt` attribute.
- Autolinks require `< >` around an absolute URI (scheme + colon + characters) or email address.
- Raw HTML tags between `< >` are rendered without escaping; tag names follow specific ASCII rules.

Entities And Concepts
- Reference Link: A link defined by a label matching a reference definition elsewhere in the document.
- Full Reference Link: `text[label]` where `label` matches a definition.
- Collapsed Reference Link: `text[]` where `text` matches a definition (equivalent to `text[text]`).
- Shortcut Reference Link: `text` where `text` matches a definition (not followed by `[]` or another label).
- Autolink: An absolute URI or email address wrapped in `< >`.
- Raw HTML Tag: Text between `< >` parsed as HTML without escaping.

Procedures And API Details
1. **Link Parsing Precedence**:
   - Link text grouping > emphasis grouping.
   - HTML tags, code spans, autolinks > link grouping.
2. **Label Normalization**:
   - Strip opening/closing brackets.
   - Perform Unicode case fold.
   - Strip leading/trailing whitespace (spaces, tabs, line endings).
   - Collapse consecutive internal whitespace to a single space.
3. **Matching Logic**:
   - Compare normalized forms of label and definition.
   - Use first matching definition if multiple exist.
4. **Image Syntax**:
   - `![description](url "title")` or `![description][]`.
   - Description may contain links, but `alt` uses plain text only.

Nuance Or Contradictions
- The spec departs from John Gruber’s original Markdown by disallowing whitespace between link text and label in reference links to prevent unintended shortcut link capture.
- Image descriptions allow nested links, but the rendered `alt` attribute ignores formatting (e.g., `<a>` tags inside description).

Candidate Wiki Hints
- Reference Link Types (Full, Collapsed, Shortcut)
- Link Label Normalization Rules
- Precedence of Inline Elements Over Links
- Autolink Syntax and Validation

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
Section 3.c defines the syntax of HTML tags, attributes, and specific constructs (comments, processing instructions, declarations, CDATA sections). It provides a series of examples illustrating valid open/closing tags, empty elements, whitespace handling, illegal tag/attribute names/values, and special character preservation. Section 6.7 introduces hard line breaks within HTML contexts.

## Local Summary
This chunk details the structural components of an HTML tag in CommonMark, distinguishing between open tags (containing a name, optional attributes, and optional closing slash) and closing tags. It enumerates valid content types allowed inside tags (comments, processing instructions, declarations, CDATA). A series of examples demonstrates parsing behavior for valid tags with attributes, illegal tag names (which are not parsed as HTML but escaped), illegal attribute syntax, and whitespace rules. The chunk also covers hard line breaks within HTML contexts, noting that they are preserved in attribute values or rendered as `<br />` outside code spans or tags.

## Key Claims
- A double-quoted attribute value consists of a `"`, zero or more characters not including `"`, and a final `"`.
- An open tag includes `<`, a tag name, optional attributes/spaces/line endings, an optional `/`, and `>`.
- Illegal tag names (e.g., starting with a digit or containing invalid characters) are not parsed as HTML tags but escaped.
- Backslash escapes do not work in HTML attributes; they are preserved as literal backslashes.
- Hard line breaks (two spaces + newline or `\` + newline) inside HTML attribute values result in the literal string including the break, unlike outside code spans where they become `<br />`.

## Entities And Concepts
- Open tag
- Closing tag
- Empty element
- HTML comment
- Processing instruction
- Declaration
- CDATA section
- Hard line break
- Attribute value (double-quoted)
- Illegal tag name
- Backslash escape

## Procedures And API Details
N/A (Conceptual parsing rules).

## Nuance Or Contradictions
Hard line breaks behave differently inside code spans (preserved literally) versus outside them or within HTML attribute values (preserved literally in attributes, rendered as `<br />` in the block context if not inside a tag). Backslash escapes function in code spans but are ignored in HTML attributes.

## Candidate Wiki Hints
- `hard-line-breaks-in-html-attributes`

## chunk-13

---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
### Chunk Context
This chunk covers **Section 3.c** (specifically `foo\` backslash escapes) and transitions into **Section 6.8** (Soft line breaks), **6.9** (Textual content), and the **Appendix: A parsing strategy**. It details the two-phase parsing model of CommonMark, the construction of block trees via lazy continuations, and the specific algorithm for handling nested emphasis and links using a delimiter stack.

### Local Summary
The document defines how trailing backslashes create soft line breaks or escape characters, followed by rules for plain text preservation. The core focus shifts to the implementation strategy: a two-phase process where Phase 1 builds a block-level tree structure (paragraphs, lists, quotes) handling indentation and markers, while Phase 2 parses inline content (links, emphasis). A specific stack-based algorithm is detailed for resolving nested delimiters like `*` and `[`.

### Key Claims
- **Soft Line Breaks**: A regular line ending not preceded by two or more spaces or a backslash is parsed as a soft break. Renderers may output this as a space or a line ending; browsers treat them identically.
- **Textual Content**: Characters not matched by specific rules (links, emphasis, code spans) are parsed as plain text, preserving internal spaces verbatim.
- **Two-Phase Parsing**:
  - *Phase 1*: Consumes lines to build block structure (document -> blocks -> children). Text is assigned but not fully parsed. Link definitions are mapped here.
  - *Phase 2*: Parses raw text of paragraphs/headings into inline elements using the link map from Phase 1.
- **Lazy Continuation**: Lines that do not start new blocks (like `>`) are added to the current open block if they satisfy its condition, even if the block is technically "closed" by a subsequent structural change on the same logical line.

### Entities And Concepts
- **Soft Break**: A line ending rendered as a space or newline.
- **Textual Content**: Uninterpreted characters preserved verbatim.
- **Block Tree**: The hierarchical representation of the document (document -> block_quote -> paragraph).
- **Lazy Continuation**: Text added to an open block that precedes a new block start on the same input line.
- **Delimiter Stack**: A doubly linked list used in inline parsing to track potential link/emphasis openers (`[`, `*`, `_`) and closers.
- **Inline Elements**: Strings, code spans, links, emphasis.

### Procedures And API Details
**Phase 1: Block Structure Procedure**
For each line processed:
1. Iterate through open blocks (root down to deepest) to check if the line satisfies their conditions (e.g., block quote needs `>`).
2. Consume continuation markers; look for new block starts. If found, close unmatched blocks and create the new block as a child of the last matched container.
3. Incorporate remaining text into the last open block.

**Phase 2: Inline Structure Procedure**
1. Close all open blocks.
2. Walk the tree visiting every node.
3. Parse raw string contents of paragraphs/headings using the link reference map.

**Delimiter Stack Algorithm (for nested emphasis/links)**
- **Insertion**: When hitting `*`, `_`, `[`, or `![`, insert a text node with literal content and add a pointer to the delimiter stack.
- **Stack Element Data**: Contains pointer to text node, delimiter type (`[`, `![, *, _`), count of delimiters, active status, and potential role (opener/closer).
- **Look for Link/Image**:
  - Start at top of stack, look back for opening `[` or `![`.
  - If inactive, remove and return literal `]`.
  - If active, parse ahead for link/image types. If found, create node, run process emphasis on children, and remove opening delimiter.
- **Process Emphasis**:
  - Iterate until potential closers are exhausted.
  - Move forward to find first potential closer (`*` or `_`).
  - Look back above `stack_bottom` for matching opener.
  - If found: Determine strength (length >= 2), insert emph/strong node, remove intermediate delimiters, update text nodes.
  - If not found: Update bounds and advance current position.

### Nuance Or Contradictions
- **Rendering Variance**: While the spec mandates a soft break is parsed as a line ending or space, it explicitly notes that in browsers these render the same, implying renderer flexibility for the visual output of soft breaks versus hard breaks.
- **Block Closure Timing**: Blocks are not strictly closed immediately upon encountering a new block start; instead, unmatched blocks from the previous step are closed *before* creating the new block, allowing for complex nesting logic within a single line.

### Candidate Wiki Hints
- **Page: CommonMark Parsing Strategy** (Concept: Two-phase parsing model)
- **Page: Delimiter Stack Algorithm** (Concept: Handling nested emphasis and links)
- **Page: Soft Line Breaks** (Concept: Rendering differences between soft and hard breaks)

