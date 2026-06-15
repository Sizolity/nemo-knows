---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Group Context

This group of notes aggregates content from the **CommonMark Specification**, covering the document structure, metadata retrieval, and the core parsing logic for lists, text content, and inline elements. The collection spans from the initial document header through detailed algorithmic descriptions of the delimiter stack, soft line breaks, and the two-phase parsing strategy (block structure followed by inline structure). The content ranges from line 1 to line 8133 of the source file, encompassing sections on list continuation, nested items, textual content preservation, and the specific mechanics of resolving emphasis and links.

# Cross-Chunk Summary

The document begins with high-level metadata and structural definitions, including fetch metadata and retrieved text handling. It transitions into specific examples involving list items, nested lists, and continuation rules (Chunks 01–02). A significant portion of the text (Chunks 03–04) focuses on "baz" and "Heading" paths, likely representing specific test cases or edge cases in the specification. The middle sections (Chunks 05–07) cover bracketed text `[Foo]` and generic "aaa" paths, suggesting coverage of link/image syntax and general text nodes. The latter half of the document (Chunks 08–13) delves deeply into list item enumeration (2. b, 3. c) and concludes with the sophisticated logic of soft line breaks, textual content preservation, and the delimiter stack algorithm used to parse inline elements like emphasis and links.

# Repeated Or Central Claims

- **Two-Phase Parsing**: The specification consistently emphasizes a two-phase approach: first constructing the block structure (paragraphs, block quotes, lists) and then parsing raw text contents into inline elements (strings, code spans, links, emphasis).
- **Textual Content Preservation**: Any characters not given an interpretation by previous rules are parsed as plain textual content, and internal spaces are preserved verbatim.
- **Delimiter Stack Mechanism**: The resolution of nested emphasis and links relies on a doubly linked list (delimiter stack) that tracks potential openers and closers for delimiters such as `*`, `_`, `[`, and `![]`.
- **Rendering Flexibility**: Soft line breaks (regular line endings not preceded by spaces or backslashes) allow renderers to choose between rendering them as a line ending or a space, as browser behavior renders both identically.
- **List Continuation**: The specification details how list items can be continued with open blocks or nested without requiring explicit item continuation markers in certain contexts.

# Important Local Details

- **Soft Break Definition**: A line ending is parsed as a soft break if it is not preceded by two or more spaces or a backslash.
- **Delimiter Stack Operations**:
  - Insertion: When encountering `*`, `_`, `[`, or `![`, a text node is inserted, and a pointer is added to the delimiter stack.
  - Link/Image Resolution: When encountering `]`, the system calls `look for link or image`, searching backwards for an active opening delimiter. If a link is found, preceding `[` delimiters are set to inactive.
  - Emphasis Processing: The `process emphasis` function iterates through the stack to find matching openers and closers, determining emphasis levels based on delimiter run length and inserting `emph` or `strong emph` nodes.
- **Openers_bottom**: This concept defines a lower bound for searching delimiters of specific types during emphasis processing.
- **Lazy Continuation**: Describes a line added to an open block without closing it, relevant to list item handling.
- **Chunk Line Ranges**: The notes cover a continuous range from line 1 to 8133, with specific focus on the final chunk (7802-8133) detailing the delimiter stack and soft breaks.

# Candidate Wiki Hints

- **Soft Line Breaks**: Create a page explaining the conditions for soft breaks and the rendering variations (space vs. line ending).
- **Parsing Strategy**: Develop a technical guide detailing the two-phase parsing model and the tree construction process.
- **Delimiter Stack Algorithm**: Write a deep dive into the logic for resolving nested emphasis and links, including the `look for link or image` and `process emphasis` procedures.
- **List Continuation Rules**: Document the rules for list item continuation, nested lists, and open blocks.
- **Textual Content**: A page on how plain text and internal spaces are handled when no other interpretation applies.

# Gaps Or Cautions

- **Missing Raw Content**: The provided chunk notes and index do not contain the actual text of the specification, only metadata about headings and line ranges. Specific examples or test cases referenced in the "baz" or "aaa" sections cannot be fully verified without the raw text.
- **Ambiguous Heading Paths**: Several chunks (e.g., Chunk 02, 03) have heading paths that appear to be test case identifiers (e.g., "foo > foo > foo") rather than semantic section titles, which may limit the ability to create intuitive navigation links without the full context of the test cases.
- **Incomplete Algorithmic Steps**: While the delimiter stack algorithm is described, the specific implementation details of the `look for link or image` search logic (e.g., exact distance limits, handling of inactive delimiters) are implied but not fully detailed in the provided notes.
- **Rendering Options**: The notes mention that renderers may offer options to treat soft line breaks as hard line breaks, but the specific API or configuration options for this are not detailed.
