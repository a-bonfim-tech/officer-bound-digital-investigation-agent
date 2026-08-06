# COMPLIANCE MAPPING BASELINE

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-COMP-001 |
| **Title** | Compliance Mapping Baseline |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative for mapping method, source governance, evidence requirements and claim boundaries; individual mappings remain design and review evidence |
| **Authority** | Project Founder |
| **Owner** | Privacy and Governance Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define mandatory authoritative-source verification, applicability, classification, mapping, strength, gap, evidence, review, update, uncertainty, jurisdiction, validation, publication and traceability requirements for OBDIA mappings to cybersecurity, AI, privacy, governance and legal reference frameworks. |
| **Scope** | NIST CSF; NIST AI RMF; NIST SP 800-207; MITRE ATT&CK; MITRE ATLAS; OWASP; ISO/IEC 27001; ISO/IEC 42001; GDPR; EU AI Act; NIS2; OBDIA Knowledge Pack 01–30; and associated architecture, implementation, test, evidence, risk, research, assumption, decision and release records. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `04_PORTFOLIO_AND_PUBLICATION_POLICY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `15_RELEASE_AND_PUBLICATION_POLICY.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; `18_EVIDENCE_MODEL.md`; `19_AGENT_LIFECYCLE_MODEL.md`; `20_CONNECTOR_SECURITY_POLICY.md`; `21_SECURE_CODING_STANDARD.md`; `22_TESTING_STANDARD.md`; `23_DIAGRAM_STANDARD.md`; `24_GITHUB_REPOSITORY_STANDARD.md`; `25_DOCUMENT_VERSIONING_POLICY.md`; `26_AI_RISK_REGISTER.md`; `27_RESEARCH_BACKLOG.md`; `28_ASSUMPTIONS_REGISTER.md`; `29_DECISION_LOG_POLICY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-RES-001`; `OBDIA-PUB-001`; `OBDIA-ARCH-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-REL-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001`; `OBDIA-EVID-001`; `OBDIA-AGENT-001`; `OBDIA-CONN-001`; `OBDIA-CODE-001`; `OBDIA-TEST-001`; `OBDIA-DIAG-001`; `OBDIA-GH-001`; `OBDIA-VER-001`; `OBDIA-RISK-001`; `OBDIA-RSCH-001`; `OBDIA-ASM-001`; `OBDIA-DEC-001` |
| **Cross-References** | Source-verification records; mapping records; applicability assessments; gaps; evidence packages; risks; assumptions; research items; ADRs; decision-log entries; change records; implementation plans; test and validation evidence; baseline manifests; releases and public claims |
| **Assumptions** | The project is an independent research architecture. Legal and regulatory applicability depends on facts, roles, jurisdiction, processing, deployment and institutional context that are not established merely by this repository. |
| **Constraints** | Mappings shall not create legal conclusions, certification, conformity assessment, regulatory approval, production readiness, institutional endorsement or independent AI authority. They shall not weaken canonical prohibitions or authorize prohibited research or operations. |
| **Security Considerations** | Incorrect, stale, overstated or selectively incomplete mappings can create false assurance, conceal gaps, misstate legal duties, substitute documentation for implementation, weaken security decisions or mislead reviewers and the public. |
| **Validation Criteria** | Every mapping identifies an immutable record ID, authoritative source and version, source type, verification date, applicability, OBDIA requirement IDs, external reference IDs or objective, mapping relationship, strength, rationale, evidence, gaps, uncertainty, jurisdiction, owner, reviewer, review trigger, status and immutable history. |
| **Implementation Relationship** | This document defines mapping governance and an initial design-level crosswalk. It does not prove implementation, control effectiveness, legal compliance, ISO conformity, certification, NIST adoption, regulatory applicability or audit readiness. |

---

## 1. Purpose

The non-Enterprise source required mappings to:

- NIST CSF;
- NIST AI RMF;
- MITRE ATT&CK;
- MITRE ATLAS;
- ISO/IEC 27001;
- ISO/IEC 42001;
- GDPR;
- EU AI Act;
- NIS2.

It stated that mappings are design references unless explicitly stated otherwise.

The Enterprise source additionally required traceability to:

- NIST SP 800-207;
- OWASP.

It stated that mappings support design and review and do not by themselves demonstrate legal compliance.

This consolidation preserves all eleven reference families and both claim limitations. It adds source versions, dates, legal-versus-voluntary classification, applicability, mapping strength, gaps, evidence, reviewer, update triggers, uncertainty, jurisdiction, validation and bidirectional traceability.

A mapping records a reasoned relationship. It does not prove that either side is implemented, effective, applicable or legally sufficient.


## 2. Fundamental Mapping Rules

- **COMP-REQ-001:** The project shall maintain mappings to NIST CSF.
- **COMP-REQ-002:** The project shall maintain mappings to NIST AI RMF.
- **COMP-REQ-003:** The project shall maintain mappings to NIST SP 800-207.
- **COMP-REQ-004:** The project shall maintain mappings to MITRE ATT&CK.
- **COMP-REQ-005:** The project shall maintain mappings to MITRE ATLAS.
- **COMP-REQ-006:** The project shall maintain mappings to OWASP.
- **COMP-REQ-007:** The project shall maintain mappings to ISO/IEC 27001.
- **COMP-REQ-008:** The project shall maintain mappings to ISO/IEC 42001.
- **COMP-REQ-009:** The project shall maintain mappings to GDPR.
- **COMP-REQ-010:** The project shall maintain mappings to the EU AI Act.
- **COMP-REQ-011:** The project shall maintain mappings to NIS2.
- **COMP-REQ-012:** Mappings shall support design and review.
- **COMP-REQ-013:** Mappings shall be treated as design references unless explicitly and validly classified otherwise.
- **COMP-REQ-014:** Mappings shall not by themselves demonstrate legal compliance.
- **COMP-REQ-015:** Mappings shall not by themselves demonstrate certification.
- **COMP-REQ-016:** Mappings shall not by themselves demonstrate implementation.
- **COMP-REQ-017:** Mappings shall not by themselves demonstrate control effectiveness.
- **COMP-REQ-018:** Mappings shall not by themselves demonstrate production readiness.
- **COMP-REQ-019:** Mappings shall not by themselves demonstrate institutional adoption.
- **COMP-REQ-020:** Mappings shall identify authoritative sources.
- **COMP-REQ-021:** Mappings shall identify source versions or dated snapshots.
- **COMP-REQ-022:** Mappings shall identify applicability.
- **COMP-REQ-023:** Mappings shall identify OBDIA requirements.
- **COMP-REQ-024:** Mappings shall identify external controls, outcomes, articles, objectives or techniques.
- **COMP-REQ-025:** Mappings shall identify mapping strength.
- **COMP-REQ-026:** Mappings shall identify gaps.
- **COMP-REQ-027:** Mappings shall identify evidence.
- **COMP-REQ-028:** Mappings shall identify reviewers.
- **COMP-REQ-029:** Mappings shall identify uncertainty.
- **COMP-REQ-030:** Mapping history shall not be silently rewritten.

## 3. Authority and Claim Boundaries

- **COMP-REQ-031:** This policy shall be normative for mapping mechanics and claim boundaries.
- **COMP-REQ-032:** Individual mapping records shall remain design and review evidence.
- **COMP-REQ-033:** A mapping record shall not override the Canonical Project Definition.
- **COMP-REQ-034:** A mapping record shall not override the Constitution.
- **COMP-REQ-035:** A mapping record shall not override an approved normative document.
- **COMP-REQ-036:** A mapping record shall not override an accepted ADR.
- **COMP-REQ-037:** A mapping record shall not create a legal conclusion.
- **COMP-REQ-038:** A mapping record shall not create a certification claim.
- **COMP-REQ-039:** A mapping record shall not create a conformity-assessment claim.
- **COMP-REQ-040:** A mapping record shall not create a regulatory-approval claim.
- **COMP-REQ-041:** A mapping record shall not create an audit opinion.
- **COMP-REQ-042:** A mapping record shall not create independent AI legal authority.
- **COMP-REQ-043:** A mapping record shall not authorize access.
- **COMP-REQ-044:** A mapping record shall not authorize evidence acquisition.
- **COMP-REQ-045:** A mapping record shall not accept risk.
- **COMP-REQ-046:** A mapping record shall not approve implementation.
- **COMP-REQ-047:** A mapping record shall not establish validation.
- **COMP-REQ-048:** A mapping record shall not publish a release.
- **COMP-REQ-049:** A mapping record shall not weaken canonical prohibitions.
- **COMP-REQ-050:** A mapping record shall not substitute for legal analysis.
- **COMP-REQ-051:** A mapping record shall not substitute for technical testing.
- **COMP-REQ-052:** A mapping record shall not substitute for implementation evidence.
- **COMP-REQ-053:** A mapping record shall not substitute for organizational context.
- **COMP-REQ-054:** Role concentration shall be disclosed.
- **COMP-REQ-055:** Internal review shall not be represented as independent external assurance.

## 4. Reference-Source Register

- **COMP-REQ-056:** Every external reference shall have an immutable source-record identifier.
- **COMP-REQ-057:** Every source record shall identify the reference family.
- **COMP-REQ-058:** Every source record shall identify the official title.
- **COMP-REQ-059:** Every source record shall identify the publisher or issuing authority.
- **COMP-REQ-060:** Every source record shall identify the edition, version or dated snapshot.
- **COMP-REQ-061:** Every source record shall identify publication or issue date.
- **COMP-REQ-062:** Every source record shall identify verification date.
- **COMP-REQ-063:** Every source record shall identify authoritative-source location.
- **COMP-REQ-064:** Every source record shall identify source type.
- **COMP-REQ-065:** Every source record shall identify legal or voluntary status.
- **COMP-REQ-066:** Every source record shall identify jurisdiction where applicable.
- **COMP-REQ-067:** Every source record shall identify amendment or corrigendum information where known.
- **COMP-REQ-068:** Every source record shall identify superseded versions.
- **COMP-REQ-069:** Every source record shall identify update mechanism.
- **COMP-REQ-070:** Every source record shall identify licensing or access limitations.
- **COMP-REQ-071:** Every source record shall identify machine-readable availability where known.
- **COMP-REQ-072:** Every source record shall identify scope.
- **COMP-REQ-073:** Every source record shall identify limitations.
- **COMP-REQ-074:** Every source record shall identify owner.
- **COMP-REQ-075:** Every source record shall identify review schedule.
- **COMP-REQ-076:** Unofficial summaries shall not replace authoritative sources.
- **COMP-REQ-077:** Search-engine snippets shall not become authoritative sources.
- **COMP-REQ-078:** Vendor crosswalks shall not replace independent source verification.
- **COMP-REQ-079:** Unavailable authoritative text shall remain an explicit limitation.
- **COMP-REQ-080:** Source-record history shall remain retrievable.

## 5. Source Classification

- **COMP-REQ-081:** Reference sources shall be classified by type.
- **COMP-REQ-082:** NIST framework publications shall be classified as voluntary guidance unless a separate authority makes them mandatory.
- **COMP-REQ-083:** NIST Special Publications shall be classified according to their source status and project use.
- **COMP-REQ-084:** MITRE ATT&CK shall be classified as a threat-knowledge reference.
- **COMP-REQ-085:** MITRE ATLAS shall be classified as an AI-threat-knowledge reference.
- **COMP-REQ-086:** OWASP resources shall be classified according to project type and release status.
- **COMP-REQ-087:** ISO standards shall be classified as voluntary consensus standards unless made applicable by contract, policy or law.
- **COMP-REQ-088:** ISO certification shall remain distinct from design alignment.
- **COMP-REQ-089:** EU regulations shall be classified as legal instruments.
- **COMP-REQ-090:** EU directives shall be classified as legal instruments requiring jurisdiction-specific implementation analysis.
- **COMP-REQ-091:** Legal applicability shall be assessed separately from technical relevance.
- **COMP-REQ-092:** Voluntary-reference relevance shall be assessed separately from legal applicability.
- **COMP-REQ-093:** Awareness documents shall remain distinct from verification standards.
- **COMP-REQ-094:** Threat taxonomies shall remain distinct from control frameworks.
- **COMP-REQ-095:** Management-system standards shall remain distinct from technical-control catalogs.
- **COMP-REQ-096:** Regulatory obligations shall remain distinct from implementation guidance.
- **COMP-REQ-097:** Source classification shall identify authority.
- **COMP-REQ-098:** Source classification shall identify uncertainty.
- **COMP-REQ-099:** Misclassified sources shall be corrected additively.
- **COMP-REQ-100:** Source classification shall not create compliance.

## 6. Authoritative-Source Verification

- **COMP-REQ-101:** Authoritative-source verification shall use official publishers or issuing authorities.
- **COMP-REQ-102:** Verification shall identify verifier.
- **COMP-REQ-103:** Verification shall identify date.
- **COMP-REQ-104:** Verification shall identify source version.
- **COMP-REQ-105:** Verification shall identify source status.
- **COMP-REQ-106:** Verification shall identify amendments.
- **COMP-REQ-107:** Verification shall identify corrigenda where material.
- **COMP-REQ-108:** Verification shall identify withdrawal or supersession.
- **COMP-REQ-109:** Verification shall identify source-language considerations.
- **COMP-REQ-110:** Verification shall identify consolidated-versus-original legal text where applicable.
- **COMP-REQ-111:** Verification shall distinguish authentic legal text from documentation tools.
- **COMP-REQ-112:** Verification shall distinguish stable releases from drafts.
- **COMP-REQ-113:** Verification shall distinguish current releases from archived releases.
- **COMP-REQ-114:** Verification shall preserve retrieved metadata.
- **COMP-REQ-115:** Verification shall preserve reference integrity where legally and technically possible.
- **COMP-REQ-116:** Verification shall not rely solely on third-party summaries.
- **COMP-REQ-117:** Verification shall not invent unavailable versions.
- **COMP-REQ-118:** Dynamic references shall use dated snapshots.
- **COMP-REQ-119:** Unknown source currency shall constrain mapping claims.
- **COMP-REQ-120:** Stale verification shall trigger review.
- **COMP-REQ-121:** Material source changes shall trigger remapping.
- **COMP-REQ-122:** Source-verification failure shall remain visible.
- **COMP-REQ-123:** Automated verification shall remain advisory.
- **COMP-REQ-124:** Human review shall confirm source authority.
- **COMP-REQ-125:** Verification history shall remain retrievable.

## 7. Mapping Record Identifiers

- **COMP-REQ-126:** Mapping records shall use `COMP-MAP-NNN` identifiers.
- **COMP-REQ-127:** Mapping-record identifiers shall be immutable.
- **COMP-REQ-128:** Mapping-record identifiers shall not be reused.
- **COMP-REQ-129:** Mapping-record identifiers shall not encode status.
- **COMP-REQ-130:** Mapping-record identifiers shall not encode strength.
- **COMP-REQ-131:** Mapping-record identifiers shall not encode applicability.
- **COMP-REQ-132:** Mapping-record identifiers shall not encode framework version.
- **COMP-REQ-133:** Mapping-record identifiers shall conform to `OBDIA-NAME-001`.
- **COMP-REQ-134:** Duplicate mapping identifiers shall be rejected.
- **COMP-REQ-135:** Superseded mappings shall retain their identifiers.
- **COMP-REQ-136:** Withdrawn mappings shall retain their identifiers.
- **COMP-REQ-137:** Archived mappings shall retain their identifiers.
- **COMP-REQ-138:** Source-record identifiers shall remain distinct from mapping identifiers.
- **COMP-REQ-139:** Gap identifiers shall remain distinct from mapping identifiers.
- **COMP-REQ-140:** Evidence identifiers shall remain distinct from mapping identifiers.
- **COMP-REQ-141:** Decision identifiers shall remain distinct from mapping identifiers.
- **COMP-REQ-142:** Risk identifiers shall remain distinct from mapping identifiers.
- **COMP-REQ-143:** Unknown identifier namespaces shall require review.
- **COMP-REQ-144:** Identifier corrections shall preserve history.
- **COMP-REQ-145:** Mapping indexes shall use immutable identifiers.

## 8. Mandatory Mapping Schema

- **COMP-REQ-146:** Every mapping shall include `Mapping ID`.
- **COMP-REQ-147:** Every mapping shall include `Title`.
- **COMP-REQ-148:** Every mapping shall include `Status`.
- **COMP-REQ-149:** Every mapping shall include `Owner`.
- **COMP-REQ-150:** Every mapping shall include `Reviewer`.
- **COMP-REQ-151:** Every mapping shall include `Created Date`.
- **COMP-REQ-152:** Every mapping shall include `Last Review Date`.
- **COMP-REQ-153:** Every active mapping shall include `Next Review Date or Trigger`.
- **COMP-REQ-154:** Every mapping shall include `Reference Family`.
- **COMP-REQ-155:** Every mapping shall include `Source Record ID`.
- **COMP-REQ-156:** Every mapping shall include `Source Version or Snapshot`.
- **COMP-REQ-157:** Every mapping shall include `Source Type`.
- **COMP-REQ-158:** Every mapping shall include `Legal or Voluntary Classification`.
- **COMP-REQ-159:** Every mapping shall include `Jurisdiction` where applicable.
- **COMP-REQ-160:** Every mapping shall include `Applicability`.
- **COMP-REQ-161:** Every mapping shall include `Applicability Rationale`.
- **COMP-REQ-162:** Every mapping shall include `OBDIA Artifact IDs`.
- **COMP-REQ-163:** Every mapping shall include `OBDIA Requirement IDs or Ranges`.
- **COMP-REQ-164:** Every mapping shall include `External Reference IDs or Objective`.
- **COMP-REQ-165:** Every mapping shall include `Mapping Relationship`.
- **COMP-REQ-166:** Every mapping shall include `Mapping Strength`.
- **COMP-REQ-167:** Every mapping shall include `Mapping Rationale`.
- **COMP-REQ-168:** Every mapping shall include `Supporting Evidence`.
- **COMP-REQ-169:** Every mapping shall include `Implementation State`.
- **COMP-REQ-170:** Every mapping shall include `Validation State`.
- **COMP-REQ-171:** Every mapping shall include `Gap Statement`.
- **COMP-REQ-172:** Every mapping shall include `Uncertainty`.
- **COMP-REQ-173:** Every mapping shall include `Assumption Links` where applicable.
- **COMP-REQ-174:** Every mapping shall include `Risk Links` where applicable.
- **COMP-REQ-175:** Every mapping shall include `Research Links` where applicable.
- **COMP-REQ-176:** Every mapping shall include `Decision Links` where applicable.
- **COMP-REQ-177:** Every mapping shall include `Test Evidence Links` where applicable.
- **COMP-REQ-178:** Every mapping shall include `Release Impact` where applicable.
- **COMP-REQ-179:** Every mapping shall include `Claim Boundary`.
- **COMP-REQ-180:** Every mapping shall include `Supersession` where applicable.
- **COMP-REQ-181:** Every mapping shall include `Correction History`.
- **COMP-REQ-182:** Every mapping shall identify exact OBDIA versions.
- **COMP-REQ-183:** Every mapping shall identify exact repository commit where material.
- **COMP-REQ-184:** Every mapping shall identify source currency.
- **COMP-REQ-185:** Incomplete mapping records shall not support authoritative claims.

## 9. Applicability Assessment

- **COMP-REQ-186:** Applicability shall be assessed explicitly.
- **COMP-REQ-187:** Permitted applicability states shall be `Applicable`, `Potentially Applicable`, `Contextual`, `Not Applicable` and `Not Assessed`.
- **COMP-REQ-188:** `Applicable` shall require documented facts supporting applicability.
- **COMP-REQ-189:** `Potentially Applicable` shall identify missing facts.
- **COMP-REQ-190:** `Contextual` shall mean useful design guidance without a legal-applicability conclusion.
- **COMP-REQ-191:** `Not Applicable` shall require rationale.
- **COMP-REQ-192:** `Not Assessed` shall not support compliance claims.
- **COMP-REQ-193:** Applicability shall identify organizational role.
- **COMP-REQ-194:** Applicability shall identify system role.
- **COMP-REQ-195:** Applicability shall identify processing role where relevant.
- **COMP-REQ-196:** Applicability shall identify jurisdiction where relevant.
- **COMP-REQ-197:** Applicability shall identify territorial scope where relevant.
- **COMP-REQ-198:** Applicability shall identify sector and entity criteria where relevant.
- **COMP-REQ-199:** Applicability shall identify deployment context.
- **COMP-REQ-200:** Applicability shall identify data categories.
- **COMP-REQ-201:** Applicability shall identify AI-system classification where relevant.
- **COMP-REQ-202:** Applicability shall identify contractual or policy obligations where relevant.
- **COMP-REQ-203:** Applicability shall not be inferred from repository location.
- **COMP-REQ-204:** Applicability shall not be inferred from framework popularity.
- **COMP-REQ-205:** Applicability shall not be inferred from technical similarity alone.
- **COMP-REQ-206:** Legal applicability shall receive qualified human review.
- **COMP-REQ-207:** Unknown facts shall remain explicit.
- **COMP-REQ-208:** Applicability changes shall be attributable.
- **COMP-REQ-209:** Applicability history shall remain retrievable.
- **COMP-REQ-210:** Application-specific legal advice remains outside this repository.

## 10. Mapping Relationship and Strength

- **COMP-REQ-211:** Permitted mapping relationships shall be `Direct`, `Partial`, `Supporting`, `Contextual`, `Conflict`, `Gap` and `Not Mapped`.
- **COMP-REQ-212:** `Direct` shall require substantially equivalent objective and scope.
- **COMP-REQ-213:** `Partial` shall identify covered and uncovered portions.
- **COMP-REQ-214:** `Supporting` shall mean the OBDIA requirement contributes without satisfying the external objective alone.
- **COMP-REQ-215:** `Contextual` shall mean conceptual relevance without control equivalence.
- **COMP-REQ-216:** `Conflict` shall identify incompatible expectations.
- **COMP-REQ-217:** `Gap` shall identify a missing or insufficient OBDIA control.
- **COMP-REQ-218:** `Not Mapped` shall remain explicit.
- **COMP-REQ-219:** Permitted mapping strengths shall be `Strong`, `Moderate`, `Weak` and `Unassessed`.
- **COMP-REQ-220:** `Strong` shall require clear objective, scope and evidence alignment.
- **COMP-REQ-221:** `Moderate` shall identify meaningful alignment with limitations.
- **COMP-REQ-222:** `Weak` shall identify indirect or incomplete alignment.
- **COMP-REQ-223:** `Unassessed` shall not support alignment claims.
- **COMP-REQ-224:** Relationship and strength shall be assessed separately.
- **COMP-REQ-225:** A Direct relationship shall not automatically be Strong.
- **COMP-REQ-226:** A Strong mapping shall not prove implementation.
- **COMP-REQ-227:** A Strong mapping shall not prove legal compliance.
- **COMP-REQ-228:** Multiple OBDIA requirements may map to one external objective.
- **COMP-REQ-229:** One OBDIA requirement may map to multiple external objectives.
- **COMP-REQ-230:** Many-to-many mappings shall preserve each relationship.
- **COMP-REQ-231:** Mapping rationale shall explain scope.
- **COMP-REQ-232:** Mapping rationale shall explain exclusions.
- **COMP-REQ-233:** Mapping rationale shall identify semantic differences.
- **COMP-REQ-234:** Mapping rationale shall identify evidence.
- **COMP-REQ-235:** Mapping rationale shall identify uncertainty.
- **COMP-REQ-236:** Mapping strength changes shall be attributable.
- **COMP-REQ-237:** Conflicting mappings shall remain visible.
- **COMP-REQ-238:** Selective omission shall be prohibited.
- **COMP-REQ-239:** Automated similarity scores shall remain advisory.
- **COMP-REQ-240:** Mapping history shall remain retrievable.

## 11. Evidence and Validation

- **COMP-REQ-241:** Every mapping shall identify supporting evidence.
- **COMP-REQ-242:** Evidence shall identify source.
- **COMP-REQ-243:** Evidence shall identify version.
- **COMP-REQ-244:** Evidence shall identify date.
- **COMP-REQ-245:** Evidence shall identify integrity where material.
- **COMP-REQ-246:** Evidence shall identify implementation state.
- **COMP-REQ-247:** Evidence shall identify validation state.
- **COMP-REQ-248:** Design documentation shall be labeled as design evidence.
- **COMP-REQ-249:** Implementation artifacts shall be labeled as implementation evidence.
- **COMP-REQ-250:** Test results shall be labeled as test evidence.
- **COMP-REQ-251:** Validation decisions shall be labeled as validation evidence.
- **COMP-REQ-252:** Audit or certification evidence shall not be claimed unless it exists and is authorized.
- **COMP-REQ-253:** Repository merge shall not count as control-effectiveness evidence.
- **COMP-REQ-254:** Passing syntax checks shall not count as semantic compliance evidence.
- **COMP-REQ-255:** Passing tests shall not generalize beyond tested scope.
- **COMP-REQ-256:** Evidence shall identify limitations.
- **COMP-REQ-257:** Evidence shall identify owner.
- **COMP-REQ-258:** Evidence shall identify reviewer.
- **COMP-REQ-259:** Evidence changes shall trigger mapping review.
- **COMP-REQ-260:** Invalidated evidence shall not support current mappings.
- **COMP-REQ-261:** Stale evidence shall constrain mapping strength.
- **COMP-REQ-262:** Missing evidence shall remain explicit.
- **COMP-REQ-263:** Evidence shall conform to `OBDIA-EVID-001` where applicable.
- **COMP-REQ-264:** Testing evidence shall conform to `OBDIA-TEST-001`.
- **COMP-REQ-265:** Evidence history shall remain retrievable.

## 12. Gap Management

- **COMP-REQ-266:** Every Partial, Gap or Conflict mapping shall identify a gap statement.
- **COMP-REQ-267:** Gap statements shall be precise.
- **COMP-REQ-268:** Gap statements shall identify affected external objectives.
- **COMP-REQ-269:** Gap statements shall identify affected OBDIA artifacts.
- **COMP-REQ-270:** Gap statements shall identify security impact.
- **COMP-REQ-271:** Gap statements shall identify privacy and rights impact.
- **COMP-REQ-272:** Gap statements shall identify evidence impact.
- **COMP-REQ-273:** Gap statements shall identify implementation impact.
- **COMP-REQ-274:** Gap statements shall identify risk.
- **COMP-REQ-275:** Gap statements shall identify owner.
- **COMP-REQ-276:** Gap statements shall identify priority.
- **COMP-REQ-277:** Gap statements shall identify treatment.
- **COMP-REQ-278:** Gap statements shall identify acceptance authority where applicable.
- **COMP-REQ-279:** Gap statements shall identify target date or review trigger.
- **COMP-REQ-280:** Gap closure shall identify evidence.
- **COMP-REQ-281:** Gap closure shall identify validation.
- **COMP-REQ-282:** Gap closure shall not be inferred from documentation change alone.
- **COMP-REQ-283:** Critical gaps shall block affected compliance claims.
- **COMP-REQ-284:** Critical safety or rights gaps shall block affected release.
- **COMP-REQ-285:** Gap acceptance shall not authorize canonical prohibitions.
- **COMP-REQ-286:** An AI system shall not accept gaps or residual risk.
- **COMP-REQ-287:** Gap changes shall be attributable.
- **COMP-REQ-288:** Gap corrections shall preserve history.
- **COMP-REQ-289:** Gap indexes shall remain current.
- **COMP-REQ-290:** Gap history shall remain retrievable.

## 13. Source Currency and Change Control

- **COMP-REQ-291:** Every source shall have a currency review.
- **COMP-REQ-292:** Dynamic references shall have dated snapshots.
- **COMP-REQ-293:** Source-version changes shall trigger impact assessment.
- **COMP-REQ-294:** Source amendments shall trigger impact assessment.
- **COMP-REQ-295:** Source corrigenda shall trigger impact assessment.
- **COMP-REQ-296:** Source withdrawal shall trigger impact assessment.
- **COMP-REQ-297:** Source supersession shall trigger impact assessment.
- **COMP-REQ-298:** Legal amendments shall trigger applicability review.
- **COMP-REQ-299:** National transposition changes shall trigger NIS2 applicability review.
- **COMP-REQ-300:** Framework taxonomy changes shall trigger mapping review.
- **COMP-REQ-301:** Technique deprecation shall trigger MITRE mapping review.
- **COMP-REQ-302:** OWASP project-version changes shall trigger profile review.
- **COMP-REQ-303:** ISO amendment changes shall trigger standard mapping review.
- **COMP-REQ-304:** OBDIA requirement changes shall trigger reverse mapping review.
- **COMP-REQ-305:** OBDIA version changes shall update mapping metadata.
- **COMP-REQ-306:** Mapping changes shall use semantic versioning.
- **COMP-REQ-307:** Material mapping changes shall use change control.
- **COMP-REQ-308:** Breaking mapping-schema changes shall require migration.
- **COMP-REQ-309:** Historical mappings shall not be silently rewritten.
- **COMP-REQ-310:** Superseded mappings shall remain retrievable.
- **COMP-REQ-311:** Unknown source currency shall block strong-current claims.
- **COMP-REQ-312:** Source drift shall remain visible.
- **COMP-REQ-313:** Change-control evidence shall identify exact commits.
- **COMP-REQ-314:** Rollback shall not restore known-invalid mappings.
- **COMP-REQ-315:** Source-change history shall remain retrievable.

## 14. Roles and Review

- **COMP-REQ-316:** The Privacy and Governance Reviewer shall own mapping governance.
- **COMP-REQ-317:** The Security Reviewer shall assess cybersecurity and threat mappings.
- **COMP-REQ-318:** The Implementation Reviewer shall assess implementation and test evidence.
- **COMP-REQ-319:** The Research Reviewer shall assess source quality and uncertainty.
- **COMP-REQ-320:** The Release Reviewer shall assess public mapping claims.
- **COMP-REQ-321:** The Documentation Authority shall assess metadata and traceability.
- **COMP-REQ-322:** The Project Founder shall approve normative baseline status.
- **COMP-REQ-323:** Legal applicability shall receive qualified human review where required.
- **COMP-REQ-324:** AI systems shall not approve mappings.
- **COMP-REQ-325:** AI systems shall not make legal applicability determinations.
- **COMP-REQ-326:** AI systems shall not accept mapping gaps or risks.
- **COMP-REQ-327:** Automated crosswalk generation shall remain advisory.
- **COMP-REQ-328:** Review shall identify exact mapping versions.
- **COMP-REQ-329:** Review shall identify exact source versions.
- **COMP-REQ-330:** Review shall identify exact OBDIA versions.
- **COMP-REQ-331:** Review shall identify findings.
- **COMP-REQ-332:** Review shall identify dispositions.
- **COMP-REQ-333:** Blocking findings shall prevent affected claims.
- **COMP-REQ-334:** Role concentration shall be disclosed.
- **COMP-REQ-335:** Review history shall remain retrievable.

## 15. NIST Cybersecurity Framework Mapping

- **COMP-REQ-336:** NIST CSF mappings shall identify the CSF version.
- **COMP-REQ-337:** The Draft baseline shall use NIST CSF 2.0 as the verified source version.
- **COMP-REQ-338:** NIST CSF mappings shall identify Functions.
- **COMP-REQ-339:** NIST CSF mappings shall identify Categories or Subcategories where detailed mapping is performed.
- **COMP-REQ-340:** NIST CSF mappings shall distinguish outcomes from implementation methods.
- **COMP-REQ-341:** NIST CSF mappings shall not treat the Framework as prescriptive implementation.
- **COMP-REQ-342:** NIST CSF Governance mappings shall include project governance and risk ownership.
- **COMP-REQ-343:** NIST CSF Identify mappings shall include assets, threats, risks and dependencies.
- **COMP-REQ-344:** NIST CSF Protect mappings shall include identity, authorization, data and secure development.
- **COMP-REQ-345:** NIST CSF Detect mappings shall include monitoring, logging and anomaly detection.
- **COMP-REQ-346:** NIST CSF Respond mappings shall include containment, communication and analysis.
- **COMP-REQ-347:** NIST CSF Recover mappings shall include restoration, lessons learned and resilience.
- **COMP-REQ-348:** NIST CSF Profiles shall not be claimed unless created and governed.
- **COMP-REQ-349:** NIST CSF Tier claims shall not be made without evidence and context.
- **COMP-REQ-350:** CSF mappings shall identify implementation gaps.
- **COMP-REQ-351:** CSF mappings shall identify validation gaps.
- **COMP-REQ-352:** CSF mappings shall identify risk-register links.
- **COMP-REQ-353:** CSF mappings shall receive Security Reviewer assessment.
- **COMP-REQ-354:** CSF mappings shall not imply NIST endorsement.
- **COMP-REQ-355:** CSF mapping history shall remain retrievable.

## 16. NIST AI RMF Mapping

- **COMP-REQ-356:** NIST AI RMF mappings shall identify the AI RMF version.
- **COMP-REQ-357:** The Draft baseline shall use NIST AI RMF 1.0 as the verified published source version.
- **COMP-REQ-358:** AI RMF mappings shall identify Govern.
- **COMP-REQ-359:** AI RMF mappings shall identify Map.
- **COMP-REQ-360:** AI RMF mappings shall identify Measure.
- **COMP-REQ-361:** AI RMF mappings shall identify Manage.
- **COMP-REQ-362:** AI RMF mappings shall distinguish voluntary guidance from legal obligations.
- **COMP-REQ-363:** AI RMF mappings shall include accountability.
- **COMP-REQ-364:** AI RMF mappings shall include context and impact.
- **COMP-REQ-365:** AI RMF mappings shall include measurement limitations.
- **COMP-REQ-366:** AI RMF mappings shall include treatment and monitoring.
- **COMP-REQ-367:** AI RMF mappings shall include human oversight.
- **COMP-REQ-368:** AI RMF mappings shall include privacy and rights.
- **COMP-REQ-369:** AI RMF mappings shall include transparency and explainability.
- **COMP-REQ-370:** AI RMF mappings shall include validity, reliability, safety and security where applicable.
- **COMP-REQ-371:** AI RMF revision activity shall trigger source review.
- **COMP-REQ-372:** AI RMF profile mappings shall identify the exact profile.
- **COMP-REQ-373:** AI RMF mappings shall receive Privacy and Governance Reviewer assessment.
- **COMP-REQ-374:** AI RMF mappings shall not imply NIST endorsement.
- **COMP-REQ-375:** AI RMF mapping history shall remain retrievable.

## 17. NIST SP 800-207 Mapping

- **COMP-REQ-376:** NIST SP 800-207 mappings shall identify the publication version and date.
- **COMP-REQ-377:** The Draft baseline shall use the August 2020 final publication.
- **COMP-REQ-378:** Zero Trust mappings shall reject implicit trust from network location.
- **COMP-REQ-379:** Zero Trust mappings shall focus on resources.
- **COMP-REQ-380:** Zero Trust mappings shall distinguish authentication and authorization.
- **COMP-REQ-381:** Zero Trust mappings shall include subject and device or workload context where applicable.
- **COMP-REQ-382:** Zero Trust mappings shall include policy decision functions.
- **COMP-REQ-383:** Zero Trust mappings shall include policy administration functions.
- **COMP-REQ-384:** Zero Trust mappings shall include policy enforcement functions.
- **COMP-REQ-385:** Zero Trust mappings shall include continuous or repeated evaluation.
- **COMP-REQ-386:** Zero Trust mappings shall include least privilege.
- **COMP-REQ-387:** Zero Trust mappings shall include session and validity constraints.
- **COMP-REQ-388:** Zero Trust mappings shall include telemetry.
- **COMP-REQ-389:** Zero Trust mappings shall include dynamic policy inputs.
- **COMP-REQ-390:** Zero Trust mappings shall include failure behavior.
- **COMP-REQ-391:** Zero Trust mappings shall not imply that encryption creates trust.
- **COMP-REQ-392:** Zero Trust mappings shall not imply that network segmentation alone satisfies Zero Trust.
- **COMP-REQ-393:** Zero Trust mappings shall receive Security Reviewer assessment.
- **COMP-REQ-394:** Zero Trust mappings shall not imply NIST endorsement.
- **COMP-REQ-395:** Zero Trust mapping history shall remain retrievable.

## 18. MITRE ATT&CK Mapping

- **COMP-REQ-396:** MITRE ATT&CK mappings shall identify the ATT&CK version.
- **COMP-REQ-397:** The Draft baseline shall use ATT&CK v19.1 as the verified current snapshot.
- **COMP-REQ-398:** ATT&CK mappings shall identify applicable domain.
- **COMP-REQ-399:** ATT&CK mappings shall identify tactic or technique identifiers where detailed mapping is performed.
- **COMP-REQ-400:** ATT&CK mappings shall distinguish threats from controls.
- **COMP-REQ-401:** ATT&CK mappings shall distinguish observed behavior from hypothetical behavior.
- **COMP-REQ-402:** ATT&CK mappings shall distinguish technique coverage from prevention claims.
- **COMP-REQ-403:** ATT&CK mappings shall include relevant credential-access threats.
- **COMP-REQ-404:** ATT&CK mappings shall include relevant defense-evasion threats.
- **COMP-REQ-405:** ATT&CK mappings shall include relevant collection and exfiltration threats.
- **COMP-REQ-406:** ATT&CK mappings shall include relevant persistence and privilege-escalation threats.
- **COMP-REQ-407:** ATT&CK mappings shall include relevant supply-chain threats.
- **COMP-REQ-408:** ATT&CK mappings shall include detection strategies where applicable.
- **COMP-REQ-409:** ATT&CK mappings shall identify mitigations separately from techniques.
- **COMP-REQ-410:** ATT&CK deprecations and revocations shall trigger review.
- **COMP-REQ-411:** ATT&CK mappings shall use defensive and authorized research methods.
- **COMP-REQ-412:** ATT&CK mappings shall not authorize adversary emulation against unauthorized targets.
- **COMP-REQ-413:** ATT&CK mappings shall receive Security Reviewer assessment.
- **COMP-REQ-414:** ATT&CK mappings shall not imply MITRE endorsement.
- **COMP-REQ-415:** ATT&CK mapping history shall remain retrievable.

## 19. MITRE ATLAS Mapping

- **COMP-REQ-416:** MITRE ATLAS mappings shall use a dated authoritative snapshot.
- **COMP-REQ-417:** The Draft baseline shall use the official ATLAS snapshot verified on 2026-08-06.
- **COMP-REQ-418:** ATLAS mappings shall identify tactics or technique identifiers where detailed mapping is performed.
- **COMP-REQ-419:** ATLAS mappings shall distinguish feasible, demonstrated and realized maturity where used.
- **COMP-REQ-420:** ATLAS mappings shall include prompt-injection threats.
- **COMP-REQ-421:** ATLAS mappings shall include retrieval and context poisoning.
- **COMP-REQ-422:** ATLAS mappings shall include model and data poisoning.
- **COMP-REQ-423:** ATLAS mappings shall include AI supply-chain compromise.
- **COMP-REQ-424:** ATLAS mappings shall include agent-tool abuse.
- **COMP-REQ-425:** ATLAS mappings shall include credential harvesting through AI components.
- **COMP-REQ-426:** ATLAS mappings shall include model or prompt extraction.
- **COMP-REQ-427:** ATLAS mappings shall include data leakage and exfiltration.
- **COMP-REQ-428:** ATLAS mappings shall include excessive or manipulated agency.
- **COMP-REQ-429:** ATLAS mappings shall identify mitigations separately from techniques.
- **COMP-REQ-430:** ATLAS mappings shall identify evidence maturity.
- **COMP-REQ-431:** ATLAS mappings shall use defensive controlled testing.
- **COMP-REQ-432:** ATLAS mappings shall not authorize harmful model exploitation.
- **COMP-REQ-433:** ATLAS mappings shall receive Security Reviewer assessment.
- **COMP-REQ-434:** ATLAS mappings shall not imply MITRE endorsement.
- **COMP-REQ-435:** ATLAS mapping history shall remain retrievable.

## 20. OWASP Mapping

- **COMP-REQ-436:** OWASP mappings shall identify the exact OWASP project.
- **COMP-REQ-437:** OWASP mappings shall identify the exact project version.
- **COMP-REQ-438:** OWASP awareness documents shall remain distinct from verification standards.
- **COMP-REQ-439:** The Draft application-security profile shall include OWASP ASVS 5.0.0.
- **COMP-REQ-440:** The Draft awareness profile shall include OWASP Top 10:2025.
- **COMP-REQ-441:** The Draft AI-security profile shall include OWASP AISVS 1.0.
- **COMP-REQ-442:** The Draft LLM-security profile shall include OWASP LLMSVS 2.0.
- **COMP-REQ-443:** OWASP ASVS mappings shall identify requirement IDs with versions where detailed mapping is performed.
- **COMP-REQ-444:** OWASP Top 10 mappings shall be classified as awareness mappings.
- **COMP-REQ-445:** OWASP AISVS mappings shall be classified according to its verification requirements.
- **COMP-REQ-446:** OWASP LLMSVS mappings shall be classified according to its verification requirements.
- **COMP-REQ-447:** OWASP mappings shall include access control.
- **COMP-REQ-448:** OWASP mappings shall include authentication.
- **COMP-REQ-449:** OWASP mappings shall include secure configuration.
- **COMP-REQ-450:** OWASP mappings shall include cryptography.
- **COMP-REQ-451:** OWASP mappings shall include injection and output handling.
- **COMP-REQ-452:** OWASP mappings shall include software supply chain.
- **COMP-REQ-453:** OWASP mappings shall include logging and exceptional conditions.
- **COMP-REQ-454:** OWASP mappings shall include AI model lifecycle.
- **COMP-REQ-455:** OWASP mappings shall include LLM integration.
- **COMP-REQ-456:** OWASP mappings shall include agents and plugins or tools.
- **COMP-REQ-457:** OWASP mappings shall include model memory and storage.
- **COMP-REQ-458:** OWASP mappings shall include monitoring and anomaly detection.
- **COMP-REQ-459:** OWASP mappings shall include prompt injection and excessive agency.
- **COMP-REQ-460:** OWASP project drafts shall not be represented as stable releases.
- **COMP-REQ-461:** OWASP mappings shall identify project-specific limitations.
- **COMP-REQ-462:** OWASP mappings shall receive Security and Implementation Reviewer assessment.
- **COMP-REQ-463:** OWASP mappings shall not imply OWASP certification or endorsement.
- **COMP-REQ-464:** OWASP version changes shall trigger review.
- **COMP-REQ-465:** OWASP mapping history shall remain retrievable.

## 21. ISO/IEC 27001 Mapping

- **COMP-REQ-466:** ISO/IEC 27001 mappings shall identify edition and amendments.
- **COMP-REQ-467:** The Draft baseline shall use ISO/IEC 27001:2022.
- **COMP-REQ-468:** The Draft baseline shall record ISO/IEC 27001:2022/Amd 1:2024.
- **COMP-REQ-469:** ISO/IEC 27001 mappings shall distinguish requirements from implementation guidance.
- **COMP-REQ-470:** ISO/IEC 27001 mappings shall identify information-security management-system context.
- **COMP-REQ-471:** ISO/IEC 27001 mappings shall identify leadership and governance.
- **COMP-REQ-472:** ISO/IEC 27001 mappings shall identify planning and risk treatment.
- **COMP-REQ-473:** ISO/IEC 27001 mappings shall identify support and competence.
- **COMP-REQ-474:** ISO/IEC 27001 mappings shall identify operational control.
- **COMP-REQ-475:** ISO/IEC 27001 mappings shall identify performance evaluation.
- **COMP-REQ-476:** ISO/IEC 27001 mappings shall identify continual improvement.
- **COMP-REQ-477:** Annex control mappings shall identify exact control references where licensed access permits.
- **COMP-REQ-478:** ISO mappings shall respect copyright and access constraints.
- **COMP-REQ-479:** Design alignment shall not be called conformity.
- **COMP-REQ-480:** Conformity shall not be called certification.
- **COMP-REQ-481:** Certification shall not be claimed without an authorized certification process.
- **COMP-REQ-482:** ISO mappings shall identify implementation and evidence gaps.
- **COMP-REQ-483:** ISO mappings shall receive Security and Governance Reviewer assessment.
- **COMP-REQ-484:** ISO mappings shall not imply ISO endorsement.
- **COMP-REQ-485:** ISO/IEC 27001 mapping history shall remain retrievable.

## 22. ISO/IEC 42001 Mapping

- **COMP-REQ-486:** ISO/IEC 42001 mappings shall identify edition.
- **COMP-REQ-487:** The Draft baseline shall use ISO/IEC 42001:2023.
- **COMP-REQ-488:** ISO/IEC 42001 mappings shall distinguish management-system requirements from technical controls.
- **COMP-REQ-489:** ISO/IEC 42001 mappings shall identify organizational context.
- **COMP-REQ-490:** ISO/IEC 42001 mappings shall identify leadership and policy.
- **COMP-REQ-491:** ISO/IEC 42001 mappings shall identify AI objectives and planning.
- **COMP-REQ-492:** ISO/IEC 42001 mappings shall identify resources and competence.
- **COMP-REQ-493:** ISO/IEC 42001 mappings shall identify AI-system lifecycle processes.
- **COMP-REQ-494:** ISO/IEC 42001 mappings shall identify risk and opportunity treatment.
- **COMP-REQ-495:** ISO/IEC 42001 mappings shall identify AI impact assessment where applicable.
- **COMP-REQ-496:** ISO/IEC 42001 mappings shall identify data and model governance.
- **COMP-REQ-497:** ISO/IEC 42001 mappings shall identify monitoring and measurement.
- **COMP-REQ-498:** ISO/IEC 42001 mappings shall identify internal review and improvement.
- **COMP-REQ-499:** ISO mappings shall respect copyright and access constraints.
- **COMP-REQ-500:** Design alignment shall not be called conformity.
- **COMP-REQ-501:** Conformity shall not be called certification.
- **COMP-REQ-502:** ISO mappings shall identify implementation and evidence gaps.
- **COMP-REQ-503:** ISO mappings shall receive AI Governance and Privacy Reviewer assessment.
- **COMP-REQ-504:** ISO mappings shall not imply ISO endorsement.
- **COMP-REQ-505:** ISO/IEC 42001 mapping history shall remain retrievable.

## 23. GDPR Mapping

- **COMP-REQ-506:** GDPR mappings shall identify Regulation (EU) 2016/679.
- **COMP-REQ-507:** GDPR mappings shall identify the authoritative or consolidated source used.
- **COMP-REQ-508:** GDPR applicability shall be assessed from facts and jurisdiction.
- **COMP-REQ-509:** GDPR mappings shall identify controller, processor or other relevant role only when supported.
- **COMP-REQ-510:** GDPR mappings shall identify personal-data categories.
- **COMP-REQ-511:** GDPR mappings shall identify processing purposes.
- **COMP-REQ-512:** GDPR mappings shall identify lawful-basis uncertainty rather than inventing a basis.
- **COMP-REQ-513:** GDPR mappings shall include data-protection principles.
- **COMP-REQ-514:** GDPR mappings shall include purpose limitation.
- **COMP-REQ-515:** GDPR mappings shall include data minimization.
- **COMP-REQ-516:** GDPR mappings shall include accuracy and correction.
- **COMP-REQ-517:** GDPR mappings shall include storage limitation.
- **COMP-REQ-518:** GDPR mappings shall include integrity and confidentiality.
- **COMP-REQ-519:** GDPR mappings shall include accountability.
- **COMP-REQ-520:** GDPR mappings shall include data-protection by design and by default.
- **COMP-REQ-521:** GDPR mappings shall include data-subject rights where applicable.
- **COMP-REQ-522:** GDPR mappings shall include security of processing.
- **COMP-REQ-523:** GDPR mappings shall include breach handling where applicable.
- **COMP-REQ-524:** GDPR mappings shall include impact-assessment triggers where applicable.
- **COMP-REQ-525:** GDPR mappings shall include international-transfer issues where applicable.
- **COMP-REQ-526:** Technical controls shall not be represented as legal compliance by themselves.
- **COMP-REQ-527:** Legal applicability shall receive qualified human review.
- **COMP-REQ-528:** GDPR mappings shall identify national-law dependencies where relevant.
- **COMP-REQ-529:** GDPR public claims shall remain qualified.
- **COMP-REQ-530:** GDPR mapping history shall remain retrievable.

## 24. EU AI Act Mapping

- **COMP-REQ-531:** EU AI Act mappings shall identify Regulation (EU) 2024/1689.
- **COMP-REQ-532:** EU AI Act mappings shall identify the authoritative source used.
- **COMP-REQ-533:** EU AI Act applicability shall be assessed from facts, roles and jurisdiction.
- **COMP-REQ-534:** AI-system classification shall not be inferred from project terminology alone.
- **COMP-REQ-535:** Prohibited-practice analysis shall receive qualified human review.
- **COMP-REQ-536:** High-risk classification shall not be claimed or rejected without factual analysis.
- **COMP-REQ-537:** Provider, deployer, importer and distributor roles shall not be assigned without evidence.
- **COMP-REQ-538:** EU AI Act mappings shall identify human-oversight relevance.
- **COMP-REQ-539:** EU AI Act mappings shall identify risk-management relevance.
- **COMP-REQ-540:** EU AI Act mappings shall identify data-governance relevance.
- **COMP-REQ-541:** EU AI Act mappings shall identify technical-documentation relevance.
- **COMP-REQ-542:** EU AI Act mappings shall identify record-keeping and logging relevance.
- **COMP-REQ-543:** EU AI Act mappings shall identify transparency relevance.
- **COMP-REQ-544:** EU AI Act mappings shall identify accuracy, robustness and cybersecurity relevance.
- **COMP-REQ-545:** EU AI Act mappings shall identify post-market and incident obligations where applicable.
- **COMP-REQ-546:** EU AI Act mappings shall identify general-purpose AI relevance only when supported.
- **COMP-REQ-547:** Application dates and transitional provisions shall be assessed separately.
- **COMP-REQ-548:** Technical design alignment shall not be represented as regulatory compliance.
- **COMP-REQ-549:** Fundamental-rights impacts shall remain explicit.
- **COMP-REQ-550:** National competent-authority or market-surveillance conclusions shall not be inferred.
- **COMP-REQ-551:** Legal applicability shall receive qualified human review.
- **COMP-REQ-552:** EU AI Act mappings shall identify evidence gaps.
- **COMP-REQ-553:** EU AI Act public claims shall remain qualified.
- **COMP-REQ-554:** EU AI Act source changes shall trigger review.
- **COMP-REQ-555:** EU AI Act mapping history shall remain retrievable.

## 25. NIS2 Mapping

- **COMP-REQ-556:** NIS2 mappings shall identify Directive (EU) 2022/2555.
- **COMP-REQ-557:** NIS2 mappings shall identify the authoritative or consolidated source used.
- **COMP-REQ-558:** NIS2 applicability shall consider national transposition.
- **COMP-REQ-559:** NIS2 applicability shall consider entity type and sector.
- **COMP-REQ-560:** NIS2 applicability shall consider size and jurisdiction criteria where relevant.
- **COMP-REQ-561:** Entity classification shall not be inferred from repository activity.
- **COMP-REQ-562:** NIS2 mappings shall include management-body accountability relevance.
- **COMP-REQ-563:** NIS2 mappings shall include cybersecurity risk-management relevance.
- **COMP-REQ-564:** NIS2 mappings shall include incident-handling relevance.
- **COMP-REQ-565:** NIS2 mappings shall include business-continuity relevance.
- **COMP-REQ-566:** NIS2 mappings shall include supply-chain security relevance.
- **COMP-REQ-567:** NIS2 mappings shall include vulnerability handling and disclosure relevance.
- **COMP-REQ-568:** NIS2 mappings shall include cryptography and access-control relevance.
- **COMP-REQ-569:** NIS2 mappings shall include security training relevance.
- **COMP-REQ-570:** NIS2 mappings shall include incident-reporting relevance where applicable.
- **COMP-REQ-571:** National reporting processes shall be assessed separately.
- **COMP-REQ-572:** Technical design alignment shall not be represented as legal compliance.
- **COMP-REQ-573:** Directive mapping shall not replace national-law analysis.
- **COMP-REQ-574:** Legal applicability shall receive qualified human review.
- **COMP-REQ-575:** NIS2 mappings shall identify evidence gaps.
- **COMP-REQ-576:** NIS2 mappings shall identify implementation gaps.
- **COMP-REQ-577:** NIS2 public claims shall remain qualified.
- **COMP-REQ-578:** National-law changes shall trigger review.
- **COMP-REQ-579:** NIS2 source changes shall trigger review.
- **COMP-REQ-580:** NIS2 mapping history shall remain retrievable.

## 26. Cross-Framework Harmonization

- **COMP-REQ-581:** Cross-framework mappings shall preserve source-specific meaning.
- **COMP-REQ-582:** Cross-framework mappings shall not flatten legal and voluntary references into one category.
- **COMP-REQ-583:** Cross-framework mappings shall distinguish governance outcomes.
- **COMP-REQ-584:** Cross-framework mappings shall distinguish technical controls.
- **COMP-REQ-585:** Cross-framework mappings shall distinguish threat knowledge.
- **COMP-REQ-586:** Cross-framework mappings shall distinguish awareness guidance.
- **COMP-REQ-587:** Cross-framework mappings shall distinguish verification requirements.
- **COMP-REQ-588:** Cross-framework mappings shall distinguish management-system requirements.
- **COMP-REQ-589:** Cross-framework mappings shall distinguish legal obligations.
- **COMP-REQ-590:** Common terminology shall use the OBDIA glossary.
- **COMP-REQ-591:** Source-specific terminology shall be retained where necessary.
- **COMP-REQ-592:** Equivalent terms shall not be assumed without analysis.
- **COMP-REQ-593:** Conflicts shall remain visible.
- **COMP-REQ-594:** Gaps shall remain visible.
- **COMP-REQ-595:** One strong mapping shall not satisfy all related frameworks.
- **COMP-REQ-596:** A legal obligation shall not be reduced to a voluntary-control analogy.
- **COMP-REQ-597:** A threat technique shall not be represented as a requirement.
- **COMP-REQ-598:** A control objective shall not be represented as implementation evidence.
- **COMP-REQ-599:** A management-system clause shall not be represented as a technical test.
- **COMP-REQ-600:** Crosswalks shall identify directionality.
- **COMP-REQ-601:** Crosswalks shall identify many-to-many relationships.
- **COMP-REQ-602:** Crosswalks shall identify scope limitations.
- **COMP-REQ-603:** Crosswalks shall identify source versions.
- **COMP-REQ-604:** Crosswalks shall receive multidisciplinary review.
- **COMP-REQ-605:** Cross-framework history shall remain retrievable.

## 27. OBDIA Requirement Mapping

- **COMP-REQ-606:** Mappings shall use immutable OBDIA Document IDs.
- **COMP-REQ-607:** Mappings shall use stable OBDIA requirement IDs.
- **COMP-REQ-608:** Mappings shall identify OBDIA document versions.
- **COMP-REQ-609:** Mappings shall identify OBDIA lifecycle statuses.
- **COMP-REQ-610:** Mappings shall identify normative authority.
- **COMP-REQ-611:** Mappings shall distinguish normative requirements from research records.
- **COMP-REQ-612:** Mappings shall distinguish design requirements from implementation evidence.
- **COMP-REQ-613:** Mappings shall distinguish implementation from validation.
- **COMP-REQ-614:** Mappings shall identify affected architecture domains.
- **COMP-REQ-615:** Mappings shall identify affected trust boundaries.
- **COMP-REQ-616:** Mappings shall identify identity and authorization controls.
- **COMP-REQ-617:** Mappings shall identify evidence and provenance controls.
- **COMP-REQ-618:** Mappings shall identify connector controls.
- **COMP-REQ-619:** Mappings shall identify AI and human-oversight controls.
- **COMP-REQ-620:** Mappings shall identify privacy and rights controls.
- **COMP-REQ-621:** Mappings shall identify secure-development controls.
- **COMP-REQ-622:** Mappings shall identify testing controls.
- **COMP-REQ-623:** Mappings shall identify repository and release controls.
- **COMP-REQ-624:** Mappings shall identify risk, research, assumption and decision records.
- **COMP-REQ-625:** Requirement changes shall trigger reverse mapping review.
- **COMP-REQ-626:** Requirement supersession shall update mappings.
- **COMP-REQ-627:** Requirement deletion shall preserve historical disposition.
- **COMP-REQ-628:** Unmapped material requirements shall remain visible.
- **COMP-REQ-629:** Mapping coverage shall not be represented as control effectiveness.
- **COMP-REQ-630:** OBDIA mapping history shall remain retrievable.

## 28. Design, Implementation and Validation Separation

- **COMP-REQ-631:** Mapping state shall distinguish design alignment.
- **COMP-REQ-632:** Mapping state shall distinguish planned implementation.
- **COMP-REQ-633:** Mapping state shall distinguish implemented controls.
- **COMP-REQ-634:** Mapping state shall distinguish tested controls.
- **COMP-REQ-635:** Mapping state shall distinguish validated controls.
- **COMP-REQ-636:** Mapping state shall distinguish published claims.
- **COMP-REQ-637:** Design alignment shall require normative documentation.
- **COMP-REQ-638:** Implementation claims shall require implementation evidence.
- **COMP-REQ-639:** Testing claims shall require test evidence.
- **COMP-REQ-640:** Validation claims shall require validation decisions.
- **COMP-REQ-641:** Publication claims shall require release evidence.
- **COMP-REQ-642:** Approved documents shall not imply implemented controls.
- **COMP-REQ-643:** Implemented controls shall not imply validated effectiveness.
- **COMP-REQ-644:** Validated controls shall not imply legal compliance.
- **COMP-REQ-645:** Legal compliance shall not be claimed through a technical crosswalk.
- **COMP-REQ-646:** Partially implemented controls shall identify missing portions.
- **COMP-REQ-647:** Partially validated controls shall identify unvalidated scope.
- **COMP-REQ-648:** Invalidated evidence shall update mappings.
- **COMP-REQ-649:** Implementation drift shall trigger review.
- **COMP-REQ-650:** Environment differences shall constrain validation.
- **COMP-REQ-651:** Provider changes shall constrain validation.
- **COMP-REQ-652:** Model changes shall constrain validation.
- **COMP-REQ-653:** Mapping status shall not be inferred from PR state.
- **COMP-REQ-654:** Mapping status shall not be inferred from release tags.
- **COMP-REQ-655:** State-separation history shall remain retrievable.

## 29. Auditability and Evidence Preservation

- **COMP-REQ-656:** Mapping changes shall be attributable.
- **COMP-REQ-657:** Mapping reviews shall be attributable.
- **COMP-REQ-658:** Mapping approvals shall identify exact versions.
- **COMP-REQ-659:** Mapping evidence shall identify exact commits.
- **COMP-REQ-660:** Mapping exports shall have integrity references.
- **COMP-REQ-661:** Mapping indexes shall have integrity references where published.
- **COMP-REQ-662:** Mapping source records shall retain verification evidence.
- **COMP-REQ-663:** Mapping gaps shall retain dispositions.
- **COMP-REQ-664:** Mapping conflicts shall retain resolutions.
- **COMP-REQ-665:** Mapping corrections shall be additive.
- **COMP-REQ-666:** Mapping supersession shall preserve predecessors.
- **COMP-REQ-667:** Mapping withdrawal shall preserve history.
- **COMP-REQ-668:** Audit logs shall not contain secrets.
- **COMP-REQ-669:** Audit logs shall minimize personal data.
- **COMP-REQ-670:** Automated mapping actions shall use governed machine identities.
- **COMP-REQ-671:** Automated mapping actions shall not approve results.
- **COMP-REQ-672:** Evidence retention shall follow classification.
- **COMP-REQ-673:** Archive migration shall preserve identifiers.
- **COMP-REQ-674:** Archive restoration shall not reactivate mappings silently.
- **COMP-REQ-675:** Audit history shall remain retrievable.

## 30. Legal, Regulatory and Rights Claim Limits

- **COMP-REQ-676:** The repository shall not claim GDPR compliance from mapping alone.
- **COMP-REQ-677:** The repository shall not claim EU AI Act compliance from mapping alone.
- **COMP-REQ-678:** The repository shall not claim NIS2 compliance from mapping alone.
- **COMP-REQ-679:** The repository shall not claim ISO conformity from mapping alone.
- **COMP-REQ-680:** The repository shall not claim ISO certification from mapping alone.
- **COMP-REQ-681:** The repository shall not claim NIST certification.
- **COMP-REQ-682:** The repository shall not claim MITRE validation or endorsement.
- **COMP-REQ-683:** The repository shall not claim OWASP certification or endorsement without a valid applicable program.
- **COMP-REQ-684:** Mappings shall not replace qualified legal advice.
- **COMP-REQ-685:** Mappings shall not determine lawful basis autonomously.
- **COMP-REQ-686:** Mappings shall not determine AI-system classification autonomously.
- **COMP-REQ-687:** Mappings shall not determine NIS2 entity classification autonomously.
- **COMP-REQ-688:** Mappings shall not determine jurisdiction autonomously.
- **COMP-REQ-689:** Mappings shall not determine admissibility autonomously.
- **COMP-REQ-690:** Mappings shall not authorize processing autonomously.
- **COMP-REQ-691:** Mappings shall not authorize surveillance.
- **COMP-REQ-692:** Mappings shall not authorize deanonymization.
- **COMP-REQ-693:** Mappings shall not authorize criminal-infrastructure interaction.
- **COMP-REQ-694:** Mappings shall not create public-authority powers.
- **COMP-REQ-695:** Mappings shall preserve human accountability.
- **COMP-REQ-696:** Fundamental-rights uncertainty shall remain visible.
- **COMP-REQ-697:** Legal-source changes shall trigger review.
- **COMP-REQ-698:** Public disclaimers shall remain adjacent to mapping claims.
- **COMP-REQ-699:** Misleading claims shall require correction or withdrawal.
- **COMP-REQ-700:** Claim-limit history shall remain retrievable.

## 31. Assumptions, Research, Risk and Decision Integration

- **COMP-REQ-701:** Mappings shall identify assumptions.
- **COMP-REQ-702:** Mappings shall identify assumption confidence.
- **COMP-REQ-703:** Mappings shall identify assumption expiry.
- **COMP-REQ-704:** Expired assumptions shall trigger mapping review.
- **COMP-REQ-705:** Mappings shall identify research items.
- **COMP-REQ-706:** Mappings shall identify research outcomes.
- **COMP-REQ-707:** Inconclusive research shall constrain mapping strength.
- **COMP-REQ-708:** Mappings shall identify risks.
- **COMP-REQ-709:** Mappings shall identify residual risk.
- **COMP-REQ-710:** Mapping gaps shall create or update risks where material.
- **COMP-REQ-711:** Mappings shall identify decisions.
- **COMP-REQ-712:** Material mapping-method decisions shall use ADRs where applicable.
- **COMP-REQ-713:** Legal-applicability decisions shall identify authority.
- **COMP-REQ-714:** Mapping adoption shall use decision records.
- **COMP-REQ-715:** Mapping changes shall not rewrite research history.
- **COMP-REQ-716:** Mapping changes shall not rewrite assumption history.
- **COMP-REQ-717:** Mapping changes shall not accept risk.
- **COMP-REQ-718:** AI systems shall not accept mapping risks.
- **COMP-REQ-719:** Integration shall conform to documents 26–29.
- **COMP-REQ-720:** Integration history shall remain retrievable.

## 32. Machine-Readable Mapping Records

- **COMP-REQ-721:** Mappings shall be machine-searchable.
- **COMP-REQ-722:** Machine-readable mappings may supplement Markdown.
- **COMP-REQ-723:** Machine-readable mappings shall identify schema version.
- **COMP-REQ-724:** Machine-readable mappings shall identify mapping ID.
- **COMP-REQ-725:** Machine-readable mappings shall identify source record.
- **COMP-REQ-726:** Machine-readable mappings shall identify OBDIA requirement IDs.
- **COMP-REQ-727:** Machine-readable mappings shall identify external reference IDs.
- **COMP-REQ-728:** Machine-readable mappings shall identify relationship.
- **COMP-REQ-729:** Machine-readable mappings shall identify strength.
- **COMP-REQ-730:** Machine-readable mappings shall identify applicability.
- **COMP-REQ-731:** Machine-readable mappings shall identify evidence.
- **COMP-REQ-732:** Machine-readable mappings shall identify gaps.
- **COMP-REQ-733:** Machine-readable mappings shall identify status.
- **COMP-REQ-734:** Machine-readable mappings shall identify review dates.
- **COMP-REQ-735:** Machine-readable mappings shall not contain secrets.
- **COMP-REQ-736:** Machine-readable mappings shall minimize personal data.
- **COMP-REQ-737:** Machine-readable and Markdown representations shall remain synchronized.
- **COMP-REQ-738:** Synchronization conflicts shall block authoritative use.
- **COMP-REQ-739:** Schema migration shall preserve history.
- **COMP-REQ-740:** Machine-readable mapping history shall remain retrievable.

## 33. Repository and Versioning Controls

- **COMP-REQ-741:** The authoritative document shall reside at `docs/knowledge/30_COMPLIANCE_MAPPING_BASELINE.md`.
- **COMP-REQ-742:** Mapping sources shall reside in governed repository paths.
- **COMP-REQ-743:** Mapping exports shall reside in governed repository paths.
- **COMP-REQ-744:** Mapping indexes shall identify versions and statuses.
- **COMP-REQ-745:** Mapping filenames shall follow `OBDIA-NAME-001`.
- **COMP-REQ-746:** Mapping versions shall follow `OBDIA-VER-001`.
- **COMP-REQ-747:** Source-record versions shall remain distinct from mapping versions.
- **COMP-REQ-748:** Baseline versions shall remain distinct from mapping versions.
- **COMP-REQ-749:** Release versions shall remain distinct from mapping versions.
- **COMP-REQ-750:** Repository moves shall preserve history.
- **COMP-REQ-751:** Repository moves shall update links.
- **COMP-REQ-752:** Generated mappings shall identify generator.
- **COMP-REQ-753:** Generated mappings shall receive human review.
- **COMP-REQ-754:** Temporary mappings shall not become authoritative.
- **COMP-REQ-755:** Duplicate authoritative mappings shall be rejected.
- **COMP-REQ-756:** Branch state shall not define mapping status.
- **COMP-REQ-757:** Merge state shall not define mapping status.
- **COMP-REQ-758:** Tags shall not create compliance claims.
- **COMP-REQ-759:** Repository validation shall identify exact commits.
- **COMP-REQ-760:** Repository mapping history shall remain retrievable.

## 34. Testing and Continuous Validation

- **COMP-REQ-761:** Mapping tests shall verify identifier uniqueness.
- **COMP-REQ-762:** Mapping tests shall verify mandatory fields.
- **COMP-REQ-763:** Mapping tests shall verify source-record links.
- **COMP-REQ-764:** Mapping tests shall verify OBDIA requirement links.
- **COMP-REQ-765:** Mapping tests shall verify external reference syntax where machine-checkable.
- **COMP-REQ-766:** Mapping tests shall verify valid applicability states.
- **COMP-REQ-767:** Mapping tests shall verify valid relationship states.
- **COMP-REQ-768:** Mapping tests shall verify valid strength states.
- **COMP-REQ-769:** Mapping tests shall verify review dates.
- **COMP-REQ-770:** Mapping tests shall verify supersession links.
- **COMP-REQ-771:** Mapping tests shall verify no silent identifier reuse.
- **COMP-REQ-772:** Mapping tests shall verify source versions.
- **COMP-REQ-773:** Mapping tests shall verify Markdown and machine-readable synchronization.
- **COMP-REQ-774:** Mapping tests shall verify public disclaimers.
- **COMP-REQ-775:** Mapping tests shall verify no Knowledge document above 30.
- **COMP-REQ-776:** Failed tests shall remain visible.
- **COMP-REQ-777:** Passing structural tests shall not establish mapping correctness.
- **COMP-REQ-778:** Human review shall assess semantic and legal-context accuracy.
- **COMP-REQ-779:** Testing shall conform to `OBDIA-TEST-001`.
- **COMP-REQ-780:** Mapping-test history shall remain retrievable.

## 35. Release and Publication Controls

- **COMP-REQ-781:** Release-relevant mappings shall receive review before release.
- **COMP-REQ-782:** Public mapping artifacts shall identify exact versions.
- **COMP-REQ-783:** Public mapping artifacts shall identify exact source versions.
- **COMP-REQ-784:** Public mapping artifacts shall identify verification dates.
- **COMP-REQ-785:** Public mapping artifacts shall identify applicability limits.
- **COMP-REQ-786:** Public mapping artifacts shall identify mapping strength.
- **COMP-REQ-787:** Public mapping artifacts shall identify gaps.
- **COMP-REQ-788:** Public mapping artifacts shall identify implementation state.
- **COMP-REQ-789:** Public mapping artifacts shall identify validation state.
- **COMP-REQ-790:** Public mapping artifacts shall include non-compliance-claim language.
- **COMP-REQ-791:** Public mapping artifacts shall not expose secrets.
- **COMP-REQ-792:** Public mapping artifacts shall not expose real case data.
- **COMP-REQ-793:** Public mapping artifacts shall not imply institutional adoption.
- **COMP-REQ-794:** Public mapping artifacts shall not imply certification.
- **COMP-REQ-795:** Public mapping artifacts shall not imply legal advice.
- **COMP-REQ-796:** Public mapping corrections shall preserve history.
- **COMP-REQ-797:** Public mapping withdrawal shall preserve records.
- **COMP-REQ-798:** Release Reviewer assessment shall be required.
- **COMP-REQ-799:** Publication shall conform to documents 04 and 15.
- **COMP-REQ-800:** Publication history shall remain retrievable.

## 36. Review Cadence and Triggers

- **COMP-REQ-801:** Active mappings shall have review schedules.
- **COMP-REQ-802:** Legal mappings shall receive source-change review.
- **COMP-REQ-803:** Dynamic threat mappings shall receive version-change review.
- **COMP-REQ-804:** OWASP mappings shall receive project-version review.
- **COMP-REQ-805:** NIST mappings shall receive publication-version review.
- **COMP-REQ-806:** ISO mappings shall receive amendment review.
- **COMP-REQ-807:** OBDIA requirement changes shall trigger review.
- **COMP-REQ-808:** OBDIA architecture changes shall trigger review.
- **COMP-REQ-809:** OBDIA implementation changes shall trigger review.
- **COMP-REQ-810:** OBDIA validation changes shall trigger review.
- **COMP-REQ-811:** Risk changes shall trigger review.
- **COMP-REQ-812:** Assumption changes shall trigger review.
- **COMP-REQ-813:** Research findings shall trigger review.
- **COMP-REQ-814:** Decision changes shall trigger review.
- **COMP-REQ-815:** Security incidents shall trigger review.
- **COMP-REQ-816:** Privacy incidents shall trigger review.
- **COMP-REQ-817:** Release planning shall trigger review.
- **COMP-REQ-818:** Missed review dates shall remain visible.
- **COMP-REQ-819:** Review results shall be attributable.
- **COMP-REQ-820:** Review history shall remain retrievable.

## 37. Metrics and Reporting

- **COMP-REQ-821:** Mapping reports shall identify total mappings.
- **COMP-REQ-822:** Mapping reports shall identify mappings by reference family.
- **COMP-REQ-823:** Mapping reports shall identify mappings by applicability.
- **COMP-REQ-824:** Mapping reports shall identify mappings by relationship.
- **COMP-REQ-825:** Mapping reports shall identify mappings by strength.
- **COMP-REQ-826:** Mapping reports shall identify mappings by implementation state.
- **COMP-REQ-827:** Mapping reports shall identify mappings by validation state.
- **COMP-REQ-828:** Mapping reports shall identify open gaps.
- **COMP-REQ-829:** Mapping reports shall identify critical gaps.
- **COMP-REQ-830:** Mapping reports shall identify overdue reviews.
- **COMP-REQ-831:** Mapping reports shall identify stale sources.
- **COMP-REQ-832:** Mapping reports shall identify unassessed mappings.
- **COMP-REQ-833:** Coverage percentages shall not imply compliance.
- **COMP-REQ-834:** Mapping counts shall not imply control effectiveness.
- **COMP-REQ-835:** Strong-mapping counts shall not imply certification.
- **COMP-REQ-836:** Dashboards shall identify freshness.
- **COMP-REQ-837:** AI-generated summaries shall receive human verification.
- **COMP-REQ-838:** Reports shall minimize personal data.
- **COMP-REQ-839:** Raw records shall remain available to authorized reviewers.
- **COMP-REQ-840:** Reporting history shall remain retrievable.

## 38. Traceability

- **COMP-REQ-841:** Every source record shall trace to an authoritative source.
- **COMP-REQ-842:** Every mapping shall trace to a source record.
- **COMP-REQ-843:** Every mapping shall trace to OBDIA requirements.
- **COMP-REQ-844:** Every mapping shall trace to external objectives.
- **COMP-REQ-845:** Every mapping shall trace to evidence.
- **COMP-REQ-846:** Every mapping shall trace to gaps.
- **COMP-REQ-847:** Every mapping shall trace to assumptions where applicable.
- **COMP-REQ-848:** Every mapping shall trace to research where applicable.
- **COMP-REQ-849:** Every mapping shall trace to risks where applicable.
- **COMP-REQ-850:** Every mapping shall trace to decisions where applicable.
- **COMP-REQ-851:** Every implementation claim shall trace to implementation evidence.
- **COMP-REQ-852:** Every validation claim shall trace to validation evidence.
- **COMP-REQ-853:** Every public claim shall trace to release evidence.
- **COMP-REQ-854:** Every superseded mapping shall trace to successor.
- **COMP-REQ-855:** Every correction shall trace to the original mapping.
- **COMP-REQ-856:** Every source update shall trace to affected mappings.
- **COMP-REQ-857:** Traceability shall be bidirectional.
- **COMP-REQ-858:** Broken traceability affecting legal, security or release claims shall be blocking.
- **COMP-REQ-859:** Traceability records shall not contain secrets.
- **COMP-REQ-860:** Traceability records shall minimize personal and case data.
- **COMP-REQ-861:** Traceability shall identify exact versions.
- **COMP-REQ-862:** Traceability shall identify exact commits where material.
- **COMP-REQ-863:** Traceability shall identify review authority.
- **COMP-REQ-864:** Baseline freeze shall validate mapping traceability across documents 01–30.
- **COMP-REQ-865:** Traceability history shall remain retrievable.

## 39. Baseline Completion and Freeze

- **COMP-REQ-866:** The compliance mapping baseline shall identify all eleven legacy reference families.
- **COMP-REQ-867:** The baseline shall identify exact source records.
- **COMP-REQ-868:** The baseline shall identify exact OBDIA document versions.
- **COMP-REQ-869:** The baseline shall identify source verification dates.
- **COMP-REQ-870:** The baseline shall identify applicability states.
- **COMP-REQ-871:** The baseline shall identify mapping strengths.
- **COMP-REQ-872:** The baseline shall identify gaps.
- **COMP-REQ-873:** The baseline shall identify evidence limitations.
- **COMP-REQ-874:** The baseline shall identify review authority.
- **COMP-REQ-875:** The baseline shall identify repository commit.
- **COMP-REQ-876:** The baseline shall identify integrity reference.
- **COMP-REQ-877:** The baseline shall not contain known misleading claims.
- **COMP-REQ-878:** The baseline shall not hide Not Assessed items.
- **COMP-REQ-879:** The baseline shall not claim legal compliance.
- **COMP-REQ-880:** The baseline shall not claim ISO certification.
- **COMP-REQ-881:** The baseline shall not claim external endorsement.
- **COMP-REQ-882:** The baseline shall not freeze with unresolved critical contradictions.
- **COMP-REQ-883:** The baseline shall receive Project Founder approval before Approved status.
- **COMP-REQ-884:** Post-freeze material changes shall use change control.
- **COMP-REQ-885:** Baseline-freeze history shall remain retrievable.

## 40. Domain-Specific Mapping Controls

These controls ensure that each material OBDIA domain is mapped with explicit authority, source, applicability, evidence, gap and claim boundaries.

- **COMP-REQ-886:** A mapping concerning officer-agent binding shall identify applicable reference families for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-887:** A mapping concerning officer-agent binding shall identify authoritative source versions for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-888:** A mapping concerning officer-agent binding shall identify legal or voluntary classification for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-889:** A mapping concerning officer-agent binding shall identify applicability and jurisdiction for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-890:** A mapping concerning officer-agent binding shall identify OBDIA requirement IDs for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-891:** A mapping concerning officer-agent binding shall identify external objectives or control IDs for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-892:** A mapping concerning officer-agent binding shall identify mapping relationship for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-893:** A mapping concerning officer-agent binding shall identify mapping strength for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-894:** A mapping concerning officer-agent binding shall identify rationale and semantic differences for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-895:** A mapping concerning officer-agent binding shall identify design evidence for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-896:** A mapping concerning officer-agent binding shall identify implementation evidence for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-897:** A mapping concerning officer-agent binding shall identify validation evidence for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-898:** A mapping concerning officer-agent binding shall identify gaps and uncertainty for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-899:** A mapping concerning officer-agent binding shall identify assumptions and research links for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-900:** A mapping concerning officer-agent binding shall identify threats and risk links for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-901:** A mapping concerning officer-agent binding shall identify reviewer and review trigger for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-902:** A mapping concerning officer-agent binding shall identify public-claim boundaries for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-903:** A mapping concerning officer-agent binding shall preserve versioned bidirectional traceability for institutional issuance, one officer, one agent identity and no credential transfer.
- **COMP-REQ-904:** A mapping concerning machine identity shall identify applicable reference families for human, agent, workload and connector identity separation.
- **COMP-REQ-905:** A mapping concerning machine identity shall identify authoritative source versions for human, agent, workload and connector identity separation.
- **COMP-REQ-906:** A mapping concerning machine identity shall identify legal or voluntary classification for human, agent, workload and connector identity separation.
- **COMP-REQ-907:** A mapping concerning machine identity shall identify applicability and jurisdiction for human, agent, workload and connector identity separation.
- **COMP-REQ-908:** A mapping concerning machine identity shall identify OBDIA requirement IDs for human, agent, workload and connector identity separation.
- **COMP-REQ-909:** A mapping concerning machine identity shall identify external objectives or control IDs for human, agent, workload and connector identity separation.
- **COMP-REQ-910:** A mapping concerning machine identity shall identify mapping relationship for human, agent, workload and connector identity separation.
- **COMP-REQ-911:** A mapping concerning machine identity shall identify mapping strength for human, agent, workload and connector identity separation.
- **COMP-REQ-912:** A mapping concerning machine identity shall identify rationale and semantic differences for human, agent, workload and connector identity separation.
- **COMP-REQ-913:** A mapping concerning machine identity shall identify design evidence for human, agent, workload and connector identity separation.
- **COMP-REQ-914:** A mapping concerning machine identity shall identify implementation evidence for human, agent, workload and connector identity separation.
- **COMP-REQ-915:** A mapping concerning machine identity shall identify validation evidence for human, agent, workload and connector identity separation.
- **COMP-REQ-916:** A mapping concerning machine identity shall identify gaps and uncertainty for human, agent, workload and connector identity separation.
- **COMP-REQ-917:** A mapping concerning machine identity shall identify assumptions and research links for human, agent, workload and connector identity separation.
- **COMP-REQ-918:** A mapping concerning machine identity shall identify threats and risk links for human, agent, workload and connector identity separation.
- **COMP-REQ-919:** A mapping concerning machine identity shall identify reviewer and review trigger for human, agent, workload and connector identity separation.
- **COMP-REQ-920:** A mapping concerning machine identity shall identify public-claim boundaries for human, agent, workload and connector identity separation.
- **COMP-REQ-921:** A mapping concerning machine identity shall preserve versioned bidirectional traceability for human, agent, workload and connector identity separation.
- **COMP-REQ-922:** A mapping concerning authorization shall identify applicable reference families for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-923:** A mapping concerning authorization shall identify authoritative source versions for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-924:** A mapping concerning authorization shall identify legal or voluntary classification for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-925:** A mapping concerning authorization shall identify applicability and jurisdiction for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-926:** A mapping concerning authorization shall identify OBDIA requirement IDs for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-927:** A mapping concerning authorization shall identify external objectives or control IDs for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-928:** A mapping concerning authorization shall identify mapping relationship for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-929:** A mapping concerning authorization shall identify mapping strength for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-930:** A mapping concerning authorization shall identify rationale and semantic differences for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-931:** A mapping concerning authorization shall identify design evidence for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-932:** A mapping concerning authorization shall identify implementation evidence for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-933:** A mapping concerning authorization shall identify validation evidence for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-934:** A mapping concerning authorization shall identify gaps and uncertainty for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-935:** A mapping concerning authorization shall identify assumptions and research links for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-936:** A mapping concerning authorization shall identify threats and risk links for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-937:** A mapping concerning authorization shall identify reviewer and review trigger for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-938:** A mapping concerning authorization shall identify public-claim boundaries for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-939:** A mapping concerning authorization shall preserve versioned bidirectional traceability for subject, resource, action, context, default deny and revocation.
- **COMP-REQ-940:** A mapping concerning agent lifecycle shall identify applicable reference families for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-941:** A mapping concerning agent lifecycle shall identify authoritative source versions for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-942:** A mapping concerning agent lifecycle shall identify legal or voluntary classification for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-943:** A mapping concerning agent lifecycle shall identify applicability and jurisdiction for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-944:** A mapping concerning agent lifecycle shall identify OBDIA requirement IDs for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-945:** A mapping concerning agent lifecycle shall identify external objectives or control IDs for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-946:** A mapping concerning agent lifecycle shall identify mapping relationship for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-947:** A mapping concerning agent lifecycle shall identify mapping strength for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-948:** A mapping concerning agent lifecycle shall identify rationale and semantic differences for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-949:** A mapping concerning agent lifecycle shall identify design evidence for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-950:** A mapping concerning agent lifecycle shall identify implementation evidence for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-951:** A mapping concerning agent lifecycle shall identify validation evidence for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-952:** A mapping concerning agent lifecycle shall identify gaps and uncertainty for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-953:** A mapping concerning agent lifecycle shall identify assumptions and research links for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-954:** A mapping concerning agent lifecycle shall identify threats and risk links for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-955:** A mapping concerning agent lifecycle shall identify reviewer and review trigger for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-956:** A mapping concerning agent lifecycle shall identify public-claim boundaries for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-957:** A mapping concerning agent lifecycle shall preserve versioned bidirectional traceability for provisioning, binding, authorization, activation, suspension, revocation and archival.
- **COMP-REQ-958:** A mapping concerning evidence integrity shall identify applicable reference families for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-959:** A mapping concerning evidence integrity shall identify authoritative source versions for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-960:** A mapping concerning evidence integrity shall identify legal or voluntary classification for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-961:** A mapping concerning evidence integrity shall identify applicability and jurisdiction for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-962:** A mapping concerning evidence integrity shall identify OBDIA requirement IDs for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-963:** A mapping concerning evidence integrity shall identify external objectives or control IDs for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-964:** A mapping concerning evidence integrity shall identify mapping relationship for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-965:** A mapping concerning evidence integrity shall identify mapping strength for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-966:** A mapping concerning evidence integrity shall identify rationale and semantic differences for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-967:** A mapping concerning evidence integrity shall identify design evidence for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-968:** A mapping concerning evidence integrity shall identify implementation evidence for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-969:** A mapping concerning evidence integrity shall identify validation evidence for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-970:** A mapping concerning evidence integrity shall identify gaps and uncertainty for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-971:** A mapping concerning evidence integrity shall identify assumptions and research links for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-972:** A mapping concerning evidence integrity shall identify threats and risk links for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-973:** A mapping concerning evidence integrity shall identify reviewer and review trigger for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-974:** A mapping concerning evidence integrity shall identify public-claim boundaries for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-975:** A mapping concerning evidence integrity shall preserve versioned bidirectional traceability for original evidence, provenance, custody, transformations and exports.
- **COMP-REQ-976:** A mapping concerning connector security shall identify applicable reference families for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-977:** A mapping concerning connector security shall identify authoritative source versions for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-978:** A mapping concerning connector security shall identify legal or voluntary classification for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-979:** A mapping concerning connector security shall identify applicability and jurisdiction for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-980:** A mapping concerning connector security shall identify OBDIA requirement IDs for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-981:** A mapping concerning connector security shall identify external objectives or control IDs for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-982:** A mapping concerning connector security shall identify mapping relationship for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-983:** A mapping concerning connector security shall identify mapping strength for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-984:** A mapping concerning connector security shall identify rationale and semantic differences for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-985:** A mapping concerning connector security shall identify design evidence for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-986:** A mapping concerning connector security shall identify implementation evidence for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-987:** A mapping concerning connector security shall identify validation evidence for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-988:** A mapping concerning connector security shall identify gaps and uncertainty for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-989:** A mapping concerning connector security shall identify assumptions and research links for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-990:** A mapping concerning connector security shall identify threats and risk links for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-991:** A mapping concerning connector security shall identify reviewer and review trigger for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-992:** A mapping concerning connector security shall identify public-claim boundaries for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-993:** A mapping concerning connector security shall preserve versioned bidirectional traceability for authentication, egress, validation, provenance, isolation and revocation.
- **COMP-REQ-994:** A mapping concerning prompt injection shall identify applicable reference families for untrusted content, retrieval, memory and tool invocation.
- **COMP-REQ-995:** A mapping concerning prompt injection shall identify authoritative source versions for untrusted content, retrieval, memory and tool invocation.
- **COMP-REQ-996:** A mapping concerning prompt injection shall identify legal or voluntary classification for untrusted content, retrieval, memory and tool invocation.
- **COMP-REQ-997:** A mapping concerning prompt injection shall identify applicability and jurisdiction for untrusted content, retrieval, memory and tool invocation.
- **COMP-REQ-998:** A mapping concerning prompt injection shall identify OBDIA requirement IDs for untrusted content, retrieval, memory and tool invocation.
- **COMP-REQ-999:** A mapping concerning prompt injection shall identify external objectives or control IDs for untrusted content, retrieval, memory and tool invocation.

## 41. Verified External Source Baseline

The following source records were verified against official issuing-authority pages on 2026-08-06. Dynamic sources require future snapshot review.

| Source Record | Reference | Verified Version or Snapshot | Source Type | Project Use |
|---|---|---|---|---|
| `COMP-SRC-001` | NIST Cybersecurity Framework | CSF 2.0; NIST CSWP 29; published 2024-02-26 | Voluntary cybersecurity risk framework | Governance, risk, protection, detection, response and recovery design reference |
| `COMP-SRC-002` | NIST Artificial Intelligence Risk Management Framework | AI RMF 1.0; NIST AI 100-1; published 2023-01-26; revision activity noted by NIST | Voluntary AI risk framework | Govern, Map, Measure and Manage design reference |
| `COMP-SRC-003` | NIST SP 800-207 Zero Trust Architecture | Final publication; August 2020 | Voluntary technical architecture guidance | Zero Trust identity, policy and enforcement design reference |
| `COMP-SRC-004` | MITRE ATT&CK | v19.1; current from 2026-04-28 at verification | Dynamic threat-knowledge base | Defensive threat, detection and mitigation traceability |
| `COMP-SRC-005` | MITRE ATLAS | Official ATLAS snapshot verified 2026-08-06 | Dynamic AI threat-knowledge base | AI, LLM and agentic threat traceability |
| `COMP-SRC-006` | OWASP Application Security Verification Standard | ASVS 5.0.0 | Open verification standard | Application-security requirements and testing reference |
| `COMP-SRC-007` | OWASP Top 10 | Top 10:2025 | Awareness document | Application-risk awareness and review reference |
| `COMP-SRC-008` | OWASP Artificial Intelligence Security Verification Standard | AISVS 1.0; released 2026-06-24 | Open AI verification standard | AI lifecycle security requirements and testing reference |
| `COMP-SRC-009` | OWASP Large Language Model Security Verification Standard | LLMSVS 2.0; publication year 2026 | Open LLM verification standard | LLM integration, model lifecycle, memory, agents and monitoring reference |
| `COMP-SRC-010` | ISO/IEC 27001 | ISO/IEC 27001:2022, Edition 3; Amendment 1:2024 | Voluntary international management-system standard | Information-security management-system design reference |
| `COMP-SRC-011` | ISO/IEC 42001 | ISO/IEC 42001:2023, Edition 1 | Voluntary international management-system standard | AI management-system design reference |
| `COMP-SRC-012` | General Data Protection Regulation | Regulation (EU) 2016/679 | EU legal instrument | Privacy and data-protection applicability and design review |
| `COMP-SRC-013` | EU Artificial Intelligence Act | Regulation (EU) 2024/1689 | EU legal instrument | AI regulatory applicability and design review |
| `COMP-SRC-014` | NIS2 | Directive (EU) 2022/2555 | EU legal instrument requiring jurisdiction-specific transposition analysis | Cybersecurity-governance applicability and design review |

## 42. Initial Design-Level Mapping Records

All initial records are `Draft`, use a `Contextual` or `Supporting` relationship unless stated otherwise, and identify a baseline gap: implementation, validation, organizational applicability and legal sufficiency are not established by the repository.

| Mapping ID | Reference Objective | Primary OBDIA Artifacts | Relationship | Strength | Applicability | Baseline Gap |
|---|---|---|---|---|---|---|
| `COMP-MAP-001` | NIST CSF Govern | 01, 05, 09, 10, 13, 25, 26, 29 | Supporting | Moderate | Contextual | Organizational governance implementation and validation absent |
| `COMP-MAP-002` | NIST CSF Identify | 07, 08, 16, 26, 27, 28 | Supporting | Moderate | Contextual | Asset inventory and operational risk evidence absent |
| `COMP-MAP-003` | NIST CSF Protect and Detect | 06, 16, 17, 19, 20, 21, 22 | Supporting | Moderate | Contextual | Implemented technical controls and monitoring evidence absent |
| `COMP-MAP-004` | NIST CSF Respond and Recover | 16, 18, 19, 22, 24 | Supporting | Weak | Contextual | Complete incident-response and recovery implementation absent |
| `COMP-MAP-005` | NIST AI RMF Govern | 01, 05, 09, 16, 26, 28, 29 | Supporting | Moderate | Contextual | Operational AI governance evidence absent |
| `COMP-MAP-006` | NIST AI RMF Map | 07, 08, 17, 18, 20, 26, 27 | Supporting | Moderate | Contextual | Use-case and stakeholder impact assessment remains synthetic |
| `COMP-MAP-007` | NIST AI RMF Measure | 18, 22, 26, 27, 28 | Supporting | Moderate | Contextual | Empirical measurement and external validation absent |
| `COMP-MAP-008` | NIST AI RMF Manage | 09, 16, 19, 20, 26, 29 | Supporting | Moderate | Contextual | Implemented treatment and monitoring evidence absent |
| `COMP-MAP-009` | SP 800-207 identity and resource focus | 06, 08, 16, 17 | Supporting | Strong | Contextual | Cryptographic implementation and validation absent |
| `COMP-MAP-010` | SP 800-207 policy decision and enforcement | 17, 20 | Supporting | Strong | Contextual | PDP, PAP, PIP and PEP implementation absent |
| `COMP-MAP-011` | SP 800-207 continuous verification | 16, 17, 19, 20 | Supporting | Moderate | Contextual | Continuous re-evaluation metrics absent |
| `COMP-MAP-012` | SP 800-207 no implicit network trust | 08, 16, 20, 23 | Supporting | Strong | Contextual | Network implementation and segmentation evidence absent |
| `COMP-MAP-013` | ATT&CK adversary behavior traceability | 07, 16, 26 | Contextual | Moderate | Contextual | Detailed technique-level mapping incomplete |
| `COMP-MAP-014` | ATT&CK detection strategies | 16, 20, 22 | Supporting | Weak | Contextual | Detection engineering and telemetry implementation absent |
| `COMP-MAP-015` | ATT&CK incident analysis | 16, 18, 24, 26 | Contextual | Weak | Contextual | Incident playbooks and exercises incomplete |
| `COMP-MAP-016` | ATT&CK authorized defensive testing | 07, 22, 27 | Supporting | Moderate | Contextual | Technique-specific safe test catalog incomplete |
| `COMP-MAP-017` | ATLAS prompt and context attacks | 07, 16, 20, 21, 22, 26 | Supporting | Moderate | Contextual | Full ATLAS technique crosswalk and tests incomplete |
| `COMP-MAP-018` | ATLAS AI supply chain and poisoning | 07, 16, 21, 22, 26 | Supporting | Moderate | Contextual | Model and data supply-chain implementation absent |
| `COMP-MAP-019` | ATLAS agent and tool abuse | 17, 19, 20, 22, 26 | Supporting | Strong | Contextual | Agentic abuse-case validation incomplete |
| `COMP-MAP-020` | ATLAS leakage, extraction and exfiltration | 16, 18, 20, 22, 26 | Supporting | Moderate | Contextual | Empirical privacy and exfiltration tests incomplete |
| `COMP-MAP-021` | OWASP ASVS application-security verification | 16, 17, 20, 21, 22 | Supporting | Moderate | Contextual | No implemented application exists for full ASVS verification |
| `COMP-MAP-022` | OWASP Top 10:2025 awareness | 07, 16, 21, 22, 26 | Contextual | Moderate | Contextual | Awareness mapping is not verification |
| `COMP-MAP-023` | OWASP AISVS AI security verification | 16, 19, 20, 21, 22, 26 | Supporting | Moderate | Contextual | Requirement-level AISVS crosswalk and implementation evidence absent |
| `COMP-MAP-024` | OWASP LLMSVS LLM and agent security | 16, 17, 19, 20, 21, 22 | Supporting | Moderate | Contextual | Requirement-level LLMSVS crosswalk and validation absent |
| `COMP-MAP-025` | ISO/IEC 27001 organizational context and leadership | 01, 05, 09, 16 | Supporting | Moderate | Contextual | No organizational ISMS scope or certification context established |
| `COMP-MAP-026` | ISO/IEC 27001 risk planning and treatment | 07, 16, 26, 28, 29 | Supporting | Moderate | Contextual | Operational risk treatment evidence absent |
| `COMP-MAP-027` | ISO/IEC 27001 operation and controls | 16–24 | Supporting | Moderate | Contextual | Implemented control environment absent |
| `COMP-MAP-028` | ISO/IEC 27001 evaluation and improvement | 22, 24, 25, 29 | Supporting | Weak | Contextual | Internal audit, management review and improvement evidence absent |
| `COMP-MAP-029` | ISO/IEC 42001 AI governance and policy | 01, 05, 09, 16, 26, 29 | Supporting | Moderate | Contextual | Organizational AIMS scope and leadership evidence absent |
| `COMP-MAP-030` | ISO/IEC 42001 AI risk and impact | 07, 16, 26, 27, 28 | Supporting | Moderate | Contextual | Formal AI impact-assessment process incomplete |
| `COMP-MAP-031` | ISO/IEC 42001 lifecycle, data and model management | 18, 19, 20, 21, 22 | Supporting | Moderate | Contextual | Implemented AI lifecycle and data controls absent |
| `COMP-MAP-032` | ISO/IEC 42001 monitoring and improvement | 16, 22, 24, 25 | Supporting | Weak | Contextual | Operational performance evaluation absent |
| `COMP-MAP-033` | GDPR principles and accountability | 01, 05, 09, 16, 18, 26 | Contextual | Moderate | Potentially Applicable | Processing roles, purposes and lawful basis not established |
| `COMP-MAP-034` | GDPR privacy by design and default | 05, 16, 17, 18, 20, 21 | Supporting | Moderate | Potentially Applicable | Implemented processing system and data inventory absent |
| `COMP-MAP-035` | GDPR rights, transparency and correction | 18, 26, 27, 28, 29 | Contextual | Weak | Potentially Applicable | Data-subject workflow and applicability analysis absent |
| `COMP-MAP-036` | GDPR security, breach and impact assessment | 16, 18, 22, 24, 26 | Supporting | Moderate | Potentially Applicable | Operational security and breach process absent |
| `COMP-MAP-037` | EU AI Act governance and human oversight | 01, 05, 09, 17, 19, 26 | Contextual | Moderate | Not Assessed | AI-system and actor classification not established |
| `COMP-MAP-038` | EU AI Act risk, data and technical documentation | 16, 18, 21, 22, 26 | Contextual | Moderate | Not Assessed | Legal applicability and required-system classification unresolved |
| `COMP-MAP-039` | EU AI Act logging, transparency and robustness | 16, 18, 19, 22, 23 | Contextual | Moderate | Not Assessed | Operational system, role and application-date analysis absent |
| `COMP-MAP-040` | EU AI Act post-market and incident obligations | 16, 22, 24, 26, 29 | Contextual | Weak | Not Assessed | Market-placement and provider/deployer facts absent |
| `COMP-MAP-041` | NIS2 governance and management accountability | 09, 16, 26, 29 | Contextual | Weak | Not Assessed | Entity and national-transposition applicability unresolved |
| `COMP-MAP-042` | NIS2 risk-management measures | 07, 16, 20, 21, 22, 26 | Supporting | Moderate | Not Assessed | Operational entity scope and implemented controls absent |
| `COMP-MAP-043` | NIS2 incident handling and reporting | 16, 18, 22, 24, 26 | Contextual | Weak | Not Assessed | National reporting process and entity facts absent |
| `COMP-MAP-044` | NIS2 supply chain and continuity | 16, 20, 21, 24, 26 | Supporting | Moderate | Not Assessed | Supplier, service and continuity environment absent |

## 43. Minimum Validation Checklist

Before an individual mapping may support an approved baseline or public claim, confirm:

- [ ] immutable `COMP-MAP-NNN` identifier exists;
- [ ] authoritative source record exists;
- [ ] source title, version, date and status are verified;
- [ ] source type and legal-versus-voluntary classification are accurate;
- [ ] jurisdiction and applicability are explicit;
- [ ] OBDIA document and requirement versions are exact;
- [ ] external control, outcome, article, objective or technique is identified;
- [ ] mapping relationship and direction are explicit;
- [ ] mapping strength has evidence-backed rationale;
- [ ] semantic differences and exclusions are recorded;
- [ ] design, implementation, testing and validation states are separate;
- [ ] evidence is current, attributable and scoped;
- [ ] gaps and conflicts remain visible;
- [ ] assumptions, research, risks and decisions are linked;
- [ ] reviewer competence and role concentration are disclosed;
- [ ] source and OBDIA change triggers are current;
- [ ] public claim boundaries are adjacent to the mapping;
- [ ] mapping does not claim certification, legal compliance or endorsement;
- [ ] legal applicability has qualified human review where required;
- [ ] machine-readable and Markdown records agree where used;
- [ ] supersession and correction history are preserved;
- [ ] traceability is bidirectional;
- [ ] exact repository commit and integrity reference are recorded;
- [ ] Project Founder approval exists before baseline status becomes Approved.

## 44. Limitations

- The initial 44 mappings are design-level records, not control-by-control legal or certification assessments.
- Source versions can change after the 2026-08-06 verification date.
- MITRE and OWASP references are dynamic and require dated snapshots.
- ISO standards are copyrighted; detailed clause and control mapping requires authorized access.
- Legal applicability depends on facts, roles, territory, sector, processing and national law.
- The repository does not establish controller, processor, provider, deployer, essential-entity, important-entity or high-risk-AI-system status.
- Mapping strength cannot prove implementation or effectiveness.
- A complete crosswalk cannot replace testing, audit, conformity assessment, certification or legal review.
- Internal mapping review is not independent external assurance.
- Future implementation may create new applicability, evidence and gap conditions.

## 45. Change Control

Every material change shall identify rationale, affected sources, versions, mappings, applicability, evidence, gaps, security and privacy impact, legal-context impact, migration, validation, rollback, authority and semantic-version effect.

- Editorial corrections shall use a patch version when meaning is unchanged.
- Backward-compatible substantive additions shall use a minor version.
- Incompatible mapping-method changes shall use a major version.
- Material mapping architecture decisions shall require an ADR where applicable.
- Changes shall require Project Founder approval.
- Changes shall receive Privacy and Governance Reviewer assessment.
- Security, implementation, research and release impacts shall receive specialist review where applicable.
- Changes shall identify affected source records, mappings, schemas, indexes and automation.
- Changes shall include migration and compatibility analysis.
- Changes shall include validation and rollback analysis.
- Changes shall not retroactively fabricate applicability, implementation, validation or compliance evidence.
- Historical mappings shall not be silently rewritten.
- Mapping and requirement identifiers shall not be reused.
- Migration shall preserve source-to-requirement and mapping-to-evidence traceability.
- The complete Knowledge Pack 01–30 shall be reconciled before this baseline becomes Approved.

## 46. Consolidation Record

Version 1.0.0 consolidates the two existing Compliance Mapping Baseline variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-COMP-001`;
- normalizes the authoritative filename to `30_COMPLIANCE_MAPPING_BASELINE.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves NIST CSF;
- preserves NIST AI RMF;
- preserves NIST SP 800-207;
- preserves MITRE ATT&CK;
- preserves MITRE ATLAS;
- preserves OWASP;
- preserves ISO/IEC 27001;
- preserves ISO/IEC 42001;
- preserves GDPR;
- preserves the EU AI Act;
- preserves NIS2;
- preserves the rule that mappings are design references unless explicitly stated otherwise;
- preserves the rule that mappings support design and review;
- preserves the prohibition on treating mappings as proof of legal compliance;
- adds authoritative-source versions, dates, legal-versus-voluntary classification, applicability, jurisdiction, mapping relationships, strength, gaps, evidence, reviewers, update triggers, uncertainty, validation, machine-readable records, publication boundaries and bidirectional traceability;
- establishes 14 verified source records;
- establishes 44 initial design-level mapping records;
- adds 999 stable mapping-governance requirements;
- closes the fixed Knowledge Pack 01–30 without creating document 31;
- treats Enterprise and non-Enterprise files as legacy source variants of the same immutable document;
- creates no implementation, validation, conformity, certification, legal-compliance, regulatory-approval, public-release or operational-authority claim.

## 47. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-06 | Draft | Privacy and Governance Reviewer; approval reserved to Project Founder | Constitutional consolidation of both legacy compliance-mapping baselines: preserved all eleven reference families and claim disclaimers; added complete source, applicability, strength, gap, evidence, legal-boundary, validation and traceability governance with 14 verified sources and 44 initial design mappings. |
