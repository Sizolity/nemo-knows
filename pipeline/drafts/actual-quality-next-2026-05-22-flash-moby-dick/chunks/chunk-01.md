---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source: raw/web/corpus-2026-05-18/102-moby-dick.md
- Chunk: 1 of 17, lines 1–26
- Heading path: Document > Moby-Dick > Fetch Metadata
- The chunk covers the YAML frontmatter, the top-level `# Moby-Dick` heading, and the `## Fetch Metadata` section listing acquisition details for the Project Gutenberg text.

## Local Summary
This chunk records metadata about the retrieval of Herman Melville’s *Moby-Dick* from Project Gutenberg. It identifies the corpus item number (102), source URLs, retrieval date, fetch status, and a test value confirming the document is a long public-domain narrative. A supplemental curl fetch was needed because the ebook landing page failed TLS when accessed via urllib.

## Key Claims
- The corpus item is “102” and belongs to the “Project Gutenberg” category.
- The source URL is `https://www.gutenberg.org/ebooks/2701`, but the final text was fetched from `https://www.gutenberg.org/files/2701/2701-0.txt`.
- The document was retrieved on 2026-05-18 with content type `text/plain; charset=utf-8`.
- Fetch status: “ok via supplemental curl fetch”, indicating the initial attempt (likely with urllib) failed due to a TLS issue on the landing page.
- The test value describes the content as “Long public-domain narrative text.”

## Entities And Concepts
- **Moby-Dick**: The novel by Herman Melville, used here as a long-form public-domain text.
- **Project Gutenberg**: The digital library from which the text was sourced.
- **Corpus item 102**: Internal identifier for this text within the local wiki’s web corpus.
- **Supplemental curl fetch**: A fallback retrieval method triggered when the primary fetch mechanism encounters a TLS error.
- **TLS failure on ebook landing page**: The cause of the initial fetch failure; the plain-text file URL worked without TLS issues.

## Procedures And API Details
- No API commands are shown in this chunk, but the acquisition note implies a two-step process:
  1. Attempt to access the Gutenberg landing page (`/ebooks/2701`) using urllib; this failed with a TLS error.
  2. Switch to a supplemental curl fetch targeting the direct plain-text URL (`/files/2701/2701-0.txt`), which succeeded.

## Nuance Or Contradictions
- The chunk does not end mid-sentence or with a truncation marker; it is a complete entry.
- No contradictions are present; the fetch status and supplemental acquisition note consistently describe a successful retrieval after a TLS problem.

## Candidate Wiki Hints
- The retrieval pattern (primary fetch failure, curl fallback) may be common for Project Gutenberg sources. A wiki page **“Project Gutenberg Fetch Issues”** could document typical TLS or redirect problems and solutions.
- The metadata fields (Corpus item, Category, Source URL, Final URL, Retrieved, Content-Type, Fetch status, Test value) suggest a standard template for source acquisition records. A page **“Corpus Fetch Metadata Standard”** might be useful for maintainers.
