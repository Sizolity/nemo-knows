---
title: Chunk 20 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
This chunk covers section "11. Historical" of the DOM Standard Living Specification. It documents interfaces and interface members that were removed from previous versions or iterations of the standard to streamline the API and align with modern usage patterns.

Local Summary
The document lists specific DOM interfaces (such as `DOMConfiguration` and `Entity`) and various member methods/properties on core nodes like `Document`, `Element`, and `Node` that are no longer part of the current standard. The section concludes with acknowledgments to contributors and copyright/license information for the Living Standard.

Key Claims
- Several interfaces have been removed from the standard entirely, including `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMImplementationSource`, `DOMLocator`, `DOMObject`, `DOMUserData`, `Entity`, `EntityReference`, `MutationEvent`, `MutationNameEvent`, `NameList`, `Notation`, `RangeException`, `TypeInfo`, and `UserDataHandler`.
- Specific members have been removed from existing interfaces:
  - **Attr**: `schemaTypeInfo`, `isId`
  - **Document**: `createEntityReference()`, `xmlEncoding`, `xmlStandalone`, `xmlVersion`, `strictErrorChecking`, `domConfig`, `normalizeDocument()`, `renameNode()`
  - **DocumentType**: `entities`, `notations`, `internalSubset`
  - **DOMImplementation**: `getFeature()`
  - **Element**: `schemaTypeInfo`, `setIdAttribute()`, `setIdAttributeNS()`, `setIdAttributeNode()`
  - **Node**: `isSupported`, `getFeature()`, `getUserData()`, `setUserData()`
  - **NodeIterator**: `expandEntityReferences`
  - **Text**: `isElementContentWhitespace`, `replaceWholeText()`
  - **TreeWalker**: `expandEntityReferences`
- The standard is written by Anne van Kesteren with substantial contributions from Aryeh Gregor and Ms2ger.
- Portions of the revision history related to custom elements are available in the w3c/webcomponents repository under the W3C Software and Document License.
- The work is licensed under a Creative Commons Attribution 4.0 International License, with source code portions under the BSD 3-Clause License.

Entities And Concepts
- DOM Standard (Living Standard)
- Removed Interfaces: `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMImplementationSource`, `DOMLocator`, `DOMObject`, `DOMUserData`, `Entity`, `EntityReference`, `MutationEvent`, `MutationNameEvent`, `NameList`, `Notation`, `RangeException`, `TypeInfo`, `UserDataHandler`
- Removed Members: `schemaTypeInfo`, `isId`, `createEntityReference()`, `xmlEncoding`, `xmlStandalone`, `xmlVersion`, `strictErrorChecking`, `domConfig`, `normalizeDocument()`, `renameNode()`, `entities`, `notations`, `internalSubset`, `getFeature()`, `setIdAttribute()`, `setIdAttributeNS()`, `setIdAttributeNode()`, `isSupported`, `getUserData()`, `setUserData()`, `expandEntityReferences`, `isElementContentWhitespace`, `replaceWholeText()`
- Contributors: Anne van Kesteren, Aryeh Gregor, Ms2ger

Procedures And API Details
- No new procedures are defined in this chunk; the focus is on historical removals.
- The `getFeature()` method was removed from `DOMImplementation` and `Node`.
- The `expandEntityReferences` member was removed from both `NodeIterator` and `TreeWalker`.

Nuance Or Contradictions
- None observed within this specific chunk; it serves purely as a record of deprecations/removals.

Candidate Wiki Hints
- Page: Historical DOM Interfaces (Summary of removed APIs)
