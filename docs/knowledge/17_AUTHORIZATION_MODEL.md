# AUTHORIZATION MODEL

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-AUTH-001 |
| **Title** | Authorization Model |
| **Version** | 1.0.1 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Security Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define the authoritative OBDIA authorization object, contextual inputs, policy-decision and enforcement responsibilities, evaluation sequence, decision outcomes, human-approval triggers, obligations, revocation, failure behavior, audit evidence and validation requirements. |
| **Scope** | Human officers, institutionally issued agents, workloads, sessions, tools, connectors, models, data, evidence, repositories, administrative operations and controlled research environments within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `06_IDENTITY_AND_DELEGATION_MODEL.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; forward dependencies `18_EVIDENCE_MODEL.md`, `19_AGENT_LIFECYCLE_MODEL.md`, `20_CONNECTOR_SECURITY_POLICY.md`, `22_TESTING_STANDARD.md`, `26_AI_RISK_REGISTER.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-ARCH-001`; `OBDIA-ID-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-SEC-001` |
| **Cross-References** | Documents 18–30; accepted ADRs; policy bundles; delegation artifacts; identity, trust, risk, evidence, audit, implementation, testing, incident and release records |
| **Assumptions** | Authorization is evaluated for explicitly identified operations using current identity, delegation, policy, trust and contextual evidence. No operation is authorized merely because it is technically possible or requested by a model. |
| **Constraints** | Authorization shall not create legal authority, waive canonical prohibitions, transfer human credentials to an agent, permit cross-case or cross-purpose reuse by default, or authorize uncontrolled criminal-infrastructure interaction. |
| **Security Considerations** | Authorization failure can cause privilege escalation, confused-deputy behavior, stale-permission use, cross-case contamination, policy bypass, unsafe tool invocation, unauthorized evidence access, silent fallback and loss of human accountability. |
| **Validation Criteria** | Every protected operation has a complete authorization object, explicit policy version, deterministic evaluation sequence, deny-by-default outcome, enforceable obligations, revocation behavior, attributable audit evidence, negative tests and bidirectional traceability. |
| **Implementation Relationship** | This model defines logical authorization requirements and responsibilities. It does not select a policy language, identity provider, token format, policy engine, product or deployment architecture. |

---

## 1. Purpose

This document defines the authorization model for the Officer-Bound Digital Investigation Agent project.

The original sources required every operation to be bound to:

- Officer;
- Institution;
- Case;
- Purpose;
- Jurisdiction;
- Time;
- Approved Tools;
- Data Scope;
- Policy Compliance.

They also required authorization to be explicit, revocable and auditable, with `Deny` as the default decision.

This consolidation preserves every original field and rule, and adds the missing object model, evaluation order, decision and enforcement roles, human-approval triggers, separation of duties, obligations, denial reasons, revocation, audit schema, cross-case isolation, failure handling and test criteria.

Authorization in this project is a technical and governance decision that a specific subject may perform a specific action on a specific resource under specific conditions. It is not legal authority.


## 2. Fundamental Authorization Rules

- **AUTH-REQ-001:** Authorization shall be explicit.
- **AUTH-REQ-002:** Authorization shall be revocable.
- **AUTH-REQ-003:** Authorization shall be auditable.
- **AUTH-REQ-004:** The default authorization decision shall be `Deny`.
- **AUTH-REQ-005:** Authentication shall not be treated as authorization.
- **AUTH-REQ-006:** Authorization shall not be treated as legal authority.
- **AUTH-REQ-007:** Model output shall not be treated as authorization.
- **AUTH-REQ-008:** Trust shall not be treated as authorization.
- **AUTH-REQ-009:** Technical capability shall not be treated as authorization.
- **AUTH-REQ-010:** Repository access shall not establish authorization to perform investigative operations.
- **AUTH-REQ-011:** Every protected operation shall be evaluated against a current authorization context.
- **AUTH-REQ-012:** Missing, stale, conflicting or unverifiable authorization context shall result in `Deny` or `Indeterminate`, never `Permit`.
- **AUTH-REQ-013:** Canonical prohibitions shall remain non-authorizable.
- **AUTH-REQ-014:** An AI agent shall not grant itself authority.
- **AUTH-REQ-015:** An AI agent shall not approve its own privilege elevation.
- **AUTH-REQ-016:** An AI agent shall not delegate permissions outside an approved human-governed path.
- **AUTH-REQ-017:** Human credentials shall not be transferred to or embedded in an agent.
- **AUTH-REQ-018:** Authorization shall be limited to one identified officer-agent binding at a time.
- **AUTH-REQ-019:** Authorization shall preserve institutional accountability.
- **AUTH-REQ-020:** Authorization shall be case-, purpose-, jurisdiction-, time-, tool- and data-scoped where applicable.
- **AUTH-REQ-021:** Authorization decisions shall be attributable to identified policies and evidence.
- **AUTH-REQ-022:** Authorization enforcement shall fail closed.
- **AUTH-REQ-023:** Authorization controls shall use least privilege.
- **AUTH-REQ-024:** Authorization changes shall be versioned and reviewable.

## 3. Authorization Object

- **AUTH-REQ-025:** Every protected request shall create or reference an authorization object.
- **AUTH-REQ-026:** The authorization object shall have an immutable request identifier.
- **AUTH-REQ-027:** The authorization object shall identify the human officer.
- **AUTH-REQ-028:** The authorization object shall identify the accountable institution.
- **AUTH-REQ-029:** The authorization object shall identify the agent identity where an agent participates.
- **AUTH-REQ-030:** The authorization object shall identify the workload or runtime identity where applicable.
- **AUTH-REQ-031:** The authorization object shall identify the active case or mandate.
- **AUTH-REQ-032:** The authorization object shall identify the approved purpose.
- **AUTH-REQ-033:** The authorization object shall identify the applicable jurisdiction.
- **AUTH-REQ-034:** The authorization object shall identify request time and validity window.
- **AUTH-REQ-035:** The authorization object shall identify the requested action.
- **AUTH-REQ-036:** The authorization object shall identify the target resource.
- **AUTH-REQ-037:** The authorization object shall identify the approved tool or connector where applicable.
- **AUTH-REQ-038:** The authorization object shall identify the requested data scope.
- **AUTH-REQ-039:** The authorization object shall identify the applicable policy bundle and version.
- **AUTH-REQ-040:** The authorization object shall identify the delegation artifact where applicable.
- **AUTH-REQ-041:** The authorization object shall identify the current trust and risk state.
- **AUTH-REQ-042:** The authorization object shall identify required human approval where applicable.
- **AUTH-REQ-043:** The authorization object shall identify relevant session or transaction identifiers.
- **AUTH-REQ-044:** The authorization object shall identify the requested result or side effect.
- **AUTH-REQ-045:** The authorization object shall identify environmental attributes used in the decision.
- **AUTH-REQ-046:** The authorization object shall identify source and freshness of material attributes.
- **AUTH-REQ-047:** The authorization object shall distinguish supplied attributes from derived attributes.
- **AUTH-REQ-048:** The authorization object shall not contain reusable human credentials.
- **AUTH-REQ-049:** The authorization object shall not contain unnecessary personal or case data.
- **AUTH-REQ-050:** The authorization object shall be integrity-protected where material.
- **AUTH-REQ-051:** The authorization object shall remain traceable to the final decision.
- **AUTH-REQ-052:** The authorization object shall remain traceable to enforcement and audit events.
- **AUTH-REQ-053:** Duplicate request identifiers shall be rejected.
- **AUTH-REQ-054:** Malformed authorization objects shall result in denial.

## 4. Subjects and Identity Inputs

- **AUTH-REQ-055:** Authorized human subjects shall be identified through approved institutional identity.
- **AUTH-REQ-056:** Agent subjects shall use institutionally issued identities.
- **AUTH-REQ-057:** Agent subjects shall remain cryptographically bound to exactly one identified officer at a time.
- **AUTH-REQ-058:** Workload subjects shall use distinct machine identities where feasible.
- **AUTH-REQ-059:** Connector subjects shall use distinct machine identities where feasible.
- **AUTH-REQ-060:** Shared machine subjects shall require an approved exception.
- **AUTH-REQ-061:** Subject identity status shall be current at decision time.
- **AUTH-REQ-062:** Suspended, revoked or expired subjects shall not receive `Permit`.
- **AUTH-REQ-063:** Subject identity proofing strength shall be proportionate to operation risk.
- **AUTH-REQ-064:** Subject attributes shall identify authoritative issuer and freshness.
- **AUTH-REQ-065:** Subject roles shall not by themselves grant unrestricted authority.
- **AUTH-REQ-066:** Subject group membership shall not bypass case, purpose or jurisdiction constraints.
- **AUTH-REQ-067:** Subject aliases shall resolve to one authoritative identity.
- **AUTH-REQ-068:** Subject rebinding shall not preserve prior authorization automatically.
- **AUTH-REQ-069:** Compromised subjects shall be containable independently.
- **AUTH-REQ-070:** Unknown subjects shall be denied.
- **AUTH-REQ-071:** Anonymous subjects shall not perform privileged or evidentiary operations.
- **AUTH-REQ-072:** Human and non-human subjects shall remain distinguishable in audit records.
- **AUTH-REQ-073:** Subject identity confidence shall not replace verification.
- **AUTH-REQ-074:** Subject identity implementation shall conform to `OBDIA-ID-001`.

## 5. Resources, Actions and Operations

- **AUTH-REQ-075:** Every protected resource shall have an authoritative resource identifier.
- **AUTH-REQ-076:** Resources shall have identified owners or custodians.
- **AUTH-REQ-077:** Resources shall have classification metadata where applicable.
- **AUTH-REQ-078:** Resources shall identify case and purpose boundaries where applicable.
- **AUTH-REQ-079:** Resources shall identify jurisdictional constraints where applicable.
- **AUTH-REQ-080:** Resources shall identify permitted action categories.
- **AUTH-REQ-081:** Resources shall identify retention and evidence requirements where applicable.
- **AUTH-REQ-082:** Resources shall identify connector or system ownership where applicable.
- **AUTH-REQ-083:** Resources shall not inherit unrestricted access solely from directory or network location.
- **AUTH-REQ-084:** Resource aliases shall resolve to one canonical target.
- **AUTH-REQ-085:** Unknown resources shall be denied.
- **AUTH-REQ-086:** Resource classification ambiguity shall result in the more restrictive handling.
- **AUTH-REQ-087:** Derived resources shall retain lineage to source resources.
- **AUTH-REQ-088:** Evidence resources shall remain distinguishable from analytical outputs.
- **AUTH-REQ-089:** Resource metadata shall be integrity-protected where material.
- **AUTH-REQ-090:** Resource ownership changes shall trigger authorization review.
- **AUTH-REQ-091:** Resource movement shall not broaden access automatically.
- **AUTH-REQ-092:** Resource deletion shall require explicit authority.
- **AUTH-REQ-093:** Resource export shall require distinct authorization where applicable.
- **AUTH-REQ-094:** Resource scope shall be enforceable at the lowest practical granularity.

## 6. Controlled Actions

- **AUTH-REQ-095:** Every protected request shall identify one or more explicit actions.
- **AUTH-REQ-096:** Actions shall use a controlled action vocabulary.
- **AUTH-REQ-097:** Ambiguous actions shall be denied.
- **AUTH-REQ-098:** Compound actions shall be decomposed when components have different authorization requirements.
- **AUTH-REQ-099:** Read, create, update, delete, execute, export, share, transform and administer shall remain distinguishable.
- **AUTH-REQ-100:** Evidence acquisition shall remain distinguishable from evidence analysis.
- **AUTH-REQ-101:** Evidence transformation shall remain distinguishable from source modification.
- **AUTH-REQ-102:** Administrative actions shall require stronger controls than ordinary use.
- **AUTH-REQ-103:** Credential issuance, rotation and revocation shall be protected actions.
- **AUTH-REQ-104:** Policy creation, approval, publication and rollback shall be protected actions.
- **AUTH-REQ-105:** Connector onboarding, enablement and disablement shall be protected actions.
- **AUTH-REQ-106:** Model or provider changes shall be protected actions where they affect trust or capability.
- **AUTH-REQ-107:** Tool execution shall be a protected action.
- **AUTH-REQ-108:** Memory read, write and delete operations shall be protected actions.
- **AUTH-REQ-109:** Cross-case search shall be a separately protected action.
- **AUTH-REQ-110:** Data export shall be a separately protected action.
- **AUTH-REQ-111:** Public release shall be a separately protected action.
- **AUTH-REQ-112:** Emergency containment shall be a separately protected action.
- **AUTH-REQ-113:** Unknown or unsupported actions shall be denied.
- **AUTH-REQ-114:** Action naming shall conform to `OBDIA-NAME-001`.

## 7. Mandatory Authorization Context

- **AUTH-REQ-115:** Officer identity shall be evaluated for every protected operation.
- **AUTH-REQ-116:** Institution shall be evaluated for every protected operation.
- **AUTH-REQ-117:** Case or mandate shall be evaluated where the operation is investigative or case-scoped.
- **AUTH-REQ-118:** Purpose shall be evaluated where data or capability is purpose-limited.
- **AUTH-REQ-119:** Jurisdiction shall be evaluated where location, law or institutional authority matters.
- **AUTH-REQ-120:** Time validity shall be evaluated for every time-bounded delegation or credential.
- **AUTH-REQ-121:** Approved tool scope shall be evaluated for every tool or connector operation.
- **AUTH-REQ-122:** Data scope shall be evaluated for every data operation.
- **AUTH-REQ-123:** Policy compliance shall be evaluated before a `Permit` decision.
- **AUTH-REQ-124:** Agent identity and binding shall be evaluated when an agent participates.
- **AUTH-REQ-125:** Workload identity shall be evaluated where a workload executes the operation.
- **AUTH-REQ-126:** Credential and revocation status shall be evaluated.
- **AUTH-REQ-127:** Session validity shall be evaluated.
- **AUTH-REQ-128:** Trust and posture attributes shall be evaluated when material.
- **AUTH-REQ-129:** Risk state shall be evaluated when policy requires it.
- **AUTH-REQ-130:** Human-approval state shall be evaluated where approval is mandatory.
- **AUTH-REQ-131:** Data classification shall be evaluated.
- **AUTH-REQ-132:** Resource ownership or custody shall be evaluated where applicable.
- **AUTH-REQ-133:** Evidence status shall be evaluated for evidence operations.
- **AUTH-REQ-134:** Connector classification shall be evaluated for external operations.
- **AUTH-REQ-135:** Network location may inform context but shall not establish authorization.
- **AUTH-REQ-136:** Source reputation may inform context but shall not establish authorization.
- **AUTH-REQ-137:** Environmental attributes shall identify their authoritative source.
- **AUTH-REQ-138:** Stale attributes shall not be used beyond defined validity.
- **AUTH-REQ-139:** Context conflicts shall result in denial or `Indeterminate`.
- **AUTH-REQ-140:** Context omissions shall not be silently replaced by broader defaults.
- **AUTH-REQ-141:** Context evaluation shall be deterministic for the same inputs and policy version.
- **AUTH-REQ-142:** Context derivation rules shall be reviewable.
- **AUTH-REQ-143:** Sensitive context shall be minimized in logs.
- **AUTH-REQ-144:** Context changes during a long operation shall trigger re-evaluation where material.

## 8. Delegation and Authority Derivation

- **AUTH-REQ-145:** Authorization shall be derived only from valid institutional and human delegation.
- **AUTH-REQ-146:** Delegation artifacts shall identify issuer, subject, institution, scope and validity.
- **AUTH-REQ-147:** Delegation artifacts shall identify case, purpose, jurisdiction, tools and data scope where applicable.
- **AUTH-REQ-148:** Delegation artifacts shall identify permitted and prohibited actions.
- **AUTH-REQ-149:** Delegation artifacts shall identify human-approval requirements.
- **AUTH-REQ-150:** Delegation artifacts shall identify revocation references.
- **AUTH-REQ-151:** Delegation artifacts shall be integrity-protected.
- **AUTH-REQ-152:** Expired delegation shall not authorize operations.
- **AUTH-REQ-153:** Suspended or revoked delegation shall not authorize operations.
- **AUTH-REQ-154:** Delegation shall not be transitive unless an explicit approved policy path exists.
- **AUTH-REQ-155:** An agent shall not create a new delegation for another agent autonomously.
- **AUTH-REQ-156:** An agent shall not broaden its own delegation.
- **AUTH-REQ-157:** Subdelegation shall require explicit policy, attributable human authority and complete audit.
- **AUTH-REQ-158:** Delegation changes shall not alter active authorization silently.
- **AUTH-REQ-159:** Delegation replacement shall invalidate or bound prior delegation use.
- **AUTH-REQ-160:** Delegation scope shall be no broader than the issuing authority.
- **AUTH-REQ-161:** Delegation shall not include canonical prohibitions.
- **AUTH-REQ-162:** Human credentials shall not serve as delegation artifacts.
- **AUTH-REQ-163:** Unknown delegation issuers shall be denied.
- **AUTH-REQ-164:** Delegation validation shall conform to `OBDIA-ID-001`.

## 9. Logical Authorization Components

- **AUTH-REQ-165:** The Policy Administration Point shall manage policy authoring and controlled publication.
- **AUTH-REQ-166:** The Policy Administration Point shall not enforce requests directly unless separation is impractical and documented.
- **AUTH-REQ-167:** The Policy Information Point shall supply authoritative contextual attributes.
- **AUTH-REQ-168:** The Policy Information Point shall identify attribute source and freshness.
- **AUTH-REQ-169:** The Policy Decision Point shall evaluate authorization requests against an exact policy version.
- **AUTH-REQ-170:** The Policy Decision Point shall return a controlled decision outcome.
- **AUTH-REQ-171:** The Policy Decision Point shall return applicable reasons and obligations.
- **AUTH-REQ-172:** The Policy Enforcement Point shall intercept protected operations.
- **AUTH-REQ-173:** The Policy Enforcement Point shall obtain a current decision before protected execution.
- **AUTH-REQ-174:** The Policy Enforcement Point shall enforce decision obligations.
- **AUTH-REQ-175:** The Policy Enforcement Point shall deny on unavailable or invalid decision where fail-closed behavior applies.
- **AUTH-REQ-176:** The Policy Enforcement Point shall record enforcement outcome.
- **AUTH-REQ-177:** The human-approval service or workflow shall record attributable approval decisions.
- **AUTH-REQ-178:** The audit service shall record decision and enforcement evidence.
- **AUTH-REQ-179:** The revocation service shall distribute revocation state.
- **AUTH-REQ-180:** Policy components shall use distinct identities where feasible.
- **AUTH-REQ-181:** Policy components shall use least privilege.
- **AUTH-REQ-182:** Policy component communication shall be authenticated and integrity-protected.
- **AUTH-REQ-183:** Policy component failure shall not broaden access.
- **AUTH-REQ-184:** Component responsibilities shall remain product-neutral and logically separable.
- **AUTH-REQ-185:** Combining roles shall require documented risk and compensating controls.
- **AUTH-REQ-186:** Policy component configuration shall be versioned.
- **AUTH-REQ-187:** Policy component changes shall trigger affected validation.
- **AUTH-REQ-188:** Policy component time sources shall be governed where time validity matters.
- **AUTH-REQ-189:** Policy component logs shall not expose secrets.
- **AUTH-REQ-190:** Policy component monitoring shall detect degraded or inconsistent state.

## 10. Evaluation Sequence

- **AUTH-REQ-191:** The enforcement point shall intercept the requested operation before execution.
- **AUTH-REQ-192:** The request shall be normalized into the authoritative authorization object.
- **AUTH-REQ-193:** Subject identity and current status shall be validated.
- **AUTH-REQ-194:** Agent-officer and institutional binding shall be validated where applicable.
- **AUTH-REQ-195:** Delegation validity and scope shall be validated.
- **AUTH-REQ-196:** Resource and action identifiers shall be validated.
- **AUTH-REQ-197:** Case, purpose, jurisdiction, time, tool and data constraints shall be evaluated.
- **AUTH-REQ-198:** Applicable policies shall be selected deterministically.
- **AUTH-REQ-199:** Policy conflicts shall be resolved using an approved combining rule.
- **AUTH-REQ-200:** Canonical prohibitions shall be evaluated before permissive rules.
- **AUTH-REQ-201:** Explicit deny rules shall take precedence over permissive rules unless a superior policy states otherwise.
- **AUTH-REQ-202:** Required human approval shall be evaluated.
- **AUTH-REQ-203:** Required step-up verification shall be evaluated.
- **AUTH-REQ-204:** Required obligations shall be generated.
- **AUTH-REQ-205:** A final controlled decision outcome shall be produced.
- **AUTH-REQ-206:** The decision shall identify policy and attribute versions.
- **AUTH-REQ-207:** The enforcement point shall verify decision freshness and integrity.
- **AUTH-REQ-208:** The enforcement point shall enforce obligations before or during execution.
- **AUTH-REQ-209:** The enforcement point shall prevent execution when obligations cannot be satisfied.
- **AUTH-REQ-210:** The enforcement point shall record execution allowed, denied, failed, cancelled or completed.
- **AUTH-REQ-211:** Long-running operations shall re-evaluate authorization at defined checkpoints.
- **AUTH-REQ-212:** Material context change shall trigger re-evaluation.
- **AUTH-REQ-213:** Revocation received before completion shall stop or constrain the operation where safe.
- **AUTH-REQ-214:** Decision and enforcement evidence shall be correlated.
- **AUTH-REQ-215:** Any sequence failure shall result in denial or containment.

## 11. Decision Outcomes and Default Deny

- **AUTH-REQ-216:** Permitted decision outcomes shall be `Permit`, `Deny`, `Not Applicable` and `Indeterminate`.
- **AUTH-REQ-217:** `Permit` shall mean the request is authorized only within returned scope and obligations.
- **AUTH-REQ-218:** `Deny` shall mean the request shall not execute.
- **AUTH-REQ-219:** `Not Applicable` shall mean the evaluated policy does not govern the request and shall not by itself authorize execution.
- **AUTH-REQ-220:** `Indeterminate` shall mean a reliable decision could not be produced.
- **AUTH-REQ-221:** The default final outcome shall be `Deny`.
- **AUTH-REQ-222:** `Not Applicable` shall require evaluation of other applicable policy before any `Permit`.
- **AUTH-REQ-223:** `Indeterminate` shall resolve to `Deny` for privileged, evidentiary or external actions.
- **AUTH-REQ-224:** A policy engine timeout shall not produce `Permit`.
- **AUTH-REQ-225:** Missing attributes shall not produce `Permit` unless an approved policy explicitly treats the attribute as optional.
- **AUTH-REQ-226:** Conflicting authoritative attributes shall not produce `Permit`.
- **AUTH-REQ-227:** Invalid policy signatures or integrity checks shall not produce `Permit`.
- **AUTH-REQ-228:** Unknown actions or resources shall produce `Deny`.
- **AUTH-REQ-229:** Revoked subjects or delegations shall produce `Deny`.
- **AUTH-REQ-230:** Expired time validity shall produce `Deny`.
- **AUTH-REQ-231:** Canonical prohibitions shall produce `Deny`.
- **AUTH-REQ-232:** Outcome representation shall be consistent across human- and machine-readable records.
- **AUTH-REQ-233:** Outcome codes shall not expose sensitive denial details to unauthorized recipients.
- **AUTH-REQ-234:** Decision outcomes shall remain attributable to exact inputs and policy versions.
- **AUTH-REQ-235:** Decision outcomes shall not be overwritten in historical audit records.

## 12. Policy Model and Lifecycle

- **AUTH-REQ-236:** Authorization policies shall have immutable identifiers.
- **AUTH-REQ-237:** Authorization policies shall have semantic versions.
- **AUTH-REQ-238:** Authorization policies shall identify authority and owner.
- **AUTH-REQ-239:** Authorization policies shall identify scope and applicability.
- **AUTH-REQ-240:** Authorization policies shall identify subjects, resources, actions and conditions.
- **AUTH-REQ-241:** Authorization policies shall identify explicit deny rules.
- **AUTH-REQ-242:** Authorization policies shall identify obligations.
- **AUTH-REQ-243:** Authorization policies shall identify human-approval triggers.
- **AUTH-REQ-244:** Authorization policies shall identify step-up requirements.
- **AUTH-REQ-245:** Authorization policies shall identify expiry or review conditions where applicable.
- **AUTH-REQ-246:** Policy rules shall be deterministic or document unavoidable nondeterminism.
- **AUTH-REQ-247:** Policy combining algorithms shall be explicit.
- **AUTH-REQ-248:** Deny-overrides shall be the default combining behavior for conflicting security rules.
- **AUTH-REQ-249:** Canonical prohibitions shall override lower-level permit rules.
- **AUTH-REQ-250:** Policy inheritance shall not broaden authority implicitly.
- **AUTH-REQ-251:** Policy references shall use immutable identifiers.
- **AUTH-REQ-252:** Policy bundles shall identify included policy versions.
- **AUTH-REQ-253:** Policy publication shall be a protected operation.
- **AUTH-REQ-254:** Policy rollback shall restore a known governed version.
- **AUTH-REQ-255:** Policy changes shall undergo review and testing.
- **AUTH-REQ-256:** Material policy-architecture changes shall require an ADR where applicable.
- **AUTH-REQ-257:** Policy deprecation shall identify replacement and migration.
- **AUTH-REQ-258:** Expired policy shall not be used silently.
- **AUTH-REQ-259:** Unknown policy version shall result in denial.
- **AUTH-REQ-260:** Policy caches shall have invalidation and freshness controls.
- **AUTH-REQ-261:** Policy distribution shall be integrity-protected.
- **AUTH-REQ-262:** Policy discrepancies across decision points shall generate a degraded state.
- **AUTH-REQ-263:** Policy authors shall not approve their own high-risk policy changes without compensating controls.
- **AUTH-REQ-264:** Policy documentation shall identify rationale and affected requirements.
- **AUTH-REQ-265:** Policy status shall not be inferred from repository location.

## 13. Human Approval and Step-Up

- **AUTH-REQ-266:** Human approval shall be required when superior policy mandates it.
- **AUTH-REQ-267:** Human approval shall be required for material privilege elevation.
- **AUTH-REQ-268:** Human approval shall be required for exceptional cross-case access.
- **AUTH-REQ-269:** Human approval shall be required for high-impact evidence modification or deletion.
- **AUTH-REQ-270:** Human approval shall be required for public release.
- **AUTH-REQ-271:** Human approval shall be required for break-glass or emergency access.
- **AUTH-REQ-272:** Human approval shall identify approver identity and role.
- **AUTH-REQ-273:** Human approval shall identify exact operation and scope.
- **AUTH-REQ-274:** Human approval shall identify case, purpose and time limits.
- **AUTH-REQ-275:** Human approval shall identify conditions and obligations.
- **AUTH-REQ-276:** Human approval shall be recorded before execution unless an approved emergency process states otherwise.
- **AUTH-REQ-277:** Approval silence shall not be treated as approval.
- **AUTH-REQ-278:** Approval expiry shall result in denial.
- **AUTH-REQ-279:** Approval reuse shall be prohibited beyond stated scope.
- **AUTH-REQ-280:** Approval delegation shall follow governance rules.
- **AUTH-REQ-281:** An AI system shall not be final approval authority.
- **AUTH-REQ-282:** Automated recommendation shall not replace approval.
- **AUTH-REQ-283:** High-risk approvals should use separation of duties.
- **AUTH-REQ-284:** Approvers shall not approve operations outside their authority.
- **AUTH-REQ-285:** Role concentration shall be disclosed.
- **AUTH-REQ-286:** Approval conflicts shall result in denial or escalation.
- **AUTH-REQ-287:** Approval revocation shall stop future use.
- **AUTH-REQ-288:** Approval evidence shall be retained.
- **AUTH-REQ-289:** Approval records shall not contain unnecessary personal data.
- **AUTH-REQ-290:** Human approval shall not authorize canonical prohibitions.

## 14. Separation of Duties

- **AUTH-REQ-291:** Policy authoring, approval and enforcement responsibilities shall be separated where feasible.
- **AUTH-REQ-292:** Identity issuance and authorization approval shall be separated where feasible.
- **AUTH-REQ-293:** Evidence custody and evidence deletion approval shall be separated where feasible.
- **AUTH-REQ-294:** Connector development and high-risk connector enablement shall be separated where feasible.
- **AUTH-REQ-295:** Risk ownership and independent review shall remain distinguishable.
- **AUTH-REQ-296:** An agent shall not approve its own requested operation.
- **AUTH-REQ-297:** A workload shall not grant itself broader permissions.
- **AUTH-REQ-298:** A connector shall not approve its own onboarding.
- **AUTH-REQ-299:** Policy administrators shall not bypass enforcement.
- **AUTH-REQ-300:** Enforcement operators shall not alter policy silently.
- **AUTH-REQ-301:** Reviewers shall identify role concentration.
- **AUTH-REQ-302:** Role concentration shall require compensating controls.
- **AUTH-REQ-303:** Compensating controls may include immutable history, dual acknowledgement, time limits and independent logs.
- **AUTH-REQ-304:** Separation-of-duties exceptions shall have identifiers, owners and expiry.
- **AUTH-REQ-305:** Expired exceptions shall result in denial or renewed review.
- **AUTH-REQ-306:** Separation-of-duties controls shall be tested.
- **AUTH-REQ-307:** Independent external assurance shall not be claimed without evidence.
- **AUTH-REQ-308:** Human accountability shall remain identifiable despite role concentration.
- **AUTH-REQ-309:** Emergency access shall not permanently collapse role separation.
- **AUTH-REQ-310:** Repeated separation exceptions shall trigger governance review.

## 15. Obligations and Constraints

- **AUTH-REQ-311:** Authorization decisions may return enforceable obligations.
- **AUTH-REQ-312:** Obligations shall have controlled identifiers.
- **AUTH-REQ-313:** Obligations shall identify responsible enforcement components.
- **AUTH-REQ-314:** Obligations shall identify timing: before, during or after execution.
- **AUTH-REQ-315:** Pre-execution obligations shall be satisfied before operation begins.
- **AUTH-REQ-316:** Continuous obligations shall be monitored during operation.
- **AUTH-REQ-317:** Post-execution obligations shall be recorded and tracked.
- **AUTH-REQ-318:** Failure to satisfy a mandatory obligation shall prevent or stop execution.
- **AUTH-REQ-319:** Obligations may require data minimization.
- **AUTH-REQ-320:** Obligations may require redaction.
- **AUTH-REQ-321:** Obligations may require human notification or approval.
- **AUTH-REQ-322:** Obligations may require enhanced logging.
- **AUTH-REQ-323:** Obligations may require evidence-integrity verification.
- **AUTH-REQ-324:** Obligations may require output quarantine or review.
- **AUTH-REQ-325:** Obligations may require time or volume limits.
- **AUTH-REQ-326:** Obligations may require no-retention or deletion behavior.
- **AUTH-REQ-327:** Obligations may require connector read-only mode.
- **AUTH-REQ-328:** Obligations may require network isolation.
- **AUTH-REQ-329:** Obligations shall not broaden the underlying permit.
- **AUTH-REQ-330:** Unknown mandatory obligations shall cause denial.
- **AUTH-REQ-331:** Conflicting obligations shall cause denial or escalation.
- **AUTH-REQ-332:** Obligation results shall be auditable.
- **AUTH-REQ-333:** Obligations shall trace to policies and requirements.
- **AUTH-REQ-334:** Obligation changes shall be versioned.
- **AUTH-REQ-335:** Obligation enforcement shall be included in validation.

## 16. Denial Reasons and Safe Disclosure

- **AUTH-REQ-336:** Every `Deny` or `Indeterminate` decision shall have a controlled reason code.
- **AUTH-REQ-337:** Reason codes shall distinguish identity, delegation, scope, policy, context, approval, revocation and system failures.
- **AUTH-REQ-338:** Reason codes shall support incident and policy analysis.
- **AUTH-REQ-339:** User-facing denial messages shall reveal only authorized information.
- **AUTH-REQ-340:** Detailed denial evidence shall be restricted to authorized reviewers.
- **AUTH-REQ-341:** Denial evidence shall identify policy version.
- **AUTH-REQ-342:** Denial evidence shall identify relevant attribute freshness without exposing secrets.
- **AUTH-REQ-343:** Denied operations shall not execute partially unless safe rollback is guaranteed and recorded.
- **AUTH-REQ-344:** Denied operations shall not be retried through a weaker policy path.
- **AUTH-REQ-345:** Repeated denials shall be monitored for misuse or misconfiguration.
- **AUTH-REQ-346:** Denials caused by policy conflict shall generate reviewable findings.
- **AUTH-REQ-347:** Denials caused by missing context shall identify required remediation where safe.
- **AUTH-REQ-348:** Denial logs shall preserve attribution.
- **AUTH-REQ-349:** Denial logs shall minimize personal and case data.
- **AUTH-REQ-350:** Denial reason taxonomies shall be versioned.
- **AUTH-REQ-351:** Unknown denial conditions shall map to a safe generic reason.
- **AUTH-REQ-352:** Denial reasons shall not be fabricated.
- **AUTH-REQ-353:** Denial evidence shall remain immutable or tamper-evident where feasible.
- **AUTH-REQ-354:** Denial analytics shall not become excessive surveillance.
- **AUTH-REQ-355:** Denial handling shall be tested.

## 17. Sessions, Tokens and Caching

- **AUTH-REQ-356:** Authorization sessions shall have unique identifiers.
- **AUTH-REQ-357:** Sessions shall be bound to verified subjects and delegation.
- **AUTH-REQ-358:** Sessions shall be bound to case and purpose where applicable.
- **AUTH-REQ-359:** Sessions shall have explicit start and expiry times.
- **AUTH-REQ-360:** Session tokens shall have defined audience and scope.
- **AUTH-REQ-361:** Session tokens shall not contain reusable human credentials.
- **AUTH-REQ-362:** Session tokens shall be integrity-protected.
- **AUTH-REQ-363:** Session tokens shall be revocable where supported.
- **AUTH-REQ-364:** Session fixation shall be prevented.
- **AUTH-REQ-365:** Session renewal shall re-evaluate current authorization context.
- **AUTH-REQ-366:** Privilege changes shall invalidate or re-evaluate affected sessions.
- **AUTH-REQ-367:** Officer-agent rebinding shall invalidate prior sessions.
- **AUTH-REQ-368:** Case closure shall invalidate case-scoped sessions.
- **AUTH-REQ-369:** Purpose completion or expiry shall invalidate purpose-scoped sessions.
- **AUTH-REQ-370:** Connector disablement shall invalidate connector-scoped sessions.
- **AUTH-REQ-371:** Session caching shall not outlive delegation or policy validity.
- **AUTH-REQ-372:** Long-running sessions shall support periodic re-evaluation.
- **AUTH-REQ-373:** Idle and absolute timeout shall be defined where applicable.
- **AUTH-REQ-374:** Session termination shall be auditable.
- **AUTH-REQ-375:** Session data shall be minimized and protected.

## 18. Revocation and Invalidation

- **AUTH-REQ-376:** Authorization shall be revocable at subject, agent, workload, credential, delegation, session, policy, connector, tool, case, purpose and resource levels where applicable.
- **AUTH-REQ-377:** Revocation events shall have immutable identifiers.
- **AUTH-REQ-378:** Revocation events shall identify authority and reason.
- **AUTH-REQ-379:** Revocation shall have an effective time.
- **AUTH-REQ-380:** Revocation shall propagate to decision and enforcement components.
- **AUTH-REQ-381:** Revocation propagation latency shall be defined and tested.
- **AUTH-REQ-382:** Critical revocation shall use immediate or near-immediate containment where technically feasible.
- **AUTH-REQ-383:** Revocation shall invalidate affected cached decisions.
- **AUTH-REQ-384:** Revocation shall invalidate affected sessions and tokens.
- **AUTH-REQ-385:** Revocation shall prevent new affected operations.
- **AUTH-REQ-386:** Revocation shall stop in-progress operations when safe and required.
- **AUTH-REQ-387:** When immediate stop risks evidence corruption, controlled containment shall preserve evidence and prevent further unauthorized action.
- **AUTH-REQ-388:** Revocation shall not erase prior audit or evidence records.
- **AUTH-REQ-389:** Revocation failure shall generate a security incident.
- **AUTH-REQ-390:** Revocation service failure shall result in degraded or denied privileged operation.
- **AUTH-REQ-391:** Revocation conflicts shall fail closed.
- **AUTH-REQ-392:** Revocation rollback shall require separate explicit authority.
- **AUTH-REQ-393:** Expired authority shall not be silently restored.
- **AUTH-REQ-394:** Reinstatement shall require fresh evaluation.
- **AUTH-REQ-395:** Revocation evidence shall trace to affected operations.

## 19. Case, Purpose and Jurisdiction Isolation

- **AUTH-REQ-396:** Authorization shall isolate cases by default.
- **AUTH-REQ-397:** Authorization shall isolate purposes by default.
- **AUTH-REQ-398:** Authorization shall isolate jurisdictions where applicable.
- **AUTH-REQ-399:** Authorization shall isolate officers and their bound agents.
- **AUTH-REQ-400:** Authorization shall isolate workloads where applicable.
- **AUTH-REQ-401:** Cross-case access shall require explicit separate authority.
- **AUTH-REQ-402:** Cross-purpose reuse shall require explicit separate authority.
- **AUTH-REQ-403:** Cross-jurisdiction use shall require explicit legal and governance review where applicable.
- **AUTH-REQ-404:** Case-scoped credentials and sessions shall not be reused across cases.
- **AUTH-REQ-405:** Case-scoped memory shall not be read by unrelated cases.
- **AUTH-REQ-406:** Case-scoped connector data shall not be exposed to unrelated cases.
- **AUTH-REQ-407:** Case-scoped evidence shall not be exported without explicit authority.
- **AUTH-REQ-408:** Aggregate analytics across cases shall require an approved minimized data design.
- **AUTH-REQ-409:** Cross-case policy shall identify necessity, scope, approval and retention.
- **AUTH-REQ-410:** Cross-case denial shall not be bypassed through search, caching or export.
- **AUTH-REQ-411:** Shared infrastructure shall enforce logical isolation.
- **AUTH-REQ-412:** Isolation failure shall trigger containment and incident response.
- **AUTH-REQ-413:** Isolation tests shall include direct and indirect access paths.
- **AUTH-REQ-414:** Isolation telemetry shall detect attempted boundary crossing.
- **AUTH-REQ-415:** Case and purpose identifiers in logs shall be minimized or pseudonymized where feasible.

## 20. Tool and Connector Authorization

- **AUTH-REQ-416:** Tool and connector access shall require explicit authorization.
- **AUTH-REQ-417:** Tool identity shall be validated.
- **AUTH-REQ-418:** Tool scope shall be included in the authorization context.
- **AUTH-REQ-419:** Tool arguments shall be constrained and validated.
- **AUTH-REQ-420:** Tool destinations shall be authorized.
- **AUTH-REQ-421:** Tool data inputs shall be minimized.
- **AUTH-REQ-422:** Tool outputs shall be treated as untrusted.
- **AUTH-REQ-423:** Tool execution shall preserve provenance.
- **AUTH-REQ-424:** Tool execution shall not inherit unrestricted model context.
- **AUTH-REQ-425:** Model requests shall not bypass tool authorization.
- **AUTH-REQ-426:** Connector retries shall not bypass denial.
- **AUTH-REQ-427:** Read-only mode shall be preferred when sufficient.
- **AUTH-REQ-428:** High-risk tools shall require human approval or step-up controls.
- **AUTH-REQ-429:** Tool credentials shall be independently revocable.
- **AUTH-REQ-430:** Tool failure shall not trigger broader fallback access.
- **AUTH-REQ-431:** Unknown tools shall be denied.
- **AUTH-REQ-432:** Unapproved destinations shall be denied.
- **AUTH-REQ-433:** Operationally isolated Dark Web research shall not interact with uncontrolled criminal infrastructure.
- **AUTH-REQ-434:** Tool use shall comply with `20_CONNECTOR_SECURITY_POLICY.md` after consolidation.
- **AUTH-REQ-435:** Tool authorization tests shall include malformed, malicious and indirect prompt-injection cases.

## 21. Data and Evidence Authorization

- **AUTH-REQ-436:** Data access shall be authorized at a granularity proportionate to risk.
- **AUTH-REQ-437:** Data scope shall identify permitted categories, fields, records or collections where applicable.
- **AUTH-REQ-438:** Data classification shall constrain authorization.
- **AUTH-REQ-439:** Personal data access shall require approved purpose and minimization.
- **AUTH-REQ-440:** Evidence acquisition shall require explicit source and scope authority.
- **AUTH-REQ-441:** Evidence viewing shall be distinct from evidence modification.
- **AUTH-REQ-442:** Evidence transformation shall be separately authorized.
- **AUTH-REQ-443:** Evidence deletion shall require exceptional explicit authority.
- **AUTH-REQ-444:** Evidence export shall be separately authorized.
- **AUTH-REQ-445:** Evidence sharing shall be separately authorized.
- **AUTH-REQ-446:** Derived analytical data shall retain lineage and authorization context.
- **AUTH-REQ-447:** AI-generated analysis shall not inherit source-evidence authority.
- **AUTH-REQ-448:** Unauthorized data fields shall be redacted or withheld.
- **AUTH-REQ-449:** Query expansion shall not exceed authorized data scope.
- **AUTH-REQ-450:** Search results shall not disclose unauthorized existence where policy requires concealment.
- **AUTH-REQ-451:** Bulk access shall require stronger controls than individual access.
- **AUTH-REQ-452:** Data retention obligations shall be enforced.
- **AUTH-REQ-453:** Data deletion obligations shall be enforced.
- **AUTH-REQ-454:** Evidence authorization shall integrate with `18_EVIDENCE_MODEL.md` after consolidation.
- **AUTH-REQ-455:** Data and evidence authorization tests shall verify field-, record-, case- and export-level controls.

## 22. Administrative and Emergency Authorization

- **AUTH-REQ-456:** Administrative operations shall use distinct administrative authorization.
- **AUTH-REQ-457:** Administrative privileges shall be time-bounded where feasible.
- **AUTH-REQ-458:** Administrative access shall require strong authentication.
- **AUTH-REQ-459:** Administrative access shall require explicit purpose.
- **AUTH-REQ-460:** Administrative actions shall be fully attributable.
- **AUTH-REQ-461:** Policy publication shall require administrative authorization.
- **AUTH-REQ-462:** Identity issuance and revocation shall require administrative authorization.
- **AUTH-REQ-463:** Connector onboarding and enablement shall require administrative authorization.
- **AUTH-REQ-464:** Audit configuration changes shall require administrative authorization.
- **AUTH-REQ-465:** Evidence-retention configuration changes shall require administrative authorization.
- **AUTH-REQ-466:** Administrative actions shall not be performed through ordinary agent delegation.
- **AUTH-REQ-467:** Administrative interfaces shall be isolated and monitored.
- **AUTH-REQ-468:** Administrative sessions shall have shorter validity where appropriate.
- **AUTH-REQ-469:** Administrative exports shall be minimized and controlled.
- **AUTH-REQ-470:** Administrative denial shall fail closed.
- **AUTH-REQ-471:** Break-glass access shall be exceptional, narrow, time-bounded and auditable.
- **AUTH-REQ-472:** Break-glass access shall require attributable human authority.
- **AUTH-REQ-473:** Break-glass access shall not authorize canonical prohibitions.
- **AUTH-REQ-474:** Break-glass use shall trigger immediate review.
- **AUTH-REQ-475:** Repeated break-glass use shall trigger governance and architectural review.

## 23. Authorization Audit Schema

- **AUTH-REQ-476:** Every authorization decision shall produce an audit event.
- **AUTH-REQ-477:** Every enforcement result shall produce or correlate to an audit event.
- **AUTH-REQ-478:** Audit events shall have immutable identifiers.
- **AUTH-REQ-479:** Audit events shall identify request identifier.
- **AUTH-REQ-480:** Audit events shall identify subject identities.
- **AUTH-REQ-481:** Audit events shall identify institution.
- **AUTH-REQ-482:** Audit events shall identify case and purpose where applicable.
- **AUTH-REQ-483:** Audit events shall identify resource and action.
- **AUTH-REQ-484:** Audit events shall identify decision outcome.
- **AUTH-REQ-485:** Audit events shall identify reason code.
- **AUTH-REQ-486:** Audit events shall identify policy bundle and version.
- **AUTH-REQ-487:** Audit events shall identify material attribute sources and freshness.
- **AUTH-REQ-488:** Audit events shall identify obligations.
- **AUTH-REQ-489:** Audit events shall identify human approval references where applicable.
- **AUTH-REQ-490:** Audit events shall identify enforcement result.
- **AUTH-REQ-491:** Audit events shall identify time and ordering information.
- **AUTH-REQ-492:** Audit events shall identify revocation state where material.
- **AUTH-REQ-493:** Audit events shall identify errors or degraded conditions.
- **AUTH-REQ-494:** Audit events shall not contain secrets.
- **AUTH-REQ-495:** Audit events shall minimize unnecessary personal data.
- **AUTH-REQ-496:** Audit events shall be append-only or tamper-evident where feasible.
- **AUTH-REQ-497:** Audit events shall be retained according to governance and evidence requirements.
- **AUTH-REQ-498:** Audit events shall support correlation across decision, enforcement and tool execution.
- **AUTH-REQ-499:** Audit failure shall not permit silent privileged operation.
- **AUTH-REQ-500:** Detailed evidence-event integration remains a forward dependency on `18_EVIDENCE_MODEL.md`.

## 24. Failure Modes and Fail-Closed Behavior

- **AUTH-REQ-501:** Policy decision service unavailability shall not result in `Permit`.
- **AUTH-REQ-502:** Policy information service unavailability shall not result in `Permit` when missing attributes are material.
- **AUTH-REQ-503:** Revocation service unavailability shall restrict or deny privileged operations.
- **AUTH-REQ-504:** Clock uncertainty beyond approved tolerance shall deny time-sensitive operations.
- **AUTH-REQ-505:** Policy integrity failure shall deny affected operations.
- **AUTH-REQ-506:** Attribute integrity failure shall deny affected operations.
- **AUTH-REQ-507:** Communication integrity failure shall deny affected operations.
- **AUTH-REQ-508:** Enforcement-point failure shall prevent protected execution.
- **AUTH-REQ-509:** Obligation-enforcement failure shall prevent or stop execution.
- **AUTH-REQ-510:** Human-approval service failure shall deny approval-dependent operations.
- **AUTH-REQ-511:** Audit-service failure shall block or degrade privileged operation according to approved policy.
- **AUTH-REQ-512:** Cache inconsistency shall trigger invalidation and safe re-evaluation.
- **AUTH-REQ-513:** Multiple decision points returning conflicting outcomes shall result in denial.
- **AUTH-REQ-514:** Unknown decision outcomes shall result in denial.
- **AUTH-REQ-515:** Malformed policies shall not be activated.
- **AUTH-REQ-516:** Partial execution after denial shall trigger containment and incident response.
- **AUTH-REQ-517:** Failure states shall be observable.
- **AUTH-REQ-518:** Failure states shall identify owners and recovery procedures.
- **AUTH-REQ-519:** Recovery shall not restore stale or revoked authority.
- **AUTH-REQ-520:** Post-recovery authorization validation shall be required.
- **AUTH-REQ-521:** Repeated failures shall trigger architectural and risk review.
- **AUTH-REQ-522:** Failure handling shall preserve evidence.
- **AUTH-REQ-523:** Fallback paths shall be explicitly authorized and no less restrictive.
- **AUTH-REQ-524:** Emergency bypass shall not be implemented as a hidden permissive fallback.
- **AUTH-REQ-525:** Failure-path behavior shall be tested.

## 25. Privacy, Ethics and Fundamental Rights

- **AUTH-REQ-526:** Authorization shall enforce purpose limitation.
- **AUTH-REQ-527:** Authorization shall enforce data minimization.
- **AUTH-REQ-528:** Authorization shall enforce case separation.
- **AUTH-REQ-529:** Authorization shall enforce retention and deletion obligations where applicable.
- **AUTH-REQ-530:** Authorization policies shall avoid discriminatory or irrelevant attributes.
- **AUTH-REQ-531:** Sensitive attributes shall be used only when necessary and authorized.
- **AUTH-REQ-532:** Policy design shall identify fundamental-rights impacts where material.
- **AUTH-REQ-533:** Automated authorization shall not become an autonomous legal or coercive decision.
- **AUTH-REQ-534:** Human review shall remain available for high-impact or contested decisions where required.
- **AUTH-REQ-535:** Denial explanations shall be proportionate and privacy-preserving.
- **AUTH-REQ-536:** Authorization telemetry shall not become excessive surveillance.
- **AUTH-REQ-537:** Authorization analytics shall use minimized or pseudonymized data where feasible.
- **AUTH-REQ-538:** Public-source accessibility shall not be treated as unrestricted authorization.
- **AUTH-REQ-539:** Cross-jurisdiction access shall identify applicable authority and limitations.
- **AUTH-REQ-540:** Authorization shall not enable unlawful deanonymization.
- **AUTH-REQ-541:** Authorization shall not enable real-person public investigations in demonstrations.
- **AUTH-REQ-542:** Privacy incidents shall trigger authorization-policy reassessment.
- **AUTH-REQ-543:** Rights-impact findings shall have owners and dispositions.
- **AUTH-REQ-544:** Compliance mappings shall not be treated as authorization.
- **AUTH-REQ-545:** Document 30 remains a forward dependency for complete compliance mapping.

## 26. Monitoring and Incident Integration

- **AUTH-REQ-546:** Authorization success and denial rates shall be monitorable.
- **AUTH-REQ-547:** Repeated denied requests shall be detectable.
- **AUTH-REQ-548:** Privilege-elevation attempts shall be detectable.
- **AUTH-REQ-549:** Cross-case access attempts shall be detectable.
- **AUTH-REQ-550:** Use of expired or revoked authority shall be detectable.
- **AUTH-REQ-551:** Policy-version divergence shall be detectable.
- **AUTH-REQ-552:** Decision-enforcement mismatches shall be detectable.
- **AUTH-REQ-553:** Obligation failures shall be detectable.
- **AUTH-REQ-554:** Human-approval anomalies shall be detectable.
- **AUTH-REQ-555:** Administrative and break-glass use shall generate alerts.
- **AUTH-REQ-556:** Monitoring shall distinguish malicious use from expected negative testing.
- **AUTH-REQ-557:** Alert rules shall be versioned.
- **AUTH-REQ-558:** Alerts shall identify owner and response path.
- **AUTH-REQ-559:** Monitoring blind spots shall be documented.
- **AUTH-REQ-560:** Monitoring failure shall create a degraded state.
- **AUTH-REQ-561:** Monitoring data shall be protected from alteration.
- **AUTH-REQ-562:** Monitoring access shall be authorized.
- **AUTH-REQ-563:** Monitoring retention shall be governed.
- **AUTH-REQ-564:** Monitoring shall minimize personal and case data.
- **AUTH-REQ-565:** Authorization incidents shall trigger threat, risk and policy reassessment.

## 27. Authorization Testing and Validation

- **AUTH-REQ-566:** Every material authorization requirement shall have testable acceptance criteria.
- **AUTH-REQ-567:** Authorization testing shall use synthetic or authorized controlled identities and data.
- **AUTH-REQ-568:** Tests shall cover `Permit` outcomes.
- **AUTH-REQ-569:** Tests shall cover explicit `Deny` outcomes.
- **AUTH-REQ-570:** Tests shall cover `Not Applicable` outcomes.
- **AUTH-REQ-571:** Tests shall cover `Indeterminate` outcomes.
- **AUTH-REQ-572:** Tests shall verify default deny.
- **AUTH-REQ-573:** Tests shall verify missing-attribute denial.
- **AUTH-REQ-574:** Tests shall verify stale-attribute denial.
- **AUTH-REQ-575:** Tests shall verify conflicting-attribute denial.
- **AUTH-REQ-576:** Tests shall verify expired credential denial.
- **AUTH-REQ-577:** Tests shall verify revoked subject denial.
- **AUTH-REQ-578:** Tests shall verify revoked delegation denial.
- **AUTH-REQ-579:** Tests shall verify session invalidation.
- **AUTH-REQ-580:** Tests shall verify policy rollback.
- **AUTH-REQ-581:** Tests shall verify deny-overrides conflict resolution.
- **AUTH-REQ-582:** Tests shall verify canonical prohibitions cannot be permitted.
- **AUTH-REQ-583:** Tests shall verify human-approval requirements.
- **AUTH-REQ-584:** Tests shall verify approval expiry and non-reuse.
- **AUTH-REQ-585:** Tests shall verify separation of duties.
- **AUTH-REQ-586:** Tests shall verify obligation enforcement.
- **AUTH-REQ-587:** Tests shall verify cross-case isolation.
- **AUTH-REQ-588:** Tests shall verify cross-purpose isolation.
- **AUTH-REQ-589:** Tests shall verify tool and connector constraints.
- **AUTH-REQ-590:** Tests shall verify data and evidence scope.
- **AUTH-REQ-591:** Tests shall verify long-running operation re-evaluation.
- **AUTH-REQ-592:** Tests shall verify policy-service failure behavior.
- **AUTH-REQ-593:** Tests shall verify revocation-service failure behavior.
- **AUTH-REQ-594:** Tests shall verify audit-service failure behavior.
- **AUTH-REQ-595:** Tests shall verify denial reason handling.
- **AUTH-REQ-596:** Tests shall identify exact identity, policy, configuration, code and environment versions.
- **AUTH-REQ-597:** Failed tests shall remain visible.
- **AUTH-REQ-598:** Passing tests shall not be generalized beyond scope.
- **AUTH-REQ-599:** Automated tests shall not replace human authorization review.
- **AUTH-REQ-600:** Document 22 remains a forward dependency for complete test governance.

## 28. Traceability

- **AUTH-REQ-601:** Every authorization requirement shall trace to governing authority or threat.
- **AUTH-REQ-602:** Every policy rule shall trace to authorization requirements.
- **AUTH-REQ-603:** Every delegation field shall trace to identity and authorization requirements.
- **AUTH-REQ-604:** Every enforcement control shall trace to policy decisions and obligations.
- **AUTH-REQ-605:** Every test shall trace to requirements and threats.
- **AUTH-REQ-606:** Every decision record shall trace to policy and context versions.
- **AUTH-REQ-607:** Every human approval shall trace to the protected operation.
- **AUTH-REQ-608:** Every revocation shall trace to affected authority and operations.
- **AUTH-REQ-609:** Every exception shall trace to compensating controls.
- **AUTH-REQ-610:** Every incident shall trace to affected policy, decisions and enforcement.
- **AUTH-REQ-611:** Traceability shall be bidirectional.
- **AUTH-REQ-612:** Broken traceability affecting authority, identity, case, purpose, evidence or release shall be blocking.
- **AUTH-REQ-613:** Planned policies shall not be represented as active policies.
- **AUTH-REQ-614:** Active policies shall not be represented as validated without evidence.
- **AUTH-REQ-615:** Superseded policies shall identify successors.
- **AUTH-REQ-616:** Deprecated actions and attributes shall identify replacements.
- **AUTH-REQ-617:** Traceability shall use immutable identifiers.
- **AUTH-REQ-618:** Traceability records shall not contain secrets.
- **AUTH-REQ-619:** Traceability records shall minimize personal and case data.
- **AUTH-REQ-620:** Baseline freeze shall validate authorization traceability across documents 01–30.

## 29. Authorization Review Gate

- **AUTH-REQ-621:** Authorization review shall identify the exact document and commit.
- **AUTH-REQ-622:** Authorization review shall verify all original context fields are preserved.
- **AUTH-REQ-623:** Authorization review shall verify default deny.
- **AUTH-REQ-624:** Authorization review shall verify explicit, revocable and auditable behavior.
- **AUTH-REQ-625:** Authorization review shall verify object and component responsibilities.
- **AUTH-REQ-626:** Authorization review shall verify evaluation order.
- **AUTH-REQ-627:** Authorization review shall verify human-approval triggers.
- **AUTH-REQ-628:** Authorization review shall verify separation of duties.
- **AUTH-REQ-629:** Authorization review shall verify revocation and session invalidation.
- **AUTH-REQ-630:** Authorization review shall verify denial reasons and obligations.
- **AUTH-REQ-631:** Authorization review shall verify cross-case and cross-purpose isolation.
- **AUTH-REQ-632:** Authorization review shall verify failure handling.
- **AUTH-REQ-633:** Authorization review shall verify audit evidence.
- **AUTH-REQ-634:** Authorization review shall verify test coverage.
- **AUTH-REQ-635:** Authorization review shall identify residual risks and forward dependencies.
- **AUTH-REQ-636:** Critical findings shall block progression.
- **AUTH-REQ-637:** Material post-review change shall invalidate affected evidence.
- **AUTH-REQ-638:** Role concentration shall be disclosed.
- **AUTH-REQ-639:** Internal review shall not be represented as independent external assurance.
- **AUTH-REQ-640:** Approval for Draft incorporation shall not establish normative approval.

## 30. Minimum Validation Checklist

Before approval of this model or an implementation, confirm:

- [ ] Officer is identified;
- [ ] Institution is identified;
- [ ] Agent and workload identities are identified where applicable;
- [ ] Case or mandate is identified;
- [ ] Purpose is identified;
- [ ] Jurisdiction is identified where applicable;
- [ ] Time validity is current;
- [ ] Approved tools and connectors are identified;
- [ ] Data scope is explicit;
- [ ] Policy compliance is evaluated;
- [ ] delegation is valid and within issuer authority;
- [ ] resource and action identifiers are canonical;
- [ ] policy bundle and version are fixed;
- [ ] canonical prohibitions are evaluated first;
- [ ] default decision is Deny;
- [ ] human approval and step-up requirements are enforced;
- [ ] obligations are enforceable;
- [ ] denial reasons are controlled;
- [ ] sessions and caches respect expiry and revocation;
- [ ] cross-case and cross-purpose isolation is tested;
- [ ] audit evidence correlates request, decision and enforcement;
- [ ] failure paths are fail-closed;
- [ ] privacy and rights impacts are reviewed;
- [ ] negative and revocation tests pass;
- [ ] forward dependencies 18, 19, 20, 22, 26 and 30 are recorded;
- [ ] Project Founder approval exists before status becomes Approved.


## 31. Limitations

- This model does not select or implement a policy engine, identity provider, policy language, token format or approval platform.
- It does not establish legal authority, institutional deployment, production readiness, certification or compliance.
- It does not replace the Identity and Delegation Model.
- It does not replace the Trust Model.
- It does not replace the Security Architecture Baseline.
- It does not define the complete Evidence Model, Agent Lifecycle Model or Connector Security Policy.
- Default deny reduces but does not eliminate policy-design or enforcement errors.
- Attribute-based decisions depend on the integrity and freshness of their sources.
- Internal review is not independent certification.
- Documents 18, 19, 20, 22, 26 and 30 remain forward dependencies for full approval.


## 32. Change Control

Every material change shall identify rationale, affected subjects, resources, actions, policies, attributes, obligations, dependencies, security and privacy impact, evidence impact, migration, validation, rollback, authority and version effect.

- **AUTH-REQ-641:** Editorial corrections shall use a patch version when meaning is unchanged.
- **AUTH-REQ-642:** Backward-compatible substantive additions shall use a minor version.
- **AUTH-REQ-643:** Incompatible authorization-semantics changes shall use a major version.
- **AUTH-REQ-644:** Material authorization architecture changes shall require an ADR where applicable.
- **AUTH-REQ-645:** Changes shall require Project Founder approval.
- **AUTH-REQ-646:** Changes shall receive Security Reviewer assessment.
- **AUTH-REQ-647:** Privacy, evidence, lifecycle and connector impacts shall receive specialist review where applicable.
- **AUTH-REQ-648:** Changes shall identify affected policies, attributes, obligations, tests and implementation.
- **AUTH-REQ-649:** Changes shall include migration, validation and rollback analysis.
- **AUTH-REQ-650:** Changes shall not retroactively fabricate authorization or approval evidence.
- **AUTH-REQ-651:** Historical decisions and audit evidence shall not be silently rewritten.
- **AUTH-REQ-652:** Forward-dependency reconciliation shall occur before this model becomes Approved.

## 33. Consolidation Record

Version 1.0.0 consolidates the two existing Authorization Model variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-AUTH-001`;
- normalizes the authoritative filename to `17_AUTHORIZATION_MODEL.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves Officer, Institution, Case, Purpose, Jurisdiction, Time, Approved Tools, Data Scope and Policy Compliance as mandatory decision inputs;
- preserves explicit, revocable and auditable authorization;
- preserves `Deny` as the default decision;
- adds the authorization object, controlled actions, subject and resource models, delegation validation, PDP/PIP/PAP/PEP responsibilities, deterministic evaluation order, outcomes, policy lifecycle, human approval, separation of duties, obligations, denial reasons, sessions, revocation, cross-case isolation, tool and evidence authorization, audit schema, failure modes, monitoring, testing and traceability;
- identifies documents 18, 19, 20, 22, 26 and 30 as forward dependencies blocking normative approval;
- treats Enterprise and non-Enterprise legacy files as source variants of the same immutable document;
- creates no implementation, policy engine selection, legal authority or operational authorization.

## 34. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Security Reviewer; approval reserved to Project Founder | Constitutional consolidation of both legacy authorization outlines: preserved all decision inputs, explicit/revocable/auditable behavior and default deny; added complete authorization object, components, evaluation, enforcement, approval, revocation, audit, failure, testing and traceability requirements. |
| 1.0.1 | 2026-08-06 | Draft | Security Reviewer; approval reserved to Project Founder | Reconciled obsolete forward-dependency filenames with the final canonical Knowledge Pack 01–30 filenames; no requirement text, authority, lifecycle status, implementation state or validation claim changed. |
