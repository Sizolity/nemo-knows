---
kind: topic
sources: [raw/web/corpus-2026-05-18/017-json-functions-and-operators.md]
status: draft
---

# Ingest Plan

## Source Summary
- **Topic**: Comprehensive documentation of SQLite's JSON functions and operators (introduced in v3.38.0).
- **Key Content**: Covers `json()` vs `jsonb()` handling, path arguments (`$`, `[#-N]`), value argument rules, aggregate functions, table-valued decomposers (`json_each`, `json_tree`), and specific update functions (`json_set`, `json_replace`, `json_insert`).
- **Technical Depth**: Details on JSON5 support, performance optimization strategies (text vs binary JSONB input), the legacy BLOB input bug fix in v3.45.1, and compatibility notes with MySQL/PostgreSQL.

## Candidate Wiki Pages
- wiki/sources/sqlite-json-functions.md — To store metadata about this specific corpus item (URL, date, confidence) and link to the canonical SQLite documentation.
- wiki/concepts/json-blob-vs-text-storage.md — To explain the internal `JSONB` binary format, its performance benefits over text JSON, and the implications of the legacy BLOB input bug.
- wiki/topics/sqlite-json-update-strategies.md — To compare and contrast `json_set`, `json_replace`, `json_insert`, and `json_patch` for different data modification scenarios (overwrite vs create).

## Suggested Links
- https://www.sqlite.org/json1.html

## Review Checklist
- [ ] Verify all function signatures and examples match the current SQLite documentation.
- [ ] Ensure distinction between `json_` (text) and `jsonb_` (binary) variants is clearly highlighted in candidate pages.
- [ ] Confirm that JSON5 extension support details are included in the concept notes.
