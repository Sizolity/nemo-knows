---
kind: source
sources: [raw/llm-wiki.md]
---

# LLM Wiki

## What It Is

A pattern for building personal knowledge bases using LLMs. Designed as an idea file to be copy-pasted to an LLM Agent (e.g., OpenAI Codex, Claude Code) to communicate the high-level concept while the agent builds specifics in collaboration with the user.

## Summary

The core idea differs from standard RAG systems where the LLM rediscover knowledge from scratch on every question. Instead, the LLM incrementally builds and maintains a persistent wiki—a structured, interlinked collection of markdown files.

The system consists of three layers:
1. **Raw sources:** Immutable source documents.
2. **The wiki:** A directory of LLM-generated markdown files (summaries, entity pages, synthesis).
3. **The schema:** A configuration file (e.g., CLAUDE.md) defining structure and workflows.

Operations include ingesting new sources, querying the wiki (with answers filed back as pages), and linting for health checks. Navigation is supported via `index.md` (catalog) and `log.md` (chronological record).

## Key Claims

- The wiki is a persistent, compounding artifact where cross-references and contradictions are pre-flagged.
- The LLM handles all bookkeeping (summarizing, cross-referencing, filing), while the human handles sourcing and questioning.
- Maintenance costs are near zero for humans because LLMs handle updates across dozens of pages.
- The pattern is modular and abstract, applicable to personal tracking, research, reading books, business teams, or competitive analysis.
- The wiki functions as a git repo of markdown files, providing version history and collaboration for free.

## Suggested Links

- [Tolkien Gateway](https://tolkiengateway.net/wiki/Main_Page)
- [qmd](https://github.com/tobi/qmd)

Other mentioned tools and concepts without explicit URLs include Obsidian, Obsidian Web Clipper, Marp, Dataview, and Vannevar Bush's Memex.
