---
title: Host Representation Types
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

# Host Representation Types

Host representation defines how network identifiers are stored internally versus how they are serialized for display or transmission. The distinction is critical because internal representations (such as case-insensitive domains) differ from serialized strings which may preserve original casing.

## Internal vs. Serialized Representation

By default, hosts are treated as **case-insensitive** domains or IP addresses. However, serialization behavior depends on the context:

-   **Default Serialization:** When `isOpaque=false` (standard URL objects), the host is serialized in a canonical form where domain names are lowercased and IPv6 addresses follow standard compression rules.
-   **Opaque Serialization:** When `isOpaque=true`, the original casing of the host string is preserved during serialization (e.g., `EXAMPLE.COM`).

This creates distinct string representations that map to different hosts depending on the parsing mode used. For instance, a serialized opaque host might differ from its canonical internal representation.

## Domain and IP Address Logic

The standard supports three primary types of host representations:

1.  **Domains:** These are case-insensitive by nature. The standard handles IDNA (Internationalized Domain Names for Applications) processing to ensure valid UTF-8 output.
2.  **IP Addresses:**
    -   **IPv4:** Serialization involves reversing octets and applying specific formatting rules.
    -   **IPv6:** Serialization finds the longest run of zeros for compression using the `::` shorthand.
3.  **Opaque Identifiers:** Used when the host cannot be resolved to a domain or IP (e.g., file system paths). These often retain their original casing and structure.

## Validation and Parsing

The lifecycle of a host includes specific validation steps:

-   **Percent-Encoding Integrity:** Sequences of percent-encoded bytes should generally decode to valid UTF-8 without a Byte Order Mark (BOM). Failure is considered fatal for hosts.
-   **Parser Resilience:** A validation error signals a mismatch but does not necessarily stop the parser unless explicitly stated as termination. This allows parsers to report errors without crashing.

## Security Considerations

Security rules strictly prohibit displaying credentials in URLs. Furthermore, rendering guidelines emphasize showing registrable domains over full hosts to mitigate spoofing attacks. The standard mandates handling invisible code points (such as LRE and RLE) and visual confusables to prevent homograph attacks where visually similar characters are used to deceive users.
