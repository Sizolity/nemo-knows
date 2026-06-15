## chunk-01

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

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk covers Sections 3 through 5 of the DCAT vocabulary specification. It defines normative and non-normative namespaces, conformance criteria for data catalogs, and provides a high-level overview of the core classes (dcat:Catalog, dcat:Dataset, etc.), RDF best practices regarding blank nodes, and illustrative examples of catalog descriptions using Turtle syntax.

Local Summary
The text establishes the technical foundation for implementing DCAT by listing the required URIs for prefixes (e.g., `dcat`, `dcterms`, `foaf`) and distinguishing between normative definitions and non-normative examples. It details how to model a data catalog as an RDF graph, emphasizing the distinction between the resource itself (`dcat:Resource`) and its metadata record (`dcat:CatalogRecord`). The section on RDF considerations strongly advises against using blank nodes for main classes to support Linked Data interoperability. Finally, it provides concrete Turtle examples of defining a catalog with multilingual labels, linking datasets to distributions, and applying thematic classification via SKOS.

Key Claims
- DCAT is an RDF vocabulary based around seven main classes designed to represent data catalogs containing datasets and data services.
- Conformance requires that metadata about the catalog, resources, and distributions be expressed using appropriate DCAT classes and properties in an RDF description.
- Blank nodes are generally discouraged for instances of DCAT main classes because they limit collaborative annotation and global identifier resolution in a Linked Data context.
- A `dcat:Dataset` is defined as a collection of data curated by a single agent, available in one or more serializations, whereas a `dcat:Distribution` represents an accessible form (e.g., a downloadable file) of that dataset.
- The scope of DCAT currently focuses on datasets and data services; extending to other resource types requires defining sub-classes of `dcat:Resource` in a profile.

Entities And Concepts
- **DCAT Vocabulary**: An RDF vocabulary for representing data catalogs.
- **Normative Namespaces**: Includes `adms`, `dc`, `dcat`, `dcterms`, `dctype`, `foaf`, `locn`, `odrl`, `owl`, `prov`, `rdf`, `rdfs`, `skos`, `spdx`, `time`, `vcard`, `xsd`.
- **Non-normative Namespaces**: Includes `dqv`, `earl`, `geosparql`, `oa`, `pav`, `sdmx-attribute`, `sdo`, `xhv`.
- **Core Classes**:
  - `dcat:Catalog`: Represents a collection of metadata records.
  - `dcat:Resource`: Parent class for datasets, services, and other resources; acts as an extension point.
  - `dcat:Dataset`: A collection of data published or curated by a single agent.
  - `dcat:DataService`: A collection of operations (APIs) providing access to datasets or processing functions.
  - `dcat:Distribution`: An accessible form of a dataset (e.g., file download).
  - `dcat:DatasetSeries`: A collection of separately published datasets sharing characteristics.
  - `dcat:CatalogRecord`: Metadata record describing registration information (optional).
- **Concept Schemes**: Used with SKOS to classify datasets thematically (e.g., accountability, transparency).

Procedures And API Details
- **Namespace Declaration**: Use standard prefixes mapped to specific IRIs (e.g., `<http://www.w3.org/ns/dcat#>` for `dcat`).
- **Catalog Definition**: Define a catalog using `a dcat:Catalog`, providing titles in multiple languages via `dcterms:title` and `rdfs:label`, and linking to a homepage.
- **Dataset Modeling**: Describe datasets with properties like `dcat:keyword`, `dcterms:creator`, temporal intervals (`dcat:startDate`, `dcat:endDate`), spatial coverage (`dcterms:spatial`), and contact points (`dcat:contactPoint`).
- **Distribution Linking**: Connect a dataset to its files using the `dcat:distribution` property, specifying the download URL (`dcat:downloadURL`) and media type (`dcat:mediaType`).
- **Thematic Classification**: Create a `skos:ConceptScheme` (e.g., `ex:themes`) and link datasets to specific concepts within it using `dcat:theme`.

Nuance Or Contradictions
- The document distinguishes between the resource itself (`dcat:Resource`) and its metadata record (`dcat:CatalogRecord`), noting that using `CatalogRecord` is optional but useful for capturing provenance.
- While DCAT 1 was limited to datasets, DCAT 2 explicitly includes data services as members of a catalog, allowing for more complex API-based access patterns.
- Blank nodes are technically allowed in RDF but are discouraged in this specification to maintain the benefits of the Linked Data model (global identifiers and open-world assumptions).

Candidate Wiki Hints
- **DCAT Core Classes**: A page explaining the seven main classes (`Catalog`, `Dataset`, `Distribution`, etc.) and their relationships.
- **Namespace Reference**: A cheat sheet for DCAT prefixes and IRIs, separating normative from non-normative usage.
- **Conformance Guide**: Notes on what makes a catalog "DCAT-compliant" versus how to create a profile with additional constraints.
- **RDF Best Practices for DCAT**: Guidelines on avoiding blank nodes and using Turtle syntax for DCAT examples.

## chunk-03

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

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk details the vocabulary properties for cataloging resources within DCAT specifications (DCAT 1, 2, and 3). It covers metadata attributes such as language, publisher, identifiers, themes, types, relations (generic and qualified), keywords, landing pages, rights/licenses, parts/policies, citations, versioning chains, lifecycle status, and series ordering. The text includes specific notes on domain relaxations from DCAT 1 to DCAT 2, the introduction of new properties in DCAT 3 (e.g., version chains), and distinctions between generic relations and specific ones like distribution or hasPart.

Local Summary
The section defines RDF properties for describing various facets of a cataloged resource. Key focus areas include linguistic identification using ISO 639 codes, entity attribution (publisher/creator) via `dcterms:publisher` and `prov:qualifiedAttribution`, categorization via `dcat:theme` and `dcterms:type`, and complex relationship modeling using `dcterms:relation` and the qualified relation pattern (`dcat:qualifiedRelation`). The chunk also addresses resource lifecycle management through properties for versions (`dcat:hasVersion`, `dcat:previousVersion`), replacement chains (`dcterms:replaces`), and series membership (`dcat:first`).

Key Claims
- Language identification should prioritize ISO 639-1 codes; if unavailable, ISO 639-2 is recommended. BCP47 tags are also noted as evolving standards for natural language in linked data.
- The domain of `dcat:theme` was relaxed in DCAT 2 to allow use beyond just datasets, unlike DCAT 1 where it was limited to `dcat:Dataset`.
- `dcterms:type` is a sub-property of `dc:type` and should draw from controlled vocabularies like DCMI Type or ISO 19115, though not all terms (e.g., Event, PhysicalObject) are strictly suitable for datasets.
- Generic relationships between resources should use `dcterms:relation`, but specific semantic links (e.g., distribution, isVersionOf) should take precedence to avoid over-generalization.
- DCAT 3 introduces specific properties for versioning lineages (`dcat:hasVersion`, `dcat:previousVersion`) and lifecycle status (`adms:status`).
- The qualified relation pattern (`dcat:qualifiedRelation`, `prov:qualifiedAttribution`) is used when the relationship nature is known but does not match standard DCTERMS or PROV-O properties.

Entities And Concepts
- **Vocabulary Standards**: DCAT 1, DCAT 2, DCAT 3, ISO 639-1, ISO 639-2, BCP47, DCMI Type, ISO 19115-1, DataCite.
- **Classes and Types**: `dcterms:LinguisticSystem`, `foaf:Agent`, `rdfs:Literal`, `skos:Concept`, `owl:ObjectProperty`.
- **Properties**: `dcterms:language`, `dcterms:publisher`, `dcterms:identifier`, `dcat:theme`, `dcterms:type`, `dcterms:relation`, `dcat:qualifiedRelation`, `dcat:keyword`, `dcat:landingPage`, `prov:qualifiedAttribution`, `dcterms:license`, `dcterms:rights`, `odrl:hasPolicy`, `dcterms:hasPart`, `dcat:isReferencedBy`, `dcat:previousVersion`, `dcat:hasVersion`, `dcat:currentVersion` (as `dcat:hasCurrentVersion`), `dcterms:replaces`, `dcat:version`, `adms:versionNotes`, `adms:status`, `dcat:first`.
- **Concepts**: Qualified relation pattern, version chains, resource lineage, dataset series.

Procedures And API Details
- **Language Usage**: If a resource is available in multiple languages, repeat the property. Values for catalog members override catalog-level values if conflicting. For separate language representations, define distinct `dcat:Distribution` instances.
- **Theme Organization**: Themes should be organized within a `skos:ConceptScheme`, `skos:Collection`, or similar ontology describing categories and relations.
- **Identifier Usage**: The identifier is a text string assigned for unambiguous reference; it may be part of the resource IRI but should still be explicitly represented.
- **Versioning Strategy**: Use `dcat:hasVersion` to link an abstract resource to versioned snapshots. Use `dcat:previousVersion` and `dcat:hasCurrentVersion` to specify lineage steps. Use `dcterms:replaces` for superseded resources.
- **Qualified Relations**: When linking to another resource where the relationship is known but not covered by standard properties, use `dcat:qualifiedRelation`. This entails that the context resource is a member of `prov:Entity`.

Nuance Or Contradictions
- DCAT 1 vs. DCAT 2: The domain of `dcat:theme` and `dcat:keyword` was restricted in DCAT 1 to `dcat:Dataset`, which limited their utility. DCAT 2 relaxed these domains to allow broader application.
- Property Semantics: `dcterms:type` is a sub-property of `dc:type`, but its range in DCAT 3 is adjusted to `rdfs:Class` for formal consistency, preventing automatic inference of objects as `skos:Concept`.
- Versioning Scope: `dcat:hasVersion` is more specific than `dcterms:hasVersion`; the former is limited to revisions during a resource's lifecycle (e.g., dataset updates), while the latter includes broader notions like editions or adaptations.

Candidate Wiki Hints
- **Data Catalog Vocabulary**: A structured set of RDF properties for describing resources in data catalogs, evolving from DCAT 1 through DCAT 3.
- **Qualified Relations**: A pattern using `prov:qualifiedInfluence` and `dcat:Relationship` to describe specific relationships that do not fit standard DCTERMS or PROV-O properties.
- **Versioning in DCAT**: Properties for managing resource versions, including lineage chains (`hasVersion`, `previousVersion`) and lifecycle status.
- **Dataset Series**: Use of `dcat:first` and related properties to define ordered collections or series of resources.

## chunk-05

---
title: Chunk NN Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
## Chunk Context
This chunk covers sections 6.4.32 through 6.7 of the DCAT vocabulary specification. It details properties for ordered resource collections (`last`, `previous`), defines the optional `Catalog Record` class to distinguish between dataset metadata and catalog entry metadata, specifies core `Dataset` properties including distribution, frequency, spatial/temporal coverage, and provenance, and introduces the new `Dataset Series` class in DCAT 3.

## Local Summary
The text outlines specific vocabulary extensions for handling resource series (first/last/previous), clarifies the distinction between a dataset's intrinsic metadata and its catalog entry via the `Catalog Record` class, lists standard properties for datasets including spatial/temporal resolution and provenance generation, and defines the `Dataset Series` class to group related datasets sharing characteristics like time or theme.

## Key Claims
- The `last` and `previous` properties were added in DCAT 3 to manage ordered collections of resources within a series.
- The `dcat:CatalogRecord` class is optional and serves to distinguish metadata about the catalog entry itself (e.g., listing date) from metadata about the underlying dataset (e.g., release date).
- Information regarding licenses and rights should primarily be provided at the `Distribution` level, though it may also appear on a `Dataset` if necessary, but differing values between them should be avoided to prevent legal conflicts.
- The `dcat:DatasetSeries` class was added in DCAT 3 to represent collections of separately published datasets that share characteristics (e.g., time series, map series).

## Entities And Concepts
- **Classes**: `dcat:CatalogRecord`, `dcat:Dataset`, `dcat:DatasetSeries`.
- **Properties**: `dcat:last`, `dcat:prev` (previous), `dcterms:title`, `dcterms:description`, `dcterms:issued` (listing date), `dcterms:modified` (update/modification date), `foaf:primaryTopic`, `dcterms:conformsTo`, `dcat:distribution`, `dcterms:accrualPeriodicity` (frequency), `dcat:inSeries`, `dcterms:spatial`, `dcat:spatialResolutionInMeters`, `dcterms:temporal`, `dcat:temporalResolution`, `prov:wasGeneratedBy`.
- **Concepts**: Dataset Series, Catalog Entry vs. Dataset Metadata, Spatial Resolution, Temporal Resolution, Provenance Activity.

## Procedures And API Details
- **Usage of `last` and `previous`**: These properties link a resource to its predecessor or successor in an ordered series (`dcat:DatasetSeries`). They are distinct from versioning properties like `previousVersion`.
- **Catalog Record vs. Dataset**: When listing dates differ (dataset publication vs. catalog addition), the date of the catalog entry should be specified on the `CatalogRecord`, while the dataset's original availability date remains on the `Dataset`.
- **Spatial Resolution Encoding**: Values for `dcat:spatialResolutionInMeters` must be typed as `xsd:decimal`. Note that JSON-LD may convert numbers to `xsd:double` or `xsd:integer`; validation schemas should accept `xsd:integer` as valid for this property.
- **Temporal Resolution vs. Frequency**: `dcterms:accrualPeriodicity` describes the rate at which the dataset is published/updated, whereas `dcat:temporalResolution` indicates the minimum time period resolvable within the data (e.g., spacing of items in a time series).

## Nuance Or Contradictions
- **Sub-class Removal**: In DCAT 1, `dcat:Dataset` was a sub-class of `dctype:Dataset`. This relationship was removed in DCAT 2 to broaden the scope of `dcat:Dataset` to include multimedia and text from the DCMI Types vocabulary.
- **Datatype Nuance**: While `xsd:integer` is semantically valid for `dcat:spatialResolutionInMeters` (as it derives from `xsd:decimal`), strict validation tools like SHACL or ShEx may treat them as distinct datatypes, requiring explicit datatype handling in schema definitions.
- **Inverse Properties**: The document references the use of inverse properties (Section 7) and specific guidance on series usage (Section 12), implying that relationships like `inSeries` might be traversed inversely depending on implementation patterns.

## Candidate Wiki Hints
- **Page: DCAT Dataset Series** - Covers the definition, usage scenarios (time series, map series), and related properties (`first`, `last`, `previous`, `inSeries`).
- **Page: DCAT Catalog Record** - Explains the optional nature of this class and the distinction between catalog entry metadata and dataset metadata.
- **Page: DCAT Spatial and Temporal Properties** - Details resolution vs. coverage, frequency vs. accrual periodicity, and encoding standards for spatial/temporal data.

## chunk-06

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

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

### Chunk Context
This chunk details the vocabulary specifications for **DCAT 2** and **DCAT 3**, covering relationships, roles, temporal intervals, spatial locations, and checksums. It further defines the use of inverse properties, dereferenceable identifiers (including HTTP IRIs, proxies, and legacy IDs), and strategies for indicating common identifier types using RDF datatypes or OWL datatypes when not HTTP-dereferenceable.

### Local Summary
The specification introduces specific classes (`Relationship`, `Role`, `PeriodOfTime`, `Location`, `Checksum`) to characterize associations between datasets that are not covered by standard DCTERMS or PROV-O properties. It defines temporal and spatial coverage using ISO 8601 literals or GeoSPARQL geometries. The section on identifiers emphasizes the use of persistent HTTP IRIs (`dcterms:identifier`) while allowing for proxy dereferenceable IRIs, legacy IDs, and locally minted identifiers via `adms:identifier`. It concludes with guidelines for encoding non-HTTP identifiers using custom datatypes.

### Key Claims
- **Relationship Class**: Added in DCAT 2 to attach additional information to relationships between DCAT resources where standard properties (DCTERMS/PROV-O) are insufficient.
- **Role Class**: Added in DCAT 2 to specify the function of a resource or agent with respect to another, recommending controlled vocabularies like ISO-19115.
- **Temporal Properties**: `dcat:startDate` and `dcat:endDate` use ISO 8601 strings; `time:hasBeginning`/`time:hasEnd` use `time:Instant` for non-Gregorian or numeric time positions.
- **Spatial Properties**: `locn:geometry` is preferred for extensive geometries, while `dcat:bbox` and `dcat:centroid` are used for bounding boxes and centers respectively.
- **Checksum Class**: Added in DCAT 3; uses `spdx:algorithm` and `spdx:checksumValue` to verify file integrity (e.g., MD5, SHA-256).
- **Inverse Properties**: Supported only in addition to standard properties, never as replacements, to ensure interoperability without OWL reasoning.
- **Identifiers**: Prefer persistent HTTP IRIs; use `owl:sameAs` for dereferenceable IDs returning RDF/OWL descriptions; distinguish primary vs. alternative identifiers based on application context (e.g., DCAT-AP).

### Entities And Concepts
- **dcat:Relationship**: An association class for relationships between datasets/resources.
- **dcat:hadRole**: Specifies the function of an entity/agent with respect to another.
- **dcat:Role**: A subclass of `skos:Concept` representing a role in attribution or relationships.
- **dcterms:PeriodOfTime**: An interval of time defined by start and end.
- **locn:geometry**: Associates a spatial thing with a geometry (WKT, GeoSPARQL).
- **spdx:Checksum**: Represents checksum algorithms and values for file integrity.
- **adms:identifier**: Used for locally minted or external identifiers (DOI, ORCID).
- **owl:sameAs**: Links resources that share the same identity when dereferenced.

### Procedures And API Details
- **Temporal Encoding**: Use `xsd:gYear`, `xsd:date`, or `xsd:dateTime` for `dcat:startDate`/`endDate`; use `time:Instant` for `time:hasBeginning`/`time:hasEnd`.
- **Spatial Encoding**: Encode geometries as literals (WKT) or classes (`geosparql:Geometry`). WKT supports non-WGS84 coordinate systems.
- **Identifier Types**:
  - HTTP IRIs: Use `dcterms:identifier` with `xsd:anyURI`.
  - Non-HTTP: Use RDF datatypes (`ex:type`) or OWL datatypes.
  - DOI Encoding: Use full URL form (e.g., `https://doi.org/10.xxxx/xxxxx/`).
- **Agency Representation**: Use `adms:schemaAgency` for the authority defining the scheme; use `dcterms:creator` if the agency has an IRI.

### Nuance Or Contradictions
- **Inverse Properties**: The spec intentionally excludes inverses in the core vocabulary to support systems not using OWL reasoning, but allows them as additions (e.g., `dcat:prev`/`dcat:next`).
- **Identifier Philosophy**: DOIs abstract sub-spaces from the registering organization to ensure stability if responsibility changes; registrants should not represent the assigning authority in the identifier value itself.
- **Primary vs. Alternative IDs**: Distinguishing between primary and legacy identifiers is application-specific and better addressed in DCAT profiles (e.g., DCAT-AP) rather than a general mandate.

### Candidate Wiki Hints
- **Page: DCAT Relationship Class** - Covers `dcat:Relationship`, `hadRole`, and usage examples from qualified relations.
- **Page: DCAT Temporal Properties** - Documents `PeriodOfTime`, `startDate`, `endDate`, and handling of non-Gregorian time scales.
- **Page: DCAT Spatial Properties** - Details geometry, bounding box, and centroid specifications using GeoSPARQL/WKT.
- **Page: DCAT Checksums** - Explains the `spdx:Checksum` class for file integrity verification in DCAT 3.
- **Page: Managing Identifiers in DCAT** - Guides on HTTP IRIs, proxy IDs, legacy identifiers, and using `adms:identifier` with schema agencies.

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

### Chunk Context
**Source Path:** `raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md`
**Chunk Range:** Lines 3413–3879
**Covered Sections:**
- Section 9: License and rights statements (including ODRL policies)
- Section 10: Time and space (Temporal properties, Spatial properties)
- Section 11: Versioning (Relationships between versions, Version chains and hierarchies)

### Local Summary
This chunk details how to express legal conditions, temporal metadata, spatial boundaries, and version histories for data cataloged resources using DCAT and related vocabularies. It distinguishes between licensing, access rights, and copyright statements, recommending specific properties from DCTERMS and ODRL. It further defines five temporal properties (release time, revision time, update schedule, temporal resolution, and temporal extent) and two spatial properties (spatial resolution and spatial extent), illustrating usage with WKT geometries and bounding boxes. Finally, it outlines a versioning model supporting version chains and hierarchies using `dcat:previousVersion`, `dcat:hasVersion`, and `dcat:hasCurrentVersion`.

### Key Claims
- Selecting conditions for resource access is complex; legal advice is required before implementation.
- Three distinct situations exist for rights statements: explicit 'license', access rights only, and other rights (e.g., copyright).
- DCAT recommends using `dcterms:rights`, `dcterms:license`, and `dcterms:accessRights` to address the three scenarios above.
- For interoperability, canonical IRIs of well-known licenses (e.g., Creative Commons) are preferred.
- Access rights can be expressed via code lists/taxonomies (e.g., [EUV-AR], Eprints Access Rights Vocabulary).
- The Open Data Rights Statement Vocabulary (ODRS) offers a more sophisticated approach extending DCTERMS for copyright notices.
- When using ODRL policies, the `odrl:hasPolicy` property links the resource to the policy.
- Five temporal properties are supported: release time (`dcterms:issued`), revision time (`dcterms:modified`), update schedule (`dcterms:accrualPeriodicity`), temporal resolution (`dcat:temporalResolution`), and temporal extent (`dcterms:temporal`).
- Two spatial properties are supported: minimum separation (`dcat:spatialResolutionInMeters`) and spatial extent (`dcterms:spatial`).
- Spatial geometries (for `locn:geometry`, `dcat:bbox`, `dcat:centroid`) should be specified using WKT; omitting CRS implies default CRS84.
- DCAT versioning focuses on revisions resulting from a resource's life-cycle, building upon [PAV], [DCTERMS], and [OWL2-OVERVIEW].
- Version chains are navigated backward via `dcat:previousVersion`, while hierarchies link abstract resources to versions via `dcat:hasVersion`.

### Entities And Concepts
- **Rights Statements:**
  - `dcterms:rights`: General property for rights statements not covered by license or access rights.
  - `dcterms:license`: Refers specifically to licenses (e.g., CC-BY 4.0).
  - `dcterms:accessRights`: Expresses access restrictions (public vs. authorized).
  - ODRL (Open Digital Rights Language): Policy expression language for permissions, prohibitions, and obligations.
  - ODRS (Open Data Rights Statement Vocabulary): Extends DCTERMS for sophisticated rights specification.
- **Temporal Properties:**
  - `dcterms:issued`: Release time (usually `xsd:date`).
  - `dcterms:modified`: Revision/update time (usually `xsd:date`).
  - `dcterms:accrualPeriodicity`: Update schedule (controlled vocabulary).
  - `dcat:temporalResolution`: Minimum temporal separation (`xsd:duration`).
  - `dcterms:temporal`: Temporal extent (`dcterms:PeriodOfTime`).
- **Spatial Properties:**
  - `dcat:spatialResolutionInMeters`: Minimum spatial separation (decimal number).
  - `dcterms:spatial`: Spatial extent (`dcterms:Location`).
  - WKT (Well-Known Text): Encoding for geometries.
  - CRS84 / WGS84: Default Coordinate Reference System.
- **Versioning Concepts:**
  - Version chain/history: Linear progression of versions.
  - Version hierarchy: Relationship between an abstract resource and its versions.
  - `dcat:previousVersion`: Links a version to its predecessor.
  - `dcat:hasVersion`: Links an abstract resource to its versions.
  - `dcat:hasCurrentVersion`: Links an abstract resource to the current snapshot.
  - `dcat:isVersionOf`: Inverse of `hasVersion`; links a version back to its abstract resource.

### Procedures And API Details
- **Expressing Rights:**
  1. Use `dcterms:license` for licenses (prefer canonical IRIs like Creative Commons).
  2. Use `dcterms:accessRights` for access restrictions (e.g., using [EUV-AR] code list).
  3. Use `dcterms:rights` with `rdfs:label` for copyright statements or other rights.
  4. For ODRL policies, use `odrl:hasPolicy` to link the resource/distribution to the policy object.
- **Expressing Temporal Properties:**
  - Encode dates as `xsd:date`.
  - Encode durations as `xsd:duration` (e.g., "PT15M").
  - Use `dcterms:PeriodOfTime` for temporal extent; can be combined with `time:ProperInterval` and `time:Instant`.
  - Support various time representations: `xsd:date`, `xsd:gYear`, or geologic age via `time:TimePosition` and specific TRS (e.g., GeologicAge).
- **Expressing Spatial Properties:**
  - Encode geometries in WKT format within `geosparql:wktLiteral`.
  - Omit CRS specification to imply default CRS84.
  - Specify polygons for detailed coverage, centroids for points, or bounding boxes (`dcat:bbox`) for general areas.
- **Defining Version Hierarchy:**
  - Define an abstract resource (e.g., `mycity-bus:stops`).
  - Link specific versions to the abstract resource using `dcat:hasVersion`.
  - Identify the current version using `dcat:hasCurrentVersion`.
  - Link individual versions backward using `dcat:previousVersion`.
  - Link versions back to the abstract resource using `dcat:isVersionOf`.

### Nuance Or Contradictions
- **Legal Advice:** The specification explicitly states that implementers must seek legal advice before deciding on access conditions, implying these recommendations are not legally binding mandates.
- **Non-Normative Sections:** Sections 9 (License), 10 (Time and space), and 11 (Versioning) are marked as non-normative, suggesting they provide guidance rather than strict requirements.
- **Complementary Versioning:** DCAT's versioning approach is meant to complement existing domain-specific practices (e.g., [OWL2-OVERVIEW] for ontologies) rather than replace them entirely. Providers decide when to release a new version based on their policies and workflows; DCAT does not define rules for triggering version releases.
- **Geometry CRS Default:** While specific CRS can be specified, omitting it defaults to CRS84 (WGS84 with longitude/latitude axis order), which may differ from expectations in other contexts using latitude/longitude first.

### Candidate Wiki Hints
- **Page: DCAT Rights Statements**
  - Topic: How to structure rights information (License vs. Access Rights vs. Copyright).
  - Content: Usage of `dcterms:rights`, `dcterms:license`, `dcterms:accessRights`, and integration with ODRL/ODRS.
- **Page: DCAT Temporal Metadata**
  - Topic: Describing time-related properties of datasets.
  - Content: Definitions of release, revision, schedule, resolution, and extent; examples of time-series data representation.
- **Page: DCAT Spatial Coverage**
  - Topic: Defining geographic boundaries for datasets.
  - Content: Usage of `dcterms:spatial`, WKT encoding, CRS handling, and examples (polygon, centroid, bbox).
- **Page: DCAT Versioning Model**
  - Topic: Managing resource versions and lifecycles in catalogs.
  - Content: Properties for version chains (`previousVersion`) and hierarchies (`hasVersion`, `hasCurrentVersion`), abstract resources, and relationship to PAV ontology.

## chunk-09

---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## Chunk Context
This chunk covers sections 11.1.2 through 12.1 of the DCAT vocabulary specification, detailing versioning relationships (replacement), version information properties, resource life-cycle management, complementary versioning approaches (OWL/DCTERMS/PROV-O), and the concept of dataset series including hierarchical linking.

## Local Summary
The text explains how to model versions that replace others using `dcterms:replaces`. It defines properties for version identifiers (`dcat:version`) and release dates (`dcterms:issued`). The resource life-cycle is treated as orthogonal to versioning, utilizing statuses like "deprecated" or "withdrawn". Finally, it introduces `dcat:DatasetSeries` for grouping related datasets (e.g., yearly budgets) using properties like `dcat:inSeries`, `dcat:first`, and `dcat:last`.

## Key Claims
- **Replacement Relationships:** DCAT uses `dcterms:replaces` and its inverse `dcterms:isReplacedBy` to indicate if a version supersedes another. This does not automatically form a strict linear chain; a version may replace any prior one, not just the immediate predecessor.
- **Version Information:** Key properties include `dcat:version` (identifier), `dcterms:issued` (release date), and `adms:versionNotes` (description of changes). DCAT defers to [DWBP] Best Practice 7 for naming conventions.
- **Life-Cycle Orthogonality:** A resource's life-cycle status (e.g., development, approved, withdrawn) is distinct from its versioning history. Statuses are managed via `adms:status` and various date properties (`dcterms:created`, `dcterms:issued`, etc.).
- **Complementary Versioning:** DCAT properties coexist with OWL (`owl:versionInfo`) and PROV-O (`prov:generalizationOf`). These have different scopes (e.g., OWL for ontologies, DCTERMS for broad editions) and are not strictly equivalent.
- **Dataset Series:** DCAT treats dataset series as first-class citizens via the `dcat:DatasetSeries` class (subclass of `dcat:Dataset`). Datasets link to a series via `dcat:inSeries`. Series can be hierarchical, with relationships for first/last members (`dcat:first`, `dcat:last`) and sequential neighbors (`dcat:prev`, `dcat:next`).

## Entities And Concepts
- **dcterms:replaces / dcterms:isReplacedBy:** Properties defining replacement relationships between versions.
- **dcat:version:** Property for the version name or identifier.
- **dcterms:issued:** Property for the release date of a version.
- **adms:versionNotes:** Property for textual descriptions of changes and compatibility issues.
- **adms:status:** Property for specifying life-cycle statuses (e.g., deprecated, withdrawn).
- **dcat:DatasetSeries:** A new class in DCAT representing a collection of interrelated datasets.
- **dcat:inSeries:** Property linking individual datasets to a dataset series.
- **dcat:first / dcat:last:** Properties identifying the first and last members of a dataset series.
- **dcat:prev / dcat:next:** Properties linking datasets sequentially within a series.
- **ISO-19135 / ISO-19115:** Standards providing example life-cycle status vocabularies (e.g., accepted, experimental, stable).

## Procedures And API Details
- **Specifying Replaced Versions:** Use `dcterms:replaces` to point from the new version to the old one. The inverse property `dcterms:isReplacedBy` can be used on the older version for back-linking.
- **Linking Dataset Series:** Create a resource typed as `dcat:DatasetSeries`. Link child datasets using `dcat:inSeries`. Use `dcat:first`, `dcat:last`, `dcat:prev`, and `dcat:next` to define the structure of the series over time.
- **Combining Series and Versions:** A dataset in a series can have its own versions. Link versions using `dcat:hasVersion` (listing all) and `dcat:hasCurrentVersion` (pointing to the active one).

## Nuance Or Contradictions
- **Non-Equivalence of Properties:** Correspondence between DCAT properties and those in OWL or PROV-O does not imply equivalence. For instance, `prov:wasRevisionOf` is semantically similar to `dcat:previousVersion` but is not intended for building version chains. Similarly, `owl:versionIRI` differs from `dcat:hasCurrentVersion` in scope and usage.
- **Life-Cycle vs. Versioning:** While related, life-cycle evolution (creation, approval, publication) does not always result in a new version, and creating a new version does not necessarily change the life-cycle status.

## Candidate Wiki Hints
- **Page: DCAT Versioning Relationships** (Concept: Explaining `dcterms:replaces` vs. linear chains).
- **Page: Dataset Series in DCAT** (Concept: Using `dcat:DatasetSeries` and hierarchical linking properties).

## chunk-10

---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

Chunk Context
- Heading path: 12. Dataset series > 12.2 Dataset series metadata, 12.3 Dataset series in existing DCAT implementations, 13. Data citation, 14. Quality information
- Lines covered: 4243–4637 (spanning dataset series inheritance rules, implementation variants, citation requirements, and quality documentation patterns).

Local Summary
This chunk details how metadata for dataset series is derived from child datasets via upstream inheritance, outlines alternative modeling approaches for series in current DCAT implementations, defines core elements for data citation, and introduces patterns for documenting quality information and conformance to standards using DQV, PROV-O, and EARL.

Key Claims
- Dataset series metadata can be split into two groups: properties describing the series itself (e.g., accrualPeriodicity) and properties inherited from child datasets via upstream inheritance.
- For inherited properties, the series value is typically the union of child values (temporal/spatial coverage), with specific rules for dates: earliest creation/publication, latest update/modification.
- Existing DCAT implementations either type the series as dcat:Dataset with children as dcat:Distribution, or both as dcat:Dataset linked by dcterms:hasPart/isPartOf; soft-typing via dcterms:type is also used.
- Data citation requires dataset identifier, creator(s), title, publisher, and publication/release date; DCAT 2 added dereferenceable identifiers and creator indication to support this.
- Quality information can be expressed using dqv:QualityAnnotation (feedback/certificates), dqv:QualityPolicy, or dqv:QualityMeasurement (metrics). Quality dimensions are not normative in DQV but may follow ISO/IEC 25012.
- Conformance to standards is modeled with dcterms:conformsTo and dcterms:Standard; best practices recommend using canonical, persistent, non-versioned IRIs from reference registries.

Entities And Concepts
- Dataset series metadata properties (dcterms:accrualPeriodicity, dcat:temporalResolution, dcat:temporal, dcat:spatial).
- Upstream inheritance mechanisms for child dataset values.
- Data citation elements (identifier, creator, title, publisher, publication date).
- Quality dimensions and DQV classes: dqv:QualityAnnotation, dqv:QualityPolicy, dqv:QualityMeasurement.
- Standard conformance modeling: dcterms:conformsTo, dcterms:Standard.
- Reference standards: EU INSPIRE Regulation (Commission Regulation (EU) No 1089/2010), OGC CRS Registry (EPSG:28992).

Procedures And API Details
- Compute series temporal coverage as the union of child start/end dates.
- Compute series spatial coverage as the union of child bounding boxes; handle multiple spatial reference systems if children differ.
- Set series dcterms:created to earliest child creation date.
- Set series dcterms:issued to earliest child publication date.
- Set series dcterms:modified to latest child publication/update date.
- Use IRIs from reference registries (W3C TR, OGC Definitions Server, ISO OBP) for standard conformance.
- Couple DCAT with DQV and PROV-O to express quality measurements and activities (e.g., ex:myQualityChecking using ex:myQualityChecker).

Nuance Or Contradictions
- DCAT does not prescribe a specific strategy for implementing upstream inheritance; mechanisms are left to implementers.
- Existing implementations may soft-type series via dcterms:type without violating DCAT, allowing coexistence with dcat:DatasetSeries during upgrades.
- DQV is non-normative and does not define a mandatory list of quality dimensions; implementers choose dimensions fitting their needs.

Candidate Wiki Hints
- Dataset Series Metadata Inheritance Rules
- Data Citation Requirements in DCAT
- Quality Information Modeling with DQV
- Documenting Conformance to Standards (dcterms:conformsTo)
- Using Canonical IRIs for Standard References

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

### Chunk Context
This section covers "Degree of conformance" (Section 14.2.2) and "Conformance test results" (Section 14.2.3), detailing how to model testing activities, plans, and results using PROV-O and EARL vocabularies. It also introduces qualified relations for datasets vs agents and resources, followed by an overview of DCAT profiles.

### Local Summary
The document outlines methods for expressing conformance degrees (e.g., conformant, not conformant) and compliance levels using `prov:Entity`, `prov:Activity`, and `dcterms:type`. It demonstrates how to link testing activities to plans derived from standards like DWBP. The text further explains the use of EARL for detailed test assertions and results, including handling failures with error messages and XML encoding. Finally, it describes "Qualified Relations" to specify roles (e.g., distributor, funder) beyond standard DCTERMS properties and lists various DCAT profiles extending the core vocabulary.

### Key Claims
- Legal contexts often require specifying a degree of conformance beyond simple binary compliance.
- Conformance tests are modeled as `prov:Activity` generating `prov:Entity` results associated with a `prov:Plan`.
- EARL vocabularies allow describing testing activities, assertions, and outcomes (passed/failed) alongside PROV-O.
- Qualified relations use an association class to attach specific roles (e.g., distributor, funder) to resources or agents.
- DCAT profiles are named sets of constraints extending the reference DCAT for specific domains (e.g., GeoDCAT-AP).

### Entities And Concepts
- **Conformance Degree**: Represented by `dcterms:type` with values like `ex:conformant`, `ex:notConformant`, or INSPIRE vocabulary terms.
- **Testing Model**: Components include `prov:Activity` (testing), `prov:Entity` (result), `prov:Plan` (test suite), and `prov:Agent` (validator).
- **EARL Vocabulary**: Used for `earl:Assertion`, `earl:TestResult`, `earl:mode` (manual/automatic), and `earl:outcome`.
- **Quality Metrics**: Defined using `dqv:Metric`, `dqv:Dimension`, and `dqv:QualityMeasurement` to quantify compliance percentages.
- **Qualified Relations**: Pattern involving `prov:Attribution` or `dcat:Relationship` with a `dcat:hadRole` property to specify roles like "funder" or "original".
- **DCAT Profiles**: Extensions of DCAT 2 for specific domains (e.g., GeoDCAT-AP, StatDCAT-AP).

### Procedures And API Details
1. **Modeling a Conformance Test**:
   - Create an `ex:testingActivity` (`prov:Activity`).
   - Link it to an `ex:conformanceTest` (`prov:Plan`) derived from a standard (e.g., DWBP).
   - Generate an `ex:testResult` (`prov:Entity`).
   - Assign the result type using `dcterms:type`.

2. **Modeling Test Results with EARL**:
   - Use `earl:Assertion` as the subject of the dataset tested.
   - Specify the test used via `earl:test`.
   - Indicate mode (`earl:automatic`) and outcome (`earl:passed` or `earl:fail`).
   - Include detailed error descriptions in `dcterms:description` (HTML/XML) for failed tests.

3. **Modeling Qualified Roles**:
   - For Agents: Use `prov:qualifiedAttribution` containing a `prov:Attribution` with `dcat:hadRole`.
   - For Resources: Use `dcat:qualifiedRelation` containing a `dcat:Relationship` with `dcterms:relation` and `dcat:hadRole`.

4. **Measuring Compliance**:
   - Define a metric (`dqv:Metric`) in a dimension (e.g., ISO compliance).
   - Create a quality measurement (`dqv:QualityMeasurement`) computed on the dataset with a numeric value (e.g., "50"^^xsd:double).

### Nuance Or Contradictions
- **Activity-Centric Limitation**: PROV-O is activity-centric and does not natively support direct Entity-to-Entity relations (except `was derived from`), necessitating new DCAT elements like `dcat:qualifiedRelation`.
- **Role Domains**: In PROV-O, roles relate to activities (`prov:hadRole` domain is `prov:Association`), whereas DCAT needs to attach roles directly to resource associations, hence the introduction of `dcat:hadRole`.
- **Vocabulary Preferences**: While non-normative URNs (e.g., from ISO codelists) are used in examples, linked data dereferenceable and normative representations are preferred.

### Candidate Wiki Hints
- **Page: Conformance Testing with DCAT**
  - Covers modeling testing activities, plans, and results using PROV-O and EARL.
  - Explains how to express partial compliance and specific conformance degrees.
- **Page: Qualified Relations in DCAT**
  - Details the pattern for assigning roles (agents vs resources) beyond standard DCTERMS properties.
  - Lists common role types (funder, distributor, original, etc.).
- **Page: DCAT Profiles Overview**
  - Catalogs existing profiles like GeoDCAT-AP, StatDCAT-AP, and regional variants.

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
- Heading: 17. Security and Privacy Considerations
- Line Range: 5050–5087

Local Summary
Section 17 addresses the handling of personal or private information within DCAT datasets and their metadata. It outlines requirements for securing sensitive data, verifying integrity via checksums (specifically borrowing from SPDX), and ensuring authenticity through HTTPS origins and separate verification channels. The section explicitly states that detailed web security implementations are out of scope for this vocabulary.

Key Claims
- DCAT metadata itself can contain personal or private information regarding creators, publishers, and qualified relations.
- Sensitive data must be stored securely and restricted to authorized parties per legal requirements.
- Detailed guidance on securing web content and user authentication is beyond the scope of DCAT.
- Checksums serve as verification for integrity and authenticity, particularly for critical data like software vulnerabilities.
- A checksum must be provided via a route separate from the data it validates to prevent manipulation by attackers.
- Integrity and authenticity ultimately depend on the trustworthiness of the source; providers must address these at both application and transport levels.

Entities And Concepts
- DCAT vocabulary
- Personal or private information
- Qualified relations
- Sensitive data
- SPDX (specifically spdx:Checksum class)
- Hash / Checksum value
- HTTPS origins
- Application level security
- Transport level security

Procedures And API Details
- Publishers should provide a checksum value and the associated algorithm for each resource in a distribution.
- The checksum route must be separate from the data it sums (e.g., including a separate metadata file within a tarfile alongside the distribution file).
- Providers should ensure their API and download endpoints are secured for integrity and authenticity.
- DCAT data and metadata files should be downloadable only from authoritative HTTPS origins.
- Checksums must be provided via a channel separate from the data they represent.

Nuance Or Contradictions
- Storing checksums within the same file as the distribution (e.g., in a tarfile) is permissible but requires a secondary, separate checksum for the metadata file itself to prevent an attacker from altering both the data and its verification simultaneously.
- Providing a checksum in DCAT metadata alone does not guarantee integrity or authenticity unless the metadata's own integrity and authenticity are independently guaranteed.

Candidate Wiki Hints
- Security Best Practices for Data Catalogs
- Managing Sensitive Metadata in DCAT
- Implementing Checksum Verification for Distributions

## chunk-13

---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
# Chunk Context
**Heading**: 18. Accessibility Considerations (Lines 5088–5500)
**Scope**: Section 18 covers accessibility, acknowledgments, alignment with Schema.org, and catalog examples including provenance and publication linking.

# Local Summary
This chunk details how DCAT supports accessibility via alternative text for non-text resources, lists working group contributors, maps DCAT to Schema.org (DCAT → sdo), and provides RDF examples of loosely structured catalogs, dataset provenance using PROV-O, and linking datasets to publications.

# Key Claims
- Enforcing alternative text for non-text data improves accessibility; this aligns with WCAG 2.0 guidelines.
- DCAT backbone classes map closely to Schema.org (e.g., `dcat:Dataset` ↔ `sdo:Dataset`, `dcat:Distribution` ↔ `sdo:DataDownload`).
- A recommended mapping from revised DCAT to Schema.org v3.4 is available as an RDF file using RDFS/OWL/SKOS predicates.
- Legacy catalogs (e.g., CKAN) often treat datasets as bags of files; DCAT distinguishes distributions from other relations via `dcat:distribution` vs. `dcterms:relation`/`dcterms:hasPart`.
- Provenance can be expressed using PROV-O (`prov:wasGeneratedBy`, `prov:wasAttributedTo`, `prov:wasDerivedFrom`).
- Datasets should link to publications via `dcterms:isReferencedBy`.

# Entities And Concepts
- **DCAT vocabulary**: Data Catalog Vocabulary for describing catalogs and datasets.
- **WCAG 2.0**: Web Content Accessibility Guidelines referenced as [UNDERSTANDING-WCAG20].
- **Schema.org**: General-purpose metadata vocabulary used by search services; DCAT aligns with it for broader exposure.
- **PROV-O**: W3C Provenance Ontology for describing dataset origins and production agents.
- **Dublin Core (dcterms)**: Used in DCAT for titles, creators, identifiers, etc.
- **RDF/OWL/SKOS**: Semantic web technologies used in the DCAT-to-Schema.org mapping.
- **CKAN**: Example of a legacy catalog where datasets are loosely structured.
- **Dryad / Nature Scientific Data**: Examples of repositories and journals linked to datasets via publications.

# Procedures And API Details
- **Mapping DCAT → Schema.org**: Use an RDF file containing axiomatizations (`rdfs:subClassOf`, `owl:equivalentClass`, etc.) and annotation properties (`sdo:domainIncludes`, `sdo:rangeIncludes`).
- **Representing distributions**: Prefer `dcat:distribution` for proper representations; use `dcterms:relation`/`dcterms:hasPart` when relationship type is unknown.
- **Expressing provenance**: Link to generating projects with `prov:wasGeneratedBy`; link agents via `prov:wasAttributedTo`; link predecessor datasets via `prov:wasDerivedFrom`.
- **Linking to publications**: Use `dcterms:isReferencedBy` to connect datasets to scholarly articles or reports.
- **Example RDF graphs**: Provided in the DXWG DCAT 3 code repository (`csiro-dap-examples.ttl`, `csiro-stratchart_dcat3.ttl`).

# Nuance Or Contradictions
- Some legacy systems conflate distributions with arbitrary file relations; DCAT clarifies this distinction.
- PROV-O properties complement but do not replace Dublin Core properties (e.g., `prov:wasAttributedTo` vs. `dcterms:creator`).
- The mapping to Schema.org is non-normative and intended for interoperability with search indexes rather than strict equivalence.

# Candidate Wiki Hints
- **Page**: DCAT–Schema.org Mapping Guide (summarize the RDF mapping table and use cases).
- **Page**: Best Practices for Accessibility in DCAT Profiles (alternative text enforcement, WCAG alignment).
- **Page**: Using PROV-O with DCAT for Dataset Provenance (patterns for agents, activities, derived datasets).
- **Page**: Linking Datasets to Publications (using `dcterms:isReferencedBy` with examples from Dryad/Nature).

## chunk-14

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

## chunk-15

---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
----------------
This chunk covers **Section 18. Accessibility Considerations** within the DCAT 2 Vocabulary document, though the provided text primarily details revision history (changes to versioning, dataset series, SPDX checksums, geometry usage, license examples, namespace prefixes, and URI/IRI terminology). It also lists changes since the W3C Recommendation of February 4, 2020, and provides a comprehensive list of normative and informative references used by the vocabulary.

Local Summary
----------------
The document has been updated to clarify versioning strategies using the [PAV] approach and to remove backward compatibility property specifications (`owl:backwardCompatibleWith`, `owl:incompatibleWith`). A new class `dcat:DatasetSeries` was introduced to treat dataset series as first-class citizens, alongside properties for linking series members (`dcat:inSeries`) and defining sequence order (`dcat:first`, `dcat:prev`, `dcat:next`, `dcat:last`). The vocabulary now includes SPDX checksums for distributions and updated geometry usage aligned with [LOCN]. Editorial changes include fixing inconsistent "URI"/"IRI" usage, replacing the `dct:` prefix with `dcterms:`, and clarifying that catalog resources are not limited to datasets. Changes since the 2020 Recommendation emphasize using specific subrelations over `dcterms:relation` and adding draft guidelines for versioning and dataset series.

Key Claims
----------------
- Versioning is now focused on versions derived from resource revision, following the [PAV] approach for version chains and hierarchies (previous, next, current, last).
- Support for specifying backward/incompatible compatibility between versions using `owl:backwardCompatibleWith` and `owl:incompatibleWith` has been dropped.
- Dataset series are first-class citizens; a new class `dcat:DatasetSeries` defines the relationship between datasets in a series.
- Properties `dcat:first`, `dcat:prev`, `dcat:next`, and `dcat:last` define the sequence of resources within a series.
- Property `spdx:checksum` was added to the Distribution class, introducing `spdx:Checksum` with properties `spdx:algorithm` and `spdx:checksumValue`.
- The range of `locn:geometry` has been revised to align with [LOCN], allowing usage with geometry literals or classes.
- Examples in the License section have been updated to address issues #676 and #1333.
- The namespace prefix for DCMI Metadata Terms was changed from `dct:` to `dcterms:` throughout the document.
- Inconsistent use of "URI" and "IRI" has been fixed to ensure terminology consistency.

Entities And Concepts
-----------------------
- **DCAT 2 Vocabulary**: The core data catalog vocabulary being maintained.
- **Versioning ([PAV])**: Approach for specifying version chains/hierarchies using `previous`, `next`, `current`, `last`.
- **Dataset Series**: A new concept where datasets are members of a series, represented by class `dcat:DatasetSeries`.
- **SPDX Checksum**: Property added to Distribution to specify checksums using the SPDX vocabulary.
- **LOCN Geometry**: Alignment with the ISA Programme Location Core Vocabulary for geometry representation.
- **W3C Recommendation (2020)**: The baseline document version from February 4, 2020.

Procedures And API Details
-----------------------------
- **Namespace Update**: Replace `dct:` with `dcterms:` in all references to DCMI Metadata Terms.
- **Geometry Usage**: Use `locn:geometry` with either geometry literals or classes, aligned with [LOCN].
- **Version Specification**: Use properties from the PAV ontology (e.g., `pav:previousVersion`) instead of OWL compatibility properties.
- **Checksum Specification**: Add `spdx:checksum` to a Distribution resource using `spdx:algorithm` and `spdx:checksumValue`.

Nuance Or Contradictions
---------------------------
- The chunk title mentions "Accessibility Considerations" (Section 18), but the content focuses on versioning, dataset series, and references. This suggests the provided text might be a fragment of a larger document where Section 18 follows or precedes these revision notes.
- The removal of `owl:backwardCompatibleWith` implies a shift away from explicitly modeling compatibility states in favor of the [PAV] approach for version relationships.

Candidate Wiki Hints
-----------------------
- **Dataset Series**: Create a page explaining `dcat:DatasetSeries`, `dcat:inSeries`, and sequence properties (`first`, `prev`, `next`, `last`).
- **Versioning Best Practices**: Summarize the shift to [PAV] for version management and the removal of OWL compatibility properties.
- **SPDX Integration**: Document how to use `spdx:checksum` within DCAT distributions.
- **Geometry Alignment**: Explain the updated usage of `locn:geometry`.

## chunk-16

---
title: Chunk 16 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
- Heading: 18. Accessibility Considerations
- Line Range: 6379-6468
- Source File: raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md

Local Summary
This chunk serves as a bibliography or reference list for the "Accessibility Considerations" section of the Data Catalog Vocabulary specification. It enumerates a collection of related W3C specifications, OGC standards, and academic papers relevant to data description, quality assessment, geospatial vocabularies, and web services that inform accessibility and interoperability practices within the DCAT ecosystem.

Key Claims
- The document lists multiple versions of the Data Catalog Vocabulary (DCAT) ranging from Version 2 (2020 Recommendation) through Version 3 (Candidate Recommendation as of Jan 2024).
- Related W3C vocabularies and ontologies cited include DQV (Data Quality), ORG (Organization), SSN (Semantic Sensor Network), VOID (Linked Datasets), XHTML Vocabulary, and Basic Geo.
- External standards included are OGC's Web Feature Service 2.0 (WFS) and Web Map Service (WMS), as well as W3C's Web Services Description Language (WSDL) 2.0.
- A specific academic survey on "Quality assessment for Linked Data" by Zaveri et al. is cited alongside these technical specifications.

Entities And Concepts
- **DCAT-2**: Data Catalog Vocabulary Version 2.
- **DCAT-3**: Data Catalog Vocabulary Version 3.
- **DQV**: Data Quality Vocabulary.
- **ORG**: Organization Ontology.
- **SSN**: Semantic Sensor Network Ontology.
- **VOID**: VoID (Describing Linked Datasets).
- **WFS**: Web Feature Service 2.0.
- **WMS**: Web Map Service.
- **WSDL 2.0**: Web Services Description Language Version 2.0.
- **Basic Geo**: Vocabulary for WGS84 latitude/longitude.
- **ZaveriEtAl**: Survey on Quality assessment for Linked Data.

Procedures And API Details
- No specific procedural steps or API calls are detailed in this chunk; it functions as a citation list of normative and informative references.
- References provided include URLs to W3C Recommendations, Working Drafts, Candidate Recommendations, and external OGC standards pages.

Nuance Or Contradictions
- The chunk lists both DCAT Version 2 (W3C Recommendation) and several drafts/CRs of Version 3, indicating an evolution in the standard's maturity status within the W3C process.
- The inclusion of diverse topics (from sensor networks to map services) suggests a broad scope of interoperability concerns relevant to accessibility considerations, though specific implementation details are absent.

Candidate Wiki Hints
- **Page: DCAT Accessibility References** – A dedicated page listing these external specifications could serve as a quick lookup for implementers needing to ensure their catalog descriptions align with broader W3C and OGC standards.

