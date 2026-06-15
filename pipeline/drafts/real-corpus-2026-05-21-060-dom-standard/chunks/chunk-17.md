---
title: Chunk 17 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
- Heading path: 5. Switch on the position of thisPoint relative to sourcePoint:
- Lines covered: 5925–6287
- Topics: deleteContents(), extractContents(), cloneContents(), insertNode(), surroundContents(), cloneRange(), detach(), comparePoint, intersectsNode, isPointInRange, stringification behavior, and NodeIterator/TreeWalker traversal.

Local Summary
The chunk details the algorithms for manipulating DOM Range objects, including deletion, extraction, cloning, insertion, surrounding contents, and point comparison. It also covers traversal mechanics via NodeIterator and TreeWalker, detailing active state management and filtering logic.

Key Claims
- `deleteContents()` returns early if the range is collapsed; it reconstructs start/end nodes to maintain a valid live range before removing contained nodes.
- Extracting or cloning contents involves splitting CharacterData nodes at boundaries and handling partially contained children by creating sub-ranges and fragments.
- If any member of `containedChildren` is a doctype, throwing a "HierarchyRequestError" DOMException occurs during extraction or cloning.
- The `insertNode` method ensures pre-insert validity and updates the range end if collapsed after insertion.
- `surroundContents` throws an "InvalidStateError" if a non-Text node is partially contained and an "InvalidNodeTypeError" if the new parent is invalid (Document, DocumentType, or DocumentFragment).
- `comparePoint`, `isPointInRange`, and `intersectsNode` methods validate document roots, reject doctypes, and handle offset bounds before performing comparisons.
- NodeIterator and TreeWalker objects maintain an `is active` flag to prevent recursive invocations and use a `whatToShow` bitmask for filtering.

Entities And Concepts
- Range (DOM)
- Live Range
- DocumentFragment
- CharacterData
- DOMException (HierarchyRequestError, InvalidStateError, InvalidNodeTypeError, WrongDocumentError, IndexSizeError)
- NodeIterator
- TreeWalker
- whatToShow (bitmask)
- FILTER_ACCEPT / FILTER_SKIP

Procedures And API Details
- `deleteContents()`:
  - Validates collapsed state.
  - Reconstructs start/end nodes if necessary.
  - Removes contained nodes in tree order.
- `extractContents()`:
  - Returns a DocumentFragment containing cloned/extracted nodes.
  - Handles partial containment by splitting CharacterData and creating sub-ranges.
- `cloneContents()`:
  - Clones the contents of a range into a fragment.
- `insertNode(node)`:
  - Validates insertion point (no ProcessingInstruction, Comment, or null-parent Text).
  - Splits Text nodes if necessary.
  - Updates range end if collapsed.
- `surroundContents(newParent)`:
  - Checks for partial containment of non-Text nodes.
  - Replaces children in newParent and appends the fragment.
- `cloneRange()`:
  - Returns a new Range with identical start/end points.
- `detach()`:
  - No-op method retained for compatibility.
- `comparePoint(node, offset)`:
  - Returns −1 (before), 0 (in range), or 1 (after).
- `intersectsNode(node)`:
  - Checks if the range intersects a given node using parent offsets.
- `isPointInRange(node, offset)`:
  - Validates document root and bounds before returning true/false.

Nuance Or Contradictions
- The `detach()` method is described as doing nothing, noting that its functionality (disabling a Range) was removed but the method remains for compatibility.
- CharacterData nodes are not explicitly checked in `surroundContents` due to historical reasons, though they may throw later errors.
- Doctype nodes cannot be boundary points or ancestors, simplifying logic regarding partial containment checks.

Candidate Wiki Hints
- DOM Range Manipulation Algorithms
- Deleting and Extracting Range Contents
- Cloning and Inserting Nodes into Ranges
- Surrounding Contents with a New Parent
- Comparing Points and Checking Intersection in DOM
- NodeIterator and TreeWalker Traversal Mechanics
