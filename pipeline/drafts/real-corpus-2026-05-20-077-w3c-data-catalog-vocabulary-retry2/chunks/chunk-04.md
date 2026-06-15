---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## Chunk Context
This chunk covers vocabulary specifications for the `Cataloged Resource` class in DCAT, spanning properties from **language** (6.4.9) through **first** (6.4.31). It details definitions, ranges, sub-properties, and usage notes for metadata describing resource attributes, relationships, versions, statuses, and series membership.

## Local Summary
The section defines RDF properties used to describe cataloged resources, ranging from basic attributes like language and publisher to complex relationship patterns involving qualified relations and versioning. It highlights transitions between DCAT 1, 2, and 3, noting domain relaxations and the introduction of new properties for version chains and resource series.

## Key Claims
- **Language**: Use ISO 639-1 codes; fall back to ISO 639-2 if ISO 639-1 is undefined. Values for members override catalog-level values.
- **Publisher**: Corresponds to general attribution with the role 'publisher'; `foaf:Agent` resources are recommended.
- **Identifier**: Represents a unique text string assigned to a resource; may be used in the IRI but explicitly represented as an `rdfs:Literal`.
- **Theme/Category**: Organized within a `skos:ConceptScheme`; DCAT 3 treats it as an OWL object property to avoid automatic inference as `skos:Concept`.
- **Type/Genre**: Values should come from controlled vocabularies (e.g., DCMI, ISO 19115). File formats and physical media should use `dcterms:format` instead.
- **Relation**: Generic link when relationship nature is unknown; specific sub-properties (e.g., `dcat:distribution`) are preferred if the relationship type is known.
- **Qualified Relation**: Used for relationships not covered by standard DCTERMS or PROV-O properties, linking to a `dcat:Relationship`.
- **Versioning**: DCAT 3 introduces specific properties (`previousVersion`, `hasVersion`, `currentVersion`) limited to revisions in a resource's life-cycle, distinct from broader notions of edition/adaptation.
- **Series**: Properties like `first` are used for resources belonging to a `dcat:DatasetSeries`.

## Entities And Concepts
- **Classes**: `Cataloged Resource`, `Dataset`, `Distribution`, `Agent`, `Organization/Person`.
- **Vocabularies**: DCMI Type, ISO 19115, DataCite, PARSE.Insight (re3data), MARC.
- **Standards**: ISO 639-1, ISO 639-2, ISO 639-3, BCP47, PROV-O, ODRL, VOCAB-ADMS.
- **Properties**: `dcterms:language`, `dcterms:publisher`, `dcterms:identifier`, `dcat:theme`, `dcterms:type`, `dcterms:relation`, `dcat:qualifiedRelation`, `prov:qualifiedAttribution`, `dcterms:license`, `dcterms:rights`, `dcterms:hasPart`, `odrl:hasPolicy`, `dcterms:isReferencedBy`, `dcat:previousVersion`, `dcat:hasVersion`, `dcat:currentVersion`, `dcterms:replaces`, `dcat:version`, `adms:versionNotes`, `adms:status`, `dcat:first`.

## Procedures And API Details
- **Language Selection**: Check for ISO 639-1 code first; if absent, use the corresponding ISO 639-2 IRI.
- **Distribution Mapping**: If a dataset has separate representations per language, define distinct `dcat:Distribution` instances with specific `dcterms:language` values. Multilingual distributions have multiple language values.
- **Qualified Relation Usage**: Construct a link to another resource where the relationship nature is known but does not match standard properties. This entails using `dcat:qualifiedRelation` and linking to a `dcat:Relationship`.
- **Version Chain Construction**: Use `dcat:previousVersion`, `dcat:hasVersion`, and `dcat:currentVersion` to specify version chains consisting of snapshots resulting from revisions.
- **Series Membership**: Use `dcat:first` to identify the first resource in an ordered collection or series (e.g., `dcat:DatasetSeries`).

## Nuance Or Contradictions
- **Theme Range Change**: In DCAT 1, `dcat:theme` domain was limited to `dcat:Dataset`. DCAT 2 relaxed this, and DCAT 3 drops the range entirely to prevent automatic inference of objects as `skos:Concept`, ensuring formalization consistency.
- **Version Semantics**: `dcat:hasVersion` is a sub-property of `dcterms:hasVersion` but is more specific, limited to revisions in the life-cycle, whereas `dcterms:hasVersion` covers broader notions like editions and adaptations.
- **Identifier vs IRI**: While the identifier might be part of the resource's IRI, explicitly representing it as an `rdfs:Literal` is still considered useful.

## Candidate Wiki Hints
- Page on **DCAT Vocabulary Properties** covering attributes like language, publisher, and identifiers.
- Guide on **Managing Resource Versions** in DCAT 3, focusing on revision chains vs. editions.
- Article on **Qualified Relations** explaining the pattern for non-standard relationships using PROV-O and DCTERMS sub-properties.
- Documentation on **Dataset Series** utilizing `dcat:first` and related ordering properties.
