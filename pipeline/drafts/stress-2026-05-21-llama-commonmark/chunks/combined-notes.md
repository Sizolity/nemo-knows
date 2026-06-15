## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
- **Source**: CommonMark Specification (Version 0.31.2, 2024-01-28)
- **Scope**: Introduction to Markdown, rationale for the specification, and preliminary definitions regarding characters, lines, and tabs.
- **Key Sections**:
  - 1.1 What is Markdown?
  - 1.2 Why is a spec needed?
  - 1.3 About this document
  - 2.1 Characters and lines
  - 2.2 Tabs
  - 2.3 Insecure characters
  - 2.4 Backslash escapes

## Local Summary
This chunk introduces the CommonMark specification, contrasting its readability with the ambiguity of John Gruber's original Markdown description. It explains the necessity of a formal spec to resolve implementation divergences (e.g., list indentation, blank line requirements). The text defines fundamental character classes (whitespace, punctuation, control characters) and establishes rules for handling tabs in block structure contexts.

## Key Claims
- **Readability**: Markdown's primary design goal is readability; documents should be publishable as plain text without looking marked up.
- **Ambiguity of Original Syntax**: John Gruber's canonical description does not unambiguously specify syntax for edge cases (e.g., sublist indentation, blank lines before block quotes, code block indentation).
- **Implementation Divergence**: Without a spec, implementations diverge significantly, causing documents to render differently across systems (e.g., GitHub wiki vs. pandoc).
- **Tab Handling**: Tabs are not expanded to spaces generally but behave as if replaced by spaces with a 4-character tab stop in contexts defining block structure (e.g., indented code blocks).
- **Security**: The Unicode character U+0000 must be replaced with the REPLACEMENT CHARACTER (U+FFFD).

## Entities And Concepts
- **Markdown**: A plain text format for structured documents based on email/usenet conventions.
- **CommonMark**: A specification for Markdown syntax intended to be unambiguous.
- **John Gruber**: Developer of the original Markdown syntax.
- **Markdown.pl**: The original Perl script for converting Markdown to HTML; noted as buggy and insufficient as a spec.
- **Unicode Code Point**: The unit of character definition used in the spec.
- **Line Ending**: Defined as U+000A, U+000D (not followed by U+000A), or U+000D followed by U+000A.
- **Blank Line**: A line with no characters or only spaces/tabs.
- **Backslash Escape**: Mechanism to treat ASCII punctuation characters as literals, stripping their Markdown meaning.

## Procedures And API Details
- **Tab Expansion Rule**: In contexts where spaces define block structure, tabs behave as if replaced by spaces with a tab stop of 4 characters.
- **Backslash Escape Rule**: Any ASCII punctuation character may be backslash-escaped to become a literal character. Backslashes before non-punctuation characters are treated as literal backslashes.
- **Test Execution**: The spec includes a script `spec_tests.py` to run conformance tests against any Markdown program using the `spec.txt` source file.

## Nuance Or Contradictions
- **Tabs vs. Spaces**: While tabs are not expanded globally, they are treated as 4 spaces specifically for block structure definition (e.g., indented code blocks). Internal tabs within code blocks are passed through as literal tabs.
- **Blank Lines**: Most implementations do not require blank lines before block quotes or headings, though the original spec was ambiguous, leading to parsing ambiguities in hard-wrapped text.
- **Original vs. Spec**: The original Markdown description suggests certain behaviors (e.g., two lists when markers change from numbers to bullets), but the CommonMark spec and many implementations produce one list in such cases.

## Candidate Wiki Hints
- **Page**: CommonMark Specification Overview
- **Topic**: Markdown Syntax Ambiguities
- **Topic**: Handling Tabs in Markdown
- **Topic**: Backslash Escapes in Markdown

## chunk-02

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

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
- **Source**: `raw/web/corpus-2026-05-18/053-commonmark-specification.md`
- **Chunk**: 3 of 13
- **Lines**: 1161-1978
- **Heading**: `baz > ` (covers sections 4.3 Setext headings, 4.4 Indented code blocks, and 4.5 Fenced code blocks)

## Local Summary
This chunk details the syntax and parsing rules for three distinct block-level Markdown constructs: **Setext headings**, **Indented code blocks**, and **Fenced code blocks**. It defines the structural requirements for each (e.g., underline characters for setext, four-space indentation for indented code, backtick/tilde sequences for fenced code) and explains how they interact with surrounding text, including rules for blank lines, indentation limits, and precedence over other block types.

## Key Claims
- **Setext Headings**: Defined by text lines followed by an underline of `=` (level 1) or `-` (level 2). They cannot interrupt paragraphs without a preceding blank line. Multiline content is supported, but the text must not be interpretable as other block constructs (like lists or thematic breaks).
- **Indented Code Blocks**: Composed of lines indented by at least four spaces. They are literal text (no Markdown parsing). They cannot interrupt paragraphs.
- **Fenced Code Blocks**: Defined by opening and closing fences of three or more backticks or tildes. They can interrupt paragraphs. The content is literal text, and the info string (language hint) is optional and not strictly mandated for rendering.
- **Precedence**: Block structure indicators (like lists or thematic breaks) take precedence over inline or heading indicators where ambiguity exists.

## Entities And Concepts
- **Setext Heading**: A heading style using underlines (`=` or `-`).
- **Indented Code Block**: A code block defined by 4+ space indentation.
- **Fenced Code Block**: A code block defined by backticks or tildes.
- **Info String**: Optional text following the opening fence in a fenced code block.
- **Code Fence**: The sequence of backticks or tildes delimiting a fenced code block.
- **Thematic Break**: A horizontal rule (`---` or `***`) that can conflict with setext heading underlines.

## Procedures And API Details
- **Setext Heading Parsing**:
  1. Check for text lines followed by an underline (`=` or `-`).
  2. Ensure the underline has no more than 3 spaces of indentation.
  3. Verify the text is not interpretable as a code fence, ATX heading, block quote, thematic break, list item, or HTML block.
  4. If followed by a paragraph, insert a blank line to separate them.
- **Indented Code Block Parsing**:
  1. Identify lines with 4+ spaces of indentation.
  2. Remove exactly 4 spaces of indentation from each line.
  3. Treat content as literal text.
  4. Stop the block if a line has fewer than 4 spaces of indentation.
- **Fenced Code Block Parsing**:
  1. Locate a code fence (3+ backticks or tildes) with up to 3 spaces of indentation.
  2. Extract the info string (trimmed of spaces/tabs).
  3. Collect subsequent lines until a matching closing fence is found.
  4. Remove indentation from content lines if the opening fence was indented.
  5. If no closing fence is found, the block extends to the end of the document/block.

## Nuance Or Contradictions
- **Multiline Setext Headings**: While the spec supports multiline headings (e.g., `Foo\nbar\n---`), most existing Markdown implementations do not. This creates a compatibility gap where users might expect a paragraph followed by a thematic break instead of a multiline heading.
- **Indentation Ambiguity**: In indented code blocks, if a line has more than 4 spaces of indentation, the extra spaces are preserved in the content. However, if a line has fewer than 4 spaces, the code block ends immediately.
- **Fence Matching**: The closing fence must use the same character (backticks or tildes) and be at least as long as the opening fence. Mixing characters or using fewer characters for the closing fence is invalid.
- **Internal Spaces in Fences**: Code fences cannot contain internal spaces or tabs (e.g., `` ``` ``` `` is invalid).

## Candidate Wiki Hints
- **Page: Setext Headings**
  - Summary: Rules for defining level 1 and 2 headings using underlines.
  - Key points: Underline characters (`=` vs `-`), indentation limits, multiline support, and interaction with paragraphs.
- **Page: Indented Code Blocks**
  - Summary: Syntax and behavior of code blocks defined by indentation.
  - Key points: 4-space indentation rule, literal text content, precedence over list items.
- **Page: Fenced Code Blocks**
  - Summary: Syntax and behavior of code blocks defined by backticks or tildes.
  - Key points: Fence length matching, info string handling, interruption of paragraphs.
- **Page: Block Precedence**
  - Summary: How CommonMark resolves conflicts between different block types (e.g., list vs. code block, thematic break vs. heading).

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
- **Heading**: `baz`
- **Line Range**: 1979–2982
- **Source Section**: Covers code block info strings, HTML blocks (types 1–7), and the beginning of link reference definitions.

## Local Summary
This chunk details the handling of info strings in code blocks, distinguishing between backtick and tilde fences regarding allowed characters. It then extensively defines HTML blocks, categorizing them into seven types based on start/end tags (e.g., `<pre>`, comments, CDATA). It explains how HTML blocks interrupt paragraphs, handle indentation, and interact with Markdown syntax. Finally, it introduces link reference definitions, outlining their structure, precedence rules, and constraints regarding indentation and placement.

## Key Claims
- Info strings for backtick code blocks cannot contain backticks; tilde code blocks can.
- HTML blocks are treated as raw HTML and are not escaped in output.
- There are seven specific kinds of HTML blocks defined by start/end conditions.
- HTML blocks of types 1–6 can interrupt a paragraph; type 7 cannot.
- Link reference definitions consist of a label, colon, destination, and optional title.
- Link reference definitions do not correspond to structural elements but define labels for later use.
- Matching of link labels is case-insensitive.

## Entities And Concepts
- **Info String**: Optional text after a code fence indicating language or metadata.
- **HTML Block**: A group of lines treated as raw HTML.
- **Link Reference Definition**: A declaration defining a label for reference links.
- **Block-level HTML Elements**: Elements like `<div>`, `<table>`, `<pre>`.
- **Inline HTML**: Tags not on their own line (e.g., `<del>text</del>`).

## Procedures And API Details
- **HTML Block Start Conditions**:
  1. `<pre`, `<script`, `<style`, or `<textarea` followed by space/tab/>.
  2. `<!--` (comment).
  3. `<?` (processing instruction).
  4. `<!` followed by an ASCII letter (declaration).
  5. `<![CDATA[` (CDATA).
  6. Block-level tags (e.g., `<div>`, `<table>`) followed by space/tab/>.
  7. Any complete open/closing tag (except types 1–4) followed by space/tab/end of line.
- **HTML Block End Conditions**:
  1. Matching end tag (`</pre>`, etc.).
  2. `-->` (comment).
  3. `?>` (processing instruction).
  4. `>` (declaration).
  5. `]]>` (CDATA).
  6. Blank line.
  7. Blank line.
- **Link Reference Definition Syntax**: `[label]: destination "title"` (with optional spaces/tabs/line breaks).

## Nuance Or Contradictions
- **Gruber’s Original Markdown vs. CommonMark**: Gruber’s specification required blank lines before HTML blocks and matching end tags, whereas CommonMark allows HTML blocks to interrupt paragraphs and does not require matching end tags for types 1–6.
- **Blank Lines in HTML Blocks**: CommonMark disallows blank lines inside HTML blocks (except types 1–5) to avoid expensive parsing of balanced tags, whereas Gruber’s rule allowed them.
- **Indentation Sensitivity**: HTML blocks can be preceded by up to three spaces of indentation; four spaces trigger a code block instead.

## Candidate Wiki Hints
- **HTML Blocks in Markdown**: A guide to understanding how HTML blocks are parsed, including the seven types and their interaction with Markdown syntax.
- **Link Reference Definitions**: Rules for defining and using reference links in CommonMark.
- **Code Block Info Strings**: Best practices for specifying language metadata in code blocks.

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
## Chunk Context
This chunk covers link reference definitions (Examples 215–227) and the start of the Paragraphs section (4.8) and Blank lines (4.9) within the CommonMark specification.

## Local Summary
The text details how link reference definitions interact with inline content, specifically regarding placement inside block containers like blockquotes. It establishes that definitions affect the entire document scope, not just the container. The section then transitions to defining paragraphs as sequences of non-blank lines that cannot be interpreted as other blocks, detailing rules for line breaks, indentation, and blank line handling.

## Key Claims
- Link reference definitions can occur consecutively without intervening blank lines.
- Definitions inside block containers (e.g., blockquotes) affect the entire document, not just the container.
- A paragraph is a sequence of non-blank lines that cannot be interpreted as other block types.
- Multiple blank lines between paragraphs have no effect on the output.
- Leading spaces or tabs are skipped when determining paragraph content.
- Indented code blocks cannot interrupt paragraphs, but four spaces of indentation on the first line creates a code block instead.
- Final spaces or tabs are stripped before inline parsing, preventing hard line breaks if two or more spaces are present.

## Entities And Concepts
- Link Reference Definition
- Block Container
- Blockquote
- Paragraph
- Indented Code Block
- Hard Line Break

## Procedures And API Details
- **Paragraph Formation**: Concatenate lines, remove initial/final spaces/tabs.
- **Link Definition Scope**: Definitions are global; placement inside a blockquote does not limit their scope.
- **Indentation Rules**:
  - 0–3 spaces: Skipped (part of paragraph).
  - 4+ spaces: Creates an indented code block (interrupts paragraph).
- **Line Break Handling**:
  - Two or more spaces at end of line: Stripped (no hard break).
  - Blank lines: Ignored except for list tightness/looseness.

## Nuance Or Contradictions
- **Indentation Ambiguity**: While lines after the first in a paragraph may be indented any amount, the first line is restricted to up to three spaces of indentation to remain part of the paragraph.
- **Blank Line Role**: Blank lines are generally ignored between block elements but are critical for determining if a list is "tight" or "loose."

## Candidate Wiki Hints
- **Link Reference Definitions**: A dedicated page explaining global vs. local link definitions and their placement rules.
- **Paragraph Parsing**: A guide on how CommonMark handles line breaks, indentation, and blank lines within paragraphs.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 5: Container blocks**, specifically detailing the syntax and parsing rules for **Block quotes (5.1)** and **List items (5.2)**. It defines how container blocks are constructed recursively from their contents, the specific markers required for block quotes (the `>` character), and the complex indentation and laziness rules governing list items.

# Local Summary

The text defines container blocks as blocks containing other blocks, specifically block quotes and list items. It establishes a recursive definition where a container is formed by transforming a sequence of blocks. Section 5.1 details block quote markers (`>`), allowing up to three spaces of indentation before the marker, and introduces the "Laziness" rule which permits omitting the marker on continuation lines of paragraphs. Section 5.2 defines list items via bullet (`-`, `+`, `*`) or ordered (1–9 digits) markers, explaining how indentation depth relative to the marker determines content inclusion, handling of indented code blocks within lists, and the "Laziness" rule for lists.

# Key Claims

- Container blocks are defined recursively by transforming a sequence of blocks.
- Block quotes require a `>` marker followed by a space or no space, preceded by up to three spaces of indentation.
- Block quotes exhibit "Laziness," allowing the `>` marker to be omitted on lines containing paragraph continuation text.
- List items are defined by markers (`-`, `+`, `*`, or `1`–`9` digits) followed by 1–4 spaces of indentation.
- Indentation depth is relative; content must be indented sufficiently to fall under the list item's edge, not just a fixed column.
- List items can contain any block type, including indented code blocks, block quotes, and thematic breaks (which terminate the list item).
- Ordered list start numbers are limited to 9 digits to avoid integer overflows in some browsers.
- List items can start with a blank line, but only one.
- Indented code blocks within list items require specific indentation relative to the list marker and the code block edge.

# Entities And Concepts

- **Container blocks**: Blocks that have other blocks as contents (block quotes, list items).
- **Block quote marker**: The `>` character.
- **Laziness**: A rule allowing omission of block quote markers on continuation lines of paragraphs.
- **List marker**: Bullet (`-`, `+`, `*`) or ordered (`1`–`9` digits + `.` or `)`) markers.
- **Indentation**: Spaces or tabs used to nest content within containers; relative indentation determines inclusion.
- **Indented code block**: A code block preceded by four spaces of indentation.
- **Thematic break**: A line that terminates a list item.
- **Paragraph continuation text**: Text parsed as part of a paragraph but not at the beginning.

# Procedures And API Details

- **Block Quote Construction**: Prepend `> ` (or just `>`) to lines. Omit `>` on continuation lines of paragraphs.
- **List Item Construction**: Prepend marker (e.g., `- ` or `1. `) to the first line. Indent subsequent lines by the width of the marker plus the spaces following it.
- **Indentation Calculation**: Calculate required indentation as `width of marker + spaces after marker`. Content must be indented at least this much to be included.
- **Ordered List Numbering**: Use 1–9 digits. Numbers starting with 0 are valid (e.g., `0.`). Negative numbers are invalid.
- **Code Block Handling**: Indented code blocks within lists require 4 spaces beyond the list item edge.

# Nuance Or Contradictions

- **Laziness vs. Structure**: The "Laziness" rule applies only to lines that would be paragraph continuations. It does not apply to lines starting new blocks (like lists or code blocks) or thematic breaks.
- **Indentation Columns**: One might assume content must align in a specific column, but the spec relies on relative indentation from the marker. Content can be in the same column as the marker but still be included if indented enough past the containing block's edge.
- **Blank Lines**: Blank lines separate block quotes but are not strictly required between a block quote and a following paragraph unless laziness rules apply. List items can contain multiple blank lines.
- **Empty List Items**: A list item can be empty (just a marker and a blank line), but an empty list item cannot interrupt a paragraph.

# Candidate Wiki Hints

- **Page: CommonMark Container Blocks**
  - Summary: Overview of block quotes and list items as container blocks.
- **Page: CommonMark Block Quote Syntax**
  - Summary: Rules for `>` markers, indentation limits, and the Laziness rule.
- **Page: CommonMark List Item Indentation**
  - Summary: How indentation depth determines content inclusion in list items.
- **Page: CommonMark Ordered List Numbers**
  - Summary: Constraints on ordered list start numbers (1–9 digits, no negatives).

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Section: 5.3 Lists
- Lines: 4230–4754
- Heading: aaa
- Scope: Rules for sublists, indentation logic, loose vs tight lists, interrupting paragraphs, and list separation.

Local Summary
- Defines how sublists are indented relative to list markers.
- Explains the four-space rule vs Markdown.pl’s two-space behavior.
- Introduces loose vs tight lists and HTML output differences.
- Details how lists interrupt paragraphs and how to separate lists.

Key Claims
- Sublists must be indented enough to fit the list marker plus any initial indentation.
- The four-space rule is preferred over fixed indentation from the margin.
- A list is loose if items are separated by blank lines or contain internal blank lines.
- Lists can interrupt paragraphs in CommonMark, unlike Markdown.pl.
- Only lists starting with "1" are allowed to interrupt paragraphs to avoid spurious captures.
- Blank HTML comments can separate consecutive lists or prevent unintended code blocks.

Entities And Concepts
- List marker: Bullet (-, +, *) or ordered (., ))
- Four-space rule: Indentation requirement for blocks under list items.
- Loose list: Items separated by blank lines; paragraphs wrapped in `<p>`.
- Tight list: No internal blank lines; paragraphs not wrapped in `<p>`.
- Principle of uniformity: Text meaning remains consistent inside containers.
- Spurious list capture: Unintended list creation from hard-wrapped numerals.

Procedures And API Details
- Indentation calculation: Measure from the start of the list marker, not the margin.
- Code block indentation: Eight spaces from the margin (or six from the marker in some proposals).
- Separating lists: Insert `<!-- -->` between lists to reset parsing.
- List interruption: No blank line needed between paragraph and list.

Nuance Or Contradictions
- Markdown.pl allowed two-space indentation for sublists inconsistently.
- Different implementations (Pandoc, discount, redcarpet) handled indentation differently.
- Fixed four-space rule feels unnatural for some layouts.
- Two-space rule risks including unintended text in list items.
- Indented code inside lists requires special handling to avoid breaking existing patterns.

Candidate Wiki Hints
- List indentation rules in CommonMark
- Loose vs tight lists
- Lists interrupting paragraphs
- Separating lists with HTML comments
- Spurious list capture prevention

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
This chunk covers sections 2.b and 3.c of the CommonMark specification, focusing on the indentation rules for list items. It demonstrates how list items are parsed when preceded by varying amounts of indentation (0, 1-3 spaces, and 4+ spaces) and how blank lines affect the interpretation of subsequent text.

## Local Summary
The text explains that list items cannot be preceded by more than three spaces of indentation. If a line is indented by more than three spaces, it is treated as a paragraph continuation rather than a new list item. If a line is indented by four or more spaces and preceded by a blank line, it is treated as an indented code block.

## Key Claims
- List items may not be preceded by more than three spaces of indentation.
- Indentation of more than three spaces causes a line to be treated as a paragraph continuation.
- Indentation of four or more spaces, when preceded by a blank line, creates an indented code block.

## Entities And Concepts
- List items
- Indentation
- Paragraph continuation
- Indented code block
- Blank line

## Procedures And API Details
- **Indentation Rule**: Check the number of spaces preceding a list item.
  - 0-3 spaces: Valid list item continuation or new item.
  - >3 spaces: Treated as paragraph continuation.
  - >=4 spaces (with blank line): Treated as indented code block.

## Nuance Or Contradictions
The specification clarifies that the distinction between a list item continuation and a code block depends on both the amount of indentation and the presence of a preceding blank line.

## Candidate Wiki Hints
- **CommonMark Indentation Rules**: A page detailing how indentation affects list parsing in CommonMark.
- **Indented Code Blocks**: A guide on when text is interpreted as code versus paragraph text.

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
- Section 6 introduces inline parsing.
- Section 6.2 details emphasis and strong emphasis rules.
- Lines 4799-5730 cover list examples and inline syntax definitions.

Local Summary
This chunk explains how Commonmark determines whether a list is "loose" or "tight" based on blank lines and block-level content within items. It then transitions to inline parsing, specifically focusing on code spans and the complex rules for emphasis and strong emphasis using asterisks and underscores.

Key Claims
- A list is loose if there is a blank line between items or if an item contains multiple block-level elements separated by blank lines.
- A list is tight if blank lines appear only within code blocks, block quotes, or between paragraphs of a sublist.
- Inline parsing proceeds sequentially from left to right.
- Code spans are delimited by backticks and normalize content by converting line endings to spaces and stripping leading/trailing spaces if present on both sides.
- Emphasis and strong emphasis rely on "delimiter runs" that are classified as left-flanking, right-flanking, or both, based on surrounding whitespace and punctuation.
- Ambiguities in emphasis parsing are resolved by minimizing nesting depth and preferring shorter spans when closing delimiters match.

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
- Unicode whitespace
- Unicode punctuation

Procedures And API Details
- **Code Span Normalization**:
  1. Convert line endings to spaces.
  2. If the string begins and ends with a space (but is not all spaces), remove one space from each end.
- **Emphasis Opening Rules**:
  - `*` opens emphasis if part of a left-flanking delimiter run.
  - `_` opens emphasis if part of a left-flanking delimiter run and either not right-flanking or right-flanking preceded by punctuation.
- **Emphasis Closing Rules**:
  - `*` closes emphasis if part of a right-flanking delimiter run.
  - `_` closes emphasis if part of a right-flanking delimiter run and either not left-flanking or left-flanking followed by punctuation.
- **Ambiguity Resolution**:
  - Minimize nestings (prefer `<strong>` over `<em><em>`).
  - Prefer `<em><strong>` over `<strong><em>`.
  - When spans overlap, the first one takes precedence.
  - When spans share a closing delimiter, the shorter one (opening later) takes precedence.
  - Inline code, links, images, and HTML tags group more tightly than emphasis.

Nuance Or Contradictions
- Intraword emphasis with `_` is generally disallowed to avoid unwanted emphasis in words containing internal underscores, unlike `*`.
- Backslash escapes do not work in code spans; all backslashes are treated literally.
- The stripping of leading/trailing spaces in code spans only occurs if spaces exist on both sides; single spaces are preserved if only one side has them.
- Browsers typically collapse consecutive spaces in `<code>` elements, so CSS `white-space: pre-wrap` is recommended.

Candidate Wiki Hints
- Page: Commonmark List Types (Loose vs Tight)
- Page: Commonmark Inline Syntax (Code Spans)
- Page: Commonmark Emphasis Rules (Delimiter Runs)

## chunk-10

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

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
- Section 3. c: Link syntax and reference link definitions.
- Section 6.4: Images syntax.
- Section 6.5: Autolinks syntax.
- Section 6.6: Raw HTML syntax.

Local Summary
This chunk details the parsing rules for reference links (full, collapsed, shortcut), image syntax, autolinks (URI and email), and raw HTML tags. It clarifies precedence rules between link types, normalization for matching, and specific constraints on content within link labels and descriptions.

Key Claims
- Links cannot contain other links at any nesting level.
- Reference link matching is case-insensitive and uses Unicode case folding.
- Link labels must not contain unescaped square brackets.
- Image descriptions may contain links, but only the plain string content is used for the `alt` attribute.
- Autolinks are absolute URIs or email addresses enclosed in `< >`.
- Raw HTML tags are parsed without escaping if they match the grammar.

Entities And Concepts
- Full reference link
- Collapsed reference link
- Shortcut reference link
- Image description
- URI autolink
- Email autolink
- Raw HTML tag
- Unicode case fold

Procedures And API Details
- **Link Label Normalization**: Strip brackets, perform Unicode case fold, strip leading/trailing whitespace, collapse internal whitespace to a single space.
- **Matching**: Compare normalized forms of link labels.
- **Image Alt Attribute**: Extract plain string content from the image description; ignore inline formatting.
- **Autolink Parsing**: Detect `<scheme:...>` or `<email>` patterns; validate scheme format (2–32 chars, ASCII letter start).

Nuance Or Contradictions
- **Whitespace in Reference Links**: Unlike original Markdown, this spec forbids whitespace between link text and label to prevent unintended shortcut link capture.
- **Image Links**: While image descriptions can contain links syntactically, the rendering recommendation is to strip formatting for the `alt` text.
- **Autolink Validity**: Many strings parsed as autolinks (e.g., `m:abc`) are not valid URIs per standard registries but are accepted by this spec.

Candidate Wiki Hints
- Reference Link Precedence
- Image Alt Text Handling
- Autolink Scheme Validation
- Raw HTML Tag Grammar

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Section 3. c: HTML Tags and Attributes
- Lines 7515-7801
- Covers definitions of HTML components (open/closing tags, comments, processing instructions, etc.) and a series of examples (613–646) demonstrating valid and invalid HTML parsing within CommonMark.

Local Summary
This chunk defines the structural components of an HTML tag in CommonMark, including open tags, closing tags, comments, processing instructions, declarations, and CDATA sections. It provides numerous examples illustrating valid tag syntax, whitespace handling, attribute formatting, and illegal characters that result in HTML entities. It also introduces the concept of hard line breaks in section 6.7, explaining how line endings preceded by spaces or backslashes are rendered as `<br />` tags, with specific rules regarding their placement within inline content versus block elements.

Key Claims
- A double-quoted attribute value consists of a starting `"`, zero or more characters not including `"`, and a final `"`.
- An open tag consists of `<`, a tag name, optional attributes/spaces, an optional `/`, and `>`.
- A closing tag consists of `</`, a tag name, optional spaces, and `>`.
- HTML comments start with `<!--` and end with `-->`.
- Processing instructions start with `<?` and end with `?>`.
- Declarations start with `<!`, followed by an ASCII letter, and end with `>`.
- CDATA sections start with `<![CDATA[` and end with `]]>`.
- Hard line breaks occur when a line ending is preceded by two or more spaces or a backslash, rendering as `<br />`.
- Hard line breaks do not occur inside code spans or at the end of a block element.

Entities And Concepts
- Open Tag
- Closing Tag
- HTML Comment
- Processing Instruction
- Declaration
- CDATA Section
- Hard Line Break
- Attribute Value
- Tag Name
- Illegal Tag Names
- Illegal Attribute Names
- Illegal Attribute Values
- Illegal Whitespace

Procedures And API Details
- To create a valid open tag: `<tagname attr="value">`
- To create a valid closing tag: `</tagname>`
- To create a valid HTML comment: `<!-- comment -->`
- To create a valid processing instruction: `<?instruction?>`
- To create a valid declaration: `<!ELEMENT br EMPTY>`
- To create a valid CDATA section: `<![CDATA[content]]>`
- To create a hard line break: Use two or more spaces or a backslash before a line ending.

Nuance Or Contradictions
- Backslash escapes do not work in HTML attributes (Example 631).
- Illegal tag names (e.g., `<33>`) are not parsed as HTML but escaped (Example 618).
- Illegal attribute names (e.g., `h*#ref`) result in escaped output (Example 619).
- Illegal attribute values (e.g., unescaped quotes) result in escaped output (Example 620).
- Illegal whitespace (e.g., spaces inside tag names) results in escaped output (Example 621).
- Missing whitespace between attributes (e.g., `href='bar'title=title`) results in escaped output (Example 622).
- Hard line breaks do not occur inside code spans (Example 640) or at the end of a block element (Example 644).

Candidate Wiki Hints
- Hard Line Breaks in Markdown
- HTML Tag Syntax in CommonMark
- CommonMark HTML Parsing Rules
- Attribute Value Quoting
- Illegal HTML Characters in CommonMark

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
- **Line range**: 7802-8133
- **Coverage**: Sections 6.8 (Soft line breaks), 6.9 (Textual content), and Appendix A (Parsing strategy: block and inline structure).

### Local Summary
This chunk details how CommonMark handles soft line breaks, preserves raw textual content (including special characters and spaces), and defines a two-phase parsing strategy. Phase 1 constructs the block structure (tree of blocks) by consuming lines and managing open/closed states. Phase 2 parses the raw text within those blocks into inline elements (strings, code spans, emphasis). The text includes a detailed algorithm for handling nested emphasis and links using a "delimiter stack."

### Key Claims
- A regular line ending not preceded by two or more spaces or a backslash is parsed as a **softbreak**.
- A conforming parser may render a soft line break in HTML either as a line ending or as a space.
- Any characters not given an interpretation by previous rules are parsed as **plain textual content**.
- Internal spaces are preserved verbatim.
- Parsing occurs in two phases: first constructing the block structure, then parsing raw text contents into inline elements.
- The delimiter stack is a doubly linked list used to track opening and closing delimiters for emphasis and links.

### Entities And Concepts
- **Soft line break**: A line ending parsed as a break or space.
- **Textual content**: Plain text characters not interpreted as Markdown syntax.
- **Block structure**: The tree of blocks (document, block quotes, lists, paragraphs) constructed in Phase 1.
- **Inline structure**: The sequence of inline elements (strings, code spans, links, emphasis) constructed in Phase 2.
- **Delimiter stack**: A data structure tracking delimiters (`*`, `_`, `[`, `!`, `]`) to resolve emphasis and links.
- **Lazy continuation**: A line that continues an open block without closing it.
- **Setext headings**: Headings formed by an underline line.
- **Reference link definitions**: Detected when a paragraph is closed.

### Procedures And API Details
- **Phase 1 Procedure**:
  1. Iterate through open blocks to check conditions for remaining open.
  2. Consume continuation markers and look for new block starts (e.g., `>`). Close unmatched blocks before creating new ones.
  3. Incorporate the remainder of the line into the last open block.
- **Phase 2 Procedure**:
  - Walk the tree and parse raw string contents of paragraphs and headings as inlines.
  - Resolve reference links using the map constructed in Phase 1.
- **Inline Parsing Algorithm**:
  - Insert text nodes for runs of `*`, `_`, `[`, or `![` and add pointers to the delimiter stack.
  - On hitting `]`, call `look for link or image`.
  - On end of input, call `process emphasis`.
- **look for link or image**:
  - Search backwards from the top of the delimiter stack for an active opening `[` or `![`.
  - If found and active, parse ahead to determine link/image type (inline, reference, collapsed, shortcut).
  - If a link is found, set all `[` delimiters before the opening delimiter to inactive.
- **process emphasis**:
  - Use `stack_bottom` to limit descent.
  - Track `openers_bottom` for each delimiter type.
  - Repeat until potential closers run out:
    - Move `current_position` to the first potential closer (`*` or `_`).
    - Look back for a matching potential opener.
    - If found, determine emphasis/strong emphasis, insert nodes, and remove delimiters.
    - If not found, update `openers_bottom` and remove inactive closers.

### Nuance Or Contradictions
- **Rendering Flexibility**: While the spec defines a soft line break, the rendering in HTML is flexible; it can be a line ending or a space without changing the result in browsers.
- **Parser Options**: Renderers may provide an option to render soft line breaks as hard line breaks.
- **Delimiter Stack Logic**: The algorithm distinguishes between "active" delimiters and those that are potential openers/closers based on preceding and following characters. Inactive delimiters are removed from the stack to prevent invalid nesting (e.g., links within links).

### Candidate Wiki Hints
- **Soft Line Breaks**: A dedicated page explaining the rules for line breaks, their HTML rendering variations, and the concept of hard line breaks.
- **Parsing Strategy**: A technical page detailing the two-phase parsing approach (block vs. inline) and the tree-based document model.
- **Delimiter Stack Algorithm**: A deep-dive page on the specific algorithm for resolving nested emphasis and links, including the logic for the delimiter stack.
- **Textual Content**: A page covering the handling of special characters, internal spaces, and plain text preservation.

