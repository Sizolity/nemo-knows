---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

# Chunk Context

**Heading Path:** `9. FTS5 Data Structures > 9.2. The FTS Index (%_idx and %_data tables) > 9.2.4. Segment B-Tree Format > 9.2.4.2. Pagination`
**Line Range:** 3167-3516

# Local Summary

This chunk details the internal storage mechanisms of the FTS5 module, specifically focusing on how segment b-tree data is paginated within the `%_data` table. It explains the header/footer structure of these pages and introduces the `%_idx` table used for indexing leaf pages to enable efficient range queries. The text further describes doclist index formats for handling large rowid lists, the structure of the `%_docsize` table for token counting, and the schema of the `%_content` and `%_config` tables. Finally, it provides a comprehensive comparison between FTS5 and legacy FTS3/4 modules regarding syntax changes, query behavior, and architectural improvements like incremental merging and handling of large instance-lists.

# Key Claims

- Segment b-trees are stored as single blobs in the `%_data` table if small (default < 4000 bytes); otherwise, they are split into pages of approximately 4000 bytes each.
- Page headers contain offsets for rowids and footers, while page footers store varints representing key offsets.
- The `%_idx` table indexes leaf pages to allow efficient searching via term prefixes without traversing internal nodes.
- Doclist indexes are b-trees stored in the `%_data` table, used primarily when a doclist spans more than 4 segment b-tree leaf pages.
- The `%_docsize` table stores token counts per column as packed varints to support ranking functions like TF-IDF.
- FTS5 replaces large single-record instance-lists (used in FTS3/4) with distributed records across multiple database entries, improving memory usage and query speed for complex queries.
- FTS5 introduces `ORDER BY rank` for relevance sorting and supports custom auxiliary functions via an API.

# Entities And Concepts

- **FTS5**: The full-text search module being documented.
- **Segment B-Tree**: An internal data structure where segments are stored, split into pages when exceeding size limits.
- **%_data Table**: Stores the actual blob data for segment b-trees and doclist index nodes/leaves.
- **%_idx Table**: Indexes leaf pages of the segment b-tree to enable efficient term-based lookups.
- **Doclist Index**: A secondary b-tree structure within FTS5 that indexes rowids associated with a specific term or prefix.
- **%_docsize Table**: Stores token counts for each column in an FTS5 document.
- **%_content Table**: Holds the actual content values of documents; schema varies based on `locale` and `UNINDEXED` options.
- **%_config Table**: Stores persistent configuration options like `crisismerge`, `pgsz`, and `usermerge`.
- **Instance-List**: A list of token occurrences identified by `(rowid, column, position)` tuples.

# Procedures And API Details

**Pagination Logic:**
1. Check if the segment b-tree blob is smaller than 4000 bytes.
2. If yes, store as a single entry in `%_data`.
3. If no, split into pages of ~4000 bytes.
4. Apply format modifications: non-prefix-compressed first keys, uncompressed rowids before the first key.
5. Add 4-byte header (rowid offset, footer offset) and variable footer (key offsets).

**Querying Leaf Pages:**
To find a leaf for segment `i` containing term `t`:
```sql
SELECT pgno FROM %_idx WHERE segid=$i AND term>=$t ORDER BY term LIMIT 1
```

**Doclist Index Creation:**
- Add index only if doclist spans > 4 segment b-tree leaf pages.
- Store in `%_data` as a b-tree (leaves and internal nodes).
- First byte of page is flags (`0x00` for root, `0x01` otherwise).
- Subsequent bytes store packed varints: leftmost child page number, smallest rowid, delta-encoded subsequent children.

**%_docsize Population:**
- For each row in FTS5 table, insert into `%_docsize`.
- Store token counts as a blob of packed varints corresponding to columns left-to-right.
- Unindexed columns get a value of zero.

**FTS3/4 Migration Syntax Changes:**
- Change module name: `fts3`/`fts4` -> `fts5`.
- Remove type/constraint info from column definitions (e.g., `VARCHAR`).
- Replace `notindexed=` with `UNINDEXED` keyword.
- Replace `matchinfo=fts3` with `columnsize=0`.
- Remove `compress=`, `uncompress=`, `languageid=` options.
- Replace `docid` alias with `rowid`.

# Nuance Or Contradictions

- **Filter Behavior:** In FTS3/4, a column filter in the MATCH operator overrides an outer filter (inner filter wins). In FTS5, both filters are applied sequentially; if they conflict (e.g., searching for 'b: string' while filtering on column 'a'), zero rows are returned.
- **Locale Handling:** The `%_content` table schema changes dynamically. If `locale=1` is set and a column is `UNINDEXED`, the corresponding "l*" column does not exist unless `contentless_unindexed=1` is specified, in which case only UNINDEXED columns get "c*" entries but no "l*" entries.
- **Memory Usage:** FTS3/4 often loads entire instance-lists into memory, whereas FTS5 loads them incrementally to reduce peak allocation.

# Candidate Wiki Hints

1. **FTS5 Pagination Format**: A detailed breakdown of how FTS5 manages storage efficiency by splitting large segment b-trees into fixed-size pages with specific header/footer structures.
2. **FTS5 vs FTS3/4 Migration Guide**: A comparison table and checklist for updating SQL schemas and queries when migrating from legacy full-text search modules to FTS5, highlighting syntax incompatibilities.
3. **Doclist Index Architecture**: Explanation of the secondary indexing strategy in FTS5 that handles large document lists by creating b-tree structures within the `%_data` table.
4. **Token Counting with %_docsize**: How FTS5 supports ranking algorithms requiring document length normalization via the `%_docsize` shadow table.
