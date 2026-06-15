---
kind: topic
sources: [raw/llm-wiki.md]
status: draft
---

# Ingest Plan

## Source Summary
*   Proposes an LLM-driven architecture for maintaining a persistent, structured wiki distinct from standard RAG systems.
*   Defines a three-layer stack: immutable raw sources, LLM-maintained wiki markdown, and a schema configuration file.
*   Outlines operational workflows for ingesting sources, querying the wiki, and periodic linting/health checks.

## Candidate Wiki Pages
*   `wiki/concepts/persistent-wiki.md` — Captures the core distinction between RAG retrieval and incremental wiki synthesis.
*   `wiki/topics/ingest-workflow.md` — Documents the specific steps for adding sources, updating entities, and logging changes.
*   `wiki/concepts/tooling-stack.md` — Catalogs recommended external tools like Obsidian, Dataview, Marp, and qmd for maintenance.

## Suggested Links
*   https://tolkiengateway.net/wiki/Main_Page
*   https://github.com/tobi/qmd

## Review Checklist
*   [ ] Validate schema file conventions against specific LLM agent capabilities (e.g., Claude Code vs. Codex).
*   [ ] Confirm directory structure supports chosen domain (e.g., text-only vs. image-heavy sources).
