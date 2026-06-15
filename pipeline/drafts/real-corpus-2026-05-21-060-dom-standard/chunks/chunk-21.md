---
title: Chunk 21 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

This chunk covers the **Historical** section (§ 11) of the DOM Standard, listing deprecated or legacy interfaces and concepts. It also includes a glossary-style index of terms from earlier sections (e.g., AbortController, Event handling, Node manipulation) that appear in the source document's metadata or cross-references, though the primary focus here is on historical features like `DOMImplementation`, `EntityReference`, and legacy activation behaviors.

# Local Summary

The text defines a collection of interfaces and attributes marked for historical context, including `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMLocator`, `DOMObject`, `DOMTokenList`, `DOMUserData`, and `entities`. It also lists legacy activation behaviors (`legacy-canceled-activation behavior`, `legacy-pre-activation behavior`) and specific historical methods like `createEntityReference()` and `getUserData()`.

# Key Claims

- The following interfaces are categorized under the **Historical** section:
  - `DOMConfiguration`
  - `DOMError`
  - `DOMErrorHandler`
  - `DOMImplementationList`
  - `DOMLocator`
  - `DOMObject`
  - `DOMTokenList`
  - `DOMUserData`
  - `entities`
  - `Entity`
  - `ENTITY_NODE`
  - `ENTITY_REFERENCE_NODE`
  - `EntityReference`
- Legacy behaviors include:
  - `legacy-canceled-activation behavior` (§ 2.7)
  - `legacy-obtain service worker fetch event listener callbacks` (§ 2.8)
  - `legacy-pre-activation behavior` (§ 2.7)
- Historical methods/functions listed:
  - `createEntityReference()` (§ 11)
  - `getUserData()` (§ 11)
  - `getFeature()` for `DOMImplementation` and `Node` (§ 11)

# Entities And Concepts

### Interfaces (Historical/Deprecated Context)
- `DOMConfiguration`: Configuration settings for the DOM implementation.
- `DOMError`: Represents a generic error in the DOM.
- `DOMErrorHandler`: Interface for handling errors during document parsing.
- `DOMImplementationList`: A collection of DOM implementations.
- `DOMLocator`: Location information within the DOM.
- `DOMObject`: Generic object interface for the DOM.
- `DOMTokenList`: Token list for attributes (e.g., class, rel).
- `DOMUserData`: User-defined data associated with nodes.
- `Entity`: Represents an entity in the document.
- `EntityReference`: Reference to an external entity.

### Attributes & Constants
- `entities`: Collection of entities defined in the document type definition.
- `ENTITY_NODE`: Constant for node type representing an entity.
- `ENTITY_REFERENCE_NODE`: Constant for node type representing an entity reference.
- `isElementContentWhitespace`: Boolean flag related to whitespace handling (legacy).
- `isSupported`: Method/attribute indicating feature support (legacy).

# Procedures And API Details

### Creation and Management
- **`createEntityReference()`**: Creates an EntityReference node. Used historically for referencing external entities defined in a DTD.
- **`getUserData()`**: Retrieves user-defined data from a node. Deprecated in favor of `dataset` or other custom attributes.
- **`getFeature(name, version)`**: (Implied via `DOMImplementation`) Gets an implementation object for a specific feature.

### Feature Support
- **`hasFeature(feature, version)`**: Checks if the DOM implementation supports a specific feature at a given version level (historical method).

### Event and Listener Contexts (Cross-referenced)
While not strictly in § 11, the chunk references legacy event handling concepts:
- `legacy-canceled-activation behavior`: Relates to Service Worker activation states.
- `legacy-pre-activation behavior`: Another state in the Service Worker lifecycle prior to modern standards.

# Nuance Or Contradictions

- **Deprecation vs. Usage**: Interfaces like `DOMTokenList` are still widely used in modern web development (e.g., for managing class lists), yet they appear here under "Historical" or alongside legacy concepts. This suggests the source document may be documenting the *evolution* of these APIs, distinguishing between their current usage and their original historical context or deprecated variants.
- **`createEntityReference()`**: This method is largely obsolete in modern HTML5/HTML parsers which no longer support external DTDs by default, making this API historically significant but practically unused in new applications.
- **Legacy Activation Behaviors**: The terms `legacy-canceled-activation behavior` and `legacy-pre-activation behavior` indicate transitional states in the Service Worker specification that have been refined or replaced in later versions of the standard.

# Candidate Wiki Hints

1. **Historical DOM Interfaces**
   - Title: `DOMImplementationList`, `DOMError`, `DOMErrorHandler`
   - Description: Overview of deprecated or legacy interfaces in the DOM Level 3 and early implementations.
   - Content Focus: Usage, deprecation status, and replacement mechanisms (e.g., using exceptions instead of `DOMError`).

2. **Entity References and DTDs**
   - Title: `EntityReference`, `createEntityReference`
   - Description: Explaining the concept of entity references in XML/HTML documents and why they are deprecated in HTML5.
   - Content Focus: Differences between internal and external entities, parsing limitations in modern browsers.

3. **Legacy Service Worker Activation**
   - Title: `legacy-pre-activation behavior`, `legacy-canceled-activation behavior`
   - Description: Historical states of Service Worker registration and activation before the current lifecycle model.
   - Content Focus: How old versions handled installation, activation, and termination compared to modern specs.

4. **DOMTokenList Utility**
   - Title: `DOMTokenList`, `classList`
   - Description: Practical usage of `DOMTokenList` for managing multiple attribute values (like `class` or `rel`).
   - Content Focus: Methods like `.add()`, `.remove()`, `.toggle()`, and its role as a bridge between DOM nodes and arrays.
