---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context

This is the first chunk (lines 1–26) of the raw source `raw/web/corpus-2026-05-18/102-moby-dick.md`. It covers the opening YAML metadata block, the document-level heading “Moby-Dick”, and the “Fetch Metadata” sub-section. The chunk introduces the source’s provenance and retrieval details.

## Local Summary

The chunk records the metadata that identifies the file as a source entry for Herman Melville’s *Moby-Dick* obtained from Project Gutenberg. The YAML frontmatter declares it with kind `source`, tags `project-gutenberg` and `web-corpus`, and a medium confidence. The “Fetch Metadata” section lists corpus item 102, the source URL `https://www.gutenberg.org/ebooks/2701`, the final fetched URL `https://www.gutenberg.org/files/2701/2701-0.txt`, retrieval date 2026-05-18, and a note that the plain-text edition was acquired via a supplemental curl fetch after the ebook landing page’s TLS handshake failed during an initial attempt with urllib.

## Key Claims

- The raw source represents a plain-text edition of *Moby-Dick* from Project Gutenberg.
- The initial fetch using urllib to the ebook landing page failed due to a TLS error; the text was subsequently obtained from the stable `/files/` path via curl.
- The file was curated as corpus item 102 with a medium confidence label.

## Entities And Concepts

- **Moby-Dick** – the public-domain novel by Herman Melville.
- **Project Gutenberg** – the digital library hosting the text.
- **Corpus item 102** – identifier within the `web-corpus` set.
- **Supplemental fetch** – a fallback retrieval method (curl) employed when the primary URL failed TLS validation.

## Procedures And API Details

- **Fetch procedure**:
  1. Attempt to fetch `https://www.gutenberg.org/ebooks/2701` with Python’s urllib.
  2. Encounter TLS failure (likely due to cipher mismatch or missing certificate).
  3. Fall back to a supplemental curl request targeting `https://www.gutenberg.org/files/2701/2701-0.txt`.
  4. Successful retrieval yields `text/plain; charset=utf-8` content.

## Nuance Or Contradictions

The chunk ends cleanly without truncation; there are no contradictions in the metadata.

## Candidate Wiki Hints

- A page on “Web Corpus Source Acquisition” could document the pattern of supplemental fetches for Project Gutenberg sources.
- The raw document itself might warrant a meta-page “Corpus Item 102: Moby-Dick” that links to this source record and any downstream notes.
- A concept note for “Project Gutenberg TLS fetch issues” might be useful if this pattern recurs across the corpus.
