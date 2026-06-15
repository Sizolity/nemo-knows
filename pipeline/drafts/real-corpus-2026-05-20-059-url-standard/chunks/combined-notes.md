## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

# Chunk Context

This chunk covers the initial sections of the **URL Standard** (Living Standard, last updated April 14, 2026). It defines the goals of the standard, which include aligning with modern implementations over obsolete RFCs, standardizing terminology (preferring "URL" over URI/IRI), and defining a robust API. The text details **Infrastructure** requirements, including writing rules, parser mechanics, percent-encoding definitions, and security considerations regarding spoofing and data leakage.

# Local Summary

The URL Standard aims to make URLs fully interoperable by updating legacy RFCs (3986, 3987) and defining a unified parsing model similar to HTML parsing. It establishes specific rules for infrastructure components: writing validation errors without terminating parsers immediately, handling percent-encoded bytes carefully to avoid security issues during UTF-8 decoding, and defining sets of characters that require encoding in different URL components (path, query, fragment, etc.). Security warnings emphasize the risks of visual spoofing (e.g., Cyrillic vs. Latin characters) and invisible code points like Left-to-Right Embedding.

# Key Claims

*   **Interoperability over Legacy:** The standard aligns with contemporary implementations and obsoletes RFC 3986 and RFC 3987, specifically addressing concepts like spaces and "illegal" code points that were previously problematic.
*   **Terminology Shift:** The term "URL" is standardized; URI and IRI are considered confusing distinctions in practice as they often use a single algorithm.
*   **Parser Resilience:** A validation error does not imply parser termination; termination must be explicitly stated (e.g., via a return statement). Reporting errors is encouraged for conformance checkers.
*   **Percent-Encoding Integrity:** Sequences of percent-encoded bytes should generally decode to valid UTF-8 without BOM or fail, particularly for hosts where failure is fatal.
*   **Security Risks:** Rendering and passing URLs requires care against "spoofing" attacks (visual similarity between characters) and ensuring that party B does not blindly trust data passed from untrusted party A.

# Entities And Concepts

*   **URL Standard:** The specification defining URLs, domains, IP addresses, and the `application/x-www-form-urlencoded` format.
*   **Infrastructure:** Dependencies on Infra, Encoding, File API, HTML, Unicode IDNA Compatibility Processing (UTS46), and Web IDL.
*   **Validation Errors:** Signals indicating mismatches between input and valid input; distinct from parser termination.
*   **Percent-Encoding:** The process of representing bytes using `U+0025 (%)` followed by two ASCII hex digits.
*   **Percent-Encode Sets:** Specific sets of code points requiring encoding in various URL components (e.g., C0 control, fragment, query, path, userinfo).
    *   *Component percent-encode set:* Used for embedding data in paths/queries.
    *   *Application/x-www-form-urlencoded percent-encode set:* Used for form data.
*   **Spoofing:** Attacks relying on visual similarities between characters (e.g., `1` vs `l`, `0` vs `O`) or invisible code points (e.g., Left-to-Right Embedding).

# Procedures And API Details

*   **Percent-Decoding Bytes:**
    1.  Initialize output as an empty byte sequence.
    2.  Iterate through input bytes. If a byte is not `%` (0x25), append it to output.
    3.  If a byte is `%`, check the next two bytes. If they are not valid hex digits (0-9, A-F, a-f), treat as literal `%`.
    4.  If valid, interpret the two following bytes as a hexadecimal number and append that value to output. Skip those two bytes.
*   **Percent-Encoding Bytes:** Return a string consisting of `U+0025 (%)` followed by two ASCII upper hex digits representing the byte value.
*   **UTF-8 Percent-Encoding:** Encode the scalar value string into UTF-8, then apply the percent-decoding/encoding logic using the specified set.
*   **Parser Pointer Logic:** A pointer is an integer referencing a code point in a string (or -1 for nowhere). `c` references the code point at the pointer; `remaining` references the substring from `pointer + 1` to the end.

# Nuance Or Contradictions

*   **Error Handling vs. Termination:** The standard explicitly states that validation errors do not mean the parser stops. This contrasts with traditional error-handling models where an invalid character might immediately abort processing.
*   **Percent-Encoding Roundtripping:** Only specific percent-encode sets (`component` and `application/x-www-form-urlencoded`) result in "roundtripable data" (encoding `%` to `%25`). Other sets used by the URL parser leave `%` untouched, requiring it to be encoded first if present.
*   **Security Trust Model:** The standard emphasizes that party B should *never* trust party A when passing URLs, as untrusted sources can inject malicious data or spoofed URLs.

# Candidate Wiki Hints

*   **Page: URL Standard Goals** – Summarizing the shift from RFCs to modern implementations and terminology changes.
*   **Page: Percent-Encoding Rules** – Detailing the specific sets of characters requiring encoding in different URL components.
*   **Page: URL Security Considerations** – Covering spoofing, invisible code points, and trust models when passing URLs between parties.

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

# Chunk Context

**Heading Path:** `3. Hosts (domains and IP addresses)`
**Line Range:** 595–989
**Coverage:** Sections 3.1 through 3.5, covering host representations, miscellaneous rules (public suffixes), IDNA processing, writing formats, and detailed parsing logic for IPv4, IPv6, and opaque hosts.

# Local Summary

This chunk defines the lifecycle of a "host" in the URL standard: how it is represented as a domain, IP address, or opaque identifier; how it is serialized; and how it is parsed. It details strict validation rules regarding forbidden code points (like control characters or `<>?`), distinguishes between domains and IP addresses, and explains the relationship between public suffixes and registrable domains using the Public Suffix List (PSL). The text also provides algorithmic steps for handling International Domain Names (IDNA) via Unicode ToASCII/ToUnicode conversions and outlines the exact parsing logic for IPv4 and IPv6 addresses embedded within host strings.

# Key Claims

- A "host" is an in-memory representation that can be serialized back to a valid ASCII string, with round-trip consistency depending on the `isOpaque` flag.
- A domain is strictly defined as a non-empty ASCII string identifying a realm; `example.com` and `example.com.` are treated as distinct entities.
- Forbidden code points include control characters (NULL, TAB, LF), whitespace, and specific punctuation marks like `#`, `/`, `:`, `?`, `@`, `[`, `\`, `]`, `^`, `|`.
- The concept of "public suffix" and "registrable domain" is provided for informational purposes but should not be relied upon for hard security boundaries due to potential client divergence.
- IDNA processing in this context uses Unicode IDNA Compatibility Processing (UTS46) rather than IDNA2008, allowing characters like `☕` to be converted to punycode (`xn--53h`).
- IPv6 addresses in host strings must be enclosed in square brackets `[...]`; unbracketed IPv6 strings are invalid.
- IPv4 addresses ending in a number (e.g., `example.255`) are treated as opaque hosts if they do not conform to valid decimal range checks or structure, while `0x` prefixes are handled specifically in parsing.

# Entities And Concepts

- **Host:** The core data type representing a network address (domain, IP, or opaque string).
- **Domain:** A non-empty ASCII string identifying a realm (e.g., `example.com`).
- **Opaque Host:** A non-empty ASCII string used as an identifier where a network address is not required (e.g., git repository URLs).
- **Public Suffix List (PSL):** The external list used to determine the public suffix of a host.
- **Registrable Domain:** The domain label preceding the public suffix, combined with the public suffix itself.
- **IDNA Compatibility Processing:** The algorithm used to convert internationalized domain names to ASCII (punycode).
- **Forbidden Code Points:** Specific Unicode characters disallowed in host strings (e.g., `U+0000`, `U+0020`, `U+003A`).
- **IPv4/IPv6 Address:** Network addresses represented as 32-bit or 128-bit integers, respectively.

# Procedures And API Details

**Host Serialization vs. Parsing Roundtrip:**
The behavior depends on the `isOpaque` argument passed to the host parser:
- `isOpaque = false`: Converts domain-like inputs (e.g., `EXAMPLE.COM`) to lowercase domains (`example.com`). Converts IPv4 hex (`0xffffffff`) to decimal (`255.255.255.255`).
- `isOpaque = true`: Preserves casing and encoding exactly (e.g., `EXAMPLE.COM` remains `EXAMPLE.COM`; `faß.example` becomes `fa%C3%9F.example`).

**Public Suffix Algorithm Steps:**
1. Return null if host is not a domain.
2. Determine trailing dot status.
3. Run Public Suffix List algorithm on the domain.
4. Assert result ends with the trailing dot (if present).
5. Return the public suffix.

**IDNA to ASCII Steps (`domain-to-ASCII`):**
1. Run Unicode ToASCII with specific flags (`CheckBidi=true`, `UseSTD3ASCIIRules=beStrict`, etc.).
2. If failure, return error.
3. If `beStrict` is false, check for forbidden code points in the result; if present, return error.
4. Assert result is non-empty and contains no forbidden code points.
5. Return result.

**IPv4 Parsing Logic:**
- Splits input on `.`.
- Validates each part is a decimal number between 0–255.
- Handles leading zeros (e.g., `09`) by flagging them as non-decimal errors if they fail standard parsing, though the text notes `09` might be caught later as an opaque host or specific error depending on context.
- Rejects parts > 255 or incorrect part counts.

**IPv6 Parsing Logic:**
- Handles bracketed notation `[...]`.
- Supports compression (`::`) but validates it appears only once.
- Allows mixed IPv4/IPv6 notation (e.g., `::ffff:192.0.2.1`) by parsing the final segment as an IPv4 address if dots are encountered after 6 pieces.
- Validates piece count and hex digit ranges.

# Nuance Or Contradictions

- **Security Boundaries:** The text explicitly warns that "public suffix" and "registrable domain" concepts cannot provide hard security boundaries because PSL lists may diverge between clients. Specifications should prefer the "origin" concept for security instead.
- **Case Sensitivity:** While domains are typically case-insensitive (e.g., `EXAMPLE.COM` vs `example.com`), the serialization with `isOpaque=true` preserves the original casing, creating distinct string representations that map back to different hosts depending on parsing mode.
- **Trailing Dots:** The distinction between `example.com` and `example.com.` is maintained strictly; they are not equivalent despite often being treated similarly in DNS lookups.
- **Hex Prefixes:** Inputs like `0x` or `09` behave differently based on the parser's strictness. `0x` is preserved as an opaque host, while `09` triggers a failure in standard parsing but might be accepted as opaque if `isOpaque` is true.

# Candidate Wiki Hints

- **Host Representation:** A page explaining the four types of hosts (Domain, IP, Opaque, Empty) and their use cases in URLs.
- **Public Suffix List Logic:** A guide to determining registrable domains versus public suffixes, including the algorithmic steps and security warnings.
- **IDNA Processing:** An explanation of how international domain names are converted to ASCII using UTS46, including compatibility processing examples.
- **IPv6 in Hosts:** Specific rules for bracketed IPv6 addresses, zone IDs (omitted here), and mixed IPv4/IPv6 literals.
- **Parsing Algorithms:** Reference documentation for the internal `domain-to-ASCII` and IPv4/IPv6 parsing logic used by URL parsers.

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---
Chunk Context
This chunk covers the serialization of hosts (IPv4, IPv6, domains) and host equivalence. It then details the high-level relationship between URL parsers, serializers, and valid strings, providing a table of input/output behaviors for various schemes and paths. The text defines URL representation components (scheme, username, password, host, port, path, query, fragment), allowed scheme/host combinations, path segment definitions, special schemes and their default ports, Windows drive letter handling, and the specific rules for writing valid URL strings including absolute/relative URLs and fragments. It concludes with character set definitions for URL units, percent-encoding behavior in legacy encodings, and a note on the inability to express credentials in valid strings.

Local Summary
The chunk provides algorithmic specifications for converting host records (IPv4, IPv6, domains) into ASCII strings, including logic for IPv6 address compression. It defines host equivalence rules for certificate comparison. A significant portion is dedicated to URL standards, explaining how parsers map input strings to internal URL records and how serializers convert those records back to strings. It includes a comprehensive table of valid/invalid inputs versus outputs for various URL formats. The text also details the structural components of a URL record, specific path segment types (single-dot, double-dot), special scheme handling, Windows drive letter normalization, and the grammar for constructing absolute and relative URL strings. Finally, it addresses character encoding constraints for URLs and legacy document encodings.

Key Claims
- Host serialization returns an ASCII string based on whether the host is IPv4, IPv6 (wrapped in brackets), or a domain/opaque host.
- IPv4 serialization involves reversing the octets by repeatedly prepending `n % 256` and dividing `n` by 256.
- IPv6 serialization requires finding the longest run of zero pieces to compress, adhering to RFC5952 recommendations.
- Host equivalence checks for certificate comparison ignore trailing dots but do not enforce DNS length constraints present in URLs.
- A URL record is a struct containing scheme, username, password, host, port, path (segment or list), query, and fragment.
- Valid URL strings are either relative-URL-with-fragment or absolute-URL-with-fragment.
- Credentials (username/password) cannot be expressed within a valid URL string according to the provided text.
- Legacy document encodings (non-UTF-8) cause issues with percent-encoding in query strings; UTF-8 is recommended to solve this.

Entities And Concepts
- Host: IPv4 address, IPv6 address, domain, opaque host, empty host.
- URL Record: In-memory representation of a universal identifier.
- Special Scheme: Schemes like "ftp", "file", "http", "https", "ws", "wss".
- Path Segment: Single-dot ("."), double-dot (".."), or other ASCII strings.
- Windows Drive Letter: Two code points (e.g., "c:").
- Base URL: Used for parsing relative URLs.
- Blob URL Entry: Supports caching for blob URLs.

Procedures And API Details
- **IPv4 Serializer**:
  1. Initialize `output` as empty string.
  2. Set `n` to the address value.
  3. Loop `i` from 1 to 4: prepend `n % 256` (serialized) to output; if `i < 4`, prepend "."; set `n = floor(n / 256)`.
  4. Return `output`.
- **IPv6 Serializer**:
  1. Initialize `output` as empty string.
  2. Find `compress` index (longest run of zeros).
  3. Iterate through address pieces: skip leading zeros if compressing; append "::" for the first compressed group or ":" otherwise; append piece as shortest lowercase hex; append ":" unless it is the last piece.
  4. Return `output`.
- **Find Compressed Piece Index**:
  1. Track `longestIndex`, `longestSize`, `foundIndex`, `foundSize`.
  2. Iterate pieces: if piece != 0, update longest if current run is longer; reset found tracking. If piece == 0, increment found run size.
  3. Return index of the longest run of zeros.

Nuance Or Contradictions
- The text states "There is no way to express a username or password of a URL record within a valid URL string," yet earlier examples show `https://user:password@example.org/` failing validation while `https://user:password@example.org/` (same content) passes in the output column, suggesting a potential inconsistency or specific constraint on how credentials are handled during serialization versus parsing in this context.
- Host equivalence ignores trailing dots for certificate comparison, whereas URLs generally do not enforce DNS length constraints, creating a divergence between URL validation and certificate matching logic.

Candidate Wiki Hints
- IPv6 Address Text Representation (RFC5952)
- Host Serialization Algorithms
- URL Record Structure and Components
- Special Schemes and Default Ports
- Windows Drive Letter Normalization in URLs
- Legacy Encoding Issues with URL Query Strings

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---
## Chunk Context
**Location:** Section 4.4 (URL parsing) within the URL Standard.
**Scope:** Defines the state machine algorithm for parsing a scalar input string into a URL record, handling schemes, hosts, paths, and special cases like `file` URLs and Windows drive letters.

## Local Summary
The basic URL parser accepts an input string, an optional base URL, and an optional encoding (defaulting to UTF-8). It initializes state variables such as `buffer`, `atSignSeen`, and `insideBrackets`. The core logic is a state machine that transitions based on the current code point `c` and the current `state` (e.g., `scheme start state`, `relative state`, `path state`).

The parser handles specific characters like colons (`:`) to detect schemes, forward slashes (`/`) for paths or authority separators, at-signs (`@`) for credentials, and brackets (`[]`) for IPv6 literals. It validates inputs against rules regarding C0 controls, tabs, newlines, and Windows drive letter quirks. If validation fails (e.g., missing scheme in non-relative URLs, invalid port ranges), it returns a failure or throws a specific validation error.

## Key Claims
- **Input Handling:** The parser removes leading/trailing C0 control or space characters but throws an error if ASCII tabs or newlines are present.
- **Scheme Detection:** A scheme starts with an ASCII alpha character and continues with alphanumeric characters, `+`, `-`, or `.`. It must be followed by a colon (`:`).
- **Base URL Interaction:** If the input lacks a scheme and a base URL is provided, the parser attempts to inherit the base's scheme and path components unless the base has an opaque path (which causes failure for non-hash characters).
- **File URLs:** `file` URLs require a leading `//`. They support Windows drive letter quirks where a platform-independent representation replaces the second code point of a drive letter with a colon.
- **Host Parsing:** Hosts are parsed after the scheme and authority slashes. If brackets are detected, IPv6 literals are handled; otherwise, host parsing assumes an IP address or hostname.
- **Path Segments:** The path state handles `.`, `..` segments by shortening the path. It also handles Windows drive letter quirks within paths for `file` URLs.

## Entities And Concepts
- **URL Record:** The structured object resulting from parsing, containing components like scheme, host, port, path, query, and fragment.
- **State Machine:** The procedural flow (`scheme start state`, `no scheme state`, `relative state`, etc.) that dictates how the parser consumes input characters.
- **Validation Errors:** Specific error codes returned or thrown (e.g., `missing-scheme-non-relative-URL`, `invalid-reverse-solidus`, `host-missing`).
- **Base URL:** An optional reference URL used to resolve relative inputs when the input lacks a scheme.
- **Opaque Path:** A path component that does not contain authority information (scheme, host, port), often seen in non-hierarchical schemes or specific file paths.

## Procedures And API Details
1.  **Initialization:**
    -   Create a new URL record.
    -   Remove leading/trailing C0 controls/spaces; reject tabs/newlines.
    -   Initialize `buffer`, `atSignSeen`, `insideBrackets`, and `pointer`.
2.  **Scheme Start State:**
    -   If `c` is ASCII alpha: Append lowercased `c` to `buffer`, move to `scheme state`.
    -   If no scheme override: Move to `no scheme state`.
3.  **Scheme State:**
    -   Accumulate alphanumeric/`+`/-/. characters.
    -   On `:` (colon): Finalize scheme, check for special scheme rules (e.g., `file`), and transition to `file state`, `special authority slashes state`, or `relative slash state`.
4.  **No Scheme State:**
    -   If base URL has an opaque path: Reject unless `c` is `#` (fragment).
    -   Otherwise, inherit scheme from base if applicable; move to `relative state` or `file state`.
5.  **Relative/Path States:**
    -   Handle `/` and `\` characters to determine if moving to authority or path segments.
    -   Process credentials (`@`) and host parsing logic.
    -   Manage port parsing (digits only, max 16-bit unsigned).
6.  **File State Specifics:**
    -   Enforce `//` prefix for special scheme `file`.
    -   Handle Windows drive letter quirks (platform-independent representation).
7.  **Path State Logic:**
    -   Process `.` and `..` segments.
    -   Accumulate path segments in `buffer`, handling percent-encoding.
    -   On `?` or `#`: Finalize path, move to query/fragment states.

## Nuance Or Contradictions
- **Windows Drive Letter Quirk:** The standard explicitly handles Windows drive letters (e.g., `C:`) in a platform-independent manner. For example, if a buffer contains a Windows drive letter and the scheme is `file`, the second code point is replaced with `:` to normalize it. This applies to both path start states and file slash states.
- **Reverse Solidus:** Backslashes (`\`) are generally invalid within authority or path segments for non-special schemes (e.g., in `relative state`), triggering an `invalid-reverse-solidus` error, though they are processed differently in `file` URLs or when paired with specific characters.
- **Base URL Opaque Paths:** If a base URL has an opaque path and the input starts with something other than `#`, parsing fails (`missing-scheme-non-relative-URL`). This is a strict constraint to prevent ambiguity in resolving relative URLs without schemes.
- **Encoding Override:** The encoding argument is legacy for HTML. For non-special schemes like `ws` or `wss`, if the provided encoding is not UTF-8, it is forcibly set to UTF-8 in the query state.

## Candidate Wiki Hints
-   **URL Parsing Algorithm**: A detailed breakdown of the URL parser's state machine and transition rules.
-   **Base URL Resolution**: How relative URLs are resolved against a base URL, including opaque path handling.
-   **File URL Syntax**: Specifics regarding `file://` URLs, including the requirement for `//` and Windows drive letter quirks.
-   **URL Validation Errors**: A list of specific validation errors (e.g., `invalid-reverse-solidus`, `host-missing`) and when they are triggered.

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---
## Chunk Context
Lines 1849–2133 of `raw/web/corpus-2026-05-18/059-url-standard.md`. Covers the EOF handling logic for URL parsing, URL serialization and equivalence rules, origin determination (including opaque origins for blob/file schemes), rendering guidelines for security display, and details on the `application/x-www-form-urlencoded` format plus its API.

## Local Summary
The chunk describes how a parser handles end-of-file (`EOF`) in various URL states, defines the steps to serialize a URL into an ASCII string (including handling of credentials, ports, paths, query strings, and fragments), explains equivalence checks via serialized comparison, outlines origin extraction for common schemes (blob, ftp, http, https, ws, wss, file), provides rendering advice to avoid spoofing (host simplification, hiding credentials, scheme replacement), and specifies parsing/serialization rules for `application/x-www-form-urlencoded` including percent-encoding hooks.

## Key Claims
- When `c` is the EOF code point, the parser must handle state transitions: special-query percent-encoding, fragment clearing, and validation of non-code-point characters or malformed percent-sequences (`%` followed by non-hex).
- URL serialization returns an ASCII string; it concatenates scheme, host (with credentials if present), optional port, path, query, and optionally fragment.
- Equivalence is determined by comparing serialized forms with fragments excluded by default.
- Origin extraction depends on scheme: tuple origins for `http`, `https`, `ws`, `wss`, `ftp`; blob origins come from the environment or are opaque; `file` origins are left as an exercise/often opaque.
- Rendering rules prioritize showing the registrable domain, hiding credentials, and handling IDNs/confusable characters to mitigate homograph attacks.
- `application/x-www-form-urlencoded` is a legacy, encoding-twisted format widely used for HTML forms; parsing splits on `&`, replaces `+` with space, percent-decodes UTF-8 without BOM, and serializing performs the inverse with specific percent-encode sets.

## Entities And Concepts
- EOF code point (end of input handling)
- URL state machine (special-query, fragment states)
- Percent-encoding sets (query, fragment, userinfo, application/x-www-form-urlencoded)
- URL serialization algorithm
- URL equivalence (serialized comparison)
- Origin (tuple vs opaque; scheme-specific rules)
- Rendering for security decisions
- Homograph spoofing / IDN confusion
- `application/x-www-form-urlencoded` (parsing, serializing, hooks)

## Procedures And API Details
**EOF Handling Steps**
When input ends:
1. If in special-query state, percent-encode remaining buffer with the query set and append to `url.query`.
2. If `c` is `#`, clear fragment and move to fragment state.
3. Otherwise (non-EOF): validate code point; if `%` appears without two hex digits following, raise validation error; else append to buffer.

**Fragment State Steps**
For each input character until EOF:
1. Validate code point or reject `%` without hex digits.
2. Percent-encode with fragment set and append to `url.fragment`.

**Username/Password Setting**
- Username: UTF-8 percent-encode using the userinfo set.
- Password: UTF-8 percent-encode using the userinfo set.

**URL Serialization Steps**
Input: `url`, optional `excludeFragment` (default false).
1. Output = scheme + `:`.
2. If host non-null:
   - Append `//`.
   - If credentials present: append username, then `:` + password if non-empty, then `@`.
3. Append serialized host.
4. If port non-null: append `:` + serialized port.
5. If path not opaque and has multiple segments starting with empty string, append `/.` to preserve path structure (prevent collapse).
6. Append serialized path.
7. If query non-null: append `?` + query.
8. If `excludeFragment` is false and fragment non-null: append `#` + fragment.
9. Return output.

**URL Path Serialization Steps**
1. If opaque path, return path as-is.
2. Else, iterate path segments, prefix each with `/`, join, return.

**Equivalence Check**
- Serialize both URLs with fragments excluded by default; compare resulting strings.

**Origin Determination**
- `blob`: environment origin from blob URL entry if available; else parse path as URL and derive or return opaque origin.
- `ftp`, `http`, `https`, `ws`, `wss`: tuple origin `(scheme, host, port, null)`.
- `file`: exercise to reader; often opaque.
- Others: opaque origin.

**Rendering Guidelines**
- Show only host (or registrable domain) where security distinction matters.
- Do not render username/password.
- May omit scheme if display surface is single-scheme; may replace with human-readable string or icon.
- Elide from lowest-level domain label when full host cannot fit.
- Render IDNs via Unicode normalization; detect confusable characters; treat bidirectional text as left-to-right for rendering.

**`application/x-www-form-urlencoded` Parsing Steps**
1. Split input on `&`.
2. For each segment:
   - If empty, skip.
   - If contains `=`, split into name/value (empty if at start/end).
   - Else treat whole as name, value empty.
3. Replace `+` with space in both name and value.
4. Percent-decode then UTF-8 decode without BOM for name and value.
5. Return list of `(nameString, valueString)` tuples.

**`application/x-www-form-urlencoded` Serialization Steps**
1. Determine output encoding (default UTF-8).
2. For each tuple:
   - Percent-encode name and value using the form-urlencoded set.
   - Join with `=`; separate tuples with `&`.
3. Return ASCII string.

**Hooks**
- String parser: UTF-8 encode input, then apply form-urlencoded parsing.

## Nuance Or Contradictions
- The `file` origin is explicitly left as an exercise, creating potential inconsistency across implementations unless defined elsewhere.
- Blob origins can be opaque if path parsing fails or scheme mismatches, leading to same-origin behavior differences.
- Rendering rules allow browser discretion (e.g., omitting `www`, replacing scheme with icon), which may cause serialized vs displayed URLs to differ significantly for security decisions.
- The form-urlencoded format is acknowledged as a legacy aberration with encoding twists; interoperability requirements override good design, leading to subtle pitfalls in nested encodings or charset handling.

## Candidate Wiki Hints
- **URL Serialization and Equivalence**: A canonical algorithm for converting URL objects to ASCII strings and checking equality by serialized comparison (fragments excluded).
- **Origin Extraction Logic**: Scheme-specific rules for determining tuple vs opaque origins, including blob environment handling.
- **Secure URL Rendering Guidelines**: Best practices for displaying URLs in address bars to prevent spoofing (hide credentials, show registrable domain, handle IDNs/confusables).
- **application/x-www-form-urlencoded Format**: Detailed parsing/serialization steps, percent-encode sets, and hook behavior; useful for form handling libraries.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---
Chunk Context
The chunk covers the **URL** and **URLSearchParams** interfaces within the URL Standard. It details the API for constructing URLs, parsing strings (absolute vs. relative), accessing components (protocol, host, pathname, etc.), and manipulating query parameters. It also explains encoding differences between the `URL` object's serialization and `URLSearchParams`.

Local Summary
The **URL** interface allows creating objects from absolute or relative URLs. The constructor accepts a URL string and an optional base URL. Static methods `parse()` and `canParse()` handle parsing without throwing exceptions (returning null/false on failure). Accessors like `href`, `origin`, and setters for components (`protocol`, `username`, `host`, etc.) are defined, with specific behaviors regarding opaque paths and port handling. The **URLSearchParams** interface manages the query component of a URL using a list of name-value tuples, supporting methods like `append()`, `delete()`, `get()`, `getAll()`, `set()`, and `sort()`. Stringification follows the application/x-www-form-urlencoded format, which differs from standard URL percent-encoding (e.g., encoding spaces as `+`).

Key Claims
- A `URL` object has an associated `URL` record and a `query` object (a `URLSearchParams` instance).
- The `new URL(url)` constructor throws a `TypeError` if the input is a relative URL without a base.
- The `href` getter returns the serialization of the internal `URL` record; the setter re-parses the string and clears the query list before repopulating it.
- Setting the `host` property does not reset the port if the given value lacks one, which differs from the behavior implied by the `host` getter returning a port-inclusive string.
- `URLSearchParams` uses application/x-www-form-urlencoded encoding, where U+0020 SPACE is encoded as U+002B (+).
- The `sort()` method on `URLSearchParams` sorts tuples based on name code unit values to potentially increase cache hits.

Entities And Concepts
- **URL**: An object representing a URL with components like scheme, host, port, pathname, query, and hash.
- **URLSearchParams**: An object representing the query component of a URL as a list of name-value tuples.
- **Opaque path**: A condition where certain setters (e.g., `host`, `pathname`) are ignored or restricted.
- **application/x-www-form-urlencoded**: The encoding format used by `URLSearchParams`.
- **Percent-encode sets**: Distinct sets for general URLs vs. query strings affecting how characters like spaces or tildes are encoded.

Procedures And API Details
- **Constructing a URL**:
  - `new URL(url, base)`: Throws `TypeError` on failure.
  - `URL.parse(url, base)`: Returns `null` on failure.
  - `URL.canParse(url, base)`: Returns `false` on failure.
- **Setting components**:
  - `protocol`, `username`, `password`, `port`: May return early if the URL cannot have these (e.g., file URLs).
  - `host`: Ignores port if not provided in the setter value.
  - `pathname`: Clears the path before reparsing; ignores if opaque.
  - `search`: Removes leading `?` from input before parsing.
  - `hash`: Removes leading `#` from input before parsing.
- **URLSearchParams operations**:
  - `append(name, value)`: Adds a tuple to the list.
  - `set(name, value)`: Replaces existing tuples with that name; appends if none exist.
  - `delete(name, value)`: Removes specific or all tuples matching the name.
  - `sort()`: Sorts the internal list by name code units and updates the object.

Nuance Or Contradictions
- **Encoding Discrepancy**: The `URL` interface's serialization (used by `href`) uses a different percent-encode set than `URLSearchParams`. Specifically, `URLSearchParams` encodes U+0020 SPACE as U+002B (+), whereas standard URL encoding often uses `%20`. Additionally, `~` is encoded differently depending on the context (query vs. path).
- **Host Setter Surprise**: The `host` getter includes the port in its return value, but the `host` setter does not update the internal port if the provided string lacks a port number. This can lead to unexpected state where the visual host changes but the internal port remains static.

Candidate Wiki Hints
- URL vs. URLSearchParams encoding differences
- Handling relative URLs with the URL constructor
- Sorting URLSearchParams for cache optimization
- Understanding opaque paths in URL setters

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

## Chunk Context

This chunk covers section **6.3. URL APIs elsewhere** and the subsequent sections of the URL Standard, including acknowledgments, intellectual property rights, an index of defined terms, normative references (such as RFC 3986), informative references, the Web IDL (IDL) interface definitions for `URL` and `URLSearchParams`, browser compatibility data from MDN, and a brief note on the `canParse` static method.

## Local Summary

The document section advises that standards exposing URLs should serialize them as strings rather than URL objects, which are reserved for manipulation. It specifies naming conventions (e.g., lowercase "url", compound names like "newURL"). The chunk lists extensive acknowledgments and copyright information under a Creative Commons Attribution 4.0 license (with BSD 3-Clause exceptions for source code). A comprehensive index of terms defined by the specification and referenced standards follows, leading into the formal IDL definitions for `URL` and `URLSearchParams`. Finally, it provides browser support tables for `URL`, `URL.canParse`, and related features across various engines.

## Key Claims

- Standards should expose URLs as strings (serialized internal URLs) rather than URL objects; URL objects are intended for manipulation.
- In IDL, the `USVString` type should be used for URL values.
- Higher-level notions require exposing values as immutable data structures.
- Naming conventions dictate that features named "URL" should use lowercase "url", while compound names prefer uppercase "URL" (e.g., "newURL").
- Examples of proper naming include `EventSource` and `HashChangeEvent` in HTML.
- The standard is written by Anne van Kesteren and licensed under Creative Commons Attribution 4.0 International, with source code portions using BSD 3-Clause License.

## Entities And Concepts

- **URL**: Interface for representing URLs.
- **USVString**: Unicode Scalar Value String type used in IDL.
- **URL object**: An object meant for URL manipulation.
- **EventSource**, **HashChangeEvent**: HTML interfaces cited as examples of proper naming.
- **canParse(url)**: Static method to check parseability (noted with MDN compatibility notes).
- **Living Standard**: The current version of the specification.
- **Creative Commons Attribution 4.0 International License**: Primary license for the standard.

## Procedures And API Details

### IDL Interface Definitions

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

### Browser Compatibility (MDN)

- **URL/URL**: Supported in Firefox 26+, Safari 14.1+, Chrome 19+, Opera?, Edge 79+, Edge (Legacy) 12+, Node.js 10.0.0+.
- **URL.canParse_static**: Supported in Firefox 115+, Safari 17+; not supported in Chrome, Opera, Edge (Legacy), or others.

## Nuance Or Contradictions

The compatibility table for `URL.canParse` shows a significant gap: it is supported in modern Firefox and Safari but explicitly noted as "None" for Chrome and other browsers listed. The text does not provide a reason for this discrepancy, which might indicate divergent implementation timelines or platform-specific constraints (e.g., Node.js support vs. browser support).

## Candidate Wiki Hints

- **URL API Naming Conventions**: Guidelines on naming URL-related features (lowercase "url" vs. uppercase in compounds).
- **URL Object Usage**: Distinction between exposing URLs as strings for standards and using objects for manipulation.
- **URL.canParse Compatibility**: A resource documenting the limited support of `canParse` across major browsers.

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

## Chunk Context
This section (Heading path: 6. API > 6.3. URL APIs elsewhere) details browser and runtime support for the `URL` interface and `URLSearchParams` object. It lists specific properties (e.g., `hostname`, `searchParams`) and methods (e.g., `append`, `delete`, `sort`) alongside their minimum supported versions for browsers like Firefox, Chrome, Safari, Edge, and Node.js.

## Local Summary
The chunk provides a comprehensive compatibility matrix for URL parsing and query string manipulation APIs. It distinguishes between the core `URL` interface (introduced in Node 10) and the more granular properties available in Node 20+. The `URLSearchParams` API is shown to be widely supported across modern browsers since around 2015-2017, with specific methods like `delete` and `sort` having slightly later adoption dates.

## Key Claims
- **Node.js Support**: The core `URL` class requires Node.js 10.0.0+. The `URLSearchParams` interface is supported from Node.js 10.0.0+, while specific properties like `URL.pathname` and methods on `URLSearchParams` generally require Node.js 7.5.0+ or higher (specifically 7.7.0+ for `sort`).
- **Browser Support**: All current engines support the `URL` interface. Firefox requires version 19+, Chrome 32+, and Safari 7+. The `URLSearchParams` object is supported from Firefox 29, Chrome 49, and Safari 10.1.
- **Method Granularity**: While basic access (e.g., `get`) is available in Firefox 29, methods like `delete` require Firefox 14 (Safari) or specific versions, and `sort` is not universally supported until newer versions (Firefox 54, Chrome 61).
- **Legacy Browsers**: Edge (Legacy) support is noted for specific versions (e.g., IE10+ for the base `URL`), while Opera support data is often marked with a question mark or specific version numbers like 79+.

## Entities And Concepts
- **URL Interface**: Represents a parsed URL string. Properties include `href`, `hostname`, `pathname`, `search`, etc.
- **URLSearchParams**: A convenient object for working with query strings. Methods include `append`, `delete`, `get`, `getAll`, `set`, `has`, `keys`, `values`, `forEach`, `toString`.
- **Node.js Versions**: Ranges from 7.0.0 to 20.0.0+ are referenced for API availability.
- **Browser Engines**: Firefox, Chrome, Safari, Edge (Legacy and Chromium), Opera, Samsung Internet.

## Procedures And API Details
- **URL Object Creation**: Supported in all current engines; Node.js 10.0.0+.
- **Property Access**:
    - `href`: Available since Firefox 19, Chrome 32, Safari 7, Node 10.
    - `hostname`: Available since Firefox 22, Chrome 32, Node 7.
    - `pathname`: Available since Firefox 22, Chrome 32, Node 7.
    - `searchParams`: Available since Firefox 29, Chrome 49, Safari 10.1, Node 7.5.
- **URLSearchParams Methods**:
    - `append`, `get`, `getAll`, `has`, `set`, `keys`, `values`, `toString`: Supported from Firefox 29/Chrome 49/Safari 10.1/Node 7.5.
    - `delete`: Supported from Firefox 29 (note: Safari requires 14 in some contexts or is implied by browser age), Chrome 49, Node 7.5.
    - `forEach`, `size`: Supported from Firefox 44/Chrome 49/Safari 10.1/Node 7.5.
    - `sort`: Supported from Firefox 54/Chrome 61/Safari 11/Node 7.7.

## Nuance Or Contradictions
- **Safari Versioning**: Some entries list Safari 10.1 for general `URLSearchParams`, while others specify Safari 17 for the `size` property, indicating that not all methods were available at the object's introduction.
- **Edge Discrepancies**: "Edge (Legacy)" often cites IE versions (e.g., IE10) alongside Edge numbers, suggesting a conflation of support timelines between Internet Explorer and early Chromium-based Edge in the source data.
- **Opera Ambiguity**: Many Opera entries are marked with "?", indicating uncertain or unverified support data for that browser compared to others.

## Candidate Wiki Hints
- **Page: URL_API_Compatibility** - Documenting the specific versions of browsers and runtimes where `URL` and `URLSearchParams` became available.
- **Page: URLSearchParams_Methods** - Detailing which methods (`append`, `delete`, `sort`) require polyfills for older environments.
- **Topic: Polyfill_Necessity** - Identifying that `URLSearchParams.sort()` might require a polyfill for browsers prior to Firefox 54 or Chrome 61.

