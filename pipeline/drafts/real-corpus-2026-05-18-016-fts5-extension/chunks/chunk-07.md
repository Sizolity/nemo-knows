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
