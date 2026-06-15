---
title: Fearless Concurrency
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://doc.rust-lang.org/book/ch16-00-concurrency.html
tags: [rust, web-corpus]
confidence: medium
---

# Fearless Concurrency

## Fetch Metadata

- Corpus item: 40
- Category: Rust
- Source URL: https://doc.rust-lang.org/book/ch16-00-concurrency.html
- Final URL: https://doc.rust-lang.org/book/ch16-00-concurrency.html
- Retrieved: 2026-05-18
- Content-Type: text/html
- Fetch status: ok
- Test value: Good cross-link with Go and SQLite concurrency.
- Fetched page title: Fearless Concurrency - The Rust Programming Language

## Retrieved Text

Keyboard shortcuts

Press ← or → to navigate between chapters

Press S or / to search in the book

Press ? to show this help

Press Esc to hide this help

[ ]

IFRAME: toc.html

(BUTTON)
* (BUTTON) Auto
* (BUTTON) Light
* (BUTTON) Rust
* (BUTTON) Coal
* (BUTTON) Navy
* (BUTTON) Ayu

(BUTTON)

The Rust Programming Language

____________________

Fearless Concurrency

Handling concurrent programming safely and efficiently is another of
Rust’s major goals. Concurrent programming, in which different parts of
a program execute independently, and parallel programming, in which
different parts of a program execute at the same time, are becoming
increasingly important as more computers take advantage of their
multiple processors. Historically, programming in these contexts has
been difficult and error-prone. Rust hopes to change that.

Initially, the Rust team thought that ensuring memory safety and
preventing concurrency problems were two separate challenges to be
solved with different methods. Over time, the team discovered that the
ownership and type systems are a powerful set of tools to help manage
memory safety and concurrency problems! By leveraging ownership and
type checking, many concurrency errors are compile-time errors in Rust
rather than runtime errors. Therefore, rather than making you spend
lots of time trying to reproduce the exact circumstances under which a
runtime concurrency bug occurs, incorrect code will refuse to compile
and present an error explaining the problem. As a result, you can fix
your code while you’re working on it rather than potentially after it
has been shipped to production. We’ve nicknamed this aspect of Rust
fearless concurrency. Fearless concurrency allows you to write code
that is free of subtle bugs and is easy to refactor without introducing
new bugs.

Note: For simplicity’s sake, we’ll refer to many of the problems as
concurrent rather than being more precise by saying concurrent and/or
parallel. For this chapter, please mentally substitute concurrent
and/or parallel whenever we use concurrent. In the next chapter, where
the distinction matters more, we’ll be more specific.

Many languages are dogmatic about the solutions they offer for handling
concurrent problems. For example, Erlang has elegant functionality for
message-passing concurrency but has only obscure ways to share state
between threads. Supporting only a subset of possible solutions is a
reasonable strategy for higher-level languages because a higher-level
language promises benefits from giving up some control to gain
abstractions. However, lower-level languages are expected to provide
the solution with the best performance in any given situation and have
fewer abstractions over the hardware. Therefore, Rust offers a variety
of tools for modeling problems in whatever way is appropriate for your
situation and requirements.

Here are the topics we’ll cover in this chapter:
* How to create threads to run multiple pieces of code at the same
time
* Message-passing concurrency, where channels send messages between
threads
* Shared-state concurrency, where multiple threads have access to
some piece of data
* The Sync and Send traits, which extend Rust’s concurrency
guarantees to user-defined types as well as types provided by the
standard library
