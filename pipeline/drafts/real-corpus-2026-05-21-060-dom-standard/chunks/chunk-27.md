---
title: Chunk 27 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

**Heading:** 11. Historical
**Source Path:** `raw/web/corpus-2026-05-18/060-dom-standard.md`
**Line Range:** 10726–11432
**Content Scope:** Compatibility tables for DOM Level 2 and Level 3 interfaces, specifically focusing on the `Document`, `Element`, `DocumentFragment`, and `XPathEvaluator` objects. The chunk lists methods such as `createEvent`, `querySelector`, `getElementsByClassName`, and properties like `documentElement`. It details support across desktop browsers (Firefox, Safari, Chrome, Opera, Edge) and legacy versions (IE), alongside mobile implementations (Android WebView, Samsung Internet).

# Local Summary

This section of the historical DOM standard documentation provides a comprehensive compatibility matrix for various Document and Element interface methods. It enumerates specific browser engine versions required to support features like `querySelector`, `prepend`, `replaceChildren`, and XPath expression creation (`createExpression`). The data distinguishes between modern engines (Firefox, Safari, Chrome) and legacy environments (Edge Legacy, Internet Explorer), often noting missing support in mobile browsers or older Android WebViews for newer APIs.

# Key Claims

- **Universal Support:** Core DOM Level 2 methods like `getElementsByClassName`, `createElement`, and `documentElement` are supported in "all current engines" starting from version 1 of the respective Firefox, Safari, and Chrome builds listed.
- **Modern API Adoption:** Features such as `prepend`, `replaceChildren`, and `firstElementChild` require significantly higher version numbers (e.g., Firefox 49+, Chrome 54+, Safari 10+).
- **Legacy Limitations:** Internet Explorer lacks support for many modern DOM methods, often marked as "None" or requiring very early versions (IE 5–6) for basic properties. Edge Legacy supports some features but not others like `replaceChildren`.
- **Mobile Disparities:** Mobile browsers (Android WebView, iOS Safari) frequently show question marks (?) indicating unknown status or gaps in support for newer APIs like `prepend` and `replaceChildren` at the time of documentation.

# Entities And Concepts

- **Document Interface:** The root object representing an HTML document in the DOM tree.
- **Element Interface:** Represents an element within a document.
- **DocumentFragment:** A lightweight container used to optimize DOM manipulation by creating a temporary node structure.
- **XPathEvaluator:** An interface for evaluating XPath expressions (e.g., `createExpression`, `evaluate`).
- **Browser Engines:** Firefox, Safari, Chrome, Opera, Edge (Modern and Legacy), Internet Explorer.
- **DOM Methods:** `querySelector`, `querySelectorAll`, `getElementsByClassName`, `getElementById`.

# Procedures And API Details

The following methods and properties are detailed with specific version requirements:

| Method/Property | Interface | Firefox | Safari | Chrome | Opera | Edge (Mod) | Edge (Leg) | IE |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| `createEvent` | Document | 1+ | 1+ | 1+ | 7+ | 79+ | 12+ | 5+ |
| `createExpression` | Document/XPathEval | 1+ | 3+ | 1+ | 12.1+ | 79+ | 12+ | None |
| `createNodeIterator` | Document | 1+ | 3+ | 1+ | 9+ | 79+ | 12+ | 5+ |
| `createNSResolver` | Document/XPathEval | 1+ | 3+ | 1+ | 12.1+ | 79+ | 12+ | None |
| `createProcessingInstruction` | Document | 1+ | 1+ | 1+ | 12.1+ | 79+ | 12+ | 9+ |
| `createRange` | Document | 1+ | 1+ | 1+ | 12.1+ | 79+ | 12+ | 9+ |
| `createTextNode` | Document | 1+ | 1+ | 1+ | 7+ | 79+ | 12+ | 5+ |
| `createTreeWalker` | Document | 1+ | 3+ | 1+ | 9+ | 79+ | 12+ | 9+ |
| `doctype` | Document | 1+ | 1+ | 1+ | 12.1+ | 79+ | 12+ | 6+ |
| `documentElement` | Document | 1+ | 1+ | 1+ | 7+ | 79+ | 12+ | 5+ |
| `documentURI` | Document | 1+ | 3+ | 1+ | 12.1+ | 79+ | 12+ | None |
| `evaluate` | Document/XPathEval | 1+ | 3+ | 1+ | 9+ | 79+ | 12+ | None |
| `firstElementChild` | Element/DocFragment | 25+ | 9+ | 29+ | ? | 79+ | 17+ | None |
| `getElementsByClassName` | Document | 3+ | 3.1+ | 1+ | 9.5+ | 79+ | 12+ | 9+ |
| `querySelector` | Element/DocFragment | 3.5+ | 3.1+ | 1+ | 10+ | 79+ | 12+ | 9+ |
| `replaceChildren` | Element/DocFragment | 78+ | 14+ | 86+ | ? | 86+ | ? | None |

# Nuance Or Contradictions

- **Version Gaps:** There are significant gaps between the support of legacy properties (e.g., `getElementById` in Firefox 1+) and modern manipulation methods (e.g., `replaceChildren` in Firefox 78+).
- **Mobile Uncertainty:** Many mobile browsers are marked with "?" for newer APIs, suggesting inconsistent implementation or lack of data compared to desktop counterparts.
- **XPath Support:** XPath expression creation (`createExpression`) is widely supported but lacks support in Internet Explorer entirely, contrasting with basic DOM properties which have minimal IE support.

# Candidate Wiki Hints

- **DOM Compatibility Matrix:** A dedicated page summarizing browser support for DOM Level 2 and 3 methods.
- **Modern DOM Methods:** Documentation focusing on newer APIs like `prepend`, `replaceChildren`, and `firstElementChild`.
- **XPath in the Browser:** A guide covering `XPathEvaluator` interfaces (`evaluate`, `createExpression`).
