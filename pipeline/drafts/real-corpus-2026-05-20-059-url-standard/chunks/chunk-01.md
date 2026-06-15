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
