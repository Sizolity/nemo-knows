---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

Chunk Context
- Location: Section 14.2.2 "Degree of conformance" through 14.2.3 "Conformance test results", followed by Section 15 "Qualified relations" and the start of Section 16 "DCAT Profiles".
- Scope: Describes modeling conformance test outcomes (full, partial, or non-compliance), uses of PROV-O, EARL, and VOCAB-DQV for quality metrics, and introduces qualified relationships between datasets, agents, and other resources.

Local Summary
The document details how to express the degree of conformance for datasets using controlled vocabularies and provenance modeling. It illustrates patterns for reporting full compliance, partial compliance (e.g., 50%), and non-conformance using `prov:Entity`, `prov:Activity`, `prov:Plan`, and EARL classes. It also covers qualified relations to assign specific roles (e.g., distributor, funder) to agents or other resources without creating an explosion of properties.

Key Claims
- Legal contexts like INSPIRE require specifying degrees of conformance beyond simple "conformant/not conformant".
- Conformance tests can be modeled using `prov:Entity` for results, `prov:Activity` for the test execution, and `prov:Plan` derived from standards like DWBP.
- Partial compliance (e.g., 50% of tests passed) is expressible via quality metrics in VOCAB-DQV.
- EARL provides specific classes (`earl:Assertion`, `earl:TestResult`) to describe testing activities, complementing PROV-O.
- Qualified relations use an additional resource (e.g., `prov:Attribution`, `dcat:Relationship`) to carry parameters like roles, avoiding property explosion.

Entities And Concepts
- **Degree of Conformance**: Represented by SKOS concepts (e.g., `ex:conformant`, `ex:notConformant`, `ex:nonConformant`) or INSPIRE vocabulary (`http://inspire.ec.europa.eu/metadata-codelist/DegreeOfConformity`).
- **Provenance Modeling**: Uses `prov:Entity` (results), `prov:Activity` (testing), `prov:Plan` (test suite derived from DWBP), and `prov:Association` to bind activities to plans/agents.
- **Quality Metrics**: Defined in VOCAB-DQV (`dqv:Metric`, `dqv:Dimension`, `dqv:Category`, `dqv:QualityMeasurement`) to measure compliance percentages.
- **EARL Testing**: Uses `earl:Assertion`, `earl:TestRequirement`, `earl:TestResult`, `earl:mode` (manual/automatic), and `earl:outcome`.
- **Qualified Relations**: Pattern using `prov:qualifiedAttribution` and `dcat:qualifiedRelation` with `dcat:hadRole` to specify roles like distributor, funder, or stereo-mate.
- **Vocabularies Referenced**: DWBP (Data on the Web Best Practices), ISO/IEC 25012 (quality dimensions), INSPIRE DoC, VOCAB-DQV, EARL, PROV-O, DCTERMS, CI_RoleCode, DS_AssociationTypeCodes.

Procedures And API Details
- **Modeling Conformance Test**:
  1. Define a `prov:Plan` (e.g., `ex:conformanceTest`) derived from a standard (e.g., DWBP).
  2. Create a `prov:Activity` (`ex:testingActivity`) using the plan.
  3. Generate an `earl:Assertion` or `prov:Entity` (`ex:testResult`) as output.
  4. Bind the activity to the plan via `prov:qualifiedAssociation` (or EARL's direct assertion).
- **Reporting Partial Compliance**:
  1. Define a `dqv:Metric` (e.g., `levelOfComplianceToDWBP`).
  2. Create a `dqv:QualityMeasurement` with a numeric value (e.g., "50"^^xsd:double) and unit (Percentage).
- **Reporting Non-Conformance**:
  1. Use `earl:Assertion` with `earl:result` containing an `earl:TestResult`.
  2. Set `dcterms:type` to `ex:nonConformant` or use INSPIRE's non-conformant URI.
  3. Include human-readable errors via `dcterms:description` (HTML/XML) and technical details in `earl:info`.
- **Assigning Agent Roles**:
  1. Use `prov:qualifiedAttribution` with a nested `prov:Attribution`.
  2. Link an agent (`prov:agent`) to the resource via `dcat:hadRole` inside the attribution.
- **Relating Datasets**:
  1. Use `dcat:qualifiedRelation` with a nested `dcat:Relationship`.
  2. Specify the target resource via `dcterms:relation` and the role via `dcat:hadRole`.

Nuance Or Contradictions
- **PROV-O Limitations**: PROV-O is activity-centric and lacks support for direct Entity-Entity relations (except "was derived from"), necessitating DCAT's new elements (`dcat:qualifiedRelation`, `dcat:Relationship`).
- **Role Property Domain**: In PROV-O, `prov:hadRole` relates to activities, not entities; thus, DCAT introduces `dcat:hadRole` on the association-class `prov:Attribution`.
- **Vocabulary Choices**: The text notes that while IRIs from non-normative codelists (e.g., CI_RoleCode via URNs) are used in examples, linked data dereferenceable and normative representations should be preferred when available.
- **Test Mode Ambiguity**: Example 51 notes uncertainty about whether a test was manual or automatic (`earl:mode earl:automatic.` with a comment "we do not know..."), highlighting the need for explicit metadata to clarify test execution modes.

Candidate Wiki Hints
- **Page: Modeling Conformance Tests in DCAT**
  - Focus: How to use PROV-O and EARL to model conformance testing activities, plans, and results.
  - Key content: Examples of `prov:Plan`, `prov:Activity`, `earl:Assertion`, and handling partial/non-conformance.
- **Page: Qualified Relations in DCAT**
  - Focus: Using `dcat:qualifiedRelation` and `prov:qualifiedAttribution` to assign roles (e.g., distributor, funder) without property explosion.
  - Key content: Patterns for relating datasets to agents and other resources with specific semantic roles.
- **Page: Quality Metrics for Data Compliance**
  - Focus: Using VOCAB-DQV to define metrics for compliance percentages and data quality dimensions.
  - Key content: Defining `dqv:Metric`, `dqv:QualityMeasurement`, and linking to ISO/IEC 25012 dimensions.
