---
title: Log Compaction Snapshots
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/093-raft-paper.md
confidence: medium
---

# Log Compaction Snapshots

In the Raft consensus algorithm, maintaining a bounded memory footprint is critical for long-running replicated state machines. Without intervention, a server's append-only log would grow indefinitely as new entries are appended to replicate leader commands. To prevent unbounded log growth, Raft employs **log compaction** via **snapshots**.

When a follower's log becomes significantly larger than the leader's (or when the leader has committed enough entries to justify it), the leader sends an `InstallSnapshot` RPC to that follower. This operation overwrites the follower's existing log with a compressed snapshot of the state machine up to the current commit index. Once installed, the follower truncates its log to match the snapshot boundary and resumes normal replication for subsequent entries.

This mechanism ensures that safety is preserved because snapshots contain all committed history, while the remaining log entries only reflect uncommitted or newly committed commands. The transition relies on the **Leader Completeness** property: since a new leader must possess all previously committed entries (either in its log or within a snapshot), no future term can overwrite an already committed state.

Cluster reconfiguration during this process often utilizes **[[cluster-membership-joint-consensus]]**, allowing the system to transition from one membership configuration to another without downtime, ensuring that snapshots and logs remain consistent across the evolving cluster topology.
