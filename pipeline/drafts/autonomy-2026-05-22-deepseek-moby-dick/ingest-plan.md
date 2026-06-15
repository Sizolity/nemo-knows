---
kind: topic
sources: [raw/web/corpus-2026-05-18/102-moby-dick.md]
status: draft
---

# Ingest Plan

## Source Summary
- The raw file is a plain-text retrieval of the early portion of *Moby-Dick* (to Chapter 93) from Project Gutenberg eBook 2701, obtained via a supplemental curl fetch after a TLS failure on the landing page.
- The source is **incomplete** due to multiple truncations: it breaks off mid-scene in Chapter 3 (chunk boundary), mid-argument in Chapter 32 (Cetology), mid-chapter in Chapter 42 (The Whiteness of the Whale) with a missing footnote, and finally with a hard `[truncated at 900000 characters]` marker that ends the file mid-sentence in Chapter 93. The novel’s later chapters and epilogue are absent.
- The covered material includes the novel’s prefatory apparatus (Etymology, Extracts), the first-person narrative from Ishmael’s arrival in New Bedford through the early part of the voyage, including Ahab’s Quarter‑Deck oath, numerous encyclopaedic digressions, and inset stories such as the Town‑Ho’s account.

## Candidate Wiki Pages
- wiki/sources/moby-dick-gutenberg-fetch.md — documents the retrieval metadata (curl fallback, final URL `https://www.gutenberg.org/files/2701/2701-0.txt`, content type) and the truncation boundaries afflicting this source; useful for provenance tracking.
- wiki/topics/moby-dick-central-themes-motifs.md — consolidates the major thematic threads present in the retrieved portion: Ahab’s monomania and the pasteboard‑mask philosophy, Moby Dick as symbol, the terror of whiteness, the lee‑shore metaphor, Ishmael’s “hyena” fatalism, the dignity of whaling, fate‑versus‑free‑will (the sword‑mat), and the use of the prefatory extracts as a panoramic overture.
- wiki/topics/queedqueg-and-ishmael-friendship.md — captures their meeting at the Spouter‑Inn, Ishmael’s evolving tolerance, Queequeg’s paradoxical nobility, their “bosom friend” bond, and the literal umbilical of the monkey‑rope; a key relational pillar of the novel.
- wiki/topics/the-crew-and-ship-life-of-the-pequod.md — covers the three mates (Starbuck, Stubb, Flask), the harpooneers (Queequeg, Tashtego, Daggoo), the hidden crew and Fedallah, Pip the ship‑keeper, the Isolatoes concept, the lay system, the Gam custom, and day‑to‑day whaling authority/hierarchy as seen in the retrieved chapters.
- wiki/topics/cetology-and-the-whale-in-moby-dick.md — addresses the novel’s extended digressions on whale taxonomy (the Folio/Octavo/Duodecimo system), sperm‑whale anatomy (head, spout, tail, eyes, blubber), the critique of pictorial depictions, and the dual identity of the whale as industrial commodity and cosmic mystery.
- wiki/topics/the-town-ho-story.md — the self‑contained inset tale of the *Town‑Ho* mutiny, Steelkilt and Radney, the white whale’s lethal intervention, and its status as a secret never revealed to Ahab; potent for studying narrative embedding and justice themes.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that the Source Summary explicitly notes the hard truncation at 900 000 characters and the earlier mid‑chapter truncations in Chapters 3 (chunk boundary), 32, and 42.
- [ ] Confirm that all candidate pages avoid describing the source as “complete,” “full text,” “entire novel,” “all chapters,” or any similar claim.
- [ ] Cross‑check the Candidate Wiki Pages against the group notes to ensure all major repeated motifs (white whale, Ahab’s mask, Queequeg, Isolatoes, cetology, trunked chapters) are covered without unnecessary duplication.
- [ ] Ensure that the wiki/topics/ pages are correctly classified as topic-type (literary/narrative source) and that no page is placed under an invalid directory.
