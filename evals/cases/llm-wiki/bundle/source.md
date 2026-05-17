---
kind: source
sources:
  - raw/llm-wiki.md
confidence: medium
---

# LLM Wiki

## What It Is

A pattern for building personal knowledge bases using LLMs where the model incrementally builds and maintains a persistent wiki of structured, interlinked markdown files rather than relying solely on retrieval-augmented generation (RAG).

## Summary

The LLM Wiki pattern uses an LLM agent to maintain a persistent, structured wiki from raw documents. Instead of re-deriving knowledge on every query like RAG systems, the LLM ingests sources, updates entity pages, and flags contradictions.

## Key Claims

- The wiki is a persistent, compounding artifact, not a temporary chat context.
- The LLM writes and maintains the wiki; humans handle sourcing and exploration.
- Answers to queries should be filed back into the wiki as new pages.

## Suggested Links

- Obsidian
- Git
