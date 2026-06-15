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
