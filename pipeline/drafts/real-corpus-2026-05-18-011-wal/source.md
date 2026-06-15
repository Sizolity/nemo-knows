---
title: Write-Ahead Logging (WAL) Source Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/011-write-ahead-logging.md
confidence: medium
---

# Write-Ahead Logging (WAL) Source Summary

## What It Is
Write-Ahead Logging (WAL) is an optional journaling mode in SQLite introduced in version 3.7.0 (2010-07-21). Unlike the traditional rollback journal, which writes changes to a temporary file before applying them to the database and then deletes that file upon commit, WAL appends all changes to a separate "Write Ahead Log" file while preserving the original database content. A commit is recorded in the WAL file, allowing readers to continue accessing the original database file while writers update the log.

## Summary
WAL mode offers significant performance and concurrency benefits over the rollback journal. It allows multiple readers to access the database simultaneously without blocking writers, and writers do not block readers. Additionally, disk I/O operations tend to be more sequential, and WAL reduces the frequency of `fsync()` calls compared to rollback journals. However, WAL requires all processes to run on the same host (due to shared memory dependencies for the WAL-index), introduces new file formats requiring SQLite 3.7.0+, and involves a checkpointing mechanism that must be managed to prevent excessive WAL file growth.

## Key Claims
- **Performance:** WAL is significantly faster in most scenarios because writes are sequential and only write content once, whereas rollback journals may require writing twice.
- **Concurrency:** Readers do not block writers, and writers do not block readers. Multiple transactions can be appended to a single WAL file while readers operate on the original database.
- **Durability Trade-off:** While WAL reduces `fsync()` calls during writes, checkpointing still requires synchronization to ensure durability. If using `PRAGMA synchronous=NORMAL`, only the checkpoint issues an I/O barrier; if power is lost between a write and a checkpoint, data may be lost.
- **File Structure:** WAL mode creates two additional files: the `-wal` file (containing changes) and the `-shm` shared memory file (used by the wal-index). These persist until the last connection closes or explicitly managed.
- **Read-Only Access:** Starting in SQLite 3.22.0, read-only access to WAL-mode databases is possible if the `-shm` and `-wal` files exist and are readable, or if the database is immutable.
- **Bug History:** A data race bug known as the "WAL-reset bug" was present in versions 3.7.0 through 3.51.2 (fixed in 3.51.3). It could cause corruption if a checkpoint occurred concurrently with a transaction reset, but its occurrence rate is estimated to be extremely low (comparable to SSD malfunctions or cosmic-ray hits).

## Suggested Links
- https://www.sqlite.org/wal.html
