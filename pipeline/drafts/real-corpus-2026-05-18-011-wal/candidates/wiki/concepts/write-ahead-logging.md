---
title: Write Ahead Logging
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/011-write-ahead-logging.md
confidence: medium
---

# Write Ahead Logging

Write-Ahead Logging (WAL) is an optional journaling mode in SQLite introduced in version 3.7.0. Unlike the traditional rollback journal, which writes changes to a temporary file before applying them and then deletes that file upon commit, WAL appends all changes to a separate "Write Ahead Log" file while preserving the original database content. A commit is recorded in the WAL file, allowing readers to continue accessing the original database file while writers update the log.

## Mechanism and Files

In WAL mode, the database structure consists of:
- The main database file containing the committed data.
- The `-wal` file, which contains uncommitted or recently committed changes appended sequentially.
- The `-shm` shared memory file, used by the wal-index to coordinate access between readers and writers.

These additional files persist until the last connection closes or are explicitly managed. Starting in SQLite 3.22.0, read-only access to WAL-mode databases is possible if the `-shm` and `-wal` files exist and are readable, or if the database is immutable.

## Performance and Concurrency Benefits

WAL mode offers significant performance and concurrency benefits over the rollback journal:
- **Concurrency:** Readers do not block writers, and writers do not block readers. Multiple transactions can be appended to a single WAL file while readers operate on the original database.
- **Performance:** WAL is significantly faster in most scenarios because writes are sequential and only write content once, whereas rollback journals may require writing twice.
- **Disk I/O:** WAL reduces the frequency of `fsync()` calls compared to rollback journals, as disk I/O operations tend to be more sequential.

## Durability and Checkpointing

While WAL reduces `fsync()` calls during writes, checkpointing still requires synchronization to ensure durability. If using `PRAGMA synchronous=NORMAL`, only the checkpoint issues an I/O barrier; if power is lost between a write and a checkpoint, data may be lost. The checkpointing mechanism must be managed to prevent excessive WAL file growth.

## History and Bugs

A data race bug known as the "WAL-reset bug" was present in SQLite versions 3.7.0 through 3.51.2 and was fixed in version 3.51.3. This bug could cause corruption if a checkpoint occurred concurrently with a transaction reset, but its occurrence rate is estimated to be extremely low (comparable to SSD malfunctions or cosmic-ray hits).

## Limitations

- WAL requires all processes to run on the same host due to shared memory dependencies for the wal-index.
- WAL introduces new file formats requiring SQLite 3.7.0+.
