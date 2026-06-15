---
title: Chunk 29 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk covers **Section 11. Historical**, detailing the browser and engine compatibility support for various DOM interfaces, attributes, and methods found in `raw/web/corpus-2026-05-18/060-dom-standard.md`. The data spans from line 12168 to 12902.

## Local Summary
The document lists compatibility matrices for a wide range of DOM APIs, including Event properties (e.g., `cancelable`, `composedPath`), EventTarget methods (`addEventListener`, `dispatchEvent`), HTMLCollection attributes (`item`, `length`), and MutationObserver interfaces. Each entry specifies the minimum version required for support across major browsers (Firefox, Safari, Chrome, Edge, Opera) and environments (Android WebView, Node.js).

## Key Claims
- **Universal Support**: Many core Event properties like `Event/type` and `Event/target` are supported in all current engines starting from version 1.5 or 1 for Firefox/Safari/Chrome respectively.
- **MutationObserver Evolution**: The `MutationObserver` interface is supported in Firefox 14, Safari 7, and Chrome 26, with specific methods like `takeRecords` requiring later versions (e.g., Chrome 20).
- **Node.js Support**: The DOM standard APIs generally require Node.js version 14.5.0 or higher for full compatibility.
- **Historical Context**: Older browsers like Edge Legacy and IE9/IE8 are listed with specific minimum versions where applicable, noting "None" for unsupported features in those environments.

## Entities And Concepts
- **Event Properties**: `cancelable`, `composed`, `composedPath`, `currentTarget`, `defaultPrevented`, `eventPhase`, `isTrusted`, `preventDefault`, `stopImmediatePropagation`, `stopPropagation`, `target`, `timeStamp`, `type`.
- **EventTarget Methods**: `addEventListener`, `dispatchEvent`, `removeEventListener`.
- **HTMLCollection Attributes**: `item`, `length`, `namedItem`.
- **MutationObserver Interfaces**: `MutationObserver`, `disconnect`, `observe`, `takeRecords`.
- **MutationRecord Attributes**: `addedNodes`, `attributeName`, `attributeNamespace`, `nextSibling`, `oldValue`, `previousSibling`, `removedNodes`, `target`, `type`.
- **Browser Engines**: Firefox, Safari, Chrome, Edge (Legacy and Chromium), Opera, Android WebView, Samsung Internet.

## Procedures And API Details
- **Event Property Compatibility**:
  - `Event/cancelable`: Supported in Firefox 1.5+, Safari 1+, Chrome 1+.
  - `Event/composedPath`: Supported in Firefox 59+, Safari 10+, Chrome 53+.
  - `Event/isTrusted`: Supported in Firefox 1.5+, Safari 10+, Chrome 46+.
- **EventTarget Methods**:
  - `dispatchEvent`: Requires Firefox 2, Safari 3.1, Chrome 4.
  - `removeEventListener`: Supported in all current engines starting from version 1 for Firefox/Safari/Chrome.
- **MutationObserver Methods**:
  - `observe`: Available since Firefox 14, Safari 6, Chrome 18.
  - `takeRecords`: Requires Firefox 14, Safari 6, Chrome 20.
- **HTMLCollection Attributes**:
  - `item`, `length`, `namedItem`: All supported in all current engines starting from version 1 for Firefox/Safari/Chrome.

## Nuance Or Contradictions
- **Android WebView Discrepancies**: Some entries show "Android WebView" with a question mark or specific versions (e.g., 37+, 46+) that differ between properties, indicating potential fragmentation in Android support compared to desktop counterparts.
- **Edge Legacy vs. Edge Chromium**: The data distinguishes between Edge Legacy (based on IE) and the modern Edge Chromium, with some features missing entirely in the Legacy version ("IENone").
- **Version Gaps**: Certain features like `Event/stopImmediatePropagation` have a gap where Firefox requires version 10 while Chrome only needs version 5, highlighting uneven adoption rates.

## Candidate Wiki Hints
- **DOM Event Properties**: A page summarizing the evolution and browser support of standard Event properties (`cancelable`, `composedPath`, etc.).
- **MutationObserver API**: Documentation on observing DOM changes, detailing method availability across browsers.
- **HTMLCollection Interface**: Notes on accessing list items via `item` and `namedItem` in legacy vs. modern browsers.
