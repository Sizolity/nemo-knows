---
kind: topic
sources: [raw/web/corpus-2026-05-18/016-fts5-extension.md]
status: draft
---

# Ingest Plan

## Source Summary
- The document details the SQLite FTS5 virtual table module, covering full-text query syntax (boolean operators, phrases, NEAR groups), table creation modes (standard, contentless, external content), and advanced configuration options.
- It extensively documents extension APIs for custom tokenizers and auxiliary functions, including lifecycle management, synonym support, and phrase iteration strategies.
- The text explains internal storage structures such as segment b-trees, shadow tables (`%_data`, `%_idx`), and binary formats (varints, doclists).
- Maintenance procedures are covered, including integrity checks, index merging strategies (`automerge`, `crisismerge`, `deletemerge`), and rebuilding indexes.

## Candidate Wiki Pages
- wiki/sources/fts5-extension.md — Consolidated overview of the FTS5 module installation, version requirements, and basic usage patterns.
- wiki/concepts/fts5-query-syntax.md — Reference for query components including BNF rules, boolean operators, implicit AND, column filters, and NEAR groups.
- wiki/concepts/fts5-tokenizers.md — Documentation of built-in tokenizers (unicode61, ascii, porter, trigram), configuration options (`tokenize`, `prefix`), and custom tokenizer API registration.
- wiki/topics/fts5-table-types.md — Comparison and usage guide for standard, contentless, external content, and contentless-delete table modes with synchronization strategies.
- wiki/topics/fts5-auxiliary-functions.md — Overview of built-in functions (`highlight`, `snippet`, `bm25`) and the API structure for registering custom auxiliary functions.
- wiki/topics/fts5-vocab-module.md — Guide to creating and querying `fts5vocab` virtual tables for term statistics and internal index inspection.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that all candidate page slugs follow the immediate child directory rule (`wiki/sources/`, `wiki/concepts/`, `wiki/topics/`).
- [ ] Ensure "External Content Tables" content is merged into the table types concept rather than a separate source page.
- [ ] Confirm that custom tokenizer and auxiliary function API details are consolidated into their respective concept pages without creating nested directories.
