---
title: Tooling Stack
kind: concept
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# Tooling Stack

The **[[tooling-stack|Tooling Stack]]** defines the software and configuration components required to build and maintain a persistent knowledge base using Large Language Models. It shifts from standard Retrieval-Augmented Generation (RAG) to incrementally building a structured, interlinked markdown repository.

## Core Components

### The Wiki
The primary artifact is a collection of immutable raw sources transformed into LLM-generated markdown files. This layer serves as the persistent memory of the system, allowing for compounding knowledge rather than rediscovering facts on every query.

### The Schema
Configuration files, such as `CLAUDE.md`, define the operational rules and context for the LLM agent. These instructions guide the system in writing, maintaining, and health-checking the wiki content.

### Operations
The stack supports three primary operations:
- **Ingest**: Processes new raw sources into the knowledge base.
- **Query**: Synthesizes answers from existing pages; good responses can be filed back as new entries.
- **Lint**: Performs automated health checks to identify contradictions or orphaned pages.

### Navigation
The system relies on specific entry points for navigation:
- `index.md`: Used for cataloging and browsing content.
- `log.md`: Maintains a chronological record of operations and changes.

## Architecture

The architecture follows a pattern where the LLM handles bookkeeping and maintenance, reducing the human burden to sourcing and asking questions. This approach relates to Vannevar Bush's concept of the Memex, utilizing associative trails to connect ideas. The system is designed to be copy-pasted into an LLM Agent (e.g., OpenAI Codex, Claude Code) to function autonomously.
