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
