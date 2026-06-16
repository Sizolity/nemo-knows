---
title: llama.cpp: Local LLM Inference Library
kind: source
sources:
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

## What It Is
llama.cpp is an open-source library and command-line toolkit, written in C/C++, for running large language model inference on personal hardware. Created by Georgi Gerganov and first released in March 2023, it was originally designed to run Meta’s LLaMA models on a laptop CPU. The project is maintained by the `ggml-org` community and supports many model families, including Qwen, Mistral, and Gemma. Its core goal is dependency‑light, local inference—a user can compile a single binary and run a quantized model without needing a GPU server or a Python runtime. It underpins popular downstream tools such as Ollama and LM Studio.

## Summary
llama.cpp runs on plain CPUs with optional acceleration via Metal (Apple Silicon), CUDA (NVIDIA GPUs), AMD GPUs, Intel GPUs, and Vulkan. It supports hybrid CPU/GPU inference by offloading specific layers, controlled by the `-ngl` flag. Models are loaded from single‑file GGUF containers that store weights and metadata (tokenizer, chat template). GGUF is the successor to GGML and the standard format for CPU‑first quantization; common quantized variants include Q4_K_M and Q8_0. The project ships with `llama-cli` for interactive and batch use, and `llama-server`, an HTTP endpoint compatible with the OpenAI API. Both accept parameters for context length, generation length, and sampling (temperature, top‑k, top‑p).

## Key Claims
- Enables LLM inference on commodity CPUs without a GPU server or Python environment.
- Uses the GGUF single‑file format for convenient distribution and CPU‑friendly inference.
- Provides portable backends (CPU, Metal, CUDA, Vulkan) and hybrid layer offloading.
- Serves as the foundation for tools like Ollama and LM Studio.
- Supports many model families, including LLaMA, Qwen, Mistral, and Gemma.
- A single compiled binary (`llama-cli` or `llama-server`) is all that is needed to run models locally.

## Suggested Links
- [llama.cpp GitHub repository](https://github.com/ggml-org/llama.cpp)
- Ollama (downstream tool)
- LM Studio (downstream tool)
- Qwen project’s llama.cpp local‑inference guide (referenced in this source)
