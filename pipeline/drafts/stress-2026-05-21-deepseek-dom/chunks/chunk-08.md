---
title: Chunk 08 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- **Source:** raw/web/corpus-2026-05-18/060-dom-standard.md
- **Chunk:** 8 of 9
- **Lines:** 11168–13614
- **Heading path:** 11. Historical

## Local Summary
This chunk lists browser‑support tables for a large set of DOM interfaces and their members. Each entry states “In all current engines.”, then gives first‑release version numbers (or “?”) for Firefox, Safari, Chrome, Opera, Edge, Edge (Legacy), IE, and various mobile browsers. No normative prose or algorithms appear—just raw compatibility data.

## Key Claims
- Many DOM APIs are supported “in all current engines,” meaning they ship in modern Firefox, Safari, Chrome, Edge, and Opera.
- Legacy support varies widely:
  - Some APIs were available since very early browser releases (e.g., `Node/appendChild` since Firefox 1, Safari 1.1, Chrome 1, IE 5).
  - Others arrived much later (e.g., `MutationObserver` appeared around Firefox 14, Chrome 26, IE 11).
  - A few, such as `Element/assignedSlot`, `Event/composed`, and `EventTarget` constructor, are absent from IE entirely or marked `IENone`.
- The tables distinguish between “Edge” (Chromium‑based) and “Edge (Legacy)” (EdgeHTML‑based).

## Entities And Concepts
The listed APIs cover these interface families:
- **Document** – `prepend`, `querySelector`, `querySelectorAll`, `replaceChildren`, constructor
- **DocumentFragment** – `prepend`, `querySelector`, `querySelectorAll`, `replaceChildren`, `getElementById`, constructor
- **Element** – `prepend`, `querySelector`, `querySelectorAll`, `replaceChildren`, attribute methods, `checkVisibility`, `classList`, `className`, `closest`, `getAttribute*`, `hasAttribute*`, `id`, `insertAdjacent*`, `localName`, `matches`, `namespaceURI`, `prefix`, `removeAttribute*`, `setAttribute*`, `shadowRoot`, `slot`, `tagName`, `toggleAttribute`, constructor
- **DocumentType** – `name`, `publicId`, `systemId`, constructor
- **Text** – `assignedSlot`
- **Event** – constructor, `bubbles`, `cancelable`, `composed`, `composedPath`, `currentTarget`, `defaultPrevented`, `eventPhase`, `isTrusted`, `preventDefault`, `stopImmediatePropagation`, `stopPropagation`, `target`, `timeStamp`, `type`
- **EventTarget** – constructor, `addEventListener`, `dispatchEvent`, `removeEventListener`
- **HTMLCollection** – `item`, `length`, `namedItem`
- **HTMLSlotElement** – `slotchange` event
- **MutationObserver** – constructor, `disconnect`, `observe`, `takeRecords`
- **MutationRecord** – `addedNodes`, `attributeName`, `attributeNamespace`, `nextSibling`, `oldValue`, `previousSibling`, `removedNodes`, `target`, `type`
- **NamedNodeMap** – `getNamedItem`, `getNamedItemNS`, `item`, `length`, `removeNamedItem`, `removeNamedItemNS`, `setNamedItem`, `setNamedItemNS`
- **Node** – `appendChild`, `baseURI`, `childNodes`, `cloneNode`, `compareDocumentPosition`, `contains`, `firstChild`, `getRootNode`, `hasChildNodes`, `insertBefore`, `isConnected`, `isDefaultNamespace`, `isEqualNode`, `isSameNode`, `lastChild`, `lookupNamespaceURI`, `lookupPrefix`, `nextSibling`, `nodeName`, `nodeType`, `nodeValue`, `normalize`, `ownerDocument`, `parentElement`, `parentNode`, `previousSibling`, `removeChild`, `replaceChild`, `textContent`
- **NodeIterator** – `filter`, `nextNode`, `pointerBeforeReferenceNode`

A global attribute `slot` is also listed.

## Procedures And API Details
- No procedural steps are present. The section is purely a data dump of version‑to‑version support.
- Each entry follows the same pattern: API path, “In all current engines.”, then rows for Firefox, Safari, Chrome, Opera, Edge, Edge (Legacy), IE, and the corresponding mobile variants, sometimes with a Node.js version noted for `Event`‑ and `EventTarget`‑related items.

## Nuance Or Contradictions
- Several API entries have “?” for mobile Firefox, Safari, Chrome for Android, Samsung Internet, or Opera Mobile, indicating unknown or unverified support data.
- Some APIs that are “in all current engines” nevertheless list `IENone` (e.g., `Element/closest`, `Element/classList`, `Event/composed`) because IE is considered either out of support or not a “current engine” for the purposes of that statement.
- `Node/isConnected` lists `Samsung Internet6.0+` instead of “?”, a rare precise mobile version.
- `Event` and `EventTarget` sub‑entries include a Node.js version row (e.g., `14.5.0+`, `15.0.0+`), while most other APIs do not.

## Candidate Wiki Hints
- A centralized **DOM Browser Compatibility** page could be built from this data, linking each API to its MDN page and providing a quick lookup of when each was first supported. The source chunk already organizes the data per‑interface, which would make a table‑based wiki page straightforward.
- No other distinct new topic emerges from this chunk alone, as the information is reference material rather than a new concept or procedure.
