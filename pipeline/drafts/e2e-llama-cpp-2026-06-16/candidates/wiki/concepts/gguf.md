---
title: GGUF
kind: concept
sources:
  - source.md
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

# GGUF

GGUF is the single-file container format that succeeded the earlier GGML layout and serves as the primary model distribution format for [Llama Cpp](../entities/llama-cpp.md). It packages model weights alongside essential metadata—including the tokenizer and chat template—so that a complete inference-ready model travels as one portable file. This design eliminates the need to assemble multiple shards or configuration files at load time.

By bundling metadata directly with the weights, GGUF files describe how a model should behave without relying on external configuration scripts. The integrated tokenizer and chat template allow the inference engine to format prompts and decode output correctly regardless of the model family, whether the underlying architecture is LLaMA, Mistral, Qwen, or another supported design. This self-contained approach makes the format particularly convenient to download, archive, and share across different machines and users.

The format is closely tied to [Quantization](quantization.md) practice. Distributors commonly publish GGUF files in several quantized variants—such as Q4_K_M and Q8_0—where each variant represents a different trade-off between memory footprint and output fidelity. Lower-bit quantizations fit into tighter RAM budgets and run faster on CPU, while higher-bit variants preserve more of the original model quality. Because the GGUF file stores the quantized weights directly, the end user does not need to run a separate quantization step before inference.

GGUF's single-file nature makes it the only practical container for pure-CPU inference at scale. Multi-file GPU-first schemes like GPTQ and AWQ assume the availability of accelerator memory and often require Python dependency chains for loading, whereas a GGUF file can be loaded by a single C++ binary without a GPU server or Python runtime. This alignment with [Llama Cpp](../entities/llama-cpp.md)'s dependency-light philosophy—pioneered by [Georgi Gerganov](../entities/georgi-gerganov.md)—is what turned the format into the default interchange for local, hybrid CPU/GPU inference across backends ranging from Metal and CUDA to Vulkan and plain x86.
