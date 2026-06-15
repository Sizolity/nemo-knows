---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk details the DCAT vocabulary specifications for classifying dataset types, describing catalog record metadata, handling datasets available via Web pages or services, and defining the core classes and properties for Catalogs and Cataloged Resources. It covers sections 5.5 through 6.4.8 of the source document.

Local Summary
The text outlines how to classify datasets using `dcterms:type` with references to external vocabularies like DCMI or DataCite, allowing multiple type assignments. It distinguishes between catalog records (`dcat:CatalogRecord`) and raw dataset instances. The chunk further categorizes access patterns: datasets behind a landing page only, those available for download alongside a landing page, and those distributed via specific services (e.g., APIs). Finally, it defines the RDF structure for `dcat:Catalog` and `dcat:Resource`, listing their specific and inherited properties such as `homepage`, `themes`, `resource`, `dataset`, `service`, and various metadata fields like `creator` and `release date`.

Key Claims
- Dataset classification should utilize well-governed, recognized vocabularies (e.g., DCMI Type Vocabulary, MARC Genre/Terms Scheme).
- Multiple classifications can coexist on a single dataset description using multiple `dcterms:type` properties.
- `dcat:CatalogRecord` is used to describe the registration of a resource within a catalog, distinct from the resource itself.
- Access patterns are modeled via `dcat:landingPage`, `dcat:distribution`, and `dcat:accessURL`.
- Services are represented using `dcat:DataService`, characterized by type, conformance, and endpoint descriptions.
- `dcat:Catalog` is a sub-class of `dcat:Dataset` in DCAT 2/3, enabling catalogs to be composed of other catalogs.
- `dcat:Resource` serves as the super-class for all cataloged items (datasets, services, catalogs).

Entities And Concepts
- **Vocabularies**: dcterms:type, DCMI Type Vocabulary, MARC Genre/Terms Scheme, ISO-19115-1 MD_Scope codes, DataCite resource types, Re3data content-types.
- **Classes**: `dcat:Dataset`, `dcat:CatalogRecord`, `dcat:Distribution`, `dcat:DataService`, `dcat:Catalog`, `dcat:Resource`.
- **Properties**: `homepage`, `themes` (themeTaxonomy), `resource`, `dataset`, `service`, `catalog`, `record`, `access rights`, `conforms to`, `contact point`, `creator`, `description`, `title`, `release date`, `update/modification date`.

Procedures And API Details
- **Multiple Typing**: To assign multiple types, use multiple triples with `dcterms:type` pointing to different URIs (e.g., combining DCMI and DataCite values).
- **Service Distribution**: A distribution linked to a service uses `dcat:accessService` pointing to the specific `dcat:DataService` instance. The service instance must define `dcat:endpointURL`, `dcterms:type`, and `dcterms:conformsTo`.
- **Landing Page Handling**: If data is only behind a Web page, define `dcat:landingPage` and a `dcat:Distribution` with the same access URL. If downloadable, use `dcat:downloadURL` on the distribution instance.

Nuance Or Contradictions
- In DCAT 1, `dcat:Catalog` was strictly for datasets; in DCAT 2/3, it is generalized as a sub-class of `dcat:Dataset` and can contain other catalogs.
- The domain of `dcat:contactPoint` was relaxed from `dcat:Dataset` to `dcat:Resource` in DCAT 2 to allow broader usage.
- Definitions for terms outside the DCAT namespace (e.g., from DCTERMS or FOAF) are provided for convenience but are not normative; authoritative definitions remain in their original specifications.

Candidate Wiki Hints
- **Page**: Classifying Dataset Types – Covers using `dcterms:type` and external vocabularies.
- **Page**: Catalog Record Metadata – Explains the difference between a dataset instance and its catalog record.
- **Page**: Access Patterns – Details modeling datasets behind landing pages, downloadable files, and service-based access.
- **Page**: DCAT Class Definitions – Defines `dcat:Catalog` and `dcat:Resource` hierarchies and properties.
