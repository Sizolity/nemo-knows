---
title: File Locking and Concurrency (SQLite)
kind: source
sources:
  - raw/web/corpus-2026-05-18/012-file-locking-and-concurrency.md
confidence: medium
---

# File Locking and Concurrency in SQLite

## What It Is
A technical document describing the locking and journaling mechanism introduced in SQLite version 3.0.0 to improve concurrency over version 2 and reduce writer starvation. The text focuses on the **rollback-mode** transaction mechanism (distinct from Write-Ahead Logging). It details how the pager module handles ACID compliance, manages database files as uniform-sized blocks (pages), and coordinates access between threads or processes using specific locking states.

## Summary
SQLite version 3 uses a new locking mechanism managed by the pager module to ensure Atomic, Consistent, Isolated, and Durable (ACID) properties. The system tracks five locking states: **UNLOCKED**, **SHARED**, **RESERVED**, **PENDING**, and **EXCLUSIVE**. Changes are recorded in a rollback journal before being written to the main database file to ensure durability and allow for atomic commits across multiple files. The document outlines algorithms for handling "hot journals" (created during crashes), deleting stale super-journals, and managing writer starvation via PENDING locks. It also covers SQL-level transaction control (`BEGIN TRANSACTION`, `COMMIT`, `ROLLBACK`) and potential risks of corruption due to hardware faults or operating system bugs.

## Key Claims
- **Locking States**: There are five states: UNLOCKED (default), SHARED (read-only, multiple readers allowed), RESERVED (planning to write, single holder), PENDING (waiting for SHARED locks to clear to get EXCLUSIVE), and EXCLUSIVE (writing, only one allowed).
- **Pager Role**: The pager module controls access for threads/processes and is responsible for ACID compliance without managing B-Trees or text encodings.
- **Journaling**: When not in WAL mode, changes are recorded in a rollback journal (`-journal` suffix) before the main file is altered. A "super-journal" coordinates multiple ATTACHed databases.
- **Hot Journals**: A journal is "hot" if it needs rollback due to a crash or power failure. It is identified by existence, non-zero header size (>512 bytes), and lack of a RESERVED lock on the database.
- **Writer Starvation Prevention**: The PENDING lock allows existing readers to finish but blocks new readers, ensuring writers eventually get an EXCLUSIVE lock.
- **Corruption Risks**: Corruption can occur via hardware faults, rogue processes, or OS bugs (e.g., `fsync()` failures, NFS locking issues). Specific risks include deleting hot journals manually or filesystem corruption moving journal files to "lost+found".
- **Transaction Control**: By default, SQLite is in autocommit mode. `BEGIN TRANSACTION` disables autocommit; locks are acquired lazily (SHARED on first SELECT, RESERVED on first write, EXCLUSIVE only when flushing cache or committing).

## Suggested Links
*   https://www.sqlite.org/lockingv3.html
