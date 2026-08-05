# ARCHITECTURE DECISION RECORD (ADR) POLICY

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-ADR-001 |
| **Title** | Architecture Decision Record (ADR) Policy |
| **Version** | 1.1.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Standardize how material architectural, security, privacy, governance and implementation decisions are proposed, evaluated, approved, recorded, superseded and traced throughout OBDIA. |
| **Scope** | Architecture Decision Records, ADR indexes, decision reviews, related repository changes, implementation traceability and decision evidence within the approved Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `02_PROJECT_GLOSSARY.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `06_IDENTITY_AND_DELEGATION_MODEL.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GLOSS-001`; `OBDIA-RES-001`; `OBDIA-ARCH-001`; `OBDIA-ID-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001` |
| **Cross-References** | Documents 11–30; accepted ADRs; architecture, threat, privacy, implementation, validation, risk and release records |
| **Assumptions** | Material decisions may evolve as research, architecture and implementation mature; historical decision context must therefore remain reviewable and immutable in substance. |
| **Constraints** | ADRs shall not override higher-level authority, grant independent legal authority to an AI agent, authorize prohibited conduct, silently expand scope or substitute for required specialist review. |
| **Security Considerations** | Weak decision records can conceal unsafe trade-offs, authorization bypasses, undocumented trust assumptions, evidence-integrity failures, privacy harm, dependency risk and non-conforming implementation. |
| **Validation Criteria** | Every Accepted ADR has an immutable identifier, complete decision context, evaluated alternatives, applicable reviews, explicit consequences and risks, bidirectional traceability, exact approval evidence and controlled supersession. |
| **Implementation Relationship** | Material implementation shall reference the Accepted ADRs that justify it; an ADR does not itself prove that implementation exists, is secure, is validated or is legally authorized. |

---

## 1. Purpose

This policy governs Architecture Decision Records for the Officer-Bound Digital Investigation Agent project.

An ADR captures a material decision, why it was made, which alternatives were considered, which requirements and risks it affects, and how the decision will be validated. ADRs preserve decision history and prevent architecture from being defined silently through implementation, diagrams, conversations or undocumented assumptions.

An ADR is subordinate to the Canonical Project Definition, Master Documentation Constitution and approved normative Knowledge documents. An ADR may select among conforming alternatives, but it may not amend or contradict higher-level authority.

This document preserves the original ADR policy requirements: mandatory ADR structure, six ADR statuses, immutable identifiers, filename convention, review requirements, implementation traceability and supersession rather than silent rewriting.

## 2. Fundamental ADR Rules

- **ADR-REQ-001:** Every material architectural, security, privacy, governance or implementation decision shall be recorded in an ADR.
- **ADR-REQ-002:** An ADR shall not contradict a higher-level authority.
- **ADR-REQ-003:** An ADR shall not grant independent legal authority to an AI agent.
- **ADR-REQ-004:** An ADR shall not authorize prohibited conduct.
- **ADR-REQ-005:** An ADR shall not silently expand the normative Knowledge Pack beyond documents 01–30.
- **ADR-REQ-006:** Implementation shall not silently define architecture.
- **ADR-REQ-007:** An Accepted ADR shall be reviewable independently from oral context or private conversation.
- **ADR-REQ-008:** ADR decisions shall be attributable to identified human roles.
- **ADR-REQ-009:** An AI agent may support drafting or analysis but shall not approve, accept, reject, supersede or deprecate an ADR.
- **ADR-REQ-010:** Repository merge alone shall not establish ADR acceptance.
- **ADR-REQ-011:** An ADR shall distinguish existing facts, assumptions, proposals, simulations, experiments and implementation status.
- **ADR-REQ-012:** Material uncertainty and unresolved limitations shall be explicit.

## 3. When an ADR Is Required

An ADR is mandatory when a decision materially affects one or more of the following:

- system structure or component boundaries;
- human accountability;
- officer-bound delegation;
- identity, credentials or machine identity;
- authentication or authorization;
- trust domains or trust boundaries;
- policy enforcement;
- model, prompt, retrieval, memory or tool controls;
- connector design or external-system access;
- evidence provenance or chain of custody;
- audit architecture;
- data classification, retention or deletion;
- privacy or fundamental rights;
- operational isolation;
- cloud, Web3 or supply-chain architecture;
- security controls or residual risk;
- interoperability or protocol selection;
- implementation technology with durable architectural consequences;
- exception or trade-off affecting an approved principle;
- migration, deprecation or replacement of an Accepted decision;
- material conflict among governed artifacts.

- **ADR-REQ-013:** ADR applicability shall be evaluated before material implementation begins.
- **ADR-REQ-014:** A material decision shall not be divided into small changes to avoid ADR review.
- **ADR-REQ-015:** Multiple tightly coupled decisions may share one ADR only when their context, alternatives, consequences and validation remain understandable.
- **ADR-REQ-016:** Independent decisions should use separate ADRs to preserve clear supersession and traceability.
- **ADR-REQ-017:** A decision affecting a canonical prohibition cannot be authorized by ADR.
- **ADR-REQ-018:** A decision that changes a normative baseline requires the baseline change first or as an explicitly governed prerequisite.
- **ADR-REQ-019:** Emergency implementation does not eliminate the need for a retrospective ADR and governance review.
- **ADR-REQ-020:** ADR non-applicability for a potentially material decision requires a recorded rationale.

## 4. When an ADR Is Normally Not Required

An ADR is normally unnecessary for:

- typographical corrections;
- formatting-only changes;
- link repair without change of meaning;
- routine dependency patches with no material architectural effect;
- implementation details fully constrained by an existing Accepted ADR and approved specification;
- reversible local experiments that do not establish repository architecture;
- ordinary backlog prioritization;
- research notes that do not create a decision.

- **ADR-REQ-021:** A change described as editorial shall be reclassified when it changes normative meaning, trust, authorization, risk, privacy, evidence or implementation behavior.
- **ADR-REQ-022:** Repeated minor changes that collectively create a material architecture change require an ADR.
- **ADR-REQ-023:** A local experiment requires an ADR before its result becomes a governing architecture or implementation choice.
- **ADR-REQ-024:** Absence of an ADR does not permit implementation to bypass approved architecture.

## 5. ADR Identifier and Naming Convention

The mandatory filename format is:

`ADR-0001-short-title.md`

The corresponding identifier is:

`ADR-0001`

Rules:

- use four digits with leading zeros;
- allocate identifiers monotonically;
- use lowercase kebab-case for the short title;
- do not reuse identifiers;
- do not renumber ADRs after allocation;
- preserve the identifier when the title is clarified;
- identify superseding and superseded ADRs explicitly.

- **ADR-REQ-025:** ADR identifiers are immutable.
- **ADR-REQ-026:** An abandoned, Rejected, Deprecated or Superseded ADR identifier shall not be reused.
- **ADR-REQ-027:** Filename and internal identifier shall match.
- **ADR-REQ-028:** Duplicate identifiers are blocking defects.
- **ADR-REQ-029:** The short title shall be concise, descriptive and free from status or version information.
- **ADR-REQ-030:** ADR filenames shall not include `_ENTERPRISE`, personal names, branch names or mutable implementation states.
- **ADR-REQ-031:** Renaming an ADR file shall preserve history and update all affected cross-references.
- **ADR-REQ-032:** ADR references should use both identifier and repository path when ambiguity is possible.

## 6. Mandatory ADR Structure

Every ADR shall contain the original sixteen mandatory sections:

1. ADR Identifier
2. Title
3. Status
4. Date
5. Context
6. Problem Statement
7. Decision
8. Alternatives Considered
9. Security Rationale
10. Privacy Considerations
11. Trade-offs
12. Consequences
13. Risks
14. Validation Criteria
15. Related Documents
16. Review History

### 6.1 ADR Identifier

The immutable `ADR-XXXX` identifier.

### 6.2 Title

A concise description of the decision.

### 6.3 Status

One of the six approved statuses defined in Section 7.

### 6.4 Date

The proposal date and, through Review History, the dates of material status decisions.

### 6.5 Context

The relevant technical, governance, security, privacy, evidentiary and project conditions.

### 6.6 Problem Statement

The precise decision problem, including scope and constraints.

### 6.7 Decision

The selected alternative, decision boundaries, applicability and effective conditions.

### 6.8 Alternatives Considered

Viable alternatives, including the status quo where relevant, and reasons for rejection or deferral.

### 6.9 Security Rationale

Threats, security requirements, control implications, trust and authorization effects, and residual risk.

### 6.10 Privacy Considerations

Data categories, purposes, minimization, retention, disclosure, jurisdiction, rights impacts and safeguards where applicable.

### 6.11 Trade-offs

Benefits, costs, limitations and quality-attribute conflicts.

### 6.12 Consequences

Positive, negative, operational, governance, migration and maintenance consequences.

### 6.13 Risks

Known risks, owners, treatments, residual risks and unresolved uncertainty.

### 6.14 Validation Criteria

Testable conditions that demonstrate whether the decision has been implemented and behaves as intended.

### 6.15 Related Documents

Affected requirements, Knowledge documents, ADRs, diagrams, threat models, specifications, implementation components, tests and evidence.

### 6.16 Review History

Reviewers, roles, dates, findings, dispositions, approvals, rejections and status transitions.

- **ADR-REQ-033:** Every mandatory section shall contain substantive content or an explicit non-applicability rationale.
- **ADR-REQ-034:** Empty headings and placeholders are prohibited.
- **ADR-REQ-035:** The Decision section shall use testable and unambiguous language.
- **ADR-REQ-036:** The Context section shall not be rewritten to conceal why the original decision was made.
- **ADR-REQ-037:** Alternatives shall be represented fairly and shall not be fabricated.
- **ADR-REQ-038:** Security and privacy sections shall not be collapsed into generic assurances.
- **ADR-REQ-039:** Risks shall not be marked mitigated without implementation and validation evidence.
- **ADR-REQ-040:** Related Documents shall support bidirectional traceability.
- **ADR-REQ-041:** Review History shall identify the exact ADR commit or immutable artifact reviewed.
- **ADR-REQ-042:** The ADR shall state whether implementation is Conceptual, Designed, Partially Implemented, Implemented, Tested, Simulated, Deferred or Out of Scope where applicable.

## 7. ADR Status Model

The approved ADR statuses are:

### 7.1 Proposed

The ADR has been created for analysis but has not entered formal review.

Entry requirements:

- identifier allocated;
- mandatory structure present;
- owner identified;
- decision problem stated.

Permitted changes:

- substantive editing;
- addition or removal of alternatives;
- refinement of the proposed decision.

### 7.2 Under Review

The ADR is undergoing the applicable review gates.

Entry requirements:

- complete Draft-quality content;
- reviewers identified;
- dependencies and affected artifacts listed;
- exact reviewed commit recorded.

Permitted changes:

- controlled changes responding to review findings;
- each material change requires affected reviewers to reassess.

### 7.3 Accepted

The decision has completed applicable review and has been approved by the authorized human role.

Entry requirements:

- applicable gates passed;
- blocking findings resolved;
- cross-references validated;
- approval evidence retained;
- exact accepted commit recorded.

Consequences:

- decision content becomes immutable in substance;
- implementation may rely on the ADR only within its stated scope and prerequisites.

### 7.4 Superseded

A later Accepted ADR replaces all or part of the decision.

Requirements:

- successor ADR identified;
- scope of replacement explicit;
- affected implementation and documents identified;
- migration or rollback impact recorded.

### 7.5 Rejected

The proposed decision was evaluated and not adopted.

Requirements:

- rejection authority and rationale recorded;
- material findings and alternatives preserved;
- identifier retained permanently.

### 7.6 Deprecated

The decision remains historically relevant but should not be selected for new work, normally because a safer or more maintainable direction is preferred and full supersession has not yet completed.

Requirements:

- deprecation rationale;
- affected scope;
- replacement direction where known;
- restrictions on new use;
- review or removal conditions.

- **ADR-REQ-043:** Only approved ADR statuses shall be used.
- **ADR-REQ-044:** Status shall reflect retained evidence, not intent.
- **ADR-REQ-045:** Proposed does not authorize implementation.
- **ADR-REQ-046:** Under Review does not authorize implementation except an explicitly approved, isolated and reversible experiment.
- **ADR-REQ-047:** Accepted does not prove implementation or validation.
- **ADR-REQ-048:** Rejected ADRs shall not be deleted merely because they were not selected.
- **ADR-REQ-049:** Superseded and Deprecated shall not be used interchangeably.
- **ADR-REQ-050:** Status transitions shall be attributable and recorded.
- **ADR-REQ-051:** Invalid status transitions shall be rejected.
- **ADR-REQ-052:** Reopening an Accepted decision requires a new ADR or a superseding ADR, not silent reversion to Proposed.

## 8. ADR Ownership and Review Roles

Each ADR shall identify:

- author or contributor;
- decision owner;
- Documentation Authority;
- Architecture Reviewer where applicable;
- Security Reviewer where applicable;
- Privacy and Governance Reviewer where applicable;
- Implementation Reviewer where applicable;
- Research Reviewer where factual or standards claims are material;
- Release Reviewer where publication or release is affected;
- Project Founder where normative scope or major architecture requires Founder approval.

- **ADR-REQ-053:** The decision owner shall be an identified human role.
- **ADR-REQ-054:** An AI agent shall not be the decision owner or approval authority.
- **ADR-REQ-055:** Contributors shall disclose material conflicts of interest.
- **ADR-REQ-056:** A contributor shall not be represented as an independent reviewer of the same decision.
- **ADR-REQ-057:** Role concentration shall be disclosed and compensating controls recorded.
- **ADR-REQ-058:** The Documentation Authority validates identifier, filename, structure, status and cross-reference integrity.
- **ADR-REQ-059:** Specialist reviewers may block acceptance within their defined domain.
- **ADR-REQ-060:** Founder approval shall not substitute for required specialist review.
- **ADR-REQ-061:** Review comments generated by tools or AI are advisory until adopted by an authorized human reviewer.

## 9. Review Requirements

An ADR may become Accepted only after:

- consistency with the Canonical Project Definition;
- consistency with the Master Documentation Constitution;
- alignment with approved normative Knowledge documents;
- architectural consistency review;
- security review when applicable;
- privacy and fundamental-rights review when applicable;
- terminology and documentation review;
- implementation-alignment review when implementation is affected;
- publication and release review when publication is affected;
- cross-reference validation;
- version and status validation;
- recorded approval.

- **ADR-REQ-062:** Applicable review gates shall be identified before the ADR enters Under Review.
- **ADR-REQ-063:** Gate non-applicability requires a written rationale.
- **ADR-REQ-064:** A failed mandatory gate blocks acceptance.
- **ADR-REQ-065:** Conditional acceptance shall identify conditions, owner and review date.
- **ADR-REQ-066:** Material changes after review invalidate affected gate evidence.
- **ADR-REQ-067:** Review evidence shall identify exact commits and artifact versions.
- **ADR-REQ-068:** Automated checks support but do not replace human review.
- **ADR-REQ-069:** An unresolved canonical conflict blocks acceptance.
- **ADR-REQ-070:** An unresolved critical security, privacy, evidence or authorization finding blocks acceptance.
- **ADR-REQ-071:** Independent external review shall not be claimed without corresponding evidence.

## 10. Decision Analysis Requirements

A decision analysis shall address:

- decision drivers;
- mandatory constraints;
- quality attributes;
- affected actors and assets;
- trust and authorization boundaries;
- threats and abuse cases;
- security, privacy and evidence impact;
- operational and maintenance impact;
- migration and rollback;
- dependency and supply-chain impact;
- cost and complexity;
- reversibility;
- uncertainty;
- residual risk.

- **ADR-REQ-072:** Decision criteria shall be stated before alternatives are evaluated where feasible.
- **ADR-REQ-073:** Alternatives shall be evaluated against consistent criteria.
- **ADR-REQ-074:** The status quo shall be considered when it is viable.
- **ADR-REQ-075:** Cost, schedule or convenience alone shall not justify violation of higher-level requirements.
- **ADR-REQ-076:** A trade-off shall not waive a canonical prohibition.
- **ADR-REQ-077:** Temporary decisions shall identify expiry, review or replacement conditions.
- **ADR-REQ-078:** Irreversible decisions require stronger evidence and rollback or containment analysis.
- **ADR-REQ-079:** Unsupported vendor or product claims shall not determine an ADR.
- **ADR-REQ-080:** Framework mappings shall not be represented as proof of compliance.
- **ADR-REQ-081:** Legal or regulatory claims shall follow `OBDIA-RES-001`.

## 11. Security, Privacy, Ethics and Legal Boundaries

Every applicable ADR shall evaluate:

- identity and delegation;
- least privilege;
- explicit authorization;
- Zero Trust;
- prompt injection and untrusted content;
- tool and connector misuse;
- evidence provenance and chain of custody;
- logging and audit integrity;
- data minimization and purpose limitation;
- case and jurisdiction separation;
- revocation and emergency containment;
- supply-chain and dependency risk;
- failure-closed behavior;
- privacy and fundamental-rights impact;
- prohibited-capability boundaries.

- **ADR-REQ-082:** An ADR shall not authorize unauthorized access, offensive operations, credential theft, malware deployment, unlawful surveillance, unlawful deanonymization, autonomous coercive decisions or uncontrolled criminal-infrastructure interaction.
- **ADR-REQ-083:** Model output shall not be treated as authorization.
- **ADR-REQ-084:** Public accessibility shall not be treated as unrestricted lawful reuse.
- **ADR-REQ-085:** AI-generated analysis shall not be represented as original evidence.
- **ADR-REQ-086:** Public demonstrations shall use synthetic, fictitious, mock, testnet, lawful public, archived-authorized or controlled-laboratory resources.
- **ADR-REQ-087:** Residual risks shall have identified human owners and dispositions.
- **ADR-REQ-088:** Risk acceptance shall not authorize prohibited conduct.
- **ADR-REQ-089:** Legal uncertainty shall be recorded rather than converted into unsupported compliance claims.

## 12. Traceability

Every ADR shall identify impacted:

- repository files;
- diagrams;
- threat models;
- governance documents;
- normative Knowledge documents;
- specifications;
- implementation components;
- policies and configurations;
- tests;
- validation evidence;
- risks and exceptions;
- releases.

Every material implementation artifact shall reference the ADRs that justify its existence.

- **ADR-REQ-090:** Traceability shall be bidirectional.
- **ADR-REQ-091:** ADR references shall use immutable identifiers.
- **ADR-REQ-092:** Broken traceability affecting security, authorization, privacy, evidence or release is blocking.
- **ADR-REQ-093:** A superseding ADR shall identify every ADR it replaces wholly or partially.
- **ADR-REQ-094:** A superseded ADR shall identify its successor through controlled status metadata or the decision index.
- **ADR-REQ-095:** Threats, controls and tests affected by the decision shall be linked.
- **ADR-REQ-096:** Implementation not justified by an applicable Accepted ADR is non-conforming when an ADR was required.
- **ADR-REQ-097:** Traceability shall distinguish planned links from implemented and validated links.

## 13. ADR Immutability and Change Control

Accepted ADR decision content shall not be silently rewritten.

After acceptance:

- substantive change requires a new ADR;
- replacement requires a superseding ADR;
- reversal requires a new ADR explaining why;
- deprecation requires an attributable status decision;
- minimal controlled status metadata may identify Superseded or Deprecated state and successor references;
- the original context, decision, alternatives, consequences and review record remain preserved.

- **ADR-REQ-098:** Accepted decision substance is immutable.
- **ADR-REQ-099:** Git history alone shall not be the only supersession mechanism.
- **ADR-REQ-100:** A correction that changes meaning requires a new or superseding ADR.
- **ADR-REQ-101:** Non-substantive metadata correction shall be explicit, reviewable and shall not alter the historical decision.
- **ADR-REQ-102:** Supersession shall preserve the original ADR and identifier.
- **ADR-REQ-103:** Partial supersession shall identify the exact clauses or decision scope replaced.
- **ADR-REQ-104:** Rejected alternatives may be reconsidered only through a new ADR with new context or evidence.
- **ADR-REQ-105:** Deletion of an Accepted, Rejected, Superseded or Deprecated ADR is prohibited except where legal or security containment requires restricted preservation.
- **ADR-REQ-106:** Containment actions affecting ADR visibility shall preserve an authorized evidence record.

## 14. ADR Index and Repository Placement

ADRs shall be stored in a governed repository location using the approved filename convention.

The ADR index should record:

- ADR identifier;
- title;
- status;
- decision owner;
- proposal date;
- acceptance, rejection, deprecation or supersession date;
- superseding or superseded identifiers;
- affected domains;
- repository path;
- exact Accepted commit where applicable.

- **ADR-REQ-107:** The ADR index is a navigation and status aid and shall not replace the ADR itself.
- **ADR-REQ-108:** Index status shall match authoritative decision evidence.
- **ADR-REQ-109:** Index changes shall be reviewable and attributable.
- **ADR-REQ-110:** The index shall not create new decision authority.
- **ADR-REQ-111:** Missing or contradictory index entries shall be corrected before baseline freeze.
- **ADR-REQ-112:** Repository paths shall preserve clear separation among normative documents, decisions, specifications and evidence.

## 15. Implementation Governance

Implementation relying on an ADR requires:

- an Accepted ADR;
- approved relevant architecture;
- completed threat modeling;
- defined authorization boundaries;
- testable security and privacy requirements;
- acceptance criteria;
- traceability from decision to implementation and tests.

- **ADR-REQ-113:** An Accepted ADR does not independently authorize implementation when other prerequisites remain incomplete.
- **ADR-REQ-114:** Implementation status shall be stated accurately.
- **ADR-REQ-115:** Code conflicting with an Accepted ADR is non-conforming until formally resolved.
- **ADR-REQ-116:** An implementation deviation requires remediation or a new ADR before acceptance.
- **ADR-REQ-117:** Tests shall validate the ADR’s decision criteria and material consequences.
- **ADR-REQ-118:** Validation evidence shall identify exact code, policy, configuration, model, dependency and environment versions where applicable.
- **ADR-REQ-119:** A proof of concept shall not be represented as production-ready or institutionally approved.
- **ADR-REQ-120:** Public demonstrations shall not use production credentials, real investigations or uncontrolled criminal infrastructure.

## 16. Validation Checklist

Before an ADR becomes Accepted, confirm:

- [ ] identifier and filename are valid and unique;
- [ ] mandatory structure is complete;
- [ ] status is valid;
- [ ] context and problem are precise;
- [ ] the decision is unambiguous and bounded;
- [ ] viable alternatives are evaluated fairly;
- [ ] security rationale is substantive;
- [ ] privacy and fundamental-rights considerations are addressed;
- [ ] trade-offs and consequences are explicit;
- [ ] risks, owners and treatments are recorded;
- [ ] validation criteria are testable;
- [ ] higher-level consistency is verified;
- [ ] applicable review gates are complete;
- [ ] affected files, diagrams, threat models, documents and components are listed;
- [ ] bidirectional traceability is established;
- [ ] implementation status is accurate;
- [ ] migration and rollback are addressed where applicable;
- [ ] prohibited outcomes are absent;
- [ ] unresolved limitations remain visible;
- [ ] review history identifies exact commits and decisions;
- [ ] authorized human approval is retained.

## 17. Limitations

- An ADR records a decision; it does not itself implement or validate that decision.
- Acceptance does not establish legal compliance, certification, operational authorization or institutional endorsement.
- ADR analysis quality depends on available evidence, assumptions and reviewer expertise.
- Internal multi-role review is not independent external assurance.
- Not every low-level implementation choice requires an ADR.
- Excessive ADR use can reduce clarity; materiality and decision independence must therefore be assessed.
- Detailed templates, indexes and automation may support this policy but shall remain subordinate to it.
- A complete ADR history cannot guarantee that every architectural decision was identified or documented.

## 18. Change Control for This Policy

Every material change to this policy shall include:

- change identifier;
- rationale;
- affected ADRs and documents;
- dependency analysis;
- security impact;
- privacy and fundamental-rights impact;
- implementation impact;
- migration requirements;
- validation plan;
- rollback plan;
- approval record;
- semantic version change.

- **ADR-REQ-121:** Editorial corrections use a patch version when meaning is unchanged.
- **ADR-REQ-122:** Backward-compatible substantive additions use a minor version.
- **ADR-REQ-123:** Incompatible normative or ADR-methodology changes use a major version.
- **ADR-REQ-124:** Changes require Project Founder approval.
- **ADR-REQ-125:** Unapproved ADR-methodology changes during an active phase are prohibited.
- **ADR-REQ-126:** Changes to this policy shall not retroactively rewrite Accepted ADR decision substance.

## 19. Consolidation Record

Version 1.1.0 consolidates the existing ADR Policy without expanding project scope. It:

- retains immutable document identifier `OBDIA-ADR-001`;
- normalizes the authoritative filename to `10_ADR_POLICY.md`;
- replaces `Enterprise Baseline (Draft)` with constitutional lifecycle status `Draft`;
- preserves the original scope requiring ADRs for significant architectural, security, governance and implementation decisions;
- preserves the original sixteen-section mandatory ADR structure;
- preserves the six original ADR statuses: Proposed, Under Review, Accepted, Superseded, Rejected and Deprecated;
- preserves filename format `ADR-0001-short-title.md` and immutable identifiers;
- preserves consistency, security review, cross-reference, version and implementation-traceability requirements;
- preserves the rule that Accepted ADRs are not silently rewritten and must be superseded through later ADRs;
- adds materiality criteria, lifecycle entry conditions, role and gate governance, decision-analysis rules, security and privacy boundaries, bidirectional traceability, controlled status metadata, implementation conformance, validation criteria and limitations;
- treats the former `_ENTERPRISE` and non-Enterprise filenames as legacy source variants of the same immutable document rather than separate normative documents;
- creates no new normative document and authorizes no implementation.

## 20. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | Prior baseline | Enterprise Baseline (Draft) | Project origin record | Original ADR Policy containing scope, mandatory structure, statuses, review requirements, traceability, naming, immutability and validation checklist. |
| 1.1.0 | 2026-08-05 | Draft | Documentation Authority; approval reserved to Project Founder | Constitutional consolidation: normalized filename, metadata and lifecycle status; preserved original ADR requirements; added materiality, status criteria, roles, gates, decision analysis, security and privacy boundaries, traceability, immutability, implementation governance, validation and limitations. |
