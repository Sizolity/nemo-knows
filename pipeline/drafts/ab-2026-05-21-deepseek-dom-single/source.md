---
title: DOM Standard
kind: source
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# DOM Standard

## What It Is

The **DOM Standard** is the WHATWG Living Standard that defines a platform-neutral model for the Document Object Model. It specifies events, aborting activities, node trees, ranges, traversal, and related APIs. The canonical URL is `https://dom.spec.whatwg.org/`.

## Summary

The standard provides the foundational data model and API surface for interacting with structured documents (HTML and XML) on the web platform. It is organized into these major areas:

- **Infrastructure** – Defines trees, tree order, ordered sets, selector scoping, and name validation rules for elements, attributes, doctypes, and namespace prefixes.
- **Events** – Specifies the `Event` and `CustomEvent` interfaces, `EventTarget`, event listener registration (`addEventListener`/`removeEventListener` with `AbortSignal` support), event phases (capture, target, bubble), dispatching, firing, and the distinction between actions and occurrences.
- **Aborting Activities** – Introduces `AbortController` and `AbortSignal` as a general mechanism to abort ongoing operations, with methods like `AbortSignal.abort()`, `AbortSignal.timeout()`, `AbortSignal.any()`, and `throwIfAborted()`.
- **Nodes** – Defines the node tree model: `Node`, `Document`, `DocumentType`, `DocumentFragment`, `ShadowRoot`, `Element`, `Attr`, `CharacterData`, `Text`, `CDATASection`, `ProcessingInstruction`, and `Comment`. Covers mutation algorithms (insert, remove, move, replace), shadow trees with slots and slottables, mutation observers (`MutationObserver`/`MutationRecord`), old-style collections (`NodeList`, `HTMLCollection`), and mixins (`ParentNode`, `ChildNode`, `NonElementParentNode`, `DocumentOrShadowRoot`, `Slottable`).
- **Ranges** – Defines `AbstractRange`, `StaticRange` (lightweight, non-live), and `Range` (live) for representing and manipulating contiguous content between two boundary points.
- **Traversal** – Specifies `NodeIterator` and `TreeWalker` with `NodeFilter` for filtered traversal of node trees.
- **Sets** – Defines `DOMTokenList` for managing ordered sets of whitespace-separated tokens (used by `classList` and similar).
- **XPath / XSLT** – Preserves legacy XPath 1.0 and XSLT 1.0 API definitions (`XPathEvaluator`, `XPathExpression`, `XPathResult`, `XSLTProcessor`) for compatibility.
- **Historical** – Documents many removed interfaces and members (e.g., `MutationEvent`, `Entity`, `DOMConfiguration`, `isSupported`, `expandEntityReferences`).

## Key Claims

1. **Events signal occurrences, not actions.** Events represent notifications from algorithms and may influence ongoing operations (via `preventDefault()`), but must not be used as initiators.
2. **`AbortController`/`AbortSignal` provide a unified abort mechanism** for promise-based and event-listener-based APIs, with support for composition (`AbortSignal.any()`) and timeouts (`AbortSignal.timeout()`).
3. **Shadow DOM enables encapsulation** through shadow trees attached to hosts, with slot-based projection (`manual` or `named` assignment) and retargeting of events across shadow boundaries.
4. **Node tree mutations follow a carefully sequenced algorithm** with insertion steps, removing steps, moving steps, post-connection steps, and children-changed steps, ensuring atomic batch operations and proper custom element lifecycle callbacks.
5. **Mutation observers (`MutationObserver`) provide asynchronous, batched reporting** of tree changes via a microtask-based notification system, replacing the deprecated synchronous `MutationEvent`.
6. **Live ranges (`Range`) are automatically maintained** across node tree mutations, while `StaticRange` offers a lightweight non-live alternative.
7. **Validation rules for names have been deliberately loosened** compared to historical XML-aligned rules, to match what the HTML parser can produce and avoid developer friction.
8. **There are no known security or privacy considerations** specific to this standard.

## Suggested Links

- **DOM Standard** – `https://dom.spec.whatwg.org/`
- **HTML Standard** – `https://html.spec.whatwg.org/multipage/`
- **Infra Standard** – `https://infra.spec.whatwg.org/`
- **Web IDL Standard** – `https://webidl.spec.whatwg.org/`
- **UI Events** – `https://w3c.github.io/uievents/`
- **Touch Events** – `https://w3c.github.io/touch-events/`
- **Selectors Level 4** – `https://drafts.csswg.org/selectors/`
- **Trusted Types** – `https://w3c.github.io/trusted-types/dist/spec/`
- **Fullscreen API Standard** – `https://fullscreen.spec.whatwg.org/`
- **DOM Parsing and Serialization** – `https://w3c.github.io/DOM-Parsing/`
- **CSSOM View Module** – `https://drafts.csswg.org/cssom-view/`
