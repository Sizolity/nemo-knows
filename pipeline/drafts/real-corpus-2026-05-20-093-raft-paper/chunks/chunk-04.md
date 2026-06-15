---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/093-raft-paper.md
confidence: medium
---
## Chunk Context
This chunk details the log consistency mechanism, leader election restrictions, and safety proofs within the Raft consensus algorithm. It covers how leaders handle inconsistent follower logs via AppendEntries RPCs, defines the "Leader Completeness Property," explains why entries from previous terms are never committed by counting replicas, and outlines the timing requirements for availability versus safety.

## Local Summary
When a leader powers up, it must ensure follower logs match its own. Inconsistencies (missing or extra entries) are resolved automatically: the leader finds the last agreed-upon entry, overwrites conflicting follower entries, and appends missing ones. To prevent split-brain scenarios where different leaders overwrite committed data, Raft enforces a restriction that a new leader must possess all committed entries from previous terms. Commitment is only declared for entries in the current term; prior entries are considered committed indirectly once the current term's entry is replicated. The chunk concludes with safety arguments proving this property holds and notes that follower crashes are handled simply by retrying RPCs indefinitely.

## Key Claims
- **Log Convergence:** A leader forces follower logs to duplicate its own by overwriting conflicting entries during AppendEntries consistency checks.
- **Leader Completeness Property:** The leader for any given term must contain all log entries committed in previous terms.
- **Commitment Rule:** Only log entries from the leader's *current* term are committed by counting replicas; older entries remain uncommitted until a future leader with those entries is elected.
- **Safety Proof:** If a leader commits an entry in term $T$, any future leader (term $U > T$) must also store that entry, ensuring no server applies a different command for the same index.
- **Timing Constraint:** Safety does not depend on timing, but availability requires `broadcastTime` $\ll$ `electionTimeout` $\ll$ `MTBF`.

## Entities And Concepts
- **AppendEntries RPC:** Used to replicate logs and check consistency; triggers log repair if it fails.
- **nextIndex:** The index of the next log entry a leader sends to a specific follower.
- **RequestVote RPC:** Implements the election restriction by allowing votes only if the candidate's log is as up-to-date as the voter's.
- **Log Matching Property:** Guarantees that if two logs have an entry at a given index, they share identical entries from the first through that index.
- **Leader Completeness Property:** Ensures a leader contains all previously committed entries.
- **State Machine Safety Property:** Ensures all servers apply exactly the same commands in the same order.
- **Idempotent RPCs:** Raft RPCs can be retried without side effects, allowing safe handling of crashed followers/candidates.

## Procedures And API Details
1.  **Log Repair Process:**
    -   Leader performs a consistency check via AppendEntries.
    -   If inconsistent, leader decrements `nextIndex` and retries.
    -   Optimization: Follower can send conflicting entry term/index to allow leader to skip directly to the last valid term.
2.  **Election Restriction Logic:**
    -   Candidate contacts majority of cluster via RequestVote.
    -   Voter denies vote if its log is more up-to-date than candidate's.
    -   "Up-to-date" definition: Later last-entry term > earlier; if terms equal, longer log > shorter log.
3.  **Commitment Verification:**
    -   Count replicas only for current-term entries.
    -   Once current-term entry is committed, all prior entries are implicitly committed due to Leader Completeness and Log Matching Properties.

## Nuance Or Contradictions
-   **Previous Term Entries:** A leader cannot immediately conclude an old entry is committed just because it exists on a majority of servers (Figure 8 scenario). It must wait for the current term's entry to be committed to guarantee safety across term boundaries.
-   **Optimization vs. Reality:** While optimizing RPC rejection by including conflict details reduces retries, the authors doubt this is necessary in practice due to infrequent failures and low probability of deep inconsistency.
-   **Timing Dependency:** Safety is strictly timing-independent, but availability (system responsiveness) relies heavily on `electionTimeout` being significantly smaller than server failure rates (`MTBF`).

## Candidate Wiki Hints
-   Raft Log Consistency Mechanism
-   Leader Completeness Property
-   Raft Election Restriction and RequestVote Logic
-   State Machine Safety in Consensus Algorithms
-   Timing Constraints in Distributed Systems (Broadcast Time vs. MTBF)
