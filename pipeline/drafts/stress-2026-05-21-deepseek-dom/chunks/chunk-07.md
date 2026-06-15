---
title: Chunk 07 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- **Source:** `raw/web/corpus-2026-05-18/060-dom-standard.md`
- **Chunk:** 7 of 9, lines 8818–11166
- **Heading path:** `11. Historical`
- **Coverage:** Full contents of the `Historical` section, combining WebIDL interface definitions with appended MDN browser‑compatibility tables.

## Local Summary
The chunk presents the DOM Standard’s “Historical” chapter. It defines legacy interfaces (`CDATASection`, `ProcessingInstruction`, `Comment`), the range and selection infrastructure (`AbstractRange`, `StaticRange`, `Range`), DOM tree traversal (`NodeIterator`, `TreeWalker`, `NodeFilter`), token list manipulation (`DOMTokenList`), and the XPath/XSLT API surface. Following the normative IDL, the chunk includes extensive MDN compatibility data for these interfaces as well as for `AbortController`/`AbortSignal`, `Attr`, `CharacterData`, `Document`, `Element`, `CustomEvent`, and `DOMImplementation`.

## Key Claims
- `CDATASection` is a historical interface (derives from `Text`).
- `ProcessingInstruction` and `Comment` are `CharacterData` subtypes.
- `AbstractRange` provides the read‑only basis for `StaticRange` (constructor with `StaticRangeInit`) and `Range` (mutable, with `commonAncestorContainer` and mutation methods).
- `Range` offers two‑point boundary manipulation, collapse, extraction/deletion/cloning/surrounding of contents, and a `detach()` method (historical).
- `NodeIterator` and `TreeWalker` allow filtered traversal; `NodeFilter` defines acceptance constants and `whatToShow` bitmask values, with several constants flagged as legacy (`SHOW_ENTITY_REFERENCE`, `SHOW_ENTITY`, `SHOW_NOTATION`).
- `DOMTokenList` represents a set of space‑separated tokens with methods (`add`, `remove`, `toggle`, `replace`, `supports`) and a stringifier `value`.
- `XPathResult` holds typed XPath evaluation results; `XPathExpression` evaluates an expression against a context node.
- `XSLTProcessor` supports importing a stylesheet and transforming a source document into a fragment or document.
- The appended compatibility tables indicate wide engine support for most DOM interfaces, with some features (e.g., `AbortSignal.timeout()`, `DOMTokenList.replace`, `AbstractRange` directly) having later adoption.

## Entities And Concepts
- **Historical interfaces:** `CDATASection`, `ProcessingInstruction`, `Comment`
- **Range model:** `AbstractRange`, `StaticRangeInit`, `StaticRange`, `Range`
- **Traversal:** `NodeIterator`, `TreeWalker`, `NodeFilter` (callback interface)
- **Token list:** `DOMTokenList`
- **XPath:** `XPathResult`, `XPathExpression`, `XPathNSResolver`, `XPathEvaluatorBase` mixin, `XPathEvaluator`
- **XSLT:** `XSLTProcessor`
- **Compatibility data:** Tables for `AbortController`, `AbortSignal`, `AbstractRange`/`Range`/`StaticRange` properties, `Attr`, `CharacterData` methods, `Comment`, `CustomEvent`, `DOMImplementation`, `DOMTokenList`, `Document` methods and properties, `XPathEvaluator`, `NodeIterator`/`TreeWalker` creation, etc.

## Procedures And API Details
- `Range` mutation: `setStart`/`setEnd`, `setStartBefore`/`setStartAfter`/`setEndBefore`/`setEndAfter`, `collapse`, `selectNode`, `selectNodeContents`, `deleteContents`, `extractContents`, `cloneContents`, `insertNode`, `surroundContents`. Boundary comparison constants `START_TO_START` (0), `START_TO_END` (1), `END_TO_END` (2), `END_TO_START` (3). `detach()` does nothing (legacy).
- `NodeIterator`/`TreeWalker`: filter using `NodeFilter.acceptNode(node)`, `whatToShow` bitmask. `TreeWalker` exposes `currentNode` and directional walk methods.
- `DOMTokenList`: `toggle(token, force)` returns whether token is present after toggling; `replace(old, new)` replaces a token; `supports(token)` checks if token is valid according to the associated attribute’s definition.
- `XPathResult`: result types defined as constants (`ANY_TYPE`, `NUMBER_TYPE`, …). `snapshotItem` and `iterateNext` for snapshots/iterators.
- `XSLTProcessor`: `importStylesheet(style)`, `transformToFragment(source, outputDocument)`, `transformToDocument(source)`, parameter management methods.

## Nuance Or Contradictions
- Several interfaces marked as historical or legacy (`CDATASection`, `detach()` methods, `SHOW_ENTITY_REFERENCE` etc.), yet they remain in the standard for reference.
- `XPathEvaluatorBase` is a mixin that both `Document` and `XPathEvaluator` include, allowing `evaluate` to be called directly on a document (legacy `createNSResolver` also available).
- Compatibility tables show near‑universal support for many interfaces, but some (like `AbortSignal.timeout()` static, `AbstractRange` independently instantiated) are noted with later browser versions.
- The source mixes formal IDL and informal MDN data; the MDN inclusion may not be part of the official DOM Standard, but it is present in this corpus dump.

## Candidate Wiki Hints
- “DOM Range Interface” — covers mutable selections, boundaries, manipulation.
- “NodeIterator and TreeWalker” — DOM tree traversal and filtering.
- “DOMTokenList” — managing space‑separated tokens (classList, relList, etc.).
- “XPath in the DOM” — evaluating XPath expressions and result types.
- “XSLTProcessor” — client‑side XSLT transformations.
- “Historical DOM Interfaces” — aliases and legacy types still defined for web compatibility.
- “DOM Living Standard — Historical chapter reference” — index of all interfaces in this section.
