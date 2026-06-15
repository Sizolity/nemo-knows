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
