---
title: Chunk 05 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
Lines 5925–6887 of `raw/web/corpus-2026-05-18/060-dom-standard.md`.
Headings covered:
- 5. Switch on the position of thisPoint relative to sourcePoint (end of Range section)
- 6. Traversal (NodeIterator, TreeWalker, NodeFilter)
- 7. Sets (DOMTokenList)
- 8. XPath (XPathResult, XPathExpression, XPathEvaluatorBase, XPathEvaluator)
- 9. XSLT (XSLTProcessor)
- 10. Security and privacy considerations

## Local Summary
This chunk completes the `Range` methods (`deleteContents`, `extractContents`, `cloneContents`, `insertNode`, `surroundContents`, `cloneRange`, `detach`), specifies point/range comparison (`comparePoint`, `isPointInRange`, `intersectsNode`), and `Range` stringification. It then defines DOM traversal via `NodeIterator` and `TreeWalker` (including `NodeFilter` callbacks and `whatToShow` bitmasks). The `DOMTokenList` interface and its backing token set are described, followed by minimal preservation of XPath and XSLT APIs with notes that full definitions are missing. The chunk closes with a statement that no known security or privacy issues exist for the DOM Standard.

## Key Claims
- `deleteContents()`, `extractContents()`, and `cloneContents()` handle partially contained nodes by splitting, cloning, and removing as appropriate; all three share a common ancestor determination and partial containment logic.
- `insertNode(node)` splits a Text start node if necessary, removes `node` from any previous parent, and inserts it before a computed reference point, adjusting the range’s end if collapsed.
- `surroundContents(newParent)` throws if a non‑Text node is partially contained; it extracts the range, clears `newParent`’s children, inserts `newParent`, and appends the extracted fragment.
- `detach()` is preserved for compatibility but performs no operation.
- `comparePoint(node, offset)` returns −1, 0, or 1 based on the point’s position relative to the range; it throws for wrong document or doctype.
- `intersectsNode(node)` returns true if any part of the node is inside the range.
- Traversal objects (`NodeIterator`, `TreeWalker`) share an `is active` flag to prevent recursive filter invocation.
- `NodeIterator` has a reference node and a pointer‑before‑reference boolean; its pre‑remove steps adjust the reference when an ancestor of it is removed.
- `TreeWalker` provides tree‑navigation methods like `parentNode()`, `firstChild()`, `nextSibling()`, `previousNode()`, and `nextNode()` that respect filtering.
- `NodeFilter` defines constants for `acceptNode` return values (`FILTER_ACCEPT`, `FILTER_REJECT`, `FILTER_SKIP`) and `whatToShow` bitmask constants (e.g., `SHOW_ELEMENT`, `SHOW_TEXT`).
- `DOMTokenList` is backed by a token set; its update steps may not always run for `toggle()` and `replace()` for web compatibility.
- XPath and XSLT interfaces are retained only for Web IDL compatibility; complete definitions are not provided.
- The DOM Standard states there are no known security or privacy considerations.

## Entities And Concepts
- **Range**: methods for deletion, extraction, cloning, insertion, and point comparison; comparison helper *(before/equal/after)*.
- **NodeIterator**: root, reference, pointerBeforeReference, whatToShow, filter; `nextNode()`, `previousNode()`, pre‑remove steps.
- **TreeWalker**: root, whatToShow, filter, currentNode; parent/child/sibling/previous/next navigation.
- **NodeFilter**: callback interface with constants for filtering and whatToShow.
- **DOMTokenList**: token set, element, attribute name; validation, update, serialize steps; methods `add`, `remove`, `toggle`, `replace`, `supports`.
- **XPathResult**, **XPathExpression**, **XPathEvaluatorBase**, **XPathEvaluator**: legacy XPath 1.0 evaluation API.
- **XSLTProcessor**: legacy XSLT transformation API (`importStylesheet`, `transformToFragment`, `transformToDocument`, etc.).
- **Security**: no known concerns.

## Procedures And API Details
- **deleteContents()**: if collapsed → return; if same CharacterData node → replace data with empty string; otherwise collect `nodesToRemove` (all contained nodes omitting those whose parent is also contained), adjust start/end to a new node/offset, then remove the collected nodes and trim CharacterData boundaries.
- **extract a live range**: builds a `DocumentFragment` of the range’s contents, removes the original nodes, and returns the fragment; uses `firstPartiallyContainedChild`/`lastPartiallyContainedChild` logic.
- **clone the contents**: similar to extract but clones instead of removing; uses `cloneNode(subtree true)` for fully contained children.
- **insert node into range**: throws if start node is a `ProcessingInstruction`, `Comment`, or orphan `Text`; splits text node if start node is `Text`; ensures pre‑insert validity; removes node if already in tree; computes new offset and inserts.
- **surroundContents(newParent)**: throws if a non‑Text node is partially contained; extracts range, clears `newParent`, inserts `newParent` at range start, appends fragment to `newParent`, then selects `newParent`.
- **comparePoint(node, offset)**: returns −1 (before start), 1 (after end), 0 (in range) using internal position checks.
- **intersectsNode(node)**: returns true if node’s parent (or the node itself if root) lies such that (parent, offset) < end and (parent, offset+1) > start.
- **Range stringification**: concatenates text data from start Text node (partial), fully contained Text nodes, and end Text node (partial).
- **NodeIterator traversal**: toggles `beforeNode` flag and moves to next/previous node in the collection, filtering until `FILTER_ACCEPT`.
- **TreeWalker.children traversal**: first/last child, with `FILTER_SKIP` causing descent into children.
- **TreeWalker.siblings traversal**: navigating next/previous siblings, managing `FILTER_REJECT` and descent.
- **DOMTokenList operations**: `add`/`remove`/`toggle`/`replace` each check for empty string or ASCII whitespace; `toggle` and `replace` may skip update steps for compatibility; `supports` uses validation steps.
- **XPath**: `createExpression`, `evaluate` on `XPathEvaluatorBase`; `evaluate` on `XPathExpression` returns `XPathResult` with types like `ANY_TYPE`, `ORDERED_NODE_ITERATOR_TYPE`, etc.
- **XSLTProcessor**: methods for importing stylesheets, transforming to fragment/document, setting/getting parameters, resetting.

## Nuance Or Contradictions
- `detach()` on both `Range` and `NodeIterator` is a no‑op kept for compatibility only.
- `DOMTokenList` name is acknowledged as an “unfortunate legacy mishap”.
- Update steps for `toggle()` and `replace()` may not always run, explicitly noted for web compatibility.
- XPath and XSLT sections contain no complete algorithm definitions; they are placeholders pending further work (`whatwg/dom#67`, `whatwg/dom#181`).
- `createNSResolver(nodeResolver)` returns the node argument directly; it exists only for historical reasons.

## Candidate Wiki Hints
- **Range Modifications** – covering `deleteContents`, `extractContents`, `cloneContents`, and `insertNode`/`surroundContents`.
- **NodeIterator vs TreeWalker** – comparison of traversal models, `pointerBeforeReference`, and pre‑remove behaviors.
- **DOMTokenList** – token set manipulation, validation, and web‑compat quirks.
- **XPath and XSLT Legacy APIs** – summary of present but incomplete DOM Level 3 XPath and XSLT interfaces.
- **Range Point Comparisons** – `isPointInRange`, `comparePoint`, `intersectsNode`, and stringification.
