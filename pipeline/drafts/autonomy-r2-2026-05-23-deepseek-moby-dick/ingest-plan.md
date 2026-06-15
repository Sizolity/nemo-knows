---
kind: topic
sources: [raw/web/corpus-2026-05-18/102-moby-dick.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source is a raw Web fetch of *Moby‑Dick* from Project Gutenberg, retrieved via curl fallback after a urllib TLS failure; chunk 1 contains the fetch metadata.
- The supplied text includes the front matter (Etymology, Extracts) and Chapters 1–93, but it is incomplete: the file ends mid‑sentence during Pip’s rescue in Chapter 93 with an explicit `[truncated at 900000 characters]` marker, and earlier chapters are broken mid‑scene (the Whiteness meditation cuts off after the first paragraph; Stubb’s supper breaks off mid‑dialogue).
- The remainder of the novel—the final chapters, the Pequod’s fate, and the resolution—is absent; the source must not be described as complete, entire, or unabridged.

## Candidate Wiki Pages
- wiki/sources/moby-dick-raw.md — Document how the text was obtained, the chunking notes, and the known truncation boundaries.
- wiki/topics/ahab-and-moby-dick.md — Ahab’s obsession, the white whale’s lore, quarter‑deck oath, pasteboard‑mask philosophy, and the Sphynx soliloquy.
- wiki/topics/ishmael-and-queequeg.md — The narrator’s friendship with Queequeg, the breakdown of prejudice, the “wedding” bond, and their shared voyage.
- wiki/topics/moby-dick-themes-of-fate-and-the-sea.md — The lee shore, loom‑of‑time, monkey‑rope, whiteness as annihilation, fatalistic humour, and the sea‑soul motif.
- wiki/topics/cetology-and-whale-lore-in-moby-dick.md — Ishmael’s taxonomic parody, comparative anatomy, known whale behavior, and the claim that the living whale cannot be painted.
- wiki/topics/whaling-practice-and-law-in-moby-dick.md — Cutting‑in procedure, Fast‑Fish/Loose‑Fish, ambergris, scrimshaw, the gam, shipboard hierarchy, and pitchpoling.
- wiki/topics/prophecy-and-symbolism-in-moby-dick.md — Elijah, Gabriel and the *Jeroboam*, the doubloon, Fedallah and the secret crew, and prophetic signs.

## Suggested Links
- [[Moby‑Dick: Fetch and corpus metadata]]
- [[Ishmael’s narrative voice and philosophical digressions]]
- [[Queequeg: character, biography, and religious practices]]
- [[The Spouter‑Inn and its symbolism]]
- [[Father Mapple’s sermon on Jonah]]
- [[Fighting Quakers: Peleg, Bildad, and Nantucket whaling piety]]
- [[The Pequod’s officers and harpooneers (Starbuck, Stubb, Flask, Tashtego, Daggoo)]]
- [[Ahab before his appearance: the lost leg, the name, and early hints]]
- [[Elijah as prophetic figure]]
- [[Lee Shore and the soul’s rejection of land]]
- [[Ishmael’s defence of whaling (economic, exploratory, and regal arguments)]]
- [[Cetology in Moby‑Dick: the opening of the classification]]
- [[The white whale’s first mention and Ahab’s obsession]]
- [[Cetology in Moby‑Dick]]
- [[Shipboard hierarchy and Specksnyder]]
- [[Mast‑head as trope]]
- [[Ahab’s quarter‑deck ritual]]
- [[The Whiteness of the Whale]]
- [[Moby Dick (whale)]]
- [[Ahab’s charts and whale migration]]
- [[Historical whale attacks]]
- [[Fate, chance, and free will (Loom of Time)]]
- [[Ahab’s secret crew / Fedallah]]
- [[Gam (whaling social call)]]
- [[Town‑Ho mutiny and Moby Dick]]
- [[Critique of whale illustrations]]
- [[Scrimshaw and sailor art]]
- [[The whale‑line]]
- [[The great live squid]]
- [[Stubb and Fleece]]
- [[Cutting‑in (whaling procedure)]]
- [[Monkey‑rope as metaphor]]
- [[Gabriel and the *Jeroboam*]]
- [[Fedallah (the Parsee)]]
- [[Spermaceti extraction (baling)]]
- [[Comparative cetology]]
- [[Whale spout controversy]]
- [[Pitchpoling]]
- [[Honorary whalemen of myth]]
- [[Fast‑Fish / Loose‑Fish]]
- [[Ambergris]]
- [[Whale schools]]

## Review Checklist
- [ ] Source summary clearly states the text is truncated at 900,000 characters, ends mid‑sentence in Chapter 93, and that earlier sections break mid‑scene.
- [ ] Candidate page descriptions avoid any claim that the source is complete, entire, unabridged, or contains the full novel.
- [ ] All candidate pages are immediate children of `wiki/sources/` or `wiki/topics/`; no nested directories or invalid root sections are proposed.
- [ ] No candidate pages for `wiki/index.md`, `wiki/log.md`, `AGENTS.md`, or schema files are included.
- [ ] The metadata page (`wiki/sources/moby-dick-raw.md`) accurately reflects the fetch method, chunking, and the truncation marker.
- [ ] Topic pages are scoped to themes, motifs, or practices, not thin chapter summaries, and do not rely on content beyond the source boundary.
- [ ] Repeated candidate hints from the group notes have been consolidated into a small, broadly useful set of pages.
