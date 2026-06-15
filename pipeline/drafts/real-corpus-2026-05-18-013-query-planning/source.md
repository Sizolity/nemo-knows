---
title: Query Planning (Source Summary)
kind: source
sources:
  - raw/web/corpus-2026-05-18/013-query-planning.md
confidence: medium
---

# Query Planning

## What It Is
Query planning in SQLite is the process by which the database engine determines the most efficient algorithm to execute a given SQL statement. As a declarative language, SQL tells the system *what* to compute, while the query planner subsystem figures out *how* to compute it. The planner acts as an AI that attempts to select the fastest and most efficient algorithm among hundreds or thousands of possible options for each specific statement.

## Summary
The document provides a technical explanation of how SQLite's query planner and engine work, focusing on the use of indices to optimize performance. Key concepts include:
- **Searching**: Techniques to avoid full table scans by using rowids or indices (single-column, multi-column).
- **Sorting**: How indices allow sorting without separate steps, including covering indexes and block sorting.
- **Combined Operations**: Strategies for searching and sorting simultaneously using multi-column or covering indexes.
- **Index Management**: Guidelines on creating effective indices, such as preferring multi-column indices over single-column prefixes and utilizing "covering indexes" that include output columns to avoid accessing the original table.

## Key Claims
- The query planner needs indices to function effectively; programmers must normally add these indices.
- Full table scans should be avoided for large tables; lookups by rowid are significantly faster than full scans but still require scanning content if not using an index on the search column.
- **Covering Indexes**: An index that includes all columns needed for a query (search terms and output) allows SQLite to avoid referencing the original table entirely, roughly doubling speed compared to non-covering indices.
- **Multi-Column Indices**: A single multi-column index can often replace multiple single-column indices, provided one does not keep an index that is a prefix of another.
- **Partial Sorting (Block Sorting)**: If an index satisfies only part of an ORDER BY clause, SQLite performs multiple small sorts rather than one large sort to save CPU cycles and temporary storage.
- **OR Clauses**: For queries with OR-connected terms in the WHERE clause, SQLite examines each term separately using available indices and combines results via a union operation; if no index exists for a term, a full table scan is required.

## Suggested Links
- [Search Documentation](https://www.sqlite.org/queryplanner.html)
