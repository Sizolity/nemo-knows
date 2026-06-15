---
title: Rollback Journal Mechanism
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/012-file-locking-and-concurrency.md
confidence: medium
---

# Rollback Journal Mechanism

The rollback journal mechanism is a concurrency control strategy introduced in SQLite version 3.0.0 to ensure ACID compliance and improve performance over version 2. Unlike Write-Ahead Logging (WAL), this method records all changes to a separate file before applying them to the main database.

## How It Works

In standard operation, the system manages the database as uniform-sized blocks called **pages**. When a transaction modifies data, it is first recorded in a rollback journal (typically named with a `-journal` suffix) rather than directly in the main database file. This allows for atomic commits even when multiple files are involved or if the system crashes during the write process.

The mechanism relies on the **pager module**, which coordinates access between threads or processes without managing B-Trees or text encodings. The pager tracks five specific locking states to manage these interactions:

- **UNLOCKED**: The default state where no locks are held.
- **SHARED**: Indicates a read-only lock, allowing multiple readers.
- **RESERVED**: Indicates the transaction is planning to write; only one holder is allowed.
- **PENDING**: A waiting state used to prevent writer starvation. It blocks new readers while existing ones finish, ensuring writers eventually acquire an **EXCLUSIVE** lock.
- **EXCLUSIVE**: Indicates the database is being written to; only one process holds this lock.

## Journal Management

The system distinguishes between active journals and "hot" journals. A journal becomes **hot** if it requires rollback due to a crash or power failure. This state is identified by:
1. The physical existence of the file.
2. A non-zero header size (greater than 512 bytes).
3. The absence of a **RESERVED** lock on the database.

To maintain integrity, the system must handle "super-journals" that coordinate changes across multiple **[[log]]**-based databases attached to the main file. If a crash occurs, the presence of a hot journal triggers a rollback upon recovery.

## Risks and Maintenance

While robust, the mechanism faces specific corruption risks:
- Manual deletion of hot journals can lead to data loss.
- Filesystem issues, such as **[[file-locking-states]]** moving journal files to `lost+found` or `fsync()` failures, can corrupt the database state.
- Rogue processes may interfere with locking states, leading to inconsistent views of the data.

By default, SQLite operates in autocommit mode, but executing `BEGIN TRANSACTION` disables this and engages the locking protocol described above.
