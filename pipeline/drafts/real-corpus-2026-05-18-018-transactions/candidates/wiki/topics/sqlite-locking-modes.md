---
title: Sqlite Locking Modes
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/018-transactions.md
confidence: medium
---

# Sqlite Locking Modes

SQLite supports three distinct transaction modes that control how and when a transaction begins. These modes interact closely with the database's locking mechanism, particularly regarding write operations and concurrent access. The available modes are `DEFERRED`, `IMMEDIATE`, and `EXCLUSIVE`.

## Transaction Mode Descriptions

### DEFERRED
This is the default mode for SQLite transactions. In this configuration, a transaction does not begin until the first data access occurs within the transaction scope. While `SELECT` statements can execute without holding locks on unmodified pages, any write operation (such as an `INSERT`, `UPDATE`, or `DELETE`) triggers the acquisition of a write lock. If the database is in journal mode, it requires exclusive access to the database file at the moment of the write.

### IMMEDIATE
In this mode, a transaction starts immediately upon invocation, even before any SQL statement executes. This means that if the connection attempts to perform a write operation while another connection holds a lock on the database (in non-WAL modes), the current transaction may fail with an `SQLITE_BUSY` error rather than waiting or proceeding with a read-only lock.

### EXCLUSIVE
The `EXCLUSIVE` mode behaves similarly to `IMMEDIATE`, starting the transaction immediately upon invocation. However, it imposes stricter constraints on concurrent access. Specifically, if the database is not using Write-Ahead Logging (WAL) mode, an exclusive transaction prevents other connections from reading the database until the current transaction commits or rolls back.

## Behavior and Interactions

The choice of locking mode affects concurrency and error handling:

*   **Read vs. Write**: Read transactions (typically initiated by `SELECT`) allow multiple simultaneous connections but do not see changes made by other connections until the transaction ends. Write statements attempt to upgrade read transactions if possible; otherwise, they fail with an `SQLITE_BUSY` error.
*   **Auto-commit**: If no explicit transaction is active, SQLite automatically starts one upon the first data access. These implicit transactions commit automatically when the last active statement finishes (e.g., cursor closes or prepared statement is finalized).
*   **Error Handling**: Certain errors such as `SQLITE_FULL`, `SQLITE_IOERR`, `SQLITE_INTERRUPT`, or `SQLITE_NOMEM` may trigger an automatic rollback of the current statement or the entire transaction depending on the context.
