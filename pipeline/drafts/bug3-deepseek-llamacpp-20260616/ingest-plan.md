---
kind: topic
sources: [pipeline/raw/web/llama-cpp.md]
status: draft
---

# Ingest Plan

## Source Summary
- Distilled technical reference for the llama.cpp project: open-source library for efficient LLM inference on commodity hardware.
- Covers origins, GGUF model format, hardware backends (CPU, Metal, CUDA, Vulkan, hybrid offloading), and tooling (`llama-cli`, `llama-server`).
- Explicitly recorded for ingest testing of entity pages; canonical project name is lowercase‑dotted `llama.cpp`.

## Candidate Wiki Pages
- wiki/sources/llama-cpp.md — distilled source page capturing the raw technical description.
- wiki/entities/llama-cpp.md — the open‑source project entity, its goals, scope, and community.
- wiki/entities/georgi-gerganov.md — creator of llama.cpp, first release, and ongoing role.
- wiki/concepts/gguf.md — the single‑file model format, metadata, quantization, and contrast with multi‑file schemes.
- wiki/concepts/llama-cpp-hardware-backends.md — CPU execution, GPU offloading, hybrid mode, and `‑ngl` flag.

## Suggested Links
- `shingo-docs/refs/llm-quantization-gptq-awq-gguf.md` (referenced for quantization contrast)
- `wiki/sources/qwen-llama-cpp.md` (referenced Qwen local‑inference guide)

## Review Checklist
- [ ] Verify source accuracy against the original llama.cpp documentation.
- [ ] Ensure no duplication with existing wiki pages (especially `wiki/sources/qwen-llama-cpp.md`).
- [ ] Confirm canonical project name (`llama.cpp`) and proper slugification for all candidate pages.
- [ ] Check that entity pages capture only widely known, verifiable facts.
