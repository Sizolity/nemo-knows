---
title: Multi Container Pods
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/044-pods.md
confidence: medium
---

# Multi Container Pods

A Pod is the smallest deployable unit of computing in Kubernetes, acting as a group of one or more containers with shared storage and network resources running in a shared context. While a Pod can function as a single container, it supports multiple container models including sidecars and init containers to enable tightly coupled application architectures.

Within a multi-container Pod, all containers share an IP address, network ports, and filesystem volumes. This allows them to communicate via `localhost` and share ephemeral data through shared volumes. Containers are launched in the same Linux namespaces (cgroups), ensuring they run in a unified environment.

Pods are generally not created directly; instead, they are managed by workload resources like Deployments or StatefulSets. Most metadata regarding a Pod is immutable after creation, though specific fields such as container images can be updated. Pods are designed as ephemeral entities that run on a specific node until execution finishes, eviction occurs, or the node fails.
