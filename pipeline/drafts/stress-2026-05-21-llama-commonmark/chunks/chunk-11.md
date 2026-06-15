---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

Chunk Context
- Section 3. c: Link syntax and reference link definitions.
- Section 6.4: Images syntax.
- Section 6.5: Autolinks syntax.
- Section 6.6: Raw HTML syntax.

Local Summary
This chunk details the parsing rules for reference links (full, collapsed, shortcut), image syntax, autolinks (URI and email), and raw HTML tags. It clarifies precedence rules between link types, normalization for matching, and specific constraints on content within link labels and descriptions.

Key Claims
- Links cannot contain other links at any nesting level.
- Reference link matching is case-insensitive and uses Unicode case folding.
- Link labels must not contain unescaped square brackets.
- Image descriptions may contain links, but only the plain string content is used for the `alt` attribute.
- Autolinks are absolute URIs or email addresses enclosed in `< >`.
- Raw HTML tags are parsed without escaping if they match the grammar.

Entities And Concepts
- Full reference link
- Collapsed reference link
- Shortcut reference link
- Image description
- URI autolink
- Email autolink
- Raw HTML tag
- Unicode case fold

Procedures And API Details
- **Link Label Normalization**: Strip brackets, perform Unicode case fold, strip leading/trailing whitespace, collapse internal whitespace to a single space.
- **Matching**: Compare normalized forms of link labels.
- **Image Alt Attribute**: Extract plain string content from the image description; ignore inline formatting.
- **Autolink Parsing**: Detect `<scheme:...>` or `<email>` patterns; validate scheme format (2–32 chars, ASCII letter start).

Nuance Or Contradictions
- **Whitespace in Reference Links**: Unlike original Markdown, this spec forbids whitespace between link text and label to prevent unintended shortcut link capture.
- **Image Links**: While image descriptions can contain links syntactically, the rendering recommendation is to strip formatting for the `alt` text.
- **Autolink Validity**: Many strings parsed as autolinks (e.g., `m:abc`) are not valid URIs per standard registries but are accepted by this spec.

Candidate Wiki Hints
- Reference Link Precedence
- Image Alt Text Handling
- Autolink Scheme Validation
- Raw HTML Tag Grammar
