---
title: Fetch Security Policies
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Fetch Security Policies

The Fetch Standard defines a unified architecture for fetching resources, prioritizing **Security First** across all web APIs. This section outlines the core security constraints and mechanisms implemented within the specification.

## Core Principles

*   **Unified Architecture**: Replaces fragmented implementations with a single algorithm governing all web platform fetching APIs, ensuring consistent security behavior across different URL schemes like HTTP(S), data, and blob.
*   **Security by Default**: The standard prioritizes security over legacy compatibility, blocking mixed content and ignoring dangerous header parameters.

## Connection Management and Partitioning

To prevent data leakage between sites, the specification enforces strict connection management:

*   **Network Partition Key**: Connections are strictly partitioned by a "Network Partition Key" (site origin + credentials). This ensures that cookies and sensitive data do not leak between different origins or security contexts.
*   **Scheme Resolution**: The fetching algorithms handle scheme resolution to ensure requests adhere to appropriate security contexts.

## Security Protocols

The standard implements specific protocols to mitigate common web vulnerabilities:

*   **CORS Safety & Preflight**: Cross-Origin Resource Sharing requires explicit opt-in via headers like `Access-Control-Allow-Origin`. This often involves a two-step preflight `OPTIONS` request. Headers are classified as "CORS-safelisted" or forbidden to control header exposure.
    *   See [[cors-preflight-cache]] for details on caching preflight requests.
*   **Mixed Content Blocking**: Requests involving mixed content (insecure resources loaded over secure contexts) are blocked by default.
*   **COEP Checks**: The standard enforces Cross-Origin-Embedder-Policy checks to restrict resource embedding.

## Response States and Header Exposure

Responses are categorized into three states which dictate header exposure, body readability, and redirect handling:

1.  **Basic**: Standard responses where all headers and the body are accessible.
2.  **CORS**: Responses subject to Cross-Origin Resource Sharing constraints.
3.  **Opaque**: Responses where headers and the body are not readable by the client (e.g., opaque redirects).

## Request and Response Handling

The Fetch API defines JavaScript interfaces (`Headers`, `Request`, `Response`) that abstract low-level operations while managing internal states:

*   **Buffering Limits**: Request bodies are buffered to 64 KiB when the source is a stream. This handles potential resends due to timeouts, with deferred fetching quotas managing memory usage hierarchically (640 KiB top-level, 64 KiB concurrent per origin).
    *   See [[deferred-fetching-quota]] for details on quota management.
*   **Body Streams**: Request and response bodies are streams (`ReadableStream`), supporting incremental reading and cloning via `tee()`. This allows content decoding without loading the entire payload into memory immediately.
    *   See [[fetch-body-streams]] for implementation details.

## Infrastructure and Lifecycle

The specification details the complete lifecycle of an HTTP request and response, from infrastructure setup to final body consumption:

*   **Infrastructure**: Defines URL parsing, method normalization (converting verbs to uppercase), and header handling rules.
*   **Http Fetch Lifecycle**: The process moves into connection management and then executes the core fetching algorithms, including redirect logic (up to 20 redirects) and security protocol enforcement.
    *   See [[http-fetch-lifecycle]] for a detailed breakdown of the stages.

## Request-Response Objects

The `Request` and `Response` objects expose attributes like `status`, `headers`, and `body`. These classes manage internal states such as timing info and abort signals, ensuring that security constraints are respected throughout the object's lifetime.
*   See [[request-response-objects]] for attribute definitions.
*   See [[log]] for debugging and monitoring request/response events.
