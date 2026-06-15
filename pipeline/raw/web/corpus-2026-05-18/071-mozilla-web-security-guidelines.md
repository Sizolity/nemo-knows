---
title: Mozilla Web Security Guidelines
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://infosec.mozilla.org/guidelines/web_security
tags: [security, web-corpus]
confidence: medium
---

# Mozilla Web Security Guidelines

## Fetch Metadata

- Corpus item: 71
- Category: Security
- Source URL: https://infosec.mozilla.org/guidelines/web_security
- Final URL: https://infosec.mozilla.org/guidelines/web_security
- Retrieved: 2026-05-18
- Content-Type: text/html; charset=utf-8
- Fetch status: ok
- Test value: Practical security checklist source.
- Fetched page title: Web Security

## Retrieved Text

Mozilla
Menu
* home
* security fundamentals
* security guidelines
* Improve this document
* Discover Firefox

Top

The goal of this document is to help operational teams with creating
secure web applications. All Mozilla sites and deployments are expected
to follow the recommendations below. Use of these recommendations by
the public is strongly encouraged. The Security Assurance team
maintains this document as a reference guide.

Table of Contents

1. Cheat Sheet
2. Transport Layer Security (TLS/SSL)
1. HTTPS
2. HTTP Strict Transport Security
3. HTTP Redirections
4. Resource Loading
3. Content Security Policy
4. contribute.json
5. Cookies
6. Cross-origin Resource Sharing
7. CSRF Prevention
8. Referrer Policy
9. robots.txt
10. Subresource Integrity
11. X-Content-Type-Options
12. X-Frame-Options
13. Version History

Web Security Cheat Sheet

Guideline Security
Benefit Implementation
Difficulty Order^† Requirements Notes
HTTPS MAXIMUM MEDIUM Mandatory Sites should use HTTPS (or other secure
protocols) for all communications
Redirections from HTTP MAXIMUM LOW 3 Mandatory Websites must redirect
to HTTPS, API endpoints should disable HTTP entirely
Resource Loading MAXIMUM LOW 2 Mandatory for all websites Both passive
and active resources should be loaded through protocols using TLS, such
as HTTPS
Strict Transport Security HIGH LOW 4 Mandatory for all websites Minimum
allowed time period of six months
TLS Configuration MEDIUM MEDIUM 1 Mandatory Use the most secure Mozilla
TLS configuration for your user base, typically Intermediate
Content Security Policy HIGH HIGH 10 Mandatory for new websites
Recommended for existing websites Disabling inline script is the
greatest concern for CSP implementation
Cookies HIGH MEDIUM 7 Mandatory for all new websites
Recommended for existing websites All cookies must be set with the
Secure flag, and set as restrictively as possible
contribute.json LOW LOW 9 Mandatory for all new Mozilla websites
Recommended for existing Mozilla sites Mozilla sites should serve
contribute.json and keep contact information up-to-date
Cross-origin Resource Sharing HIGH LOW 11 Mandatory Origin sharing
headers and files should not be present, except for specific use cases
Cross-site Request Forgery Tokenization HIGH UNKNOWN 6 Varies Mandatory
for websites that allow destructive changes
Unnecessary for all other websites
Most application frameworks have built-in CSRF tokenization to ease
implementation
Referrer Policy LOW LOW 12 Recommended for all websites Improves
privacy for users, prevents the leaking of internal URLs via Referer
header
robots.txt LOW LOW 14 Optional Websites that implement robots.txt must
use it only for noted purposes
Subresource Integrity MEDIUM MEDIUM 15 Recommended^‡ ^‡ Only for
websites that load JavaScript or stylesheets from foreign origins
X-Content-Type-Options LOW LOW 8 Recommended for all websites Websites
should verify that they are setting the proper MIME types for all
resources
X-Frame-Options HIGH LOW 5 Mandatory for all websites Websites that
don't use DENY or SAMEORIGIN must employ clickjacking defenses
^† Suggested order that administrators implement the web security
guidelines. It is based on a combination of the security impact and the
ease of implementation from an operational and developmental
perspective.

Transport Layer Security (TLS/SSL)

Transport Layer Security provides assurances about the confidentiality,
authentication, and integrity of all communications both inside and
outside of Mozilla. To protect our users and networked systems, the
support and use of encrypted communications using TLS is mandatory for
all systems.

HTTPS

Websites or API endpoints that only communicate with modern browsers
and systems should use the Mozilla modern TLS configuration.

Websites intended for general public consumption should use the Mozilla
intermediate TLS configuration.

Websites that require backwards compatibility with extremely old
browsers and operating systems may use the Mozilla backwards compatible
TLS configuration. This is not recommended, and use of this
compatibility level should be done only after consulting the security
team for a risk assessment.

Compatibility

Configuration Oldest compatible clients
Modern Firefox 63, Android 10.0, Chrome 70, Edge 75, Java 11, OpenSSL
1.1.1, Opera 57, and Safari 12.1
Intermediate Firefox 27, Android 4.4.2, Chrome 31, Edge, IE 11 on
Windows 7, Java 8u31, OpenSSL 1.0.1, Opera 20, Safari 9
Backwards Compatible (Old) Firefox 1, Android 2.3, Chrome 1, Edge 12,
IE8 on Windows XP, Java 6, OpenSSL 0.9.8, Opera 5, and Safari 1

See Also

* Mozilla Server Side TLS Guidelines
* Mozilla Server Side TLS Configuration Generator - generates
software configurations for the three levels of compatibility

HTTP Strict Transport Security

HTTP Strict Transport Security (HSTS) is an HTTP header that notifies
user agents to only connect to a given site over HTTPS, even if the
scheme chosen was HTTP. Browsers that have had HSTS set for a given
site will transparently upgrade all requests to HTTPS. HSTS also tells
the browser to treat TLS and certificate-related errors more strictly
by disabling the ability for users to bypass the error page.

The header consists of one mandatory parameter (max-age) and two
optional parameters (includeSubDomains and preload), separated by
semicolons.

Directives

* max-age: how long user agents will redirect to HTTPS, in seconds
* includeSubDomains: whether user agents should upgrade requests on
subdomains
* preload: whether the site should be included in the HSTS preload
list

max-age must be set to a minimum of six months (15768000), but longer
periods such as two years (63072000) are recommended. Note that once
this value is set, the site must continue to support HTTPS until the
expiry time has been reached.

includeSubDomains notifies the browser that all subdomains of the
current origin should also be upgraded via HSTS. For example, setting
includeSubDomains on domain.example.com will also set it on
host1.domain.example.com and host2.domain.example.com. Extreme care is
needed when setting the includeSubDomains flag, as it could disable
sites on subdomains that don’t yet have HTTPS enabled.

preload allows the website to be included in the HSTS preload list,
upon submission. As a result, web browsers will do HTTPS upgrades to
the site without ever having to receive the initial HSTS header. This
prevents downgrade attacks upon first use and is recommended for all
high risk websites. Note that being included in the HSTS preload list
requires that includeSubDomains also be set.

Examples

# Only connect to this site via HTTPS for the two years (recommended)
Strict-Transport-Security: max-age=63072000

# Only connect to this site and subdomains via HTTPS for the next two years and
also include in the preload list
Strict-Transport-Security: max-age=63072000; includeSubDomains; preload

See Also

* MDN on HTTP Strict Transport Security
* RFC6797: HTTP Strict Transport Security (HSTS)

HTTP Redirections

Websites may continue to listen on port 80 (HTTP) so that users do not
get connection errors when typing a URL into their address bar, as
browsers currently connect via HTTP for their initial request. Sites
that listen on port 80 should only redirect to the same resource on
HTTPS. Once the redirection has occurred, HSTS should ensure that all
future attempts go to the site via HTTP are instead sent directly to
the secure site. APIs or websites not intended for public consumption
should disable the use of HTTP entirely.

Sites should avoid redirections from HTTP to HTTPS on a different host,
as this prevents HSTS from being set. Instead, for example, first
redirect from http://example.com/ to https://example.com/ and then in a
second redirect, redirect from https://example.com/ to
https://example.org/

Examples

# Redirect all incoming http requests to the same site and URI on https, using n
ginx
server {
listen 80;

return 301 https://$host$request_uri;
}

# Redirect for site.example.org from http to https, using Apache
<VirtualHost *:80>
ServerName site.example.org
Redirect permanent / https://site.example.org/
</VirtualHost>

Resource Loading

All resources — whether on the same origin or not — should be loaded
over secure channels. Secure (HTTPS) websites that attempt to load
active resources such as JavaScript insecurely will be blocked by
browsers. As a result, users will experience degraded UIs and “mixed
content” warnings. Attempts to load passive content (such as images)
insecurely, although less risky, will still lead to degraded UIs and
can allow active attackers to deface websites or phish users.

Despite the fact that modern browsers make it evident that websites are
loading resources insecurely, these errors still occur with significant
frequency. To prevent this from occurring, developers should verify
that all resources are loaded securely prior to deployment.

Examples

<!-- HTTPS is a fantastic way to load a JavaScript resource -->
<script src="https://code.jquery.com/jquery-1.12.0.min.js"></script>

<!-- Attempts to load over HTTP will be blocked and will generate mixed content
warnings -->
<script src="http://code.jquery.com/jquery-1.12.0.min.js"></script>

<!-- Although passive content won't be blocked, it will still generate mixed con
tent warnings -->
<img src="http://very.badssl.com/image.jpg">

See Also

* MDN on Mixed Content

Content Security Policy

Content Security Policy (CSP) is an HTTP header that allows site
operators fine-grained control over where resources on their site can
be loaded from. The use of this header is the best method to prevent
cross-site scripting (XSS) vulnerabilities. Due to the difficulty in
retrofitting CSP into existing websites, CSP is mandatory for all new
websites and is strongly recommended for all existing high-risk sites.

The primary benefit of CSP comes from disabling the use of unsafe
inline JavaScript. Inline JavaScript – either reflected or stored –
means that improperly escaped user-inputs can generate code that is
interpreted by the web browser as JavaScript. By using CSP to disable
inline JavaScript, you can effectively eliminate almost all XSS attacks
against your site.

Note that disabling inline JavaScript means that all JavaScript must be
loaded from <script> src tags . Event handlers such as onclick used
directly on a tag will fail to work, as will JavaScript inside <script>
tags but not loaded via src. Furthermore, inline stylesheets using
either <style> tags or the style attribute will also fail to load. As
such, care must be taken when designing sites so that CSP becomes
easier to implement.

Implementation Notes

* Aiming for default-src https: is a great first goal, as it disables
inline code and requires https.
* For existing websites with large codebases that would require too
much work to disable inline scripts, default-src https:
'unsafe-inline' is still helpful, as it keeps resources from being
accidentally loaded over http. However, it does not provide any XSS
protection.
* It is recommended to start with a reasonably locked down policy
such as default-src 'none'; img-src 'self'; script-src 'self';
style-src 'self' and then add in sources as revealed during
testing.
* In lieu of the preferred HTTP header, pages can instead include a
<meta http-equiv=“Content-Security-Policy” content=“…”> tag. If
they do, it should be the first <meta> tag that appears inside
<head>.
* Care needs to be taken with data: URIs, as these are unsafe inside
script-src and object-src (or inherited from default-src).
* Similarly, the use of script-src 'self' can be unsafe for sites
with JSONP endpoints. These sites should use a script-src that
includes the path to their JavaScript source folder(s).
* Unless sites need the ability to execute plugins such as Flash or
Silverlight, they should disable their execution with object-src
'none'.
* Sites should ideally use the report-uri directive, which POSTs JSON
reports about CSP violations that do occur. This allows CSP
violations to be caught and repaired quickly.
* Prior to implementation, it is recommended to use the
Content-Security-Policy-Report-Only HTTP header, to see if any
violations would have occurred with that policy.
* Set form-action directive to none or self or to specific sources
which are allowed to be used in forms. This directive does not fall
back to default-src, therefore, not explicitly setting the
directive would allow forms to be submitted to any source.

Examples

# Disable unsafe inline/eval, only allow loading of resources (images, fonts, sc
ripts, etc.) over https
# Note that this does not provide any XSS protection
Content-Security-Policy: default-src https:

<!-- Do the same thing, but with a <meta> tag -->
<meta http-equiv="Content-Security-Policy" content="default-src https:">

# Disable the use of unsafe inline/eval, allow everything else except plugin exe
cution
Content-Security-Policy: default-src *; object-src 'none'

# Disable unsafe inline/eval, only load resources from same origin except also a
llow images from imgur
# Also disables the execution of plugins
Content-Security-Policy: default-src 'self'; img-src 'self' https://i.imgur.com;
object-src 'none'

# Disable unsafe inline/eval and plugins, only load scripts and stylesheets from
same origin, fonts from google,
# and images from same origin and imgur. Sites should aim for policies like this
.
Content-Security-Policy: default-src 'none'; font-src https://fonts.gstatic.com;
img-src 'self' https://i.imgur.com; object-src 'none';
script-src 'self'; style-src 'self'

# Pre-existing site that uses too much inline code to fix
# but wants to ensure resources are loaded only over https and disable plugins
Content-Security-Policy: default-src https: 'unsafe-eval' 'unsafe-inline'; objec
t-src 'none'

# Don't implement the above policy yet; instead just report violations that woul
d have occurred
Content-Security-Policy-Report-Only: default-src https:; report-uri /csp-violati
on-report-endpoint/

# Disable the loading of any resources and disable framing, recommended for APIs
to use
Content-Security-Policy: default-src 'none'; frame-ancestors 'none'

See Also

* An Introduction to Content Security Policy
* Content Security Policy Level 2 Standard
* Google CSP Evaluator
* Using the frame-ancestors directive to prevent framing

contribute.json

contribute.json is a text file placed within the root directory of a
website that describes what it is, where its source exists, what
technologies it uses, and how to reach support and contribute.
contribute.json is a Mozilla standard used to describe all active
Mozilla websites and projects.

Its existence can greatly speed up the process of bug triage,
particularly for smaller websites with just a handful of maintainers.
It further assists security researchers to find testable websites and
instructs them on where to file their bugs against. As such,
contribute.json is mandatory for all Mozilla websites, and must be
maintained as contributors join and depart projects.

Require subkeys include name, description, bugs, participate
(particularly irc and irc-contacts), and urls.

Examples

{
"name": "Bedrock",
"description": "The app powering www.mozilla.org.",
"repository": {
"url": "https://github.com/mozilla/bedrock",
"license": "MPL2",
"tests": "https://travis-ci.org/mozilla/bedrock/"
},
"participate": {
"home": "https://wiki.mozilla.org/Webdev/GetInvolved/mozilla.org",
"docs": "https://bedrock.readthedocs.io/en/latest/",
"mailing-list": "https://www.mozilla.org/about/forums/#dev-mozilla-org",
"irc": "irc://irc.mozilla.org/#www",
"irc-contacts": [
"someperson1",
"someperson2",
"someperson3"
]
},
"bugs": {
"list": "https://bugzilla.mozilla.org/describecomponents.cgi?product=www.m
ozilla.org",
"report": "https://bugzilla.mozilla.org/enter_bug.cgi?product=www.mozilla.
org",
"mentored": "https://bugzilla.mozilla.org/buglist.cgi?f1=bug_mentor&o1=isn
otempty
&query_format=advanced&bug_status=NEW&product=www.mozilla.org
&list_id=10866041"
},
"urls": {
"prod": "https://www.mozilla.org",
"stage": "https://www.allizom.org",
"dev": "https://www-dev.allizom.org",
"demo1": "https://www-demo1.allizom.org"
},
"keywords": [
"python",
"less-css",
"django",
"html5",
"jquery"
]
}

See Also

* The contribute.json Standard

Cookies

All cookies should be created such that their access is as limited as
possible. This can help minimize damage from cross-site scripting (XSS)
vulnerabilities, as these cookies often contain session identifiers or
other sensitive information.

Directives

* Name: Cookie names may be either be prepended with either __Secure-
or __Host- to prevent cookies from being overwritten by insecure
sources
+ Use __Host- for all cookies needed only on a specific domain
(no subdomains) where Path is set to /
+ Use __Secure- for all other cookies sent from secure origins
(such as HTTPS)
* Secure: All cookies must be set with the Secure flag, indicating
that they should only be sent over HTTPS
* HttpOnly: Cookies that don’t require access from JavaScript should
be set with the HttpOnly flag
* Expiration: Cookies should expire as soon as is necessary: session
identifiers in particular should expire quickly
+ Expires: Sets an absolute expiration date for a given cookie
+ Max-Age: Sets a relative expiration date for a given cookie
(not supported by IE <8)
* Domain: Cookies should only be set with this if they need to be
accessible on other domains, and should be set to the most
restrictive domain possible
* Path: Cookies should be set to the most restrictive path possible,
but for most applications this will be set to the root directory
* SameSite: Forbid sending the cookie via cross-origin requests (such
as from <img> tags, etc.), as a strong anti-CSRF measure
+ SameSite=Strict: Only send the cookie when site is directly
navigated to
+ SameSite=Lax: Send the cookie when navigating to your site
from another site

Examples

# Session identifier cookie only accessible on this host that gets purged when t
he user closes their browser
Set-Cookie: MOZSESSIONID=980e5da39d4b472b9f504cac9; Path=/; Secure; HttpOnly

# Session identifier for all example.org sites that expires in 30 days using the
__Secure- prefix
# This cookie is not sent cross-origin, but is sent when navigating to any Mozil
la site from another site
Set-Cookie: __Secure-MOZSESSIONID=7307d70a86bd4ab5a00499762; Max-Age=2592000; Do
main=example.org; Path=/; Secure; HttpOnly; SameSite=Lax

# Sets a long-lived cookie for the current host, accessible by Javascript, when
the user accepts the ToS
# This cookie is sent when navigating to your sent from another site, such as by
clicking a link
Set-Cookie: __Host-ACCEPTEDTOS=true; Expires=Fri, 31 Dec 9999 23:59:59 GMT; Path
=/; Secure; SameSite=Lax

# Session identifier used for a secure site, such as bugzilla.example.org. It is
n't sent from cross-origin
# requests, nor is it sent when navigating to bugzilla.example.org from another
site. Used in conjunction with
# other anti-CSRF measures, this is a very strong way to defend your site agains
t CSRF attacks.
Set-Cookie: __Host-BMOSESSIONID=YnVnemlsbGE=; Max-Age=2592000; Path=/; Secure; H
ttpOnly; SameSite=Strict

See Also

* RFC 6265 (HTTP Cookies)
* HTTP Cookie Prefixes
* Same-site Cookies

Cross-origin Resource Sharing

Access-Control-Allow-Origin is an HTTP header that defines which
foreign origins are allowed to access the content of pages on your
domain via scripts using methods such as XMLHttpRequest.
crossdomain.xml and clientaccesspolicy.xml provide similar
functionality, but for Flash and Silverlight-based applications,
respectively.

These should not be present unless specifically needed. Use cases
include content delivery networks (CDNs) that provide hosting for
JavaScript/CSS libraries and public API endpoints. If present, they
should be locked down to as few origins and resources as is needed for
proper function. For example, if your server provides both a website
and an API intended for XMLHttpRequest access on a remote websites,
only the API resources should return the Access-Control-Allow-Origin
header. Failure to do so will allow foreign origins to read the
contents of any page on your origin.

Examples

# Allow any site to read the contents of this JavaScript library, so that subres
ource integrity works
Access-Control-Allow-Origin: *

# Allow https://random-dashboard.example.org to read the returned results of thi
s API
Access-Control-Allow-Origin: https://random-dashboard.example.org

<!-- Allow Flash from https://random-dashboard.example.org to read page contents
-->
<cross-domain-policy xsi:noNamespaceSchemaLocation="http://www.adobe.com/xml/sch
emas/PolicyFile.xsd">
<allow-access-from domain="random-dashboard.example.org"/>
<site-control permitted-cross-domain-policies="master-only"/>
<allow-http-request-headers-from domain="random-dashboard.example.org" headers
="*" secure="true"/>
</cross-domain-policy>

<!-- The same thing, but for Silverlight-->
<?xml version="1.0" encoding="utf-8"?>
<access-policy>
<cross-domain-access>
<policy>
<allow-from http-request-headers="*">
<domain uri="https://random-dashboard.example.org"/>
</allow-from>
<grant-to>
<resource path="/" include-subpaths="true"/>
</grant-to>
</policy>
</cross-domain-access>
</access-policy>

See Also

* CORS for Developers
* MDN on HTTP access control (CORS)
* Adobe on Setting crossdomain.xml
* Microsoft on Setting clientaccesspolicy.xml

CSRF Prevention

Cross-site request forgeries are a class of attacks where unauthorized
commands are transmitted to a website from a trusted user. Because they
inherit the users cookies (and hence session information), they appear
to be validly issued commands. A CSRF attack might look like this:
<!-- Attempt to delete a user's account -->
<img src="https://accounts.example.org/management/delete?confirm=true">

When a user visits a page with that HTML fragment, the browser will
attempt to make a GET request to that URL. If the user is logged in,
the browser will provide their session cookies and the account deletion
attempt will be successful.

While there are a variety of mitigation strategies such as
Origin/Referrer checking and challenge-response systems (such as
CAPTCHA), the most common and transparent method of CSRF mitigation is
through the use of anti-CSRF tokens. Anti-CSRF tokens prevent CSRF
attacks by requiring the existence of a secret, unique, and
unpredictable token on all destructive changes. These tokens can be set
for an entire user session, rotated on a regular basis, or be created
uniquely for each request. Although SameSite cookies are the best
defense against CSRF attacks, they are not yet fully supported in all
browsers and should be used in conjunction with other anti-CSRF
defenses.

Examples

<!-- A secret anti-CSRF token, included in the form to delete an account -->
<input type="hidden" name="csrftoken" value="1df93e1eafa42012f9a8aff062eeb1db038
0b">

# Server-side: set an anti-CSRF cookie that JavaScript must send as an X header,
which can't be done cross-origin
Set-Cookie: CSRFTOKEN=1df93e1eafa42012f9a8aff062eeb1db0380b; Path=/; Secure; Sam
eSite=Strict

// Client-side, have JavaScript add it as an X header to the XMLHttpRequest
var token = readCookie(CSRFTOKEN); // read the cookie
httpRequest.setRequestHeader('X-CSRF-Token', token); // add it as an X-CSRF-Toke
n header

See Also

* Wikipedia on CRSF Attacks and Prevention
* OWASP CSRF Prevention Cheat Sheet

Referrer Policy

When a user navigates to a site via a hyperlink or a website loads an
external resource, browsers inform the destination site of the origin
of the requests through the use of the HTTP Referer (sic) header.
Although this can be useful for a variety of purposes, it can also
place the privacy of users at risk. HTTP Referrer Policy allows sites
to have fine-grained control over how and when browsers transmit the
HTTP Referer header.

In legacy browsers, if a page at https://example.com/page.html contains
<img src="https://not.example.com/image.jpg">, then the browser will
send a request like this:
GET /image.jpg HTTP/1.1
Host: not.example.com
Referer: https://example.com/page.html

In addition to the privacy risks that this entails, the browser may
also transmit internal-use-only URLs that it may not have intended to
reveal. If you as the site operator want to limit the exposure of this
information, you can use HTTP Referrer Policy to either eliminate the
Referer header or reduce the amount of information that it contains.

Directives

* no-referrer: never send the Referer header
* same-origin: send referrer, but only on requests to the same origin
* strict-origin: send referrer to all origins, but only the URL
without the path (e.g. https://example.com/)
* strict-origin-when-cross-origin: send full referrer on same origin,
URL without the path on foreign origin

Notes

Although there are other options for referrer policies, they do not
protect user privacy and limit exposure in the same way as the options
above.

strict-origin-when-cross-origin is the default behavior for all current
browsers.

Referrer Policy has good support across modern browsers. In recent
versions of Firefox and Safari, “unsafe” directives
(no-referrer-when-downgrade, origin-when-cross-origin, and unsafe-url)
will behave like the default of strict-origin-when-cross-origin.

Examples

# On example.com, only send the Referer header when loading or linking to other
example.com resources
Referrer-Policy: same-origin

# Only send the shortened referrer to a foreign origin, full referrer to a local
host
Referrer-Policy: strict-origin-when-cross-origin

# Disable referrers for browsers that don't support strict-origin-when-cross-ori
gin
# Uses strict-origin-when-cross-origin for browsers that do
Referrer-Policy: no-referrer, strict-origin-when-cross-origin

<!-- Do the same, but with a meta tag -->
<meta http-equiv="Referrer-Policy" content="no-referrer, strict-origin-when-cros
s-origin">

<!-- referrerpolicy doesn't support multiple values, so only set no-referrer -->
<a href="https://example.org/" referrerpolicy="no-referrer">

See Also

* Referrer Policy standard
* MDN on Referrer Policy

robots.txt

robots.txt is a text file placed within the root directory of a site
that tells robots (such as indexers employed by search engines) how to
behave, by instructing them not to crawl certain paths on the website.
This is particularly useful for reducing load on your website through
disabling the crawling of automatically generated content. It can also
be helpful for preventing the pollution of search results, for
resources that don’t benefit from being searchable.

Sites may optionally use robots.txt, but should only use it for these
purposes. It should not be used as a way to prevent the disclosure of
private information or to hide portions of a website. Although this
does prevent these sites from appearing in search engines, it does not
prevent its discovery from attackers, as robots.txt is frequently used
for reconnaissance.

Examples

# Stop all search engines from crawling this site
User-agent: *
Disallow: /

# Using robots.txt to hide certain directories is a terrible idea
User-agent: *
Disallow: /secret/admin-interface

See Also

* About robots.txt

Subresource Integrity

Subresource integrity is a recent W3C standard that protects against
attackers modifying the contents of JavaScript libraries hosted on
content delivery networks (CDNs) in order to create vulnerabilities in
all websites that make use of that hosted library.

For example, JavaScript code on jquery.org that is loaded from
example.org has access to the entire contents of everything of
example.org. If this resource was successfully attacked, it could
modify download links, deface the site, steal credentials, cause
denial-of-service attacks, and more.

Subresource integrity locks an external JavaScript resource to its
known contents at a specific point in time. If the file is modified at
any point thereafter, supporting web browsers will refuse to load it.
As such, the use of subresource integrity is mandatory for all external
JavaScript resources loaded from sources not hosted on
Mozilla-controlled systems.

Note that CDNs must support the Cross Origin Resource Sharing (CORS)
standard by setting the Access-Control-Allow-Origin header. Most CDNs
already do this, but if the CDN you are loading does not support CORS,
please contact the Security Assurance team. We are happy to contact the
CDN on your behalf.

Directives

* integrity: a cryptographic hash of the file, prepended with the
hash function used to generate it
* crossorigin: should be anonymous to inform browsers to send
anonymous requests without cookies

Examples

<!-- Load jQuery 2.1.4 from their CDN -->
<script src="https://code.jquery.com/jquery-2.1.4.min.js"
integrity="sha384-R4/ztc4ZlRqWjqIuvf6RX5yb/v90qNGx6fS48N0tRxiGkqveZETq72KgDVJC
p2TC"
crossorigin="anonymous"></script>

<!-- Load AngularJS 1.4.8 from their CDN -->
<script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular.min.j
s"
integrity="sha384-r1y8TJcloKTvouxnYsi4PJAx+nHNr90ibsEn3zznzDzWBN9X3o3kbHLSgcIP
tzAp"
crossorigin="anonymous"></script>

# Generate the hash myself
$ curl -s https://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular.min.js |
\
openssl dgst -sha384 -binary | \
openssl base64 -A

r1y8TJcloKTvouxnYsi4PJAx+nHNr90ibsEn3zznzDzWBN9X3o3kbHLSgcIPtzAp

See Also

* SRI Hash Generator - generates <script> tags for you, and informs
you if the CDN lacks CORS support
* Subresource Integrity W3C Standard
* MDN on Subresource Integrity

X-Content-Type-Options

X-Content-Type-Options is a header supported by Internet Explorer,
Chrome and Firefox 50+ that tells it not to load scripts and
stylesheets unless the server indicates the correct MIME type. Without
this header, these browsers can incorrectly detect files as scripts and
stylesheets, leading to XSS attacks. As such, all sites must set the
X-Content-Type-Options header and the appropriate MIME types for files
that they serve.

Examples

# Prevent browsers from incorrectly detecting non-scripts as scripts
X-Content-Type-Options: nosniff

See Also

* Microsoft on Reducing MIME Type Security Risks

X-Frame-Options

X-Frame-Options is an HTTP header that allows sites control over how
your site may be framed within an iframe. Clickjacking is a practical
attack that allows malicious sites to trick users into clicking links
on your site even though they may appear to not be on your site at all.
As such, the use of the X-Frame-Options header is mandatory for all new
websites, and all existing websites are expected to add support for
X-Frame-Options as soon as possible.

Note that X-Frame-Options has been superseded by the Content Security
Policy’s frame-ancestors directive, which allows considerably more
granular control over the origins allowed to frame a site. As
frame-ancestors is not yet supported in IE11 and older, Edge, Safari
9.1 (desktop), and Safari 9.2 (iOS), it is recommended that sites
employ X-Frame-Options in addition to using CSP.

Sites that require the ability to be iframed must use either Content
Security Policy and/or employ JavaScript defenses to prevent
clickjacking from malicious origins.

Directives

* DENY: disallow allow attempts to iframe site (recommended)
* SAMEORIGIN: allow the site to iframe itself
* ALLOW-FROM uri: deprecated; instead use CSP’s frame-ancestors
directive

Examples

# Block site from being framed with X-Frame-Options and CSP
Content-Security-Policy: frame-ancestors 'none'
X-Frame-Options: DENY

# Only allow my site to frame itself
Content-Security-Policy: frame-ancestors 'self'
X-Frame-Options: SAMEORIGIN

# Allow only framer.example.org to frame site
# Note that this blocks framing from browsers that don't support CSP2+
Content-Security-Policy: frame-ancestors https://framer.example.org
X-Frame-Options: DENY

See Also

* MDN on X-Frame-Options
* CSP standard on ‘frame-ancestors’
* OWASP Clickjacking Defense Cheat Sheet

Version History

Date Editor Changes
March, 2024 arroway Remove HPKP and XXSSP obsolete guidelines
March, 2024 Frida Updates to CSP recommendations
February, 2024 LeoMcA Update Referrer Policy section
August, 2023 arroway Update HPKP, XXSSP, TLS and redirects sections
April, 2020 Mads J. Update TLS configurations client support
September, 2019 rmccue Fix CSP examples
September, 2018 Chris M. Update Referrer Policy support
July, 2018 April Link to CORS for Developers
April, 2018 April Added SameSite cookies and syntax highlighting
June, 2017 April Moved cookie prefixes to no longer be experimental
November, 2016 April Added Referrer Policy, tidied up XFO examples
October, 2016 April Updates to CSP recommendations
July, 2016 April Updates to CSP for APIs, and CSP’s deprecation of XFO,
and XXSSP
February, 2016 April Initial document creation

Mozilla

©2017 by individual contributors. Content available under a Creative
Commons license
* Contact Us

* Privacy
* Cookies

* Security:
+ Twitter (@mozillaSecurity)
* Mozilla:
+ Twitter (@mozilla)
+ Facebook (Mozilla)
* Firefox:
+ Twitter (@firefox)
+ Facebook (Firefox)
+ YouTube (firefoxchannel)
