---
title: WHATWG HTML Standard
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://html.spec.whatwg.org/multipage/
tags: [web-standards, web-corpus]
confidence: medium
---

# WHATWG HTML Standard

## Fetch Metadata

- Corpus item: 57
- Category: Web Standards
- Source URL: https://html.spec.whatwg.org/multipage/
- Final URL: https://html.spec.whatwg.org/multipage/
- Retrieved: 2026-05-18
- Content-Type: text/html; charset=utf-8
- Fetch status: ok
- Test value: Very long living standard.
- Fetched page title: HTML Standard

## Retrieved Text

WHATWG

HTML

Living Standard — Last Updated 17 May 2026

One-Page Version html.spec.whatwg.org Multipage Version /multipage
Version for Web Devs /dev PDF Version /print.pdf Translations 日本語 •
简体中文

FAQ on GitHub Chat on Matrix

Contribute on GitHub whatwg/html repository Commits on GitHub Snapshot
as of this commit Twitter Updates @htmlstandard

Open Issues filed on GitHub Open an Issue whatwg.org/newbug Tests
web-platform-tests html/ Issues for Tests ongoing work
__________________________________________________________________

Table of contents

1. 1 Introduction
2. 2 Common infrastructure
3. 3 Semantics, structure, and APIs of HTML documents
4. 4 The elements of HTML
5. 5 Microdata
6. 6 User interaction
7. 7 Loading web pages
8. 8 Web application APIs
9. 9 Communication
10. 10 Web workers
11. 11 Worklets
12. 12 Web storage
13. 13 The HTML syntax
14. 14 The XML syntax
15. 15 Rendering
16. 16 Obsolete features
17. 17 IANA considerations
18. Index
19. References
20. Acknowledgments
21. Intellectual property rights

Full table of contents

1. 1 Introduction
1. 1.1 Where does this specification fit?
2. 1.2 Is this HTML5?
3. 1.3 Background
4. 1.4 Audience
5. 1.5 Scope
6. 1.6 History
7. 1.7 Design notes
1. 1.7.1 Serializability of script execution
2. 1.7.2 Extensibility
8. 1.8 HTML vs XML syntax
9. 1.9 Structure of this specification
1. 1.9.1 How to read this specification
2. 1.9.2 Typographic conventions
10. 1.10 A quick introduction to HTML
1. 1.10.1 Writing secure applications with HTML
2. 1.10.2 Common pitfalls to avoid when using the scripting
APIs
3. 1.10.3 How to catch mistakes when writing HTML:
validators and conformance checkers
11. 1.11 Conformance requirements for authors
1. 1.11.1 Presentational markup
2. 1.11.2 Syntax errors
3. 1.11.3 Restrictions on content models and on attribute
values
12. 1.12 Suggested reading
2. 2 Common infrastructure
1. 2.1 Terminology
1. 2.1.1 Parallelism
2. 2.1.2 Resources
3. 2.1.3 XML compatibility
4. 2.1.4 DOM trees
5. 2.1.5 Scripting
6. 2.1.6 Plugins
7. 2.1.7 Character encodings
8. 2.1.8 Conformance classes
9. 2.1.9 Dependencies
10. 2.1.10 Extensibility
11. 2.1.11 Interactions with XPath and XSLT
2. 2.2 Policy-controlled features
3. 2.3 Common microsyntaxes
1. 2.3.1 Common parser idioms
2. 2.3.2 Boolean attributes
3. 2.3.3 Keywords and enumerated attributes
4. 2.3.4 Numbers
1. 2.3.4.1 Signed integers
2. 2.3.4.2 Non-negative integers
3. 2.3.4.3 Floating-point numbers
4. 2.3.4.4 Percentages and lengths
5. 2.3.4.5 Nonzero percentages and lengths
6. 2.3.4.6 Lists of floating-point numbers
7. 2.3.4.7 Lists of dimensions
5. 2.3.5 Dates and times
1. 2.3.5.1 Months
2. 2.3.5.2 Dates
3. 2.3.5.3 Yearless dates
4. 2.3.5.4 Times
5. 2.3.5.5 Local dates and times
6. 2.3.5.6 Time zones
7. 2.3.5.7 Global dates and times
8. 2.3.5.8 Weeks
9. 2.3.5.9 Durations
10. 2.3.5.10 Vaguer moments in time
6. 2.3.6 Legacy colors
7. 2.3.7 Space-separated tokens
8. 2.3.8 Comma-separated tokens
9. 2.3.9 References
10. 2.3.10 Media queries
11. 2.3.11 Unique internal values
4. 2.4 URLs
1. 2.4.1 Terminology
2. 2.4.2 Parsing URLs
3. 2.4.3 Document base URLs
5. 2.5 Fetching resources
1. 2.5.1 Terminology
2. 2.5.2 Determining the type of a resource
3. 2.5.3 Extracting character encodings from meta elements
4. 2.5.4 CORS settings attributes
5. 2.5.5 Referrer policy attributes
6. 2.5.6 Nonce attributes
7. 2.5.7 Lazy loading attributes
8. 2.5.8 Blocking attributes
9. 2.5.9 Fetch priority attributes
6. 2.6 Common DOM interfaces
1. 2.6.1 Reflecting content attributes in IDL attributes
2. 2.6.2 Using reflect via IDL extended attributes
3. 2.6.3 Using reflect in specifications
4. 2.6.4 Collections
1. 2.6.4.1 The HTMLAllCollection interface
1. 2.6.4.1.1 [[Call]] ( thisArgument,
argumentsList )
2. 2.6.4.2 The HTMLFormControlsCollection interface
3. 2.6.4.3 The HTMLOptionsCollection interface
5. 2.6.5 The DOMStringList interface
7. 2.7 Safe passing of structured data
1. 2.7.1 Serializable objects
2. 2.7.2 Transferable objects
3. 2.7.3 StructuredSerializeInternal ( value, forStorage [ ,
memory ] )
4. 2.7.4 StructuredSerialize ( value )
5. 2.7.5 StructuredSerializeForStorage ( value )
6. 2.7.6 StructuredDeserialize ( serialized, targetRealm [ ,
memory ] )
7. 2.7.7 StructuredSerializeWithTransfer ( value,
transferList )
8. 2.7.8 StructuredDeserializeWithTransfer (
serializeWithTransferResult, targetRealm )
9. 2.7.9 Performing serialization and transferring from
other specifications
10. 2.7.10 Structured cloning API
3. 3 Semantics, structure, and APIs of HTML documents
1. 3.1 Documents
1. 3.1.1 The Document object
2. 3.1.2 The DocumentOrShadowRoot interface
3. 3.1.3 Ancestor origins
4. 3.1.4 Resource metadata management
5. 3.1.5 Reporting document loading status
6. 3.1.6 Render-blocking mechanism
7. 3.1.7 DOM tree accessors
2. 3.2 Elements
1. 3.2.1 Semantics
2. 3.2.2 Elements in the DOM
3. 3.2.3 HTML element constructors
4. 3.2.4 Element definitions
1. 3.2.4.1 Attributes
5. 3.2.5 Content models
1. 3.2.5.1 The "nothing" content model
2. 3.2.5.2 Kinds of content
1. 3.2.5.2.1 Metadata content
2. 3.2.5.2.2 Flow content
3. 3.2.5.2.3 Sectioning content
4. 3.2.5.2.4 Heading content
5. 3.2.5.2.5 Phrasing content
6. 3.2.5.2.6 Embedded content
7. 3.2.5.2.7 Interactive content
8. 3.2.5.2.8 Palpable content
9. 3.2.5.2.9 Script-supporting elements
10. 3.2.5.2.10 select element inner content
elements
11. 3.2.5.2.11 optgroup element inner content
elements
12. 3.2.5.2.12 option element inner content
elements
3. 3.2.5.3 Transparent content models
4. 3.2.5.4 Paragraphs
6. 3.2.6 Global attributes
1. 3.2.6.1 The title attribute
2. 3.2.6.2 The lang and xml:lang attributes
3. 3.2.6.3 The translate attribute
4. 3.2.6.4 The dir attribute
5. 3.2.6.5 The style attribute
6. 3.2.6.6 Embedding custom non-visible data with the
data-* attributes
7. 3.2.7 The innerText and outerText properties
8. 3.2.8 Requirements relating to the bidirectional
algorithm
1. 3.2.8.1 Authoring conformance criteria for
bidirectional-algorithm formatting characters
2. 3.2.8.2 User agent conformance criteria
9. 3.2.9 Requirements related to ARIA and to platform
accessibility APIs
4. 4 The elements of HTML
1. 4.1 The document element
1. 4.1.1 The html element
2. 4.2 Document metadata
1. 4.2.1 The head element
2. 4.2.2 The title element
3. 4.2.3 The base element
4. 4.2.4 The link element
1. 4.2.4.1 Processing the media attribute
2. 4.2.4.2 Processing the type attribute
3. 4.2.4.3 Fetching and processing a resource from a
link element
4. 4.2.4.4 Processing `Link` headers
5. 4.2.4.5 Early hints
6. 4.2.4.6 Providing users with a means to follow
hyperlinks created using the link element
5. 4.2.5 The meta element
1. 4.2.5.1 Standard metadata names
2. 4.2.5.2 Other metadata names
3. 4.2.5.3 Pragma directives
4. 4.2.5.4 Specifying the document's character encoding
6. 4.2.6 The style element
7. 4.2.7 Interactions of styling and scripting
3. 4.3 Sections
1. 4.3.1 The body element
2. 4.3.2 The article element
3. 4.3.3 The section element
4. 4.3.4 The nav element
5. 4.3.5 The aside element
6. 4.3.6 The h1, h2, h3, h4, h5, and h6 elements
7. 4.3.7 The hgroup element
8. 4.3.8 The header element
9. 4.3.9 The footer element
10. 4.3.10 The address element
11. 4.3.11 Headings and outlines
1. 4.3.11.1 Heading levels & offsets
2. 4.3.11.2 Sample outlines
3. 4.3.11.3 Exposing outlines to users
12. 4.3.12 Usage summary
1. 4.3.12.1 Article or section?
4. 4.4 Grouping content
1. 4.4.1 The p element
2. 4.4.2 The hr element
3. 4.4.3 The pre element
4. 4.4.4 The blockquote element
5. 4.4.5 The ol element
6. 4.4.6 The ul element
7. 4.4.7 The menu element
8. 4.4.8 The li element
9. 4.4.9 The dl element
10. 4.4.10 The dt element
11. 4.4.11 The dd element
12. 4.4.12 The figure element
13. 4.4.13 The figcaption element
14. 4.4.14 The main element
15. 4.4.15 The search element
16. 4.4.16 The div element
5. 4.5 Text-level semantics
1. 4.5.1 The a element
2. 4.5.2 The em element
3. 4.5.3 The strong element
4. 4.5.4 The small element
5. 4.5.5 The s element
6. 4.5.6 The cite element
7. 4.5.7 The q element
8. 4.5.8 The dfn element
9. 4.5.9 The abbr element
10. 4.5.10 The ruby element
11. 4.5.11 The rt element
12. 4.5.12 The rp element
13. 4.5.13 The data element
14. 4.5.14 The time element
15. 4.5.15 The code element
16. 4.5.16 The var element
17. 4.5.17 The samp element
18. 4.5.18 The kbd element
19. 4.5.19 The sub and sup elements
20. 4.5.20 The i element
21. 4.5.21 The b element
22. 4.5.22 The u element
23. 4.5.23 The mark element
24. 4.5.24 The bdi element
25. 4.5.25 The bdo element
26. 4.5.26 The span element
27. 4.5.27 The br element
28. 4.5.28 The wbr element
29. 4.5.29 Usage summary
6. 4.6 Links
1. 4.6.1 Introduction
2. 4.6.2 Links created by a and area elements
3. 4.6.3 API for hyperlink elements
4. 4.6.4 API for a and area elements
5. 4.6.5 Following hyperlinks
6. 4.6.6 Downloading resources
7. 4.6.7 Hyperlink auditing
1. 4.6.7.1 The `Ping-From` and `Ping-To` headers
8. 4.6.8 Link types
1. 4.6.8.1 Link type "alternate"
2. 4.6.8.2 Link type "author"
3. 4.6.8.3 Link type "bookmark"
4. 4.6.8.4 Link type "canonical"
5. 4.6.8.5 Link type "dns-prefetch"
6. 4.6.8.6 Link type "expect"
7. 4.6.8.7 Link type "external"
8. 4.6.8.8 Link type "help"
9. 4.6.8.9 Link type "icon"
10. 4.6.8.10 Link type "license"
11. 4.6.8.11 Link type "manifest"
12. 4.6.8.12 Link type "modulepreload"
13. 4.6.8.13 Link type "nofollow"
14. 4.6.8.14 Link type "noopener"
15. 4.6.8.15 Link type "noreferrer"
16. 4.6.8.16 Link type "opener"
17. 4.6.8.17 Link type "pingback"
18. 4.6.8.18 Link type "preconnect"
19. 4.6.8.19 Link type "prefetch"
20. 4.6.8.20 Link type "preload"
21. 4.6.8.21 Link type "privacy-policy"
22. 4.6.8.22 Link type "search"
23. 4.6.8.23 Link type "stylesheet"
24. 4.6.8.24 Link type "tag"
25. 4.6.8.25 Link Type "terms-of-service"
26. 4.6.8.26 Sequential link types
1. 4.6.8.26.1 Link type "next"
2. 4.6.8.26.2 Link type "prev"
27. 4.6.8.27 Other link types
7. 4.7 Edits
1. 4.7.1 The ins element
2. 4.7.2 The del element
3. 4.7.3 Attributes common to ins and del elements
4. 4.7.4 Edits and paragraphs
5. 4.7.5 Edits and lists
6. 4.7.6 Edits and tables
8. 4.8 Embedded content
1. 4.8.1 The picture element
2. 4.8.2 The source element
3. 4.8.3 The img element
4. 4.8.4 Images
1. 4.8.4.1 Introduction
1. 4.8.4.1.1 Adaptive images
2. 4.8.4.2 Attributes common to source, img, and link
elements
1. 4.8.4.2.1 Srcset attributes
2. 4.8.4.2.2 Sizes attributes
3. 4.8.4.3 Processing model
1. 4.8.4.3.1 When to obtain images
2. 4.8.4.3.2 Reacting to DOM mutations
3. 4.8.4.3.3 The list of available images
4. 4.8.4.3.4 Decoding images
5. 4.8.4.3.5 Updating the image data
6. 4.8.4.3.6 Preparing an image for presentation
7. 4.8.4.3.7 Selecting an image source
8. 4.8.4.3.8 Creating a source set from attributes
9. 4.8.4.3.9 Updating the source set
10. 4.8.4.3.10 Parsing a srcset attribute
11. 4.8.4.3.11 Parsing a sizes attribute
12. 4.8.4.3.12 Normalizing the source densities
13. 4.8.4.3.13 Reacting to environment changes
4. 4.8.4.4 Requirements for providing text to act as an
alternative for images
1. 4.8.4.4.1 General guidelines
2. 4.8.4.4.2 A link or button containing nothing
but the image
3. 4.8.4.4.3 A phrase or paragraph with an
alternative graphical representation: charts,
diagrams, graphs, maps, illustrations
4. 4.8.4.4.4 A short phrase or label with an
alternative graphical representation: icons,
logos
5. 4.8.4.4.5 Text that has been rendered to a
graphic for typographical effect
6. 4.8.4.4.6 A graphical representation of some of
the surrounding text
7. 4.8.4.4.7 Ancillary images
8. 4.8.4.4.8 A purely decorative image that
doesn't add any information
9. 4.8.4.4.9 A group of images that form a single
larger picture with no links
10. 4.8.4.4.10 A group of images that form a single
larger picture with links
11. 4.8.4.4.11 A key part of the content
12. 4.8.4.4.12 An image not intended for the user
13. 4.8.4.4.13 An image in an email or private
document intended for a specific person who is
known to be able to view images
14. 4.8.4.4.14 Guidance for markup generators
15. 4.8.4.4.15 Guidance for conformance checkers
5. 4.8.5 The iframe element
6. 4.8.6 The embed element
7. 4.8.7 The object element
8. 4.8.8 The video element
9. 4.8.9 The audio element
10. 4.8.10 The track element
11. 4.8.11 Media elements
1. 4.8.11.1 Error codes
2. 4.8.11.2 Location of the media resource
3. 4.8.11.3 MIME types
4. 4.8.11.4 Network states
5. 4.8.11.5 Loading the media resource
6. 4.8.11.6 Offsets into the media resource
7. 4.8.11.7 Ready states
8. 4.8.11.8 Playing the media resource
9. 4.8.11.9 Seeking
10. 4.8.11.10 Media resources with multiple media tracks
1. 4.8.11.10.1 AudioTrackList and VideoTrackList
objects
2. 4.8.11.10.2 Selecting specific audio and video
tracks declaratively
11. 4.8.11.11 Timed text tracks
1. 4.8.11.11.1 Text track model
2. 4.8.11.11.2 Sourcing in-band text tracks
3. 4.8.11.11.3 Sourcing out-of-band text tracks
4. 4.8.11.11.4 Guidelines for exposing cues in
various formats as text track cues
5. 4.8.11.11.5 Text track API
6. 4.8.11.11.6 Event handlers for objects of the
text track APIs
7. 4.8.11.11.7 Best practices for metadata text
tracks
12. 4.8.11.12 Identifying a track kind through a URL
13. 4.8.11.13 User interface
14. 4.8.11.14 Time ranges
15. 4.8.11.15 The TrackEvent interface
16. 4.8.11.16 Events summary
17. 4.8.11.17 Security and privacy considerations
18. 4.8.11.18 Best practices for authors using media
elements
19. 4.8.11.19 Best practices for implementers of media
elements
12. 4.8.12 The map element
13. 4.8.13 The area element
14. 4.8.14 Image maps
1. 4.8.14.1 Authoring
2. 4.8.14.2 Processing model
15. 4.8.15 MathML
16. 4.8.16 SVG
17. 4.8.17 Dimension attributes
9. 4.9 Tabular data
1. 4.9.1 The table element
1. 4.9.1.1 Techniques for describing tables
2. 4.9.1.2 Techniques for table design
2. 4.9.2 The caption element
3. 4.9.3 The colgroup element
4. 4.9.4 The col element
5. 4.9.5 The tbody element
6. 4.9.6 The thead element
7. 4.9.7 The tfoot element
8. 4.9.8 The tr element
9. 4.9.9 The td element
10. 4.9.10 The th element
11. 4.9.11 Attributes common to td and th elements
12. 4.9.12 Processing model
1. 4.9.12.1 Forming a table
2. 4.9.12.2 Forming relationships between data cells
and header cells
13. 4.9.13 Examples
10. 4.10 Forms
1. 4.10.1 Introduction
1. 4.10.1.1 Writing a form's user interface
2. 4.10.1.2 Implementing the server-side processing for
a form
3. 4.10.1.3 Configuring a form to communicate with a
server
4. 4.10.1.4 Client-side form validation
5. 4.10.1.5 Enabling client-side automatic filling of
form controls
6. 4.10.1.6 Improving the user experience on mobile
devices
7. 4.10.1.7 The difference between the field type, the
autofill field name, and the input modality
8. 4.10.1.8 Date, time, and number formats
2. 4.10.2 Categories
3. 4.10.3 The form element
4. 4.10.4 The label element
5. 4.10.5 The input element
1. 4.10.5.1 States of the type attribute
1. 4.10.5.1.1 Hidden state (type=hidden)
2. 4.10.5.1.2 Text (type=text) state and Search
state (type=search)
3. 4.10.5.1.3 Telephone state (type=tel)
4. 4.10.5.1.4 URL state (type=url)
5. 4.10.5.1.5 Email state (type=email)
6. 4.10.5.1.6 Password state (type=password)
7. 4.10.5.1.7 Date state (type=date)
8. 4.10.5.1.8 Month state (type=month)
9. 4.10.5.1.9 Week state (type=week)
10. 4.10.5.1.10 Time state (type=time)
11. 4.10.5.1.11 Local Date and Time state
(type=datetime-local)
12. 4.10.5.1.12 Number state (type=number)
13. 4.10.5.1.13 Range state (type=range)
14. 4.10.5.1.14 Color state (type=color)
15. 4.10.5.1.15 Checkbox state (type=checkbox)
16. 4.10.5.1.16 Radio Button state (type=radio)
17. 4.10.5.1.17 File Upload state (type=file)
18. 4.10.5.1.18 Submit Button state (type=submit)
19. 4.10.5.1.19 Image Button state (type=image)
20. 4.10.5.1.20 Reset Button state (type=reset)
21. 4.10.5.1.21 Button state (type=button)
2. 4.10.5.2 Implementation notes regarding localization
of form controls
3. 4.10.5.3 Common input element attributes
1. 4.10.5.3.1 The maxlength and minlength
attributes
2. 4.10.5.3.2 The size attribute
3. 4.10.5.3.3 The readonly attribute
4. 4.10.5.3.4 The required attribute
5. 4.10.5.3.5 The multiple attribute
6. 4.10.5.3.6 The pattern attribute
7. 4.10.5.3.7 The min and max attributes
8. 4.10.5.3.8 The step attribute
9. 4.10.5.3.9 The list attribute
10. 4.10.5.3.10 The placeholder attribute
4. 4.10.5.4 Common input element APIs
5. 4.10.5.5 Common event behaviors
6. 4.10.6 The button element
7. 4.10.7 The select element
8. 4.10.8 The datalist element
9. 4.10.9 The optgroup element
10. 4.10.10 The option element
11. 4.10.11 The textarea element
12. 4.10.12 The output element
13. 4.10.13 The progress element
14. 4.10.14 The meter element
15. 4.10.15 The fieldset element
16. 4.10.16 The legend element
17. 4.10.17 The selectedcontent element
18. 4.10.18 Form control infrastructure
1. 4.10.18.1 A form control's value
2. 4.10.18.2 Mutability
3. 4.10.18.3 Association of controls and forms
19. 4.10.19 Attributes common to form controls
1. 4.10.19.1 Naming form controls: the name attribute
2. 4.10.19.2 Submitting element directionality: the
dirname attribute
3. 4.10.19.3 Limiting user input length: the maxlength
attribute
4. 4.10.19.4 Setting minimum input length requirements:
the minlength attribute
5. 4.10.19.5 Enabling and disabling form controls: the
disabled attribute
6. 4.10.19.6 Form submission attributes
7. 4.10.19.7 Autofill
1. 4.10.19.7.1 Autofilling form controls: the
autocomplete attribute
2. 4.10.19.7.2 Processing model
20. 4.10.20 APIs for the text control selections
21. 4.10.21 Constraints
1. 4.10.21.1 Definitions
2. 4.10.21.2 Constraint validation
3. 4.10.21.3 The constraint validation API
4. 4.10.21.4 Security
22. 4.10.22 Form submission
1. 4.10.22.1 Introduction
2. 4.10.22.2 Implicit submission
3. 4.10.22.3 Form submission algorithm
4. 4.10.22.4 Constructing the entry list
5. 4.10.22.5 Selecting a form submission encoding
6. 4.10.22.6 Converting an entry list to a list of
name-value pairs
7. 4.10.22.7 URL-encoded form data
8. 4.10.22.8 Multipart form data
9. 4.10.22.9 Plain text form data
10. 4.10.22.10 The SubmitEvent interface
11. 4.10.22.11 The FormDataEvent interface
23. 4.10.23 Resetting a form
11. 4.11 Interactive elements
1. 4.11.1 The details element
2. 4.11.2 The summary element
3. 4.11.3 Commands
1. 4.11.3.1 Facets
2. 4.11.3.2 Using the a element to define a command
3. 4.11.3.3 Using the button element to define a
command
4. 4.11.3.4 Using the input element to define a command
5. 4.11.3.5 Using the option element to define a
command
6. 4.11.3.6 Using the accesskey attribute on a legend
element to define a command
7. 4.11.3.7 Using the accesskey attribute to define a
command on other elements
4. 4.11.4 The dialog element
5. 4.11.5 Dialog light dismiss
12. 4.12 Scripting
1. 4.12.1 The script element
1. 4.12.1.1 Processing model
2. 4.12.1.2 Scripting languages
3. 4.12.1.3 Restrictions for contents of script
elements
4. 4.12.1.4 Inline documentation for external scripts
5. 4.12.1.5 Interaction of script elements and XSLT
2. 4.12.2 The noscript element
3. 4.12.3 The template element
1. 4.12.3.1 Interaction of template elements with XSLT
and XPath
4. 4.12.4 The slot element
5. 4.12.5 The canvas element
1. 4.12.5.1 The 2D rendering context
1. 4.12.5.1.1 Implementation notes
2. 4.12.5.1.2 The canvas settings
3. 4.12.5.1.3 The canvas state
4. 4.12.5.1.4 Line styles
5. 4.12.5.1.5 Text styles
6. 4.12.5.1.6 Building paths
7. 4.12.5.1.7 Path2D objects
8. 4.12.5.1.8 Transformations
9. 4.12.5.1.9 Image sources for 2D rendering
contexts
10. 4.12.5.1.10 Fill and stroke styles
11. 4.12.5.1.11 Drawing rectangles to the bitmap
12. 4.12.5.1.12 Drawing text to the bitmap
13. 4.12.5.1.13 Drawing paths to the canvas
14. 4.12.5.1.14 Drawing focus rings
15. 4.12.5.1.15 Drawing images
16. 4.12.5.1.16 Pixel manipulation
17. 4.12.5.1.17 Compositing
18. 4.12.5.1.18 Image smoothing
19. 4.12.5.1.19 Shadows
20. 4.12.5.1.20 Filters
21. 4.12.5.1.21 Working with externally-defined SVG
filters
22. 4.12.5.1.22 Drawing model
23. 4.12.5.1.23 Best practices
24. 4.12.5.1.24 Examples
2. 4.12.5.2 The ImageBitmap rendering context
1. 4.12.5.2.1 Introduction
2. 4.12.5.2.2 The ImageBitmapRenderingContext
interface
3. 4.12.5.3 The OffscreenCanvas interface
1. 4.12.5.3.1 The offscreen 2D rendering context
4. 4.12.5.4 Color spaces and color space conversion
5. 4.12.5.5 Serializing bitmaps to a file
6. 4.12.5.6 Security with canvas elements
7. 4.12.5.7 Premultiplied alpha and the 2D rendering
context
13. 4.13 Custom elements
1. 4.13.1 Introduction
1. 4.13.1.1 Creating an autonomous custom element
2. 4.13.1.2 Creating a form-associated custom element
3. 4.13.1.3 Creating a custom element with default
accessible roles, states, and properties
4. 4.13.1.4 Creating a customized built-in element
5. 4.13.1.5 Drawbacks of autonomous custom elements
6. 4.13.1.6 Upgrading elements after their creation
7. 4.13.1.7 Scoped custom element registries
8. 4.13.1.8 Exposing custom element states
2. 4.13.2 Requirements for custom element constructors and
reactions
1. 4.13.2.1 Preserving custom element state when moved
3. 4.13.3 Core concepts
4. 4.13.4 The CustomElementRegistry interface
5. 4.13.5 Upgrades
6. 4.13.6 Custom element reactions
7. 4.13.7 Element internals
1. 4.13.7.1 The ElementInternals interface
2. 4.13.7.2 Shadow root access
3. 4.13.7.3 Form-associated custom elements
4. 4.13.7.4 Accessibility semantics
5. 4.13.7.5 Custom state pseudo-class
14. 4.14 Common idioms without dedicated elements
1. 4.14.1 Breadcrumb navigation
2. 4.14.2 Tag clouds
3. 4.14.3 Conversations
4. 4.14.4 Footnotes
15. 4.15 Disabled elements
16. 4.16 Matching HTML elements using selectors and CSS
1. 4.16.1 Case-sensitivity of the CSS 'attr()' function
2. 4.16.2 Case-sensitivity of selectors
3. 4.16.3 Pseudo-classes
5. 5 Microdata
1. 5.1 Introduction
1. 5.1.1 Overview
2. 5.1.2 The basic syntax
3. 5.1.3 Typed items
4. 5.1.4 Global identifiers for items
5. 5.1.5 Selecting names when defining vocabularies
2. 5.2 Encoding microdata
1. 5.2.1 The microdata model
2. 5.2.2 Items
3. 5.2.3 Names: the itemprop attribute
4. 5.2.4 Values
5. 5.2.5 Associating names with items
6. 5.2.6 Microdata and other namespaces
3. 5.3 Sample microdata vocabularies
1. 5.3.1 vCard
1. 5.3.1.1 Conversion to vCard
2. 5.3.1.2 Examples
2. 5.3.2 vEvent
1. 5.3.2.1 Conversion to iCalendar
2. 5.3.2.2 Examples
3. 5.3.3 Licensing works
1. 5.3.3.1 Examples
4. 5.4 Converting HTML to other formats
1. 5.4.1 JSON
6. 6 User interaction
1. 6.1 The hidden attribute
2. 6.2 Page visibility
1. 6.2.1 The VisibilityStateEntry interface
3. 6.3 Inert subtrees
1. 6.3.1 Modal dialogs and inert subtrees
2. 6.3.2 The inert attribute
4. 6.4 Tracking user activation
1. 6.4.1 Data model
2. 6.4.2 Processing model
3. 6.4.3 APIs gated by user activation
4. 6.4.4 The UserActivation interface
5. 6.4.5 User agent automation
5. 6.5 Activation behavior of elements
1. 6.5.1 The ToggleEvent interface
2. 6.5.2 The CommandEvent interface
6. 6.6 Focus
1. 6.6.1 Introduction
2. 6.6.2 Data model
3. 6.6.3 The tabindex attribute
4. 6.6.4 Processing model
5. 6.6.5 Sequential focus navigation
6. 6.6.6 Focus management APIs
7. 6.6.7 The autofocus attribute
7. 6.7 Assigning keyboard shortcuts
1. 6.7.1 Introduction
2. 6.7.2 The accesskey attribute
3. 6.7.3 Processing model
8. 6.8 Editing
1. 6.8.1 Making document regions editable: The
contenteditable content attribute
2. 6.8.2 Making entire documents editable: the designMode
getter and setter
3. 6.8.3 Best practices for in-page editors
4. 6.8.4 Editing APIs
5. 6.8.5 Spelling and grammar checking
6. 6.8.6 Writing suggestions
7. 6.8.7 Autocapitalization
8. 6.8.8 Autocorrection
9. 6.8.9 Input modalities: the inputmode attribute
10. 6.8.10 Input modalities: the enterkeyhint attribute
9. 6.9 Find-in-page
1. 6.9.1 Introduction
2. 6.9.2 Interaction with details and hidden=until-found
3. 6.9.3 Interaction with selection
10. 6.10 Close requests and close watchers
1. 6.10.1 Close requests
2. 6.10.2 Close watcher infrastructure
3. 6.10.3 The CloseWatcher interface
11. 6.11 Drag and drop
1. 6.11.1 Introduction
2. 6.11.2 The drag data store
3. 6.11.3 The DataTransfer interface
1. 6.11.3.1 The DataTransferItemList interface
2. 6.11.3.2 The DataTransferItem interface
4. 6.11.4 The DragEvent interface
5. 6.11.5 Processing model
6. 6.11.6 Events summary
7. 6.11.7 The draggable attribute
8. 6.11.8 Security risks in the drag-and-drop model
12. 6.12 The popover attribute
1. 6.12.1 The popover target attributes
2. 6.12.2 Popover light dismiss
7. 7 Loading web pages
1. 7.1 Supporting concepts
1. 7.1.1 Origins
1. 7.1.1.1 Sites
2. 7.1.1.2 Relaxing the same-origin restriction
3. 7.1.1.3 The Origin interface
2. 7.1.2 Origin-keyed agent clusters
3. 7.1.3 Cross-origin opener policies
1. 7.1.3.1 The headers
2. 7.1.3.2 Browsing context group switches due to
opener policy
3. 7.1.3.3 Reporting
4. 7.1.4 Cross-origin embedder policies
1. 7.1.4.1 The headers
2. 7.1.4.2 Embedder policy checks
5. 7.1.5 Sandboxing
6. 7.1.6 iframe element referrer policy
7. 7.1.7 Policy containers
2. 7.2 APIs related to navigation and session history
1. 7.2.1 Security infrastructure for Window, WindowProxy,
and Location objects
1. 7.2.1.1 Integration with IDL
2. 7.2.1.2 Shared internal slot:
[[CrossOriginPropertyDescriptorMap]]
3. 7.2.1.3 Shared abstract operations
1. 7.2.1.3.1 CrossOriginProperties ( O )
2. 7.2.1.3.2 CrossOriginPropertyFallback ( P )
3. 7.2.1.3.3 IsPlatformObjectSameOrigin ( O )
4. 7.2.1.3.4 CrossOriginGetOwnPropertyHelper ( O,
P )
5. 7.2.1.3.5 CrossOriginGet ( O, P, Receiver )
6. 7.2.1.3.6 CrossOriginSet ( O, P, V, Receiver )
7. 7.2.1.3.7 CrossOriginOwnPropertyKeys ( O )
2. 7.2.2 The Window object
1. 7.2.2.1 Opening and closing windows
2. 7.2.2.2 Indexed access on the Window object
3. 7.2.2.3 Named access on the Window object
4. 7.2.2.4 Accessing related windows
5. 7.2.2.5 Historical browser interface element APIs
6. 7.2.2.6 Script settings for Window objects
3. 7.2.3 The WindowProxy exotic object
1. 7.2.3.1 [[GetPrototypeOf]] ( )
2. 7.2.3.2 [[SetPrototypeOf]] ( V )
3. 7.2.3.3 [[IsExtensible]] ( )
4. 7.2.3.4 [[PreventExtensions]] ( )
5. 7.2.3.5 [[GetOwnProperty]] ( P )
6. 7.2.3.6 [[DefineOwnProperty]] ( P, Desc )
7. 7.2.3.7 [[Get]] ( P, Receiver )
8. 7.2.3.8 [[Set]] ( P, V, Receiver )
9. 7.2.3.9 [[Delete]] ( P )
10. 7.2.3.10 [[OwnPropertyKeys]] ( )
4. 7.2.4 The Location interface
1. 7.2.4.1 [[GetPrototypeOf]] ( )
2. 7.2.4.2 [[SetPrototypeOf]] ( V )
3. 7.2.4.3 [[IsExtensible]] ( )
4. 7.2.4.4 [[PreventExtensions]] ( )
5. 7.2.4.5 [[GetOwnProperty]] ( P )
6. 7.2.4.6 [[DefineOwnProperty]] ( P, Desc )
7. 7.2.4.7 [[Get]] ( P, Receiver )
8. 7.2.4.8 [[Set]] ( P, V, Receiver )
9. 7.2.4.9 [[Delete]] ( P )
10. 7.2.4.10 [[OwnPropertyKeys]] ( )
5. 7.2.5 The History interface
6. 7.2.6 The navigation API
1. 7.2.6.1 Introduction
2. 7.2.6.2 The Navigation interface
3. 7.2.6.3 Core infrastructure
4. 7.2.6.4 Initializing and updating the entry list
5. 7.2.6.5 The NavigationHistoryEntry interface
6. 7.2.6.6 The history entry list
7. 7.2.6.7 Initiating navigations
8. 7.2.6.8 Ongoing navigation tracking
9. 7.2.6.9 The NavigationActivation interface
10. 7.2.6.10 The navigate event
1. 7.2.6.10.1 The NavigateEvent interface
2. 7.2.6.10.2 The NavigationPrecommitController
interface
3. 7.2.6.10.3 The NavigationDestination interface
4. 7.2.6.10.4 Firing the event
5. 7.2.6.10.5 Scroll and focus behavior
7. 7.2.7 Event interfaces
1. 7.2.7.1 The NavigationCurrentEntryChangeEvent
interface
2. 7.2.7.2 The PopStateEvent interface
3. 7.2.7.3 The HashChangeEvent interface
4. 7.2.7.4 The PageSwapEvent interface
5. 7.2.7.5 The PageRevealEvent interface
6. 7.2.7.6 The PageTransitionEvent interface
7. 7.2.7.7 The BeforeUnloadEvent interface
8. 7.2.8 The NotRestoredReasons interface
3. 7.3 Infrastructure for sequences of documents
1. 7.3.1 Navigables
1. 7.3.1.1 Traversable navigables
2. 7.3.1.2 Top-level traversables
3. 7.3.1.3 Child navigables
4. 7.3.1.4 Jake diagrams
5. 7.3.1.5 Related navigable collections
6. 7.3.1.6 Navigable destruction
7. 7.3.1.7 Navigable target names
2. 7.3.2 Browsing contexts
1. 7.3.2.1 Creating browsing contexts
2. 7.3.2.2 Related browsing contexts
3. 7.3.2.3 Groupings of browsing contexts
3. 7.3.3 Fully active documents
4. 7.4 Navigation and session history
1. 7.4.1 Session history
1. 7.4.1.1 Session history entries
2. 7.4.1.2 Document state
3. 7.4.1.3 Centralized modifications of session history
4. 7.4.1.4 Low-level operations on session history
2. 7.4.2 Navigation
1. 7.4.2.1 Supporting concepts
2. 7.4.2.2 Beginning navigation
3. 7.4.2.3 Ending navigation
1. 7.4.2.3.1 The usual cross-document navigation
case
2. 7.4.2.3.2 The javascript: URL special case
3. 7.4.2.3.3 Fragment navigations
4. 7.4.2.3.4 Non-fetch schemes and external
software
4. 7.4.2.4 Preventing navigation
5. 7.4.2.5 Aborting navigation
3. 7.4.3 Reloading and traversing
4. 7.4.4 Non-fragment synchronous "navigations"
5. 7.4.5 Populating a session history entry
6. 7.4.6 Applying the history step
1. 7.4.6.1 Updating the traversable
2. 7.4.6.2 Updating the document
3. 7.4.6.3 Revealing the document
4. 7.4.6.4 Scrolling to a fragment
5. 7.4.6.5 Persisted history entry state
5. 7.5 Document lifecycle
1. 7.5.1 Shared document creation infrastructure
2. 7.5.2 Loading HTML documents
3. 7.5.3 Loading XML documents
4. 7.5.4 Loading text documents
5. 7.5.5 Loading multipart/x-mixed-replace documents
6. 7.5.6 Loading media documents
7. 7.5.7 Loading a document for inline content that doesn't
have a DOM
8. 7.5.8 Finishing the loading process
9. 7.5.9 Unloading documents
10. 7.5.10 Destroying documents
11. 7.5.11 Aborting a document load
6. 7.6 Speculative loading
1. 7.6.1 Speculation rules
1. 7.6.1.1 Data model
2. 7.6.1.2 Parsing
3. 7.6.1.3 Processing model
2. 7.6.2 Navigational prefetching
3. 7.6.3 The `Speculation-Rules` header
4. 7.6.4 The `Sec-Speculation-Tags` header
5. 7.6.5 Security considerations
1. 7.6.5.1 Cross-site requests
2. 7.6.5.2 Injected content
3. 7.6.5.3 IP anonymization
6. 7.6.6 Privacy considerations
1. 7.6.6.1 Heuristics and optionality
2. 7.6.6.2 State partitioning
3. 7.6.6.3 Identity joining
7. 7.7 The `X-Frame-Options` header
8. 7.8 The `Refresh` header
9. 7.9 Browser user interface considerations
8. 8 Web application APIs
1. 8.1 Scripting
1. 8.1.1 Introduction
2. 8.1.2 Agents and agent clusters
1. 8.1.2.1 Integration with the JavaScript agent
formalism
2. 8.1.2.2 Integration with the JavaScript agent
cluster formalism
3. 8.1.3 Realms and their counterparts
1. 8.1.3.1 Environments
2. 8.1.3.2 Environment settings objects
3. 8.1.3.3 Realms, settings objects, and global objects
1. 8.1.3.3.1 Entry
2. 8.1.3.3.2 Incumbent
3. 8.1.3.3.3 Current
4. 8.1.3.3.4 Relevant
4. 8.1.3.4 Enabling and disabling scripting
5. 8.1.3.5 Secure contexts
4. 8.1.4 Script processing model
1. 8.1.4.1 Scripts
2. 8.1.4.2 Fetching scripts
3. 8.1.4.3 Creating scripts
4. 8.1.4.4 Calling scripts
5. 8.1.4.5 Killing scripts
6. 8.1.4.6 Runtime script errors
7. 8.1.4.7 Unhandled promise rejections
8. 8.1.4.8 Import map parse results
9. 8.1.4.9 Speculation rules parse results
5. 8.1.5 Module specifier resolution
1. 8.1.5.1 The resolution algorithm
2. 8.1.5.2 Import maps
3. 8.1.5.3 Import map processing model
6. 8.1.6 JavaScript specification host hooks
1. 8.1.6.1 HostEnsureCanAddPrivateElement(O)
2. 8.1.6.2 HostEnsureCanCompileStrings(realm,
parameterStrings, bodyString, codeString,
compilationType, parameterArgs, bodyArg)
3. 8.1.6.3 HostGetCodeForEval(argument)
4. 8.1.6.4 HostPromiseRejectionTracker(promise,
operation)
5. 8.1.6.5 HostSystemUTCEpochNanoseconds(global)
6. 8.1.6.6 Job-related host hooks
1. 8.1.6.6.1 HostCallJobCallback(callback, V,
argumentsList)
2. 8.1.6.6.2
HostEnqueueFinalizationRegistryCleanupJob(final
izationRegistry)
3. 8.1.6.6.3 HostEnqueueGenericJob(job, realm)
4. 8.1.6.6.4 HostEnqueuePromiseJob(job, realm)
5. 8.1.6.6.5 HostEnqueueTimeoutJob(job, realm,
milliseconds)
6. 8.1.6.6.6 HostMakeJobCallback(callable)
7. 8.1.6.7 Module-related host hooks
1. 8.1.6.7.1
HostGetImportMetaProperties(moduleRecord)
2. 8.1.6.7.2 HostGetSupportedImportAttributes()
3. 8.1.6.7.3 HostLoadImportedModule(referrer,
moduleRequest, loadState, payload)
7. 8.1.7 Event loops
1. 8.1.7.1 Definitions
2. 8.1.7.2 Queuing tasks
3. 8.1.7.3 Processing model
4. 8.1.7.4 Generic task sources
5. 8.1.7.5 Dealing with the event loop from other
specifications
8. 8.1.8 Events
1. 8.1.8.1 Event handlers
2. 8.1.8.2 Event handlers on elements, Document
objects, and Window objects
1. 8.1.8.2.1 IDL definitions
3. 8.1.8.3 Event firing
2. 8.2 The WindowOrWorkerGlobalScope mixin
3. 8.3 Base64 utility methods
4. 8.4 Dynamic markup insertion
1. 8.4.1 Opening the input stream
2. 8.4.2 Closing the input stream
3. 8.4.3 document.write()
4. 8.4.4 document.writeln()
5. 8.5 DOM parsing and serialization APIs
1. 8.5.1 The DOMParser interface
2. 8.5.2 Unsafe HTML parsing methods
3. 8.5.3 HTML serialization methods
4. 8.5.4 The innerHTML property
5. 8.5.5 The outerHTML property
6. 8.5.6 The insertAdjacentHTML() method
7. 8.5.7 The createContextualFragment() method
8. 8.5.8 The XMLSerializer interface
6. 8.6 Timers
7. 8.7 Microtask queuing
8. 8.8 User prompts
1. 8.8.1 Simple dialogs
2. 8.8.2 Printing
9. 8.9 System state and capabilities
1. 8.9.1 The Navigator object
1. 8.9.1.1 Client identification
2. 8.9.1.2 Language preferences
3. 8.9.1.3 Browser state
4. 8.9.1.4 Custom scheme handlers: the
registerProtocolHandler() method
1. 8.9.1.4.1 Security and privacy
2. 8.9.1.4.2 User agent automation
5. 8.9.1.5 Cookies
6. 8.9.1.6 PDF viewing support
10. 8.10 Images
1. 8.10.1 The ImageData interface
2. 8.10.2 The ImageBitmap interface
11. 8.11 Animation frames
9. 9 Communication
1. 9.1 The MessageEvent interface
2. 9.2 Server-sent events
1. 9.2.1 Introduction
2. 9.2.2 The EventSource interface
3. 9.2.3 Processing model
4. 9.2.4 The `Last-Event-ID` header
5. 9.2.5 Parsing an event stream
6. 9.2.6 Interpreting an event stream
7. 9.2.7 Authoring notes
8. 9.2.8 Connectionless push and other features
9. 9.2.9 Garbage collection
10. 9.2.10 Implementation advice
3. 9.3 Cross-document messaging
1. 9.3.1 Introduction
2. 9.3.2 Security
1. 9.3.2.1 Authors
2. 9.3.2.2 User agents
3. 9.3.3 Posting messages
4. 9.4 Channel messaging
1. 9.4.1 Introduction
1. 9.4.1.1 Examples
2. 9.4.1.2 Ports as the basis of an object-capability
model on the web
3. 9.4.1.3 Ports as the basis of abstracting out
service implementations
2. 9.4.2 Message channels
3. 9.4.3 The MessageEventTarget mixin
4. 9.4.4 Message ports
5. 9.4.5 Ports and garbage collection
5. 9.5 Broadcasting to other browsing contexts
10. 10 Web workers
1. 10.1 Introduction
1. 10.1.1 Scope
2. 10.1.2 Examples
1. 10.1.2.1 A background number-crunching worker
2. 10.1.2.2 Using a JavaScript module as a worker
3. 10.1.2.3 Shared workers introduction
4. 10.1.2.4 Shared state using a shared worker
5. 10.1.2.5 Delegation
6. 10.1.2.6 Providing libraries
3. 10.1.3 Tutorials
1. 10.1.3.1 Creating a dedicated worker
2. 10.1.3.2 Communicating with a dedicated worker
3. 10.1.3.3 Shared workers
2. 10.2 Infrastructure
1. 10.2.1 The global scope
1. 10.2.1.1 The WorkerGlobalScope common interface
2. 10.2.1.2 Dedicated workers and the
DedicatedWorkerGlobalScope interface
3. 10.2.1.3 Shared workers and the
SharedWorkerGlobalScope interface
2. 10.2.2 The event loop
3. 10.2.3 The worker's lifetime
4. 10.2.4 Processing model
5. 10.2.5 Runtime script errors
6. 10.2.6 Creating workers
1. 10.2.6.1 The AbstractWorker mixin
2. 10.2.6.2 Script settings for workers
3. 10.2.6.3 Dedicated workers and the Worker interface
4. 10.2.6.4 Shared workers and the SharedWorker
interface
7. 10.2.7 Concurrent hardware capabilities
3. 10.3 APIs available to workers
1. 10.3.1 Importing scripts and libraries
2. 10.3.2 The WorkerNavigator interface
3. 10.3.3 The WorkerLocation interface
11. 11 Worklets
1. 11.1 Introduction
1. 11.1.1 Motivations
2. 11.1.2 Code idempotence
3. 11.1.3 Speculative evaluation
2. 11.2 Examples
1. 11.2.1 Loading scripts
2. 11.2.2 Registering a class and invoking its methods
3. 11.3 Infrastructure
1. 11.3.1 The global scope
1. 11.3.1.1 Agents and event loops
2. 11.3.1.2 Creation and termination
3. 11.3.1.3 Script settings for worklets
2. 11.3.2 The Worklet class
3. 11.3.3 The worklet's lifetime
12. 12 Web storage
1. 12.1 Introduction
2. 12.2 The API
1. 12.2.1 The Storage interface
2. 12.2.2 The sessionStorage getter
3. 12.2.3 The localStorage getter
4. 12.2.4 The StorageEvent interface
3. 12.3 Privacy
1. 12.3.1 User tracking
2. 12.3.2 Sensitivity of data
4. 12.4 Security
1. 12.4.1 DNS spoofing attacks
2. 12.4.2 Cross-directory attacks
3. 12.4.3 Implementation risks
13. 13 The HTML syntax
1. 13.1 Writing HTML documents
1. 13.1.1 The DOCTYPE
2. 13.1.2 Elements
1. 13.1.2.1 Start tags
2. 13.1.2.2 End tags
3. 13.1.2.3 Attributes
4. 13.1.2.4 Optional tags
5. 13.1.2.5 Restrictions on content models
6. 13.1.2.6 Restrictions on the contents of raw text
and escapable raw text elements
3. 13.1.3 Text
1. 13.1.3.1 Newlines
4. 13.1.4 Character references
5. 13.1.5 CDATA sections
6. 13.1.6 Comments
2. 13.2 Parsing HTML documents
1. 13.2.1 Overview of the parsing model
2. 13.2.2 Parse errors
3. 13.2.3 The input byte stream
1. 13.2.3.1 Parsing with a known character encoding
2. 13.2.3.2 Determining the character encoding
3. 13.2.3.3 Character encodings
4. 13.2.3.4 Changing the encoding while parsing
5. 13.2.3.5 Preprocessing the input stream
4. 13.2.4 Parse state
1. 13.2.4.1 The insertion mode
2. 13.2.4.2 The stack of open elements
3. 13.2.4.3 The list of active formatting elements
4. 13.2.4.4 The element pointers
5. 13.2.4.5 Other parsing state flags
5. 13.2.5 Tokenization
1. 13.2.5.1 Data state
2. 13.2.5.2 RCDATA state
3. 13.2.5.3 RAWTEXT state
4. 13.2.5.4 Script data state
5. 13.2.5.5 PLAINTEXT state
6. 13.2.5.6 Tag open state
7. 13.2.5.7 End tag open state
8. 13.2.5.8 Tag name state
9. 13.2.5.9 RCDATA less-than sign state
10. 13.2.5.10 RCDATA end tag open state
11. 13.2.5.11 RCDATA end tag name state
12. 13.2.5.12 RAWTEXT less-than sign state
13. 13.2.5.13 RAWTEXT end tag open state
14. 13.2.5.14 RAWTEXT end tag name state
15. 13.2.5.15 Script data less-than sign state
16. 13.2.5.16 Script data end tag open state
17. 13.2.5.17 Script data end tag name state
18. 13.2.5.18 Script data escape start state
19. 13.2.5.19 Script data escape start dash state
20. 13.2.5.20 Script data escaped state
21. 13.2.5.21 Script data escaped dash state
22. 13.2.5.22 Script data escaped dash dash state
23. 13.2.5.23 Script data escaped less-than sign state
24. 13.2.5.24 Script data escaped end tag open state
25. 13.2.5.25 Script data escaped end tag name state
26. 13.2.5.26 Script data double escape start state
27. 13.2.5.27 Script data double escaped state
28. 13.2.5.28 Script data double escaped dash state
29. 13.2.5.29 Script data double escaped dash dash state
30. 13.2.5.30 Script data double escaped less-than sign
state
31. 13.2.5.31 Script data double escape end state
32. 13.2.5.32 Before attribute name state
33. 13.2.5.33 Attribute name state
34. 13.2.5.34 After attribute name state
35. 13.2.5.35 Before attribute value state
36. 13.2.5.36 Attribute value (double-quoted) state
37. 13.2.5.37 Attribute value (single-quoted) state
38. 13.2.5.38 Attribute value (unquoted) state
39. 13.2.5.39 After attribute value (quoted) state
40. 13.2.5.40 Self-closing start tag state
41. 13.2.5.41 Bogus comment state
42. 13.2.5.42 Markup declaration open state
43. 13.2.5.43 Comment start state
44. 13.2.5.44 Comment start dash state
45. 13.2.5.45 Comment state
46. 13.2.5.46 Comment less-than sign state
47. 13.2.5.47 Comment less-than sign bang state
48. 13.2.5.48 Comment less-than sign bang dash state
49. 13.2.5.49 Comment less-than sign bang dash dash
state
50. 13.2.5.50 Comment end dash state
51. 13.2.5.51 Comment end state
52. 13.2.5.52 Comment end bang state
53. 13.2.5.53 DOCTYPE state
54. 13.2.5.54 Before DOCTYPE name state
55. 13.2.5.55 DOCTYPE name state
56. 13.2.5.56 After DOCTYPE name state
57. 13.2.5.57 After DOCTYPE public keyword state
58. 13.2.5.58 Before DOCTYPE public identifier state
59. 13.2.5.59 DOCTYPE public identifier (double-quoted)
state
60. 13.2.5.60 DOCTYPE public identifier (single-quoted)
state
61. 13.2.5.61 After DOCTYPE public identifier state
62. 13.2.5.62 Between DOCTYPE public and system
identifiers state
63. 13.2.5.63 After DOCTYPE system keyword state
64. 13.2.5.64 Before DOCTYPE system identifier state
65. 13.2.5.65 DOCTYPE system identifier (double-quoted)
state
66. 13.2.5.66 DOCTYPE system identifier (single-quoted)
state
67. 13.2.5.67 After DOCTYPE system identifier state
68. 13.2.5.68 Bogus DOCTYPE state
69. 13.2.5.69 CDATA section state
70. 13.2.5.70 CDATA section bracket state
71. 13.2.5.71 CDATA section end state
72. 13.2.5.72 Character reference state
73. 13.2.5.73 Named character reference state
74. 13.2.5.74 Ambiguous ampersand state
75. 13.2.5.75 Numeric character reference state
76. 13.2.5.76 Hexadecimal character reference start
state
77. 13.2.5.77 Decimal character reference start state
78. 13.2.5.78 Hexadecimal character reference state
79. 13.2.5.79 Decimal character reference state
80. 13.2.5.80 Numeric character reference end state
6. 13.2.6 Tree construction
1. 13.2.6.1 Creating and inserting nodes
2. 13.2.6.2 Parsing elements that contain only text
3. 13.2.6.3 Closing elements that have implied end tags
4. 13.2.6.4 The rules for parsing tokens in HTML
content
1. 13.2.6.4.1 The "initial" insertion mode
2. 13.2.6.4.2 The "before html" insertion mode
3. 13.2.6.4.3 The "before head" insertion mode
4. 13.2.6.4.4 The "in head" insertion mode
5. 13.2.6.4.5 The "in head noscript" insertion
mode
6. 13.2.6.4.6 The "after head" insertion mode
7. 13.2.6.4.7 The "in body" insertion mode
8. 13.2.6.4.8 The "text" insertion mode
9. 13.2.6.4.9 The "in table" insertion mode
10. 13.2.6.4.10 The "in table text" insertion mode
11. 13.2.6.4.11 The "in caption" insertion mode
12. 13.2.6.4.12 The "in column group" insertion
mode
13. 13.2.6.4.13 The "in table body" insertion mode
14. 13.2.6.4.14 The "in row" insertion mode
15. 13.2.6.4.15 The "in cell" insertion mode
16. 13.2.6.4.16 The "in template" insertion mode
17. 13.2.6.4.17 The "after body" insertion mode
18. 13.2.6.4.18 The "in frameset" insertion mode
19. 13.2.6.4.19 The "after frameset" insertion mode
20. 13.2.6.4.20 The "after after body" insertion
mode
21. 13.2.6.4.21 The "after after frameset"
insertion mode
5. 13.2.6.5 The rules for parsing tokens in foreign
content
7. 13.2.7 The end
8. 13.2.8 Speculative HTML parsing
9. 13.2.9 Coercing an HTML DOM into an infoset
10. 13.2.10 An introduction to error handling and strange
cases in the parser
1. 13.2.10.1 Misnested tags: <b><i></b></i>
2. 13.2.10.2 Misnested tags: <b><p></b></p>
3. 13.2.10.3 Unexpected markup in tables
4. 13.2.10.4 Scripts that modify the page as it is
being parsed
5. 13.2.10.5 The execution of scripts that are moving
across multiple documents
6. 13.2.10.6 Unclosed formatting elements
3. 13.3 Serializing HTML fragments
4. 13.4 Parsing HTML fragments
5. 13.5 Named character references
14. 14 The XML syntax
1. 14.1 Writing documents in the XML syntax
2. 14.2 Parsing XML documents
3. 14.3 Serializing XML fragments
4. 14.4 Parsing XML fragments
15. 15 Rendering
1. 15.1 Introduction
2. 15.2 The CSS user agent style sheet and presentational hints
3. 15.3 Non-replaced elements
1. 15.3.1 Hidden elements
2. 15.3.2 The page
3. 15.3.3 Flow content
4. 15.3.4 Phrasing content
5. 15.3.5 Bidirectional text
6. 15.3.6 Sections and headings
7. 15.3.7 Lists
8. 15.3.8 Tables
9. 15.3.9 Margin collapsing quirks
10. 15.3.10 Form controls
11. 15.3.11 The hr element
12. 15.3.12 The fieldset and legend elements
4. 15.4 Replaced elements
1. 15.4.1 Embedded content
2. 15.4.2 Images
3. 15.4.3 Attributes for embedded content and images
4. 15.4.4 Image maps
5. 15.5 Widgets
1. 15.5.1 Native appearance
2. 15.5.2 Writing mode
3. 15.5.3 Button layout
4. 15.5.4 The button element
5. 15.5.5 The details and summary elements
6. 15.5.6 The input element as a text entry widget
7. 15.5.7 The input element as domain-specific widgets
8. 15.5.8 The input element as a range control
9. 15.5.9 The input element as a color well
10. 15.5.10 The input element as a checkbox and radio button
widgets
11. 15.5.11 The input element as a file upload control
12. 15.5.12 The input element as a button
13. 15.5.13 The marquee element
14. 15.5.14 The meter element
15. 15.5.15 The progress element
16. 15.5.16 The select element
17. 15.5.17 The textarea element
6. 15.6 Frames and framesets
7. 15.7 Interactive media
1. 15.7.1 Links, forms, and navigation
2. 15.7.2 The title attribute
3. 15.7.3 Editing hosts
4. 15.7.4 Text rendered in native user interfaces
8. 15.8 Print media
9. 15.9 Unstyled XML documents
16. 16 Obsolete features
1. 16.1 Obsolete but conforming features
1. 16.1.1 Warnings for obsolete but conforming features
2. 16.2 Non-conforming features
3. 16.3 Requirements for implementations
1. 16.3.1 The marquee element
2. 16.3.2 Frames
3. 16.3.3 Other elements, attributes and APIs
17. 17 IANA considerations
1. 17.1 text/html
2. 17.2 multipart/x-mixed-replace
3. 17.3 application/xhtml+xml
4. 17.4 text/ping
5. 17.5 application/microdata+json
6. 17.6 application/speculationrules+json
7. 17.7 text/event-stream
8. 17.8 web+ scheme prefix
18. Index
1. Elements
2. Element content categories
3. Attributes
4. Element interfaces
5. All interfaces
6. Events
7. HTTP headers
8. MIME types
19. References
20. Acknowledgments
21. Intellectual property rights
