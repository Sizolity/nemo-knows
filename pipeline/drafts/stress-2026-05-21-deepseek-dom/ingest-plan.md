---
kind: topic
sources: [raw/web/corpus-2026-05-18/060-dom-standard.md]
status: draft
---

# Ingest Plan

## Source Summary
- This is a complete snapshot of the DOM Standard, covering the core document object model, event handling, aborting, ranges, traversal, sets, and legacy XPath/XSLT.
- It defines the fundamental tree model, node hierarchy, shadow DOM with slot assignment, live mutation algorithms, and the cooperative cancellation pattern via `AbortController`/`AbortSignal`.
- The final chapters aggregate an exhaustive list of historically removed features and a large set of browser‑compatibility tables derived from MDN.

## Candidate Wiki Pages
- wiki/sources/dom-standard-snapshot.md — Source metadata page describing the provenance, scope, and snapshot date of the raw specification.
- wiki/concepts/dom-node-tree.md — Core data model: trees, node kinds, document tree, shadow tree, slots, slottables, and flattened tree order.
- wiki/concepts/dom-event-model.md — Event flow, `Event`/`EventTarget` IDL, dispatching algorithm, legacy window event, and integration with aborting.
- wiki/concepts/dom-abort-signal.md — `AbortController`/`AbortSignal` cooperative cancellation, dependent signals, static methods, and API integration pattern.
- wiki/concepts/dom-ranges.md — Live (`Range`) and static (`StaticRange`) range creation, boundary manipulation, delete/extract/clone operations, and consistency guarantees.
- wiki/concepts/dom-mutation-algorithms.md — Low‑level insert/move/replace/remove primitives, interaction with custom elements and `MutationObserver`, and range pre‑remove steps.
- wiki/topics/dom-legacy-and-removed-apis.md — Historically removed interfaces and members (e.g., `DOMConfiguration`, `MutationEvent`, `Entity`), plus purpose‑preserved but legacy‑classed APIs such as XPath and XSLT.
- wiki/topics/dom-compatibility-snapshot.md — At‑a‑glance engine‑support matrix for core DOM interfaces, drawn from the standard’s bundled MDN‑sourced compatibility tables.

## Suggested Links
none

## Review Checklist
- [ ] Verify that candidate pages cover the main areas identified in the group notes (tree, events, aborting, ranges, mutation observers, token lists, legacy features, compatibility data).
- [ ] Confirm that each candidate page is placed under wiki/sources/, wiki/concepts/, or wiki/topics/ and does not introduce nested directories.
- [ ] Cross‑check important local details (e.g., `AbortSignal` garbage‑collection condition, `move` primitive behaviour, `DOMTokenList` update‑step suppression) and ensure they are referenced in the appropriate concept page descriptions.
- [ ] Note that XPath/XSLT sections are incomplete and only provide Web IDL; the legacy topic page should clearly flag this.
- [ ] Ensure the source page records the snapshot date (2026‑05‑18) and warns that later changes are not reflected.
- [ ] Validate that the ingest plan does not propose index, log, AGENTS, or schema files.
