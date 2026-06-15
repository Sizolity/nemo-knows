---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
- Heading path: 6. Vocabulary specification > 6.8 Class: Distribution
- Line range: 2409–2920
- Covers properties of the `dcat:Distribution` class, definitions for `dcat:DataService`, and brief notes on `skos:ConceptScheme`, `skos:Concept`, and `foaf:Agent` classes.

Local Summary
- Defines the `dcat:Distribution` class as a specific representation (serialization) of a dataset, distinguishing it from the dataset itself.
- Lists 20 specific properties for describing distributions (e.g., title, license, access URLs, byte size).
- Clarifies the distinction between `accessURL` (service/location) and `downloadURL` (direct file link).
- Introduces new DCAT 2/3 properties like `spatialResolutionInMeters`, `temporalResolution`, `compressFormat`, `packageFormat`, and `checksum`.
- Defines the `dcat:DataService` class for describing API endpoints or services that serve datasets.
- Briefly introduces classes for categorization (`skos:ConceptScheme`, `skos:Concept`) and agents (`foaf:Person`, `foaf:Organization`).

Key Claims
- A dataset may have multiple distributions differing by format, resolution, or profile; they are not always fully informationally equivalent.
- `dcat:Distribution` implies general availability but does not specify the access method (download vs. API) without using `dcat:downloadURL` or `dcat:accessService`.
- License and rights information should be provided at the Distribution level; applying different license info to a Dataset than its Distributions creates legal conflicts.
- `dcat:mediaType` is preferred over generic `dcterms:format` when the media type is defined by IANA.
- `spatialResolutionInMeters` and `temporalResolution` are intended as single-value summaries; complex precision data belongs in the Data Quality Vocabulary.
- `checksum` (DCAT 3) links to a download URL to verify file integrity.

Entities And Concepts
- **dcat:Distribution**: A specific representation of a dataset (e.g., CSV, netCDF, JSON).
- **dcat:DataService**: A collection of operations providing access to datasets or processing functions.
- **skos:ConceptScheme**: A knowledge organization system for themes/categories.
- **skos:Concept**: A category used to describe datasets in a catalog.
- **foaf:Person / foaf:Organization**: Agents describing people or entities.

Procedures And API Details
- **Access URL (`dcat:accessURL`)**: Used for services, landing pages, or SPARQL endpoints. If only accessible via a landing page, duplicate the dataset's landing page URL here.
- **Download URL (`dcat:downloadURL`)**: Used for direct HTTP GET requests to downloadable files.
- **Service Linking**: Use `dcat:accessService` to link a distribution to a `dcat:DataService` description (e.g., OpenAPI, WSDL).
- **Resolution Properties**: Set `dcat:spatialResolutionInMeters` for grid/image spacing or minimum distance; set `dcat:temporalResolution` as an ISO 8601 duration for time-series spacing.

Nuance Or Contradictions
- **Format vs. Media Type**: `dcat:mediaType` is a sub-property of `dcterms:format`. Use `dcat:mediaType` for IANA-defined types; otherwise, use `dcterms:format`.
- **License vs. Rights**: `dcterms:license` links to a specific document (sub-property of rights). `dcterms:rights` allows broader statements including attribution. Both should generally be on the Distribution, not the Dataset, unless consistent.
- **Equivalence of Distributions**: While some distributions are losslessly transformable (e.g., RDF/XML vs. Turtle), others (e.g., CSV summary) may lose information but still represent distributions of the same dataset. Judgment on equivalence is application-specific.

Candidate Wiki Hints
- Page: DCAT Distribution Class Properties
  - Summary: Reference for all 20 properties of `dcat:Distribution`, including usage notes on license, resolution, and format.
- Page: DCAT Data Services
  - Summary: Overview of `dcat:DataService` and its relation to datasets via `dcat:servesDataset`.
