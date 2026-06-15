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
