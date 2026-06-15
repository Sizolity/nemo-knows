---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---

Chunk Context
- **Chunk ID:** 1 of 10
- **Line Range:** 1-26
- **Heading Path:** Document > Alice's Adventures in Wonderland > Fetch Metadata
- **Source File:** raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md

Local Summary
This chunk establishes the metadata and acquisition context for the source document "Alice's Adventures in Wonderland" (Corpus item 104). It identifies the content as a Project Gutenberg plain-text edition retrieved from `https://www.gutenberg.org/files/11/11-0.txt`. The record notes a specific technical failure during initial retrieval where the ebook landing page failed TLS verification via `urllib`, necessitating a supplemental fetch of the direct file URL.

Key Claims
- The document is categorized under Project Gutenberg with tags including `project-gutenberg` and `web-corpus`.
- The content type is identified as `text/plain; charset=utf-8`.
- The retrieval status for the final file URL is reported as "ok via supplemental web fetch".
- The specific edition retrieved is described as a "Shorter narrative source."

Entities And Concepts
- **Alice's Adventures in Wonderland** (Title)
- **Project Gutenberg** (Source Category)
- **urllib** (Python Library referenced in error context)
- **TLS** (Transport Layer Security protocol involved in the fetch failure)
- **Corpus item 104** (Internal identifier)

Procedures And API Details
- **Initial Fetch Failure:** The landing page at `https://www.gutenberg.org/ebooks/11` failed TLS verification when accessed via `urllib`.
- **Resolution Procedure:** A supplemental acquisition was performed directly from the file URL: `https://www.gutenberg.org/files/11/11-0.txt`.

Nuance Or Contradictions
- The metadata describes the retrieved content as a "Shorter narrative source," which contrasts with the typical full-length nature of Project Gutenberg editions, though this may refer to a specific truncation or versioning within the corpus.

Candidate Wiki Hints
- **Acquisition Troubleshooting:** A note on handling TLS failures when fetching raw text files from legacy web archives using Python's `urllib`.
- **Project Gutenberg File Structure:** The distinction between the landing page (`/ebooks/<id>`) and the actual text file location (`/files/<id>/<id>-<version>.txt`).
