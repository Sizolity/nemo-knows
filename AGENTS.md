# AGENTS.md — schema for LLM agents working in `nemo-knows`

This file is the **contract** between the human and any LLM agent
operating on this repository. If you are an LLM, read this in full
before touching any file. If a rule here conflicts with a user
instruction in chat, the user wins — but tell the user there is a
conflict so the schema can be updated.

## 0. Mental model

- `raw/` is **immutable**. Read it. Never modify, rename, or delete
  anything under `raw/` without explicit user permission.
- `wiki/` is **yours to maintain**. Create, update, link, and rearrange
  pages here freely, following the conventions below.
- `AGENTS.md` (this file) is the **schema**. Treat changes to it as
  significant: surface the diff in chat before committing.

The wiki is a *compounding artifact*, not a chat scratchpad. Every edit
should make the wiki more accurate, more interconnected, or more
navigable. If an edit doesn't do at least one of those, don't make it.

## 1. Directory conventions

```
nemo-knows/
├── raw/                        # source material (immutable)
│   └── <free-form>             # filename echoes the source; subdirs OK
├── wiki/                       # LLM-maintained Markdown (your domain)
│   ├── index.md                # content catalogue (you maintain it)
│   ├── log.md                  # append-only operation log
│   ├── sources/                # one page per ingested source
│   ├── entities/               # people, organisations, products, places
│   ├── concepts/               # ideas, mechanisms, definitions
│   ├── topics/                 # cross-cutting syntheses, comparisons
│   └── assets/                 # images and other binary attachments
└── AGENTS.md                   # this file
```

The four wiki subdirectories (`sources/`, `entities/`, `concepts/`,
`topics/`) are the default vocabulary. Add a new subdirectory only when
a category genuinely doesn't fit and you've discussed it with the user;
when you do, document it in this section.

## 2. File conventions

**Naming.** Lowercase, hyphenated, no spaces. `entities/ada-lovelace.md`,
`concepts/retrieval-augmented-generation.md`. Filenames are stable
identifiers — renaming a file means updating every inbound link.

**Frontmatter.** Every wiki page starts with YAML frontmatter:

```yaml
---
title: Ada Lovelace
kind: entity            # one of: source | entity | concept | topic
created: 2026-04-22
updated: 2026-04-22
sources:                # paths under raw/ or wiki/sources/
  - raw/lovelace-bio.md
tags: [history, mathematics]
confidence: high        # high | medium | low
---
```

`kind` mirrors the subdirectory. `sources` lists every `raw/` document
that contributed material to this page, plus any `wiki/sources/` summary
that contributed. `confidence` reflects how solid the page's claims are
after the latest ingest — drop it to `medium` or `low` when sources
disagree or when you're inferring beyond what the sources say.

**Links.** Use Obsidian-style `[[wikilinks]]` for cross-references inside
`wiki/`. Use standard Markdown links with relative paths for references
into `raw/`. Inline citations look like `(see [[ada-lovelace]])` or
`(source: raw/lovelace-bio.md §3)`.

**Length.** Prefer many short, focused pages over one long page. If a
page exceeds ~600 lines or starts covering more than one subject, split
it and update inbound links.

## 3. The ingest workflow

When the user asks you to ingest a source under `raw/<path>`:

1. **Read the source in full.** If it is too long for a single read,
   read it in sections and keep notes; do not summarise from the table
   of contents alone.
2. **Discuss takeaways briefly** with the user. Confirm which angles
   matter to them before writing a permanent page.
3. **Write a source page** at `wiki/sources/<source-slug>.md`. This page
   captures the source itself: what it is, who wrote it, when, key
   claims, and a one-paragraph summary. Frontmatter `kind: source`,
   `sources: [raw/...]`.
4. **Update or create entity / concept / topic pages** that the source
   touches. A typical ingest updates 5–15 wiki pages. For each touched
   page, append the source to its `sources` list and update `updated`.
5. **Update `wiki/index.md`** so any new pages are listed in the right
   category with a one-line summary.
6. **Append to `wiki/log.md`** in the format defined in §6.
7. **Report back** in chat with: which pages were created, which were
   updated, and any contradictions surfaced (see §7).

If the source is short and clearly low-value (e.g., a tweet quoting a
well-covered fact), it is OK to skip step 4 and only write the source
page plus the log entry. Tell the user when you do this.

## 4. The query workflow

When the user asks a question:

1. **Read `wiki/index.md` first.** Decide which pages are relevant.
2. **Read those pages in full** before answering. Do not answer from the
   index entries alone.
3. **Synthesise the answer** in chat. Cite specific wiki pages and, when
   the wiki page itself cites a `raw/` document, propagate the citation.
4. **Offer to file the answer back** as a new `wiki/topics/<slug>.md`
   page when the answer involved real synthesis (e.g., a comparison
   across pages, a derived insight). The user decides whether to file
   it. If they say yes, write the page, update `wiki/index.md`, and
   append to `wiki/log.md` with a `query-filed` action.

If the question cannot be answered from the wiki, say so plainly. Then
suggest sources to ingest that would close the gap.

## 5. The lint workflow

When the user asks for a lint pass (or on a regular cadence agreed in
chat):

1. **Find contradictions.** Read pairs of pages with overlapping
   `sources` and surface any factual disagreements.
2. **Find orphans.** List pages with no inbound `[[wikilinks]]`.
3. **Find stubs.** List pages mentioned in `[[wikilinks]]` but missing
   files.
4. **Find stale claims.** Read pages whose `updated` is older than any
   source they cite; flag for re-review.
5. **Find missing concepts.** Identify terms that recur across many
   pages without their own `concepts/` page.
6. **Report findings** as a chat message with proposed actions. Make no
   edits in this step — wait for the user to approve which findings to
   act on, then do those edits as a normal ingest-style pass and append
   a `lint` entry to `wiki/log.md`.

## 6. The log format

`wiki/log.md` is append-only. Every entry starts with a heading line in
this exact format so it is greppable:

```
## [YYYY-MM-DD] <action> | <subject>
```

`<action>` is one of: `ingest`, `query-filed`, `lint`, `schema-change`,
`note`. `<subject>` is a short human-readable handle (a source title, a
question, etc.).

Below the heading, write a short body: what changed, which pages were
touched (as a bulleted list of relative paths), and any open questions.

Example:

```
## [2026-04-22] ingest | "Attention Is All You Need"
Source: raw/papers/attention-is-all-you-need.pdf
Touched:
- wiki/sources/attention-is-all-you-need.md (created)
- wiki/concepts/self-attention.md (created)
- wiki/concepts/transformer.md (updated)
- wiki/entities/vaswani-et-al.md (created)
- wiki/index.md (updated)
Open: relation to earlier seq2seq work needs its own topic page.
```

To skim recent activity:

```sh
grep "^## \[" wiki/log.md | tail -10
```

## 7. Contradictions and confidence

When a new source contradicts an existing wiki claim, do **not**
silently overwrite. Instead:

1. Keep both claims on the page in a short `## Disagreements` section,
   each with its source citation.
2. Drop the page's `confidence` to `medium` (or `low` if the
   disagreement is fundamental).
3. Note the contradiction in the ingest log entry.

The user decides which side to elevate, by chat. Default policy: **newer
source wins only when the older page's `confidence` is not `high`**;
otherwise wait for the user.

## 8. Writing rules

- **Source the wiki, not the LLM.** Every non-trivial claim on a wiki
  page must trace to either a `raw/` document or another wiki page that
  itself traces to one.
- **Prefer plain prose.** No marketing voice, no rhetorical questions,
  no bullets-of-bullets unless the content is genuinely list-shaped.
- **Be concise.** A wiki page is a reference, not a blog post. Cut every
  sentence that doesn't add information.
- **No emojis** in wiki pages.
- **No fabricated dates, names, or numbers.** If you don't know, say so.
- **No hidden agent commentary** in committed pages. Reasoning belongs
  in chat or in the log; the wiki is for facts and synthesis.

## 9. What not to do

- Do not modify, rename, or delete anything under `raw/`.
- Do not invent `kind` values, subdirectories, or frontmatter fields
  without updating this schema and notifying the user.
- Do not commit secrets, credentials, or anything from `raw/` that the
  user marked private.
- Do not run `git push` automatically. The user controls what leaves
  this machine.
- Do not collapse `wiki/log.md` or rewrite past entries. It is an
  audit trail; only append.

## 10. Open questions for the schema itself

These are deliberate gaps that will be filled as the wiki grows:

- Search over `wiki/` once `wiki/index.md` outgrows its context budget
  (candidate: BM25 + the qmd CLI).
- Ingestion of long sources (books, large repos) that don't fit a
  single read.
- Per-source confidence weights when multiple sources contradict.
- A canonical workflow for ingesting structured operational logs (e.g.
  agent run traces, meeting transcripts, chat exports) where each file
  often produces an entity update rather than a new source page.

When you propose changes to this file, append them under a
`schema-change` log entry.
