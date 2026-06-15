---
title: Index
kind: index
updated: 2026-06-15
---

# Index

The catalogue of every maintained page in the production wiki, organised
by category. This is the entry point for all operations: start here
before any query, ingest, or lint pass.

Development and stability-evaluation artifacts live under `pipeline/`.
Production wiki workflows operate on `wiki/`; short-lived debug or review
artifacts may be written under `tmp/` and cleaned up after use.

When this file outgrows its useful size (rough threshold: it stops
fitting comfortably in an LLM context window), split each category into
its own `index-<category>.md` and have this file link to them.

## Sources

_One page per ingested external document. Each entry uses a standard Markdown
relative link for direct navigation, followed by a one-line description._

- [git-branching](sources/git-branching.md) — Notes on Git's lightweight branch pointers,
  `HEAD`, divergent history, and branch workflow implications.
- [llm-wiki](sources/llm-wiki.md) — Karpathy's LLM-maintained wiki pattern for compounding
  knowledge bases.
- [qwen-llama-cpp](sources/qwen-llama-cpp.md) — Notes on running Qwen models locally with
  llama.cpp, GGUF files, and generation parameters.
- [sqlite-wal](sources/sqlite-wal.md) — Notes on SQLite write-ahead logging, reader/writer
  concurrency, checkpoints, and operational trade-offs.

## Entities

_People, organisations, products, places. One page per entity._

(none yet)

## Concepts

_Ideas, mechanisms, definitions. One page per concept._

(none yet)

## Topics

_Cross-cutting syntheses, comparisons, derived insights — including
high-value query answers filed back from chat._

(none yet)
