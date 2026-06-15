---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context

This chunk details section **4.5. Interface Document** of the DOM Standard specification. It defines the `Document` interface, its attributes (such as `URL`, `compatMode`, and `doctype`), and creation methods (`createElement`, `createTextNode`, etc.). It also covers `XMLDocument`, element creation options (including custom element registries), namespace handling, cloning/importing nodes, and the adoption algorithm for moving nodes between documents.

## Local Summary

The section establishes the `Document` interface as a subclass of `Node`. It defines default values for document properties (e.g., encoding is UTF-8, type is "xml", mode defaults to "no-quirks"). The chunk details how to create elements with specific namespaces or local names, handling legacy string options and custom element registries. It explains the behavior of `getElementsByTagName` and `getElementsByClassName`, including case-sensitivity nuances in HTML documents. Finally, it outlines the algorithms for importing nodes (deep cloning), adopting nodes from other documents, and managing custom element registries within shadow roots during adoption.

## Key Claims

- A document's default encoding is UTF-8, content type is "application/xml", URL is "about:blank", origin is opaque, type is "xml", mode is "no-quirks", declarative shadow roots are false, and the custom element registry is null.
- The `Document` interface has legacy aliases: `charset` and `inputEncoding` map to `characterSet`.
- `compatMode` returns "BackCompat" if the document mode is "quirks"; otherwise "CSS1Compat".
- `createElement(localName)` lowercases `localName` in HTML documents; `createElementNS` handles namespace prefixes.
- `getElementsByClassName` interprets the argument as a space-separated list of classes and requires all specified classes to be present on an element.
- Cloning a document or shadow root via `importNode` throws a "NotSupportedError".
- Adopting a node into a new document updates the `node document` for the node and its inclusive descendants (shadow-including tree order).

## Entities And Concepts

- **Document Interface**: The base interface for documents, inheriting from `Node`.
- **XMLDocument**: An interface extending `Document` for XML content.
- **Modes**: "no-quirks", "quirks", and "limited-quirks" (formerly "standards mode" and "almost standards mode").
- **CustomElementRegistry**: Used to define custom elements, passed via options during creation or import.
- **Adoption Algorithm**: The process of moving a node from one document to another, updating references for the node and its inclusive descendants.
- **Namespace Handling**: Rules for `createElementNS` regarding prefixes, empty namespaces, and reserved names like "xmlns".

## Procedures And API Details

### Document Attributes
- `implementation`: Returns the associated `DOMImplementation`.
- `URL` / `documentURI`: Return the document's URL.
- `compatMode`: Returns "BackCompat" (quirks mode) or "CSS1Compat".
- `characterSet` / `charset` / `inputEncoding`: Return the encoding name.
- `doctype`: Returns the document type declaration or null.
- `documentElement`: Returns the root element of the document.

### Creation Methods
- `createElement(localName, options)`: Creates an element. In HTML documents, local names are lowercased. Options can specify a custom element registry or a customized built-in element via `is`.
- `createElementNS(namespace, qualifiedName, options)`: Creates an element in a specific namespace. Handles prefix extraction and validates namespace usage.
- `createDocumentFragment()`: Returns a new fragment node.
- `createTextNode(data)`, `createCDATASection(data)`, `createComment(data)`, `createProcessingInstruction(target, data)`: Return corresponding node types.

### Collection Methods
- `getElementsByTagName(qualifiedName)`: Returns an `HTMLCollection` of matching elements. Case-insensitive matching for HTML namespace elements in HTML documents.
- `getElementsByTagNameNS(namespace, localName)`: Filters by namespace and local name. Wildcards (`*`) allow partial matching.
- `getElementsByClassName(classNames)`: Returns elements containing all classes in the space-separated string argument.

### Cloning and Adoption
- `importNode(node, options)`: Creates a deep copy of a node (descendants included if `selfOnly` is false or option is true). Throws "NotSupportedError" for documents or shadow roots. Supports setting custom element registry on cloned elements.
- `adoptNode(node)`: Moves a node to the current document. Throws errors for documents or shadow roots with incompatible registries.

### Flattening Options
The `flatten element creation options` algorithm extracts `customElementRegistry` and `is` from an options object or string, validating that if a registry is provided, it matches the document's registry or has a scoped `is` flag set to false.

## Nuance Or Contradictions

- **Mode Renaming**: "Standards mode" and "almost standards mode" were renamed to "no-quirks" and "limited-quirks" respectively due to semantic issues and Ian Hickson's veto.
- **Case Sensitivity in `getElementsByTagName`**: In an HTML document, `<FOO>` (non-HTML namespace) and `<foo>` (HTML namespace) are matched separately from `<FOO>` (HTML namespace). The method matches case-insensitively only within the HTML namespace for HTML documents.
- **Class Matching Logic**: `getElementsByClassName("aaa bbb")` matches elements having both "aaa" and "bbb". Spaces are delimiters; commas or other characters do not split classes in the same way (e.g., `"aaa,bbb"` returns no nodes if no element has exactly those comma-separated classes).
- **Legacy String Options**: The `options` parameter for `createElement` and `createElementNS` can be a string for web compatibility, though the primary definition uses dictionaries.

## Candidate Wiki Hints

1. **Document Interface Attributes**: A page summarizing `compatMode`, `URL`, `characterSet`, and other read-only properties of the `Document` interface.
2. **Element Creation Options**: Documentation on how to use `customElementRegistry` and `is` in `createElement` options, including error cases like "NotSupportedError".
3. **getElementsByClassName Behavior**: A guide explaining space-separated class lists, case sensitivity, and examples of matching logic (e.g., `"aaa bbb"` vs `"aaa,bbb"`).
4. **Node Import and Adoption**: An explanation of the difference between `importNode` (cloning) and `adoptNode` (moving), including restrictions on shadow roots and document nodes.
