# OBDIA Reference Slice Security Data Contracts

| Field | Value |
|---|---|
| Contract ID | `OBDIA-CONTRACT-RS-001` |
| Version | `1.0.0` |
| Status | `Approved / retained by DEC-CONTRACT-001` |
| Decision date | `2026-08-11` |
| Accountable Human | André Luiz Vieira Bonfim |
| Governed baseline | `b07048df9b811a74731b3af90f89356206b9f9ba` |
| Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Classification | Category B internal governed contract; Evidence Level D |

## Authority and Boundary

This contract freezes the definition semantics for a later separately authorized implementation. It creates no source code, implementation, schema-validator runtime, dependency, test execution, CI, effectiveness evidence, risk acceptance or Implementation Entry Gate authorization.

The governing invariant is: **NO TOOL OR CONNECTOR EXECUTION BEFORE A VALID POSITIVE AUTHORIZATION DECISION.** Model and connector content are untrusted data and cannot authorize.

```text
SCHEMA_STRATEGY=JSON_SCHEMA_DRAFT_2020_12_COMPOUND_SCHEMA_PLUS_STRICT_PRE_SCHEMA_DECODER
SCHEMA_VERSION=1.0.0
UNKNOWN_FIELD_POLICY=REJECT_BEFORE_POLICY_EVALUATION
DUPLICATE_FIELD_POLICY=REJECT_BEFORE_POLICY_EVALUATION
PRESENCE_SEMANTICS=ABSENT_NULL_EMPTY_STRING_EMPTY_ARRAY_EMPTY_OBJECT_ZERO_FALSE_AND_RAW_EMPTY_ARE_DISTINCT
STRING_POLICY=EXACT_RAW_VALUE_NO_CASEFOLD_NO_TRIM
UNICODE_POLICY=VALID_UTF8_REQUIRED_NO_NORMALIZATION_CONTROLS_NUL_AND_UNPAIRED_SURROGATES_REJECTED
REQUEST_ID_POLICY=REQUIRED_SINGLE_USE_ATOMICALLY_RESERVED_CASE_SCOPED_AND_BOUND_TO_CANONICAL_INPUT_HASH
NONCE_POLICY=NO_SEPARATE_NONCE_REQUEST_ID_IS_THE_SINGLE_REPLAY_TOKEN
POLICY_DECISION_MODEL=CLOSED_ALLOW_OR_DENY_ONLY_ERRORS_AND_INDETERMINATE_MAP_TO_DENY
REPLAY_STORE=CASE_SCOPED_APPEND_ONLY_LOCAL_RECORD_WITH_ATOMIC_RESERVE_BEFORE_EVALUATION
CANONICALIZATION_STANDARD=RFC_8785_JCS_WITH_OBDIA_RESTRICTED_JSON_PROFILE
HASH_ALGORITHM=SHA-256
NEW_RUNTIME_DEPENDENCY_REQUIRED=false
```

## Strict Decode and Presence

The mandatory sequence is:

```text
RAW_BYTE_LIMIT_VALIDATION
-> UTF8_VALIDATION
-> DUPLICATE_KEY_DETECTION
-> JSON_GRAMMAR_VALIDATION
-> JSON_SCHEMA_VALIDATION
-> GOVERNED_SEMANTIC_VALIDATION
```

Unknown schema/object/enum values, unknown or duplicate fields, trailing content, multiple top-level values, wrong types, unexpected null, missing required values, invalid identifiers/timestamps, malformed UTF-8 and oversized input reject before policy evaluation. JSON Schema does not detect duplicate raw object members; that remains a mandatory pre-schema control.

For authorization-bearing required fields, absent, null, empty string and an unapproved zero value are invalid. Boolean `false` is distinct from absence. Empty arrays and objects are permitted only where expressly allowed by the schema. Go zero values never establish authority.

Security strings are byte-exact: no trimming, case folding, Unicode normalization or line-ending rewriting. Security identifiers are ASCII constrained. NUL, disallowed controls, bidi controls, malformed UTF-8 and unpaired surrogates reject.

## Versioning

Every persisted or integrity-relevant object contains `schema_version` and `object_type`. Version `1.0.0` is the only accepted initial version. Unknown versions reject. No implicit migration, best-effort interpretation or silent coercion is permitted. A breaking change requires a new major version, new governed decision and regenerated golden vectors.

## Domain Objects

The compound schema defines: `SyntheticOfficerIdentity`, `SyntheticCaseContext`, `AuthorizationGrant`, `ScopedDelegation`, `RequestedAction`, `AuthorizationContext`, `PolicyDecision`, `ConnectorCapability`, `ConnectorRequest`, `ConnectorResult`, `EvidenceEnvelope`, `AuditEvent`, `ProvenanceRecord`, `ReplayRecord` and `FixtureProvenance`. Objects are closed and required fields are explicit.

## Identifier Contract

Identifiers are case-sensitive, immutable, canonical lowercase ASCII prefix plus 128-bit lowercase hexadecimal value. Their exact length is prefix-dependent and at most 38 bytes (`grant_` plus 32 hexadecimal characters). They are not reused after terminal lifecycle states.

| Identifier | Prefix | Generation authority | Uniqueness/reuse |
|---|---|---|---|
| Officer | `off_` | Synthetic fixture authority | Unique in fixture corpus; no reassignment |
| Case | `case_` | Synthetic case authority | Global in slice; no reuse |
| Grant | `grant_` | Authorization authority | Global; expired/revoked IDs never reused |
| Delegation | `dlg_` | Delegation authority | Global; one grant binding |
| Request | `req_` | Orchestrator using 128-bit CSPRNG | Global and single-use |
| Decision | `dec_` | Policy evaluator | One per evaluation |
| Connector | `conn_` | Governed allowlist authority | Stable governed identity |
| Operation | `op_` | Gated orchestrator | One authorized attempt |
| Evidence | `evid_` | Evidence producer | Global; no reuse |
| Audit event | `evt_` | Audit producer | Global; append-only |

## Request, Replay, Freshness and Revocation

`request_id` is the sole replay token; no separate nonce exists. It is bound to exact case and canonical input hash and atomically reserved before evaluation. Duplicate reservation or replay-store failure produces `DENY`, zero connector calls and zero operation execution.

Replay states are closed: `RESERVED`, `DENIED`, `CONNECTOR_INVOKED`, `RESULT_REJECTED`, `COMPLETED`, `INDETERMINATE_OUTCOME`. A record never returns to `RESERVED`.

```text
clock_source=INJECTED_DETERMINISTIC_TIME_AUTHORITY
clock_skew=0
VALIDITY_RULE=not_before <= evaluated_at < expires_at
DECISION_VALIDITY_RULE=evaluated_at <= decision_valid_until <= expires_at
TIMESTAMP=UTC_RFC3339_FIXED_MILLISECONDS_YYYY-MM-DDTHH:MM:SS.mmmZ
```

Non-UTC offsets, leap seconds and noncanonical timestamps reject. `REVOKED` is terminal and permanently ineligible. A later flow requires a new distinct grant, complete validation, fresh evaluation and fresh exact `ALLOW`.

## Policy Decision and ALLOW Binding

Only `ALLOW` and `DENY` are decision values. Evaluator `ERROR` or `INDETERMINATE` maps externally to `DENY`. Only exact `ALLOW` can reach a connector.

Every `ALLOW` binds officer, case, authorization grant, delegation, requested action, connector, connector capability/version, policy version, request ID, canonical input hash, `evaluated_at` and `decision_valid_until`. It is not transferable.

Every retry requires a new request ID, fresh complete validation, fresh evaluation and fresh exact `ALLOW`. Timeout or unknown execution becomes `INDETERMINATE_OUTCOME`; automatic retry is prohibited.

## Storage and Persistence

Storage uses structural case namespaces plus an exact recorded `case_id` match. Cross-case read and write deny. The replay record is local, append-only and atomically reserved. Security-state storage failure fails closed.

| Object | Initial classification |
|---|---|
| AuthorizationGrant | `PERSISTED_FOR_TEST` |
| PolicyDecision | `PERSISTED_AS_EVIDENCE` |
| RequestID / ReplayRecord | `PERSISTED_AS_EVIDENCE` |
| Nonce | `NOT_STORED` |
| ConnectorResult | `EPHEMERAL_EXCEPT_VALIDATED_DIGEST_AND_REQUIRED_RESULT_DATA` |
| EvidenceEnvelope | `PERSISTED_AS_EVIDENCE` |
| AuditEvent | `PERSISTED_AS_EVIDENCE` |
| FixtureProvenance | `PERSISTED_FOR_TEST` |

Retention duration and disposal remain governed by `PRIV-C01`; this contract invents no retention period.

## Error Contract

Error taxonomy version `1` is closed:

```text
VALIDATION_ERROR
AUTHORIZATION_DENIED
EXPIRED_AUTHORIZATION
NOT_YET_VALID_AUTHORIZATION
REVOKED_AUTHORIZATION
IDENTITY_MISMATCH
CASE_MISMATCH
SCOPE_MISMATCH
CONNECTOR_NOT_ALLOWED
REPLAY_DETECTED
STALE_DECISION
MALFORMED_CONNECTOR_RESULT
INTEGRITY_FAILURE
STORAGE_FAILURE
INTERNAL_INVARIANT_FAILURE
UNSUPPORTED_SCHEMA
OVERSIZED_INPUT
```

External errors expose only a machine code, correlation ID and minimized diagnostic. They do not expose raw sensitive input, internal paths, stack traces, credentials or secrets. Synthetic minimized audit detail is distinct.

Every pre-ALLOW failure maps to `DENY`, zero connector calls, zero operations, no evidence envelope and an attributable audit where available. Invalid connector output makes the prior `ALLOW` non-reusable, permits no follow-up operation and creates no valid evidence.

## Connector Output Validation

Raw result maximum is `65536` bytes. Results use a known closed versioned schema and must match connector, request, operation and schema identifiers. They are always untrusted data and contain no authority fields. They cannot change officer, case, authorization, delegation, scope, allowlist, decision or next operation. Malformed, malicious, oversized or instruction-bearing results are `REJECT_AND_QUARANTINE`; no follow-up and no valid evidence envelope are permitted.

## Canonicalization and Integrity

The canonical form is RFC 8785 JCS with the restricted OBDIA profile: valid UTF-8, no normalization, no null in integrity objects, no output whitespace or final newline, RFC 8785 UTF-16 code-unit lexical member ordering, and schema-semantic array order.

Numbers are integers only in `[-9007199254740991, 9007199254740991]`. Floats, exponents, NaN, infinities and negative zero are prohibited.

SHA-256 input is:

```text
ASCII("OBDIA") || 0x00
|| uint16be(len(object_type)) || UTF8(object_type)
|| uint16be(len(schema_version)) || ASCII(schema_version)
|| uint64be(len(canonical_bytes)) || canonical_bytes
```

Canonicalization failure means `NO_HASH`, `NO_ALLOW`, `NO_VALID_EVIDENCE`, `FAIL_CLOSED`; no fallback encoding exists.

## State Machine and Invariants

```text
RECEIVED -> RAW_VALIDATED -> SCHEMA_VALIDATED -> REPLAY_RESERVED
-> AUTHORIZATION_EVALUATED -> ALLOWED -> CONNECTOR_INVOKED
-> RESULT_VALIDATED -> OPERATION_COMPLETED -> EVIDENCE_CREATED
```

Terminal/failure states are `INPUT_REJECTED`, `REPLAY_REJECTED`, `DENIED`, `CONNECTOR_FAILED`, `RESULT_REJECTED`, `INTEGRITY_FAILED`, `STORAGE_FAILED`, `INDETERMINATE_OUTCOME`. Direct `RECEIVED -> CONNECTOR_INVOKED`, `DENIED -> CONNECTOR_INVOKED`, `REPLAY_REJECTED -> ALLOWED`, `REVOKED -> ACTIVE` and `RESULT_REJECTED -> EVIDENCE_CREATED` transitions are prohibited.

```text
I-01 connector call requires current exact ALLOW
I-02 evidence requires completed authorized operation
I-03 DENY implies zero connector calls
I-04 malformed context cannot become valid evaluation context
I-05 revoked grant is terminal
I-06 request_id cannot execute twice
I-07 cross-case read/write is denied
I-08 untrusted content cannot mutate authority
I-09 invalid connector result cannot create valid evidence
I-10 integrity bytes are deterministic
I-11 replay-store failure is fail-closed
I-12 INDETERMINATE never equals ALLOW
```

## Golden Vectors and Negative Tests

The governed manifest is `testdata/contracts/canonical-vectors-v1.json`. It is specification evidence, not executed validation. Unknown future canonical bytes or digests use `TO_BE_COMPUTED_BY_FUTURE_AUTHORIZED_REFERENCE_ENCODER`; fictional hashes are prohibited.

Future table-driven negative tests cover missing/expired/revoked authorization, identity/case/scope mismatch, non-allowlisted connector, unknown state, malformed/duplicate/unknown fields, authority mutation, cross-case access, evidence ordering/integrity, duplicate request, stale context and malicious connector output. AC-17a replay, AC-17b stale context and AC-17c revoked grant remain distinct.

## Governance Effect

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
material_threat_count=12
material_threats_blocking=12
implementation_entry_gate=BLOCKED
implementation_authorized=false
implementation_performed=false
tests_executed=false
implementation_evidence=false
effectiveness_evidence=false
risk_accepted=false
```
