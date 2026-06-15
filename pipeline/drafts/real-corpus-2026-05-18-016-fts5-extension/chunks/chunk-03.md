---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---
Chunk Context
Heading path: 4. FTS5 Table Creation and Initialization > 4.4. External Content and Contentless Tables > 4.4.3. External Content Tables
Line range: 841-1217

Local Summary
This section details the mechanics of FTS5 external content tables, where the virtual table derives data from a physical "content table." It explains the SQL query pattern used to fetch values, the necessity of maintaining consistency between the index and source (often via triggers), and the consequences of inconsistency. The text also covers limitations regarding REPLACE conflict handling and provides examples of inconsistent query results when the index and content table diverge. Finally, it discusses `UPDATE`/`DELETE` semantics on contentless tables and the behavior of the `columnsize` option.

Key Claims
- External content FTS5 tables fetch data from a named "content table" using a specific SQL query pattern involving `<content_rowid>` and `<cols>`.
- Users are responsible for keeping the external content FTS5 table synchronized with its source content table.
- Triggers on the content table are recommended to maintain consistency but only handle new/changed rows, not existing ones upon creation.
- The `rebuild` command can be used to discard an FTS index and rebuild it based on the current content table.
- External content tables do not support `REPLACE` conflict handling; such operations result in `ABORT`.
- Contentless tables support `UPDATE` and `DELETE`, but the FTS5 table must be updated *before* the content table to ensure data availability for tokenization.

Entities And Concepts
- External Content Table: An FTS5 virtual table that queries a separate physical table for values.
- Content Table: The physical table (or view) providing data to an external content FTS5 table.
- `content_rowid`: The name of the column in the content table used as the primary key for joining; defaults to "rowid".
- Triggers: Database triggers (`AFTER INSERT`, `AFTER DELETE`, `AFTER UPDATE`) used to keep the FTS index up to date.
- `rebuild` command: Used to clear an FTS index and repopulate it from the content table.
- Contentless Table: An FTS5 table without a content table dependency (context for comparison).
- `columnsize`: Option controlling whether column token counts are stored in a backing table (`0` or `1`).

Procedures And API Details
- **External Content Query Pattern**:
  ```sql
  SELECT <content_rowid>, <cols> FROM <content> WHERE <content_rowid> = ?;
  ```
- **Trigger Example for External Content Synchronization**:
  ```sql
  CREATE TRIGGER t1_ai AFTER INSERT ON t1 BEGIN
    INSERT INTO fts_idx(rowid, b, c) VALUES (new.a, new.b, new.c);
  END;
  CREATE TRIGGER t1_ad AFTER DELETE ON t1 BEGIN
    INSERT INTO fts_idx(fts_idx, rowid, b, c) VALUES('delete', old.a, old.b, old.c);
  END;
  CREATE TRIGGER t1_au AFTER UPDATE ON t1 BEGIN
    INSERT INTO fts_idx(fts_idx, rowid, b, c) VALUES('delete', old.a, old.b, old.c);
    INSERT INTO fts_idx(rowid, b, c) VALUES (new.a, new.b, new.c);
  END;
  ```
- **Rebuilding an Index**: Execute `REINDEX` or equivalent rebuild command on the FTS5 table to discard existing index entries and regenerate them from the content table.
- **Columnsize Configuration**:
  - `columnsize=0`: Omits backing table for token counts; `xColumnSize()` runs slowly by counting tokens on demand.
  - `columnsize=1`: Stores `xColumnSize()` values in a `<name>_docsize` table.
  - Error: Setting `columnsize` to any value other than `0` or `1`.

Nuance Or Contradictions
- **Inconsistency Handling**: If the FTS index and content table are inconsistent, queries without a `MATCH` operator return results from the content table directly (ignoring the index). Queries with a `MATCH` operator use the index; if entries are missing or stale, they may return no rows or incorrect data.
- **Trigger Limitations**: Triggers only react to changes *after* they are created. They cannot copy existing rows from the content table into the FTS index upon creation, leading to an initial state of inconsistency that must be manually resolved or rebuilt.
- **Contentless Table Updates**: Unlike standard tables, `UPDATE` on a contentless FTS5 table is implemented as `DELETE` followed by `INSERT`. Therefore, the FTS5 row must be updated *before* the corresponding content table row to prevent data loss during tokenization.

Candidate Wiki Hints
- **External Content Tables in SQLite**: A guide covering setup, synchronization strategies using triggers, and troubleshooting inconsistencies between the index and source data.
- **FTS5 Index Optimization**: Notes on the `columnsize` option for reducing disk usage versus query performance trade-offs.
- **Locale Support in FTS5**: How to associate locale values with tokens using `fts5_locale()` and `locale=1`.
