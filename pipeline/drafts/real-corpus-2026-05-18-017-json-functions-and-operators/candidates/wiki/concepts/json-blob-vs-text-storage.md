---
title: Json Blob Vs Text Storage
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/017-json-functions-and-operators.md
confidence: medium
---

# Json Blob Vs Text Storage

In SQLite, the handling of JSON data involves a choice between standard text storage and a specialized binary format designed for performance. While SQLite lacks a dedicated `JSON` column type due to backwards compatibility constraints, it stores JSON as ordinary text by default. This text storage supports NULL, integers, floating-point numbers, text, and BLOBs.

Starting with version 3.45.0, SQLite introduced a binary internal format known as **JSONB** (or `jsonb`) to optimize performance. The system utilizes a recursive descent parser that avoids excessive stack space but enforces a nesting limit of 1000 levels for standard inputs.

## Storage Formats

### Text Storage
Historically and by default, SQLite stores JSON documents as text strings within columns.
- **Input**: Supports standard RFC-8259 JSON text and JSON5 extensions (unquoted keys, trailing commas) added in v3.42.0.
- **Output**: All functions output strict canonical JSON regardless of input style.
- **Parsing**: Requires parsing the text string to extract values or modify structure.

### Binary Storage (JSONB)
Introduced to bypass the text parser during read and update operations, this format stores the internal parse tree on disk.
- **Availability**: Available as a distinct binary format starting in v3.45.0.
- **Functions**: Functions prefixed with `jsonb_` operate on this binary format.
- **Performance**: Significantly faster execution when inputs are already in JSONB format, avoiding the overhead of text parsing and canonicalization.

## Operators

SQLite provides specific operators to navigate these structures:
- `$` followed by `.objectlabel` or `[arrayindex]`: Used for path arguments.
- `->`: Returns a JSON representation of a subcomponent.
- `->>`: Returns an SQL scalar value (TEXT, INTEGER, REAL, or NULL).

## References

For detailed documentation on these functions and operators, refer to the official SQLite JSON1 documentation.
