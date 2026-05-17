# Qwen3.5 Hugging Face Parameter Guidance

This note records the Qwen3.5 sampling and generation settings found on
Hugging Face model cards. It is about inference-time parameter tuning, not
training or fine-tuning hyperparameters.

Checked on 2026-05-16. Updated with llama.cpp/Qwen official guidance on
2026-05-17.

## Sources

- [Qwen/Qwen3.5-0.8B](https://huggingface.co/Qwen/Qwen3.5-0.8B)
- [Qwen/Qwen3.5-27B](https://huggingface.co/Qwen/Qwen3.5-27B)
- [Qwen/Qwen3.5-35B-A3B](https://huggingface.co/Qwen/Qwen3.5-35B-A3B)
- [Qwen/Qwen3.5-397B-A17B](https://huggingface.co/Qwen/Qwen3.5-397B-A17B)
- [Qwen3.5 collection](https://huggingface.co/collections/Qwen/qwen35)
- [Qwen/Qwen3.5-35B-A3B discussion #74](https://huggingface.co/Qwen/Qwen3.5-35B-A3B/discussions/74)
- [Qwen llama.cpp guide](https://qwen.readthedocs.io/en/v3.0/run_locally/llama.cpp.html)
- [llama.cpp CLI options](https://github.com/ggml-org/llama.cpp/blob/master/tools/cli/README.md)
- [llama.cpp PR #13196: chat template kwargs](https://github.com/ggml-org/llama.cpp/pull/13196)
- [llama.cpp PR #20297: reasoning budget](https://github.com/ggml-org/llama.cpp/pull/20297)
- [llama.cpp issue #20833: Qwen3.5 thinking by default](https://github.com/ggml-org/llama.cpp/issues/20833)

## General Pattern

Qwen3.5 model cards recommend choosing parameters by mode and task type:

- Thinking mode usually uses higher nucleus sampling: `top_p=0.95`,
  `top_k=20`, and `min_p=0.0`.
- Non-thinking or instruct mode usually uses lower temperature/nucleus
  sampling for general direct answers: `temperature=0.7`, `top_p=0.8`,
  `top_k=20`, and `min_p=0.0`.
- Precise coding or visual-language tasks in thinking mode use
  `temperature=0.6` and no presence penalty.
- `presence_penalty` is used to reduce endless repetitions when the serving
  framework supports it. Qwen notes that values between `0` and `2` can help,
  but higher values may cause language mixing or a small performance drop.
- Qwen3.5 does not officially support Qwen3's soft prompt switches
  `/think` and `/nothink`. Use API or chat-template parameters instead.

## Recommended Sampling Parameters

| Model group | Mode / task | Recommended parameters |
| --- | --- | --- |
| `Qwen3.5-397B-A17B` | Thinking mode | `temperature=0.6`, `top_p=0.95`, `top_k=20`, `min_p=0.0`, `presence_penalty=0.0`, `repetition_penalty=1.0` |
| `Qwen3.5-397B-A17B` | Instruct / non-thinking mode | `temperature=0.7`, `top_p=0.8`, `top_k=20`, `min_p=0.0`, `presence_penalty=1.5`, `repetition_penalty=1.0` |
| `Qwen3.5-27B`, `Qwen3.5-35B-A3B` | Thinking mode for general tasks | `temperature=1.0`, `top_p=0.95`, `top_k=20`, `min_p=0.0`, `presence_penalty=1.5`, `repetition_penalty=1.0` |
| `Qwen3.5-27B`, `Qwen3.5-35B-A3B` | Thinking mode for precise coding, e.g. WebDev | `temperature=0.6`, `top_p=0.95`, `top_k=20`, `min_p=0.0`, `presence_penalty=0.0`, `repetition_penalty=1.0` |
| `Qwen3.5-27B`, `Qwen3.5-35B-A3B` | Instruct / non-thinking mode for general tasks | `temperature=0.7`, `top_p=0.8`, `top_k=20`, `min_p=0.0`, `presence_penalty=1.5`, `repetition_penalty=1.0` |
| `Qwen3.5-27B`, `Qwen3.5-35B-A3B` | Instruct / non-thinking mode for reasoning tasks | See conflict note below. |
| `Qwen3.5-0.8B` | Non-thinking text tasks | `temperature=1.0`, `top_p=1.0`, `top_k=20`, `min_p=0.0`, `presence_penalty=2.0`, `repetition_penalty=1.0` |
| `Qwen3.5-0.8B` | Non-thinking visual-language tasks | `temperature=0.7`, `top_p=0.8`, `top_k=20`, `min_p=0.0`, `presence_penalty=1.5`, `repetition_penalty=1.0` |
| `Qwen3.5-0.8B` | Thinking text tasks | `temperature=1.0`, `top_p=0.95`, `top_k=20`, `min_p=0.0`, `presence_penalty=1.5`, `repetition_penalty=1.0` |
| `Qwen3.5-0.8B` | Thinking visual-language or precise coding tasks | `temperature=0.6`, `top_p=0.95`, `top_k=20`, `min_p=0.0`, `presence_penalty=0.0`, `repetition_penalty=1.0` |

## Thinking Mode Control

For `Qwen3.5-27B`, `Qwen3.5-35B-A3B`, and `Qwen3.5-397B-A17B`, the model card
states that Qwen3.5 thinks by default. To get a direct non-thinking response
through vLLM or SGLang-compatible OpenAI APIs, pass:

```json
{
  "chat_template_kwargs": {
    "enable_thinking": false
  }
}
```

Alibaba Cloud Model Studio uses a different shape:

```json
{
  "enable_thinking": false
}
```

The `Qwen3.5-0.8B` card presents the inverse example for thinking mode: pass
`"enable_thinking": true` in the API body.

### llama.cpp Control

Qwen's llama.cpp guide recommends using the GGUF embedded chat template with
`--jinja`. It also notes that, when the hard switch is not exposed cleanly in
llama.cpp, a practical workaround is a custom non-thinking chat template passed
with `--chat-template-file`.

Current llama.cpp builds expose several relevant controls:

```text
--reasoning off
--reasoning-budget 0
--chat-template-kwargs '{"enable_thinking":false}'
--reasoning-budget-message "reasoning budget exceeded, now write the final answer"
```

`--reasoning off` is the intended high-level switch. `--reasoning-budget 0` is
the stronger budget-based path that immediately closes or prevents the thinking
section when the template path is insufficient. `--chat-template-kwargs` passes
Qwen's `enable_thinking=false` value into templates that support it.

The llama.cpp issue tracker shows that Qwen3.5 behavior has varied across
builds and templates. In some `llama-cli` builds, `--reasoning off` or
`--chat-template-kwargs` alone did not prevent `[Start thinking]`; in other
builds, the same settings worked. Treat local verification against the exact
`llama-cli` binary and GGUF as required.

Local verification on 2026-05-17 with:

```text
/home/karo/src/llama.cpp/build/bin/llama-cli
/home/karo/models/qwen3.5-9b-q4_k_m.gguf
```

confirmed that this combination suppressed visible thinking for a short prompt:

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

## Output Length

The model cards recommend:

- `32768` output tokens for most queries.
- `81920` output tokens for highly complex math or programming competition
  benchmarks.

For this project, `81920` is usually too expensive and increases cleanup risk.
Use it only when the task explicitly needs long reasoning or benchmark parity.

## Conflict Note

The `Qwen3.5-27B` and `Qwen3.5-35B-A3B` model cards currently contain conflicting
recommendations for instruct / non-thinking reasoning tasks:

- The Chat Completions API section says:
  `temperature=1.0`, `top_p=0.95`, `top_k=20`, `min_p=0.0`,
  `presence_penalty=1.5`, `repetition_penalty=1.0`.
- The Best Practices section says:
  `temperature=1.0`, `top_p=1.0`, `top_k=40`, `min_p=0.0`,
  `presence_penalty=2.0`, `repetition_penalty=1.0`.

Hugging Face discussion #74 asks which value is correct, but there was no
visible maintainer answer in the fetched page. Until Qwen resolves this, prefer
the conservative Chat Completions API values for production-like use and treat
the Best Practices values as an experimental reasoning profile.

## Suggested Local Defaults

For `nemo-knows`, maintenance profiles now use non-thinking behavior for ingest
drafts:

```text
temperature=0.7
top_p=0.8
top_k=20
min_p=0.0
presence_penalty=1.5
repetition_penalty=1.0
reasoning=off
reasoning_budget=0
chat_template_kwargs={"enable_thinking":false}
max_tokens=32768
```

Use thinking mode only for complex synthesis or contradiction analysis. The
`deep` profile keeps reasoning enabled but bounds it locally:

```text
temperature=0.6
top_p=0.95
top_k=20
min_p=0.0
presence_penalty=0.0
repetition_penalty=1.0
reasoning=on
reasoning_budget=1000-2000 for bounded local runs
reasoning_budget_message="reasoning budget exceeded, now write the final Markdown"
max_tokens=65536
```

For `Qwen3.5-0.8B`, monitor streaming output if thinking mode is enabled. Its
model card warns that the recommended thinking parameters are more likely to
enter thinking loops than other Qwen3.5 models.

## nemo E2E Observation

During the 2026-05-16 `drafts/e2e-real-llm-wiki` candidate-generation run,
`tooling-stack.raw.txt` and `tooling-stack.fallback.raw.txt` showed the model
spending the output budget on visible `[Start thinking]` content. The prompt
asked for `Tooling Stack`, while the source material was titled `LLM Wiki`; the
model repeatedly analyzed that mismatch and did not reach complete final
Markdown before the configured output budget ended.

This indicates that the slow path was not a Go deadlock. It was normal
llama.cpp generation at roughly 61 tokens/sec, with the token budget consumed by
reasoning. Because `nemo` currently captures llama.cpp output with
`CombinedOutput()`, the user sees no incremental progress until the process
exits.

After wiring the llama.cpp reasoning controls into `nemo`, the same
candidate-generation path completed in about 26 seconds with
`NEMO_MAX_TOKENS=2048` and no `[Start thinking]` markers in the primary raw
outputs. The remaining candidate-eval issues were content-contract issues
(heading/title match, wikilink presence, and originality), not runaway
reasoning.
