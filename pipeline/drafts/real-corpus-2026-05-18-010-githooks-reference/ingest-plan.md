---
kind: topic
sources: [raw/web/corpus-2026-05-18/010-githooks-reference.md]
status: draft
---

# Ingest Plan

## Source Summary
- Comprehensive documentation of Git hooks available in the default template directory and how to configure them.
- Details the invocation timing, parameters, environment variables, and exit status semantics for each hook type.
- Covers server-side hooks (pre-receive, update, post-receive) including protocol details for `proc-receive`.
- Lists specialized hooks like `fsmonitor-watchman`, `p4-changelist`, and `sendemail-validate` with specific usage patterns.

## Candidate Wiki Pages
- wiki/sources/githooks-reference.md — Store the full reference text and metadata from the fetched source document.
- wiki/concepts/git-hooks.md — Create a conceptual page explaining the lifecycle, security model, and basic structure of Git hooks.
- wiki/topics/server-side-hooks.md — Dedicate a topic to server-enforced hooks (pre-receive, update, post-receive) and their protocols.

## Suggested Links
- https://git-scm.com/docs/githooks

## Review Checklist
- [ ] Verify all hook names listed in the source are included in the wiki pages.
- [ ] Ensure protocol details for `proc-receive` are accurately represented.
- [ ] Confirm that environment variable descriptions match the source text exactly.
