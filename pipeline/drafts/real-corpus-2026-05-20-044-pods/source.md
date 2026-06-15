---
title: Pods
kind: source
sources:
  - raw/web/corpus-2026-05-18/044-pods.md
confidence: medium
---

# Pods

## What It Is
A Pod is the smallest deployable unit of computing in Kubernetes. It acts as a group of one or more containers with shared storage and network resources, running in a shared context (Linux namespaces, cgroups). A Pod models an application-specific "logical host," containing tightly coupled application containers, init containers for startup, and ephemeral containers for debugging.

## Summary
Pods are generally not created directly; instead, they are managed by workload resources like Deployments, StatefulSets, DaemonSets, or Jobs. They are designed as ephemeral entities that run on a specific node until execution finishes, eviction occurs, or the node fails. Within a Pod, containers share an IP address, network ports, and can communicate via `localhost`. Pods enable data sharing through shared volumes and support multiple container models, including single-container wrappers, sidecars, and init containers.

## Key Claims
- **Smallest Unit:** Pods are the smallest deployable units; Kubernetes manages Pods rather than individual containers.
- **Shared Context:** Containers in a Pod share a network namespace (IP address) and filesystem volumes.
- **Workload Management:** Direct Pod creation is rare; workload resources handle replication, rollout, and automatic healing (e.g., replacing failed Pods on a dead node).
- **Immutability:** Most metadata (name, UID, creation timestamp) is immutable after creation. Only specific fields like container images or tolerations can be updated in place.
- **Resource Enforcement:** CPU limits are enforced via throttling, and memory limits are enforced via OOM kills.
- **Static Pods:** Unlike regular Pods managed by the control plane, static Pods are supervised directly by the kubelet on a specific node (often used for control plane components).

## Suggested Links
*   https://kubernetes.io/docs/concepts/workloads/pods/
*   Kubernetes Documentation
*   Kubernetes Blog
*   Training
*   Careers
*   Partners
*   Community
*   Versions
