---
title: llama.cpp
kind: source
sources:
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

## What It Is

llama.cpp is an open-source C/C++ library and command‑line toolset for running
large language model inference efficiently on commodity hardware. Started by
Georgi Gerganov in March 2023 to run Meta LLaMA models on laptop CPUs, it now
supports many model families (e.g., Qwen, Mistral, Gemma) and is maintained by
the `ggml-org` community. Its design prioritises local, dependency‑light
execution: a single binary compiles and runs quantized models without a GPU
server or Python runtime. The project underpins popular downstream tools such as
Ollama and LM Studio.

## Summary

llama.cpp provides portable CPU‑first inference with optional GPU acceleration
for Apple Silicon (Metal), NVIDIA (CUDA), AMD, Intel, and Vulkan. It can offload
a configurable number of layers to the GPU (`-ngl`) to run models larger than
VRAM. Models are stored in the GGUF single‑file format, which bundles weights,
tokenizer, and chat template—making it convenient for distribution and the only
practical route for pure‑CPU inference. Common quantized variants like Q4_K_M
and Q8_0 trade memory footprint for quality. The project ships with
`llama‑cli` for interactive/batch use and `llama‑server` for an
OpenAI‑compatible HTTP endpoint, both accepting context length, generation
length, temperature, top‑k, top‑p, and other sampling controls.

## Key Claims

- Designed for local, dependency‑light inference on consumer hardware; no GPU
  server or Python required.
- Written in C and C++, delivering broad portability and multiple backend
  accelerators (Metal, CUDA, AMD, Intel GPU, Vulkan).
- Hybrid CPU/GPU inference allows a model bigger than GPU memory to run by
  controlling offloaded layers.
- Uses the single‑file GGUF format (successor to GGML) that simplifies sharing
  and enables pure‑CPU inference; quantized variants range from low‑bit to
  higher‑fidelity.
- Exposes standard generation controls (context length, generation length,
  temperature, top‑k, top‑p) across CLI and HTTP server interfaces.
- Widely embedded in downstream tools like Ollama and LM Studio.

## Suggested Links

- Project repository: https://github.com/ggml-org/llama.cpp
