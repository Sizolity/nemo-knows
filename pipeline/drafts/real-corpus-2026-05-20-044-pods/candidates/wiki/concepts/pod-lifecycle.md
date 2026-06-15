---
title: Pod Lifecycle
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/044-pods.md
confidence: medium
---

# Pod Lifecycle

A Pod is the smallest deployable unit of computing in Kubernetes. It acts as a group of one or more containers with shared storage and network resources, running in a shared context such as Linux namespaces and cgroups. A Pod models an application-specific "logical host," containing tightly coupled application containers, init containers for startup, and ephemeral containers for debugging.

## Management and Creation

Pods are generally not created directly; instead, they are managed by workload resources like Deployments, StatefulSets, DaemonSets, or Jobs. Direct Pod creation is rare, as these workload resources handle replication, rollout, and automatic healing (e.g., replacing failed Pods on a dead node). Unlike regular Pods managed by the control plane, static Pods are supervised directly by the kubelet on a specific node, often used for control plane components.

## Runtime Characteristics

Within a Pod, containers share an IP address, network ports, and can communicate via `localhost`. Pods enable data sharing through shared volumes. They support multiple container models, including single-container wrappers, sidecars, and init containers.

A Pod is designed as an ephemeral entity that runs on a specific node until execution finishes, eviction occurs, or the node fails. Most metadata (name, UID, creation timestamp) is immutable after creation; only specific fields like container images or tolerations can be updated in place.

## Resource Enforcement

- **CPU Limits:** Enforced via throttling.
- **Memory Limits:** Enforced via OOM kills.

## Related Concepts

[[multi-container-pods]]

{{log}}
