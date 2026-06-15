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
- Scope: Guidelines for handling personal data within DCAT, integrity verification via checksums, and trust model requirements.

Local Summary
This section addresses how the Data Catalog Vocabulary (DCAT) handles datasets containing personal information and metadata that might reveal private details about creators or publishers. It outlines responsibilities for implementers to ensure secure storage and access control, noting that detailed web security mechanisms are out of scope. The text emphasizes using checksums from SPDX to verify data integrity, requiring separate delivery channels for checksums to prevent tampering. Ultimately, the trustworthiness of DCAT data relies on securing API endpoints and serving content over authoritative HTTPS origins.

Key Claims
- DCAT vocabulary supports datasets containing personal or private information.
- Metadata expressed with DCAT may itself contain personal/private information (e.g., resource creators, publishers).
- Implementers must take steps to ensure security and privacy considerations are addressed for sensitive data.
- Sensitive data and metadata must be stored securely and made available only to authorized parties per legal/functional requirements.
- Detailing how to secure web content and authenticate users is beyond the scope of DCAT.
- Checksums can serve as verification for datasets requiring assurances of integrity and authenticity (e.g., software vulnerabilities).
- DCAT borrows `spdx:Checksum` from SPDX to ensure integrity and authenticity of DCAT distributions.
- A checksum must be provided via a route separate from the data it sums.
- If a checksum is included in metadata with the data, a checksum for that metadata must also be provided separately to foil attackers.
- Integrity and authenticity of DCAT metadata depend on the trustworthiness of the source.
- Providers should address integrity/authenticity at application and transport levels (e.g., HTTPS origins).

Entities And Concepts
- DCAT (Data Catalog Vocabulary)
- Personal/private information
- Resource creators, publishers, agents
- Sensitive data/metadata
- Checksums (hashes), algorithms
- `spdx:Checksum` class
- Integrity and authenticity assurances
- API endpoints, download endpoints
- Authoritative HTTPS origins
- Application level security
- Transport level security

Procedures And API Details
- Store sensitive data securely; restrict access to authorized parties.
- Provide a checksum value (hash) and the generating algorithm for each resource in a distribution.
- Deliver checksums via a separate channel from the data they represent.
  - Example: A tarfile containing a distribution file and a metadata file with the distribution's checksum.
  - Requirement: Also provide a checksum for the metadata file itself if embedded with the data.
- Ensure integrity and authenticity of API and download endpoints.
- Make DCAT data/metadata files downloadable from authoritative HTTPS origins.

Nuance Or Contradictions
- No direct contradictions found; recommendations align with standard security practices (separation of data and verification, use of HTTPS).
- The distinction between verifying the data checksum and the metadata checksum is critical to prevent manipulation attacks.
- DCAT explicitly defers detailed web security implementation details to external standards or organizational policies.

Candidate Wiki Hints
- Security best practices for DCAT implementations
- Handling privacy in catalog metadata
- Integrity verification using SPDX checksums in DCAT distributions
