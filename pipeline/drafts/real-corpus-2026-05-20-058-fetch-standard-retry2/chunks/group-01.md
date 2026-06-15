---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This group of notes synthesizes the **Fetch Standard** (Living Standard) document, specifically covering the infrastructure and core algorithms governing HTTP fetching in the web platform. The content spans from the document preface through detailed specifications for URL schemes, HTTP methods, headers, request/response objects, bodies, authentication, CORS protocols, cache management, and various fetch modes (scheme, network, redirect). It also covers specialized topics like data URLs, cookies, cross-origin resource policies (CORS, COEP), and response filtering.

The document unifies fetching logic across APIs such as `fetch()`, `<img>`, `<script>`, `navigator.sendBeacon()`, and Service Workers, ensuring consistent behavior for redirects, caching, and security constraints. It supersedes previous inconsistent semantics, particularly regarding the `Origin` header.

# Cross-Chunk Summary

The Fetch Standard defines a unified architecture where fetching is conceptually simple (request in, response out) but involves complex underlying details. The flow generally proceeds as follows:

1.  **Infrastructure & Request Setup**: The process begins with URL parsing and method normalization (uppercase for standard verbs). Requests are constructed with attributes like `method`, `url`, `headers`, `body`, `mode` (same-origin, cors, no-cors, navigate), and `credentials`.
2.  **Fetching Algorithm**: A request is processed through a series of fetch steps:
    *   **Scheme Fetch**: Determines the protocol (`http`, `https`, `data`, etc.).
    *   **HTTP Fetch**: Handles the network interaction, potentially checking the cache first (if mode allows).
    *   **Redirects**: If redirects are followed, the process loops back to scheme fetch.
    *   **Network Error**: Returns a network error if the request fails or is blocked (e.g., mixed content, forbidden headers).
3.  **Response Handling**: The response is constructed with a status code, headers, and body. It has an associated type (`basic`, `cors`, `default`, `error`, `opaque`).
4.  **Security & Policies**: Throughout the process, security checks occur, including CORS preflight validation, Content Security Policy (CSP) hooking, Cross-Origin-Embedder-Policy (COEP), and Mixed Content blocking.
5.  **Body Processing**: Request and response bodies are handled as streams (`ReadableStream`), supporting incremental reading, cloning via `tee()`, and content decoding.

# Repeated Or Central Claims

*   **Unified Fetching Architecture**: The standard provides a single definition for fetching that applies to all web APIs, replacing fragmented implementations.
*   **Method Normalization**: HTTP methods are technically case-sensitive, but the standard normalizes them to uppercase (e.g., `GET`, `POST`) for compatibility. Custom methods must be carefully handled.
*   **Header Handling**: Headers are stored as ordered lists (multimaps). Non-`Set-Cookie` headers are combined into single values when exposed to JavaScript, while `Set-Cookie` is preserved individually.
*   **Request/Response Objects**: Both `Request` and `Response` objects encapsulate the state of a fetch operation, including attributes like `body`, `headers`, `status`, and various flags (`keepalive`, `aborted`).
*   **CORS Safety**: Headers are categorized as "CORS-safelisted" (safe to send cross-origin) or forbidden. This classification depends on character sets and length limits.
*   **Body Streams**: Bodies are represented by streams, allowing for incremental reading without loading the entire payload into memory immediately. Cloning a body shares the underlying stream via `tee()`.
*   **Security First**: The standard emphasizes security through CSP directives, COEP checks, origin serialization (preventing redirect taint leakage), and filtering responses to prevent information leakage.

# Important Local Details

*   **URL Schemes**: The standard supports "about", "blob", "data", "file" (fetch scheme), and HTTP(S).
*   **Method Categories**:
    *   *CORS-safelisted*: `GET`, `HEAD`, `POST` (and others not in forbidden lists).
    *   *Forbidden*: `CONNECT`, `TRACE`, `TRACK`.
    *   *Normalized*: Uppercase versions of standard verbs.
*   **Cache Modes**:
    *   `default`: Uses HTTP cache, revalidates if needed.
    *   `no-store`: Does not use the cache.
    *   `only-if-cached`: Returns cached response or network error (requires same-origin).
*   **Response Types**:
    *   `basic`/`cors`: Standard responses with headers accessible.
    *   `default`: Responses from `fetch()` in non-CORS modes.
    *   `opaque`: Responses where status and headers are hidden (often used for `no-cors`).
    *   `error`: Indicates a network error occurred.
*   **CORS Preflight**: A two-step process involving an `OPTIONS` request to check permissions before sending the actual request. This involves caching preflight results.
*   **Range Requests**: Support for `Range` headers and `Content-Range` responses, allowing partial content retrieval.
*   **Filtered Responses**: Used to expose limited views of a response (e.g., only image data) to prevent leaking sensitive header info or full body content in certain contexts.

# Candidate Wiki Hints

*   **Page: Fetch Standard Overview** – Introduction to goals, unified architecture, and API coverage.
*   **Page: HTTP Method Normalization** – Rules for uppercase conversion, forbidden methods, and CORS-safelisted lists.
*   **Page: Request & Response Objects** – Detailed attributes, default values, cloning mechanics, and lifecycle states.
*   **Page: Fetch Modes & Cache Interaction** – Explaining `same-origin`, `cors`, `no-cors`, `navigate` modes and cache behavior (`default`, `reload`, etc.).
*   **Page: Body Streams & Cloning** – How bodies work as streams, `tee()` usage, incremental reading algorithms.
*   **Page: CORS Protocol Deep Dive** – Preflight logic, header filtering, credentials handling, and exceptions.
*   **Page: Security Policies (CSP/COEP)** – Content Security Policy hooks, Cross-Origin-Embedder-Policy checks, and Mixed Content blocking.
*   **Page: Origin Resolution & Taint** – How origins are resolved, serialized, and how redirect-taint is computed to prevent leakage.

# Gaps Or Cautions

*   **Missing Raw Text**: While the chunk outline provides headings and summaries, specific algorithmic pseudocode (e.g., exact microstep sequences for `fetch()` implementation) is not fully detailed in the provided notes beyond high-level descriptions.
*   **Data URL Specifics**: Chunks 18–24 cover "data: URLs" but only list them as a heading without detailed content in the notes, potentially missing nuances on parsing or security constraints specific to data URIs.
*   **Service Worker Interaction**: The preface mentions Service Workers, but the detailed interaction between fetch algorithms and Service Worker event handlers (e.g., `fetch` events) is not elaborated upon in these chunks.
*   **Implementation Variations**: The notes mention "implementation-defined operations" (e.g., DNS queries) which may vary across browsers or environments; caution is needed when relying on specific IP resolution orders.
*   **Legacy Compatibility**: Some sections note that certain features (like opaque filtered responses for legacy image decoding) are discouraged for new specifications due to architectural limitations, implying potential deprecation or migration paths not fully detailed here.
