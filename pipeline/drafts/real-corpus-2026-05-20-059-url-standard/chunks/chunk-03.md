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
