---
title: Gguf
kind: concept
sources:
  - source.md
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

# Gguf

GGUF is the single-file model container that serves as the primary input format for [[llama-cpp]]. It succeeds the earlier GGML layout, bundling quantized weights together with essential metadata such as the tokenizer and chat template in one self-contained file.

The container’s unified design makes models straightforward to distribute and share. Unlike the multi-file layouts common in GPU-first quantization schemes like GPTQ and AWQ, GGUF avoids scattered data files, which simplifies loading on local machines and lowers the barrier to running models without specialised tooling.

Because the format is built for CPU-first inference, it underpins quantised variants such as Q4_K_M and Q8_0. Lower-bit quantisations reduce memory usage at the expense of a slight decline in output quality, enabling large language models to fit within the memory budgets of consumer hardware.

The format emerged from the [[llama-cpp]] project, originally started by [[georgi-gerganov]]. As the backbone of llama.cpp, GGUF enables the library’s [[llama-cpp-hardware-backends]]—from pure CPU execution to mixed CPU/GPU offloading—to operate predictably across devices. Numerous downstream tools, including Ollama and LM Studio, rely on GGUF for friendly local model serving.
