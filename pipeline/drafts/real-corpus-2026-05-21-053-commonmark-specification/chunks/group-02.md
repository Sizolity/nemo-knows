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
