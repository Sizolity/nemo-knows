---
title: Chunk 26 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 11. Historical**, specifically detailing browser compatibility tables for various DOM API methods and properties within the `Document`, `DOMImplementation`, `DOMTokenList`, `Element`, `DocumentFragment`, `CharacterData`, `Comment`, `CustomEvent`, and related interfaces. The data presents version numbers for browsers including Firefox, Safari, Chrome, Opera, Edge (Legacy), Edge, and various mobile variants (Android WebView, Samsung Internet, etc.).

# Local Summary

The section provides a comprehensive compatibility matrix for DOM APIs. It lists specific methods such as `DOMImplementation.createDocument`, `Element.appendChild` (implied via similar patterns in other sections not fully shown but referenced in logic), `DOMTokenList.add`, and properties like `Document.characterSet`. The tables indicate the minimum browser version required to support these features, noting "In all current engines" for widely supported methods and providing specific legacy versions (e.g., IE6+, Edge Legacy) for older implementations. Some entries mark certain platforms (like iOS Safari or Android WebView) with question marks or empty checks where data is unavailable.

# Key Claims

-   **Universal Support**: Many core DOM APIs (e.g., `DOMImplementation`, `CharacterData`, `Comment` in specific contexts, `Element.appendChild`) are marked as supported "In all current engines" across Firefox, Safari, Chrome, Opera, and Edge.
-   **Legacy Browser Support**: Specific versions for legacy browsers are listed, such as `Edge (Legacy) 12+` or `IE6+`, indicating when these APIs became available in older environments.
-   **Mobile Browser Variance**: Mobile browsers often have different version numbers than their desktop counterparts (e.g., Firefox for Android vs. standard Firefox). Some mobile platforms like iOS Safari or specific Android WebViews are sometimes marked with question marks, suggesting uncertain or unverified support data in the source text.
-   **Version Specificity**: Support is quantified by browser version (e.g., `Firefox 49+`, `Chrome 54+`), highlighting when a feature was introduced or became stable enough for general use.

# Entities And Concepts

-   **DOM APIs**: Methods and properties belonging to the Document Object Model (e.g., `createDocument`, `appendChild`, `removeChild`).
-   **Browser Engines**: The underlying rendering engines of web browsers (Gecko, WebKit, Blink).
-   **Legacy Browsers**: Older versions of browsers like Internet Explorer (IE) and early Edge releases.
-   **Mobile Browsers**: Browser variants designed for mobile devices (Android WebView, iOS Safari, Samsung Internet).
-   **Compatibility Matrix**: A structured table showing feature support across different browser versions.

# Procedures And API Details

-   **`DOMImplementation.createDocument`**: Supported in all current engines; legacy Edge 12+, IE9+.
-   **`DOMImplementation.createDocumentType`**: Supported in all current engines; legacy Edge 12+, IE9+.
-   **`DOMTokenList.add` / `remove` / `toggle`**: Added around Firefox 3.6, Safari 5.1, Chrome 8.
-   **`DOMTokenList.replace`**: Added later, specifically Firefox 49+, Safari 10.1, Chrome 61+.
-   **`Document.createAttributeNS`**: Supported in all current engines; legacy Edge 12+, IE9+.
-   **`Document.createElement`**: Supported in all current engines; legacy Edge 12+, IE5+.

# Nuance Or Contradictions

-   **Missing Data Indicators**: Question marks (e.g., `Firefox for Android?`, `iOS Safari?`) indicate that the source text does not have definitive compatibility data for these specific platforms, contrasting with the "In all current engines" claim for others.
-   **Legacy vs. Modern Edge**: Distinction is made between "Edge (Legacy)" and modern "Edge", sometimes showing different version requirements or noting "IENone" (Internet Explorer None) for features not supported in IE at all.
-   **Version Gaps**: Some APIs appear in later versions of major browsers (e.g., `DOMTokenList.supports` in Firefox 49+) compared to others that support them from earlier versions, reflecting incremental feature adoption.

# Candidate Wiki Hints

1.  **DOM Compatibility Tables**: A dedicated page or section explaining how to read MDN compatibility tables and interpreting version numbers.
2.  **Legacy Browser Support**: Documentation on writing for legacy browsers like Internet Explorer and early Edge, focusing on fallback strategies for APIs listed here.
3.  **DOMTokenList API Reference**: A detailed guide covering `add`, `remove`, `toggle`, and other methods with historical adoption timelines.
