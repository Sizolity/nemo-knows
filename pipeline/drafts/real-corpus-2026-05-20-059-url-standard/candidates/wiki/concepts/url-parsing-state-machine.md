---
title: Url Parsing State Machine
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

# Url Parsing State Machine

The **URL Parsing State Machine** is a state-based algorithm defined by the URL Standard for converting scalar strings into structured URL records. This machine handles the lifecycle of network identifiers, managing infrastructure requirements and error handling to ensure robust parsing across modern implementations.

## States and Transitions

The parser operates through distinct states, beginning with specific entry conditions:
- **Scheme Start**: Requires an ASCII alpha character. If present, the machine transitions to the scheme state; otherwise, it moves to the `no scheme` state.
- **No Scheme**: Handles inputs where a scheme is absent. Unless the fragment starts the input and the base URL has an opaque path, this state rejects the parsing attempt.
- **Path State**: Processes path segments, specifically handling `.`, `..`, and rejecting backslashes in non-file contexts.

## Parsing Logic

The state machine defines valid schemes and applies special handling for specific URL types:
- **File URLs**: Must start with `//`.
- **Path Components**: The logic processes the path segment by segment, managing normalization rules for relative paths.

## Error Handling

A critical design principle of this state machine is the distinction between error types:
- **Validation Errors**: Signal a mismatch (e.g., invalid character in a component) but do not necessarily stop the parser. Reporting these errors is encouraged for conformance checkers.
- **Parser Termination**: Must be explicit, typically occurring when a state explicitly rejects the input or encounters fatal encoding issues.

## Encoding and Serialization

The machine ensures integrity during the roundtrip of data:
- **Percent-Encoding**: Sequences generally decode to valid UTF-8 without BOM for hosts, where failure is considered fatal. Different components utilize specific encoding sets (e.g., `application/x-www-form-urlencoded`).
- **Serialization**: Converts the internal URL record back into ASCII strings, adhering to strict rules about preserving casing or reversing octets for IP addresses.

## Security Boundaries

The parsing logic enforces security boundaries against visual spoofing and data leakage:
- **Untrusted Operations**: URL passing is treated as an untrusted operation.
- **Rendering Rules**: Strictly prohibit displaying credentials and emphasize showing registrable domains over full hosts.
- **Invisible Code Points**: The machine mandates handling of invisible code points to prevent homograph attacks.
