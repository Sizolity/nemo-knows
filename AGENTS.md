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

- **`wiki/`** is the LLM-maintained Markdown knowledge base that `nemocli`
  builds and maintains. An agent maintaining wiki *content* reads
  `wiki/AGENTS.md`, not this file.
- **Everything else** is development infrastructure for the Go CLI
  (`nemocli`) that builds and maintains the wiki.

> **Current phase — `wiki/` is the test corpus.** Its contents are
> presently **test data, not production product content**. Test the core
> framework **black-box, directly against `wiki/`** (run real `nemocli`
> commands and observe real effects); treat `wiki/` changes made while
> testing as expected, not as production edits.

```
nemo-knows/
├── wiki/                 # LLM-maintained KB; CURRENT TEST CORPUS (see wiki/AGENTS.md)
├── pipeline/             # legacy ingest-pipeline scaffold (deprecated as primary test path)
│   ├── raw/              #   test source material (immutable)
│   ├── drafts/           #   model-output buffers
│   └── evals/            #   evaluation harness
├── tmp/                  # transient scratch (gitignored); clean up periodically
├── cmd/                  # Go CLI entry points (nemocli, nemo-web, nemo-server)
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
- **Production wiki workflows are wiki-first.** They read and write `wiki/`
  directly, may use `tmp/` for transient debug or review artifacts, and should
  not depend on `pipeline/` unless the user is explicitly running a development
  or stability-evaluation pipeline.
- **`wiki/` is the maintained knowledge base** (currently the test corpus,
  see Mental model). The autonomous maintainer (`nemocli maintain --mode safe`) uses
  `wiki/` as its knowledge input, never reads `pipeline/`, and may write
  transient reports or debug artifacts under `tmp/`.
- **Do not run `git push` automatically.** The user controls what
  leaves this machine.
- **Do not commit secrets.** `.env` is gitignored; `.env.example` is a
  template.

## 2. Go CLI

The `nemocli` CLI supports draft generation, review, deterministic
evaluation, approved wiki writes, and autonomous maintenance. It works
with local `llama.cpp` models and DeepSeek's API.

```sh
go build -o .bin/nemocli ./cmd/nemocli
go build -o .bin/nemo-web ./cmd/nemo-web

go test ./...
```

**Testing convention (current).** Prefer **black-box testing directly
against `wiki/`**: run the real `nemocli` pipeline end-to-end on `wiki/`
content and observe actual effects (apply → index/log → lint). The
`pipeline/` scaffold below is **legacy / deprecated as the primary test
path** — kept for reference and optional reuse. Put intermediate or debug
artifacts under `tmp/` (transient quality/stability verification only),
clean them up when done, and keep production-runtime artifacts minimal.

**Two independent products.** `cmd/nemo-server` and `cmd/nemo-web` continue
to be maintained under their original names; they are not in nemocli's
scope. See `deploy/systemd/install-user-units.sh` for their service units.

Legacy development pipeline (optional, for prompt/review-logic experiments):

```sh
.bin/nemocli init                                  # 一次,首跑骨架
.bin/nemocli ingest pipeline/raw/example.md        # 端到端 ingest
.bin/nemocli lint                                  # read-only 审计
.bin/nemocli maintain --mode safe                  # 确定性维护
```

Wiki maintenance:

```sh
.bin/nemocli lint --out-dir tmp/wiki-lint
.bin/nemocli maintain --mode report --out-dir tmp/wiki-maint
.bin/nemocli maintain --mode safe --out-dir tmp/wiki-maint
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
- Do not treat `tmp/` as durable storage. It holds transient debug /
  quality-verification artifacts; clean it up periodically.
- Do not commit secrets or credentials.
- Do not run `git push` automatically.
- Do not collapse `wiki/log.md` or rewrite past entries.
- Do not point `$NEMO_WIKI_ROOT` at the repo root in production; the in-repo
  `wiki/` is the development corpus.

## 6. Open questions

- When does nemocli's V1 daemon (`nemocli serve`) replace the systemd timer
  + `nemocli once` pattern in production?
- How should the web console handle the pipeline directory migration?
