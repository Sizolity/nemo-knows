---
title: Checkpointing
kind: concept
sources:
  - source.md
  - raw/web/sqlite-wal.md
confidence: medium
---

# Checkpointing

In SQLite's Write-Ahead Logging (WAL) mode, changes are appended to a separate log file rather than being written directly into the main database file. While readers access the original database file alongside the WAL, committed changes remain in the log until they are explicitly moved back. A **[[checkpointing]]** operation is required to copy these committed frames from the WAL back into the main database file, ensuring data durability and persistence at the file level.

This process allows for a distinct operational profile where writers append to the log while readers utilize the snapshot up to a recorded end mark. However, managing this separation introduces trade-offs, such as the requirement for additional `-wal` and `-shm` files and constraints on shared memory usage that typically necessitate processes running on the same host.
