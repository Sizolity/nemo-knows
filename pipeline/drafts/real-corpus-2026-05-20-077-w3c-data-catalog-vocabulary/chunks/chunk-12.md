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
