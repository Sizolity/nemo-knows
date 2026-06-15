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
