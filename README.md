# nemo-knows

> A persistent, LLM-curated Markdown wiki with a local ingest pipeline, review
> gates, deterministic evaluation, and an optional browser console.

`nemo-knows` turns raw source files into a maintained knowledge layer. Instead
of re-running retrieval and synthesis on every question, it pays the synthesis
cost at ingest time, stores the result as interlinked Markdown, and lets future
queries read from a wiki that already contains summaries, concepts, topic
syntheses, citations, and known disagreements.

The pattern follows Andrej Karpathy's
[`llm-wiki.md`](https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f),
with an opinionated file schema and a small Go toolchain around it.

## What This Project Is

`nemo-knows` is both a knowledge-base convention and a local tool:

- A plain-file wiki under `wiki/`, maintained by LLM agents according to
  [`AGENTS.md`](AGENTS.md).
- An immutable source area under `raw/`, where original documents live.
- A review buffer under `drafts/`, where model output lands before it can be
  accepted into the wiki.
- A Go CLI, `nemo`, for draft generation, review, evaluation, linting, and
  explicitly approved wiki writes.
- A local web console, `nemo-web`, for browsing the wiki graph, importing
  source-like pages, and starting reviewable ingest jobs.

The project deliberately uses Markdown, Git, and local files first. It does not
require a database, vector store, search service, or always-on backend.

## Repository Structure

```text
nemo-knows/
├── raw/                  # Immutable source material; automation reads only.
├── wiki/                 # Accepted LLM-maintained knowledge base.
│   ├── index.md          # Human-readable catalogue.
│   ├── log.md            # Append-only audit log for formal wiki workflows.
│   ├── sources/          # One page per ingested source.
│   ├── entities/         # People, organisations, products, places.
│   ├── concepts/         # Definitions and mechanisms.
│   └── topics/           # Cross-cutting synthesis pages.
├── drafts/               # Generated bundles and candidate pages for review.
├── evals/
│   ├── cases/            # Regression fixtures.
│   └── runs/             # Local evaluation outputs, ignored by Git.
├── prompts/              # Prompt templates used by the ingest pipeline.
├── cmd/
│   ├── nemo/             # Main CLI.
│   └── nemo-web/         # Standalone web-console launcher.
├── internal/             # Go packages for apply, eval, web, config, lint, etc.
├── docs/                 # Architecture and development notes.
├── .cloudflare/          # Cloudflare Tunnel + Worker deployment guide.
└── AGENTS.md             # Contract for agents maintaining the wiki.
```

The most important boundary is `raw/ -> drafts/ -> wiki/`: source material is
preserved, generated output is reviewed, and accepted wiki writes are explicit.

## Core Capabilities

### Agent-Maintained Wiki

LLM agents use [`AGENTS.md`](AGENTS.md) as the maintenance contract. The
contract defines the canonical workflows:

- **Ingest:** read a source, write a source summary, update concept/topic/entity
  pages, refresh `wiki/index.md`, and append a formal entry to `wiki/log.md`.
- **Query:** answer from `wiki/index.md` and relevant wiki pages first, with
  citations back to maintained pages and raw sources.
- **Lint:** report broken links, orphans, stale claims, missing concepts, and
  contradictions before changing wiki content.

### Local Ingest CLI

The `nemo` CLI can generate source summaries, ingest plans, reviewed apply
plans, candidate concept/topic pages, deterministic evaluation reports, and
approved wiki writes. It supports local `llama.cpp` models and DeepSeek's
OpenAI-compatible API.

Typical reviewable pipeline:

```sh
go build -o .bin/nemo ./cmd/nemo

.bin/nemo -provider llama -source raw/example.md \
  -bundle-dir drafts/example -profile stable

.bin/nemo -review-bundle drafts/example \
  -out drafts/example/apply-plan.md

.bin/nemo -provider llama -generate-candidates drafts/example \
  -profile stable

.bin/nemo -eval-bundle drafts/example \
  -out-dir evals/runs/example/bundle

.bin/nemo -eval-candidates drafts/example \
  -out-dir evals/runs/example/candidates

.bin/nemo -review-candidates drafts/example \
  -out-dir evals/runs/example/review
```

Applying to `wiki/` is intentionally separate and requires explicit approval:

```sh
.bin/nemo -apply-approved drafts/example -approve
```

### Deterministic Evaluation And Linting

The evaluation harness checks generated artifacts without calling a model. It is
used to compare prompt and pipeline changes over time, and to keep apply
readiness auditable. Current checks cover frontmatter, durable source
references, candidate paths, duplicated content, wikilink safety, index
consistency, and apply gates.

Useful commands:

```sh
.bin/nemo -eval-regression evals/cases -out-dir evals/runs/regression
.bin/nemo -lint-wiki -out-dir evals/runs/wiki-lint
.bin/nemo -maintain-wiki -mode report -out-dir evals/runs/wiki-maint
```

`-maintain-wiki -mode safe` is conservative: it only applies mechanical safe
fixes, such as index consistency. Semantic repairs remain review tasks.

### Local Web Console

The web console provides a browser UI for browsing the wiki and starting a
reviewable ingest job.

```sh
# In-process web server from the main CLI.
go run ./cmd/nemo -serve

# Dedicated launcher; uses .bin/nemo or NEMO_BINARY for ingest subprocesses.
go run ./cmd/nemo-web -addr 127.0.0.1:8787
```

Open `http://127.0.0.1:8787`. The console can:

- Render `wiki/`, `raw/`, and `drafts/` Markdown pages.
- Resolve `[[wikilinks]]` for local navigation.
- Show a simple knowledge graph.
- Save imported pages under `wiki/sources/`.
- Start a background bundle-generation and review job.

The web console does not apply reviewed output to `wiki/`. Accepted writes still
go through the explicit CLI apply workflow.

### Cloudflare Exposure

The optional deployment shape is:

```text
User -> Cloudflare Worker cache -> Cloudflare Tunnel -> nemo-web on localhost
```

See [`.cloudflare/README.md`](.cloudflare/README.md) for setup details. The
Worker caches read-heavy GET routes such as `/view`, `/graph`, and `/static`,
while POST routes such as `/run` and `/build` pass through to the local console.
The documented server path uses a Git-over-SSH `systemd` timer that pulls the
default branch and builds locally, so GitHub never SSHes into the host. Release
packages remain an optional fallback for servers that can download HTTPS assets.

## Configuration

`nemo` reads `.env` and environment variables, then lets CLI flags override the
provider per invocation.

Provider selection:

```text
NEMO_MODEL_PROVIDER=llama     # default local llama.cpp backend
NEMO_MODEL_PROVIDER=deepseek  # hosted DeepSeek backend
```

Common overrides:

```text
NEMO_LLAMA_CLI=/path/to/llama-cli
NEMO_LLAMA_MODEL=/path/to/model.gguf
NEMO_DEEPSEEK_API_KEY=...
NEMO_DEEPSEEK_BASE_URL=https://api.deepseek.com
NEMO_DEEPSEEK_MODEL=deepseek-v4-pro
NEMO_CHUNKED_THRESHOLD_CHARS=90000
NEMO_MAX_CHUNK_CHARS=18000
```

Generation profiles are `fast`, `stable`, `deep`, and `fallback`. For
multi-stage runs, pass `-provider llama` or `-provider deepseek` on every stage
so later environment changes cannot switch backends mid-pipeline.

More detail:

- [`docs/development/local-ingest-mvp.md`](docs/development/local-ingest-mvp.md)
- [`docs/development/deepseek-model-config.md`](docs/development/deepseek-model-config.md)
- [`docs/architecture/llm-wiki-core-concept.md`](docs/architecture/llm-wiki-core-concept.md)

## Usage Patterns

### Use Only The Wiki Convention

1. Put source documents under `raw/`.
2. Open the repository in an LLM agent that will follow [`AGENTS.md`](AGENTS.md).
3. Ask it to ingest a source, query the wiki, or run a lint pass.
4. Review the resulting `wiki/` changes and keep them in Git.

This path requires no Go binary and no server.

### Use The CLI For Reviewable Drafts

Use the CLI when you want a repeatable local pipeline, deterministic evaluation,
or an explicit apply gate. Generated artifacts stay under `drafts/` and
`evals/runs/` until approved.

### Use The Browser Console

Use the web console when you want local browsing, graph navigation, and a UI for
starting ingest jobs. It is a convenience layer over the same draft-first
workflow, not a replacement for reviewed apply.

## Capability Boundaries

`nemo-knows` is useful today for building and maintaining a Markdown knowledge
base, but its boundaries are intentional:

- It is not a RAG server, chatbot service, vector database, or search engine.
- It does not automatically trust model output; model-generated pages go to
  `drafts/` first.
- It does not modify `raw/`; optional web-source persistence is create-only.
- It does not provide a top-level `nemo query` command yet. Querying is still an
  agent workflow over the maintained wiki.
- The web console does not run the full candidate/eval/apply chain.
- Only one ingest pipeline should run at a time, because local model execution
  can saturate GPU memory.
- Default local llama paths are machine-specific; real installations should set
  their own environment overrides.
- Cloudflare exposure does not include built-in authentication. Use Cloudflare
  Access or another access-control layer before exposing private material.
- The schema is still pre-v0.1 and may change as more sources and workflows are
  tested.

Known open design areas include wiki search at scale, better long-book ingest,
per-source confidence weighting, and structured operational-log ingestion.

## Development

This repository is a Go module:

```sh
go test ./...
go build -o .bin/nemo ./cmd/nemo
go build -o .bin/nemo-web ./cmd/nemo-web
```

CI runs `gofmt`, `go test ./...`, and `staticcheck`. CD builds Linux `nemo` and
`nemo-web` packages for `amd64` and `arm64` and publishes the latest successful
default-branch package to the `main-latest` GitHub Release channel. The primary
server path pulls source over SSH and builds locally; GitHub does not deploy over
SSH.

## Status

Early, vision-driven, and actively evolving. The file contracts are designed to
be auditable and simple, but conventions should be considered unstable until a
v0.1 release.

## License

TBD by the repo owner. Until a license is added, treat the contents as "all
rights reserved": public reading is fine, derivative use is not yet granted.
