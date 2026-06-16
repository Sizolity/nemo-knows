---
kind: topic
sources: [pipeline/raw/web/llama-cpp.md]
status: draft
---

# Ingest Plan

## Source Summary
- llama.cpp is an open‑source C/C++ library and CLI ecosystem for local LLM inference, started by Georgi Gerganov to run Meta’s LLaMA on a laptop CPU.
- It runs on commodity hardware, supports CPU and GPU backends (Metal, CUDA, Vulkan, …), and enables hybrid CPU/GPU offloading via the `-ngl` flag.
- The project reads models in the single‑file GGUF format (successor to GGML) and distributes quantised variants (Q4_K_M, Q8_0, …).
- Important tools include `llama-cli` for interactive/batch use and `llama-server` with OpenAI‑compatible HTTP endpoint.

## Candidate Wiki Pages
- wiki/sources/llama-cpp.md — raw distilled description of the project and its ecosystem.
- wiki/entities/llama-cpp.md — the project entity, capturing purpose, history, maintainer, and community.
- wiki/entities/georgi-gerganov.md — the primary creator of llama.cpp.
- wiki/concepts/gguf.md — the GGUF container format, its metadata, and its role in CPU‑first distribution.
- wiki/concepts/llama-cpp-backends.md — overview of supported hardware backends and layer‑offloading.
- wiki/concepts/quantization.md — quantisation schemes used by llama.cpp and their quality‑size trade‑offs.

## Suggested Links
- https://github.com/ggml-org/llama.cpp
- shingo-docs/refs/llm-quantization-gptq-awq-gguf.md
- wiki/sources/qwen-llama-cpp.md

## Review Checklist
- [ ] Verify that the llama.cpp entity page uses the canonical lowercase, dotted name `llama.cpp`.
- [ ] Confirm that the GGUF concept page correctly describes the transition from GGML and the single‑file design.
- [ ] Check that all cross‑references (refs/llm-quantization-… and wiki/sources/qwen-llama-cpp.md) exist at the specified paths.
- [ ] Ensure the hardware backends page matches the current backends (Metal, CUDA, Vulkan, etc.) and covers hybrid offloading.
- [ ] Validate the tool descriptions (`llama-cli`, `llama-server`) against the latest project documentation.
