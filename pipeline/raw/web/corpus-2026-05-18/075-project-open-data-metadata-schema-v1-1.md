---
title: Project Open Data Metadata Schema v1.1
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://resources.data.gov/resources/dcat-us/
tags: [data, web-corpus]
confidence: medium
---

# Project Open Data Metadata Schema v1.1

## Fetch Metadata

- Corpus item: 75
- Category: Data
- Source URL: https://resources.data.gov/resources/dcat-us/
- Final URL: https://resources.data.gov/resources/dcat-us/
- Retrieved: 2026-05-18
- Content-Type: text/html; charset=utf-8
- Fetch status: ok
- Test value: Useful for comparing schema versions.
- Fetched page title: DCAT-US Schema v1.1 (Project Open Data Metadata Schema) | resources.data.gov

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

* Overview
* Details
+ Introduction
+ Standard Metadata Vocabulary
+ What to Document – Datasets and Web APIs
+ Metadata File Format – JSON
+ Catalog Fields
+ Dataset Fields
+ Dataset Distribution Fields
+ Further Metadata Field Guidance
+ Federal Government Fields
+ Rationale for Metadata Nomenclature

DCAT-US Schema v1.1 (Project Open Data Metadata Schema)

Overview

How to use Project Open Data Metadata Schema guidelines to document and
list agency datasets and application programming interfaces (APIs) for
hosting at agency.gov/data and currently in use at data.gov

Source

data.gov

Category

Data standards

Keywords

data schema,
open data,
DCAT,
Project Open Data Metadata Schema,
data standards,
data inventory

Links

* Metadata Resources
* DCAT

Details

Specification Name: DCAT-US Schema v1.1 (Project Open Data Metadata
Schema)
This version: 1.1
Latest version: This version
Publication date: November 6th 2014
__________________________________________________________________

Introduction

This section contains guidance to support the use of the Project Open
Data metadata to list agency datasets and application programming
interfaces (APIs) as hosted at agency.gov/data. Additional technical
information about the schema can be found on the Metadata Resources
page.

Standard Metadata Vocabulary

Metadata is structured information that describes, explains, locates,
or otherwise makes it easier to retrieve, use, or manage an information
resource (NISO 2004, ISBN: 1-880124-62-9). The challenge is to define
and name standard metadata fields so that a data consumer has
sufficient information to process and understand the described data.
The more information that can be conveyed in a standardized regular
format, the more valuable data becomes. Metadata can range from basic
to advanced, from allowing one to discover the mere fact that a certain
data asset exists and is about a general subject all the way to
providing detailed information documenting the structure, processing
history, quality, relationships, and other properties of a dataset.
Making metadata machine readable greatly increases its utility, but
requires more detailed standardization, defining not only field names,
but also how information is encoded in the metadata fields.

Establishing a common vocabulary is the key to communication. The
metadata schema specified in this memorandum is based on DCAT, a
hierarchical vocabulary specific to datasets. This specification
defines three types of metadata elements: Required, Required-if
(conditionally required), and Expanded fields. These elements were
selected to represent information that is most often looked for on the
web. To assist users of other metadata standards, field mappings to
equivalent elements in other standards are provided.

What to Document – Datasets and Web APIs

A dataset is an identifiable collection of structured data objects
unified by some criteria (authorship, subject, scope, spatial or
temporal extent…). A catalog is a collection of descriptions of
datasets; each description is a metadata record. The intention of a
data catalog is to facilitate data access by users who are first
interested in a particular kind of data, and upon finding a
fit-for-purpose dataset, will next want to know how to get the data.

A Web API (Application Programming Interface) allows computer programs
to dynamically query a dataset using the World Wide Web. For example, a
dataset of farmers markets may be made available for download as a
single file (e.g., a CSV), or may be made available to developers
through a Web API, such that a computer program could use a ZIP Code to
retrieve a list of farmers markets in the ZIP Code area.

The catalog file for each agency should list all of the agency’s
datasets that can be made public, regardless of whether they are
distributed by a file download or a Web API. Please also see the
extended guidance on documenting Web APIs in your data.json files.

Metadata File Format – JSON

The Implementation Guidance available as a part of Project Open Data
describes Agency requirements for the development of metadata as per
the Open Data Policy. A quick primer on the file format involved:

JSON is a lightweight data-exchange format that is very easy to read,
parse and generate. Based on a subset of the JavaScript programming
language, JSON is a text format that is optimized for data interchange.
JSON is built on two structures: (1) a collection of name/value pairs
and (2) an ordered list of values.

Where optional fields are included in a catalog file but are
unpopulated, they may be represented by a null value. They should not
be represented by an empty string ("").

When a record has an accessURL or downloadURL, they should be contained
as objects within a distribution. Any object may be described by title,
description, format, or mediaType, though when an object contains
downloadURL, it must be accompanied by mediaType.

The Project Open Data schema is case sensitive. The schema uses a camel
case convention where the first letter of some words within a field are
capitalized (usually all words but the first one). While it may seem
subtle which characters are uppercase and lowercase, it is necessary to
follow the exact same casing as defined in the schema documented here.
For example:

Correct: contactPoint
Incorrect: ContactPoint
Incorrect: contactpoint
incorrect: CONTACTPOINT

Links to downloadable examples of metadata files developed in this and
other formats are in the metadata resources. Agencies can validate
their metadata using this validator.

Catalog Fields

These fields describe the entire Public Data Listing catalog file.
Publishers can also use the describedBy field to reference the default
JSON Schema file used to define the schema
(https://project-open-data.cio.gov/v1.1/schema/catalog.json) or they
may refer to their own JSON Schema file if they have extended the
schema with additional schema definitions. Similarly, @context can be
used to reference the default JSON-LD Context used to define the schema
(https://project-open-data.cio.gov/v1.1/schema/catalog.jsonld) or
publishers can refer to their own if they have extended the schema with
additional linked data vocabularies. See the Catalog section under
Further Metadata Field Guidance for more details.
Field Label Definition Required
@context Metadata Context URL or JSON object for the JSON-LD Context
that defines the schema used. No
@id Metadata Catalog ID IRI for the JSON-LD Node Identifier of the
Catalog. This should be the URL of the data.json file itself. No
@type Metadata Type IRI for the JSON-LD data type. This should be
dcat:Catalog for the Catalog. No
conformsTo Schema Version URI that identifies the version of the
Project Open Data schema being used. Always
describedBy Data Dictionary URL for the JSON Schema file that defines
the schema used. No
dataset Dataset A container for the array of Dataset objects. See
Dataset Fields below for details. Always

Dataset Fields

See the Further Metadata Field Guidance section to learn more about the
use of each element, including the range of valid entries where
appropriate. Consult the field mappings to find the equivalent v1.0,
DCAT, Schema.org, and CKAN fields.
Field Label Definition Required
@type Metadata Type IRI for the JSON-LD data type. This should be
dcat:Dataset for each Dataset. No
title Title Human-readable name of the asset. Should be in plain
English and include sufficient detail to facilitate search and
discovery. Always
description Description Human-readable description (e.g., an abstract)
with sufficient detail to enable a user to quickly understand whether
the asset is of interest. Always
keyword Tags Tags (or keywords) help users discover your dataset;
please include terms that would be used by technical and non-technical
users. Always
modified Last Update Most recent date on which the dataset was changed,
updated or modified. Always
publisher Publisher The publishing entity and optionally their parent
organization(s). Always
contactPoint Contact Name and Email Contact person’s name and email for
the asset. Always
identifier Unique Identifier A unique identifier for the dataset or API
as maintained within an Agency catalog or database. Always
accessLevel Public Access Level The degree to which this dataset could
be made publicly-available, regardless of whether it has been made
available. Choices: public (Data asset is or could be made publicly
available to all without restrictions), restricted public (Data asset
is available under certain use restrictions), or non-public (Data asset
is not available to members of the public). Always
bureauCode^USG Bureau Code Federal agencies, combined agency and bureau
code from OMB Circular A-11, Appendix C (PDF, CSV in the format of
015:11. Always
programCode^USG Program Code Federal agencies, list the primary program
related to this data asset, from the Federal Program Inventory. Use the
format of 015:001. Always
license License The license or non-license (i.e. Public Domain) status
with which the dataset or API has been published. See Open Licenses for
more information. If-Applicable
rights Rights This may include information regarding access or
restrictions based on privacy, security, or other policies. This should
also serve as an explanation for the selected “accessLevel” including
instructions for how to access a restricted file, if applicable, or
explanation for why a “non-public” or “restricted public” data asset is
not “public,” if applicable. Text, 255 characters. If-Applicable
spatial Spatial The range of spatial applicability of a dataset. Could
include a spatial region like a bounding box or a named place.
If-Applicable
temporal Temporal The range of temporal applicability of a dataset
(i.e., a start and end date of applicability for the data).
If-Applicable
distribution Distribution A container for the array of Distribution
objects. See Dataset Distribution Fields below for details.
If-Applicable
accrualPeriodicity Frequency The frequency with which dataset is
published. No
conformsTo Data Standard URI used to identify a standardized
specification the dataset conforms to. No
dataQuality^USG Data Quality Whether the dataset meets the agency’s
Information Quality Guidelines (true/false). No
describedBy Data Dictionary URL to the data dictionary for the dataset.
Note that documentation other than a data dictionary can be referenced
using Related Documents (references). No
describedByType Data Dictionary Type The machine-readable file format
(IANA Media Type also known as MIME Type) of the dataset’s Data
Dictionary (describedBy). No
isPartOf Collection The collection of which the dataset is a subset. No
issued Release Date Date of formal issuance. No
language Language The language of the dataset. No
landingPage Homepage URL This field is not intended for an agency’s
homepage (e.g. www.agency.gov), but rather if a dataset has a
human-friendly hub or landing page that users can be directed to for
all resources tied to the dataset. No
primaryITInvestmentUII^USG Primary IT Investment UII For linking a
dataset with an IT Unique Investment Identifier (UII). No
references Related Documents Related documents such as technical
information about a dataset, developer documentation, etc. No
systemOfRecords^USG System of Records If the system is designated as a
system of records under the Privacy Act of 1974, provide the URL to the
System of Records Notice related to this dataset. No
theme Category Main thematic category of the dataset. No

Dataset Distribution Fields

Within a dataset, distribution is used to aggregate the metadata
specific to a dataset’s resources (accessURL and downloadURL), which
may be described using the following fields. Each distribution should
contain one accessURL or downloadURL. A downloadURL should always be
accompanied by mediaType.
Field Label Definition Required
@type Metadata Type IRI for the JSON-LD data type. This should be
dcat:Distribution for each Distribution. No
accessURL Access URL URL providing indirect access to a dataset, for
example via API or a graphical interface. If-Applicable
conformsTo Data Standard URI used to identify a standardized
specification the distribution conforms to. No
describedBy Data Dictionary URL to the data dictionary for the
distribution found at the downloadURL. Note that documentation other
than a data dictionary can be referenced using Related Documents as
shown in the expanded fields. No
describedByType Data Dictionary Type The machine-readable file format
(IANA Media Type or MIME Type) of the distribution’s describedBy URL.
No
description Description Human-readable description of the distribution.
No
downloadURL Download URL URL providing direct access to a downloadable
file of a dataset. If-Applicable
format Format A human-readable description of the file format of a
distribution. No
mediaType Media Type The machine-readable file format (IANA Media Type
or MIME Type) of the distribution’s downloadURL. If-Applicable
title Title Human-readable name of the distribution. No

Extending the Schema

“Extensional” and/or domain specific metadata can easily be added using
other vocabularies even if it is not a term (entity/property) that will
get indexed by the major search engines - it could still be indexed by
other custom search engines and by Data.gov. Publishers are encouraged
to extend their metadata descriptions using elements from the “Expanded
Fields” list shown below, or from any well-known vocabulary (including
Dublin Core, Schema.org, FGDC, ISO 19115, and NIEM) as long as they are
properly assigned. It’s also recommended that these extensions be
defined through the describedBy and @context fields at the top of the
Catalog metadata.

Further Metadata Field Guidance

Additional details for each field are provided here broken down into
sections for the overarching Catalog, each dataset, and each dataset’s
distribution. Consult the field mappings to find the equivalent v1.0,
DCAT, Schema.org, and CKAN fields.
* Key
+ Required
+ Required if Applicable
+ Expanded (optional)

* Catalog
+ @context
+ @id
+ @type
+ conformsTo
+ describedBy
+ dataset
o @type
o accessLevel
o accrualPeriodicity
o bureauCode
o conformsTo
o contactPoint
# @type
# fn
# hasEmail
o dataQuality
o describedBy
o describedByType
o description
o distribution
# @type
# accessURL
# conformsTo
# downloadURL
# describedBy
# describedByType
# description
# format
# mediaType
# title
o identifier
o isPartOf
o issued
o keyword
o landingPage
o language
o license
o modified
o primaryITInvestmentUII
o programCode
o publisher
# @type
# name
# subOrganizationOf
o references
o rights
o spatial
o systemOfRecords
o temporal
o theme
o title

Catalog Fields
__________________________________________________________________

Field # @context
Cardinality (0,1)
Required No
Accepted Values String (URL)
Usage Notes The URL or JSON object for the JSON-LD Context that defines
the schema used. The URL for version 1.1 of the schema is
https://project-open-data.cio.gov/v1.1/schema/catalog.jsonld
Example {"@context":
"https://project-open-data.cio.gov/v1.1/schema/catalog.jsonld"}
Field # @id
Cardinality (0,1)
Required No
Accepted Values String (IRI)
Usage Notes A unique identifier for the Catalog as defined by JSON-LD
Node Identifiers. This should be the URL of the data.json file itself
Example {"@id": "https://www.agency.gov/data.json"}
Field # @type
Cardinality (0,1)
Required No
Accepted Values String (IRI)
Usage Notes The metadata type as defined by JSON-LD data types. This
should be dcat:Catalog for the Catalog
Example {"@type": "dcat:Catalog"}
Field # conformsTo
Cardinality (1,1)
Required Yes, always
Accepted Values String (URI)
Usage Notes This is used to identify the schema version using a URI.
The URI for version 1.1 of the schema is
https://project-open-data.cio.gov/v1.1/schema
Example {"conformsTo": "https://project-open-data.cio.gov/v1.1/schema"}
Field # describedBy
Cardinality (0,1)
Required No
Accepted Values String (URL)
Usage Notes This is used to specify a JSON Schema file that defines all
fields. By default, it is recommended that the canonical JSON Schema
file is referenced
(https://project-open-data.cio.gov/v1.1/schema/catalog.json) but if the
schema had been extended, publishers may reference a file that defines
those extensions.
Example {"describedBy":
"https://project-open-data.cio.gov/v1.1/schema/catalog.json"}
Field # dataset
Cardinality (1,n)
Required Yes, always
Accepted Values Array of Objects
Usage Notes This field is a container for an array of Dataset objects.
See Dataset Fields below for details
Example {"dataset": [...]}

Dataset Fields
__________________________________________________________________

Field # @type
Cardinality (0,1)
Required No
Accepted Values String (IRI)
Usage Notes The metadata type as defined by JSON-LD data types. This
should be dcat:Dataset for the Dataset
Example {"@type": "dcat:Dataset"}
Field # accessLevel
Cardinality (1,1)
Required Yes, always
Accepted Values Must be one of the following: “public”, “restricted
public”, “non-public”
Usage Notes This field refers to the degree to which this dataset could
be made available to the public, regardless of whether it is currently
available to the public. For example, if a member of the public can
walk into your agency and obtain a dataset, that entry is public even
if there are no files online. A restricted public dataset is one only
available under certain conditions or to certain audiences (such as
researchers who sign a waiver). A non-public dataset is one that could
never be made available to the public for privacy, security, or other
reasons as determined by your agency.
Example {"accessLevel":"public"}
Field # accrualPeriodicity
Cardinality (0,1)
Required No
Accepted Values ISO 8601 Repeating Duration (or irregular)
Usage Notes Must be an ISO 8601 repeating duration unless this is not
possible because the accrual periodicity is completely irregular, in
which case the value should simply be irregular. The value should not
include a start or end date but rather simply express the duration of
time between data publishing. For example, a dataset which is published
on an annual basis would be R/P1Y; every three months would be R/P3M;
weekly would be R/P1W; and daily would be R/P1D. Further examples and
documentation can be found here.
Example {"accrualPeriodicity":"R/P1Y"}
Field # bureauCode
Cardinality (0,n)
Required Yes, for United States Federal Government agencies
Accepted Values Array of Strings
Usage Notes Represent each bureau responsible for the dataset according
to the codes found in OMB Circular A-11, Appendix C (PDF, CSV). Start
with the agency code, then a colon, then the bureau code.
Example The Office of the Solicitor (86) at the Department of the
Interior (010) would be: {"bureauCode":["010:86"]}. If a second bureau
was also responsible, the format like this:
{"bureauCode":["010:86","010:04"]}.
Field # conformsTo
Cardinality (0,1)
Required No
Accepted Values String (URI)
Usage Notes This is used to identify a standardized specification the
dataset conforms to. If this is a technical specification associated
with a particular serialization of a distribution, this should be
specified with conformsTo at the distribution level. It’s recommended
that this be a URI that serves as a unique identifier for the standard.
The URI may or may not also be a URL that provides documentation of the
specification.
Example {"conformsTo":
"http://www.agency.gov/common-vegetable-analysis-model/"}
Field # contactPoint
Cardinality (1,1)
Required Yes, always
Accepted Values vCard object
Usage Notes This is a container for two fields that together make up
the contact information for the dataset. contactPoint should always
contain both the person’s appropriately formatted full name (fn) and
email (hasEmail).
Example See below
"contactPoint": {
"@type": "vcard:Contact",
"fn": "Jane Doe",
"hasEmail": "mailto:jane.doe@agency.gov"
}

Field # contactPoint → @type
Cardinality (0,1)
Required No
Accepted Values String (IRI)
Usage Notes The metadata type as defined by JSON-LD data types. This
should be vcard:Contact for contactPoint
Example {"@type": "vcard:Contact"}
Field # contactPoint → fn
Cardinality (1,1)
Required Yes, always
Accepted Values String
Usage Notes This should include included with hasEmail as part of a
record’s contactPoint (see above example).
Example {"fn": "Jane Doe"}
Field # contactPoint → hasEmail
Cardinality (1,1)
Required Yes, always
Accepted Values String
Usage Notes This should be formatted per vCard specifications (see
example below) and included with fn as part of a record’s contactPoint
(see above example).
Example {"hasEmail": "mailto:jane.doe@agency.gov"}
Field # dataQuality
Cardinality (0,1)
Required No
Accepted Values Must be a boolean value of true or false (not contained
within quote marks)
Usage Notes Indicates whether a dataset conforms to the agency’s
information quality guidelines.
Example {"dataQuality":true}
Field # describedBy
Cardinality (0,1)
Required No
Accepted Values String (URL)
Usage Notes This is used to specify a data dictionary or schema that
defines fields or column headings in the dataset. If this is a machine
readable file, it’s recommended to be specified with describedBy at the
distribution level along with the associated describedByType. At the
dataset level it’s assumed to be a human readable HTML webpage or PDF
document. Documentation that is not specifically a data dictionary
belongs in “references”
Example {"describedBy":
"http://www.agency.gov/vegetables/definitions.pdf"}
Field # describedByType
Cardinality (0,1)
Required No
Accepted Values String (IANA Media Type)
Usage Notes This is used to identify the media type (IANA Media Type
also known as MIME Type) of the URL used for the dataset’s describedBy
field. This should be specified if describedBy is not an HTML webpage.
Example {"describedByType": "application/pdf"}
Field # description
Cardinality (1,1)
Required Yes, always
Accepted Values String
Usage Notes This should be human-readable and understandable to an
average person.
Example {"description":"This dataset contains a list of vegetables,
including nutrition information and seasonality. Includes details on
tomatoes, which are really fruit but considered a vegetable in this
dataset."}
Field # distribution
Cardinality (0,n)
Required Yes, if the dataset has an accessURL or downloadURL.
Accepted Values Array of Objects
Usage Notes This is a container for one or multiple distribution
objects which group together the fields: accessURL, conformsTo,
downloadURL, describedBy, describedByType, description, format,
mediaType, and title.
Example See below
"distribution": [
{
"@type": "dcat:Distribution",
"description": "Vegetable data as a CSV file",
"downloadURL": "http://www.agency.gov/vegetables/listofvege
tables.csv",
"format": "CSV",
"mediaType": "text/csv",
"title": "listofvegetables.csv"
},
{
"@type": "dcat:Distribution",
"conformsTo": "http://www.agency.gov/vegetables-data-standa
rd/",
"describedBy": "http://www.agency.gov/vegetables/schema.xsd
",
"describedByType": "text/xml",
"description": "Vegetable data as an XML file",
"downloadURL": "http://www.agency.gov/vegetables/listofvege
tables.xml",
"format": "XML",
"mediaType": "text/xml",
"title": "listofvegetables.xml"
},
{
"@type": "dcat:Distribution",
"description": "Vegetable data as a zipped CSV file with at
tached data dictionary",
"downloadURL": "http://www.agency.gov/vegetables/vegetables
-all.zip",
"format": "Zipped CSV",
"mediaType": "application/zip",
"title": "vegetables-all.zip"
},
{
"@type": "dcat:Distribution",
"accessURL": "http://www.agency.gov/api/vegetables/",
"description": "A fully queryable REST API with JSON and XM
L output",
"format": "API",
"title": "Vegetables REST API"
}
]

Field # distribution → @type
Cardinality (0,1)
Required No
Accepted Values String (IRI)
Usage Notes The metadata type as defined by JSON-LD data types. This
should be dcat:Distribution for each distribution
Example {"@type": "dcat:Distribution"}
Field # distribution → accessURL
Cardinality (0,1)
Required Yes, if the file is accessible indirectly, through means other
than direct download.
Accepted Values String (URL)
Usage Notes This should be the URL for an indirect means of accessing
the data, such as API documentation, a ‘wizard’ or other graphical
interface which is used to generate a download, feed, or a request form
for the data. When accessLevel is “restricted public” but the dataset
is available online indirectly, this field should be the URL that
provides indirect access. This should not be a direct download URL. It
is usually assumed that accessURL is an HTML webpage.
Example {"accessURL":"http://www.agency.gov/api/vegetables/"}
Field # distribution → conformsTo
Cardinality (0,1)
Required No
Accepted Values String (URI)
Usage Notes This is used to identify a standardized specification the
distribution conforms to. It’s recommended that this be a URI that
serves as a unique identifier for the standard. The URI may or may not
also be a URL that provides documentation of the specification.
Example {"conformsTo":
"http://www.agency.gov/vegetables-data-standard/"}
Field # distribution → downloadURL
Cardinality (0,1)
Required Yes, if the file is available for public download.
Accepted Values String (URL)
Usage Notes This must be the direct download URL. Other means of
accessing the dataset should be expressed using accessURL. This should
always be accompanied by mediaType.
Example
{"downloadURL":"http://www.agency.gov/vegetables/listofvegetables.csv"}
Field # distribution → describedBy
Cardinality (0,1)
Required No
Accepted Values String (URL)
Usage Notes This is used to specify a data dictionary or schema that
defines fields or column headings in the distribution. If this is a
machine readable file the media type should be specified with
describedByType - otherwise it’s assumed to be a human readable HTML
webpage.
Example {"describedBy": "http://www.agency.gov/vegetables/schema.json"}
Field # distribution → describedByType
Cardinality (0,1)
Required No
Accepted Values String (IANA Media Type)
Usage Notes This is used to identify the media type (IANA Media Type
also known as MIME Type) of the URL used for the distribution’s
describedBy field. This is especially important if describedBy is a
machine readable file.
Example {"describedByType": "application/schema+json"}
Field # distribution → description
Cardinality (0,1)
Required No
Accepted Values String
Usage Notes This should be a human-readable description of the
distribution.
Example {"description":"Vegetable data as a zipped CSV file with
attached data dictionary"}
Field # distribution → format
Cardinality (0,1)
Required No
Accepted Values String
Usage Notes This should be a human-readable description of the file
format of the dataset, that provides useful information that might not
be apparent from mediaType. Note that API should always be used to
distinguish web APIs.
Example {"format":"CSV"}
Field # distribution → mediaType
Cardinality (0,1)
Required Yes, if the file is available for public download.
Accepted Values String (IANA Media Type)
Usage Notes This must describe the exact files available at downloadURL
using a media type (IANA Media Type also known as MIME Type). For
common Microsoft Office files, see Office Open XML MIME types
Example {"mediaType":"text/csv"}
Field # distribution → title
Cardinality (0,1)
Required No
Accepted Values String
Usage Notes This should be a useful title for the distribution.
Acronyms should be avoided.
Example {"title":"listofvegetables.csv"}
Field # identifier
Cardinality (1,1)
Required Yes, always
Accepted Values String
Usage Notes This field allows third parties to maintain a consistent
record for datasets even if title or URLs are updated. Agencies may
integrate an existing system for maintaining unique identifiers. Each
identifier must be unique across the agency’s catalog and remain fixed.
It is highly recommended that a URI (preferably an HTTP URL) be used to
provide a globally unique identifier. Identifier URLs should be
designed and maintained to persist indefinitely regardless of whether
the URL of the resource itself changes.
Example {"identifier":"http://dx.doi.org/10.7927/H4PZ56R2"}
Field # isPartOf
Cardinality (0,1)
Required No
Accepted Values String
Usage Notes This field allows the grouping of multiple datasets into a
“collection”. This field should be employed by the individual datasets
that together make up a collection. The value for this field should
match the identifier of the parent dataset.
Example {"isPartOf":"http://dx.doi.org/10.7927/H4PZ56R2"}
Field # issued
Cardinality (0,1)
Required No
Accepted Values ISO 8601 Date
Usage Notes Dates should be ISO 8601 of least resolution. In other
words, as much of YYYY-MM-DDThh:mm:ss.sTZD as is relevant to this
dataset.
Example {"issued":"2001-01-15"}
Field # keyword
Cardinality (1,n)
Required Yes, always
Accepted Values Array of strings
Usage Notes Surround each keyword with quotes. Separate keywords with
commas. Avoid duplicate keywords in the same record.
Example
{"keyword":["vegetables","veggies","greens","leafy","spinach","kale","n
utrition"]}
Field # landingPage
Cardinality (0,1)
Required No
Accepted Values String (URL)
Usage Notes This field is not intended for an agency’s homepage (e.g.
www.agency.gov), but rather if a dataset has a human-friendly hub or
landing page that users can be directed to for all resources tied to
the dataset.
Example {"landingPage":"http://www.agency.gov/vegetables"}
Field # language
Cardinality (0,n)
Required No
Accepted Values Array of strings
Usage Notes This should adhere to the RFC 5646 standard. This language
subtag lookup provides a good tool for checking and verifying language
codes. A language tag is comprised of either one or two parts, the
language subtag (such as en for English, sp for Spanish, wo for Wolof)
and the regional subtag (such as US for United States, GB for Great
Britain, MX for Mexico), separated by a hyphen. Regional subtags should
only be provided when needed to distinguish a language tag from another
one (such as American vs. British English).
Example {"language":["en-US"]} or if multiple languages,
{"language":["es-MX","wo","nv","en-US"]}
Field # license
Cardinality (0,1)
Required Yes, if applicable
Accepted Values String (URL)
Usage Notes See list of license-free declarations and licenses.
Example {"license":"http://creativecommons.org/publicdomain/zero/1.0/"}
Field # modified
Cardinality (1,1)
Required Yes, always
Accepted Values ISO 8601 Date
Usage Notes Dates should be ISO 8601 of highest resolution. In other
words, as much of YYYY-MM-DDThh:mm:ss.sTZD as is relevant to this
dataset. If there is a need to reflect that the dataset is continually
updated, ISO 8601 formatting can account for this with repeating
intervals. For instance, R/P1D for daily, R/P2W for every two weeks,
and R/PT5M for every five minutes.
+Example {"modified":"2012-01-15"} or {"modified":"R/P1D"}
Field # primaryITInvestmentUII
Cardinality (0,1)
Required No
Accepted Values String
Usage Notes Use to link a given dataset with its related IT Unique
Investment Identifier, which can often be found in Exhibit 53
documents.
Example {"primaryITInvestmentUII":"023-000000001"}
Field # programCode
Cardinality (0,n)
Required Yes, for United States Federal Government Agencies
Accepted Values Array of strings
Usage Notes Provide an array of programs related to this data asset,
from the Federal Program Inventory.
Example {"programCode":["015:001"]} or if multiple programs,
{"programCode":["015:001","015:002"]}
Field # publisher
Cardinality (1,1)
Required Yes, always
Accepted Values Object
Usage Notes This is a container for a publisher object which groups
together the fields: name and subOrganization. The subOrganization
field can also contain a publisher object which allows one to describe
an organization’s hierarchy. Where greater specificity is desired,
include as many levels of publisher as is useful, in ascending order,
using the below format.
Example See below
"publisher": {
"@type": "org:Organization",
"name": "Widget Services",
"subOrganizationOf": {
"@type": "org:Organization",
"name": "Office of Citizen Services and Innovative Technologies",
"subOrganizationOf": {
"@type": "org:Organization",
"name": "General Services Administration",
"subOrganizationOf": {
"@type": "org:Organization",
"name": "U.S. Government"
}
}
}
}

Field # publisher → @type
Cardinality (0,1)
Required No
Accepted Values String (IRI)
Usage Notes The metadata type as defined by JSON-LD data types. This
should be org:Organization for each publisher
Example {"@type": "org:Organization"}
Field # publisher → name
Cardinality (1,1)
Required Yes, always
Accepted Values String
Usage Notes The plaintext name of the entity publishing this dataset.
Example {"name": "U.S. Department of Commerce"}
Field # publisher → subOrganizationOf
Cardinality (0,1)
Required No
Accepted Values publisher object
Usage Notes A parent organizational entity described using the same
publisher object fields.
Example "subOrganizationOf": {"name": "General Services
Administration", "subOrganizationOf": {"name": "U.S. Government"}}
Field # references
Cardinality (0,n)
Required No
Accepted Values Array of strings (URLs)
Usage Notes Enclose each URL within strings. Separate multiple URLs
with a comma.
Example
{"references":["http://www.agency.gov/legumes/legumes_data_documentatio
n.html"]} or if multiple URLs,
{"references":["http://www.agency.gov/legumes/legumes_data_documentatio
n.html","http://www.agency.gov/fruits/fruit_data_documentation.html"]}
Field # rights
Cardinality (0,1)
Required Yes, if accessLevel is “restricted public” or “non-public”
Accepted Values String
Usage Notes This may include information regarding access or
restrictions based on privacy, security, or other policies. This should
also serve as an explanation for the selected “accessLevel” including
instructions for how to access a restricted file, if applicable, or
explanation for why a “non-public” or “restricted public” data asset is
not “public,” if applicable. If the dataset can be made available
through a website indirectly, use accessURL for the URL that provides
such access.
Example {"rights":"This dataset contains Personally Identifiable
Information and could not be released for public access."}
Field # spatial
Cardinality (0,1)
Required Yes, if the dataset is spatial
Accepted Values See Usage Notes
Usage Notes This field supports either Strings or GeoJSON Objects. As a
string; This field should contain one of the following types of
content: (1) a bounding coordinate box for the dataset represented in
latitude / longitude pairs where the coordinates are specified in
decimal degrees and in the order of: minimum longitude, minimum
latitude, maximum longitude, maximum latitude; (2) a latitude /
longitude pair (in decimal degrees) representing a point where the
dataset is relevant; (3) a geographic feature from the GeoNames
database. (4) As GeoJSON Objects we support the following "types":
"Polygon" or "Point". (4.1) For a "Polygon" type, the "coordinates"
field must be an array of arrays, of points, where the array defines
the outer boundary of the "Polygon". This array should contain at least
four [longitude, latitude] points (as numbers, either integers or
floating-point numbers), and the first and last points must be the same
to close the shape.; or (4.2) "Point" objects where the "coordinates"
field is an array [longitude, latitude] expressed (as numbers, either
integers or floating-point numbers). (4.3) We also support GeoJSON if
it is also expressed as a string. Currently we do not support
geographic features expressed in Geography Markup Language using the
Simple Features Profile. Other JSON types and/or strings will pass
validation (such as “envelope”), however they may not be spatially
searchable or have a map of the dataset bounding region.
Example See Below
(1) {"spatial": "137.5488,3.8128,163.3647,10.2284"}
(2) {"spatial": "-88.9718,36.52033"}
(3) {"spatial":"Lincoln, Nebraska"}
(4.1) {"spatial":{"type":"Polygon","coordinates":[[[-77.119759,38.791645],[-76.9
09393,38.791645],[-76.909393,38.995548],[-77.119759,38.995548],[-77.119759,38.79
1645]]]}}
(4.2) {"spatial":{"type":"Point","coordinates":[-77.0369,38.9072]}}
(4.3) {"spatial":"{\"type\": \"Point\", \"coordinates\": [-80.162, 25.731333]}"}

Field # systemOfRecords
Cardinality (0,1)
Required No
Accepted Values String (URL)
Usage Notes This field should contain a URL to the System of Records
Notice (SORN) that relates to the dataset, specifically from
FederalRegister.gov.
Example
{"systemOfRecords":"https://www.federalregister.gov/articles/2002/04/08
/02-7376/privacy-act-of-1974-publication-in-full-of-all-notices-of-syst
ems-of-records-including-several-new#p-361"}
Field # temporal
Cardinality (0,1)
Required Yes, if applicable
Accepted Values ISO 8601 Date
Usage Notes This field should contain an interval of time defined by
the start and end dates for which the dataset is applicable. Dates
should be formatted as pairs of {start datetime/end datetime} in the
ISO 8601 format. ISO 8601 specifies that datetimes can be formatted in
a number of ways, including a simple four-digit year (eg. 2013) to a
much more specific YYYY-MM-DDTHH:MM:SSZ, where the T specifies a
seperator between the date and time and time is expressed in 24 hour
notation in the UTC (Zulu) time zone. (e.g.,
2011-02-14T12:00:00Z/2013-07-04T19:34:00Z). Use a solidus (“/”) to
separate start and end times. If there is a need to define the start or
end of applicability using a duration rather than a date, ISO 8601
formatting can account for this with duration based intervals. For
instance, applicability starting in January 2010 and continuing for one
month could be represented as 2010-01/P1M or 2010-01/2010-02. However,
when possible, full dates are preferred for both start and end times.
Example {"temporal":"2000-01-15T00:45:00Z/2010-01-15T00:06:00Z"} or
{"temporal":"2000-01-15T00:45:00Z/P1W"}
Field # theme
Cardinality (0,n)
Required No
Accepted Values Array of strings
Usage Notes Separate multiple categories with a comma. Could include
ISO Topic Categories.
Examples {"theme":["vegetables"]} or if multiple categories,
{"theme":["vegetables","produce"]}
Field # title
Cardinality (1,1)
Required Yes, always
Accepted Values String
Usage Notes Acronyms should be avoided.
Example {"title":"Types of Vegetables"}

Federal Government Fields
__________________________________________________________________

^USG — Fields specific to the U.S. Federal Government have been denoted
with the ^USG superscript. The Project Open Data schema has been
developed as part of a U.S Federal Government open data policy.
However, every attempt has been made to align the schema with existing
international standards and to provide opportunities for re-use and
interoperability with state and local government as well as
non-profits, academic institutions, and businesses. There are however
some fields that have been introduced specifically for use by the U.S.
Federal Government and have special meaning in that context. These
fields are: bureauCode, programCode, dataQuality,
primaryITInvestmentUII, and systemOfRecords. Non-federal data
publishers are encouraged to make use of this schema, but these fields
should not be seen as required and may not be relevant for those
entities.

Rationale for Metadata Nomenclature

We sought to be platform-independent and to align as much as possible
with existing open standards.

To that end, our JSON key names are directly drawn from DCAT, with a
few exceptions.

We added the accessLevel field to help easily sort datasets into our
three existing categories: public, restricted public, and non-public.
This field means an agency can run a basic filter against its
enterprise data catalog to generate a public-facing list of datasets
that are, or could one day be, made publicly available (or, in the case
of restricted data, available under certain conditions). This field
also makes it easy for anyone to generate a list of datasets that could
be made available but have not yet been released by filtering
accessLevel to public and accessURL to blank.

We added the rights field (formerly accessLevelComment) for data
stewards to explain how to access restricted public datasets, and for
agencies to have a place to record (even if only internally) the reason
for not releasing a non-public dataset.

We added the systemOfRecords field for data stewards to optionally link
to a relevant System of Records Notice URL. A System of Records is a
group of any records under the control of any agency from which
information is retrieved by the name of the individual or by some
identifying number, symbol, or other identifier assigned to the
individual.

We added the bureauCode field to ensure every dataset is connected in a
standard way with an agency bureau.

We added the programCode field to ensure that when applicable, every
dataset is connected in a standard way with an agency program office.

We added the dataQuality to indicate whether or not the data meets an
agency’s Information Quality Guidelines.

Additional Information
__________________________________________________________________

* Metadata Resources (including starter template and sample files)
* DCAT

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
