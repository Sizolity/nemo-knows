---
kind: topic
sources: [raw/web/corpus-2026-05-18/102-moby-dick.md]
status: draft
---

# Ingest Plan

## Source Summary
- The raw document is a Project Gutenberg edition of *Moby‑Dick* fetched on 2026‑05‑18 via a curl fallback after TLS failure in the primary fetch. It contains entire novel text plus front matter (Etymology, Extracts) and full metadata.
- The novel follows Ishmael, a schoolmaster who goes to sea on a whaler, the *Pequod*, commanded by the monomaniacal Captain Ahab, who seeks revenge on the white whale Moby Dick that took his leg.
- Rich in digressive cetology, whaling procedures, philosophical reflections, and ship‑to‑ship encounters, the text explores themes of obsession, whiteness, fate, isolation, and the limits of knowledge.
- The source is literary and narrative; its value for this wiki lies in character studies, symbolic motifs, whaling practices, and the novel’s structural features.

## Candidate Wiki Pages
- `wiki/sources/moby-dick.md` — documents the raw file, its retrieval provenance, chunking, and overall structure (Etymology, Extracts, 135+ chapters).
- `wiki/topics/moby-dick-themes.md` — overarching themes: whiteness and annihilation, fate vs. free will, the lee shore, isolation, the limits of systematic knowledge, and the sublime terror of the whale.
- `wiki/topics/ahab-and-the-white-whale.md` — Ahab’s monomania, pasteboard mask philosophy, quarter‑deck oath, the doubloon, Moby Dick as symbol of hidden malice; includes Fedallah and prophetic elements.
- `wiki/topics/ishmael-and-quequeeg.md` — their friendship across cultural/religious divides, Queequeg’s nobility and background, Ishmael’s sea‑compulsion and narrative voice, the Spouter‑Inn and Whaleman’s Chapel.
- `wiki/topics/whaling-in-moby-dick.md` — whaling procedures (cutting‑in, pitchpoling, fast‑fish/loose‑fish, ambergris extraction), economic aspects (lay system), material culture (scrimshaw, whale‑line, druggs and waifs), cetological classifications, and critique of whale art.
- `wiki/topics/pequod-crew-and-encounters.md` — mates (Starbuck, Stubb, Flask), harpooneers (Queequeg, Tashtego, Daggoo), Pip, other ships (the *Goney*, *Town‑Ho*, *Jeroboam*, *Jungfrau*, *Rose‑Bud*), the gam custom, and the multi‑ethnic “Isolato” community.

## Suggested Links
- none

## Review Checklist
- [ ] Verify all character summaries, thematic interpretations, and whaling‑procedure details against the source text before building final pages.
- [ ] Confirm that the source page captures the fetch metadata (TLS failure, curl fallback, fetch date) and the document’s chunk structure.
- [ ] Ensure candidate pages avoid overlap—e.g., keep Ahab’s monomania separate from general themes, and crew bios consolidated in one page.
- [ ] Validate that all proposed slugs reside directly under `wiki/topics/` or `wiki/sources/` and no additional directories are introduced.
- [ ] Assess whether any highly distinctive concept (e.g., Fast‑Fish / Loose‑Fish) merits its own `wiki/concepts/` page without expanding the page set beyond a small, broadly useful collection.
