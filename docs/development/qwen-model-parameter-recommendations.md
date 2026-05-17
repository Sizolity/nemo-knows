# Qwen Model Parameter Recommendations

This note preserves the Qwen model parameter guidance used by `nemo-knows`.
It is copied from already-gathered project notes and should not be treated as a
live upstream reference.

## Source Scope

The recommendations were gathered from Qwen model cards and Qwen llama.cpp
guidance during earlier project work, then verified against the local
Qwen3.5-9B GGUF through llama.cpp.

Alignment checked against Qwen's current inference and llama.cpp docs on
2026-05-17.

Relevant local model:

```text
/home/karo/models/qwen3.5-9b-q4_k_m.gguf
```

## General Pattern

Qwen models should be tuned by mode:

- Non-thinking or instruct-style maintenance tasks use direct-answer sampling:
  `temperature=0.7`, `top_p=0.8`, `top_k=20`, `min_p=0`.
- Thinking tasks use wider nucleus sampling:
  `temperature=0.6`, `top_p=0.95`, `top_k=20`, `min_p=0` for precise work.
- Greedy decoding is not recommended for thinking mode because it can degrade
  quality and increase repetition.
- Presence penalty can reduce endless repetition when the runtime supports it.
  Qwen notes values between `0` and `2`; `nemo` uses `1.5` for non-thinking
  maintenance profiles.
- Repetition penalty should remain `1.0` unless a concrete repetition problem is
  reproduced and tested.

## Thinking Control

Qwen's current Qwen3 inference docs state that Qwen3 thinks before responding by
default. They document two prompt-level controls:

- append a final assistant message containing only blank thinking delimiters to
  strictly prevent thinking for one turn;
- add `/no_think` or `/think` to the user or system message to switch mode
  across turns.

For API runtimes that support Qwen chat template kwargs, non-thinking mode can
also be requested with:

```json
{
  "chat_template_kwargs": {
    "enable_thinking": false
  }
}
```

llama.cpp exposes this through:

```text
--chat-template-kwargs '{"enable_thinking":false}'
```

Qwen's llama.cpp guide adds an important caveat: the soft switch is always
available, but the hard switch implemented in the chat template may not be
exposed cleanly in llama.cpp. Its documented workaround is a custom
non-thinking Jinja template passed with `--chat-template-file`.

For this project, do not rely on `/think` or `/no_think` as the primary control
for Qwen3.5 through llama.cpp. Prefer the runtime/template controls that were
verified locally:

```text
--reasoning off
--reasoning-budget 0
--chat-template-kwargs '{"enable_thinking":false}'
```

## nemo Profiles

### fast

Use for smoke tests and strict short-format tasks:

```text
temperature=0.2
top_p=0.9
top_k=20
min_p=0
presence_penalty=0
repeat_penalty=1.0
reasoning=off
reasoning_budget=0
max_tokens=2048
```

### stable

Use for default local wiki maintenance:

```text
temperature=0.7
top_p=0.8
top_k=20
min_p=0
presence_penalty=1.5
repeat_penalty=1.0
reasoning=off
reasoning_budget=0
chat_template_kwargs={"enable_thinking":false}
max_tokens=32768
```

This is the main profile for source summaries, ingest drafts, and candidate
pages. It should avoid visible thinking and produce reviewable Markdown.

### deep

Use only for complex synthesis, contradiction analysis, or tasks where reasoning
is explicitly useful:

```text
temperature=0.6
top_p=0.95
top_k=20
min_p=0
presence_penalty=0
repeat_penalty=1.0
reasoning=on
reasoning_budget=2000
reasoning_budget_message="reasoning budget exceeded, now write the final Markdown"
max_tokens=65536
```

The budget keeps local reasoning bounded so the model does not spend the full
output allowance on hidden or visible thought.

### fallback

Use after malformed output or cleaning failure:

```text
temperature=0.2
top_p=0.8
top_k=20
min_p=0
presence_penalty=1.5
repeat_penalty=1.0
reasoning=off
reasoning_budget=0
chat_template_kwargs={"enable_thinking":false}
max_tokens=16384
```

Fallback should remain shorter and more deterministic than `stable`, while
still leaving enough room for local reviewable Markdown when the source is long.

## Output Length

Qwen docs and model cards commonly use large output budgets such as `32768`
tokens for general Qwen3 generation and higher budgets for benchmark-style
complex math or programming. For `nemo-knows`, large budgets are risky unless
reasoning is strictly controlled: they increase latency and give the model room
to spend the entire budget on thinking instead of final Markdown.

Default maintenance should therefore prefer larger local budgets with strict
non-thinking controls:

```text
stable: 32768
fallback: 16384
fast: 2048
deep: 65536 with bounded reasoning
```

## Project Observation

Before reasoning controls were wired into `nemo`, the `tooling-stack` candidate
generation spent its output budget on visible `[Start thinking]` content while
trying to reconcile `Tooling Stack` with a source titled `LLM Wiki`. After
passing non-thinking llama.cpp parameters, the same E2E candidate-generation path
completed in about 26 seconds with no visible thinking markers in the primary
raw outputs.

## Alignment Notes

This local recommendation intentionally differs from the general Qwen examples
in a few places:

- Qwen examples often show thinking-capable settings because they demonstrate
  the model's reasoning behavior. `nemo` defaults to non-thinking for wiki
  maintenance because final Markdown is the artifact.
- Qwen's llama.cpp example uses `--color`; `nemo` omits it because colored
  subprocess output complicates raw-output debugging and cleaning.
- Qwen's example uses `-hf`; `nemo` uses a pinned local GGUF path for
  reproducibility.
- Qwen documents `-fa` and `-sm row` as useful speed settings. They are not yet
  default in `nemo` because the current single-GPU local path is already stable;
  they remain safe candidates for a later performance pass.
