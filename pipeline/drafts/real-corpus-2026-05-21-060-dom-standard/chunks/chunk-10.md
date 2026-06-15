---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
The source text defines algorithms for locating namespaces and prefixes on DOM nodes (Interface Node), describes behavior for `HTMLCollection` lists of elements based on qualified names, namespaces, and local names, and specifies filtering logic for class names. It covers steps for `Element`, `Document`, `Attr`, and other node types, including fallback to parent elements.

Local Summary
This chunk details the logic for resolving namespace prefixes and URIs across different DOM interfaces (`Element`, `Document`, `Attr`, etc.) by checking specific attributes, namespaces, or delegating to a parent element. It also outlines how to construct `HTMLCollection` objects that filter descendant elements based on qualified names, namespace/local name combinations, and class names, with specific handling for HTML documents and quirks mode.

Key Claims
- Namespace prefix resolution involves checking the element's own namespace, an "xmlns" attribute, or delegating to the parent element.
- `lookupNamespaceURI` converts an empty string prefix to null before resolving the namespace.
- `isDefaultNamespace` compares a given namespace against the result of locating a namespace with a null prefix.
- DOM mutation methods (`insertBefore`, `appendChild`, `replaceChild`, `removeChild`) delegate to underlying pre-insertion/pre-removal logic.
- `HTMLCollection` filters for qualified names distinguish between HTML namespace elements (case-insensitive matching) and non-HTML namespace elements.
- Class name filtering requires all specified classes to be present on an element, with ASCII case-insensitive comparison in "quirks" mode.

Entities And Concepts
- Interface Node: The interface defining methods for locating namespaces and prefixes.
- `locate a namespace prefix`: Algorithm to find the prefix string for a given namespace URI.
- `locate a namespace`: Algorithm to find the namespace URI for a given prefix string.
- `lookupPrefix`: Method returning the prefix for a specific namespace.
- `lookupNamespaceURI`: Method returning the namespace URI for a specific prefix.
- `isDefaultNamespace`: Method checking if a namespace is the default (empty) namespace.
- `HTMLCollection`: A collection object returned by element listing algorithms.
- Qualified Name: Combination of namespace and local name used to identify elements.

Procedures And API Details
- **Algorithm for locating namespace prefix:**
  1. Return the element's prefix if namespace matches and prefix is non-null.
  2. Return the local name of an attribute with namespace "xmlns" and value matching namespace.
  3. Recursively call on the parent element if it exists.
  4. Return null otherwise.
- **Algorithm for locating namespace URI:**
  1. Set empty string prefix to null.
  2. Check specific prefixes ("xml", "xmlns") or attributes with local name matching the prefix.
  3. Return the namespace value if non-empty, or null otherwise.
  4. Delegate to parent element if current node has no parent.
- **Algorithm for `HTMLCollection` by qualified name:**
  1. If argument is "*", return all descendant elements.
  2. For HTML documents, filter descendants where namespace is HTML (ASCII lowercase) or not HTML matching the qualified name.
  3. Otherwise, match any element with the specified qualified name.
- **Algorithm for `HTMLCollection` by namespace and local name:**
  1. Convert empty strings to null.
  2. Handle wildcard (*) for namespace or local name individually or together.
  3. Return a collection filtering descendants matching the specific namespace/local name combination.
- **Algorithm for `HTMLCollection` by class names:**
  1. Parse class names into an ordered set.
  2. If empty, return empty collection.
  3. Return collection where elements have all classes in the set; comparisons are ASCII case-insensitive in "quirks" mode.

Nuance Or Contradictions
- The logic for locating a namespace prefix prioritizes the element's own namespace attribute over inherited context from parents unless specific conditions (like an "xmlns" attribute) are met.
- `HTMLCollection` caching is mentioned: the same object may be returned for identical arguments as long as the document type hasn't changed.
- Case sensitivity for qualified name matching depends on whether the document is an HTML document and the specific namespace involved (HTML namespace requires ASCII lowercase comparison).

Candidate Wiki Hints
- DOM Namespace Resolution Algorithms
- Interface Node Methods (`lookupPrefix`, `lookupNamespaceURI`)
- HTMLCollection Filtering Strategies
- Handling Class Names in DOM Queries
