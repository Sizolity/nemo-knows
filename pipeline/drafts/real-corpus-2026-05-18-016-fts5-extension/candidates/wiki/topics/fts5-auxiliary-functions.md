---
title: Fts5 Auxiliary Functions
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

# Fts5 Auxiliary Functions

The SQLite FTS5 (Full-Text Search 5) extension provides a suite of auxiliary functions designed to enhance the usability and analysis of search results. These functions are restricted to full-text queries and operate on virtual tables created by the FTS5 module.

## Function Overview

FTS5 supports specific built-in functions that return values derived from the internal index structure or the BM25 scoring algorithm. These tools allow developers to generate readable output directly within SQL queries.

### Scoring and Ranking
The system utilizes the **BM25** algorithm for relevance ranking. While a hidden `rank` column exists in FTS5 tables to provide fast access to scores, the public interface exposes this via the `bm25()` function. The scoring is inverted, meaning a lower value indicates higher relevance.

### Text Processing
Two primary functions assist in presenting search results:
*   **Highlighting**: The `highlight()` function identifies and marks matching terms within the document text.
*   **Snippets**: The `snippet()` function generates short excerpts of text containing the query terms, useful for displaying previews in search result lists.

### Locale Information
The `fts5_get_locale()` function allows queries to retrieve the current locale settings configured for the FTS5 module, which affects tokenization and matching behavior.

## Usage Context

These functions are typically used in conjunction with:
*   [[fts5-query-syntax]]: To format results based on match locations or proximity.
*   [[index]]: When analyzing term frequencies stored in the index structure.
*   [[query]]: For constructing complex search result sets that include relevance scores and formatted text.

For detailed information on how these functions interact with specific column types, see [[fts5-table-types]].
