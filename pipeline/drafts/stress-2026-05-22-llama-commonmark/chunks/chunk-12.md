---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
This chunk (lines 7515-7801) covers section 3.c of the CommonMark specification, detailing the parsing rules for HTML tags, attributes, comments, processing instructions, declarations, and CDATA sections. It includes numerous examples (613–646) demonstrating valid and invalid syntax for open tags, closing tags, whitespace handling, illegal characters, and the specific behavior of hard line breaks within HTML contexts.

Local Summary
The specification defines the structural components of HTML tags, including open tags, closing tags, and various embedded content types like comments and CDATA sections. It provides concrete examples of valid tag syntax, including custom tag names and attributes with various quote styles and values. The text explicitly lists illegal tag names, illegal attribute names, and illegal attribute values, showing how the parser escapes these in the output. A significant portion is dedicated to "Hard line breaks" (section 6.7), explaining how line endings preceded by spaces or backslashes are rendered as `<br />` tags, with specific constraints regarding where these breaks can occur (e.g., not inside code spans or at the end of a block).

Key Claims
- A double-quoted attribute value consists of a starting `"`, zero or more characters not including `"`, and a final `"`.
- An open tag consists of `<`, a tag name, optional attributes/spaces, an optional `/`, and `>`.
- An HTML comment consists of `<!--`, a string not including `-->`, and `-->`.
- A processing instruction consists of `<?`, a string not including `?>`, and `?>`.
- A declaration consists of `<!`, an ASCII letter, zero or more characters not including `>`, and `>`.
- A CDATA section consists of `<![CDATA[`, a string not including `]]>`, and `]]>`.
- An HTML tag is defined as an open tag, a closing tag, an HTML comment, a processing instruction, a declaration, or a CDATA section.
- Hard line breaks (two or more spaces or a backslash before a line ending) are parsed as `<br />` tags, except inside code spans or at the end of a block.
- Backslash escapes do not work in HTML attributes.
- Entity and numeric character references are preserved in HTML attributes.

Entities And Concepts
- Open tag
- Closing tag
- HTML comment
- Processing instruction
- Declaration
- CDATA section
- HTML tag
- Hard line break
- Entity reference
- Numeric character reference
- Attribute value
- Tag name

Procedures And API Details
- To form a valid open tag: Start with `<`, followed by a tag name, then zero or more attributes, optional spaces/tabs/line endings, an optional `/`, and end with `>`.
- To form a valid closing tag: Start with `</`, followed by a tag name, optional spaces/tabs/line endings, and end with `>`.
- To form a hard line break: Place two or more spaces or a backslash before a line ending (not in a code span or HTML tag, and not at the end of a block).
- To form an HTML comment: Use `<!--`, insert content not containing `-->`, and close with `-->`.
- To form a processing instruction: Use `<?`, insert content not containing `?>`, and close with `?>`.
- To form a declaration: Use `<!`, an ASCII letter, insert content not containing `>`, and close with `>`.
- To form a CDATA section: Use `<![CDATA[`, insert content not containing `]]>`, and close with `]]>`.

Nuance Or Contradictions
- The specification notes that illegal tag names (e.g., `<33>`, `<__>`) are not parsed as HTML tags but are escaped in the output.
- Illegal attribute names (e.g., `h*#ref`) and illegal attribute values (e.g., missing closing quote) result in escaped output.
- Illegal whitespace (e.g., space after `<`, space before `>`) causes the tag to be treated as text.
- Missing whitespace between attributes (e.g., `href='bar'title=title`) is considered illegal and results in escaped output.
- Hard line breaks do not occur inside code spans or HTML tags, nor at the end of a paragraph or other block element.
- Backslash escapes are explicitly stated to not work in HTML attributes.

Candidate Wiki Hints
- HTML Tag Syntax in CommonMark
- Hard Line Breaks in Markdown
- HTML Comments and CDATA in CommonMark
- Illegal Characters in HTML Tags
