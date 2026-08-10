# Synthetic Officer-Bound Reference Slice — Threat-Model Applicability Review

| Field | Value |
|---|---|
| Record ID | `TM-REVIEW-RS-001` |
| Version / Status | `0.1.0` / `Proposed` |
| Classification | Category C, Evidence Level D review proposal |
| Authority | Review evidence only |
| Owner / Approval | Security Reviewer / Project Founder after applicable reviews |
| Reviewed architecture | Proposed `ADR-0001` at base `e98859f7f4d39552a283cbe3ca12bcaa8c57df7f` |
| Implementation | Not Implemented |

## Scope and Boundaries

This bounded review applies Draft `OBDIA-TM-001` to a local, isolated, synthetic-only, mock-only slice. It neither modifies nor approves that baseline. Assets are synthetic identity/case/delegation/authorization, policy result, allowlist, operation result, evidence and audit/provenance. The model and all fixture/connector content are untrusted, never authority.

```text
Untrusted content -> typed request construction
-> authorization boundary (officer + case + scope + action + time + connector)
-> exact ALLOW only -> local mock connector
-> synthetic result -> evidence envelope -> audit/provenance
```

## Applicability Register

| Review item | Source | Component | Proposed control/test | Uncertainty | Status |
|---|---|---|---|---|---|
| Authorization bypass | `TM-REQ-022`; `AI-RISK-004` | Evaluator/gate | Exact ALLOW; AC-02/03/06/08/09 | No implementation | BLOCKING |
| Identity confusion | `TM-REQ-028`, `053`–`057`; `AI-RISK-002` | Identity/context | Immutable binding; AC-04/11 | Identity schema pending | BLOCKING |
| Case confusion/contamination | `TM-REQ-014`, `021`, `057`; `AI-RISK-014` | Context/evidence | Per-case binding; AC-05/12/13 | Persistence absent | BLOCKING |
| Prompt injection | `TM-REQ-046`–`052`; `AI-RISK-006`, `007` | Untrusted-content boundary | Content cannot alter typed authority; AC-10–12 | No model in initial slice | BLOCKING |
| Connector misuse/privilege expansion | `TM-REQ-048`; `CONN-REQ-206`–`235`; `AI-RISK-020` | Gate/mock | Allowlist and exact scope; AC-06/07 | Interface absent | BLOCKING |
| Evidence/audit tampering | `TM-REQ-027`, `058`–`063`; `AI-RISK-012`, `028` | Evidence/audit | SHA-256, immutable links; AC-14–16 | Serialization pending | BLOCKING |
| Replay/stale decision | `AUTH-REQ-012`, `AUTH-REQ-053`, `AUTH-REQ-207`; `IMP-REQ-081`; `AI-RISK-005` | Gate/clock | Distinct stale-context and duplicate-request denial/re-evaluation; AC-03/17 | Nonce design pending | BLOCKING |
| Malformed context | `TM-REQ-022`; `AUTH-REQ-054`; `IMP-REQ-011` | Validator | Reject before evaluation; AC-09 | Schema pending | BLOCKING |
| Compromised dependency | `TM-REQ-064`–`068`; `AI-RISK-024` | Toolchain | No runtime dependency by default; integrity review | Toolchain pending | BLOCKING |
| Malicious connector result | `TM-REQ-017`, `047`; `AI-RISK-021` | Result boundary | Schema/size validation, no instruction execution | Result schema pending | BLOCKING |
| Secret exposure | Supply-chain category; `AI-RISK-019` | Fixtures/repo | Synthetic-only and secret scan; AC-19/20 | Enforcement not tested | BLOCKING |
| Insecure/unsafe failure | `TM-REQ-022`; `IMP-REQ-175`–`184` | All boundaries | Empty allowlist, DENY, halt/preserve audit | Implementation absent | BLOCKING |

No new baseline threat ID is created. These are derived review items within `TM-REVIEW-RS-001`.

## Abuse Cases and Review

- Hostile content requests another officer/case: typed bindings remain unchanged and connector is not called.
- Authorization is missing, malformed, revoked or expired: DENY and denial audit only.
- Connector is not allowlisted or resembles a live service: DENY before resolution.
- Mock result contains instructions/malformed data: reject or quarantine, never execute.
- Evidence changes: integrity verification fails.
- Retry uses stale/mismatched decision: re-evaluate or DENY.

Architecture review found no drafting conflict with the synthetic scope. Security review found the right fail-closed design but no implementation proof. Privacy review found no required real-person processing. Evidence review found proposed binding fields but unresolved serialization/storage. Governance review confirmed Proposed/BLOCKED language. Scope review confirmed one policy path, one mock and one evidence path is the minimum useful slice.

Human Architecture, Security and Privacy/Governance Reviewer dispositions remain required. Critical unresolved threats block implementation. Merge alone does not establish review or approval.

Current status: `PENDING HUMAN REVIEW`.

Repository: a-bonfim-tech/officer-bound-digital-investigation-agent

Version: e98859f7f4d39552a283cbe3ca12bcaa8c57df7f
