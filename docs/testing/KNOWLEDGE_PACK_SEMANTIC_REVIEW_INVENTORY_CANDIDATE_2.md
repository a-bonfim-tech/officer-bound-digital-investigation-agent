# Candidate.2 Semantic Review Inventory

| Field | Value |
|---|---|
| Repository commit | `290a84bb85aef357186bc8fa097da05df4c45f32` |
| Repository tree | `3bf8c42419bc17b75a5a5f41edb5d6c26c4bcf92` |
| Canonical identity | `OBDIA-CANON-001` |
| Constitution identity | `OBDIA-CONST-001` |
| Candidate members | `30` |
| Documents scanned | `31` |
| Raw dependency references | `800` |
| Unique dependency edges | `448` |
| Dependency cycles | `2` |
| Blocking resolution findings | `0` |
| Non-authoritative path advisories | `73` |
| Inventory result | **PASS** |
| Human semantic review | **REQUIRED** |
| Baseline freeze | **BLOCKED** |

## Dependency Direction

- `BACKWARD`: `322`
- `FORWARD`: `96`
- `GOVERNANCE_ANCHOR`: `30`

## Cycle Components

### Cycle 1

- `docs/knowledge/02_PROJECT_GLOSSARY.md`
- `docs/knowledge/03_RESEARCH_AND_SOURCE_POLICY.md`
- `docs/knowledge/04_PORTFOLIO_AND_PUBLICATION_POLICY.md`

### Cycle 2

- `docs/knowledge/12_IMPLEMENTATION_POLICY.md`
- `docs/knowledge/13_REPOSITORY_STATE_MODEL.md`
- `docs/knowledge/14_TERMINOLOGY_AND_NAMING.md`
- `docs/knowledge/15_RELEASE_AND_PUBLICATION_POLICY.md`
- `docs/knowledge/16_SECURITY_ARCHITECTURE_BASELINE.md`
- `docs/knowledge/17_AUTHORIZATION_MODEL.md`
- `docs/knowledge/18_EVIDENCE_MODEL.md`
- `docs/knowledge/19_AGENT_LIFECYCLE_MODEL.md`
- `docs/knowledge/20_CONNECTOR_SECURITY_POLICY.md`
- `docs/knowledge/21_SECURE_CODING_STANDARD.md`
- `docs/knowledge/22_TESTING_STANDARD.md`
- `docs/knowledge/23_DIAGRAM_STANDARD.md`
- `docs/knowledge/24_GITHUB_REPOSITORY_STANDARD.md`
- `docs/knowledge/25_DOCUMENT_VERSIONING_POLICY.md`
- `docs/knowledge/26_AI_RISK_REGISTER.md`
- `docs/knowledge/27_RESEARCH_BACKLOG.md`
- `docs/knowledge/28_ASSUMPTIONS_REGISTER.md`
- `docs/knowledge/29_DECISION_LOG_POLICY.md`
- `docs/knowledge/30_COMPLIANCE_MAPPING_BASELINE.md`

## Human Review Items

| Review ID | Source | Relation | Target | Direction | Cycle | Disposition |
|---|---|---|---|---|---|---|
| `SEM-EDGE-0001` | `OBDIA-CONST-001` | Dependencies | `OBDIA-CANON-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0002` | `OBDIA-GLOSS-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0003` | `OBDIA-GLOSS-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0004` | `OBDIA-GLOSS-001` | Dependencies, Normative References | `OBDIA-RES-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0005` | `OBDIA-RES-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0006` | `OBDIA-RES-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0007` | `OBDIA-RES-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0008` | `OBDIA-RES-001` | Cross-References | `OBDIA-PUB-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0009` | `OBDIA-PUB-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0010` | `OBDIA-PUB-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0011` | `OBDIA-PUB-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0012` | `OBDIA-PUB-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0013` | `OBDIA-ARCH-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0014` | `OBDIA-ARCH-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0015` | `OBDIA-ARCH-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0016` | `OBDIA-ARCH-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0017` | `OBDIA-ID-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0018` | `OBDIA-ID-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0019` | `OBDIA-ID-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0020` | `OBDIA-ID-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0021` | `OBDIA-ID-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0022` | `OBDIA-TM-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0023` | `OBDIA-TM-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0024` | `OBDIA-TM-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0025` | `OBDIA-TM-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0026` | `OBDIA-TM-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0027` | `OBDIA-TM-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0028` | `OBDIA-TRUST-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0029` | `OBDIA-TRUST-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0030` | `OBDIA-TRUST-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0031` | `OBDIA-TRUST-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0032` | `OBDIA-TRUST-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0033` | `OBDIA-TRUST-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0034` | `OBDIA-TRUST-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0035` | `OBDIA-GOV-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0036` | `OBDIA-GOV-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0037` | `OBDIA-GOV-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0038` | `OBDIA-GOV-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0039` | `OBDIA-GOV-001` | Dependencies, Normative References | `OBDIA-PUB-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0040` | `OBDIA-GOV-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0041` | `OBDIA-GOV-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0042` | `OBDIA-GOV-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0043` | `OBDIA-GOV-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0044` | `OBDIA-ADR-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0045` | `OBDIA-ADR-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0046` | `OBDIA-ADR-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0047` | `OBDIA-ADR-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0048` | `OBDIA-ADR-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0049` | `OBDIA-ADR-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0050` | `OBDIA-ADR-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0051` | `OBDIA-ADR-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0052` | `OBDIA-ADR-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0053` | `OBDIA-DOC-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0054` | `OBDIA-DOC-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0055` | `OBDIA-DOC-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0056` | `OBDIA-DOC-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0057` | `OBDIA-DOC-001` | Dependencies, Normative References | `OBDIA-PUB-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0058` | `OBDIA-DOC-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0059` | `OBDIA-DOC-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0060` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0061` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0062` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0063` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0064` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-PUB-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0065` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0066` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0067` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0068` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0069` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0070` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0071` | `OBDIA-IMP-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0072` | `OBDIA-IMP-001` | Dependencies | `OBDIA-SEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0073` | `OBDIA-IMP-001` | Dependencies | `OBDIA-AUTH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0074` | `OBDIA-IMP-001` | Dependencies | `OBDIA-EVID-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0075` | `OBDIA-IMP-001` | Dependencies | `OBDIA-CONN-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0076` | `OBDIA-STATE-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0077` | `OBDIA-STATE-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0078` | `OBDIA-STATE-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0079` | `OBDIA-STATE-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0080` | `OBDIA-STATE-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0081` | `OBDIA-NAME-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0082` | `OBDIA-NAME-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0083` | `OBDIA-NAME-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0084` | `OBDIA-NAME-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0085` | `OBDIA-NAME-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0086` | `OBDIA-NAME-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0087` | `OBDIA-NAME-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0088` | `OBDIA-REL-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0089` | `OBDIA-REL-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0090` | `OBDIA-REL-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0091` | `OBDIA-REL-001` | Dependencies, Normative References | `OBDIA-PUB-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0092` | `OBDIA-REL-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0093` | `OBDIA-REL-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0094` | `OBDIA-REL-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0095` | `OBDIA-REL-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0096` | `OBDIA-REL-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0097` | `OBDIA-REL-001` | Dependencies | `OBDIA-TEST-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0098` | `OBDIA-REL-001` | Dependencies | `OBDIA-GH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0099` | `OBDIA-REL-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0100` | `OBDIA-REL-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0101` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0102` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0103` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0104` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0105` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0106` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0107` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0108` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0109` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0110` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0111` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0112` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0113` | `OBDIA-SEC-001` | Dependencies, Normative References | `OBDIA-REL-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0114` | `OBDIA-SEC-001` | Dependencies | `OBDIA-AUTH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0115` | `OBDIA-SEC-001` | Dependencies | `OBDIA-EVID-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0116` | `OBDIA-SEC-001` | Dependencies | `OBDIA-AGENT-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0117` | `OBDIA-SEC-001` | Dependencies | `OBDIA-CONN-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0118` | `OBDIA-SEC-001` | Dependencies | `OBDIA-TEST-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0119` | `OBDIA-SEC-001` | Dependencies | `OBDIA-RISK-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0120` | `OBDIA-SEC-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0121` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0122` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0123` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0124` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0125` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0126` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0127` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0128` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0129` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0130` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0131` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0132` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0133` | `OBDIA-AUTH-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0134` | `OBDIA-AUTH-001` | Dependencies | `OBDIA-EVID-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0135` | `OBDIA-AUTH-001` | Dependencies | `OBDIA-AGENT-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0136` | `OBDIA-AUTH-001` | Dependencies | `OBDIA-CONN-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0137` | `OBDIA-AUTH-001` | Dependencies | `OBDIA-TEST-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0138` | `OBDIA-AUTH-001` | Dependencies | `OBDIA-RISK-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0139` | `OBDIA-AUTH-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0140` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0141` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0142` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0143` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0144` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0145` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0146` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0147` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0148` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0149` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0150` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0151` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0152` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0153` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0154` | `OBDIA-EVID-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0155` | `OBDIA-EVID-001` | Dependencies | `OBDIA-AGENT-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0156` | `OBDIA-EVID-001` | Dependencies | `OBDIA-CONN-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0157` | `OBDIA-EVID-001` | Dependencies | `OBDIA-TEST-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0158` | `OBDIA-EVID-001` | Dependencies | `OBDIA-DIAG-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0159` | `OBDIA-EVID-001` | Dependencies | `OBDIA-GH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0160` | `OBDIA-EVID-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0161` | `OBDIA-EVID-001` | Dependencies | `OBDIA-RISK-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0162` | `OBDIA-EVID-001` | Dependencies | `OBDIA-RSCH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0163` | `OBDIA-EVID-001` | Dependencies | `OBDIA-ASM-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0164` | `OBDIA-EVID-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0165` | `OBDIA-EVID-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0166` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0167` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0168` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0169` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0170` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0171` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0172` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0173` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0174` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0175` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0176` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0177` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0178` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0179` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0180` | `OBDIA-AGENT-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0181` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-CONN-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0182` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-CODE-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0183` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-TEST-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0184` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-GH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0185` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0186` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-RISK-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0187` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-RSCH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0188` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-ASM-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0189` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0190` | `OBDIA-AGENT-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0191` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0192` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0193` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0194` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-ID-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0195` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0196` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0197` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0198` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0199` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0200` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0201` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0202` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0203` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-REL-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0204` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0205` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0206` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0207` | `OBDIA-CONN-001` | Dependencies, Normative References | `OBDIA-AGENT-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0208` | `OBDIA-CONN-001` | Dependencies | `OBDIA-CODE-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0209` | `OBDIA-CONN-001` | Dependencies | `OBDIA-TEST-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0210` | `OBDIA-CONN-001` | Dependencies | `OBDIA-GH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0211` | `OBDIA-CONN-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0212` | `OBDIA-CONN-001` | Dependencies | `OBDIA-RISK-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0213` | `OBDIA-CONN-001` | Dependencies | `OBDIA-RSCH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0214` | `OBDIA-CONN-001` | Dependencies | `OBDIA-ASM-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0215` | `OBDIA-CONN-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0216` | `OBDIA-CONN-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0217` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0218` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0219` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0220` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0221` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0222` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0223` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0224` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0225` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0226` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0227` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0228` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0229` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-REL-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0230` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0231` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0232` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0233` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-AGENT-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0234` | `OBDIA-CODE-001` | Dependencies, Normative References | `OBDIA-CONN-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0235` | `OBDIA-CODE-001` | Dependencies | `OBDIA-TEST-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0236` | `OBDIA-CODE-001` | Dependencies | `OBDIA-GH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0237` | `OBDIA-CODE-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0238` | `OBDIA-CODE-001` | Dependencies | `OBDIA-RISK-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0239` | `OBDIA-CODE-001` | Dependencies | `OBDIA-RSCH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0240` | `OBDIA-CODE-001` | Dependencies | `OBDIA-ASM-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0241` | `OBDIA-CODE-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0242` | `OBDIA-CODE-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0243` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0244` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0245` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0246` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0247` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0248` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0249` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0250` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-REL-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0251` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0252` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0253` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0254` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-AGENT-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0255` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-CONN-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0256` | `OBDIA-TEST-001` | Dependencies, Normative References | `OBDIA-CODE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0257` | `OBDIA-TEST-001` | Dependencies | `OBDIA-DIAG-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0258` | `OBDIA-TEST-001` | Dependencies | `OBDIA-GH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0259` | `OBDIA-TEST-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0260` | `OBDIA-TEST-001` | Dependencies | `OBDIA-RISK-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0261` | `OBDIA-TEST-001` | Dependencies | `OBDIA-RSCH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0262` | `OBDIA-TEST-001` | Dependencies | `OBDIA-ASM-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0263` | `OBDIA-TEST-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0264` | `OBDIA-TEST-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0265` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0266` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0267` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-GLOSS-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0268` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0269` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0270` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0271` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0272` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0273` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0274` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0275` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0276` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0277` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0278` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-AGENT-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0279` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-CONN-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0280` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-CODE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0281` | `OBDIA-DIAG-001` | Dependencies, Normative References | `OBDIA-TEST-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0282` | `OBDIA-DIAG-001` | Dependencies | `OBDIA-GH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0283` | `OBDIA-DIAG-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0284` | `OBDIA-DIAG-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0285` | `OBDIA-DIAG-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0286` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0287` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0288` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-PUB-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0289` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0290` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0291` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0292` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0293` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0294` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0295` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-REL-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0296` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0297` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0298` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-CONN-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0299` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-CODE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0300` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-TEST-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0301` | `OBDIA-GH-001` | Dependencies, Normative References | `OBDIA-DIAG-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0302` | `OBDIA-GH-001` | Dependencies | `OBDIA-VER-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0303` | `OBDIA-GH-001` | Dependencies | `OBDIA-RISK-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0304` | `OBDIA-GH-001` | Dependencies | `OBDIA-RSCH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0305` | `OBDIA-GH-001` | Dependencies | `OBDIA-ASM-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0306` | `OBDIA-GH-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0307` | `OBDIA-GH-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0308` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0309` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0310` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0311` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0312` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0313` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0314` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0315` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-REL-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0316` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0317` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-TEST-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0318` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-DIAG-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0319` | `OBDIA-VER-001` | Dependencies, Normative References | `OBDIA-GH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0320` | `OBDIA-VER-001` | Dependencies | `OBDIA-RISK-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0321` | `OBDIA-VER-001` | Dependencies | `OBDIA-RSCH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0322` | `OBDIA-VER-001` | Dependencies | `OBDIA-ASM-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0323` | `OBDIA-VER-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0324` | `OBDIA-VER-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0325` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0326` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0327` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0328` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0329` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0330` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0331` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-REL-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0332` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0333` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0334` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0335` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-AGENT-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0336` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-CONN-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0337` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-CODE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0338` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-TEST-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0339` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-GH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0340` | `OBDIA-RISK-001` | Dependencies, Normative References | `OBDIA-VER-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0341` | `OBDIA-RISK-001` | Dependencies | `OBDIA-RSCH-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0342` | `OBDIA-RISK-001` | Dependencies | `OBDIA-ASM-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0343` | `OBDIA-RISK-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0344` | `OBDIA-RISK-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0345` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0346` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0347` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0348` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0349` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0350` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0351` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0352` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0353` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0354` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0355` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0356` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0357` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-AGENT-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0358` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-CONN-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0359` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-CODE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0360` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-TEST-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0361` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-DIAG-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0362` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-GH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0363` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-VER-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0364` | `OBDIA-RSCH-001` | Dependencies, Normative References | `OBDIA-RISK-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0365` | `OBDIA-RSCH-001` | Dependencies | `OBDIA-ASM-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0366` | `OBDIA-RSCH-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0367` | `OBDIA-RSCH-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0368` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0369` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0370` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0371` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0372` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0373` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0374` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0375` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0376` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0377` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0378` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0379` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0380` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0381` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0382` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-AGENT-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0383` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-CONN-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0384` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-CODE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0385` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-TEST-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0386` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-DIAG-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0387` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-GH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0388` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-VER-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0389` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-RISK-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0390` | `OBDIA-ASM-001` | Dependencies, Normative References | `OBDIA-RSCH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0391` | `OBDIA-ASM-001` | Dependencies | `OBDIA-DEC-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0392` | `OBDIA-ASM-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0393` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0394` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0395` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0396` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-PUB-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0397` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0398` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0399` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0400` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0401` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0402` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0403` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0404` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0405` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0406` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-REL-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0407` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0408` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0409` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0410` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-AGENT-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0411` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-CONN-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0412` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-CODE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0413` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-TEST-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0414` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-DIAG-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0415` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-GH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0416` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-VER-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0417` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-RISK-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0418` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-RSCH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0419` | `OBDIA-DEC-001` | Dependencies, Normative References | `OBDIA-ASM-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0420` | `OBDIA-DEC-001` | Dependencies | `OBDIA-COMP-001` | FORWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0421` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-CANON-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0422` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-CONST-001` | GOVERNANCE_ANCHOR | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0423` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-RES-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0424` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-PUB-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0425` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-ARCH-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0426` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-TM-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0427` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-TRUST-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0428` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-GOV-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0429` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-ADR-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0430` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-DOC-001` | BACKWARD | No | PENDING HUMAN REVIEW |
| `SEM-EDGE-0431` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-IMP-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0432` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-STATE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0433` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-NAME-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0434` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-REL-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0435` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-SEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0436` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-AUTH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0437` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-EVID-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0438` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-AGENT-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0439` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-CONN-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0440` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-CODE-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0441` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-TEST-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0442` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-DIAG-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0443` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-GH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0444` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-VER-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0445` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-RISK-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0446` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-RSCH-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0447` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-ASM-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |
| `SEM-EDGE-0448` | `OBDIA-COMP-001` | Dependencies, Normative References | `OBDIA-DEC-001` | BACKWARD | Yes | PENDING HUMAN REVIEW |

## Required Human Disposition

Each edge requires review of scope, authority, lifecycle, security, privacy, rights impact, legal applicability and normative appropriateness.

Inventory completion does not change document status and does not authorize baseline freeze.
