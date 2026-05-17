---
title: SQLite Write-Ahead Logging (WAL) Summary
kind: source
sources:
  - raw/web/sqlite-wal.md
confidence: medium
---

# SQLite Write-Ahead Logging (WAL) Summary

## What It Is
SQLite's write-ahead logging (WAL) is an alternative to the default rollback journal mode. In WAL mode, changes are appended to a separate write-ahead log file rather than being written directly into the database file or preserved in a rollback journal.

## Summary
The documentation highlights that WAL improves concurrency by allowing readers to continue accessing the original database file while writers append changes to the WAL. Readers utilize the database file plus the WAL up to a recorded end mark, ensuring newer commits do not disturb their snapshot. A checkpoint operation is required to copy committed frames from the WAL back into the main database file.

## Key Claims
- **Concurrency**: Readers can read the original database while writers append changes to the WAL, allowing readers and writers to overlap.
- **Durability and Persistence**: WAL is persistent at the database-file level until the journal mode is changed again.
- **Trade-offs**:
  - Requires extra `-wal` and `-shm` files.
  - Normally requires all processes using the database to be on the same host due to shared memory usage for coordination.
  - Very large transactions are no longer a special disadvantage in modern SQLite, though checkpoint behavior can still affect latency.
- **Operational Profile**: WAL offers a different concurrency and durability profile that is often faster but requires understanding of checkpointing, shared-memory constraints, and additional file management.

## Suggested Links
- https://www.sqlite.org/wal.html
