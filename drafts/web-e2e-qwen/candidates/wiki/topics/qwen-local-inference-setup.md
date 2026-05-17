---
title: Qwen Local Inference Setup
kind: topic
sources:
  - source.md
  - raw/web/qwen-llama-cpp.md
confidence: medium
---

# Qwen Local Inference Setup

This guide covers running Qwen models locally using the `llama-cli` and `llama-server` example programs within the lightweight C/C++ inference ecosystem provided by llama.cpp. It outlines support for various hardware backends, including CPUs, Apple Silicon, NVIDIA CUDA, AMD GPU, Intel GPU, Vulkan, and hybrid CPU/GPU configurations.

## Getting Started

To begin local inference, you can pull GGUF model files directly from the Qwen Hugging Face organization or convert existing models. A commonly used example is `Qwen/Qwen3-8B-GGUF`, specifically the `qwen3-8b-q4_k_m.gguf` quantized model.

## Running Inference

The following command demonstrates a basic local inference setup using `llama-cli`. This example utilizes GPU offloading, manages context size, and sets generation parameters:

```bash
./llama-cli -hf Qwen/Qwen3-8B-GGUF:Q8_0 --jinja --color -ngl 99 -fa -sm row --temp 0.6 --top-k 20 --top-p 0.95 --min-p 0 -c 40960 -n 32768 --no-context-shift
```

### Command Flags

- `--jinja`: Utilizes the embedded chat template for interaction.
- `-ngl`: Controls GPU layer offloading (e.g., `-ngl 99` attempts to offload all layers).
- `-fa`: Enables flash attention.
- `-sm row`: Sets the memory layout strategy.
- `--chat-template-file`: Allows passing a custom file if default templates do not support specific mode switches.

## Advanced Configuration and Troubleshooting

### Hardware Support
llama.cpp supports Qwen3 and Qwen3MoE starting from version `b5092`. You can leverage this to run models on diverse hardware, ensuring compatibility with your specific GPU or CPU setup.

### Handling Generation Issues
If generation repeats or runs endlessly, consider adjusting parameters such as setting `--presence-penalty` up to `2.0` to help mitigate infinite loops.

### Chat Template Considerations
A caveat exists regarding hard switches between thinking and non-thinking modes implemented in the chat template. These may not be exposed in llama.cpp by default. If necessary, use a custom non-thinking Jinja chat template passed via `--chat-template-file`.
