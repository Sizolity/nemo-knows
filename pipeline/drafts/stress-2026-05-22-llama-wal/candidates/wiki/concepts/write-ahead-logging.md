---
title: Write Ahead Logging
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/011-write-ahead-logging.md
confidence: medium
---

# Write Ahead Logging

Write-Ahead Logging (WAL) is a journaling mechanism introduced in SQLite version 3.7.0 that differs from the traditional rollback journal approach. Instead of modifying the main database file directly and then applying changes, WAL appends transaction data to a separate `-wal` file while keeping the original database file intact. A commit is finalized by appending a record to this log file, allowing readers to continue accessing the unmodified database file even while writers are committing changes.

## Mechanism and Operation

In WAL mode, the system utilizes three primitive operations: reading, writing, and checkpointing. When a transaction commits, a special record is appended to the WAL file without immediately writing to the main database. This separation enables concurrent access where readers operate on the original file and writers append to the log. Eventually, the contents of the WAL file must be transferred back into the main database file through a process called checkpointing. By default, SQLite performs an automatic checkpoint when the WAL file reaches a threshold of 1000 pages, though this can be adjusted or triggered manually.

## Performance and Concurrency

WAL mode provides significant performance advantages, particularly for write-heavy workloads. Because write transactions involve sequential writes to the log file, disk I/O operations are more efficient. Additionally, WAL reduces the frequency of `fsync()` system calls, which are often a bottleneck for durability. The architecture allows readers and writers to operate concurrently without blocking each other, improving overall throughput compared to the rollback journal method.

## Limitations and Constraints

Despite its benefits, WAL mode has specific limitations. It requires all processes to reside on the same host computer because it relies on shared memory for the WAL index, which cannot be shared across network filesystems. Furthermore, the page size cannot be changed once WAL mode is enabled, whether on an empty database or an existing one. Older versions of SQLite prior to 3.7.0 cannot recover databases in WAL mode. There is also a historical "WAL-reset bug" present in versions 3.7.0 through 3.51.2 that could cause corruption in rare scenarios involving simultaneous checkpoints and commits; this issue was resolved in version 3.51.3.

## Configuration and Management

Managing WAL involves handling the separate WAL file and a shared memory file. To prevent excessive file growth, checkpointing is essential. While writes are fast, ensuring durability requires syncing to disk, which is not performed for every write unless the `PRAGMA synchronous` setting is explicitly configured to `FULL`. Applications can adjust the automatic checkpoint threshold or initiate checkpoints manually. Note that while most read-only databases can be opened in WAL mode, specific conditions regarding the existence of the `-shm` and `-wal` files may apply depending on the SQLite version.
