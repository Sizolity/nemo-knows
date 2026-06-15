---
title: URL Standard Source Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

## What It Is

The **URL Standard** (Living Standard) is a specification that defines the lifecycle of network identifiers, including infrastructure requirements, host representations, URL parsing and serialization, origin determination, secure rendering guidelines, and the `application/x-www-form-urlencoded` format. Last updated April 14, 2026, it aims to unify modern implementations by obsoleting legacy RFCs (3986, 3987) while establishing robust APIs (`URL`, `URLSearchParams`) and strict security boundaries against visual spoofing and data leakage.

## Summary

The document defines a complete lifecycle for network identifiers:
1.  **Infrastructure:** Establishes the parser model, percent-encoding rules, and error handling (validation errors vs. termination).
2.  **Hosts:** Defines how domains, IP addresses (IPv4/IPv6), and opaque identifiers are represented, parsed, serialized, and checked for equivalence. It details IDNA processing and public suffix logic.
3.  **URL Structure:** Breaks down the URL record into components (scheme, username, password, host, port, path, query, fragment). It defines valid schemes, special handling for `file` URLs, and Windows drive letter quirks.
4.  **Parsing & Serialization:** Provides state machine algorithms for parsing scalar strings into URL records and serializing them back to ASCII. Equivalence is determined by serialized comparison (excluding fragments).
5.  **Origins & Rendering:** Defines how to extract origins (tuple vs. opaque) and guidelines for secure display (hiding credentials, simplifying hosts, handling IDNs).
6.  **APIs:** Exposes the `URL` and `URLSearchParams` interfaces, detailing constructor behavior, property accessors/setters, and the specific encoding differences between query strings and general URL paths.

The standard prioritizes alignment with modern implementations over strict adherence to obsolete RFCs, explicitly addressing issues like spaces and "illegal" code points that caused fragmentation in the past. A critical distinction is made between a validation error (which signals a mismatch but does not stop the parser) and parser termination (which must be explicit). The standard treats URL passing as an untrusted operation; rendering rules strictly prohibit displaying credentials, emphasize showing registrable domains over full hosts, and mandate handling of invisible code points to prevent homograph attacks.

## Key Claims

*   **Interoperability over Legacy:** The standard aligns with contemporary implementations and obsoletes RFC 3986 and RFC 3987, specifically addressing concepts like spaces and "illegal" code points that were previously problematic.
*   **Terminology Shift:** The term "URL" is standardized; URI and IRI are considered confusing distinctions in practice as they often use a single algorithm.
*   **Parser Resilience:** A validation error does not imply parser termination; termination must be explicitly stated (e.g., via a return statement). Reporting errors is encouraged for conformance checkers.
*   **Percent-Encoding Integrity:** Sequences of percent-encoded bytes should generally decode to valid UTF-8 without BOM or fail, particularly for hosts where failure is fatal.
*   **Security Risks:** Rendering and passing URLs requires care against "spoofing" attacks (visual similarity between characters) and ensuring that party B does not blindly trust data passed from untrusted party A.
*   **Host Representation Consistency:** Hosts are case-insensitive domains or IP addresses by default, but serialization with `isOpaque=true` preserves original casing (e.g., `EXAMPLE.COM`). This creates distinct string representations that map to different hosts depending on parsing mode.
*   **Validation Errors vs. Termination:** A critical distinction is made between a validation error (which signals a mismatch but does not stop the parser) and parser termination (which must be explicit). This allows parsers to report errors without crashing or silently failing.
*   **Security First (Spoofing):** The standard treats URL passing as an untrusted operation. Rendering rules strictly prohibit displaying credentials, emphasize showing registrable domains over full hosts, and mandate handling of invisible code points (e.g., LRE, RLE) and visual confusables to prevent homograph attacks.
*   **Percent-Encoding Sets:** Different components use different encoding sets (`component`, `application/x-www-form-urlencoded`). The standard emphasizes that only specific sets result in roundtripable data; others require careful handling of the literal `%` character.
*   **Host Serialization Logic:** IPv4 serialization reverses octets by prepending `n % 256` and dividing `n` by 256 repeatedly. IPv6 serialization finds the longest run of zeros for compression (`::`).
*   **URL Parsing State Machine:** Scheme start requires ASCII alpha; moves to scheme state or `no scheme` state. No scheme rejects unless fragment starts input (if base URL has opaque path). File URLs must start with `//`. Path state processes `.`, `..` segments and rejects backslashes in non-file contexts.
*   **Origin Extraction Rules:** Tuple origins `(scheme, host, port, null)` for `http`, `https`, `ws`, `wss`, `ftp`. Blob origins come from environment origin if available; otherwise opaque. File origins are often opaque.
*   **`application/x-www-form-urlencoded`:** Parsing splits on `&`, replaces `+` with space, percent-decodes UTF-8 without BOM. Serialization percent-encodes using the form set (space becomes `%20` or `+` depending on context), joins with `=` and `&`.
*   **String Serialization Preference:** Standards should expose URLs as strings (serialized internal representations) rather than directly exposing URL objects, which are intended for internal manipulation.
*   **Licensing:** The specification is written by Anne van Kesteren, licensed under Creative Commons Attribution 4.0 International for the standard text, while source code portions use the BSD 3-Clause License.

## Suggested Links

- [URL Standard Goals](#)
- [Host Representation Types](#)
- [IDNA & Public Suffix Logic](#)
- [URL Parsing State Machine](#)
- [Secure URL Rendering](#)
- [application/x-www-form-urlencoded Format](#)
- [URL vs. URLSearchParams API](#)
- [URL_API_Naming_Conventions](#)
- [URL_API_Compatibility](#)
- [URLSearchParams_Methods_Polyfills](#)
- [Web_IDL_URL_Definition](#)

none
