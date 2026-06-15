---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Group Context

This group of notes synthesizes the **DOM Standard** specification, specifically covering the transition from general infrastructure to detailed node tree manipulation, event systems, and legacy API maintenance. The source material spans from fundamental tree structures (Trees, Ordered Sets) through the comprehensive definition of the `Element` interface, attribute handling, shadow DOM integration, range selection algorithms, and traversal mechanisms (`NodeIterator`, `TreeWalker`). The latter portions of the document address legacy features such as XPath/XSLT and historical notes, while maintaining conformance definitions for modern APIs like Custom Elements and Mutation Observers.

# Cross-Chunk Summary

The document is structured as a progression from abstract infrastructure to concrete DOM interfaces:

1.  **Infrastructure (Chunks 01–04):** Establishes the foundational concepts of trees, ordered sets, selectors, and name validation. It introduces the event system lifecycle: from introduction to interface definitions (`Event`, `CustomEvent`), target management (`EventTarget`), listener observation, dispatching/firing, and finally, activity abortion via `AbortController`/`AbortSignal`.
2.  **Node Tree & Mixins (Chunks 05–08):** Defines the hierarchy of node interfaces. It details the `Node` interface, document/shadow tree structures (including slots/slottables), mutation algorithms, and a suite of mixin interfaces (`NonElementParentNode`, `DocumentOrShadowRoot`, `ParentNode`, `ChildNode`, `Slottable`). This section also covers old-style collections (`NodeList`, `HTMLCollection`) and introduces Mutation Observers.
3.  **Core Interfaces (Chunks 09–15):** Focuses on specific node types (`Node`, `Document`, `DocumentType`, `DocumentFragment`, `ShadowRoot`, `Element`). It provides deep dives into attribute handling via `NamedNodeMap` and `Attr`, character data nodes (`Text`, `Comment`, etc.), and the introduction of DOM Ranges.
4.  **Range & Traversal (Chunks 16–18):** Details the `Range` interface, including boundary points, containment logic, and manipulation algorithms (delete, extract, clone). It covers traversal interfaces (`NodeIterator`, `TreeWalker`) with filtering logic via `NodeFilter` and the set-like behavior of `DOMTokenList`.
5.  **Legacy & Security (Chunks 19–32):** Concludes with XPath/XSLT interfaces, security/privacy considerations, and a large section dedicated to historical context regarding deprecated or legacy features.

# Repeated Or Central Claims

- **Element Definition:** An element is fundamentally defined by its custom element state ("uncustomized" or "custom"). It possesses a tag name, optional namespace/prefix, and an associated shadow root (null by default).
- **Attribute Semantics:** Attributes like `id`, `class`, and `slot` are super-global content attributes. The `id` concept is formalized as unique per element (replacing historical DTD-based multiple identifiers). Qualified names are computed from local name and prefix, with HTML documents requiring uppercase normalization for tag names.
- **Live Ranges:** `Range` objects are "live," meaning they automatically update their boundary points when the underlying DOM tree mutates (insertion, removal, replacement). `StaticRange` is the exception that does not update.
- **Shadow Host Restrictions:** Only elements with valid local names (e.g., standard HTML elements or registered custom elements) can act as shadow hosts. Attaching a shadow root to an invalid host throws a `NotSupportedError`. The `shadowRoot` getter returns null if the mode is "closed".
- **Mutation & Callbacks:** Changing attributes triggers mutation records. For custom elements, this also enqueues upgrade reactions and invokes `attributeChangedCallback` if the element is defined.
- **Traversal Mechanics:** Both `NodeIterator` and `TreeWalker` rely on a `whatToShow` bitmask and a `filter` (via `NodeFilter`) to determine node visibility. The `detach()` method exists in both for compatibility but performs no action in the current spec.
- **DOMTokenList Validation:** Token lists strictly forbid empty strings and tokens containing ASCII whitespace, throwing `SyntaxError` or `InvalidCharacterError` respectively.

# Important Local Details

- **Interface Hierarchy:** The spec uses a mixin approach where interfaces like `ParentNode`, `ChildNode`, and `NonDocumentTypeChildNode` are layered onto base nodes to provide specific traversal capabilities (e.g., `closest()`, `insertBefore()`).
- **Range Algorithms:** Operations like `deleteContents()` and `extractContents()` involve complex logic to reconstruct start/end nodes if they become invalid after tree modification. `surroundContents()` throws errors for partial containment of non-Text nodes or invalid parent types.
- **Shadow DOM Init:** The `attachShadow` method accepts a dictionary (`ShadowRootInit`) allowing configuration of mode, delegates focus, and slot assignment options. `customElementRegistry` can be passed to pass a node directly in some contexts.
- **Boundary Point Logic:** A range's containment is strictly defined by offsets within nodes. A node is *not* contained if it defines the boundary itself; only its descendants (or content for CharacterData) are included.
- **Historical Aliases:** The spec retains legacy aliases such as `webkitMatchesSelector` alongside standard `matches`, and `insertAdjacentElement`/`insertAdjacentText` with legacy string mappings.
- **XPath Status:** While defined in the DOM spec, XPath Level 3 APIs (`XPathResult`, `XPathEvaluator`) are noted as legacy and not actively maintained, though definitions persist for future updates.

# Candidate Wiki Hints

- **Page: Element Interface Overview**: Summarize attributes, shadow root association, and custom element lifecycle states.
- **Page: Custom Element Lifecycle**: Explain the transition from "undefined" to "custom", upgrade mechanisms, and reaction handling.
- **Page: DOM Range Selection**: Detail live vs. static ranges, boundary points, and containment rules.
- **Page: Traversal Interfaces**: Compare `NodeIterator` (fixed root) vs. `TreeWalker` (mutable current node).
- **Page: Shadow DOM Host Requirements**: List valid element names and conditions for attaching shadow roots.
- **Page: Attribute Manipulation**: Guide on `setAttribute`, `removeAttribute`, and trusted type validation steps.

# Gaps Or Cautions

- **Missing Content Data:** The provided chunk notes summarize headings and claims but do not include the full text of algorithms or specific error message strings (e.g., exact `DOMException` codes beyond names).
- **Incomplete Historical Section:** Chunks 20–32 cover "Historical" content but the notes only indicate the heading. Specific details on deprecated APIs within this section are not elaborated in the provided summaries.
- **Algorithm Implementation Details:** While high-level logic for ranges and traversal is described, specific internal algorithm steps (e.g., exact order of mutation record queuing) are inferred from context rather than explicitly detailed in the notes.
- **Security Specifics:** The "Security and privacy considerations" section is listed but lacks detailed content in the provided notes, requiring cross-referencing with external security documentation for implementation guidance.
