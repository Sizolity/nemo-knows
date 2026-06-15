---
title: Deployment Controller
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/045-deployments.md
confidence: medium
---

# Deployment Controller

A Deployment Controller is the Kubernetes controller responsible for reconciling a Deployment object. The Deployment describes the desired application rollout, and the controller manages ReplicaSets and Pods until the cluster state matches that desired shape.

## Purpose

Deployments are mainly used for replaceable application replicas. They keep the requested number of Pods available, coordinate changes to the Pod template, and provide operational controls for rollout, pause, resume, and rollback workflows.

## Rollout Mechanics

A rollout begins when the Deployment's Pod template changes. Changes such as a new container image or template label cause the controller to create or update ReplicaSets. Scaling the Deployment changes replica counts, but does not by itself create a new rollout revision.

The controller coordinates several lifecycle actions:
- creating ReplicaSets for the desired Pod template;
- shifting replicas between old and new ReplicaSets during updates;
- returning to an earlier revision when a rollout fails;
- cleaning up older ReplicaSets according to history limits.

## Strategies

The Deployment Controller supports two strategy types:

- **RollingUpdate** gradually moves capacity from old ReplicaSets to new ones. `maxUnavailable` and `maxSurge` control the disruption and temporary extra capacity allowed.
- **Recreate** removes the existing Pods before starting replacements, which is simpler but does not preserve availability during the transition.

## Status Conditions

Deployment status conditions summarize rollout health:

- **Progressing**: Indicates the Deployment is creating a new ReplicaSet, scaling up/down, or waiting for Pods to become ready.
- **Available**: Indicates that the minimum number of replicas required are available.
- **Failed**: Indicates that the progress deadline has been exceeded or an error occurred (e.g., insufficient quota).

## Progress Deadline

`progressDeadlineSeconds` bounds how long the controller waits for rollout progress before reporting failure. The default is 600 seconds.

## Pausing Rollouts

Operators can pause a rollout, batch several template changes, and resume later. Rollback remains unavailable while the rollout is paused.

## Constraints

For `apps/v1`, the selector is effectively part of the Deployment's identity. Changing it requires recreating the Deployment rather than editing the selector in place.
