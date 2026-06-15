---
kind: topic
sources: [raw/web/corpus-2026-05-18/100-martin-fowler-microservices.md]
status: draft
---

# Ingest Plan

## Source Summary
- Defines microservice architecture as suites of independently deployable services organized around business capabilities, with decentralized governance, data management, and infrastructure automation.
- Contrasts microservices with monolithic applications and discusses characteristics: componentization via services, smart endpoints and dumb pipes, evolutionary design, and design for failure.
- Highlights trade-offs, including operational complexity, distributed data consistency, and the need for sophisticated monitoring and team maturity.
- References related patterns and concepts such as bounded contexts, circuit breaker, continuous delivery, and consumer-driven contracts.

## Candidate Wiki Pages
- wiki/sources/martin-fowler-microservices.md — capture source metadata, publication context, and a structured summary of this foundational article.
- wiki/concepts/microservices.md — define the architectural style, its common characteristics, and key trade-offs as presented in the source.
- wiki/concepts/monolithic-architecture.md — document the monolithic counterpart used as a baseline for comparison, including its strengths and the reasons for migration.

## Suggested Links
- Conway’s Law
- Domain-Driven Design (Bounded Context)
- Continuous Delivery
- Polyglot Persistence
- Tolerant Reader pattern
- Consumer-Driven Contracts pattern
- Circuit Breaker pattern
- Smart endpoints and dumb pipes

## Review Checklist
- [ ] Confirm that all claims and definitions align with the original article.
- [ ] Check for existing candidate pages to avoid duplication.
- [ ] Validate that all suggested links correspond to terms explicitly mentioned or described in the source.
- [ ] Ensure candidate page titles and slugs follow wiki naming conventions.
