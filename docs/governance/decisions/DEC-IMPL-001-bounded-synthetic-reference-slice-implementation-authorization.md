# DEC-IMPL-001 — Bounded Synthetic Reference Slice Implementation Authorization

| Field | Value |
|---|---|
| Decision ID | `DEC-IMPL-001` |
| Decision Type | `BOUNDED_SYNTHETIC_REFERENCE_SLICE_IMPLEMENTATION_AUTHORIZATION` |
| Decision Date | `2026-08-12` |
| Status | `Effective` |
| Accountable Human | André Luiz Vieira Bonfim |
| Governed Baseline | `cd31d0d662f73ed93f47cc167c028a9a34d43626` |
| Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Purpose | `CONTROLLED_IMPLEMENTATION_AND_EVIDENCE_GENERATION` |

## Decision

André Luiz Vieira Bonfim, acting as Accountable Human and retained Human Implementation Owner, approves bounded implementation of the Synthetic Officer-Bound Reference Slice solely for controlled implementation and evidence generation. This authorization is synthetic-only, local, isolated and non-production. It does not authorize general implementation, production or investigative use, risk acceptance, security-tool execution, CI, release or publication.

```text
decision_id=DEC-IMPL-001
decision_type=BOUNDED_SYNTHETIC_REFERENCE_SLICE_IMPLEMENTATION_AUTHORIZATION
status=Effective
decision_date=2026-08-12
accountable_human=André Luiz Vieira Bonfim
governed_baseline=cd31d0d662f73ed93f47cc167c028a9a34d43626
scope=ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE
purpose=CONTROLLED_IMPLEMENTATION_AND_EVIDENCE_GENERATION
DEC_IMPL_001_HUMAN_DISPOSITION=APPROVED
DEC_IMPL_001_RETAINED=true
REPOSITORY_BOUNDED_IMPLEMENTATION_AUTHORIZED=true
bounded_implementation_authorized=true
controlled_evidence_generation_authorized=true
implementation_authorization_scope=SYNTHETIC_REFERENCE_SLICE_ONLY
general_or_unbounded_implementation_authorized=false
production_implementation_authorized=false
investigative_use_authorized=false
```

## Environmental Boundary

```text
environment=LOCAL_ISOLATED_NON_PRODUCTION
data=SYNTHETIC_ONLY
identities=SYNTHETIC_ONLY
cases=SYNTHETIC_ONLY
connectors=ALLOWLISTED_LOCAL_MOCKS_ONLY
runtime_network=DENY
external_export=PROHIBITED
```

Prohibited: real person or case data; real or production credentials and secrets; live connectors, Dark Web or Tor investigative use; live scraping; authenticated external APIs; real government-system access; real wallets, private keys or blockchain transactions; production deployment; autonomous legal or investigative authority.

## Supreme Authorization Invariant

```text
NO_TOOL_OR_CONNECTOR_EXECUTION_BEFORE_A_VALID_POSITIVE_AUTHORIZATION_DECISION
authorization_default=DENY
positive_authorization_required=true
ERROR_or_INDETERMINATE=DENY
model_output_creates_authority=false
connector_output_creates_authority=false
untrusted_content_creates_authority=false
```

Any violation is an immediate stop condition.

## Authorized Source Scope and Processing Order

Authorized activities are Go module initialization and bounded Go source authoring for typed governed domain objects; strict raw JSON decoding; raw byte-limit, UTF-8, duplicate-key, JSON grammar, schema and semantic validation; RFC 8785 restricted canonicalization; domain-separated SHA-256 integrity; global request reservation; case-scoped replay lifecycle; deterministic clock; authorization evaluator and exact ALLOW gate; connector allowlist and local mock connector; connector-result validation; evidence envelope; audit and provenance records; privacy-minimized logging; integrity verification; and rollback/containment implementation.

No governed specification may be silently altered. The mandatory order is:

```text
RAW_BYTE_LIMIT_VALIDATION
-> UTF8_VALIDATION
-> DUPLICATE_KEY_DETECTION
-> JSON_GRAMMAR_VALIDATION
-> JSON_SCHEMA_VALIDATION
-> GOVERNED_SEMANTIC_VALIDATION
-> GLOBAL_REQUEST_ID_RESERVATION
-> CASE_SCOPED_REPLAY_RECORD_RESERVATION
-> AUTHORIZATION_EVALUATION
-> exact ALLOW
-> CONNECTOR_INVOCATION
```

## Dependencies

```text
initial_third_party_runtime_dependencies=0
third_party_test_dependencies=0
runtime_dependency_policy=STANDARD_LIBRARY_FIRST
new_dependency_response=STOP_AND_REOPEN_GOVERNANCE
```

## Fixtures

```text
synthetic_fixture_authoring_authorized=true
fixture_provenance_authoring_authorized=true
fixture_digest_computation_authorized=true
fixture_manifest_update_authorized=true
unknown_future_fixture_preapproved=false
```

Every executable fixture must be synthetic, must not derive from real-person, live-case, production-export or live-connector data, must contain no real credentials, and must have a reproduced content digest plus attributable human `APPROVED` review with real reviewer and timestamp. Required sequence:

```text
AUTHOR
-> PROVENANCE
-> DIGEST
-> ORIGIN VALIDATION
-> HUMAN REVIEW
-> APPROVED
-> ADMISSION
-> EXECUTION
```

Fixtures with `PENDING`, `REJECTED`, missing, ambiguous or invalid provenance must not execute.

## Tests and Local Execution

AC-01 through AC-20 test authoring, unit and table-driven negative tests, deterministic serialization, golden-byte and hash-preimage tests, local-mock integration tests, bounded fuzz tests, rollback tests and containment tests are authorized. AC-17 remains distinct:

```text
AC_17a=REPLAYED_REQUEST_IDENTIFIER
AC_17b=STALE_AUTHORIZATION_CONTEXT
AC_17c=REVOKED_AUTHORIZATION_GRANT
```

Conditional on an already-installed and verified exact Go `1.26.5` toolchain, local `gofmt`, `go vet`, `go test` and `testing.F` execution are authorized with:

```text
GOENV=off
GOTOOLCHAIN=local
GOPROXY=off
GOSUMDB=off
GOVCS=*:off
CGO_ENABLED=0
runtime_network=DENY
```

Any attempted dependency or tool download requires STOP.

## Separately Governed Activities

```text
Go_1_26_5_bootstrap_authorized=false
bootstrap_network_authorized=false
gosec_execution_authorized=false
govulncheck_execution_authorized=false
TruffleHog_execution_authorized=false
cyclonedx_gomod_execution_authorized=false
security_tool_execution_authorized=false
CI_creation_authorized=false
CI_modification_authorized=false
CI_execution_authorized=false
```

## Implementation Topology

```text
IMPLEMENTATION_TOPOLOGY=NEW_IMPLEMENTATION_BRANCH_AND_NEW_DRAFT_PR_FROM_GOVERNED_BASELINE
governed_implementation_baseline=cd31d0d662f73ed93f47cc167c028a9a34d43626
preferred_implementation_branch=impl/bounded-synthetic-reference-slice-2026-08
implementation_PR_head=impl/bounded-synthetic-reference-slice-2026-08
implementation_PR_base=docs/propose-reference-slice-adr-entry-gate-2026-08
```

Implementation must not occur directly in PR #48. If PR #48 merges, advances materially or changes topology before implementation branch creation, stop and reconcile the baseline and PR topology. Future implementation commits must be signed, attributable and audit-friendly; force-push, history rewrite, squash during active evidence generation and rebase after evidence hashing are prohibited.

## Immediate Stop Conditions

Stop affected implementation activity upon real data, real-person data, real-case data, real credentials or production secrets; live connector introduction; unauthorized network access; a required new runtime dependency or an unapproved dependency; toolchain version mismatch or integrity failure; authorization-invariant failure; connector invocation without exact ALLOW; cross-case isolation failure; global request-ID uniqueness or atomic-reservation failure; fixture-provenance failure or unapproved-fixture execution; canonicalization or hash-preimage divergence; governed-specification ambiguity; material security finding; or material scope or architecture change.

Required response:

```text
STOP_IMPLEMENTATION_ACTIVITY
DENY_WHERE_APPLICABLE
HALT_CONNECTOR_EXECUTION
PRESERVE_MINIMIZED_SYNTHETIC_AUDIT_EVIDENCE
DO_NOT_EXPAND_SCOPE
REOPEN_APPROPRIATE_GOVERNANCE_REVIEW
```

## Entry Gate, Conditions and Evidence Boundary

The global gate concerns evidence-based completion/readiness. This separate authorization permits only the controlled synthetic work necessary to generate that evidence.

```text
GLOBAL_IMPLEMENTATION_ENTRY_GATE=BLOCKED_PENDING_IMPLEMENTATION_AND_EFFECTIVENESS_EVIDENCE
implementation_entry_gate=BLOCKED
IEG_04_THREAT_TREATMENT=SATISFIED_WITH_CONDITIONS
IEG_10_ROLLBACK_CONTAINMENT=SATISFIED_WITH_CONDITIONS
IEG_12_SECURITY_DATA_CONTRACTS=PARTIALLY_SATISFIED
IEG_13_PRIVACY_DATA_GOVERNANCE=PARTIALLY_SATISFIED
SEC_C01=OPEN
SEC_C02=PARTIALLY_SATISFIED
PRIV_C02=PARTIALLY_SATISFIED
IEG_10_EFFECTIVENESS_EVIDENCE=PENDING
material_threat_count=12
material_threats_blocking=12
threat_mitigation_claimed=false
implementation_performed=false
source_code_created=false
source_code_modified=false
Go_module_created=false
fixtures_created=false
tests_created=false
tests_executed=false
deterministic_validation_executed=false
negative_tests_executed=false
rollback_effectiveness_tested=false
containment_effectiveness_tested=false
security_scanners_executed=false
CI_created=false
CI_modified=false
implementation_evidence=false
operating_evidence=false
effectiveness_evidence=false
risk_accepted=false
RISK_ACCEPTANCE_AUTHORIZED=false
production_ready=false
formal_compliance_determined=false
Ready_authorized=false
merge_authorized=false
release_authorized=false
publication_authorized=false
```

Authorization is not implementation or effectiveness evidence and closes no evidence-dependent condition.
