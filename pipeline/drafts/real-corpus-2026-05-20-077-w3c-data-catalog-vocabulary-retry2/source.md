---
title: W3C Data Catalog Vocabulary (DCAT) 3
kind: source
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## What It Is

The **W3C Data Catalog Vocabulary (DCAT)** is an RDF-based vocabulary designed to enable interoperable data cataloging across federated systems. Published as a W3C Recommendation in August 2024, DCAT 3 supersedes DCAT 1 and DCAT 2 while maintaining backward compatibility. It provides a standardized framework for describing datasets, distributions, services, and catalogs using common query mechanisms like SPARQL and HTML-RDFa.

## Summary

The specification defines a core set of classes and properties to model data catalog records. Key components include the **Catalog** (collection of metadata), **Dataset** (abstract collection of data), **Distribution** (specific serializations or access points like files or APIs), and **DataService** (service endpoints). DCAT 3 introduces explicit support for dataset series, versioning chains, checksums via SPDX, and qualified relations between resources and agents. It relies heavily on external vocabularies such as Dublin Core (`dcterms`), SKOS, FOAF, and Schema.org for extended semantics. The document concludes with guidelines on accessibility, security considerations, and conformance testing.

## Key Claims

- **Interoperability & Federation:** A core goal is to enable decentralized publishing and federated search across multiple organizations using uniform query mechanisms (SPARQL, HTML-RDFa).
- **Backward Compatibility:** DCAT 3 preserves definitions from previous versions, allowing existing implementations to remain compliant without mandatory modification.
- **Resource vs. Distribution Distinction:** A fundamental distinction is maintained between the abstract `Dataset` and its `Distribution`. Distributions can represent files (with download URLs) or services (via access URLs).
- **External Vocabulary Integration:** DCAT relies on external vocabularies for specific semantics. Changes to these external definitions do not affect DCAT conformance.
- **Non-Prescriptive Deployment:** Implementations may choose among various RDF serialization formats and exposure mechanisms; no specific deployment method is mandated.
- **Blank Node Avoidance:** The specification strongly discourages the use of blank nodes for DCAT instances to ensure global identifiers (IRIs) are used for collaborative annotation.
- **Dataset Series & Versioning:** The vocabulary distinguishes between version history (`previousVersion`, `currentVersion`) and ordered series (`first`, `last`, `prev`). It introduces `dcat:DatasetSeries` as a first-class entity for grouping related datasets.
- **Checksums:** DCAT 3 integrates SPDX checksums to represent file digests, linking the checksum to a specific distribution to ensure integrity.
- **Accessibility:** DCAT profiles should enforce alternative text for non-text resources to comply with guidelines like UNDERSTANDING-WCAG20.
- **Schema.org Alignment:** A recommended mapping exists between revised DCAT elements and Schema.org 3.4 to support tools like Google Dataset Search, though this mapping is non-normative.

## Suggested Links

- https://www.w3.org/TR/skos-reference/#notations
- https://www.w3.org/
