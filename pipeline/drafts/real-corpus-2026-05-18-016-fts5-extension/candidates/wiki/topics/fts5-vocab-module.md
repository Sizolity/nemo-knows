---
title: Fts5 Vocab Module
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/016-fts5-extension.md
confidence: medium
---

# Fts5 Vocab Module

The **fts5vocab** module is a virtual table feature within the SQLite FTS5 extension that allows users to query term statistics without scanning the main content table. This capability provides efficient access to data regarding term frequency and occurrence patterns.

## Usage Types

When querying via the vocab module, three specific modes are supported:
- **row**: Returns term frequency per document.
- **col**: Returns term frequency per column.
- **instance**: Returns exact occurrences of terms.

To retrieve offset values in **instance** mode, the table configuration must include `detail='full'`.

## Relationship to Index Structure

The underlying index data is stored as immutable segment B-trees in shadow tables (e.g., `%_data`, `%_idx`). The vocab module interacts with these structures to generate statistics. This approach complements standard full-text search capabilities, which utilize auxiliary functions like `bm25()`, `highlight()`, and `snippet()` for scoring and presentation.

## Configuration Context

The vocab module operates within the broader FTS5 architecture, which supports:
- Custom **tokenizers** (e.g., `unicode61`, `porter`) configured via the `tokenize` option.
- Different **table types**, including standard tables, contentless tables (`content=''`), and contentless-delete tables.
- Index maintenance options such as `automerge`, `crisismerge`, and `rebuild`.

## Security Note

When using features like secure deletion or specific version upgrades (e.g., 3.42.0+), users should be aware that physical removal of entries may prevent reconstruction of deleted rows, which impacts the consistency of term statistics if not managed carefully.
