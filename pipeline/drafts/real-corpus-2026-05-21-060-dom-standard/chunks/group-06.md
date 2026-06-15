---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Group Context

This group of notes synthesizes the concluding sections of the DOM Standard document, specifically focusing on **Section 11: Historical**. This section aggregates extensive browser compatibility data for a wide range of DOM interfaces and methods previously detailed in the main standard. The coverage includes traversal APIs (`TreeWalker`, `NodeIterator`), query languages (`XPath`, `XSLT`), legacy collection types (`NodeList`, `HTMLCollection`), Shadow DOM properties, and various text/range manipulation methods.

The data serves to map the implementation timeline of these features across major browser engines (Firefox, Safari, Chrome, Opera, Edge) and their mobile counterparts (Android WebView, iOS Safari, Samsung Internet). A significant portion of this group addresses the dichotomy between modern Chromium-based browsers and legacy environments (Internet Explorer, Edge Legacy), as well as the varying levels of support in mobile WebViews compared to desktop versions.

# Cross-Chunk Summary

The progression from Chunk 31 to Chunk 32 shifts focus from general DOM interfaces and Range/Text manipulation details toward advanced query and transformation APIs.

*   **Chunk 31** covers broad compatibility tables for core interfaces like `NodeList`, `Range`, `ShadowRoot`, and `Text`. It highlights early adoption (e.g., Firefox 1, Chrome 1) for basic features but notes specific version requirements for Shadow DOM properties (Firefox 63+, Chrome 53+). It explicitly distinguishes between "Edge" and "Edge (Legacy)," marking the latter as lacking support for modern standards.
*   **Chunk 32** narrows the scope to advanced traversal (`TreeWalker`), query languages (`XPath`, `XSLT`), and specific Shadow DOM attributes (`Element.slot`). It reinforces that while `TreeWalker` is universally supported in current engines, `XPath` and `XSLT` are absent or limited in legacy Edge and IE.

Together, these chunks provide a comprehensive historical map of the DOM API ecosystem, identifying which features are stable across all modern browsers and which remain fragmented due to legacy engine constraints or mobile WebView implementations.

# Repeated Or Central Claims

*   **Universal Modern Support:** Core interfaces such as `TreeWalker` and basic properties like `Element.slot` (Shadow DOM v1) are supported in all current engines, though specific methods within these interfaces may have higher minimum version requirements (e.g., Firefox 92 for certain Shadow DOM features).
*   **Legacy Engine Fragmentation:** There is a consistent distinction between "Edge" (Chromium-based) and "Edge (Legacy)" (IE-based). Legacy Edge and Internet Explorer consistently show no support ("None") or partial/unknown support ("?") for modern features like XPath, XSLT, Shadow DOM properties, and advanced Range methods.
*   **Mobile WebView Variability:** Mobile WebViews (Firefox for Android, iOS Safari, Chrome for Android) frequently display missing data ("?") or specific version gaps compared to their desktop counterparts. Samsung Internet is also noted with varying support levels depending on the underlying WebView version.
*   **Early Adoption vs. Removal:** While many features are listed as supported in "all current engines," the inclusion of very early version numbers (e.g., Firefox 1, Safari 1) implies long-term stability for core APIs, whereas specific methods like `Range.deleteContents` show historical support ranges that may imply deprecation or removal in later versions.

# Important Local Details

*   **TreeWalker Interface:**
    *   **Properties/Methods:** `currentNode`, `filter`, `firstChild`, `lastChild`, `nextNode`, `nextSibling`, `parentNode`, `previousNode`, `previousSibling`, `root`, `whatToShow`.
    *   **Usage:** Instantiated with a root node, optional filter function, and `whatToShow` flag.
*   **XPath API:**
    *   **Interfaces:** `XPathEvaluator`, `XPathExpression`, `XPathResult`.
    *   **Properties:** `booleanValue`, `numberValue`, `stringValue`, `singleNodeValue`, `snapshotItem`, `snapshotLength`.
    *   **Usage:** `XPathEvaluator.evaluate()` returns an `XPathResult` object for querying.
*   **XSLT API:**
    *   **Interface:** `XSLTProcessor`.
    *   **Methods:** `clearParameters`, `getParameter`, `importStylesheet`, `removeParameter`, `reset`, `setParameter`, `transformToDocument`, `transformToFragment`.
*   **Shadow DOM & Slots:**
    *   **Property:** `Element.slot` (indicates Shadow DOM v1 support).
    *   **History:** Supported from Firefox 63, Safari 10.1, Chrome 53.
    *   **Usage:** Assign a slot name to an element in a Shadow Root (e.g., `element.slot = "name"`).
*   **Range & Text APIs:**
    *   **Methods:** `toString`, `deleteContents`, `extractContents`.
    *   **Note:** `deleteContents` had specific Firefox version ranges (1–15) in historical data, suggesting potential removal or change.
    *   **Properties:** `Text.wholeText`.
*   **Legacy Collections:**
    *   `NodeList` and `HTMLCollection` are identified as "Old-style collections."
    *   `NodeList/forEach` supported in Firefox 50+, Safari 10+, Chrome 51+.

# Candidate Wiki Hints

*   **DOM API Compatibility Tables:** A summary page listing support status for core DOM interfaces across browsers.
*   **TreeWalker API Reference:** Documentation covering properties and filtering examples for `TreeWalker`.
*   **XPath in JavaScript:** A guide on using `XPathEvaluator` and `XPathResult` for DOM queries.
*   **XSLT Transformation Guide:** Instructions for using `XSLTProcessor` to transform XML/HTML documents.
*   **Shadow DOM Slots:** An article covering the `slot` property, light DOM distribution, and browser compatibility.
*   **Legacy Browser Support:** A section detailing the differences between modern Edge and Edge Legacy, as well as IE limitations.

# Gaps Or Cautions

*   **Incomplete Mobile Data:** Many entries for mobile WebViews (Firefox for Android, iOS Safari, Chrome for Android) are marked with "?", indicating incomplete tracking or unknown status compared to desktop browsers.
*   **Edge Discrepancy:** Care must be taken when referencing "Edge" compatibility; data often conflates the modern Chromium version with the legacy IE-based version, which lacks significant support for newer standards.
*   **Historical Context:** The presence of very early version numbers (e.g., Firefox 1) alongside notes on method removal or change implies that some APIs have evolved significantly or been deprecated since their initial introduction.
*   **Data Consistency:** The source distinguishes between "supported" (✔MDN) and "unknown/unlisted" (?), suggesting that the absence of a checkmark does not necessarily mean unsupported, but rather unverified in the context of the data compilation.
