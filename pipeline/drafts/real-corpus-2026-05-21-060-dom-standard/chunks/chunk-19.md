---
title: Chunk 19 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk covers the DOM Level 3 XPath and XSLT interfaces, detailing `XPathResult`, `XPathExpression`, `XPathEvaluator`, and `XSLTProcessor`. It also notes that no known security or privacy considerations exist for this standard. The coverage spans sections 8.1 through 10.

## Local Summary
The specification defines the `XPathResult` interface with various result types (e.g., `ANY_TYPE`, `STRING_TYPE`, `UNORDERED_NODE_ITERATOR_TYPE`) and attributes for accessing values (`numberValue`, `stringValue`, `singleNodeValue`). The `XPathEvaluatorBase` mixin provides methods like `createExpression` and `evaluate`, noting that `Document` includes this mixin for historical reasons. The `XSLTProcessor` interface is defined for transforming XML documents, with methods such as `importStylesheet`, `transformToFragment`, and parameter management (`setParameter`, `clearParameters`).

## Key Claims
- The `XPathResult` interface supports multiple result types via constants like `ANY_TYPE = 0` through `FIRST_ORDERED_NODE_TYPE = 9`.
- The `createNSResolver(nodeResolver)` method is provided only for historical reasons and simply returns the provided node resolver.
- Both `XPathEvaluator` and `Document` instances allow access to XPath evaluation methods due to `Document` including `XPathEvaluatorBase`.
- The `XSLTProcessor` interface includes a constructor and methods for importing stylesheets, transforming documents/fragments, and managing parameters.
- No known security or privacy considerations are identified for the XSLT and XPath APIs in this standard.

## Entities And Concepts
- **XPathResult**: An interface representing the result of an XPath evaluation.
- **XPathExpression**: An interface representing a compiled XPath expression.
- **XPathEvaluatorBase**: A mixin providing core XPath evaluation functionality (`createExpression`, `evaluate`).
- **XPathEvaluator**: An interface that includes `XPathEvaluatorBase`.
- **XSLTProcessor**: An interface for performing XSLT transformations on XML documents.
- **XPathNSResolver**: A callback interface used to resolve namespace URIs during expression creation.

## Procedures And API Details
- **Creating an XPath Expression**: Use `XPathEvaluatorBase.createExpression(expression, resolver)` where the expression is a DOMString and resolver is optional.
- **Evaluating an XPath Expression**: Call `XPathResult.evaluate(expression, contextNode, resolver, type, result)`. The `type` argument defaults to `0` (`ANY_TYPE`).
- **Processing XSLT**: Instantiate `XSLTProcessor`, call `importStylesheet(Node)` to load stylesheets, then use `transformToFragment(source)` or `transformToDocument(source)`. Parameters are set via `setParameter(namespaceURI, localName, value)`.

## Nuance Or Contradictions
- The `createNSResolver` method is explicitly marked as existing only for historical reasons and acts as a pass-through (returns the input node).
- Historically, `XPathEvaluator` could be constructed directly; however, since `Document` includes `XPathEvaluatorBase`, one can also access XPath methods directly on a document object.

## Candidate Wiki Hints
- **Page**: XPath in DOM / XSLT Processing
  - **Content**: Overview of the XPath and XSLT interfaces, result types, and common usage patterns for XML transformation.
