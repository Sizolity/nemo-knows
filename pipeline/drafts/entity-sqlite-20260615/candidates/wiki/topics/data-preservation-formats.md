---
title: Data Preservation Formats
kind: topic
sources:
  - source.md
  - pipeline/raw/web/sqlite.md
confidence: medium
---

# Data Preservation Formats

SQLite represents a unique approach to data storage through its serverless architecture. Unlike traditional database systems that require a separate server process, this engine is embedded directly into host applications written in C. This design eliminates the need for external configuration or administration of a standalone server instance, making it highly portable and easy to integrate into various software environments.

The entire database ecosystem, including tables, indexes, triggers, and views, resides within a single portable file. This compact structure ensures compatibility across different hardware architectures, supporting both 32-bit and 64-bit systems as well as big-endian and little-endian processors. Due to this self-contained nature, the US Library of Congress has specifically recommended the SQLite file format for long-term data preservation efforts.

Originating in the year 2000 under the stewardship of D. Richard Hipp and Hwaci, the project operates under a public domain license. This legal status allows anyone to copy, modify, and distribute the software for commercial use without restriction or attribution requirements. Such freedom has facilitated its widespread adoption across nearly every Android and iOS device, major web browsers, popular operating systems, and programming language runtimes like Python.

Reliability is maintained through strict adherence to ACID compliance, ensuring that committed changes survive crashes and power failures. The system supports transactional integrity via rollback journals or write-ahead logging modes, which allow for concurrent reading by multiple threads while a single thread handles writing. With an estimated billions of running instances globally, the project maintains a rigorous automated test suite achieving full branch coverage with support commitments extending to at least 2050.
