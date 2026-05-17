# Local Ingest MVP

This document describes the second MVP: using the local 9B llama.cpp model to
prepare an ingest draft bundle without directly modifying the maintained wiki.

## Goal

Validate whether the local model can support basic knowledge-base maintenance
work by producing reviewable ingest artifacts:

```text
raw source -> local model -> draft bundle -> human/agent review -> wiki edits
```

This stage is about maintenance assistance, not fully autonomous maintenance.

## Command

```sh
go run ./cmd/nemo \
  -source raw/llm-wiki.md \
  -bundle-dir drafts/llm-wiki-ingest \
  -profile stable
```

Expected outputs:

```text
drafts/llm-wiki-ingest/source.raw.txt
drafts/llm-wiki-ingest/source.md
drafts/llm-wiki-ingest/ingest-plan.raw.txt
drafts/llm-wiki-ingest/ingest-plan.md
```

For public web material that starts outside `raw/`, the bundle command can
explicitly persist a create-only source copy first:

```sh
go run ./cmd/nemo \
  -source drafts/web-e2e-sources/qwen-llama-cpp.md \
  -bundle-dir drafts/web-e2e-qwen \
  -profile stable \
  -persist-raw-web
```

This writes `raw/web/qwen-llama-cpp.md` if it does not already exist, then uses
that durable raw path for prompt rendering. It must not overwrite existing
`raw/` files.

## Bundle Contents

`source.md` is a candidate source summary page. The command normalizes its YAML
frontmatter before writing the draft so reviewed promotion has the fields wiki
lint requires: `title`, `kind: source`, durable `sources`, and
`confidence: medium`. It should still be reviewed before copying any content
into `wiki/sources/`.

`ingest-plan.md` is a candidate plan for wiki maintenance. It should identify
likely source, concept, and topic pages, suggested links, and a review checklist.
It is not itself an accepted wiki page.

The `.raw.txt` files preserve complete llama.cpp output for debugging. They are
useful when the cleaned output is incomplete, over-truncated, or contains
unexpected claims.

## Review Boundary

The bundle command must not write to `wiki/`.

Accepted wiki changes remain a separate reviewed operation:

```text
draft bundle -> review -> wiki source/concept/topic edits -> index update -> log append
```

This protects the knowledge base from model artifacts such as:

- invented dates or links,
- incomplete sections,
- missed contradictions,
- wrong page categorization,
- prompt echoes or reasoning text.

## Prompt Variables

Prompt templates should use braced placeholders:

```text
{{RAW_SOURCE_PATH}}
{{RAW_SOURCE_CONTENT}}
{{CONCEPT_NAME}}
{{SOURCE_LIST}}
{{SOURCE_CONTENT}}
```

The renderer keeps legacy unbraced placeholders for compatibility, but new
templates should use braced placeholders to avoid partial replacement bugs such
as `RAW_SOURCE_CONTENT` becoming `RAW_`.

## Local Model Role

The local 9B model is suitable for:

- source summary drafts,
- first-pass ingest plans,
- extracting candidate concepts,
- generating review checklists,
- producing Markdown that can be cleaned and inspected.

The local model should not yet be trusted to:

- automatically update `wiki/index.md`,
- automatically append to `wiki/log.md`,
- resolve contradictions,
- rewrite multiple existing pages without review,
- decide confidence levels without human or stronger-model review.

## Qwen Generation Guidance

The local model is a Qwen GGUF served through llama.cpp. Qwen's official
guidance is important because many failures in this project come from the model
spending the output budget on reasoning instead of the final Markdown.

Official Qwen3 guidance:

- Qwen3-Instruct-2507 is non-thinking only and does not emit `<think></think>`
  blocks. Recommended parameters: `temperature=0.7`, `top_p=0.8`, `top_k=20`,
  `min_p=0`.
- Qwen3-Thinking-2507 is thinking only. Recommended parameters:
  `temperature=0.6`, `top_p=0.95`, `top_k=20`, `min_p=0`.
- Qwen3 hybrid models can switch thinking behavior. For non-thinking mode,
  Qwen recommends `temperature=0.7`, `top_p=0.8`, `top_k=20`, `min_p=0`.
  For thinking mode, Qwen recommends `temperature=0.6`, `top_p=0.95`,
  `top_k=20`, `min_p=0`.
- Qwen recommends avoiding greedy decoding for thinking mode because it can
  degrade performance and cause endless repetitions.
- Qwen recommends an output length of 16,384 tokens for most non-thinking
  queries, and examples for thinking mode often use 32,768 tokens.
- When supported by the inference runtime, a presence penalty between `0` and
  `2` can reduce repetitions. Higher values can occasionally cause language
  mixing or slightly worse performance.

For Qwen3.5-specific Hugging Face model-card guidance, see
[`docs/development/qwen3-5-huggingface-parameters.md`](qwen3-5-huggingface-parameters.md).

For this project, the default maintenance profile should favor non-thinking
behavior: it is faster, easier to clean, and less likely to spend the entire
budget on reasoning. Thinking mode should be an explicit opt-in for complex
analysis, not the default ingest path.

The 2026-05-16 end-to-end candidate-generation test confirmed this boundary in
practice. The `tooling-stack` candidate used the `stable` profile but the local
Qwen3.5 GGUF emitted visible `[Start thinking]` content and spent its output
budget analyzing a mismatch between the requested page title (`Tooling Stack`)
and the source title (`LLM Wiki`). The generation was not a Go deadlock; it was
llama.cpp continuing normal token generation, while `nemo` waited for
`CombinedOutput()` to return.

For llama.cpp, non-thinking maintenance calls now pass the reasoning controls
that the local binary supports through `internal/config` and `internal/llama`:

```text
--jinja
--reasoning off
--reasoning-budget 0
--chat-template-kwargs {"enable_thinking":false}
--presence-penalty 1.5
--repeat-penalty 1.0
```

If a future llama.cpp build or GGUF template ignores these controls, the next
fallback should be Qwen's documented custom non-thinking chat template via
`--chat-template-file`.

## Task Profiles

The command should use named profiles rather than hard-coded generation
settings:

```text
fast      smoke tests and strict short-format tasks
stable    default local wiki maintenance tasks
deep      complex analysis where thinking is acceptable
fallback  retry path with shorter, more template-like prompts
```

Recommended mapping:

| Profile | Intended use | Max tokens | Temperature | Top-p | Top-k | Min-p | Reasoning |
| --- | --- | ---: | ---: | ---: | ---: | ---: | --- |
| `fast` | smoke tests and quick format checks | 2048 | 0.2 | 0.9 | 20 | 0 | off |
| `stable` | default source and ingest drafts | 32768 | 0.7 | 0.8 | 20 | 0 | off |
| `deep` | complex reasoning or synthesis | 65536 | 0.6 | 0.95 | 20 | 0 | on, budget 2000 |
| `fallback` | retry after malformed output | 16384 | 0.2 | 0.8 | 20 | 0 | off |

`stable` is the default profile for `nemo`. `fast` remains useful for tiny
smoke tests. `deep` is intentionally slower and should be requested explicitly.
`fallback` is used by tool-controlled retries when validation fails.

## Stability Strategy

The system should be more permissive with time and output length, but stricter
about accepting results. The desired behavior is:

1. Generate with the selected profile, defaulting to `stable`.
2. If `-persist-raw-web` is set, copy the source into `raw/web/<slug>.md` before
   rendering prompts. This is explicit, create-only, and keeps candidate source
   attribution durable.
3. Disable thinking for maintenance profiles unless `deep` is explicitly
   selected. The `stable`, `fast`, and `fallback` profiles pass
   `--reasoning off`, `--reasoning-budget 0`, and
   `--chat-template-kwargs {"enable_thinking":false}`.
4. Preserve raw output regardless of success or failure.
5. Clean and validate the Markdown draft.
6. If cleaning or validation fails, retry once with a shorter fallback prompt
   and the `fallback` profile.
7. If the retry still fails, keep the raw output and report the validation
   failure instead of writing to `wiki/`.
8. When a primary generation succeeds, remove any stale `.fallback.raw.txt`
   file from an earlier failed run so debugging artifacts reflect the latest
   execution.

This keeps the model free to use more tokens while ensuring only clean,
reviewable drafts move forward.

## Acceptance Criteria

This MVP is successful when:

- `go test ./...` passes.
- `nemo -bundle-dir` writes the four expected bundle files.
- cleaned drafts contain YAML frontmatter.
- cleaned drafts do not contain prompt echoes, thinking blocks, or llama.cpp
  runtime logs.
- `wiki/` is unchanged by the bundle command.
- web-derived bundle sources can be explicitly persisted to new `raw/web/`
  files, but existing raw files are never overwritten.
- raw outputs remain available for debugging.

## Next Step

After this MVP, the next command can be a reviewed ingest helper. It should read
a draft bundle, validate page frontmatter and required sections, and propose a
file-by-file patch plan. It still should not apply wiki edits without explicit
approval.
