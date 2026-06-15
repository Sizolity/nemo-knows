---
title: Isolation in SQLite
kind: source
sources:
  - raw/web/corpus-2026-05-18/019-isolation-in-sqlite.md
confidence: medium
---

# Isolation in SQLite

## What It Is
Isolation in SQLite determines when changes made by one operation become visible to other concurrent operations. The behavior depends heavily on whether database connections share a cache and whether the `read_uncommitted` pragma is enabled. By default, separate connections are isolated from each other unless specific configuration overrides this.

## Summary
SQLite transactions are serializable. Changes made in one database connection are invisible to all other database connections until the writer commits. Within a single database connection, a query sees all changes completed prior to its start, regardless of commit status; however, seeing changes occurring while the query is running is undefined behavior. Shared cache mode combined with `PRAGMA read_uncommitted` effectively treats multiple connections as a single one regarding visibility.

## Key Claims
- **Serializable Isolation:** All transactions in SQLite show "serializable" isolation by serializing writes to ensure only a single writer exists at a time via automatic locking.
- **Connection Isolation:** Partial changes by a writer are invisible to readers on different connections unless `read_uncommitted` is turned on and shared cache mode is used.
- **WAL Mode Behavior:** Write-Ahead Log (WAL) mode permits simultaneous readers and writers, exhibiting "snapshot isolation" where a read transaction sees an unchanging snapshot of the database at the start time.
- **Single Connection Visibility:** Within the same connection, subsequent SELECT statements see changes made by UPDATE/INSERT/DELETE prior to their execution, but behavior regarding concurrent modifications during query execution is undefined.
- **Unsafe Concurrent Reads/Writes on Same Connection:** Developers should avoid assuming a SELECT statement will or won't see changes made to the same table after the query starts but before it completes.

## Suggested Links
- https://www.sqlite.org/isolation.html
