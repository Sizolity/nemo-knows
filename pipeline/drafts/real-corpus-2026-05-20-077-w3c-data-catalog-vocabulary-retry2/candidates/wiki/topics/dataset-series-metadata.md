---
title: Dataset Series Metadata
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

# Dataset Series Metadata

Dataset series represent an ordered grouping of related datasets within a data catalog. In the W3C Data Catalog Vocabulary (DCAT) 3 framework, these are modeled as first-class entities using the `dcat:DatasetSeries` class to distinguish between version history and logical series groupings.

## Core Concepts

A dataset series functions as a container for multiple datasets that share a common thematic or structural relationship. The vocabulary provides specific properties to manage the ordering and lifecycle of these series, distinguishing them from simple version chains.

### Ordering Properties

To define the sequence within a series, the following properties are utilized:

- `first`: Identifies the initial dataset in the series.
- `last`: Identifies the final dataset in the series.
- `prev`: Links to the immediately preceding dataset in the sequence.

These properties allow for the creation of navigable paths through time-series or iterative data releases.

## Relationship to Versioning

While DCAT 3 supports versioning chains via `previousVersion` and `currentVersion`, these are distinct from the ordered series defined by `first`, `last`, and `prev`. Implementations should carefully distinguish between a dataset that is simply an updated iteration of another (versioning) and a dataset that belongs to a logical collection or timeline (series).

## Integration with External Standards

The specification encourages the use of SPDX checksums for distributions within these series to ensure data integrity. Additionally, external vocabularies such as Dublin Core (`dcterms`) and SKOS are relied upon to provide extended semantics for agents and classifications associated with the series members.
