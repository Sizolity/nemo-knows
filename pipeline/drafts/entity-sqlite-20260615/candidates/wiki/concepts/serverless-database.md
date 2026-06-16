---
title: Serverless Database
kind: concept
sources:
  - source.md
  - pipeline/raw/web/sqlite.md
confidence: medium
---

# Serverless Database

A serverless database is a self-contained SQL engine that embeds directly into host applications without requiring a separate server process. Unlike traditional systems that run as standalone services, this architecture links the database code to the application itself, eliminating the need for external installation or configuration. This design simplifies deployment by ensuring the database runs within the same memory space and execution context as the application logic.

The entire database structure, including tables, indexes, triggers, and views, resides in a single portable file on disk. This unified file format is stable and backwards compatible across different architectures, supporting both 32-bit and 64-bit systems as well as big-endian and little-endian processors. Due to its simplicity and stability, this storage format has been recommended by the US Library of Congress for long-term data preservation efforts using [[data-preservation-formats]].

Ubiquity is a defining characteristic of this technology, with billions of instances running globally across mobile devices, web browsers, and operating systems. Major platforms like Android and iOS bundle the engine natively, while programming languages such as Python include it in their standard libraries. This widespread adoption ensures that developers have access to a reliable tool without needing to manage complex infrastructure or licensing restrictions, as the software is often released into the public domain.

Reliability is maintained through robust transaction support that adheres to ACID principles, ensuring data integrity even during crashes or power failures. The system utilizes mechanisms like rollback journals or write-ahead logging to handle atomic commits and crash recovery. In particular, write-ahead logging allows multiple readers and a single writer to operate concurrently, enhancing performance while maintaining consistency without the overhead of managing separate server processes.
