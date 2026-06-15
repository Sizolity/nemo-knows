---
title: Query Planner Cost Analysis
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/014-query-optimizer-overview.md
confidence: medium
---

# Query Planner Cost Analysis

The SQLite query planner and optimizer select algorithms that minimize disk I/O and CPU overhead for a given SQL statement. Its primary task is to analyze the database schema and statement complexity to choose the fastest plan from competing options, such as deciding between full table scans and index usage.

## Cost Estimation
SQLite uses a cost-based approach to estimate CPU and disk I/O costs. The optimizer evaluates multiple execution plans and selects the one with the lowest estimated cost.

## Optimization Techniques

### Index Usage Rules
An index is considered usable if:
- The left-most columns appear in the `WHERE` clause with equality operators (`=`, `IN`, `IS`).
- Inequalities are allowed only on the right-most column of an index prefix.
- There are no gaps in constrained columns.

### OR Optimizations
OR-connected constraints can be converted to `IN` operators or evaluated separately using a UNION-like mechanism. This allows potentially utilizing different indexes for each subterm.

### Skip-Scan Optimization
SQLite may use an index even if the left-most column is unconstrained but later columns are constrained. This requires the left-most column to have many duplicate values and typically necessitates running `ANALYZE`.

### Join Order
- Inner joins are reordered automatically to optimize performance.
- Outer joins maintain their original order.
- Join order can be manually forced using `CROSS JOIN` or by manipulating `sqlite_stat` tables.

### Subquery Flattening
Subqueries in the `FROM` clause are flattened into the outer query if possible to avoid creating transient tables. This is subject to constraints such as the absence of `LIMIT` or aggregates in subqueries.

### Automatic Indexes
If no persistent indexes exist and a query requires multiple lookups, SQLite creates transient "automatic indexes" for the duration of the statement.

### Covering Indexes
If an index contains all columns needed for a query, SQLite avoids looking up the original table row entirely, improving performance by reducing I/O.
