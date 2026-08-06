# AI RISK REGISTER

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-RISK-001 |
| **Title** | AI Risk Register |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Privacy and Governance Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define the authoritative OBDIA AI-risk record schema, rating method, ownership, treatment, acceptance, review, escalation, release gating, evidence and traceability requirements, and establish the initial synthetic baseline risk catalogue. |
| **Scope** | Risks arising from AI models, prompts, agents, human oversight, identity and delegation, authorization, evidence handling, connectors, external services, data, privacy, fundamental rights, software, supply chain, repository governance, testing, releases and research assumptions within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `07_THREAT_MODEL_BASELINE.md`; `09_GOVERNANCE_MODEL.md`; `15_RELEASE_AND_PUBLICATION_POLICY.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; `18_EVIDENCE_MODEL.md`; `19_AGENT_LIFECYCLE_MODEL.md`; `20_CONNECTOR_SECURITY_POLICY.md`; `21_SECURE_CODING_STANDARD.md`; `22_TESTING_STANDARD.md`; `24_GITHUB_REPOSITORY_STANDARD.md`; `25_DOCUMENT_VERSIONING_POLICY.md`; forward dependencies `27_RESEARCH_BACKLOG.md`, `28_ASSUMPTIONS_REGISTER.md`, `29_DECISION_LOG_POLICY.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-RES-001`; `OBDIA-ARCH-001`; `OBDIA-TM-001`; `OBDIA-GOV-001`; `OBDIA-REL-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001`; `OBDIA-EVID-001`; `OBDIA-AGENT-001`; `OBDIA-CONN-001`; `OBDIA-CODE-001`; `OBDIA-TEST-001`; `OBDIA-GH-001`; `OBDIA-VER-001` |
| **Cross-References** | Threats, controls, assumptions, research items, ADRs, decisions, requirements, tests, incidents, exceptions, release records, evidence packages, baseline manifests and compliance mappings |
| **Assumptions** | The initial risk catalogue is a research baseline derived from the current OBDIA architecture and uses synthetic scenarios. It is not a record of an operational institution, real investigation or deployed system. |
| **Constraints** | Risk treatment or acceptance shall not authorize canonical prohibitions, unlawful access, offensive operations, real criminal-infrastructure interaction, unlawful surveillance, autonomous legal or coercive action, credential transfer, or independent AI legal authority. |
| **Security Considerations** | A stale, incomplete or manipulated risk register can conceal critical threats, normalize unsafe residual risk, weaken release gates, misdirect testing and create misleading assurance claims. |
| **Validation Criteria** | Every active risk has a unique immutable identifier, complete schema, justified inherent and residual ratings, human owner, linked controls and evidence, review date, treatment or acceptance decision, release impact and bidirectional traceability. |
| **Implementation Relationship** | This document governs risk records and contains an initial research baseline. It does not accept risk, implement controls, authorize deployment, prove compliance or establish production readiness. |

---

## 1. Purpose

The two legacy source variants required the register to track:

- risk or description;
- category;
- cause;
- impact;
- likelihood;
- risk level;
- controls;
- residual risk;
- owner;
- review date;
- status.

The Enterprise source additionally required an identifier and review before every release.

This consolidation preserves every field and the pre-release review requirement. It adds rating definitions, acceptance authority, treatment states, escalation, evidence, privacy and fundamental-rights impacts, immutable history, and traceability to threats, controls, tests and compliance mappings.

The policy sections are normative. Individual risk records are governed risk evidence; they do not change architecture, authorize implementation or waive higher authority.

## 2. Fundamental Risk-Governance Rules

- **RISK-REQ-001:** Every material AI-related risk shall be recorded.
- **RISK-REQ-002:** Every risk shall have an attributable human owner.
- **RISK-REQ-003:** An AI system shall not own or accept risk.
- **RISK-REQ-004:** Risk acceptance shall remain distinct from risk identification.
- **RISK-REQ-005:** Risk acceptance shall remain distinct from control implementation.
- **RISK-REQ-006:** Risk acceptance shall remain distinct from legal compliance.
- **RISK-REQ-007:** Risk ratings shall support prioritization and shall not be represented as mathematical certainty.
- **RISK-REQ-008:** Risk records shall distinguish inherent risk from residual risk.
- **RISK-REQ-009:** Risk records shall distinguish control design from control effectiveness.
- **RISK-REQ-010:** Risk records shall distinguish current evidence from expected evidence.
- **RISK-REQ-011:** Missing material risk information shall not default to a low rating.
- **RISK-REQ-012:** Critical unresolved risk shall block affected approval, activation, validation or release.
- **RISK-REQ-013:** Risk governance shall preserve human accountability.
- **RISK-REQ-014:** Risk governance shall preserve evidence integrity and chain-of-custody requirements.
- **RISK-REQ-015:** Risk governance shall preserve officer-agent binding and scoped delegation.
- **RISK-REQ-016:** Risk governance shall preserve connector isolation and revocation.
- **RISK-REQ-017:** Risk governance shall preserve canonical safety prohibitions.
- **RISK-REQ-018:** Risk records shall minimize personal and case data.
- **RISK-REQ-019:** Risk records shall not use real criminal infrastructure as test evidence.
- **RISK-REQ-020:** Every release shall include a review of applicable open risks.

## 3. Authority and Register Status

- **RISK-REQ-021:** This document shall define the normative risk-register method.
- **RISK-REQ-022:** Individual risk records shall be governed records subordinate to this method.
- **RISK-REQ-023:** Risk records shall not alter approved requirements directly.
- **RISK-REQ-024:** Risk treatment requiring architecture change shall use the applicable ADR and change process.
- **RISK-REQ-025:** Risk treatment requiring research shall create or link a governed research item.
- **RISK-REQ-026:** Risk acceptance decisions shall identify the approving authority.
- **RISK-REQ-027:** Risk closure decisions shall identify the closing authority.
- **RISK-REQ-028:** Risk escalation decisions shall identify the escalating authority.
- **RISK-REQ-029:** The Security Reviewer shall review security-critical risks.
- **RISK-REQ-030:** The Implementation Reviewer shall review implementation and supply-chain risks.
- **RISK-REQ-031:** The Release Reviewer shall review release-affecting risks.
- **RISK-REQ-032:** The Project Founder shall approve high residual-risk acceptance where permitted.
- **RISK-REQ-033:** Critical residual risk shall not be accepted for an approved baseline or release.
- **RISK-REQ-034:** Repository merge shall not establish risk acceptance.
- **RISK-REQ-035:** Automated scoring shall not establish risk acceptance.
- **RISK-REQ-036:** Role concentration shall be disclosed.
- **RISK-REQ-037:** Internal review shall not be represented as independent external assurance.
- **RISK-REQ-038:** Register status shall remain distinct from individual risk status.
- **RISK-REQ-039:** Risk records shall identify the method version used for rating.
- **RISK-REQ-040:** Changes to rating methodology shall not silently recalculate historical ratings.

## 4. Mandatory Risk Record Schema

The following fields form the minimum authoritative schema. A field may be marked not applicable only with an attributable rationale.

- **RISK-REQ-041:** Every risk record shall include `ID` as immutable risk identifier.
- **RISK-REQ-042:** Every risk record shall include `Title` as concise risk title.
- **RISK-REQ-043:** Every risk record shall include `Description` as clear risk scenario.
- **RISK-REQ-044:** Every risk record shall include `Category` as controlled risk category.
- **RISK-REQ-045:** Every risk record shall include `Cause` as credible initiating conditions.
- **RISK-REQ-046:** Every risk record shall include `Threat references` as linked threat identifiers.
- **RISK-REQ-047:** Every risk record shall include `Asset references` as affected assets or capabilities.
- **RISK-REQ-048:** Every risk record shall include `Affected parties` as people, groups or institutions that may be harmed.
- **RISK-REQ-049:** Every risk record shall include `Impact` as impact narrative.
- **RISK-REQ-050:** Every risk record shall include `Impact dimensions` as affected impact dimensions.
- **RISK-REQ-051:** Every risk record shall include `Likelihood` as likelihood rating.
- **RISK-REQ-052:** Every risk record shall include `Impact rating` as impact severity rating.
- **RISK-REQ-053:** Every risk record shall include `Risk score` as derived prioritization score.
- **RISK-REQ-054:** Every risk record shall include `Risk level` as qualitative level.
- **RISK-REQ-055:** Every risk record shall include `Rating confidence` as confidence in the assessment.
- **RISK-REQ-056:** Every risk record shall include `Controls` as current control references.
- **RISK-REQ-057:** Every risk record shall include `Planned treatments` as planned treatment references.
- **RISK-REQ-058:** Every risk record shall include `Control effectiveness` as supported effectiveness assessment.
- **RISK-REQ-059:** Every risk record shall include `Residual likelihood` as post-control likelihood.
- **RISK-REQ-060:** Every risk record shall include `Residual impact` as post-control impact.
- **RISK-REQ-061:** Every risk record shall include `Residual risk` as residual score and level.
- **RISK-REQ-062:** Every risk record shall include `Owner` as accountable human risk owner.
- **RISK-REQ-063:** Every risk record shall include `Contributors` as supporting reviewers or contributors.
- **RISK-REQ-064:** Every risk record shall include `Review date` as next mandatory review date.
- **RISK-REQ-065:** Every risk record shall include `Status` as current risk lifecycle state.
- **RISK-REQ-066:** Every risk record shall include `Acceptance authority` as human authority where acceptance is permitted.
- **RISK-REQ-067:** Every risk record shall include `Acceptance expiry` as expiry of any temporary acceptance.
- **RISK-REQ-068:** Every risk record shall include `Evidence` as supporting evidence references.
- **RISK-REQ-069:** Every risk record shall include `Test references` as linked validation and negative tests.
- **RISK-REQ-070:** Every risk record shall include `Incident references` as linked incidents.
- **RISK-REQ-071:** Every risk record shall include `Assumption references` as linked assumptions.
- **RISK-REQ-072:** Every risk record shall include `Research references` as linked research items.
- **RISK-REQ-073:** Every risk record shall include `Decision references` as linked ADRs or decisions.
- **RISK-REQ-074:** Every risk record shall include `Release impact` as effect on release eligibility.
- **RISK-REQ-075:** Every risk record shall include `Closure evidence` as evidence supporting closure.
- **RISK-REQ-076:** Every risk record shall include `Revision history` as append-only record of material changes.

## 6. Risk Identifier Rules

- **RISK-REQ-077:** Risk record identifiers shall use the controlled format `AI-RISK-NNN`.
- **RISK-REQ-078:** The `AI` prefix shall identify the OBDIA AI-risk register namespace.
- **RISK-REQ-079:** Risk numbers shall be zero-padded to at least three digits.
- **RISK-REQ-080:** Risk identifiers shall be unique within the register.
- **RISK-REQ-081:** Risk identifiers shall not be reused.
- **RISK-REQ-082:** Risk identifiers shall not encode severity.
- **RISK-REQ-083:** Risk identifiers shall not encode lifecycle status.
- **RISK-REQ-084:** Risk identifiers shall not encode owner.
- **RISK-REQ-085:** Risk identifiers shall not contain case names.
- **RISK-REQ-086:** Risk identifiers shall not contain personal names.
- **RISK-REQ-087:** Duplicate identifiers shall be rejected.
- **RISK-REQ-088:** Temporary spreadsheet row numbers shall not replace risk identifiers.
- **RISK-REQ-089:** GitHub issue numbers may supplement but shall not replace risk identifiers.
- **RISK-REQ-090:** Threat identifiers may supplement but shall not replace risk identifiers.
- **RISK-REQ-091:** Risk mergers shall preserve all predecessor identifiers and disposition.
- **RISK-REQ-092:** Risk splits shall preserve the original record and identify successor records.
- **RISK-REQ-093:** Superseded risk records shall identify successors.
- **RISK-REQ-094:** Closed risk identifiers shall remain reserved.
- **RISK-REQ-095:** Unknown legacy identifiers shall remain documented as unknown.
- **RISK-REQ-096:** Identifier validation shall be automated where practical.

## 7. Risk Taxonomy

- **RISK-REQ-097:** `Authority and Accountability` shall be an available controlled risk category for loss of human control, unclear responsibility or unauthorized AI authority.
- **RISK-REQ-098:** Risks classified as `Authority and Accountability` shall identify the specific affected requirements and assets.
- **RISK-REQ-099:** Use of `Authority and Accountability` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-100:** `Identity and Delegation` shall be an available controlled risk category for binding, delegation, credential or machine-identity failure.
- **RISK-REQ-101:** Risks classified as `Identity and Delegation` shall identify the specific affected requirements and assets.
- **RISK-REQ-102:** Use of `Identity and Delegation` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-103:** `Authorization and Access Control` shall be an available controlled risk category for policy, scope, obligation or enforcement failure.
- **RISK-REQ-104:** Risks classified as `Authorization and Access Control` shall identify the specific affected requirements and assets.
- **RISK-REQ-105:** Use of `Authorization and Access Control` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-106:** `AI Model and Prompt` shall be an available controlled risk category for hallucination, prompt injection, model behavior or model-change risk.
- **RISK-REQ-107:** Risks classified as `AI Model and Prompt` shall identify the specific affected requirements and assets.
- **RISK-REQ-108:** Use of `AI Model and Prompt` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-109:** `Data and Privacy` shall be an available controlled risk category for overcollection, leakage, unlawful processing or retention risk.
- **RISK-REQ-110:** Risks classified as `Data and Privacy` shall identify the specific affected requirements and assets.
- **RISK-REQ-111:** Use of `Data and Privacy` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-112:** `Fundamental Rights` shall be an available controlled risk category for discrimination, unlawful surveillance, deanonymization or coercive impact.
- **RISK-REQ-113:** Risks classified as `Fundamental Rights` shall identify the specific affected requirements and assets.
- **RISK-REQ-114:** Use of `Fundamental Rights` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-115:** `Evidence and Forensics` shall be an available controlled risk category for integrity, provenance, custody, authenticity or derived-output confusion.
- **RISK-REQ-116:** Risks classified as `Evidence and Forensics` shall identify the specific affected requirements and assets.
- **RISK-REQ-117:** Use of `Evidence and Forensics` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-118:** `Connector and External Service` shall be an available controlled risk category for external integration, endpoint, credential or provider risk.
- **RISK-REQ-119:** Risks classified as `Connector and External Service` shall identify the specific affected requirements and assets.
- **RISK-REQ-120:** Use of `Connector and External Service` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-121:** `Software and Supply Chain` shall be an available controlled risk category for code, dependency, build, package or workflow compromise.
- **RISK-REQ-122:** Risks classified as `Software and Supply Chain` shall identify the specific affected requirements and assets.
- **RISK-REQ-123:** Use of `Software and Supply Chain` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-124:** `Repository and Governance` shall be an available controlled risk category for history, approval, documentation, decision or release-governance failure.
- **RISK-REQ-125:** Risks classified as `Repository and Governance` shall identify the specific affected requirements and assets.
- **RISK-REQ-126:** Use of `Repository and Governance` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-127:** `Operational Resilience` shall be an available controlled risk category for availability, recovery, revocation propagation or resource exhaustion.
- **RISK-REQ-128:** Risks classified as `Operational Resilience` shall identify the specific affected requirements and assets.
- **RISK-REQ-129:** Use of `Operational Resilience` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-130:** `Human Factors` shall be an available controlled risk category for automation bias, skill gaps, misuse or inadequate oversight.
- **RISK-REQ-131:** Risks classified as `Human Factors` shall identify the specific affected requirements and assets.
- **RISK-REQ-132:** Use of `Human Factors` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-133:** `Legal and Regulatory` shall be an available controlled risk category for jurisdiction, legal basis, retention or compliance-claim risk.
- **RISK-REQ-134:** Risks classified as `Legal and Regulatory` shall identify the specific affected requirements and assets.
- **RISK-REQ-135:** Use of `Legal and Regulatory` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-136:** `Publication and Reputation` shall be an available controlled risk category for unsafe disclosure, misleading claims or reputational harm.
- **RISK-REQ-137:** Risks classified as `Publication and Reputation` shall identify the specific affected requirements and assets.
- **RISK-REQ-138:** Use of `Publication and Reputation` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-139:** `Research and Assumptions` shall be an available controlled risk category for unsupported hypothesis, stale assumption or uncertain source risk.
- **RISK-REQ-140:** Risks classified as `Research and Assumptions` shall identify the specific affected requirements and assets.
- **RISK-REQ-141:** Use of `Research and Assumptions` shall not exclude secondary categories when multiple dimensions apply.
- **RISK-REQ-142:** `Web3 and Distributed Systems` shall be an available controlled risk category for blockchain, smart-contract, wallet, finality or distributed-state risk.
- **RISK-REQ-143:** Risks classified as `Web3 and Distributed Systems` shall identify the specific affected requirements and assets.
- **RISK-REQ-144:** Use of `Web3 and Distributed Systems` shall not exclude secondary categories when multiple dimensions apply.

## 8. Risk Source and Intake

- **RISK-REQ-145:** Risk intake may originate from threat modeling.
- **RISK-REQ-146:** Risk intake may originate from architecture review.
- **RISK-REQ-147:** Risk intake may originate from test failure.
- **RISK-REQ-148:** Risk intake may originate from security incidents.
- **RISK-REQ-149:** Risk intake may originate from privacy review.
- **RISK-REQ-150:** Risk intake may originate from evidence-integrity findings.
- **RISK-REQ-151:** Risk intake may originate from dependency or supply-chain findings.
- **RISK-REQ-152:** Risk intake may originate from research findings.
- **RISK-REQ-153:** Risk intake may originate from assumption invalidation.
- **RISK-REQ-154:** Risk intake may originate from external standards or legal-source updates.
- **RISK-REQ-155:** Risk intake may originate from user or contributor reports.
- **RISK-REQ-156:** Risk intake shall identify its source.
- **RISK-REQ-157:** Risk intake shall identify the date received.
- **RISK-REQ-158:** Risk intake shall identify the submitting actor or channel.
- **RISK-REQ-159:** Unverified submissions shall be labeled.
- **RISK-REQ-160:** Duplicate submissions shall link to the canonical risk.
- **RISK-REQ-161:** Potential critical risks shall be triaged immediately.
- **RISK-REQ-162:** Sensitive reports shall use restricted handling.
- **RISK-REQ-163:** Intake shall not include active credentials.
- **RISK-REQ-164:** Intake shall preserve source provenance.

## 9. Risk Lifecycle States

- **RISK-REQ-165:** `Identified` shall be a controlled risk status meaning recorded but not fully assessed.
- **RISK-REQ-166:** Entry into `Identified` shall be attributable and timestamped.
- **RISK-REQ-167:** Exit from `Identified` shall identify the next status and transition evidence.
- **RISK-REQ-168:** `Assessing` shall be a controlled risk status meaning under structured analysis.
- **RISK-REQ-169:** Entry into `Assessing` shall be attributable and timestamped.
- **RISK-REQ-170:** Exit from `Assessing` shall identify the next status and transition evidence.
- **RISK-REQ-171:** `Assessed` shall be a controlled risk status meaning inherent and residual ratings completed.
- **RISK-REQ-172:** Entry into `Assessed` shall be attributable and timestamped.
- **RISK-REQ-173:** Exit from `Assessed` shall identify the next status and transition evidence.
- **RISK-REQ-174:** `Treatment Planned` shall be a controlled risk status meaning approved treatment plan exists.
- **RISK-REQ-175:** Entry into `Treatment Planned` shall be attributable and timestamped.
- **RISK-REQ-176:** Exit from `Treatment Planned` shall identify the next status and transition evidence.
- **RISK-REQ-177:** `Treating` shall be a controlled risk status meaning treatment implementation or validation is in progress.
- **RISK-REQ-178:** Entry into `Treating` shall be attributable and timestamped.
- **RISK-REQ-179:** Exit from `Treating` shall identify the next status and transition evidence.
- **RISK-REQ-180:** `Monitoring` shall be a controlled risk status meaning residual risk is under active observation.
- **RISK-REQ-181:** Entry into `Monitoring` shall be attributable and timestamped.
- **RISK-REQ-182:** Exit from `Monitoring` shall identify the next status and transition evidence.
- **RISK-REQ-183:** `Accepted` shall be a controlled risk status meaning residual risk has time-bounded human acceptance.
- **RISK-REQ-184:** Entry into `Accepted` shall be attributable and timestamped.
- **RISK-REQ-185:** Exit from `Accepted` shall identify the next status and transition evidence.
- **RISK-REQ-186:** `Closed` shall be a controlled risk status meaning risk is eliminated, no longer applicable or otherwise closed with evidence.
- **RISK-REQ-187:** Entry into `Closed` shall be attributable and timestamped.
- **RISK-REQ-188:** Exit from `Closed` shall identify the next status and transition evidence.
- **RISK-REQ-189:** `Superseded` shall be a controlled risk status meaning record replaced by one or more successor risks.
- **RISK-REQ-190:** Entry into `Superseded` shall be attributable and timestamped.
- **RISK-REQ-191:** Exit from `Superseded` shall identify the next status and transition evidence.

## 10. Risk Lifecycle Transitions

- **RISK-REQ-192:** New risk records shall enter `Identified`.
- **RISK-REQ-193:** `Identified` risks may transition to `Assessing`.
- **RISK-REQ-194:** `Assessed` risks may transition to `Treatment Planned`.
- **RISK-REQ-195:** `Assessed` risks may transition to `Monitoring` when no additional treatment is currently required and evidence supports the decision.
- **RISK-REQ-196:** `Assessed` risks may transition to `Accepted` only through authorized acceptance.
- **RISK-REQ-197:** `Treatment Planned` risks may transition to `Treating`.
- **RISK-REQ-198:** `Treating` risks may return to `Treatment Planned` when treatment fails or changes.
- **RISK-REQ-199:** `Monitoring` risks may return to `Assessing` when conditions change.
- **RISK-REQ-200:** `Accepted` risks shall return to `Assessing` before acceptance expiry.
- **RISK-REQ-201:** `Accepted` risks shall return to `Assessing` when a review trigger occurs.
- **RISK-REQ-202:** Any risk may transition to `Superseded` only when successor records are identified.
- **RISK-REQ-203:** `Closed` risks may be reopened through a new attributable transition.
- **RISK-REQ-204:** `Superseded` risks shall not be reopened; successor records shall be used.
- **RISK-REQ-205:** Status transitions shall not erase prior states.
- **RISK-REQ-206:** Concurrent status changes shall resolve through authorized review.
- **RISK-REQ-207:** More restrictive release treatment shall prevail during status conflict.
- **RISK-REQ-208:** Risk status shall not be inferred from issue status.
- **RISK-REQ-209:** Risk status shall not be inferred from PR status.
- **RISK-REQ-210:** Risk status shall not be inferred from absence of incidents.
- **RISK-REQ-211:** Lifecycle events shall remain retrievable.

## 11. Impact Dimensions

- **RISK-REQ-212:** Every risk assessment shall consider `Security` impact covering confidentiality, integrity, availability or control compromise.
- **RISK-REQ-213:** A non-applicable `Security` impact shall have rationale.
- **RISK-REQ-214:** Material `Security` impact shall identify affected parties and assets.
- **RISK-REQ-215:** Every risk assessment shall consider `Privacy` impact covering personal-data misuse, overcollection, exposure or unlawful processing.
- **RISK-REQ-216:** A non-applicable `Privacy` impact shall have rationale.
- **RISK-REQ-217:** Material `Privacy` impact shall identify affected parties and assets.
- **RISK-REQ-218:** Every risk assessment shall consider `Fundamental Rights` impact covering discrimination, surveillance, deanonymization, due-process or autonomy harm.
- **RISK-REQ-219:** A non-applicable `Fundamental Rights` impact shall have rationale.
- **RISK-REQ-220:** Material `Fundamental Rights` impact shall identify affected parties and assets.
- **RISK-REQ-221:** Every risk assessment shall consider `Human Safety` impact covering physical or psychological harm to people.
- **RISK-REQ-222:** A non-applicable `Human Safety` impact shall have rationale.
- **RISK-REQ-223:** Material `Human Safety` impact shall identify affected parties and assets.
- **RISK-REQ-224:** Every risk assessment shall consider `Evidence Integrity` impact covering provenance, custody, authenticity, reliability or admissibility-related harm.
- **RISK-REQ-225:** A non-applicable `Evidence Integrity` impact shall have rationale.
- **RISK-REQ-226:** Material `Evidence Integrity` impact shall identify affected parties and assets.
- **RISK-REQ-227:** Every risk assessment shall consider `Legal and Governance` impact covering lack of authority, jurisdictional conflict, policy breach or misleading compliance claim.
- **RISK-REQ-228:** A non-applicable `Legal and Governance` impact shall have rationale.
- **RISK-REQ-229:** Material `Legal and Governance` impact shall identify affected parties and assets.
- **RISK-REQ-230:** Every risk assessment shall consider `Operational` impact covering service disruption, workflow failure, recovery difficulty or loss of mission capability.
- **RISK-REQ-231:** A non-applicable `Operational` impact shall have rationale.
- **RISK-REQ-232:** Material `Operational` impact shall identify affected parties and assets.
- **RISK-REQ-233:** Every risk assessment shall consider `Financial and Resource` impact covering direct cost, fraud, asset loss or uncontrolled consumption.
- **RISK-REQ-234:** A non-applicable `Financial and Resource` impact shall have rationale.
- **RISK-REQ-235:** Material `Financial and Resource` impact shall identify affected parties and assets.
- **RISK-REQ-236:** Every risk assessment shall consider `Reputation and Trust` impact covering loss of public, institutional or stakeholder confidence.
- **RISK-REQ-237:** A non-applicable `Reputation and Trust` impact shall have rationale.
- **RISK-REQ-238:** Material `Reputation and Trust` impact shall identify affected parties and assets.
- **RISK-REQ-239:** Every risk assessment shall consider `Research Validity` impact covering invalid conclusions, irreproducibility or unsupported architectural decisions.
- **RISK-REQ-240:** A non-applicable `Research Validity` impact shall have rationale.
- **RISK-REQ-241:** Material `Research Validity` impact shall identify affected parties and assets.

## 12. Likelihood Scale

- **RISK-REQ-242:** Likelihood score `1` shall be labeled `Rare`.
- **RISK-REQ-243:** `Rare` shall mean: Not expected under current conditions; credible occurrence requires exceptional circumstances.
- **RISK-REQ-244:** Use of `Rare` shall identify supporting evidence and assumptions.
- **RISK-REQ-245:** Likelihood `Rare` shall be reassessed when relevant controls, threats or exposure change.
- **RISK-REQ-246:** Likelihood score `2` shall be labeled `Unlikely`.
- **RISK-REQ-247:** `Unlikely` shall mean: Could occur, but evidence and conditions indicate infrequent occurrence.
- **RISK-REQ-248:** Use of `Unlikely` shall identify supporting evidence and assumptions.
- **RISK-REQ-249:** Likelihood `Unlikely` shall be reassessed when relevant controls, threats or exposure change.
- **RISK-REQ-250:** Likelihood score `3` shall be labeled `Possible`.
- **RISK-REQ-251:** `Possible` shall mean: Plausible under foreseeable conditions or observed occasionally in comparable contexts.
- **RISK-REQ-252:** Use of `Possible` shall identify supporting evidence and assumptions.
- **RISK-REQ-253:** Likelihood `Possible` shall be reassessed when relevant controls, threats or exposure change.
- **RISK-REQ-254:** Likelihood score `4` shall be labeled `Likely`.
- **RISK-REQ-255:** `Likely` shall mean: Expected under several normal or adversarial conditions, or repeatedly observed.
- **RISK-REQ-256:** Use of `Likely` shall identify supporting evidence and assumptions.
- **RISK-REQ-257:** Likelihood `Likely` shall be reassessed when relevant controls, threats or exposure change.
- **RISK-REQ-258:** Likelihood score `5` shall be labeled `Almost Certain`.
- **RISK-REQ-259:** `Almost Certain` shall mean: Expected frequently, currently occurring, or strongly indicated by direct evidence.
- **RISK-REQ-260:** Use of `Almost Certain` shall identify supporting evidence and assumptions.
- **RISK-REQ-261:** Likelihood `Almost Certain` shall be reassessed when relevant controls, threats or exposure change.

## 13. Impact Scale

- **RISK-REQ-262:** Impact score `1` shall be labeled `Negligible`.
- **RISK-REQ-263:** `Negligible` shall mean: Minimal reversible impact with no material authority, evidence, privacy or operational consequence.
- **RISK-REQ-264:** Use of `Negligible` shall identify the affected impact dimensions.
- **RISK-REQ-265:** Impact `Negligible` shall use the highest credible material dimension unless a documented method justifies another aggregation.
- **RISK-REQ-266:** Impact score `2` shall be labeled `Minor`.
- **RISK-REQ-267:** `Minor` shall mean: Limited and contained impact requiring routine correction with no serious rights or evidence consequence.
- **RISK-REQ-268:** Use of `Minor` shall identify the affected impact dimensions.
- **RISK-REQ-269:** Impact `Minor` shall use the highest credible material dimension unless a documented method justifies another aggregation.
- **RISK-REQ-270:** Impact score `3` shall be labeled `Moderate`.
- **RISK-REQ-271:** `Moderate` shall mean: Material impact requiring managed response, remediation or restricted operation.
- **RISK-REQ-272:** Use of `Moderate` shall identify the affected impact dimensions.
- **RISK-REQ-273:** Impact `Moderate` shall use the highest credible material dimension unless a documented method justifies another aggregation.
- **RISK-REQ-274:** Impact score `4` shall be labeled `Major`.
- **RISK-REQ-275:** `Major` shall mean: Serious impact to security, privacy, rights, evidence, governance or operations requiring senior intervention.
- **RISK-REQ-276:** Use of `Major` shall identify the affected impact dimensions.
- **RISK-REQ-277:** Impact `Major` shall use the highest credible material dimension unless a documented method justifies another aggregation.
- **RISK-REQ-278:** Impact score `5` shall be labeled `Severe`.
- **RISK-REQ-279:** `Severe` shall mean: Catastrophic, unlawful, systemic or irreversible impact, including canonical-prohibition scenarios.
- **RISK-REQ-280:** Use of `Severe` shall identify the affected impact dimensions.
- **RISK-REQ-281:** Impact `Severe` shall use the highest credible material dimension unless a documented method justifies another aggregation.

## 14. Risk Score and Level

- **RISK-REQ-282:** Risk score shall equal likelihood score multiplied by impact score.
- **RISK-REQ-283:** Risk score shall range from 1 to 25.
- **RISK-REQ-284:** Risk score shall not be represented as exact probability or expected loss.
- **RISK-REQ-285:** Risk score shall not override canonical prohibitions.
- **RISK-REQ-286:** Risk score shall not override qualitative expert judgment.
- **RISK-REQ-287:** The rating method version shall be recorded.
- **RISK-REQ-288:** Alternative quantitative methods shall require an ADR and compatibility plan.
- **RISK-REQ-289:** Risks with severe rights, authority or evidence impacts may be elevated regardless of arithmetic score.
- **RISK-REQ-290:** Multiple affected dimensions shall be recorded even when one drives the level.
- **RISK-REQ-291:** Score calculation errors shall be corrected additively.
- **RISK-REQ-292:** Historical scores shall remain visible.
- **RISK-REQ-293:** `Low` shall be the qualitative risk level for scores `1–4`.
- **RISK-REQ-294:** `Low` shall not be assigned solely from intuition without likelihood and impact records.
- **RISK-REQ-295:** `Moderate` shall be the qualitative risk level for scores `5–9`.
- **RISK-REQ-296:** `Moderate` risks shall receive treatment and review proportional to their level.
- **RISK-REQ-297:** `High` shall be the qualitative risk level for scores `10–16`.
- **RISK-REQ-298:** `High` risks shall receive treatment and review proportional to their level.
- **RISK-REQ-299:** `High` shall not be assigned solely from intuition without likelihood and impact records.
- **RISK-REQ-300:** `Critical` risks shall receive treatment and review proportional to their level.
- **RISK-REQ-301:** `Critical` shall not be assigned solely from intuition without likelihood and impact records.

## 15. Confidence and Uncertainty

- **RISK-REQ-302:** Every likelihood and impact rating shall include a confidence assessment.
- **RISK-REQ-303:** Unknown information shall be recorded explicitly.
- **RISK-REQ-304:** Conflicting evidence shall be recorded.
- **RISK-REQ-305:** Evidence age shall be considered.
- **RISK-REQ-306:** Model and provider opacity shall be considered.
- **RISK-REQ-307:** Novelty shall be considered.
- **RISK-REQ-308:** Environmental differences shall be considered.
- **RISK-REQ-309:** Assumption dependence shall be considered.
- **RISK-REQ-310:** Low-confidence high-impact risks shall be treated conservatively.
- **RISK-REQ-311:** Uncertainty reduction may be a treatment objective.
- **RISK-REQ-312:** Confidence shall remain distinct from evidence integrity.
- **RISK-REQ-313:** Confidence shall remain distinct from compliance.
- **RISK-REQ-314:** Rating confidence `Low` shall mean material evidence is missing, assumptions dominate or the system is highly novel.
- **RISK-REQ-315:** `Low` confidence shall identify the principal uncertainty.
- **RISK-REQ-316:** `Low` confidence shall not alter the numerical score silently.
- **RISK-REQ-317:** Rating confidence `Medium` shall mean some relevant evidence exists but uncertainty remains.
- **RISK-REQ-318:** `Medium` confidence shall not alter the numerical score silently.
- **RISK-REQ-319:** Rating confidence `High` shall mean direct, current and reproducible evidence supports the rating.
- **RISK-REQ-320:** `High` confidence shall identify the principal uncertainty.
- **RISK-REQ-321:** `High` confidence shall not alter the numerical score silently.

## 16. Inherent Risk Assessment

- **RISK-REQ-322:** Inherent risk shall be assessed before crediting current controls.
- **RISK-REQ-323:** Inherent risk shall identify the uncontrolled scenario.
- **RISK-REQ-324:** Inherent likelihood shall be recorded.
- **RISK-REQ-325:** Inherent impact shall be recorded.
- **RISK-REQ-326:** Inherent score shall be calculated.
- **RISK-REQ-327:** Inherent level shall be recorded.
- **RISK-REQ-328:** Inherent rating confidence shall be recorded.
- **RISK-REQ-329:** Inherent risk shall identify exposed assets.
- **RISK-REQ-330:** Inherent risk shall identify affected parties.
- **RISK-REQ-331:** Inherent risk shall identify threat and cause references.
- **RISK-REQ-332:** Inherent risk shall not assume planned controls exist.
- **RISK-REQ-333:** Inherent risk shall not assume human review succeeds without evidence.
- **RISK-REQ-334:** Inherent risk shall not assume provider claims are effective controls.
- **RISK-REQ-335:** Inherent risk shall not assume network location is trust.
- **RISK-REQ-336:** Inherent risk shall not assume encryption establishes authorization.
- **RISK-REQ-337:** Inherent risk shall not assume public availability establishes lawful processing.
- **RISK-REQ-338:** Inherent risk shall identify canonical-prohibition relationships.
- **RISK-REQ-339:** Inherent risk changes shall preserve history.
- **RISK-REQ-340:** Inherent risk shall be reviewed after material scope change.
- **RISK-REQ-341:** Inherent risk shall be reviewed after material threat change.

## 17. Control Records and Linkage

- **RISK-REQ-342:** Every credited control shall record `Control identifier`.
- **RISK-REQ-343:** Every credited control shall record `Control owner`.
- **RISK-REQ-344:** Every credited control shall record `Control type`.
- **RISK-REQ-345:** Every credited control shall record `Preventive, detective, corrective or recovery function`.
- **RISK-REQ-346:** Every credited control shall record `Applicable requirements`.
- **RISK-REQ-347:** Every credited control shall record `Test evidence`.
- **RISK-REQ-348:** Every credited control shall record `Effectiveness conclusion`.
- **RISK-REQ-349:** Every credited control shall record `Dependencies`.
- **RISK-REQ-350:** Every credited control shall record `Monitoring method`.
- **RISK-REQ-351:** Every credited control shall record `Review date`.
- **RISK-REQ-352:** Controls shall use immutable identifiers.
- **RISK-REQ-353:** Implemented controls shall be distinguished from validated controls.
- **RISK-REQ-354:** Control presence shall not imply effectiveness.
- **RISK-REQ-355:** Control evidence shall identify exact versions.
- **RISK-REQ-356:** Control failure shall trigger risk reassessment.
- **RISK-REQ-357:** Compensating controls shall identify the control gap they address.
- **RISK-REQ-358:** Controls shall not authorize canonical prohibitions.
- **RISK-REQ-359:** Controls shall not rely solely on documentation claims.
- **RISK-REQ-360:** Controls shall preserve least privilege and default deny.
- **RISK-REQ-361:** Controls shall preserve human accountability.

## 18. Control Effectiveness

- **RISK-REQ-362:** Control-effectiveness status `Not Implemented` shall mean no current effective control is available.
- **RISK-REQ-363:** `Not Implemented` shall identify supporting evidence or the reason evidence is absent.
- **RISK-REQ-364:** Control-effectiveness status `Implemented, Not Tested` shall mean the control exists but effectiveness evidence is absent.
- **RISK-REQ-365:** `Implemented, Not Tested` shall identify supporting evidence or the reason evidence is absent.
- **RISK-REQ-366:** Control-effectiveness status `Partially Effective` shall mean evidence shows limited effectiveness or material gaps.
- **RISK-REQ-367:** `Partially Effective` shall identify supporting evidence or the reason evidence is absent.
- **RISK-REQ-368:** Control-effectiveness status `Effective` shall mean evidence supports expected operation within the tested scope.
- **RISK-REQ-369:** `Effective` shall identify supporting evidence or the reason evidence is absent.
- **RISK-REQ-370:** `Effective` shall be reviewed when the control or environment changes.
- **RISK-REQ-371:** `Ineffective` shall identify supporting evidence or the reason evidence is absent.
- **RISK-REQ-372:** `Ineffective` shall be reviewed when the control or environment changes.
- **RISK-REQ-373:** `Unknown` shall identify supporting evidence or the reason evidence is absent.
- **RISK-REQ-374:** `Unknown` shall be reviewed when the control or environment changes.
- **RISK-REQ-375:** Residual-risk calculations shall credit only current controls.
- **RISK-REQ-376:** Untested controls shall receive limited or no effectiveness credit unless justified.
- **RISK-REQ-377:** Unknown effectiveness shall not be treated as effective.
- **RISK-REQ-378:** Control evidence shall remain traceable.
- **RISK-REQ-379:** Control-effectiveness review shall identify reviewer.
- **RISK-REQ-380:** Control-effectiveness automation shall remain advisory.
- **RISK-REQ-381:** Material post-test changes shall invalidate affected effectiveness evidence.

## 19. Residual Risk Assessment

- **RISK-REQ-382:** Residual risk shall be assessed after applying current credited controls.
- **RISK-REQ-383:** Residual likelihood shall be recorded.
- **RISK-REQ-384:** Residual impact shall be recorded.
- **RISK-REQ-385:** Residual score shall be calculated.
- **RISK-REQ-386:** Residual level shall be recorded.
- **RISK-REQ-387:** Residual rating confidence shall be recorded.
- **RISK-REQ-388:** Residual risk shall identify credited controls.
- **RISK-REQ-389:** Residual risk shall identify unmitigated causes.
- **RISK-REQ-390:** Residual risk shall identify control limitations.
- **RISK-REQ-391:** Residual risk shall identify remaining affected parties.
- **RISK-REQ-392:** Residual risk shall identify remaining evidence or privacy impacts.
- **RISK-REQ-393:** Residual risk shall identify release impact.
- **RISK-REQ-394:** Residual risk shall identify acceptance eligibility.
- **RISK-REQ-395:** Residual risk shall not be lower than evidence supports.
- **RISK-REQ-396:** Residual risk shall not credit planned treatments.
- **RISK-REQ-397:** Residual risk shall not be accepted by an AI system.
- **RISK-REQ-398:** Residual risk shall be recalculated after control change.
- **RISK-REQ-399:** Residual risk shall be recalculated after incident.
- **RISK-REQ-400:** Residual risk history shall remain visible.
- **RISK-REQ-401:** Residual risk closure shall require evidence.

## 20. Risk Treatment Options

- **RISK-REQ-402:** `Avoid` shall be a controlled treatment option meaning remove the activity, capability, data flow or condition creating the risk.
- **RISK-REQ-403:** Selection of `Avoid` shall identify owner, rationale and review date.
- **RISK-REQ-404:** `Avoid` shall not waive canonical prohibitions.
- **RISK-REQ-405:** `Reduce` shall be a controlled treatment option meaning implement controls that lower likelihood or impact.
- **RISK-REQ-406:** Selection of `Reduce` shall identify owner, rationale and review date.
- **RISK-REQ-407:** `Reduce` shall not waive canonical prohibitions.
- **RISK-REQ-408:** `Transfer` shall be a controlled treatment option meaning allocate defined financial or operational consequences through lawful arrangements without transferring accountability.
- **RISK-REQ-409:** Selection of `Transfer` shall identify owner, rationale and review date.
- **RISK-REQ-410:** `Transfer` shall not waive canonical prohibitions.
- **RISK-REQ-411:** `Accept` shall be a controlled treatment option meaning time-bounded human decision to retain residual risk within authority.
- **RISK-REQ-412:** Selection of `Accept` shall identify owner, rationale and review date.
- **RISK-REQ-413:** `Accept` shall not waive canonical prohibitions.
- **RISK-REQ-414:** `Research` shall be a controlled treatment option meaning perform governed research to reduce uncertainty before another treatment decision.
- **RISK-REQ-415:** Selection of `Research` shall identify owner, rationale and review date.
- **RISK-REQ-416:** `Research` shall not waive canonical prohibitions.
- **RISK-REQ-417:** `Monitor` shall be a controlled treatment option meaning observe an accepted or currently tolerable residual risk with defined indicators.
- **RISK-REQ-418:** Selection of `Monitor` shall identify owner, rationale and review date.
- **RISK-REQ-419:** `Monitor` shall not waive canonical prohibitions.

## 21. Treatment Plan

- **RISK-REQ-420:** Every High or Critical risk shall have a treatment plan.
- **RISK-REQ-421:** Every Moderate risk shall have treatment or acceptance rationale.
- **RISK-REQ-422:** Treatment plans shall have immutable identifiers.
- **RISK-REQ-423:** Treatment plans shall identify target risk.
- **RISK-REQ-424:** Treatment plans shall identify selected treatment option.
- **RISK-REQ-425:** Treatment plans shall identify actions.
- **RISK-REQ-426:** Treatment plans shall identify control requirements.
- **RISK-REQ-427:** Treatment plans shall identify owners.
- **RISK-REQ-428:** Treatment plans shall identify dependencies.
- **RISK-REQ-429:** Treatment plans shall identify milestones.
- **RISK-REQ-430:** Treatment plans shall identify target dates.
- **RISK-REQ-431:** Treatment plans shall identify resources.
- **RISK-REQ-432:** Treatment plans shall identify success criteria.
- **RISK-REQ-433:** Treatment plans shall identify validation plan.
- **RISK-REQ-434:** Treatment plans shall identify rollback or containment.
- **RISK-REQ-435:** Treatment plans shall identify expected residual risk.
- **RISK-REQ-436:** Treatment plans shall identify assumptions.
- **RISK-REQ-437:** Treatment plans shall identify research needs.
- **RISK-REQ-438:** Treatment plans shall identify release impact.
- **RISK-REQ-439:** Treatment plans shall identify closure criteria.
- **RISK-REQ-440:** Treatment-plan delays shall trigger review.
- **RISK-REQ-441:** Treatment-plan failure shall trigger reassessment.
- **RISK-REQ-442:** Treatment changes shall preserve history.
- **RISK-REQ-443:** Treatment completion shall require evidence.
- **RISK-REQ-444:** Treatment completion shall not automatically close the risk.

## 22. Risk Acceptance

- **RISK-REQ-445:** Risk acceptance shall be an explicit human decision.
- **RISK-REQ-446:** Risk acceptance shall identify the exact residual risk.
- **RISK-REQ-447:** Risk acceptance shall identify the acceptance authority.
- **RISK-REQ-448:** Risk acceptance shall identify rationale.
- **RISK-REQ-449:** Risk acceptance shall identify scope.
- **RISK-REQ-450:** Risk acceptance shall identify conditions.
- **RISK-REQ-451:** Risk acceptance shall identify compensating controls.
- **RISK-REQ-452:** Risk acceptance shall identify evidence.
- **RISK-REQ-453:** Risk acceptance shall identify review date.
- **RISK-REQ-454:** Risk acceptance shall identify expiry.
- **RISK-REQ-455:** Risk acceptance shall identify release impact.
- **RISK-REQ-456:** Risk acceptance shall identify revocation conditions.
- **RISK-REQ-457:** Low residual risk may be accepted by a delegated domain owner within documented authority.
- **RISK-REQ-458:** Moderate residual risk shall require the Privacy and Governance Reviewer and applicable domain reviewer.
- **RISK-REQ-459:** High residual risk shall require Project Founder approval after applicable specialist review.
- **RISK-REQ-460:** Critical residual risk shall not be accepted for an approved baseline, operational activation or public release.
- **RISK-REQ-461:** Canonical-prohibition risk shall not be accepted.
- **RISK-REQ-462:** Risk acceptance shall not create legal authority.
- **RISK-REQ-463:** Risk acceptance shall not establish compliance.
- **RISK-REQ-464:** Risk acceptance shall not replace implementation or testing.
- **RISK-REQ-465:** Expired acceptance shall return the risk to `Assessing`.
- **RISK-REQ-466:** Material change shall invalidate affected acceptance.
- **RISK-REQ-467:** Incident occurrence shall trigger acceptance review.
- **RISK-REQ-468:** Accepted risk shall remain visible.
- **RISK-REQ-469:** Acceptance withdrawal shall be attributable.

## 23. Risk Ownership

- **RISK-REQ-470:** Every risk owner shall be an identified human role or accountable institutional role.
- **RISK-REQ-471:** Risk ownership shall be accepted explicitly.
- **RISK-REQ-472:** Risk owners shall maintain record completeness.
- **RISK-REQ-473:** Risk owners shall coordinate assessment.
- **RISK-REQ-474:** Risk owners shall coordinate treatment.
- **RISK-REQ-475:** Risk owners shall monitor review dates.
- **RISK-REQ-476:** Risk owners shall escalate overdue treatment.
- **RISK-REQ-477:** Risk owners shall ensure evidence references remain current.
- **RISK-REQ-478:** Risk owners shall disclose conflicts of interest.
- **RISK-REQ-479:** Risk owners shall disclose role concentration.
- **RISK-REQ-480:** Risk owners shall not self-certify independent assurance.
- **RISK-REQ-481:** Risk owners shall not delegate final accountability to an AI system.
- **RISK-REQ-482:** Risk owners may assign contributors without transferring accountability silently.
- **RISK-REQ-483:** Owner changes shall preserve history.
- **RISK-REQ-484:** Departing owners shall transfer or escalate open risks.
- **RISK-REQ-485:** Orphaned risks shall be escalated.
- **RISK-REQ-486:** Critical risks shall have an executive accountable owner.
- **RISK-REQ-487:** Privacy risks shall include Privacy and Governance Reviewer involvement.
- **RISK-REQ-488:** Security risks shall include Security Reviewer involvement.
- **RISK-REQ-489:** Implementation risks shall include Implementation Reviewer involvement.
- **RISK-REQ-490:** Release risks shall include Release Reviewer involvement.
- **RISK-REQ-491:** Evidence risks shall include evidence-domain review.
- **RISK-REQ-492:** Research risks shall include Research Reviewer involvement where applicable.
- **RISK-REQ-493:** Ownership shall not imply acceptance authority beyond delegated governance.
- **RISK-REQ-494:** Ownership records shall minimize personal data.

## 24. Review Cadence and Triggers

- **RISK-REQ-495:** Applicable risks shall be reviewed scheduled review date.
- **RISK-REQ-496:** Applicable risks shall be reviewed before every release.
- **RISK-REQ-497:** Applicable risks shall be reviewed before baseline approval or freeze.
- **RISK-REQ-498:** Applicable risks shall be reviewed before operational activation.
- **RISK-REQ-499:** Applicable risks shall be reviewed after a material architecture change.
- **RISK-REQ-500:** Applicable risks shall be reviewed after a material model or provider change.
- **RISK-REQ-501:** Applicable risks shall be reviewed after a material policy change.
- **RISK-REQ-502:** Applicable risks shall be reviewed after a connector or endpoint change.
- **RISK-REQ-503:** Applicable risks shall be reviewed after a credential or key compromise.
- **RISK-REQ-504:** Applicable risks shall be reviewed after a security incident.
- **RISK-REQ-505:** Applicable risks shall be reviewed after a privacy incident.
- **RISK-REQ-506:** Applicable risks shall be reviewed after an evidence-integrity incident.
- **RISK-REQ-507:** Applicable risks shall be reviewed after a failed critical test.
- **RISK-REQ-508:** Applicable risks shall be reviewed after a control-effectiveness change.
- **RISK-REQ-509:** Applicable risks shall be reviewed after an assumption is invalidated.
- **RISK-REQ-510:** Applicable risks shall be reviewed after a legal or standards applicability change.
- **RISK-REQ-511:** Applicable risks shall be reviewed after a significant threat-intelligence change.
- **RISK-REQ-512:** Applicable risks shall be reviewed after risk acceptance expiry.
- **RISK-REQ-513:** Applicable risks shall be reviewed after treatment completion.
- **RISK-REQ-514:** Applicable risks shall be reviewed after an owner or accountable institution change.
- **RISK-REQ-515:** Critical risks shall be reviewed continuously or at the shortest practical interval.
- **RISK-REQ-516:** High risks shall be reviewed at least monthly while unresolved.
- **RISK-REQ-517:** Moderate risks shall be reviewed at least quarterly while active.
- **RISK-REQ-518:** Low risks shall be reviewed at least annually or when triggered.
- **RISK-REQ-519:** Review cadence may be shortened based on uncertainty.
- **RISK-REQ-520:** Missed review dates shall trigger escalation.
- **RISK-REQ-521:** Review completion shall update the next review date.
- **RISK-REQ-522:** Review shall identify exact evidence considered.
- **RISK-REQ-523:** Review shall identify rating changes.
- **RISK-REQ-524:** Review shall identify treatment changes.
- **RISK-REQ-525:** Review shall identify release impact.
- **RISK-REQ-526:** Review shall preserve prior assessments.

## 25. Release Risk Gate

- **RISK-REQ-527:** Every release candidate shall include an applicable risk-register review.
- **RISK-REQ-528:** The release review shall identify all open Critical risks.
- **RISK-REQ-529:** The release review shall identify all open High risks.
- **RISK-REQ-530:** The release review shall identify expired acceptances.
- **RISK-REQ-531:** The release review shall identify overdue treatments.
- **RISK-REQ-532:** The release review shall identify stale reviews.
- **RISK-REQ-533:** The release review shall identify unresolved security-test failures.
- **RISK-REQ-534:** The release review shall identify unresolved privacy findings.
- **RISK-REQ-535:** The release review shall identify unresolved evidence-integrity findings.
- **RISK-REQ-536:** The release review shall identify unresolved connector risks.
- **RISK-REQ-537:** The release review shall identify unresolved supply-chain risks.
- **RISK-REQ-538:** The release review shall identify public-claim risks.
- **RISK-REQ-539:** Critical residual risk shall block release.
- **RISK-REQ-540:** High residual risk shall block release unless explicitly accepted within authority and permitted by this policy.
- **RISK-REQ-541:** Canonical-prohibition risk shall block release.
- **RISK-REQ-542:** Missing risk evidence shall block affected release claims.
- **RISK-REQ-543:** Release Reviewer shall record the risk-gate outcome.
- **RISK-REQ-544:** Risk-gate pass shall not replace release approval.
- **RISK-REQ-545:** Release withdrawal shall trigger risk review.
- **RISK-REQ-546:** Correction release shall review the originating risk.
- **RISK-REQ-547:** Public release shall not expose restricted risk details unnecessarily.
- **RISK-REQ-548:** Release records shall link to the exact register version.
- **RISK-REQ-549:** Release review shall preserve open-risk limitations.
- **RISK-REQ-550:** Every release shall satisfy the legacy requirement for risk review before release.
- **RISK-REQ-551:** Release risk history shall remain retrievable.

## 26. Escalation

- **RISK-REQ-552:** Critical risk identification shall trigger immediate escalation.
- **RISK-REQ-553:** High risk with overdue treatment shall trigger escalation.
- **RISK-REQ-554:** Expired High-risk acceptance shall trigger escalation.
- **RISK-REQ-555:** Risk owner absence shall trigger escalation.
- **RISK-REQ-556:** Risk rating dispute shall trigger escalation.
- **RISK-REQ-557:** Control failure affecting multiple risks shall trigger escalation.
- **RISK-REQ-558:** Evidence-integrity uncertainty affecting legal or governance decisions shall trigger escalation.
- **RISK-REQ-559:** Potential canonical-prohibition exposure shall trigger escalation.
- **RISK-REQ-560:** Potential unlawful surveillance or deanonymization shall trigger escalation.
- **RISK-REQ-561:** Potential uncontrolled criminal-infrastructure interaction shall trigger escalation.
- **RISK-REQ-562:** Potential credential transfer or compromise shall trigger escalation.
- **RISK-REQ-563:** Potential autonomous legal or coercive action shall trigger escalation.
- **RISK-REQ-564:** Escalation shall identify recipient authority.
- **RISK-REQ-565:** Escalation shall identify reason.
- **RISK-REQ-566:** Escalation shall identify time.
- **RISK-REQ-567:** Escalation shall identify containment.
- **RISK-REQ-568:** Escalation shall preserve evidence.
- **RISK-REQ-569:** Escalation shall not publish sensitive details unnecessarily.
- **RISK-REQ-570:** Escalation shall not wait for complete certainty when immediate containment is required.
- **RISK-REQ-571:** Escalation closure shall identify disposition.
- **RISK-REQ-572:** Unresolved escalations shall remain visible.
- **RISK-REQ-573:** Escalation failure shall trigger governance review.
- **RISK-REQ-574:** AI systems may recommend escalation but shall not close it.
- **RISK-REQ-575:** Emergency containment shall not authorize offensive action.
- **RISK-REQ-576:** Escalation records shall remain retrievable.

## 27. Critical and Non-Acceptable Risks

- **RISK-REQ-577:** Risk involving independent AI legal authority shall be treated as non-acceptable.
- **RISK-REQ-578:** Risk involving autonomous policing decisions shall be treated as non-acceptable.
- **RISK-REQ-579:** Risk involving autonomous coercive action shall be treated as non-acceptable.
- **RISK-REQ-580:** Risk involving unauthorized access shall be treated as non-acceptable.
- **RISK-REQ-581:** Risk involving credential theft shall be treated as non-acceptable.
- **RISK-REQ-582:** Risk involving malware deployment shall be treated as non-acceptable.
- **RISK-REQ-583:** Risk involving offensive cyber operations shall be treated as non-acceptable.
- **RISK-REQ-584:** Risk involving unlawful surveillance shall be treated as non-acceptable.
- **RISK-REQ-585:** Risk involving unlawful deanonymization shall be treated as non-acceptable.
- **RISK-REQ-586:** Risk involving uncontrolled criminal-infrastructure interaction shall be treated as non-acceptable.
- **RISK-REQ-587:** Risk involving human credential transfer to an agent shall be treated as non-acceptable.
- **RISK-REQ-588:** Risk involving silent destruction or alteration of original evidence shall be treated as non-acceptable.
- **RISK-REQ-589:** Risk involving unbounded cross-case data reuse shall be treated as non-acceptable.
- **RISK-REQ-590:** Risk involving release of real case data in public demonstrations shall be treated as non-acceptable.
- **RISK-REQ-591:** Non-acceptable risk shall be avoided, removed or contained.
- **RISK-REQ-592:** Non-acceptable risk shall not receive acceptance approval.
- **RISK-REQ-593:** Non-acceptable risk shall block affected implementation and release.
- **RISK-REQ-594:** Non-acceptable risk shall remain traceable to canonical authority.
- **RISK-REQ-595:** Attempted exception for non-acceptable risk shall be rejected.
- **RISK-REQ-596:** Closure shall require evidence that the prohibited condition is absent or inaccessible.

## 28. Identity, Delegation and Authorization Risk

- **RISK-REQ-597:** Risk assessments shall consider officer-agent binding failure.
- **RISK-REQ-598:** Risk assessments shall consider shared or orphaned identities.
- **RISK-REQ-599:** Risk assessments shall consider human credential transfer.
- **RISK-REQ-600:** Risk assessments shall consider stale or revoked delegation.
- **RISK-REQ-601:** Risk assessments shall consider policy-service failure.
- **RISK-REQ-602:** Risk assessments shall consider cached-decision staleness.
- **RISK-REQ-603:** Risk assessments shall consider revocation propagation latency.
- **RISK-REQ-604:** Risk assessments shall consider cross-case access.
- **RISK-REQ-605:** Risk assessments shall consider unapproved tool or connector use.
- **RISK-REQ-606:** Identity risks shall identify affected credentials and sessions.
- **RISK-REQ-607:** Authorization risks shall identify affected actions and resources.
- **RISK-REQ-608:** Authorization risk tests shall include denied, expired and revoked paths.

## 29. AI Model, Prompt and Human-Oversight Risk

- **RISK-REQ-609:** Risk assessments shall consider hallucination and fabricated claims.
- **RISK-REQ-610:** Risk assessments shall consider direct prompt injection.
- **RISK-REQ-611:** Risk assessments shall consider indirect prompt injection.
- **RISK-REQ-612:** Risk assessments shall consider model overconfidence.
- **RISK-REQ-613:** Risk assessments shall consider provider-side silent change.
- **RISK-REQ-614:** Risk assessments shall consider context-window loss.
- **RISK-REQ-615:** Risk assessments shall consider cross-case memory leakage.
- **RISK-REQ-616:** Risk assessments shall consider training-data or provider-retention exposure.
- **RISK-REQ-617:** Risk assessments shall consider inadequate human review.
- **RISK-REQ-618:** Model risks shall identify model, provider, prompt and policy versions.
- **RISK-REQ-619:** AI output shall remain untrusted and derived.
- **RISK-REQ-620:** AI risk treatment shall not rely solely on the model detecting its own failure.

## 30. Evidence and Forensic Risk

- **RISK-REQ-621:** Risk assessments shall consider source and analysis conflation.
- **RISK-REQ-622:** Risk assessments shall consider integrity mismatch.
- **RISK-REQ-623:** Risk assessments shall consider chain-of-custody gaps.
- **RISK-REQ-624:** Risk assessments shall consider timestamp uncertainty.
- **RISK-REQ-625:** Risk assessments shall consider unsupported admissibility claims.
- **RISK-REQ-626:** Risk assessments shall consider export-manifest errors.
- **RISK-REQ-627:** Risk assessments shall consider retention and legal-hold failure.
- **RISK-REQ-628:** Risk assessments shall consider AI-generated content being stored as original evidence.
- **RISK-REQ-629:** Risk assessments shall consider backup and restoration corruption.
- **RISK-REQ-630:** Evidence risks shall identify affected objects and custody events.
- **RISK-REQ-631:** Evidence control tests shall include alteration and substitution.
- **RISK-REQ-632:** Evidence integrity risk shall not be accepted when it undermines authoritative evidence.

## 31. Privacy and Fundamental-Rights Risk

- **RISK-REQ-633:** Risk assessments shall consider overcollection.
- **RISK-REQ-634:** Risk assessments shall consider excessive retention.
- **RISK-REQ-635:** Risk assessments shall consider unauthorized disclosure.
- **RISK-REQ-636:** Risk assessments shall consider sensitive personal-data processing.
- **RISK-REQ-637:** Risk assessments shall consider unlawful surveillance.
- **RISK-REQ-638:** Risk assessments shall consider chilling effects.
- **RISK-REQ-639:** Risk assessments shall consider absence of meaningful human review.
- **RISK-REQ-640:** Risk assessments shall consider public disclosure of real persons.
- **RISK-REQ-641:** Risk assessments shall consider telemetry overcollection.
- **RISK-REQ-642:** Privacy risks shall identify purpose and legal-authority assumptions without asserting legal conclusions.
- **RISK-REQ-643:** Rights risks shall receive Privacy and Governance Reviewer assessment.
- **RISK-REQ-644:** Privacy risk evidence shall minimize further personal-data exposure.

## 32. Connector and External-Service Risk

- **RISK-REQ-645:** Risk assessments shall consider connector credential compromise.
- **RISK-REQ-646:** Risk assessments shall consider unapproved endpoints.
- **RISK-REQ-647:** Risk assessments shall consider server-side request forgery.
- **RISK-REQ-648:** Risk assessments shall consider prompt injection through retrieved content.
- **RISK-REQ-649:** Risk assessments shall consider rate and quota exhaustion.
- **RISK-REQ-650:** Risk assessments shall consider retry storms.
- **RISK-REQ-651:** Risk assessments shall consider external-service outage.
- **RISK-REQ-652:** Risk assessments shall consider cross-tenant or cross-account access.
- **RISK-REQ-653:** Risk assessments shall consider uncontrolled Dark Web interaction.
- **RISK-REQ-654:** Connector risks shall identify machine identities and endpoints.
- **RISK-REQ-655:** Connector treatment shall preserve controlled egress.
- **RISK-REQ-656:** Connector risk closure shall verify credential, session and endpoint revocation.

## 33. Software, Supply-Chain and Repository Risk

- **RISK-REQ-657:** Risk assessments shall consider insecure code defects.
- **RISK-REQ-658:** Risk assessments shall consider dependency confusion.
- **RISK-REQ-659:** Risk assessments shall consider typosquatting.
- **RISK-REQ-660:** Risk assessments shall consider compromised CI workflows.
- **RISK-REQ-661:** Risk assessments shall consider secret exposure in repository history.
- **RISK-REQ-662:** Risk assessments shall consider branch-protection gaps.
- **RISK-REQ-663:** Risk assessments shall consider release artifact substitution.
- **RISK-REQ-664:** Risk assessments shall consider AI-assisted code hallucinations.
- **RISK-REQ-665:** Risk assessments shall consider missing dependency provenance.
- **RISK-REQ-666:** Repository risks shall identify exact commits and settings evidence.
- **RISK-REQ-667:** Critical workflow risks shall block release.
- **RISK-REQ-668:** Supply-chain risk closure shall include regression and integrity evidence.

## 34. Operational Resilience and Availability Risk

- **RISK-REQ-669:** Risk assessments shall consider identity-service outage.
- **RISK-REQ-670:** Risk assessments shall consider audit-service outage.
- **RISK-REQ-671:** Risk assessments shall consider monitoring-service outage.
- **RISK-REQ-672:** Risk assessments shall consider network partition.
- **RISK-REQ-673:** Risk assessments shall consider partial activation.
- **RISK-REQ-674:** Risk assessments shall consider partial revocation.
- **RISK-REQ-675:** Risk assessments shall consider stale replicas.
- **RISK-REQ-676:** Risk assessments shall consider resource exhaustion.
- **RISK-REQ-677:** Risk assessments shall consider recovery restoring revoked authority.
- **RISK-REQ-678:** Resilience treatments shall identify recovery time and data-integrity needs where applicable.
- **RISK-REQ-679:** Resilience tests shall use controlled fault injection.
- **RISK-REQ-680:** Availability risk shall not justify bypassing authorization or evidence controls.

## 35. Publication, Portfolio and Compliance-Claim Risk

- **RISK-REQ-681:** Risk assessments shall consider proposed architecture being represented as implemented.
- **RISK-REQ-682:** Risk assessments shall consider internal review being represented as independent assurance.
- **RISK-REQ-683:** Risk assessments shall consider version numbers being represented as approval.
- **RISK-REQ-684:** Risk assessments shall consider release of secrets.
- **RISK-REQ-685:** Risk assessments shall consider release of sensitive endpoints.
- **RISK-REQ-686:** Risk assessments shall consider outdated public documentation.
- **RISK-REQ-687:** Risk assessments shall consider unsupported production-readiness claims.
- **RISK-REQ-688:** Risk assessments shall consider misleading institutional-adoption claims.
- **RISK-REQ-689:** Risk assessments shall consider withdrawn artifacts remaining discoverable without warning.
- **RISK-REQ-690:** Release review shall consider every open publication risk.
- **RISK-REQ-691:** Public limitations shall remain visible.
- **RISK-REQ-692:** Publication risk closure shall verify all distributed copies within project control.

## 36. Assumptions, Research and Decision Integration

- **RISK-REQ-693:** Every risk dependent on an assumption shall link the assumption identifier.
- **RISK-REQ-694:** Assumption invalidation shall trigger risk review.
- **RISK-REQ-695:** Assumption expiry shall trigger risk review.
- **RISK-REQ-696:** Research backlog priority shall consider linked risk level.
- **RISK-REQ-697:** Research closure shall provide evidence or an explicit inconclusive result.
- **RISK-REQ-698:** Risk acceptance requiring governance decision shall link the decision record.
- **RISK-REQ-699:** Risk records shall not substitute for ADRs.
- **RISK-REQ-700:** Decision logs shall index material risk decisions where applicable.
- **RISK-REQ-701:** Risk findings shall update tests when control verification changes.
- **RISK-REQ-702:** Risk findings shall update compliance mappings where design evidence changes.
- **RISK-REQ-703:** Forward document 28 shall govern assumption-record mechanics.
- **RISK-REQ-704:** Forward document 30 shall govern compliance-mapping mechanics.
- **RISK-REQ-705:** Research uncertainty shall not default to acceptance.
- **RISK-REQ-706:** Assumption-dependent acceptance shall expire no later than the assumption review date.
- **RISK-REQ-707:** Cross-record links shall be bidirectional.

## 37. Testing and Validation

- **RISK-REQ-708:** Every High and Critical risk shall have linked validation activities.
- **RISK-REQ-709:** Risk tests shall identify exact requirements.
- **RISK-REQ-710:** Risk tests shall identify exact control versions.
- **RISK-REQ-711:** Risk tests shall identify exact environment and fixtures.
- **RISK-REQ-712:** Risk tests shall include observed results.
- **RISK-REQ-713:** Risk tests shall include failure and recovery paths where applicable.
- **RISK-REQ-714:** Risk tests shall include revocation paths where applicable.
- **RISK-REQ-715:** Risk tests shall use synthetic or authorized data.
- **RISK-REQ-716:** Failed risk tests shall trigger reassessment.
- **RISK-REQ-717:** Skipped mandatory risk tests shall remain visible.
- **RISK-REQ-718:** Passing tests shall not prove absence of risk.
- **RISK-REQ-719:** Material changes shall invalidate affected risk evidence.
- **RISK-REQ-720:** Test evidence shall preserve integrity.
- **RISK-REQ-721:** AI-generated test summaries shall not replace raw evidence.
- **RISK-REQ-722:** Testing shall conform to `OBDIA-TEST-001`.

## 38. Monitoring and Key Risk Indicators

- **RISK-REQ-723:** Active High and Critical risks shall define monitoring indicators.
- **RISK-REQ-724:** Indicators shall have data sources.
- **RISK-REQ-725:** Indicators shall have thresholds.
- **RISK-REQ-726:** Indicators shall have alert actions.
- **RISK-REQ-727:** Indicators shall not expose secrets.
- **RISK-REQ-728:** Indicators shall distinguish leading and lagging signals where useful.
- **RISK-REQ-729:** Indicators shall identify baseline values.
- **RISK-REQ-730:** Threshold changes shall be attributable.
- **RISK-REQ-731:** Monitoring failure shall trigger risk review.
- **RISK-REQ-732:** Missing indicator data shall not be treated as absence of risk.
- **RISK-REQ-733:** Risk dashboards shall identify data freshness.
- **RISK-REQ-734:** AI-generated monitoring analysis shall remain advisory.
- **RISK-REQ-735:** Indicator breaches shall create review events.
- **RISK-REQ-736:** Repeated breaches shall trigger escalation.
- **RISK-REQ-737:** Monitoring limitations shall remain visible.

## 39. Incidents and Emerging Risks

- **RISK-REQ-738:** Incidents shall link to affected risks.
- **RISK-REQ-739:** Incident occurrence shall trigger inherent and residual reassessment.
- **RISK-REQ-740:** Incident evidence shall remain distinct from risk summaries.
- **RISK-REQ-741:** Incident remediation shall update control-effectiveness evidence.
- **RISK-REQ-742:** Near misses shall be considered.
- **RISK-REQ-743:** Emerging risks shall be recorded before complete certainty when credible.
- **RISK-REQ-744:** Emerging risks shall identify uncertainty.
- **RISK-REQ-745:** Threat-intelligence changes shall be evaluated for applicability.
- **RISK-REQ-746:** Incident severity and risk level shall remain distinct.
- **RISK-REQ-747:** Security incidents shall trigger Security Reviewer involvement.
- **RISK-REQ-748:** Evidence incidents shall trigger evidence-domain review.
- **RISK-REQ-749:** Public incident claims shall follow release policy.
- **RISK-REQ-750:** Incident evidence shall not contain unnecessary exploit detail.
- **RISK-REQ-751:** Resolved incidents shall generate regression tests where applicable.
- **RISK-REQ-752:** Emerging-risk closure shall identify why the scenario is no longer credible.

## 40. Exceptions and Compensating Controls

- **RISK-REQ-753:** Risk exceptions shall be explicit.
- **RISK-REQ-754:** Risk exceptions shall identify rationale.
- **RISK-REQ-755:** Risk exceptions shall identify owner.
- **RISK-REQ-756:** Risk exceptions shall identify residual risk.
- **RISK-REQ-757:** Risk exceptions shall identify expiry.
- **RISK-REQ-758:** Risk exceptions shall identify release impact.
- **RISK-REQ-759:** Risk exceptions shall not waive canonical prohibitions.
- **RISK-REQ-760:** Risk exceptions shall not fabricate control effectiveness.
- **RISK-REQ-761:** Risk exceptions shall not transfer accountability to an AI system.
- **RISK-REQ-762:** Expired exceptions shall trigger reassessment.
- **RISK-REQ-763:** Exception closure shall preserve history.
- **RISK-REQ-764:** Compensating controls shall be tested.
- **RISK-REQ-765:** Exception status shall remain distinct from risk status.
- **RISK-REQ-766:** Missing detailed exception policy shall not be used to bypass this section.
- **RISK-REQ-767:** Exception records shall remain retrievable.

## 41. Risk Closure and Reopening

- **RISK-REQ-768:** Risk closure shall identify closure rationale.
- **RISK-REQ-769:** Risk closure shall identify closure date.
- **RISK-REQ-770:** Risk closure shall identify closure evidence.
- **RISK-REQ-771:** Risk closure shall identify control disposition.
- **RISK-REQ-772:** Risk closure shall identify linked assumption disposition.
- **RISK-REQ-773:** Risk closure shall identify release implications.
- **RISK-REQ-774:** Risk closure shall not erase prior acceptance.
- **RISK-REQ-775:** Risk closure shall not be based solely on absence of recent incidents.
- **RISK-REQ-776:** Risk closure shall not occur while critical evidence is missing.
- **RISK-REQ-777:** Risk closure may use eliminated, no longer applicable or duplicate/superseded rationale.
- **RISK-REQ-778:** Superseded closure shall identify successor risks.
- **RISK-REQ-779:** Closed risks shall remain in release history where previously relevant.
- **RISK-REQ-780:** Reopening shall identify trigger.
- **RISK-REQ-781:** Reopening shall preserve the old closure record.
- **RISK-REQ-782:** Reopened risk shall return to `Assessing`.

## 42. Versioning and Immutable History

- **RISK-REQ-783:** The risk-register method shall have a semantic version.
- **RISK-REQ-784:** Risk-record schema changes shall have versions.
- **RISK-REQ-785:** Risk-register baseline snapshots shall have versions.
- **RISK-REQ-786:** Risk-record content changes shall preserve revision history.
- **RISK-REQ-787:** Owner changes shall preserve prior owners.
- **RISK-REQ-788:** Control changes shall preserve prior control references.
- **RISK-REQ-789:** Acceptance changes shall preserve prior decisions.
- **RISK-REQ-790:** Historical risk snapshots shall not be silently rewritten.
- **RISK-REQ-791:** Baseline manifests shall identify the exact risk-register version.
- **RISK-REQ-792:** Release records shall identify the exact risk-register version reviewed.
- **RISK-REQ-793:** Markdown and machine-readable records shall remain synchronized.
- **RISK-REQ-794:** Risk-register tags shall not create acceptance.
- **RISK-REQ-795:** Version `1.0.0` shall not imply approval.
- **RISK-REQ-796:** Risk history shall conform to `OBDIA-VER-001`.
- **RISK-REQ-797:** Version history shall remain retrievable.

## 43. Automation and Tooling

- **RISK-REQ-798:** Automation may validate risk-record schema.
- **RISK-REQ-799:** Automation may detect overdue reviews.
- **RISK-REQ-800:** Automation may detect expired acceptance.
- **RISK-REQ-801:** Automation may validate cross-references.
- **RISK-REQ-802:** Automation may propose duplicate risks.
- **RISK-REQ-803:** Automation shall not assign final qualitative impact autonomously.
- **RISK-REQ-804:** Automation shall not assign final likelihood autonomously.
- **RISK-REQ-805:** Automation shall not close risk.
- **RISK-REQ-806:** Automation shall not alter historical records silently.
- **RISK-REQ-807:** Automation configuration shall be versioned.
- **RISK-REQ-808:** Automation failure shall remain visible.
- **RISK-REQ-809:** AI-generated risk narratives shall receive human review.
- **RISK-REQ-810:** Risk tools shall use least privilege.
- **RISK-REQ-811:** Risk tools shall not receive secrets unnecessarily.
- **RISK-REQ-812:** Automation limitations shall remain visible.

## 44. Reporting and Metrics

- **RISK-REQ-813:** Risk reporting shall identify open risks by level.
- **RISK-REQ-814:** Risk reporting shall identify overdue treatments.
- **RISK-REQ-815:** Risk reporting shall identify expired acceptances.
- **RISK-REQ-816:** Risk reporting shall identify control-effectiveness gaps.
- **RISK-REQ-817:** Risk reporting shall identify trend methodology.
- **RISK-REQ-818:** Average risk score shall not conceal Critical risks.
- **RISK-REQ-819:** Closed-risk counts shall not incentivize premature closure.
- **RISK-REQ-820:** Metrics shall identify the register version.
- **RISK-REQ-821:** Metrics shall identify excluded records.
- **RISK-REQ-822:** Metrics shall distinguish inherent and residual risk.
- **RISK-REQ-823:** Metrics shall distinguish planned and implemented controls.
- **RISK-REQ-824:** Metrics shall minimize personal and case data.
- **RISK-REQ-825:** Public reporting shall not claim compliance.
- **RISK-REQ-826:** AI-generated reports shall receive human verification.
- **RISK-REQ-827:** Reporting limitations shall remain visible.

## 45. Public and Portfolio Boundaries

- **RISK-REQ-828:** Public risk summaries shall use synthetic or abstract scenarios.
- **RISK-REQ-829:** Public risk summaries shall not identify real officers.
- **RISK-REQ-830:** Public risk summaries shall not identify real case subjects.
- **RISK-REQ-831:** Public risk summaries shall not expose internal credentials.
- **RISK-REQ-832:** Public risk summaries shall not expose sensitive endpoints.
- **RISK-REQ-833:** Public risk summaries shall not provide unnecessary operational attack instructions.
- **RISK-REQ-834:** Public risk summaries shall distinguish planned from implemented controls.
- **RISK-REQ-835:** Public risk summaries shall distinguish implemented from validated controls.
- **RISK-REQ-836:** Public risk summaries shall distinguish internal review from independent assurance.
- **RISK-REQ-837:** Public risk summaries shall not claim zero risk.
- **RISK-REQ-838:** Public risk summaries shall not claim legal compliance.
- **RISK-REQ-839:** Public risk summaries shall not imply institutional adoption.
- **RISK-REQ-840:** Public risk summaries shall state limitations.
- **RISK-REQ-841:** Public release shall review open publication risks.
- **RISK-REQ-842:** Redaction shall preserve the substantive risk meaning.
- **RISK-REQ-843:** Redaction shall be documented.
- **RISK-REQ-844:** Withdrawal shall preserve records.
- **RISK-REQ-845:** Correction shall use governed versioning.
- **RISK-REQ-846:** Portfolio claims shall trace to exact evidence.
- **RISK-REQ-847:** Risk-register publication shall require Release Reviewer assessment.

## 46. Traceability

- **RISK-REQ-848:** Every risk shall trace to its source.
- **RISK-REQ-849:** Every risk shall trace to affected requirements.
- **RISK-REQ-850:** Every risk shall trace to relevant threats.
- **RISK-REQ-851:** Every risk shall trace to affected assets.
- **RISK-REQ-852:** Every risk shall trace to affected parties.
- **RISK-REQ-853:** Every risk shall trace to current controls.
- **RISK-REQ-854:** Every risk shall trace to planned treatments.
- **RISK-REQ-855:** Every risk shall trace to control-effectiveness evidence.
- **RISK-REQ-856:** Every risk shall trace to applicable tests.
- **RISK-REQ-857:** Every risk shall trace to incidents.
- **RISK-REQ-858:** Every risk shall trace to assumptions.
- **RISK-REQ-859:** Every risk shall trace to research items.
- **RISK-REQ-860:** Every risk shall trace to ADRs and decisions where applicable.
- **RISK-REQ-861:** Every risk shall trace to acceptance decisions.
- **RISK-REQ-862:** Every risk shall trace to release records where applicable.
- **RISK-REQ-863:** Every risk shall trace to closure evidence.
- **RISK-REQ-864:** Every risk shall trace to revision history.
- **RISK-REQ-865:** Every control shall trace back to affected risks.
- **RISK-REQ-866:** Every test shall trace back to affected risks.
- **RISK-REQ-867:** Every acceptance shall trace back to the exact residual rating.
- **RISK-REQ-868:** Traceability shall be bidirectional.
- **RISK-REQ-869:** Broken traceability affecting Critical or High risk shall be blocking.
- **RISK-REQ-870:** Traceability records shall not contain secrets.
- **RISK-REQ-871:** Traceability records shall minimize personal and case data.
- **RISK-REQ-872:** Baseline freeze shall validate risk traceability across documents 01–30.

## 47. Review Gate

- **RISK-REQ-873:** AI Risk Register review shall identify the exact document and commit.
- **RISK-REQ-874:** Review shall verify review-before-release requirements.
- **RISK-REQ-875:** Review shall verify identifier and schema rules.
- **RISK-REQ-876:** Review shall verify lifecycle states.
- **RISK-REQ-877:** Review shall verify likelihood and impact definitions.
- **RISK-REQ-878:** Review shall verify control-effectiveness rules.
- **RISK-REQ-879:** Review shall verify treatment and acceptance authority.
- **RISK-REQ-880:** Review shall verify review triggers and release gating.
- **RISK-REQ-881:** Review shall verify evidence and authorization risks.
- **RISK-REQ-882:** Review shall verify AI, connector and supply-chain risks.
- **RISK-REQ-883:** Review shall verify testing and monitoring.
- **RISK-REQ-884:** Review shall verify immutable history.
- **RISK-REQ-885:** Review shall verify traceability.
- **RISK-REQ-886:** Review shall identify residual limitations.
- **RISK-REQ-887:** Critical findings shall block progression.
- **RISK-REQ-888:** Role concentration shall be disclosed.
- **RISK-REQ-889:** Internal review shall not be represented as independent external assurance.
- **RISK-REQ-890:** Security Reviewer shall assess security-risk coverage.
- **RISK-REQ-891:** Implementation Reviewer shall assess implementation and supply-chain coverage.
- **RISK-REQ-892:** Project Founder approval shall not substitute for required specialist review.
- **RISK-REQ-893:** Automated findings shall remain advisory.
- **RISK-REQ-894:** Approval for Draft incorporation shall not accept any individual risk.

## 48. Initial Baseline Risk Catalogue

The following entries are an initial synthetic research baseline. Ratings are provisional and shall be reassessed when implementations, environments, controls or evidence exist.

| Risk ID | Title | Category | Inherent scenario | Initial L | Initial I | Initial level | Initial treatment direction |
|---|---|---|---|---:|---:|---|---|
| `AI-RISK-001` | Independent AI authority or autonomous consequential action | Authority and Accountability | The agent is treated as having legal, policing or coercive authority independent of the officer. | 2 | 5 | Critical | Avoid; enforce human approval and explicit non-authority boundaries. |
| `AI-RISK-002` | Officer-agent binding failure | Identity and Delegation | The agent identity cannot be reliably bound to exactly one authorized officer and institution. | 3 | 4 | High | Reduce; cryptographic binding, verification and revocation. |
| `AI-RISK-003` | Human credential transfer to agent | Identity and Delegation | Human passwords, tokens or reusable credentials are embedded in agent or connector workflows. | 2 | 5 | Critical | Avoid; dedicated machine identities and secret controls. |
| `AI-RISK-004` | Authorization bypass or permissive fallback | Authorization and Access Control | Missing or failed policy evaluation permits a protected action. | 3 | 5 | Critical | Reduce; default deny, complete mediation and failure tests. |
| `AI-RISK-005` | Stale delegation, policy or revocation state | Authorization and Access Control | Cached or replicated state permits operation after authority changed or expired. | 3 | 4 | High | Reduce; bounded caches, re-evaluation and tested propagation. |
| `AI-RISK-006` | Direct prompt injection | AI Model and Prompt | User input manipulates the model into ignoring policy or requesting unsafe actions. | 4 | 4 | High | Reduce; constrained tool schemas, policy enforcement and adversarial tests. |
| `AI-RISK-007` | Indirect prompt injection from connector content | AI Model and Prompt | Retrieved content contains instructions that influence agent behavior or tool use. | 4 | 5 | Critical | Reduce; content/control separation, untrusted labeling and tool authorization. |
| `AI-RISK-008` | Hallucination or fabricated investigative claim | AI Model and Prompt | The model produces unsupported facts, citations, entities or conclusions. | 4 | 4 | High | Reduce; provenance, source checking, uncertainty and human review. |
| `AI-RISK-009` | Automation bias and overreliance | Human Factors | Investigators defer to AI output despite uncertainty or conflicting evidence. | 3 | 4 | High | Reduce; training, interfaces, mandatory review and challenge paths. |
| `AI-RISK-010` | Inadequate human oversight | Human Factors | Human approval becomes superficial, delayed or unable to detect material errors. | 3 | 4 | High | Reduce; defined approval criteria, evidence and separation of duties. |
| `AI-RISK-011` | Original evidence and AI output conflation | Evidence and Forensics | Derived model output is stored or presented as original evidence. | 3 | 5 | Critical | Avoid; separate identifiers, stores, labels and lineage. |
| `AI-RISK-012` | Evidence integrity or chain-of-custody loss | Evidence and Forensics | Objects are altered, substituted or transferred without verifiable custody. | 3 | 5 | Critical | Reduce; immutable events, hashes, custody validation and quarantine. |
| `AI-RISK-013` | Unsupported authenticity or admissibility claim | Evidence and Forensics | Integrity or provenance is overstated as authenticity, legality or admissibility. | 3 | 4 | High | Reduce; claim boundaries, legal review and evidence-model controls. |
| `AI-RISK-014` | Cross-case data contamination | Data and Privacy | Data, memory, outputs or credentials cross between unrelated cases. | 3 | 5 | Critical | Reduce; case isolation, scoped storage, authorization and tests. |
| `AI-RISK-015` | Excessive or incompatible personal-data processing | Data and Privacy | The system collects, retains or reuses more personal data than justified. | 4 | 4 | High | Reduce; minimization, purpose controls, retention and privacy review. |
| `AI-RISK-016` | Unlawful surveillance or deanonymization | Fundamental Rights | Capabilities are used to monitor or identify people without lawful authority. | 2 | 5 | Critical | Avoid; prohibit capability and require rights review. |
| `AI-RISK-017` | Bias or discriminatory impact | Fundamental Rights | Model or workflow disadvantages protected or vulnerable groups. | 3 | 4 | High | Reduce; rights-impact analysis, representative tests and human review. |
| `AI-RISK-018` | Provider retention or secondary use of submitted data | Data and Privacy | External model or service providers retain or reuse sensitive inputs. | 3 | 4 | High | Reduce; contractual/technical controls, minimization and approved providers. |
| `AI-RISK-019` | Secret or credential exposure | Software and Supply Chain | Secrets appear in code, prompts, logs, fixtures, workflows or history. | 3 | 5 | Critical | Reduce; scanning, least privilege, rotation and incident response. |
| `AI-RISK-020` | Connector privilege escalation | Connector and External Service | A connector obtains broader permissions, endpoints or actions than authorized. | 3 | 5 | Critical | Reduce; manifests, least privilege, endpoint control and revocation. |
| `AI-RISK-021` | Malicious or malformed connector response | Connector and External Service | External content exploits parsers, model context or tool logic. | 4 | 4 | High | Reduce; schemas, bounds, isolation, quarantine and negative tests. |
| `AI-RISK-022` | Uncontrolled criminal-infrastructure interaction | Connector and External Service | Research tooling communicates with live criminal infrastructure outside an approved isolated design. | 2 | 5 | Critical | Avoid; simulations, archives, test services and egress denial. |
| `AI-RISK-023` | Dark Web isolation failure | Connector and External Service | Operationally isolated research loses network, identity or data separation. | 2 | 5 | Critical | Avoid or reduce; strict isolation, mocks, monitoring and shutdown. |
| `AI-RISK-024` | Dependency or build supply-chain compromise | Software and Supply Chain | A package, action, build tool or artifact is malicious or substituted. | 3 | 5 | Critical | Reduce; provenance, pinning, review, SBOM and integrity verification. |
| `AI-RISK-025` | Repository governance or workflow compromise | Repository and Governance | Protected documents, CI workflows or releases are modified without valid review. | 3 | 4 | High | Reduce; branch protection, CODEOWNERS, least privilege and audit. |
| `AI-RISK-026` | Model or provider drift | AI Model and Prompt | Model behavior changes without corresponding risk, test or policy review. | 4 | 4 | High | Reduce; version capture, change triggers and revalidation. |
| `AI-RISK-027` | Policy or schema version incompatibility | Repository and Governance | Mixed versions cause incorrect authorization, evidence or lifecycle behavior. | 3 | 4 | High | Reduce; manifests, compatibility tests and baseline controls. |
| `AI-RISK-028` | Audit or monitoring failure | Operational Resilience | Security-relevant actions occur without complete audit or detection. | 3 | 4 | High | Reduce; fail-closed behavior, health monitoring and evidence preservation. |
| `AI-RISK-029` | Revocation propagation failure | Operational Resilience | Credentials, sessions, workloads or connectors remain usable after revocation. | 3 | 5 | Critical | Reduce; tested propagation, containment and bounded latency. |
| `AI-RISK-030` | Availability or external-service outage | Operational Resilience | Critical dependencies fail and disrupt authorized work or evidence handling. | 3 | 3 | Moderate | Reduce; graceful degradation, isolation, recovery and no unsafe fallback. |
| `AI-RISK-031` | Resource or cost exhaustion | Operational Resilience | Unbounded model, connector or storage use causes denial of service or excessive cost. | 4 | 3 | High | Reduce; quotas, rate limits, budgets and alerts. |
| `AI-RISK-032` | Unsafe public claim or disclosure | Publication and Reputation | The repository overstates implementation, validation, compliance or releases sensitive data. | 3 | 4 | High | Reduce; release gates, redaction, claim review and correction. |
| `AI-RISK-033` | License or third-party rights failure | Legal and Regulatory | Code, data, diagrams or generated content are published without compatible rights. | 3 | 3 | Moderate | Reduce; provenance, license review and release checks. |
| `AI-RISK-034` | Jurisdiction or legal-authority mismatch | Legal and Regulatory | Data or operations exceed applicable mandate, jurisdiction or legal basis. | 2 | 5 | Critical | Avoid; explicit scope, human legal authority and denial. |
| `AI-RISK-035` | Insufficient testing or false assurance | Research and Assumptions | Passing narrow tests are generalized into unsupported security or readiness claims. | 3 | 4 | High | Reduce; risk-based coverage, limitations and validation governance. |
| `AI-RISK-036` | Stale risk register or overdue review | Repository and Governance | Material changes occur without updating ratings, owners or treatments. | 3 | 4 | High | Reduce; triggers, release review and automated overdue detection. |
| `AI-RISK-037` | Invalid or stale architectural assumption | Research and Assumptions | Design depends on an assumption that is unsupported, expired or false. | 3 | 4 | High | Reduce; assumption register, validation and conservative treatment. |
| `AI-RISK-038` | Role concentration and lack of independent assurance | Authority and Accountability | One person performs multiple governance roles and internal review is overstated. | 4 | 3 | High | Reduce; disclosure, immutable evidence and external review when available. |
| `AI-RISK-039` | Blockchain or smart-contract side effect | Web3 and Distributed Systems | A transaction, key use or contract call causes unintended irreversible state or asset change. | 2 | 4 | High | Avoid in public demos; use testnets, simulation and human approval. |
| `AI-RISK-040` | Communication misdelivery or impersonation | Connector and External Service | Automated communication reaches the wrong recipient or appears to impersonate an officer. | 3 | 4 | High | Reduce; recipient validation, approval, idempotency and explicit attribution. |

## 49. Baseline Risk Record Requirements

- **RISK-REQ-895:** `AI-RISK-001` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-896:** `AI-RISK-001` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-897:** `AI-RISK-002` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-898:** `AI-RISK-002` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-899:** `AI-RISK-003` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-900:** `AI-RISK-003` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-901:** `AI-RISK-004` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-902:** `AI-RISK-004` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-903:** `AI-RISK-005` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-904:** `AI-RISK-005` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-905:** `AI-RISK-006` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-906:** `AI-RISK-006` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-907:** `AI-RISK-007` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-908:** `AI-RISK-007` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-909:** `AI-RISK-008` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-910:** `AI-RISK-008` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-911:** `AI-RISK-009` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-912:** `AI-RISK-009` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-913:** `AI-RISK-010` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-914:** `AI-RISK-010` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-915:** `AI-RISK-011` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-916:** `AI-RISK-011` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-917:** `AI-RISK-012` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-918:** `AI-RISK-012` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-919:** `AI-RISK-013` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-920:** `AI-RISK-013` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-921:** `AI-RISK-014` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-922:** `AI-RISK-014` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-923:** `AI-RISK-015` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-924:** `AI-RISK-015` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-925:** `AI-RISK-016` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-926:** `AI-RISK-016` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-927:** `AI-RISK-017` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-928:** `AI-RISK-017` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-929:** `AI-RISK-018` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-930:** `AI-RISK-018` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-931:** `AI-RISK-019` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-932:** `AI-RISK-019` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-933:** `AI-RISK-020` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-934:** `AI-RISK-020` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-935:** `AI-RISK-021` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-936:** `AI-RISK-021` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-937:** `AI-RISK-022` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-938:** `AI-RISK-022` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-939:** `AI-RISK-023` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-940:** `AI-RISK-023` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-941:** `AI-RISK-024` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-942:** `AI-RISK-024` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-943:** `AI-RISK-025` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-944:** `AI-RISK-025` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-945:** `AI-RISK-026` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-946:** `AI-RISK-026` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-947:** `AI-RISK-027` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-948:** `AI-RISK-027` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-949:** `AI-RISK-028` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-950:** `AI-RISK-028` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-951:** `AI-RISK-029` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-952:** `AI-RISK-029` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-953:** `AI-RISK-030` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-954:** `AI-RISK-030` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-955:** `AI-RISK-031` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-956:** `AI-RISK-031` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-957:** `AI-RISK-032` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-958:** `AI-RISK-032` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-959:** `AI-RISK-033` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-960:** `AI-RISK-033` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-961:** `AI-RISK-034` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-962:** `AI-RISK-034` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-963:** `AI-RISK-035` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-964:** `AI-RISK-035` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-965:** `AI-RISK-036` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-966:** `AI-RISK-036` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-967:** `AI-RISK-037` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-968:** `AI-RISK-037` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-969:** `AI-RISK-038` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-970:** `AI-RISK-038` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-971:** `AI-RISK-039` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-972:** `AI-RISK-039` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.
- **RISK-REQ-973:** `AI-RISK-040` shall remain in the baseline register until formally closed or superseded.
- **RISK-REQ-974:** `AI-RISK-040` shall have a human owner and shall link rating evidence, controls, tests, review dates and release impact.

## 50. Minimum Validation Checklist

Before a risk is treated as Assessed, Accepted, Closed or release-reviewed, confirm:

- [ ] immutable `AI-RISK-NNN` identifier exists;
- [ ] description, category and cause are complete;
- [ ] threats, assets and affected parties are linked;
- [ ] inherent likelihood, impact, score and level are justified;
- [ ] rating confidence and uncertainty are recorded;
- [ ] current and planned controls are separated;
- [ ] control effectiveness has evidence;
- [ ] residual likelihood, impact, score and level are justified;
- [ ] owner is an accountable human;
- [ ] review date and triggers are current;
- [ ] treatment or acceptance authority is recorded;
- [ ] acceptance is time-bounded and permitted by policy;
- [ ] Critical and canonical-prohibition risks are not accepted;
- [ ] privacy and fundamental-rights impacts are considered;
- [ ] evidence integrity and authorization impacts are considered;
- [ ] applicable tests and incidents are linked;
- [ ] assumptions, research items and decisions are linked;
- [ ] release impact is explicit;
- [ ] revision history is append-only;
- [ ] no secrets or unnecessary personal data are present;
- [ ] traceability is bidirectional;
- [ ] exact method and register versions are recorded.

## 51. Limitations

- The initial catalogue is a research baseline, not an operational institutional risk assessment.
- The qualitative 5×5 method supports prioritization but does not produce objective probability or loss estimates.
- Risk ratings depend on evidence, assumptions, environment and reviewer judgment.
- Absence of a recorded risk does not prove absence of risk.
- Control implementation does not prove effectiveness.
- Passing tests do not prove complete security or legal compliance.
- Legal and regulatory applicability requires competent institutional or legal review.
- External model and service opacity may limit confidence.
- Internal role concentration is disclosed and is not independent external assurance.
- Documents 27–30 remain forward dependencies for research, assumptions, decision logging and compliance mapping.

## 52. Change Control

- **RISK-REQ-975:** Changes to this document shall identify rationale.
- **RISK-REQ-976:** Changes shall identify affected risk records.
- **RISK-REQ-977:** Changes shall identify affected rating semantics.
- **RISK-REQ-978:** Changes shall identify affected categories.
- **RISK-REQ-979:** Changes shall identify affected statuses.
- **RISK-REQ-980:** Changes shall identify security impact.
- **RISK-REQ-981:** Changes shall identify privacy and rights impact.
- **RISK-REQ-982:** Changes shall identify evidence impact.
- **RISK-REQ-983:** Changes shall identify implementation impact.
- **RISK-REQ-984:** Changes shall identify testing impact.
- **RISK-REQ-985:** Changes shall identify release impact.
- **RISK-REQ-986:** Changes shall identify migration.
- **RISK-REQ-987:** Changes shall identify rollback.
- **RISK-REQ-988:** Editorial corrections shall use a patch version when meaning is unchanged.
- **RISK-REQ-989:** Backward-compatible substantive additions shall use a minor version.
- **RISK-REQ-990:** Incompatible risk-governance changes shall use a major version.
- **RISK-REQ-991:** Material risk-method changes shall require an ADR where applicable.
- **RISK-REQ-992:** Changes shall require Project Founder approval.
- **RISK-REQ-993:** Changes shall receive Privacy and Governance Reviewer assessment.
- **RISK-REQ-994:** Security-impacting changes shall receive Security Reviewer assessment.
- **RISK-REQ-995:** Implementation-impacting changes shall receive Implementation Reviewer assessment.
- **RISK-REQ-996:** Release-impacting changes shall receive Release Reviewer assessment.
- **RISK-REQ-997:** Historical risk records shall not be silently recalculated.
- **RISK-REQ-998:** Risk and requirement identifiers shall not be reused.
- **RISK-REQ-999:** Forward-dependency reconciliation shall occur before this document becomes Approved.

## 53. Consolidation Record

Version 1.0.0 consolidates both legacy AI Risk Register variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-RISK-001`;
- normalizes the authoritative filename to `26_AI_RISK_REGISTER.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves ID, description or risk, category, cause, impact, likelihood, risk level, controls, residual risk, owner, review date and status;
- preserves mandatory risk review before every release;
- adds immutable risk-record identifiers, complete record schema, taxonomy, lifecycle, 5×5 rating definitions, confidence, inherent and residual assessment, control effectiveness, treatment, acceptance authority, escalation, monitoring, incidents, closure, versioning, automation, publication limits and traceability;
- adds an initial synthetic baseline catalogue of forty OBDIA risks;
- identifies documents 27–30 as forward dependencies;
- treats Enterprise and non-Enterprise legacy files as source variants of the same immutable document;
- creates no operational risk acceptance, implementation, deployment, certification, compliance or legal-authority claim.

## 54. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-06 | Draft | Privacy and Governance Reviewer; approval reserved to Project Founder | Constitutional consolidation of both legacy risk-register variants: preserved all fields and release review; added rating, ownership, treatment, acceptance, evidence, escalation, baseline catalogue and traceability. |
