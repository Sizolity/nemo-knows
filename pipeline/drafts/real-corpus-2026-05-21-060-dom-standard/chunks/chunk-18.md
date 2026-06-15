---
title: Chunk 18 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

**Heading Path:** 6. Traversal > 6.1. Interface NodeIterator, 6.2. Interface TreeWalker, 6.3. Interface NodeFilter, 7. Sets, 7.1. Interface DOMTokenList, 8. XPath
**Source File:** `raw/web/corpus-2026-05-18/060-dom-standard.md`
**Line Range:** 6288–6784

This chunk details the DOM traversal interfaces (`NodeIterator`, `TreeWalker`, `NodeFilter`) and the set-like interface `DOMTokenList`. It concludes with a note on the legacy XPath API.

# Local Summary

The document defines algorithms and IDL for tree traversal mechanisms. `NodeIterator` provides forward/backward traversal from a specific root with a filter, while `TreeWalker` maintains a current node position to traverse descendants, ancestors, siblings, or children based on filters. `NodeFilter` supplies the filtering logic (`acceptNode`) and bitmask constants (`whatToShow`). The chunk also covers `DOMTokenList`, an interface for managing sets of tokens (like class names), including methods to add, remove, toggle, and replace tokens with specific error handling for whitespace and empty strings. Finally, it notes that DOM Level 3 XPath APIs are legacy but maintained in the spec for future updates.

# Key Claims

- `NodeIterator` is created via `createNodeIterator()` on a `Document`.
- `TreeWalker` is created via `createTreeWalker()` on a `Document`.
- Both iterators use a filter (`NodeFilter`) and a `whatToShow` bitmask to determine node visibility.
- The `detach()` method on `NodeIterator` exists for compatibility but performs no action; its functionality was removed.
- `DOMTokenList` methods throw `SyntaxError` for empty strings and `InvalidCharacterError` for tokens containing ASCII whitespace.
- XPath Level 3 APIs are considered legacy and not actively maintained, though definitions remain for potential future updates.

# Entities And Concepts

- **Interfaces:** `NodeIterator`, `TreeWalker`, `NodeFilter`, `DOMTokenList`.
- **Constants:** `FILTER_ACCEPT` (1), `FILTER_REJECT` (2), `FILTER_SKIP` (3), `SHOW_ALL`, `SHOW_ELEMENT`, `SHOW_TEXT`, etc.
- **Exceptions:** `SyntaxError` DOMException, `InvalidCharacterError` DOMException, `TypeError`.
- **Attributes/Properties:** `root`, `referenceNode`, `pointerBeforeReferenceNode`, `whatToShow`, `filter`, `currentNode`, `length`, `value`.
- **Methods:** `nextNode()`, `previousNode()`, `parentNode()`, `firstChild()`, `lastChild()`, `add()`, `remove()`, `toggle()`, `replace()`.

# Procedures And API Details

### NodeIterator Traversal Logic
To traverse, the algorithm iterates while true:
1. Determine direction (`next` or `previous`).
2. If moving forward and `pointerBeforeReferenceNode` is false, get the first following node. If true, move it to false and get the next node.
3. Filter the candidate node.
4. If filtered as `FILTER_ACCEPT`, break the loop.
5. Update `referenceNode` and return the node.

### TreeWalker Traversal Logic
- **Children:** Traverse first/last child. Skip children if filtered, moving to the next/previous sibling of the parent if necessary.
- **Siblings:** Move to next/previous sibling. If filtered, descend into children or move to the next sibling of the current node.
- **Ancestors:** Walk up the tree until a filter accepts a node or the root is reached.

### DOMTokenList Operations
- **add(tokens...):** Appends tokens if not present. Throws `SyntaxError` for empty strings, `InvalidCharacterError` for whitespace.
- **remove(tokens...):** Removes tokens if present. Same error handling as `add`.
- **toggle(token, force):** If `force` is omitted/false and token exists, removes it (returns false). If `force` is true/omitted and token does not exist, adds it (returns true).
- **replace(token, newToken):** Swaps tokens. Returns false if token doesn't exist; throws errors for empty/whitespace strings.

# Nuance Or Contradictions

- **detach() Methodality:** The `detach()` method on `NodeIterator` is preserved strictly for compatibility but currently does nothing, as its original functionality was removed from the spec logic.
- **DOMTokenList Name:** The text explicitly notes that the name "DOMTokenList" is an "unfortunate legacy mishap," suggesting a potential future rename or refactoring in related specifications (like HTML).
- **Update Steps Execution:** The standard update steps for `DOMTokenList` are not always executed for `toggle()` and `replace()` methods to ensure web compatibility, implying implementation-specific behaviors may differ from strict algorithmic definitions.
- **XPath Status:** While the XPath interface definitions are maintained in the DOM spec, the underlying functionality (DOM Level 3 XPath) is widely implemented but considered unmaintained and legacy.

# Candidate Wiki Hints

- **NodeIterator vs TreeWalker:** Create a comparison page explaining that `NodeIterator` traverses from a fixed root without changing its current position context during iteration, whereas `TreeWalker` maintains a mutable `currentNode` state allowing navigation to parents, children, and siblings dynamically.
- **DOMTokenList Token Validation:** Document the strict validation rules for `DOMTokenList`, specifically the prohibition of whitespace characters and empty strings, which distinguishes it from standard string lists.
- **Traversal Filter Constants:** Create a reference table for `NodeFilter` constants (`FILTER_ACCEPT`, `SHOW_ELEMENT`, etc.) and explain how bitwise operations combine `whatToShow` flags.
