---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk covers the **Deferred Fetching** algorithm within the Fetch specification, detailing how requests are queued to execute at a later time (e.g., when a fetch group terminates or after a timeout). It also includes the **Fetch API** overview.

The content spans from section 4.12 ("Deferred fetching") through 4.12.1 ("Deferred fetching quota"), providing detailed algorithms for queuing, processing, and managing quotas for deferred requests. Finally, it introduces section 5, offering a high-level overview of the `fetch()` method and practical usage examples.

# Local Summary

The specification defines a mechanism to defer network requests until necessary, prioritizing them in a specific task source to ensure scripts reflecting the latest state run before dependent scripts. A strict quota system limits memory usage:
- **Top-level quota**: 640 kibibytes.
- **Same-origin nested documents**: Share the parent's quota.
- **Cross-origin nested documents**: Receive a default allocation of 8 kibibytes, configurable via `Permissions-Policy`.
- **Per-origin concurrent limit**: Only 64 kibibytes can be used concurrently for the same reporting origin to prevent third-party libraries from hoarding memory.

The chunk concludes with an overview of the `fetch()` API, highlighting its capabilities for low-level resource fetching, handling responses as Blobs or JSON, working with URL parameters, and progressive consumption of request bodies.

# Key Claims

- **Deferred Fetching Task Source**: User agents must prioritize tasks in the deferred fetch task source before other sources (like DOM manipulation) to ensure the most recent state of a `fetchLater()` call is reflected before running dependent scripts.
- **Quota Allocation**: The default top-level quota is 640 kibibytes. By default, 128 kibibytes are reserved for delegating to cross-origin nested documents (`deferred-fetch-minimal` policy).
- **Concurrent Usage Limit**: Out of the allocated quota, only 64 kibibytes can be used concurrently for a specific reporting origin (the request's URL's origin).
- **Quota Calculation**: Total request length includes the URL (without fragment), referrer, header list lengths, and body length.
- **Policy Control**: The feature is identified by `"deferred-fetch"` (default allowlist: `"self"`) and `"deferred-fetch-minimal"` (default allowlist: `"*"`).

# Entities And Concepts

- **Deferred Fetch Record**: A record created when `fetchLater()` is called, containing the request and a notification function.
- **Fetch Group**: The container managing deferred fetch records for a client.
- **Task Source**: Specifically the "deferred fetch task source," prioritized for executing deferred requests.
- **Permissions Policy**: Used to control how much quota is delegated to cross-origin nested documents (e.g., `Permissions-Policy: deferred-fetch=(self "https://fratop.example.com")`).
- **Reporting Origin**: The origin derived from the request's URL, which has its own concurrent usage limit of 64 kibibytes.
- **Top-level Traversable**: A tab or window acting as the root for quota allocation (default 640 kibibytes).

# Procedures And API Details

### Queuing a Deferred Fetch
To queue a deferred fetch given a `request`, `activateAfter` (optional timestamp), and `onActivatedWithoutTermination`:
1. Populate request from client.
2. Set service-workers mode to "none".
3. Set keepalive to true.
4. Create a new deferred fetch record.
5. Append the record to the fetch group's deferred fetch records.
6. If `activateAfter` is provided, wait for the timeout or a reason to believe scripts are about to be lost (e.g., backgrounding) before processing.

### Computing Total Request Length
Used to check against quota limits:
1. Length of request’s URL (exclude fragment).
2. Plus length of request’s referrer.
3. Plus sum of name and value lengths for all headers.
4. Plus length of request’s body.

### Processing Deferred Fetches
When a fetch group is terminated or timeout occurs:
1. Iterate through deferred fetch records in the group.
2. For each record with "pending" invoke state:
   - Set state to "sent".
   - Fetch the request.
   - Queue a global task on the deferred fetch task source to run the notification function.

### Managing Quota (Algorithm Overview)
To get available quota for a document and origin:
1. Determine if it is a top-level traversable and if policies allow usage.
2. Calculate base quota (e.g., 640, 512, or 0 based on policy).
3. Subtract reserved quotas for cross-origin children and pending requests.
4. Return the remaining quota or 0 if exhausted.

# Nuance Or Contradictions

- **Quota Reset on Navigation**: If a navigable container navigates to a same-origin document, any reserved quota for that container is freed. Conversely, navigating away from a cross-origin frame releases its specific reservation back into the pool.
- **Redirect Handling**: Quota calculations assume redirects are handled before origin checks. A document created via redirect knows its final origin only after redirects are resolved, affecting when quota can be reserved or checked.
- **Cross-Origin vs Same-Origin Frames**: Same-origin frames share the parent's quota pool. Cross-origin frames default to 8 kibibytes unless explicitly delegated more via `Permissions-Policy`. If a cross-origin frame navigates to a different origin, it loses any previously granted specific quota for that path.

# Candidate Wiki Hints

- **Topic: Deferred Fetching**
  - Explain the purpose of deferring requests (memory management, user interaction timing).
  - Detail the `fetchLater()` API and its parameters.

- **Topic: Fetch Quota Management**
  - Break down the quota hierarchy (Top-level -> Same-origin children -> Cross-origin children).
  - Explain the `Permissions-Policy` header usage for controlling cross-origin delegation.
  - Provide examples of request sizes that trigger quota exhaustion.

- **Topic: Fetch API Overview**
  - Summarize the capabilities of `fetch()` compared to `XMLHttpRequest`.
  - Include code snippets for Blob extraction, JSON handling, and progressive reading.
