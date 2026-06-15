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
