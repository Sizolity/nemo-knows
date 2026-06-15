# AGENTS.md — project schema for `nemo-knows`

This file is the **contract** for LLM agents operating on this
repository. It covers the development infrastructure — the Go CLI,
pipeline testing, deployment, and repository conventions.

For the **wiki maintenance contract**, see [`wiki/AGENTS.md`](wiki/AGENTS.md).
The wiki is a self-contained system; its schema lives inside it.

If a rule here conflicts with a user instruction in chat, the user wins
— but tell the user there is a conflict so the schema can be updated.

## 0. Mental model

This repository has two distinct surfaces:

- **`wiki/`** is the product — a self-contained, LLM-maintained
  Markdown knowledge base. An agent maintaining the wiki reads
  `wiki/AGENTS.md`, not this file.
- **Everything else** is development infrastructure for the Go CLI
  (`nemo`) that builds and maintains the wiki.

```
nemo-knows/
├── wiki/                 # THE PRODUCT (see wiki/AGENTS.md)
├── pipeline/             # ingest pipeline test infrastructure
│   ├── raw/              #   test source material (immutable)
│   ├── drafts/           #   model-output buffers
│   └── evals/            #   evaluation harness
├── tmp/                  # ad-hoc test scratch space (gitignored)
├── cmd/                  # Go CLI entry points (nemo, nemo-web, nemo-server)
├── internal/             # Go packages
├── prompts/              # prompt templates for the ingest pipeline
├── deploy/               # deployment scripts (release, systemd)
├── docs/                 # architecture and development notes
└── AGENTS.md             # this file
```

## 1. Key boundaries

- **Do not modify `pipeline/raw/`.** It is immutable test source
  material.
- **`pipeline/drafts/` and `pipeline/evals/runs/` are development
  artifacts.** They are not part of the production runtime. Pipeline
  evaluations and stress runs do not produce wiki log entries. A log
  entry in `wiki/log.md` is required only when reviewed content is
  actually applied to `wiki/`.
- **`wiki/` is the product.** The autonomous maintainer
  (`nemo -maintain-wiki`) operates on `wiki/` alone and never reads
  `pipeline/` or `tmp/`.
- **Do not run `git push` automatically.** The user controls what
  leaves this machine.
- **Do not commit secrets.** `.env` is gitignored; `.env.example` is a
  template.

## 2. Go CLI

The `nemo` CLI supports draft generation, review, deterministic
evaluation, approved wiki writes, and autonomous maintenance. It works
with local `llama.cpp` models and DeepSeek's API.

```sh
go build -o .bin/nemo ./cmd/nemo
go build -o .bin/nemo-web ./cmd/nemo-web

go test ./...
```

Development pipeline (for testing prompts and review logic):

```sh
.bin/nemo -provider llama -source pipeline/raw/example.md \
  -bundle-dir pipeline/drafts/example -profile stable

.bin/nemo -review-bundle pipeline/drafts/example \
  -out pipeline/drafts/example/apply-plan.md

.bin/nemo -eval-bundle pipeline/drafts/example \
  -out-dir pipeline/evals/runs/example

.bin/nemo -apply-approved pipeline/drafts/example -approve
```

Wiki maintenance:

```sh
.bin/nemo -lint-wiki -out-dir pipeline/evals/runs/wiki-lint
.bin/nemo -maintain-wiki -mode report -out-dir pipeline/evals/runs/wiki-maint
.bin/nemo -maintain-wiki -mode safe -out-dir pipeline/evals/runs/wiki-maint
```

## 3. Prompt templates

Templates under `prompts/` use `{{PLACEHOLDER}}` syntax. They are
rendered by `internal/prompt` and sent to the model backend. See
`docs/development/local-ingest-mvp.md` for the full list of variables.

## 4. Deployment

See `deploy/` for release scripts and systemd units. The primary server
path pulls source over SSH and builds locally.

## 5. What not to do

- Do not modify, rename, or delete anything under `pipeline/raw/`.
- Do not treat `pipeline/drafts/` or `pipeline/evals/runs/` as durable
  storage. They are development buffers.
- Do not commit secrets or credentials.
- Do not run `git push` automatically.
- Do not collapse `wiki/log.md` or rewrite past entries.

## 6. Open questions

- When does a development pipeline stage graduate to a production
  workflow?
- How should the web console handle the pipeline directory migration?
