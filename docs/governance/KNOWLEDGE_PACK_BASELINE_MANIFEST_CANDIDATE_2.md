# KNOWLEDGE PACK BASELINE MANIFEST — CANDIDATE 2

| Metadata Field | Value |
|---|---|
| **Record ID** | `VER-MANIFEST-002` |
| **Manifest Version** | `0.1.0` |
| **Status** | Draft |
| **Classification** | Governance record |
| **Candidate Baseline Version** | `1.0.0-candidate.2` |
| **Baseline Disposition** | Candidate; not frozen |
| **Predecessor Record** | `VER-MANIFEST-001` |
| **Predecessor Baseline Version** | `1.0.0-candidate.1` |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Record** | None; approval has not been granted |
| **Evaluated Repository Commit** | `622706ac050e96aacba12eab27fbc2e78b198342` |
| **Evaluated Repository Tree** | `ed077ab1d2f395c02cf2147660f8012d547368bb` |
| **Recorded At** | `2026-08-06T21:03:45Z` |
| **Member Count** | `30` |
| **Requirement Definitions** | `15703` |
| **Machine Reference Result** | PASS |
| **Human Semantic Review** | REQUIRED |
| **Baseline Freeze** | BLOCKED |
| **Evidence Record** | `TEST-EVID-002` |
| **Supersession** | None; this is a subsequent candidate record |

## 1. Purpose

This manifest records the exact candidate Knowledge Pack member set evaluated after dependency-filename reconciliation.

It does not amend or overwrite `VER-MANIFEST-001`. Candidate.1 remains an immutable historical record of its own evaluated commit and tree.

This record does not approve, implement, validate, publish or freeze the baseline. Individual document lifecycle states remain authoritative and independent.

## 2. Candidate Lineage

| Field | Candidate 1 | Candidate 2 |
|---|---|---|
| Manifest Record | `VER-MANIFEST-001` | `VER-MANIFEST-002` |
| Baseline Version | `1.0.0-candidate.1` | `1.0.0-candidate.2` |
| Evaluated Commit | `32abb5f4ddf202ee168c9faa02ff56faaee20df1` | `622706ac050e96aacba12eab27fbc2e78b198342` |
| Evaluated Tree | `791ff25649c1c9efa51269bffeeba8c8c4977a6f` | `ed077ab1d2f395c02cf2147660f8012d547368bb` |
| Machine Reference Result | Not executed in candidate.1 record | PASS after adjudication |
| Freeze State | BLOCKED | BLOCKED |

Candidate.2 is a later candidate snapshot. It does not retroactively alter candidate.1.

## 3. Governance Anchor

| Field | Value |
|---|---|
| **Governance Anchor ID** | `OBDIA-CONST-001` |
| **Title** | MASTER DOCUMENTATION CONSTITUTION |
| **Version** | `1.0.0` |
| **Status** | `Approved Constitutional Baseline` |
| **Authoritative Path** | `docs/governance/MASTER_DOCUMENTATION_CONSTITUTION.md` |
| **SHA-256** | `ba013979fefd589ccfb505ea5fcbc2c9020ddd91df4c12f838683d83b7bf46bf` |

The Constitution is the governance anchor and is not a numbered Knowledge Pack member.

## 4. Candidate Baseline Members

| No. | Document ID | Title | Version | Status | Classification | Authoritative Path | Requirement Definitions | SHA-256 |
|---:|---|---|---|---|---|---|---:|---|
| 01 | `OBDIA-CANON-001` | Canonical Project Definition | `1.0.0` | `Approved Baseline` | Not stated in source metadata | `docs/governance/01_CANONICAL_PROJECT_DEFINITION.md` | `0` | `06d99f0f22276cda863445c91f2c1bd591def9dedcd287f1ee17639e2ebdd5d5` |
| 02 | `OBDIA-GLOSS-001` | PROJECT GLOSSARY | `1.1.0` | `Draft` | Normative | `docs/knowledge/02_PROJECT_GLOSSARY.md` | `27` | `76d12d06a34d0ec4c2ef83854c058057c30f33ffbae323501cad82d31f31495c` |
| 03 | `OBDIA-RES-001` | RESEARCH AND SOURCE POLICY | `1.1.0` | `Draft` | Normative | `docs/knowledge/03_RESEARCH_AND_SOURCE_POLICY.md` | `51` | `66e08b81f37162263e9804696cd67ab0481e9dbbc40ef296219ab2685cfc0358` |
| 04 | `OBDIA-PUB-001` | PORTFOLIO AND PUBLICATION POLICY | `1.1.0` | `Draft` | Normative | `docs/knowledge/04_PORTFOLIO_AND_PUBLICATION_POLICY.md` | `54` | `84a2f170b5b415fe601bbc3620081254456e2828d8dc2802a62ecbe835424579` |
| 05 | `OBDIA-ARCH-001` | ARCHITECTURE PRINCIPLES | `1.1.0` | `Draft` | Normative | `docs/knowledge/05_ARCHITECTURE_PRINCIPLES.md` | `50` | `3be5310f2155722fd94124eea50d281b7833bae019dc445aadb09a5e7dd118cc` |
| 06 | `OBDIA-ID-001` | IDENTITY AND DELEGATION MODEL | `1.1.0` | `Draft` | Normative | `docs/knowledge/06_IDENTITY_AND_DELEGATION_MODEL.md` | `112` | `e688525b5fdace7543116192b8aee35dc04ef7f6de1aa3bdbddf1fe8d1f50a89` |
| 07 | `OBDIA-TM-001` | THREAT MODEL BASELINE | `1.1.0` | `Draft` | Normative | `docs/knowledge/07_THREAT_MODEL_BASELINE.md` | `96` | `db0ef1e515b7209fa2387212934176c97d4042542010554c382e6cce3ff868e3` |
| 08 | `OBDIA-TRUST-001` | TRUST MODEL | `1.1.0` | `Draft` | Normative | `docs/knowledge/08_TRUST_MODEL.md` | `144` | `8a72d0bf73717d6cb2754582184581e122326493b92b09913073dcc9898f8b40` |
| 09 | `OBDIA-GOV-001` | GOVERNANCE MODEL | `1.1.0` | `Draft` | Normative | `docs/knowledge/09_GOVERNANCE_MODEL.md` | `129` | `4b4d074973b13c3de9419ceb56dd2695ee64dbd9681192cc38e7c92923d8efa8` |
| 10 | `OBDIA-ADR-001` | ARCHITECTURE DECISION RECORD (ADR) POLICY | `1.1.0` | `Draft` | Normative | `docs/knowledge/10_ADR_POLICY.md` | `126` | `67296bd5b8c734012e25502b0ec54ba5ab4e711905adbbe8e4f1553e531c2878` |
| 11 | `OBDIA-DOC-001` | DOCUMENTATION STANDARD | `1.1.0` | `Draft` | Normative | `docs/knowledge/11_DOCUMENTATION_STANDARD.md` | `182` | `1a940a0c0cb486955e0c55fcae65f19274b603c8a1a2a8afee3cc33b703cc6b7` |
| 12 | `OBDIA-IMP-001` | IMPLEMENTATION POLICY | `1.0.0` | `Draft` | Normative | `docs/knowledge/12_IMPLEMENTATION_POLICY.md` | `208` | `f5900f62b3136923649b560d3ee873fedd862be2c0d41189189b47c07424051c` |
| 13 | `OBDIA-STATE-001` | REPOSITORY STATE MODEL | `1.0.0` | `Draft` | Normative | `docs/knowledge/13_REPOSITORY_STATE_MODEL.md` | `190` | `72ed5cf10ed63de0d67a9f6c6c3e1cbf9590df765c613bb86d06bb9d192f5163` |
| 14 | `OBDIA-NAME-001` | TERMINOLOGY AND NAMING | `1.0.0` | `Draft` | Normative | `docs/knowledge/14_TERMINOLOGY_AND_NAMING.md` | `266` | `938e35425d353ee843c758949b85a7a5e7218b75f6f5f4237dbc7f9eded65069` |
| 15 | `OBDIA-REL-001` | RELEASE AND PUBLICATION POLICY | `1.0.1` | `Draft` | Normative | `docs/knowledge/15_RELEASE_AND_PUBLICATION_POLICY.md` | `317` | `195454653f16eb636f61fce980a25e7081525f3ba05683798f8eb600e3ac5f96` |
| 16 | `OBDIA-SEC-001` | SECURITY ARCHITECTURE BASELINE | `1.0.1` | `Draft` | Normative | `docs/knowledge/16_SECURITY_ARCHITECTURE_BASELINE.md` | `427` | `0d58affc8a0e5c928728fda7f725d489daa3e44e3c55b00d588d14de3cd88fe2` |
| 17 | `OBDIA-AUTH-001` | AUTHORIZATION MODEL | `1.0.1` | `Draft` | Normative | `docs/knowledge/17_AUTHORIZATION_MODEL.md` | `652` | `5791ed7565d5adb70c600d3e520d19f9f6cd4cb72c0647ea2c1d16f5f4d692b6` |
| 18 | `OBDIA-EVID-001` | EVIDENCE MODEL | `1.0.1` | `Draft` | Normative | `docs/knowledge/18_EVIDENCE_MODEL.md` | `989` | `8f62480b69ae125ff4fab4b1e49f72c048eb0265226335671031a7d615ccbc47` |
| 19 | `OBDIA-AGENT-001` | AGENT LIFECYCLE MODEL | `1.0.1` | `Draft` | Normative | `docs/knowledge/19_AGENT_LIFECYCLE_MODEL.md` | `883` | `b56c2acfcee6cccd3de1322b6900ddee96866b50555a65423ff3a6bd56e7fc51` |
| 20 | `OBDIA-CONN-001` | CONNECTOR SECURITY POLICY | `1.0.1` | `Draft` | Normative | `docs/knowledge/20_CONNECTOR_SECURITY_POLICY.md` | `988` | `45b8534cc6d8c7af29f9ef62f7066b822cce07d0af76f0318e7d3db0cedf0453` |
| 21 | `OBDIA-CODE-001` | SECURE CODING STANDARD | `1.0.1` | `Draft` | Normative | `docs/knowledge/21_SECURE_CODING_STANDARD.md` | `965` | `c68842690a8d78549a56a7be28e76bf462045995d13b5727549b8e24763cd496` |
| 22 | `OBDIA-TEST-001` | TESTING STANDARD | `1.0.1` | `Draft` | Normative | `docs/knowledge/22_TESTING_STANDARD.md` | `999` | `f9e7480c05a41cf00d2ee49dc28bee284f4d52bca8f0a7d4c0428f61aad2de59` |
| 23 | `OBDIA-DIAG-001` | DIAGRAM STANDARD | `1.0.1` | `Draft` | Normative | `docs/knowledge/23_DIAGRAM_STANDARD.md` | `855` | `76bbf52126a9ddf86825eb45dc6261f13a5d81a57bd222f462ec286dcfd10705` |
| 24 | `OBDIA-GH-001` | GITHUB REPOSITORY STANDARD | `1.0.1` | `Draft` | Normative | `docs/knowledge/24_GITHUB_REPOSITORY_STANDARD.md` | `999` | `c17f46532b4c04d795d75e6c295a5b38b1fdc27632297313d589ed35a4c14ffc` |
| 25 | `OBDIA-VER-001` | DOCUMENT VERSIONING POLICY | `1.0.1` | `Draft` | Normative | `docs/knowledge/25_DOCUMENT_VERSIONING_POLICY.md` | `999` | `938a536b344238c4ad0d53dae39f51886a9b155a1c62322a6e62c2c9af4383a7` |
| 26 | `OBDIA-RISK-001` | AI RISK REGISTER | `1.0.0` | `Draft` | Normative | `docs/knowledge/26_AI_RISK_REGISTER.md` | `999` | `59b90294c15803144d55b9ccda516eebd4c71e9e4086872f2c75587c524af33e` |
| 27 | `OBDIA-RSCH-001` | RESEARCH BACKLOG | `1.0.0` | `Draft` | Normative for backlog-governance rules; individual backlog items are Research records | `docs/knowledge/27_RESEARCH_BACKLOG.md` | `999` | `c77472c77bd05dabfac970565aab05dfc55c37c73fae403c791fd24899965a30` |
| 28 | `OBDIA-ASM-001` | ASSUMPTIONS REGISTER | `1.0.0` | `Draft` | Normative for assumption-governance rules; individual assumptions are Research records until validated and adopted through applicable governance | `docs/knowledge/28_ASSUMPTIONS_REGISTER.md` | `999` | `efb27950f2ff189cf29b4129cf58140021e1d74e05aa827e0ee078837aba764f` |
| 29 | `OBDIA-DEC-001` | DECISION LOG POLICY | `1.0.0` | `Draft` | Normative | `docs/knowledge/29_DECISION_LOG_POLICY.md` | `999` | `c45b33f9bcbb89f2af5e305c460d4b03ed66a1ddd2d7c77629285376955705df` |
| 30 | `OBDIA-COMP-001` | COMPLIANCE MAPPING BASELINE | `1.0.0` | `Draft` | Normative for mapping method, source governance, evidence requirements and claim boundaries; individual mappings remain design and review evidence | `docs/knowledge/30_COMPLIANCE_MAPPING_BASELINE.md` | `999` | `12c72c6237280f7aa48ccd25ab42bc63d38e3766f8a015382961893bbddee8fe` |

## 5. Validation Findings

| Check | Result | Boundary |
|---|---|---|
| Numbered member inventory | PASS | Exactly documents 01–30 are represented |
| Candidate member count | PASS | `30` members |
| Requirement-definition uniqueness | PASS | `15703` unique definitions |
| Primary Document ID uniqueness | PASS | No duplicate primary Document IDs |
| Constitution source uniqueness | PASS | One authoritative `OBDIA-CONST-001` source |
| Dependency filename reconciliation | PASS | No legacy dependency filename remains |
| Documents 15–25 version markers | PASS | All are `1.0.1` |
| Required Document ID references | PASS | No unresolved required reference |
| Requirement references | PASS | No unresolved requirement reference |
| Strong Markdown paths | PASS | No unresolved or ambiguous strong path |
| Numbered documents above 30 | PASS | None detected |
| Candidate.1 integrity | PASS | Manifest and evidence hashes unchanged |
| Machine semantic-reference precheck | PASS | After validator-defect adjudication |
| Unknown Document-ID-shaped advisory tokens | PASS | `0` observed |
| Non-authoritative unresolved path advisories | RECORDED | `73` observations |
| Self-dependency advisories | PASS | `0` observations |
| Human semantic compatibility review | REQUIRED | Not replaced by machine validation |
| Baseline-freeze decision | BLOCKED | No explicit approval exists |

## 6. Lifecycle and Version Distribution

- Member status distribution: Approved Baseline=1; Draft=29.
- Member version distribution: 1.0.0=9; 1.0.1=11; 1.1.0=10.
- Inclusion in this manifest does not change any member status.
- Documents 15–25 remain Draft after their patch-version reconciliation.

## 7. Dependencies

This candidate manifest depends on:

- `OBDIA-CANON-001`;
- `OBDIA-CONST-001`;
- Knowledge documents 02–30;
- `OBDIA-VER-001`;
- `OBDIA-TEST-001`;
- `OBDIA-GH-001`;
- `VER-MANIFEST-001`;
- `TEST-EVID-001`;
- `TEST-EVID-002`.

## 8. Exclusions

The following are not candidate baseline members:

- the Constitution;
- historical consolidation execution records;
- candidate manifests and evidence records;
- ADRs;
- implementation artifacts;
- test implementations;
- repository settings;
- generated diagrams;
- releases.

No Knowledge document 31 is created or implied.

## 9. Unresolved Gates

Before any baseline-freeze decision, the project still requires:

1. human review of semantic compatibility;
2. normative-dependency appropriateness review;
3. security and privacy blocking-finding review;
4. rights-impact and legal-applicability review where relevant;
5. approval-record completeness;
6. explicit Project Founder baseline-freeze decision.

## 10. Known Limitations

- Machine reference consistency does not establish semantic compatibility.
- Hash integrity does not establish correctness or legal sufficiency.
- Draft documents remain Draft.
- The 73 non-authoritative path observations remain advisory and require contextual review.
- No implementation or control effectiveness is demonstrated.
- Compliance mappings remain design and review evidence.
- Internal review is not independent external assurance.
- No release, certification, legal-compliance or institutional-adoption claim is created.

## 11. Approval Boundary

This is a Draft candidate manifest.

It shall not be described as Approved, Validated, Published or frozen until the remaining governance gates and explicit approval evidence exist.
