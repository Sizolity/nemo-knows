---
title: Chunk 09 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/060-dom-standard.md` (DOM Standard)
- Chunk: 9 of 9, lines 13616–15051
- Heading path: 11. Historical
- This chunk covers the entire “Historical” section of the standard and consists entirely of browser‑compatibility data for a large set of legacy DOM APIs.

## Local Summary
The chunk is a dense collection of compatibility tables (often identified by “✔MDN”), listing each API member and its support in current and legacy browser engines. No descriptive prose is present; the content is purely machine‑readable browser‑support information. Every entry follows the pattern: interface/member name, then a block of engine‑version lines (Firefox, Safari, Chrome, Opera, Edge, IE, and mobile equivalents).

## Key Claims
- The DOM Standard’s Historical section enumerates a wide range of interfaces and their members, and this chunk asserts their current implementation status.
- Nearly all listed historical APIs are marked “In all current engines”, implying broad availability despite their historical designation.
- A few features show later adoption (e.g., `ShadowRoot`, `StaticRange`, `Element/slot`) or platform limitations (e.g., IE missing `NodeIterator.referenceNode`, `NodeList.forEach`, `XPathEvaluator`, etc.).
- The chunk’s data is a verbatim extraction from the standard’s source, likely generated from the MDN Browser Compatibility Data format.

## Entities And Concepts
Prominent historical DOM interfaces whose compatibility data appear in this chunk:
- `NodeIterator` (and members: `previousNode`, `referenceNode`, `root`, `whatToShow`)
- `NodeList` (including `forEach`, `item`, `length`)
- `ProcessingInstruction` (`target`)
- `Range` (constructor and all manipulation/query methods: `cloneContents`, `cloneRange`, `collapse`, `commonAncestorContainer`, `compareBoundaryPoints`, `comparePoint`, `deleteContents`, `detach`, `extractContents`, `insertNode`, `intersectsNode`, `isPointInRange`, `selectNode`, `selectNodeContents`, `setEnd`, `setEndAfter`, `setEndBefore`, `setStart`, `setStartAfter`, `setStartBefore`, `surroundContents`, `toString`)
- `ShadowRoot` (`delegatesFocus`, `host`, `mode`, `slotAssignment`)
- `StaticRange` (constructor and interface)
- `Text` (constructor, `splitText`, `wholeText`)
- `TreeWalker` (`currentNode`, `filter`, `firstChild`, `lastChild`, `nextNode`, `nextSibling`, `parentNode`, `previousNode`, `previousSibling`, `root`, `whatToShow`)
- `XMLDocument`
- `XPathEvaluator` (constructor)
- `XPathExpression` (`evaluate`)
- `XPathResult` (`booleanValue`, `invalidIteratorState`, `iterateNext`, `numberValue`, `resultType`, `singleNodeValue`, `snapshotItem`, `snapshotLength`, `stringValue`)
- `XSLTProcessor` (constructor, `clearParameters`, `getParameter`, `importStylesheet`, `removeParameter`, `reset`, `setParameter`, `transformToDocument`, `transformToFragment`)
- `Element/slot` (now in the `shadow DOM` section, but appears under Historical in this raw source)

## Procedures And API Details
- No procedural steps or algorithmic descriptions are present; the chunk only reports compatibility version numbers.

## Nuance Or Contradictions
- The chunk shows many “?” entries for mobile‑browser variants (e.g., Firefox for Android, Safari for iOS, etc.), indicating missing or uncertain data in the original source.
- Some legacy engines (IE, Edge Legacy) are explicitly noted as “None” for certain methods (e.g., `range.comparePoint`, `range.isPointInRange`) while other historical APIs were supported even in IE5/IE9.
- The `Range.detach` line has a Firefox version range “1–15” while other lines just state a number; this suggests that `detach` was later removed from Firefox.
- The `ShadowRoot` interface and its members appear in the Historical section despite being part of modern shadow DOM; this may reflect a structural decision in the standard where all compatibility data is aggregated under “Historical”.

## Candidate Wiki Hints
- A source‑backed page could aggregate the compatibility snapshot for DOM Historical APIs, especially focusing on the “all current engines” status and notable exceptions.
- The raw data is effectively a snapshot of the MDN Browser Compatibility Data for historical interfaces; a wiki page could link to the specific entries as a reference for legacy feature support.
