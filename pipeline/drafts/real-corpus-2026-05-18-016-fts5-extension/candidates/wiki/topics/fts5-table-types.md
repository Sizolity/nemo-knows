---
title: Fts5 Table Types
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

# Fts5 Table Types

The SQLite FTS5 (Full-Text Search 5) extension introduces virtual table modules that enable efficient full-text searching. Unlike standard SQL tables, FTS5 tables are implemented via extensions and support advanced capabilities such as prefix matching, phrase queries, proximity searches (`NEAR`), boolean logic, and relevance ranking using BM25 algorithms.

FTS5 supports three distinct table types based on how they store content and handle modifications:

## Standard Tables
Standard FTS5 tables store the original column values alongside index entries. They support standard `INSERT` and `UPDATE` operations to populate both the content and the index simultaneously. This is the default behavior for most use cases where the full text content needs to be retrievable.

## Contentless Tables
Contentless tables (specified with `content=''`) store only index entries rather than the original column values. Accessing the original text columns is disabled, except for the implicit `rowid`. Consequently, standard `INSERT` and `UPDATE` commands are disabled because there is no content to update; modifications must be performed via external triggers or mechanisms that directly manipulate the index.

## Contentless-Delete Tables
Contentless-delete tables combine the storage characteristics of contentless tables with support for deletion operations. They allow `DELETE`, `UPDATE`, and `INSERT OR REPLACE` commands on index-only tables. This type is useful when applications need to manage deletions without storing redundant data in the virtual table itself.

## External Content
FTS5 tables can derive their data from a physical "content table" using a specific SQL query pattern defined during table creation. Users must maintain consistency between the external content and the FTS5 index, often via triggers. It is important to note that external content tables do not support `REPLACE` operations.

## Query Syntax
Queries for these tables utilize `MATCH` operators, equality checks, and table-valued functions. The system supports complex query syntax including barewords (unquoted alphanumeric terms), phrases, negations (`^`), column filters (`colname:`), and boolean operators (`AND`, `OR`, `NOT`) with defined precedence rules. Case-insensitive matching is the default behavior.

## Auxiliary Functions
Restricted to full-text queries, auxiliary functions include `bm25()` for scoring (where lower values are better), `highlight()`, and `snippet()`. The hidden `rank` column provides faster access to BM25 scores. Additional tools like `fts5_get_locale()` allow for locale-specific operations.

## Index Maintenance
The internal index structure uses segment B-trees stored in shadow tables (`%_data`, `%_idx`, `%_docsize`, etc.). Maintenance is controlled by configuration options such as:
- **automerge**: Controls automatic merging of level-0 B-trees (default 4).
- **crisismerge**: Forces immediate merge when a threshold (default 16) is reached.
- **deletemerge**: Controls when B-trees containing tombstones are eligible for merging (default 10%).

Commands like `rebuild`, `optimize`, and `integrity-check` allow users to manage index consistency and performance. The `secure-delete` option physically removes entries upon deletion, preventing reconstruction of deleted rows but breaking compatibility with FTS5 versions prior to 3.42.0.
