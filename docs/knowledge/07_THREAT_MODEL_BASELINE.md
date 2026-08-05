# THREAT MODEL BASELINE

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-TM-001 |
| **Title** | Threat Model Baseline |
| **Version** | 1.1.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Security Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define the mandatory threat-modeling scope, method, deliverables, review gates, traceability and validation requirements for all material OBDIA architecture and implementation artifacts. |
| **Scope** | Human, institutional, agent, workload, model, runtime, connector, data-source, storage, evidence, audit, cloud, Web3, research and repository components within the approved project scope. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `02_PROJECT_GLOSSARY.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `06_IDENTITY_AND_DELEGATION_MODEL.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GLOSS-001`; `OBDIA-RES-001`; `OBDIA-ARCH-001`; `OBDIA-ID-001` |
| **Cross-References** | Documents 08–30; accepted ADRs; architecture models; identity, trust, privacy, evidence, runtime and validation specifications |
| **Assumptions** | The project uses defensive, synthetic, authorized, simulated, testnet, archived-authorized or controlled-laboratory environments and does not rely on real criminal infrastructure or unauthorized access. |
| **Constraints** | Threat modeling shall not authorize offensive operations, live exploitation, credential theft, malware deployment, unlawful surveillance, unlawful deanonymization or autonomous coercive behavior. |
| **Security Considerations** | Incomplete or stale threat models can conceal identity compromise, prompt injection, unsafe tool use, evidence tampering, cross-case contamination, supply-chain compromise, privacy harm and authorization failure. |
| **Validation Criteria** | Every material threat model identifies assets, actors, entry points, trust and authorization boundaries, threats, abuse cases, controls, residual risks, owners, tests, assumptions, limitations and implementation status. |
| **Implementation Relationship** | No material implementation may begin until the applicable threat model and security requirements are complete, reviewable and linked to acceptance criteria. |

---

## 1. Purpose

This document establishes the mandatory threat-modeling baseline for the Officer-Bound Digital Investigation Agent project.

Threat modeling is required to identify how the project’s human-accountability, officer-bound delegation, authorization, privacy, evidence-integrity and operational-isolation guarantees may fail or be abused.

Threat modeling does not authorize attack activity. Analysis shall use architecture review, synthetic scenarios, safe test fixtures, controlled laboratories and defensive validation methods.

This document is subordinate to `01_CANONICAL_PROJECT_DEFINITION.md`, `MASTER_DOCUMENTATION_CONSTITUTION.md`, `05_ARCHITECTURE_PRINCIPLES.md` and `06_IDENTITY_AND_DELEGATION_MODEL.md`.

## 2. Fundamental Threat-Modeling Rules

- **TM-REQ-001:** Every material architecture component, feature, trust boundary, authorization path and evidence-processing flow shall have an applicable threat model.
- **TM-REQ-002:** Threat modeling shall begin before implementation.
- **TM-REQ-003:** Threat models shall be updated when architecture, trust, authorization, data flow, dependencies, implementation or threat assumptions materially change.
- **TM-REQ-004:** Threat models shall distinguish existing controls, proposed controls and unimplemented controls.
- **TM-REQ-005:** Threats shall not be treated as mitigated without control implementation and validation evidence.
- **TM-REQ-006:** Threat modeling shall include misuse by external attackers, insiders, compromised components, malicious data sources, supply-chain actors and unintended operator error.
- **TM-REQ-007:** Threat modeling shall preserve the project’s prohibition on unauthorized access and offensive cyber operations.
- **TM-REQ-008:** A missing or unreviewed threat model blocks implementation of a material feature.
- **TM-REQ-009:** Threat-model records shall identify the exact architecture, version, commit or artifact reviewed.
- **TM-REQ-010:** Threat-model limitations and unresolved uncertainty shall be explicit.

## 3. Protected Assets

Threat models shall consider at least the following assets where applicable:

- human officer identity;
- issuing institution identity;
- agent identity;
- workload identities;
- delegation artifacts;
- authorization records;
- credentials, tokens and cryptographic keys;
- policy definitions and policy versions;
- case and mandate context;
- purpose and jurisdiction attributes;
- source and tool permissions;
- evidence and derived analytical products;
- chain-of-custody records;
- audit logs;
- models and model configuration;
- prompts and system instructions;
- persistent and temporary memory;
- connector configuration;
- source data;
- cloud and local runtime environments;
- source code;
- build and release artifacts;
- research records;
- test fixtures and validation evidence;
- secrets and environment configuration.

- **TM-REQ-011:** Each threat model shall identify the confidentiality, integrity, availability, authenticity, accountability, privacy and evidentiary properties relevant to each material asset.
- **TM-REQ-012:** Asset ownership and responsible roles shall be identified.
- **TM-REQ-013:** Evidence and audit assets shall identify integrity and retention requirements.
- **TM-REQ-014:** Case, purpose and jurisdiction context shall be treated as security-relevant assets, not merely metadata.

## 4. Threat Actors

Threat actors include, where applicable:

- external attacker;
- malicious insider;
- negligent or mistaken authorized user;
- compromised human account;
- credential thief;
- malicious or compromised connector;
- malicious data provider;
- compromised external source;
- supply-chain attacker;
- malicious or compromised model, plugin or dependency;
- unauthorized administrator;
- compromised workload;
- malicious tenant or neighboring workload;
- fraudulent issuer;
- attacker controlling prompt or retrieved content;
- attacker attempting evidence tampering or audit suppression.

- **TM-REQ-015:** Threat actors shall be described by capability, access, intent and relevant constraints.
- **TM-REQ-016:** Insider threats shall include both intentional abuse and accidental misuse.
- **TM-REQ-017:** Compromised trusted components shall be modeled as hostile within affected boundaries.
- **TM-REQ-018:** The AI agent itself shall be modeled as a potentially fallible and manipulable component, not as a trusted authority.

## 5. Trust and Authorization Boundaries

Threat models shall identify at least the following boundaries where applicable:

- Human Officer ↔ Agent;
- Institution ↔ Human Officer;
- Institution ↔ Agent;
- Agent ↔ Workload;
- Agent ↔ Connector;
- Connector ↔ External Data Source;
- Runtime ↔ Storage;
- Runtime ↔ Model Service;
- Runtime ↔ Policy Engine;
- Runtime ↔ Evidence Store;
- Runtime ↔ Audit System;
- Institution ↔ Cloud Platform;
- Case Context ↔ Other Case Contexts;
- Authorized Research Environment ↔ External Networks;
- Public Demonstration Environment ↔ Operational or Private Data;
- Repository ↔ Build and Release Pipeline.

Every boundary requires explicit consideration of:

- authentication;
- authorization;
- integrity;
- confidentiality;
- logging;
- failure behavior;
- revocation;
- data minimization;
- case and purpose separation;
- trust assumptions.

- **TM-REQ-019:** Trust shall not be assumed transitive.
- **TM-REQ-020:** Authorization boundaries shall be modeled separately from network boundaries.
- **TM-REQ-021:** A boundary that allows context loss or privilege broadening is a material threat.
- **TM-REQ-022:** Failure to verify trust or authorization shall result in denial or containment.
- **TM-REQ-023:** Trust-boundary diagrams shall be included when they materially improve technical understanding.

## 6. Mandatory Threat Categories

Every applicable threat model shall evaluate at least:

- identity impersonation;
- institution or issuer spoofing;
- credential theft;
- key compromise;
- delegation forgery;
- stale authorization;
- privilege escalation;
- confused-deputy behavior;
- prompt injection;
- indirect prompt injection;
- memory poisoning;
- data poisoning;
- model manipulation;
- unsafe tool use;
- connector misuse;
- supply-chain compromise;
- evidence tampering;
- chain-of-custody corruption;
- audit-log manipulation;
- data exfiltration;
- unauthorized disclosure;
- cross-case contamination;
- cross-purpose reuse;
- jurisdiction violation;
- authorization bypass;
- malicious or malformed external content;
- denial of service;
- revocation failure;
- containment failure;
- dependency or update compromise;
- insecure defaults;
- configuration drift;
- unsafe publication;
- privacy and fundamental-rights harm.

- **TM-REQ-024:** Non-applicable categories require a documented rationale.
- **TM-REQ-025:** Threat categories shall be refined into concrete abuse cases for the reviewed architecture.
- **TM-REQ-026:** AI-specific threats shall be linked to tool, model, prompt, retrieval, memory and authorization controls.
- **TM-REQ-027:** Evidence-specific threats shall include provenance loss, transformation ambiguity, timestamp failure and unauthorized modification.
- **TM-REQ-028:** Identity-specific threats shall include officer impersonation, agent rebinding, token substitution and incomplete revocation.

## 7. Required Analysis Method

Each threat model shall document:

1. purpose and scope;
2. reviewed architecture and version;
3. actors and roles;
4. assets;
5. entry points;
6. interfaces and dependencies;
7. data flows;
8. trust boundaries;
9. authorization boundaries;
10. threat actors;
11. threats and abuse cases;
12. existing controls;
13. proposed controls;
14. assumptions;
15. security requirements;
16. validation tests;
17. residual risks;
18. risk owners;
19. unresolved findings;
20. limitations;
21. implementation status;
22. review and approval history.

- **TM-REQ-029:** A data-flow diagram shall be included for material data-processing or trust-boundary analysis unless non-applicability is documented.
- **TM-REQ-030:** Abuse cases shall describe attacker objective, preconditions, path, affected assets, expected controls and failure outcome.
- **TM-REQ-031:** Threats shall be traceable to controls and tests.
- **TM-REQ-032:** Controls shall be stated as testable requirements where feasible.
- **TM-REQ-033:** Threat models shall distinguish prevention, detection, correction, recovery and governance controls.
- **TM-REQ-034:** Threat models shall identify control dependencies and single points of failure.
- **TM-REQ-035:** Accepted residual risk shall identify an authorized owner and rationale.

## 8. Risk Evaluation

Risk evaluation shall consider:

- likelihood or plausibility;
- impact;
- affected assets;
- affected persons or rights;
- authorization and jurisdiction consequences;
- evidence-integrity consequences;
- exploitability;
- detection capability;
- containment capability;
- recovery complexity;
- dependency and supply-chain exposure;
- uncertainty.

- **TM-REQ-036:** Risk scales and criteria shall be documented and used consistently.
- **TM-REQ-037:** Risk scores shall not replace qualitative analysis.
- **TM-REQ-038:** High-impact, low-frequency threats shall not be dismissed solely because likelihood is uncertain.
- **TM-REQ-039:** Privacy, fundamental-rights and evidentiary impact shall be considered independently from conventional technical impact.
- **TM-REQ-040:** Unknown risk shall remain unresolved rather than being assigned an unsupported low rating.
- **TM-REQ-041:** Risk acceptance shall not authorize a canonical prohibition.

## 9. Control Families

Threat models shall identify controls across:

### 9.1 Preventive Controls

Examples include:

- explicit authentication and authorization;
- short-lived credentials;
- cryptographic binding;
- input isolation;
- least privilege;
- case and purpose separation;
- policy enforcement;
- network and runtime segmentation;
- read-only connectors;
- safe defaults;
- provenance preservation.

### 9.2 Detective Controls

Examples include:

- security logging;
- anomaly detection;
- authorization-denial monitoring;
- integrity verification;
- prompt-injection detection;
- evidence-hash verification;
- configuration-drift detection;
- suspicious connector or source monitoring.

### 9.3 Corrective Controls

Examples include:

- credential rotation;
- policy correction;
- evidence quarantine;
- revocation;
- configuration repair;
- compromised-component replacement;
- reprocessing from verified source data.

### 9.4 Recovery Controls

Examples include:

- trusted restoration;
- audit reconstruction;
- evidence recovery;
- state reconciliation;
- safe reauthorization;
- tested incident recovery.

### 9.5 Governance Controls

Examples include:

- ADRs;
- review gates;
- separation of duties;
- risk ownership;
- approval records;
- release controls;
- research and source policy;
- revision history.

- **TM-REQ-042:** Controls shall identify owner, implementation state and validation method.
- **TM-REQ-043:** A control described only in documentation shall not be treated as implemented.
- **TM-REQ-044:** Compensating controls shall identify the unavailable preferred control and residual limitation.
- **TM-REQ-045:** Defense in depth shall avoid multiple controls depending on the same unprotected failure point.

## 10. AI and Agent Threats

Threat models for AI-assisted functions shall address:

- prompt injection;
- indirect prompt injection;
- memory poisoning;
- retrieval poisoning;
- malicious tool output;
- fabricated or unsupported model output;
- excessive tool autonomy;
- context-window manipulation;
- policy bypass through natural-language ambiguity;
- unsafe model or plugin updates;
- data leakage through prompts or outputs;
- cross-case memory contamination;
- insecure model fallback;
- model-service compromise;
- hidden instruction conflicts.

- **TM-REQ-046:** Model output shall not be treated as authorization.
- **TM-REQ-047:** External content shall be treated as untrusted input.
- **TM-REQ-048:** Tool requests shall be constrained by explicit policy and authorization context.
- **TM-REQ-049:** High-risk actions shall require additional control and human review.
- **TM-REQ-050:** Persistent memory shall be scoped, validated, revocable and separated by case and purpose.
- **TM-REQ-051:** Prompt and retrieval pipelines shall preserve source attribution and processing history.
- **TM-REQ-052:** AI-output uncertainty and verification requirements shall be visible to the human reviewer.

## 11. Identity and Delegation Threats

Threat models shall address:

- officer impersonation;
- agent impersonation;
- fraudulent issuance;
- agent rebinding;
- delegation forgery;
- stale delegation;
- token replay;
- token substitution;
- credential theft;
- key misuse;
- unauthorized privilege broadening;
- missing case or purpose context;
- cross-jurisdiction access;
- incomplete revocation;
- compromised policy decision points;
- audit attribution failure.

- **TM-REQ-053:** Identity and delegation threats shall trace to `OBDIA-ID-001`.
- **TM-REQ-054:** Authentication success shall not be treated as authorization success.
- **TM-REQ-055:** Rebinding an agent to a different officer shall be modeled as prohibited.
- **TM-REQ-056:** Revocation latency and failure paths shall be tested.
- **TM-REQ-057:** Compromise of one agent or workload shall be contained from unrelated officers, cases and mandates.

## 12. Evidence and Audit Threats

Threat models shall address:

- source substitution;
- evidence tampering;
- provenance loss;
- unauthorized transformation;
- timestamp manipulation;
- hash mismatch;
- audit deletion or modification;
- chain-of-custody gaps;
- evidence and analysis conflation;
- unauthorized evidence access;
- retention failure;
- evidence contamination after trust failure;
- incomplete reproduction records.

- **TM-REQ-058:** Original source data and derived analysis shall remain distinguishable.
- **TM-REQ-059:** Evidence integrity failures shall trigger quarantine or qualified handling.
- **TM-REQ-060:** Audit systems shall be protected from the components they record where feasible.
- **TM-REQ-061:** Evidence access shall be scoped, attributable and logged.
- **TM-REQ-062:** AI-generated content shall not be represented as original evidence.
- **TM-REQ-063:** Chain-of-custody threats shall trace to evidence and audit controls and validation tests.

## 13. Supply-Chain and Dependency Threats

Threat models shall address:

- compromised packages;
- malicious updates;
- dependency confusion;
- unsigned or unverified artifacts;
- build-pipeline compromise;
- compromised model or dataset supply chain;
- vendor-service compromise;
- unsupported or deprecated dependencies;
- secret leakage in pipelines;
- provenance loss in release artifacts.

- **TM-REQ-064:** Material dependencies shall have provenance, version and integrity records.
- **TM-REQ-065:** Dependency updates shall trigger risk-based reassessment.
- **TM-REQ-066:** Build and release artifacts shall correspond to reviewed source and configuration.
- **TM-REQ-067:** Vendor assurance shall not replace project-specific threat analysis.
- **TM-REQ-068:** Unsupported dependencies shall have a migration, containment or rejection decision.

## 14. Operational Isolation and Research Boundaries

- **TM-REQ-069:** High-risk research shall use segregated, authorized and controlled environments.
- **TM-REQ-070:** Public demonstrations shall use synthetic, fictitious, mock, testnet, lawful public, archived-authorized or controlled-laboratory resources.
- **TM-REQ-071:** Threat validation shall not rely on real criminal marketplaces or uncontrolled criminal infrastructure.
- **TM-REQ-072:** Real-person investigations shall not be used in public demonstrations.
- **TM-REQ-073:** Offensive exploitation, credential theft and malware deployment are prohibited validation methods.
- **TM-REQ-074:** Egress, data transfer, reset and evidence-handling boundaries shall be documented for isolated environments.
- **TM-REQ-075:** A disclaimer shall not substitute for effective technical isolation.

## 15. Required Deliverables

Each material threat model shall include:

- threat-model identifier;
- reviewed artifact and version;
- scope statement;
- architecture or data-flow representation;
- trust-boundary representation where applicable;
- attack-surface inventory;
- asset inventory;
- threat-actor profiles;
- threat and abuse-case register;
- control mapping;
- security requirements;
- validation plan;
- residual-risk summary;
- assumptions and limitations;
- review history;
- approval status.

- **TM-REQ-076:** Deliverables shall be stored in governed repository locations.
- **TM-REQ-077:** Threat identifiers shall be immutable.
- **TM-REQ-078:** Closed threats shall retain disposition and evidence.
- **TM-REQ-079:** Superseded threat models shall remain historically traceable.
- **TM-REQ-080:** Machine-readable registers may supplement but shall not replace human-reviewable analysis.

## 16. Review and Validation Gate

No material implementation progresses until:

- scope is approved;
- applicable architecture is available;
- assets and boundaries are documented;
- threats and abuse cases are identified;
- security requirements are testable;
- controls and owners are identified;
- residual risks are recorded;
- validation criteria exist;
- Security Reviewer findings are resolved or formally dispositioned.

- **TM-REQ-081:** Critical unresolved threats block implementation.
- **TM-REQ-082:** A review shall identify the exact commit or artifact evaluated.
- **TM-REQ-083:** Founder approval is required before this baseline becomes Approved.
- **TM-REQ-084:** Repository merge alone does not establish threat-model approval.
- **TM-REQ-085:** Validation evidence shall identify exact implementation and environment versions.
- **TM-REQ-086:** Failed security tests shall reopen the related threat or control finding.

## 17. Traceability

Every material threat shall trace to:

- affected asset;
- related architecture;
- trust or authorization boundary;
- originating requirement;
- control;
- implementation;
- test;
- validation evidence;
- residual-risk disposition.

- **TM-REQ-087:** Traceability shall be bidirectional.
- **TM-REQ-088:** Broken traceability affecting authorization, evidence, privacy, security or release status is blocking.
- **TM-REQ-089:** Control changes shall trigger review of linked threats and tests.
- **TM-REQ-090:** Threat-model references shall use immutable identifiers where available.

## 18. Validation Checklist

Before approval of this baseline or a derived threat model, confirm:

- [ ] mandatory metadata is complete;
- [ ] scope and reviewed artifacts are explicit;
- [ ] actors, assets, entry points and dependencies are identified;
- [ ] trust and authorization boundaries are identified;
- [ ] threat actors and abuse cases are concrete;
- [ ] mandatory threat categories are addressed or marked non-applicable with rationale;
- [ ] AI, identity, evidence, supply-chain and privacy threats are addressed where relevant;
- [ ] existing and proposed controls are distinguished;
- [ ] control owners and implementation states are recorded;
- [ ] security requirements are testable;
- [ ] residual risks and owners are explicit;
- [ ] validation methods remain defensive, simulated and authorized;
- [ ] diagrams are included where materially useful;
- [ ] traceability is bidirectional;
- [ ] assumptions and limitations are visible;
- [ ] implementation status is accurate;
- [ ] review evidence identifies exact commits or versions;
- [ ] prohibited validation methods are absent;
- [ ] required Security Reviewer and Founder approvals are recorded.

## 19. Limitations

- This baseline does not constitute a complete threat model for the full OBDIA system.
- It does not select a specific risk-scoring framework, modeling tool or diagram notation.
- Threat identification cannot guarantee discovery of every vulnerability or abuse path.
- Risk ratings depend on documented assumptions and available evidence.
- Internal review does not constitute independent penetration testing, certification or operational accreditation.
- Threat analysis does not authorize live exploitation or interaction with unauthorized systems.
- Detailed component threat models remain required for documents 08–30 and future approved implementation artifacts.

## 20. Change Control

Every material change shall include:

- change identifier;
- rationale;
- affected threat categories and documents;
- dependency analysis;
- security impact;
- privacy and fundamental-rights impact;
- implementation impact;
- migration requirements;
- validation plan;
- rollback plan;
- approval record;
- semantic version change.

- **TM-REQ-091:** Editorial corrections use a patch version when meaning is unchanged.
- **TM-REQ-092:** Backward-compatible substantive additions use a minor version.
- **TM-REQ-093:** Incompatible methodology or normative changes use a major version.
- **TM-REQ-094:** Material threat-modeling decisions require an ADR where constitutionally required.
- **TM-REQ-095:** Changes require Project Founder approval.
- **TM-REQ-096:** Unapproved threat-model methodology changes during an active phase are prohibited.

## 21. Consolidation Record

Version 1.1.0 consolidates the existing Threat Model Baseline without expanding project scope. It:

- retains immutable document identifier `OBDIA-TM-001`;
- normalizes the authoritative filename to `07_THREAT_MODEL_BASELINE.md`;
- replaces `Enterprise Baseline (Draft)` with constitutional lifecycle status `Draft`;
- preserves the original protected assets, threat actors, trust boundaries, mandatory threat categories, required analysis, control families, deliverables and implementation gate;
- adds explicit authorization-boundary, AI-agent, identity, evidence, supply-chain, operational-isolation, risk, traceability and validation requirements;
- reconciles the original diagram requirement with the constitutional rule by requiring diagrams where materially applicable and permitting documented non-applicability;
- treats the former `_ENTERPRISE` and non-Enterprise filenames as legacy source variants of the same immutable document rather than separate normative documents;
- creates no new normative document and authorizes no offensive testing or implementation.

## 22. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | Prior baseline | Enterprise Baseline (Draft) | Project origin record | Original Threat Model Baseline covering assets, actors, boundaries, threat categories, required analysis, control families, deliverables and implementation gate. |
| 1.1.0 | 2026-08-05 | Draft | Security Reviewer; approval reserved to Project Founder | Constitutional consolidation: normalized filename, metadata and lifecycle status; preserved original intent; added traceable threat, control, risk, AI, identity, evidence, supply-chain, review, validation and limitation requirements. |
