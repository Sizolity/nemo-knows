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
