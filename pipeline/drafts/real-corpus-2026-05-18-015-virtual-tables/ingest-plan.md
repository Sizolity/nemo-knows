---
kind: topic
sources: [raw/web/corpus-2026-05-18/015-virtual-tables.md]
status: draft
---

# Ingest Plan

## Source Summary
- The document details the SQLite Virtual Table Mechanism, covering creation, implementation, and lifecycle management.
- It defines the C API (`sqlite3_module`) methods required for custom virtual table development (e.g., `xCreate`, `xBestIndex`).
- Security considerations include preventing bypasses of schema trust settings and handling shadow tables defensively.

## Candidate Wiki Pages
- wiki/sources/virtual-tables.md — To store the raw corpus metadata, fetch status, and specific URL references from the source file.
- wiki/concepts/sqlite-vtab-methods.md — To document the specific C function pointers (`xCreate`, `xOpen`, etc.) required for virtual table implementations.
- wiki/topics/custom-virtual-table-design.md — To synthesize the usage patterns, eponymous tables, and security constraints into a design guide.

## Suggested Links
- https://www.sqlite.org/vtab.html

## Review Checklist
- [ ] Verify that all C struct definitions are accurately transcribed to the `wiki/concepts/` directory.
- [ ] Ensure security warnings regarding `SQLITE_VTAB_DIRECTONLY` are highlighted in the design guide.
