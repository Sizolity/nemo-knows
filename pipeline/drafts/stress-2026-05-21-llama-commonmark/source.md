---
title: CommonMark Specification Overview
kind: source
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## What It Is

The CommonMark Specification (Version 0.31.2, 2024-01-28) is a formal definition of the Markdown syntax designed to resolve implementation divergences found in the original Markdown description. It provides a rigorous parsing model that distinguishes between block-level structure (headings, code blocks, lists) and inline-level syntax (emphasis, links, code spans). The document is structured as a tree of blocks and defines a two-phase parsing strategy: first constructing the block structure by consuming lines, then parsing the raw text within those blocks into inline elements.

## Summary

The specification begins by establishing the necessity of a formal standard to ensure readability as plain text without markup artifacts. It defines fundamental character classes, handling rules for tabs (treated as 4 spaces in specific contexts), and backslash escapes (valid outside code contexts). The core of the document details the syntax for block-level elements, including ATX and Setext headings, indented and fenced code blocks, thematic breaks, and HTML blocks. It further defines container blocks like block quotes and lists, explaining complex rules regarding indentation, "laziness," and the inclusion of content. Finally, it covers inline elements such as links, emphasis, and code spans, along with the global scope of link reference definitions.

The document outlines the evolution of a Markdown parser's logic, moving from high-level document structure to low-level parsing algorithms. Early sections cover document metadata and basic list handling, while subsequent sections introduce complexity through nested headings, code spans, and raw text preservation. Later sections provide the rigorous algorithmic definitions for how these structures are interpreted, specifically detailing the handling of soft line breaks versus hard breaks, the definition of textual content, and the specific mechanics of the delimiter stack used to resolve nested emphasis and links.

## Key Claims

- **Readability Goal**: Markdown is designed to be readable as plain text without markup artifacts.
- **Ambiguity Resolution**: The CommonMark spec aims to be unambiguous, contrasting with the original Markdown description which allowed for implementation divergence.
- **Tab Handling**: Tabs are not globally expanded but are treated as 4 spaces specifically in contexts defining block structure (e.g., indented code blocks).
- **Backslash Escaping**: Backslashes escape ASCII punctuation characters to treat them as literals, but this mechanism is invalid within code blocks, code spans, autolinks, and raw HTML.
- **Indentation Rules**: Indentation is relative to the start of a block or container. 4 spaces generally trigger a code block; 3 spaces or fewer are often skipped or used for container markers.
- **Laziness**: Container blocks (block quotes, lists) allow the omission of markers on continuation lines of paragraphs to improve readability.
- **Global Scope**: Link reference definitions affect the entire document scope, not just the container in which they appear.
- **HTML Blocks**: HTML blocks are treated as raw text and can interrupt paragraphs, with specific rules for start/end conditions and indentation.
- **Two-Phase Parsing**: The specification consistently emphasizes that parsing occurs in two distinct phases: first constructing the block structure, then parsing raw text contents into inline elements.
- **Textual Content**: Any characters not given an interpretation by previous rules are parsed as plain textual content, with internal spaces preserved verbatim.
- **Delimiter Stack**: The resolution of emphasis and links relies on a specific algorithm using a delimiter stack to track state, distinguishing between active and inactive delimiters.
- **Soft Line Breaks**: A regular line ending not preceded by two or more spaces or a backslash is parsed as a soft break, which may render as a line ending or a space in HTML.
- **Block Structure**: The document is modeled as a tree of blocks (document, block quotes, lists, paragraphs).

## Suggested Links

none
