---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
This chunk covers the `baz` heading and details CommonMark rules for code blocks (info strings, tilde vs backtick fences) and HTML blocks. It defines seven types of HTML block start/end conditions, illustrates how they interrupt or contain paragraphs, and contrasts these rules with John Gruber’s original Markdown syntax regarding blank lines and indentation. The section concludes by introducing link reference definitions.

Local Summary
The specification clarifies that info strings in backtick code fences cannot contain backticks, whereas tilde fences can. HTML blocks are defined by seven distinct start/end patterns (e.g., `<pre>...</pre>`, comments `<!-- -->`, or generic tags). Unlike Gruber’s original syntax which required blank lines around block-level HTML and forbade indentation, this spec allows indented HTML blocks to interrupt paragraphs (except for type 7) and treats content inside tags as raw HTML until a matching end tag or blank line terminates the block.

Key Claims
- Info strings for backtick code fences cannot contain backticks; tilde fences allow them.
- There are seven kinds of HTML block defined by specific start/end conditions.
- HTML blocks continue until an appropriate end condition, the document end, or a container boundary.
- HTML blocks of types 1–6 may interrupt a paragraph without preceding blank lines (except at document end).
- Blocks of type 7 cannot interrupt a paragraph.
- This spec disallows blank lines inside HTML blocks to avoid expensive balanced tag parsing and allow Markdown content insertion via blank lines.
- Link reference definitions consist of a label, colon, destination, and optional title; they do not correspond to structural elements.

Entities And Concepts
- Code block info strings (language specification)
- Backtick code fences (` ``` `)
- Tilde code fences (` ~~~ `)
- HTML blocks (types 1–7)
- Block-level HTML tags (e.g., `<div>`, `<table>`)
- Inline HTML tags (e.g., `<del>`)
- Link reference definitions
- Raw HTML vs. escaped HTML

Procedures And API Details
- **HTML Block Detection**:
  - Type 1: Starts with `<pre`, `<script`, `<style`, or `<textarea` followed by space/tab/`>`/EOL; ends at matching `</pre>`, etc.
  - Type 2: Starts with `<!--`; ends at `-->`.
  - Type 3: Starts with `<?`; ends at `?>`.
  - Type 4: Starts with `<!` followed by ASCII letter; ends at `>`.
  - Type 5: Starts with `<![CDATA[`; ends at `]]>`.
  - Type 6: Starts with block-level tags (e.g., `<div>`, `<table>`); ends at blank line.
  - Type 7: Starts with any complete open/closing tag; ends at blank line.
- **Link Reference Definition Syntax**:
  - Label + `:` + Destination + Optional Title.
  - Matching is case-insensitive.
  - Definitions can precede or follow usage.

Nuance Or Contradictions
- Gruber’s original Markdown requires blank lines around block-level HTML and forbids indentation; this spec relaxes indentation rules but removes blank line flexibility inside blocks to simplify parsing.
- In type 1–6 HTML blocks, a partial tag on the first line is valid if split where whitespace would occur (e.g., `<div id="foo"`).

Candidate Wiki Hints
- Page: CommonMark HTML Blocks
  - Explain the seven types of HTML block start/end conditions.
  - Discuss how HTML blocks interact with paragraphs and code blocks.
  - Compare with Gruber’s original Markdown restrictions.
