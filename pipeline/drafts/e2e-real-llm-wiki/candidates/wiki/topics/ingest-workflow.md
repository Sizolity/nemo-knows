---
title: Ingest Workflow
kind: topic
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# Ingest Workflow

The ingest workflow is a core operation within the **[[llm-wiki]]** system, designed to process new sources and incrementally build a **[[persistent-wiki]]**. Unlike standard RAG systems that rediscover knowledge from scratch, this approach treats the wiki as a **compounding-artifact**, where the LLM writes and maintains structured, interlinked markdown files while humans handle sourcing.

The workflow operates on three layers: immutable raw sources, the generated wiki content, and a configuration schema. When new material is introduced, the system processes it to update the knowledge base without losing historical context. This process often involves navigating the **[[tooling-stack]]** to ensure proper integration with existing structures like **[[llm-wiki-core-concept]]**.

Key outcomes of the ingest phase include:
- Shifting from temporary retrieval to persistent storage.
- Automatically maintaining cross-references and resolving contradictions via **[[lint]]**.
- Creating associative trails reminiscent of Vannevar Bush's Memex concept.

The resulting knowledge is accessible through an **[[index]]** for cataloging and tracked chronologically in the **[[log]]**. Good answers synthesized during this phase can be filed back into the wiki as new pages, further enriching the **[[persistent-wiki-architecture]]**. For detailed configuration, see the **[[llm-maintenance-pattern]]**, which guides the agent on how to handle these updates alongside standard **[[query]]** operations.
