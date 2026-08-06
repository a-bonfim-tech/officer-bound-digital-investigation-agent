# AGENT LIFECYCLE MODEL

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-AGENT-001 |
| **Title** | Agent Lifecycle Model |
| **Version** | 1.0.1 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Security Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define the authoritative lifecycle states, transition criteria, invariants, responsible roles, evidence, failure behavior, suspension, revocation, archival and validation requirements for an institutionally issued OBDIA agent. |
| **Scope** | Agent identity, officer and institutional binding, delegation, authorization, activation, runtime operation, periodic and event-driven review, suspension, revocation, archival, related workload and session handling, connector dependencies, evidence and audit records within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `06_IDENTITY_AND_DELEGATION_MODEL.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; `18_EVIDENCE_MODEL.md`; forward dependencies `20_CONNECTOR_SECURITY_POLICY.md`, `21_SECURE_CODING_STANDARD.md`, `22_TESTING_STANDARD.md`, `24_GITHUB_REPOSITORY_STANDARD.md`, `25_DOCUMENT_VERSIONING_POLICY.md`, `26_AI_RISK_REGISTER.md`, `27_RESEARCH_BACKLOG.md`, `28_ASSUMPTIONS_REGISTER.md`, `29_DECISION_LOG_POLICY.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-ARCH-001`; `OBDIA-ID-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001`; `OBDIA-EVID-001` |
| **Cross-References** | Documents 20–30; accepted ADRs; identity, delegation, authorization, connector, workload, session, evidence, audit, risk, exception, incident, change, validation and release records |
| **Assumptions** | The lifecycle is demonstrated through synthetic identities, mock institutions, local laboratories, test environments and non-operational research configurations. |
| **Constraints** | An agent shall remain bound to exactly one identified human officer and one accountable institution; shall not receive independent legal authority; shall not transfer human credentials; shall not authorize prohibited conduct; and shall not be rebound to another officer. |
| **Security Considerations** | Lifecycle failure can cause orphaned agents, stale authority, unauthorized activation, credential misuse, loss of officer attribution, cross-case contamination, uncontrolled connector access, evidence corruption, incomplete revocation, audit gaps and misleading operational status. |
| **Validation Criteria** | Every lifecycle state and transition has explicit entry and exit criteria, authorized roles, exact identity and policy references, logged transition evidence, secure failure behavior, revocation propagation, negative tests and bidirectional traceability. |
| **Implementation Relationship** | This model defines lifecycle semantics and controls. It does not provision a real agent, create credentials, select an identity platform, authorize investigative activity, deploy a runtime or establish production readiness. |

---

## 1. Purpose

This document defines the lifecycle of an institutionally issued Officer-Bound Digital Investigation Agent.

The two legacy sources defined these sequences:

- `Provision → Authorize → Activate → Operate → Review → Suspend → Revoke → Archive`
- `Provision → Bind → Authorize → Activate → Operate → Review → Suspend → Revoke → Archive`

The consolidated authoritative sequence is:

`Provision → Bind → Authorize → Activate → Operate → Review → Suspend → Revoke → Archive`

`Bind` is mandatory because the project requires cryptographic binding to exactly one identified human officer and one accountable institution.

The sequence describes governed lifecycle states, not a mandatory straight-line execution path. Approved transitions include review loops, emergency suspension and terminal revocation. Every transition shall be logged and auditable.


## 2. Fundamental Lifecycle Rules

- **AGENT-REQ-001:** The authoritative agent lifecycle states shall be `Provision`, `Bind`, `Authorize`, `Activate`, `Operate`, `Review`, `Suspend`, `Revoke` and `Archive`.
- **AGENT-REQ-002:** Every lifecycle transition shall be logged.
- **AGENT-REQ-003:** Every lifecycle transition shall be auditable.
- **AGENT-REQ-004:** Every lifecycle transition shall identify an authorized human or institutional authority.
- **AGENT-REQ-005:** An AI agent shall not approve its own lifecycle transition.
- **AGENT-REQ-006:** An AI agent shall not create, broaden or restore its own authority.
- **AGENT-REQ-007:** An agent shall remain bound to exactly one identified human officer at a time.
- **AGENT-REQ-008:** An agent shall remain bound to one accountable institution.
- **AGENT-REQ-009:** Rebinding an existing agent identity to another officer shall be prohibited.
- **AGENT-REQ-010:** Changing the responsible officer shall require revocation and archival of the prior identity and provisioning of a new identity.
- **AGENT-REQ-011:** Human credentials shall not be transferred to or embedded in the agent.
- **AGENT-REQ-012:** Lifecycle state shall not create independent legal authority.
- **AGENT-REQ-013:** Lifecycle state shall not waive case, purpose, jurisdiction, time, tool or data constraints.
- **AGENT-REQ-014:** Canonical prohibitions shall remain non-authorizable in every state.
- **AGENT-REQ-015:** Missing, stale, conflicting or unverifiable lifecycle context shall result in denial, suspension or containment.
- **AGENT-REQ-016:** Lifecycle transition controls shall fail closed.
- **AGENT-REQ-017:** Lifecycle state shall remain distinct from repository document lifecycle state.
- **AGENT-REQ-018:** Lifecycle state shall remain distinct from session state, workload state, connector state and evidence handling state.
- **AGENT-REQ-019:** Repository merge, branch state, tag or deployment shall not establish an agent lifecycle state.
- **AGENT-REQ-020:** Automated checks may provide evidence but shall not establish final lifecycle authority.
- **AGENT-REQ-021:** Every active state shall have a responsible human owner and institutional custodian.
- **AGENT-REQ-022:** Every state shall have explicit entry and exit criteria.
- **AGENT-REQ-023:** Every state shall have secure failure and containment behavior.
- **AGENT-REQ-024:** Every material transition shall identify the exact agent identity and lifecycle version.
- **AGENT-REQ-025:** Historical transition records shall not be silently overwritten.

## 3. Lifecycle Object

- **AGENT-REQ-026:** Every agent lifecycle instance shall have an immutable lifecycle identifier.
- **AGENT-REQ-027:** Every lifecycle instance shall identify the agent identity.
- **AGENT-REQ-028:** Every lifecycle instance shall identify the bound human officer.
- **AGENT-REQ-029:** Every lifecycle instance shall identify the accountable institution.
- **AGENT-REQ-030:** Every lifecycle instance shall identify the agent identity issuer.
- **AGENT-REQ-031:** Every lifecycle instance shall identify the cryptographic-binding reference.
- **AGENT-REQ-032:** Every lifecycle instance shall identify the current lifecycle state.
- **AGENT-REQ-033:** Every lifecycle instance shall identify the prior lifecycle state.
- **AGENT-REQ-034:** Every lifecycle instance shall identify state-effective time.
- **AGENT-REQ-035:** Every lifecycle instance shall identify the lifecycle-policy version.
- **AGENT-REQ-036:** Every lifecycle instance shall identify applicable delegation references.
- **AGENT-REQ-037:** Every lifecycle instance shall identify applicable authorization-policy references.
- **AGENT-REQ-038:** Every lifecycle instance shall identify approved case or mandate scope where applicable.
- **AGENT-REQ-039:** Every lifecycle instance shall identify approved purpose scope.
- **AGENT-REQ-040:** Every lifecycle instance shall identify jurisdictional constraints where applicable.
- **AGENT-REQ-041:** Every lifecycle instance shall identify approved tools and connector classes.
- **AGENT-REQ-042:** Every lifecycle instance shall identify approved data scope.
- **AGENT-REQ-043:** Every lifecycle instance shall identify time validity and expiry.
- **AGENT-REQ-044:** Every lifecycle instance shall identify current credential and key references without containing secrets.
- **AGENT-REQ-045:** Every lifecycle instance shall identify workload and runtime references where applicable.
- **AGENT-REQ-046:** Every lifecycle instance shall identify active session references where applicable.
- **AGENT-REQ-047:** Every lifecycle instance shall identify current trust and risk state.
- **AGENT-REQ-048:** Every lifecycle instance shall identify unresolved findings and restrictions.
- **AGENT-REQ-049:** Every lifecycle instance shall identify transition authority.
- **AGENT-REQ-050:** Every lifecycle instance shall identify transition evidence.
- **AGENT-REQ-051:** Every lifecycle instance shall identify suspension, revocation or archive references where applicable.
- **AGENT-REQ-052:** Every lifecycle instance shall identify retention requirements.
- **AGENT-REQ-053:** Every lifecycle instance shall identify the current schema version.
- **AGENT-REQ-054:** Duplicate lifecycle identifiers shall be rejected.
- **AGENT-REQ-055:** Malformed lifecycle objects shall be rejected or quarantined.
- **AGENT-REQ-056:** Lifecycle metadata shall not contain reusable credentials.
- **AGENT-REQ-057:** Lifecycle metadata shall minimize personal and case data.
- **AGENT-REQ-058:** Lifecycle objects shall be integrity-protected where material.
- **AGENT-REQ-059:** Lifecycle objects shall remain traceable to all state-transition events.
- **AGENT-REQ-060:** Lifecycle objects shall remain retrievable after revocation and archival.

## 4. Lifecycle Invariants

- **AGENT-REQ-061:** An agent shall not be authorized before successful binding.
- **AGENT-REQ-062:** An agent shall not be activated before successful authorization.
- **AGENT-REQ-063:** An agent shall not operate before successful activation.
- **AGENT-REQ-064:** An agent shall not operate with expired authorization.
- **AGENT-REQ-065:** An agent shall not operate with revoked credentials.
- **AGENT-REQ-066:** An agent shall not operate with suspended identity or delegation.
- **AGENT-REQ-067:** An agent shall not operate when officer binding cannot be verified.
- **AGENT-REQ-068:** An agent shall not operate when institutional binding cannot be verified.
- **AGENT-REQ-069:** An agent shall not operate outside approved case or mandate scope.
- **AGENT-REQ-070:** An agent shall not operate outside approved purpose.
- **AGENT-REQ-071:** An agent shall not operate outside applicable jurisdictional constraints.
- **AGENT-REQ-072:** An agent shall not use unapproved tools or connectors.
- **AGENT-REQ-073:** An agent shall not access data outside approved scope.
- **AGENT-REQ-074:** An agent shall not continue privileged operation after material revocation.
- **AGENT-REQ-075:** An agent shall not create or approve another agent identity autonomously.
- **AGENT-REQ-076:** An agent shall not transfer authority to another agent autonomously.
- **AGENT-REQ-077:** An agent shall not retain human authentication secrets.
- **AGENT-REQ-078:** An agent shall not silently change model, policy, connector or runtime identity.
- **AGENT-REQ-079:** An agent shall not represent its output as original evidence.
- **AGENT-REQ-080:** An agent shall not perform autonomous legal, coercive or policing decisions.
- **AGENT-REQ-081:** An agent shall preserve officer and institutional attribution for every protected operation.
- **AGENT-REQ-082:** An agent shall preserve auditability in every active state.
- **AGENT-REQ-083:** An agent shall preserve evidence provenance and chain of custody where evidence is handled.
- **AGENT-REQ-084:** An agent shall fail closed when material authorization or lifecycle evidence is unavailable.
- **AGENT-REQ-085:** An agent shall support suspension and revocation without erasing historical evidence.
- **AGENT-REQ-086:** An agent shall not resume from suspension without review and fresh authorization.
- **AGENT-REQ-087:** A revoked agent identity shall not return to an active state.
- **AGENT-REQ-088:** An archived agent identity shall not return to an active state.
- **AGENT-REQ-089:** Lifecycle invariants shall be evaluated continuously or at risk-appropriate checkpoints.
- **AGENT-REQ-090:** Invariant violation shall trigger denial, suspension, revocation or incident handling.

## 5. Transition Event Model

- **AGENT-REQ-091:** Every lifecycle transition shall create an immutable transition-event identifier.
- **AGENT-REQ-092:** Every transition event shall identify the lifecycle identifier.
- **AGENT-REQ-093:** Every transition event shall identify the agent identity.
- **AGENT-REQ-094:** Every transition event shall identify prior state.
- **AGENT-REQ-095:** Every transition event shall identify requested state.
- **AGENT-REQ-096:** Every transition event shall identify final state.
- **AGENT-REQ-097:** Every transition event shall identify transition time.
- **AGENT-REQ-098:** Every transition event shall identify requesting actor.
- **AGENT-REQ-099:** Every transition event shall identify approving authority.
- **AGENT-REQ-100:** Every transition event shall identify accountable institution.
- **AGENT-REQ-101:** Every transition event shall identify bound officer.
- **AGENT-REQ-102:** Every transition event shall identify reason.
- **AGENT-REQ-103:** Every transition event shall identify applicable policy and version.
- **AGENT-REQ-104:** Every transition event shall identify delegation and authorization references.
- **AGENT-REQ-105:** Every transition event shall identify evidence used to evaluate entry criteria.
- **AGENT-REQ-106:** Every transition event shall identify required human approvals.
- **AGENT-REQ-107:** Every transition event shall identify obligations and conditions.
- **AGENT-REQ-108:** Every transition event shall identify expiry where applicable.
- **AGENT-REQ-109:** Every transition event shall identify transition result.
- **AGENT-REQ-110:** Every transition event shall identify errors and degraded conditions.
- **AGENT-REQ-111:** Every transition event shall identify related credential, session, workload and connector effects.
- **AGENT-REQ-112:** Every transition event shall identify related evidence and audit records.
- **AGENT-REQ-113:** Every transition event shall identify risk or exception references where applicable.
- **AGENT-REQ-114:** Every transition event shall identify rollback or containment references where applicable.
- **AGENT-REQ-115:** Every transition event shall be integrity-protected where material.
- **AGENT-REQ-116:** Transition events shall be append-only or tamper-evident where feasible.
- **AGENT-REQ-117:** Transition-event corrections shall be additive.
- **AGENT-REQ-118:** Failed transitions shall remain visible.
- **AGENT-REQ-119:** Retries shall create new attributable transition events.
- **AGENT-REQ-120:** Transition evidence shall identify the exact evaluated agent and policy versions.
- **AGENT-REQ-121:** Unknown transition outcomes shall resolve to the safer prior or contained state.
- **AGENT-REQ-122:** Transition records shall not expose secrets.
- **AGENT-REQ-123:** Transition records shall minimize unnecessary personal data.
- **AGENT-REQ-124:** Transition events shall support bidirectional traceability.
- **AGENT-REQ-125:** Transition events shall be retained after archival.

## 6. Permitted Transition Matrix

| From | To | Normal or Exceptional | Minimum condition |
|---|---|---|---|
| None | Provision | Normal | Approved issuance request and institutional authority |
| Provision | Bind | Normal | Verified officer and institution; cryptographic binding established |
| Bind | Authorize | Normal | Valid delegation and authorization context |
| Authorize | Activate | Normal | Authorization Permit; required obligations satisfied |
| Activate | Operate | Normal | Runtime, policy, monitoring and connector readiness verified |
| Operate | Review | Normal | Scheduled, event-driven or requested review |
| Review | Operate | Normal | Review passes; authorization remains valid |
| Review | Authorize | Normal | Material scope or policy change requires fresh authorization |
| Operate | Suspend | Exceptional or planned | Containment, expiry, review finding or administrative decision |
| Review | Suspend | Normal or exceptional | Blocking finding, expired condition or unresolved risk |
| Suspend | Review | Normal | Remediation evidence and authorized review request |
| Authorize | Suspend | Exceptional | Pre-activation issue or authorization hold |
| Activate | Suspend | Exceptional | Readiness or activation failure |
| Provision | Revoke | Exceptional | Issuance cancelled, compromised or invalid |
| Bind | Revoke | Exceptional | Binding invalid, compromised or withdrawn |
| Authorize | Revoke | Exceptional | Authority withdrawn or identity compromised |
| Activate | Revoke | Exceptional | Critical failure or authority withdrawal |
| Operate | Revoke | Exceptional | Critical incident, terminal withdrawal or binding loss |
| Review | Revoke | Normal or exceptional | Terminal review decision |
| Suspend | Revoke | Normal | Remediation rejected, authority ended or retirement decision |
| Revoke | Archive | Normal | Revocation effects completed and records preserved |

- **AGENT-REQ-126:** Only transitions listed in the governed transition matrix shall be accepted.
- **AGENT-REQ-127:** Direct transition from `Provision` to `Authorize` shall be prohibited.
- **AGENT-REQ-128:** Direct transition from `Bind` to `Activate` shall be prohibited.
- **AGENT-REQ-129:** Direct transition from `Authorize` to `Operate` shall be prohibited.
- **AGENT-REQ-130:** Direct transition from `Suspend` to `Activate` shall be prohibited.
- **AGENT-REQ-131:** Direct transition from `Suspend` to `Operate` shall be prohibited.
- **AGENT-REQ-132:** Direct transition from `Revoke` to any active state shall be prohibited.
- **AGENT-REQ-133:** Direct transition from `Archive` to any other state shall be prohibited.
- **AGENT-REQ-134:** Emergency transition to `Suspend` may occur from any non-terminal active or preparatory state when containment is required.
- **AGENT-REQ-135:** Emergency transition to `Revoke` may occur from any non-archived state under authorized terminal conditions.
- **AGENT-REQ-136:** A transition not represented in the matrix shall require a model change rather than an informal override.
- **AGENT-REQ-137:** Transition guards shall be evaluated against current, not historical, context.
- **AGENT-REQ-138:** Conflicting transition requests shall resolve to the more restrictive state.
- **AGENT-REQ-139:** Concurrent transition processing shall be serialized or use equivalent consistency controls.
- **AGENT-REQ-140:** Transition completion shall be confirmed before dependent activity begins.

## 7. Provision State

- **AGENT-REQ-141:** `Provision` shall represent institutional creation of a new agent identity and lifecycle record.
- **AGENT-REQ-142:** Provisioning shall require an approved issuance request.
- **AGENT-REQ-143:** Provisioning shall identify the accountable institution.
- **AGENT-REQ-144:** Provisioning shall identify the intended officer before binding.
- **AGENT-REQ-145:** Provisioning shall assign a unique immutable agent identity.
- **AGENT-REQ-146:** Provisioning shall assign a unique lifecycle identifier.
- **AGENT-REQ-147:** Provisioning shall identify identity issuer.
- **AGENT-REQ-148:** Provisioning shall identify identity assurance requirements.
- **AGENT-REQ-149:** Provisioning shall identify initial key and credential requirements.
- **AGENT-REQ-150:** Provisioning shall identify intended runtime class.
- **AGENT-REQ-151:** Provisioning shall identify intended tool and connector classes.
- **AGENT-REQ-152:** Provisioning shall identify initial case or mandate constraints where known.
- **AGENT-REQ-153:** Provisioning shall identify initial purpose constraints.
- **AGENT-REQ-154:** Provisioning shall identify jurisdictional constraints where known.
- **AGENT-REQ-155:** Provisioning shall identify data-classification constraints.
- **AGENT-REQ-156:** Provisioning shall identify expiry and review requirements.
- **AGENT-REQ-157:** Provisioning shall identify owners and custodians.
- **AGENT-REQ-158:** Provisioning shall identify required threat and security controls.
- **AGENT-REQ-159:** Provisioning shall use non-production synthetic identities in public demonstrations.
- **AGENT-REQ-160:** Provisioning shall not embed human credentials.
- **AGENT-REQ-161:** Provisioning shall not activate the agent.
- **AGENT-REQ-162:** Provisioning shall not grant operational authorization.
- **AGENT-REQ-163:** Provisioning shall not establish legal authority.
- **AGENT-REQ-164:** Provisioning shall not create shared or orphaned ownership.
- **AGENT-REQ-165:** Provisioning shall use approved identifier and naming conventions.
- **AGENT-REQ-166:** Provisioning shall create an audit record.
- **AGENT-REQ-167:** Provisioning failure shall leave no partially active identity.
- **AGENT-REQ-168:** Duplicate identity allocation shall fail closed.
- **AGENT-REQ-169:** Provisioning cancellation shall transition to `Revoke` or an attributable rejected record.
- **AGENT-REQ-170:** Provisioning tests shall verify duplicate, malformed, unauthorized and partial requests.

## 8. Bind State

- **AGENT-REQ-171:** `Bind` shall represent cryptographic and governance binding of the agent to one identified officer and one institution.
- **AGENT-REQ-172:** Binding shall verify officer identity.
- **AGENT-REQ-173:** Binding shall verify institutional affiliation or authority.
- **AGENT-REQ-174:** Binding shall identify binding authority.
- **AGENT-REQ-175:** Binding shall identify binding method.
- **AGENT-REQ-176:** Binding shall identify cryptographic-binding reference.
- **AGENT-REQ-177:** Binding shall identify agent identity.
- **AGENT-REQ-178:** Binding shall identify officer identity.
- **AGENT-REQ-179:** Binding shall identify institution.
- **AGENT-REQ-180:** Binding shall identify effective time.
- **AGENT-REQ-181:** Binding shall identify expiry or review conditions.
- **AGENT-REQ-182:** Binding shall identify case or mandate relationship where applicable.
- **AGENT-REQ-183:** Binding shall identify restrictions.
- **AGENT-REQ-184:** Binding shall identify revocation reference.
- **AGENT-REQ-185:** Binding shall be integrity-protected.
- **AGENT-REQ-186:** Binding shall be attributable.
- **AGENT-REQ-187:** Binding shall not transfer officer credentials.
- **AGENT-REQ-188:** Binding shall not duplicate the officer’s identity.
- **AGENT-REQ-189:** Binding shall not create independent agent authority.
- **AGENT-REQ-190:** Binding shall not be shared across multiple officers.
- **AGENT-REQ-191:** Binding to multiple institutions shall be prohibited unless a superior approved design explicitly defines accountable primary authority and legal basis.
- **AGENT-REQ-192:** Officer identity uncertainty shall prevent binding.
- **AGENT-REQ-193:** Institutional authority uncertainty shall prevent binding.
- **AGENT-REQ-194:** Binding-key compromise shall trigger suspension or revocation.
- **AGENT-REQ-195:** Binding validation failure shall prevent authorization.
- **AGENT-REQ-196:** Binding replacement for the same officer shall preserve history and invalidate prior binding material.
- **AGENT-REQ-197:** Binding to a different officer shall require a new agent identity.
- **AGENT-REQ-198:** Binding records shall remain after revocation.
- **AGENT-REQ-199:** Binding status shall be continuously or periodically verifiable.
- **AGENT-REQ-200:** Binding tests shall include wrong officer, wrong institution, expired binding, substituted key and duplicate binding.

## 9. Authorize State

- **AGENT-REQ-201:** `Authorize` shall represent a valid, explicit and current authorization decision for bounded operation.
- **AGENT-REQ-202:** Authorization shall occur only after successful binding.
- **AGENT-REQ-203:** Authorization shall conform to `OBDIA-AUTH-001`.
- **AGENT-REQ-204:** Authorization shall identify officer.
- **AGENT-REQ-205:** Authorization shall identify institution.
- **AGENT-REQ-206:** Authorization shall identify agent identity.
- **AGENT-REQ-207:** Authorization shall identify case or mandate.
- **AGENT-REQ-208:** Authorization shall identify purpose.
- **AGENT-REQ-209:** Authorization shall identify jurisdiction where applicable.
- **AGENT-REQ-210:** Authorization shall identify validity period.
- **AGENT-REQ-211:** Authorization shall identify approved tools and connectors.
- **AGENT-REQ-212:** Authorization shall identify approved actions.
- **AGENT-REQ-213:** Authorization shall identify approved data scope.
- **AGENT-REQ-214:** Authorization shall identify human-approval requirements.
- **AGENT-REQ-215:** Authorization shall identify obligations.
- **AGENT-REQ-216:** Authorization shall identify policy bundle and version.
- **AGENT-REQ-217:** Authorization shall identify delegation reference.
- **AGENT-REQ-218:** Authorization shall identify revocation conditions.
- **AGENT-REQ-219:** Authorization shall default to deny.
- **AGENT-REQ-220:** Authorization shall not include canonical prohibitions.
- **AGENT-REQ-221:** Authorization shall not be inferred from binding alone.
- **AGENT-REQ-222:** Authorization shall not be inferred from authentication alone.
- **AGENT-REQ-223:** Authorization shall not be inferred from repository access.
- **AGENT-REQ-224:** Authorization shall not be inferred from model output.
- **AGENT-REQ-225:** Authorization shall be revocable.
- **AGENT-REQ-226:** Authorization shall be auditable.
- **AGENT-REQ-227:** Authorization expiry shall prevent activation or continued operation.
- **AGENT-REQ-228:** Authorization changes shall invalidate affected sessions and cached decisions.
- **AGENT-REQ-229:** Authorization failure shall result in `Suspend`, `Revoke` or remaining in a non-operational state.
- **AGENT-REQ-230:** Authorization tests shall verify scope, expiry, revocation, policy conflict and default deny.

## 10. Activate State

- **AGENT-REQ-231:** `Activate` shall represent controlled preparation of an authorized agent for operation.
- **AGENT-REQ-232:** Activation shall require current successful authorization.
- **AGENT-REQ-233:** Activation shall verify officer-agent binding.
- **AGENT-REQ-234:** Activation shall verify institutional binding.
- **AGENT-REQ-235:** Activation shall verify credential validity.
- **AGENT-REQ-236:** Activation shall verify policy availability and integrity.
- **AGENT-REQ-237:** Activation shall verify runtime identity.
- **AGENT-REQ-238:** Activation shall verify runtime configuration.
- **AGENT-REQ-239:** Activation shall verify model and prompt bundle where material.
- **AGENT-REQ-240:** Activation shall verify approved connector availability.
- **AGENT-REQ-241:** Activation shall verify connector identity and revocation state.
- **AGENT-REQ-242:** Activation shall verify monitoring readiness.
- **AGENT-REQ-243:** Activation shall verify audit readiness.
- **AGENT-REQ-244:** Activation shall verify evidence-handling readiness where applicable.
- **AGENT-REQ-245:** Activation shall verify network and environment isolation.
- **AGENT-REQ-246:** Activation shall verify secret and key access boundaries.
- **AGENT-REQ-247:** Activation shall verify session creation controls.
- **AGENT-REQ-248:** Activation shall verify time and expiry context.
- **AGENT-REQ-249:** Activation shall identify exact code, model, policy, configuration and dependency versions.
- **AGENT-REQ-250:** Activation shall create an activation event.
- **AGENT-REQ-251:** Activation shall not begin protected work before readiness confirmation.
- **AGENT-REQ-252:** Activation shall not broaden authorized scope.
- **AGENT-REQ-253:** Activation shall not silently substitute a provider, model, connector or policy.
- **AGENT-REQ-254:** Activation failure shall result in suspension or return to authorization review.
- **AGENT-REQ-255:** Partial activation shall not produce `Operate` state.
- **AGENT-REQ-256:** Activation timeout shall not produce `Operate` state.
- **AGENT-REQ-257:** Activation shall be idempotent or detect duplicate requests safely.
- **AGENT-REQ-258:** Activation shall establish a bounded operational session.
- **AGENT-REQ-259:** Activation evidence shall remain traceable to later operation.
- **AGENT-REQ-260:** Activation tests shall verify unavailable policy, expired credential, failed monitoring, failed audit and connector mismatch.

## 11. Operate State

- **AGENT-REQ-261:** `Operate` shall represent authorized execution within current lifecycle, delegation and policy constraints.
- **AGENT-REQ-262:** Operation shall remain bound to the identified officer.
- **AGENT-REQ-263:** Operation shall remain attributable to the accountable institution.
- **AGENT-REQ-264:** Operation shall remain within approved case or mandate.
- **AGENT-REQ-265:** Operation shall remain within approved purpose.
- **AGENT-REQ-266:** Operation shall remain within applicable jurisdiction.
- **AGENT-REQ-267:** Operation shall remain within approved time.
- **AGENT-REQ-268:** Operation shall use only approved tools and connectors.
- **AGENT-REQ-269:** Operation shall remain within approved data scope.
- **AGENT-REQ-270:** Operation shall continuously enforce authorization obligations.
- **AGENT-REQ-271:** Operation shall preserve auditability.
- **AGENT-REQ-272:** Operation shall preserve evidence provenance and custody where applicable.
- **AGENT-REQ-273:** Operation shall preserve separation between source evidence and model output.
- **AGENT-REQ-274:** Operation shall not create independent legal authority.
- **AGENT-REQ-275:** Operation shall not perform prohibited conduct.
- **AGENT-REQ-276:** Operation shall not transfer authority to another agent autonomously.
- **AGENT-REQ-277:** Operation shall not use human credentials.
- **AGENT-REQ-278:** Operation shall not interact with uncontrolled criminal infrastructure.
- **AGENT-REQ-279:** Operation shall re-evaluate authorization at defined checkpoints.
- **AGENT-REQ-280:** Operation shall react to revocation.
- **AGENT-REQ-281:** Operation shall react to binding invalidation.
- **AGENT-REQ-282:** Operation shall react to policy changes.
- **AGENT-REQ-283:** Operation shall react to critical trust degradation.
- **AGENT-REQ-284:** Operation shall react to credential expiry.
- **AGENT-REQ-285:** Operation shall react to monitoring or audit failure.
- **AGENT-REQ-286:** Operation shall support immediate suspension.
- **AGENT-REQ-287:** Operation shall support terminal revocation.
- **AGENT-REQ-288:** Long-running operation shall have bounded duration and renewal rules.
- **AGENT-REQ-289:** Operation completion shall create a result event.
- **AGENT-REQ-290:** Operation tests shall include normal, denied, expired, revoked, degraded and interrupted paths.

## 12. Review State

- **AGENT-REQ-291:** `Review` shall represent formal reassessment of agent identity, binding, authority, operation, controls and risk.
- **AGENT-REQ-292:** Review may be scheduled.
- **AGENT-REQ-293:** Review may be triggered by expiry.
- **AGENT-REQ-294:** Review may be triggered by material change.
- **AGENT-REQ-295:** Review may be triggered by an incident.
- **AGENT-REQ-296:** Review may be triggered by anomalous behavior.
- **AGENT-REQ-297:** Review may be triggered by failed validation.
- **AGENT-REQ-298:** Review may be triggered by connector or provider change.
- **AGENT-REQ-299:** Review may be triggered by policy change.
- **AGENT-REQ-300:** Review may be triggered by officer or institutional change.
- **AGENT-REQ-301:** Review may be triggered by case or purpose change.
- **AGENT-REQ-302:** Review may be triggered by jurisdictional change.
- **AGENT-REQ-303:** Review may be triggered by risk or exception expiry.
- **AGENT-REQ-304:** Review shall identify exact agent identity and lifecycle state.
- **AGENT-REQ-305:** Review shall verify officer-agent binding.
- **AGENT-REQ-306:** Review shall verify institutional authority.
- **AGENT-REQ-307:** Review shall verify delegation and authorization.
- **AGENT-REQ-308:** Review shall verify credentials and revocation status.
- **AGENT-REQ-309:** Review shall verify runtime, model, policy and connector versions.
- **AGENT-REQ-310:** Review shall verify operation history and audit completeness.
- **AGENT-REQ-311:** Review shall verify evidence handling where applicable.
- **AGENT-REQ-312:** Review shall verify unresolved findings, incidents, risks and exceptions.
- **AGENT-REQ-313:** Review shall verify privacy and fundamental-rights constraints.
- **AGENT-REQ-314:** Review shall identify decision: return to `Operate`, return to `Authorize`, transition to `Suspend` or transition to `Revoke`.
- **AGENT-REQ-315:** Review shall identify reviewer identity and role.
- **AGENT-REQ-316:** Review shall identify exact evidence and time.
- **AGENT-REQ-317:** Review shall preserve role-concentration disclosure.
- **AGENT-REQ-318:** An AI system shall not be final lifecycle reviewer.
- **AGENT-REQ-319:** Material review findings shall remain visible.
- **AGENT-REQ-320:** Review tests shall verify each possible disposition.

## 13. Suspend State

- **AGENT-REQ-321:** `Suspend` shall represent reversible containment that prevents ordinary operation.
- **AGENT-REQ-322:** Suspension shall stop new protected operations.
- **AGENT-REQ-323:** Suspension shall prevent new sessions.
- **AGENT-REQ-324:** Suspension shall invalidate or restrict active sessions.
- **AGENT-REQ-325:** Suspension shall prevent new connector actions.
- **AGENT-REQ-326:** Suspension shall preserve evidence and audit records.
- **AGENT-REQ-327:** Suspension shall preserve lifecycle and binding history.
- **AGENT-REQ-328:** Suspension shall preserve revocation capability.
- **AGENT-REQ-329:** Suspension shall identify reason.
- **AGENT-REQ-330:** Suspension shall identify authority.
- **AGENT-REQ-331:** Suspension shall identify effective time.
- **AGENT-REQ-332:** Suspension shall identify scope.
- **AGENT-REQ-333:** Suspension shall identify conditions for review.
- **AGENT-REQ-334:** Suspension shall identify remediation owner.
- **AGENT-REQ-335:** Suspension shall identify expiry or next-review time where applicable.
- **AGENT-REQ-336:** Emergency suspension shall be available for critical containment.
- **AGENT-REQ-337:** Emergency suspension shall be attributable.
- **AGENT-REQ-338:** Emergency suspension shall not authorize otherwise prohibited actions.
- **AGENT-REQ-339:** Suspension shall not erase credentials silently; credentials shall be disabled, restricted or revoked according to policy.
- **AGENT-REQ-340:** Suspension shall not imply final retirement.
- **AGENT-REQ-341:** Suspension shall not permit operation through a fallback identity.
- **AGENT-REQ-342:** Suspension shall not permit operation through another connector or environment.
- **AGENT-REQ-343:** Suspension shall propagate to authorization, session, workload and connector controls.
- **AGENT-REQ-344:** Suspension propagation latency shall be defined and tested.
- **AGENT-REQ-345:** Suspension failure shall trigger incident response.
- **AGENT-REQ-346:** Suspended agents shall not return directly to `Activate` or `Operate`.
- **AGENT-REQ-347:** Suspended agents shall enter `Review` before fresh authorization.
- **AGENT-REQ-348:** Suspension records shall remain retrievable.
- **AGENT-REQ-349:** Suspension shall be visible to authorized operators.
- **AGENT-REQ-350:** Suspension tests shall verify propagation, blocked operations, evidence preservation and review-only recovery.

## 14. Revoke State

- **AGENT-REQ-351:** `Revoke` shall represent terminal withdrawal of authority and usability for the agent identity.
- **AGENT-REQ-352:** Revocation shall be attributable to authorized human or institutional authority.
- **AGENT-REQ-353:** Revocation shall identify reason.
- **AGENT-REQ-354:** Revocation shall identify effective time.
- **AGENT-REQ-355:** Revocation shall identify affected identity, credentials, bindings, delegations, sessions, workloads and connectors.
- **AGENT-REQ-356:** Revocation shall prevent new authorization.
- **AGENT-REQ-357:** Revocation shall prevent activation.
- **AGENT-REQ-358:** Revocation shall prevent operation.
- **AGENT-REQ-359:** Revocation shall invalidate active sessions.
- **AGENT-REQ-360:** Revocation shall invalidate cached authorization decisions.
- **AGENT-REQ-361:** Revocation shall disable or revoke agent credentials.
- **AGENT-REQ-362:** Revocation shall disable or revoke binding material where applicable.
- **AGENT-REQ-363:** Revocation shall disable connector-scoped credentials where applicable.
- **AGENT-REQ-364:** Revocation shall stop in-progress operations when safe and required.
- **AGENT-REQ-365:** Revocation shall preserve evidence and audit history.
- **AGENT-REQ-366:** Revocation shall preserve chain-of-custody records.
- **AGENT-REQ-367:** Revocation shall preserve incident and review records.
- **AGENT-REQ-368:** Revocation shall not delete lawfully retained evidence.
- **AGENT-REQ-369:** Revocation shall not be reversible for the same agent identity.
- **AGENT-REQ-370:** Reinstatement shall require provisioning a new agent identity.
- **AGENT-REQ-371:** Revocation shall not transfer prior authority to a new identity.
- **AGENT-REQ-372:** Revocation propagation latency shall be defined and tested.
- **AGENT-REQ-373:** Revocation-service failure shall trigger containment.
- **AGENT-REQ-374:** Partial revocation shall remain a critical unresolved condition.
- **AGENT-REQ-375:** Revocation conflicts shall fail closed.
- **AGENT-REQ-376:** Revocation shall generate a terminal lifecycle event.
- **AGENT-REQ-377:** Revocation shall identify archive prerequisites.
- **AGENT-REQ-378:** Revocation shall identify retention and destruction obligations.
- **AGENT-REQ-379:** Revoked identity references shall remain reserved and shall not be reused.
- **AGENT-REQ-380:** Revocation tests shall verify every dependent identity, session, policy and connector path.

## 15. Archive State

- **AGENT-REQ-381:** `Archive` shall be the terminal retained state after revocation effects are complete.
- **AGENT-REQ-382:** Archival shall require completed revocation.
- **AGENT-REQ-383:** Archival shall identify archive authority.
- **AGENT-REQ-384:** Archival shall identify archive time.
- **AGENT-REQ-385:** Archival shall identify retained lifecycle records.
- **AGENT-REQ-386:** Archival shall identify retained binding and authorization records.
- **AGENT-REQ-387:** Archival shall identify retained audit and evidence records.
- **AGENT-REQ-388:** Archival shall identify retention schedule.
- **AGENT-REQ-389:** Archival shall identify access controls.
- **AGENT-REQ-390:** Archival shall identify lawful deletion or disposition rules.
- **AGENT-REQ-391:** Archived agent identities shall remain unique and reserved.
- **AGENT-REQ-392:** Archived agent identities shall not be reactivated.
- **AGENT-REQ-393:** Archived agent identities shall not be rebound.
- **AGENT-REQ-394:** Archived credentials shall not remain usable.
- **AGENT-REQ-395:** Archived sessions shall not remain usable.
- **AGENT-REQ-396:** Archived connectors shall not retain agent-scoped access.
- **AGENT-REQ-397:** Archive integrity shall be verifiable.
- **AGENT-REQ-398:** Archive access shall use least privilege.
- **AGENT-REQ-399:** Archive access shall be logged.
- **AGENT-REQ-400:** Archive restoration shall not restore authority.
- **AGENT-REQ-401:** Archive restoration for review shall create an attributable access event.
- **AGENT-REQ-402:** Archived lifecycle history shall remain chronological.
- **AGENT-REQ-403:** Archived records shall not be silently edited.
- **AGENT-REQ-404:** Archive corrections shall be additive.
- **AGENT-REQ-405:** Archive migration shall preserve integrity and traceability.
- **AGENT-REQ-406:** Archive failure shall trigger incident handling.
- **AGENT-REQ-407:** Archive visibility shall respect privacy and classification.
- **AGENT-REQ-408:** Public release of archive material shall require separate authorization and release review.
- **AGENT-REQ-409:** Archive retention shall not be represented as legal compliance without evidence.
- **AGENT-REQ-410:** Archive tests shall verify terminal state, credential unusability, integrity and controlled retrieval.

## 16. Review, Reauthorization and Reactivation

- **AGENT-REQ-411:** Return from `Review` to `Operate` shall require confirmation that current authorization remains valid.
- **AGENT-REQ-412:** Return from `Review` to `Operate` shall require no blocking finding.
- **AGENT-REQ-413:** Return from `Review` to `Operate` shall identify reviewer and evidence.
- **AGENT-REQ-414:** Material scope change shall require transition from `Review` to `Authorize`.
- **AGENT-REQ-415:** Material policy change shall require fresh authorization.
- **AGENT-REQ-416:** Material case or purpose change shall require fresh authorization.
- **AGENT-REQ-417:** Material jurisdiction change shall require fresh authorization.
- **AGENT-REQ-418:** Material tool or connector change shall require fresh authorization.
- **AGENT-REQ-419:** Material data-scope change shall require fresh authorization.
- **AGENT-REQ-420:** Material officer or institution change shall not use reauthorization; it shall require new identity issuance.
- **AGENT-REQ-421:** Suspended agents shall transition to `Review` before reauthorization.
- **AGENT-REQ-422:** Fresh authorization shall verify binding and current identity status.
- **AGENT-REQ-423:** Fresh authorization shall invalidate stale sessions and cached decisions.
- **AGENT-REQ-424:** Reactivation shall use `Authorize → Activate → Operate`.
- **AGENT-REQ-425:** Reactivation shall not bypass activation readiness checks.
- **AGENT-REQ-426:** Reactivation shall identify remediation evidence.
- **AGENT-REQ-427:** Reactivation shall identify residual risks and exceptions.
- **AGENT-REQ-428:** Reactivation shall identify changed versions and dependencies.
- **AGENT-REQ-429:** Reactivation shall be prohibited when revocation has occurred.
- **AGENT-REQ-430:** Reactivation tests shall verify fresh policy, session replacement, connector readiness and denial of stale credentials.

## 17. Officer Change and Rebinding Prohibition

- **AGENT-REQ-431:** An existing agent identity shall not be rebound to a different officer.
- **AGENT-REQ-432:** An existing lifecycle record shall not replace the bound officer identifier.
- **AGENT-REQ-433:** Officer reassignment shall trigger suspension and review of the existing agent.
- **AGENT-REQ-434:** Terminal officer reassignment shall trigger revocation.
- **AGENT-REQ-435:** A successor officer shall receive a separately provisioned agent identity.
- **AGENT-REQ-436:** A successor agent shall receive a separately allocated lifecycle identifier.
- **AGENT-REQ-437:** Prior authorization shall not transfer to the successor agent.
- **AGENT-REQ-438:** Prior sessions shall not transfer to the successor agent.
- **AGENT-REQ-439:** Prior credentials shall not transfer to the successor agent.
- **AGENT-REQ-440:** Prior connector tokens shall not transfer to the successor agent.
- **AGENT-REQ-441:** Prior memory shall not transfer automatically.
- **AGENT-REQ-442:** Prior evidence access shall require separate authorization.
- **AGENT-REQ-443:** Prior case knowledge transfer shall be explicit, minimized, authorized and audited.
- **AGENT-REQ-444:** Historical records shall preserve the prior officer-agent relationship.
- **AGENT-REQ-445:** Officer departure shall trigger defined containment and revocation procedures.
- **AGENT-REQ-446:** Officer suspension shall trigger agent suspension or revocation according to policy.
- **AGENT-REQ-447:** Officer credential compromise shall trigger agent containment review.
- **AGENT-REQ-448:** Institutional change shall require revocation and new provisioning unless superior approved authority defines otherwise.
- **AGENT-REQ-449:** Rebinding attempts shall be detected and logged.
- **AGENT-REQ-450:** Rebinding tests shall verify denial across identity, session, connector and data paths.

## 18. Credential and Key Lifecycle

- **AGENT-REQ-451:** Agent credentials shall have distinct identifiers.
- **AGENT-REQ-452:** Agent credentials shall have identified issuer and audience.
- **AGENT-REQ-453:** Agent credentials shall have explicit purpose and scope.
- **AGENT-REQ-454:** Agent credentials shall have validity periods.
- **AGENT-REQ-455:** Agent credentials shall be revocable.
- **AGENT-REQ-456:** Agent credentials shall not contain human passwords or reusable authentication secrets.
- **AGENT-REQ-457:** Agent keys shall have identified owners and custodians.
- **AGENT-REQ-458:** Agent keys shall be generated through approved mechanisms.
- **AGENT-REQ-459:** Agent keys shall be protected proportionate to risk.
- **AGENT-REQ-460:** Agent keys shall support rotation.
- **AGENT-REQ-461:** Key rotation shall preserve lifecycle history.
- **AGENT-REQ-462:** Key rotation shall not change officer binding.
- **AGENT-REQ-463:** Key rotation shall not broaden authorization.
- **AGENT-REQ-464:** Credential rotation shall invalidate superseded credentials.
- **AGENT-REQ-465:** Credential expiry shall prevent activation and operation.
- **AGENT-REQ-466:** Credential compromise shall trigger suspension or revocation.
- **AGENT-REQ-467:** Binding-key compromise shall trigger binding review and likely revocation.
- **AGENT-REQ-468:** Credentials shall be distinct across development, testing, demonstration and operational environments.
- **AGENT-REQ-469:** Credential backups shall be governed.
- **AGENT-REQ-470:** Credential recovery shall not bypass lifecycle authorization.
- **AGENT-REQ-471:** Credential issuance shall create audit events.
- **AGENT-REQ-472:** Credential use shall be attributable.
- **AGENT-REQ-473:** Credential revocation shall propagate to sessions and connectors.
- **AGENT-REQ-474:** Credential state shall be verified during activation and operation.
- **AGENT-REQ-475:** Credential material shall not appear in logs, documents or evidence metadata.
- **AGENT-REQ-476:** Credential inventory shall remain accurate.
- **AGENT-REQ-477:** Orphaned credentials shall be detected and revoked.
- **AGENT-REQ-478:** Expired credentials shall not be silently extended.
- **AGENT-REQ-479:** Credential failure shall not trigger permissive fallback.
- **AGENT-REQ-480:** Credential-lifecycle tests shall verify issue, rotate, expire, revoke, compromise and recovery paths.

## 19. Session and Workload Lifecycle Integration

- **AGENT-REQ-481:** Agent lifecycle state shall constrain session creation.
- **AGENT-REQ-482:** Only authorized and activating or operating agents may create operational sessions.
- **AGENT-REQ-483:** Sessions shall identify agent, officer, institution, case and purpose where applicable.
- **AGENT-REQ-484:** Sessions shall have bounded lifetime.
- **AGENT-REQ-485:** Sessions shall be invalidated on suspension.
- **AGENT-REQ-486:** Sessions shall be invalidated on revocation.
- **AGENT-REQ-487:** Sessions shall be re-evaluated on material authorization change.
- **AGENT-REQ-488:** Sessions shall not survive officer rebinding because rebinding is prohibited.
- **AGENT-REQ-489:** Sessions shall not transfer to successor agents.
- **AGENT-REQ-490:** Workload identities shall remain distinct from agent identities.
- **AGENT-REQ-491:** Workloads shall be created only for authorized agent operations.
- **AGENT-REQ-492:** Workloads shall inherit no more authority than the agent operation.
- **AGENT-REQ-493:** Workload authority shall expire no later than the agent session.
- **AGENT-REQ-494:** Workloads shall be suspended or terminated when the agent is suspended.
- **AGENT-REQ-495:** Workloads shall be terminated when the agent is revoked.
- **AGENT-REQ-496:** Workload credentials shall be independently revocable where feasible.
- **AGENT-REQ-497:** Orphaned workloads shall be detected and contained.
- **AGENT-REQ-498:** Long-running workloads shall re-evaluate lifecycle and authorization state.
- **AGENT-REQ-499:** Workload completion shall create attributable results.
- **AGENT-REQ-500:** Session and workload tests shall verify expiry, suspension, revocation, orphan detection and authority bounds.

## 20. Runtime, Model and Configuration Changes

- **AGENT-REQ-501:** Runtime identity shall be recorded during activation and operation.
- **AGENT-REQ-502:** Runtime version shall be recorded.
- **AGENT-REQ-503:** Model provider and model version shall be recorded where material.
- **AGENT-REQ-504:** Prompt and policy bundle versions shall be recorded where material.
- **AGENT-REQ-505:** Configuration version shall be recorded.
- **AGENT-REQ-506:** Dependency versions shall be recorded where material.
- **AGENT-REQ-507:** Runtime change shall trigger lifecycle impact assessment.
- **AGENT-REQ-508:** Model change shall trigger trust and risk reassessment.
- **AGENT-REQ-509:** Provider change shall trigger privacy, security and availability reassessment.
- **AGENT-REQ-510:** Prompt-policy change shall trigger authorization and safety reassessment.
- **AGENT-REQ-511:** Configuration change shall not broaden authority silently.
- **AGENT-REQ-512:** Dependency change shall trigger supply-chain review.
- **AGENT-REQ-513:** Material runtime change shall require review.
- **AGENT-REQ-514:** Material runtime change may require fresh authorization.
- **AGENT-REQ-515:** Unapproved runtime substitution shall trigger suspension.
- **AGENT-REQ-516:** Fallback models or providers shall be pre-approved and no more permissive.
- **AGENT-REQ-517:** Unknown runtime version shall prevent activation.
- **AGENT-REQ-518:** Runtime integrity failure shall trigger suspension or revocation.
- **AGENT-REQ-519:** Runtime state shall remain traceable to operation and evidence outputs.
- **AGENT-REQ-520:** Runtime-change tests shall verify approval, rollback, fallback and denial of unknown versions.

## 21. Connector and Tool Lifecycle Integration

- **AGENT-REQ-521:** Connector access shall depend on current agent lifecycle and authorization state.
- **AGENT-REQ-522:** Connectors shall identify the agent, workload and session using them.
- **AGENT-REQ-523:** Connector credentials shall not outlive applicable agent authorization.
- **AGENT-REQ-524:** Connector sessions shall be invalidated on agent suspension.
- **AGENT-REQ-525:** Connector sessions shall be revoked on agent revocation.
- **AGENT-REQ-526:** Connector enablement shall be a protected transition-related action.
- **AGENT-REQ-527:** Connector disablement shall support emergency containment.
- **AGENT-REQ-528:** Connector change shall trigger review where capability or risk changes.
- **AGENT-REQ-529:** Connector compromise shall trigger agent suspension or scoped containment.
- **AGENT-REQ-530:** Connector failure shall not trigger broader fallback permissions.
- **AGENT-REQ-531:** Unapproved connectors shall not be available during activation or operation.
- **AGENT-REQ-532:** Read-only connector mode shall be preferred where sufficient.
- **AGENT-REQ-533:** Connector outputs shall remain untrusted until validated.
- **AGENT-REQ-534:** Connector use shall preserve provenance.
- **AGENT-REQ-535:** Connector use shall preserve authorization boundaries.
- **AGENT-REQ-536:** Connector logs shall correlate with lifecycle and operation records.
- **AGENT-REQ-537:** Operationally isolated Dark Web research shall not interact with uncontrolled criminal infrastructure.
- **AGENT-REQ-538:** Connector lifecycle shall conform to document 20 after consolidation.
- **AGENT-REQ-539:** Connector retirement shall remove agent access.
- **AGENT-REQ-540:** Connector-lifecycle tests shall verify activation, suspension, revocation, substitution and failure isolation.

## 22. Case, Purpose, Jurisdiction and Time Boundaries

- **AGENT-REQ-541:** Agent lifecycle authorization shall identify case or mandate where applicable.
- **AGENT-REQ-542:** Agent lifecycle authorization shall identify purpose.
- **AGENT-REQ-543:** Agent lifecycle authorization shall identify jurisdictional constraints where applicable.
- **AGENT-REQ-544:** Agent lifecycle authorization shall identify validity period.
- **AGENT-REQ-545:** Case closure shall trigger review, suspension or revocation.
- **AGENT-REQ-546:** Purpose completion shall trigger review, suspension or revocation.
- **AGENT-REQ-547:** Jurisdictional change shall trigger review and fresh authorization where applicable.
- **AGENT-REQ-548:** Authorization expiry shall trigger suspension or revocation.
- **AGENT-REQ-549:** Case-scoped sessions shall not cross into unrelated cases.
- **AGENT-REQ-550:** Purpose-scoped memory shall not cross into unrelated purposes.
- **AGENT-REQ-551:** Case-scoped connector access shall not cross into unrelated cases.
- **AGENT-REQ-552:** Cross-case access shall require explicit separate authorization.
- **AGENT-REQ-553:** Cross-purpose reuse shall require explicit separate authorization.
- **AGENT-REQ-554:** Cross-jurisdiction use shall require explicit review where applicable.
- **AGENT-REQ-555:** Time extension shall require fresh authorization.
- **AGENT-REQ-556:** Clock uncertainty beyond approved tolerance shall prevent time-sensitive operation.
- **AGENT-REQ-557:** Lifecycle records shall identify controlling time source.
- **AGENT-REQ-558:** Expired exceptions shall trigger suspension, denial or renewed review.
- **AGENT-REQ-559:** Expired risk acceptance shall trigger review.
- **AGENT-REQ-560:** Boundary changes shall remain auditable.

## 23. Monitoring and Health

- **AGENT-REQ-561:** Agent lifecycle state shall be observable to authorized operators.
- **AGENT-REQ-562:** Monitoring shall identify the current lifecycle state.
- **AGENT-REQ-563:** Monitoring shall identify state-transition failures.
- **AGENT-REQ-564:** Monitoring shall identify binding-verification failures.
- **AGENT-REQ-565:** Monitoring shall identify authorization expiry.
- **AGENT-REQ-566:** Monitoring shall identify credential expiry and revocation.
- **AGENT-REQ-567:** Monitoring shall identify session and workload mismatch.
- **AGENT-REQ-568:** Monitoring shall identify orphaned workloads.
- **AGENT-REQ-569:** Monitoring shall identify unauthorized connector use.
- **AGENT-REQ-570:** Monitoring shall identify policy and runtime version drift.
- **AGENT-REQ-571:** Monitoring shall identify audit gaps.
- **AGENT-REQ-572:** Monitoring shall identify anomalous operation volume or duration.
- **AGENT-REQ-573:** Monitoring shall identify cross-case or cross-purpose attempts.
- **AGENT-REQ-574:** Monitoring shall identify repeated denied transitions.
- **AGENT-REQ-575:** Monitoring shall identify suspension-propagation failure.
- **AGENT-REQ-576:** Monitoring shall identify revocation-propagation failure.
- **AGENT-REQ-577:** Monitoring shall identify archive-integrity failure.
- **AGENT-REQ-578:** Monitoring failure shall produce a degraded state.
- **AGENT-REQ-579:** Critical monitoring failure shall trigger suspension where required.
- **AGENT-REQ-580:** Monitoring shall minimize personal and case data.
- **AGENT-REQ-581:** Monitoring access shall be authorized.
- **AGENT-REQ-582:** Monitoring events shall correlate with lifecycle events.
- **AGENT-REQ-583:** Alert rules shall be versioned and tested.
- **AGENT-REQ-584:** Alert suppression shall be authorized and time-bounded.
- **AGENT-REQ-585:** Lifecycle health metrics shall not be represented as legal or operational approval.

## 24. Failure and Degraded-State Handling

- **AGENT-REQ-586:** Identity-service failure shall not permit activation or operation.
- **AGENT-REQ-587:** Binding-verification failure shall not permit authorization or operation.
- **AGENT-REQ-588:** Authorization-service failure shall not permit protected operation.
- **AGENT-REQ-589:** Revocation-service failure shall restrict or suspend privileged operation.
- **AGENT-REQ-590:** Audit-service failure shall block or restrict privileged operation.
- **AGENT-REQ-591:** Monitoring-service failure shall trigger degraded-state handling.
- **AGENT-REQ-592:** Policy-integrity failure shall trigger suspension.
- **AGENT-REQ-593:** Credential-integrity failure shall trigger suspension or revocation.
- **AGENT-REQ-594:** Runtime-integrity failure shall trigger suspension or revocation.
- **AGENT-REQ-595:** Connector-integrity failure shall trigger scoped containment.
- **AGENT-REQ-596:** Clock failure shall prevent time-sensitive transitions.
- **AGENT-REQ-597:** Storage failure shall preserve or quarantine lifecycle records.
- **AGENT-REQ-598:** Network failure shall not trigger uncontrolled direct connection.
- **AGENT-REQ-599:** Partial activation shall remain non-operational.
- **AGENT-REQ-600:** Partial suspension shall remain a critical incident.
- **AGENT-REQ-601:** Partial revocation shall remain a critical incident.
- **AGENT-REQ-602:** Unknown lifecycle state shall resolve to containment.
- **AGENT-REQ-603:** Conflicting lifecycle state shall resolve to the more restrictive state.
- **AGENT-REQ-604:** Stale lifecycle state shall not authorize operation.
- **AGENT-REQ-605:** Recovery shall restore a known governed state.
- **AGENT-REQ-606:** Recovery shall not restore revoked credentials or authority.
- **AGENT-REQ-607:** Post-recovery review shall be required.
- **AGENT-REQ-608:** Failure events shall preserve evidence.
- **AGENT-REQ-609:** Failure events shall identify owners and remediation.
- **AGENT-REQ-610:** Failure-path behavior shall be tested.

## 25. Emergency Containment

- **AGENT-REQ-611:** Emergency containment shall support immediate agent suspension.
- **AGENT-REQ-612:** Emergency containment shall support immediate credential revocation.
- **AGENT-REQ-613:** Emergency containment shall support connector disablement.
- **AGENT-REQ-614:** Emergency containment shall support session invalidation.
- **AGENT-REQ-615:** Emergency containment shall support workload termination or isolation.
- **AGENT-REQ-616:** Emergency containment shall support network isolation.
- **AGENT-REQ-617:** Emergency containment shall support evidence quarantine.
- **AGENT-REQ-618:** Emergency containment shall identify authorized initiator.
- **AGENT-REQ-619:** Emergency containment shall identify reason and scope.
- **AGENT-REQ-620:** Emergency containment shall create an immutable event.
- **AGENT-REQ-621:** Emergency containment shall not authorize offensive action.
- **AGENT-REQ-622:** Emergency containment shall not erase evidence.
- **AGENT-REQ-623:** Emergency containment shall not erase audit records.
- **AGENT-REQ-624:** Emergency containment shall minimize unrelated disruption.
- **AGENT-REQ-625:** Emergency containment shall preserve human accountability.
- **AGENT-REQ-626:** Emergency containment shall be reviewed retrospectively.
- **AGENT-REQ-627:** Emergency containment shall identify recovery or revocation path.
- **AGENT-REQ-628:** Emergency containment shall not become a hidden bypass.
- **AGENT-REQ-629:** Repeated emergency containment shall trigger architecture and risk review.
- **AGENT-REQ-630:** Emergency-containment tests shall verify speed, scope, propagation, evidence preservation and recovery.

## 26. Privacy and Fundamental-Rights Controls

- **AGENT-REQ-631:** Lifecycle records shall minimize personal data.
- **AGENT-REQ-632:** Officer identity references shall use controlled identifiers where feasible.
- **AGENT-REQ-633:** Case and subject identifiers shall be opaque where feasible.
- **AGENT-REQ-634:** Lifecycle monitoring shall not become excessive surveillance.
- **AGENT-REQ-635:** Lifecycle reviews shall consider purpose limitation.
- **AGENT-REQ-636:** Lifecycle reviews shall consider data minimization.
- **AGENT-REQ-637:** Lifecycle reviews shall consider jurisdiction.
- **AGENT-REQ-638:** Lifecycle reviews shall consider retention.
- **AGENT-REQ-639:** Lifecycle reviews shall consider fundamental-rights impact.
- **AGENT-REQ-640:** Agent operation shall not become autonomous policing.
- **AGENT-REQ-641:** Agent operation shall not become autonomous legal decision-making.
- **AGENT-REQ-642:** Agent operation shall not become autonomous coercion.
- **AGENT-REQ-643:** Agent operation shall not enable unlawful surveillance.
- **AGENT-REQ-644:** Agent operation shall not enable unlawful deanonymization.
- **AGENT-REQ-645:** Public demonstrations shall not use real officer identities or real cases.
- **AGENT-REQ-646:** Lifecycle analytics shall use minimized or pseudonymized data where feasible.
- **AGENT-REQ-647:** Suspension or revocation decisions shall remain attributable.
- **AGENT-REQ-648:** Automated recommendations shall not replace consequential human decisions.
- **AGENT-REQ-649:** Privacy incidents shall trigger lifecycle review.
- **AGENT-REQ-650:** Rights-impact findings shall have owners and dispositions.

## 27. Separation of Duties and Administrative Control

- **AGENT-REQ-651:** Provisioning authority and final activation approval shall be separated where feasible.
- **AGENT-REQ-652:** Identity issuance and lifecycle review shall be separated where feasible.
- **AGENT-REQ-653:** Policy authoring and policy approval shall be separated where feasible.
- **AGENT-REQ-654:** Connector development and high-risk connector enablement shall be separated where feasible.
- **AGENT-REQ-655:** Evidence custody and evidence disposition approval shall be separated where feasible.
- **AGENT-REQ-656:** An agent shall not approve its own transition.
- **AGENT-REQ-657:** A workload shall not approve its own lifecycle.
- **AGENT-REQ-658:** Lifecycle administrators shall not bypass authorization enforcement.
- **AGENT-REQ-659:** Lifecycle reviewers shall identify role concentration.
- **AGENT-REQ-660:** Role concentration shall require compensating controls.
- **AGENT-REQ-661:** Compensating controls may include immutable history, dual acknowledgement, time limits and independent logs.
- **AGENT-REQ-662:** Administrative lifecycle actions shall use strong authentication.
- **AGENT-REQ-663:** Administrative lifecycle actions shall use explicit authorization.
- **AGENT-REQ-664:** Administrative lifecycle actions shall be fully auditable.
- **AGENT-REQ-665:** Break-glass lifecycle action shall be exceptional and time-bounded.
- **AGENT-REQ-666:** Break-glass action shall not authorize canonical prohibitions.
- **AGENT-REQ-667:** Break-glass action shall trigger immediate review.
- **AGENT-REQ-668:** Expired administrative exceptions shall fail closed.
- **AGENT-REQ-669:** Repeated administrative overrides shall trigger governance review.
- **AGENT-REQ-670:** Internal multi-role review shall not be described as independent external assurance.

## 28. Expiry and Scheduled Events

- **AGENT-REQ-671:** Lifecycle authorizations shall have explicit validity or review conditions.
- **AGENT-REQ-672:** Credential expiry shall be scheduled and observable.
- **AGENT-REQ-673:** Binding review dates shall be scheduled where applicable.
- **AGENT-REQ-674:** Authorization review dates shall be scheduled.
- **AGENT-REQ-675:** Case and purpose expiry shall be scheduled where applicable.
- **AGENT-REQ-676:** Risk-acceptance expiry shall be scheduled.
- **AGENT-REQ-677:** Exception expiry shall be scheduled.
- **AGENT-REQ-678:** Connector approval expiry shall be scheduled where applicable.
- **AGENT-REQ-679:** Model and policy review dates shall be scheduled where applicable.
- **AGENT-REQ-680:** Scheduled events shall have immutable identifiers.
- **AGENT-REQ-681:** Scheduled events shall identify responsible owner.
- **AGENT-REQ-682:** Scheduled events shall identify expected transition or review.
- **AGENT-REQ-683:** Missed critical events shall trigger alerting.
- **AGENT-REQ-684:** Missed authorization expiry shall not extend authority.
- **AGENT-REQ-685:** Missed credential rotation shall trigger restriction or suspension.
- **AGENT-REQ-686:** Missed review shall trigger suspension where required.
- **AGENT-REQ-687:** Schedule changes shall be attributable.
- **AGENT-REQ-688:** Clock-source limitations shall be documented.
- **AGENT-REQ-689:** Scheduled-event automation shall fail closed.
- **AGENT-REQ-690:** Expiry tests shall verify time boundaries, delays and missed-event handling.

## 29. Concurrency, Idempotency and Ordering

- **AGENT-REQ-691:** Lifecycle transitions shall prevent conflicting concurrent state changes.
- **AGENT-REQ-692:** Transition processing shall use serialization, compare-and-set or equivalent consistency controls.
- **AGENT-REQ-693:** Every transition request shall have an idempotency identifier.
- **AGENT-REQ-694:** Duplicate transition requests shall not create duplicate credentials or authority.
- **AGENT-REQ-695:** Repeated successful transition requests shall return or reference the existing result safely.
- **AGENT-REQ-696:** Repeated failed transition requests shall remain attributable.
- **AGENT-REQ-697:** Out-of-order transition events shall be detected.
- **AGENT-REQ-698:** Late transition events shall not overwrite newer state.
- **AGENT-REQ-699:** State reads shall identify version or sequence.
- **AGENT-REQ-700:** Transition writes shall validate expected prior state.
- **AGENT-REQ-701:** Concurrent suspension and activation shall resolve to suspension.
- **AGENT-REQ-702:** Concurrent revocation and any active transition shall resolve to revocation.
- **AGENT-REQ-703:** Concurrent archive and recovery shall resolve to archive unless an authorized correction process applies.
- **AGENT-REQ-704:** Transition-event ordering shall identify its basis.
- **AGENT-REQ-705:** Clock time alone shall not be the only ordering control where stronger sequencing is available.
- **AGENT-REQ-706:** Retry policies shall not bypass transition guards.
- **AGENT-REQ-707:** Distributed duplicate processing shall remain safe.
- **AGENT-REQ-708:** Concurrency failures shall trigger reviewable findings.
- **AGENT-REQ-709:** Idempotency records shall have retention rules.
- **AGENT-REQ-710:** Concurrency and ordering shall be tested under race conditions.

## 30. Distributed Consistency and Propagation

- **AGENT-REQ-711:** Lifecycle state shall have one authoritative source of truth.
- **AGENT-REQ-712:** Replicated lifecycle state shall identify source version.
- **AGENT-REQ-713:** Replicas shall not independently broaden state.
- **AGENT-REQ-714:** State propagation latency shall be defined.
- **AGENT-REQ-715:** Suspension propagation latency shall be defined.
- **AGENT-REQ-716:** Revocation propagation latency shall be defined.
- **AGENT-REQ-717:** Policy propagation latency shall be defined.
- **AGENT-REQ-718:** Credential-status propagation latency shall be defined.
- **AGENT-REQ-719:** Connector-status propagation latency shall be defined.
- **AGENT-REQ-720:** Stale replicas shall not permit privileged operation.
- **AGENT-REQ-721:** Conflicting replicas shall resolve to the more restrictive state.
- **AGENT-REQ-722:** Partitioned components shall fail closed for privileged operations.
- **AGENT-REQ-723:** Offline operation shall require an explicitly approved bounded design.
- **AGENT-REQ-724:** Offline authorization caches shall have strict expiry and revocation assumptions.
- **AGENT-REQ-725:** Reconnection shall trigger state reconciliation.
- **AGENT-REQ-726:** Reconciliation shall preserve all transition evidence.
- **AGENT-REQ-727:** Reconciliation shall not fabricate missing transitions.
- **AGENT-REQ-728:** Propagation failure shall generate monitoring and incident events.
- **AGENT-REQ-729:** Distributed recovery shall restore a known governed state.
- **AGENT-REQ-730:** Propagation and partition tests shall verify denial, suspension and revocation behavior.

## 31. Evidence and Audit Integration

- **AGENT-REQ-731:** Every lifecycle transition shall produce an audit event.
- **AGENT-REQ-732:** Lifecycle audit events shall identify lifecycle and agent identifiers.
- **AGENT-REQ-733:** Lifecycle audit events shall identify prior and final state.
- **AGENT-REQ-734:** Lifecycle audit events shall identify actor and authority.
- **AGENT-REQ-735:** Lifecycle audit events shall identify time and policy version.
- **AGENT-REQ-736:** Lifecycle audit events shall identify result and reason.
- **AGENT-REQ-737:** Lifecycle audit events shall correlate with identity, authorization, session, workload and connector events.
- **AGENT-REQ-738:** Lifecycle audit events shall correlate with evidence handling where applicable.
- **AGENT-REQ-739:** Lifecycle audit events shall not contain secrets.
- **AGENT-REQ-740:** Lifecycle audit events shall minimize personal data.
- **AGENT-REQ-741:** Lifecycle audit events shall be append-only or tamper-evident where feasible.
- **AGENT-REQ-742:** Audit gaps shall remain visible.
- **AGENT-REQ-743:** Audit failure shall block or restrict privileged transitions.
- **AGENT-REQ-744:** Lifecycle evidence shall preserve provenance.
- **AGENT-REQ-745:** Lifecycle evidence shall preserve integrity references.
- **AGENT-REQ-746:** Lifecycle evidence shall preserve reviewer and approval records.
- **AGENT-REQ-747:** Suspension and revocation shall not erase evidence.
- **AGENT-REQ-748:** Archival shall preserve lifecycle evidence.
- **AGENT-REQ-749:** AI-generated lifecycle summaries shall not replace source audit events.
- **AGENT-REQ-750:** Lifecycle-evidence integration shall conform to `OBDIA-EVID-001`.

## 32. Incident, Risk and Exception Integration

- **AGENT-REQ-751:** Material lifecycle incidents shall have immutable identifiers.
- **AGENT-REQ-752:** Incidents shall identify affected agent and lifecycle states.
- **AGENT-REQ-753:** Incidents shall identify containment and transition actions.
- **AGENT-REQ-754:** Incidents shall preserve evidence.
- **AGENT-REQ-755:** Incidents shall trigger threat and risk reassessment.
- **AGENT-REQ-756:** Critical incidents shall support emergency suspension or revocation.
- **AGENT-REQ-757:** Risk records shall identify affected lifecycle states and transitions.
- **AGENT-REQ-758:** Residual risks shall have human owners.
- **AGENT-REQ-759:** An AI agent shall not accept lifecycle risk.
- **AGENT-REQ-760:** Risk acceptance shall not waive canonical prohibitions.
- **AGENT-REQ-761:** Critical unresolved risk shall block activation or operation.
- **AGENT-REQ-762:** Exceptions shall be narrow, attributable, time-bounded and revocable.
- **AGENT-REQ-763:** Exceptions shall identify compensating controls.
- **AGENT-REQ-764:** Exceptions shall identify affected transition guards.
- **AGENT-REQ-765:** Expired exceptions shall trigger denial, suspension or renewed review.
- **AGENT-REQ-766:** Exceptions shall not permit rebinding.
- **AGENT-REQ-767:** Exceptions shall not permit operation after revocation.
- **AGENT-REQ-768:** Repeated exceptions shall trigger architectural review.
- **AGENT-REQ-769:** Documents 26–28 remain forward dependencies for detailed risk and exception governance.
- **AGENT-REQ-770:** Lifecycle incident, risk and exception traceability shall be bidirectional.

## 33. Testing and Validation

- **AGENT-REQ-771:** Every material lifecycle requirement shall have testable acceptance criteria.
- **AGENT-REQ-772:** Lifecycle tests shall use synthetic identities and mock institutional context.
- **AGENT-REQ-773:** Tests shall verify the complete normal lifecycle sequence.
- **AGENT-REQ-774:** Tests shall verify mandatory `Bind` before `Authorize`.
- **AGENT-REQ-775:** Tests shall verify mandatory `Authorize` before `Activate`.
- **AGENT-REQ-776:** Tests shall verify mandatory `Activate` before `Operate`.
- **AGENT-REQ-777:** Tests shall verify every permitted transition.
- **AGENT-REQ-778:** Tests shall verify every prohibited transition.
- **AGENT-REQ-779:** Tests shall verify emergency suspension.
- **AGENT-REQ-780:** Tests shall verify terminal revocation.
- **AGENT-REQ-781:** Tests shall verify archival terminality.
- **AGENT-REQ-782:** Tests shall verify officer-rebinding denial.
- **AGENT-REQ-783:** Tests shall verify new identity requirement for a new officer.
- **AGENT-REQ-784:** Tests shall verify credential expiry.
- **AGENT-REQ-785:** Tests shall verify key rotation.
- **AGENT-REQ-786:** Tests shall verify binding-key compromise.
- **AGENT-REQ-787:** Tests shall verify authorization expiry.
- **AGENT-REQ-788:** Tests shall verify session invalidation.
- **AGENT-REQ-789:** Tests shall verify workload containment.
- **AGENT-REQ-790:** Tests shall verify connector disablement.
- **AGENT-REQ-791:** Tests shall verify evidence and audit preservation.
- **AGENT-REQ-792:** Tests shall verify monitoring failure.
- **AGENT-REQ-793:** Tests shall verify audit failure.
- **AGENT-REQ-794:** Tests shall verify authorization-service failure.
- **AGENT-REQ-795:** Tests shall verify revocation-service failure.
- **AGENT-REQ-796:** Tests shall verify clock failure.
- **AGENT-REQ-797:** Tests shall verify partial activation denial.
- **AGENT-REQ-798:** Tests shall verify partial suspension and revocation incidents.
- **AGENT-REQ-799:** Tests shall verify concurrent transition races.
- **AGENT-REQ-800:** Tests shall verify idempotent retries.
- **AGENT-REQ-801:** Tests shall verify distributed partition behavior.
- **AGENT-REQ-802:** Tests shall verify stale replica denial.
- **AGENT-REQ-803:** Tests shall verify review and reauthorization loops.
- **AGENT-REQ-804:** Tests shall verify suspension-to-review-to-authorization recovery.
- **AGENT-REQ-805:** Tests shall verify that revoked identities cannot return.
- **AGENT-REQ-806:** Tests shall verify archive retrieval without authority restoration.
- **AGENT-REQ-807:** Tests shall identify exact code, policy, model, connector, configuration and environment versions.
- **AGENT-REQ-808:** Failed tests shall remain visible.
- **AGENT-REQ-809:** Passing tests shall not be generalized beyond scope.
- **AGENT-REQ-810:** Document 22 remains a forward dependency for complete testing governance.

## 34. Traceability

- **AGENT-REQ-811:** Every lifecycle requirement shall trace to governing authority or threat.
- **AGENT-REQ-812:** Every state shall trace to entry and exit criteria.
- **AGENT-REQ-813:** Every transition shall trace to policy and evidence.
- **AGENT-REQ-814:** Every binding shall trace to officer, institution and cryptographic reference.
- **AGENT-REQ-815:** Every authorization shall trace to delegation and policy.
- **AGENT-REQ-816:** Every activation shall trace to readiness evidence.
- **AGENT-REQ-817:** Every operation shall trace to lifecycle and authorization state.
- **AGENT-REQ-818:** Every review shall trace to findings and disposition.
- **AGENT-REQ-819:** Every suspension shall trace to reason and containment.
- **AGENT-REQ-820:** Every revocation shall trace to authority and propagation evidence.
- **AGENT-REQ-821:** Every archive shall trace to completed revocation and retention.
- **AGENT-REQ-822:** Every credential shall trace to lifecycle state.
- **AGENT-REQ-823:** Every session and workload shall trace to agent lifecycle.
- **AGENT-REQ-824:** Every connector use shall trace to lifecycle and authorization.
- **AGENT-REQ-825:** Every evidence action shall trace to lifecycle state.
- **AGENT-REQ-826:** Every incident shall trace to affected transitions.
- **AGENT-REQ-827:** Every risk and exception shall trace to affected requirements.
- **AGENT-REQ-828:** Every test shall trace to lifecycle requirements and threats.
- **AGENT-REQ-829:** Traceability shall be bidirectional.
- **AGENT-REQ-830:** Broken traceability affecting identity, authorization, suspension, revocation, evidence or archive shall be blocking.
- **AGENT-REQ-831:** Planned lifecycle controls shall not be represented as implemented.
- **AGENT-REQ-832:** Implemented lifecycle controls shall not be represented as validated without evidence.
- **AGENT-REQ-833:** Superseded lifecycle schemas shall identify successors.
- **AGENT-REQ-834:** Traceability shall use immutable identifiers.
- **AGENT-REQ-835:** Traceability records shall not contain secrets.
- **AGENT-REQ-836:** Traceability records shall minimize personal and case data.
- **AGENT-REQ-837:** Baseline freeze shall validate lifecycle traceability across documents 01–30.
- **AGENT-REQ-838:** No document above 30 shall acquire normative lifecycle authority by implication.

## 35. Lifecycle Review Gate

- **AGENT-REQ-839:** Lifecycle-model review shall identify the exact document and commit.
- **AGENT-REQ-840:** Review shall verify preservation of both legacy lifecycle sequences.
- **AGENT-REQ-841:** Review shall verify inclusion of mandatory `Bind`.
- **AGENT-REQ-842:** Review shall verify every state definition.
- **AGENT-REQ-843:** Review shall verify entry and exit criteria.
- **AGENT-REQ-844:** Review shall verify the permitted transition matrix.
- **AGENT-REQ-845:** Review shall verify prohibited transitions.
- **AGENT-REQ-846:** Review shall verify officer-rebinding prohibition.
- **AGENT-REQ-847:** Review shall verify authorization and activation ordering.
- **AGENT-REQ-848:** Review shall verify suspension and recovery behavior.
- **AGENT-REQ-849:** Review shall verify terminal revocation and archival.
- **AGENT-REQ-850:** Review shall verify credential, session, workload and connector propagation.
- **AGENT-REQ-851:** Review shall verify evidence and audit preservation.
- **AGENT-REQ-852:** Review shall verify privacy and rights controls.
- **AGENT-REQ-853:** Review shall verify failure, containment and recovery.
- **AGENT-REQ-854:** Review shall verify concurrency and distributed consistency.
- **AGENT-REQ-855:** Review shall verify testing and traceability.
- **AGENT-REQ-856:** Review shall identify residual risks and forward dependencies.
- **AGENT-REQ-857:** Critical findings shall block progression.
- **AGENT-REQ-858:** Material post-review change shall invalidate affected review evidence.
- **AGENT-REQ-859:** Role concentration shall be disclosed.
- **AGENT-REQ-860:** Internal review shall not be represented as independent external assurance.
- **AGENT-REQ-861:** Security Reviewer shall assess binding, authorization, revocation and failure behavior.
- **AGENT-REQ-862:** Implementation Reviewer shall assess state-machine implementation where implementation exists.
- **AGENT-REQ-863:** Privacy and Governance Reviewer shall assess personal-data and rights impacts.
- **AGENT-REQ-864:** Release Reviewer shall assess public representations of lifecycle capability.
- **AGENT-REQ-865:** Project Founder approval shall not substitute for required specialist review.
- **AGENT-REQ-866:** Automated findings shall remain advisory until adopted by an authorized human reviewer.
- **AGENT-REQ-867:** Review non-applicability shall have rationale.
- **AGENT-REQ-868:** Approval for Draft incorporation shall not establish operational agent approval.

## 36. Minimum Validation Checklist

Before approval of this model or a lifecycle implementation, confirm:

- [ ] unique agent and lifecycle identifiers exist;
- [ ] accountable institution is identified;
- [ ] exactly one human officer is bound;
- [ ] cryptographic binding is verifiable;
- [ ] human credentials are not transferred;
- [ ] `Provision → Bind → Authorize → Activate → Operate` ordering is enforced;
- [ ] review, suspension, revocation and archival are governed;
- [ ] every transition is logged and auditable;
- [ ] transition authority and exact policy versions are recorded;
- [ ] canonical prohibitions remain non-authorizable;
- [ ] suspended agents cannot operate;
- [ ] suspended agents require review and fresh authorization;
- [ ] revoked agents cannot reactivate;
- [ ] archived agents cannot reactivate or rebind;
- [ ] officer change requires a new agent identity;
- [ ] credentials, sessions, workloads and connectors follow lifecycle state;
- [ ] evidence and audit records survive suspension and revocation;
- [ ] monitoring and containment are operationally testable;
- [ ] failure paths are fail-closed;
- [ ] concurrency and idempotency controls are defined;
- [ ] distributed-state conflicts resolve restrictively;
- [ ] privacy and fundamental-rights impacts are reviewed;
- [ ] negative and race-condition tests exist;
- [ ] forward dependencies 20–30 are recorded where applicable;
- [ ] Project Founder approval exists before status becomes Approved.


## 37. Limitations

- This model does not provision or deploy a real agent.
- It does not select an identity provider, credential format, key-management platform, policy engine, runtime, model provider, connector platform or monitoring system.
- It does not create legal authority or institutional investigative authority.
- It does not prove production readiness, certification, accreditation or legal compliance.
- It does not replace the Identity and Delegation Model, Authorization Model, Evidence Model or Connector Security Policy.
- Suspension and revocation depend on correctly implemented enforcement and propagation.
- Distributed systems may have non-zero propagation latency that must be measured and bounded.
- Internal lifecycle review is not independent certification.
- Documents 20–30 remain forward dependencies where they govern connectors, coding, testing, structure, versioning, risk, exceptions, changes and compliance.


## 38. Change Control

Every material change shall identify rationale, affected states and transitions, identity and binding impact, authorization impact, credential and connector impact, security and privacy impact, evidence impact, migration, validation, rollback, authority and version effect.

- **AGENT-REQ-869:** Editorial corrections shall use a patch version when meaning is unchanged.
- **AGENT-REQ-870:** Backward-compatible substantive additions shall use a minor version.
- **AGENT-REQ-871:** Incompatible lifecycle-state or transition changes shall use a major version.
- **AGENT-REQ-872:** Material lifecycle-architecture decisions shall require an ADR where applicable.
- **AGENT-REQ-873:** Changes shall require Project Founder approval.
- **AGENT-REQ-874:** Changes shall receive Security Reviewer assessment.
- **AGENT-REQ-875:** Identity, authorization, evidence, connector, implementation and privacy impacts shall receive specialist review where applicable.
- **AGENT-REQ-876:** Changes shall identify affected policies, credentials, sessions, workloads, connectors, tests and records.
- **AGENT-REQ-877:** Changes shall include migration and compatibility analysis.
- **AGENT-REQ-878:** Changes shall include validation and rollback analysis.
- **AGENT-REQ-879:** Changes shall not retroactively fabricate transition, authorization or review evidence.
- **AGENT-REQ-880:** Historical lifecycle records shall not be silently rewritten.
- **AGENT-REQ-881:** Agent and lifecycle identifiers shall not be reused.
- **AGENT-REQ-882:** Migration shall preserve prior state and transition history.
- **AGENT-REQ-883:** Forward-dependency reconciliation shall occur before this model becomes Approved.

## 39. Consolidation Record

Version 1.0.0 consolidates the two existing Agent Lifecycle Model variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-AGENT-001`;
- normalizes the authoritative filename to `19_AGENT_LIFECYCLE_MODEL.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves `Provision`, `Authorize`, `Activate`, `Operate`, `Review`, `Suspend`, `Revoke` and `Archive`;
- preserves the Enterprise variant’s mandatory `Bind` state;
- establishes the authoritative sequence `Provision → Bind → Authorize → Activate → Operate → Review → Suspend → Revoke → Archive`;
- preserves the original requirement that every transition be logged and auditable;
- adds lifecycle objects, invariants, transition events, permitted and prohibited transitions, entry and exit criteria, reauthorization, rebinding prohibition, credential, session, workload, runtime, connector, monitoring, containment, failure, distributed-consistency, evidence, risk, testing and traceability controls;
- defines suspension as reversible only through review and fresh authorization;
- defines revocation as terminal for the agent identity and archival as the terminal retained state;
- identifies documents 20–30 as forward dependencies where applicable;
- treats Enterprise and non-Enterprise legacy files as source variants of the same immutable document;
- creates no operational agent, credential, legal authority, implementation or deployment claim.

## 40. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Security Reviewer; approval reserved to Project Founder | Constitutional consolidation of both legacy lifecycle sequences: preserved every state and audit requirement; added mandatory binding, state-machine semantics, transition guards, suspension, terminal revocation, archival, propagation, validation and traceability. |
| 1.0.1 | 2026-08-06 | Draft | Security Reviewer; approval reserved to Project Founder | Reconciled obsolete forward-dependency filenames with the final canonical Knowledge Pack 01–30 filenames; no requirement text, authority, lifecycle status, implementation state or validation claim changed. |
