---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

Chunk Context
- Source: W3C Recommendation for the Data Catalog Vocabulary (DCAT) Version 3, published 22 August 2024.
- Scope: Defines an RDF vocabulary for publishing and aggregating metadata about datasets and data services on the Web.
- Evolution: DCAT 3 supersedes DCAT 2 while maintaining backward compatibility via the same namespace; it does not render DCAT 2 obsolete.

Local Summary
DCAT 3 is a W3C Recommendation designed to facilitate interoperability between data catalogs published on the Web. It provides a standard schema and vocabulary for describing datasets, data services, and their relationships. The document emphasizes that while implementations should adopt DCAT 3 for new features (like versioning and dataset series), existing deployments can continue using DCAT 2 without modification unless they wish to leverage the new capabilities.

Key Claims
- **Interoperability:** Enables consumption and aggregation of metadata from multiple catalogs, increasing discoverability.
- **Federated Search:** Supports searching for datasets across catalogs in multiple sites using a unified query mechanism.
- **Decentralized Publishing:** Facilitates a decentralized approach to publishing data catalogs.
- **Manifest for Preservation:** Aggregated DCAT metadata can serve as a manifest file in digital preservation processes.
- **Backward Compatibility:** DCAT 3 preserves the DCAT namespace and definition of previous terms; existing implementations do not need to upgrade unless using new features.

Entities And Concepts
- **DCAT (Data Catalog Vocabulary):** An RDF vocabulary for describing datasets and data services.
- **Namespace:** `http://www.w3.org/ns/dcat#` with suggested prefix `dcat`.
- **DCAT 1, DCAT 2, DCAT 3:** Historical versions of the vocabulary.
- **Distribution:** A manifestation or concrete access point for an abstract dataset (e.g., a download link).
- **Dataset Series:** A new class in DCAT 3 for representing series of related datasets.
- **Data Service:** An abstraction for describing services that provide access to data (e.g., APIs).
- **Profiles:** Extensions like DCAT-AP or HCLS-Dataset that build upon the base standard.

Procedures And API Details
- **Deployment Methods:** DCAT information can be deployed via SPARQL endpoints, embedded in HTML pages (using RDFa), or serialized as RDF/XML, N3, Turtle, or JSON-LD.
- **Implementation Reporting:** Publishers are invited to report implementations to the Dataset Exchange Working Group for analysis in an implementation report.
- **Feedback Mechanism:** Comments and issues should be submitted via GitHub pull requests/issues or email (`public-dxwg-comments@w3.org`).

Nuance Or Contradictions
- **Versioning Strategy:** DCAT 3 updates the specification but explicitly states it does not make DCAT 2 obsolete. Current deployments using only existing features (without overlapping with new features like versioning) remain in conformance without changes.
- **External Terms:** While DCAT incorporates stable terms from external vocabularies (e.g., `foaf:homepage`, `dcterms:title`), conformance to DCAT is based solely on terms defined within the DCAT specification itself. Changes to external definitions do not affect DCAT conformance.

Candidate Wiki Hints
- **Page Suggestion:** Create a dedicated page for "DCAT 3 Vocabulary" summarizing the schema, key classes (Catalog, Dataset, Distribution, Data Service), and the transition path from DCAT 2.
- **Concept Page:** A separate entry for "Data Catalog Interoperability" explaining how federated search and aggregation work using DCAT metadata.
