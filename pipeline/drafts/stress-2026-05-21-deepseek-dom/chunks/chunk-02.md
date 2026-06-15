---
title: Chunk 02 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/060-dom-standard.md`, lines 1580–2900.
- Heading path: `3. Aborting ongoing activities > 3.2. Interface AbortSignal > 3.2.1. Garbage collection` through `4.3.3. Interface MutationRecord`.
- Covers: AbortSignal GC behaviour, requirements for using AbortSignal in APIs, node tree model, shadow trees (slots, slottables, assignment, signalling), mutation algorithms (insert, move, replace, remove), ParentNode/ChildNode/other mixins, NodeList/HTMLCollection, and the MutationObserver system.

## Local Summary
This chunk defines how an `AbortSignal` must not be garbage‑collected while it depends on source signals and still has listeners or abort algorithms. It then sets out the rules for any promise‑based web API to accept an `AbortSignal` via a `signal` dictionary member, reject with the signal’s abort reason, use abort steps, and handle the already‑aborted case.
The bulk of the chunk details the DOM node tree. It introduces nodes, child constraints per node type, document trees, shadow trees (with host, light tree, and `connected` definition), slots (created by `slot` elements, named, assigned nodes), slottables (Element and Text, optionally with manual assignment), and algorithms for finding, assigning, and signalling slot changes.
A suite of low‑level mutation algorithms is defined: **pre‑insert validity checks**, **pre‑insert**, **insert** (including insertion steps, post‑connection steps, custom element callbacks, and shadow‑slot re‑assignment), **move** (with separate moving steps, `moveBefore`), **replace**, **replace all**, **pre‑remove**, and **remove** (with removing steps, `disconnectedCallback`, and transient registered observers).
Several mixins are described: `NonElementParentNode` (exposes `getElementById` on `Document` and `DocumentFragment`), `DocumentOrShadowRoot` (`customElementRegistry`), `ParentNode` (`children`, `firstElementChild`, `lastElementChild`, `childElementCount`, `prepend`, `append`, `replaceChildren`, `moveBefore`, `querySelector`, `querySelectorAll`), `NonDocumentTypeChildNode` (`previousElementSibling`, `nextElementSibling`), `ChildNode` (`before`, `after`, `replaceWith`, `remove`), and `Slottable` (`assignedSlot`).
Finally, it covers old‑style live collections (`NodeList`, `HTMLCollection`) and the MutationObserver system: per‑agent flags, pending observers, the notify algorithm, slot‑change events, `MutationObserver` interface (`observe`, `disconnect`, `takeRecords`), options dictionary validation, and `MutationRecord`.

## Key Claims
- A non‑aborted dependent `AbortSignal` must not be garbage‑collected while its source signals are non‑empty **and** it has registered `abort` event listeners or non‑empty abort algorithms.
- Any promise‑returning web API that supports aborting must:
  - Accept an `AbortSignal` through a `signal` dictionary member.
  - Reject the promise with the `AbortSignal` object’s **abort reason**.
  - Reject immediately if the signal is already aborted.
  - Use the **abort algorithms** mechanism without clashing with other observers.
- DOM documents are represented as **node trees**; node interfaces respect strict child constraints (Document, DocumentFragment, Element, CharacterData, etc.).
- A **shadow tree** is always attached to a host; a node is **connected** if its shadow‑including root is a document.
- Slots are created via HTML’s `slot` element, have a `name` (updated via attribute change steps), and maintain an **assigned nodes** list.
- **Slottables** (Element and Text) have an `assigned slot` and an optional **manual slot assignment** (weak reference).
- **Insertion steps** must not modify the node tree, create browsing contexts, fire events, or execute JavaScript; they may queue tasks. **Post‑connection steps** allow those actions atomically after a batch insert.
- A **move** primitive (e.g., `moveBefore`) is distinct from remove‑+‑insert; it does **not** invoke insertion/removing steps, but runs **moving steps** and a `connectedMoveCallback` on custom elements.
- The `remove` algorithm includes transient registered observers to retain `subtree` observation after a node is removed.
- `NodeList` is a collection of nodes (live by default); `HTMLCollection` is a legacy element collection with `namedItem`.
- Mutation observers rely on per‑agent **pending mutation observers** and a **microtask**; slot‑change events are fired from the same microtask.

## Entities And Concepts
- **AbortSignal** – dependent signal GC rule
- **AbortController / AbortSignal in APIs** – `signal` dictionary member, abort reason, abort steps
- **Node tree** – constraints per interface (Document, DocumentFragment, Element, etc.)
- **Document tree**, **document element**, **in a document tree** vs **in a document**
- **Shadow tree**, **shadow root**, **host**, **light tree**, **connected**
- **Slot** – name, attribute change steps, default slot, assigned nodes
- **Slottable** – Element/Text, name, assigned slot, manual slot assignment (weak reference)
- **Find a slot** / **find slottables** / **find flattened slottables**
- **Assign slottables for a slot**, **assign slottables for a tree**, **assign a slot**
- **Signal slots** (per agent) and **slot change** signalling
- **Insertion steps**, **post‑connection steps**, **children changed steps**
- **Moving steps** (separate from insertion/removing)
- **Removing steps**, **transient registered observer**
- **Pre‑insert validity**, **pre‑insert**, **insert**, **move**, **replace**, **replace all**, **pre‑remove**, **remove**
- **Pre‑insert validity** checks for Document constraints (DocumentFragment with multiple elements, doctype after element, etc.)
- `NonElementParentNode` – exposed on Document and DocumentFragment
- `DocumentOrShadowRoot` – `customElementRegistry`
- `ParentNode` – `children`, `firstElementChild`, `lastElementChild`, `childElementCount`, `prepend`, `append`, `replaceChildren`, `moveBefore`, `querySelector`, `querySelectorAll`
- `NonDocumentTypeChildNode` – `previousElementSibling`, `nextElementSibling`
- `ChildNode` – `before`, `after`, `replaceWith`, `remove`
- `Slottable` – `assignedSlot`
- **Collection** – live vs static, filter and root
- `NodeList` – length, `item`, iterable
- `HTMLCollection` – historical, `namedItem`, supported property names based on ID and `name` attribute
- **MutationObserver** – callback, node list (weak references), record queue
- **MutationObserverInit** – `childList`, `attributes`, `characterData`, `subtree`, `attributeOldValue`, `characterDataOldValue`, `attributeFilter`
- **Registered observer**, **transient registered observer**
- **Pending mutation observers** (per agent), **mutation observer microtask queued**
- **Notify mutation observers** – empty queues, invoke callbacks, fire `slotchange` events
- `MutationRecord` – type, target, addedNodes, removedNodes, previousSibling, nextSibling, attributeName, attributeNamespace, oldValue

## Procedures And API Details
- **AbortSignal GC**: The condition ties collection to source signals and presence of listeners/algorithms.
- **Example promise method** `doAmazingness(options)`:
  1. Get `global`.
  2. Create promise `p`.
  3. If `options["signal"]` exists:
     - If signal already aborted → reject `p` with `signal`’s abort reason, return `p`.
     - Add abort steps: stop amazing things; reject `p` with `signal`’s abort reason.
  4. In parallel: compute result, queue a global task to resolve `p`.
  5. Return `p`.
- **Node tree child constraints** (rules for each interface).
- **Length of a node**: 0 for DocumentType/Attr; CharacterData’s `data` length; number of children otherwise.
- **`connected`** = shadow‑including root is a document.
- **Slot name update**: attribute change steps; run assign slottables for a tree on the root.
- **Slottable name update**: attribute change steps; if slottable is assigned, run assign slottables for assigned slot; then run assign a slot.
- **Find a slot** (slottable, open flag): uses parent’s shadow root, checks mode for `"open"`, handles `"manual"` slot assignment.
- **Find slottables for a slot**: manual assignment vs named slot matching.
- **Find flattened slottables**: recursive, includes slottables from nested slots.
- **Assign slottables for a slot**: update assigned nodes; if changed → signal a slot change; update `assigned slot` on each slottable.
- **Assign a slot (slottable)**: find slot; if found, assign slottables for it.
- **Signal a slot change**: append slot to agent’s signal slots; queue a mutation observer microtask.
- **Insert algorithm** (detailed): adopts nodes, handles live ranges, shadow hosts, slot assignment, custom element callbacks (`connectedCallback`), registry scoping, suppression of observers, queues tree mutation record, collects nodes for `post‑connection` steps.
- **Move algorithm**: checks same shadow‑including root, host‑including ancestor, child validity; live range/NodeIterator pre‑remove; remove from old parent; slot re‑assignment; run moving steps (with `isSubtreeRoot`, `oldAncestor`); if custom and connected → `connectedMoveCallback`; queue mutation records.
- **Replace** child: uses nodes, respects document constraints, removes child with suppression, inserts new nodes, queues mutation record.
- **Replace all** with node/null: removes all children, inserts if non‑null; does **not** check node tree constraints — spec authors must use wisely.
- **Remove** node: live range/iterator steps; slot re‑assignment; runs removing steps on node and shadow‑including descendants; calls `disconnectedCallback` if custom and parent connected; appends transient registered observers up the ancestor chain for `subtree` observers.
- **`prepend`, `append`, `replaceChildren`** on `ParentNode`: convert nodes into a single node, use pre‑insert/append/replace‑all.
- **`moveBefore(node, child)`**: use the move primitive.
- **`querySelector` / `querySelectorAll`**: scope‑match a selectors string against the context node.
- **`ChildNode`**: `before`, `after`, `replaceWith` use `convert nodes into a node`, handle parent, viable sibling.
- **`Slottable.assignedSlot`**: calls find a slot with `open` = true.
- **`HTMLCollection.namedItem`**: supports ID and `name` attribute in HTML namespace.
- **`MutationObserver.observe(target, options)`**: validates options, deduplicates, sets up registered observer; stores weak reference to target.
- **`MutationObserver.disconnect()`**: removes all registered observers, empties record queue.
- **`MutationObserver.takeRecords()`**: clones and empties the record queue.
- **Queue a mutation record**: walks inclusive ancestors, checks options (subtree, type filters, attributeFilter, oldValue), enqueues to interested observers, appends to pending observers, queues microtask.
- **Queue a tree mutation record**: specialization for `childList` with added/removed nodes.
- **Notify mutation observers microtask**: empties signal slots, fires `slotchange` events on each slot.

## Nuance Or Contradictions
- The note about `in a document` being deprecated because older specs have not been updated to account for shadow trees.
- `Attr` nodes “participate in a tree for historical reasons” but never have a parent or children.
- Insertion steps are deliberately restricted from running JavaScript; the example shows that `script` and `style` insertion still has script‑observable side effects through post‑connection steps.
- The `move` operation is a distinct primitive that **does not** invoke insertion or removing steps; it uses moving steps and `connectedMoveCallback` for custom elements — preserving state that would be lost by remove‑and‑insert.
- `replace all` does **not** check node tree constraints, placing the onus on specification authors.
- Transient registered observers exist solely to preserve `subtree` observation for a node after it is removed from its parent, so mutations within descendants are not lost.
- The `moveBefore` method on `ParentNode` uses the move primitive, while `before`/`after`/`replaceWith` on `ChildNode` use insert/replace (via `convert nodes into a node` and `pre‑insert`/`replace`).
- `slottable`’s `manual slot assignment` is implemented via a weak reference because it is not directly accessible from script.
- The algorithm for `assign slottables for a slot` compares the old and new assigned nodes lists; it only signals a slot change if they are not identical.

## Candidate Wiki Hints
- **AbortSignal garbage collection** – a dedicated note on the GC rule for dependent signals.
- **AbortSignal in web APIs** – the requirements and pattern for promise‑returning APIs that accept `AbortSignal`.
- **Node tree constraints** – a reference page listing the allowed children for each node interface.
- **Shadow DOM: slots and slottables** – the lifecycle of slot assignment, flattened slottables, and signalling.
- **Mutation algorithms (insert, move, remove)** – detailed breakdown of the insert, move, and remove primitives, their step restrictions, custom element callbacks, and live range handling.
- **ParentNode mixin** – the procedural `prepend`/`append`/`replaceChildren`/`moveBefore` methods and `convert nodes into a node`.
- **HTMLCollection vs NodeList** – live collections, `namedItem` and its supported property names.
- **MutationObserver lifecycle** – registration, transient observers, microtask notification, and `slotchange` event.
