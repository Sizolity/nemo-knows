---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
This chunk represents the metadata acquisition record for "Moby-Dick" (Corpus item 102). It documents the technical retrieval process from Project Gutenberg, noting a TLS failure on the landing page which necessitated a fallback to the raw text file. The content type is identified as UTF-8 plain text.

## Local Summary
The source is a Project Gutenberg ebook of "Moby-Dick," retrieved successfully via a supplemental curl fetch after an initial urllib attempt failed due to TLS issues. The document is stored as plain text with a confidence rating of medium.

## Key Claims
- The corpus item ID for Moby-Dick is 102.
- The source URL for the ebook landing page was https://www.gutenberg.org/ebooks/2701.
- The final content retrieved is located at https://www.gutenberg.org/files/2701/2701-0.txt.
- The acquisition occurred on 2026-05-18.
- A TLS error occurred with the landing page, requiring a direct fetch of the text file.

## Entities And Concepts
- **Moby-Dick**: The subject work.
- **Project Gutenberg**: The hosting organization.
- **TLS**: Transport Layer Security (context: connection failure).
- **urllib**: Python library used for initial retrieval attempt.
- **curl**: Command-line tool used for supplemental acquisition.

## Procedures And API Details
- **Fetch Logic**: Initial attempt using `urllib` failed; fallback to direct fetch of the text file using `curl`.
- **Content Type**: `text/plain; charset=utf-8`.
- **File Path**: `2701-0.txt`.

## Nuance Or Contradictions
The metadata indicates a successful retrieval ("Fetch status: ok") despite an initial "failed TLS" error on the landing page, implying the fallback mechanism worked correctly. The confidence level is explicitly set to medium.

## Candidate Wiki Hints
- **Project Gutenberg Acquisition**: A note on handling TLS failures when fetching public domain texts from Project Gutenberg using Python `urllib` versus `curl`.
