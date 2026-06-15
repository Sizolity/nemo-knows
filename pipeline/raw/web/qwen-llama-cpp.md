# Qwen llama.cpp Local Inference Notes

Source URL: https://qwen.readthedocs.io/en/v3.0/run_locally/llama.cpp.html

Qwen's llama.cpp guide describes how to run Qwen models locally with the
`llama-cli` and `llama-server` example programs. It explains that llama.cpp is a
lightweight C/C++ inference ecosystem with support for CPUs, Apple Silicon,
NVIDIA CUDA, AMD GPU backends, Intel GPU backends, Vulkan, and hybrid CPU/GPU
inference. The guide notes that llama.cpp supports Qwen3 and Qwen3MoE from
version `b5092`.

The guide recommends obtaining GGUF model files from the Qwen Hugging Face
organization or converting Hugging Face model files to GGUF with llama.cpp's
conversion script. It gives an example download command for
`Qwen/Qwen3-8B-GGUF` and the `qwen3-8b-q4_k_m.gguf` quantized model.

For `llama-cli`, the guide shows a command shaped like:

```sh
./llama-cli -hf Qwen/Qwen3-8B-GGUF:Q8_0 --jinja --color -ngl 99 -fa -sm row --temp 0.6 --top-k 20 --top-p 0.95 --min-p 0 -c 40960 -n 32768 --no-context-shift
```

It explains that `--jinja` uses the embedded chat template, `-ngl` controls GPU
layer offload, `-fa` may speed generation, `-sm row` controls split mode for
multi-GPU setups, `-c` controls context length, `-n` controls generation length,
and `--no-context-shift` prevents rotating context behavior when the context is
full. It also notes that if generation repeats or runs endlessly, a
`--presence-penalty` up to `2.0` may help.

The guide includes a thinking-mode caveat. It says the soft switch between
thinking and non-thinking is available, but the hard switch implemented in the
chat template may not be exposed in llama.cpp. As a workaround, it recommends a
custom non-thinking Jinja chat template passed with `--chat-template-file` when
needed.
