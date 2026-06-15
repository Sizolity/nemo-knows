---
kind: source
sources:
  - web:https://www.sqlite.org/wal.html
confidence: medium
---

# SQLite Write-Ahead Logging

## What It Is

SQLite write-ahead logging is an alternative transaction journaling mode in
which changes are appended to a separate WAL file before being checkpointed back
into the main database file.

## Summary

The WAL mode changes SQLite's write path from rollback-journal updates to an
append-first design. Readers can continue reading from the original database
while writers append new changes to the WAL. A checkpoint operation later moves
committed changes back into the main database. The documentation emphasizes
concurrency benefits, checkpoint behavior, and operational trade-offs such as
extra files and local-filesystem assumptions.

## Key Claims

- WAL allows readers and writers to proceed concurrently in more cases than the
  rollback journal mode.
- Committed changes are appended to the WAL and later transferred into the main
  database by checkpointing.
- WAL requires all processes using the database to be on the same host because
  shared memory coordination is involved.
- WAL can improve performance but introduces checkpoint management and extra
  files.

## Suggested Links

- SQLite
- write-ahead logging
- checkpointing
- transaction journaling
