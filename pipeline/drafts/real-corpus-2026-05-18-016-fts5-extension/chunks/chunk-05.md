---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

## Chunk Context
This chunk (lines 1582-1906) covers sections **6.5 through 6.14** of the SQLite FTS5 documentation, detailing specific configuration options (`deletemerge`, `insttoken`, `pgsz`, `rank`, `secure-delete`, `usermerge`) and maintenance commands (`integrity-check`, `merge`, `optimize`, `rebuild`). It concludes with **Section 7**, outlining how to extend FTS5 via C APIs for custom tokenizers and auxiliary functions.

## Local Summary
The text describes mechanisms for managing b-tree lifecycle in FTS5, specifically regarding tombstones for deleted rows in contentless tables (`deletemerge`), integrity verification (`integrity-check`), and merging strategies (`merge`, `optimize`). It also covers security implications of deletion (`secure-delete`) and extension points allowing developers to register custom tokenizers and functions using the `fts5_api` structure.

## Key Claims
- **Deletemerge**: Controls when b-trees containing deleted row markers ("tombstones") become eligible for merging; default threshold is 10%.
- **Integrity-check**: Verifies index consistency against the content table (if not contentless) or internal structures; returns `SQLITE_CORRUPT_VTAB` on failure.
- **Merge/Optimize Logic**: Positive merge parameters respect `usermerge` counts for eligibility, while negative parameters force a deep reorganization regardless of level structure.
- **Secure-delete**: Setting this to 1 physically removes entries from the database file upon deletion/updates to prevent reconstruction of deleted rows, but breaks compatibility with FTS5 versions earlier than 3.42.0.
- **Extension API**: New tokenizers and auxiliary functions must be registered via C callbacks provided through the `fts5_api` structure obtained via the SQL function `fts5()`.

## Entities And Concepts
- **Contentless-delete tables**: Tables where rows have no content stored in the main table, relying solely on the FTS index.
- **Tombstone marker**: A marker attached to a b-tree indicating that associated tokens belong to a deleted row; these are omitted from query results but kept until merging.
- **B-tree levels**: Logical grouping of b-trees within the FTS5 index structure, relevant for merge eligibility.
- **xInstToken API**: An API method requiring extra data collection for prefix queries when enabled via `insttoken`.
- **fts5_api**: The C structure containing function pointers to register new tokenizers (`xCreateTokenizer`) and auxiliary functions (`xCreateFunction`).

## Procedures And API Details
- **Setting deletemerge threshold**:
  ```sql
  INSERT INTO ft(ft, rank) VALUES('deletemerge', 15);
  ```
- **Enabling xInstToken data collection**:
  ```sql
  INSERT INTO ft(ft, rank) VALUES('insttoken', 1);
  ```
- **Running integrity check**:
  ```sql
  INSERT INTO ft(ft) VALUES('integrity-check');
  -- With rank 1 for external content tables
  INSERT INTO ft(ft, rank) VALUES('integrity-check', 1);
  ```
- **Forcing merge with specific page count**:
  ```sql
  INSERT INTO ft(ft, rank) VALUES('merge', 500);
  ```
- **Optimizing via merge steps**:
  1. Start deep merge: `INSERT INTO ft(ft) VALUES('optimize');` (or equivalent negative merge logic).
  2. Monitor `sqlite3_total_changes()`; stop when difference < 2.
- **Enabling secure-delete**:
  ```sql
  INSERT INTO ft(ft, rank) VALUES('secure-delete', 1);
  ```
- **Obtaining fts5_api pointer (C)**:
  ```c
  sqlite3_stmt *pStmt;
  fts5_api *pRet = 0;
  // Prepare "SELECT fts5(?1)"
  sqlite3_bind_pointer(pStmt, 1, (void*)&pRet, "fts5_api_ptr", NULL);
  sqlite3_step(pStmt);
  ```

## Nuance Or Contradictions
- **Merge Parameter Signs**: A positive merge parameter respects the `usermerge` count threshold and b-tree levels, whereas a negative parameter ignores `usermerge`, flattens levels, and merges aggressively.
- **Secure-delete Compatibility**: Enabling `secure-delete` (value 1) creates a file format incompatibility with FTS5 versions prior to 3.42.0, necessitating a `rebuild` command if downgrading or accessing via older tools.
- **Integrity Check Scope**: By default, `integrity-check` only validates internal index structures for contentless tables; validating against the external content table requires explicitly setting `rank` to 1.

## Candidate Wiki Hints
- Page: **FTS5 Configuration Options** (Summary of `deletemerge`, `insttoken`, `pgsz`, `rank`, `secure-delete`, `usermerge`).
- Page: **FTS5 Maintenance Commands** (Details on `integrity-check`, `merge`, `optimize`, `rebuild` strategies).
- Page: **FTS5 Security and Deletion Policies** (Explanation of tombstones, secure-delete, and data reconstruction risks).
- Page: **Extending FTS5 with C APIs** (Guide to using `fts5_api`, registering custom tokenizers and functions).
