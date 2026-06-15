---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
This chunk covers sections 4.2.5 through 4.3 of the DOM standard, detailing mixins for document and shadow root handling, parent node manipulation (including `prepend`, `append`, `replaceChildren`), sibling navigation (`previousElementSibling`), child node insertion/removal (`before`, `after`, `remove`), slottable elements, old-style collections (`NodeList`, `HTMLCollection`), and the mechanics of mutation observers.

Local Summary
The standard defines specific interface mixins to share functionality between `Document` and `ShadowRoot`. It details how nodes are converted into lists, how parent nodes manage their children via new methods like `prepend` and `append`, and how to navigate siblings excluding doctypes. It clarifies the behavior of legacy collections versus modern iterables and explains the microtask queue mechanism used by mutation observers to notify callbacks without losing data during subtree removals.

Key Claims
- The `DocumentOrShadowRoot` mixin provides a `customElementRegistry` attribute shared by both `Document` and `ShadowRoot`.
- New parent node methods (`prepend`, `append`, `replaceChildren`) automatically convert string arguments into `Text` nodes.
- These mutation methods throw a "HierarchyRequestError" DOMException if tree constraints are violated.
- The `previousElementSibling` and `nextElementSibling` attributes are excluded from `DocumentType` nodes for web compatibility reasons.
- `HTMLCollection` is described as a historical artifact; new API designers should use `sequence<T>` instead.
- Mutation observers operate via a microtask queue to ensure callbacks are invoked even if the DOM changes during notification.

Entities And Concepts
- **Mixin**: A mechanism to add shared attributes and methods to interfaces (e.g., `DocumentOrShadowRoot`, `ParentNode`).
- **CustomElementRegistry**: The registry object for custom elements, accessible via the mixin.
- **HierarchyRequestError**: A DOMException thrown when structural constraints are violated during manipulation.
- **NodeList vs HTMLCollection**: Distinguishes between live collections of nodes and historical element-only collections with named item lookup.
- **MutationObserver Microtask**: The internal queue mechanism ensuring observer callbacks fire correctly despite DOM mutations.

Procedures And API Details
- **Parent Node Conversion**: Strings in a list are replaced with `Text` nodes; single-node lists return the node directly, while multiple nodes are wrapped in a `DocumentFragment`.
- **Prepend/Append Logic**: Inserts converted nodes before the first child or after the last child respectively.
- **MoveBefore Logic**: Moves a node into a parent before a specified reference child (or the last child if none specified), preserving state.
- **NamedItem Lookup**: Returns the first element with a matching ID or `name` attribute (in HTML namespace) that hasn't been returned previously in the iteration.
- **Observer Notification Steps**: Clone pending observers, empty the queue, remove transient observers from nodes, invoke callbacks with records, and fire `slotchange` events.

Nuance Or Contradictions
- **Doctype Siblings**: The standard explicitly notes that sibling element attributes are not exposed on doctypes to maintain web compatibility, preventing potential inconsistencies.
- **Live vs Static Collections**: While most collections must be live, the text implies a distinction exists for specific cases where a snapshot might be acceptable unless otherwise stated (though the default requirement is liveness).
- **Legacy Artifacts**: The standard explicitly advises against using `HTMLCollection` in new designs, marking it as a legacy artifact to be phased out.

Candidate Wiki Hints
- **DocumentOrShadowRoot Mixin**: A page explaining how custom element registries are shared between the main document and shadow roots.
- **ParentNode Methods Deep Dive**: Documentation on the new `prepend`, `append`, and `replaceChildren` methods, including their error handling and string-to-node conversion behavior.
- **MutationObserver Microtask Queue**: An explanation of the internal microtask mechanism that prevents observer callbacks from being skipped during rapid DOM changes.
