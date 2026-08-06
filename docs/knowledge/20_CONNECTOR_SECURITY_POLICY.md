# CONNECTOR SECURITY POLICY

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-CONN-001 |
| **Title** | Connector Security Policy |
| **Version** | 1.0.1 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Security Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define mandatory security, identity, authorization, data, provenance, isolation, audit, revocation, lifecycle, validation and failure-handling requirements for every OBDIA connector and mediated external tool. |
| **Scope** | Connectors to authorized APIs, authenticated Internet services, cloud platforms, threat-intelligence platforms, digital-evidence systems, communication platforms, blockchains, smart contracts, controlled research environments, local mock services and other approved external systems within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `06_IDENTITY_AND_DELEGATION_MODEL.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `15_RELEASE_AND_PUBLICATION_POLICY.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; `18_EVIDENCE_MODEL.md`; `19_AGENT_LIFECYCLE_MODEL.md`; forward dependencies `21_SECURE_CODING_STANDARD.md`, `22_TESTING_STANDARD.md`, `24_GITHUB_REPOSITORY_STANDARD.md`, `25_DOCUMENT_VERSIONING_POLICY.md`, `26_AI_RISK_REGISTER.md`, `27_RESEARCH_BACKLOG.md`, `28_ASSUMPTIONS_REGISTER.md`, `29_DECISION_LOG_POLICY.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-ARCH-001`; `OBDIA-ID-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-REL-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001`; `OBDIA-EVID-001`; `OBDIA-AGENT-001` |
| **Cross-References** | Documents 21–30; accepted ADRs; connector manifests; identity, authorization, policy, evidence, audit, incident, risk, exception, change, testing and release records |
| **Assumptions** | Public demonstrations and research use synthetic datasets, mock services, testnets, local laboratories, read-only APIs, lawful public sources, archived-authorized resources or otherwise explicitly authorized systems. |
| **Constraints** | Connectors shall not authorize unlawful access, credential theft, malware deployment, offensive cyber operations, uncontrolled criminal-infrastructure interaction, real-person public investigations, autonomous legal decisions or independent AI legal authority. |
| **Security Considerations** | Connectors cross trust boundaries and can expose credentials, permit privilege escalation, leak data, import malicious content, cause confused-deputy behavior, corrupt evidence, bypass policy, create uncontrolled side effects, expand network reach or conceal third-party failure. |
| **Validation Criteria** | Every connector has a unique manifest, distinct machine identity, bounded authorization, least-privileged permissions, validated schemas, controlled egress, provenance preservation, failure isolation, auditable use, revocation, negative tests and bidirectional traceability. |
| **Implementation Relationship** | This policy defines connector requirements. It does not approve a specific connector, vendor, protocol, API, credential, endpoint, production integration or operational deployment. |

---

## 1. Purpose

This document governs every connector through which an OBDIA agent or workload interacts with an external or separately governed system.

The two legacy source variants required connectors to:

- authenticate securely;
- enforce least privilege or minimize permissions;
- validate outputs;
- respect authorization or authorization boundaries;
- isolate failures;
- log usage or generate audit logs;
- preserve provenance;
- support revocation.

This consolidation preserves all eight controls and turns them into a complete connector governance and security model.

A connector is a constrained integration boundary. It is not an independent authority, a substitute for human authorization, or permission for an agent to interact with any technically reachable system.


## 2. Fundamental Connector Rules

- **CONN-REQ-001:** Every connector shall authenticate securely.
- **CONN-REQ-002:** Every connector shall enforce least privilege.
- **CONN-REQ-003:** Every connector shall minimize permissions.
- **CONN-REQ-004:** Every connector shall validate outputs.
- **CONN-REQ-005:** Every connector shall respect authorization boundaries.
- **CONN-REQ-006:** Every connector shall isolate failures.
- **CONN-REQ-007:** Every connector shall generate audit logs.
- **CONN-REQ-008:** Every connector shall log governed usage.
- **CONN-REQ-009:** Every connector shall preserve provenance.
- **CONN-REQ-010:** Every connector shall support revocation.
- **CONN-REQ-011:** Every connector shall default to disabled until explicitly authorized.
- **CONN-REQ-012:** Every connector shall default to deny for unrecognized operations.
- **CONN-REQ-013:** Read-only access shall be preferred whenever sufficient.
- **CONN-REQ-014:** A connector shall not create independent AI legal authority.
- **CONN-REQ-015:** A connector shall not convert technical reachability into authorization.
- **CONN-REQ-016:** A connector shall not bypass officer-agent binding.
- **CONN-REQ-017:** A connector shall not receive reusable human credentials.
- **CONN-REQ-018:** A connector shall not broaden case, purpose, jurisdiction, time, tool or data scope.
- **CONN-REQ-019:** A connector shall not authorize canonical prohibitions.
- **CONN-REQ-020:** Model output shall not authorize connector execution.
- **CONN-REQ-021:** Instructions received through connector content shall not override system, policy or human authority.
- **CONN-REQ-022:** Missing, stale, conflicting or unverifiable connector context shall result in denial, suspension or quarantine.
- **CONN-REQ-023:** Connector controls shall fail closed for privileged, evidentiary or side-effecting operations.
- **CONN-REQ-024:** Connector activity shall remain attributable to officer, agent, workload, institution and policy where applicable.
- **CONN-REQ-025:** Repository merge shall not establish connector approval.
- **CONN-REQ-026:** Passing automated checks shall not establish connector approval.
- **CONN-REQ-027:** Every material connector control shall be testable.
- **CONN-REQ-028:** Every connector shall have an authorized human owner.
- **CONN-REQ-029:** Every connector shall have a documented shutdown and revocation path.
- **CONN-REQ-030:** Unresolved critical connector risk shall block enablement or release.

## 3. Connector Definition and Boundary

- **CONN-REQ-031:** A connector shall be a separately identifiable integration component.
- **CONN-REQ-032:** A connector shall mediate access to one or more explicitly approved external capabilities.
- **CONN-REQ-033:** A connector shall have a unique immutable connector identifier.
- **CONN-REQ-034:** A connector shall have one canonical name.
- **CONN-REQ-035:** A connector shall have one authoritative manifest.
- **CONN-REQ-036:** A connector boundary shall distinguish caller, connector, network, external system and returned content.
- **CONN-REQ-037:** A connector boundary shall identify all trust crossings.
- **CONN-REQ-038:** A connector boundary shall identify all policy enforcement points.
- **CONN-REQ-039:** A connector boundary shall identify all credential boundaries.
- **CONN-REQ-040:** A connector boundary shall identify all data ingress and egress paths.
- **CONN-REQ-041:** A connector boundary shall identify all evidence and provenance transformations.
- **CONN-REQ-042:** A connector boundary shall identify all failure and containment paths.
- **CONN-REQ-043:** A connector boundary shall identify all human approval points.
- **CONN-REQ-044:** A connector shall remain distinct from a model.
- **CONN-REQ-045:** A connector shall remain distinct from an agent identity.
- **CONN-REQ-046:** A connector shall remain distinct from a workload identity.
- **CONN-REQ-047:** A connector shall remain distinct from an authorization decision.
- **CONN-REQ-048:** A connector shall remain distinct from the external service it accesses.
- **CONN-REQ-049:** A connector shall remain distinct from the data returned by the external service.
- **CONN-REQ-050:** A connector shall remain distinct from a local tool that has no external trust boundary unless policy classifies the tool as connector-equivalent.
- **CONN-REQ-051:** Connector-equivalent tools shall meet this policy when they access separately governed resources or perform material side effects.
- **CONN-REQ-052:** Connector naming shall conform to `OBDIA-NAME-001`.
- **CONN-REQ-053:** Connector boundaries shall be diagrammed when a diagram materially improves review.
- **CONN-REQ-054:** Boundary ambiguity shall block implementation.
- **CONN-REQ-055:** Boundary changes shall trigger threat-model and authorization review.

## 4. Connector Manifest

- **CONN-REQ-056:** Every connector shall have a version-controlled manifest.
- **CONN-REQ-057:** The connector manifest shall identify connector identifier.
- **CONN-REQ-058:** The connector manifest shall identify connector name and title.
- **CONN-REQ-059:** The connector manifest shall identify owner.
- **CONN-REQ-060:** The connector manifest shall identify approval authority.
- **CONN-REQ-061:** The connector manifest shall identify implementation status.
- **CONN-REQ-062:** The connector manifest shall identify operational lifecycle state.
- **CONN-REQ-063:** The connector manifest shall identify connector class.
- **CONN-REQ-064:** The connector manifest shall identify external system or service class.
- **CONN-REQ-065:** The connector manifest shall identify approved endpoints or endpoint patterns.
- **CONN-REQ-066:** The connector manifest shall identify supported protocols.
- **CONN-REQ-067:** The connector manifest shall identify supported operations.
- **CONN-REQ-068:** The connector manifest shall identify explicitly prohibited operations.
- **CONN-REQ-069:** The connector manifest shall identify authentication method.
- **CONN-REQ-070:** The connector manifest shall identify machine-identity reference.
- **CONN-REQ-071:** The connector manifest shall identify credential scope and audience.
- **CONN-REQ-072:** The connector manifest shall identify authorization-policy references.
- **CONN-REQ-073:** The connector manifest shall identify case, purpose, jurisdiction and data constraints.
- **CONN-REQ-074:** The connector manifest shall identify read-only or write-capable behavior.
- **CONN-REQ-075:** The connector manifest shall identify side effects.
- **CONN-REQ-076:** The connector manifest shall identify data ingress schema.
- **CONN-REQ-077:** The connector manifest shall identify data egress schema.
- **CONN-REQ-078:** The connector manifest shall identify output-validation rules.
- **CONN-REQ-079:** The connector manifest shall identify provenance-generation rules.
- **CONN-REQ-080:** The connector manifest shall identify logging and audit requirements.
- **CONN-REQ-081:** The connector manifest shall identify rate, volume, retry and timeout limits.
- **CONN-REQ-082:** The connector manifest shall identify network egress controls.
- **CONN-REQ-083:** The connector manifest shall identify dependency and supply-chain references.
- **CONN-REQ-084:** The connector manifest shall identify failure and isolation behavior.
- **CONN-REQ-085:** The connector manifest shall identify suspension and revocation mechanisms.
- **CONN-REQ-086:** The connector manifest shall identify retention and deletion behavior.
- **CONN-REQ-087:** The connector manifest shall identify privacy and classification constraints.
- **CONN-REQ-088:** The connector manifest shall identify known limitations.
- **CONN-REQ-089:** The connector manifest shall identify threat-model references.
- **CONN-REQ-090:** The connector manifest shall identify validation and test evidence.
- **CONN-REQ-091:** The connector manifest shall identify release eligibility.
- **CONN-REQ-092:** The connector manifest shall identify version and change history.
- **CONN-REQ-093:** The connector manifest shall not contain secrets.
- **CONN-REQ-094:** Malformed or incomplete manifests shall block connector enablement.
- **CONN-REQ-095:** Manifest changes shall trigger risk-based review.

## 5. Connector Classification

- **CONN-REQ-096:** Every connector shall have a controlled connector class.
- **CONN-REQ-097:** Connector classes shall distinguish local mock, internal service, public API, authenticated Internet service, cloud service, threat-intelligence platform, digital-evidence system, communication platform, blockchain, smart contract, controlled research environment and other approved classes.
- **CONN-REQ-098:** Connector classes shall identify read-only or side-effecting capability.
- **CONN-REQ-099:** Connector classes shall identify synchronous, asynchronous or streaming behavior.
- **CONN-REQ-100:** Connector classes shall identify data sensitivity.
- **CONN-REQ-101:** Connector classes shall identify whether evidence may be acquired.
- **CONN-REQ-102:** Connector classes shall identify whether personal data may be processed.
- **CONN-REQ-103:** Connector classes shall identify whether external state may be changed.
- **CONN-REQ-104:** Connector classes shall identify whether funds or digital assets may be affected.
- **CONN-REQ-105:** Connector classes shall identify whether administrative capability exists.
- **CONN-REQ-106:** Connector classes shall identify risk tier.
- **CONN-REQ-107:** Connector risk tier shall determine minimum review and testing depth.
- **CONN-REQ-108:** Unknown connector classes shall remain disabled.
- **CONN-REQ-109:** Class changes shall trigger reauthorization.
- **CONN-REQ-110:** Class changes shall trigger threat-model review.
- **CONN-REQ-111:** Class changes shall not broaden permissions silently.
- **CONN-REQ-112:** Public accessibility shall not imply low risk.
- **CONN-REQ-113:** Read-only classification shall be verified technically, not inferred from documentation alone.
- **CONN-REQ-114:** A connector capable of state change shall not be represented as read-only.
- **CONN-REQ-115:** A connector using a testnet shall identify the exact test network.
- **CONN-REQ-116:** A connector using a mock service shall identify the mock boundary.
- **CONN-REQ-117:** A connector using archived-authorized content shall identify archive authority.
- **CONN-REQ-118:** A connector for operationally isolated Dark Web research shall identify isolation and non-interaction constraints.
- **CONN-REQ-119:** Class labels shall not imply compliance or production readiness.
- **CONN-REQ-120:** Connector classification history shall be retained.

## 6. Operational Connector Lifecycle

- **CONN-REQ-121:** Connector operational states shall be `Defined`, `Registered`, `Authorized`, `Enabled`, `Suspended`, `Revoked` and `Retired`.
- **CONN-REQ-122:** Connector operational state shall remain distinct from repository document lifecycle state.
- **CONN-REQ-123:** `Defined` shall indicate that design information exists but registration is incomplete.
- **CONN-REQ-124:** `Registered` shall indicate that identity, manifest and ownership records exist.
- **CONN-REQ-125:** `Authorized` shall indicate an explicit bounded authorization exists.
- **CONN-REQ-126:** `Enabled` shall indicate the connector may accept authorized operations.
- **CONN-REQ-127:** `Suspended` shall indicate reversible containment preventing ordinary use.
- **CONN-REQ-128:** `Revoked` shall indicate terminal withdrawal of connector authority or credentials.
- **CONN-REQ-129:** `Retired` shall indicate terminal retained status after revocation and archival.
- **CONN-REQ-130:** A connector shall not transition from `Defined` directly to `Enabled`.
- **CONN-REQ-131:** A connector shall not transition from `Registered` directly to `Enabled` without authorization.
- **CONN-REQ-132:** A connector shall not transition from `Suspended` directly to `Enabled` without review and fresh authorization.
- **CONN-REQ-133:** A revoked connector identity shall not return to an active state.
- **CONN-REQ-134:** A retired connector identity shall not return to an active state.
- **CONN-REQ-135:** Every lifecycle transition shall create an immutable event.
- **CONN-REQ-136:** Every transition shall identify actor, authority, reason, time and evidence.
- **CONN-REQ-137:** Every transition shall identify affected credentials, sessions and endpoints.
- **CONN-REQ-138:** Emergency suspension shall be supported.
- **CONN-REQ-139:** Emergency revocation shall be supported.
- **CONN-REQ-140:** Lifecycle conflict shall resolve to the more restrictive state.
- **CONN-REQ-141:** Unknown lifecycle state shall result in disabled behavior.
- **CONN-REQ-142:** Connector enablement shall validate current manifest and policy versions.
- **CONN-REQ-143:** Connector enablement shall validate monitoring and audit readiness.
- **CONN-REQ-144:** Connector suspension shall stop new operations.
- **CONN-REQ-145:** Connector revocation shall invalidate credentials and sessions.
- **CONN-REQ-146:** Connector retirement shall preserve audit and provenance records.
- **CONN-REQ-147:** Operational state shall be observable to authorized operators.
- **CONN-REQ-148:** Operational state changes shall not erase historical events.
- **CONN-REQ-149:** Operational lifecycle tests shall cover every permitted and prohibited transition.
- **CONN-REQ-150:** Document 19 agent lifecycle state shall constrain connector availability.

## 7. Connector Identity

- **CONN-REQ-151:** Every connector shall use a distinct machine identity where feasible.
- **CONN-REQ-152:** Connector identities shall be institutionally issued or explicitly approved.
- **CONN-REQ-153:** Connector identities shall be distinguishable from human, agent and workload identities.
- **CONN-REQ-154:** Connector identities shall have unique immutable identifiers.
- **CONN-REQ-155:** Shared connector identities shall require an approved exception.
- **CONN-REQ-156:** Connector identity issuers shall be documented.
- **CONN-REQ-157:** Connector identity audience shall be documented.
- **CONN-REQ-158:** Connector identity purpose shall be documented.
- **CONN-REQ-159:** Connector identity scope shall be documented.
- **CONN-REQ-160:** Connector identity validity shall be documented.
- **CONN-REQ-161:** Connector identity assurance limitations shall be documented.
- **CONN-REQ-162:** Connector identity shall not contain human credentials.
- **CONN-REQ-163:** Connector identity shall not impersonate the bound officer.
- **CONN-REQ-164:** Connector identity shall remain attributable to the invoking agent and workload.
- **CONN-REQ-165:** Connector identity status shall be verified before every protected use or at bounded checkpoints.
- **CONN-REQ-166:** Suspended connector identities shall not authenticate.
- **CONN-REQ-167:** Revoked connector identities shall not authenticate.
- **CONN-REQ-168:** Expired connector identities shall not authenticate.
- **CONN-REQ-169:** Unknown connector identities shall be denied.
- **CONN-REQ-170:** Connector identity rebinding shall not preserve prior authority silently.
- **CONN-REQ-171:** Connector identity replacement shall invalidate prior sessions.
- **CONN-REQ-172:** Connector identity compromise shall trigger suspension or revocation.
- **CONN-REQ-173:** Connector identity logs shall not expose secrets.
- **CONN-REQ-174:** Connector identity metadata shall minimize personal data.
- **CONN-REQ-175:** Connector identity implementation shall conform to `OBDIA-ID-001`.

## 8. Authentication and Credential Security

- **CONN-REQ-176:** Connector authentication shall use mechanisms proportionate to risk.
- **CONN-REQ-177:** Static long-lived credentials shall be avoided.
- **CONN-REQ-178:** Short-lived credentials shall be preferred where supported.
- **CONN-REQ-179:** Credential audience shall be restricted.
- **CONN-REQ-180:** Credential scope shall be restricted.
- **CONN-REQ-181:** Credential validity shall be bounded.
- **CONN-REQ-182:** Credential issuance shall be attributable.
- **CONN-REQ-183:** Credential retrieval shall be authenticated and authorized.
- **CONN-REQ-184:** Credential storage shall use approved secret protections.
- **CONN-REQ-185:** Credentials shall not be stored in source code.
- **CONN-REQ-186:** Credentials shall not be stored in documentation.
- **CONN-REQ-187:** Credentials shall not be stored in prompts or fixtures.
- **CONN-REQ-188:** Credentials shall not appear in logs or evidence metadata.
- **CONN-REQ-189:** Development, test, demonstration and operational credentials shall remain distinct.
- **CONN-REQ-190:** Credential rotation shall be supported.
- **CONN-REQ-191:** Credential revocation shall be supported.
- **CONN-REQ-192:** Credential compromise shall trigger immediate containment.
- **CONN-REQ-193:** Credential expiry shall fail closed.
- **CONN-REQ-194:** Credential refresh shall not broaden scope.
- **CONN-REQ-195:** Credential fallback shall not use human credentials.
- **CONN-REQ-196:** Authentication success shall not be treated as authorization success.
- **CONN-REQ-197:** Mutual authentication shall be used where risk requires it.
- **CONN-REQ-198:** Server or endpoint identity shall be verified.
- **CONN-REQ-199:** Certificate or key validation failure shall block connection.
- **CONN-REQ-200:** Deprecated authentication methods shall have migration plans.
- **CONN-REQ-201:** Credential inventory shall remain accurate.
- **CONN-REQ-202:** Orphaned connector credentials shall be detected and revoked.
- **CONN-REQ-203:** Secret scanning shall include repository history.
- **CONN-REQ-204:** Credential tests shall cover expiry, rotation, revocation, substitution and unauthorized audience.
- **CONN-REQ-205:** Authentication limitations shall remain visible.

## 9. Authorization Enforcement

- **CONN-REQ-206:** Every connector operation shall require explicit authorization.
- **CONN-REQ-207:** Connector authorization shall default to deny.
- **CONN-REQ-208:** Connector authorization shall evaluate officer identity.
- **CONN-REQ-209:** Connector authorization shall evaluate institution.
- **CONN-REQ-210:** Connector authorization shall evaluate agent and workload identity.
- **CONN-REQ-211:** Connector authorization shall evaluate case or mandate.
- **CONN-REQ-212:** Connector authorization shall evaluate purpose.
- **CONN-REQ-213:** Connector authorization shall evaluate jurisdiction where applicable.
- **CONN-REQ-214:** Connector authorization shall evaluate action.
- **CONN-REQ-215:** Connector authorization shall evaluate destination.
- **CONN-REQ-216:** Connector authorization shall evaluate tool and connector identifier.
- **CONN-REQ-217:** Connector authorization shall evaluate data scope.
- **CONN-REQ-218:** Connector authorization shall evaluate time validity.
- **CONN-REQ-219:** Connector authorization shall evaluate policy version.
- **CONN-REQ-220:** Connector authorization shall evaluate connector operational state.
- **CONN-REQ-221:** Connector authorization shall evaluate credential and revocation state.
- **CONN-REQ-222:** Connector authorization shall evaluate human approval where required.
- **CONN-REQ-223:** Connector authorization shall evaluate risk and trust state where required.
- **CONN-REQ-224:** Permit shall apply only to the exact returned scope and obligations.
- **CONN-REQ-225:** Denied operations shall not be retried through a weaker path.
- **CONN-REQ-226:** Connector retries shall preserve the original authorization decision or obtain a fresh valid decision.
- **CONN-REQ-227:** Connector redirects shall not bypass authorization.
- **CONN-REQ-228:** Connector pagination shall not broaden authorized scope.
- **CONN-REQ-229:** Connector batching shall not combine differently authorized operations silently.
- **CONN-REQ-230:** Connector write operations shall require stronger controls than equivalent reads.
- **CONN-REQ-231:** Administrative connector operations shall require separate authorization.
- **CONN-REQ-232:** Cross-case connector operations shall be denied by default.
- **CONN-REQ-233:** Cross-purpose connector operations shall be denied by default.
- **CONN-REQ-234:** Model output shall not directly authorize connector operations.
- **CONN-REQ-235:** Connector authorization shall conform to `OBDIA-AUTH-001`.

## 10. Least Privilege and Permission Minimization

- **CONN-REQ-236:** Connector permissions shall be limited to required operations.
- **CONN-REQ-237:** Connector permissions shall be limited to required resources.
- **CONN-REQ-238:** Connector permissions shall be limited to required data fields.
- **CONN-REQ-239:** Connector permissions shall be limited to required accounts, tenants, projects, chains or services.
- **CONN-REQ-240:** Connector permissions shall be time-bounded where supported.
- **CONN-REQ-241:** Connector permissions shall be case-bounded where applicable.
- **CONN-REQ-242:** Connector permissions shall be purpose-bounded.
- **CONN-REQ-243:** Connector permissions shall be environment-bounded.
- **CONN-REQ-244:** Read-only permission shall be preferred when sufficient.
- **CONN-REQ-245:** List permission shall not imply read-content permission.
- **CONN-REQ-246:** Read permission shall not imply write permission.
- **CONN-REQ-247:** Write permission shall not imply delete permission.
- **CONN-REQ-248:** Delete permission shall require exceptional explicit authority.
- **CONN-REQ-249:** Export permission shall be separately controlled.
- **CONN-REQ-250:** Administrative permission shall be separately controlled.
- **CONN-REQ-251:** Wildcard permissions shall be prohibited unless justified and approved.
- **CONN-REQ-252:** Permission inheritance shall not broaden authority silently.
- **CONN-REQ-253:** Unused permissions shall be removed.
- **CONN-REQ-254:** Permission reviews shall occur at risk-based intervals.
- **CONN-REQ-255:** Permission drift shall be detectable.
- **CONN-REQ-256:** Permission changes shall be attributable.
- **CONN-REQ-257:** Permission changes shall trigger validation.
- **CONN-REQ-258:** Connector compromise blast radius shall be minimized.
- **CONN-REQ-259:** Separate connector identities should be used for materially different permission sets.
- **CONN-REQ-260:** Emergency access shall be narrow and time-bounded.
- **CONN-REQ-261:** Emergency access shall not authorize canonical prohibitions.
- **CONN-REQ-262:** Permission exceptions shall have owners and expiry.
- **CONN-REQ-263:** Expired exceptions shall result in restriction or suspension.
- **CONN-REQ-264:** Least-privilege tests shall verify denied out-of-scope operations.
- **CONN-REQ-265:** Vendor default permissions shall not be accepted without review.

## 11. Endpoint and Destination Control

- **CONN-REQ-266:** Connector destinations shall be explicitly approved.
- **CONN-REQ-267:** Approved endpoints shall be recorded in the manifest.
- **CONN-REQ-268:** Endpoint patterns shall be narrow and reviewable.
- **CONN-REQ-269:** Dynamic destination selection shall be constrained.
- **CONN-REQ-270:** User-controlled destinations shall be rejected or validated against approved policy.
- **CONN-REQ-271:** Model-generated destinations shall not be trusted.
- **CONN-REQ-272:** Redirect destinations shall be revalidated.
- **CONN-REQ-273:** DNS resolution shall not bypass destination policy.
- **CONN-REQ-274:** Resolved addresses shall be checked against prohibited or unexpected ranges where applicable.
- **CONN-REQ-275:** Loopback, link-local, metadata-service and internal administrative endpoints shall be denied unless explicitly required and isolated.
- **CONN-REQ-276:** Server-side request forgery risks shall be included in threat modeling.
- **CONN-REQ-277:** Protocol changes shall trigger revalidation.
- **CONN-REQ-278:** Port changes shall trigger revalidation.
- **CONN-REQ-279:** TLS hostname and certificate validation shall be enforced where applicable.
- **CONN-REQ-280:** Certificate pinning may be used when justified and maintainable.
- **CONN-REQ-281:** Endpoint failover shall be pre-approved.
- **CONN-REQ-282:** Failover endpoints shall be no more permissive.
- **CONN-REQ-283:** Unknown endpoints shall be denied.
- **CONN-REQ-284:** Endpoint inventory shall remain current.
- **CONN-REQ-285:** Endpoint changes shall be attributable.
- **CONN-REQ-286:** Endpoint deprecation shall identify migration.
- **CONN-REQ-287:** Endpoint health failure shall not trigger uncontrolled direct access.
- **CONN-REQ-288:** Proxy use shall be documented.
- **CONN-REQ-289:** Proxy bypass shall be prohibited unless approved.
- **CONN-REQ-290:** Destination-control tests shall include redirects, DNS changes, private ranges and malformed URLs.

## 12. Network Egress and Segmentation

- **CONN-REQ-291:** Connector network egress shall be explicitly controlled.
- **CONN-REQ-292:** Default egress shall deny unnecessary destinations.
- **CONN-REQ-293:** Connector egress rules shall align with approved endpoints.
- **CONN-REQ-294:** Connector network identity shall be attributable where feasible.
- **CONN-REQ-295:** Connector traffic shall be segmented from unrelated workloads.
- **CONN-REQ-296:** Development, test, demonstration and operational connector traffic shall remain separated.
- **CONN-REQ-297:** High-risk research connectors shall use operational isolation.
- **CONN-REQ-298:** Dark Web research connectors shall not interact with uncontrolled criminal infrastructure.
- **CONN-REQ-299:** Tor or equivalent anonymity networks shall not be used unless an approved isolated research design exists.
- **CONN-REQ-300:** Public demonstrations shall not require access to real criminal infrastructure.
- **CONN-REQ-301:** Connector ingress shall be denied unless explicitly required.
- **CONN-REQ-302:** Callback endpoints shall be authenticated and authorized.
- **CONN-REQ-303:** Webhook sources shall be verified.
- **CONN-REQ-304:** Webhook replay shall be controlled.
- **CONN-REQ-305:** Network encryption shall be used where appropriate.
- **CONN-REQ-306:** Network encryption shall not replace authorization.
- **CONN-REQ-307:** Network timeout behavior shall be bounded.
- **CONN-REQ-308:** Network failure shall not broaden egress.
- **CONN-REQ-309:** Network fallback shall not bypass policy enforcement.
- **CONN-REQ-310:** Egress telemetry shall support incident response.
- **CONN-REQ-311:** Egress telemetry shall minimize personal data.
- **CONN-REQ-312:** Unexpected destinations shall trigger alerting.
- **CONN-REQ-313:** Connector quarantine shall support egress shutdown.
- **CONN-REQ-314:** Network-control configuration shall be versioned.
- **CONN-REQ-315:** Network-control tests shall verify denied destinations and isolation boundaries.

## 13. Request Construction and Input Validation

- **CONN-REQ-316:** Connector requests shall use explicit schemas.
- **CONN-REQ-317:** Request fields shall be allowlisted.
- **CONN-REQ-318:** Unknown request fields shall be rejected unless explicitly supported.
- **CONN-REQ-319:** Required request fields shall be validated.
- **CONN-REQ-320:** Request types shall be validated.
- **CONN-REQ-321:** Request lengths shall be bounded.
- **CONN-REQ-322:** Request counts shall be bounded.
- **CONN-REQ-323:** Request encodings shall be validated.
- **CONN-REQ-324:** Identifiers shall use canonical formats.
- **CONN-REQ-325:** Paths shall be normalized safely.
- **CONN-REQ-326:** Path traversal shall be prevented.
- **CONN-REQ-327:** Command injection shall be prevented.
- **CONN-REQ-328:** SQL, query-language and template injection shall be prevented.
- **CONN-REQ-329:** Header injection shall be prevented.
- **CONN-REQ-330:** URL injection shall be prevented.
- **CONN-REQ-331:** Serialized-object injection shall be prevented.
- **CONN-REQ-332:** File upload types and sizes shall be constrained.
- **CONN-REQ-333:** Archive extraction shall prevent traversal and resource exhaustion.
- **CONN-REQ-334:** User-supplied and model-supplied input shall be treated as untrusted.
- **CONN-REQ-335:** Connector input shall not include secrets unless explicitly required and protected.
- **CONN-REQ-336:** Input minimization shall occur before transmission.
- **CONN-REQ-337:** Personal data shall be minimized before transmission.
- **CONN-REQ-338:** Case and purpose constraints shall shape request construction.
- **CONN-REQ-339:** Request canonicalization shall be documented where signatures or hashes depend on it.
- **CONN-REQ-340:** Request transformations shall preserve provenance.
- **CONN-REQ-341:** Request validation failure shall produce denial.
- **CONN-REQ-342:** Malformed input shall not be forwarded partially.
- **CONN-REQ-343:** Input-validation errors shall be logged safely.
- **CONN-REQ-344:** Input-validation tests shall include malformed, oversized, encoded and adversarial values.
- **CONN-REQ-345:** Validation shall occur before any side effect.

## 14. Response and Output Validation

- **CONN-REQ-346:** Every connector response shall be treated as untrusted input.
- **CONN-REQ-347:** Response transport integrity shall be verified where applicable.
- **CONN-REQ-348:** Response source identity shall be verified.
- **CONN-REQ-349:** Response status and protocol semantics shall be validated.
- **CONN-REQ-350:** Response media type shall be validated.
- **CONN-REQ-351:** Response schema shall be validated.
- **CONN-REQ-352:** Required response fields shall be validated.
- **CONN-REQ-353:** Response field types shall be validated.
- **CONN-REQ-354:** Response size shall be bounded.
- **CONN-REQ-355:** Response item count shall be bounded.
- **CONN-REQ-356:** Response nesting depth shall be bounded.
- **CONN-REQ-357:** Unexpected fields shall be rejected, ignored safely or quarantined according to policy.
- **CONN-REQ-358:** Malformed responses shall not be treated as successful.
- **CONN-REQ-359:** Partial responses shall be labeled.
- **CONN-REQ-360:** Paginated responses shall preserve ordering and completeness metadata.
- **CONN-REQ-361:** Streaming responses shall enforce per-message validation.
- **CONN-REQ-362:** Signed responses shall have signatures verified when applicable.
- **CONN-REQ-363:** Response-provided hashes shall remain distinct from project-calculated hashes.
- **CONN-REQ-364:** Returned URLs shall not be followed automatically without authorization.
- **CONN-REQ-365:** Returned executable content shall not be executed automatically.
- **CONN-REQ-366:** Returned scripts, macros and active content shall be isolated.
- **CONN-REQ-367:** Returned instructions shall not override policy or system authority.
- **CONN-REQ-368:** Response content shall be scanned or constrained where risk requires it.
- **CONN-REQ-369:** Response normalization shall preserve raw-source linkage.
- **CONN-REQ-370:** Response redaction shall be attributable.
- **CONN-REQ-371:** Output validation failure shall trigger denial, quarantine or qualified handling.
- **CONN-REQ-372:** Validation failure shall not be silently repaired.
- **CONN-REQ-373:** Output-validation rules shall be versioned.
- **CONN-REQ-374:** Output-validation tests shall include malicious, malformed, truncated and oversized responses.
- **CONN-REQ-375:** Validated structure shall not be represented as validated truth.

## 15. Prompt Injection and Content-Borne Instructions

- **CONN-REQ-376:** Connector content shall be treated as potentially adversarial.
- **CONN-REQ-377:** Instructions embedded in connector content shall not override system instructions.
- **CONN-REQ-378:** Instructions embedded in connector content shall not override policy.
- **CONN-REQ-379:** Instructions embedded in connector content shall not override officer authorization.
- **CONN-REQ-380:** Instructions embedded in connector content shall not trigger tool execution directly.
- **CONN-REQ-381:** Instructions embedded in connector content shall not request or receive secrets.
- **CONN-REQ-382:** Instructions embedded in connector content shall not alter connector configuration.
- **CONN-REQ-383:** Instructions embedded in connector content shall not broaden data scope.
- **CONN-REQ-384:** Instructions embedded in connector content shall not change case or purpose.
- **CONN-REQ-385:** Instructions embedded in connector content shall not authorize external communication.
- **CONN-REQ-386:** Indirect prompt injection shall be included in connector threat models.
- **CONN-REQ-387:** Content and control channels shall be separated.
- **CONN-REQ-388:** Untrusted content shall be clearly labeled in model context.
- **CONN-REQ-389:** Connector metadata shall not be treated as trusted instructions automatically.
- **CONN-REQ-390:** Retrieved documents shall not become policy.
- **CONN-REQ-391:** Retrieved code shall not be executed without separate approval and isolation.
- **CONN-REQ-392:** Retrieved links shall not be followed automatically.
- **CONN-REQ-393:** Retrieved credentials or tokens shall trigger containment rather than use.
- **CONN-REQ-394:** Model summaries of connector content shall remain derived analytical output.
- **CONN-REQ-395:** Human review shall be required for consequential decisions based on untrusted connector content.
- **CONN-REQ-396:** Prompt-injection detection shall not be represented as complete protection.
- **CONN-REQ-397:** Safe parsing and constrained context shall be preferred over content-based trust.
- **CONN-REQ-398:** High-risk content shall support quarantine.
- **CONN-REQ-399:** Content-origin provenance shall be preserved.
- **CONN-REQ-400:** Prompt-injection tests shall use synthetic defensive fixtures.
- **CONN-REQ-401:** Tests shall include hidden text, encoded instructions, metadata injection and cross-document attacks.
- **CONN-REQ-402:** Connector failure under prompt injection shall remain fail-closed.
- **CONN-REQ-403:** Model output shall not bypass connector authorization.
- **CONN-REQ-404:** AI-assisted connector use shall preserve human accountability.
- **CONN-REQ-405:** Prompt and policy versions shall be recorded where material.

## 16. Data Minimization and Privacy

- **CONN-REQ-406:** Connector requests shall transmit only data necessary for the approved purpose.
- **CONN-REQ-407:** Connector responses shall retain only data necessary for the approved purpose.
- **CONN-REQ-408:** Personal data fields shall be minimized.
- **CONN-REQ-409:** Sensitive personal data shall receive heightened restrictions.
- **CONN-REQ-410:** Real case data shall not be used in public demonstrations.
- **CONN-REQ-411:** Real officer credentials or identities shall not be used in public demonstrations.
- **CONN-REQ-412:** Public-source accessibility shall not imply unrestricted processing authority.
- **CONN-REQ-413:** Cross-case data transfer shall be denied by default.
- **CONN-REQ-414:** Cross-purpose data transfer shall be denied by default.
- **CONN-REQ-415:** Cross-jurisdiction transfer shall require explicit review where applicable.
- **CONN-REQ-416:** Connector data maps shall identify transmitted and received fields.
- **CONN-REQ-417:** Connector data maps shall identify storage and retention.
- **CONN-REQ-418:** Connector data maps shall identify third-party processing.
- **CONN-REQ-419:** Connector data maps shall identify international or cross-region transfer where applicable.
- **CONN-REQ-420:** Connector telemetry shall minimize personal data.
- **CONN-REQ-421:** Error messages shall not expose sensitive content.
- **CONN-REQ-422:** Support bundles shall not contain secrets or unnecessary personal data.
- **CONN-REQ-423:** Data redaction shall occur before transmission where feasible.
- **CONN-REQ-424:** Pseudonymization shall preserve controlled re-linking only when authorized.
- **CONN-REQ-425:** Anonymization limitations shall be documented.
- **CONN-REQ-426:** Connector caching shall have retention and deletion rules.
- **CONN-REQ-427:** Connector temporary files shall be governed.
- **CONN-REQ-428:** Third-party retention limitations shall be disclosed.
- **CONN-REQ-429:** Third-party model training or reuse shall be disabled or prohibited where required.
- **CONN-REQ-430:** Privacy incidents shall trigger connector suspension and reassessment.
- **CONN-REQ-431:** Rights-impact findings shall have owners and dispositions.
- **CONN-REQ-432:** Connector analytics shall not become excessive surveillance.
- **CONN-REQ-433:** Data deletion requests shall be enforceable where applicable.
- **CONN-REQ-434:** Data minimization shall be tested.
- **CONN-REQ-435:** Document 30 mappings shall not substitute for lawful authorization.

## 17. Provenance and Evidence Preservation

- **CONN-REQ-436:** Every material connector operation shall preserve provenance.
- **CONN-REQ-437:** Provenance shall identify connector identifier.
- **CONN-REQ-438:** Provenance shall identify connector version.
- **CONN-REQ-439:** Provenance shall identify external service or system.
- **CONN-REQ-440:** Provenance shall identify approved endpoint.
- **CONN-REQ-441:** Provenance shall identify request identifier.
- **CONN-REQ-442:** Provenance shall identify response identifier or correlation reference.
- **CONN-REQ-443:** Provenance shall identify officer, agent, workload and institution where applicable.
- **CONN-REQ-444:** Provenance shall identify case and purpose where applicable.
- **CONN-REQ-445:** Provenance shall identify authorization decision.
- **CONN-REQ-446:** Provenance shall identify policy version.
- **CONN-REQ-447:** Provenance shall identify request time and response time.
- **CONN-REQ-448:** Provenance shall identify source-provided timestamps separately.
- **CONN-REQ-449:** Provenance shall identify request scope.
- **CONN-REQ-450:** Provenance shall identify response scope.
- **CONN-REQ-451:** Provenance shall identify filtering and normalization.
- **CONN-REQ-452:** Provenance shall identify retries and pagination.
- **CONN-REQ-453:** Provenance shall identify errors and partial results.
- **CONN-REQ-454:** Provenance shall identify raw-response preservation where lawful and feasible.
- **CONN-REQ-455:** Normalized output shall remain linked to raw output.
- **CONN-REQ-456:** Derived outputs shall identify transformation lineage.
- **CONN-REQ-457:** AI-generated connector summaries shall remain derived analytical output.
- **CONN-REQ-458:** Connector content shall not be represented as original evidence without governed incorporation.
- **CONN-REQ-459:** Evidence acquired through connectors shall conform to `OBDIA-EVID-001`.
- **CONN-REQ-460:** Connector-provided integrity values shall remain distinct from locally verified values.
- **CONN-REQ-461:** Provenance gaps shall remain visible.
- **CONN-REQ-462:** Provenance shall not be fabricated.
- **CONN-REQ-463:** Provenance records shall be append-only or tamper-evident where feasible.
- **CONN-REQ-464:** Provenance tests shall verify request-response and raw-normalized traceability.
- **CONN-REQ-465:** Connector revocation shall not erase provenance.

## 18. Logging and Audit

- **CONN-REQ-466:** Every protected connector operation shall create an audit event.
- **CONN-REQ-467:** Connector audit events shall have immutable identifiers.
- **CONN-REQ-468:** Connector audit events shall identify connector identifier and version.
- **CONN-REQ-469:** Connector audit events shall identify invoking officer, agent and workload where applicable.
- **CONN-REQ-470:** Connector audit events shall identify institution.
- **CONN-REQ-471:** Connector audit events shall identify case and purpose where applicable.
- **CONN-REQ-472:** Connector audit events shall identify action.
- **CONN-REQ-473:** Connector audit events shall identify destination class.
- **CONN-REQ-474:** Connector audit events shall identify authorization-decision reference.
- **CONN-REQ-475:** Connector audit events shall identify policy version.
- **CONN-REQ-476:** Connector audit events shall identify request correlation identifier.
- **CONN-REQ-477:** Connector audit events shall identify result.
- **CONN-REQ-478:** Connector audit events shall identify response-validation outcome.
- **CONN-REQ-479:** Connector audit events shall identify provenance reference.
- **CONN-REQ-480:** Connector audit events shall identify retry, timeout and rate-limit behavior.
- **CONN-REQ-481:** Connector audit events shall identify errors and degraded conditions.
- **CONN-REQ-482:** Connector audit events shall identify operational lifecycle state.
- **CONN-REQ-483:** Connector audit events shall identify credential identifier without secret material.
- **CONN-REQ-484:** Connector audit events shall not contain credentials.
- **CONN-REQ-485:** Connector audit events shall not contain full sensitive payloads by default.
- **CONN-REQ-486:** Connector audit events shall minimize personal data.
- **CONN-REQ-487:** Connector audit events shall be append-only or tamper-evident where feasible.
- **CONN-REQ-488:** Audit-event integrity shall be protected.
- **CONN-REQ-489:** Audit retention shall follow governance and evidence requirements.
- **CONN-REQ-490:** Audit access shall be authorized and logged.
- **CONN-REQ-491:** Audit failure shall block or restrict privileged connector operations.
- **CONN-REQ-492:** Audit gaps shall remain visible.
- **CONN-REQ-493:** Audit corrections shall be additive.
- **CONN-REQ-494:** Audit events shall correlate with agent lifecycle and evidence events.
- **CONN-REQ-495:** Audit tests shall verify success, denial, failure, retry, suspension and revocation.

## 19. Rate, Volume, Quota and Cost Controls

- **CONN-REQ-496:** Every connector shall define request-rate limits.
- **CONN-REQ-497:** Every connector shall define concurrency limits.
- **CONN-REQ-498:** Every connector shall define request-size limits.
- **CONN-REQ-499:** Every connector shall define response-size limits.
- **CONN-REQ-500:** Every connector shall define item-count limits.
- **CONN-REQ-501:** Every connector shall define pagination limits.
- **CONN-REQ-502:** Every connector shall define time-window limits.
- **CONN-REQ-503:** Every connector shall define cost or quota limits where applicable.
- **CONN-REQ-504:** Every connector shall define token or compute limits where applicable.
- **CONN-REQ-505:** Limits shall be scoped by officer, agent, case, purpose or workload where appropriate.
- **CONN-REQ-506:** Limits shall prevent one case from exhausting shared capacity.
- **CONN-REQ-507:** Limit increases shall require authorization.
- **CONN-REQ-508:** Unknown or unavailable quota state shall not permit unbounded use.
- **CONN-REQ-509:** Rate-limit responses shall not trigger uncontrolled retries.
- **CONN-REQ-510:** Quota exhaustion shall produce a controlled failure.
- **CONN-REQ-511:** Cost thresholds shall generate alerts.
- **CONN-REQ-512:** Unexpected cost growth shall trigger suspension or review.
- **CONN-REQ-513:** Volume limits shall protect data minimization.
- **CONN-REQ-514:** Bulk export shall require separate authorization.
- **CONN-REQ-515:** Streaming connectors shall enforce continuous limits.
- **CONN-REQ-516:** Limits shall be versioned.
- **CONN-REQ-517:** Limit changes shall be attributable.
- **CONN-REQ-518:** Limit bypass shall be prohibited.
- **CONN-REQ-519:** Emergency overrides shall be narrow, time-bounded and audited.
- **CONN-REQ-520:** Abuse detection shall distinguish authorized load testing.
- **CONN-REQ-521:** Limits shall not expose sensitive usage to unauthorized users.
- **CONN-REQ-522:** Rate and volume telemetry shall support incident response.
- **CONN-REQ-523:** Limit failure shall isolate the connector.
- **CONN-REQ-524:** Rate-limit tests shall include burst, sustained, concurrent and retry conditions.
- **CONN-REQ-525:** Cost-control limitations shall remain visible.

## 20. Timeout, Retry and Idempotency

- **CONN-REQ-526:** Every connector operation shall have bounded connection and execution timeouts.
- **CONN-REQ-527:** Timeout values shall be risk-based.
- **CONN-REQ-528:** Timeout shall not be interpreted as success.
- **CONN-REQ-529:** Timeout shall not trigger broader fallback access.
- **CONN-REQ-530:** Retries shall be bounded.
- **CONN-REQ-531:** Retry intervals shall use controlled backoff where appropriate.
- **CONN-REQ-532:** Retries shall respect external rate limits.
- **CONN-REQ-533:** Retries shall preserve authorization scope.
- **CONN-REQ-534:** Retries shall not reuse expired authorization.
- **CONN-REQ-535:** Retries shall not reuse revoked credentials.
- **CONN-REQ-536:** Side-effecting operations shall use idempotency controls where supported.
- **CONN-REQ-537:** Idempotency identifiers shall be unique and scoped.
- **CONN-REQ-538:** Duplicate side effects shall be detected or prevented.
- **CONN-REQ-539:** Uncertain side-effect results shall trigger reconciliation rather than blind retry.
- **CONN-REQ-540:** Retry after partial response shall preserve provenance.
- **CONN-REQ-541:** Retry after authentication failure shall not rotate to human credentials.
- **CONN-REQ-542:** Retry after authorization denial shall be prohibited unless context materially changes and a fresh decision is obtained.
- **CONN-REQ-543:** Retry after schema failure shall not bypass validation.
- **CONN-REQ-544:** Retry after integrity failure shall not accept altered output.
- **CONN-REQ-545:** Connector cancellation shall be supported where feasible.
- **CONN-REQ-546:** Cancellation shall not fabricate rollback.
- **CONN-REQ-547:** Long-running operations shall expose status safely.
- **CONN-REQ-548:** Polling shall be bounded.
- **CONN-REQ-549:** Webhook and polling duplication shall be reconciled.
- **CONN-REQ-550:** Idempotency records shall have retention rules.
- **CONN-REQ-551:** Timeout and retry events shall be audited.
- **CONN-REQ-552:** Timeout and retry configuration shall be versioned.
- **CONN-REQ-553:** Retry storms shall be detectable.
- **CONN-REQ-554:** Timeout and idempotency tests shall include uncertain and duplicate outcomes.
- **CONN-REQ-555:** External-system limitations shall be documented.

## 21. Failure Isolation and Circuit Breaking

- **CONN-REQ-556:** Connector failures shall be isolated from unrelated connectors.
- **CONN-REQ-557:** Connector failures shall be isolated from unrelated agents and cases.
- **CONN-REQ-558:** Connector failures shall not corrupt caller state.
- **CONN-REQ-559:** Connector failures shall not corrupt evidence.
- **CONN-REQ-560:** Connector failures shall not broaden network reach.
- **CONN-REQ-561:** Connector failures shall not expose secrets.
- **CONN-REQ-562:** Connector failures shall not bypass authorization.
- **CONN-REQ-563:** Connector failures shall not trigger uncontrolled fallback.
- **CONN-REQ-564:** Circuit breakers shall be used where repeated failure can amplify harm.
- **CONN-REQ-565:** Circuit-breaker state shall be observable.
- **CONN-REQ-566:** Circuit-breaker thresholds shall be versioned.
- **CONN-REQ-567:** Circuit-breaker opening shall stop or restrict operations.
- **CONN-REQ-568:** Circuit-breaker closure shall require recovery evidence.
- **CONN-REQ-569:** Half-open testing shall use constrained operations.
- **CONN-REQ-570:** Bulkhead isolation shall be used where shared resources create blast-radius risk.
- **CONN-REQ-571:** Queue failure shall not drop protected operations silently.
- **CONN-REQ-572:** Backpressure shall be applied where capacity is exceeded.
- **CONN-REQ-573:** Dead-letter or quarantine handling shall preserve provenance.
- **CONN-REQ-574:** Malformed external responses shall not crash unrelated services.
- **CONN-REQ-575:** Dependency failure shall produce controlled degraded state.
- **CONN-REQ-576:** External-service outage shall not produce fabricated data.
- **CONN-REQ-577:** Fallback services shall be pre-approved.
- **CONN-REQ-578:** Fallback services shall be no more permissive.
- **CONN-REQ-579:** Fallback data shall be labeled and provenance-preserved.
- **CONN-REQ-580:** Failure isolation shall include memory, CPU, storage, network and credential boundaries where applicable.
- **CONN-REQ-581:** Critical failure shall support connector suspension.
- **CONN-REQ-582:** Repeated critical failure shall support revocation.
- **CONN-REQ-583:** Failure events shall trigger monitoring and incident handling.
- **CONN-REQ-584:** Failure-isolation tests shall include cascading and resource-exhaustion scenarios.
- **CONN-REQ-585:** Known isolation limitations shall remain visible.

## 22. Revocation and Emergency Disablement

- **CONN-REQ-586:** Every connector shall support administrative suspension.
- **CONN-REQ-587:** Every connector shall support credential revocation.
- **CONN-REQ-588:** Every connector shall support endpoint disablement.
- **CONN-REQ-589:** Every connector shall support session invalidation.
- **CONN-REQ-590:** Every connector shall support workload-access removal.
- **CONN-REQ-591:** Every connector shall support network-egress shutdown.
- **CONN-REQ-592:** Every connector shall support policy-based denial.
- **CONN-REQ-593:** Revocation shall identify authority and reason.
- **CONN-REQ-594:** Revocation shall identify effective time.
- **CONN-REQ-595:** Revocation shall identify affected identities and credentials.
- **CONN-REQ-596:** Revocation shall identify affected sessions and operations.
- **CONN-REQ-597:** Revocation shall propagate to policy and enforcement points.
- **CONN-REQ-598:** Revocation latency shall be defined and tested.
- **CONN-REQ-599:** Critical revocation shall occur immediately or as near immediately as technically feasible.
- **CONN-REQ-600:** Revocation shall prevent new operations.
- **CONN-REQ-601:** Revocation shall stop in-progress operations when safe and required.
- **CONN-REQ-602:** Revocation shall preserve evidence and audit history.
- **CONN-REQ-603:** Revocation shall not erase provenance.
- **CONN-REQ-604:** Revocation shall invalidate cached authorization.
- **CONN-REQ-605:** Revocation failure shall trigger an incident.
- **CONN-REQ-606:** Partial revocation shall remain a critical unresolved condition.
- **CONN-REQ-607:** Revoked connector identifiers shall not be reused.
- **CONN-REQ-608:** Revoked connectors shall not return to enabled state.
- **CONN-REQ-609:** Replacement connectors shall receive new identity or controlled successor identity.
- **CONN-REQ-610:** Emergency disablement shall be attributable.
- **CONN-REQ-611:** Emergency disablement shall not authorize offensive action.
- **CONN-REQ-612:** Emergency disablement shall not delete evidence.
- **CONN-REQ-613:** Re-enablement after suspension shall require review and fresh authorization.
- **CONN-REQ-614:** Revocation tests shall verify credentials, sessions, endpoints, egress and in-progress operations.
- **CONN-REQ-615:** Revocation evidence shall remain retrievable after retirement.

## 23. Supply-Chain and Dependency Security

- **CONN-REQ-616:** Connector source shall be attributable.
- **CONN-REQ-617:** Connector builds shall correspond to reviewed source.
- **CONN-REQ-618:** Connector dependencies shall be inventoried.
- **CONN-REQ-619:** Connector dependencies shall have version and provenance records.
- **CONN-REQ-620:** Connector dependencies shall be assessed for security and maintenance.
- **CONN-REQ-621:** Connector dependencies shall be assessed for license compatibility.
- **CONN-REQ-622:** Unsupported dependencies shall have migration, containment or rejection decisions.
- **CONN-REQ-623:** Dependency updates shall trigger risk-based reassessment.
- **CONN-REQ-624:** Build pipelines shall use least privilege.
- **CONN-REQ-625:** Build credentials shall be protected.
- **CONN-REQ-626:** Build logs shall not expose secrets.
- **CONN-REQ-627:** Build artifacts shall have integrity references.
- **CONN-REQ-628:** Build provenance shall identify source, environment and process.
- **CONN-REQ-629:** Generated code shall receive human review.
- **CONN-REQ-630:** AI-assisted code shall receive human review.
- **CONN-REQ-631:** Vendor SDKs shall not be trusted solely on vendor claims.
- **CONN-REQ-632:** Vendor API specifications shall be verified against observed controlled behavior where feasible.
- **CONN-REQ-633:** Package substitution and dependency confusion shall be addressed.
- **CONN-REQ-634:** Typosquatting and namespace-confusion risks shall be addressed.
- **CONN-REQ-635:** Artifact signatures may supplement but shall not replace review.
- **CONN-REQ-636:** Model, schema and policy dependencies shall be versioned.
- **CONN-REQ-637:** Container or runtime base images shall be governed where applicable.
- **CONN-REQ-638:** Third-party outage and compromise shall be included in threat modeling.
- **CONN-REQ-639:** Supply-chain exceptions shall have owners and expiry.
- **CONN-REQ-640:** Connector SBOM or equivalent dependency record shall be maintained where implementation exists.
- **CONN-REQ-641:** Dependency vulnerabilities shall have dispositions.
- **CONN-REQ-642:** Critical unresolved dependency risk shall block enablement or release.
- **CONN-REQ-643:** Supply-chain tests shall verify artifact integrity and dependency pinning.
- **CONN-REQ-644:** Document 21 shall govern detailed secure coding after consolidation.
- **CONN-REQ-645:** Supply-chain limitations shall remain visible.

## 24. Schema and Protocol Governance

- **CONN-REQ-646:** Connector schemas shall have unique identifiers.
- **CONN-REQ-647:** Connector schemas shall have versions.
- **CONN-REQ-648:** Connector schemas shall identify compatibility rules.
- **CONN-REQ-649:** Connector schemas shall identify required and optional fields.
- **CONN-REQ-650:** Connector schemas shall identify field types and constraints.
- **CONN-REQ-651:** Connector schemas shall identify sensitive fields.
- **CONN-REQ-652:** Connector schemas shall identify deprecated fields.
- **CONN-REQ-653:** Schema changes shall be reviewed.
- **CONN-REQ-654:** Breaking schema changes shall require migration analysis.
- **CONN-REQ-655:** Unknown schema versions shall be rejected or quarantined.
- **CONN-REQ-656:** Schema downgrade shall not bypass controls.
- **CONN-REQ-657:** Schema negotiation shall not select unapproved versions.
- **CONN-REQ-658:** Protocol versions shall be explicit.
- **CONN-REQ-659:** Deprecated protocol versions shall have migration plans.
- **CONN-REQ-660:** Protocol fallback shall be no less secure.
- **CONN-REQ-661:** Serialization formats shall be parsed safely.
- **CONN-REQ-662:** Deserialization of executable objects shall be prohibited unless explicitly isolated and approved.
- **CONN-REQ-663:** Canonicalization rules shall be documented.
- **CONN-REQ-664:** Signature verification shall use canonical representation where required.
- **CONN-REQ-665:** Character encoding shall be validated.
- **CONN-REQ-666:** Compression handling shall be bounded.
- **CONN-REQ-667:** Nested archive handling shall be bounded.
- **CONN-REQ-668:** Schema validation shall occur before business logic.
- **CONN-REQ-669:** Schema errors shall not expose sensitive internals.
- **CONN-REQ-670:** Schema and protocol changes shall trigger tests.
- **CONN-REQ-671:** External schema claims shall not replace local validation.
- **CONN-REQ-672:** Schema provenance shall identify source and adoption authority.
- **CONN-REQ-673:** Schema examples shall use synthetic data.
- **CONN-REQ-674:** Documented compatibility shall not imply security compatibility.
- **CONN-REQ-675:** Schema and protocol governance shall be traceable to connector versions.

## 25. Cloud and Authenticated-Service Connectors

- **CONN-REQ-676:** Cloud connectors shall identify provider, tenant, account, subscription or project context.
- **CONN-REQ-677:** Cloud connectors shall identify region where material.
- **CONN-REQ-678:** Cloud connectors shall identify service and API version.
- **CONN-REQ-679:** Cloud connectors shall use dedicated machine identities.
- **CONN-REQ-680:** Cloud connector roles shall use least privilege.
- **CONN-REQ-681:** Cross-account or cross-tenant access shall be denied by default.
- **CONN-REQ-682:** Cloud resource tags or labels shall not be sole authorization controls.
- **CONN-REQ-683:** Cloud metadata services shall be inaccessible unless explicitly required and protected.
- **CONN-REQ-684:** Cloud administrative APIs shall require stronger controls.
- **CONN-REQ-685:** Cloud data export shall require separate authorization.
- **CONN-REQ-686:** Cloud audit sources shall be identified.
- **CONN-REQ-687:** Cloud connector logs shall correlate with provider logs where feasible.
- **CONN-REQ-688:** Provider-side retention and deletion limitations shall be documented.
- **CONN-REQ-689:** Provider-side encryption claims shall not replace project authorization and key governance.
- **CONN-REQ-690:** Cloud connector regions shall align with privacy and jurisdiction constraints.
- **CONN-REQ-691:** Cloud service changes shall trigger review.
- **CONN-REQ-692:** Cloud provider fallback shall be pre-approved.
- **CONN-REQ-693:** Cloud connector tests shall use sandbox or controlled accounts.
- **CONN-REQ-694:** Authenticated Internet services shall use dedicated non-human credentials.
- **CONN-REQ-695:** Authenticated-service terms and technical constraints shall be documented.
- **CONN-REQ-696:** Account lockout and suspension behavior shall be handled safely.
- **CONN-REQ-697:** Multi-factor requirements for human administrators shall not be bypassed by shared credentials.
- **CONN-REQ-698:** Session cookies or browser tokens shall not be reused as connector credentials unless explicitly approved and protected.
- **CONN-REQ-699:** Cloud and authenticated-service connectors shall avoid scraping when an authorized API exists and policy requires it.
- **CONN-REQ-700:** Cloud connector failure shall not broaden access.
- **CONN-REQ-701:** Cloud evidence acquisition shall conform to `OBDIA-EVID-001`.
- **CONN-REQ-702:** Cloud connector revocation shall remove provider-side access.
- **CONN-REQ-703:** Cloud connector inventory shall remain current.
- **CONN-REQ-704:** Cloud connector test data shall be synthetic or explicitly authorized.
- **CONN-REQ-705:** Provider limitations shall remain visible.

## 26. Web and Public-Source Connectors

- **CONN-REQ-706:** Web connectors shall identify approved domains or source classes.
- **CONN-REQ-707:** Web connectors shall distinguish surface-web access from authenticated deep-web services.
- **CONN-REQ-708:** Deep Web shall not be treated as Dark Web.
- **CONN-REQ-709:** Dark Web connectors shall use separate explicitly approved isolated designs.
- **CONN-REQ-710:** Public availability shall not imply unrestricted collection or reuse.
- **CONN-REQ-711:** Robots, terms, access controls and applicable authorization constraints shall be assessed where relevant.
- **CONN-REQ-712:** Web connectors shall respect case and purpose limits.
- **CONN-REQ-713:** Web connectors shall limit crawling depth and breadth.
- **CONN-REQ-714:** Web connectors shall limit request rates.
- **CONN-REQ-715:** Web connectors shall prevent uncontrolled link following.
- **CONN-REQ-716:** Web connectors shall validate redirects.
- **CONN-REQ-717:** Web connectors shall prevent access to prohibited local and internal destinations.
- **CONN-REQ-718:** Web content shall be treated as untrusted.
- **CONN-REQ-719:** Web active content shall not execute in privileged context.
- **CONN-REQ-720:** Browser automation shall be isolated.
- **CONN-REQ-721:** Downloads shall be type-, size- and integrity-checked.
- **CONN-REQ-722:** Web snapshots shall preserve retrieval time and source URL.
- **CONN-REQ-723:** Dynamic content limitations shall be documented.
- **CONN-REQ-724:** Archived content shall identify archive provider and retrieval reference.
- **CONN-REQ-725:** Authentication-wall bypass shall be prohibited.
- **CONN-REQ-726:** Paywall or access-control bypass shall be prohibited.
- **CONN-REQ-727:** Credential stuffing and credential reuse shall be prohibited.
- **CONN-REQ-728:** Web connectors shall not collect real-person data for public demonstrations.
- **CONN-REQ-729:** Web connectors shall not perform unlawful deanonymization.
- **CONN-REQ-730:** Web connector provenance shall preserve raw and normalized relationships.
- **CONN-REQ-731:** Web extraction errors shall remain visible.
- **CONN-REQ-732:** Web content summaries shall remain derived analytical outputs.
- **CONN-REQ-733:** Web connector tests shall include redirect, SSRF, malicious HTML and prompt-injection cases.
- **CONN-REQ-734:** Dark Web demonstrations shall use simulations or archived-authorized resources.
- **CONN-REQ-735:** Uncontrolled criminal-infrastructure interaction shall remain prohibited.

## 27. Blockchain and Smart-Contract Connectors

- **CONN-REQ-736:** Blockchain connectors shall identify network and chain.
- **CONN-REQ-737:** Blockchain connectors shall identify mainnet, testnet or local network.
- **CONN-REQ-738:** Public demonstrations shall prefer testnets or local networks.
- **CONN-REQ-739:** Blockchain read operations shall be distinguished from transaction submission.
- **CONN-REQ-740:** Transaction submission shall require explicit separate authorization.
- **CONN-REQ-741:** Asset transfer shall require explicit human approval and remain out of public demonstration scope unless testnet-only.
- **CONN-REQ-742:** Private keys shall not be exposed to agent prompts or logs.
- **CONN-REQ-743:** Signing shall occur in a separately controlled component.
- **CONN-REQ-744:** An agent shall not autonomously sign transactions.
- **CONN-REQ-745:** Transaction parameters shall be validated.
- **CONN-REQ-746:** Recipient addresses shall be validated.
- **CONN-REQ-747:** Network identifiers shall be validated.
- **CONN-REQ-748:** Gas, fee or cost limits shall be bounded.
- **CONN-REQ-749:** Nonce and replay behavior shall be controlled.
- **CONN-REQ-750:** Smart-contract calls shall identify contract address and interface version.
- **CONN-REQ-751:** Contract code or verified-source assumptions shall be documented.
- **CONN-REQ-752:** Contract upgradeability shall be considered.
- **CONN-REQ-753:** Chain reorganization and finality limitations shall be documented.
- **CONN-REQ-754:** Block and transaction timestamps shall remain distinct from local observation time.
- **CONN-REQ-755:** Node or provider provenance shall be recorded.
- **CONN-REQ-756:** Provider responses shall be treated as untrusted.
- **CONN-REQ-757:** Multiple providers may be compared for validation where appropriate.
- **CONN-REQ-758:** Provider disagreement shall remain visible.
- **CONN-REQ-759:** Blockchain analytics shall not imply identity attribution without evidence.
- **CONN-REQ-760:** Wallet-address analysis shall not be represented as lawful deanonymization.
- **CONN-REQ-761:** Testnet data shall be labeled.
- **CONN-REQ-762:** Smart-contract simulation shall be preferred before any authorized state change.
- **CONN-REQ-763:** Blockchain connector revocation shall disable signing and provider access.
- **CONN-REQ-764:** Blockchain connector tests shall include wrong-chain, replay, reorg, malformed response and high-fee cases.
- **CONN-REQ-765:** Mainnet interaction in public demonstrations shall be read-only unless a separate approved safe design exists.

## 28. Communication-Platform Connectors

- **CONN-REQ-766:** Communication-platform connectors shall identify platform and account context.
- **CONN-REQ-767:** Connector accounts shall be institutionally controlled where applicable.
- **CONN-REQ-768:** Personal accounts shall not be reused as connector identities.
- **CONN-REQ-769:** Message read, search, export, send, edit and delete operations shall remain distinct.
- **CONN-REQ-770:** Sending messages shall require explicit separate authorization.
- **CONN-REQ-771:** An agent shall not impersonate the bound officer.
- **CONN-REQ-772:** Automated outbound communication shall identify its authorized nature where policy requires it.
- **CONN-REQ-773:** Recipient scope shall be validated.
- **CONN-REQ-774:** Attachment handling shall follow input and output controls.
- **CONN-REQ-775:** Message content shall be treated as untrusted.
- **CONN-REQ-776:** Message instructions shall not override policy.
- **CONN-REQ-777:** Communication metadata shall be minimized.
- **CONN-REQ-778:** Retention and deletion limitations shall be documented.
- **CONN-REQ-779:** Platform audit capabilities shall be identified.
- **CONN-REQ-780:** Platform rate limits shall be respected.
- **CONN-REQ-781:** Account suspension behavior shall be handled safely.
- **CONN-REQ-782:** Cross-workspace or cross-tenant access shall be denied by default.
- **CONN-REQ-783:** Public demonstrations shall use mock or test workspaces.
- **CONN-REQ-784:** Real-person messaging shall not be used in public demonstrations.
- **CONN-REQ-785:** Communication evidence acquisition shall preserve provenance and authorization.
- **CONN-REQ-786:** Connector exports shall preserve message identifiers and time-source context.
- **CONN-REQ-787:** Outbound communication shall not be autonomous coercive action.
- **CONN-REQ-788:** Outbound communication shall not make legal determinations.
- **CONN-REQ-789:** Emergency disablement shall stop new messages.
- **CONN-REQ-790:** Communication connector tests shall include wrong recipient, duplicate send, malicious attachment and prompt injection.

## 29. Threat-Intelligence and Evidence-System Connectors

- **CONN-REQ-791:** Threat-intelligence connectors shall identify source, feed and license constraints.
- **CONN-REQ-792:** Threat-intelligence confidence shall remain distinct from evidence integrity.
- **CONN-REQ-793:** Threat-intelligence indicators shall preserve source and timestamp.
- **CONN-REQ-794:** Indicator ingestion shall not trigger blocking or coercive action autonomously.
- **CONN-REQ-795:** Indicator enrichment shall preserve provenance.
- **CONN-REQ-796:** Conflicting intelligence shall remain attributable.
- **CONN-REQ-797:** Vendor scoring shall not replace project-specific assessment.
- **CONN-REQ-798:** Threat-intelligence data shall not be treated as verified fact automatically.
- **CONN-REQ-799:** Evidence-system connectors shall preserve evidence identifiers.
- **CONN-REQ-800:** Evidence-system connectors shall preserve custody and provenance references.
- **CONN-REQ-801:** Evidence-system write operations shall require stronger authorization.
- **CONN-REQ-802:** Evidence-system deletion shall require exceptional explicit authority.
- **CONN-REQ-803:** Evidence-system exports shall include manifests and integrity values.
- **CONN-REQ-804:** Evidence-system schema mappings shall preserve semantic distinctions.
- **CONN-REQ-805:** Connector normalization shall not overwrite source evidence.
- **CONN-REQ-806:** Evidence-system connectors shall support read-only modes.
- **CONN-REQ-807:** Evidence-system connector failures shall not corrupt authoritative evidence.
- **CONN-REQ-808:** Evidence-system access shall be case- and purpose-scoped.
- **CONN-REQ-809:** Evidence-system connectors shall log all protected operations.
- **CONN-REQ-810:** Threat-intelligence and evidence connectors shall support revocation.
- **CONN-REQ-811:** Imported data shall be quarantined when integrity or provenance is materially uncertain.
- **CONN-REQ-812:** Indicator expiry and deprecation shall be preserved.
- **CONN-REQ-813:** Feed outages shall not fabricate absence of threat.
- **CONN-REQ-814:** Connector tests shall include stale indicators, conflicting feeds, custody mismatch and export failure.
- **CONN-REQ-815:** Public demonstrations shall use synthetic indicators and evidence.

## 30. Monitoring and Detection

- **CONN-REQ-816:** Connector availability shall be monitored.
- **CONN-REQ-817:** Connector operational state shall be monitored.
- **CONN-REQ-818:** Connector authentication failures shall be monitored.
- **CONN-REQ-819:** Connector authorization denials shall be monitored.
- **CONN-REQ-820:** Connector credential expiry shall be monitored.
- **CONN-REQ-821:** Connector revocation state shall be monitored.
- **CONN-REQ-822:** Unexpected endpoints shall be monitored.
- **CONN-REQ-823:** Unexpected protocols shall be monitored.
- **CONN-REQ-824:** Unusual request rate or volume shall be monitored.
- **CONN-REQ-825:** Unusual response size or schema shall be monitored.
- **CONN-REQ-826:** Repeated retries and timeouts shall be monitored.
- **CONN-REQ-827:** Circuit-breaker state shall be monitored.
- **CONN-REQ-828:** Prompt-injection and malicious-content indicators shall be monitored where feasible.
- **CONN-REQ-829:** Output-validation failures shall be monitored.
- **CONN-REQ-830:** Provenance gaps shall be monitored.
- **CONN-REQ-831:** Audit gaps shall be monitored.
- **CONN-REQ-832:** Cross-case and cross-purpose attempts shall be monitored.
- **CONN-REQ-833:** Unexpected write or delete attempts shall be monitored.
- **CONN-REQ-834:** Network egress anomalies shall be monitored.
- **CONN-REQ-835:** Dependency and version drift shall be monitored.
- **CONN-REQ-836:** Monitoring alerts shall have severity and owner.
- **CONN-REQ-837:** Alert suppression shall be authorized and time-bounded.
- **CONN-REQ-838:** Monitoring failure shall create a degraded state.
- **CONN-REQ-839:** Critical monitoring failure shall trigger suspension where required.
- **CONN-REQ-840:** Monitoring data shall be protected from alteration.
- **CONN-REQ-841:** Monitoring access shall be authorized.
- **CONN-REQ-842:** Monitoring shall minimize personal and case data.
- **CONN-REQ-843:** Monitoring shall distinguish approved security tests.
- **CONN-REQ-844:** Monitoring effectiveness shall be reviewed after incidents.
- **CONN-REQ-845:** Monitoring tests shall verify alerts for critical connector conditions.

## 31. Incident Response and Containment

- **CONN-REQ-846:** Connector incidents shall have immutable identifiers.
- **CONN-REQ-847:** Connector incidents shall identify affected connector and version.
- **CONN-REQ-848:** Connector incidents shall identify affected identities and credentials.
- **CONN-REQ-849:** Connector incidents shall identify affected endpoints and operations.
- **CONN-REQ-850:** Connector incidents shall identify affected officers, agents, workloads, cases and purposes where applicable.
- **CONN-REQ-851:** Connector incidents shall preserve evidence.
- **CONN-REQ-852:** Connector incidents shall preserve audit and provenance.
- **CONN-REQ-853:** Credential exposure shall trigger rotation or revocation.
- **CONN-REQ-854:** Unauthorized endpoint access shall trigger suspension.
- **CONN-REQ-855:** Unauthorized data transfer shall trigger containment and privacy review.
- **CONN-REQ-856:** Malicious response processing shall trigger quarantine.
- **CONN-REQ-857:** Prompt-injection compromise shall trigger containment.
- **CONN-REQ-858:** Supply-chain compromise shall trigger revocation or suspension.
- **CONN-REQ-859:** Evidence corruption shall trigger quarantine and evidence review.
- **CONN-REQ-860:** Audit failure shall trigger restricted operation.
- **CONN-REQ-861:** Incident containment shall minimize unrelated disruption.
- **CONN-REQ-862:** Incident containment shall not authorize offensive activity.
- **CONN-REQ-863:** Incident containment shall not erase evidence.
- **CONN-REQ-864:** Recovery shall restore a known governed connector state.
- **CONN-REQ-865:** Recovery shall verify credentials, policy, endpoints, schemas and monitoring.
- **CONN-REQ-866:** Recovery shall not restore revoked authority.
- **CONN-REQ-867:** Post-incident review shall update threats, risks and controls.
- **CONN-REQ-868:** Repeated incidents shall trigger architecture review.
- **CONN-REQ-869:** Incident communications shall protect sensitive information.
- **CONN-REQ-870:** Incident closure shall require evidence of containment and follow-up ownership.

## 32. Testing and Validation

- **CONN-REQ-871:** Every material connector requirement shall have testable acceptance criteria.
- **CONN-REQ-872:** Connector tests shall use mock, sandbox, testnet or otherwise authorized controlled systems.
- **CONN-REQ-873:** Tests shall verify secure authentication.
- **CONN-REQ-874:** Tests shall verify expired credential denial.
- **CONN-REQ-875:** Tests shall verify revoked credential denial.
- **CONN-REQ-876:** Tests shall verify least privilege.
- **CONN-REQ-877:** Tests shall verify out-of-scope action denial.
- **CONN-REQ-878:** Tests shall verify cross-case denial.
- **CONN-REQ-879:** Tests shall verify cross-purpose denial.
- **CONN-REQ-880:** Tests shall verify endpoint allowlisting.
- **CONN-REQ-881:** Tests shall verify redirect validation.
- **CONN-REQ-882:** Tests shall verify SSRF protections.
- **CONN-REQ-883:** Tests shall verify request-schema validation.
- **CONN-REQ-884:** Tests shall verify output-schema validation.
- **CONN-REQ-885:** Tests shall verify malformed response handling.
- **CONN-REQ-886:** Tests shall verify oversized response handling.
- **CONN-REQ-887:** Tests shall verify active-content isolation.
- **CONN-REQ-888:** Tests shall verify prompt-injection resistance.
- **CONN-REQ-889:** Tests shall verify provenance capture.
- **CONN-REQ-890:** Tests shall verify raw-normalized lineage.
- **CONN-REQ-891:** Tests shall verify audit-event creation.
- **CONN-REQ-892:** Tests shall verify rate and volume limits.
- **CONN-REQ-893:** Tests shall verify timeout behavior.
- **CONN-REQ-894:** Tests shall verify bounded retries.
- **CONN-REQ-895:** Tests shall verify idempotency.
- **CONN-REQ-896:** Tests shall verify uncertain side-effect reconciliation.
- **CONN-REQ-897:** Tests shall verify circuit breaking.
- **CONN-REQ-898:** Tests shall verify failure isolation.
- **CONN-REQ-899:** Tests shall verify network egress control.
- **CONN-REQ-900:** Tests shall verify suspension propagation.
- **CONN-REQ-901:** Tests shall verify revocation propagation.
- **CONN-REQ-902:** Tests shall verify connector lifecycle transitions.
- **CONN-REQ-903:** Tests shall verify supply-chain artifact integrity.
- **CONN-REQ-904:** Tests shall verify monitoring alerts.
- **CONN-REQ-905:** Tests shall verify audit-service failure behavior.
- **CONN-REQ-906:** Tests shall verify privacy minimization.
- **CONN-REQ-907:** Tests shall verify evidence handling where applicable.
- **CONN-REQ-908:** Tests shall identify exact connector, code, dependency, schema, policy, configuration and environment versions.
- **CONN-REQ-909:** Failed tests shall remain visible.
- **CONN-REQ-910:** Passing tests shall not be generalized beyond scope.
- **CONN-REQ-911:** Automated tests shall not replace human connector review.
- **CONN-REQ-912:** Document 22 remains a forward dependency for complete testing governance.
- **CONN-REQ-913:** Material connector changes shall invalidate affected validation evidence.
- **CONN-REQ-914:** Validation limitations shall remain visible.
- **CONN-REQ-915:** Negative and abuse-case testing shall be mandatory for side-effecting connectors.

## 33. Traceability

- **CONN-REQ-916:** Every connector requirement shall trace to governing authority or threat.
- **CONN-REQ-917:** Every connector manifest field shall trace to a requirement or documented rationale.
- **CONN-REQ-918:** Every connector identity shall trace to issuance and owner.
- **CONN-REQ-919:** Every credential shall trace to identity and scope.
- **CONN-REQ-920:** Every permission shall trace to an authorized operation.
- **CONN-REQ-921:** Every endpoint shall trace to approved purpose.
- **CONN-REQ-922:** Every connector operation shall trace to authorization.
- **CONN-REQ-923:** Every response shall trace to request and source.
- **CONN-REQ-924:** Every normalized output shall trace to raw output.
- **CONN-REQ-925:** Every evidence object shall trace to connector provenance where applicable.
- **CONN-REQ-926:** Every audit event shall trace to connector operation.
- **CONN-REQ-927:** Every suspension shall trace to reason and authority.
- **CONN-REQ-928:** Every revocation shall trace to affected credentials and sessions.
- **CONN-REQ-929:** Every incident shall trace to affected connector requirements and controls.
- **CONN-REQ-930:** Every risk and exception shall trace to affected connector requirements.
- **CONN-REQ-931:** Every test shall trace to requirements and threats.
- **CONN-REQ-932:** Every release shall trace to connector validation where connectors are included.
- **CONN-REQ-933:** Traceability shall be bidirectional.
- **CONN-REQ-934:** Broken traceability affecting identity, authorization, data, evidence, revocation or release shall be blocking.
- **CONN-REQ-935:** Planned connector controls shall not be represented as implemented.
- **CONN-REQ-936:** Implemented connector controls shall not be represented as validated without evidence.
- **CONN-REQ-937:** Superseded connector versions shall identify successors.
- **CONN-REQ-938:** Deprecated schemas and endpoints shall identify replacements.
- **CONN-REQ-939:** Traceability shall use immutable identifiers.
- **CONN-REQ-940:** Traceability records shall not contain secrets.
- **CONN-REQ-941:** Traceability records shall minimize personal and case data.
- **CONN-REQ-942:** Baseline freeze shall validate connector traceability across documents 01–30.
- **CONN-REQ-943:** No document above 30 shall acquire normative connector authority by implication.

## 34. Connector Review Gate

- **CONN-REQ-944:** Connector-policy review shall identify the exact document and commit.
- **CONN-REQ-945:** Review shall verify preservation of all eight legacy controls.
- **CONN-REQ-946:** Review shall verify manifest completeness.
- **CONN-REQ-947:** Review shall verify connector identity and credential security.
- **CONN-REQ-948:** Review shall verify least privilege.
- **CONN-REQ-949:** Review shall verify authorization enforcement.
- **CONN-REQ-950:** Review shall verify endpoint and egress controls.
- **CONN-REQ-951:** Review shall verify input and output validation.
- **CONN-REQ-952:** Review shall verify prompt-injection controls.
- **CONN-REQ-953:** Review shall verify data minimization and privacy.
- **CONN-REQ-954:** Review shall verify provenance and evidence handling.
- **CONN-REQ-955:** Review shall verify audit logging.
- **CONN-REQ-956:** Review shall verify limits, retries and idempotency.
- **CONN-REQ-957:** Review shall verify failure isolation.
- **CONN-REQ-958:** Review shall verify revocation and emergency disablement.
- **CONN-REQ-959:** Review shall verify supply-chain controls.
- **CONN-REQ-960:** Review shall verify connector-class-specific controls.
- **CONN-REQ-961:** Review shall verify monitoring, incident response and testing.
- **CONN-REQ-962:** Review shall identify residual risks and forward dependencies.
- **CONN-REQ-963:** Critical findings shall block progression.
- **CONN-REQ-964:** Material post-review changes shall invalidate affected review evidence.
- **CONN-REQ-965:** Role concentration shall be disclosed.
- **CONN-REQ-966:** Internal review shall not be represented as independent external assurance.
- **CONN-REQ-967:** Security Reviewer shall assess identity, authorization, failure and revocation.
- **CONN-REQ-968:** Privacy and Governance Reviewer shall assess data minimization, transfer, retention and rights impact.
- **CONN-REQ-969:** Implementation Reviewer shall assess implementation alignment where implementation exists.
- **CONN-REQ-970:** Research Reviewer shall assess source and external-system claims.
- **CONN-REQ-971:** Release Reviewer shall assess distributed connectors and public demonstrations.
- **CONN-REQ-972:** Project Founder approval shall not substitute for required specialist review.
- **CONN-REQ-973:** Approval for Draft incorporation shall not approve any operational connector.

## 35. Minimum Validation Checklist

Before approval of this policy or a connector, confirm:

- [ ] unique connector identifier and manifest exist;
- [ ] connector owner and approval authority are identified;
- [ ] connector class and operational state are correct;
- [ ] distinct machine identity is used;
- [ ] no reusable human credential is present;
- [ ] authentication is secure;
- [ ] permissions are least-privileged;
- [ ] authorization defaults to deny;
- [ ] case, purpose, jurisdiction, time, action and data scopes are enforced;
- [ ] approved endpoints and protocols are explicit;
- [ ] network egress is constrained;
- [ ] request schemas and bounds are enforced;
- [ ] responses are treated as untrusted and validated;
- [ ] indirect prompt injection cannot override authority;
- [ ] data minimization and privacy controls are applied;
- [ ] provenance and raw-to-normalized lineage are preserved;
- [ ] audit events are complete and secret-free;
- [ ] rate, volume, timeout, retry and idempotency controls are defined;
- [ ] failures are isolated;
- [ ] suspension and revocation propagate;
- [ ] supply-chain provenance is recorded;
- [ ] connector-class-specific restrictions are satisfied;
- [ ] monitoring and incident response are operationally testable;
- [ ] negative and abuse-case tests pass;
- [ ] forward dependencies 21–30 are recorded where applicable;
- [ ] Project Founder approval exists before status becomes Approved.


## 36. Limitations

- This policy does not approve any connector, endpoint, protocol, vendor, cloud service, threat-intelligence source, evidence system, communication platform, blockchain network or smart contract.
- It does not provide credentials or authorize access.
- It does not establish legal authority, institutional deployment, production readiness, certification or compliance.
- It does not guarantee that external systems return correct, complete or lawful data.
- Secure authentication does not establish authorization.
- Schema validation does not establish truth.
- Provenance does not establish authenticity or admissibility.
- Failure isolation reduces but does not eliminate shared-infrastructure risk.
- Revocation may have non-zero propagation latency that must be bounded and tested.
- Internal connector review is not independent certification.
- Documents 21–30 remain forward dependencies where they govern coding, testing, structure, versioning, risk, exceptions, changes and compliance.


## 37. Change Control

Every material change shall identify rationale, affected connectors, manifests, identities, credentials, endpoints, schemas, permissions, data flows, dependencies, security and privacy impact, evidence impact, migration, validation, rollback, authority and version effect.

- **CONN-REQ-974:** Editorial corrections shall use a patch version when meaning is unchanged.
- **CONN-REQ-975:** Backward-compatible substantive additions shall use a minor version.
- **CONN-REQ-976:** Incompatible connector-governance changes shall use a major version.
- **CONN-REQ-977:** Material connector-architecture decisions shall require an ADR where applicable.
- **CONN-REQ-978:** Changes shall require Project Founder approval.
- **CONN-REQ-979:** Changes shall receive Security Reviewer assessment.
- **CONN-REQ-980:** Privacy, authorization, evidence, implementation and release impacts shall receive specialist review where applicable.
- **CONN-REQ-981:** Changes shall identify affected manifests, credentials, policies, schemas, tests and integrations.
- **CONN-REQ-982:** Changes shall include migration and compatibility analysis.
- **CONN-REQ-983:** Changes shall include validation and rollback analysis.
- **CONN-REQ-984:** Changes shall not retroactively fabricate connector approval, provenance or audit evidence.
- **CONN-REQ-985:** Historical connector events shall not be silently rewritten.
- **CONN-REQ-986:** Connector identifiers shall not be reused.
- **CONN-REQ-987:** Endpoint, schema and credential migration shall preserve traceability.
- **CONN-REQ-988:** Forward-dependency reconciliation shall occur before this policy becomes Approved.

## 38. Consolidation Record

Version 1.0.0 consolidates the two existing Connector Security Policy variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-CONN-001`;
- normalizes the authoritative filename to `20_CONNECTOR_SECURITY_POLICY.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves secure authentication;
- preserves least privilege and permission minimization;
- preserves output validation;
- preserves authorization-boundary enforcement;
- preserves failure isolation;
- preserves usage logging and audit logs;
- preserves provenance;
- preserves revocation support;
- adds connector definitions, manifests, classifications, operational lifecycle, identities, credentials, endpoint and egress controls, request and response validation, prompt-injection defenses, data minimization, evidence handling, limits, retries, idempotency, circuit breaking, supply-chain, schema, environment-specific controls, monitoring, incidents, testing and traceability;
- identifies documents 21–30 as forward dependencies where applicable;
- treats Enterprise and non-Enterprise legacy files as source variants of the same immutable document;
- creates no connector, credential, endpoint authorization, operational integration or production claim.

## 39. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Security Reviewer; approval reserved to Project Founder | Constitutional consolidation of both legacy connector policies: preserved all eight original controls; added complete manifest, lifecycle, identity, authorization, validation, isolation, revocation, environment-specific, testing and traceability requirements. |
| 1.0.1 | 2026-08-06 | Draft | Security Reviewer; approval reserved to Project Founder | Reconciled obsolete forward-dependency filenames with the final canonical Knowledge Pack 01–30 filenames; no requirement text, authority, lifecycle status, implementation state or validation claim changed. |
