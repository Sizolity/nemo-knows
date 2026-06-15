---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
### Chunk Context
This chunk covers sections 5.5 through 6.4 of the W3C Data Catalog Vocabulary (DCAT) specification. It details how to classify dataset types using external vocabularies, describe catalog records, represent datasets available via web pages or services, and defines the core RDF classes (`dcat:Catalog`, `dcat:Resource`) and their properties (e.g., title, creator, access rights).

### Local Summary
The text outlines methods for assigning dataset genres using recognized external vocabularies like DCMI Type or DataCite. It distinguishes between datasets available only via a landing page versus those downloadable directly or accessed through specific services (APIs). The section transitions into the formal vocabulary specification, defining RDF representations, usage of external vocabularies (VoID, PROV-O), and detailed definitions for classes like `dcat:Catalog` and `dcat:Resource`, along with their associated properties such as `homepage`, `themes`, and `accessRights`.

### Key Claims
- Dataset types should be indicated using the `dcterms:type` property, preferably drawing from well-governed vocabularies (DCMI Type, MARC Genre, ISO-19115, DataCite, Re3data).
- Multiple classifications can coexist for a single dataset instance.
- Catalog records (`dcat:CatalogRecord`) allow publishers to track metadata about the registration of datasets distinct from the dataset itself.
- Datasets available only via a web page are represented using `dcat:landingPage` without defining a separate distribution, whereas downloadable files require a `dcat:distribution` with `dcat:downloadURL`.
- Services (e.g., APIs) are linked to distributions using `dcat:accessService` and characterized by `dcterms:type` (e.g., INSPIRE spatial data service types) and `dcterms:conformsTo`.
- The `dcat:Catalog` class was generalized in DCAT 2 to include datasets, services, and other catalogs as sub-resources.
- Properties like `access rights`, `conforms to`, and `creator` were added or expanded in DCAT 2 to address data citation and broader resource management needs.

### Entities And Concepts
- **dcterms:type**: Property for indicating the genre/type of a dataset.
- **Vocabularies**: DCMI Type Vocabulary, MARC Genre/Terms Scheme, ISO-19115-1 MD_Scope codes, DataCite resource types, Re3data content-types.
- **dcat:CatalogRecord**: Class for metadata describing the registration of a resource within a catalog.
- **dcat:landingPage**: Property pointing to the canonical Web page for accessing a dataset.
- **dcat:distribution**: Represents an instance of access (download or service).
- **dcat:accessURL**: Generic URL for accessing a distribution.
- **dcat:downloadURL**: Specific property for direct download links.
- **dcat:accessService**: Links a distribution to a specific data service.
- **dcat:endpointURL**: URL of the service endpoint.
- **dcat:endpointDescription**: Link to detailed parameter descriptions for an endpoint.
- **dcat:Catalog**: Class representing a curated collection of metadata; sub-class of `dcat:Dataset`.
- **dcat:Resource**: Super-class for all cataloged resources (`dcat:Dataset`, `dcat:DataService`, etc.).
- **External Vocabularies**: VoID (statistics), PROV-O (provenance/workflow), Organization Ontology (agents).
- **Properties of dcat:Resource**: `accessRights`, `conformsTo`, `contactPoint`, `creator`, `description`, `title`, `release date` (`dcterms:issued`), `update/modification date` (`dcterms:modified`).

### Procedures And API Details
1.  **Classifying Types**: Assign values to `dcterms:type` from an external vocabulary URI (e.g., `<http://purl.org/dc/dcmitype/Dataset>`). Multiple URIs can be assigned if the resource fits multiple genres.
2.  **Representing Catalog Records**: Use `dcat:record` to link a catalog to a record instance (`dcat:CatalogRecord`). The record uses `foaf:primaryTopic` to point to the actual dataset and `dcterms:issued` for the date the record was added.
3.  **Accessing via Web Page**: Define the dataset with `a dcat:Dataset` and assign a `dcat:landingPage`. If the data is not downloadable directly, do not define a separate distribution instance.
4.  **Mixing Access Methods**: For datasets available both via a landing page and download, define one `dcat:distribution` instance using `dcat:downloadURL`. The landing page remains the primary access point unless specified otherwise.
5.  **Service-Based Distribution**: Create a `dcat:DataService` instance (e.g., `ex:table-service-005`). Link it to a distribution via `dcat:accessService`. Specify the API definition using `dcterms:conformsTo`, the service type using `dcterms:type`, and link parameters via `dcat:endpointDescription`. A single service can serve multiple datasets (e.g., `dcat:servesDataset`).
6.  **Defining Resource Properties**: When instantiating a resource, populate standard properties like `dcterms:title`, `dcterms:description`, `dcterms:creator` (using `foaf:Agent`), and `dcterms:accessRights`.

### Nuance Or Contradictions
- **Catalog Evolution**: The scope of `dcat:Catalog` changed significantly between DCAT 1 and DCAT 2. In DCAT 1, it was limited to catalogs of datasets. In DCAT 2, it became a sub-class of `dcat:Dataset`, allowing for nested catalogs and the inclusion of data services as resources within a catalog.
- **Property Domains**: Properties defined by DCAT generally do not have specified domains to remain flexible for use with any resource type, unlike specific subclasses like `dcat:Catalog` which have specific domain/range constraints.
- **External Definitions**: Definitions and domains for terms outside the DCAT namespace (e.g., those from DCTERMS or FOAF) are provided in the spec only for convenience and are not normative; authoritative definitions lie in their original specifications.
- **Access Rights vs. Licenses**: While `dcterms:accessRights` is used for security status or access permissions, licenses are discussed separately (referenced in section 9), implying a distinction between who *can* access (rights) and under what terms (license).

### Candidate Wiki Hints
- **Page: DCAT Vocabulary Specification** – Documenting the RDF structure, external vocabularies integration, and class definitions.
- **Page: Dataset Access Patterns** – Comparing landing page-only, direct download, and service-based access representations in DCAT.
- **Page: Catalog vs. Resource Classes** – Explaining the hierarchy between `dcat:Catalog`, `dcat:Resource`, and specific resource types like `dcat:Dataset`.
