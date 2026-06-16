# llama.cpp

Source: distilled from the public llama.cpp project documentation
(https://github.com/ggml-org/llama.cpp) and cross-checked against the LLM
quantization reference notes in `shingo-docs/refs/llm-quantization-gptq-awq-gguf.md`
and the Qwen llama.cpp local-inference guide already summarized in
`wiki/sources/qwen-llama-cpp.md`. This file records widely known, verifiable
facts about llama.cpp for ingest testing of entity pages. The canonical spelling
of the project name is the lowercase, dotted form `llama.cpp`.

## What llama.cpp is

llama.cpp is an open-source library and command-line ecosystem, written in C and
C++, for running large language model inference efficiently on commodity
hardware. It was started by Georgi Gerganov and first released in March 2023,
originally to run Meta's LLaMA models on a laptop CPU. The project has since
grown to support many model families, including Qwen, Mistral, and Gemma, and is
maintained by a large community of contributors under the `ggml-org` organization.

The defining goal of llama.cpp is local, dependency-light inference: a user can
compile a single binary and run a quantized model on a personal machine without a
GPU server or a Python runtime. Because of this, llama.cpp underpins many
downstream tools, including Ollama and LM Studio, that wrap it behind friendlier
interfaces.

## Hardware backends

llama.cpp is deliberately portable. It runs on plain CPUs and adds optional
accelerated backends for Apple Silicon (via Metal), NVIDIA GPUs (via CUDA), AMD
GPUs, Intel GPUs, and Vulkan. It also supports hybrid CPU/GPU inference, in which
some model layers are offloaded to the GPU while the rest stay on the CPU. The
`-ngl` (number of GPU layers) flag controls how many layers are offloaded, which
lets a large model run on a machine whose GPU memory alone could not hold it.

## The GGUF model format

llama.cpp reads models in the GGUF file format, a single-file container that
stores model weights together with metadata such as the tokenizer and chat
template. GGUF is the successor to the earlier GGML format. Compared with the
multi-file layouts used by GPU-first quantization schemes such as GPTQ and AWQ,
GGUF's single-file design is convenient to download and share, and it is the only
practical format for pure-CPU inference. Models are commonly distributed in
quantized GGUF variants such as Q4_K_M and Q8_0, where lower-bit quantization
trades some quality for a smaller memory footprint.

## Tooling

The project ships several example programs. `llama-cli` is the interactive and
batch command-line runner, and `llama-server` exposes an OpenAI-compatible HTTP
endpoint for serving completions. These tools accept generation controls such as
context length (`-c`), generation length (`-n`), and sampling parameters like
temperature, top-k, and top-p. The same flags appear in the Qwen project's
recommended `llama-cli` invocation for running Qwen3 GGUF models locally.
