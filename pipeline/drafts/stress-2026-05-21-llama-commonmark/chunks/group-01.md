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
