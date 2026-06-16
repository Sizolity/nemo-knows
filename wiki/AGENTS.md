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
  every page by category and is navigation-only: under each category
  heading it lists one entry per page (or `(none yet)`), with no
  explanatory prose. What each category holds is defined in this list,
  not repeated inside the index. Start here before any query, ingest, or
  lint pass.
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

**Links.** Every cross-reference between wiki pages — both the `index.md`
navigation catalogue and the semantic links inside body pages — is a standard
Markdown link whose target is a path relative to the page that contains it. A
sibling page in the same folder is `[Label](name.md)`; a page in another wiki
folder is `[Label](../folder/name.md)`; `index.md` links down into a category
as `[sqlite-wal](sources/sqlite-wal.md)`. Inline citations look like
`(see [Ada Lovelace](../entities/ada-lovelace.md))` or
`(source: wiki/sources/some-source.md §3)`. Obsidian-style `[[wikilinks]]` are
**not allowed** anywhere in `wiki/` content; the lint pass reports any residual
`[[...]]` as `forbidden-wikilink` (error). Because these links are real file
paths rather than slugs, they double as direct file jumps and need no
slug→path resolution layer in the renderer. The trade-off is that a relative
link breaks when its target page is renamed or moved; the lint pass guards this
by reporting any relative link whose target file does not exist as
`missing-link-target` (error), so a move must update its inbound links in the
same change. Slugs (page filenames) stay globally unique and are still used for
page identity and de-duplication — only the cross-reference syntax changed.

**Index format.** `index.md` is navigation-only. It contains the four
category headings — `## Sources`, `## Entities`, `## Concepts`,
`## Topics` — and, under each, one catalogue entry per page, or the literal
placeholder `(none yet)` when the category has no pages. Each entry is a
single line: a standard Markdown relative link followed by a one-line
description, e.g.
`- [sqlite-wal](sources/sqlite-wal.md) — Notes on SQLite write-ahead logging.`
Do not write per-section usage descriptions or any other explanatory prose
in the index; what each category holds is defined once in §0 (How the wiki
works), which is the single source for that guidance. Tooling generates and
normalizes the index from this format, so prose placed under a heading is
treated as noise and removed on the next maintenance pass. Adding the first
entry to a category replaces its `(none yet)` placeholder, and a category
that loses its last entry has the placeholder restored.

**Images and assets.** Binary assets — images, diagrams, slides — live under
`wiki/assets/`, grouped in a per-source or per-topic subdirectory, such as
`wiki/assets/<group>/<file>.png`. Pages reference them with standard Markdown
image syntax and a path relative to the page, so a page in `sources/`,
`entities/`, `concepts/`, or `topics/` writes
`![alt text](../assets/<group>/<file>.png)`. Always give the image meaningful
alt text. A local image target must resolve to an existing file under
`wiki/assets/`; the lint pass reports any missing local image as
`missing-image`. External images may be referenced by absolute `http(s)://` URL
and are not existence-checked. The web console serves `wiki/assets/` read-only
under `/assets/` and renders these references as inline `<img>` tags;
unresolvable or unsafe targets degrade to the alt text instead of an image.

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
2. **Find orphans.** List pages with no inbound semantic Markdown relative
   links and no `index.md` Markdown catalogue entry.
3. **Find stubs.** List relative links that point at missing target files
   (`missing-link-target`), and any residual `[[...]]` (`forbidden-wikilink`).
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
- Do not add per-section usage notes or other explanatory prose to
  `index.md`; it is navigation-only (see §2, Index format). Category
  meanings belong in §0 of this schema.
- Do not invent `kind` values, subdirectories, or frontmatter fields
  without updating this schema and notifying the user.
- Do not collapse `log.md` or rewrite past entries. It is an audit
  trail; only append.
- Do not run `git push` automatically. The user controls what leaves
  this machine.
- Do not reference files outside `wiki/` in wiki pages. The wiki is
  self-contained; external references belong in source pages, not as
  internal relative links.

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
