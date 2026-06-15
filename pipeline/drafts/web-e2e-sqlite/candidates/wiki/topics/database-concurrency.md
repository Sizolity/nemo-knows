---
title: Database Concurrency
kind: topic
sources:
  - source.md
  - raw/web/sqlite-wal.md
confidence: medium
---

# Database Concurrency

Database concurrency refers to the ability of a database system to allow multiple readers and writers to access data simultaneously without compromising data integrity or performance. One prominent implementation strategy is Write-Ahead Logging (WAL), often used in conjunction with standard rollback journaling modes.

## Write-Ahead Logging (WAL)

In WAL mode, changes are appended to a separate write-ahead log file rather than being written directly into the main database file or preserved in a rollback journal. This architectural shift enables readers to continue accessing the original database file while writers append changes to the WAL. By utilizing the database file plus the WAL up to a recorded end mark, readers ensure that newer commits do not disturb their snapshot.

### Trade-offs and Constraints

While WAL improves concurrency by allowing readers and writers to overlap, it introduces specific operational requirements:
- **File Management**: Requires extra `-wal` and `-shm` files.
- **Shared Memory**: Normally requires all processes using the database to be on the same host due to shared memory usage for coordination.
- **Checkpointing**: A checkpoint operation is required to copy committed frames from the WAL back into the main database file. While very large transactions are no longer a special disadvantage in modern implementations, checkpoint behavior can still affect latency.

## Related Concepts

- [[checkpointing]]
- [[log]]
