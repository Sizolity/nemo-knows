---
title: Chunk 06 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- Source file: `raw/web/corpus-2026-05-18/060-dom-standard.md`
- Chunk 6 of 9, lines 6888–8816
- Heading path: **11. Historical**
- Includes the historical removals section, acknowledgments, intellectual property rights, the full Index of terms, normative/informative references, and the IDL Index.

## Local Summary
This chunk enumerates every interface and interface member that has been removed from the DOM Standard. It then provides acknowledgments, copyright/licensing information, an exhaustive Index of all terms defined in the specification, references, and the complete IDL Index for the current standard.

## Key Claims
- The following interfaces have been removed from the DOM Standard:
  `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMImplementationSource`, `DOMLocator`, `DOMObject`, `DOMUserData`, `Entity`, `EntityReference`, `MutationEvent`, `MutationNameEvent`, `NameList`, `Notation`, `RangeException`, `TypeInfo`, `UserDataHandler`.
- The following interface members have been removed:

  **Attr**: `schemaTypeInfo`, `isId`

  **Document**: `createEntityReference()`, `xmlEncoding`, `xmlStandalone`, `xmlVersion`, `strictErrorChecking`, `domConfig`, `normalizeDocument()`, `renameNode()`

  **DocumentType**: `entities`, `notations`, `internalSubset`

  **DOMImplementation**: `getFeature()`

  **Element**: `schemaTypeInfo`, `setIdAttribute()`, `setIdAttributeNS()`, `setIdAttributeNode()`

  **Node**: `isSupported`, `getFeature()`, `getUserData()`, `setUserData()`

  **NodeIterator**: `expandEntityReferences`

  **Text**: `isElementContentWhitespace`, `replaceWholeText()`

  **TreeWalker**: `expandEntityReferences`
- The standard is written by Anne van Kesteren, with substantial contributions from Aryeh Gregor and Ms2ger.
- Copyright belongs to WHATWG (Apple, Google, Mozilla, Microsoft) under CC BY 4.0, with source code portions under BSD 3-Clause.
- The remaining content (Index, references, IDL Index) is the current complete definitional and cross-reference material for the entire DOM Standard.

## Entities And Concepts
- **Removed interfaces**: `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMImplementationSource`, `DOMLocator`, `DOMObject`, `DOMUserData`, `Entity`, `EntityReference`, `MutationEvent`, `MutationNameEvent`, `NameList`, `Notation`, `RangeException`, `TypeInfo`, `UserDataHandler`
- **Removed member groups**: by interface (see above)
- **Acknowledgments**: lists ~200 contributors
- **Intellectual property rights**: CC BY 4.0 for text, BSD 3-Clause for code; history in w3c/webcomponents repo under W3C Software and Document License
- **Index**: every term defined by the specification (alphabetical, with references to sections)
- **Terms defined by reference**: from other specifications (HTML, WebIDL, etc.)
- **References**: normative and informative
- **IDL Index**: complete Web IDL definitions of all current DOM interfaces

## Procedures And API Details
- None (this chunk documents removals, acknowledgments, and reference material, not operational procedures).

## Nuance Or Contradictions
- The list explicitly states that `ENTITY_REFERENCE_NODE` and `ENTITY_NODE` are “legacy” constants still present on `Node`, whereas the interfaces `Entity` and `EntityReference` are entirely removed.
- `Node`’s `ENTITY_REFERENCE_NODE` and `ENTITY_NODE` constants are marked “legacy” in the current IDL, even though the corresponding interfaces no longer exist.
- The acknowledgments and IP clauses are unchanged historical records, not normative requirements.

## Candidate Wiki Hints
- A wiki page **“Deprecated and removed DOM interfaces”** could list the removed interfaces and members along with their original version, reasons for removal, and migration notes.
- A **“DOM Standard history”** page could link to this removal list and explain the evolution of DOM levels.
