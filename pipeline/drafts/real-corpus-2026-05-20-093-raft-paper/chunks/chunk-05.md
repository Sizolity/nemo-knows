---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/093-raft-paper.md
confidence: medium
---
Chunk Context
This chunk (Lines 184-233) covers Section 6 ("Cluster membership changes") and the beginning of Section 7 ("Log compaction"). It details how Raft handles dynamic cluster reconfiguration safely using "joint consensus" to prevent split-brain scenarios, manages server additions/removals without availability gaps, and introduces snapshotting for log compaction.

Local Summary
Raft automates configuration changes to avoid downtime and operator error. Safety is maintained by using a two-phase transition involving a "joint consensus" state where the cluster operates under both old and new configurations simultaneously until the change is committed. The chunk also addresses challenges like adding new servers (which initially join as non-voting members), leaders not being in the new configuration, and preventing removed servers from disrupting elections. Finally, it introduces log compaction via snapshots to manage growing log sizes, describing the `InstallSnapshot` RPC mechanism for bringing followers up to date.

Key Claims
- Configuration changes must be automated; manual updates risk downtime and error.
- Direct switching between configurations is unsafe because atomic switches are impossible, risking split majorities.
- Joint consensus combines old ($C_{old}$) and new ($C_{new}$) configurations: log entries replicate to both, any server from either can lead, and agreement requires separate majorities from both.
- New servers join as non-voting members initially to avoid availability gaps while catching up on logs.
- If the current leader is not in the new configuration, it steps down once the new configuration entry is committed.
- Removed servers are prevented from disrupting elections by ignoring `RequestVote` RPCs received within the minimum election timeout of a current leader's heartbeat.
- Snapshots replace committed log entries to prevent unbounded log growth; followers take snapshots independently rather than relying solely on the leader to send them.
- Linearizable reads require leaders to commit a no-op entry at term start and verify leadership via heartbeats before responding.

Entities And Concepts
- Joint Consensus: A transitional configuration ($C_{old, new}$) combining servers from both old and new configurations to ensure safety during reconfiguration.
- Non-voting Members: Servers added to the cluster initially without voting rights to prevent availability issues while syncing logs.
- RequestVote RPC Handling: Logic where followers ignore votes from removed servers if they have recently heard from a current leader.
- Snapshotting: A compaction method where the state machine writes current state to stable storage, replacing old log entries.
- InstallSnapshot RPC: An RPC used by leaders to send snapshot chunks to followers that are too far behind.
- Leader Completeness Property: Guarantees a leader has all committed entries, crucial for safety during term transitions and reads.

Procedures And API Details
- Configuration Change Process:
  1. Leader creates joint consensus entry ($C_{old, new}$) and commits it to a majority of both old and new configurations.
  2. Once committed, leader creates the final new configuration entry ($C_{new}$).
  3. After $C_{new}$ is committed, old servers can be shut down.
- Non-voting Member Join: New servers join first as non-voting members; once logs are caught up, they transition to full voting status via standard reconfiguration.
- InstallSnapshot RPC Protocol:
  - Arguments: `term`, `leaderId`, `lastIncludedIndex`, `lastIncludedTerm`, `offset`, `data[]`, `done`.
  - Receiver Logic:
    1. Reply immediately if received term < current term.
    2. If first chunk (offset 0), create new snapshot file.
    3. Write data at specified offset.
    4. Wait for more chunks if not done.
    5. Save snapshot and discard existing/partial snapshots with smaller indices.
    6. Retain log entries following the snapshot if they match the snapshot's last included index/term.
    7. Discard entire log if conflicting or superseded.
- Leader Read Safety: Before responding to read-only requests, a leader commits a blank no-op entry and exchanges heartbeats with a majority of the cluster.

Nuance Or Contradictions
- Independence vs. Leadership: Snapshotting allows followers to reorganize data independently, departing from Raft's strong "leader principle." This is justified because consensus (and thus state consistency) has already been reached before snapshotting occurs.
- Log Replacement: Unlike traditional append-only logs, Raft servers may delete entire log segments once replaced by a snapshot, provided the remaining log entries maintain consistency with the snapshot's metadata.
- Timing Assumptions: The alternative approach of using a heartbeat-based "lease" for read safety relies on bounded clock skew, whereas Raft prefers explicit no-op commits and heartbeat checks to avoid timing dependencies.

Candidate Wiki Hints
- Page: **Raft Configuration Changes** (Focus on joint consensus mechanism and safety guarantees)
- Page: **Raft Log Compaction** (Focus on snapshotting strategy and `InstallSnapshot` RPC details)
- Page: **Raft Linearizable Semantics** (Focus on handling retries, unique serial numbers, and read-only request safety)
