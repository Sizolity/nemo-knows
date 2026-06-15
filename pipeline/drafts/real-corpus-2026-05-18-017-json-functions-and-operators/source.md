---
title: JSON Functions and Operators Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/017-json-functions-and-operators.md
confidence: medium
---

# JSON Functions and Operators Summary

## What It Is
A comprehensive documentation of the built-in JSON functions and operators available in SQLite. These tools allow for parsing, generating, extracting, modifying, and validating JSON data within SQL queries. The implementation supports both standard RFC-8259 JSON text and a binary internal format known as JSONB for performance optimization.

## Summary
Since version 3.38.0, SQLite includes thirty scalar functions and two operators by default to handle JSON values, plus four table-valued functions for decomposing JSON strings. The library supports standard canonical JSON syntax and reads input containing JSON5 extensions (added in v3.42.0). A binary "JSONB" format was introduced in v3.45.0 to store the internal parse tree on disk, allowing faster read/update operations by bypassing the parser. While SQLite stores JSON as ordinary text with a limitation to NULL, integers, floats, text, and BLOBs (no dedicated JSON type), it offers a robust set of functions for manipulation, including object creation (`json_object`), extraction (`json_extract`, `->`, `->>`), modification (`json_insert`, `json_replace`, `json_set`), and aggregation.

## Key Claims
- **Default Availability**: JSON support is built-in by default starting from SQLite version 3.38.0 (2022-02-22). Prior versions required the `-DSQLITE_ENABLE_JSON1` compile-time option.
- **Storage Format**: SQLite stores JSON as ordinary text, supporting NULL, integers, floating-point numbers, text, and BLOBs. A new "JSON" type cannot be added due to backwards compatibility constraints.
- **JSONB Performance**: Starting in v3.45.0 (2024-01-15), SQLite uses a binary JSONB format internally. Functions prefixed with `jsonb_` return this binary format, enabling faster execution when inputs are already in JSONB format.
- **JSON5 Support**: Beginning in version 3.42.0 (2023-05-16), SQLite routines can read and interpret JSON5 extensions (e.g., unquoted keys, trailing commas, single quotes) but always output strict canonical JSON.
- **Nesting Limit**: Due to the recursive descent parser used to avoid excess stack space, JSON inputs with more than 1000 levels of nesting are considered invalid.
- **Path Syntax**: Path arguments must begin with `$` followed by `.objectlabel` or `[arrayindex]`. Array indices can be negative (e.g., `#-N`) relative to the end of an array, supported from v3.47.0 onwards.
- **Operators**: The `->` operator returns a JSON representation of a subcomponent, while `->>` returns an SQL scalar value (TEXT, INTEGER, REAL, or NULL).

## Suggested Links
https://www.sqlite.org/json1.html
