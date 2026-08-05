# MASTER DOCUMENTATION CONSTITUTION

| Metadata Field | Value |
|---|---|
| **Document Title** | Master Documentation Constitution |
| **Document ID** | OBDIA-CONST-001 |
| **Version** | 1.0.0 |
| **Status** | Approved Constitutional Baseline |
| **Classification** | Supreme Documentation Governance |
| **Authority** | Project Founder |
| **Effective Date** | 2026-08-05 |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md` |
| **Supersession Rules** | This Constitution supersedes any lower-level documentation-governance rule that conflicts with it. It does not supersede, amend, reinterpret, or reduce the authority of `01_CANONICAL_PROJECT_DEFINITION.md`. A superseded version of this Constitution remains immutable in project history and is replaced only by an explicitly approved later version. |
| **Change-Control Authority** | Project Founder. The Documentation Authority administers change records and reviews but cannot approve a constitutional amendment without explicit Founder approval. |

---

## 1. Constitutional Purpose and Status

This Constitution is the supreme documentation-governance authority for the Officer-Bound Digital Investigation Agent project immediately below `01_CANONICAL_PROJECT_DEFINITION.md`.

Its purpose is to preserve a stable, auditable, internally consistent, and security-first body of project knowledge. It governs the creation, classification, approval, use, implementation, validation, publication, supersession, and archival of:

- normative Knowledge documents;
- policies and governance records;
- architecture models;
- Architecture Decision Records;
- repository architecture and governance documentation;
- technical specifications;
- diagrams;
- implementation;
- tests and validation records;
- evidence and audit records;
- releases;
- research records, assumptions, and backlog items.

This Constitution governs documentation methodology. It does not grant investigative authority, legal authority, operational authority, or institutional status to any person, component, or AI agent.

The Canonical Project Definition remains the supreme conceptual authority. The Officer-Bound Digital Investigation Agent remains an institutionally issued, cryptographically and operationally bound digital extension of one identified human officer, with no independent legal authority. Human and institutional accountability cannot be delegated to the agent.

### 1.1 Binding Language

The terms **shall**, **must**, **must not**, and **prohibited** express binding requirements. **Should** expresses a strong recommendation that requires documented justification when not followed. **May** expresses a permitted option.

### 1.2 Stability During Active Phases

Once an active documentation, architecture, implementation, validation, or release phase has an approved scope, no contributor may introduce an unapproved methodological change, alternative governance structure, additional Knowledge layer, expanded normative document set, or strategic reformulation. A material methodological change requires formal change control under this Constitution before it may affect active work.

An unresolved proposal remains informative only. It cannot redirect approved work, alter authority, or create an implicit baseline.

---

## 2. Normative Hierarchy

The project shall apply the following hierarchy exactly:

1. `01_CANONICAL_PROJECT_DEFINITION.md`
2. `MASTER_DOCUMENTATION_CONSTITUTION.md`
3. Approved normative Knowledge documents
4. Accepted Architecture Decision Records
5. Repository architecture and governance documentation
6. Technical specifications and diagrams
7. Implementation
8. Tests, validation records and evidence
9. Informative research notes and backlog items

### 2.1 Precedence

A higher-level authority prevails over every lower-level artifact. No lower-level artifact may contradict, weaken, bypass, reinterpret, or silently expand a higher-level authority.

A lower-level artifact may add detail only when the detail:

- remains within approved scope;
- is compatible with every applicable higher-level requirement;
- does not create new legal or institutional authority;
- preserves officer-bound delegation and human accountability;
- is traceable to an approved requirement or decision;
- is identified accurately as normative, proposed, implemented, tested, simulated, or informative.

### 2.2 Conflict Handling

When a conflict is identified:

1. the conflict shall be recorded explicitly;
2. the higher-level authority shall control;
3. affected lower-level decisions or implementation shall be suspended where continued use could create non-conformance, security risk, privacy risk, evidence-integrity risk, or misleading publication;
4. the responsible owner shall identify all affected artifacts and dependencies;
5. the resolution shall be approved at the authority level required for the affected material;
6. cross-references and versions shall be updated;
7. an ADR shall be issued when the conflict concerns a material architectural decision;
8. the audit record shall preserve the original conflict and its resolution.

No conflict may be resolved by silently editing history or treating implementation behavior as an implicit architectural decision.

### 2.3 Authority of Pack Membership

Membership in the Knowledge Pack does not override a document's classification or lifecycle state. A draft Knowledge document is not approved. An informative or research record does not become normative because it is stored beside normative documents.

---

## 3. Knowledge Pack Scope

The initial governed baseline is frozen at **30 Knowledge documents**:

- **01–04 — Foundation**
- **05–15 — Architecture and Governance**
- **16–25 — Engineering**
- **26–30 — Research, Risk and Compliance**

The baseline name is:

> **OBDIA Knowledge Pack Enterprise v1.0**

Documents **31–35** are auxiliary research materials. They are not part of the normative v1.0 baseline and shall not be treated as approved architecture, policy, control, implementation requirement, or compliance evidence unless incorporated through formal change control.

No contributor may expand the normative Knowledge Pack, reserve additional normative layers, renumber the baseline, or create new normative Knowledge documents outside 01–30 without an approved constitutional change. Research needs shall be recorded within the existing governed backlog and assumptions mechanisms unless the Constitution is formally amended.

Legacy manifests, roadmaps, build orders, indexes, or archived packs that propose documents above 30 are informative historical materials only and have no authority to expand the v1.0 baseline.

---

## 4. Document Classification

Every governed artifact shall declare one primary classification. An artifact may reference other classes but shall not use multiple classifications to avoid clear authority.

| Class | Authority | Permitted Content | Approval Requirements | Change Control | Relationship to Implementation |
|---|---|---|---|---|---|
| **Constitutional** | Highest governance authority assigned to the specific constitutional artifact, subject only to any expressly superior constitutional source | Project purpose, immutable boundaries, normative hierarchy, governance authority, amendment rules | Explicit Project Founder approval | Formal constitutional amendment; impact assessment; recorded effective date; version increment | Implementation shall conform. Implementation cannot amend constitutional content. |
| **Normative** | Binding within its approved scope at hierarchy level 3 | Mandatory principles, policies, architecture requirements, control requirements, governance rules, required models | Applicable review gates and explicit approval by the stated authority; Founder approval for the v1.0 normative baseline | Change record; impact analysis; affected reviews; semantic version increment | Implementation shall satisfy traceable requirements. Non-conforming code remains non-conforming until formally resolved. |
| **Decision** | Binding for the decision scope at hierarchy level 4 when accepted | Context, alternatives, decision, rationale, consequences, risks, validation, supersession | Acceptance by the designated decision authority after required reviews | Accepted records are immutable; change occurs through a new or superseding ADR | Implementation shall reference and conform to the accepted decision. |
| **Specification** | Binding only when approved and only within delegated detail at hierarchy level 6 | Interfaces, schemas, protocols, state transitions, control parameters, data models, diagrammed technical behavior | Owner approval plus architecture, security, and other applicable reviews | Controlled revision with compatibility and migration analysis | Implementation realizes the specification; the specification cannot contradict normative architecture. |
| **Operational** | Procedural authority for approved operations and repository processes | Runbooks, release procedures, review procedures, maintenance steps, incident and containment procedures | Process owner approval and applicable security, privacy, governance, or release review | Controlled operational change; rollback and training impact documented | Guides execution but cannot redefine architecture or authority. |
| **Evidence** | Evidentiary value rather than policy authority | Review records, approvals, test results, release records, risk acceptance, provenance, chain-of-custody records | Authenticated creation by an authorized role; integrity and retention controls | Immutable or append-only where feasible; corrections are additive and traceable | Demonstrates whether requirements were satisfied; does not create requirements by itself. |
| **Informative** | No independent normative authority | Explanations, summaries, indexes, examples, non-binding mappings, background | Owner review for accuracy; publication review if public | Normal documented revision; status must remain informative | May explain implementation but cannot authorize or require it. |
| **Template** | No independent authority beyond the approved source requirements it represents | Reusable structure for records, reviews, tests, ADRs, evidence, and reports | Approval by the owner of the governing requirement | Controlled when changes affect required fields or evidence quality | Completed instances receive their own classification and status. |
| **Research** | No normative authority unless incorporated through change control | Research questions, hypotheses, literature notes, experiments, simulations, unresolved findings | Research Reviewer approval for method and accurate classification; other reviews when risk or publication requires | Preserve sources, method, uncertainty, and revisions; formal change required before findings alter architecture | May inform proposals; cannot directly define production or normative behavior. |

### 4.1 Classification Integrity

Normative requirements shall be clearly distinguishable from guidance, examples, research hypotheses, and implementation notes. A document shall not use an informative label to avoid required approval for binding content.

---

## 5. Mandatory Document Metadata and Structure

Every governed document shall contain:

- Document ID;
- Title;
- Version;
- Status;
- Classification;
- Authority or owner;
- Purpose;
- Scope;
- Dependencies;
- Normative references;
- Cross-references;
- Assumptions;
- Constraints;
- Security considerations;
- Validation criteria;
- Revision history.

Metadata shall be explicit, current, internally consistent, and machine-searchable in plain Markdown or an approved machine-readable representation.

### 5.1 Approved Exceptions

The following limited exceptions are permitted:

1. **Canonical authority exception.** `01_CANONICAL_PROJECT_DEFINITION.md` is superior to this Constitution. Its existing approved metadata is accepted as sufficient for constitutional precedence. This Constitution cannot require a modification to its conceptual content. Any canonical metadata change requires explicit Founder approval under the Canonical Project Definition's own change control.
2. **Short records.** A short ADR index entry, decision-log entry, review record, or evidence receipt may use a compact schema when its governing policy defines all mandatory fields and the record contains a unique identifier, status, owner or signer, timestamp, related authority, and integrity reference.
3. **Templates.** A template may represent fields as instructions rather than completed values, but it shall identify its governing document, version, classification, owner, and intended record type.
4. **Machine-readable files.** A machine-readable artifact may store required metadata in a sidecar manifest, header, schema field, signed envelope, or repository manifest when direct embedding would break the format. The linkage shall be unambiguous and version-controlled.
5. **Generated evidence.** Automated test or audit output may rely on an authenticated evidence envelope that supplies identity, timestamp, version, environment, and integrity metadata.

An exception shall not remove traceability, ownership, status, security, or integrity requirements.

### 5.2 Status Integrity

A document must not be treated as approved merely because it exists, is committed, is named “baseline,” is included in an archive, or has a version number. Approval requires evidence of the applicable gates and approval authority.

---

## 6. Document Lifecycle

The governed lifecycle is:

> **Draft → Under Review → Approved → Implemented → Validated → Published → Superseded → Archived**

| State | Entry Criteria | Exit Criteria | Responsible Roles | Permitted Changes | Required Evidence |
|---|---|---|---|---|---|
| **Draft** | Owner assigned; purpose and scope identified; source authority known | Mandatory metadata substantially complete; dependencies identified; review package opened | Contributor; document owner | Broad edits permitted; no reliance as approved authority | Draft history; source notes; identified assumptions and limitations |
| **Under Review** | Draft passes completeness screening; affected reviewers assigned | Required gates completed; issues resolved or explicitly accepted; approval decision recorded | Documentation Authority and applicable reviewers | Controlled edits addressing review findings; substantive changes may restart affected gates | Review comments; findings; dispositions; gate records |
| **Approved** | Required approvals recorded; version fixed; no unresolved blocking conflict | Implementation applicability assessed and, where applicable, implementation work authorized | Stated authority; Project Founder for constitutional and v1.0 normative baselines | Only controlled changes; no silent substantive edits | Approval record; approved version hash or equivalent integrity reference |
| **Implemented** | Approved requirements have an identified conforming implementation, or the document records “not implementation-applicable” with justification | Acceptance criteria executed or validation package prepared | Implementation Reviewer; implementation owner | Implementation corrections under change control; documentation remains authoritative | Traceability to implementation; build or configuration identity; implementation status |
| **Validated** | Implementation or process has testable criteria and validation evidence | Validation accepted; residual limitations recorded; publication eligibility assessed | Security, privacy, architecture, implementation, and other applicable reviewers | Corrections require a new evidence record and affected version review | Test results; security validation; review sign-off; residual-risk record |
| **Published** | Publication and release gate passed; public classification and limitations accurate | Replaced by a later approved version or formally withdrawn | Release Reviewer; Documentation Authority | No alteration of published history; corrections require a new version | Release record; tag or immutable reference; publication checklist |
| **Superseded** | A later approved artifact explicitly replaces all or part of the document | Archival package complete; references redirected without erasing history | Documentation Authority; owning authority | No substantive edits; annotations may point to successor | Supersession record; successor linkage; dependency impact record |
| **Archived** | Retention, integrity, and access controls confirmed | Terminal state unless restored through a recorded governance decision | Documentation Authority; evidence custodian where applicable | Read-only except additive archival metadata | Archive manifest; integrity verification; retention and access record |

### 6.1 Lifecycle Applicability

Where a state is not applicable, the document shall record **Not Applicable** with justification and approval. States shall not be skipped merely for convenience.

### 6.2 Rollback

Rollback shall restore a previously approved and validated state through a new controlled action. It shall not erase the failed or withdrawn version. Rollback records shall identify cause, affected artifacts, security and privacy effects, evidence preservation, and the restored version.

### 6.3 Archival

Superseded and archived artifacts shall remain retrievable for audit, subject to lawful retention and access restrictions. Links shall identify the successor without rewriting historical content.

---

## 7. Roles and Responsibilities

One person may hold multiple roles in a small research project only when the conflict is disclosed and compensating review is documented. No person may self-approve a material change in a role that requires independent review when another qualified reviewer is available. The Project Founder retains final authority over constitutional and normative baselines.

| Role | Authority | Responsibilities | Prohibited Conflicts | Required Approvals or Participation |
|---|---|---|---|---|
| **Project Founder** | Final authority over the Canonical Project Definition, Constitution, normative v1.0 baseline, major scope, and constitutional amendments | Approve constitutional and normative baselines; resolve escalated authority conflicts; approve major change and risk acceptance at Founder level | Shall not represent unreviewed work as approved; shall not bypass recorded gates | Explicit approval for constitutional amendments, v1.0 freeze, and changes requiring Founder authority |
| **Documentation Authority** | Administrative authority over document governance and integrity | Enforce metadata, lifecycle, naming, versioning, cross-references, review records, supersession, and archives | Cannot amend the Constitution or approve technical correctness outside competence | Documentation and terminology gate; administration of change and approval records |
| **Architecture Reviewer** | Review authority for architectural consistency | Verify scope, actors, assets, boundaries, dependencies, authorization, trade-offs, ADR alignment, and implementation consistency | Shall not approve an architecture solely because they authored its implementation | Architectural consistency gate for architecture-affecting artifacts |
| **Security Reviewer** | Review authority for security requirements and threat treatment | Review threats, controls, trust boundaries, secure defaults, evidence integrity, failure handling, residual risk, and validation | Shall not accept unresolved material risk without authorized risk acceptance | Security gate for architecture, implementation, connectors, evidence, releases, and material changes |
| **Privacy and Governance Reviewer** | Review authority for privacy, purpose limitation, fundamental rights, governance, and accountability | Review data minimization, lawful and authorized scope, human oversight, role allocation, retention, publication risk, and rights impacts | Shall not substitute a framework mapping for legal analysis or claim legal compliance without evidence | Privacy and fundamental-rights gate when data, people, monitoring, evidence, or public claims are affected |
| **Implementation Reviewer** | Review authority for conformance between approved design and implementation | Verify code, configuration, interfaces, dependencies, tests, operational controls, and declared implementation status | Shall not let implementation silently redefine architecture | Implementation-alignment gate and implemented-state evidence |
| **Research Reviewer** | Review authority for research method and claim classification | Verify sources, citation fit, facts versus proposals, experiment design, uncertainty, reproducibility, and limitations | Shall not promote hypotheses or simulations as validated architecture | Research review for research records and material external claims |
| **Release Reviewer** | Review authority for release readiness and publication integrity | Verify status, evidence, version, secrets, personal data, safety boundaries, reproducibility, changelog, and rollback readiness | Shall not release artifacts with unresolved blocking findings or misleading status | Publication and release gate |
| **Contributor** | Authoring authority within assigned scope; no approval authority unless separately assigned | Produce accurate, traceable, secure, and reviewable artifacts; disclose assumptions and limitations; resolve findings | Shall not self-declare approval, conceal conflicts, fabricate evidence, or expand scope | Submit work to required reviewers; preserve review history |

### 7.1 Separation of Duties

At minimum, the author of a material architectural, security, privacy, or release change shall not be its sole reviewer and approver. Where project staffing makes independent review impossible, the conflict shall be recorded, approval shall be escalated to the Project Founder, and the artifact shall disclose the limitation.

---

## 8. Review and Approval Gates

| Gate | Mandatory When | Minimum Retained Evidence |
|---|---|---|
| **1. Scope and dependency review** | Every new document; every major or minor change; every consolidation action | Approved scope; dependency list; affected artifacts; out-of-scope confirmation; conflict check |
| **2. Architectural consistency review** | Architecture, identity, authorization, trust, lifecycle, evidence, connector, implementation, or material technical changes | Architecture review record; boundary and dependency analysis; ADR references; consistency findings |
| **3. Security review** | Any change affecting identity, credentials, authorization, data, models, tools, connectors, runtime, evidence, logs, network, supply chain, release, or threat assumptions | Threat analysis; control assessment; residual risk; required tests; reviewer decision |
| **4. Privacy and fundamental-rights review** | Any processing of personal or case-related data; monitoring; correlation; retention; publication; human-impacting output; jurisdiction or purpose change | Data-purpose analysis; minimization; rights and accountability assessment; constraints; reviewer decision |
| **5. Terminology and documentation review** | Every governed document and public release | Metadata check; terminology check; classification and status check; cross-reference validation; editorial findings |
| **6. Implementation-alignment review** | Before implementation begins; when implementation status changes; before validation or release of implemented features | Requirement-to-implementation traceability; conformance findings; exception records; implementation status evidence |
| **7. Publication and release review** | Every public or tagged release; every portfolio publication; every release of evidence or demonstration material | Release checklist; secret and personal-data checks; safety review; status accuracy; version and rollback record |

A gate may be marked not applicable only with documented rationale approved by the Documentation Authority and the relevant specialist reviewer.

Blocking findings shall prevent advancement. Non-blocking findings shall have an owner, due condition, and accepted disposition.

---

## 9. Change Control

Every material change shall have a change record containing:

- change identifier;
- rationale;
- affected documents;
- dependency analysis;
- security impact;
- privacy impact;
- implementation impact;
- migration requirements;
- validation plan;
- rollback plan;
- approval record;
- version change.

### 9.1 Change Classes

| Change Class | Definition | Minimum Control |
|---|---|---|
| **Editorial** | Spelling, formatting, broken-link repair, or clarification that does not alter requirements, authority, scope, meaning, control strength, dependencies, interfaces, or expected behavior | Change identifier; reviewer confirmation of non-material effect; patch increment when the published artifact changes |
| **Minor** | Backward-compatible substantive addition, clarification with operational value, new traceability, or strengthened control that does not invalidate conforming dependencies | Full impact record; affected review gates; minor version increment; migration notes when needed |
| **Major** | Incompatible normative or architectural change; authority, scope, trust, identity, authorization, evidence, privacy, lifecycle, interface, classification, or methodology change that invalidates prior conformance | Full impact record; ADR where architectural; all applicable gates; explicit Founder approval for normative baseline; major version increment; migration and rollback plan |

A change described as editorial shall be reclassified when it changes interpretation, obligation, permitted behavior, or evidence requirements.

### 9.2 ADR Requirement

A material architectural decision requires an ADR. This includes decisions affecting trust boundaries, identity, delegation, authorization, data flow, evidence, security controls, privacy constraints, component responsibilities, interfaces, state transitions, or material trade-offs.

Accepted ADRs shall not be silently rewritten. A correction that changes an accepted decision requires a new ADR that supersedes the earlier ADR. The earlier ADR remains immutable and linked to its successor.

### 9.3 Change Freeze During Active Work

During an approved consolidation, implementation, validation, or release phase:

- unapproved methodological changes are prohibited;
- scope shall remain fixed;
- proposed alternatives shall be logged without redirecting active work;
- only approved change records may alter dependencies, sequence, required evidence, or acceptance criteria;
- urgent security containment may temporarily suspend work, but permanent changes still require retrospective formalization and approval.

---

## 10. Versioning

Governed documents shall use semantic versioning:

- **Major** — incompatible normative or architectural change;
- **Minor** — backward-compatible substantive addition;
- **Patch** — correction, clarification, or non-material improvement.

Version numbers shall use `MAJOR.MINOR.PATCH`. Pre-release labels may be used for drafts but do not confer approval.

### 10.1 Version Synchronization

Interdependent documents need not have identical version numbers. They shall, however:

- identify the version or minimum compatible version of material dependencies;
- be reviewed together when one change makes another dependency incomplete, inconsistent, or incompatible;
- be included in the same controlled change set when compatibility requires coordinated revision;
- update cross-references and traceability records before approval;
- prevent release of a mixed baseline known to contain incompatible versions.

A baseline version is separate from each document version. The baseline manifest shall record the exact approved version and integrity reference of every member.

### 10.2 Status and Version

Status and version are independent. Version `1.0.0` does not imply approval. A draft may have a stable version number; an approved document requires approval evidence.

---

## 11. Identifiers and Naming

### 11.1 Document Identifiers

Document IDs shall use:

> `OBDIA-<DOMAIN>-<NUMBER>`

The domain shall be a stable uppercase abbreviation. The number shall be a stable zero-padded sequence consistent with existing identifiers. An assigned identifier is immutable and shall not be reused.

Renaming a file, changing a title, changing a classification, or superseding a document shall not change its identifier.

### 11.2 File Names

- Normative document file names: `UPPER_SNAKE_CASE.md`, with the approved numerical prefix where the Knowledge Pack sequence requires it.
- ADR file names: `ADR-XXXX-short-title.md`.
- Diagram file names: `kebab-case.<extension>`.
- Machine-readable evidence and record names shall include or resolve unambiguously to an immutable unique identifier.

A filename shall not claim an approval state or classification that differs from metadata.

### 11.3 Evidence Identifiers

Each evidence record shall have an immutable unique identifier generated by an approved scheme. The identifier shall remain stable across storage, export, verification, and archival. Corrections shall create linked additive records rather than reuse or overwrite the identifier.

### 11.4 Status Notation

The authoritative status shall appear in metadata using an approved lifecycle term. Informal labels such as “enterprise,” “final,” “baseline,” or archive placement shall not substitute for lifecycle status.

---

## 12. Traceability

Every material requirement shall be traceable to:

- its originating authority;
- related architecture;
- relevant threats and risks;
- controls;
- implementation;
- tests;
- validation evidence.

Traceability shall be bidirectional:

- a requirement shall point forward to decisions, controls, implementation, tests, and evidence;
- implementation, tests, and evidence shall point backward to the governing requirement and decision;
- risks and threats shall identify the controls and validation that address them;
- superseding artifacts shall link to predecessors and affected dependents.

### 12.1 Requirement Identifiers

Material normative requirements shall have stable identifiers within their document. Requirement identifiers remain stable across editorial changes. When a requirement is removed or superseded, its identifier shall not be silently reassigned.

### 12.2 Cross-Reference Integrity

Before approval and release:

- referenced files and identifiers shall exist;
- titles, versions, and statuses shall match authoritative metadata;
- circular dependencies shall be identified and justified or removed;
- superseded references shall be redirected without erasing history;
- broken references are blocking when they impair authority, implementation, validation, or audit.

---

## 13. Architectural Consistency

Every architecture artifact shall document:

- purpose;
- scope;
- actors;
- assets;
- trust boundaries;
- dependencies;
- authorization boundaries;
- threats;
- controls;
- residual risks;
- validation criteria;
- implementation status.

Every architecture artifact shall remain consistent with:

- human accountability;
- officer-bound delegation;
- Zero Trust;
- least privilege;
- purpose limitation;
- evidence integrity;
- chain of custody;
- privacy and fundamental rights.

### 13.1 Human Authority Boundary

No architecture artifact may attribute independent legal authority, autonomous coercive authority, unrestricted investigative authority, or institutional discretion to an AI agent. Consequential decisions remain assigned to authorized human and institutional actors.

### 13.2 Architecture and Implementation Separation

Architecture defines approved responsibilities, boundaries, constraints, and behavior. Implementation realizes them. Neither prototypes nor existing code shall be treated as architecture unless the relevant decision is formally approved and documented.

### 13.3 Status Accuracy

Architecture shall state accurately whether each element is conceptual, designed, partially implemented, implemented, tested, simulated, deferred, or out of scope.

---

## 14. Research and Evidence Discipline

Documents shall distinguish clearly among:

- established facts;
- legal requirements;
- standards;
- guidance;
- existing implementations;
- proposed architecture;
- assumptions;
- forecasts;
- simulations;
- experimental results;
- research hypotheses.

### 14.1 Source Discipline

Citations shall support the exact claims they accompany. Primary and authoritative sources shall be prioritized. Current technologies, active standards, and legal or regulatory claims shall be verified for currency before reliance.

No fabricated source, title, identifier, date, quotation, legal conclusion, metric, test result, validation, certification, deployment claim, or compliance claim is permitted.

### 14.2 Legal and Compliance Accuracy

Documentation shall distinguish legally binding requirements from standards, frameworks, guidance, recommendations, and project-specific design choices. A mapping does not by itself demonstrate compliance. Legal conclusions require an identified basis and appropriate qualified review; the project shall not imply official approval.

### 14.3 Experimental and Simulation Records

An experiment or simulation shall record:

- question or hypothesis;
- scope and environment;
- synthetic or authorized data source;
- method;
- inputs and versions;
- controls;
- results;
- limitations;
- reproducibility information;
- reviewer;
- relationship to architecture.

An experimental result cannot become normative architecture without formal review and change control.

---

## 15. Diagram Governance

Acceptable diagram formats include:

- Mermaid;
- C4;
- Data Flow Diagrams;
- Sequence Diagrams;
- Trust Boundary Diagrams;
- State Diagrams;
- Attack Trees;
- Authorization Flows.

Every governed diagram shall include or link unambiguously to:

- purpose;
- scope;
- version;
- related documents;
- legend where necessary;
- trust boundaries where applicable.

Diagrams shall identify actors, data flows, authorization boundaries, and external systems accurately. A diagram shall not imply implemented capability when it is conceptual.

No document is required to contain every diagram type. A diagram is required only when it materially improves technical understanding, resolves ambiguity, supports threat modeling, or enables validation.

The source representation and rendered output shall be version-linked. A diagram conflict with approved text is a documented conflict; neither is silently treated as authoritative without resolution under the normative hierarchy.

---

## 16. Security, Ethics, and Legal Boundaries

All documentation, research, implementation, tests, and demonstrations shall prohibit:

- unauthorized access;
- offensive cyber operations;
- credential theft;
- malware deployment;
- uncontrolled interaction with criminal infrastructure;
- unlawful surveillance;
- unlawful deanonymization;
- autonomous coercive decisions;
- independent legal authority for AI agents;
- real-person investigations in public demonstrations;
- use of unlawfully obtained data.

All public demonstrations shall use only:

- synthetic data;
- fictitious identities;
- mock services;
- local laboratories;
- testnets;
- lawful public sources;
- archived authorized datasets;
- controlled cybersecurity environments.

### 16.1 Environment Distinction

Documentation shall distinguish the Surface Web, authenticated or non-indexed Deep Web services, operationally isolated Dark Web research environments, Web3 systems, cloud platforms, and institutional systems. Deep Web shall not be used as a synonym for Dark Web.

### 16.2 Dual-Use Treatment

A feature with material abuse potential shall be constrained to defensive, read-only, simulated, synthetic, testnet, mock, or controlled-laboratory use. Documentation shall record the constraint, abuse case, control, and validation.

### 16.3 Data Legitimacy

Data provenance, authorization, purpose, minimization, retention, access, and disclosure constraints shall be recorded. Uncertain authorization is treated as not authorized.

---

## 17. Implementation Governance

Implementation may begin only when:

- scope is approved;
- relevant architecture is approved;
- an ADR exists where required;
- threat modeling is complete;
- authorization boundaries are defined;
- security requirements are testable;
- acceptance criteria exist.

Additional privacy, evidence, release, or operational approvals shall be completed when applicable.

Implementation shall not silently define architecture. When code, configuration, deployment behavior, tests, or operational practice conflicts with approved documentation, the implementation is non-conforming until the conflict is formally resolved.

### 17.1 Secure Defaults

Implementations shall default to deny, minimum privilege, explicit authorization, revocable credentials, controlled data access, auditable actions, isolated high-risk operations, and safe failure.

### 17.2 Implementation Status

Claims of implementation shall identify the exact component, version, environment, limitations, and validation status. A mock, stub, diagram, test fixture, or simulation is not an implemented operational capability.

### 17.3 Exceptions

An implementation exception requires:

- identifier;
- affected requirement;
- rationale;
- security and privacy impact;
- compensating controls;
- expiration or review condition;
- owner;
- approval;
- validation evidence.

An exception cannot grant prohibited capability or independent agent authority.

---

## 18. Quality Criteria

Every governed artifact shall meet minimum standards for:

- completeness;
- internal consistency;
- technical accuracy;
- clarity;
- traceability;
- security;
- privacy;
- maintainability;
- reproducibility;
- testability;
- recruiter and reviewer readability;
- explicit limitations;
- accurate implementation status.

The following are prohibited:

- empty placeholders;
- superficial policy statements presented as enterprise baselines;
- unsupported superlatives;
- fabricated compliance claims;
- duplicated rules without authoritative ownership;
- hidden contradictions;
- ambiguous status;
- untraceable requirements;
- misleading implementation or validation claims;
- unexplained scope expansion.

A concise document may be acceptable only when it remains complete for its purpose and satisfies all applicable requirements. Brevity does not excuse omission of authority, controls, evidence, or limitations.

---

## 19. Conflict Resolution Procedure

When two governed artifacts conflict:

1. apply the normative hierarchy;
2. identify the conflict explicitly;
3. suspend affected lower-level decisions or behavior where necessary;
4. assign an owner and record the resolution path;
5. assess dependencies, security, privacy, evidence, and implementation impact;
6. issue an ADR if the conflict concerns architecture;
7. obtain required approvals;
8. increment affected versions;
9. update cross-references and traceability;
10. retain evidence of the conflict and resolution.

No conflict may be resolved silently. A conflict that affects authorization, evidence integrity, privacy, human accountability, or prohibited capability is blocking until resolved.

---

## 20. Baseline Freeze

The baseline declaration is:

> **OBDIA Knowledge Pack Enterprise v1.0**

Freeze requires:

- all 30 governed documents reviewed;
- metadata normalized;
- identifiers validated;
- terminology harmonized;
- dependencies verified;
- conflicts resolved;
- gaps documented;
- cross-references validated;
- Founder approval.

Freeze also requires an exact manifest of file names, document IDs, versions, classifications, statuses, integrity references, and approval evidence.

A freeze does not mean that every proposed capability is implemented. It means the documentation baseline is approved, internally consistent, correctly classified, traceable, and suitable to govern subsequent work.

After freeze, every change requires formal change control and an appropriate version increment. No post-freeze artifact may be substituted into the baseline by archive replacement, filename reuse, or undocumented edit.

---

## 21. Audit and Evidence

The project shall retain evidence of:

- reviews;
- approvals;
- rejected changes;
- ADR decisions;
- risk acceptance;
- validation results;
- releases;
- supersession;
- archival.

Audit history shall be tamper-evident where technically feasible. At minimum, records shall identify the actor or authenticated workload, timestamp, action, affected artifact, previous and resulting state, reason, approval, and integrity reference.

Rejected and withdrawn changes shall remain discoverable for governance and audit, subject to lawful retention and access restrictions.

Evidence shall be protected against unauthorized alteration, deletion, cross-case contamination, misleading reclassification, and loss of provenance.

---

## 22. Constitutional Amendments

An amendment requires:

- written proposal;
- rationale;
- impact assessment;
- review of all affected normative documents;
- explicit Project Founder approval;
- major or minor version increment as appropriate;
- recorded effective date.

A major increment is required when the amendment changes constitutional authority, hierarchy, baseline scope, accountability, prohibited boundaries, lifecycle, or mandatory control model in an incompatible way. A minor increment may be used for a backward-compatible substantive addition. A patch may correct a non-material error without altering constitutional meaning.

No implicit amendment is valid. Custom, convenience, implementation behavior, repository structure, a legacy roadmap, an ADR, or a lower-level policy cannot amend this Constitution.

---

## 23. Definitive Artifact Validation Checklist

An artifact is eligible to advance only when every applicable item below is answered **Yes**, **Not Applicable with approved rationale**, or **Blocked**. A blocked item prevents advancement.

### Authority and Scope

- [ ] The originating authority is identified.
- [ ] The artifact's purpose and scope are explicit.
- [ ] The artifact remains within the 01–30 baseline and approved project scope.
- [ ] The artifact does not amend the Canonical Project Definition.
- [ ] The artifact does not introduce independent legal authority for an AI agent.
- [ ] The classification and hierarchy level are correct.
- [ ] The owner and approval authority are identified.
- [ ] The lifecycle status is accurate and evidenced.

### Metadata and Structure

- [ ] Document ID, title, version, status, and classification are present.
- [ ] Purpose, scope, dependencies, and normative references are present.
- [ ] Cross-references, assumptions, and constraints are present.
- [ ] Security considerations and validation criteria are present.
- [ ] Revision history is present.
- [ ] Any metadata exception is authorized and traceable.
- [ ] The filename and identifier comply with naming rules.
- [ ] The identifier has not been changed or reused.

### Consistency and Traceability

- [ ] No lower-level statement contradicts a higher-level authority.
- [ ] Dependencies exist, use correct identifiers, and have compatible versions.
- [ ] Circular dependencies are removed or explicitly justified.
- [ ] Material requirements have stable identifiers.
- [ ] Requirements trace to threats, risks, controls, implementation, tests, and evidence where applicable.
- [ ] Implementation and evidence trace back to requirements and ADRs.
- [ ] Superseded references identify their successors without erasing history.
- [ ] Terminology matches the approved glossary.

### Architecture and Security

- [ ] Actors, assets, trust boundaries, and authorization boundaries are identified where applicable.
- [ ] Threats, controls, residual risks, and failure handling are documented.
- [ ] Human accountability and officer-bound delegation are preserved.
- [ ] Zero Trust, least privilege, purpose limitation, and secure defaults are preserved.
- [ ] Evidence provenance, integrity, and chain of custody are addressed where applicable.
- [ ] Prompt injection, tool misuse, memory or data poisoning, and connector risk are addressed where applicable.
- [ ] Revocation, containment, and recovery are addressed where applicable.
- [ ] Security requirements are testable.

### Privacy, Ethics, and Legal Boundaries

- [ ] Data purpose, minimization, source authorization, retention, and access are documented where applicable.
- [ ] Privacy and fundamental-rights impacts are reviewed where applicable.
- [ ] Jurisdiction and authorization limits are explicit.
- [ ] Prohibited activities are neither authorized nor enabled.
- [ ] Demonstrations use only approved synthetic, fictitious, mock, testnet, lawful, archived-authorized, or controlled-laboratory resources.
- [ ] Deep Web and Dark Web terminology is used correctly.
- [ ] Legal requirements, standards, guidance, and project choices are distinguished.
- [ ] No mapping is represented as proof of legal compliance.

### Research and Evidence Quality

- [ ] Facts, requirements, standards, guidance, implementations, proposals, assumptions, forecasts, simulations, experiments, and hypotheses are distinguished.
- [ ] Citations support the exact claims they accompany.
- [ ] Primary and authoritative sources are prioritized.
- [ ] Current external claims have been checked for currency.
- [ ] No source, metric, result, validation, or legal conclusion is fabricated.
- [ ] Methods, inputs, versions, results, limitations, and reproducibility are recorded for experiments and simulations.
- [ ] Evidence has an immutable unique identifier and integrity protection where applicable.

### Implementation and Validation

- [ ] Implementation began only after required approvals and threat modeling.
- [ ] An ADR exists for each material architectural decision.
- [ ] Accepted ADRs have not been silently rewritten.
- [ ] Acceptance criteria are explicit.
- [ ] Implementation status is accurate and limited to evidenced capability.
- [ ] Tests include positive, negative, security, regression, and boundary cases as applicable.
- [ ] Validation evidence identifies the tested version and environment.
- [ ] Non-conformance, exceptions, residual risks, and limitations are explicit.
- [ ] Rollback and archival requirements are satisfied.

### Review, Approval, and Release

- [ ] Required review gates are complete.
- [ ] Reviewer conflicts of interest are disclosed and mitigated.
- [ ] Blocking findings are resolved.
- [ ] The approval record identifies the approved version and authority.
- [ ] Publication contains no secrets, real case data, unauthorized material, or misleading claims.
- [ ] Release and supersession records preserve audit history.
- [ ] The artifact is recruiter- and reviewer-readable without sacrificing technical accuracy.
- [ ] The artifact contains no empty placeholders or superficial claims presented as complete controls.

---

## 24. Revision History

| Version | Date | Status | Change | Authority |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Approved Constitutional Baseline | Initial constitutional baseline establishing hierarchy, fixed 30-document scope, classifications, metadata, lifecycle, roles, gates, change control, versioning, traceability, architecture and research discipline, safety boundaries, implementation governance, conflict resolution, freeze criteria, audit, amendments, and validation. | Project Founder |
