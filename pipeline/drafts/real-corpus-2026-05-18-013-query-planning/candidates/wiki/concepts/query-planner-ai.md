---
title: Query Planner Ai
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/013-query-planning.md
confidence: medium
---

# Query Planner Ai

The **Query Planner AI** acts as an intelligent subsystem within the SQLite database engine. Its primary function is to determine the most efficient algorithm to execute a given SQL statement. Because SQL is a declarative language, it specifies *what* data needs to be computed, leaving the system to figure out *how* to compute it.

The planner evaluates hundreds or thousands of possible execution options for each specific statement and selects the fastest path. It relies heavily on available **indices** to function effectively, avoiding full table scans for large tables whenever possible.

## Core Mechanisms

### Search Optimization
To avoid scanning entire tables, the planner utilizes rowids or indices (single-column and multi-column). If an index exists on a search column, lookups are significantly faster than scanning content without one.

### Sorting Strategies
Indices allow sorting operations to occur without separate steps. The planner employs **covering indexes** that include all output columns, enabling SQLite to avoid referencing the original table entirely. This strategy can roughly double speed compared to non-covering indices.

When an index satisfies only part of an `ORDER BY` clause, the system performs multiple small sorts (partial sorting or block sorting) rather than one large sort to save CPU cycles and temporary storage.

### Handling Complex Logic
For queries involving `OR` clauses in the `WHERE` condition, the planner examines each term separately using available indices and combines results via a union operation. If no index exists for a specific term, a full table scan is required for that portion of the query.

## Index Management Guidelines

Programmers must normally add indices to support the planner's work. Effective management includes:
- Preferring multi-column indices over single-column prefixes.
- Avoiding the creation of an index that is a prefix of another existing index.
- Utilizing covering indexes to minimize table access.

While a single multi-column index can often replace multiple single-column indices, this optimization depends on the specific query structure and available data.
