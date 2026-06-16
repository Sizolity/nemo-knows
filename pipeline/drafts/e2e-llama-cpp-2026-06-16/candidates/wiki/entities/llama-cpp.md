---
title: llama.cpp
kind: entity
sources:
  - source.md
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

# llama.cpp

llama.cpp is an open‑source C/C++ library and command‑line toolkit for running large language model inference on commodity hardware. It was created by [Georgi Gerganov](georgi-gerganov.md) and first released in March 2023, initially to run Meta’s LLaMA models on a laptop CPU. The project, now maintained by the `ggml-org` community, has grown to support many model families (such as Qwen, Mistral, and Gemma) and prioritises local, dependency‑light execution: a single binary compiles and runs quantized models without a GPU server or Python runtime.

The library is deliberately portable, offering CPU‑first inference with optional GPU acceleration. It provides accelerated backends for Apple Silicon (Metal), NVIDIA GPUs (CUDA), AMD GPUs, Intel GPUs, and Vulkan. Hybrid CPU/GPU inference is controlled by the `-ngl` flag, which offloads a configurable number of layers to the GPU, allowing models that exceed VRAM to run on a system. These hardware backends are detailed in [Llama Cpp Backends](../concepts/llama-cpp-backends.md).

Models are stored in the [Gguf](../concepts/gguf.md) single‑file format, which bundles weights, tokenizer, and chat template into one file. GGUF is the successor to the earlier GGML format and is the only practical container for pure‑CPU inference. Quantized variants, such as Q4_K_M and Q8_0, trade memory footprint for quality, and are discussed in the [Quantization](../concepts/quantization.md) reference.

The project ships several tools, including `llama-cli` for interactive and batch use and `llama-server`, which exposes an OpenAI‑compatible HTTP endpoint. Both accept standard sampling controls—context length, generation length, temperature, top‑k, and top‑p—and are designed to make local inference straightforward. Because of its lightweight design, llama.cpp underpins popular downstream tools like Ollama and LM Studio.
