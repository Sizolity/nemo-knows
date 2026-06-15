---
title: Browser Compatibility Matrix
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Browser Compatibility Matrix

The DOM Standard defines a hierarchical, tree-based interface for representing and manipulating documents in web browsers. The specification establishes a structural model comprising light trees (document) and shadow trees, organized through specific node types, mixins, and interfaces.

## Core Architecture

### Tree Hierarchy
The DOM is fundamentally a finite hierarchical tree structure where nodes have parents, children, and siblings, with relationships defined as inclusive or exclusive (e.g., ancestor vs. parent). The specification begins with infrastructure concepts, defining trees, ordered sets, selectors, and name validation rules.

### Shadow DOM Isolation
Shadow trees are attached to light trees but can be closed to isolate their internal structure. Events and nodes within closed shadow trees require specific retargeting logic via the `composedPath()` algorithm. The standard covers the complete lifecycle of events—from definition and listener management to dispatching and abort mechanisms—alongside utilities for text manipulation (Ranges), traversal (TreeWalker), querying (XPath/XSLT), and mutation observation.

## Event System

Events are objects that signal occurrences; they do not initiate actions. They can be synthetic (created by code) or native. Propagation involves traversing ancestors in two phases: capture (downwards) and bubble (upwards), with specific handling for shadow DOM boundaries via the `composedPath()` algorithm.

## Mutation Handling

DOM mutations are handled via algorithms that separate "insertion steps" (structural changes) from "post-connection steps" (side effects like style application or script execution). This separation ensures atomicity for batch operations and triggers custom element lifecycle callbacks.

## Abort Mechanism

Asynchronous APIs should use `AbortController` and `AbortSignal` to support cancellation. Promises associated with these signals must reject immediately if the signal is aborted.

## Mixin Architecture

The DOM relies heavily on mixin interfaces (`DocumentOrShadowRoot`, `ParentNode`, `ChildNode`, `Slottable`) to share functionality between disparate node types without bloating individual interface definitions.

## Legacy and Namespace Considerations

The standard includes legacy extensions (e.g., `Window.event`, `initEvent`) marked as deprecated or replaceable to ensure backward compatibility while encouraging modern API usage. Additionally, CSS selectors within the DOM standard do not currently support namespaces, which may limit certain querying strategies.
