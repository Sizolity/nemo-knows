---
kind: source
sources:
  - raw/llm-wiki.md
---

# LLM Wiki

## What It Is
A pattern for building personal knowledge bases using LLMs. Designed to be copy-pasted to an LLM Agent (e.g., OpenAI Codex, Claude Code). It shifts from standard RAG to incrementally building a persistent wiki of structured, interlinked markdown files. The LLM writes and maintains the wiki while humans handle sourcing and asking questions.

## Summary
The system incrementally builds and maintains a persistent wiki rather than rediscovering knowledge from scratch on every query. It consists of three layers: Raw sources (immutable), The wiki (LLM-generated markdown), and The schema (configuration file like CLAUDE.md). Operations include Ingest (processing new sources), Query (synthesizing answers that can be filed back), and Lint (health-checking for contradictions or orphan pages). Navigation relies on `index.md` for content cataloging and `log.md` for chronological records.

## Key Claims
- The wiki is a persistent, compounding artifact rather than a temporary retrieval system.
- Cross-references and contradictions are maintained automatically by the LLM.
- Maintenance burden is near zero for humans as the LLM handles bookkeeping.
- Good answers to queries can be filed back into the wiki as new pages.
- The approach relates to Vannevar Bush's Memex concept regarding associative trails.

## Suggested Links
- [Tolkien Gateway](https://tolkiengateway.net/wiki/Main_Page)
- [qmd](https://github.com/tobi/qmd)
