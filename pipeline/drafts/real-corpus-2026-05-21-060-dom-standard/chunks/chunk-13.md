---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

Chunk Context
- Source section: 4.9. Interface Element (part of the Return document specification).
- Line range: 4463–4929.
- Focus: Definition and behavior of the `Element` interface, including attributes, shadow roots, custom element states, attribute handling algorithms, and ID/class/slot semantics.

Local Summary
The `Element` interface represents DOM nodes that have a name (tag), optional namespace, prefix, local name, and an associated shadow root. The section details:
- Attributes like `id`, `className`, `classList`, and `slot`.
- Methods for querying descendants (`closest`, `matches`, `getElementsByTagName*`).
- Shadow DOM integration via `attachShadow` and `shadowRoot`.
- Custom element lifecycle states ("undefined", "failed", "uncustomized", "precustomized", "custom") and upgrade mechanisms.
- Algorithms for creating elements, handling attribute changes, and managing qualified names (including HTML-uppercased variants).
- Reflection of string attributes (`id`, `class`, `slot`) and their associated IDL behavior.

Key Claims
- An element is defined if its custom element state is "uncustomized" or "custom".
- Only one ID per element is allowed; the concept of ID is formalized in the DOM spec.
- Qualified names are computed from local name and prefix, with HTML namespace elements uppercased in HTML documents.
- Custom element upgrade reactions are enqueued when an element transitions to a defined state (except for synchronous construction).
- Attribute changes trigger mutation records and `attributeChangedCallback` if the element is custom.
- Shadow roots are null by default; non-null indicates a shadow host.

Entities And Concepts
- **Element**: Base interface for DOM nodes with tag names, namespaces, and attributes.
- **ShadowRootInit**: Dictionary used in `attachShadow()` to configure shadow root mode, focus delegation, slot assignment, etc.
- **Custom element states**: "undefined", "failed", "uncustomized", "precustomized", "custom".
- **Qualified name**: Computed as `[prefix]:localName` or just `localName`; HTML-uppercased variant for HTML documents.
- **Attribute list**: Exposed via `NamedNodeMap`; managed through append/remove/replace algorithms.
- **ID, class, slot**: Super-global content attributes reflectable on any element.
- **Shadow host**: An element with a non-null `shadowRoot`.

Procedures And API Details
- **Creating an element**:
  - Input: document, localName, namespace, prefix, synchronousCustomElements flag, registry.
  - Steps involve looking up custom element definitions, constructing internal elements, handling upgrade reactions or immediate construction if synchronous.
  - Exceptions during synchronous upgrades are reported to the global object; failed states prevent re-execution.
- **Attribute operations**:
  - `setAttribute`, `getAttribute`, `removeAttribute`, `toggleAttribute`.
  - Internal algorithms for appending/removing/replacing attributes with mutation record queuing.
  - Trusted type validation on attribute values via `get trusted type compliant attribute value`.
- **Query methods**:
  - `closest(selectors)`, `matches(selectors)`, legacy `webkitMatchesSelector`.
  - `getElementsByTagName*` returning `HTMLCollection`.
  - `insertAdjacentElement` (legacy), `insertAdjacentText` (legacy).
- **Shadow DOM**:
  - `attachShadow(init)` returns a `ShadowRoot`.
  - `shadowRoot` getter returns the associated shadow root or null.

Nuance Or Contradictions
- `customElementRegistry` in `ShadowRootInit` allows both `undefined` and `null` to pass a `ShadowRoot` node directly instead of a dictionary.
- The spec does not claim conformance for using `id`, `class`, or `slot` on elements; it only defines their behavior.
- Historical DTD-based multiple identifiers are superseded by the single DOM `id` concept.
- Legacy aliases like `webkitMatchesSelector` exist alongside standard `matches`.

Candidate Wiki Hints
- **Element Interface Overview**: Summary of `Element` attributes, methods, and shadow root association.
- **Custom Element Lifecycle**: Explanation of states ("undefined" to "custom") and upgrade mechanisms.
- **Attribute Handling Algorithms**: Deep dive into how `setAttribute`, `removeAttribute`, etc., trigger mutations and callbacks.
- **Qualified Name Computation**: How tag names are formatted, including HTML-uppercase rules.
- **Shadow DOM Integration**: Usage of `attachShadow` and shadow host detection.
