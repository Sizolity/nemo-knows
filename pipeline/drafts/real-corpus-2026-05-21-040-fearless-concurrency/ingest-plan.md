---
kind: topic
sources: [raw/web/corpus-2026-05-18/040-fearless-concurrency.md]
status: draft
---

# Ingest Plan

## Source Summary
- Covers Rust's approach to "fearless concurrency," leveraging ownership and type systems to shift concurrency errors from runtime to compile-time.
- Discusses the distinction between concurrent (independent execution) and parallel (simultaneous execution) programming contexts.
- Outlines a curriculum covering thread creation, message-passing via channels, shared-state concurrency, and the `Sync`/`Send` traits.
- Highlights Rust's philosophy of offering a variety of low-level tools compared to higher-level languages that enforce specific abstractions.

## Candidate Wiki Pages
- wiki/sources/fearless-concurrency.md — To store the raw content and metadata from the Rust Book chapter on concurrency.
- wiki/concepts/fearless-concurrency.md — To define the core concept of "fearless concurrency" as a compile-time safety guarantee in Rust.
- wiki/topics/threading-and-safety.md — To document specific topics like `Send`, `Sync`, channels, and thread creation extracted from the source.

## Suggested Links
- https://doc.rust-lang.org/book/ch16-00-concurrency.html

## Review Checklist
- [ ] Verify that all code examples in the source are transcribed accurately to wiki/sources/fearless-concurrency.md.
- [ ] Ensure the distinction between concurrent and parallel is clarified in wiki/concepts/fearless-concurrency.md.
- [ ] Confirm that `Send` and `Sync` trait definitions align with current Rust standard library documentation.
