---
title: Quantization
kind: concept
sources:
  - source.md
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

# Quantization

Quantization is the process of reducing the numerical precision of a model's parameters—typically from 16‑bit floating point down to 8‑bit, 4‑bit, or even lower integer ranges. For large language models, this compression dramatically shrinks the memory footprint, making it possible to load and run models on consumer hardware that would otherwise require far more VRAM or RAM.

In the [Llama Cpp](../entities/llama-cpp.md) ecosystem, quantization is implemented through the [Gguf](gguf.md) single‑file format. A GGUF file bundles the model weights together with the tokenizer and chat template, all stored at a chosen precision. Widely used quantized variants such as Q4_K_M and Q8_0 illustrate the range: the former uses a more aggressive 4‑bit scheme to save space, while the latter retains higher fidelity at 8 bits.

The fundamental trade‑off is between resource efficiency and output quality. Lower‑bit quantization compresses the model more, reducing memory use and often speeding up inference on CPU‑only machines, but it can introduce subtle degradation in the model’s reasoning and generation quality. Hybrid inference, which offloads some layers to a GPU through backends supported by [Llama Cpp Backends](llama-cpp-backends.md), relies on quantized CPU‑resident layers to keep the total memory budget within the capacity of the available hardware.

The early work on llama.cpp by [Georgi Gerganov](../entities/georgi-gerganov.md) established a practical baseline for portable quantized inference. By coupling a lightweight C/C++ codebase with the single‑file GGUF container, the project made it routine to distribute models in quantized form. This approach underpins the local, dependency‑light execution that defines the llama.cpp toolchain and the downstream tools that build on it.
