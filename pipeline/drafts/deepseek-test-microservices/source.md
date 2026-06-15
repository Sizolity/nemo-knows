---
title: Martin Fowler - Microservices
kind: source
sources:
  - raw/web/corpus-2026-05-18/100-martin-fowler-microservices.md
confidence: medium
---

## What It Is
The seminal article **"Microservices"** by Martin Fowler and James Lewis (published 25 March 2014), defining the microservice architectural style. It is hosted at `https://martinfowler.com/articles/microservices.html` and describes an approach to building single applications as suites of small, independently deployable services.

## Summary
The article introduces the microservice style by contrasting it with monolithic architectures. It outlines nine common characteristics: componentization via services, organization around business capabilities, product mindset over project mindset, smart endpoints and dumb pipes, decentralized governance, decentralized data management, infrastructure automation, design for failure, and evolutionary design. The authors discuss benefits, trade-offs, and the relationship between microservices and SOA. The piece ends with cautious optimism, noting that while early results are positive, long-term outcomes are not yet known, and team skill remains critical.

## Key Claims
- **Definition**: A microservice architecture is a suite of small services, each running in its own process and communicating via lightweight mechanisms (often HTTP APIs), built around business capabilities and independently deployable with full automation.
- **Componentization via Services**: Services are the primary unit of componentization, allowing independent deployment and explicit interfaces, at the cost of remote-call overhead and coarser-grained APIs.
- **Business Capability Alignment**: Teams are cross-functional and organized around business capabilities, not technology layers; service boundaries reinforce team boundaries (Conway’s Law).
- **Products not Projects**: Teams own the entire product lifecycle ("you build, you run it"), increasing operational responsibility and user empathy.
- **Smart Endpoints, Dumb Pipes**: Communication uses lightweight protocols (HTTP, lightweight messaging) with intelligence in the endpoints, avoiding heavy ESB-style orchestration.
- **Decentralized Governance**: Languages, frameworks, and data storage are chosen per service (polyglot persistence). Standards grow from shared libraries and internal open-source, not central mandates.
- **Decentralized Data Management**: Each service manages its own database; consistency is often eventual, with compensating operations rather than distributed transactions.
- **Infrastructure Automation**: Continuous Delivery and deployment pipelines, automated testing, and provisioning tools are essential; they make frequent, boring deployments possible.
- **Design for Failure**: Services must gracefully handle failures of other services; patterns like Circuit Breaker, real-time monitoring, and "Simian Army" chaos testing are used.
- **Evolutionary Design**: Services support independent replaceability; boundaries follow patterns of change. The article recommends starting with a well-modularized monolith and splitting when needed.
- **Trade-offs and Maturity**: Microservices are not a silver bullet; poorly defined boundaries, complex inter-service communication, and team skill constraints can lead to worse outcomes.

## Suggested Links
- **Papers referenced in the raw document:**
  - L. Lamport, “The Implementation of Reliable Distributed Multiprocess Systems”, 1978 - [http://research.microsoft.com/en-us/um/people/lamport/pubs/implementation.pdf](http://research.microsoft.com/en-us/um/people/lamport/pubs/implementation.pdf)
  - L. Lamport, R. Shostak, M. Pease, “The Byzantine Generals Problem”, 1982 - [http://www.cs.cornell.edu/courses/cs614/2004sp/papers/lsp82.pdf](http://www.cs.cornell.edu/courses/cs614/2004sp/papers/lsp82.pdf)
  - R.T. Fielding, “Architectural Styles and the Design of Network-based Software Architectures”, 2000 - [http://www.ics.uci.edu/~fielding/pubs/dissertation/top.htm](http://www.ics.uci.edu/~fielding/pubs/dissertation/top.htm)
  - E. A. Brewer, “Towards Robust Distributed Systems”, 2000 - [http://www.cs.berkeley.edu/~brewer/cs262b-2004/PODC-keynote.pdf](http://www.cs.berkeley.edu/~brewer/cs262b-2004/PODC-keynote.pdf)
  - E. Brewer, “CAP Twelve Years Later: How the 'Rules' Have Changed”, 2012 - [http://www.infoq.com/articles/cap-twelve-years-later-how-the-rules-have-changed](http://www.infoq.com/articles/cap-twelve-years-later-how-the-rules-have-changed)
- **Books and articles mentioned (no URLs in raw document):**
  - *Release It!* (Michael Nygard) – introduces Circuit Breaker, Bulkhead, Timeout patterns.
  - *Continuous Delivery* (Jez Humble, David Farley) – foundational for deployment automation.
  - *Domain-Driven Design* (Eric Evans) – inspires bounded context and service boundaries.
  - *The Art of Unix Programming* (Eric S. Raymond) – cited for monolith definition and Unix philosophy.
  - Martin Fowler’s Microservice Resource Guide and Microservice Trade-Offs article – referenced as further reading without explicit URL.
- **Tooling mentioned:** Netflix OSS, Dropwizard, RabbitMQ, ZeroMQ, protobufs.
