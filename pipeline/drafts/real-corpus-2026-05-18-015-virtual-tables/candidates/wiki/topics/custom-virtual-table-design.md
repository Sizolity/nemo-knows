---
title: Custom Virtual Table Design
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/015-virtual-tables.md
confidence: medium
---

# Custom Virtual Table Design

A custom virtual table lets SQLite expose non-standard data sources through the table interface. SQL statements can query or modify the object as though it were a table, while the actual behavior is supplied by extension callbacks rather than by SQLite's ordinary table storage.

## Implementation Structure

Virtual table modules are registered with SQLite and then instantiated with `CREATE VIRTUAL TABLE`. The module can represent data held in memory, external files, computed views, search indexes, or other custom stores. This is the mechanism behind features such as full-text search and spatial indexing.

The C API centers on `sqlite3_module`, a table of callbacks. Methods such as `xCreate`, `xConnect`, `xBestIndex`, and `xFilter` let the module participate in table creation, planning, scanning, updates, and cleanup. Query planning is especially important because `xBestIndex` tells SQLite which constraints and orderings the module can use efficiently.

## Creation Syntax

The SQL creation form names a module and passes module-specific arguments:

```sql
CREATE VIRTUAL TABLE IF NOT EXISTS schema-name . table-name USING module-name (...)
```

Temporary virtual tables use the `temp` schema. The source also identifies restrictions that matter for design: virtual tables do not accept separate indexes, cannot have triggers attached to them, and do not support adding columns with `ALTER TABLE`.

## Types of Virtual Tables

- **Eponymous modules** can be used by name without an explicit `CREATE VIRTUAL TABLE` instance when their implementation supports that pattern.
- **Eponymous-only modules** are available only through direct module invocation, which makes them useful for table-valued function behavior.

## Security and Shadow Tables

Virtual table implementations can expose hidden columns to support function-like arguments. They can also opt into safety markers that constrain where the table may be used or indicate that a module is safe in less-trusted SQL contexts.

Some modules maintain auxiliary storage in shadow tables. Because direct writes to those tables can corrupt module invariants, SQLite includes protections around how shadow tables are accessed.
