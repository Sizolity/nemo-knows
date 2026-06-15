---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
- Heading path: 2. Events > 2.9. Dispatching events through 3.1. Interface AbortController
- Lines covered: 1069–1406
- Topics: Event dispatch algorithms, event paths, activation behavior, legacy target overrides, event firing vs dispatching, action versus occurrence semantics, and the AbortController/AbortSignal abort mechanism.

Local Summary
This chunk details how an event is dispatched across a document tree, including handling of shadow DOM, touch targets, and activation behaviors (e.g., click on buttons). It defines algorithms for appending to an event path, invoking listeners during capturing/bubbling phases, and managing legacy behavior flags. It clarifies that events signal occurrences rather than initiating actions, distinguishing them from "default actions." The section concludes by introducing AbortController/AbortSignal as the standard mechanism for aborting ongoing activities, providing interface definitions and example usage patterns for promise-based APIs.

Key Claims
- Dispatching involves setting an event’s dispatch flag, constructing an event path with targets (including shadow roots), and invoking listeners in phases (capturing → at target → bubbling).
- Activation behavior (e.g., `click`) triggers only on specific elements (like buttons) and respects the `bubbles` attribute.
- Shadow DOM is handled via "slot-in-closed-tree" flags and special retargeting logic when traversing closed shadow roots.
- Events are notifications that influence future algorithm steps; they do not cause actions to start.
- AbortController provides an `abort()` method to signal cancellation, which should reject unsettled promises with an `AbortError` (reason: "AbortError" if no reason given).

Entities And Concepts
- Event: The object representing a user or programmatic action.
- Event Path: A list of structs describing the targets and contexts through which an event will travel during dispatch.
- Activation Behavior: Special handling for elements like buttons that respond to activation events (e.g., clicks).
- Shadow DOM: Encapsulated trees where events may be retargeted or blocked depending on mode ("closed").
- AbortController / AbortSignal: A pair of objects used to signal cancellation of asynchronous operations.
- Legacy Target Override: An optional flag used by HTML when the target is a Window object.

Procedures And API Details
- Dispatching an event involves:
  - Setting the event’s dispatch flag.
  - Determining the target and activationTarget (if applicable).
  - Building an event path by traversing parents, handling shadow roots, and adjusting for slottable elements.
  - Invoking listeners in reverse order of structs in the path during capturing phase, then bubbling phase.
  - Clearing targets if `clearTargets` is true.
- Firing an event: Creates an event object, initializes its attributes (including IDL attributes), and dispatches it at a target.
- AbortController usage:
  ```js
  const controller = new AbortController();
  const signal = controller.signal;
  // ... use signal in async operations ...
  controller.abort(); // signals abort
  ```
- `abort(reason)` stores the reason in the associated `AbortSignal`; if no reason is provided, it uses a default "AbortError".

Nuance Or Contradictions
- Legacy behavior flags (`legacy target override flag`, `legacyOutputDidListenersThrowFlag`) are specific to certain APIs (HTML, Indexed Database) and are not part of the general event model.
- Events do not represent actions; using them as initiators is incorrect per the DOM spec’s philosophy.
- The distinction between "fire" (create + initialize + dispatch) and "dispatch" (run through the algorithm on an existing event) is important for clarity.

Candidate Wiki Hints
- Page: `dom-events-dispatch-algorithm` – Summarize steps 2.9, including event path construction and phase invocation.
- Page: `dom-events-firing-vs-dispatching` – Clarify the difference between creating/firing an event and dispatching it.
- Page: `dom-action-versus-occurrence` – Explain that events are notifications, not actions.
- Page: `dom-abortcontroller-signaling` – Document the AbortController interface and abort semantics for promise-based APIs.
