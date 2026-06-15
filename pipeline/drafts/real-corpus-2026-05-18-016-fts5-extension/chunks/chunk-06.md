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
