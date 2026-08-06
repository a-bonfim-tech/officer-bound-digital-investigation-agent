# DIAGRAM STANDARD

| Metadata Field | Value |
|---|---|
| **Document ID** | OBDIA-DIAG-001 |
| **Title** | Diagram Standard |
| **Version** | 1.0.1 |
| **Status** | Draft |
| **Classification** | Normative |
| **Authority** | Project Founder |
| **Owner** | Documentation Authority |
| **Approval Authority** | Project Founder |
| **Effective Date** | Not applicable while Draft |
| **Purpose** | Define mandatory diagram types, metadata, identifiers, filenames, source and rendered artifacts, notation, trust and authorization semantics, evidence and lifecycle representation, accessibility, validation, review, change control and traceability for OBDIA architecture and engineering diagrams. |
| **Scope** | Mermaid, C4 Model, Data Flow Diagrams, Sequence Diagrams, Trust Boundary Diagrams, Attack Trees, state and lifecycle diagrams, deployment and network diagrams, evidence and chain-of-custody diagrams, authorization flows, connector diagrams, implementation diagrams and supporting rendered artifacts within Knowledge Pack 01–30. |
| **Dependencies** | `01_CANONICAL_PROJECT_DEFINITION.md`; `MASTER_DOCUMENTATION_CONSTITUTION.md`; `02_PROJECT_GLOSSARY.md`; `05_ARCHITECTURE_PRINCIPLES.md`; `07_THREAT_MODEL_BASELINE.md`; `08_TRUST_MODEL.md`; `10_ADR_POLICY.md`; `11_DOCUMENTATION_STANDARD.md`; `13_REPOSITORY_STATE_MODEL.md`; `14_TERMINOLOGY_AND_NAMING.md`; `16_SECURITY_ARCHITECTURE_BASELINE.md`; `17_AUTHORIZATION_MODEL.md`; `18_EVIDENCE_MODEL.md`; `19_AGENT_LIFECYCLE_MODEL.md`; `20_CONNECTOR_SECURITY_POLICY.md`; `21_SECURE_CODING_STANDARD.md`; `22_TESTING_STANDARD.md`; forward dependencies `24_GITHUB_REPOSITORY_STANDARD.md`, `25_DOCUMENT_VERSIONING_POLICY.md`, `29_DECISION_LOG_POLICY.md` and `30_COMPLIANCE_MAPPING_BASELINE.md` |
| **Normative References** | `OBDIA-CANON-001`; `OBDIA-CONST-001`; `OBDIA-GLOSS-001`; `OBDIA-ARCH-001`; `OBDIA-TM-001`; `OBDIA-TRUST-001`; `OBDIA-ADR-001`; `OBDIA-DOC-001`; `OBDIA-STATE-001`; `OBDIA-NAME-001`; `OBDIA-SEC-001`; `OBDIA-AUTH-001`; `OBDIA-EVID-001`; `OBDIA-AGENT-001`; `OBDIA-CONN-001`; `OBDIA-CODE-001`; `OBDIA-TEST-001` |
| **Cross-References** | Documents 24–30; accepted ADRs; threat models; trust models; authorization and evidence schemas; lifecycle records; connector manifests; implementation plans; test evidence; risk, exception, change and release records |
| **Assumptions** | Diagrams describe project architecture, proposed designs, controlled implementations, simulations and test environments using governed terminology and synthetic or authorized examples. |
| **Constraints** | Diagrams shall not imply independent AI legal authority, hidden operational capability, uncontrolled criminal-infrastructure access, real-person public investigations, certification, compliance, production deployment or validation beyond supporting evidence. |
| **Security Considerations** | Diagrams can expose secrets, internal endpoints, real identities, case data, credential paths, privileged trust assumptions, vulnerable dependencies, misleading boundaries and operationally sensitive architecture. |
| **Validation Criteria** | Every authoritative diagram has an immutable identifier, version, status, scope, related documents, notation, legend, source file, rendered representation where needed, review evidence, accurate trust and authority semantics, accessible text and bidirectional traceability. |
| **Implementation Relationship** | This standard governs visual documentation. It does not implement architecture, authorize access, select a diagramming product, certify a design or establish production readiness. |

---

## 1. Purpose

This document consolidates the OBDIA Diagram Standard.

The original baseline preferred:

- Mermaid;
- C4 Model;
- Data Flow Diagram;
- Sequence Diagram;
- Trust Boundary Diagram;
- Attack Tree.

It also required every diagram to identify:

- version;
- scope;
- related documents.

This consolidation preserves all six preferred formats and all three mandatory metadata elements. It adds the notation, repository, security, accessibility, review, validation and traceability controls required for professional and reproducible architecture documentation.

A diagram is a governed representation of a defined scope. It is not the architecture itself, not implementation evidence, not authorization and not proof that a control is effective.


## 2. Fundamental Diagram Rules

- **DIAG-REQ-001:** Every authoritative diagram shall identify its version.
- **DIAG-REQ-002:** Every authoritative diagram shall identify its scope.
- **DIAG-REQ-003:** Every authoritative diagram shall identify related documents.
- **DIAG-REQ-004:** Every authoritative diagram shall have an immutable diagram identifier.
- **DIAG-REQ-005:** Every authoritative diagram shall identify its status.
- **DIAG-REQ-006:** Every authoritative diagram shall identify its owner.
- **DIAG-REQ-007:** Every authoritative diagram shall identify its notation or diagram type.
- **DIAG-REQ-008:** Every authoritative diagram shall identify its authoritative source file.
- **DIAG-REQ-009:** Rendered artifacts shall remain traceable to their source files.
- **DIAG-REQ-010:** Diagram source shall be preferred over manually edited rendered output.
- **DIAG-REQ-011:** Diagram semantics shall remain subordinate to governing normative documents.
- **DIAG-REQ-012:** A diagram shall not create architecture by implication.
- **DIAG-REQ-013:** A diagram shall not create authorization by implication.
- **DIAG-REQ-014:** A diagram shall not create legal authority by implication.
- **DIAG-REQ-015:** A diagram shall not establish implementation status by implication.
- **DIAG-REQ-016:** A diagram shall not establish validation status by implication.
- **DIAG-REQ-017:** A diagram shall not establish compliance by implication.
- **DIAG-REQ-018:** A diagram shall not conceal known uncertainty.
- **DIAG-REQ-019:** A diagram shall not combine incompatible scopes without explicit boundaries.
- **DIAG-REQ-020:** A diagram shall use governed terminology.
- **DIAG-REQ-021:** A diagram shall distinguish human, agent, workload, connector and external-system actors.
- **DIAG-REQ-022:** A diagram shall distinguish identity from authorization.
- **DIAG-REQ-023:** A diagram shall distinguish source evidence from derived analysis.
- **DIAG-REQ-024:** A diagram shall distinguish repository lifecycle from agent, connector and evidence states.
- **DIAG-REQ-025:** A diagram shall identify trust boundaries explicitly when trust is material.
- **DIAG-REQ-026:** A diagram shall identify authorization enforcement points when authorization is material.
- **DIAG-REQ-027:** A diagram shall identify data and evidence flows when data handling is material.
- **DIAG-REQ-028:** A diagram shall identify failure or containment paths when security behavior is material.
- **DIAG-REQ-029:** A diagram shall be reviewable without relying on undocumented visual conventions.
- **DIAG-REQ-030:** Critical diagram inaccuracies shall block approval or release of the affected artifact.

## 3. Diagram Record and Metadata

- **DIAG-REQ-031:** Every diagram record shall have an immutable identifier.
- **DIAG-REQ-032:** Every diagram record shall identify title.
- **DIAG-REQ-033:** Every diagram record shall identify version.
- **DIAG-REQ-034:** Every diagram record shall identify status.
- **DIAG-REQ-035:** Every diagram record shall identify classification.
- **DIAG-REQ-036:** Every diagram record shall identify owner.
- **DIAG-REQ-037:** Every diagram record shall identify approval authority.
- **DIAG-REQ-038:** Every diagram record shall identify diagram type.
- **DIAG-REQ-039:** Every diagram record shall identify purpose.
- **DIAG-REQ-040:** Every diagram record shall identify scope.
- **DIAG-REQ-041:** Every diagram record shall identify audience.
- **DIAG-REQ-042:** Every diagram record shall identify related normative documents.
- **DIAG-REQ-043:** Every diagram record shall identify related ADRs.
- **DIAG-REQ-044:** Every diagram record shall identify related requirements where material.
- **DIAG-REQ-045:** Every diagram record shall identify assumptions.
- **DIAG-REQ-046:** Every diagram record shall identify exclusions.
- **DIAG-REQ-047:** Every diagram record shall identify notation.
- **DIAG-REQ-048:** Every diagram record shall identify legend.
- **DIAG-REQ-049:** Every diagram record shall identify source-file path.
- **DIAG-REQ-050:** Every diagram record shall identify rendered-file paths where applicable.
- **DIAG-REQ-051:** Every diagram record shall identify generation tool and version where material.
- **DIAG-REQ-052:** Every diagram record shall identify last governed review.
- **DIAG-REQ-053:** Every diagram record shall identify known limitations.
- **DIAG-REQ-054:** Every diagram record shall identify security considerations.
- **DIAG-REQ-055:** Every diagram record shall identify privacy considerations where material.
- **DIAG-REQ-056:** Every diagram record shall identify accessibility alternative.
- **DIAG-REQ-057:** Every diagram record shall identify change-history reference.
- **DIAG-REQ-058:** Every diagram record shall identify supersession relationship where applicable.
- **DIAG-REQ-059:** Every diagram record shall identify validation evidence where a validated rendering is claimed.
- **DIAG-REQ-060:** Malformed or incomplete diagram records shall not be treated as authoritative.

## 4. Diagram Identifiers and Filenames

- **DIAG-REQ-061:** Diagram identifiers shall conform to `OBDIA-NAME-001`.
- **DIAG-REQ-062:** Diagram identifiers shall be immutable.
- **DIAG-REQ-063:** Diagram identifiers shall not be reused.
- **DIAG-REQ-064:** Diagram identifiers shall remain stable when titles change.
- **DIAG-REQ-065:** Diagram identifiers shall remain stable when rendering formats change.
- **DIAG-REQ-066:** Materially different semantic diagrams shall receive distinct identifiers.
- **DIAG-REQ-067:** Diagram source filenames shall use kebab-case.
- **DIAG-REQ-068:** Diagram filenames shall describe subject rather than lifecycle state alone.
- **DIAG-REQ-069:** Diagram filenames shall avoid real person names.
- **DIAG-REQ-070:** Diagram filenames shall avoid real case identifiers.
- **DIAG-REQ-071:** Diagram filenames shall avoid credentials and secrets.
- **DIAG-REQ-072:** Diagram filenames shall avoid ambiguous abbreviations.
- **DIAG-REQ-073:** Diagram filenames shall identify view or level where useful.
- **DIAG-REQ-074:** Diagram filenames shall identify notation where ambiguity would otherwise exist.
- **DIAG-REQ-075:** Rendered filenames shall preserve the source basename.
- **DIAG-REQ-076:** Rendered filenames shall use an extension appropriate to their format.
- **DIAG-REQ-077:** Source and rendered filenames shall not collide.
- **DIAG-REQ-078:** Legacy filenames shall be mapped through documented aliases or migration records.
- **DIAG-REQ-079:** Filename changes shall preserve repository history.
- **DIAG-REQ-080:** Filename changes shall update cross-references.
- **DIAG-REQ-081:** Filename changes shall not alter immutable diagram identifiers.
- **DIAG-REQ-082:** Generated filenames shall be deterministic where practical.
- **DIAG-REQ-083:** Temporary rendering files shall not become authoritative artifacts.
- **DIAG-REQ-084:** Duplicate canonical filenames shall be rejected.
- **DIAG-REQ-085:** Case-only filename differences shall be prohibited.

## 5. Authoritative Source and Rendered Artifacts

- **DIAG-REQ-086:** Every diagram shall identify one authoritative source representation.
- **DIAG-REQ-087:** Text-based source shall be preferred where it preserves required semantics.
- **DIAG-REQ-088:** Mermaid source shall use `.mmd` or an approved Markdown-embedded representation.
- **DIAG-REQ-089:** Other text-based diagram sources shall use documented extensions.
- **DIAG-REQ-090:** Rendered SVG shall be preferred for scalable repository display where supported.
- **DIAG-REQ-091:** PNG may be used when SVG is unavailable or unsuitable.
- **DIAG-REQ-092:** PDF may be used for publication packages when required.
- **DIAG-REQ-093:** Rendered output shall not be edited independently of source.
- **DIAG-REQ-094:** Manual annotations to rendered output shall be incorporated into authoritative source or recorded as a distinct artifact.
- **DIAG-REQ-095:** Rendered output shall identify or remain traceable to source version.
- **DIAG-REQ-096:** Render commands or procedures shall be documented.
- **DIAG-REQ-097:** Render-tool versions shall be recorded where output differences are material.
- **DIAG-REQ-098:** Rendering shall be reproducible where practical.
- **DIAG-REQ-099:** Rendering failures shall remain visible.
- **DIAG-REQ-100:** Rendered artifacts shall not contain hidden layers with secrets or sensitive data.
- **DIAG-REQ-101:** Rendered metadata shall be reviewed for sensitive information.
- **DIAG-REQ-102:** Embedded fonts or assets shall comply with licensing and repository policy.
- **DIAG-REQ-103:** External image dependencies shall be avoided for authoritative diagrams.
- **DIAG-REQ-104:** Remote diagram rendering shall not receive secrets or real case data.
- **DIAG-REQ-105:** Generated artifacts shall be regenerated after source changes.
- **DIAG-REQ-106:** Stale rendered artifacts shall not be represented as current.
- **DIAG-REQ-107:** Source/render mismatches shall block release of the affected diagram.
- **DIAG-REQ-108:** Render validation shall compare identifiers, version and content where feasible.
- **DIAG-REQ-109:** Authoritative source shall remain available after archival.
- **DIAG-REQ-110:** Rendered artifacts shall remain accessible to reviewers without proprietary software where feasible.

## 6. Preferred Diagram Types

- **DIAG-REQ-111:** Mermaid shall be a preferred format for repository-native text diagrams.
- **DIAG-REQ-112:** C4 Model shall be a preferred format for architecture abstraction levels.
- **DIAG-REQ-113:** Data Flow Diagram shall be a preferred format for data movement and trust boundaries.
- **DIAG-REQ-114:** Sequence Diagram shall be a preferred format for ordered interactions.
- **DIAG-REQ-115:** Trust Boundary Diagram shall be a preferred format for security boundary analysis.
- **DIAG-REQ-116:** Attack Tree shall be a preferred format for adversarial decomposition.
- **DIAG-REQ-117:** State diagrams may be used for governed lifecycle and transition semantics.
- **DIAG-REQ-118:** Deployment diagrams may be used for environment and hosting views.
- **DIAG-REQ-119:** Network diagrams may be used for segmentation and egress views.
- **DIAG-REQ-120:** Evidence-flow diagrams may be used for acquisition, provenance and custody.
- **DIAG-REQ-121:** Authorization-flow diagrams may be used for PDP, PEP, PAP and PIP interactions.
- **DIAG-REQ-122:** Connector diagrams may be used for external integration boundaries.
- **DIAG-REQ-123:** Decision trees may be used when decisions are deterministic and governed.
- **DIAG-REQ-124:** Flowcharts may be used for procedural logic when another notation is not more precise.
- **DIAG-REQ-125:** Entity-relationship diagrams may be used for schema and data relationships.
- **DIAG-REQ-126:** Timing diagrams may be used when ordering and time constraints are central.
- **DIAG-REQ-127:** Tables shall be preferred over diagrams when tabular form is clearer.
- **DIAG-REQ-128:** Prose shall be preferred over diagrams when visual representation adds no clarity.
- **DIAG-REQ-129:** Multiple complementary diagrams may be used for complex scopes.
- **DIAG-REQ-130:** One diagram shall not combine every architectural concern if separation improves review.
- **DIAG-REQ-131:** Diagram type selection shall be justified by communication need.
- **DIAG-REQ-132:** Diagram type shall not be selected solely for visual appearance.
- **DIAG-REQ-133:** Unsupported notation shall require documented rationale.
- **DIAG-REQ-134:** Notation changes shall preserve semantic traceability.
- **DIAG-REQ-135:** Preferred formats shall not override security, accessibility or reproducibility requirements.

## 7. Mermaid Requirements

- **DIAG-REQ-136:** Mermaid diagrams shall declare a supported diagram type.
- **DIAG-REQ-137:** Mermaid source shall be syntactically valid.
- **DIAG-REQ-138:** Mermaid source shall use stable identifiers for material nodes.
- **DIAG-REQ-139:** Mermaid node labels shall use governed terminology.
- **DIAG-REQ-140:** Mermaid source shall avoid embedding secrets.
- **DIAG-REQ-141:** Mermaid source shall avoid real case data in public artifacts.
- **DIAG-REQ-142:** Mermaid links shall use explicit labels when flow meaning is material.
- **DIAG-REQ-143:** Mermaid subgraphs shall represent meaningful boundaries.
- **DIAG-REQ-144:** Mermaid styling shall not be the sole carrier of semantic meaning.
- **DIAG-REQ-145:** Mermaid colors shall not be the sole carrier of status or trust meaning.
- **DIAG-REQ-146:** Mermaid classes shall have documented semantics where used.
- **DIAG-REQ-147:** Mermaid comments shall not contain sensitive information.
- **DIAG-REQ-148:** Mermaid rendering shall use a documented tool or supported platform.
- **DIAG-REQ-149:** Mermaid version differences affecting layout or syntax shall be recorded.
- **DIAG-REQ-150:** Mermaid diagrams shall have text alternatives where visual interpretation is material.
- **DIAG-REQ-151:** Mermaid sequence diagrams shall identify actors explicitly.
- **DIAG-REQ-152:** Mermaid state diagrams shall identify invalid transitions in accompanying text or tables.
- **DIAG-REQ-153:** Mermaid flowcharts shall label decision branches.
- **DIAG-REQ-154:** Mermaid diagrams shall avoid excessive node density.
- **DIAG-REQ-155:** Large Mermaid diagrams shall be decomposed into views.
- **DIAG-REQ-156:** Mermaid source embedded in Markdown shall remain the authoritative source only when the file is explicitly identified.
- **DIAG-REQ-157:** Generated Mermaid source shall receive human review.
- **DIAG-REQ-158:** Mermaid security directives shall not relax platform protections.
- **DIAG-REQ-159:** Mermaid external links shall be reviewed.
- **DIAG-REQ-160:** Mermaid validation shall be included in documentation checks.

## 8. C4 Model Requirements

- **DIAG-REQ-161:** C4 diagrams shall identify the C4 level.
- **DIAG-REQ-162:** System Context diagrams shall identify persons and external systems.
- **DIAG-REQ-163:** Container diagrams shall identify major deployable or runtime boundaries.
- **DIAG-REQ-164:** Component diagrams shall identify material internal responsibilities.
- **DIAG-REQ-165:** Code diagrams shall be used sparingly and remain synchronized with implementation.
- **DIAG-REQ-166:** C4 elements shall identify names and responsibilities.
- **DIAG-REQ-167:** C4 relationships shall identify direction and purpose.
- **DIAG-REQ-168:** C4 relationships shall identify protocols or mechanisms where material.
- **DIAG-REQ-169:** C4 boundaries shall identify ownership and trust implications.
- **DIAG-REQ-170:** C4 persons shall distinguish human officer roles from other human actors.
- **DIAG-REQ-171:** C4 software systems shall distinguish OBDIA from external institutions and services.
- **DIAG-REQ-172:** C4 containers shall not imply deployment where only proposed design exists.
- **DIAG-REQ-173:** C4 components shall not imply implementation where only proposed design exists.
- **DIAG-REQ-174:** C4 diagrams shall identify proposed, implemented or validated status where mixed states exist.
- **DIAG-REQ-175:** C4 diagrams shall not use deployment terms for conceptual-only elements without qualification.
- **DIAG-REQ-176:** C4 scope shall be consistent across related levels.
- **DIAG-REQ-177:** C4 element identifiers shall remain stable across levels where they refer to the same element.
- **DIAG-REQ-178:** C4 abstractions shall not conceal security enforcement points.
- **DIAG-REQ-179:** C4 abstractions shall not conceal evidence or authorization boundaries where material.
- **DIAG-REQ-180:** C4 diagrams shall reference related deployment or data-flow views where needed.
- **DIAG-REQ-181:** C4 diagrams shall identify the accountable institution.
- **DIAG-REQ-182:** C4 diagrams shall identify agent, workload and connector boundaries.
- **DIAG-REQ-183:** C4 diagrams shall identify model or provider boundaries where material.
- **DIAG-REQ-184:** C4 diagrams shall identify storage and evidence systems where material.
- **DIAG-REQ-185:** C4 reviews shall verify consistency between levels.

## 9. Data Flow Diagram Requirements

- **DIAG-REQ-186:** Data Flow Diagrams shall identify external entities.
- **DIAG-REQ-187:** Data Flow Diagrams shall identify processes.
- **DIAG-REQ-188:** Data Flow Diagrams shall identify data stores.
- **DIAG-REQ-189:** Data Flow Diagrams shall identify data flows.
- **DIAG-REQ-190:** Data Flow Diagrams shall identify trust boundaries.
- **DIAG-REQ-191:** Data Flow Diagrams shall label data categories.
- **DIAG-REQ-192:** Data Flow Diagrams shall label flow direction.
- **DIAG-REQ-193:** Data Flow Diagrams shall identify authorization at protected crossings.
- **DIAG-REQ-194:** Data Flow Diagrams shall identify encryption where relevant without implying authorization.
- **DIAG-REQ-195:** Data Flow Diagrams shall identify data classification where material.
- **DIAG-REQ-196:** Data Flow Diagrams shall identify personal-data flows where material.
- **DIAG-REQ-197:** Data Flow Diagrams shall identify evidence and derived-data flows separately.
- **DIAG-REQ-198:** Data Flow Diagrams shall identify secret and credential flows without revealing values.
- **DIAG-REQ-199:** Data Flow Diagrams shall identify logs and audit flows.
- **DIAG-REQ-200:** Data Flow Diagrams shall identify retention or deletion destinations where material.
- **DIAG-REQ-201:** Data Flow Diagrams shall identify third-party processing.
- **DIAG-REQ-202:** Data Flow Diagrams shall identify cross-case or cross-purpose boundaries.
- **DIAG-REQ-203:** Data Flow Diagrams shall identify jurisdiction or region boundaries where material.
- **DIAG-REQ-204:** Data Flow Diagrams shall identify model and provider data flows.
- **DIAG-REQ-205:** Data Flow Diagrams shall identify connector ingress and egress.
- **DIAG-REQ-206:** Data Flow Diagrams shall identify rejected or quarantine paths where material.
- **DIAG-REQ-207:** Data Flow Diagrams shall avoid unlabeled bidirectional arrows.
- **DIAG-REQ-208:** Data Flow Diagrams shall identify scope and decomposition level.
- **DIAG-REQ-209:** Data Flow Diagrams shall trace to privacy and threat analysis.
- **DIAG-REQ-210:** Data Flow Diagram reviews shall verify that every material flow has a source, destination and purpose.

## 10. Sequence Diagram Requirements

- **DIAG-REQ-211:** Sequence Diagrams shall identify participants.
- **DIAG-REQ-212:** Participants shall distinguish human, agent, workload, connector, policy service and external system.
- **DIAG-REQ-213:** Sequence Diagrams shall identify message direction.
- **DIAG-REQ-214:** Sequence Diagrams shall label protected operations.
- **DIAG-REQ-215:** Sequence Diagrams shall identify authentication where material.
- **DIAG-REQ-216:** Sequence Diagrams shall identify authorization decisions where material.
- **DIAG-REQ-217:** Sequence Diagrams shall identify human approvals where material.
- **DIAG-REQ-218:** Sequence Diagrams shall identify obligations where material.
- **DIAG-REQ-219:** Sequence Diagrams shall identify timeouts and failures where material.
- **DIAG-REQ-220:** Sequence Diagrams shall identify retries where material.
- **DIAG-REQ-221:** Sequence Diagrams shall identify revocation or cancellation paths where material.
- **DIAG-REQ-222:** Sequence Diagrams shall identify evidence and audit events where material.
- **DIAG-REQ-223:** Sequence Diagrams shall distinguish synchronous and asynchronous interactions.
- **DIAG-REQ-224:** Sequence Diagrams shall distinguish request, acknowledgement and completion.
- **DIAG-REQ-225:** Sequence Diagrams shall identify alternative and error branches.
- **DIAG-REQ-226:** Sequence Diagrams shall identify loop bounds where relevant.
- **DIAG-REQ-227:** Sequence Diagrams shall not imply successful execution when only the happy path is shown.
- **DIAG-REQ-228:** Happy-path-only diagrams shall be labeled as such.
- **DIAG-REQ-229:** Sequence Diagrams shall identify the point of policy enforcement.
- **DIAG-REQ-230:** Sequence Diagrams shall identify connector mediation.
- **DIAG-REQ-231:** Sequence Diagrams shall identify source evidence versus derived output.
- **DIAG-REQ-232:** Sequence Diagrams shall identify lifecycle state prerequisites where material.
- **DIAG-REQ-233:** Sequence Diagrams shall identify long-running re-evaluation where material.
- **DIAG-REQ-234:** Sequence Diagrams shall preserve ordering assumptions.
- **DIAG-REQ-235:** Sequence Diagram reviews shall verify consistency with authorization and lifecycle policies.

## 11. Trust Boundary Diagram Requirements

- **DIAG-REQ-236:** Trust Boundary Diagrams shall identify every material trust domain.
- **DIAG-REQ-237:** Trust boundaries shall have stable identifiers.
- **DIAG-REQ-238:** Trust boundaries shall identify owner or authority.
- **DIAG-REQ-239:** Trust boundaries shall identify entry and exit points.
- **DIAG-REQ-240:** Trust boundaries shall identify authentication requirements.
- **DIAG-REQ-241:** Trust boundaries shall identify authorization requirements.
- **DIAG-REQ-242:** Trust boundaries shall identify data and evidence crossing.
- **DIAG-REQ-243:** Trust boundaries shall identify protocols where material.
- **DIAG-REQ-244:** Trust boundaries shall identify machine identities where material.
- **DIAG-REQ-245:** Trust boundaries shall identify human approval points where material.
- **DIAG-REQ-246:** Trust boundaries shall identify monitoring and audit points.
- **DIAG-REQ-247:** Trust boundaries shall identify failure and containment paths.
- **DIAG-REQ-248:** Trust boundaries shall identify connector and external-service boundaries.
- **DIAG-REQ-249:** Trust boundaries shall identify model-provider boundaries.
- **DIAG-REQ-250:** Trust boundaries shall identify repository, CI and build boundaries where material.
- **DIAG-REQ-251:** Trust boundaries shall identify development, test, demonstration and operational separation.
- **DIAG-REQ-252:** Trust boundaries shall identify isolated Dark Web research boundaries when represented.
- **DIAG-REQ-253:** Trust boundaries shall not imply trust from network location.
- **DIAG-REQ-254:** Trust boundaries shall not imply authorization from encryption.
- **DIAG-REQ-255:** Trust boundaries shall not imply lawful authority from authentication.
- **DIAG-REQ-256:** Nested trust boundaries shall be explained.
- **DIAG-REQ-257:** Overlapping trust boundaries shall be explained.
- **DIAG-REQ-258:** Unknown or disputed boundaries shall remain explicit.
- **DIAG-REQ-259:** Boundary changes shall trigger threat-model review.
- **DIAG-REQ-260:** Trust Boundary Diagram reviews shall verify alignment with `OBDIA-TRUST-001` and `OBDIA-SEC-001`.

## 12. Attack Tree Requirements

- **DIAG-REQ-261:** Attack Trees shall identify the adversarial objective.
- **DIAG-REQ-262:** Attack Trees shall identify root goal.
- **DIAG-REQ-263:** Attack Trees shall identify intermediate goals.
- **DIAG-REQ-264:** Attack Trees shall identify leaf conditions or techniques.
- **DIAG-REQ-265:** Attack Trees shall identify AND and OR relationships explicitly.
- **DIAG-REQ-266:** Attack Trees shall identify assumptions.
- **DIAG-REQ-267:** Attack Trees shall identify scope.
- **DIAG-REQ-268:** Attack Trees shall identify affected assets.
- **DIAG-REQ-269:** Attack Trees shall identify applicable threat actors or capabilities where relevant.
- **DIAG-REQ-270:** Attack Trees shall identify existing mitigations separately from attack paths.
- **DIAG-REQ-271:** Attack Trees shall identify residual or unmitigated paths.
- **DIAG-REQ-272:** Attack Trees shall identify detection and containment opportunities.
- **DIAG-REQ-273:** Attack Trees shall identify evidence and audit targets where relevant.
- **DIAG-REQ-274:** Attack Trees shall identify authorization-bypass paths where relevant.
- **DIAG-REQ-275:** Attack Trees shall identify prompt-injection and connector paths where relevant.
- **DIAG-REQ-276:** Attack Trees shall identify supply-chain paths where relevant.
- **DIAG-REQ-277:** Attack Trees shall not provide unnecessary operational instructions for harmful activity.
- **DIAG-REQ-278:** Attack Trees shall remain defensive and research-oriented.
- **DIAG-REQ-279:** Attack Trees shall use synthetic or abstract examples for public artifacts.
- **DIAG-REQ-280:** Attack Trees shall not identify real criminal infrastructure.
- **DIAG-REQ-281:** Attack Trees shall not encourage exploitation.
- **DIAG-REQ-282:** Attack Trees shall trace to threat-model identifiers.
- **DIAG-REQ-283:** Attack Trees shall trace mitigations to security requirements.
- **DIAG-REQ-284:** Attack Trees shall identify uncertainty.
- **DIAG-REQ-285:** Attack Tree reviews shall verify ethical and publication boundaries.

## 13. State and Lifecycle Diagrams

- **DIAG-REQ-286:** State diagrams shall identify the governed state machine.
- **DIAG-REQ-287:** State diagrams shall identify every valid state.
- **DIAG-REQ-288:** State diagrams shall identify every valid transition.
- **DIAG-REQ-289:** State diagrams shall identify terminal states.
- **DIAG-REQ-290:** State diagrams shall identify guards where material.
- **DIAG-REQ-291:** State diagrams shall identify transition authority where material.
- **DIAG-REQ-292:** State diagrams shall identify failure transitions.
- **DIAG-REQ-293:** State diagrams shall identify rollback or recovery transitions where applicable.
- **DIAG-REQ-294:** State diagrams shall identify prohibited transitions in accompanying text or tables.
- **DIAG-REQ-295:** State diagrams shall not imply transitions absent from the normative model.
- **DIAG-REQ-296:** Repository lifecycle diagrams shall use the exact constitutional lifecycle.
- **DIAG-REQ-297:** Agent lifecycle diagrams shall include mandatory `Bind`.
- **DIAG-REQ-298:** Agent lifecycle diagrams shall represent revocation and archive as terminal.
- **DIAG-REQ-299:** Connector lifecycle diagrams shall represent revoked and retired terminality.
- **DIAG-REQ-300:** Evidence handling-state diagrams shall remain distinct from repository lifecycle.
- **DIAG-REQ-301:** Session and workload states shall remain distinct from agent lifecycle.
- **DIAG-REQ-302:** State diagrams shall identify concurrency assumptions where material.
- **DIAG-REQ-303:** State diagrams shall identify idempotency or duplicate-transition behavior where material.
- **DIAG-REQ-304:** State diagrams shall identify suspension and revocation propagation where material.
- **DIAG-REQ-305:** State diagrams shall identify state source of truth.
- **DIAG-REQ-306:** State diagrams shall identify stale-state handling where material.
- **DIAG-REQ-307:** State diagrams shall trace to the governing lifecycle document.
- **DIAG-REQ-308:** State changes in diagrams shall require reconciliation with normative text.
- **DIAG-REQ-309:** State diagrams shall not create a new normative state by visual convenience.
- **DIAG-REQ-310:** State Diagram reviews shall verify semantic equivalence with transition tables.

## 14. Authorization Diagrams

- **DIAG-REQ-311:** Authorization diagrams shall identify subject.
- **DIAG-REQ-312:** Authorization diagrams shall identify resource.
- **DIAG-REQ-313:** Authorization diagrams shall identify action.
- **DIAG-REQ-314:** Authorization diagrams shall identify context.
- **DIAG-REQ-315:** Authorization diagrams shall identify policy version.
- **DIAG-REQ-316:** Authorization diagrams shall identify PDP.
- **DIAG-REQ-317:** Authorization diagrams shall identify PEP.
- **DIAG-REQ-318:** Authorization diagrams shall identify PAP where represented.
- **DIAG-REQ-319:** Authorization diagrams shall identify PIP where represented.
- **DIAG-REQ-320:** Authorization diagrams shall identify human approval where required.
- **DIAG-REQ-321:** Authorization diagrams shall identify decision outcomes.
- **DIAG-REQ-322:** Authorization diagrams shall identify obligations.
- **DIAG-REQ-323:** Authorization diagrams shall identify default-deny behavior.
- **DIAG-REQ-324:** Authorization diagrams shall identify revocation and cache invalidation where material.
- **DIAG-REQ-325:** Authorization diagrams shall identify case, purpose, jurisdiction, time, tool and data scope where material.
- **DIAG-REQ-326:** Authorization diagrams shall distinguish authentication from authorization.
- **DIAG-REQ-327:** Authorization diagrams shall distinguish authorization from legal authority.
- **DIAG-REQ-328:** Authorization diagrams shall distinguish model recommendation from policy decision.
- **DIAG-REQ-329:** Authorization diagrams shall identify failure-closed paths.
- **DIAG-REQ-330:** Authorization diagrams shall identify audit-event generation.
- **DIAG-REQ-331:** Authorization diagrams shall identify long-running re-evaluation where material.
- **DIAG-REQ-332:** Authorization diagrams shall not show direct model-to-privileged-action paths.
- **DIAG-REQ-333:** Authorization diagrams shall trace to `OBDIA-AUTH-001`.
- **DIAG-REQ-334:** Authorization diagrams shall use controlled decision labels.
- **DIAG-REQ-335:** Authorization Diagram reviews shall verify that `Permit` is never inferred from missing context.

## 15. Evidence and Chain-of-Custody Diagrams

- **DIAG-REQ-336:** Evidence diagrams shall identify source evidence.
- **DIAG-REQ-337:** Evidence diagrams shall identify acquisition events.
- **DIAG-REQ-338:** Evidence diagrams shall identify collector.
- **DIAG-REQ-339:** Evidence diagrams shall identify integrity verification.
- **DIAG-REQ-340:** Evidence diagrams shall identify custody events.
- **DIAG-REQ-341:** Evidence diagrams shall identify storage locations or domains.
- **DIAG-REQ-342:** Evidence diagrams shall identify derived artifacts.
- **DIAG-REQ-343:** Evidence diagrams shall identify transformation events.
- **DIAG-REQ-344:** Evidence diagrams shall identify reviewer actions where material.
- **DIAG-REQ-345:** Evidence diagrams shall identify export packages.
- **DIAG-REQ-346:** Evidence diagrams shall identify retention and disposition where material.
- **DIAG-REQ-347:** Evidence diagrams shall identify quarantine paths.
- **DIAG-REQ-348:** Evidence diagrams shall distinguish evidence from analysis.
- **DIAG-REQ-349:** Evidence diagrams shall distinguish source-provided and collector-generated metadata.
- **DIAG-REQ-350:** Evidence diagrams shall identify provenance links.
- **DIAG-REQ-351:** Evidence diagrams shall identify authorization references where material.
- **DIAG-REQ-352:** Evidence diagrams shall identify case and purpose boundaries.
- **DIAG-REQ-353:** Evidence diagrams shall not display real evidence content in public artifacts.
- **DIAG-REQ-354:** Evidence diagrams shall not reveal real case identifiers.
- **DIAG-REQ-355:** Evidence diagrams shall not imply admissibility.
- **DIAG-REQ-356:** Evidence diagrams shall not imply authenticity from integrity alone.
- **DIAG-REQ-357:** Evidence diagrams shall identify confidence separately from integrity where represented.
- **DIAG-REQ-358:** Evidence diagrams shall trace to `OBDIA-EVID-001`.
- **DIAG-REQ-359:** Chain-of-custody diagrams shall preserve chronological direction.
- **DIAG-REQ-360:** Evidence Diagram reviews shall verify that original and derived artifacts remain distinct.

## 16. Connector and External-System Diagrams

- **DIAG-REQ-361:** Connector diagrams shall identify connector identifier.
- **DIAG-REQ-362:** Connector diagrams shall identify connector operational state where material.
- **DIAG-REQ-363:** Connector diagrams shall identify machine identity.
- **DIAG-REQ-364:** Connector diagrams shall identify credential boundary without revealing credentials.
- **DIAG-REQ-365:** Connector diagrams shall identify authorization enforcement.
- **DIAG-REQ-366:** Connector diagrams shall identify approved endpoints or endpoint classes.
- **DIAG-REQ-367:** Connector diagrams shall identify network egress controls.
- **DIAG-REQ-368:** Connector diagrams shall identify request validation.
- **DIAG-REQ-369:** Connector diagrams shall identify response validation.
- **DIAG-REQ-370:** Connector diagrams shall identify provenance and audit generation.
- **DIAG-REQ-371:** Connector diagrams shall identify rate, timeout, retry and idempotency controls where material.
- **DIAG-REQ-372:** Connector diagrams shall identify failure isolation.
- **DIAG-REQ-373:** Connector diagrams shall identify suspension and revocation.
- **DIAG-REQ-374:** Connector diagrams shall identify external-service ownership.
- **DIAG-REQ-375:** Connector diagrams shall identify data classifications crossing the boundary.
- **DIAG-REQ-376:** Connector diagrams shall identify prompt-injection boundary where content enters model context.
- **DIAG-REQ-377:** Connector diagrams shall not show human credentials flowing to connectors.
- **DIAG-REQ-378:** Connector diagrams shall not imply unrestricted external access.
- **DIAG-REQ-379:** Connector diagrams shall not depict uncontrolled criminal-infrastructure interaction as an approved path.
- **DIAG-REQ-380:** Dark Web connector diagrams shall show simulation, archive or operational isolation boundaries.
- **DIAG-REQ-381:** Blockchain connector diagrams shall identify testnet, mainnet or local network.
- **DIAG-REQ-382:** Communication connector diagrams shall identify outbound-message approval where material.
- **DIAG-REQ-383:** Connector diagrams shall trace to connector manifests.
- **DIAG-REQ-384:** Connector diagrams shall trace to `OBDIA-CONN-001`.
- **DIAG-REQ-385:** Connector Diagram reviews shall verify that fallbacks are no more permissive.

## 17. Deployment, Network and Environment Diagrams

- **DIAG-REQ-386:** Deployment diagrams shall identify environment.
- **DIAG-REQ-387:** Deployment diagrams shall identify hosting or execution boundaries.
- **DIAG-REQ-388:** Deployment diagrams shall identify development, test, demonstration and operational separation.
- **DIAG-REQ-389:** Deployment diagrams shall identify workload identities where material.
- **DIAG-REQ-390:** Deployment diagrams shall identify storage and secret-management boundaries.
- **DIAG-REQ-391:** Deployment diagrams shall identify monitoring and audit components.
- **DIAG-REQ-392:** Deployment diagrams shall identify model and provider boundaries where material.
- **DIAG-REQ-393:** Deployment diagrams shall identify external dependencies.
- **DIAG-REQ-394:** Network diagrams shall identify segments and zones.
- **DIAG-REQ-395:** Network diagrams shall identify ingress and egress paths.
- **DIAG-REQ-396:** Network diagrams shall identify administrative interfaces.
- **DIAG-REQ-397:** Network diagrams shall identify proxies and gateways where material.
- **DIAG-REQ-398:** Network diagrams shall identify encryption in transit where material.
- **DIAG-REQ-399:** Network diagrams shall not imply authorization from network placement.
- **DIAG-REQ-400:** Network diagrams shall identify default-deny or restricted paths where material.
- **DIAG-REQ-401:** Network diagrams shall identify quarantine and containment zones.
- **DIAG-REQ-402:** Network diagrams shall identify high-risk isolated research environments.
- **DIAG-REQ-403:** Network diagrams shall avoid real public IP addresses in public artifacts unless intentionally public and safe.
- **DIAG-REQ-404:** Deployment diagrams shall avoid real internal hostnames in public artifacts.
- **DIAG-REQ-405:** Deployment diagrams shall not expose secret paths or credential values.
- **DIAG-REQ-406:** Conceptual deployments shall be labeled as conceptual.
- **DIAG-REQ-407:** Implemented deployments shall identify evidence and version.
- **DIAG-REQ-408:** Production-like diagrams shall not be labeled production without evidence and authority.
- **DIAG-REQ-409:** Environment diagrams shall trace to implementation and test environment definitions.
- **DIAG-REQ-410:** Deployment and Network Diagram reviews shall verify least privilege and segmentation.

## 18. Security Labels and Visual Semantics

- **DIAG-REQ-411:** Security labels shall use controlled terminology.
- **DIAG-REQ-412:** Trust levels shall not be represented only by color.
- **DIAG-REQ-413:** Authorization state shall not be represented only by color.
- **DIAG-REQ-414:** Lifecycle state shall not be represented only by color.
- **DIAG-REQ-415:** Risk severity shall not be represented only by color.
- **DIAG-REQ-416:** Line styles shall have documented meaning.
- **DIAG-REQ-417:** Arrowheads shall have documented meaning where nonstandard.
- **DIAG-REQ-418:** Dashed lines shall not be used ambiguously.
- **DIAG-REQ-419:** Boundary shapes shall have documented meaning.
- **DIAG-REQ-420:** Icons shall have documented meaning.
- **DIAG-REQ-421:** Human actors shall use a consistent symbol.
- **DIAG-REQ-422:** AI agents shall use a distinct consistent symbol.
- **DIAG-REQ-423:** Workloads shall use a distinct consistent symbol.
- **DIAG-REQ-424:** Connectors shall use a distinct consistent symbol.
- **DIAG-REQ-425:** External systems shall use a distinct consistent symbol.
- **DIAG-REQ-426:** Evidence stores shall use a distinct consistent symbol.
- **DIAG-REQ-427:** Policy decision and enforcement points shall be visually distinguishable.
- **DIAG-REQ-428:** Proposed and implemented elements shall be visually distinguishable.
- **DIAG-REQ-429:** Implemented and validated elements shall be visually distinguishable.
- **DIAG-REQ-430:** Deprecated or superseded elements shall be visibly qualified.
- **DIAG-REQ-431:** Unknown or unresolved elements shall be visibly qualified.
- **DIAG-REQ-432:** Sensitive labels shall be minimized.
- **DIAG-REQ-433:** Legend entries shall define every non-obvious visual convention.
- **DIAG-REQ-434:** Visual semantics shall remain consistent across related diagrams.
- **DIAG-REQ-435:** Semantic changes to visual conventions shall be versioned.

## 19. Trust, Authority and Accountability Semantics

- **DIAG-REQ-436:** Diagrams shall identify the accountable human officer where officer-bound operation is in scope.
- **DIAG-REQ-437:** Diagrams shall identify the accountable institution.
- **DIAG-REQ-438:** Diagrams shall represent the AI agent as a delegated technical actor.
- **DIAG-REQ-439:** Diagrams shall not represent the AI agent as an independent legal authority.
- **DIAG-REQ-440:** Diagrams shall not depict autonomous policing authority.
- **DIAG-REQ-441:** Diagrams shall not depict autonomous legal decision authority.
- **DIAG-REQ-442:** Diagrams shall not depict autonomous coercive authority.
- **DIAG-REQ-443:** Diagrams shall not depict human credentials as agent-held assets.
- **DIAG-REQ-444:** Diagrams shall distinguish institutional issuance from human identity.
- **DIAG-REQ-445:** Diagrams shall distinguish delegation from credential sharing.
- **DIAG-REQ-446:** Diagrams shall distinguish authentication from authorization.
- **DIAG-REQ-447:** Diagrams shall distinguish authorization from lawful authority.
- **DIAG-REQ-448:** Diagrams shall distinguish trust evaluation from access permission.
- **DIAG-REQ-449:** Diagrams shall identify human approval points for consequential operations.
- **DIAG-REQ-450:** Diagrams shall identify separation-of-duties boundaries where material.
- **DIAG-REQ-451:** Diagrams shall identify risk owners where risk decisions are represented.
- **DIAG-REQ-452:** Diagrams shall identify that AI systems do not accept risk.
- **DIAG-REQ-453:** Diagrams shall identify role concentration where it affects interpretation.
- **DIAG-REQ-454:** Diagrams shall not imply independent external review where none occurred.
- **DIAG-REQ-455:** Diagrams shall identify revocation authority where material.
- **DIAG-REQ-456:** Diagrams shall identify audit attribution paths.
- **DIAG-REQ-457:** Diagrams shall identify who can suspend or contain components where material.
- **DIAG-REQ-458:** Diagrams shall identify policy authority where material.
- **DIAG-REQ-459:** Diagrams shall preserve one-agent-to-one-officer binding.
- **DIAG-REQ-460:** Trust and accountability semantics shall trace to documents 06, 08, 09, 17 and 19.

## 20. Scope, Abstraction and Decomposition

- **DIAG-REQ-461:** Every diagram shall define its system boundary.
- **DIAG-REQ-462:** Every diagram shall define its time or lifecycle scope where material.
- **DIAG-REQ-463:** Every diagram shall define its environment scope.
- **DIAG-REQ-464:** Every diagram shall define its implementation-status scope.
- **DIAG-REQ-465:** Every diagram shall define its data-classification scope where material.
- **DIAG-REQ-466:** Every diagram shall define its audience.
- **DIAG-REQ-467:** Every diagram shall state material exclusions.
- **DIAG-REQ-468:** Abstraction level shall match the stated purpose.
- **DIAG-REQ-469:** Conceptual and implementation details shall not be mixed without explicit labels.
- **DIAG-REQ-470:** Cross-level diagrams shall identify the level of each element.
- **DIAG-REQ-471:** Large diagrams shall be decomposed when reviewability suffers.
- **DIAG-REQ-472:** Decomposed diagrams shall preserve stable element identifiers.
- **DIAG-REQ-473:** Decomposed diagrams shall include navigation or cross-references.
- **DIAG-REQ-474:** Overview diagrams shall not conceal critical trust boundaries.
- **DIAG-REQ-475:** Detailed diagrams shall not contradict overview diagrams.
- **DIAG-REQ-476:** Multiple views shall have a consistency review.
- **DIAG-REQ-477:** Partial views shall be labeled as partial.
- **DIAG-REQ-478:** Happy-path views shall be labeled as happy-path views.
- **DIAG-REQ-479:** Threat-only views shall be labeled as threat views.
- **DIAG-REQ-480:** Future-state views shall be labeled as proposed or future research.
- **DIAG-REQ-481:** Simulation views shall be labeled as simulation.
- **DIAG-REQ-482:** Test-environment views shall be labeled as test environment.
- **DIAG-REQ-483:** Public views may redact sensitive details but shall not fabricate architecture.
- **DIAG-REQ-484:** Redaction shall preserve enough semantics for the stated purpose.
- **DIAG-REQ-485:** Scope changes shall trigger version and review assessment.

## 21. Legends, Notes and Annotations

- **DIAG-REQ-486:** Every diagram using non-obvious notation shall include a legend.
- **DIAG-REQ-487:** Legends shall define shapes.
- **DIAG-REQ-488:** Legends shall define line styles.
- **DIAG-REQ-489:** Legends shall define arrow meanings.
- **DIAG-REQ-490:** Legends shall define colors where used.
- **DIAG-REQ-491:** Legends shall define status markers.
- **DIAG-REQ-492:** Legends shall define trust markers.
- **DIAG-REQ-493:** Legends shall define data-classification markers.
- **DIAG-REQ-494:** Legends shall define proposed, implemented and validated markers.
- **DIAG-REQ-495:** Notes shall identify assumptions.
- **DIAG-REQ-496:** Notes shall identify unresolved questions.
- **DIAG-REQ-497:** Notes shall identify limitations.
- **DIAG-REQ-498:** Notes shall identify redactions.
- **DIAG-REQ-499:** Notes shall identify external dependencies.
- **DIAG-REQ-500:** Notes shall identify time or ordering caveats.
- **DIAG-REQ-501:** Notes shall identify whether the diagram is normative or informative.
- **DIAG-REQ-502:** Annotations shall not contain secrets.
- **DIAG-REQ-503:** Annotations shall not contain unnecessary personal or case data.
- **DIAG-REQ-504:** Annotations shall not claim compliance without evidence.
- **DIAG-REQ-505:** Annotations shall not claim production deployment without evidence.
- **DIAG-REQ-506:** Annotations shall not conceal failed or missing controls.
- **DIAG-REQ-507:** Callouts shall remain linked to the elements they describe.
- **DIAG-REQ-508:** Numbered callouts shall have corresponding explanations.
- **DIAG-REQ-509:** Legend and note text shall be readable at normal repository scale.
- **DIAG-REQ-510:** Legend semantics shall remain consistent across rendered formats.

## 22. Accessibility and Readability

- **DIAG-REQ-511:** Diagrams shall provide meaningful text alternatives.
- **DIAG-REQ-512:** Diagram titles shall be descriptive.
- **DIAG-REQ-513:** Diagram descriptions shall summarize purpose and scope.
- **DIAG-REQ-514:** Color shall not be the sole means of conveying meaning.
- **DIAG-REQ-515:** Text contrast shall be sufficient for ordinary review.
- **DIAG-REQ-516:** Text size shall remain readable in repository rendering.
- **DIAG-REQ-517:** Labels shall avoid excessive abbreviation.
- **DIAG-REQ-518:** Abbreviations shall be defined.
- **DIAG-REQ-519:** Arrow crossings shall be minimized.
- **DIAG-REQ-520:** Overlapping labels shall be avoided.
- **DIAG-REQ-521:** Layouts shall support a clear reading order.
- **DIAG-REQ-522:** Complex diagrams shall provide a walkthrough or companion text.
- **DIAG-REQ-523:** Interactive-only semantics shall have static alternatives.
- **DIAG-REQ-524:** Hover-only information shall not be required for understanding.
- **DIAG-REQ-525:** SVG text should remain selectable where practical.
- **DIAG-REQ-526:** Raster diagrams shall have sufficient resolution.
- **DIAG-REQ-527:** Screen-reader alternatives shall identify major nodes and flows.
- **DIAG-REQ-528:** Text alternatives shall describe trust and authorization boundaries.
- **DIAG-REQ-529:** Text alternatives shall describe important failure paths.
- **DIAG-REQ-530:** Text alternatives shall not expose redacted sensitive details.
- **DIAG-REQ-531:** Accessibility checks shall be included in diagram review.
- **DIAG-REQ-532:** Accessible alternatives shall remain synchronized with diagrams.
- **DIAG-REQ-533:** Diagram captions shall identify version and identifier.
- **DIAG-REQ-534:** Rendered diagrams shall avoid clipping and truncation.
- **DIAG-REQ-535:** Readability limitations shall remain visible.

## 23. Security and Privacy Redaction

- **DIAG-REQ-536:** Public diagrams shall not expose secrets.
- **DIAG-REQ-537:** Public diagrams shall not expose credentials.
- **DIAG-REQ-538:** Public diagrams shall not expose private keys or tokens.
- **DIAG-REQ-539:** Public diagrams shall not expose real officer identities unless explicitly authorized.
- **DIAG-REQ-540:** Public diagrams shall not expose real victim, suspect or case-subject identities.
- **DIAG-REQ-541:** Public diagrams shall not expose real case identifiers.
- **DIAG-REQ-542:** Public diagrams shall not expose internal administrative endpoints unnecessarily.
- **DIAG-REQ-543:** Public diagrams shall not expose real sensitive hostnames unnecessarily.
- **DIAG-REQ-544:** Public diagrams shall not expose confidential tenant or account identifiers.
- **DIAG-REQ-545:** Public diagrams shall not expose sensitive storage paths.
- **DIAG-REQ-546:** Public diagrams shall not expose hidden offensive capabilities.
- **DIAG-REQ-547:** Redaction shall use controlled placeholders.
- **DIAG-REQ-548:** Redaction shall not change architectural meaning materially.
- **DIAG-REQ-549:** Redaction shall not invent components or boundaries.
- **DIAG-REQ-550:** Redaction shall be documented.
- **DIAG-REQ-551:** Redacted and internal versions shall have distinct access controls.
- **DIAG-REQ-552:** Redacted diagrams shall remain traceable to authoritative internal sources where permitted.
- **DIAG-REQ-553:** Diagram metadata shall be reviewed for sensitive information.
- **DIAG-REQ-554:** Embedded links shall be reviewed.
- **DIAG-REQ-555:** Embedded images shall be reviewed.
- **DIAG-REQ-556:** Version-control history shall be considered during sensitive-data review.
- **DIAG-REQ-557:** Sensitive diagram incidents shall trigger containment and history review.
- **DIAG-REQ-558:** Privacy redaction shall minimize personal data.
- **DIAG-REQ-559:** Security redaction shall not be described as complete assurance.
- **DIAG-REQ-560:** Release review shall confirm public-diagram safety.

## 24. Diagram Validation

- **DIAG-REQ-561:** Diagram source shall pass syntax validation where tooling supports it.
- **DIAG-REQ-562:** Rendered output shall be visually inspected.
- **DIAG-REQ-563:** Diagram identifiers shall be validated.
- **DIAG-REQ-564:** Diagram version shall be validated.
- **DIAG-REQ-565:** Diagram status shall be validated.
- **DIAG-REQ-566:** Related-document references shall be validated.
- **DIAG-REQ-567:** Element identifiers shall be validated where used.
- **DIAG-REQ-568:** Terminology shall be validated against the glossary.
- **DIAG-REQ-569:** Trust-boundary semantics shall be validated.
- **DIAG-REQ-570:** Authorization semantics shall be validated.
- **DIAG-REQ-571:** Lifecycle semantics shall be validated.
- **DIAG-REQ-572:** Evidence semantics shall be validated.
- **DIAG-REQ-573:** Connector semantics shall be validated.
- **DIAG-REQ-574:** Source and rendered artifacts shall be compared for consistency.
- **DIAG-REQ-575:** Broken links shall be detected.
- **DIAG-REQ-576:** Missing assets shall be detected.
- **DIAG-REQ-577:** Clipped or unreadable output shall be rejected.
- **DIAG-REQ-578:** Stale rendered artifacts shall be detected.
- **DIAG-REQ-579:** Sensitive-data checks shall be performed.
- **DIAG-REQ-580:** Accessibility checks shall be performed.
- **DIAG-REQ-581:** Diagram validation shall identify the exact source commit.
- **DIAG-REQ-582:** Diagram validation shall identify tool versions where material.
- **DIAG-REQ-583:** Validation failures shall remain visible.
- **DIAG-REQ-584:** Passing syntax validation shall not establish semantic correctness.
- **DIAG-REQ-585:** Diagram validation evidence shall trace to test records where automated.

## 25. Diagram Review

- **DIAG-REQ-586:** Material diagrams shall receive human review.
- **DIAG-REQ-587:** Review shall identify the exact diagram source commit.
- **DIAG-REQ-588:** Review shall identify diagram identifier and version.
- **DIAG-REQ-589:** Review shall verify scope.
- **DIAG-REQ-590:** Review shall verify related documents.
- **DIAG-REQ-591:** Review shall verify notation.
- **DIAG-REQ-592:** Review shall verify legend and annotations.
- **DIAG-REQ-593:** Review shall verify terminology.
- **DIAG-REQ-594:** Review shall verify architecture consistency.
- **DIAG-REQ-595:** Review shall verify trust and authorization semantics.
- **DIAG-REQ-596:** Review shall verify evidence and lifecycle semantics where applicable.
- **DIAG-REQ-597:** Review shall verify implementation-status labels.
- **DIAG-REQ-598:** Review shall verify security redaction.
- **DIAG-REQ-599:** Review shall verify privacy redaction.
- **DIAG-REQ-600:** Review shall verify accessibility.
- **DIAG-REQ-601:** Review shall verify source/render consistency.
- **DIAG-REQ-602:** Security-sensitive diagrams shall receive Security Reviewer assessment.
- **DIAG-REQ-603:** Privacy-sensitive diagrams shall receive Privacy and Governance Reviewer assessment.
- **DIAG-REQ-604:** Implementation diagrams shall receive Implementation Reviewer assessment where implementation exists.
- **DIAG-REQ-605:** Threat diagrams shall receive Security Reviewer assessment.
- **DIAG-REQ-606:** Release diagrams shall receive Release Reviewer assessment.
- **DIAG-REQ-607:** Role concentration shall be disclosed.
- **DIAG-REQ-608:** Internal review shall not be represented as independent external assurance.
- **DIAG-REQ-609:** Material changes after review shall invalidate affected review evidence.
- **DIAG-REQ-610:** Critical review findings shall block approval or publication.

## 26. Diagram Status and Lifecycle

- **DIAG-REQ-611:** Diagram lifecycle status shall follow `OBDIA-STATE-001`.
- **DIAG-REQ-612:** Diagram lifecycle status shall remain distinct from the state depicted.
- **DIAG-REQ-613:** A Draft diagram may depict an Approved normative model only when the relationship is explicit.
- **DIAG-REQ-614:** An Approved diagram shall not imply that depicted implementation exists.
- **DIAG-REQ-615:** An Implemented diagram shall identify implementation evidence.
- **DIAG-REQ-616:** A Validated diagram shall identify validation evidence for the representation and scope.
- **DIAG-REQ-617:** A Published diagram shall identify the publication decision.
- **DIAG-REQ-618:** A Superseded diagram shall identify its successor.
- **DIAG-REQ-619:** An Archived diagram shall remain read-only.
- **DIAG-REQ-620:** Diagram status shall be declared in metadata.
- **DIAG-REQ-621:** Diagram status shall not be inferred from repository path.
- **DIAG-REQ-622:** Diagram status shall not be inferred from merge state.
- **DIAG-REQ-623:** Diagram status shall not be inferred from rendered format.
- **DIAG-REQ-624:** Status changes shall be attributable.
- **DIAG-REQ-625:** Status changes shall have evidence.
- **DIAG-REQ-626:** Invalid status transitions shall be rejected.
- **DIAG-REQ-627:** Diagram withdrawal shall preserve history.
- **DIAG-REQ-628:** Diagram correction shall preserve prior records.
- **DIAG-REQ-629:** Diagram publication shall follow `OBDIA-REL-001`.
- **DIAG-REQ-630:** Diagram archival shall preserve source and rendered artifacts.
- **DIAG-REQ-631:** Restoration from archive shall not reactivate status silently.
- **DIAG-REQ-632:** Status badges shall match authoritative metadata.
- **DIAG-REQ-633:** Mixed-status elements shall be labeled individually.
- **DIAG-REQ-634:** Status uncertainty shall remain explicit.
- **DIAG-REQ-635:** Lifecycle tests or validation shall verify status representation where automated.

## 27. Versioning and Compatibility

- **DIAG-REQ-636:** Every authoritative diagram shall have a version.
- **DIAG-REQ-637:** Diagram versions shall follow document 25 after consolidation.
- **DIAG-REQ-638:** Editorial layout-only changes shall be distinguishable from semantic changes.
- **DIAG-REQ-639:** Semantic changes shall identify affected elements and relationships.
- **DIAG-REQ-640:** Breaking semantic changes shall receive impact analysis.
- **DIAG-REQ-641:** Diagram versions shall not be reused.
- **DIAG-REQ-642:** Rendered artifacts shall identify the source diagram version.
- **DIAG-REQ-643:** Related diagrams shall identify compatibility expectations where material.
- **DIAG-REQ-644:** C4 levels describing the same system shall remain version-compatible.
- **DIAG-REQ-645:** Sequence diagrams shall remain compatible with authorization and lifecycle versions.
- **DIAG-REQ-646:** Data-flow diagrams shall remain compatible with data and privacy models.
- **DIAG-REQ-647:** Trust-boundary diagrams shall remain compatible with trust and security models.
- **DIAG-REQ-648:** Attack trees shall identify the threat-model version.
- **DIAG-REQ-649:** State diagrams shall identify the lifecycle-model version.
- **DIAG-REQ-650:** Diagram version changes shall update cross-references where needed.
- **DIAG-REQ-651:** Version changes shall preserve prior versions according to repository policy.
- **DIAG-REQ-652:** Silent replacement of published diagrams shall be prohibited.
- **DIAG-REQ-653:** Correction releases shall use governed versioning.
- **DIAG-REQ-654:** Diagram aliases shall not conceal version changes.
- **DIAG-REQ-655:** Unknown version shall block authoritative use.
- **DIAG-REQ-656:** Version compatibility limitations shall be documented.
- **DIAG-REQ-657:** Forward dependency on document 25 shall remain explicit.
- **DIAG-REQ-658:** Version metadata shall be machine-readable where practical.
- **DIAG-REQ-659:** Version validation shall be part of diagram checks.
- **DIAG-REQ-660:** Version history shall identify rationale.

## 28. Repository Placement and Indexing

- **DIAG-REQ-661:** Diagram source files shall reside in governed repository locations.
- **DIAG-REQ-662:** Rendered diagrams shall reside in governed repository locations.
- **DIAG-REQ-663:** Source and rendered directories shall be distinguishable where practical.
- **DIAG-REQ-664:** Diagram indexes shall list identifier, title, version, status and paths.
- **DIAG-REQ-665:** Diagram indexes shall identify related documents.
- **DIAG-REQ-666:** Diagram indexes shall identify superseded and archived diagrams.
- **DIAG-REQ-667:** Diagram indexes shall not list temporary artifacts as authoritative.
- **DIAG-REQ-668:** Documentation shall link to diagrams by governed paths.
- **DIAG-REQ-669:** Cross-references shall prefer stable relative repository links.
- **DIAG-REQ-670:** Broken diagram links shall block affected release.
- **DIAG-REQ-671:** Repository moves shall preserve history.
- **DIAG-REQ-672:** Repository moves shall update indexes and links.
- **DIAG-REQ-673:** Generated diagrams shall not be stored in uncontrolled locations.
- **DIAG-REQ-674:** Large binary artifacts shall be governed.
- **DIAG-REQ-675:** External diagram-hosting shall not be the sole authoritative source.
- **DIAG-REQ-676:** Private diagrams shall not be linked from public artifacts without access consideration.
- **DIAG-REQ-677:** Repository placement shall support code ownership and review routing.
- **DIAG-REQ-678:** Repository placement shall support automated validation.
- **DIAG-REQ-679:** Repository placement shall support archival.
- **DIAG-REQ-680:** Repository placement shall support release packaging.
- **DIAG-REQ-681:** Document 24 shall govern final repository paths after consolidation.
- **DIAG-REQ-682:** Temporary placement before document 24 shall be explicitly transitional.
- **DIAG-REQ-683:** Path changes shall not alter diagram identifiers.
- **DIAG-REQ-684:** Repository indexes shall be validated.
- **DIAG-REQ-685:** Duplicate authoritative sources shall be prohibited.

## 29. Automation and CI Validation

- **DIAG-REQ-686:** Diagram automation shall operate on exact source commits.
- **DIAG-REQ-687:** Diagram automation configuration shall be versioned.
- **DIAG-REQ-688:** Automation shall validate syntax where supported.
- **DIAG-REQ-689:** Automation shall validate required metadata.
- **DIAG-REQ-690:** Automation shall validate identifiers.
- **DIAG-REQ-691:** Automation shall validate filenames.
- **DIAG-REQ-692:** Automation shall validate cross-references.
- **DIAG-REQ-693:** Automation shall render supported source formats.
- **DIAG-REQ-694:** Automation shall detect stale rendered output where feasible.
- **DIAG-REQ-695:** Automation shall detect missing rendered output where required.
- **DIAG-REQ-696:** Automation shall detect source/render mismatch where feasible.
- **DIAG-REQ-697:** Automation shall detect broken links.
- **DIAG-REQ-698:** Automation shall perform secret scanning.
- **DIAG-REQ-699:** Automation shall not expose sensitive diagrams to external services without authorization.
- **DIAG-REQ-700:** Automation credentials shall use least privilege.
- **DIAG-REQ-701:** Automation logs shall not expose secrets.
- **DIAG-REQ-702:** Automation-tool failure shall not count as success.
- **DIAG-REQ-703:** Automation shall preserve validation evidence.
- **DIAG-REQ-704:** Automation shall identify tool versions.
- **DIAG-REQ-705:** Automation shall enforce timeouts.
- **DIAG-REQ-706:** Automation shall not approve diagrams autonomously.
- **DIAG-REQ-707:** AI-generated diagram summaries shall not replace source review.
- **DIAG-REQ-708:** Required checks shall reflect risk-based diagram controls.
- **DIAG-REQ-709:** CI reruns shall preserve prior failures.
- **DIAG-REQ-710:** Automation limitations shall remain visible.

## 30. Diagram Testing

- **DIAG-REQ-711:** Diagram tests shall trace to diagram requirements.
- **DIAG-REQ-712:** Syntax tests shall verify parsability.
- **DIAG-REQ-713:** Metadata tests shall verify required fields.
- **DIAG-REQ-714:** Filename tests shall verify naming rules.
- **DIAG-REQ-715:** Identifier tests shall verify uniqueness.
- **DIAG-REQ-716:** Cross-reference tests shall verify related documents.
- **DIAG-REQ-717:** Render tests shall verify successful rendering.
- **DIAG-REQ-718:** Render tests shall verify expected output paths.
- **DIAG-REQ-719:** Visual review shall verify readability.
- **DIAG-REQ-720:** Accessibility tests shall verify text alternatives and non-color semantics.
- **DIAG-REQ-721:** Security tests shall verify absence of secrets.
- **DIAG-REQ-722:** Privacy tests shall verify absence of unauthorized personal or case data.
- **DIAG-REQ-723:** Semantic tests shall verify lifecycle state names where machine-checkable.
- **DIAG-REQ-724:** Semantic tests shall verify controlled authorization outcomes where machine-checkable.
- **DIAG-REQ-725:** Semantic tests shall verify one-agent-to-one-officer representation where applicable.
- **DIAG-REQ-726:** Semantic tests shall verify source and derived evidence distinction where applicable.
- **DIAG-REQ-727:** Regression tests shall protect corrected diagram defects.
- **DIAG-REQ-728:** Repository tests shall verify index completeness.
- **DIAG-REQ-729:** Release tests shall verify source/render integrity.
- **DIAG-REQ-730:** Tests shall identify exact diagram and tool versions.
- **DIAG-REQ-731:** Failed tests shall remain visible.
- **DIAG-REQ-732:** Passing render tests shall not establish architectural correctness.
- **DIAG-REQ-733:** Manual semantic review shall remain required for material diagrams.
- **DIAG-REQ-734:** Diagram tests shall conform to `OBDIA-TEST-001`.
- **DIAG-REQ-735:** Test evidence shall be retained.

## 31. Change Impact and Synchronization

- **DIAG-REQ-736:** Architecture changes shall identify affected diagrams.
- **DIAG-REQ-737:** ADR changes shall identify affected diagrams.
- **DIAG-REQ-738:** Requirement changes shall identify affected diagrams.
- **DIAG-REQ-739:** Terminology changes shall identify affected diagrams.
- **DIAG-REQ-740:** Trust-model changes shall identify affected diagrams.
- **DIAG-REQ-741:** Authorization changes shall identify affected diagrams.
- **DIAG-REQ-742:** Evidence-model changes shall identify affected diagrams.
- **DIAG-REQ-743:** Agent-lifecycle changes shall identify affected diagrams.
- **DIAG-REQ-744:** Connector-policy changes shall identify affected diagrams.
- **DIAG-REQ-745:** Implementation changes shall identify affected diagrams.
- **DIAG-REQ-746:** Test-environment changes shall identify affected diagrams.
- **DIAG-REQ-747:** Repository-structure changes shall identify affected diagram paths.
- **DIAG-REQ-748:** Versioning-policy changes shall identify affected diagram metadata.
- **DIAG-REQ-749:** Risk and exception changes shall identify affected threat diagrams.
- **DIAG-REQ-750:** Release changes shall identify affected public diagrams.
- **DIAG-REQ-751:** Affected diagrams shall be updated or explicitly marked stale.
- **DIAG-REQ-752:** Stale diagrams shall not be represented as current.
- **DIAG-REQ-753:** Synchronization status shall be visible.
- **DIAG-REQ-754:** Diagram updates shall not silently alter normative text.
- **DIAG-REQ-755:** Normative-text updates shall not rely solely on diagram changes.
- **DIAG-REQ-756:** Conflicts between diagram and normative text shall resolve in favor of higher authority.
- **DIAG-REQ-757:** Conflicts shall be recorded and corrected.
- **DIAG-REQ-758:** Material synchronization gaps shall block approval or release.
- **DIAG-REQ-759:** Change impact shall be bidirectionally traceable.
- **DIAG-REQ-760:** Document 29 remains a forward dependency for complete change governance.

## 32. Public and Portfolio Diagrams

- **DIAG-REQ-761:** Public diagrams shall comply with documents 04 and 15.
- **DIAG-REQ-762:** Public diagrams shall state independent research status where material.
- **DIAG-REQ-763:** Public diagrams shall not imply official institutional adoption.
- **DIAG-REQ-764:** Public diagrams shall not imply real-investigation deployment.
- **DIAG-REQ-765:** Public diagrams shall not imply production readiness without evidence.
- **DIAG-REQ-766:** Public diagrams shall not imply certification or guaranteed compliance.
- **DIAG-REQ-767:** Public diagrams shall distinguish proposed architecture from implementation.
- **DIAG-REQ-768:** Public diagrams shall distinguish implementation from testing.
- **DIAG-REQ-769:** Public diagrams shall distinguish testing from validation.
- **DIAG-REQ-770:** Public diagrams shall use synthetic names and identifiers.
- **DIAG-REQ-771:** Public diagrams shall avoid real officer and case data.
- **DIAG-REQ-772:** Public diagrams shall avoid internal sensitive endpoints.
- **DIAG-REQ-773:** Public diagrams shall not expose credentials or secrets.
- **DIAG-REQ-774:** Public threat diagrams shall remain defensive.
- **DIAG-REQ-775:** Public Dark Web diagrams shall use simulation or authorized archival framing.
- **DIAG-REQ-776:** Public blockchain diagrams shall identify testnet or read-only scope where applicable.
- **DIAG-REQ-777:** Public screenshots shall match current diagram status.
- **DIAG-REQ-778:** Public renders shall trace to reviewed source.
- **DIAG-REQ-779:** Public diagram captions shall identify version and status.
- **DIAG-REQ-780:** Public diagram limitations shall remain visible.
- **DIAG-REQ-781:** Public diagram packages shall include accessibility alternatives.
- **DIAG-REQ-782:** Public diagram corrections shall follow release policy.
- **DIAG-REQ-783:** Withdrawn public diagrams shall retain governed withdrawal records.
- **DIAG-REQ-784:** Public-diagram release shall receive Release Reviewer assessment.
- **DIAG-REQ-785:** Portfolio use shall not exceed supporting evidence.

## 33. Diagram Traceability

- **DIAG-REQ-786:** Every diagram shall trace to governing documents.
- **DIAG-REQ-787:** Every diagram shall trace to its scope.
- **DIAG-REQ-788:** Every material element shall trace to an architecture element, requirement or documented rationale.
- **DIAG-REQ-789:** Every trust boundary shall trace to trust or security analysis.
- **DIAG-REQ-790:** Every authorization flow shall trace to authorization requirements.
- **DIAG-REQ-791:** Every evidence flow shall trace to evidence requirements.
- **DIAG-REQ-792:** Every lifecycle state shall trace to the governing lifecycle model.
- **DIAG-REQ-793:** Every connector boundary shall trace to a connector manifest or policy.
- **DIAG-REQ-794:** Every attack-tree node shall trace to a threat identifier where applicable.
- **DIAG-REQ-795:** Every mitigation shown shall trace to a control or requirement.
- **DIAG-REQ-796:** Every implemented element shall trace to implementation evidence.
- **DIAG-REQ-797:** Every validated element shall trace to validation evidence.
- **DIAG-REQ-798:** Every published diagram shall trace to release approval.
- **DIAG-REQ-799:** Every superseded diagram shall trace to its successor.
- **DIAG-REQ-800:** Every rendered artifact shall trace to source.
- **DIAG-REQ-801:** Every diagram review shall identify exact source commit.
- **DIAG-REQ-802:** Every diagram test shall trace to requirements.
- **DIAG-REQ-803:** Traceability shall be bidirectional.
- **DIAG-REQ-804:** Broken traceability affecting authority, security, evidence, lifecycle or release shall be blocking.
- **DIAG-REQ-805:** Planned elements shall not be represented as implemented.
- **DIAG-REQ-806:** Implemented elements shall not be represented as validated without evidence.
- **DIAG-REQ-807:** Unknown relationships shall remain explicit.
- **DIAG-REQ-808:** Traceability records shall use immutable identifiers.
- **DIAG-REQ-809:** Traceability records shall not contain secrets.
- **DIAG-REQ-810:** Baseline freeze shall validate diagram traceability across documents 01–30.

## 34. Diagram Review Gate

- **DIAG-REQ-811:** Diagram-standard review shall identify the exact document and commit.
- **DIAG-REQ-812:** Review shall verify preservation of all six preferred formats.
- **DIAG-REQ-813:** Review shall verify preservation of version, scope and related-document metadata.
- **DIAG-REQ-814:** Review shall verify identifier and filename rules.
- **DIAG-REQ-815:** Review shall verify source and rendered-artifact rules.
- **DIAG-REQ-816:** Review shall verify notation-specific requirements.
- **DIAG-REQ-817:** Review shall verify trust and authorization semantics.
- **DIAG-REQ-818:** Review shall verify evidence and lifecycle semantics.
- **DIAG-REQ-819:** Review shall verify connector and deployment semantics.
- **DIAG-REQ-820:** Review shall verify legends and annotations.
- **DIAG-REQ-821:** Review shall verify accessibility.
- **DIAG-REQ-822:** Review shall verify security and privacy redaction.
- **DIAG-REQ-823:** Review shall verify validation and testing.
- **DIAG-REQ-824:** Review shall verify versioning and repository dependencies.
- **DIAG-REQ-825:** Review shall verify public and portfolio boundaries.
- **DIAG-REQ-826:** Review shall verify traceability.
- **DIAG-REQ-827:** Review shall identify residual risks and forward dependencies.
- **DIAG-REQ-828:** Critical findings shall block progression.
- **DIAG-REQ-829:** Material post-review changes shall invalidate affected evidence.
- **DIAG-REQ-830:** Role concentration shall be disclosed.
- **DIAG-REQ-831:** Internal review shall not be represented as independent external assurance.
- **DIAG-REQ-832:** Documentation Authority shall assess documentation consistency.
- **DIAG-REQ-833:** Security Reviewer shall assess trust, attack and sensitive architecture views.
- **DIAG-REQ-834:** Privacy and Governance Reviewer shall assess data and personal-information exposure.
- **DIAG-REQ-835:** Implementation Reviewer shall assess implementation diagrams where implementation exists.
- **DIAG-REQ-836:** Release Reviewer shall assess public diagrams.
- **DIAG-REQ-837:** Project Founder approval shall not substitute for required specialist review.
- **DIAG-REQ-838:** Automated validation shall remain advisory for semantic correctness.
- **DIAG-REQ-839:** Review non-applicability shall have rationale.
- **DIAG-REQ-840:** Approval for Draft incorporation shall not validate any depicted architecture.

## 35. Minimum Validation Checklist

Before approval or publication of a material diagram, confirm:

- [ ] immutable diagram identifier exists;
- [ ] version is declared;
- [ ] status is declared;
- [ ] scope is declared;
- [ ] related documents are declared;
- [ ] owner and approval authority are declared;
- [ ] notation and legend are declared;
- [ ] authoritative source is identified;
- [ ] rendered artifacts trace to source;
- [ ] filename follows governed naming;
- [ ] terminology matches the glossary;
- [ ] human, agent, workload, connector and external actors are distinguishable;
- [ ] trust boundaries are explicit where material;
- [ ] authentication and authorization are distinct;
- [ ] authorization and legal authority are distinct;
- [ ] evidence and derived analysis are distinct;
- [ ] lifecycle states match governing models;
- [ ] proposed, implemented and validated elements are distinguished;
- [ ] failure, denial, suspension or revocation paths are shown where material;
- [ ] secrets and sensitive operational details are absent from public artifacts;
- [ ] personal and case data are minimized;
- [ ] text alternatives and non-color semantics are provided;
- [ ] source syntax and rendering are validated;
- [ ] source and rendered output are consistent;
- [ ] review evidence identifies the exact source commit;
- [ ] traceability is bidirectional;
- [ ] forward dependencies 24, 25, 29 and 30 are recorded;
- [ ] Project Founder approval exists before status becomes Approved.


## 36. Limitations

- This standard does not select a diagramming product, rendering service, icon library, visual theme or CI platform.
- It does not make a diagram operationally accurate without review.
- It does not prove architecture implementation, control effectiveness, production readiness, certification or compliance.
- A complete diagram can still omit unknown dependencies or threats.
- Visual simplicity can obscure technical nuance and shall be balanced with companion text.
- Rendered similarity does not prove semantic equivalence.
- Redaction can reduce reviewability and shall be documented.
- Automated syntax and rendering checks cannot establish architectural correctness.
- Internal review is not independent certification.
- Documents 24, 25, 29 and 30 remain forward dependencies for repository placement, versioning, change and compliance governance.


## 37. Change Control

Every material change shall identify rationale, affected notations, diagrams, identifiers, filenames, source and rendered artifacts, security and privacy impact, accessibility impact, migration, validation, rollback, authority and version effect.

- **DIAG-REQ-841:** Editorial corrections shall use a patch version when meaning is unchanged.
- **DIAG-REQ-842:** Backward-compatible substantive additions shall use a minor version.
- **DIAG-REQ-843:** Incompatible diagram semantics shall use a major version.
- **DIAG-REQ-844:** Material architecture-representation changes shall require an ADR where applicable.
- **DIAG-REQ-845:** Changes shall require Project Founder approval.
- **DIAG-REQ-846:** Changes shall receive Documentation Authority assessment.
- **DIAG-REQ-847:** Security, privacy, implementation and release impacts shall receive specialist review where applicable.
- **DIAG-REQ-848:** Changes shall identify affected diagrams and cross-references.
- **DIAG-REQ-849:** Changes shall include source and rendering migration analysis.
- **DIAG-REQ-850:** Changes shall include accessibility validation.
- **DIAG-REQ-851:** Changes shall include rollback analysis.
- **DIAG-REQ-852:** Changes shall not retroactively fabricate diagram review or validation evidence.
- **DIAG-REQ-853:** Historical diagram versions shall not be silently rewritten.
- **DIAG-REQ-854:** Diagram identifiers shall not be reused.
- **DIAG-REQ-855:** Forward-dependency reconciliation shall occur before this standard becomes Approved.

## 38. Consolidation Record

Version 1.0.0 consolidates the existing Diagram Standard without expanding project scope. It:

- retains immutable document identifier `OBDIA-DIAG-001`;
- normalizes the authoritative filename to `23_DIAGRAM_STANDARD.md`;
- establishes constitutional lifecycle status `Draft`;
- preserves Mermaid;
- preserves C4 Model;
- preserves Data Flow Diagram;
- preserves Sequence Diagram;
- preserves Trust Boundary Diagram;
- preserves Attack Tree;
- preserves mandatory version metadata;
- preserves mandatory scope metadata;
- preserves mandatory related-document metadata;
- adds diagram records, identifiers, filenames, source and rendered artifacts, notation-specific controls, lifecycle, authorization, evidence, connector, deployment, visual semantics, accountability, accessibility, redaction, validation, review, testing, repository, automation, synchronization, publication and traceability requirements;
- identifies documents 24, 25, 29 and 30 as forward dependencies;
- treats the `_ENTERPRISE` file as a legacy source variant of the same immutable document;
- creates no architecture implementation, validation, publication, certification or operational authorization.

## 39. Revision History

| Version | Date | Status | Authority / Owner | Change |
|---|---|---|---|---|
| 1.0.0 | 2026-08-05 | Draft | Documentation Authority; approval reserved to Project Founder | Constitutional consolidation of the original Diagram Standard: preserved all six preferred formats and all required metadata; added complete notation, security, accessibility, validation, repository and traceability controls. |
| 1.0.1 | 2026-08-06 | Draft | Documentation Authority; approval reserved to Project Founder | Reconciled obsolete forward-dependency filenames with the final canonical Knowledge Pack 01–30 filenames; no requirement text, authority, lifecycle status, implementation state or validation claim changed. |
