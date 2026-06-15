## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

## Chunk Context
This chunk introduces the SQLite FTS5 Extension, detailing its purpose as a virtual table module for full-text search. It covers basic usage patterns (creation, querying), query syntax components (strings, phrases, prefixes, NEAR groups), and column filtering mechanisms.

## Local Summary
FTS5 enables efficient full-text searching within SQLite databases by creating virtual tables with one or more columns. Users can populate these tables like standard SQL tables and execute queries using `MATCH` operators, equality checks, or table-valued functions. The system supports advanced search types including prefix matching, phrase queries, proximity searches (NEAR), and boolean logic.

## Key Claims
- FTS5 is included in SQLite version 3.9.0+ via the `--enable-fts5` configure option or as a loadable extension.
- Queries are case-independent by default unless specified otherwise.
- Results can be sorted by relevance using the `rank` auxiliary function or column.
- Barewords (unquoted terms) are restricted to alphanumeric characters, underscores, and specific Unicode ranges.

## Entities And Concepts
- **FTS5**: Full-text search virtual table module.
- **Virtual Table**: A special type of table in SQLite implemented via extensions.
- **Tokenizer**: Module that extracts tokens from text strings (e.g., default, Porter, Trigram).
- **Auxiliary Functions**: Scalar functions accessible to FTS5 queries (e.g., `highlight`, `snippet`).
- **Bareword**: Unquoted string used in FTS5 queries containing only allowed characters.
- **Phrase**: Ordered list of tokens forming a query segment.
- **NEAR Query**: Proximity search grouping phrases within N tokens.

## Procedures And API Details
### Creating an FTS5 Table
```sql
CREATE VIRTUAL TABLE email USING fts5(sender, title, body);
```
*Note: Types, constraints, or PRIMARY KEY declarations are not allowed.*

### Basic Query Syntax
- **Match Operator**: `SELECT * FROM table WHERE table MATCH 'term';`
- **Equality Operator**: `SELECT * FROM table WHERE table = 'term';`
- **Table-valued Function**: `SELECT * FROM table('term');`

### Highlighting Matches
```sql
SELECT highlight(email, 2, '<b>', '</b>') FROM email('fts5');
```

### Query Syntax Components (BNF Summary)
- `<phrase>`: Single string or concatenated phrases (`+`).
- `<neargroup>`: `NEAR ( <phrase> <phrase> ... [, N] )`.
- `<query>`: Combines phrases, negations (`^`), column filters (`colname :`), and boolean operators.

## Nuance Or Contradictions
- **Bareword Limitations**: Characters outside ASCII letters, digits, underscore, or specific Unicode ranges must be quoted. Future versions may expand bareword rules.
- **Prefix Token Behavior**: The `*` suffix marks a token as a prefix match. Placing `*` inside double quotes passes it to the tokenizer, which may discard it depending on the tokenizer used.
- **Column Filter Scope**: Column filters apply to all phrases within an expression; nested filters only further restrict columns and cannot re-enable previously filtered ones.

## Candidate Wiki Hints
- **Page: SQLite FTS5 Extension Overview**
  Summarize installation, basic creation, and query mechanics.
- **Page: FTS5 Query Syntax Reference**
  Document BNF rules for phrases, NEAR groups, prefixes, and column filters.
- **Page: FTS5 Tokenizers and Configuration**
  Detail available tokenizers (Unicode61, Ascii, Porter, Trigram) and options like `detail`, `tokendata`.

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---
Chunk Context
This chunk covers section 3.7 regarding FTS5 Boolean Operators (precedence, parentheses, implicit AND) and sections 4 through 4.4.2 detailing the creation of FTS5 virtual tables, including column options, configuration options (tokenize, prefix, content), specific tokenizer implementations (unicode61, ascii, porter, trigram), and the distinction between standard, contentless, and contentless-delete table types.

Local Summary
The document explains how to construct complex full-text queries using boolean operators (`AND`, `OR`, `NOT`) with their specific precedence rules and grouping via parentheses. It then transitions to table creation syntax, defining column names/options and configuration options for tokenization, prefix indexing, and content storage strategies. Detailed subsections describe the behavior of built-in tokenizers (unicode61, ascii, porter, trigram) and their specific arguments. Finally, it defines three table modes: standard (stores content), contentless (empty string for `content` option), and contentless-delete (adds update support to contentless tables).

Key Claims
- Boolean operators have a strict precedence order: implicit AND > NOT > OR. Parentheses can override this.
- Implicit AND operators exist between whitespace-separated phrases or NEAR groups but are never inserted inside parentheses.
- The `UNINDEXED` column option prevents column data from being added to the FTS index, effectively making it invisible to MATCH queries.
- Prefix indexes speed up queries for tokens starting with a specific string by storing separate indices for token prefixes of specified lengths.
- Tokenizers are configured via the `tokenize` option, which accepts a bareword or quoted SQL literal specifying the tokenizer name and its arguments.
- The trigram tokenizer allows substring matching rather than exact token matching, supporting GLOB and LIKE patterns (unless using ESCAPE).
- Contentless tables (`content=''`) store only index entries, saving space but preventing reading of original column values (except rowid) and disabling standard INSERT/UPDATE logic.
- Contentless-delete tables combine the storage efficiency of contentless tables with support for DELETE, UPDATE (full replacement), and `INSERT OR REPLACE` statements.

Entities And Concepts
- **FTS5 Boolean Operators**: Logical connectors (`AND`, `OR`, `NOT`) used in full-text queries.
- **Implicit AND**: Unspoken logical conjunction between adjacent phrases or NEAR groups separated by whitespace.
- **Prefix Index**: A secondary index structure optimizing range scans for tokens starting with a specific character sequence.
- **Tokenizer**: Modules that process text into tokens (e.g., `unicode61`, `porter`, `trigram`).
- **Contentless Table**: An FTS5 table where the original row content is not stored in the table itself (`content=''`).
- **External Content**: A reference to an external database object used to retrieve column values for an FTS5 table.
- **Diacritics**: Accents on characters, handled differently depending on the tokenizer and `remove_diacritics` option.

Procedures And API Details
- **Creating a Table with Prefix Indexes**: Set the `prefix` configuration option to a list of integers (e.g., `'2 3'`) or multiple `prefix=` arguments.
    - Example: `CREATE VIRTUAL TABLE ft USING fts5(a, b, prefix='2 3');`
- **Configuring Tokenizers**: Use the `tokenize` option with a quoted string specifying the tokenizer and its arguments.
    - Examples:
        - Default Porter: `tokenize = 'porter'`
        - Custom Tokenizer: `tokenize = "unicode61 remove_diacritics 0 tokenchars '-_'"`
- **Creating a Contentless Table**: Set `content` to an empty string.
    - Example: `CREATE VIRTUAL TABLE ft USING fts5(a, b, c, content='');`
- **Creating a Contentless-Delete Table**: Combine `content=''` with `contentless_delete=1`.
    - Example: `CREATE VIRTUAL TABLE ft USING fts5(a, b, c, content='', contentless_delete=1);`

Nuance Or Contradictions
- **Trigram Case Sensitivity**: The trigram tokenizer supports case-insensitive matching by default. If `case_sensitive` is set to 1, GLOB queries are supported, but LIKE queries are not.
- **Diacritic Handling Bug**: In the unicode61 tokenizer, setting `remove_diacritics` to "1" fails to correctly remove diacritics from certain multi-codepoint characters (e.g., combining marks), a known technical bug that can be fixed by setting it to "2" but may impact backwards compatibility.
- **Token Length Limits**: If the trigram tokenizer is used with `detail=none` or `detail=column`, full-text queries cannot contain tokens longer than 3 unicode characters.

Candidate Wiki Hints
- FTS5 Boolean Operator Precedence
- Creating Prefix Indexes in SQLite
- Configuring Tokenizers for FTS5
- Contentless vs. External Content Tables
- Trigram Substring Matching

## chunk-03

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

## chunk-04

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

## chunk-05

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

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

# Chunk Context

**Heading:** 7. Extending FTS5 > 7.1. Custom Tokenizers
**Line Range:** 1907–2233
**Source Path:** `raw/web/corpus-2026-05-18/016-fts5-extension.md`

# Local Summary

This chunk details the API for implementing custom tokenizers in SQLite FTS5, defining the `fts5_tokenizer_v2` structure and its required callbacks (`xCreate`, `xDelete`, `xTokenize`). It explains the lifecycle of a tokenizer instance, the meaning of tokenization flags (e.g., `FTS5_TOKENIZE_QUERY`, `FTS5_TOKENIZE_DOCUMENT`), and how to handle locale arguments. The section also covers synonym support via the `FTS5_TOKEN_COLOCATED` flag, outlining three implementation strategies for synonyms. Finally, it briefly introduces custom auxiliary functions as a related extension mechanism.

# Key Claims

- A custom tokenizer requires three callbacks: `xCreate`, `xDelete`, and `xTokenize`.
- The `fts5_tokenizer_v2` struct includes an `iVersion` field (always 2) and function pointers for the lifecycle and tokenization logic.
- Tokenization flags include `FTS5_TOKENIZE_QUERY`, `FTS5_TOKENIZE_PREFIX`, `FTS5_TOKENIZE_DOCUMENT`, and `FTS5_TOKENIZE_AUX`.
- The `xTokenize()` method may receive a locale buffer (`pLocale`) of size `nLocale`; if both are 0, the default locale is used.
- Synonyms are supported by invoking `xToken()` with the `FTS5_TOKEN_COLOCATED` flag for subsequent tokens in a sequence.
- Three methods exist for synonym support: mapping synonyms to a single token, expanding query terms into multiple synonyms during querying, or indexing all synonyms of a term.
- Method 1 is efficient but lacks prefix support; Method 3 supports prefixes but uses more disk space; Method 2 offers a middle ground.
- Custom auxiliary functions are implemented using the `fts5_extension_function` type and registered via `xCreateFunction()`.

# Entities And Concepts

- **fts5_tokenizer_v2**: Structure defining a custom tokenizer implementation.
- **xCreate**: Callback to create and initialize a tokenizer instance.
- **xDelete**: Callback to destroy a tokenizer instance; guaranteed once per successful `xCreate`.
- **xTokenize**: Callback invoked to tokenize input text; receives flags and locale info.
- **xToken**: Internal callback provided by the tokenizer implementation to return tokens.
- **FTS5_TOKENIZE_* Flags**: Masks indicating why tokenization is requested (document, query, prefix, aux).
- **FTS5_TOKEN_COLOCATED**: Flag indicating a synonym for the previous token.
- **fts5_api.xCreateTokenizer_v2()**: Method to register a new tokenizer with FTS5.
- **Synonyms**: Alternative forms of tokens that should be treated as equivalent during search.

# Procedures And API Details

### Registering a Custom Tokenizer

1. Populate an `fts5_tokenizer_v2` struct with:
   - `xCreate`: Constructor for the tokenizer instance.
   - `xDelete`: Destructor for the tokenizer instance.
   - `xTokenize`: Main tokenization logic, invoking `xToken()` for each token found.
2. Call `fts5_api.xCreateTokenizer_v2()` with a pointer to the struct and optional user data.
3. On success (`SQLITE_OK`), the tokenizer is registered; on failure, cleanup does not occur.

### Tokenizer Lifecycle

- **Initialization**: `xCreate()` is called once per table insertion/update or query start.
  - Arguments: User-provided pointer, array of tokenizer arguments (null-terminated strings), output handle pointer.
- **Tokenization**: `xTokenize()` is called zero or more times with text and flags.
  - Returns `SQLITE_OK` on completion or error code if an error occurs.
- **Cleanup**: `xDelete()` is called exactly once per successful `xCreate()`.

### Handling Locale

- Arguments `pLocale` and `nLocale` specify the locale string (e.g., `"en_US"`).
- If `pLocale == NULL`, use default locale (`nLocale` ignored).
- The `pLocale` buffer is not null-terminated.

### Synonym Implementation

To support synonyms, invoke `xToken()` with `FTS5_TOKEN_COLOCATED` for each synonym:

```c
xToken(pCtx, 0, "first", ...);
xToken(pCtx, FTS5_TOKEN_COLOCATED, "1st", ...);
xToken(pCtx, FTS5_TOKEN_COLOCATED, "one", ...);
```

- First call must not use `FTS5_TOKEN_COLOCATED`.
- Multiple synonyms allowed per token via sequential calls.

# Nuance Or Contradictions

- Legacy `fts5_tokenizer` lacks locale support and uses `xCreateTokenizer()` instead of `_v2()`.
- Synonym handling is inefficient if provided during both document and query tokenization; should be restricted to one context.
- Method 1 (synonym mapping) does not support prefix queries well unless synonyms include prefixes explicitly.

# Candidate Wiki Hints

- **Custom Tokenizer API**: Document the `fts5_tokenizer_v2` structure and required callbacks.
- **FTS5 Flags**: Create a reference for tokenization flags (`FTS5_TOKENIZE_*`) and synonym flag (`FTS5_TOKEN_COLOCATED`).
- **Synonym Strategies**: Summarize the three approaches to synonym support with trade-offs (space vs. query performance).
- **Locale Handling**: Explain how to pass locale strings to `xTokenize()` and when defaults apply.

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

## Chunk Context
This chunk documents the **Custom Auxiliary Functions API Overview** for FTS5 (Section 7.2.1). It details how custom functions invoked during an FTS5 query can query metadata about the table, current row, tokens, and specific phrases in the query. The text explains available APIs, tokenization concepts, phrase iteration methods, and includes examples demonstrating matches within rows, including handling of column filters and NEAR filters.

## Local Summary
The auxiliary function API allows implementations to inspect FTS5 tables and queries without needing external data access. Key capabilities include querying row/column counts, accessing current row text/rowid, counting tokens in columns (current vs. total), inspecting query phrases and tokens, and iterating through phrase matches within a specific row. Two primary methods for iterating phrase matches are described: one that iterates per phrase (`xPhraseFirst`/`xPhraseNext`) and one that provides a flat array of all matches sorted by occurrence (`xInstCount`/`xInst`). The chunk clarifies how column filters (e.g., `y:`) and proximity filters (`NEAR`) affect which matches are reported.

## Key Claims
- Auxiliary functions are invoked for rows matching the query (e.g., containing token "ab").
- Column numbering starts at 0 from left to right, excluding the implicit `rowid`.
- Tokens are extracted based on the configured tokenizer (default `unicode61`); punctuation is removed and casing may be folded.
- `xColumnSize` returns tokens in the current row; `xColumnTotalSize` returns tokens across all table rows.
- Queries can contain multiple phrases (e.g., `"ab"`, `"cd ef gh"`).
- Matches reported by phrase iteration APIs exclude those filtered out by column filters or `NEAR` constraints.
- The flat array method (`xInstCount`/`xInst`) collates all matches into a single list sorted by occurrence order within the row.

## Entities And Concepts
- **FTS5**: SQLite full-text search virtual table module.
- **Auxiliary Function**: A custom function invoked during FTS5 query execution to access internal data.
- **Token**: Basic unit of text after tokenization (words, possibly folded/lowercased).
- **Phrase**: A contiguous sequence of tokens in a query.
- **Column Filter**: Syntax like `column_name:` restricting matches to specific columns.
- **NEAR Filter**: Proximity constraint limiting matches within a certain distance.
- **Rowid**: Internal unique identifier for FTS5 rows (not counted as a user column).

## Procedures And API Details
- **xRowCount**: Returns total rows in the FTS5 table.
- **xColumnCount**: Returns number of user-defined columns in the table.
- **xRowid**: Retrieves the `rowid` of the current row being visited.
- **xColumnText(column_index)**: Gets text stored in a specific column for the current row.
- **xColumnSize(column_index)**: Counts tokens in a specified column for the current row.
- **xColumnTotalSize(column_index)**: Counts total tokens in a specified column across all table rows.
- **xPhraseCount**: Returns number of phrases in the current query.
- **xPhraseSize(phrase_index)**: Returns token count for a specific phrase (0-indexed).
- **xQueryToken(phrase_index, token_index)**: Retrieves text of a specific token within a phrase.
- **xPhraseFirst(phrase_index)** / **xPhraseNext(phrase_index)**: Iterate matches for a single phrase in the current row, returning `(column_number, token_offset)`.
- **xInstCount** / **xInst**: Provide access to a flat array of all phrase matches in the current row. Each element is `(phrase_number, column_number, token_offset)`.

## Nuance Or Contradictions
- **Filtering Behavior**: Matches filtered by column filters (e.g., `y:`) or `NEAR` filters are explicitly excluded from iteration results (`xPhraseFirst`, `xInst`). This is not a contradiction but a filtering rule.
- **Column Numbering**: The `rowid` is never counted as a column index; only user-declared columns start at 0.
- **Token Counting**: Token counts differ between current row (`xColumnSize`) and entire table (`xColumnTotalSize`).
- **Phrase Ordering**: In the flat array method (`xInst`), matches are sorted by occurrence within the row, not necessarily by phrase number or query order.

## Candidate Wiki Hints
- **Page**: `SQLite_FTS5_Custom_Auxiliary_Functions_API`
  - Focus: Comprehensive guide to FTS5 auxiliary function capabilities, including API reference and usage patterns for token/phrase inspection.
- **Page**: `SQLite_FTS5_Query_Match_Iteration`
  - Focus: Explaining how to iterate through phrase matches in a row, comparing per-phrase iteration vs. flat array approaches, with examples of filter interactions.

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

## Chunk Context
This chunk documents section **7.2.2. Custom Auxiliary Functions API Reference** of the SQLite FTS5 extension documentation. It provides the C struct definition `struct Fts5ExtensionApi` and detailed descriptions for each callback function pointer contained within it. The content details how application developers can access metadata, tokenize text, iterate through phrase matches, and manage auxiliary data during a full-text search query execution.

## Local Summary
The section defines the interface for custom functions that interact with an FTS5 table during a `MATCH` query. It outlines methods to retrieve column counts, row IDs, token counts (global or per-row), specific text content, and phrase details. It also covers iteration APIs for phrase instances, token retrieval within phrases, and locale handling. Special attention is given to performance caveats regarding tables created with `detail=none` or `detail=column`.

## Key Claims
- The `iVersion` field in `struct Fts5ExtensionApi` is currently always set to 4.
- Functions marked as "Below this point are iVersion>=3 only" include `xQueryToken` and `xInstToken`.
- Functions marked as "Below this point are iVersion>=4 only" include `xColumnLocale` and `xTokenize_v2`.
- The `xSetAuxdata()` API allows storing a pointer for retrieval by the same extension function during a single query, with an optional cleanup callback.
- Iteration APIs (`xPhraseFirst`, `xPhraseNext`) are generally faster than `xInstCount`/`xInst` but behave differently with contentless tables or specific detail options.
- The `xInstToken()` API may be slow for prefix tokens unless FTS5 is configured to pre-collect data via the `insttoken` option.

## Entities And Concepts
- **struct Fts5ExtensionApi**: The C struct defining the interface for custom auxiliary functions.
- **Fts5Context**: A handle representing the context of the current FTS5 query.
- **xUserData**: Callback to retrieve the user data pointer passed during function registration.
- **Phrase Iteration**: Mechanisms (`xPhraseFirst`, `xPhraseNext`, `xPhraseFirstColumn`) to traverse phrase matches within a row or across columns.
- **Locale Support**: Capability via `xColumnLocale` and `xTokenize_v2` to associate locale strings with column values.
- **Auxiliary Data**: Temporary storage (`xSetAuxdata`, `xGetAuxdata`) available per query invocation.

## Procedures And API Details
### xUserData
Returns a copy of the pointer passed to `xCreateFunction()` when registering the extension function.

### Column Statistics and Text Access
- **xColumnCount**: Returns the number of columns in the FTS5 table.
- **xColumnTotalSize(iCol, pnToken)**: Returns total token count for column `iCol` (or all tokens if `iCol < 0`). Returns `SQLITE_RANGE` if `iCol` is out of bounds.
- **xColumnSize(iCol, pnToken)**: Returns token count for the current row in column `iCol`. Inefficient with `columnsize=0`.
- **xColumnText(iCol, pz, pn)**: Retrieves text for column `iCol`. Returns `SQLITE_RANGE` if `iCol` is invalid.

### Phrase Analysis
- **xPhraseCount**: Returns the number of phrases in the current query expression.
- **xPhraseSize(iPhrase)**: Returns the number of tokens in phrase `iPhrase`. Returns 0 if `iPhrase` is invalid.
- **xInstCount(pnInst)**: Sets `*pnInst` to the total number of occurrences of all phrases in the current row. Returns `SQLITE_OK` on success or `SQLITE_NOMEM` on error. Returns 0 for contentless tables (`content=` option). Slow with `detail=none` or `detail=column`.
- **xInst(iIdx, piPhrase, piCol, piOff)**: Details phrase match at index `iIdx`. Outputs phrase number, column index, and token offset. Returns `SQLITE_RANGE` if `iIdx` is invalid. Slow with `detail=none` or `detail=column`.
- **xPhraseFirst(iPhrase, &iter, &iCol, &iOff)**: Initiates iteration over instances of phrase `iPhrase`. Outputs column and offset.
- **xPhraseNext(&iter, &iCol, &iOff)**: Advances iterator for phrase `iPhrase`.
- **xPhraseFirstColumn(iPhrase, &iter, &iCol)**: Iterates over columns containing at least one instance of phrase `iPhrase`. More efficient than `xInst`/`xInstCount` for `detail=column` tables.
- **xPhraseNextColumn(&iter, &iCol)**: Advances iterator for column-based phrase iteration.

### Token and Query Execution
- **xTokenize(pText, nText, pCtx, xToken)**: Tokenizes text using the table's tokenizer. Invokes callback `xToken` for each token.
- **xQueryPhrase(iPhrase, pUserData, xCallback)**: Executes a query equivalent to `WHERE ftstable MATCH $p` for phrase `iPhrase`. Calls `xCallback` for each matched row. Returns `SQLITE_RANGE` if `iPhrase` is invalid. Returns `SQLITE_OK` on success or propagated error code on failure.
- **xQueryToken(iPhrase, iToken, ppToken, pnToken)**: Retrieves token text at index `iToken` of phrase `iPhrase`. Returns tokenizer output (includes embedded nulls if `tokendata=1`). Returns `SQLITE_RANGE` if indices are invalid.
- **xInstToken(iIdx, iToken, ppToken, pnToken)**: Retrieves token text at index `iToken` within phrase hit `iIdx`. May require scanning the full-text index for prefix tokens unless configured via `insttoken`.

### Auxiliary Data Management
- **xSetAuxdata(pAux, xDelete)**: Saves pointer `pAux` as auxiliary data. Invokes `xDelete` if existing data had a delete callback or on query completion.
- **xGetAuxdata(bClear)**: Retrieves stored auxiliary data pointer. If `bClear` is non-zero, sets data to NULL without invoking `xDelete`.

### Locale Support
- **xColumnLocale(iCol, pz, pn)**: Retrieves the locale string associated with column `iCol` if set via `fts5_locale()`. Returns `SQLITE_RANGE` if `iCol` is invalid.

### Advanced Tokenization
- **xTokenize_v2(pText, nText, pLocale, nLocale, pCtx, xToken)**: Same as `xTokenize` but accepts a locale string and length for the tokenizer.

## Nuance Or Contradictions
- **Performance Variance**: Many APIs (`xInstCount`, `xInst`, `xPhraseFirst`, etc.) are explicitly noted as "quite slow" when used with tables created using `detail=none` or `detail=column`. This is likely due to the lack of pre-computed positional data in these configurations.
- **Contentless Tables**: For tables created with the `content=` option, APIs returning row-level details (like `xInstCount`) always return 0 or iterate over an empty set because there are no document rows to inspect.
- **Data Representation**: Output text from token access functions (`xQueryToken`, `xInstToken`) is raw tokenizer output, not necessarily the original UTF-8 document text. This includes embedded null bytes (`0x00`) and trailing data for tables with `tokendata=1`.
- **Prefix Token Overhead**: Using `xInstToken` on prefix tokens can force a secondary scan of the index. Configuration options (`insttoken` option or `fts5_insttoken()` function) exist to pre-collect this data at the cost of increased memory usage and slower queries that do not use `xInstToken`.

## Candidate Wiki Hints
- **Page**: FTS5 Custom Auxiliary Functions API
  - **Summary**: Comprehensive reference for accessing internal FTS5 state during query execution.
  - **Key Sections**: Column stats, Phrase iteration strategies, Token retrieval, Auxiliary data lifecycle.
- **Concept**: FTS5 Detail Options Performance Impact
  - **Summary**: How `detail=none` and `detail=column` affect the efficiency of auxiliary function calls.
- **Concept**: FTS5 Prefix Token Optimization
  - **Summary**: Managing memory vs speed trade-offs for prefix token queries using `xInstToken`.

## chunk-09

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

## chunk-10

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

