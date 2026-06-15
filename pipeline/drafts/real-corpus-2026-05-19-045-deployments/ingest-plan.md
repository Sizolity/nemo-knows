---
kind: topic
sources: [raw/web/corpus-2026-05-18/045-deployments.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source is official Kubernetes documentation detailing the `Deployment` workload controller.
- It covers use cases, creation, updating, rolling back, scaling, and managing rollout states (progressing, complete, failed).
- Detailed explanations of strategies like RollingUpdate vs Recreate, maxSurge/maxUnavailable tuning, and pause/resume operations are included.
- The document provides a comprehensive API reference for the Deployment object and its status conditions.

## Candidate Wiki Pages
- wiki/sources/kubernetes-deployments-guide.md — To archive the full documentation content regarding Kubernetes Deployments for future reference.
- wiki/concepts/deployment-controller.md — To define the core concept of a Deployment as a declarative controller managing ReplicaSets and Pods.
- wiki/topics/deployment-strategies-and-lifecycle.md — To document specific operational procedures including rollout strategies, scaling logic, and status monitoring.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that candidate page slugs adhere to the immediate child directory rule.
- [ ] Ensure no nested directories are created for wiki pages.
- [ ] Confirm that source metadata (URL, date) is accurately reflected in the summary.
