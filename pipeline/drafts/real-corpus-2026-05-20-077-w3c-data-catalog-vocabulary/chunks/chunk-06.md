---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk covers Section 6 of the DCAT vocabulary specification, detailing classes and properties for distributions, data services, and taxonomies. It spans from the definition of `dcat:Distribution` through its specific properties (title, license, access URLs, byte size, resolution), to the new `dcat:DataService` class in DCAT 2, and finally introduces `skos:ConceptScheme`, `skos:Concept`, and `foaf:Agent` classes for catalog organization.

Local Summary
The text defines the `dcat:Distribution` class as a specific representation of a dataset, distinguishing between informational equivalence (e.g., different RDF serializations) and semantic equivalence (e.g., CSV vs. graphical representation). It details 20 properties covering metadata, access methods, licenses, formats, resolutions, and conformance. The chunk also introduces the `dcat:DataService` class for describing API endpoints or services that provide access to datasets. Finally, it outlines classes for taxonomic classification (`skos:ConceptScheme`, `skos:Concept`) and entity descriptions (`foaf:Person`, `foaf:Organization`).

Key Claims
- A `dcat:Distribution` represents a specific serialization of a dataset, which may differ in media type, format, or profile.
- All distributions of a single dataset should broadly contain the same data; distinct budget years are typically modeled as different datasets.
- `dcat:accessURL` is for services or landing pages (e.g., SPARQL endpoints), while `dcat:downloadURL` is preferred for direct file downloads via HTTP GET.
- In DCAT 2, `dcat:DataService` was added to describe collections of operations providing access to datasets.
- Properties like `dcat:spatialResolutionInMeters` and `dcat:temporalResolution` were added in DCAT 2 to summarize data granularity.
- `dcat:packageFormat` and `dcat:compressFormat` (DCAT 2 additions) handle grouped or compressed files, using IANA media types.
- `spdx:checksum` was added in DCAT 3 for file integrity verification.

Entities And Concepts
- **Class**: `dcat:Distribution` (RDF Class), representing a specific representation of a dataset.
- **Class**: `dcat:DataService` (RDF Class, added in DCAT 2), a collection of operations providing access to datasets.
- **Class**: `skos:ConceptScheme`, a knowledge organization system for themes/categories.
- **Class**: `skos:Concept`, a category or theme used to classify datasets.
- **Class**: `foaf:Person` and `foaf:Organization`, sub-classes of `foaf:Agent` for describing entities.
- **Property**: `dcat:accessURL`, URL for accessing a distribution (landing page, feed, endpoint).
- **Property**: `dcat:downloadURL`, URL for directly downloading a file.
- **Property**: `dcterms:title`, name given to the distribution.
- **Property**: `dcterms:license`, legal document under which the distribution is made available.
- **Property**: `dcterms:rights`, information about rights held in and over the distribution.
- **Property**: `odrl:hasPolicy`, an ODRL conformant policy expressing rights.
- **Property**: `dcat:byteSize`, size of the distribution in bytes.
- **Property**: `dcat:spatialResolutionInMeters`, minimum spatial separation resolvable (DCAT 2).
- **Property**: `dcat:temporalResolution`, minimum time period resolvable (DCAT 2).
- **Property**: `dcat:packageFormat`, format of a package containing multiple files (e.g., ZIP, TAR).
- **Property**: `dcat:compressFormat`, compression format used for the file.

Procedures And API Details
- **Access Logic**: Use `dcat:accessURL` for services (APIs, landing pages) and `dcat:downloadURL` for direct downloads. If only a landing page exists without known download URLs, duplicate the landing page URL as the access URL on the distribution.
- **Resolution Reporting**: For images/grids, spatial resolution corresponds to item spacing; for other spatial data, it indicates the smallest distance between items. Similarly, temporal resolution indicates the smallest time difference between items in a series.
- **Checksum Verification**: Use `spdx:checksum` (DCAT 3) to verify file contents have not changed, linked to the download URL.

Nuance Or Contradictions
- **License vs. Rights**: `dcterms:license` is a sub-property of `dcterms:rights`. While `dcterms:license` links to a license document, `dcterms:rights` allows linking to broader rights statements including attribution. Information on licenses/rights should primarily reside at the Distribution level; adding conflicting info at the Dataset level is discouraged.
- **Format vs. Media Type**: `dcat:mediaType` (IANA-defined) is preferred over `dcterms:format`. If the type is not IANA-defined, `dcterms:format` may be used. `dcat:compressFormat` and `dcat:packageFormat` are distinct from the base media type of the content inside.
- **Equivalence**: Distributions can be fully informationally equivalent (lossless transformations possible, e.g., Turtle to RDF/XML) or have different levels of fidelity (e.g., a CSV summary vs. the full raw data).

Candidate Wiki Hints
- **Page**: `dcat:Distribution` class properties and usage guidelines.
- **Page**: `dcat:DataService` definition and endpoint descriptions.
- **Page**: Guidelines on distinguishing between dataset versions and distinct datasets (e.g., budget years).
