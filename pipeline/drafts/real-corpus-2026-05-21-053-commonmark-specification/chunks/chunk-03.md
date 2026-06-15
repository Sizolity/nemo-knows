---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
- Source Section: **4.3 Setext headings** and **4.4 Indented code blocks**.
- Line Range: 1161–1978.
- Scope: Defines syntax for Setext-style headers (underline-based), rules for multiline content, indentation limits, precedence over other block types, compatibility notes regarding existing implementations, and the transition to indented code blocks.

Local Summary
This chunk details the mechanics of **Setext headings**, which rely on underlines (`=` or `-`) rather than the `#` syntax used in ATX headings. It clarifies that heading content can span multiple lines (unlike many legacy Markdown parsers), provided a blank line separates preceding paragraphs. The text also introduces **Indented code blocks**, defined as sequences of non-blank lines preceded by at least four spaces, and notes their precedence over list items when ambiguity exists.

Key Claims
- A Setext heading consists of one or more lines of text followed by an underline (`=` for level 1, `-` for level 2).
- The underline must not exceed three spaces of indentation; four spaces terminates the block as code or paragraph content.
- Multiline headings are supported if the text lines do not form other block constructs (e.g., lists, thematic breaks) before the underline.
- A blank line is required between a preceding paragraph and a following Setext heading to prevent the paragraph from becoming part of the heading's content.
- Indented code blocks require four or more spaces of indentation per line and treat content as literal text (no Markdown parsing).
- List item interpretations take precedence over indented code block interpretations if ambiguity exists regarding indentation.

Entities And Concepts
- **Setext Heading**: A header style using underlines (`=`, `-`).
- **Indented Code Block**: A code block formed by lines indented four or more spaces.
- **ATX Heading**: Implicitly referenced as the alternative heading syntax (using `#`).
- **Thematic Break**: A horizontal rule defined by three or more dashes, which can interrupt a potential Setext heading if not properly delimited.

Procedures And API Details
- **Heading Level Determination**: Check underline character (`=` -> Level 1, `-` -> Level 2).
- **Indentation Rule**:
    - Content lines: Up to 3 spaces indentation allowed.
    - Underlines: Up to 3 spaces indentation allowed; trailing spaces/tabs ignored for the underline itself but not for content.
    - Code blocks: Minimum 4 spaces indentation required.
- **Closing Condition**: A line with fewer than 4 spaces of indentation ends an indented code block immediately.

Nuance Or Contradictions
- **Multiline Compatibility**: The specification allows multiline headings (e.g., "Foo\nbar\n---\nbaz"), but notes that most existing Markdown implementations do not support this. These legacy tools often interpret such text as separate paragraphs or thematic breaks.
- **Lazy Continuation**: Setext underlines cannot function as lazy continuation lines within list items or block quotes; they must appear on their own structural level.
- **Internal Spaces in Underlines**: An underline cannot contain internal spaces or tabs (e.g., `= =` is invalid for a heading underline).

Candidate Wiki Hints
- **Topic: Setext Headings** – A page explaining the specific syntax of underlined headers, including multiline examples and indentation constraints.
- **Topic: Indented Code Blocks** – Documentation on how four-space indentation defines code blocks and their interaction with list items.
