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
