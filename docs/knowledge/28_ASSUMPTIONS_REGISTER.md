# ASSUMPTIONS REGISTER

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-ASM-001 |
| **Title** | Assumptions Register |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative for assumption-governance rules; individual assumptions are Research records until validated and adopted through applicable governance |
| **Authority** | Project Founder |
| **Owner** | Research Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define mandatory identification, recording, ownership, classification, rationale, evidence, confidence, consequence, dependency, risk, validation, review, expiry, disposition, supersession, closure, archival and traceability requirements for OBDIA assumptions, and establish an initial synthetic assumptions baseline. |
| **Scope** | Architectural, identity, authorization, agent-lifecycle, evidence, connector, AI/model, privacy, security, software, supply-chain, repository, testing, release, cloud, Web3, external-provider, legal-context and research-method assumptions within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; `18_EVIDENCE_MODEL.md`; `19_AGENT_LIFECYCLE_MODEL.md`; `20_CONNECTOR_SECURITY_POLICY.md`; `21_SECURE_CODING_STANDARD.md`; `22_TESTING_STANDARD.md`; `23_DIAGRAM_STANDARD.md`; `24_GITHUB_REPOSITORY_STANDARD.md`; `25_DOCUMENT_VERSIONING_POLICY.md`; `26_AI_RISK_REGISTER.md`; `27_RESEARCH_BACKLOG.md`; forward dependencies `29_DECISION_LOG_POLICY.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-RES-001`; `OBDIA-ARCH-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001`; `OBDIA-EVID-001`; `OBDIA-AGENT-001`; `OBDIA-CONN-001`; `OBDIA-CODE-001`; `OBDIA-TEST-001`; `OBDIA-DIAG-001`; `OBDIA-GH-001`; `OBDIA-VER-001`; `OBDIA-RISK-001`; `OBDIA-RSCH-001` |
| **Cross-References** | Assumption records; research items; source records; hypotheses; threats; risks; requirements; ADRs; decision-log entries; implementation plans; test plans; validation packages; exceptions; changes; baseline manifests; releases and compliance mappings |
| **Assumptions** | The initial register is a synthetic governance baseline derived from current OBDIA design questions. It does not describe a deployed institution, real investigation, accepted risk or validated implementation. |
| **Constraints** | Assumptions shall not create independent AI legal authority, authorize prohibited conduct, substitute for evidence, override higher authority, conceal uncertainty, bypass validation or permit uncontrolled criminal-infrastructure interaction. |
| **Security Considerations** | Hidden, stale, false or weakly owned assumptions can invalidate authorization, officer-agent binding, evidence integrity, privacy controls, connector isolation, testing, releases and public claims. |
| **Validation Criteria** | Every active assumption has an immutable identifier, precise statement, rationale, owner, affected artifacts, consequence if false, confidence basis, evidence state, validation strategy, review schedule, expiry or trigger, risk links, current status and immutable history; no assumption becomes architecture without validation and governed adoption. |
| **Implementation Relationship** | This document governs assumption records and contains an initial synthetic register. It does not validate any assumption, approve architecture, authorize implementation, accept risk, establish compliance or create operational authority. |

---

## 1. Purpose

The non-Enterprise source required every architectural assumption to be recorded with:

- rationale;
- validation plan;
- review status.

The Enterprise source additionally required:

- Identifier;
- Statement;
- Evidence;
- Validation strategy;
- Review schedule;
- Current status.

This consolidation preserves every original field. It adds owner, affected artifacts, consequence if false, risk linkage, confidence and evidence separation, expiry, acceptance authority, validation result, supersession, closure and immutable history.

The policy sections are normative. Individual assumptions are research records. An assumption does not become a fact, requirement, architectural decision, control, legal conclusion or implementation merely because it is recorded.


## 2. Fundamental Assumption-Governance Rules

- **ASM-REQ-001:** Every material architectural assumption shall be recorded.
- **ASM-REQ-002:** Every material security assumption shall be recorded.
- **ASM-REQ-003:** Every material privacy assumption shall be recorded.
- **ASM-REQ-004:** Every material operational assumption shall be recorded.
- **ASM-REQ-005:** Every assumption shall have an immutable identifier.
- **ASM-REQ-006:** Every assumption shall have a precise statement.
- **ASM-REQ-007:** Every assumption shall have a rationale.
- **ASM-REQ-008:** Every assumption shall identify evidence.
- **ASM-REQ-009:** Every assumption shall identify a validation strategy.
- **ASM-REQ-010:** Every assumption shall identify a review schedule.
- **ASM-REQ-011:** Every assumption shall identify a current status.
- **ASM-REQ-012:** Every active assumption shall have an accountable human owner.
- **ASM-REQ-013:** Every assumption shall identify affected artifacts.
- **ASM-REQ-014:** Every assumption shall identify the consequence if false.
- **ASM-REQ-015:** Every assumption shall identify linked risks where material.
- **ASM-REQ-016:** Every assumption shall identify confidence and its basis.
- **ASM-REQ-017:** Evidence strength shall remain distinct from confidence.
- **ASM-REQ-018:** An assumption shall remain distinct from a verified fact.
- **ASM-REQ-019:** An assumption shall remain distinct from a hypothesis.
- **ASM-REQ-020:** An assumption shall remain distinct from a risk.
- **ASM-REQ-021:** An assumption shall remain distinct from an ADR.
- **ASM-REQ-022:** An assumption shall remain distinct from a requirement.
- **ASM-REQ-023:** An assumption shall remain distinct from an implementation decision.
- **ASM-REQ-024:** An assumption shall not possess normative authority.
- **ASM-REQ-025:** An assumption shall not override the Canonical Project Definition.
- **ASM-REQ-026:** An assumption shall not override the Constitution.
- **ASM-REQ-027:** An assumption shall not override an approved normative document.
- **ASM-REQ-028:** An assumption shall not authorize canonical prohibitions.
- **ASM-REQ-029:** No assumption shall become architecture without validation and governed adoption.
- **ASM-REQ-030:** Assumption history shall not be silently rewritten.

## 3. Authority and Classification

- **ASM-REQ-031:** This policy shall be normative for assumption-governance mechanics.
- **ASM-REQ-032:** Individual assumption records shall remain Research records.
- **ASM-REQ-033:** Individual assumption records shall remain below normative architecture in the authority hierarchy.
- **ASM-REQ-034:** Recording an assumption shall not approve it.
- **ASM-REQ-035:** Supporting an assumption shall not approve architecture derived from it.
- **ASM-REQ-036:** Validating an assumption shall not authorize implementation automatically.
- **ASM-REQ-037:** Closing an assumption shall not accept associated risk automatically.
- **ASM-REQ-038:** Publishing an assumption shall not establish factual truth.
- **ASM-REQ-039:** An assumption shall not create legal authority.
- **ASM-REQ-040:** An assumption shall not create institutional authority.
- **ASM-REQ-041:** An assumption shall not create operational authority.
- **ASM-REQ-042:** An assumption shall not authorize access.
- **ASM-REQ-043:** An assumption shall not authorize evidence acquisition.
- **ASM-REQ-044:** An assumption shall not authorize connector execution.
- **ASM-REQ-045:** An assumption shall not authorize autonomous decisions.
- **ASM-REQ-046:** Assumption classifications shall identify public, internal or restricted handling.
- **ASM-REQ-047:** Restricted assumptions shall not be copied into public artifacts without review.
- **ASM-REQ-048:** Assumption authority shall remain attributable to human roles.
- **ASM-REQ-049:** The Research Reviewer shall assess assumption quality.
- **ASM-REQ-050:** The Security Reviewer shall assess security-critical assumptions.
- **ASM-REQ-051:** The Privacy and Governance Reviewer shall assess rights and privacy assumptions.
- **ASM-REQ-052:** The Implementation Reviewer shall assess implementation-sensitive assumptions.
- **ASM-REQ-053:** The Release Reviewer shall assess assumptions underlying public claims.
- **ASM-REQ-054:** The Project Founder shall approve normative changes derived from assumptions where required.
- **ASM-REQ-055:** Role concentration shall be disclosed.

## 4. Identifier and Record Types

- **ASM-REQ-056:** Assumption identifiers shall use `ASM-NNN`.
- **ASM-REQ-057:** Assumption identifiers shall be immutable.
- **ASM-REQ-058:** Assumption identifiers shall not be reused.
- **ASM-REQ-059:** Assumption identifiers shall not encode status.
- **ASM-REQ-060:** Assumption identifiers shall not encode confidence.
- **ASM-REQ-061:** Assumption identifiers shall not encode priority.
- **ASM-REQ-062:** Assumption identifiers shall not encode owner.
- **ASM-REQ-063:** Assumption identifiers shall conform to `OBDIA-NAME-001`.
- **ASM-REQ-064:** Duplicate assumption identifiers shall be rejected.
- **ASM-REQ-065:** Superseded assumptions shall retain their identifiers.
- **ASM-REQ-066:** Archived assumptions shall retain their identifiers.
- **ASM-REQ-067:** Assumption-validation records shall have immutable identifiers.
- **ASM-REQ-068:** Assumption-review records shall have immutable identifiers.
- **ASM-REQ-069:** Assumption-correction records shall have immutable identifiers.
- **ASM-REQ-070:** Assumption-closure records shall have immutable identifiers.
- **ASM-REQ-071:** Evidence records shall not reuse assumption identifiers.
- **ASM-REQ-072:** Research items shall not reuse assumption identifiers.
- **ASM-REQ-073:** Risk records shall not reuse assumption identifiers.
- **ASM-REQ-074:** Decision records shall not reuse assumption identifiers.
- **ASM-REQ-075:** Unknown record types shall not acquire authority silently.

## 5. Mandatory Assumption Schema

- **ASM-REQ-076:** Every assumption shall include `Identifier`.
- **ASM-REQ-077:** Every assumption shall include `Title`.
- **ASM-REQ-078:** Every assumption shall include `Statement`.
- **ASM-REQ-079:** Every assumption shall include `Rationale`.
- **ASM-REQ-080:** Every assumption shall include `Category`.
- **ASM-REQ-081:** Every assumption shall include `Owner`.
- **ASM-REQ-082:** Every assumption shall include `Created Date`.
- **ASM-REQ-083:** Every assumption shall include `Current Status`.
- **ASM-REQ-084:** Every active assumption shall include `Review Schedule`.
- **ASM-REQ-085:** Every active assumption shall include `Next Review Date`.
- **ASM-REQ-086:** Every assumption shall include `Expiry Date or Trigger` where material.
- **ASM-REQ-087:** Every assumption shall include `Affected Artifacts`.
- **ASM-REQ-088:** Every assumption shall include `Affected Requirements` where known.
- **ASM-REQ-089:** Every assumption shall include `Dependencies`.
- **ASM-REQ-090:** Every assumption shall include `Evidence`.
- **ASM-REQ-091:** Every assumption shall include `Evidence State`.
- **ASM-REQ-092:** Every assumption shall include `Confidence`.
- **ASM-REQ-093:** Every assumption shall include `Confidence Basis`.
- **ASM-REQ-094:** Every assumption shall include `Consequence if False`.
- **ASM-REQ-095:** Every assumption shall include `Criticality`.
- **ASM-REQ-096:** Every assumption shall include `Risk Links` where applicable.
- **ASM-REQ-097:** Every assumption shall include `Threat Links` where applicable.
- **ASM-REQ-098:** Every assumption shall include `Research Links` where applicable.
- **ASM-REQ-099:** Every assumption shall include `Decision Links` where applicable.
- **ASM-REQ-100:** Every assumption shall include `Validation Strategy`.
- **ASM-REQ-101:** Every assumption shall include `Validation Criteria`.
- **ASM-REQ-102:** Every assumption shall include `Required Environment` where applicable.
- **ASM-REQ-103:** Every assumption shall include `Required Data` where applicable.
- **ASM-REQ-104:** Every assumption shall include `Security Constraints`.
- **ASM-REQ-105:** Every assumption shall include `Privacy Constraints`.
- **ASM-REQ-106:** Every assumption shall include `Known Limitations`.
- **ASM-REQ-107:** Every assumption shall include `Validation Result` after validation.
- **ASM-REQ-108:** Every validated assumption shall include `Validation Evidence`.
- **ASM-REQ-109:** Every validated assumption shall include `Validator`.
- **ASM-REQ-110:** Every validated assumption shall include `Validation Date`.
- **ASM-REQ-111:** Every closed assumption shall include `Disposition`.
- **ASM-REQ-112:** Every closed assumption shall include `Closure Authority`.
- **ASM-REQ-113:** Every closed assumption shall include `Closure Date`.
- **ASM-REQ-114:** Every superseded assumption shall include `Successor`.
- **ASM-REQ-115:** Every assumption shall maintain immutable `Revision History`.

## 6. Assumption Statement Quality

- **ASM-REQ-116:** Assumption statements shall be concise.
- **ASM-REQ-117:** Assumption statements shall be specific.
- **ASM-REQ-118:** Assumption statements shall be testable where technically feasible.
- **ASM-REQ-119:** Assumption statements shall identify the object or condition assumed.
- **ASM-REQ-120:** Assumption statements shall identify scope.
- **ASM-REQ-121:** Assumption statements shall identify environment where material.
- **ASM-REQ-122:** Assumption statements shall identify temporal conditions where material.
- **ASM-REQ-123:** Assumption statements shall identify actors where material.
- **ASM-REQ-124:** Assumption statements shall identify trust boundaries where material.
- **ASM-REQ-125:** Assumption statements shall avoid vague terms.
- **ASM-REQ-126:** Assumption statements shall use glossary terminology.
- **ASM-REQ-127:** Assumption statements shall not embed desired conclusions.
- **ASM-REQ-128:** Assumption statements shall not combine unrelated propositions.
- **ASM-REQ-129:** Compound assumptions shall be decomposed when independent validation is needed.
- **ASM-REQ-130:** Assumption statements shall distinguish technical feasibility from lawful authority.
- **ASM-REQ-131:** Assumption statements shall distinguish authentication from authorization.
- **ASM-REQ-132:** Assumption statements shall distinguish integrity from authenticity.
- **ASM-REQ-133:** Assumption statements shall distinguish model confidence from factual truth.
- **ASM-REQ-134:** Assumption statements shall distinguish Deep Web from Dark Web.
- **ASM-REQ-135:** Assumption statements shall not presume compliance.
- **ASM-REQ-136:** Assumption statements shall not presume institutional adoption.
- **ASM-REQ-137:** Assumption statements shall not presume production deployment.
- **ASM-REQ-138:** Material statement changes shall create a revision.
- **ASM-REQ-139:** Incompatible statement changes shall create a successor assumption where appropriate.
- **ASM-REQ-140:** Ambiguous statements shall not enter active validation.

## 7. Assumption Categories

- **ASM-REQ-141:** Permitted assumption categories shall be controlled.
- **ASM-REQ-142:** Architecture assumptions shall concern structural design conditions.
- **ASM-REQ-143:** Identity assumptions shall concern human, agent, workload or connector identity.
- **ASM-REQ-144:** Authorization assumptions shall concern policy inputs, decisions and enforcement.
- **ASM-REQ-145:** Lifecycle assumptions shall concern agent, connector, evidence or repository states.
- **ASM-REQ-146:** Evidence assumptions shall concern acquisition, integrity, provenance or custody.
- **ASM-REQ-147:** Security assumptions shall concern threats, controls and failure behavior.
- **ASM-REQ-148:** Privacy assumptions shall concern data, purpose, retention and rights.
- **ASM-REQ-149:** AI assumptions shall concern models, prompts, retrieval, memory or tools.
- **ASM-REQ-150:** Connector assumptions shall concern external integrations.
- **ASM-REQ-151:** Software assumptions shall concern code, dependencies, runtime or builds.
- **ASM-REQ-152:** Testing assumptions shall concern fixtures, environments, methods or acceptance.
- **ASM-REQ-153:** Repository assumptions shall concern settings, workflows, branches or releases.
- **ASM-REQ-154:** Provider assumptions shall concern external vendors or services.
- **ASM-REQ-155:** Cloud assumptions shall concern tenant, account, region or managed services.
- **ASM-REQ-156:** Web assumptions shall distinguish surface, authenticated deep-web and isolated Dark Web contexts.
- **ASM-REQ-157:** Web3 assumptions shall concern chains, contracts, wallets or providers.
- **ASM-REQ-158:** Research assumptions shall concern methods, sources or reproducibility.
- **ASM-REQ-159:** Legal-context assumptions shall be labeled as contextual assumptions rather than legal conclusions.
- **ASM-REQ-160:** Compliance assumptions shall not be represented as proof of compliance.
- **ASM-REQ-161:** An assumption may have multiple categories when justified.
- **ASM-REQ-162:** Category changes shall be attributable.
- **ASM-REQ-163:** Unknown categories shall require review.
- **ASM-REQ-164:** Category assignment shall not determine truth.
- **ASM-REQ-165:** Category assignment shall not determine authority.

## 8. Rationale

- **ASM-REQ-166:** Rationale shall explain why the assumption is necessary.
- **ASM-REQ-167:** Rationale shall identify the decision or activity relying on the assumption.
- **ASM-REQ-168:** Rationale shall identify why direct evidence is unavailable or incomplete.
- **ASM-REQ-169:** Rationale shall identify expected benefit.
- **ASM-REQ-170:** Rationale shall identify potential harm if incorrect.
- **ASM-REQ-171:** Rationale shall identify alternatives considered where material.
- **ASM-REQ-172:** Rationale shall identify temporary versus enduring use.
- **ASM-REQ-173:** Rationale shall identify scope boundaries.
- **ASM-REQ-174:** Rationale shall not restate the assumption without explanation.
- **ASM-REQ-175:** Rationale shall not use popularity as proof.
- **ASM-REQ-176:** Rationale shall not use model fluency as proof.
- **ASM-REQ-177:** Rationale shall not use vendor claims as independent evidence.
- **ASM-REQ-178:** Rationale shall not use urgency to bypass validation.
- **ASM-REQ-179:** Rationale shall not imply legal conclusions.
- **ASM-REQ-180:** Rationale shall identify unresolved uncertainty.
- **ASM-REQ-181:** Rationale shall identify source limitations.
- **ASM-REQ-182:** Rationale changes shall be attributable.
- **ASM-REQ-183:** Rationale changes shall trigger impact review where material.
- **ASM-REQ-184:** Missing rationale shall block active use.
- **ASM-REQ-185:** Weak rationale shall reduce confidence.

## 9. Evidence

- **ASM-REQ-186:** Evidence shall be attributable.
- **ASM-REQ-187:** Evidence shall identify source.
- **ASM-REQ-188:** Evidence shall identify date or version.
- **ASM-REQ-189:** Evidence shall identify relevance.
- **ASM-REQ-190:** Evidence shall identify limitations.
- **ASM-REQ-191:** Evidence shall identify whether it supports or contradicts the assumption.
- **ASM-REQ-192:** Evidence shall identify whether it is direct or indirect.
- **ASM-REQ-193:** Evidence shall identify whether it is empirical, documentary, analytical or testimonial.
- **ASM-REQ-194:** Primary sources shall be preferred where appropriate.
- **ASM-REQ-195:** Official sources shall be preferred for standards and platform behavior.
- **ASM-REQ-196:** Academic sources shall be preferred for empirical or theoretical claims.
- **ASM-REQ-197:** Marketing sources shall not be treated as independent confirmation.
- **ASM-REQ-198:** Model-generated text shall not be treated as evidence by itself.
- **ASM-REQ-199:** Repository documentation shall not prove platform settings by itself.
- **ASM-REQ-200:** Test output shall identify exact scope and versions.
- **ASM-REQ-201:** Integrity evidence shall not prove authenticity by itself.
- **ASM-REQ-202:** Evidence conflicts shall remain visible.
- **ASM-REQ-203:** Negative evidence shall remain visible.
- **ASM-REQ-204:** Missing evidence shall remain explicit.
- **ASM-REQ-205:** Evidence gaps shall not be filled with invented facts.
- **ASM-REQ-206:** Evidence shall conform to `OBDIA-RES-001`.
- **ASM-REQ-207:** Evidence derived from tests shall conform to `OBDIA-TEST-001`.
- **ASM-REQ-208:** Evidence involving digital objects shall conform to `OBDIA-EVID-001` where applicable.
- **ASM-REQ-209:** Evidence corrections shall preserve history.
- **ASM-REQ-210:** Evidence links shall remain valid.

## 10. Evidence State

- **ASM-REQ-211:** Permitted evidence states shall be `None`, `Preliminary`, `Partial`, `Sufficient for Planning`, `Sufficient for Validation`, `Conflicted`, `Stale` and `Invalid`.
- **ASM-REQ-212:** `None` shall mean no supporting evidence is recorded.
- **ASM-REQ-213:** `Preliminary` shall mean initial observations exist but are insufficient for reliance.
- **ASM-REQ-214:** `Partial` shall mean evidence supports only part of the statement or scope.
- **ASM-REQ-215:** `Sufficient for Planning` shall mean evidence supports temporary planning use within declared limits.
- **ASM-REQ-216:** `Sufficient for Validation` shall mean evidence meets defined validation needs for the stated scope.
- **ASM-REQ-217:** `Conflicted` shall mean material evidence disagrees.
- **ASM-REQ-218:** `Stale` shall mean evidence currency is inadequate.
- **ASM-REQ-219:** `Invalid` shall mean evidence cannot support reliable interpretation.
- **ASM-REQ-220:** Evidence state shall remain distinct from confidence.
- **ASM-REQ-221:** Evidence state shall remain distinct from assumption status.
- **ASM-REQ-222:** Evidence state shall remain distinct from validation outcome.
- **ASM-REQ-223:** Evidence state shall not be inferred from source count alone.
- **ASM-REQ-224:** Evidence state shall identify evaluator.
- **ASM-REQ-225:** Evidence state shall identify evaluation date.
- **ASM-REQ-226:** Conflicted evidence shall trigger analysis.
- **ASM-REQ-227:** Stale evidence shall trigger refresh or limitation.
- **ASM-REQ-228:** Invalid evidence shall not support validation.
- **ASM-REQ-229:** Evidence-state changes shall be attributable.
- **ASM-REQ-230:** Evidence-state history shall remain retrievable.

## 11. Confidence

- **ASM-REQ-231:** Permitted confidence levels shall be `Unknown`, `Low`, `Moderate` and `High`.
- **ASM-REQ-232:** `Unknown` shall mean confidence has not been supportably assessed.
- **ASM-REQ-233:** `Low` shall mean evidence or reasoning is weak, indirect, stale or highly uncertain.
- **ASM-REQ-234:** `Moderate` shall mean evidence and reasoning are useful but contain material limitations.
- **ASM-REQ-235:** `High` shall mean multiple strong and applicable evidence sources support the stated scope.
- **ASM-REQ-236:** Confidence shall identify its basis.
- **ASM-REQ-237:** Confidence shall identify uncertainty.
- **ASM-REQ-238:** Confidence shall identify evaluator.
- **ASM-REQ-239:** Confidence shall identify evaluation date.
- **ASM-REQ-240:** Confidence shall not be inferred from document length.
- **ASM-REQ-241:** Confidence shall not be inferred from source count alone.
- **ASM-REQ-242:** Confidence shall not be inferred from model certainty language.
- **ASM-REQ-243:** Confidence shall not be confused with evidence integrity.
- **ASM-REQ-244:** Confidence shall not be confused with authenticity.
- **ASM-REQ-245:** Confidence shall not be confused with admissibility.
- **ASM-REQ-246:** Confidence shall not be confused with probability unless a justified method exists.
- **ASM-REQ-247:** Confidence shall not be confused with risk likelihood.
- **ASM-REQ-248:** High confidence shall not create authority.
- **ASM-REQ-249:** Low confidence shall constrain reliance.
- **ASM-REQ-250:** Unknown confidence shall constrain reliance.
- **ASM-REQ-251:** Confidence changes shall identify new evidence or reasoning.
- **ASM-REQ-252:** Confidence reductions shall remain visible.
- **ASM-REQ-253:** Confidence history shall remain retrievable.
- **ASM-REQ-254:** Automated confidence recommendations shall remain advisory.
- **ASM-REQ-255:** Public claims shall not exceed confidence and evidence.

## 12. Criticality and Consequence if False

- **ASM-REQ-256:** Every assumption shall identify the consequence if false.
- **ASM-REQ-257:** Consequence analysis shall identify affected safety properties.
- **ASM-REQ-258:** Consequence analysis shall identify affected security properties.
- **ASM-REQ-259:** Consequence analysis shall identify affected privacy rights.
- **ASM-REQ-260:** Consequence analysis shall identify affected authorization.
- **ASM-REQ-261:** Consequence analysis shall identify affected evidence.
- **ASM-REQ-262:** Consequence analysis shall identify affected human accountability.
- **ASM-REQ-263:** Consequence analysis shall identify affected implementation.
- **ASM-REQ-264:** Consequence analysis shall identify affected tests.
- **ASM-REQ-265:** Consequence analysis shall identify affected releases.
- **ASM-REQ-266:** Consequence analysis shall identify affected public claims.
- **ASM-REQ-267:** Permitted criticality levels shall be `Critical`, `High`, `Medium` and `Low`.
- **ASM-REQ-268:** `Critical` shall mean falsity could invalidate canonical safety, legal-authority, rights, evidence-integrity or baseline foundations.
- **ASM-REQ-269:** `High` shall mean falsity could materially weaken security, privacy, authorization or core architecture.
- **ASM-REQ-270:** `Medium` shall mean falsity could require significant redesign or revalidation.
- **ASM-REQ-271:** `Low` shall mean falsity has limited and contained impact.
- **ASM-REQ-272:** Criticality shall not be inferred from priority alone.
- **ASM-REQ-273:** Criticality shall not be reduced without evidence.
- **ASM-REQ-274:** Critical assumptions shall receive enhanced review.
- **ASM-REQ-275:** Critical unsupported assumptions shall block approval.
- **ASM-REQ-276:** High unsupported assumptions shall block affected implementation unless constrained.
- **ASM-REQ-277:** Consequence changes shall trigger risk review.
- **ASM-REQ-278:** Consequence analysis shall identify containment.
- **ASM-REQ-279:** Consequence analysis shall identify recovery or rollback where material.
- **ASM-REQ-280:** Consequence history shall remain retrievable.

## 13. Status Model

- **ASM-REQ-281:** Permitted assumption statuses shall be `Proposed`, `Recorded`, `Active`, `Under Validation`, `Supported`, `Partially Supported`, `Not Supported`, `Inconclusive`, `Expired`, `Superseded`, `Closed` and `Archived`.
- **ASM-REQ-282:** `Proposed` shall mean an assumption has been submitted but not fully reviewed.
- **ASM-REQ-283:** `Recorded` shall mean the mandatory record exists.
- **ASM-REQ-284:** `Active` shall mean project work currently relies on the assumption within declared limits.
- **ASM-REQ-285:** `Under Validation` shall mean a governed validation activity is in progress.
- **ASM-REQ-286:** `Supported` shall mean evidence satisfies validation criteria for the stated scope.
- **ASM-REQ-287:** `Partially Supported` shall identify supported and unsupported portions.
- **ASM-REQ-288:** `Not Supported` shall mean evidence does not support the assumption.
- **ASM-REQ-289:** `Inconclusive` shall mean available evidence cannot support a reliable conclusion.
- **ASM-REQ-290:** `Expired` shall mean the review or validity condition has lapsed.
- **ASM-REQ-291:** `Superseded` shall mean another assumption replaces it.
- **ASM-REQ-292:** `Closed` shall mean disposition and closure evidence are complete.
- **ASM-REQ-293:** `Archived` shall mean the closed historical record is retained read-only.
- **ASM-REQ-294:** Assumption status shall remain distinct from repository lifecycle status.
- **ASM-REQ-295:** Assumption status shall remain distinct from research-item status.
- **ASM-REQ-296:** Assumption status shall remain distinct from validation outcome.
- **ASM-REQ-297:** Assumption status shall not be inferred from branch or PR state.
- **ASM-REQ-298:** Status transitions shall be attributable.
- **ASM-REQ-299:** Status transitions shall record rationale.
- **ASM-REQ-300:** Invalid transitions shall be rejected.
- **ASM-REQ-301:** A Proposed assumption shall not transition directly to Supported.
- **ASM-REQ-302:** An Active assumption shall identify reliance and controls.
- **ASM-REQ-303:** An Expired assumption shall not support new decisions.
- **ASM-REQ-304:** A Not Supported assumption shall trigger impact review.
- **ASM-REQ-305:** A Superseded assumption shall identify its successor.
- **ASM-REQ-306:** An Archived assumption shall remain read-only.

## 14. Reliance and Temporary Acceptance for Planning

- **ASM-REQ-307:** Reliance on an assumption shall be explicit.
- **ASM-REQ-308:** Reliance shall identify affected decisions.
- **ASM-REQ-309:** Reliance shall identify affected requirements.
- **ASM-REQ-310:** Reliance shall identify affected implementation.
- **ASM-REQ-311:** Reliance shall identify affected validation.
- **ASM-REQ-312:** Reliance shall identify reliance period.
- **ASM-REQ-313:** Reliance shall identify constraints.
- **ASM-REQ-314:** Reliance shall identify compensating controls.
- **ASM-REQ-315:** Reliance shall identify owner.
- **ASM-REQ-316:** Reliance shall identify approval authority where required.
- **ASM-REQ-317:** Temporary planning reliance shall not be called risk acceptance.
- **ASM-REQ-318:** Temporary planning reliance shall not create factual truth.
- **ASM-REQ-319:** Temporary planning reliance shall not authorize canonical prohibitions.
- **ASM-REQ-320:** Temporary planning reliance shall not authorize operational deployment.
- **ASM-REQ-321:** Temporary planning reliance shall not authorize access.
- **ASM-REQ-322:** Critical low-confidence assumptions shall not support approval.
- **ASM-REQ-323:** High-impact assumptions shall require stronger reliance controls.
- **ASM-REQ-324:** Reliance shall cease on expiry.
- **ASM-REQ-325:** Reliance shall cease when the assumption is Not Supported.
- **ASM-REQ-326:** Reliance shall be reassessed when evidence becomes Conflicted or Stale.
- **ASM-REQ-327:** Reliance decisions shall be attributable.
- **ASM-REQ-328:** Reliance decisions shall preserve role-concentration disclosure.
- **ASM-REQ-329:** AI systems shall not approve reliance.
- **ASM-REQ-330:** Reliance history shall remain retrievable.
- **ASM-REQ-331:** Reliance limitations shall remain visible.

## 15. Validation Strategy

- **ASM-REQ-332:** Every active assumption shall have a validation strategy.
- **ASM-REQ-333:** The validation strategy shall identify the assumption.
- **ASM-REQ-334:** The validation strategy shall identify method.
- **ASM-REQ-335:** The validation strategy shall identify evidence needs.
- **ASM-REQ-336:** The validation strategy shall identify data.
- **ASM-REQ-337:** The validation strategy shall identify environment.
- **ASM-REQ-338:** The validation strategy shall identify tools and versions.
- **ASM-REQ-339:** The validation strategy shall identify dependencies.
- **ASM-REQ-340:** The validation strategy shall identify responsible humans.
- **ASM-REQ-341:** The validation strategy shall identify reviewers.
- **ASM-REQ-342:** The validation strategy shall identify acceptance criteria.
- **ASM-REQ-343:** The validation strategy shall identify falsification criteria.
- **ASM-REQ-344:** The validation strategy shall identify partial-support criteria.
- **ASM-REQ-345:** The validation strategy shall identify inconclusive conditions.
- **ASM-REQ-346:** The validation strategy shall identify invalidation conditions.
- **ASM-REQ-347:** The validation strategy shall identify security constraints.
- **ASM-REQ-348:** The validation strategy shall identify privacy constraints.
- **ASM-REQ-349:** The validation strategy shall identify ethical constraints.
- **ASM-REQ-350:** The validation strategy shall identify authorization prerequisites.
- **ASM-REQ-351:** The validation strategy shall identify stop conditions.
- **ASM-REQ-352:** The validation strategy shall identify cleanup.
- **ASM-REQ-353:** The validation strategy shall identify retention.
- **ASM-REQ-354:** The validation strategy shall identify reproducibility requirements.
- **ASM-REQ-355:** The validation strategy shall identify expected evidence.
- **ASM-REQ-356:** The validation strategy shall not require prohibited operations.

## 16. Validation Execution

- **ASM-REQ-357:** Validation shall follow the approved strategy.
- **ASM-REQ-358:** Validation shall identify strategy version.
- **ASM-REQ-359:** Validation shall identify executor.
- **ASM-REQ-360:** Validation shall identify start and completion times.
- **ASM-REQ-361:** Validation shall identify environment.
- **ASM-REQ-362:** Validation shall identify tools and versions.
- **ASM-REQ-363:** Validation shall identify data and fixture versions.
- **ASM-REQ-364:** Validation shall identify deviations.
- **ASM-REQ-365:** Validation shall preserve raw observations.
- **ASM-REQ-366:** Validation shall preserve negative results.
- **ASM-REQ-367:** Validation shall preserve failures.
- **ASM-REQ-368:** Validation shall preserve interruptions.
- **ASM-REQ-369:** Validation shall preserve provenance.
- **ASM-REQ-370:** Validation shall preserve integrity where material.
- **ASM-REQ-371:** Validation shall not broaden scope silently.
- **ASM-REQ-372:** Validation shall not use unauthorized systems.
- **ASM-REQ-373:** Validation shall not use production credentials by default.
- **ASM-REQ-374:** Validation shall not use real case data in public demonstrations.
- **ASM-REQ-375:** Validation shall not interact with uncontrolled criminal infrastructure.
- **ASM-REQ-376:** Validation shall isolate high-risk content.
- **ASM-REQ-377:** Material method changes shall trigger strategy review.
- **ASM-REQ-378:** Invalid execution shall not support a Supported outcome.
- **ASM-REQ-379:** Aborted execution shall remain visible.
- **ASM-REQ-380:** Validation evidence shall conform to `OBDIA-TEST-001` where tests are used.
- **ASM-REQ-381:** Validation history shall remain retrievable.

## 17. Validation Outcomes

- **ASM-REQ-382:** Permitted validation outcomes shall be `Supported`, `Partially Supported`, `Not Supported`, `Inconclusive` and `Invalid`.
- **ASM-REQ-383:** `Supported` shall mean evidence satisfies defined criteria for the stated scope.
- **ASM-REQ-384:** `Partially Supported` shall identify supported and unsupported portions.
- **ASM-REQ-385:** `Not Supported` shall mean evidence does not support the assumption.
- **ASM-REQ-386:** `Inconclusive` shall mean available evidence cannot support a reliable conclusion.
- **ASM-REQ-387:** `Invalid` shall mean method, data or execution defects prevent reliable interpretation.
- **ASM-REQ-388:** Validation outcome shall identify evaluator.
- **ASM-REQ-389:** Validation outcome shall identify date.
- **ASM-REQ-390:** Validation outcome shall identify evidence.
- **ASM-REQ-391:** Validation outcome shall identify scope.
- **ASM-REQ-392:** Validation outcome shall identify limitations.
- **ASM-REQ-393:** Validation outcome shall identify residual uncertainty.
- **ASM-REQ-394:** Validation outcome shall identify confidence impact.
- **ASM-REQ-395:** Validation outcome shall identify risk impact.
- **ASM-REQ-396:** Validation outcome shall identify architecture impact.
- **ASM-REQ-397:** Validation outcome shall identify implementation impact.
- **ASM-REQ-398:** Validation outcome shall identify test impact.
- **ASM-REQ-399:** Validation outcome shall identify release impact.
- **ASM-REQ-400:** Validation outcome shall not amend architecture automatically.
- **ASM-REQ-401:** Validation outcome shall not amend requirements automatically.
- **ASM-REQ-402:** Validation outcome shall not approve implementation automatically.
- **ASM-REQ-403:** Validation outcome shall not accept risk automatically.
- **ASM-REQ-404:** Validation outcome shall not establish compliance.
- **ASM-REQ-405:** Outcome changes shall be additive and attributable.
- **ASM-REQ-406:** Invalid outcomes shall preserve failure evidence.

## 18. Review Schedule and Triggers

- **ASM-REQ-407:** Every active assumption shall have a review schedule.
- **ASM-REQ-408:** Review schedule shall reflect criticality.
- **ASM-REQ-409:** Review schedule shall reflect confidence.
- **ASM-REQ-410:** Review schedule shall reflect evidence state.
- **ASM-REQ-411:** Review schedule shall reflect dependency volatility.
- **ASM-REQ-412:** Review schedule shall reflect provider volatility.
- **ASM-REQ-413:** Review schedule shall reflect legal-context volatility where material.
- **ASM-REQ-414:** Critical assumptions shall receive frequent review.
- **ASM-REQ-415:** High assumptions shall receive regular review.
- **ASM-REQ-416:** Medium and Low assumptions shall receive risk-based review.
- **ASM-REQ-417:** Every review shall identify reviewer.
- **ASM-REQ-418:** Every review shall identify date.
- **ASM-REQ-419:** Every review shall identify evidence changes.
- **ASM-REQ-420:** Every review shall identify confidence changes.
- **ASM-REQ-421:** Every review shall identify risk changes.
- **ASM-REQ-422:** Every review shall identify reliance changes.
- **ASM-REQ-423:** A material threat change shall trigger review.
- **ASM-REQ-424:** A material architecture change shall trigger review.
- **ASM-REQ-425:** A material provider change shall trigger review.
- **ASM-REQ-426:** A material model change shall trigger review.
- **ASM-REQ-427:** A security or privacy incident shall trigger review.
- **ASM-REQ-428:** A failed validation shall trigger review.
- **ASM-REQ-429:** Source staleness shall trigger review.
- **ASM-REQ-430:** Release planning shall trigger review of release-relevant assumptions.
- **ASM-REQ-431:** Missed review dates shall remain visible.

## 19. Expiry

- **ASM-REQ-432:** Every time-sensitive assumption shall have an expiry date or trigger.
- **ASM-REQ-433:** Expiry shall identify rationale.
- **ASM-REQ-434:** Expiry shall identify owner.
- **ASM-REQ-435:** Expiry shall identify required action.
- **ASM-REQ-436:** Expired assumptions shall transition to Expired.
- **ASM-REQ-437:** Expired assumptions shall not support new approval.
- **ASM-REQ-438:** Expired assumptions shall not support new implementation.
- **ASM-REQ-439:** Expired assumptions shall not support new validation.
- **ASM-REQ-440:** Expired assumptions shall not support public claims.
- **ASM-REQ-441:** Expired assumptions shall trigger impact review.
- **ASM-REQ-442:** Expiry extension shall require rationale.
- **ASM-REQ-443:** Expiry extension shall require current evidence review.
- **ASM-REQ-444:** Expiry extension shall be attributable.
- **ASM-REQ-445:** Automatic expiry shall not delete the record.
- **ASM-REQ-446:** Automatic expiry shall not create a validation result.
- **ASM-REQ-447:** Assumptions tied to provider versions shall expire on material provider change.
- **ASM-REQ-448:** Assumptions tied to legal or standards versions shall expire on relevant change.
- **ASM-REQ-449:** Assumptions tied to threat conditions shall expire on material threat change.
- **ASM-REQ-450:** Assumptions tied to environments shall expire on material environment change.
- **ASM-REQ-451:** Expiry history shall remain retrievable.

## 20. Risk Integration

- **ASM-REQ-452:** Assumptions shall link to affected risks.
- **ASM-REQ-453:** False-assumption consequences shall inform impact assessment.
- **ASM-REQ-454:** Evidence uncertainty shall inform risk confidence.
- **ASM-REQ-455:** Assumption volatility shall inform risk review cadence.
- **ASM-REQ-456:** Critical unsupported assumptions shall create or update risks.
- **ASM-REQ-457:** Not Supported outcomes shall trigger risk review.
- **ASM-REQ-458:** Inconclusive outcomes shall not reduce risk without justification.
- **ASM-REQ-459:** Expired assumptions shall trigger risk review.
- **ASM-REQ-460:** Conflicted evidence shall trigger risk review.
- **ASM-REQ-461:** Stale evidence shall trigger risk review.
- **ASM-REQ-462:** Assumption reliance shall identify residual risk.
- **ASM-REQ-463:** Assumption validation shall identify control-effectiveness implications.
- **ASM-REQ-464:** Assumption closure shall identify unresolved risk.
- **ASM-REQ-465:** An assumption shall not accept risk.
- **ASM-REQ-466:** An AI system shall not accept assumption-related risk.
- **ASM-REQ-467:** Risk owners shall remain human.
- **ASM-REQ-468:** Canonical-prohibition risks shall not be accepted.
- **ASM-REQ-469:** Critical residual risk shall block affected approval.
- **ASM-REQ-470:** Risk updates shall use `OBDIA-RISK-001`.
- **ASM-REQ-471:** Assumption-risk history shall remain retrievable.

## 21. Research Integration

- **ASM-REQ-472:** Assumptions shall link to relevant research items.
- **ASM-REQ-473:** Research items shall identify assumptions they depend on.
- **ASM-REQ-474:** Research may propose new assumptions.
- **ASM-REQ-475:** Research may validate assumptions.
- **ASM-REQ-476:** Research may invalidate assumptions.
- **ASM-REQ-477:** Research may reduce assumption scope.
- **ASM-REQ-478:** Research may identify successor assumptions.
- **ASM-REQ-479:** Research results shall not rewrite assumption history.
- **ASM-REQ-480:** Assumption validation plans may be managed through research items.
- **ASM-REQ-481:** Research evidence shall identify exact assumption versions.
- **ASM-REQ-482:** Assumption changes shall trigger research impact review.
- **ASM-REQ-483:** Research closure shall identify assumption outcomes.
- **ASM-REQ-484:** Research priority shall not create assumption authority.
- **ASM-REQ-485:** Supported research conclusions shall not amend architecture automatically.
- **ASM-REQ-486:** Inconclusive research shall not convert an assumption into fact.
- **ASM-REQ-487:** Negative research results shall remain visible.
- **ASM-REQ-488:** Research methods shall conform to `OBDIA-RSCH-001`.
- **ASM-REQ-489:** Source handling shall conform to `OBDIA-RES-001`.
- **ASM-REQ-490:** Research and assumption identifiers shall remain distinct.
- **ASM-REQ-491:** Research-assumption history shall remain retrievable.

## 22. Decision and ADR Integration

- **ASM-REQ-492:** Decisions relying on assumptions shall identify exact assumption identifiers.
- **ASM-REQ-493:** ADRs relying on assumptions shall identify exact assumption identifiers.
- **ASM-REQ-494:** Decisions shall identify assumption status and confidence.
- **ASM-REQ-495:** Decisions shall identify consequence if the assumption is false.
- **ASM-REQ-496:** Decisions shall identify validation requirements.
- **ASM-REQ-497:** Decisions shall identify fallback or rollback where material.
- **ASM-REQ-498:** Accepted ADRs shall not convert assumptions into facts automatically.
- **ASM-REQ-499:** Assumption invalidation shall trigger ADR impact review.
- **ASM-REQ-500:** Assumption expiry shall trigger decision impact review.
- **ASM-REQ-501:** Assumption supersession shall update decision links.
- **ASM-REQ-502:** Decision records shall preserve historical assumption versions.
- **ASM-REQ-503:** Research records shall not substitute for ADRs.
- **ASM-REQ-504:** Assumption records shall not substitute for ADRs.
- **ASM-REQ-505:** Document 29 shall provide the chronological decision index.
- **ASM-REQ-506:** Decision changes shall not rewrite assumption history.
- **ASM-REQ-507:** Rejected decisions shall preserve assumption links.
- **ASM-REQ-508:** Partial reliance shall identify scope.
- **ASM-REQ-509:** Decision authority shall remain human.
- **ASM-REQ-510:** AI systems shall not approve assumption-dependent decisions.
- **ASM-REQ-511:** Decision-assumption history shall remain retrievable.

## 23. Architecture and Requirements Impact

- **ASM-REQ-512:** Assumptions shall identify affected architecture.
- **ASM-REQ-513:** Assumptions shall identify affected requirements.
- **ASM-REQ-514:** Assumptions shall identify affected trust boundaries.
- **ASM-REQ-515:** Assumptions shall identify affected identities.
- **ASM-REQ-516:** Assumptions shall identify affected authorization.
- **ASM-REQ-517:** Assumptions shall identify affected evidence.
- **ASM-REQ-518:** Assumptions shall identify affected connectors.
- **ASM-REQ-519:** Assumptions shall identify affected models and prompts.
- **ASM-REQ-520:** Assumptions shall identify affected privacy controls.
- **ASM-REQ-521:** Assumptions shall identify affected software and dependencies.
- **ASM-REQ-522:** Assumptions shall identify affected tests.
- **ASM-REQ-523:** Assumptions shall identify affected releases.
- **ASM-REQ-524:** Architecture shall not conceal critical assumptions.
- **ASM-REQ-525:** Requirements shall not present assumptions as facts.
- **ASM-REQ-526:** Architecture diagrams shall label material assumptions or link to them.
- **ASM-REQ-527:** Implementation plans shall identify active assumptions.
- **ASM-REQ-528:** Validation plans shall identify active assumptions.
- **ASM-REQ-529:** Release reviews shall identify release-relevant assumptions.
- **ASM-REQ-530:** Assumption invalidation shall trigger change control.
- **ASM-REQ-531:** Assumption support shall not amend requirements automatically.
- **ASM-REQ-532:** Architecture adoption shall require applicable ADRs.
- **ASM-REQ-533:** Requirement adoption shall require governed change.
- **ASM-REQ-534:** Assumption changes shall preserve traceability.
- **ASM-REQ-535:** Critical broken traceability shall block approval.
- **ASM-REQ-536:** Architecture-assumption history shall remain retrievable.

## 24. Identity and Binding Assumptions

- **ASM-REQ-537:** Identity assumptions shall distinguish human, agent, workload and connector identities.
- **ASM-REQ-538:** Identity assumptions shall preserve accountable institution.
- **ASM-REQ-539:** Binding assumptions shall preserve one-agent-to-one-officer binding.
- **ASM-REQ-540:** Binding assumptions shall not presume credential sharing.
- **ASM-REQ-541:** Binding assumptions shall not presume legal authority from authentication.
- **ASM-REQ-542:** Binding assumptions shall identify cryptographic mechanisms where material.
- **ASM-REQ-543:** Binding assumptions shall identify issuance authority.
- **ASM-REQ-544:** Binding assumptions shall identify revocation authority.
- **ASM-REQ-545:** Binding assumptions shall identify key-rotation effects.
- **ASM-REQ-546:** Binding assumptions shall identify re-binding prohibitions.
- **ASM-REQ-547:** Identity assumptions shall identify directory or registry dependencies.
- **ASM-REQ-548:** Identity assumptions shall identify freshness.
- **ASM-REQ-549:** Identity assumptions shall identify compromise behavior.
- **ASM-REQ-550:** Identity assumptions shall identify recovery limitations.
- **ASM-REQ-551:** Identity assumptions shall identify privacy effects.
- **ASM-REQ-552:** Critical identity assumptions shall receive Security Reviewer assessment.
- **ASM-REQ-553:** Identity validation shall use synthetic identities.
- **ASM-REQ-554:** Identity validation shall not transfer human credentials.
- **ASM-REQ-555:** Identity assumptions shall conform to `OBDIA-ID-001`.
- **ASM-REQ-556:** Identity-assumption history shall remain retrievable.

## 25. Authorization Assumptions

- **ASM-REQ-557:** Authorization assumptions shall default to deny.
- **ASM-REQ-558:** Authorization assumptions shall identify subject.
- **ASM-REQ-559:** Authorization assumptions shall identify resource.
- **ASM-REQ-560:** Authorization assumptions shall identify action.
- **ASM-REQ-561:** Authorization assumptions shall identify context.
- **ASM-REQ-562:** Authorization assumptions shall identify policy version.
- **ASM-REQ-563:** Authorization assumptions shall identify attribute sources.
- **ASM-REQ-564:** Authorization assumptions shall identify freshness.
- **ASM-REQ-565:** Authorization assumptions shall identify failure behavior.
- **ASM-REQ-566:** Authorization assumptions shall identify revocation behavior.
- **ASM-REQ-567:** Authorization assumptions shall identify cache behavior.
- **ASM-REQ-568:** Authorization assumptions shall identify human-approval boundaries.
- **ASM-REQ-569:** Authorization assumptions shall identify case and purpose boundaries.
- **ASM-REQ-570:** Authorization assumptions shall identify jurisdiction and time constraints where material.
- **ASM-REQ-571:** Authorization assumptions shall not infer Permit from missing context.
- **ASM-REQ-572:** Authorization assumptions shall not infer legal authority from Permit.
- **ASM-REQ-573:** Authorization validation shall include denial and failure cases.
- **ASM-REQ-574:** Critical authorization assumptions shall receive Security Reviewer assessment.
- **ASM-REQ-575:** Authorization assumptions shall conform to `OBDIA-AUTH-001`.
- **ASM-REQ-576:** Authorization-assumption history shall remain retrievable.

## 26. Evidence Assumptions

- **ASM-REQ-577:** Evidence assumptions shall distinguish original and derived artifacts.
- **ASM-REQ-578:** Evidence assumptions shall identify source.
- **ASM-REQ-579:** Evidence assumptions shall identify collection method.
- **ASM-REQ-580:** Evidence assumptions shall identify time source.
- **ASM-REQ-581:** Evidence assumptions shall identify collector.
- **ASM-REQ-582:** Evidence assumptions shall identify integrity mechanism.
- **ASM-REQ-583:** Evidence assumptions shall identify provenance.
- **ASM-REQ-584:** Evidence assumptions shall identify custody.
- **ASM-REQ-585:** Evidence assumptions shall identify storage.
- **ASM-REQ-586:** Evidence assumptions shall identify retention.
- **ASM-REQ-587:** Evidence assumptions shall identify export behavior.
- **ASM-REQ-588:** Evidence assumptions shall not equate integrity with authenticity.
- **ASM-REQ-589:** Evidence assumptions shall not equate confidence with admissibility.
- **ASM-REQ-590:** Evidence assumptions shall not imply lawful acquisition.
- **ASM-REQ-591:** Evidence assumptions shall identify transformation behavior.
- **ASM-REQ-592:** Evidence assumptions shall identify quarantine behavior.
- **ASM-REQ-593:** Evidence validation shall use synthetic evidence.
- **ASM-REQ-594:** Critical evidence assumptions shall receive evidence and security review.
- **ASM-REQ-595:** Evidence assumptions shall conform to `OBDIA-EVID-001`.
- **ASM-REQ-596:** Evidence-assumption history shall remain retrievable.

## 27. Agent-Lifecycle Assumptions

- **ASM-REQ-597:** Agent-lifecycle assumptions shall preserve the authoritative state sequence.
- **ASM-REQ-598:** Agent-lifecycle assumptions shall preserve mandatory `Bind`.
- **ASM-REQ-599:** Agent-lifecycle assumptions shall preserve authorization before activation.
- **ASM-REQ-600:** Agent-lifecycle assumptions shall preserve terminal revocation.
- **ASM-REQ-601:** Agent-lifecycle assumptions shall preserve terminal archival.
- **ASM-REQ-602:** Agent-lifecycle assumptions shall prohibit officer re-binding.
- **ASM-REQ-603:** Agent-lifecycle assumptions shall identify state source of truth.
- **ASM-REQ-604:** Agent-lifecycle assumptions shall identify concurrency.
- **ASM-REQ-605:** Agent-lifecycle assumptions shall identify stale-state behavior.
- **ASM-REQ-606:** Agent-lifecycle assumptions shall identify propagation.
- **ASM-REQ-607:** Agent-lifecycle assumptions shall identify credential effects.
- **ASM-REQ-608:** Agent-lifecycle assumptions shall identify session effects.
- **ASM-REQ-609:** Agent-lifecycle assumptions shall identify workload effects.
- **ASM-REQ-610:** Agent-lifecycle assumptions shall identify connector effects.
- **ASM-REQ-611:** Agent-lifecycle assumptions shall identify recovery limitations.
- **ASM-REQ-612:** Lifecycle validation shall include invalid transitions.
- **ASM-REQ-613:** Critical lifecycle assumptions shall receive Security Reviewer assessment.
- **ASM-REQ-614:** Lifecycle assumptions shall conform to `OBDIA-AGENT-001`.
- **ASM-REQ-615:** Lifecycle changes shall trigger assumption review.
- **ASM-REQ-616:** Lifecycle-assumption history shall remain retrievable.

## 28. Connector and External-Service Assumptions

- **ASM-REQ-617:** Connector assumptions shall identify external system ownership.
- **ASM-REQ-618:** Connector assumptions shall identify authentication.
- **ASM-REQ-619:** Connector assumptions shall identify authorization.
- **ASM-REQ-620:** Connector assumptions shall identify credential scope.
- **ASM-REQ-621:** Connector assumptions shall identify endpoint scope.
- **ASM-REQ-622:** Connector assumptions shall identify egress controls.
- **ASM-REQ-623:** Connector assumptions shall identify request validation.
- **ASM-REQ-624:** Connector assumptions shall identify response validation.
- **ASM-REQ-625:** Connector assumptions shall identify provenance.
- **ASM-REQ-626:** Connector assumptions shall identify failure isolation.
- **ASM-REQ-627:** Connector assumptions shall identify revocation.
- **ASM-REQ-628:** Connector assumptions shall identify provider version risk.
- **ASM-REQ-629:** Connector assumptions shall identify rate and quota limits.
- **ASM-REQ-630:** Connector assumptions shall identify retention and privacy behavior.
- **ASM-REQ-631:** Connector assumptions shall treat returned content as untrusted.
- **ASM-REQ-632:** Connector assumptions shall not presume service availability.
- **ASM-REQ-633:** Connector assumptions shall not use human credentials.
- **ASM-REQ-634:** Connector validation shall use mocks, sandboxes or testnets.
- **ASM-REQ-635:** Connector assumptions shall conform to `OBDIA-CONN-001`.
- **ASM-REQ-636:** Connector-assumption history shall remain retrievable.

## 29. AI and Model Assumptions

- **ASM-REQ-637:** AI assumptions shall identify model and provider.
- **ASM-REQ-638:** AI assumptions shall identify model version where available.
- **ASM-REQ-639:** AI assumptions shall identify prompt version where material.
- **ASM-REQ-640:** AI assumptions shall identify sampling settings where material.
- **ASM-REQ-641:** AI assumptions shall identify retrieval and memory context.
- **ASM-REQ-642:** AI assumptions shall identify tool boundaries.
- **ASM-REQ-643:** AI assumptions shall treat model output as untrusted.
- **ASM-REQ-644:** AI assumptions shall not grant independent legal authority.
- **ASM-REQ-645:** AI assumptions shall not grant authorization authority.
- **ASM-REQ-646:** AI assumptions shall not grant risk-acceptance authority.
- **ASM-REQ-647:** AI assumptions shall identify nondeterminism.
- **ASM-REQ-648:** AI assumptions shall identify drift risk.
- **ASM-REQ-649:** AI assumptions shall identify provider silent-change risk.
- **ASM-REQ-650:** AI assumptions shall identify reproducibility limits.
- **ASM-REQ-651:** AI assumptions shall identify prompt-injection risk.
- **ASM-REQ-652:** AI assumptions shall identify privacy and retention risk.
- **ASM-REQ-653:** AI validation shall use synthetic inputs.
- **ASM-REQ-654:** AI validation shall preserve human accountability.
- **ASM-REQ-655:** Critical AI assumptions shall receive security and privacy review.
- **ASM-REQ-656:** AI-assumption history shall remain retrievable.

## 30. Privacy and Rights Assumptions

- **ASM-REQ-657:** Privacy assumptions shall identify personal-data categories.
- **ASM-REQ-658:** Privacy assumptions shall identify purpose.
- **ASM-REQ-659:** Privacy assumptions shall identify data scope.
- **ASM-REQ-660:** Privacy assumptions shall identify retention.
- **ASM-REQ-661:** Privacy assumptions shall identify deletion.
- **ASM-REQ-662:** Privacy assumptions shall identify access.
- **ASM-REQ-663:** Privacy assumptions shall identify transfer.
- **ASM-REQ-664:** Privacy assumptions shall identify jurisdiction where material.
- **ASM-REQ-665:** Privacy assumptions shall identify pseudonymization or anonymization limits.
- **ASM-REQ-666:** Privacy assumptions shall identify cross-case isolation.
- **ASM-REQ-667:** Privacy assumptions shall identify cross-purpose isolation.
- **ASM-REQ-668:** Privacy assumptions shall identify human-review needs.
- **ASM-REQ-669:** Privacy assumptions shall identify discrimination risk.
- **ASM-REQ-670:** Privacy assumptions shall identify surveillance risk.
- **ASM-REQ-671:** Privacy assumptions shall not treat public availability as unrestricted authority.
- **ASM-REQ-672:** Privacy assumptions shall not overstate anonymization.
- **ASM-REQ-673:** Privacy validation shall minimize personal data.
- **ASM-REQ-674:** Critical privacy assumptions shall receive Privacy and Governance Reviewer assessment.
- **ASM-REQ-675:** Privacy incidents shall trigger assumption review.
- **ASM-REQ-676:** Privacy-assumption history shall remain retrievable.

## 31. Software, Supply-Chain and Repository Assumptions

- **ASM-REQ-677:** Software assumptions shall identify language and runtime versions where material.
- **ASM-REQ-678:** Software assumptions shall identify dependency versions.
- **ASM-REQ-679:** Software assumptions shall identify build tools.
- **ASM-REQ-680:** Software assumptions shall identify deployment environment.
- **ASM-REQ-681:** Software assumptions shall identify secure defaults.
- **ASM-REQ-682:** Software assumptions shall identify error behavior.
- **ASM-REQ-683:** Software assumptions shall identify resource limits.
- **ASM-REQ-684:** Supply-chain assumptions shall identify provenance.
- **ASM-REQ-685:** Supply-chain assumptions shall identify registry and source trust.
- **ASM-REQ-686:** Supply-chain assumptions shall identify update behavior.
- **ASM-REQ-687:** Repository assumptions shall distinguish desired from observed settings.
- **ASM-REQ-688:** Repository assumptions shall identify branch and review protections.
- **ASM-REQ-689:** Repository assumptions shall identify workflow permissions.
- **ASM-REQ-690:** Repository assumptions shall identify secret handling.
- **ASM-REQ-691:** Repository assumptions shall identify release behavior.
- **ASM-REQ-692:** Repository assumptions shall not infer enforcement from documentation alone.
- **ASM-REQ-693:** Validation shall identify exact commits and settings evidence.
- **ASM-REQ-694:** Critical software assumptions shall receive implementation and security review.
- **ASM-REQ-695:** Software assumptions shall conform to `OBDIA-CODE-001` and `OBDIA-GH-001`.
- **ASM-REQ-696:** Software-assumption history shall remain retrievable.

## 32. Testing and Validation Assumptions

- **ASM-REQ-697:** Testing assumptions shall identify implementation version.
- **ASM-REQ-698:** Testing assumptions shall identify test definitions.
- **ASM-REQ-699:** Testing assumptions shall identify fixtures.
- **ASM-REQ-700:** Testing assumptions shall identify environment.
- **ASM-REQ-701:** Testing assumptions shall identify tools and versions.
- **ASM-REQ-702:** Testing assumptions shall identify expected results.
- **ASM-REQ-703:** Testing assumptions shall identify acceptance criteria.
- **ASM-REQ-704:** Testing assumptions shall identify negative cases.
- **ASM-REQ-705:** Testing assumptions shall identify reproducibility.
- **ASM-REQ-706:** Testing assumptions shall identify nondeterminism.
- **ASM-REQ-707:** Testing assumptions shall identify external dependencies.
- **ASM-REQ-708:** Testing assumptions shall identify coverage limitations.
- **ASM-REQ-709:** Testing assumptions shall not infer security from coverage percentages alone.
- **ASM-REQ-710:** Testing assumptions shall not infer validation from passing tests alone.
- **ASM-REQ-711:** Testing assumptions shall identify flaky-test effects.
- **ASM-REQ-712:** Testing assumptions shall identify evidence retention.
- **ASM-REQ-713:** Validation changes shall trigger assumption review.
- **ASM-REQ-714:** Critical test assumptions shall receive Implementation Reviewer assessment.
- **ASM-REQ-715:** Testing assumptions shall conform to `OBDIA-TEST-001`.
- **ASM-REQ-716:** Testing-assumption history shall remain retrievable.

## 33. Web, Cloud and Web3 Assumptions

- **ASM-REQ-717:** Web assumptions shall distinguish public surface-web access.
- **ASM-REQ-718:** Web assumptions shall distinguish authorized authenticated-service access.
- **ASM-REQ-719:** Web assumptions shall distinguish operationally isolated Dark Web research.
- **ASM-REQ-720:** Web assumptions shall not conflate Deep Web with Dark Web.
- **ASM-REQ-721:** Dark Web assumptions shall prohibit uncontrolled criminal-infrastructure interaction.
- **ASM-REQ-722:** Cloud assumptions shall identify provider.
- **ASM-REQ-723:** Cloud assumptions shall identify account and tenant.
- **ASM-REQ-724:** Cloud assumptions shall identify region.
- **ASM-REQ-725:** Cloud assumptions shall identify identity and network boundaries.
- **ASM-REQ-726:** Cloud assumptions shall identify managed-service behavior.
- **ASM-REQ-727:** Web3 assumptions shall identify chain or network.
- **ASM-REQ-728:** Web3 assumptions shall identify testnet, local or mainnet scope.
- **ASM-REQ-729:** Web3 assumptions shall identify provider and node behavior.
- **ASM-REQ-730:** Web3 assumptions shall identify finality and reorganization.
- **ASM-REQ-731:** Smart-contract assumptions shall identify address and interface versions.
- **ASM-REQ-732:** Smart-contract assumptions shall identify upgradeability.
- **ASM-REQ-733:** Web3 assumptions shall distinguish read-only analysis from state change.
- **ASM-REQ-734:** Validation shall use testnets or simulation by default.
- **ASM-REQ-735:** Critical external-environment assumptions shall receive security review.
- **ASM-REQ-736:** Environment-assumption history shall remain retrievable.

## 34. Public Claims and Compliance Assumptions

- **ASM-REQ-737:** Public claims shall identify underlying assumptions where material.
- **ASM-REQ-738:** Public claims shall not present assumptions as verified facts.
- **ASM-REQ-739:** Public claims shall identify uncertainty.
- **ASM-REQ-740:** Public claims shall identify synthetic scope.
- **ASM-REQ-741:** Public claims shall identify implementation status.
- **ASM-REQ-742:** Public claims shall identify validation status.
- **ASM-REQ-743:** Public claims shall not imply institutional adoption.
- **ASM-REQ-744:** Public claims shall not imply production readiness.
- **ASM-REQ-745:** Public claims shall not imply certification.
- **ASM-REQ-746:** Public claims shall not imply legal compliance.
- **ASM-REQ-747:** Compliance mappings shall identify assumptions.
- **ASM-REQ-748:** Compliance assumptions shall identify applicable source versions.
- **ASM-REQ-749:** Compliance assumptions shall identify jurisdiction.
- **ASM-REQ-750:** Compliance assumptions shall identify mapping strength.
- **ASM-REQ-751:** Compliance assumptions shall identify evidence gaps.
- **ASM-REQ-752:** Compliance assumptions shall not replace legal advice.
- **ASM-REQ-753:** Release review shall assess public-facing assumptions.
- **ASM-REQ-754:** Expired assumptions shall not support public claims.
- **ASM-REQ-755:** Document 30 shall govern complete compliance-mapping controls.
- **ASM-REQ-756:** Public-claim assumption history shall remain retrievable.

## 35. Supersession

- **ASM-REQ-757:** Supersession shall identify predecessor.
- **ASM-REQ-758:** Supersession shall identify successor.
- **ASM-REQ-759:** Supersession shall identify authority.
- **ASM-REQ-760:** Supersession shall identify rationale.
- **ASM-REQ-761:** Supersession shall identify effective date.
- **ASM-REQ-762:** Supersession shall identify affected artifacts.
- **ASM-REQ-763:** Supersession shall identify risk impact.
- **ASM-REQ-764:** Supersession shall identify validation impact.
- **ASM-REQ-765:** Supersession shall identify reliance migration.
- **ASM-REQ-766:** Superseded assumptions shall remain immutable.
- **ASM-REQ-767:** Superseded assumptions shall remain retrievable.
- **ASM-REQ-768:** Successor assumptions shall receive independent validation.
- **ASM-REQ-769:** Supersession shall not transfer validation automatically.
- **ASM-REQ-770:** Supersession shall not transfer reliance approval automatically.
- **ASM-REQ-771:** Partial supersession shall identify retained scope.
- **ASM-REQ-772:** Ambiguous supersession shall block reliance.
- **ASM-REQ-773:** Supersession shall update indexes.
- **ASM-REQ-774:** Supersession shall update dependency links.
- **ASM-REQ-775:** Supersession corrections shall be additive.
- **ASM-REQ-776:** Supersession history shall remain retrievable.

## 36. Closure and Archival

- **ASM-REQ-777:** Closure shall require a disposition.
- **ASM-REQ-778:** Closure shall identify final status.
- **ASM-REQ-779:** Closure shall identify validation outcome.
- **ASM-REQ-780:** Closure shall identify evidence.
- **ASM-REQ-781:** Closure shall identify confidence.
- **ASM-REQ-782:** Closure shall identify residual uncertainty.
- **ASM-REQ-783:** Closure shall identify risk impact.
- **ASM-REQ-784:** Closure shall identify architecture impact.
- **ASM-REQ-785:** Closure shall identify resulting decisions.
- **ASM-REQ-786:** Closure shall identify resulting changes.
- **ASM-REQ-787:** Closure shall identify remaining reliance.
- **ASM-REQ-788:** Closure shall identify owner and authority.
- **ASM-REQ-789:** Closure shall identify date.
- **ASM-REQ-790:** Closure shall not fabricate validation.
- **ASM-REQ-791:** Closure shall not hide Not Supported outcomes.
- **ASM-REQ-792:** Closed assumptions shall remain retrievable.
- **ASM-REQ-793:** Closed assumptions shall be read-only except for additive corrections.
- **ASM-REQ-794:** Archived assumptions shall preserve evidence.
- **ASM-REQ-795:** Archived assumptions shall preserve versions.
- **ASM-REQ-796:** Archived assumptions shall preserve links.
- **ASM-REQ-797:** Archive migration shall preserve identifiers.
- **ASM-REQ-798:** Archive restoration shall not reactivate status silently.
- **ASM-REQ-799:** Archive corruption shall trigger incident handling.
- **ASM-REQ-800:** Closure and archival history shall remain auditable.
- **ASM-REQ-801:** Disposition shall not accept risk automatically.

## 37. Corrections and Immutable History

- **ASM-REQ-802:** Assumption records shall maintain immutable history.
- **ASM-REQ-803:** Corrections shall be additive.
- **ASM-REQ-804:** Corrections shall identify the original error.
- **ASM-REQ-805:** Corrections shall identify authority.
- **ASM-REQ-806:** Corrections shall identify date.
- **ASM-REQ-807:** Corrections shall identify affected fields.
- **ASM-REQ-808:** Corrections shall identify impact.
- **ASM-REQ-809:** Corrections shall not erase prior statements.
- **ASM-REQ-810:** Corrections shall not fabricate prior evidence.
- **ASM-REQ-811:** Corrections shall not fabricate prior confidence.
- **ASM-REQ-812:** Corrections shall not fabricate prior validation.
- **ASM-REQ-813:** Corrections shall not fabricate prior review dates.
- **ASM-REQ-814:** Corrections shall not reuse identifiers.
- **ASM-REQ-815:** Material statement correction may require a successor assumption.
- **ASM-REQ-816:** Evidence corrections shall preserve prior references.
- **ASM-REQ-817:** Status corrections shall preserve transition history.
- **ASM-REQ-818:** Owner corrections shall preserve prior accountability.
- **ASM-REQ-819:** Public corrections shall follow release policy.
- **ASM-REQ-820:** Correction records shall remain retrievable.
- **ASM-REQ-821:** Silent rewriting shall be prohibited.

## 38. Metrics and Reporting

- **ASM-REQ-822:** Register reporting shall identify active assumptions.
- **ASM-REQ-823:** Register reporting shall identify assumptions by status.
- **ASM-REQ-824:** Register reporting shall identify assumptions by criticality.
- **ASM-REQ-825:** Register reporting shall identify assumptions by confidence.
- **ASM-REQ-826:** Register reporting shall identify assumptions by evidence state.
- **ASM-REQ-827:** Register reporting shall identify overdue reviews.
- **ASM-REQ-828:** Register reporting shall identify expired assumptions.
- **ASM-REQ-829:** Register reporting shall identify orphaned assumptions.
- **ASM-REQ-830:** Register reporting shall identify Not Supported assumptions.
- **ASM-REQ-831:** Register reporting shall identify Inconclusive assumptions.
- **ASM-REQ-832:** Register reporting shall identify critical low-confidence assumptions.
- **ASM-REQ-833:** Register reporting shall identify assumptions supporting releases.
- **ASM-REQ-834:** Metrics shall not conceal individual critical assumptions.
- **ASM-REQ-835:** Metrics shall not convert quantity into quality.
- **ASM-REQ-836:** High confidence rates shall not prove architecture correctness.
- **ASM-REQ-837:** Closure rates shall not incentivize premature closure.
- **ASM-REQ-838:** Dashboards shall identify freshness.
- **ASM-REQ-839:** AI-generated summaries shall be verified.
- **ASM-REQ-840:** Reports shall minimize personal data.
- **ASM-REQ-841:** Raw records shall remain available to authorized reviewers.

## 39. Traceability

- **ASM-REQ-842:** Every assumption shall trace to its origin.
- **ASM-REQ-843:** Every assumption shall trace to its owner.
- **ASM-REQ-844:** Every assumption shall trace to affected artifacts.
- **ASM-REQ-845:** Every assumption shall trace to evidence.
- **ASM-REQ-846:** Every assumption shall trace to confidence basis.
- **ASM-REQ-847:** Every assumption shall trace to validation strategy.
- **ASM-REQ-848:** Every assumption shall trace to validation evidence.
- **ASM-REQ-849:** Every assumption shall trace to risks.
- **ASM-REQ-850:** Every assumption shall trace to research items where applicable.
- **ASM-REQ-851:** Every assumption shall trace to decisions where applicable.
- **ASM-REQ-852:** Every assumption shall trace to requirements where applicable.
- **ASM-REQ-853:** Every assumption shall trace to implementation where applicable.
- **ASM-REQ-854:** Every assumption shall trace to tests where applicable.
- **ASM-REQ-855:** Every assumption shall trace to releases where applicable.
- **ASM-REQ-856:** Every successor shall trace to predecessor.
- **ASM-REQ-857:** Every closed assumption shall trace to closure authority.
- **ASM-REQ-858:** Every correction shall trace to the corrected record.
- **ASM-REQ-859:** Traceability shall be bidirectional.
- **ASM-REQ-860:** Broken traceability affecting safety, authority or evidence shall be blocking.
- **ASM-REQ-861:** Traceability records shall not contain secrets.
- **ASM-REQ-862:** Traceability records shall minimize personal and case data.
- **ASM-REQ-863:** Baseline freeze shall validate assumption traceability across documents 01–30.
- **ASM-REQ-864:** Traceability shall identify exact versions.
- **ASM-REQ-865:** Traceability shall identify exact commits where material.
- **ASM-REQ-866:** Traceability history shall remain retrievable.

## 40. Domain-Specific Assumption Controls

These controls ensure that material assumptions in every OBDIA domain have explicit evidence, consequence, validation and governance boundaries.

- **ASM-REQ-867:** An assumption concerning officer-agent binding shall state the assumed condition concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-868:** An assumption concerning officer-agent binding shall identify the rationale concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-869:** An assumption concerning officer-agent binding shall identify the evidence supporting or contradicting one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-870:** An assumption concerning officer-agent binding shall identify the confidence basis concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-871:** An assumption concerning officer-agent binding shall identify the consequence if false concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-872:** An assumption concerning officer-agent binding shall identify affected requirements concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-873:** An assumption concerning officer-agent binding shall identify linked threats and risks concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-874:** An assumption concerning officer-agent binding shall identify dependencies concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-875:** An assumption concerning officer-agent binding shall identify the validation strategy concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-876:** An assumption concerning officer-agent binding shall identify acceptance and falsification criteria concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-877:** An assumption concerning officer-agent binding shall identify safe environments and fixtures concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-878:** An assumption concerning officer-agent binding shall identify prohibited validation methods concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-879:** An assumption concerning officer-agent binding shall identify review and expiry triggers concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-880:** An assumption concerning officer-agent binding shall identify human ownership and approval boundaries concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-881:** An assumption concerning officer-agent binding shall identify architecture and implementation impact concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-882:** An assumption concerning officer-agent binding shall identify test and release impact concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-883:** An assumption concerning officer-agent binding shall identify privacy and fundamental-rights impact concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-884:** An assumption concerning officer-agent binding shall identify closure, supersession and archival evidence concerning one accountable officer, one issued agent identity and institutional issuance.
- **ASM-REQ-885:** An assumption concerning machine identity shall state the assumed condition concerning human, agent, workload and connector identity separation.
- **ASM-REQ-886:** An assumption concerning machine identity shall identify the rationale concerning human, agent, workload and connector identity separation.
- **ASM-REQ-887:** An assumption concerning machine identity shall identify the evidence supporting or contradicting human, agent, workload and connector identity separation.
- **ASM-REQ-888:** An assumption concerning machine identity shall identify the confidence basis concerning human, agent, workload and connector identity separation.
- **ASM-REQ-889:** An assumption concerning machine identity shall identify the consequence if false concerning human, agent, workload and connector identity separation.
- **ASM-REQ-890:** An assumption concerning machine identity shall identify affected requirements concerning human, agent, workload and connector identity separation.
- **ASM-REQ-891:** An assumption concerning machine identity shall identify linked threats and risks concerning human, agent, workload and connector identity separation.
- **ASM-REQ-892:** An assumption concerning machine identity shall identify dependencies concerning human, agent, workload and connector identity separation.
- **ASM-REQ-893:** An assumption concerning machine identity shall identify the validation strategy concerning human, agent, workload and connector identity separation.
- **ASM-REQ-894:** An assumption concerning machine identity shall identify acceptance and falsification criteria concerning human, agent, workload and connector identity separation.
- **ASM-REQ-895:** An assumption concerning machine identity shall identify safe environments and fixtures concerning human, agent, workload and connector identity separation.
- **ASM-REQ-896:** An assumption concerning machine identity shall identify prohibited validation methods concerning human, agent, workload and connector identity separation.
- **ASM-REQ-897:** An assumption concerning machine identity shall identify review and expiry triggers concerning human, agent, workload and connector identity separation.
- **ASM-REQ-898:** An assumption concerning machine identity shall identify human ownership and approval boundaries concerning human, agent, workload and connector identity separation.
- **ASM-REQ-899:** An assumption concerning machine identity shall identify architecture and implementation impact concerning human, agent, workload and connector identity separation.
- **ASM-REQ-900:** An assumption concerning machine identity shall identify test and release impact concerning human, agent, workload and connector identity separation.
- **ASM-REQ-901:** An assumption concerning machine identity shall identify privacy and fundamental-rights impact concerning human, agent, workload and connector identity separation.
- **ASM-REQ-902:** An assumption concerning machine identity shall identify closure, supersession and archival evidence concerning human, agent, workload and connector identity separation.
- **ASM-REQ-903:** An assumption concerning authorization shall state the assumed condition concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-904:** An assumption concerning authorization shall identify the rationale concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-905:** An assumption concerning authorization shall identify the evidence supporting or contradicting case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-906:** An assumption concerning authorization shall identify the confidence basis concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-907:** An assumption concerning authorization shall identify the consequence if false concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-908:** An assumption concerning authorization shall identify affected requirements concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-909:** An assumption concerning authorization shall identify linked threats and risks concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-910:** An assumption concerning authorization shall identify dependencies concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-911:** An assumption concerning authorization shall identify the validation strategy concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-912:** An assumption concerning authorization shall identify acceptance and falsification criteria concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-913:** An assumption concerning authorization shall identify safe environments and fixtures concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-914:** An assumption concerning authorization shall identify prohibited validation methods concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-915:** An assumption concerning authorization shall identify review and expiry triggers concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-916:** An assumption concerning authorization shall identify human ownership and approval boundaries concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-917:** An assumption concerning authorization shall identify architecture and implementation impact concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-918:** An assumption concerning authorization shall identify test and release impact concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-919:** An assumption concerning authorization shall identify privacy and fundamental-rights impact concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-920:** An assumption concerning authorization shall identify closure, supersession and archival evidence concerning case, purpose, jurisdiction, time, tool and data scope.
- **ASM-REQ-921:** An assumption concerning agent lifecycle shall state the assumed condition concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-922:** An assumption concerning agent lifecycle shall identify the rationale concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-923:** An assumption concerning agent lifecycle shall identify the evidence supporting or contradicting provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-924:** An assumption concerning agent lifecycle shall identify the confidence basis concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-925:** An assumption concerning agent lifecycle shall identify the consequence if false concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-926:** An assumption concerning agent lifecycle shall identify affected requirements concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-927:** An assumption concerning agent lifecycle shall identify linked threats and risks concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-928:** An assumption concerning agent lifecycle shall identify dependencies concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-929:** An assumption concerning agent lifecycle shall identify the validation strategy concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-930:** An assumption concerning agent lifecycle shall identify acceptance and falsification criteria concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-931:** An assumption concerning agent lifecycle shall identify safe environments and fixtures concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-932:** An assumption concerning agent lifecycle shall identify prohibited validation methods concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-933:** An assumption concerning agent lifecycle shall identify review and expiry triggers concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-934:** An assumption concerning agent lifecycle shall identify human ownership and approval boundaries concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-935:** An assumption concerning agent lifecycle shall identify architecture and implementation impact concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-936:** An assumption concerning agent lifecycle shall identify test and release impact concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-937:** An assumption concerning agent lifecycle shall identify privacy and fundamental-rights impact concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-938:** An assumption concerning agent lifecycle shall identify closure, supersession and archival evidence concerning provisioning, binding, activation, suspension, revocation and archival.
- **ASM-REQ-939:** An assumption concerning evidence integrity shall state the assumed condition concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-940:** An assumption concerning evidence integrity shall identify the rationale concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-941:** An assumption concerning evidence integrity shall identify the evidence supporting or contradicting original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-942:** An assumption concerning evidence integrity shall identify the confidence basis concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-943:** An assumption concerning evidence integrity shall identify the consequence if false concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-944:** An assumption concerning evidence integrity shall identify affected requirements concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-945:** An assumption concerning evidence integrity shall identify linked threats and risks concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-946:** An assumption concerning evidence integrity shall identify dependencies concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-947:** An assumption concerning evidence integrity shall identify the validation strategy concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-948:** An assumption concerning evidence integrity shall identify acceptance and falsification criteria concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-949:** An assumption concerning evidence integrity shall identify safe environments and fixtures concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-950:** An assumption concerning evidence integrity shall identify prohibited validation methods concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-951:** An assumption concerning evidence integrity shall identify review and expiry triggers concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-952:** An assumption concerning evidence integrity shall identify human ownership and approval boundaries concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-953:** An assumption concerning evidence integrity shall identify architecture and implementation impact concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-954:** An assumption concerning evidence integrity shall identify test and release impact concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-955:** An assumption concerning evidence integrity shall identify privacy and fundamental-rights impact concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-956:** An assumption concerning evidence integrity shall identify closure, supersession and archival evidence concerning original bytes, provenance, custody, transformations and exports.
- **ASM-REQ-957:** An assumption concerning connector security shall state the assumed condition concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-958:** An assumption concerning connector security shall identify the rationale concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-959:** An assumption concerning connector security shall identify the evidence supporting or contradicting authentication, egress, validation, isolation and revocation.
- **ASM-REQ-960:** An assumption concerning connector security shall identify the confidence basis concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-961:** An assumption concerning connector security shall identify the consequence if false concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-962:** An assumption concerning connector security shall identify affected requirements concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-963:** An assumption concerning connector security shall identify linked threats and risks concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-964:** An assumption concerning connector security shall identify dependencies concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-965:** An assumption concerning connector security shall identify the validation strategy concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-966:** An assumption concerning connector security shall identify acceptance and falsification criteria concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-967:** An assumption concerning connector security shall identify safe environments and fixtures concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-968:** An assumption concerning connector security shall identify prohibited validation methods concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-969:** An assumption concerning connector security shall identify review and expiry triggers concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-970:** An assumption concerning connector security shall identify human ownership and approval boundaries concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-971:** An assumption concerning connector security shall identify architecture and implementation impact concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-972:** An assumption concerning connector security shall identify test and release impact concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-973:** An assumption concerning connector security shall identify privacy and fundamental-rights impact concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-974:** An assumption concerning connector security shall identify closure, supersession and archival evidence concerning authentication, egress, validation, isolation and revocation.
- **ASM-REQ-975:** An assumption concerning prompt injection shall state the assumed condition concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-976:** An assumption concerning prompt injection shall identify the rationale concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-977:** An assumption concerning prompt injection shall identify the evidence supporting or contradicting untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-978:** An assumption concerning prompt injection shall identify the confidence basis concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-979:** An assumption concerning prompt injection shall identify the consequence if false concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-980:** An assumption concerning prompt injection shall identify affected requirements concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-981:** An assumption concerning prompt injection shall identify linked threats and risks concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-982:** An assumption concerning prompt injection shall identify dependencies concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-983:** An assumption concerning prompt injection shall identify the validation strategy concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-984:** An assumption concerning prompt injection shall identify acceptance and falsification criteria concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-985:** An assumption concerning prompt injection shall identify safe environments and fixtures concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-986:** An assumption concerning prompt injection shall identify prohibited validation methods concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-987:** An assumption concerning prompt injection shall identify review and expiry triggers concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-988:** An assumption concerning prompt injection shall identify human ownership and approval boundaries concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-989:** An assumption concerning prompt injection shall identify architecture and implementation impact concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-990:** An assumption concerning prompt injection shall identify test and release impact concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-991:** An assumption concerning prompt injection shall identify privacy and fundamental-rights impact concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-992:** An assumption concerning prompt injection shall identify closure, supersession and archival evidence concerning untrusted content, retrieval, memory and tool invocation.
- **ASM-REQ-993:** An assumption concerning model behavior shall state the assumed condition concerning versioning, drift, uncertainty and reproducibility.
- **ASM-REQ-994:** An assumption concerning model behavior shall identify the rationale concerning versioning, drift, uncertainty and reproducibility.
- **ASM-REQ-995:** An assumption concerning model behavior shall identify the evidence supporting or contradicting versioning, drift, uncertainty and reproducibility.
- **ASM-REQ-996:** An assumption concerning model behavior shall identify the confidence basis concerning versioning, drift, uncertainty and reproducibility.
- **ASM-REQ-997:** An assumption concerning model behavior shall identify the consequence if false concerning versioning, drift, uncertainty and reproducibility.
- **ASM-REQ-998:** An assumption concerning model behavior shall identify affected requirements concerning versioning, drift, uncertainty and reproducibility.
- **ASM-REQ-999:** An assumption concerning model behavior shall identify linked threats and risks concerning versioning, drift, uncertainty and reproducibility.

## 41. Initial Synthetic Assumptions Register

The following records are initial research assumptions. None is validated, approved as architecture, accepted as risk or authorized for operational reliance.


### `ASM-001`

| Field | Value |
|---|---|
| **Statement** | An institution can issue a unique agent identity for one officer without transferring the officer's reusable credentials. |
| **Category** | Identity |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Credential transfer would undermine accountability and enable impersonation. |
| **Research Links** | RSCH-ITEM-001; RSCH-ITEM-022 |
| **Risk Links** | AI-RISK-001; AI-RISK-002 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-002`

| Field | Value |
|---|---|
| **Statement** | The institution can maintain an authoritative binding record linking one agent identity to one officer and one institution. |
| **Category** | Identity |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Binding disputes could invalidate attribution and authorization. |
| **Research Links** | RSCH-ITEM-001 |
| **Risk Links** | AI-RISK-001 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-003`

| Field | Value |
|---|---|
| **Statement** | Revocation can propagate to all material enforcement points within a measurable bounded interval. |
| **Category** | Authorization |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Revoked agents or connectors could continue operating. |
| **Research Links** | RSCH-ITEM-003 |
| **Risk Links** | AI-RISK-004 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-004`

| Field | Value |
|---|---|
| **Statement** | Authorization attributes can be refreshed during long-running operations. |
| **Category** | Authorization |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Stale case, purpose, time or jurisdiction context could permit unauthorized action. |
| **Research Links** | RSCH-ITEM-002 |
| **Risk Links** | AI-RISK-003 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-005`

| Field | Value |
|---|---|
| **Statement** | Restrictive lifecycle transitions can win during concurrent or partitioned state updates. |
| **Category** | Agent Lifecycle |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | An agent could remain active after suspension or revocation. |
| **Research Links** | RSCH-ITEM-004 |
| **Risk Links** | AI-RISK-005 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-006`

| Field | Value |
|---|---|
| **Statement** | Original evidence bytes can remain immutable while metadata corrections are recorded additively. |
| **Category** | Evidence |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Corrections could overwrite or obscure original evidence. |
| **Research Links** | RSCH-ITEM-005; RSCH-ITEM-006 |
| **Risk Links** | AI-RISK-008 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-007`

| Field | Value |
|---|---|
| **Statement** | A deterministic canonicalization profile can be defined for structured evidence metadata. |
| **Category** | Evidence |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Different systems could calculate inconsistent integrity values. |
| **Research Links** | RSCH-ITEM-005 |
| **Risk Links** | AI-RISK-009 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-008`

| Field | Value |
|---|---|
| **Statement** | Derived AI and non-AI artifacts can preserve complete parent-child lineage. |
| **Category** | Evidence |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Analysis could be detached from source evidence and transformations. |
| **Research Links** | RSCH-ITEM-006 |
| **Risk Links** | AI-RISK-010 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-009`

| Field | Value |
|---|---|
| **Statement** | A proportionate tamper-evident audit mechanism can be demonstrated in a local laboratory. |
| **Category** | Security |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Audit changes could be undetectable or assurance claims misleading. |
| **Research Links** | RSCH-ITEM-007; RSCH-ITEM-020 |
| **Risk Links** | AI-RISK-011 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-010`

| Field | Value |
|---|---|
| **Statement** | Connector content can be isolated so embedded instructions cannot override policy or trigger unauthorized tools. |
| **Category** | AI Security |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Indirect prompt injection could cause unauthorized disclosure or action. |
| **Research Links** | RSCH-ITEM-008 |
| **Risk Links** | AI-RISK-012 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-011`

| Field | Value |
|---|---|
| **Statement** | Connector failures can be isolated by process, credential, network and queue boundaries. |
| **Category** | Connector |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | One connector failure could contaminate cases or agent sessions. |
| **Research Links** | RSCH-ITEM-009 |
| **Risk Links** | AI-RISK-013 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-012`

| Field | Value |
|---|---|
| **Statement** | Model uncertainty can be communicated without implying factual truth, integrity or admissibility. |
| **Category** | AI Governance |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Users could over-rely on model output. |
| **Research Links** | RSCH-ITEM-010 |
| **Risk Links** | AI-RISK-014 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-013`

| Field | Value |
|---|---|
| **Statement** | Consequential actions can be classified into approval and step-up tiers. |
| **Category** | Human Oversight |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | High-impact actions could occur without adequate human control. |
| **Research Links** | RSCH-ITEM-011 |
| **Risk Links** | AI-RISK-015 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-014`

| Field | Value |
|---|---|
| **Statement** | Purpose and case boundaries can be enforced consistently across storage, memory, logs and connectors. |
| **Category** | Privacy |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Cross-case or cross-purpose contamination could occur. |
| **Research Links** | RSCH-ITEM-012; RSCH-ITEM-013 |
| **Risk Links** | AI-RISK-016 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-015`

| Field | Value |
|---|---|
| **Statement** | Jurisdiction can be represented as a policy attribute without encoding autonomous legal conclusions. |
| **Category** | Governance |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | The system could falsely claim lawful authority. |
| **Research Links** | RSCH-ITEM-014 |
| **Risk Links** | AI-RISK-017 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-016`

| Field | Value |
|---|---|
| **Statement** | Minimal telemetry can support accountability and incident response without excessive personal or case data. |
| **Category** | Privacy |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Monitoring could become disproportionate surveillance or omit critical evidence. |
| **Research Links** | RSCH-ITEM-015 |
| **Risk Links** | AI-RISK-018 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-017`

| Field | Value |
|---|---|
| **Statement** | Model and provider changes can be detected or bounded sufficiently to qualify prior validation evidence. |
| **Category** | AI Assurance |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Validation could silently become stale. |
| **Research Links** | RSCH-ITEM-016 |
| **Risk Links** | AI-RISK-019 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-018`

| Field | Value |
|---|---|
| **Statement** | AI-assisted analysis can be reproduced or explicitly qualified using recorded versions and settings. |
| **Category** | Testing |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Results could not be independently reviewed or repeated. |
| **Research Links** | RSCH-ITEM-017 |
| **Risk Links** | AI-RISK-020 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-019`

| Field | Value |
|---|---|
| **Statement** | The qualitative 5×5 method can be applied consistently without implying mathematical probability. |
| **Category** | Risk |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Medium` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Risk priorities could be inconsistent or overstated. |
| **Research Links** | RSCH-ITEM-018 |
| **Risk Links** | AI-RISK-021 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-020`

| Field | Value |
|---|---|
| **Statement** | Authorization invariants can be tested with property-based or formal techniques proportionate to the project. |
| **Category** | Authorization |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Critical policy errors could evade example-based tests. |
| **Research Links** | RSCH-ITEM-019 |
| **Risk Links** | AI-RISK-022 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-021`

| Field | Value |
|---|---|
| **Statement** | Time sources and uncertainty can be preserved for authorization and evidence events. |
| **Category** | Evidence |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Event ordering and validity periods could be disputed. |
| **Research Links** | RSCH-ITEM-021 |
| **Risk Links** | AI-RISK-023 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-022`

| Field | Value |
|---|---|
| **Statement** | Keys can rotate without changing the officer-agent binding identity. |
| **Category** | Identity |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Rotation could break continuity or enable unauthorized rebinding. |
| **Research Links** | RSCH-ITEM-022 |
| **Risk Links** | AI-RISK-024 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-023`

| Field | Value |
|---|---|
| **Statement** | A public research repository can maintain sufficient dependency and build provenance without a proprietary platform. |
| **Category** | Supply Chain |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Medium` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Released artifacts could not be traced to reviewed source. |
| **Research Links** | RSCH-ITEM-023 |
| **Risk Links** | AI-RISK-025 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-024`

| Field | Value |
|---|---|
| **Statement** | Policy rollback can restore a known version without restoring revoked authority. |
| **Category** | Authorization |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Rollback could re-enable unsafe access. |
| **Research Links** | RSCH-ITEM-024 |
| **Risk Links** | AI-RISK-026 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-025`

| Field | Value |
|---|---|
| **Statement** | A high-risk research environment can be isolated sufficiently for safe defensive experimentation. |
| **Category** | Security |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Research could escape containment or contact prohibited infrastructure. |
| **Research Links** | RSCH-ITEM-025 |
| **Risk Links** | AI-RISK-027 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-026`

| Field | Value |
|---|---|
| **Statement** | Testnet blockchain observations can be represented as synthetic evidence with explicit finality limits. |
| **Category** | Web3 |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Medium` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Testnet artifacts could be mistaken for real-world evidence. |
| **Research Links** | RSCH-ITEM-026; RSCH-ITEM-028 |
| **Risk Links** | AI-RISK-028 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-027`

| Field | Value |
|---|---|
| **Statement** | Smart-contract inspection and simulation can remain read-only without autonomous signing. |
| **Category** | Web3 |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Research could create unintended state changes or transactions. |
| **Research Links** | RSCH-ITEM-027 |
| **Risk Links** | AI-RISK-029 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-028`

| Field | Value |
|---|---|
| **Statement** | Authenticated deep-web research can use dedicated sandbox accounts and scoped sessions. |
| **Category** | Connector |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Research could violate account, purpose or data boundaries. |
| **Research Links** | RSCH-ITEM-029 |
| **Risk Links** | AI-RISK-030 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-029`

| Field | Value |
|---|---|
| **Statement** | A Dark Web simulation can be sufficiently realistic without uncontrolled criminal-infrastructure interaction. |
| **Category** | Research |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Critical` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Research could become unsafe, unlawful or misleading. |
| **Research Links** | RSCH-ITEM-030 |
| **Risk Links** | AI-RISK-031 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-030`

| Field | Value |
|---|---|
| **Statement** | Compliance mappings can express applicability, strength, gaps and uncertainty without becoming legal conclusions. |
| **Category** | Compliance |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Mappings could be misrepresented as proof of compliance. |
| **Research Links** | RSCH-ITEM-031 |
| **Risk Links** | AI-RISK-032 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-031`

| Field | Value |
|---|---|
| **Statement** | A fully synthetic scenario can demonstrate officer binding, authorization and evidence integrity credibly. |
| **Category** | Publication |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Portfolio artifacts could overstate operational capability or use sensitive data. |
| **Research Links** | RSCH-ITEM-032 |
| **Risk Links** | AI-RISK-033 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-032`

| Field | Value |
|---|---|
| **Statement** | External specialist review can be scoped without implying certification or institutional endorsement. |
| **Category** | Governance |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `Medium` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Review claims could mislead readers. |
| **Research Links** | RSCH-ITEM-033 |
| **Risk Links** | AI-RISK-034 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-033`

| Field | Value |
|---|---|
| **Statement** | A machine-readable baseline manifest can represent all Knowledge Pack versions, hashes, statuses and dependencies. |
| **Category** | Repository |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Baseline integrity and compatibility could not be verified reliably. |
| **Research Links** | RSCH-ITEM-034 |
| **Risk Links** | AI-RISK-035 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-034`

| Field | Value |
|---|---|
| **Statement** | Structural repository checks can detect missing artifacts, duplicate normative sources and invalid numbering. |
| **Category** | Repository |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Repository drift could undermine governance. |
| **Research Links** | RSCH-ITEM-035 |
| **Risk Links** | AI-RISK-036 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-035`

| Field | Value |
|---|---|
| **Statement** | Synthetic fixtures can be realistic enough for security testing without representing real persons or cases. |
| **Category** | Testing |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Tests could be weak or privacy-invasive. |
| **Research Links** | RSCH-ITEM-036 |
| **Risk Links** | AI-RISK-037 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-036`

| Field | Value |
|---|---|
| **Statement** | Incident evidence can be retained without exposing secrets or excessive case data. |
| **Category** | Incident Response |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Containment could destroy evidence or violate privacy. |
| **Research Links** | RSCH-ITEM-037 |
| **Risk Links** | AI-RISK-038 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-037`

| Field | Value |
|---|---|
| **Statement** | Retention and minimization can be balanced through explicit evidence, audit and legal-hold rules. |
| **Category** | Privacy |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Low` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Data could be retained excessively or deleted prematurely. |
| **Research Links** | RSCH-ITEM-038 |
| **Risk Links** | AI-RISK-039 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-038`

| Field | Value |
|---|---|
| **Statement** | Human review of AI-generated code can be evidenced beyond compilation and automated checks. |
| **Category** | Secure Coding |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Unsafe generated code could be merged under superficial assurance. |
| **Research Links** | RSCH-ITEM-039 |
| **Risk Links** | AI-RISK-040 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-039`

| Field | Value |
|---|---|
| **Statement** | User interfaces can communicate human responsibility and agent non-authority clearly. |
| **Category** | Human Oversight |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `Moderate` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | Users could misunderstand the agent as an autonomous authority. |
| **Research Links** | RSCH-ITEM-040 |
| **Risk Links** | AI-RISK-015 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


### `ASM-040`

| Field | Value |
|---|---|
| **Statement** | Knowledge Pack documents 29 and 30 can be consolidated without changing the established 01–30 authority boundary. |
| **Category** | Governance |
| **Rationale** | The current OBDIA design depends on this condition for one or more proposed security, governance, evidence or implementation properties. |
| **Owner** | Research Reviewer |
| **Criticality** | `High` |
| **Confidence** | `High` |
| **Evidence State** | `Sufficient for Planning` from internal normative and research records; external or experimental evidence remains required where applicable. |
| **Current Status** | `Active` |
| **Consequence if False** | The baseline could remain incomplete or internally inconsistent. |
| **Research Links** | OBDIA-RSCH-001 |
| **Risk Links** | AI-RISK-035 |
| **Validation Strategy** | Perform governed source review, threat and risk analysis, controlled synthetic experiment or simulation where applicable, negative testing, reproducibility review and accountable human assessment. |
| **Review Schedule** | Review before affected architecture approval, implementation, validation or release, and upon material evidence, provider, threat, incident or dependency change. |
| **Reliance Boundary** | Temporary planning input only; no operational or normative authority. |


## 42. Minimum Validation Checklist

Before an assumption may be marked `Supported`, confirm:

- [ ] immutable `ASM-NNN` identifier exists;
- [ ] statement is precise, scoped and testable where feasible;
- [ ] rationale identifies why reliance is needed;
- [ ] owner and required reviewers are identified;
- [ ] category and criticality are accurate;
- [ ] affected artifacts and requirements are linked;
- [ ] consequence if false is explicit;
- [ ] risks, threats and research items are linked;
- [ ] evidence source, strength, state and limitations are recorded;
- [ ] confidence and confidence basis are explicit;
- [ ] validation strategy defines method, data, environment and tools;
- [ ] acceptance, partial-support, falsification and inconclusive criteria are defined;
- [ ] security, privacy and ethical constraints are satisfied;
- [ ] no unauthorized or uncontrolled infrastructure is used;
- [ ] raw observations, failures and negative evidence are retained;
- [ ] reproducibility requirements are met or limitations are stated;
- [ ] validator, date and exact evidence are recorded;
- [ ] residual uncertainty is explicit;
- [ ] risk impact is reviewed;
- [ ] architecture and requirements remain unchanged until governed adoption;
- [ ] reliance, expiry and review schedule are current;
- [ ] resulting ADRs, decisions, changes and tests are linked;
- [ ] closure, supersession or archival evidence is complete where applicable;
- [ ] history remains immutable.


## 43. Limitations

- This document does not validate any initial assumption.
- The initial register is synthetic and derived from current project design dependencies.
- Evidence sufficient for planning may remain insufficient for validation.
- A Supported assumption remains scoped and can later become stale, conflicted or false.
- Validation cannot prove universal truth beyond its method, data, environment and time.
- Provider, model, threat, legal and standards changes can invalidate assumptions.
- Assumption confidence is not mathematical probability unless a justified method is recorded.
- An assumption record does not accept risk or create architecture.
- Internal review is not independent certification.
- Documents 29 and 30 remain forward dependencies for decision chronology and compliance mapping.

## 44. Change Control

Every material change shall identify rationale, affected assumptions, evidence, confidence, risks, research, decisions, architecture, security and privacy impact, migration, validation, rollback, authority and version effect.

- Editorial corrections shall use a patch version when meaning is unchanged.
- Backward-compatible substantive additions shall use a minor version.
- Incompatible assumption-governance changes shall use a major version.
- Material methodology decisions shall require an ADR where applicable.
- Changes shall require Project Founder approval.
- Changes shall receive Research Reviewer assessment.
- Security, privacy, implementation and release impacts shall receive specialist review where applicable.
- Changes shall identify affected records, schemas, indexes and automation.
- Changes shall include migration and compatibility analysis.
- Changes shall include validation and rollback analysis.
- Changes shall not retroactively fabricate evidence, confidence or validation.
- Historical assumption records shall not be silently rewritten.
- Assumption and requirement identifiers shall not be reused.
- Migration shall preserve assumption-to-evidence and assumption-to-decision traceability.
- Forward-dependency reconciliation shall occur before this policy becomes Approved.

## 45. Consolidation Record

Version 1.0.0 consolidates the two existing Assumptions Register variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-ASM-001`;
- normalizes the authoritative filename to `28_ASSUMPTIONS_REGISTER.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves Identifier;
- preserves Statement;
- preserves Rationale;
- preserves Evidence;
- preserves Validation plan and Validation strategy;
- preserves Review status, Review schedule and Current status;
- adds ownership, category, affected artifacts, consequence if false, criticality, confidence, evidence state, dependencies, risk, threat, research and decision links;
- adds controlled statuses, validation outcomes, reliance boundaries, expiry, supersession, closure, corrections and immutable history;
- adds identity, authorization, evidence, lifecycle, connector, AI, privacy, software, testing, repository, cloud, web and Web3 controls;
- establishes 40 synthetic initial assumption records;
- adds 999 stable assumption-governance requirements;
- identifies documents 29 and 30 as forward dependencies;
- treats Enterprise and non-Enterprise files as legacy source variants of the same immutable document;
- creates no validated assumption, architecture decision, implementation, risk acceptance, compliance conclusion, publication or operational authority.

## 46. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-06 | Draft | Research Reviewer; approval reserved to Project Founder | Constitutional consolidation of both legacy Assumptions Register variants: preserved every original field; added complete ownership, confidence, consequence, validation, expiry, disposition and traceability governance plus 40 synthetic initial assumptions. |
