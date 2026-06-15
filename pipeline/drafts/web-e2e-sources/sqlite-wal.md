# SQLite Write-Ahead Logging Notes

Source URL: https://www.sqlite.org/wal.html

SQLite's write-ahead logging documentation describes WAL as an alternative to
the default rollback journal. In rollback-journal mode, the original unchanged
database content is preserved in a rollback journal and changes are written
directly into the database file. In WAL mode, changes are appended to a separate
write-ahead log file and are later moved back into the main database by a
checkpoint.

The documentation emphasizes that WAL can improve concurrency because readers
can continue reading the original database while writers append changes to the
WAL. A reader uses the database file plus the WAL up to a recorded end mark, so
newer commits do not disturb the reader's snapshot. Writers append new content
to the WAL. A checkpoint copies committed frames from the WAL back into the main
database file.

The page also describes trade-offs. WAL requires extra `-wal` and `-shm` files.
It normally requires all processes using the database to be on the same host
because shared memory is used for coordination. Very large transactions are no
longer a special disadvantage in modern SQLite, but checkpoint behavior can
still affect latency. WAL is persistent at the database-file level until the
journal mode is changed again.

Operationally, WAL gives SQLite a different concurrency and durability profile:
it is often faster and allows readers and writers to overlap, but users need to
understand checkpointing, shared-memory constraints, and the additional files
that accompany the database.
