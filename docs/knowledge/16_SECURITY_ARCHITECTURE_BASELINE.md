# SECURITY ARCHITECTURE BASELINE

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-SEC-001 |
| **Title** | Security Architecture Baseline |
| **Version** | 1.0.1 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Security Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Establish mandatory security architecture objectives, domains, boundaries, controls, ownership, failure behavior, containment, validation and traceability for OBDIA. |
| **Scope** | Governance, identity, delegation, authorization, model and runtime, connectors and tools, storage, network, evidence, audit, monitoring, supply chain, implementation, testing, release and controlled research environments within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `06_IDENTITY_AND_DELEGATION_MODEL.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `15_RELEASE_AND_PUBLICATION_POLICY.md`; forward dependencies `17_AUTHORIZATION_MODEL.md`, `18_EVIDENCE_MODEL.md`, `19_AGENT_LIFECYCLE_MODEL.md`, `20_CONNECTOR_SECURITY_POLICY.md`, `22_TESTING_STANDARD.md`, `26_AI_RISK_REGISTER.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-ARCH-001`; `OBDIA-ID-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-REL-001` |
| **Cross-References** | Documents 17–30; accepted ADRs; threat models; authorization policies; evidence records; risk and exception records; implementation, validation, incident and release evidence |
| **Assumptions** | The project is an independent defensive research architecture using synthetic, simulated, testnet, mock, local-laboratory, lawful public, archived-authorized or otherwise explicitly authorized resources. |
| **Constraints** | Security controls shall not create independent AI legal authority, authorize prohibited conduct, transfer human credentials to an agent, permit uncontrolled criminal-infrastructure interaction, or substitute technical capability for lawful authorization. |
| **Security Considerations** | Failure can cause unauthorized access, privilege escalation, cross-case contamination, prompt injection, unsafe tool use, credential compromise, evidence corruption, privacy harm, supply-chain compromise, monitoring failure, misleading assurance and loss of human accountability. |
| **Validation Criteria** | Every security domain has identified assets, owners, trust boundaries, threats, requirements, controls, secure-failure behavior, containment paths, testable criteria, residual-risk decisions and bidirectional traceability to implementation and evidence. |
| **Implementation Relationship** | This baseline defines security requirements but does not itself implement, validate, certify, deploy or authorize an operational system. |

---

## 1. Purpose

This document defines the mandatory security architecture baseline for the Officer-Bound Digital Investigation Agent project.

The two legacy sources established the following security domains:

- Identity;
- Runtime;
- Connectors;
- Storage;
- Network;
- Evidence;
- Governance;
- Monitoring.

They also established these mandatory controls:

- Zero Trust;
- Defense in Depth;
- Secure by Design;
- Privacy by Design;
- Least Privilege;
- Explicit authorization;
- Cryptographic identity;
- immutable or tamper-evident auditability;
- threat modeling before implementation;
- secure defaults;
- continuous verification.

This consolidation preserves all of those controls, resolves the conflicting legacy status labels and turns named principles into enforceable architecture requirements.

The unreviewed legacy label `Approved Baseline` is not accepted as valid approval evidence. The consolidated artifact remains `Draft` until all applicable constitutional review gates are completed.


## 2. Fundamental Security Rules

- **SEC-REQ-001:** Security architecture shall remain subordinate to the Canonical Project Definition and Master Documentation Constitution.
- **SEC-REQ-002:** Human and institutional accountability shall remain explicit throughout every security control path.
- **SEC-REQ-003:** An AI agent shall not possess or exercise independent legal authority.
- **SEC-REQ-004:** Security architecture shall not authorize prohibited conduct.
- **SEC-REQ-005:** Technical capability shall not be treated as lawful authority.
- **SEC-REQ-006:** Authentication shall not be treated as authorization.
- **SEC-REQ-007:** Authorization shall not be treated as legal authority.
- **SEC-REQ-008:** Model output shall not be treated as authorization.
- **SEC-REQ-009:** Trust shall not be assumed from network location, component identity or prior success.
- **SEC-REQ-010:** Missing, stale, conflicting or unverifiable security context shall result in denial or containment.
- **SEC-REQ-011:** Security controls shall fail closed for privileged or evidentiary operations.
- **SEC-REQ-012:** Security architecture shall use defense in depth.
- **SEC-REQ-013:** Security architecture shall use least privilege.
- **SEC-REQ-014:** Security architecture shall use secure defaults.
- **SEC-REQ-015:** Security architecture shall minimize data and privilege exposure.
- **SEC-REQ-016:** Security architecture shall preserve evidence integrity and auditability.
- **SEC-REQ-017:** Security controls shall be reviewable, testable and attributable.
- **SEC-REQ-018:** Security requirements shall trace to threats, controls, implementation, tests and evidence.
- **SEC-REQ-019:** Repository merge shall not establish security approval.
- **SEC-REQ-020:** Automated checks shall not replace required human security review.

## 3. Security Objectives

- **SEC-REQ-021:** Security objectives shall include confidentiality, integrity, availability, accountability, authenticity, authorization correctness, privacy, safety, resilience and evidence integrity.
- **SEC-REQ-022:** Security objectives shall identify protected persons, assets, rights and institutional obligations.
- **SEC-REQ-023:** Security objectives shall be scoped by case, purpose, jurisdiction, time, tool, data source and authorization context where applicable.
- **SEC-REQ-024:** Security objectives shall identify unacceptable outcomes.
- **SEC-REQ-025:** Security objectives shall distinguish prevention, detection, containment, recovery and assurance.
- **SEC-REQ-026:** Security objectives shall identify assumptions and dependencies.
- **SEC-REQ-027:** Security objectives shall identify measurable validation criteria.
- **SEC-REQ-028:** Security objectives shall identify control owners.
- **SEC-REQ-029:** Security objectives shall identify residual-risk owners.
- **SEC-REQ-030:** Security objectives shall not be reduced to technical availability alone.
- **SEC-REQ-031:** Privacy and fundamental-rights harms shall be treated as security-relevant outcomes.
- **SEC-REQ-032:** Evidence corruption or loss of provenance shall be treated as a security failure.
- **SEC-REQ-033:** Loss of officer-agent binding shall be treated as a critical security failure.
- **SEC-REQ-034:** Loss of revocation capability shall be treated as a critical security failure.
- **SEC-REQ-035:** Security objectives shall be reassessed when architecture, threats, data, models, dependencies or intended use materially change.

## 4. Security Domains and Ownership

- **SEC-REQ-036:** Every material security domain shall have an identified owner.
- **SEC-REQ-037:** Governance security shall control authority, approvals, exceptions, risk acceptance and change.
- **SEC-REQ-038:** Identity security shall control human, institutional, agent and workload identities.
- **SEC-REQ-039:** Authorization security shall control policy decisions, enforcement and obligations.
- **SEC-REQ-040:** Runtime security shall control models, prompts, retrieval, memory, tools and execution.
- **SEC-REQ-041:** Connector security shall control external-system and API interactions.
- **SEC-REQ-042:** Storage security shall control data, configuration, secrets, logs and evidence at rest.
- **SEC-REQ-043:** Network security shall control communication paths, segmentation, ingress and egress.
- **SEC-REQ-044:** Evidence security shall control provenance, integrity, chain of custody and transformation.
- **SEC-REQ-045:** Monitoring security shall control telemetry, detection, alerting and review.
- **SEC-REQ-046:** Supply-chain security shall control source, dependencies, builds, models, datasets and artifacts.
- **SEC-REQ-047:** Release security shall control distributed content, manifests, integrity and withdrawal.
- **SEC-REQ-048:** Domain boundaries shall be explicit.
- **SEC-REQ-049:** Cross-domain dependencies shall be documented.
- **SEC-REQ-050:** No security domain shall rely on undocumented implicit trust.
- **SEC-REQ-051:** Domain controls shall identify failure and degradation behavior.
- **SEC-REQ-052:** Domain controls shall identify containment scope.
- **SEC-REQ-053:** Domain controls shall identify required evidence.
- **SEC-REQ-054:** Domain ownership shall remain assigned to authorized human roles.
- **SEC-REQ-055:** AI systems may support analysis but shall not own or approve a security domain.

## 5. Trust and Authorization Boundaries

- **SEC-REQ-056:** Trust boundaries shall be identified before implementation.
- **SEC-REQ-057:** Authorization boundaries shall be distinguished from network boundaries.
- **SEC-REQ-058:** Case boundaries shall be explicit.
- **SEC-REQ-059:** Purpose boundaries shall be explicit.
- **SEC-REQ-060:** Jurisdiction boundaries shall be explicit where applicable.
- **SEC-REQ-061:** Human-to-agent boundaries shall preserve attribution.
- **SEC-REQ-062:** Agent-to-workload boundaries shall use distinct identities where feasible.
- **SEC-REQ-063:** Workload-to-connector boundaries shall enforce least privilege.
- **SEC-REQ-064:** Model-to-tool boundaries shall require separate policy enforcement.
- **SEC-REQ-065:** External-content boundaries shall treat content as untrusted.
- **SEC-REQ-066:** Evidence-to-analysis boundaries shall preserve original and derived distinctions.
- **SEC-REQ-067:** Development, test, demonstration and operational boundaries shall remain separated.
- **SEC-REQ-068:** Public and private repository boundaries shall be documented.
- **SEC-REQ-069:** Boundary crossings shall be authenticated, authorized, logged and constrained.
- **SEC-REQ-070:** Boundary failure shall not broaden access.
- **SEC-REQ-071:** Boundary diagrams shall use governed terminology.
- **SEC-REQ-072:** Boundary changes shall trigger threat-model review.
- **SEC-REQ-073:** Boundary changes with material architecture impact shall require an ADR.
- **SEC-REQ-074:** Boundary assumptions shall be validated rather than inferred.
- **SEC-REQ-075:** Unresolved boundary ambiguity shall block implementation or release.

## 6. Zero Trust and Continuous Verification

- **SEC-REQ-076:** Zero Trust shall apply to every human, agent, workload, connector, model, source, network and repository interaction.
- **SEC-REQ-077:** Access shall be explicitly evaluated for each protected operation.
- **SEC-REQ-078:** Continuous verification shall be proportionate to risk and operation duration.
- **SEC-REQ-079:** Prior authentication shall not create permanent trust.
- **SEC-REQ-080:** Prior authorization shall not create permanent authorization.
- **SEC-REQ-081:** Cached decisions shall have bounded lifetime and invalidation rules.
- **SEC-REQ-082:** Credential status shall be re-evaluated when material risk changes.
- **SEC-REQ-083:** Policy version shall be part of material authorization evidence.
- **SEC-REQ-084:** Device, runtime and workload posture shall be considered where applicable.
- **SEC-REQ-085:** Source reputation shall not replace content validation.
- **SEC-REQ-086:** Internal components shall not be trusted solely because they are internal.
- **SEC-REQ-087:** Encrypted communication shall not be treated as authorization.
- **SEC-REQ-088:** Successful model execution shall not establish trust in output.
- **SEC-REQ-089:** Successful tool execution shall not establish trust in returned data.
- **SEC-REQ-090:** Trust elevation shall require explicit evidence and policy.
- **SEC-REQ-091:** Trust degradation shall trigger step-up, restriction, quarantine or denial.
- **SEC-REQ-092:** Silence or timeout shall not be interpreted as Permit.
- **SEC-REQ-093:** Zero Trust decisions shall be explainable to an authorized reviewer.
- **SEC-REQ-094:** Zero Trust telemetry shall avoid unnecessary personal data.
- **SEC-REQ-095:** Zero Trust implementation shall trace to `OBDIA-TRUST-001` and future `OBDIA-AUTH-001`.

## 7. Identity and Cryptographic Binding

- **SEC-REQ-096:** Every agent identity shall be institutionally issued.
- **SEC-REQ-097:** Every agent identity shall be bound to exactly one identified human officer at a time.
- **SEC-REQ-098:** Every agent identity shall be bound to one accountable institution.
- **SEC-REQ-099:** Human credentials shall not be transferred to or embedded in an agent.
- **SEC-REQ-100:** Agent and workload identities shall be cryptographically distinguishable from human identities.
- **SEC-REQ-101:** Shared machine identities shall be avoided.
- **SEC-REQ-102:** Identity credentials shall be scoped, time-bounded and revocable where supported.
- **SEC-REQ-103:** Identity proofing strength shall be proportionate to risk.
- **SEC-REQ-104:** Credential issuance shall be attributable.
- **SEC-REQ-105:** Credential rotation shall be supported.
- **SEC-REQ-106:** Credential revocation shall propagate to dependent sessions and permissions.
- **SEC-REQ-107:** Rebinding an existing agent identity to another officer shall be prohibited.
- **SEC-REQ-108:** Compromised identities shall be containable without disabling unrelated officers or cases.
- **SEC-REQ-109:** Identity metadata shall not expose secrets.
- **SEC-REQ-110:** Identity logs shall preserve attribution without unnecessary personal data.
- **SEC-REQ-111:** Identity provider failure shall not result in broader fallback access.
- **SEC-REQ-112:** Identity federation shall define issuer, audience, trust and revocation boundaries.
- **SEC-REQ-113:** Cryptographic identity shall use approved algorithms and key-management controls.
- **SEC-REQ-114:** Identity assurance limitations shall be documented.
- **SEC-REQ-115:** Identity implementation shall conform to `OBDIA-ID-001`.

## 8. Authorization and Policy Enforcement

- **SEC-REQ-116:** Authorization shall default to deny.
- **SEC-REQ-117:** Authorization shall be enforced at the point of use.
- **SEC-REQ-118:** Authorization decisions shall evaluate officer, agent, institution, case, purpose, jurisdiction, resource, action, tool, data, time and risk where applicable.
- **SEC-REQ-119:** Authorization decisions shall identify the applicable policy version.
- **SEC-REQ-120:** Prohibited actions shall remain non-authorizable.
- **SEC-REQ-121:** Denied actions shall not be retried through a less restrictive path.
- **SEC-REQ-122:** Privilege elevation shall require an explicit authorized decision.
- **SEC-REQ-123:** High-impact operations shall require step-up controls or human approval.
- **SEC-REQ-124:** Human approval shall be attributable and scoped.
- **SEC-REQ-125:** Authorization obligations shall be enforced before or during action execution.
- **SEC-REQ-126:** Authorization context loss shall result in denial.
- **SEC-REQ-127:** Conflicting policy shall result in denial and review.
- **SEC-REQ-128:** Policy decision points and enforcement points shall have explicit responsibilities.
- **SEC-REQ-129:** Policy administration shall be separated from enforcement where feasible.
- **SEC-REQ-130:** Policy changes shall be versioned and reviewed.
- **SEC-REQ-131:** Authorization evidence shall include decision, reason, subject, resource, action and policy.
- **SEC-REQ-132:** Authorization denials shall be observable without exposing sensitive details.
- **SEC-REQ-133:** Cross-case and cross-purpose access shall be denied by default.
- **SEC-REQ-134:** Authorization revocation shall invalidate applicable cached decisions.
- **SEC-REQ-135:** Detailed authorization semantics remain governed by forward dependency `17_AUTHORIZATION_MODEL.md`.

## 9. Cryptography, Keys and Secrets

- **SEC-REQ-136:** Secrets shall not be stored in source code.
- **SEC-REQ-137:** Secrets shall not be stored in documentation, prompts, fixtures, screenshots or logs.
- **SEC-REQ-138:** Cryptographic keys shall have identified owners and purposes.
- **SEC-REQ-139:** Key material shall be generated using approved mechanisms.
- **SEC-REQ-140:** Key material shall be stored using protections proportionate to risk.
- **SEC-REQ-141:** Keys shall be scoped to the minimum required purpose.
- **SEC-REQ-142:** Keys shall have defined rotation and revocation procedures.
- **SEC-REQ-143:** Compromised keys shall be independently revocable.
- **SEC-REQ-144:** Key identifiers may be logged, but private key material shall not.
- **SEC-REQ-145:** Development and test keys shall be distinct from production or institutional keys.
- **SEC-REQ-146:** Placeholder secrets shall be unmistakably non-secret.
- **SEC-REQ-147:** Secret retrieval shall be authenticated, authorized and logged.
- **SEC-REQ-148:** Secret access shall fail closed.
- **SEC-REQ-149:** Secret export shall be restricted.
- **SEC-REQ-150:** Backup and recovery of keys shall preserve confidentiality and authorization.
- **SEC-REQ-151:** Cryptographic algorithm selection shall be governed and reviewable.
- **SEC-REQ-152:** Deprecated algorithms shall have migration plans.
- **SEC-REQ-153:** Certificate, token or key expiry shall result in denial rather than silent extension.
- **SEC-REQ-154:** Repository history shall be included in secret scanning.
- **SEC-REQ-155:** Secret exposure shall trigger containment, rotation, impact review and evidence preservation.

## 10. Model and Runtime Security

- **SEC-REQ-156:** Model and runtime components shall be treated as potentially fallible and manipulable.
- **SEC-REQ-157:** System instructions shall remain protected from untrusted content.
- **SEC-REQ-158:** Retrieved content shall not override system, policy or operator authority.
- **SEC-REQ-159:** Prompt injection shall be included in threat modeling.
- **SEC-REQ-160:** Model output shall be validated before security-sensitive use.
- **SEC-REQ-161:** Model output shall not directly authorize privileged action.
- **SEC-REQ-162:** Model output shall not be represented as original evidence.
- **SEC-REQ-163:** Tool invocations shall pass through separate authorization enforcement.
- **SEC-REQ-164:** Tool arguments shall be constrained or typed where feasible.
- **SEC-REQ-165:** Runtime permissions shall be least-privileged.
- **SEC-REQ-166:** Persistent memory shall be scoped by officer, case, purpose and retention.
- **SEC-REQ-167:** Cross-case memory access shall be denied by default.
- **SEC-REQ-168:** Stored memory shall remain untrusted until re-evaluated.
- **SEC-REQ-169:** Model, prompt, retrieval and policy versions shall be recorded when material.
- **SEC-REQ-170:** Provider or fallback changes shall trigger trust and risk reassessment.
- **SEC-REQ-171:** Runtime isolation shall limit blast radius.
- **SEC-REQ-172:** Unsafe runtime behavior shall support suspension and containment.
- **SEC-REQ-173:** Runtime telemetry shall preserve attribution.
- **SEC-REQ-174:** Model uncertainty and verification requirements shall be visible.
- **SEC-REQ-175:** AI runtime implementation shall remain subject to document 34 only as non-normative future research unless formally incorporated within the 01–30 baseline.

## 11. Connector and Tool Security

- **SEC-REQ-176:** Every connector shall have an identified owner and purpose.
- **SEC-REQ-177:** Every connector shall use a distinct machine identity where feasible.
- **SEC-REQ-178:** Connector permissions shall be scoped to required resources and actions.
- **SEC-REQ-179:** Read-only access shall be preferred where sufficient.
- **SEC-REQ-180:** Connector credentials shall be independently revocable.
- **SEC-REQ-181:** Connector inputs and outputs shall be treated as untrusted.
- **SEC-REQ-182:** Connector schemas shall be validated.
- **SEC-REQ-183:** Connector responses shall preserve provenance.
- **SEC-REQ-184:** Connector errors shall not trigger broader fallback permissions.
- **SEC-REQ-185:** Connector retries shall not bypass denial or human-approval requirements.
- **SEC-REQ-186:** Connector rate, volume and time limits shall be enforced where applicable.
- **SEC-REQ-187:** Connector network egress shall be constrained.
- **SEC-REQ-188:** Connector allowlists shall be explicit where feasible.
- **SEC-REQ-189:** Connector data minimization shall be enforced.
- **SEC-REQ-190:** Connector logs shall not expose secrets.
- **SEC-REQ-191:** High-risk connectors shall support immediate disablement.
- **SEC-REQ-192:** Connector compromise shall be containable from unrelated systems and cases.
- **SEC-REQ-193:** Connector dependencies shall be included in supply-chain review.
- **SEC-REQ-194:** Connector tests shall include malicious and malformed responses.
- **SEC-REQ-195:** Detailed connector requirements remain governed by forward dependency `20_CONNECTOR_SECURITY_POLICY.md`.

## 12. Network and Environment Isolation

- **SEC-REQ-196:** Network architecture shall use segmentation proportionate to risk.
- **SEC-REQ-197:** Development, test, demonstration and high-risk research networks shall be separated.
- **SEC-REQ-198:** Network location shall not establish trust.
- **SEC-REQ-199:** Ingress shall be denied unless explicitly required.
- **SEC-REQ-200:** Egress shall be denied or constrained according to purpose.
- **SEC-REQ-201:** High-risk research environments shall use controlled egress.
- **SEC-REQ-202:** Operationally isolated Dark Web research shall not interact with uncontrolled criminal infrastructure.
- **SEC-REQ-203:** Network paths shall use authenticated encryption where appropriate.
- **SEC-REQ-204:** Encryption shall not replace authorization.
- **SEC-REQ-205:** Service exposure shall be minimized.
- **SEC-REQ-206:** Administrative interfaces shall not be publicly exposed without explicit approved design.
- **SEC-REQ-207:** Default network rules shall deny unnecessary traffic.
- **SEC-REQ-208:** DNS, proxy and routing dependencies shall be included in threat modeling.
- **SEC-REQ-209:** Network telemetry shall support incident investigation.
- **SEC-REQ-210:** Network telemetry shall avoid unnecessary data collection.
- **SEC-REQ-211:** Network failure shall not result in uncontrolled direct connections.
- **SEC-REQ-212:** Remote administration shall require strong authentication and explicit authorization.
- **SEC-REQ-213:** Environment reset and destruction procedures shall be documented.
- **SEC-REQ-214:** Compromised segments shall support quarantine.
- **SEC-REQ-215:** Network architecture changes shall require applicable review and validation.

## 13. Data, Storage and Privacy Security

- **SEC-REQ-216:** Data shall be classified before processing.
- **SEC-REQ-217:** Data collection shall be limited to approved purpose and scope.
- **SEC-REQ-218:** Personal data shall be minimized.
- **SEC-REQ-219:** Real case data shall not be used in public demonstrations.
- **SEC-REQ-220:** Public demonstrations shall use synthetic or otherwise explicitly authorized safe data.
- **SEC-REQ-221:** Data at rest shall use protections proportionate to classification.
- **SEC-REQ-222:** Data in transit shall use protections proportionate to classification.
- **SEC-REQ-223:** Data access shall be authenticated, authorized and logged.
- **SEC-REQ-224:** Cross-case and cross-purpose reuse shall require separate authorization.
- **SEC-REQ-225:** Retention shall be defined and enforced.
- **SEC-REQ-226:** Deletion shall be verifiable where applicable.
- **SEC-REQ-227:** Backups shall preserve security, privacy and retention requirements.
- **SEC-REQ-228:** Temporary files and caches shall be governed.
- **SEC-REQ-229:** Data exports shall be controlled and attributable.
- **SEC-REQ-230:** Sensitive metadata shall be minimized.
- **SEC-REQ-231:** Storage failure shall not silently corrupt authoritative data.
- **SEC-REQ-232:** Data-integrity checks shall be applied where material.
- **SEC-REQ-233:** Untrusted data shall be quarantined or processed through constrained paths.
- **SEC-REQ-234:** Storage location shall not establish trust.
- **SEC-REQ-235:** Privacy by Design shall apply to architecture and implementation.

## 14. Evidence, Audit and Chain of Custody

- **SEC-REQ-236:** Original evidence and derived analysis shall remain distinguishable.
- **SEC-REQ-237:** Evidence acquisition shall preserve provenance.
- **SEC-REQ-238:** Evidence transformations shall be attributable and reproducible.
- **SEC-REQ-239:** Chain-of-custody events shall identify actor, time, action and artifact.
- **SEC-REQ-240:** Evidence access shall be least-privileged.
- **SEC-REQ-241:** Evidence access shall be logged.
- **SEC-REQ-242:** Evidence integrity shall use approved cryptographic or equivalent controls where appropriate.
- **SEC-REQ-243:** Evidence-integrity failure shall trigger quarantine or qualified handling.
- **SEC-REQ-244:** AI-generated content shall not be represented as source evidence.
- **SEC-REQ-245:** Evidence deletion or alteration shall require explicit authority.
- **SEC-REQ-246:** Revocation shall not erase lawfully retained evidence or audit records.
- **SEC-REQ-247:** Audit logs shall distinguish human, agent, workload, connector and model-related actors where applicable.
- **SEC-REQ-248:** Audit logs shall be append-only or tamper-evident where feasible.
- **SEC-REQ-249:** Audit-system failure shall not permit silent privileged operation.
- **SEC-REQ-250:** Audit records shall not expose secrets.
- **SEC-REQ-251:** Audit records shall minimize unnecessary personal data.
- **SEC-REQ-252:** Timestamp and clock limitations shall be documented.
- **SEC-REQ-253:** Evidence retention shall follow purpose, law, policy and privacy requirements.
- **SEC-REQ-254:** Evidence exports shall preserve manifest and integrity references.
- **SEC-REQ-255:** Detailed evidence requirements remain governed by forward dependency `18_EVIDENCE_MODEL.md`.

## 15. Monitoring, Detection and Alerting

- **SEC-REQ-256:** Security-relevant events shall be observable.
- **SEC-REQ-257:** Monitoring objectives shall trace to threats and controls.
- **SEC-REQ-258:** Monitoring shall distinguish human, agent, workload and connector activity.
- **SEC-REQ-259:** Monitoring shall identify authorization denials and policy failures.
- **SEC-REQ-260:** Monitoring shall identify credential failures and suspicious use.
- **SEC-REQ-261:** Monitoring shall identify integrity failures.
- **SEC-REQ-262:** Monitoring shall identify connector anomalies.
- **SEC-REQ-263:** Monitoring shall identify unusual data access and export.
- **SEC-REQ-264:** Monitoring shall identify model and tool safety events where feasible.
- **SEC-REQ-265:** Monitoring data shall be protected from unauthorized alteration.
- **SEC-REQ-266:** Monitoring access shall be least-privileged.
- **SEC-REQ-267:** Monitoring retention shall be proportionate and governed.
- **SEC-REQ-268:** Monitoring shall not become unlawful or excessive surveillance.
- **SEC-REQ-269:** Alerts shall have severity, owner and response path.
- **SEC-REQ-270:** Alert suppression shall be authorized and time-bounded.
- **SEC-REQ-271:** Monitoring blind spots shall be documented.
- **SEC-REQ-272:** Monitoring failure shall generate a visible degraded state.
- **SEC-REQ-273:** Critical monitoring failure shall block or restrict privileged operation where feasible.
- **SEC-REQ-274:** Detection rules shall be versioned and validated.
- **SEC-REQ-275:** Monitoring effectiveness shall be reviewed against incidents and threat changes.

## 16. Incident Response, Containment and Recovery

- **SEC-REQ-276:** Security incidents shall have immutable identifiers.
- **SEC-REQ-277:** Incident roles and escalation paths shall be defined.
- **SEC-REQ-278:** Incident handling shall preserve evidence.
- **SEC-REQ-279:** Containment shall be scoped to minimize unrelated disruption.
- **SEC-REQ-280:** Compromised identities shall support suspension and revocation.
- **SEC-REQ-281:** Compromised connectors shall support disablement.
- **SEC-REQ-282:** Compromised runtimes shall support isolation.
- **SEC-REQ-283:** Compromised data or evidence shall support quarantine.
- **SEC-REQ-284:** Emergency containment shall not authorize prohibited conduct.
- **SEC-REQ-285:** Containment actions shall be attributable and logged.
- **SEC-REQ-286:** Recovery shall restore a known governed state.
- **SEC-REQ-287:** Recovery shall not restore revoked credentials or obsolete authority.
- **SEC-REQ-288:** Post-recovery validation shall confirm identity, authorization, integrity and configuration.
- **SEC-REQ-289:** Incident communications shall protect sensitive information.
- **SEC-REQ-290:** Material incidents shall trigger threat-model and risk reassessment.
- **SEC-REQ-291:** Material incidents shall trigger affected state and release reassessment.
- **SEC-REQ-292:** Post-incident review shall identify control improvements.
- **SEC-REQ-293:** Repeated incidents shall trigger architectural or governance review.
- **SEC-REQ-294:** Unresolved incident impact shall remain visible.
- **SEC-REQ-295:** Incident closure shall require evidence of containment, recovery and follow-up ownership.

## 17. Supply-Chain and Build Security

- **SEC-REQ-296:** Source code changes shall be attributable.
- **SEC-REQ-297:** Material changes shall use reviewed change paths.
- **SEC-REQ-298:** Dependencies shall have version and provenance records.
- **SEC-REQ-299:** Dependencies shall be assessed for security, maintenance, licensing and applicability.
- **SEC-REQ-300:** Unsupported dependencies shall have migration, containment or rejection decisions.
- **SEC-REQ-301:** Dependency updates shall trigger risk-based reassessment.
- **SEC-REQ-302:** Build processes shall be reviewable and reproducible where feasible.
- **SEC-REQ-303:** Build artifacts shall correspond to reviewed source and configuration.
- **SEC-REQ-304:** Artifact integrity shall be verifiable.
- **SEC-REQ-305:** Build pipelines shall use least privilege.
- **SEC-REQ-306:** Pipeline secrets shall be scoped and excluded from logs.
- **SEC-REQ-307:** Model and dataset provenance shall be recorded where material.
- **SEC-REQ-308:** Prompt templates and policy bundles shall be versioned.
- **SEC-REQ-309:** Generated code shall receive human review.
- **SEC-REQ-310:** AI-assisted changes shall receive human review.
- **SEC-REQ-311:** Third-party services shall receive only approved minimum data.
- **SEC-REQ-312:** Vendor claims shall not replace project-specific validation.
- **SEC-REQ-313:** Supply-chain compromise shall be included in threat modeling.
- **SEC-REQ-314:** Release manifests shall identify material dependencies.
- **SEC-REQ-315:** Supply-chain exceptions shall have owners, expiry and residual-risk records.

## 18. Secure Configuration and Defaults

- **SEC-REQ-316:** Security-relevant configuration shall be versioned.
- **SEC-REQ-317:** Configuration changes shall be attributable.
- **SEC-REQ-318:** Default configuration shall deny unnecessary privilege and exposure.
- **SEC-REQ-319:** Dangerous optional behavior shall be disabled by default.
- **SEC-REQ-320:** Development shortcuts shall not become release defaults silently.
- **SEC-REQ-321:** Debug interfaces shall be disabled outside approved environments.
- **SEC-REQ-322:** Configuration secrets shall be separated from configuration data.
- **SEC-REQ-323:** Configuration schema shall be validated.
- **SEC-REQ-324:** Configuration drift shall be detectable where feasible.
- **SEC-REQ-325:** Policy and configuration version shall be included in material validation evidence.
- **SEC-REQ-326:** Environment-specific configuration shall be explicit.
- **SEC-REQ-327:** Unknown configuration values shall not default to broader permission.
- **SEC-REQ-328:** Configuration rollback shall restore a known governed state.
- **SEC-REQ-329:** Configuration changes affecting architecture shall require an ADR where applicable.
- **SEC-REQ-330:** Configuration changes affecting threats shall trigger threat-model review.
- **SEC-REQ-331:** Configuration changes affecting privacy shall trigger privacy review.
- **SEC-REQ-332:** Configuration examples shall use non-secret placeholders.
- **SEC-REQ-333:** Secure baseline deviations shall be recorded.
- **SEC-REQ-334:** Expired deviations shall result in rollback, denial or renewed review.
- **SEC-REQ-335:** Configuration validation shall include negative and failure cases.

## 19. Threat Modeling and Risk Treatment

- **SEC-REQ-336:** Threat modeling shall begin before implementation.
- **SEC-REQ-337:** A missing or unreviewed threat model shall block material implementation.
- **SEC-REQ-338:** Threat models shall identify assets, actors, boundaries, abuse cases and assumptions.
- **SEC-REQ-339:** Threat models shall include AI-specific threats.
- **SEC-REQ-340:** Threat models shall include identity and authorization threats.
- **SEC-REQ-341:** Threat models shall include evidence-integrity threats.
- **SEC-REQ-342:** Threat models shall include privacy and fundamental-rights harms.
- **SEC-REQ-343:** Threat models shall include supply-chain threats.
- **SEC-REQ-344:** Threat models shall include monitoring and containment failure.
- **SEC-REQ-345:** Security controls shall map to threat identifiers.
- **SEC-REQ-346:** Residual risks shall have identified human owners.
- **SEC-REQ-347:** An AI agent shall not accept risk.
- **SEC-REQ-348:** Risk acceptance shall be scoped, time-bounded and attributable.
- **SEC-REQ-349:** Risk acceptance shall not waive a canonical prohibition.
- **SEC-REQ-350:** Critical unresolved risk shall block implementation or release.
- **SEC-REQ-351:** Unknown risk shall not be assigned an unsupported low rating.
- **SEC-REQ-352:** Risk treatment status shall be evidence-based.
- **SEC-REQ-353:** Risk reassessment shall occur when assumptions or dependencies change.
- **SEC-REQ-354:** Forward dependency `26_AI_RISK_REGISTER.md` shall be reconciled before normative approval.
- **SEC-REQ-355:** Security exceptions shall remain narrow, revocable, auditable and time-bounded.

## 20. Security Validation

- **SEC-REQ-356:** Security requirements shall have testable acceptance criteria.
- **SEC-REQ-357:** Security testing shall trace to threats and controls.
- **SEC-REQ-358:** Authorization tests shall include Permit, Deny, Not Applicable and Indeterminate outcomes where supported.
- **SEC-REQ-359:** Negative tests shall be mandatory for privileged operations.
- **SEC-REQ-360:** Failure-path tests shall be mandatory for critical controls.
- **SEC-REQ-361:** Revocation tests shall verify propagation.
- **SEC-REQ-362:** Credential-expiry tests shall verify denial.
- **SEC-REQ-363:** Prompt-injection tests shall use synthetic or authorized defensive fixtures.
- **SEC-REQ-364:** Connector tests shall include malformed and malicious responses.
- **SEC-REQ-365:** Evidence tests shall verify provenance and integrity behavior.
- **SEC-REQ-366:** Monitoring tests shall verify alert generation and degraded-state behavior.
- **SEC-REQ-367:** Containment tests shall verify identity, connector and runtime disablement.
- **SEC-REQ-368:** Rollback tests shall verify restoration without obsolete authority.
- **SEC-REQ-369:** Supply-chain tests shall verify artifact provenance and integrity.
- **SEC-REQ-370:** Tests shall identify exact source, configuration, policy, model, dependency and environment versions.
- **SEC-REQ-371:** Failed tests shall remain visible.
- **SEC-REQ-372:** Passing tests shall not be generalized beyond scope.
- **SEC-REQ-373:** Automated tests shall not replace human security review.
- **SEC-REQ-374:** Document 22 remains a forward dependency for complete testing governance.
- **SEC-REQ-375:** Validation evidence shall be retained in governed repository locations.

## 21. Security Traceability

- **SEC-REQ-376:** Every security requirement shall trace to one or more threats, principles or governing constraints.
- **SEC-REQ-377:** Every implemented security control shall trace back to requirements.
- **SEC-REQ-378:** Every security test shall trace to requirements and threats.
- **SEC-REQ-379:** Every validation result shall identify exact implementation versions.
- **SEC-REQ-380:** Every accepted residual risk shall trace to affected requirements and controls.
- **SEC-REQ-381:** Every exception shall trace to its compensating controls.
- **SEC-REQ-382:** Every incident shall trace to affected threats and controls.
- **SEC-REQ-383:** Every release shall trace to security-review evidence.
- **SEC-REQ-384:** Traceability shall be bidirectional.
- **SEC-REQ-385:** Broken traceability affecting authority, identity, authorization, privacy, evidence or release shall be blocking.
- **SEC-REQ-386:** Planned controls shall not be represented as implemented controls.
- **SEC-REQ-387:** Implemented controls shall not be represented as validated without evidence.
- **SEC-REQ-388:** Control status shall be accurate at feature level.
- **SEC-REQ-389:** Superseded requirements and controls shall identify successors.
- **SEC-REQ-390:** Traceability records shall use immutable identifiers.
- **SEC-REQ-391:** Traceability records shall not contain secrets.
- **SEC-REQ-392:** Traceability records shall minimize personal and case data.
- **SEC-REQ-393:** Cross-reference validation shall occur before approval.
- **SEC-REQ-394:** Baseline freeze shall verify security traceability across documents 01–30.
- **SEC-REQ-395:** No document above 30 shall acquire normative security authority by implication.

## 22. Security Review Gate

- **SEC-REQ-396:** Security review shall identify the exact artifact and commit.
- **SEC-REQ-397:** Security review shall identify applicable domains and boundaries.
- **SEC-REQ-398:** Security review shall verify threat-model completeness.
- **SEC-REQ-399:** Security review shall verify control ownership.
- **SEC-REQ-400:** Security review shall verify authorization and failure behavior.
- **SEC-REQ-401:** Security review shall verify privacy and evidence protections.
- **SEC-REQ-402:** Security review shall verify monitoring and containment.
- **SEC-REQ-403:** Security review shall verify supply-chain and dependency controls.
- **SEC-REQ-404:** Security review shall verify implementation-status accuracy.
- **SEC-REQ-405:** Security review shall verify validation evidence.
- **SEC-REQ-406:** Security review shall identify residual risks and exceptions.
- **SEC-REQ-407:** Critical findings shall block progression.
- **SEC-REQ-408:** Material post-review change shall invalidate affected review evidence.
- **SEC-REQ-409:** Role concentration shall be disclosed.
- **SEC-REQ-410:** Internal review shall not be represented as independent external assurance.
- **SEC-REQ-411:** Founder approval shall not substitute for required specialist review.
- **SEC-REQ-412:** Automated findings shall remain advisory until adopted by a human reviewer.
- **SEC-REQ-413:** Review non-applicability shall have rationale.
- **SEC-REQ-414:** Review approval shall identify actor, role, date, scope and exact artifact.
- **SEC-REQ-415:** Approval of this Draft for repository incorporation shall not establish normative security-baseline approval.

## 23. Minimum Validation Checklist

Before approval of this baseline or a material implementation, confirm:

- [ ] human and institutional accountability are explicit;
- [ ] the AI agent has no independent legal authority;
- [ ] prohibited conduct remains non-authorizable;
- [ ] security domains and owners are identified;
- [ ] assets and trust boundaries are documented;
- [ ] threat modeling is complete;
- [ ] Zero Trust and continuous verification are applied;
- [ ] identity and cryptographic binding requirements are satisfied;
- [ ] authorization defaults to deny;
- [ ] secrets, keys and credentials are protected;
- [ ] model, prompt, retrieval, memory and tool controls are defined;
- [ ] connectors use least privilege and constrained egress;
- [ ] environments and networks are isolated appropriately;
- [ ] data minimization and privacy controls are present;
- [ ] evidence provenance and audit integrity are preserved;
- [ ] monitoring, containment and recovery are testable;
- [ ] dependencies and builds have provenance;
- [ ] secure defaults and drift controls are defined;
- [ ] risks and exceptions have human owners;
- [ ] validation evidence identifies exact versions;
- [ ] traceability is bidirectional;
- [ ] forward dependencies 17, 18, 19, 20, 22, 26 and 30 are recorded;
- [ ] limitations and unresolved findings remain visible;
- [ ] Project Founder approval exists before status becomes Approved.


## 24. Limitations

- This baseline does not implement security controls.
- It does not select identity, policy, cloud, model, connector, monitoring, storage or cryptographic products.
- It does not establish legal compliance, certification, accreditation, production readiness or institutional authorization.
- It does not replace the detailed Authorization Model, Evidence Model, Agent Lifecycle Model or Connector Security Policy.
- It does not guarantee absence of vulnerabilities, misuse or human error.
- Internal review is not independent certification.
- Continuous verification can reduce but cannot eliminate uncertainty.
- Cryptographic integrity does not prove lawful collection or semantic correctness.
- Network isolation does not establish authorization.
- Documents 17, 18, 19, 20, 22, 26 and 30 remain forward dependencies for full approval.


## 25. Change Control

Every material change shall identify rationale, affected domains and artifacts, dependencies, threats, security and privacy impact, evidence impact, implementation impact, migration, validation, rollback, authority and version effect.

- **SEC-REQ-416:** Editorial corrections shall use a patch version when meaning is unchanged.
- **SEC-REQ-417:** Backward-compatible substantive additions shall use a minor version.
- **SEC-REQ-418:** Incompatible normative security changes shall use a major version.
- **SEC-REQ-419:** Material security-architecture decisions shall require an ADR where applicable.
- **SEC-REQ-420:** Changes shall require Project Founder approval.
- **SEC-REQ-421:** Changes shall receive Security Reviewer assessment.
- **SEC-REQ-422:** Privacy, evidence and authorization impacts shall receive specialist review where applicable.
- **SEC-REQ-423:** Changes shall identify affected threats, controls, tests, risks and implementation.
- **SEC-REQ-424:** Changes shall include migration, validation and rollback analysis.
- **SEC-REQ-425:** Changes shall not retroactively fabricate security approval or validation evidence.
- **SEC-REQ-426:** Historical approved decisions and evidence shall not be silently rewritten.
- **SEC-REQ-427:** Forward-dependency reconciliation shall occur before this baseline becomes Approved.

## 26. Consolidation Record

Version 1.0.0 consolidates the two existing Security Architecture Baseline variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-SEC-001`;
- normalizes the authoritative filename to `16_SECURITY_ARCHITECTURE_BASELINE.md`;
- rejects the unreviewed legacy `Approved Baseline` label as insufficient evidence;
- establishes constitutional lifecycle status `Draft`;
- preserves all original security domains: Identity, Runtime, Connectors, Storage, Network, Evidence, Governance and Monitoring;
- preserves all original mandatory controls: Zero Trust, Defense in Depth, Secure by Design, Privacy by Design, Least Privilege, Explicit Authorization, Cryptographic Identity, audit integrity, threat modeling before implementation, secure defaults and continuous verification;
- adds domain ownership, trust-boundary mapping, secure failure, keys and secrets, model and tool security, connector constraints, network isolation, data and evidence protection, monitoring, containment, recovery, supply-chain, validation, risk treatment and traceability;
- identifies documents 17, 18, 19, 20, 22, 26 and 30 as forward dependencies blocking normative approval;
- treats Enterprise and non-Enterprise legacy files as source variants of the same immutable document;
- creates no implementation, production claim, security certification or operational authorization.

## 27. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Security Reviewer; approval reserved to Project Founder | Constitutional consolidation of both legacy baselines: preserved all domains and mandatory controls; resolved conflicting status; added enforceable security architecture, validation, containment, risk and traceability requirements. |
| 1.0.1 | 2026-08-06 | Draft | Security Reviewer; approval reserved to Project Founder | Reconciled obsolete forward-dependency filenames with the final canonical Knowledge Pack 01–30 filenames; no requirement text, authority, lifecycle status, implementation state or validation claim changed. |
