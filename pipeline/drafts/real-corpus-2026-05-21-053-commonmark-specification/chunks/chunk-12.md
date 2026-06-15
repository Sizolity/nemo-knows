---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
Section 3.c defines the syntax of HTML tags, attributes, and specific constructs (comments, processing instructions, declarations, CDATA sections). It provides a series of examples illustrating valid open/closing tags, empty elements, whitespace handling, illegal tag/attribute names/values, and special character preservation. Section 6.7 introduces hard line breaks within HTML contexts.

## Local Summary
This chunk details the structural components of an HTML tag in CommonMark, distinguishing between open tags (containing a name, optional attributes, and optional closing slash) and closing tags. It enumerates valid content types allowed inside tags (comments, processing instructions, declarations, CDATA). A series of examples demonstrates parsing behavior for valid tags with attributes, illegal tag names (which are not parsed as HTML but escaped), illegal attribute syntax, and whitespace rules. The chunk also covers hard line breaks within HTML contexts, noting that they are preserved in attribute values or rendered as `<br />` outside code spans or tags.

## Key Claims
- A double-quoted attribute value consists of a `"`, zero or more characters not including `"`, and a final `"`.
- An open tag includes `<`, a tag name, optional attributes/spaces/line endings, an optional `/`, and `>`.
- Illegal tag names (e.g., starting with a digit or containing invalid characters) are not parsed as HTML tags but escaped.
- Backslash escapes do not work in HTML attributes; they are preserved as literal backslashes.
- Hard line breaks (two spaces + newline or `\` + newline) inside HTML attribute values result in the literal string including the break, unlike outside code spans where they become `<br />`.

## Entities And Concepts
- Open tag
- Closing tag
- Empty element
- HTML comment
- Processing instruction
- Declaration
- CDATA section
- Hard line break
- Attribute value (double-quoted)
- Illegal tag name
- Backslash escape

## Procedures And API Details
N/A (Conceptual parsing rules).

## Nuance Or Contradictions
Hard line breaks behave differently inside code spans (preserved literally) versus outside them or within HTML attribute values (preserved literally in attributes, rendered as `<br />` in the block context if not inside a tag). Backslash escapes function in code spans but are ignored in HTML attributes.

## Candidate Wiki Hints
- `hard-line-breaks-in-html-attributes`
