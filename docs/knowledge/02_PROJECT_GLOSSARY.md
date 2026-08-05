# PROJECT GLOSSARY

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-GLOSS-001 |
| **Title** | Project Glossary |
| **Version** | 1.1.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Establish the authoritative project terminology used across the OBDIA Enterprise Architecture Framework. |
| **Scope** | Terminology governing the 01–30 Knowledge Pack, ADRs, repository documentation, specifications, diagrams, implementation, tests, evidence and research records. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `03_RESEARCH_AND_SOURCE_POLICY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-RES-001` |
| **Cross-References** | Knowledge documents 03–30 and all lower-level terminology-dependent artifacts |
| **Assumptions** | English is the authoritative repository language for normative terms. Legal meaning may depend on jurisdiction and qualified review. |
| **Constraints** | Definitions may not grant authority, expand project scope, replace legal advice, or weaken canonical security, privacy, ethics or human-accountability boundaries. |
| **Security Considerations** | Ambiguous terminology can weaken authorization, evidence integrity, privacy controls and auditability; controlled vocabulary is therefore a security and governance control. |
| **Validation Criteria** | Definitions are internally consistent, traceable to governing authorities, free of prohibited conflations, referenced terms exist, changes are impact-assessed, and approval evidence is retained. |
| **Implementation Relationship** | Implementation and tests must use these preferred terms and may not redefine them silently. |

---

## 1. Purpose

This document is the authoritative terminology source for the Officer-Bound Digital Investigation Agent project once approved. It standardizes project-specific meaning, prevents conflicting definitions, preserves the distinction between legal authority and delegated technical capability, and supports consistent architecture, security, privacy, evidence and research records.

This glossary is subordinate to `01_CANONICAL_PROJECT_DEFINITION.md` and `MASTER_DOCUMENTATION_CONSTITUTION.md`. Where a conflict exists, the higher-level authority prevails and this glossary must be corrected through formal change control.

## 2. Scope

The glossary governs terminology used in:

- normative Knowledge documents 01–30;
- Architecture Decision Records;
- repository architecture and governance documentation;
- technical specifications and diagrams;
- implementation and configuration;
- tests, validation records and evidence;
- publication and research records.

The glossary does not create legal definitions for external jurisdictions, independently authorize system behavior, or establish implementation as complete.

## 3. Terminology Authority and Use

- **GLOSS-REQ-001:** Preferred terms defined in this document shall be used consistently in governed artifacts.
- **GLOSS-REQ-002:** A lower-level artifact shall not redefine an approved glossary term without an approved change to this document or an explicit, non-conflicting domain qualification.
- **GLOSS-REQ-003:** Deep Web and Dark Web shall not be used as synonyms.
- **GLOSS-REQ-004:** The Officer-Bound Digital Investigation Agent shall not be described as an autonomous police officer, legal person, independent investigative authority or independent coercive actor.
- **GLOSS-REQ-005:** Terms describing authorization shall identify the relevant human, institutional, case, purpose, jurisdiction, source, tool, permission and time boundaries where applicable.
- **GLOSS-REQ-006:** Normative requirements, informative guidance, proposed architecture, implementation status, simulation, experimental result, assumption, forecast and research hypothesis shall remain distinguishable.
- **GLOSS-REQ-007:** Terms used in public material shall preserve the project’s non-operational, independent-research and controlled-demonstration status.
- **GLOSS-REQ-008:** A terminology conflict shall be resolved according to the constitutional hierarchy and shall not be resolved silently.

## 4. Term Ownership, Synonyms and Deprecation

- **GLOSS-REQ-009:** The Documentation Authority owns glossary administration; substantive approval remains with the Project Founder.
- **GLOSS-REQ-010:** Each term identifier is immutable. A removed or superseded identifier shall not be reused.
- **GLOSS-REQ-011:** A synonym may be listed only when it does not introduce ambiguity. The preferred term remains authoritative.
- **GLOSS-REQ-012:** A deprecated term shall remain recorded with its replacement, rationale, effective version and affected cross-references.
- **GLOSS-REQ-013:** A new term or changed definition requires dependency analysis, terminology-collision review, security and privacy impact review where applicable, and a semantic version increment.
- **GLOSS-REQ-014:** A domain-specific qualification may narrow a term for a defined context but may not contradict the unqualified glossary definition or higher-level authority.

### 4.1 Approved Synonym Rules

- `AI agent` may be used as a generic technical description only when the officer-bound and no-independent-authority constraints remain explicit in context.
- `Human Officer` and `identified human officer` refer to the same project actor when identity and accountability remain explicit.
- `Issuing Institution` and `institution` may be used interchangeably only when the institution’s issuing, supervisory and revocation role is unambiguous.
- `Dark Web research environment` is acceptable only when qualified as authorized, controlled, simulated, archived or operationally isolated.

### 4.2 Prohibited Conflations

- Deep Web is not Dark Web.
- Authentication is not authorization.
- Authorization is not legal authority.
- Agent identity is not human identity.
- Workload identity is not a permanent human credential.
- Human review is not a transfer of accountability to the AI system.
- A framework mapping is not proof of compliance.
- A proposed architecture is not an implemented system.
- A simulated capability is not an operational capability.
- Tamper-evident is not immutable.
- Publicly accessible is not equivalent to unrestricted lawful reuse.

## 5. Controlled Vocabulary

### GLOSS-TERM-001 — Officer-Bound Digital Investigation Agent

An institutionally issued AI agent cryptographically and operationally bound to one identified human law-enforcement officer. It operates only within an authorized mandate and has no independent legal authority.

### GLOSS-TERM-002 — Human Officer

The identified professional who owns the delegated operational responsibility, reviews relevant outputs and remains accountable for decisions taken with support from the agent.

### GLOSS-TERM-003 — Issuing Institution

The legally competent organization that creates, authorizes, manages, monitors and revokes the agent identity and its permissions.

### GLOSS-TERM-004 — Officer-Bound

A technical and governance condition establishing that the agent cannot operate as an independent identity. Its actions remain attributable to the authorized human officer and issuing institution.

### GLOSS-TERM-005 — Digital Professional Extension

A controlled digital representation of the officer’s professional function. It is not a psychological clone, legal person or autonomous authority.

### GLOSS-TERM-006 — Agent Identity

The verifiable machine identity assigned to the agent, including its institutional issuer, officer binding, validity period, authorization context and cryptographic credentials.

### GLOSS-TERM-007 — Workload Identity

A machine identity used by software components, services, agents or processes to authenticate securely without relying on permanent human credentials.

### GLOSS-TERM-008 — Scoped Delegation

The limited transfer of specific permissions to the agent for a defined purpose, case, jurisdiction, period, dataset and set of tools.

### GLOSS-TERM-009 — Case-Bound Authorization

Authorization that restricts the agent’s operations to one identified investigation, incident, mandate or legally approved operational context.

### GLOSS-TERM-010 — Purpose Limitation

The requirement that data and capabilities may be used only for the declared and approved investigative purpose.

### GLOSS-TERM-011 — Jurisdiction

The legal and institutional boundary within which the agent and human officer may operate.

### GLOSS-TERM-012 — Surface Web

Publicly accessible Internet content that can generally be discovered through conventional search engines.

### GLOSS-TERM-013 — Deep Web

Internet content that is not publicly indexed, including authenticated portals, private databases, institutional systems, subscription services and internal platforms. Deep Web does not inherently mean illegal content.

### GLOSS-TERM-014 — Dark Web

Services and resources accessed through anonymity networks or specialized software. Project demonstrations must use lawful, authorized, archived, synthetic or isolated environments.

### GLOSS-TERM-015 — Web3

Decentralized digital ecosystems involving blockchains, smart contracts, wallets, decentralized applications, token systems and decentralized storage.

### GLOSS-TERM-016 — Digital Investigation

The authorized collection, preservation, correlation, analysis and documentation of information from digital environments.

### GLOSS-TERM-017 — Digital Evidence

Information with potential evidentiary relevance that is collected, preserved, analyzed and documented according to integrity and accountability requirements.

### GLOSS-TERM-018 — Evidence Provenance

The documented origin and processing history of evidence, including source, collection method, timestamps, transformations and responsible entities.

### GLOSS-TERM-019 — Chain of Custody

The chronological and verifiable record describing who or what collected, accessed, transferred, analyzed or stored a piece of evidence.

### GLOSS-TERM-020 — Tamper-Evident Audit Log

A record designed so that unauthorized modifications can be detected through cryptographic or equivalent integrity controls.

### GLOSS-TERM-021 — Human Accountability

The principle that legal, professional and operational responsibility remains assigned to identifiable human and institutional actors.

### GLOSS-TERM-022 — Human Review

The examination of agent outputs or proposed actions by an authorized human before consequential decisions are accepted or executed.

### GLOSS-TERM-023 — Zero Trust

A security model based on continuous verification, explicit authorization, limited trust and least-privilege access.

### GLOSS-TERM-024 — Least Privilege

The principle that the agent receives only the minimum permissions required for the specific authorized task.

### GLOSS-TERM-025 — Separation of Duties

The distribution of sensitive responsibilities among different actors or systems to prevent unchecked authority or abuse.

### GLOSS-TERM-026 — Policy as Code

Security, authorization and governance policies expressed in machine-readable and enforceable form.

### GLOSS-TERM-027 — Authorization Gate

A technical or procedural checkpoint that must be satisfied before an action, tool or dataset can be accessed.

### GLOSS-TERM-028 — Prompt Injection

An attack in which malicious or untrusted content attempts to alter the instructions or behavior of an AI system.

### GLOSS-TERM-029 — Indirect Prompt Injection

Prompt injection delivered through external content such as websites, documents, messages, APIs or retrieved data.

### GLOSS-TERM-030 — Memory Poisoning

The manipulation of persistent agent memory by inserting false, malicious or unauthorized information.

### GLOSS-TERM-031 — Tool Misuse

The use of an approved tool in an unauthorized, unsafe or purpose-incompatible manner.

### GLOSS-TERM-032 — Operational Isolation

The separation of high-risk activities into controlled environments such as sandboxes, containers, segregated networks or dedicated research systems.

### GLOSS-TERM-033 — Emergency Containment

The immediate suspension, isolation or revocation of an agent, tool, connector or credential when risk is detected.

### GLOSS-TERM-034 — Reproducibility

The ability to reconstruct or repeat an investigative process using documented inputs, tools, versions, policies and execution records.

### GLOSS-TERM-035 — Synthetic Data

Artificially generated data that does not represent real persons, credentials, investigations or criminal infrastructure.

### GLOSS-TERM-036 — Simulated Dark Web Environment

A controlled local or laboratory environment reproducing selected technical characteristics of anonymity-network services without interacting with real criminal infrastructure.

### GLOSS-TERM-037 — Testnet

A blockchain network designed for testing and experimentation without the economic and operational consequences of a production blockchain.

### GLOSS-TERM-038 — Proposed Architecture

A design or model created by the project that has not necessarily been implemented, deployed or independently validated.

### GLOSS-TERM-039 — Experimental Implementation

A limited proof of concept used to test selected architectural assumptions.

### GLOSS-TERM-040 — Simulated Capability

A function demonstrated through synthetic data, mock systems or controlled environments rather than real operational deployment.

### GLOSS-TERM-041 — Established Fact

A claim supported by authoritative documentation, recognized standards, verified implementations or reliable academic evidence.

### GLOSS-TERM-042 — Research Hypothesis

A proposition requiring further investigation, validation or empirical testing.

### GLOSS-TERM-043 — Authorized Mandate

A formally recognized institutional, legal or operational basis that defines the investigation, incident, case or research activity the agent may support, including its purpose, jurisdiction, scope, validity period and applicable restrictions.

### GLOSS-TERM-044 — Independent Legal Authority

Legal or coercive authority exercised in the agent’s own right rather than through a valid delegation from an identified human officer and issuing institution. The Officer-Bound Digital Investigation Agent never possesses independent legal authority.

### GLOSS-TERM-045 — Cryptographic Binding

A verifiable association between the agent identity, the identified human officer, the issuing institution and the authorized mandate, protected through cryptographic credentials, signatures, attestations or equivalent integrity mechanisms.

### GLOSS-TERM-046 — Institutionally Issued

Created, authorized, governed, monitored and revocable by the issuing institution rather than self-created, privately assumed or independently asserted by the agent.

### GLOSS-TERM-047 — Authorization Boundary

The explicit limit separating permitted actions, data, tools, environments, purposes, jurisdictions or time periods from actions requiring additional authorization or from prohibited actions.

### GLOSS-TERM-048 — Additional Legal Authorization

A legal or institutional approval beyond the agent’s current delegated permissions that must be obtained before a restricted action, source, dataset, system or investigative technique may be used.

### GLOSS-TERM-049 — Prohibited Action

An action that the project, architecture, implementation or demonstration must not authorize or perform, including unauthorized access, offensive intrusion, credential theft, malware deployment, unlawful surveillance, unlawful deanonymization, autonomous coercive decisions and uncontrolled interaction with criminal infrastructure.

### GLOSS-TERM-050 — Public Demonstration

A repository demonstration, presentation, publication or reproducible exercise intended for public or portfolio use. Public demonstrations must use synthetic, fictitious, mock, testnet, lawful public, archived-authorized or controlled-laboratory resources.

### GLOSS-TERM-051 — Operationally Isolated Dark Web Research Environment

A segregated, controlled and authorized research environment used to study selected technical characteristics associated with anonymity networks without uncontrolled interaction with real criminal infrastructure.

### GLOSS-TERM-052 — Normative Requirement

A binding project rule established by an approved authority at the applicable level of the documentation hierarchy. A normative requirement is distinguishable from guidance, examples, research findings and implementation notes.

### GLOSS-TERM-053 — Informative Guidance

Non-binding explanatory material intended to aid understanding or implementation without creating independent authority or changing approved requirements.

### GLOSS-TERM-054 — Documented Assumption

A condition accepted for the purpose of analysis or design that has not been established as a verified fact and whose uncertainty, consequences and validation needs are recorded.

### GLOSS-TERM-055 — Existing Implementation

A technology, product, service, pattern or system that is documented as currently implemented or available. Its existence does not by itself establish suitability, security, legality or conformance for OBDIA.

### GLOSS-TERM-056 — Experimental Result

An observed outcome produced by a controlled experiment with recorded method, inputs, environment, versions, limitations and reproducibility information.

### GLOSS-TERM-057 — Forecast

A reasoned statement about a possible future condition or development that remains uncertain and must not be presented as an established fact.

### GLOSS-TERM-058 — Research Record

A governed record containing a research question, date, sources, method, findings, uncertainties, limitations and implications for the architecture.

### GLOSS-TERM-059 — Cross-Reference Integrity

The condition in which referenced files, identifiers, titles, versions and statuses exist, match their authoritative metadata and preserve valid predecessor, successor and dependency relationships.

### GLOSS-TERM-060 — Bidirectional Traceability

Traceability in which requirements point forward to decisions, controls, implementation, tests and evidence, while those downstream artifacts point back to their governing requirements and decisions.

### GLOSS-TERM-061 — Residual Risk

Risk remaining after applicable controls have been selected and applied. Residual risk must be recorded, owned and accepted, mitigated, transferred or rejected by an authorized role.

### GLOSS-TERM-062 — Non-Conforming Implementation

Code, configuration, workflow or deployed behavior that conflicts with an approved higher-level authority or lacks required approval, traceability, validation or evidence.

## 6. Term Change Procedure

- **GLOSS-REQ-015:** Every material terminology change shall have a change identifier, rationale, affected terms, affected documents, dependency analysis, security impact, privacy impact, migration requirements, validation plan, rollback plan, approval record and version change.
- **GLOSS-REQ-016:** Editorial corrections that do not change meaning may use a patch version; backward-compatible substantive additions require a minor version; incompatible normative meaning changes require a major version.
- **GLOSS-REQ-017:** Accepted terminology changes shall update affected cross-references and shall not silently rewrite accepted ADRs or historical evidence.
- **GLOSS-REQ-018:** During an active consolidation phase, unapproved methodological or vocabulary expansion is prohibited.

## 7. Cross-Reference Requirements

- **GLOSS-REQ-019:** Each normative Knowledge document shall use the preferred terminology and identify any necessary domain qualification.
- **GLOSS-REQ-020:** Before approval, referenced term identifiers, filenames, document identifiers, versions and statuses shall be validated.
- **GLOSS-REQ-021:** A broken or conflicting terminology reference is blocking when it affects authority, authorization, security, privacy, evidence integrity, implementation or auditability.
- **GLOSS-REQ-022:** Superseded terminology shall point to its successor without erasing historical use.

## 8. Security, Privacy, Ethics and Legal Considerations

Terminology is a control surface. Ambiguous or inflated wording can conceal unauthorized scope, misstate implementation status, weaken chain-of-custody records, create privacy risk, or imply legal authority that the agent does not possess.

- **GLOSS-REQ-023:** Terms shall preserve human accountability and the issuing institution’s responsibility.
- **GLOSS-REQ-024:** Terms shall not imply autonomous policing, autonomous coercive decision-making, unrestricted investigative authority or independent legal authority for an AI agent.
- **GLOSS-REQ-025:** Environment terms shall distinguish public Surface Web resources, authenticated or non-indexed Deep Web services, operationally isolated Dark Web research environments, Web3 systems, cloud platforms and institutional systems.
- **GLOSS-REQ-026:** Data terms shall not imply that public accessibility eliminates purpose, authorization, minimization, retention, disclosure or privacy obligations.
- **GLOSS-REQ-027:** Public-demonstration terminology shall identify synthetic, fictitious, mock, testnet, lawful public, archived-authorized or controlled-laboratory constraints.

## 9. Validation Criteria

This document is eligible to advance from Draft to Under Review only when:

- all mandatory metadata is present and internally consistent;
- all original glossary definitions are preserved or any change is explicitly justified;
- new governance-critical terms are traceable to the Canonical Project Definition, Constitution or Research and Source Policy;
- stable requirement and term identifiers are unique;
- prohibited conflations are explicit;
- dependencies and cross-references are valid;
- security, privacy, ethics and legal-boundary terminology is consistent with higher-level authority;
- the document contains no implementation or compliance claim;
- review findings and dispositions are retained;
- Project Founder approval is recorded before status becomes Approved.

## 10. Limitations

- This glossary is a project terminology authority, not a jurisdiction-independent legal dictionary.
- Definitions concerning law, jurisdiction, evidence or authorization may require qualified legal review before operational use.
- Inclusion of a technology or environment term does not authorize access or implementation.
- A definition does not demonstrate that a corresponding control or capability has been implemented or validated.
- The glossary cannot resolve an external standards conflict without source verification and formal change control.

## 11. Consolidation Record

Version 1.1.0 consolidates the existing glossary without changing the approved project scope. It:

- retains the immutable document identifier `OBDIA-GLOSS-001`;
- replaces the non-constitutional status `Initial Baseline` with lifecycle state `Draft`;
- preserves the 42 original definitions;
- adds mandatory constitutional metadata and governance sections;
- adds stable identifiers for requirements and terms;
- defines terminology ownership, synonym control, deprecation, collision resolution and cross-reference validation;
- adds only governance-critical terms already required by higher-level project authorities;
- does not approve the document, implement any capability, or expand the 01–30 normative baseline.

## 12. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | Prior baseline | Initial Baseline | Project origin record | Original glossary containing 42 project definitions. |
| 1.1.0 | 2026-08-05 | Draft | Documentation Authority; approval reserved to Project Founder | Constitutional consolidation: normalized metadata and lifecycle status, preserved original definitions, added governance rules, stable identifiers, controlled terminology procedures, validation criteria and limitations. |
