---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- Chunk: 1 of 17, lines 1-26
- Document path: Document → Moby-Dick → Fetch Metadata
- Covers YAML frontmatter of the source file, the top-level heading, and the “Fetch Metadata” section.

## Local Summary
This chunk records acquisition metadata for the Project Gutenberg edition of *Moby-Dick*. The text was obtained via a supplemental curl fetch after the ebook landing page failed TLS with urllib.

## Key Claims
- Corpus item 102 is in the “Project Gutenberg” category.
- The source URL is `https://www.gutenberg.org/ebooks/2701`.
- The final plain-text URL is `https://www.gutenberg.org/files/2701/2701-0.txt`.
- The retrieval date is 2026-05-18.
- The content type is text/plain in UTF-8 encoding.
- The fetch status is “ok via supplemental curl fetch”.
- The test value describes the item as “Long public-domain narrative text.”

## Entities And Concepts
- **Moby-Dick**: Herman Melville’s novel, obtained from Project Gutenberg.
- **Corpus item 102**: A long plain-text narrative.
- **Project Gutenberg ebook 2701**: The specific edition used.
- **Supplemental curl fetch**: Method used to bypass a TLS failure in the primary urllib-based fetch.
- **test value**: A description used to verify content type/length, not a content preview.

## Procedures And API Details
- No programming interfaces described.
- The acquisition workflow involved an initial attempt via urllib that failed TLS, then a fallback using curl to retrieve the plain-text file directly.

## Nuance Or Contradictions
- The fetch succeeded only after the landing page failed TLS in urllib, which suggests a TLS compatibility issue with the ebook landing page but not the file server.

## Candidate Wiki Hints
- A page on “Project Gutenberg text retrieval” could capture the pattern of falling back to curl for plain-text files when ebook landing pages fail.
- A page on “Corpus acquisition supplemental fetch” for documenting similar TLS workarounds.
