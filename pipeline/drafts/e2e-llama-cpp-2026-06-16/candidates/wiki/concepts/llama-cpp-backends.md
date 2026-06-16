---
title: Llama Cpp Backends
kind: concept
sources:
  - source.md
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

# Llama Cpp Backends

The [Llama Cpp](../entities/llama-cpp.md) project, initiated by [Georgi Gerganov](../entities/georgi-gerganov.md), is built around a portable inference engine that runs on a deliberately broad set of hardware targets. Its backends span plain CPU execution and a collection of GPU-accelerated paths, enabling the same codebase to serve anything from a laptop to a workstation with a dedicated graphics card.

On the GPU side, the library offers accelerated backends for Apple Silicon through Metal, for NVIDIA cards via CUDA, and for AMD, Intel, and Vulkan-capable devices. Each backend is optional at build time, so users compile only the acceleration their machine actually needs. The plain CPU path remains the fallback and still powers the project's signature dependency-light experience.

A defining feature of the backend architecture is hybrid CPU/GPU inference. The `-ngl` parameter controls how many model layers are offloaded to the GPU, while the remaining layers stay on the CPU. This division allows a model whose total size exceeds available VRAM to run successfully, since only a portion of the computation must fit on the graphics card.

These backends operate on models stored in the [Gguf](gguf.md) single-file format, which bundles weights, tokenizer metadata, and chat templates together. GGUF is the practical format for CPU inference and pairs naturally with [Quantization](quantization.md) variants like Q4_K_M or Q8_0, giving each backend an efficient memory-versus-quality trade-off without changing the model file itself.
