---
kind: topic
sources: [raw/web/corpus-2026-05-18/052-horizontal-pod-autoscaling.md]
status: draft
---

# Ingest Plan

## Source Summary
- The document details the HorizontalPodAutoscaler (HPA) implementation in Kubernetes, covering its role as a control loop to automatically scale workloads based on demand.
- It explains the underlying algorithms for calculating replica counts using resource metrics (CPU/Memory), custom metrics, and external metrics APIs.
- Key operational behaviors are described, including stabilization windows to prevent thrashing, tolerance thresholds for minor fluctuations, and handling of Pod readiness during startup.
- The text provides guidance on configuring scaling policies (scale-up/down rates) and migrating from manual replica counts to managed autoscaling.

## Candidate Wiki Pages
- wiki/sources/052-horizontal-pod-autoscaling.md — Raw ingestion artifact documenting the Kubernetes HPA feature and its source URL.
- wiki/concepts/horizontal-pod-autoscaler-algorithm.md — Explains the mathematical logic behind scaling decisions, metric averaging, and readiness handling.
- wiki/topics/configuring-hpa-scaling-policies.md — Practical guide for setting stabilization windows, tolerances, and policies to manage scale-up/down behavior.

## Suggested Links
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ (Explicitly present in source metadata)

## Review Checklist
- [ ] Verify that all metric API references (metrics.k8s.io, custom.metrics.k8s.io) are linked to appropriate wiki pages or external resources.
- [ ] Ensure the distinction between autoscaling/v1 and autoscaling/v2 is clearly documented in the concept page.
- [ ] Confirm that examples for container resource metrics align with the Kubernetes v1.30 feature state mentioned in the source.
