---
kind: source
sources:
  - fixture:technical-doc
confidence: medium
---

# Local Model Ingest Pipeline

## What It Is

A local CLI pipeline that renders prompts, calls a local llama.cpp model, cleans model output, reviews generated artifacts, and applies approved wiki edits.

## Summary

The pipeline separates generation from approval. Draft commands produce `source.md` and `ingest-plan.md`; review commands produce `apply-plan.md`; evaluation commands score the artifacts; approved apply writes only after explicit approval.

## Key Claims

- Prompt rendering should use explicit braced variables.
- Raw model output should be preserved next to cleaned Markdown.
- Deterministic review and eval steps reduce reliance on prompt quality alone.

## Suggested Links

- llama.cpp
- GGUF
- Qwen
