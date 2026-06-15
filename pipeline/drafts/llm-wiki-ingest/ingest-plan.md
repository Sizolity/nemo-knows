---
kind: topic
sources: [raw/llm-wiki.md]
status: draft
---

# Ingest Plan

## Source Summary
*   Proposes an LLM-driven architecture for maintaining a persistent, structured knowledge base distinct from standard RAG systems.
*   Defines three layers: immutable raw sources, LLM-generated wiki content, and a schema configuration file.
*   Outlines operational workflows for ingesting sources, querying the wiki, and periodic linting to maintain health.

## Candidate Wiki Pages
- wiki/topics/rag-vs-wiki.md — Explains the core distinction between retrieval-augmented generation and incremental wiki building.
- wiki/concepts/maintenance-workflow.md — Details the ingest, query, and lint operations for keeping the knowledge base current.
- wiki/sources/llm-wiki-pattern.md — Documents the high-level pattern described in the raw source for future reference.

## Suggested Links
- [Tolkien Gateway](https://tolkiengateway.net/wiki/Main_Page)
- [qmd](https://github.com/tobi/qmd)

## Review Checklist
- [ ] Verify schema conventions align with chosen LLM agent capabilities.
- [ ] Confirm candidate pages cover core concepts without duplicating index or log files.
- [ ] Validate that suggested links resolve correctly and are relevant to the wiki structure.
