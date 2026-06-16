---
title: Llama Cpp Hardware Backends
kind: concept
sources:
  - source.md
  - pipeline/raw/web/llama-cpp.md
confidence: medium
---

# Llama Cpp Hardware Backends

[[llama-cpp]] achieves broad portability through a layered backend architecture. It runs on plain CPUs using highly optimised scalar and SIMD code paths, then layers optional acceleration for specialised hardware on top. This design means the same binary can execute from the same GGUF model file without platform‑specific forks, keeping the local‑first promise of the library.

Apple Silicon devices receive Metal‑based GPU acceleration that is integrated into the project’s own compute graph. On macOS, the Metal backend uses the unified memory of M‑series chips to avoid expensive host‑to‑device transfers, providing a near‑transparent speed‑up for laptop and desktop inference. The backend is implemented without external library dependencies beyond what the operating system already provides.

For discrete GPUs, [[llama-cpp]] ships CUDA support for NVIDIA cards and Vulkan support that extends to AMD and Intel GPUs. These backends are compiled‑in at build time and can be selected dynamically at runtime. By using Vulkan, the library reaches a wide GPU ecosystem without requiring vendor‑specific libraries on every platform, reflecting the original portability goals of [[georgi-gerganov]].

A defining feature of the backend system is hybrid CPU‑GPU inference. The `-ngl` flag tells the runtime how many model layers to offload to the GPU, leaving the rest on the CPU. This allows users to run models whose total memory requirements exceed a GPU’s VRAM, combining the speed of a partial offload with the large capacity of system RAM. The same [[gguf]]‑packed model can be used with any number of layers offloaded, making the backend selection a runtime decision rather than a model‑conversion step.
