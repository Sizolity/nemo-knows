---
kind: topic
sources: [raw/web/corpus-2026-05-18/100-martin-fowler-microservices.md]
status: draft
---

# Ingest Plan

## Source Summary
- Defines microservice architecture as a suite of independently deployable services organized around business capabilities.
- Contrasts microservices with monolithic applications, highlighting trade-offs in deployment, scalability, and complexity.
- Covers key principles including smart endpoints/dumb pipes, decentralized governance, polyglot persistence, and infrastructure automation.
- Discusses resilience patterns (circuit breakers), evolutionary design, and the risks of shifting complexity to service connections.

## Candidate Wiki Pages
- wiki/sources/martin-fowler-microservices.md — Store raw ingestion metadata and source text for this canonical architecture article.
- wiki/concepts/business-capability-services.md — Document the principle of organizing services around business domains rather than technical layers.
- wiki/topics/resilience-patterns.md — Capture patterns like circuit breakers, bulkheads, and timeouts mentioned in the "Design for failure" section.

## Suggested Links
- https://martinfowler.com/articles/microservices.html

## Review Checklist
- [ ] Verify that service boundary definitions align with Conway's Law discussion.
- [ ] Ensure polyglot persistence examples are cross-referenced with a database technology page if one exists.
- [ ] Confirm that the distinction between SOA and microservices is clearly articulated in the concepts draft.
