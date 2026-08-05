# IMPLEMENTATION POLICY

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-IMP-001 |
| **Title** | Implementation Policy |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Implementation Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Govern when and how OBDIA implementation may begin, how it shall conform to approved architecture and decisions, which environments and data are permitted, how secure development and validation evidence are produced, and how deviations, rollback and implementation status are controlled. |
| **Scope** | Source code, configuration, infrastructure-as-code, policies, schemas, connectors, models, prompts, test harnesses, local laboratories, testnets, mock services, build artifacts, deployment artifacts and implementation documentation within the approved OBDIA scope. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `02_PROJECT_GLOSSARY.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `04_PORTFOLIO_AND_PUBLICATION_POLICY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `06_IDENTITY_AND_DELEGATION_MODEL.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; forward dependencies `16_SECURITY_ARCHITECTURE_BASELINE.md`, `17_AUTHORIZATION_MODEL.md`, `18_EVIDENCE_MODEL.md` and `20_CONNECTOR_SECURITY_POLICY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GLOSS-001`; `OBDIA-RES-001`; `OBDIA-PUB-001`; `OBDIA-ARCH-001`; `OBDIA-ID-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001` |
| **Cross-References** | Documents 13–30; accepted ADRs; threat models; authorization, evidence, connector, testing, versioning, risk and compliance artifacts |
| **Assumptions** | Implementation is a controlled research activity and uses synthetic, simulated, testnet, mock, local-laboratory, archived-authorized or otherwise explicitly authorized resources. |
| **Constraints** | Implementation shall not create independent AI legal authority, enable prohibited conduct, use real investigations in public demonstrations, interact with uncontrolled criminal infrastructure or silently expand the approved project scope. |
| **Security Considerations** | Implementation can introduce privilege escalation, credential exposure, prompt injection, unsafe tool use, evidence corruption, privacy harm, supply-chain compromise, cross-case contamination, insecure defaults and misleading capability claims. |
| **Validation Criteria** | Implementation begins only after entry gates are satisfied; changes trace to requirements and ADRs; environments and data are approved; controls fail safely; tests and evidence identify exact versions; deviations are governed; and implementation status is represented accurately. |
| **Implementation Relationship** | This document governs implementation activity but does not itself authorize production deployment, operational investigations, legal access, institutional adoption or public release. |

---

## 1. Purpose

This policy defines the mandatory controls for implementing the Officer-Bound Digital Investigation Agent project.

The original implementation baseline established that implementation may begin only after:

1. architecture approval;
2. ADR acceptance;
3. threat-model review.

It also required the use of:

- synthetic data;
- testnets;
- mock services;
- local laboratories;

and prohibited:

- production claims;
- unauthorized capabilities;
- interaction with criminal infrastructure.

This consolidation preserves those rules and expands them into a complete implementation-governance policy without creating implementation itself.

## 2. Fundamental Implementation Rules

- **IMP-REQ-001:** Implementation shall remain subordinate to the Canonical Project Definition and Master Documentation Constitution.
- **IMP-REQ-002:** Implementation shall not silently define architecture.
- **IMP-REQ-003:** Implementation shall not grant independent legal authority to an AI agent.
- **IMP-REQ-004:** Implementation shall not authorize or enable prohibited conduct.
- **IMP-REQ-005:** Implementation shall remain within the approved Knowledge Pack 01–30 scope.
- **IMP-REQ-006:** An implemented capability shall not be represented as approved, tested, validated, secure, compliant, production-ready or operational unless corresponding evidence exists.
- **IMP-REQ-007:** Human and institutional accountability shall remain explicit in every material implementation path.
- **IMP-REQ-008:** Model output shall not be treated as authorization.
- **IMP-REQ-009:** Authentication shall not be treated as authorization.
- **IMP-REQ-010:** Authorization shall not be treated as legal authority.
- **IMP-REQ-011:** Missing or invalid authorization context shall fail closed.
- **IMP-REQ-012:** No agent, model or workload may self-approve a privilege increase, exception, deployment or release.

## 3. Implementation Entry Gate

Material implementation may begin only when all applicable prerequisites are satisfied:

1. approved scope exists;
2. relevant architecture is approved or formally authorized for a bounded experiment;
3. required ADRs are Accepted;
4. applicable threat model is reviewed;
5. trust and authorization boundaries are defined;
6. security and privacy requirements are testable;
7. acceptance criteria exist;
8. implementation owner is assigned;
9. data and environment classification is known;
10. rollback or containment requirements are defined.

- **IMP-REQ-013:** Architecture approval is mandatory before implementation of a material feature.
- **IMP-REQ-014:** Required ADRs shall be Accepted before implementation relies on their decisions.
- **IMP-REQ-015:** Threat modeling shall be completed before implementation of a material feature.
- **IMP-REQ-016:** Critical unresolved threat findings block implementation.
- **IMP-REQ-017:** Authorization boundaries shall exist before implementation begins.
- **IMP-REQ-018:** Entry-gate evidence shall identify the exact requirements, architecture, ADRs and threat model.
- **IMP-REQ-019:** Gate non-applicability requires documented rationale.
- **IMP-REQ-020:** A repository branch or issue shall not be treated as implementation authorization.
- **IMP-REQ-021:** A local experiment may proceed only when isolated, reversible, non-operational and explicitly labeled.
- **IMP-REQ-022:** Experimental implementation shall not establish normative architecture by implication.
- **IMP-REQ-023:** Material changes after gate approval require affected gate re-evaluation.
- **IMP-REQ-024:** Forward dependencies 16, 17, 18 and 20 block approval of this policy until their relevant requirements are consolidated and reconciled.

## 4. Permitted Implementation Environments

Permitted environments include:

- local development workstations using controlled test data;
- isolated local laboratories;
- containerized or virtualized test environments;
- mock APIs and services;
- synthetic identity providers and policy engines;
- blockchain testnets;
- authorized CTF or defensive training environments;
- controlled cloud sandboxes;
- archived-authorized datasets;
- read-only lawful public-source demonstrations;
- operationally isolated Dark Web research simulations;
- other explicitly authorized controlled environments.

- **IMP-REQ-025:** Every environment shall have a documented purpose, owner, classification and access boundary.
- **IMP-REQ-026:** Development and test environments shall be separated from operational or evidentiary environments.
- **IMP-REQ-027:** Public demonstrations shall be separated from private and operational data.
- **IMP-REQ-028:** High-risk research shall use appropriate network, egress, identity and data isolation.
- **IMP-REQ-029:** Dark Web research implementations shall remain authorized, simulated or operationally isolated and shall not interact with uncontrolled criminal infrastructure.
- **IMP-REQ-030:** Testnets shall not be represented as production networks.
- **IMP-REQ-031:** Mock services shall not be represented as live institutional integrations.
- **IMP-REQ-032:** Environment reset and destruction procedures shall be documented.
- **IMP-REQ-033:** Environment compromise or integrity failure shall trigger containment and revalidation.
- **IMP-REQ-034:** Environment promotion shall require a distinct review decision.
- **IMP-REQ-035:** Network location shall not establish trust by itself.
- **IMP-REQ-036:** Shared environments shall preserve case, purpose and identity separation.

## 5. Prohibited Implementation Contexts

Implementation shall not depend on:

- unauthorized systems or accounts;
- stolen credentials or tokens;
- unlawfully obtained datasets;
- real criminal marketplaces;
- uncontrolled criminal infrastructure;
- active malware deployment;
- offensive exploitation;
- unlawful surveillance;
- uncontrolled deanonymization;
- real-person public investigations;
- production officer identities in public demonstrations;
- production secrets in source code or documentation;
- unapproved institutional systems;
- undisclosed high-risk external services.

- **IMP-REQ-037:** Prohibited contexts shall remain non-authorizable within the project.
- **IMP-REQ-038:** A disclaimer shall not compensate for prohibited implementation.
- **IMP-REQ-039:** Dual-use functionality shall be redesigned into defensive, read-only, simulated or controlled-laboratory operation.
- **IMP-REQ-040:** Implementation shall not include hidden offensive capability.
- **IMP-REQ-041:** Implementation shall not provide instructions designed to bypass authorization or monitoring controls.
- **IMP-REQ-042:** Production claims are prohibited until production status is independently established by authorized governance and evidence outside this research baseline.

## 6. Data Requirements

Implementation data shall be one or more of:

- synthetic;
- fictitious;
- generated test fixtures;
- anonymized or de-identified under an approved method;
- testnet data;
- mock-service data;
- lawful public data with documented purpose and restrictions;
- archived-authorized data;
- institutionally authorized controlled data.

- **IMP-REQ-043:** Data origin, purpose, classification and authorization shall be recorded.
- **IMP-REQ-044:** Public demonstrations shall use synthetic or otherwise explicitly approved safe data.
- **IMP-REQ-045:** Real case data shall not be used in public demonstrations.
- **IMP-REQ-046:** Real personal data shall be minimized and shall not be used without explicit authorization and governance.
- **IMP-REQ-047:** Test fixtures shall avoid realistic secrets, credentials and sensitive identifiers.
- **IMP-REQ-048:** Cross-case and cross-purpose data reuse requires a separate authorization decision.
- **IMP-REQ-049:** Data retention and deletion behavior shall be testable.
- **IMP-REQ-050:** Source evidence and derived analysis shall remain distinguishable.
- **IMP-REQ-051:** AI-generated content shall not be represented as original evidence.
- **IMP-REQ-052:** Data transformations shall preserve provenance and reproducibility.
- **IMP-REQ-053:** Suspicious or malformed data shall be quarantined or processed through an explicitly untrusted path.
- **IMP-REQ-054:** Test data shall not silently become production or evidentiary data.

## 7. Identity, Credential and Secrets Controls

- **IMP-REQ-055:** Human credentials shall not be embedded in an agent or workload.
- **IMP-REQ-056:** Material workloads shall use distinct machine identities where feasible.
- **IMP-REQ-057:** Credentials shall be scoped, revocable and time-limited where supported.
- **IMP-REQ-058:** Secrets shall not be committed to source control.
- **IMP-REQ-059:** Secrets shall not appear in prompts, logs, screenshots, fixtures or documentation.
- **IMP-REQ-060:** Repository history shall be considered during secret review.
- **IMP-REQ-061:** Development credentials shall be non-production and independently revocable.
- **IMP-REQ-062:** Shared static credentials require an approved exception, owner, expiry and compensating controls.
- **IMP-REQ-063:** Credential rotation and revocation shall be testable.
- **IMP-REQ-064:** Missing, expired, suspended or revoked credentials shall fail closed.
- **IMP-REQ-065:** Credential identifiers may be logged, but secret material shall not.
- **IMP-REQ-066:** A compromised credential shall be contained from unrelated cases, officers and workloads.
- **IMP-REQ-067:** Secret-management implementation shall follow an Accepted ADR when the choice is architecturally material.
- **IMP-REQ-068:** Example configuration shall use unmistakably non-secret placeholders.

## 8. Authorization and Policy Enforcement

Every privileged operation shall evaluate, where applicable:

- officer identity;
- agent identity;
- institution;
- case or mandate;
- purpose;
- jurisdiction;
- resource;
- requested action;
- tool or connector;
- credential status;
- policy version;
- time;
- risk state;
- human-approval requirement;
- revocation or containment state.

- **IMP-REQ-069:** Authorization shall be enforced at the point of use.
- **IMP-REQ-070:** Authorization shall default to deny.
- **IMP-REQ-071:** Context loss shall result in denial rather than broader fallback.
- **IMP-REQ-072:** Cached authorization shall have defined expiry and invalidation.
- **IMP-REQ-073:** Cross-case or cross-purpose operations require distinct authorization.
- **IMP-REQ-074:** Prohibited actions shall not be expressible as permitted policy outcomes.
- **IMP-REQ-075:** Policy decisions shall identify the applicable policy version.
- **IMP-REQ-076:** Denied operations shall not be retried through a less restrictive path.
- **IMP-REQ-077:** Human approval shall be explicit and attributable.
- **IMP-REQ-078:** Silence, timeout or service failure shall not be interpreted as approval.
- **IMP-REQ-079:** Policy conflicts shall fail closed and generate reviewable evidence.
- **IMP-REQ-080:** Policy engines and enforcement points shall use distinct responsibilities where feasible.
- **IMP-REQ-081:** Authorization tests shall include positive, negative, boundary and stale-context cases.
- **IMP-REQ-082:** Authorization implementation shall trace to `OBDIA-ID-001`, `OBDIA-TRUST-001` and the future consolidated `OBDIA-AUTH-001`.

## 9. Model, Prompt, Retrieval, Memory and Tool Controls

- **IMP-REQ-083:** External content shall be treated as untrusted input.
- **IMP-REQ-084:** Instructions embedded in retrieved content shall not override system, constitutional, policy or operator authority.
- **IMP-REQ-085:** Model output shall not directly invoke a privileged tool without separate policy enforcement.
- **IMP-REQ-086:** Tool requests shall be typed or otherwise constrained where feasible.
- **IMP-REQ-087:** Tools shall receive only the minimum data and permissions required.
- **IMP-REQ-088:** Tool outputs shall be validated and retain provenance.
- **IMP-REQ-089:** High-impact tool use shall require additional control and human review.
- **IMP-REQ-090:** Persistent memory shall be scoped by officer, case, purpose and retention rules.
- **IMP-REQ-091:** Memory content shall not become trusted solely because it was stored previously.
- **IMP-REQ-092:** Cross-case memory access shall be denied by default.
- **IMP-REQ-093:** Prompt, model, retrieval and tool versions relevant to validation shall be recorded.
- **IMP-REQ-094:** Model-provider or fallback changes shall trigger trust and risk reassessment.
- **IMP-REQ-095:** Model uncertainty and verification requirements shall be visible to human reviewers.
- **IMP-REQ-096:** Unsafe model or tool behavior shall support containment and disablement.
- **IMP-REQ-097:** Prompt-injection tests shall use synthetic or authorized defensive fixtures.
- **IMP-REQ-098:** Implementation shall not rely solely on natural-language instructions for a critical authorization boundary.

## 10. Evidence, Audit and Chain of Custody

- **IMP-REQ-099:** Evidence acquisition and transformation paths shall preserve provenance.
- **IMP-REQ-100:** Original data and derived artifacts shall remain distinguishable.
- **IMP-REQ-101:** Material transformations shall be reproducible.
- **IMP-REQ-102:** Evidence access shall be authenticated, authorized and logged.
- **IMP-REQ-103:** Chain-of-custody events shall identify actor, time, action and artifact.
- **IMP-REQ-104:** Integrity checks shall use approved mechanisms and record results.
- **IMP-REQ-105:** Evidence-integrity failure shall trigger quarantine or qualified handling.
- **IMP-REQ-106:** Audit records shall distinguish human, agent, workload, connector and model-related actors where applicable.
- **IMP-REQ-107:** Audit records shall not expose secrets or unnecessary personal data.
- **IMP-REQ-108:** Audit history shall be tamper-evident where technically feasible.
- **IMP-REQ-109:** Audit-system failure shall not permit silent privileged operation.
- **IMP-REQ-110:** Clock, ordering and timestamp limitations shall be documented.
- **IMP-REQ-111:** Evidence and audit implementation shall trace to the future consolidated `OBDIA-EVID-001`.
- **IMP-REQ-112:** Implementation tests shall verify that revocation does not erase required evidence or audit history.

## 11. Secure Development Requirements

- **IMP-REQ-113:** Source changes shall be attributable through version control.
- **IMP-REQ-114:** Material changes shall use reviewed branches or equivalent controlled change paths.
- **IMP-REQ-115:** Code review shall evaluate security, privacy, authorization, evidence and failure behavior where applicable.
- **IMP-REQ-116:** Secure defaults shall be used.
- **IMP-REQ-117:** Dangerous optional behavior shall not be enabled by default.
- **IMP-REQ-118:** Input validation shall be appropriate to the interface and threat model.
- **IMP-REQ-119:** Output encoding and data handling shall prevent unintended execution or disclosure where relevant.
- **IMP-REQ-120:** Error messages shall not expose secrets or unnecessary sensitive context.
- **IMP-REQ-121:** Logging shall be sufficient for accountability without excessive data collection.
- **IMP-REQ-122:** Security-relevant configuration shall be versioned and reviewable.
- **IMP-REQ-123:** Configuration drift shall be detectable where feasible.
- **IMP-REQ-124:** Development shortcuts shall not silently become release configuration.
- **IMP-REQ-125:** Unsafe debug interfaces shall be disabled outside their approved environment.
- **IMP-REQ-126:** Generated code and AI-assisted changes require human review.
- **IMP-REQ-127:** A coding standard shall govern implementation when document 21 is consolidated.
- **IMP-REQ-128:** A testing standard shall govern validation when document 22 is consolidated.

## 12. Dependency and Supply-Chain Controls

- **IMP-REQ-129:** Material dependencies shall have version and provenance records.
- **IMP-REQ-130:** Dependencies shall be evaluated for maintenance, licensing, security and applicability.
- **IMP-REQ-131:** Unsupported or deprecated dependencies require migration, containment or rejection.
- **IMP-REQ-132:** Dependency updates shall trigger risk-based reassessment.
- **IMP-REQ-133:** Build artifacts shall correspond to reviewed source and configuration.
- **IMP-REQ-134:** Artifact integrity shall be verifiable where feasible.
- **IMP-REQ-135:** Build and release pipelines shall use least privilege.
- **IMP-REQ-136:** Pipeline secrets shall be scoped and excluded from logs.
- **IMP-REQ-137:** Third-party services shall not receive data beyond approved purpose and necessity.
- **IMP-REQ-138:** Vendor claims shall not replace project-specific validation.
- **IMP-REQ-139:** Model, dataset and prompt-template provenance shall be recorded when material.
- **IMP-REQ-140:** Supply-chain compromise assumptions shall be included in threat modeling.
- **IMP-REQ-141:** Lockfiles, manifests or equivalent dependency records should be retained where applicable.
- **IMP-REQ-142:** Dependency exceptions require an owner, rationale, expiry and residual-risk record.

## 13. Implementation Status Model

Approved feature-level implementation states are:

- **Conceptual**
- **Designed**
- **Partially Implemented**
- **Implemented**
- **Tested**
- **Simulated**
- **Deferred**
- **Out of Scope**

- **IMP-REQ-143:** Status shall be assigned at the feature or control level when aggregate status would be misleading.
- **IMP-REQ-144:** `Implemented` shall not mean `Tested`, `Validated`, `Secure`, `Compliant` or `Production Ready`.
- **IMP-REQ-145:** `Tested` shall identify test scope, environment, version and result.
- **IMP-REQ-146:** `Simulated` shall identify the mock, testnet, synthetic or laboratory boundary.
- **IMP-REQ-147:** `Partially Implemented` shall identify missing required components.
- **IMP-REQ-148:** `Deferred` shall identify rationale, owner and dependency or review point.
- **IMP-REQ-149:** `Out of Scope` shall identify the controlling scope decision.
- **IMP-REQ-150:** Status changes shall be supported by repository evidence.
- **IMP-REQ-151:** Production claims remain prohibited unless a future authorized institutional governance process establishes them.
- **IMP-REQ-152:** Screenshots and demonstrations shall not imply a higher status than the underlying evidence.

## 14. Testing and Acceptance Evidence

Implementation acceptance evidence shall include, where applicable:

- requirement identifiers;
- ADR identifiers;
- threat and control identifiers;
- implementation commit;
- configuration and policy version;
- dependency versions;
- model and prompt version;
- environment;
- test data classification;
- test procedure;
- expected result;
- observed result;
- failures;
- limitations;
- reviewer;
- date.

- **IMP-REQ-153:** Tests shall trace to requirements, threats, controls and ADRs.
- **IMP-REQ-154:** Negative and failure-path tests are mandatory for material authorization and security controls.
- **IMP-REQ-155:** Test evidence shall identify exact versions.
- **IMP-REQ-156:** Failed tests shall remain visible and block acceptance when material.
- **IMP-REQ-157:** Passing tests shall not be generalized beyond their scope.
- **IMP-REQ-158:** Test fixtures shall use synthetic or authorized controlled data.
- **IMP-REQ-159:** Reproducibility limitations shall be documented.
- **IMP-REQ-160:** Acceptance criteria shall be defined before final implementation review.
- **IMP-REQ-161:** Acceptance shall be attributable to an authorized human reviewer.
- **IMP-REQ-162:** An AI system shall not serve as acceptance authority.
- **IMP-REQ-163:** Automated checks may provide evidence but shall not replace required human review.
- **IMP-REQ-164:** Validation evidence shall be retained in governed repository locations.

## 15. Deviations and Non-Conformance

A deviation record shall include:

- identifier;
- affected requirement or ADR;
- rationale;
- scope;
- owner;
- security and privacy impact;
- evidence impact;
- compensating controls;
- duration;
- remediation plan;
- validation plan;
- approval.

- **IMP-REQ-165:** Implementation conflicting with approved architecture or ADRs is non-conforming.
- **IMP-REQ-166:** Non-conformance shall not be hidden through documentation changes made after implementation.
- **IMP-REQ-167:** A deviation shall be narrow, explicit and time-bound.
- **IMP-REQ-168:** No deviation may authorize a canonical prohibition.
- **IMP-REQ-169:** Repeated deviations indicating structural weakness require architectural review.
- **IMP-REQ-170:** Expired deviations shall result in denial, rollback or renewed formal review.
- **IMP-REQ-171:** Critical non-conformance blocks release.
- **IMP-REQ-172:** Remediation shall preserve evidence and audit history.
- **IMP-REQ-173:** A permanent architecture change requires an ADR and normative update where applicable.
- **IMP-REQ-174:** Deviation approval shall not be represented as baseline conformance.

## 16. Rollback, Revocation and Emergency Containment

- **IMP-REQ-175:** Material implementation shall define rollback or containment behavior.
- **IMP-REQ-176:** Rollback shall restore a known governed state.
- **IMP-REQ-177:** Rollback shall not restore revoked credentials or broader authority.
- **IMP-REQ-178:** Emergency containment shall support disabling affected identities, sessions, tools, connectors and workloads.
- **IMP-REQ-179:** Containment shall preserve required evidence and audit records.
- **IMP-REQ-180:** Failed rollback or containment shall generate an incident and unresolved finding.
- **IMP-REQ-181:** Recovery shall require revalidation of identity, authorization, configuration and integrity.
- **IMP-REQ-182:** Emergency changes shall be narrow, attributable, time-bounded and retrospectively reviewed.
- **IMP-REQ-183:** Emergency procedures shall not authorize prohibited conduct.
- **IMP-REQ-184:** Rollback and containment shall be tested in controlled environments.

## 17. Implementation Review Gate

Before a material implementation is accepted, the Implementation Reviewer shall verify:

- approved scope;
- applicable architecture;
- Accepted ADRs;
- reviewed threat model;
- defined trust and authorization boundaries;
- secure environment and data;
- requirement traceability;
- dependency and secrets controls;
- test results;
- residual risks;
- deviations;
- rollback and containment;
- status accuracy;
- documentation completeness.

- **IMP-REQ-185:** Review evidence shall identify the exact implementation commit.
- **IMP-REQ-186:** Review shall identify code, configuration, policy, model, dependency and environment versions where applicable.
- **IMP-REQ-187:** Material post-review changes invalidate affected evidence.
- **IMP-REQ-188:** Blocking findings prevent acceptance.
- **IMP-REQ-189:** Role concentration shall be disclosed.
- **IMP-REQ-190:** Internal review shall not be represented as independent external assurance.
- **IMP-REQ-191:** Repository merge shall not establish implementation validation automatically.
- **IMP-REQ-192:** Public release requires the separate publication and release gate.
- **IMP-REQ-193:** Operational deployment requires governance beyond this independent research baseline.
- **IMP-REQ-194:** Implementation acceptance shall not alter normative document status by implication.

## 18. Traceability

Every material implementation component shall trace to:

- canonical or normative requirement;
- architecture;
- ADR;
- threat;
- control;
- data and environment classification;
- source code or configuration;
- test;
- validation evidence;
- risk or deviation;
- release artifact where applicable.

- **IMP-REQ-195:** Traceability shall be bidirectional.
- **IMP-REQ-196:** Broken traceability affecting authority, security, privacy, evidence or release is blocking.
- **IMP-REQ-197:** Planned implementation links shall not be represented as implemented or tested.
- **IMP-REQ-198:** Superseded implementation shall identify its replacement or removal status.
- **IMP-REQ-199:** Traceability records shall not contain secrets or sensitive case identifiers.
- **IMP-REQ-200:** Changes to requirements, ADRs or threats shall trigger review of linked implementation and tests.

## 19. Minimum Validation Checklist

Before implementation acceptance, confirm:

- [ ] entry gates are satisfied;
- [ ] architecture and required ADRs are approved;
- [ ] threat model is reviewed;
- [ ] trust and authorization boundaries are enforced;
- [ ] the agent has no independent legal authority;
- [ ] environment and data are authorized and controlled;
- [ ] no prohibited context or capability is present;
- [ ] secrets and credentials are protected;
- [ ] permissions are least-privileged and revocable;
- [ ] prompt, retrieval, memory and tool controls are present;
- [ ] evidence provenance and auditability are preserved;
- [ ] dependencies and build artifacts are traceable;
- [ ] implementation status is accurate;
- [ ] tests include negative and failure paths;
- [ ] exact versions and environment are recorded;
- [ ] deviations and residual risks have owners;
- [ ] rollback and containment are tested;
- [ ] documentation is complete;
- [ ] role concentration and limitations are disclosed;
- [ ] publication and operational claims remain within evidence.

## 20. Limitations

- This policy does not authorize implementation by itself.
- It does not select a programming language, framework, cloud provider, identity platform, policy engine, model provider or deployment product.
- It does not establish production readiness, institutional adoption, operational accreditation or legal compliance.
- It does not replace documents 16, 17, 18 and 20, which remain forward dependencies for full approval.
- A secure design and passing tests cannot guarantee absence of defects or misuse.
- Internal implementation review is not independent certification.
- Public proof-of-concept behavior shall not be generalized to real investigative environments.
- Detailed coding, testing, repository and versioning requirements remain governed by documents 21–25 after consolidation.

## 21. Change Control

Every material change to this policy shall include:

- change identifier;
- rationale;
- affected implementation domains and documents;
- dependency analysis;
- security impact;
- privacy and fundamental-rights impact;
- evidence and audit impact;
- migration requirements;
- validation plan;
- rollback plan;
- approval record;
- semantic version change.

- **IMP-REQ-201:** Editorial corrections use a patch version when meaning is unchanged.
- **IMP-REQ-202:** Backward-compatible substantive additions use a minor version.
- **IMP-REQ-203:** Incompatible normative or implementation-method changes use a major version.
- **IMP-REQ-204:** Material implementation architecture decisions require an ADR.
- **IMP-REQ-205:** Changes require Project Founder approval.
- **IMP-REQ-206:** Unapproved implementation-methodology changes during an active phase are prohibited.
- **IMP-REQ-207:** Changes shall not retroactively fabricate conformance or validation evidence.
- **IMP-REQ-208:** Migration shall update affected tests, documentation, risks and traceability.

## 22. Consolidation Record

Version 1.0.0 consolidates the existing Implementation Policy without expanding project scope. It:

- retains immutable document identifier `OBDIA-IMP-001`;
- normalizes the authoritative filename to `12_IMPLEMENTATION_POLICY.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves the original three prerequisites: approved architecture, Accepted ADR and reviewed threat model;
- preserves the original safe-environment requirements for synthetic data, testnets, mock services and local laboratories;
- preserves the prohibition on production claims;
- preserves the prohibition on unauthorized capabilities and criminal-infrastructure interaction;
- adds entry gates, environment, data, identity, secrets, authorization, AI, evidence, secure-development, dependency, status, testing, deviation, rollback, review and traceability requirements;
- identifies documents 16, 17, 18 and 20 as forward dependencies that block approval until reconciled;
- treats legacy Enterprise and non-Enterprise source variants as the same immutable document;
- creates no implementation and authorizes no operational use.

## 23. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Implementation Reviewer; approval reserved to Project Founder | Constitutional consolidation of the existing minimal policy: preserved original prerequisites, safe environments and prohibitions; added complete implementation governance, 208 stable requirements, validation criteria and explicit forward-dependency limitations. |
