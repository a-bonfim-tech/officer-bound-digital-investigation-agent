# GITHUB REPOSITORY STANDARD

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-GH-001 |
| **Title** | GitHub Repository Standard |
| **Version** | 1.0.1 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define the mandatory structure, ownership, access, branch, pull-request, workflow, documentation, code, test, diagram, evidence, dependency, secret, release, archival and validation requirements for the OBDIA GitHub repository. |
| **Scope** | Repository root, documentation, Knowledge Pack 01–30, ADRs, diagrams, source code, tests, fixtures, workflows, issue and pull-request templates, security reporting, dependency metadata, generated artifacts, releases, tags, branches, repository settings and retained governance evidence. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `04_PORTFOLIO_AND_PUBLICATION_POLICY.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `15_RELEASE_AND_PUBLICATION_POLICY.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `18_EVIDENCE_MODEL.md`; `20_CONNECTOR_SECURITY_POLICY.md`; `21_SECURE_CODING_STANDARD.md`; `22_TESTING_STANDARD.md`; `23_DIAGRAM_STANDARD.md`; forward dependencies `25_DOCUMENT_VERSIONING_POLICY.md`, `26_AI_RISK_REGISTER.md`, `27_RESEARCH_BACKLOG.md`, `28_ASSUMPTIONS_REGISTER.md`, `29_DECISION_LOG_POLICY.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-PUB-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-REL-001`; `OBDIA-SEC-001`; `OBDIA-EVID-001`; `OBDIA-CONN-001`; `OBDIA-CODE-001`; `OBDIA-TEST-001`; `OBDIA-DIAG-001` |
| **Cross-References** | Documents 25–30; accepted ADRs; repository settings; branch protection or rulesets; CODEOWNERS; workflows; issue and pull-request records; dependency manifests; build and test evidence; releases; risk, exception and change records |
| **Assumptions** | The repository is a professional research repository that may remain private during development and may later publish selected approved artifacts. |
| **Constraints** | Repository structure and automation shall not grant independent legal authority to an AI agent, authorize prohibited conduct, expose secrets or real case data, bypass human review, misrepresent implementation state or depend on uncontrolled criminal infrastructure. |
| **Security Considerations** | Repository compromise can alter normative requirements, inject malicious code or workflows, expose secrets, falsify evidence, weaken authorization, corrupt releases, erase review history or create misleading public claims. |
| **Validation Criteria** | The repository contains every mandatory root artifact and governed directory, uses stable ownership and review controls, isolates generated and authoritative sources, constrains automation, protects secrets, preserves immutable history and provides machine-checkable structural and traceability evidence. |
| **Implementation Relationship** | This standard governs repository organization and controls. It does not activate a production system, select a hosting plan, prove repository settings are currently enforced, or establish public-release approval. |

---

## 1. Purpose

This document consolidates the OBDIA GitHub Repository Standard.

The legacy sources required the repository to include:

- `README.md`;
- `SECURITY.md`;
- `LICENSE`;
- `CONTRIBUTING.md`;
- `docs/`;
- `adr/`;
- `diagrams/`;
- `tests/`;
- continuous integration through `.github/workflows/`;
- versioning.

They also required the repository structure to reflect architectural dependencies.

This consolidation preserves every original artifact and requirement and defines the repository controls necessary for secure, auditable and professionally credible research publication.

The repository is an evidence-bearing engineering environment. Repository presence, merge, tag or workflow success does not by itself prove approval, implementation, validation, compliance or production readiness.


## 2. Fundamental Repository Rules

- **GH-REQ-001:** The repository shall include `README.md`.
- **GH-REQ-002:** The repository shall include `SECURITY.md`.
- **GH-REQ-003:** The repository shall include a repository-root license file named `LICENSE` or a platform-compatible license filename.
- **GH-REQ-004:** The repository shall include `CONTRIBUTING.md`.
- **GH-REQ-005:** The repository shall include `docs/`.
- **GH-REQ-006:** The repository shall include `adr/`.
- **GH-REQ-007:** The repository shall include `diagrams/`.
- **GH-REQ-008:** The repository shall include `tests/`.
- **GH-REQ-009:** The repository shall include governed continuous-integration workflows under `.github/workflows/` when automation is introduced.
- **GH-REQ-010:** The repository shall define versioning.
- **GH-REQ-011:** The repository structure shall reflect architectural dependencies.
- **GH-REQ-012:** The repository shall preserve the authority hierarchy defined by the Constitution.
- **GH-REQ-013:** The repository shall preserve Knowledge Pack numbering 01–30.
- **GH-REQ-014:** The repository shall not create a normative Knowledge document numbered above 30 without constitutional amendment.
- **GH-REQ-015:** The repository shall preserve immutable document and requirement identifiers.
- **GH-REQ-016:** The repository shall distinguish normative content from informative content.
- **GH-REQ-017:** The repository shall distinguish proposed architecture from implementation.
- **GH-REQ-018:** The repository shall distinguish implementation from testing.
- **GH-REQ-019:** The repository shall distinguish testing from validation.
- **GH-REQ-020:** The repository shall distinguish internal review from independent external assurance.
- **GH-REQ-021:** The repository shall not contain real case data for public demonstrations.
- **GH-REQ-022:** The repository shall not contain reusable credentials or secrets.
- **GH-REQ-023:** The repository shall not contain unlawful or unauthorized access tooling.
- **GH-REQ-024:** The repository shall not depend on uncontrolled criminal infrastructure.
- **GH-REQ-025:** The repository shall preserve human ownership and accountability.
- **GH-REQ-026:** An AI system shall not approve its own repository changes.
- **GH-REQ-027:** An AI system shall not merge or release privileged changes without authorized human control.
- **GH-REQ-028:** Unknown repository status shall not be represented as approved.
- **GH-REQ-029:** Critical repository-integrity failures shall block merge or release.
- **GH-REQ-030:** Repository history shall not be silently rewritten.

## 3. Canonical Minimum Layout

```text
.
├── README.md
├── SECURITY.md
├── LICENSE
├── CONTRIBUTING.md
├── docs/
│   ├── governance/
│   ├── knowledge/
│   ├── architecture/
│   ├── security/
│   ├── implementation/
│   ├── testing/
│   ├── research/
│   └── publication/
├── adr/
├── diagrams/
│   ├── source/
│   └── rendered/
├── tests/
│   ├── unit/
│   ├── integration/
│   ├── security/
│   ├── fixtures/
│   └── evidence/
└── .github/
    ├── workflows/
    ├── ISSUE_TEMPLATE/
    ├── PULL_REQUEST_TEMPLATE.md
    └── CODEOWNERS
```

- **GH-REQ-031:** The canonical minimum layout shall be treated as a target structure.
- **GH-REQ-032:** Directories not yet applicable may be absent only with documented rationale.
- **GH-REQ-033:** Directories shall not imply implemented content when empty.
- **GH-REQ-034:** Placeholder files shall not be used to create misleading implementation claims.
- **GH-REQ-035:** Repository-root files shall remain easy to discover.
- **GH-REQ-036:** Normative Knowledge documents shall reside under `docs/knowledge/`.
- **GH-REQ-037:** Constitutional and governance authority documents shall reside under `docs/governance/`.
- **GH-REQ-038:** Architecture documentation shall reside under `docs/architecture/` when separated from Knowledge documents.
- **GH-REQ-039:** Security documentation shall reside under `docs/security/` when separated from Knowledge documents.
- **GH-REQ-040:** Implementation documentation shall reside under `docs/implementation/`.
- **GH-REQ-041:** Testing documentation and validation packages shall reside under `docs/testing/` or governed test-evidence paths.
- **GH-REQ-042:** Research notes and source analysis shall reside under `docs/research/`.
- **GH-REQ-043:** Publication and portfolio records shall reside under `docs/publication/`.
- **GH-REQ-044:** ADRs shall reside under `adr/`.
- **GH-REQ-045:** Authoritative diagram sources shall reside under `diagrams/source/` when this split is adopted.
- **GH-REQ-046:** Rendered diagrams shall reside under `diagrams/rendered/` when this split is adopted.
- **GH-REQ-047:** Tests shall reside under `tests/` or language-standard test locations mapped by repository documentation.
- **GH-REQ-048:** Test fixtures shall be distinguishable from test code.
- **GH-REQ-049:** Test evidence shall be distinguishable from mutable temporary output.
- **GH-REQ-050:** Workflow files shall reside under `.github/workflows/`.
- **GH-REQ-051:** Issue templates shall reside under `.github/ISSUE_TEMPLATE/` when enabled.
- **GH-REQ-052:** The pull-request template shall reside at a platform-supported governed path.
- **GH-REQ-053:** CODEOWNERS shall reside at a platform-supported governed path.
- **GH-REQ-054:** Implementation source directories shall be added only when implementation exists.
- **GH-REQ-055:** Structural migrations shall preserve history and traceability.

## 4. Root README

- **GH-REQ-056:** `README.md` shall identify the project name.
- **GH-REQ-057:** `README.md` shall state the project purpose.
- **GH-REQ-058:** `README.md` shall state that the AI agent is bound to one authorized human investigator.
- **GH-REQ-059:** `README.md` shall state that the human remains legally and operationally responsible.
- **GH-REQ-060:** `README.md` shall state that the AI agent has no independent legal authority.
- **GH-REQ-061:** `README.md` shall state the independent research or institutional status accurately.
- **GH-REQ-062:** `README.md` shall distinguish proposed architecture from implementation.
- **GH-REQ-063:** `README.md` shall distinguish tests from validation.
- **GH-REQ-064:** `README.md` shall identify safe demonstration boundaries.
- **GH-REQ-065:** `README.md` shall prohibit unauthorized access and offensive operations.
- **GH-REQ-066:** `README.md` shall distinguish Deep Web from Dark Web when both are mentioned.
- **GH-REQ-067:** `README.md` shall link to the canonical project definition.
- **GH-REQ-068:** `README.md` shall link to the Constitution.
- **GH-REQ-069:** `README.md` shall link to the Knowledge Pack index.
- **GH-REQ-070:** `README.md` shall link to architecture documentation.
- **GH-REQ-071:** `README.md` shall link to threat and trust models.
- **GH-REQ-072:** `README.md` shall link to security reporting.
- **GH-REQ-073:** `README.md` shall link to contribution guidance.
- **GH-REQ-074:** `README.md` shall identify the license.
- **GH-REQ-075:** `README.md` shall identify current repository maturity without overstatement.
- **GH-REQ-076:** `README.md` shall identify synthetic, mock, testnet and local-laboratory assumptions.
- **GH-REQ-077:** `README.md` shall not claim production readiness without evidence.
- **GH-REQ-078:** `README.md` shall not claim certification or compliance without evidence.
- **GH-REQ-079:** `README.md` shall not expose secrets or sensitive operational details.
- **GH-REQ-080:** `README.md` shall remain synchronized with material repository status.

## 5. Security Policy

- **GH-REQ-081:** `SECURITY.md` shall define a safe vulnerability-reporting channel.
- **GH-REQ-082:** `SECURITY.md` shall define supported versions or current support scope.
- **GH-REQ-083:** `SECURITY.md` shall define expected report contents.
- **GH-REQ-084:** `SECURITY.md` shall prohibit public disclosure of active secrets.
- **GH-REQ-085:** `SECURITY.md` shall prohibit testing without authorization.
- **GH-REQ-086:** `SECURITY.md` shall prohibit interaction with real criminal infrastructure.
- **GH-REQ-087:** `SECURITY.md` shall prohibit social engineering of real persons.
- **GH-REQ-088:** `SECURITY.md` shall prohibit destructive or privacy-invasive testing.
- **GH-REQ-089:** `SECURITY.md` shall define handling of accidental sensitive-data discovery.
- **GH-REQ-090:** `SECURITY.md` shall define handling of exposed credentials.
- **GH-REQ-091:** `SECURITY.md` shall identify response expectations without guaranteeing unavailable capacity.
- **GH-REQ-092:** `SECURITY.md` shall distinguish vulnerability reporting from general support.
- **GH-REQ-093:** `SECURITY.md` shall identify coordinated-disclosure expectations where applicable.
- **GH-REQ-094:** `SECURITY.md` shall not publish exploit instructions unnecessarily.
- **GH-REQ-095:** `SECURITY.md` shall not claim a bug-bounty program unless one exists.
- **GH-REQ-096:** `SECURITY.md` shall not claim service-level commitments without authority.
- **GH-REQ-097:** `SECURITY.md` shall identify that repository research is non-operational unless changed by governance.
- **GH-REQ-098:** `SECURITY.md` shall link to canonical safety boundaries.
- **GH-REQ-099:** `SECURITY.md` shall be reviewed after security incidents.
- **GH-REQ-100:** `SECURITY.md` shall be included in release review.
- **GH-REQ-101:** `SECURITY.md` shall remain accessible from the repository interface.
- **GH-REQ-102:** `SECURITY.md` shall not contain private contact credentials.
- **GH-REQ-103:** `SECURITY.md` shall use a monitored and authorized contact path.
- **GH-REQ-104:** `SECURITY.md` changes shall receive security review.
- **GH-REQ-105:** Security-report records shall remain confidential according to classification.

## 6. License and Intellectual Property

- **GH-REQ-106:** The repository shall have an explicit license decision.
- **GH-REQ-107:** The license file shall use a recognized or legally reviewed text.
- **GH-REQ-108:** The repository shall not imply a license that is absent.
- **GH-REQ-109:** License scope shall identify whether code, documentation and data share the same terms.
- **GH-REQ-110:** Third-party content shall preserve required notices.
- **GH-REQ-111:** Third-party diagrams and assets shall preserve required attribution.
- **GH-REQ-112:** Third-party code shall preserve license obligations.
- **GH-REQ-113:** Generated content shall not be assumed free of third-party rights.
- **GH-REQ-114:** AI-assisted content shall receive provenance and license review where material.
- **GH-REQ-115:** Dataset licenses shall be documented.
- **GH-REQ-116:** Test-fixture licenses shall be documented where external material is used.
- **GH-REQ-117:** Archived-source permissions shall be documented.
- **GH-REQ-118:** Public-source accessibility shall not be treated as permission to republish.
- **GH-REQ-119:** License incompatibilities shall block distribution.
- **GH-REQ-120:** Unknown license status shall block public release.
- **GH-REQ-121:** License changes shall receive governance and release review.
- **GH-REQ-122:** License changes shall not be retroactively represented without legal basis.
- **GH-REQ-123:** Forked or vendored dependencies shall preserve notices.
- **GH-REQ-124:** SBOM or dependency records shall identify applicable licenses where feasible.
- **GH-REQ-125:** The README shall link to the license.
- **GH-REQ-126:** Release artifacts shall include required license notices.
- **GH-REQ-127:** Private repository status shall not substitute for license analysis.
- **GH-REQ-128:** License scanners may support but shall not replace human review.
- **GH-REQ-129:** License exceptions shall have owners and rationale.
- **GH-REQ-130:** Intellectual-property limitations shall remain visible.

## 7. Contribution Policy

- **GH-REQ-131:** `CONTRIBUTING.md` shall define the contribution workflow.
- **GH-REQ-132:** `CONTRIBUTING.md` shall define branch expectations.
- **GH-REQ-133:** `CONTRIBUTING.md` shall define commit expectations.
- **GH-REQ-134:** `CONTRIBUTING.md` shall define pull-request expectations.
- **GH-REQ-135:** `CONTRIBUTING.md` shall define review expectations.
- **GH-REQ-136:** `CONTRIBUTING.md` shall define test expectations.
- **GH-REQ-137:** `CONTRIBUTING.md` shall define documentation expectations.
- **GH-REQ-138:** `CONTRIBUTING.md` shall define security and privacy boundaries.
- **GH-REQ-139:** `CONTRIBUTING.md` shall prohibit secrets and real case data.
- **GH-REQ-140:** `CONTRIBUTING.md` shall prohibit unauthorized or offensive contributions.
- **GH-REQ-141:** `CONTRIBUTING.md` shall define use of synthetic fixtures.
- **GH-REQ-142:** `CONTRIBUTING.md` shall define issue and requirement references.
- **GH-REQ-143:** `CONTRIBUTING.md` shall define status and claim accuracy.
- **GH-REQ-144:** `CONTRIBUTING.md` shall define generated and AI-assisted content review.
- **GH-REQ-145:** `CONTRIBUTING.md` shall define license and provenance responsibilities.
- **GH-REQ-146:** `CONTRIBUTING.md` shall define how to report security issues privately.
- **GH-REQ-147:** `CONTRIBUTING.md` shall identify required local validation.
- **GH-REQ-148:** `CONTRIBUTING.md` shall identify formatting and naming rules.
- **GH-REQ-149:** `CONTRIBUTING.md` shall identify when an ADR is required.
- **GH-REQ-150:** `CONTRIBUTING.md` shall identify when threat-model review is required.
- **GH-REQ-151:** `CONTRIBUTING.md` shall identify when specialist review is required.
- **GH-REQ-152:** `CONTRIBUTING.md` shall not promise contributor acceptance.
- **GH-REQ-153:** `CONTRIBUTING.md` shall not imply that contribution grants authority.
- **GH-REQ-154:** `CONTRIBUTING.md` changes shall receive documentation review.
- **GH-REQ-155:** Contribution history shall remain attributable.

## 8. Knowledge Pack and Normative Documents

- **GH-REQ-156:** Knowledge Pack documents 01–30 shall retain their assigned numbers.
- **GH-REQ-157:** Knowledge Pack document identifiers shall remain immutable.
- **GH-REQ-158:** Knowledge Pack filenames shall follow the canonical naming profile.
- **GH-REQ-159:** Legacy `_ENTERPRISE` variants shall not remain competing normative sources.
- **GH-REQ-160:** Legacy source variants shall be retained only in governed archives or consolidation evidence where needed.
- **GH-REQ-161:** Each normative Knowledge document shall declare metadata.
- **GH-REQ-162:** Each normative Knowledge document shall declare status.
- **GH-REQ-163:** Each normative Knowledge document shall declare version.
- **GH-REQ-164:** Each normative Knowledge document shall declare authority and owner.
- **GH-REQ-165:** Each normative Knowledge document shall declare dependencies.
- **GH-REQ-166:** Each normative Knowledge document shall declare limitations.
- **GH-REQ-167:** Each normative Knowledge document shall declare validation criteria.
- **GH-REQ-168:** Each normative Knowledge document shall use stable requirement identifiers.
- **GH-REQ-169:** Requirement identifiers shall not be reused.
- **GH-REQ-170:** Requirement deletions shall preserve historical disposition.
- **GH-REQ-171:** Normative documents shall not use filenames to imply approval.
- **GH-REQ-172:** Merge shall not change document status automatically.
- **GH-REQ-173:** Approval shall require documented authority and evidence.
- **GH-REQ-174:** Forward dependencies shall remain visible.
- **GH-REQ-175:** Cross-document contradictions shall block approval.
- **GH-REQ-176:** Knowledge Pack indexes shall identify document number, ID, version and status.
- **GH-REQ-177:** Knowledge Pack indexes shall identify authoritative paths.
- **GH-REQ-178:** Knowledge Pack indexes shall identify superseded variants.
- **GH-REQ-179:** Knowledge Pack indexes shall not list drafts as approved.
- **GH-REQ-180:** No normative document 31 or higher shall be created without constitutional authority.

## 9. Governance Documents

- **GH-REQ-181:** The canonical project definition shall remain the highest project-specific authority.
- **GH-REQ-182:** The Constitution shall remain immediately below the canonical definition.
- **GH-REQ-183:** Governance documents shall identify their authority level.
- **GH-REQ-184:** Governance documents shall identify owners.
- **GH-REQ-185:** Governance documents shall identify approval authority.
- **GH-REQ-186:** Governance documents shall preserve revision history.
- **GH-REQ-187:** Governance documents shall identify supersession.
- **GH-REQ-188:** Governance documents shall not be edited through generated output alone.
- **GH-REQ-189:** Governance changes shall use controlled pull requests.
- **GH-REQ-190:** Governance changes shall identify affected requirements.
- **GH-REQ-191:** Governance changes shall identify affected documents and implementation.
- **GH-REQ-192:** Governance changes shall identify security and privacy impacts.
- **GH-REQ-193:** Governance changes shall identify migration and rollback.
- **GH-REQ-194:** Governance changes shall receive required specialist review.
- **GH-REQ-195:** Governance decisions shall remain attributable.
- **GH-REQ-196:** Role concentration shall be disclosed.
- **GH-REQ-197:** Internal review shall not be represented as independent external assurance.
- **GH-REQ-198:** Governance exceptions shall be explicit and time-bounded.
- **GH-REQ-199:** Governance records shall remain accessible to authorized reviewers.
- **GH-REQ-200:** Governance history shall not be erased.
- **GH-REQ-201:** Conflicts shall resolve according to the authority hierarchy.
- **GH-REQ-202:** Governance indexes shall identify current and superseded records.
- **GH-REQ-203:** Governance paths shall be stable.
- **GH-REQ-204:** Governance files shall not contain secrets.
- **GH-REQ-205:** Governance release shall follow publication controls.

## 10. ADR Directory

- **GH-REQ-206:** ADRs shall reside under `adr/`.
- **GH-REQ-207:** ADR filenames shall follow `ADR-XXXX-short-title.md`.
- **GH-REQ-208:** ADR numbers shall not be reused.
- **GH-REQ-209:** ADRs shall identify status.
- **GH-REQ-210:** ADRs shall identify context.
- **GH-REQ-211:** ADRs shall identify decision.
- **GH-REQ-212:** ADRs shall identify alternatives.
- **GH-REQ-213:** ADRs shall identify consequences.
- **GH-REQ-214:** ADRs shall identify security and privacy impact.
- **GH-REQ-215:** ADRs shall identify implementation and migration impact.
- **GH-REQ-216:** ADRs shall identify validation and rollback.
- **GH-REQ-217:** ADRs shall identify related requirements and documents.
- **GH-REQ-218:** Accepted ADRs shall not be silently edited.
- **GH-REQ-219:** Accepted ADR corrections shall preserve history.
- **GH-REQ-220:** Superseded ADRs shall identify successors.
- **GH-REQ-221:** Rejected ADRs shall remain retrievable.
- **GH-REQ-222:** Proposed ADRs shall not be represented as accepted.
- **GH-REQ-223:** ADR indexes shall identify number, title and status.
- **GH-REQ-224:** ADR links shall use stable relative paths.
- **GH-REQ-225:** Architectural implementation shall trace to accepted ADRs.
- **GH-REQ-226:** Material architectural divergence shall require a new or superseding ADR.
- **GH-REQ-227:** ADR review shall remain attributable.
- **GH-REQ-228:** ADR directories shall not contain unrelated informal notes.
- **GH-REQ-229:** ADRs shall not contain secrets.
- **GH-REQ-230:** ADR governance shall conform to `OBDIA-ADR-001`.

## 11. Diagram Directory

- **GH-REQ-231:** Diagrams shall reside under `diagrams/` or approved subordinate documentation paths.
- **GH-REQ-232:** Diagram sources shall be distinguishable from rendered artifacts.
- **GH-REQ-233:** Authoritative source paths shall be documented.
- **GH-REQ-234:** Rendered artifacts shall trace to source.
- **GH-REQ-235:** Diagram filenames shall use governed kebab-case.
- **GH-REQ-236:** Diagram identifiers shall remain immutable.
- **GH-REQ-237:** Diagram indexes shall identify version, status and paths.
- **GH-REQ-238:** Diagram sources shall not contain secrets.
- **GH-REQ-239:** Public diagrams shall not contain real case data.
- **GH-REQ-240:** Public diagrams shall not expose sensitive internal endpoints unnecessarily.
- **GH-REQ-241:** Diagram metadata shall identify related documents.
- **GH-REQ-242:** Diagram metadata shall identify scope.
- **GH-REQ-243:** Diagram metadata shall identify version.
- **GH-REQ-244:** Diagram metadata shall identify status.
- **GH-REQ-245:** Source/render mismatches shall block release.
- **GH-REQ-246:** Stale rendered artifacts shall be detected.
- **GH-REQ-247:** Generated diagrams shall receive human review.
- **GH-REQ-248:** Diagram render workflows shall use least privilege.
- **GH-REQ-249:** External render services shall not receive sensitive content without authorization.
- **GH-REQ-250:** Diagram assets shall preserve license obligations.
- **GH-REQ-251:** Archived diagrams shall remain read-only.
- **GH-REQ-252:** Superseded diagrams shall identify successors.
- **GH-REQ-253:** Diagram links shall be validated.
- **GH-REQ-254:** Diagram tests shall conform to `OBDIA-TEST-001`.
- **GH-REQ-255:** Diagram structure shall conform to `OBDIA-DIAG-001`.

## 12. Source Code Structure

- **GH-REQ-256:** Source-code directories shall be introduced only when implementation exists.
- **GH-REQ-257:** Source-code paths shall be documented.
- **GH-REQ-258:** Source-code modules shall reflect architectural boundaries.
- **GH-REQ-259:** Security-critical modules shall have clear ownership.
- **GH-REQ-260:** Generated source shall be distinguishable from authored source.
- **GH-REQ-261:** Experimental source shall be distinguishable from supported implementation.
- **GH-REQ-262:** Demonstration source shall be distinguishable from reusable libraries.
- **GH-REQ-263:** Connector implementations shall be distinguishable from core policy logic.
- **GH-REQ-264:** Authorization code shall be distinguishable from identity and authentication code.
- **GH-REQ-265:** Evidence-handling code shall be distinguishable from analytical code.
- **GH-REQ-266:** Model and prompt assets shall be distinguishable from policy assets.
- **GH-REQ-267:** Configuration shall be distinguishable from secrets.
- **GH-REQ-268:** Build output shall be distinguishable from source.
- **GH-REQ-269:** Temporary output shall not be committed as authoritative source.
- **GH-REQ-270:** Source paths shall avoid real case identifiers.
- **GH-REQ-271:** Source examples shall use synthetic data.
- **GH-REQ-272:** Source modules shall not embed credentials.
- **GH-REQ-273:** Source structure shall support tests and review.
- **GH-REQ-274:** Source structure shall support dependency scanning.
- **GH-REQ-275:** Source structure shall support code ownership.
- **GH-REQ-276:** Source moves shall preserve history.
- **GH-REQ-277:** Source moves shall update imports and documentation.
- **GH-REQ-278:** Source structure changes with architectural impact shall require ADR review.
- **GH-REQ-279:** Unsupported source directories shall have rationale.
- **GH-REQ-280:** Source code shall conform to `OBDIA-CODE-001`.

## 13. Test Directory

- **GH-REQ-281:** The repository shall include `tests/`.
- **GH-REQ-282:** Test organization shall distinguish unit tests.
- **GH-REQ-283:** Test organization shall distinguish integration tests.
- **GH-REQ-284:** Test organization shall distinguish security tests.
- **GH-REQ-285:** Test organization shall distinguish fixtures.
- **GH-REQ-286:** Test organization shall distinguish retained evidence.
- **GH-REQ-287:** Language-standard test paths may supplement `tests/` when documented.
- **GH-REQ-288:** Test files shall use controlled naming.
- **GH-REQ-289:** Test fixtures shall be synthetic or explicitly authorized.
- **GH-REQ-290:** Test fixtures shall not contain production credentials.
- **GH-REQ-291:** Test fixtures shall not contain real case data for public demonstrations.
- **GH-REQ-292:** Malicious fixtures shall be defensive and isolated.
- **GH-REQ-293:** Test evidence shall not be confused with fixtures.
- **GH-REQ-294:** Generated test output shall not be committed unless retention is governed.
- **GH-REQ-295:** Authoritative validation packages shall reside in governed evidence paths.
- **GH-REQ-296:** Test paths shall support requirement traceability.
- **GH-REQ-297:** Test paths shall support defect traceability.
- **GH-REQ-298:** Test paths shall support CI discovery.
- **GH-REQ-299:** Test paths shall support local reproduction.
- **GH-REQ-300:** Test cleanup shall not delete retained evidence.
- **GH-REQ-301:** Skipped or quarantined tests shall remain visible.
- **GH-REQ-302:** Regression tests shall identify protected defects where practical.
- **GH-REQ-303:** Test directory indexes shall identify suites when scale requires it.
- **GH-REQ-304:** Test artifacts shall not expose secrets.
- **GH-REQ-305:** Test structure shall conform to `OBDIA-TEST-001`.

## 14. GitHub Metadata Directory

- **GH-REQ-306:** GitHub-specific configuration shall reside under `.github/` where supported.
- **GH-REQ-307:** Workflow definitions shall reside under `.github/workflows/`.
- **GH-REQ-308:** Issue templates shall reside under `.github/ISSUE_TEMPLATE/` where used.
- **GH-REQ-309:** Pull-request templates shall use a supported `.github/` or repository path.
- **GH-REQ-310:** CODEOWNERS shall use a supported repository path.
- **GH-REQ-311:** Funding configuration shall not be added without governance approval.
- **GH-REQ-312:** Discussion templates shall not imply official support unless support exists.
- **GH-REQ-313:** Community-health files shall remain synchronized.
- **GH-REQ-314:** GitHub-specific configuration shall be version-controlled.
- **GH-REQ-315:** GitHub-specific configuration shall receive review.
- **GH-REQ-316:** Configuration shall not contain secrets.
- **GH-REQ-317:** Configuration shall use least-privileged permissions.
- **GH-REQ-318:** Configuration changes shall identify platform assumptions.
- **GH-REQ-319:** Platform-specific behavior shall not replace project governance.
- **GH-REQ-320:** Repository settings not expressible in Git shall have evidence records.
- **GH-REQ-321:** Settings evidence shall identify capture time.
- **GH-REQ-322:** Settings evidence shall minimize sensitive administrative details.
- **GH-REQ-323:** Disabled features shall not be described as enabled.
- **GH-REQ-324:** Template changes shall preserve required governance fields.
- **GH-REQ-325:** GitHub configuration shall not override canonical authority.
- **GH-REQ-326:** Unrecognized GitHub configuration shall receive security review.
- **GH-REQ-327:** Deprecated platform configuration shall be migrated.
- **GH-REQ-328:** Automation-generated metadata shall remain attributable.
- **GH-REQ-329:** Repository metadata changes shall be auditable.
- **GH-REQ-330:** GitHub metadata shall be included in release review where material.

## 15. CODEOWNERS and Ownership

- **GH-REQ-331:** Material repository paths shall have identified owners.
- **GH-REQ-332:** CODEOWNERS shall reflect actual review responsibility where used.
- **GH-REQ-333:** Canonical and constitutional files shall have restricted ownership.
- **GH-REQ-334:** Knowledge Pack files shall have documentation and domain ownership.
- **GH-REQ-335:** Security-sensitive files shall include Security Reviewer ownership.
- **GH-REQ-336:** Implementation files shall include Implementation Reviewer ownership.
- **GH-REQ-337:** Privacy-sensitive files shall include Privacy and Governance Reviewer ownership where applicable.
- **GH-REQ-338:** Release files shall include Release Reviewer ownership.
- **GH-REQ-339:** Workflow files shall have security and implementation ownership.
- **GH-REQ-340:** Dependency files shall have implementation and security ownership.
- **GH-REQ-341:** Diagram files shall have Documentation Authority ownership.
- **GH-REQ-342:** Test evidence shall have implementation and applicable specialist ownership.
- **GH-REQ-343:** CODEOWNERS shall not substitute for required human approval.
- **GH-REQ-344:** CODEOWNERS shall not claim independent review.
- **GH-REQ-345:** Role concentration shall be disclosed outside CODEOWNERS where material.
- **GH-REQ-346:** Ownership entries shall use maintained accounts or teams.
- **GH-REQ-347:** Orphaned ownership entries shall be corrected.
- **GH-REQ-348:** Ownership changes shall preserve historical accountability.
- **GH-REQ-349:** Wildcard ownership shall not conceal critical-path specialization.
- **GH-REQ-350:** Ownership precedence shall be reviewed.
- **GH-REQ-351:** Sensitive path changes shall require applicable owners.
- **GH-REQ-352:** Emergency ownership bypass shall be documented.
- **GH-REQ-353:** Ownership configuration shall be tested where platform support permits.
- **GH-REQ-354:** Ownership records shall not expose unnecessary personal data.
- **GH-REQ-355:** Ownership review shall occur at risk-based intervals.

## 16. Branch Model

- **GH-REQ-356:** The repository shall identify one default branch.
- **GH-REQ-357:** The default branch shall be named consistently.
- **GH-REQ-358:** The default branch shall represent the current integrated repository state.
- **GH-REQ-359:** Feature and documentation work shall use short-lived branches where practical.
- **GH-REQ-360:** Branch names shall follow governed naming.
- **GH-REQ-361:** Branch names shall not contain secrets.
- **GH-REQ-362:** Branch names shall not contain real case identifiers.
- **GH-REQ-363:** Branch names shall identify purpose concisely.
- **GH-REQ-364:** Direct changes to protected authoritative content shall be restricted.
- **GH-REQ-365:** Branches shall not create independent approval status.
- **GH-REQ-366:** Branch existence shall not imply implementation.
- **GH-REQ-367:** Branch deletion shall not erase merged history.
- **GH-REQ-368:** Unmerged branch deletion shall follow retention and evidence needs.
- **GH-REQ-369:** Long-lived branches shall have documented purpose.
- **GH-REQ-370:** Stale branches shall be reviewed.
- **GH-REQ-371:** Abandoned branches shall not be represented as current work.
- **GH-REQ-372:** Protected branches shall use applicable review and check controls.
- **GH-REQ-373:** Emergency branch bypass shall be narrow and attributable.
- **GH-REQ-374:** Branch protection evidence shall be retained when material.
- **GH-REQ-375:** Fork-based contributions shall not receive privileged secrets.
- **GH-REQ-376:** Branches from untrusted contributors shall use restricted automation.
- **GH-REQ-377:** Branch synchronization shall not silently discard reviewed changes.
- **GH-REQ-378:** Force pushes to protected branches shall be prohibited or tightly governed.
- **GH-REQ-379:** Default-branch history rewrite shall be prohibited except under approved incident recovery.
- **GH-REQ-380:** Branch-model changes shall require governance review.

## 17. Commit Requirements

- **GH-REQ-381:** Commits shall be attributable to identifiable contributors or governed machine identities.
- **GH-REQ-382:** Commit messages shall summarize the change.
- **GH-REQ-383:** Commit messages shall not contain secrets.
- **GH-REQ-384:** Commit messages shall not contain unnecessary personal or case data.
- **GH-REQ-385:** Commits shall be scoped coherently.
- **GH-REQ-386:** Unrelated material changes shall not be bundled without rationale.
- **GH-REQ-387:** Generated changes shall identify their generator where material.
- **GH-REQ-388:** AI-assisted changes shall remain attributable to a human committer.
- **GH-REQ-389:** Commits shall not fabricate review or validation evidence.
- **GH-REQ-390:** Commits shall preserve file history where practical.
- **GH-REQ-391:** Commit signing may supplement but shall not replace review.
- **GH-REQ-392:** Unsigned commits shall not be represented as untrusted solely because signing is absent without a governing policy.
- **GH-REQ-393:** Commit-verification requirements shall be documented if enforced.
- **GH-REQ-394:** Failed local signing shall not justify bypassing repository governance.
- **GH-REQ-395:** Commit authorship shall not by itself establish authority.
- **GH-REQ-396:** Commit timestamps shall not be treated as sole evidence of sequence.
- **GH-REQ-397:** Commit amendments after review shall invalidate affected review evidence.
- **GH-REQ-398:** Fixup and squash behavior shall preserve PR-level review traceability.
- **GH-REQ-399:** Merge methods shall be governed.
- **GH-REQ-400:** Merge commits shall identify the integrated PR where the platform supports it.
- **GH-REQ-401:** Reverted commits shall remain visible.
- **GH-REQ-402:** Sensitive accidental commits shall trigger incident and history review.
- **GH-REQ-403:** Secret removal shall include rotation and exposure assessment.
- **GH-REQ-404:** Commit policy changes shall be reviewed.
- **GH-REQ-405:** Commit history shall remain retrievable.

## 18. Pull-Request Requirements

- **GH-REQ-406:** Material changes shall use pull requests.
- **GH-REQ-407:** Pull requests shall identify purpose.
- **GH-REQ-408:** Pull requests shall identify scope.
- **GH-REQ-409:** Pull requests shall identify changed authoritative artifacts.
- **GH-REQ-410:** Pull requests shall identify document or implementation status.
- **GH-REQ-411:** Pull requests shall identify related requirements.
- **GH-REQ-412:** Pull requests shall identify related ADRs.
- **GH-REQ-413:** Pull requests shall identify security impact.
- **GH-REQ-414:** Pull requests shall identify privacy impact.
- **GH-REQ-415:** Pull requests shall identify test evidence.
- **GH-REQ-416:** Pull requests shall identify validation limitations.
- **GH-REQ-417:** Pull requests shall identify forward dependencies.
- **GH-REQ-418:** Pull requests shall identify migration and rollback where material.
- **GH-REQ-419:** Pull requests shall identify exact integrity values for controlled consolidation artifacts where applicable.
- **GH-REQ-420:** Pull requests shall not claim approval merely because they are mergeable.
- **GH-REQ-421:** Draft pull requests shall not be represented as approved changes.
- **GH-REQ-422:** Pull-request templates shall preserve mandatory review fields.
- **GH-REQ-423:** Pull requests shall contain only expected files for controlled one-artifact consolidations.
- **GH-REQ-424:** Unexpected files shall block merge until explained.
- **GH-REQ-425:** Large pull requests shall be decomposed where practical.
- **GH-REQ-426:** Pull-request descriptions shall remain synchronized after material changes.
- **GH-REQ-427:** Automated summaries shall not replace human-readable rationale.
- **GH-REQ-428:** AI-generated descriptions shall receive human verification.
- **GH-REQ-429:** Sensitive information shall not be posted in public PRs.
- **GH-REQ-430:** Pull-request history shall remain available.

## 19. Pull-Request Review

- **GH-REQ-431:** Material pull requests shall receive human review.
- **GH-REQ-432:** Review shall identify the exact head commit.
- **GH-REQ-433:** Review shall identify affected requirements and authority.
- **GH-REQ-434:** Review shall verify scope isolation.
- **GH-REQ-435:** Review shall verify file paths.
- **GH-REQ-436:** Review shall verify integrity values where applicable.
- **GH-REQ-437:** Review shall verify status claims.
- **GH-REQ-438:** Review shall verify security impact.
- **GH-REQ-439:** Review shall verify privacy impact.
- **GH-REQ-440:** Review shall verify evidence and test impact.
- **GH-REQ-441:** Review shall verify release impact.
- **GH-REQ-442:** Security-sensitive changes shall receive Security Reviewer assessment.
- **GH-REQ-443:** Implementation changes shall receive Implementation Reviewer assessment.
- **GH-REQ-444:** Privacy-sensitive changes shall receive Privacy and Governance Reviewer assessment.
- **GH-REQ-445:** Documentation changes shall receive Documentation Authority assessment.
- **GH-REQ-446:** Release changes shall receive Release Reviewer assessment.
- **GH-REQ-447:** Review comments shall remain attributable.
- **GH-REQ-448:** Review findings shall have dispositions.
- **GH-REQ-449:** Blocking findings shall prevent merge.
- **GH-REQ-450:** Material changes after review shall invalidate affected approval.
- **GH-REQ-451:** Self-review shall not be represented as independent review.
- **GH-REQ-452:** Role concentration shall be disclosed.
- **GH-REQ-453:** Automated review shall remain advisory.
- **GH-REQ-454:** Review approval shall not establish lifecycle approval.
- **GH-REQ-455:** Review records shall remain retrievable.

## 20. Merge Governance

- **GH-REQ-456:** Merge shall require the reviewed head commit.
- **GH-REQ-457:** Head changes after review shall require renewed validation.
- **GH-REQ-458:** Required checks shall complete before merge unless a governed exception exists.
- **GH-REQ-459:** Failed required checks shall block merge.
- **GH-REQ-460:** Pending required checks shall block merge.
- **GH-REQ-461:** Cancelled required checks shall block merge unless rerun or dispositioned.
- **GH-REQ-462:** Skipped required checks shall block merge unless explicitly authorized.
- **GH-REQ-463:** Mergeability conflicts shall be resolved without discarding reviewed changes.
- **GH-REQ-464:** Merge method shall be documented.
- **GH-REQ-465:** Merge shall preserve attributable PR history.
- **GH-REQ-466:** Merge shall not automatically change document lifecycle status.
- **GH-REQ-467:** Merge shall not automatically establish validation.
- **GH-REQ-468:** Merge shall not automatically publish a release.
- **GH-REQ-469:** Merge shall not automatically accept residual risk.
- **GH-REQ-470:** Merge shall not automatically close unresolved security findings.
- **GH-REQ-471:** Administrative bypass shall be exceptional.
- **GH-REQ-472:** Administrative bypass shall identify authority and reason.
- **GH-REQ-473:** Administrative bypass shall trigger retrospective review.
- **GH-REQ-474:** Emergency merge shall preserve evidence.
- **GH-REQ-475:** Emergency merge shall not authorize canonical prohibitions.
- **GH-REQ-476:** Revert paths shall be available.
- **GH-REQ-477:** Reverts shall preserve original history.
- **GH-REQ-478:** Post-merge verification shall confirm expected files and integrity.
- **GH-REQ-479:** Branch cleanup shall occur after successful merge where appropriate.
- **GH-REQ-480:** Merge records shall remain auditable.

## 21. Branch Protection and Rulesets

- **GH-REQ-481:** The default branch shall use protection controls where repository capabilities permit.
- **GH-REQ-482:** Protection shall require pull requests for material changes where feasible.
- **GH-REQ-483:** Protection shall require applicable reviews.
- **GH-REQ-484:** Protection shall require applicable status checks.
- **GH-REQ-485:** Protection shall prevent force pushes where feasible.
- **GH-REQ-486:** Protection shall prevent branch deletion where feasible.
- **GH-REQ-487:** Protection shall restrict administrative bypass where feasible.
- **GH-REQ-488:** Protection shall require resolution of blocking conversations where applicable.
- **GH-REQ-489:** Protection shall apply to workflow files and governance files with heightened care.
- **GH-REQ-490:** Rulesets shall be documented.
- **GH-REQ-491:** Ruleset scope shall be documented.
- **GH-REQ-492:** Ruleset bypass actors shall be documented.
- **GH-REQ-493:** Ruleset changes shall receive review.
- **GH-REQ-494:** Ruleset changes shall have evidence.
- **GH-REQ-495:** Protection gaps shall be treated as residual risk.
- **GH-REQ-496:** Private-repository plan limitations shall be documented.
- **GH-REQ-497:** Platform limitations shall not be concealed.
- **GH-REQ-498:** Protection configuration shall not be inferred solely from documentation.
- **GH-REQ-499:** Protection verification shall use current repository evidence.
- **GH-REQ-500:** Protection tests shall avoid destructive actions.
- **GH-REQ-501:** Protection drift shall be detected where feasible.
- **GH-REQ-502:** Ruleset conflicts shall be reviewed.
- **GH-REQ-503:** Emergency protection changes shall be time-bounded.
- **GH-REQ-504:** Protection restoration shall be verified.
- **GH-REQ-505:** Protection evidence shall minimize administrative sensitive detail.

## 22. Issue Management

- **GH-REQ-506:** Issues shall use controlled categories or labels.
- **GH-REQ-507:** Security vulnerabilities shall not be reported through public issues when sensitive.
- **GH-REQ-508:** Issues shall not contain secrets.
- **GH-REQ-509:** Issues shall not contain unnecessary personal or case data.
- **GH-REQ-510:** Issues shall identify scope.
- **GH-REQ-511:** Issues shall identify affected requirements where applicable.
- **GH-REQ-512:** Issues shall identify owner.
- **GH-REQ-513:** Issues shall identify priority or severity where applicable.
- **GH-REQ-514:** Issues shall distinguish defect, enhancement, research, risk and governance work.
- **GH-REQ-515:** Issues shall distinguish implementation from documentation work.
- **GH-REQ-516:** Issues shall identify acceptance criteria where actionable.
- **GH-REQ-517:** Issues shall identify dependencies.
- **GH-REQ-518:** Issues shall identify blocking conditions.
- **GH-REQ-519:** Issues shall identify closure evidence.
- **GH-REQ-520:** Closing an issue shall not fabricate implementation or validation.
- **GH-REQ-521:** Duplicate issues shall preserve canonical references.
- **GH-REQ-522:** Transferred issues shall preserve traceability.
- **GH-REQ-523:** Sensitive issue content shall use restricted channels.
- **GH-REQ-524:** Automated issue creation shall remain attributable.
- **GH-REQ-525:** AI-generated issues shall receive human verification.
- **GH-REQ-526:** Stale issues shall not be closed solely by age when risk remains.
- **GH-REQ-527:** Risk and exception issues shall preserve owners and expiry.
- **GH-REQ-528:** Incident issues shall preserve evidence and access controls.
- **GH-REQ-529:** Issue templates shall guide safe reporting.
- **GH-REQ-530:** Issue history shall remain retrievable.

## 23. Labels, Milestones and Projects

- **GH-REQ-531:** Labels shall use controlled names.
- **GH-REQ-532:** Labels shall have documented meaning.
- **GH-REQ-533:** Lifecycle labels shall not override authoritative metadata.
- **GH-REQ-534:** Security-severity labels shall have documented criteria.
- **GH-REQ-535:** Privacy labels shall identify review needs.
- **GH-REQ-536:** Implementation-status labels shall be accurate.
- **GH-REQ-537:** Validation labels shall identify evidence scope.
- **GH-REQ-538:** Release labels shall not publish artifacts.
- **GH-REQ-539:** Labels shall not contain personal or case identifiers.
- **GH-REQ-540:** Label colors shall not be the sole carrier of meaning.
- **GH-REQ-541:** Label changes shall preserve historical interpretation where material.
- **GH-REQ-542:** Milestones shall identify scope and exit criteria.
- **GH-REQ-543:** Milestones shall not imply approval.
- **GH-REQ-544:** Milestone completion shall not imply validation.
- **GH-REQ-545:** Project boards shall not become authoritative governance sources.
- **GH-REQ-546:** Project status shall trace to issues and PRs.
- **GH-REQ-547:** Automated project updates shall not alter normative status.
- **GH-REQ-548:** Roadmaps shall distinguish planned and committed work.
- **GH-REQ-549:** Future research shall be labeled as future research.
- **GH-REQ-550:** Archived projects shall preserve relevant links.
- **GH-REQ-551:** Duplicate status systems shall be reconciled.
- **GH-REQ-552:** Unknown status shall remain explicit.
- **GH-REQ-553:** Reporting dashboards shall not conceal blocking findings.
- **GH-REQ-554:** Labels and milestones shall be reviewed periodically.
- **GH-REQ-555:** Administrative metadata shall remain attributable.

## 24. Workflow Security

- **GH-REQ-556:** Workflows shall reside under `.github/workflows/`.
- **GH-REQ-557:** Workflow files shall be version-controlled.
- **GH-REQ-558:** Workflow changes shall receive human review.
- **GH-REQ-559:** Workflow permissions shall default to least privilege.
- **GH-REQ-560:** Workflow permissions shall be explicit where practical.
- **GH-REQ-561:** Write permissions shall be granted only when required.
- **GH-REQ-562:** Workflow secrets shall not be available to untrusted changes.
- **GH-REQ-563:** Untrusted pull-request code shall not execute with privileged secrets.
- **GH-REQ-564:** Workflows shall avoid unsafe evaluation of untrusted input.
- **GH-REQ-565:** Workflow shell commands shall quote and validate untrusted values.
- **GH-REQ-566:** Workflow actions shall use pinned versions or immutable references where appropriate.
- **GH-REQ-567:** Third-party actions shall receive supply-chain review.
- **GH-REQ-568:** Workflow dependencies shall be inventoried.
- **GH-REQ-569:** Workflow logs shall not expose secrets.
- **GH-REQ-570:** Workflow artifacts shall have retention rules.
- **GH-REQ-571:** Workflow artifacts shall not contain real case data.
- **GH-REQ-572:** Workflow caches shall not bypass integrity or trust boundaries.
- **GH-REQ-573:** Workflow tokens shall have bounded scope.
- **GH-REQ-574:** Workflow environments shall have controlled approvals where supported and required.
- **GH-REQ-575:** Workflow concurrency shall prevent unsafe duplicate releases.
- **GH-REQ-576:** Workflow timeouts shall be bounded.
- **GH-REQ-577:** Workflow failures shall remain visible.
- **GH-REQ-578:** Workflow reruns shall preserve prior failures.
- **GH-REQ-579:** Workflow disablement shall be auditable.
- **GH-REQ-580:** Compromised workflow infrastructure shall invalidate affected evidence.

## 25. Continuous Integration

- **GH-REQ-581:** Continuous integration shall validate exact commits.
- **GH-REQ-582:** CI shall identify workflow and run identifiers.
- **GH-REQ-583:** CI shall validate documentation syntax where applicable.
- **GH-REQ-584:** CI shall validate required document metadata.
- **GH-REQ-585:** CI shall validate requirement identifiers.
- **GH-REQ-586:** CI shall validate filenames and paths.
- **GH-REQ-587:** CI shall validate cross-references.
- **GH-REQ-588:** CI shall validate diagram syntax and rendering where configured.
- **GH-REQ-589:** CI shall run applicable unit tests.
- **GH-REQ-590:** CI shall run applicable integration tests.
- **GH-REQ-591:** CI shall run applicable security tests.
- **GH-REQ-592:** CI shall run applicable negative tests.
- **GH-REQ-593:** CI shall run applicable regression tests.
- **GH-REQ-594:** CI shall perform secret scanning.
- **GH-REQ-595:** CI shall perform dependency checks where implementation exists.
- **GH-REQ-596:** CI shall perform static analysis where implementation exists.
- **GH-REQ-597:** CI shall preserve raw results.
- **GH-REQ-598:** CI shall distinguish pass, fail, blocked, skipped, aborted and invalid outcomes.
- **GH-REQ-599:** CI tool failure shall not count as pass.
- **GH-REQ-600:** CI shall not automatically establish lifecycle `Validated`.
- **GH-REQ-601:** CI shall not automatically accept risk.
- **GH-REQ-602:** CI shall not automatically approve release.
- **GH-REQ-603:** CI summaries shall not replace raw evidence.
- **GH-REQ-604:** CI environments shall use synthetic or controlled data.
- **GH-REQ-605:** CI shall not interact with uncontrolled criminal infrastructure.

## 26. Continuous Delivery and Release Automation

- **GH-REQ-606:** Release automation shall remain disabled until governed release requirements exist.
- **GH-REQ-607:** Release automation shall operate on an exact reviewed commit.
- **GH-REQ-608:** Release automation shall verify version and tag consistency.
- **GH-REQ-609:** Release automation shall verify artifact manifests.
- **GH-REQ-610:** Release automation shall verify integrity values.
- **GH-REQ-611:** Release automation shall verify test and validation evidence.
- **GH-REQ-612:** Release automation shall verify secret scanning.
- **GH-REQ-613:** Release automation shall verify license and attribution.
- **GH-REQ-614:** Release automation shall verify public-content safety.
- **GH-REQ-615:** Release automation shall distinguish draft, prerelease and published release.
- **GH-REQ-616:** Release automation shall not publish from untrusted pull requests.
- **GH-REQ-617:** Release credentials shall use least privilege.
- **GH-REQ-618:** Release credentials shall be separately protected.
- **GH-REQ-619:** Release environments shall use required human approval where supported and mandated.
- **GH-REQ-620:** Release automation shall support cancellation.
- **GH-REQ-621:** Release automation shall support withdrawal or correction procedures.
- **GH-REQ-622:** Release automation shall not rewrite existing release evidence silently.
- **GH-REQ-623:** Release artifacts shall trace to source.
- **GH-REQ-624:** Release artifacts shall not contain real case data.
- **GH-REQ-625:** Release artifacts shall not contain credentials.
- **GH-REQ-626:** Release artifacts shall not claim production readiness without evidence.
- **GH-REQ-627:** Automated changelogs shall receive human verification.
- **GH-REQ-628:** Release failures shall block publication.
- **GH-REQ-629:** Release reruns shall preserve prior failure evidence.
- **GH-REQ-630:** Release automation shall conform to `OBDIA-REL-001`.

## 27. Dependency Files and Automation

- **GH-REQ-631:** Dependency manifests shall reside at standard governed paths.
- **GH-REQ-632:** Resolved dependency files or lockfiles shall be retained where applicable.
- **GH-REQ-633:** Dependency files shall not contain credentials.
- **GH-REQ-634:** Dependency files shall use approved registries.
- **GH-REQ-635:** Dependency updates shall use pull requests.
- **GH-REQ-636:** Automated dependency updates shall receive human review.
- **GH-REQ-637:** Automated dependency updates shall run applicable tests.
- **GH-REQ-638:** Dependency update grouping shall not conceal critical changes.
- **GH-REQ-639:** Major updates shall receive impact analysis.
- **GH-REQ-640:** Security updates shall preserve urgency and review.
- **GH-REQ-641:** Dependency provenance shall be recorded.
- **GH-REQ-642:** Dependency licenses shall be reviewed.
- **GH-REQ-643:** Unsupported dependencies shall have dispositions.
- **GH-REQ-644:** Critical unresolved vulnerabilities shall block release.
- **GH-REQ-645:** False-positive dispositions shall have evidence.
- **GH-REQ-646:** Dependency exceptions shall have owners and expiry.
- **GH-REQ-647:** Dependency automation credentials shall use least privilege.
- **GH-REQ-648:** Dependency automation shall not merge autonomously unless explicitly governed.
- **GH-REQ-649:** Dependency bots shall not approve their own changes.
- **GH-REQ-650:** Dependency change history shall remain attributable.
- **GH-REQ-651:** Vendored dependencies shall identify source and version.
- **GH-REQ-652:** Forked dependencies shall identify divergence.
- **GH-REQ-653:** Dependency-confusion controls shall be applied.
- **GH-REQ-654:** Dependency inventories shall remain current.
- **GH-REQ-655:** Dependency governance shall conform to `OBDIA-CODE-001`.

## 28. Secret Management and Scanning

- **GH-REQ-656:** The repository shall not contain plaintext secrets.
- **GH-REQ-657:** The repository shall not contain private keys.
- **GH-REQ-658:** The repository shall not contain reusable human credentials.
- **GH-REQ-659:** The repository shall not contain production tokens.
- **GH-REQ-660:** The repository shall not contain secret-bearing configuration examples.
- **GH-REQ-661:** Secret placeholders shall be unmistakably synthetic.
- **GH-REQ-662:** Secret scanning shall cover current content.
- **GH-REQ-663:** Secret scanning shall cover repository history where feasible.
- **GH-REQ-664:** Secret scanning shall cover workflow files.
- **GH-REQ-665:** Secret scanning shall cover fixtures.
- **GH-REQ-666:** Secret scanning shall cover generated artifacts.
- **GH-REQ-667:** Secret scanning findings shall have dispositions.
- **GH-REQ-668:** Detected secrets shall be rotated or revoked.
- **GH-REQ-669:** Deleting a secret from the latest commit shall not be treated as sufficient remediation.
- **GH-REQ-670:** Secret incidents shall assess forks, clones, logs and artifacts.
- **GH-REQ-671:** Secret remediation shall preserve incident evidence.
- **GH-REQ-672:** Secret-management references shall use identifiers rather than values.
- **GH-REQ-673:** Workflow secrets shall have bounded scope.
- **GH-REQ-674:** Environment secrets shall remain separated.
- **GH-REQ-675:** Development and test secrets shall remain distinct from production secrets.
- **GH-REQ-676:** Public demonstrations shall use non-secret values.
- **GH-REQ-677:** AI prompts shall not receive repository secrets.
- **GH-REQ-678:** External analysis services shall not receive secret-bearing files without authorization.
- **GH-REQ-679:** Secret exceptions shall not permit hard-coded production credentials.
- **GH-REQ-680:** Secret policy shall be tested.

## 29. Data and Evidence in the Repository

- **GH-REQ-681:** Repository data shall be classified.
- **GH-REQ-682:** Real case data shall not be committed to public repository content.
- **GH-REQ-683:** Public demonstrations shall use synthetic data.
- **GH-REQ-684:** Fixtures shall avoid real victims, suspects and case subjects.
- **GH-REQ-685:** Personal data shall be minimized.
- **GH-REQ-686:** Sensitive data shall use restricted governed storage outside ordinary repository paths when required.
- **GH-REQ-687:** Evidence objects shall not be committed casually as documentation examples.
- **GH-REQ-688:** Evidence fixtures shall be unmistakably synthetic.
- **GH-REQ-689:** Evidence metadata shall not expose real case identifiers.
- **GH-REQ-690:** Evidence integrity values shall identify synthetic scope where applicable.
- **GH-REQ-691:** Derived analytical outputs shall not be represented as source evidence.
- **GH-REQ-692:** Repository history shall be considered when removing sensitive data.
- **GH-REQ-693:** Large evidence artifacts shall have explicit retention and storage decisions.
- **GH-REQ-694:** Encrypted sensitive artifacts shall not be committed solely because encrypted.
- **GH-REQ-695:** Encryption keys shall not reside in the same repository.
- **GH-REQ-696:** Evidence exports shall follow `OBDIA-EVID-001`.
- **GH-REQ-697:** Test evidence shall remain distinct from investigation evidence.
- **GH-REQ-698:** CI artifacts shall not become authoritative evidence automatically.
- **GH-REQ-699:** Repository backups shall preserve classification.
- **GH-REQ-700:** Repository mirrors shall preserve access restrictions.
- **GH-REQ-701:** Data incidents shall trigger history and release review.
- **GH-REQ-702:** Data deletion shall preserve required audit evidence.
- **GH-REQ-703:** Repository archival shall not publish restricted data.
- **GH-REQ-704:** Data handling shall comply with publication policy.
- **GH-REQ-705:** Data limitations shall remain visible.

## 30. Generated Files and Build Artifacts

- **GH-REQ-706:** Generated files shall identify their authoritative source.
- **GH-REQ-707:** Generated files shall identify generator and version where material.
- **GH-REQ-708:** Generated files shall not be manually edited unless explicitly allowed.
- **GH-REQ-709:** Generated files shall be reproducible where practical.
- **GH-REQ-710:** Generated files shall not contain secrets.
- **GH-REQ-711:** Generated files shall not contain unauthorized personal data.
- **GH-REQ-712:** Generated files shall not create normative authority.
- **GH-REQ-713:** Rendered diagrams shall remain generated from governed source.
- **GH-REQ-714:** Compiled artifacts shall not be committed unless repository policy requires it.
- **GH-REQ-715:** Build artifacts shall reside in release or artifact storage when appropriate.
- **GH-REQ-716:** Temporary output shall be ignored through governed ignore rules.
- **GH-REQ-717:** Ignore rules shall not conceal authoritative evidence.
- **GH-REQ-718:** Generated documentation shall receive review.
- **GH-REQ-719:** AI-generated documents shall receive human review.
- **GH-REQ-720:** Generated indexes shall trace to authoritative metadata.
- **GH-REQ-721:** Generated integrity manifests shall identify exact inputs.
- **GH-REQ-722:** Stale generated output shall be detected.
- **GH-REQ-723:** Source/generated mismatches shall block release.
- **GH-REQ-724:** Generated output changes shall be included in review when committed.
- **GH-REQ-725:** Build artifacts shall have integrity values where released.
- **GH-REQ-726:** Generated file licensing shall be considered.
- **GH-REQ-727:** Generator compromise shall invalidate affected output.
- **GH-REQ-728:** Generated output retention shall be governed.
- **GH-REQ-729:** Generated files shall not conceal implementation changes.
- **GH-REQ-730:** Generated-artifact rules shall be documented.

## 31. Ignore and Attribute Files

- **GH-REQ-731:** `.gitignore` shall exclude local secrets and transient artifacts.
- **GH-REQ-732:** `.gitignore` shall exclude environment-specific sensitive files.
- **GH-REQ-733:** `.gitignore` shall exclude build output unless intentionally governed.
- **GH-REQ-734:** `.gitignore` shall exclude caches.
- **GH-REQ-735:** `.gitignore` shall exclude local virtual environments where applicable.
- **GH-REQ-736:** `.gitignore` shall not be relied on after sensitive data is committed.
- **GH-REQ-737:** Ignore rules shall not exclude required normative artifacts.
- **GH-REQ-738:** Ignore rules shall not conceal failed test evidence required for retention.
- **GH-REQ-739:** Ignore rules shall be reviewed when new tools are introduced.
- **GH-REQ-740:** Global ignore behavior shall not be assumed by repository contributors.
- **GH-REQ-741:** `.gitattributes` shall define text normalization where needed.
- **GH-REQ-742:** `.gitattributes` shall identify binary files where needed.
- **GH-REQ-743:** Line-ending policy shall preserve deterministic content hashes where material.
- **GH-REQ-744:** Diff drivers may be configured for reviewability.
- **GH-REQ-745:** Generated files may be marked where platform support exists.
- **GH-REQ-746:** Large-file handling shall be governed.
- **GH-REQ-747:** Git LFS or equivalent use shall identify availability and retention implications.
- **GH-REQ-748:** Pointer files shall not be mistaken for complete artifacts.
- **GH-REQ-749:** Attribute changes shall receive review.
- **GH-REQ-750:** Merge strategies for generated or structured files shall be safe.
- **GH-REQ-751:** Archive export behavior shall be reviewed.
- **GH-REQ-752:** Submodule behavior shall be explicit where used.
- **GH-REQ-753:** Submodules shall not introduce uncontrolled dependencies.
- **GH-REQ-754:** Ignore and attribute files shall not contain secrets.
- **GH-REQ-755:** Ignore and attribute behavior shall be tested where material.

## 32. Repository Settings

- **GH-REQ-756:** Repository visibility shall be explicitly governed.
- **GH-REQ-757:** Default branch shall be explicitly configured.
- **GH-REQ-758:** Merge methods shall be explicitly governed.
- **GH-REQ-759:** Automatic branch deletion shall be governed.
- **GH-REQ-760:** Issue availability shall be governed.
- **GH-REQ-761:** Discussion availability shall be governed.
- **GH-REQ-762:** Wiki availability shall be governed.
- **GH-REQ-763:** Project-board availability shall be governed.
- **GH-REQ-764:** Forking availability shall be governed where controls permit.
- **GH-REQ-765:** Actions permissions shall be governed.
- **GH-REQ-766:** Workflow token permissions shall be governed.
- **GH-REQ-767:** Artifact retention shall be governed.
- **GH-REQ-768:** Environment protections shall be governed.
- **GH-REQ-769:** Security-feature availability shall be assessed.
- **GH-REQ-770:** Dependabot or equivalent availability shall be assessed where implementation exists.
- **GH-REQ-771:** Secret-scanning availability shall be assessed.
- **GH-REQ-772:** Code-scanning availability shall be assessed.
- **GH-REQ-773:** Repository settings shall not be inferred from desired documentation.
- **GH-REQ-774:** Settings evidence shall identify the actual observed configuration.
- **GH-REQ-775:** Settings drift shall be reviewed.
- **GH-REQ-776:** Settings changes shall be attributable.
- **GH-REQ-777:** Administrative access shall use least privilege.
- **GH-REQ-778:** Inactive administrators shall be removed.
- **GH-REQ-779:** Administrative changes shall be logged where platform support permits.
- **GH-REQ-780:** Repository-settings limitations shall remain visible.

## 33. Access Control

- **GH-REQ-781:** Repository access shall use individual accounts or governed machine identities.
- **GH-REQ-782:** Shared human accounts shall be prohibited.
- **GH-REQ-783:** Human credentials shall not be transferred to agents.
- **GH-REQ-784:** Access shall use least privilege.
- **GH-REQ-785:** Administrative access shall be limited.
- **GH-REQ-786:** Write access shall be limited to contributors who require it.
- **GH-REQ-787:** Triage and read roles shall be used where sufficient.
- **GH-REQ-788:** Machine identities shall have narrow scopes.
- **GH-REQ-789:** Automation tokens shall have bounded audience and validity.
- **GH-REQ-790:** Access shall be revocable.
- **GH-REQ-791:** Access reviews shall occur at risk-based intervals.
- **GH-REQ-792:** Officer or contributor departure shall trigger access review.
- **GH-REQ-793:** Compromised accounts shall trigger containment.
- **GH-REQ-794:** Multi-factor authentication shall be required where project or organization policy mandates it.
- **GH-REQ-795:** Authentication shall not be treated as authorization to approve governance.
- **GH-REQ-796:** Repository role shall not create legal authority.
- **GH-REQ-797:** External collaborator access shall be time-bounded where feasible.
- **GH-REQ-798:** Fork access shall not expose private secrets.
- **GH-REQ-799:** Access changes shall remain auditable.
- **GH-REQ-800:** Emergency access shall be narrow and time-bounded.
- **GH-REQ-801:** Emergency access shall trigger retrospective review.
- **GH-REQ-802:** Orphaned machine identities shall be revoked.
- **GH-REQ-803:** Inactive tokens shall be revoked.
- **GH-REQ-804:** Access exceptions shall have owners and expiry.
- **GH-REQ-805:** Access limitations shall remain visible.

## 34. Tags and Versions

- **GH-REQ-806:** Tags shall follow governed versioning.
- **GH-REQ-807:** Release tags shall identify exact commits.
- **GH-REQ-808:** Tags shall not be reused.
- **GH-REQ-809:** Published tags shall not be moved silently.
- **GH-REQ-810:** Draft-document versions shall not imply release versions automatically.
- **GH-REQ-811:** Repository version and document versions shall remain distinguishable.
- **GH-REQ-812:** Artifact versions shall remain distinguishable from schema versions.
- **GH-REQ-813:** Tag names shall use controlled format.
- **GH-REQ-814:** Prerelease tags shall be distinguishable.
- **GH-REQ-815:** Experimental tags shall not imply support.
- **GH-REQ-816:** Signed tags may supplement but shall not replace release review.
- **GH-REQ-817:** Tag creation shall be attributable.
- **GH-REQ-818:** Tag deletion shall require governance review.
- **GH-REQ-819:** Incorrect tags shall be corrected through documented procedure.
- **GH-REQ-820:** Tag corrections shall preserve prior evidence.
- **GH-REQ-821:** Version metadata shall match release manifests.
- **GH-REQ-822:** Version changes shall identify compatibility impact.
- **GH-REQ-823:** Version changes shall identify migration.
- **GH-REQ-824:** Version changes shall identify rollback.
- **GH-REQ-825:** Unknown version state shall block release.
- **GH-REQ-826:** Version indexes shall be maintained where scale requires it.
- **GH-REQ-827:** Document 25 shall govern final version semantics.
- **GH-REQ-828:** Version automation shall not publish without approval.
- **GH-REQ-829:** Version history shall remain retrievable.
- **GH-REQ-830:** Versioning limitations shall remain visible.

## 35. Releases

- **GH-REQ-831:** Releases shall follow `OBDIA-REL-001`.
- **GH-REQ-832:** Every release shall have a unique version.
- **GH-REQ-833:** Every release shall identify exact commit.
- **GH-REQ-834:** Every release shall identify included artifacts.
- **GH-REQ-835:** Every release shall include an integrity manifest where required.
- **GH-REQ-836:** Every release shall identify approval evidence.
- **GH-REQ-837:** Every release shall identify test and validation evidence.
- **GH-REQ-838:** Every release shall identify known limitations.
- **GH-REQ-839:** Every release shall identify security and privacy review.
- **GH-REQ-840:** Every release shall identify license and attribution.
- **GH-REQ-841:** Every release shall identify changes and migration.
- **GH-REQ-842:** Every release shall identify rollback or withdrawal procedure.
- **GH-REQ-843:** Draft releases shall not be represented as published.
- **GH-REQ-844:** Prereleases shall be labeled.
- **GH-REQ-845:** Public releases shall use approved safe examples.
- **GH-REQ-846:** Public releases shall not contain secrets.
- **GH-REQ-847:** Public releases shall not contain real case data.
- **GH-REQ-848:** Public releases shall not claim production readiness without evidence.
- **GH-REQ-849:** Public releases shall not claim certification or compliance without evidence.
- **GH-REQ-850:** Release assets shall trace to source.
- **GH-REQ-851:** Release assets shall have integrity values.
- **GH-REQ-852:** Release corrections shall preserve history.
- **GH-REQ-853:** Release withdrawal shall preserve records.
- **GH-REQ-854:** Release automation shall not replace Release Reviewer approval.
- **GH-REQ-855:** Release history shall remain retrievable.

## 36. Archival and Retention

- **GH-REQ-856:** Repository archival shall be an explicit governance decision.
- **GH-REQ-857:** Archived repositories shall be read-only where platform support permits.
- **GH-REQ-858:** Archival shall identify authority.
- **GH-REQ-859:** Archival shall identify reason.
- **GH-REQ-860:** Archival shall identify effective time.
- **GH-REQ-861:** Archival shall identify retained artifacts.
- **GH-REQ-862:** Archival shall identify external dependencies.
- **GH-REQ-863:** Archival shall identify security contacts or successor channels.
- **GH-REQ-864:** Archival shall identify successor repository where applicable.
- **GH-REQ-865:** Archival shall preserve canonical and governance documents.
- **GH-REQ-866:** Archival shall preserve ADRs.
- **GH-REQ-867:** Archival shall preserve release records.
- **GH-REQ-868:** Archival shall preserve test and validation evidence according to retention.
- **GH-REQ-869:** Archival shall preserve security incident records according to classification.
- **GH-REQ-870:** Archival shall not publish previously private data automatically.
- **GH-REQ-871:** Archival shall not reactivate revoked credentials.
- **GH-REQ-872:** Archival shall revoke unnecessary automation tokens.
- **GH-REQ-873:** Archival shall disable unnecessary workflows.
- **GH-REQ-874:** Archival shall preserve integrity manifests.
- **GH-REQ-875:** Archived links shall remain stable where feasible.
- **GH-REQ-876:** Archive restoration shall not restore approval or operational state silently.
- **GH-REQ-877:** Archive migration shall preserve history.
- **GH-REQ-878:** Archive corruption shall trigger incident handling.
- **GH-REQ-879:** Retention conflicts shall resolve to the more protective rule until reviewed.
- **GH-REQ-880:** Archival limitations shall remain visible.

## 37. Forks, Mirrors and Copies

- **GH-REQ-881:** Fork policy shall be governed.
- **GH-REQ-882:** Private forks shall preserve access restrictions.
- **GH-REQ-883:** Public forks shall not receive private secrets.
- **GH-REQ-884:** Forks shall not be treated as authoritative project repositories.
- **GH-REQ-885:** Official mirrors shall be explicitly identified.
- **GH-REQ-886:** Unofficial mirrors shall not be represented as official.
- **GH-REQ-887:** Mirrors shall preserve license and attribution.
- **GH-REQ-888:** Mirrors shall preserve release integrity where applicable.
- **GH-REQ-889:** Mirrors shall identify synchronization status.
- **GH-REQ-890:** Stale mirrors shall not be represented as current.
- **GH-REQ-891:** Mirror credentials shall use least privilege.
- **GH-REQ-892:** Mirror automation shall not expose secrets.
- **GH-REQ-893:** Repository copies shall preserve provenance.
- **GH-REQ-894:** Exported archives shall identify source commit.
- **GH-REQ-895:** Exported archives shall not include hidden secrets.
- **GH-REQ-896:** Template repositories shall not include operational credentials.
- **GH-REQ-897:** Template content shall use synthetic values.
- **GH-REQ-898:** Downstream reuse shall not imply institutional endorsement.
- **GH-REQ-899:** Fork pull requests shall use untrusted-workflow protections.
- **GH-REQ-900:** Fork deletion shall not erase upstream records.
- **GH-REQ-901:** Security incidents affecting forks shall be assessed.
- **GH-REQ-902:** License changes shall consider forks and releases.
- **GH-REQ-903:** Fork-specific modifications shall remain distinguishable.
- **GH-REQ-904:** Mirror failure shall not alter the authoritative repository.
- **GH-REQ-905:** Fork and mirror limitations shall remain visible.

## 38. Submodules and External References

- **GH-REQ-906:** Submodules shall be avoided unless they provide clear governed value.
- **GH-REQ-907:** Submodules shall use explicit immutable revisions.
- **GH-REQ-908:** Submodule sources shall be reviewed.
- **GH-REQ-909:** Submodule licenses shall be reviewed.
- **GH-REQ-910:** Submodule availability risk shall be documented.
- **GH-REQ-911:** Submodule updates shall use pull requests.
- **GH-REQ-912:** Submodule updates shall trigger applicable tests.
- **GH-REQ-913:** Submodules shall not contain secrets.
- **GH-REQ-914:** Submodules shall not introduce unauthorized tooling.
- **GH-REQ-915:** Submodules shall not point to uncontrolled criminal infrastructure.
- **GH-REQ-916:** External documentation links shall use authoritative sources where possible.
- **GH-REQ-917:** Critical external references shall be archived or summarized where legally permitted and necessary for reproducibility.
- **GH-REQ-918:** External references shall identify access date where temporal relevance matters.
- **GH-REQ-919:** Broken critical references shall be corrected.
- **GH-REQ-920:** External scripts shall not be executed directly without review.
- **GH-REQ-921:** Remote includes shall not create hidden build behavior.
- **GH-REQ-922:** Vendored external content shall preserve provenance.
- **GH-REQ-923:** External schemas shall be versioned.
- **GH-REQ-924:** External assets shall preserve license obligations.
- **GH-REQ-925:** External references shall not become normative by hyperlink alone.
- **GH-REQ-926:** Dependency on proprietary access shall be disclosed.
- **GH-REQ-927:** Access-controlled references shall identify review limitations.
- **GH-REQ-928:** Removed external content shall not be fabricated.
- **GH-REQ-929:** External reference checks shall be included in documentation validation.
- **GH-REQ-930:** External-reference limitations shall remain visible.

## 39. Documentation Links and Indexes

- **GH-REQ-931:** The repository shall maintain a Knowledge Pack index.
- **GH-REQ-932:** The repository shall maintain an ADR index.
- **GH-REQ-933:** The repository shall maintain a diagram index where diagrams exist.
- **GH-REQ-934:** The repository shall maintain an implementation index where implementation exists.
- **GH-REQ-935:** The repository shall maintain a testing and validation index where evidence exists.
- **GH-REQ-936:** Indexes shall identify immutable IDs.
- **GH-REQ-937:** Indexes shall identify titles.
- **GH-REQ-938:** Indexes shall identify versions.
- **GH-REQ-939:** Indexes shall identify statuses.
- **GH-REQ-940:** Indexes shall identify authoritative paths.
- **GH-REQ-941:** Indexes shall identify owners.
- **GH-REQ-942:** Indexes shall identify supersession where applicable.
- **GH-REQ-943:** Indexes shall not infer status from path.
- **GH-REQ-944:** Indexes shall not list absent artifacts as complete.
- **GH-REQ-945:** Broken internal links shall be detected.
- **GH-REQ-946:** Renamed files shall update inbound links.
- **GH-REQ-947:** Relative links shall be preferred for repository portability.
- **GH-REQ-948:** External links shall be reviewed.
- **GH-REQ-949:** Link text shall be descriptive.
- **GH-REQ-950:** Links shall not expose restricted paths in public artifacts.
- **GH-REQ-951:** Generated indexes shall trace to authoritative metadata.
- **GH-REQ-952:** Generated indexes shall receive review.
- **GH-REQ-953:** Index corrections shall preserve history.
- **GH-REQ-954:** Stale indexes shall block release where they materially mislead.
- **GH-REQ-955:** Documentation navigation shall support recruiter and reviewer readability.

## 40. Repository Security Incidents

- **GH-REQ-956:** Repository security incidents shall have immutable identifiers.
- **GH-REQ-957:** Incidents shall identify affected repositories, branches, commits and artifacts.
- **GH-REQ-958:** Incidents shall identify affected credentials.
- **GH-REQ-959:** Incidents shall identify affected workflows.
- **GH-REQ-960:** Incidents shall identify affected releases.
- **GH-REQ-961:** Incidents shall identify affected documents and evidence.
- **GH-REQ-962:** Incidents shall preserve logs and audit evidence.
- **GH-REQ-963:** Secret exposure shall trigger rotation or revocation.
- **GH-REQ-964:** Malicious commit discovery shall trigger containment.
- **GH-REQ-965:** Workflow compromise shall trigger token revocation.
- **GH-REQ-966:** Release compromise shall trigger withdrawal or correction.
- **GH-REQ-967:** History rewrite shall not occur without approved incident procedure.
- **GH-REQ-968:** Public disclosure shall follow security and release governance.
- **GH-REQ-969:** Incident response shall not destroy evidence.
- **GH-REQ-970:** Incident access shall be restricted.
- **GH-REQ-971:** Incident recovery shall restore a known governed state.
- **GH-REQ-972:** Recovered branches shall be verified.
- **GH-REQ-973:** Recovered workflows shall be reviewed.
- **GH-REQ-974:** Recovered artifacts shall have new integrity evidence.
- **GH-REQ-975:** Post-incident review shall update threats and controls.
- **GH-REQ-976:** Regression tests shall be added where code changes correct incidents.
- **GH-REQ-977:** Repeated incident classes shall trigger architecture review.
- **GH-REQ-978:** Incident closure shall require containment and follow-up ownership.
- **GH-REQ-979:** Incident records shall remain retrievable.
- **GH-REQ-980:** Incident limitations shall remain visible.

## 41. Repository Validation

- **GH-REQ-981:** Repository validation shall identify the exact commit.
- **GH-REQ-982:** Repository validation shall verify mandatory root files.
- **GH-REQ-983:** Repository validation shall verify mandatory governed directories.
- **GH-REQ-984:** Repository validation shall verify Knowledge Pack numbering.
- **GH-REQ-985:** Repository validation shall verify immutable document IDs.
- **GH-REQ-986:** Repository validation shall verify requirement-ID uniqueness.
- **GH-REQ-987:** Repository validation shall verify required metadata.
- **GH-REQ-988:** Repository validation shall verify filenames.
- **GH-REQ-989:** Repository validation shall verify internal links.
- **GH-REQ-990:** Repository validation shall verify ADR naming.
- **GH-REQ-991:** Repository validation shall verify diagram source and render paths.
- **GH-REQ-992:** Repository validation shall verify test structure.
- **GH-REQ-993:** Repository validation shall verify workflow paths.
- **GH-REQ-994:** Repository validation shall verify CODEOWNERS syntax where used.
- **GH-REQ-995:** Repository validation shall verify secret scanning.
- **GH-REQ-996:** Repository validation shall verify dependency manifests where applicable.
- **GH-REQ-997:** Repository validation shall verify generated-artifact consistency where applicable.
- **GH-REQ-998:** Repository validation shall verify release manifests where applicable.
- **GH-REQ-999:** Repository validation shall verify prohibited files and content.
- **GH-REQ-1000:** Repository validation shall verify status-claim accuracy.
- **GH-REQ-1001:** Repository validation shall identify actual platform settings separately from file-based validation.
- **GH-REQ-1002:** Repository validation shall record tool versions.
- **GH-REQ-1003:** Repository validation failures shall remain visible.
- **GH-REQ-1004:** Passing structural validation shall not establish semantic correctness.
- **GH-REQ-1005:** Repository validation evidence shall conform to `OBDIA-TEST-001`.

## 42. Structural Conformance Tests

- **GH-REQ-1006:** Structural tests shall verify `README.md`.
- **GH-REQ-1007:** Structural tests shall verify `SECURITY.md`.
- **GH-REQ-1008:** Structural tests shall verify the license file.
- **GH-REQ-1009:** Structural tests shall verify `CONTRIBUTING.md`.
- **GH-REQ-1010:** Structural tests shall verify `docs/`.
- **GH-REQ-1011:** Structural tests shall verify `adr/`.
- **GH-REQ-1012:** Structural tests shall verify `diagrams/`.
- **GH-REQ-1013:** Structural tests shall verify `tests/`.
- **GH-REQ-1014:** Structural tests shall verify `.github/workflows/` when CI is declared.
- **GH-REQ-1015:** Structural tests shall verify versioning metadata.
- **GH-REQ-1016:** Structural tests shall verify that repository structure reflects architectural dependencies.
- **GH-REQ-1017:** Structural tests shall verify no duplicate normative Knowledge filenames.
- **GH-REQ-1018:** Structural tests shall verify no Knowledge document above 30.
- **GH-REQ-1019:** Structural tests shall verify no `_ENTERPRISE` file is treated as a competing normative source.
- **GH-REQ-1020:** Structural tests shall verify no secret-bearing example files.
- **GH-REQ-1021:** Structural tests shall verify no real case data patterns in public fixtures where detection is feasible.
- **GH-REQ-1022:** Structural tests shall verify required indexes.
- **GH-REQ-1023:** Structural tests shall verify path ownership where configured.
- **GH-REQ-1024:** Structural tests shall verify branch and PR templates where required.
- **GH-REQ-1025:** Structural tests shall verify release paths where releases exist.
- **GH-REQ-1026:** Structural tests shall be reproducible.
- **GH-REQ-1027:** Structural-test failures shall block affected conformance claims.
- **GH-REQ-1028:** Structural tests shall not modify authoritative content.
- **GH-REQ-1029:** Structural tests shall retain evidence.
- **GH-REQ-1030:** Structural-test limitations shall remain visible.

## 43. Traceability

- **GH-REQ-1031:** Every authoritative repository path shall trace to an owner.
- **GH-REQ-1032:** Every normative document shall trace to its authority.
- **GH-REQ-1033:** Every requirement shall trace to governing documents or rationale.
- **GH-REQ-1034:** Every ADR shall trace to affected architecture and implementation.
- **GH-REQ-1035:** Every diagram shall trace to source and related documents.
- **GH-REQ-1036:** Every implementation component shall trace to requirements and ADRs.
- **GH-REQ-1037:** Every test shall trace to requirements or threats.
- **GH-REQ-1038:** Every validation decision shall trace to evidence.
- **GH-REQ-1039:** Every workflow shall trace to a repository control or automation purpose.
- **GH-REQ-1040:** Every dependency shall trace to a purpose and manifest.
- **GH-REQ-1041:** Every secret reference shall trace to approved secret management.
- **GH-REQ-1042:** Every release shall trace to exact source and validation evidence.
- **GH-REQ-1043:** Every exception shall trace to affected requirements.
- **GH-REQ-1044:** Every risk shall trace to affected repository assets and controls.
- **GH-REQ-1045:** Every change shall trace to issues, PRs or governance records.
- **GH-REQ-1046:** Every superseded artifact shall trace to its successor.
- **GH-REQ-1047:** Every archived artifact shall trace to retention authority.
- **GH-REQ-1048:** Traceability shall be bidirectional.
- **GH-REQ-1049:** Broken traceability affecting authority, security, evidence or release shall be blocking.
- **GH-REQ-1050:** Planned repository controls shall not be represented as enforced.
- **GH-REQ-1051:** Configured controls shall not be represented as tested without evidence.
- **GH-REQ-1052:** Tested controls shall not be represented as validated beyond scope.
- **GH-REQ-1053:** Traceability records shall not contain secrets.
- **GH-REQ-1054:** Traceability records shall minimize personal and case data.
- **GH-REQ-1055:** Baseline freeze shall validate repository traceability across documents 01–30.

## 44. Repository Review Gate

- **GH-REQ-1056:** Repository-standard review shall identify the exact document and commit.
- **GH-REQ-1057:** Review shall verify preservation of every legacy root artifact.
- **GH-REQ-1058:** Review shall verify preservation of CI and versioning requirements.
- **GH-REQ-1059:** Review shall verify that structure reflects architectural dependencies.
- **GH-REQ-1060:** Review shall verify canonical minimum layout.
- **GH-REQ-1061:** Review shall verify README, SECURITY, license and contribution requirements.
- **GH-REQ-1062:** Review shall verify Knowledge Pack and governance placement.
- **GH-REQ-1063:** Review shall verify ADR and diagram placement.
- **GH-REQ-1064:** Review shall verify code and test placement.
- **GH-REQ-1065:** Review shall verify branches, PRs and merge governance.
- **GH-REQ-1066:** Review shall verify workflow and CI security.
- **GH-REQ-1067:** Review shall verify dependency and secret controls.
- **GH-REQ-1068:** Review shall verify data and evidence controls.
- **GH-REQ-1069:** Review shall verify releases, tags and archival.
- **GH-REQ-1070:** Review shall verify settings and access governance.
- **GH-REQ-1071:** Review shall verify structural validation and traceability.
- **GH-REQ-1072:** Review shall identify residual risks and forward dependencies.
- **GH-REQ-1073:** Critical findings shall block progression.
- **GH-REQ-1074:** Material post-review changes shall invalidate affected evidence.
- **GH-REQ-1075:** Role concentration shall be disclosed.
- **GH-REQ-1076:** Internal review shall not be represented as independent external assurance.
- **GH-REQ-1077:** Documentation Authority shall assess repository information architecture.
- **GH-REQ-1078:** Security Reviewer shall assess access, workflow, secret and integrity controls.
- **GH-REQ-1079:** Implementation Reviewer shall assess code, dependency, test and automation structure.
- **GH-REQ-1080:** Privacy and Governance Reviewer shall assess data and public-content controls.
- **GH-REQ-1081:** Release Reviewer shall assess tags, releases and public artifacts.
- **GH-REQ-1082:** Project Founder approval shall not substitute for required specialist review.
- **GH-REQ-1083:** Automated structural checks shall remain advisory for governance semantics.
- **GH-REQ-1084:** Review non-applicability shall have rationale.
- **GH-REQ-1085:** Approval for Draft incorporation shall not prove repository-setting enforcement.

## 45. Minimum Validation Checklist

Before approval of this standard or a repository-conformance decision, confirm:

- [ ] `README.md` exists and states project boundaries accurately;
- [ ] `SECURITY.md` exists and defines safe reporting;
- [ ] a valid license decision exists;
- [ ] `CONTRIBUTING.md` exists;
- [ ] `docs/` exists;
- [ ] `adr/` exists;
- [ ] `diagrams/` exists;
- [ ] `tests/` exists;
- [ ] `.github/workflows/` exists when CI is declared;
- [ ] versioning is defined;
- [ ] structure reflects architectural dependencies;
- [ ] Knowledge Pack 01–30 identifiers and filenames are unique;
- [ ] no unauthorized normative document above 30 exists;
- [ ] governance authority hierarchy is preserved;
- [ ] CODEOWNERS or equivalent ownership is current;
- [ ] protected changes use pull requests and human review;
- [ ] required checks block failed or pending merges;
- [ ] workflows use least privilege;
- [ ] untrusted changes do not receive privileged secrets;
- [ ] secret scanning covers content and history where feasible;
- [ ] dependencies and generated files have provenance;
- [ ] tests use synthetic or authorized data;
- [ ] public content contains no real case data or credentials;
- [ ] tags and releases trace to exact commits;
- [ ] archive and retention paths are governed;
- [ ] structural validation is reproducible;
- [ ] traceability is bidirectional;
- [ ] forward dependencies 25–30 are recorded;
- [ ] Project Founder approval exists before status becomes Approved.


## 46. Limitations

- This standard does not prove that current GitHub repository settings are configured as described.
- It does not select a GitHub plan, organization model, CI product, dependency bot, scanning product or release mechanism.
- Some protections depend on platform and account capabilities.
- File-based configuration cannot represent every repository setting.
- CODEOWNERS does not guarantee review unless repository controls enforce it.
- Branch protection and workflow success cannot establish semantic correctness.
- Secret scanning cannot guarantee detection of every sensitive value.
- Repository history is widely replicated after cloning or forking and may not be fully retractable.
- Internal repository review is not independent certification.
- Documents 25–30 remain forward dependencies for versioning, risk, exception, change and compliance governance.


## 47. Change Control

Every material change shall identify rationale, affected paths, ownership, repository settings, workflows, branches, PRs, dependencies, data, security and privacy impact, release impact, migration, validation, rollback, authority and version effect.

- **GH-REQ-1086:** Editorial corrections shall use a patch version when meaning is unchanged.
- **GH-REQ-1087:** Backward-compatible substantive additions shall use a minor version.
- **GH-REQ-1088:** Incompatible repository-governance changes shall use a major version.
- **GH-REQ-1089:** Material repository-architecture decisions shall require an ADR where applicable.
- **GH-REQ-1090:** Changes shall require Project Founder approval.
- **GH-REQ-1091:** Changes shall receive Documentation Authority assessment.
- **GH-REQ-1092:** Security, implementation, privacy and release impacts shall receive specialist review where applicable.
- **GH-REQ-1093:** Changes shall identify affected paths, templates, rulesets, workflows and automation.
- **GH-REQ-1094:** Changes shall include migration and compatibility analysis.
- **GH-REQ-1095:** Changes shall include validation and rollback analysis.
- **GH-REQ-1096:** Changes shall not retroactively fabricate repository-setting or review evidence.
- **GH-REQ-1097:** Historical PR, workflow and release records shall not be silently rewritten.
- **GH-REQ-1098:** Document and requirement identifiers shall not be reused.
- **GH-REQ-1099:** Repository migrations shall preserve history and traceability.
- **GH-REQ-1100:** Forward-dependency reconciliation shall occur before this standard becomes Approved.

## 48. Consolidation Record

Version 1.0.0 consolidates the two existing GitHub Repository Standard variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-GH-001`;
- normalizes the authoritative filename to `24_GITHUB_REPOSITORY_STANDARD.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves `README.md`;
- preserves `SECURITY.md`;
- preserves `LICENSE`;
- preserves `CONTRIBUTING.md`;
- preserves `docs/`;
- preserves `adr/`;
- preserves `diagrams/`;
- preserves `tests/`;
- preserves continuous integration and `.github/workflows/`;
- preserves versioning;
- preserves the requirement that repository structure reflect architectural dependencies;
- adds canonical layout, ownership, branches, commits, pull requests, reviews, merge governance, rulesets, issues, workflows, CI, releases, dependencies, secrets, evidence, settings, access, tags, archival, incidents, validation and traceability controls;
- identifies documents 25–30 as forward dependencies;
- treats Enterprise and non-Enterprise legacy files as source variants of the same immutable document;
- creates no repository-setting claim, production deployment, release approval, compliance claim or operational authority.

## 49. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-06 | Draft | Documentation Authority; approval reserved to Project Founder | Constitutional consolidation of both legacy repository baselines: preserved all required root artifacts, CI, versioning and architecture-aligned structure; added complete ownership, workflow, review, security, release, validation and traceability requirements. |
| 1.0.1 | 2026-08-06 | Draft | Documentation Authority; approval reserved to Project Founder | Reconciled obsolete forward-dependency filenames with the final canonical Knowledge Pack 01–30 filenames; no requirement text, authority, lifecycle status, implementation state or validation claim changed. |
