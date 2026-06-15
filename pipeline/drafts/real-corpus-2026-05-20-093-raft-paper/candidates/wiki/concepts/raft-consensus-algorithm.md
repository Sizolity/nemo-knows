---
title: Raft Consensus Algorithm
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/093-raft-paper.md
confidence: medium
---

# Raft Consensus Algorithm

**Raft** is a consensus algorithm designed for distributed systems to manage replicated state machines. Unlike Paxos, which prioritizes mathematical elegance, Raft emphasizes understandability and practical implementation. It achieves this through structural decomposition—separating leader election, log replication, and safety—and the use of randomized leader election timeouts.

## Core Properties

### Understandability
Raft is fundamentally easier to understand and implement than Paxos. Empirical studies indicate higher comprehension scores among students with prior experience in consensus algorithms when studying Raft compared to Paxos.

### Safety via Log Matching
The algorithm ensures that all servers maintain identical logs up to the last committed index. If inconsistencies arise, conflicting entries on followers are overwritten by the leader's log during replication.

### Leader Completeness
A critical invariant known as "Leader Completeness" requires that any new leader must contain all previously committed entries. This prevents split-brain scenarios where a future term could overwrite an already committed state.

### Randomized Timers
Leader election relies on randomized timeouts rather than fixed intervals to prevent split votes and ensure rapid convergence. This design reduces median downtime from over 10 seconds to approximately 287ms. The system availability relies heavily on hardware reliability and ensuring the `electionTimeout` is significantly smaller than the Mean Time Between Failures (MTBF).

### Joint Consensus
Cluster reconfiguration is not atomic; it proceeds through a transitional "joint consensus" phase rather than an immediate switch. During this phase, the cluster maintains agreement under both old and new configurations until the change is fully committed. For more details on cluster membership changes, see [[cluster-membership-joint-consensus]].

## Implementation Details

### Log Compaction
To prevent unbounded log growth, Raft achieves log compaction via snapshots using the `InstallSnapshot` RPC. This process replaces a large portion of the log with a snapshot file. For details on this mechanism, see [[log-compaction-snapshots]].

### Implementation Size
The referenced implementation is roughly 2000 lines of C++ (excluding tests and comments), indicating high conciseness relative to the algorithm's complexity. The source material originates from a PDF retrieved on 2026-05-18.

## Formal Verification Limits
While Log Completeness has been mechanically proven in TLA+, type safety remains unverified mechanically, relying on informal proofs for State Machine Safety.
