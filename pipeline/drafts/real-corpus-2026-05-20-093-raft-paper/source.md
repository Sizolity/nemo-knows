---
title: Raft Consensus Algorithm
kind: source
sources:
  - raw/web/corpus-2026-05-18/093-raft-paper.md
confidence: medium
---

## What It Is

Raft is a consensus algorithm designed for distributed systems to manage replicated state machines. Unlike Paxos, which prioritizes mathematical elegance, Raft emphasizes understandability and practical implementation through structural decomposition (separating leader election, log replication, and safety) and randomized leader election timeouts. The source material originates from a PDF retrieved on 2026-05-18 (`raw/web/corpus-2026-05-18/assets/093-raft-paper.pdf`), categorized under "Systems."

## Summary

The document details Raft's evolution as a practical alternative to Paxos, addressing the latter's complexity and lack of intuitive explanations. The paper progresses from high-level motivation regarding multi-decree consensus to low-level implementation mechanics including RPC protocols, state management, log compaction, and cluster membership changes. A key design philosophy is that while safety proofs are timing-independent, system availability relies heavily on hardware reliability and `electionTimeout` being significantly smaller than the Mean Time Between Failures (MTBF).

The algorithm ensures all servers maintain identical logs up to the last committed index by overwriting conflicting entries during replication. It enforces a "Leader Completeness" property where any new leader must contain all previously committed entries, preventing split-brain scenarios. Cluster reconfiguration proceeds through a transitional "joint consensus" phase rather than atomic switches. Log compaction is achieved via snapshots (`InstallSnapshot` RPC) to prevent unbounded log growth.

## Key Claims

- **Understandability:** Raft is fundamentally easier to understand and implement than Paxos, supported by empirical studies showing higher quiz scores (mean 25.7 vs 20.8) among students with prior Paxos experience.
- **Safety via Log Matching:** All servers must maintain identical logs up to the last committed index; inconsistencies are resolved by overwriting follower entries that conflict with the leader's log.
- **Leader Completeness:** A critical invariant is that a new leader must contain all previously committed entries, ensuring no future term can overwrite an already committed state.
- **Randomized Timers:** Leader election relies on randomized timeouts rather than fixed intervals to prevent split votes and ensure rapid convergence, reducing median downtime from over 10 seconds to ~287ms.
- **Joint Consensus:** Cluster reconfiguration is not atomic; it proceeds through a transitional "joint consensus" phase where the cluster maintains agreement under both old and new configurations until the change is fully committed.
- **Implementation Size:** The Raft implementation discussed is roughly 2000 lines of C++ (excluding tests and comments), indicating high conciseness relative to the algorithm's complexity.
- **Formal Verification Limits:** While Log Completeness has been mechanically proven in TLA+, type safety remains unverified mechanically, relying on informal proofs for State Machine Safety.

## Suggested Links

- none
