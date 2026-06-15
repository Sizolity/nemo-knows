---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

### Chunk Context
This is chunk 1 of 17 from `raw/web/corpus-2026-05-18/102-moby-dick.md`.
It covers lines 1–26, the YAML frontmatter and the “Moby-Dick > Fetch Metadata” section.
The chunk precedes any narrative content from the novel.

### Local Summary
The chunk records provenance metadata for the source document: corpus item 102, Project Gutenberg ebook 2701, source and final URLs, retrieval date (2026-05-18), content type (`text/plain; charset=utf-8`), and a fetch status.
The fetch succeeded via a supplemental curl request after the ebook landing page failed TLS when accessed with urllib.
A test value labels the content as “Long public-domain narrative text.”

### Key Claims
- The primary source is a Project Gutenberg plain-text edition of *Moby-Dick*, ebook 2701.
- The final URL used is `https://www.gutenberg.org/files/2701/2701-0.txt`.
- Retrieval occurred on 2026-05-18; content type is `text/plain; charset=utf-8`.
- Fetch status is “ok via supplemental curl fetch” because the landing page (`https://www.gutenberg.org/ebooks/2701`) failed TLS in urllib.
- The source is tagged `project-gutenberg` and `web-corpus` with confidence “medium.”

### Entities And Concepts
- **Moby-Dick** – Project Gutenberg etext 2701.
- **Project Gutenberg** – repository hosting the plain-text file.
- **Supplemental curl fetch** – fallback retrieval method after a TLS failure with urllib.
- **Corpus metadata** – item 102, category “Project Gutenberg”, test value “Long public-domain narrative text.”

### Procedures And API Details
- No API.
- The acquisition process:
  1. Attempt to fetch the landing page (`https://www.gutenberg.org/ebooks/2701`) via urllib – failed due to TLS.
  2. Supplemental curl fetch from the direct plain-text URL (`https://www.gutenberg.org/files/2701/2701-0.txt`) succeeded, yielding the source document.

### Nuance Or Contradictions
- The chunk is entirely metadata and contains no content from the novel.
- The fetch status documents a workaround for a TLS issue, which matters for reproducibility of the corpus build.

### Candidate Wiki Hints
- A page on **corpus acquisition methods** could describe the urllib‑TLS fallback pattern observed with Project Gutenberg items.
- A reference page for **Project Gutenberg source URLs** in the curated corpus could record the direct text‑path pattern (`/files/{ebook_id}/{ebook_id}-0.txt`).
