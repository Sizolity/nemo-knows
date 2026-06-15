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
