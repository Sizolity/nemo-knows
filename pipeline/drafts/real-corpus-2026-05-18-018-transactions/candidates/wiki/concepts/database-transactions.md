---
title: Database Transactions
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/018-transactions.md
confidence: medium
---

# Database Transactions

A SQLite feature that groups SQL commands into atomic units. Commands accessing the database automatically start a transaction if none is active; these auto-started transactions commit upon completion of the last statement. Transactions can also be started manually using `BEGIN`.

## Control Statements

Standard control statements include:
- `BEGIN`: Starts a new transaction.
- `COMMIT`: Finalizes and saves changes.
- `ROLLBACK`: Undoes changes since the last commit or savepoint.

Savepoints are managed using `ROLLBACK TRANSACTION TO SAVEPOINT savepoint-name`.

## Transaction Modes

Transaction modes define isolation levels and concurrency behavior:
- **DEFERRED** (default): The transaction starts only upon the first access to the database.
- **IMMEDIATE**: Starts a new write transaction immediately; it may fail with `SQLITE_BUSY` if another connection holds a write lock.
- **EXCLUSIVE**: Similar to IMMEDIATE but prevents other connections from reading the database in non-WAL modes.

## Read vs. Write Behavior

- **Read transactions** (started by `SELECT`) see a historic snapshot and do not see changes from other connections until the transaction ends.
- **Write statements** upgrade read transactions if possible; otherwise, they fail with `SQLITE_BUSY`.

Only one write transaction is allowed at a time. Read transactions allow multiple simultaneous connections.

## Error Handling

Errors such as `SQLITE_FULL`, `SQLITE_IOERR`, `SQLITE_INTERRUPT`, or `SQLITE_NOMEM` may trigger an automatic rollback of the current statement or the entire transaction depending on context.
