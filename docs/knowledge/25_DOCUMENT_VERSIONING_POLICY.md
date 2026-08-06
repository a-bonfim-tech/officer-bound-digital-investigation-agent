# DOCUMENT VERSIONING POLICY

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-VER-001 |
| **Title** | Document Versioning Policy |
| **Version** | 1.0.1 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define mandatory semantic versioning, lifecycle independence, dependency compatibility, coordinated change, baseline manifests, release and artifact versioning, revision history, supersession, rollback, migration, integrity, validation and traceability requirements for OBDIA governed artifacts. |
| **Scope** | Normative and informative documents, Knowledge Pack 01–30, Constitution, ADRs, diagrams, policies, schemas, policy bundles, source code, implementations, models, prompts, connectors, tests, evidence packages, repository baselines, releases, tags and archived versions. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `15_RELEASE_AND_PUBLICATION_POLICY.md`; `18_EVIDENCE_MODEL.md`; `22_TESTING_STANDARD.md`; `23_DIAGRAM_STANDARD.md`; `24_GITHUB_REPOSITORY_STANDARD.md`; forward dependencies `26_AI_RISK_REGISTER.md`, `27_RESEARCH_BACKLOG.md`, `28_ASSUMPTIONS_REGISTER.md`, `29_DECISION_LOG_POLICY.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-REL-001`; `OBDIA-EVID-001`; `OBDIA-TEST-001`; `OBDIA-DIAG-001`; `OBDIA-GH-001` |
| **Cross-References** | Documents 26–30; accepted ADRs; document metadata; revision histories; dependency declarations; baseline manifests; release manifests; repository tags; validation packages; risk, exception and change records |
| **Assumptions** | OBDIA artifacts evolve through controlled, attributable and reviewable changes while preserving immutable identifiers and historical versions. |
| **Constraints** | Version numbers shall not create approval, implementation, validation, publication, compliance, legal authority or operational authority; historical versions shall not be silently rewritten; canonical prohibitions shall not be weakened through versioning. |
| **Security Considerations** | Version confusion can cause stale-policy use, mixed incompatible baselines, rollback to vulnerable artifacts, substitution of unreviewed content, loss of provenance, incorrect authorization, evidence corruption, misleading release claims and audit failure. |
| **Validation Criteria** | Every governed artifact has a valid version, accurate lifecycle status, attributable revision history, compatible dependency declarations, integrity reference, change classification, migration and rollback analysis where material, and bidirectional traceability to exact repository commits and approval evidence. |
| **Implementation Relationship** | This policy defines version semantics and records. It does not approve any artifact, create a release, select a versioning tool, configure repository protections or establish production readiness. |

---

## 1. Purpose

This policy consolidates the two legacy Document Versioning Policy variants.

The legacy sources required every document to define:

- Document ID;
- Version;
- Status;
- Author;
- Dependencies;
- Revision History;
- Authoritative source;
- Change history.

They also stated that semantic versioning was recommended and that major revisions required impact assessment and cross-reference review.

The Constitution supersedes the recommendation and makes semantic versioning mandatory. The authoritative core format is:

`MAJOR.MINOR.PATCH`

- **Major** — incompatible normative or architectural change;
- **Minor** — backward-compatible substantive addition;
- **Patch** — correction, clarification or non-material improvement.

Pre-release labels may be used for drafts but do not confer approval. Interdependent documents do not need identical versions, but they shall declare compatibility, be revised together when required and shall not form a knowingly incompatible baseline.

A baseline version is separate from each member document version. The baseline manifest shall record the exact approved version and integrity reference of every member.


## 2. Fundamental Versioning Rules

- **VER-REQ-001:** Every governed document shall define a version.
- **VER-REQ-002:** Every governed document shall define a lifecycle status.
- **VER-REQ-003:** Every governed document shall identify its immutable Document ID.
- **VER-REQ-004:** Every governed document shall identify its author, contributor or attributable origin.
- **VER-REQ-005:** Every governed document shall identify its authoritative source.
- **VER-REQ-006:** Every governed document shall identify material dependencies.
- **VER-REQ-007:** Every governed document shall maintain revision history.
- **VER-REQ-008:** Every governed document shall maintain change history.
- **VER-REQ-009:** Governed artifacts shall use semantic versioning.
- **VER-REQ-010:** Semantic versioning shall be mandatory rather than advisory.
- **VER-REQ-011:** Core versions shall use `MAJOR.MINOR.PATCH`.
- **VER-REQ-012:** Version and lifecycle status shall remain independent.
- **VER-REQ-013:** Version and classification shall remain independent.
- **VER-REQ-014:** Version and implementation status shall remain independent.
- **VER-REQ-015:** Version and validation status shall remain independent.
- **VER-REQ-016:** Version and publication status shall remain independent.
- **VER-REQ-017:** Version numbers shall not create approval.
- **VER-REQ-018:** Version numbers shall not create legal authority.
- **VER-REQ-019:** Version numbers shall not create institutional authority.
- **VER-REQ-020:** Version numbers shall not create implementation evidence.
- **VER-REQ-021:** Version numbers shall not create validation evidence.
- **VER-REQ-022:** Version numbers shall not create compliance evidence.
- **VER-REQ-023:** Historical versions shall not be silently rewritten.
- **VER-REQ-024:** Published versions shall not be silently replaced.
- **VER-REQ-025:** Approved versions shall have integrity references.
- **VER-REQ-026:** Every material version change shall be attributable.
- **VER-REQ-027:** Every material version change shall have rationale.
- **VER-REQ-028:** Every material version change shall preserve traceability.
- **VER-REQ-029:** Version identifiers shall not be reused.
- **VER-REQ-030:** Unknown or invalid version state shall block authoritative use.

## 3. Versioned Artifact Classes

- **VER-REQ-031:** Normative documents shall have document versions.
- **VER-REQ-032:** Informative documents shall have document versions when governed.
- **VER-REQ-033:** The Constitution shall have a constitutional document version.
- **VER-REQ-034:** The Canonical Project Definition shall follow its own superior change authority.
- **VER-REQ-035:** ADRs shall have record versions where their governing policy permits controlled correction metadata.
- **VER-REQ-036:** Diagrams shall have diagram versions.
- **VER-REQ-037:** Schemas shall have schema versions.
- **VER-REQ-038:** Policy bundles shall have policy-bundle versions.
- **VER-REQ-039:** Authorization policies shall have policy versions.
- **VER-REQ-040:** Agent lifecycle definitions shall have lifecycle-model versions.
- **VER-REQ-041:** Connector manifests shall have connector-manifest versions.
- **VER-REQ-042:** Connector implementations shall have implementation versions.
- **VER-REQ-043:** Source-code packages shall have implementation or package versions where released.
- **VER-REQ-044:** Model configurations shall have model-configuration versions.
- **VER-REQ-045:** Prompt templates shall have prompt-template versions where material.
- **VER-REQ-046:** Test plans shall have test-plan versions.
- **VER-REQ-047:** Test cases and suites shall have test-definition versions.
- **VER-REQ-048:** Test evidence packages shall have evidence-package versions.
- **VER-REQ-049:** Validation decisions shall identify the exact versions evaluated.
- **VER-REQ-050:** Repository integrated states shall have commit identifiers.
- **VER-REQ-051:** Repository baselines shall have baseline versions.
- **VER-REQ-052:** Release packages shall have release versions.
- **VER-REQ-053:** Export packages shall have package versions where format evolution matters.
- **VER-REQ-054:** Machine-readable records shall have schema versions.
- **VER-REQ-055:** Generated artifacts shall identify source and generator versions.
- **VER-REQ-056:** Archived artifacts shall retain their historical versions.
- **VER-REQ-057:** Different artifact classes shall not share a version namespace implicitly.
- **VER-REQ-058:** Document version shall not be confused with baseline version.
- **VER-REQ-059:** Baseline version shall not be confused with release version.
- **VER-REQ-060:** Release version shall not be confused with schema or implementation version.

## 4. Core Version Syntax

- **VER-REQ-061:** The core version syntax shall be `MAJOR.MINOR.PATCH`.
- **VER-REQ-062:** `MAJOR` shall be a non-negative integer.
- **VER-REQ-063:** `MINOR` shall be a non-negative integer.
- **VER-REQ-064:** `PATCH` shall be a non-negative integer.
- **VER-REQ-065:** All three numeric components shall be present.
- **VER-REQ-066:** Version components shall not be omitted.
- **VER-REQ-067:** Version components shall not use negative values.
- **VER-REQ-068:** Version components shall not use decimal fractions.
- **VER-REQ-069:** Version components shall not contain whitespace.
- **VER-REQ-070:** Core versions shall not use a leading `v` in document metadata.
- **VER-REQ-071:** Prose and tags may use a governed `v` prefix where allowed by naming policy.
- **VER-REQ-072:** Core version strings shall be machine-searchable.
- **VER-REQ-073:** Canonical document filenames shall not contain versions.
- **VER-REQ-074:** Normative document titles shall not contain versions.
- **VER-REQ-075:** Display labels shall not replace machine-readable versions.
- **VER-REQ-076:** Invalid version syntax shall be rejected.
- **VER-REQ-077:** Version parsing shall fail safely.
- **VER-REQ-078:** Unknown version components shall not default to zero.
- **VER-REQ-079:** Missing versions shall not be inferred from repository history.
- **VER-REQ-080:** Version comparison shall compare numeric core components.
- **VER-REQ-081:** Version comparison shall consider pre-release status where used.
- **VER-REQ-082:** Build metadata shall not alter semantic precedence.
- **VER-REQ-083:** Version normalization shall not silently change meaning.
- **VER-REQ-084:** Legacy malformed versions shall require explicit migration.
- **VER-REQ-085:** Version syntax changes shall require major-policy review.

## 5. Major Version Semantics

- **VER-REQ-086:** A major increment shall represent an incompatible normative or architectural change.
- **VER-REQ-087:** A major increment shall be required when prior conforming dependencies become non-conforming.
- **VER-REQ-088:** A major increment shall be required for incompatible authority changes.
- **VER-REQ-089:** A major increment shall be required for incompatible scope changes.
- **VER-REQ-090:** A major increment shall be required for incompatible trust-model changes.
- **VER-REQ-091:** A major increment shall be required for incompatible identity-model changes.
- **VER-REQ-092:** A major increment shall be required for incompatible authorization-model changes.
- **VER-REQ-093:** A major increment shall be required for incompatible evidence-model changes.
- **VER-REQ-094:** A major increment shall be required for incompatible privacy-control changes.
- **VER-REQ-095:** A major increment shall be required for incompatible lifecycle changes.
- **VER-REQ-096:** A major increment shall be required for incompatible interface or schema changes.
- **VER-REQ-097:** A major increment shall be required for incompatible classification changes.
- **VER-REQ-098:** A major increment shall be required for incompatible documentation-methodology changes.
- **VER-REQ-099:** A major increment shall be required when a mandatory field is removed or redefined incompatibly.
- **VER-REQ-100:** A major increment shall be required when a requirement is weakened incompatibly.
- **VER-REQ-101:** A major increment shall be required when a prohibited boundary is changed incompatibly.
- **VER-REQ-102:** A major increment shall be required when prior validation evidence is broadly invalidated.
- **VER-REQ-103:** A major increment shall require full impact assessment.
- **VER-REQ-104:** A major increment shall require cross-reference review.
- **VER-REQ-105:** A major increment shall require dependency compatibility review.
- **VER-REQ-106:** A major increment shall require migration planning.
- **VER-REQ-107:** A major increment shall require rollback planning.
- **VER-REQ-108:** A major increment shall require all applicable review gates.
- **VER-REQ-109:** A major normative baseline increment shall require explicit Project Founder approval.
- **VER-REQ-110:** A major architectural change shall require an ADR where applicable.
- **VER-REQ-111:** A major increment shall reset `MINOR` and `PATCH` to zero.
- **VER-REQ-112:** A major increment shall identify superseded versions.
- **VER-REQ-113:** A major increment shall not erase prior versions.
- **VER-REQ-114:** Major-version acceptance shall identify affected implementations and releases.
- **VER-REQ-115:** Major-version release shall not occur while known dependency incompatibility remains unresolved.

## 6. Minor Version Semantics

- **VER-REQ-116:** A minor increment shall represent a backward-compatible substantive addition.
- **VER-REQ-117:** A minor increment may add new requirements that do not invalidate prior conformance.
- **VER-REQ-118:** A minor increment may add new traceability.
- **VER-REQ-119:** A minor increment may strengthen controls without invalidating conforming dependencies.
- **VER-REQ-120:** A minor increment may add optional fields with safe defaults.
- **VER-REQ-121:** A minor increment may add compatible schemas or views.
- **VER-REQ-122:** A minor increment may add operationally valuable clarification.
- **VER-REQ-123:** A minor increment may add new review or validation detail.
- **VER-REQ-124:** A minor increment shall not silently weaken existing controls.
- **VER-REQ-125:** A minor increment shall not remove mandatory fields.
- **VER-REQ-126:** A minor increment shall not redefine existing identifiers incompatibly.
- **VER-REQ-127:** A minor increment shall not broaden authority incompatibly.
- **VER-REQ-128:** A minor increment shall require impact assessment.
- **VER-REQ-129:** A minor increment shall require dependency review.
- **VER-REQ-130:** A minor increment shall require cross-reference review.
- **VER-REQ-131:** A minor increment shall identify affected requirements.
- **VER-REQ-132:** A minor increment shall identify affected tests.
- **VER-REQ-133:** A minor increment shall identify affected implementation.
- **VER-REQ-134:** A minor increment shall identify migration notes when needed.
- **VER-REQ-135:** A minor increment shall identify validation impact.
- **VER-REQ-136:** A minor increment shall identify release impact.
- **VER-REQ-137:** A minor increment shall require applicable specialist review.
- **VER-REQ-138:** A minor increment shall reset `PATCH` to zero.
- **VER-REQ-139:** A minor increment shall preserve `MAJOR`.
- **VER-REQ-140:** A minor increment shall preserve historical versions.
- **VER-REQ-141:** Minor-version additions shall use new immutable requirement identifiers.
- **VER-REQ-142:** Removed draft-only text shall be classified according to semantic impact.
- **VER-REQ-143:** Minor-version compatibility shall be demonstrated rather than assumed.
- **VER-REQ-144:** Minor-version approval shall remain separate from the version increment.
- **VER-REQ-145:** Minor-version publication shall require a release decision.

## 7. Patch Version Semantics

- **VER-REQ-146:** A patch increment shall represent a correction, clarification or non-material improvement.
- **VER-REQ-147:** A patch increment may correct spelling.
- **VER-REQ-148:** A patch increment may correct formatting.
- **VER-REQ-149:** A patch increment may repair broken links.
- **VER-REQ-150:** A patch increment may correct non-semantic metadata errors.
- **VER-REQ-151:** A patch increment may clarify wording without changing requirements.
- **VER-REQ-152:** A patch increment may update revision history for an already authorized correction.
- **VER-REQ-153:** A patch increment may correct a rendering defect.
- **VER-REQ-154:** A patch increment may correct an integrity-reference transcription error through governed evidence.
- **VER-REQ-155:** A patch increment shall not change authority.
- **VER-REQ-156:** A patch increment shall not change scope materially.
- **VER-REQ-157:** A patch increment shall not add incompatible requirements.
- **VER-REQ-158:** A patch increment shall not weaken controls.
- **VER-REQ-159:** A patch increment shall not change lifecycle semantics.
- **VER-REQ-160:** A patch increment shall not change authorization behavior.
- **VER-REQ-161:** A patch increment shall not change evidence semantics.
- **VER-REQ-162:** A patch increment shall not change implementation interfaces incompatibly.
- **VER-REQ-163:** A patch increment shall require confirmation that meaning is unchanged.
- **VER-REQ-164:** A patch increment shall identify the corrected defect.
- **VER-REQ-165:** A patch increment shall identify affected references.
- **VER-REQ-166:** A patch increment shall identify whether revalidation is required.
- **VER-REQ-167:** A patch increment shall identify whether republication is required.
- **VER-REQ-168:** A patch increment shall preserve `MAJOR` and `MINOR`.
- **VER-REQ-169:** A patch increment shall increase `PATCH` by at least one controlled step.
- **VER-REQ-170:** Multiple unrelated corrections may be grouped only with clear rationale.
- **VER-REQ-171:** Patch changes shall preserve historical versions.
- **VER-REQ-172:** Patch changes to published artifacts shall use correction-release controls.
- **VER-REQ-173:** Security patches may require expanded review despite patch numbering.
- **VER-REQ-174:** Patch numbering shall not minimize the severity of the corrected defect.
- **VER-REQ-175:** Patch approval shall remain separate from the version increment.

## 8. Initial Versions and Zero Versions

- **VER-REQ-176:** Initial consolidated normative Knowledge documents may use version `1.0.0` while remaining Draft.
- **VER-REQ-177:** Version `1.0.0` shall not imply approval.
- **VER-REQ-178:** Version `1.0.0` shall not imply publication.
- **VER-REQ-179:** Version `1.0.0` shall not imply implementation.
- **VER-REQ-180:** Version `1.0.0` shall not imply validation.
- **VER-REQ-181:** Version `0.MINOR.PATCH` may be used for experimental or unstable artifacts when explicitly governed.
- **VER-REQ-182:** Zero-major versions shall not be used to evade review.
- **VER-REQ-183:** Zero-major versions shall not be represented as approved normative baselines unless authority explicitly permits it.
- **VER-REQ-184:** Experimental zero-major artifacts shall identify instability.
- **VER-REQ-185:** Experimental zero-major artifacts shall identify compatibility limitations.
- **VER-REQ-186:** Transition from zero-major to `1.0.0` shall require baseline-readiness review.
- **VER-REQ-187:** Initial version selection shall identify rationale.
- **VER-REQ-188:** Imported legacy artifacts shall not receive fabricated historical versions.
- **VER-REQ-189:** Unknown legacy versions shall be recorded as unknown rather than inferred.
- **VER-REQ-190:** Consolidated versions shall identify the legacy sources incorporated.
- **VER-REQ-191:** Consolidation shall not erase source provenance.
- **VER-REQ-192:** Initial version assignment shall preserve immutable Document IDs.
- **VER-REQ-193:** Initial version assignment shall preserve historical source records.
- **VER-REQ-194:** Initial version assignment shall identify status independently.
- **VER-REQ-195:** Initial version assignment shall identify approval evidence separately.

## 9. Pre-Release Labels

- **VER-REQ-196:** Pre-release labels may be used for drafts.
- **VER-REQ-197:** Pre-release labels shall not confer approval.
- **VER-REQ-198:** Pre-release labels shall follow the project semantic-version profile.
- **VER-REQ-199:** Pre-release labels shall be appended after a hyphen.
- **VER-REQ-200:** Pre-release identifiers shall use controlled alphanumeric and hyphen tokens.
- **VER-REQ-201:** Dot-separated pre-release identifiers may be used.
- **VER-REQ-202:** Pre-release labels shall not contain whitespace.
- **VER-REQ-203:** Pre-release labels shall not contain secrets.
- **VER-REQ-204:** Pre-release labels shall not contain personal or case identifiers.
- **VER-REQ-205:** `draft` may be used as a pre-release label when needed.
- **VER-REQ-206:** `rc` may be used for a release candidate when governed by release policy.
- **VER-REQ-207:** `alpha` and `beta` may be used only when their project meaning is documented.
- **VER-REQ-208:** Pre-release sequence identifiers shall be monotonically controlled within the same core version.
- **VER-REQ-209:** A later pre-release shall not reuse an earlier complete version string.
- **VER-REQ-210:** Pre-release labels shall not replace lifecycle status.
- **VER-REQ-211:** A Draft document may use a stable core version without a pre-release label.
- **VER-REQ-212:** An Approved document shall not rely on a pre-release label to express approval.
- **VER-REQ-213:** Release-candidate labels shall identify exact candidate artifacts.
- **VER-REQ-214:** Release-candidate rejection shall preserve history.
- **VER-REQ-215:** Pre-release promotion shall identify the source candidate.
- **VER-REQ-216:** Pre-release changes shall receive semantic-impact classification.
- **VER-REQ-217:** Pre-release artifacts shall not be represented as published stable releases.
- **VER-REQ-218:** Pre-release dependencies shall be declared explicitly.
- **VER-REQ-219:** Mixed stable and pre-release baselines shall require compatibility review.
- **VER-REQ-220:** Pre-release label conventions shall be machine-validated.

## 10. Build Metadata

- **VER-REQ-221:** Build metadata may be appended after a plus sign where useful.
- **VER-REQ-222:** Build metadata shall identify non-semantic build information.
- **VER-REQ-223:** Build metadata shall not alter semantic version precedence.
- **VER-REQ-224:** Build metadata shall not replace repository commit identifiers.
- **VER-REQ-225:** Build metadata shall not replace integrity references.
- **VER-REQ-226:** Build metadata shall not create approval.
- **VER-REQ-227:** Build metadata shall not create validation.
- **VER-REQ-228:** Build metadata shall not contain secrets.
- **VER-REQ-229:** Build metadata shall not contain personal or case identifiers.
- **VER-REQ-230:** Build metadata may identify a shortened commit only when the full commit remains available elsewhere.
- **VER-REQ-231:** Build metadata may identify a controlled build sequence.
- **VER-REQ-232:** Build metadata may identify a platform or environment class where governed.
- **VER-REQ-233:** Build metadata shall be reproducible or attributable.
- **VER-REQ-234:** Build metadata collisions shall be avoided.
- **VER-REQ-235:** Build metadata shall not be used as the sole release identifier.
- **VER-REQ-236:** Published artifact manifests shall identify full source and integrity references.
- **VER-REQ-237:** Generated build metadata shall identify the generator.
- **VER-REQ-238:** Build metadata format shall be documented.
- **VER-REQ-239:** Unknown build metadata shall not broaden compatibility assumptions.
- **VER-REQ-240:** Build metadata shall be retained with the artifact record where material.

## 11. Status and Version Independence

- **VER-REQ-241:** Lifecycle status and semantic version shall be stored as separate fields.
- **VER-REQ-242:** Draft shall not be encoded solely through the version.
- **VER-REQ-243:** Under Review shall not be encoded solely through the version.
- **VER-REQ-244:** Approved shall not be encoded solely through the version.
- **VER-REQ-245:** Implemented shall not be encoded solely through the version.
- **VER-REQ-246:** Validated shall not be encoded solely through the version.
- **VER-REQ-247:** Published shall not be encoded solely through the version.
- **VER-REQ-248:** Superseded shall not be encoded solely through the version.
- **VER-REQ-249:** Archived shall not be encoded solely through the version.
- **VER-REQ-250:** Version `1.0.0` may have Draft status.
- **VER-REQ-251:** Multiple lifecycle transitions may occur without a version increment when content is unchanged.
- **VER-REQ-252:** Content changes shall receive semantic-impact assessment even when lifecycle status is unchanged.
- **VER-REQ-253:** Lifecycle transition evidence shall identify the exact version.
- **VER-REQ-254:** Approval evidence shall identify the exact version and integrity reference.
- **VER-REQ-255:** Implementation evidence shall identify the exact approved version implemented.
- **VER-REQ-256:** Validation evidence shall identify the exact implementation and governing versions.
- **VER-REQ-257:** Publication evidence shall identify the exact approved artifacts released.
- **VER-REQ-258:** Supersession records shall identify old and successor versions.
- **VER-REQ-259:** Archive records shall identify retained versions.
- **VER-REQ-260:** Status labels in repository interfaces shall not override document metadata.
- **VER-REQ-261:** PR state shall not override lifecycle status.
- **VER-REQ-262:** Branch name shall not override lifecycle status.
- **VER-REQ-263:** Tag existence shall not override lifecycle status.
- **VER-REQ-264:** Release existence shall not override lifecycle status.
- **VER-REQ-265:** Status-version conflicts shall block progression.

## 12. Mandatory Version Metadata

- **VER-REQ-266:** Document metadata shall include Document ID.
- **VER-REQ-267:** Document metadata shall include Title.
- **VER-REQ-268:** Document metadata shall include Version.
- **VER-REQ-269:** Document metadata shall include Status.
- **VER-REQ-270:** Document metadata shall include Classification.
- **VER-REQ-271:** Document metadata shall include Authority or owner.
- **VER-REQ-272:** Document metadata shall identify author or contributor attribution.
- **VER-REQ-273:** Document metadata shall include Purpose.
- **VER-REQ-274:** Document metadata shall include Scope.
- **VER-REQ-275:** Document metadata shall include Dependencies.
- **VER-REQ-276:** Document metadata shall include Normative references.
- **VER-REQ-277:** Document metadata shall include Cross-references.
- **VER-REQ-278:** Document metadata shall include Assumptions.
- **VER-REQ-279:** Document metadata shall include Constraints.
- **VER-REQ-280:** Document metadata shall include Security considerations.
- **VER-REQ-281:** Document metadata shall include Validation criteria.
- **VER-REQ-282:** Document metadata shall include Revision history.
- **VER-REQ-283:** Document metadata shall identify authoritative source.
- **VER-REQ-284:** Document metadata shall identify supersession where applicable.
- **VER-REQ-285:** Document metadata shall identify approval authority.
- **VER-REQ-286:** Document metadata shall identify effective date where applicable.
- **VER-REQ-287:** Document metadata shall be internally consistent.
- **VER-REQ-288:** Document metadata shall be machine-searchable.
- **VER-REQ-289:** Metadata versions shall match revision-history current entries.
- **VER-REQ-290:** Metadata and filename shall not conflict.
- **VER-REQ-291:** Metadata changes shall receive semantic-impact assessment.
- **VER-REQ-292:** Machine-readable sidecars may supply metadata only through unambiguous linkage.
- **VER-REQ-293:** Short records may use compact version metadata only when their governing schema permits it.
- **VER-REQ-294:** Generated evidence may use authenticated envelopes for version metadata.
- **VER-REQ-295:** Missing mandatory version metadata shall block approval.

## 13. Author and Authoritative Source

- **VER-REQ-296:** Every governed version shall have attributable authorship or contribution origin.
- **VER-REQ-297:** Author attribution shall not confer approval authority.
- **VER-REQ-298:** Author attribution shall not confer legal authority.
- **VER-REQ-299:** AI assistance shall not replace accountable human authorship.
- **VER-REQ-300:** AI-generated content shall identify accountable human review.
- **VER-REQ-301:** Machine-generated artifacts shall identify the generating system or workflow.
- **VER-REQ-302:** Authoritative source shall identify the controlling repository path or governed record.
- **VER-REQ-303:** Rendered artifacts shall identify their authoritative source.
- **VER-REQ-304:** Generated indexes shall identify their authoritative metadata source.
- **VER-REQ-305:** Released binaries shall identify source commit and build evidence.
- **VER-REQ-306:** Mirrors shall not become authoritative without governance approval.
- **VER-REQ-307:** Copies shall not become authoritative through recency alone.
- **VER-REQ-308:** Downloaded filenames shall not define authoritative source.
- **VER-REQ-309:** Local edits shall not become authoritative until governed incorporation.
- **VER-REQ-310:** Authoritative-source changes shall require migration and cross-reference updates.
- **VER-REQ-311:** Authoritative-source conflicts shall resolve according to the hierarchy.
- **VER-REQ-312:** Unknown authorship shall remain explicit.
- **VER-REQ-313:** Unknown authoritative source shall block approval.
- **VER-REQ-314:** Authorship corrections shall preserve historical attribution.
- **VER-REQ-315:** Author and source metadata shall not contain unnecessary personal data.

## 14. Revision History

- **VER-REQ-316:** Every governed document shall maintain revision history.
- **VER-REQ-317:** Revision history shall identify version.
- **VER-REQ-318:** Revision history shall identify date or governed time reference.
- **VER-REQ-319:** Revision history shall identify lifecycle status at the recorded event.
- **VER-REQ-320:** Revision history shall identify author, authority or owner as applicable.
- **VER-REQ-321:** Revision history shall summarize the change.
- **VER-REQ-322:** Revision history shall identify change classification.
- **VER-REQ-323:** Revision history shall identify approval reference where applicable.
- **VER-REQ-324:** Revision history shall identify repository commit where material.
- **VER-REQ-325:** Revision history shall identify integrity reference where material.
- **VER-REQ-326:** Revision history shall identify superseded version where applicable.
- **VER-REQ-327:** Revision history shall identify rollback where applicable.
- **VER-REQ-328:** Revision history shall identify migration where applicable.
- **VER-REQ-329:** Revision history shall not be rewritten to conceal prior states.
- **VER-REQ-330:** Revision history corrections shall be additive.
- **VER-REQ-331:** Revision history shall not fabricate dates.
- **VER-REQ-332:** Revision history shall not fabricate approvals.
- **VER-REQ-333:** Revision history shall not infer unknown source versions.
- **VER-REQ-334:** Consolidation entries shall identify incorporated legacy variants.
- **VER-REQ-335:** Published-version history shall remain immutable.
- **VER-REQ-336:** Archived-version history shall remain retrievable.
- **VER-REQ-337:** Revision-history tables shall use consistent columns.
- **VER-REQ-338:** Revision history shall remain readable in plain Markdown.
- **VER-REQ-339:** Revision-history entries shall trace to change records.
- **VER-REQ-340:** Current metadata shall match the latest valid revision-history entry.

## 15. Change History and Change Records

- **VER-REQ-341:** Material version changes shall have change records.
- **VER-REQ-342:** Change records shall have immutable identifiers.
- **VER-REQ-343:** Change records shall identify affected artifacts.
- **VER-REQ-344:** Change records shall identify prior versions.
- **VER-REQ-345:** Change records shall identify proposed versions.
- **VER-REQ-346:** Change records shall identify semantic-impact classification.
- **VER-REQ-347:** Change records shall identify rationale.
- **VER-REQ-348:** Change records shall identify affected requirements.
- **VER-REQ-349:** Change records shall identify affected dependencies.
- **VER-REQ-350:** Change records shall identify security impact.
- **VER-REQ-351:** Change records shall identify privacy impact.
- **VER-REQ-352:** Change records shall identify evidence impact.
- **VER-REQ-353:** Change records shall identify implementation impact.
- **VER-REQ-354:** Change records shall identify testing impact.
- **VER-REQ-355:** Change records shall identify release impact.
- **VER-REQ-356:** Change records shall identify migration.
- **VER-REQ-357:** Change records shall identify rollback.
- **VER-REQ-358:** Change records shall identify reviewers.
- **VER-REQ-359:** Change records shall identify approval authority.
- **VER-REQ-360:** Change records shall identify exact commits.
- **VER-REQ-361:** Change records shall identify integrity references.
- **VER-REQ-362:** Change records shall identify unresolved findings.
- **VER-REQ-363:** Change records shall identify residual risk.
- **VER-REQ-364:** Change records shall identify effective time.
- **VER-REQ-365:** Change records shall remain retrievable.

## 16. Dependency Version Declarations

- **VER-REQ-366:** Every material dependency shall identify an exact version or compatibility constraint.
- **VER-REQ-367:** Dependencies shall use immutable document or artifact identifiers.
- **VER-REQ-368:** Dependency declarations shall distinguish exact from minimum-compatible versions.
- **VER-REQ-369:** Dependency declarations shall distinguish required from optional dependencies.
- **VER-REQ-370:** Dependency declarations shall distinguish normative from informative dependencies.
- **VER-REQ-371:** Dependency declarations shall identify pre-release dependencies explicitly.
- **VER-REQ-372:** Dependency declarations shall identify superseded dependencies.
- **VER-REQ-373:** Dependency declarations shall identify unresolved forward dependencies.
- **VER-REQ-374:** Dependency declarations shall not rely solely on filenames.
- **VER-REQ-375:** Dependency declarations shall not rely solely on repository branch names.
- **VER-REQ-376:** Dependency declarations shall not infer compatibility from identical version numbers.
- **VER-REQ-377:** Interdependent documents need not have identical versions.
- **VER-REQ-378:** Compatibility shall be assessed by requirements and interfaces.
- **VER-REQ-379:** Material dependency changes shall trigger impact review.
- **VER-REQ-380:** Removed dependencies shall have rationale.
- **VER-REQ-381:** Added dependencies shall receive security and governance review where material.
- **VER-REQ-382:** Dependency cycles shall be identified.
- **VER-REQ-383:** Unresolvable dependency cycles shall block approval.
- **VER-REQ-384:** Unknown dependency versions shall block an approved baseline.
- **VER-REQ-385:** Dependency constraints shall be machine-searchable where practical.
- **VER-REQ-386:** Dependency declarations shall remain synchronized with cross-references.
- **VER-REQ-387:** Dependency declarations shall identify the authoritative source.
- **VER-REQ-388:** Dependency-version drift shall be detectable.
- **VER-REQ-389:** Dependency compatibility tests shall identify exact versions.
- **VER-REQ-390:** Dependency history shall remain retrievable.

## 17. Compatibility Semantics

- **VER-REQ-391:** Compatibility shall mean conformance without invalidating applicable requirements.
- **VER-REQ-392:** Backward compatibility shall be defined for the affected artifact class.
- **VER-REQ-393:** Forward compatibility shall not be assumed.
- **VER-REQ-394:** Document compatibility shall consider normative meaning.
- **VER-REQ-395:** Schema compatibility shall consider required and optional fields.
- **VER-REQ-396:** Policy compatibility shall consider decision semantics and obligations.
- **VER-REQ-397:** Implementation compatibility shall consider interfaces and expected behavior.
- **VER-REQ-398:** Evidence compatibility shall consider identifiers, integrity, provenance and custody.
- **VER-REQ-399:** Test compatibility shall consider fixtures, environments and expected results.
- **VER-REQ-400:** Release compatibility shall consider installation, migration and public claims.
- **VER-REQ-401:** Compatibility shall consider security controls.
- **VER-REQ-402:** Compatibility shall consider privacy controls.
- **VER-REQ-403:** Compatibility shall consider authority and accountability.
- **VER-REQ-404:** Compatibility shall consider case and purpose isolation.
- **VER-REQ-405:** Compatibility shall consider lifecycle semantics.
- **VER-REQ-406:** Compatibility shall consider deprecation and removal.
- **VER-REQ-407:** Compatibility claims shall identify tested or reviewed scope.
- **VER-REQ-408:** Compatibility claims shall identify limitations.
- **VER-REQ-409:** Compatibility shall not be inferred from successful parsing alone.
- **VER-REQ-410:** Compatibility shall not be inferred from identical titles.
- **VER-REQ-411:** Compatibility shall not be inferred from identical major versions without review.
- **VER-REQ-412:** Known incompatible versions shall not form an approved baseline.
- **VER-REQ-413:** Unknown compatibility shall be treated as unresolved.
- **VER-REQ-414:** Compatibility conflicts shall identify owners.
- **VER-REQ-415:** Compatibility evidence shall remain traceable.

## 18. Coordinated Version Changes

- **VER-REQ-416:** Interdependent artifacts shall be reviewed together when one change makes another incomplete.
- **VER-REQ-417:** Interdependent artifacts shall be reviewed together when one change makes another inconsistent.
- **VER-REQ-418:** Interdependent artifacts shall be reviewed together when one change makes another incompatible.
- **VER-REQ-419:** Compatible coordinated changes shall use one controlled change set when required.
- **VER-REQ-420:** Coordinated changes shall identify every affected artifact.
- **VER-REQ-421:** Coordinated changes shall identify every prior and successor version.
- **VER-REQ-422:** Coordinated changes shall identify merge and release ordering.
- **VER-REQ-423:** Coordinated changes shall identify temporary mixed-state risks.
- **VER-REQ-424:** Coordinated changes shall prevent release of known incompatible mixtures.
- **VER-REQ-425:** Coordinated changes shall update cross-references before approval.
- **VER-REQ-426:** Coordinated changes shall update traceability before approval.
- **VER-REQ-427:** Coordinated changes shall update indexes and manifests.
- **VER-REQ-428:** Coordinated changes shall update tests.
- **VER-REQ-429:** Coordinated changes shall update diagrams where affected.
- **VER-REQ-430:** Coordinated changes shall update implementation where affected.
- **VER-REQ-431:** Coordinated changes shall update release notes where affected.
- **VER-REQ-432:** Coordinated changes shall identify rollback order.
- **VER-REQ-433:** Coordinated changes shall identify partial-failure behavior.
- **VER-REQ-434:** Coordinated changes shall identify authority for each artifact.
- **VER-REQ-435:** Coordinated changes shall preserve distinct document versions.
- **VER-REQ-436:** Coordinated changes shall not force identical version numbers without need.
- **VER-REQ-437:** Coordinated changes shall preserve individual revision histories.
- **VER-REQ-438:** Coordinated changes shall preserve integrity references.
- **VER-REQ-439:** Unresolved coordinated-change failure shall block baseline freeze.
- **VER-REQ-440:** Coordinated-change evidence shall remain retrievable.

## 19. Baseline Versions

- **VER-REQ-441:** A baseline version shall be separate from each member document version.
- **VER-REQ-442:** A baseline shall represent a governed set of compatible artifact versions.
- **VER-REQ-443:** A baseline shall have an immutable baseline identifier.
- **VER-REQ-444:** A baseline shall have a semantic version.
- **VER-REQ-445:** A baseline shall have a lifecycle status.
- **VER-REQ-446:** A baseline shall identify authority and owner.
- **VER-REQ-447:** A baseline shall identify scope.
- **VER-REQ-448:** A baseline shall identify every member artifact.
- **VER-REQ-449:** A baseline shall identify each member version.
- **VER-REQ-450:** A baseline shall identify each member integrity reference.
- **VER-REQ-451:** A baseline shall identify each member authoritative path.
- **VER-REQ-452:** A baseline shall identify dependency compatibility.
- **VER-REQ-453:** A baseline shall identify unresolved exclusions.
- **VER-REQ-454:** A baseline shall identify approval evidence.
- **VER-REQ-455:** A baseline shall identify repository commit.
- **VER-REQ-456:** A baseline shall identify creation or freeze time.
- **VER-REQ-457:** A baseline shall identify release relationship where applicable.
- **VER-REQ-458:** A baseline shall identify superseded baseline where applicable.
- **VER-REQ-459:** A baseline shall not be inferred from a branch.
- **VER-REQ-460:** A baseline shall not be inferred from a tag alone.
- **VER-REQ-461:** A baseline shall not be inferred from repository recency.
- **VER-REQ-462:** A baseline shall not contain known incompatible versions.
- **VER-REQ-463:** A baseline shall not omit blocking forward dependencies silently.
- **VER-REQ-464:** A baseline version increment shall reflect baseline-level semantic impact.
- **VER-REQ-465:** Baseline history shall remain immutable.

## 20. Baseline Manifest

- **VER-REQ-466:** Every approved baseline shall have a baseline manifest.
- **VER-REQ-467:** The baseline manifest shall have an immutable identifier.
- **VER-REQ-468:** The baseline manifest shall identify baseline version.
- **VER-REQ-469:** The baseline manifest shall identify baseline status.
- **VER-REQ-470:** The baseline manifest shall identify authority.
- **VER-REQ-471:** The baseline manifest shall identify repository commit.
- **VER-REQ-472:** The baseline manifest shall list every member identifier.
- **VER-REQ-473:** The baseline manifest shall list every member version.
- **VER-REQ-474:** The baseline manifest shall list every member integrity reference.
- **VER-REQ-475:** The baseline manifest shall list every member authoritative path.
- **VER-REQ-476:** The baseline manifest shall identify classification.
- **VER-REQ-477:** The baseline manifest shall identify dependencies.
- **VER-REQ-478:** The baseline manifest shall identify compatibility findings.
- **VER-REQ-479:** The baseline manifest shall identify review records.
- **VER-REQ-480:** The baseline manifest shall identify approval record.
- **VER-REQ-481:** The baseline manifest shall identify exclusions and rationale.
- **VER-REQ-482:** The baseline manifest shall identify forward dependencies.
- **VER-REQ-483:** The baseline manifest shall identify known limitations.
- **VER-REQ-484:** The baseline manifest shall identify supersession.
- **VER-REQ-485:** The baseline manifest shall identify release linkage where applicable.
- **VER-REQ-486:** The baseline manifest shall be machine-readable or machine-searchable.
- **VER-REQ-487:** The baseline manifest shall be integrity-protected.
- **VER-REQ-488:** The baseline manifest shall not contain secrets.
- **VER-REQ-489:** The baseline manifest shall minimize personal and case data.
- **VER-REQ-490:** Manifest correction shall preserve prior versions.

## 21. Baseline Freeze

- **VER-REQ-491:** Baseline freeze shall require complete manifest evidence.
- **VER-REQ-492:** Baseline freeze shall require compatibility review.
- **VER-REQ-493:** Baseline freeze shall require cross-reference validation.
- **VER-REQ-494:** Baseline freeze shall require requirement-ID validation.
- **VER-REQ-495:** Baseline freeze shall require version validation.
- **VER-REQ-496:** Baseline freeze shall require status validation.
- **VER-REQ-497:** Baseline freeze shall require integrity verification.
- **VER-REQ-498:** Baseline freeze shall require review-gate completion.
- **VER-REQ-499:** Baseline freeze shall require explicit Project Founder approval for the normative v1.0 baseline.
- **VER-REQ-500:** Baseline freeze shall identify exact repository commit.
- **VER-REQ-501:** Baseline freeze shall identify exact member versions.
- **VER-REQ-502:** Baseline freeze shall identify exact member hashes or equivalent integrity references.
- **VER-REQ-503:** Baseline freeze shall identify unresolved non-blocking limitations.
- **VER-REQ-504:** Baseline freeze shall not proceed with canonical conflicts.
- **VER-REQ-505:** Baseline freeze shall not proceed with known incompatible dependencies.
- **VER-REQ-506:** Baseline freeze shall not proceed with unresolved critical security findings.
- **VER-REQ-507:** Baseline freeze shall not proceed with unresolved critical privacy findings.
- **VER-REQ-508:** Baseline freeze shall not proceed with fabricated or missing approval evidence.
- **VER-REQ-509:** Freeze shall not convert implementation status automatically.
- **VER-REQ-510:** Freeze shall not convert validation status automatically.
- **VER-REQ-511:** Freeze shall not publish artifacts automatically.
- **VER-REQ-512:** Post-freeze material changes shall require formal change control.
- **VER-REQ-513:** Post-freeze material changes shall increment affected versions.
- **VER-REQ-514:** Freeze records shall remain immutable.
- **VER-REQ-515:** Freeze reversal shall use supersession or rollback records rather than deletion.

## 22. Release Versions

- **VER-REQ-516:** Release versions shall be separate from document versions.
- **VER-REQ-517:** Release versions shall identify exact included artifact versions.
- **VER-REQ-518:** Release versions shall identify exact source commit.
- **VER-REQ-519:** Release versions shall identify release status.
- **VER-REQ-520:** Release versions shall identify prerelease status where applicable.
- **VER-REQ-521:** Release versions shall identify integrity manifest.
- **VER-REQ-522:** Release versions shall identify test and validation evidence.
- **VER-REQ-523:** Release versions shall identify known limitations.
- **VER-REQ-524:** Release versions shall identify compatibility.
- **VER-REQ-525:** Release versions shall identify migration.
- **VER-REQ-526:** Release versions shall identify rollback or withdrawal.
- **VER-REQ-527:** Release versions shall identify public claims.
- **VER-REQ-528:** Release versions shall not imply production readiness.
- **VER-REQ-529:** Release versions shall not imply certification.
- **VER-REQ-530:** Release versions shall not imply compliance.
- **VER-REQ-531:** Release versions shall not imply institutional adoption.
- **VER-REQ-532:** Draft releases shall remain distinguishable.
- **VER-REQ-533:** Release candidates shall remain distinguishable.
- **VER-REQ-534:** Published releases shall be immutable.
- **VER-REQ-535:** Published release correction shall use a new version.
- **VER-REQ-536:** Withdrawn releases shall retain version and history.
- **VER-REQ-537:** Release tags shall trace to release versions.
- **VER-REQ-538:** Release automation shall not create approval.
- **VER-REQ-539:** Release versions shall conform to `OBDIA-REL-001`.
- **VER-REQ-540:** Release history shall remain retrievable.

## 23. Repository Tags

- **VER-REQ-541:** Tags shall identify exact commits.
- **VER-REQ-542:** Release tags shall use governed version formats.
- **VER-REQ-543:** Tags shall not be reused.
- **VER-REQ-544:** Published tags shall not be moved silently.
- **VER-REQ-545:** Tag deletion shall require governance review.
- **VER-REQ-546:** Incorrect tags shall be corrected through documented procedures.
- **VER-REQ-547:** Tag correction shall preserve prior evidence.
- **VER-REQ-548:** Tag creation shall be attributable.
- **VER-REQ-549:** Tag signatures may supplement but shall not replace review.
- **VER-REQ-550:** Tag presence shall not create approval.
- **VER-REQ-551:** Tag presence shall not create validation.
- **VER-REQ-552:** Tag presence shall not create publication without release evidence.
- **VER-REQ-553:** Document versions shall not require one tag per draft revision unless governance chooses it.
- **VER-REQ-554:** Baseline tags shall identify baseline manifests.
- **VER-REQ-555:** Release tags shall identify release manifests.
- **VER-REQ-556:** Pre-release tags shall be distinguishable.
- **VER-REQ-557:** Experimental tags shall not imply support.
- **VER-REQ-558:** Tag names shall not contain secrets.
- **VER-REQ-559:** Tag names shall not contain personal or case identifiers.
- **VER-REQ-560:** Tag version and manifest version shall match.
- **VER-REQ-561:** Tag validation shall verify commit and manifest.
- **VER-REQ-562:** Tag collisions shall be rejected.
- **VER-REQ-563:** Tag history shall remain retrievable.
- **VER-REQ-564:** Repository tag policy shall conform to `OBDIA-GH-001`.
- **VER-REQ-565:** Unknown tag provenance shall block authoritative release use.

## 24. Schema and Interface Versions

- **VER-REQ-566:** Material schemas shall have schema versions.
- **VER-REQ-567:** Material interfaces shall have interface versions.
- **VER-REQ-568:** Schema version shall remain distinct from document version.
- **VER-REQ-569:** Interface version shall remain distinct from implementation version.
- **VER-REQ-570:** Breaking schema changes shall require a major schema increment.
- **VER-REQ-571:** Backward-compatible required-field additions shall be classified carefully and shall not be assumed compatible.
- **VER-REQ-572:** Optional backward-compatible fields may use a minor increment.
- **VER-REQ-573:** Non-semantic schema corrections may use a patch increment.
- **VER-REQ-574:** Schema migration shall preserve identifiers and provenance.
- **VER-REQ-575:** Schema downgrade shall not bypass controls.
- **VER-REQ-576:** Unknown schema versions shall be rejected or quarantined.
- **VER-REQ-577:** Deprecated fields shall identify replacement and removal version.
- **VER-REQ-578:** Removed fields shall require impact analysis.
- **VER-REQ-579:** Interface consumers shall declare supported versions.
- **VER-REQ-580:** Interface providers shall declare supported versions.
- **VER-REQ-581:** Version negotiation shall not select an unapproved insecure version.
- **VER-REQ-582:** Fallback versions shall be no less secure.
- **VER-REQ-583:** Serialization shall preserve schema version.
- **VER-REQ-584:** Evidence schemas shall preserve chain-of-custody semantics.
- **VER-REQ-585:** Authorization schemas shall preserve controlled decision semantics.
- **VER-REQ-586:** Connector schemas shall preserve request and response provenance.
- **VER-REQ-587:** Schema tests shall identify exact versions.
- **VER-REQ-588:** Interface tests shall identify exact versions.
- **VER-REQ-589:** Compatibility evidence shall remain traceable.
- **VER-REQ-590:** Schema and interface history shall remain retrievable.

## 25. Policy and Authorization Versions

- **VER-REQ-591:** Policy bundles shall have immutable version identifiers.
- **VER-REQ-592:** Authorization policies shall identify exact versions.
- **VER-REQ-593:** Policy version shall remain distinct from document version.
- **VER-REQ-594:** Policy decisions shall record the evaluated policy version.
- **VER-REQ-595:** Policy enforcement shall reject unknown mandatory policy versions.
- **VER-REQ-596:** Policy cache keys shall include version where required.
- **VER-REQ-597:** Policy publication shall be attributable.
- **VER-REQ-598:** Policy rollback shall identify restored version.
- **VER-REQ-599:** Policy rollback shall not restore revoked authority.
- **VER-REQ-600:** Policy changes shall receive semantic-impact classification.
- **VER-REQ-601:** Policy changes affecting decision outcomes shall receive security review.
- **VER-REQ-602:** Policy changes affecting obligations shall receive compatibility review.
- **VER-REQ-603:** Policy changes affecting canonical prohibitions shall require superior authority and shall not be permitted through lower-level versioning.
- **VER-REQ-604:** Policy versions shall identify applicable subject, action and resource schemas.
- **VER-REQ-605:** Policy versions shall identify applicable attribute schemas.
- **VER-REQ-606:** Policy versions shall identify dependencies.
- **VER-REQ-607:** Policy versions shall identify effective time.
- **VER-REQ-608:** Policy versions shall identify expiry where applicable.
- **VER-REQ-609:** Stale policy versions shall not authorize privileged actions.
- **VER-REQ-610:** Mixed policy versions shall be prevented or reconciled.
- **VER-REQ-611:** Policy-version propagation shall be tested.
- **VER-REQ-612:** Policy-version drift shall be monitored where implemented.
- **VER-REQ-613:** Policy evidence shall remain immutable.
- **VER-REQ-614:** Authorization tests shall identify policy versions.
- **VER-REQ-615:** Policy history shall remain retrievable.

## 26. Implementation and Artifact Versions

- **VER-REQ-616:** Implemented components shall identify implementation versions.
- **VER-REQ-617:** Implementation versions shall trace to source commits.
- **VER-REQ-618:** Build artifacts shall identify implementation versions.
- **VER-REQ-619:** Implementation versions shall remain distinct from document versions.
- **VER-REQ-620:** Implementation versions shall remain distinct from baseline versions.
- **VER-REQ-621:** Implementation versions shall identify governing requirement versions.
- **VER-REQ-622:** Implementation versions shall identify dependency versions.
- **VER-REQ-623:** Implementation versions shall identify configuration versions.
- **VER-REQ-624:** Implementation versions shall identify model and connector versions where material.
- **VER-REQ-625:** Implementation versions shall identify build provenance.
- **VER-REQ-626:** Implementation versions shall identify test evidence.
- **VER-REQ-627:** Implementation versions shall identify known limitations.
- **VER-REQ-628:** Breaking implementation interfaces shall require major-version assessment.
- **VER-REQ-629:** Backward-compatible features shall require minor-version assessment.
- **VER-REQ-630:** Non-semantic fixes shall require patch-version assessment.
- **VER-REQ-631:** Security fixes may be patch versions while requiring elevated review.
- **VER-REQ-632:** Implementation version shall not imply validation.
- **VER-REQ-633:** Implementation version shall not imply deployment.
- **VER-REQ-634:** Implementation version shall not imply production readiness.
- **VER-REQ-635:** Artifacts with the same version shall not have different bytes without governed build metadata and manifests.
- **VER-REQ-636:** Artifact integrity shall be recorded.
- **VER-REQ-637:** Artifact replacement shall use a new attributable build or version.
- **VER-REQ-638:** Artifact version collisions shall be rejected.
- **VER-REQ-639:** Generated artifacts shall identify generator versions.
- **VER-REQ-640:** Implementation history shall remain retrievable.

## 27. Model, Prompt and AI Configuration Versions

- **VER-REQ-641:** Material model configurations shall have versioned identifiers.
- **VER-REQ-642:** Prompt templates shall have versions where they affect governed behavior.
- **VER-REQ-643:** System instructions shall have versions where they affect governed behavior.
- **VER-REQ-644:** Tool schemas shall have versions.
- **VER-REQ-645:** Retrieval configurations shall have versions where material.
- **VER-REQ-646:** Memory schemas shall have versions.
- **VER-REQ-647:** Model-provider identifiers shall be recorded.
- **VER-REQ-648:** Model-version availability limitations shall be documented.
- **VER-REQ-649:** Model aliases shall not replace exact version evidence where exact versions are available.
- **VER-REQ-650:** Model changes shall trigger risk and validation impact assessment.
- **VER-REQ-651:** Prompt changes shall receive semantic-impact classification.
- **VER-REQ-652:** Tool-schema changes shall receive interface compatibility review.
- **VER-REQ-653:** AI configuration versions shall identify applicable policy versions.
- **VER-REQ-654:** AI configuration versions shall identify applicable connector versions.
- **VER-REQ-655:** AI configuration versions shall identify applicable test evidence.
- **VER-REQ-656:** AI output shall record versions where material to reproducibility.
- **VER-REQ-657:** Unknown model version shall limit validation claims.
- **VER-REQ-658:** Provider-side silent changes shall be treated as external version risk.
- **VER-REQ-659:** Model fallback shall identify fallback version.
- **VER-REQ-660:** Fallback versions shall be pre-approved and no more permissive.
- **VER-REQ-661:** Prompt rollback shall not restore revoked authority.
- **VER-REQ-662:** AI configuration version shall not create legal authority.
- **VER-REQ-663:** AI configuration version shall not imply model certification.
- **VER-REQ-664:** AI-version limitations shall remain visible.
- **VER-REQ-665:** AI configuration history shall remain retrievable.

## 28. Test and Validation Versions

- **VER-REQ-666:** Test plans shall identify versions.
- **VER-REQ-667:** Test cases shall identify versions.
- **VER-REQ-668:** Test suites shall identify versions.
- **VER-REQ-669:** Fixtures shall identify versions or immutable references.
- **VER-REQ-670:** Test environments shall identify versions or immutable definitions.
- **VER-REQ-671:** Test tools shall identify versions.
- **VER-REQ-672:** Test runs shall identify exact tested versions.
- **VER-REQ-673:** Validation decisions shall identify exact evaluated versions.
- **VER-REQ-674:** Test-result status shall remain distinct from version.
- **VER-REQ-675:** Validation outcome shall remain distinct from version.
- **VER-REQ-676:** Passing tests on one version shall not validate another version automatically.
- **VER-REQ-677:** Material implementation changes shall invalidate affected validation evidence.
- **VER-REQ-678:** Material policy changes shall invalidate affected authorization-test evidence.
- **VER-REQ-679:** Material fixture changes shall invalidate affected reproducibility claims.
- **VER-REQ-680:** Test-definition changes after execution shall create new versions.
- **VER-REQ-681:** Expected-result changes shall preserve prior test versions.
- **VER-REQ-682:** Reruns shall not overwrite prior run evidence.
- **VER-REQ-683:** Regression tests shall identify defect and protected versions.
- **VER-REQ-684:** Validation packages shall identify versions and integrity.
- **VER-REQ-685:** Conditional validation shall identify version scope and expiry.
- **VER-REQ-686:** Test versions shall not encode pass or fail as authoritative lifecycle state.
- **VER-REQ-687:** Validation versions shall not create approval.
- **VER-REQ-688:** Archived validation evidence shall preserve tested versions.
- **VER-REQ-689:** Test and validation versioning shall conform to `OBDIA-TEST-001`.
- **VER-REQ-690:** Test-version history shall remain retrievable.

## 29. Evidence and Export Versions

- **VER-REQ-691:** Evidence object identifiers shall remain immutable and shall not be reused as versions.
- **VER-REQ-692:** Evidence metadata schemas shall have versions.
- **VER-REQ-693:** Derived evidence shall identify its own object and transformation versions where applicable.
- **VER-REQ-694:** Evidence bytes shall not be versioned through silent overwrite.
- **VER-REQ-695:** Corrected evidence metadata shall use additive correction records.
- **VER-REQ-696:** Evidence export packages shall have package versions where format or content selection evolves.
- **VER-REQ-697:** Evidence export manifests shall identify object and schema versions.
- **VER-REQ-698:** Evidence integrity values shall identify exact bytes or canonical representation.
- **VER-REQ-699:** Evidence-version metadata shall not imply authenticity.
- **VER-REQ-700:** Evidence-version metadata shall not imply admissibility.
- **VER-REQ-701:** Evidence-version metadata shall not imply lawful acquisition.
- **VER-REQ-702:** Evidence schema migration shall preserve provenance.
- **VER-REQ-703:** Evidence schema migration shall preserve chain of custody.
- **VER-REQ-704:** Evidence schema migration shall preserve prior integrity references.
- **VER-REQ-705:** Unknown evidence schema versions shall be quarantined or qualified.
- **VER-REQ-706:** Export corrections shall preserve prior export manifests.
- **VER-REQ-707:** Redacted derivatives shall receive distinct identifiers and versions where appropriate.
- **VER-REQ-708:** AI-generated analysis shall remain distinct from evidence versions.
- **VER-REQ-709:** Evidence versioning shall not erase original objects.
- **VER-REQ-710:** Evidence versioning shall conform to `OBDIA-EVID-001`.
- **VER-REQ-711:** Evidence retention shall preserve version history.
- **VER-REQ-712:** Evidence disposal shall preserve required version and disposition records.
- **VER-REQ-713:** Evidence-version conflicts shall trigger review.
- **VER-REQ-714:** Export-version history shall remain retrievable.
- **VER-REQ-715:** Evidence-version limitations shall remain visible.

## 30. Deprecation

- **VER-REQ-716:** Deprecated versions shall remain identifiable.
- **VER-REQ-717:** Deprecation shall identify authority.
- **VER-REQ-718:** Deprecation shall identify reason.
- **VER-REQ-719:** Deprecation shall identify effective time.
- **VER-REQ-720:** Deprecation shall identify successor.
- **VER-REQ-721:** Deprecation shall identify migration path.
- **VER-REQ-722:** Deprecation shall identify support period where applicable.
- **VER-REQ-723:** Deprecation shall identify security implications.
- **VER-REQ-724:** Deprecation shall identify privacy implications.
- **VER-REQ-725:** Deprecation shall identify evidence implications.
- **VER-REQ-726:** Deprecation shall identify release implications.
- **VER-REQ-727:** Deprecated versions shall not be removed silently.
- **VER-REQ-728:** Deprecated versions shall not receive new authority.
- **VER-REQ-729:** Deprecated insecure versions shall be disabled or restricted where implemented.
- **VER-REQ-730:** Deprecation warnings shall not expose secrets.
- **VER-REQ-731:** Deprecation shall not imply supersession unless successor authority is explicit.
- **VER-REQ-732:** Deprecation shall remain distinct from archival.
- **VER-REQ-733:** Deprecation shall remain distinct from revocation.
- **VER-REQ-734:** Dependency declarations shall identify deprecated versions.
- **VER-REQ-735:** Tests shall cover deprecated behavior where support remains.
- **VER-REQ-736:** Deprecation exceptions shall have owners and expiry.
- **VER-REQ-737:** Deprecation completion shall preserve history.
- **VER-REQ-738:** Deprecated version use shall be auditable where implemented.
- **VER-REQ-739:** Deprecation records shall remain retrievable.
- **VER-REQ-740:** Unknown deprecation status shall not be treated as supported.

## 31. Supersession

- **VER-REQ-741:** Supersession shall identify the superseded artifact and version.
- **VER-REQ-742:** Supersession shall identify the successor artifact and version.
- **VER-REQ-743:** Supersession shall identify authority.
- **VER-REQ-744:** Supersession shall identify effective time.
- **VER-REQ-745:** Supersession shall identify full or partial scope.
- **VER-REQ-746:** Supersession shall identify dependency impacts.
- **VER-REQ-747:** Supersession shall identify migration.
- **VER-REQ-748:** Supersession shall identify rollback limitations.
- **VER-REQ-749:** Supersession shall identify security and privacy impacts.
- **VER-REQ-750:** Supersession shall identify implementation and validation impacts.
- **VER-REQ-751:** Superseded versions shall remain immutable.
- **VER-REQ-752:** Superseded versions shall remain retrievable.
- **VER-REQ-753:** Superseded versions shall not be edited substantively.
- **VER-REQ-754:** Annotations may point to successors.
- **VER-REQ-755:** Supersession shall not erase historical approval.
- **VER-REQ-756:** Supersession shall not transfer approval automatically to the successor.
- **VER-REQ-757:** Successor approval shall require independent applicable evidence.
- **VER-REQ-758:** Partial supersession shall identify retained authoritative sections.
- **VER-REQ-759:** Ambiguous supersession shall block authoritative use.
- **VER-REQ-760:** Supersession records shall update indexes.
- **VER-REQ-761:** Supersession records shall update dependency declarations.
- **VER-REQ-762:** Supersession shall preserve integrity references.
- **VER-REQ-763:** Supersession shall preserve release history.
- **VER-REQ-764:** Supersession shall conform to `OBDIA-STATE-001`.
- **VER-REQ-765:** Supersession history shall remain retrievable.

## 32. Rollback

- **VER-REQ-766:** Rollback shall be a new controlled action.
- **VER-REQ-767:** Rollback shall not erase the failed or withdrawn version.
- **VER-REQ-768:** Rollback shall identify cause.
- **VER-REQ-769:** Rollback shall identify affected artifacts.
- **VER-REQ-770:** Rollback shall identify restored version.
- **VER-REQ-771:** Rollback shall identify current version after the action.
- **VER-REQ-772:** Rollback shall identify security impact.
- **VER-REQ-773:** Rollback shall identify privacy impact.
- **VER-REQ-774:** Rollback shall identify evidence impact.
- **VER-REQ-775:** Rollback shall identify implementation impact.
- **VER-REQ-776:** Rollback shall identify dependency compatibility.
- **VER-REQ-777:** Rollback shall identify migration or data conversion.
- **VER-REQ-778:** Rollback shall identify credential and policy implications.
- **VER-REQ-779:** Rollback shall identify validation evidence.
- **VER-REQ-780:** Rollback shall identify approval authority.
- **VER-REQ-781:** Rollback shall preserve audit history.
- **VER-REQ-782:** Rollback shall not restore revoked credentials.
- **VER-REQ-783:** Rollback shall not restore revoked agent or connector identities.
- **VER-REQ-784:** Rollback shall not restore prohibited behavior.
- **VER-REQ-785:** Rollback to a known vulnerable version shall require explicit containment and authority.
- **VER-REQ-786:** Rollback shall not create silent version reuse.
- **VER-REQ-787:** Rollback shall update manifests and indexes.
- **VER-REQ-788:** Rollback shall update release and publication records where applicable.
- **VER-REQ-789:** Rollback failure shall remain visible.
- **VER-REQ-790:** Rollback tests shall identify exact versions.

## 33. Migration

- **VER-REQ-791:** Material version changes shall identify migration requirements.
- **VER-REQ-792:** Migration shall identify source version.
- **VER-REQ-793:** Migration shall identify target version.
- **VER-REQ-794:** Migration shall identify prerequisites.
- **VER-REQ-795:** Migration shall identify affected artifacts.
- **VER-REQ-796:** Migration shall identify transformation steps.
- **VER-REQ-797:** Migration shall identify authorization requirements.
- **VER-REQ-798:** Migration shall identify security and privacy controls.
- **VER-REQ-799:** Migration shall identify evidence and provenance preservation.
- **VER-REQ-800:** Migration shall identify downtime or transition conditions where applicable.
- **VER-REQ-801:** Migration shall identify compatibility windows.
- **VER-REQ-802:** Migration shall identify fallback and rollback.
- **VER-REQ-803:** Migration shall identify validation criteria.
- **VER-REQ-804:** Migration shall identify owner.
- **VER-REQ-805:** Migration shall identify completion evidence.
- **VER-REQ-806:** Migration shall not fabricate historical metadata.
- **VER-REQ-807:** Migration shall not reuse immutable identifiers.
- **VER-REQ-808:** Migration shall not silently drop mandatory fields.
- **VER-REQ-809:** Migration shall not broaden authority.
- **VER-REQ-810:** Migration shall not weaken canonical prohibitions.
- **VER-REQ-811:** Partial migration shall remain visible.
- **VER-REQ-812:** Mixed-version operation shall be bounded and reviewed.
- **VER-REQ-813:** Migration failures shall trigger containment.
- **VER-REQ-814:** Migration history shall remain retrievable.
- **VER-REQ-815:** Migration tests shall use controlled data and exact versions.

## 34. Integrity and Version Binding

- **VER-REQ-816:** Approved versions shall have integrity references.
- **VER-REQ-817:** Published versions shall have integrity references.
- **VER-REQ-818:** Baseline members shall have integrity references.
- **VER-REQ-819:** Release artifacts shall have integrity references.
- **VER-REQ-820:** Integrity references shall identify algorithms.
- **VER-REQ-821:** Integrity references shall bind exact bytes or documented canonical representation.
- **VER-REQ-822:** Integrity values shall not replace version identifiers.
- **VER-REQ-823:** Version identifiers shall not replace integrity values.
- **VER-REQ-824:** Identical versions with differing authoritative content shall be rejected.
- **VER-REQ-825:** Identical content under different versions shall require rationale.
- **VER-REQ-826:** Integrity mismatch shall block approval or release.
- **VER-REQ-827:** Integrity mismatch shall trigger incident or correction review where material.
- **VER-REQ-828:** Integrity mismatch shall not be corrected silently.
- **VER-REQ-829:** Integrity calculation shall identify tool or method where material.
- **VER-REQ-830:** Integrity verification shall be attributable.
- **VER-REQ-831:** Integrity verification shall identify time.
- **VER-REQ-832:** Integrity verification shall identify expected and observed values.
- **VER-REQ-833:** Integrity records shall not contain secrets.
- **VER-REQ-834:** Integrity references shall remain available after archival.
- **VER-REQ-835:** Algorithm migration shall preserve old and new values.
- **VER-REQ-836:** Signatures may supplement but shall not replace version and integrity records.
- **VER-REQ-837:** Signature verification failures shall remain visible.
- **VER-REQ-838:** Commit identifiers shall supplement but shall not replace artifact hashes where separate artifacts exist.
- **VER-REQ-839:** Integrity-version bindings shall be included in manifests.
- **VER-REQ-840:** Integrity tests shall use exact versions.

## 35. Machine-Readable Version Records

- **VER-REQ-841:** Version metadata shall be machine-searchable.
- **VER-REQ-842:** Machine-readable version records may use approved sidecars or manifests.
- **VER-REQ-843:** Machine-readable records shall identify schema version.
- **VER-REQ-844:** Machine-readable records shall identify artifact identifier.
- **VER-REQ-845:** Machine-readable records shall identify artifact version.
- **VER-REQ-846:** Machine-readable records shall identify lifecycle status.
- **VER-REQ-847:** Machine-readable records shall identify integrity reference.
- **VER-REQ-848:** Machine-readable records shall identify authoritative source.
- **VER-REQ-849:** Machine-readable records shall identify dependencies.
- **VER-REQ-850:** Machine-readable records shall identify supersession.
- **VER-REQ-851:** Machine-readable records shall identify approval evidence where applicable.
- **VER-REQ-852:** Machine-readable records shall not contain secrets.
- **VER-REQ-853:** Machine-readable records shall minimize personal and case data.
- **VER-REQ-854:** Machine-readable records shall use deterministic field names.
- **VER-REQ-855:** Unknown fields shall be handled safely.
- **VER-REQ-856:** Missing mandatory fields shall fail validation.
- **VER-REQ-857:** Duplicate identifiers shall fail validation.
- **VER-REQ-858:** Duplicate complete version strings within one namespace shall fail validation.
- **VER-REQ-859:** Machine-readable records shall remain synchronized with Markdown metadata.
- **VER-REQ-860:** Synchronization conflicts shall block approval.
- **VER-REQ-861:** Generated version indexes shall identify their source.
- **VER-REQ-862:** Generated version indexes shall receive human review.
- **VER-REQ-863:** Schema migration shall preserve history.
- **VER-REQ-864:** Machine-readable records shall be version-controlled.
- **VER-REQ-865:** Machine-readable validation evidence shall remain retrievable.

## 36. Version Validation

- **VER-REQ-866:** Version validation shall verify core syntax.
- **VER-REQ-867:** Version validation shall verify pre-release syntax where used.
- **VER-REQ-868:** Version validation shall verify build metadata syntax where used.
- **VER-REQ-869:** Version validation shall verify metadata and revision-history consistency.
- **VER-REQ-870:** Version validation shall verify immutable identifiers.
- **VER-REQ-871:** Version validation shall verify no version reuse.
- **VER-REQ-872:** Version validation shall verify semantic-impact classification.
- **VER-REQ-873:** Version validation shall verify dependency compatibility.
- **VER-REQ-874:** Version validation shall verify cross-references.
- **VER-REQ-875:** Version validation shall verify integrity references.
- **VER-REQ-876:** Version validation shall verify status independence.
- **VER-REQ-877:** Version validation shall verify authoritative source.
- **VER-REQ-878:** Version validation shall verify supersession and rollback records.
- **VER-REQ-879:** Version validation shall verify baseline manifests.
- **VER-REQ-880:** Version validation shall verify release manifests.
- **VER-REQ-881:** Version validation shall verify repository tags where applicable.
- **VER-REQ-882:** Version validation shall verify machine-readable records.
- **VER-REQ-883:** Version validation shall identify exact repository commit.
- **VER-REQ-884:** Version validation shall identify tools and versions.
- **VER-REQ-885:** Version-validation failure shall remain visible.
- **VER-REQ-886:** Tool failure shall not count as validation success.
- **VER-REQ-887:** Passing syntax validation shall not establish semantic correctness.
- **VER-REQ-888:** Automated validation shall not approve versions.
- **VER-REQ-889:** Human review shall assess semantic impact.
- **VER-REQ-890:** Version-validation evidence shall conform to `OBDIA-TEST-001`.

## 37. Version Review Gates

- **VER-REQ-891:** Every major change shall receive full scope and dependency review.
- **VER-REQ-892:** Every major change shall receive cross-reference review.
- **VER-REQ-893:** Every major change shall receive applicable security review.
- **VER-REQ-894:** Every major change shall receive applicable privacy review.
- **VER-REQ-895:** Every major change shall receive implementation and validation impact review.
- **VER-REQ-896:** Every minor change shall receive scope and compatibility review.
- **VER-REQ-897:** Every minor change shall receive affected specialist review.
- **VER-REQ-898:** Every patch change shall receive confirmation of non-material impact.
- **VER-REQ-899:** Security patches shall receive security review.
- **VER-REQ-900:** Published corrections shall receive release review.
- **VER-REQ-901:** Baseline version changes shall receive manifest review.
- **VER-REQ-902:** Release version changes shall receive release-gate review.
- **VER-REQ-903:** Schema version changes shall receive compatibility review.
- **VER-REQ-904:** Policy version changes shall receive authorization review.
- **VER-REQ-905:** Evidence version changes shall receive evidence-integrity review.
- **VER-REQ-906:** AI configuration changes shall receive risk and validation review.
- **VER-REQ-907:** Version review shall identify exact commit.
- **VER-REQ-908:** Version review shall identify exact artifact hash.
- **VER-REQ-909:** Version review shall identify affected requirements.
- **VER-REQ-910:** Version review shall identify residual risk.
- **VER-REQ-911:** Version review shall identify unresolved dependencies.
- **VER-REQ-912:** Role concentration shall be disclosed.
- **VER-REQ-913:** Internal review shall not be represented as independent external assurance.
- **VER-REQ-914:** Material changes after review shall invalidate affected evidence.
- **VER-REQ-915:** Project Founder approval shall remain mandatory where required by authority.

## 38. Governance Roles

- **VER-REQ-916:** The Project Founder shall approve constitutional and v1.0 normative baselines.
- **VER-REQ-917:** The Project Founder shall approve major changes requiring Founder authority.
- **VER-REQ-918:** The Documentation Authority shall administer versioning records.
- **VER-REQ-919:** The Documentation Authority shall validate metadata and revision history.
- **VER-REQ-920:** The Documentation Authority shall administer baseline manifests.
- **VER-REQ-921:** The Documentation Authority shall administer supersession and archive records.
- **VER-REQ-922:** Domain owners shall classify semantic impact within their competence.
- **VER-REQ-923:** The Security Reviewer shall assess security-impacting versions.
- **VER-REQ-924:** The Privacy and Governance Reviewer shall assess privacy-impacting versions.
- **VER-REQ-925:** The Implementation Reviewer shall assess implementation and schema compatibility.
- **VER-REQ-926:** The Release Reviewer shall assess release and publication versions.
- **VER-REQ-927:** Evidence custodians shall assess evidence-version integrity where applicable.
- **VER-REQ-928:** Authors and contributors shall propose but not self-approve normative versions.
- **VER-REQ-929:** AI systems shall not approve versions.
- **VER-REQ-930:** AI systems shall not accept version-related risk.
- **VER-REQ-931:** Automation shall not assign final semantic impact without human review.
- **VER-REQ-932:** Role concentration shall require disclosure.
- **VER-REQ-933:** Role concentration shall not be represented as independent assurance.
- **VER-REQ-934:** Emergency version actions shall identify authority.
- **VER-REQ-935:** Role changes shall not erase historical accountability.

## 39. Risk, Exceptions and Versioning

- **VER-REQ-936:** Version-related risks shall have immutable identifiers.
- **VER-REQ-937:** Version risks shall identify affected artifacts and versions.
- **VER-REQ-938:** Version risks shall identify compatibility impact.
- **VER-REQ-939:** Version risks shall identify security and privacy impact.
- **VER-REQ-940:** Version risks shall identify owner.
- **VER-REQ-941:** Version risks shall identify treatment.
- **VER-REQ-942:** Version risks shall identify expiry where accepted temporarily.
- **VER-REQ-943:** An AI system shall not accept version risk.
- **VER-REQ-944:** Version exceptions shall be explicit.
- **VER-REQ-945:** Version exceptions shall be narrow.
- **VER-REQ-946:** Version exceptions shall identify affected requirements.
- **VER-REQ-947:** Version exceptions shall identify rationale.
- **VER-REQ-948:** Version exceptions shall identify compensating controls.
- **VER-REQ-949:** Version exceptions shall identify owner.
- **VER-REQ-950:** Version exceptions shall identify expiry.
- **VER-REQ-951:** Version exceptions shall not authorize version reuse.
- **VER-REQ-952:** Version exceptions shall not authorize silent rewriting.
- **VER-REQ-953:** Version exceptions shall not bypass canonical prohibitions.
- **VER-REQ-954:** Version exceptions shall not fabricate approval.
- **VER-REQ-955:** Expired exceptions shall block progression or require renewed review.
- **VER-REQ-956:** Repeated exceptions shall trigger policy review.
- **VER-REQ-957:** Exceptions shall preserve status-version independence.
- **VER-REQ-958:** Exception use shall remain auditable.
- **VER-REQ-959:** Documents 26–28 remain forward dependencies for detailed risk and exception governance.
- **VER-REQ-960:** Version-risk history shall remain retrievable.

## 40. Change-Control Integration

- **VER-REQ-961:** Every material version increment shall trace to a change record.
- **VER-REQ-962:** Change control shall classify semantic impact.
- **VER-REQ-963:** Change control shall identify affected versions.
- **VER-REQ-964:** Change control shall identify affected dependencies.
- **VER-REQ-965:** Change control shall identify coordinated-change needs.
- **VER-REQ-966:** Change control shall identify migration.
- **VER-REQ-967:** Change control shall identify rollback.
- **VER-REQ-968:** Change control shall identify testing and validation.
- **VER-REQ-969:** Change control shall identify security and privacy impact.
- **VER-REQ-970:** Change control shall identify release impact.
- **VER-REQ-971:** Change control shall identify approval authority.
- **VER-REQ-972:** Emergency changes shall receive retrospective formalization.
- **VER-REQ-973:** Emergency changes shall not erase prior versions.
- **VER-REQ-974:** Unauthorized changes shall not receive authoritative versions.
- **VER-REQ-975:** Abandoned version proposals shall preserve disposition when materially reviewed.
- **VER-REQ-976:** Rejected version proposals shall not become current versions.
- **VER-REQ-977:** Change closure shall verify metadata and manifests.
- **VER-REQ-978:** Change closure shall verify repository tags where applicable.
- **VER-REQ-979:** Change closure shall verify integrity.
- **VER-REQ-980:** Change closure shall verify cross-references.
- **VER-REQ-981:** Document 29 remains the forward authority for complete change-control mechanics.
- **VER-REQ-982:** Versioning shall not substitute for change control.
- **VER-REQ-983:** Change control shall not substitute for versioning.
- **VER-REQ-984:** Version-change evidence shall remain retrievable.
- **VER-REQ-985:** Unresolved change-control conflict shall block approval.

## 41. Publication and Public Claims

- **VER-REQ-986:** Public artifacts shall display accurate versions.
- **VER-REQ-987:** Public artifacts shall display accurate status.
- **VER-REQ-988:** Public artifacts shall identify pre-release status.
- **VER-REQ-989:** Public artifacts shall not represent Draft as Approved.
- **VER-REQ-990:** Public artifacts shall not represent Approved as Implemented.
- **VER-REQ-991:** Public artifacts shall not represent Implemented as Validated without evidence.
- **VER-REQ-992:** Public artifacts shall not represent Validated as Published without release evidence.
- **VER-REQ-993:** Public artifacts shall not imply production readiness from version numbers.
- **VER-REQ-994:** Public artifacts shall not imply certification from version numbers.
- **VER-REQ-995:** Public artifacts shall not imply compliance from version numbers.
- **VER-REQ-996:** Public artifacts shall not imply institutional adoption from version numbers.
- **VER-REQ-997:** Public download names may include release versions when governed.
- **VER-REQ-998:** Canonical repository document filenames shall remain version-free.
- **VER-REQ-999:** Published correction shall use a new version.
- **VER-REQ-1000:** Withdrawn public versions shall retain withdrawal records.
- **VER-REQ-1001:** Public indexes shall identify current and superseded versions.
- **VER-REQ-1002:** Public changelogs shall trace to change records.
- **VER-REQ-1003:** Public manifests shall identify integrity.
- **VER-REQ-1004:** Public release notes shall identify compatibility and limitations.
- **VER-REQ-1005:** Public version claims shall receive Release Reviewer assessment.
- **VER-REQ-1006:** Public version history shall remain retrievable.
- **VER-REQ-1007:** Private incorporation shall remain distinct from public publication.
- **VER-REQ-1008:** Portfolio use shall not exceed versioned evidence.
- **VER-REQ-1009:** Publication versioning shall conform to documents 04 and 15.
- **VER-REQ-1010:** Version disclosure shall not expose restricted information.

## 42. Archival and Retention

- **VER-REQ-1011:** Archived artifacts shall retain their original versions.
- **VER-REQ-1012:** Archived artifacts shall retain status.
- **VER-REQ-1013:** Archived artifacts shall retain integrity references.
- **VER-REQ-1014:** Archived artifacts shall retain authoritative-source references.
- **VER-REQ-1015:** Archived artifacts shall retain revision history.
- **VER-REQ-1016:** Archived artifacts shall retain supersession records.
- **VER-REQ-1017:** Archived artifacts shall retain approval and release evidence.
- **VER-REQ-1018:** Archived artifacts shall remain read-only except for additive archival metadata.
- **VER-REQ-1019:** Archive migration shall preserve version identity.
- **VER-REQ-1020:** Archive restoration shall not create a new approval automatically.
- **VER-REQ-1021:** Archive restoration shall create an attributable event.
- **VER-REQ-1022:** Archived version identifiers shall not be reused.
- **VER-REQ-1023:** Archived tags shall not be moved silently.
- **VER-REQ-1024:** Archived manifests shall remain verifiable.
- **VER-REQ-1025:** Retention rules shall identify versioned artifacts.
- **VER-REQ-1026:** Disposition shall preserve required version and audit records.
- **VER-REQ-1027:** Repository archival shall not publish private versions automatically.
- **VER-REQ-1028:** Archive corruption shall trigger incident handling.
- **VER-REQ-1029:** Archive corrections shall be additive.
- **VER-REQ-1030:** Archive-version history shall remain retrievable.

## 43. Traceability

- **VER-REQ-1031:** Every document version shall trace to an immutable Document ID.
- **VER-REQ-1032:** Every version increment shall trace to a change record.
- **VER-REQ-1033:** Every change record shall trace to exact commits.
- **VER-REQ-1034:** Every approved version shall trace to approval evidence.
- **VER-REQ-1035:** Every implemented version shall trace to governing approved versions.
- **VER-REQ-1036:** Every test run shall trace to exact tested versions.
- **VER-REQ-1037:** Every validation decision shall trace to exact evidence versions.
- **VER-REQ-1038:** Every published version shall trace to a release record.
- **VER-REQ-1039:** Every baseline version shall trace to a baseline manifest.
- **VER-REQ-1040:** Every baseline member shall trace to version and integrity.
- **VER-REQ-1041:** Every release version shall trace to source and manifest.
- **VER-REQ-1042:** Every schema version shall trace to compatibility evidence.
- **VER-REQ-1043:** Every policy version shall trace to authorization tests.
- **VER-REQ-1044:** Every evidence version or schema shall trace to provenance.
- **VER-REQ-1045:** Every rollback shall trace to restored and failed versions.
- **VER-REQ-1046:** Every supersession shall trace to predecessor and successor.
- **VER-REQ-1047:** Every deprecation shall trace to migration.
- **VER-REQ-1048:** Every exception shall trace to affected version requirements.
- **VER-REQ-1049:** Every risk shall trace to affected versions.
- **VER-REQ-1050:** Every archive record shall trace to retained versions.
- **VER-REQ-1051:** Traceability shall be bidirectional.
- **VER-REQ-1052:** Broken version traceability shall block approval or release where material.
- **VER-REQ-1053:** Traceability records shall not contain secrets.
- **VER-REQ-1054:** Traceability records shall minimize personal and case data.
- **VER-REQ-1055:** Baseline freeze shall validate version traceability across documents 01–30.

## 44. Document Versioning Review Gate

- **VER-REQ-1056:** Versioning-policy review shall identify the exact document and commit.
- **VER-REQ-1057:** Review shall verify preservation of all legacy metadata requirements.
- **VER-REQ-1058:** Review shall verify mandatory semantic versioning.
- **VER-REQ-1059:** Review shall verify major, minor and patch semantics.
- **VER-REQ-1060:** Review shall verify pre-release and build metadata rules.
- **VER-REQ-1061:** Review shall verify status-version independence.
- **VER-REQ-1062:** Review shall verify revision and change history.
- **VER-REQ-1063:** Review shall verify author and authoritative-source attribution.
- **VER-REQ-1064:** Review shall verify dependency compatibility.
- **VER-REQ-1065:** Review shall verify coordinated-change requirements.
- **VER-REQ-1066:** Review shall verify baseline and manifest requirements.
- **VER-REQ-1067:** Review shall verify release and tag requirements.
- **VER-REQ-1068:** Review shall verify schema, policy, implementation, AI, test and evidence versions.
- **VER-REQ-1069:** Review shall verify deprecation, supersession, rollback and migration.
- **VER-REQ-1070:** Review shall verify integrity binding.
- **VER-REQ-1071:** Review shall verify machine-readable metadata.
- **VER-REQ-1072:** Review shall verify validation and traceability.
- **VER-REQ-1073:** Review shall identify residual risks and forward dependencies.
- **VER-REQ-1074:** Critical findings shall block progression.
- **VER-REQ-1075:** Material post-review changes shall invalidate affected evidence.
- **VER-REQ-1076:** Role concentration shall be disclosed.
- **VER-REQ-1077:** Internal review shall not be represented as independent external assurance.
- **VER-REQ-1078:** Documentation Authority shall assess metadata and baseline mechanics.
- **VER-REQ-1079:** Security Reviewer shall assess stale, rollback and substitution risk.
- **VER-REQ-1080:** Implementation Reviewer shall assess compatibility and migration.
- **VER-REQ-1081:** Release Reviewer shall assess tags and public versions.
- **VER-REQ-1082:** Project Founder approval shall not substitute for required specialist review.
- **VER-REQ-1083:** Automated findings shall remain advisory for semantic classification.
- **VER-REQ-1084:** Review non-applicability shall have rationale.
- **VER-REQ-1085:** Approval for Draft incorporation shall not approve any baseline or release.

## 45. Minimum Validation Checklist

Before approval of this policy or a governed version transition, confirm:

- [ ] immutable artifact identifier exists;
- [ ] `MAJOR.MINOR.PATCH` syntax is valid;
- [ ] status is declared separately;
- [ ] classification is declared separately;
- [ ] author or contributor origin is attributable;
- [ ] authoritative source is identified;
- [ ] dependencies identify versions or compatibility constraints;
- [ ] revision history is current;
- [ ] change history and change record exist;
- [ ] semantic impact is correctly classified;
- [ ] major changes have full impact and cross-reference review;
- [ ] minor changes remain backward-compatible;
- [ ] patch changes do not alter normative meaning;
- [ ] pre-release labels do not imply approval;
- [ ] build metadata does not alter precedence;
- [ ] coordinated revisions include all affected artifacts;
- [ ] compatibility is demonstrated;
- [ ] baseline manifest identifies every member version and integrity reference;
- [ ] baseline and document versions remain distinct;
- [ ] release and document versions remain distinct;
- [ ] tags identify exact commits and are not reused;
- [ ] schemas, policies, implementations, tests and evidence identify their own versions;
- [ ] deprecation, supersession, migration and rollback preserve history;
- [ ] approved or published artifacts have integrity references;
- [ ] machine-readable metadata matches Markdown metadata;
- [ ] validation evidence identifies exact versions;
- [ ] traceability is bidirectional;
- [ ] forward dependencies 26–30 are recorded;
- [ ] Project Founder approval exists before status becomes Approved.


## 46. Limitations

- This policy does not select a versioning tool, changelog generator, package registry, release platform, signing system or manifest format.
- It does not approve any existing version, baseline or release.
- Semantic versioning communicates intended compatibility but cannot prove compatibility.
- A version number cannot establish authenticity, correctness, security, legal compliance or operational fitness.
- Integrity hashes cannot establish semantic correctness.
- External model and service providers may not expose stable version identifiers.
- Distributed systems may temporarily observe mixed versions and require explicit containment.
- Historical Git data can be replicated outside project control.
- Internal version review is not independent certification.
- Documents 26–30 remain forward dependencies for detailed risk, exception, change and compliance governance.


## 47. Change Control

Every material change to this policy shall identify rationale, affected version domains, compatibility, security and privacy impact, baseline and release impact, migration, validation, rollback, authority and semantic-version effect.

- **VER-REQ-1086:** Editorial corrections shall use a patch version when meaning is unchanged.
- **VER-REQ-1087:** Backward-compatible substantive additions shall use a minor version.
- **VER-REQ-1088:** Incompatible versioning-policy changes shall use a major version.
- **VER-REQ-1089:** Material versioning-architecture decisions shall require an ADR where applicable.
- **VER-REQ-1090:** Changes shall require Project Founder approval.
- **VER-REQ-1091:** Changes shall receive Documentation Authority assessment.
- **VER-REQ-1092:** Security, implementation, evidence and release impacts shall receive specialist review where applicable.
- **VER-REQ-1093:** Changes shall identify affected metadata schemas, manifests, tags and automation.
- **VER-REQ-1094:** Changes shall include migration and compatibility analysis.
- **VER-REQ-1095:** Changes shall include validation and rollback analysis.
- **VER-REQ-1096:** Changes shall not retroactively fabricate revision or approval evidence.
- **VER-REQ-1097:** Historical versions and manifests shall not be silently rewritten.
- **VER-REQ-1098:** Version and requirement identifiers shall not be reused.
- **VER-REQ-1099:** Policy migration shall preserve cross-document traceability.
- **VER-REQ-1100:** Forward-dependency reconciliation shall occur before this policy becomes Approved.

## 48. Consolidation Record

Version 1.0.0 consolidates the two existing Document Versioning Policy variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-VER-001`;
- normalizes the authoritative filename to `25_DOCUMENT_VERSIONING_POLICY.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves Document ID;
- preserves Version;
- preserves Status;
- preserves Author;
- preserves Dependencies;
- preserves Revision History;
- preserves Authoritative source;
- preserves Change history;
- converts semantic versioning from a legacy recommendation into a constitutional requirement;
- preserves major-revision impact assessment and cross-reference review;
- implements mandatory `MAJOR.MINOR.PATCH`;
- defines major, minor, patch, pre-release and build metadata semantics;
- separates version from lifecycle status, classification, implementation, validation and publication;
- adds dependency compatibility, coordinated changes, baseline versions, baseline manifests, freeze, releases, tags, schemas, policies, implementations, AI configurations, tests, evidence, deprecation, supersession, rollback, migration, integrity, machine-readable metadata, validation and traceability;
- identifies documents 26–30 as forward dependencies;
- treats Enterprise and non-Enterprise files as legacy source variants of the same immutable document;
- creates no approved baseline, release, implementation, validation or operational authority.

## 49. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-06 | Draft | Documentation Authority; approval reserved to Project Founder | Constitutional consolidation of both legacy versioning policies: preserved all original fields and major-change review; made semantic versioning mandatory; added complete compatibility, baseline, manifest, release, rollback, validation and traceability governance. |
| 1.0.1 | 2026-08-06 | Draft | Documentation Authority; approval reserved to Project Founder | Reconciled obsolete forward-dependency filenames with the final canonical Knowledge Pack 01–30 filenames; no requirement text, authority, lifecycle status, implementation state or validation claim changed. |
