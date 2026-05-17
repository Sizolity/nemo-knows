# Minimal Go Implementation

This document describes the first code implementation for `nemo-knows`.
It is a development design, not a wiki knowledge page.

## Goal

Build a small Go command that makes the current manual draft workflow
repeatable:

```text
raw source -> prompt render -> llama.cpp -> raw model output -> cleaned draft
```

The command should write both:

```text
drafts/<name>.raw.txt
drafts/<name>.md
```

The raw file keeps the full llama.cpp output for debugging. The Markdown file
is the cleaned candidate draft for human or agent review before anything is
moved into `wiki/`.

## Non-Goals

The first implementation should not include:

- Cobra or another CLI framework.
- Eino or another LLM orchestration framework.
- A web server.
- A database.
- Automatic git commits.
- Automatic writes from model output directly into `wiki/`.
- Embeddings, vector search, or retrieval infrastructure.

## Why No CLI Framework Yet

The initial command surface is small enough for Go's standard library:

- `flag` for CLI arguments.
- `os/exec` for calling `llama-cli`.
- `os`, `io/fs`, and `path/filepath` for file operations.
- `strings` or `text/template` for prompt rendering.
- `regexp` and string processing for output cleaning.
- `testing` for unit tests.

Avoiding a framework keeps the implementation close to the workflow and makes
it easier to inspect, test, and replace individual pieces.

## Proposed Command

```sh
go run ./cmd/nemo \
  -source raw/llm-wiki.md \
  -prompt prompts/source-page.md \
  -out drafts/llm-wiki-source.md
```

Expected outputs:

```text
drafts/llm-wiki-source.raw.txt
drafts/llm-wiki-source.md
```

The next local-only maintenance check uses bundle mode:

```sh
go run ./cmd/nemo \
  -source raw/llm-wiki.md \
  -bundle-dir drafts/llm-wiki-ingest \
  -profile stable
```

Bundle mode generates source and ingest-plan drafts without writing to `wiki/`.
See [`local-ingest-mvp.md`](local-ingest-mvp.md) for the detailed workflow.

## Proposed Layout

```text
cmd/nemo/
  main.go

internal/config/
  config.go

internal/prompt/
  render.go
  render_test.go

internal/llama/
  cli.go
  cli_test.go

internal/draft/
  clean.go
  clean_test.go
  write.go

internal/wiki/
  paths.go
  validate.go
```

## Package Responsibilities

### `cmd/nemo`

The entrypoint only:

- Parses flags.
- Constructs dependencies.
- Runs the flow.
- Prints errors.
- Returns the process exit code.

It should not contain prompt rendering, llama.cpp process logic, or Markdown
cleaning logic.

### `internal/config`

Owns default local settings:

```text
LlamaCLI=/home/karo/src/llama.cpp/build/bin/llama-cli
LlamaModel=/home/karo/models/qwen3.5-9b-q4_k_m.gguf
GPULayers=all
MaxTokens=2048
Temp=0.2
TopP=0.9
```

Later, these defaults can be overridden by environment variables:

```text
NEMO_LLAMA_CLI
NEMO_LLAMA_MODEL
```

### `internal/prompt`

Renders prompt templates. The first renderer only needs to replace:

```text
RAW_SOURCE_PATH
RAW_SOURCE_CONTENT
```

Later it can support:

```text
CONCEPT_NAME
SOURCE_LIST
SOURCE_CONTENT
```

### `internal/llama`

Wraps local model generation behind an interface:

```go
type Generator interface {
    Generate(ctx context.Context, prompt string) (string, error)
}
```

The first implementation calls the verified CUDA build:

```text
/home/karo/src/llama.cpp/build/bin/llama-cli
```

Command shape:

```text
llama-cli -m <model> -p <prompt> -n 2048 -ngl all --single-turn --temp 0.2 --top-p 0.9
```

### `internal/draft`

Cleans and writes model outputs. It should:

- Preserve the raw output.
- Remove `[Start thinking] ... [End thinking]` from the final draft.
- Remove llama.cpp performance lines such as `[ Prompt: ... ]`.
- Remove `Exiting...`.
- Remove `common_memory_breakdown_print` runtime logs.
- Extract content from a complete YAML frontmatter block.
- If final output is incomplete, recover a complete fenced Markdown block when one exists.

### `internal/wiki`

Starts with validation only:

- Validate raw paths.
- Validate wiki paths.
- Validate frontmatter.

It should not automatically write model output into `wiki/` in the first
version. Moving a draft into `wiki/` remains a reviewed step.

## First Implementation Sequence

1. Initialize the Go module.
2. Implement `internal/draft.Clean`.
3. Translate the existing Python cleaner tests into Go tests.
4. Implement `internal/prompt.Render`.
5. Implement `internal/llama.CLI`.
6. Implement `cmd/nemo/main.go`.
7. Run the command on `raw/llm-wiki.md` with `prompts/source-page.md`.
8. Review `drafts/llm-wiki-source.raw.txt` and `drafts/llm-wiki-source.md`.

Start with `internal/draft.Clean`; it is independent of llama.cpp, already has
observed failure cases, and can be covered with unit tests.

## Future Extension Points

Cobra becomes useful when the command surface grows:

```text
nemo draft
nemo ingest
nemo query
nemo lint
nemo config
```

Eino becomes useful when the project needs workflow orchestration:

```text
multi-step LLM flows
tool calling
retrievers
agent graphs
human-in-the-loop review
multiple model backends
```

Until then, a small Go standard-library implementation keeps the project close
to its current file-first workflow.
