---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---
## Chunk Context
**Heading:** 5. Auxiliary Functions (Lines 1218–1581)
**Scope:** Defines auxiliary functions for FTS5, details four built-in functions (`bm25`, `highlight`, `snippet`, `fts5_get_locale`), discusses sorting by results via the `rank` column, and introduces special INSERT commands (`automerge`, `crisismerge`, `delete`, `delete-all`).

## Local Summary
This section explains that FTS5 auxiliary functions are similar to SQL scalar functions but operate exclusively within full-text queries (using `MATCH` or `LIKE/GLOB` with trigram tokenizer). Their results depend on both arguments and the current match context. The text details four built-in functions: `bm25()` for scoring, `highlight()` for markup, `snippet()` for fragments, and `fts5_get_locale()` for locale retrieval. It further explains how to sort results using the hidden `rank` column for performance and describes configuration options (`automerge`, `crisismerge`) and maintenance commands (`delete`, `delete-all`) for managing the internal b-tree index structure.

## Key Claims
- Auxiliary functions are restricted to full-text queries on FTS5 tables.
- The `bm25()` function returns a value where lower numbers indicate better matches (inverted compared to standard BM25).
- The `rank` column is a hidden alias for `bm25()` that offers faster sorting than calling the function directly.
- The `automerge` parameter controls how many level-0 b-trees are merged automatically; default is 4.
- The `crisismerge` option forces an immediate merge when a threshold (default 16) of b-trees exists on a level, potentially causing slow write operations.
- The `delete` command requires inserting specific values into the row to accurately remove index entries from contentless tables.

## Entities And Concepts
- **FTS5 Auxiliary Functions:** Specialized SQL functions usable only within full-text queries.
- **bm25():** Returns a numeric score for match accuracy (lower is better).
- **highlight():** Wraps matched terms in specified markup tags within a column's text.
- **snippet():** Extracts a short fragment of text containing matches, with configurable length and boundary markers.
- **fts5_get_locale():** Retrieves the locale associated with a stored value if the table supports locales.
- **rank:** A hidden column acting as an alias for the default `bm25()` score to optimize sorting performance.
- **automerge:** Configuration option controlling incremental merging of b-trees.
- **crisismerge:** Configuration option triggering immediate merging when a critical number of b-trees accumulates.
- **delete / delete-all:** Special INSERT commands for removing or clearing entries from the full-text index in contentless tables.

## Procedures And API Details
**Invoking Auxiliary Functions:**
The table name is the first argument. Subsequent arguments vary by function.
*   **Syntax Example:** `SELECT highlight(ft, 2, '<b>', '</b>') FROM ft WHERE ft MATCH 'fts5'`
*   **bm25() Weighting:** Arguments specify weights for columns left-to-right (e.g., `ORDER BY bm25(email, 10.0, 5.0)`).

**highlight() Parameters:**
1.  Integer column index (0-based).
2.  Text before match.
3.  Text after match.
*   **Behavior:** Overlapping phrases share markers; non-overlapping instances get separate markers.

**snippet() Parameters:**
1.  Column index (-1 for auto-selection).
2.  Pre-markup text.
3.  Post-markup text.
4.  Start/End indicator text (if fragment is not at the boundary).
5.  Maximum token count (0 < count ≤ 64).

**Configuring Rank Column:**
*   **Per-query:** `WHERE ft MATCH ? AND rank MATCH 'bm25(10.0, 5.0)' ORDER BY rank`
*   **Table-level:** Set via FTS5 rank configuration option during table creation.

**Managing B-Trees (INSERT Commands):**
*   **Automerge:** Insert `'automerge', <value>` to set how many b-trees merge per transaction. Max value is 16; default is 4. Value 0 disables auto-merge.
*   **Crisismerge:** Insert `'crisismerge', <value>` to set the threshold for immediate merging. Default is 16.
*   **Delete (Contentless):** `INSERT INTO ft(ft, rowid, col...) VALUES('delete', <rowid>, ...)` where column values must match current row data exactly.
*   **Delete All:** `INSERT INTO ft(ft) VALUES('delete-all')`.

## Nuance Or Contradictions
- **BM25 Scoring Direction:** Standard BM25 assigns higher scores to better matches. FTS5 inverts this (multiplies by -1) so that `ORDER BY bm25()` returns best matches first (ascending order). Without the inversion, descending order would be required.
- **Rank Column Performance:** Using `rank` is faster than `bm25()` for sorting, but they are logically equivalent when called with identical arguments.
- **Crisismerge Impact:** Triggering a crisis merge can cause an INSERT/UPDATE/DELETE to take significantly longer as it performs an immediate, heavy merge operation.

## Candidate Wiki Hints
- **FTS5 Auxiliary Functions Overview**
- **BM25 Scoring and Ranking in SQLite**
- **FTS5 Highlight and Snippet Functionality**
- **Managing FTS5 B-Tree Merging (Automerge/Crisismerge)**
- **Contentless Table Maintenance Commands**
