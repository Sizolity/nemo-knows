---
title: Ingest
kind: concept
created: 2026-05-16
updated: 2026-05-16
sources:
  - raw/llm-wiki.md
  - wiki/sources/llm-wiki.md
tags: [workflow, wiki-maintenance]
confidence: high
---

# Ingest

Ingest is the workflow for integrating a new raw source into the wiki.

During ingest, the LLM reads a source from `raw/`, writes a source summary, updates relevant concept, entity, or topic pages, refreshes the index, and records the change in the log. A single source may touch many wiki pages because the goal is integration, not just summarisation.

The source remains immutable. The wiki changes so that future queries can reuse the compiled knowledge instead of re-deriving it from raw text.

## Related

- [[persistent-wiki]]
- [[llm-wiki-core-concept]]
- [[query]]
- [[lint]]
