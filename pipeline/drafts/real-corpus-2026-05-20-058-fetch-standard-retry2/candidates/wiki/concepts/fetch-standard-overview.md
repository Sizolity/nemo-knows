---
title: Fetch Standard Overview
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Fetch Standard Overview

The Fetch Standard is the WHATWG Living Standard that consolidates resource retrieval on the web platform into a single algorithm. Where browsers previously inherited different rules from XHR, image loading, script loading, and other APIs, this specification provides one definition that every fetching surface — including the JavaScript `fetch()` function, `<img>`, `<script>`, navigation, and Service Worker interception — is expected to share. The algorithm applies uniformly across HTTP and HTTPS as well as schemes such as `data:` and `blob:`, so behaviour around caching, redirects, and security checks remains consistent regardless of which entry point initiated the request.

## Specification Shape

Internally the standard layers its definitions from low to high. The Infrastructure section pins down URL parsing rules, normalises HTTP method tokens by uppercasing them, and defines how header names and values are represented and compared. Connection management sits above that: rather than reusing connections by host alone, sockets are keyed by a network partition that includes the requesting site's origin and credential state, so that two top-level documents on different sites cannot share a connection and inadvertently leak signals between them.

## Fetching Algorithms

The bulk of the document specifies the fetching algorithms themselves. These cover scheme dispatch (different schemes hand off to different sub-procedures), HTTP cache interaction with modes such as `default`, `no-store`, and `reload`, the chain of permitted redirects with a hard upper bound of twenty hops, and the cross-origin checks that gate whether a response is exposed to a caller. Mixed-content blocking, port blocking against a fixed list of disallowed network ports such as FTP and Telnet, and Cross-Origin-Embedder-Policy enforcement are all integrated into this main path rather than living as separate add-ons.

## Fetch API Surface

Finally the JavaScript-facing API layer defines `Headers`, `Request`, and `Response` as the visible interfaces. They wrap the low-level algorithm's intermediate state so that web content can observe values like the status code, response headers, and the body stream while the implementation keeps internal bookkeeping such as timing information and abort signals out of view. Closing sections of the specification handle ancillary topics: how `data:` URLs decode into a synthetic response, when bodies become eligible for garbage collection, and per-implementation compatibility tables documenting browser support.
