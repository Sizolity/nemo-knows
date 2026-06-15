# nemo-knows

> A persistent, LLM-curated Markdown wiki — a knowledge base that compounds
> over time instead of being re-discovered on every question.

`nemo-knows` is a self-contained, LLM-maintained knowledge base under `wiki/`.
Instead of re-running retrieval and synthesis on every question (RAG), it pays
the synthesis cost once at ingest time, stores the result as interlinked
Markdown, and lets future queries read from a layer that already contains
summaries, concepts, topic syntheses, citations, and known disagreements.

The pattern follows Andrej Karpathy's
[`llm-wiki.md`](https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f),
with an opinionated file schema and a Go toolchain for ingest, evaluation, and
maintenance.

## What This Project Is

The **product** is `wiki/` — a self-contained knowledge base:

```
wiki/
├── index.md          # Catalogue — entry point for every operation
├── log.md            # Append-only audit trail
├── sources/          # One page per ingested external document
├── entities/         # People, organisations, products, places
├── concepts/         # Ideas, mechanisms, definitions
├── topics/           # Cross-cutting syntheses and comparisons
└── assets/           # Generated assets (images, diagrams)
```

The wiki describes itself (index + log), contains its content (sources,
entities, concepts, topics), and tracks its own history (log). `wiki/index.md`
uses ordinary Markdown relative links for direct navigation; body pages may
still use `[[wikilinks]]` as semantic cross-references. All knowledge operations
— adding documents, updating pages, maintaining structure, answering questions
— happen inside it.

Everything else in the repository is **development infrastructure** for the
Go CLI that maintains the wiki:

- `pipeline/` — ingest pipeline test infrastructure: `raw/` (test source
  material, immutable), `drafts/` (model-output buffers), `evals/`
  (evaluation harness)
- `tmp/` — scratch space for ad-hoc development tests
- `cmd/`, `internal/`, `prompts/`, `docs/`, `deploy/` — Go toolchain and docs

The Go CLI (`nemo`) and the autonomous maintainer (`nemo -maintain-wiki`)
operate on `wiki/`. The maintainer never reads `pipeline/` or `tmp/`.

## Core Workflows

The wiki is maintained by LLM agents following [`AGENTS.md`](AGENTS.md):

- **Ingest:** read external material, write a source page under
  `wiki/sources/`, update entity/concept/topic pages, refresh
  `wiki/index.md`, append to `wiki/log.md`.
- **Query:** start from `wiki/index.md`, read relevant pages, synthesize an
  answer with citations. File valuable answers back into `wiki/topics/`.
- **Lint:** scan for contradictions, orphans, stale claims, missing concepts,
  and broken links. Report findings before making changes.

## Go CLI

The `nemo` CLI supports draft generation, review, deterministic evaluation,
approved wiki writes, and autonomous maintenance. It works with local
`llama.cpp` models and DeepSeek's API.

Quick start:

```sh
go build -o .bin/nemo ./cmd/nemo

# Lint the wiki
.bin/nemo -lint-wiki -out-dir tmp/wiki-lint

# Autonomous maintenance (report only)
.bin/nemo -maintain-wiki -mode report -out-dir tmp/wiki-maint

# Autonomous maintenance (apply safe fixes)
.bin/nemo -maintain-wiki -mode safe -out-dir tmp/wiki-maint

# Query the maintained wiki without writing
.bin/nemo -query "How does WAL affect SQLite readers?"

# Draft a filed query answer for review under tmp/query-drafts/
.bin/nemo -query "How does WAL affect SQLite readers?" -file-query

# File a reviewed query answer back into wiki/topics/ with an audit log entry
.bin/nemo -query "How does WAL affect SQLite readers?" -file-query -approve
```

The full development pipeline (bundle → review → eval → candidates → apply) is
documented in `docs/development/`. It routes model output through
`pipeline/drafts/` and `pipeline/evals/` for development testing before anything
reaches `wiki/`. Production wiki-only debugging output should go under `tmp/`.

## Web Console

```sh
go run ./cmd/nemo-web -addr 127.0.0.1:8787
```

Open `http://127.0.0.1:8787` to browse the wiki, follow Markdown index links,
resolve semantic `[[wikilinks]]`, view the knowledge graph, and start background
ingest jobs. The console does not apply output to `wiki/` — accepted writes go
through the explicit CLI apply workflow.

## Configuration

`nemo` reads `.env` and environment variables:

```text
NEMO_MODEL_PROVIDER=llama     # local llama.cpp (default)
NEMO_MODEL_PROVIDER=deepseek  # hosted DeepSeek API
NEMO_LLAMA_CLI=/path/to/llama-cli
NEMO_LLAMA_MODEL=/path/to/model.gguf
NEMO_DEEPSEEK_API_KEY=...
```

See `.env.example` for all options. More detail in
[`docs/development/local-ingest-mvp.md`](docs/development/local-ingest-mvp.md)
and
[`docs/development/deepseek-model-config.md`](docs/development/deepseek-model-config.md).

## Usage Patterns

**Wiki-only (no CLI):** Open the repository in an LLM agent, ask it to ingest a
source, query the wiki, or run a lint pass. The agent follows `AGENTS.md`.

**CLI for reviewable drafts:** Use `nemo` when you want deterministic evaluation
and an explicit apply gate before wiki writes.

**Browser console:** Use `nemo-web` for local browsing, graph navigation, and a
UI for starting ingest jobs.

## Development

```sh
go test ./...
go build -o .bin/nemo ./cmd/nemo
go build -o .bin/nemo-web ./cmd/nemo-web
```

CI runs `gofmt`, `go test ./...`, and `staticcheck`.

## Status

Early and actively evolving. File contracts are designed to be auditable and
simple, but conventions should be considered unstable until a v0.1 release.

## License

TBD by the repo owner.
