---
title: Dataset Class Reference
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

# Dataset Class Reference

The W3C Data Catalog Vocabulary (DCAT) provides a standardized framework for describing datasets, distributions, services, and catalogs using common query mechanisms like SPARQL and HTML-RDFa. DCAT 3 supersedes previous versions while maintaining backward compatibility, enabling interoperable data cataloging across federated systems.

## Core Classes

The specification defines a core set of classes to model data catalog records:

- **Catalog**: Represents a collection of metadata.
- **Dataset**: An abstract collection of data.
- **Distribution**: Specific serializations or access points, such as files or APIs.
- **DataService**: Service endpoints.

## Series and Versioning

DCAT 3 distinguishes between version history and ordered series:

- **Version History**: Managed via `previousVersion` and `currentVersion`.
- **Ordered Series**: Managed via `first`, `last`, and `prev`.
- **DatasetSeries**: Introduced as a first-class entity (`dcat:DatasetSeries`) for grouping related datasets.

## Integrity and Access

- **Checksums**: The vocabulary integrates SPDX checksums to represent file digests, linking the checksum to a specific distribution to ensure integrity.
- **Distribution vs. Service**: Distributions can represent files (with download URLs) or services (via access URLs).

## External Dependencies

DCAT relies heavily on external vocabularies for extended semantics:

- Dublin Core (`dcterms`)
- SKOS
- FOAF
- Schema.org

Changes to these external definitions do not affect DCAT conformance. A recommended mapping exists between revised DCAT elements and Schema.org 3.4 to support tools like Google Dataset Search, though this mapping is non-normative.
