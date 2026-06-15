---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
This chunk details the CommonMark specification rules for backslash escaping, entity and numeric character references, and the structure of block-level elements. It covers thematic breaks (horizontal rules), ATX headings, and precedence rules between block and inline structures.

Local Summary
The document explains that backslashes escape only specific characters (like `*`, `#`) in certain contexts but not others (like code blocks). It defines valid HTML entity references and numeric character references (`&#` or `&#x`), noting their restrictions in structural elements like lists and headings. The chunk also introduces the concept of block structure (paragraphs, lists, headings) versus inline content, establishing that block indicators take precedence over inline ones. Specific rules for thematic breaks (`***`, `---`) and ATX headings (`# foo`) are provided, including requirements for spacing, indentation limits (3 spaces), and character matching.

Key Claims
- Backslash escapes do not work in code blocks, code spans, autolinks, or raw HTML.
- Valid HTML entity references and numeric character references can be used in place of Unicode characters, except in code contexts and structural elements like emphasis delimiters or list markers.
- Entity/numeric references cannot replace special characters defining structure (e.g., `&#42;` cannot create an emphasis bullet).
- Indicators of block structure always take precedence over indicators of inline structure.
- Thematic breaks require 3+ matching `-`, `_`, or `*` with optional indentation (up to 3 spaces).
- ATX headings use 1–6 unescaped `#` characters; more than six is not a heading.
- ATX headings must have at least one space between the opening `#` and content unless empty.

Entities And Concepts
- **Backslash Escaping**: Context-dependent escaping mechanism.
- **Entity References**: HTML5 named entities (e.g., `&copy;`).
- **Numeric Character References**: Decimal (`&#35;`) or hexadecimal (`&#x22;`).
- **Thematic Breaks**: Horizontal rules defined by sequences of `-`, `_`, or `*`.
- **ATX Headings**: Headings defined by leading/trailing `#` characters.
- **Block Structure**: Structural elements like paragraphs, lists, and headings.
- **Inline Content**: Text, links, emphasized text within blocks.

Procedures And API Details
- **Thematic Break Validation**: Line must contain 3+ matching characters (`-`, `_`, `*`) with optional spaces/tabs; no other characters allowed.
- **ATX Heading Parsing**:
  - Opening sequence: 1–6 unescaped `#`.
  - Closing sequence: Optional, any number of unescaped `#`.
  - Spacing: Spaces or tabs required after opening and before closing `#`.
  - Indentation: Up to 3 spaces allowed; 4+ spaces treats it as a code block.
- **Entity Parsing**: Check against HTML5 entity list; invalid codes replaced by U+FFFD.

Nuance Or Contradictions
- Original ATX implementation required a space after `#`, but the spec allows headings without it if strictly following the grammar, though many implementations still require it to avoid ambiguity (e.g., `#5 bolt`).
- HTML5 accepts entities without semicolons (e.g., `&copy`), but CommonMark requires the semicolon (`&copy;`) to avoid grammar ambiguity.

Candidate Wiki Hints
- **Thematic Breaks**: Rules for horizontal rules in Markdown.
- **ATX Headings**: Syntax and parsing rules for `#` based headings.
- **Backslash Escaping**: Limitations and valid contexts for backslashes.
- **Entity References**: Usage of HTML entities in Markdown text.
