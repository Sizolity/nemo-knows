---
title: Llama Cli Parameters
kind: concept
sources:
  - source.md
  - raw/web/qwen-llama-cpp.md
confidence: medium
---

# Llama Cli Parameters

The `llama-cli` utility is an example program within the `llama.cpp` ecosystem used to run local inference on Qwen models. It supports various hardware backends including CPUs, Apple Silicon, NVIDIA CUDA, AMD GPU, Intel GPU, and Vulkan.

## Usage Example

A typical command for running local inference includes flags for model selection, chat template usage, color output, GPU offloading, and generation parameters:

```bash
./llama-cli -hf Qwen/Qwen3-8B-GGUF:Q8_0 --jinja --color -ngl 99 -fa -sm row --temp 0.6 --top-k 20 --top-p 0.95 --min-p 0 -c 40960 -n 32768 --no-context-shift
```

## Parameter Flags

- `-hf`: Specifies the Hugging Face URL for the model file (e.g., `Qwen/Qwen3-8B-GGUF:Q8_0`).
- `--jinja`: Utilizes the embedded Jinja chat template.
- `--color`: Enables colored output in the terminal.
- `-ngl`: Controls the number of GPU layers to offload (e.g., `99` for full offload).
- `-fa`: Applies flash attention optimization.
- `-sm row`: Sets the memory space management strategy (e.g., row-major).
- `--temp`: Sets the sampling temperature (e.g., `0.6`).
- `--top-k`: Limits the number of tokens considered for sampling (e.g., `20`).
- `--top-p`: Applies nucleus sampling probability threshold (e.g., `0.95`).
- `--min-p`: Sets the minimum presence penalty value (e.g., `0`).
- `-c`: Defines the context window size in tokens (e.g., `40960`).
- `-n`: Specifies the number of tokens to generate (e.g., `32768`).
- `--no-context-shift`: Disables context shifting mechanisms.

## Troubleshooting

If generation repeats or runs endlessly, increasing the `--presence-penalty` parameter up to `2.0` may help mitigate repetition issues.

For models requiring a specific thinking mode that is not exposed by default in the chat template, users can pass a custom non-thinking Jinja chat template using the `--chat-template-file` flag.
