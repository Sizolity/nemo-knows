---
title: Url Api Compatibility
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

# Url Api Compatibility

The **URL Standard** (Living Standard) establishes a unified specification for network identifiers, aiming to unify modern implementations by obsoleting legacy RFCs (3986, 3987). It defines robust APIs and strict security boundaries against visual spoofing and data leakage.

## Core Objectives

The standard prioritizes alignment with contemporary implementations over strict adherence to obsolete RFCs. Key goals include:

*   **Unified Lifecycle:** Defining the complete lifecycle for network identifiers, from infrastructure requirements to serialization.
*   **API Exposure:** Providing standardized interfaces such as `URL` and `URLSearchParams`, detailing constructor behavior and property accessors.
*   **Security Boundaries:** Implementing strict rules against visual spoofing, data leakage, and homograph attacks during rendering and passing of URLs.

## Parsing and Serialization

The standard introduces a parser model with specific resilience characteristics:

*   **Validation vs. Termination:** A validation error signals a mismatch but does not stop the parser; termination must be explicitly stated. This allows parsers to report errors without crashing or silently failing.
*   **State Machine Algorithms:** Provides algorithms for parsing scalar strings into URL records and serializing them back to ASCII.
    *   Scheme start requires ASCII alpha characters.
    *   Path state processes `.`, `..` segments and rejects backslashes in non-file contexts.
*   **Equivalence Determination:** Equivalence is determined by serialized comparison, excluding fragments.

## Host Representations

The standard details how domains, IP addresses (IPv4/IPv6), and opaque identifiers are represented:

*   **Case Sensitivity:** Hosts are case-insensitive domains or IP addresses by default. Serialization with `isOpaque=true` preserves original casing (e.g., `EXAMPLE.COM`).
*   **Serialization Logic:**
    *   IPv4 serialization reverses octets.
    *   IPv6 serialization finds the longest run of zeros for compression (`::`).
*   **IDNA Processing:** Includes logic for IDNA processing and public suffix determination.

## Secure Rendering Guidelines

Rendering rules strictly prohibit displaying credentials to prevent leakage. The standard emphasizes:

*   **Domain Display:** Showing registrable domains over full hosts.
*   **Invisible Code Points:** Mandating handling of invisible code points (e.g., LRE, RLE) and visual confusables to prevent homograph attacks.
*   **Untrusted Operations:** Treating URL passing as an untrusted operation between parties.

## Encoding Formats

Different components utilize different encoding sets:

*   **`application/x-www-form-urlencoded`:** Used for form data. Parsing splits on `&`, replaces `+` with space, and percent-decodes UTF-8 without BOM. Serialization uses the form set where space becomes `%20` or `+`.
*   **Percent-Encoding Integrity:** Sequences of percent-encoded bytes should generally decode to valid UTF-8 without BOM or fail, particularly for hosts where failure is fatal.

## Related Concepts

*   [[application-x-www-form-urlencoded]]
*   [[host-representation-types]]
*   [[query]]
*   [[secure-url-rendering]]
*   [[url-parsing-state-machine]]
