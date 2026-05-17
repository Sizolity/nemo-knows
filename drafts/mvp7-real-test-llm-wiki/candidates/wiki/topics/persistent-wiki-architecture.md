---
title: Persistent Wiki Architecture
kind: topic
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# Persistent Wiki Architecture

A pattern for building personal knowledge bases using LLMs where the model incrementally builds and maintains a persistent wiki of structured, interlinked markdown files rather than relying solely on retrieval-augmented generation ([[RAG]]).

## Overview

The [[LLM Wiki]] pattern utilizes an LLM agent to maintain a persistent, structured wiki from raw documents. The system architecture consists of raw sources, the wiki, and a schema. Instead of re-deriving knowledge on every query, the LLM ingests sources, updates entity pages, and flags contradictions.

## Operations

Core operations include ingesting, querying (with answers filed back), and linting. Indexing typically utilizes `index.md` and `log.md`.

## Principles

*   **Persistence:** The wiki is a persistent, compounding artifact, not a temporary chat context.
*   **Agency:** The LLM writes and maintains the wiki; humans handle sourcing and exploration.
*   **Documentation:** Answers to queries should be filed back into the wiki as new pages.
*   **Maintenance:** Maintenance cost is near zero for humans because the LLM handles bookkeeping.
*   **Concept:** The pattern relates to Vannevar Bush's [[Memex]] concept.

## Ecosystem

Tools like [[Obsidian]], [[Git]], [[qmd]], [[Marp]], and [[Dataview]] support the workflow. The [[Obsidian Web Clipper]] assists in ingestion.

## References

*   [[Tolkien Gateway]]
*   [[LLM Wiki]]
