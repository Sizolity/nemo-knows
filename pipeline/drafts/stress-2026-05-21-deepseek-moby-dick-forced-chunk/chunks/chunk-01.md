---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context

This chunk covers the YAML frontmatter of the raw source document (lines 1–8) and the subsequent “Fetch Metadata” section (lines 10–26) under the heading **Moby-Dick**. It is the first chunk of 17, establishing the document’s identity and acquisition details.

## Local Summary

The chunk defines the source document as a **source**‑kind item titled “Moby-Dick,” created on 2026‑05‑18. It declares the document’s own lineage via two listed sources: a local curated‑web corpus file and the Project Gutenberg ebook landing page URL. The “Fetch Metadata” section records the retrieval specifics: corpus item number 102, category “Project Gutenberg,” the source landing‑page URL, the final plain‑text URL (`/files/2701/2701-0.txt`), retrieval date, content type, and a fetch status of `ok` obtained through a supplemental `curl` fetch after a TLS failure with `urllib`.

## Key Claims

- The raw document is a `kind: source` item titled “Moby-Dick” and was created/updated on 2026‑05‑18.
- Its lineage includes two sources: the curated web corpus file `raw/web/curated-web-corpus-2026-05-18.md` and the Gutenberg URL `https://www.gutenberg.org/ebooks/2701`.
- The final content was retrieved from `https://www.gutenberg.org/files/2701/2701-0.txt` as `text/plain; charset=utf-8`.
- The retrieval succeeded only after a TLS failure in `urllib` prompted a supplemental `curl` fetch.
- The fetched content is described as “Long public‑domain narrative text.”

## Entities And Concepts

- **Moby-Dick** – the source document title.
- **Project Gutenberg** – category of the source item.
- **Corpus item 102** – internal identifier.
- **Source URL**: `https://www.gutenberg.org/ebooks/2701` (ebook landing page).
- **Final URL**: `https://www.gutenberg.org/files/2701/2701-0.txt` (plain‑text edition).
- **Supplemental curl fetch** – fallback method after a TLS‑related failure with `urllib`.
- **raw/web/curated-web-corpus-2026-05-18.md** – another source from which this document draws its own metadata.
- **urllib TLS failure** – the technical obstacle that triggered the fallback.
- Concepts: source‑document metadata structure, acquisition fallback pattern, curation of Project Gutenberg texts.

## Procedures And API Details

- The original acquisition attempt used Python’s `urllib` to fetch the ebook landing page but encountered a TLS error.
- As a supplement, a `curl` command was used to directly retrieve the plain‑text file from the Gutenberg `/files/` path.
- The resulting response had `Content-Type: text/plain; charset=utf-8` and was saved as the raw source content.

## Nuance Or Contradictions

- The chunk lists both a curated‑corpus file and the Gutenberg URL as sources; they are complementary, not conflicting.
- No contradictions are present; the need for a supplemental `curl` fetch is explicitly documented without any contradictory status.

## Candidate Wiki Hints

- A candidate procedural note: “Project Gutenberg Acquisition Fallback (urllib → curl)” describing the TLS failure and the plain‑text fetch pattern.
- A concept note: “Source Document Acquisition Metadata” that explains the fields (`source URL`, `final URL`, `fetch status`, etc.) and the supplemental‑fetch strategy.
- The source entity “Moby-Dick” itself, but only as a concrete example—no generic wiki page warranted from this metadata chunk alone.
