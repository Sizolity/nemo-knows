# AGENTS.md — schema for LLM agents maintaining this wiki

This file is the **contract** between the human and any LLM agent
maintaining this knowledge base. It lives inside the wiki because the
wiki is a self-contained system — copy this directory anywhere and an
LLM agent will know how to maintain it.

This is a living document. You and the human co-evolve it over time as
you discover what works for your domain. If a rule here conflicts with a
user instruction in chat, the user wins — but tell the user there is a
conflict so the schema can be updated.

## 0. Mental model

This wiki is a **persistent, compounding artifact**. Unlike RAG, where
the LLM re-discovers knowledge from raw chunks on every query, this wiki
pays the synthesis cost once at ingest time. Knowledge is compiled,
cross-referenced, and kept current. Future queries read from a layer
that already contains summaries, concepts, connections, and known
disagreements.

### Who does what

- **The human** curates sources, directs the analysis, asks the
  questions, and thinks about what it all means. The human decides what
  enters the wiki and how to resolve contradictions.
- **The LLM** does everything else: summarizing, cross-referencing,
  filing, bookkeeping, and maintenance. The LLM owns the wiki layer — it
  writes and maintains all pages.

The pattern works because the tedious part of maintaining a knowledge
base is not the reading or the thinking — it's the bookkeeping. LLMs
don't get bored, don't forget to update a cross-reference, and can
touch many files in one pass.

### How the wiki works

- **`index.md`** is the entry point for every operation. It catalogues
  every page by category. Start here before any query, ingest, or lint
  pass.
- **`log.md`** is the append-only audit trail. Every ingest, filed
  query, lint pass, and schema change is recorded here.
- **`sources/`** holds one page per ingested external document — what it
  is, who wrote it, key claims, a summary. This is where external
  material enters the system.
- **`entities/`** holds pages for people, organisations, products, and
  places.
- **`concepts/`** holds pages for ideas, mechanisms, and definitions.
- **`topics/`** holds cross-cutting syntheses, comparisons, and derived
  insights — including high-value query answers filed back from chat.
- **`assets/`** holds generated assets (images, diagrams).

The wiki describes itself (index + log), contains its content (sources,
entities, concepts, topics), and tracks its own history (log).

## 1. Directory conventions

```
wiki/
├── index.md                # catalogue — entry point for every operation
├── log.md                  # append-only audit trail
├── sources/                # one page per ingested external document
├── entities/               # people, organisations, products, places
├── concepts/               # ideas, mechanisms, definitions
├── topics/                 # cross-cutting syntheses, comparisons
├── assets/                 # generated assets (images, diagrams)
└── AGENTS.md               # this file
```

The four content subdirectories (`sources/`, `entities/`, `concepts/`,
`topics/`) are the default vocabulary. Add a new subdirectory only when a
category genuinely doesn't fit and you've discussed it with the user;
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
sources:                # paths to source material
  - wiki/sources/some-source.md
tags: [history, mathematics]
confidence: high        # high | medium | low
---
```

`kind` mirrors the subdirectory. `sources` lists every document that
contributed material to this page. `confidence` reflects how solid the
page's claims are after the latest ingest — drop it to `medium` or `low`
when sources disagree or when you're inferring beyond what the sources
say.

**Links.** `index.md` is a navigation catalogue and uses standard Markdown
links with paths relative to `index.md`, such as
`[sqlite-wal](sources/sqlite-wal.md)`. Body pages may use Obsidian-style
`[[wikilinks]]` as semantic cross-references between wiki concepts; those
wikilinks identify related page slugs and are not the index navigation format.
Use standard Markdown links with relative paths for non-semantic navigation
or references outside the wiki. Inline citations look like
`(see [[ada-lovelace]])` or `(source: wiki/sources/some-source.md §3)`.

**Length.** Prefer many short, focused pages over one long page. If a
page exceeds ~600 lines or starts covering more than one subject, split
it and update inbound links.

## 3. The ingest workflow

When the user asks you to ingest external material into the wiki:

1. **Read the material in full.** If it is too long for a single read,
   read it in sections and keep notes; do not summarise from the table
   of contents alone.
2. **Discuss takeaways briefly** with the user. Confirm which angles
   matter to them before writing permanent pages.
3. **Write a source page** at `sources/<source-slug>.md`. This page
   captures the material itself: what it is, who wrote it, when, key
   claims, and a one-paragraph summary. Frontmatter `kind: source`,
   `sources` listing the origin.
4. **Update or create entity / concept / topic pages** that the material
   touches. A typical ingest updates 5–15 wiki pages. For each touched
   page, append the source to its `sources` list and update `updated`.
5. **Update `index.md`** so any new pages are listed in the right
   category with a one-line summary.
6. **Append to `log.md`** in the format defined in §6.
7. **Report back** in chat with: which pages were created, which were
   updated, and any contradictions surfaced (see §7).

If the material is short and clearly low-value (e.g., a tweet quoting a
well-covered fact), it is OK to skip step 4 and only write the source
page plus the log entry. Tell the user when you do this.

## 4. The query workflow

When the user asks a question:

1. **Read `index.md` first.** Decide which pages are relevant.
2. **Read those pages in full** before answering. Do not answer from the
   index entries alone.
3. **Synthesise the answer** in chat. Cite specific wiki pages and
   propagate their source citations.
4. **File the answer back into the wiki** when it involved real
   synthesis — a comparison across pages, an analysis, a connection you
   discovered. These are valuable and should not disappear into chat
   history. They compound in the knowledge base just like ingested
   sources do. Write the page under `topics/<slug>.md`, update
   `index.md`, and append to `log.md` with a `query-filed` action.
   Confirm with the user before filing.

If the question cannot be answered from the wiki, say so plainly. Then
suggest material to ingest that would close the gap.

## 5. The lint workflow

When the user asks for a lint pass (or on a regular cadence agreed in
chat):

1. **Find contradictions.** Read pairs of pages with overlapping
   `sources` and surface any factual disagreements.
2. **Find orphans.** List pages with no inbound semantic `[[wikilinks]]`
   and no `index.md` Markdown catalogue entry.
3. **Find stubs.** List pages mentioned in semantic `[[wikilinks]]` but
   missing files.
4. **Find stale claims.** Read pages whose `updated` is older than any
   source they cite; flag for re-review.
5. **Find missing concepts.** Identify terms that recur across many
   pages without their own `concepts/` page.
6. **Report findings** as a chat message with proposed actions. Make no
   edits in this step — wait for the user to approve which findings to
   act on, then do those edits as a normal ingest-style pass and append
   a `lint` entry to `log.md`.

## 6. The log format

`log.md` is append-only. Every entry starts with a heading line in this
exact format so it is greppable:

```
## [YYYY-MM-DD] <action> | <subject>
```

`<action>` is one of: `ingest`, `query-filed`, `lint`, `schema-change`,
`note`. `<subject>` is a short human-readable handle (a source title, a
question, etc.).

Below the heading, write a short body: what changed, which pages were
touched (as a bulleted list of relative paths within the wiki), and any
open questions.

Example:

```
## [2026-04-22] ingest | "Attention Is All You Need"
Touched:
- sources/attention-is-all-you-need.md (created)
- concepts/self-attention.md (created)
- concepts/transformer.md (updated)
- entities/vaswani-et-al.md (created)
- index.md (updated)
Open: relation to earlier seq2seq work needs its own topic page.
```

To skim recent activity:

```sh
grep "^## \[" log.md | tail -10
```

## 7. Contradictions and confidence

When new material contradicts an existing wiki claim, do **not**
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
  page must trace to a documented source (a `sources/` page or another
  wiki page that itself traces to one).
- **Prefer plain prose.** No marketing voice, no rhetorical questions,
  no bullets-of-bullets unless the content is genuinely list-shaped.
- **Be concise.** A wiki page is a reference, not a blog post. Cut every
  sentence that doesn't add information.
- **No emojis** in wiki pages.
- **No fabricated dates, names, or numbers.** If you don't know, say so.
- **No hidden agent commentary** in committed pages. Reasoning belongs
  in chat or in the log; the wiki is for facts and synthesis.

## 9. What not to do

- Do not modify or delete `index.md` or `log.md` structural sections;
  only append to the log and update the index catalogue.
- Do not invent `kind` values, subdirectories, or frontmatter fields
  without updating this schema and notifying the user.
- Do not collapse `log.md` or rewrite past entries. It is an audit
  trail; only append.
- Do not run `git push` automatically. The user controls what leaves
  this machine.
- Do not reference files outside `wiki/` in wiki pages. The wiki is
  self-contained; external references belong in source pages, not as
  wikilinks.

## 10. Open questions for the schema itself

These are deliberate gaps that will be filled as the wiki grows:

- Search over the wiki once `index.md` outgrows its context budget.
- Ingestion of long sources (books, large repos) that don't fit a
  single read.
- Per-source confidence weights when multiple sources contradict.
- A canonical workflow for ingesting structured operational logs (e.g.
  agent run traces, meeting transcripts, chat exports).

When you propose changes to this file, append them under a
`schema-change` log entry.
