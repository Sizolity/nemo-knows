---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
- Heading path: 3. c
- Line range: 6637-7513
- Section: Reference links, image syntax, and autolinks.

Local Summary
This chunk details the parsing rules for reference links (full, collapsed, shortcut), image syntax, and autolinks. It establishes precedence rules where link text grouping overrides emphasis, HTML tags, code spans, and autolinks. It defines the normalization process for link labels (case folding, whitespace collapsing) and specifies constraints such as the 999-character limit for link labels and the prohibition of unescaped brackets within labels. The section concludes with the grammar for raw HTML tags.

Key Claims
- Links cannot contain other links at any level of nesting.
- Link text grouping takes precedence over emphasis grouping.
- HTML tags, code spans, and autolinks take precedence over link grouping.
- Link labels must be normalized (case fold, strip whitespace) for matching.
- Matching is case-insensitive.
- Reference links are case-insensitive.
- Spaces, tabs, or line endings are not allowed between the link text and the link label.
- Image descriptions may contain links, but only the plain string content is used for the `alt` attribute.
- Autolinks consist of absolute URIs or email addresses enclosed in `< >`.
- Raw HTML text between `<` and `>` is parsed as tags without escaping.

Entities And Concepts
- Reference links (full, collapsed, shortcut)
- Link label
- Link text
- Image description
- Autolink (URI autolink, email autolink)
- Raw HTML tag
- Normalization (Unicode case fold, whitespace collapsing)
- Scheme (for URIs)

Procedures And API Details
- **Link Label Normalization**: Strip opening/closing brackets, perform Unicode case fold, strip leading/trailing whitespace, collapse internal whitespace to a single space.
- **Link Matching**: Compare normalized forms of the link text and label. Use the first matching definition if multiple exist.
- **Image Alt Attribute**: Extract only the plain string content of the image description; ignore inline formatting.
- **Autolink Parsing**: Match `<` followed by an absolute URI or email address followed by `>`.
- **Raw HTML Grammar**:
  - Tag name: ASCII letter followed by letters, digits, or hyphens.
  - Attribute name: ASCII letter, `_`, or `:` followed by letters, digits, `_`, `.`, `:`, or `-`.
  - Attribute value: Unquoted, single-quoted, or double-quoted.

Nuance Or Contradictions
- **Whitespace in Reference Links**: This spec forbids whitespace between link text and label, departing from John Gruber’s original Markdown syntax which allowed it. This change prevents unintended capture of consecutive shortcut reference links.
- **Image Formatting**: While image descriptions can contain links, the rendered `alt` attribute uses only the plain text, stripping all formatting.
- **Autolink Validity**: Many strings parsed as autolinks (e.g., `a+b+c:d`) are not valid URIs according to strict standards but are accepted by this spec.
- **Escaping in Autolinks**: Backslash-escapes do not work inside autolinks; they are treated as literal characters.

Candidate Wiki Hints
- **Reference Link Precedence**: How full and collapsed references take precedence over shortcut references.
- **Image Alt Text Rules**: Why only plain text is used for `alt` attributes in Markdown images.
- **Autolink Syntax**: Rules for valid URI and email autolinks, including scheme requirements.
- **Raw HTML Parsing**: Grammar for tags and attributes in the CommonMark spec.
