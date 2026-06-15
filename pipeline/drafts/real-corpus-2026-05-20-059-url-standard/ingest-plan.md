---
kind: topic
sources: [raw/web/corpus-2026-05-18/059-url-standard.md]
status: draft
---

# Ingest Plan

## Source Summary
- The document defines the **URL Standard** (Living Standard, April 2026), unifying legacy RFCs with modern browser implementations and establishing robust APIs (`URL`, `URLSearchParams`).
- It covers the full lifecycle of network identifiers: infrastructure rules, host representations (domains/IPs/IDNA), URL parsing/serialization state machines, origin determination, secure rendering guidelines, and form encoding formats.
- Includes extensive **compatibility matrices** for browser/runtime support across Firefox, Chrome, Safari, Edge, Node.js, and Opera, highlighting polyfill needs for features like `URL.canParse` and `sort()`.

## Candidate Wiki Pages
- wiki/sources/url-standard-intro.md — Summarizes goals, terminology shifts (URI vs URL), and infrastructure dependencies from Section 1.
- wiki/concepts/host-representation-types.md — Explains Domain, IP Address, Opaque Host, Empty Host, serialization modes (`isOpaque`), and IDNA processing.
- wiki/concepts/url-parsing-state-machine.md — Details the parser states (`scheme start`, `relative`, `path`) and transition rules for schemes, hosts, paths, and file URLs.
- wiki/concepts/secure-url-rendering.md — Best practices for address bar display: hiding credentials, eliding hosts, handling IDNs/confusables to prevent spoofing.
- wiki/topics/application-x-www-form-urlencoded.md — Canonical algorithm for parsing/serializing form data, highlighting `%20` vs `+` encoding differences.
- wiki/concepts/url-api-compatibility.md — Comprehensive resource detailing browser/runtime versions where `URL`, `URLSearchParams`, and methods like `sort()` became available.

## Suggested Links
- none

## Review Checklist
- [ ] Verify candidate pages are immediate children of `wiki/sources/`, `wiki/concepts/`, or `wiki/topics/`.
- [ ] Ensure no nested directories are created (e.g., avoid `wiki/sources/url-standard/intro.md`).
- [ ] Confirm that "public suffix" logic is noted as informational rather than a security boundary in the host page.
- [ ] Validate that the encoding discrepancy between `URL` and `URLSearchParams` is documented in the API or form encoding pages.
