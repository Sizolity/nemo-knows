---
title: Chunk 32 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

Chunk Context
- **Section:** 11. Historical (DOM Compatibility Data)
- **Coverage:** Line range 14387–15051 of `060-dom-standard.md`.
- **Content Scope:** A comprehensive compatibility table listing DOM APIs, specifically focusing on the TreeWalker interface, XPath support (XPathEvaluator, XPathExpression, XPathResult), XSLT processing (XSLTProcessor), and the Element.slot property.

Local Summary
This section provides historical browser support data for various DOM Level 3 traversal and manipulation interfaces. It details version requirements for major engines (Firefox, Safari, Chrome, Opera, Edge, IE) and their mobile counterparts (Android WebView, Samsung Internet). The data confirms broad support for `TreeWalker` methods across modern browsers while noting legacy limitations in older Internet Explorer versions and specific gaps in early Edge builds regarding XPath and XSLT.

Key Claims
- **Universal Support:** The `TreeWalker` interface and its primary properties (`currentNode`, `filter`, `firstChild`, `lastChild`, `nextNode`, `nextSibling`, `parentNode`, `previousNode`, `previousSibling`, `root`, `whatToShow`) are supported in all current engines.
- **Legacy Browser Limits:** Older versions of Internet Explorer (IE5, IE9) lack support for these advanced traversal features. Legacy Edge (version 12+) shows partial or no support depending on the specific API.
- **XPath Support:** The `XPathEvaluator` and associated `XPathExpression`, `XPathResult` interfaces are present in modern engines but absent or limited in legacy Edge and IE environments.
- **XSLT Capabilities:** `XSLTProcessor` is widely supported in current versions, with `transformToDocument` and `transformToFragment` being key methods listed.
- **Shadow DOM Indicator:** The presence of the `Element.slot` property indicates support for Shadow DOM v1, available from Firefox 63, Safari 10, and Chrome 53 onwards.

Entities And Concepts
- **TreeWalker:** An interface used to traverse a tree structure (Document) using filters and specific traversal rules.
  - *Properties/Methods:* `currentNode`, `filter`, `firstChild`, `lastChild`, `nextNode`, `nextSibling`, `parentNode`, `previousNode`, `previousSibling`, `root`, `whatToShow`.
- **XPath API:** A set of interfaces for evaluating XPath expressions within the DOM.
  - *Entities:* `XPathEvaluator`, `XPathExpression`, `XPathResult` (with properties like `booleanValue`, `numberValue`, `stringValue`, `singleNodeValue`, `snapshotItem`, `snapshotLength`).
- **XSLT API:** Interfaces for applying XSLT transformations to XML documents.
  - *Entity:* `XSLTProcessor`.
  - *Methods:* `clearParameters`, `getParameter`, `importStylesheet`, `removeParameter`, `reset`, `setParameter`, `transformToDocument`, `transformToFragment`.
- **Shadow DOM:** A DOM feature allowing encapsulated sub-trees.
  - *Property:* `slot` (on Element).

Procedures And API Details
- **TreeWalker Initialization:** Developers can instantiate a `TreeWalker` by specifying a root node, a filter function (optional), and a `whatToShow` flag to define which nodes to visit during traversal.
- **XPath Evaluation:** Use `XPathEvaluator.evaluate()` with an `XPathExpression` instance to query the DOM. The result is typically wrapped in an `XPathResult` object, allowing access to results via properties like `singleNodeValue` for single-node matches or iteration methods like `iterateNext`.
- **XSLT Transformation:** Instantiate an `XSLTProcessor`, optionally load stylesheets via `importStylesheet()` or set parameters using `setParameter()`, and execute transformations using `transformToFragment()` (for DOM fragments) or `transformToDocument()` (for full documents).
- **Slot Assignment:** Assign a slot name to an element in a Shadow Root using the `slot` property (e.g., `element.slot = "name"`), enabling content distribution within shadow host elements.

Nuance Or Contradictions
- **Edge Discrepancy:** There is a noted distinction between "Edge" (Chromium-based) and "Edge (Legacy)" (IE-based). Legacy Edge often lacks XPath and XSLT support entirely, whereas modern Edge supports them.
- **Mobile Variability:** Mobile browsers like "Firefox for Android," "iOS Safari," and "Samsung Internet" show varying levels of support depending on the underlying WebView version or specific mobile engine implementation (e.g., Samsung Internet 10.1+ vs. older versions).
- **IE Limitations:** Internet Explorer is consistently marked with "?" or "None" for modern features like XPath, XSLT, and Shadow DOM, highlighting a clear divide between legacy and modern web standards support.

Candidate Wiki Hints
- **TreeWalker API Reference:** A page documenting the properties and methods of `TreeWalker`, including examples of filtering nodes.
- **XPath in JavaScript:** A guide explaining how to use `XPathEvaluator` and `XPathResult` for DOM queries.
- **XSLT Transformation Guide:** Instructions on using `XSLTProcessor` to transform XML/HTML documents.
- **Shadow DOM Slots:** An article covering the `slot` property, light DOM distribution, and compatibility across browsers.
