---
title: Fetch Body Streams
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Fetch Body Streams

In the **Fetch Standard**, request and response bodies are defined as **ReadableStreams**. This architecture enables incremental reading of data without loading the entire payload into memory immediately.

The stream-based approach supports specific operations such as cloning via `tee()` for simultaneous consumption by multiple consumers. It also facilitates content decoding processes directly within the stream pipeline.

## Memory Management and Buffering

To handle potential resends due to timeouts or other network interruptions, request bodies are buffered to a limit of 64 KiB when the source is a stream.

Memory usage is managed hierarchically through **deferred fetching quotas**:
- Top-level quota: 640 KiB
- Concurrent per origin quota: 64 KiB

## Lifecycle Context

Body streams are integral to the broader **http-fetch-lifecycle**, interacting with infrastructure setup, connection management, and security protocols like CORS. The standard ensures that these streams adhere to strict security constraints and buffering limits throughout their existence.
