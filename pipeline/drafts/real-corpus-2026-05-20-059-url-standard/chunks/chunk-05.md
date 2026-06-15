---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---
## Chunk Context
Lines 1849–2133 of `raw/web/corpus-2026-05-18/059-url-standard.md`. Covers the EOF handling logic for URL parsing, URL serialization and equivalence rules, origin determination (including opaque origins for blob/file schemes), rendering guidelines for security display, and details on the `application/x-www-form-urlencoded` format plus its API.

## Local Summary
The chunk describes how a parser handles end-of-file (`EOF`) in various URL states, defines the steps to serialize a URL into an ASCII string (including handling of credentials, ports, paths, query strings, and fragments), explains equivalence checks via serialized comparison, outlines origin extraction for common schemes (blob, ftp, http, https, ws, wss, file), provides rendering advice to avoid spoofing (host simplification, hiding credentials, scheme replacement), and specifies parsing/serialization rules for `application/x-www-form-urlencoded` including percent-encoding hooks.

## Key Claims
- When `c` is the EOF code point, the parser must handle state transitions: special-query percent-encoding, fragment clearing, and validation of non-code-point characters or malformed percent-sequences (`%` followed by non-hex).
- URL serialization returns an ASCII string; it concatenates scheme, host (with credentials if present), optional port, path, query, and optionally fragment.
- Equivalence is determined by comparing serialized forms with fragments excluded by default.
- Origin extraction depends on scheme: tuple origins for `http`, `https`, `ws`, `wss`, `ftp`; blob origins come from the environment or are opaque; `file` origins are left as an exercise/often opaque.
- Rendering rules prioritize showing the registrable domain, hiding credentials, and handling IDNs/confusable characters to mitigate homograph attacks.
- `application/x-www-form-urlencoded` is a legacy, encoding-twisted format widely used for HTML forms; parsing splits on `&`, replaces `+` with space, percent-decodes UTF-8 without BOM, and serializing performs the inverse with specific percent-encode sets.

## Entities And Concepts
- EOF code point (end of input handling)
- URL state machine (special-query, fragment states)
- Percent-encoding sets (query, fragment, userinfo, application/x-www-form-urlencoded)
- URL serialization algorithm
- URL equivalence (serialized comparison)
- Origin (tuple vs opaque; scheme-specific rules)
- Rendering for security decisions
- Homograph spoofing / IDN confusion
- `application/x-www-form-urlencoded` (parsing, serializing, hooks)

## Procedures And API Details
**EOF Handling Steps**
When input ends:
1. If in special-query state, percent-encode remaining buffer with the query set and append to `url.query`.
2. If `c` is `#`, clear fragment and move to fragment state.
3. Otherwise (non-EOF): validate code point; if `%` appears without two hex digits following, raise validation error; else append to buffer.

**Fragment State Steps**
For each input character until EOF:
1. Validate code point or reject `%` without hex digits.
2. Percent-encode with fragment set and append to `url.fragment`.

**Username/Password Setting**
- Username: UTF-8 percent-encode using the userinfo set.
- Password: UTF-8 percent-encode using the userinfo set.

**URL Serialization Steps**
Input: `url`, optional `excludeFragment` (default false).
1. Output = scheme + `:`.
2. If host non-null:
   - Append `//`.
   - If credentials present: append username, then `:` + password if non-empty, then `@`.
3. Append serialized host.
4. If port non-null: append `:` + serialized port.
5. If path not opaque and has multiple segments starting with empty string, append `/.` to preserve path structure (prevent collapse).
6. Append serialized path.
7. If query non-null: append `?` + query.
8. If `excludeFragment` is false and fragment non-null: append `#` + fragment.
9. Return output.

**URL Path Serialization Steps**
1. If opaque path, return path as-is.
2. Else, iterate path segments, prefix each with `/`, join, return.

**Equivalence Check**
- Serialize both URLs with fragments excluded by default; compare resulting strings.

**Origin Determination**
- `blob`: environment origin from blob URL entry if available; else parse path as URL and derive or return opaque origin.
- `ftp`, `http`, `https`, `ws`, `wss`: tuple origin `(scheme, host, port, null)`.
- `file`: exercise to reader; often opaque.
- Others: opaque origin.

**Rendering Guidelines**
- Show only host (or registrable domain) where security distinction matters.
- Do not render username/password.
- May omit scheme if display surface is single-scheme; may replace with human-readable string or icon.
- Elide from lowest-level domain label when full host cannot fit.
- Render IDNs via Unicode normalization; detect confusable characters; treat bidirectional text as left-to-right for rendering.

**`application/x-www-form-urlencoded` Parsing Steps**
1. Split input on `&`.
2. For each segment:
   - If empty, skip.
   - If contains `=`, split into name/value (empty if at start/end).
   - Else treat whole as name, value empty.
3. Replace `+` with space in both name and value.
4. Percent-decode then UTF-8 decode without BOM for name and value.
5. Return list of `(nameString, valueString)` tuples.

**`application/x-www-form-urlencoded` Serialization Steps**
1. Determine output encoding (default UTF-8).
2. For each tuple:
   - Percent-encode name and value using the form-urlencoded set.
   - Join with `=`; separate tuples with `&`.
3. Return ASCII string.

**Hooks**
- String parser: UTF-8 encode input, then apply form-urlencoded parsing.

## Nuance Or Contradictions
- The `file` origin is explicitly left as an exercise, creating potential inconsistency across implementations unless defined elsewhere.
- Blob origins can be opaque if path parsing fails or scheme mismatches, leading to same-origin behavior differences.
- Rendering rules allow browser discretion (e.g., omitting `www`, replacing scheme with icon), which may cause serialized vs displayed URLs to differ significantly for security decisions.
- The form-urlencoded format is acknowledged as a legacy aberration with encoding twists; interoperability requirements override good design, leading to subtle pitfalls in nested encodings or charset handling.

## Candidate Wiki Hints
- **URL Serialization and Equivalence**: A canonical algorithm for converting URL objects to ASCII strings and checking equality by serialized comparison (fragments excluded).
- **Origin Extraction Logic**: Scheme-specific rules for determining tuple vs opaque origins, including blob environment handling.
- **Secure URL Rendering Guidelines**: Best practices for displaying URLs in address bars to prevent spoofing (hide credentials, show registrable domain, handle IDNs/confusables).
- **application/x-www-form-urlencoded Format**: Detailed parsing/serialization steps, percent-encode sets, and hook behavior; useful for form handling libraries.
