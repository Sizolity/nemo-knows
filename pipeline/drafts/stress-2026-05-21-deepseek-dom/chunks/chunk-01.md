---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- Source: raw/web/corpus-2026-05-18/060-dom-standard.md
- Chunk: 1 of 9
- Lines: 1–1579
- Heading path: Document
- Covers the DOM Standard sections 1 (Infrastructure) through 3 (Aborting ongoing activities), including trees, events, AbortController/Signal.

## Local Summary
The chunk lays out foundational DOM infrastructure: tree definitions, ordered set parsing/serialization, selector scoping, and name validation rules. It then introduces the DOM Events model — the Event, CustomEvent, EventTarget interfaces, event constructors, dispatching, firing, and the principle that events signal occurrences, not actions. Finally, it defines the AbortController/AbortSignal API for cooperative cancellation, along with dependent abort signals and static factory methods (`AbortSignal.abort()`, `.timeout()`, `.any()`).

## Key Claims
- The DOM Standard depends on the Infra Standard and extends concepts for trees, events, and aborting.
- A **tree** is finite hierarchical structure; tree order is preorder, depth-first traversal.
- **Event** objects signal an occurrence; they are not actions. The old “default actions” concept was misleading.
- `EventTarget` provides methods to add/remove listeners and dispatch events. Listeners have capture, passive, once, and signal options.
- `addEventListener()` with an `AbortSignal` allows removal on abort.
- Event dispatching builds a path through the target’s ancestors, invoking capture and bubble listeners in tree order (reverse for bubble).
- `AbortController` and `AbortSignal` enable cooperative cancellation; `AbortSignal` can be used as a promise rejection reason.
- `AbortSignal.any()` creates a dependent signal that aborts when any source signal aborts.
- Name validations for namespaces, elements, attributes, and doctypes are loosened compared to older specifications to match HTML parser capabilities.

## Entities And Concepts
- **Tree**: object with parent/children; definitions: root, descendant, ancestor, sibling, first/last child, previous/next sibling, index.
- **Tree order**: preorder depth-first traversal.
- **Ordered set**: parsed by splitting on ASCII whitespace; serialized with U+0020 SPACE.
- **Scope-match a selectors string**: parse selector, match against node’s root using scoping root node; no namespace support planned.
- **Valid namespace prefix**: length ≥ 1, no ASCII whitespace, NULL, `/`, or `>`.
- **Valid attribute local name**: length ≥ 1, no whitespace, NULL, `/`, `=`, or `>`.
- **Valid element local name**: defined by rules allowing a wide range after the first character, with a JavaScript-compatible regex.
- **Event**: has type, target, relatedTarget, touch target list, path, flags (stop propagation, stop immediate propagation, canceled, in passive listener, composed, initialized, dispatch).
- **EventTarget**: has event listener list (listeners with type, callback, capture, passive, once, signal, removed), a get-the-parent algorithm (default null), activation behavior algorithms.
- **Event listener** structure: type, callback, capture, passive, once, signal, removed.
- **AbortController**: associated with an `AbortSignal`; `abort(reason)` stores reason and signals abort.
- **AbortSignal**: has abort reason (undefined until aborted), abort algorithms, dependent flag, source/dependent signals; provides `aborted`, `reason`, `throwIfAborted()`, `onabort`.
- **Dependent abort signal**: created via `AbortSignal.any()`; aggregates signals.

## Procedures And API Details
- **Ordered set parser**: input split on ASCII whitespace → tokens appended to new ordered set.
- **Name validation and extraction**: `validate and extract a namespace and qualifiedName` algorithm: checks namespace prefix validity, local name validity depending on context (attribute/element), enforces XML/XMLNS namespace constraints.
- **Event constructor**: calls inner event creation steps with `EventInit` dictionary.
- **Inner event creation steps**: create object → set initialized flag, `timeStamp`, dictionary members, run event constructing steps.
- **Create an event**: used by specs to create and dispatch separately; initializes `isTrusted` to true.
- **Event dispatching**: sets dispatch flag, builds path via get-the-parent, invokes listeners in capture then bubble phase (reverse tree order for bubble). Handles shadow DOM and closed-tree visibility via `composedPath()`.
- **Event path**: list of structs with invocation target, shadow-adjusted target, relatedTarget, etc.
- **Invoke**: for each path struct, sets `currentTarget`, clones listener list, runs inner invoke; if no listener found and event is trusted, tries legacy event type mapping (e.g., `animationend` → `webkitAnimationEnd`).
- **Fire an event**: convenience to create and dispatch an event with given constructor and attribute initializations.
- **AbortController constructor**: creates a new `AbortSignal`.
- **signal abort**: if not already aborted, set reason, propagate to dependent signals, run abort steps (execute algorithms, empty them, fire `abort` event).
- **AbortSignal.abort(reason)**: returns already-aborted signal with given reason or `"AbortError"` `DOMException`.
- **AbortSignal.timeout(milliseconds)**: returns signal that aborts with `"TimeoutError"` after timeout.
- **AbortSignal.any(signals)**: returns dependent signal that aborts when any source signal aborts; uses `create a dependent abort signal`.
- **add an event listener**: checks for ServiceWorker warnings, ignores null callback or aborted signal, sets default passive if not specified, adds to list if not duplicate, attaches abort steps if signal present.
- **remove an event listener**: marks listener as removed, deletes from list.

## Nuance Or Contradictions
- Name validation rules were loosened to align with what the HTML parser can create, avoiding developer frustration about names valid in parser but rejected by DOM APIs.
- The `EventTarget` interface’s `dispatchEvent()` can influence cancelable operations: returning `false` indicates the event was canceled.
- Passive event listeners exist because some APIs (touch, wheel) need to know whether preventDefault will be called to optimize scrolling; the presence of non-passive listeners can make the event cancelable.
- `isTrusted` is `false` for synthetic events except for the legacy `click()` method which dispatches an untrusted click event.
- The attribute `event` on `Window` (`window.event`) returns the current event; it is not available in workers or worklets and inaccurate for shadow tree events. Developers are encouraged to use the event parameter.
- Activation behavior and legacy pre-activation/canceled-activation algorithms exist for historical compatibility with `area` and checkbox/radio elements.
- Events signify an occurrence, not an action; they cannot start an algorithm, only influence an ongoing one.

## Candidate Wiki Hints
- “DOM Tree Concepts” – definitions of tree, tree order, ancestors, siblings, index.
- “DOM Event Model” – overview of event flow, capture/bubble phases, target, `EventTarget`.
- “EventTarget and EventListener” – details on adding/removing listeners, options (passive, once, signal), default passive value.
- “AbortController and AbortSignal” – API design, dependent signals, static methods, cooperative cancellation.
- “DOM Name Validation Rules” – valid namespace prefix, attribute/element local names, algorithm for validate-and-extract.
- “Event Constructing and Firing” – inner event creation steps, `create an event`, `fire an event` patterns.
- “Event Dispatching Algorithm” – path building, invocation, legacy event type fallback.
