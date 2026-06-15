---
title: Chunk 31 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
The provided text covers "11. Historical," listing browser compatibility data for DOM interfaces and methods (e.g., NodeIterator, NodeList, Range, ShadowRoot). The data includes support status ("✔MDN" or "?") and version numbers for Firefox, Safari, Chrome, Opera, Edge, and legacy IE/Android/iOS WebViews.

Local Summary
This section documents the historical browser support for various DOM API properties and methods. It details which browsers implemented specific features like `NodeList.item`, `Range.extractContents`, and `ShadowRoot.mode`, often distinguishing between modern engines and legacy versions (Edge Legacy, IE). The data frequently indicates full MDN support with a checkmark or lack thereof with a question mark for mobile WebViews.

Key Claims
- Most DOM interfaces listed (e.g., `NodeList`, `Range`, `Text`) are supported in all current engines as of the time of the source's compilation, often dating back to very early versions (Firefox 1, Safari 1, Chrome 1).
- Some specific methods like `ShadowRoot/slotAssignment` have higher minimum version requirements (e.g., Firefox 92, Chrome 86) compared to the core interface itself.
- Legacy browsers (IE, Edge Legacy) show inconsistent or non-existent support ("IENone", "Edge (Legacy)?") for newer standards features like Shadow DOM properties.
- Mobile WebViews (Firefox for Android, iOS Safari, Chrome for Android) often show missing data ("?") or specific version gaps compared to desktop counterparts.

Entities And Concepts
- **DOM Interfaces**: `NodeIterator`, `NodeList`, `Range`, `ProcessingInstruction`, `ShadowRoot`, `StaticRange`, `Text`.
- **Browser Engines**: Firefox, Safari, Chrome, Opera, Edge (Legacy), IE.
- **Mobile Platforms**: Firefox for Android, iOS Safari, Chrome for Android, Samsung Internet, WebView.
- **Compatibility Markers**: "✔MDN" (supported), "?" (unknown/unlisted), version numbers indicating first support.

Procedures And API Details
- `NodeList/forEach`: Supported in Firefox 50+, Safari 10+, Chrome 51+.
- `Range/toString`: Supported in all current engines, with early versions (Firefox 1, Safari 1, Chrome 1).
- `ShadowRoot/mode`: Supported in Firefox 63+, Safari 10.1+, Chrome 53+.
- `Text/wholeText`: Supported in Firefox 3.5+, Safari 4+, Chrome 2+.
- `Range/deleteContents`: Supported in Firefox 1–15, Safari 1+, Chrome 1+ (note the range for Firefox).

Nuance Or Contradictions
- The source distinguishes between "Edge" and "Edge (Legacy)," with the latter showing no support ("IENone") for many features that modern Edge supports.
- Data for mobile WebViews is often marked with "?", suggesting incomplete tracking compared to desktop browsers where specific version numbers are provided.
- Some entries list "In all current engines" but then provide early version numbers, implying the feature existed long ago in those engines but may have been removed or changed in others (though the text mostly implies stability).

Candidate Wiki Hints
- **DOM API Compatibility**: A page summarizing browser support tables for core DOM interfaces.
- **Shadow DOM History**: Tracking the introduction and version requirements of ShadowRoot properties across browsers.
- **Range API Evolution**: Notes on `Range` method availability, particularly older methods like `deleteContents` which had specific Firefox version ranges.
