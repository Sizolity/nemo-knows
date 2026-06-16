# SQLite

Source: distilled from the public SQLite documentation at https://www.sqlite.org/
(about page, "Most Widely Deployed and Used Database Engine", testing, and
long-term support notes). This file records widely known, verifiable facts about
SQLite for ingest testing of entity pages.

## What SQLite is

SQLite is a software library, written in the C programming language, that
implements a small, fast, self-contained, full-featured SQL database engine. It
is not a standalone server process. Instead, the engine is linked directly into
the host application and runs in the same process, so there is no separate
database server to install, configure, or administer. The project describes this
design as serverless and zero-configuration.

An entire SQLite database — tables, indexes, triggers, and views — is stored in a
single cross-platform file on disk. That file format is stable, backwards
compatible, and portable between 32-bit and 64-bit machines and between
big-endian and little-endian architectures. The United States Library of Congress
has recommended the SQLite file format as a storage format for long-term data
preservation.

## Origin and stewardship

SQLite was created by D. Richard Hipp, who first released it in the year 2000.
Development is led by a small core team and is overseen by Hipp's company, Hwaci.
The source code is released into the public domain rather than under a
conventional open-source license, so anyone may copy, modify, and distribute it,
including for commercial use, without restriction or attribution.

## Deployment and reach

The SQLite developers describe it as the most widely deployed database engine in
the world. Copies are bundled inside every Android and iOS device, the major web
browsers, popular operating systems, and many programming-language runtimes — for
example, Python ships an `sqlite3` module in its standard library. Because each
application embeds its own copy, the total number of running SQLite instances is
estimated in the many billions.

## Reliability and transactions

SQLite provides ACID transactions, so committed changes survive crashes and power
failures. Atomic commit is implemented through either a rollback journal or, in
write-ahead log (WAL) mode, by appending changes to a separate log file and later
checkpointing them back into the main database. WAL mode allows readers and a
single writer to proceed concurrently.

The project is known for an unusually thorough automated test suite. The tests
achieve 100% branch test coverage (measured as modified condition/decision
coverage) of the core engine, and the harness exercises the library against many
billions of test cases, including simulated I/O errors and out-of-memory
conditions. The maintainers have publicly committed to supporting SQLite through
at least the year 2050.
