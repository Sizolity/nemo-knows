---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/060-dom-standard.md` (DOM Standard)
- Chunk 3 of 9, lines 2901–4462
- Heading path: `Text: Why yes.⏎ > 4.4. Interface Node` through `4.8. Interface ShadowRoot`
- Covers: `Node`, `Document`, `DOMImplementation`, `DocumentType`, `DocumentFragment`, `ShadowRoot` interfaces, plus associated algorithms, constants, and concepts.

## Local Summary
This chunk defines the fundamental `Node` interface and its descendants: `Document`, `DocumentType`, `DocumentFragment`, and `ShadowRoot`. It documents the DOM tree structure, node properties (type, name, relationships), mutation methods, cloning, comparison, namespace utilities, and document creation and manipulation procedures. The `DOMImplementation` factory interface and the shadow DOM (`ShadowRoot`) infrastructure are also detailed, including shadow‑including tree traversals and retargeting for event dispatch.

## Key Claims
- `Node` is an abstract interface; every node has a node document, a registered observer list, and a `get the parent` algorithm that may route through assigned slots.
- Node type constants (`ELEMENT_NODE`, `TEXT_NODE`, etc.) and `nodeName` values are defined per interface.
- `isConnected`, `ownerDocument`, `getRootNode()`, `parentNode`, `childNodes`, and sibling accessors reflect the live tree.
- `cloneNode()` triggers the “clone a node” algorithm, which may run cloning steps and optionally attach a shadow root.
- `compareDocumentPosition()` returns a bitmask; attribute ordering relative to children is special‑cased.
- `createElement` and `createElementNS` support `ElementCreationOptions` (custom element registry, `is`); `createDocument` returns an `XMLDocument` and sets content type based on namespace.
- `importNode()` and `adoptNode()` move or copy nodes across documents; `adopt` updates node documents and custom element registries, enqueuing `adoptedCallback` for custom elements.
- `ShadowRoot` is a `DocumentFragment` with a non‑null host, mode (`"open"`/`"closed"`), slot assignment (`"manual"`/`"named"`), and flags for focus, declarative, clonable, serializable, and custom element registry.
- Shadow‑including tree order and retargeting algorithm are defined for event path computation.

## Entities And Concepts
- **Interfaces**: `Node`, `Document`, `XMLDocument`, `DOMImplementation`, `DocumentType`, `DocumentFragment`, `ShadowRoot`.
- **Dictionaries**: `GetRootNodeOptions`, `ElementCreationOptions`, `ImportNodeOptions`.
- **Enums**: `ShadowRootMode` (`"open"`, `"closed"`), `SlotAssignmentMode` (`"manual"`, `"named"`).
- **Constants**: `ELEMENT_NODE` (1) through `NOTATION_NODE` (12); `DOCUMENT_POSITION_DISCONNECTED` (0x01), `PRECEDING` (0x02), `FOLLOWING` (0x04), `CONTAINS` (0x08), `CONTAINED_BY` (0x10), `IMPLEMENTATION_SPECIFIC` (0x20).
- **Concepts**: node document, shadow‑including root, shadow‑including tree order, inclusive descendant, connected node, cloning steps, adopting steps, global vs. scoped custom element registry, flatten element creation options, retargeting, closed‑shadow‑hidden.

## Procedures And API Details
- **`nodeType` getter**: dispatches on the interface implemented (Element → `ELEMENT_NODE`, etc.).
- **`nodeName` getter**: returns HTML‑uppercased qualified name for elements, `"#text"` for Text, etc.
- **`getRootNode(options)`**: returns root if `composed` false, else shadow‑including root.
- **`normalize()`**: merges contiguous exclusive `Text` nodes, removing empty ones and adjusting live range positions.
- **Clone a node**: deep‑clones with optional shadow root attachment; preserves element custom element registry or falls back to `fallbackRegistry`.
- **`compareDocumentPosition`**: attribute handling leads to special `IMPLEMENTATION_SPECIFIC` | `PRECEDING`/`FOLLOWING` results.
- **`createElement`**: if HTML document, lowercases local name; uses HTML namespace unless content type is `"application/xhtml+xml"`; supports `customElementRegistry` and `is` options.
- **`adoptNode`**: throws for documents and shadow roots; updates node documents and custom element registries for all shadow‑including descendants, then enqueues `adoptedCallback` for custom elements.
- **`DOMImplementation`**: `createDocument` returns `XMLDocument` with content type mapped from namespace; `createHTMLDocument` builds a minimal HTML document.
- **ShadowRoot properties**: `host`, `mode`, `delegatesFocus`, `slotAssignment`, `clonable`, `serializable`; `onslotchange` event handler.
- **Shadow‑including tree order**: preorder, depth‑first, diving into shadow trees after encountering a shadow host.

## Nuance Or Contradictions
- `isSameNode` is a legacy alias of `===` and should be avoided.
- `hasFeature()` always returns `true`; it was unreliable and is retained only for legacy compatibility.
- `createEvent()` is legacy; event constructors are preferred, but the method still maps string names to interfaces (e.g., `"htmlevents"` → `Event`).
- `importNode` and `adoptNode` treat scoped custom element registries differently: global registries are replaced with the document’s effective global registry; scoped registries are preserved.
- `compareDocumentPosition` for attributes returns `IMPLEMENTATION_SPECIFIC` combined with ordering bits; the exact ordering between attributes is implementation‑specific and may use pointer comparison or `Math.random()` in JavaScript implementations.
- `createCDATASection` throws `"NotSupportedError"` in HTML documents.

## Candidate Wiki Hints
- **Node (DOM)** – central abstract interface, type constants, tree traversal, cloning, comparison.
- **Document (DOM)** – document creation, mode (`no‑quirks`, `quirks`, `limited‑quirks`), factory methods, `createEvent` legacy.
- **ShadowRoot** – shadow DOM encapsulation, retargeting, slot assignment, declarative shadow roots.
- **DOMImplementation** – document type and document factory, `hasFeature` deprecation.
- **Tree Traversal & Position** – `compareDocumentPosition` bitmask, node relationships, shadow‑including order.
- **Custom Element Registry & Adoption** – registry scoping, `adoptedCallback` lifecycle, global vs. scoped registries.
- **Cloning Nodes** – `cloneNode`, shadow root clonable, `importNode` fallback registry.
