---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source file: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- Lines: 1‑26 (chunk 1 of 17)
- Heading path: **Document** → **Moby-Dick** → **Fetch Metadata**
- The chunk covers the YAML frontmatter and the `Fetch Metadata` section at the top of the document.

## Local Summary
This initial chunk is administrative metadata for the Moby‑Dick source note. It declares the document’s kind (`source`), publication info (title, creation/update dates, source pointers), and a `Fetch Metadata` section that records how the Gutenberg plain‑text edition was acquired and its retrieval details.

## Key Claims
- The document is a corpus item with ID `102` in the Project Gutenberg category.
- The source URL is `https://www.gutenberg.org/ebooks/2701`, but the final resolved URL for the text is `https://www.gutenberg.org/files/2701/2701-0.txt`.
- The content was retrieved on **2026‑05‑18** with content type `text/plain; charset=utf-8`.
- Fetch status was **ok via supplemental curl fetch**; the standard `urllib`‑based fetch failed because the landing page (ebook/2701) encountered a TLS error.
- The test value describes the resource as a “long public‑domain narrative text.”

## Entities And Concepts
- **Moby‑Dick** – the novel (public‑domain, Project Gutenberg ebook #2701)
- **Corpus item 102** – identifier within the curated web corpus
- **Project Gutenberg** – category/source of the item
- **Source URLs**: landing page `https://www.gutenberg.org/ebooks/2701` and direct text file `https://www.gutenberg.org/files/2701/2701-0.txt`
- **Supplemental acquisition** – fallback curl fetch triggered by a TLS failure in urllib
- **YAML frontmatter metadata**: `title`, `kind`, `created`, `updated`, `sources`, `tags`, `confidence`

## Procedures And API Details
- **Fetching procedure**: initial fetch of the ebook landing page via `urllib` failed due to a TLS issue. A supplemental curl command was then used to retrieve the plain‑text edition directly from the `/files/2701/2701-0.txt` path. The fetch result was recorded as “ok” with the content type and retrieval timestamp.
- No explicit APIs or command‑line options are documented here, only the high‑level strategy.

## Nuance Or Contradictions
- The chunk itself is part of a note that will eventually contain the full novel text, but this header is only about source provenance. The actual narrative content will appear in later chunks.
- The “test value” line is ambiguous – it may be a manually entered descriptor rather than an automated extraction.
- The failure of `urllib` for the landing page suggests a potential TLS compatibility issue with that specific endpoint; the direct text URL did not exhibit the same problem.

## Candidate Wiki Hints
- A page on **Project Gutenberg corpus items** could collect observed patterns (ID, category, retrieval strategies).
- A page on **supplemental fetch methods for Gutenberg texts** could detail when to fall back to direct `/files/...` URLs after encountering TLS errors on the eBook landing pages.
- A page about **curated web corpus metadata** that explains the fields `corpus item`, `category`, `source URL`, `final URL`, and retrieval status.
