---
title: SQLite Database Engine Overview
kind: source
sources:
  - pipeline/raw/web/sqlite.md
confidence: medium
---

## What It Is

SQLite is a serverless, self-contained SQL database engine written in C that embeds directly into host applications without requiring a separate server process. The entire database, including tables, indexes, triggers, and views, resides in a single portable file format compatible across 32-bit/64-bit architectures and big-endian/little-endian systems.

## Summary

Created by D. Richard Hipp and released in 2000, SQLite is maintained under the stewardship of Hwaci with source code placed in the public domain. It is bundled within nearly every Android and iOS device, major web browsers, popular operating systems, and programming language runtimes like Python. The project boasts a highly rigorous automated test suite achieving 100% branch coverage and commits to long-term support through at least 2050.

## Key Claims

- **Serverless Design**: Runs within the application process with zero configuration required.
- **Single File Storage**: Stores all data structures in one file, recommended by the US Library of Congress for long-term preservation.
- **Public Domain**: No license restrictions apply; free for commercial use without attribution.
- **Ubiquity**: Estimated to have billions of running instances globally across mobile and web platforms.
- **ACID Compliance**: Supports transactions with crash recovery via rollback journal or Write-Ahead Logging (WAL) modes.
- **Concurrency**: WAL mode enables concurrent reading by multiple threads and writing by a single thread.

## Suggested Links

- https://www.sqlite.org/
