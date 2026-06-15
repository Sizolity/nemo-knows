---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

This chunk covers the beginning of the **DOM Standard** (Living Standard), specifically focusing on foundational concepts such as trees, ordered sets, selectors, name validation, and an introduction to DOM events. It establishes the structural definitions required for the rest of the standard.

# Local Summary

The document defines the core infrastructure for the DOM, starting with abstract data structures like finite hierarchical trees and ordered sets. It outlines how nodes relate to one another (parent, child, ancestor, descendant) and introduces utility concepts like selectors and name validation rules. The section concludes by introducing the event system, explaining how events are dispatched, captured, and bubbled through the tree structure.

# Key Claims

-   **Tree Structure**: A tree is a finite hierarchical structure where nodes have parents (null or object) and ordered children.
-   **Relationships**: Relationships are defined as inclusive (e.g., "inclusive descendant" includes the node itself). Precedence and following order are determined by tree order (preorder, depth-first).
-   **Ordered Sets**: These are parsed from whitespace-separated strings and serialized back to such strings using space delimiters.
-   **Selectors**: Scoping-match is performed against a parsed selector string; namespace support within selectors is explicitly not planned.
-   **Name Validation**: Valid element local names must start with an ASCII alpha or specific non-ASCII code points, allowing broader construction possibilities than strict XML parsers but maintaining compatibility with the HTML parser branch.
-   **Events**: Events are objects implementing `EventTarget` and `Event`. They can be synthetic (created by the application) or dispatched by the user agent.
-   **Event Propagation**: Events traverse ancestors in two phases: first capturing (downwards/inclusive ancestors, `capture: true`) in tree order, then bubbling (upwards/inclusive ancestors, `capture: false`) in reverse tree order.

# Entities And Concepts

-   **Tree**: Finite hierarchical structure.
-   **Node Object**: Has a parent and children; participates in the tree.
-   **Root**: Any object in a tree whose parent is null.
-   **Descendant/Ancestor**: Hierarchical relationships (child, parent, etc.).
-   **Sibling**: Objects sharing the same non-null parent.
-   **Ordered Set**: Collection of tokens parsed from whitespace-separated strings.
-   **Selector**: String used for scoping-match; namespaces not supported.
-   **Event**: Object signaling an occurrence; implements `EventTarget`.
-   **EventTarget**: Interface allowing addition/removal of listeners via `addEventListener()` and `removeEventListener()`.
-   **Synthetic Event**: An event created and dispatched by the application (e.g., using `CustomEvent`).
-   **Bubbles/Capture**: Modes of event propagation through the tree.

# Procedures And API Details

-   **Tree Order Traversal**: Depth-first traversal starting from the root.
-   **Ordered Set Parsing**:
    1.  Split input on ASCII whitespace.
    2.  Append tokens to a new ordered set.
    3.  Return tokens.
-   **Selector Scoping-Match**:
    1.  Parse selector string.
    2.  Throw "SyntaxError" if parsing fails.
    3.  Match against tree with specified scoping root.
-   **Event Listener Removal**:
    -   Call `removeEventListener()` with same arguments used to add.
    -   Pass an `AbortSignal` to `addEventListener()` and call `abort()` on the controller.
-   **Synthetic Event Dispatch**:
    1.  Create event (e.g., `new CustomEvent(...)`).
    2.  Call `obj.dispatchEvent(event)`.
    3.  Check return value of `dispatchEvent()` to see if the event was canceled.
-   **Event Propagation Flow**:
    1.  Invoke listeners on inclusive ancestors where `capture` is `true` (tree order).
    2.  If `event.bubbles` is true, invoke listeners on inclusive ancestors where `capture` is `false` (reverse tree order).

# Nuance Or Contradictions

-   **Namespace Support**: The standard explicitly states that support for namespaces within selectors is not planned and will not be added.
-   **Name Validation Loosening**: Previous strict validations aligned with XML specifications were found annoying for web developers because they prevented creating names possible via the HTML parser. Validations have been loosened to allow names constructible by the HTML parser plus additional possibilities, while restricting ASCII ranges for historical reasons beyond that.

# Candidate Wiki Hints

-   **DOM Tree Structure**: Explaining parent/child/sibling relationships and tree order.
-   **Ordered Sets in DOM**: How whitespace-separated strings are treated as collections.
-   **DOM Selectors**: Basic usage and the explicit lack of namespace support.
-   **Name Validation Rules**: The relaxed rules for element local names compared to strict XML.
-   **DOM Events Overview**: Introduction to `EventTarget`, synthetic events, and propagation phases (capture/bubble).
