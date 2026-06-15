---
title: Sqlite Page Structure
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/020-database-file-format.md
confidence: medium
---

# Sqlite Page Structure

SQLite stores a database as a page-oriented file format. The main database file contains the persistent database state, while transaction modes may add temporary companion files such as a rollback journal or write-ahead log.

## Page Organization

Every database file is divided into same-sized pages. The page size is fixed for the database and falls within SQLite's supported range. Page 1 is special because it begins with the database header, which records file-level metadata including the format signature, page size, counters, encoding, and schema information.

## B-tree Storage

Tables and indexes are represented with b-tree pages. Page types separate navigation, stored cells, free space management, and overflow storage:

- **Interior pages** route lookup through the tree.
- **Leaf pages** hold table or index entries.
- **Freelist pages** track unused pages that can be reused.
- **Overflow pages** store payload bytes that do not fit in the owning b-tree page.

### Table and Index B-trees

Table b-trees and index b-trees organize their keys differently. Ordinary table b-trees use integer rowids as keys and store row payloads in leaves. Index b-trees store index keys instead. `WITHOUT ROWID` tables use an index-style b-tree keyed by the primary key columns.

## Schema Representation

The schema is itself represented as table content. SQLite stores schema records in `sqlite_schema` and recognizes historical aliases such as `sqlite_master`. Because page 1 contains the database header, the schema b-tree root shares that first page with file metadata.

## Transaction Management Files

The page format interacts with transaction files:

- **Rollback journals** preserve pre-change page images so a failed transaction can be undone.
- **WAL files** receive new page versions before those changes are checkpointed back into the main database file.

## References

- [[sqlite-journal-modes]]
