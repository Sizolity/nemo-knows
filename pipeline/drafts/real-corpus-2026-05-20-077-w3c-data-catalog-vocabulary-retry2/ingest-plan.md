---
kind: topic
sources: [raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md]
status: draft
---

# Ingest Plan

## Source Summary
- The document is the W3C Data Catalog Vocabulary (DCAT) 3 specification, defining a standard RDF vocabulary for describing datasets and catalogs.
- It covers core classes (`Catalog`, `Dataset`, `Distribution`, `DataService`), namespace definitions, versioning strategies (distinguishing between series and versions), and integration with external vocabularies like Dublin Core, SKOS, and SPDX.
- Key sections include accessibility considerations, quality information modeling, conformance testing documentation, and security/privacy guidelines.

## Candidate Wiki Pages
- wiki/sources/dcat-3-specification.md — High-level overview of the DCAT 3 recommendation, evolution from DCAT 2, and core interoperability goals.
- wiki/concepts/dcatesp-normative-namespaces.md — Reference table for normative vs. non-normative prefixes (dcat, dcterms, foaf, skos, spdx) to prevent namespace collisions.
- wiki/topics/dataset-class-reference.md — Comprehensive guide on the `Dataset` class properties, including relations, themes, licenses, and versioning chains.
- wiki/concepts/distribution-vs-dataset-service.md — Clarification on modeling access methods: landing pages, download URLs, and API service endpoints (`DataService`).
- wiki/topics/versioning-strategies.md — Guide to distinguishing ordered series (`first`, `last`, `prev`) from version histories (`previousVersion`, `replaces`).
- wiki/concepts/quality-and-conformance-modeling.md — How to represent quality metrics, conformance test results (EARL), and validator agents using PROV-O patterns.
- wiki/topics/dataset-series-metadata.md — Best practices for using `dcat:DatasetSeries` to group related datasets with series-level temporal/spatial coverage.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that all candidate pages use valid paths under `wiki/sources/`, `wiki/concepts/`, or `wiki/topics/`.
- [ ] Ensure no nested directories are created in the wiki structure.
- [ ] Confirm that tool/API concepts are placed under `wiki/concepts/` or `wiki/topics/` as per rules.
- [ ] Check that repeated hints (e.g., versioning, access patterns) are consolidated into single pages rather than duplicated.
- [ ] Review that accessibility and security considerations are captured in the source summary without creating invalid page paths.
