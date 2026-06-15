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
