---
title: Horizontal Pod Autoscaler Algorithm
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/052-horizontal-pod-autoscaling.md
confidence: medium
---

# Horizontal Pod Autoscaler Algorithm

The Horizontal Pod Autoscaler (HPA) is a Kubernetes feature that automatically manages workload scale by adjusting the number of Pods. It functions as a control loop with a default sync period of 15 seconds, periodically calculating desired replica counts based on observed metrics such as CPU utilization, memory usage, or custom metrics.

## Scaling Logic

The controller determines the target replica count using a specific formula that compares current metric values against desired targets:

```text
desiredReplicas = ceil(currentReplicas * currentMetricValue / desiredMetricValue)
```

During these calculations, the algorithm ignores Pods with missing metrics or those not yet ready to prevent misleading scaling decisions.

## Stabilization Features

To prevent rapid fluctuations known as "thrashing," the HPA applies stabilization windows and tolerances:

- **Stabilization Window**: A default window of 5 minutes is used for downscale operations to smooth out rapid metric changes.
- **Tolerance**: A default setting of 10% allows the controller to ignore minor metric variations below a specified threshold.
- **Scaling Policies**: Administrators can configure separate behaviors for scale-up and scale-down operations, including limiting the rate of change.

## Supported Metrics

Scaling decisions are based on various metric sources fetched via specific APIs:
- Resource metrics (CPU/Memory)
- Container-level metrics
- Custom metrics
- External metrics

The HPA supports target resources such as Deployments and StatefulSets but does not apply to DaemonSets.
