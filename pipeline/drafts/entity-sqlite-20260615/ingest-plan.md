---
kind: topic
sources: [pipeline/raw/web/sqlite.md]
status: draft
---

# Ingest Plan

## Source Summary
- Describes SQLite as a C library implementing a serverless, self-contained SQL database engine stored in a single file.
- Highlights the project's public domain licensing, stewardship by D. Richard Hipp and Hwaci, and deployment in billions of devices (Android, iOS, browsers).
- Notes robust reliability features including ACID transactions via rollback journal or WAL mode, and a test suite with 100% branch coverage supporting operations until 2050.

## Candidate Wiki Pages
- wiki/sources/sqlite.md — To document the raw source content regarding SQLite's definition, history, and deployment metrics.
- wiki/entities/sqlite.md — To capture the specific software entity details, licensing status (public domain), and stewardship information.
- wiki/concepts/serverless-database.md — To define the architectural concept of serverless database engines as exemplified by SQLite's embedded design.
- wiki/topics/data-preservation-formats.md — To discuss SQLite's file format stability and its recommendation by the US Library of Congress for long-term archival.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that candidate pages reside only in immediate subdirectories (sources, entities, concepts, topics).
- [ ] Ensure no nested directory structures are created within wiki/sources or wiki/entities.
- [ ] Confirm that the source kind (technical documentation) correctly maps to concept and topic candidates rather than narrative themes.
- [ ] Check that wiki/index.md and schema files are not included in the candidate list.
