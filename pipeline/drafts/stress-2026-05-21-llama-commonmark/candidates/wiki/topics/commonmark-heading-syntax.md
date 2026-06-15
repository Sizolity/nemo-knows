---
title: Commonmark Heading Syntax
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Heading Syntax

The CommonMark specification establishes a formal standard to ensure Markdown remains readable as plain text, resolving ambiguities present in the original description. Within this framework, headings are defined as block-level elements that sit alongside other structural components like code blocks and lists. The document distinguishes between two primary heading formats: ATX headings, which use hash symbols, and Setext headings, which rely on underlining and overlining.

Parsing proceeds in two distinct phases. First, the parser constructs the block structure by consuming lines to identify headings and other containers. Once the block hierarchy is established, the raw text within those blocks is processed into inline elements. This rigorous model ensures that the interpretation of headings is consistent across different implementations.

The specification details specific rules for handling tabs, treating them as four spaces in contexts that define block structure, such as indented code blocks. Backslash escapes allow punctuation characters to be treated as literals, though this mechanism is invalid within code contexts. Furthermore, indentation relative to the start of a block or container determines whether content is treated as a code block or part of a paragraph. Container blocks, such as block quotes and lists, permit the omission of markers on continuation lines to improve readability, a concept known as laziness.

Global link reference definitions affect the entire document scope rather than just the container where they appear. HTML blocks are treated as raw text and can interrupt paragraphs, subject to specific start and end conditions. Finally, any characters not interpreted by previous rules are parsed as plain textual content, preserving internal spaces verbatim. The resolution of nested emphasis and links relies on a delimiter stack algorithm to track active and inactive states.
