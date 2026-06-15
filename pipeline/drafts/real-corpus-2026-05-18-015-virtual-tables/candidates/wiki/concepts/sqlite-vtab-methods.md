---
title: Sqlite Vtab Methods
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/015-virtual-tables.md
confidence: medium
---

# Sqlite Vtab Methods

In SQLite, a virtual table is a registered object accessible via SQL statements that behaves like a standard table or view. Unlike real tables, it does not store data in the database file; instead, queries invoke callback methods defined by the virtual table implementation. They can represent in-memory structures, external disk files (e.g., CSV), or computed results.

## Implementation Structure

Virtual tables are managed through a specific C API (`sqlite3_module`). Implementations define methods for creation, connection, data access (reading/writing), index optimization, and lifecycle management (destruction). The structure contains method pointers such as `xCreate`, `xBestIndex`, `xFilter`, and others.

## Usage Constraints

- Triggers cannot be created on virtual tables.
- Additional indices cannot be added separately.
- `ALTER TABLE ... ADD COLUMN` is not supported.

## Creation Syntax

Virtual tables are created using `CREATE VIRTUAL TABLE`. Temporary versions use the `temp` schema prefix. The syntax follows:

```sql
CREATE VIRTUAL TABLE IF NOT EXISTS schema-name . table-name USING module-name (...)
```

## Types

- **Eponymous**: Exist automatically when their module is registered (if `xCreate` equals `xConnect`).
- **Eponymous-only**: Cannot be created with `CREATE VIRTUAL TABLE`; exist only via module name. These are useful for table-valued functions.

## Security and Shadow Tables

Implementations can restrict usage to direct SQL (`SQLITE_VTAB_DIRECTONLY`) or mark as innocuous (`SQLITE_VTAB_INNOCUOUS`). Some implementations (e.g., FTS3, RTREE) use auxiliary real tables called "shadow tables" to store content; these are protected from direct SQL modification.

[[sqlite-vtab-methods]]
