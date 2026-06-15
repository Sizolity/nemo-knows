---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source file: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- Chunk 1 of 17, lines 1–26
- Heading path: Document > Moby-Dick > Fetch Metadata
- The chunk starts with YAML frontmatter (title, kind, created/updated, sources, tags) then the document title and a metadata section.

## Local Summary
Introduces the Moby-Dick source note with its own metadata, then describes how the raw text was fetched. The initial fetch via the Project Gutenberg ebook landing page (`https://www.gutenberg.org/ebooks/2701`) encountered a TLS error with urllib, so the system performed a supplemental curl fetch of the plain-text file at `https://www.gutenberg.org/files/2701/2701-0.txt`.

## Key Claims
- The Moby-Dick item is part of a web corpus (corpus item 102, category Project Gutenberg).
- The primary URL (`https://www.gutenberg.org/ebooks/2701`) was tried but failed TLS in urllib.
- A supplemental curl fetch succeeded using the direct plain-text URL `https://www.gutenberg.org/files/2701/2701-0.txt`.
- The retrieved content type was `text/plain; charset=utf-8`.
- The narrative text is described as a “Long public-domain narrative text” (test value).
- Fetch status is recorded as “ok via supplemental curl fetch”.

## Entities And Concepts
- **Moby-Dick**: Source note about the Project Gutenberg edition.
- **Corpus item 102**: Identifier within the curated-web-corpus-2026-05-18.
- **Project Gutenberg**: Source category; ebook ID 2701.
- **Supplemental curl fetch**: Fallback retrieval method after urllib TLS failure.
- **Plain-text URL**: `https://www.gutenberg.org/files/2701/2701-0.txt` (direct text).
- **Ebook landing page**: `https://www.gutenberg.org/ebooks/2701` (primary but failed).

## Procedures And API Details
1. **Primary fetch attempt**: Use urllib to retrieve `https://www.gutenberg.org/ebooks/2701`.
2. **TLS failure**: urllib could not establish a secure connection (TLS error).
3. **Supplemental acquisition**: Execute a curl command to fetch `https://www.gutenberg.org/files/2701/2701-0.txt`.
4. **Result**: Got `text/plain; charset=utf-8` content, mark fetch status “ok via supplemental curl fetch”.

## Nuance Or Contradictions
- No truncation markers; the chunk ends at a complete bullet point.
- The raw source does not detail the exact urllib error or why curl succeeded; only that a TLS failure occurred.

## Candidate Wiki Hints
- **Project Gutenberg**: A reusable source for public-domain texts.
- **Corpus ingestion fallback strategies**: A concept page describing when and how supplemental curl fetches are triggered after urllib failures.
- **`raw/web/corpus-2026-05-18/102-moby-dick.md`**: The specific source note could be referenced if discussing how Moby-Dick was acquired.
