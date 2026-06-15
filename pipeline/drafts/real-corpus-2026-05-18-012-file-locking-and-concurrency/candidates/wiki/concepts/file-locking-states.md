---
title: File Locking States
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/012-file-locking-and-concurrency.md
confidence: medium
---

# File Locking States

In SQLite version 3.0.0 and later, the pager module manages database concurrency using a specific set of locking states to ensure ACID compliance and prevent writer starvation. These states coordinate access between threads or processes for uniform-sized blocks (pages) within the database file.

## Locking States

The system tracks five distinct locking states:

- **UNLOCKED**: The default state, indicating no locks are held on the database.
- **SHARED**: A read-only state where multiple readers are allowed simultaneously.
- **RESERVED**: A planning-to-write state held by a single process while it prepares for a write operation.
- **PENDING**: A waiting state used to prevent writer starvation; a transaction holds this lock while waiting for SHARED locks to clear so it can acquire an EXCLUSIVE lock.
- **EXCLUSIVE**: The writing state, where only one holder is allowed to modify the database.

## Transaction Lifecycle

Locks are acquired lazily during transaction execution:

1.  **Read Operations**: When a `SELECT` statement is executed in autocommit mode, the pager acquires a **SHARED** lock.
2.  **Write Planning**: Upon the first write operation within a transaction (disabling autocommit via `BEGIN TRANSACTION`), the state transitions to **RESERVED**.
3.  **Waiting for Write**: If readers are still holding SHARED locks, the transaction enters the **PENDING** state, blocking new readers but allowing existing ones to finish.
4.  **Writing**: Once all conflicting locks are cleared, the state becomes **EXCLUSIVE**, allowing the cache to be flushed or the commit to proceed.
5.  **Completion**: After `COMMIT` or `ROLLBACK`, the database returns to the **UNLOCKED** state.

## Journaling and Durability

When not in Write-Ahead Logging (WAL) mode, changes are recorded in a rollback journal before being written to the main database file. The pager coordinates access using these states to ensure durability and allow for atomic commits across multiple files, including handling "super-journals" for ATTACHed databases.

## Risks and Maintenance

Corruption risks include hardware faults, rogue processes, or operating system bugs (e.g., `fsync()` failures). Specific maintenance concerns involve deleting "hot journals" manually or filesystem corruption moving journal files to "lost+found". The pager manages these scenarios to maintain data integrity without managing B-Trees or text encodings.
