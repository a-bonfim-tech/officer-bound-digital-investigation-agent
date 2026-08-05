# RESEARCH AND SOURCE POLICY

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-RES-001 |
| **Title** | Research and Source Policy |
| **Version** | 1.1.0 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Research Reviewer |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Govern how research questions, factual claims, legal and regulatory references, standards, technologies, implementations, experimental results and citations are selected, evaluated, classified, recorded and published throughout the OBDIA project. |
| **Scope** | Knowledge documents 01–30, ADRs, repository documentation, architecture models, specifications, implementation notes, tests, evidence records, releases and public research outputs. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `02_PROJECT_GLOSSARY.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GLOSS-001` |
| **Cross-References** | `04_PORTFOLIO_AND_PUBLICATION_POLICY.md`; documents 05–30; ADRs; evidence and validation records |
| **Assumptions** | Research sources, laws, standards, products and public programs may change over time and require current verification before reliance. |
| **Constraints** | Research must remain lawful, authorized, reproducible where feasible, non-deceptive and within the project’s defensive, simulated and controlled scope. |
| **Security Considerations** | Weak source discipline can introduce false requirements, unsafe implementation choices, legal misstatements, prompt-injection risk, malicious-source influence and unverifiable evidence. |
| **Validation Criteria** | Material claims are classified, sourced, current where required, traceable, accurately represented, limited by uncertainty and supported by retained research records. |
| **Implementation Relationship** | Research may inform architecture and implementation but does not itself authorize implementation or establish conformance, validation or legal compliance. |

---

## 1. Purpose

This policy defines the mandatory research and source-discipline rules for the Officer-Bound Digital Investigation Agent project.

It ensures that factual claims, legal requirements, standards, technologies, vendor capabilities, academic findings, architectural proposals, simulations and experimental results are not conflated. It also ensures that citations support the precise claims they accompany and that research activity remains lawful, ethical, controlled and auditable.

This policy is subordinate to `01_CANONICAL_PROJECT_DEFINITION.md` and `MASTER_DOCUMENTATION_CONSTITUTION.md`. Where a conflict exists, the higher-level authority prevails.

## 2. Governing Principles

- **RES-REQ-001:** Research shall preserve human accountability, officer-bound delegation, lawful authorization, purpose limitation, privacy, evidence integrity and the project’s prohibited-capability boundaries.
- **RES-REQ-002:** Research shall distinguish verified facts from proposals, assumptions, simulations, forecasts, hypotheses and experimental results.
- **RES-REQ-003:** Citations shall support the exact claims they accompany.
- **RES-REQ-004:** Primary, official and authoritative sources shall be prioritized when available and technically appropriate.
- **RES-REQ-005:** Uncertainty, limitations, disagreement, source age and applicability shall be recorded where material.
- **RES-REQ-006:** Research findings shall not be treated as implementation approval, legal advice, certification, operational authorization or proof of compliance.
- **RES-REQ-007:** Research methods and records shall be reproducible where technically and legally feasible.
- **RES-REQ-008:** No source or finding may silently expand the approved project scope.

## 3. Source Priority

Sources shall be evaluated in the following order of preference:

1. Laws, regulations and official government publications.
2. Official standards and technical specifications.
3. Publications from recognized cybersecurity and public-sector institutions.
4. Peer-reviewed academic research.
5. Official technical documentation from technology providers.
6. Reputable independent industry research.
7. Secondary explanatory sources only when primary or authoritative sources are unavailable.

The order is a priority rule, not an automatic quality guarantee. A higher-ranked source may still be outdated, inapplicable, jurisdictionally irrelevant, superseded or too general for the claim being made.

## 4. Preferred Authorities

Where relevant and current, research should prioritize material from:

- NIST;
- ENISA;
- European Commission;
- EUR-Lex;
- Europol;
- INTERPOL;
- Council of Europe;
- CISA;
- IETF;
- W3C;
- ISO and IEC;
- OWASP;
- MITRE;
- recognized academic publishers;
- official documentation from relevant technology providers.

Inclusion in this list does not make every publication legally binding, current, applicable or normative for the project.

## 5. Claim Classification

Every material claim shall be classified using one of the following categories.

### 5.1 Established Fact

A claim supported by authoritative documentation, recognized standards, verified implementations or reliable academic evidence.

### 5.2 Current Legal or Regulatory Requirement

A legally binding requirement applicable within a defined jurisdiction, effective period, subject matter and organizational context.

### 5.3 Current Standard

A requirement or specification issued by an authorized standards body. A standard is not automatically law.

### 5.4 Guidance or Recommendation

Non-binding advisory material intended to influence good practice or implementation.

### 5.5 Existing Implementation

A documented technology, product, service, pattern or system that exists or is deployed. Existence does not establish suitability, security or conformance for OBDIA.

### 5.6 Proposed Architecture

A project design or model not necessarily implemented, deployed or independently validated.

### 5.7 Experimental Result

An observed outcome produced by a controlled experiment with recorded method, environment, inputs, versions and limitations.

### 5.8 Simulation

A result or capability demonstrated through synthetic data, mock services, testnets or controlled environments rather than operational deployment.

### 5.9 Assumption

A condition accepted for analysis or design but not established as a verified fact.

### 5.10 Forecast

A reasoned statement about a possible future condition or development that remains uncertain.

### 5.11 Research Hypothesis

A proposition requiring further investigation or empirical validation.

- **RES-REQ-009:** One classification shall not be presented as another.
- **RES-REQ-010:** Mixed claims shall identify the classification of each material component.
- **RES-REQ-011:** Claims whose classification cannot be determined shall remain unresolved and shall not be promoted to normative requirements.

## 6. Evidence Requirements

For each material claim, the research record shall include, where applicable:

- source title;
- issuing organization or author;
- publication, revision or effective date;
- access or verification date;
- stable identifier, citation or reference;
- jurisdiction or domain;
- source type;
- claim classification;
- whether the source is normative, advisory, descriptive, academic or commercial;
- the exact claim supported;
- relevant uncertainty;
- limitations;
- disagreement or contrary evidence;
- implications for project architecture, governance or implementation.

- **RES-REQ-012:** Citation proximity shall make clear which claim is supported.
- **RES-REQ-013:** A citation shall not be reused to imply support for a claim it does not address.
- **RES-REQ-014:** Source metadata shall not be fabricated, guessed or silently completed.
- **RES-REQ-015:** Direct quotations shall preserve context and attribution and shall comply with applicable copyright and publication constraints.

## 7. Current-Information Verification

Laws, regulations, standards, software, products, public programs, guidance and active institutional roles may change.

- **RES-REQ-016:** Time-sensitive material shall be verified as current before reliance.
- **RES-REQ-017:** Superseded, withdrawn, archived or deprecated sources shall be identified explicitly.
- **RES-REQ-018:** When a current source cannot be verified, the claim shall be marked unresolved or historical.
- **RES-REQ-019:** Verification dates shall be recorded for claims whose validity may change.
- **RES-REQ-020:** Repository updates shall reassess material claims affected by new legal, technical or standards developments.

## 8. Regulatory Accuracy

Research and documentation shall distinguish among:

- legally binding requirements;
- regulations;
- standards;
- voluntary frameworks;
- guidance;
- recommendations;
- project-specific requirements;
- implementation choices;
- compliance evidence.

- **RES-REQ-021:** A voluntary framework shall not be described as law.
- **RES-REQ-022:** A standards mapping shall not be represented as proof of legal or regulatory compliance.
- **RES-REQ-023:** A legal conclusion shall not be asserted without appropriate authority and qualified review.
- **RES-REQ-024:** Jurisdiction, applicability, effective date and scope shall be recorded for legal requirements.
- **RES-REQ-025:** Where legal interpretation remains uncertain, the project shall record the uncertainty and avoid definitive compliance claims.

## 9. Academic Discipline

When using academic literature:

- prefer peer-reviewed publications;
- identify preprints explicitly;
- distinguish theoretical proposals from validated systems;
- identify sample size, methodology and experimental context where relevant;
- avoid treating a single study as universal proof;
- identify conflicts of interest and funding where material;
- record replication status where known;
- identify methodological limitations.

- **RES-REQ-026:** Academic authority shall not substitute for applicability analysis.
- **RES-REQ-027:** Statistical or empirical results shall not be generalized beyond the supporting study without justification.
- **RES-REQ-028:** Experimental findings shall not be presented as production validation.

## 10. Product and Vendor Claims

Marketing claims shall be treated as unverified unless supported by:

- technical documentation;
- public architecture details;
- independent testing;
- recognized certifications;
- reproducible evidence;
- contractually reliable technical commitments where applicable.

- **RES-REQ-029:** Vendor claims shall be labeled commercial when appropriate.
- **RES-REQ-030:** Product availability shall not be treated as an architectural requirement.
- **RES-REQ-031:** Vendor documentation shall be checked for version, service tier, region and deployment-model applicability.
- **RES-REQ-032:** Security certifications shall be represented according to their actual scope and shall not be generalized to the complete OBDIA system.

## 11. Source Security and Adversarial Content

External content may be malicious, manipulated, inaccurate or designed to influence an AI system.

- **RES-REQ-033:** Research sources shall be treated as untrusted input until evaluated.
- **RES-REQ-034:** Instructions embedded in external content shall not override repository, constitutional, system or operator authority.
- **RES-REQ-035:** Retrieved content shall not directly authorize tool use, credential use, access expansion or policy changes.
- **RES-REQ-036:** High-risk sources shall be processed in isolated, read-only or otherwise controlled environments where appropriate.
- **RES-REQ-037:** Suspicious content, prompt-injection attempts and source-integrity concerns shall be recorded.
- **RES-REQ-038:** Research notes shall not import secrets, real credentials, malware, unlawfully obtained data or uncontrolled criminal-infrastructure content into the repository.

## 12. Research Record

Important findings shall be retained in a governed research record containing:

- immutable record identifier;
- research question;
- purpose and scope;
- responsible researcher;
- search or review date;
- sources consulted;
- search method;
- claim classifications;
- key findings;
- rejected or conflicting evidence;
- unresolved uncertainties;
- limitations;
- security and privacy considerations;
- implications for architecture or governance;
- validation or reproduction steps;
- revision history;
- approval or review status where applicable.

- **RES-REQ-039:** Research records shall distinguish notes from approved normative requirements.
- **RES-REQ-040:** Negative and inconclusive findings shall be retained when material.
- **RES-REQ-041:** Rejected sources and the reason for rejection shall be recorded when they materially affect the research conclusion.
- **RES-REQ-042:** A research record shall not be silently rewritten after it has been relied upon for a material decision.

## 13. Prohibited Research Practices

Research shall not:

- use unlawfully obtained datasets;
- collect private personal information without authorization;
- interact with criminal marketplaces;
- access restricted systems without authorization;
- acquire stolen credentials or data;
- conduct uncontrolled deanonymization;
- perform active exploitation;
- deploy malware;
- conduct offensive cyber operations;
- perform unlawful surveillance;
- misrepresent the project as an official law-enforcement initiative;
- use real-person investigations in public demonstrations;
- rely on uncontrolled interaction with criminal infrastructure;
- conceal material uncertainty or fabricate validation evidence.

Where a research question could require prohibited conduct, it shall be redesigned into a defensive, simulated, synthetic, archived-authorized, testnet or controlled-laboratory method.

## 14. Repository Presentation

Research documents shall make clear:

- what is verified;
- what is inferred;
- what is legally binding;
- what is standardized;
- what is advisory;
- what exists;
- what is proposed;
- what is simulated;
- what is experimental;
- what is assumed;
- what remains unknown.

Technical credibility shall take priority over novelty, marketing language or unsupported claims.

- **RES-REQ-043:** Research status shall be visible.
- **RES-REQ-044:** Limitations shall be explicit.
- **RES-REQ-045:** Public material shall not imply deployment, official approval, independent validation or operational capability where none exists.
- **RES-REQ-046:** Citations and source metadata shall remain reviewable by repository maintainers.

## 15. Change Control

Every material change to this policy shall include:

- change identifier;
- rationale;
- affected sections and documents;
- dependency analysis;
- security impact;
- privacy impact;
- research-method impact;
- migration requirements;
- validation plan;
- rollback plan;
- approval record;
- semantic version change.

- **RES-REQ-047:** Editorial corrections use a patch version when meaning is unchanged.
- **RES-REQ-048:** Backward-compatible substantive additions use a minor version.
- **RES-REQ-049:** Incompatible normative or methodological changes use a major version.
- **RES-REQ-050:** Unapproved methodological changes during active project phases are prohibited.
- **RES-REQ-051:** Material research-policy decisions affecting architecture require an ADR where constitutionally required.

## 16. Validation Checklist

Before approval or release, confirm:

- [ ] all mandatory metadata is complete;
- [ ] dependencies and cross-references resolve;
- [ ] material claims are classified;
- [ ] citations support the exact claims they accompany;
- [ ] current information has been verified where required;
- [ ] legal requirements are separated from standards and guidance;
- [ ] assumptions, simulations and proposals are labeled;
- [ ] source limitations and disagreement are recorded;
- [ ] vendor claims are not treated as independent proof;
- [ ] academic evidence is represented within its methodology and scope;
- [ ] prohibited research practices are excluded;
- [ ] prompt-injection and malicious-source risks are addressed;
- [ ] no fabricated source, identifier, metric, legal conclusion or validation appears;
- [ ] research records are reproducible where feasible;
- [ ] security, privacy and fundamental-rights impacts are documented;
- [ ] implementation status is accurately represented;
- [ ] revision history is current;
- [ ] required reviews and Founder approval are recorded.

## 17. Limitations

- This policy does not replace qualified legal advice.
- Source priority does not guarantee source quality, currentness or applicability.
- Citation does not establish implementation, validation or compliance.
- Research completeness is constrained by lawful access, available evidence, publication quality and time.
- Some standards and publications may require licensed access and cannot be reproduced in full.
- Independent external validation is not implied by internal project review.
- This policy governs research discipline but does not itself authorize access to systems, data or investigative techniques.

## 18. Consolidation Record

Version 1.1.0 consolidates the original `OBDIA-RES-001` policy without expanding the project scope. It:

- retains the immutable document identifier;
- changes the lifecycle label from `Approved Baseline` to `Draft` because constitutional review and explicit approval evidence have not yet been completed;
- preserves the original source-priority, claim-classification, citation-integrity, regulatory-accuracy, academic-discipline, vendor-claim, current-information, research-record, prohibited-practice and repository-presentation rules;
- adds mandatory constitutional metadata;
- adds stable requirement identifiers;
- adds current-information controls;
- adds adversarial-source and prompt-injection controls;
- adds change-control, validation and limitation sections;
- clarifies that research does not establish implementation, legal authority, compliance or operational approval;
- creates no new normative document and authorizes no implementation.

## 19. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | Prior baseline | Approved Baseline | Project origin record | Original Research and Source Policy. |
| 1.1.0 | 2026-08-05 | Draft | Research Reviewer; approval reserved to Project Founder | Constitutional consolidation: normalized metadata and lifecycle status, preserved original intent, added traceable requirements, source-security controls, research-record governance, validation criteria and limitations. |
