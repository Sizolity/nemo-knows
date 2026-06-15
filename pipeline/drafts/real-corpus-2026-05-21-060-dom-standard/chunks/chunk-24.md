---
title: Chunk 24 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 11: Historical** of the DOM Standard, defining legacy interfaces and their current status. It details interfaces such as `DocumentType`, `ShadowRoot`, `Element`, `Attr`, character data types (`Text`, `Comment`), range manipulation (`Range`, `NodeIterator`), XPath processing, XSLT transformation, and the modernization of the Abort mechanism via `AbortController`. The text includes specific engine support tables (e.g., Firefox 57+, Chrome 66+) for these features.

# Local Summary

This section catalogs historical DOM interfaces that have been retained or superseded in modern standards. It defines the structure for `DocumentType`, `ShadowRoot` (with modes like "open"/"closed"), and the extensive attribute manipulation API on `Element`. The chunk also covers tree traversal tools (`TreeWalker`, `NodeIterator`) and legacy XPath/XSLT support. A significant portion is dedicated to the `AbortController` and `AbortSignal` interfaces, listing their specific properties (like `aborted`, `reason`) and providing a breakdown of browser compatibility for various Abort-related methods.

# Key Claims

- **Historical Interfaces**: The standard includes definitions for `DocumentType`, `ShadowRoot`, `Element`, `Attr`, `CharacterData`, `Text`, `CDATASection`, `ProcessingInstruction`, `Comment`, `AbstractRange`, `StaticRange`, `Range`, `NodeIterator`, `TreeWalker`, `XPathResult`, `XPathExpression`, `XSLTProcessor`.
- **Abort Mechanism**: The `AbortController` interface and its associated `signal` are supported in all current engines (Firefox 57+, Safari 12.1+, Chrome 66+).
- **Signal Properties**: Specific properties on `AbortSignal` such as `aborted`, `reason`, and methods like `throwIfAborted` have varying support dates (e.g., `reason` requires Firefox 97+).
- **Legacy Support**: Some legacy aliases exist, such as `webkitMatchesSelector` for `.matches()` and specific XPath result types.

# Entities And Concepts

- **Interfaces**: `DocumentType`, `ShadowRoot`, `Element`, `Attr`, `CharacterData`, `Text`, `Range`, `NodeIterator`, `TreeWalker`, `XPathEvaluator`, `XSLTProcessor`, `AbortController`, `AbortSignal`.
- **Enums**: `ShadowRootMode` ("open", "closed"), `SlotAssignmentMode` ("manual", "named").
- **Constants**: XPath result types (e.g., `ANY_TYPE`, `SNAPSHOT_TYPE`), Node filter flags (e.g., `SHOW_ELEMENT`, `SHOW_TEXT`).
- **Abort Features**: `signal`, `abort_event`, `aborted`, `reason`, `throwIfAborted`.

# Procedures And API Details

**Element Interface Methods**:
- `closest(selectors)`: Finds the closest ancestor matching a selector.
- `matches(selectors)`: Checks if the element matches a CSS selector (legacy alias: `webkitMatchesSelector`).
- `getElementsByTagName(qualifiedName)`: Returns an HTMLCollection of elements by tag name.
- `insertAdjacentElement(where, element)`: Inserts an element adjacent to the current one (marked legacy).

**Range Interface Methods**:
- `setStart(node, offset)`, `setEnd(node, offset)`: Sets range boundaries.
- `selectNode(node)`, `selectNodeContents(node)`: Selects a node or its contents within the range.
- `deleteContents()`, `extractContents()`, `cloneContents()`: Manipulates the range's contents (creates new objects).

**XPath Support**:
- `XPathEvaluator.createExpression(expression, resolver)`: Creates an expression object.
- `XPathEvaluator.evaluate(expression, contextNode, ...)`: Evaluates an XPath expression returning an `XPathResult`.

**Abort Signal Properties**:
- `signal.aborted`: Boolean indicating if the signal has been aborted.
- `signal.reason`: The reason for aborting (requires newer browser versions).
- `signal.throwIfAborted()`: Throws a DOMException if aborted.

# Nuance Or Contradictions

- **Legacy vs Modern**: Many methods like `insertAdjacentElement` are marked as legacy, while their modern counterparts or aliases (`closest`, `matches`) are preferred.
- **Engine Variance**: While the text states "In all current engines" for `AbortController`, specific properties like `signal.reason` have strict version requirements (e.g., Firefox 97+, Chrome 98+), indicating partial implementation history.
- **Useless Attributes**: Comments in the source mark attributes like `hasFeature()` and `Attr.specified` as "useless; always returns true", implying they are deprecated or obsolete for practical use.

# Candidate Wiki Hints

- **Page: DOM Element API** – Summarize attribute handling (`id`, `className`, `classList`) and query methods (`querySelector` logic implied via `closest/matches`).
- **Page: Range Manipulation** – Explain `Range` operations, including `deleteContents`, `extractContents`, and boundary setting.
- **Page: AbortController Guide** – Detail the lifecycle of `AbortController`, signal properties (`aborted`, `reason`), and usage in async contexts.
- **Page: Shadow DOM Basics** – Cover `ShadowRoot` attributes (`mode`, `delegatesFocus`) and initialization via `attachShadow`.
