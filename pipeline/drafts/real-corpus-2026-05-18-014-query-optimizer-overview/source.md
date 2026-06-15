---
title: Query Optimizer Overview Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/014-query-optimizer-overview.md
confidence: medium
---

# Query Optimizer Overview Summary

## What It Is
This document provides an overview of the SQLite query planner and optimizer. Its primary task is to select algorithms that minimize disk I/O and CPU overhead for a given SQL statement by analyzing the database schema and statement complexity.

## Summary
SQLite uses a cost-based query planner to estimate CPU and disk I/O costs, choosing the fastest plan from competing options (e.g., full table scans vs. index usage). The optimizer processes `WHERE` clauses by converting join constraints into conjuncts, analyzing terms for index usage, and applying specific optimizations like skip-scan, OR-clause handling, and subquery flattening. It supports manual control of join order via `CROSS JOIN` or statistical manipulation of `sqlite_stat` tables.

## Key Claims
- **Index Usage Rules**: An index is usable if the left-most columns appear in the `WHERE` clause with equality (`=`, `IN`, `IS`) operators. Inequalities are allowed only on the right-most column of an index prefix. Gaps in constrained columns disqualify the index.
- **OR Optimizations**: OR-connected constraints can be converted to `IN` operators or evaluated separately using a UNION-like mechanism, potentially utilizing different indexes for each subterm.
- **Skip-Scan Optimization**: SQLite uses an index even if the left-most column is unconstrained but later columns are constrained, provided the left-most column has many duplicate values (requires `ANALYZE`).
- **Join Order**: Inner joins are reordered automatically; outer joins maintain order. Join order can be manually forced using `CROSS JOIN`.
- **Automatic Indexes**: If no persistent indexes exist and a query requires multiple lookups, SQLite creates transient "automatic indexes" for the duration of the statement.
- **Subquery Flattening**: Subqueries in the `FROM` clause are flattened into the outer query if possible to avoid creating transient tables, subject to specific constraints (e.g., no `LIMIT` or aggregates in subqueries).
- **Covering Indexes**: If an index contains all columns needed for a query, SQLite avoids looking up the original table row entirely.

## Suggested Links
*   https://www.sqlite.org/optoverview.html
