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
