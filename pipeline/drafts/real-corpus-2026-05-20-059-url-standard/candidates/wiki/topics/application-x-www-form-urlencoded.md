---
title: Application X Www Form Urlencoded
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

# Application X Www Form Urlencoded

The `application/x-www-form-urlencoded` MIME type is a standard format used to represent form data in network identifiers. Defined within the **URL Standard**, this specification outlines strict parsing and serialization algorithms that differ from those used for general URL components.

### Parsing Rules
When parsing content encoded as `application/x-www-form-urlencoded`, the process follows these steps:
1.  The input string is split on the `&` character to separate key-value pairs.
2.  Within each pair, a `+` character is replaced with a literal space.
3.  The resulting sequences are percent-decoded into valid UTF-8 characters without a BOM.

### Serialization Rules
Conversely, serializing data for this MIME type requires adherence to specific encoding sets:
*   Spaces may be encoded as either `%20` or `+`, depending on the context.
*   Special characters are percent-encoded using the form set rules.
*   Key-value pairs are joined using `=` and `&` delimiters.

This format is distinct from general URL paths, which utilize different encoding sets. The standard emphasizes that only specific encoding sets result in roundtripable data; other sets require careful handling of the literal `%` character to prevent corruption.
