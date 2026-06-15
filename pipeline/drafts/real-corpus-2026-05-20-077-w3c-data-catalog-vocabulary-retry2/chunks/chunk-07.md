---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk covers the vocabulary specification for relationship, role, time, and location classes in DCAT 2 (and 3), plus sections on inverse properties and dereferenceable identifiers.

Local Summary
The document defines new RDF classes and properties introduced in DCAT 2 for expressing relationships between datasets, roles of agents or entities, temporal intervals, spatial locations, and file integrity checksums. It also clarifies how to use inverse properties, manage various identifier types (including non-HTTP IDs), and link external metadata.

Key Claims
- Class dcat:Relationship is an association class for attaching additional information to a relationship between DCAT resources; it subclasses prov:EntityInfluence.
- dcat:relation points to another dataset or cataloged resource in the context of dcat:Relationship.
- dcat:hadRole expresses the function of an entity or agent with respect to another entity or resource; recommended values come from controlled vocabularies like ISO-19115 CI_RoleCode, DS_AssociationTypeCode, IANA RELATIONS, DataCite, or MARC relators.
- dcat:Role is a subclass of skos:Concept used in qualified-attribution or qualified-relation; values should be managed as controlled vocabularies.
- dcterms:PeriodOfTime represents an interval defined by start and end; DCAT 2 uses dcat:startDate/dcat:endDate (ISO 8601) or time:hasBeginning/time:hasEnd (OWL-TIME).
- dcat:bbox, dcat:centroid, locn:geometry are used for spatial coverage; WKT literals support non-WGS84 coordinate systems.
- spdx:Checksum and its properties algorithm/checksumValue allow verification of file integrity using SPDX-defined algorithms (MD5, SHA families, etc.).
- DCAT intentionally omits inverse properties in the core vocabulary but allows them as additions, not replacements.
- HTTP IRIs are preferred for dereferenceable identifiers; legacy or non-HTTP identifiers can be expressed via dcterms:identifier, adms:identifier, and related metadata (adms:schemaAgency, skos:notation).
- owl:sameAs may link to external RDF descriptions (e.g., DOIs) that return machine-readable metadata.

Entities And Concepts
- dcat:Relationship, dcat:Role, dcterms:PeriodOfTime, dcterms:Location, spdx:Checksum
- Properties: dcat:relation, dcat:hadRole, dcat:startDate, dcat:endDate, time:hasBeginning, time:hasEnd, locn:geometry, dcat:bbox, dcat:centroid, spdx:algorithm, spdx:checksumValue
- Inverse properties: dcat:prev/dcat:next, dcat:distribution/dcat:isDistributionOf, dcterms:hasPart/dcterms:isPartOf, etc.
- Identifier types: HTTP IRIs, DOI, ORCID, VIAF, ISNI, locally minted identifiers, legacy identifiers
- Controlled vocabularies: ISO-19115 CI_RoleCode, DS_AssociationTypeCode, IANA RELATIONS, DataCite, MARC relators

Procedures And API Details
- Use dcat:startDate/dcat:endDate with ISO 8601 datatypes (xsd:gYear, xsd:date, xsd:dateTime) for temporal intervals.
- Use time:hasBeginning/time:hasEnd when the value is a time:Instant from OWL-TIME; implies dcterms:PeriodOfTime can be treated as time:ProperInterval.
- Encode geometry as WKT literals (geosparql:wktLiteral) within rdfs:Literal ranges for bbox/centroid; supports non-WGS84 CRS.
- For checksums, pair spdx:algorithm with spdx:checksumValue (xsd:hexBinary) to verify file integrity.
- Represent non-HTTP identifiers using adms:identifier with skos:notation and optionally adms:schemaAgency; use dcterms:publisher/creator for authority metadata.
- When an external IRI returns RDF/OWL, consider owl:sameAs to link the dataset to that description.

Nuance Or Contradictions
- DCAT 2 introduces dcat:hadRole and dcat:Role; spdx:Checksum is new in DCAT 3.
- Inverse properties are allowed only as additions to the core vocabulary, never as replacements.
- For DOIs, the authority (e.g., DOI Foundation) is recorded via adms:schemaAgency or dcterms:creator, not by naming the registrant organization, to preserve DOI stability across organizational changes.
- Identifier schemes registered as IRIs (e.g., info:doi/) do not need a separate 'type' annotation per RFC 3986.

Candidate Wiki Hints
- Page on DCAT relationship modeling (dcat:Relationship, dcat:relation, dcat:hadRole)
- Page on role vocabularies and controlled lists (ISO-19115, DataCite, MARC relators)
- Guide to temporal properties in DCAT (start/end dates vs time:hasBeginning/hasEnd)
- Spatial coverage patterns (geometry, bbox, centroid; WKT literals)
- File integrity with SPDX checksums (algorithm + hex digest)
- Identifier management strategies (HTTP IRIs, legacy IDs, adms:identifier, schemaAgency)
- Using owl:sameAs for external RDF metadata enrichment
