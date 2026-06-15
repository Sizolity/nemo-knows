---
kind: topic
sources: [raw/web/corpus-2026-05-18/093-raft-paper.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source document is a PDF of the Raft consensus algorithm paper, retrieved on 2026-05-18 and converted to text via supplemental processing.
- Content covers metadata, introduction, state machine logic (Leader/Follower/Candidate), RPC protocols (`RequestVote`, `AppendEntries`), safety properties, cluster membership changes (joint consensus), log compaction (snapshots), and empirical evaluation studies comparing Raft vs. Paxos.
- The paper emphasizes that Raft is designed for understandability and practical implementation over the complexity of original Paxos, supported by student study results and formal TLA+ proofs for specific properties like Log Completeness.

## Candidate Wiki Pages
- wiki/sources/raft-paper.md — Primary source documentation for the Raft consensus algorithm paper metadata, acquisition path, and content scope.
- wiki/concepts/raft-consensus-algorithm.md — Core page covering the separation of concerns (election, replication, safety), state machines, RPC protocols, and leader-based architecture.
- wiki/concepts/paxos-vs-raft-comparison.md — Comparative analysis based on user study results, highlighting why Raft is easier to understand despite students' prior exposure to Paxos.
- wiki/topics/cluster-membership-joint-consensus.md — Detailed explanation of the two-phase configuration change process using joint consensus states to prevent split majorities during reconfiguration.
- wiki/topics/log-compaction-snapshots.md — Documentation on `InstallSnapshot` RPC, replacing log entries with state machine checkpoints, and managing unbounded log growth.

## Suggested Links
- raw/web/corpus-2026-05-18/assets/093-raft-paper.pdf

## Review Checklist
- [ ] Verify that all candidate pages are placed under valid directories (`wiki/sources/`, `wiki/concepts/`, `wiki/topics/`).
- [ ] Ensure no nested directories were created for candidate pages.
- [ ] Confirm that the source path matches the provided chunk index exactly.
