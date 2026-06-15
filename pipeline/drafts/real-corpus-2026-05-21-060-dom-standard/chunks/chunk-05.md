---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

Chunk Context
This chunk details the **AbortSignal** interface, its static factory methods (`abort`, `timeout`, `any`), and internal state management (abort algorithms, dependencies). It covers garbage collection rules for dependent signals and guidelines for APIs using promises. The latter half of the chunk introduces **Nodes**, defining the Document tree structure, Shadow trees, and the complex mechanics of **slots** and **slottables** within that context.

Local Summary
The AbortSignal interface allows asynchronous operations to be cancelled via a shared state. It supports creating signals for specific reasons (`abort`), timeouts (`timeout`), or when any of a group of signals is aborted (`any`). The specification defines how these signals propagate abort states through dependencies and source signals, ensuring proper garbage collection only after all listeners and algorithms are cleared. For APIs, the standard mandates rejecting promises immediately if a signal is already aborted. The document then shifts to defining the Node tree hierarchy, distinguishing between light trees (Document) and shadow trees, and explaining how elements and text nodes function as slots or slottables to distribute content within Shadow DOM contexts.

Key Claims
- An `AbortSignal` object has an associated abort reason (initially `undefined`) and a set of abort algorithms that execute upon abortion.
- The static `abort(reason)` method creates a signal with a specific reason or defaults to an "AbortError" DOMException.
- The static `timeout(milliseconds)` method schedules a global task to abort the signal after the specified duration, setting the reason to a "TimeoutError".
- A dependent AbortSignal will be aborted when any of its source signals is aborted, inheriting that specific reason.
- A non-aborted dependent AbortSignal must not be garbage collected while it has active source signals, registered event listeners, or abort algorithms.
- Web platform APIs using promises must accept an `AbortSignal` via a signal dictionary member and reject the promise with the signal's abort reason upon cancellation.
- Nodes are objects implementing the `Node` interface; every node belongs to a primary interface such as `Document`, `Element`, `Text`, or `ShadowRoot`.
- A shadow tree is attached to a light tree (host) and can itself be part of another shadow tree hierarchy.
- Slots are created via HTML's `<slot>` element, while slottables are nodes (Elements, Text, Slots) that can be assigned to those slots.

Entities And Concepts
- **AbortSignal**: Interface for signaling cancellation in asynchronous operations.
- **AbortController**: Controller object (implied context) associated with an AbortSignal.
- **DOMException**: Exception type used for abort reasons ("AbortError", "TimeoutError").
- **Node**: Base interface for the tree structure; includes `Document`, `Element`, `Text`, etc.
- **ShadowRoot**: Root of a shadow tree.
- **Light Tree**: The node tree of a host element containing a shadow root.
- **Slot**: An element that defines a placeholder for slottables in a shadow DOM.
- **Slottable**: A node (Element or Text) capable of being assigned to a slot.
- **Assigned Nodes**: List of slottables currently inside a specific slot.

Procedures And API Details
- **Creating Signals**:
  - `AbortSignal.abort(reason)`: Creates a signal with the given reason.
  - `AbortSignal.timeout(ms)`: Creates a signal that aborts after `ms` milliseconds.
  - `AbortSignal.any(signals)`: Returns a signal aborted when any in the input list is aborted.
- **Checking State**:
  - `signal.aborted`: Boolean indicating if the signal has been aborted.
  - `signal.reason`: The associated DOMException or value causing abort.
  - `signal.throwIfAborted()`: Throws the abort reason if aborted; useful for synchronous checks.
- **Promise Handling Pattern**:
  1. Check if `options["signal"]` exists and is aborted.
  2. If so, reject the promise immediately with `signal.reason`.
  3. Otherwise, add abort steps to the signal: stop operations and reject the promise.
- **Finding Slots**: Use `find a slot for slottable` to locate where a node belongs in a shadow tree based on name or manual assignment.
- **Assigning Nodes**: Run "assign slottables" to update the `assigned nodes` list of slots when the DOM structure changes.

Nuance Or Contradictions
- The term "in a document" is noted as deprecated usage; it now implies being in a document tree that does not account for shadow trees, whereas modern specs use "connected" (shadow-including root is a document).
- Garbage collection rules are strict: dependent signals must remain alive while source signals exist and active listeners/algorithms are present.
- API authors have flexibility to ignore abort wishes if the operation has already completed, provided they adhere to promise rejection rules when possible.

Candidate Wiki Hints
- **AbortSignal**: A dedicated page explaining the interface, methods, and lifecycle.
- **Shadow DOM Slots**: A guide on how slots and slottables interact within shadow trees.
- **Node Tree Structure**: Documentation covering Document vs. Shadow trees and node types.
