# DeepSeek Model Configuration

DeepSeek exposes an OpenAI-compatible API at `https://api.deepseek.com`.
`nemo` can use it by selecting the `deepseek` model provider and supplying an
API key through the environment.

```sh
export NEMO_MODEL_PROVIDER=deepseek
export NEMO_DEEPSEEK_API_KEY="$DEEPSEEK_API_KEY"
export NEMO_DEEPSEEK_BASE_URL=https://api.deepseek.com
```

The settings below are DeepSeek-specific. They should not be treated as shared
defaults for the local `llama.cpp` backend, because DeepSeek uses hosted
OpenAI-compatible chat parameters while `llama.cpp` uses CLI flags and a local
chat template.

## Nemo Profiles

Official model limits for both `deepseek-v4-flash` and `deepseek-v4-pro`:

```json
{
  "context_length": 1000000,
  "max_output_tokens": 384000
}
```

`nemo` keeps the local llama profile budgets separate from DeepSeek. For
DeepSeek, `max_tokens` defaults to the official maximum output length. Thinking
mode is enabled by default in the DeepSeek API; the `fast` and `fallback`
profiles disable it explicitly, while `stable` and `deep` keep it enabled with
the official default `high` reasoning effort.

```json
{
  "provider": "deepseek",
  "base_url": "https://api.deepseek.com",
  "api_key_env": "NEMO_DEEPSEEK_API_KEY",
  "response_format": {
    "type": "text"
  },
  "profiles": {
    "fast": {
      "model": "deepseek-v4-flash",
      "context_length": 1000000,
      "max_tokens": 384000,
      "thinking": {
        "type": "disabled"
      },
      "temperature": 0.2,
      "top_p": 0.9
    },
    "stable": {
      "model": "deepseek-v4-pro",
      "context_length": 1000000,
      "max_tokens": 384000,
      "thinking": {
        "type": "enabled"
      },
      "reasoning_effort": "high"
    },
    "deep": {
      "model": "deepseek-v4-pro",
      "context_length": 1000000,
      "max_tokens": 384000,
      "thinking": {
        "type": "enabled"
      },
      "reasoning_effort": "high"
    },
    "fallback": {
      "model": "deepseek-v4-flash",
      "context_length": 1000000,
      "max_tokens": 384000,
      "thinking": {
        "type": "disabled"
      },
      "temperature": 0.2,
      "top_p": 0.8
    }
  }
}
```

The current DeepSeek docs also list an Anthropic-compatible base URL,
`https://api.deepseek.com/anthropic`, but `nemo` uses the OpenAI-compatible
`/chat/completions` shape.

## Request Shape

DeepSeek's `/chat/completions` endpoint is stateless. Each request must include
the messages required for that turn. `nemo` sends the rendered prompt as a
`user` message and can optionally prepend a DeepSeek-only system prompt:

```json
{
  "model": "deepseek-v4-pro",
  "messages": [
    {
      "role": "system",
      "content": "Return concise Markdown suitable for a wiki draft."
    },
    {
      "role": "user",
      "content": "<rendered nemo prompt>"
    }
  ],
  "thinking": {
    "type": "enabled"
  },
  "reasoning_effort": "high",
  "max_tokens": 384000,
  "stream": false
}
```

The optional `system` message is configured with
`NEMO_DEEPSEEK_SYSTEM_PROMPT`. It is intentionally separate from local prompt
templates so DeepSeek-specific steering does not alter the llama backend.

## Overrides

Use these environment variables to override the generated profile defaults:

```sh
export NEMO_DEEPSEEK_MODEL=deepseek-v4-pro
export NEMO_DEEPSEEK_MAX_TOKENS=384000
export NEMO_DEEPSEEK_THINKING=enabled
export NEMO_DEEPSEEK_REASONING_EFFORT=high
export NEMO_DEEPSEEK_RESPONSE_FORMAT=text
export NEMO_DEEPSEEK_USER_ID=nemo-local
export NEMO_DEEPSEEK_SYSTEM_PROMPT="Return concise Markdown suitable for a wiki draft."
```

When `thinking.type` is `enabled`, DeepSeek ignores `temperature`, `top_p`,
`presence_penalty`, and `frequency_penalty`; `nemo` omits sampling parameters in
that mode. Sampling parameters are only sent for non-thinking profiles.

For non-thinking profiles, these DeepSeek-only sampling overrides can be used:

```sh
export NEMO_DEEPSEEK_THINKING=disabled
export NEMO_DEEPSEEK_TEMPERATURE=0.2
export NEMO_DEEPSEEK_TOP_P=0.9
```

## JSON Output

DeepSeek supports structured JSON output via:

```json
{
  "response_format": {
    "type": "json_object"
  }
}
```

When this mode is enabled, the official docs require the prompt itself to
contain the word `json` and an example of the expected JSON shape. Otherwise the
model may keep emitting whitespace until the token limit. Use it only with
prompts designed for JSON:

```sh
export NEMO_DEEPSEEK_RESPONSE_FORMAT=json_object
export NEMO_DEEPSEEK_SYSTEM_PROMPT="Return only valid JSON matching the schema in the user prompt."
```

## Usage And Cache Fields

DeepSeek responses include usage fields that are useful for cost and cache
debugging:

```json
{
  "usage": {
    "prompt_tokens": 1000,
    "prompt_cache_hit_tokens": 750,
    "prompt_cache_miss_tokens": 250,
    "completion_tokens": 300,
    "total_tokens": 1300,
    "completion_tokens_details": {
      "reasoning_tokens": 120
    }
  }
}
```

The request-side `user_id` is also DeepSeek-specific. It can isolate KVCache
usage between application users, but should not contain private user data.

## Tool Calls Reference

DeepSeek also supports OpenAI-style `tools` and `tool_choice`. `nemo` does not
send tools yet, but future DeepSeek-only integrations should follow this shape:

```json
{
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "read_wiki_page",
        "description": "Read a wiki page by path.",
        "parameters": {
          "type": "object",
          "properties": {
            "path": {
              "type": "string",
              "description": "The wiki page path."
            }
          },
          "required": ["path"],
          "additionalProperties": false
        }
      }
    }
  ],
  "tool_choice": "auto"
}
```

For DeepSeek's Beta strict mode, use the beta base URL and set
`function.strict` to `true`. Strict schemas require every object property to be
listed in `required` and `additionalProperties` to be `false`.

The DeepSeek documentation notes that the older `deepseek-chat` and
`deepseek-reasoner` model names are compatibility aliases and are scheduled for
deprecation on 2026-07-24, so new configuration should use `deepseek-v4-flash`
or `deepseek-v4-pro`.

## Chunked Bundle Thresholds

The bundle command switches from single-shot generation to a multi-stage chunked
path when the raw source exceeds a configurable character threshold. The
defaults are provider-aware because DeepSeek's input window dwarfs the local
`llama.cpp` 24576-token context window.

| Provider | `ChunkedBundleCharThreshold` | `MaxChunkChars` | Rationale |
| --- | ---: | ---: | --- |
| `llama` | 90,000 | 18,000 | Empirically the largest single-shot prompt that the local 24576-token context can finish without dropping frontmatter or mid-document detail. |
| `deepseek` | 300,000 | 60,000 | DeepSeek-V4 accepts 128K input tokens (~460K ASCII chars). A 300K-char ceiling keeps a >50% safety margin for prompt scaffolding and reasoning tokens while avoiding the N-times API-call cost of unnecessary chunking. |

Why chunking still matters even when the model can technically fit the input:

- Long-context LLMs still exhibit "lost in the middle" attention degradation.
- Per-chunk `.raw.txt` artifacts preserve audit granularity for human review.
- A failed single chunk can be re-run without redoing the whole document.

But chunking has a real cost on DeepSeek that it does not have on local llama:
each chunk is a separately billed API call with its own thinking-mode reasoning
budget, and per-chunk latency adds up sequentially. Lifting the threshold for
DeepSeek mostly trades chunk-level auditability for raw cost and wall-clock
time.

### Overriding the thresholds

Both knobs are exposed as environment variables:

```sh
export NEMO_CHUNKED_THRESHOLD_CHARS=200000
export NEMO_MAX_CHUNK_CHARS=40000
```

`NEMO_CHUNKED_THRESHOLD_CHARS` overrides the source-size cutoff that triggers
the chunked path. `NEMO_MAX_CHUNK_CHARS` overrides the per-chunk size cap used
by the chunker once the chunked path is active. Both accept positive integers;
invalid or non-positive values fall back to the provider default.
