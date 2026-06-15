---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/093-raft-paper.md
confidence: medium
---
Chunk Context
This chunk details the state machine behaviors for all Raft servers (Leader, Follower, Candidate) and defines the specific RPC protocols used to maintain consensus. It covers the `AppendEntries` and `RequestVote` RPC structures, receiver implementation logic, and the core safety properties (Election Safety, Leader Append-Only, Log Matching, Leader Completeness, State Machine Safety).

Local Summary
The Raft algorithm maintains a single leader that handles client requests while followers passively replicate logs. Servers transition between states based on election timeouts or incoming RPCs. The leader replicates log entries using `AppendEntries` RPCs, which also serve as heartbeats to prevent elections. Consistency is enforced via term numbers and log indices in the RPC arguments. Safety properties ensure that committed entries are durable and identical across all logs before being applied to state machines.

Key Claims
- All servers increment `lastApplied` if `commitIndex > lastApplied`.
- Any server receiving an RPC with a term greater than its current term must update its term and become a follower.
- A leader sends `AppendEntries` RPCs to replicate logs and maintain authority (heartbeats).
- An election occurs if a follower receives no communication for the duration of an election timeout.
- A candidate wins election upon receiving votes from a majority of servers.
- `AppendEntries` is used for both log replication and heartbeats.
- If a leader crashes, logs may become inconsistent; followers must reject entries that do not match the expected term/index in incoming RPCs.

Entities And Concepts
- Raft Cluster: A set of servers with states (Leader, Follower, Candidate).
- Term: A logical clock unit; terms are numbered consecutively and increase monotonically.
- Election Timeout: The duration after which a follower initiates an election if no communication is received.
- RequestVote RPC: Sent by candidates to request votes during elections.
- AppendEntries RPC: Sent by leaders to replicate logs, apply commits, and maintain heartbeats.
- State Machine: The application logic executed upon applying committed log entries.

Procedures And API Details
AppendEntries RPC Structure:
Arguments:
- term: Leader's current term.
- leaderId: Leader's server ID.
- prevLogIndex: Index of the log entry immediately preceding new ones.
- prevLogTerm: Term of the entry at prevLogIndex.
- entries[]: Log entries to store (empty for heartbeats).
- leaderCommit: Leader's commitIndex.

Results:
- term: Current term (updated by follower if necessary).
- success: Boolean indicating if the follower accepted the entries.

Receiver Implementation Logic:
1. Reply false if `term < currentTerm`.
2. Reply false if log lacks an entry at `prevLogIndex` with matching `prevLogTerm`.
3. If an existing entry conflicts (same index, different term), delete the existing entry and all following entries.
4. Append any new entries not already in the log.
5. If `leaderCommit > commitIndex`, set `commitIndex = min(leaderCommit, index of last new entry)`.

Leader Behavior:
- Sends initial empty `AppendEntries` RPCs (heartbeats) upon election and during idle periods.
- Appends command to local log for client requests; responds after applying to state machine.
- Sends `AppendEntries` starting at `nextIndex` if `last log index >= nextIndex`.
- Decrement `nextIndex` and retry if `AppendEntries` fails due to inconsistency.
- Updates `commitIndex` if a majority of servers have replicated entries up to an index N where `log[N].term == currentTerm`.

Candidate Behavior:
- Increment `currentTerm`, vote for self, reset election timer.
- Send `RequestVote` RPCs to all other servers.
- Become leader if votes received from majority.
- Convert to follower upon receiving valid `AppendEntries` from a new leader.
- Start new election if election timeout elapses without becoming leader.

Follower Behavior:
- Respond to RPCs from candidates and leaders.
- Convert to candidate if election timeout elapses without receiving `AppendEntries` or granting a vote.

Nuance Or Contradictions
- Raft avoids ranking systems for elections in favor of randomized election timeouts to prevent split votes and ensure rapid leader election.
- A candidate may receive an `AppendEntries` RPC from a new leader; if the leader's term is >= candidate's term, the candidate reverts to follower state immediately.
- Split votes can occur if many followers become candidates simultaneously, but randomized timeouts reduce likelihood of indefinite splitting.

Candidate Wiki Hints
- Raft Consensus Algorithm Overview
- Raft State Machine Safety Properties
- Raft Log Replication Mechanism
