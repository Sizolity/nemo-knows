## group-01

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Group Context
- Source: `raw/web/corpus-2026-05-18/060-dom-standard.md` – the full DOM Standard snapshot.
- Chunks covered: 1–9 (lines 1–15051) – the entire document.
- Major heading coverage:
  - Infrastructure (trees, ordered sets, selectors, name validation)
  - Events (Event, CustomEvent, EventTarget, dispatching, firing)
  - Aborting (AbortController, AbortSignal, API integration patterns)
  - Nodes (node tree, shadow tree, slots, mutation algorithms, mixins, NodeList/HTMLCollection, MutationObserver)
  - Core interfaces (Node, Document, DocumentType, DocumentFragment, ShadowRoot, Element, CharacterData, Text, Comment, etc.)
  - Ranges (AbstractRange, StaticRange, Range, manipulation and comparison)
  - Traversal (NodeIterator, TreeWalker, NodeFilter)
  - Sets (DOMTokenList)
  - Legacy XPath / XSLT APIs
  - Security & privacy considerations
  - Historical removals, acknowledgments, index, references, IDL index

## Cross-Chunk Summary
This snapshot provides the complete normative definition of the DOM Standard. Early chunks establish the foundational data model: trees, ordered sets, selectors, and naming rules. The event model is then specified in detail – how events flow, are constructed, dispatched, and how they relate to aborting concepts. The `AbortController`/`AbortSignal` pattern is introduced as a cooperative cancellation mechanism, with explicit guidance for promise‑based APIs.

The core of the standard resides in the node hierarchy. Chunks 2‑4 define the node tree, including document trees, shadow trees with named/manual slot assignment, and the low‑level mutation algorithms (insert, move, replace, remove) that underpin all DOM operations. These algorithms handle live range adjustments, custom element callbacks, and `MutationObserver` signalling. The abstract `Node`, `Document`, and `ShadowRoot` interfaces are specified with their APIs, followed by `Element`, attribute handling, `NamedNodeMap`, `CharacterData`, and concrete leaf types.

Ranges, traversal, and token‑list APIs complete the set of core interfaces. The `Range` family provides both static and live‑range behaviour with detailed text‑manipulation primitives. `NodeIterator` and `TreeWalker` offer alternative ways to walk the DOM, while `DOMTokenList` wraps ordered‑set semantics for class‑like attributes. The final chunks preserve legacy XPath and XSLT interfaces (documented as incomplete), record the exhaustive list of historically removed features, and consolidate acknowledgments, terms, references, and the IDL index.

Across all chunks, the standard emphasizes live collections, strict mutation rules, shadow‑DOM encapsulation, and a clear separation between infrastructure and higher‑level APIs. No known security or privacy concerns are identified.

## Repeated Or Central Claims
- **Event dispatching is a path‑based, capture‑then‑bubble flow** – events do not represent actions. This model is introduced in chunk 1 and remains the foundation for all event‑related algorithms, including the `EventTarget` listener list, `composedPath()`, and legacy `window.event`.
- **The DOM is a tree of nodes with strict parent‑child contracts** – tree definitions, child constraints per interface, and traversal order are first established in chunks 1‑2 and then referenced throughout the rest of the document (especially in mutation algorithms, `compareDocumentPosition`, and traversal).
- **Mutation algorithms (insert/move/replace/remove) are low‑level primitives** – defined in chunk 2, they are reused by higher‑level methods (e.g., `ParentNode.append`, `ChildNode.before`) and integrate with live ranges, custom element callbacks, shadow DOM slot re‑assignment, and `MutationObserver`.
- **`AbortSignal` objects are intended for cooperative cancellation** – chunks 1‑2 detail the controller/signal API, dependent signals via `AbortSignal.any()`, static factory methods, and the binding pattern for promise‑returning web APIs. The same pattern is expected to be reused by any specification that wants to offer abort support.
- **Shadow trees provide encapsulation via slots and assigned nodes** – the slot/slottable machinery, signal‑and‑assign cycle, and the flattened tree model appear in chunks 2‑3 and are integral to custom element and event‑retargeting logic.
- **Live ranges and `MutationObserver` maintain consistency after tree changes** – the algorithms for range pre‑remove steps and mutation‑record queuing are interwoven with the mutation primitives (chunks 2‑4) and guarantee that scripts see a coherent tree even as nodes are moved or removed.
- **Name validation rules are intentionally loose** – chunk 1 relaxes historical XML‑era constraints to match what the HTML parser can produce, and this policy is respected by `Element.setAttribute` and `Document.createElement` (chunks 3‑4).

## Important Local Details
- **`AbortSignal` garbage‑collection condition** (chunk 2): a non‑aborted dependent `AbortSignal` must not be collected while its source signals are non‑empty and it has listeners or abort algorithms. This subtlety is critical for memory management.
- **`move` primitive vs. remove‑and‑insert** (chunk 2): `move` runs moving steps and `connectedMoveCallback` but explicitly avoids insertion/removing steps, preserving state (e.g., `<iframe>` content) that would be destroyed by a remove‑and‑insert round‑trip.
- **`replace all` does not check node‑tree constraints** (chunk 2): specification authors are warned to use it carefully, as it bypasses the safety‑valve validations present in `pre‑insert`.
- **`Serializable`/`clonable` shadow roots and declarative shadow DOM** (chunk 3): `ShadowRoot` can be marked `serializable` or `clonable`, enabling declarative shadow tree serialization and `cloneNode`‑preserving shadow trees, respectively.
- **`compareDocumentPosition` and attribute ordering** (chunk 3): attributes return `IMPLEMENTATION_SPECIFIC` combined with `PRECEDING`/`FOLLOWING` bits; exact ordering is implementation‑defined, often based on pointer comparison or a random fallback.
- **`StaticRange` validation does not auto‑adjust** (chunk 4): a static range becomes invalid when its boundaries cross documents or offsets exceed node lengths, with no automatic correction, unlike live `Range`.
- **`DOMTokenList` update‑step suppression** (chunk 5): for web compatibility, `toggle` and `replace` may skip running update steps, meaning the backing attribute is not always immediately synchronised.
- **Exhaustive list of removed interfaces and members** (chunks 6‑9): the historical section catalogs every dropped feature such as `DOMConfiguration`, `MutationEvent`, `Entity`, `schemaTypeInfo`, `getUserData`, `isElementContentWhitespace`, and others. This is valuable for migration and understanding deprecated APIs.

## Candidate Wiki Hints
- **DOM Tree Concepts** – trees, tree order, node types, child constraints, shadow‑including order.
- **DOM Event Model** – `Event`/`EventTarget` APIs, capture/bubble, dispatching algorithm, legacy `window.event`.
- **AbortController and AbortSignal** – cooperative cancellation, dependent signals, static methods, API integration pattern.
- **Node Tree Mutation Algorithms** – insert, move, replace, remove primitives and their integration with custom elements and ranges.
- **Shadow DOM and Slot Assignment** – shadow trees, host, slots, slottables, assignment cycle, signal‑slot‑change microtask.
- **Ranges** – live vs. static ranges, containment, extraction/deletion/cloning, boundary‑point comparisons.
- **MutationObserver** – registration, transient observers, notification microtask, `slotchange` events.
- **DOM Token Lists** – `DOMTokenList` as ordered set, validation, update‑step quirks.
- **Deprecated and Removed DOM Features** – historical removals, rationale, migration notes.
- **Legacy XPath and XSLT APIs** – presence but incomplete specification, placeholder status.

## Gaps Or Cautions
- Confidence is **medium** because the chunk notes are synthetic summaries of a snapshot; subtle algorithm details or edge cases may be omitted.
- XPath and XSLT sections are explicitly marked as incomplete; their definitions are not fully specified, and only the Web IDL is provided for compatibility.
- The document uses the snapshot date `2026-05-18`; any later changes to the DOM Standard are not reflected.
- Some local details (e.g., `EntityReference` legacy constants preserved on `Node` despite interface removal) may appear contradictory unless cross‑referenced with the full removal list.
- The notes do not capture every algorithmic step, error‑handling branch, or informal warning present in the original source; direct reference to the standard is recommended for normative use.

## group-02

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Group Context
This group covers chunks 7–9 (lines 8818–15051) of the DOM Standard source. All three chunks fall within the top-level heading **11. Historical**. Together they form the complete *Historical* chapter, which specifies legacy DOM interfaces, traversal, ranges, token lists, XPath, XSLT, and a comprehensive set of browser‑compatibility tables derived from MDN data.

Chunk 07 supplies normative IDL and algorithmic descriptions for a range of historical interfaces, along with compatibility tables for some of those same interfaces. Chunk 08 and Chunk 09 consist entirely of compatibility data—a detailed snapshot of browser support for a large set of DOM APIs, many of which are not algorithmically described in this section but are referenced here for legacy completeness.

## Cross-Chunk Summary
The three chunks can be read as a single reference document for “what was once specified and how it is implemented now.” Chunk 07 defines the behaviour of:

- `CDATASection`, `ProcessingInstruction`, `Comment`
- `AbstractRange`, `StaticRange`, `Range`
- `NodeIterator`, `TreeWalker`, `NodeFilter`
- `DOMTokenList`
- `XPathResult`, `XPathExpression`, `XPathEvaluator` (and related mixin)
- `XSLTProcessor`

Each interface appears with its WebIDL, attributes and methods, and (where appropriate) step‑by‑step algorithms. Interspersed with that specification are compatibility tables that often start with “In all current engines.” followed by version‑number rows for desktop and mobile browsers.

Chunks 08 and 09 extend this compatibility layer to a much wider set of interfaces: `Document`, `DocumentFragment`, `Element`, `DocumentType`, `Text`, `Event`, `EventTarget`, `HTMLCollection`, `HTMLSlotElement`, `MutationObserver`, `MutationRecord`, `NamedNodeMap`, `Node`, `NodeIterator`, `NodeList`, `ShadowRoot`, `StaticRange`, `TreeWalker`, `XMLDocument`, `XPathEvaluator`, `XPathExpression`, `XPathResult`, `XSLTProcessor`, and many of their individual members. Taken together, the Historical section is a dual‑purpose chapter: it preserves the normative text of legacy APIs and simultaneously serves as a compatibility‑data appendix for the whole DOM.

## Repeated Or Central Claims
- **Broad engine support:** The overwhelming majority of listed APIs carry the statement “In all current engines.”, meaning Firefox, Safari, Chrome, Edge (Chromium) and Opera all ship these features.
- **Version‑annotated support:** For every API member the tables record the first‑release version for each engine family (Firefox, Safari, Chrome, Opera, Edge, Edge Legacy, IE) and frequently their mobile counterparts and Node.js for `Event`/`EventTarget` items.
- **Legacy gaps:** Internet Explorer is noted as “IENone” for many APIs that are otherwise universally available (`Element/closest`, `Element/classList`, `Event/composed`, etc.), highlighting that IE is no longer considered a current engine.
- **Data provenance:** The compatibility tables are clearly derived from MDN’s Browser Compatibility Data; several entries include “✔MDN” markers and occasionally “?” for mobile browsers where data was unverified at the time of the source snapshot.

## Important Local Details
### From Chunk 07 (normative + tables)
- `Range` provides full boundary‑point manipulation: `setStart`, `setEnd`, and their `Before`/`After` variants; collapse; `selectNode`/`selectNodeContents`; deletion, extraction, cloning, and `surroundContents`. The constants `START_TO_START`, `START_TO_END`, `END_TO_END`, `END_TO_START` (0–3) are defined for `compareBoundaryPoints`. `detach()` is specified as doing nothing—a historical artifact.
- `NodeIterator` and `TreeWalker` use `NodeFilter` (a callback interface) and a `whatToShow` bitmask. Several constants (`SHOW_ENTITY_REFERENCE`, `SHOW_ENTITY`, `SHOW_NOTATION`) are labelled as legacy.
- `DOMTokenList` supports `add`, `remove`, `toggle` (with optional `force`), `replace`, and `supports`. The `replace` method replaces one token with another; `supports` checks if a token is valid per the linked attribute’s definition.
- `XPathResult` declares result type constants (`ANY_TYPE`, `NUMBER_TYPE`, …) and provides `iterateNext` and `snapshotItem` for iterative/snapshot results.
- `XSLTProcessor` imports a stylesheet and transforms a source node into either a fragment or a document.
- Compatibility tables within this chunk note that `AbortSignal.timeout()` and the direct creation of `AbstractRange` gained support later than core members.

### From Chunk 08 (compatibility tables only)
- Tables cover:
  - `Document`(`prepend`, `querySelector`, `querySelectorAll`, `replaceChildren`, constructor)
  - `DocumentFragment`(similar members + `getElementById`)
  - `Element`(attribute methods, `checkVisibility`, `classList`, `closest`, etc.)
  - `DocumentType`(`name`, `publicId`, `systemId`)
  - `Text`(`assignedSlot`)
  - `Event`(constructor, `bubbles`, `cancelable`, `composed`, `composedPath`, `currentTarget`, `defaultPrevented`, `eventPhase`, `isTrusted`, `preventDefault`, `stopImmediatePropagation`, `stopPropagation`, `target`, `timeStamp`, `type`)
  - `EventTarget`(constructor, `addEventListener`, `dispatchEvent`, `removeEventListener`)
  - `HTMLCollection`(`item`, `length`, `namedItem`)
  - `HTMLSlotElement`(`slotchange` event)
  - `MutationObserver`(constructor, `disconnect`, `observe`, `takeRecords`)
  - `MutationRecord`(all properties)
  - `NamedNodeMap`(named item getters/setters, `item`, `length`)
  - `Node`(extensive list: `appendChild`, `baseURI`, `childNodes`, `cloneNode`, `compareDocumentPosition`, `contains`, `firstChild`, `getRootNode`, `hasChildNodes`, `insertBefore`, `isConnected`, `isDefaultNamespace`, `isEqualNode`, `isSameNode`, `lastChild`, `lookupNamespaceURI`, `lookupPrefix`, `nextSibling`, `nodeName`, `nodeType`, `nodeValue`, `normalize`, `ownerDocument`, `parentElement`, `parentNode`, `previousSibling`, `removeChild`, `replaceChild`, `textContent`)
  - `NodeIterator`(`filter`, `nextNode`, `pointerBeforeReferenceNode`)
- `Node.isConnected` shows a rare precise Samsung Internet version (`6.0+`) instead of a generic “?”.
- `Event` and `EventTarget` members include a Node.js version row (e.g., `14.5.0+`, `15.0.0+`), unique among the majority of entries.

### From Chunk 09 (compatibility tables only, complementing Chunk 07)
- Additional compatibility data for interfaces already described algorithmically in Chunk 07:
  - `NodeIterator` (`previousNode`, `referenceNode`, `root`, `whatToShow`)
  - `Range` (constructor and all manipulation/query methods)
  - `TreeWalker` (`currentNode`, filter, directional walk methods, `root`, `whatToShow`)
  - `XPathEvaluator` (constructor), `XPathExpression` (`evaluate`), `XPathResult` (all result-access members)
  - `XSLTProcessor` (constructor, parameter management, transformation methods)
  - `StaticRange` (constructor and interface)
  - `Text` (constructor, `splitText`, `wholeText`)
- `ShadowRoot` and its members (`delegatesFocus`, `host`, `mode`, `slotAssignment`) appear under Historical despite belonging to modern shadow DOM, probably because the standard aggregates all compatibility data here.
- `Range.detach` is listed with a Firefox version range `1–15`, indicating it was removed after Firefox 15, while other entries simply give a version.

## Candidate Wiki Hints
- **“Historical DOM Interfaces”** – A landing page that links to all legacy but still‑specified interfaces: `CDATASection`, `ProcessingInstruction`, `Comment`, the Range family, traversal, token lists, XPath, XSLT, and compatibility considerations.
- **“DOM Range API”** – Detailed reference for `AbstractRange`, `StaticRange`, and `Range`, including boundary manipulation, content extraction/deletion/cloning, and the `detach()` legacy.
- **“NodeIterator & TreeWalker”** – DOM tree traversal with filtering using `NodeFilter` callbacks and `whatToShow` bitmasks.
- **“DOMTokenList”** – Managing space‑separated tokens with `add`, `remove`, `toggle`, `replace`, and `supports`.
- **“XPath in the DOM”** – Evaluating XPath expressions, handling result types (`XPathResult` constants), and using `XPathEvaluator` on documents.
- **“XSLTProcessor”** – Client‑side XSLT transformations with parameter management.
- **“DOM Compatibility Snapshot”** – A wiki page derived from the combined tables of chunks 08 and 09, providing an at‑a‑glance engine support matrix for core DOM interfaces, with notes on IE gaps, mobile “?” entries, and Node.js versions where available.
- **“Shadow DOM in the Historical Section”** – Note on how `ShadowRoot` compatibility data ended up here; a cross‑reference to the modern shadow DOM spec.

## Gaps Or Cautions
- The source material intermingles normative DOM Standard text with external MDN compatibility tables; the MDN data may not be part of the official living standard and could be a snapshot frozen in time.
- Numerous mobile‑browser entries carry “?”, indicating unverified or missing support information. Those blanks should not be taken as evidence of non‑support.
- The “Historical” designation for some interfaces (`CDATASection`, `Range.detach`, `SHOW_ENTITY_REFERENCE`) is clear, but the presence of `ShadowRoot` and `StaticRange` compatibility tables under the same section may confuse readers; in the standard’s source these are placed there simply because all compatibility data was aggregated into the Historical chapter.
- Several version numbers reflect the state of browsers at the time the corpus was captured; more recent releases are not reflected. Similarly, the Firefox version range for `Range.detach` (1–15) is a historical note that detach was removed, which may not be obvious from the table alone.
- No procedural steps appear in chunks 08 and 09; they are purely reference tables. Any algorithmic detail must be drawn from chunk 07.

