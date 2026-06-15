---
title: Commonmark Line Breaks
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Line Breaks

In CommonMark, a line break is a structural element that represents a newline in the source text. The specification distinguishes between hard line breaks, which are explicitly marked, and soft line breaks, which occur at the end of a line without a preceding space or backslash.

Hard line breaks are created by placing two or more spaces at the end of a line. This syntax signals to the parser that the line should terminate with a line break character rather than a space. Conversely, soft line breaks occur when a line ends without these markers. The parser treats these as optional breaks, allowing renderers to decide whether to display them as a visual line break or simply as a space, depending on the desired output format.

The specification also defines rules for escaping special characters using backslashes. While backslashes can escape punctuation in most contexts, they do not function inside autolinks. Additionally, entity and numeric character references are treated as literal text within code spans and code blocks, preventing them from being interpreted as structural elements like emphasis markers or list bullets.

For more details on the broader block structure, see [[commonmark-block-structure]]. Information on how special characters are handled can be found in [[commonmark-escaping-and-entities]].
