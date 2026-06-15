---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
- **Heading**: `baz`
- **Line Range**: 1979–2982
- **Coverage**: Section 4.6 (HTML blocks) and the beginning of Section 4.7 (Link reference definitions).

Local Summary
This chunk details the CommonMark specification for handling HTML blocks. It defines seven types of HTML blocks based on start and end conditions, explaining how they interrupt or continue paragraphs. It covers specific behaviors for tags like `<pre>`, `<script>`, and comments, noting that certain blocks end at matching end tags rather than blank lines. The text also contrasts these rules with John Gruber's original Markdown syntax, highlighting differences regarding indentation and blank lines. The section concludes by introducing link reference definitions, outlining their structure (label, colon, destination, optional title) and parsing rules.

Key Claims
- HTML blocks are treated as raw HTML and are not escaped in output.
- There are seven kinds of HTML blocks defined by specific start and end conditions.
- Blocks of type 1–6 (e.g., `<pre>`, `<script>`, comments) end at the first line containing a corresponding end tag, allowing blank lines inside them.
- Blocks of type 7 (generic tags) end at the first blank line following the block.
- HTML blocks of types 1–6 can interrupt a paragraph; type 7 blocks cannot.
- Link reference definitions consist of a label, a colon, a destination, and an optional title.
- Link reference definitions do not correspond to structural elements but define labels for reference links.
- Matching of link labels is case-insensitive.
- A link reference definition cannot interrupt a paragraph.

Entities And Concepts
- **HTML Block**: A group of lines treated as raw HTML.
- **Start/End Conditions**: Rules determining when an HTML block begins and ends.
- **Type 1–6 Blocks**: Specific HTML constructs (pre, script, style, textarea, comments, processing instructions, declarations) ending at matching tags.
- **Type 7 Blocks**: Generic HTML blocks ending at a blank line.
- **Link Reference Definition**: A structural element defining a label for reference links.
- **Link Label**: The text used to identify a link.
- **Link Destination**: The URL or anchor target.
- **Link Title**: Optional text describing the link.

Procedures And API Details
- **HTML Block Parsing**:
  1. Check if a line meets a start condition (e.g., begins with `<pre`, `<!--`, `<?`).
  2. If matched, treat subsequent lines as raw HTML until the matching end condition is met.
  3. For types 1–5, look for the corresponding end tag (e.g., `</pre>`, `-->`).
  4. For type 6, look for a blank line.
  5. For type 7, look for a blank line.
- **Link Reference Definition Parsing**:
  1. Identify a line starting with a label followed by a colon (`:`).
  2. Parse the destination and optional title.
  3. Ensure no further characters occur after the title.
  4. Store the definition for use in reference links later in the document.

Nuance Or Contradictions
- **Indentation**: HTML blocks of types 1–6 can be preceded by up to three spaces of indentation, but not four. Type 7 blocks also allow indentation up to three spaces.
- **Blank Lines**: Unlike Gruber's original specification, CommonMark does not require blank lines before HTML blocks of types 1–6, nor does it allow blank lines inside type 7 blocks.
- **Tag Matching**: End tags need not match the start tag exactly (case-insensitive), but the content between them is treated as raw HTML.
- **Garbage In, Garbage Out**: Partial or invalid tags are passed through as-is if they start like a valid tag.

Candidate Wiki Hints
- **HTML Blocks in Markdown**: Explaining how to embed raw HTML safely and the differences between block types.
- **Link Reference Definitions**: A guide to defining custom links using labels and destinations.
- **CommonMark vs. Original Markdown**: Comparing the handling of HTML blocks and indentation rules.
