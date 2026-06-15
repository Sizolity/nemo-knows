---
title: Database File Format Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/020-database-file-format.md
confidence: medium
---

# Database File Format

## What It Is
The **Database File Format** describes the on-disk structure used by SQLite versions 3.0.0 and later. The complete state of an SQLite database is typically contained in a single file called the "main database file," which consists of one or more pages. During transactions, additional information may be stored in a second file known as the "rollback journal" or, if WAL mode is active, a "write-ahead log" (WAL) file.

## Summary
SQLite databases are organized into fixed-size pages (powers of two between 512 and 65536 bytes). The main database file begins with a 100-byte header containing metadata such as the magic string ("SQLite format 3"), page size, file change counter, text encoding, and schema format number. Internally, data is stored using b-tree structures (for tables and indexes), which utilize interior pages, leaf pages, freelist pages for unused space, and overflow pages for large payloads. The document details the binary layout of these pages, including cell formats, varint encodings for serial types, and pointer maps used in vacuum modes. It also covers the schema layer stored in the `sqlite_schema` table (also known as `sqlite_master`) and the mechanisms for transaction management via rollback journals and WAL files.

## Key Claims
- **Single File Architecture:** The main database file contains the complete state of the database, optionally supplemented by a rollback journal or write-ahead log during transactions.
- **Page Structure:** All pages within a database are the same size; reads and writes generally occur at page boundaries. Page 1 always contains the 100-byte database header and is the root of the schema b-tree.
- **B-tree Storage:** Data is stored in b-trees which can be table b-trees (storing data in leaves) or index b-trees (storing only keys). WITHOUT ROWID tables use index b-trees where the key includes primary key columns.
- **Schema Representation:** The database schema is stored in a special table named `sqlite_schema` (aliases include `sqlite_master`, `sqlite_temp_schema`, `sqlite_temp_master`) located on page 1.
- **WAL Mode:** In Write-Ahead Log mode, changes are written to a separate `-wal` file rather than the main database file immediately. Checkpoints transfer WAL content back to the main database file.
- **Rollback Journals:** A rollback journal (named with `-journal` appended) holds original page contents before modification to allow for atomic transaction commits and crash recovery.

## Suggested Links
*   https://www.sqlite.org/fileformat2.html
