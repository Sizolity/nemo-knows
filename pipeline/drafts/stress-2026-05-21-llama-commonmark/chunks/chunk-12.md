---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Section 3. c: HTML Tags and Attributes
- Lines 7515-7801
- Covers definitions of HTML components (open/closing tags, comments, processing instructions, etc.) and a series of examples (613–646) demonstrating valid and invalid HTML parsing within CommonMark.

Local Summary
This chunk defines the structural components of an HTML tag in CommonMark, including open tags, closing tags, comments, processing instructions, declarations, and CDATA sections. It provides numerous examples illustrating valid tag syntax, whitespace handling, attribute formatting, and illegal characters that result in HTML entities. It also introduces the concept of hard line breaks in section 6.7, explaining how line endings preceded by spaces or backslashes are rendered as `<br />` tags, with specific rules regarding their placement within inline content versus block elements.

Key Claims
- A double-quoted attribute value consists of a starting `"`, zero or more characters not including `"`, and a final `"`.
- An open tag consists of `<`, a tag name, optional attributes/spaces, an optional `/`, and `>`.
- A closing tag consists of `</`, a tag name, optional spaces, and `>`.
- HTML comments start with `<!--` and end with `-->`.
- Processing instructions start with `<?` and end with `?>`.
- Declarations start with `<!`, followed by an ASCII letter, and end with `>`.
- CDATA sections start with `<![CDATA[` and end with `]]>`.
- Hard line breaks occur when a line ending is preceded by two or more spaces or a backslash, rendering as `<br />`.
- Hard line breaks do not occur inside code spans or at the end of a block element.

Entities And Concepts
- Open Tag
- Closing Tag
- HTML Comment
- Processing Instruction
- Declaration
- CDATA Section
- Hard Line Break
- Attribute Value
- Tag Name
- Illegal Tag Names
- Illegal Attribute Names
- Illegal Attribute Values
- Illegal Whitespace

Procedures And API Details
- To create a valid open tag: `<tagname attr="value">`
- To create a valid closing tag: `</tagname>`
- To create a valid HTML comment: `<!-- comment -->`
- To create a valid processing instruction: `<?instruction?>`
- To create a valid declaration: `<!ELEMENT br EMPTY>`
- To create a valid CDATA section: `<![CDATA[content]]>`
- To create a hard line break: Use two or more spaces or a backslash before a line ending.

Nuance Or Contradictions
- Backslash escapes do not work in HTML attributes (Example 631).
- Illegal tag names (e.g., `<33>`) are not parsed as HTML but escaped (Example 618).
- Illegal attribute names (e.g., `h*#ref`) result in escaped output (Example 619).
- Illegal attribute values (e.g., unescaped quotes) result in escaped output (Example 620).
- Illegal whitespace (e.g., spaces inside tag names) results in escaped output (Example 621).
- Missing whitespace between attributes (e.g., `href='bar'title=title`) results in escaped output (Example 622).
- Hard line breaks do not occur inside code spans (Example 640) or at the end of a block element (Example 644).

Candidate Wiki Hints
- Hard Line Breaks in Markdown
- HTML Tag Syntax in CommonMark
- CommonMark HTML Parsing Rules
- Attribute Value Quoting
- Illegal HTML Characters in CommonMark
