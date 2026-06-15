---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---
Chunk Context
Section 3.c details the rules for parsing reference links, including full, collapsed, and shortcut variants. It covers link text grouping precedence over emphasis, HTML/code span/autolink precedence, label normalization (case folding, whitespace collapse), matching logic, and image syntax (6.4). Section 6.5 defines autolinks (URI and email) and Section 6.6 describes raw HTML tag parsing.

Local Summary
The chunk explains that links cannot nest other links at any level. It defines three types of reference links (full, collapsed, shortcut) with specific rules for label matching, normalization, and whitespace handling. Precedence rules determine whether text is treated as link text or inline content. Image syntax mirrors link syntax but allows embedded links in descriptions (though only plain text is used for `alt`). Autolinks are defined by `< >` wrappers around absolute URIs or email addresses. Raw HTML tags between `< >` are passed through without escaping.

Key Claims
- Links may not contain other links, at any level of nesting.
- Reference link labels must be normalized (case fold, strip/normalize whitespace) before matching.
- Matching is case-insensitive; consecutive internal spaces/tabs/line endings count as one space.
- Spaces/tabs/line endings are not allowed between the link text and link label for full/collapsed references.
- Shortcut reference links cannot be followed by `[]` or another link label.
- Full and collapsed references take precedence over shortcut references; inline links also take precedence.
- Image descriptions start with `!` and may contain links, but only plain string content is used for the `alt` attribute.
- Autolinks require `< >` around an absolute URI (scheme + colon + characters) or email address.
- Raw HTML tags between `< >` are rendered without escaping; tag names follow specific ASCII rules.

Entities And Concepts
- Reference Link: A link defined by a label matching a reference definition elsewhere in the document.
- Full Reference Link: `text[label]` where `label` matches a definition.
- Collapsed Reference Link: `text[]` where `text` matches a definition (equivalent to `text[text]`).
- Shortcut Reference Link: `text` where `text` matches a definition (not followed by `[]` or another label).
- Autolink: An absolute URI or email address wrapped in `< >`.
- Raw HTML Tag: Text between `< >` parsed as HTML without escaping.

Procedures And API Details
1. **Link Parsing Precedence**:
   - Link text grouping > emphasis grouping.
   - HTML tags, code spans, autolinks > link grouping.
2. **Label Normalization**:
   - Strip opening/closing brackets.
   - Perform Unicode case fold.
   - Strip leading/trailing whitespace (spaces, tabs, line endings).
   - Collapse consecutive internal whitespace to a single space.
3. **Matching Logic**:
   - Compare normalized forms of label and definition.
   - Use first matching definition if multiple exist.
4. **Image Syntax**:
   - `![description](url "title")` or `![description][]`.
   - Description may contain links, but `alt` uses plain text only.

Nuance Or Contradictions
- The spec departs from John Gruber’s original Markdown by disallowing whitespace between link text and label in reference links to prevent unintended shortcut link capture.
- Image descriptions allow nested links, but the rendered `alt` attribute ignores formatting (e.g., `<a>` tags inside description).

Candidate Wiki Hints
- Reference Link Types (Full, Collapsed, Shortcut)
- Link Label Normalization Rules
- Precedence of Inline Elements Over Links
- Autolink Syntax and Validation
