# RESEARCH BACKLOG

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-RSCH-001 |
| **Title** | Research Backlog |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative for backlog-governance rules; individual backlog items are Research records |
| **Authority** | Project Founder |
| **Owner** | Research Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define the governed intake, triage, prioritization, ownership, sourcing, hypothesis, dependency, validation, evidence, disposition, closure, archival, publication and traceability requirements for OBDIA research items, and establish an initial synthetic research backlog. |
| **Scope** | Open research questions concerning officer-agent binding, identity, authorization, evidence, AI security, connectors, privacy, human oversight, software and supply-chain security, testing, repository governance, release assurance, Web3, cloud and isolated research environments within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `07_THREAT_MODEL_BASELINE.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; `18_EVIDENCE_MODEL.md`; `19_AGENT_LIFECYCLE_MODEL.md`; `20_CONNECTOR_SECURITY_POLICY.md`; `21_SECURE_CODING_STANDARD.md`; `22_TESTING_STANDARD.md`; `23_DIAGRAM_STANDARD.md`; `24_GITHUB_REPOSITORY_STANDARD.md`; `25_DOCUMENT_VERSIONING_POLICY.md`; `26_AI_RISK_REGISTER.md`; forward dependencies `28_ASSUMPTIONS_REGISTER.md`, `29_DECISION_LOG_POLICY.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-RES-001`; `OBDIA-TM-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001`; `OBDIA-EVID-001`; `OBDIA-AGENT-001`; `OBDIA-CONN-001`; `OBDIA-CODE-001`; `OBDIA-TEST-001`; `OBDIA-DIAG-001`; `OBDIA-GH-001`; `OBDIA-VER-001`; `OBDIA-RISK-001` |
| **Cross-References** | Research items; source records; assumptions; hypotheses; threats; risks; requirements; ADRs; decision-log entries; implementation plans; test plans; validation packages; exceptions; changes; baseline manifests; releases and compliance mappings |
| **Assumptions** | The initial backlog is a synthetic research baseline derived from the current OBDIA architecture. It does not describe a deployed institution, real investigation or validated implementation. |
| **Constraints** | Research shall use synthetic data, mocks, testnets, local laboratories, read-only demonstrations, archived-authorized material and explicitly authorized isolated environments. Research shall not require unauthorized access, malware deployment, credential theft, offensive operations, unlawful surveillance, unlawful deanonymization, autonomous legal or coercive action, or interaction with uncontrolled criminal infrastructure. |
| **Security Considerations** | A manipulated, stale or weakly governed backlog can turn speculation into architecture, conceal unsupported assumptions, create unsafe experiments, misdirect resources, bypass change control, overstate evidence or normalize prohibited research. |
| **Validation Criteria** | Every active research item has a unique immutable identifier, precise question, motivation, owner, priority, dependencies, source state, uncertainty, validation plan, expected deliverables, risk and assumption links, status, review date and closure evidence; no item changes architecture without governed validation and decision records. |
| **Implementation Relationship** | This document governs research records and contains an initial synthetic backlog. It does not validate hypotheses, select architecture, approve implementation, authorize operational research, accept risk or establish production readiness. |

---

## 1. Purpose

The non-Enterprise source required the backlog to track open research questions, dependencies, references, validation status and priority.

The Enterprise source additionally required:

- motivation;
- validation plan;
- status;
- expected deliverables.

It also established the essential boundary that no hypothesis becomes architecture without validation.

This consolidation preserves every original field and boundary. It adds governed intake, triage, ownership, source discipline, uncertainty, assumptions, risks, validation evidence, closure, archival, publication and change-control integration.

The policy sections are normative. Individual backlog items are hierarchy-9 research records. They do not possess normative authority and cannot amend architecture, requirements, controls or legal boundaries.


## 2. Fundamental Research-Governance Rules

- **RSCH-REQ-001:** Every material open research question shall be recorded.
- **RSCH-REQ-002:** Every research item shall have an immutable identifier.
- **RSCH-REQ-003:** Every research item shall identify a precise research question.
- **RSCH-REQ-004:** Every research item shall identify motivation.
- **RSCH-REQ-005:** Every research item shall identify dependencies.
- **RSCH-REQ-006:** Every research item shall identify references or an explicit reference-acquisition state.
- **RSCH-REQ-007:** Every research item shall identify a validation plan.
- **RSCH-REQ-008:** Every research item shall identify priority.
- **RSCH-REQ-009:** Every research item shall identify status.
- **RSCH-REQ-010:** Every research item shall identify expected deliverables.
- **RSCH-REQ-011:** Every active research item shall have an accountable human owner.
- **RSCH-REQ-012:** An AI system shall not own or approve a research item.
- **RSCH-REQ-013:** An AI system shall not accept research-related risk.
- **RSCH-REQ-014:** Backlog entries shall have no normative authority.
- **RSCH-REQ-015:** No hypothesis shall become architecture without validation.
- **RSCH-REQ-016:** No research result shall become architecture without an accepted ADR or other governing decision where required.
- **RSCH-REQ-017:** No research result shall become a requirement without governed change control.
- **RSCH-REQ-018:** No research result shall authorize implementation by implication.
- **RSCH-REQ-019:** No research result shall authorize access by implication.
- **RSCH-REQ-020:** No research result shall create independent AI legal authority.
- **RSCH-REQ-021:** Research shall remain within canonical safety boundaries.
- **RSCH-REQ-022:** Research shall default to synthetic, simulated, read-only or otherwise controlled methods.
- **RSCH-REQ-023:** Research uncertainty shall remain explicit.
- **RSCH-REQ-024:** Negative or inconclusive results shall remain visible.
- **RSCH-REQ-025:** Unsupported claims shall not be converted into design facts.
- **RSCH-REQ-026:** Repository merge shall not establish research validation.
- **RSCH-REQ-027:** Publication shall not establish research validation.
- **RSCH-REQ-028:** Priority shall not establish truth or authority.
- **RSCH-REQ-029:** Critical safety or rights concerns shall block unsafe experimentation.
- **RSCH-REQ-030:** Research history shall not be silently rewritten.

## 3. Authority Boundary and Record Classification

- **RSCH-REQ-031:** This policy shall be normative for backlog-governance mechanics.
- **RSCH-REQ-032:** Individual research items shall remain research records.
- **RSCH-REQ-033:** Individual research items shall remain below normative architecture in the authority hierarchy.
- **RSCH-REQ-034:** A backlog item shall not override the Canonical Project Definition.
- **RSCH-REQ-035:** A backlog item shall not override the Constitution.
- **RSCH-REQ-036:** A backlog item shall not override an approved normative document.
- **RSCH-REQ-037:** A backlog item shall not override an accepted ADR.
- **RSCH-REQ-038:** A backlog item shall not override authorization policy.
- **RSCH-REQ-039:** A backlog item shall not override evidence requirements.
- **RSCH-REQ-040:** A backlog item shall not override canonical prohibitions.
- **RSCH-REQ-041:** A research hypothesis shall remain distinguishable from a verified fact.
- **RSCH-REQ-042:** A research proposal shall remain distinguishable from an experiment.
- **RSCH-REQ-043:** An experiment shall remain distinguishable from an implementation.
- **RSCH-REQ-044:** An implementation shall remain distinguishable from validation.
- **RSCH-REQ-045:** Validation shall remain distinguishable from approval.
- **RSCH-REQ-046:** Approval shall remain distinguishable from publication.
- **RSCH-REQ-047:** Research records shall identify whether they are public, internal or restricted.
- **RSCH-REQ-048:** Classification shall reflect source, method and potential harm.
- **RSCH-REQ-049:** Restricted research records shall not be copied into public artifacts without review.
- **RSCH-REQ-050:** Research authority shall remain attributable to authorized human roles.
- **RSCH-REQ-051:** The Research Reviewer shall assess methodological quality.
- **RSCH-REQ-052:** The Security Reviewer shall assess security-sensitive methods.
- **RSCH-REQ-053:** The Privacy and Governance Reviewer shall assess rights and privacy impacts.
- **RSCH-REQ-054:** The Project Founder shall approve normative changes derived from research where required.
- **RSCH-REQ-055:** Role concentration shall be disclosed.

## 4. Research Record Types

- **RSCH-REQ-056:** The backlog shall distinguish research items from source records.
- **RSCH-REQ-057:** The backlog shall distinguish research items from assumptions.
- **RSCH-REQ-058:** The backlog shall distinguish research items from hypotheses.
- **RSCH-REQ-059:** The backlog shall distinguish research items from risks.
- **RSCH-REQ-060:** The backlog shall distinguish research items from threats.
- **RSCH-REQ-061:** The backlog shall distinguish research items from ADRs.
- **RSCH-REQ-062:** The backlog shall distinguish research items from implementation tasks.
- **RSCH-REQ-063:** The backlog shall distinguish research items from test cases.
- **RSCH-REQ-064:** The backlog shall distinguish research items from validation decisions.
- **RSCH-REQ-065:** The backlog shall distinguish research items from exceptions.
- **RSCH-REQ-066:** The backlog shall distinguish research items from change records.
- **RSCH-REQ-067:** Research questions shall use `RSCH-ITEM-NNN` identifiers.
- **RSCH-REQ-068:** Source records shall use the identifier format governed by the Research and Source Policy.
- **RSCH-REQ-069:** Hypothesis records shall have immutable identifiers where maintained separately.
- **RSCH-REQ-070:** Experiment records shall have immutable identifiers.
- **RSCH-REQ-071:** Research evidence packages shall have immutable identifiers.
- **RSCH-REQ-072:** Research review records shall have immutable identifiers.
- **RSCH-REQ-073:** Research closure records shall have immutable identifiers.
- **RSCH-REQ-074:** Identifier namespaces shall not overlap ambiguously.
- **RSCH-REQ-075:** Identifiers shall not encode lifecycle status.
- **RSCH-REQ-076:** Identifiers shall not encode priority.
- **RSCH-REQ-077:** Identifiers shall not be reused.
- **RSCH-REQ-078:** Superseded records shall retain their identifiers.
- **RSCH-REQ-079:** Duplicate identifiers shall be rejected.
- **RSCH-REQ-080:** Unknown record types shall not acquire authority silently.

## 5. Mandatory Research-Item Schema

- **RSCH-REQ-081:** Every research item shall include `ID`.
- **RSCH-REQ-082:** Every research item shall include `Title`.
- **RSCH-REQ-083:** Every research item shall include `Research Question`.
- **RSCH-REQ-084:** Every research item shall include `Motivation`.
- **RSCH-REQ-085:** Every research item shall include `Scope`.
- **RSCH-REQ-086:** Every research item shall include `Out of Scope`.
- **RSCH-REQ-087:** Every research item shall include `Owner`.
- **RSCH-REQ-088:** Every research item shall include `Contributors` where applicable.
- **RSCH-REQ-089:** Every research item shall include `Priority`.
- **RSCH-REQ-090:** Every research item shall include `Status`.
- **RSCH-REQ-091:** Every research item shall include `Created Date`.
- **RSCH-REQ-092:** Every active research item shall include `Next Review Date`.
- **RSCH-REQ-093:** Every research item shall include `Dependencies`.
- **RSCH-REQ-094:** Every research item shall include `Affected Documents`.
- **RSCH-REQ-095:** Every research item shall include `Affected Requirements` where known.
- **RSCH-REQ-096:** Every research item shall include `Risk Links` where applicable.
- **RSCH-REQ-097:** Every research item shall include `Assumption Links` where applicable.
- **RSCH-REQ-098:** Every research item shall include `Threat Links` where applicable.
- **RSCH-REQ-099:** Every research item shall include `Source Records`.
- **RSCH-REQ-100:** Every research item shall include `Reference State`.
- **RSCH-REQ-101:** Every research item shall include `Known Facts`.
- **RSCH-REQ-102:** Every research item shall include `Open Uncertainties`.
- **RSCH-REQ-103:** Every research item shall include `Hypotheses` where applicable.
- **RSCH-REQ-104:** Every research item shall include `Validation Plan`.
- **RSCH-REQ-105:** Every research item shall include `Validation Criteria`.
- **RSCH-REQ-106:** Every research item shall include `Required Environment`.
- **RSCH-REQ-107:** Every research item shall include `Required Data`.
- **RSCH-REQ-108:** Every research item shall include `Security Constraints`.
- **RSCH-REQ-109:** Every research item shall include `Privacy Constraints`.
- **RSCH-REQ-110:** Every research item shall include `Ethical Constraints`.
- **RSCH-REQ-111:** Every research item shall include `Expected Deliverables`.
- **RSCH-REQ-112:** Every research item shall include `Decision Boundary`.
- **RSCH-REQ-113:** Every research item shall include `Publication Boundary`.
- **RSCH-REQ-114:** Every research item shall include `Known Limitations`.
- **RSCH-REQ-115:** Every research item shall include `Result Summary` after execution.
- **RSCH-REQ-116:** Every completed research item shall include `Evidence References`.
- **RSCH-REQ-117:** Every completed research item shall include `Validation Outcome`.
- **RSCH-REQ-118:** Every completed research item shall include `Disposition`.
- **RSCH-REQ-119:** Every closed research item shall include `Closure Authority`.
- **RSCH-REQ-120:** Every closed research item shall include `Closure Date`.

## 6. Research-Question Formulation

- **RSCH-REQ-121:** Research questions shall be answerable through defined evidence or analysis.
- **RSCH-REQ-122:** Research questions shall identify the object of study.
- **RSCH-REQ-123:** Research questions shall identify the intended decision supported.
- **RSCH-REQ-124:** Research questions shall identify the relevant environment.
- **RSCH-REQ-125:** Research questions shall identify relevant actors.
- **RSCH-REQ-126:** Research questions shall identify relevant trust boundaries.
- **RSCH-REQ-127:** Research questions shall identify relevant security properties.
- **RSCH-REQ-128:** Research questions shall identify relevant privacy or rights properties.
- **RSCH-REQ-129:** Research questions shall identify temporal scope where material.
- **RSCH-REQ-130:** Research questions shall avoid presupposing the desired answer.
- **RSCH-REQ-131:** Research questions shall avoid embedding unsupported conclusions.
- **RSCH-REQ-132:** Research questions shall avoid vague terms without definitions.
- **RSCH-REQ-133:** Research questions shall use glossary terminology.
- **RSCH-REQ-134:** Research questions shall distinguish feasibility from desirability.
- **RSCH-REQ-135:** Research questions shall distinguish technical capability from lawful authority.
- **RSCH-REQ-136:** Research questions shall distinguish security from compliance.
- **RSCH-REQ-137:** Research questions shall distinguish evidence integrity from authenticity.
- **RSCH-REQ-138:** Research questions shall distinguish model confidence from factual truth.
- **RSCH-REQ-139:** Research questions shall distinguish Deep Web from Dark Web.
- **RSCH-REQ-140:** Research questions shall not require canonical prohibitions.
- **RSCH-REQ-141:** Overly broad questions shall be decomposed.
- **RSCH-REQ-142:** Duplicate questions shall be linked and consolidated where appropriate.
- **RSCH-REQ-143:** Question changes shall preserve revision history.
- **RSCH-REQ-144:** Material question changes shall trigger renewed triage.
- **RSCH-REQ-145:** Ambiguous questions shall not enter active validation.

## 7. Motivation and Decision Relevance

- **RSCH-REQ-146:** Motivation shall explain why the question matters.
- **RSCH-REQ-147:** Motivation shall identify the affected project objective.
- **RSCH-REQ-148:** Motivation shall identify affected stakeholders.
- **RSCH-REQ-149:** Motivation shall identify affected architecture or governance areas.
- **RSCH-REQ-150:** Motivation shall identify the harm of leaving the question unresolved.
- **RSCH-REQ-151:** Motivation shall identify potential benefit without overstating certainty.
- **RSCH-REQ-152:** Motivation shall identify whether the item reduces risk.
- **RSCH-REQ-153:** Motivation shall identify whether the item resolves an assumption.
- **RSCH-REQ-154:** Motivation shall identify whether the item supports an ADR.
- **RSCH-REQ-155:** Motivation shall identify whether the item supports testing.
- **RSCH-REQ-156:** Motivation shall identify whether the item supports publication.
- **RSCH-REQ-157:** Motivation shall identify whether the item supports compliance mapping without claiming compliance.
- **RSCH-REQ-158:** Motivation shall remain separate from expected result.
- **RSCH-REQ-159:** Motivation shall not manufacture urgency.
- **RSCH-REQ-160:** Motivation shall not use fear or reputational pressure as evidence.
- **RSCH-REQ-161:** Motivation shall not imply official institutional demand without evidence.
- **RSCH-REQ-162:** Motivation shall not imply operational deployment.
- **RSCH-REQ-163:** Motivation shall identify decision deadlines only when real.
- **RSCH-REQ-164:** Motivation shall identify dependencies on external developments where material.
- **RSCH-REQ-165:** Motivation changes shall be attributable.

## 8. Intake

- **RSCH-REQ-166:** Research intake shall use a governed submission path.
- **RSCH-REQ-167:** Intake shall capture the submitter.
- **RSCH-REQ-168:** Intake shall capture the initial question.
- **RSCH-REQ-169:** Intake shall capture motivation.
- **RSCH-REQ-170:** Intake shall capture known dependencies.
- **RSCH-REQ-171:** Intake shall capture known risks.
- **RSCH-REQ-172:** Intake shall capture proposed sources.
- **RSCH-REQ-173:** Intake shall capture proposed validation approach.
- **RSCH-REQ-174:** Intake shall capture expected deliverables.
- **RSCH-REQ-175:** Intake shall capture desired priority with rationale.
- **RSCH-REQ-176:** Intake shall identify whether sensitive data may be involved.
- **RSCH-REQ-177:** Intake shall identify whether external systems may be involved.
- **RSCH-REQ-178:** Intake shall identify whether human subjects may be implicated.
- **RSCH-REQ-179:** Intake shall identify whether isolated research is required.
- **RSCH-REQ-180:** Intake shall reject canonical-prohibition requests.
- **RSCH-REQ-181:** Intake shall reject unauthorized-access requests.
- **RSCH-REQ-182:** Intake shall reject malware-deployment requests.
- **RSCH-REQ-183:** Intake shall reject unlawful surveillance requests.
- **RSCH-REQ-184:** Intake shall reject unlawful deanonymization requests.
- **RSCH-REQ-185:** Intake shall reject uncontrolled criminal-infrastructure interaction.
- **RSCH-REQ-186:** Rejected intake shall preserve safe rationale.
- **RSCH-REQ-187:** Accepted intake shall receive an immutable item identifier.
- **RSCH-REQ-188:** Intake shall not assign architecture authority.
- **RSCH-REQ-189:** Intake automation shall remain advisory.
- **RSCH-REQ-190:** Intake records shall minimize personal data.

## 9. Triage

- **RSCH-REQ-191:** Every accepted intake shall receive triage.
- **RSCH-REQ-192:** Triage shall verify scope alignment.
- **RSCH-REQ-193:** Triage shall verify authority alignment.
- **RSCH-REQ-194:** Triage shall verify source-policy compatibility.
- **RSCH-REQ-195:** Triage shall verify method safety.
- **RSCH-REQ-196:** Triage shall verify data safety.
- **RSCH-REQ-197:** Triage shall verify legal and ethical boundaries where material.
- **RSCH-REQ-198:** Triage shall identify duplicate items.
- **RSCH-REQ-199:** Triage shall identify prerequisite items.
- **RSCH-REQ-200:** Triage shall identify blocking assumptions.
- **RSCH-REQ-201:** Triage shall identify linked risks.
- **RSCH-REQ-202:** Triage shall identify linked threats.
- **RSCH-REQ-203:** Triage shall identify required reviewers.
- **RSCH-REQ-204:** Triage shall identify required environment.
- **RSCH-REQ-205:** Triage shall identify whether external approval is required.
- **RSCH-REQ-206:** Triage shall identify expected evidence strength.
- **RSCH-REQ-207:** Triage shall identify whether the question can be answered from existing evidence.
- **RSCH-REQ-208:** Triage shall identify whether the item should become a documentation correction instead.
- **RSCH-REQ-209:** Triage shall identify whether the item should become a defect or change record instead.
- **RSCH-REQ-210:** Triage shall identify whether the item should become an ADR proposal instead.
- **RSCH-REQ-211:** Triage shall not elevate speculation to fact.
- **RSCH-REQ-212:** Triage shall assign an initial priority.
- **RSCH-REQ-213:** Triage shall assign an initial status.
- **RSCH-REQ-214:** Triage shall identify the next review date.
- **RSCH-REQ-215:** Triage decisions shall be attributable.

## 10. Priority Model

- **RSCH-REQ-216:** Permitted priorities shall be `P0`, `P1`, `P2`, `P3` and `P4`.
- **RSCH-REQ-217:** `P0` shall mean an urgent blocker involving critical safety, authority, rights, integrity or baseline incompatibility.
- **RSCH-REQ-218:** `P1` shall mean a high-priority item needed for a near-term governance, security or baseline decision.
- **RSCH-REQ-219:** `P2` shall mean an important item with material architectural or validation value.
- **RSCH-REQ-220:** `P3` shall mean a useful item without an immediate blocking effect.
- **RSCH-REQ-221:** `P4` shall mean exploratory or future research.
- **RSCH-REQ-222:** Priority shall reflect decision need.
- **RSCH-REQ-223:** Priority shall reflect risk reduction.
- **RSCH-REQ-224:** Priority shall reflect dependency criticality.
- **RSCH-REQ-225:** Priority shall reflect safety and rights impact.
- **RSCH-REQ-226:** Priority shall reflect evidence availability.
- **RSCH-REQ-227:** Priority shall reflect validation feasibility.
- **RSCH-REQ-228:** Priority shall not reflect contributor seniority.
- **RSCH-REQ-229:** Priority shall not imply truth.
- **RSCH-REQ-230:** Priority shall not imply approval.
- **RSCH-REQ-231:** Priority shall not override canonical prohibitions.
- **RSCH-REQ-232:** Priority changes shall include rationale.
- **RSCH-REQ-233:** Priority changes shall be attributable.
- **RSCH-REQ-234:** P0 items shall receive prompt human review.
- **RSCH-REQ-235:** P0 status shall not authorize unsafe shortcuts.
- **RSCH-REQ-236:** P1 and P2 items shall identify expected decision windows.
- **RSCH-REQ-237:** P3 and P4 items shall remain reviewable for staleness.
- **RSCH-REQ-238:** Priority conflicts shall be resolved by authorized human roles.
- **RSCH-REQ-239:** Automated priority recommendations shall remain advisory.
- **RSCH-REQ-240:** Backlog reports shall preserve the rationale behind priority.

## 11. Research-Item Status Model

- **RSCH-REQ-241:** Permitted research-item statuses shall be `Proposed`, `Triaged`, `Planned`, `In Progress`, `Blocked`, `Under Review`, `Validated`, `Inconclusive`, `Rejected`, `Superseded`, `Closed` and `Archived`.
- **RSCH-REQ-242:** `Proposed` shall mean intake exists but triage is incomplete.
- **RSCH-REQ-243:** `Triaged` shall mean scope, safety and priority have been assessed.
- **RSCH-REQ-244:** `Planned` shall mean an approved validation plan exists.
- **RSCH-REQ-245:** `In Progress` shall mean governed research execution has started.
- **RSCH-REQ-246:** `Blocked` shall mean a documented prerequisite prevents progress.
- **RSCH-REQ-247:** `Under Review` shall mean results and evidence are being assessed.
- **RSCH-REQ-248:** `Validated` shall mean the research conclusion met defined validation criteria for its stated scope.
- **RSCH-REQ-249:** `Inconclusive` shall mean available evidence did not support a reliable conclusion.
- **RSCH-REQ-250:** `Rejected` shall mean the item is out of scope, unsafe, duplicative or otherwise unsuitable.
- **RSCH-REQ-251:** `Superseded` shall mean another research item replaces it.
- **RSCH-REQ-252:** `Closed` shall mean disposition and closure evidence are complete.
- **RSCH-REQ-253:** `Archived` shall mean the closed historical record is retained read-only.
- **RSCH-REQ-254:** Research-item status shall remain distinct from repository document lifecycle status.
- **RSCH-REQ-255:** Research-item status shall remain distinct from test-run status.
- **RSCH-REQ-256:** Research-item status shall remain distinct from validation-decision outcome.
- **RSCH-REQ-257:** Research-item status shall not be inferred from branch or PR state.
- **RSCH-REQ-258:** Status transitions shall be attributable.
- **RSCH-REQ-259:** Status transitions shall record rationale.
- **RSCH-REQ-260:** Invalid transitions shall be rejected.
- **RSCH-REQ-261:** A Proposed item shall not transition directly to Validated.
- **RSCH-REQ-262:** A Triaged item shall not transition to In Progress without a validation plan.
- **RSCH-REQ-263:** A Blocked item shall identify blocking conditions and owner.
- **RSCH-REQ-264:** An Inconclusive item may return to Planned only with a revised plan.
- **RSCH-REQ-265:** A Rejected item shall preserve safe rationale.
- **RSCH-REQ-266:** A Superseded item shall identify its successor.
- **RSCH-REQ-267:** A Closed item shall identify disposition.
- **RSCH-REQ-268:** An Archived item shall remain read-only.
- **RSCH-REQ-269:** Status corrections shall be additive.
- **RSCH-REQ-270:** Unknown status shall block authoritative use.

## 12. Ownership and Accountability

- **RSCH-REQ-271:** Every active research item shall have one accountable human owner.
- **RSCH-REQ-272:** The owner shall be identifiable.
- **RSCH-REQ-273:** The owner shall understand the research question and boundaries.
- **RSCH-REQ-274:** The owner shall maintain status and review dates.
- **RSCH-REQ-275:** The owner shall maintain dependencies.
- **RSCH-REQ-276:** The owner shall maintain source and evidence links.
- **RSCH-REQ-277:** The owner shall escalate safety and rights concerns.
- **RSCH-REQ-278:** The owner shall not accept risk outside delegated authority.
- **RSCH-REQ-279:** The owner shall not self-approve consequential architecture changes.
- **RSCH-REQ-280:** Contributors shall be distinguished from owners.
- **RSCH-REQ-281:** Reviewers shall be distinguished from owners.
- **RSCH-REQ-282:** AI assistance shall be distinguished from human authorship.
- **RSCH-REQ-283:** Automation shall not become the accountable owner.
- **RSCH-REQ-284:** Ownership changes shall be attributable.
- **RSCH-REQ-285:** Ownership changes shall preserve prior accountability.
- **RSCH-REQ-286:** Orphaned active items shall be blocked or reassigned.
- **RSCH-REQ-287:** Inactive owners shall trigger review.
- **RSCH-REQ-288:** Role concentration shall be disclosed.
- **RSCH-REQ-289:** Self-review shall not be described as independent review.
- **RSCH-REQ-290:** Owner identity records shall minimize personal data.

## 13. Dependencies

- **RSCH-REQ-291:** Every research item shall declare material dependencies.
- **RSCH-REQ-292:** Dependencies shall use immutable identifiers where available.
- **RSCH-REQ-293:** Dependencies shall distinguish research items from assumptions.
- **RSCH-REQ-294:** Dependencies shall distinguish risks from threats.
- **RSCH-REQ-295:** Dependencies shall distinguish normative documents from informative sources.
- **RSCH-REQ-296:** Dependencies shall distinguish prerequisites from related work.
- **RSCH-REQ-297:** Dependencies shall identify status.
- **RSCH-REQ-298:** Dependencies shall identify version where material.
- **RSCH-REQ-299:** Dependencies shall identify blocking effect.
- **RSCH-REQ-300:** Dependencies shall identify owner where material.
- **RSCH-REQ-301:** Unknown dependencies shall remain explicit.
- **RSCH-REQ-302:** Circular dependencies shall be identified.
- **RSCH-REQ-303:** Unresolvable circular dependencies shall block validation.
- **RSCH-REQ-304:** Completed dependency work shall not imply compatibility automatically.
- **RSCH-REQ-305:** Dependency changes shall trigger impact review.
- **RSCH-REQ-306:** Removed dependencies shall have rationale.
- **RSCH-REQ-307:** New dependencies shall receive source and security review where material.
- **RSCH-REQ-308:** External dependencies shall identify availability and reproducibility limitations.
- **RSCH-REQ-309:** Provider dependencies shall identify version or change risk.
- **RSCH-REQ-310:** Research environment dependencies shall identify authorization requirements.
- **RSCH-REQ-311:** Data dependencies shall identify provenance and classification.
- **RSCH-REQ-312:** Tool dependencies shall identify version and supply-chain considerations.
- **RSCH-REQ-313:** Dependency drift shall be reviewed.
- **RSCH-REQ-314:** Dependency history shall remain retrievable.
- **RSCH-REQ-315:** Broken dependency links shall block closure where material.

## 14. Source Discipline

- **RSCH-REQ-316:** Research shall conform to `OBDIA-RES-001`.
- **RSCH-REQ-317:** Verified facts shall be distinguished from proposed interpretations.
- **RSCH-REQ-318:** Primary sources shall be preferred for technical claims.
- **RSCH-REQ-319:** Official standards and documentation shall be preferred for normative framework claims.
- **RSCH-REQ-320:** Academic publications shall be preferred for empirical or theoretical claims.
- **RSCH-REQ-321:** Marketing claims shall not be treated as independent evidence.
- **RSCH-REQ-322:** Unverified blog claims shall not become architecture.
- **RSCH-REQ-323:** Source authority shall be assessed.
- **RSCH-REQ-324:** Source currency shall be assessed.
- **RSCH-REQ-325:** Source applicability shall be assessed.
- **RSCH-REQ-326:** Source jurisdiction shall be assessed where material.
- **RSCH-REQ-327:** Source methodology shall be assessed.
- **RSCH-REQ-328:** Source limitations shall be recorded.
- **RSCH-REQ-329:** Conflicting sources shall remain visible.
- **RSCH-REQ-330:** Source disagreement shall not be resolved by selective omission.
- **RSCH-REQ-331:** Source records shall identify title.
- **RSCH-REQ-332:** Source records shall identify publisher or author.
- **RSCH-REQ-333:** Source records shall identify publication or revision date where available.
- **RSCH-REQ-334:** Source records shall identify access or verification date where material.
- **RSCH-REQ-335:** Source records shall identify stable reference.
- **RSCH-REQ-336:** Source records shall identify source type.
- **RSCH-REQ-337:** Source records shall identify relevance.
- **RSCH-REQ-338:** Source records shall identify claims supported.
- **RSCH-REQ-339:** Source records shall identify reliability limitations.
- **RSCH-REQ-340:** Source records shall identify license or use constraints where applicable.
- **RSCH-REQ-341:** Direct quotations shall remain limited and attributable.
- **RSCH-REQ-342:** Copyright restrictions shall be respected.
- **RSCH-REQ-343:** Unavailable sources shall not be fabricated.
- **RSCH-REQ-344:** Source gaps shall remain explicit.
- **RSCH-REQ-345:** Reference acquisition shall be completed before claims requiring those references are validated.

## 15. Reference State

- **RSCH-REQ-346:** Permitted reference states shall be `Not Started`, `Collecting`, `Sufficient for Planning`, `Sufficient for Validation`, `Conflicted`, `Stale` and `Unavailable`.
- **RSCH-REQ-347:** `Not Started` shall mean no governed source collection has begun.
- **RSCH-REQ-348:** `Collecting` shall mean source acquisition is active.
- **RSCH-REQ-349:** `Sufficient for Planning` shall mean sources support a safe validation plan but not a conclusion.
- **RSCH-REQ-350:** `Sufficient for Validation` shall mean sources meet defined evidence needs for the stated scope.
- **RSCH-REQ-351:** `Conflicted` shall mean material sources disagree.
- **RSCH-REQ-352:** `Stale` shall mean source currency is no longer adequate.
- **RSCH-REQ-353:** `Unavailable` shall mean required sources cannot be obtained lawfully or reliably.
- **RSCH-REQ-354:** Reference state shall not be inferred from source count alone.
- **RSCH-REQ-355:** Reference state shall identify evaluator.
- **RSCH-REQ-356:** Reference state shall identify evaluation date.
- **RSCH-REQ-357:** Reference sufficiency shall match the research question.
- **RSCH-REQ-358:** Reference sufficiency shall match claim strength.
- **RSCH-REQ-359:** Conflicted sources shall trigger explicit analysis.
- **RSCH-REQ-360:** Stale sources shall trigger refresh or limitation.
- **RSCH-REQ-361:** Unavailable sources shall not be replaced with invented facts.
- **RSCH-REQ-362:** Reference-state changes shall be attributable.
- **RSCH-REQ-363:** Reference-state history shall be retained.
- **RSCH-REQ-364:** Sufficient for Planning shall not be represented as validation.
- **RSCH-REQ-365:** Sufficient for Validation shall not guarantee a validated conclusion.

## 16. Hypotheses

- **RSCH-REQ-366:** Hypotheses shall be clearly labeled.
- **RSCH-REQ-367:** Hypotheses shall be falsifiable where technically feasible.
- **RSCH-REQ-368:** Hypotheses shall identify supporting observations.
- **RSCH-REQ-369:** Hypotheses shall identify contradicting observations.
- **RSCH-REQ-370:** Hypotheses shall identify uncertainty.
- **RSCH-REQ-371:** Hypotheses shall identify assumptions.
- **RSCH-REQ-372:** Hypotheses shall identify expected observations if supported.
- **RSCH-REQ-373:** Hypotheses shall identify expected observations if falsified.
- **RSCH-REQ-374:** Hypotheses shall identify alternative explanations.
- **RSCH-REQ-375:** Hypotheses shall not be represented as facts.
- **RSCH-REQ-376:** Hypotheses shall not become architecture directly.
- **RSCH-REQ-377:** Hypotheses shall not authorize implementation.
- **RSCH-REQ-378:** Hypotheses shall not authorize access.
- **RSCH-REQ-379:** Hypotheses shall not override risk controls.
- **RSCH-REQ-380:** Hypothesis changes shall preserve history.
- **RSCH-REQ-381:** Rejected hypotheses shall remain retrievable.
- **RSCH-REQ-382:** Partially supported hypotheses shall identify scope.
- **RSCH-REQ-383:** Inconclusive hypotheses shall remain inconclusive.
- **RSCH-REQ-384:** AI-generated hypotheses shall receive human review.
- **RSCH-REQ-385:** Hypothesis validation shall preserve raw evidence.
- **RSCH-REQ-386:** Hypothesis confidence shall not be confused with model confidence.
- **RSCH-REQ-387:** Hypothesis confidence shall not be represented as mathematical certainty without a justified method.
- **RSCH-REQ-388:** Hypotheses involving people shall avoid harmful profiling.
- **RSCH-REQ-389:** Hypotheses involving attribution shall require heightened evidence.
- **RSCH-REQ-390:** Hypothesis records shall link to the research item.

## 17. Assumptions Integration

- **RSCH-REQ-391:** Research items shall identify material assumptions.
- **RSCH-REQ-392:** Assumptions shall use identifiers governed by document 28 after consolidation.
- **RSCH-REQ-393:** Assumptions shall identify owner.
- **RSCH-REQ-394:** Assumptions shall identify consequence if false.
- **RSCH-REQ-395:** Assumptions shall identify validation state.
- **RSCH-REQ-396:** Assumptions shall identify expiry or review trigger where material.
- **RSCH-REQ-397:** Unvalidated assumptions shall remain visible.
- **RSCH-REQ-398:** Critical unvalidated assumptions shall block architecture adoption.
- **RSCH-REQ-399:** Research shall not silently convert assumptions into facts.
- **RSCH-REQ-400:** Research results shall identify which assumptions remain.
- **RSCH-REQ-401:** Assumption invalidation shall trigger research impact review.
- **RSCH-REQ-402:** Assumption confirmation shall identify evidence.
- **RSCH-REQ-403:** Assumption confirmation shall not exceed evidence scope.
- **RSCH-REQ-404:** Assumption conflicts shall be resolved or retained explicitly.
- **RSCH-REQ-405:** Assumption changes shall preserve history.
- **RSCH-REQ-406:** Research closure shall identify unresolved assumptions.
- **RSCH-REQ-407:** Public claims shall not hide unresolved assumptions.
- **RSCH-REQ-408:** AI systems shall not accept assumptions on behalf of human authority.
- **RSCH-REQ-409:** Assumption risk shall link to the risk register where material.
- **RSCH-REQ-410:** Document 28 remains a forward dependency for complete assumption governance.

## 18. Uncertainty and Confidence

- **RSCH-REQ-411:** Every material research conclusion shall identify uncertainty.
- **RSCH-REQ-412:** Uncertainty shall identify source limitations.
- **RSCH-REQ-413:** Uncertainty shall identify method limitations.
- **RSCH-REQ-414:** Uncertainty shall identify data limitations.
- **RSCH-REQ-415:** Uncertainty shall identify model limitations.
- **RSCH-REQ-416:** Uncertainty shall identify environment limitations.
- **RSCH-REQ-417:** Uncertainty shall identify temporal limitations.
- **RSCH-REQ-418:** Uncertainty shall identify jurisdictional limitations where material.
- **RSCH-REQ-419:** Uncertainty shall identify reproducibility limitations.
- **RSCH-REQ-420:** Confidence shall be justified.
- **RSCH-REQ-421:** Confidence shall not be inferred from document length.
- **RSCH-REQ-422:** Confidence shall not be inferred from source count alone.
- **RSCH-REQ-423:** Confidence shall not be inferred from model fluency.
- **RSCH-REQ-424:** Confidence shall not be confused with evidence integrity.
- **RSCH-REQ-425:** Confidence shall not be confused with authenticity.
- **RSCH-REQ-426:** Confidence shall not be confused with legal admissibility.
- **RSCH-REQ-427:** Confidence shall not be confused with compliance.
- **RSCH-REQ-428:** Qualitative confidence labels shall have defined meanings.
- **RSCH-REQ-429:** Quantitative confidence shall identify the statistical or analytical method.
- **RSCH-REQ-430:** Unknown uncertainty shall remain explicit.
- **RSCH-REQ-431:** High uncertainty shall constrain claims.
- **RSCH-REQ-432:** High uncertainty shall constrain architecture adoption.
- **RSCH-REQ-433:** Uncertainty reduction shall identify new evidence.
- **RSCH-REQ-434:** Uncertainty changes shall be attributable.
- **RSCH-REQ-435:** Residual uncertainty shall be included in closure records.

## 19. Validation Plan

- **RSCH-REQ-436:** Every Planned research item shall have a validation plan.
- **RSCH-REQ-437:** The validation plan shall identify the question.
- **RSCH-REQ-438:** The validation plan shall identify hypotheses.
- **RSCH-REQ-439:** The validation plan shall identify method.
- **RSCH-REQ-440:** The validation plan shall identify data.
- **RSCH-REQ-441:** The validation plan shall identify sources.
- **RSCH-REQ-442:** The validation plan shall identify environment.
- **RSCH-REQ-443:** The validation plan shall identify tools and versions.
- **RSCH-REQ-444:** The validation plan shall identify expected evidence.
- **RSCH-REQ-445:** The validation plan shall identify acceptance criteria.
- **RSCH-REQ-446:** The validation plan shall identify falsification criteria.
- **RSCH-REQ-447:** The validation plan shall identify inconclusive conditions.
- **RSCH-REQ-448:** The validation plan shall identify safety controls.
- **RSCH-REQ-449:** The validation plan shall identify privacy controls.
- **RSCH-REQ-450:** The validation plan shall identify ethical controls.
- **RSCH-REQ-451:** The validation plan shall identify authorization prerequisites.
- **RSCH-REQ-452:** The validation plan shall identify stop conditions.
- **RSCH-REQ-453:** The validation plan shall identify cleanup.
- **RSCH-REQ-454:** The validation plan shall identify retention.
- **RSCH-REQ-455:** The validation plan shall identify reproducibility requirements.
- **RSCH-REQ-456:** The validation plan shall identify reviewers.
- **RSCH-REQ-457:** The validation plan shall identify risks.
- **RSCH-REQ-458:** The validation plan shall identify assumptions.
- **RSCH-REQ-459:** The validation plan shall identify external dependencies.
- **RSCH-REQ-460:** The validation plan shall identify expected deliverables.
- **RSCH-REQ-461:** The validation plan shall identify decision use.
- **RSCH-REQ-462:** The validation plan shall identify publication limits.
- **RSCH-REQ-463:** The validation plan shall not require unauthorized systems.
- **RSCH-REQ-464:** The validation plan shall not require real criminal infrastructure.
- **RSCH-REQ-465:** The validation plan shall not require real-person public investigations.

## 20. Validation Execution

- **RSCH-REQ-466:** Research execution shall follow the approved validation plan.
- **RSCH-REQ-467:** Execution shall identify the exact plan version.
- **RSCH-REQ-468:** Execution shall identify responsible humans and workloads.
- **RSCH-REQ-469:** Execution shall identify start and completion times.
- **RSCH-REQ-470:** Execution shall identify environment.
- **RSCH-REQ-471:** Execution shall identify tools and versions.
- **RSCH-REQ-472:** Execution shall identify data and fixture versions.
- **RSCH-REQ-473:** Execution shall identify deviations.
- **RSCH-REQ-474:** Execution shall identify stop-condition events.
- **RSCH-REQ-475:** Execution shall preserve raw observations.
- **RSCH-REQ-476:** Execution shall preserve failed and negative results.
- **RSCH-REQ-477:** Execution shall preserve error and interruption records.
- **RSCH-REQ-478:** Execution shall not conceal unsafe deviations.
- **RSCH-REQ-479:** Execution shall not broaden scope silently.
- **RSCH-REQ-480:** Execution shall not use unapproved credentials.
- **RSCH-REQ-481:** Execution shall not use real case data without explicit authorization.
- **RSCH-REQ-482:** Execution shall not interact with uncontrolled criminal infrastructure.
- **RSCH-REQ-483:** Execution shall isolate high-risk content.
- **RSCH-REQ-484:** Execution shall maintain provenance.
- **RSCH-REQ-485:** Execution shall maintain auditability.
- **RSCH-REQ-486:** Execution changes shall be attributable.
- **RSCH-REQ-487:** Material method changes shall trigger plan review.
- **RSCH-REQ-488:** Aborted execution shall remain visible.
- **RSCH-REQ-489:** Invalid execution shall not support validation.
- **RSCH-REQ-490:** Execution evidence shall conform to `OBDIA-TEST-001` where test methods are used.

## 21. Validation Outcomes

- **RSCH-REQ-491:** Permitted research validation outcomes shall be `Supported`, `Partially Supported`, `Not Supported`, `Inconclusive` and `Invalid`.
- **RSCH-REQ-492:** `Supported` shall mean evidence satisfies the defined criteria for the stated scope.
- **RSCH-REQ-493:** `Partially Supported` shall identify supported and unsupported portions.
- **RSCH-REQ-494:** `Not Supported` shall mean evidence does not satisfy the hypothesis or proposed conclusion.
- **RSCH-REQ-495:** `Inconclusive` shall mean available evidence cannot support a reliable conclusion.
- **RSCH-REQ-496:** `Invalid` shall mean method, data or execution defects prevent reliable interpretation.
- **RSCH-REQ-497:** Validation outcome shall identify evaluator.
- **RSCH-REQ-498:** Validation outcome shall identify date.
- **RSCH-REQ-499:** Validation outcome shall identify evidence.
- **RSCH-REQ-500:** Validation outcome shall identify scope.
- **RSCH-REQ-501:** Validation outcome shall identify limitations.
- **RSCH-REQ-502:** Validation outcome shall identify unresolved uncertainty.
- **RSCH-REQ-503:** Validation outcome shall identify assumptions.
- **RSCH-REQ-504:** Validation outcome shall identify risk impact.
- **RSCH-REQ-505:** Validation outcome shall identify architecture implications.
- **RSCH-REQ-506:** Validation outcome shall identify whether an ADR is required.
- **RSCH-REQ-507:** Validation outcome shall identify whether change control is required.
- **RSCH-REQ-508:** Validation outcome shall not change architecture automatically.
- **RSCH-REQ-509:** Validation outcome shall not approve implementation automatically.
- **RSCH-REQ-510:** Validation outcome shall not accept risk automatically.
- **RSCH-REQ-511:** Validation outcome shall not establish compliance.
- **RSCH-REQ-512:** Negative outcomes shall remain retrievable.
- **RSCH-REQ-513:** Inconclusive outcomes shall not be converted to supported through wording.
- **RSCH-REQ-514:** Invalid outcomes shall preserve failure evidence.
- **RSCH-REQ-515:** Outcome changes shall be additive and attributable.

## 22. Expected Deliverables

- **RSCH-REQ-516:** Every Planned item shall identify expected deliverables.
- **RSCH-REQ-517:** Deliverables shall identify format.
- **RSCH-REQ-518:** Deliverables shall identify owner.
- **RSCH-REQ-519:** Deliverables shall identify intended audience.
- **RSCH-REQ-520:** Deliverables shall identify classification.
- **RSCH-REQ-521:** Deliverables shall identify acceptance criteria.
- **RSCH-REQ-522:** Deliverables shall identify authoritative source.
- **RSCH-REQ-523:** Deliverables shall identify dependencies.
- **RSCH-REQ-524:** Deliverables shall identify versioning.
- **RSCH-REQ-525:** Deliverables shall identify review requirements.
- **RSCH-REQ-526:** Deliverables shall identify publication eligibility.
- **RSCH-REQ-527:** Deliverables may include source analyses.
- **RSCH-REQ-528:** Deliverables may include experiment records.
- **RSCH-REQ-529:** Deliverables may include prototypes.
- **RSCH-REQ-530:** Deliverables may include simulations.
- **RSCH-REQ-531:** Deliverables may include diagrams.
- **RSCH-REQ-532:** Deliverables may include test plans and evidence.
- **RSCH-REQ-533:** Deliverables may include ADR proposals.
- **RSCH-REQ-534:** Deliverables may include requirement-change proposals.
- **RSCH-REQ-535:** Deliverables may include risk-register updates.
- **RSCH-REQ-536:** Deliverables may include assumption dispositions.
- **RSCH-REQ-537:** Deliverables shall distinguish proposal from implementation.
- **RSCH-REQ-538:** Deliverables shall not contain secrets.
- **RSCH-REQ-539:** Deliverables shall not contain unauthorized personal or case data.
- **RSCH-REQ-540:** Missing deliverables shall remain visible.

## 23. Architecture and Requirements Boundary

- **RSCH-REQ-541:** Research items shall not directly amend architecture.
- **RSCH-REQ-542:** Research items shall not directly amend requirements.
- **RSCH-REQ-543:** Research results shall identify proposed architecture implications.
- **RSCH-REQ-544:** Proposed architecture implications shall remain non-normative until governed adoption.
- **RSCH-REQ-545:** Architecture adoption shall require applicable ADRs.
- **RSCH-REQ-546:** Requirement adoption shall require change control.
- **RSCH-REQ-547:** Security-control adoption shall require Security Reviewer assessment.
- **RSCH-REQ-548:** Privacy-control adoption shall require Privacy and Governance Reviewer assessment.
- **RSCH-REQ-549:** Evidence-control adoption shall conform to `OBDIA-EVID-001`.
- **RSCH-REQ-550:** Authorization-control adoption shall conform to `OBDIA-AUTH-001`.
- **RSCH-REQ-551:** Agent-lifecycle changes shall conform to `OBDIA-AGENT-001`.
- **RSCH-REQ-552:** Connector changes shall conform to `OBDIA-CONN-001`.
- **RSCH-REQ-553:** Implementation changes shall conform to `OBDIA-CODE-001`.
- **RSCH-REQ-554:** Testing changes shall conform to `OBDIA-TEST-001`.
- **RSCH-REQ-555:** Diagram changes shall conform to `OBDIA-DIAG-001`.
- **RSCH-REQ-556:** Repository changes shall conform to `OBDIA-GH-001`.
- **RSCH-REQ-557:** Version changes shall conform to `OBDIA-VER-001`.
- **RSCH-REQ-558:** Risk changes shall conform to `OBDIA-RISK-001`.
- **RSCH-REQ-559:** Research evidence shall remain linked after adoption.
- **RSCH-REQ-560:** Adoption shall identify exact research versions.
- **RSCH-REQ-561:** Rejected adoption shall preserve rationale.
- **RSCH-REQ-562:** Partial adoption shall identify scope.
- **RSCH-REQ-563:** Architecture shall not cite a research question as if it were validation evidence.
- **RSCH-REQ-564:** Research backlog priority shall not control architecture priority automatically.
- **RSCH-REQ-565:** Conflicts with higher authority shall block adoption.

## 24. ADR and Decision Integration

- **RSCH-REQ-566:** Material architecture decisions derived from research shall use ADRs where applicable.
- **RSCH-REQ-567:** Research items shall identify related ADRs.
- **RSCH-REQ-568:** Research results shall identify proposed ADR needs.
- **RSCH-REQ-569:** ADRs shall cite exact research items and evidence.
- **RSCH-REQ-570:** ADRs shall not cite unsupported hypotheses as facts.
- **RSCH-REQ-571:** ADR acceptance shall remain distinct from research validation.
- **RSCH-REQ-572:** ADR rejection shall not invalidate underlying evidence automatically.
- **RSCH-REQ-573:** Research items shall not substitute for ADRs.
- **RSCH-REQ-574:** Decision-log entries shall not substitute for research evidence.
- **RSCH-REQ-575:** Document 29 shall govern the chronological decision index.
- **RSCH-REQ-576:** Research closure shall identify resulting ADRs.
- **RSCH-REQ-577:** Research closure shall identify rejected decision proposals.
- **RSCH-REQ-578:** Superseded ADRs shall trigger research impact review where material.
- **RSCH-REQ-579:** Decision changes shall not rewrite research history.
- **RSCH-REQ-580:** Research-derived decisions shall identify authority.
- **RSCH-REQ-581:** Research-derived decisions shall identify affected artifacts.
- **RSCH-REQ-582:** Research-derived decisions shall identify implementation and validation needs.
- **RSCH-REQ-583:** Research-derived decisions shall identify residual uncertainty.
- **RSCH-REQ-584:** Research-derived decisions shall identify risk impact.
- **RSCH-REQ-585:** Research-derived decisions shall preserve role-concentration disclosure.

## 25. Risk Integration

- **RSCH-REQ-586:** Research items shall link to affected risks.
- **RSCH-REQ-587:** Research findings shall identify new risks.
- **RSCH-REQ-588:** Research findings shall identify changed likelihood or impact evidence.
- **RSCH-REQ-589:** Research findings shall identify control-effectiveness evidence.
- **RSCH-REQ-590:** Research findings shall identify residual risk implications.
- **RSCH-REQ-591:** Research findings shall not accept risk.
- **RSCH-REQ-592:** An AI system shall not accept research-related risk.
- **RSCH-REQ-593:** Critical research risk shall block unsafe execution.
- **RSCH-REQ-594:** Canonical-prohibition risk shall not be accepted.
- **RSCH-REQ-595:** Research uncertainty shall inform risk confidence.
- **RSCH-REQ-596:** Risk-register changes shall use controlled records.
- **RSCH-REQ-597:** Research closure shall identify unresolved risks.
- **RSCH-REQ-598:** Negative results shall not be hidden because they increase risk.
- **RSCH-REQ-599:** Inconclusive results shall not reduce risk without justification.
- **RSCH-REQ-600:** Source uncertainty shall inform risk assessment.
- **RSCH-REQ-601:** Method limitations shall inform risk assessment.
- **RSCH-REQ-602:** Research environment limitations shall inform risk assessment.
- **RSCH-REQ-603:** Risk owners shall remain human.
- **RSCH-REQ-604:** Release-relevant research risks shall be reviewed before release.
- **RSCH-REQ-605:** Research-risk history shall remain retrievable.

## 26. Security and Dual-Use Boundary

- **RSCH-REQ-606:** Security research shall be defensive, simulated or otherwise explicitly authorized.
- **RSCH-REQ-607:** Research shall not perform unauthorized access.
- **RSCH-REQ-608:** Research shall not perform credential theft.
- **RSCH-REQ-609:** Research shall not deploy malware.
- **RSCH-REQ-610:** Research shall not conduct offensive cyber operations.
- **RSCH-REQ-611:** Research shall not conduct unlawful surveillance.
- **RSCH-REQ-612:** Research shall not conduct unlawful deanonymization.
- **RSCH-REQ-613:** Research shall not interact with uncontrolled criminal infrastructure.
- **RSCH-REQ-614:** Research shall not perform autonomous policing.
- **RSCH-REQ-615:** Research shall not make autonomous legal decisions.
- **RSCH-REQ-616:** Research shall not perform autonomous coercive actions.
- **RSCH-REQ-617:** Dual-use methods shall be redesigned into safe demonstrations where feasible.
- **RSCH-REQ-618:** Security tests shall use synthetic targets.
- **RSCH-REQ-619:** Network research shall use local labs or explicitly authorized sandboxes.
- **RSCH-REQ-620:** Dark Web research shall use operational isolation, simulation or archived-authorized material.
- **RSCH-REQ-621:** Deep Web research shall be treated according to authenticated-service constraints.
- **RSCH-REQ-622:** Exploit development shall not be required for public research artifacts.
- **RSCH-REQ-623:** Harmful payloads shall be minimized.
- **RSCH-REQ-624:** High-risk methods shall have stop conditions.
- **RSCH-REQ-625:** High-risk methods shall have Security Reviewer assessment.
- **RSCH-REQ-626:** Unexpected sensitive data shall trigger containment.
- **RSCH-REQ-627:** Unexpected credentials shall not be used.
- **RSCH-REQ-628:** Unexpected criminal content shall be quarantined according to policy.
- **RSCH-REQ-629:** Research logs shall not expose harmful operational details unnecessarily.
- **RSCH-REQ-630:** Public deliverables shall preserve defensive framing.

## 27. Privacy and Fundamental Rights

- **RSCH-REQ-631:** Research shall minimize personal data.
- **RSCH-REQ-632:** Research shall use synthetic identities by default.
- **RSCH-REQ-633:** Research shall avoid real case subjects in public artifacts.
- **RSCH-REQ-634:** Research shall avoid profiling real persons.
- **RSCH-REQ-635:** Research shall avoid biometric or sensitive-trait inference without explicit lawful authority.
- **RSCH-REQ-636:** Research shall preserve purpose limitation.
- **RSCH-REQ-637:** Research shall preserve case isolation.
- **RSCH-REQ-638:** Research shall assess data retention.
- **RSCH-REQ-639:** Research shall assess data deletion.
- **RSCH-REQ-640:** Research shall assess redaction.
- **RSCH-REQ-641:** Research shall assess cross-jurisdiction transfer where material.
- **RSCH-REQ-642:** Research shall assess human-review requirements.
- **RSCH-REQ-643:** Research shall assess potential discrimination.
- **RSCH-REQ-644:** Research shall assess chilling or surveillance effects.
- **RSCH-REQ-645:** Research shall assess mistaken attribution.
- **RSCH-REQ-646:** Research shall assess contestability and correction where material.
- **RSCH-REQ-647:** Research shall not treat public availability as unrestricted processing authority.
- **RSCH-REQ-648:** Privacy impact shall influence method selection.
- **RSCH-REQ-649:** Rights impact shall influence stop conditions.
- **RSCH-REQ-650:** Research evidence shall minimize personal data.
- **RSCH-REQ-651:** Research reports shall not identify real people unnecessarily.
- **RSCH-REQ-652:** Privacy incidents shall trigger containment and risk review.
- **RSCH-REQ-653:** Privacy limitations shall remain visible.
- **RSCH-REQ-654:** Privacy and Governance Reviewer assessment shall be required for material rights impact.
- **RSCH-REQ-655:** Public claims shall not overstate anonymization.

## 28. Reproducibility and Testing

- **RSCH-REQ-656:** Research validation shall be reproducible where technically feasible.
- **RSCH-REQ-657:** Reproduction instructions shall identify source commit.
- **RSCH-REQ-658:** Reproduction instructions shall identify artifact versions.
- **RSCH-REQ-659:** Reproduction instructions shall identify tools and versions.
- **RSCH-REQ-660:** Reproduction instructions shall identify environment.
- **RSCH-REQ-661:** Reproduction instructions shall identify data and fixture versions.
- **RSCH-REQ-662:** Reproduction instructions shall identify random seeds where material.
- **RSCH-REQ-663:** Reproduction instructions shall identify model versions where available.
- **RSCH-REQ-664:** Reproduction instructions shall identify prompt versions where material.
- **RSCH-REQ-665:** Reproduction instructions shall identify external-service dependencies.
- **RSCH-REQ-666:** Reproduction instructions shall identify expected outputs.
- **RSCH-REQ-667:** Reproduction instructions shall distinguish commands from expected outputs.
- **RSCH-REQ-668:** Reproduction attempts shall create new records.
- **RSCH-REQ-669:** Reproduction failures shall remain visible.
- **RSCH-REQ-670:** Model nondeterminism shall be documented.
- **RSCH-REQ-671:** External-service nondeterminism shall be documented.
- **RSCH-REQ-672:** Equivalent results shall use defined equivalence criteria.
- **RSCH-REQ-673:** Reproducibility shall not rely on production credentials.
- **RSCH-REQ-674:** Reproducibility shall not rely on unauthorized systems.
- **RSCH-REQ-675:** Research tests shall conform to `OBDIA-TEST-001`.
- **RSCH-REQ-676:** Failed tests shall remain visible.
- **RSCH-REQ-677:** Skipped tests shall remain visible.
- **RSCH-REQ-678:** Flaky tests shall not count as reliable validation.
- **RSCH-REQ-679:** Passing tests shall not generalize beyond scope.
- **RSCH-REQ-680:** Automated testing shall not replace Research Reviewer judgment.

## 29. Research Environments

- **RSCH-REQ-681:** Research environments shall be explicitly identified.
- **RSCH-REQ-682:** Research environments shall have owners.
- **RSCH-REQ-683:** Research environments shall use least privilege.
- **RSCH-REQ-684:** Research environments shall use non-production credentials.
- **RSCH-REQ-685:** Research environments shall use synthetic or authorized data.
- **RSCH-REQ-686:** Research environments shall restrict network egress by default.
- **RSCH-REQ-687:** Research environments shall separate development, test and demonstration contexts.
- **RSCH-REQ-688:** High-risk research environments shall be isolated.
- **RSCH-REQ-689:** Dark Web research environments shall be operationally isolated.
- **RSCH-REQ-690:** Isolated Dark Web environments shall not interact with uncontrolled criminal infrastructure.
- **RSCH-REQ-691:** Web3 research shall use local chains or testnets by default.
- **RSCH-REQ-692:** Cloud research shall use sandbox accounts.
- **RSCH-REQ-693:** Connector research shall use mocks or authorized sandboxes.
- **RSCH-REQ-694:** Evidence research shall use synthetic evidence.
- **RSCH-REQ-695:** Research environments shall identify monitoring.
- **RSCH-REQ-696:** Research environments shall identify logging.
- **RSCH-REQ-697:** Research environments shall identify cleanup.
- **RSCH-REQ-698:** Research environments shall identify retention.
- **RSCH-REQ-699:** Environment drift shall be detectable where feasible.
- **RSCH-REQ-700:** Environment changes shall trigger reproducibility review.
- **RSCH-REQ-701:** Environment failures shall not create false positive results.
- **RSCH-REQ-702:** Shared environment interference shall be controlled.
- **RSCH-REQ-703:** Environment snapshots shall not contain secrets.
- **RSCH-REQ-704:** Environment access shall be authorized.
- **RSCH-REQ-705:** Environment limitations shall remain visible.

## 30. External-System and Provider Research

- **RSCH-REQ-706:** Research involving external systems shall identify the provider or owner.
- **RSCH-REQ-707:** External-system access shall be explicitly authorized.
- **RSCH-REQ-708:** External-system terms and technical constraints shall be assessed.
- **RSCH-REQ-709:** External-system versions shall be recorded where available.
- **RSCH-REQ-710:** Provider aliases shall not replace exact version evidence where exact versions exist.
- **RSCH-REQ-711:** Provider silent-change risk shall be recorded.
- **RSCH-REQ-712:** Provider availability shall not be treated as guaranteed.
- **RSCH-REQ-713:** Provider output shall be treated as untrusted.
- **RSCH-REQ-714:** Provider documentation shall not replace controlled validation.
- **RSCH-REQ-715:** Provider marketing claims shall not be treated as independent evidence.
- **RSCH-REQ-716:** External-system data retention shall be assessed.
- **RSCH-REQ-717:** External-system model-training or reuse shall be assessed.
- **RSCH-REQ-718:** External-system jurisdiction shall be assessed where material.
- **RSCH-REQ-719:** External-system credentials shall be dedicated and least-privileged.
- **RSCH-REQ-720:** External-system failures shall not broaden access.
- **RSCH-REQ-721:** External-system fallback shall be pre-approved.
- **RSCH-REQ-722:** External-system research shall preserve provenance.
- **RSCH-REQ-723:** External-system cost and quota shall be bounded.
- **RSCH-REQ-724:** External-system incidents shall trigger containment.
- **RSCH-REQ-725:** External-system limitations shall remain visible.

## 31. Domain-Specific Research Controls

- **RSCH-REQ-726:** Identity research shall preserve distinctions among human, agent, workload and connector identities.
- **RSCH-REQ-727:** Binding research shall preserve one-agent-to-one-officer accountability.
- **RSCH-REQ-728:** Credential research shall prohibit transfer of reusable human credentials.
- **RSCH-REQ-729:** Authorization research shall default to deny.
- **RSCH-REQ-730:** Authorization research shall distinguish authentication from authorization.
- **RSCH-REQ-731:** Authorization research shall preserve case, purpose, jurisdiction, time, tool and data scope.
- **RSCH-REQ-732:** Agent-lifecycle research shall preserve mandatory `Bind` before authorization.
- **RSCH-REQ-733:** Agent-lifecycle research shall preserve terminal revocation.
- **RSCH-REQ-734:** Evidence research shall preserve original-versus-derived distinctions.
- **RSCH-REQ-735:** Evidence research shall preserve provenance and chain of custody.
- **RSCH-REQ-736:** Evidence research shall distinguish integrity from authenticity and admissibility.
- **RSCH-REQ-737:** Connector research shall preserve endpoint and egress controls.
- **RSCH-REQ-738:** Connector research shall treat external content as untrusted.
- **RSCH-REQ-739:** Connector research shall preserve revocation and failure isolation.
- **RSCH-REQ-740:** AI research shall treat model output as untrusted.
- **RSCH-REQ-741:** AI research shall preserve human accountability.
- **RSCH-REQ-742:** AI research shall test indirect prompt injection.
- **RSCH-REQ-743:** Cloud research shall preserve tenant and account isolation.
- **RSCH-REQ-744:** Web research shall distinguish surface, authenticated deep-web and isolated Dark Web contexts.
- **RSCH-REQ-745:** Web3 research shall distinguish read-only analysis from transaction submission.
- **RSCH-REQ-746:** Smart-contract research shall use simulation or testnets by default.
- **RSCH-REQ-747:** Repository research shall preserve change and review evidence.
- **RSCH-REQ-748:** Supply-chain research shall preserve dependency provenance.
- **RSCH-REQ-749:** Release research shall preserve safe-publication boundaries.
- **RSCH-REQ-750:** Compliance research shall not present mappings as proof of compliance.

## 32. Review Cadence and Triggers

- **RSCH-REQ-751:** Active research items shall have review dates.
- **RSCH-REQ-752:** Review cadence shall reflect priority.
- **RSCH-REQ-753:** P0 items shall receive frequent review until resolved or safely blocked.
- **RSCH-REQ-754:** P1 items shall receive scheduled near-term review.
- **RSCH-REQ-755:** P2 items shall receive regular review.
- **RSCH-REQ-756:** P3 and P4 items shall receive staleness review.
- **RSCH-REQ-757:** A material threat change shall trigger research review.
- **RSCH-REQ-758:** A material risk change shall trigger research review.
- **RSCH-REQ-759:** A material assumption change shall trigger research review.
- **RSCH-REQ-760:** A material architecture change shall trigger research review.
- **RSCH-REQ-761:** A material provider or model change shall trigger research review.
- **RSCH-REQ-762:** A material legal or standards change shall trigger research review where applicable.
- **RSCH-REQ-763:** A security incident shall trigger affected research review.
- **RSCH-REQ-764:** A privacy incident shall trigger affected research review.
- **RSCH-REQ-765:** A failed validation shall trigger review.
- **RSCH-REQ-766:** A new blocking dependency shall trigger review.
- **RSCH-REQ-767:** Source staleness shall trigger review.
- **RSCH-REQ-768:** Release planning shall trigger review of release-relevant items.
- **RSCH-REQ-769:** Review outcomes shall be attributable.
- **RSCH-REQ-770:** Missed review dates shall remain visible.

## 33. Staleness and Aging

- **RSCH-REQ-771:** Backlog age shall be measured from created date.
- **RSCH-REQ-772:** Active age shall be measured from the most recent substantive action.
- **RSCH-REQ-773:** Age shall not determine priority alone.
- **RSCH-REQ-774:** Stale items shall be identified.
- **RSCH-REQ-775:** Staleness criteria shall consider source currency.
- **RSCH-REQ-776:** Staleness criteria shall consider architecture changes.
- **RSCH-REQ-777:** Staleness criteria shall consider provider changes.
- **RSCH-REQ-778:** Staleness criteria shall consider threat changes.
- **RSCH-REQ-779:** Staleness criteria shall consider unresolved ownership.
- **RSCH-REQ-780:** Stale items shall receive triage.
- **RSCH-REQ-781:** Stale items may be refreshed.
- **RSCH-REQ-782:** Stale items may be superseded.
- **RSCH-REQ-783:** Stale items may be rejected or closed with rationale.
- **RSCH-REQ-784:** Stale items shall not be silently deleted.
- **RSCH-REQ-785:** Old high-risk items shall not be closed solely by age.
- **RSCH-REQ-786:** Age metrics shall not incentivize unsafe closure.
- **RSCH-REQ-787:** Aging reports shall preserve priority and owner context.
- **RSCH-REQ-788:** Archived items shall remain retrievable.
- **RSCH-REQ-789:** Staleness decisions shall be attributable.
- **RSCH-REQ-790:** Staleness history shall remain visible.

## 34. Closure and Archival

- **RSCH-REQ-791:** Research closure shall require a disposition.
- **RSCH-REQ-792:** Closure shall identify the final status.
- **RSCH-REQ-793:** Closure shall identify the final research outcome.
- **RSCH-REQ-794:** Closure shall identify evidence.
- **RSCH-REQ-795:** Closure shall identify unresolved uncertainty.
- **RSCH-REQ-796:** Closure shall identify remaining assumptions.
- **RSCH-REQ-797:** Closure shall identify risk impact.
- **RSCH-REQ-798:** Closure shall identify affected architecture.
- **RSCH-REQ-799:** Closure shall identify resulting ADRs.
- **RSCH-REQ-800:** Closure shall identify resulting change records.
- **RSCH-REQ-801:** Closure shall identify resulting tests.
- **RSCH-REQ-802:** Closure shall identify resulting publications.
- **RSCH-REQ-803:** Closure shall identify owner and authority.
- **RSCH-REQ-804:** Closure shall identify date.
- **RSCH-REQ-805:** Closure shall identify supersession where applicable.
- **RSCH-REQ-806:** Closure shall not fabricate validation.
- **RSCH-REQ-807:** Closure shall not hide negative results.
- **RSCH-REQ-808:** Closed items shall remain retrievable.
- **RSCH-REQ-809:** Closed items shall be read-only except for additive correction metadata.
- **RSCH-REQ-810:** Archived items shall preserve evidence references.
- **RSCH-REQ-811:** Archived items shall preserve source references.
- **RSCH-REQ-812:** Archived items shall preserve versions.
- **RSCH-REQ-813:** Archive migration shall preserve identifiers.
- **RSCH-REQ-814:** Archive restoration shall not reactivate status silently.
- **RSCH-REQ-815:** Closure and archival history shall remain auditable.

## 35. Metrics and Reporting

- **RSCH-REQ-816:** Backlog reporting shall identify total active items.
- **RSCH-REQ-817:** Backlog reporting shall identify items by priority.
- **RSCH-REQ-818:** Backlog reporting shall identify items by status.
- **RSCH-REQ-819:** Backlog reporting shall identify blocked items.
- **RSCH-REQ-820:** Backlog reporting shall identify overdue reviews.
- **RSCH-REQ-821:** Backlog reporting shall identify orphaned items.
- **RSCH-REQ-822:** Backlog reporting shall identify source-state gaps.
- **RSCH-REQ-823:** Backlog reporting shall identify validation outcomes.
- **RSCH-REQ-824:** Backlog reporting shall identify unresolved critical risks.
- **RSCH-REQ-825:** Backlog reporting shall identify assumption dependencies.
- **RSCH-REQ-826:** Backlog reporting shall identify decision dependencies.
- **RSCH-REQ-827:** Metrics shall not conceal individual critical items.
- **RSCH-REQ-828:** Metrics shall not convert quantity into quality claims.
- **RSCH-REQ-829:** Closure rate shall not incentivize premature closure.
- **RSCH-REQ-830:** Priority metrics shall preserve rationale.
- **RSCH-REQ-831:** AI-generated summaries shall be verified.
- **RSCH-REQ-832:** Dashboards shall identify freshness.
- **RSCH-REQ-833:** Dashboard errors shall not overwrite authoritative records.
- **RSCH-REQ-834:** Reports shall minimize personal data.
- **RSCH-REQ-835:** Raw records shall remain available to authorized reviewers.

## 36. Publication and Portfolio Use

- **RSCH-REQ-836:** Public research outputs shall comply with documents 04 and 15.
- **RSCH-REQ-837:** Public outputs shall distinguish verified facts from proposed architecture.
- **RSCH-REQ-838:** Public outputs shall distinguish hypotheses from conclusions.
- **RSCH-REQ-839:** Public outputs shall distinguish simulation from implementation.
- **RSCH-REQ-840:** Public outputs shall distinguish testing from validation.
- **RSCH-REQ-841:** Public outputs shall identify limitations.
- **RSCH-REQ-842:** Public outputs shall identify synthetic data.
- **RSCH-REQ-843:** Public outputs shall not expose secrets.
- **RSCH-REQ-844:** Public outputs shall not expose real case data.
- **RSCH-REQ-845:** Public outputs shall not expose real officer credentials.
- **RSCH-REQ-846:** Public outputs shall not enable harmful operational misuse unnecessarily.
- **RSCH-REQ-847:** Public outputs shall not claim institutional adoption without evidence.
- **RSCH-REQ-848:** Public outputs shall not claim production readiness without evidence.
- **RSCH-REQ-849:** Public outputs shall not claim certification or compliance without evidence.
- **RSCH-REQ-850:** Public Dark Web research shall use simulation or authorized archival framing.
- **RSCH-REQ-851:** Public Web3 research shall identify testnet or read-only scope.
- **RSCH-REQ-852:** Public source citations shall remain accurate.
- **RSCH-REQ-853:** Publication shall require Release Reviewer assessment.
- **RSCH-REQ-854:** Publication corrections shall preserve history.
- **RSCH-REQ-855:** Portfolio claims shall not exceed research evidence.

## 37. Exceptions and Non-Applicability

- **RSCH-REQ-856:** Research exceptions shall be explicit.
- **RSCH-REQ-857:** Research exceptions shall be narrow.
- **RSCH-REQ-858:** Research exceptions shall identify affected requirements.
- **RSCH-REQ-859:** Research exceptions shall identify rationale.
- **RSCH-REQ-860:** Research exceptions shall identify owner.
- **RSCH-REQ-861:** Research exceptions shall identify risk.
- **RSCH-REQ-862:** Research exceptions shall identify compensating controls.
- **RSCH-REQ-863:** Research exceptions shall identify expiry.
- **RSCH-REQ-864:** Research exceptions shall identify validation impact.
- **RSCH-REQ-865:** Research exceptions shall not authorize canonical prohibitions.
- **RSCH-REQ-866:** Research exceptions shall not authorize unauthorized access.
- **RSCH-REQ-867:** Research exceptions shall not authorize uncontrolled criminal-infrastructure interaction.
- **RSCH-REQ-868:** Research exceptions shall not convert unsupported claims into facts.
- **RSCH-REQ-869:** Non-applicability shall have technical rationale.
- **RSCH-REQ-870:** Non-applicability shall identify reviewer.
- **RSCH-REQ-871:** Expired exceptions shall block progression or require renewed review.
- **RSCH-REQ-872:** Repeated exceptions shall trigger policy review.
- **RSCH-REQ-873:** Exception closure shall preserve history.
- **RSCH-REQ-874:** An AI system shall not approve exceptions.
- **RSCH-REQ-875:** Document 28 remains the forward authority for detailed exception-related assumption impacts where applicable.

## 38. Traceability

- **RSCH-REQ-876:** Every research item shall trace to its intake record.
- **RSCH-REQ-877:** Every research item shall trace to its owner.
- **RSCH-REQ-878:** Every research item shall trace to dependencies.
- **RSCH-REQ-879:** Every research item shall trace to sources.
- **RSCH-REQ-880:** Every research item shall trace to assumptions.
- **RSCH-REQ-881:** Every research item shall trace to risks.
- **RSCH-REQ-882:** Every research item shall trace to threats where applicable.
- **RSCH-REQ-883:** Every hypothesis shall trace to a research item.
- **RSCH-REQ-884:** Every validation plan shall trace to a research item.
- **RSCH-REQ-885:** Every experiment shall trace to a validation plan.
- **RSCH-REQ-886:** Every evidence package shall trace to execution.
- **RSCH-REQ-887:** Every validation outcome shall trace to evidence.
- **RSCH-REQ-888:** Every architecture proposal shall trace to validated research.
- **RSCH-REQ-889:** Every ADR derived from research shall trace to exact items and evidence.
- **RSCH-REQ-890:** Every requirement change derived from research shall trace to change control.
- **RSCH-REQ-891:** Every risk update derived from research shall trace to evidence.
- **RSCH-REQ-892:** Every assumption disposition shall trace to research.
- **RSCH-REQ-893:** Every publication shall trace to reviewed deliverables.
- **RSCH-REQ-894:** Every superseded item shall trace to its successor.
- **RSCH-REQ-895:** Every closed item shall trace to closure authority.
- **RSCH-REQ-896:** Traceability shall be bidirectional.
- **RSCH-REQ-897:** Broken traceability affecting safety, authority, evidence or publication shall be blocking.
- **RSCH-REQ-898:** Traceability records shall not contain secrets.
- **RSCH-REQ-899:** Traceability records shall minimize personal and case data.
- **RSCH-REQ-900:** Baseline freeze shall validate research traceability across documents 01–30.

## 39. Domain-Specific Research Planning Controls

The following controls ensure that each material OBDIA research domain is planned with explicit authority, safety, evidence and adoption boundaries.

- **RSCH-REQ-901:** Research concerning officer-agent binding shall identify a precise research question concerning the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-902:** Research concerning officer-agent binding shall identify the governing documents for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-903:** Research concerning officer-agent binding shall identify the threat and risk links for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-904:** Research concerning officer-agent binding shall identify assumptions that materially affect the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-905:** Research concerning officer-agent binding shall identify safe and authorized validation methods for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-906:** Research concerning officer-agent binding shall identify prohibited methods when studying the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-907:** Research concerning officer-agent binding shall identify required synthetic data or fixtures for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-908:** Research concerning officer-agent binding shall identify exact environments and dependencies for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-909:** Research concerning officer-agent binding shall identify expected evidence and acceptance criteria for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-910:** Research concerning officer-agent binding shall identify reproducibility limitations for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-911:** Research concerning officer-agent binding shall identify privacy and fundamental-rights implications of the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-912:** Research concerning officer-agent binding shall identify human accountability and approval boundaries for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-913:** Research concerning officer-agent binding shall identify failure, misuse and abuse cases involving the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-914:** Research concerning officer-agent binding shall identify implementation and change-control boundaries for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-915:** Research concerning officer-agent binding shall identify publication and portfolio limitations for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-916:** Research concerning officer-agent binding shall identify closure and archival evidence for the unique accountable human investigator and institutional issuer.
- **RSCH-REQ-917:** Research concerning machine identity shall identify a precise research question concerning agent, workload and connector identity provenance.
- **RSCH-REQ-918:** Research concerning machine identity shall identify the governing documents for agent, workload and connector identity provenance.
- **RSCH-REQ-919:** Research concerning machine identity shall identify the threat and risk links for agent, workload and connector identity provenance.
- **RSCH-REQ-920:** Research concerning machine identity shall identify assumptions that materially affect agent, workload and connector identity provenance.
- **RSCH-REQ-921:** Research concerning machine identity shall identify safe and authorized validation methods for agent, workload and connector identity provenance.
- **RSCH-REQ-922:** Research concerning machine identity shall identify prohibited methods when studying agent, workload and connector identity provenance.
- **RSCH-REQ-923:** Research concerning machine identity shall identify required synthetic data or fixtures for agent, workload and connector identity provenance.
- **RSCH-REQ-924:** Research concerning machine identity shall identify exact environments and dependencies for agent, workload and connector identity provenance.
- **RSCH-REQ-925:** Research concerning machine identity shall identify expected evidence and acceptance criteria for agent, workload and connector identity provenance.
- **RSCH-REQ-926:** Research concerning machine identity shall identify reproducibility limitations for agent, workload and connector identity provenance.
- **RSCH-REQ-927:** Research concerning machine identity shall identify privacy and fundamental-rights implications of agent, workload and connector identity provenance.
- **RSCH-REQ-928:** Research concerning machine identity shall identify human accountability and approval boundaries for agent, workload and connector identity provenance.
- **RSCH-REQ-929:** Research concerning machine identity shall identify failure, misuse and abuse cases involving agent, workload and connector identity provenance.
- **RSCH-REQ-930:** Research concerning machine identity shall identify implementation and change-control boundaries for agent, workload and connector identity provenance.
- **RSCH-REQ-931:** Research concerning machine identity shall identify publication and portfolio limitations for agent, workload and connector identity provenance.
- **RSCH-REQ-932:** Research concerning machine identity shall identify closure and archival evidence for agent, workload and connector identity provenance.
- **RSCH-REQ-933:** Research concerning authorization shall identify a precise research question concerning case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-934:** Research concerning authorization shall identify the governing documents for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-935:** Research concerning authorization shall identify the threat and risk links for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-936:** Research concerning authorization shall identify assumptions that materially affect case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-937:** Research concerning authorization shall identify safe and authorized validation methods for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-938:** Research concerning authorization shall identify prohibited methods when studying case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-939:** Research concerning authorization shall identify required synthetic data or fixtures for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-940:** Research concerning authorization shall identify exact environments and dependencies for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-941:** Research concerning authorization shall identify expected evidence and acceptance criteria for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-942:** Research concerning authorization shall identify reproducibility limitations for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-943:** Research concerning authorization shall identify privacy and fundamental-rights implications of case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-944:** Research concerning authorization shall identify human accountability and approval boundaries for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-945:** Research concerning authorization shall identify failure, misuse and abuse cases involving case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-946:** Research concerning authorization shall identify implementation and change-control boundaries for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-947:** Research concerning authorization shall identify publication and portfolio limitations for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-948:** Research concerning authorization shall identify closure and archival evidence for case, purpose, jurisdiction, time, tool and data constraints.
- **RSCH-REQ-949:** Research concerning agent lifecycle shall identify a precise research question concerning provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-950:** Research concerning agent lifecycle shall identify the governing documents for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-951:** Research concerning agent lifecycle shall identify the threat and risk links for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-952:** Research concerning agent lifecycle shall identify assumptions that materially affect provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-953:** Research concerning agent lifecycle shall identify safe and authorized validation methods for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-954:** Research concerning agent lifecycle shall identify prohibited methods when studying provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-955:** Research concerning agent lifecycle shall identify required synthetic data or fixtures for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-956:** Research concerning agent lifecycle shall identify exact environments and dependencies for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-957:** Research concerning agent lifecycle shall identify expected evidence and acceptance criteria for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-958:** Research concerning agent lifecycle shall identify reproducibility limitations for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-959:** Research concerning agent lifecycle shall identify privacy and fundamental-rights implications of provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-960:** Research concerning agent lifecycle shall identify human accountability and approval boundaries for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-961:** Research concerning agent lifecycle shall identify failure, misuse and abuse cases involving provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-962:** Research concerning agent lifecycle shall identify implementation and change-control boundaries for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-963:** Research concerning agent lifecycle shall identify publication and portfolio limitations for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-964:** Research concerning agent lifecycle shall identify closure and archival evidence for provisioning, binding, activation, suspension, revocation and archival.
- **RSCH-REQ-965:** Research concerning evidence integrity shall identify a precise research question concerning hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-966:** Research concerning evidence integrity shall identify the governing documents for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-967:** Research concerning evidence integrity shall identify the threat and risk links for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-968:** Research concerning evidence integrity shall identify assumptions that materially affect hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-969:** Research concerning evidence integrity shall identify safe and authorized validation methods for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-970:** Research concerning evidence integrity shall identify prohibited methods when studying hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-971:** Research concerning evidence integrity shall identify required synthetic data or fixtures for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-972:** Research concerning evidence integrity shall identify exact environments and dependencies for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-973:** Research concerning evidence integrity shall identify expected evidence and acceptance criteria for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-974:** Research concerning evidence integrity shall identify reproducibility limitations for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-975:** Research concerning evidence integrity shall identify privacy and fundamental-rights implications of hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-976:** Research concerning evidence integrity shall identify human accountability and approval boundaries for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-977:** Research concerning evidence integrity shall identify failure, misuse and abuse cases involving hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-978:** Research concerning evidence integrity shall identify implementation and change-control boundaries for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-979:** Research concerning evidence integrity shall identify publication and portfolio limitations for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-980:** Research concerning evidence integrity shall identify closure and archival evidence for hashing, provenance, custody and derived-artifact lineage.
- **RSCH-REQ-981:** Research concerning connector security shall identify a precise research question concerning authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-982:** Research concerning connector security shall identify the governing documents for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-983:** Research concerning connector security shall identify the threat and risk links for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-984:** Research concerning connector security shall identify assumptions that materially affect authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-985:** Research concerning connector security shall identify safe and authorized validation methods for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-986:** Research concerning connector security shall identify prohibited methods when studying authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-987:** Research concerning connector security shall identify required synthetic data or fixtures for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-988:** Research concerning connector security shall identify exact environments and dependencies for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-989:** Research concerning connector security shall identify expected evidence and acceptance criteria for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-990:** Research concerning connector security shall identify reproducibility limitations for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-991:** Research concerning connector security shall identify privacy and fundamental-rights implications of authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-992:** Research concerning connector security shall identify human accountability and approval boundaries for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-993:** Research concerning connector security shall identify failure, misuse and abuse cases involving authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-994:** Research concerning connector security shall identify implementation and change-control boundaries for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-995:** Research concerning connector security shall identify publication and portfolio limitations for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-996:** Research concerning connector security shall identify closure and archival evidence for authentication, egress, validation, provenance, isolation and revocation.
- **RSCH-REQ-997:** Research concerning prompt injection shall identify a precise research question concerning untrusted retrieved content and tool-mediated actions.
- **RSCH-REQ-998:** Research concerning prompt injection shall identify the governing documents for untrusted retrieved content and tool-mediated actions.
- **RSCH-REQ-999:** Research concerning prompt injection shall identify the threat and risk links for untrusted retrieved content and tool-mediated actions.

## 40. Initial Synthetic Research Backlog

The following records are initial research inputs. They are not validated conclusions, approved architecture, implementation commitments or risk acceptances.


### `RSCH-ITEM-001` — Cryptographic officer-agent binding assurance

| Field | Value |
|---|---|
| **Research Question** | Which cryptographic binding pattern best proves that one institutionally issued agent identity is bound to exactly one identified officer without transferring human credentials? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P0` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-ID-001; OBDIA-AUTH-001; OBDIA-AGENT-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Architecture options analysis, threat analysis, prototype plan and ADR proposal. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-002` — Authorization-context freshness

| Field | Value |
|---|---|
| **Research Question** | What freshness and re-evaluation controls are required for officer, case, purpose, jurisdiction, time, tool and data attributes during long-running tasks? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P0` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-AUTH-001; OBDIA-SEC-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Context-freshness model, failure cases and test plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-003` — Revocation propagation latency

| Field | Value |
|---|---|
| **Research Question** | What measurable upper bounds are technically defensible for revocation propagation across identity, policy, runtime, workload and connector enforcement points? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P0` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-ID-001; OBDIA-AUTH-001; OBDIA-AGENT-001; OBDIA-CONN-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Latency model, controlled fault-injection plan and acceptance criteria. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-004` — Distributed lifecycle consistency

| Field | Value |
|---|---|
| **Research Question** | How should restrictive lifecycle transitions win during partitions, retries and concurrent activation, suspension or revocation events? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-AGENT-001; OBDIA-SEC-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | State-consistency analysis, transition invariants and simulation plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-005` — Evidence canonicalization

| Field | Value |
|---|---|
| **Research Question** | Which canonicalization rules are required before hashing structured evidence metadata without obscuring original bytes or source representation? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-EVID-001; OBDIA-CODE-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Canonicalization profile, ambiguity analysis and test vectors. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-006` — Derived-evidence lineage

| Field | Value |
|---|---|
| **Research Question** | How can multi-step AI and non-AI transformations preserve complete parent-child lineage while keeping original evidence immutable? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-EVID-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Lineage model, transformation-event schema and validation plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-007` — Tamper-evident custody events

| Field | Value |
|---|---|
| **Research Question** | Which append-only or cryptographically linked event approach is proportionate for synthetic chain-of-custody demonstrations? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-EVID-001; OBDIA-SEC-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Options assessment, threat model and prototype plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-008` — Indirect prompt-injection containment

| Field | Value |
|---|---|
| **Research Question** | Which layered controls best prevent connector-returned content from changing policy, obtaining secrets or triggering unauthorized tool actions? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P0` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-CONN-001; OBDIA-CODE-001; OBDIA-TEST-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Attack taxonomy, defensive fixtures, control comparison and negative tests. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-009` — Connector failure isolation

| Field | Value |
|---|---|
| **Research Question** | What connector-process, credential, network and queue isolation boundaries minimize cross-case and cross-agent blast radius? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-CONN-001; OBDIA-SEC-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Isolation architecture, failure matrix and test plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-010` — Model uncertainty communication

| Field | Value |
|---|---|
| **Research Question** | How should model uncertainty be represented without conflating confidence, integrity, authenticity, admissibility or factual truth? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-EVID-001; OBDIA-RES-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Terminology proposal, interface examples and validation study. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-011` — Human-approval thresholds

| Field | Value |
|---|---|
| **Research Question** | Which actions require mandatory human approval, step-up authentication or separation of duties in a safe research demonstration? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P0` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-AUTH-001; OBDIA-GOV-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Action taxonomy, approval matrix and abuse-case review. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-012` — Purpose-limitation enforcement

| Field | Value |
|---|---|
| **Research Question** | Which policy and data-model patterns most reliably prevent cross-purpose reuse of collected or derived information? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-AUTH-001; OBDIA-EVID-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Policy patterns, negative tests and residual-risk analysis. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-013` — Cross-case isolation

| Field | Value |
|---|---|
| **Research Question** | How should storage, memory, caches, logs and connectors enforce case isolation under concurrent research workflows? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P0` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-AUTH-001; OBDIA-EVID-001; OBDIA-CONN-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Isolation model, test fixtures and failure-injection plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-014` — Jurisdiction representation

| Field | Value |
|---|---|
| **Research Question** | How can jurisdiction constraints be represented as explicit policy attributes without pretending to encode legal conclusions autonomously? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-AUTH-001; OBDIA-COMP-001 forward dependency |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Attribute model, limitation statement and governance review. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-015` — Privacy-preserving telemetry

| Field | Value |
|---|---|
| **Research Question** | What telemetry is minimally necessary for accountability, security monitoring and incident response without excessive case or personal data? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-SEC-001; OBDIA-EVID-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Telemetry schema, minimization analysis and retention proposal. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-016` — Model and provider drift

| Field | Value |
|---|---|
| **Research Question** | How should silent model, API and provider behavior changes invalidate or qualify prior validation evidence? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-TEST-001; OBDIA-VER-001; OBDIA-RISK-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Drift-detection model, evidence invalidation rules and test plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-017` — Reproducible AI-assisted analysis

| Field | Value |
|---|---|
| **Research Question** | Which inputs, prompts, model identifiers, settings and evidence are necessary to reproduce or appropriately qualify AI-assisted analysis? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-TEST-001; OBDIA-EVID-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Reproducibility profile and controlled experiment plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-018` — Risk-rating calibration

| Field | Value |
|---|---|
| **Research Question** | How should the qualitative 5×5 risk method be calibrated against scenario evidence without implying mathematical probability? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-RISK-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Calibration workshop design, example scenarios and reviewer guidance. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-019` — Policy verification

| Field | Value |
|---|---|
| **Research Question** | Which formal or property-based techniques are proportionate for validating deny-by-default and non-transitive authorization invariants? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-AUTH-001; OBDIA-TEST-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Technique comparison, candidate properties and prototype plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-020` — Tamper-evident audit architecture

| Field | Value |
|---|---|
| **Research Question** | Which audit-linking, storage and verification patterns provide credible tamper evidence for a local synthetic laboratory? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-SEC-001; OBDIA-EVID-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Architecture options, test vectors and operational limitations. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-021` — Trusted time and uncertainty

| Field | Value |
|---|---|
| **Research Question** | How should local, provider and blockchain times be recorded with source and uncertainty for evidence and authorization decisions? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-EVID-001; OBDIA-AUTH-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Time-source model, uncertainty rules and test cases. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-022` — Key rotation without rebinding

| Field | Value |
|---|---|
| **Research Question** | How can agent, workload and connector keys rotate while preserving identity continuity and prohibiting officer rebinding? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-ID-001; OBDIA-AGENT-001; OBDIA-CONN-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Key-lifecycle model, failure modes and migration plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-023` — Software supply-chain provenance

| Field | Value |
|---|---|
| **Research Question** | Which minimum build and dependency provenance records are feasible for a public research repository? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-CODE-001; OBDIA-GH-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Provenance profile, tool-neutral schema and validation plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-024` — Policy rollback safety

| Field | Value |
|---|---|
| **Research Question** | How can policy rollback restore a known version without restoring revoked authority or deprecated insecure behavior? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-AUTH-001; OBDIA-VER-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Rollback invariants, simulation and negative tests. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-025` — Isolated research environment design

| Field | Value |
|---|---|
| **Research Question** | Which segmentation, egress, identity, logging and cleanup controls are required for high-risk but lawful isolated research? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-SEC-001; OBDIA-CONN-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Environment blueprint, threat model and validation checklist. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-026` — Web3 testnet evidence semantics

| Field | Value |
|---|---|
| **Research Question** | How should testnet transactions, blocks, provider observations and reorgs be represented as synthetic research evidence? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-EVID-001; OBDIA-CONN-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Evidence profile, provenance rules and test cases. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-027` — Smart-contract read-only verification

| Field | Value |
|---|---|
| **Research Question** | How can contract addresses, interfaces, upgradeability and simulated calls be verified without autonomous signing or mainnet state change? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-CONN-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Read-only workflow, simulation plan and limitations. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-028` — Blockchain finality and reorganization

| Field | Value |
|---|---|
| **Research Question** | How should chain finality assumptions and reorganization events affect evidence confidence and validation? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-EVID-001; OBDIA-RISK-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Finality model, scenario analysis and evidence-update rules. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-029` — Authenticated deep-web connector controls

| Field | Value |
|---|---|
| **Research Question** | Which account, session, rate, purpose and evidence controls are required for lawful authenticated-service research? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-CONN-001; OBDIA-AUTH-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Connector profile, threat analysis and sandbox validation plan. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-030` — Dark Web simulation fidelity

| Field | Value |
|---|---|
| **Research Question** | How realistic can a safe Dark Web simulation be without interacting with uncontrolled criminal infrastructure or reproducing harmful content? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-CONN-001; OBDIA-PUB-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Simulation specification, ethical boundaries and validation criteria. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-031` — Compliance-mapping evidence strength

| Field | Value |
|---|---|
| **Research Question** | How should mapping strength, gaps and uncertainty be represented so mappings remain design evidence rather than compliance proof? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-COMP-001 forward dependency |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Mapping-strength model and reviewer guidance. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-032` — Safe portfolio demonstration

| Field | Value |
|---|---|
| **Research Question** | Which synthetic scenario best demonstrates officer binding, authorization, evidence integrity and human accountability without real investigative activity? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-PUB-001; OBDIA-REL-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Scenario specification, data fixtures, diagrams and release checklist. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-033` — Independent-review pathway

| Field | Value |
|---|---|
| **Research Question** | What external review profiles would add meaningful assurance without overstating certification or institutional endorsement? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P3` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-GOV-001; OBDIA-REL-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Reviewer-role model, evidence package and claim boundaries. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-034` — Baseline-manifest automation

| Field | Value |
|---|---|
| **Research Question** | Which machine-readable manifest structure can validate document versions, hashes, statuses and dependencies across Knowledge Pack 01–30? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-VER-001; OBDIA-GH-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Manifest schema, validation script design and test fixtures. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-035` — Repository structural conformance

| Field | Value |
|---|---|
| **Research Question** | Which automated checks most reliably detect missing root artifacts, duplicate normative files, broken IDs and invalid Knowledge numbering? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-GH-001; OBDIA-TEST-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Conformance-test specification and synthetic repository fixtures. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-036` — Synthetic fixture realism

| Field | Value |
|---|---|
| **Research Question** | How can test fixtures remain realistic enough for meaningful security testing while avoiding real identities, cases and criminal infrastructure? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-TEST-001; OBDIA-PUB-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Fixture-generation method, safety review and quality criteria. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-037` — Incident evidence preservation

| Field | Value |
|---|---|
| **Research Question** | Which repository, runtime and connector evidence must be retained during containment without exposing secrets or case data? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-EVID-001; OBDIA-SEC-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Incident evidence profile, retention rules and recovery tests. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-038` — Retention-minimization trade-off

| Field | Value |
|---|---|
| **Research Question** | How should evidence, audit and research retention be balanced against privacy minimization and deletion obligations? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-EVID-001; OBDIA-RISK-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Decision framework, scenario analysis and governance proposal. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-039` — AI-assisted code review evidence

| Field | Value |
|---|---|
| **Research Question** | What evidence demonstrates accountable human review of AI-generated code beyond compilation and automated checks? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P2` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-CODE-001; OBDIA-TEST-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Review-evidence schema, examples and acceptance criteria. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


### `RSCH-ITEM-040` — Human-accountability interface

| Field | Value |
|---|---|
| **Research Question** | Which interface cues and approval records most clearly communicate that the officer remains responsible and the agent has no independent authority? |
| **Motivation** | Resolve a material OBDIA architecture, security, governance or validation uncertainty using controlled research. |
| **Owner** | Research Reviewer |
| **Priority** | `P1` |
| **Status** | `Triaged` |
| **Dependencies** | OBDIA-CANON-001; OBDIA-AUTH-001 |
| **Reference State** | `Sufficient for Planning` from internal normative sources; external primary-source selection remains required before external factual claims are validated. |
| **Validation Plan** | Produce a scoped source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative cases, reproducibility notes and human review. |
| **Expected Deliverables** | Interaction model, accessibility review and controlled usability study. |
| **Decision Boundary** | Results may inform an ADR, requirement change, risk update or implementation plan only through applicable governance. |
| **Publication Boundary** | Public use requires synthetic content, limitation statements and Release Reviewer assessment. |


## 41. Minimum Validation Checklist

Before a research item may be marked `Validated`, confirm:

- [ ] immutable `RSCH-ITEM-NNN` identifier exists;
- [ ] research question is precise and unbiased;
- [ ] motivation and intended decision are explicit;
- [ ] owner and required reviewers are identified;
- [ ] scope and exclusions are explicit;
- [ ] priority and rationale are current;
- [ ] dependencies, risks, threats and assumptions are linked;
- [ ] source records satisfy the required reference state;
- [ ] verified facts, hypotheses and interpretations are separated;
- [ ] validation plan defines method, data, environment and tools;
- [ ] acceptance, falsification and inconclusive criteria are defined;
- [ ] security, privacy and ethical constraints are satisfied;
- [ ] no unauthorized or uncontrolled infrastructure is used;
- [ ] raw observations and negative results are retained;
- [ ] deviations and failures remain visible;
- [ ] reproducibility requirements are met or limitations are stated;
- [ ] validation outcome is attributable and evidence-backed;
- [ ] residual uncertainty is explicit;
- [ ] architecture and requirements remain unchanged until governed adoption;
- [ ] resulting ADRs, changes, risks and assumptions are linked;
- [ ] public claims do not exceed evidence;
- [ ] closure and archival records are complete.


## 42. Limitations

- This document does not answer the initial research questions.
- It does not validate any hypothesis, architecture or control.
- It does not authorize implementation, operational investigation or access to external systems.
- The initial backlog is synthetic and may require reprioritization as documents 28–30 are consolidated.
- Source sufficiency is specific to each question and claim.
- Controlled experiments cannot reproduce every production, institutional or adversarial condition.
- Model and provider changes may invalidate prior research evidence.
- Negative results can be method-specific and shall not be generalized without justification.
- Internal review is not independent certification.
- A validated research conclusion still requires applicable ADR, change, implementation, testing, approval and release governance before becoming operational architecture.

## 43. Change Control

Every material change shall identify rationale, affected research items, sources, assumptions, risks, methods, evidence, architecture, security and privacy impact, migration, validation, rollback, authority and version effect.

- Editorial corrections shall use a patch version when meaning is unchanged.
- Backward-compatible substantive additions shall use a minor version.
- Incompatible research-governance changes shall use a major version.
- Material methodology decisions shall require an ADR where applicable.
- Changes shall require Project Founder approval.
- Changes shall receive Research Reviewer assessment.
- Security, privacy, implementation and release impacts shall receive specialist review where applicable.
- Changes shall identify affected records, schemas, indexes and automation.
- Changes shall include migration and compatibility analysis.
- Changes shall include validation and rollback analysis.
- Changes shall not retroactively fabricate research evidence.
- Historical research records shall not be silently rewritten.
- Research-item and requirement identifiers shall not be reused.
- Migration shall preserve item-to-source and item-to-decision traceability.
- Forward-dependency reconciliation shall occur before this policy becomes Approved.

## 44. Consolidation Record

Version 1.0.0 consolidates the two existing Research Backlog variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-RSCH-001`;
- normalizes the authoritative filename to `27_RESEARCH_BACKLOG.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves open research questions;
- preserves motivation;
- preserves dependencies;
- preserves references;
- preserves validation plan and validation status;
- preserves priority;
- preserves status;
- preserves expected deliverables;
- preserves the rule that no hypothesis becomes architecture without validation;
- adds immutable research-item identifiers, intake, triage, ownership, priority and status models;
- adds source discipline, reference states, hypotheses, assumptions, uncertainty and evidence;
- adds safe validation, reproducibility, environment, privacy, dual-use and publication controls;
- adds ADR, change, risk, closure, archival, metrics and bidirectional traceability;
- establishes an initial synthetic backlog of 40 research items;
- adds 999 stable research-governance requirements;
- identifies documents 28–30 as forward dependencies;
- treats Enterprise and non-Enterprise files as legacy source variants of the same immutable document;
- creates no validated hypothesis, architecture decision, implementation, risk acceptance, publication or operational authority.

## 45. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-06 | Draft | Research Reviewer; approval reserved to Project Founder | Constitutional consolidation of both legacy Research Backlog variants: preserved every original field and the validation-before-architecture boundary; added complete intake, sourcing, validation, safety, disposition, archival and traceability governance plus 40 synthetic initial research items. |
