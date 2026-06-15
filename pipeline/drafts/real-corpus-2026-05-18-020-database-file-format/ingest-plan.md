---
kind: topic
sources: [raw/web/corpus-2026-05-18/020-database-file-format.md]
status: draft
---

# Ingest Plan

## Source Summary
- The document defines the on-disk file format for SQLite databases since version 3.0.0.
- It details the structure of the main database file, rollback journals, and write-ahead logs (WAL).
- Coverage includes page structures, B-tree algorithms, schema storage, record formats, and internal statistics tables.
- The source explicitly covers low-level byte-level specifications for headers, cells, and overflow pages.

## Candidate Wiki Pages
- wiki/sources/020-database-file-format.md — To store the raw metadata and ingestion status of this specific corpus item.
- wiki/concepts/sqlite-page-structure.md — To document the 100-byte header, page size limits, and B-tree cell layouts defined in sections 1.3 and 1.6.
- wiki/topics/sqlite-journal-modes.md — To summarize the rollback journal and WAL file formats described in sections 3 and 4.

## Suggested Links
- https://www.sqlite.org/fileformat2.html

## Review Checklist
- [ ] Verify that all internal schema object names (sqlite_schema, sqlite_sequence) match the source definitions.
- [ ] Ensure byte-order conventions (big-endian vs native) for the WAL-index are clearly distinguished from other formats.
