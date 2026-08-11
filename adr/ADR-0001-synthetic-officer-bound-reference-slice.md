# ADR-0001 — Synthetic Officer-Bound Reference Slice

## 1. ADR Identifier

`ADR-0001`

Version: `0.2.0`

Classification: Decision

Decision Owner: Project Founder / Accountable Human

Contributor: AI-assisted drafting and analysis only

Implementation State: Conceptual / Designed Proposal

Base Commit: `e98859f7f4d39552a283cbe3ca12bcaa8c57df7f`

Base Tree: `e7cfea306926e65524c6e716c158de7df2763f5d`

## 2. Title

Synthetic Officer-Bound Reference Slice Architecture

## 3. Status

`Accepted`

Human Disposition: `ACCEPT_WITH_CONDITIONS`

Governance Effective Date: `2026-08-11`

This status was established prospectively through the retained governance package recorded on `2026-08-11`. It creates no implementation authority. The ADR is not Under Review, Proposed, Rejected, Deprecated or Superseded. Acceptance is limited to the technology-neutral architecture and remains subject to the governed condition register and independent downstream gates.

## 4. Date

Proposed: `2026-08-10`. Prospectively Accepted with disposition `ACCEPT_WITH_CONDITIONS`: `2026-08-11`.

## 5. Context

OBDIA has no executable reference implementation. Candidate.2 remains Draft and unfrozen, Wave 1 remains incomplete, and baseline freeze remains BLOCKED. The proposed narrow slice is:

```text
Synthetic Officer Identity -> Synthetic Case Context
-> Synthetic Authorization Grant -> Scoped Delegation
-> Deterministic Policy Evaluation -> Allowlisted Mock Connector
-> Synthetic Operation -> Evidence Envelope -> Audit / Provenance Record
```

The governing invariant is: **NO TOOL OR CONNECTOR EXECUTION BEFORE A VALID POSITIVE AUTHORIZATION DECISION.** The experiment is local, isolated, reversible, non-operational, synthetic-data-only, mock-only and secret-free. Technical authorization is a security-test object, not institutional or legal authority.

## 6. Problem Statement

Select the architecture for the smallest executable slice that could demonstrate officer, case, scope, operation, time and connector binding; deterministic ALLOW/DENY evaluation; fail-closed behavior; and evidence/provenance creation after authorized execution. Binding selection of language, runtime, framework, package manager and toolchain is deferred to a later governed implementation decision. Production connectors, real identities/cases, CI, network services and deployment are excluded.

## 7. Decision

### Accepted Decision

Adopt the technology-neutral **Synthetic Officer-Bound Reference Slice architecture** described in this ADR. This ADR does not select an implementation language, runtime, framework, package manager or toolchain.

If separately authorized, the design will use immutable or readonly-equivalent domain objects; closed, explicit `ALLOW`/`DENY` outcomes; a deterministic policy evaluator; an injected deterministic time authority; an in-memory mock-connector allowlist; and an authorization-gated orchestrator. Unknown, malformed, expired, revoked, mismatched, unsupported or indeterminate states map to `DENY`. Prompt, model and connector content are untrusted and cannot alter officer, case, delegation, policy, allowlist or authorization outcome.

Evidence and audit records are produced only after authorized synthetic execution; denial attempts use separate audit events. SHA-256 covers exact bytes or a documented deterministic representation. No runtime dependency is authorized. A later governed technology decision must fix exact language, runtime, package-manager, dependency, formatting, linting, typing or equivalent static-analysis, testing, reproducible-build and supply-chain controls before implementation.

Architecture disposition, technology selection, bounded-experiment authorization, Implementation Entry Gate disposition and implementation authorization are distinct decisions. Acceptance of this architecture does not satisfy any later gate and authorizes no code, tests, dependencies, CI, release or merge.

## 8. Alternatives Considered

| Alternative | Type/domain modeling | Determinism | Supply-chain exposure | Readability | Relative complexity | Result |
|---|---|---|---|---|---|---|
| TypeScript / Node.js 24 LTS | Strong with strict mode and discriminated unions | Strong for pure functions and injected time | Moderate; npm must be constrained | High for agent/tool boundaries | Low-moderate | Previously preferred candidate; evaluated, non-binding |
| Python with strict tooling | Good; enforcement depends on separate tooling/runtime validation | Strong | Moderate; environment controls required | High | Low | Evaluated, non-binding alternative |
| Go | Strong structs/interfaces | Strong | Low-moderate | High | Moderate for this slice | Evaluated, non-binding alternative |
| Defer technology selection | Technology-neutral architecture only | No executable validation until a later selection and authorization | No current implementation exposure | Preserves documentation-only state | None currently | Current posture; not an implementation choice |

The earlier analysis preferred TypeScript for the state-heavy tool-boundary design, lower anticipated slice complexity than Go and stronger default compile-time discrimination than ordinary Python execution. That preference is preserved as historical evaluation, not as a current selection. Node.js 24 was verified as official Active LTS on 2026-08-10; installed Node.js 26 was not treated as applicability evidence. No alternative is selected by this ADR. A later governed implementation decision must revalidate current lifecycle, security, supply-chain, maintainability and reproducibility evidence before selecting a technology profile.

## 9. Security Rationale

Complete mediation requires an exact `ALLOW` immediately before mock invocation. Authentication is not authorization; authorization is not legal authority. Officer, institution, case, purpose, action, connector, policy version and validity interval are bound. Cross-case access, non-allowlisted connectors and untrusted instructions deny. Connector results are schema-validated untrusted data. Retries require a still-valid or fresh decision. Invariant failure halts execution and preserves available audit evidence.

Applicable risks: `AI-RISK-002`, `AI-RISK-004`–`007`, `AI-RISK-012`, `AI-RISK-014`, `AI-RISK-019`–`021`, `AI-RISK-024`, `AI-RISK-027`, `AI-RISK-028`, `AI-RISK-032`, `AI-RISK-035`, `AI-RISK-038`. No risk is accepted.

## 10. Privacy Considerations

Only synthetic identities, cases and operations are permitted. Real persons, investigations, law-enforcement identities, credentials, surveillance, live Dark Web interaction and external investigative services are prohibited. Fixtures and logs must contain only fields needed to test binding and must be unmistakably synthetic. The attributable Privacy/Governance disposition is `PASS_WITH_CONDITIONS`; `PRIV-C01` through `PRIV-C03` remain open and govern implementation or scope expansion as applicable.

## 11. Trade-offs

- TypeScript could improve explicit states and reviewability but would add npm/toolchain exposure; it is not selected.
- Python or Go could satisfy the architecture with different enforcement, complexity and supply-chain trade-offs; neither is selected.
- Deferral prevents premature technology binding but postpones executable validation and requires a separate governed selection.
- No runtime dependencies reduces exposure but may require small deterministic internal validation/serialization helpers.
- Mock-only operation proves containment, not real connector behavior.
- In-memory fixtures simplify rollback but do not test distributed persistence.
- Fail-closed behavior may reject unavailable or malformed contexts; this is intentional.

## 12. Consequences

If separately authorized through independent downstream gates, a later plan may create a bounded local slice using a technology profile selected by a separate governed decision. No model, database, network, production connector, CI or deployment is implied. Any scope expansion requires new or superseding decision and renewed threat/privacy/risk review. No migration exists. Rollback uses ordinary Git history to the pre-implementation governed commit; history rewrite is prohibited.

## 13. Risks

| Risk | Existing record | Proposed treatment | State |
|---|---|---|---|
| Implicit allow | `AI-RISK-004` | Closed decisions, default deny, boundary tests | Proposed only |
| Identity/case spoofing | `AI-RISK-002`, `014` | Immutable binding, mismatch tests | Proposed only |
| Content expands authority | `AI-RISK-006`, `007`, `021` | Content/data separation | Proposed only |
| Evidence/audit tampering | `AI-RISK-012`, `028` | Integrity and provenance links | Proposed only |
| Dependency compromise | `AI-RISK-024` | Later technology decision must define dependency and supply-chain controls | Proposed only |
| Mock/live confusion | `AI-RISK-020`, `032` | Mock manifest, no network, claim boundary | Proposed only |
| False assurance | `AI-RISK-035` | Explicit scope limitations | Proposed only |
| Role concentration | `AI-RISK-038` | Disclosure and attributable gates | Open limitation |

Residual risk is unresolved because nothing is implemented or validated. No risk is accepted by the architectural acceptance.

## 14. Validation Criteria

Eligibility requires: approval of AC-01–AC-20; attributable threat/privacy/architecture reviews; a governed and versioned implementation technology profile; proof that the connector is unreachable before exact ALLOW; exhaustive DENY mapping; identity/case/scope/time/connector mismatch tests; hostile-content containment; detectable evidence modification; synthetic/network-free fixtures; bidirectional traceability; and an explicit Implementation Entry Gate decision. These are proposed criteria, not test evidence.

## 15. Related Documents

- `OBDIA-CANON-001`, `OBDIA-CONST-001`, `OBDIA-ARCH-001`, `OBDIA-ID-001`, `OBDIA-TM-001`, `OBDIA-TRUST-001`, `OBDIA-ADR-001`, `OBDIA-IMP-001`, `OBDIA-SEC-001`, `OBDIA-AUTH-001`, `OBDIA-EVID-001`, `OBDIA-CONN-001`, `OBDIA-CODE-001`, `OBDIA-TEST-001`, `OBDIA-RISK-001`.
- `docs/architecture/reference-slice/SYNTHETIC_OFFICER_BOUND_REFERENCE_SLICE_THREAT_MODEL_APPLICABILITY_REVIEW.md`.
- `docs/governance/reference-slice/SYNTHETIC_OFFICER_BOUND_REFERENCE_SLICE_IMPLEMENTATION_ENTRY_GATE_PACKAGE.md`.
- Official Node.js release schedule, consulted 2026-08-10; historical external technical evidence for an evaluated non-binding alternative only.

## 16. Review History

| Date | Activity | Result |
|---|---|---|
| 2026-08-10 | AI-assisted architecture, security, privacy, evidence, governance and scope-minimization drafting passes | Advisory draft; no human disposition |
| 2026-08-10 | Formal technical review and correction cycles on frozen commit `b42af7b74c28284cbedc49b8e8645ba7d3265d5b` | Technical review PASS for accountable-human disposition; ADR remained `Proposed` |
| 2026-08-10 | Accountable-human acceptance intent prepared and staged locally | Pre-effective evidence only; no acceptance commit and no effective status transition |
| 2026-08-10 | Accountable Human and Project Founder André Luiz Vieira Bonfim issued disposition `CORRECT` | Prior pre-effective acceptance intent superseded; correction class `TECHNOLOGY_SELECTION_DEFERMENT_AND_ARCHITECTURE_IMPLEMENTATION_DECOUPLING`; ADR remains `Proposed` |
| 2026-08-10 | Controlled technology-neutral correction prepared as version `0.2.0` | Security architecture, requirement mappings, threat model, AC-01–AC-20 and authorization invariant preserved; technology selection deferred; fresh separate review required |
| 2026-08-10 | Purported six object-bound human specialist dispositions and final disposition were described in the ADR | Later reconciliation found the aggregated assertion insufficient as retained evidence; no retroactive validation |
| 2026-08-11 | New read-only technical review of historical immutable object `77ad16819ecc716949ea47035d556bcda478c08c` | `TECH-REVIEW-ADR-0001-001=PASS`; current review, not backdated historical evidence |
| 2026-08-11 | Six individually identifiable specialist dispositions retained under `ROLE_CONCENTRATION=true` | Architecture and Documentation `PASS`; Security, Privacy/Governance, Implementation and Research `PASS_WITH_CONDITIONS`; internal governance review only |
| 2026-08-11 | Accountable Human and Project Founder André Luiz Vieira Bonfim issued new final disposition `ACCEPT_WITH_CONDITIONS` in `DEC-ENTRY-001` | ADR transitions prospectively to `Accepted`; implementation, bounded experiment, technology selection, risk acceptance, Ready, merge, release and publication remain unauthorized |

### 16.1 Accepted Conditions and Review Dates

| ID | Condition | Owner | Governance review date | Status | Evidence required | Closure/review mechanism | Effect |
|---|---|---|---|---|---|---|---|
| `SEC-C01` | Implement and test treatments for all 12 threat classes. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-17` | Open | Implementation and effectiveness evidence for all threat treatments | Security Reviewer evidence disposition | Blocks implementation/effectiveness claims, not ADR acceptance |
| `SEC-C02` | Define schemas, deterministic serialization, replay/nonce and output validation. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-17` | Open | Governed schemas and deterministic negative-test evidence | Security and Implementation Reviewer disposition | Blocks implementation/effectiveness claims, not ADR acceptance |
| `SEC-C03` | Define the supply-chain profile after the future technology selection. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-17` | Open | Exact dependency, integrity and reproducibility profile | Security Reviewer disposition after technology decision | Blocks implementation/effectiveness claims, not ADR acceptance |
| `PRIV-C01` | Define log minimization, retention and disposal. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-17` | Open | Approved minimization, retention and disposal specification | Privacy/Governance Reviewer disposition | Blocks applicable implementation, not ADR acceptance |
| `PRIV-C02` | Demonstrate provenance and synthetic classification of fixtures. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-17` | Open | Fixture manifest and provenance evidence | Privacy/Governance Reviewer disposition | Blocks applicable implementation, not ADR acceptance |
| `PRIV-C03` | Reopen Privacy/Governance Review for any material expansion of data, identity, connector or environment. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-17` | Open | New scope and privacy/governance assessment when triggered | Renewed review on material expansion | Blocks applicable scope expansion, not ADR acceptance |
| `IMP-C01` | Issue a separate governed technology decision. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-24` | Open | Accepted technology-selection decision | Separate accountable-human disposition | Blocks implementation, not ADR acceptance |
| `IMP-C02` | Define schemas and deterministic serialization. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-24` | Open | Versioned schemas and serialization specification | Implementation Reviewer disposition | Blocks implementation, not ADR acceptance |
| `IMP-C03` | Define nonce/request identifier, storage, errors and replay behavior. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-24` | Open | Versioned behavior specification and negative-test plan | Implementation and Security Reviewer disposition | Blocks implementation, not ADR acceptance |
| `IMP-C04` | Fix runtime, dependencies, test runner, lint, format, static analysis and CI. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-24` | Open | Exact governed toolchain and supply-chain profile | Implementation Reviewer disposition | Blocks implementation, not ADR acceptance |
| `RES-C01` | Revalidate lifecycle, security, supply chain, maintainability and reproducibility when the future technology-selection decision is made, using current appropriate sources. | André Luiz Vieira Bonfim — Accountable Human | `2026-08-17` | Open | Current authoritative sources and documented comparison | Research Reviewer disposition during technology selection | Blocks future technology selection, not ADR acceptance |

Condition ownership establishes accountability for review, disposition, evidence and escalation; it does not designate a Human Implementation Owner or authorize implementation. Review dates are governance deadlines, not implementation completion or authorization dates. `SEPARATE_TECHNOLOGY_SELECTION_DECISION_REQUIRED=true`.

The accepted scope is exclusively the technology-neutral architecture at the reviewed object above. The retained records are `TECH-REVIEW-ADR-0001-001`, `SPEC-DISP-ADR-0001-001` and `DEC-ENTRY-001`. No risk is accepted. Internal role concentration is disclosed; independent external assurance is not claimed. Human Implementation Owner remains unassigned. Bounded experiment, Implementation Entry Gate, implementation, technology selection, dependencies, CI, source code, Ready, merge, release, tag, publication and visibility change remain unauthorized. Independent post-retention reconciliation is required before PR metadata reconciliation or any later readiness decision.
