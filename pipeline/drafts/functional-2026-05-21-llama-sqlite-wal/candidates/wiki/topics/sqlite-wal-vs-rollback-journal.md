---
title: Sqlite Wal Vs Rollback Journal
kind: topic
sources:
  - source.md
  - raw/web/sqlite-wal.md
confidence: medium
---

# Sqlite Wal Vs Rollback Journal

SQLite supports multiple journaling strategies to manage transaction safety, with the primary distinction lying between the default rollback journal and the write-ahead logging (WAL) mode. While both mechanisms ensure data integrity, they employ different file structures and concurrency models to handle database access patterns.

## Concurrency Models

The fundamental advantage of WAL is its ability to facilitate overlapping read and write operations. In this configuration, readers can access the original database file without being blocked by writers, as new changes are appended to a separate log file instead of modifying the main data file directly. This separation ensures that active commits do not disrupt existing snapshots held by readers, provided they utilize the database file combined with the WAL up to the recorded end mark. Conversely, the default rollback journal mode typically restricts concurrent access more strictly, as writers must acquire locks that prevent simultaneous reads from accessing uncommitted changes.

## File Structure and Persistence

When operating in WAL mode, the database environment requires three specific files: the main database file, a write-ahead log file (named with the `-wal` suffix), and a shared memory file (named with the `-shm` suffix). These additional components allow for efficient coordination between processes. The WAL itself remains persistent at the database-file level until the journal mode is explicitly changed again or a checkpoint occurs. In contrast, the rollback journal usually exists only temporarily during a transaction to record changes before being either merged into the main file or rolled back upon failure.

## Performance and Host Constraints

Modern implementations of SQLite often find that very large transactions are no longer a significant disadvantage in WAL mode, though checkpoint behavior can still influence latency. The performance benefits generally stem from reduced locking contention, allowing for faster overall throughput compared to the rollback journal approach. However, this setup introduces host-level constraints; because WAL relies on shared memory files for coordination, it typically requires all processes accessing the database to reside on the same host machine.
