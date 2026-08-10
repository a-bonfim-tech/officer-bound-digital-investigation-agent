# Synthetic Officer-Bound Reference Slice — Implementation Entry Gate Package

| Field | Value |
|---|---|
| Record ID / Version | `IMP-GATE-RS-001` / `0.1.0` |
| Status | `Proposed / BLOCKED` |
| Classification | Category C, Evidence Level D governance proposal |
| Owner | Project Founder / Accountable Human |
| Base | `e98859f7f4d39552a283cbe3ca12bcaa8c57df7f` / tree `e7cfea306926e65524c6e716c158de7df2763f5d` |
| Proposed ADR | `ADR-0001` — Proposed |

## Purpose and Verified Baseline

This package prepares a later decision and creates no authority. At baseline: main matched local/remote commit above; worktree was clean; Candidate.2 was Draft/unfrozen; Wave 1 was 105 complete and 142 pending; freeze was BLOCKED; no actual ADR, related branch or open PR existed; examples of `ADR-0001` were not allocations.

## Implementation Entry Gate

| Gate | Requirement | Evidence | Status | Missing/action | Accountable role | Effect |
|---|---|---|---|---|---|---|
| `IEG-01` scope | `IMP:71`, `IMP-REQ-013` | Proposed bounded scope | PENDING HUMAN DECISION | Approve/correct scope | Founder | BLOCKING |
| `IEG-02` architecture/experiment | `IMP:72`, `IMP-REQ-013`, `021` | ADR/request proposed | BLOCKED | Accept architecture or authorize experiment | Founder after reviews | BLOCKING |
| `IEG-03` ADR Accepted | `IMP:73`, `IMP-REQ-014` | ADR-0001 Proposed | BLOCKED | Governed acceptance | ADR decision authority | BLOCKING |
| `IEG-04` threat model reviewed | `IMP:74`, `IMP-REQ-015`, `TM-REQ-081`–`085` | Review proposed | PENDING HUMAN DECISION | Specialist dispositions | Security Reviewer | BLOCKING |
| `IEG-05` boundaries | `IMP:75`, `IMP-REQ-017` | Design boundaries defined | SATISFIED WITH CONDITIONS | Human validation | Architecture/Security | BLOCKING until confirmed |
| `IEG-06` testable security/privacy | `IMP:76` | AC-01–20 | SATISFIED WITH CONDITIONS | Sufficiency review | Security/Privacy-Governance | BLOCKING until confirmed |
| `IEG-07` acceptance criteria | `IMP:77`, `IMP-REQ-160` | AC-01–20 | SATISFIED WITH CONDITIONS | Human adoption | Implementation Reviewer | BLOCKING until adopted |
| `IEG-08` owner | `IMP:78` | Role only | PENDING HUMAN DECISION | Assign human owner | Founder | BLOCKING |
| `IEG-09` data/environment | `IMP:79` | Synthetic/local/mock/secret-free proposal | SATISFIED WITH CONDITIONS | Classification approval | Security/Privacy-Governance | BLOCKING until confirmed |
| `IEG-10` rollback/containment | `IMP:80`, `IMP-REQ-175`–`184` | Design below | SATISFIED WITH CONDITIONS | Approve and later test | Security/Implementation | BLOCKING until confirmed |

Overall: `BLOCKED`. Branch, commit, Draft PR or merge cannot satisfy the gate (`IMP-REQ-020`).

## Bounded Experiment Authorization Request

**PENDING ACCOUNTABLE-HUMAN AUTHORIZATION**

Purpose: demonstrate the officer-bound authorization invariant. Environment: isolated local development. Data and identities: synthetic only. Connectors: named local mocks only. Operations: non-destructive synthetic operations in Accepted ADR-0001. Network: no investigative-service dependency. Real cases, subjects, law-enforcement identities, credentials, production secrets, criminal infrastructure, live Dark Web interaction, exploitation and autonomous legal/investigative authority are prohibited.

Rollback: revert/delete bounded artifacts through ordinary Git history to the pre-implementation commit; no history rewrite. Containment: DENY, halt connector and preserve audit evidence on invariant failure. Expiry, if authorized: earliest of 30 days, material design/toolchain change, scope expansion request or completion of the bounded evidence run. Review triggers include failed security tests, dependencies, network/live connector, real data, identity/case/evidence model or owner changes.

## Proposed Acceptance Criteria

| ID | Criterion / proposed test |
|---|---|
| `AC-01` | A valid Synthetic Officer and Synthetic Case, valid Authorization Grant and valid Scoped Delegation are evaluated; the policy returns explicit `ALLOW`; the selected connector is allowlisted and mock-only; only after `ALLOW` the Synthetic Operation executes; only after that authorized operation an Evidence Envelope is generated; and the Audit/Provenance Record binds officer, case, authorization grant, scoped delegation/scope, policy decision, connector, operation and evidence |
| `AC-02` | Missing authorization DENY; zero connector calls |
| `AC-03` | Expired authorization DENY; injected clock |
| `AC-04` | Officer mismatch DENY |
| `AC-05` | Case mismatch DENY |
| `AC-06` | Out-of-scope action DENY |
| `AC-07` | Non-allowlisted connector DENY |
| `AC-08` | Unknown policy state DENY |
| `AC-09` | Malformed context DENY |
| `AC-10` | Untrusted content cannot expand scope |
| `AC-11` | Untrusted content cannot change officer |
| `AC-12` | Untrusted content cannot change case |
| `AC-13` | Cross-case artifact access denied |
| `AC-14` | Operation evidence only after authorized execution |
| `AC-15` | Evidence modification detectable |
| `AC-16` | Audit binds officer, case, operation, policy and mock connector |
| `AC-17` | Three distinct cases: a replayed request identifier is rejected with `DENY`, zero connector calls and zero operation execution; stale authorization context is re-evaluated and, unless a fresh explicit `ALLOW` is produced, results in `DENY`, zero connector calls and zero operation execution; authorization state `REVOKED` results in `DENY`, zero connector calls and zero operation execution; each case produces attributable denial/re-evaluation audit evidence |
| `AC-18` | No real investigative infrastructure required |
| `AC-19` | No production credential required |
| `AC-20` | Fixture manifest proves all fixtures synthetic |

No tests or results exist.

## Traceability Matrix

| Principle | Threat/risk | Security/auth requirements | Evidence requirements | AC | Proposed component |
|---|---|---|---|---|---|
| `AP-02`, `ARCH-REQ-032` | `TM-REQ-028`, `AI-RISK-002` | `ID-REQ-001`, `ID-REQ-004`, `AUTH-REQ-001`, `AUTH-REQ-018` | `EVID-REQ-011`, `EVID-REQ-023` | 04,11 | SyntheticOfficerIdentity, AuthorizationContext |
| `AP-09` | `AI-RISK-004` | `IMP-REQ-011`, `AUTH-REQ-004`, `AUTH-REQ-011`, `AUTH-REQ-012`, `AUTH-REQ-054` | `EVID-REQ-014`, `EVID-REQ-036`, `EVID-REQ-103` | 01–09 | PolicyEvaluator, AuthorizationGate |
| `AP-11` | `TM-REQ-014`, `057`, `AI-RISK-014` | `AUTH-REQ-020`, `AUTH-REQ-031`, `AUTH-REQ-117` | `EVID-REQ-033`, `EVID-REQ-100` | 05,12,13 | SyntheticCaseContext, case-scoped store abstraction |
| `AP-13` | `TM-REQ-046`–`052`, `AI-RISK-006`, `007` | `AUTH-REQ-007` | `EVID-REQ-079`, `080` | 10–12 | UntrustedContent, typed request builder |
| `AP-07` | `AI-RISK-020` | `CONN-REQ-012`, `023`, `206`–`235` | `CONN-REQ-436`–`465` | 06,07 | MockConnector, allowlist |
| `AP-08` | `TM-REQ-058`–`063`, `AI-RISK-012` | `EVID-REQ-013`, `014` | `EVID-REQ-025`–`060`, `205`–`220` | 14–16 | EvidenceEnvelope, AuditEvent, IntegrityVerifier |
| `AP-12` | `AI-RISK-005` | `AUTH-REQ-012`, `AUTH-REQ-053`, `AUTH-REQ-153`, `AUTH-REQ-207`; `IMP-REQ-081` | `EVID-REQ-014` | 03,17 | DeterministicClock, AuthorizationGate |

Requirement IDs above resolve in the repository. Component names are design labels, not implementation.

## Roles, Risks and Reviews

Decision Owner/Founder decides ADR and experiment. Architecture, Security, Privacy/Governance, Implementation and Documentation roles must record their own review. A human Implementation Owner is unassigned and blocking. Founder role concentration must be disclosed; AI is advisory, not owner, approver or risk acceptor.

Existing risks cover the slice; link `AI-RISK-002`, `004`–`007`, `012`, `014`, `019`–`021`, `024`, `027`, `028`, `032`, `035`, `038`. No new risk ID or acceptance is created.

Adversarial drafting passes found no scope conflict, implicit allow or real-person requirement. They found no implementation, operating or effectiveness evidence; unresolved serialization, replay and toolchain design; and all human gates pending.

## Governance Boundary

```text
adr_status=Proposed
adr_accepted=false
architecture_approved=false
bounded_experiment_authorized=false
implementation_entry_gate=BLOCKED
implementation_authorized=false
implementation_performed=false
source_code_created=false
tests_created=false
ci_created=false
dependencies_installed=false
risk_accepted=false
candidate2_status=Draft
candidate2_frozen=false
wave1_complete=false
baseline_freeze=BLOCKED
production_ready=false
compliance_established=false
legal_authority_created=false
autonomous_investigative_authority_created=false
```

Hashes are reported in the PR to avoid self-reference. Human decisions on ADR disposition, specialist review, owner assignment and bounded-experiment authorization remain pending.
