---
title: Virtual Tables Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/015-virtual-tables.md
confidence: medium
---

# Virtual Tables

## What It Is
A virtual table in SQLite is a registered object accessible via SQL statements that behaves like a standard table or view. Unlike real tables, it does not store data in the database file; instead, queries invoke callback methods defined by the virtual table implementation. They can represent in-memory structures, external disk files (e.g., CSV), or computed results.

## Summary
Virtual tables extend SQLite's capabilities by allowing custom interfaces for SQL manipulation, such as full-text search, spatial indexing, or reading host filesystems. They are created using `CREATE VIRTUAL TABLE` and managed through a specific C API (`sqlite3_module`). Implementations define methods for creation, connection, data access (reading/writing), index optimization, and lifecycle management (destruction). Security features exist to prevent misuse from triggers or views, and mechanisms allow defining hidden columns for function-like behavior.

## Key Claims
- **Functionality**: Virtual tables appear as standard tables to SQL but operate via callback methods rather than file I/O.
- **Constraints**: Triggers cannot be created on virtual tables; additional indices cannot be added separately; `ALTER TABLE ... ADD COLUMN` is not supported.
- **Creation**: Created with `CREATE VIRTUAL TABLE IF NOT EXISTS schema-name . table-name USING module-name (...)`. Temporary versions use the `temp` schema prefix.
- **Types**:
  - *Eponymous*: Exist automatically when their module is registered (if `xCreate` equals `xConnect`).
  - *Eponymous-only*: Cannot be created with `CREATE VIRTUAL TABLE`; exist only via module name (useful for table-valued functions).
- **Implementation**: Defined by the `sqlite3_module` structure containing method pointers like `xCreate`, `xBestIndex`, `xFilter`, etc.
- **Security**: Implementations can restrict usage to direct SQL (`SQLITE_VTAB_DIRECTONLY`) or mark as innocuous (`SQLITE_VTAB_INNOCUOUS`).
- **Shadow Tables**: Some implementations (e.g., FTS3, RTREE) use auxiliary real tables called "shadow tables" to store content; these are protected from direct SQL modification.

## Suggested Links
- https://www.sqlite.org/vtab.html
