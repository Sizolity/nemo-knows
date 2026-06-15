## group-01

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

# Group Context

This group of notes synthesizes the **URL Standard** (Living Standard, last updated April 14, 2026), covering infrastructure requirements, host representations, URL parsing and serialization, origin determination, secure rendering guidelines, and the `application/x-www-form-urlencoded` format. The standard aims to unify modern implementations, obsoleting legacy RFCs (3986, 3987) while establishing robust APIs (`URL`, `URLSearchParams`) and strict security boundaries against visual spoofing and data leakage.

# Cross-Chunk Summary

The document defines a complete lifecycle for network identifiers:
1.  **Infrastructure:** Establishes the parser model, percent-encoding rules, and error handling (validation errors vs. termination).
2.  **Hosts:** Defines how domains, IP addresses (IPv4/IPv6), and opaque identifiers are represented, parsed, serialized, and checked for equivalence. It details IDNA processing and public suffix logic.
3.  **URL Structure:** Breaks down the URL record into components (scheme, username, password, host, port, path, query, fragment). It defines valid schemes, special handling for `file` URLs, and Windows drive letter quirks.
4.  **Parsing & Serialization:** Provides state machine algorithms for parsing scalar strings into URL records and serializing them back to ASCII. Equivalence is determined by serialized comparison (excluding fragments).
5.  **Origins & Rendering:** Defines how to extract origins (tuple vs. opaque) and guidelines for secure display (hiding credentials, simplifying hosts, handling IDNs).
6.  **APIs:** Exposes the `URL` and `URLSearchParams` interfaces, detailing constructor behavior, property accessors/setters, and the specific encoding differences between query strings and general URL paths.

# Repeated Or Central Claims

*   **Interoperability over Legacy:** The standard prioritizes alignment with modern implementations over strict adherence to obsolete RFCs, explicitly addressing issues like spaces and "illegal" code points that caused fragmentation in the past.
*   **Validation Errors vs. Termination:** A critical distinction is made between a validation error (which signals a mismatch but does not stop the parser) and parser termination (which must be explicit). This allows parsers to report errors without crashing or silently failing.
*   **Security First (Spoofing):** The standard treats URL passing as an untrusted operation. Rendering rules strictly prohibit displaying credentials, emphasize showing registrable domains over full hosts, and mandate handling of invisible code points (e.g., LRE, RLE) and visual confusables to prevent homograph attacks.
*   **Percent-Encoding Sets:** Different components use different encoding sets (`component`, `application/x-www-form-urlencoded`). The standard emphasizes that only specific sets result in roundtripable data; others require careful handling of the literal `%` character.
*   **Host Representation Consistency:** Hosts are case-insensitive domains or IP addresses by default, but serialization with `isOpaque=true` preserves original casing (e.g., `EXAMPLE.COM`). This creates distinct string representations that map to different hosts depending on parsing mode.

# Important Local Details

*   **Percent-Decoding Algorithm:**
    1.  Initialize output as an empty byte sequence.
    2.  Iterate through input bytes. If not `%` (0x25), append to output.
    3.  If `%`, check the next two bytes. If invalid hex, treat `%` as literal.
    4.  If valid hex, interpret the pair and append value; skip the two bytes.
*   **Host Serialization Logic:**
    *   **IPv4:** Reverses octets by prepending `n % 256` and dividing `n` by 256 repeatedly.
    *   **IPv6:** Finds the longest run of zeros for compression (`::`). Iterates pieces, skipping leading zeros if compressing, appending shortest lowercase hex, and handling delimiters carefully.
*   **URL Parsing State Machine:**
    *   **Scheme Start:** Requires ASCII alpha; moves to scheme state or `no scheme` state.
    *   **No Scheme:** If base URL has an opaque path, reject unless fragment (`#`) starts input. Otherwise, inherit base scheme if applicable.
    *   **File URLs:** Must start with `//`. Handles Windows drive letter quirks (replacing second code point with `:`).
    *   **Path State:** Processes `.`, `..` segments. Rejects backslashes (`\`) in non-file contexts (`invalid-reverse-solidus`).
*   **Origin Extraction Rules:**
    *   **Tuple Origins:** `(scheme, host, port, null)` for `http`, `https`, `ws`, `wss`, `ftp`.
    *   **Blob Origins:** Environment origin if available; otherwise opaque.
    *   **File Origins:** Left as an exercise (often opaque).
*   **`application/x-www-form-urlencoded`:**
    *   **Parsing:** Splits on `&`, replaces `+` with space, percent-decodes UTF-8 without BOM.
    *   **Serialization:** Percent-encodes using the form set (space becomes `%20` or `+` depending on context), joins with `=` and `&`.

# Candidate Wiki Hints

*   **URL Standard Goals & Terminology:** Page summarizing the shift from RFCs to modern implementations, terminology changes (URI vs. URL), and infrastructure dependencies.
*   **Host Representation Types:** Page explaining Domain, IP Address, Opaque Host, and Empty Host, including serialization modes (`isOpaque`).
*   **IDNA & Public Suffix Logic:** Guide to International Domain Name processing (UTS46) and the algorithm for determining registrable domains vs. public suffixes.
*   **URL Parsing State Machine:** Detailed breakdown of parser states (`scheme start`, `relative`, `path`, etc.) and transition rules.
*   **Secure URL Rendering:** Best practices for address bar display, including hiding credentials, simplifying hosts (elision), and handling IDNs/confusables.
*   **application/x-www-form-urlencoded Format:** Canonical algorithm for parsing/serializing form data, highlighting differences from standard URL percent-encoding.
*   **URL vs. URLSearchParams API:** Documentation on `href`, `origin` getters/setters, opaque path behaviors, and query parameter manipulation.

# Gaps Or Cautions

*   **File Origin Ambiguity:** The origin for `file` URLs is explicitly left as an exercise to the reader, potentially leading to inconsistency across implementations unless defined elsewhere.
*   **Public Suffix Reliability:** The concept of "public suffix" and "registrable domain" should not be relied upon for hard security boundaries because Public Suffix List (PSL) lists may diverge between clients. The "origin" concept is preferred for security.
*   **Credential Serialization Contradiction:** There is a noted inconsistency where the text states credentials cannot be expressed in valid URL strings, yet examples show parsing behavior that might accept them depending on context, suggesting specific constraints during serialization vs. parsing.
*   **Encoding Discrepancies:** The `URL` interface's `href` getter uses a different percent-encode set than `URLSearchParams` (e.g., space encoding as `+` vs `%20`, tilde handling). This can lead to mismatched strings when comparing query components directly with path components.
*   **Legacy Encoding Pitfalls:** The `application/x-www-form-urlencoded` format is acknowledged as a legacy aberration. Interoperability requirements often override good design, creating subtle pitfalls in nested encodings or charset handling.

## group-02

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

