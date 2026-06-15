---
kind: source
sources:
  - web:https://qwen.readthedocs.io/en/v3.0/run_locally/llama.cpp.html
confidence: medium
---

# Qwen llama.cpp Guide

## What It Is

Qwen's official guide for running Qwen models locally with llama.cpp, including
`llama-cli`, `llama-server`, GGUF model files, GPU offload, sampling parameters,
chat templates, context management, and thinking-mode caveats.

## Summary

The guide explains that llama.cpp is a lightweight C/C++ inference ecosystem and
that users can run Qwen GGUF models with `llama-cli` or `llama-server`. It
recommends using the embedded Jinja chat template, GPU layer offload, optional
Flash Attention, explicit sampling parameters, context sizing, and
`--no-context-shift`. It also notes that a custom non-thinking chat template may
be needed when the hard thinking switch is not exposed cleanly.

## Key Claims

- Qwen3 and Qwen3MoE are supported by llama.cpp from version `b5092`.
- Official Qwen GGUF models can be downloaded from Hugging Face or converted
  from Hugging Face model files.
- `--jinja` uses the GGUF embedded chat template and is preferred for Qwen chat
  models.
- `-c` controls context length, `-n` controls generation length, and
  `--no-context-shift` prevents rotating context continuation.
- Presence penalty can help when generation repeats or runs endlessly.

## Suggested Links

- llama.cpp
- GGUF
- Qwen
- Jinja chat template
- context management
