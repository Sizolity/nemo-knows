---
kind: topic
sources: [raw/web/qwen-llama-cpp.md]
status: draft
---

# Ingest Plan

## Source Summary
- The guide details running Qwen models locally using `llama-cli` and `llama-server` via the lightweight C/C++ llama.cpp ecosystem.
- It covers hardware support for CPUs, Apple Silicon, NVIDIA/AMD/Intel GPUs, Vulkan, and hybrid inference modes.
- Specific usage examples include GGUF model downloads, quantization flags (`Q8_0`), and advanced command-line parameters like `--jinja`, `-ngl`, and context management.
- A caveat regarding the "thinking-mode" soft switch is noted, recommending a custom Jinja chat template as a workaround for hard switches in llama.cpp.

## Candidate Wiki Pages
- wiki/sources/qwen-llama-cpp-guide.md — To document the specific source URL and version (v3.0) containing local inference notes.
- wiki/concepts/llama-cli-parameters.md — To catalog command-line flags such as `--jinja`, `-ngl`, `--no-context-shift`, and generation penalties.
- wiki/topics/qwen-local-inference-setup.md — To aggregate instructions for downloading GGUF models and configuring local inference environments.

## Suggested Links
- https://qwen.readthedocs.io/en/v3.0/run_locally/llama.cpp.html

## Review Checklist
- [ ] Verify GGUF model version compatibility (e.g., Qwen3 support from b5092).
- [ ] Confirm accuracy of `--chat-template-file` workaround for thinking mode.
- [ ] Validate GPU offload examples against current hardware constraints.
