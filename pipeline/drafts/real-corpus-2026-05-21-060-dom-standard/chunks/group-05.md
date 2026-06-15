---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Group Context

This group of notes covers **Section 11. Historical** of the DOM Standard, which serves as a comprehensive compatibility reference for the Document Object Model across major browser engines (Gecko/Firefox, WebKit/Safari, Blink/Chrome/Edge) and legacy environments (Internet Explorer, Edge Legacy). The content transitions from general infrastructure concepts in earlier chunks to detailed historical support matrices for specific interfaces (`Document`, `Element`, `Node`, `Event`, `MutationObserver`) and methods. It highlights the evolution of feature adoption, distinguishing between abstract base classes and concrete implementations, and notes significant version gaps between modern engines and legacy or mobile browsers.

# Cross-Chunk Summary

The documents detail the historical support status of DOM APIs, ranging from fundamental properties (e.g., `id`, `tagName`) to advanced manipulation methods (e.g., `replaceChildren`, `getRootNode`). A recurring theme is the distinction between "all current engines" (indicating modern universal support) and specific legacy version requirements. The notes consistently track support across desktop browsers, mobile webviews (Android WebView, iOS Safari), and Node.js environments. There is a clear separation in data regarding Edge Legacy versus Chromium-based Edge, with many modern features marked as unsupported or having limited support in the legacy Microsoft browser. Uncertainty is frequently denoted by question marks for mobile platforms where specific version data is unavailable.

# Repeated Or Central Claims

- **Universal Modern Support:** Core DOM Level 2 and 3 interfaces (e.g., `DocumentType`, `Element` attributes, `Node` properties) are supported in "all current engines" starting from the earliest versions of Firefox, Safari, and Chrome (often version 1+).
- **Legacy Browser Limitations:** Internet Explorer and Edge Legacy lack support for many modern APIs, often marked as "None" or requiring very specific early versions. Features like `baseURI` or `contains` are explicitly noted as unsupported in IE ("IENone").
- **Feature Adoption Gaps:** There are significant version gaps between browsers for newer features. For example, `replaceChildren` requires Firefox 78+ and Chrome 86+, while older methods like `appendChild` are available from version 1+.
- **Mobile Fragmentation:** Mobile browsers (Android WebView, iOS Safari) frequently show uncertain status ("?") or specific version splits compared to their desktop counterparts, indicating inconsistent implementation or lack of data.
- **Event and Mutation Evolution:** Properties like `composedPath` and methods like `takeRecords` have distinct adoption timelines, with `MutationObserver` appearing in Firefox 14/Chrome 26, while basic event handling is supported from earlier versions.

# Important Local Details

- **Interface Specifics:**
    - **Document:** Includes support for `createEvent`, `querySelector`, `getElementsByClassName`, `prepend`, and `replaceChildren`. XPath expression creation (`createExpression`) is widely supported but absent in IE.
    - **Element:** Covers core attributes (`id`, `className`, `classList`), methods (`getAttribute`, `insertAdjacentText`), and Shadow DOM features (`assignedSlot`, `attachShadow`). `classList` requires newer versions than the legacy string property `className`.
    - **Node:** Lists properties like `baseURI`, `childNodes`, and methods like `cloneNode`, `compareDocumentPosition`. `getRootNode` is noted as a modern addition requiring higher version numbers.
    - **Event:** Tracks properties (`cancelable`, `isTrusted`) and methods (`addEventListener`, `dispatchEvent`). `composedPath` is a later addition (Firefox 59+).
    - **MutationObserver:** Details support for `observe`, `disconnect`, and `takeRecords`.
- **Browser Versions:** Specific version thresholds are critical:
    - Firefox often lags behind Chrome in adopting certain modern APIs (e.g., `replaceChildren` vs. standard methods).
    - Safari generally supports features from versions 3–10 for core DOM, with higher numbers for Shadow DOM or specific event path properties.
    - Edge Legacy aligns with IE capabilities, while Chromium-based Edge aligns with Chrome.
- **Node.js:** Full compatibility with DOM APIs generally requires Node.js version 14.5.0 or higher.

# Candidate Wiki Hints

- **DOM Compatibility Matrix:** A central reference page explaining how to read the version tables, distinguishing between "In all current engines," legacy support, and mobile uncertainties.
- **Legacy Browser Migration Guide:** Documentation focusing on fallback strategies for Internet Explorer and Edge Legacy, highlighting APIs marked as unsupported or requiring polyfills (e.g., `querySelector` vs. `getElementsByClassName`).
- **Modern DOM Methods Guide:** A section dedicated to newer APIs like `prepend`, `replaceChildren`, `getRootNode`, and Shadow DOM slots (`assignedSlot`), detailing their introduction history.
- **Event API Evolution:** A page summarizing the timeline of Event properties, specifically highlighting the shift from basic event handling to composed paths and trusted checks.
- **MutationObserver Reference:** Detailed documentation on observing DOM changes, covering method availability across browsers and Node.js versions.

# Gaps Or Cautions

- **Mobile Data Uncertainty:** Many entries for Android WebView, iOS Safari, and Samsung Internet are marked with question marks (`?`), indicating that the source data lacks definitive compatibility information for these specific platforms at the time of documentation. Users should treat mobile support as potentially partial or unverified for newer APIs.
- **Legacy Edge Distinction:** The notes distinguish between "Edge (Legacy)" and modern "Edge." Features supported in Chromium-based Edge may not exist in the legacy version, which is based on IE. Entries marked "IENone" specifically indicate a lack of support in the legacy environment.
- **Version Gaps:** There are notable discrepancies where Firefox requires significantly higher versions than Chrome or Safari for certain features (e.g., `replaceChildren`), suggesting uneven implementation standards or delays in Gecko's adoption of newer specs compared to Blink/WebKit.
- **Abstract vs. Concrete Interfaces:** The data distinguishes between abstract base classes (like `AbstractRange`) and concrete implementations (`StaticRange`, `Range`). Features defined on abstract bases may have later release dates than their concrete counterparts, which is important for API design compatibility.
