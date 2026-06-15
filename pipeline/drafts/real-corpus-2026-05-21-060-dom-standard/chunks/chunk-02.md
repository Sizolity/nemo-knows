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
