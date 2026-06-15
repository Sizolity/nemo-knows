---
title: Sqlite Isolation Modes
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/019-isolation-in-sqlite.md
confidence: medium
---

# Sqlite Isolation Modes

SQLite isolation describes what one database connection can observe while another connection is changing the database. The default behavior keeps uncommitted work private to the writer, with special cases only when applications deliberately combine shared cache mode with `PRAGMA read_uncommitted`.

## Summary

For separate connections, SQLite presents serializable behavior by ensuring only one writer proceeds at a time. Readers do not see another connection's partial transaction. WAL mode changes concurrency by allowing readers and a writer to overlap, but each reader keeps a stable view from the start of its read transaction.

Visibility inside one connection is different from visibility across connections. Statements on the same connection can observe changes already made by earlier statements on that connection, even if the transaction has not committed. The source cautions that changing a table while a query on that same table is still running leads to undefined visibility.

## Isolation Characteristics

### Serializable Isolation
SQLite's locking rules serialize writes. That prevents concurrent writers from exposing interleaved partial updates as independent transactions.

### Connection Isolation
Different connections normally do not see each other's uncommitted changes. The documented exception requires both shared cache mode and the `read_uncommitted` pragma.

### WAL Mode Behavior
In WAL mode, readers and a writer can operate at the same time. A reader continues using the database snapshot that was current when its read transaction began.

### Single Connection Visibility
Statements on one connection see prior changes from that same connection. They should not rely on predictable results if rows are changed while an earlier scan is still in progress.

## Unsafe Concurrent Reads/Writes on Same Connection

The practical rule is to avoid mixing long-running scans with writes to the same table on the same connection unless the application does not depend on whether those new changes are visible to the scan.
