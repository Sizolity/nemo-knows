---
title: Sqlite Journal Modes
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/020-database-file-format.md
confidence: medium
---

# Sqlite Journal Modes

SQLite manages data integrity and crash recovery through specific journaling mechanisms. The on-disk structure typically consists of a single "main database file" containing one or more pages, optionally supplemented by auxiliary files during transactions.

## Transactional Storage

SQLite supports two primary methods for handling transactional changes:

- **Rollback Journal**: A separate file (named with `-journal` appended) stores the original page contents before modification. This allows for atomic transaction commits and crash recovery by reverting to the pre-transaction state if necessary.
- **Write-Ahead Log (WAL)**: In WAL mode, changes are written to a separate `-wal` file rather than modifying the main database file immediately. Checkpoints are used to transfer content from the WAL back into the main database file.

## Page Architecture

The main database file is organized into fixed-size pages, with sizes being powers of two between 512 and 65536 bytes. Reads and writes generally occur at page boundaries. The internal layout utilizes b-tree structures for storing data in table b-trees or keys in index b-trees. Unused space is managed via freelist pages, while large payloads may utilize overflow pages.

## Schema Management

The database schema is stored in a special table named `sqlite_schema` (also known as `sqlite_master`, `sqlite_temp_schema`, or `sqlite_temp_master`). This metadata is located on page 1 of the main database file and includes information about tables, indexes, triggers, and views.
