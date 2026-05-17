---
kind: source
sources:
  - fixture:meeting-notes
confidence: medium
---

# Wiki Maintenance Meeting

## What It Is

Notes from a planning discussion about moving from source-summary-only ingest toward reviewed candidate page generation, candidate evaluation, and post-apply wiki lint.

## Summary

The discussion identified the need to keep raw sources immutable, write only reviewed wiki pages, and verify each stage with deterministic checks. The next milestones were candidate draft generation, candidate draft evaluation, wiki lint, and multi-source regression evals.

## Key Claims

- Real functional tests matter more than only rerunning unit tests.
- Candidate drafts should be generated before approved apply writes them.
- Post-apply lint should report issues before humans decide whether to fix them.

## Suggested Links

- AGENTS.md
- eval harness
