---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

**Heading Path:** `2. Append entry to the user agent’s CORS-preflight cache. > 4.12. Deferred fetching`
**Line Range:** 4503–4863
**Coverage:** Section 4.12 (Deferred fetching algorithm and task sources), Section 4.12.1 (Deferred fetching quota mechanics, policies, and examples), and the beginning of Section 5 (Fetch API usage patterns).

# Local Summary

This chunk details the "deferred fetch" feature in the Fetch Standard. It defines how to queue a fetch operation to execute later—specifically when a fetch group terminates or after a timeout. The text specifies the task source priority for deferred fetches, algorithms for queuing and processing these requests, and the calculation of total request length (including URL, referrer, headers, and body).

Section 4.12.1 focuses on quota management. It establishes a default top-level quota of 640 kibibytes, with sub-allocations for cross-origin nested documents (8 kibibytes each) via the `deferred-fetch-minimal` policy. The chunk explains how to calculate available quota per reporting origin (capped at 64 kibibytes), how permissions policies (`Permissions-Policy: deferred-fetch`) allow delegation of this quota, and provides examples of navigation scenarios affecting quota reservation and release. Finally, it introduces the high-level `fetch()` API, contrasting it with XMLHttpRequest and providing code snippets for handling blobs, headers, JSON parsing, URL parameters, and progressive streaming.

# Key Claims

- Deferred fetching allows requests to be invoked at the latest possible moment (e.g., when a fetch group terminates).
- The deferred fetch task source prioritizes tasks before other sources (like DOM manipulation) to reflect the most recent state of `fetchLater()` calls before running dependent scripts.
- To compute total request length, the algorithm sums the serialized URL (excluding fragment), referrer, header list lengths, and body length.
- The top-level traversable is allocated a default deferred-fetch quota of 640 kibibytes.
- By default, 128 kibibytes are reserved for delegating to cross-origin nested documents.
- Only 64 kibibytes of the allocated quota can be used concurrently for the same reporting origin to prevent opportunistic reservation by third-party libraries.
- The `Permissions-Policy` header controls delegation; `deferred-fetch-minimal` is enabled by default for all origins, while `deferred-fetch` is enabled only for the top-level document's origin by default.
- Cross-origin or cross-agent iframes receive a default of 8 kibibytes of quota.
- Same-origin nested documents share their parent's quota and do not reserve additional space unless navigation occurs under specific conditions.

# Entities And Concepts

- **Deferred Fetch**: A mechanism to queue fetch requests for later execution.
- **`fetchLater()`**: The conceptual API (represented in algorithms) used to queue a deferred fetch.
- **Deferred Fetch Record**: An object containing the request, notify function, and invoke state.
- **Deferred Fetch Task Source**: A specific task source prioritized before script-dependent sources.
- **Total Request Length**: A metric including URL, referrer, headers, and body used to check quota limits.
- **Quota**: The memory limit (default 640 kibibytes) for deferred fetches per top-level traversable.
- **Reporting Origin Quota**: The sub-limit of 64 kibibytes per origin.
- **`Permissions-Policy`**: Header used to control quota delegation (`deferred-fetch`, `deferred-fetch-minimal`).
- **Top-Level Traversable**: A "tab" or top-level document context holding the main quota pool.
- **Navigable Container**: An element (like an iframe) capable of having its own reserved quota.
- **`fetch()`**: The standard Fetch API method for fetching resources, supporting blobs, headers, JSON, and streaming.

# Procedures And API Details

**Algorithm: Queue a Deferred Fetch**
1. Populate request from client given request.
2. Set request’s service-workers mode to "none".
3. Set request’s keepalive to true.
4. Create a new deferred fetch record (containing request and `notify invoked` function).
5. Append record to request’s client’s fetch group’s deferred fetch records.
6. If `activateAfter` is non-null: Wait until `activateAfter` ms passes OR user agent believes scripts will be lost (background, hidden state), then process deferredRecord.
7. Return deferredRecord.

**Algorithm: Compute Total Request Length**
1. Get length of request’s URL (serialized with exclude fragment true).
2. Increment by length of request’s referrer (serialized).
3. For each header (name, value), increment by `name.length + value.length`.
4. Increment by request’s body’s length.
5. Return total.

**Algorithm: Process Deferred Fetches**
1. Iterate over fetch group’s deferred fetch records.
2. Call process a deferred fetch for each record.

**Algorithm: Process a Deferred Fetch**
1. If invoke state is not "pending", return.
2. Set invoke state to "sent".
3. Fetch the request.
4. Queue a global task on the deferred fetch task source to run `notify invoked`.

**Algorithm: Get Available Quota (Simplified Logic)**
- Inputs: Document, Origin.
- Logic checks if top-level, policy allowed, and current reserved quota.
- Returns available quota (e.g., 640, 512, or 0 depending on state).

**Code Example: Fetch Blob**
```javascript
fetch("/music/pk/altes-kamuffel.flac")
  .then(res => res.blob())
  .then(playBlob)
```

**Code Example: Check Header and Parse JSON**
```javascript
fetch("https://pk.example/berlin-calling.json", {mode:"cors"})
.then(res => {
  if (res.headers.get("content-type") &&
      res.headers.get("content-type").toLowerCase().indexOf("application/json") >= 0) {
    return res.json()
  } else {
    throw new TypeError()
  }
})
.then(processJSON)
```

**Code Example: Progressive Streaming**
```javascript
function consume(reader) {
  var total = 0
  return pump()
  function pump() {
    return reader.read().then(({done, value}) => {
      if (done) return
      total += value.byteLength
      log(`received ${value.byteLength} bytes (${total} bytes in total)`)
      return pump()
    })
  }
}

fetch("/music/pk/altes-kamuffel.flac")
.then(res => consume(res.body.getReader()))
.then(() => log("consumed the entire body without keeping the whole thing in memory!"))
.catch(e => log("something went wrong: " + e))
```

# Nuance Or Contradictions

- **Quota Calculation Complexity**: The algorithm for getting available quota involves checking multiple nested conditions regarding top-level status, policy permissions (`deferred-fetch` vs `deferred-fetch-minimal`), and current reserved quotas. The text notes that 640kb should be enough but clarifies the calculation results in specific values (512, etc.) based on these flags.
- **Policy Defaults**: There is a distinction between the default allowlist for `deferred-fetch` ("self") and `deferred-fetch-minimal` ("*"). The text explicitly states that disabling `deferred-fetch-minimal` for the top-level document collects the 128 kibibytes back into the main pool.
- **Navigation Effects**: When a navigable container navigates to a cross-origin URL, it may reserve quota (64kb or 8kb). If it later navigates away or changes state, that reserved quota is freed or lost, which affects subsequent calculations.
- **Request Length Definition**: The definition of "total request length" explicitly includes the referrer and header lengths, which might be counter-intuitive if one assumes only payload size matters for quota.

# Candidate Wiki Hints

1. **Page: Deferred Fetching**
   - *Content*: Explain the concept of deferred fetching, its use case (lazy loading heavy resources), and the `fetchLater()` mechanism.
   - *Sections*: Task Sources, Quota Management, Permissions Policy (`Permissions-Policy: deferred-fetch`).

2. **Page: Fetch API Overview**
   - *Content*: High-level introduction to `fetch()`, comparison with XMLHttpRequest, and common patterns (blobs, JSON, streaming).
   - *Sections*: Basic Usage, Response Handling, Streaming Bodies.

3. **Page: Request Length Calculation**
   - *Content*: Technical detail on how request size is computed for quota purposes (URL + referrer + headers + body).
   - *Note*: This is a niche but important implementation detail for developers managing large payloads.
