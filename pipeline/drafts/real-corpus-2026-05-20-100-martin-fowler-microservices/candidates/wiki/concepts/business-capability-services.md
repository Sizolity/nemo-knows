---
title: Business Capability Services
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/100-martin-fowler-microservices.md
confidence: medium
---

# Business Capability Services

In a microservice architecture, services are organized around **business capabilities** rather than technical layers such as UI, server, or database. This organizational approach ensures that the system structure aligns with the organization's communication patterns, effectively reinforcing Conway's Law.

By grouping functionality into distinct units of business value, teams can become cross-functional and focus on specific domains. This contrasts with organizing around technology stacks, which often leads to silos that do not reflect actual business needs. The resulting firm module boundaries allow for independent deployment and scalability of specific capabilities without requiring full application rebuilds.

## Key Characteristics

- **Organizational Alignment**: Teams are structured around the capabilities they support, ensuring the software architecture mirrors the business organization.
- **Independent Lifecycle**: Each capability can be developed, deployed, and scaled independently using fully automated machinery.
- **Technology Agnosticism**: Services within a capability may utilize different programming languages and data storage technologies (polyglot persistence) while maintaining clear boundaries.

## Design Implications

Designing services around business capabilities requires accepting increased complexity in distributed systems. It demands that teams manage this complexity effectively, often utilizing **resilience-patterns** such as circuit breakers and bulkheads to handle failures inherent in remote communication between these independent units.
