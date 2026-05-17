# nemo-knows

> A persistent, LLM-curated wiki — plain Markdown, no servers, no embeddings required.

`nemo-knows` is a **knowledge layer** you grow over time: a directory of
interlinked Markdown files that an LLM agent writes and maintains as you
feed it sources. It is the long-term counterpart to ephemeral chat
sessions. You read it; the LLM writes it.

The pattern is the one described in Andrej Karpathy's
[`llm-wiki.md`](https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f) —
this repo is one concrete, opinionated instantiation of it.

## Why this exists

Most LLM-with-documents workflows look like RAG: every query re-discovers
relevant fragments from raw sources and re-synthesises an answer. Nothing
accumulates. `nemo-knows` makes the opposite trade: pay the synthesis
cost **once at ingest time**, store the result as plain Markdown, and let
every future query read from a layer that already cross-references,
flags contradictions, and reflects everything you've read.

The wiki is the compounding artifact. Sessions come and go; `nemo-knows`
remembers.

## Three layers

```
nemo-knows/
├── raw/        # immutable source material — the LLM reads, never edits
├── wiki/       # LLM-maintained Markdown — entity / concept / synthesis pages
└── AGENTS.md   # the schema — conventions and workflows for the LLM
```

- **`raw/`** is your source of truth. Articles, papers, transcripts,
  meeting notes, agent run logs — anything you want the wiki to be
  derived from. Files here are **never modified** by the LLM.
- **`wiki/`** is the LLM's working area. It owns every file under here —
  it creates pages, updates them when new sources arrive, maintains
  cross-references with `[[wikilinks]]`, and keeps everything consistent.
- **`AGENTS.md`** is what makes the LLM behave like a disciplined wiki
  maintainer rather than a generic chatbot. It describes directory
  conventions, ingest/query/lint workflows, and writing rules. You and
  the LLM co-evolve it over time.

Two special files inside `wiki/` help navigation as the wiki grows:

- **`wiki/index.md`** — content-oriented catalogue, organised by category.
- **`wiki/log.md`** — chronological, append-only record of every ingest,
  query worth filing, and lint pass.

## Three operations

**Ingest.** You drop a source into `raw/` and tell the LLM to process it.
The LLM reads it, discusses key takeaways with you, writes a summary
page in `wiki/`, updates `wiki/index.md`, updates whichever entity and
concept pages it touches, and appends an entry to `wiki/log.md`. A single
source can touch 5–15 wiki pages.

**Query.** You ask the LLM a question. It reads `wiki/index.md` first
to scope, then drills into the relevant pages, then synthesises an
answer with citations back to `wiki/` and `raw/` paths. **Good answers
should be filed back into the wiki** so future you doesn't redo the work.

**Lint.** Periodically, ask the LLM to health-check the wiki. Look for
contradictions between pages, stale claims that newer sources have
superseded, orphan pages with no inbound links, missing concept pages,
and gaps that a follow-up source could fill.

The exact procedures for each of these live in [`AGENTS.md`](AGENTS.md).

For the engineering-facing version of this concept, see
[`docs/architecture/llm-wiki-core-concept.md`](docs/architecture/llm-wiki-core-concept.md).

## Getting started

There is no installer and no service to run. The whole repo is a Git
tree of Markdown files plus a schema that any capable LLM agent (Claude
Code, Cursor, Codex, OpenCode, etc.) can follow. Concretely:

1. Open this repo in your LLM agent of choice.
2. Drop a source file into `raw/` — for example `raw/some-article.md`.
3. Ask the agent: *"Ingest `raw/some-article.md` per `AGENTS.md`."*
4. Read what it produces under `wiki/`. Push back. Iterate.

For this to work, the agent must follow [`AGENTS.md`](AGENTS.md). That
file is the contract.

There is also a small Go CLI under [`cmd/nemo`](cmd/nemo) for local ingest
drafts, deterministic evaluation, and explicitly approved writes to `wiki/`.
Start with [`docs/development/local-ingest-mvp.md`](docs/development/local-ingest-mvp.md).

## Status

Early, vision-driven. The schema in `AGENTS.md` is the first draft and
will tighten as real sources reveal what's missing. Expect breaking
changes to conventions until v0.1.

## License

TBD by the repo owner. Until a license is added, treat the contents as
"all rights reserved" — public reading is fine, derivative use is not
yet granted.
