---
title: Write-Ahead Logging
kind: source
sources:
  - raw/web/corpus-2026-05-18/011-write-ahead-logging.md
confidence: medium
---

# Write-Ahead Logging

## What It Is
Write-Ahead Logging (WAL) is a journaling method available in SQLite since version 3.7.0 that differs from the traditional rollback journal. Instead of writing changes to a separate rollback journal and then applying them to the main database file, WAL appends changes to a separate `-wal` file while preserving the original database content. A commit is recorded in the WAL file, allowing readers to continue accessing the unaltered database file while writers commit changes.

## Summary
WAL mode offers significant performance and concurrency benefits, particularly for write-heavy workloads. It allows readers and writers to operate concurrently without blocking each other, as readers access the original database file while writers append to the WAL file. However, WAL mode requires all processes to reside on the same host computer due to its reliance on shared memory for the WAL index. It introduces a separate WAL file and a shared memory file, which must be managed via checkpointing to prevent excessive file growth.

## Key Claims
- **Performance:** WAL is significantly faster in most scenarios because write transactions only involve writing content once and are sequential. It reduces the number of `fsync()` operations.
- **Concurrency:** Readers do not block writers, and writers do not block readers, enabling concurrent read and write operations.
- **Durability Trade-off:** While writes are fast, syncing to disk is not required for every write unless `PRAGMA synchronous` is set to `FULL`. Checkpointing requires sync operations to ensure durability.
- **Limitations:** WAL does not work over network filesystems. It is not possible to change the page size after entering WAL mode. Older versions of SQLite (prior to 3.7.0) cannot recover databases in WAL mode.
- **Bug History:** A "WAL-reset bug" existed in versions 3.7.0 through 3.51.2, which could cause corruption in rare cases involving simultaneous checkpoints and commits. This was fixed in version 3.51.3.
- **Checkpointing:** By default, SQLite automatically checkpoints when the WAL file reaches 1000 pages. Checkpoints can be manual or automatic, and they transfer content from the WAL back into the database file.

## Suggested Links
- https://www.sqlite.org/wal.html
