---
title: Event System Guide
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Event System Guide

The DOM Standard defines a complete lifecycle for events, ranging from definition and listener management to dispatching and abort mechanisms. Events are objects that signal occurrences and do not initiate actions themselves; they can be either synthetic (created by code) or native.

## Core Concepts

### Event Objects
An `Event` is the fundamental object used to represent an occurrence. It serves as a carrier for information about what happened, including details such as the target element, timing, and modifiers.

### Custom Events
The standard provides the `CustomEvent` interface, allowing developers to define their own event types with associated data payloads, extending the base functionality of the event system.

### Event Targets
The `EventTarget` interface is the base for objects that can receive events. Any object implementing this interface can have event listeners attached to it.

## Listener Management

Developers attach handlers to targets to react to specific occurrences. The system supports adding and removing these listeners dynamically during the application's lifecycle.

## Propagation Phases

Event propagation involves traversing ancestors in two distinct phases:
1.  **Capture Phase**: Traverses downwards from the root to the target.
2.  **Bubble Phase**: Traverses upwards from the target to the root.

### Shadow DOM Boundaries
When dealing with Shadow DOMs, specific retargeting logic is required via the `composedPath()` algorithm to correctly handle events crossing shadow boundary lines.

## Abort Mechanisms

Asynchronous APIs within the event system should utilize `AbortController` and `AbortSignal` to support cancellation. Promises associated with these signals must reject immediately if the signal is aborted, ensuring clean resource management.
