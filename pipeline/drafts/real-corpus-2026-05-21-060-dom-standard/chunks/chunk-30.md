---
title: Chunk 30 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk covers the "Historical" section of a DOM API compatibility table, specifically detailing support matrices for `NamedNodeMap` methods and various `Node` properties and methods (including attributes like `baseURI`, `appendChild`, `childNodes`, `cloneNode`, etc.) across major browsers (Firefox, Safari, Chrome, Opera, Edge) and their legacy versions. It also includes data for `NodeIterator` properties.

## Local Summary
The document provides a granular breakdown of feature support for DOM Level 2/3 APIs within the `NamedNodeMap` interface and the `Node` interface. Each entry lists specific API names (e.g., `getNamedItemNS`, `insertBefore`) alongside browser engine versions required for compatibility, distinguishing between desktop browsers, mobile webviews, and legacy Internet Explorer versions.

## Key Claims
- **Universal Support**: Many core `Node` properties (like `nodeValue`, `nodeType`, `firstChild`) and methods (like `appendChild`, `removeChild`) are supported in "all current engines" starting from very early versions (Firefox 1+, Safari 1+, Chrome 1+).
- **Legacy IE Variance**: Support for specific `NamedNodeMap` methods often diverges in Internet Explorer, with some requiring IE6 and others requiring IE9.
- **Mobile WebView Gaps**: Android WebViews and certain mobile browsers sometimes lack support or have unknown (`?`) status compared to desktop counterparts.
- **Modern DOM Additions**: Features like `getRootNode` are noted as having significantly higher version requirements (e.g., Firefox 53+, Chrome 54+).

## Entities And Concepts
- **Interfaces**: `NamedNodeMap`, `Node`, `NodeIterator`.
- **Properties**: `baseURI`, `childNodes`, `firstChild`, `lastChild`, `nodeValue`, `parentNode`, `previousSibling`, `nextSibling`, `ownerDocument`, `parentElement`.
- **Methods**: `appendChild`, `cloneNode`, `compareDocumentPosition`, `contains`, `insertBefore`, `isConnected`, `normalize`, `removeChild`, `replaceChild`, `getNamedItemNS`, `setNamedItemNS`.
- **Iterators**: `filter`, `nextNode`, `pointerBeforeReferenceNode`, `previousNode`, `referenceNode`.

## Procedures And API Details
The chunk lists specific method signatures and property accessors without detailed syntax, focusing solely on browser compatibility matrices:
- **`NamedNodeMap` Methods**: Includes `getNamedItemNS`, `item`, `length`, `removeNamedItem`, `removeNamedItemNS`, `setNamedItem`, `setNamedItemNS`.
- **`Node` Properties**: Includes `baseURI` (Safari 4+), `compareDocumentPosition` (Chrome 2+), `getRootNode` (Firefox 53+, Chrome 54+).
- **`NodeIterator` Methods**: Includes `filter`, `nextNode`, `pointerBeforeReferenceNode`, `previousNode`.

## Nuance Or Contradictions
- **Inconsistent IE Support**: While many features are marked "IE6+" or "IE9+", some entries list "IENone" (e.g., `baseURI`, `contains`), indicating a complete lack of support in legacy Internet Explorer environments.
- **Unknown Mobile Status**: Several mobile browsers (Firefox for Android, iOS Safari, Chrome for Android) are marked with a question mark (`?`), signifying that the source data lacks definitive compatibility information for these specific platforms at the time of writing.

## Candidate Wiki Hints
- **Page: DOM API Compatibility Matrix** – A comprehensive table comparing browser support for various DOM Level 2 and 3 features.
- **Page: Node Interface** – Documentation covering standard properties and methods of the `Node` interface, including historical browser support.
- **Page: NamedNodeMap** – Details on collection interfaces for attributes, specifically focusing on namespaced item handling (`getNamedItemNS`).
