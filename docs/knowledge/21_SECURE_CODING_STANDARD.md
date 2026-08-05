# SECURE CODING STANDARD

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-CODE-001 |
| **Title** | Secure Coding Standard |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Implementation Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define mandatory secure-coding, readability, maintainability, determinism, complexity, interface, validation, error, secret, dependency, testing, review, supply-chain and implementation-evidence requirements for OBDIA software and code-like artifacts. |
| **Scope** | Application code, libraries, command-line tools, services, connector implementations, policy code, prompts and tool schemas, data-processing code, infrastructure-as-code, build scripts, shell scripts, CI/CD definitions, configuration parsers, migrations, tests, generated code and AI-assisted changes within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `02_PROJECT_GLOSSARY.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `07_THREAT_MODEL_BASELINE.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `15_RELEASE_AND_PUBLICATION_POLICY.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; `18_EVIDENCE_MODEL.md`; `19_AGENT_LIFECYCLE_MODEL.md`; `20_CONNECTOR_SECURITY_POLICY.md`; forward dependencies `22_TESTING_STANDARD.md`, `24_GITHUB_REPOSITORY_STANDARD.md`, `25_DOCUMENT_VERSIONING_POLICY.md`, `26_AI_RISK_REGISTER.md`, `27_RISK_ASSESSMENT_METHOD.md`, `28_EXCEPTION_MANAGEMENT_POLICY.md`, `29_CHANGE_CONTROL_POLICY.md` and `30_COMPLIANCE_MAPPING_POLICY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GLOSS-001`; `OBDIA-RES-001`; `OBDIA-ARCH-001`; `OBDIA-TM-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-REL-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001`; `OBDIA-EVID-001`; `OBDIA-AGENT-001`; `OBDIA-CONN-001` |
| **Cross-References** | Documents 22–30; accepted ADRs; threat models; implementation plans; code-review records; dependency manifests; build and test evidence; risk, exception, incident, change and release records |
| **Assumptions** | Implementations use synthetic data, mock services, testnets, local laboratories, isolated research environments and other explicitly authorized resources until separate governance establishes a broader environment. |
| **Constraints** | Code shall not create independent AI legal authority, authorize prohibited conduct, embed human credentials, conceal unsafe behavior, bypass authorization or evidence controls, interact with uncontrolled criminal infrastructure, or overstate implementation and validation status. |
| **Security Considerations** | Coding defects can cause privilege escalation, injection, secret exposure, authorization bypass, evidence corruption, privacy harm, supply-chain compromise, cross-case contamination, unsafe tool execution, nondeterministic behavior, denial of service and misleading audit records. |
| **Validation Criteria** | Material code has documented interfaces, validated inputs and outputs, explicit errors, secure defaults, bounded resources, protected secrets, controlled dependencies, review evidence, security-focused tests, reproducible builds where applicable and bidirectional traceability to requirements, threats and validation evidence. |
| **Implementation Relationship** | This standard governs implementation quality. It does not choose a programming language, framework, cloud provider, policy engine, model provider, connector vendor or production deployment architecture. |

---

## 1. Purpose

This document consolidates the OBDIA Coding Standard into an enforceable Secure Coding Standard.

The original baseline established five principles:

- Secure by Design;
- Readability;
- Maintainability;
- Deterministic behavior;
- Least complexity.

It also required:

- documentation of every public interface;
- input validation;
- explicit error handling;
- avoidance of hard-coded secrets;
- dependency pinning where appropriate;
- security-focused tests.

This consolidation preserves every original principle and requirement and expands them into complete implementation controls suitable for security-sensitive AI-assisted digital-investigation research.

The normalized title `Secure Coding Standard` does not create a new document. It retains immutable identifier `OBDIA-CODE-001` and aligns the canonical filename with existing cross-references.


## 2. Fundamental Coding Rules

- **CODE-REQ-001:** Code shall remain subordinate to approved architecture, ADRs, authorization, evidence, lifecycle and connector policies.
- **CODE-REQ-002:** Code shall not silently define architecture.
- **CODE-REQ-003:** Code shall not create independent legal authority for an AI agent.
- **CODE-REQ-004:** Code shall not authorize canonical prohibitions.
- **CODE-REQ-005:** Code shall not transfer human credentials to an agent or workload.
- **CODE-REQ-006:** Code shall use secure defaults.
- **CODE-REQ-007:** Code shall fail closed for privileged, evidentiary and side-effecting operations.
- **CODE-REQ-008:** Code shall use least privilege.
- **CODE-REQ-009:** Code shall minimize data collection, processing, retention and disclosure.
- **CODE-REQ-010:** Code shall preserve human and institutional accountability.
- **CODE-REQ-011:** Code shall preserve officer-agent binding where applicable.
- **CODE-REQ-012:** Code shall preserve case and purpose separation.
- **CODE-REQ-013:** Code shall preserve evidence provenance and chain of custody where applicable.
- **CODE-REQ-014:** Code shall distinguish source evidence from derived and AI-generated output.
- **CODE-REQ-015:** Code shall be readable by an appropriately skilled reviewer.
- **CODE-REQ-016:** Code shall be maintainable without relying on undocumented author knowledge.
- **CODE-REQ-017:** Code shall be deterministic where the problem and platform permit it.
- **CODE-REQ-018:** Necessary nondeterminism shall be explicit, bounded and testable.
- **CODE-REQ-019:** Code shall use the least complexity sufficient for the approved requirement.
- **CODE-REQ-020:** Code shall avoid hidden control paths.
- **CODE-REQ-021:** Code shall avoid surprising privilege or side effects.
- **CODE-REQ-022:** Code shall be attributable through version control.
- **CODE-REQ-023:** Material code shall trace to approved requirements and ADRs.
- **CODE-REQ-024:** Generated and AI-assisted code shall receive human review.
- **CODE-REQ-025:** Passing automated checks shall not replace human code review.
- **CODE-REQ-026:** Repository merge shall not establish implementation validation.
- **CODE-REQ-027:** Unknown security state shall resolve to denial, containment or explicit failure.
- **CODE-REQ-028:** Critical unresolved coding defects shall block acceptance or release.
- **CODE-REQ-029:** Code examples in public documentation shall use synthetic and non-secret values.
- **CODE-REQ-030:** Code shall not interact with uncontrolled criminal infrastructure.

## 3. Core Principles

- **CODE-REQ-031:** `Secure by Design` shall be applied from requirements through implementation and validation.
- **CODE-REQ-032:** `Readability` shall be treated as a security and maintainability property.
- **CODE-REQ-033:** `Maintainability` shall include clear ownership, modularity, tests and change safety.
- **CODE-REQ-034:** `Deterministic behavior` shall be preferred for authorization, evidence and policy-critical logic.
- **CODE-REQ-035:** `Least complexity` shall be preferred over unnecessary abstraction, indirection and dependency.
- **CODE-REQ-036:** Security controls shall be explicit rather than implied by naming.
- **CODE-REQ-037:** Privacy by Design shall apply to data structures, interfaces, logs and retention.
- **CODE-REQ-038:** Defense in Depth shall prevent reliance on a single code-level control.
- **CODE-REQ-039:** Zero Trust shall apply to inputs, outputs, dependencies, services and runtime context.
- **CODE-REQ-040:** Least privilege shall apply to code execution, files, network, credentials and tools.
- **CODE-REQ-041:** Fail-safe defaults shall deny or contain when policy or context is unavailable.
- **CODE-REQ-042:** Complete mediation shall apply to protected operations.
- **CODE-REQ-043:** Separation of duties shall be preserved where code implements approval or governance paths.
- **CODE-REQ-044:** Economy of mechanism shall favor small, understandable security mechanisms.
- **CODE-REQ-045:** Open design shall avoid relying on secrecy of implementation for protection.
- **CODE-REQ-046:** Traceability shall connect code to requirements, threats, tests and evidence.
- **CODE-REQ-047:** Reproducibility shall be preserved where it materially supports verification.
- **CODE-REQ-048:** Observability shall support accountability without excessive data collection.
- **CODE-REQ-049:** Resilience shall include bounded failure, containment, recovery and rollback.
- **CODE-REQ-050:** Explicit limitations shall be documented rather than concealed.

## 4. Applicability and Code Classification

- **CODE-REQ-051:** Every material code artifact shall have an identified owner.
- **CODE-REQ-052:** Every material code artifact shall have an implementation status.
- **CODE-REQ-053:** Every material code artifact shall identify its security criticality.
- **CODE-REQ-054:** Every material code artifact shall identify applicable data classification.
- **CODE-REQ-055:** Every material code artifact shall identify applicable trust boundaries.
- **CODE-REQ-056:** Every material code artifact shall identify applicable threat-model references.
- **CODE-REQ-057:** Every material code artifact shall identify applicable authorization requirements.
- **CODE-REQ-058:** Every material code artifact shall identify applicable evidence requirements.
- **CODE-REQ-059:** Every material code artifact shall identify applicable connector requirements.
- **CODE-REQ-060:** Every material code artifact shall identify supported execution environments.
- **CODE-REQ-061:** Security-critical code shall receive enhanced review and testing.
- **CODE-REQ-062:** Authorization-critical code shall be classified as security-critical.
- **CODE-REQ-063:** Identity and credential code shall be classified as security-critical.
- **CODE-REQ-064:** Evidence integrity and custody code shall be classified as security-critical.
- **CODE-REQ-065:** Cryptographic code shall be classified as security-critical.
- **CODE-REQ-066:** Connector request, response and credential code shall be classified as security-critical.
- **CODE-REQ-067:** Policy parsing and enforcement code shall be classified as security-critical.
- **CODE-REQ-068:** Release and update code shall be classified as security-critical where it can distribute executable artifacts.
- **CODE-REQ-069:** Test-only code shall not be assumed harmless when it accesses real resources or credentials.
- **CODE-REQ-070:** Demonstration code shall remain governed when it exercises security boundaries.
- **CODE-REQ-071:** Generated code shall inherit the criticality of its use.
- **CODE-REQ-072:** Unknown criticality shall be treated conservatively.
- **CODE-REQ-073:** Non-applicability of a material section shall have documented rationale.
- **CODE-REQ-074:** Code classification changes shall trigger review.
- **CODE-REQ-075:** Classification shall not be used to bypass mandatory controls.

## 5. Repository and Module Structure

- **CODE-REQ-076:** Code shall reside in governed repository locations.
- **CODE-REQ-077:** Authoritative source shall be distinguishable from generated output.
- **CODE-REQ-078:** Production-like source shall be distinguishable from experiments and demonstrations.
- **CODE-REQ-079:** Tests shall be distinguishable from application code.
- **CODE-REQ-080:** Fixtures shall be distinguishable from real data.
- **CODE-REQ-081:** Configuration shall be distinguishable from secret material.
- **CODE-REQ-082:** Migrations shall be distinguishable from ordinary runtime code.
- **CODE-REQ-083:** Security-critical modules shall have clear ownership.
- **CODE-REQ-084:** Module boundaries shall reflect architectural boundaries where practical.
- **CODE-REQ-085:** Module boundaries shall avoid cyclic dependencies.
- **CODE-REQ-086:** Cross-layer access shall be explicit and justified.
- **CODE-REQ-087:** Internal modules shall not become public interfaces accidentally.
- **CODE-REQ-088:** Package names shall conform to repository naming rules.
- **CODE-REQ-089:** File names shall conform to language and repository conventions.
- **CODE-REQ-090:** Directory traversal through dynamic module loading shall be prohibited.
- **CODE-REQ-091:** Generated files shall identify their generator where practical.
- **CODE-REQ-092:** Generated files shall not be edited manually unless the process explicitly permits it.
- **CODE-REQ-093:** Build artifacts shall not be committed unless an approved repository rule requires them.
- **CODE-REQ-094:** Temporary files shall not become authoritative source.
- **CODE-REQ-095:** Local environment files shall not be committed when they contain secrets or machine-specific state.
- **CODE-REQ-096:** Repository structure shall support code ownership and review routing.
- **CODE-REQ-097:** Repository structure shall support test discovery.
- **CODE-REQ-098:** Repository structure shall support dependency and secret scanning.
- **CODE-REQ-099:** Repository structure shall remain compatible with document 24 after consolidation.
- **CODE-REQ-100:** Structural changes with architectural impact shall require an ADR where applicable.

## 6. Public Interfaces

- **CODE-REQ-101:** Every public interface shall be documented.
- **CODE-REQ-102:** Public-interface documentation shall state purpose.
- **CODE-REQ-103:** Public-interface documentation shall state inputs.
- **CODE-REQ-104:** Public-interface documentation shall state outputs.
- **CODE-REQ-105:** Public-interface documentation shall state side effects.
- **CODE-REQ-106:** Public-interface documentation shall state authorization requirements.
- **CODE-REQ-107:** Public-interface documentation shall state error behavior.
- **CODE-REQ-108:** Public-interface documentation shall state security considerations.
- **CODE-REQ-109:** Public-interface documentation shall state privacy and data-handling considerations where applicable.
- **CODE-REQ-110:** Public-interface documentation shall state evidence and provenance effects where applicable.
- **CODE-REQ-111:** Public-interface documentation shall state idempotency behavior where applicable.
- **CODE-REQ-112:** Public-interface documentation shall state timeout and cancellation behavior where applicable.
- **CODE-REQ-113:** Public-interface documentation shall state version and compatibility expectations.
- **CODE-REQ-114:** Public interfaces shall use explicit types or schemas where the language supports them.
- **CODE-REQ-115:** Public interfaces shall reject unsupported inputs.
- **CODE-REQ-116:** Public interfaces shall not expose internal secret or credential representations.
- **CODE-REQ-117:** Public interfaces shall not expose mutable internal state unnecessarily.
- **CODE-REQ-118:** Public interfaces shall not broaden authorization silently.
- **CODE-REQ-119:** Public interfaces shall not depend on undocumented global state.
- **CODE-REQ-120:** Public interfaces shall use stable identifiers rather than display labels where authority is material.
- **CODE-REQ-121:** Breaking interface changes shall have migration and impact analysis.
- **CODE-REQ-122:** Deprecated interfaces shall identify replacement and removal conditions.
- **CODE-REQ-123:** Public interfaces shall have positive and negative tests.
- **CODE-REQ-124:** Security-critical public interfaces shall have abuse-case tests.
- **CODE-REQ-125:** Examples shall use synthetic and non-secret data.

## 7. Types, Schemas and Data Contracts

- **CODE-REQ-126:** Material data structures shall have explicit schemas or type definitions.
- **CODE-REQ-127:** Security-critical identifiers shall use distinct types or validated wrappers where feasible.
- **CODE-REQ-128:** Human, agent, workload and connector identities shall remain distinguishable in code.
- **CODE-REQ-129:** Case, purpose, jurisdiction and resource identifiers shall remain distinguishable.
- **CODE-REQ-130:** Evidence identifiers shall remain distinguishable from filenames and display labels.
- **CODE-REQ-131:** Authorization decisions shall use controlled outcome types.
- **CODE-REQ-132:** Unknown enum values shall be handled explicitly.
- **CODE-REQ-133:** Optional fields shall not silently default to broader authority.
- **CODE-REQ-134:** Missing fields shall not be confused with empty or zero values.
- **CODE-REQ-135:** Null handling shall be explicit.
- **CODE-REQ-136:** Schema versions shall be identified where compatibility matters.
- **CODE-REQ-137:** Schema migrations shall preserve semantic meaning.
- **CODE-REQ-138:** Schema downgrade shall not bypass controls.
- **CODE-REQ-139:** Deserialization shall validate schema before business logic.
- **CODE-REQ-140:** Serialization shall preserve required provenance and identifiers.
- **CODE-REQ-141:** Sensitive fields shall be identified in schemas.
- **CODE-REQ-142:** Secrets shall not be represented by ordinary printable string types where stronger abstractions are feasible.
- **CODE-REQ-143:** Time values shall identify timezone or offset where material.
- **CODE-REQ-144:** Durations shall not be confused with timestamps.
- **CODE-REQ-145:** Monetary or digital-asset values shall use safe numeric representations.
- **CODE-REQ-146:** Unbounded collections shall be prohibited at external boundaries.
- **CODE-REQ-147:** Recursive structures shall have bounded depth.
- **CODE-REQ-148:** Data contracts shall identify maximum sizes where risk requires it.
- **CODE-REQ-149:** Schema examples shall use synthetic data.
- **CODE-REQ-150:** Schema changes shall trigger compatibility and security tests.

## 8. Input Validation

- **CODE-REQ-151:** All untrusted inputs shall be validated.
- **CODE-REQ-152:** Validation shall occur before privileged or side-effecting behavior.
- **CODE-REQ-153:** Validation shall use allowlists where practical.
- **CODE-REQ-154:** Validation shall enforce type.
- **CODE-REQ-155:** Validation shall enforce length.
- **CODE-REQ-156:** Validation shall enforce range.
- **CODE-REQ-157:** Validation shall enforce format.
- **CODE-REQ-158:** Validation shall enforce encoding.
- **CODE-REQ-159:** Validation shall enforce cardinality.
- **CODE-REQ-160:** Validation shall enforce nesting depth.
- **CODE-REQ-161:** Validation shall enforce required fields.
- **CODE-REQ-162:** Validation shall reject unknown security-sensitive fields.
- **CODE-REQ-163:** Validation shall canonicalize only through documented rules.
- **CODE-REQ-164:** Canonicalization shall precede comparisons that depend on canonical form.
- **CODE-REQ-165:** Validation shall prevent path traversal.
- **CODE-REQ-166:** Validation shall prevent command injection.
- **CODE-REQ-167:** Validation shall prevent query-language injection.
- **CODE-REQ-168:** Validation shall prevent template injection.
- **CODE-REQ-169:** Validation shall prevent header injection.
- **CODE-REQ-170:** Validation shall prevent URL and redirect injection.
- **CODE-REQ-171:** Validation shall prevent unsafe deserialization.
- **CODE-REQ-172:** Validation shall prevent archive traversal.
- **CODE-REQ-173:** Validation shall prevent resource-exhaustion payloads.
- **CODE-REQ-174:** Validation shall treat model-generated values as untrusted.
- **CODE-REQ-175:** Validation shall treat connector responses as untrusted.
- **CODE-REQ-176:** Validation shall treat stored data as potentially untrusted on re-entry.
- **CODE-REQ-177:** Validation failure shall produce explicit safe errors.
- **CODE-REQ-178:** Validation failure shall not partially execute protected operations.
- **CODE-REQ-179:** Validation rules shall be versioned when material.
- **CODE-REQ-180:** Validation tests shall include malformed, boundary, encoded, duplicate and adversarial values.

## 9. Output Validation and Encoding

- **CODE-REQ-181:** Outputs crossing trust boundaries shall be validated.
- **CODE-REQ-182:** Outputs shall conform to declared schemas.
- **CODE-REQ-183:** Outputs shall not expose secrets.
- **CODE-REQ-184:** Outputs shall not expose unnecessary personal or case data.
- **CODE-REQ-185:** Outputs shall preserve provenance where applicable.
- **CODE-REQ-186:** Outputs shall distinguish source evidence from derived analysis.
- **CODE-REQ-187:** Outputs shall distinguish model-generated content.
- **CODE-REQ-188:** Outputs shall use context-appropriate encoding.
- **CODE-REQ-189:** HTML output shall be safely encoded.
- **CODE-REQ-190:** Shell output used as input shall not be concatenated unsafely.
- **CODE-REQ-191:** SQL or query fragments shall not be produced for direct execution without separate validation and parameterization.
- **CODE-REQ-192:** URLs shall be validated before use.
- **CODE-REQ-193:** File paths shall be normalized and constrained.
- **CODE-REQ-194:** Logs shall use structured fields where practical.
- **CODE-REQ-195:** Error outputs shall not expose internal stack traces to unauthorized recipients.
- **CODE-REQ-196:** Export outputs shall include manifests and integrity values where required.
- **CODE-REQ-197:** Authorization outputs shall use controlled decision values.
- **CODE-REQ-198:** Unknown output states shall fail safely.
- **CODE-REQ-199:** Output size shall be bounded.
- **CODE-REQ-200:** Output collection size shall be bounded.
- **CODE-REQ-201:** Model outputs shall not be treated as executable instructions.
- **CODE-REQ-202:** Connector outputs shall not be treated as trusted solely because schema validation passed.
- **CODE-REQ-203:** Output redaction shall be attributable where material.
- **CODE-REQ-204:** Output transformations shall preserve lineage.
- **CODE-REQ-205:** Output validation shall have negative tests.

## 10. Error Handling

- **CODE-REQ-206:** Errors shall be handled explicitly.
- **CODE-REQ-207:** Security-critical failures shall not be ignored.
- **CODE-REQ-208:** Exceptions shall not be swallowed silently.
- **CODE-REQ-209:** Error handling shall preserve the original cause where safe.
- **CODE-REQ-210:** Error handling shall avoid exposing secrets.
- **CODE-REQ-211:** Error handling shall avoid exposing unnecessary personal or case data.
- **CODE-REQ-212:** Error handling shall distinguish user errors from system failures.
- **CODE-REQ-213:** Error handling shall distinguish authorization denial from technical failure.
- **CODE-REQ-214:** Error handling shall distinguish integrity failure from parsing failure.
- **CODE-REQ-215:** Error handling shall distinguish retryable from terminal failure.
- **CODE-REQ-216:** Error handling shall avoid broad catch-all behavior that hides programming defects.
- **CODE-REQ-217:** Catch-all handling used at process boundaries shall log or preserve safe diagnostic context.
- **CODE-REQ-218:** Error messages shall be actionable for authorized operators.
- **CODE-REQ-219:** External error messages shall be minimized and safe.
- **CODE-REQ-220:** Errors shall use stable codes where automated handling is required.
- **CODE-REQ-221:** Unknown errors shall resolve to safe failure.
- **CODE-REQ-222:** Cleanup shall occur after failure.
- **CODE-REQ-223:** Cleanup failure shall remain visible.
- **CODE-REQ-224:** Rollback shall not restore revoked authority.
- **CODE-REQ-225:** Partial success shall be represented explicitly.
- **CODE-REQ-226:** Retries shall not occur after authorization denial without fresh context.
- **CODE-REQ-227:** Retries shall be bounded.
- **CODE-REQ-228:** Repeated failure shall trigger circuit breaking or containment where appropriate.
- **CODE-REQ-229:** Error paths shall be tested.
- **CODE-REQ-230:** Failure metrics shall not conceal individual critical errors.

## 11. Secrets and Credentials

- **CODE-REQ-231:** Hard-coded secrets shall be prohibited.
- **CODE-REQ-232:** Secrets shall not be committed to version control.
- **CODE-REQ-233:** Secrets shall not appear in source comments.
- **CODE-REQ-234:** Secrets shall not appear in test fixtures.
- **CODE-REQ-235:** Secrets shall not appear in example configuration.
- **CODE-REQ-236:** Secrets shall not appear in logs.
- **CODE-REQ-237:** Secrets shall not appear in exceptions.
- **CODE-REQ-238:** Secrets shall not appear in telemetry.
- **CODE-REQ-239:** Secrets shall not appear in prompts.
- **CODE-REQ-240:** Secrets shall not appear in screenshots or generated documentation.
- **CODE-REQ-241:** Human credentials shall not be embedded in agent or connector code.
- **CODE-REQ-242:** Secret identifiers may be logged, but secret values shall not.
- **CODE-REQ-243:** Secret retrieval shall be authenticated and authorized.
- **CODE-REQ-244:** Secret access shall use least privilege.
- **CODE-REQ-245:** Secret lifetime shall be bounded where supported.
- **CODE-REQ-246:** Secrets shall be independently revocable.
- **CODE-REQ-247:** Secret rotation shall be supported.
- **CODE-REQ-248:** Secret rotation shall not require source changes.
- **CODE-REQ-249:** Development, test, demonstration and operational secrets shall remain separate.
- **CODE-REQ-250:** Secret placeholders shall be unmistakably non-secret.
- **CODE-REQ-251:** Secret comparison shall avoid avoidable timing leakage where relevant.
- **CODE-REQ-252:** Secret material shall be cleared from temporary storage where feasible.
- **CODE-REQ-253:** Secret storage failure shall fail closed.
- **CODE-REQ-254:** Repository history shall be included in secret scanning.
- **CODE-REQ-255:** Secret exposure shall trigger containment, rotation and impact review.
- **CODE-REQ-256:** Generated code shall not invent or embed credentials.
- **CODE-REQ-257:** Environment variables shall not be assumed secure storage by themselves.
- **CODE-REQ-258:** Command-line arguments shall not carry secrets when process listings or logs can expose them.
- **CODE-REQ-259:** Secret-bearing files shall use restrictive permissions.
- **CODE-REQ-260:** Secret-handling tests shall verify redaction and failure behavior.

## 12. Cryptographic Code

- **CODE-REQ-261:** Custom cryptographic algorithms shall not be created for protection purposes.
- **CODE-REQ-262:** Approved well-reviewed cryptographic libraries shall be used.
- **CODE-REQ-263:** Cryptographic algorithm selection shall be governed by approved architecture or ADRs.
- **CODE-REQ-264:** Deprecated algorithms shall not be introduced.
- **CODE-REQ-265:** Key sizes and parameters shall be explicit and reviewed.
- **CODE-REQ-266:** Random values used for security shall come from cryptographically secure generators.
- **CODE-REQ-267:** Security tokens shall have sufficient entropy.
- **CODE-REQ-268:** Nonces shall not be reused where reuse breaks security properties.
- **CODE-REQ-269:** Initialization vectors shall follow algorithm requirements.
- **CODE-REQ-270:** Keys shall not be reused across incompatible purposes.
- **CODE-REQ-271:** Encryption shall use integrity protection where required.
- **CODE-REQ-272:** Cryptographic verification shall fail closed.
- **CODE-REQ-273:** Signature verification shall validate algorithm, key, context and canonical representation.
- **CODE-REQ-274:** Hash comparison shall use safe comparison where relevant.
- **CODE-REQ-275:** Password storage, where applicable, shall use approved password-hashing methods rather than reversible encryption.
- **CODE-REQ-276:** Key derivation shall use approved methods.
- **CODE-REQ-277:** Certificate validation shall not be disabled.
- **CODE-REQ-278:** Hostname verification shall not be disabled.
- **CODE-REQ-279:** Test-only certificate bypasses shall not reach release configurations.
- **CODE-REQ-280:** Cryptographic errors shall not expose key material.
- **CODE-REQ-281:** Cryptographic metadata shall identify algorithms and versions.
- **CODE-REQ-282:** Key identifiers shall remain distinct from key values.
- **CODE-REQ-283:** Cryptographic migrations shall preserve old and new verification evidence where required.
- **CODE-REQ-284:** Cryptographic code shall receive specialist review.
- **CODE-REQ-285:** Cryptographic tests shall include invalid keys, altered data, expiry and algorithm mismatch.

## 13. Authentication, Authorization and Policy Code

- **CODE-REQ-286:** Authentication code shall remain distinct from authorization code.
- **CODE-REQ-287:** Authorization code shall default to deny.
- **CODE-REQ-288:** Authorization shall occur at the point of use.
- **CODE-REQ-289:** Policy decisions shall identify exact policy versions.
- **CODE-REQ-290:** Policy enforcement shall reject unknown decision values.
- **CODE-REQ-291:** `Not Applicable` shall not be treated as `Permit`.
- **CODE-REQ-292:** `Indeterminate` shall resolve safely.
- **CODE-REQ-293:** Authorization caches shall have expiry and invalidation.
- **CODE-REQ-294:** Revocation shall invalidate affected cached decisions.
- **CODE-REQ-295:** Policy failure shall not produce permissive fallback.
- **CODE-REQ-296:** Subject, resource and action identifiers shall be validated.
- **CODE-REQ-297:** Case, purpose, jurisdiction, time, tool and data scope shall be enforced where applicable.
- **CODE-REQ-298:** Human approval shall be explicit and attributable.
- **CODE-REQ-299:** Silence and timeout shall not be treated as approval.
- **CODE-REQ-300:** Obligations shall be enforced before or during execution.
- **CODE-REQ-301:** Unknown mandatory obligations shall cause denial.
- **CODE-REQ-302:** Denied actions shall not execute partially.
- **CODE-REQ-303:** Denied actions shall not retry through weaker paths.
- **CODE-REQ-304:** Administrative authorization shall be distinct from ordinary use.
- **CODE-REQ-305:** Break-glass paths shall be explicit, narrow and audited.
- **CODE-REQ-306:** Policy parsing shall reject malformed policies.
- **CODE-REQ-307:** Policy publication and rollback shall be protected operations.
- **CODE-REQ-308:** Authorization logs shall avoid unnecessary sensitive details.
- **CODE-REQ-309:** Authorization code shall have exhaustive negative tests.
- **CODE-REQ-310:** Authorization code shall conform to `OBDIA-AUTH-001`.

## 14. AI, Prompt, Retrieval, Memory and Tool Code

- **CODE-REQ-311:** Model output shall be treated as untrusted.
- **CODE-REQ-312:** Retrieved content shall be treated as untrusted.
- **CODE-REQ-313:** Stored memory shall be treated as untrusted on reuse.
- **CODE-REQ-314:** Prompt content shall remain separate from policy and authorization data.
- **CODE-REQ-315:** External instructions shall not override system or governance authority.
- **CODE-REQ-316:** Tool invocations shall pass through separate authorization enforcement.
- **CODE-REQ-317:** Tool arguments shall use constrained schemas.
- **CODE-REQ-318:** Tool destinations shall be validated.
- **CODE-REQ-319:** Tool outputs shall be validated.
- **CODE-REQ-320:** Model output shall not directly execute shell commands.
- **CODE-REQ-321:** Model output shall not directly execute database queries.
- **CODE-REQ-322:** Model output shall not directly issue credentials.
- **CODE-REQ-323:** Model output shall not directly approve lifecycle or governance transitions.
- **CODE-REQ-324:** Model output shall not directly alter original evidence.
- **CODE-REQ-325:** Model output shall not directly authorize exports, deletion or public release.
- **CODE-REQ-326:** Prompt templates shall be versioned where material.
- **CODE-REQ-327:** Model and provider versions shall be recorded where material.
- **CODE-REQ-328:** Memory shall be scoped by officer, case and purpose.
- **CODE-REQ-329:** Cross-case memory access shall be denied by default.
- **CODE-REQ-330:** Prompt-injection boundaries shall be explicit.
- **CODE-REQ-331:** Untrusted content shall be labeled in model context where feasible.
- **CODE-REQ-332:** AI-generated summaries shall be marked as derived analytical output.
- **CODE-REQ-333:** AI-assisted code paths shall preserve human accountability.
- **CODE-REQ-334:** Unsafe model behavior shall support suspension and containment.
- **CODE-REQ-335:** AI and tool code shall have adversarial and prompt-injection tests.

## 15. Connector Implementation Code

- **CODE-REQ-336:** Connector code shall conform to `OBDIA-CONN-001`.
- **CODE-REQ-337:** Connector code shall use distinct machine identities where feasible.
- **CODE-REQ-338:** Connector code shall not use human credentials.
- **CODE-REQ-339:** Connector code shall enforce endpoint allowlists.
- **CODE-REQ-340:** Connector code shall validate redirects.
- **CODE-REQ-341:** Connector code shall constrain network egress.
- **CODE-REQ-342:** Connector code shall validate request schemas.
- **CODE-REQ-343:** Connector code shall validate response schemas.
- **CODE-REQ-344:** Connector code shall bound response sizes.
- **CODE-REQ-345:** Connector code shall bound retries.
- **CODE-REQ-346:** Connector code shall implement timeouts.
- **CODE-REQ-347:** Side-effecting connector code shall implement idempotency or reconciliation where supported.
- **CODE-REQ-348:** Connector code shall preserve request-response correlation.
- **CODE-REQ-349:** Connector code shall preserve raw-to-normalized provenance.
- **CODE-REQ-350:** Connector code shall isolate failures.
- **CODE-REQ-351:** Connector code shall support suspension.
- **CODE-REQ-352:** Connector code shall support credential revocation.
- **CODE-REQ-353:** Connector code shall support session invalidation.
- **CODE-REQ-354:** Connector code shall not automatically follow untrusted links.
- **CODE-REQ-355:** Connector code shall not execute returned active content.
- **CODE-REQ-356:** Connector code shall not broaden permission on fallback.
- **CODE-REQ-357:** Connector code shall handle partial and paginated responses explicitly.
- **CODE-REQ-358:** Connector code shall log safely.
- **CODE-REQ-359:** Connector code shall have malicious-response tests.
- **CODE-REQ-360:** Connector code shall have network and revocation failure tests.

## 16. Evidence-Handling Code

- **CODE-REQ-361:** Evidence-handling code shall conform to `OBDIA-EVID-001`.
- **CODE-REQ-362:** Original evidence shall not be modified in place.
- **CODE-REQ-363:** Derived artifacts shall receive distinct identifiers.
- **CODE-REQ-364:** Derived artifacts shall preserve parent lineage.
- **CODE-REQ-365:** Acquisition events shall be attributable.
- **CODE-REQ-366:** Transformation events shall be attributable.
- **CODE-REQ-367:** Custody events shall be attributable.
- **CODE-REQ-368:** Integrity algorithms and values shall be explicit.
- **CODE-REQ-369:** Integrity mismatches shall trigger quarantine.
- **CODE-REQ-370:** Integrity mismatch shall not be repaired silently.
- **CODE-REQ-371:** Evidence storage shall prevent silent overwrite.
- **CODE-REQ-372:** Evidence exports shall include manifests where required.
- **CODE-REQ-373:** Evidence access shall default to deny.
- **CODE-REQ-374:** Evidence deletion shall require exceptional explicit authorization.
- **CODE-REQ-375:** Retention and legal-hold behavior shall be explicit.
- **CODE-REQ-376:** AI-generated content shall not be stored as original evidence.
- **CODE-REQ-377:** Confidence shall remain distinct from integrity and authenticity.
- **CODE-REQ-378:** Timestamp sources and uncertainty shall be preserved.
- **CODE-REQ-379:** Evidence errors shall not fabricate complete acquisition.
- **CODE-REQ-380:** Partial acquisition shall be labeled.
- **CODE-REQ-381:** Evidence identifiers shall not be derived solely from unsafe filenames.
- **CODE-REQ-382:** Evidence filenames shall be normalized safely.
- **CODE-REQ-383:** Evidence logs shall not contain unnecessary source content.
- **CODE-REQ-384:** Evidence code shall have altered-object and custody-gap tests.
- **CODE-REQ-385:** Evidence code shall have backup, restore and quarantine tests.

## 17. Privacy and Data Protection Code

- **CODE-REQ-386:** Code shall collect only data necessary for approved purpose.
- **CODE-REQ-387:** Code shall process only authorized data scope.
- **CODE-REQ-388:** Code shall minimize personal data in memory, storage and logs.
- **CODE-REQ-389:** Code shall avoid copying sensitive data unnecessarily.
- **CODE-REQ-390:** Code shall define retention for temporary and persistent data.
- **CODE-REQ-391:** Code shall implement deletion where required.
- **CODE-REQ-392:** Deletion code shall preserve required audit and disposition evidence.
- **CODE-REQ-393:** Code shall distinguish pseudonymization from anonymization.
- **CODE-REQ-394:** Code shall not represent anonymized data as risk-free.
- **CODE-REQ-395:** Code shall enforce case and purpose isolation.
- **CODE-REQ-396:** Cross-case queries shall require explicit authorization.
- **CODE-REQ-397:** Cross-purpose reuse shall require explicit authorization.
- **CODE-REQ-398:** Data exports shall be separately authorized.
- **CODE-REQ-399:** Telemetry shall be disabled, minimized or explicitly governed.
- **CODE-REQ-400:** Analytics identifiers shall be minimized.
- **CODE-REQ-401:** Debugging shall not expose sensitive data.
- **CODE-REQ-402:** Error reporting shall redact sensitive values.
- **CODE-REQ-403:** Test data shall be synthetic or explicitly authorized.
- **CODE-REQ-404:** Fixtures shall not include real credentials or real case subjects.
- **CODE-REQ-405:** Public demonstrations shall not process real case data.
- **CODE-REQ-406:** Data maps shall remain consistent with code behavior.
- **CODE-REQ-407:** Privacy-impacting changes shall trigger review.
- **CODE-REQ-408:** Data-protection defaults shall be restrictive.
- **CODE-REQ-409:** Privacy controls shall have tests.
- **CODE-REQ-410:** Privacy incidents shall trigger containment and reassessment.

## 18. Logging, Audit and Observability Code

- **CODE-REQ-411:** Security-relevant operations shall produce structured audit events.
- **CODE-REQ-412:** Audit events shall use immutable identifiers.
- **CODE-REQ-413:** Audit events shall identify actor type and identifier.
- **CODE-REQ-414:** Audit events shall identify action and resource.
- **CODE-REQ-415:** Audit events shall identify decision and result.
- **CODE-REQ-416:** Audit events shall identify policy version where applicable.
- **CODE-REQ-417:** Audit events shall identify time and ordering information.
- **CODE-REQ-418:** Audit events shall preserve correlation identifiers.
- **CODE-REQ-419:** Audit events shall not contain secrets.
- **CODE-REQ-420:** Audit events shall minimize personal and case data.
- **CODE-REQ-421:** Logs shall distinguish audit, diagnostic and business telemetry.
- **CODE-REQ-422:** Diagnostic logs shall not substitute for audit records.
- **CODE-REQ-423:** Log levels shall be used consistently.
- **CODE-REQ-424:** Security failures shall not be logged only at debug level.
- **CODE-REQ-425:** Logging failures shall be visible.
- **CODE-REQ-426:** Audit failures shall block or restrict privileged operation where required.
- **CODE-REQ-427:** Logging shall avoid unsafe format-string construction.
- **CODE-REQ-428:** Log injection shall be prevented.
- **CODE-REQ-429:** Untrusted values shall be encoded or structured safely.
- **CODE-REQ-430:** Stack traces shall be restricted.
- **CODE-REQ-431:** Trace sampling shall not omit mandatory audit events.
- **CODE-REQ-432:** Metrics shall not replace event evidence.
- **CODE-REQ-433:** Monitoring identifiers shall correlate with lifecycle, connector and evidence records.
- **CODE-REQ-434:** Log retention shall be governed.
- **CODE-REQ-435:** Logging and audit behavior shall be tested.

## 19. File, Path and Archive Safety

- **CODE-REQ-436:** File paths from untrusted input shall be validated.
- **CODE-REQ-437:** Path traversal shall be prevented.
- **CODE-REQ-438:** Symbolic-link behavior shall be considered.
- **CODE-REQ-439:** File operations shall use least-privileged directories.
- **CODE-REQ-440:** Temporary files shall use safe creation mechanisms.
- **CODE-REQ-441:** Temporary files shall have restrictive permissions where required.
- **CODE-REQ-442:** Temporary files shall be removed according to policy.
- **CODE-REQ-443:** Sensitive temporary content shall not persist longer than necessary.
- **CODE-REQ-444:** File names shall be normalized without losing provenance.
- **CODE-REQ-445:** File extensions shall not be trusted as content type.
- **CODE-REQ-446:** Media types shall be verified where material.
- **CODE-REQ-447:** File sizes shall be bounded.
- **CODE-REQ-448:** File counts shall be bounded.
- **CODE-REQ-449:** Archive extraction shall prevent traversal.
- **CODE-REQ-450:** Archive extraction shall prevent decompression bombs.
- **CODE-REQ-451:** Nested archives shall be bounded.
- **CODE-REQ-452:** Executable permissions shall not be added unnecessarily.
- **CODE-REQ-453:** Downloaded files shall not execute automatically.
- **CODE-REQ-454:** File writes shall be atomic where integrity requires it.
- **CODE-REQ-455:** Concurrent writes shall be controlled.
- **CODE-REQ-456:** Evidence source files shall not be overwritten.
- **CODE-REQ-457:** File deletion shall be explicit and authorized.
- **CODE-REQ-458:** File-system errors shall be handled explicitly.
- **CODE-REQ-459:** File and archive code shall have adversarial tests.
- **CODE-REQ-460:** Platform-specific path behavior shall be tested where supported.

## 20. Network and Inter-Process Code

- **CODE-REQ-461:** Network destinations shall be explicitly constrained.
- **CODE-REQ-462:** Network timeouts shall be bounded.
- **CODE-REQ-463:** Network retries shall be bounded.
- **CODE-REQ-464:** TLS verification shall not be disabled.
- **CODE-REQ-465:** Hostname verification shall not be disabled.
- **CODE-REQ-466:** Redirects shall be revalidated.
- **CODE-REQ-467:** Proxy use shall be explicit.
- **CODE-REQ-468:** Proxy bypass shall be controlled.
- **CODE-REQ-469:** SSRF protections shall cover local, link-local and metadata destinations.
- **CODE-REQ-470:** Request and response sizes shall be bounded.
- **CODE-REQ-471:** Network code shall handle partial reads and writes.
- **CODE-REQ-472:** Network code shall distinguish timeout, refusal and protocol error.
- **CODE-REQ-473:** Network failure shall not trigger insecure fallback.
- **CODE-REQ-474:** Fallback destinations shall be pre-approved.
- **CODE-REQ-475:** Inter-process communication shall authenticate peers where required.
- **CODE-REQ-476:** IPC messages shall have schemas.
- **CODE-REQ-477:** IPC messages shall be validated.
- **CODE-REQ-478:** IPC permissions shall be least-privileged.
- **CODE-REQ-479:** Socket and pipe paths shall be protected.
- **CODE-REQ-480:** Child-process execution shall avoid shell interpretation where practical.
- **CODE-REQ-481:** Command arguments shall be passed as structured arrays.
- **CODE-REQ-482:** Environment inheritance shall be minimized.
- **CODE-REQ-483:** Child-process output shall be treated as untrusted.
- **CODE-REQ-484:** Process termination and cleanup shall be handled.
- **CODE-REQ-485:** Network and IPC code shall have failure and injection tests.

## 21. Database, Query and Persistence Code

- **CODE-REQ-486:** Database queries shall use parameterization.
- **CODE-REQ-487:** Untrusted values shall not be concatenated into queries.
- **CODE-REQ-488:** Dynamic identifiers shall use allowlists and safe construction.
- **CODE-REQ-489:** Database accounts shall use least privilege.
- **CODE-REQ-490:** Read and write accounts shall be separated where appropriate.
- **CODE-REQ-491:** Administrative database access shall be separately controlled.
- **CODE-REQ-492:** Transactions shall be used where atomicity is required.
- **CODE-REQ-493:** Transaction boundaries shall be explicit.
- **CODE-REQ-494:** Rollback failure shall remain visible.
- **CODE-REQ-495:** Retry behavior shall account for duplicate side effects.
- **CODE-REQ-496:** Migration scripts shall be versioned.
- **CODE-REQ-497:** Migration scripts shall have rollback or forward-recovery plans.
- **CODE-REQ-498:** Migrations shall preserve security constraints.
- **CODE-REQ-499:** Migrations shall preserve evidence and audit history where applicable.
- **CODE-REQ-500:** Data constraints shall be enforced at appropriate layers.
- **CODE-REQ-501:** Unique and foreign-key constraints shall support integrity where applicable.
- **CODE-REQ-502:** Sensitive fields shall use appropriate storage protection.
- **CODE-REQ-503:** Connection strings shall not contain hard-coded secrets.
- **CODE-REQ-504:** Database errors shall not expose schema or sensitive values externally.
- **CODE-REQ-505:** Query result sizes shall be bounded.
- **CODE-REQ-506:** Pagination shall preserve authorization scope.
- **CODE-REQ-507:** Bulk operations shall require stronger controls.
- **CODE-REQ-508:** Data deletion shall respect retention and legal hold.
- **CODE-REQ-509:** Persistence code shall have concurrency and failure tests.
- **CODE-REQ-510:** Migration tests shall use controlled data.

## 22. Serialization and Deserialization

- **CODE-REQ-511:** Serialization formats shall be explicitly selected.
- **CODE-REQ-512:** Deserialization shall use safe libraries and modes.
- **CODE-REQ-513:** Executable-object deserialization shall be prohibited unless explicitly isolated and approved.
- **CODE-REQ-514:** Deserialization shall validate schema before object use.
- **CODE-REQ-515:** Unknown types shall be rejected.
- **CODE-REQ-516:** Unknown security-sensitive fields shall be rejected.
- **CODE-REQ-517:** Object graphs shall have bounded depth.
- **CODE-REQ-518:** Collection sizes shall be bounded.
- **CODE-REQ-519:** Numeric ranges shall be validated.
- **CODE-REQ-520:** Encoding shall be validated.
- **CODE-REQ-521:** Duplicate keys shall have defined handling.
- **CODE-REQ-522:** Canonicalization shall be documented where signatures or hashes depend on it.
- **CODE-REQ-523:** Version fields shall be validated.
- **CODE-REQ-524:** Schema downgrade shall not bypass controls.
- **CODE-REQ-525:** Polymorphic type selection shall be allowlisted.
- **CODE-REQ-526:** Deserialization errors shall fail safely.
- **CODE-REQ-527:** Partial objects shall not reach protected logic.
- **CODE-REQ-528:** Serialized secrets shall be prohibited unless an approved protected format is required.
- **CODE-REQ-529:** Serialized personal data shall be minimized.
- **CODE-REQ-530:** Integrity protection shall be applied where serialized objects cross untrusted boundaries and risk requires it.
- **CODE-REQ-531:** Compression shall be bounded.
- **CODE-REQ-532:** Serialization shall preserve required identifiers and provenance.
- **CODE-REQ-533:** Serialization changes shall have compatibility tests.
- **CODE-REQ-534:** Malicious payload tests shall be included.
- **CODE-REQ-535:** Fuzz testing should be considered for high-risk parsers.

## 23. Concurrency, State and Idempotency

- **CODE-REQ-536:** Shared mutable state shall be minimized.
- **CODE-REQ-537:** State ownership shall be explicit.
- **CODE-REQ-538:** Concurrent access shall use appropriate synchronization.
- **CODE-REQ-539:** Lock ordering shall be documented where deadlock risk is material.
- **CODE-REQ-540:** Race conditions affecting authorization or lifecycle shall be treated as critical.
- **CODE-REQ-541:** State transitions shall validate expected prior state.
- **CODE-REQ-542:** Concurrent restrictive and permissive transitions shall resolve restrictively.
- **CODE-REQ-543:** Revocation shall win over activation or continued operation.
- **CODE-REQ-544:** Idempotency shall be implemented for retried side-effecting operations where feasible.
- **CODE-REQ-545:** Idempotency keys shall be unique and scoped.
- **CODE-REQ-546:** Duplicate execution shall be detected or reconciled.
- **CODE-REQ-547:** Uncertain results shall not be retried blindly.
- **CODE-REQ-548:** State machines shall reject invalid transitions.
- **CODE-REQ-549:** Unknown state shall resolve safely.
- **CODE-REQ-550:** State persistence shall be atomic where required.
- **CODE-REQ-551:** Distributed state shall identify authoritative source and version.
- **CODE-REQ-552:** Stale state shall not permit privileged operation.
- **CODE-REQ-553:** Cache invalidation shall include revocation and policy changes.
- **CODE-REQ-554:** Time-of-check/time-of-use risks shall be addressed.
- **CODE-REQ-555:** Long-running operations shall re-evaluate material context.
- **CODE-REQ-556:** Cancellation shall preserve consistency.
- **CODE-REQ-557:** Shutdown shall flush or preserve required audit evidence.
- **CODE-REQ-558:** Concurrency failures shall be observable.
- **CODE-REQ-559:** Concurrency behavior shall have race and stress tests.
- **CODE-REQ-560:** Deterministic test harnesses shall be used where practical.

## 24. Resource Management and Denial-of-Service Resistance

- **CODE-REQ-561:** External inputs shall have size limits.
- **CODE-REQ-562:** Collections shall have item-count limits.
- **CODE-REQ-563:** Recursion shall have depth limits.
- **CODE-REQ-564:** Loops over untrusted data shall have bounded work.
- **CODE-REQ-565:** Network operations shall have timeouts.
- **CODE-REQ-566:** Subprocesses shall have timeouts.
- **CODE-REQ-567:** Memory usage shall be bounded where practical.
- **CODE-REQ-568:** CPU-intensive work shall be constrained.
- **CODE-REQ-569:** Disk usage shall be constrained.
- **CODE-REQ-570:** Temporary storage shall have limits.
- **CODE-REQ-571:** Concurrency shall have limits.
- **CODE-REQ-572:** Queues shall have capacity limits.
- **CODE-REQ-573:** Backpressure shall be implemented where appropriate.
- **CODE-REQ-574:** Rate limits shall be applied where abuse or cost risk exists.
- **CODE-REQ-575:** Cost and quota limits shall be applied to paid external services where applicable.
- **CODE-REQ-576:** Large model inputs and outputs shall be bounded.
- **CODE-REQ-577:** Archive and decompression work shall be bounded.
- **CODE-REQ-578:** Regex patterns shall avoid catastrophic behavior.
- **CODE-REQ-579:** User-controlled algorithms shall be evaluated for complexity.
- **CODE-REQ-580:** Resource exhaustion shall fail safely.
- **CODE-REQ-581:** Partial work shall not be represented as complete.
- **CODE-REQ-582:** Capacity failure shall be observable.
- **CODE-REQ-583:** Resource cleanup shall occur on success and failure.
- **CODE-REQ-584:** Resource-limit configuration shall be versioned.
- **CODE-REQ-585:** Resource-exhaustion tests shall be included.

## 25. Determinism and Reproducibility

- **CODE-REQ-586:** Authorization and policy-critical logic shall be deterministic for the same inputs and policy version.
- **CODE-REQ-587:** Evidence integrity calculations shall be deterministic.
- **CODE-REQ-588:** Identifier generation shall have defined uniqueness behavior.
- **CODE-REQ-589:** Randomness shall be explicit.
- **CODE-REQ-590:** Security randomness shall use approved secure generators.
- **CODE-REQ-591:** Test randomness shall support recorded seeds where reproducibility matters.
- **CODE-REQ-592:** Time dependencies shall be injectable or controllable in tests where practical.
- **CODE-REQ-593:** Timezone behavior shall be explicit.
- **CODE-REQ-594:** Locale-dependent behavior shall be avoided or controlled.
- **CODE-REQ-595:** Iteration ordering shall not be assumed when the language does not guarantee it.
- **CODE-REQ-596:** Floating-point behavior shall not be used for exact legal, financial or integrity values.
- **CODE-REQ-597:** Model nondeterminism shall be documented.
- **CODE-REQ-598:** Model parameters affecting reproducibility shall be recorded where available.
- **CODE-REQ-599:** External-service nondeterminism shall be documented.
- **CODE-REQ-600:** Retries shall not hide nondeterministic side effects.
- **CODE-REQ-601:** Build inputs shall be identifiable.
- **CODE-REQ-602:** Dependency versions shall be recorded.
- **CODE-REQ-603:** Configuration versions shall be recorded.
- **CODE-REQ-604:** Generated artifacts shall trace to source and generator versions.
- **CODE-REQ-605:** Reproduction instructions shall identify environment assumptions.
- **CODE-REQ-606:** Determinism shall not be claimed when external or model behavior prevents it.
- **CODE-REQ-607:** Reproducibility limitations shall remain visible.
- **CODE-REQ-608:** Deterministic behavior shall not be achieved by unsafe fixed secrets or reused nonces.
- **CODE-REQ-609:** Reproducibility tests shall compare governed outputs where applicable.
- **CODE-REQ-610:** Non-reproducible failures shall receive diagnostic evidence.

## 26. Configuration and Secure Defaults

- **CODE-REQ-611:** Configuration shall be separate from source code where appropriate.
- **CODE-REQ-612:** Security-relevant configuration shall be versioned.
- **CODE-REQ-613:** Configuration schemas shall be explicit.
- **CODE-REQ-614:** Configuration values shall be validated.
- **CODE-REQ-615:** Unknown security-relevant configuration shall fail safely.
- **CODE-REQ-616:** Missing security-relevant configuration shall fail safely.
- **CODE-REQ-617:** Default configuration shall deny unnecessary privilege.
- **CODE-REQ-618:** Default configuration shall disable dangerous optional features.
- **CODE-REQ-619:** Debug mode shall be disabled outside approved environments.
- **CODE-REQ-620:** Verbose sensitive logging shall be disabled by default.
- **CODE-REQ-621:** External network access shall be disabled or constrained by default.
- **CODE-REQ-622:** Write and delete capabilities shall be disabled by default where read-only is sufficient.
- **CODE-REQ-623:** Telemetry shall be disabled or governed by default.
- **CODE-REQ-624:** Configuration shall not contain hard-coded secrets.
- **CODE-REQ-625:** Environment-specific configuration shall be explicit.
- **CODE-REQ-626:** Development defaults shall not reach release environments silently.
- **CODE-REQ-627:** Configuration precedence shall be documented.
- **CODE-REQ-628:** Command-line overrides shall not bypass protected policy.
- **CODE-REQ-629:** Remote configuration shall be authenticated and integrity-protected.
- **CODE-REQ-630:** Configuration reload shall be atomic where required.
- **CODE-REQ-631:** Configuration changes shall be attributable.
- **CODE-REQ-632:** Configuration rollback shall restore a known governed state.
- **CODE-REQ-633:** Configuration drift shall be detectable where feasible.
- **CODE-REQ-634:** Sample configuration shall use synthetic values.
- **CODE-REQ-635:** Secure-default tests shall verify omitted and invalid configuration.

## 27. Dependency Management

- **CODE-REQ-636:** Dependencies shall be minimized.
- **CODE-REQ-637:** Every material dependency shall have an identified purpose.
- **CODE-REQ-638:** Dependencies shall be pinned where appropriate.
- **CODE-REQ-639:** Lockfiles or equivalent resolved manifests shall be retained where applicable.
- **CODE-REQ-640:** Direct and transitive dependencies shall be inventoried.
- **CODE-REQ-641:** Dependency provenance shall be recorded.
- **CODE-REQ-642:** Dependency licenses shall be reviewed.
- **CODE-REQ-643:** Dependency maintenance status shall be reviewed.
- **CODE-REQ-644:** Unsupported dependencies shall have migration, containment or rejection decisions.
- **CODE-REQ-645:** Dependency vulnerabilities shall have tracked dispositions.
- **CODE-REQ-646:** Critical unresolved dependency risk shall block acceptance or release.
- **CODE-REQ-647:** Dependency updates shall be reviewed.
- **CODE-REQ-648:** Dependency updates shall trigger tests.
- **CODE-REQ-649:** Major dependency updates shall receive impact analysis.
- **CODE-REQ-650:** Automatic dependency updates shall not merge without required review.
- **CODE-REQ-651:** Alternative dependencies shall be considered when complexity or risk is excessive.
- **CODE-REQ-652:** Unused dependencies shall be removed.
- **CODE-REQ-653:** Development-only dependencies shall not reach runtime artifacts unnecessarily.
- **CODE-REQ-654:** Optional dependencies shall remain disabled unless required.
- **CODE-REQ-655:** Package sources and registries shall be controlled.
- **CODE-REQ-656:** Dependency confusion and namespace substitution shall be addressed.
- **CODE-REQ-657:** Typosquatting checks shall be applied where feasible.
- **CODE-REQ-658:** Vendored code shall preserve origin and version.
- **CODE-REQ-659:** Forked dependencies shall identify divergence and maintenance owner.
- **CODE-REQ-660:** Dependency exceptions shall have owner, expiry and residual risk.

## 28. Build and Supply-Chain Security

- **CODE-REQ-661:** Builds shall use attributable source commits.
- **CODE-REQ-662:** Build configuration shall be versioned.
- **CODE-REQ-663:** Build environments shall be identified.
- **CODE-REQ-664:** Build tools and versions shall be identified.
- **CODE-REQ-665:** Build pipelines shall use least privilege.
- **CODE-REQ-666:** Build credentials shall be scoped.
- **CODE-REQ-667:** Build credentials shall not appear in logs.
- **CODE-REQ-668:** Build artifacts shall correspond to reviewed source.
- **CODE-REQ-669:** Build artifacts shall have integrity references where feasible.
- **CODE-REQ-670:** Generated artifacts shall identify generator and source.
- **CODE-REQ-671:** Build steps shall not download unverified executable content silently.
- **CODE-REQ-672:** External build inputs shall have provenance.
- **CODE-REQ-673:** Build caches shall not bypass integrity controls.
- **CODE-REQ-674:** Build outputs shall not contain source secrets.
- **CODE-REQ-675:** Build outputs shall not contain unnecessary test data.
- **CODE-REQ-676:** Build outputs shall not include debug interfaces unintentionally.
- **CODE-REQ-677:** Release packaging shall not modify reviewed behavior silently.
- **CODE-REQ-678:** Reproducible or documented deterministic builds should be used where feasible.
- **CODE-REQ-679:** Non-reproducible build limitations shall be documented.
- **CODE-REQ-680:** Artifact signing may supplement but shall not replace review and validation.
- **CODE-REQ-681:** SBOM or equivalent dependency records shall be generated where implementation exists and risk warrants it.
- **CODE-REQ-682:** Build-pipeline changes shall receive review.
- **CODE-REQ-683:** Compromised build infrastructure shall trigger containment and artifact reassessment.
- **CODE-REQ-684:** Supply-chain tests shall verify artifact and dependency integrity.
- **CODE-REQ-685:** Release artifacts shall comply with `OBDIA-REL-001`.

## 29. Generated Code and AI-Assisted Development

- **CODE-REQ-686:** AI-generated code shall be treated as untrusted until reviewed.
- **CODE-REQ-687:** AI-generated code shall have an accountable human reviewer.
- **CODE-REQ-688:** AI-generated code shall not be merged solely because it compiles or passes tests.
- **CODE-REQ-689:** AI-generated code shall be checked against approved requirements.
- **CODE-REQ-690:** AI-generated code shall be checked for invented interfaces and dependencies.
- **CODE-REQ-691:** AI-generated code shall be checked for hard-coded secrets.
- **CODE-REQ-692:** AI-generated code shall be checked for unsafe defaults.
- **CODE-REQ-693:** AI-generated code shall be checked for authorization bypass.
- **CODE-REQ-694:** AI-generated code shall be checked for insecure error handling.
- **CODE-REQ-695:** AI-generated code shall be checked for injection risks.
- **CODE-REQ-696:** AI-generated code shall be checked for privacy and evidence impact.
- **CODE-REQ-697:** AI-generated tests shall not be assumed complete.
- **CODE-REQ-698:** AI-generated documentation shall not fabricate implementation or validation.
- **CODE-REQ-699:** AI-generated dependency recommendations shall receive supply-chain review.
- **CODE-REQ-700:** Prompts containing proprietary or sensitive code shall follow approved data-handling rules.
- **CODE-REQ-701:** Secret or credential material shall not be sent to external models.
- **CODE-REQ-702:** Real case data shall not be sent to coding assistants.
- **CODE-REQ-703:** AI-assistance records shall identify tool or model where material.
- **CODE-REQ-704:** AI-assisted changes shall remain attributable to the human committer.
- **CODE-REQ-705:** Large generated changes shall be decomposed for review where practical.
- **CODE-REQ-706:** Generated code shall follow formatting, linting and type-checking rules.
- **CODE-REQ-707:** Generated code shall not disable security checks.
- **CODE-REQ-708:** Generated code shall not create hidden network access.
- **CODE-REQ-709:** Generated code shall not introduce uncontrolled telemetry.
- **CODE-REQ-710:** AI-assisted coding limitations shall be disclosed in review evidence where material.

## 30. Comments and Documentation in Code

- **CODE-REQ-711:** Comments shall explain intent, constraints and non-obvious security reasoning.
- **CODE-REQ-712:** Comments shall not merely restate obvious syntax.
- **CODE-REQ-713:** Comments shall remain synchronized with behavior.
- **CODE-REQ-714:** Security assumptions shall be documented.
- **CODE-REQ-715:** Authorization assumptions shall be documented.
- **CODE-REQ-716:** Evidence and provenance assumptions shall be documented.
- **CODE-REQ-717:** Concurrency and ordering assumptions shall be documented.
- **CODE-REQ-718:** Time and clock assumptions shall be documented.
- **CODE-REQ-719:** External-service assumptions shall be documented.
- **CODE-REQ-720:** Known limitations shall be documented.
- **CODE-REQ-721:** Temporary workarounds shall identify owner and removal condition.
- **CODE-REQ-722:** TODO items affecting security shall have tracking references.
- **CODE-REQ-723:** Deprecated code shall identify replacement and removal condition.
- **CODE-REQ-724:** Public interfaces shall have structured documentation appropriate to the language.
- **CODE-REQ-725:** Examples shall use synthetic and non-secret values.
- **CODE-REQ-726:** Comments shall not contain secrets.
- **CODE-REQ-727:** Comments shall not contain unnecessary personal or case data.
- **CODE-REQ-728:** Comments shall not claim legal compliance without evidence.
- **CODE-REQ-729:** Comments shall not describe proposed controls as implemented.
- **CODE-REQ-730:** Comments shall not describe implemented controls as validated without evidence.
- **CODE-REQ-731:** Generated documentation shall identify authoritative source.
- **CODE-REQ-732:** Code diagrams shall follow document 23 after consolidation.
- **CODE-REQ-733:** Architecture references shall use exact identifiers.
- **CODE-REQ-734:** Documentation changes shall accompany material behavior changes.
- **CODE-REQ-735:** Misleading stale comments shall be treated as defects.

## 31. Testing Requirements for Code

- **CODE-REQ-736:** Security-focused tests shall be included.
- **CODE-REQ-737:** Every material feature shall have acceptance criteria.
- **CODE-REQ-738:** Unit tests shall cover critical logic.
- **CODE-REQ-739:** Integration tests shall cover trust boundaries.
- **CODE-REQ-740:** Negative tests shall cover denied and invalid inputs.
- **CODE-REQ-741:** Regression tests shall cover corrected defects.
- **CODE-REQ-742:** Authorization code shall test default deny.
- **CODE-REQ-743:** Revocation code shall test propagation.
- **CODE-REQ-744:** Evidence code shall test integrity and lineage.
- **CODE-REQ-745:** Connector code shall test malicious and malformed responses.
- **CODE-REQ-746:** Secret handling shall test redaction.
- **CODE-REQ-747:** Error paths shall be tested.
- **CODE-REQ-748:** Failure-closed behavior shall be tested.
- **CODE-REQ-749:** Boundary values shall be tested.
- **CODE-REQ-750:** Resource limits shall be tested.
- **CODE-REQ-751:** Concurrency and race conditions shall be tested where material.
- **CODE-REQ-752:** Idempotency shall be tested where retries can cause side effects.
- **CODE-REQ-753:** Rollback and recovery shall be tested.
- **CODE-REQ-754:** Tests shall use synthetic or authorized controlled data.
- **CODE-REQ-755:** Tests shall not depend on real criminal infrastructure.
- **CODE-REQ-756:** Tests shall not use production credentials.
- **CODE-REQ-757:** Test data shall be reproducible where feasible.
- **CODE-REQ-758:** Test environments shall be isolated.
- **CODE-REQ-759:** Test results shall identify exact code and dependency versions.
- **CODE-REQ-760:** Failed tests shall remain visible.
- **CODE-REQ-761:** Flaky security tests shall be fixed or treated as unresolved findings.
- **CODE-REQ-762:** Coverage metrics shall not replace risk-based test design.
- **CODE-REQ-763:** Automated tests shall not replace human review.
- **CODE-REQ-764:** Document 22 remains the detailed forward authority for testing.
- **CODE-REQ-765:** Security-critical code shall not be accepted without applicable negative tests.

## 32. Static Analysis, Linting and Formatting

- **CODE-REQ-766:** Supported languages shall use consistent automated formatting where practical.
- **CODE-REQ-767:** Formatting configuration shall be versioned.
- **CODE-REQ-768:** Linting configuration shall be versioned.
- **CODE-REQ-769:** Type-checking configuration shall be versioned where the language supports it.
- **CODE-REQ-770:** Security-focused static analysis shall be applied where practical.
- **CODE-REQ-771:** Secret scanning shall be applied.
- **CODE-REQ-772:** Dependency scanning shall be applied.
- **CODE-REQ-773:** Generated-code scanning exclusions shall be narrow and documented.
- **CODE-REQ-774:** False-positive suppressions shall have rationale.
- **CODE-REQ-775:** Security suppressions shall have owner and review condition.
- **CODE-REQ-776:** Broad rule disablement shall be prohibited without approval.
- **CODE-REQ-777:** New critical findings shall block acceptance or release.
- **CODE-REQ-778:** Analysis shall run on the exact reviewed commit.
- **CODE-REQ-779:** Tool versions shall be recorded where material.
- **CODE-REQ-780:** Tool failures shall not be treated as passing results.
- **CODE-REQ-781:** Tool output shall be retained where required for validation.
- **CODE-REQ-782:** Passing static analysis shall not prove security.
- **CODE-REQ-783:** Human review shall assess semantic and architectural issues.
- **CODE-REQ-784:** Formatter output shall not be assumed behavior-preserving without tests for material generated changes.
- **CODE-REQ-785:** Lint rules shall not encourage unsafe shortcuts.
- **CODE-REQ-786:** Type errors affecting security boundaries shall be blocking.
- **CODE-REQ-787:** Unreachable or dead security code shall be investigated.
- **CODE-REQ-788:** Unused privilege-related code shall be removed.
- **CODE-REQ-789:** Baseline suppressions shall not conceal new issues.
- **CODE-REQ-790:** Static-analysis limitations shall remain visible.

## 33. Code Review

- **CODE-REQ-791:** Material code changes shall receive human review.
- **CODE-REQ-792:** Review shall identify the exact commit.
- **CODE-REQ-793:** Review shall verify requirement and ADR alignment.
- **CODE-REQ-794:** Review shall verify threat-model impact.
- **CODE-REQ-795:** Review shall verify authorization impact.
- **CODE-REQ-796:** Review shall verify evidence impact.
- **CODE-REQ-797:** Review shall verify privacy impact.
- **CODE-REQ-798:** Review shall verify connector impact.
- **CODE-REQ-799:** Review shall verify secure defaults.
- **CODE-REQ-800:** Review shall verify input and output validation.
- **CODE-REQ-801:** Review shall verify error handling.
- **CODE-REQ-802:** Review shall verify secret handling.
- **CODE-REQ-803:** Review shall verify dependency changes.
- **CODE-REQ-804:** Review shall verify logging and audit behavior.
- **CODE-REQ-805:** Review shall verify failure and rollback behavior.
- **CODE-REQ-806:** Review shall verify tests and negative paths.
- **CODE-REQ-807:** Security-critical code shall receive Security Reviewer assessment.
- **CODE-REQ-808:** Cryptographic code shall receive specialist review.
- **CODE-REQ-809:** Generated and AI-assisted code shall receive the same review standard.
- **CODE-REQ-810:** Reviewers shall disclose role concentration.
- **CODE-REQ-811:** Self-review alone shall not be represented as independent review.
- **CODE-REQ-812:** Large changes shall be decomposed where practical.
- **CODE-REQ-813:** Unrelated changes shall not be bundled without rationale.
- **CODE-REQ-814:** Review comments and dispositions shall be retained.
- **CODE-REQ-815:** Blocking findings shall prevent merge or acceptance.
- **CODE-REQ-816:** Material changes after review shall invalidate affected approval.
- **CODE-REQ-817:** Review approval shall not establish validation.
- **CODE-REQ-818:** Review shall not rely solely on screenshots or generated summaries.
- **CODE-REQ-819:** Review exceptions shall follow document 28 after consolidation.
- **CODE-REQ-820:** Code review shall remain attributable.

## 34. Language and Artifact Profiles

- **CODE-REQ-821:** Each adopted language shall have a repository-approved profile.
- **CODE-REQ-822:** Language profiles shall identify supported versions.
- **CODE-REQ-823:** Language profiles shall identify formatter and linter expectations.
- **CODE-REQ-824:** Language profiles shall identify type-checking expectations.
- **CODE-REQ-825:** Language profiles shall identify package and dependency management.
- **CODE-REQ-826:** Language profiles shall identify testing tools and conventions.
- **CODE-REQ-827:** Language profiles shall identify secure standard-library usage.
- **CODE-REQ-828:** Language profiles shall identify prohibited or restricted features.
- **CODE-REQ-829:** Language profiles shall identify serialization risks.
- **CODE-REQ-830:** Language profiles shall identify concurrency risks.
- **CODE-REQ-831:** Language profiles shall identify secret-handling patterns.
- **CODE-REQ-832:** Language profiles shall identify error-handling conventions.
- **CODE-REQ-833:** Language profiles shall identify logging conventions.
- **CODE-REQ-834:** Language profiles shall identify build and release behavior.
- **CODE-REQ-835:** Python profiles shall require explicit dependency environments and safe subprocess use where Python is adopted.
- **CODE-REQ-836:** JavaScript or TypeScript profiles shall control dynamic evaluation and package lifecycle scripts where adopted.
- **CODE-REQ-837:** Shell profiles shall prohibit unsafe unquoted expansion and unchecked pipeline failure where shell is adopted.
- **CODE-REQ-838:** SQL profiles shall require parameterized queries where SQL is adopted.
- **CODE-REQ-839:** Infrastructure-as-code profiles shall prohibit embedded secrets and overly broad permissions where IaC is adopted.
- **CODE-REQ-840:** YAML and JSON profiles shall validate schemas and avoid unsafe implicit typing where relevant.
- **CODE-REQ-841:** Container profiles shall use least privilege and controlled base images where containers are adopted.
- **CODE-REQ-842:** Workflow profiles shall scope tokens and permissions where CI/CD workflows are adopted.
- **CODE-REQ-843:** Language-profile exceptions shall be documented.
- **CODE-REQ-844:** Unsupported language introduction shall require review.
- **CODE-REQ-845:** Language profiles shall remain implementation-specific and shall not alter this normative baseline silently.

## 35. Vulnerability and Defect Handling

- **CODE-REQ-846:** Security defects shall have immutable tracking identifiers.
- **CODE-REQ-847:** Defects shall identify affected code and versions.
- **CODE-REQ-848:** Defects shall identify severity and rationale.
- **CODE-REQ-849:** Defects shall identify exploitability assumptions.
- **CODE-REQ-850:** Defects shall identify affected threats and controls.
- **CODE-REQ-851:** Defects shall identify data, evidence and privacy impact.
- **CODE-REQ-852:** Defects shall identify owner.
- **CODE-REQ-853:** Defects shall identify remediation or containment.
- **CODE-REQ-854:** Critical defects shall block activation or release.
- **CODE-REQ-855:** Known exploitable defects shall not be concealed.
- **CODE-REQ-856:** Temporary mitigations shall have expiry and validation.
- **CODE-REQ-857:** Fixes shall include regression tests.
- **CODE-REQ-858:** Fixes shall preserve evidence and audit history.
- **CODE-REQ-859:** Security fixes shall receive review.
- **CODE-REQ-860:** Security fixes shall not introduce undocumented breaking behavior.
- **CODE-REQ-861:** Embargoed vulnerability details shall be access-controlled.
- **CODE-REQ-862:** Disclosure timing shall not fabricate security status.
- **CODE-REQ-863:** Dependency vulnerabilities shall have dispositions.
- **CODE-REQ-864:** False-positive dispositions shall include evidence.
- **CODE-REQ-865:** Accepted residual risk shall have human owner and expiry.
- **CODE-REQ-866:** An AI system shall not accept coding risk.
- **CODE-REQ-867:** Repeated defect classes shall trigger systemic review.
- **CODE-REQ-868:** Post-incident fixes shall update threat models where applicable.
- **CODE-REQ-869:** Corrected releases shall follow `OBDIA-REL-001`.
- **CODE-REQ-870:** Defect closure shall require verification evidence.

## 36. Exceptions and Deviations

- **CODE-REQ-871:** Exceptions shall be narrow.
- **CODE-REQ-872:** Exceptions shall be explicit.
- **CODE-REQ-873:** Exceptions shall be attributable.
- **CODE-REQ-874:** Exceptions shall be time-bounded.
- **CODE-REQ-875:** Exceptions shall identify affected requirements.
- **CODE-REQ-876:** Exceptions shall identify rationale.
- **CODE-REQ-877:** Exceptions shall identify security and privacy impact.
- **CODE-REQ-878:** Exceptions shall identify compensating controls.
- **CODE-REQ-879:** Exceptions shall identify validation.
- **CODE-REQ-880:** Exceptions shall identify owner.
- **CODE-REQ-881:** Exceptions shall identify expiry.
- **CODE-REQ-882:** Exceptions shall identify remediation plan.
- **CODE-REQ-883:** Exceptions shall not authorize canonical prohibitions.
- **CODE-REQ-884:** Exceptions shall not permit hard-coded production secrets.
- **CODE-REQ-885:** Exceptions shall not permit transfer of human credentials to an agent.
- **CODE-REQ-886:** Exceptions shall not permit silent authorization bypass.
- **CODE-REQ-887:** Exceptions shall not permit uncontrolled criminal-infrastructure interaction.
- **CODE-REQ-888:** Expired exceptions shall block acceptance or require renewed review.
- **CODE-REQ-889:** Repeated exceptions shall trigger architecture review.
- **CODE-REQ-890:** Exception use shall be auditable.
- **CODE-REQ-891:** Exception closure shall preserve history.
- **CODE-REQ-892:** Exception approval shall not be represented as baseline conformance.
- **CODE-REQ-893:** Emergency code changes shall receive retrospective review.
- **CODE-REQ-894:** Emergency changes shall remain narrow and reversible.
- **CODE-REQ-895:** Detailed exception governance remains a forward dependency on document 28.

## 37. Traceability

- **CODE-REQ-896:** Every material code component shall trace to requirements.
- **CODE-REQ-897:** Every security-critical component shall trace to threats.
- **CODE-REQ-898:** Every architectural implementation shall trace to accepted ADRs.
- **CODE-REQ-899:** Every public interface shall trace to its owning capability.
- **CODE-REQ-900:** Every authorization control shall trace to authorization requirements.
- **CODE-REQ-901:** Every evidence control shall trace to evidence requirements.
- **CODE-REQ-902:** Every connector control shall trace to connector requirements.
- **CODE-REQ-903:** Every test shall trace to code requirements and threats.
- **CODE-REQ-904:** Every dependency shall trace to an identified purpose.
- **CODE-REQ-905:** Every secret reference shall trace to approved secret management.
- **CODE-REQ-906:** Every build artifact shall trace to source and build evidence.
- **CODE-REQ-907:** Every code review shall identify exact commits.
- **CODE-REQ-908:** Every security finding shall trace to affected code.
- **CODE-REQ-909:** Every exception shall trace to affected requirements.
- **CODE-REQ-910:** Every release shall trace to reviewed and validated code.
- **CODE-REQ-911:** Traceability shall be bidirectional.
- **CODE-REQ-912:** Broken traceability affecting authority, security, privacy, evidence or release shall be blocking.
- **CODE-REQ-913:** Planned code shall not be represented as implemented.
- **CODE-REQ-914:** Implemented code shall not be represented as tested without evidence.
- **CODE-REQ-915:** Tested code shall not be represented as validated beyond test scope.
- **CODE-REQ-916:** Superseded code shall identify successor or removal disposition.
- **CODE-REQ-917:** Traceability shall use immutable identifiers.
- **CODE-REQ-918:** Traceability records shall not contain secrets.
- **CODE-REQ-919:** Traceability records shall minimize personal and case data.
- **CODE-REQ-920:** Baseline freeze shall validate coding traceability across documents 01–30.

## 38. Secure Coding Review Gate

- **CODE-REQ-921:** Secure-coding review shall identify the exact document and commit.
- **CODE-REQ-922:** Review shall verify preservation of all five original principles.
- **CODE-REQ-923:** Review shall verify preservation of all six original requirements.
- **CODE-REQ-924:** Review shall verify public-interface documentation rules.
- **CODE-REQ-925:** Review shall verify input and output validation rules.
- **CODE-REQ-926:** Review shall verify explicit error handling.
- **CODE-REQ-927:** Review shall verify secret and credential rules.
- **CODE-REQ-928:** Review shall verify dependency controls.
- **CODE-REQ-929:** Review shall verify security-focused testing.
- **CODE-REQ-930:** Review shall verify authorization, AI, connector and evidence code boundaries.
- **CODE-REQ-931:** Review shall verify logging and privacy controls.
- **CODE-REQ-932:** Review shall verify concurrency, determinism and resource limits.
- **CODE-REQ-933:** Review shall verify generated and AI-assisted code controls.
- **CODE-REQ-934:** Review shall verify build and supply-chain requirements.
- **CODE-REQ-935:** Review shall verify code-review and defect processes.
- **CODE-REQ-936:** Review shall verify exceptions and traceability.
- **CODE-REQ-937:** Review shall identify residual risks and forward dependencies.
- **CODE-REQ-938:** Critical findings shall block progression.
- **CODE-REQ-939:** Material post-review changes shall invalidate affected evidence.
- **CODE-REQ-940:** Role concentration shall be disclosed.
- **CODE-REQ-941:** Internal review shall not be represented as independent external assurance.
- **CODE-REQ-942:** Implementation Reviewer shall assess maintainability and implementation alignment.
- **CODE-REQ-943:** Security Reviewer shall assess security-critical controls.
- **CODE-REQ-944:** Privacy and Governance Reviewer shall assess data and rights impacts.
- **CODE-REQ-945:** Release Reviewer shall assess distributed code and examples.
- **CODE-REQ-946:** Project Founder approval shall not substitute for required specialist review.
- **CODE-REQ-947:** Automated findings shall remain advisory until adopted by an authorized human reviewer.
- **CODE-REQ-948:** Review non-applicability shall have rationale.
- **CODE-REQ-949:** Approval for Draft incorporation shall not establish code implementation or validation.
- **CODE-REQ-950:** Document 22 remains blocking for complete test-governance approval.

## 39. Minimum Validation Checklist

Before approval of this standard or acceptance of material code, confirm:

- [ ] Secure by Design is applied;
- [ ] readability and maintainability are adequate;
- [ ] deterministic behavior is used or nondeterminism is documented;
- [ ] complexity is no greater than necessary;
- [ ] every public interface is documented;
- [ ] all untrusted inputs are validated;
- [ ] outputs crossing trust boundaries are validated;
- [ ] errors are handled explicitly;
- [ ] no hard-coded secrets exist;
- [ ] human credentials are not embedded;
- [ ] dependencies are minimized, inventoried and pinned where appropriate;
- [ ] secure defaults and fail-closed behavior are present;
- [ ] authentication and authorization remain distinct;
- [ ] model output cannot authorize or execute privileged behavior directly;
- [ ] connector inputs and outputs remain untrusted;
- [ ] evidence remains immutable and provenance-preserving;
- [ ] privacy and case isolation are enforced;
- [ ] logging is attributable and secret-free;
- [ ] file, network, database and serialization boundaries are safe;
- [ ] concurrency, retries and idempotency are controlled;
- [ ] resources are bounded;
- [ ] builds and dependencies have provenance;
- [ ] generated and AI-assisted code received human review;
- [ ] security-focused and negative tests exist;
- [ ] static analysis and secret scanning completed where applicable;
- [ ] review findings and exceptions are resolved or owned;
- [ ] forward dependencies 22 and 24–30 are recorded where applicable;
- [ ] Project Founder approval exists before status becomes Approved.


## 40. Limitations

- This standard does not select a programming language, framework, dependency manager, formatter, linter, type checker, static-analysis tool, build system or CI/CD platform.
- It does not implement any control.
- It does not prove that code is secure, correct, compliant or production-ready.
- Static analysis, tests and review cannot guarantee absence of defects.
- Determinism may be limited by models, external services, distributed systems and platform behavior.
- Dependency pinning improves reproducibility but does not establish dependency trust.
- Cryptographic libraries can still be misconfigured or misused.
- Internal code review is not independent certification.
- Document 22 remains a forward dependency for complete testing governance.
- Documents 24–30 remain forward dependencies where they govern repository structure, versioning, risk, exceptions, changes and compliance.


## 41. Change Control

Every material change shall identify rationale, affected code classes, languages, interfaces, dependencies, security and privacy impact, evidence impact, migration, validation, rollback, authority and version effect.

- **CODE-REQ-951:** Editorial corrections shall use a patch version when meaning is unchanged.
- **CODE-REQ-952:** Backward-compatible substantive additions shall use a minor version.
- **CODE-REQ-953:** Incompatible coding requirements shall use a major version.
- **CODE-REQ-954:** Material secure-coding architecture decisions shall require an ADR where applicable.
- **CODE-REQ-955:** Changes shall require Project Founder approval.
- **CODE-REQ-956:** Changes shall receive Implementation Reviewer assessment.
- **CODE-REQ-957:** Security, privacy, evidence, authorization and release impacts shall receive specialist review where applicable.
- **CODE-REQ-958:** Changes shall identify affected language profiles, tools, tests and repository automation.
- **CODE-REQ-959:** Changes shall include migration and compatibility analysis.
- **CODE-REQ-960:** Changes shall include validation and rollback analysis.
- **CODE-REQ-961:** Changes shall not retroactively fabricate code-review or test evidence.
- **CODE-REQ-962:** Historical code-review and validation records shall not be silently rewritten.
- **CODE-REQ-963:** Requirement identifiers shall not be reused.
- **CODE-REQ-964:** Migration shall preserve requirement-to-code traceability.
- **CODE-REQ-965:** Forward-dependency reconciliation shall occur before this standard becomes Approved.

## 42. Consolidation Record

Version 1.0.0 consolidates the existing Coding Standard without expanding project scope. It:

- retains immutable document identifier `OBDIA-CODE-001`;
- normalizes the authoritative title and filename to `Secure Coding Standard` and `21_SECURE_CODING_STANDARD.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves Secure by Design;
- preserves Readability;
- preserves Maintainability;
- preserves Deterministic behavior;
- preserves Least complexity;
- preserves documentation of every public interface;
- preserves input validation;
- preserves explicit error handling;
- preserves the prohibition on hard-coded secrets;
- preserves dependency pinning where appropriate;
- preserves security-focused tests;
- adds applicability, code classification, interfaces, schemas, output validation, cryptography, authorization, AI, connector, evidence, privacy, audit, file, network, database, serialization, concurrency, resource, reproducibility, configuration, supply-chain, generated-code, review, language-profile, vulnerability, exception and traceability requirements;
- identifies documents 22 and 24–30 as forward dependencies where applicable;
- treats the former `_ENTERPRISE` source as a legacy source variant of the same immutable document;
- creates no implementation, language selection, production-readiness claim or operational capability.

## 43. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Implementation Reviewer; approval reserved to Project Founder | Constitutional consolidation of the original Coding Standard: preserved all five principles and six requirements; normalized the secure-coding title; added complete implementation, review, validation, supply-chain and traceability controls. |
