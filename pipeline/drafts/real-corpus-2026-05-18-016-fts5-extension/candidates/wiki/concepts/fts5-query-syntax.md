---
title: Fts5 Query Syntax
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

# Fts5 Query Syntax

The SQLite FTS5 (Full-Text Search 5) extension enables efficient full-text searching within database tables using specialized `MATCH` operators. Queries are case-independent by default and support complex syntax for retrieving rows containing specific terms.

## Basic Terms

Queries utilize **barewords**, which are unquoted alphanumeric terms, underscores, and specific Unicode ranges. Prefix matching is supported by appending the `*` suffix to a token (e.g., `term*`).

## Operators and Logic

FTS5 supports various operators and logic constructs:

- **Phrases**: Enclosed in double quotes (e.g., `"exact phrase"`).
- **Negations**: Use the `^` prefix (e.g., `^term`) to exclude terms.
- **Column Filters**: Specify columns using the `colname:` syntax (e.g., `title:hello`).
- **Boolean Operators**:
  - Implicit `AND` exists between whitespace-separated phrases.
  - Explicit operators include `AND`, `OR`, and `NOT`.
  - Precedence rules apply: implicit AND > NOT > OR.
  - Parentheses can be used to override precedence.

## Auxiliary Functions

The system includes built-in auxiliary functions for advanced query capabilities:

- **Scoring**: The `bm25()` function scores results using BM25 algorithms (inverted so lower is better). The hidden `rank` column provides faster access to these scores.
- **Highlighting**: The `highlight()` function marks matches within text.
- **Snippets**: The `snippet()` function generates excerpts from matching documents.

## Index Maintenance

The internal index structure uses segment B-trees stored in shadow tables (`%_data`, `%_idx`, etc.). Maintenance commands such as `rebuild`, `optimize`, and `integrity-check` manage index consistency. Configuration options like `automerge` control automatic merging of level-0 B-trees, while `crisismerge` forces immediate merges when thresholds are reached.

## Table Types

FTS5 distinguishes between different virtual table types based on storage and update capabilities:

- **Standard**: Stores original column values and supports standard `INSERT`/`UPDATE`.
- **Contentless** (`content=''`): Stores only index entries; prevents reading original column values (except `rowid`) and disables standard `INSERT`/`UPDATE`.
- **Contentless-Delete**: Combines contentless storage with support for `DELETE`, `UPDATE`, and `INSERT OR REPLACE`.

## External Content

FTS5 tables can derive data from a physical "content table" using a specific SQL query pattern. Users must maintain consistency, often via triggers. External content tables do not support `REPLACE` operations.
