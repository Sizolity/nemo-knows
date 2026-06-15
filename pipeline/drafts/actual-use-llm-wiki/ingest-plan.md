---
kind: topic
sources: [raw/llm-wiki.md]
status: draft
---

# Ingest Plan

## Source Summary
- Shifts from standard RAG retrieval to a persistent, compounding wiki maintained by an LLM.
- Architecture consists of immutable raw sources, an LLM-owned wiki directory, and a schema configuration file.
- Operations include ingesting sources to update entities, querying to synthesize answers, and linting for consistency.

## Candidate Wiki Pages
- wiki/topics/persistent-wiki-architecture.md — Defines the core distinction between retrieval-based systems and the proposed persistent wiki model.
- wiki/concepts/llm-maintenance-pattern.md — Documents the human-LLM collaboration workflow for bookkeeping and synthesis.
- wiki/sources/llm-wiki-pattern.md — Archives the methodology source for future reference and adaptation.

## Suggested Links
- https://tolkiengateway.net/wiki/Main_Page
- https://github.com/tobi/qmd

## Review Checklist
- [ ] Verify schema file conventions are defined before first ingestion.
- [ ] Ensure index.md and log.md are treated as operational tools, not candidate content pages.
