# TERMINOLOGY AND NAMING

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-NAME-001 |
| **Title** | Terminology and Naming |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define authoritative naming mechanics for OBDIA identifiers, filenames, repository paths, titles, requirement labels, cross-references, aliases, migrations and collision handling without redefining glossary semantics, lifecycle states or version semantics. |
| **Scope** | Constitutional artifacts, Knowledge documents 01–30, ADRs, requirements, diagrams, specifications, evidence and audit records, risk and change records, implementation artifacts, tests, releases, indexes, repository paths and generated representations. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `02_PROJECT_GLOSSARY.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `13_REPOSITORY_STATE_MODEL.md`; forward dependency `25_DOCUMENT_VERSIONING_POLICY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GLOSS-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-STATE-001` |
| **Cross-References** | Documents 15–30; accepted ADRs; controlled identifier registries; repository indexes; migration records; validation evidence |
| **Assumptions** | English is the authoritative language for normative identifiers, filenames and titles unless a controlled translation artifact explicitly states otherwise. |
| **Constraints** | This document governs naming mechanics only. Semantic definitions remain owned by the Project Glossary; lifecycle meaning remains owned by the Constitution and Repository State Model; version semantics remain owned by the Constitution and document 25. |
| **Security Considerations** | Ambiguous, colliding or sensitive names can cause authority confusion, cross-case disclosure, path manipulation, duplicate normative sources, broken traceability, homoglyph attacks, misleading status claims and accidental exposure of secrets or personal data. |
| **Validation Criteria** | Every governed artifact uses a unique immutable identifier, one canonical filename and path, approved notation, resolvable cross-references, controlled aliases, collision-free allocation and retained migration evidence. |
| **Implementation Relationship** | Repository tooling may lint and index names, but automation shall not allocate authority, redefine terminology, approve migrations or silently rename governed artifacts. |

---

## 1. Purpose

This document defines the authoritative naming mechanics for the Officer-Bound Digital Investigation Agent project.

The original baseline established four rules:

- documents use `UPPER_SNAKE_CASE.md`;
- ADRs use `ADR-0001-short-title.md`;
- Mermaid diagram sources use `kebab-case.mmd`;
- governed identifiers use `OBDIA-<DOMAIN>-<NUMBER>`.

This consolidation preserves those rules and adds the controls required to make them reliable across a governed repository.

Terminology meaning is controlled by `02_PROJECT_GLOSSARY.md`. This document controls how terms, artifacts and records are named, identified, referenced, migrated and validated.

## 2. Fundamental Naming Rules

- **NAME-REQ-001:** Every governed artifact shall have one authoritative name and, where applicable, one immutable identifier.
- **NAME-REQ-002:** Naming shall preserve human accountability and institutional attribution.
- **NAME-REQ-003:** A name shall not create legal authority, approval, validation, implementation or publication status by implication.
- **NAME-REQ-004:** A lower-level naming rule shall not contradict the Canonical Project Definition or Master Documentation Constitution.
- **NAME-REQ-005:** Naming shall not silently expand the normative Knowledge Pack beyond documents 01–30.
- **NAME-REQ-006:** Names shall be deterministic, reviewable and suitable for version control.
- **NAME-REQ-007:** Names shall avoid ambiguity across case-sensitive and case-insensitive filesystems.
- **NAME-REQ-008:** Identifiers shall remain distinct from human-readable titles.
- **NAME-REQ-009:** Filenames shall remain distinct from document identifiers.
- **NAME-REQ-010:** Lifecycle state shall remain distinct from filename, identifier and version.
- **NAME-REQ-011:** Implementation status shall not be encoded as authoritative lifecycle state.
- **NAME-REQ-012:** A repository alias shall not become a second normative authority.
- **NAME-REQ-013:** Naming collisions shall fail closed and block incorporation until resolved.
- **NAME-REQ-014:** Automated naming suggestions remain advisory until accepted by an authorized human role.
- **NAME-REQ-015:** Names shall not contain secrets, credentials or unnecessary personal or case information.
- **NAME-REQ-016:** Unresolved naming ambiguity affecting authority, security, privacy, evidence or release shall be blocking.

## 3. Authority Boundaries

### 3.1 Semantic Meaning

`02_PROJECT_GLOSSARY.md` owns controlled meanings, synonyms, deprecated terms and prohibited conflations.

### 3.2 Naming Mechanics

This document owns:

- identifier syntax;
- filename syntax;
- path syntax;
- title and label mechanics;
- domain-token allocation;
- cross-reference notation;
- alias, rename and migration mechanics;
- collision and case handling.

### 3.3 ADR Naming

`10_ADR_POLICY.md` owns ADR applicability and decision status. This document preserves ADR filename and identifier mechanics.

### 3.4 Documentation Structure

`11_DOCUMENTATION_STANDARD.md` owns document structure and authoring requirements. This document owns mechanical naming conventions.

### 3.5 Lifecycle State

`13_REPOSITORY_STATE_MODEL.md` owns lifecycle meaning. This document owns exact state-token spelling and notation.

### 3.6 Version Semantics

The Constitution and `25_DOCUMENT_VERSIONING_POLICY.md` own semantic-version meaning. This document owns only formatting and placement.

- **NAME-REQ-017:** This document shall not redefine Project Glossary meanings.
- **NAME-REQ-018:** This document shall not create synonyms that bypass glossary governance.
- **NAME-REQ-019:** This document shall not redefine ADR status meaning.
- **NAME-REQ-020:** This document shall not redefine lifecycle-state meaning.
- **NAME-REQ-021:** This document shall not redefine semantic-version compatibility.
- **NAME-REQ-022:** A naming change that alters meaning requires glossary or normative change control as applicable.
- **NAME-REQ-023:** A naming change that alters architecture requires an ADR where constitutionally required.
- **NAME-REQ-024:** A title clarification shall not silently alter requirement meaning.
- **NAME-REQ-025:** A repository path change shall not silently alter artifact authority.
- **NAME-REQ-026:** Naming ownership shall remain with identified human governance roles.
- **NAME-REQ-027:** An AI system shall not serve as final naming authority.
- **NAME-REQ-028:** Conflicts among naming authorities shall be resolved according to the constitutional hierarchy.
- **NAME-REQ-029:** Naming non-applicability requires documented rationale.
- **NAME-REQ-030:** Informative display labels shall not replace authoritative identifiers.
- **NAME-REQ-031:** Generated navigation shall identify its authoritative naming source.
- **NAME-REQ-032:** Authority-boundary conflicts shall not be resolved through undocumented renaming.

## 4. Character and Encoding Rules

Authoritative identifiers and repository filenames shall use a conservative character set.

Permitted identifier characters are normally:

- uppercase ASCII letters `A–Z`;
- digits `0–9`;
- hyphen `-`.

Permitted canonical filename characters are normally:

- uppercase or lowercase ASCII letters according to the applicable convention;
- digits;
- underscore `_`;
- hyphen `-`;
- one approved extension separator `.`.

- **NAME-REQ-033:** Governed identifiers shall use ASCII characters.
- **NAME-REQ-034:** Canonical repository filenames shall use ASCII characters.
- **NAME-REQ-035:** Spaces are prohibited in canonical filenames.
- **NAME-REQ-036:** Parentheses are prohibited in canonical filenames.
- **NAME-REQ-037:** Control characters are prohibited.
- **NAME-REQ-038:** Path separators shall not appear inside an identifier or filename component.
- **NAME-REQ-039:** Leading or trailing whitespace is prohibited.
- **NAME-REQ-040:** Leading or trailing periods are prohibited in filename components.
- **NAME-REQ-041:** Unicode homoglyphs and visually confusable characters are prohibited in authoritative identifiers.
- **NAME-REQ-042:** Unicode may appear in human-readable prose but shall not create a second authoritative identifier.
- **NAME-REQ-043:** Filesystem-reserved names shall not be used.
- **NAME-REQ-044:** Shell metacharacters shall not appear in canonical filenames.
- **NAME-REQ-045:** Percent-encoded or escaped representations shall not replace canonical names.
- **NAME-REQ-046:** Case shall be treated as significant even when the local filesystem is case-insensitive.
- **NAME-REQ-047:** Names differing only by case are collisions.
- **NAME-REQ-048:** Normalization or transliteration shall preserve a recorded mapping to the original human-readable term when material.

## 5. Governed Identifier Model

The base governed identifier pattern is:

`OBDIA-<DOMAIN>-<NUMBER>`

Where:

- `OBDIA` identifies the project namespace;
- `<DOMAIN>` is an approved uppercase domain token;
- `<NUMBER>` is a zero-padded allocation number.

Examples:

- `OBDIA-CANON-001`
- `OBDIA-GLOSS-001`
- `OBDIA-ADR-001`
- `OBDIA-NAME-001`

- **NAME-REQ-049:** Governed identifiers shall begin with `OBDIA-` unless a specialized approved policy defines another project-controlled family.
- **NAME-REQ-050:** Domain tokens shall use uppercase ASCII letters and digits only.
- **NAME-REQ-051:** Domain tokens shall be unique within the project namespace.
- **NAME-REQ-052:** Domain-token allocation shall be recorded in a controlled registry.
- **NAME-REQ-053:** Numeric allocations shall use at least three digits unless a superior existing convention requires otherwise.
- **NAME-REQ-054:** Existing assigned identifiers shall be preserved exactly.
- **NAME-REQ-055:** Identifiers shall be immutable after allocation.
- **NAME-REQ-056:** Retired, rejected, superseded or archived identifiers shall not be reused.
- **NAME-REQ-057:** A title change shall not change the identifier.
- **NAME-REQ-058:** A filename change shall not change the identifier.
- **NAME-REQ-059:** A repository move shall not change the identifier.
- **NAME-REQ-060:** Duplicate identifiers are blocking defects.
- **NAME-REQ-061:** Identifier allocation shall be attributable.
- **NAME-REQ-062:** An allocation gap shall not be filled by reusing an earlier identifier.
- **NAME-REQ-063:** Sequential numbering shall not be interpreted as priority, authority or lifecycle state.
- **NAME-REQ-064:** Identifier registries shall preserve allocation and retirement history.

## 6. Domain Tokens

A domain token identifies the governed namespace, not a repository directory.

Existing tokens include:

- `CANON`
- `CONST`
- `GLOSS`
- `RES`
- `PUB`
- `ARCH`
- `ID`
- `TM`
- `TRUST`
- `GOV`
- `ADR`
- `DOC`
- `IMP`
- `STATE`
- `NAME`

Future tokens shall be allocated only for documents and record families already authorized within Knowledge Pack 01–30 or approved lower-level artifacts.

- **NAME-REQ-065:** Domain tokens shall be concise and semantically stable.
- **NAME-REQ-066:** A token shall not be allocated solely to shorten a temporary branch or filename.
- **NAME-REQ-067:** Tokens differing only by abbreviation style shall not coexist when they refer to the same authority.
- **NAME-REQ-068:** Token collisions shall be resolved before identifier allocation.
- **NAME-REQ-069:** A new token shall identify its owning document or registry.
- **NAME-REQ-070:** A token shall not encode a person’s name.
- **NAME-REQ-071:** A token shall not encode a transient vendor or product unless an approved artifact family specifically requires it.
- **NAME-REQ-072:** A token shall not encode lifecycle status.
- **NAME-REQ-073:** A token shall not encode classification.
- **NAME-REQ-074:** Deprecated tokens shall remain reserved.
- **NAME-REQ-075:** Token migration shall preserve aliases only as non-authoritative history.
- **NAME-REQ-076:** Token allocation shall not create a new normative document.
- **NAME-REQ-077:** The controlled registry shall identify token, meaning, owner, allocation date and status.
- **NAME-REQ-078:** Unknown or unregistered domain tokens shall fail validation.

## 7. Requirement and Record Identifiers

Normative requirements use:

`<PREFIX>-REQ-<NUMBER>`

Examples:

- `NAME-REQ-001`
- `ADR-REQ-025`
- `STATE-REQ-150`

Other controlled records may use:

`<PREFIX>-<TYPE>-<NUMBER>`

where `<TYPE>` is defined by the owning policy, such as:

- `TERM`
- `THREAT`
- `RISK`
- `CTRL`
- `TEST`
- `EVID`
- `CHANGE`
- `DECISION`
- `EXCEPTION`
- `INCIDENT`

- **NAME-REQ-079:** Requirement prefixes shall map unambiguously to the owning normative document.
- **NAME-REQ-080:** Requirement numbers shall be zero-padded to at least three digits.
- **NAME-REQ-081:** Requirement identifiers shall be immutable.
- **NAME-REQ-082:** Removed requirements shall retain historical identifiers.
- **NAME-REQ-083:** Superseded requirements shall identify their successors or disposition.
- **NAME-REQ-084:** Requirement identifiers shall not be renumbered to close gaps.
- **NAME-REQ-085:** Requirement identifiers shall not encode severity, priority or lifecycle state.
- **NAME-REQ-086:** Record-type tokens shall be controlled by the owning policy.
- **NAME-REQ-087:** A record identifier shall remain unique within its owning registry.
- **NAME-REQ-088:** Evidence identifiers shall not contain case names, personal names or secret values.
- **NAME-REQ-089:** Risk, threat, control and test identifiers shall preserve bidirectional traceability.
- **NAME-REQ-090:** Temporary spreadsheet row numbers or issue numbers shall not replace governed record identifiers.
- **NAME-REQ-091:** GitHub issue and Pull Request numbers may supplement but shall not replace governed identifiers.
- **NAME-REQ-092:** A generated identifier shall not become authoritative until accepted into the controlled registry.
- **NAME-REQ-093:** Record families shall document allocation and retirement.
- **NAME-REQ-094:** Duplicate requirement or record identifiers are blocking.
- **NAME-REQ-095:** Cross-document requirement references shall use the full identifier.
- **NAME-REQ-096:** Human-readable labels may accompany identifiers but shall not replace them.

## 8. Canonical Document Filenames

### 8.1 Numbered Knowledge Documents

Numbered Knowledge documents use:

`NN_UPPER_SNAKE_CASE.md`

Example:

`14_TERMINOLOGY_AND_NAMING.md`

### 8.2 Non-Numbered Governance Documents

Non-numbered authoritative governance documents use:

`UPPER_SNAKE_CASE.md`

Example:

`MASTER_DOCUMENTATION_CONSTITUTION.md`

### 8.3 Source Variant Prohibition

Legacy suffixes such as `_ENTERPRISE` and temporary suffixes such as `_CONSOLIDATED` shall not remain in authoritative filenames.

- **NAME-REQ-097:** Numbered Knowledge filenames shall begin with the assigned two-digit document number.
- **NAME-REQ-098:** The number shall match the authoritative Knowledge Pack position.
- **NAME-REQ-099:** The filename title component shall use uppercase ASCII words separated by single underscores.
- **NAME-REQ-100:** Canonical Markdown documents shall use the lowercase `.md` extension.
- **NAME-REQ-101:** Legacy `_ENTERPRISE` suffixes are prohibited in canonical filenames.
- **NAME-REQ-102:** Temporary `_CONSOLIDATED` suffixes are prohibited in canonical filenames.
- **NAME-REQ-103:** Download collision suffixes such as `(1)` are prohibited in canonical filenames.
- **NAME-REQ-104:** Lifecycle states shall not appear in canonical filenames.
- **NAME-REQ-105:** Semantic versions shall not appear in canonical filenames unless a specialized approved release artifact requires them.
- **NAME-REQ-106:** Personal names shall not appear in canonical normative filenames.
- **NAME-REQ-107:** Branch names shall not be embedded in canonical filenames.
- **NAME-REQ-108:** One document identifier shall have only one canonical filename at a time.
- **NAME-REQ-109:** A filename shall remain stable unless clarity, conflict, migration or constitutional normalization requires change.
- **NAME-REQ-110:** A filename change shall update affected cross-references.
- **NAME-REQ-111:** A filename change shall preserve repository history.
- **NAME-REQ-112:** A filename shall not imply implementation or approval absent corresponding state evidence.

## 9. Repository Paths

Repository paths provide organization and do not create independent authority.

- **NAME-REQ-113:** Repository paths shall use lowercase ASCII directory components unless an approved repository standard defines otherwise.
- **NAME-REQ-114:** Directory components shall use lowercase kebab-case or an explicitly approved stable convention.
- **NAME-REQ-115:** Path traversal components such as `..` shall not appear in governed references.
- **NAME-REQ-116:** Absolute local workstation paths shall not appear as authoritative repository references.
- **NAME-REQ-117:** Internal references should use repository-relative paths.
- **NAME-REQ-118:** A path shall not include secrets, personal names or real case identifiers.
- **NAME-REQ-119:** Moving a document shall not change its identifier.
- **NAME-REQ-120:** Moving a document shall update navigation and cross-reference indexes.
- **NAME-REQ-121:** Duplicate authoritative copies in different directories are prohibited.
- **NAME-REQ-122:** Generated output shall be stored separately from authoritative source where feasible.
- **NAME-REQ-123:** Archive paths shall not be interpreted as current authority.
- **NAME-REQ-124:** Path conventions shall remain compatible with document 24 after consolidation.

## 10. Titles, Headings and Labels

Human-readable document titles use clear professional English.

- **NAME-REQ-125:** A document title shall describe the artifact’s subject and authority accurately.
- **NAME-REQ-126:** The top-level Markdown heading shall match the authoritative title in substance.
- **NAME-REQ-127:** Titles shall not contain lifecycle state unless the title describes state as a subject.
- **NAME-REQ-128:** Titles shall not contain semantic version numbers.
- **NAME-REQ-129:** Titles shall not imply certification, production readiness or official endorsement without evidence.
- **NAME-REQ-130:** Titles shall not use unexplained acronyms.
- **NAME-REQ-131:** Headings shall use stable descriptive wording.
- **NAME-REQ-132:** Heading changes that break links shall trigger link validation.
- **NAME-REQ-133:** Labels shall use Project Glossary terminology.
- **NAME-REQ-134:** Informative labels shall not be mistaken for normative identifiers.
- **NAME-REQ-135:** A title change shall be recorded in revision history when material.
- **NAME-REQ-136:** Translated titles shall identify the authoritative English title and identifier.
- **NAME-REQ-137:** Marketing phrases and unsupported superlatives are prohibited in normative titles.
- **NAME-REQ-138:** A title shall not conceal a narrower scope or material limitation.

## 11. ADR Naming

ADRs use:

- identifier: `ADR-0001`;
- filename: `ADR-0001-short-title.md`.

- **NAME-REQ-139:** ADR numbers shall use four digits with leading zeros.
- **NAME-REQ-140:** ADR numbers shall be allocated monotonically.
- **NAME-REQ-141:** ADR numbers shall not be reused.
- **NAME-REQ-142:** ADR short titles shall use lowercase kebab-case.
- **NAME-REQ-143:** ADR filenames shall not contain status, version, personal names or branch names.
- **NAME-REQ-144:** ADR filename and internal identifier shall match.
- **NAME-REQ-145:** ADR title clarification shall preserve the identifier.
- **NAME-REQ-146:** ADR renaming shall preserve history and update references.
- **NAME-REQ-147:** Superseded and rejected ADR identifiers shall remain reserved.
- **NAME-REQ-148:** ADR naming mechanics shall remain subordinate to `10_ADR_POLICY.md`.

## 12. Diagram Naming

Mermaid source files use:

`kebab-case.mmd`

Example:

`officer-agent-trust-boundaries.mmd`

Rendered diagrams may use a corresponding basename with an approved rendered extension.

- **NAME-REQ-149:** Mermaid source filenames shall use lowercase kebab-case.
- **NAME-REQ-150:** Mermaid source files shall use the lowercase `.mmd` extension.
- **NAME-REQ-151:** A rendered diagram should preserve the source basename.
- **NAME-REQ-152:** A rendered diagram shall remain traceable to its source.
- **NAME-REQ-153:** Diagram filenames shall describe subject rather than implementation status.
- **NAME-REQ-154:** Diagram filenames shall not contain secrets, infrastructure addresses or case identifiers.
- **NAME-REQ-155:** Diagram identifiers, where assigned, shall follow the controlled project identifier model.
- **NAME-REQ-156:** Duplicate diagram basenames within the same governed scope are prohibited.
- **NAME-REQ-157:** A diagram rename shall update embeds and references.
- **NAME-REQ-158:** Detailed diagram naming remains subordinate to document 23 after consolidation.

## 13. State, Version and Classification Notation

Authoritative lifecycle labels are:

- `Draft`
- `Under Review`
- `Approved`
- `Implemented`
- `Validated`
- `Published`
- `Superseded`
- `Archived`

Version metadata uses:

`MAJOR.MINOR.PATCH`

Prose may prefix a version with lowercase `v`, for example `v1.0.0`.

- **NAME-REQ-159:** Lifecycle labels shall use exact constitutional spelling.
- **NAME-REQ-160:** `Reviewed` shall not replace `Under Review`.
- **NAME-REQ-161:** `Tested` shall not replace `Validated` as a lifecycle state.
- **NAME-REQ-162:** State labels shall not be abbreviated in authoritative metadata.
- **NAME-REQ-163:** State labels shall not be encoded in canonical filenames.
- **NAME-REQ-164:** Metadata version values shall omit the `v` prefix.
- **NAME-REQ-165:** Prose version references may use the `v` prefix consistently.
- **NAME-REQ-166:** Version strings shall use three numeric components.
- **NAME-REQ-167:** Pre-release and build notation require document 25 governance before normative use.
- **NAME-REQ-168:** Classification labels shall use exact values defined by the governing policy.
- **NAME-REQ-169:** Status, classification and implementation state shall not be conflated.
- **NAME-REQ-170:** A badge or generated label shall match authoritative metadata.
- **NAME-REQ-171:** Disagreement among displayed states shall resolve to the authoritative evidence record.
- **NAME-REQ-172:** Version-format compliance shall not be represented as version-governance approval.

## 14. Cross-Reference Syntax

A complete governed reference should include:

- immutable identifier;
- exact title or filename where useful;
- repository-relative path or link where navigation is required;
- version or commit where a fixed historical reference is required.

Examples:

- `OBDIA-NAME-001`
- `14_TERMINOLOGY_AND_NAMING.md`
- `docs/knowledge/14_TERMINOLOGY_AND_NAMING.md`
- `NAME-REQ-001`

- **NAME-REQ-173:** Normative references shall use immutable identifiers where available.
- **NAME-REQ-174:** File references shall use exact canonical filenames.
- **NAME-REQ-175:** Repository links shall use relative paths where feasible.
- **NAME-REQ-176:** A title alone shall not be used when ambiguity is possible.
- **NAME-REQ-177:** Historical references shall identify version, commit or immutable evidence where material.
- **NAME-REQ-178:** Requirement references shall use full identifiers.
- **NAME-REQ-179:** Cross-references shall preserve case exactly.
- **NAME-REQ-180:** Broken references affecting authority, security, privacy, evidence or release are blocking.
- **NAME-REQ-181:** Informative references shall not be presented as normative dependencies.
- **NAME-REQ-182:** External references shall not create a project identifier implicitly.
- **NAME-REQ-183:** Aliases shall resolve to a single canonical target.
- **NAME-REQ-184:** Cross-reference validation shall occur before approval and baseline freeze.

## 15. Alias, Rename and Migration Rules

A canonical rename or migration shall use a governed migration record containing:

- affected identifier;
- prior filename and path;
- new filename and path;
- reason;
- authority;
- effective commit;
- affected references;
- compatibility period;
- alias or redirect plan;
- validation evidence;
- rollback plan.

- **NAME-REQ-185:** A rename shall preserve the immutable identifier.
- **NAME-REQ-186:** A rename shall not create a second normative document.
- **NAME-REQ-187:** Legacy files may remain only as clearly non-authoritative redirects or archived evidence when necessary.
- **NAME-REQ-188:** An alias shall identify the canonical target.
- **NAME-REQ-189:** An alias shall not contain independent normative content.
- **NAME-REQ-190:** Temporary aliases shall have an owner and removal condition.
- **NAME-REQ-191:** Alias chains are prohibited.
- **NAME-REQ-192:** Circular aliases are prohibited.
- **NAME-REQ-193:** A migration shall update indexes, links, diagrams, ADRs and implementation references.
- **NAME-REQ-194:** A migration shall preserve Git history where feasible.
- **NAME-REQ-195:** Case-only renames shall use a version-control-safe migration procedure.
- **NAME-REQ-196:** Legacy `_ENTERPRISE` source variants shall not remain as duplicate normative files.
- **NAME-REQ-197:** Migration validation shall confirm one canonical path per identifier.
- **NAME-REQ-198:** Failed migration shall use controlled rollback rather than silent duplication.

## 16. Collision Handling

A collision exists when two names or identifiers are identical, case-equivalent, visually confusable or semantically indistinguishable within the same governed namespace.

- **NAME-REQ-199:** Identifier collisions shall block incorporation.
- **NAME-REQ-200:** Canonical filename collisions shall block incorporation.
- **NAME-REQ-201:** Case-only filename collisions shall block incorporation.
- **NAME-REQ-202:** Confusable Unicode collisions shall block incorporation.
- **NAME-REQ-203:** Domain-token collisions shall be referred to the Documentation Authority.
- **NAME-REQ-204:** Semantic collisions shall be referred to the Project Glossary owner.
- **NAME-REQ-205:** Collision resolution shall preserve previously assigned identifiers.
- **NAME-REQ-206:** Collision resolution shall not renumber accepted historical artifacts.
- **NAME-REQ-207:** A collision record shall identify affected artifacts, authority and disposition.
- **NAME-REQ-208:** Automated collision detection shall not replace human semantic review.

## 17. Deprecation and Retirement

Naming deprecation applies to aliases, tokens, labels or conventions. It does not erase historical artifacts.

- **NAME-REQ-209:** Deprecated names shall identify their canonical replacements.
- **NAME-REQ-210:** Deprecated identifiers shall remain permanently reserved.
- **NAME-REQ-211:** Deprecated tokens shall not be allocated to new artifacts.
- **NAME-REQ-212:** Deprecation shall identify effective date, owner and migration guidance.
- **NAME-REQ-213:** Retirement shall preserve historical references.
- **NAME-REQ-214:** A deprecated name shall not be removed while governed references still depend on it without a migration plan.
- **NAME-REQ-215:** Deprecation shall not be used to conceal a superseded authority relationship.
- **NAME-REQ-216:** Replacement names shall pass collision and cross-reference validation.

## 18. Security, Privacy and Evidence Naming

- **NAME-REQ-217:** Names shall not contain passwords, tokens, API keys or cryptographic secrets.
- **NAME-REQ-218:** Names shall not contain real officer names unless an authorized internal registry explicitly requires them.
- **NAME-REQ-219:** Public names shall not contain real case names, victim names, suspect names or personal identifiers.
- **NAME-REQ-220:** Case-scoped internal identifiers shall use opaque values where feasible.
- **NAME-REQ-221:** Evidence identifiers shall not reveal sensitive content by themselves.
- **NAME-REQ-222:** Filenames shall not disclose classified or restricted data unnecessarily.
- **NAME-REQ-223:** Environment names shall not imply trust based solely on network location.
- **NAME-REQ-224:** Names such as `safe`, `trusted`, `secure`, `validated` or `production` require corresponding governed evidence when used as status-bearing labels.
- **NAME-REQ-225:** Untrusted external filenames shall be normalized before repository use.
- **NAME-REQ-226:** External filenames shall not directly determine tool behavior.
- **NAME-REQ-227:** Path and filename input shall be validated against traversal and injection risks.
- **NAME-REQ-228:** Sanitization shall preserve provenance linking the normalized name to the source name.
- **NAME-REQ-229:** Redaction shall not break required chain-of-custody traceability.
- **NAME-REQ-230:** Naming logs shall avoid unnecessary personal data.
- **NAME-REQ-231:** Security-sensitive aliases shall not disclose restricted canonical locations.
- **NAME-REQ-232:** Naming exceptions affecting evidence shall receive evidence-policy review after document 18 is consolidated.

## 19. Branches, Tags and Release Labels

Repository references are not artifact identifiers.

Recommended branch components use lowercase kebab-case with a stable category prefix, for example:

`docs/consolidate-14-terminology-and-naming`

- **NAME-REQ-233:** Branch names shall not establish artifact lifecycle state.
- **NAME-REQ-234:** Branch names shall not contain secrets or real case identifiers.
- **NAME-REQ-235:** Branch names should identify change category and subject.
- **NAME-REQ-236:** Pull Request numbers shall not replace governed identifiers.
- **NAME-REQ-237:** Git tags shall not establish Approved, Validated or Published state by themselves.
- **NAME-REQ-238:** Release labels shall follow document 15 and document 25 after consolidation.
- **NAME-REQ-239:** A release label shall not overstate implementation or assurance.
- **NAME-REQ-240:** Temporary worktree and download names shall not be incorporated as canonical repository names.

## 20. Validation and Automation

Naming validation should check:

- identifier syntax and uniqueness;
- filename syntax;
- path syntax;
- case collisions;
- domain registry;
- requirement identifiers;
- canonical-path uniqueness;
- legacy suffixes;
- alias chains;
- broken references;
- state and version notation;
- prohibited sensitive content patterns.

- **NAME-REQ-241:** Naming validation shall operate on the exact reviewed commit.
- **NAME-REQ-242:** Validation results shall identify rule, artifact and outcome.
- **NAME-REQ-243:** Validation shall distinguish errors, warnings and advisory findings.
- **NAME-REQ-244:** Identifier and canonical-filename collisions shall be errors.
- **NAME-REQ-245:** Legacy suffixes in authoritative filenames shall be errors.
- **NAME-REQ-246:** Broken normative references shall be errors.
- **NAME-REQ-247:** Tool configuration shall be versioned and reviewable.
- **NAME-REQ-248:** Automated renaming shall not occur without an approved migration plan.
- **NAME-REQ-249:** Automation shall not allocate final identifiers without registry control.
- **NAME-REQ-250:** Validation tooling shall fail safely when the registry is unavailable.
- **NAME-REQ-251:** False positives and tool limitations shall be documented.
- **NAME-REQ-252:** Passing automated validation shall not establish approval.
- **NAME-REQ-253:** Human review shall assess semantic and authority collisions.
- **NAME-REQ-254:** Validation evidence shall be retained for baseline freeze.
- **NAME-REQ-255:** Machine-readable and human-readable names shall agree.
- **NAME-REQ-256:** Disagreement shall resolve to the authoritative registry and generate correction work.

## 21. Minimum Validation Checklist

Before approval or migration, confirm:

- [ ] Document ID is unique and immutable;
- [ ] domain token is registered;
- [ ] numeric allocation is valid;
- [ ] filename follows the applicable canonical pattern;
- [ ] no `_ENTERPRISE`, `_CONSOLIDATED` or download suffix remains;
- [ ] canonical path is unique;
- [ ] identifier, title and filename are distinct but consistent;
- [ ] requirement and record identifiers are unique;
- [ ] case and Unicode collision checks pass;
- [ ] lifecycle labels use exact constitutional spelling;
- [ ] version notation is syntactically valid;
- [ ] ADR and diagram naming follows approved formats;
- [ ] cross-references use exact identifiers and filenames;
- [ ] aliases resolve directly to one canonical target;
- [ ] rename and migration evidence is retained;
- [ ] no secret, personal name or real case identifier is exposed;
- [ ] automated checks and human semantic review are complete;
- [ ] forward dependency on document 25 is recorded;
- [ ] revision history is current;
- [ ] Project Founder approval exists before status becomes Approved.

## 22. Limitations

- This document does not define the semantic meaning of controlled terms.
- It does not define ADR applicability or decision status.
- It does not define lifecycle-state criteria.
- It does not define semantic-version compatibility or release-version mechanics.
- It does not define the complete repository directory structure.
- It does not define full diagram governance.
- It does not allocate every future domain or record-type token.
- Automated naming validation cannot detect every semantic or authority conflict.
- Internal naming review does not constitute independent certification.
- Document 25 remains a forward dependency for complete version-notation governance.

## 23. Change Control

Every material change shall include:

- change identifier;
- rationale;
- affected conventions and registries;
- dependency analysis;
- migration and compatibility impact;
- security and privacy impact;
- evidence and traceability impact;
- automation impact;
- validation plan;
- rollback plan;
- approval record;
- semantic version change.

- **NAME-REQ-257:** Editorial corrections use a patch version when meaning is unchanged.
- **NAME-REQ-258:** Backward-compatible substantive additions use a minor version.
- **NAME-REQ-259:** Incompatible naming or identifier changes use a major version.
- **NAME-REQ-260:** Material naming-architecture decisions require an ADR where applicable.
- **NAME-REQ-261:** Changes require Project Founder approval.
- **NAME-REQ-262:** Unapproved naming-methodology changes during an active phase are prohibited.
- **NAME-REQ-263:** Changes shall not retroactively rewrite immutable identifiers.
- **NAME-REQ-264:** Changes shall preserve migration, collision and retirement evidence.
- **NAME-REQ-265:** A convention change shall include repository-wide impact analysis.
- **NAME-REQ-266:** Document 25 reconciliation is required before this policy becomes Approved.

## 24. Consolidation Record

Version 1.0.0 consolidates the existing Terminology and Naming baseline without expanding project scope. It:

- retains immutable document identifier `OBDIA-NAME-001`;
- normalizes the authoritative filename to `14_TERMINOLOGY_AND_NAMING.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves `UPPER_SNAKE_CASE.md` for documents;
- refines numbered Knowledge documents to `NN_UPPER_SNAKE_CASE.md`;
- preserves `ADR-0001-short-title.md`;
- preserves `kebab-case.mmd`;
- preserves `OBDIA-<DOMAIN>-<NUMBER>`;
- adds immutable-identifier, domain-registry, requirement, record, path, title, state-notation, version-notation, cross-reference, alias, migration, collision, deprecation, security, privacy and validation controls;
- treats the former `_ENTERPRISE` source as a legacy source variant rather than a separate normative document;
- identifies document 25 as a forward dependency blocking approval;
- creates no new normative document and authorizes no implementation or release.

## 25. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Documentation Authority; approval reserved to Project Founder | Constitutional consolidation of the original naming baseline: preserved document, ADR, diagram and identifier formats; added 266 stable requirements for identifiers, filenames, paths, titles, references, migration, collision, deprecation, security, validation and change control. |
