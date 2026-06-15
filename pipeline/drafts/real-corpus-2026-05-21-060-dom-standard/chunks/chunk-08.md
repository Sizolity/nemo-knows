---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk defines the `MutationObserver` interface and related mechanisms for observing DOM mutations. It covers the `MutationObserver` constructor, methods (`observe`, `disconnect`, `takeRecords`), configuration options (`MutationObserverInit`), algorithmic steps for queuing mutation records based on specific criteria (attributes, childList, characterData), and the `MutationRecord` interface structure.

## Local Summary
The section details how to construct a `MutationObserver`, configure what changes to watch via options like `subtree` or `attributeFilter`, and stop observation. It specifies the logic for determining which observers receive a record when a mutation occurs, distinguishing between attribute, character data, and child list mutations. Finally, it defines the `MutationRecord` properties that describe the specific change (e.g., old value, affected node, sibling context).

## Key Claims
- A `MutationObserver` maintains an internal queue of records until explicitly retrieved via `takeRecords()`.
- The observer's callback is invoked after nodes registered with `observe()` are mutated.
- Specific validation rules apply to the `options` object passed to `observe()`, throwing a `TypeError` if mutually exclusive options (e.g., `attributeOldValue` without `attributes`) are used or if no observation type is selected.
- When queuing records, the system iterates through inclusive ancestors of the target node to determine which observers should be notified based on their `subtree` setting and specific filter criteria.
- The `MutationRecord` interface distinguishes between "attributes", "characterData", and "childList" mutation types via its `type` attribute.

## Entities And Concepts
- **MutationObserver**: Interface for observing DOM mutations.
- **MutationRecord**: Object representing a single mutation event.
- **observe()**: Method to start watching a specific node.
- **disconnect()**: Method to stop watching.
- **takeRecords()**: Method to retrieve and clear the internal queue.
- **MutationObserverInit**: Dictionary defining observation options (`childList`, `attributes`, `characterData`, `subtree`, etc.).
- **MutationCallback**: Type definition for the function passed to the constructor.

## Procedures And API Details
- **Constructor**: `new MutationObserver(callback)` sets up the callback which receives a sequence of `MutationRecord` objects and the observer itself.
- **observe()**:
  - Automatically enables `attributes` if `attributeOldValue` or `attributeFilter` is set but `attributes` is omitted.
  - Automatically enables `characterData` if `characterDataOldValue` is set but `characterData` is omitted.
  - Throws `TypeError` if no observation type (`childList`, `attributes`, `characterData`) is true.
  - Throws `TypeError` if old value options are used without enabling their respective types.
  - Throws `TypeError` if an attribute filter is present without enabling `attributes`.
- **disconnect()**: Removes the observer from all registered node lists and empties the record queue.
- **Queuing Logic**:
  - Iterates through ancestors of the target node.
  - Checks if the mutation type matches the observer's options (e.g., `subtree` is false, or specific attribute filters don't match).
  - If the mutation should be observed, creates a `MutationRecord` with appropriate `oldValue` (if requested) and enqueues it.
  - Queues a microtask to process these records.

## Nuance Or Contradictions
- **Option Defaults**: The specification explicitly states that setting `attributeOldValue` implies `attributes` is true, even if not explicitly set in the object literal, preventing silent failures where an old value is requested but no attribute watching is active.
- **Weak References**: The observer's internal node list consists of weak references, implying automatic cleanup if the observed nodes are garbage collected, though the text focuses on the removal logic during `disconnect()`.

## Candidate Wiki Hints
- Page: **MutationObserver** (Overview of usage and lifecycle)
- Page: **MutationRecord** (Data structure details)
- Concept: **DOM Mutation Observation Algorithms** (The internal logic for filtering and queuing records)
