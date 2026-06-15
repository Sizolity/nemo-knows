---
title: Resilience Patterns
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/100-martin-fowler-microservices.md
confidence: medium
---

# Resilience Patterns

In microservice architectures, services communicate via lightweight mechanisms such as HTTP APIs or RPCs. Because these systems are distributed and inherently complex, applications must be designed to tolerate service failures. This approach ensures that the failure of one component does not cascade into a total system outage.

## Core Principles

### Design for Failure
Applications must explicitly assume that network partitions, timeouts, and individual service crashes will occur. Techniques used to handle these scenarios include:
- **Circuit Breakers**: Preventing calls to failing services from exhausting resources.
- **Bulkheads**: Isolating failures within specific components (similar to ship bulkheads) to prevent total collapse.
- **Timeouts**: Ensuring that requests do not hang indefinitely if a service is unresponsive.

### Monitoring and Emergent Behavior
Monitoring is vital for detecting emergent behavior resulting from choreographed interactions between multiple services. Without visibility into these distributed interactions, failures may go unnoticed until they cause significant impact.

## Implementation Context

Resilience patterns are often discussed alongside broader microservice principles:
- **Smart Endpoints and Dumb Pipes**: While logic resides in the endpoints, communication pipes must remain simple to reduce points of failure.
- **Decentralized Governance**: Teams should choose tools that fit their specific reliability needs rather than enforcing rigid mandates across the entire organization.
- **Infrastructure Automation**: Continuous Delivery and CI/CD pipelines are essential for managing the deployment of resilient services reliably across environments.

## Related Concepts

[[business-capability-services]]
[[log]]
[[resilience-patterns]]
