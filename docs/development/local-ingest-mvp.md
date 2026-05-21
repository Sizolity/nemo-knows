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

## Long-Source Ingest

When the raw source is longer than the active model can reliably process in one
prompt, bundle generation uses a structure-aware chunked path instead of
truncating the document. The trigger is provider-aware:

- Local `llama.cpp` (default): source size above **90,000 characters**, with
  each chunk capped at **18,000 characters**. These limits were established
  against the 24576-token context window where larger single-shot prompts
  empirically dropped frontmatter or mid-document detail.
- DeepSeek (`NEMO_MODEL_PROVIDER=deepseek`): source size above **300,000
  characters**, with each chunk capped at **60,000 characters**. DeepSeek-V4's
  128K-token input window (~460K ASCII chars) keeps the safety margin >50%
  while avoiding the API-call multiplier of unnecessary chunking. See
  [`deepseek-model-config.md`](deepseek-model-config.md#chunked-bundle-thresholds)
  for the rationale and override knobs.

Either threshold can be overridden by environment variable:

```sh
NEMO_CHUNKED_THRESHOLD_CHARS=200000 NEMO_MAX_CHUNK_CHARS=40000 \
    nemo -provider llama -source raw/large.md -bundle-dir drafts/large -profile stable
```

For multi-stage runs, prefer the CLI `-provider` flag over relying on `.env`
alone. Each pipeline stage is a new `nemo` process, so an edited `.env` can
otherwise change the backend between bundle generation and candidate
generation.

The long-source path writes the normal bundle files plus chunk artifacts:

```text
drafts/<run>/source.md
drafts/<run>/source.raw.txt
drafts/<run>/ingest-plan.md
drafts/<run>/ingest-plan.raw.txt
drafts/<run>/chunks/outline.md
drafts/<run>/chunks/chunk-index.json
drafts/<run>/chunks/chunk-01.md
drafts/<run>/chunks/chunk-01.raw.txt
drafts/<run>/chunks/combined-notes.md
drafts/<run>/chunks/group-01.md
drafts/<run>/chunks/group-01.raw.txt
drafts/<run>/chunks/combined-group-notes.md
```

The group files are present only when the source produces more than one group
of chunk notes.

### Chunking Strategy

The chunker preserves source structure before it considers size:

1. Split on Markdown headings and recognized numbered/named section headings.
2. Pack adjacent short sections into one chunk until the chunk approaches the
   configured maximum size. This avoids wasting context when a long document is
   made of many short sections.
3. If a section is too large, split it by paragraphs.
4. If a paragraph is still too large, split by soft boundaries: line, sentence,
   then word.
5. Fall back to character cuts only for boundaryless text such as generated
   blobs or very long tokens.

The fallback cut does not discard content. It creates multiple segment chunks
with the same heading context so the model can summarize each segment and later
synthesis can recombine the notes.

### Heading Coverage

Packed chunks keep both a primary `heading_path` and complete `heading_paths`
coverage in `chunk-index.json`, `outline.md`, and the rendered chunk prompt.
This matters when one chunk contains many short sections: the model gets the
full local outline instead of only the first section title.

### Group Notes

For very long sources, the final source and ingest-plan synthesis should not
infer whole-document themes from a flat list of many chunk notes. The chunked
path therefore adds a middle layer:

```text
chunk notes -> group notes -> final source.md and ingest-plan.md
```

Group notes summarize adjacent chunks into regional themes while the original
chunk notes remain available for source-backed detail. This is a quality
tradeoff: it adds model calls and runtime, but gives final synthesis a stronger
view of early, middle, and late document structure.

The final chunk synthesis prompts receive:

```text
{{CHUNK_OUTLINE}}
{{CHUNK_INDEX}}
{{CHUNK_GROUP_NOTES}}
{{CHUNK_NOTES}}
```

When group notes are present they are treated as the authoritative
whole-document summary and `{{CHUNK_NOTES}}` is intentionally left empty for
the final source/ingest synthesis. Sending both layers into a single prompt
exceeded the local model's 24,576-token context window on long specifications
(observed 36k-43k tokens for 16-24 chunk sources during round-3 stress
testing) and caused `clean fallback draft: draft output does not contain
complete frontmatter` failures. The combined chunk notes are still written to
disk for human inspection and are still used when the source is short enough
that group notes are skipped.

Use `combined-group-notes.md` to inspect the regional summaries and
`combined-notes.md` to inspect the raw per-chunk notes.

### Real Long-Source Validation

The group-notes path was validated against real public-web corpus sources on
2026-05-19:

| Source | Size | Chunk notes | Group notes | Result |
| --- | ---: | ---: | ---: | --- |
| `raw/web/corpus-2026-05-18/031-effective-go.md` | 99 KB | 7 | 2 | pass |
| `raw/web/corpus-2026-05-18/032-go-modules-reference.md` | 191 KB | 13 | 3 | pass |
| `raw/web/corpus-2026-05-18/060-dom-standard.md` | 468 KB | 32 | 6 | pass |
| `raw/web/corpus-2026-05-18/102-moby-dick.md` | 892 KB | 53 | 9 | pass |

These real runs completed bundle generation, bundle review, bundle eval,
candidate generation, candidate eval, and candidate review. Candidate review
reported `overall: pass` and `items: 0` for 031, 032, and 102; DOM surfaced one
reviewable title/originality issue in a candidate page while the bundle itself
passed.

Round-4 validation on 2026-05-21 raised the real-source ceiling to 892 KB and
the mechanical synthetic ceiling to 1.8 MB (102 chunks, 17 group notes) using a
fake local backend. The synthetic run validates chunk/group plumbing, not model
context quality.

## Resuming A Reviewed Bundle

If a multi-stage run fails after `source.md` and `ingest-plan.md` exist, resume
the post-bundle stages instead of regenerating chunks:

```sh
nemo -provider llama \
  -resume drafts/<run-id> \
  -out-dir evals/runs/<run-id>
```

Resume checks for `apply-plan.md`, bundle eval output, candidate drafts,
candidate eval, and candidate review, then runs only the missing stages. Delete
the bundle manually when a full regeneration is intended.

Detailed local artifacts are recorded in
`evals/runs/real-corpus-2026-05-19-group-notes-summary.md`. These files are test
artifacts, not maintained wiki pages.

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
{{CHUNK_CONTENT}}
{{CHUNK_NOTES}}
{{CHUNK_GROUP_NOTES}}
{{CHUNK_OUTLINE}}
{{CHUNK_INDEX}}
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

The same ingest pipeline can also use DeepSeek through its OpenAI-compatible
API by setting `NEMO_MODEL_PROVIDER=deepseek` and `NEMO_DEEPSEEK_API_KEY`.
The generated model/profile mapping lives in
[`docs/development/deepseek-model-config.md`](deepseek-model-config.md).

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
