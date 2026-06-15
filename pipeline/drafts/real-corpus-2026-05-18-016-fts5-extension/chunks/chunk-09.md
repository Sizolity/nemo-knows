---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---
Chunk Context
This chunk covers Section 8 ("The fts5vocab Virtual Table Module") and the beginning of Section 9 ("FTS5 Data Structures"). It details how to create virtual tables associated with FTS5 indexes for querying term statistics, describes the three types of `fts5vocab` tables (`row`, `col`, `instance`), explains shadow table structures created by FTS5, and defines the binary formats for varints, keys, doclists, and position lists used internally.

Local Summary
The fts5vocab module allows direct querying of FTS5 index data without scanning the main table. It supports three modes: 'row' (term statistics per document), 'col' (term statistics per column), and 'instance' (exact term occurrences with rowid/column/offset). The chunk also introduces the underlying storage architecture, including shadow tables (`%_data`, `%_idx`, `%_config`, `%_docsize`, `%_content`) and the varint encoding scheme used for compact binary storage of keys and position lists.

Key Claims
- The `fts5vocab` module is part of FTS5 and requires FTS5 to be available.
- An `fts5vocab` table must be in the same database as its associated FTS5 table unless created in the `temp` database.
- Three `fts5vocab` types exist: 'row', 'col', and 'instance'.
- Creating an `fts5vocab` table with three arguments (database, table, type) is required for non-temp databases; two arguments suffice for temp or when the DB matches.
- The `fts5vocab` "row" type returns term frequency counts (`doc`, `cnt`) per document.
- The `fts5vocab` "col" type distinguishes terms by column name and reports counts per column.
- The `fts5vocab` "instance" type requires the FTS5 table to be created with `detail='full'` to return offset values; if `detail='none'`, offsets are NULL.
- FTS5 uses shadow tables (specifically `%_data`, `%_idx`, etc.) to store persistent data, which users should not access directly.
- FTS5 stores the index as a series of immutable "segment b-trees" merged over time.
- Keys in the index are stored with a common-prefix compression scheme.
- Doclists store rowids using delta-encoded varints and position lists (poslist) for term offsets.

Entities And Concepts
- fts5vocab: A virtual table module for querying FTS5 index data.
- Types: 'row', 'col', 'instance'.
- Shadow tables: `%_data`, `%_idx`, `%_config`, `%_docsize`, `%_content`.
- Segment b-trees: Immutable structures holding index entries, merged incrementally.
- Varint: A variable-length integer encoding used for compact storage.
- Doclist: A packed array of varints representing term occurrences (rowids).
- Poslist: A structure within a doclist identifying term offsets per column.
- Prefix indexes: Additional indices keyed with prefixes (e.g., "1do" for "document").

Procedures And API Details
- Create `fts5vocab` row table: `CREATE VIRTUAL TABLE ft1_v USING fts5vocab('ft1', 'row');`
- Create `fts5vocab` col table: `CREATE VIRTUAL TABLE ft2_v USING fts5vocab(ft2, col);`
- Create `fts5vocab` instance table: `CREATE VIRTUAL TABLE ft3_v USING fts5vocab(ft3, instance);`
- Cross-database creation (requires temp db): `CREATE VIRTUAL TABLE temp.ft1_v USING fts5vocab(main, 'ft1', 'row');`
- FTS5 shadow table creation pattern (internal):
  - `CREATE TABLE %_data(id INTEGER PRIMARY KEY, block BLOB);`
  - `CREATE TABLE %_idx(segid, term, pgno, PRIMARY KEY(segid, term)) WITHOUT ROWID;`
  - `CREATE TABLE %_config(k PRIMARY KEY, v) WITHOUT ROWID;`
  - `CREATE TABLE %_docsize(id INTEGER PRIMARY KEY, sz BLOB);`
  - `CREATE TABLE %_content(id INTEGER PRIMARY KEY, c0, c1...);`

Nuance Or Contradictions
- Offset values in the "instance" type are NULL if the FTS5 table was created with `detail='col'` or `detail='none'`.
- The `%_docsize` shadow table is absent if the "columnsize" option is set to 0.
- `%_content` is absent for contentless or external content FTS5 tables.
- Specifying three arguments for a non-temp database results in an error; two are required unless specifying a specific attached database name.

Candidate Wiki Hints
- Page: FTS5Vocab Module (Overview of types and creation syntax)
- Page: FTS5 Internal Storage (Shadow tables, segment b-trees, varint format)
