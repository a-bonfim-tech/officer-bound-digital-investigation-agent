# DEC-CONTRACT-001 — Reference Slice Security Data Contracts

| Field | Value |
|---|---|
| Decision ID | `DEC-CONTRACT-001` |
| Decision Date | `2026-08-11` |
| Status | `Effective` |
| Accountable Human | André Luiz Vieira Bonfim |
| Governed Baseline | `b07048df9b811a74731b3af90f89356206b9f9ba` |
| Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Schema Version | `1.0.0` |
| Human Contract | `docs/governance/contracts/REFERENCE_SLICE_SECURITY_DATA_CONTRACTS.md` |
| Machine Schema | `schemas/reference-slice/security-data-contracts.schema.json` |
| Vector Manifest | `testdata/contracts/canonical-vectors-v1.json` |

## Decision

André Luiz Vieira Bonfim, acting as Accountable Human, approves the exact versioned security-data-contract profile retained with this record. The decision freezes schema, serialization, presence, identifiers, replay, freshness, revocation, storage, errors, connector output, hashing and negative-test semantics for any later separately authorized implementation.

```text
ACCOUNTABLE_HUMAN_DECISION=APPROVE_REFERENCE_SLICE_SECURITY_DATA_CONTRACTS
ACCOUNTABLE_HUMAN=André Luiz Vieira Bonfim
DECISION_DATE=2026-08-11
GOVERNED_BASELINE=b07048df9b811a74731b3af90f89356206b9f9ba
SCHEMA_VERSION=1.0.0
SCHEMA_STRATEGY=JSON_SCHEMA_DRAFT_2020_12_COMPOUND_SCHEMA_PLUS_STRICT_PRE_SCHEMA_DECODER
UNKNOWN_FIELD_POLICY=REJECT_BEFORE_POLICY_EVALUATION
DUPLICATE_FIELD_POLICY=REJECT_BEFORE_POLICY_EVALUATION
STRING_POLICY=EXACT_RAW_VALUE_NO_CASEFOLD_NO_TRIM
REQUEST_ID_POLICY=REQUIRED_SINGLE_USE_ATOMICALLY_RESERVED_CASE_SCOPED_AND_BOUND_TO_CANONICAL_INPUT_HASH
NONCE_POLICY=NO_SEPARATE_NONCE_REQUEST_ID_IS_THE_SINGLE_REPLAY_TOKEN
POLICY_DECISION_MODEL=CLOSED_ALLOW_OR_DENY_ONLY_ERRORS_AND_INDETERMINATE_MAP_TO_DENY
REPLAY_STORE=CASE_SCOPED_APPEND_ONLY_LOCAL_RECORD_WITH_ATOMIC_RESERVE_BEFORE_EVALUATION
CANONICALIZATION_STANDARD=RFC_8785_JCS_WITH_OBDIA_RESTRICTED_JSON_PROFILE
HASH_ALGORITHM=SHA-256
NEW_RUNTIME_DEPENDENCY_REQUIRED=false
```

## Approved Semantics

- Strict decoding is raw size, UTF-8, duplicate-key, JSON grammar, schema and governed semantic validation in that order. Unknown versions, types, enums and fields reject.
- Presence states remain distinct. Required authority fields reject absent, null, empty and invalid zero states. Go zero values never create authority.
- Security strings are exact raw values with no trimming, case folding, Unicode normalization or line-ending rewriting. Identifiers are governed lowercase prefixes plus 128-bit lowercase hex.
- `request_id` is the sole replay token, single-use, case/input-bound and atomically reserved. No separate nonce exists. Replay-store failure is fail-closed.
- Time is injected and deterministic with zero skew: `not_before <= evaluated_at < expires_at` and `evaluated_at <= decision_valid_until <= expires_at`. Timestamps are UTC RFC3339 fixed milliseconds; non-UTC offsets and leap seconds reject.
- Revocation is terminal. A later flow requires a new grant, full validation, fresh evaluation and fresh exact `ALLOW`.
- Decisions are closed `ALLOW`/`DENY`; `ERROR` and `INDETERMINATE` map to `DENY`. Every `ALLOW` binds officer, case, grant, delegation, action, connector capability/version, policy version, request, canonical input hash and validity times.
- Every retry uses a new request and fresh evaluation. Unknown outcome is not automatically retried.
- Storage is local, append-only, case-scoped and exact-case checked. Cross-case reads/writes deny. Retention/disposal remains governed by `PRIV-C01`.
- Error taxonomy version 1 and minimized external output are frozen by the human contract. All pre-ALLOW failures deny with zero connector calls and operations.
- Connector output is untrusted, closed-schema, maximum 65536 raw bytes, exact-ID-bound and cannot mutate authority. Malicious/malformed output is rejected and quarantined.
- Canonicalization is restricted RFC 8785 JCS, integers only in the approved safe range, UTF-16 code-unit lexical field order, fixed UTC timestamps and no fallback. SHA-256 uses the approved length-prefixed type/version domain separation.
- The approved state machine, I-01 through I-12 invariants, versioned vector format and distinct AC-17a/17b/17c negative-test plan are integral to this decision.

## Condition Disposition Upon Retention

```text
IMP_C01=SATISFIED
IMP_C02=SATISFIED
IMP_C03=SATISFIED
IMP_C04=SATISFIED
SEC_C01=OPEN
SEC_C02=PARTIALLY_SATISFIED
SEC_C03=SATISFIED
RES_C01=SATISFIED
PRIV_C01=OPEN
PRIV_C02=OPEN
PRIV_C03=OPEN
```

`SEC-C02` is not satisfied because deterministic-validation and negative-test evidence has not been executed. This retained package is design/governance evidence only.

## Explicit Non-Effects

```text
DEC_CONTRACT_001_RETAINED=true
implementation_entry_gate=BLOCKED
implementation_authorized=false
implementation_performed=false
source_code_authorized=false
schema_implementation_authorized=false
test_implementation_authorized=false
test_execution_authorized=false
dependency_installation_authorized=false
CI_modification_authorized=false
implementation_evidence=false
effectiveness_evidence=false
risk_accepted=false
```

The twelve material threat classes remain implementation-blocking. This decision does not establish privacy retention/disposal, fixture-provenance disposition, compliance, certification, production readiness, Ready, merge, release or publication.
