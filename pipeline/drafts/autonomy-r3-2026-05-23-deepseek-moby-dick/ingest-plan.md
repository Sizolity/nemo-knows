---
kind: topic
sources: [raw/web/corpus-2026-05-18/102-moby-dick.md]
status: draft
---

# Ingest Plan

## Source Summary
- Plain‑text retrieval of *Moby‑Dick* from Project Gutenberg (file `2701-0.txt`) acquired via a supplemental `curl` after a TLS handshake failure on the landing page.
- The raw source contains the full front matter (title, Etymology, Extracts) and the narrative from Ishmael’s opening through Chapter 93 (The Castaway).
- The extraction is truncated at 900,000 characters. The final chunk breaks off mid‑sentence during Pip’s abandonment scene, marked `[truncated at 900000 characters]`. The remainder of the novel is not present; the source is incomplete.
- The structure moves from editorial apparatus into a single continuous heading block (`Moby‑Dick > Retrieved Text`) covering all narrative chunks.

## Candidate Wiki Pages
- wiki/sources/moby-dick-gutenberg-retrieval.md — Documents the acquisition history, TLS fallback, chunk structure, and the 900k‑character truncation point that leaves the novel incomplete.
- wiki/topics/ahab-in-moby-dick.md — Captain Ahab’s delayed entrance, physical symbolism (ivory leg, cruciform scar, pivot‑hole stance), monomania, pasteboard‑mask metaphysics, discarded pipe, and his suspected pact with Fedallah.
- wiki/topics/moby-dick-as-legend-and-symbol.md — The White Whale’s markings, ubiquity, alleged immortality, intelligent malignity, historical attacks (Essex, Town‑Ho, affidavit), and the meditation on whiteness as a terror beyond blood.
- wiki/topics/ishmaels-philosophical-cosmology.md — Recurring philosophical passages: the Lee Shore and the soul’s peril, whiteness and annihilation, mast‑head reverie, monkey‑rope interdependence, loom of time, Isolatoes, and the dignity of whaling.
- wiki/topics/whaling-economy-and-practices.md — The lay system, cutting‑in and spermaceti extraction, fast‑fish/loose‑fish legal satire, ambergris recovery, pitchpoling, cetological classification, whale‑line hazards, and the gam ritual.
- wiki/topics/quequeg-and-cross-cultural-friendship.md — Queequeg’s origin, tattooing, idol Yojo, harpoon skill, innate dignity, and his bond with Ishmael as a model of “civilized” friendship beyond religious and racial divides.
- wiki/topics/supernatural-prophecies-and-doom.md — Elijah’s riddles, Father Mapple’s sermon, Gabriel and the Jeroboam plague, Fedallah’s tiger‑yellow crew and the two‑heads charm, and Ahab’s Faustian undercurrent.

## Suggested Links
- none

## Review Checklist
- [ ] Source Summary accurately reports the 900k‑character truncation mid‑Chapter 93 and the resulting incompleteness of the text.
- [ ] No candidate page describes the source as complete, full, or containing all chapters.
- [ ] All candidate pages are immediate children of `wiki/sources/`, `wiki/concepts/`, or `wiki/topics/` and do not invent nested directories.
- [ ] The candidate set consolidates the group‑note hints into a modest number of broad pages, avoiding excessive fragmentation.
- [ ] Missing endings (Cetology cut‑off in group‑01, Town‑Ho’s story cut‑off, Stubb’s Supper cut‑off, Pip scene cut‑off) are noted where relevant in candidate descriptions and do not imply closure.
- [ ] No tool/API, index, log, AGENTS, or schema pages have been proposed.
