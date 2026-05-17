---
kind: topic
sources: [web:https://qwen.readthedocs.io/en/v3.0/run_locally/llama.cpp.html]
status: draft
---

# Ingest Plan

## Source Summary

- The source is official Qwen documentation for local llama.cpp inference.
- It describes obtaining llama.cpp programs, downloading or converting Qwen GGUF
  models, and running `llama-cli` or `llama-server`.
- It covers sampling, speed, chat templates, context management, and
  thinking-mode workarounds.

## Candidate Wiki Pages

- wiki/sources/qwen-llama-cpp-guide.md — Source page for the official local
  inference guide.
- wiki/concepts/gguf-model-format.md — Concept page for GGUF in local inference.
- wiki/concepts/jinja-chat-template.md — Concept page for llama.cpp chat
  templates and Qwen thinking controls.
- wiki/topics/local-qwen-llama-cpp-inference.md — Topic page for the full local
  Qwen inference workflow.

## Suggested Links

- llama.cpp
- Qwen
- GGUF
- Jinja
- reasoning control

## Review Checklist

- Verify whether the current local llama.cpp build supports every documented
  parameter before turning guide examples into defaults.
- Keep Qwen official examples distinct from `nemo` project-specific profile
  choices.
