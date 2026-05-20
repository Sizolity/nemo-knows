# llama.cpp Parameter Reference

This note preserves the llama.cpp parameter knowledge used by `nemo-knows`.
It is based on already-gathered project research and local `llama-cli --help`
output. It should be updated only after a deliberate version check.

Alignment checked against Qwen's llama.cpp guide and the llama.cpp
`tools/cli/README.md` on 2026-05-17.

## Local Binary

```text
/home/karo/src/llama.cpp/build/bin/llama-cli
build: b8895-b76429a69
model: /home/karo/models/qwen3.5-9b-q4_k_m.gguf
```

The local build supports the parameters needed for Qwen reasoning control:

```text
--jinja
--reasoning [on|off|auto]
--reasoning-budget N
--reasoning-budget-message MESSAGE
--chat-template-kwargs STRING
--presence-penalty N
--repeat-penalty N
--ctx-size N
--no-context-shift
```

## Core Runtime Parameters

| Parameter | Purpose | nemo use |
| --- | --- | --- |
| `-m`, `--model` | Local GGUF model path. | Always set from `Config.LlamaModel`. |
| `-hf`, `--hf-repo` | Load a model from Hugging Face, optionally with a quant suffix. | Not used by `nemo`; local paths are more reproducible. |
| `-mu`, `--model-url` | Load a model from a direct URL. | Not used by `nemo`. |
| `-p`, `--prompt` | Prompt text. | Not used by `nemo`; large rendered prompts can exceed OS argument length limits. |
| `-f`, `--file` | Read prompt from a file. | Used for every rendered prompt so long raw sources do not become command-line arguments. |
| `-n`, `--predict` | Maximum generated tokens. | Profile-specific token budget. |
| `-c`, `--ctx-size` | Prompt context size. `0` means model default. | `stable` and `fallback` use 24576 after Batch 1 corpus testing showed the model default was too small for medium-length SQLite docs. |
| `-ngl`, `--gpu-layers` | Number of layers offloaded to GPU, or `all`. | Defaults to `all`. |
| `-t`, `--threads` | CPU threads for generation. | Not currently set; GPU offload is the main path. |
| `-tb`, `--threads-batch` | CPU threads for prompt processing. | Not currently set. |
| `-fa`, `--flash-attn` | Flash Attention mode. | Official Qwen examples use it; not currently set by `nemo`. |
| `-sm`, `--split-mode` | Multi-GPU split mode: `none`, `layer`, `row`, or `tensor`. | Not currently set; single-GPU local default is sufficient. |
| `-dev`, `--device` | Select offload devices. | Not currently set. |
| `--single-turn` | One generation turn. | Always enabled for draft generation. |
| `--simple-io` | Basic subprocess-friendly IO. | Always enabled. |
| `--no-display-prompt` | Do not echo the prompt as display output. | Always enabled, though raw logs may still contain llama.cpp shell framing. |
| `--no-context-shift` | Stop instead of rotating context when context fills. | Enabled by default in `nemo` profiles. |
| `-sys`, `--system-prompt` | System prompt for chat-template-aware models. | Not currently used; prompts are rendered as one user prompt. |
| `-co`, `--color` | Colorize CLI output. | Avoided for subprocess raw output. |

## Sampling Parameters

| Parameter | Purpose | Notes |
| --- | --- | --- |
| `--temp` | Sampling temperature. | Lower values are more deterministic. |
| `--top-p` | Nucleus sampling cutoff. | Qwen non-thinking defaults use `0.8`; thinking uses `0.95`. |
| `--top-k` | Keep only top K token candidates. | Qwen guidance commonly uses `20`. |
| `--min-p` | Minimum probability filter. | Qwen guidance uses `0`. |
| `--presence-penalty` | Penalizes tokens that already appeared. | Qwen notes `0` to `2` can reduce repetition; `1.5` is useful for non-thinking maintenance. |
| `--repeat-penalty` | Penalizes repeated sequences. | Qwen recommendations keep this at `1.0`. |
| `--repeat-last-n` | Window for repeat penalty. | Not currently set by `nemo`. |

## Reasoning Controls

| Parameter | Purpose | Recommended use |
| --- | --- | --- |
| `--reasoning off` | High-level llama.cpp switch to disable reasoning. | Default for `fast`, `stable`, and `fallback`. |
| `--reasoning on` | High-level switch to enable reasoning. | Used only by `deep`. |
| `--reasoning auto` | Let llama.cpp infer from template. | Avoid for maintenance paths because Qwen3.5 can think unexpectedly. |
| `--reasoning-budget 0` | Immediate reasoning cutoff. | Used with `--reasoning off` for non-thinking profiles. |
| `--reasoning-budget N` | Bound reasoning to N tokens. | `deep` uses a bounded budget instead of unlimited reasoning. |
| `--reasoning-budget-message` | Message inserted when reasoning budget is exhausted. | Should tell the model to write final Markdown. |
| `--chat-template-kwargs '{"enable_thinking":false}'` | Qwen template-level thinking control. | Used with non-thinking profiles. |
| `--reasoning-format` | Controls whether thought tags are parsed/extracted, e.g. `deepseek`. | Useful for `llama-server`; not needed by `nemo` CLI raw capture. |
| `--chat-template-file` | Custom Jinja chat template. | Fallback if a future build/template ignores the regular controls. |

`--reasoning off`, `--reasoning-budget 0`, and
`--chat-template-kwargs '{"enable_thinking":false}'` are intentionally used
together for maintenance profiles. llama.cpp/Qwen3.5 behavior has varied across
builds and templates, so layered controls are safer than relying on one flag.

## nemo Profile Mapping

| Profile | Max tokens | Context | Temp | Top-p | Top-k | Min-p | Presence | Repeat | Reasoning |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| `fast` | 2048 | model default | 0.2 | 0.9 | 20 | 0 | 0 | 1.0 | off, budget 0 |
| `stable` | 32768 | 24576 | 0.7 | 0.8 | 20 | 0 | 1.5 | 1.0 | off, budget 0 |
| `deep` | 65536 | model default | 0.6 | 0.95 | 20 | 0 | 0 | 1.0 | on, budget 2000 |
| `fallback` | 16384 | 24576 | 0.2 | 0.8 | 20 | 0 | 1.5 | 1.0 | off, budget 0 |

## Official Qwen llama.cpp Example

Qwen's llama.cpp guide uses this shape for local Qwen3 GGUF inference:

```sh
./llama-cli \
  -hf Qwen/Qwen3-8B-GGUF:Q8_0 \
  --jinja \
  --color \
  -ngl 99 \
  -fa \
  -sm row \
  --temp 0.6 \
  --top-k 20 \
  --top-p 0.95 \
  --min-p 0 \
  -c 40960 \
  -n 32768 \
  --no-context-shift
```

For `nemo`, the equivalent is adapted for subprocess use and the local model:

- use `-m /home/karo/models/qwen3.5-9b-q4_k_m.gguf` rather than `-hf`;
- omit `--color` so raw output stays parseable;
- keep `--single-turn`, `--simple-io`, and `--no-display-prompt`;
- default maintenance profiles use non-thinking reasoning controls;
- `-fa` and `-sm row` remain candidate future optimizations, not current
  defaults.

## Context And Long Context

Qwen's llama.cpp guide describes context behavior this way:

- `-c` controls maximum context length.
- `-n` controls maximum generation length.
- if context shifting is enabled and context fills, llama.cpp can keep initial
  prompt tokens and discard part of the rest, then continue generation.
- `--no-context-shift` prevents this rotating behavior and stops once context
  capacity is reached.
- YaRN can extend context, for example:

```text
-c 131072 --rope-scaling yarn --rope-scale 4 --yarn-orig-ctx 32768
```

`nemo` does not enable YaRN automatically. It should be explicit because long
context settings can change quality and performance on short inputs.

## llama-server Notes

Qwen's guide also documents `llama-server`:

```sh
./llama-server \
  -hf Qwen/Qwen3-8B-GGUF:Q8_0 \
  --jinja \
  --reasoning-format deepseek \
  -ngl 99 \
  -fa \
  -sm row \
  --temp 0.6 \
  --top-k 20 \
  --top-p 0.95 \
  --min-p 0 \
  -c 40960 \
  -n 32768 \
  --no-context-shift
```

`llama-server` can expose an OpenAI-compatible API and parse reasoning content.
`nemo` currently uses `llama-cli` because draft generation is local,
subprocess-based, and stores raw outputs next to cleaned drafts.

## Local Verification

This command verified that the local build can suppress visible thinking for a
short prompt:

```sh
llama-cli \
  -m /home/karo/models/qwen3.5-9b-q4_k_m.gguf \
  -p "Return exactly: OK" \
  -n 256 \
  -ngl all \
  --single-turn \
  --simple-io \
  --no-display-prompt \
  --temp 0.7 \
  --top-p 0.8 \
  --top-k 20 \
  --min-p 0 \
  --reasoning off \
  --reasoning-budget 0 \
  --chat-template-kwargs '{"enable_thinking":false}'
```

The output was `OK` without `[Start thinking]`.
