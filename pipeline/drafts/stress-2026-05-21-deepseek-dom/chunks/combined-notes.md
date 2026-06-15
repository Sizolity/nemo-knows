## chunk-01

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

## chunk-02

---
title: Chunk 02 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/060-dom-standard.md`, lines 1580–2900.
- Heading path: `3. Aborting ongoing activities > 3.2. Interface AbortSignal > 3.2.1. Garbage collection` through `4.3.3. Interface MutationRecord`.
- Covers: AbortSignal GC behaviour, requirements for using AbortSignal in APIs, node tree model, shadow trees (slots, slottables, assignment, signalling), mutation algorithms (insert, move, replace, remove), ParentNode/ChildNode/other mixins, NodeList/HTMLCollection, and the MutationObserver system.

## Local Summary
This chunk defines how an `AbortSignal` must not be garbage‑collected while it depends on source signals and still has listeners or abort algorithms. It then sets out the rules for any promise‑based web API to accept an `AbortSignal` via a `signal` dictionary member, reject with the signal’s abort reason, use abort steps, and handle the already‑aborted case.
The bulk of the chunk details the DOM node tree. It introduces nodes, child constraints per node type, document trees, shadow trees (with host, light tree, and `connected` definition), slots (created by `slot` elements, named, assigned nodes), slottables (Element and Text, optionally with manual assignment), and algorithms for finding, assigning, and signalling slot changes.
A suite of low‑level mutation algorithms is defined: **pre‑insert validity checks**, **pre‑insert**, **insert** (including insertion steps, post‑connection steps, custom element callbacks, and shadow‑slot re‑assignment), **move** (with separate moving steps, `moveBefore`), **replace**, **replace all**, **pre‑remove**, and **remove** (with removing steps, `disconnectedCallback`, and transient registered observers).
Several mixins are described: `NonElementParentNode` (exposes `getElementById` on `Document` and `DocumentFragment`), `DocumentOrShadowRoot` (`customElementRegistry`), `ParentNode` (`children`, `firstElementChild`, `lastElementChild`, `childElementCount`, `prepend`, `append`, `replaceChildren`, `moveBefore`, `querySelector`, `querySelectorAll`), `NonDocumentTypeChildNode` (`previousElementSibling`, `nextElementSibling`), `ChildNode` (`before`, `after`, `replaceWith`, `remove`), and `Slottable` (`assignedSlot`).
Finally, it covers old‑style live collections (`NodeList`, `HTMLCollection`) and the MutationObserver system: per‑agent flags, pending observers, the notify algorithm, slot‑change events, `MutationObserver` interface (`observe`, `disconnect`, `takeRecords`), options dictionary validation, and `MutationRecord`.

## Key Claims
- A non‑aborted dependent `AbortSignal` must not be garbage‑collected while its source signals are non‑empty **and** it has registered `abort` event listeners or non‑empty abort algorithms.
- Any promise‑returning web API that supports aborting must:
  - Accept an `AbortSignal` through a `signal` dictionary member.
  - Reject the promise with the `AbortSignal` object’s **abort reason**.
  - Reject immediately if the signal is already aborted.
  - Use the **abort algorithms** mechanism without clashing with other observers.
- DOM documents are represented as **node trees**; node interfaces respect strict child constraints (Document, DocumentFragment, Element, CharacterData, etc.).
- A **shadow tree** is always attached to a host; a node is **connected** if its shadow‑including root is a document.
- Slots are created via HTML’s `slot` element, have a `name` (updated via attribute change steps), and maintain an **assigned nodes** list.
- **Slottables** (Element and Text) have an `assigned slot` and an optional **manual slot assignment** (weak reference).
- **Insertion steps** must not modify the node tree, create browsing contexts, fire events, or execute JavaScript; they may queue tasks. **Post‑connection steps** allow those actions atomically after a batch insert.
- A **move** primitive (e.g., `moveBefore`) is distinct from remove‑+‑insert; it does **not** invoke insertion/removing steps, but runs **moving steps** and a `connectedMoveCallback` on custom elements.
- The `remove` algorithm includes transient registered observers to retain `subtree` observation after a node is removed.
- `NodeList` is a collection of nodes (live by default); `HTMLCollection` is a legacy element collection with `namedItem`.
- Mutation observers rely on per‑agent **pending mutation observers** and a **microtask**; slot‑change events are fired from the same microtask.

## Entities And Concepts
- **AbortSignal** – dependent signal GC rule
- **AbortController / AbortSignal in APIs** – `signal` dictionary member, abort reason, abort steps
- **Node tree** – constraints per interface (Document, DocumentFragment, Element, etc.)
- **Document tree**, **document element**, **in a document tree** vs **in a document**
- **Shadow tree**, **shadow root**, **host**, **light tree**, **connected**
- **Slot** – name, attribute change steps, default slot, assigned nodes
- **Slottable** – Element/Text, name, assigned slot, manual slot assignment (weak reference)
- **Find a slot** / **find slottables** / **find flattened slottables**
- **Assign slottables for a slot**, **assign slottables for a tree**, **assign a slot**
- **Signal slots** (per agent) and **slot change** signalling
- **Insertion steps**, **post‑connection steps**, **children changed steps**
- **Moving steps** (separate from insertion/removing)
- **Removing steps**, **transient registered observer**
- **Pre‑insert validity**, **pre‑insert**, **insert**, **move**, **replace**, **replace all**, **pre‑remove**, **remove**
- **Pre‑insert validity** checks for Document constraints (DocumentFragment with multiple elements, doctype after element, etc.)
- `NonElementParentNode` – exposed on Document and DocumentFragment
- `DocumentOrShadowRoot` – `customElementRegistry`
- `ParentNode` – `children`, `firstElementChild`, `lastElementChild`, `childElementCount`, `prepend`, `append`, `replaceChildren`, `moveBefore`, `querySelector`, `querySelectorAll`
- `NonDocumentTypeChildNode` – `previousElementSibling`, `nextElementSibling`
- `ChildNode` – `before`, `after`, `replaceWith`, `remove`
- `Slottable` – `assignedSlot`
- **Collection** – live vs static, filter and root
- `NodeList` – length, `item`, iterable
- `HTMLCollection` – historical, `namedItem`, supported property names based on ID and `name` attribute
- **MutationObserver** – callback, node list (weak references), record queue
- **MutationObserverInit** – `childList`, `attributes`, `characterData`, `subtree`, `attributeOldValue`, `characterDataOldValue`, `attributeFilter`
- **Registered observer**, **transient registered observer**
- **Pending mutation observers** (per agent), **mutation observer microtask queued**
- **Notify mutation observers** – empty queues, invoke callbacks, fire `slotchange` events
- `MutationRecord` – type, target, addedNodes, removedNodes, previousSibling, nextSibling, attributeName, attributeNamespace, oldValue

## Procedures And API Details
- **AbortSignal GC**: The condition ties collection to source signals and presence of listeners/algorithms.
- **Example promise method** `doAmazingness(options)`:
  1. Get `global`.
  2. Create promise `p`.
  3. If `options["signal"]` exists:
     - If signal already aborted → reject `p` with `signal`’s abort reason, return `p`.
     - Add abort steps: stop amazing things; reject `p` with `signal`’s abort reason.
  4. In parallel: compute result, queue a global task to resolve `p`.
  5. Return `p`.
- **Node tree child constraints** (rules for each interface).
- **Length of a node**: 0 for DocumentType/Attr; CharacterData’s `data` length; number of children otherwise.
- **`connected`** = shadow‑including root is a document.
- **Slot name update**: attribute change steps; run assign slottables for a tree on the root.
- **Slottable name update**: attribute change steps; if slottable is assigned, run assign slottables for assigned slot; then run assign a slot.
- **Find a slot** (slottable, open flag): uses parent’s shadow root, checks mode for `"open"`, handles `"manual"` slot assignment.
- **Find slottables for a slot**: manual assignment vs named slot matching.
- **Find flattened slottables**: recursive, includes slottables from nested slots.
- **Assign slottables for a slot**: update assigned nodes; if changed → signal a slot change; update `assigned slot` on each slottable.
- **Assign a slot (slottable)**: find slot; if found, assign slottables for it.
- **Signal a slot change**: append slot to agent’s signal slots; queue a mutation observer microtask.
- **Insert algorithm** (detailed): adopts nodes, handles live ranges, shadow hosts, slot assignment, custom element callbacks (`connectedCallback`), registry scoping, suppression of observers, queues tree mutation record, collects nodes for `post‑connection` steps.
- **Move algorithm**: checks same shadow‑including root, host‑including ancestor, child validity; live range/NodeIterator pre‑remove; remove from old parent; slot re‑assignment; run moving steps (with `isSubtreeRoot`, `oldAncestor`); if custom and connected → `connectedMoveCallback`; queue mutation records.
- **Replace** child: uses nodes, respects document constraints, removes child with suppression, inserts new nodes, queues mutation record.
- **Replace all** with node/null: removes all children, inserts if non‑null; does **not** check node tree constraints — spec authors must use wisely.
- **Remove** node: live range/iterator steps; slot re‑assignment; runs removing steps on node and shadow‑including descendants; calls `disconnectedCallback` if custom and parent connected; appends transient registered observers up the ancestor chain for `subtree` observers.
- **`prepend`, `append`, `replaceChildren`** on `ParentNode`: convert nodes into a single node, use pre‑insert/append/replace‑all.
- **`moveBefore(node, child)`**: use the move primitive.
- **`querySelector` / `querySelectorAll`**: scope‑match a selectors string against the context node.
- **`ChildNode`**: `before`, `after`, `replaceWith` use `convert nodes into a node`, handle parent, viable sibling.
- **`Slottable.assignedSlot`**: calls find a slot with `open` = true.
- **`HTMLCollection.namedItem`**: supports ID and `name` attribute in HTML namespace.
- **`MutationObserver.observe(target, options)`**: validates options, deduplicates, sets up registered observer; stores weak reference to target.
- **`MutationObserver.disconnect()`**: removes all registered observers, empties record queue.
- **`MutationObserver.takeRecords()`**: clones and empties the record queue.
- **Queue a mutation record**: walks inclusive ancestors, checks options (subtree, type filters, attributeFilter, oldValue), enqueues to interested observers, appends to pending observers, queues microtask.
- **Queue a tree mutation record**: specialization for `childList` with added/removed nodes.
- **Notify mutation observers microtask**: empties signal slots, fires `slotchange` events on each slot.

## Nuance Or Contradictions
- The note about `in a document` being deprecated because older specs have not been updated to account for shadow trees.
- `Attr` nodes “participate in a tree for historical reasons” but never have a parent or children.
- Insertion steps are deliberately restricted from running JavaScript; the example shows that `script` and `style` insertion still has script‑observable side effects through post‑connection steps.
- The `move` operation is a distinct primitive that **does not** invoke insertion or removing steps; it uses moving steps and `connectedMoveCallback` for custom elements — preserving state that would be lost by remove‑and‑insert.
- `replace all` does **not** check node tree constraints, placing the onus on specification authors.
- Transient registered observers exist solely to preserve `subtree` observation for a node after it is removed from its parent, so mutations within descendants are not lost.
- The `moveBefore` method on `ParentNode` uses the move primitive, while `before`/`after`/`replaceWith` on `ChildNode` use insert/replace (via `convert nodes into a node` and `pre‑insert`/`replace`).
- `slottable`’s `manual slot assignment` is implemented via a weak reference because it is not directly accessible from script.
- The algorithm for `assign slottables for a slot` compares the old and new assigned nodes lists; it only signals a slot change if they are not identical.

## Candidate Wiki Hints
- **AbortSignal garbage collection** – a dedicated note on the GC rule for dependent signals.
- **AbortSignal in web APIs** – the requirements and pattern for promise‑returning APIs that accept `AbortSignal`.
- **Node tree constraints** – a reference page listing the allowed children for each node interface.
- **Shadow DOM: slots and slottables** – the lifecycle of slot assignment, flattened slottables, and signalling.
- **Mutation algorithms (insert, move, remove)** – detailed breakdown of the insert, move, and remove primitives, their step restrictions, custom element callbacks, and live range handling.
- **ParentNode mixin** – the procedural `prepend`/`append`/`replaceChildren`/`moveBefore` methods and `convert nodes into a node`.
- **HTMLCollection vs NodeList** – live collections, `namedItem` and its supported property names.
- **MutationObserver lifecycle** – registration, transient observers, microtask notification, and `slotchange` event.

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/060-dom-standard.md` (DOM Standard)
- Chunk 3 of 9, lines 2901–4462
- Heading path: `Text: Why yes.⏎ > 4.4. Interface Node` through `4.8. Interface ShadowRoot`
- Covers: `Node`, `Document`, `DOMImplementation`, `DocumentType`, `DocumentFragment`, `ShadowRoot` interfaces, plus associated algorithms, constants, and concepts.

## Local Summary
This chunk defines the fundamental `Node` interface and its descendants: `Document`, `DocumentType`, `DocumentFragment`, and `ShadowRoot`. It documents the DOM tree structure, node properties (type, name, relationships), mutation methods, cloning, comparison, namespace utilities, and document creation and manipulation procedures. The `DOMImplementation` factory interface and the shadow DOM (`ShadowRoot`) infrastructure are also detailed, including shadow‑including tree traversals and retargeting for event dispatch.

## Key Claims
- `Node` is an abstract interface; every node has a node document, a registered observer list, and a `get the parent` algorithm that may route through assigned slots.
- Node type constants (`ELEMENT_NODE`, `TEXT_NODE`, etc.) and `nodeName` values are defined per interface.
- `isConnected`, `ownerDocument`, `getRootNode()`, `parentNode`, `childNodes`, and sibling accessors reflect the live tree.
- `cloneNode()` triggers the “clone a node” algorithm, which may run cloning steps and optionally attach a shadow root.
- `compareDocumentPosition()` returns a bitmask; attribute ordering relative to children is special‑cased.
- `createElement` and `createElementNS` support `ElementCreationOptions` (custom element registry, `is`); `createDocument` returns an `XMLDocument` and sets content type based on namespace.
- `importNode()` and `adoptNode()` move or copy nodes across documents; `adopt` updates node documents and custom element registries, enqueuing `adoptedCallback` for custom elements.
- `ShadowRoot` is a `DocumentFragment` with a non‑null host, mode (`"open"`/`"closed"`), slot assignment (`"manual"`/`"named"`), and flags for focus, declarative, clonable, serializable, and custom element registry.
- Shadow‑including tree order and retargeting algorithm are defined for event path computation.

## Entities And Concepts
- **Interfaces**: `Node`, `Document`, `XMLDocument`, `DOMImplementation`, `DocumentType`, `DocumentFragment`, `ShadowRoot`.
- **Dictionaries**: `GetRootNodeOptions`, `ElementCreationOptions`, `ImportNodeOptions`.
- **Enums**: `ShadowRootMode` (`"open"`, `"closed"`), `SlotAssignmentMode` (`"manual"`, `"named"`).
- **Constants**: `ELEMENT_NODE` (1) through `NOTATION_NODE` (12); `DOCUMENT_POSITION_DISCONNECTED` (0x01), `PRECEDING` (0x02), `FOLLOWING` (0x04), `CONTAINS` (0x08), `CONTAINED_BY` (0x10), `IMPLEMENTATION_SPECIFIC` (0x20).
- **Concepts**: node document, shadow‑including root, shadow‑including tree order, inclusive descendant, connected node, cloning steps, adopting steps, global vs. scoped custom element registry, flatten element creation options, retargeting, closed‑shadow‑hidden.

## Procedures And API Details
- **`nodeType` getter**: dispatches on the interface implemented (Element → `ELEMENT_NODE`, etc.).
- **`nodeName` getter**: returns HTML‑uppercased qualified name for elements, `"#text"` for Text, etc.
- **`getRootNode(options)`**: returns root if `composed` false, else shadow‑including root.
- **`normalize()`**: merges contiguous exclusive `Text` nodes, removing empty ones and adjusting live range positions.
- **Clone a node**: deep‑clones with optional shadow root attachment; preserves element custom element registry or falls back to `fallbackRegistry`.
- **`compareDocumentPosition`**: attribute handling leads to special `IMPLEMENTATION_SPECIFIC` | `PRECEDING`/`FOLLOWING` results.
- **`createElement`**: if HTML document, lowercases local name; uses HTML namespace unless content type is `"application/xhtml+xml"`; supports `customElementRegistry` and `is` options.
- **`adoptNode`**: throws for documents and shadow roots; updates node documents and custom element registries for all shadow‑including descendants, then enqueues `adoptedCallback` for custom elements.
- **`DOMImplementation`**: `createDocument` returns `XMLDocument` with content type mapped from namespace; `createHTMLDocument` builds a minimal HTML document.
- **ShadowRoot properties**: `host`, `mode`, `delegatesFocus`, `slotAssignment`, `clonable`, `serializable`; `onslotchange` event handler.
- **Shadow‑including tree order**: preorder, depth‑first, diving into shadow trees after encountering a shadow host.

## Nuance Or Contradictions
- `isSameNode` is a legacy alias of `===` and should be avoided.
- `hasFeature()` always returns `true`; it was unreliable and is retained only for legacy compatibility.
- `createEvent()` is legacy; event constructors are preferred, but the method still maps string names to interfaces (e.g., `"htmlevents"` → `Event`).
- `importNode` and `adoptNode` treat scoped custom element registries differently: global registries are replaced with the document’s effective global registry; scoped registries are preserved.
- `compareDocumentPosition` for attributes returns `IMPLEMENTATION_SPECIFIC` combined with ordering bits; the exact ordering between attributes is implementation‑specific and may use pointer comparison or `Math.random()` in JavaScript implementations.
- `createCDATASection` throws `"NotSupportedError"` in HTML documents.

## Candidate Wiki Hints
- **Node (DOM)** – central abstract interface, type constants, tree traversal, cloning, comparison.
- **Document (DOM)** – document creation, mode (`no‑quirks`, `quirks`, `limited‑quirks`), factory methods, `createEvent` legacy.
- **ShadowRoot** – shadow DOM encapsulation, retargeting, slot assignment, declarative shadow roots.
- **DOMImplementation** – document type and document factory, `hasFeature` deprecation.
- **Tree Traversal & Position** – `compareDocumentPosition` bitmask, node relationships, shadow‑including order.
- **Custom Element Registry & Adoption** – registry scoping, `adoptedCallback` lifecycle, global vs. scoped registries.
- **Cloning Nodes** – `cloneNode`, shadow root clonable, `importNode` fallback registry.

## chunk-04

---
title: Chunk 04 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
Lines 4463-5924 of the DOM Standard, covering sections 4.9 (Element interface) through 4.14 (Comment interface) and the entirety of section 5 (Ranges). Headings range from “4.9. Interface Element” and its subsections (NamedNodeMap, Attr) to “5.5. Interface Range”. The chunk details element‑related algorithms, attribute handling, shadow DOM attachment, text and character data interfaces, and the abstract and live range models.

## Local Summary
This chunk defines the `Element` interface and its supporting types (`NamedNodeMap`, `Attr`), the abstract `CharacterData` interface and its concrete sub‑interfaces (`Text`, `CDATASection`, `ProcessingInstruction`, `Comment`), and the Range family (`AbstractRange`, `StaticRange`, `Range`). Key algorithmic content includes custom element states, attribute lifecycle, `attachShadow`, adjacent insertion, text splitting, and live range adjustment during tree mutations.

## Key Claims
- Elements possess an attribute list exposed via a `NamedNodeMap`. Attribute change steps propagate mutation records, custom element callbacks, and ID updates.
- Custom element states are `"undefined"`, `"failed"`, `"uncustomized"`, `"precustomized"`, and `"custom"`. An element is **defined** if its state is `"uncustomized"` or `"custom"`, and **custom** only if the state is `"custom"`.
- Shadow roots can be attached only to elements in the HTML namespace with a valid shadow host name, and only if the custom element definition (if any) does not disable shadow.
- `CharacterData` is an abstract interface; its `data` is a mutable string. The `replace data` algorithm adjusts live ranges when the data changes.
- A `StaticRange` does not update with tree mutations; a `Range` (live range) does. Live ranges ensure that the represented content remains coherent after tree changes.
- A boundary point is a (node, offset) tuple. A range’s start and end define a sequence between them. Collapsed ranges have identical start and end.
- Live range pre‑remove steps re‑parent range endpoints that are inside a removed node to its parent and adjust offsets.

## Entities And Concepts
- **Element**: node with a namespace, prefix, local name, custom element state/definition, shadow root, attribute list, and optional ID.
- **NamedNodeMap**: ordered map of attributes; length, item, getNamedItem, setNamedItem, removeNamedItem (and NS variants).
- **Attr**: node representing a content attribute; has namespace, prefix, local name, value, owner element. The `specified` getter always returns `true`.
- **CharacterData**: abstract base for text‑like nodes; has `data`, `length`, substring/insert/delete/replace methods.
- **Text**: extends `CharacterData`; includes `splitText` and `wholeText`. `CDATASection` is a non‑exclusive `Text` subclass.
- **ProcessingInstruction**: `CharacterData` with a `target`.
- **Comment**: `CharacterData` with a constructor.
- **AbstractRange**: has start/end boundary points (node + offset) and a `collapsed` flag.
- **StaticRange**: immutable range; must be valid (start before/equal end, offsets within length, same root).
- **Range (live range)**: mutable range that tracks tree mutations; supports `setStart/End`, `selectNode`, `extractContents`, `insertNode`, `surroundContents`, etc.
- **Contained/partially contained nodes**: concepts used to describe which nodes lie between a live range’s start and end.

## Procedures And API Details
- **`setAttribute(qualifiedName, value)`**: if no existing attribute with that qualified name, treats the argument as a local name and creates a new attribute; validates only the local name.
- **`toggleAttribute(qualifiedName, force)`**: validates as a local name; toggles presence (or uses `force`) and returns whether the attribute is now present.
- **`attachShadow(init)`**: validates namespace, host name, disables‑shadow flag; sets up a new shadow root with mode, delegatesFocus, slotAssignment, clonable, serializable, and a `CustomElementRegistry`.
- **`insertAdjacentElement`/`insertAdjacentText`**: insert nodes relative to `"beforebegin"`, `"afterbegin"`, `"beforeend"`, `"afterend"`.
- **`splitText(offset)`**: splits a `Text` node at the given offset, updates live ranges, returns the new node containing the remainder.
- **`StaticRange` constructor**: validates that the boundary nodes are not `DocumentType` or `Attr`; sets start/end.
- **`Range.setStart`/`setEnd`**: after validation, if the boundary is after the range’s end (for start) or before start (for end), the opposite endpoint is adjusted.
- **`compareBoundaryPoints(how, sourceRange)`**: compares two ranges’ boundary points using constants `START_TO_START`, `START_TO_END`, `END_TO_END`, `END_TO_START`.
- **Live range containment rules**: the start node and end node are never contained; partially contained nodes exist only when start and end nodes differ.
- **Pre‑remove steps**: move endpoints inside a removed node to its parent, decrease offsets after the removed index.

## Nuance Or Contradictions
- `insertAdjacentText` returns nothing because the method predates a well‑designed return value.
- Despite the parameter name `qualifiedName`, `setAttribute` treats it as a local name when no attribute matching the full name exists, leading to the local‑name‑only validation.
- `NamedNodeMap`’s `supported property names` will omit mixed‑case duplicates for HTML elements in HTML documents by removing entries where the lowercase version differs.
- `Attr` designed today would reportedly have only `name` and `value`, without the separate `namespaceURI`, `prefix`, `localName`.
- A `StaticRange` remains valid only if its boundaries do not cross document trees and offsets remain within node lengths; no automatic adjustment occurs on mutation.

## Candidate Wiki Hints
- **Element interface** – core properties, attribute methods, classList, slot, custom element states, shadow root.
- **Attribute change steps** – mutation records, custom element callbacks, ID updates.
- **Shadow DOM attachment** – valid host names, `ShadowRootInit`, delegation, slot assignment, serializable/clonable options.
- **NamedNodeMap and Attr** – legacy attribute mapping, `specified` always true.
- **Text and CharacterData** – `splitText`, `wholeText`, data replacement and live range adjustments.
- **StaticRange vs Range** – immutability, validation, maintenance cost, use cases.
- **Live range algorithms** – containment definitions, pre‑remove steps, tree mutation effects, `commonAncestorContainer`.
- **Custom element states** – `undefined`, `failed`, `uncustomized`, `precustomized`, `custom`, defined vs custom.

## chunk-05

---
title: Chunk 05 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
Lines 5925–6887 of `raw/web/corpus-2026-05-18/060-dom-standard.md`.
Headings covered:
- 5. Switch on the position of thisPoint relative to sourcePoint (end of Range section)
- 6. Traversal (NodeIterator, TreeWalker, NodeFilter)
- 7. Sets (DOMTokenList)
- 8. XPath (XPathResult, XPathExpression, XPathEvaluatorBase, XPathEvaluator)
- 9. XSLT (XSLTProcessor)
- 10. Security and privacy considerations

## Local Summary
This chunk completes the `Range` methods (`deleteContents`, `extractContents`, `cloneContents`, `insertNode`, `surroundContents`, `cloneRange`, `detach`), specifies point/range comparison (`comparePoint`, `isPointInRange`, `intersectsNode`), and `Range` stringification. It then defines DOM traversal via `NodeIterator` and `TreeWalker` (including `NodeFilter` callbacks and `whatToShow` bitmasks). The `DOMTokenList` interface and its backing token set are described, followed by minimal preservation of XPath and XSLT APIs with notes that full definitions are missing. The chunk closes with a statement that no known security or privacy issues exist for the DOM Standard.

## Key Claims
- `deleteContents()`, `extractContents()`, and `cloneContents()` handle partially contained nodes by splitting, cloning, and removing as appropriate; all three share a common ancestor determination and partial containment logic.
- `insertNode(node)` splits a Text start node if necessary, removes `node` from any previous parent, and inserts it before a computed reference point, adjusting the range’s end if collapsed.
- `surroundContents(newParent)` throws if a non‑Text node is partially contained; it extracts the range, clears `newParent`’s children, inserts `newParent`, and appends the extracted fragment.
- `detach()` is preserved for compatibility but performs no operation.
- `comparePoint(node, offset)` returns −1, 0, or 1 based on the point’s position relative to the range; it throws for wrong document or doctype.
- `intersectsNode(node)` returns true if any part of the node is inside the range.
- Traversal objects (`NodeIterator`, `TreeWalker`) share an `is active` flag to prevent recursive filter invocation.
- `NodeIterator` has a reference node and a pointer‑before‑reference boolean; its pre‑remove steps adjust the reference when an ancestor of it is removed.
- `TreeWalker` provides tree‑navigation methods like `parentNode()`, `firstChild()`, `nextSibling()`, `previousNode()`, and `nextNode()` that respect filtering.
- `NodeFilter` defines constants for `acceptNode` return values (`FILTER_ACCEPT`, `FILTER_REJECT`, `FILTER_SKIP`) and `whatToShow` bitmask constants (e.g., `SHOW_ELEMENT`, `SHOW_TEXT`).
- `DOMTokenList` is backed by a token set; its update steps may not always run for `toggle()` and `replace()` for web compatibility.
- XPath and XSLT interfaces are retained only for Web IDL compatibility; complete definitions are not provided.
- The DOM Standard states there are no known security or privacy considerations.

## Entities And Concepts
- **Range**: methods for deletion, extraction, cloning, insertion, and point comparison; comparison helper *(before/equal/after)*.
- **NodeIterator**: root, reference, pointerBeforeReference, whatToShow, filter; `nextNode()`, `previousNode()`, pre‑remove steps.
- **TreeWalker**: root, whatToShow, filter, currentNode; parent/child/sibling/previous/next navigation.
- **NodeFilter**: callback interface with constants for filtering and whatToShow.
- **DOMTokenList**: token set, element, attribute name; validation, update, serialize steps; methods `add`, `remove`, `toggle`, `replace`, `supports`.
- **XPathResult**, **XPathExpression**, **XPathEvaluatorBase**, **XPathEvaluator**: legacy XPath 1.0 evaluation API.
- **XSLTProcessor**: legacy XSLT transformation API (`importStylesheet`, `transformToFragment`, `transformToDocument`, etc.).
- **Security**: no known concerns.

## Procedures And API Details
- **deleteContents()**: if collapsed → return; if same CharacterData node → replace data with empty string; otherwise collect `nodesToRemove` (all contained nodes omitting those whose parent is also contained), adjust start/end to a new node/offset, then remove the collected nodes and trim CharacterData boundaries.
- **extract a live range**: builds a `DocumentFragment` of the range’s contents, removes the original nodes, and returns the fragment; uses `firstPartiallyContainedChild`/`lastPartiallyContainedChild` logic.
- **clone the contents**: similar to extract but clones instead of removing; uses `cloneNode(subtree true)` for fully contained children.
- **insert node into range**: throws if start node is a `ProcessingInstruction`, `Comment`, or orphan `Text`; splits text node if start node is `Text`; ensures pre‑insert validity; removes node if already in tree; computes new offset and inserts.
- **surroundContents(newParent)**: throws if a non‑Text node is partially contained; extracts range, clears `newParent`, inserts `newParent` at range start, appends fragment to `newParent`, then selects `newParent`.
- **comparePoint(node, offset)**: returns −1 (before start), 1 (after end), 0 (in range) using internal position checks.
- **intersectsNode(node)**: returns true if node’s parent (or the node itself if root) lies such that (parent, offset) < end and (parent, offset+1) > start.
- **Range stringification**: concatenates text data from start Text node (partial), fully contained Text nodes, and end Text node (partial).
- **NodeIterator traversal**: toggles `beforeNode` flag and moves to next/previous node in the collection, filtering until `FILTER_ACCEPT`.
- **TreeWalker.children traversal**: first/last child, with `FILTER_SKIP` causing descent into children.
- **TreeWalker.siblings traversal**: navigating next/previous siblings, managing `FILTER_REJECT` and descent.
- **DOMTokenList operations**: `add`/`remove`/`toggle`/`replace` each check for empty string or ASCII whitespace; `toggle` and `replace` may skip update steps for compatibility; `supports` uses validation steps.
- **XPath**: `createExpression`, `evaluate` on `XPathEvaluatorBase`; `evaluate` on `XPathExpression` returns `XPathResult` with types like `ANY_TYPE`, `ORDERED_NODE_ITERATOR_TYPE`, etc.
- **XSLTProcessor**: methods for importing stylesheets, transforming to fragment/document, setting/getting parameters, resetting.

## Nuance Or Contradictions
- `detach()` on both `Range` and `NodeIterator` is a no‑op kept for compatibility only.
- `DOMTokenList` name is acknowledged as an “unfortunate legacy mishap”.
- Update steps for `toggle()` and `replace()` may not always run, explicitly noted for web compatibility.
- XPath and XSLT sections contain no complete algorithm definitions; they are placeholders pending further work (`whatwg/dom#67`, `whatwg/dom#181`).
- `createNSResolver(nodeResolver)` returns the node argument directly; it exists only for historical reasons.

## Candidate Wiki Hints
- **Range Modifications** – covering `deleteContents`, `extractContents`, `cloneContents`, and `insertNode`/`surroundContents`.
- **NodeIterator vs TreeWalker** – comparison of traversal models, `pointerBeforeReference`, and pre‑remove behaviors.
- **DOMTokenList** – token set manipulation, validation, and web‑compat quirks.
- **XPath and XSLT Legacy APIs** – summary of present but incomplete DOM Level 3 XPath and XSLT interfaces.
- **Range Point Comparisons** – `isPointInRange`, `comparePoint`, `intersectsNode`, and stringification.

## chunk-06

---
title: Chunk 06 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- Source file: `raw/web/corpus-2026-05-18/060-dom-standard.md`
- Chunk 6 of 9, lines 6888–8816
- Heading path: **11. Historical**
- Includes the historical removals section, acknowledgments, intellectual property rights, the full Index of terms, normative/informative references, and the IDL Index.

## Local Summary
This chunk enumerates every interface and interface member that has been removed from the DOM Standard. It then provides acknowledgments, copyright/licensing information, an exhaustive Index of all terms defined in the specification, references, and the complete IDL Index for the current standard.

## Key Claims
- The following interfaces have been removed from the DOM Standard:
  `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMImplementationSource`, `DOMLocator`, `DOMObject`, `DOMUserData`, `Entity`, `EntityReference`, `MutationEvent`, `MutationNameEvent`, `NameList`, `Notation`, `RangeException`, `TypeInfo`, `UserDataHandler`.
- The following interface members have been removed:

  **Attr**: `schemaTypeInfo`, `isId`

  **Document**: `createEntityReference()`, `xmlEncoding`, `xmlStandalone`, `xmlVersion`, `strictErrorChecking`, `domConfig`, `normalizeDocument()`, `renameNode()`

  **DocumentType**: `entities`, `notations`, `internalSubset`

  **DOMImplementation**: `getFeature()`

  **Element**: `schemaTypeInfo`, `setIdAttribute()`, `setIdAttributeNS()`, `setIdAttributeNode()`

  **Node**: `isSupported`, `getFeature()`, `getUserData()`, `setUserData()`

  **NodeIterator**: `expandEntityReferences`

  **Text**: `isElementContentWhitespace`, `replaceWholeText()`

  **TreeWalker**: `expandEntityReferences`
- The standard is written by Anne van Kesteren, with substantial contributions from Aryeh Gregor and Ms2ger.
- Copyright belongs to WHATWG (Apple, Google, Mozilla, Microsoft) under CC BY 4.0, with source code portions under BSD 3-Clause.
- The remaining content (Index, references, IDL Index) is the current complete definitional and cross-reference material for the entire DOM Standard.

## Entities And Concepts
- **Removed interfaces**: `DOMConfiguration`, `DOMError`, `DOMErrorHandler`, `DOMImplementationList`, `DOMImplementationSource`, `DOMLocator`, `DOMObject`, `DOMUserData`, `Entity`, `EntityReference`, `MutationEvent`, `MutationNameEvent`, `NameList`, `Notation`, `RangeException`, `TypeInfo`, `UserDataHandler`
- **Removed member groups**: by interface (see above)
- **Acknowledgments**: lists ~200 contributors
- **Intellectual property rights**: CC BY 4.0 for text, BSD 3-Clause for code; history in w3c/webcomponents repo under W3C Software and Document License
- **Index**: every term defined by the specification (alphabetical, with references to sections)
- **Terms defined by reference**: from other specifications (HTML, WebIDL, etc.)
- **References**: normative and informative
- **IDL Index**: complete Web IDL definitions of all current DOM interfaces

## Procedures And API Details
- None (this chunk documents removals, acknowledgments, and reference material, not operational procedures).

## Nuance Or Contradictions
- The list explicitly states that `ENTITY_REFERENCE_NODE` and `ENTITY_NODE` are “legacy” constants still present on `Node`, whereas the interfaces `Entity` and `EntityReference` are entirely removed.
- `Node`’s `ENTITY_REFERENCE_NODE` and `ENTITY_NODE` constants are marked “legacy” in the current IDL, even though the corresponding interfaces no longer exist.
- The acknowledgments and IP clauses are unchanged historical records, not normative requirements.

## Candidate Wiki Hints
- A wiki page **“Deprecated and removed DOM interfaces”** could list the removed interfaces and members along with their original version, reasons for removal, and migration notes.
- A **“DOM Standard history”** page could link to this removal list and explain the evolution of DOM levels.

## chunk-07

---
title: Chunk 07 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- **Source:** `raw/web/corpus-2026-05-18/060-dom-standard.md`
- **Chunk:** 7 of 9, lines 8818–11166
- **Heading path:** `11. Historical`
- **Coverage:** Full contents of the `Historical` section, combining WebIDL interface definitions with appended MDN browser‑compatibility tables.

## Local Summary
The chunk presents the DOM Standard’s “Historical” chapter. It defines legacy interfaces (`CDATASection`, `ProcessingInstruction`, `Comment`), the range and selection infrastructure (`AbstractRange`, `StaticRange`, `Range`), DOM tree traversal (`NodeIterator`, `TreeWalker`, `NodeFilter`), token list manipulation (`DOMTokenList`), and the XPath/XSLT API surface. Following the normative IDL, the chunk includes extensive MDN compatibility data for these interfaces as well as for `AbortController`/`AbortSignal`, `Attr`, `CharacterData`, `Document`, `Element`, `CustomEvent`, and `DOMImplementation`.

## Key Claims
- `CDATASection` is a historical interface (derives from `Text`).
- `ProcessingInstruction` and `Comment` are `CharacterData` subtypes.
- `AbstractRange` provides the read‑only basis for `StaticRange` (constructor with `StaticRangeInit`) and `Range` (mutable, with `commonAncestorContainer` and mutation methods).
- `Range` offers two‑point boundary manipulation, collapse, extraction/deletion/cloning/surrounding of contents, and a `detach()` method (historical).
- `NodeIterator` and `TreeWalker` allow filtered traversal; `NodeFilter` defines acceptance constants and `whatToShow` bitmask values, with several constants flagged as legacy (`SHOW_ENTITY_REFERENCE`, `SHOW_ENTITY`, `SHOW_NOTATION`).
- `DOMTokenList` represents a set of space‑separated tokens with methods (`add`, `remove`, `toggle`, `replace`, `supports`) and a stringifier `value`.
- `XPathResult` holds typed XPath evaluation results; `XPathExpression` evaluates an expression against a context node.
- `XSLTProcessor` supports importing a stylesheet and transforming a source document into a fragment or document.
- The appended compatibility tables indicate wide engine support for most DOM interfaces, with some features (e.g., `AbortSignal.timeout()`, `DOMTokenList.replace`, `AbstractRange` directly) having later adoption.

## Entities And Concepts
- **Historical interfaces:** `CDATASection`, `ProcessingInstruction`, `Comment`
- **Range model:** `AbstractRange`, `StaticRangeInit`, `StaticRange`, `Range`
- **Traversal:** `NodeIterator`, `TreeWalker`, `NodeFilter` (callback interface)
- **Token list:** `DOMTokenList`
- **XPath:** `XPathResult`, `XPathExpression`, `XPathNSResolver`, `XPathEvaluatorBase` mixin, `XPathEvaluator`
- **XSLT:** `XSLTProcessor`
- **Compatibility data:** Tables for `AbortController`, `AbortSignal`, `AbstractRange`/`Range`/`StaticRange` properties, `Attr`, `CharacterData` methods, `Comment`, `CustomEvent`, `DOMImplementation`, `DOMTokenList`, `Document` methods and properties, `XPathEvaluator`, `NodeIterator`/`TreeWalker` creation, etc.

## Procedures And API Details
- `Range` mutation: `setStart`/`setEnd`, `setStartBefore`/`setStartAfter`/`setEndBefore`/`setEndAfter`, `collapse`, `selectNode`, `selectNodeContents`, `deleteContents`, `extractContents`, `cloneContents`, `insertNode`, `surroundContents`. Boundary comparison constants `START_TO_START` (0), `START_TO_END` (1), `END_TO_END` (2), `END_TO_START` (3). `detach()` does nothing (legacy).
- `NodeIterator`/`TreeWalker`: filter using `NodeFilter.acceptNode(node)`, `whatToShow` bitmask. `TreeWalker` exposes `currentNode` and directional walk methods.
- `DOMTokenList`: `toggle(token, force)` returns whether token is present after toggling; `replace(old, new)` replaces a token; `supports(token)` checks if token is valid according to the associated attribute’s definition.
- `XPathResult`: result types defined as constants (`ANY_TYPE`, `NUMBER_TYPE`, …). `snapshotItem` and `iterateNext` for snapshots/iterators.
- `XSLTProcessor`: `importStylesheet(style)`, `transformToFragment(source, outputDocument)`, `transformToDocument(source)`, parameter management methods.

## Nuance Or Contradictions
- Several interfaces marked as historical or legacy (`CDATASection`, `detach()` methods, `SHOW_ENTITY_REFERENCE` etc.), yet they remain in the standard for reference.
- `XPathEvaluatorBase` is a mixin that both `Document` and `XPathEvaluator` include, allowing `evaluate` to be called directly on a document (legacy `createNSResolver` also available).
- Compatibility tables show near‑universal support for many interfaces, but some (like `AbortSignal.timeout()` static, `AbstractRange` independently instantiated) are noted with later browser versions.
- The source mixes formal IDL and informal MDN data; the MDN inclusion may not be part of the official DOM Standard, but it is present in this corpus dump.

## Candidate Wiki Hints
- “DOM Range Interface” — covers mutable selections, boundaries, manipulation.
- “NodeIterator and TreeWalker” — DOM tree traversal and filtering.
- “DOMTokenList” — managing space‑separated tokens (classList, relList, etc.).
- “XPath in the DOM” — evaluating XPath expressions and result types.
- “XSLTProcessor” — client‑side XSLT transformations.
- “Historical DOM Interfaces” — aliases and legacy types still defined for web compatibility.
- “DOM Living Standard — Historical chapter reference” — index of all interfaces in this section.

## chunk-08

---
title: Chunk 08 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- **Source:** raw/web/corpus-2026-05-18/060-dom-standard.md
- **Chunk:** 8 of 9
- **Lines:** 11168–13614
- **Heading path:** 11. Historical

## Local Summary
This chunk lists browser‑support tables for a large set of DOM interfaces and their members. Each entry states “In all current engines.”, then gives first‑release version numbers (or “?”) for Firefox, Safari, Chrome, Opera, Edge, Edge (Legacy), IE, and various mobile browsers. No normative prose or algorithms appear—just raw compatibility data.

## Key Claims
- Many DOM APIs are supported “in all current engines,” meaning they ship in modern Firefox, Safari, Chrome, Edge, and Opera.
- Legacy support varies widely:
  - Some APIs were available since very early browser releases (e.g., `Node/appendChild` since Firefox 1, Safari 1.1, Chrome 1, IE 5).
  - Others arrived much later (e.g., `MutationObserver` appeared around Firefox 14, Chrome 26, IE 11).
  - A few, such as `Element/assignedSlot`, `Event/composed`, and `EventTarget` constructor, are absent from IE entirely or marked `IENone`.
- The tables distinguish between “Edge” (Chromium‑based) and “Edge (Legacy)” (EdgeHTML‑based).

## Entities And Concepts
The listed APIs cover these interface families:
- **Document** – `prepend`, `querySelector`, `querySelectorAll`, `replaceChildren`, constructor
- **DocumentFragment** – `prepend`, `querySelector`, `querySelectorAll`, `replaceChildren`, `getElementById`, constructor
- **Element** – `prepend`, `querySelector`, `querySelectorAll`, `replaceChildren`, attribute methods, `checkVisibility`, `classList`, `className`, `closest`, `getAttribute*`, `hasAttribute*`, `id`, `insertAdjacent*`, `localName`, `matches`, `namespaceURI`, `prefix`, `removeAttribute*`, `setAttribute*`, `shadowRoot`, `slot`, `tagName`, `toggleAttribute`, constructor
- **DocumentType** – `name`, `publicId`, `systemId`, constructor
- **Text** – `assignedSlot`
- **Event** – constructor, `bubbles`, `cancelable`, `composed`, `composedPath`, `currentTarget`, `defaultPrevented`, `eventPhase`, `isTrusted`, `preventDefault`, `stopImmediatePropagation`, `stopPropagation`, `target`, `timeStamp`, `type`
- **EventTarget** – constructor, `addEventListener`, `dispatchEvent`, `removeEventListener`
- **HTMLCollection** – `item`, `length`, `namedItem`
- **HTMLSlotElement** – `slotchange` event
- **MutationObserver** – constructor, `disconnect`, `observe`, `takeRecords`
- **MutationRecord** – `addedNodes`, `attributeName`, `attributeNamespace`, `nextSibling`, `oldValue`, `previousSibling`, `removedNodes`, `target`, `type`
- **NamedNodeMap** – `getNamedItem`, `getNamedItemNS`, `item`, `length`, `removeNamedItem`, `removeNamedItemNS`, `setNamedItem`, `setNamedItemNS`
- **Node** – `appendChild`, `baseURI`, `childNodes`, `cloneNode`, `compareDocumentPosition`, `contains`, `firstChild`, `getRootNode`, `hasChildNodes`, `insertBefore`, `isConnected`, `isDefaultNamespace`, `isEqualNode`, `isSameNode`, `lastChild`, `lookupNamespaceURI`, `lookupPrefix`, `nextSibling`, `nodeName`, `nodeType`, `nodeValue`, `normalize`, `ownerDocument`, `parentElement`, `parentNode`, `previousSibling`, `removeChild`, `replaceChild`, `textContent`
- **NodeIterator** – `filter`, `nextNode`, `pointerBeforeReferenceNode`

A global attribute `slot` is also listed.

## Procedures And API Details
- No procedural steps are present. The section is purely a data dump of version‑to‑version support.
- Each entry follows the same pattern: API path, “In all current engines.”, then rows for Firefox, Safari, Chrome, Opera, Edge, Edge (Legacy), IE, and the corresponding mobile variants, sometimes with a Node.js version noted for `Event`‑ and `EventTarget`‑related items.

## Nuance Or Contradictions
- Several API entries have “?” for mobile Firefox, Safari, Chrome for Android, Samsung Internet, or Opera Mobile, indicating unknown or unverified support data.
- Some APIs that are “in all current engines” nevertheless list `IENone` (e.g., `Element/closest`, `Element/classList`, `Event/composed`) because IE is considered either out of support or not a “current engine” for the purposes of that statement.
- `Node/isConnected` lists `Samsung Internet6.0+` instead of “?”, a rare precise mobile version.
- `Event` and `EventTarget` sub‑entries include a Node.js version row (e.g., `14.5.0+`, `15.0.0+`), while most other APIs do not.

## Candidate Wiki Hints
- A centralized **DOM Browser Compatibility** page could be built from this data, linking each API to its MDN page and providing a quick lookup of when each was first supported. The source chunk already organizes the data per‑interface, which would make a table‑based wiki page straightforward.
- No other distinct new topic emerges from this chunk alone, as the information is reference material rather than a new concept or procedure.

## chunk-09

---
title: Chunk 09 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/060-dom-standard.md` (DOM Standard)
- Chunk: 9 of 9, lines 13616–15051
- Heading path: 11. Historical
- This chunk covers the entire “Historical” section of the standard and consists entirely of browser‑compatibility data for a large set of legacy DOM APIs.

## Local Summary
The chunk is a dense collection of compatibility tables (often identified by “✔MDN”), listing each API member and its support in current and legacy browser engines. No descriptive prose is present; the content is purely machine‑readable browser‑support information. Every entry follows the pattern: interface/member name, then a block of engine‑version lines (Firefox, Safari, Chrome, Opera, Edge, IE, and mobile equivalents).

## Key Claims
- The DOM Standard’s Historical section enumerates a wide range of interfaces and their members, and this chunk asserts their current implementation status.
- Nearly all listed historical APIs are marked “In all current engines”, implying broad availability despite their historical designation.
- A few features show later adoption (e.g., `ShadowRoot`, `StaticRange`, `Element/slot`) or platform limitations (e.g., IE missing `NodeIterator.referenceNode`, `NodeList.forEach`, `XPathEvaluator`, etc.).
- The chunk’s data is a verbatim extraction from the standard’s source, likely generated from the MDN Browser Compatibility Data format.

## Entities And Concepts
Prominent historical DOM interfaces whose compatibility data appear in this chunk:
- `NodeIterator` (and members: `previousNode`, `referenceNode`, `root`, `whatToShow`)
- `NodeList` (including `forEach`, `item`, `length`)
- `ProcessingInstruction` (`target`)
- `Range` (constructor and all manipulation/query methods: `cloneContents`, `cloneRange`, `collapse`, `commonAncestorContainer`, `compareBoundaryPoints`, `comparePoint`, `deleteContents`, `detach`, `extractContents`, `insertNode`, `intersectsNode`, `isPointInRange`, `selectNode`, `selectNodeContents`, `setEnd`, `setEndAfter`, `setEndBefore`, `setStart`, `setStartAfter`, `setStartBefore`, `surroundContents`, `toString`)
- `ShadowRoot` (`delegatesFocus`, `host`, `mode`, `slotAssignment`)
- `StaticRange` (constructor and interface)
- `Text` (constructor, `splitText`, `wholeText`)
- `TreeWalker` (`currentNode`, `filter`, `firstChild`, `lastChild`, `nextNode`, `nextSibling`, `parentNode`, `previousNode`, `previousSibling`, `root`, `whatToShow`)
- `XMLDocument`
- `XPathEvaluator` (constructor)
- `XPathExpression` (`evaluate`)
- `XPathResult` (`booleanValue`, `invalidIteratorState`, `iterateNext`, `numberValue`, `resultType`, `singleNodeValue`, `snapshotItem`, `snapshotLength`, `stringValue`)
- `XSLTProcessor` (constructor, `clearParameters`, `getParameter`, `importStylesheet`, `removeParameter`, `reset`, `setParameter`, `transformToDocument`, `transformToFragment`)
- `Element/slot` (now in the `shadow DOM` section, but appears under Historical in this raw source)

## Procedures And API Details
- No procedural steps or algorithmic descriptions are present; the chunk only reports compatibility version numbers.

## Nuance Or Contradictions
- The chunk shows many “?” entries for mobile‑browser variants (e.g., Firefox for Android, Safari for iOS, etc.), indicating missing or uncertain data in the original source.
- Some legacy engines (IE, Edge Legacy) are explicitly noted as “None” for certain methods (e.g., `range.comparePoint`, `range.isPointInRange`) while other historical APIs were supported even in IE5/IE9.
- The `Range.detach` line has a Firefox version range “1–15” while other lines just state a number; this suggests that `detach` was later removed from Firefox.
- The `ShadowRoot` interface and its members appear in the Historical section despite being part of modern shadow DOM; this may reflect a structural decision in the standard where all compatibility data is aggregated under “Historical”.

## Candidate Wiki Hints
- A source‑backed page could aggregate the compatibility snapshot for DOM Historical APIs, especially focusing on the “all current engines” status and notable exceptions.
- The raw data is effectively a snapshot of the MDN Browser Compatibility Data for historical interfaces; a wiki page could link to the specific entries as a reference for legacy feature support.

