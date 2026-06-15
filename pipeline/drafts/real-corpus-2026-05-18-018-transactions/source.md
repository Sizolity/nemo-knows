---
title: Transactions
kind: source
sources:
  - raw/web/corpus-2026-05-18/018-transactions.md
confidence: medium
---

# Transactions

## What It Is
A SQLite feature that groups SQL commands into atomic units. Commands accessing the database automatically start a transaction if none is active; these auto-started transactions commit upon completion of the last statement. Transactions can also be started manually using `BEGIN`.

## Summary
The document details the syntax and behavior of SQLite transactions, covering control statements (`BEGIN`, `COMMIT`, `ROLLBACK`), transaction types (read vs. write, DEFERRED/IMMEDIATE/EXCLUSIVE), and error handling. It notes that read transactions allow multiple simultaneous connections but only one write transaction is allowed at a time.

## Key Claims
- **Syntax**: Standard commands include `BEGIN EXCLUSIVE TRANSACTION DEFERRED IMMEDIATE`, `COMMIT TRANSACTION END`, and `ROLLBACK TRANSACTION TO SAVEPOINT savepoint-name`.
- **Auto-commit**: Implicit transactions commit automatically when the last active statement finishes (cursor closes or prepared statement is reset/finalized).
- **Transaction Modes**:
  - `DEFERRED` (default): Transaction starts only upon first access.
  - `IMMEDIATE`: Starts a new write immediately; may fail with `SQLITE_BUSY`.
  - `EXCLUSIVE`: Similar to IMMEDIATE but prevents other connections from reading the database in non-WAL modes.
- **Read vs. Write**: Read transactions (started by `SELECT`) see a historic snapshot and do not see changes from other connections until the transaction ends. Write statements upgrade read transactions if possible; otherwise, they fail with `SQLITE_BUSY`.
- **Error Handling**: Errors like `SQLITE_FULL`, `SQLITE_IOERR`, `SQLITE_INTERRUPT`, or `SQLITE_NOMEM` may trigger an automatic rollback of the current statement or the entire transaction depending on context.

## Suggested Links
- https://www.sqlite.org/lang_transaction.html
