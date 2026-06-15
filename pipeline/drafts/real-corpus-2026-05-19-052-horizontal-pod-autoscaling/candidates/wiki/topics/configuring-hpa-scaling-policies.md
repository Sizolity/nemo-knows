---
title: Configuring Hpa Scaling Policies
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/052-horizontal-pod-autoscaling.md
confidence: medium
---

# Configuring Hpa Scaling Policies

Horizontal Pod Autoscaling (HPA) changes the replica count of a scalable Kubernetes workload in response to measured demand. It is the horizontal counterpart to adding resources to a single Pod: instead of making one replica larger, the controller asks the target workload to run more or fewer replicas.

## How It Works

The control plane periodically compares observed metric values with the target values configured on the `HorizontalPodAutoscaler`. From that ratio, it computes a desired replica count and writes the result through the target's scale subresource. The calculation is adjusted for readiness and missing metrics so that transient or incomplete measurements do not cause over-eager scaling.

Scaling behavior is configurable in both directions. Operators can use policies to limit how quickly replica counts change, tolerances to ignore small metric deviations, and stabilization windows to reduce oscillation when load rises and falls quickly.

## Supported Target Resources

HPA works with resources that expose a scale subresource, including common workload controllers such as Deployments and StatefulSets. DaemonSets are outside this model because their Pod count is tied to node placement rather than a desired replica field.

## Metrics Sources

Policies can use several classes of metrics:
- Resource metrics (CPU/Memory)
- Container-level metrics
- Custom metrics
- External metrics fetched via specific APIs (`metrics.k8s.io`, `custom.metrics.k8s.io`, `external.metrics.k8s.io`)

## Scaling Algorithm

The controller uses the following formula to calculate desired replicas:

```text
desiredReplicas = ceil(currentReplicas * currentMetricValue / desiredMetricValue)
```

The formula is only the starting point. The controller then accounts for not-yet-ready Pods and unavailable metrics before deciding whether to scale.

## Stabilization and Tolerance

Two safeguards shape practical scaling behavior:
- **Stabilization windows** reuse recent recommendations to avoid abrupt reversals, especially during downscale.
- **Tolerance settings** suppress changes when the metric ratio is close enough to the target that scaling would add noise rather than stability.

## Management with kubectl

HPAs can be inspected and removed with ordinary `kubectl` resource commands. `kubectl autoscale` is a shortcut for creating an autoscaler with minimum and maximum replica bounds plus a metric target.
