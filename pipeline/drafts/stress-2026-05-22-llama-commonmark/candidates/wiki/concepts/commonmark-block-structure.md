---
title: Commonmark Block Structure
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Block Structure

The CommonMark Specification defines a rigorous, unambiguous syntax for Markdown, authored by John MacFarlane to resolve historical ambiguities in the original description. It ensures consistent rendering across different implementations by establishing a two-phase parsing strategy: first constructing the block structure, then parsing raw text contents into inline elements.

The specification covers the entire lifecycle of parsing, from character-level preprocessing to the final tree structure. It begins by defining fundamental units like characters, lines, and tabs, and establishes rules for escaping special characters using backslashes and parsing HTML entities.

The core of the document details block-level structures, including thematic breaks, ATX and Setext headings, indented and fenced code blocks, HTML blocks, and link reference definitions. A significant portion is dedicated to container blocks, specifically block quotes and list items. It defines strict rules for indentation, distinguishing between list continuations, paragraph continuations, and indented code blocks using a "four-space rule." The specification also rigorously defines "loose" versus "tight" lists and handles the "laziness" of markers in block quotes and lists.

Inline parsing logic is detailed with a focus on precedence, where code spans, links, and HTML tags interrupt or take precedence over emphasis markers. Specific rules govern the matching of emphasis delimiters (`*`, `_`, `**`, `__`), the normalization of link labels, and the strict syntax requirements for raw HTML tags. The document concludes with rules for soft line breaks, textual content preservation, and the complex algorithmic logic of the delimiter stack used to resolve nested emphasis and links.

## Key Claims

- **Readability and Unambiguity:** Markdown source text must remain readable as plain text, and the specification aims to eliminate ambiguity to ensure consistent behavior across all compliant renderers.
- **Escaping Mechanisms:** Backslashes escape ASCII punctuation characters in most contexts but are invalid within code blocks, code spans, autolinks, and raw HTML.
- **Indentation Hierarchy:** Content under a list item must be indented four spaces from the margin to be treated as part of the list item; three spaces or fewer continue the paragraph. Four spaces after a blank line create a code block.
- **HTML Block Termination:** Unlike the original Markdown spec, CommonMark HTML blocks (types 1–6) end at the matching end tag rather than requiring a blank line to terminate.
- **List Tightness:** A list is considered "loose" if items are separated by blank lines or if an item contains two block-level elements separated by a blank line; otherwise, it is "tight."
- **Emphasis Precedence:** Inline elements like code spans, links, and HTML tags interrupt or take precedence over emphasis markers. Links bind more tightly than brackets, which bind more tightly than emphasis.
- **Delimiter Matching:** Emphasis requires matching delimiter runs that are left-flanking (to open) and right-flanking (to close). Intraword emphasis is generally forbidden for underscores but allowed for asterisks under specific conditions.
- **Soft Line Breaks:** Regular line endings not preceded by spaces or backslashes are parsed as soft breaks, allowing renderers to choose between rendering them as a line ending or a space.
- **Autolink Validity:** The spec accepts strings like `a+b+c:d` as autolinks, which may not be valid URIs under strict RFC standards, and backslash escapes do not function inside autolinks.

## Related Concepts

- [[commonmark-inline-elements]]
- [[commonmark-escaping-and-entities]]
- [[commonmark-lists]]
- [[commonmark-line-breaks]]
