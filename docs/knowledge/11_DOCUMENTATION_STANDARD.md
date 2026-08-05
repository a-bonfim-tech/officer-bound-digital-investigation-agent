# DOCUMENTATION STANDARD

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-DOC-001 |
| **Title** | Documentation Standard |
| **Version** | 1.1.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define mandatory structure, metadata, writing, naming, traceability, review, security, privacy, accessibility and lifecycle standards for governed OBDIA repository documentation. |
| **Scope** | Knowledge documents 01–30, the Master Documentation Constitution, ADRs, architecture and governance records, specifications, diagrams, implementation documentation, tests, evidence, research records, release records and supporting navigation files. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `02_PROJECT_GLOSSARY.md`; `03_RESEARCH_AND_SOURCE_POLICY.md`; `04_PORTFOLIO_AND_PUBLICATION_POLICY.md`; `09_GOVERNANCE_MODEL.md`; `10_ADR_POLICY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GLOSS-001`; `OBDIA-RES-001`; `OBDIA-PUB-001`; `OBDIA-GOV-001`; `OBDIA-ADR-001` |
| **Cross-References** | Documents 12–30; accepted ADRs; repository documentation profiles; validation, release and publication evidence |
| **Assumptions** | English is the authoritative language for normative repository documentation unless an approved artifact establishes a controlled translation. |
| **Constraints** | Documentation shall not grant authority, conceal status, fabricate evidence, silently define architecture, weaken canonical boundaries or create a normative document above number 30. |
| **Security Considerations** | Documentation can expose secrets, personal data, unsafe procedures, trust assumptions, attack paths and misleading implementation claims; content and publication therefore require risk-based review. |
| **Validation Criteria** | Governed documents use valid identifiers and filenames, complete applicable metadata and sections, accurate lifecycle and implementation status, valid cross-references, traceable requirements, retained review evidence and no prohibited or sensitive content. |
| **Implementation Relationship** | Documentation governs and explains implementation but does not by itself prove implementation, testing, validation, compliance, deployment or operational authorization. |

---

## 1. Purpose

This standard defines how repository documentation is structured, written, reviewed, versioned, linked, validated and maintained for the Officer-Bound Digital Investigation Agent project.

It preserves the original Documentation Standard requirements that every governed document include:

- Document ID;
- Version;
- Status;
- Purpose;
- Scope;
- Dependencies;
- Cross-references;
- Assumptions;
- Security considerations;
- Validation criteria;
- Revision history.

It also preserves the original writing principles:

- precision over marketing;
- facts separated from hypotheses;
- consistent terminology;
- GitHub-ready Markdown.

This standard is subordinate to the Canonical Project Definition and Master Documentation Constitution. It does not replace the specialized controls in the Research and Source Policy, Portfolio and Publication Policy, Governance Model or ADR Policy.

## 2. Fundamental Documentation Rules

- **DOC-REQ-001:** Governed documentation shall be accurate, reviewable, maintainable and attributable.
- **DOC-REQ-002:** Documentation shall preserve human and institutional accountability.
- **DOC-REQ-003:** Documentation shall not represent an AI agent as possessing independent legal authority.
- **DOC-REQ-004:** Documentation shall not authorize prohibited conduct.
- **DOC-REQ-005:** Documentation shall not silently expand project scope or the normative Knowledge Pack.
- **DOC-REQ-006:** Documentation shall distinguish authority, requirements, decisions, proposals, implementation, tests and evidence.
- **DOC-REQ-007:** Repository presence or merge shall not be represented as normative approval.
- **DOC-REQ-008:** Documentation shall not silently define architecture through examples, diagrams, code excerpts or implementation notes.
- **DOC-REQ-009:** Material claims shall follow `03_RESEARCH_AND_SOURCE_POLICY.md`.
- **DOC-REQ-010:** Public claims shall follow `04_PORTFOLIO_AND_PUBLICATION_POLICY.md`.
- **DOC-REQ-011:** Material architectural decisions shall follow `10_ADR_POLICY.md`.
- **DOC-REQ-012:** Unresolved ambiguity affecting authority, security, privacy, evidence or implementation shall be visible and escalated.

## 3. Documentation Classes

### 3.1 Canonical and Constitutional Authorities

Includes:

- `01_CANONICAL_PROJECT_DEFINITION.md`;
- `MASTER_DOCUMENTATION_CONSTITUTION.md`.

These artifacts have the authority and amendment rules assigned by the Constitution.

### 3.2 Normative Knowledge Documents

Includes the approved numbered Knowledge Pack from 02 through 30.

These documents define binding project requirements within the constitutional hierarchy.

### 3.3 Architecture Decision Records

ADRs record material decisions under `OBDIA-ADR-001`.

### 3.4 Architecture, Governance and Technical Specifications

These artifacts define lower-level structures, interfaces, policies, data models, protocols, diagrams and implementation constraints.

### 3.5 Implementation and Operational Documentation

Includes component READMEs, configuration guidance, local laboratory instructions, runbooks and maintenance guidance.

### 3.6 Research Records

Research records contain questions, methods, sources, findings, uncertainty and implications under `OBDIA-RES-001`.

### 3.7 Test, Validation and Evidence Records

These artifacts record methods, environments, exact versions, outcomes, limitations and approval status.

### 3.8 Release and Publication Records

These artifacts identify reviewed commits, release content, disclosure checks, status claims and approval evidence.

### 3.9 Navigation and Support Files

Includes repository README files, indexes and contribution navigation that do not create independent normative authority.

- **DOC-REQ-013:** Every governed artifact shall declare or unambiguously inherit its documentation class.
- **DOC-REQ-014:** A lower documentation class shall not contradict a higher authority.
- **DOC-REQ-015:** Navigation files shall not create requirements or decisions by implication.
- **DOC-REQ-016:** Informative examples shall be labeled and shall not override normative text.
- **DOC-REQ-017:** Test or evidence records shall not be rewritten as requirements.
- **DOC-REQ-018:** Research records shall not be represented as approved architecture unless incorporated through formal control.
- **DOC-REQ-019:** Implementation documentation shall state the applicable implementation status.
- **DOC-REQ-020:** A document serving multiple classes shall identify which sections are normative, informative, evidentiary or experimental.

## 4. Mandatory Metadata

Every governed document shall include, where applicable:

- Document ID;
- Title;
- Version;
- Status;
- Classification;
- Authority;
- Owner;
- Approval Authority;
- Effective Date;
- Purpose;
- Scope;
- Dependencies;
- Normative References;
- Cross-References;
- Assumptions;
- Constraints;
- Security Considerations;
- Validation Criteria;
- Implementation Relationship;
- Revision History.

Navigation-only support files may use a reduced metadata profile only when they contain no independent requirements, decisions, evidence or status claims.

- **DOC-REQ-021:** Document identifiers shall be unique and immutable.
- **DOC-REQ-022:** Metadata shall appear near the beginning of the document.
- **DOC-REQ-023:** Metadata values shall be explicit rather than inferred from branch, directory or filename.
- **DOC-REQ-024:** Missing applicable metadata is a blocking documentation defect.
- **DOC-REQ-025:** A non-applicable field shall state `Not applicable` and provide rationale when material.
- **DOC-REQ-026:** Authority and owner shall identify governed human roles rather than AI systems.
- **DOC-REQ-027:** Approval Authority shall match the applicable governance rule.
- **DOC-REQ-028:** Effective Date shall not imply approval while status is Draft or Under Review.
- **DOC-REQ-029:** Dependencies shall identify artifacts required to interpret or implement the document.
- **DOC-REQ-030:** Normative References shall identify controlling authorities by immutable identifier where available.
- **DOC-REQ-031:** Cross-References shall identify related but not necessarily controlling artifacts.
- **DOC-REQ-032:** Metadata shall not contain secrets, personal credentials or unnecessary personal information.

## 5. Mandatory Document Sections

A normative Knowledge document shall include:

1. Purpose;
2. Scope or applicability;
3. authority and governing relationship;
4. definitions or terminology references where needed;
5. normative requirements;
6. roles and responsibilities where applicable;
7. security considerations;
8. privacy, ethics and legal-boundary considerations where applicable;
9. implementation relationship;
10. traceability;
11. validation criteria or checklist;
12. limitations;
13. change control;
14. consolidation or migration record where applicable;
15. revision history.

Specialized document classes may require additional sections defined by their controlling policy.

- **DOC-REQ-033:** Every required section shall contain substantive content or an explicit non-applicability rationale.
- **DOC-REQ-034:** Empty headings, `TODO`, `TBD` and deceptive placeholders are prohibited in approved artifacts.
- **DOC-REQ-035:** Draft placeholders shall identify owner, resolution condition and target review point.
- **DOC-REQ-036:** Purpose shall explain why the artifact exists.
- **DOC-REQ-037:** Scope shall define inclusions, exclusions and applicability boundaries.
- **DOC-REQ-038:** Assumptions shall not be presented as verified facts.
- **DOC-REQ-039:** Constraints shall include applicable canonical and constitutional boundaries.
- **DOC-REQ-040:** Limitations shall identify what the artifact does not establish.
- **DOC-REQ-041:** Validation criteria shall be observable or testable where feasible.
- **DOC-REQ-042:** Revision history shall record every material version and status change.

## 6. Normative Language

The preferred normative terms are:

- **shall** — mandatory requirement;
- **shall not** — prohibition;
- **should** — recommended practice requiring rationale when not followed;
- **should not** — discouraged practice requiring rationale when followed;
- **may** — permitted option;
- **can** — statement of capability or possibility, not permission.

- **DOC-REQ-043:** Normative language shall be used consistently.
- **DOC-REQ-044:** `Must` may appear in quotations or external requirements but project-authored requirements should use `shall`.
- **DOC-REQ-045:** Informative text shall not use ambiguous wording that appears mandatory.
- **DOC-REQ-046:** Requirements shall identify the responsible subject and expected behavior.
- **DOC-REQ-047:** Requirements shall avoid undefined qualifiers such as `appropriate`, `secure`, `timely` or `sufficient` without criteria.
- **DOC-REQ-048:** Prohibitions shall be explicit.
- **DOC-REQ-049:** Exceptions shall not be embedded informally inside examples or notes.
- **DOC-REQ-050:** Requirement identifiers shall remain stable after allocation.
- **DOC-REQ-051:** Requirement identifiers shall not be reused after removal or supersession.
- **DOC-REQ-052:** A substantive requirement change shall be reflected in version and revision history.

## 7. Writing Principles

Repository documentation shall use:

- precision over marketing;
- direct and professional language;
- consistent terminology;
- facts separated from assumptions and proposals;
- short, meaningful headings;
- examples that preserve security and privacy boundaries;
- sufficient context for independent review;
- explicit limitations and uncertainty.

- **DOC-REQ-053:** Unsupported superlatives and marketing claims are prohibited.
- **DOC-REQ-054:** Acronyms shall be expanded at first material use unless defined in the Project Glossary.
- **DOC-REQ-055:** Deep Web and Dark Web shall not be conflated.
- **DOC-REQ-056:** Authentication, authorization, legal authority and trust shall remain distinct.
- **DOC-REQ-057:** Existing technology, proposed architecture, implementation, simulation and experimental result shall remain distinct.
- **DOC-REQ-058:** Active voice should be used where it improves responsibility and clarity.
- **DOC-REQ-059:** Vague passive constructions shall not conceal the responsible actor or decision authority.
- **DOC-REQ-060:** Material uncertainty shall be stated near the affected claim.
- **DOC-REQ-061:** Documentation shall not imply operational deployment, certification or endorsement without evidence.
- **DOC-REQ-062:** Legal language shall identify jurisdiction and review limitations where material.

## 8. GitHub-Ready Markdown

Documents shall use portable GitHub-compatible Markdown.

- **DOC-REQ-063:** A document shall contain one top-level `#` heading.
- **DOC-REQ-064:** Heading levels shall not skip levels without a documented structural reason.
- **DOC-REQ-065:** Headings shall be descriptive and stable enough for reviewable links.
- **DOC-REQ-066:** Tables shall have header rows and remain readable in source form where feasible.
- **DOC-REQ-067:** Code fences shall identify a language or format when known.
- **DOC-REQ-068:** Commands shall be distinguishable from expected output.
- **DOC-REQ-069:** Expected output shall not be presented as a command to execute.
- **DOC-REQ-070:** Shell examples shall avoid destructive or interactive failure behavior unless explicitly explained.
- **DOC-REQ-071:** Commands shall not request users to paste secrets into the terminal or repository.
- **DOC-REQ-072:** Relative repository links are preferred for internal artifacts.
- **DOC-REQ-073:** Raw external URLs should be minimized in normative documents in favor of governed references.
- **DOC-REQ-074:** Rendered Markdown and source Markdown shall both remain understandable.
- **DOC-REQ-075:** Unicode characters shall be used consistently and shall not create identifier ambiguity.
- **DOC-REQ-076:** Markdown validation shall include broken-link, duplicate-heading and formatting checks where tooling exists.

## 9. Naming and Repository Placement

Numbered Knowledge documents use:

`NN_UPPER_SNAKE_CASE.md`

ADRs use:

`ADR-0001-short-title.md`

Mermaid diagram sources use:

`kebab-case.mmd`

Identifiers use:

`OBDIA-<DOMAIN>-<NUMBER>`

- **DOC-REQ-077:** The authoritative numbered filename shall not include the legacy `_ENTERPRISE` suffix.
- **DOC-REQ-078:** A legacy filename variant shall not become a second normative authority.
- **DOC-REQ-079:** Filenames shall not encode transient status, branch names, personal names or review dates.
- **DOC-REQ-080:** Case-sensitive naming shall be consistent.
- **DOC-REQ-081:** Documents shall be stored in the governed directory appropriate to their class.
- **DOC-REQ-082:** Moving a governed artifact requires cross-reference and navigation updates.
- **DOC-REQ-083:** Duplicate document identifiers or authoritative filenames are blocking.
- **DOC-REQ-084:** Generated artifacts shall be distinguishable from authoritative source files.
- **DOC-REQ-085:** Temporary download suffixes such as `(1)` or `_CONSOLIDATED` shall not remain in authoritative repository filenames.
- **DOC-REQ-086:** Repository placement shall preserve separation among governance, Knowledge, ADR, specification, evidence and generated-output records.

## 10. Status and Version Control

Document lifecycle status follows the Master Documentation Constitution.

Versioning follows semantic versioning:

- major — incompatible normative or structural change;
- minor — backward-compatible substantive addition;
- patch — correction or clarification without changed normative meaning.

- **DOC-REQ-087:** Status and version are independent.
- **DOC-REQ-088:** Status shall reflect completed evidence, not intended future state.
- **DOC-REQ-089:** Draft incorporation shall not be represented as approval.
- **DOC-REQ-090:** A merged Pull Request shall not change status automatically.
- **DOC-REQ-091:** Material changes require a version increment.
- **DOC-REQ-092:** Editorial changes shall not conceal a substantive change.
- **DOC-REQ-093:** Revision history shall identify date, status, authority or owner and change summary.
- **DOC-REQ-094:** Superseded artifacts shall identify successors.
- **DOC-REQ-095:** Archived artifacts shall not be used as current authority.
- **DOC-REQ-096:** Rollback shall restore a known governed version and preserve the failed change record.

## 11. Cross-References and Traceability

- **DOC-REQ-097:** Cross-references shall use exact filenames and immutable identifiers where available.
- **DOC-REQ-098:** References shall distinguish normative dependency from informative relationship.
- **DOC-REQ-099:** Cross-reference integrity shall be validated before approval and baseline freeze.
- **DOC-REQ-100:** Broken references affecting authority, authorization, security, privacy, evidence or implementation are blocking.
- **DOC-REQ-101:** Requirements shall trace forward to decisions, controls, implementation, tests and evidence where applicable.
- **DOC-REQ-102:** Downstream artifacts shall trace back to governing requirements and ADRs.
- **DOC-REQ-103:** Supersession shall preserve historical references and identify the successor.
- **DOC-REQ-104:** Planned traceability shall not be represented as implemented or validated traceability.
- **DOC-REQ-105:** A cross-reference update shall not silently change the meaning of the referenced requirement.
- **DOC-REQ-106:** Traceability identifiers shall not contain secrets or sensitive case information.

## 12. Diagrams and Visual Documentation

Diagrams shall be used where they materially improve understanding of architecture, data flow, trust, authorization, lifecycle or threat relationships.

- **DOC-REQ-107:** Diagrams shall have a source format that can be reviewed and versioned where feasible.
- **DOC-REQ-108:** Every material diagram shall identify purpose, scope and status.
- **DOC-REQ-109:** Trust and authorization boundaries shall be shown when material.
- **DOC-REQ-110:** Diagram labels shall use Project Glossary terminology.
- **DOC-REQ-111:** Diagrams shall distinguish conceptual, proposed, implemented and validated elements.
- **DOC-REQ-112:** A diagram shall not introduce a component, trust relationship or authority absent from governing text.
- **DOC-REQ-113:** Rendered images shall retain accessible explanatory text.
- **DOC-REQ-114:** Sensitive infrastructure, credentials, personal data and unnecessary abuse-enabling detail shall not appear in public diagrams.
- **DOC-REQ-115:** Diagram changes affecting architecture require applicable ADR and review assessment.
- **DOC-REQ-116:** Documented non-applicability is acceptable when a diagram would not materially improve understanding.

## 13. Code, Command and Configuration Examples

- **DOC-REQ-117:** Examples shall be safe, minimal and appropriate to the documented environment.
- **DOC-REQ-118:** Examples shall use synthetic identities, mock services, testnets, local laboratories or other authorized controlled resources.
- **DOC-REQ-119:** Examples shall not contain production credentials, real personal data, real case data or unlawfully obtained material.
- **DOC-REQ-120:** Placeholder secrets shall be unmistakably non-secret.
- **DOC-REQ-121:** Potentially destructive commands shall include scope, effect, preconditions and recovery guidance.
- **DOC-REQ-122:** Interactive commands shall be identified to prevent accidental pasted input.
- **DOC-REQ-123:** Failure-handling examples shall not terminate an unrelated interactive shell without warning.
- **DOC-REQ-124:** Expected terminal output shall be labeled as output.
- **DOC-REQ-125:** Example code shall not implement prohibited capability.
- **DOC-REQ-126:** A documentation disclaimer shall not compensate for unsafe example behavior.

## 14. Security, Privacy and Disclosure Controls

Documentation shall be reviewed for:

- secrets and credentials;
- personal and case data;
- internal-only identifiers;
- unsafe procedures;
- abuse-enabling detail;
- vulnerable configuration;
- trust assumptions;
- legal and jurisdictional claims;
- implementation overstatement;
- supply-chain and dependency disclosure risk.

- **DOC-REQ-127:** Secrets shall not be stored in documentation, examples, history or generated output.
- **DOC-REQ-128:** Personal or case data shall be excluded unless expressly authorized and governed.
- **DOC-REQ-129:** Public documents shall use synthetic, fictitious or otherwise approved safe data.
- **DOC-REQ-130:** Sensitive material shall be generalized, restricted or withheld when publication creates disproportionate risk.
- **DOC-REQ-131:** Withholding detail shall not conceal a material limitation or fabricate security.
- **DOC-REQ-132:** Evidence integrity and chain-of-custody language shall distinguish source evidence from analysis.
- **DOC-REQ-133:** AI-generated content shall not be represented as original evidence.
- **DOC-REQ-134:** A security property shall identify scope, assumptions and validation state.
- **DOC-REQ-135:** A framework mapping shall not be represented as proof of compliance.
- **DOC-REQ-136:** Disclosure review shall consider repository history, not only the current tree.

## 15. Source, Citation and Evidence Presentation

- **DOC-REQ-137:** Material factual claims shall cite sources according to `OBDIA-RES-001`.
- **DOC-REQ-138:** Citations shall support the exact claims they accompany.
- **DOC-REQ-139:** Laws, standards, guidance and project requirements shall remain distinct.
- **DOC-REQ-140:** Time-sensitive claims shall include or link to verification dates.
- **DOC-REQ-141:** Quotations shall preserve context and attribution.
- **DOC-REQ-142:** Evidence records shall identify exact artifacts, versions, environments and methods.
- **DOC-REQ-143:** Test results shall not be generalized beyond their tested scope.
- **DOC-REQ-144:** Failed, negative and inconclusive results shall remain visible when material.
- **DOC-REQ-145:** Fabricated sources, identifiers, metrics, approvals and validation evidence are prohibited.
- **DOC-REQ-146:** Research uncertainty and contrary evidence shall be recorded where material.

## 16. AI-Assisted Documentation

AI tools may support drafting, editing, terminology review, consistency analysis and formatting, subject to human accountability.

- **DOC-REQ-147:** An identified human owner remains responsible for AI-assisted content.
- **DOC-REQ-148:** An AI system shall not be listed as approval authority or independent reviewer.
- **DOC-REQ-149:** AI-generated normative content shall receive human review before incorporation.
- **DOC-REQ-150:** AI output shall be treated as untrusted until checked against governing authorities and sources.
- **DOC-REQ-151:** Secrets, real credentials, private case data and unauthorized personal data shall not be submitted to unapproved AI services.
- **DOC-REQ-152:** Material AI limitations or uncertain claims shall be disclosed.
- **DOC-REQ-153:** AI assistance shall not be used to fabricate independent assurance, citations, tests, approvals or implementation.
- **DOC-REQ-154:** Tool-generated changes shall remain attributable through repository history.
- **DOC-REQ-155:** External content used by an AI tool shall not override constitutional or repository authority.
- **DOC-REQ-156:** Human reviewers shall verify security, privacy, evidence and legal-boundary statements.

## 17. Accessibility and Reviewer Readability

- **DOC-REQ-157:** Documents shall use descriptive headings and logical navigation.
- **DOC-REQ-158:** Images and diagrams shall have meaningful alternative or adjacent explanatory text.
- **DOC-REQ-159:** Meaning shall not depend solely on color, position or visual styling.
- **DOC-REQ-160:** Tables shall not be used when they materially reduce readability or accessibility.
- **DOC-REQ-161:** Long documents should include a table of contents or clear section structure.
- **DOC-REQ-162:** Critical security, privacy and limitation information shall not be hidden in footnotes or appendices.
- **DOC-REQ-163:** Recruiter readability shall not remove technical depth or material limitations.
- **DOC-REQ-164:** Abbreviated navigation documents shall link to authoritative details.
- **DOC-REQ-165:** Terminology shall remain consistent across prose, tables, diagrams and examples.

## 18. Review and Validation Gate

Before a governed document advances beyond Draft, reviewers shall verify:

- metadata completeness;
- authority and hierarchy;
- purpose and scope;
- dependencies and cross-references;
- terminology;
- normative-language consistency;
- security, privacy, ethics and legal boundaries;
- implementation and publication status;
- source and citation integrity;
- traceability;
- Markdown and rendering quality;
- accessibility;
- revision history;
- unresolved findings and limitations.

- **DOC-REQ-166:** The review shall identify the exact artifact and commit.
- **DOC-REQ-167:** Applicable constitutional gates shall be identified.
- **DOC-REQ-168:** Gate non-applicability requires rationale.
- **DOC-REQ-169:** Blocking findings shall prevent lifecycle progression.
- **DOC-REQ-170:** Automated linting may provide evidence but shall not replace human review.
- **DOC-REQ-171:** Reviewer role concentration shall be disclosed.
- **DOC-REQ-172:** Independent review shall not be claimed without evidence.
- **DOC-REQ-173:** A material post-review change invalidates affected review evidence.
- **DOC-REQ-174:** Approval shall identify actor, role, date, scope and exact artifact.
- **DOC-REQ-175:** Publication requires the separate release gate.

## 19. Minimum Validation Checklist

Before approval, confirm:

- [ ] filename and Document ID are unique and valid;
- [ ] version and lifecycle status are accurate;
- [ ] classification, authority, owner and approval authority are explicit;
- [ ] purpose, scope, dependencies and assumptions are complete;
- [ ] mandatory sections contain substantive content;
- [ ] terminology conforms to `OBDIA-GLOSS-001`;
- [ ] facts, proposals, implementation, tests and evidence are distinct;
- [ ] requirements use stable identifiers and normative language;
- [ ] cross-references resolve;
- [ ] traceability is present where applicable;
- [ ] diagrams and examples are safe and accurately labeled;
- [ ] no secret, credential, personal data or real case data is present;
- [ ] security, privacy and fundamental-rights considerations are complete;
- [ ] limitations and unresolved findings are visible;
- [ ] citations support material claims;
- [ ] implementation and publication status are accurate;
- [ ] Markdown renders correctly and remains source-readable;
- [ ] applicable review gates and exact commit evidence are recorded;
- [ ] revision history is current;
- [ ] Project Founder approval is recorded before status becomes Approved.

## 20. Change Control

Every material change to this standard shall include:

- change identifier;
- rationale;
- affected document classes and artifacts;
- dependency analysis;
- security impact;
- privacy and fundamental-rights impact;
- implementation and publication impact;
- migration requirements;
- validation plan;
- rollback plan;
- approval record;
- semantic version change.

- **DOC-REQ-176:** Editorial corrections use a patch version when meaning is unchanged.
- **DOC-REQ-177:** Backward-compatible substantive additions use a minor version.
- **DOC-REQ-178:** Incompatible normative or documentation-method changes use a major version.
- **DOC-REQ-179:** Changes require Project Founder approval.
- **DOC-REQ-180:** Unapproved documentation-methodology changes during an active phase are prohibited.
- **DOC-REQ-181:** Accepted ADRs and historical evidence shall not be silently rewritten to conform retroactively.
- **DOC-REQ-182:** Migration shall update affected templates, indexes, automation and cross-references.

## 21. Limitations

- This standard does not replace specialized architecture, security, privacy, legal, evidentiary, implementation or publication review.
- Conforming formatting does not prove technical correctness.
- Complete documentation does not prove implementation or validation.
- Internal review does not constitute independent certification or institutional endorsement.
- Automated linting cannot determine all semantic, legal, security or privacy defects.
- A navigation-only README may use a reduced metadata profile but cannot create normative authority.
- Future operational use may require institution-specific documentation controls beyond this independent research baseline.

## 22. Consolidation Record

Version 1.1.0 consolidates the existing Documentation Standard without expanding project scope. It:

- retains immutable document identifier `OBDIA-DOC-001`;
- normalizes the authoritative filename to `11_DOCUMENTATION_STANDARD.md`;
- replaces `Enterprise Baseline (Draft)` with constitutional lifecycle status `Draft`;
- preserves all original mandatory metadata fields;
- preserves precision over marketing, separation of facts from hypotheses, consistent terminology and GitHub-ready Markdown;
- adds documentation classes, metadata profiles, mandatory sections, normative language, naming, lifecycle, traceability, diagram, example, security, privacy, citation, AI-assistance, accessibility, review and validation controls;
- treats the former `_ENTERPRISE` filename as a legacy source variant rather than a second normative document;
- creates no new normative document and authorizes no implementation or publication.

## 23. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | Prior baseline | Enterprise Baseline (Draft) | Project origin record | Original Documentation Standard defining mandatory metadata and four writing principles. |
| 1.1.0 | 2026-08-05 | Draft | Documentation Authority; approval reserved to Project Founder | Constitutional consolidation: normalized filename, metadata and lifecycle status; preserved original requirements; added document classes, structure, normative language, naming, traceability, security, privacy, AI-assistance, accessibility, validation and change-control requirements. |
