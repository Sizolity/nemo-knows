---
title: DOM Standard
kind: source
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## What It Is

The DOM Standard is the **Living Standard** maintained by the WHATWG. It defines a platform‑neutral model for events, aborting activities, and node trees — the core Document Object Model used across web browsers.

## Summary

The standard specifies the fundamental interfaces and algorithms for working with HTML and XML documents:

- **Events:** `Event` and `CustomEvent` interfaces, `EventTarget`, event phases (capture, target, bubble), passive listeners, composed paths, and synthetic event dispatching.
- **Aborting Activities:** `AbortController` and `AbortSignal` for signalling cancellation of operations (e.g., fetch, promises) and integrating with event listeners via the `signal` option.
- **Nodes and Trees:** `Node`, `Document`, `Element`, `Text`, `Comment`, `DocumentFragment`, `Attr`, and processing instructions. Tree order, mutation algorithms (insert, remove, replace, move), and mixins such as `ParentNode`, `ChildNode`, `Slottable`.
- **Shadow DOM:** shadow trees, slots, slottables, open/closed shadow roots, slot assignment, retargeting, and composed event paths.
- **Custom Elements:** element creation, custom element states (`"undefined"`, `"failed"`, `"uncustomized"`, `"custom"`), upgrades, and lifecycle callbacks.
- **Ranges:** `StaticRange` (lightweight snapshot) and live `Range` objects representing content between boundary points; methods to manipulate, extract, and compare ranges.
- **Traversal:** `NodeIterator` and `TreeWalker` for filtering and walking node trees.
- **Collections and Sets:** `NodeList`, `HTMLCollection`, and `DOMTokenList` (used by `classList`).
- **Mutation Observers:** `MutationObserver` and `MutationRecord` for asynchronously watching DOM changes.
- **Legacy APIs:** XPath (`XPathEvaluator`, `XPathResult`, `XPathExpression`) and XSLT (`XSLTProcessor`) maintained for compatibility.
- **Infrastructure:** tree definitions, ordered sets, selector scoping, name validation rules.

## Key Claims

- The DOM provides a **platform‑neutral model** for events, aborting, and node trees.
- Events propagate through capture → target → bubble phases; they can be cancelled and may traverse shadow boundaries when `composed` is `true`.
- `AbortController` offers a standard abort primitive, integrated with `addEventListener` and promise‑based APIs via `AbortSignal`.
- Node tree mutations follow strict algorithms, preserving constraints and notifying mutation observers.
- Shadow DOM enables **encapsulation** with slots and slot assignment, affecting event retargeting and `composedPath()`.
- Live `Range` objects **automatically adjust** to tree mutations; `StaticRange` provides a lightweight, immutable alternative.
- Custom element registries support **autonomous** and **customized built-in** elements, with lifecycle callbacks like `connectedCallback` and `attributeChangedCallback`.
- The standard depends on other specifications (Infra, Web IDL, HTML, Encoding, Selectors) and provides **extensibility hooks** for other standards.
- The standard introduces **no known security or privacy considerations** on its own.

## Suggested Links

- [DOM Standard](https://dom.spec.whatwg.org/)
- [Infra Standard](https://infra.spec.whatwg.org/)
- [HTML Standard](https://html.spec.whatwg.org/multipage/)
- [Web IDL Standard](https://webidl.spec.whatwg.org/)
- [Selectors Level 4](https://drafts.csswg.org/selectors/)
- [Encoding Standard](https://encoding.spec.whatwg.org/)
- [UI Events](https://w3c.github.io/uievents/)
- [Trusted Types](https://w3c.github.io/trusted-types/dist/spec/)
- [Touch Events](https://w3c.github.io/touch-events/)
- [Service Workers](https://w3c.github.io/ServiceWorker/)
