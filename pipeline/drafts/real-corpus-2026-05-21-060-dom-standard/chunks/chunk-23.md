---
title: Chunk 23 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk defines the "Historical" section of the standard, listing terms borrowed from various specifications (HTML, ECMAScript, Web IDL, etc.) and providing normative references for those standards. It also includes an IDL Index defining core interfaces like `Event`, `CustomEvent`, `EventTarget`, `AbortController`, `MutationObserver`, `Node`, and `Document`.

## Local Summary
The chunk outlines which external specifications contribute specific terminology to the DOM Standard and lists the full Web IDL definitions for fundamental eventing, node manipulation, document creation, and abort mechanisms. It serves as a glossary of dependencies and a foundational interface definition block.

## Key Claims
- The standard defines terms by reference to other living standards and W3C recommendations (e.g., HTML, ECMAScript, URL).
- `Event` is the base interface for all events, including `CustomEvent`.
- `EventTarget` provides the core lifecycle methods: `addEventListener`, `removeEventListener`, and `dispatchEvent`.
- `AbortController` and `AbortSignal` allow cancellation of operations.
- `MutationObserver` allows observation of changes to a tree of nodes.
- `Node` is the base interface for DOM nodes, defining structural relationships (parent/child/sibling) and creation methods (`appendChild`, `insertBefore`).
- `Document` extends `Node` with specific document-level methods like `createElement`, `getElementById`, and `querySelector`.

## Entities And Concepts
- **Event System**: `Event`, `CustomEvent`, `EventInit`, `EventTarget`, `EventListener`, `AbortController`, `AbortSignal`.
- **Tree Traversal & Mutation**: `MutationObserver`, `MutationRecord`, `MutationObserverInit`, `ParentNode` (mixin), `ChildNode` (mixin).
- **Core Node Interface**: `Node`, `Element`, `Document`, `DocumentFragment`, `CharacterData`, `Text`, `Comment`.
- **Abort Mechanism**: `signal`, `aborted`, `reason`, `abort()`.
- **DOM Construction**: `createElement`, `createTextNode`, `createDocumentFragment`, `importNode`, `adoptNode`.
- **Querying**: `querySelector`, `querySelectorAll`, `getElementById`, `getElementsByTagName`.
- **External Specs Referenced**: HTML, ECMAScript, URL, Web IDL, Infra, Encoding, Console, DeviceOrientation, etc.

## Procedures And API Details
- **Creating Events**: `new Event(type, options)`, `new CustomEvent(type, options)`.
- **Handling Events**: `target.addEventListener(type, callback, options)`, `target.removeEventListener(type, callback, options)`, `target.dispatchEvent(event)`.
- **Observing Changes**: `observer.observe(target, config)`, `observer.disconnect()`, `observer.takeRecords()`.
- **Node Manipulation**: `node.appendChild(newNode)`, `node.insertBefore(newNode, refNode)`, `node.replaceWith(node)`, `node.remove()`.
- **Document Creation**: `doc.createElement(name)`, `doc.createTextNode(data)`, `doc.createComment(data)`, `doc.createRange()`.
- **Abort Signals**: `AbortController.abort([reason])`, `AbortSignal.timeout(ms)`, `AbortSignal.any(signals)`.

## Nuance Or Contradictions
- The standard distinguishes between legacy properties (e.g., `srcElement`, `cancelBubble`) and modern ones, often marking the former as deprecated or aliases.
- Some methods like `createEvent` are marked as legacy.
- Certain attributes have multiple names (e.g., `charset` vs `characterSet`).

## Candidate Wiki Hints
- **Event System Architecture**: A page explaining the `Event`/`CustomEvent` inheritance and the event loop phases (`CAPTURING_PHASE`, `AT_TARGET`, `BUBBLING_PHASE`).
- **MutationObserver API**: A guide on observing DOM changes for performance-sensitive applications.
- **AbortController Usage**: How to use `AbortSignal` to cancel asynchronous operations (fetch, timers).
- **Node Manipulation Mixins**: Explaining how `ParentNode`, `ChildNode`, and `NonElementParentNode` extend the base `Node` interface.
- **Document Creation API**: Methods for creating elements and text nodes via the `Document` interface.
