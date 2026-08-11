# OBDIA Reference Slice Privacy and Data Governance Profile

| Field | Value |
|---|---|
| Profile ID | `OBDIA-PRIV-RS-001` |
| Version | `1.1.0` |
| Status | `Approved / reconciled by DEC-SPEC-001; substantive authority retained from DEC-PRIV-001` |
| Corrective decision date | `2026-08-12` |
| Accountable Human | André Luiz Vieira Bonfim |
| Governed baseline | `d0c5c6f6dacf7afbcd7515edc97cd06ecccc552e` |
| Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Classification | Category B internal governed profile; Evidence Level D |

## Authority and Legal Boundary

This profile freezes privacy and data-governance definitions for the bounded synthetic reference slice. Version `1.1.0`, retained by `DEC-SPEC-001`, changes only the FixtureProvenance machine-schema alignment, manifest completeness semantics and the AuditEventType/ErrorCode cross-reference. Retention periods, export policy, telemetry policy, scope-expansion triggers and legal/compliance boundaries are unchanged. It authorizes no implementation, test execution, executable fixture, dependency, CI, export, external processing, risk acceptance or scope expansion.

```text
GDPR_compliance_claimed=false
privacy_compliance_certified=false
DPIA_determination_made=false
legal_anonymization_determination_made=false
legal_opinion_created=false
CURRENT_DESIGN_INTENT=SYNTHETIC_ONLY_NO_REAL_PERSON_DATA
```

Privacy principles are conservative design constraints, not a formal legal-compliance determination.

## Closed Data Classification

```text
SYNTHETIC_DOMAIN
SYNTHETIC_OPERATION
SYNTHETIC_FIXTURE
TECHNICAL_METADATA
SECURITY_AUDIT
GOVERNANCE_EVIDENCE
DERIVED_DIGEST
TOOL_OUTPUT
PROHIBITED_REAL_PERSON_DATA
PROHIBITED_SECRET
UNKNOWN_CLASSIFICATION=FAIL_CLOSED
```

The final two values are prohibited states, not admissible processing classes. Unknown classification prohibits use, persistence, execution and repository commit.

## Data Minimization and Field Disposition

```text
DATA_MINIMIZATION_RULE=COLLECT_OR_PERSIST_ONLY_FIELDS_REQUIRED_FOR_A_GOVERNED_TEST_OR_EVIDENCE_PURPOSE
```

Every field is classified as `REQUIRED`, `OPTIONAL`, `EPHEMERAL_ONLY`, `PERSIST_DIGEST_ONLY` or `PROHIBITED`. Convenience alone is not a retention purpose.

```text
DIAGNOSTIC_LOG_POLICY=MINIMIZED_EPHEMERAL_NO_RAW_PAYLOADS_NO_SECRETS_NO_STACK_TRACE_BY_DEFAULT
AUDIT_LOG_POLICY=EVENT_SPECIFIC_ALLOWLIST_CASE_BOUND_SYNTHETIC_IDS_CODES_DIGESTS_AND_COMPONENT_VERSION_ONLY
RAW_INPUT_RETENTION_POLICY=PROHIBITED_UNLESS_EXACT_SYNTHETIC_SCHEMA_TEST_REQUIRES_EPHEMERAL_COPY
RAW_CONNECTOR_RESULT_POLICY=PROHIBITED_PERSIST_DIGEST_AND_EXPLICITLY_ALLOWLISTED_RESULT_FIELDS_ONLY
EVIDENCE_MINIMIZATION_POLICY=BINDING_IDENTIFIERS_TIMESTAMPS_DIGESTS_AND_PROVENANCE_NO_REDUNDANT_RAW_CONTENT
```

Ordinary diagnostic logs may contain only minimum operational metadata and are removed no later than one hour after test-run termination. They exclude raw connector payloads, authorization objects, evidence envelopes, credentials, secrets, environment dumps, unnecessary paths, real-person data and full synthetic payloads when identifiers or digests suffice.

Raw-input exceptions are synthetic-only, purpose-bound, test-specific, ephemeral and non-exported. `REQUIRED_RESULT_DATA` means only schema-defined, explicitly allowlisted synthetic fields required for a governed evidence purpose; it is not arbitrary raw-result retention.

## Security Audit Allowlists

Common allowed fields are `schema_version`, `audit_event_id`, `event_type`, `occurred_at`, `request_id`, `case_id`, `component`, `component_version`, `machine_outcome`, `error_code` and `digest_reference`.

| Event | Required additional fields | Optional fields | Prohibited fields | Retention class |
|---|---|---|---|---|
| `VALIDATION_REJECTED` | error code | object type | raw input | `SHORT_LIVED_TEST_EVIDENCE` |
| `AUTHORIZATION_DENIED` | decision ID, error code | grant/delegation IDs | raw context | `SHORT_LIVED_TEST_EVIDENCE` |
| `AUTHORIZATION_ALLOWED` | decision/grant/delegation IDs | policy version | raw grant | `SHORT_LIVED_TEST_EVIDENCE` |
| `REPLAY_REJECTED` | `REPLAY_DETECTED` | prior-record digest | request payload | `SHORT_LIVED_TEST_EVIDENCE` |
| `STALE_CONTEXT_REJECTED` | evaluated/valid-until times, `STALE_DECISION` | decision ID | raw context | `SHORT_LIVED_TEST_EVIDENCE` |
| `REVOKED_GRANT_REJECTED` | grant ID, `REVOKED_AUTHORIZATION` | decision ID | grant content | `SHORT_LIVED_TEST_EVIDENCE` |
| `CONNECTOR_INVOKED` | connector/operation IDs | capability version | connector request payload | `SHORT_LIVED_TEST_EVIDENCE` |
| `CONNECTOR_RESULT_REJECTED` | connector/operation IDs, error code | result digest | raw result | `GOVERNED_REVIEW_EVIDENCE` |
| `OPERATION_COMPLETED` | operation/decision IDs | integer duration | result content | `SHORT_LIVED_TEST_EVIDENCE` |
| `EVIDENCE_CREATED` | evidence/operation IDs, digest | provenance digest | embedded envelope | `GOVERNED_REVIEW_EVIDENCE` |
| `INTEGRITY_FAILURE` | object type, digest reference, error code | component version | corrupt raw object | `GOVERNED_REVIEW_EVIDENCE` |
| `STORAGE_FAILURE` | storage class, error code | component | paths, stack trace | `GOVERNED_REVIEW_EVIDENCE` |
| `INTERNAL_INVARIANT_FAILURE` | invariant ID, error code | correlation ID | raw memory/state | `GOVERNED_REVIEW_EVIDENCE` |

The closed `AuditEventType` vocabulary has 13 values and is distinct from the 17-value `ErrorCode` vocabulary in security contract/schema `2.0.0`. Rejection and failure events require the exact mapped error code. Success events do not contain fictional error codes.

## Retention Classes and Exact Lifecycle Rules

These are project-governance choices, not statutory periods.

```text
EPHEMERAL_PROCESS=memory only; ends with process lifetime
EPHEMERAL_TEST_RUN=delete no later than 1 hour after test-run termination
SHORT_LIVED_TEST_EVIDENCE=30 calendar days after run; if an attributable review remains open, retain until review closure + 7 calendar days, whichever is later
GOVERNED_REVIEW_EVIDENCE=retain until attributable review closure + 30 calendar days
GOVERNED_REPOSITORY_ARTIFACT=retain through its governed authoritative lifecycle and Git historical record
QUARANTINED_REJECTED_DATA=retain until 7 calendar days after detection or incident disposition + 24 hours, whichever is later
PROHIBITED_NOT_RETAINED=do not persist; if detected, place only into governed quarantine
```

```text
GIT_ARTIFACT_RETENTION_POLICY=GOVERNED_LIFECYCLE_NO_RAW_PERSONAL_DATA_OR_SECRETS_ADMITTED
```

Git is not used for real-person raw data, production secrets, raw connector payloads or unminimized evidence. Ordinary deletion is not represented as secure erasure of historical objects; admission prevention is mandatory.

## Disposal and Verification

```text
DISPOSAL_POLICY=STORAGE_SPECIFIC_DELETE_OR_EXPIRE_NO_UNPROVABLE_SECURE_ERASURE_CLAIM
DISPOSAL_VERIFICATION_POLICY=CLOSED_REQUESTED_COMPLETED_FAILED_ALREADY_ABSENT_OUTCOME
DISPOSAL_FAILURE_POLICY=FAIL_VISIBLE_AUDIT_ESCALATE_AND_BLOCK_SCOPE_EXPANSION
```

- Memory: release references and terminate the process.
- Temporary/test files: delete the explicit path and verify absence.
- Mutable evidence: delete after the governed period and record a minimized outcome.
- Quarantine: dispose only after attributable human incident disposition.
- Git artifacts: governed lifecycle; no secure-erasure claim for ordinary deletion.
- External stores: prohibited in the current scope.

Closed outcomes are `DELETION_REQUESTED`, `DELETION_COMPLETED`, `DELETION_FAILED` and `ALREADY_ABSENT`. Failure produces minimized audit evidence, escalates to the Accountable Human, remains visible and blocks silent closure and related expansion.

Synthetic test retention may use the injected deterministic time authority. Governance and repository records use real attributable record time. These time domains remain distinct.

## Fixture Definition and Provenance

```text
FIXTURE_DEFINITION=PURPOSE_BOUND_SYNTHETIC_INPUT_EXPECTED_OUTPUT_OR_MOCK_RESPONSE_USED_FOR_A_GOVERNED_TEST
FIXTURE_MANIFEST_POLICY=ONE_VERSIONED_MACHINE_READABLE_MANIFEST_COVERING_EVERY_ADMITTED_FIXTURE
SYNTHETIC_CLASSIFICATION_STANDARD=MANUALLY_INVENTED_OR_GENERATED_ONLY_FROM_SYNTHETIC_SEED_WITH_PROHIBITED_ORIGINS_FALSE
FIXTURE_ADMISSION_GATE=SCHEMA_VALID_CLASSIFICATION_EXPLICIT_PROVENANCE_COMPLETE_DIGEST_VERIFIED_REVIEW_APPROVED
FIXTURE_MUTATION_POLICY=DIGEST_CHANGE_REOPENS_PROVENANCE_AND_RETURNS_REVIEW_TO_PENDING
```

Admitted types are `SYNTHETIC_DOMAIN_FIXTURE`, `CANONICAL_SERIALIZATION_VECTOR`, `NEGATIVE_TEST_INPUT`, `MOCK_CONNECTOR_RESPONSE`, `AUTHORIZATION_FIXTURE` and `CASE_FIXTURE`.

Every admitted fixture requires `fixture_id`, `schema_version`, `fixture_type`, `classification`, `synthetic`, `creation_method`, `creator_or_generator`, `generator_version`, `creation_timestamp`, `source_description`, prohibited-origin flags, `content_digest`, `review_status`, `reviewer` and `review_timestamp`.

The authoritative machine representation is `FixtureProvenance` schema version `2.0.0`. `creation_method` is closed to `MANUAL_INVENTION` or `SYNTHETIC_GENERATOR`; `review_status` is closed to `PENDING`, `APPROVED` or `REJECTED`. Approved and rejected records require an attributable reviewer and review timestamp. Only `APPROVED` may enter execution; `PENDING` and `REJECTED` are `DO_NOT_EXECUTE`.

```text
EVERY_EXECUTABLE_FIXTURE_FILE_MUST_HAVE_EXACTLY_ONE_MANIFEST_ENTRY=true
EVERY_EXECUTABLE_MANIFEST_ENTRY_MUST_RESOLVE_TO_EXACTLY_ONE_FIXTURE_FILE=true
MANIFEST_DIGEST_MUST_EQUAL_RECOMPUTED_FIXTURE_DIGEST=true
ORPHAN_DUPLICATE_MISSING_OR_DIGEST_MISMATCH=DO_NOT_EXECUTE
```

Synthetic classification requires documented human invention or generation solely from synthetic seed; all prohibited-origin flags false; reproduced digest; attributable provenance; and `review_status=APPROVED`. A filename, directory or `synthetic=true` assertion alone is insufficient.

Unknown, ambiguous or missing provenance means `REJECT` and `DO_NOT_EXECUTE`. Any digest change reopens provenance and returns review to pending.

Current reality is retained in `testdata/fixtures/fixture-manifest.json`: executable fixture count is zero. `testdata/contracts/canonical-vectors-v1.json` is a `TEST_VECTOR_SPECIFICATION`, not an executable fixture or complete fixture-provenance evidence.

## Incident, Export and Network Defaults

```text
REAL_DATA_INCIDENT_POLICY=HALT_QUARANTINE_DO_NOT_EXECUTE_DO_NOT_COMMIT_OPEN_PRIVACY_AND_SECURITY_REVIEW
EXTERNAL_EXPORT_POLICY=PROHIBITED
external_data_export_allowed=false
TELEMETRY_POLICY=DISABLED_BY_DEFAULT
application_telemetry=DISABLED_BY_DEFAULT
CRASH_REPORT_POLICY=EXTERNAL_CRASH_REPORTING_DISABLED
runtime_network_policy=RUNTIME_DENY
external_log_sink_allowed=false
external_backup_allowed=false
```

Suspected secrets follow `HALT`, `QUARANTINE`, `DO_NOT_EXECUTE`, `DO_NOT_COMMIT`, `SECURITY_REVIEW`. Personal data and secrets remain separate classifications. No automatic Git-history rewrite is authorized.

Before a future data/evidence artifact is committed:

```text
SYNTHETIC_CLASSIFICATION_VERIFIED=true
PROVENANCE_COMPLETE=true
SECRET_SCAN_REQUIRED=true
REAL_PERSON_DATA_PRESENT=false
```

These are required policy controls, not currently automated controls.

## Material Scope Expansion and Continuing Review

Trigger families are `DATA`, `IDENTITY`, `CONNECTOR`, `ENVIRONMENT`, `PURPOSE`, `AI_MODEL`, `SHARING_EXPORT` and `RETENTION`.

- Data: real/pseudonymized person data, real cases, new category, raw retention, expanded logs, longer retention, profiling, external source or sharing.
- Identity: real officer/institution/subject, persistent account, new identity attributes or credentials.
- Connector: mock-to-live, network, new capability, authenticated API, scraping, Tor/Dark Web, blockchain/wallet, SaaS or real external data.
- Environment: shared/cloud/CI-retained/staging/production, database, object store, external logging, new region or processor.
- Purpose: material investigative, processing, evidence or research-purpose change.
- AI model: LLM, embeddings, vector store, external inference, telemetry, training or fine-tuning.
- Sharing/export and retention: new recipient/export/visibility or longer/different retention.

```text
MATERIAL_PRIVACY_SCOPE_EXPANSION
=> PRIVACY_GOVERNANCE_REVIEW_REQUIRED
=> EXPANSION_NOT_AUTHORIZED_UNTIL_HUMAN_DISPOSITION
PRIVACY_REVIEW_REOPEN_RULE=MATERIAL_TRIGGER_BLOCKS_EXPANSION_UNTIL_ATTRIBUTABLE_HUMAN_DISPOSITION
```

Reopened-review outcomes are `APPROVE`, `APPROVE_WITH_CONDITIONS` or `REJECT`; unknown/indeterminate means `NO_SCOPE_EXPANSION`. AI analysis cannot issue the human disposition.

```text
CONTINUING_OBLIGATION_MODEL=PRIV_C03_SATISFIED_WHILE_TRIGGER_CONTROL_REMAINS_BINDING_AND_AUTOMATICALLY_REOPENS_WHEN_TRIGGERED
```

`PRIV-C03=SATISFIED` does not pre-approve any future expansion.

## Governance Effect and Non-Effects

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
material_threat_count=12
material_threats_blocking=12
implementation_entry_gate=BLOCKED
implementation_authorized=false
implementation_performed=false
tests_executed=false
implementation_evidence=false
operating_evidence=false
effectiveness_evidence=false
risk_accepted=false
formal_compliance_determined=false
PREIMPL_GAP_001=RESOLVED_AT_SPECIFICATION_LEVEL
PREIMPL_GAP_002=RESOLVED_AT_SPECIFICATION_LEVEL
PREIMPL_GAP_003=RESOLVED_AT_SPECIFICATION_LEVEL
preimplementation_specification_gaps=0
```
