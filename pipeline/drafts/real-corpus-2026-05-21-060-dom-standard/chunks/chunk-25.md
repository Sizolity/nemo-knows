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
