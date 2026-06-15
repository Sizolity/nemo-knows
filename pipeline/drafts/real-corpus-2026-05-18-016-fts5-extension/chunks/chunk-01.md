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
