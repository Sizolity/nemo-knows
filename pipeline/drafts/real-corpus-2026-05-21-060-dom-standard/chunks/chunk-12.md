---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context
This chunk covers DOM operations related to custom element registries, node adoption, document creation, and shadow roots. It includes definitions for `DOMImplementation`, `DocumentType`, `DocumentFragment`, and `ShadowRoot` interfaces, along with algorithms for attribute creation, event creation, range creation, node iteration, and tree walking.

# Local Summary
The text details how custom element registries are managed within documents and elements, specifically handling scoped vs. global registries. It outlines the `adoptNode()` method steps for moving nodes between documents, throwing specific errors for unsupported types like documents or shadow roots. The chunk also defines the `DOMImplementation` interface methods (`createDocumentType`, `createDocument`, `createHTMLDocument`) and notes that `hasFeature()` is obsolete. Definitions for `DocumentFragment` (host concept) and `ShadowRoot` (attributes like mode, delegatesFocus, host) are provided, including traversal orders for shadow-including trees.

# Key Claims
- A document's effective global custom element registry is the document's custom element registry if it is a global registry; otherwise, it is null.
- The `adoptNode()` method throws a "NotSupportedError" if the node is a document and a "HierarchyRequestError" if it is a shadow root.
- The `hasFeature()` method on `DOMImplementation` always returns true and is no longer reliable for feature detection; it exists for backward compatibility.
- Shadow roots have an associated mode ("open" or "closed") and an associated host which is never null.
- In shadow-including tree order, traversal includes a depth-first walk of the shadow root's node tree immediately after encountering the shadow host.

# Entities And Concepts
- **Custom Element Registry**: Can be global or scoped; determines how custom elements are registered.
- **DOMImplementation**: Interface for creating document types and documents.
  - `createDocumentType()`: Creates a doctype.
  - `createDocument()`: Creates an XMLDocument.
  - `createHTMLDocument()`: Creates an HTML document with basic structure.
- **DocumentType**: Represents the doctype of a document (name, publicId, systemId).
- **DocumentFragment**: A node used for fragment manipulation; has an associated host.
- **ShadowRoot**: A special type of DocumentFragment acting as the root of a shadow tree.
  - Attributes: `mode`, `delegatesFocus`, `slotAssignment`, `clonable`, `serializable`, `host`.
- **Shadow-including Tree Order**: Traversal order including shadow DOM subtrees.

# Procedures And API Details
- **`createAttribute(localName)`**: Returns a new attribute; throws "InvalidCharacterError" if localName is invalid. In HTML documents, the local name is lowercased.
- **`createAttributeNS(namespace, qualifiedName)`**: Validates and extracts namespace/prefix/local name; returns a new attribute.
- **`createEvent(interface)`**: Creates an event object. Supports interfaces like `BeforeUnloadEvent`, `CustomEvent`, `MouseEvent`, etc. Throws "NotSupportedError" if the interface is not supported or exposed. Initializes type, timeStamp, and isTrusted attributes.
- **`createRange()`**: Returns a new live range starting and ending at index 0 of the current context.
- **`createNodeIterator(root, whatToShow, filter)`**: Creates a NodeIterator with specified root, visibility filter, and optional filter.
- **`createTreeWalker(root, whatToShow, filter)`**: Creates a TreeWalker with specified root, visibility filter, and optional filter.
- **`adoptNode(node)`**: Moves a node into the current document. Throws errors for documents or shadow roots. Returns the adopted node.

# Nuance Or Contradictions
- The `hasFeature()` method is described as "useless" and always returning true, yet it remains part of the interface specification to maintain backward compatibility with old pages.
- Event creation via `createEvent()` is noted as deprecated in favor of using constructors directly (e.g., `new MouseEvent()`).
- Shadow roots have an associated custom element registry which is initially null, but can be set; this interacts with the "keep custom element registry null" boolean flag relevant only for declarative shadow roots.

# Candidate Wiki Hints
- **Page: Custom Element Registry Scope** – Explaining global vs. scoped registries and how `adoptedCallback` reactions handle registry changes.
- **Page: DOMImplementation Interface** – Documenting the creation of documents and doctypes, including content type determination based on namespace.
- **Page: ShadowRoot Attributes** – Detailing `mode`, `delegatesFocus`, `host`, and other shadow root specific attributes.
- **Page: Event Creation API** – Listing supported event interfaces and initialization steps for `createEvent()`.
