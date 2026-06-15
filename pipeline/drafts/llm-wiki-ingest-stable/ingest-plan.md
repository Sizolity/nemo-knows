---
kind: topic
sources: [raw/llm-wiki.md]
status: draft
---

# Ingest Plan

## Source Summary
- Shifts from standard RAG retrieval to an LLM-maintained persistent wiki that accumulates knowledge across queries.
- Architecture relies on three layers: immutable raw sources, LLM-generated wiki pages, and a schema configuration file.
- Operations include ingesting sources to update the wiki, querying with synthesis, and periodic linting for health checks.
- Human role is curation and questioning; LLM handles maintenance, cross-referencing, and bookkeeping.

## Candidate Wiki Pages
- wiki/sources/llm-wiki-pattern.md — Documents the core pattern for using LLMs to maintain a personal knowledge base.
- wiki/concepts/persistent-wiki.md — Defines the core concept of incremental knowledge accumulation versus re-derivation.
- wiki/topics/architecture.md — Details the three-layer structure (Raw, Wiki, Schema) and their specific responsibilities.

## Suggested Links
- https://tolkiengateway.net/wiki/Main_Page
- https://github.com/tobi/qmd

## Review Checklist
- [ ] Define specific schema conventions (e.g., CLAUDE.md or AGENTS.md) for the LLM agent.
- [ ] Verify tooling compatibility (Obsidian plugins, CLI search engines) with the proposed workflow.
- [ ] Establish linting rules for detecting contradictions and orphan pages.
