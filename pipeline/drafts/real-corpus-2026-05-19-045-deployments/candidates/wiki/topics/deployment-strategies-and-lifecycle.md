---
title: Deployment Strategies And Lifecycle
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/045-deployments.md
confidence: medium
---

# Deployment Strategies And Lifecycle

In Kubernetes, a **Deployment** is a declarative controller responsible for managing a set of Pods to run stateless application workloads. It operates by describing a desired state and automatically adjusting the actual cluster state to match it at a controlled rate. This mechanism allows users to manage the lifecycle of Pods, ensuring the correct number of replicas are running while applying updates safely, typically without downtime.

## Core Capabilities

The Deployment controller enables several key operations:

- **Rollout Management**: It creates ReplicaSets in the background to roll out new applications.
- **State Updates**: By updating the PodTemplateSpec (e.g., changing container images or labels), it triggers a gradual scaling process where old ReplicaSets scale down while new ones scale up.
- **Rollbacks**: If the current state becomes unstable, such as when pods crash-loop, users can rollback to an earlier revision using a specific revision number or the last stable state. Revision history is preserved in the associated ReplicaSets.
- **Scaling**: Deployments support horizontal scaling or integration with HorizontalPodAutoscaler based on metrics like CPU utilization.
- **Pause and Resume**: Users can pause rollouts, apply multiple fixes to the PodTemplateSpec, and then resume the update process. Note that rollbacks are not possible while a rollout is paused until it is resumed.
- **Cleanup**: Once old ReplicaSets are no longer needed, they are removed via garbage collection.

## Strategy Types

Deployments support two primary strategies for updating Pods:

1.  **RollingUpdate (Default)**: This strategy gradually scales down old ReplicaSets while scaling up new ones. It supports configurable parameters such as `maxUnavailable` and `maxSurge` to control the update process.
2.  **Recreate**: This strategy terminates all existing Pods before creating new ones. It is useful for upgrades where zero downtime is not required.

A rollout is triggered specifically when the Deployment's Pod template (`.spec.template`) changes, such as updates to labels or container images. Scaling alone does not trigger a rollout.

## Status Conditions

The controller exposes several status conditions to indicate the health and progress of a deployment:

- **Progressing**: Indicates that the Deployment is creating a new ReplicaSet, scaling up/down, or waiting for Pods to become ready.
- **Available**: Indicates that the minimum number of required replicas are available.
- **Failed**: Indicates that the progress deadline has been exceeded or an error occurred (e.g., insufficient quota).

## Progress Deadline

By default, the Deployment waits up to 600 seconds (`progressDeadlineSeconds`) before marking a rollout as failed if it stalls. This timeout can be adjusted via patching. Additionally, the label selector is immutable after creation in `apps/v1`; changing it requires deleting and recreating the Deployment.
