# DECISION LOG POLICY

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-DEC-001 |
| **Title** | Decision Log Policy |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define mandatory identification, chronology, classification, authority, rationale, alternatives, impact, evidence, implementation, validation, publication, supersession, reversal, correction, retention and traceability requirements for OBDIA architectural, governance, security and publication decisions. |
| **Scope** | Major and material decisions affecting architecture, governance, security, privacy, identity, authorization, evidence, agent lifecycle, connectors, AI models, software, testing, repository controls, research, assumptions, risks, releases, publication and compliance mapping within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `04_PORTFOLIO_AND_PUBLICATION_POLICY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `12_IMPLEMENTATION_POLICY.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `15_RELEASE_AND_PUBLICATION_POLICY.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; `18_EVIDENCE_MODEL.md`; `19_AGENT_LIFECYCLE_MODEL.md`; `20_CONNECTOR_SECURITY_POLICY.md`; `21_SECURE_CODING_STANDARD.md`; `22_TESTING_STANDARD.md`; `23_DIAGRAM_STANDARD.md`; `24_GITHUB_REPOSITORY_STANDARD.md`; `25_DOCUMENT_VERSIONING_POLICY.md`; `26_AI_RISK_REGISTER.md`; `27_RESEARCH_BACKLOG.md`; `28_ASSUMPTIONS_REGISTER.md`; forward dependency `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-RES-001`; `OBDIA-PUB-001`; `OBDIA-ARCH-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-IMP-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-REL-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001`; `OBDIA-EVID-001`; `OBDIA-AGENT-001`; `OBDIA-CONN-001`; `OBDIA-CODE-001`; `OBDIA-TEST-001`; `OBDIA-DIAG-001`; `OBDIA-GH-001`; `OBDIA-VER-001`; `OBDIA-RISK-001`; `OBDIA-RSCH-001`; `OBDIA-ASM-001` |
| **Cross-References** | ADRs; chronological decision entries; research items; assumptions; risks; threats; requirements; change records; implementation plans; test plans; validation packages; exception records; baseline manifests; releases and compliance mappings |
| **Assumptions** | Decision entries document governed human decisions and their evidence. They do not create authority merely through recording and do not replace the governing ADR, policy, approval or release record. |
| **Constraints** | Decisions shall not authorize canonical prohibitions, independent AI legal authority, unauthorized access, offensive operations, unlawful surveillance, unlawful deanonymization, uncontrolled criminal-infrastructure interaction or misleading implementation, validation, compliance or publication claims. |
| **Security Considerations** | An incomplete, altered, ambiguous or weakly governed decision log can conceal authority, substitute unreviewed decisions, erase dissent, hide implementation impact, restore insecure choices, misstate chronology or undermine auditability and accountability. |
| **Validation Criteria** | Every material decision has an immutable entry identifier, chronological timestamps, decision type, scope, accountable authority, rationale, alternatives, evidence, assumptions, risk and impact analysis, corresponding ADR when applicable, implementation impact, validation needs, effective status, supersession or reversal links and immutable history. |
| **Implementation Relationship** | This policy governs decision records. It does not accept any decision, approve an ADR, authorize implementation, accept risk, establish validation, publish a release or prove compliance. |

---

## 1. Purpose

The non-Enterprise source required every major decision to reference:

- the corresponding ADR; and
- implementation impact.

The Enterprise source additionally required a chronological log of:

- architectural decisions;
- governance decisions;
- security decisions;
- publication decisions.

It also required each entry to reference the corresponding ADR when applicable.

This consolidation preserves every original requirement. It defines the decision log as a chronological, attributable and tamper-evident index of governed decisions. The log does not replace ADRs, normative policies, approvals, risk records, implementation evidence, validation evidence or release decisions.


## 2. Fundamental Decision-Log Rules

- **DEC-REQ-001:** The project shall maintain a chronological decision log.
- **DEC-REQ-002:** The decision log shall include architectural decisions.
- **DEC-REQ-003:** The decision log shall include governance decisions.
- **DEC-REQ-004:** The decision log shall include security decisions.
- **DEC-REQ-005:** The decision log shall include publication decisions.
- **DEC-REQ-006:** Every major decision shall reference the corresponding ADR.
- **DEC-REQ-007:** Every major decision shall record implementation impact.
- **DEC-REQ-008:** Every decision entry shall reference the corresponding ADR when applicable.
- **DEC-REQ-009:** Every material decision shall have an immutable entry identifier.
- **DEC-REQ-010:** Every material decision shall identify the accountable human authority.
- **DEC-REQ-011:** Every material decision shall identify its scope.
- **DEC-REQ-012:** Every material decision shall identify rationale.
- **DEC-REQ-013:** Every material decision shall identify alternatives considered.
- **DEC-REQ-014:** Every material decision shall identify affected artifacts.
- **DEC-REQ-015:** Every material decision shall identify security impact.
- **DEC-REQ-016:** Every material decision shall identify privacy impact where applicable.
- **DEC-REQ-017:** Every material decision shall identify evidence impact where applicable.
- **DEC-REQ-018:** Every material decision shall identify implementation impact.
- **DEC-REQ-019:** Every material decision shall identify testing and validation impact.
- **DEC-REQ-020:** Every material decision shall identify release and publication impact.
- **DEC-REQ-021:** Decision entries shall remain chronological.
- **DEC-REQ-022:** Decision entries shall remain attributable.
- **DEC-REQ-023:** Decision entries shall preserve immutable history.
- **DEC-REQ-024:** A decision entry shall not replace an ADR.
- **DEC-REQ-025:** A decision entry shall not replace an approval record.
- **DEC-REQ-026:** A decision entry shall not replace a risk-acceptance record.
- **DEC-REQ-027:** A decision entry shall not create legal authority.
- **DEC-REQ-028:** An AI system shall not make or approve a material decision.
- **DEC-REQ-029:** Repository merge shall not establish decision acceptance.
- **DEC-REQ-030:** Decision history shall not be silently rewritten.

## 3. Authority and Decision Boundaries

- **DEC-REQ-031:** Decision authority shall derive from the governing authority hierarchy.
- **DEC-REQ-032:** The Canonical Project Definition shall override lower-level decisions.
- **DEC-REQ-033:** The Constitution shall override lower-level decisions.
- **DEC-REQ-034:** Approved normative documents shall override conflicting lower-level decisions.
- **DEC-REQ-035:** Accepted ADRs shall remain subordinate to superior normative authority.
- **DEC-REQ-036:** Repository documentation shall remain subordinate to accepted ADRs.
- **DEC-REQ-037:** Implementation choices shall remain subordinate to governing requirements.
- **DEC-REQ-038:** A decision log entry shall not create authority by recency alone.
- **DEC-REQ-039:** A decision log entry shall not create authority by position alone.
- **DEC-REQ-040:** A decision log entry shall not create authority by repository merge alone.
- **DEC-REQ-041:** A decision log entry shall not create authority by publication alone.
- **DEC-REQ-042:** A decision log entry shall not create institutional endorsement.
- **DEC-REQ-043:** A decision log entry shall not create independent AI legal authority.
- **DEC-REQ-044:** A decision log entry shall not authorize access by implication.
- **DEC-REQ-045:** A decision log entry shall not authorize evidence acquisition by implication.
- **DEC-REQ-046:** A decision log entry shall not accept risk by implication.
- **DEC-REQ-047:** A decision log entry shall not establish compliance.
- **DEC-REQ-048:** Decision authority shall be explicit.
- **DEC-REQ-049:** Decision authority shall be attributable.
- **DEC-REQ-050:** Decision authority shall be within the decision-maker's delegated competence.
- **DEC-REQ-051:** Decisions outside delegated authority shall be rejected or escalated.
- **DEC-REQ-052:** Conflicting authority claims shall block implementation.
- **DEC-REQ-053:** Emergency authority shall be narrow and time-bounded.
- **DEC-REQ-054:** Role concentration shall be disclosed.
- **DEC-REQ-055:** Internal review shall not be represented as independent external assurance.

## 4. Decision Entry Identifiers

- **DEC-REQ-056:** Decision entries shall use `DEC-ENTRY-NNN` identifiers.
- **DEC-REQ-057:** Decision entry identifiers shall be immutable.
- **DEC-REQ-058:** Decision entry identifiers shall not be reused.
- **DEC-REQ-059:** Decision entry identifiers shall not encode status.
- **DEC-REQ-060:** Decision entry identifiers shall not encode decision outcome.
- **DEC-REQ-061:** Decision entry identifiers shall not encode priority.
- **DEC-REQ-062:** Decision entry identifiers shall not encode owner.
- **DEC-REQ-063:** Decision entry identifiers shall conform to `OBDIA-NAME-001`.
- **DEC-REQ-064:** Duplicate decision identifiers shall be rejected.
- **DEC-REQ-065:** Superseded decision entries shall retain their identifiers.
- **DEC-REQ-066:** Reversed decision entries shall retain their identifiers.
- **DEC-REQ-067:** Archived decision entries shall retain their identifiers.
- **DEC-REQ-068:** ADR identifiers shall remain distinct from decision entry identifiers.
- **DEC-REQ-069:** Risk identifiers shall remain distinct from decision entry identifiers.
- **DEC-REQ-070:** Assumption identifiers shall remain distinct from decision entry identifiers.
- **DEC-REQ-071:** Research identifiers shall remain distinct from decision entry identifiers.
- **DEC-REQ-072:** Change-record identifiers shall remain distinct from decision entry identifiers.
- **DEC-REQ-073:** Release identifiers shall remain distinct from decision entry identifiers.
- **DEC-REQ-074:** Unknown identifier namespaces shall require review.
- **DEC-REQ-075:** Identifier corrections shall preserve the original record.

## 5. Mandatory Decision Entry Schema

- **DEC-REQ-076:** Every decision entry shall include `Decision Entry ID`.
- **DEC-REQ-077:** Every decision entry shall include `Title`.
- **DEC-REQ-078:** Every decision entry shall include `Decision Type`.
- **DEC-REQ-079:** Every decision entry shall include `Decision Status`.
- **DEC-REQ-080:** Every decision entry shall include `Decision Statement`.
- **DEC-REQ-081:** Every decision entry shall include `Scope`.
- **DEC-REQ-082:** Every decision entry shall include `Authority`.
- **DEC-REQ-083:** Every decision entry shall include `Owner`.
- **DEC-REQ-084:** Every decision entry shall include `Decision Date`.
- **DEC-REQ-085:** Every decision entry shall include `Recorded Date`.
- **DEC-REQ-086:** Every decision entry shall include `Effective Date` where applicable.
- **DEC-REQ-087:** Every decision entry shall include `Rationale`.
- **DEC-REQ-088:** Every decision entry shall include `Alternatives Considered`.
- **DEC-REQ-089:** Every decision entry shall include `Evidence References`.
- **DEC-REQ-090:** Every decision entry shall include `Source References` where applicable.
- **DEC-REQ-091:** Every decision entry shall include `Research Links` where applicable.
- **DEC-REQ-092:** Every decision entry shall include `Assumption Links` where applicable.
- **DEC-REQ-093:** Every decision entry shall include `Risk Links` where applicable.
- **DEC-REQ-094:** Every decision entry shall include `Threat Links` where applicable.
- **DEC-REQ-095:** Every decision entry shall include `ADR Reference` when applicable.
- **DEC-REQ-096:** Every major decision shall include `ADR Reference`.
- **DEC-REQ-097:** Every decision entry shall include `Affected Documents`.
- **DEC-REQ-098:** Every decision entry shall include `Affected Requirements` where applicable.
- **DEC-REQ-099:** Every decision entry shall include `Security Impact`.
- **DEC-REQ-100:** Every decision entry shall include `Privacy and Rights Impact` where applicable.
- **DEC-REQ-101:** Every decision entry shall include `Evidence Impact` where applicable.
- **DEC-REQ-102:** Every decision entry shall include `Implementation Impact`.
- **DEC-REQ-103:** Every decision entry shall include `Testing and Validation Impact`.
- **DEC-REQ-104:** Every decision entry shall include `Release and Publication Impact`.
- **DEC-REQ-105:** Every decision entry shall include `Compatibility Impact` where applicable.
- **DEC-REQ-106:** Every decision entry shall include `Migration` where applicable.
- **DEC-REQ-107:** Every decision entry shall include `Rollback or Reversal` where applicable.
- **DEC-REQ-108:** Every decision entry shall include `Known Limitations`.
- **DEC-REQ-109:** Every decision entry shall include `Reviewers`.
- **DEC-REQ-110:** Every decision entry shall include `Approval Evidence` where accepted.
- **DEC-REQ-111:** Every decision entry shall include `Implementation Evidence` where implemented.
- **DEC-REQ-112:** Every decision entry shall include `Validation Evidence` where validated.
- **DEC-REQ-113:** Every superseded entry shall include `Successor`.
- **DEC-REQ-114:** Every reversed entry shall include `Reversal Record`.
- **DEC-REQ-115:** Every decision entry shall maintain `Revision History`.

## 6. Decision Types

- **DEC-REQ-116:** Permitted decision types shall be controlled.
- **DEC-REQ-117:** Architectural decisions shall concern structural or technology architecture.
- **DEC-REQ-118:** Governance decisions shall concern authority, roles, policies or lifecycle.
- **DEC-REQ-119:** Security decisions shall concern threats, controls, identity or containment.
- **DEC-REQ-120:** Privacy decisions shall concern data, purpose, retention or rights.
- **DEC-REQ-121:** Evidence decisions shall concern acquisition, provenance, custody or export.
- **DEC-REQ-122:** Authorization decisions shall concern policy, attributes or enforcement.
- **DEC-REQ-123:** Agent-lifecycle decisions shall concern provisioning, binding, activation, suspension, revocation or archival.
- **DEC-REQ-124:** Connector decisions shall concern external integration boundaries.
- **DEC-REQ-125:** AI decisions shall concern models, prompts, retrieval, memory or tool use.
- **DEC-REQ-126:** Implementation decisions shall concern code, runtime, dependency or deployment choices.
- **DEC-REQ-127:** Testing decisions shall concern validation methods, fixtures or acceptance.
- **DEC-REQ-128:** Repository decisions shall concern branches, workflows, rulesets or releases.
- **DEC-REQ-129:** Research decisions shall concern methods, scope or disposition.
- **DEC-REQ-130:** Risk decisions shall concern treatment, escalation or acceptance authority.
- **DEC-REQ-131:** Publication decisions shall concern public content, releases or claims.
- **DEC-REQ-132:** Compliance-mapping decisions shall concern applicability and mapping interpretation.
- **DEC-REQ-133:** Incident decisions shall concern containment, recovery or disclosure.
- **DEC-REQ-134:** Emergency decisions shall concern time-bounded urgent action.
- **DEC-REQ-135:** Administrative decisions shall concern non-architectural operational governance.
- **DEC-REQ-136:** A decision may have multiple types when justified.
- **DEC-REQ-137:** Decision type shall not create authority.
- **DEC-REQ-138:** Decision type changes shall be attributable.
- **DEC-REQ-139:** Unknown decision types shall require review.
- **DEC-REQ-140:** Type assignment shall remain consistent with the decision statement.

## 7. Decision Materiality

- **DEC-REQ-141:** Decision materiality shall be assessed.
- **DEC-REQ-142:** Major decisions shall require an ADR.
- **DEC-REQ-143:** Major decisions shall record implementation impact.
- **DEC-REQ-144:** Materiality shall consider authority impact.
- **DEC-REQ-145:** Materiality shall consider architectural impact.
- **DEC-REQ-146:** Materiality shall consider security impact.
- **DEC-REQ-147:** Materiality shall consider privacy and rights impact.
- **DEC-REQ-148:** Materiality shall consider evidence impact.
- **DEC-REQ-149:** Materiality shall consider interoperability impact.
- **DEC-REQ-150:** Materiality shall consider implementation impact.
- **DEC-REQ-151:** Materiality shall consider migration impact.
- **DEC-REQ-152:** Materiality shall consider validation impact.
- **DEC-REQ-153:** Materiality shall consider release impact.
- **DEC-REQ-154:** Materiality shall consider public-claim impact.
- **DEC-REQ-155:** Materiality shall consider reversibility.
- **DEC-REQ-156:** Materiality shall consider affected users and systems.
- **DEC-REQ-157:** Materiality shall consider dependency breadth.
- **DEC-REQ-158:** Materiality shall not be reduced to avoid ADR review.
- **DEC-REQ-159:** Unknown materiality shall receive review.
- **DEC-REQ-160:** Materiality decisions shall be attributable.
- **DEC-REQ-161:** Minor administrative decisions may use a decision entry without an ADR when the rationale is recorded.
- **DEC-REQ-162:** Repeated minor decisions with cumulative material impact shall trigger ADR review.
- **DEC-REQ-163:** Emergency decisions with material impact shall receive retrospective ADR review.
- **DEC-REQ-164:** Materiality changes shall update the decision entry.
- **DEC-REQ-165:** Materiality history shall remain retrievable.

## 8. Decision Status Model

- **DEC-REQ-166:** Permitted decision statuses shall be `Proposed`, `Under Review`, `Accepted`, `Rejected`, `Deferred`, `Superseded`, `Reversed` and `Archived`.
- **DEC-REQ-167:** `Proposed` shall mean a decision option has been recorded but not reviewed.
- **DEC-REQ-168:** `Under Review` shall mean applicable evaluation is active.
- **DEC-REQ-169:** `Accepted` shall mean the authorized decision-maker has approved the decision for its stated scope.
- **DEC-REQ-170:** `Rejected` shall mean the proposal was not accepted.
- **DEC-REQ-171:** `Deferred` shall mean no current decision has been made.
- **DEC-REQ-172:** `Superseded` shall mean a successor decision replaces the entry.
- **DEC-REQ-173:** `Reversed` shall mean an authorized later action withdraws the decision.
- **DEC-REQ-174:** `Archived` shall mean the historical record is retained read-only.
- **DEC-REQ-175:** Decision status shall remain distinct from repository lifecycle status.
- **DEC-REQ-176:** Decision status shall remain distinct from ADR status.
- **DEC-REQ-177:** Decision status shall remain distinct from implementation state.
- **DEC-REQ-178:** Decision status shall remain distinct from validation outcome.
- **DEC-REQ-179:** Decision status shall remain distinct from publication state.
- **DEC-REQ-180:** Decision status shall not be inferred from branch state.
- **DEC-REQ-181:** Decision status shall not be inferred from PR state.
- **DEC-REQ-182:** Decision status shall not be inferred from merge state.
- **DEC-REQ-183:** Decision status shall not be inferred from tag existence.
- **DEC-REQ-184:** Status transitions shall be attributable.
- **DEC-REQ-185:** Status transitions shall record rationale.
- **DEC-REQ-186:** Invalid status transitions shall be rejected.
- **DEC-REQ-187:** A Proposed decision shall not become Accepted without authority evidence.
- **DEC-REQ-188:** A Rejected decision shall remain retrievable.
- **DEC-REQ-189:** A Deferred decision shall identify review trigger.
- **DEC-REQ-190:** An Archived decision shall remain read-only.

## 9. Implementation State

- **DEC-REQ-191:** Decision implementation state shall be recorded separately from decision status.
- **DEC-REQ-192:** Permitted implementation states shall be `Not Planned`, `Planned`, `In Progress`, `Implemented`, `Partially Implemented`, `Rolled Back` and `Not Applicable`.
- **DEC-REQ-193:** Implementation state shall identify affected artifacts.
- **DEC-REQ-194:** Implementation state shall identify owner.
- **DEC-REQ-195:** Implementation state shall identify exact versions.
- **DEC-REQ-196:** Implementation state shall identify repository commits where material.
- **DEC-REQ-197:** Implementation state shall identify deployment environment where material.
- **DEC-REQ-198:** Implementation state shall identify implementation evidence.
- **DEC-REQ-199:** Implementation state shall identify incomplete scope.
- **DEC-REQ-200:** Implementation state shall identify rollback evidence where applicable.
- **DEC-REQ-201:** Accepted shall not imply Implemented.
- **DEC-REQ-202:** Implemented shall not imply Validated.
- **DEC-REQ-203:** Partially Implemented shall identify implemented and missing portions.
- **DEC-REQ-204:** Rolled Back shall preserve prior implementation evidence.
- **DEC-REQ-205:** Not Applicable shall include rationale.
- **DEC-REQ-206:** Implementation-state changes shall be attributable.
- **DEC-REQ-207:** Implementation-state changes shall not rewrite the decision.
- **DEC-REQ-208:** Implementation evidence shall conform to `OBDIA-IMP-001`.
- **DEC-REQ-209:** Implementation code shall conform to `OBDIA-CODE-001`.
- **DEC-REQ-210:** Implementation state shall remain auditable.

## 10. Validation State

- **DEC-REQ-211:** Decision validation state shall be recorded separately.
- **DEC-REQ-212:** Permitted validation states shall be `Not Evaluated`, `Planned`, `Under Validation`, `Validated`, `Partially Validated`, `Not Validated` and `Invalidated`.
- **DEC-REQ-213:** Validation state shall identify scope.
- **DEC-REQ-214:** Validation state shall identify evidence.
- **DEC-REQ-215:** Validation state shall identify exact implementation versions.
- **DEC-REQ-216:** Validation state shall identify governing document versions.
- **DEC-REQ-217:** Validation state shall identify evaluator.
- **DEC-REQ-218:** Validation state shall identify date.
- **DEC-REQ-219:** Validation state shall identify limitations.
- **DEC-REQ-220:** Validation state shall identify residual uncertainty.
- **DEC-REQ-221:** Accepted shall not imply Validated.
- **DEC-REQ-222:** Implemented shall not imply Validated.
- **DEC-REQ-223:** Published shall not imply Validated.
- **DEC-REQ-224:** Partially Validated shall identify validated and unvalidated scope.
- **DEC-REQ-225:** Invalidated shall identify trigger and impact.
- **DEC-REQ-226:** Material implementation changes shall invalidate affected validation.
- **DEC-REQ-227:** Material assumption changes shall invalidate affected validation.
- **DEC-REQ-228:** Material provider or model changes shall invalidate affected validation where applicable.
- **DEC-REQ-229:** Validation evidence shall conform to `OBDIA-TEST-001`.
- **DEC-REQ-230:** Validation history shall remain retrievable.

## 11. Chronology and Time

- **DEC-REQ-231:** Decision entries shall be ordered chronologically.
- **DEC-REQ-232:** Chronology shall use decision date.
- **DEC-REQ-233:** Chronology shall retain recorded date.
- **DEC-REQ-234:** Chronology shall retain effective date where applicable.
- **DEC-REQ-235:** Decision date shall identify when authority made the decision.
- **DEC-REQ-236:** Recorded date shall identify when the log entry was created.
- **DEC-REQ-237:** Effective date shall identify when the decision takes effect.
- **DEC-REQ-238:** Decision date and recorded date differences shall be explained where material.
- **DEC-REQ-239:** Retroactive decisions shall identify authority and rationale.
- **DEC-REQ-240:** Emergency decisions shall identify exact time where material.
- **DEC-REQ-241:** Time sources shall be documented where material.
- **DEC-REQ-242:** Known clock uncertainty shall be recorded.
- **DEC-REQ-243:** Chronology shall not rely solely on commit timestamps.
- **DEC-REQ-244:** Chronology shall not be rewritten to improve appearance.
- **DEC-REQ-245:** Late entries shall remain labeled as late entries.
- **DEC-REQ-246:** Unknown dates shall remain unknown.
- **DEC-REQ-247:** Dates shall not be fabricated.
- **DEC-REQ-248:** Time-zone context shall be recorded where ambiguity matters.
- **DEC-REQ-249:** Supersession chronology shall preserve predecessor ordering.
- **DEC-REQ-250:** Reversal chronology shall preserve original and reversal ordering.
- **DEC-REQ-251:** Correction chronology shall preserve original and corrected dates.
- **DEC-REQ-252:** Machine-readable chronology shall match human-readable records.
- **DEC-REQ-253:** Chronological indexes shall be reproducible.
- **DEC-REQ-254:** Chronology failures shall trigger review.
- **DEC-REQ-255:** Decision time history shall remain retrievable.

## 12. Rationale

- **DEC-REQ-256:** Rationale shall explain why the decision is needed.
- **DEC-REQ-257:** Rationale shall identify the problem addressed.
- **DEC-REQ-258:** Rationale shall identify desired outcomes.
- **DEC-REQ-259:** Rationale shall identify constraints.
- **DEC-REQ-260:** Rationale shall identify trade-offs.
- **DEC-REQ-261:** Rationale shall identify evidence relied upon.
- **DEC-REQ-262:** Rationale shall identify assumptions relied upon.
- **DEC-REQ-263:** Rationale shall identify risks accepted, transferred, mitigated or avoided where applicable.
- **DEC-REQ-264:** Rationale shall identify why alternatives were not chosen.
- **DEC-REQ-265:** Rationale shall not merely restate the decision.
- **DEC-REQ-266:** Rationale shall not rely on authority alone.
- **DEC-REQ-267:** Rationale shall not rely on vendor preference alone.
- **DEC-REQ-268:** Rationale shall not rely on model output alone.
- **DEC-REQ-269:** Rationale shall not conceal dissent.
- **DEC-REQ-270:** Rationale shall not fabricate urgency.
- **DEC-REQ-271:** Rationale shall not imply compliance without evidence.
- **DEC-REQ-272:** Rationale shall identify uncertainty.
- **DEC-REQ-273:** Rationale shall identify limitations.
- **DEC-REQ-274:** Rationale changes shall be attributable.
- **DEC-REQ-275:** Material rationale changes shall trigger renewed review.

## 13. Alternatives

- **DEC-REQ-276:** Material decisions shall identify alternatives considered.
- **DEC-REQ-277:** Alternatives shall include maintaining the current state where relevant.
- **DEC-REQ-278:** Alternatives shall include safer or simpler options where relevant.
- **DEC-REQ-279:** Alternatives shall identify benefits.
- **DEC-REQ-280:** Alternatives shall identify costs.
- **DEC-REQ-281:** Alternatives shall identify security impact.
- **DEC-REQ-282:** Alternatives shall identify privacy impact.
- **DEC-REQ-283:** Alternatives shall identify evidence impact.
- **DEC-REQ-284:** Alternatives shall identify implementation impact.
- **DEC-REQ-285:** Alternatives shall identify validation impact.
- **DEC-REQ-286:** Alternatives shall identify migration impact.
- **DEC-REQ-287:** Alternatives shall identify rollback impact.
- **DEC-REQ-288:** Alternatives shall identify dependencies.
- **DEC-REQ-289:** Alternatives shall identify assumptions.
- **DEC-REQ-290:** Alternatives shall identify risks.
- **DEC-REQ-291:** Rejected alternatives shall preserve rationale.
- **DEC-REQ-292:** Alternatives shall not be fabricated after acceptance.
- **DEC-REQ-293:** Material new alternatives shall trigger review.
- **DEC-REQ-294:** Alternative comparison shall use consistent criteria.
- **DEC-REQ-295:** Alternative history shall remain retrievable.

## 14. ADR Integration

- **DEC-REQ-296:** Every major decision shall reference a corresponding ADR.
- **DEC-REQ-297:** Every architectural decision with material impact shall reference an ADR.
- **DEC-REQ-298:** Every material governance decision shall reference an ADR or superior normative change record as applicable.
- **DEC-REQ-299:** Every material security decision shall reference an ADR where architectural choice is involved.
- **DEC-REQ-300:** Every publication decision shall reference an ADR only when architectural or governance trade-offs require it.
- **DEC-REQ-301:** Decision entries shall identify ADR identifier.
- **DEC-REQ-302:** Decision entries shall identify ADR version or commit where material.
- **DEC-REQ-303:** Decision entries shall identify ADR status.
- **DEC-REQ-304:** Decision entries shall not represent a Proposed ADR as Accepted.
- **DEC-REQ-305:** Decision entries shall not replace ADR context.
- **DEC-REQ-306:** Decision entries shall not replace ADR alternatives.
- **DEC-REQ-307:** Decision entries shall not replace ADR consequences.
- **DEC-REQ-308:** ADR acceptance shall update the decision entry.
- **DEC-REQ-309:** ADR rejection shall update the decision entry.
- **DEC-REQ-310:** ADR supersession shall trigger decision review.
- **DEC-REQ-311:** ADR reversal shall trigger decision review.
- **DEC-REQ-312:** Decision and ADR statements shall remain consistent.
- **DEC-REQ-313:** Conflicts shall resolve in favor of the higher authority.
- **DEC-REQ-314:** Broken ADR links shall block material decision closure.
- **DEC-REQ-315:** ADR integration shall conform to `OBDIA-ADR-001`.

## 15. Research Integration

- **DEC-REQ-316:** Decisions shall identify supporting research items.
- **DEC-REQ-317:** Decisions shall identify research validation outcomes.
- **DEC-REQ-318:** Decisions shall identify research limitations.
- **DEC-REQ-319:** Decisions shall identify unresolved research questions.
- **DEC-REQ-320:** Decisions shall not present hypotheses as facts.
- **DEC-REQ-321:** Decisions shall not present Inconclusive research as Supported.
- **DEC-REQ-322:** Decisions shall not hide negative research results.
- **DEC-REQ-323:** Research priority shall not create decision authority.
- **DEC-REQ-324:** Research validation shall not accept a decision automatically.
- **DEC-REQ-325:** Decision authority shall assess research evidence.
- **DEC-REQ-326:** Decision rejection shall not erase research evidence.
- **DEC-REQ-327:** Decision deferral shall identify research gaps.
- **DEC-REQ-328:** Decision reversal shall trigger research impact review where material.
- **DEC-REQ-329:** Research links shall identify exact item identifiers.
- **DEC-REQ-330:** Research evidence shall identify exact versions.
- **DEC-REQ-331:** Research-derived changes shall preserve source traceability.
- **DEC-REQ-332:** AI-generated research summaries shall receive human verification.
- **DEC-REQ-333:** Research integration shall conform to `OBDIA-RSCH-001`.
- **DEC-REQ-334:** Source integration shall conform to `OBDIA-RES-001`.
- **DEC-REQ-335:** Research-decision history shall remain retrievable.

## 16. Assumption Integration

- **DEC-REQ-336:** Decisions shall identify assumptions relied upon.
- **DEC-REQ-337:** Assumption links shall use immutable identifiers.
- **DEC-REQ-338:** Decisions shall identify assumption status.
- **DEC-REQ-339:** Decisions shall identify assumption confidence.
- **DEC-REQ-340:** Decisions shall identify evidence state.
- **DEC-REQ-341:** Decisions shall identify consequence if assumptions are false.
- **DEC-REQ-342:** Critical unsupported assumptions shall block decision acceptance.
- **DEC-REQ-343:** Expired assumptions shall not support new decisions.
- **DEC-REQ-344:** Not Supported assumptions shall trigger decision review.
- **DEC-REQ-345:** Inconclusive assumptions shall constrain decision scope.
- **DEC-REQ-346:** Assumption supersession shall update decision links.
- **DEC-REQ-347:** Assumption invalidation shall trigger impact review.
- **DEC-REQ-348:** Decision rationale shall not convert assumptions into facts.
- **DEC-REQ-349:** Decision acceptance shall not validate assumptions automatically.
- **DEC-REQ-350:** Assumption reliance shall identify compensating controls.
- **DEC-REQ-351:** Assumption reliance shall identify expiry or review trigger.
- **DEC-REQ-352:** AI systems shall not approve assumption-dependent decisions.
- **DEC-REQ-353:** Assumption integration shall conform to `OBDIA-ASM-001`.
- **DEC-REQ-354:** Broken assumption links shall block affected closure.
- **DEC-REQ-355:** Assumption-decision history shall remain retrievable.

## 17. Risk Integration

- **DEC-REQ-356:** Decisions shall identify affected risks.
- **DEC-REQ-357:** Decisions shall identify new risks.
- **DEC-REQ-358:** Decisions shall identify inherent-risk impact.
- **DEC-REQ-359:** Decisions shall identify control-effectiveness impact.
- **DEC-REQ-360:** Decisions shall identify residual-risk impact.
- **DEC-REQ-361:** Decisions shall identify risk owners.
- **DEC-REQ-362:** Decisions shall identify treatment implications.
- **DEC-REQ-363:** Decisions shall identify risk-acceptance authority where applicable.
- **DEC-REQ-364:** Decision acceptance shall not accept risk automatically.
- **DEC-REQ-365:** An AI system shall not accept risk.
- **DEC-REQ-366:** Critical residual risk shall block affected acceptance.
- **DEC-REQ-367:** Canonical-prohibition risk shall not be accepted.
- **DEC-REQ-368:** Risk exceptions shall remain explicit.
- **DEC-REQ-369:** Risk changes shall be attributable.
- **DEC-REQ-370:** Security incidents shall trigger affected decision review.
- **DEC-REQ-371:** Privacy incidents shall trigger affected decision review.
- **DEC-REQ-372:** Decision reversal shall trigger risk reassessment.
- **DEC-REQ-373:** Risk integration shall conform to `OBDIA-RISK-001`.
- **DEC-REQ-374:** Broken risk links shall block affected release.
- **DEC-REQ-375:** Risk-decision history shall remain retrievable.

## 18. Threat and Security Integration

- **DEC-REQ-376:** Security decisions shall identify affected threats.
- **DEC-REQ-377:** Security decisions shall identify affected assets.
- **DEC-REQ-378:** Security decisions shall identify trust boundaries.
- **DEC-REQ-379:** Security decisions shall identify attack paths.
- **DEC-REQ-380:** Security decisions shall identify controls.
- **DEC-REQ-381:** Security decisions shall identify failure behavior.
- **DEC-REQ-382:** Security decisions shall identify detection.
- **DEC-REQ-383:** Security decisions shall identify containment.
- **DEC-REQ-384:** Security decisions shall identify recovery.
- **DEC-REQ-385:** Security decisions shall identify residual risk.
- **DEC-REQ-386:** Security decisions shall preserve Zero Trust.
- **DEC-REQ-387:** Security decisions shall preserve least privilege.
- **DEC-REQ-388:** Security decisions shall preserve defense in depth.
- **DEC-REQ-389:** Security decisions shall preserve secure defaults.
- **DEC-REQ-390:** Security decisions shall preserve explicit authorization.
- **DEC-REQ-391:** Security decisions shall not authorize offensive operations.
- **DEC-REQ-392:** Security decisions shall not authorize malware deployment.
- **DEC-REQ-393:** Security decisions shall not authorize uncontrolled criminal-infrastructure interaction.
- **DEC-REQ-394:** Security decisions shall receive Security Reviewer assessment.
- **DEC-REQ-395:** Security integration shall conform to `OBDIA-SEC-001` and `OBDIA-TM-001`.

## 19. Privacy and Fundamental-Rights Integration

- **DEC-REQ-396:** Decisions shall identify personal-data impact.
- **DEC-REQ-397:** Decisions shall identify purpose-limitation impact.
- **DEC-REQ-398:** Decisions shall identify data-minimization impact.
- **DEC-REQ-399:** Decisions shall identify retention impact.
- **DEC-REQ-400:** Decisions shall identify deletion impact.
- **DEC-REQ-401:** Decisions shall identify access impact.
- **DEC-REQ-402:** Decisions shall identify transfer impact.
- **DEC-REQ-403:** Decisions shall identify jurisdiction impact where material.
- **DEC-REQ-404:** Decisions shall identify human-review impact.
- **DEC-REQ-405:** Decisions shall identify discrimination risk.
- **DEC-REQ-406:** Decisions shall identify surveillance risk.
- **DEC-REQ-407:** Decisions shall identify contestability and correction impact where material.
- **DEC-REQ-408:** Decisions shall not authorize unlawful surveillance.
- **DEC-REQ-409:** Decisions shall not authorize unlawful deanonymization.
- **DEC-REQ-410:** Decisions shall not infer lawful authority from technical capability.
- **DEC-REQ-411:** Decisions shall not treat public availability as unrestricted processing authority.
- **DEC-REQ-412:** Material rights-impacting decisions shall receive Privacy and Governance Reviewer assessment.
- **DEC-REQ-413:** Privacy incidents shall trigger affected decision review.
- **DEC-REQ-414:** Privacy limitations shall remain visible.
- **DEC-REQ-415:** Privacy-decision history shall remain retrievable.

## 20. Identity and Officer-Agent Binding Decisions

- **DEC-REQ-416:** Identity decisions shall distinguish human identities.
- **DEC-REQ-417:** Identity decisions shall distinguish agent identities.
- **DEC-REQ-418:** Identity decisions shall distinguish workload identities.
- **DEC-REQ-419:** Identity decisions shall distinguish connector identities.
- **DEC-REQ-420:** Binding decisions shall preserve one agent identity to one officer.
- **DEC-REQ-421:** Binding decisions shall preserve accountable institution.
- **DEC-REQ-422:** Binding decisions shall prohibit transfer of reusable human credentials.
- **DEC-REQ-423:** Binding decisions shall identify issuance authority.
- **DEC-REQ-424:** Binding decisions shall identify revocation authority.
- **DEC-REQ-425:** Binding decisions shall identify key rotation.
- **DEC-REQ-426:** Binding decisions shall prohibit officer rebinding of an existing agent identity.
- **DEC-REQ-427:** Binding decisions shall identify compromise behavior.
- **DEC-REQ-428:** Binding decisions shall identify recovery.
- **DEC-REQ-429:** Identity decisions shall identify directory or registry dependencies.
- **DEC-REQ-430:** Identity decisions shall identify privacy impact.
- **DEC-REQ-431:** Identity decisions shall identify implementation impact.
- **DEC-REQ-432:** Identity decisions shall identify validation tests.
- **DEC-REQ-433:** Critical identity decisions shall receive Security Reviewer assessment.
- **DEC-REQ-434:** Identity decisions shall conform to `OBDIA-ID-001`.
- **DEC-REQ-435:** Identity-decision history shall remain retrievable.

## 21. Authorization Decisions

- **DEC-REQ-436:** Authorization decisions shall identify subject.
- **DEC-REQ-437:** Authorization decisions shall identify resource.
- **DEC-REQ-438:** Authorization decisions shall identify action.
- **DEC-REQ-439:** Authorization decisions shall identify context.
- **DEC-REQ-440:** Authorization decisions shall identify policy version.
- **DEC-REQ-441:** Authorization decisions shall identify attribute sources.
- **DEC-REQ-442:** Authorization decisions shall identify decision outcomes.
- **DEC-REQ-443:** Authorization decisions shall preserve default deny.
- **DEC-REQ-444:** Authorization decisions shall identify obligations.
- **DEC-REQ-445:** Authorization decisions shall identify human approvals.
- **DEC-REQ-446:** Authorization decisions shall identify revocation.
- **DEC-REQ-447:** Authorization decisions shall identify cache invalidation.
- **DEC-REQ-448:** Authorization decisions shall identify failure-closed behavior.
- **DEC-REQ-449:** Authorization decisions shall preserve case and purpose isolation.
- **DEC-REQ-450:** Authorization decisions shall identify jurisdiction and time where material.
- **DEC-REQ-451:** Authorization decisions shall not infer legal authority from Permit.
- **DEC-REQ-452:** Authorization decisions shall identify negative tests.
- **DEC-REQ-453:** Critical authorization decisions shall receive Security Reviewer assessment.
- **DEC-REQ-454:** Authorization decisions shall conform to `OBDIA-AUTH-001`.
- **DEC-REQ-455:** Authorization-decision history shall remain retrievable.

## 22. Evidence Decisions

- **DEC-REQ-456:** Evidence decisions shall distinguish source and derived artifacts.
- **DEC-REQ-457:** Evidence decisions shall identify acquisition impact.
- **DEC-REQ-458:** Evidence decisions shall identify integrity impact.
- **DEC-REQ-459:** Evidence decisions shall identify provenance impact.
- **DEC-REQ-460:** Evidence decisions shall identify custody impact.
- **DEC-REQ-461:** Evidence decisions shall identify storage impact.
- **DEC-REQ-462:** Evidence decisions shall identify access impact.
- **DEC-REQ-463:** Evidence decisions shall identify retention impact.
- **DEC-REQ-464:** Evidence decisions shall identify export impact.
- **DEC-REQ-465:** Evidence decisions shall identify correction impact.
- **DEC-REQ-466:** Evidence decisions shall identify quarantine impact.
- **DEC-REQ-467:** Evidence decisions shall not equate integrity with authenticity.
- **DEC-REQ-468:** Evidence decisions shall not imply admissibility.
- **DEC-REQ-469:** Evidence decisions shall not imply lawful acquisition.
- **DEC-REQ-470:** Evidence decisions shall identify AI-generated artifact handling.
- **DEC-REQ-471:** Evidence decisions shall identify migration and rollback.
- **DEC-REQ-472:** Evidence decisions shall identify validation tests.
- **DEC-REQ-473:** Critical evidence decisions shall receive security and evidence review.
- **DEC-REQ-474:** Evidence decisions shall conform to `OBDIA-EVID-001`.
- **DEC-REQ-475:** Evidence-decision history shall remain retrievable.

## 23. Agent-Lifecycle Decisions

- **DEC-REQ-476:** Agent-lifecycle decisions shall preserve the authoritative lifecycle.
- **DEC-REQ-477:** Agent-lifecycle decisions shall preserve mandatory `Bind`.
- **DEC-REQ-478:** Agent-lifecycle decisions shall preserve authorization before activation.
- **DEC-REQ-479:** Agent-lifecycle decisions shall preserve suspension.
- **DEC-REQ-480:** Agent-lifecycle decisions shall preserve terminal revocation.
- **DEC-REQ-481:** Agent-lifecycle decisions shall preserve terminal archival.
- **DEC-REQ-482:** Agent-lifecycle decisions shall prohibit officer rebinding.
- **DEC-REQ-483:** Agent-lifecycle decisions shall identify transition authority.
- **DEC-REQ-484:** Agent-lifecycle decisions shall identify state source of truth.
- **DEC-REQ-485:** Agent-lifecycle decisions shall identify concurrency behavior.
- **DEC-REQ-486:** Agent-lifecycle decisions shall identify propagation.
- **DEC-REQ-487:** Agent-lifecycle decisions shall identify session impact.
- **DEC-REQ-488:** Agent-lifecycle decisions shall identify workload impact.
- **DEC-REQ-489:** Agent-lifecycle decisions shall identify connector impact.
- **DEC-REQ-490:** Agent-lifecycle decisions shall identify rollback limits.
- **DEC-REQ-491:** Agent-lifecycle decisions shall identify invalid-transition tests.
- **DEC-REQ-492:** Critical lifecycle decisions shall receive Security Reviewer assessment.
- **DEC-REQ-493:** Lifecycle decisions shall conform to `OBDIA-AGENT-001`.
- **DEC-REQ-494:** Lifecycle changes shall trigger dependency review.
- **DEC-REQ-495:** Lifecycle-decision history shall remain retrievable.

## 24. Connector Decisions

- **DEC-REQ-496:** Connector decisions shall identify connector identity.
- **DEC-REQ-497:** Connector decisions shall identify authentication.
- **DEC-REQ-498:** Connector decisions shall identify authorization.
- **DEC-REQ-499:** Connector decisions shall identify credential scope.
- **DEC-REQ-500:** Connector decisions shall identify endpoint scope.
- **DEC-REQ-501:** Connector decisions shall identify egress controls.
- **DEC-REQ-502:** Connector decisions shall identify request validation.
- **DEC-REQ-503:** Connector decisions shall identify response validation.
- **DEC-REQ-504:** Connector decisions shall identify provenance.
- **DEC-REQ-505:** Connector decisions shall identify prompt-injection handling.
- **DEC-REQ-506:** Connector decisions shall identify failure isolation.
- **DEC-REQ-507:** Connector decisions shall identify suspension and revocation.
- **DEC-REQ-508:** Connector decisions shall identify rate, quota and cost controls.
- **DEC-REQ-509:** Connector decisions shall identify external-provider risk.
- **DEC-REQ-510:** Connector decisions shall not use human credentials.
- **DEC-REQ-511:** Connector decisions shall not authorize uncontrolled criminal-infrastructure interaction.
- **DEC-REQ-512:** Connector decisions shall identify sandbox or testnet validation.
- **DEC-REQ-513:** Critical connector decisions shall receive Security Reviewer assessment.
- **DEC-REQ-514:** Connector decisions shall conform to `OBDIA-CONN-001`.
- **DEC-REQ-515:** Connector-decision history shall remain retrievable.

## 25. AI, Model and Prompt Decisions

- **DEC-REQ-516:** AI decisions shall identify model and provider.
- **DEC-REQ-517:** AI decisions shall identify model version where available.
- **DEC-REQ-518:** AI decisions shall identify prompt version where material.
- **DEC-REQ-519:** AI decisions shall identify retrieval configuration.
- **DEC-REQ-520:** AI decisions shall identify memory configuration.
- **DEC-REQ-521:** AI decisions shall identify tool schemas.
- **DEC-REQ-522:** AI decisions shall identify nondeterminism.
- **DEC-REQ-523:** AI decisions shall identify drift risk.
- **DEC-REQ-524:** AI decisions shall identify prompt-injection risk.
- **DEC-REQ-525:** AI decisions shall identify privacy and retention risk.
- **DEC-REQ-526:** AI decisions shall treat model output as untrusted.
- **DEC-REQ-527:** AI decisions shall not grant authorization authority.
- **DEC-REQ-528:** AI decisions shall not grant risk-acceptance authority.
- **DEC-REQ-529:** AI decisions shall not grant independent legal authority.
- **DEC-REQ-530:** AI decisions shall preserve human approval boundaries.
- **DEC-REQ-531:** AI decisions shall identify fallback behavior.
- **DEC-REQ-532:** AI decisions shall identify validation limits.
- **DEC-REQ-533:** Critical AI decisions shall receive security and privacy review.
- **DEC-REQ-534:** AI decision evidence shall identify exact configuration versions.
- **DEC-REQ-535:** AI-decision history shall remain retrievable.

## 26. Implementation Decisions

- **DEC-REQ-536:** Implementation decisions shall identify architecture dependencies.
- **DEC-REQ-537:** Implementation decisions shall identify source-code scope.
- **DEC-REQ-538:** Implementation decisions shall identify language and runtime where material.
- **DEC-REQ-539:** Implementation decisions shall identify dependencies.
- **DEC-REQ-540:** Implementation decisions shall identify configuration.
- **DEC-REQ-541:** Implementation decisions shall identify secrets handling.
- **DEC-REQ-542:** Implementation decisions shall identify interfaces.
- **DEC-REQ-543:** Implementation decisions shall identify data models.
- **DEC-REQ-544:** Implementation decisions shall identify failure behavior.
- **DEC-REQ-545:** Implementation decisions shall identify resource limits.
- **DEC-REQ-546:** Implementation decisions shall identify migration.
- **DEC-REQ-547:** Implementation decisions shall identify rollback.
- **DEC-REQ-548:** Implementation decisions shall identify tests.
- **DEC-REQ-549:** Implementation decisions shall identify release impact.
- **DEC-REQ-550:** Implementation decisions shall identify maintainability impact.
- **DEC-REQ-551:** Implementation decisions shall identify supply-chain impact.
- **DEC-REQ-552:** Implementation decisions shall identify secure defaults.
- **DEC-REQ-553:** Implementation decisions shall receive Implementation Reviewer assessment.
- **DEC-REQ-554:** Implementation decisions shall conform to `OBDIA-CODE-001`.
- **DEC-REQ-555:** Implementation-decision history shall remain retrievable.

## 27. Testing and Validation Decisions

- **DEC-REQ-556:** Testing decisions shall identify acceptance criteria.
- **DEC-REQ-557:** Testing decisions shall identify test levels.
- **DEC-REQ-558:** Testing decisions shall identify positive cases.
- **DEC-REQ-559:** Testing decisions shall identify negative cases.
- **DEC-REQ-560:** Testing decisions shall identify regression cases.
- **DEC-REQ-561:** Testing decisions shall identify fixtures.
- **DEC-REQ-562:** Testing decisions shall identify environments.
- **DEC-REQ-563:** Testing decisions shall identify tools and versions.
- **DEC-REQ-564:** Testing decisions shall identify reproducibility.
- **DEC-REQ-565:** Testing decisions shall identify evidence retention.
- **DEC-REQ-566:** Testing decisions shall identify failure disposition.
- **DEC-REQ-567:** Testing decisions shall identify flaky-test treatment.
- **DEC-REQ-568:** Testing decisions shall identify coverage limitations.
- **DEC-REQ-569:** Testing decisions shall not infer validation from passing tests alone.
- **DEC-REQ-570:** Testing decisions shall not infer security from coverage percentages alone.
- **DEC-REQ-571:** Validation decisions shall identify exact scope.
- **DEC-REQ-572:** Validation decisions shall identify exact implementation versions.
- **DEC-REQ-573:** Testing decisions shall receive Implementation Reviewer assessment.
- **DEC-REQ-574:** Testing decisions shall conform to `OBDIA-TEST-001`.
- **DEC-REQ-575:** Testing-decision history shall remain retrievable.

## 28. Repository and Workflow Decisions

- **DEC-REQ-576:** Repository decisions shall identify affected paths.
- **DEC-REQ-577:** Repository decisions shall identify ownership.
- **DEC-REQ-578:** Repository decisions shall identify branch impact.
- **DEC-REQ-579:** Repository decisions shall identify PR impact.
- **DEC-REQ-580:** Repository decisions shall identify review impact.
- **DEC-REQ-581:** Repository decisions shall identify ruleset impact.
- **DEC-REQ-582:** Repository decisions shall identify workflow impact.
- **DEC-REQ-583:** Repository decisions shall identify token and secret impact.
- **DEC-REQ-584:** Repository decisions shall identify dependency impact.
- **DEC-REQ-585:** Repository decisions shall identify artifact impact.
- **DEC-REQ-586:** Repository decisions shall identify release impact.
- **DEC-REQ-587:** Repository decisions shall distinguish desired and observed settings.
- **DEC-REQ-588:** Repository decisions shall not infer enforcement from documentation alone.
- **DEC-REQ-589:** Repository decisions shall identify rollback.
- **DEC-REQ-590:** Workflow decisions shall preserve least privilege.
- **DEC-REQ-591:** Workflow decisions shall protect untrusted pull requests.
- **DEC-REQ-592:** Critical workflow decisions shall receive Security Reviewer assessment.
- **DEC-REQ-593:** Repository decisions shall conform to `OBDIA-GH-001`.
- **DEC-REQ-594:** Repository decision evidence shall identify exact commits.
- **DEC-REQ-595:** Repository-decision history shall remain retrievable.

## 29. Publication Decisions

- **DEC-REQ-596:** Publication decisions shall identify content scope.
- **DEC-REQ-597:** Publication decisions shall identify audience.
- **DEC-REQ-598:** Publication decisions shall identify classification.
- **DEC-REQ-599:** Publication decisions shall identify exact artifacts.
- **DEC-REQ-600:** Publication decisions shall identify exact versions.
- **DEC-REQ-601:** Publication decisions shall identify integrity references.
- **DEC-REQ-602:** Publication decisions shall identify implementation status.
- **DEC-REQ-603:** Publication decisions shall identify validation status.
- **DEC-REQ-604:** Publication decisions shall identify known limitations.
- **DEC-REQ-605:** Publication decisions shall identify security review.
- **DEC-REQ-606:** Publication decisions shall identify privacy review.
- **DEC-REQ-607:** Publication decisions shall identify license and attribution.
- **DEC-REQ-608:** Publication decisions shall identify correction and withdrawal procedures.
- **DEC-REQ-609:** Publication decisions shall not expose secrets.
- **DEC-REQ-610:** Publication decisions shall not expose real case data.
- **DEC-REQ-611:** Publication decisions shall not imply institutional adoption.
- **DEC-REQ-612:** Publication decisions shall not imply production readiness.
- **DEC-REQ-613:** Publication decisions shall not imply certification or compliance.
- **DEC-REQ-614:** Publication decisions shall receive Release Reviewer assessment.
- **DEC-REQ-615:** Publication decisions shall conform to documents 04 and 15.

## 30. Compliance-Mapping Decisions

- **DEC-REQ-616:** Compliance-mapping decisions shall identify applicable reference.
- **DEC-REQ-617:** Compliance-mapping decisions shall identify reference version.
- **DEC-REQ-618:** Compliance-mapping decisions shall identify applicability.
- **DEC-REQ-619:** Compliance-mapping decisions shall identify mapping strength.
- **DEC-REQ-620:** Compliance-mapping decisions shall identify evidence.
- **DEC-REQ-621:** Compliance-mapping decisions shall identify gaps.
- **DEC-REQ-622:** Compliance-mapping decisions shall identify uncertainty.
- **DEC-REQ-623:** Compliance-mapping decisions shall identify jurisdiction where material.
- **DEC-REQ-624:** Compliance-mapping decisions shall not present mapping as legal compliance.
- **DEC-REQ-625:** Compliance-mapping decisions shall not replace legal advice.
- **DEC-REQ-626:** Compliance-mapping decisions shall not claim certification.
- **DEC-REQ-627:** Compliance-mapping decisions shall identify owner.
- **DEC-REQ-628:** Compliance-mapping decisions shall identify review date.
- **DEC-REQ-629:** Compliance-mapping decisions shall identify source-change triggers.
- **DEC-REQ-630:** Compliance-mapping decisions shall identify implementation impact.
- **DEC-REQ-631:** Compliance-mapping decisions shall identify validation impact.
- **DEC-REQ-632:** Compliance-mapping decisions shall identify public-claim limits.
- **DEC-REQ-633:** Document 30 shall govern complete compliance-mapping controls.
- **DEC-REQ-634:** Compliance decisions shall receive Privacy and Governance Reviewer assessment where applicable.
- **DEC-REQ-635:** Compliance-decision history shall remain retrievable.

## 31. Emergency Decisions

- **DEC-REQ-636:** Emergency decisions shall be exceptional.
- **DEC-REQ-637:** Emergency decisions shall identify the emergency.
- **DEC-REQ-638:** Emergency decisions shall identify authority.
- **DEC-REQ-639:** Emergency decisions shall identify scope.
- **DEC-REQ-640:** Emergency decisions shall identify time limit.
- **DEC-REQ-641:** Emergency decisions shall identify safety constraints.
- **DEC-REQ-642:** Emergency decisions shall identify security impact.
- **DEC-REQ-643:** Emergency decisions shall identify privacy impact.
- **DEC-REQ-644:** Emergency decisions shall identify evidence preservation.
- **DEC-REQ-645:** Emergency decisions shall identify rollback.
- **DEC-REQ-646:** Emergency decisions shall identify retrospective review.
- **DEC-REQ-647:** Emergency decisions shall not authorize canonical prohibitions.
- **DEC-REQ-648:** Emergency decisions shall not transfer human credentials.
- **DEC-REQ-649:** Emergency decisions shall not create permanent bypass silently.
- **DEC-REQ-650:** Emergency decisions shall expire automatically where feasible.
- **DEC-REQ-651:** Emergency decisions shall preserve audit records.
- **DEC-REQ-652:** Emergency decisions shall trigger risk review.
- **DEC-REQ-653:** Emergency decisions shall trigger affected ADR review.
- **DEC-REQ-654:** Emergency-decision corrections shall be additive.
- **DEC-REQ-655:** Emergency-decision history shall remain retrievable.

## 32. Rejection and Deferral

- **DEC-REQ-656:** Rejected decisions shall remain in the log.
- **DEC-REQ-657:** Rejected decisions shall identify authority.
- **DEC-REQ-658:** Rejected decisions shall identify rationale.
- **DEC-REQ-659:** Rejected decisions shall identify evidence.
- **DEC-REQ-660:** Rejected decisions shall identify alternatives.
- **DEC-REQ-661:** Rejected decisions shall identify risks.
- **DEC-REQ-662:** Rejected decisions shall not be represented as accepted.
- **DEC-REQ-663:** Rejected decisions shall not be deleted silently.
- **DEC-REQ-664:** Deferred decisions shall identify reason.
- **DEC-REQ-665:** Deferred decisions shall identify prerequisites.
- **DEC-REQ-666:** Deferred decisions shall identify owner.
- **DEC-REQ-667:** Deferred decisions shall identify review trigger.
- **DEC-REQ-668:** Deferred decisions shall identify interim controls.
- **DEC-REQ-669:** Deferred decisions shall identify risk impact.
- **DEC-REQ-670:** Deferred decisions shall not create implementation authority.
- **DEC-REQ-671:** Deferred decisions shall not accept risk automatically.
- **DEC-REQ-672:** Repeated deferral shall trigger governance review.
- **DEC-REQ-673:** Deferral expiry shall trigger review.
- **DEC-REQ-674:** Rejection and deferral history shall remain retrievable.
- **DEC-REQ-675:** Late changes shall not fabricate prior rationale.

## 33. Supersession

- **DEC-REQ-676:** Supersession shall identify predecessor.
- **DEC-REQ-677:** Supersession shall identify successor.
- **DEC-REQ-678:** Supersession shall identify authority.
- **DEC-REQ-679:** Supersession shall identify rationale.
- **DEC-REQ-680:** Supersession shall identify effective date.
- **DEC-REQ-681:** Supersession shall identify affected artifacts.
- **DEC-REQ-682:** Supersession shall identify implementation impact.
- **DEC-REQ-683:** Supersession shall identify migration.
- **DEC-REQ-684:** Supersession shall identify rollback limitations.
- **DEC-REQ-685:** Supersession shall identify security and privacy impact.
- **DEC-REQ-686:** Superseded entries shall remain immutable.
- **DEC-REQ-687:** Superseded entries shall remain retrievable.
- **DEC-REQ-688:** Successor acceptance shall require independent authority evidence.
- **DEC-REQ-689:** Supersession shall not transfer validation automatically.
- **DEC-REQ-690:** Supersession shall not transfer risk acceptance automatically.
- **DEC-REQ-691:** Partial supersession shall identify retained scope.
- **DEC-REQ-692:** Ambiguous supersession shall block implementation.
- **DEC-REQ-693:** Indexes shall update successor links.
- **DEC-REQ-694:** Supersession corrections shall be additive.
- **DEC-REQ-695:** Supersession history shall remain retrievable.

## 34. Reversal and Rollback

- **DEC-REQ-696:** Reversal shall be a new governed decision.
- **DEC-REQ-697:** Reversal shall identify the original decision.
- **DEC-REQ-698:** Reversal shall identify authority.
- **DEC-REQ-699:** Reversal shall identify rationale.
- **DEC-REQ-700:** Reversal shall identify effective date.
- **DEC-REQ-701:** Reversal shall identify affected artifacts.
- **DEC-REQ-702:** Reversal shall identify implementation rollback.
- **DEC-REQ-703:** Reversal shall identify data and evidence impact.
- **DEC-REQ-704:** Reversal shall identify security and privacy impact.
- **DEC-REQ-705:** Reversal shall identify dependency impact.
- **DEC-REQ-706:** Reversal shall identify validation.
- **DEC-REQ-707:** Reversal shall preserve original history.
- **DEC-REQ-708:** Reversal shall not delete the original decision.
- **DEC-REQ-709:** Rollback shall not restore revoked authority.
- **DEC-REQ-710:** Rollback shall not restore prohibited behavior.
- **DEC-REQ-711:** Rollback shall not restore known insecure configuration without explicit containment.
- **DEC-REQ-712:** Rollback shall identify exact versions.
- **DEC-REQ-713:** Rollback shall identify tests.
- **DEC-REQ-714:** Rollback failure shall remain visible.
- **DEC-REQ-715:** Reversal history shall remain retrievable.

## 35. Corrections and Amendments

- **DEC-REQ-716:** Decision corrections shall be additive.
- **DEC-REQ-717:** Corrections shall identify the original entry.
- **DEC-REQ-718:** Corrections shall identify the error.
- **DEC-REQ-719:** Corrections shall identify authority.
- **DEC-REQ-720:** Corrections shall identify date.
- **DEC-REQ-721:** Corrections shall identify affected fields.
- **DEC-REQ-722:** Corrections shall identify impact.
- **DEC-REQ-723:** Corrections shall not alter original chronology silently.
- **DEC-REQ-724:** Corrections shall not fabricate prior authority.
- **DEC-REQ-725:** Corrections shall not fabricate prior evidence.
- **DEC-REQ-726:** Corrections shall not fabricate prior alternatives.
- **DEC-REQ-727:** Corrections shall not fabricate prior review.
- **DEC-REQ-728:** Material amendments shall require renewed review.
- **DEC-REQ-729:** Incompatible amendments shall require a successor decision.
- **DEC-REQ-730:** Editorial corrections shall not change decision meaning.
- **DEC-REQ-731:** Public corrections shall follow release policy.
- **DEC-REQ-732:** Correction records shall preserve integrity references.
- **DEC-REQ-733:** Correction records shall remain retrievable.
- **DEC-REQ-734:** Silent rewriting shall be prohibited.
- **DEC-REQ-735:** Correction history shall remain auditable.

## 36. Machine-Readable Decision Log

- **DEC-REQ-736:** The decision log shall be machine-searchable.
- **DEC-REQ-737:** A machine-readable representation may supplement Markdown.
- **DEC-REQ-738:** Machine-readable entries shall identify schema version.
- **DEC-REQ-739:** Machine-readable entries shall identify decision entry ID.
- **DEC-REQ-740:** Machine-readable entries shall identify status.
- **DEC-REQ-741:** Machine-readable entries shall identify dates.
- **DEC-REQ-742:** Machine-readable entries shall identify authority.
- **DEC-REQ-743:** Machine-readable entries shall identify ADR references.
- **DEC-REQ-744:** Machine-readable entries shall identify affected artifacts.
- **DEC-REQ-745:** Machine-readable entries shall identify links.
- **DEC-REQ-746:** Machine-readable entries shall identify supersession and reversal.
- **DEC-REQ-747:** Machine-readable entries shall not contain secrets.
- **DEC-REQ-748:** Machine-readable entries shall minimize personal data.
- **DEC-REQ-749:** Machine-readable and Markdown representations shall remain synchronized.
- **DEC-REQ-750:** Synchronization conflicts shall block authoritative use.
- **DEC-REQ-751:** Generated indexes shall identify authoritative source.
- **DEC-REQ-752:** Generated indexes shall receive human review.
- **DEC-REQ-753:** Schema migration shall preserve history.
- **DEC-REQ-754:** Unknown schema versions shall fail safely.
- **DEC-REQ-755:** Machine-readable history shall remain retrievable.

## 37. Repository Placement and Indexing

- **DEC-REQ-756:** The authoritative policy shall reside at `docs/knowledge/29_DECISION_LOG_POLICY.md`.
- **DEC-REQ-757:** Decision entries shall reside in a governed repository location.
- **DEC-REQ-758:** The decision-log index shall identify each entry.
- **DEC-REQ-759:** The index shall identify title.
- **DEC-REQ-760:** The index shall identify type.
- **DEC-REQ-761:** The index shall identify status.
- **DEC-REQ-762:** The index shall identify decision date.
- **DEC-REQ-763:** The index shall identify authority.
- **DEC-REQ-764:** The index shall identify ADR reference where applicable.
- **DEC-REQ-765:** The index shall identify supersession where applicable.
- **DEC-REQ-766:** The index shall identify authoritative path.
- **DEC-REQ-767:** The index shall not infer status from path.
- **DEC-REQ-768:** The index shall not list absent entries as complete.
- **DEC-REQ-769:** Decision-entry filenames shall follow `OBDIA-NAME-001`.
- **DEC-REQ-770:** Repository moves shall preserve history.
- **DEC-REQ-771:** Repository moves shall update links.
- **DEC-REQ-772:** Temporary files shall not become authoritative entries.
- **DEC-REQ-773:** Generated indexes shall not replace source records.
- **DEC-REQ-774:** Document 24 shall govern repository controls.
- **DEC-REQ-775:** Repository decision records shall remain auditable.

## 38. Review Gates

- **DEC-REQ-776:** Material decisions shall receive human review.
- **DEC-REQ-777:** Review shall identify exact decision entry.
- **DEC-REQ-778:** Review shall identify exact ADR where applicable.
- **DEC-REQ-779:** Review shall identify exact repository commit.
- **DEC-REQ-780:** Review shall verify authority.
- **DEC-REQ-781:** Review shall verify scope.
- **DEC-REQ-782:** Review shall verify rationale.
- **DEC-REQ-783:** Review shall verify alternatives.
- **DEC-REQ-784:** Review shall verify evidence.
- **DEC-REQ-785:** Review shall verify assumptions.
- **DEC-REQ-786:** Review shall verify risks.
- **DEC-REQ-787:** Review shall verify security impact.
- **DEC-REQ-788:** Review shall verify privacy impact.
- **DEC-REQ-789:** Review shall verify implementation impact.
- **DEC-REQ-790:** Review shall verify testing and validation impact.
- **DEC-REQ-791:** Review shall verify release impact.
- **DEC-REQ-792:** Review shall verify migration and rollback.
- **DEC-REQ-793:** Security-sensitive decisions shall receive Security Reviewer assessment.
- **DEC-REQ-794:** Privacy-sensitive decisions shall receive Privacy and Governance Reviewer assessment.
- **DEC-REQ-795:** Implementation decisions shall receive Implementation Reviewer assessment.
- **DEC-REQ-796:** Publication decisions shall receive Release Reviewer assessment.
- **DEC-REQ-797:** Role concentration shall be disclosed.
- **DEC-REQ-798:** Material post-review changes shall invalidate affected approval.
- **DEC-REQ-799:** Automated review shall remain advisory.
- **DEC-REQ-800:** Review history shall remain retrievable.

## 39. Approval and Acceptance

- **DEC-REQ-801:** Decision acceptance shall identify authority.
- **DEC-REQ-802:** Decision acceptance shall identify exact entry version.
- **DEC-REQ-803:** Decision acceptance shall identify exact ADR version where applicable.
- **DEC-REQ-804:** Decision acceptance shall identify integrity reference where material.
- **DEC-REQ-805:** Decision acceptance shall identify date.
- **DEC-REQ-806:** Decision acceptance shall identify scope.
- **DEC-REQ-807:** Decision acceptance shall identify conditions.
- **DEC-REQ-808:** Decision acceptance shall identify expiry where conditional.
- **DEC-REQ-809:** Decision acceptance shall identify required follow-up.
- **DEC-REQ-810:** Decision acceptance shall not imply implementation.
- **DEC-REQ-811:** Decision acceptance shall not imply validation.
- **DEC-REQ-812:** Decision acceptance shall not imply publication.
- **DEC-REQ-813:** Decision acceptance shall not imply compliance.
- **DEC-REQ-814:** Decision acceptance shall not imply risk acceptance.
- **DEC-REQ-815:** An AI system shall not accept decisions.
- **DEC-REQ-816:** Unauthorized acceptance shall be invalid.
- **DEC-REQ-817:** Conditional acceptance shall identify unresolved conditions.
- **DEC-REQ-818:** Expired conditional acceptance shall trigger review.
- **DEC-REQ-819:** Acceptance corrections shall preserve history.
- **DEC-REQ-820:** Acceptance evidence shall remain retrievable.

## 40. Testing and Validation of the Decision Log

- **DEC-REQ-821:** Decision-log tests shall verify identifier uniqueness.
- **DEC-REQ-822:** Decision-log tests shall verify chronological fields.
- **DEC-REQ-823:** Decision-log tests shall verify mandatory fields.
- **DEC-REQ-824:** Decision-log tests shall verify valid statuses.
- **DEC-REQ-825:** Decision-log tests shall verify valid types.
- **DEC-REQ-826:** Decision-log tests shall verify ADR references for major decisions.
- **DEC-REQ-827:** Decision-log tests shall verify implementation-impact fields.
- **DEC-REQ-828:** Decision-log tests shall verify affected-artifact links.
- **DEC-REQ-829:** Decision-log tests shall verify assumption links.
- **DEC-REQ-830:** Decision-log tests shall verify risk links.
- **DEC-REQ-831:** Decision-log tests shall verify supersession links.
- **DEC-REQ-832:** Decision-log tests shall verify reversal links.
- **DEC-REQ-833:** Decision-log tests shall verify no silent identifier reuse.
- **DEC-REQ-834:** Decision-log tests shall verify no knowledge-document number above 30.
- **DEC-REQ-835:** Decision-log tests shall verify machine-readable synchronization where used.
- **DEC-REQ-836:** Failed tests shall remain visible.
- **DEC-REQ-837:** Passing structural tests shall not establish decision correctness.
- **DEC-REQ-838:** Human review shall assess authority and rationale.
- **DEC-REQ-839:** Testing shall conform to `OBDIA-TEST-001`.
- **DEC-REQ-840:** Decision-log validation evidence shall remain retrievable.

## 41. Retention and Archival

- **DEC-REQ-841:** Decision entries shall have retention rules.
- **DEC-REQ-842:** Accepted decisions shall remain retrievable.
- **DEC-REQ-843:** Rejected decisions shall remain retrievable.
- **DEC-REQ-844:** Deferred decisions shall remain retrievable.
- **DEC-REQ-845:** Superseded decisions shall remain retrievable.
- **DEC-REQ-846:** Reversed decisions shall remain retrievable.
- **DEC-REQ-847:** Archived decisions shall remain read-only.
- **DEC-REQ-848:** Archive records shall preserve identifiers.
- **DEC-REQ-849:** Archive records shall preserve versions.
- **DEC-REQ-850:** Archive records shall preserve chronology.
- **DEC-REQ-851:** Archive records shall preserve authority evidence.
- **DEC-REQ-852:** Archive records shall preserve ADR links.
- **DEC-REQ-853:** Archive records shall preserve evidence links.
- **DEC-REQ-854:** Archive migration shall preserve integrity.
- **DEC-REQ-855:** Archive restoration shall not reactivate decisions silently.
- **DEC-REQ-856:** Archived decisions shall not be deleted solely because they are obsolete.
- **DEC-REQ-857:** Retention shall preserve incident and release relevance.
- **DEC-REQ-858:** Disposition shall preserve required audit evidence.
- **DEC-REQ-859:** Archive corruption shall trigger incident handling.
- **DEC-REQ-860:** Retention history shall remain retrievable.

## 42. Publication and Transparency

- **DEC-REQ-861:** Public decision logs shall use approved content.
- **DEC-REQ-862:** Public decision logs shall not expose secrets.
- **DEC-REQ-863:** Public decision logs shall not expose real case data.
- **DEC-REQ-864:** Public decision logs shall not expose sensitive administrative details unnecessarily.
- **DEC-REQ-865:** Public decision logs shall distinguish accepted and proposed decisions.
- **DEC-REQ-866:** Public decision logs shall distinguish implemented and unimplemented decisions.
- **DEC-REQ-867:** Public decision logs shall distinguish validated and unvalidated decisions.
- **DEC-REQ-868:** Public decision logs shall identify limitations.
- **DEC-REQ-869:** Public decision logs shall identify role concentration where material.
- **DEC-REQ-870:** Public decision logs shall not claim independent assurance without evidence.
- **DEC-REQ-871:** Public decision logs shall not claim institutional adoption without evidence.
- **DEC-REQ-872:** Public decision logs shall not claim production readiness without evidence.
- **DEC-REQ-873:** Public decision logs shall not claim certification or compliance without evidence.
- **DEC-REQ-874:** Public decision logs shall preserve chronological integrity.
- **DEC-REQ-875:** Public decision corrections shall preserve history.
- **DEC-REQ-876:** Public decision withdrawal shall preserve records.
- **DEC-REQ-877:** Publication decisions shall receive Release Reviewer assessment.
- **DEC-REQ-878:** Public artifacts shall conform to documents 04 and 15.
- **DEC-REQ-879:** Transparency shall not override security or privacy classification.
- **DEC-REQ-880:** Publication history shall remain retrievable.

## 43. Traceability

- **DEC-REQ-881:** Every decision entry shall trace to its origin.
- **DEC-REQ-882:** Every decision entry shall trace to authority.
- **DEC-REQ-883:** Every major decision shall trace to an ADR.
- **DEC-REQ-884:** Every decision shall trace to affected artifacts.
- **DEC-REQ-885:** Every decision shall trace to affected requirements.
- **DEC-REQ-886:** Every decision shall trace to research where applicable.
- **DEC-REQ-887:** Every decision shall trace to assumptions where applicable.
- **DEC-REQ-888:** Every decision shall trace to risks where applicable.
- **DEC-REQ-889:** Every decision shall trace to threats where applicable.
- **DEC-REQ-890:** Every accepted decision shall trace to approval evidence.
- **DEC-REQ-891:** Every implemented decision shall trace to implementation evidence.
- **DEC-REQ-892:** Every validated decision shall trace to validation evidence.
- **DEC-REQ-893:** Every published decision shall trace to release evidence.
- **DEC-REQ-894:** Every superseded decision shall trace to successor.
- **DEC-REQ-895:** Every reversed decision shall trace to reversal.
- **DEC-REQ-896:** Every correction shall trace to the original entry.
- **DEC-REQ-897:** Every archive record shall trace to retained entries.
- **DEC-REQ-898:** Traceability shall be bidirectional.
- **DEC-REQ-899:** Broken traceability affecting authority, security or release shall be blocking.
- **DEC-REQ-900:** Traceability records shall not contain secrets.
- **DEC-REQ-901:** Traceability records shall minimize personal and case data.
- **DEC-REQ-902:** Traceability shall identify exact versions.
- **DEC-REQ-903:** Traceability shall identify exact commits where material.
- **DEC-REQ-904:** Baseline freeze shall validate decision traceability across documents 01–30.
- **DEC-REQ-905:** Traceability history shall remain retrievable.

## 44. Domain-Specific Decision Controls

These controls ensure that material decisions in every OBDIA domain carry explicit authority, evidence, impact, implementation and lifecycle boundaries.

- **DEC-REQ-906:** A material decision concerning officer-agent binding shall identify the decision statement concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-907:** A material decision concerning officer-agent binding shall identify accountable authority concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-908:** A material decision concerning officer-agent binding shall identify scope and exclusions concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-909:** A material decision concerning officer-agent binding shall identify corresponding ADR requirements concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-910:** A material decision concerning officer-agent binding shall identify rationale and alternatives concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-911:** A material decision concerning officer-agent binding shall identify supporting research and evidence concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-912:** A material decision concerning officer-agent binding shall identify assumptions and confidence concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-913:** A material decision concerning officer-agent binding shall identify threats and risks concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-914:** A material decision concerning officer-agent binding shall identify security impact concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-915:** A material decision concerning officer-agent binding shall identify privacy and rights impact concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-916:** A material decision concerning officer-agent binding shall identify evidence impact concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-917:** A material decision concerning officer-agent binding shall identify implementation impact concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-918:** A material decision concerning officer-agent binding shall identify testing and validation impact concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-919:** A material decision concerning officer-agent binding shall identify migration and rollback concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-920:** A material decision concerning officer-agent binding shall identify release and publication impact concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-921:** A material decision concerning officer-agent binding shall identify status, chronology and effective date concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-922:** A material decision concerning officer-agent binding shall identify supersession or reversal conditions concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-923:** A material decision concerning officer-agent binding shall identify retention and traceability concerning institutional issuance, one officer, one agent identity and no credential transfer.
- **DEC-REQ-924:** A material decision concerning machine identity shall identify the decision statement concerning human, agent, workload and connector identity separation.
- **DEC-REQ-925:** A material decision concerning machine identity shall identify accountable authority concerning human, agent, workload and connector identity separation.
- **DEC-REQ-926:** A material decision concerning machine identity shall identify scope and exclusions concerning human, agent, workload and connector identity separation.
- **DEC-REQ-927:** A material decision concerning machine identity shall identify corresponding ADR requirements concerning human, agent, workload and connector identity separation.
- **DEC-REQ-928:** A material decision concerning machine identity shall identify rationale and alternatives concerning human, agent, workload and connector identity separation.
- **DEC-REQ-929:** A material decision concerning machine identity shall identify supporting research and evidence concerning human, agent, workload and connector identity separation.
- **DEC-REQ-930:** A material decision concerning machine identity shall identify assumptions and confidence concerning human, agent, workload and connector identity separation.
- **DEC-REQ-931:** A material decision concerning machine identity shall identify threats and risks concerning human, agent, workload and connector identity separation.
- **DEC-REQ-932:** A material decision concerning machine identity shall identify security impact concerning human, agent, workload and connector identity separation.
- **DEC-REQ-933:** A material decision concerning machine identity shall identify privacy and rights impact concerning human, agent, workload and connector identity separation.
- **DEC-REQ-934:** A material decision concerning machine identity shall identify evidence impact concerning human, agent, workload and connector identity separation.
- **DEC-REQ-935:** A material decision concerning machine identity shall identify implementation impact concerning human, agent, workload and connector identity separation.
- **DEC-REQ-936:** A material decision concerning machine identity shall identify testing and validation impact concerning human, agent, workload and connector identity separation.
- **DEC-REQ-937:** A material decision concerning machine identity shall identify migration and rollback concerning human, agent, workload and connector identity separation.
- **DEC-REQ-938:** A material decision concerning machine identity shall identify release and publication impact concerning human, agent, workload and connector identity separation.
- **DEC-REQ-939:** A material decision concerning machine identity shall identify status, chronology and effective date concerning human, agent, workload and connector identity separation.
- **DEC-REQ-940:** A material decision concerning machine identity shall identify supersession or reversal conditions concerning human, agent, workload and connector identity separation.
- **DEC-REQ-941:** A material decision concerning machine identity shall identify retention and traceability concerning human, agent, workload and connector identity separation.
- **DEC-REQ-942:** A material decision concerning authorization shall identify the decision statement concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-943:** A material decision concerning authorization shall identify accountable authority concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-944:** A material decision concerning authorization shall identify scope and exclusions concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-945:** A material decision concerning authorization shall identify corresponding ADR requirements concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-946:** A material decision concerning authorization shall identify rationale and alternatives concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-947:** A material decision concerning authorization shall identify supporting research and evidence concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-948:** A material decision concerning authorization shall identify assumptions and confidence concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-949:** A material decision concerning authorization shall identify threats and risks concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-950:** A material decision concerning authorization shall identify security impact concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-951:** A material decision concerning authorization shall identify privacy and rights impact concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-952:** A material decision concerning authorization shall identify evidence impact concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-953:** A material decision concerning authorization shall identify implementation impact concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-954:** A material decision concerning authorization shall identify testing and validation impact concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-955:** A material decision concerning authorization shall identify migration and rollback concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-956:** A material decision concerning authorization shall identify release and publication impact concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-957:** A material decision concerning authorization shall identify status, chronology and effective date concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-958:** A material decision concerning authorization shall identify supersession or reversal conditions concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-959:** A material decision concerning authorization shall identify retention and traceability concerning subject, resource, action, context, default deny and revocation.
- **DEC-REQ-960:** A material decision concerning agent lifecycle shall identify the decision statement concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-961:** A material decision concerning agent lifecycle shall identify accountable authority concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-962:** A material decision concerning agent lifecycle shall identify scope and exclusions concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-963:** A material decision concerning agent lifecycle shall identify corresponding ADR requirements concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-964:** A material decision concerning agent lifecycle shall identify rationale and alternatives concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-965:** A material decision concerning agent lifecycle shall identify supporting research and evidence concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-966:** A material decision concerning agent lifecycle shall identify assumptions and confidence concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-967:** A material decision concerning agent lifecycle shall identify threats and risks concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-968:** A material decision concerning agent lifecycle shall identify security impact concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-969:** A material decision concerning agent lifecycle shall identify privacy and rights impact concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-970:** A material decision concerning agent lifecycle shall identify evidence impact concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-971:** A material decision concerning agent lifecycle shall identify implementation impact concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-972:** A material decision concerning agent lifecycle shall identify testing and validation impact concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-973:** A material decision concerning agent lifecycle shall identify migration and rollback concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-974:** A material decision concerning agent lifecycle shall identify release and publication impact concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-975:** A material decision concerning agent lifecycle shall identify status, chronology and effective date concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-976:** A material decision concerning agent lifecycle shall identify supersession or reversal conditions concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-977:** A material decision concerning agent lifecycle shall identify retention and traceability concerning provisioning, binding, authorization, activation, suspension, revocation and archival.
- **DEC-REQ-978:** A material decision concerning evidence integrity shall identify the decision statement concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-979:** A material decision concerning evidence integrity shall identify accountable authority concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-980:** A material decision concerning evidence integrity shall identify scope and exclusions concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-981:** A material decision concerning evidence integrity shall identify corresponding ADR requirements concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-982:** A material decision concerning evidence integrity shall identify rationale and alternatives concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-983:** A material decision concerning evidence integrity shall identify supporting research and evidence concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-984:** A material decision concerning evidence integrity shall identify assumptions and confidence concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-985:** A material decision concerning evidence integrity shall identify threats and risks concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-986:** A material decision concerning evidence integrity shall identify security impact concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-987:** A material decision concerning evidence integrity shall identify privacy and rights impact concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-988:** A material decision concerning evidence integrity shall identify evidence impact concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-989:** A material decision concerning evidence integrity shall identify implementation impact concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-990:** A material decision concerning evidence integrity shall identify testing and validation impact concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-991:** A material decision concerning evidence integrity shall identify migration and rollback concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-992:** A material decision concerning evidence integrity shall identify release and publication impact concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-993:** A material decision concerning evidence integrity shall identify status, chronology and effective date concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-994:** A material decision concerning evidence integrity shall identify supersession or reversal conditions concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-995:** A material decision concerning evidence integrity shall identify retention and traceability concerning original evidence, provenance, custody, transformations and exports.
- **DEC-REQ-996:** A material decision concerning connector security shall identify the decision statement concerning authentication, egress, validation, provenance, isolation and revocation.
- **DEC-REQ-997:** A material decision concerning connector security shall identify accountable authority concerning authentication, egress, validation, provenance, isolation and revocation.
- **DEC-REQ-998:** A material decision concerning connector security shall identify scope and exclusions concerning authentication, egress, validation, provenance, isolation and revocation.
- **DEC-REQ-999:** A material decision concerning connector security shall identify corresponding ADR requirements concerning authentication, egress, validation, provenance, isolation and revocation.

## 45. Minimum Validation Checklist

Before a material decision may be marked `Accepted`, confirm:

- [ ] immutable `DEC-ENTRY-NNN` identifier exists;
- [ ] decision type and materiality are accurate;
- [ ] decision statement and scope are precise;
- [ ] accountable authority is identified;
- [ ] decision, recorded and effective dates are distinguishable;
- [ ] rationale and alternatives are complete;
- [ ] supporting sources and research are linked;
- [ ] assumptions and confidence are explicit;
- [ ] threats, risks and residual risk are linked;
- [ ] corresponding ADR exists for every major decision;
- [ ] ADR status and decision status are consistent;
- [ ] affected documents and requirements are identified;
- [ ] security, privacy, evidence and rights impacts are assessed;
- [ ] implementation impact is explicit;
- [ ] testing and validation impact is explicit;
- [ ] migration and rollback are defined where material;
- [ ] release and publication impact is explicit;
- [ ] role concentration and review limitations are disclosed;
- [ ] prohibited actions are not authorized;
- [ ] decision acceptance does not imply implementation, validation or publication;
- [ ] approval evidence identifies the exact entry and commit;
- [ ] supersession and reversal paths are defined;
- [ ] machine-readable and Markdown records agree where used;
- [ ] immutable history and bidirectional traceability are preserved;
- [ ] forward dependency on document 30 is recorded.

## 46. Limitations

- This policy does not accept any decision.
- It does not replace ADRs, normative policies, risk records, implementation plans, test evidence or release approvals.
- Chronological recording does not prove that a decision was correct.
- Complete rationale does not prove that all alternatives were known.
- Repository timestamps and signatures do not by themselves establish authority.
- Internal review is not independent certification.
- An accepted decision can later become stale, superseded, reversed or invalidated by new evidence.
- Decision logs cannot encode legal conclusions autonomously.
- Document 30 remains a forward dependency for compliance-mapping decisions.

## 47. Change Control

Every material change shall identify rationale, affected decision types, statuses, entry schemas, chronology, ADR integration, security and privacy impact, implementation and release impact, migration, validation, rollback, authority and version effect.

- Editorial corrections shall use a patch version when meaning is unchanged.
- Backward-compatible substantive additions shall use a minor version.
- Incompatible decision-governance changes shall use a major version.
- Material decision-governance architecture changes shall require an ADR where applicable.
- Changes shall require Project Founder approval.
- Changes shall receive Documentation Authority assessment.
- Security, privacy, implementation and release impacts shall receive specialist review where applicable.
- Changes shall identify affected records, schemas, indexes and automation.
- Changes shall include migration and compatibility analysis.
- Changes shall include validation and rollback analysis.
- Changes shall not retroactively fabricate decision authority or evidence.
- Historical decision records shall not be silently rewritten.
- Decision-entry and requirement identifiers shall not be reused.
- Migration shall preserve decision-to-ADR and decision-to-implementation traceability.
- Forward-dependency reconciliation shall occur before this policy becomes Approved.

## 48. Consolidation Record

Version 1.0.0 consolidates the two existing Decision Log Policy variants without expanding project scope. It:

- retains immutable document identifier `OBDIA-DEC-001`;
- normalizes the authoritative filename to `29_DECISION_LOG_POLICY.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves the requirement for a chronological decision log;
- preserves architectural decisions;
- preserves governance decisions;
- preserves security decisions;
- preserves publication decisions;
- preserves mandatory ADR references when applicable;
- preserves the requirement that every major decision reference its corresponding ADR;
- preserves mandatory implementation-impact recording for major decisions;
- adds immutable `DEC-ENTRY-NNN` identifiers;
- adds controlled decision types, materiality and statuses;
- separates decision status, implementation state, validation state and publication state;
- adds chronology, authority, rationale, alternatives, research, assumptions, risks, threats, security, privacy, evidence, implementation, testing, release, supersession, reversal, correction, retention and traceability controls;
- adds machine-readable and repository-index requirements;
- adds 999 stable decision-governance requirements;
- identifies document 30 as a forward dependency;
- treats Enterprise and non-Enterprise files as legacy source variants of the same immutable document;
- creates no accepted decision, ADR approval, implementation, validation, risk acceptance, publication or operational authority.

## 49. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-06 | Draft | Documentation Authority; approval reserved to Project Founder | Constitutional consolidation of both legacy Decision Log Policy variants: preserved chronological architectural, governance, security and publication logging, ADR references and implementation impact; added complete authority, evidence, impact, lifecycle, correction and traceability governance. |
