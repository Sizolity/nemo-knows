---
title: Sqlite Optimizations Indexing
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/014-query-optimizer-overview.md
confidence: medium
---

# Sqlite Optimizations Indexing

SQLite employs a cost-based query planner to minimize disk I/O and CPU overhead. The optimizer analyzes the database schema and statement complexity to select algorithms, such as choosing between full table scans and index usage. This process involves estimating costs for CPU and disk I/O operations to determine the fastest execution plan.

## Index Usage Rules

An index is considered usable if specific conditions regarding column order and operators are met:
- The left-most columns of an index must appear in the `WHERE` clause with equality operators (`=`, `IN`, `IS`).
- Inequalities are permitted only on the right-most column of an index prefix.
- Gaps in constrained columns disqualify the index from use.

## Advanced Optimizations

### Skip-Scan Optimization
SQLite utilizes skip-scan optimization when the left-most column of an index is unconstrained but later columns are constrained. This is effective provided the left-most column has many duplicate values, a condition that typically requires running `ANALYZE`.

### OR Clause Handling
Constraints connected by `OR` can be optimized by converting them into `IN` operators or evaluating them separately using a UNION-like mechanism. This approach allows the optimizer to potentially utilize different indexes for each subterm of the expression.

### Covering Indexes
If an index contains all columns required for a specific query, SQLite avoids looking up the original table row entirely, improving performance by reading directly from the index structure.

### Automatic Indexes
When no persistent indexes exist and a query requires multiple lookups, SQLite creates transient "automatic indexes" for the duration of the statement to facilitate efficient execution.

## Subquery Flattening

Subqueries located in the `FROM` clause may be flattened into the outer query if possible. This optimization avoids creating transient tables. However, this flattening is subject to specific constraints, such as restrictions on using `LIMIT` or aggregates within subqueries.
