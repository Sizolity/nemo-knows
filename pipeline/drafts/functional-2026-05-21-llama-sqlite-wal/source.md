---
title: SQLite Write-Ahead Logging (WAL)
kind: source
sources:
  - raw/web/sqlite-wal.md
confidence: medium
---

# SQLite Write-Ahead Logging (WAL)

## What It Is
SQLite Write-Ahead Logging (WAL) is an alternative journaling mode to the default rollback journal. In this mode, changes are appended to a separate write-ahead log file rather than being written directly into the main database file.

## Summary
The documentation outlines how WAL improves concurrency by allowing readers to access the original database file while writers append changes to the log. Readers utilize the database file combined with the WAL up to a recorded end mark, ensuring that newer commits do not disrupt their snapshot. Periodically, a checkpoint operation copies committed frames from the WAL back into the main database file.

## Key Claims
- **Concurrency**: WAL allows readers and writers to overlap operations, often resulting in faster performance compared to rollback-journal mode.
- **File Structure**: WAL mode requires the presence of additional `-wal` and `-shm` files alongside the main database file.
- **Host Constraints**: It typically requires all processes accessing the database to reside on the same host due to the use of shared memory for coordination.
- **Transaction Handling**: While very large transactions are no longer a specific disadvantage in modern SQLite, checkpoint behavior can still impact latency.
- **Persistence**: The WAL is persistent at the database-file level until the journal mode is explicitly changed again.

## Suggested Links
- https://www.sqlite.org/wal.html
