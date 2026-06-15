---
title: LLM Maintenance Pattern
kind: concept
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# LLM Maintenance Pattern

## Definition

A pattern for building personal knowledge bases using LLMs where the model incrementally builds and maintains a persistent wiki of structured, interlinked markdown files rather than relying solely on retrieval-augmented generation ([[RAG]]).

## Workflow

The [[LLM Wiki]] pattern uses an LLM agent to maintain a persistent, structured wiki from raw documents. Instead of re-deriving knowledge on every query, the LLM ingests sources, updates entity pages, and flags contradictions.

The system consists of raw sources, the wiki, and a schema. Operations include ingesting, querying (with answers filed back), and linting. Indexing uses `index.md` and `log.md`.

## Key Claims

- The wiki is a persistent, compounding artifact, not a temporary chat context.
- The LLM writes and maintains the wiki; humans handle sourcing and exploration.
- Answers to queries should be filed back into the wiki as new pages.
- Maintenance cost is near zero for humans because the LLM handles bookkeeping.
- The pattern relates to Vannevar Bush's [[Memex]] concept.

## Tools

- [[Obsidian]]
- [[Obsidian Web Clipper]]
- [[qmd]]
- [[Marp]]
- [[Dataview]]
- [[Git]]

## External References

- [Tolkien Gateway](https://tolkiengateway.net/wiki/Main_Page)
- [qmd](https://github.com/tobi/qmd)
