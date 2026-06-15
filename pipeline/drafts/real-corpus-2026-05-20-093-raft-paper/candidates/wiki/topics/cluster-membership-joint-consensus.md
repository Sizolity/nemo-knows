---
title: Cluster Membership Joint Consensus
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/093-raft-paper.md
confidence: medium
---

# Cluster Membership Joint Consensus

In the Raft consensus algorithm, cluster membership changes are not performed atomically. Instead, the system transitions through a specialized **joint consensus** phase to safely reconfigure the cluster without risking data loss or split-brain scenarios.

During this transitional period, the cluster operates under both the old and new configurations simultaneously. This design ensures that any leader elected during the transition must have logs consistent with the previously committed state, regardless of which configuration they are running. The joint consensus phase continues until a leader is successfully elected that belongs to the new configuration and can commit entries confirming the transition. Once this condition is met, the system fully adopts the new configuration, and the old leaders leave the cluster.

This approach contrasts with atomic switches, which could lead to inconsistencies if a leader from the old configuration were to be re-elected before the change was fully propagated. By maintaining agreement under dual configurations, Raft guarantees safety while allowing for smooth membership updates.
