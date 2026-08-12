# DEC-SPEC-001 — Preimplementation Specification Reconciliation

| Field | Value |
|---|---|
| Decision ID | `DEC-SPEC-001` |
| Decision Type | `PREIMPLEMENTATION_SPECIFICATION_RECONCILIATION` |
| Decision Date | `2026-08-12` |
| Recorded Date | `2026-08-12` |
| Status | `Effective` |
| Accountable Human | André Luiz Vieira Bonfim |
| Governed Baseline | `218bc749cc7f65d222917f2a1487af2b2ea8a1a2` |
| Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Corrective Contract / Schema Version | `2.0.0` |
| Privacy Profile Version | `1.1.0` |
| Fixture Manifest Version | `1.1.0` |

## Decision

André Luiz Vieira Bonfim, acting as Accountable Human, prospectively approves the controlled resolution of exactly three preimplementation specification gaps. This decision creates no implementation, fixture, test, tool, CI, risk-acceptance, Ready or merge authority.

```text
ACCOUNTABLE_HUMAN_DECISION=AUTHORIZE_CONTROLLED_PREIMPLEMENTATION_SPECIFICATION_RECONCILIATION
decision_id=DEC-SPEC-001
decision_type=PREIMPLEMENTATION_SPECIFICATION_RECONCILIATION
status=Effective
decision_date=2026-08-12
accountable_human=André Luiz Vieira Bonfim
governed_baseline=218bc749cc7f65d222917f2a1487af2b2ea8a1a2
```

## PREIMPL-GAP-001 — Global Request Reservation

Global `request_id` uniqueness is preserved. A global atomic, append-only, single-use reservation index keyed only by `request_id` is authoritative for uniqueness and executes before the separate case-scoped replay lifecycle reservation and authorization evaluation.

```text
REQUEST_ID_UNIQUENESS_DOMAIN=GLOBAL_WITHIN_BOUNDED_REFERENCE_SLICE
REQUEST_ID_REUSE_ALLOWED=false
GLOBAL_REQUEST_RESERVATION_INDEX=GLOBAL_ATOMIC_APPEND_ONLY_SINGLE_USE_INDEX_KEYED_ONLY_BY_REQUEST_ID
CASE_SCOPED_REPLAY_RECORD=CASE_NAMESPACE_LIFECYCLE_RECORD
cross_case_request_id_reuse=DENY
REQUEST_ID_REPLAY_SEMANTICS_UNAMBIGUOUS=true
```

Same-case, cross-case or different-input-hash reuse produces `REPLAY_DETECTED`, `DENY`, zero connector calls and zero operations. Global reservation unavailability or loss of atomicity fails closed as `STORAGE_FAILURE`. `RequestReservationRecord` does not replace the case-scoped `ReplayRecord`.

## PREIMPL-GAP-002 — Fixture Provenance

The closed `FixtureProvenance` schema `2.0.0` now represents the complete `DEC-PRIV-001` fixture admission contract: type, classification, creation method, generator identity/version, timestamp, source, prohibited-origin flags, digest, review status and attributable review fields.

Only `APPROVED` fixtures with complete provenance, exact one-to-one manifest/file correspondence and a reproduced matching digest may execute. `PENDING`, `REJECTED`, missing, duplicate, orphaned or digest-mismatched fixtures are `DO_NOT_EXECUTE`.

```text
executable_fixture_count=0
actual_fixture_provenance_complete=false
fictional_fixture_provenance_created=false
fixtures_created=false
fixture_execution=false
```

## PREIMPL-GAP-003 — Error and Audit Vocabulary

`ErrorCode` and `AuditEventType` are separate closed machine-readable vocabularies. The corrective contract/schema freezes exactly 17 error codes, exactly 13 event types and one unambiguous mapping for every failure error code. Rejection and failure events require an error code; successful events prohibit fictional errors. Canonical vectors use separate `expected_audit_event_type` and `expected_error_code` fields.

```text
ERROR_TAXONOMY_COUNT=17
ERROR_TAXONOMY_INTERNAL_CONSISTENCY=true
AUDIT_EVENT_TYPE_COUNT=13
error_audit_mapping_complete=true
canonical_vectors_use_separate_event_and_error_fields=true
```

## Versioning and Historical Decisions

This is a breaking governed contract correction:

```text
SECURITY_CONTRACT_VERSION=2.0.0
SECURITY_SCHEMA_VERSION=2.0.0
CANONICAL_VECTOR_MANIFEST_VERSION=2.0.0
PRIVACY_PROFILE_VERSION=1.1.0
FIXTURE_MANIFEST_VERSION=1.1.0
DEC_CONTRACT_001_REWRITTEN=false
DEC_PRIV_001_REWRITTEN=false
DEC_SPEC_001_SUPERSEDES_ONLY_IDENTIFIED_INCONSISTENT_SPECIFICATION_PORTIONS=true
```

Version `1.0.0` remains historical. No implicit migration is permitted. This decision does not supersede the officer-bound authorization invariant, synthetic-only boundary, technology or toolchain decisions, privacy minimization/retention policy or risk boundary.

## Condition and Gate Effect

```text
PREIMPL_GAP_001=RESOLVED_AT_SPECIFICATION_LEVEL
PREIMPL_GAP_002=RESOLVED_AT_SPECIFICATION_LEVEL
PREIMPL_GAP_003=RESOLVED_AT_SPECIFICATION_LEVEL
preimplementation_specification_gaps=0
CONF_BOUNDARY_001=PENDING
CONF_SUFFICIENCY_001=PENDING
CONF_ADOPTION_001=PENDING
preimplementation_human_confirmations_required=3
DOCUMENTATION_DESIGN_GATE=BLOCKED_PENDING_HUMAN_CONFIRMATIONS
implementation_entry_gate=BLOCKED
implementation_authorized=false
```

Condition states are preserved:

```text
IMP_C01=SATISFIED
IMP_C02=SATISFIED
IMP_C03=SATISFIED
IMP_C04=SATISFIED
SEC_C01=OPEN
SEC_C02=PARTIALLY_SATISFIED
SEC_C03=SATISFIED
RES_C01=SATISFIED
PRIV_C01=SATISFIED
PRIV_C02=PARTIALLY_SATISFIED
PRIV_C03=SATISFIED
PRIV_C03_CONTINUING_OBLIGATION=true
```

`SEC-C02` remains partial only because executed deterministic-validation and negative-test evidence is absent. `PRIV-C02` remains partial only because actual admitted-fixture provenance evidence is absent. `SEC-C01` remains open pending implementation and effectiveness evidence for all 12 threat treatments.

## Explicit Non-Effects

```text
source_code_created=false
source_code_modified=false
Go_module_created=false
fixtures_created=false
tests_created=false
tests_executed=false
tools_installed=false
dependencies_installed=false
security_scanners_executed=false
CI_created=false
CI_modified=false
implementation_authorized=false
implementation_performed=false
implementation_evidence=false
operating_evidence=false
effectiveness_evidence=false
material_threat_count=12
material_threats_blocking=12
risk_accepted=false
production_ready=false
ready_authorized=false
merge_authorized=false
release_authorized=false
publication_authorized=false
```
