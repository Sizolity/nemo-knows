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
