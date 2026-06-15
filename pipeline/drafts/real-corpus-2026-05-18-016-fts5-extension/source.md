---
title: SQLite FTS5 Extension Overview
kind: source
sources:
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

## What It Is

The SQLite FTS5 (Full-Text Search 5) extension is a virtual table module introduced in SQLite 3.9.0 that enables efficient full-text searching within database tables. Unlike standard SQL tables, FTS5 tables are implemented via extensions and support advanced query capabilities such as prefix matching, phrase queries, proximity searches (`NEAR`), boolean logic, and relevance ranking. It includes built-in auxiliary functions for highlighting matches, generating snippets, and scoring results using BM25 algorithms. The module also supports custom tokenizers, external content tables, and virtual tables for querying index statistics via `fts5vocab`.

## Summary

FTS5 allows users to create virtual tables with one or more columns that can be populated and queried like standard SQL tables. Queries utilize `MATCH` operators, equality checks, and table-valued functions to retrieve rows containing specific terms. The system supports complex query syntax including barewords (unquoted alphanumeric terms), phrases, negations (`^`), column filters (`colname:`), and boolean operators (`AND`, `OR`, `NOT`) with defined precedence rules.

The module manages its internal index structure using segment B-trees stored in shadow tables (`%_data`, `%_idx`, `%_docsize`, etc.). These structures are optimized through incremental merging, controlled by configuration options like `automerge` and `crisismerge`. Maintenance commands such as `rebuild`, `optimize`, and `integrity-check` allow users to manage index consistency and performance. FTS5 distinguishes between standard tables (storing content), contentless tables (index-only storage), and contentless-delete tables (supporting updates on index-only tables).

## Key Claims

- **Installation**: FTS5 is included in SQLite 3.9.0+ via the `--enable-fts5` configure option or as a loadable extension.
- **Query Syntax**: Queries are case-independent by default. Barewords are restricted to alphanumeric characters, underscores, and specific Unicode ranges. Prefix matching uses the `*` suffix on tokens.
- **Boolean Operators**: Implicit `AND` exists between whitespace-separated phrases; explicit operators (`AND`, `OR`, `NOT`) follow precedence: implicit AND > NOT > OR. Parentheses can override precedence.
- **Table Types**:
  - *Standard*: Stores original column values and supports standard `INSERT`/`UPDATE`.
  - *Contentless* (`content=''`): Stores only index entries; prevents reading original column values (except `rowid`) and disables standard `INSERT`/`UPDATE`.
  - *Contentless-Delete*: Combines contentless storage with support for `DELETE`, `UPDATE`, and `INSERT OR REPLACE`.
- **External Content**: FTS5 tables can derive data from a physical "content table" using a specific SQL query pattern. Users must maintain consistency, often via triggers. External content tables do not support `REPLACE` operations.
- **Tokenizers**: Configured via the `tokenize` option. Built-in options include `unicode61`, `ascii`, `porter`, and `trigram`. Custom tokenizers can be registered via C APIs (`fts5_tokenizer_v2`).
- **Auxiliary Functions**: Restricted to full-text queries. Includes `bm25()` (scoring, inverted so lower is better), `highlight()`, `snippet()`, and `fts5_get_locale()`. The hidden `rank` column provides faster access to BM25 scores.
- **Index Maintenance**:
  - `automerge`: Controls automatic merging of level-0 B-trees (default 4).
  - `crisismerge`: Forces immediate merge when a threshold (default 16) is reached.
  - `deletemerge`: Controls when B-trees containing tombstones are eligible for merging (default 10%).
  - `rebuild`: Discards the index and rebuilds from the content table.
- **Security**: The `secure-delete` option physically removes entries upon deletion, preventing reconstruction of deleted rows but breaking compatibility with FTS5 versions prior to 3.42.0.
- **fts5vocab Module**: A virtual table module for querying term statistics without scanning the main table. Supports 'row' (term frequency per document), 'col' (per column), and 'instance' (exact occurrences) types. Requires `detail='full'` for offset values in 'instance' mode.
- **Internal Storage**: Index data is stored as immutable segment B-trees in `%_data` and indexed by leaf pages in `%_idx`. Large doclists are split into secondary B-tree structures. Token counts are stored in `%_docsize`.

## Suggested Links

none
