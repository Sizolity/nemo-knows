---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

# Group Context

This group of notes covers the latter portion of the URL Standard specification, specifically focusing on **Section 6: API**. The content transitions from core URL manipulation logic into practical implementation details, including IDL (Web IDL) interface definitions for `URL` and `URLSearchParams`, naming conventions for standards authors, licensing information (Creative Commons Attribution 4.0 with BSD exceptions), and a comprehensive compatibility matrix detailing browser and runtime support across Firefox, Chrome, Safari, Edge, Opera, Node.js, and other environments. The document concludes with references to normative and informative standards like RFC 3986.

# Cross-Chunk Summary

The notes synthesize the evolution from theoretical interface design to practical deployment constraints.
- **Chunk 07** establishes the foundational rules for exposing APIs: URLs should be serialized as strings for standards, while objects are reserved for internal manipulation. It introduces the `URL` and `URLSearchParams` interfaces with their Web IDL definitions.
- **Chunk 08** expands on the compatibility landscape, providing granular version data for specific properties (e.g., `href`, `hostname`) and methods (e.g., `delete`, `sort`). It highlights discrepancies in legacy browser support (Edge/IE conflation) and uncertainty regarding Opera support.
- Together, these chunks define the complete lifecycle of the URL API within the standard: from the conceptual definition of the object and its stringification rules to the specific engine versions required to utilize its full feature set without polyfills.

# Repeated Or Central Claims

- **String Serialization Preference**: Standards should expose URLs as strings (serialized internal representations) rather than directly exposing URL objects, which are intended for internal manipulation.
- **Naming Conventions**: Features related to the `URL` interface should use lowercase "url" in naming (e.g., `newURL`), while compound names often prefer uppercase (e.g., `EventSource`).
- **Interface Definitions**: The core interfaces are `URL` and `URLSearchParams`. `URL` includes properties like `href`, `origin`, `protocol`, etc., and methods like `canParse`. `URLSearchParams` handles query string manipulation with methods like `append`, `delete`, `get`, `set`, and `sort`.
- **Licensing**: The specification is written by Anne van Kesteren, licensed under Creative Commons Attribution 4.0 International for the standard text, while source code portions use the BSD 3-Clause License.
- **Type Usage**: In Web IDL, the `USVString` type is used for URL values to handle Unicode scalar values correctly.

# Important Local Details

## Interface Definitions (Web IDL)

**URL Interface**
```webidl
[Exposed=*, LegacyWindowAlias=webkitURL]
interface URL {
  constructor(USVString url, optional USVString base);
  static URL? parse(USVString url, optional USVString base);
  static boolean canParse(USVString url, optional USVString base);

  stringifier attribute USVString href;
  readonly attribute USVString origin;
  attribute USVString protocol;
  attribute USVString username;
  attribute USVString password;
  attribute USVString host;
  attribute USVString hostname;
  attribute USVString port;
  attribute USVString pathname;
  attribute USVString search;
  [SameObject] readonly attribute URLSearchParams searchParams;
  attribute USVString hash;

  USVString toJSON();
};
```

**URLSearchParams Interface**
```webidl
[Exposed=*]
interface URLSearchParams {
  constructor(optional (sequence<sequence<USVString>> or record<USVString, USVString> or USVString) init = "");

  readonly attribute unsigned long size;

  undefined append(USVString name, USVString value);
  undefined delete(USVString name, optional USVString value);
  USVString? get(USVString name);
  sequence<USVString> getAll(USVString name);
  boolean has(USVString name, optional USVString value);
  undefined set(USVString name, USVString value);

  undefined sort();

  iterable<USVString, USVString>;
  stringifier;
};
```

## Browser and Runtime Compatibility Highlights

| Feature | Firefox | Chrome | Safari | Node.js | Edge (Legacy) | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **URL Interface** | 26+ | 19+ | 14.1+ | 10.0.0+ | 79+ (Chromium) | Core parsing and properties. |
| **URL.canParse** | 115+ | None* | 17+ | - | - | Static method; significant gap in Chrome support. |
| **URLSearchParams** | 29+ | 49+ | 10.1+ | 10.0.0+ | - | Core query string manipulation. |
| **`URL.pathname`** | 22+ | 32+ | 7+ | 7.5.0+ | - | Property access. |
| **`URLSearchParams.sort()`** | 54+ | 61+ | 11+ | 7.7.0+ | - | Requires polyfill for older engines. |

*\*Note: "None" indicates lack of support or specific constraints noted in the source data.*

## Nuances and Contradictions

- **Safari Versioning Discrepancy**: The `size` property on `URLSearchParams` is listed as supported from Safari 17, while other features appear from Safari 10.1, suggesting incremental feature rollout rather than simultaneous availability.
- **Edge Conflation**: Data for "Edge (Legacy)" often conflates Internet Explorer versions (e.g., IE10) with early Chromium-based Edge numbers, making version comparisons difficult without cross-referencing the engine lineage.
- **Opera Uncertainty**: Support data for Opera is frequently marked with a question mark (`?`), indicating that support levels are unverified or uncertain compared to other major engines.
- **canParse Gap**: `URL.canParse` is supported in modern Firefox and Safari but explicitly noted as unsupported ("None") in Chrome, highlighting a potential divergence in implementation priorities between engine vendors.

# Candidate Wiki Hints

- **Page: URL_API_Naming_Conventions**: Documenting the guidelines for naming URL-related features (lowercase "url" vs. uppercase compounds) to maintain consistency with standards like `EventSource` and `HashChangeEvent`.
- **Page: URL_API_Compatibility**: A comprehensive resource detailing specific versions of browsers and runtimes where `URL` and `URLSearchParams` became available, including legacy support notes.
- **Page: URLSearchParams_Methods_Polyfills**: Detailing which methods (specifically `sort()`) require polyfills for environments prior to Firefox 54 or Chrome 61.
- **Topic: Web_IDL_URL_Definition**: The formal IDL interface definitions for `URL` and `URLSearchParams`, including type usage (`USVString`) and exposed contexts.

# Gaps Or Cautions

- **Missing Implementation Details**: While the compatibility tables list version numbers, they do not explain *why* certain features (like `canParse`) are missing in specific engines (e.g., Chrome) beyond stating "None".
- **Legacy Edge Ambiguity**: The source data mixes Internet Explorer and legacy Edge support timelines. Users must be cautious when citing "Edge (Legacy)" versions, as they may refer to IE10/11 rather than early Chromium-based Edge.
- **Incomplete Opera Data**: Support for Opera is largely anecdotal or marked as uncertain in the provided notes; relying on these figures for production compatibility planning requires external verification.
- **Source Code License Distinction**: The notes mention that source code portions use BSD 3-Clause, distinct from the CC-BY-4.0 license of the standard text. This distinction is important for developers contributing to or distributing the specification's implementation.
