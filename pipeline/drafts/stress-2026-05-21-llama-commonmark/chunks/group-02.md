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
