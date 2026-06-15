---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
The chunk is the initial YAML frontmatter and first headings of `raw/web/corpus-2026-05-18/102-moby-dick.md`. It sets up the Moby‑Dick source document, identifying it as corpus item 102 and documenting how the plain‑text edition was acquired.

## Local Summary
This chunk provides metadata for the Moby‑Dick source entry: corpus number 102, Project Gutenberg category, source and final URLs, retrieval date (2026‑05‑18), content‑type (`text/plain; charset=utf-8`), and fetch status. It explains that the primary fetch attempt on the ebook landing page failed due to TLS issues in urllib, so a supplemental curl fetch retrieved the plain‑text file directly.

## Key Claims
- Corpus item 102 is a long public‑domain narrative text (Moby‑Dick) from Project Gutenberg.
- The final fetched plain‑text URL is `https://www.gutenberg.org/files/2701/2701-0.txt`.
- Fetch status is “ok” only after a fallback: the landing page (`https://www.gutenberg.org/ebooks/2701`) failed TLS in urllib.
- Retrieval date is 2026‑05‑18; content‑type is `text/plain; charset=utf-8`.

## Entities And Concepts
- **Corpus item 102** – the Moby‑Dick entry.
- **Project Gutenberg** – category and source organization.
- **Moby‑Dick** – the public‑domain narrative text.
- **urllib** – Python HTTP client that experienced a TLS failure on the landing page.
- **Supplemental curl fetch** – fallback method used to obtain the plain‑text edition.
- **Final URL** – `https://www.gutenberg.org/files/2701/2701-0.txt`

## Procedures And API Details
- Primary attempt: urllib against the ebook landing page (`https://www.gutenberg.org/ebooks/2701`) — failed due to TLS.
- Fallback: supplemental curl fetch of the plain‑text file at `https://www.gutenberg.org/files/2701/2701-0.txt`.
- Fallback status recorded as “ok via supplemental curl fetch.”

## Nuance Or Contradictions
- The chunk’s own frontmatter lists sources (`raw/web/curated-web-corpus-2026-05-18.md` and the Gutenberg URL), indicating this file was generated from a curated corpus entry; the relationship is not detailed.
- No contradictions within the chunk itself.

## Candidate Wiki Hints
- **Corpus item 102 (Moby‑Dick)** – a page cataloguing this entry, its retrieval method, and source metadata.
- **Project Gutenberg plain‑text fallback via curl** – reusable note on fetching etexts when the standard landing page fails TLS.
- **Supplemental curl fetch pattern** – general technique for bypassing urllib TLS issues in corpus pipelines.
