---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## Chunk Context
**Heading:** 18. Accessibility Considerations
**Line Range:** 5502–5860
**Scope:** This chunk covers the introduction of accessibility considerations, followed by detailed examples of data services (C.4), compressed/packaged distributions (C.5), and a comprehensive change history detailing revisions from the Candidate Recommendation Snapshot through multiple public working drafts up to May 2021.

## Local Summary
The text introduces **Section 18**, noting its addition in changes since the second public working draft of 4 May 2021 (Issue #1358). It provides concrete RDF examples for describing data services using DCAT properties like `dcterms:type`, `dcat:endpointDescription`, and `dcat:endpointURL`. The section on distributions demonstrates how to model compressed files (GZIP) and packaged archives (TAR, ZIP) using `dcat:compressFormat` and `dcat:packageFormat`. Finally, the change history logs specific editorial fixes, definition revisions (e.g., replacing "item" with "resource"), and removal of properties in response to implementation feedback.

## Key Claims
- **Section 18 (Accessibility Considerations)** was added to the specification following Issue #1358 during updates since May 2021.
- Data services can be described using classifiers such as `dcterms:type`, `dcterms:conformsTo`, and `dcat:endpointDescription` to provide progressive detail about a service.
- The actual endpoint of a service is identified via the property **`dcat:endpointURL`**.
- Distributions support modeling for compressed files (e.g., GZIP) using **`dcat:compressFormat`** and packaged archives (e.g., TAR) using **`dcat:packageFormat`**.
- Property **`dcat:resource`** was introduced to link a `dcat:Catalog` to a `dcat:Resource`, replacing the usage of `dcterms:hasPart` from DCAT 2.
- The property **`dcat:inCatalog`** serves as the inverse of `dcat:resource`.

## Entities And Concepts
- **Accessibility Considerations**: A new section addressing accessibility in data catalogs.
- **Data Services**: Described using DCAT, including examples for the European Environment Agency (EEA) and Geoscience Australia.
- **Distributions**: Models for downloadable files, compressed archives (GZIP), and packaged files (TAR).
- **DCAT Properties**:
  - `dcat:endpointURL`: Identifies the actual endpoint of a service.
  - `dcat:compressFormat`: Specifies compression types (e.g., GZIP).
  - `dcat:packageFormat`: Specifies packaging formats (e.g., TAR, ZIP).
  - `dcat:resource`: Links catalog to resource instances.
  - `dcat:inCatalog`: Inverse property for resources.
- **INSPIRE Classification**: Referenced for spatial data service types (`discovery`, `download`, `view`).

## Procedures And API Details
- **Modeling a Data Service**: Use `rdf:type dcat:DataService` and set `dcterms:type` (e.g., INSPIRE codes) and `dcat:endpointURL`.
- **Modeling Distributions**:
  - For compressed files: Set `dcat:downloadURL` and add `dcat:compressFormat`.
  - For packaged files: Set `dcat:packageFormat`.
  - Combined compression/packaging: Use both properties.
- **Service Types**: Examples use `dcterms:type` values from the INSPIRE Spatial Data Service Category list (e.g., `infoCatalogueService`, `download`, `view`).

## Nuance Or Contradictions
- **Terminology Shift**: The term "item" was replaced with "resource" when referring to instances of `dcat:Resource` to ensure consistent terminology throughout the document.
- **Property Replacement**: `dcterms:hasPart` (from DCAT 2) is no longer recommended for linking a catalog to a resource; `dcat:resource` is now defined for this purpose.
- **Inverse Properties**: Section 7 defines inverse properties like `dcat:inCatalog`, while older properties like `dcat:isVersionOf` and `dcterms:isReplacedBy` were removed in the May 2021 draft.

## Candidate Wiki Hints
- **Page: DCAT Accessibility**
  - *Content*: Summarize Section 18, explaining why accessibility considerations were added and how they integrate into catalog metadata standards.
- **Page: Data Service Endpoints**
  - *Content*: Document the use of `dcat:endpointURL` and related descriptors (`dcterms:conformsTo`, `dcat:endpointDescription`) for defining service endpoints in DCAT.
- **Page: Distribution Packaging Formats**
  - *Content*: Explain how to model downloadable, compressed (GZIP), and packaged (TAR/ZIP) distributions using specific DCAT properties.
