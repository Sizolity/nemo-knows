---
title: Horizontal Pod Autoscaling
kind: source
sources:
  - raw/web/corpus-2026-05-18/052-horizontal-pod-autoscaling.md
confidence: medium
---

# Horizontal Pod Autoscaling

## What It Is

Horizontal Pod Autoscaling (HPA) is a Kubernetes feature that automatically manages the scale of workloads by adjusting the number of Pods. Unlike vertical scaling, which increases resources for existing Pods, horizontal scaling deploys more or fewer Pods to match demand. The `HorizontalPodAutoscaler` acts as a controller within the Kubernetes control plane that periodically adjusts the desired replica count of target resources (such as Deployments or StatefulSets) based on observed metrics like CPU utilization, memory usage, or custom metrics.

## Summary

The HPA operates as a control loop with a default sync period of 15 seconds. It calculates the ratio between current and desired metric values to determine scaling actions. The controller considers Pod readiness, handles missing metrics conservatively, and applies stabilization windows to prevent rapid fluctuations ("thrashing"). Administrators can configure scaling policies to limit the rate of change, set tolerances for minor metric variations, and define separate behaviors for scale-up and scale-down operations.

## Key Claims

- **Target Resources**: HPA supports objects that support scaling, such as Deployments and StatefulSets. It does not apply to DaemonSets.
- **Metrics Sources**: Scaling can be based on resource metrics (CPU/Memory), container-level metrics, custom metrics, or external metrics fetched via specific APIs (`metrics.k8s.io`, `custom.metrics.k8s.io`, `external.metrics.k8s.io`).
- **Scaling Algorithm**: The controller uses a formula to calculate desired replicas: `desiredReplicas = ceil(currentReplicas * currentMetricValue / desiredMetricValue)`. It ignores Pods with missing metrics or those not yet ready during calculations to avoid misleading scaling decisions.
- **Stabilization and Tolerance**: Features include a stabilization window (default 5 minutes for downscale) to smooth out rapid metric changes and a tolerance setting (default 10%) to ignore minor fluctuations below a specified threshold.
- **kubectl Support**: Standard `kubectl` commands (`get`, `describe`, `delete`) apply to HPAs, and the specialized `kubectl autoscale` command allows quick creation of autoscalers with specific min/max replica counts and metric targets.

## Suggested Links

*   Kubernetes Documentation - Concepts - Workloads
*   Horizontal Pod Autoscaler Object
*   Metrics Server
*   Walkthrough example for horizontal pod autoscaling
*   kubectl autoscale documentation
*   Custom metrics adapter boilerplate
*   HorizontalPodAutoscaler API reference
