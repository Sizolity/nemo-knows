---
title: Commonmark Escaping And Entities
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Escaping And Entities

The CommonMark specification defines how to handle special characters and HTML entities within Markdown documents. This mechanism ensures that structural symbols retain their intended meaning while allowing authors to include literal punctuation or Unicode characters without breaking the document's layout.

## Escaping Special Characters

Authors can prevent a character from being interpreted as a structural element by prefixing it with a backslash. This technique is effective for ASCII punctuation marks in most contexts. However, backslash escapes are invalid inside code blocks, code spans, autolinks, and raw HTML blocks.

## Entity and Numeric Character References

The specification permits the use of valid HTML entity references and numeric character references in place of corresponding Unicode characters. These references are treated as literal text within code spans and code blocks. Furthermore, entities cannot substitute for symbols that define structural elements, such as asterisks in emphasis delimiters or bullets in list markers.

## Precedence and Structure

When parsing inline content, code spans, links, and HTML tags interrupt or take precedence over emphasis markers. Links bind more tightly than brackets, which bind more tightly than emphasis. This hierarchy ensures that structural elements are recognized correctly regardless of surrounding punctuation.

## Line Breaks and Soft Breaks

Regular line endings not preceded by spaces or backslashes are parsed as soft breaks. Renderers may choose to render these as a line ending or a space, providing flexibility in how text is displayed.

## Related Concepts

- [[commonmark-block-structure]]
- [[commonmark-inline-elements]]
- [[commonmark-line-breaks]]
