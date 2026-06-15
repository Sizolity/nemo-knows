## group-01

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Group Context

This document group synthesizes the core structural and behavioral specifications of the DOM (Document Object Model) Standard. It covers the foundational infrastructure including tree hierarchies, ordered sets, selectors, and name validation rules. The notes progress through the event system lifecycle—from interface definitions (`Event`, `CustomEvent`) and listener management (`EventTarget`) to dispatching algorithms, propagation phases, and abort mechanisms (`AbortController`, `AbortSignal`). Finally, it details the node tree architecture, distinguishing between document trees and shadow DOM, mutation algorithms for structural changes, and specific interfaces for nodes, elements, attributes, ranges, traversal, XPath, XSLT, and historical considerations.

# Cross-Chunk Summary

The standard establishes a hierarchical tree model where nodes are organized into light trees (Document) and shadow trees, with complex interactions regarding slots and slottables within the latter. The event system is defined as a notification mechanism rather than an action initiator, utilizing `EventTarget` to manage listeners that can be passive, once-triggered, or signal-aborted. Events propagate in two phases: capturing (downwards) and bubbling (upwards), with specific handling for shadow DOM boundaries via the `composedPath()` algorithm. Structural changes to the tree are managed through distinct mutation algorithms that separate structural modification from side-effect execution, triggering custom element lifecycle callbacks and observer notifications. The specification also defines utility interfaces for text manipulation (Ranges), traversal (`NodeIterator`, `TreeWalker`), querying (XPath), and legacy compatibility layers.

# Repeated Or Central Claims

- **Tree Hierarchy**: The DOM is fundamentally a finite hierarchical tree structure where nodes have parents, children, and siblings, with relationships defined as inclusive or exclusive based on the specific property (e.g., ancestor vs. parent).
- **Event Semantics**: Events are objects that signal occurrences; they do not initiate actions. They can be synthetic (created by code) or native (dispatched by the user agent). Propagation involves traversing ancestors in two phases: capture and bubble.
- **Mutation Algorithms**: DOM mutations are handled via algorithms that separate "insertion steps" (structural changes) from "post-connection steps" (side effects like style application or script execution). This separation ensures atomicity for batch operations.
- **Shadow DOM Isolation**: Shadow trees are attached to light trees but can be closed to isolate their internal structure. Events and nodes within closed shadow trees are handled specifically, often requiring retargeting or filtering via `composedPath()`.
- **Abort Mechanism**: Asynchronous APIs should use `AbortController` and `AbortSignal` to support cancellation. Promises associated with these signals must reject immediately if the signal is aborted.
- **Legacy Compatibility**: The standard includes legacy extensions (e.g., `Window.event`, `initEvent`) marked as deprecated or replaceable to ensure backward compatibility while encouraging modern API usage.

# Important Local Details

- **Selectors**: Scoping-match is performed against parsed selector strings, but namespace support within selectors is explicitly not planned and will not be added in the future.
- **Name Validation**: Element local name validation has been loosened compared to strict XML specifications to allow names constructible by the HTML parser, though ASCII ranges are still restricted for historical reasons.
- **Ordered Sets**: These collections are parsed from whitespace-separated strings and serialized back using space delimiters, providing a simple mechanism for managing sets of tokens (e.g., class lists).
- **Event Options**: The `addEventListener` method accepts either a boolean or a dictionary for options. Booleans are treated as legacy flags for the `capture` state, while dictionaries allow explicit settings for `passive`, `once`, and `signal`.
- **Custom Element Lifecycle**: Custom elements trigger specific callbacks (`connectedCallback`, `disconnectedCallback`, `connectedMoveCallback`) upon connection or movement within the DOM tree.
- **DocumentFragment Optimization**: `DocumentFragment` nodes are optimized for batch insertion; their children are detached before re-insertion into a parent to ensure all side effects occur after the entire batch is connected.
- **NonElementParentNode Mixin**: The `getElementById()` method is exposed via this mixin only on `Document` and `DocumentFragment`, not on regular `Element` nodes, to prevent potential conflicts or unintended behavior.
- **Range Algorithms**: Ranges are defined by boundary points (before/after sets) rather than just offset positions, allowing for precise selection of node fragments.

# Candidate Wiki Hints

- **DOM Tree Fundamentals**: A page explaining parent/child/sibling relationships, tree order (preorder), and the distinction between light and shadow trees.
- **Event System Guide**: Comprehensive documentation on `Event`, `CustomEvent`, propagation phases, listener management, and control methods (`stopPropagation`, `preventDefault`).
- **Mutation Mechanics**: An article detailing the separation of insertion steps and post-connection steps, along with observer notifications and custom element lifecycle triggers.
- **Shadow DOM Architecture**: A guide covering slots, slottables, slot assignment algorithms, and how they interact with the document tree.
- **Abort API Usage**: Best practices for using `AbortController` and `AbortSignal` in promise-based APIs, including garbage collection rules for dependent signals.
- **Legacy vs Modern Events**: A reference page listing deprecated attributes (`srcElement`, `cancelBubble`, `Window.event`) and explaining why modern listeners are preferred.

# Gaps Or Cautions

- **Namespace Limitations**: Developers must be aware that CSS selectors within the DOM standard do not currently support namespaces, which may limit certain querying strategies.
- **Shadow DOM Closure**: When working with shadow trees, developers must handle closed shadow roots carefully, as events and nodes inside them may be inaccessible or require specific retargeting logic via `composedPath()`.
- **Passive Listener Performance**: The default passive value logic for touch/mousewheel events is complex; relying on passive listeners can optimize scrolling performance but may have unintended side effects if not configured correctly.
- **Garbage Collection Rules**: Dependent `AbortSignal` objects must remain alive while their source signals exist and active algorithms are running. Premature garbage collection could lead to unexpected behavior in asynchronous operations.
- **Historical Sections**: Chunks 20 through 32 cover "Historical" sections which may contain obsolete specifications or deprecated features; these should be treated as informational rather than current requirements for implementation.

## group-02

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Group Context

This group of notes synthesizes the core DOM specification covering infrastructure, event handling, node manipulation, and document structure. The content spans from high-level infrastructure concepts (trees, ordered sets, selectors) through detailed event interfaces (Event, EventTarget, AbortController), to comprehensive node tree mechanics including shadow DOM slots, mutation observers, and node traversal. It concludes with definitions for Document, Element, and specialized interfaces like XPath and XSLT processors, alongside a historical section documenting legacy features.

## Cross-Chunk Summary

The document is structured into four primary domains:
1.  **Infrastructure (Section 1)**: Defines the foundational concepts of the DOM including tree structures, ordered sets for collections, selector mechanisms, and name validation rules.
2.  **Events (Section 2)**: Details the lifecycle of events from introduction to interface definitions (`Event`, `CustomEvent`), listener observation via `EventTarget`, dispatching logic, and the abort mechanism using `AbortController`/`AbortSignal`.
3.  **Nodes & Trees (Sections 4 & 5)**: Describes the node hierarchy, including document and shadow trees, slot mechanics for custom elements, mutation algorithms, mixin interfaces (`ParentNode`, `ChildNode`, `Slottable`), legacy collections (`NodeList`, `HTMLCollection`), mutation observers, and specific node interfaces (`Node`, `Document`, `Element`).
4.  **Advanced & Legacy (Sections 6-11)**: Covers ranges, traversal algorithms (`NodeIterator`, `TreeWalker`), sets, XPath/XSLT processing, security considerations, and a substantial historical section tracking deprecated features and legacy aliases.

## Repeated Or Central Claims

*   **Mixin Architecture**: The DOM relies heavily on mixin interfaces to share functionality between disparate node types. Specifically, `DocumentOrShadowRoot`, `ParentNode`, `ChildNode`, and `Slottable` allow shared methods like `prepend`, `append`, and custom element registry access without bloating individual interface definitions.
*   **Mutation Observer Microtask Queue**: Mutation observers utilize an internal microtask queue to ensure that callback functions are invoked even if the DOM tree changes during the notification process, preventing skipped updates or race conditions.
*   **Shadow Tree Integration**: The specification explicitly handles the integration of shadow DOMs into the main document tree. Concepts like `getRootNode({ composed: true })`, shadow-including traversal orders, and slot assignment logic are central to managing the composition of light and shadow trees.
*   **Legacy vs. Modern Collections**: There is a strong distinction between modern iterable collections (sequences) and legacy artifacts. `HTMLCollection` and `NodeList` are treated as historical; new API designers are advised to use standard iterables, though backward compatibility logic preserves their behavior.
*   **Error Handling Consistency**: Specific DOMExceptions are consistently used for structural violations: `HierarchyRequestError` for tree constraint breaches (e.g., moving nodes across documents), and `NotSupportedError` for unsupported operations (e.g., cloning shadow roots, invalid element creation).

## Important Local Details

*   **Node Type Constants**: The `Node` interface uses unsigned short constants to represent node types (e.g., `ELEMENT_NODE = 1`, `TEXT_NODE`). The `nodeName` property returns specific strings like "#text" for Text nodes or null/qualified names for others.
*   **Document Attributes**: Key document properties include `URL` (document URI), `compatMode` (returns "BackCompat" or "CSS1Compat"), `characterSet`, and `doctype`. The default encoding is UTF-8, and the content type is "application/xml".
*   **Element Creation Logic**: When creating elements via `createElement()`, the local name is lowercased in HTML documents. Options allow passing a custom element registry or an `is` flag to customize built-in elements.
*   **Class Matching Nuance**: The `getElementsByClassName` method requires all space-separated classes to be present on an element. Commas within class names (e.g., `"aaa,bbb"`) do not act as delimiters; they are treated as part of the class name string.
*   **ShadowRoot Attributes**: Shadow roots possess specific attributes such as `mode` ("open" or "closed"), `delegatesFocus`, `slotAssignment`, and a reference to the host node (`host`).
*   **Namespace Resolution**: The logic for resolving namespace prefixes involves checking the element's own namespace, an "xmlns" attribute, or recursively delegating to the parent element. Empty string prefixes are converted to null during resolution.

## Candidate Wiki Hints

*   **Mixin Interfaces Deep Dive**: A dedicated page explaining how `DocumentOrShadowRoot`, `ParentNode`, and other mixins function to provide shared methods across different node types.
*   **MutationObserver Lifecycle**: Documentation covering the constructor, `observe()` validation rules (including automatic enabling of options), `disconnect()`, and the internal queuing logic for records.
*   **Shadow DOM Composition**: A guide on managing shadow trees, including slot assignment, finding slottables, signaling slot changes, and traversing shadow-including trees.
*   **Document Interface Properties**: An overview of `Document` attributes like `compatMode`, `URL`, and the adoption algorithm for moving nodes between documents.
*   **Legacy Collection Handling**: A section explaining `NodeList` vs. `HTMLCollection`, their live collection behavior, and why they are considered legacy artifacts in modern API design.

## Gaps Or Cautions

*   **Incomplete Historical Content**: Chunks 20 through 32 cover a "Historical" section with repetitive headings but no specific sub-headers or content details provided in the notes. This suggests a large block of deprecated features or legacy definitions that are not fully synthesized here.
*   **Missing Range Algorithms**: While ranges are mentioned, the specific algorithms for range creation and manipulation (beyond basic `createRange`) are less detailed compared to node mutation logic.
*   **XPath/XSLT Scope**: The coverage of XPath and XSLT is limited to interface definitions (`XPathResult`, `XPathExpression`) and processor interfaces without deep algorithmic detail on query execution or transformation steps.
*   **Security Considerations Vague**: Section 10 on "Security and privacy considerations" is listed as a heading but lacks specific claims or details in the provided notes, representing a potential gap for a security-focused wiki entry.

## group-03

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Group Context

This group of notes synthesizes the **DOM Standard** specification, specifically covering the transition from general infrastructure to detailed node tree manipulation, event systems, and legacy API maintenance. The source material spans from fundamental tree structures (Trees, Ordered Sets) through the comprehensive definition of the `Element` interface, attribute handling, shadow DOM integration, range selection algorithms, and traversal mechanisms (`NodeIterator`, `TreeWalker`). The latter portions of the document address legacy features such as XPath/XSLT and historical notes, while maintaining conformance definitions for modern APIs like Custom Elements and Mutation Observers.

# Cross-Chunk Summary

The document is structured as a progression from abstract infrastructure to concrete DOM interfaces:

1.  **Infrastructure (Chunks 01–04):** Establishes the foundational concepts of trees, ordered sets, selectors, and name validation. It introduces the event system lifecycle: from introduction to interface definitions (`Event`, `CustomEvent`), target management (`EventTarget`), listener observation, dispatching/firing, and finally, activity abortion via `AbortController`/`AbortSignal`.
2.  **Node Tree & Mixins (Chunks 05–08):** Defines the hierarchy of node interfaces. It details the `Node` interface, document/shadow tree structures (including slots/slottables), mutation algorithms, and a suite of mixin interfaces (`NonElementParentNode`, `DocumentOrShadowRoot`, `ParentNode`, `ChildNode`, `Slottable`). This section also covers old-style collections (`NodeList`, `HTMLCollection`) and introduces Mutation Observers.
3.  **Core Interfaces (Chunks 09–15):** Focuses on specific node types (`Node`, `Document`, `DocumentType`, `DocumentFragment`, `ShadowRoot`, `Element`). It provides deep dives into attribute handling via `NamedNodeMap` and `Attr`, character data nodes (`Text`, `Comment`, etc.), and the introduction of DOM Ranges.
4.  **Range & Traversal (Chunks 16–18):** Details the `Range` interface, including boundary points, containment logic, and manipulation algorithms (delete, extract, clone). It covers traversal interfaces (`NodeIterator`, `TreeWalker`) with filtering logic via `NodeFilter` and the set-like behavior of `DOMTokenList`.
5.  **Legacy & Security (Chunks 19–32):** Concludes with XPath/XSLT interfaces, security/privacy considerations, and a large section dedicated to historical context regarding deprecated or legacy features.

# Repeated Or Central Claims

- **Element Definition:** An element is fundamentally defined by its custom element state ("uncustomized" or "custom"). It possesses a tag name, optional namespace/prefix, and an associated shadow root (null by default).
- **Attribute Semantics:** Attributes like `id`, `class`, and `slot` are super-global content attributes. The `id` concept is formalized as unique per element (replacing historical DTD-based multiple identifiers). Qualified names are computed from local name and prefix, with HTML documents requiring uppercase normalization for tag names.
- **Live Ranges:** `Range` objects are "live," meaning they automatically update their boundary points when the underlying DOM tree mutates (insertion, removal, replacement). `StaticRange` is the exception that does not update.
- **Shadow Host Restrictions:** Only elements with valid local names (e.g., standard HTML elements or registered custom elements) can act as shadow hosts. Attaching a shadow root to an invalid host throws a `NotSupportedError`. The `shadowRoot` getter returns null if the mode is "closed".
- **Mutation & Callbacks:** Changing attributes triggers mutation records. For custom elements, this also enqueues upgrade reactions and invokes `attributeChangedCallback` if the element is defined.
- **Traversal Mechanics:** Both `NodeIterator` and `TreeWalker` rely on a `whatToShow` bitmask and a `filter` (via `NodeFilter`) to determine node visibility. The `detach()` method exists in both for compatibility but performs no action in the current spec.
- **DOMTokenList Validation:** Token lists strictly forbid empty strings and tokens containing ASCII whitespace, throwing `SyntaxError` or `InvalidCharacterError` respectively.

# Important Local Details

- **Interface Hierarchy:** The spec uses a mixin approach where interfaces like `ParentNode`, `ChildNode`, and `NonDocumentTypeChildNode` are layered onto base nodes to provide specific traversal capabilities (e.g., `closest()`, `insertBefore()`).
- **Range Algorithms:** Operations like `deleteContents()` and `extractContents()` involve complex logic to reconstruct start/end nodes if they become invalid after tree modification. `surroundContents()` throws errors for partial containment of non-Text nodes or invalid parent types.
- **Shadow DOM Init:** The `attachShadow` method accepts a dictionary (`ShadowRootInit`) allowing configuration of mode, delegates focus, and slot assignment options. `customElementRegistry` can be passed to pass a node directly in some contexts.
- **Boundary Point Logic:** A range's containment is strictly defined by offsets within nodes. A node is *not* contained if it defines the boundary itself; only its descendants (or content for CharacterData) are included.
- **Historical Aliases:** The spec retains legacy aliases such as `webkitMatchesSelector` alongside standard `matches`, and `insertAdjacentElement`/`insertAdjacentText` with legacy string mappings.
- **XPath Status:** While defined in the DOM spec, XPath Level 3 APIs (`XPathResult`, `XPathEvaluator`) are noted as legacy and not actively maintained, though definitions persist for future updates.

# Candidate Wiki Hints

- **Page: Element Interface Overview**: Summarize attributes, shadow root association, and custom element lifecycle states.
- **Page: Custom Element Lifecycle**: Explain the transition from "undefined" to "custom", upgrade mechanisms, and reaction handling.
- **Page: DOM Range Selection**: Detail live vs. static ranges, boundary points, and containment rules.
- **Page: Traversal Interfaces**: Compare `NodeIterator` (fixed root) vs. `TreeWalker` (mutable current node).
- **Page: Shadow DOM Host Requirements**: List valid element names and conditions for attaching shadow roots.
- **Page: Attribute Manipulation**: Guide on `setAttribute`, `removeAttribute`, and trusted type validation steps.

# Gaps Or Cautions

- **Missing Content Data:** The provided chunk notes summarize headings and claims but do not include the full text of algorithms or specific error message strings (e.g., exact `DOMException` codes beyond names).
- **Incomplete Historical Section:** Chunks 20–32 cover "Historical" content but the notes only indicate the heading. Specific details on deprecated APIs within this section are not elaborated in the provided summaries.
- **Algorithm Implementation Details:** While high-level logic for ranges and traversal is described, specific internal algorithm steps (e.g., exact order of mutation record queuing) are inferred from context rather than explicitly detailed in the notes.
- **Security Specifics:** The "Security and privacy considerations" section is listed but lacks detailed content in the provided notes, requiring cross-referencing with external security documentation for implementation guidance.

## group-04

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## Group Context

This document group synthesizes the **DOM Standard Living Specification**, specifically focusing on the transition from legacy interfaces to modern APIs. The content spans infrastructure concepts (trees, ordered sets), a comprehensive eventing system (Event, EventTarget, AbortController), node manipulation (Node, Document, Element, MutationObserver), and historical sections detailing deprecated features.

The group covers the specification's evolution, noting that while many legacy interfaces (e.g., `DOMConfiguration`, `EntityReference`) have been removed or marked as historical, core APIs like `AbortController` and `MutationObserver` are central to modern web development. The document also includes a significant portion dedicated to XPath and XSLT processing, which remains part of the standard but is often considered legacy for HTML contexts.

## Cross-Chunk Summary

The specification is structured into logical layers:
1.  **Infrastructure & Events**: Defines the foundational tree structure, event loops, and the `AbortController` mechanism for managing asynchronous operations.
2.  **Node Manipulation**: Details the `Node` hierarchy, including specific mixins (`ParentNode`, `ChildNode`) that enable fluent DOM manipulation APIs like `.appendChild()` and `.querySelector()`.
3.  **Mutation & Traversal**: Introduces `MutationObserver` for performance-friendly change detection and `TreeWalker`/`NodeIterator` for structured traversal.
4.  **Document & Element**: Defines the document tree, shadow DOM capabilities (slots, slottables), and element attribute handling via `NamedNodeMap`.
5.  **Legacy & Historical**: A dedicated section (§11) catalogues removed interfaces (e.g., `DOMError`, `MutationEvent`) and deprecated methods, distinguishing them from their modern replacements or noting their status in the living standard.

## Repeated Or Central Claims

-   **Historical Section (§11)**: Multiple chunks repeatedly define a section dedicated to "Historical" interfaces and members. These are explicitly removed from the current standard or retained only for legacy compatibility (e.g., `DOMConfiguration`, `EntityReference`, `MutationEvent`).
-   **Abort Mechanism**: The `AbortController` and `AbortSignal` pair is consistently described as the modern, preferred method for aborting ongoing activities (fetches, timers), replacing older mechanisms.
-   **Mixin Architecture**: The document frequently references mixin interfaces (`ParentNode`, `ChildNode`, `NonElementParentNode`) that augment the base `Node` interface to provide specific manipulation capabilities without bloating the core interface.
-   **Event System**: The distinction between legacy event handling (e.g., `MutationEvent`) and modern observation patterns (`MutationObserver`) is a recurring theme, emphasizing performance and cleaner separation of concerns.

## Important Local Details

-   **XPath & XSLT**: Despite being marked as historical or legacy in some contexts, the specification maintains full definitions for `XPathResult`, `XPathEvaluator`, and `XSLTProcessor`. These interfaces allow for XML transformation and querying within the DOM environment.
-   **Shadow DOM Internals**: Detailed coverage of Shadow DOM includes concepts like `slots`, `slottables`, and shadow tree assignment modes (`manual`, `named`), alongside the `serializable` attribute on `ShadowRootInit`.
-   **Node Types**: Specific node type constants are defined (e.g., `ELEMENT_NODE`, `TEXT_NODE`, `DOCUMENT_FRAGMENT_NODE`) and used throughout the tree structure definitions.
-   **Range API**: The `Range` interface is detailed with methods for setting boundaries (`setStart`, `setEnd`) and manipulating contents (`extractContents`, `deleteContents`), serving as the basis for selection and replacement operations.
-   **Browser Compatibility Nuance**: While core features like `AbortController` are supported in all current engines, specific properties (e.g., `signal.reason`) have versioned support requirements (e.g., Firefox 97+).

## Candidate Wiki Hints

-   **Page: DOM Standard Overview**
    -   **Content**: High-level summary of the Living Standard structure, distinguishing between active APIs and historical sections.
-   **Page: Event System & Abort Signals**
    -   **Content**: Guide on using `EventTarget`, `CustomEvent`, and `AbortController` for managing async tasks and event propagation phases.
-   **Page: Mutation Observation**
    -   **Content**: Explanation of `MutationObserver` usage, microtask queuing, and the structure of `MutationRecord`.
-   **Page: Shadow DOM & Slots**
    -   **Content**: Deep dive into Shadow Root modes, slot assignment algorithms, and slottable element requirements.
-   **Page: Historical DOM Interfaces**
    -   **Content**: Catalog of removed interfaces (`DOMError`, `EntityReference`) and deprecated methods with migration paths or context.
-   **Page: XPath & XSLT in DOM**
    -   **Content**: Overview of the legacy XML processing APIs available within the DOM environment.

## Gaps Or Cautions

-   **Incomplete Historical Context**: While Chunk 20–32 cover the "Historical" section extensively, they often list interfaces without detailed algorithmic descriptions for every removed member. Users must cross-reference with original spec versions for full algorithmic understanding of deprecated methods like `createEntityReference`.
-   **Versioned Property Support**: Claims that features are supported in "all current engines" may be slightly misleading regarding specific properties (e.g., `signal.reason`) which have strict version gates. Developers should check browser compatibility tables for these specific nuances.
-   **XPath/XSLT Relevance**: The inclusion of XPath and XSLT might confuse readers expecting an HTML-focused DOM guide. These sections are retained for XML processing but are less relevant for typical web page manipulation without specific use cases.
-   **Legacy Terminology**: Terms like "legacy-canceled-activation behavior" in Service Workers appear in historical contexts; these should not be used in new implementations as they refer to transitional states replaced by the modern lifecycle model.

## group-05

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

## group-06

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Group Context

This group of notes synthesizes the concluding sections of the DOM Standard document, specifically focusing on **Section 11: Historical**. This section aggregates extensive browser compatibility data for a wide range of DOM interfaces and methods previously detailed in the main standard. The coverage includes traversal APIs (`TreeWalker`, `NodeIterator`), query languages (`XPath`, `XSLT`), legacy collection types (`NodeList`, `HTMLCollection`), Shadow DOM properties, and various text/range manipulation methods.

The data serves to map the implementation timeline of these features across major browser engines (Firefox, Safari, Chrome, Opera, Edge) and their mobile counterparts (Android WebView, iOS Safari, Samsung Internet). A significant portion of this group addresses the dichotomy between modern Chromium-based browsers and legacy environments (Internet Explorer, Edge Legacy), as well as the varying levels of support in mobile WebViews compared to desktop versions.

# Cross-Chunk Summary

The progression from Chunk 31 to Chunk 32 shifts focus from general DOM interfaces and Range/Text manipulation details toward advanced query and transformation APIs.

*   **Chunk 31** covers broad compatibility tables for core interfaces like `NodeList`, `Range`, `ShadowRoot`, and `Text`. It highlights early adoption (e.g., Firefox 1, Chrome 1) for basic features but notes specific version requirements for Shadow DOM properties (Firefox 63+, Chrome 53+). It explicitly distinguishes between "Edge" and "Edge (Legacy)," marking the latter as lacking support for modern standards.
*   **Chunk 32** narrows the scope to advanced traversal (`TreeWalker`), query languages (`XPath`, `XSLT`), and specific Shadow DOM attributes (`Element.slot`). It reinforces that while `TreeWalker` is universally supported in current engines, `XPath` and `XSLT` are absent or limited in legacy Edge and IE.

Together, these chunks provide a comprehensive historical map of the DOM API ecosystem, identifying which features are stable across all modern browsers and which remain fragmented due to legacy engine constraints or mobile WebView implementations.

# Repeated Or Central Claims

*   **Universal Modern Support:** Core interfaces such as `TreeWalker` and basic properties like `Element.slot` (Shadow DOM v1) are supported in all current engines, though specific methods within these interfaces may have higher minimum version requirements (e.g., Firefox 92 for certain Shadow DOM features).
*   **Legacy Engine Fragmentation:** There is a consistent distinction between "Edge" (Chromium-based) and "Edge (Legacy)" (IE-based). Legacy Edge and Internet Explorer consistently show no support ("None") or partial/unknown support ("?") for modern features like XPath, XSLT, Shadow DOM properties, and advanced Range methods.
*   **Mobile WebView Variability:** Mobile WebViews (Firefox for Android, iOS Safari, Chrome for Android) frequently display missing data ("?") or specific version gaps compared to their desktop counterparts. Samsung Internet is also noted with varying support levels depending on the underlying WebView version.
*   **Early Adoption vs. Removal:** While many features are listed as supported in "all current engines," the inclusion of very early version numbers (e.g., Firefox 1, Safari 1) implies long-term stability for core APIs, whereas specific methods like `Range.deleteContents` show historical support ranges that may imply deprecation or removal in later versions.

# Important Local Details

*   **TreeWalker Interface:**
    *   **Properties/Methods:** `currentNode`, `filter`, `firstChild`, `lastChild`, `nextNode`, `nextSibling`, `parentNode`, `previousNode`, `previousSibling`, `root`, `whatToShow`.
    *   **Usage:** Instantiated with a root node, optional filter function, and `whatToShow` flag.
*   **XPath API:**
    *   **Interfaces:** `XPathEvaluator`, `XPathExpression`, `XPathResult`.
    *   **Properties:** `booleanValue`, `numberValue`, `stringValue`, `singleNodeValue`, `snapshotItem`, `snapshotLength`.
    *   **Usage:** `XPathEvaluator.evaluate()` returns an `XPathResult` object for querying.
*   **XSLT API:**
    *   **Interface:** `XSLTProcessor`.
    *   **Methods:** `clearParameters`, `getParameter`, `importStylesheet`, `removeParameter`, `reset`, `setParameter`, `transformToDocument`, `transformToFragment`.
*   **Shadow DOM & Slots:**
    *   **Property:** `Element.slot` (indicates Shadow DOM v1 support).
    *   **History:** Supported from Firefox 63, Safari 10.1, Chrome 53.
    *   **Usage:** Assign a slot name to an element in a Shadow Root (e.g., `element.slot = "name"`).
*   **Range & Text APIs:**
    *   **Methods:** `toString`, `deleteContents`, `extractContents`.
    *   **Note:** `deleteContents` had specific Firefox version ranges (1–15) in historical data, suggesting potential removal or change.
    *   **Properties:** `Text.wholeText`.
*   **Legacy Collections:**
    *   `NodeList` and `HTMLCollection` are identified as "Old-style collections."
    *   `NodeList/forEach` supported in Firefox 50+, Safari 10+, Chrome 51+.

# Candidate Wiki Hints

*   **DOM API Compatibility Tables:** A summary page listing support status for core DOM interfaces across browsers.
*   **TreeWalker API Reference:** Documentation covering properties and filtering examples for `TreeWalker`.
*   **XPath in JavaScript:** A guide on using `XPathEvaluator` and `XPathResult` for DOM queries.
*   **XSLT Transformation Guide:** Instructions for using `XSLTProcessor` to transform XML/HTML documents.
*   **Shadow DOM Slots:** An article covering the `slot` property, light DOM distribution, and browser compatibility.
*   **Legacy Browser Support:** A section detailing the differences between modern Edge and Edge Legacy, as well as IE limitations.

# Gaps Or Cautions

*   **Incomplete Mobile Data:** Many entries for mobile WebViews (Firefox for Android, iOS Safari, Chrome for Android) are marked with "?", indicating incomplete tracking or unknown status compared to desktop browsers.
*   **Edge Discrepancy:** Care must be taken when referencing "Edge" compatibility; data often conflates the modern Chromium version with the legacy IE-based version, which lacks significant support for newer standards.
*   **Historical Context:** The presence of very early version numbers (e.g., Firefox 1) alongside notes on method removal or change implies that some APIs have evolved significantly or been deprecated since their initial introduction.
*   **Data Consistency:** The source distinguishes between "supported" (✔MDN) and "unknown/unlisted" (?), suggesting that the absence of a checkmark does not necessarily mean unsupported, but rather unverified in the context of the data compilation.

