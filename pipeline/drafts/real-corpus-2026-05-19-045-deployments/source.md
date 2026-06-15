---
title: Deployments (Kubernetes)
kind: source
sources:
  - raw/web/corpus-2026-05-18/045-deployments.md
confidence: medium
---

# Deployments (Kubernetes)

## What It Is
A Kubernetes Deployment is a declarative controller that manages a set of Pods to run an application workload, typically one that does not maintain state. It provides declarative updates for Pods and ReplicaSets by describing a desired state; the Deployment Controller then changes the actual state to match the desired state at a controlled rate. Deployments can be used to create new ReplicaSets or remove existing ones while adopting their resources.

## Summary
Deployments allow users to manage the lifecycle of Pods, ensuring that the correct number of replicas are running and that updates are applied safely without downtime (by default). Key capabilities include:
- Creating a Deployment to rollout a ReplicaSet which creates Pods in the background.
- Updating the PodTemplateSpec to declare new states, triggering a gradual scaling process where old ReplicaSets are scaled down while new ones scale up.
- Rolling back to an earlier revision if the current state is unstable (e.g., crash looping).
- Scaling deployments horizontally or using HorizontalPodAutoscaler based on metrics like CPU utilization.
- Pausing and resuming rollouts to apply multiple fixes before restarting updates.
- Cleaning up older ReplicaSets via garbage collection once they are no longer needed.

## Key Claims
- **Use Case**: Deployments are ideal for stateless applications where Pods can be replaced easily. They manage ReplicaSets to bring up a specific number of replicas (default is 1).
- **Rollout Mechanics**: A rollout is triggered only when the Deployment's Pod template (`.spec.template`) changes, such as updates to labels or container images. Scaling alone does not trigger a rollout.
- **Strategy Types**:
  - `RollingUpdate` (default): Gradually scales down old ReplicaSets and scales up new ones. Supports configurable `maxUnavailable` and `maxSurge` parameters.
  - `Recreate`: Terminates all existing Pods before creating new ones; useful for upgrades where zero downtime is not required.
- **Rollback**: Users can rollback to a previous revision by specifying a specific revision number or rolling back to the last stable state. Revision history is stored in the ReplicaSets controlled by the Deployment.
- **Status Conditions**:
  - `Progressing`: Indicates the Deployment is creating a new ReplicaSet, scaling up/down, or waiting for Pods to become ready.
  - `Available`: Indicates that the minimum number of replicas required are available.
  - `Failed`: Indicates that the progress deadline has been exceeded or an error occurred (e.g., insufficient quota).
- **Progress Deadline**: By default, the Deployment waits up to 600 seconds (`progressDeadlineSeconds`) before marking a rollout as failed if it stalls. This can be adjusted via patching.
- **Label Selector Updates**: The label selector is immutable after creation in `apps/v1`. Changing the selector requires deleting and recreating the Deployment.
- **Pausing Rollouts**: Users can pause rollouts using `kubectl rollout pause`, make changes to the PodTemplateSpec, and then resume with `kubectl rollout resume`. Rollbacks are not possible while paused until resumed.

## Suggested Links
- [Deployments | Kubernetes](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
