# EVIDENCE MODEL

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-EVID-001 |
| **Title** | Evidence Model |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Security Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define the authoritative OBDIA evidence object, event, provenance, integrity, custody, derivation, access, export, retention, correction, quarantine, validation and audit requirements. |
| **Scope** | Digital evidence, source records, acquired artifacts, derived artifacts, analytical outputs, collection events, custody events, verification events, transformations, exports, evidence packages, storage, access, retention and controlled demonstrations within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `02_PROJECT_GLOSSARY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `06_IDENTITY_AND_DELEGATION_MODEL.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; forward dependencies `19_AGENT_LIFECYCLE_MODEL.md`, `20_CONNECTOR_SECURITY_POLICY.md`, `22_TESTING_STANDARD.md`, `23_DIAGRAM_STANDARD.md`, `24_REPOSITORY_STRUCTURE_STANDARD.md`, `25_DOCUMENT_VERSIONING_POLICY.md`, `26_AI_RISK_REGISTER.md`, `27_RISK_ASSESSMENT_METHOD.md`, `28_EXCEPTION_MANAGEMENT_POLICY.md`, `29_CHANGE_CONTROL_POLICY.md` and `30_COMPLIANCE_MAPPING_POLICY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GLOSS-001`; `OBDIA-ARCH-001`; `OBDIA-ID-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001` |
| **Cross-References** | Documents 19–30; accepted ADRs; authorization decisions; connector records; acquisition records; custody records; storage records; validation packages; incident, risk, exception, change and release records |
| **Assumptions** | Evidence handling in this repository is demonstrated through synthetic, simulated, testnet, mock, local-laboratory, lawful public, archived-authorized or otherwise explicitly authorized resources. |
| **Constraints** | Evidence handling shall not authorize unlawful collection, real-person public investigations, uncontrolled criminal-infrastructure interaction, silent source modification, fabricated provenance, hidden transformation or autonomous legal conclusions. |
| **Security Considerations** | Evidence can be corrupted, substituted, misattributed, cross-contaminated, over-collected, unlawfully accessed, misleadingly transformed, detached from provenance, exposed through export, or confused with AI-generated analysis. |
| **Validation Criteria** | Every evidence object has a unique immutable identifier, source, collection method, timestamp, integrity reference, collector, provenance, custody history, classification, confidence metadata where used, reviewer information, authorization context, verification events, access controls, retention rules, export metadata and failure handling appropriate to its type. |
| **Implementation Relationship** | This model defines evidence semantics and controls but does not establish legal admissibility, forensic-tool certification, institutional evidence-system accreditation, production deployment or investigative authority. |

---

## 1. Purpose

This document defines the evidence model for the Officer-Bound Digital Investigation Agent project.

The two legacy sources required the following fields:

- unique identifier or identifier;
- source;
- collection method;
- timestamp;
- integrity hash or hash;
- collector;
- provenance;
- chain of custody;
- integrity verification;
- classification;
- confidence;
- reviewer;
- export metadata.

This consolidation preserves every original field and adds the object and event model required to manage evidence throughout acquisition, verification, storage, access, analysis, derivation, custody transfer, export, retention, quarantine, correction and disposition.

`Confidence` is preserved as optional analytical metadata. It shall never be confused with integrity, authenticity, provenance completeness, lawful acquisition, evidentiary relevance, admissibility or factual truth.


## 2. Fundamental Evidence Rules

- **EVID-REQ-001:** Every evidence object shall have a unique immutable identifier.
- **EVID-REQ-002:** Evidence shall preserve documented source and acquisition context.
- **EVID-REQ-003:** Evidence shall preserve provenance.
- **EVID-REQ-004:** Evidence shall preserve chain of custody.
- **EVID-REQ-005:** Evidence shall preserve integrity-verification history.
- **EVID-REQ-006:** Original evidence and derived artifacts shall remain distinguishable.
- **EVID-REQ-007:** AI-generated content shall not be represented as original evidence.
- **EVID-REQ-008:** Analytical confidence shall not be treated as integrity assurance.
- **EVID-REQ-009:** Integrity verification shall not be treated as proof of authenticity, legality, truth or admissibility.
- **EVID-REQ-010:** Evidence accessibility shall not be treated as authorization to collect, process, export or publish it.
- **EVID-REQ-011:** Evidence handling shall remain attributable to identified human, agent, workload and connector actors.
- **EVID-REQ-012:** An AI agent shall not make final legal or evidentiary-admissibility decisions.
- **EVID-REQ-013:** Evidence handling shall remain within case, purpose, jurisdiction, time, tool and data authorization.
- **EVID-REQ-014:** Missing or invalid authorization context shall result in denial or quarantine.
- **EVID-REQ-015:** Canonical prohibitions shall remain non-authorizable.
- **EVID-REQ-016:** Evidence records shall not be silently overwritten.
- **EVID-REQ-017:** Corrections shall be additive and traceable.
- **EVID-REQ-018:** Evidence-integrity failure shall trigger quarantine or qualified handling.
- **EVID-REQ-019:** Evidence handling shall use least privilege and data minimization.
- **EVID-REQ-020:** Repository merge shall not establish evidence validity.
- **EVID-REQ-021:** Automated verification shall not replace required human review.
- **EVID-REQ-022:** Evidence controls shall be testable and bidirectionally traceable.
- **EVID-REQ-023:** Evidence handling shall preserve human and institutional accountability.
- **EVID-REQ-024:** Unresolved material evidence ambiguity shall remain visible.

## 3. Evidence Object Model

- **EVID-REQ-025:** Every evidence object shall have an immutable evidence identifier.
- **EVID-REQ-026:** Every evidence object shall identify its evidence-object type.
- **EVID-REQ-027:** Every evidence object shall identify its authoritative source.
- **EVID-REQ-028:** Every evidence object shall identify the source-system or source-location class where applicable.
- **EVID-REQ-029:** Every evidence object shall identify the collection method.
- **EVID-REQ-030:** Every evidence object shall identify collection time or time interval.
- **EVID-REQ-031:** Every evidence object shall identify collector identity.
- **EVID-REQ-032:** Every evidence object shall identify accountable institution.
- **EVID-REQ-033:** Every evidence object shall identify case or mandate where applicable.
- **EVID-REQ-034:** Every evidence object shall identify approved purpose.
- **EVID-REQ-035:** Every evidence object shall identify applicable jurisdiction where relevant.
- **EVID-REQ-036:** Every evidence object shall identify authorization-request or decision references.
- **EVID-REQ-037:** Every evidence object shall identify the acquiring tool or connector where applicable.
- **EVID-REQ-038:** Every evidence object shall identify the acquiring workload or agent where applicable.
- **EVID-REQ-039:** Every evidence object shall identify original filename or source label where available.
- **EVID-REQ-040:** Every evidence object shall identify canonical stored-object name.
- **EVID-REQ-041:** Every evidence object shall identify media type or content type where applicable.
- **EVID-REQ-042:** Every evidence object shall identify byte length or equivalent size where applicable.
- **EVID-REQ-043:** Every evidence object shall identify integrity algorithm.
- **EVID-REQ-044:** Every evidence object shall identify integrity value.
- **EVID-REQ-045:** Every evidence object shall identify integrity-verification status.
- **EVID-REQ-046:** Every evidence object shall identify provenance-record references.
- **EVID-REQ-047:** Every evidence object shall identify chain-of-custody record references.
- **EVID-REQ-048:** Every evidence object shall identify classification.
- **EVID-REQ-049:** Every evidence object shall identify access-control references.
- **EVID-REQ-050:** Every evidence object shall identify retention and disposition references.
- **EVID-REQ-051:** Every evidence object shall identify reviewer or reviewer-role references where review occurred.
- **EVID-REQ-052:** Every evidence object shall identify confidence metadata only when a defined method exists.
- **EVID-REQ-053:** Every evidence object shall identify known limitations.
- **EVID-REQ-054:** Every evidence object shall identify current handling state.
- **EVID-REQ-055:** Every evidence object shall identify creation and last-governed-update timestamps.
- **EVID-REQ-056:** Every evidence object shall identify parent and child lineage where applicable.
- **EVID-REQ-057:** Every evidence object shall identify export-package references where exported.
- **EVID-REQ-058:** Every evidence object shall identify quarantine references where quarantined.
- **EVID-REQ-059:** Every evidence object shall identify correction records where corrected.
- **EVID-REQ-060:** Malformed evidence objects shall be rejected or quarantined.
- **EVID-REQ-061:** Duplicate evidence identifiers shall be rejected.
- **EVID-REQ-062:** Evidence identifiers shall conform to `OBDIA-NAME-001`.
- **EVID-REQ-063:** Evidence metadata shall not contain secrets.
- **EVID-REQ-064:** Evidence metadata shall minimize unnecessary personal data.

## 4. Evidence Object Types

- **EVID-REQ-065:** Source evidence shall represent information acquired from an identified source.
- **EVID-REQ-066:** Native evidence shall preserve source-native bytes or representation where available.
- **EVID-REQ-067:** Logical evidence shall represent a governed logical extraction.
- **EVID-REQ-068:** Metadata evidence shall represent metadata collected as evidence.
- **EVID-REQ-069:** Event evidence shall represent an observed or recorded event.
- **EVID-REQ-070:** Network evidence shall identify capture scope and collection mechanism.
- **EVID-REQ-071:** Blockchain evidence shall identify network, chain, block or transaction context as applicable.
- **EVID-REQ-072:** Smart-contract evidence shall identify contract, network and code or state context.
- **EVID-REQ-073:** Cloud evidence shall identify provider, tenant, account, region and service context where authorized.
- **EVID-REQ-074:** Web evidence shall identify URL, retrieval time, response context and capture method.
- **EVID-REQ-075:** Communication evidence shall identify platform, account context and authorized acquisition method.
- **EVID-REQ-076:** Connector-acquired evidence shall identify connector identity and policy version.
- **EVID-REQ-077:** Agent-observed evidence shall identify agent, officer binding and runtime context.
- **EVID-REQ-078:** Derived evidence shall identify all parent evidence and transformations.
- **EVID-REQ-079:** Analytical output shall remain distinct from evidence unless formally incorporated through an authorized evidence process.
- **EVID-REQ-080:** AI-generated summary shall be classified as derived analytical output, not source evidence.
- **EVID-REQ-081:** Screenshot evidence shall identify capture method, target, time and limitations.
- **EVID-REQ-082:** Document export evidence shall identify originating system and export parameters.
- **EVID-REQ-083:** Database export evidence shall identify query or selection scope where permissible.
- **EVID-REQ-084:** File-system evidence shall identify source path or logical location without exposing restricted information unnecessarily.
- **EVID-REQ-085:** Memory evidence shall identify acquisition scope and volatility limitations.
- **EVID-REQ-086:** Log evidence shall identify log source, time basis and retention limitations.
- **EVID-REQ-087:** Test or synthetic evidence shall be unmistakably labeled.
- **EVID-REQ-088:** Demonstration evidence shall not be confused with operational evidence.
- **EVID-REQ-089:** Archived-authorized evidence shall identify archive authority and retrieval reference.
- **EVID-REQ-090:** Unknown evidence types shall be quarantined until classified.
- **EVID-REQ-091:** Type conversion shall be recorded as a transformation.
- **EVID-REQ-092:** Evidence type shall not imply trustworthiness or admissibility.
- **EVID-REQ-093:** Evidence type shall not broaden authorization.
- **EVID-REQ-094:** Evidence-type definitions shall be versioned and reviewable.

## 5. Acquisition Event Model

- **EVID-REQ-095:** Every acquisition shall produce an immutable acquisition-event identifier.
- **EVID-REQ-096:** An acquisition event shall identify evidence identifier.
- **EVID-REQ-097:** An acquisition event shall identify collector identity.
- **EVID-REQ-098:** An acquisition event shall identify officer, agent, workload and connector identities where applicable.
- **EVID-REQ-099:** An acquisition event shall identify institution.
- **EVID-REQ-100:** An acquisition event shall identify case or mandate.
- **EVID-REQ-101:** An acquisition event shall identify purpose.
- **EVID-REQ-102:** An acquisition event shall identify jurisdiction where applicable.
- **EVID-REQ-103:** An acquisition event shall identify authorization-decision reference.
- **EVID-REQ-104:** An acquisition event shall identify collection method.
- **EVID-REQ-105:** An acquisition event shall identify source.
- **EVID-REQ-106:** An acquisition event shall identify source ownership or custody where known.
- **EVID-REQ-107:** An acquisition event shall identify collection start time.
- **EVID-REQ-108:** An acquisition event shall identify collection completion time.
- **EVID-REQ-109:** An acquisition event shall identify time source and known clock limitations.
- **EVID-REQ-110:** An acquisition event shall identify tool, connector and version.
- **EVID-REQ-111:** An acquisition event shall identify configuration and policy version where material.
- **EVID-REQ-112:** An acquisition event shall identify requested and acquired scope.
- **EVID-REQ-113:** An acquisition event shall identify filtering or minimization applied.
- **EVID-REQ-114:** An acquisition event shall identify errors, omissions and partial results.
- **EVID-REQ-115:** An acquisition event shall identify original integrity value when available.
- **EVID-REQ-116:** An acquisition event shall identify calculated acquisition integrity value.
- **EVID-REQ-117:** An acquisition event shall identify storage destination.
- **EVID-REQ-118:** An acquisition event shall identify immediate custody recipient.
- **EVID-REQ-119:** An acquisition event shall identify source-side modification risk.
- **EVID-REQ-120:** An acquisition event shall identify whether collection was read-only.
- **EVID-REQ-121:** An acquisition event shall identify whether source interaction occurred.
- **EVID-REQ-122:** An acquisition event shall identify whether the source was synthetic, mock, testnet, lawful public, archived-authorized or otherwise authorized.
- **EVID-REQ-123:** An acquisition event shall not conceal failed or partial collection.
- **EVID-REQ-124:** Acquisition retries shall create new attributable events.
- **EVID-REQ-125:** Acquisition shall not exceed authorized scope.
- **EVID-REQ-126:** Collection method changes shall trigger authorization and provenance review.
- **EVID-REQ-127:** Uncontrolled criminal-infrastructure interaction shall not be used for acquisition.
- **EVID-REQ-128:** Real-person public demonstrations shall not be used for acquisition.
- **EVID-REQ-129:** Acquisition events shall be integrity-protected.
- **EVID-REQ-130:** Acquisition events shall remain linked to all produced evidence objects.
- **EVID-REQ-131:** Acquisition failure shall not fabricate a complete evidence object.
- **EVID-REQ-132:** Acquisition uncertainty shall be recorded.
- **EVID-REQ-133:** Acquisition events shall not contain reusable credentials.
- **EVID-REQ-134:** Acquisition tooling limitations shall be documented.

## 6. Source and Collection Method

- **EVID-REQ-135:** Source shall identify the origin from which evidence was acquired.
- **EVID-REQ-136:** Source descriptions shall distinguish system, service, person-provided, archive, network, blockchain, file, log and generated-test origins.
- **EVID-REQ-137:** Source identifiers shall use opaque or minimized values where disclosure is restricted.
- **EVID-REQ-138:** Source authority shall be recorded where applicable.
- **EVID-REQ-139:** Source accessibility shall not be treated as lawful authority.
- **EVID-REQ-140:** Source reputation shall not be treated as evidence integrity.
- **EVID-REQ-141:** Source-provided metadata shall remain distinguishable from collector-generated metadata.
- **EVID-REQ-142:** Source-provided timestamps shall identify their origin.
- **EVID-REQ-143:** Source modification after acquisition shall not alter the preserved evidence object.
- **EVID-REQ-144:** Source disappearance shall not invalidate preserved evidence automatically.
- **EVID-REQ-145:** Collection method shall identify the technical and procedural acquisition mechanism.
- **EVID-REQ-146:** Collection method shall identify whether acquisition was physical, logical, exported, captured, queried, observed or generated.
- **EVID-REQ-147:** Collection method shall identify read-only or write-capable behavior.
- **EVID-REQ-148:** Collection method shall identify known side effects.
- **EVID-REQ-149:** Collection method shall identify filtering, normalization and decompression where applied.
- **EVID-REQ-150:** Collection method shall identify tool and version.
- **EVID-REQ-151:** Collection method shall identify configuration where material.
- **EVID-REQ-152:** Collection method shall identify validation status.
- **EVID-REQ-153:** Collection method shall identify whether it is proposed, implemented, tested or validated.
- **EVID-REQ-154:** Collection method shall identify known lossiness.
- **EVID-REQ-155:** Collection method shall identify unsupported data types or fields.
- **EVID-REQ-156:** Collection method shall identify time and ordering limitations.
- **EVID-REQ-157:** Collection method shall identify privilege required.
- **EVID-REQ-158:** Collection method shall identify authorization obligations.
- **EVID-REQ-159:** Collection method shall identify network or connector dependencies.
- **EVID-REQ-160:** Collection method shall not imply forensic soundness without evidence.
- **EVID-REQ-161:** Collection method shall not imply legal admissibility.
- **EVID-REQ-162:** Material collection-method changes shall require review.
- **EVID-REQ-163:** Collection-method implementation shall comply with documents 16, 17 and 20.
- **EVID-REQ-164:** Collection methods used in demonstrations shall remain safe and authorized.

## 7. Provenance Model

- **EVID-REQ-165:** Provenance shall document evidence origin and processing history.
- **EVID-REQ-166:** Provenance shall identify every material acquisition event.
- **EVID-REQ-167:** Provenance shall identify every material transformation.
- **EVID-REQ-168:** Provenance shall identify every derived artifact.
- **EVID-REQ-169:** Provenance shall identify responsible entities.
- **EVID-REQ-170:** Provenance shall identify material tools and versions.
- **EVID-REQ-171:** Provenance shall identify material policy and configuration versions.
- **EVID-REQ-172:** Provenance shall identify storage transitions.
- **EVID-REQ-173:** Provenance shall identify custody transfers.
- **EVID-REQ-174:** Provenance shall identify verification events.
- **EVID-REQ-175:** Provenance shall identify corrections.
- **EVID-REQ-176:** Provenance shall identify exports.
- **EVID-REQ-177:** Provenance shall identify quarantine and release-from-quarantine events.
- **EVID-REQ-178:** Provenance shall identify retention and disposition events.
- **EVID-REQ-179:** Provenance shall preserve parent-child lineage.
- **EVID-REQ-180:** Provenance shall distinguish observed facts from inferred metadata.
- **EVID-REQ-181:** Provenance shall distinguish source-provided metadata from collector-generated metadata.
- **EVID-REQ-182:** Provenance shall distinguish human actions from agent, workload and connector actions.
- **EVID-REQ-183:** Provenance shall distinguish deterministic transformations from model-assisted transformations.
- **EVID-REQ-184:** Provenance shall identify uncertainty and missing links.
- **EVID-REQ-185:** Provenance shall not be fabricated to fill gaps.
- **EVID-REQ-186:** Unknown provenance shall remain explicitly unknown.
- **EVID-REQ-187:** Partial provenance shall not be described as complete.
- **EVID-REQ-188:** Provenance records shall use immutable identifiers.
- **EVID-REQ-189:** Provenance records shall be append-only or tamper-evident where feasible.
- **EVID-REQ-190:** Provenance records shall not be silently rewritten.
- **EVID-REQ-191:** Provenance corrections shall preserve prior assertions.
- **EVID-REQ-192:** Provenance shall support reproducibility where technically feasible.
- **EVID-REQ-193:** Provenance shall remain accessible to authorized reviewers.
- **EVID-REQ-194:** Provenance access shall be logged.
- **EVID-REQ-195:** Provenance shall minimize unnecessary sensitive data.
- **EVID-REQ-196:** Provenance shall remain traceable to authorization evidence.
- **EVID-REQ-197:** Provenance shall remain traceable to chain-of-custody records.
- **EVID-REQ-198:** Provenance shall remain traceable to export manifests.
- **EVID-REQ-199:** Broken provenance affecting integrity or authority shall trigger quarantine or qualified handling.
- **EVID-REQ-200:** Provenance validation shall be included in evidence tests.
- **EVID-REQ-201:** Generated provenance shall receive human review where material.
- **EVID-REQ-202:** An AI-generated provenance narrative shall not replace source event records.
- **EVID-REQ-203:** Provenance interoperability mappings shall not replace the authoritative project record.
- **EVID-REQ-204:** Provenance retention shall be at least as protective as the governed evidence retention requirement.

## 8. Integrity Hash and Verification

- **EVID-REQ-205:** Every byte-addressable evidence object shall have an integrity value where technically feasible.
- **EVID-REQ-206:** Integrity values shall identify the algorithm.
- **EVID-REQ-207:** SHA-256 or a formally approved equivalent shall be used for repository research artifacts unless an approved ADR specifies otherwise.
- **EVID-REQ-208:** Integrity values shall be calculated over exact bytes or a documented canonical representation.
- **EVID-REQ-209:** Canonicalization shall be documented when integrity is not calculated over exact source bytes.
- **EVID-REQ-210:** Integrity calculation shall identify tool and version where material.
- **EVID-REQ-211:** Integrity calculation shall identify time.
- **EVID-REQ-212:** Integrity calculation shall identify responsible actor.
- **EVID-REQ-213:** Integrity verification shall create an immutable verification event.
- **EVID-REQ-214:** Verification events shall identify expected and observed values.
- **EVID-REQ-215:** Verification events shall identify algorithm.
- **EVID-REQ-216:** Verification events shall identify tool and version.
- **EVID-REQ-217:** Verification events shall identify time and verifier.
- **EVID-REQ-218:** Verification events shall identify verification result.
- **EVID-REQ-219:** Verification events shall identify errors and limitations.
- **EVID-REQ-220:** Successful integrity verification shall not prove authenticity.
- **EVID-REQ-221:** Successful integrity verification shall not prove lawful acquisition.
- **EVID-REQ-222:** Successful integrity verification shall not prove factual truth.
- **EVID-REQ-223:** Successful integrity verification shall not prove admissibility.
- **EVID-REQ-224:** Successful integrity verification shall not prove provenance completeness.
- **EVID-REQ-225:** Integrity mismatch shall trigger quarantine.
- **EVID-REQ-226:** Integrity mismatch shall not be silently corrected.
- **EVID-REQ-227:** Re-hashing shall not erase prior mismatch events.
- **EVID-REQ-228:** Algorithm migration shall preserve old and new integrity values.
- **EVID-REQ-229:** Deprecated algorithms shall trigger migration review.
- **EVID-REQ-230:** Integrity metadata shall itself be integrity-protected.
- **EVID-REQ-231:** Integrity values shall be included in export manifests.
- **EVID-REQ-232:** Integrity values shall be verified after transfer or packaging.
- **EVID-REQ-233:** Integrity verification shall occur before and after material transformation where applicable.
- **EVID-REQ-234:** Integrity verification shall occur before release from quarantine.
- **EVID-REQ-235:** Integrity verification shall occur before export where applicable.
- **EVID-REQ-236:** Integrity verification shall occur during periodic preservation checks where required.
- **EVID-REQ-237:** Unknown integrity status shall remain explicit.
- **EVID-REQ-238:** Non-byte evidence shall define an equivalent verifiable representation or record non-applicability.
- **EVID-REQ-239:** Integrity failure shall generate an incident when material.
- **EVID-REQ-240:** Integrity-verification automation shall operate on exact object references.
- **EVID-REQ-241:** Passing automated verification shall not establish complete evidence validity.
- **EVID-REQ-242:** Verification failures shall remain visible.
- **EVID-REQ-243:** Integrity controls shall be tested with altered and substituted objects.
- **EVID-REQ-244:** Integrity records shall remain traceable to evidence and custody events.

## 9. Chain of Custody

- **EVID-REQ-245:** Every custody-relevant action shall produce a custody event.
- **EVID-REQ-246:** Custody events shall have immutable identifiers.
- **EVID-REQ-247:** Custody events shall identify evidence identifier.
- **EVID-REQ-248:** Custody events shall identify event type.
- **EVID-REQ-249:** Custody events shall identify prior custodian where applicable.
- **EVID-REQ-250:** Custody events shall identify receiving custodian where applicable.
- **EVID-REQ-251:** Custody events shall identify human, agent, workload and institutional actors as applicable.
- **EVID-REQ-252:** Custody events shall identify event time.
- **EVID-REQ-253:** Custody events shall identify time source and known limitations.
- **EVID-REQ-254:** Custody events shall identify location or storage domain at an appropriate classification level.
- **EVID-REQ-255:** Custody events shall identify action performed.
- **EVID-REQ-256:** Custody events shall identify authorization reference.
- **EVID-REQ-257:** Custody events shall identify integrity status before transfer.
- **EVID-REQ-258:** Custody events shall identify integrity status after transfer.
- **EVID-REQ-259:** Custody events shall identify packaging or container reference.
- **EVID-REQ-260:** Custody events shall identify seal or equivalent control where used.
- **EVID-REQ-261:** Custody events shall identify reason and purpose.
- **EVID-REQ-262:** Custody events shall identify conditions and limitations.
- **EVID-REQ-263:** Custody events shall identify acknowledgement by receiving custodian where applicable.
- **EVID-REQ-264:** Custody events shall identify failed or refused transfer.
- **EVID-REQ-265:** Custody transfer shall not occur to an unauthorized custodian.
- **EVID-REQ-266:** Custody transfer shall not broaden case or purpose access.
- **EVID-REQ-267:** Custody transfer shall preserve provenance.
- **EVID-REQ-268:** Custody transfer shall preserve integrity references.
- **EVID-REQ-269:** Custody transfer shall preserve classification and handling obligations.
- **EVID-REQ-270:** Custody records shall be chronological and verifiable.
- **EVID-REQ-271:** Custody gaps shall be recorded explicitly.
- **EVID-REQ-272:** Custody gaps shall not be fabricated away.
- **EVID-REQ-273:** Unresolved material custody gaps shall trigger quarantine or qualified handling.
- **EVID-REQ-274:** Automated system custody shall identify the responsible institution and system identity.
- **EVID-REQ-275:** Cloud or external-service custody shall identify service and account context.
- **EVID-REQ-276:** Connector custody shall identify connector identity and destination.
- **EVID-REQ-277:** Temporary custody shall have expiry or return conditions.
- **EVID-REQ-278:** Custody relinquishment shall be recorded.
- **EVID-REQ-279:** Custody records shall be append-only or tamper-evident where feasible.
- **EVID-REQ-280:** Custody corrections shall be additive.
- **EVID-REQ-281:** Custody history shall remain retrievable.
- **EVID-REQ-282:** Custody access shall be authorized and logged.
- **EVID-REQ-283:** Custody records shall minimize unnecessary personal data.
- **EVID-REQ-284:** Chain-of-custody tests shall verify transfer, rejection, mismatch and recovery cases.

## 10. Handling States and Evidence Lifecycle

- **EVID-REQ-285:** Evidence handling state shall remain distinct from repository document lifecycle state.
- **EVID-REQ-286:** Permitted evidence handling states shall be explicitly defined.
- **EVID-REQ-287:** `Pending Verification` shall indicate evidence not yet integrity-verified.
- **EVID-REQ-288:** `Verified` shall indicate successful defined integrity verification only.
- **EVID-REQ-289:** `In Custody` shall indicate controlled custody and storage.
- **EVID-REQ-290:** `Under Analysis` shall indicate authorized analytical use.
- **EVID-REQ-291:** `Quarantined` shall indicate restricted handling due to integrity, provenance, authorization or safety concern.
- **EVID-REQ-292:** `Restricted` shall indicate access limited by classification or legal constraint.
- **EVID-REQ-293:** `Exported` shall indicate inclusion in an authorized export package.
- **EVID-REQ-294:** `Retained` shall indicate governed preservation under a retention rule.
- **EVID-REQ-295:** `Disposed` shall indicate authorized terminal disposition while retaining required disposition evidence.
- **EVID-REQ-296:** Handling state shall not imply authenticity, truth, admissibility or legal sufficiency.
- **EVID-REQ-297:** Handling-state changes shall create events.
- **EVID-REQ-298:** Handling-state events shall identify actor, time, reason and authorization.
- **EVID-REQ-299:** Invalid state transitions shall be rejected.
- **EVID-REQ-300:** Pending verification shall not be represented as verified.
- **EVID-REQ-301:** Quarantined evidence shall not be used operationally without explicit qualified authorization.
- **EVID-REQ-302:** Restricted evidence shall not be exported without explicit authority.
- **EVID-REQ-303:** Disposed evidence shall not be silently restored.
- **EVID-REQ-304:** Restoration from backup shall create a new controlled event.
- **EVID-REQ-305:** Derived artifacts shall have independent handling state.
- **EVID-REQ-306:** Parent evidence state changes shall trigger child-impact review where material.
- **EVID-REQ-307:** Case closure shall trigger retention and disposition review.
- **EVID-REQ-308:** Authorization revocation shall restrict future access without erasing lawful history.
- **EVID-REQ-309:** Handling-state automation shall fail closed.
- **EVID-REQ-310:** State conflicts shall resolve to the more restrictive state.
- **EVID-REQ-311:** State displays shall identify authoritative event records.
- **EVID-REQ-312:** Handling-state history shall be retained.
- **EVID-REQ-313:** Handling state shall be testable.
- **EVID-REQ-314:** Document 19 shall govern agent runtime lifecycle and shall not replace evidence handling state.

## 11. Derived Evidence and Lineage

- **EVID-REQ-315:** Every derived evidence object shall have its own immutable identifier.
- **EVID-REQ-316:** Derived evidence shall identify every parent evidence object.
- **EVID-REQ-317:** Derived evidence shall identify transformation-event references.
- **EVID-REQ-318:** Derived evidence shall identify transformation method.
- **EVID-REQ-319:** Derived evidence shall identify tool, model or script version.
- **EVID-REQ-320:** Derived evidence shall identify parameters and configuration where material.
- **EVID-REQ-321:** Derived evidence shall identify responsible actor.
- **EVID-REQ-322:** Derived evidence shall identify transformation time.
- **EVID-REQ-323:** Derived evidence shall identify authorization reference.
- **EVID-REQ-324:** Derived evidence shall identify information loss or approximation.
- **EVID-REQ-325:** Derived evidence shall identify whether transformation was deterministic.
- **EVID-REQ-326:** Derived evidence shall identify whether a model contributed.
- **EVID-REQ-327:** Derived evidence shall identify human review where material.
- **EVID-REQ-328:** Derived evidence shall identify integrity value.
- **EVID-REQ-329:** Derived evidence shall not replace or overwrite parent evidence.
- **EVID-REQ-330:** Derived evidence shall not be represented as native source evidence.
- **EVID-REQ-331:** Derived evidence shall preserve parent integrity references.
- **EVID-REQ-332:** Derived evidence shall preserve source provenance links.
- **EVID-REQ-333:** Derived evidence shall preserve case and purpose restrictions.
- **EVID-REQ-334:** Derived evidence shall not broaden data scope.
- **EVID-REQ-335:** Derived evidence shall not silently combine unrelated cases.
- **EVID-REQ-336:** Cross-source correlation shall identify every source object.
- **EVID-REQ-337:** Aggregate evidence shall identify inclusion and exclusion criteria.
- **EVID-REQ-338:** Redacted evidence shall identify redaction method and authority.
- **EVID-REQ-339:** Normalized evidence shall identify normalization rules.
- **EVID-REQ-340:** Decoded evidence shall identify decoder and version.
- **EVID-REQ-341:** Decompressed evidence shall identify archive or compression source.
- **EVID-REQ-342:** Translated evidence shall identify language, method and reviewer where material.
- **EVID-REQ-343:** OCR-derived text shall identify OCR method, language and known limitations.
- **EVID-REQ-344:** Summarized evidence shall identify summarization method and reviewer.
- **EVID-REQ-345:** AI-generated summaries shall remain analytical derivatives.
- **EVID-REQ-346:** Confidence in derived evidence shall not replace lineage.
- **EVID-REQ-347:** Derived-object verification shall be independent from parent verification.
- **EVID-REQ-348:** Derived evidence errors shall not alter parent evidence.
- **EVID-REQ-349:** Correction of derived evidence shall create a new version or object with preserved history.
- **EVID-REQ-350:** Derived evidence export shall include lineage metadata.
- **EVID-REQ-351:** Derived evidence deletion shall preserve required lineage and disposition records.
- **EVID-REQ-352:** Lineage gaps shall remain visible.
- **EVID-REQ-353:** Material lineage failure shall trigger quarantine.
- **EVID-REQ-354:** Derived-evidence tests shall verify reproducibility and parent-child traceability.

## 12. Transformation Event Model

- **EVID-REQ-355:** Every material transformation shall create a transformation event.
- **EVID-REQ-356:** Transformation events shall have immutable identifiers.
- **EVID-REQ-357:** Transformation events shall identify input evidence.
- **EVID-REQ-358:** Transformation events shall identify output evidence.
- **EVID-REQ-359:** Transformation events shall identify method.
- **EVID-REQ-360:** Transformation events shall identify tool, script, model and version where applicable.
- **EVID-REQ-361:** Transformation events shall identify configuration and parameters where material.
- **EVID-REQ-362:** Transformation events shall identify responsible actor.
- **EVID-REQ-363:** Transformation events shall identify time.
- **EVID-REQ-364:** Transformation events shall identify authorization reference.
- **EVID-REQ-365:** Transformation events shall identify purpose.
- **EVID-REQ-366:** Transformation events shall identify deterministic or nondeterministic behavior.
- **EVID-REQ-367:** Transformation events shall identify expected loss, approximation or alteration.
- **EVID-REQ-368:** Transformation events shall identify pre-transformation integrity status.
- **EVID-REQ-369:** Transformation events shall identify post-transformation integrity values.
- **EVID-REQ-370:** Transformation events shall identify validation result.
- **EVID-REQ-371:** Transformation events shall identify errors and partial output.
- **EVID-REQ-372:** Transformation events shall identify human review where required.
- **EVID-REQ-373:** Transformation events shall not conceal failed steps.
- **EVID-REQ-374:** Transformation retries shall create new events.
- **EVID-REQ-375:** Transformations shall not modify source evidence in place.
- **EVID-REQ-376:** Transformations shall not remove case or purpose restrictions.
- **EVID-REQ-377:** Transformations shall not invent missing provenance.
- **EVID-REQ-378:** Transformations shall not be represented as collection events.
- **EVID-REQ-379:** Model-assisted transformation shall preserve prompt and model context where material and lawful.
- **EVID-REQ-380:** Model-assisted transformation shall identify uncertainty and verification requirements.
- **EVID-REQ-381:** Transformation tooling shall be included in supply-chain review.
- **EVID-REQ-382:** Transformation events shall be append-only or tamper-evident where feasible.
- **EVID-REQ-383:** Transformation-event corrections shall preserve prior records.
- **EVID-REQ-384:** Transformation events shall be retained with derived evidence.
- **EVID-REQ-385:** Transformation tests shall include corrupted, unsupported and partial inputs.
- **EVID-REQ-386:** Transformation reproducibility limitations shall remain visible.

## 13. Confidence and Analytical Assessment

- **EVID-REQ-387:** Confidence metadata shall be optional unless an approved method requires it.
- **EVID-REQ-388:** Confidence shall identify the subject of the assessment.
- **EVID-REQ-389:** Confidence shall identify the assessment method.
- **EVID-REQ-390:** Confidence shall identify the scale or controlled vocabulary.
- **EVID-REQ-391:** Confidence shall identify the assessor.
- **EVID-REQ-392:** Confidence shall identify assessment time.
- **EVID-REQ-393:** Confidence shall identify evidence and assumptions used.
- **EVID-REQ-394:** Confidence shall identify uncertainty and limitations.
- **EVID-REQ-395:** Confidence shall distinguish source reliability from content credibility where applicable.
- **EVID-REQ-396:** Confidence shall distinguish human assessment from model-generated scoring.
- **EVID-REQ-397:** Confidence shall not be inferred from integrity verification.
- **EVID-REQ-398:** Confidence shall not be inferred from a valid hash.
- **EVID-REQ-399:** Confidence shall not be treated as authenticity.
- **EVID-REQ-400:** Confidence shall not be treated as lawful acquisition.
- **EVID-REQ-401:** Confidence shall not be treated as admissibility.
- **EVID-REQ-402:** Confidence shall not be treated as factual truth.
- **EVID-REQ-403:** High confidence shall not broaden authorization.
- **EVID-REQ-404:** Low confidence shall not permit destruction or suppression without authority.
- **EVID-REQ-405:** Unknown confidence shall remain `Unknown`, not a fabricated numeric value.
- **EVID-REQ-406:** Confidence scales shall be versioned.
- **EVID-REQ-407:** Confidence methods shall be reviewable.
- **EVID-REQ-408:** Confidence changes shall preserve prior assessments.
- **EVID-REQ-409:** Conflicting confidence assessments shall coexist with attribution.
- **EVID-REQ-410:** Model-generated confidence shall receive human review before consequential use.
- **EVID-REQ-411:** Confidence shall not become an autonomous legal conclusion.
- **EVID-REQ-412:** Confidence shall not conceal integrity, provenance or custody defects.
- **EVID-REQ-413:** Confidence records shall minimize personal data.
- **EVID-REQ-414:** Confidence records shall trace to reviewer and evidence.
- **EVID-REQ-415:** Confidence tests shall verify scale boundaries and unknown handling.
- **EVID-REQ-416:** Absence of confidence metadata shall not invalidate an otherwise governed evidence object automatically.

## 14. Classification and Handling Controls

- **EVID-REQ-417:** Every evidence object shall have an approved classification or explicit pending-classification state.
- **EVID-REQ-418:** Classification shall identify the governing classification scheme.
- **EVID-REQ-419:** Classification shall identify owner or authority.
- **EVID-REQ-420:** Classification shall identify handling restrictions.
- **EVID-REQ-421:** Classification shall identify access constraints.
- **EVID-REQ-422:** Classification shall identify export constraints.
- **EVID-REQ-423:** Classification shall identify retention constraints.
- **EVID-REQ-424:** Classification shall identify redaction requirements where applicable.
- **EVID-REQ-425:** Classification shall identify public-release eligibility.
- **EVID-REQ-426:** Unknown classification shall receive restrictive handling.
- **EVID-REQ-427:** Classification shall not imply evidentiary weight.
- **EVID-REQ-428:** Classification shall not imply integrity.
- **EVID-REQ-429:** Classification shall not imply authorization.
- **EVID-REQ-430:** Classification changes shall be attributable.
- **EVID-REQ-431:** Classification downgrades shall require explicit authority.
- **EVID-REQ-432:** Classification upgrades shall trigger access review.
- **EVID-REQ-433:** Classification conflict shall resolve to more restrictive handling until resolved.
- **EVID-REQ-434:** Classification metadata shall not expose restricted content unnecessarily.
- **EVID-REQ-435:** Synthetic and demonstration evidence shall be unmistakably classified.
- **EVID-REQ-436:** Public-source evidence shall not be classified as unrestricted automatically.
- **EVID-REQ-437:** Evidence containing personal data shall identify privacy handling requirements.
- **EVID-REQ-438:** Evidence containing credentials or secrets shall be quarantined and contained.
- **EVID-REQ-439:** Evidence containing malware or unsafe active content shall use isolated handling.
- **EVID-REQ-440:** Evidence from high-risk research environments shall identify isolation requirements.
- **EVID-REQ-441:** Classification shall propagate to derived evidence unless an approved determination states otherwise.
- **EVID-REQ-442:** Classification shall propagate to export packages.
- **EVID-REQ-443:** Classification shall be validated before release or sharing.
- **EVID-REQ-444:** Classification history shall be retained.
- **EVID-REQ-445:** Classification errors shall generate correction records.
- **EVID-REQ-446:** Classification controls shall be tested.

## 15. Access Control and Authorization

- **EVID-REQ-447:** Evidence access shall require explicit authorization.
- **EVID-REQ-448:** Evidence access shall default to deny.
- **EVID-REQ-449:** Evidence access shall evaluate officer, institution, case, purpose, jurisdiction, action, time, tool and data scope where applicable.
- **EVID-REQ-450:** Evidence access shall evaluate classification.
- **EVID-REQ-451:** Evidence access shall evaluate custody and handling state.
- **EVID-REQ-452:** Evidence access shall evaluate retention and legal-hold obligations.
- **EVID-REQ-453:** Evidence access shall evaluate human-approval requirements.
- **EVID-REQ-454:** Evidence access shall identify policy version.
- **EVID-REQ-455:** Evidence access shall enforce least privilege.
- **EVID-REQ-456:** Read access shall remain distinct from modify, transform, export, share and dispose actions.
- **EVID-REQ-457:** Evidence modification shall be prohibited for source objects unless an explicit governed process applies.
- **EVID-REQ-458:** Evidence transformation shall require separate authorization.
- **EVID-REQ-459:** Evidence export shall require separate authorization.
- **EVID-REQ-460:** Evidence deletion or disposition shall require exceptional explicit authorization.
- **EVID-REQ-461:** Cross-case evidence access shall be denied by default.
- **EVID-REQ-462:** Cross-purpose evidence reuse shall be denied by default.
- **EVID-REQ-463:** Cross-jurisdiction access shall require explicit review where applicable.
- **EVID-REQ-464:** Evidence access through connectors shall comply with document 20.
- **EVID-REQ-465:** Evidence access by agents shall preserve officer binding.
- **EVID-REQ-466:** Evidence access by workloads shall use distinct identities where feasible.
- **EVID-REQ-467:** Evidence access shall not rely solely on network location.
- **EVID-REQ-468:** Evidence access denial shall be logged.
- **EVID-REQ-469:** Evidence access approvals shall be attributable.
- **EVID-REQ-470:** Access revocation shall invalidate affected sessions and cached decisions.
- **EVID-REQ-471:** Access-control failure shall not broaden access.
- **EVID-REQ-472:** Unknown evidence identifiers shall be denied.
- **EVID-REQ-473:** Quarantined evidence shall require qualified access.
- **EVID-REQ-474:** Restricted evidence shall require appropriate clearance and purpose.
- **EVID-REQ-475:** Evidence access audit shall correlate request, decision and result.
- **EVID-REQ-476:** Evidence authorization shall conform to `OBDIA-AUTH-001`.
- **EVID-REQ-477:** Evidence access tests shall include field, object, case, purpose and export boundaries.
- **EVID-REQ-478:** Evidence access analytics shall not become excessive surveillance.

## 16. Storage and Preservation

- **EVID-REQ-479:** Evidence storage shall have an identified owner or custodian.
- **EVID-REQ-480:** Evidence storage shall identify storage domain and classification.
- **EVID-REQ-481:** Evidence storage shall preserve object integrity.
- **EVID-REQ-482:** Evidence storage shall preserve metadata integrity.
- **EVID-REQ-483:** Evidence storage shall preserve provenance.
- **EVID-REQ-484:** Evidence storage shall preserve chain of custody.
- **EVID-REQ-485:** Evidence storage shall preserve parent-child lineage.
- **EVID-REQ-486:** Evidence storage shall preserve authorization and audit references.
- **EVID-REQ-487:** Evidence storage shall use access controls proportionate to classification.
- **EVID-REQ-488:** Evidence storage shall use encryption at rest where appropriate.
- **EVID-REQ-489:** Encryption shall not replace authorization.
- **EVID-REQ-490:** Evidence storage shall use secure defaults.
- **EVID-REQ-491:** Evidence storage shall prevent silent overwrite.
- **EVID-REQ-492:** Evidence storage shall use immutable or versioned storage where feasible.
- **EVID-REQ-493:** Evidence storage shall identify replication and backup behavior.
- **EVID-REQ-494:** Evidence backups shall preserve integrity references.
- **EVID-REQ-495:** Evidence backups shall preserve classification and retention.
- **EVID-REQ-496:** Evidence backup access shall be authorized and logged.
- **EVID-REQ-497:** Evidence restoration shall create a governed event.
- **EVID-REQ-498:** Evidence restoration shall verify integrity.
- **EVID-REQ-499:** Evidence storage migration shall create custody and provenance events.
- **EVID-REQ-500:** Evidence storage migration shall verify source and destination integrity.
- **EVID-REQ-501:** Storage failure shall trigger alerting and incident handling.
- **EVID-REQ-502:** Storage corruption shall trigger quarantine.
- **EVID-REQ-503:** Storage deletion shall require authorization and disposition evidence.
- **EVID-REQ-504:** Temporary files, caches and working copies shall be governed.
- **EVID-REQ-505:** Working copies shall not become authoritative evidence silently.
- **EVID-REQ-506:** Storage location shall not establish trust.
- **EVID-REQ-507:** Cloud storage shall identify tenant, account, region and service context where authorized.
- **EVID-REQ-508:** External storage shall identify custody and contractual limitations.
- **EVID-REQ-509:** Evidence preservation checks shall occur at a risk-based cadence.
- **EVID-REQ-510:** Preservation results shall be retained.
- **EVID-REQ-511:** Preservation failures shall remain visible.
- **EVID-REQ-512:** Storage capacity failure shall not silently discard evidence.
- **EVID-REQ-513:** Storage-system changes shall trigger validation.
- **EVID-REQ-514:** Storage implementation shall conform to documents 16, 17, 21, 22 and 24 after consolidation.

## 17. Evidence Review

- **EVID-REQ-515:** Evidence review shall identify reviewer identity and role.
- **EVID-REQ-516:** Evidence review shall identify evidence object and exact version.
- **EVID-REQ-517:** Evidence review shall identify review purpose.
- **EVID-REQ-518:** Evidence review shall identify review time.
- **EVID-REQ-519:** Evidence review shall identify authorization reference.
- **EVID-REQ-520:** Evidence review shall verify mandatory metadata.
- **EVID-REQ-521:** Evidence review shall verify provenance completeness.
- **EVID-REQ-522:** Evidence review shall verify custody history.
- **EVID-REQ-523:** Evidence review shall verify integrity status.
- **EVID-REQ-524:** Evidence review shall verify classification.
- **EVID-REQ-525:** Evidence review shall verify confidence metadata where used.
- **EVID-REQ-526:** Evidence review shall verify known limitations.
- **EVID-REQ-527:** Evidence review shall verify derived lineage where applicable.
- **EVID-REQ-528:** Evidence review shall verify export restrictions.
- **EVID-REQ-529:** Evidence review shall verify retention obligations.
- **EVID-REQ-530:** Evidence review shall identify findings.
- **EVID-REQ-531:** Evidence review shall identify unresolved ambiguity.
- **EVID-REQ-532:** Evidence review shall identify qualified-use conditions.
- **EVID-REQ-533:** Evidence review shall not establish legal admissibility automatically.
- **EVID-REQ-534:** Evidence review shall not establish factual truth automatically.
- **EVID-REQ-535:** Evidence review shall not conceal collection or provenance defects.
- **EVID-REQ-536:** Human review shall be required for material evidence qualification decisions.
- **EVID-REQ-537:** An AI system shall not be final evidence reviewer.
- **EVID-REQ-538:** AI-assisted review shall identify model and prompt context where material.
- **EVID-REQ-539:** Reviewer conflicts of interest shall be disclosed.
- **EVID-REQ-540:** Role concentration shall be disclosed.
- **EVID-REQ-541:** Internal review shall not be represented as independent external assurance.
- **EVID-REQ-542:** Evidence-review records shall be retained.
- **EVID-REQ-543:** Material post-review changes shall invalidate affected review evidence.
- **EVID-REQ-544:** Review corrections shall be additive.

## 18. Export and Evidence Package Model

- **EVID-REQ-545:** Every evidence export shall have an immutable export identifier.
- **EVID-REQ-546:** Every evidence export shall identify export authority.
- **EVID-REQ-547:** Every evidence export shall identify requesting and approving actors.
- **EVID-REQ-548:** Every evidence export shall identify case and purpose.
- **EVID-REQ-549:** Every evidence export shall identify recipient or destination class.
- **EVID-REQ-550:** Every evidence export shall identify included evidence objects.
- **EVID-REQ-551:** Every evidence export shall identify excluded evidence or redactions where material.
- **EVID-REQ-552:** Every evidence export shall identify evidence versions.
- **EVID-REQ-553:** Every evidence export shall identify integrity algorithms and values.
- **EVID-REQ-554:** Every evidence export shall identify provenance references.
- **EVID-REQ-555:** Every evidence export shall identify custody references.
- **EVID-REQ-556:** Every evidence export shall identify classification and handling restrictions.
- **EVID-REQ-557:** Every evidence export shall identify retention and deletion obligations.
- **EVID-REQ-558:** Every evidence export shall identify packaging method.
- **EVID-REQ-559:** Every evidence export shall identify export tool and version.
- **EVID-REQ-560:** Every evidence export shall identify export time.
- **EVID-REQ-561:** Every evidence export shall identify encryption or access-protection method where used.
- **EVID-REQ-562:** Every evidence export shall identify manifest integrity.
- **EVID-REQ-563:** Every evidence export shall identify known limitations.
- **EVID-REQ-564:** Every evidence export shall identify derived-evidence lineage.
- **EVID-REQ-565:** Every evidence export shall identify authorization decision.
- **EVID-REQ-566:** Every evidence export shall identify transfer and acknowledgement events.
- **EVID-REQ-567:** Export shall not broaden authorization.
- **EVID-REQ-568:** Export shall not remove classification or handling obligations silently.
- **EVID-REQ-569:** Export shall not conceal redaction or transformation.
- **EVID-REQ-570:** Export packages shall be verified before transfer.
- **EVID-REQ-571:** Export packages shall be verified after transfer where feasible.
- **EVID-REQ-572:** Export destinations shall be authorized.
- **EVID-REQ-573:** Export failure shall not produce a misleading complete package.
- **EVID-REQ-574:** Partial export shall be identified explicitly.
- **EVID-REQ-575:** Export retries shall create attributable events.
- **EVID-REQ-576:** Export packages shall not contain secrets unrelated to evidence purpose.
- **EVID-REQ-577:** Export metadata shall minimize unnecessary personal data.
- **EVID-REQ-578:** Public release shall require document 15 controls.
- **EVID-REQ-579:** Evidence export shall not imply legal admissibility.
- **EVID-REQ-580:** Export records shall remain linked to later withdrawal or correction.
- **EVID-REQ-581:** Export package formats shall be documented.
- **EVID-REQ-582:** Package-format migration shall preserve manifest and integrity evidence.
- **EVID-REQ-583:** Export testing shall include missing, altered and unauthorized-object cases.
- **EVID-REQ-584:** Export manifests shall remain retrievable.

## 19. Retention, Legal Hold and Disposition

- **EVID-REQ-585:** Every evidence object shall have a retention rule or documented pending determination.
- **EVID-REQ-586:** Retention rules shall identify authority and owner.
- **EVID-REQ-587:** Retention rules shall identify start event.
- **EVID-REQ-588:** Retention rules shall identify duration or review condition.
- **EVID-REQ-589:** Retention rules shall identify legal-hold interaction where applicable.
- **EVID-REQ-590:** Retention rules shall identify privacy and minimization constraints.
- **EVID-REQ-591:** Retention rules shall identify archive and disposal behavior.
- **EVID-REQ-592:** Retention rules shall identify derived-object implications.
- **EVID-REQ-593:** Retention rules shall identify export-package implications.
- **EVID-REQ-594:** Legal hold shall suspend conflicting disposition.
- **EVID-REQ-595:** Legal hold shall have an immutable identifier.
- **EVID-REQ-596:** Legal hold shall identify authority, scope and effective time.
- **EVID-REQ-597:** Legal-hold release shall be attributable.
- **EVID-REQ-598:** Retention expiry shall not trigger silent deletion.
- **EVID-REQ-599:** Disposition shall require explicit authorization.
- **EVID-REQ-600:** Disposition shall identify object, authority, method, time and result.
- **EVID-REQ-601:** Disposition shall verify targeted scope.
- **EVID-REQ-602:** Disposition shall not delete unrelated evidence.
- **EVID-REQ-603:** Disposition shall preserve required provenance and disposition records.
- **EVID-REQ-604:** Disposition shall preserve audit history.
- **EVID-REQ-605:** Disposition shall preserve evidence of integrity status before disposal where required.
- **EVID-REQ-606:** Failed disposition shall remain visible.
- **EVID-REQ-607:** Restoration after disposition shall not occur silently.
- **EVID-REQ-608:** Backups shall follow retention and legal-hold rules.
- **EVID-REQ-609:** Derived evidence shall not outlive parent restrictions without explicit authority.
- **EVID-REQ-610:** Retention changes shall be versioned.
- **EVID-REQ-611:** Retention conflicts shall resolve to the more protective rule until reviewed.
- **EVID-REQ-612:** Excessive retention shall be treated as a privacy risk.
- **EVID-REQ-613:** Premature deletion shall be treated as an evidence-integrity risk.
- **EVID-REQ-614:** Retention automation shall be tested.
- **EVID-REQ-615:** Disposition automation shall fail closed.
- **EVID-REQ-616:** Retention and disposition records shall minimize personal data.
- **EVID-REQ-617:** Case closure shall trigger retention review.
- **EVID-REQ-618:** Project archival shall not convert evidence into public content.
- **EVID-REQ-619:** Retention shall not be represented as legal compliance without evidence.

## 20. Privacy and Fundamental-Rights Controls

- **EVID-REQ-620:** Evidence collection shall be limited to authorized purpose.
- **EVID-REQ-621:** Evidence collection shall be limited to necessary scope.
- **EVID-REQ-622:** Personal data shall be minimized.
- **EVID-REQ-623:** Sensitive personal data shall receive heightened controls.
- **EVID-REQ-624:** Public accessibility shall not be treated as unrestricted processing authority.
- **EVID-REQ-625:** Real-person public investigations shall not be used in demonstrations.
- **EVID-REQ-626:** Evidence metadata shall avoid unnecessary names and identifiers.
- **EVID-REQ-627:** Opaque case and subject identifiers should be used where feasible.
- **EVID-REQ-628:** Evidence review interfaces shall minimize disclosure.
- **EVID-REQ-629:** Evidence exports shall minimize included data.
- **EVID-REQ-630:** Redaction shall be authorized and traceable.
- **EVID-REQ-631:** Redaction shall not fabricate source content.
- **EVID-REQ-632:** Redaction shall preserve required provenance.
- **EVID-REQ-633:** Anonymization or pseudonymization methods shall be documented.
- **EVID-REQ-634:** Anonymized data shall not be represented as risk-free.
- **EVID-REQ-635:** Cross-case analytics shall require approved minimized design.
- **EVID-REQ-636:** Evidence retention shall be proportionate.
- **EVID-REQ-637:** Evidence access monitoring shall not become excessive surveillance.
- **EVID-REQ-638:** Rights-impact findings shall have owners and dispositions.
- **EVID-REQ-639:** Unlawful surveillance shall remain prohibited.
- **EVID-REQ-640:** Unlawful deanonymization shall remain prohibited.
- **EVID-REQ-641:** Automated evidence classification shall not become an autonomous legal decision.
- **EVID-REQ-642:** Human review shall remain required for consequential evidence use.
- **EVID-REQ-643:** Evidence confidence shall not become a proxy for protected attributes.
- **EVID-REQ-644:** Jurisdictional uncertainty shall remain explicit.
- **EVID-REQ-645:** Privacy incidents shall trigger evidence and authorization reassessment.
- **EVID-REQ-646:** Privacy corrections shall preserve audit and provenance history.
- **EVID-REQ-647:** Evidence publication shall comply with documents 04 and 15.
- **EVID-REQ-648:** Compliance mappings shall not establish lawful evidence handling.
- **EVID-REQ-649:** Document 30 remains a forward dependency for complete compliance mapping.

## 21. Quarantine and Failure Handling

- **EVID-REQ-650:** Evidence shall enter quarantine when integrity verification fails.
- **EVID-REQ-651:** Evidence shall enter quarantine when provenance is materially inconsistent.
- **EVID-REQ-652:** Evidence shall enter quarantine when custody is materially unresolved.
- **EVID-REQ-653:** Evidence shall enter quarantine when authorization is invalid or unknown.
- **EVID-REQ-654:** Evidence shall enter quarantine when classification is unknown and risk is material.
- **EVID-REQ-655:** Evidence shall enter quarantine when malware or unsafe active content is detected.
- **EVID-REQ-656:** Evidence shall enter quarantine when secrets or credentials are exposed unexpectedly.
- **EVID-REQ-657:** Evidence shall enter quarantine when source substitution is suspected.
- **EVID-REQ-658:** Evidence shall enter quarantine when storage corruption is detected.
- **EVID-REQ-659:** Evidence shall enter quarantine when export-package integrity fails.
- **EVID-REQ-660:** Quarantine shall have an immutable event identifier.
- **EVID-REQ-661:** Quarantine shall identify reason and triggering evidence.
- **EVID-REQ-662:** Quarantine shall identify responsible owner.
- **EVID-REQ-663:** Quarantine shall identify access restrictions.
- **EVID-REQ-664:** Quarantine shall identify permitted diagnostic actions.
- **EVID-REQ-665:** Quarantine shall identify release criteria.
- **EVID-REQ-666:** Quarantine shall preserve original bytes and metadata where safe.
- **EVID-REQ-667:** Quarantine shall preserve provenance and custody history.
- **EVID-REQ-668:** Quarantine shall not be treated as deletion.
- **EVID-REQ-669:** Quarantined evidence shall not be used as verified evidence.
- **EVID-REQ-670:** Quarantined evidence shall not be exported without explicit exceptional authority.
- **EVID-REQ-671:** Quarantine access shall be logged.
- **EVID-REQ-672:** Release from quarantine shall require human review.
- **EVID-REQ-673:** Release from quarantine shall require successful defined verification.
- **EVID-REQ-674:** Release from quarantine shall create an event.
- **EVID-REQ-675:** Quarantine failure shall trigger incident handling.
- **EVID-REQ-676:** Failure handling shall preserve evidence.
- **EVID-REQ-677:** Failure handling shall not fabricate missing metadata.
- **EVID-REQ-678:** Failure handling shall not broaden access.
- **EVID-REQ-679:** Unknown failure states shall resolve to restrictive handling.
- **EVID-REQ-680:** Tool failure shall not produce a misleading success record.
- **EVID-REQ-681:** Audit failure shall block or restrict privileged evidence operations.
- **EVID-REQ-682:** Clock failure shall be recorded and may require quarantine.
- **EVID-REQ-683:** Revocation-service failure shall restrict evidence access.
- **EVID-REQ-684:** Authorization-service failure shall deny protected evidence operations.
- **EVID-REQ-685:** Storage-service failure shall not silently discard evidence.
- **EVID-REQ-686:** Recovery shall restore a known governed state.
- **EVID-REQ-687:** Recovery shall verify integrity and authorization.
- **EVID-REQ-688:** Repeated failures shall trigger architectural and risk review.
- **EVID-REQ-689:** Quarantine and failure behavior shall be tested.

## 22. Corrections and Disputed Metadata

- **EVID-REQ-690:** Evidence bytes shall not be silently corrected.
- **EVID-REQ-691:** Metadata corrections shall create additive correction records.
- **EVID-REQ-692:** Correction records shall have immutable identifiers.
- **EVID-REQ-693:** Correction records shall identify affected object and field.
- **EVID-REQ-694:** Correction records shall identify prior and corrected values where lawful.
- **EVID-REQ-695:** Correction records shall identify reason.
- **EVID-REQ-696:** Correction records shall identify authority.
- **EVID-REQ-697:** Correction records shall identify time.
- **EVID-REQ-698:** Correction records shall identify supporting evidence.
- **EVID-REQ-699:** Correction records shall identify impact on provenance, custody and exports.
- **EVID-REQ-700:** Correction records shall identify whether re-verification is required.
- **EVID-REQ-701:** Correction records shall preserve original metadata.
- **EVID-REQ-702:** Disputed metadata shall remain visible.
- **EVID-REQ-703:** Conflicting assertions shall be attributable.
- **EVID-REQ-704:** Disputed source information shall not be resolved by silent preference.
- **EVID-REQ-705:** Disputed timestamps shall preserve each source and basis.
- **EVID-REQ-706:** Disputed collector identity shall trigger investigation and possible quarantine.
- **EVID-REQ-707:** Disputed custody events shall trigger qualified handling.
- **EVID-REQ-708:** Disputed integrity values shall trigger re-verification and quarantine.
- **EVID-REQ-709:** Disputed classification shall resolve to restrictive handling.
- **EVID-REQ-710:** Disputed confidence assessments shall coexist with attribution.
- **EVID-REQ-711:** Corrections shall propagate to future exports.
- **EVID-REQ-712:** Past exports shall retain historical manifest and correction notice.
- **EVID-REQ-713:** Material corrections shall trigger reviewer reassessment.
- **EVID-REQ-714:** Material corrections shall trigger affected authorization review.
- **EVID-REQ-715:** Correction automation shall not decide substantive disputes autonomously.
- **EVID-REQ-716:** An AI system shall not be final correction authority.
- **EVID-REQ-717:** Correction records shall be retained.
- **EVID-REQ-718:** Correction access shall be authorized.
- **EVID-REQ-719:** Correction tests shall verify historical preservation.

## 23. Time, Ordering and Clock Integrity

- **EVID-REQ-720:** Evidence timestamps shall identify their source.
- **EVID-REQ-721:** Evidence timestamps shall identify timezone or offset.
- **EVID-REQ-722:** Evidence timestamps shall identify precision where material.
- **EVID-REQ-723:** Evidence timestamps shall identify known clock uncertainty.
- **EVID-REQ-724:** Collector-generated and source-generated timestamps shall remain distinguishable.
- **EVID-REQ-725:** Reception, collection, processing, verification and storage times shall remain distinguishable.
- **EVID-REQ-726:** Unknown time shall remain unknown.
- **EVID-REQ-727:** Estimated time shall be labeled as estimated.
- **EVID-REQ-728:** Normalized time shall preserve source representation.
- **EVID-REQ-729:** Time conversion shall be recorded as a transformation.
- **EVID-REQ-730:** Clock synchronization status shall be recorded where material.
- **EVID-REQ-731:** Clock drift shall be considered in event ordering.
- **EVID-REQ-732:** Conflicting timestamps shall not be silently reconciled.
- **EVID-REQ-733:** Event ordering shall identify its basis.
- **EVID-REQ-734:** Sequence numbers shall not be treated as wall-clock time.
- **EVID-REQ-735:** Blockchain block time shall remain distinguishable from local observation time.
- **EVID-REQ-736:** Cloud-provider event time shall identify provider source.
- **EVID-REQ-737:** Connector receipt time shall remain distinguishable from source event time.
- **EVID-REQ-738:** AI-generated temporal inference shall remain analytical output.
- **EVID-REQ-739:** Time uncertainty shall be included in confidence or limitation metadata where relevant.
- **EVID-REQ-740:** Time-sensitive authorization shall use governed time sources.
- **EVID-REQ-741:** Clock failure shall not produce fabricated timestamps.
- **EVID-REQ-742:** Time corrections shall be additive.
- **EVID-REQ-743:** Time evidence shall be included in provenance.
- **EVID-REQ-744:** Time and ordering limitations shall be visible to reviewers.
- **EVID-REQ-745:** Time-based retention shall identify the controlling timestamp.
- **EVID-REQ-746:** Legal-hold timing shall identify effective time.
- **EVID-REQ-747:** Export manifests shall identify package creation time.
- **EVID-REQ-748:** Timestamp validation shall be tested.
- **EVID-REQ-749:** Clock-manipulation threats shall be included in threat modeling.

## 24. Audit Event Schema

- **EVID-REQ-750:** Every material evidence action shall produce an audit event.
- **EVID-REQ-751:** Evidence audit events shall have immutable identifiers.
- **EVID-REQ-752:** Evidence audit events shall identify evidence object.
- **EVID-REQ-753:** Evidence audit events shall identify event type.
- **EVID-REQ-754:** Evidence audit events shall identify actor identities.
- **EVID-REQ-755:** Evidence audit events shall identify institution.
- **EVID-REQ-756:** Evidence audit events shall identify case and purpose where applicable.
- **EVID-REQ-757:** Evidence audit events shall identify action.
- **EVID-REQ-758:** Evidence audit events shall identify authorization decision.
- **EVID-REQ-759:** Evidence audit events shall identify policy version.
- **EVID-REQ-760:** Evidence audit events shall identify time and ordering information.
- **EVID-REQ-761:** Evidence audit events shall identify source and destination where applicable.
- **EVID-REQ-762:** Evidence audit events shall identify integrity status.
- **EVID-REQ-763:** Evidence audit events shall identify custody impact.
- **EVID-REQ-764:** Evidence audit events shall identify classification impact.
- **EVID-REQ-765:** Evidence audit events shall identify result.
- **EVID-REQ-766:** Evidence audit events shall identify errors and degraded conditions.
- **EVID-REQ-767:** Evidence audit events shall identify related acquisition, transformation, verification, export, correction, quarantine or disposition events.
- **EVID-REQ-768:** Evidence audit events shall not contain secrets.
- **EVID-REQ-769:** Evidence audit events shall minimize personal data.
- **EVID-REQ-770:** Evidence audit events shall be append-only or tamper-evident where feasible.
- **EVID-REQ-771:** Evidence audit events shall remain searchable by authorized reviewers.
- **EVID-REQ-772:** Evidence audit events shall support correlation.
- **EVID-REQ-773:** Evidence audit events shall support incident investigation.
- **EVID-REQ-774:** Evidence audit retention shall follow evidence and governance requirements.
- **EVID-REQ-775:** Audit-service failure shall not permit silent privileged evidence action.
- **EVID-REQ-776:** Audit gaps shall remain visible.
- **EVID-REQ-777:** Audit corrections shall be additive.
- **EVID-REQ-778:** Audit-event schemas shall be versioned.
- **EVID-REQ-779:** Audit-event validation shall be tested.

## 25. Connector and External-System Evidence

- **EVID-REQ-780:** Connector-acquired evidence shall identify connector identity.
- **EVID-REQ-781:** Connector-acquired evidence shall identify connector version.
- **EVID-REQ-782:** Connector-acquired evidence shall identify connector policy version.
- **EVID-REQ-783:** Connector-acquired evidence shall identify destination and source scope.
- **EVID-REQ-784:** Connector-acquired evidence shall identify request and response correlation.
- **EVID-REQ-785:** Connector-acquired evidence shall preserve raw response where lawful and feasible.
- **EVID-REQ-786:** Connector-normalized evidence shall preserve raw-source linkage.
- **EVID-REQ-787:** Connector filters shall be recorded.
- **EVID-REQ-788:** Connector pagination shall be recorded.
- **EVID-REQ-789:** Connector retries shall be recorded.
- **EVID-REQ-790:** Connector rate-limit effects shall be recorded.
- **EVID-REQ-791:** Connector errors and partial results shall be recorded.
- **EVID-REQ-792:** Connector timestamps shall distinguish source and receipt times.
- **EVID-REQ-793:** Connector provenance shall include endpoint or service context without exposing secrets.
- **EVID-REQ-794:** Connector authentication material shall not appear in evidence metadata.
- **EVID-REQ-795:** Connector output shall be treated as untrusted until validated.
- **EVID-REQ-796:** Indirect prompt-injection content shall not alter evidence-handling policy.
- **EVID-REQ-797:** Connector-provided hashes shall remain distinguishable from project-calculated hashes.
- **EVID-REQ-798:** Connector evidence shall not exceed authorized scope.
- **EVID-REQ-799:** Connector evidence shall not broaden case or purpose access.
- **EVID-REQ-800:** Connector disablement shall prevent new acquisition.
- **EVID-REQ-801:** Connector revocation shall invalidate affected sessions.
- **EVID-REQ-802:** Connector compromise shall trigger quarantine and incident review.
- **EVID-REQ-803:** External-system deletion shall not erase preserved evidence history.
- **EVID-REQ-804:** External-system updates shall not overwrite preserved evidence.
- **EVID-REQ-805:** External-system exports shall identify export mechanism.
- **EVID-REQ-806:** Connector supply-chain provenance shall be recorded.
- **EVID-REQ-807:** Connector evidence tests shall include malformed and malicious responses.
- **EVID-REQ-808:** Document 20 remains the forward authority for connector-specific controls.
- **EVID-REQ-809:** Connector evidence handling shall remain safe, authorized, read-only or simulated where required.

## 26. AI-Assisted Evidence Processing

- **EVID-REQ-810:** AI-assisted processing shall identify model and provider where material.
- **EVID-REQ-811:** AI-assisted processing shall identify model version where available.
- **EVID-REQ-812:** AI-assisted processing shall identify prompt or task template version where material.
- **EVID-REQ-813:** AI-assisted processing shall identify input evidence.
- **EVID-REQ-814:** AI-assisted processing shall identify output artifact.
- **EVID-REQ-815:** AI-assisted processing shall identify human reviewer where required.
- **EVID-REQ-816:** AI-assisted outputs shall be classified as derived analytical outputs unless formally incorporated through a governed process.
- **EVID-REQ-817:** AI-assisted outputs shall not be represented as source evidence.
- **EVID-REQ-818:** AI-assisted outputs shall preserve lineage.
- **EVID-REQ-819:** AI-assisted outputs shall identify uncertainty and limitations.
- **EVID-REQ-820:** AI-assisted processing shall not broaden authorization.
- **EVID-REQ-821:** AI-assisted processing shall not access unrelated cases.
- **EVID-REQ-822:** AI-assisted processing shall not persist memory outside approved scope.
- **EVID-REQ-823:** AI-assisted processing shall not directly alter original evidence.
- **EVID-REQ-824:** AI-assisted processing shall not directly authorize export, deletion or disposition.
- **EVID-REQ-825:** AI-assisted processing shall not make final admissibility decisions.
- **EVID-REQ-826:** AI-assisted processing shall not make autonomous legal conclusions.
- **EVID-REQ-827:** AI-assisted processing shall be resilient to prompt injection from evidence content.
- **EVID-REQ-828:** Evidence content instructions shall not override system or policy authority.
- **EVID-REQ-829:** Model output shall be validated before consequential use.
- **EVID-REQ-830:** AI-assisted transformation shall record nondeterminism where relevant.
- **EVID-REQ-831:** Repeatability limitations shall be documented.
- **EVID-REQ-832:** Model changes shall trigger affected validation.
- **EVID-REQ-833:** Provider changes shall trigger trust and privacy reassessment.
- **EVID-REQ-834:** AI-generated confidence shall remain distinct from human confidence assessment.
- **EVID-REQ-835:** AI-assisted review shall not be represented as independent external review.
- **EVID-REQ-836:** AI risk shall trace to document 26 after consolidation.
- **EVID-REQ-837:** AI-assisted evidence tests shall use synthetic or authorized controlled fixtures.
- **EVID-REQ-838:** Unsafe model behavior shall support containment.
- **EVID-REQ-839:** AI-assisted processing shall preserve human accountability.

## 27. Demonstration and Synthetic Evidence

- **EVID-REQ-840:** Synthetic evidence shall be unmistakably labeled.
- **EVID-REQ-841:** Synthetic evidence shall not reuse real credentials.
- **EVID-REQ-842:** Synthetic evidence shall not use real victims, suspects or case subjects.
- **EVID-REQ-843:** Synthetic evidence shall not be represented as operational evidence.
- **EVID-REQ-844:** Mock-service evidence shall identify the mock boundary.
- **EVID-REQ-845:** Testnet evidence shall identify the test network.
- **EVID-REQ-846:** Local-laboratory evidence shall identify the laboratory boundary.
- **EVID-REQ-847:** Simulated Dark Web evidence shall identify the simulation boundary.
- **EVID-REQ-848:** Archived-authorized evidence shall identify archive authority.
- **EVID-REQ-849:** Lawful public-source evidence shall identify collection and use limitations.
- **EVID-REQ-850:** Demonstration evidence shall identify purpose and audience.
- **EVID-REQ-851:** Demonstration evidence shall not contain real case data.
- **EVID-REQ-852:** Demonstration evidence shall not interact with uncontrolled criminal infrastructure.
- **EVID-REQ-853:** Demonstration evidence shall not enable offensive operations.
- **EVID-REQ-854:** Demonstration exports shall remain controlled.
- **EVID-REQ-855:** Demonstration screenshots shall not imply production deployment.
- **EVID-REQ-856:** Demonstration chain of custody may be simplified only when clearly labeled and authorized.
- **EVID-REQ-857:** Demonstration integrity controls shall remain testable.
- **EVID-REQ-858:** Demonstration provenance shall remain complete enough for reproducibility.
- **EVID-REQ-859:** Demonstration confidence labels shall not imply factual findings.
- **EVID-REQ-860:** Demonstration cleanup shall be documented.
- **EVID-REQ-861:** Demonstration retention shall be defined.
- **EVID-REQ-862:** Demonstration artifacts shall follow publication policy.
- **EVID-REQ-863:** Demonstration findings shall not be generalized to real investigations without evidence.
- **EVID-REQ-864:** Public demonstrations shall pass document 15 release controls.
- **EVID-REQ-865:** Synthetic-data generation shall avoid realistic sensitive identifiers.
- **EVID-REQ-866:** Demonstration evidence shall preserve safe defaults.
- **EVID-REQ-867:** Unsafe fixtures shall be quarantined.
- **EVID-REQ-868:** Demonstration environments shall be isolated.
- **EVID-REQ-869:** Demonstration evidence validation shall be retained.

## 28. Testing and Validation

- **EVID-REQ-870:** Every material evidence requirement shall have testable acceptance criteria.
- **EVID-REQ-871:** Evidence testing shall use synthetic or authorized controlled fixtures.
- **EVID-REQ-872:** Tests shall verify unique identifier enforcement.
- **EVID-REQ-873:** Tests shall verify duplicate identifier rejection.
- **EVID-REQ-874:** Tests shall verify mandatory field validation.
- **EVID-REQ-875:** Tests shall verify source and collection-method capture.
- **EVID-REQ-876:** Tests shall verify acquisition-event creation.
- **EVID-REQ-877:** Tests shall verify integrity calculation.
- **EVID-REQ-878:** Tests shall verify successful integrity verification.
- **EVID-REQ-879:** Tests shall verify altered-object detection.
- **EVID-REQ-880:** Tests shall verify substituted-object detection.
- **EVID-REQ-881:** Tests shall verify custody transfer.
- **EVID-REQ-882:** Tests shall verify refused or failed custody transfer.
- **EVID-REQ-883:** Tests shall verify custody-gap handling.
- **EVID-REQ-884:** Tests shall verify provenance completeness.
- **EVID-REQ-885:** Tests shall verify derived lineage.
- **EVID-REQ-886:** Tests shall verify transformation event capture.
- **EVID-REQ-887:** Tests shall verify AI-generated output classification.
- **EVID-REQ-888:** Tests shall verify confidence separation from integrity.
- **EVID-REQ-889:** Tests shall verify classification enforcement.
- **EVID-REQ-890:** Tests shall verify default-deny evidence access.
- **EVID-REQ-891:** Tests shall verify cross-case isolation.
- **EVID-REQ-892:** Tests shall verify cross-purpose isolation.
- **EVID-REQ-893:** Tests shall verify revocation behavior.
- **EVID-REQ-894:** Tests shall verify quarantine triggers.
- **EVID-REQ-895:** Tests shall verify release-from-quarantine controls.
- **EVID-REQ-896:** Tests shall verify export manifest completeness.
- **EVID-REQ-897:** Tests shall verify export integrity before and after transfer.
- **EVID-REQ-898:** Tests shall verify retention and legal-hold behavior.
- **EVID-REQ-899:** Tests shall verify disposition authorization.
- **EVID-REQ-900:** Tests shall verify correction history.
- **EVID-REQ-901:** Tests shall verify timestamp-source capture.
- **EVID-REQ-902:** Tests shall verify audit-event correlation.
- **EVID-REQ-903:** Tests shall verify connector malformed-response handling.
- **EVID-REQ-904:** Tests shall verify prompt-injection resistance.
- **EVID-REQ-905:** Tests shall verify storage corruption handling.
- **EVID-REQ-906:** Tests shall verify backup restoration and integrity.
- **EVID-REQ-907:** Tests shall verify audit-service failure handling.
- **EVID-REQ-908:** Tests shall identify exact source, policy, tool, connector, model, configuration and environment versions.
- **EVID-REQ-909:** Failed tests shall remain visible.
- **EVID-REQ-910:** Passing tests shall not be generalized beyond scope.
- **EVID-REQ-911:** Automated tests shall not replace human evidence review.
- **EVID-REQ-912:** Document 22 remains a forward dependency for complete testing governance.
- **EVID-REQ-913:** Validation evidence shall be retained in governed repository locations.
- **EVID-REQ-914:** Material implementation changes shall invalidate affected validation evidence.

## 29. Evidence Traceability

- **EVID-REQ-915:** Every evidence object shall trace to its acquisition or creation event.
- **EVID-REQ-916:** Every acquisition event shall trace to authorization.
- **EVID-REQ-917:** Every evidence object shall trace to provenance records.
- **EVID-REQ-918:** Every evidence object shall trace to custody records.
- **EVID-REQ-919:** Every integrity value shall trace to calculation and verification events.
- **EVID-REQ-920:** Every derived artifact shall trace to parent evidence.
- **EVID-REQ-921:** Every transformation shall trace to method and tool.
- **EVID-REQ-922:** Every confidence assessment shall trace to assessor and method.
- **EVID-REQ-923:** Every access event shall trace to authorization and enforcement.
- **EVID-REQ-924:** Every export shall trace to included evidence and approval.
- **EVID-REQ-925:** Every retention rule shall trace to authority.
- **EVID-REQ-926:** Every disposition shall trace to authorization.
- **EVID-REQ-927:** Every correction shall trace to prior metadata and supporting evidence.
- **EVID-REQ-928:** Every quarantine event shall trace to trigger and disposition.
- **EVID-REQ-929:** Every review shall trace to exact evidence object and state.
- **EVID-REQ-930:** Every incident shall trace to affected evidence.
- **EVID-REQ-931:** Every risk and exception shall trace to affected requirements and evidence.
- **EVID-REQ-932:** Every release shall trace to evidence-package validation where applicable.
- **EVID-REQ-933:** Traceability shall be bidirectional.
- **EVID-REQ-934:** Broken traceability affecting authorization, integrity, provenance, custody, privacy or release shall be blocking.
- **EVID-REQ-935:** Planned evidence controls shall not be represented as implemented.
- **EVID-REQ-936:** Implemented controls shall not be represented as validated without evidence.
- **EVID-REQ-937:** Superseded schemas shall identify successors.
- **EVID-REQ-938:** Deprecated evidence types and fields shall identify replacements.
- **EVID-REQ-939:** Traceability shall use immutable identifiers.
- **EVID-REQ-940:** Traceability records shall not contain secrets.
- **EVID-REQ-941:** Traceability records shall minimize personal and case data.
- **EVID-REQ-942:** Cross-reference validation shall occur before approval.
- **EVID-REQ-943:** Baseline freeze shall validate evidence traceability across documents 01–30.
- **EVID-REQ-944:** No document above 30 shall gain normative evidence authority by implication.

## 30. Evidence Review Gate

- **EVID-REQ-945:** Evidence-model review shall identify the exact document and commit.
- **EVID-REQ-946:** Review shall verify preservation of all original evidence fields.
- **EVID-REQ-947:** Review shall verify the evidence object model.
- **EVID-REQ-948:** Review shall verify acquisition-event requirements.
- **EVID-REQ-949:** Review shall verify provenance requirements.
- **EVID-REQ-950:** Review shall verify integrity-verification requirements.
- **EVID-REQ-951:** Review shall verify chain-of-custody requirements.
- **EVID-REQ-952:** Review shall verify derived-evidence lineage.
- **EVID-REQ-953:** Review shall verify access and authorization controls.
- **EVID-REQ-954:** Review shall verify storage and retention controls.
- **EVID-REQ-955:** Review shall verify export-package requirements.
- **EVID-REQ-956:** Review shall verify confidence separation from integrity.
- **EVID-REQ-957:** Review shall verify privacy and fundamental-rights controls.
- **EVID-REQ-958:** Review shall verify quarantine and correction behavior.
- **EVID-REQ-959:** Review shall verify AI-assisted processing boundaries.
- **EVID-REQ-960:** Review shall verify testing and traceability.
- **EVID-REQ-961:** Review shall identify residual risks and forward dependencies.
- **EVID-REQ-962:** Critical findings shall block progression.
- **EVID-REQ-963:** Material post-review change shall invalidate affected review evidence.
- **EVID-REQ-964:** Role concentration shall be disclosed.
- **EVID-REQ-965:** Internal review shall not be represented as independent external assurance.
- **EVID-REQ-966:** Security Reviewer shall assess integrity, custody, authorization and failure handling.
- **EVID-REQ-967:** Privacy and Governance Reviewer shall assess minimization, retention, rights and publication impacts.
- **EVID-REQ-968:** Implementation Reviewer shall assess implementation alignment where implementation exists.
- **EVID-REQ-969:** Research Reviewer shall assess source, method, claims and uncertainty where research evidence is involved.
- **EVID-REQ-970:** Release Reviewer shall assess export and publication evidence where release is proposed.
- **EVID-REQ-971:** Project Founder approval shall not substitute for required specialist review.
- **EVID-REQ-972:** Automated findings shall remain advisory until adopted by an authorized human reviewer.
- **EVID-REQ-973:** Review non-applicability shall have rationale.
- **EVID-REQ-974:** Approval for Draft incorporation shall not establish legal admissibility or normative approval.

## 31. Minimum Validation Checklist

Before approval of this model or an evidence implementation, confirm:

- [ ] unique immutable evidence identifier exists;
- [ ] source is recorded;
- [ ] collection method is recorded;
- [ ] timestamp and time-source limitations are recorded;
- [ ] collector and accountable institution are recorded;
- [ ] case, purpose, jurisdiction and authorization context are recorded;
- [ ] integrity algorithm and value are recorded;
- [ ] integrity verification has attributable evidence;
- [ ] provenance is complete or gaps are explicit;
- [ ] chain of custody is chronological and verifiable;
- [ ] classification and handling restrictions are defined;
- [ ] confidence, where used, has a method and is distinct from integrity;
- [ ] reviewer information is recorded where review occurred;
- [ ] original and derived evidence remain distinguishable;
- [ ] transformations preserve lineage;
- [ ] access defaults to deny;
- [ ] cross-case and cross-purpose isolation is enforced;
- [ ] storage, backup and restoration preserve integrity;
- [ ] export metadata and manifest are complete;
- [ ] retention, legal hold and disposition are governed;
- [ ] corrections preserve history;
- [ ] quarantine behavior is fail-closed;
- [ ] AI-generated content is not represented as original evidence;
- [ ] privacy and rights impacts are reviewed;
- [ ] tests cover alteration, substitution, custody, lineage, export, retention and failure;
- [ ] forward dependencies 19–30 are recorded where applicable;
- [ ] Project Founder approval exists before status becomes Approved.


## 32. Limitations

- This model does not establish legal admissibility or evidentiary weight.
- It does not certify forensic tools or collection methods.
- It does not select a storage platform, evidence management system, hash library, signing technology, connector, model provider or export format.
- It does not authorize evidence acquisition.
- It does not replace the Authorization Model.
- It does not replace the Agent Lifecycle Model or Connector Security Policy.
- It does not prove authenticity, truth, lawful collection or provenance completeness.
- Cryptographic integrity cannot establish factual correctness.
- Confidence metadata cannot compensate for broken provenance, custody or integrity.
- Internal evidence review is not independent certification.
- Documents 19–30 remain forward dependencies where they govern lifecycle, connectors, testing, structure, versioning, risk, exceptions, change and compliance.


## 33. Change Control

Every material change shall identify rationale, affected evidence types, fields, events, schemas, dependencies, security and privacy impact, chain-of-custody impact, retention impact, migration, validation, rollback, authority and version effect.

- **EVID-REQ-975:** Editorial corrections shall use a patch version when meaning is unchanged.
- **EVID-REQ-976:** Backward-compatible substantive additions shall use a minor version.
- **EVID-REQ-977:** Incompatible evidence-schema or lifecycle changes shall use a major version.
- **EVID-REQ-978:** Material evidence-architecture decisions shall require an ADR where applicable.
- **EVID-REQ-979:** Changes shall require Project Founder approval.
- **EVID-REQ-980:** Changes shall receive Security Reviewer assessment.
- **EVID-REQ-981:** Privacy, authorization, implementation and release impacts shall receive specialist review where applicable.
- **EVID-REQ-982:** Changes shall identify affected evidence objects, events, exports, tests and implementation.
- **EVID-REQ-983:** Changes shall include schema migration and backward-compatibility analysis.
- **EVID-REQ-984:** Changes shall include validation and rollback analysis.
- **EVID-REQ-985:** Changes shall not retroactively fabricate provenance, custody, integrity or review evidence.
- **EVID-REQ-986:** Historical evidence records shall not be silently rewritten.
- **EVID-REQ-987:** Identifier reuse shall remain prohibited.
- **EVID-REQ-988:** Migration shall preserve old and new schema references.
- **EVID-REQ-989:** Forward-dependency reconciliation shall occur before this model becomes Approved.

## 34. Consolidation Record

Version 1.0.0 consolidates the two existing Evidence Model variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-EVID-001`;
- normalizes the authoritative filename to `18_EVIDENCE_MODEL.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves identifier, source, collection method, timestamp, integrity hash, collector, provenance, chain of custody, integrity verification, classification, confidence, reviewer and export metadata;
- distinguishes confidence from integrity, authenticity, legality, admissibility and truth;
- adds evidence object types, acquisition events, handling states, provenance, verification events, custody transfers, derived-artifact lineage, transformations, authorization, storage, review, export packages, retention, privacy, quarantine, corrections, time integrity, audit, connectors, AI-assisted processing, demonstrations, validation and traceability;
- identifies documents 19–30 as forward dependencies where they govern lifecycle, connectors, implementation, testing, diagrams, repository structure, versioning, risk, exceptions, change and compliance;
- treats Enterprise and non-Enterprise legacy files as source variants of the same immutable document;
- creates no evidence, legal-admissibility claim, forensic certification, investigative authority or operational implementation.

## 35. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Security Reviewer; approval reserved to Project Founder | Constitutional consolidation of both legacy evidence outlines: preserved all original fields; clarified confidence; added complete object, event, provenance, integrity, custody, derivation, access, export, retention, quarantine, validation and traceability requirements. |
