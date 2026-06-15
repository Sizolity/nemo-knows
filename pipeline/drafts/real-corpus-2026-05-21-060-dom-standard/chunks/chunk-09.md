---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

Chunk Context
The chunk covers section 4.4, "Interface Node," from the DOM Standard. It details the `Node` interface definition, including its attributes, methods, and associated dictionaries. The text explains the behavior of node types (element, attribute, text, etc.), their properties like `nodeType`, `nodeName`, `baseURI`, and relationships such as `parentNode`, `childNodes`, and siblings. It also covers utility methods for cloning nodes (`cloneNode`), comparing positions within the document tree (`compareDocumentPosition`), and manipulating text content (`textContent`, `nodeValue`).

Local Summary
This section defines the `Node` interface, which is abstract and implemented by all DOM nodes. It lists constants for node types (e.g., `ELEMENT_NODE`, `TEXT_NODE`) and provides getter/setter logic for attributes like `nodeType`, `nodeName`, `baseURI`, `isConnected`, `ownerDocument`, and traversal properties like `parentNode`, `firstChild`, `previousSibling`. Methods include `getRootNode()`, `hasChildNodes()`, `appendChild()`, `removeChild()`, `cloneNode()`, `isEqualNode()`, `compareDocumentPosition()`, and `contains()`. Special attention is given to text content handling via `textContent` and the normalization of text nodes using `normalize()`. The cloning algorithm accounts for custom element registries and shadow roots. Node equality checks are defined based on structural properties like namespace, local name, attribute lists, and children.

Key Claims
- The `Node` interface is abstract; direct instances cannot be obtained.
- Every node has an associated "node document," set upon creation and potentially changed via the adopt algorithm.
- Node types are represented by unsigned short constants (e.g., `ELEMENT_NODE = 1`, `ATTRIBUTE_NODE = 2`).
- The `nodeName` property returns specific strings for different node types (e.g., "#text" for exclusive Text nodes, "null" or qualified names for others).
- The `baseURI` getter returns the serialized document base URL of the node's document.
- `isConnected` returns true if the node is part of a live tree.
- `getRootNode()` returns the root; with `{ composed: true }`, it includes shadow roots.
- `cloneNode(subtree)` creates a copy, including descendants if `subtree` is true. It handles custom element registries and shadow roots specifically.
- `isEqualNode(otherNode)` compares structural equality (interfaces, attributes, children) but does not consider live ranges or text offsets.
- `isSameNode(otherNode)` checks reference equality (strict `===`).
- `compareDocumentPosition(other)` returns a bitmask indicating relative position (preceding, following, contained by, contains).
- `normalize()` removes empty exclusive Text nodes and concatenates contiguous ones.

Entities And Concepts
- **Node**: Abstract interface for DOM nodes.
- **Node Types**: Element, Attr, Text (exclusive), CDATASection, ProcessingInstruction, Comment, Document, DocumentType, DocumentFragment, Notation (legacy).
- **Node Type Constants**: `ELEMENT_NODE`, `ATTRIBUTE_NODE`, `TEXT_NODE`, etc.
- **Traversal Properties**: `parentNode`, `parentElement`, `childNodes`, `firstChild`, `lastChild`, `previousSibling`, `nextSibling`.
- **Text Content**: Managed via `textContent` getter/setter and `nodeValue`.
- **Cloning**: `cloneNode()` with support for shadow roots and custom element registries.
- **Equality**: `isEqualNode()` (structural), `isSameNode()` (reference).
- **Position**: `compareDocumentPosition()`, `contains()`.
- **Normalization**: `normalize()` method for cleaning up text nodes.

Procedures And API Details
- **`getRootNode(options)`**: Returns the root node. If `options.composed` is true, returns the shadow-including root.
- **`cloneNode(subtree = false)`**: Clones a node and optionally its descendants. Handles shadow hosts and custom element registries. Throws "NotSupportedError" if called on a shadow root.
- **`isEqualNode(otherNode)`**: Returns true if `otherNode` is non-null and structurally equal to the current node (same interfaces, attributes, children).
- **`isSameNode(otherNode)`**: Legacy alias for strict equality (`===`).
- **`compareDocumentPosition(other)`**: Returns a bitmask combining flags like `DOCUMENT_POSITION_PRECEDING`, `DOCUMENT_POSITION_FOLLOWING`, etc. Handles attributes as preceding their element's children.
- **`contains(other)`**: Returns true if `other` is an inclusive descendant (or null).
- **`normalize()`**: Iterates descendants; removes empty exclusive Text nodes and merges contiguous ones, updating live ranges.
- **`textContent` Setter**: If value is null, treats as empty string; delegates to internal logic based on node type (Element, Attr, CharacterData).

Nuance Or Contradictions
- **Node Document**: A node's document is immutable except via the adopt algorithm. For documents themselves, `ownerDocument` returns null.
- **Attribute Handling in Position**: Attributes are treated as preceding their element's children in `compareDocumentPosition`, even though they don't participate in the same tree structure for traversal purposes.
- **Cloning Shadow Roots**: If a shadow root is cloned and its `clonable` flag is true, the clone includes a new shadow root with specific properties copied from the original (mode, serializable, delegates focus, etc.).
- **Legacy Constants**: Some node types like `ENTITY_REFERENCE_NODE`, `ENTITY_NODE`, and `NOTATION_NODE` are marked as legacy.
- **Text Content Logic**: The logic for getting/setting text content differs between Element descendants (descendant text), Attr (value), and CharacterData (data).

Candidate Wiki Hints
- **Node Interface Overview**: A page explaining the `Node` interface, its abstract nature, and common usage patterns.
- **Node Types and Constants**: Documentation on node type constants (`ELEMENT_NODE`, `TEXT_NODE`) and their corresponding types.
- **Cloning Nodes**: Deep dive into `cloneNode()`, including handling of shadow DOMs and custom elements.
- **Document Position**: Explanation of the bitmask returned by `compareDocumentPosition()` and how to interpret it.
- **Text Manipulation**: Guide on using `textContent`, `nodeValue`, and `normalize()` for text content management.
