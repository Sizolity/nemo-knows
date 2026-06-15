## group-01

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

- **Source**: CommonMark Specification (Version 0.31.2, 2024-01-28)
- **Scope**: A comprehensive overview of the CommonMark specification, covering its rationale, fundamental character definitions, block-level syntax (headings, code blocks, thematic breaks), container blocks (block quotes, lists), and inline syntax (links, emphasis).
- **Key Sections**:
  - Introduction (What is Markdown, Why a spec is needed)
  - Characters and Lines (Tabs, Insecure characters, Backslash escapes)
  - Block Structure (Headings, Code blocks, Thematic breaks, HTML blocks)
  - Container Blocks (Block quotes, Lists)
  - Inline Elements (Links, Emphasis, Code spans)
  - Parsing Rules (Precedence, Laziness, Indentation)

# Cross-Chunk Summary

The document begins by establishing the necessity of a formal specification to resolve implementation divergences found in the original Markdown description. It defines fundamental character classes and rules for handling tabs and backslash escapes. The specification then details the syntax for various block-level elements, including ATX and Setext headings, indented and fenced code blocks, thematic breaks, and HTML blocks. It further defines container blocks like block quotes and lists, explaining complex rules regarding indentation, "laziness," and the inclusion of content. Finally, it covers inline elements such as links, emphasis, and code spans, along with the global scope of link reference definitions.

# Repeated Or Central Claims

- **Readability Goal**: Markdown is designed to be readable as plain text without markup artifacts.
- **Ambiguity Resolution**: The CommonMark spec aims to be unambiguous, contrasting with the original Markdown description which allowed for implementation divergence (e.g., list indentation, blank line requirements).
- **Tab Handling**: Tabs are not globally expanded but are treated as 4 spaces specifically in contexts defining block structure (e.g., indented code blocks).
- **Backslash Escaping**: Backslashes escape ASCII punctuation characters to treat them as literals, but this mechanism is invalid within code blocks, code spans, autolinks, and raw HTML.
- **Indentation Rules**: Indentation is relative to the start of a block or container. 4 spaces generally trigger a code block; 3 spaces or fewer are often skipped or used for container markers.
- **Laziness**: Container blocks (block quotes, lists) allow the omission of markers on continuation lines of paragraphs to improve readability.
- **Global Scope**: Link reference definitions affect the entire document scope, not just the container in which they appear.
- **HTML Blocks**: HTML blocks are treated as raw text and can interrupt paragraphs (unlike Gruber's original spec), with specific rules for start/end conditions and indentation.

# Important Local Details

- **Line Endings**: Defined as U+000A, U+000D (not followed by U+000A), or U+000D followed by U+000A.
- **Unicode Handling**: The Unicode character U+0000 must be replaced with the REPLACEMENT CHARACTER (U+FFFD).
- **ATX Headings**: Require 1–6 unescaped `#` characters at the start, followed by spaces or tabs. Indentation of 4 spaces converts a heading to a code block.
- **Setext Headings**: Defined by text lines followed by an underline of `=` (level 1) or `-` (level 2). Multiline content is supported by the spec but not widely implemented.
- **Fenced Code Blocks**: Defined by 3+ backticks or tildes. Info strings are optional. Closing fences must match the opening character and length.
- **Indented Code Blocks**: Composed of lines indented by at least four spaces. Content is literal text.
- **Thematic Breaks**: Consist of 3+ matching `-`, `_`, or `*` characters with optional indentation (up to 3 spaces).
- **Link Reference Definitions**: Consist of a label, colon, destination, and optional title. Matching is case-insensitive.
- **Paragraph Formation**: A sequence of non-blank lines that cannot be interpreted as other blocks. Leading spaces/tabs are skipped; final spaces/tabs are stripped.
- **Block Quote Markers**: The `>` character preceded by up to three spaces of indentation.
- **List Markers**: Bullet (`-`, `+`, `*`) or ordered (1–9 digits) markers. Ordered list numbers are limited to 9 digits.

# Candidate Wiki Hints

- **Page: CommonMark Specification Overview**
  - Summary: Introduction to Markdown, rationale for the spec, and fundamental definitions.
- **Page: Backslash Escaping Rules**
  - Summary: When and how to use backslashes to escape special characters, excluding code contexts.
- **Page: Entity Reference Usage**
  - Summary: Best practices for using HTML entities, noting restrictions in structural contexts.
- **Page: ATX Heading Syntax**
  - Summary: Rules for creating headings with `#` characters, including closing sequences and indentation limits.
- **Page: Setext Headings**
  - Summary: Rules for defining level 1 and 2 headings using underlines (`=` or `-`).
- **Page: Indented Code Blocks**
  - Summary: Syntax and behavior of code blocks defined by 4+ space indentation.
- **Page: Fenced Code Blocks**
  - Summary: Syntax and behavior of code blocks defined by backticks or tildes, including info strings.
- **Page: HTML Blocks in Markdown**
  - Summary: How HTML blocks are parsed, including the seven types and their interaction with Markdown syntax.
- **Page: Link Reference Definitions**
  - Summary: Rules for defining and using reference links, including global scope and placement.
- **Page: Paragraph Parsing**
  - Summary: How CommonMark handles line breaks, indentation, and blank lines within paragraphs.
- **Page: CommonMark Container Blocks**
  - Summary: Overview of block quotes and list items as container blocks.
- **Page: CommonMark Block Quote Syntax**
  - Summary: Rules for `>` markers, indentation limits, and the Laziness rule.
- **Page: CommonMark List Item Indentation**
  - Summary: How indentation depth determines content inclusion in list items.
- **Page: CommonMark Ordered List Numbers**
  - Summary: Constraints on ordered list start numbers (1–9 digits, no negatives).

# Gaps Or Cautions

- **Multiline Setext Headings**: While the spec supports multiline headings, most existing implementations do not, creating a compatibility gap.
- **Tab Expansion**: Tabs are not expanded globally; users must be aware that tabs behave as 4 spaces only in specific block structure contexts.
- **Entity Ambiguity**: HTML5 allows entities without semicolons, but CommonMark requires the semicolon to avoid grammar ambiguity.
- **Blank Lines in HTML Blocks**: CommonMark disallows blank lines inside HTML blocks (except types 1–5) to avoid expensive parsing, differing from Gruber's original rule.
- **Indentation Sensitivity**: HTML blocks can be preceded by up to three spaces; four spaces trigger a code block instead.
- **Fence Matching**: The closing fence must use the same character and be at least as long as the opening fence; mixing characters or lengths is invalid.
- **Internal Spaces in Fences**: Code fences cannot contain internal spaces or tabs.
- **Link Definition Scope**: Definitions inside block containers affect the entire document, which may be unintuitive for users expecting local scope.
- **Hard Line Breaks**: Two or more spaces at the end of a line are stripped before inline parsing, preventing hard line breaks in many contexts.

## group-02

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

This group of notes synthesizes sections of the CommonMark specification covering list indentation rules, loose versus tight lists, inline parsing (emphasis, strong emphasis, links, images, autolinks), and raw HTML syntax. The content spans from the definition of list markers and indentation logic (chunks 07–08) through detailed inline element parsing and HTML tag grammar (chunks 09–12).

# Cross-Chunk Summary

The specification defines a rigorous parsing model where block structure (lists, paragraphs, code blocks) is determined by indentation and blank lines, while inline structure (emphasis, links, HTML) is determined by delimiter runs and nesting precedence.

- **Lists**: Sublists must be indented relative to the list marker, not the margin. The "four-space rule" is the standard for code blocks within lists. Lists are classified as "loose" (separated by blank lines) or "tight" (no internal blank lines).
- **Inline Elements**: Parsing proceeds left-to-right. Emphasis and strong emphasis rely on "delimiter runs" (left/right-flanking) and ambiguity resolution (minimizing nesting depth). Links bind tighter than emphasis but looser than code spans.
- **HTML**: Raw HTML tags are parsed if they match the grammar. Hard line breaks (preceded by two spaces or a backslash) render as `<br />` but are forbidden inside code spans or at the end of block elements.

# Repeated Or Central Claims

- **Indentation Rules**: List items cannot be preceded by more than three spaces; more than three spaces implies paragraph continuation. Four or more spaces (with a preceding blank line) implies an indented code block.
- **Loose vs. Tight Lists**: A list is loose if items are separated by blank lines or contain internal block-level elements separated by blank lines. A list is tight if blank lines appear only within code blocks, block quotes, or between sublist paragraphs.
- **Delimiter Runs**: Emphasis and strong emphasis are determined by the classification of delimiter runs (left-flanking, right-flanking, both). Ambiguities are resolved by minimizing nesting depth and preferring shorter spans.
- **Link Precedence**: Links cannot contain nested links. Link brackets bind tighter than emphasis markers but looser than code spans.
- **HTML Parsing**: Raw HTML tags are accepted if they match the grammar. Illegal characters (e.g., spaces in tag names, unescaped quotes) result in escaped output rather than parsing errors.

# Important Local Details

- **Four-Space Rule**: Indentation for blocks under list items is measured from the start of the list marker. Code blocks require eight spaces from the margin (or six from the marker in some proposals).
- **Code Span Normalization**: Line endings are converted to spaces. Leading/trailing spaces are stripped only if spaces exist on both sides.
- **Emphasis Ambiguity**: Intraword emphasis with `_` is generally disallowed to avoid unwanted emphasis in words containing internal underscores. Backslash escapes do not work inside code spans.
- **Link Structure**: Link destinations may contain spaces only if enclosed in `<...>`; line endings are forbidden. Titles use `""`, `''`, or `()`.
- **Hard Line Breaks**: Occur when a line ending is preceded by two or more spaces or a backslash. They do not occur inside code spans or at the end of a block element.

# Candidate Wiki Hints

- **CommonMark List Indentation**: Rules for sublists, the four-space rule, and loose/tight classification.
- **CommonMark Inline Syntax**: Delimiter runs, nesting logic, and precedence of inline elements.
- **CommonMark Link Syntax**: Reference link matching, image descriptions, and autolink validation.
- **CommonMark HTML Tags**: Grammar for open/closing tags, comments, and processing instructions.
- **CommonMark Hard Line Breaks**: Rules for rendering `<br />` and restrictions on placement.

# Gaps Or Cautions

- **Implementation Variance**: Some implementations (Markdown.pl, Pandoc, discount) historically handled indentation differently (e.g., two-space vs. four-space), which may cause confusion when migrating to CommonMark.
- **Escaping Limitations**: Backslash escapes are not functional within code spans or HTML attributes in the same way they are in plain text; they are treated literally or result in escaped output.
- **Whitespace Sensitivity**: Whitespace handling is strict; spaces inside tag names or unescaped quotes in attributes break parsing and result in entity escaping rather than valid HTML.
- **Autolink Validity**: The spec accepts strings like `m:abc` as autolinks even if they are not valid URIs per standard registries, which may lead to rendering issues in strict environments.

## group-03

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

This group of notes covers the **CommonMark Specification**, specifically focusing on the transition from high-level document structure to low-level parsing algorithms. The content spans from the initial document metadata and list structures (Chunks 01-02) through various edge cases involving nested headings, code spans, and raw text (Chunks 03-07), culminating in the detailed parsing strategies for block and inline structures, soft line breaks, and the delimiter stack algorithm (Chunks 08-13).

The source material defines a two-phase parsing strategy:
1.  **Phase 1**: Constructing the block structure (tree of blocks) by consuming lines and managing open/closed states.
2.  **Phase 2**: Parsing the raw text within those blocks into inline elements (strings, code spans, emphasis, links).

# Cross-Chunk Summary

The document outlines the evolution of a Markdown parser's logic. Early chunks establish the document root, metadata fetching, and basic list handling. Subsequent chunks introduce complexity through nested headings, code spans (`foo *bar* \*baz\*`), and raw text preservation. The later chunks (08-13) provide the rigorous algorithmic definitions for how these structures are interpreted, specifically detailing:
-   The handling of soft line breaks versus hard breaks.
-   The definition of textual content as any character not interpreted by syntax rules.
-   The specific mechanics of the **delimiter stack**, a doubly linked list used to track opening and closing delimiters (`*`, `_`, `[`, `!`, `]`) to resolve nested emphasis and links.
-   The distinction between active delimiters and potential openers/closers to prevent invalid nesting (e.g., links within links).

# Repeated Or Central Claims

-   **Two-Phase Parsing**: The specification consistently emphasizes that parsing occurs in two distinct phases: first constructing the block structure, then parsing raw text contents into inline elements.
-   **Textual Content**: Any characters not given an interpretation by previous rules are parsed as plain textual content, with internal spaces preserved verbatim.
-   **Delimiter Stack**: The resolution of emphasis and links relies on a specific algorithm using a delimiter stack to track state, distinguishing between active and inactive delimiters.
-   **Soft Line Breaks**: A regular line ending not preceded by two or more spaces or a backslash is parsed as a soft break, which may render as a line ending or a space in HTML.
-   **Block Structure**: The document is modeled as a tree of blocks (document, block quotes, lists, paragraphs).

# Important Local Details

-   **Chunk 01**: Covers the document root, fetch metadata, retrieved text, and basic list item continuations (e.g., "List item two continued with an open block").
-   **Chunk 02**: Introduces complex nested heading paths (e.g., `foo > foo > foo > foo > foo > foo`) and raw text variations including code spans and HTML-like tags (`bar</p>`).
-   **Chunk 03-04**: Focuses on the "baz" heading path and the transition to "Heading" and "baz" structures.
-   **Chunk 05-07**: Covers specific bracketed headings (`[Foo]`) and the "aaa" heading path, representing sections of the spec dealing with specific syntax elements.
-   **Chunk 08-12**: Deep dive into list items (`2. b`, `3. c`) and their continuations across multiple line ranges (4756-7801).
-   **Chunk 13**: Details the "3. c > foo\" path, covering soft line breaks, textual content, and the full algorithm for the delimiter stack and lazy continuation.

# Candidate Wiki Hints

-   **Soft Line Breaks**: A dedicated page explaining the rules for line breaks, their HTML rendering variations, and the concept of hard line breaks.
-   **Parsing Strategy**: A technical page detailing the two-phase parsing approach (block vs. inline) and the tree-based document model.
-   **Delimiter Stack Algorithm**: A deep-dive page on the specific algorithm for resolving nested emphasis and links, including the logic for the delimiter stack.
-   **Textual Content**: A page covering the handling of special characters, internal spaces, and plain text preservation.
-   **Block Structure**: A page defining the tree of blocks (document, block quotes, lists, paragraphs) and how open/closed states are managed.

# Gaps Or Cautions

-   **Rendering Flexibility**: While the spec defines a soft line break, the rendering in HTML is flexible; it can be a line ending or a space without changing the result in browsers. Renderers may provide an option to render soft line breaks as hard line breaks.
-   **Parser Options**: The specification notes that renderers may offer options regarding how soft line breaks are handled, implying implementation variance.
-   **Delimiter Logic Nuance**: The algorithm distinguishes between "active" delimiters and those that are potential openers/closers based on preceding and following characters. Inactive delimiters are removed from the stack to prevent invalid nesting.
-   **Line Range Continuity**: The chunk notes indicate specific line ranges (e.g., 7802-8133 for Chunk 13), suggesting that the full document is segmented. Care must be taken to ensure synthesis does not assume content exists outside the provided line ranges or chunk notes.

