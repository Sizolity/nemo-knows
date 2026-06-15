---
title: Chunk 16 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk details the **Range** interface within the DOM Standard, specifically covering section 5.5 "Interface Range". It defines the `Range` object used to represent a selection or range of nodes in a document tree, including its attributes (like `commonAncestorContainer`), methods for setting boundaries (`setStart`, `setEnd`, etc.), and algorithms for maintaining consistency with the live tree structure during modifications.

## Local Summary
The section defines the `Range` interface as representing a selection of nodes within a DOM tree. It specifies how ranges interact with node containment, common ancestors, and tree modifications (pre-remove steps). The text provides the full interface definition, constructor behavior, getter/setter logic for boundary points, methods for selecting nodes or collapsing ranges, and comparison algorithms between ranges.

## Key Claims
- Objects implementing the Range interface are known as **live ranges**.
- Algorithms that modify a tree (insert, remove, move, replace data, split) automatically modify associated live ranges.
- A node is contained in a live range if its root matches the range's root and it lies between the start and end boundaries.
- The `commonAncestorContainer` attribute returns the furthest ancestor common to both the start and end nodes of the range.
- Setting a range's start or end involves ensuring both points share the same document root; otherwise, the other boundary is adjusted.
- Methods like `selectNode` and `selectNodeContents` automatically adjust boundaries to encompass specific nodes or their contents.

## Entities And Concepts
- **Range**: The interface representing a selection in a document tree.
- **Live Range**: A range object associated with a live DOM tree that updates automatically when the tree changes.
- **Boundary Point**: Defined as a pair `(node, offset)` specifying a position within a node.
- **Common Ancestor Container**: The deepest common ancestor of the range's start and end nodes.
- **Containment Rules**: Logic defining which nodes are strictly inside, partially inside, or outside a range based on ancestry and index order.
- **DOMException**: Error types thrown for invalid operations (e.g., `InvalidNodeTypeError`, `IndexSizeError`, `WrongDocumentError`).

## Procedures And API Details
- **Constructor**: `new Range()` initializes the range with start and end at `(current document, 0)`.
- **Boundary Setting**:
    - `setStart(node, offset)`: Sets the start boundary. If roots differ or the point is after the current end, the end is updated first.
    - `setEnd(node, offset)`: Sets the end boundary. Adjusts the start if necessary to maintain root consistency.
    - `setStartBefore/After(node)`: Positions the start relative to a specific node's index within its parent.
    - `setEndBefore/After(node)`: Positions the end relative to a specific node's index within its parent.
- **Selection**:
    - `selectNode(node)`: Sets boundaries immediately before and after the given node.
    - `selectNodeContents(node)`: Sets boundaries at the start (0) and end (length) of the node, throwing error if the node is a doctype.
- **Collapse**: `collapse(toStart)` collapses the range by setting end to start if true, or start to end otherwise.
- **Comparison**: `compareBoundaryPoints(how, sourceRange)` compares two ranges based on start/end points relative to each other (START_TO_START, START_TO_END, etc.), throwing errors if roots differ or `how` is invalid.
- **Tree Modification Logic**: Before removing a node, the spec details steps to adjust start/end offsets and indices of all live ranges affected by the removal.

## Nuance Or Contradictions
- **Root Consistency**: A range cannot have its start and end in different documents. Any attempt to set a boundary outside the current document's root forces the other boundary to move into alignment with the valid root, effectively clamping the range.
- **Containment Paradox**: The start and end nodes themselves are *never* considered contained within the range, even though they define its boundaries. Only their descendants (or the node itself for CharacterData) may be included in the selection content.
- **Partial Containment**: A node is partially contained if it is an ancestor of either the start or end node but not both. This occurs only when the start and end nodes are different and neither is an ancestor of the other (requiring a distinct common inclusive ancestor).

## Candidate Wiki Hints
- **Page: Range Interface** – Documenting the `Range` API, boundary points, and containment logic for developers implementing text selection or manipulation.
- **Page: Live Ranges** – Explaining how DOM mutations trigger automatic updates to associated range objects.
