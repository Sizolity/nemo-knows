---
title: Commonmark List Structure
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark List Structure

In the CommonMark specification, lists are treated as container blocks that organize content into distinct structural units. These containers allow for the omission of markers on continuation lines to enhance readability, a concept known as laziness. The structure of a list is defined relative to the start of the block or container, where indentation plays a critical role in determining whether content belongs to the list item or triggers a code block.

The specification distinguishes between different types of lists, including ordered and unordered varieties, and defines how they interact with other block-level elements. When a list is nested within a block quote or another container, specific rules apply to ensure the correct interpretation of indentation and markers. The parser constructs the block structure first, consuming lines to identify list items before processing the raw text within those items into inline elements.

Content within list items can include paragraphs, thematic breaks, code blocks, and other block elements, provided they adhere to the indentation and marker requirements. The handling of tabs is specific to contexts defining block structure, such as indented code blocks, while backslash escapes operate outside code contexts to treat punctuation as literals. This rigorous parsing model ensures that the resulting Markdown is unambiguous and readable as plain text without markup artifacts.
