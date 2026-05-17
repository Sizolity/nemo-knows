---
kind: source
sources: [raw/llm-wiki.md]
---

# LLM Wiki

## What It Is

A pattern for building personal knowledge bases using LLMs. It is designed as an idea file to be pasted into an LLM Agent (e.g., OpenAI Codex, Claude Code) to communicate the high-level concept. The system functions as an IDE where the user is in charge of sourcing and exploration, the LLM acts as the programmer, and the wiki serves as the codebase.

## Summary

Unlike standard RAG systems that retrieve from raw documents at query time, this approach has the LLM incrementally build and maintain a persistent wiki. The architecture consists of three layers: raw sources (immutable), the wiki (LLM-generated markdown), and the schema (configuration file). Operations include ingesting sources to update the wiki, querying to synthesize answers (which can be filed back as new pages), and periodic linting to health-check the system. Navigation is aided by `index.md` (content catalog) and `log.md` (chronological record).

## Key Claims

- The wiki is a persistent, compounding artifact where knowledge is compiled once and kept current rather than re-derived.
- The LLM should write and maintain all wiki content, including summarizing, cross-referencing, and updating entity pages.
- Answers to queries should be filed back into the wiki to ensure explorations compound in the knowledge base.
- The maintenance burden is reduced because LLMs do not get bored or forget to update cross-references.
- The concept is related in spirit to Vannevar Bush's Memex (1945), where connections between documents are as valuable as the documents themselves.

## Suggested Links

- [Tolkien Gateway](https://tolkiengateway.net/wiki/Main_Page)
- [qmd](https://github.com/tobi/qmd)
