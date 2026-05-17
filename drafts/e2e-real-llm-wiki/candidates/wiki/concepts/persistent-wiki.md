---
title: Persistent Wiki
kind: concept
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# Persistent Wiki

A pattern for building personal knowledge bases using Large Language Models (LLMs) that shifts from standard Retrieval Augmented Generation (RAG) to incrementally building a persistent wiki of structured, interlinked markdown files. The system functions as an LLM Agent (e.g., OpenAI Codex, Claude Code) where the AI writes and maintains the wiki while humans handle sourcing and asking questions.

## Architecture

The system operates on three distinct layers:
- **Raw sources**: Immutable input data.
- **The wiki**: Generated markdown files containing structured knowledge.
- **The schema**: A configuration file (e.g., `CLAUDE.md`) defining system behavior.

## Operations

Key operations include:
- **[[ingest]]**: Processing new sources to update the knowledge base.
- **[[query]]**: Synthesizing answers from the wiki, which can be filed back as new pages.
- **[[lint]]**: Health-checking for contradictions or orphan pages.
- **Index**: Content cataloging via `index.md`.
- **Log**: Chronological records via `log.md`.

## Principles

- The wiki is a persistent, compounding artifact rather than a temporary retrieval system.
- Cross-references and contradictions are maintained automatically by the LLM.
- Maintenance burden is near zero for humans as the LLM handles bookkeeping.
- Good answers to queries can be filed back into the wiki as new pages.
- The approach relates to Vannevar Bush's Memex concept regarding associative trails.

## Related Concepts

- [[llm-wiki]]
- [[llm-wiki-core-concept]]
- [[llm-maintenance-pattern]]
- [[tooling-stack]]
- [[nemo-knows-mvp]]
- [[wiki-as-compounding-artifact]]
- [[persistent-wiki-architecture]]
- [[ingest-workflow]]
