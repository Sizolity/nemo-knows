---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
## Chunk Context
This chunk covers **Section 18: Accessibility Considerations** within the W3C Data Catalog Vocabulary (DCAT 3) specification, followed by change history logs documenting updates from Candidate Recommendation snapshots and working drafts between December 2020 and January 2024. It also details examples of data services (CSW, WFS, WMS) and compressed distributions.

## Local Summary
The primary focus is the addition of **Section 18: Accessibility Considerations** in the DCAT 3 vocabulary, introduced to address accessibility requirements for spatial data services. The text provides RDF examples for a thermal tolerance database (`GlobTherm`) and demonstrates how data services (Catalog Search Web, Web Feature Service, Web Map Service) are described using `dcterms:type` classifiers from INSPIRE. Additionally, it illustrates how distributions can be modeled with compression formats (GZIP, TAR) and packaging types. The latter part of the chunk lists specific editorial and technical changes made to the specification since various public working drafts, including property renames (e.g., `dcat:resource` replacing `dcterms:hasPart`), inverse property additions, and corrections to URI references.

## Key Claims
- **Section 18** was added to the DCAT 3 vocabulary specifically to include **Accessibility Considerations**.
- Data services can be classified using `dcterms:type` values from the **INSPIRE classification of spatial data service types** (e.g., "discovery", "download", "view").
- The property `dcat:resource` was introduced to link a catalog to a resource, replacing `dcterms:hasPart` from DCAT 2.
- New properties such as `dcat:inCatalog`, `dcat:seriesMember`, and `dcat:isVersionOf` (removed) were added or modified in various draft revisions.
- The property `dcat:theme` was explicitly defined as an OWL object property with its range dropped in later revisions.
- Examples show that a single dataset can be served via multiple endpoints (CSW, WFS, WMS) and described with distinct service types.

## Entities And Concepts
- **DCAT 3**: The Data Catalog Vocabulary specification being maintained by the W3C DXWG.
- **Section 18: Accessibility Considerations**: A newly added section addressing accessibility in data catalogs.
- **INSPIRE-SDST**: INSPIRE Spatial Data Service Type classification used for `dcterms:type` values.
- **GlobTherm**: A global database on thermal tolerances (example of a cataloged resource).
- **EEA-CSW-Endpoint**: An example of a Catalog Search Web service hosted by the European Environment Agency.
- **ga-courts**: A dataset regarding Australian judicial courts served via multiple protocols (MapServer, WFS, WMS).
- **dcat:resource**: The unified class for cataloged resources introduced in DCAT 3.
- **dcat:theme**: An OWL object property used for thematic classification.
- **Inverse Properties**: A new section added to define relationships like `dcat:inCatalog` and `dcat:seriesMember`.

## Procedures And API Details
- **Service Description**: Use `dcterms:type` with URIs from `<http://inspire.ec.europa.eu/metadata-codelist/SpatialDataServiceCategory/>` (e.g., `infoCatalogueService`, `download`, `view`).
- **Endpoint Identification**: Use `dcat:endpointURL` to point to the actual service endpoint and `dcat:endpointDescription` for the capabilities document.
- **Compression/Package Modeling**:
  - Use `dcat:compressFormat` (e.g., `application/gzip`) for compressed files.
  - Use `dcat:packageFormat` (e.g., `file-type/TAR`) for archived files.
  - Combine both properties if a file is both packaged and compressed (e.g., `.tar.gz`).
- **Property Replacement**: When modeling resources, use `dcat:resource` instead of `dcterms:hasPart`. Use `dcat:inCatalog` as the inverse of `dcat:resource`.

## Nuance Or Contradictions
- **Property Evolution**: The specification explicitly notes changes where properties were removed (e.g., `dcat:isVersionOf`, `dcat:next`) or renamed (e.g., `dcterms:hasPart` -> `dcat:resource`). Users must adhere to the latest version definitions.
- **Range Changes**: The range for `dcterms:byte size` changed from `xsd:decimal` to `xsd:nonNegativeInteger`, affecting data validation rules.
- **Inverse Properties**: Section 7 was added specifically to define inverse properties, removing previous implicit or conflicting usage of properties like `dcat:isVersionOf`.

## Candidate Wiki Hints
- **Page: DCAT 3 Accessibility**
  - *Summary*: Document the introduction of Section 18 and best practices for accessibility in data catalogs.
  - *Source Link*: raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md (Section 18)
- **Page: DCAT 3 Vocabulary Changes**
  - *Summary*: Track the evolution of properties and sections from the first public working draft to the Candidate Recommendation.
  - *Source Link*: raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md (Sections D, E, F, G, H, I)
