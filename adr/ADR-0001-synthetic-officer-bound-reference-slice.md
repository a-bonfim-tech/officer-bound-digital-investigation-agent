# ADR-0001 — Synthetic Officer-Bound Reference Slice

## 1. ADR Identifier

`ADR-0001`

Version: `0.1.0`

Classification: Decision proposal

Decision Owner: Project Founder / Accountable Human

Contributor: AI-assisted drafting and analysis only

Implementation State: Conceptual / Designed Proposal

Base Commit: `e98859f7f4d39552a283cbe3ca12bcaa8c57df7f`

Base Tree: `e7cfea306926e65524c6e716c158de7df2763f5d`

## 2. Title

Synthetic Officer-Bound Reference Slice Architecture and Implementation Technology

## 3. Status

`Proposed`

This status creates no implementation authority. The ADR is not Under Review, Accepted, Rejected, Deprecated or Superseded.

## 4. Date

Proposed: `2026-08-10`. No acceptance date exists.

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

Select one architecture and technology for the smallest executable slice that can demonstrate officer, case, scope, operation, time and connector binding; deterministic ALLOW/DENY evaluation; fail-closed behavior; and evidence/provenance creation after authorized execution. Production connectors, real identities/cases, CI, network services and deployment are excluded.

## 7. Decision

### Proposed Decision

Use **TypeScript with strict compiler settings on the supported Node.js 24 LTS line**.

If separately authorized, the design will use readonly domain objects; discriminated `ALLOW`/`DENY` outcomes; a pure policy evaluator; an injected deterministic clock; an in-memory mock-connector allowlist; and an authorization-gated orchestrator. Unknown, malformed, expired, revoked, mismatched, unsupported or indeterminate states map to `DENY`. Prompt, model and connector content are untrusted and cannot alter officer, case, delegation, policy, allowlist or authorization outcome.

Evidence and audit records are produced only after authorized synthetic execution; denial attempts use separate audit events. SHA-256 covers exact bytes or a documented deterministic representation. No runtime dependency is proposed by default. A repository-approved language profile must fix exact TypeScript, Node.js, package-manager, formatting, linting, typing, static-analysis and testing controls before implementation.

This decision is unusable unless later Accepted and every Implementation Entry Gate prerequisite independently passes. It authorizes no code, tests, dependencies, CI, release or merge.

## 8. Alternatives Considered

| Alternative | Type/domain modeling | Determinism | Supply-chain exposure | Readability | Relative complexity | Result |
|---|---|---|---|---|---|---|
| TypeScript / Node.js 24 LTS | Strong with strict mode and discriminated unions | Strong for pure functions and injected time | Moderate; npm must be constrained | High for agent/tool boundaries | Low-moderate | Recommended proposal |
| Python with strict tooling | Good; enforcement depends on separate tooling/runtime validation | Strong | Moderate; environment controls required | High | Low | Viable, not selected |
| Go | Strong structs/interfaces | Strong | Low-moderate | High | Moderate for this slice | Viable, not selected |
| Defer/status quo | None | No executable validation | None | Preserves documentation-only state | None | Safe but cannot demonstrate invariant |

TypeScript best fits the state-heavy tool-boundary design with lower slice complexity than Go and stronger default compile-time discrimination than ordinary Python execution. Node.js 24 was verified as official Active LTS on 2026-08-10; installed Node.js 26 is not treated as applicability evidence. The choice must be revisited if later dependency or language-profile review identifies unacceptable risk.

## 9. Security Rationale

Complete mediation requires an exact `ALLOW` immediately before mock invocation. Authentication is not authorization; authorization is not legal authority. Officer, institution, case, purpose, action, connector, policy version and validity interval are bound. Cross-case access, non-allowlisted connectors and untrusted instructions deny. Connector results are schema-validated untrusted data. Retries require a still-valid or fresh decision. Invariant failure halts execution and preserves available audit evidence.

Applicable risks: `AI-RISK-002`, `AI-RISK-004`–`007`, `AI-RISK-012`, `AI-RISK-014`, `AI-RISK-019`–`021`, `AI-RISK-024`, `AI-RISK-027`, `AI-RISK-028`, `AI-RISK-032`, `AI-RISK-035`, `AI-RISK-038`. No risk is accepted.

## 10. Privacy Considerations

Only synthetic identities, cases and operations are permitted. Real persons, investigations, law-enforcement identities, credentials, surveillance, live Dark Web interaction and external investigative services are prohibited. Fixtures and logs must contain only fields needed to test binding and must be unmistakably synthetic. Human privacy/fundamental-rights review remains pending.

## 11. Trade-offs

- TypeScript improves explicit states and reviewability but adds npm/toolchain exposure.
- No runtime dependencies reduces exposure but may require small deterministic internal validation/serialization helpers.
- Mock-only operation proves containment, not real connector behavior.
- In-memory fixtures simplify rollback but do not test distributed persistence.
- Fail-closed behavior may reject unavailable or malformed contexts; this is intentional.

## 12. Consequences

If Accepted and separately authorized, a later plan may create a bounded local TypeScript slice. No model, database, network, production connector, CI or deployment is implied. Any scope expansion requires new or superseding decision and renewed threat/privacy/risk review. No migration exists. Rollback uses ordinary Git history to the pre-implementation governed commit; history rewrite is prohibited.

## 13. Risks

| Risk | Existing record | Proposed treatment | State |
|---|---|---|---|
| Implicit allow | `AI-RISK-004` | Closed decisions, default deny, boundary tests | Proposed only |
| Identity/case spoofing | `AI-RISK-002`, `014` | Immutable binding, mismatch tests | Proposed only |
| Content expands authority | `AI-RISK-006`, `007`, `021` | Content/data separation | Proposed only |
| Evidence/audit tampering | `AI-RISK-012`, `028` | Integrity and provenance links | Proposed only |
| Dependency compromise | `AI-RISK-024` | No runtime dependency by default | Proposed only |
| Mock/live confusion | `AI-RISK-020`, `032` | Mock manifest, no network, claim boundary | Proposed only |
| False assurance | `AI-RISK-035` | Explicit scope limitations | Proposed only |
| Role concentration | `AI-RISK-038` | Disclosure and attributable gates | Open limitation |

Residual risk is unresolved because nothing is implemented or validated. Acceptance remains a separate human action.

## 14. Validation Criteria

Eligibility requires: approval of AC-01–AC-20; attributable threat/privacy/architecture reviews; a versioned language profile; proof that the connector is unreachable before exact ALLOW; exhaustive DENY mapping; identity/case/scope/time/connector mismatch tests; hostile-content containment; detectable evidence modification; synthetic/network-free fixtures; bidirectional traceability; and an explicit Implementation Entry Gate decision. These are proposed criteria, not test evidence.

## 15. Related Documents

- `OBDIA-CANON-001`, `OBDIA-CONST-001`, `OBDIA-ARCH-001`, `OBDIA-ID-001`, `OBDIA-TM-001`, `OBDIA-TRUST-001`, `OBDIA-ADR-001`, `OBDIA-IMP-001`, `OBDIA-SEC-001`, `OBDIA-AUTH-001`, `OBDIA-EVID-001`, `OBDIA-CONN-001`, `OBDIA-CODE-001`, `OBDIA-TEST-001`, `OBDIA-RISK-001`.
- `docs/architecture/reference-slice/SYNTHETIC_OFFICER_BOUND_REFERENCE_SLICE_THREAT_MODEL_APPLICABILITY_REVIEW.md`.
- `docs/governance/reference-slice/SYNTHETIC_OFFICER_BOUND_REFERENCE_SLICE_IMPLEMENTATION_ENTRY_GATE_PACKAGE.md`.
- Official Node.js release schedule, consulted 2026-08-10; external technical evidence only.

## 16. Review History

| Date | Activity | Result |
|---|---|---|
| 2026-08-10 | AI-assisted architecture, security, privacy, evidence, governance and scope-minimization drafting passes | Advisory draft; no human disposition |

No accepting reviewer, approval, rejection, risk acceptance or status transition exists. Internal role concentration is disclosed; independent external assurance is not claimed.
