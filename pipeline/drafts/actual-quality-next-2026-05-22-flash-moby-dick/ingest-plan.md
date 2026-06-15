---
kind: topic
sources: [raw/web/corpus-2026-05-18/102-moby-dick.md]
status: draft
---

# Ingest Plan

## Source Summary
- The raw file is a partial retrieval of Herman Melville's *Moby-Dick* fetched from Project Gutenberg (`2701-0.txt`) on 2026‑05‑18 via a supplemental curl fallback after a TLS failure on the landing page.
- The text spans from the front matter (ETYMOLOGY, EXTRACTS, Table of Contents) through Chapter 93 ("The Castaway"), covering Ishmael's embarkation, Queequeg's friendship, the *Pequod*'s outfitting, Ahab's quarter‑deck oath, extensive cetological and philosophical digressions, multiple ship encounters, the Grand Armada, and the beginning of Pip's abandonment.
- **The source is truncated.** The final chunk ends mid‑sentence at the word `rescued` with the explicit marker `[truncated at 900000 characters]`. All narrative content beyond this point—including the climactic chase, the three‑day battle, and the novel's conclusion—is absent. The source does not contain the complete work.

## Candidate Wiki Pages
- **wiki/topics/moby-dick-characters.md** — Broad survey of major figures (Ishmael, Queequeg, Ahab, Starbuck, Stubb, Flask, Fedallah, Pip, Peleg, Bildad, Elijah, Father Mapple, Gabriel) drawn from all three note groups; characters are introduced, developed, or referenced across the full span of the retrieved text.
- **wiki/topics/ahab-vengeance-and-the-white-whale.md** — Ahab's monomania, the "pasteboard mask" philosophy, the quarter‑deck oath ritual (doubloon, crossed lances, murderous chalices), Fedallah's phantom crew, the sphinx‑head soliloquy, and Moby Dick as both physical quarry and cosmic antagonist (groups 1–3).
- **wiki/topics/whiteness-in-moby-dick.md** — The "Whiteness of the Whale" chapter's argument that whiteness evokes annihilation and the void; catalogue of natural, historical, and artistic associations; links to the novel's broader terror of the unknown (group 2).
- **wiki/topics/cetology-and-whale-anatomy-in-moby-dick.md** — Ishmael's ironical Folio/Octavo/Duodecimo taxonomy, the sperm‑whale's asserted supremacy, comparative head anatomy (case, junk, spermaceti), spout uncertainty, tail motions, whale vision and spinal phrenology, and the encyclopaedic impulse as failure (groups 1–3).
- **wiki/topics/pequod-encounters-and-whaling-customs.md** — The gam ritual, ship encounters (*Goney*/Albatross, *Jeroboam* and Gabriel, *Jungfrau*, *Bouton de Rose* and the ambergris trick), the Town‑Ho's inserted story, Fast‑Fish and Loose‑Fish law, scrimshaw, cutting‑in and processing, pitchpoling, drugg and waif, and the whale‑line as mortal metaphor (groups 1–3).
- **wiki/topics/pip-the-castaway.md** — Pip's background as a Connecticut ship‑keeper, his forced substitution into a whaleboat, two panic‑driven jumps, Stubb's calculated abandonment and the devaluation of his life ("a whale is worth thirty Pips in Alabama"), and his solitary drift on the open sea. **Note:** the source breaks mid‑rescue; Pip's complete fate and any later narrative role are not present in the raw text.
- **wiki/sources/moby-dick-fetched-text.md** — Acquisition record for corpus item 102: target ebook landing page TLS failure, supplemental curl fallback to the plain‑text file, fetch date, and the 900 000‑character truncation boundary that cuts off the novel in Chapter 93.

## Suggested Links
- none

## Review Checklist
- [ ] Truncation boundary (`[truncated at 900000 characters]`, mid‑sentence at `rescued`) is stated in Source Summary, candidate descriptions where relevant, and this checklist item.
- [ ] No candidate page describes the source as complete, entire, full text, all chapters, final chapters, or unabridged.
- [ ] All candidate pages are immediate children of `wiki/sources/`, `wiki/concepts/`, or `wiki/topics/`; no nested directories or invented directory names are present.
- [ ] Literary/narrative source classification is applied: topic pages are used for themes, motifs, characters, and practices; concept pages are proposed only if clearly warranted and rare.
- [ ] Group notes are treated as the authoritative whole‑document summary; no contradictory claims from per‑chunk notes are retained.
- [ ] Consolidated candidate set avoids duplication; repeated hints across groups are merged into the small set of broadly useful pages listed above.
- [ ] The `sources` frontmatter field matches the raw file path exactly.
