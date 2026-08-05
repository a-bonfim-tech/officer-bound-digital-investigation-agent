# REPOSITORY STATE MODEL

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-STATE-001 |
| **Title** | Repository State Model |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define the authoritative lifecycle states, transition rules, state evidence, responsible roles, guards, rollback, supersession, archival and repository representation for governed OBDIA artifacts. |
| **Scope** | Knowledge documents 01–30, constitutional artifacts, ADRs, specifications, diagrams, implementation, tests, evidence, research records, releases, indexes and other governed repository artifacts. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `09_GOVERNANCE_MODEL.md`; `11_DOCUMENTATION_STANDARD.md`; forward dependency `25_DOCUMENT_VERSIONING_POLICY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GOV-001`; `OBDIA-DOC-001` |
| **Cross-References** | Documents 14–30; accepted ADRs; status indexes; review evidence; implementation evidence; validation records; release records; supersession and archive records |
| **Assumptions** | Repository artifacts may have different lifecycle applicability, but every governed artifact must declare an accurate state or an approved `Not Applicable` disposition. |
| **Constraints** | This model shall implement the constitutional lifecycle exactly, shall not create alternative normative states, and shall not treat branches, Pull Requests, commits, tags or automation results as approval by themselves. |
| **Security Considerations** | Incorrect or stale state can cause unauthorized implementation, misleading publication, use of superseded controls, evidence loss, unreviewed changes, unsafe releases and fabricated assurance. |
| **Validation Criteria** | Every governed artifact has an attributable current state, evidence-backed transitions, permitted-state guards, exact artifact references, accurate status representation, preserved history and controlled rollback, supersession and archival. |
| **Implementation Relationship** | Repository automation may validate and display state, but only authorized human decisions and retained evidence can establish governed lifecycle transitions. |

---

## 1. Purpose

This document defines the authoritative repository state model for the Officer-Bound Digital Investigation Agent project.

The exact constitutional lifecycle is:

`Draft → Under Review → Approved → Implemented → Validated → Published → Superseded → Archived`

The original state-model sources used `Reviewed` in place of `Under Review`, used `Tested` in place of `Validated`, and omitted `Superseded`. Those legacy terms are not normative states. This consolidation replaces them with the constitutional lifecycle while preserving the original intent that every artifact declare its current state and responsible reviewer.

A lifecycle state represents evidenced governance status. It does not describe a Git branch, a Pull Request workflow state, an issue label, a commit, a tag, a deployment environment or an AI-generated assessment.

## 2. Fundamental State Rules

- **STATE-REQ-001:** Every governed artifact shall declare its current lifecycle state.
- **STATE-REQ-002:** Every state declaration shall identify or reference the responsible owner.
- **STATE-REQ-003:** Every material transition shall identify an authorized human decision-maker.
- **STATE-REQ-004:** The authoritative lifecycle states are exactly `Draft`, `Under Review`, `Approved`, `Implemented`, `Validated`, `Published`, `Superseded` and `Archived`.
- **STATE-REQ-005:** Alternative spellings, synonyms or abbreviated states shall not replace the authoritative states.
- **STATE-REQ-006:** `Reviewed` is a legacy term and shall not be used as a normative lifecycle state.
- **STATE-REQ-007:** `Tested` is an implementation-transparency or evidence descriptor and shall not be used as a normative lifecycle state.
- **STATE-REQ-008:** Repository existence shall not establish approval.
- **STATE-REQ-009:** Pull Request merge shall not establish approval automatically.
- **STATE-REQ-010:** Automated checks shall not establish approval, validation or publication independently.
- **STATE-REQ-011:** An AI agent shall not approve, validate, publish, supersede or archive an artifact.
- **STATE-REQ-012:** State and semantic version shall remain distinct.
- **STATE-REQ-013:** State and implementation status shall remain distinct.
- **STATE-REQ-014:** State and document classification shall remain distinct.
- **STATE-REQ-015:** An unsupported state declaration shall be treated as a governance defect.
- **STATE-REQ-016:** Ambiguous state shall resolve to the lower-assurance state until evidence is established.

## 3. State Record

A governed state record shall include, where applicable:

- artifact identifier;
- artifact title;
- repository path or immutable reference;
- version;
- classification;
- current state;
- state-effective date and time;
- owner;
- transition initiator;
- approving or responsible role;
- prior state;
- transition decision;
- applicable review gates;
- evidence references;
- unresolved findings;
- conditions or expiry;
- successor or predecessor references;
- rollback reference;
- integrity reference;
- repository commit or tag where applicable.

- **STATE-REQ-017:** State records shall use immutable artifact identifiers where available.
- **STATE-REQ-018:** State records shall identify the exact artifact version or commit evaluated.
- **STATE-REQ-019:** State evidence shall remain attributable.
- **STATE-REQ-020:** State records shall not contain secrets or unnecessary personal data.
- **STATE-REQ-021:** A state record shall distinguish decision evidence from supporting automated evidence.
- **STATE-REQ-022:** Missing mandatory transition evidence shall block the transition.
- **STATE-REQ-023:** Historical state records shall not be silently overwritten.
- **STATE-REQ-024:** Corrections to state records shall preserve the original record and explain the correction.
- **STATE-REQ-025:** State records affecting evidence or release shall use tamper-evident integrity controls where technically feasible.
- **STATE-REQ-026:** Repository state indexes shall reference authoritative records rather than replace them.

## 4. Lifecycle Applicability

Not every artifact requires every lifecycle state.

Examples:

- a normative document may progress through the complete lifecycle;
- an ADR uses the status model defined by `OBDIA-ADR-001`, while its repository artifact may also declare governance state where required;
- an evidence record may not have an `Implemented` state;
- a template may not have `Validated` or `Published` states;
- an internal record may never become `Published`.

- **STATE-REQ-027:** State applicability shall be defined by artifact class and governing policy.
- **STATE-REQ-028:** A non-applicable state shall be recorded as `Not Applicable` with rationale.
- **STATE-REQ-029:** `Not Applicable` is a disposition, not a ninth lifecycle state.
- **STATE-REQ-030:** A state shall not be skipped merely for convenience.
- **STATE-REQ-031:** The authority approving non-applicability shall be recorded.
- **STATE-REQ-032:** Non-applicability shall not be used to avoid a mandatory review, implementation, validation or release gate.
- **STATE-REQ-033:** A change in artifact class or intended use shall trigger lifecycle-applicability review.
- **STATE-REQ-034:** Forward lifecycle planning shall not be represented as completed state.
- **STATE-REQ-035:** An artifact may remain indefinitely in a state only when the state remains accurate and review requirements are satisfied.
- **STATE-REQ-036:** Stale state shall be reviewed when dependencies, authority, implementation or risk materially change.

## 5. Draft

### 5.1 Meaning

`Draft` means the artifact is being created or materially revised and has not completed formal review.

### 5.2 Entry Criteria

An artifact may enter Draft when:

- a valid artifact identifier is assigned or an approved exception exists;
- purpose and intended scope are identified;
- an owner is assigned;
- classification is proposed;
- dependencies are initially identified;
- no higher-assurance claim is made.

### 5.3 Permitted Changes

Substantive changes are permitted, subject to repository controls and preservation of attributable history.

### 5.4 Exit Criteria

Draft may exit to `Under Review` only when:

- mandatory metadata is complete;
- required sections contain substantive content;
- dependencies and cross-references are identified;
- requirements are identifiable;
- known limitations are recorded;
- reviewers and gates are identified;
- the exact review artifact is fixed.

- **STATE-REQ-037:** Draft shall not be represented as approved authority.
- **STATE-REQ-038:** Draft requirements shall not be treated as binding implementation authority.
- **STATE-REQ-039:** Draft incorporation into `main` shall remain visibly labeled Draft.
- **STATE-REQ-040:** Draft placeholders shall have owners and resolution criteria.
- **STATE-REQ-041:** Draft may be abandoned without approval, but any material review history shall be preserved.
- **STATE-REQ-042:** A Draft artifact shall not authorize operational activity.
- **STATE-REQ-043:** Draft publication shall not occur unless an approved policy explicitly permits publication as a labeled draft.
- **STATE-REQ-044:** Material dependencies missing from Draft shall be recorded as blockers or forward dependencies.

## 6. Under Review

### 6.1 Meaning

`Under Review` means the artifact is formally submitted to identified reviewers for the applicable constitutional gates.

### 6.2 Entry Criteria

- Draft exit criteria are satisfied;
- exact artifact version or commit is recorded;
- reviewers and review scope are identified;
- applicable gates are listed;
- unresolved findings have tracking identifiers.

### 6.3 Permitted Changes

Changes are controlled. A material change invalidates affected review evidence and requires re-review.

### 6.4 Exit Criteria

Under Review may exit to `Approved` only when:

- all applicable gates pass;
- blocking findings are resolved;
- non-applicable gates have approved rationales;
- cross-references are validated;
- approval authority records an explicit decision;
- the exact approved artifact is identified.

It may return to `Draft` when substantial revision is required.

- **STATE-REQ-045:** Under Review shall not be represented as Approved.
- **STATE-REQ-046:** Reviewer assignment shall not itself constitute completed review.
- **STATE-REQ-047:** Review evidence shall identify reviewer role, date, artifact and disposition.
- **STATE-REQ-048:** Critical unresolved findings block approval.
- **STATE-REQ-049:** Material changes after review begin shall trigger affected reviewer reassessment.
- **STATE-REQ-050:** Conditional review results shall identify conditions, owner and review date.
- **STATE-REQ-051:** Role concentration shall be disclosed.
- **STATE-REQ-052:** Internal review shall not be labeled independent external review.
- **STATE-REQ-053:** Under Review may return to Draft without erasing findings.
- **STATE-REQ-054:** Automated review output remains advisory until adopted by an authorized human reviewer.

## 7. Approved

### 7.1 Meaning

`Approved` means the artifact has completed applicable review and has been explicitly approved by the authorized human authority.

### 7.2 Entry Criteria

- applicable review gates passed;
- exact artifact reference recorded;
- approval decision attributable;
- blocking findings resolved;
- residual limitations recorded;
- version and cross-references validated.

### 7.3 Permitted Changes

Only non-substantive metadata corrections are permitted without a new controlled version. Substantive change returns the successor version to Draft.

### 7.4 Exit Criteria

Approved may proceed to `Implemented` when implementation exists and is traceable, or may move toward `Superseded` when replaced by a later approved artifact.

- **STATE-REQ-055:** Approval shall not be inferred from silence.
- **STATE-REQ-056:** Approval shall not be inferred from merge, signature, tag or deployment alone.
- **STATE-REQ-057:** Approval shall identify actor, role, date, scope and exact artifact.
- **STATE-REQ-058:** Approved status shall not imply implementation.
- **STATE-REQ-059:** Approved status shall not imply validation.
- **STATE-REQ-060:** Approved status shall not imply publication.
- **STATE-REQ-061:** Approved status shall not imply legal compliance, certification or institutional deployment.
- **STATE-REQ-062:** An approved artifact shall remain subordinate to superior authority.
- **STATE-REQ-063:** Discovery of a material defect shall trigger review, suspension, correction or supersession.
- **STATE-REQ-064:** An approval decision shall not authorize a canonical prohibition.
- **STATE-REQ-065:** A materially changed approved artifact shall use a new version and re-enter Draft.
- **STATE-REQ-066:** Approved historical content shall not be silently rewritten.

## 8. Implemented

### 8.1 Meaning

`Implemented` means the approved requirements or decisions have a corresponding implementation that is traceable to the artifact.

### 8.2 Entry Criteria

- approved governing artifact exists;
- implementation commit and relevant configuration are identified;
- required ADRs are Accepted;
- applicable threat model and entry gates are complete;
- implementation traceability exists;
- implementation status is accurately scoped.

### 8.3 Permitted Changes

Implementation changes follow controlled development, review and versioning. A material implementation change may require updated architecture, ADRs, threat modeling and artifact versions.

### 8.4 Exit Criteria

Implemented may proceed to `Validated` only when validation criteria are executed and evidence is retained.

- **STATE-REQ-067:** Implemented shall not mean Tested, Validated, Secure, Compliant or Production Ready.
- **STATE-REQ-068:** Implemented state shall identify the exact implementation scope.
- **STATE-REQ-069:** Partial implementation shall not establish artifact-level Implemented state when mandatory requirements remain absent.
- **STATE-REQ-070:** Implementation shall trace to approved requirements and Accepted ADRs.
- **STATE-REQ-071:** Non-conforming implementation shall not support Implemented state.
- **STATE-REQ-072:** Implementation evidence shall identify code, configuration, policy, model and environment versions where applicable.
- **STATE-REQ-073:** Repository code existence shall not establish Implemented state.
- **STATE-REQ-074:** A proof of concept shall not be represented as full implementation.
- **STATE-REQ-075:** Implementation changes invalidating alignment shall trigger state reassessment.
- **STATE-REQ-076:** Implemented state shall not authorize operational deployment.
- **STATE-REQ-077:** Implementation evidence shall comply with `12_IMPLEMENTATION_POLICY.md`.
- **STATE-REQ-078:** Forward dependency on `25_DOCUMENT_VERSIONING_POLICY.md` shall be resolved before this model becomes Approved.

## 9. Validated

### 9.1 Meaning

`Validated` means the artifact or corresponding implementation has been evaluated against defined validation criteria, with retained evidence and reviewer disposition.

### 9.2 Entry Criteria

- applicable Approved or Implemented artifact exists;
- validation plan and acceptance criteria exist;
- exact versions and environment are identified;
- validation is executed;
- results and limitations are retained;
- blocking failures are resolved or formally dispositioned;
- authorized reviewer confirms the result.

### 9.3 Permitted Changes

Validation records are immutable evidence except for additive corrections preserving original results. Material artifact or implementation change invalidates affected validation.

### 9.4 Exit Criteria

Validated may proceed to `Published` only when the publication and release gate passes, or may remain internal.

- **STATE-REQ-079:** Validated shall replace the legacy lifecycle term `Tested`.
- **STATE-REQ-080:** A test execution alone shall not establish Validated state.
- **STATE-REQ-081:** Validation shall identify exact artifact and environment versions.
- **STATE-REQ-082:** Validation shall include material negative and failure-path results where applicable.
- **STATE-REQ-083:** Failed validation shall remain visible.
- **STATE-REQ-084:** Passing results shall not be generalized beyond tested scope.
- **STATE-REQ-085:** Validation shall record limitations and residual risk.
- **STATE-REQ-086:** An AI system shall not serve as final validation authority.
- **STATE-REQ-087:** Automated checks may support but shall not replace authorized validation review.
- **STATE-REQ-088:** Material post-validation change invalidates affected validation evidence.
- **STATE-REQ-089:** Validated shall not imply Published.
- **STATE-REQ-090:** Validated shall not imply production readiness, certification or legal compliance.

## 10. Published

### 10.1 Meaning

`Published` means the artifact has passed the publication and release gate and is made available through an approved release channel with accurate status and limitations.

### 10.2 Entry Criteria

- underlying artifact state supports publication;
- release reviewer completes the publication gate;
- exact release artifact, commit or tag is identified;
- secret, privacy, legal, security and claim reviews pass;
- release limitations and status are accurate;
- withdrawal and correction paths are defined.

### 10.3 Permitted Changes

Published historical content is immutable. Corrections require a new version, correction notice, superseding release or withdrawal record.

### 10.4 Exit Criteria

Published may become `Superseded` when replaced, or may be withdrawn under the release policy while preserving history.

- **STATE-REQ-091:** Private repository incorporation shall not establish Published state.
- **STATE-REQ-092:** Public branch visibility shall not establish Published state.
- **STATE-REQ-093:** Published state requires an attributable release decision.
- **STATE-REQ-094:** Release evidence shall identify exact artifacts and integrity references.
- **STATE-REQ-095:** Published limitations and implementation status shall remain visible.
- **STATE-REQ-096:** Published content shall not imply official endorsement without evidence.
- **STATE-REQ-097:** Published content shall not conceal unresolved material security or privacy limitations.
- **STATE-REQ-098:** Publication shall not expose secrets, real case data or prohibited capability.
- **STATE-REQ-099:** A correction shall not silently alter published history.
- **STATE-REQ-100:** Withdrawal shall preserve a governed record and reason.
- **STATE-REQ-101:** Published state shall follow documents 04 and 15.
- **STATE-REQ-102:** A release tag shall not establish publication without gate evidence.

## 11. Superseded

### 11.1 Meaning

`Superseded` means a later approved artifact explicitly replaces all or part of the artifact.

### 11.2 Entry Criteria

- successor artifact is identified;
- replacement scope is explicit;
- dependency impact is assessed;
- cross-references are redirected without erasing history;
- implementation and migration impact are recorded;
- supersession decision is attributable.

### 11.3 Permitted Changes

No substantive edits. Additive metadata may identify the successor, scope and effective date.

### 11.4 Exit Criteria

Superseded may proceed to `Archived` when archival controls and successor references are complete.

- **STATE-REQ-103:** Superseded is mandatory in the constitutional lifecycle and shall not be omitted.
- **STATE-REQ-104:** Superseded shall not be used when no approved successor exists.
- **STATE-REQ-105:** Partial supersession shall identify exact replaced scope.
- **STATE-REQ-106:** The superseded artifact shall remain retrievable.
- **STATE-REQ-107:** Supersession shall not rewrite the original decision or requirements.
- **STATE-REQ-108:** Successor and predecessor shall reference each other where feasible.
- **STATE-REQ-109:** Active implementation shall migrate, receive an approved exception or be blocked.
- **STATE-REQ-110:** Superseded artifacts shall not be used as current authority.
- **STATE-REQ-111:** Dependency indexes shall reflect supersession.
- **STATE-REQ-112:** Published superseded artifacts shall display successor information through approved release mechanisms.

## 12. Archived

### 12.1 Meaning

`Archived` is the terminal governed state for artifacts retained for history, audit, evidence, legal, research or reproducibility purposes.

### 12.2 Entry Criteria

- retention and access requirements are defined;
- integrity is verified;
- successor or terminal disposition is recorded;
- active references are redirected;
- archive owner or custodian is identified;
- secrets and restricted data are appropriately controlled.

### 12.3 Permitted Changes

Archived artifacts are read-only except for additive archival metadata that preserves the original content.

### 12.4 Exit Criteria

Archived is terminal unless restoration occurs through a recorded governance decision creating a new controlled version or state record.

- **STATE-REQ-113:** Archived artifacts shall remain retrievable subject to lawful access and retention controls.
- **STATE-REQ-114:** Archival shall not erase review, approval, implementation, validation or release history.
- **STATE-REQ-115:** Archived artifacts shall not be used as current authority.
- **STATE-REQ-116:** Archive integrity shall be verifiable.
- **STATE-REQ-117:** Archive access shall follow least privilege.
- **STATE-REQ-118:** Archive metadata shall identify retention and disposal rules.
- **STATE-REQ-119:** Restoration shall not reactivate obsolete authority silently.
- **STATE-REQ-120:** Restoration shall use a new controlled action and evidence record.
- **STATE-REQ-121:** Legal or security restriction may limit visibility but shall preserve authorized evidence.
- **STATE-REQ-122:** Final deletion, where lawfully required, shall have an attributable disposition record.

## 13. State Transition Matrix

| From | To | Permitted | Minimum Authority | Required Evidence |
|---|---|---:|---|---|
| Draft | Under Review | Yes | Owner and Documentation Authority | Complete review artifact; reviewers; gates; exact commit |
| Under Review | Draft | Yes | Owner or blocking reviewer | Findings; revision rationale; preserved review history |
| Under Review | Approved | Yes | Approval Authority after applicable reviewers | Passed gates; resolved findings; exact artifact; approval record |
| Approved | Implemented | Yes, where applicable | Implementation Reviewer | Traceable implementation; Accepted ADRs; threat and entry-gate evidence |
| Implemented | Validated | Yes, where applicable | Implementation or Validation Reviewer | Executed validation; exact versions; results; limitations |
| Validated | Published | Yes, where applicable | Release Reviewer and required authority | Release gate; manifest; exact release reference |
| Approved, Implemented, Validated or Published | Superseded | Yes | Owning authority and Documentation Authority | Approved successor; impact analysis; successor linkage |
| Superseded | Archived | Yes | Documentation Authority or custodian | Archive manifest; integrity; retention and access record |
| Any active state | Prior state | Not as direct erasure | Authorized rollback decision | New rollback record; cause; impact; restored version |
| Archived | Active state | Not directly | Authorized restoration decision | New version or record; rationale; review and validation |

- **STATE-REQ-123:** Only permitted transitions shall be accepted.
- **STATE-REQ-124:** A transition not listed as direct requires an explicit governance decision and preserved intermediate evidence.
- **STATE-REQ-125:** State transitions shall not be retroactively fabricated.
- **STATE-REQ-126:** Transition authority shall be no broader than the role’s governance authority.
- **STATE-REQ-127:** A blocked transition shall not proceed through an informal override.
- **STATE-REQ-128:** A canonical prohibition cannot be waived by a state transition.
- **STATE-REQ-129:** Transition conditions shall be evaluated against current dependencies and risk.
- **STATE-REQ-130:** Invalid transitions shall generate a reviewable finding.

## 14. State Guards and Blocking Conditions

A state transition shall be blocked when applicable conditions include:

- missing owner or authority;
- unresolved canonical conflict;
- invalid identifier or filename;
- missing mandatory metadata;
- unresolved critical security finding;
- unresolved material privacy or fundamental-rights finding;
- broken authority or cross-reference chain;
- missing required ADR;
- missing threat model;
- non-conforming implementation;
- failed material validation;
- missing release review;
- inaccurate implementation status;
- exposed secrets or unauthorized data;
- missing successor for supersession;
- missing retention or integrity evidence for archival.

- **STATE-REQ-131:** Blocking conditions shall have identifiers and owners.
- **STATE-REQ-132:** Blocking findings shall remain visible until resolved or formally rejected by authorized authority.
- **STATE-REQ-133:** Risk acceptance shall not remove a canonical prohibition.
- **STATE-REQ-134:** Conditional transitions shall identify enforceable conditions and expiry.
- **STATE-REQ-135:** Expired conditions shall trigger denial, rollback or re-review.
- **STATE-REQ-136:** Automation shall fail closed when required state evidence is missing.
- **STATE-REQ-137:** Manual override shall require authority, rationale, scope and evidence.
- **STATE-REQ-138:** Repeated overrides shall trigger governance review.

## 15. Rollback

Rollback restores a previously approved and validated state through a new controlled action. It does not erase the failed, withdrawn or superseded version.

A rollback record shall include:

- rollback identifier;
- initiating event;
- affected artifacts and versions;
- prior and restored states;
- security impact;
- privacy and rights impact;
- implementation impact;
- evidence-preservation actions;
- migration or data impact;
- responsible roles;
- exact restored reference;
- validation result;
- approval.

- **STATE-REQ-139:** Rollback shall not rewrite repository history.
- **STATE-REQ-140:** Rollback shall not restore revoked credentials or expired authority.
- **STATE-REQ-141:** Rollback shall preserve evidence and audit records.
- **STATE-REQ-142:** Restored artifacts shall be revalidated where affected assumptions or environments changed.
- **STATE-REQ-143:** Rollback shall not create an unsupported Approved, Implemented or Validated state.
- **STATE-REQ-144:** Failed rollback shall generate an incident and unresolved finding.
- **STATE-REQ-145:** Emergency rollback shall be narrow, attributable and retrospectively reviewed.
- **STATE-REQ-146:** Rollback behavior shall be tested where implementation exists.

## 16. Repository Representation

State may be represented in:

- document metadata;
- governed state indexes;
- Pull Request descriptions;
- review evidence;
- release manifests;
- machine-readable metadata;
- repository badges;
- generated documentation portals.

- **STATE-REQ-147:** Document metadata is the primary human-readable state representation unless a higher authority defines otherwise.
- **STATE-REQ-148:** A generated state display shall identify its authoritative source.
- **STATE-REQ-149:** Branch names shall not determine lifecycle state.
- **STATE-REQ-150:** Pull Request states such as Draft, Open, Closed or Merged shall not be conflated with artifact lifecycle.
- **STATE-REQ-151:** Issue labels shall not establish lifecycle state independently.
- **STATE-REQ-152:** Git tags shall not establish Approved, Validated or Published state independently.
- **STATE-REQ-153:** Badges shall not overstate state or assurance.
- **STATE-REQ-154:** Machine-readable and human-readable state shall agree.
- **STATE-REQ-155:** State-display disagreement shall resolve to the authoritative evidence record and generate correction work.
- **STATE-REQ-156:** State automation shall preserve manual authorized decisions and audit history.
- **STATE-REQ-157:** Public state displays shall not expose restricted information.
- **STATE-REQ-158:** Repository navigation shall identify superseded and archived artifacts accurately.

## 17. Artifact-Class Profiles

### 17.1 Normative Knowledge Documents

Expected states:

`Draft → Under Review → Approved`

`Implemented`, `Validated` and `Published` apply when corresponding implementation, validation and release evidence exist.

### 17.2 ADRs

ADR decision status follows `OBDIA-ADR-001`. Repository lifecycle metadata shall not replace ADR status.

### 17.3 Specifications and Diagrams

Approval and validation depend on governing authority, consistency and implementation relationship.

### 17.4 Implementation

Implementation status follows `OBDIA-IMP-001`; repository lifecycle records reflect governance and validation evidence.

### 17.5 Evidence Records

Evidence records usually progress through creation, review and validated retention; non-applicable lifecycle states require rationale.

### 17.6 Templates

Templates may be Approved and Published without Implemented state when implementation is not applicable.

### 17.7 Research Records

Research records may remain Draft, Under Review, Approved or Published according to purpose and source policy. Approval does not transform findings into normative architecture.

- **STATE-REQ-159:** Artifact-class profiles shall remain subordinate to the Constitution.
- **STATE-REQ-160:** A class profile shall not create unauthorized alternative states.
- **STATE-REQ-161:** ADR status and repository lifecycle state shall not be conflated.
- **STATE-REQ-162:** Feature implementation status and document lifecycle state shall not be conflated.
- **STATE-REQ-163:** Evidence validity and document approval shall not be conflated.
- **STATE-REQ-164:** A published research record shall not become normative by publication alone.
- **STATE-REQ-165:** A template shall not create authority beyond its approved use.
- **STATE-REQ-166:** Class-specific non-applicability shall be documented.

## 18. State Review and Reconciliation

State shall be reviewed when:

- a material artifact change occurs;
- a dependency changes;
- an ADR is superseded;
- implementation diverges;
- validation fails;
- a security incident occurs;
- an exception expires;
- a release is withdrawn;
- a successor is approved;
- retention obligations change;
- a state-display discrepancy is detected.

- **STATE-REQ-167:** State reconciliation shall identify affected artifacts and dependencies.
- **STATE-REQ-168:** Reconciliation shall not silently elevate state.
- **STATE-REQ-169:** State reduction or suspension shall preserve the prior state record and reason.
- **STATE-REQ-170:** Uncertain state shall be represented conservatively.
- **STATE-REQ-171:** Stale state findings shall have owners and target review dates.
- **STATE-REQ-172:** Critical state inaccuracies affecting security, privacy, authority or publication shall block continued use or release.
- **STATE-REQ-173:** Periodic state review cadence shall be proportionate to risk.
- **STATE-REQ-174:** Baseline freeze shall validate state across all 30 Knowledge documents.
- **STATE-REQ-175:** No document 31 or above shall acquire normative state through repository representation.
- **STATE-REQ-176:** Auxiliary records shall remain non-normative unless formally incorporated.

## 19. Traceability

Every material state shall trace to:

- originating authority;
- artifact identifier and version;
- owner;
- review gates;
- approvals;
- implementation;
- validation evidence;
- release record;
- successor or archive record;
- rollback decisions;
- risks and exceptions.

- **STATE-REQ-177:** State traceability shall be bidirectional.
- **STATE-REQ-178:** Broken state traceability affecting authority, security, privacy, evidence or publication is blocking.
- **STATE-REQ-179:** State records shall use immutable identifiers.
- **STATE-REQ-180:** Planned future state shall not be represented as achieved.
- **STATE-REQ-181:** Supersession and archival links shall preserve historical navigation.
- **STATE-REQ-182:** State evidence shall identify exact repository references where applicable.

## 20. Validation Checklist

Before this model or a state transition is approved, confirm:

- [ ] the artifact identifier, version and classification are accurate;
- [ ] the lifecycle uses only constitutional states;
- [ ] `Reviewed` is not used as a normative state;
- [ ] `Tested` is not used as a normative lifecycle state;
- [ ] `Superseded` is included where applicable;
- [ ] owner and transition authority are identified;
- [ ] exact artifact and commit references are recorded;
- [ ] entry criteria are satisfied;
- [ ] exit criteria are satisfied;
- [ ] applicable review gates have evidence;
- [ ] blocking findings are resolved;
- [ ] non-applicable states have rationale and approval;
- [ ] semantic version and state remain distinct;
- [ ] implementation status and lifecycle state remain distinct;
- [ ] repository and PR workflow states are not conflated with lifecycle;
- [ ] rollback preserves failed history and evidence;
- [ ] supersession identifies the approved successor;
- [ ] archival preserves integrity, retention and access controls;
- [ ] state displays agree with authoritative evidence;
- [ ] forward dependency on document 25 is recorded;
- [ ] Project Founder approval exists before this model becomes Approved.

## 21. Limitations

- This model does not replace the Master Documentation Constitution.
- It does not define semantic-version compatibility in full; document 25 remains a forward dependency.
- It does not replace the ADR status model.
- It does not replace feature-level implementation-transparency states.
- It does not prove that repository automation correctly enforces state.
- It does not create legal authority, implementation approval, operational readiness or publication authorization.
- Internal state review does not constitute independent certification.
- Artifact-specific lifecycle profiles may require additional approved policies.

## 22. Change Control

Every material change to this model shall include:

- change identifier;
- rationale;
- affected lifecycle states and artifact classes;
- dependency analysis;
- security impact;
- privacy and fundamental-rights impact;
- implementation and publication impact;
- migration requirements;
- state-record migration;
- validation plan;
- rollback plan;
- approval record;
- semantic version change.

- **STATE-REQ-183:** Editorial corrections use a patch version when meaning is unchanged.
- **STATE-REQ-184:** Backward-compatible substantive additions use a minor version.
- **STATE-REQ-185:** Incompatible lifecycle or normative changes use a major version.
- **STATE-REQ-186:** Changes shall remain consistent with the exact constitutional lifecycle.
- **STATE-REQ-187:** Material repository-state architecture decisions require an ADR where applicable.
- **STATE-REQ-188:** Changes require Project Founder approval.
- **STATE-REQ-189:** Unapproved state-methodology changes during an active phase are prohibited.
- **STATE-REQ-190:** Historical state records shall not be rewritten retroactively.

## 23. Consolidation Record

Version 1.0.0 consolidates the existing Repository State Model without expanding project scope. It:

- retains immutable document identifier `OBDIA-STATE-001`;
- normalizes the authoritative filename to `13_REPOSITORY_STATE_MODEL.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves the original requirement that every artifact declare its current state;
- preserves the original requirement to identify a responsible reviewer or owner;
- replaces legacy `Reviewed` with constitutional `Under Review`;
- replaces legacy lifecycle `Tested` with constitutional `Validated`;
- adds the constitutionally required `Superseded` state;
- defines state entry and exit criteria, responsible roles, evidence, permitted changes, transition guards, non-applicability, rollback, repository representation, artifact profiles, reconciliation, traceability and archival;
- identifies `25_DOCUMENT_VERSIONING_POLICY.md` as a forward dependency that blocks approval until reconciled;
- treats the former `_ENTERPRISE` and non-Enterprise files as legacy source variants of the same immutable document;
- creates no new normative document and authorizes no implementation or release.

## 24. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Documentation Authority; approval reserved to Project Founder | Constitutional consolidation of the original state model: retained artifact-state declaration and responsible-reviewer intent; corrected lifecycle conflicts; added 190 stable requirements, state criteria, evidence, transitions, guards, rollback, supersession, archival, repository representation and limitations. |
