---
kind: topic
sources: [raw/web/corpus-2026-05-18/102-moby-dick.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source is a complete ingestion of Herman Melville's *Moby-Dick* (Project Gutenberg eBook ID 2701), split into 53 chunks totaling approximately 16,000 characters.
- The text covers the full narrative arc from Ishmael's introduction in New Bedford and his friendship with Queequeg, through the philosophical digressions on Cetology and whaling history, to the fatal chase of the White Whale and the sinking of the *Pequod*.
- Initial acquisition involved overcoming a TLS failure on the Project Gutenberg landing page by falling back to a direct fetch of the raw text file.

## Candidate Wiki Pages
- wiki/sources/project-gutenberg-moby-dick.md — Documenting the technical handling of TLS failures when fetching public domain texts from Project Gutenberg and the specific corpus ID structure.
- wiki/concepts/whaling-economy-and-law.md — Exploring the social hierarchy, share-based profit systems ("lays"), and the legal concepts of "Fast-Fish" versus "Loose-Fish" found in the text.
- wiki/topics/cetology-classification-system.md — A summary of Ishmael's rejection of Linnaean taxonomy in favor of a size-based classification system (Folio/Octavo/Duodecimo) and his definition of whales as spouting fish.
- wiki/concepts/whaling-safety-and-tools.md — Detailing specific whaling tools like the "Drugg," "Waif," and "Monkey-Rope," along with safety protocols for line coiling and harpooning mechanics.
- wiki/topics/ahab-monomania-and-prophecy.md — An entry covering Ahab's psychological state, his ivory leg prosthesis, the prophecy of Elijah, and the ritualistic oath binding the crew to his vengeance.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that all candidate pages are placed strictly under `wiki/sources/`, `wiki/concepts/`, or `wiki/topics/`.
- [ ] Ensure no nested directories are created within the wiki structure.
- [ ] Confirm that "Project Gutenberg Acquisition" is documented as a source technicality rather than a literary concept.
