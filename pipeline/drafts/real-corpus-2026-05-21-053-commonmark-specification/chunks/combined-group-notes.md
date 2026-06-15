## group-01

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

This group of notes synthesizes the **CommonMark Specification** (Version 0.31.2), authored by John MacFarlane under a Creative Commons BY-SA license. The document addresses the lack of unambiguous syntax in the original Markdown description, which led to divergent implementations across platforms. The specification aims to resolve ambiguities regarding lists, code blocks, headings, and inline structure through side-by-side Markdown/HTML examples that serve as conformance tests.

The covered text spans from the introduction and metadata (Chunk 01) through definitions of block-level elements like thematic breaks and ATX headings (Chunk 02), Setext-style headers and indented code blocks (Chunk 03), HTML blocks and link reference definitions (Chunks 04-05), container blocks such as block quotes and lists (Chunk 06), and continues through subsequent sections on lists and inline processing (Chunks 07-13).

# Cross-Chunk Summary

The specification is structured into distinct phases of parsing: **Block Structure** followed by **Inline Structure**.
- **Phase 1 (Block Structure)**: Identifies container blocks (lists, block quotes) and leaf blocks (paragraphs, code blocks, thematic breaks). This phase handles the "hard" boundaries between elements.
- **Phase 2 (Inline Structure)**: Processes inlines (links, emphasis, code spans) within identified blocks.

The document systematically defines syntax for:
1.  **Metadata & Introduction**: Rationale for the spec, conformance testing tools (`spec_tests.py`), and character definitions (Unicode code points).
2.  **Block-Level Elements**:
    -   **Thematic Breaks**: Horizontal rules using `-`, `_`, or `*`.
    -   **ATX Headings**: Using `#` characters with specific spacing rules.
    -   **Setext Headings**: Underline-based (`=` or `-`) supporting multiline content.
    -   **Code Blocks**: Fenced (backtick/tilde) and indented (4+ spaces).
    -   **HTML Blocks**: Seven types defined by start/end conditions, allowing interruption of paragraphs.
3.  **Container Blocks**: Block quotes (`>`) and List items (bullets/ordered), including rules for "laziness" (omitting markers on continuation lines) and relative indentation.
4.  **Inline Elements**: Links, emphasis, code spans, and entities.

# Repeated Or Central Claims

-   **Readability Goal**: Markdown is designed so that a formatted document is readable as plain text without appearing marked up with tags or instructions.
-   **Ambiguity in Original Syntax**: Gruber’s original description does not specify rules for sublist indentation, blank lines before block quotes/headings, code block requirements, list item wrapping, right-aligned markers, thematic breaks within lists, marker precedence, section headings inside lists, empty list items, link reference scope, or definition precedence.
-   **Implementation Divergence**: Without a formal spec, implementations consulted buggy scripts (Markdown.pl) or made arbitrary choices, leading to documents rendering differently on different systems without triggering syntax errors.
-   **HTML as Test Representation**: The spec uses HTML for side-by-side examples because it represents structural distinctions well; however, not every HTML feature in the examples is mandated by the spec.
-   **Block Structure Precedence**: Indicators of block structure always take precedence over indicators of inline structure. For example, a `>` marker starts a block quote even if preceded by text that looks like emphasis.
-   **Backslash Escaping Limits**: Backslashes escape only specific characters (like `*`, `#`) in certain contexts but not others (like code blocks). Escapes do not work in code blocks, code spans, autolinks, or raw HTML.
-   **Indentation Rules**: Tabs are not expanded to spaces generally but behave as if replaced by four spaces in contexts defining block structure (e.g., indented code blocks, list item continuation). Internal tabs within content are passed through literally.
-   **Blank Line Semantics**: Multiple blank lines between paragraphs have no effect on paragraph boundaries; they define list tightness/looseness or separate distinct block quotes. Blank lines inside HTML blocks are disallowed to avoid expensive balanced tag parsing.

# Important Local Details

-   **Conformance Testing**: The document outlines a two-phase parsing strategy and provides the command `python test/spec_tests.py --spec spec.txt --program PROGRAM` for execution.
-   **Tooling**: `tools/makespec.py` converts the source text file (`spec.txt`) into HTML or CommonMark.
-   **Entity Parsing**: Valid HTML entity references (e.g., `&copy;`) and numeric character references (`&#35;`, `&#x22;`) can be used in place of Unicode characters, except in code contexts and structural elements like emphasis delimiters or list markers. CommonMark requires the semicolon for entities to avoid grammar ambiguity.
-   **ATX Heading Nuances**: ATX headings use 1–6 unescaped `#` characters. More than six is not a heading. They must have at least one space between the opening `#` and content unless empty. While some original implementations required a space after `#`, the spec allows headings without it if strictly following the grammar, though many implementations still require it to avoid ambiguity (e.g., `#5 bolt`).
-   **Setext Heading Multiline**: A Setext heading can span multiple lines provided a blank line separates preceding paragraphs. The underline must not exceed three spaces of indentation; four spaces terminates the block as code or paragraph content. Underlines cannot contain internal spaces or tabs.
-   **HTML Block Types**:
    -   Type 1: `<pre>`, `<script>`, `<style>`, `<textarea>` (ends at matching tag).
    -   Type 2: `<!--` ... `-->`.
    -   Type 3: `<?` ... `?>`.
    -   Type 4: `<!` followed by ASCII letter ... `>`.
    -   Type 5: `<![CDATA[` ... `]]>`.
    -   Type 6: Block-level tags (ends at blank line).
    -   Type 7: Any complete open/closing tag (ends at blank line).
-   **List Item Indentation**: List item indentation is relative. Content must be indented sufficiently past the list marker and any preceding containers. Lines may be uniformly indented by up to three spaces without changing the list item status. Ordered list start numbers are limited to 9 digits due to browser integer overflow concerns; leading zeros are allowed.
-   **Lazy Continuation**: In block quotes, the `>` marker can be omitted on lines containing paragraph continuation text ("laziness"), provided indentation requirements relative to containing blocks are met. This does not apply to lines starting new block types (like code blocks).
-   **Link Reference Definitions**: Consist of a label, colon, destination, and optional title. They do not correspond to structural elements but can follow each other without blank lines and appear inside block containers.

# Candidate Wiki Hints

-   **Topic: CommonMark Specification Overview** – Summary of the rationale, versioning, and conformance testing methodology.
-   **Topic: Block Structure vs. Inline Structure** – Explanation of the two-phase parsing strategy and precedence rules.
-   **Topic: Escaping and Entities** – Rules for backslash escaping, HTML entities, and numeric character references.
-   **Topic: Heading Syntax (ATX and Setext)** – Comparison of `#` based headings versus underline-based headings, including multiline support.
-   **Topic: Code Blocks** – Distinction between fenced (backtick/tilde) and indented code blocks, info string rules, and indentation limits.
-   **Topic: HTML Blocks** – The seven types of HTML block start/end conditions and their interaction with paragraphs.
-   **Topic: Block Quotes** – Syntax for `>` markers, laziness rules, nesting, and separation.
-   **Topic: List Items** – Bullet and ordered markers, indentation math, relative positioning, and empty items.

# Gaps Or Cautions

-   **Missing Chunk Content**: The provided chunk notes cover the specification up through "3. c > foo\" (Chunk 13). While the outline suggests chunks continue beyond this, the detailed notes for Chunks 07 through 12 are missing from the input data (only summaries of headings like "aaa" or "3. c" are present in the index, but no specific content notes were provided in the `chunk-notes` section for these indices). This group note relies heavily on the detailed notes for Chunks 01-06 and the general synthesis from the outline.
-   **Original vs. Spec Contradictions**: Be aware that Gruber’s original Markdown syntax often differs significantly from CommonMark (e.g., requiring blank lines around block-level HTML, forbidding indentation in lists). This spec relaxes some rules but removes others (like blank lines inside HTML blocks) to simplify parsing and ensure interoperability.
-   **Implementation Variance**: Some legacy implementations do not support features allowed by the spec, such as multiline Setext headings or specific behaviors of lazy continuation. Users relying on these features must be aware that their documents may not render correctly on older parsers.

## group-02

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

This document group covers **Section 3.c** of the CommonMark specification, focusing on the syntax and parsing rules for lists. It details how to distinguish between "tight" and "loose" lists based on blank lines and block content, defines strict indentation rules (the "four-space rule" vs. relative indentation from list markers), and explains how lists interact with paragraphs and code blocks. Additionally, it covers the transition to inline parsing within lists, including the complex algorithms for emphasis (`*`, `_`), strong emphasis (`**`, `__`), code spans (backticks), links (inline, reference, image, autolink), and raw HTML tags.

# Cross-Chunk Summary

The text systematically moves from block-level list structure to inline content parsing:
1.  **Block Structure**: Defines how lists are formed, how they interrupt paragraphs, and the specific indentation requirements for sublists and indented code blocks (distinguishing between paragraph continuations and code blocks).
2.  **List Classification**: Establishes criteria for "loose" vs. "tight" lists based on internal blank lines and block separation.
3.  **Inline Syntax**: Details the precedence hierarchy for inline elements (Code Spans > Links > Emphasis) and provides specific parsing algorithms for emphasis delimiters, link references (full, collapsed, shortcut), image syntax, and HTML tags within list items.

# Repeated Or Central Claims

*   **Indentation Rules**: Sublists must be indented relative to the parent list marker; a fixed "four-space rule" from the margin is often arbitrary, though supported. Indenting more than three spaces before a list item changes parsing (often treating subsequent text as a continuation or code block).
*   **List Interruption**: Lists may interrupt paragraphs in CommonMark, specifically for ordered lists starting with `1`, to prevent capturing unintended lists from hard-wrapped text (e.g., "is 14.").
*   **No Nested Links**: Links cannot contain other links at any nesting level. If nested, the innermost link definition takes precedence, though this contradicts some legacy Markdown behaviors where reference definitions were treated differently.
*   **Precedence Hierarchy**: Inline elements have a strict parsing order: Code Spans > Autolinks/HTML Tags > Links > Emphasis. This ensures that backticks and HTML tags are not misinterpreted as emphasis markers or link text.
*   **Delimiter Flanking**: Emphasis delimiters (`*`, `_`) require specific "flanking" conditions (left or right) based on surrounding whitespace and punctuation, allowing for intraword emphasis with asterisks but restricting underscores to word boundaries.

# Important Local Details

*   **Tight vs. Loose Lists**: A list is "loose" if there is a blank line between items, or if an item contains two block-level elements separated by a blank line (even if that blank line is not physically present in the raw text but implied by block separation). A "tight" list has no such internal separation.
*   **Code Span Normalization**: Contents of code spans delimited by backticks are normalized by converting line endings to spaces and stripping single leading/trailing spaces if they exist on both sides. Backslashes inside code spans are literal (no escaping).
*   **Reference Link Matching**: Labels are normalized via case folding, whitespace collapse, and stripping brackets before matching. Full and collapsed references take precedence over shortcuts. Spaces between link text and label are forbidden for full/collapsed references.
*   **HTML Attribute Escaping**: Backslash escapes do not work inside HTML attributes; they are preserved as literal characters. Hard line breaks (two spaces + newline or `\` + newline) inside attribute values result in the literal string including the break, unlike outside tags where they might become `<br />`.
*   **Image Syntax**: Image descriptions can contain links, but the `alt` attribute is derived from plain text only, ignoring any nested HTML structure like `<a>` tags.

# Candidate Wiki Hints

*   **CommonMark List Indentation**: Rules for calculating sublist indentation relative to list markers.
*   **Tight vs. Loose Lists**: Criteria and examples distinguishing these two list types in CommonMark.
*   **Emphasis Algorithms**: Delimiter runs, flanking rules, and nesting constraints for `*`, `_`, `**`, `__`.
*   **Link Reference Syntax**: Full, collapsed, and shortcut reference link definitions and matching logic.
*   **Inline Precedence**: How code spans, autolinks, and HTML tags override emphasis and link parsing.
*   **HTML Tag Parsing**: Rules for valid tag names, attributes, comments, and handling of illegal syntax within list contexts.

# Gaps Or Cautions

*   **Markdown.pl Legacy**: CommonMark adopts a more forgiving indentation strategy to handle legacy `Markdown.pl` behavior alongside the strict four-space rule, which may confuse users expecting consistent spacing requirements.
*   **Ambiguity Resolution**: The spec minimizes nesting depth for emphasis (preferring `<strong>` over `<em><em>`) and resolves ambiguities in link parsing by preferring definitions that appear first or are more specific.
*   **Hard-Wrapped Text Risk**: While allowing lists to interrupt paragraphs improves readability, it risks capturing unintended lists; the restriction on ordered lists starting with `1` is a necessary mitigation for this edge case.
*   **Backslash Limitations**: Users should be aware that backslash escapes are effective in code spans and inline text but are ignored within HTML attributes, potentially leading to unexpected rendering if not careful.

## group-03

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

This document group constitutes a substantial portion of the **CommonMark Specification**, specifically focusing on the transition from initial metadata fetching and introductory examples into the core parsing algorithms. The content spans from the top-level document structure through specific edge cases involving list continuations, nested emphasis, soft line breaks, and textual content preservation. The sequence covers the implementation strategy for a two-phase parser (block structure followed by inline structure) and details the stack-based algorithm for resolving nested delimiters.

# Cross-Chunk Summary

The notes progress from high-level document identification to specific technical deep-dives:
1.  **Initial Structure (Chunks 01–04):** Establishes the document root, fetch metadata, retrieved text handling, and list item continuations. It covers basic structural elements like nested lists and block transitions.
2.  **Edge Cases & Examples (Chunks 05–07):** Explores specific formatting scenarios including bracketed items (`[Foo]`), deeply nested "foo" headers illustrating structural complexity, and textual content handling.
3.  **Parsing Logic (Chunks 08–13):** Shifts focus to the mechanical implementation of the parser. This includes the definition of soft line breaks, the two-phase parsing model (building block trees via lazy continuations), and the detailed stack-based algorithm for processing inline elements like links and emphasis.

# Repeated Or Central Claims

-   **Two-Phase Parsing Model:** The document consistently asserts that CommonMark parsing occurs in two distinct phases: Phase 1 constructs a block-level tree structure (handling indentation, markers, and lazy continuations) without fully parsing inline content; Phase 2 then parses the raw text of paragraphs and headings into inline elements using a pre-computed link map.
-   **Soft Line Breaks:** It is repeatedly noted that line endings not preceded by two or more spaces or a backslash are treated as soft breaks. While renderers have flexibility in outputting these as spaces or newlines, browsers typically treat them identically to hard breaks visually.
-   **Lazy Continuation:** The concept of "lazy continuation" is central to the block structure phase, where lines that do not start new blocks (e.g., block quotes marked by `>`) are added to the current open block even if the block's structural condition changes mid-line, ensuring complex nesting logic is handled correctly.
-   **Delimiter Stack:** The use of a doubly linked list (delimiter stack) to track potential openers (`[`, `*`, `_`) and closers for resolving nested emphasis and links is a persistent technical detail across the inline parsing sections.

# Important Local Details

-   **Block Closure Timing:** Blocks are not strictly closed immediately upon encountering a new block start on the same logical line. Instead, unmatched blocks from the previous step are closed *before* creating the new block, allowing for complex nesting logic within a single line.
-   **Delimiter Stack Algorithm Mechanics:**
    -   **Insertion:** When hitting `*`, `_`, `[`, or `![`, a text node with literal content is inserted, and a pointer is added to the delimiter stack.
    -   **Stack Element Data:** Contains pointers to text nodes, delimiter type, count of delimiters, active status, and potential role (opener/closer).
    -   **Link/Image Resolution:** The algorithm starts at the top of the stack, looking back for opening brackets. If inactive, it returns a literal `]`. If active, it parses ahead for link/image types.
    -   **Emphasis Processing:** Iterates until potential closers are exhausted, moves forward to find the first potential closer (`*` or `_`), looks back above `stack_bottom` for matching openers, determines strength (length >= 2), inserts emph/strong nodes, and updates text nodes.
-   **Textual Content Preservation:** Characters not matched by specific rules (links, emphasis, code spans) are parsed as plain text, preserving internal spaces verbatim.

# Candidate Wiki Hints

-   **Page: CommonMark Parsing Strategy** (Concept: Two-phase parsing model of block vs. inline processing).
-   **Page: Delimiter Stack Algorithm** (Concept: Handling nested emphasis and links using a stack-based approach).
-   **Page: Soft Line Breaks** (Concept: Rules for soft breaks and rendering variance between browsers and renderers).
-   **Page: Block Tree Construction** (Concept: Building hierarchical document representations via lazy continuations).

# Gaps Or Cautions

-   **Rendering Variance:** The specification explicitly notes that while it mandates how soft breaks are parsed, the visual output in browsers is identical for soft and hard breaks, implying significant renderer flexibility regarding the final CSS/HTML generation of these line endings.
-   **Incomplete Header Paths:** Several chunks (specifically Chunks 02–13) show heading paths that are either "not a heading" or generic placeholders like "foo", "[Foo]", or "aaa". This suggests these sections may be testing edge cases, examples, or specific internal states rather than defining the primary structural headers of the document itself.
-   **Line Continuation Ambiguity:** The handling of lines where a block quote marker (`>`) appears but the content continues with text that might suggest a new block requires careful attention to the "lazy continuation" rules to avoid misinterpreting block boundaries.

