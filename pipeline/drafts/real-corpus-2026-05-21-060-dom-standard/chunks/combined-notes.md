## chunk-01

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

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context

This chunk covers the specification for DOM events, focusing on the `Event` interface definition, its legacy attributes, and behavior. It details the `CustomEvent` interface for carrying custom data, explains the internal steps for constructing events (including handling shadow DOM paths), and provides guidelines for defining new event interfaces. The section also notes legacy extensions to the `Window` interface and deprecated methods like `initEvent()`.

## Local Summary

The `Event` interface represents an occurrence in the DOM. It is defined with a constructor accepting a type and optional initialization dictionary. Key attributes include `type`, `target`, `currentTarget`, `bubbles`, `cancelable`, and `timeStamp`. The chunk details the lifecycle of event propagation phases (`NONE`, `CAPTURING_PHASE`, `AT_TARGET`, `BUBBLING_PHASE`) and methods to control propagation (`stopPropagation()`, `stopImmediatePropagation()`) or default behavior (`preventDefault()`). It also defines `CustomEvent` for passing custom data via a `detail` attribute. The specification outlines the algorithmic steps for constructing events, handling shadow DOM paths in `composedPath()`, and legacy compatibility considerations.

## Key Claims

- An `Event` object signals that something has occurred (e.g., an image download completion).
- A potential event target is either `null` or an `EventTarget` object.
- The `target` attribute returns the object to which the event was dispatched.
- The `currentTarget` attribute returns the object whose listener callback is currently being invoked.
- The `composedPath()` method returns a list of invocation targets, excluding nodes in closed shadow trees not reachable from the current target.
- Invoking `stopPropagation()` prevents the event from reaching other objects in the tree after the current one.
- Invoking `stopImmediatePropagation()` stops listeners after the current one and halts further propagation.
- The `bubbles` attribute indicates if the event traverses ancestors in reverse tree order.
- The `cancelable` attribute indicates if `preventDefault()` can cancel the operation that caused the event.
- `CustomEvent` extends `Event` and includes a `detail` attribute for custom data.
- Legacy methods like `initEvent()` are redundant with constructors and should be avoided, though supported for legacy content.
- The `Window` interface has a legacy `event` attribute (deprecated in favor of passing events to listeners).

## Entities And Concepts

- **Event**: Interface representing an occurrence.
- **EventTarget**: Object to which an event is dispatched.
- **CustomEvent**: Sub-interface of `Event` for carrying custom data.
- **EventInit**: Dictionary defining initialization options (`bubbles`, `cancelable`, `composed`).
- **CustomEventInit**: Dictionary extending `EventInit` with a `detail` attribute.
- **EventPhase**: Enumerated phases: `NONE` (0), `CAPTURING_PHASE` (1), `AT_TARGET` (2), `BUBBLING_PHASE` (3).
- **composedPath()**: Method returning the path of invocation targets.
- **isTrusted**: Boolean indicating if the event was dispatched by the user agent.
- **stopPropagation()**: Method to halt propagation to subsequent nodes.
- **stopImmediatePropagation()**: Method to halt listeners and propagation.
- **preventDefault()**: Method to cancel the default action of an event.
- **Window.event**: Legacy attribute holding the current event (deprecated).

## Procedures And API Details

### Constructing an Event
To create an `Event`:
1. Call `new Event(type, eventInitDict)`.
2. The constructor initializes `type`, `bubbles`, and `cancelable` via the dictionary.

### Using CustomEvent
To create a custom event with data:
1. Call `new CustomEvent(type, { detail: yourData, bubbles: false, cancelable: false })`.
2. Access custom data via `event.detail`.

### Controlling Propagation
- **Stop propagation**: Call `event.stopPropagation()`.
- **Stop immediate propagation**: Call `event.stopImmediatePropagation()`.
- **Prevent default action**: Call `event.preventDefault()` if `cancelable` is true.

### Retrieving Path Information
- `event.target`: Returns the target element or null.
- `event.currentTarget`: Returns the listener's host element.
- `event.composedPath()`: Returns an array of `EventTarget` objects representing the path, excluding closed shadow tree nodes unreachable from the current target.

### Event Initialization Steps
When initializing an event (manually or via constructor):
1. Set `initialized` flag to true.
2. Unset propagation stop flags and canceled flag.
3. Set `isTrusted` to false (unless created by user agent).
4. Set `target` to null.
5. Set `type`, `bubbles`, and `cancelable` attributes.

### Shadow DOM Path Handling (`composedPath`)
The algorithm filters the internal path:
1. Identify closed shadow trees.
2. Traverse the path, skipping nodes in closed trees not reachable from `currentTarget`.
3. Collect invocation targets that are accessible.

## Nuance Or Contradictions

- **Legacy vs Modern API**: The specification maintains legacy attributes like `srcElement` (alias for `target`) and `cancelBubble` (alias for `stopPropagation()`), but strongly encourages using modern APIs.
- **Window.event Attribute**: This attribute is marked as `[Replaceable]` and deprecated. It is inaccurate for shadow tree events and unavailable in workers/worklets. Developers should rely on the event passed to listeners.
- **initEvent() Redundancy**: The `initEvent()` method is described as redundant with constructors and incapable of setting the `composed` flag, yet must be supported for legacy content compatibility.
- **isTrusted Initialization**: Normally initialized to `false` upon creation, but exceptions exist (e.g., click events dispatched by user agents might differ in older contexts).
- **preventDefault() Side Effects**: Invoking `preventDefault()` has no effect if the event is not cancelable or if invoked outside a passive listener context. User agents are encouraged to log causes for debugging.

## Candidate Wiki Hints

1. **Event Interface Overview**: A page summarizing the `Event` interface, its attributes, and common methods.
2. **CustomEvent Guide**: Documentation on using `CustomEvent` for custom data passing in modern web apps.
3. **Event Propagation Control**: A guide explaining `stopPropagation()`, `stopImmediatePropagation()`, and `preventDefault()` with usage examples.
4. **Legacy Event Attributes**: A reference page listing deprecated attributes (`srcElement`, `cancelBubble`, `Window.event`) and their modern equivalents.
5. **Shadow DOM and Events**: An article detailing how events interact with shadow DOM, including the `composedPath()` algorithm logic.
6. **Event Construction Best Practices**: Guidelines for creating custom event types, emphasizing constructor usage over legacy `init*Event` methods.

## chunk-03

---
title: Chunk 03 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
- Heading path: 2. Events > 2.7. Interface EventTarget, 2.8. Observing event listeners
- Lines: 819–1068
- Coverage: Definition of `EventTarget`, event listener mechanics, activation behaviors, and performance implications of observing listeners.

Local Summary
This chunk defines the `EventTarget` interface as the foundational object for attaching listeners to events in the DOM. It details how listeners are added (`addEventListener`) and removed (`removeEventListener`), explaining the flattening logic that handles legacy boolean options versus structured dictionaries. The text distinguishes between the listener object and the broader event listener concept, which includes state flags like `capture`, `passive`, `once`, and `signal`. Special attention is given to activation behaviors for specific elements (like checkboxes) and the performance trade-offs involved in observing listener presence to optimize event handling (e.g., making touch events uncancelable if all listeners are passive).

Key Claims
- An `EventTarget` object represents a target to which an event can be dispatched.
- Each `EventTarget` has an associated event listener list, initially empty.
- The `addEventListener` method appends a listener only if no listener with the same type, callback, and capture state already exists.
- Options passed to `addEventListener` can be a boolean (legacy) or a dictionary; booleans are treated as setting the `capture` flag.
- A `passive` listener indicates that the callback will not call `preventDefault()`, allowing for performance optimizations.
- An `once` listener is automatically removed after being invoked once.
- An `AbortSignal` can be used to abort and remove a listener when the signal is aborted.
- Activation behaviors exist for specific elements (e.g., `<area>`) to trigger synthetic events, distinct from general event handling.
- Legacy pre-activation and legacy-canceled-activation behaviors are reserved strictly for checkbox and radio input elements.

Entities And Concepts
- **EventTarget**: The base interface for objects that can receive events.
- **EventListener**: An object with a `handleEvent` method invoked by the target.
- **EventListenerOptions**: A dictionary defining listener behavior (`capture`, `passive`, `once`, `signal`).
- **Activation Behavior**: Actions taken on specific elements (e.g., clicking an area) in response to synthetic events.
- **Passive Event**: A listener that does not prevent default actions, enabling parallel scrolling processing.
- **AbortSignal**: Used to detach listeners when the associated signal is aborted.

Procedures And API Details
- **Creating a Target**: `target = new EventTarget()` creates an object capable of dispatching and listening for events; the constructor steps currently do nothing.
- **Adding a Listener**:
  1. Flatten options: Convert boolean to capture value, extract `once`, `passive`, and `signal` from dictionary if present.
  2. Check signal: If the signal is aborted, return immediately.
  3. Set default passive value based on event type (e.g., `touchstart`) and target context.
  4. Append listener to the list if a duplicate (same type, callback, capture) does not exist.
  5. Add abort steps to the listener if a signal is provided.
- **Removing a Listener**: Flatten options to determine capture state; remove the specific listener matching type, callback, and capture from the list.
- **Dispatching an Event**: `dispatchEvent(event)` initializes the event's trusted flag to false and returns true unless `preventDefault()` was called on a cancelable event or the event was invalid.
- **Flattening Options Algorithm**:
  - If options is a boolean, return it (treated as capture).
  - Otherwise, extract `capture`, `once`, `passive`, and `signal` from the dictionary properties.

Nuance Or Contradictions
- The text notes that while an `EventListener` object is passed to the callback, the "event listener" concept is broader, encompassing state flags like `capture` and `passive` which are not part of the `EventListener` interface itself.
- Performance optimization relies on observing listener presence; specifically, touch events can block asynchronous scrolling unless all listeners are passive, in which case the event becomes uncancelable.
- The default passive value logic is complex: it defaults to true for certain touch/mousewheel events on specific target contexts (Window, node, document, body), otherwise false.

Candidate Wiki Hints
- **EventTarget Interface**: A foundational page explaining what an EventTarget is and its core methods.
- **addEventListener Options**: Documentation detailing the difference between boolean and object options, specifically `passive` and `once`.
- **Performance: Observing Listeners**: A section discussing why observing listeners is necessary for efficiency and how passive listeners mitigate blocking behavior in scrolling events.

## chunk-04

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

## chunk-05

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

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
## Chunk Context
**Heading path:** Text: Why yes.⏎ > 4.2. Node tree > 4.2.3. Mutation algorithms > 4.2.4. Mixin NonElementParentNode
**Line range:** 1878–2271

## Local Summary
This chunk details the DOM mutation algorithms for inserting, moving, and replacing nodes within a parent node. It defines strict pre-insertion validity checks (e.g., hierarchy constraints, node types) to prevent invalid tree structures. The text distinguishes between "insertion steps" (which must not execute JavaScript or modify the tree) and "post-connection steps" (which handle side effects like style application and script execution). It also covers shadow DOM slot assignment, custom element lifecycle callbacks (`connectedCallback`, `disconnectedCallback`), and observer notifications. Finally, it introduces the `NonElementParentNode` mixin, which exposes `getElementById()` on `Document` and `DocumentFragment` but not on regular elements for web compatibility reasons.

## Key Claims
- **Pre-insert Validity:** Before inserting a node, algorithms must verify hierarchy constraints (e.g., no circular ancestry, correct parent types). Violations throw `HierarchyRequestError` or `NotFoundError`.
- **Separation of Concerns:** Insertion steps modify the tree structure without executing scripts. Post-connection steps handle side effects asynchronously after the structural change is complete.
- **Atomicity:** Batch insertions (like via a `DocumentFragment`) ensure all major side effects occur after the entire batch is inserted into the tree, maintaining consistency for observers and custom elements.
- **Shadow DOM Integration:** Algorithms account for shadow roots, slot assignment, and scoped document sets during node movement or insertion.
- **Observer Management:** Tree mutation records are queued to notify observers of structural changes. Observer lists are updated when nodes are moved or removed.

## Entities And Concepts
- **DOMException:** Exception types like `HierarchyRequestError` and `NotFoundError` used for invalid operations.
- **DocumentFragment:** A node type optimized for batch insertion; its children are detached before re-insertion into a parent.
- **Insertion Steps / Post-Connection Steps:** Distinct phases in the mutation algorithm; insertion steps modify structure, post-connection steps handle side effects (e.g., CSS application, script execution).
- **Slot Assignment:** Mechanism for mapping slottable nodes to named slots in shadow DOM hosts.
- **Custom Element Callbacks:** `connectedCallback` triggers on connection; `disconnectedCallback` triggers on disconnection. `connectedMoveCallback` is queued if a custom element moves into a connected parent.
- **Live Range:** A concept for tracking offset-based ranges (e.g., in `DocumentFragment` or text nodes) that need adjustment when nodes are inserted.
- **NonElementParentNode:** A mixin providing the `getElementById()` method, available on `Document` and `DocumentFragment` but not `Element`.

## Procedures And API Details
**Pre-insert Validity Checks (for inserting node into parent):**
1. Validate parent type (`Document`, `DocumentFragment`, or `Element`).
2. Ensure no circular ancestry between `node` and `parent`.
3. Verify `child`'s parent matches `parent` if `child` is non-null.
4. Confirm `node` type allows insertion (e.g., not a forbidden combination of Text/Doctype contexts).
5. Handle specific cases for `DocumentFragment`, `Element`, and `DocumentType` nodes regarding element children and doctypes.

**Insertion Algorithm:**
1. Validate pre-insertion.
2. Adjust live range offsets if `child` exists.
3. Adopt `node` into the parent's document.
4. Append or insert `node` before the specified `child`.
5. Handle slot assignment for shadow hosts.
6. Queue tree mutation records and run children changed steps.
7. Collect connected descendants to run post-connection steps safely (avoiding tree traversal during side effects).

**Move Algorithm:**
1. Verify roots match (same shadow-including root).
2. Check for circular ancestry between `node` and `newParent`.
3. Remove `node` from old parent, updating live ranges and slot assignments.
4. Insert `node` into new parent.
5. Queue mutation records for both old and new parents.
6. Trigger moving steps on descendants and enqueue `connectedMoveCallback` if applicable.

**Replace Algorithm:**
1. Validate pre-insertion constraints (similar to insert).
2. Determine reference child for insertion point.
3. Remove the existing `child` node.
4. Insert the new `node` at the reference position.
5. Queue mutation record with added and removed nodes.

**Remove Algorithm:**
1. Validate parent-child relationship.
2. Run pre-remove steps (live ranges, iterators).
3. Remove node from parent's children.
4. Handle slot assignments and shadow root checks.
5. Run removing steps on descendants.
6. Queue `disconnectedCallback` for custom elements if parent is connected.
7. Update observer lists for subtree observers.
8. Queue tree mutation record.

**NonElementParentNode API:**
- **Method:** `getElementById(DOMString elementId)`
- **Availability:** Included in `Document` and `DocumentFragment`. Not included in `Element`.
- **Behavior:** Returns the first descendant element with the matching ID in tree order, or `null` if none exists.

## Nuance Or Contradictions
- **Pre-insert vs. Replace:** The validity checks for replacing a node differ slightly from pre-insertion (e.g., checking if a doctype follows a specific child). This distinction ensures that replacement operations respect the current tree state more strictly than simple insertion.
- **SuppressObservers Flag:** Algorithms support an optional `suppressObservers` flag to batch mutations without immediate notification, which is crucial for performance but requires careful handling of mutation records afterward.
- **Web Compatibility Constraint:** The `getElementById()` method is intentionally restricted to `Document` and `DocumentFragment` via the `NonElementParentNode` mixin to prevent potential conflicts or unintended behavior if exposed on all elements.

## Candidate Wiki Hints
- **DOM Mutation Algorithms Overview:** A summary page explaining the separation between insertion steps, post-connection steps, and removing steps.
- **Pre-insert Validity Rules:** A detailed reference for the specific checks performed before adding nodes to a DOM tree.
- **Shadow DOM Slot Assignment:** Documentation on how slot assignment interacts with node insertion, movement, and replacement.
- **Custom Element Lifecycle in Mutation:** How `connectedCallback`, `disconnectedCallback`, and `connectedMoveCallback` are triggered during DOM mutations.
- **NonElementParentNode Mixin:** A page explaining why `getElementById()` is only available on certain parent nodes.

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
This chunk covers sections 4.2.5 through 4.3 of the DOM standard, detailing mixins for document and shadow root handling, parent node manipulation (including `prepend`, `append`, `replaceChildren`), sibling navigation (`previousElementSibling`), child node insertion/removal (`before`, `after`, `remove`), slottable elements, old-style collections (`NodeList`, `HTMLCollection`), and the mechanics of mutation observers.

Local Summary
The standard defines specific interface mixins to share functionality between `Document` and `ShadowRoot`. It details how nodes are converted into lists, how parent nodes manage their children via new methods like `prepend` and `append`, and how to navigate siblings excluding doctypes. It clarifies the behavior of legacy collections versus modern iterables and explains the microtask queue mechanism used by mutation observers to notify callbacks without losing data during subtree removals.

Key Claims
- The `DocumentOrShadowRoot` mixin provides a `customElementRegistry` attribute shared by both `Document` and `ShadowRoot`.
- New parent node methods (`prepend`, `append`, `replaceChildren`) automatically convert string arguments into `Text` nodes.
- These mutation methods throw a "HierarchyRequestError" DOMException if tree constraints are violated.
- The `previousElementSibling` and `nextElementSibling` attributes are excluded from `DocumentType` nodes for web compatibility reasons.
- `HTMLCollection` is described as a historical artifact; new API designers should use `sequence<T>` instead.
- Mutation observers operate via a microtask queue to ensure callbacks are invoked even if the DOM changes during notification.

Entities And Concepts
- **Mixin**: A mechanism to add shared attributes and methods to interfaces (e.g., `DocumentOrShadowRoot`, `ParentNode`).
- **CustomElementRegistry**: The registry object for custom elements, accessible via the mixin.
- **HierarchyRequestError**: A DOMException thrown when structural constraints are violated during manipulation.
- **NodeList vs HTMLCollection**: Distinguishes between live collections of nodes and historical element-only collections with named item lookup.
- **MutationObserver Microtask**: The internal queue mechanism ensuring observer callbacks fire correctly despite DOM mutations.

Procedures And API Details
- **Parent Node Conversion**: Strings in a list are replaced with `Text` nodes; single-node lists return the node directly, while multiple nodes are wrapped in a `DocumentFragment`.
- **Prepend/Append Logic**: Inserts converted nodes before the first child or after the last child respectively.
- **MoveBefore Logic**: Moves a node into a parent before a specified reference child (or the last child if none specified), preserving state.
- **NamedItem Lookup**: Returns the first element with a matching ID or `name` attribute (in HTML namespace) that hasn't been returned previously in the iteration.
- **Observer Notification Steps**: Clone pending observers, empty the queue, remove transient observers from nodes, invoke callbacks with records, and fire `slotchange` events.

Nuance Or Contradictions
- **Doctype Siblings**: The standard explicitly notes that sibling element attributes are not exposed on doctypes to maintain web compatibility, preventing potential inconsistencies.
- **Live vs Static Collections**: While most collections must be live, the text implies a distinction exists for specific cases where a snapshot might be acceptable unless otherwise stated (though the default requirement is liveness).
- **Legacy Artifacts**: The standard explicitly advises against using `HTMLCollection` in new designs, marking it as a legacy artifact to be phased out.

Candidate Wiki Hints
- **DocumentOrShadowRoot Mixin**: A page explaining how custom element registries are shared between the main document and shadow roots.
- **ParentNode Methods Deep Dive**: Documentation on the new `prepend`, `append`, and `replaceChildren` methods, including their error handling and string-to-node conversion behavior.
- **MutationObserver Microtask Queue**: An explanation of the internal microtask mechanism that prevents observer callbacks from being skipped during rapid DOM changes.

## chunk-08

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

## chunk-09

---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

Chunk Context
The chunk covers section 4.4, "Interface Node," from the DOM Standard. It details the `Node` interface definition, including its attributes, methods, and associated dictionaries. The text explains the behavior of node types (element, attribute, text, etc.), their properties like `nodeType`, `nodeName`, `baseURI`, and relationships such as `parentNode`, `childNodes`, and siblings. It also covers utility methods for cloning nodes (`cloneNode`), comparing positions within the document tree (`compareDocumentPosition`), and manipulating text content (`textContent`, `nodeValue`).

Local Summary
This section defines the `Node` interface, which is abstract and implemented by all DOM nodes. It lists constants for node types (e.g., `ELEMENT_NODE`, `TEXT_NODE`) and provides getter/setter logic for attributes like `nodeType`, `nodeName`, `baseURI`, `isConnected`, `ownerDocument`, and traversal properties like `parentNode`, `firstChild`, `previousSibling`. Methods include `getRootNode()`, `hasChildNodes()`, `appendChild()`, `removeChild()`, `cloneNode()`, `isEqualNode()`, `compareDocumentPosition()`, and `contains()`. Special attention is given to text content handling via `textContent` and the normalization of text nodes using `normalize()`. The cloning algorithm accounts for custom element registries and shadow roots. Node equality checks are defined based on structural properties like namespace, local name, attribute lists, and children.

Key Claims
- The `Node` interface is abstract; direct instances cannot be obtained.
- Every node has an associated "node document," set upon creation and potentially changed via the adopt algorithm.
- Node types are represented by unsigned short constants (e.g., `ELEMENT_NODE = 1`, `ATTRIBUTE_NODE = 2`).
- The `nodeName` property returns specific strings for different node types (e.g., "#text" for exclusive Text nodes, "null" or qualified names for others).
- The `baseURI` getter returns the serialized document base URL of the node's document.
- `isConnected` returns true if the node is part of a live tree.
- `getRootNode()` returns the root; with `{ composed: true }`, it includes shadow roots.
- `cloneNode(subtree)` creates a copy, including descendants if `subtree` is true. It handles custom element registries and shadow roots specifically.
- `isEqualNode(otherNode)` compares structural equality (interfaces, attributes, children) but does not consider live ranges or text offsets.
- `isSameNode(otherNode)` checks reference equality (strict `===`).
- `compareDocumentPosition(other)` returns a bitmask indicating relative position (preceding, following, contained by, contains).
- `normalize()` removes empty exclusive Text nodes and concatenates contiguous ones.

Entities And Concepts
- **Node**: Abstract interface for DOM nodes.
- **Node Types**: Element, Attr, Text (exclusive), CDATASection, ProcessingInstruction, Comment, Document, DocumentType, DocumentFragment, Notation (legacy).
- **Node Type Constants**: `ELEMENT_NODE`, `ATTRIBUTE_NODE`, `TEXT_NODE`, etc.
- **Traversal Properties**: `parentNode`, `parentElement`, `childNodes`, `firstChild`, `lastChild`, `previousSibling`, `nextSibling`.
- **Text Content**: Managed via `textContent` getter/setter and `nodeValue`.
- **Cloning**: `cloneNode()` with support for shadow roots and custom element registries.
- **Equality**: `isEqualNode()` (structural), `isSameNode()` (reference).
- **Position**: `compareDocumentPosition()`, `contains()`.
- **Normalization**: `normalize()` method for cleaning up text nodes.

Procedures And API Details
- **`getRootNode(options)`**: Returns the root node. If `options.composed` is true, returns the shadow-including root.
- **`cloneNode(subtree = false)`**: Clones a node and optionally its descendants. Handles shadow hosts and custom element registries. Throws "NotSupportedError" if called on a shadow root.
- **`isEqualNode(otherNode)`**: Returns true if `otherNode` is non-null and structurally equal to the current node (same interfaces, attributes, children).
- **`isSameNode(otherNode)`**: Legacy alias for strict equality (`===`).
- **`compareDocumentPosition(other)`**: Returns a bitmask combining flags like `DOCUMENT_POSITION_PRECEDING`, `DOCUMENT_POSITION_FOLLOWING`, etc. Handles attributes as preceding their element's children.
- **`contains(other)`**: Returns true if `other` is an inclusive descendant (or null).
- **`normalize()`**: Iterates descendants; removes empty exclusive Text nodes and merges contiguous ones, updating live ranges.
- **`textContent` Setter**: If value is null, treats as empty string; delegates to internal logic based on node type (Element, Attr, CharacterData).

Nuance Or Contradictions
- **Node Document**: A node's document is immutable except via the adopt algorithm. For documents themselves, `ownerDocument` returns null.
- **Attribute Handling in Position**: Attributes are treated as preceding their element's children in `compareDocumentPosition`, even though they don't participate in the same tree structure for traversal purposes.
- **Cloning Shadow Roots**: If a shadow root is cloned and its `clonable` flag is true, the clone includes a new shadow root with specific properties copied from the original (mode, serializable, delegates focus, etc.).
- **Legacy Constants**: Some node types like `ENTITY_REFERENCE_NODE`, `ENTITY_NODE`, and `NOTATION_NODE` are marked as legacy.
- **Text Content Logic**: The logic for getting/setting text content differs between Element descendants (descendant text), Attr (value), and CharacterData (data).

Candidate Wiki Hints
- **Node Interface Overview**: A page explaining the `Node` interface, its abstract nature, and common usage patterns.
- **Node Types and Constants**: Documentation on node type constants (`ELEMENT_NODE`, `TEXT_NODE`) and their corresponding types.
- **Cloning Nodes**: Deep dive into `cloneNode()`, including handling of shadow DOMs and custom elements.
- **Document Position**: Explanation of the bitmask returned by `compareDocumentPosition()` and how to interpret it.
- **Text Manipulation**: Guide on using `textContent`, `nodeValue`, and `normalize()` for text content management.

## chunk-10

---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
The source text defines algorithms for locating namespaces and prefixes on DOM nodes (Interface Node), describes behavior for `HTMLCollection` lists of elements based on qualified names, namespaces, and local names, and specifies filtering logic for class names. It covers steps for `Element`, `Document`, `Attr`, and other node types, including fallback to parent elements.

Local Summary
This chunk details the logic for resolving namespace prefixes and URIs across different DOM interfaces (`Element`, `Document`, `Attr`, etc.) by checking specific attributes, namespaces, or delegating to a parent element. It also outlines how to construct `HTMLCollection` objects that filter descendant elements based on qualified names, namespace/local name combinations, and class names, with specific handling for HTML documents and quirks mode.

Key Claims
- Namespace prefix resolution involves checking the element's own namespace, an "xmlns" attribute, or delegating to the parent element.
- `lookupNamespaceURI` converts an empty string prefix to null before resolving the namespace.
- `isDefaultNamespace` compares a given namespace against the result of locating a namespace with a null prefix.
- DOM mutation methods (`insertBefore`, `appendChild`, `replaceChild`, `removeChild`) delegate to underlying pre-insertion/pre-removal logic.
- `HTMLCollection` filters for qualified names distinguish between HTML namespace elements (case-insensitive matching) and non-HTML namespace elements.
- Class name filtering requires all specified classes to be present on an element, with ASCII case-insensitive comparison in "quirks" mode.

Entities And Concepts
- Interface Node: The interface defining methods for locating namespaces and prefixes.
- `locate a namespace prefix`: Algorithm to find the prefix string for a given namespace URI.
- `locate a namespace`: Algorithm to find the namespace URI for a given prefix string.
- `lookupPrefix`: Method returning the prefix for a specific namespace.
- `lookupNamespaceURI`: Method returning the namespace URI for a specific prefix.
- `isDefaultNamespace`: Method checking if a namespace is the default (empty) namespace.
- `HTMLCollection`: A collection object returned by element listing algorithms.
- Qualified Name: Combination of namespace and local name used to identify elements.

Procedures And API Details
- **Algorithm for locating namespace prefix:**
  1. Return the element's prefix if namespace matches and prefix is non-null.
  2. Return the local name of an attribute with namespace "xmlns" and value matching namespace.
  3. Recursively call on the parent element if it exists.
  4. Return null otherwise.
- **Algorithm for locating namespace URI:**
  1. Set empty string prefix to null.
  2. Check specific prefixes ("xml", "xmlns") or attributes with local name matching the prefix.
  3. Return the namespace value if non-empty, or null otherwise.
  4. Delegate to parent element if current node has no parent.
- **Algorithm for `HTMLCollection` by qualified name:**
  1. If argument is "*", return all descendant elements.
  2. For HTML documents, filter descendants where namespace is HTML (ASCII lowercase) or not HTML matching the qualified name.
  3. Otherwise, match any element with the specified qualified name.
- **Algorithm for `HTMLCollection` by namespace and local name:**
  1. Convert empty strings to null.
  2. Handle wildcard (*) for namespace or local name individually or together.
  3. Return a collection filtering descendants matching the specific namespace/local name combination.
- **Algorithm for `HTMLCollection` by class names:**
  1. Parse class names into an ordered set.
  2. If empty, return empty collection.
  3. Return collection where elements have all classes in the set; comparisons are ASCII case-insensitive in "quirks" mode.

Nuance Or Contradictions
- The logic for locating a namespace prefix prioritizes the element's own namespace attribute over inherited context from parents unless specific conditions (like an "xmlns" attribute) are met.
- `HTMLCollection` caching is mentioned: the same object may be returned for identical arguments as long as the document type hasn't changed.
- Case sensitivity for qualified name matching depends on whether the document is an HTML document and the specific namespace involved (HTML namespace requires ASCII lowercase comparison).

Candidate Wiki Hints
- DOM Namespace Resolution Algorithms
- Interface Node Methods (`lookupPrefix`, `lookupNamespaceURI`)
- HTMLCollection Filtering Strategies
- Handling Class Names in DOM Queries

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context

This chunk details section **4.5. Interface Document** of the DOM Standard specification. It defines the `Document` interface, its attributes (such as `URL`, `compatMode`, and `doctype`), and creation methods (`createElement`, `createTextNode`, etc.). It also covers `XMLDocument`, element creation options (including custom element registries), namespace handling, cloning/importing nodes, and the adoption algorithm for moving nodes between documents.

## Local Summary

The section establishes the `Document` interface as a subclass of `Node`. It defines default values for document properties (e.g., encoding is UTF-8, type is "xml", mode defaults to "no-quirks"). The chunk details how to create elements with specific namespaces or local names, handling legacy string options and custom element registries. It explains the behavior of `getElementsByTagName` and `getElementsByClassName`, including case-sensitivity nuances in HTML documents. Finally, it outlines the algorithms for importing nodes (deep cloning), adopting nodes from other documents, and managing custom element registries within shadow roots during adoption.

## Key Claims

- A document's default encoding is UTF-8, content type is "application/xml", URL is "about:blank", origin is opaque, type is "xml", mode is "no-quirks", declarative shadow roots are false, and the custom element registry is null.
- The `Document` interface has legacy aliases: `charset` and `inputEncoding` map to `characterSet`.
- `compatMode` returns "BackCompat" if the document mode is "quirks"; otherwise "CSS1Compat".
- `createElement(localName)` lowercases `localName` in HTML documents; `createElementNS` handles namespace prefixes.
- `getElementsByClassName` interprets the argument as a space-separated list of classes and requires all specified classes to be present on an element.
- Cloning a document or shadow root via `importNode` throws a "NotSupportedError".
- Adopting a node into a new document updates the `node document` for the node and its inclusive descendants (shadow-including tree order).

## Entities And Concepts

- **Document Interface**: The base interface for documents, inheriting from `Node`.
- **XMLDocument**: An interface extending `Document` for XML content.
- **Modes**: "no-quirks", "quirks", and "limited-quirks" (formerly "standards mode" and "almost standards mode").
- **CustomElementRegistry**: Used to define custom elements, passed via options during creation or import.
- **Adoption Algorithm**: The process of moving a node from one document to another, updating references for the node and its inclusive descendants.
- **Namespace Handling**: Rules for `createElementNS` regarding prefixes, empty namespaces, and reserved names like "xmlns".

## Procedures And API Details

### Document Attributes
- `implementation`: Returns the associated `DOMImplementation`.
- `URL` / `documentURI`: Return the document's URL.
- `compatMode`: Returns "BackCompat" (quirks mode) or "CSS1Compat".
- `characterSet` / `charset` / `inputEncoding`: Return the encoding name.
- `doctype`: Returns the document type declaration or null.
- `documentElement`: Returns the root element of the document.

### Creation Methods
- `createElement(localName, options)`: Creates an element. In HTML documents, local names are lowercased. Options can specify a custom element registry or a customized built-in element via `is`.
- `createElementNS(namespace, qualifiedName, options)`: Creates an element in a specific namespace. Handles prefix extraction and validates namespace usage.
- `createDocumentFragment()`: Returns a new fragment node.
- `createTextNode(data)`, `createCDATASection(data)`, `createComment(data)`, `createProcessingInstruction(target, data)`: Return corresponding node types.

### Collection Methods
- `getElementsByTagName(qualifiedName)`: Returns an `HTMLCollection` of matching elements. Case-insensitive matching for HTML namespace elements in HTML documents.
- `getElementsByTagNameNS(namespace, localName)`: Filters by namespace and local name. Wildcards (`*`) allow partial matching.
- `getElementsByClassName(classNames)`: Returns elements containing all classes in the space-separated string argument.

### Cloning and Adoption
- `importNode(node, options)`: Creates a deep copy of a node (descendants included if `selfOnly` is false or option is true). Throws "NotSupportedError" for documents or shadow roots. Supports setting custom element registry on cloned elements.
- `adoptNode(node)`: Moves a node to the current document. Throws errors for documents or shadow roots with incompatible registries.

### Flattening Options
The `flatten element creation options` algorithm extracts `customElementRegistry` and `is` from an options object or string, validating that if a registry is provided, it matches the document's registry or has a scoped `is` flag set to false.

## Nuance Or Contradictions

- **Mode Renaming**: "Standards mode" and "almost standards mode" were renamed to "no-quirks" and "limited-quirks" respectively due to semantic issues and Ian Hickson's veto.
- **Case Sensitivity in `getElementsByTagName`**: In an HTML document, `<FOO>` (non-HTML namespace) and `<foo>` (HTML namespace) are matched separately from `<FOO>` (HTML namespace). The method matches case-insensitively only within the HTML namespace for HTML documents.
- **Class Matching Logic**: `getElementsByClassName("aaa bbb")` matches elements having both "aaa" and "bbb". Spaces are delimiters; commas or other characters do not split classes in the same way (e.g., `"aaa,bbb"` returns no nodes if no element has exactly those comma-separated classes).
- **Legacy String Options**: The `options` parameter for `createElement` and `createElementNS` can be a string for web compatibility, though the primary definition uses dictionaries.

## Candidate Wiki Hints

1. **Document Interface Attributes**: A page summarizing `compatMode`, `URL`, `characterSet`, and other read-only properties of the `Document` interface.
2. **Element Creation Options**: Documentation on how to use `customElementRegistry` and `is` in `createElement` options, including error cases like "NotSupportedError".
3. **getElementsByClassName Behavior**: A guide explaining space-separated class lists, case sensitivity, and examples of matching logic (e.g., `"aaa bbb"` vs `"aaa,bbb"`).
4. **Node Import and Adoption**: An explanation of the difference between `importNode` (cloning) and `adoptNode` (moving), including restrictions on shadow roots and document nodes.

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context
This chunk covers DOM operations related to custom element registries, node adoption, document creation, and shadow roots. It includes definitions for `DOMImplementation`, `DocumentType`, `DocumentFragment`, and `ShadowRoot` interfaces, along with algorithms for attribute creation, event creation, range creation, node iteration, and tree walking.

# Local Summary
The text details how custom element registries are managed within documents and elements, specifically handling scoped vs. global registries. It outlines the `adoptNode()` method steps for moving nodes between documents, throwing specific errors for unsupported types like documents or shadow roots. The chunk also defines the `DOMImplementation` interface methods (`createDocumentType`, `createDocument`, `createHTMLDocument`) and notes that `hasFeature()` is obsolete. Definitions for `DocumentFragment` (host concept) and `ShadowRoot` (attributes like mode, delegatesFocus, host) are provided, including traversal orders for shadow-including trees.

# Key Claims
- A document's effective global custom element registry is the document's custom element registry if it is a global registry; otherwise, it is null.
- The `adoptNode()` method throws a "NotSupportedError" if the node is a document and a "HierarchyRequestError" if it is a shadow root.
- The `hasFeature()` method on `DOMImplementation` always returns true and is no longer reliable for feature detection; it exists for backward compatibility.
- Shadow roots have an associated mode ("open" or "closed") and an associated host which is never null.
- In shadow-including tree order, traversal includes a depth-first walk of the shadow root's node tree immediately after encountering the shadow host.

# Entities And Concepts
- **Custom Element Registry**: Can be global or scoped; determines how custom elements are registered.
- **DOMImplementation**: Interface for creating document types and documents.
  - `createDocumentType()`: Creates a doctype.
  - `createDocument()`: Creates an XMLDocument.
  - `createHTMLDocument()`: Creates an HTML document with basic structure.
- **DocumentType**: Represents the doctype of a document (name, publicId, systemId).
- **DocumentFragment**: A node used for fragment manipulation; has an associated host.
- **ShadowRoot**: A special type of DocumentFragment acting as the root of a shadow tree.
  - Attributes: `mode`, `delegatesFocus`, `slotAssignment`, `clonable`, `serializable`, `host`.
- **Shadow-including Tree Order**: Traversal order including shadow DOM subtrees.

# Procedures And API Details
- **`createAttribute(localName)`**: Returns a new attribute; throws "InvalidCharacterError" if localName is invalid. In HTML documents, the local name is lowercased.
- **`createAttributeNS(namespace, qualifiedName)`**: Validates and extracts namespace/prefix/local name; returns a new attribute.
- **`createEvent(interface)`**: Creates an event object. Supports interfaces like `BeforeUnloadEvent`, `CustomEvent`, `MouseEvent`, etc. Throws "NotSupportedError" if the interface is not supported or exposed. Initializes type, timeStamp, and isTrusted attributes.
- **`createRange()`**: Returns a new live range starting and ending at index 0 of the current context.
- **`createNodeIterator(root, whatToShow, filter)`**: Creates a NodeIterator with specified root, visibility filter, and optional filter.
- **`createTreeWalker(root, whatToShow, filter)`**: Creates a TreeWalker with specified root, visibility filter, and optional filter.
- **`adoptNode(node)`**: Moves a node into the current document. Throws errors for documents or shadow roots. Returns the adopted node.

# Nuance Or Contradictions
- The `hasFeature()` method is described as "useless" and always returning true, yet it remains part of the interface specification to maintain backward compatibility with old pages.
- Event creation via `createEvent()` is noted as deprecated in favor of using constructors directly (e.g., `new MouseEvent()`).
- Shadow roots have an associated custom element registry which is initially null, but can be set; this interacts with the "keep custom element registry null" boolean flag relevant only for declarative shadow roots.

# Candidate Wiki Hints
- **Page: Custom Element Registry Scope** – Explaining global vs. scoped registries and how `adoptedCallback` reactions handle registry changes.
- **Page: DOMImplementation Interface** – Documenting the creation of documents and doctypes, including content type determination based on namespace.
- **Page: ShadowRoot Attributes** – Detailing `mode`, `delegatesFocus`, `host`, and other shadow root specific attributes.
- **Page: Event Creation API** – Listing supported event interfaces and initialization steps for `createEvent()`.

## chunk-13

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

## chunk-14

---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

Chunk Context
- **Heading**: `8. Return document. > 4.9. Interface Element`
- **Line Range**: 4931–5225
- **Scope**: Defines attribute manipulation methods (`setAttribute`, `removeAttribute`, etc.), shadow DOM attachment logic, selector matching utilities, and adjacent insertion mechanisms for DOM elements.

Local Summary
This chunk details the API for managing attributes on a generic `element` interface, including validation rules for qualified names and trusted types integration. It specifies the behavior of the shadow DOM host methods (`attachShadow`, `shadowRoot`) and lists valid host element names. Additionally, it covers selector-based traversal (`closest`, `matches`) and list retrieval via tag/class names. Finally, it outlines the logic for inserting nodes or text adjacent to an element using specific string positions.

Key Claims
- **Attribute Validation**: `setAttribute` throws an "InvalidCharacterError" if the qualified name is not a valid attribute local name; in HTML documents, this name is normalized to ASCII lowercase before storage.
- **Trusted Types**: The `get trusted type compliant attribute value` step is invoked during attribute setting to ensure content safety.
- **Shadow Host Restrictions**: Only elements with specific local names (e.g., "article", "div", custom element names) can act as shadow hosts; attaching a shadow root to an invalid host throws a "NotSupportedError".
- **Closed Mode Behavior**: The `shadowRoot` getter returns `null` if the shadow root's mode is "closed".
- **Adjacent Insertion**: The `insertAdjacentElement` and `insertAdjacentText` methods rely on ASCII case-insensitive matching of strings like "beforebegin" or "afterend" to determine insertion points.

Entities And Concepts
- **Interface Element**: The generic host for DOM manipulation methods.
- **Attribute List**: A collection associated with an element, returned by the `attributes` getter as a `NamedNodeMap`.
- **Shadow Root**: An encapsulated tree attached to a shadow host element via `attachShadow`.
- **CustomElementRegistry**: Used to manage custom element definitions and shadow DOM capabilities.
- **Selectors4**: Referenced for parsing selector strings within `closest` and `matches`.

Procedures And API Details
- **setAttributeNS**: Validates namespace and local name extraction, checks trusted type compliance, and updates or creates the attribute.
- **attachShadow(init)**: Checks if the element is in the HTML namespace and has a valid shadow host name. It verifies custom element definitions for "disable shadow" flags. If a shadow root exists, it validates mode compatibility before clearing children and updating properties.
- **closest(selectors)**: Parses the selector string; throws "SyntaxError" on failure. Iterates through inclusive ancestors in reverse tree order to find the first match.
- **insertAdjacentElement(where, element)**: Maps `where` strings to specific insertion logic (e.g., inserting before the parent if "beforebegin" and no parent exists). Throws "SyntaxError" for unrecognized strings.

Nuance Or Contradictions
- **Qualified Name vs. Local Name**: The parameter name `qualifiedName` in `setAttribute` is misleading; it acts as a local name when adding a new attribute but represents the full qualified name if an attribute with that name already exists.
- **Return Values**: Methods like `removeAttribute` and `insertAdjacentText` return `undefined` or nothing, which differs from methods that return booleans or nodes, sometimes due to historical design before standardization.

Candidate Wiki Hints
- **DOM Attribute Manipulation**: Guide for using `setAttribute`, `removeAttribute`, and `toggleAttribute`.
- **Shadow DOM Host Requirements**: List of valid element names and conditions for attaching a shadow root.
- **Adjacent Insertion Patterns**: How to use `insertAdjacentElement` with different position strings.

## chunk-15

---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
- Heading path: 8. Return document. > 4.9. Interface Element > 4.9.1. Interface NamedNodeMap
- Line range: 5227–5703
- Topics covered: NamedNodeMap, Attr, CharacterData (Text, CDATASection, ProcessingInstruction, Comment), and DOM Ranges (AbstractRange, StaticRange).

Local Summary
This chunk defines the API for attribute collections (`NamedNodeMap`), attribute nodes (`Attr`), character data containers (`CharacterData`), specific text-like nodes (`Text`, `CDATASection`, `ProcessingInstruction`, `Comment`), and the mechanism for selecting node tree content via `Range` objects (including immutable `StaticRange`).

Key Claims
- A `NamedNodeMap` acts as a collection of attributes associated with an element, exposing methods to get/set/remove items by index or name.
- `Attr` nodes are distinct from IDL attributes and include properties for namespace, prefix, local name, value, and owner element.
- `CharacterData` is an abstract interface implemented by `Text`, `CDATASection`, `ProcessingInstruction`, and `Comment`, allowing mutable string manipulation (`data`, `length`, `substringData`).
- `Range` objects represent a sequence of content between two boundary points (start/end nodes and offsets). They are "live" and update on mutations, whereas `StaticRange` does not.
- Mutating the node tree requires updating all affected live ranges, which can be expensive.

Entities And Concepts
- NamedNodeMap: Interface for accessing an element's attributes.
- Attr: Represents a content attribute of an element.
- CharacterData: Abstract interface for nodes containing character data.
- Text: A `CharacterData` node representing textual content (excluding CDATA).
- CDATASection: A `Text` node specifically for CDATA sections.
- ProcessingInstruction: A `CharacterData` node with a target.
- Comment: A `CharacterData` node for XML comments.
- Range: An object representing a selection within the DOM tree.
- StaticRange: An immutable range that does not update on DOM mutations.
- Boundary point: A tuple of a node and an offset defining a position in the tree.

Procedures And API Details
- NamedNodeMap methods:
  - `getNamedItem(qualifiedName)`: Returns an `Attr` by qualified name.
  - `setNamedItem(attr)`: Sets an attribute; throws if invalid.
  - `removeNamedItem(qualifiedName)`: Removes and returns the attribute or throws "NotFoundError".
- Attr properties:
  - `namespaceURI`, `prefix`, `localName`, `name`, `value`, `ownerElement`.
  - `specified` always returns true (noted as useless).
- CharacterData methods:
  - `appendData()`, `insertData()`, `deleteData()`, `replaceData()`: Mutate the underlying string.
  - `substringData(offset, count)`: Returns a substring without mutating.
- Text methods:
  - `splitText(offset)`: Splits the text node at the given offset.
  - `wholeText`: Returns concatenated data of contiguous sibling Text nodes.
- Range concepts:
  - A range is defined by `(startContainer, startOffset)` and `(endContainer, endOffset)`.
  - `collapsed` is true if start equals end.

Nuance Or Contradictions
- The `specified` property on `Attr` always returns `true`, which the source notes as "useless".
- Attributes (e.g., `src`, `alt`) cannot be represented by a `Range`; ranges apply only to nodes.
- Live ranges attempt to maintain validity during mutations but can be modified themselves if the content they represent changes.

Candidate Wiki Hints
- Page: **DOM NamedNodeMap** – Explains attribute collection mechanics and methods.
- Page: **DOM Attr Node** – Details attribute properties and namespace handling.
- Page: **DOM CharacterData Interface** – Covers text manipulation algorithms.
- Page: **DOM Range Selection** – Introduces live vs static ranges and boundary points.

## chunk-16

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

## chunk-17

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

## chunk-18

---
title: Chunk 18 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

**Heading Path:** 6. Traversal > 6.1. Interface NodeIterator, 6.2. Interface TreeWalker, 6.3. Interface NodeFilter, 7. Sets, 7.1. Interface DOMTokenList, 8. XPath
**Source File:** `raw/web/corpus-2026-05-18/060-dom-standard.md`
**Line Range:** 6288–6784

This chunk details the DOM traversal interfaces (`NodeIterator`, `TreeWalker`, `NodeFilter`) and the set-like interface `DOMTokenList`. It concludes with a note on the legacy XPath API.

# Local Summary

The document defines algorithms and IDL for tree traversal mechanisms. `NodeIterator` provides forward/backward traversal from a specific root with a filter, while `TreeWalker` maintains a current node position to traverse descendants, ancestors, siblings, or children based on filters. `NodeFilter` supplies the filtering logic (`acceptNode`) and bitmask constants (`whatToShow`). The chunk also covers `DOMTokenList`, an interface for managing sets of tokens (like class names), including methods to add, remove, toggle, and replace tokens with specific error handling for whitespace and empty strings. Finally, it notes that DOM Level 3 XPath APIs are legacy but maintained in the spec for future updates.

# Key Claims

- `NodeIterator` is created via `createNodeIterator()` on a `Document`.
- `TreeWalker` is created via `createTreeWalker()` on a `Document`.
- Both iterators use a filter (`NodeFilter`) and a `whatToShow` bitmask to determine node visibility.
- The `detach()` method on `NodeIterator` exists for compatibility but performs no action; its functionality was removed.
- `DOMTokenList` methods throw `SyntaxError` for empty strings and `InvalidCharacterError` for tokens containing ASCII whitespace.
- XPath Level 3 APIs are considered legacy and not actively maintained, though definitions remain for potential future updates.

# Entities And Concepts

- **Interfaces:** `NodeIterator`, `TreeWalker`, `NodeFilter`, `DOMTokenList`.
- **Constants:** `FILTER_ACCEPT` (1), `FILTER_REJECT` (2), `FILTER_SKIP` (3), `SHOW_ALL`, `SHOW_ELEMENT`, `SHOW_TEXT`, etc.
- **Exceptions:** `SyntaxError` DOMException, `InvalidCharacterError` DOMException, `TypeError`.
- **Attributes/Properties:** `root`, `referenceNode`, `pointerBeforeReferenceNode`, `whatToShow`, `filter`, `currentNode`, `length`, `value`.
- **Methods:** `nextNode()`, `previousNode()`, `parentNode()`, `firstChild()`, `lastChild()`, `add()`, `remove()`, `toggle()`, `replace()`.

# Procedures And API Details

### NodeIterator Traversal Logic
To traverse, the algorithm iterates while true:
1. Determine direction (`next` or `previous`).
2. If moving forward and `pointerBeforeReferenceNode` is false, get the first following node. If true, move it to false and get the next node.
3. Filter the candidate node.
4. If filtered as `FILTER_ACCEPT`, break the loop.
5. Update `referenceNode` and return the node.

### TreeWalker Traversal Logic
- **Children:** Traverse first/last child. Skip children if filtered, moving to the next/previous sibling of the parent if necessary.
- **Siblings:** Move to next/previous sibling. If filtered, descend into children or move to the next sibling of the current node.
- **Ancestors:** Walk up the tree until a filter accepts a node or the root is reached.

### DOMTokenList Operations
- **add(tokens...):** Appends tokens if not present. Throws `SyntaxError` for empty strings, `InvalidCharacterError` for whitespace.
- **remove(tokens...):** Removes tokens if present. Same error handling as `add`.
- **toggle(token, force):** If `force` is omitted/false and token exists, removes it (returns false). If `force` is true/omitted and token does not exist, adds it (returns true).
- **replace(token, newToken):** Swaps tokens. Returns false if token doesn't exist; throws errors for empty/whitespace strings.

# Nuance Or Contradictions

- **detach() Methodality:** The `detach()` method on `NodeIterator` is preserved strictly for compatibility but currently does nothing, as its original functionality was removed from the spec logic.
- **DOMTokenList Name:** The text explicitly notes that the name "DOMTokenList" is an "unfortunate legacy mishap," suggesting a potential future rename or refactoring in related specifications (like HTML).
- **Update Steps Execution:** The standard update steps for `DOMTokenList` are not always executed for `toggle()` and `replace()` methods to ensure web compatibility, implying implementation-specific behaviors may differ from strict algorithmic definitions.
- **XPath Status:** While the XPath interface definitions are maintained in the DOM spec, the underlying functionality (DOM Level 3 XPath) is widely implemented but considered unmaintained and legacy.

# Candidate Wiki Hints

- **NodeIterator vs TreeWalker:** Create a comparison page explaining that `NodeIterator` traverses from a fixed root without changing its current position context during iteration, whereas `TreeWalker` maintains a mutable `currentNode` state allowing navigation to parents, children, and siblings dynamically.
- **DOMTokenList Token Validation:** Document the strict validation rules for `DOMTokenList`, specifically the prohibition of whitespace characters and empty strings, which distinguishes it from standard string lists.
- **Traversal Filter Constants:** Create a reference table for `NodeFilter` constants (`FILTER_ACCEPT`, `SHOW_ELEMENT`, etc.) and explain how bitwise operations combine `whatToShow` flags.

## chunk-19

---
title: Chunk 19 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk covers the DOM Level 3 XPath and XSLT interfaces, detailing `XPathResult`, `XPathExpression`, `XPathEvaluator`, and `XSLTProcessor`. It also notes that no known security or privacy considerations exist for this standard. The coverage spans sections 8.1 through 10.

## Local Summary
The specification defines the `XPathResult` interface with various result types (e.g., `ANY_TYPE`, `STRING_TYPE`, `UNORDERED_NODE_ITERATOR_TYPE`) and attributes for accessing values (`numberValue`, `stringValue`, `singleNodeValue`). The `XPathEvaluatorBase` mixin provides methods like `createExpression` and `evaluate`, noting that `Document` includes this mixin for historical reasons. The `XSLTProcessor` interface is defined for transforming XML documents, with methods such as `importStylesheet`, `transformToFragment`, and parameter management (`setParameter`, `clearParameters`).

## Key Claims
- The `XPathResult` interface supports multiple result types via constants like `ANY_TYPE = 0` through `FIRST_ORDERED_NODE_TYPE = 9`.
- The `createNSResolver(nodeResolver)` method is provided only for historical reasons and simply returns the provided node resolver.
- Both `XPathEvaluator` and `Document` instances allow access to XPath evaluation methods due to `Document` including `XPathEvaluatorBase`.
- The `XSLTProcessor` interface includes a constructor and methods for importing stylesheets, transforming documents/fragments, and managing parameters.
- No known security or privacy considerations are identified for the XSLT and XPath APIs in this standard.

## Entities And Concepts
- **XPathResult**: An interface representing the result of an XPath evaluation.
- **XPathExpression**: An interface representing a compiled XPath expression.
- **XPathEvaluatorBase**: A mixin providing core XPath evaluation functionality (`createExpression`, `evaluate`).
- **XPathEvaluator**: An interface that includes `XPathEvaluatorBase`.
- **XSLTProcessor**: An interface for performing XSLT transformations on XML documents.
- **XPathNSResolver**: A callback interface used to resolve namespace URIs during expression creation.

## Procedures And API Details
- **Creating an XPath Expression**: Use `XPathEvaluatorBase.createExpression(expression, resolver)` where the expression is a DOMString and resolver is optional.
- **Evaluating an XPath Expression**: Call `XPathResult.evaluate(expression, contextNode, resolver, type, result)`. The `type` argument defaults to `0` (`ANY_TYPE`).
- **Processing XSLT**: Instantiate `XSLTProcessor`, call `importStylesheet(Node)` to load stylesheets, then use `transformToFragment(source)` or `transformToDocument(source)`. Parameters are set via `setParameter(namespaceURI, localName, value)`.

## Nuance Or Contradictions
- The `createNSResolver` method is explicitly marked as existing only for historical reasons and acts as a pass-through (returns the input node).
- Historically, `XPathEvaluator` could be constructed directly; however, since `Document` includes `XPathEvaluatorBase`, one can also access XPath methods directly on a document object.

## Candidate Wiki Hints
- **Page**: XPath in DOM / XSLT Processing
  - **Content**: Overview of the XPath and XSLT interfaces, result types, and common usage patterns for XML transformation.

## chunk-20

---
title: Chunk 20 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
This chunk covers section "11. Historical" of the DOM Standard Living Specification. It documents interfaces and interface members that were removed from previous versions or iterations of the standard to streamline the API and align with modern usage patterns.

Local Summary
The document lists specific DOM interfaces (such as `DOMConfiguration` and `Entity`) and various member methods/properties on core nodes like `Document`, `Element`, and `Node` that are no longer part of the current standard. The section concludes with acknowledgments to contributors and copyright/license information for the Living Standard.

Key Claims
- Several interfaces have been removed from the standard entirely, including `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMImplementationSource`, `DOMLocator`, `DOMObject`, `DOMUserData`, `Entity`, `EntityReference`, `MutationEvent`, `MutationNameEvent`, `NameList`, `Notation`, `RangeException`, `TypeInfo`, and `UserDataHandler`.
- Specific members have been removed from existing interfaces:
  - **Attr**: `schemaTypeInfo`, `isId`
  - **Document**: `createEntityReference()`, `xmlEncoding`, `xmlStandalone`, `xmlVersion`, `strictErrorChecking`, `domConfig`, `normalizeDocument()`, `renameNode()`
  - **DocumentType**: `entities`, `notations`, `internalSubset`
  - **DOMImplementation**: `getFeature()`
  - **Element**: `schemaTypeInfo`, `setIdAttribute()`, `setIdAttributeNS()`, `setIdAttributeNode()`
  - **Node**: `isSupported`, `getFeature()`, `getUserData()`, `setUserData()`
  - **NodeIterator**: `expandEntityReferences`
  - **Text**: `isElementContentWhitespace`, `replaceWholeText()`
  - **TreeWalker**: `expandEntityReferences`
- The standard is written by Anne van Kesteren with substantial contributions from Aryeh Gregor and Ms2ger.
- Portions of the revision history related to custom elements are available in the w3c/webcomponents repository under the W3C Software and Document License.
- The work is licensed under a Creative Commons Attribution 4.0 International License, with source code portions under the BSD 3-Clause License.

Entities And Concepts
- DOM Standard (Living Standard)
- Removed Interfaces: `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMImplementationSource`, `DOMLocator`, `DOMObject`, `DOMUserData`, `Entity`, `EntityReference`, `MutationEvent`, `MutationNameEvent`, `NameList`, `Notation`, `RangeException`, `TypeInfo`, `UserDataHandler`
- Removed Members: `schemaTypeInfo`, `isId`, `createEntityReference()`, `xmlEncoding`, `xmlStandalone`, `xmlVersion`, `strictErrorChecking`, `domConfig`, `normalizeDocument()`, `renameNode()`, `entities`, `notations`, `internalSubset`, `getFeature()`, `setIdAttribute()`, `setIdAttributeNS()`, `setIdAttributeNode()`, `isSupported`, `getUserData()`, `setUserData()`, `expandEntityReferences`, `isElementContentWhitespace`, `replaceWholeText()`
- Contributors: Anne van Kesteren, Aryeh Gregor, Ms2ger

Procedures And API Details
- No new procedures are defined in this chunk; the focus is on historical removals.
- The `getFeature()` method was removed from `DOMImplementation` and `Node`.
- The `expandEntityReferences` member was removed from both `NodeIterator` and `TreeWalker`.

Nuance Or Contradictions
- None observed within this specific chunk; it serves purely as a record of deprecations/removals.

Candidate Wiki Hints
- Page: Historical DOM Interfaces (Summary of removed APIs)

## chunk-21

---
title: Chunk 21 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

This chunk covers the **Historical** section (§ 11) of the DOM Standard, listing deprecated or legacy interfaces and concepts. It also includes a glossary-style index of terms from earlier sections (e.g., AbortController, Event handling, Node manipulation) that appear in the source document's metadata or cross-references, though the primary focus here is on historical features like `DOMImplementation`, `EntityReference`, and legacy activation behaviors.

# Local Summary

The text defines a collection of interfaces and attributes marked for historical context, including `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMLocator`, `DOMObject`, `DOMTokenList`, `DOMUserData`, and `entities`. It also lists legacy activation behaviors (`legacy-canceled-activation behavior`, `legacy-pre-activation behavior`) and specific historical methods like `createEntityReference()` and `getUserData()`.

# Key Claims

- The following interfaces are categorized under the **Historical** section:
  - `DOMConfiguration`
  - `DOMError`
  - `DOMErrorHandler`
  - `DOMImplementationList`
  - `DOMLocator`
  - `DOMObject`
  - `DOMTokenList`
  - `DOMUserData`
  - `entities`
  - `Entity`
  - `ENTITY_NODE`
  - `ENTITY_REFERENCE_NODE`
  - `EntityReference`
- Legacy behaviors include:
  - `legacy-canceled-activation behavior` (§ 2.7)
  - `legacy-obtain service worker fetch event listener callbacks` (§ 2.8)
  - `legacy-pre-activation behavior` (§ 2.7)
- Historical methods/functions listed:
  - `createEntityReference()` (§ 11)
  - `getUserData()` (§ 11)
  - `getFeature()` for `DOMImplementation` and `Node` (§ 11)

# Entities And Concepts

### Interfaces (Historical/Deprecated Context)
- `DOMConfiguration`: Configuration settings for the DOM implementation.
- `DOMError`: Represents a generic error in the DOM.
- `DOMErrorHandler`: Interface for handling errors during document parsing.
- `DOMImplementationList`: A collection of DOM implementations.
- `DOMLocator`: Location information within the DOM.
- `DOMObject`: Generic object interface for the DOM.
- `DOMTokenList`: Token list for attributes (e.g., class, rel).
- `DOMUserData`: User-defined data associated with nodes.
- `Entity`: Represents an entity in the document.
- `EntityReference`: Reference to an external entity.

### Attributes & Constants
- `entities`: Collection of entities defined in the document type definition.
- `ENTITY_NODE`: Constant for node type representing an entity.
- `ENTITY_REFERENCE_NODE`: Constant for node type representing an entity reference.
- `isElementContentWhitespace`: Boolean flag related to whitespace handling (legacy).
- `isSupported`: Method/attribute indicating feature support (legacy).

# Procedures And API Details

### Creation and Management
- **`createEntityReference()`**: Creates an EntityReference node. Used historically for referencing external entities defined in a DTD.
- **`getUserData()`**: Retrieves user-defined data from a node. Deprecated in favor of `dataset` or other custom attributes.
- **`getFeature(name, version)`**: (Implied via `DOMImplementation`) Gets an implementation object for a specific feature.

### Feature Support
- **`hasFeature(feature, version)`**: Checks if the DOM implementation supports a specific feature at a given version level (historical method).

### Event and Listener Contexts (Cross-referenced)
While not strictly in § 11, the chunk references legacy event handling concepts:
- `legacy-canceled-activation behavior`: Relates to Service Worker activation states.
- `legacy-pre-activation behavior`: Another state in the Service Worker lifecycle prior to modern standards.

# Nuance Or Contradictions

- **Deprecation vs. Usage**: Interfaces like `DOMTokenList` are still widely used in modern web development (e.g., for managing class lists), yet they appear here under "Historical" or alongside legacy concepts. This suggests the source document may be documenting the *evolution* of these APIs, distinguishing between their current usage and their original historical context or deprecated variants.
- **`createEntityReference()`**: This method is largely obsolete in modern HTML5/HTML parsers which no longer support external DTDs by default, making this API historically significant but practically unused in new applications.
- **Legacy Activation Behaviors**: The terms `legacy-canceled-activation behavior` and `legacy-pre-activation behavior` indicate transitional states in the Service Worker specification that have been refined or replaced in later versions of the standard.

# Candidate Wiki Hints

1. **Historical DOM Interfaces**
   - Title: `DOMImplementationList`, `DOMError`, `DOMErrorHandler`
   - Description: Overview of deprecated or legacy interfaces in the DOM Level 3 and early implementations.
   - Content Focus: Usage, deprecation status, and replacement mechanisms (e.g., using exceptions instead of `DOMError`).

2. **Entity References and DTDs**
   - Title: `EntityReference`, `createEntityReference`
   - Description: Explaining the concept of entity references in XML/HTML documents and why they are deprecated in HTML5.
   - Content Focus: Differences between internal and external entities, parsing limitations in modern browsers.

3. **Legacy Service Worker Activation**
   - Title: `legacy-pre-activation behavior`, `legacy-canceled-activation behavior`
   - Description: Historical states of Service Worker registration and activation before the current lifecycle model.
   - Content Focus: How old versions handled installation, activation, and termination compared to modern specs.

4. **DOMTokenList Utility**
   - Title: `DOMTokenList`, `classList`
   - Description: Practical usage of `DOMTokenList` for managing multiple attribute values (like `class` or `rel`).
   - Content Focus: Methods like `.add()`, `.remove()`, `.toggle()`, and its role as a bridge between DOM nodes and arrays.

## chunk-22

---
title: Chunk 22 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
- **Heading**: 11. Historical
- **Scope**: Covers a comprehensive list of DOM interfaces, attributes, methods, and concepts referenced in the specification's historical section.
- **Content Type**: An enumerated index of terms (e.g., `MutationObserver`, `ShadowRoot`, `NodeIterator`) linked to their respective sections (§) within the document.

Local Summary
This chunk serves as an extensive glossary or index for the DOM Standard, specifically focusing on terminology and identifiers associated with historical versions or legacy concepts. It lists interfaces like `XPathEvaluator` and `XSLTProcessor`, shadow DOM components (`ShadowRootInit`, `SlotAssignmentMode`), mutation tracking APIs (`MutationRecord`, `MutationObserver`), node traversal utilities (`TreeWalker`, `NodeIterator`), and various attribute/method combinations (e.g., `setAttributeNS`, `querySelectorAll`). The entries reference specific sections (§) where definitions or algorithms for these terms are detailed.

Key Claims
- The DOM standard includes support for legacy processing instructions such as XSLT (`XSLTProcessor`) and XPath evaluation (`XPathEvaluator`).
- Shadow DOM functionality is represented by interfaces like `ShadowRoot`, `SlotAssignmentMode`, and attributes like `serializable` and `slot`.
- Mutation tracking is handled via the `MutationObserver` API, generating `MutationRecord` objects that track changes to the DOM tree.
- Node traversal is facilitated by `NodeIterator` and `TreeWalker` interfaces, which support filtering (`whatToShow`) and ordering modes (e.g., `SHOW_ELEMENT`, `SHOW_COMMENT`).
- Event handling mechanisms include standard listeners, abort signals (`AbortController`, `AbortSignal`), and propagation control methods like `stopPropagation()`.

Entities And Concepts
- **Interfaces**: `ShadowRoot`, `MutationObserver`, `NodeIterator`, `TreeWalker`, `XPathEvaluator`, `XSLTProcessor`, `Element`, `Attr`, `DocumentType`.
- **Attributes/Properties**: `namespaceURI`, `localName`, `textContent`, `ownerDocument`, `shadowRoot`, `slot`, `parentElement`, `nextSibling`.
- **Methods**: `appendChild`, `removeChild`, `replaceChild`, `querySelector`, `setAttribute`, `observe`, `disconnect`, `takeRecords`.
- **Constants/Types**: `ELEMENT_NODE`, `TEXT_NODE`, `COMMENT_NODE`, `DOCUMENT_FRAGMENT_NODE`, `SHOW_ALL`, `ORDERED_NODE_ITERATOR_TYPE`.
- **Events**: `MutationEvent`, `MutationNameEvent`, `SlotChangeEvent` (implied via `slotchange`).

Procedures And API Details
- **Mutation Observation**: The `MutationObserver` constructor accepts a callback and an optional `MutationObserverInit` dict. It queues microtasks when changes occur, creating `MutationRecord` objects with properties like `addedNodes`, `removedNodes`, and `type`.
- **Shadow DOM Initialization**: A `ShadowRoot` is initialized with an `ShadowRootInit` dictionary containing options like `slots` (for legacy slot assignment) and `serializable`.
- **Node Traversal**: `NodeIterator` and `TreeWalker` are configured using a filter function (`whatToShow`) and a root node. They support methods like `nextNode()`, `previousNode()`, and properties like `currentNode`.
- **Attribute Manipulation**: Methods include `setAttributeNS(namespace, qualifiedName, value)`, `removeAttribute(qualifiedName)`, and `toggleAttribute(qualifiedName, force)`.
- **Range Operations**: The `Range` interface allows manipulating text via `setStart(node, offset)`, `setEnd(node, offset)`, and `selectNodeContents(node)`.

Nuance Or Contradictions
- The section is titled "Historical," implying some listed concepts (like explicit `MutationEvent` or legacy slot assignment modes) may be deprecated or superseded by newer mechanisms like `MutationObserver` or the HTML slot element API, though they remain part of the standard's history.
- Some terms appear with both "dfn" (definition) and "attribute/method" markers, indicating they are defined concepts that also possess specific properties or behaviors in the DOM tree structure.

Candidate Wiki Hints
- **Shadow DOM & Slots**: Create a page explaining `ShadowRoot` internals, specifically focusing on `slotAssignment`, `serializable` attributes, and the legacy vs. modern slot handling mechanisms.
- **Mutation Tracking**: Document the `MutationObserver` API workflow, detailing how it interacts with `MutationRecord` and the microtask queue (`queue a mutation observer microtask`).
- **Tree Traversal**: Summarize the differences between `NodeIterator` and `TreeWalker`, including their configuration via `whatToShow` constants (e.g., `SHOW_ELEMENT`) and traversal order options.
- **XPath & XSLT**: Provide an overview of the legacy `XPathEvaluator` and `XSLTProcessor` interfaces, noting their section references and historical context in the DOM spec.

## chunk-23

---
title: Chunk 23 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk defines the "Historical" section of the standard, listing terms borrowed from various specifications (HTML, ECMAScript, Web IDL, etc.) and providing normative references for those standards. It also includes an IDL Index defining core interfaces like `Event`, `CustomEvent`, `EventTarget`, `AbortController`, `MutationObserver`, `Node`, and `Document`.

## Local Summary
The chunk outlines which external specifications contribute specific terminology to the DOM Standard and lists the full Web IDL definitions for fundamental eventing, node manipulation, document creation, and abort mechanisms. It serves as a glossary of dependencies and a foundational interface definition block.

## Key Claims
- The standard defines terms by reference to other living standards and W3C recommendations (e.g., HTML, ECMAScript, URL).
- `Event` is the base interface for all events, including `CustomEvent`.
- `EventTarget` provides the core lifecycle methods: `addEventListener`, `removeEventListener`, and `dispatchEvent`.
- `AbortController` and `AbortSignal` allow cancellation of operations.
- `MutationObserver` allows observation of changes to a tree of nodes.
- `Node` is the base interface for DOM nodes, defining structural relationships (parent/child/sibling) and creation methods (`appendChild`, `insertBefore`).
- `Document` extends `Node` with specific document-level methods like `createElement`, `getElementById`, and `querySelector`.

## Entities And Concepts
- **Event System**: `Event`, `CustomEvent`, `EventInit`, `EventTarget`, `EventListener`, `AbortController`, `AbortSignal`.
- **Tree Traversal & Mutation**: `MutationObserver`, `MutationRecord`, `MutationObserverInit`, `ParentNode` (mixin), `ChildNode` (mixin).
- **Core Node Interface**: `Node`, `Element`, `Document`, `DocumentFragment`, `CharacterData`, `Text`, `Comment`.
- **Abort Mechanism**: `signal`, `aborted`, `reason`, `abort()`.
- **DOM Construction**: `createElement`, `createTextNode`, `createDocumentFragment`, `importNode`, `adoptNode`.
- **Querying**: `querySelector`, `querySelectorAll`, `getElementById`, `getElementsByTagName`.
- **External Specs Referenced**: HTML, ECMAScript, URL, Web IDL, Infra, Encoding, Console, DeviceOrientation, etc.

## Procedures And API Details
- **Creating Events**: `new Event(type, options)`, `new CustomEvent(type, options)`.
- **Handling Events**: `target.addEventListener(type, callback, options)`, `target.removeEventListener(type, callback, options)`, `target.dispatchEvent(event)`.
- **Observing Changes**: `observer.observe(target, config)`, `observer.disconnect()`, `observer.takeRecords()`.
- **Node Manipulation**: `node.appendChild(newNode)`, `node.insertBefore(newNode, refNode)`, `node.replaceWith(node)`, `node.remove()`.
- **Document Creation**: `doc.createElement(name)`, `doc.createTextNode(data)`, `doc.createComment(data)`, `doc.createRange()`.
- **Abort Signals**: `AbortController.abort([reason])`, `AbortSignal.timeout(ms)`, `AbortSignal.any(signals)`.

## Nuance Or Contradictions
- The standard distinguishes between legacy properties (e.g., `srcElement`, `cancelBubble`) and modern ones, often marking the former as deprecated or aliases.
- Some methods like `createEvent` are marked as legacy.
- Certain attributes have multiple names (e.g., `charset` vs `characterSet`).

## Candidate Wiki Hints
- **Event System Architecture**: A page explaining the `Event`/`CustomEvent` inheritance and the event loop phases (`CAPTURING_PHASE`, `AT_TARGET`, `BUBBLING_PHASE`).
- **MutationObserver API**: A guide on observing DOM changes for performance-sensitive applications.
- **AbortController Usage**: How to use `AbortSignal` to cancel asynchronous operations (fetch, timers).
- **Node Manipulation Mixins**: Explaining how `ParentNode`, `ChildNode`, and `NonElementParentNode` extend the base `Node` interface.
- **Document Creation API**: Methods for creating elements and text nodes via the `Document` interface.

## chunk-24

---
title: Chunk 24 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 11: Historical** of the DOM Standard, defining legacy interfaces and their current status. It details interfaces such as `DocumentType`, `ShadowRoot`, `Element`, `Attr`, character data types (`Text`, `Comment`), range manipulation (`Range`, `NodeIterator`), XPath processing, XSLT transformation, and the modernization of the Abort mechanism via `AbortController`. The text includes specific engine support tables (e.g., Firefox 57+, Chrome 66+) for these features.

# Local Summary

This section catalogs historical DOM interfaces that have been retained or superseded in modern standards. It defines the structure for `DocumentType`, `ShadowRoot` (with modes like "open"/"closed"), and the extensive attribute manipulation API on `Element`. The chunk also covers tree traversal tools (`TreeWalker`, `NodeIterator`) and legacy XPath/XSLT support. A significant portion is dedicated to the `AbortController` and `AbortSignal` interfaces, listing their specific properties (like `aborted`, `reason`) and providing a breakdown of browser compatibility for various Abort-related methods.

# Key Claims

- **Historical Interfaces**: The standard includes definitions for `DocumentType`, `ShadowRoot`, `Element`, `Attr`, `CharacterData`, `Text`, `CDATASection`, `ProcessingInstruction`, `Comment`, `AbstractRange`, `StaticRange`, `Range`, `NodeIterator`, `TreeWalker`, `XPathResult`, `XPathExpression`, `XSLTProcessor`.
- **Abort Mechanism**: The `AbortController` interface and its associated `signal` are supported in all current engines (Firefox 57+, Safari 12.1+, Chrome 66+).
- **Signal Properties**: Specific properties on `AbortSignal` such as `aborted`, `reason`, and methods like `throwIfAborted` have varying support dates (e.g., `reason` requires Firefox 97+).
- **Legacy Support**: Some legacy aliases exist, such as `webkitMatchesSelector` for `.matches()` and specific XPath result types.

# Entities And Concepts

- **Interfaces**: `DocumentType`, `ShadowRoot`, `Element`, `Attr`, `CharacterData`, `Text`, `Range`, `NodeIterator`, `TreeWalker`, `XPathEvaluator`, `XSLTProcessor`, `AbortController`, `AbortSignal`.
- **Enums**: `ShadowRootMode` ("open", "closed"), `SlotAssignmentMode` ("manual", "named").
- **Constants**: XPath result types (e.g., `ANY_TYPE`, `SNAPSHOT_TYPE`), Node filter flags (e.g., `SHOW_ELEMENT`, `SHOW_TEXT`).
- **Abort Features**: `signal`, `abort_event`, `aborted`, `reason`, `throwIfAborted`.

# Procedures And API Details

**Element Interface Methods**:
- `closest(selectors)`: Finds the closest ancestor matching a selector.
- `matches(selectors)`: Checks if the element matches a CSS selector (legacy alias: `webkitMatchesSelector`).
- `getElementsByTagName(qualifiedName)`: Returns an HTMLCollection of elements by tag name.
- `insertAdjacentElement(where, element)`: Inserts an element adjacent to the current one (marked legacy).

**Range Interface Methods**:
- `setStart(node, offset)`, `setEnd(node, offset)`: Sets range boundaries.
- `selectNode(node)`, `selectNodeContents(node)`: Selects a node or its contents within the range.
- `deleteContents()`, `extractContents()`, `cloneContents()`: Manipulates the range's contents (creates new objects).

**XPath Support**:
- `XPathEvaluator.createExpression(expression, resolver)`: Creates an expression object.
- `XPathEvaluator.evaluate(expression, contextNode, ...)`: Evaluates an XPath expression returning an `XPathResult`.

**Abort Signal Properties**:
- `signal.aborted`: Boolean indicating if the signal has been aborted.
- `signal.reason`: The reason for aborting (requires newer browser versions).
- `signal.throwIfAborted()`: Throws a DOMException if aborted.

# Nuance Or Contradictions

- **Legacy vs Modern**: Many methods like `insertAdjacentElement` are marked as legacy, while their modern counterparts or aliases (`closest`, `matches`) are preferred.
- **Engine Variance**: While the text states "In all current engines" for `AbortController`, specific properties like `signal.reason` have strict version requirements (e.g., Firefox 97+, Chrome 98+), indicating partial implementation history.
- **Useless Attributes**: Comments in the source mark attributes like `hasFeature()` and `Attr.specified` as "useless; always returns true", implying they are deprecated or obsolete for practical use.

# Candidate Wiki Hints

- **Page: DOM Element API** – Summarize attribute handling (`id`, `className`, `classList`) and query methods (`querySelector` logic implied via `closest/matches`).
- **Page: Range Manipulation** – Explain `Range` operations, including `deleteContents`, `extractContents`, and boundary setting.
- **Page: AbortController Guide** – Detail the lifecycle of `AbortController`, signal properties (`aborted`, `reason`), and usage in async contexts.
- **Page: Shadow DOM Basics** – Cover `ShadowRoot` attributes (`mode`, `delegatesFocus`) and initialization via `attachShadow`.

## chunk-25

---
title: Chunk 25 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

Chunk Context
This chunk covers "11. Historical" compatibility data for various DOM APIs and methods across major browser engines (Gecko, WebKit/Blink) and their versions. It lists support status for Range properties, Attr attributes, CharacterData methods, and Element manipulation methods.

Local Summary
The document details the historical version support for specific DOM interfaces and methods. It distinguishes between abstract base classes (e.g., `AbstractRange`) and concrete implementations (e.g., `StaticRange`, `Range`). For each API, it lists versions for Firefox, Safari, Chrome, Opera, Edge, and legacy Edge/IE, often noting "IENone" or specific version numbers like "Firefox69+" for newer features.

Key Claims
- Many modern APIs (e.g., `AbstractRange` methods) are supported in all current engines but had delayed support in Firefox (v69), Safari (v10.1/14.1), and Chrome (v90).
- Legacy implementations (`Range`, `Attr`) have been supported since the earliest versions of these browsers (e.g., Firefox 1, IE 5.5+).
- Certain methods like `CharacterData/after` and `Element/replaceWith` show a significant gap in support between older and newer engines (e.g., Firefox 49 vs. Firefox 69 for similar abstract features).
- Legacy Edge (pre-Chromium) had limited or no support for many modern APIs compared to Chromium-based Edge.

Entities And Concepts
- **DOM Interfaces**: `AbstractRange`, `StaticRange`, `Range`, `Attr`, `CDATASection`, `CharacterData`, `Element`, `DocumentType`.
- **Methods**: `endContainer`, `endOffset`, `startContainer`, `startOffset`, `localName`, `namespaceURI`, `ownerElement`, `prefix`, `value`, `after`, `before`, `appendData`, `deleteData`, `insertData`, `length`, `nextElementSibling`, `previousElementSibling`, `remove`, `replaceData`, `replaceWith`.
- **Browser Versions**: Firefox (desktop and Android), Safari (iOS and desktop), Chrome (desktop and Android), Opera, Edge (legacy and Chromium-based), IE.

Procedures And API Details
- **Range Properties**:
  - `endContainer` / `startContainer`: Supported in all current engines from Firefox 69+, Safari 14.1+/10.1+, Chrome 90+/60+. Legacy Edge supports it from v79+ or v18 depending on the property variant.
  - `endOffset` / `startOffset`: Similar version splits, with older concrete `Range` interface supported since Firefox 1.
- **Attr Attributes**:
  - Properties like `localName`, `name`, `namespaceURI`, `ownerElement`, `prefix`, `value` are supported in all current engines from Firefox 1+, Safari 1+, Chrome 1+. Legacy Edge support varies (v12+ or v6+).
- **CharacterData Methods**:
  - `appendData`, `deleteData`, `insertData`, `replaceData`: Supported since early versions (Firefox 1+).
  - `after`, `before`, `nextElementSibling`, `previousElementSibling`, `remove`, `replaceWith`: Show later adoption in Firefox (v49/v25) and Chrome (v54/v29) compared to Safari (v10/v9).
- **Element Methods**:
  - `after`, `before`, `nextElementSibling`, `previousElementSibling`, `remove`, `replaceWith`: Follow similar version trends as CharacterData methods, with older implementations supported since Firefox 3.5+.

Nuance Or Contradictions
- The term "In all current engines" seems to imply broad modern support, yet specific version numbers (e.g., Firefox 69) suggest that these features are not universally available in the entire history of those browsers.
- There is a distinction between `AbstractRange` and `StaticRange`/`Range`; the abstract versions often have later release dates for specific properties than the concrete ones.
- Legacy Edge entries often show "None" or very early versions (e.g., IE 5.5, IE 6) for older APIs, contrasting with modern Edge support which aligns more closely with Chromium-based standards.

Candidate Wiki Hints
- **DOM Compatibility Tables**: This data is ideal for creating a compatibility table on the wiki page for specific DOM methods like `Range.startContainer` or `Attr.ownerElement`.
- **Browser Version Tracking**: Useful for documenting the timeline of feature adoption in Firefox, Chrome, and Safari.
- **Legacy Browser Support**: Highlights when features dropped support for IE or Legacy Edge, useful for migration guides.

## chunk-26

---
title: Chunk 26 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 11. Historical**, specifically detailing browser compatibility tables for various DOM API methods and properties within the `Document`, `DOMImplementation`, `DOMTokenList`, `Element`, `DocumentFragment`, `CharacterData`, `Comment`, `CustomEvent`, and related interfaces. The data presents version numbers for browsers including Firefox, Safari, Chrome, Opera, Edge (Legacy), Edge, and various mobile variants (Android WebView, Samsung Internet, etc.).

# Local Summary

The section provides a comprehensive compatibility matrix for DOM APIs. It lists specific methods such as `DOMImplementation.createDocument`, `Element.appendChild` (implied via similar patterns in other sections not fully shown but referenced in logic), `DOMTokenList.add`, and properties like `Document.characterSet`. The tables indicate the minimum browser version required to support these features, noting "In all current engines" for widely supported methods and providing specific legacy versions (e.g., IE6+, Edge Legacy) for older implementations. Some entries mark certain platforms (like iOS Safari or Android WebView) with question marks or empty checks where data is unavailable.

# Key Claims

-   **Universal Support**: Many core DOM APIs (e.g., `DOMImplementation`, `CharacterData`, `Comment` in specific contexts, `Element.appendChild`) are marked as supported "In all current engines" across Firefox, Safari, Chrome, Opera, and Edge.
-   **Legacy Browser Support**: Specific versions for legacy browsers are listed, such as `Edge (Legacy) 12+` or `IE6+`, indicating when these APIs became available in older environments.
-   **Mobile Browser Variance**: Mobile browsers often have different version numbers than their desktop counterparts (e.g., Firefox for Android vs. standard Firefox). Some mobile platforms like iOS Safari or specific Android WebViews are sometimes marked with question marks, suggesting uncertain or unverified support data in the source text.
-   **Version Specificity**: Support is quantified by browser version (e.g., `Firefox 49+`, `Chrome 54+`), highlighting when a feature was introduced or became stable enough for general use.

# Entities And Concepts

-   **DOM APIs**: Methods and properties belonging to the Document Object Model (e.g., `createDocument`, `appendChild`, `removeChild`).
-   **Browser Engines**: The underlying rendering engines of web browsers (Gecko, WebKit, Blink).
-   **Legacy Browsers**: Older versions of browsers like Internet Explorer (IE) and early Edge releases.
-   **Mobile Browsers**: Browser variants designed for mobile devices (Android WebView, iOS Safari, Samsung Internet).
-   **Compatibility Matrix**: A structured table showing feature support across different browser versions.

# Procedures And API Details

-   **`DOMImplementation.createDocument`**: Supported in all current engines; legacy Edge 12+, IE9+.
-   **`DOMImplementation.createDocumentType`**: Supported in all current engines; legacy Edge 12+, IE9+.
-   **`DOMTokenList.add` / `remove` / `toggle`**: Added around Firefox 3.6, Safari 5.1, Chrome 8.
-   **`DOMTokenList.replace`**: Added later, specifically Firefox 49+, Safari 10.1, Chrome 61+.
-   **`Document.createAttributeNS`**: Supported in all current engines; legacy Edge 12+, IE9+.
-   **`Document.createElement`**: Supported in all current engines; legacy Edge 12+, IE5+.

# Nuance Or Contradictions

-   **Missing Data Indicators**: Question marks (e.g., `Firefox for Android?`, `iOS Safari?`) indicate that the source text does not have definitive compatibility data for these specific platforms, contrasting with the "In all current engines" claim for others.
-   **Legacy vs. Modern Edge**: Distinction is made between "Edge (Legacy)" and modern "Edge", sometimes showing different version requirements or noting "IENone" (Internet Explorer None) for features not supported in IE at all.
-   **Version Gaps**: Some APIs appear in later versions of major browsers (e.g., `DOMTokenList.supports` in Firefox 49+) compared to others that support them from earlier versions, reflecting incremental feature adoption.

# Candidate Wiki Hints

1.  **DOM Compatibility Tables**: A dedicated page or section explaining how to read MDN compatibility tables and interpreting version numbers.
2.  **Legacy Browser Support**: Documentation on writing for legacy browsers like Internet Explorer and early Edge, focusing on fallback strategies for APIs listed here.
3.  **DOMTokenList API Reference**: A detailed guide covering `add`, `remove`, `toggle`, and other methods with historical adoption timelines.

## chunk-27

---
title: Chunk 27 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

**Heading:** 11. Historical
**Source Path:** `raw/web/corpus-2026-05-18/060-dom-standard.md`
**Line Range:** 10726–11432
**Content Scope:** Compatibility tables for DOM Level 2 and Level 3 interfaces, specifically focusing on the `Document`, `Element`, `DocumentFragment`, and `XPathEvaluator` objects. The chunk lists methods such as `createEvent`, `querySelector`, `getElementsByClassName`, and properties like `documentElement`. It details support across desktop browsers (Firefox, Safari, Chrome, Opera, Edge) and legacy versions (IE), alongside mobile implementations (Android WebView, Samsung Internet).

# Local Summary

This section of the historical DOM standard documentation provides a comprehensive compatibility matrix for various Document and Element interface methods. It enumerates specific browser engine versions required to support features like `querySelector`, `prepend`, `replaceChildren`, and XPath expression creation (`createExpression`). The data distinguishes between modern engines (Firefox, Safari, Chrome) and legacy environments (Edge Legacy, Internet Explorer), often noting missing support in mobile browsers or older Android WebViews for newer APIs.

# Key Claims

- **Universal Support:** Core DOM Level 2 methods like `getElementsByClassName`, `createElement`, and `documentElement` are supported in "all current engines" starting from version 1 of the respective Firefox, Safari, and Chrome builds listed.
- **Modern API Adoption:** Features such as `prepend`, `replaceChildren`, and `firstElementChild` require significantly higher version numbers (e.g., Firefox 49+, Chrome 54+, Safari 10+).
- **Legacy Limitations:** Internet Explorer lacks support for many modern DOM methods, often marked as "None" or requiring very early versions (IE 5–6) for basic properties. Edge Legacy supports some features but not others like `replaceChildren`.
- **Mobile Disparities:** Mobile browsers (Android WebView, iOS Safari) frequently show question marks (?) indicating unknown status or gaps in support for newer APIs like `prepend` and `replaceChildren` at the time of documentation.

# Entities And Concepts

- **Document Interface:** The root object representing an HTML document in the DOM tree.
- **Element Interface:** Represents an element within a document.
- **DocumentFragment:** A lightweight container used to optimize DOM manipulation by creating a temporary node structure.
- **XPathEvaluator:** An interface for evaluating XPath expressions (e.g., `createExpression`, `evaluate`).
- **Browser Engines:** Firefox, Safari, Chrome, Opera, Edge (Modern and Legacy), Internet Explorer.
- **DOM Methods:** `querySelector`, `querySelectorAll`, `getElementsByClassName`, `getElementById`.

# Procedures And API Details

The following methods and properties are detailed with specific version requirements:

| Method/Property | Interface | Firefox | Safari | Chrome | Opera | Edge (Mod) | Edge (Leg) | IE |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| `createEvent` | Document | 1+ | 1+ | 1+ | 7+ | 79+ | 12+ | 5+ |
| `createExpression` | Document/XPathEval | 1+ | 3+ | 1+ | 12.1+ | 79+ | 12+ | None |
| `createNodeIterator` | Document | 1+ | 3+ | 1+ | 9+ | 79+ | 12+ | 5+ |
| `createNSResolver` | Document/XPathEval | 1+ | 3+ | 1+ | 12.1+ | 79+ | 12+ | None |
| `createProcessingInstruction` | Document | 1+ | 1+ | 1+ | 12.1+ | 79+ | 12+ | 9+ |
| `createRange` | Document | 1+ | 1+ | 1+ | 12.1+ | 79+ | 12+ | 9+ |
| `createTextNode` | Document | 1+ | 1+ | 1+ | 7+ | 79+ | 12+ | 5+ |
| `createTreeWalker` | Document | 1+ | 3+ | 1+ | 9+ | 79+ | 12+ | 9+ |
| `doctype` | Document | 1+ | 1+ | 1+ | 12.1+ | 79+ | 12+ | 6+ |
| `documentElement` | Document | 1+ | 1+ | 1+ | 7+ | 79+ | 12+ | 5+ |
| `documentURI` | Document | 1+ | 3+ | 1+ | 12.1+ | 79+ | 12+ | None |
| `evaluate` | Document/XPathEval | 1+ | 3+ | 1+ | 9+ | 79+ | 12+ | None |
| `firstElementChild` | Element/DocFragment | 25+ | 9+ | 29+ | ? | 79+ | 17+ | None |
| `getElementsByClassName` | Document | 3+ | 3.1+ | 1+ | 9.5+ | 79+ | 12+ | 9+ |
| `querySelector` | Element/DocFragment | 3.5+ | 3.1+ | 1+ | 10+ | 79+ | 12+ | 9+ |
| `replaceChildren` | Element/DocFragment | 78+ | 14+ | 86+ | ? | 86+ | ? | None |

# Nuance Or Contradictions

- **Version Gaps:** There are significant gaps between the support of legacy properties (e.g., `getElementById` in Firefox 1+) and modern manipulation methods (e.g., `replaceChildren` in Firefox 78+).
- **Mobile Uncertainty:** Many mobile browsers are marked with "?" for newer APIs, suggesting inconsistent implementation or lack of data compared to desktop counterparts.
- **XPath Support:** XPath expression creation (`createExpression`) is widely supported but lacks support in Internet Explorer entirely, contrasting with basic DOM properties which have minimal IE support.

# Candidate Wiki Hints

- **DOM Compatibility Matrix:** A dedicated page summarizing browser support for DOM Level 2 and 3 methods.
- **Modern DOM Methods:** Documentation focusing on newer APIs like `prepend`, `replaceChildren`, and `firstElementChild`.
- **XPath in the Browser:** A guide covering `XPathEvaluator` interfaces (`evaluate`, `createExpression`).

## chunk-28

---
title: Chunk 28 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Chunk Context

This chunk (lines 11434–12166) covers the **Historical** section of the DOM Standard compatibility tables. It details browser support for various DOM interfaces and methods, specifically focusing on `DocumentType`, `Element` attributes and methods, and the `Event` interface. The data indicates support status across major browsers (Firefox, Safari, Chrome, Opera, Edge, IE) and their mobile counterparts, often citing specific version numbers where support was introduced or noting lack thereof with question marks or "None".

# Local Summary

The text presents a series of compatibility tables for DOM components. It begins with `DocumentType` public/system IDs, followed by a comprehensive list of `Element` properties (like `id`, `className`, `classList`) and methods (such as `getAttribute`, `setAttribute`, `insertAdjacentText`). The section concludes with the start of the `Event` interface compatibility data. A recurring pattern shows that core Element properties are supported in very early browser versions (Firefox 1+, Safari 1+, Chrome 1+), while newer features like `assignedSlot` or `toggleAttribute` require significantly higher version numbers (e.g., Firefox 63, Chrome 69).

# Key Claims

- **DocumentType Support**: The `publicId` and `systemId` attributes of the `DocumentType` interface are supported in all current engines starting from Firefox 1, Safari 3, and Chrome 1.
- **Element Core Attributes**: Fundamental properties like `id`, `className`, `tagName`, `localName`, and `namespaceURI` have been supported since the earliest versions of major browsers (Firefox 1, Safari 1, Chrome 1).
- **Element Methods**: Methods such as `getAttribute`, `setAttribute`, `removeAttribute`, and `getElementsByTagName` are universally supported in early browser iterations. Conversely, newer methods like `getElementsByClassName` require slightly newer versions (e.g., Firefox 3, Chrome 1).
- **Shadow DOM & Slots**: The `assignedSlot` attribute and `attachShadow` method appear only in modern browsers (Firefox 63+, Safari 10+, Chrome 53+), indicating they are part of the Shadow DOM specification history.
- **Event Interface**: The base `Event` interface requires more recent versions compared to DOM elements (e.g., Firefox 11, Safari 6, Chrome 15).
- **Browser Variance**: There is significant variance in support for mobile webviews and legacy browsers like IE and Edge Legacy, often marked with question marks or "None" for specific features.

# Entities And Concepts

- **DOM Interfaces**: `DocumentType`, `Element`, `Event`.
- **Properties**: `publicId`, `systemId`, `id`, `className`, `classList`, `tagName`, `localName`, `namespaceURI`, `prefix`, `shadowRoot`, `slot`, `assignedSlot`.
- **Methods**: `getAttribute`, `setAttribute`, `removeAttribute`, `hasAttribute`, `getElementsByClassName`, `insertAdjacentElement`, `insertAdjacentText`, `toggleAttribute`.
- **Browser Engines**: Firefox, Safari, Chrome, Opera, Edge (Legacy and Current), IE.
- **Mobile Environments**: Android WebView, iOS Safari, Samsung Internet, Opera Mobile.

# Procedures And API Details

No procedural steps are described in this chunk; it is strictly a compatibility matrix listing feature availability against specific browser versions.

# Nuance Or Contradictions

The data contains several instances of uncertainty denoted by question marks (e.g., `Firefox for Android?`, `Opera?`), suggesting missing or unverified data points for those platforms. Additionally, some entries show "None" or "?IE9+" for Legacy Edge and IE, indicating a lack of support for specific modern DOM features in older Microsoft browsers. The distinction between `className` (string) and `classList` (collection) is implied by their separate entries, with `classList` requiring newer browser versions than the basic string property.

# Candidate Wiki Hints

- **Topic: Browser Compatibility Matrix** – A page summarizing how different DOM features are supported across browsers.
- **Topic: Shadow DOM History** – Focusing on the introduction of `attachShadow`, `shadowRoot`, and `assignedSlot`.
- **Topic: Element Attribute Evolution** – Comparing legacy attributes (`className`) with modern collections (`classList`).

## chunk-29

---
title: Chunk 29 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk covers **Section 11. Historical**, detailing the browser and engine compatibility support for various DOM interfaces, attributes, and methods found in `raw/web/corpus-2026-05-18/060-dom-standard.md`. The data spans from line 12168 to 12902.

## Local Summary
The document lists compatibility matrices for a wide range of DOM APIs, including Event properties (e.g., `cancelable`, `composedPath`), EventTarget methods (`addEventListener`, `dispatchEvent`), HTMLCollection attributes (`item`, `length`), and MutationObserver interfaces. Each entry specifies the minimum version required for support across major browsers (Firefox, Safari, Chrome, Edge, Opera) and environments (Android WebView, Node.js).

## Key Claims
- **Universal Support**: Many core Event properties like `Event/type` and `Event/target` are supported in all current engines starting from version 1.5 or 1 for Firefox/Safari/Chrome respectively.
- **MutationObserver Evolution**: The `MutationObserver` interface is supported in Firefox 14, Safari 7, and Chrome 26, with specific methods like `takeRecords` requiring later versions (e.g., Chrome 20).
- **Node.js Support**: The DOM standard APIs generally require Node.js version 14.5.0 or higher for full compatibility.
- **Historical Context**: Older browsers like Edge Legacy and IE9/IE8 are listed with specific minimum versions where applicable, noting "None" for unsupported features in those environments.

## Entities And Concepts
- **Event Properties**: `cancelable`, `composed`, `composedPath`, `currentTarget`, `defaultPrevented`, `eventPhase`, `isTrusted`, `preventDefault`, `stopImmediatePropagation`, `stopPropagation`, `target`, `timeStamp`, `type`.
- **EventTarget Methods**: `addEventListener`, `dispatchEvent`, `removeEventListener`.
- **HTMLCollection Attributes**: `item`, `length`, `namedItem`.
- **MutationObserver Interfaces**: `MutationObserver`, `disconnect`, `observe`, `takeRecords`.
- **MutationRecord Attributes**: `addedNodes`, `attributeName`, `attributeNamespace`, `nextSibling`, `oldValue`, `previousSibling`, `removedNodes`, `target`, `type`.
- **Browser Engines**: Firefox, Safari, Chrome, Edge (Legacy and Chromium), Opera, Android WebView, Samsung Internet.

## Procedures And API Details
- **Event Property Compatibility**:
  - `Event/cancelable`: Supported in Firefox 1.5+, Safari 1+, Chrome 1+.
  - `Event/composedPath`: Supported in Firefox 59+, Safari 10+, Chrome 53+.
  - `Event/isTrusted`: Supported in Firefox 1.5+, Safari 10+, Chrome 46+.
- **EventTarget Methods**:
  - `dispatchEvent`: Requires Firefox 2, Safari 3.1, Chrome 4.
  - `removeEventListener`: Supported in all current engines starting from version 1 for Firefox/Safari/Chrome.
- **MutationObserver Methods**:
  - `observe`: Available since Firefox 14, Safari 6, Chrome 18.
  - `takeRecords`: Requires Firefox 14, Safari 6, Chrome 20.
- **HTMLCollection Attributes**:
  - `item`, `length`, `namedItem`: All supported in all current engines starting from version 1 for Firefox/Safari/Chrome.

## Nuance Or Contradictions
- **Android WebView Discrepancies**: Some entries show "Android WebView" with a question mark or specific versions (e.g., 37+, 46+) that differ between properties, indicating potential fragmentation in Android support compared to desktop counterparts.
- **Edge Legacy vs. Edge Chromium**: The data distinguishes between Edge Legacy (based on IE) and the modern Edge Chromium, with some features missing entirely in the Legacy version ("IENone").
- **Version Gaps**: Certain features like `Event/stopImmediatePropagation` have a gap where Firefox requires version 10 while Chrome only needs version 5, highlighting uneven adoption rates.

## Candidate Wiki Hints
- **DOM Event Properties**: A page summarizing the evolution and browser support of standard Event properties (`cancelable`, `composedPath`, etc.).
- **MutationObserver API**: Documentation on observing DOM changes, detailing method availability across browsers.
- **HTMLCollection Interface**: Notes on accessing list items via `item` and `namedItem` in legacy vs. modern browsers.

## chunk-30

---
title: Chunk 30 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
This chunk covers the "Historical" section of a DOM API compatibility table, specifically detailing support matrices for `NamedNodeMap` methods and various `Node` properties and methods (including attributes like `baseURI`, `appendChild`, `childNodes`, `cloneNode`, etc.) across major browsers (Firefox, Safari, Chrome, Opera, Edge) and their legacy versions. It also includes data for `NodeIterator` properties.

## Local Summary
The document provides a granular breakdown of feature support for DOM Level 2/3 APIs within the `NamedNodeMap` interface and the `Node` interface. Each entry lists specific API names (e.g., `getNamedItemNS`, `insertBefore`) alongside browser engine versions required for compatibility, distinguishing between desktop browsers, mobile webviews, and legacy Internet Explorer versions.

## Key Claims
- **Universal Support**: Many core `Node` properties (like `nodeValue`, `nodeType`, `firstChild`) and methods (like `appendChild`, `removeChild`) are supported in "all current engines" starting from very early versions (Firefox 1+, Safari 1+, Chrome 1+).
- **Legacy IE Variance**: Support for specific `NamedNodeMap` methods often diverges in Internet Explorer, with some requiring IE6 and others requiring IE9.
- **Mobile WebView Gaps**: Android WebViews and certain mobile browsers sometimes lack support or have unknown (`?`) status compared to desktop counterparts.
- **Modern DOM Additions**: Features like `getRootNode` are noted as having significantly higher version requirements (e.g., Firefox 53+, Chrome 54+).

## Entities And Concepts
- **Interfaces**: `NamedNodeMap`, `Node`, `NodeIterator`.
- **Properties**: `baseURI`, `childNodes`, `firstChild`, `lastChild`, `nodeValue`, `parentNode`, `previousSibling`, `nextSibling`, `ownerDocument`, `parentElement`.
- **Methods**: `appendChild`, `cloneNode`, `compareDocumentPosition`, `contains`, `insertBefore`, `isConnected`, `normalize`, `removeChild`, `replaceChild`, `getNamedItemNS`, `setNamedItemNS`.
- **Iterators**: `filter`, `nextNode`, `pointerBeforeReferenceNode`, `previousNode`, `referenceNode`.

## Procedures And API Details
The chunk lists specific method signatures and property accessors without detailed syntax, focusing solely on browser compatibility matrices:
- **`NamedNodeMap` Methods**: Includes `getNamedItemNS`, `item`, `length`, `removeNamedItem`, `removeNamedItemNS`, `setNamedItem`, `setNamedItemNS`.
- **`Node` Properties**: Includes `baseURI` (Safari 4+), `compareDocumentPosition` (Chrome 2+), `getRootNode` (Firefox 53+, Chrome 54+).
- **`NodeIterator` Methods**: Includes `filter`, `nextNode`, `pointerBeforeReferenceNode`, `previousNode`.

## Nuance Or Contradictions
- **Inconsistent IE Support**: While many features are marked "IE6+" or "IE9+", some entries list "IENone" (e.g., `baseURI`, `contains`), indicating a complete lack of support in legacy Internet Explorer environments.
- **Unknown Mobile Status**: Several mobile browsers (Firefox for Android, iOS Safari, Chrome for Android) are marked with a question mark (`?`), signifying that the source data lacks definitive compatibility information for these specific platforms at the time of writing.

## Candidate Wiki Hints
- **Page: DOM API Compatibility Matrix** – A comprehensive table comparing browser support for various DOM Level 2 and 3 features.
- **Page: Node Interface** – Documentation covering standard properties and methods of the `Node` interface, including historical browser support.
- **Page: NamedNodeMap** – Details on collection interfaces for attributes, specifically focusing on namespaced item handling (`getNamedItemNS`).

## chunk-31

---
title: Chunk 31 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---
Chunk Context
The provided text covers "11. Historical," listing browser compatibility data for DOM interfaces and methods (e.g., NodeIterator, NodeList, Range, ShadowRoot). The data includes support status ("✔MDN" or "?") and version numbers for Firefox, Safari, Chrome, Opera, Edge, and legacy IE/Android/iOS WebViews.

Local Summary
This section documents the historical browser support for various DOM API properties and methods. It details which browsers implemented specific features like `NodeList.item`, `Range.extractContents`, and `ShadowRoot.mode`, often distinguishing between modern engines and legacy versions (Edge Legacy, IE). The data frequently indicates full MDN support with a checkmark or lack thereof with a question mark for mobile WebViews.

Key Claims
- Most DOM interfaces listed (e.g., `NodeList`, `Range`, `Text`) are supported in all current engines as of the time of the source's compilation, often dating back to very early versions (Firefox 1, Safari 1, Chrome 1).
- Some specific methods like `ShadowRoot/slotAssignment` have higher minimum version requirements (e.g., Firefox 92, Chrome 86) compared to the core interface itself.
- Legacy browsers (IE, Edge Legacy) show inconsistent or non-existent support ("IENone", "Edge (Legacy)?") for newer standards features like Shadow DOM properties.
- Mobile WebViews (Firefox for Android, iOS Safari, Chrome for Android) often show missing data ("?") or specific version gaps compared to desktop counterparts.

Entities And Concepts
- **DOM Interfaces**: `NodeIterator`, `NodeList`, `Range`, `ProcessingInstruction`, `ShadowRoot`, `StaticRange`, `Text`.
- **Browser Engines**: Firefox, Safari, Chrome, Opera, Edge (Legacy), IE.
- **Mobile Platforms**: Firefox for Android, iOS Safari, Chrome for Android, Samsung Internet, WebView.
- **Compatibility Markers**: "✔MDN" (supported), "?" (unknown/unlisted), version numbers indicating first support.

Procedures And API Details
- `NodeList/forEach`: Supported in Firefox 50+, Safari 10+, Chrome 51+.
- `Range/toString`: Supported in all current engines, with early versions (Firefox 1, Safari 1, Chrome 1).
- `ShadowRoot/mode`: Supported in Firefox 63+, Safari 10.1+, Chrome 53+.
- `Text/wholeText`: Supported in Firefox 3.5+, Safari 4+, Chrome 2+.
- `Range/deleteContents`: Supported in Firefox 1–15, Safari 1+, Chrome 1+ (note the range for Firefox).

Nuance Or Contradictions
- The source distinguishes between "Edge" and "Edge (Legacy)," with the latter showing no support ("IENone") for many features that modern Edge supports.
- Data for mobile WebViews is often marked with "?", suggesting incomplete tracking compared to desktop browsers where specific version numbers are provided.
- Some entries list "In all current engines" but then provide early version numbers, implying the feature existed long ago in those engines but may have been removed or changed in others (though the text mostly implies stability).

Candidate Wiki Hints
- **DOM API Compatibility**: A page summarizing browser support tables for core DOM interfaces.
- **Shadow DOM History**: Tracking the introduction and version requirements of ShadowRoot properties across browsers.
- **Range API Evolution**: Notes on `Range` method availability, particularly older methods like `deleteContents` which had specific Firefox version ranges.

## chunk-32

---
title: Chunk 32 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

Chunk Context
- **Section:** 11. Historical (DOM Compatibility Data)
- **Coverage:** Line range 14387–15051 of `060-dom-standard.md`.
- **Content Scope:** A comprehensive compatibility table listing DOM APIs, specifically focusing on the TreeWalker interface, XPath support (XPathEvaluator, XPathExpression, XPathResult), XSLT processing (XSLTProcessor), and the Element.slot property.

Local Summary
This section provides historical browser support data for various DOM Level 3 traversal and manipulation interfaces. It details version requirements for major engines (Firefox, Safari, Chrome, Opera, Edge, IE) and their mobile counterparts (Android WebView, Samsung Internet). The data confirms broad support for `TreeWalker` methods across modern browsers while noting legacy limitations in older Internet Explorer versions and specific gaps in early Edge builds regarding XPath and XSLT.

Key Claims
- **Universal Support:** The `TreeWalker` interface and its primary properties (`currentNode`, `filter`, `firstChild`, `lastChild`, `nextNode`, `nextSibling`, `parentNode`, `previousNode`, `previousSibling`, `root`, `whatToShow`) are supported in all current engines.
- **Legacy Browser Limits:** Older versions of Internet Explorer (IE5, IE9) lack support for these advanced traversal features. Legacy Edge (version 12+) shows partial or no support depending on the specific API.
- **XPath Support:** The `XPathEvaluator` and associated `XPathExpression`, `XPathResult` interfaces are present in modern engines but absent or limited in legacy Edge and IE environments.
- **XSLT Capabilities:** `XSLTProcessor` is widely supported in current versions, with `transformToDocument` and `transformToFragment` being key methods listed.
- **Shadow DOM Indicator:** The presence of the `Element.slot` property indicates support for Shadow DOM v1, available from Firefox 63, Safari 10, and Chrome 53 onwards.

Entities And Concepts
- **TreeWalker:** An interface used to traverse a tree structure (Document) using filters and specific traversal rules.
  - *Properties/Methods:* `currentNode`, `filter`, `firstChild`, `lastChild`, `nextNode`, `nextSibling`, `parentNode`, `previousNode`, `previousSibling`, `root`, `whatToShow`.
- **XPath API:** A set of interfaces for evaluating XPath expressions within the DOM.
  - *Entities:* `XPathEvaluator`, `XPathExpression`, `XPathResult` (with properties like `booleanValue`, `numberValue`, `stringValue`, `singleNodeValue`, `snapshotItem`, `snapshotLength`).
- **XSLT API:** Interfaces for applying XSLT transformations to XML documents.
  - *Entity:* `XSLTProcessor`.
  - *Methods:* `clearParameters`, `getParameter`, `importStylesheet`, `removeParameter`, `reset`, `setParameter`, `transformToDocument`, `transformToFragment`.
- **Shadow DOM:** A DOM feature allowing encapsulated sub-trees.
  - *Property:* `slot` (on Element).

Procedures And API Details
- **TreeWalker Initialization:** Developers can instantiate a `TreeWalker` by specifying a root node, a filter function (optional), and a `whatToShow` flag to define which nodes to visit during traversal.
- **XPath Evaluation:** Use `XPathEvaluator.evaluate()` with an `XPathExpression` instance to query the DOM. The result is typically wrapped in an `XPathResult` object, allowing access to results via properties like `singleNodeValue` for single-node matches or iteration methods like `iterateNext`.
- **XSLT Transformation:** Instantiate an `XSLTProcessor`, optionally load stylesheets via `importStylesheet()` or set parameters using `setParameter()`, and execute transformations using `transformToFragment()` (for DOM fragments) or `transformToDocument()` (for full documents).
- **Slot Assignment:** Assign a slot name to an element in a Shadow Root using the `slot` property (e.g., `element.slot = "name"`), enabling content distribution within shadow host elements.

Nuance Or Contradictions
- **Edge Discrepancy:** There is a noted distinction between "Edge" (Chromium-based) and "Edge (Legacy)" (IE-based). Legacy Edge often lacks XPath and XSLT support entirely, whereas modern Edge supports them.
- **Mobile Variability:** Mobile browsers like "Firefox for Android," "iOS Safari," and "Samsung Internet" show varying levels of support depending on the underlying WebView version or specific mobile engine implementation (e.g., Samsung Internet 10.1+ vs. older versions).
- **IE Limitations:** Internet Explorer is consistently marked with "?" or "None" for modern features like XPath, XSLT, and Shadow DOM, highlighting a clear divide between legacy and modern web standards support.

Candidate Wiki Hints
- **TreeWalker API Reference:** A page documenting the properties and methods of `TreeWalker`, including examples of filtering nodes.
- **XPath in JavaScript:** A guide explaining how to use `XPathEvaluator` and `XPathResult` for DOM queries.
- **XSLT Transformation Guide:** Instructions on using `XSLTProcessor` to transform XML/HTML documents.
- **Shadow DOM Slots:** An article covering the `slot` property, light DOM distribution, and compatibility across browsers.

