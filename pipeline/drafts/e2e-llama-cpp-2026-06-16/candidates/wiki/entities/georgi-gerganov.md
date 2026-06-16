---
title: Georgi Gerganov
kind: entity
sources:
  - source.md
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

# Georgi Gerganov

Georgi Gerganov is the person who initiated the open-source [Llama Cpp](llama-cpp.md) inference library. The project’s C/​C++ codebase and its focus on local, dependency‑light execution—no GPU server or Python required—originated with his early design choices.

He first released the library in March 2023, targeting Meta’s LLaMA models on consumer laptop CPUs. That release established the single‑binary approach and the goal of efficient CPU‑first inference that remain central to the project.

After the initial launch the library attracted a wide contributor community, and maintenance later moved under the `ggml-org` organization. Gerganov’s original architecture continues to anchor the project even as it has grown to support many model families.

The code he started now underpins popular downstream tools such as Ollama and LM Studio, placing his initial work at the foundation of a broad local LLM inference ecosystem.
