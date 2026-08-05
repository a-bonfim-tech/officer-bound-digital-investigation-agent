# RELEASE AND PUBLICATION POLICY

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-REL-001 |
| **Title** | Release and Publication Policy |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Release Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define the release mechanics, evidence, manifests, integrity controls, approvals, version and tag handling, channels, corrections, withdrawal, supersession and archival required before any OBDIA artifact is distributed or represented as published. |
| **Scope** | Private and public releases, release candidates, source and documentation bundles, diagrams, demonstrations, laboratory packages, generated artifacts, tags, changelogs, manifests, validation evidence, corrections, withdrawals and archives within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `04_PORTFOLIO_AND_PUBLICATION_POLICY.md`; `09_GOVERNANCE_MODEL.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; forward dependencies `22_TESTING_STANDARD.md`, `24_REPOSITORY_STRUCTURE_STANDARD.md`, `25_DOCUMENT_VERSIONING_POLICY.md` and `30_COMPLIANCE_MAPPING_POLICY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-RES-001`; `OBDIA-PUB-001`; `OBDIA-GOV-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001` |
| **Cross-References** | Documents 16–30; accepted ADRs; threat models; release manifests; validation records; risk and exception records; correction, withdrawal, supersession and archive records |
| **Assumptions** | The repository is an independent research and engineering project. Public distribution is optional and shall occur only after an explicit release decision for an exact artifact set. |
| **Constraints** | A release shall not create independent AI legal authority, authorize prohibited conduct, imply institutional adoption, expose secrets or real case data, interact with uncontrolled criminal infrastructure, or overstate implementation, validation, compliance or production readiness. |
| **Security Considerations** | Release activity can disclose secrets, personal data, vulnerable configurations, unsafe instructions, supply-chain artifacts, evidentiary material, internal paths and misleading assurance claims. |
| **Validation Criteria** | Each release has an exact reviewed commit, manifest, integrity references, changelog, accurate status, completed security and privacy checks, safe-demonstration evidence, attributable approval, reproducible instructions where applicable, and tested correction or withdrawal procedures. |
| **Implementation Relationship** | This policy governs release mechanics. It does not authorize implementation, production deployment, institutional use, investigative activity or legal access. |

---

## 1. Purpose

This policy governs how OBDIA artifacts become release candidates, approved releases, corrected releases, withdrawn releases, superseded releases and archived release records.

The original baseline required:

- completed security review;
- complete documentation;
- no secrets;
- no personal data;
- safe demonstrations;
- a version tag;
- an updated changelog;
- accurate distinction between implemented features and proposed architecture.

This consolidation preserves those controls and establishes the evidence and governance needed to apply them reliably.

`04_PORTFOLIO_AND_PUBLICATION_POLICY.md` remains authoritative for public positioning, portfolio claims, safe-demonstration content and misrepresentation limits. This document is authoritative for release mechanics, manifests, integrity, tagging, approval evidence, correction, withdrawal, supersession and archival.


## 2. Fundamental Release Rules

- **REL-REQ-001:** Every release decision shall be attributable to an authorized human role.
- **REL-REQ-002:** An AI agent shall not approve, publish, withdraw, supersede or archive a release.
- **REL-REQ-003:** A repository merge shall not establish a release.
- **REL-REQ-004:** A Git tag shall not establish a release by itself.
- **REL-REQ-005:** Public repository visibility shall not establish Published state by itself.
- **REL-REQ-006:** Release approval shall apply only to the exact artifacts and versions reviewed.
- **REL-REQ-007:** Release status shall be supported by retained evidence.
- **REL-REQ-008:** Release content shall remain subordinate to the Canonical Project Definition and Constitution.
- **REL-REQ-009:** Release mechanics shall not override the public-positioning authority of `OBDIA-PUB-001`.
- **REL-REQ-010:** A release shall not authorize prohibited conduct.
- **REL-REQ-011:** A release shall not imply independent legal authority for an AI agent.
- **REL-REQ-012:** A release shall not imply institutional adoption, certification or operational deployment without evidence.
- **REL-REQ-013:** Private incorporation and public publication shall remain distinct.
- **REL-REQ-014:** Source release, binary release, documentation release and demonstration release shall remain distinguishable.
- **REL-REQ-015:** Failure of a mandatory release gate shall block release.
- **REL-REQ-016:** Silence, timeout or automation failure shall not be interpreted as release approval.

## 3. Authority Boundary with Document 04

- **REL-REQ-017:** `OBDIA-PUB-001` shall govern public narrative, portfolio positioning and professional claims.
- **REL-REQ-018:** `OBDIA-REL-001` shall govern release candidates, manifests, tags, approvals, channels, corrections, withdrawal and archival.
- **REL-REQ-019:** A release record shall not amend public-positioning rules.
- **REL-REQ-020:** A public-positioning decision shall not substitute for release-gate evidence.
- **REL-REQ-021:** A technical release shall identify whether it is public, private, internal, experimental or archival.
- **REL-REQ-022:** An experimental release shall remain explicitly labeled and shall not imply production readiness.
- **REL-REQ-023:** A documentation-only release shall not imply implementation.
- **REL-REQ-024:** A source-only release shall not imply that builds were reproduced or validated.
- **REL-REQ-025:** A demonstration release shall identify synthetic, mock, testnet, local-laboratory or other controlled boundaries.
- **REL-REQ-026:** Conflict between documents 04 and 15 shall be resolved according to the constitutional hierarchy and recorded.

## 4. Release Classes and Channels

Permitted classes include release candidate, private research release, public source release, documentation release, demonstration release, corrective release, archival release and withdrawal record.

- **REL-REQ-027:** Release classes shall be defined before gate review.
- **REL-REQ-028:** A release candidate shall be immutable in content after review begins, except through a new candidate.
- **REL-REQ-029:** A private research release shall identify its intended recipients and access restrictions.
- **REL-REQ-030:** A public source release shall identify included source, documentation and generated artifacts.
- **REL-REQ-031:** A documentation release shall identify whether implementation artifacts are excluded.
- **REL-REQ-032:** A demonstration release shall identify execution environment and data classification.
- **REL-REQ-033:** An archival release shall be read-only and shall not be represented as current authority.
- **REL-REQ-034:** A corrective release shall identify the defect and affected predecessor.
- **REL-REQ-035:** A security correction may restrict technical detail while preserving an accountable record.
- **REL-REQ-036:** A withdrawn release shall remain represented by a governed withdrawal record.
- **REL-REQ-037:** A superseding release shall identify every release it replaces wholly or partially.
- **REL-REQ-038:** Release channels shall be approved and documented.
- **REL-REQ-039:** Channel promotion shall require a separate decision when the audience or risk changes.
- **REL-REQ-040:** A release copied to another channel shall preserve manifest and integrity references.
- **REL-REQ-041:** Unofficial mirrors shall not be represented as authoritative release channels.

## 5. Roles and Separation of Duties

- **REL-REQ-042:** The Release Reviewer shall own completion of the release gate.
- **REL-REQ-043:** The Project Founder shall approve public release of normative or project-defining artifacts.
- **REL-REQ-044:** The Security Reviewer shall review security-sensitive release content.
- **REL-REQ-045:** The Privacy and Governance Reviewer shall review personal-data, rights, jurisdiction and governance impacts.
- **REL-REQ-046:** The Documentation Authority shall verify metadata, filenames, cross-references, status and changelog integrity.
- **REL-REQ-047:** The Implementation Reviewer shall verify implementation-status claims where implementation is included.
- **REL-REQ-048:** The Research Reviewer shall verify material factual, standards and regulatory claims.
- **REL-REQ-049:** Contributors shall disclose known release limitations and unresolved findings.
- **REL-REQ-050:** Role concentration shall be disclosed.
- **REL-REQ-051:** Internal multi-role review shall not be represented as independent external assurance.
- **REL-REQ-052:** A release approver shall identify the exact candidate reviewed.
- **REL-REQ-053:** Repository write access shall not confer release authority.
- **REL-REQ-054:** Release authority shall not be delegated to an AI system.
- **REL-REQ-055:** Conflicts of interest shall be disclosed and addressed.
- **REL-REQ-056:** A blocking specialist reviewer finding shall prevent release until resolved or lawfully dispositioned.

## 6. Release Candidate Requirements

- **REL-REQ-057:** A release candidate shall have a unique candidate identifier.
- **REL-REQ-058:** A release candidate shall identify the exact source commit.
- **REL-REQ-059:** A release candidate shall identify included and excluded artifacts.
- **REL-REQ-060:** A release candidate shall identify its proposed version and channel.
- **REL-REQ-061:** A release candidate shall identify lifecycle and implementation status.
- **REL-REQ-062:** A release candidate shall identify dependencies and required tools.
- **REL-REQ-063:** A release candidate shall identify known limitations and residual risks.
- **REL-REQ-064:** A release candidate shall identify data classifications and demonstration boundaries.
- **REL-REQ-065:** A release candidate shall identify required reviewers and gates.
- **REL-REQ-066:** A release candidate shall identify rollback, correction and withdrawal paths.
- **REL-REQ-067:** A material candidate change shall create a new candidate or new immutable candidate reference.
- **REL-REQ-068:** Review evidence shall not be reused after a material candidate change without reassessment.

## 7. Release Manifest

- **REL-REQ-069:** Every release shall include or reference a release manifest.
- **REL-REQ-070:** The manifest shall identify release identifier, title, version and date.
- **REL-REQ-071:** The manifest shall identify the exact source commit.
- **REL-REQ-072:** The manifest shall identify included files and artifact paths.
- **REL-REQ-073:** The manifest shall identify cryptographic integrity values for material artifacts where feasible.
- **REL-REQ-074:** The manifest shall identify generated artifacts and their authoritative sources.
- **REL-REQ-075:** The manifest shall identify dependency manifests, lockfiles or equivalent records where applicable.
- **REL-REQ-076:** The manifest shall identify model, prompt, policy and configuration versions where material.
- **REL-REQ-077:** The manifest shall identify test and validation evidence.
- **REL-REQ-078:** The manifest shall identify applicable ADRs and normative requirements.
- **REL-REQ-079:** The manifest shall identify release approvers and gate decisions.
- **REL-REQ-080:** The manifest shall identify known limitations, deviations and accepted residual risks.
- **REL-REQ-081:** The manifest shall identify license and attribution information where applicable.
- **REL-REQ-082:** The manifest shall identify release channel and access classification.
- **REL-REQ-083:** The manifest shall identify predecessor and successor releases where applicable.
- **REL-REQ-084:** The manifest shall not contain secrets, credentials or unnecessary personal data.
- **REL-REQ-085:** Manifest corrections shall preserve prior records and explain the correction.
- **REL-REQ-086:** Manifest integrity shall be verifiable.

## 8. Integrity, Provenance and Packaging

- **REL-REQ-087:** Material release artifacts shall have recorded SHA-256 or an approved equivalent integrity value.
- **REL-REQ-088:** Integrity values shall be calculated from the exact distributed bytes.
- **REL-REQ-089:** Integrity records shall identify the algorithm.
- **REL-REQ-090:** Integrity verification instructions shall be documented.
- **REL-REQ-091:** A mismatch shall block distribution or trigger withdrawal.
- **REL-REQ-092:** Signed tags or attestations may supplement but shall not replace release approval.
- **REL-REQ-093:** Signing keys shall be protected and independently revocable.
- **REL-REQ-094:** Unsigned releases shall not be represented as signed.
- **REL-REQ-095:** Build provenance shall identify source, build process and environment where builds are distributed.
- **REL-REQ-096:** Generated artifacts shall remain traceable to authoritative sources.
- **REL-REQ-097:** Release packaging shall not modify reviewed content silently.
- **REL-REQ-098:** Repackaging shall create a new integrity reference.
- **REL-REQ-099:** Release archives shall use deterministic or documented packaging where feasible.
- **REL-REQ-100:** Artifact integrity limitations shall be disclosed.
- **REL-REQ-101:** Integrity evidence shall be retained with the release record.

## 9. Version, Tag and Changelog Controls

- **REL-REQ-102:** Every release shall have an approved version identifier.
- **REL-REQ-103:** Version syntax shall remain consistent with the Constitution and document 25 after consolidation.
- **REL-REQ-104:** A release tag shall correspond to the reviewed release commit.
- **REL-REQ-105:** A tag shall not be moved or reused after publication.
- **REL-REQ-106:** A corrected tag shall use a new version or governed correction process.
- **REL-REQ-107:** Pre-release identifiers shall not be used normatively until document 25 is reconciled.
- **REL-REQ-108:** Release version and document versions shall be distinguishable.
- **REL-REQ-109:** Interdependent document-version compatibility shall be recorded.
- **REL-REQ-110:** Version synchronization requirements shall be reviewed before release.
- **REL-REQ-111:** A changelog shall identify material additions, changes, removals, security effects and limitations.
- **REL-REQ-112:** The changelog shall not fabricate implementation or validation.
- **REL-REQ-113:** The changelog shall identify breaking normative or interface changes.
- **REL-REQ-114:** The changelog shall identify superseded or withdrawn artifacts.
- **REL-REQ-115:** Changelog corrections shall preserve historical traceability.
- **REL-REQ-116:** Document 25 remains a blocking forward dependency for normative approval of release-version mechanics.

## 10. Documentation Completeness

- **REL-REQ-117:** Release documentation shall be complete for the intended release class.
- **REL-REQ-118:** Installation or execution instructions shall state prerequisites and supported environment.
- **REL-REQ-119:** Commands shall distinguish user input from expected output.
- **REL-REQ-120:** Instructions shall not request production credentials or secrets.
- **REL-REQ-121:** Demonstration instructions shall use synthetic or authorized controlled resources.
- **REL-REQ-122:** Known failures and limitations shall remain visible.
- **REL-REQ-123:** Documentation shall identify what is conceptual, designed, implemented, tested, simulated, deferred and out of scope.
- **REL-REQ-124:** Implemented shall not be used as a synonym for Tested, Validated, Secure, Compliant or Production Ready.
- **REL-REQ-125:** Documentation shall identify exact versions where reproducibility depends on them.
- **REL-REQ-126:** Broken critical links shall block release.
- **REL-REQ-127:** Documentation shall comply with `OBDIA-DOC-001`.
- **REL-REQ-128:** Release notes shall identify applicable security and privacy considerations.
- **REL-REQ-129:** Release notes shall identify correction and reporting channels.
- **REL-REQ-130:** Documentation-only changes shall not overstate technical capability.

## 11. Security Review

- **REL-REQ-131:** Security review shall be completed before release.
- **REL-REQ-132:** Release review shall examine the current tree and repository history for secrets.
- **REL-REQ-133:** No password, token, API key, private key or reusable credential shall be released.
- **REL-REQ-134:** No real case data shall be released.
- **REL-REQ-135:** No unauthorized personal data shall be released.
- **REL-REQ-136:** No unlawfully obtained material shall be released.
- **REL-REQ-137:** No hidden offensive capability shall be released.
- **REL-REQ-138:** No malware deployment capability shall be released.
- **REL-REQ-139:** No functionality facilitating unauthorized access shall be released.
- **REL-REQ-140:** No uncontrolled interaction with criminal infrastructure shall be released.
- **REL-REQ-141:** Unsafe debug interfaces and insecure defaults shall not be enabled in distributed configurations.
- **REL-REQ-142:** Dependency and supply-chain risks shall be reviewed.
- **REL-REQ-143:** Known exploitable defects affecting the release shall be resolved, contained or block release.
- **REL-REQ-144:** Security limitations shall not be concealed.
- **REL-REQ-145:** Redaction shall not fabricate security or invalidate required evidence.
- **REL-REQ-146:** Security-sensitive withdrawal procedures shall support rapid containment.
- **REL-REQ-147:** Document 30 remains a forward dependency for complete compliance-mapping review.

## 12. Privacy and Fundamental-Rights Review

- **REL-REQ-148:** Privacy and fundamental-rights review shall be completed when release content may affect persons or protected data.
- **REL-REQ-149:** Public releases shall not identify real officers, victims, suspects or case subjects without explicit lawful authorization.
- **REL-REQ-150:** Personal data shall be minimized.
- **REL-REQ-151:** Sensitive metadata shall be removed from distributed artifacts where not required.
- **REL-REQ-152:** Example identities shall be fictitious.
- **REL-REQ-153:** Telemetry shall be disabled by default or clearly disclosed and controlled.
- **REL-REQ-154:** Release analytics shall not collect unnecessary personal data.
- **REL-REQ-155:** Retention and deletion behavior shall be documented where applicable.
- **REL-REQ-156:** Public-source data shall not be represented as unrestricted for reuse.
- **REL-REQ-157:** Jurisdictional and legal uncertainty shall be disclosed rather than converted into compliance claims.
- **REL-REQ-158:** Privacy redaction shall preserve technical meaning and provenance where required.
- **REL-REQ-159:** A privacy incident shall trigger release reassessment and possible withdrawal.

## 13. Safe Demonstration Gate

- **REL-REQ-160:** Public demonstrations shall use synthetic, fictitious, mock, testnet, local-laboratory, lawful public, archived-authorized or controlled cybersecurity resources.
- **REL-REQ-161:** Demonstrations shall not investigate real persons.
- **REL-REQ-162:** Demonstrations shall not use production officer identities.
- **REL-REQ-163:** Demonstrations shall not use production credentials.
- **REL-REQ-164:** Dark Web demonstrations shall be simulated, archived-authorized or operationally isolated.
- **REL-REQ-165:** Demonstrations shall not transact with criminal marketplaces.
- **REL-REQ-166:** Blockchain demonstrations shall use testnets unless an explicitly authorized read-only public analysis is documented.
- **REL-REQ-167:** Mock integrations shall be labeled as mock.
- **REL-REQ-168:** Simulated capabilities shall not be represented as operational deployment.
- **REL-REQ-169:** Demonstration reset and cleanup procedures shall be documented.
- **REL-REQ-170:** Demonstration outputs shall not be represented as original evidence.
- **REL-REQ-171:** Abuse-enabling details shall be minimized without concealing material limitations.
- **REL-REQ-172:** Safe-demonstration evidence shall be retained in the release record.

## 14. Claim and Status Accuracy

- **REL-REQ-173:** Public release descriptions shall accurately distinguish proposed architecture from implementation.
- **REL-REQ-174:** Claims shall distinguish implementation from testing and validation.
- **REL-REQ-175:** Claims shall distinguish internal review from independent assurance.
- **REL-REQ-176:** Claims shall not describe the project as an official police system.
- **REL-REQ-177:** Claims shall not describe the project as deployed in real investigations.
- **REL-REQ-178:** Claims shall not claim public-authority approval without evidence.
- **REL-REQ-179:** Claims shall not state guaranteed compliance.
- **REL-REQ-180:** Standards mappings shall not be represented as proof of compliance.
- **REL-REQ-181:** Claims shall not state production readiness without evidence and authorized governance.
- **REL-REQ-182:** Claims shall not use unsupported superlatives.
- **REL-REQ-183:** Metrics shall identify method, scope and limitations.
- **REL-REQ-184:** Screenshots shall reflect actual reviewed artifacts and status.
- **REL-REQ-185:** Release pages shall display the project’s independent-research status.
- **REL-REQ-186:** Release pages shall display material limitations.
- **REL-REQ-187:** Claims shall follow `OBDIA-RES-001` and `OBDIA-PUB-001`.

## 15. Mandatory Release Gate

- **REL-REQ-188:** The release gate shall identify the exact release candidate.
- **REL-REQ-189:** The release gate shall include scope and dependency review.
- **REL-REQ-190:** The release gate shall include architecture consistency review when architecture is represented.
- **REL-REQ-191:** The release gate shall include security review.
- **REL-REQ-192:** The release gate shall include privacy and fundamental-rights review where applicable.
- **REL-REQ-193:** The release gate shall include terminology and documentation review.
- **REL-REQ-194:** The release gate shall include implementation-alignment review when implementation is included.
- **REL-REQ-195:** The release gate shall include publication and release review.
- **REL-REQ-196:** The release gate shall verify complete documentation.
- **REL-REQ-197:** The release gate shall verify no secrets.
- **REL-REQ-198:** The release gate shall verify no unauthorized personal data.
- **REL-REQ-199:** The release gate shall verify safe demonstrations only.
- **REL-REQ-200:** The release gate shall verify version, tag and changelog consistency.
- **REL-REQ-201:** The release gate shall verify manifest and integrity references.
- **REL-REQ-202:** The release gate shall verify correction and withdrawal capability.
- **REL-REQ-203:** Gate non-applicability shall include rationale.
- **REL-REQ-204:** Conditional approval shall identify conditions, owner and expiry.
- **REL-REQ-205:** A failed mandatory gate shall block release.
- **REL-REQ-206:** Approval shall identify actor, role, date, scope and exact candidate.
- **REL-REQ-207:** Automated checks may provide evidence but shall not replace human release approval.

## 16. Approval Evidence

- **REL-REQ-208:** A release decision shall be recorded in a governed release record.
- **REL-REQ-209:** The record shall identify approved, rejected, conditionally approved or withdrawn outcome.
- **REL-REQ-210:** The record shall identify reviewed commit and artifact manifest.
- **REL-REQ-211:** The record shall identify all required reviewers.
- **REL-REQ-212:** The record shall identify unresolved limitations and residual risks.
- **REL-REQ-213:** The record shall identify release channel and effective time.
- **REL-REQ-214:** The record shall identify the release tag or immutable release reference.
- **REL-REQ-215:** The record shall identify correction, withdrawal and contact procedures.
- **REL-REQ-216:** The record shall preserve role-concentration disclosure.
- **REL-REQ-217:** The record shall not claim independent external review when none occurred.
- **REL-REQ-218:** Rejected candidates shall retain their decision evidence.
- **REL-REQ-219:** Approval shall not alter normative document status by implication.

## 17. Release Automation

- **REL-REQ-220:** Release automation shall operate on an exact reviewed commit.
- **REL-REQ-221:** Release automation shall use least privilege.
- **REL-REQ-222:** Automation credentials shall be scoped and excluded from logs.
- **REL-REQ-223:** Automation configuration shall be versioned and reviewed.
- **REL-REQ-224:** Automation shall fail closed when mandatory evidence is missing.
- **REL-REQ-225:** Automation shall not create final release approval.
- **REL-REQ-226:** Automation shall not move or overwrite an existing release tag silently.
- **REL-REQ-227:** Automation shall verify artifact integrity after packaging.
- **REL-REQ-228:** Automation shall produce auditable logs without exposing secrets.
- **REL-REQ-229:** Generated release notes shall receive human review.
- **REL-REQ-230:** AI-generated summaries shall not fabricate changes, tests, risks or approvals.
- **REL-REQ-231:** Automation failure shall leave the candidate unreleased or clearly incomplete.
- **REL-REQ-232:** Manual override shall require authority, rationale and evidence.
- **REL-REQ-233:** Release automation remains experimental until implemented and validated under documents 21–25.

## 18. Corrections

- **REL-REQ-234:** A material release defect shall receive a tracked correction decision.
- **REL-REQ-235:** A correction shall identify affected versions and artifacts.
- **REL-REQ-236:** A correction shall identify security, privacy, evidence and user impact.
- **REL-REQ-237:** A correction shall preserve the original release record.
- **REL-REQ-238:** A replacement release shall use a new immutable version or approved correction identifier.
- **REL-REQ-239:** A silent replacement of distributed artifacts is prohibited.
- **REL-REQ-240:** Correction notices shall be distributed through appropriate channels.
- **REL-REQ-241:** Correction urgency shall be proportionate to risk.
- **REL-REQ-242:** Security corrections may coordinate disclosure timing without fabricating status.
- **REL-REQ-243:** Corrected documentation shall identify the corrected claim or instruction.
- **REL-REQ-244:** Correction evidence shall include validation of the replacement.
- **REL-REQ-245:** A correction shall not be represented as independent certification.

## 19. Withdrawal and Takedown

- **REL-REQ-246:** A release shall be withdrawn when continued distribution creates unacceptable security, privacy, legal, evidence or misrepresentation risk.
- **REL-REQ-247:** A withdrawal decision shall be attributable.
- **REL-REQ-248:** A withdrawal record shall identify affected release identifiers and channels.
- **REL-REQ-249:** Withdrawal shall preserve manifest, approval and reason records.
- **REL-REQ-250:** Withdrawal shall remove or restrict distributable artifacts where feasible.
- **REL-REQ-251:** Withdrawal shall not erase repository or audit history.
- **REL-REQ-252:** Withdrawal shall identify safe replacement or mitigation where available.
- **REL-REQ-253:** Credential or secret exposure shall trigger immediate containment and rotation.
- **REL-REQ-254:** Personal-data exposure shall trigger privacy-incident handling.
- **REL-REQ-255:** Unsafe capability exposure shall trigger containment and security review.
- **REL-REQ-256:** Users shall receive clear notice when reliance on a withdrawn release is unsafe.
- **REL-REQ-257:** Republication after withdrawal requires a new release gate.
- **REL-REQ-258:** Failed withdrawal shall remain an unresolved incident.

## 20. Supersession and Archival

- **REL-REQ-259:** A superseding release shall identify predecessor releases.
- **REL-REQ-260:** Supersession scope shall be explicit.
- **REL-REQ-261:** Superseded releases shall not be represented as current.
- **REL-REQ-262:** Compatibility and migration implications shall be documented.
- **REL-REQ-263:** Partial supersession shall identify unaffected components.
- **REL-REQ-264:** Release indexes shall identify current, superseded, withdrawn and archived releases.
- **REL-REQ-265:** Archive records shall preserve integrity values and manifests.
- **REL-REQ-266:** Archived releases shall be read-only.
- **REL-REQ-267:** Archive access shall follow classification and least privilege.
- **REL-REQ-268:** Retention and lawful deletion rules shall be documented.
- **REL-REQ-269:** Restoration from archive shall not reactivate Published state silently.
- **REL-REQ-270:** Supersession and archival shall remain consistent with `OBDIA-STATE-001`.

## 21. Release Incident Handling

- **REL-REQ-271:** A release-related incident shall have an immutable incident identifier.
- **REL-REQ-272:** Incident handling shall identify affected releases and artifacts.
- **REL-REQ-273:** Incident handling shall preserve evidence.
- **REL-REQ-274:** Incident handling shall distinguish containment, correction, withdrawal and supersession.
- **REL-REQ-275:** Incident communications shall be accurate and proportionate.
- **REL-REQ-276:** Incident communications shall not expose unnecessary sensitive detail.
- **REL-REQ-277:** Material incidents shall trigger threat, risk and governance reassessment.
- **REL-REQ-278:** Post-incident review shall identify control and process improvements.
- **REL-REQ-279:** Repeated release failures shall trigger architectural or governance review.
- **REL-REQ-280:** Incident closure shall require evidence that required release actions completed.

## 22. Traceability

- **REL-REQ-281:** Every release requirement shall trace to applicable policy, control, evidence and decision.
- **REL-REQ-282:** Release traceability shall be bidirectional.
- **REL-REQ-283:** Manifest entries shall trace to repository artifacts.
- **REL-REQ-284:** Validation results shall trace to exact release artifacts.
- **REL-REQ-285:** Release tags shall trace to exact reviewed commits.
- **REL-REQ-286:** Changelog entries shall trace to material changes.
- **REL-REQ-287:** Risk and exception records shall trace to release decisions.
- **REL-REQ-288:** Correction and withdrawal records shall trace to affected releases.
- **REL-REQ-289:** Broken traceability affecting security, privacy, authority or publication shall block release.
- **REL-REQ-290:** Planned validation shall not be represented as completed validation.
- **REL-REQ-291:** Release evidence shall use immutable identifiers where available.
- **REL-REQ-292:** Traceability records shall not contain secrets or unnecessary personal data.

## 23. Validation Requirements

- **REL-REQ-293:** Release validation shall use the exact candidate artifacts.
- **REL-REQ-294:** Validation shall identify environment and tool versions.
- **REL-REQ-295:** Validation shall verify manifest completeness.
- **REL-REQ-296:** Validation shall verify integrity values.
- **REL-REQ-297:** Validation shall verify installation or execution instructions where applicable.
- **REL-REQ-298:** Validation shall verify safe-demonstration boundaries.
- **REL-REQ-299:** Validation shall verify failure and cleanup procedures where applicable.
- **REL-REQ-300:** Validation shall verify no secret or unauthorized data is present.
- **REL-REQ-301:** Validation shall verify status and claim accuracy.
- **REL-REQ-302:** Validation shall verify correction and withdrawal contacts.
- **REL-REQ-303:** Validation limitations shall remain visible.
- **REL-REQ-304:** Document 22 remains a forward dependency for complete testing governance.
- **REL-REQ-305:** Document 24 remains a forward dependency for repository packaging and placement.
- **REL-REQ-306:** Document 25 remains a forward dependency for complete release versioning.
- **REL-REQ-307:** Document 30 remains a forward dependency for compliance-mapping claims.

## 24. Minimum Release Checklist

Before release, confirm:

- [ ] exact release candidate and source commit are fixed;
- [ ] release class and channel are defined;
- [ ] manifest is complete;
- [ ] integrity references are verified;
- [ ] security review is complete;
- [ ] privacy and fundamental-rights review is complete where applicable;
- [ ] documentation is complete;
- [ ] no secrets or credentials are present;
- [ ] no unauthorized personal or case data is present;
- [ ] only safe demonstrations are included;
- [ ] implementation and lifecycle states are accurate;
- [ ] proposed architecture is distinguishable from implementation;
- [ ] version, tag and changelog are consistent;
- [ ] dependencies, licenses and attributions are reviewed;
- [ ] unresolved risks and limitations are visible;
- [ ] correction and withdrawal procedures are ready;
- [ ] reviewers and role concentration are disclosed;
- [ ] exact approval evidence is retained;
- [ ] forward dependencies are recorded;
- [ ] Project Founder approval is recorded when required.


## 25. Limitations

- This policy does not require the project to publish publicly.
- It does not authorize implementation, deployment, institutional adoption or investigative use.
- It does not replace the public-positioning authority of document 04.
- It does not define all testing mechanics; document 22 remains a forward dependency.
- It does not define the complete repository structure; document 24 remains a forward dependency.
- It does not define complete semantic-version and pre-release rules; document 25 remains a forward dependency.
- It does not prove legal or regulatory compliance.
- Internal release review is not independent certification.
- Cryptographic integrity does not prove content safety or correctness.
- Withdrawal cannot guarantee removal of every third-party copy.


## 26. Change Control

Every material change shall identify rationale, affected releases and documents, security and privacy impact, implementation impact, migration, validation, rollback, authority and semantic-version effect.

- **REL-REQ-308:** Editorial corrections use a patch version when meaning is unchanged.
- **REL-REQ-309:** Backward-compatible substantive additions use a minor version.
- **REL-REQ-310:** Incompatible release-governance changes use a major version.
- **REL-REQ-311:** Material release-architecture decisions require an ADR where applicable.
- **REL-REQ-312:** Changes require Project Founder approval.
- **REL-REQ-313:** Unapproved release-methodology changes during an active phase are prohibited.
- **REL-REQ-314:** Changes shall preserve historical manifests, decisions, corrections and withdrawals.
- **REL-REQ-315:** Changes shall include migration, validation and rollback analysis.
- **REL-REQ-316:** Forward-dependency reconciliation is required before this policy becomes Approved.
- **REL-REQ-317:** Changes shall not retroactively fabricate release evidence or approval.

## 27. Consolidation Record

Version 1.0.0 consolidates the existing Release and Publication Policy without expanding project scope. It:

- retains immutable document identifier `OBDIA-REL-001`;
- normalizes the authoritative filename to `15_RELEASE_AND_PUBLICATION_POLICY.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves the original release checklist: security review, complete documentation, no secrets, no personal data, safe demonstrations, version tag and changelog;
- preserves the original requirement to distinguish implemented features from proposed architecture;
- makes document 04 authoritative for public positioning and document 15 authoritative for release mechanics;
- adds release classes, candidates, manifests, integrity, provenance, channels, approval evidence, automation, correction, withdrawal, supersession, archival, incident handling, validation and traceability;
- identifies documents 22, 24, 25 and 30 as forward dependencies that block normative approval until reconciled;
- treats legacy Enterprise and non-Enterprise source variants as the same immutable document;
- creates no release, implementation, public-availability decision or operational authorization.

## 28. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Release Reviewer; approval reserved to Project Founder | Constitutional consolidation of the original release checklist: preserved all original controls and implementation-transparency rule; added complete release mechanics, evidence, integrity, correction, withdrawal, supersession, archival, validation and change control. |
