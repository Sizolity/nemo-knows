---
title: Chunk 22 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
- **Heading**: 11. Historical
- **Scope**: Covers a comprehensive list of DOM interfaces, attributes, methods, and concepts referenced in the specification's historical section.
- **Content Type**: An enumerated index of terms (e.g., `MutationObserver`, `ShadowRoot`, `NodeIterator`) linked to their respective sections (§) within the document.

Local Summary
This chunk serves as an extensive glossary or index for the DOM Standard, specifically focusing on terminology and identifiers associated with historical versions or legacy concepts. It lists interfaces like `XPathEvaluator` and `XSLTProcessor`, shadow DOM components (`ShadowRootInit`, `SlotAssignmentMode`), mutation tracking APIs (`MutationRecord`, `MutationObserver`), node traversal utilities (`TreeWalker`, `NodeIterator`), and various attribute/method combinations (e.g., `setAttributeNS`, `querySelectorAll`). The entries reference specific sections (§) where definitions or algorithms for these terms are detailed.

Key Claims
- The DOM standard includes support for legacy processing instructions such as XSLT (`XSLTProcessor`) and XPath evaluation (`XPathEvaluator`).
- Shadow DOM functionality is represented by interfaces like `ShadowRoot`, `SlotAssignmentMode`, and attributes like `serializable` and `slot`.
- Mutation tracking is handled via the `MutationObserver` API, generating `MutationRecord` objects that track changes to the DOM tree.
- Node traversal is facilitated by `NodeIterator` and `TreeWalker` interfaces, which support filtering (`whatToShow`) and ordering modes (e.g., `SHOW_ELEMENT`, `SHOW_COMMENT`).
- Event handling mechanisms include standard listeners, abort signals (`AbortController`, `AbortSignal`), and propagation control methods like `stopPropagation()`.

Entities And Concepts
- **Interfaces**: `ShadowRoot`, `MutationObserver`, `NodeIterator`, `TreeWalker`, `XPathEvaluator`, `XSLTProcessor`, `Element`, `Attr`, `DocumentType`.
- **Attributes/Properties**: `namespaceURI`, `localName`, `textContent`, `ownerDocument`, `shadowRoot`, `slot`, `parentElement`, `nextSibling`.
- **Methods**: `appendChild`, `removeChild`, `replaceChild`, `querySelector`, `setAttribute`, `observe`, `disconnect`, `takeRecords`.
- **Constants/Types**: `ELEMENT_NODE`, `TEXT_NODE`, `COMMENT_NODE`, `DOCUMENT_FRAGMENT_NODE`, `SHOW_ALL`, `ORDERED_NODE_ITERATOR_TYPE`.
- **Events**: `MutationEvent`, `MutationNameEvent`, `SlotChangeEvent` (implied via `slotchange`).

Procedures And API Details
- **Mutation Observation**: The `MutationObserver` constructor accepts a callback and an optional `MutationObserverInit` dict. It queues microtasks when changes occur, creating `MutationRecord` objects with properties like `addedNodes`, `removedNodes`, and `type`.
- **Shadow DOM Initialization**: A `ShadowRoot` is initialized with an `ShadowRootInit` dictionary containing options like `slots` (for legacy slot assignment) and `serializable`.
- **Node Traversal**: `NodeIterator` and `TreeWalker` are configured using a filter function (`whatToShow`) and a root node. They support methods like `nextNode()`, `previousNode()`, and properties like `currentNode`.
- **Attribute Manipulation**: Methods include `setAttributeNS(namespace, qualifiedName, value)`, `removeAttribute(qualifiedName)`, and `toggleAttribute(qualifiedName, force)`.
- **Range Operations**: The `Range` interface allows manipulating text via `setStart(node, offset)`, `setEnd(node, offset)`, and `selectNodeContents(node)`.

Nuance Or Contradictions
- The section is titled "Historical," implying some listed concepts (like explicit `MutationEvent` or legacy slot assignment modes) may be deprecated or superseded by newer mechanisms like `MutationObserver` or the HTML slot element API, though they remain part of the standard's history.
- Some terms appear with both "dfn" (definition) and "attribute/method" markers, indicating they are defined concepts that also possess specific properties or behaviors in the DOM tree structure.

Candidate Wiki Hints
- **Shadow DOM & Slots**: Create a page explaining `ShadowRoot` internals, specifically focusing on `slotAssignment`, `serializable` attributes, and the legacy vs. modern slot handling mechanisms.
- **Mutation Tracking**: Document the `MutationObserver` API workflow, detailing how it interacts with `MutationRecord` and the microtask queue (`queue a mutation observer microtask`).
- **Tree Traversal**: Summarize the differences between `NodeIterator` and `TreeWalker`, including their configuration via `whatToShow` constants (e.g., `SHOW_ELEMENT`) and traversal order options.
- **XPath & XSLT**: Provide an overview of the legacy `XPathEvaluator` and `XSLTProcessor` interfaces, noting their section references and historical context in the DOM spec.
