---
title: Qwen llama.cpp Local Inference Notes
kind: source
sources:
  - raw/web/qwen-llama-cpp.md
confidence: medium
---

# Qwen llama.cpp Local Inference Notes

## What It Is
A guide describing how to run Qwen models locally using the `llama-cli` and `llama-server` example programs within the lightweight C/C++ inference ecosystem provided by llama.cpp.

## Summary
The document outlines support for various hardware backends including CPUs, Apple Silicon, NVIDIA CUDA, AMD GPU, Intel GPU, Vulkan, and hybrid CPU/GPU configurations. It details how to obtain GGUF model files from the Qwen Hugging Face organization or convert existing models, specifically mentioning `Qwen/Qwen3-8B-GGUF` and the `qwen3-8b-q4_k_m.gguf` quantized model. The guide provides a specific command example for `llama-cli` to run inference with various flags for GPU offloading, context management, and generation parameters. It also addresses a caveat regarding thinking-mode switches, recommending the use of a custom non-thinking Jinja chat template passed via `--chat-template-file` if necessary.

## Key Claims
- llama.cpp supports Qwen3 and Qwen3MoE starting from version `b5092`.
- The command `./llama-cli -hf Qwen/Qwen3-8B-GGUF:Q8_0 --jinja --color -ngl 99 -fa -sm row --temp 0.6 --top-k 20 --top-p 0.95 --min-p 0 -c 40960 -n 32768 --no-context-shift` is an example for local inference.
- The `--jinja` flag utilizes the embedded chat template, while `-ngl` controls GPU layer offload.
- If generation repeats or runs endlessly, setting `--presence-penalty` up to `2.0` may help.
- A hard switch between thinking and non-thinking modes implemented in the chat template may not be exposed in llama.cpp by default; a workaround involves passing a custom file with `--chat-template-file`.

## Suggested Links
https://qwen.readthedocs.io/en/v3.0/run_locally/llama.cpp.html
