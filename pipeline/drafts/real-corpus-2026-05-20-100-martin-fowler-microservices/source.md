---
title: Microservices Architecture Principles
kind: source
sources:
  - raw/web/corpus-2026-05-18/100-martin-fowler-microservices.md
confidence: medium
---

# Microservices Architecture Principles

## What It Is
A **Microservice Architecture** is an approach to developing a single application as a suite of small services, each running in its own process and communicating with lightweight mechanisms, often an HTTP resource API. These services are built around business capabilities and are independently deployable by fully automated deployment machinery. The style features a bare minimum of centralized management; services may be written in different programming languages and use different data storage technologies.

## Summary
Martin Fowler describes microservices as a response to the frustrations of monolithic applications, where change cycles are tied together (requiring full rebuilds) and scaling requires scaling the entire application. Microservices provide firm module boundaries, allowing for independent deployment and scalability. While the style is not novel (rooted in Unix design principles), it offers significant benefits in organizational flexibility and technical modularity, provided teams manage the increased complexity of distributed systems effectively.

## Key Claims
- **Componentization via Services**: Services are out-of-process components that communicate remotely (e.g., web service requests or RPCs), unlike libraries which are linked into a program. This allows independent redeployment of specific services rather than the entire application.
- **Organized around Business Capabilities**: Teams should be cross-functional and organized around business capabilities, not technology layers (UI, server, database). This reinforces Conway's Law by ensuring the system structure matches the organization's communication structure effectively.
- **Smart Endpoints and Dumb Pipes**: Applications should own their own domain logic and act as filters. Communication mechanisms (pipes) should be simple (e.g., REST or lightweight messaging), avoiding complex central orchestration tools like Enterprise Service Buses (ESBs).
- **Decentralized Governance**: Teams should use the right tool for the job, potentially using multiple languages and data storage technologies per service (Polyglot Persistence). Standards should be battle-tested and enforced by tools rather than rigid corporate mandates.
- **Design for Failure**: Applications must tolerate service failures. Techniques include Circuit Breakers, Bulkheads, and Timeouts. Monitoring is vital to detect emergent behavior from choreographed interactions.
- **Evolutionary Design**: Services should be designed for replaceability rather than just evolution. Frequent, fast changes are possible, but moving code across service boundaries is harder than in-process libraries.
- **Infrastructure Automation**: Continuous Delivery and automation (CI/CD) are essential to manage the deployment of multiple services and environments reliably.

## Suggested Links
- [martinfowler.com/articles/microservices.html](https://martinfowler.com/articles/microservices.html)
- Release It! (Book reference)
- 12 Factor Apps (from Heroku)
- UK Government Digital Service design principles
- Domain-Driven Design (Bounded Contexts)
- Consumer-Driven Contracts
