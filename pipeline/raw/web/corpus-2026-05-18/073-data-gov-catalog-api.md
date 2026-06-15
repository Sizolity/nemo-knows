---
title: Data.gov Catalog API
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://resources.data.gov/catalog-api/
tags: [data, web-corpus]
confidence: medium
---

# Data.gov Catalog API

## Fetch Metadata

- Corpus item: 73
- Category: Data
- Source URL: https://resources.data.gov/catalog-api/
- Final URL: https://resources.data.gov/catalog-api/
- Retrieved: 2026-05-18
- Content-Type: text/html; charset=utf-8
- Fetch status: ok
- Test value: Tests API documentation summarisation.
- Fetched page title: Catalog API | resources.data.gov

## Retrieved Text

Skip to main content

U.S. flag

An official website of the United States government

Here’s how you know
(BUTTON) Here's how you know
Dot gov

The .gov means it’s official.
Federal government websites often end in .gov or .mil. Before sharing
sensitive information, make sure you’re on a federal government site.
Https

The site is secure.
The https:// ensures that you are connecting to the official website
and that any information you provide is encrypted and transmitted
securely.

resources.data.gov
(BUTTON) Menu

(BUTTON) close
* (BUTTON) Resources
+ Data management & governance
+ Data tools
+ Guidance
+ Harvester Help
* (BUTTON) Data Standards
+ DCAT-US 3
+ DCAT-US v1.1
* (BUTTON) API Documentation
+ General Information
+ Catalog API
+ API.Data.Gov
* (BUTTON) Agency Links
+ Catalog.Data.Gov
+ Harvest.Data.gov
+ Inventory.data.gov
* (BUTTON) About
+ About this repository
+ Governance

* Glossary

Search small ____________________ (BUTTON) Search Search

Catalog API

The Data.gov Catalog API provides access to metadata about datasets
published by federal, state, local, and tribal governments. You can use
it to search for datasets, filter by organization or topic, and
retrieve detailed information about individual records.

4/9/2026: Note on upcoming changes: The base URL for the data.gov APIs
will change when Data.gov routes its APIs through api.data.gov. This is
expected to occur soon. If you’re building a new integration, avoid
hard-coding the base URL and check this page for updates before going
to production. Email us with any questions at datagovhelp@gsa.gov

Base URL: https://catalog.data.gov

No API key is required. All endpoints are publicly accessible and
return JSON.

Note: This API replaces the previous CKAN-based API. The prior endpoint
remains available in a read-only state for existing integrations, but
new development should use this API.
__________________________________________________________________

Getting Started

Here is a complete working example that walks through finding NASA
climate datasets step by step.

Step 1: Find the organization slug for NASA

To filter results by organization, you first need the organization’s
slug. Fetch the full list of organizations and find the one you want.

Request:
GET https://catalog.data.gov/api/organizations

Look for NASA in the response:
{
"organizations": [
{
"id": "f4ca4614-8901-409b-8553-2e994ad10023",
"name": "National Aeronautics and Space Administration",
"slug": "nasa",
"organization_type": "Federal Government",
"dataset_count": 27040
}
]
}

The slug is nasa. You will use this to filter search results.

Step 2: Search for climate datasets from NASA

Use the q parameter for your keyword and org_slug to filter by
organization.

Request:
GET https://catalog.data.gov/search?q=climate&org_slug=nasa&per_page=3

Response:
{
"sort": "relevance",
"after": "WzY5LjM0NDY5NiwwLCJiYmRhZGNmYi00NDM1LTQzZWUtYjhlMy0yMzZiZjBlZDEwODIi
XQ==",
"results": [
{
"title": "Amazon Web Services: Downscaled Climate Projections (NEX-DCP30)"
,
"publisher": "AWS NEX",
"accessLevel": "public",
"keyword": ["amazon-web-services", "aws", "climate", "earth-science"],
"last_harvested_date": "2025-08-04T13:35:12.398986",
"landingPage": "https://aws.amazon.com/nasa/nex/"
},
{
"title": "Mirador - Climate Variability and Change",
"publisher": "National Aeronautics and Space Administration",
"accessLevel": "public",
"keyword": ["aerosols", "atmospheric-height", "atmospheric-radiation"],
"last_harvested_date": "2025-08-03T15:49:47.819080"
}
]
}

The response includes an after value, which means there are more
results available.

Step 3: Get the next page

Pass the after value from the previous response to retrieve the next
page. Keep all other parameters the same.

Request:
GET https://catalog.data.gov/search?q=climate&org_slug=nasa&per_page=3&after=WzY
5LjM0NDY5NiwwLCJiYmRhZGNmYi00NDM1LTQzZWUtYjhlMy0yMzZiZjBlZDEwODIiXQ==

Continue repeating this step until the response no longer includes an
after field, which means you have reached the last page. For more
details see Pagination below.
__________________________________________________________________

Table of Contents

1. Search Datasets
2. Get Keywords
3. Search Locations
4. Get Location Geometry
5. Get Organizations
6. Get Harvest Record
7. Get Harvest Record Raw
8. Get Harvest Record Transformed
9. Pagination
10. Response Codes and Errors
__________________________________________________________________

Search Datasets

Search the catalog for datasets using keywords, filters, and sorting
options.

Endpoint: /search Method: GET

Query Parameters

Parameter Type Required Default Description
q string No "" Full-text search query
sort string No relevance Sort order: relevance, popularity, distance,
or last_harvested_date
per_page integer No 10 Number of results to return per page (minimum:
1)
org_slug string No - Filter by organization slug (e.g., nasa). Use the
Get Organizations endpoint to find valid slugs.
org_type string No - Filter by organization type. Valid values: Federal
Government, City Government, State Government, County Government,
University, Tribal, Non-Profit
keyword array No - Filter by one or more keywords (exact match). Repeat
the parameter for multiple values.
spatial_filter string No - Limit results to datasets with or without
spatial data: geospatial or non-geospatial
spatial_geometry string (GeoJSON) No - A GeoJSON geometry object
defining a geographic shape to filter by. Use with spatial_within to
control how datasets are matched against the shape.
spatial_within boolean No false Controls how datasets are matched
against spatial_geometry. When false (default), returns datasets whose
spatial extent intersects the shape. When true, returns only datasets
whose spatial extent falls completely within the shape.
after string No - Pagination cursor returned from a previous response.
See Pagination.

Example Requests

GET https://catalog.data.gov/search?q=water+quality
GET https://catalog.data.gov/search?q=climate&sort=popularity&per_page=25
GET https://catalog.data.gov/search?org_slug=nasa&per_page=10
GET https://catalog.data.gov/search?org_type=Federal+Government&spatial_filter=g
eospatial
GET https://catalog.data.gov/search?q=education&after=WzEwMC4wNjEzNiwwLCJiMWEzOT
Y3YzJhMTExZjE2NzgxN2IwMTI0YzUyYjBhYyJd
GET https://catalog.data.gov/search?spatial_geometry={"type":"Polygon","coordina
tes":[[[-109.05,37.0],[-102.05,37.0],[-102.05,41.0],[-109.05,41.0],[-109.05,37.0
]]]}&spatial_within=true

Response

Status Code: 200 OK
{
"after": "WzEwMC4wNjEzNiwwLCJiMWEzOTY3YzJhMTExZjE2NzgxN2IwMTI0YzUyYjBhYyJd",
"sort": "relevance",
"results": [
{
"title": "National Household Education Surveys Program, 2012 Parent and Fa
mily Involvement in Education Survey",
"description": "A cross-sectional survey collecting data from households o
n educational issues...",
"identifier": "bdf82c61-0027-4d50-9505-44fc57f2fd12",
"slug": "national-household-education-surveys-program-2012-parent-and-fami
ly-involvement-in-educati",
"publisher": "National Center for Education Statistics (NCES)",
"keyword": ["education", "homeschooling", "households", "parental-involvem
ent-in-education"],
"has_spatial": true,
"popularity": 2,
"last_harvested_date": "2025-08-02T21:17:47.154806",
"distribution_titles": [
"National Household Education Surveys Program (NHES):2012 Restricted-Use
Data Files",
"2012PFIascii.zip"
],
"theme": [],
"spatial_centroid": null,
"spatial_shape": null,
"organization": {
"id": "217e855b-cd64-4ebc-958b-abbbb0f57ac2",
"name": "Department of Education",
"slug": "ed",
"organization_type": "Federal Government",
"logo": "https://raw.githubusercontent.com/GSA/logo/refs/heads/master/ed
.png",
"aliases": ["dept"],
"description": null
},
"dcat": {
"@type": "dcat:Dataset",
"title": "National Household Education Surveys Program, 2012...",
"description": "...",
"accessLevel": "restricted public",
"accrualPeriodicity": "irregular",
"bureauCode": ["018:50"],
"contactPoint": {
"@type": "vcard:Contact",
"fn": "Sarah Grady",
"hasEmail": "mailto:sarah.grady@ed.gov"
},
"dataQuality": true,
"distribution": ["..."],
"identifier": "bdf82c61-0027-4d50-9505-44fc57f2fd12",
"issued": "2014-05-21",
"keyword": ["education", "homeschooling"],
"language": ["en-US"],
"license": "https://creativecommons.org/publicdomain/zero/1.0/",
"modified": "2023-06-22T20:25:39.652070",
"programCode": ["018:000"],
"publisher": {
"@type": "org:Organization",
"name": "National Center for Education Statistics (NCES)"
},
"rights": "This dataset has restricted access.",
"spatial": "United States",
"systemOfRecords": "https://www2.ed.gov/notices/pai/pai-18-13-01.pdf",
"temporal": "2012/2012"
},
"harvest_record": "https://catalog.data.gov/harvest_record/0be6d0c0-8383-4
966-acd1-38b0d7baea3c",
"harvest_record_raw": "https://catalog.data.gov/harvest_record/0be6d0c0-83
83-4966-acd1-38b0d7baea3c/raw"
}
]
}

Response Fields

Field Type Description
results array List of matching datasets
sort string The sort order applied to this response
after string Cursor for retrieving the next page of results. Absent if
there are no more results.
results[].title string Dataset title
results[].description string Dataset description
results[].identifier string (UUID) Unique dataset identifier
results[].slug string URL-friendly identifier for the dataset
results[].publisher string Name of the publishing organization
results[].keyword array List of keywords associated with the dataset
results[].theme array List of themes associated with the dataset
results[].has_spatial boolean Whether the dataset has a spatial
component
results[].spatial_centroid object or null Geographic center point of
the dataset's spatial coverage, if available
results[].spatial_shape object or null GeoJSON shape representing the
dataset's spatial coverage, if available
results[].popularity integer Relative popularity score for the dataset
results[].last_harvested_date string (ISO 8601) Date and time the
dataset was last ingested into the catalog
results[].distribution_titles array Titles of the dataset's available
distributions (downloads, APIs, etc.)
results[].organization object Information about the publishing
organization. See Get Organizations for field definitions.
results[].dcat object Full DCAT-US metadata for the dataset. See DCAT
Object Fields below.
results[].harvest_record string (URL) Link to the harvest record for
this dataset, if available
results[].harvest_record_raw string (URL) Link to the raw source
payload for this dataset's harvest record, if available

DCAT Object Fields

The dcat object contains the full DCAT-US metadata as submitted by the
publishing organization. Not all fields are present for every dataset.
Fields marked as always present appear in virtually every result;
others are optional and vary by publisher.
Field Type Always Present Description
@type string No Always dcat:Dataset when present
title string Yes Dataset title
description string Yes Full dataset description
identifier string Yes Unique dataset identifier as assigned by the
source system
accessLevel string Yes One of: public, restricted public, or non-public
modified string Yes Date the dataset was last modified (ISO 8601)
publisher object Yes Publishing organization. Contains @type
(org:Organization) and name
contactPoint object Yes Contact information. Contains @type
(vcard:Contact), fn (name), and hasEmail
keyword array Yes List of keywords describing the dataset
distribution array No Available downloads and access methods, each
following the dcat:Distribution structure
landingPage string (URL) No URL of the dataset's home page
license string (URL) No URL of the license under which the dataset is
published
bureauCode array No Federal bureau code(s) in DDD:XX format. Present on
federal datasets.
programCode array No Federal program code(s). Present on federal
datasets.
issued string No Date the dataset was first published (ISO 8601)
theme array No Thematic categories for the dataset (e.g.,
Transportation, Health)
spatial string No Geographic coverage of the dataset (e.g., United
States or a bounding box)
temporal string No Time period covered by the dataset (e.g.,
2018-01-01/2018-09-28)
accrualPeriodicity string No How frequently the dataset is updated,
using ISO 8601 duration format (e.g., R/P1Y for annual, R/P1D for
daily)
language array No Language(s) the dataset is available in (e.g., en-US)
rights string No Description of any access restrictions or rights
statement
describedBy string (URL) No URL of a data dictionary or schema
describing the dataset
describedByType string No MIME type of the resource at describedBy
(e.g., application/pdf)
references array No URLs of related documents or resources
isPartOf string No Identifier of a parent dataset that this dataset
belongs to
dataQuality boolean No Whether the dataset meets the publisher's data
quality guidelines
conformsTo string (URL) No URL of a standard or specification the
dataset conforms to
primaryITInvestmentUII string No Federal IT investment identifier in
DDD-XXXXXXXXX format
systemOfRecords string (URL) No URL of the Privacy Act System of
Records Notice, if applicable
phone string No Contact phone number for the dataset
__________________________________________________________________

Get Keywords

Retrieve a list of the most commonly used keywords across all datasets,
along with how many datasets use each one.

Endpoint: /api/keywords Method: GET

Query Parameters

Parameter Type Required Default Valid Range Description
size integer No 100 1-1000 Maximum number of keywords to return
min_count integer No 1 ≥1 Only return keywords used by at least this
many datasets

Example Requests

GET https://catalog.data.gov/api/keywords
GET https://catalog.data.gov/api/keywords?size=50&min_count=100

Response

Status Code: 200 OK
{
"keywords": [
{ "keyword": "County or Equivalent Entity", "count": 90507 },
{ "keyword": "State FIPS Code", "count": 48578 }
],
"size": 2,
"min_count": 1,
"total": 2
}

Keywords are sorted by count, highest first.
__________________________________________________________________

Search Locations

Search for location names to use with spatial filtering. This endpoint
is designed for autocomplete — pass a partial place name and receive
matching location IDs.

Endpoint: /api/locations/search Method: GET

Query Parameters

Parameter Type Required Default Description
q string No - Partial or full location name to search for
size integer No - Maximum number of results to return

Example Request

GET https://catalog.data.gov/api/locations/search?q=Colorado&size=5

Response

Status Code: 200 OK
{
"locations": [
{
"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
"display_name": "Colorado, United States"
}
],
"size": 1,
"total": 1
}

Use the id value from this response with the Get Location Geometry
endpoint.
__________________________________________________________________

Get Location Geometry

Retrieve the geographic boundary (GeoJSON geometry) for a specific
location by its ID.

Endpoint: /api/location/{location_id} Method: GET

Path Parameters

Parameter Type Required Description
location_id string (UUID) Yes The location ID returned from
/api/locations/search

Example Request

GET https://catalog.data.gov/api/location/a1b2c3d4-e5f6-7890-abcd-ef1234567890

Response

Status Code: 200 OK or 404 Not Found
{
"id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
"geometry": {
"type": "Polygon",
"coordinates": [[[...]]]
}
}
__________________________________________________________________

Get Organizations

Retrieve the complete list of organizations that publish datasets in
the catalog.

Endpoint: /api/organizations Method: GET

No query parameters. Returns all organizations.

Example Request

GET https://catalog.data.gov/api/organizations

Response

Status Code: 200 OK
{
"organizations": [
{
"id": "f4ca4614-8901-409b-8553-2e994ad10023",
"name": "National Aeronautics and Space Administration",
"slug": "nasa",
"organization_type": "Federal Government",
"aliases": [""],
"dataset_count": 27040
}
],
"total": 312
}

Response Fields

Field Type Description
id string (UUID) Unique organization identifier
name string Organization display name
slug string URL-friendly identifier, usable as org_slug in search
organization_type string Type of organization: Federal Government,
State Government, City Government, County Government, University,
Tribal, or Non-Profit
aliases array Alternative names or abbreviations for the organization
dataset_count integer Number of datasets published by this organization
__________________________________________________________________

Get Harvest Record

Retrieve metadata about a specific harvest record by its ID. Harvest
records track how individual datasets were ingested into the catalog.

Endpoint: /harvest_record/{record_id} Method: GET

Path Parameters

Parameter Type Required Description
record_id string (UUID) Yes The harvest record ID

Example Request

GET https://catalog.data.gov/harvest_record/d0e03fb2-f885-4b1d-8feb-2d8acc93f4f8

Response

Status Code: 200 OK or 404 Not Found
{
"id": "d0e03fb2-f885-4b1d-8feb-2d8acc93f4f8",
"identifier": "http://datainventory.doi.gov/id/dataset/bsee-0000000070",
"status": "error",
"action": "create",
"date_created": "2025-11-26T07:46:13.673655",
"date_finished": null,
"harvest_job_id": "de2010f9-d9ec-4211-9690-5b3bbc9fe1f3",
"harvest_source_id": "14348973-07a5-4661-8341-02230f2f6cbb",
"source_hash": "47ca2dd5471e659e4cd1c83d79adb0b0c2c8c013a1e03d629d56b0541e3072
67",
"source_raw": { },
"source_transform": null,
"ckan_id": null,
"ckan_name": null,
"parent_identifier": null
}

Notes

* record_id must be a valid UUID format
* Date fields are returned in ISO 8601 format
* source_raw is parsed as JSON when possible

Error Response:
{ "error": "Not Found" }
__________________________________________________________________

Get Harvest Record Raw

Retrieve the original, unmodified source payload from a harvest record
exactly as it was received.

Endpoint: /harvest_record/{record_id}/raw Method: GET

Path Parameters

Parameter Type Required Description
record_id string (UUID) Yes The harvest record ID

Example Request

GET https://catalog.data.gov/harvest_record/d0e03fb2-f885-4b1d-8feb-2d8acc93f4f8
/raw

Response

Status Code: 200 OK or 404 Not Found

The Content-Type of the response is detected automatically based on the
payload:
* application/json for JSON payloads
* application/xml for XML payloads
* text/plain for all other content

Returns 404 if the record does not exist or has no raw source data.
__________________________________________________________________

Get Harvest Record Transformed

Retrieve the transformed DCAT-US payload for a harvest record. This is
the version of the metadata after any source-specific transformations
have been applied.

Endpoint: /harvest_record/{record_id}/transformed Method: GET

Path Parameters

Parameter Type Required Description
record_id string (UUID) Yes The harvest record ID

Example Request

GET https://catalog.data.gov/harvest_record/000c4ce7-90c6-405c-8ed7-3ae06c45005c
/transformed

Response

Status Code: 200 OK or 404 Not Found Content-Type: application/json

Returns 404 if the record does not exist or has no transformed data.
__________________________________________________________________

Pagination

The /search endpoint uses cursor-based pagination. This approach is
more reliable than page-number pagination for large result sets because
it maintains consistent ordering even as the catalog changes.

How it works:
1. Make a request to /search. If more results exist beyond what was
returned, the response will include an after field.
2. To get the next page, add after=<value> to your next request using
the value from the previous response.
3. Continue until the response no longer includes an after field,
which means you have reached the last page.

Example:
# First request
GET https://catalog.data.gov/search?q=water&per_page=10

# Response includes: "after": "WzEwMC4wNjEzNiwwLCJiMWEz..."

# Second request
GET https://catalog.data.gov/search?q=water&per_page=10&after=WzEwMC4wNjEzNiwwLC
JiMWEz...

Keep all other parameters the same across pages. Changing q, sort, or
filter parameters while paginating will return inconsistent results.
__________________________________________________________________

Response Codes and Errors

Status Code Meaning
200 OK Request was successful
404 Not Found The requested resource does not exist, or the ID provided
is not valid
422 Unprocessable Entity The request was understood but contains
invalid parameter values
500 Internal Server Error An unexpected error occurred on the server

All error responses use this JSON format:
{
"error": "A description of what went wrong"
}

For validation errors (422), additional detail is provided:
{
"message": "Validation error",
"detail": {
"<location>": {
"<field_name>": ["error message"]
}
}
}

Return to top

Office of Management and Budget
General Services Administration
NARA Office of Government Information Services

The Federal Enterprise Data Resources content is maintained by the
Data.gov Program Management Office in GSA TTS, the Office of Government
and Information Services (OGIS), and the Office of Management and
Budget (OMB).

Help us improve this site on GitHub.
* Data.gov
* Privacy policy
* Contact us

resources.data.gov

An official website of the Office of Management and Budget, the General
Services Administration, and the Office of Government Information
Services.
* About OMB
* Accessibility statement
* FOIA requests
* No FEAR Act data
* Office of the Inspector General
* Performance reports
* Privacy policy

Looking for U.S. government information and services?
Visit USA.gov

(BUTTON) Close Hide glossary

Glossary

This section contains explanations of common terms referenced on
resources.data.gov.
Search for a term: ____________________
