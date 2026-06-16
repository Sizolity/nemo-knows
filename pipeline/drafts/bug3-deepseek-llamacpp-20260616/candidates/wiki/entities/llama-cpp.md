---
title: llama.cpp
kind: entity
sources:
  - source.md
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

# llama.cpp

llama.cpp is an open-source C/C++ library and command-line toolkit for running large language model inference on commodity hardware without a GPU server or Python runtime. Created by [[georgi-gerganov]] and first released in March 2023, it was originally designed to run Meta’s LLaMA models on a laptop CPU. The project is maintained by the `ggml-org` community and underlies numerous downstream tools, including Ollama and LM Studio.

The library is portable across plain CPUs and adds optional acceleration through backends for Metal, CUDA, AMD GPUs, Intel GPUs, and Vulkan. It can split computation between CPU and GPU by offloading specific layers, controlled by the `-ngl` flag, allowing a model to run even when GPU memory alone is insufficient. ([[llama-cpp-hardware-backends]])

Models are loaded from single-file containers in the [[gguf]] format, which stores weights together with metadata such as the tokenizer and chat template. GGUF is the successor to GGML and the standard format for CPU‑first quantization; commonly distributed variants include Q4_K_M and Q8_0, where lower-bit quantizations reduce memory use at a modest quality trade-off.

The project provides `llama-cli` for interactive and batch mode and `llama-server` as an OpenAI-compatible HTTP endpoint. Both tools accept parameters for context length, generation length, and sampling (temperature, top‑k, top‑p). The toolkit supports many model families—including LLaMA, Qwen, Mistral, and Gemma—and is often the recommended inference path for running local GGUF quantizations.
