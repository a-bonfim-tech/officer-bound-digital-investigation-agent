# DEC-PRIV-001 — Reference Slice Privacy and Data Governance Profile

| Field | Value |
|---|---|
| Decision ID | `DEC-PRIV-001` |
| Decision Date | `2026-08-11` |
| Status | `Effective` |
| Accountable Human | André Luiz Vieira Bonfim |
| Governed Baseline | `d0c5c6f6dacf7afbcd7515edc97cd06ecccc552e` |
| Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Privacy Profile | `docs/governance/privacy/REFERENCE_SLICE_PRIVACY_AND_DATA_GOVERNANCE_PROFILE.md` |
| Fixture Manifest | `testdata/fixtures/fixture-manifest.json` |

## Decision

André Luiz Vieira Bonfim, acting as Accountable Human, approves the exact privacy and data-governance profile retained with this record for the bounded synthetic reference slice.

```text
ACCOUNTABLE_HUMAN_DECISION=APPROVE_REFERENCE_SLICE_PRIVACY_AND_DATA_GOVERNANCE_PROFILE
ACCOUNTABLE_HUMAN=André Luiz Vieira Bonfim
DECISION_DATE=2026-08-11
GOVERNED_BASELINE=d0c5c6f6dacf7afbcd7515edc97cd06ecccc552e
```

The decision freezes closed data classification, minimization, diagnostic and audit logging, raw-data restrictions, exact retention classes and lifecycle periods, disposal and verification, fixture provenance and admission, scope-expansion triggers, review reopening, accidental-real-data handling, export restrictions, telemetry and crash-report defaults.

## Approved Core Profile

```text
DATA_MINIMIZATION_RULE=COLLECT_OR_PERSIST_ONLY_FIELDS_REQUIRED_FOR_A_GOVERNED_TEST_OR_EVIDENCE_PURPOSE
DIAGNOSTIC_LOG_POLICY=MINIMIZED_EPHEMERAL_NO_RAW_PAYLOADS_NO_SECRETS_NO_STACK_TRACE_BY_DEFAULT
AUDIT_LOG_POLICY=EVENT_SPECIFIC_ALLOWLIST_CASE_BOUND_SYNTHETIC_IDS_CODES_DIGESTS_AND_COMPONENT_VERSION_ONLY
RAW_INPUT_RETENTION_POLICY=PROHIBITED_UNLESS_EXACT_SYNTHETIC_SCHEMA_TEST_REQUIRES_EPHEMERAL_COPY
RAW_CONNECTOR_RESULT_POLICY=PROHIBITED_PERSIST_DIGEST_AND_EXPLICITLY_ALLOWLISTED_RESULT_FIELDS_ONLY
EVIDENCE_MINIMIZATION_POLICY=BINDING_IDENTIFIERS_TIMESTAMPS_DIGESTS_AND_PROVENANCE_NO_REDUNDANT_RAW_CONTENT
GIT_ARTIFACT_RETENTION_POLICY=GOVERNED_LIFECYCLE_NO_RAW_PERSONAL_DATA_OR_SECRETS_ADMITTED
DISPOSAL_FAILURE_POLICY=FAIL_VISIBLE_AUDIT_ESCALATE_AND_BLOCK_SCOPE_EXPANSION
EXTERNAL_EXPORT_POLICY=PROHIBITED
TELEMETRY_POLICY=DISABLED_BY_DEFAULT
CRASH_REPORT_POLICY=EXTERNAL_CRASH_REPORTING_DISABLED
```

The seven exact retention classes and lifecycle rules, audit-event field allowlists, fixture manifest/provenance requirements, synthetic-classification test, admission/mutation gates, incident defaults and eight material-expansion trigger families in the retained privacy profile are integral to this decision.

## Fixture and Continuing-Control Boundary

Current governed evidence contains no executable fixtures. The canonical-vector file remains a specification, not complete fixture provenance. Therefore `PRIV-C02` cannot be fully satisfied.

```text
MATERIAL_PRIVACY_SCOPE_EXPANSION
=> PRIVACY_GOVERNANCE_REVIEW_REQUIRED
=> EXPANSION_NOT_AUTHORIZED_UNTIL_HUMAN_DISPOSITION
```

`PRIV-C03` remains a continuing obligation and automatically reopens when triggered. Satisfaction does not pre-approve expansion.

## Condition Disposition Upon Retention

```text
PRIV_C01=SATISFIED
PRIV_C02=PARTIALLY_SATISFIED
PRIV_C03=SATISFIED
PRIV_C03_CONTINUING_OBLIGATION=true
IMP_C01=SATISFIED
IMP_C02=SATISFIED
IMP_C03=SATISFIED
IMP_C04=SATISFIED
SEC_C01=OPEN
SEC_C02=PARTIALLY_SATISFIED
SEC_C03=SATISFIED
RES_C01=SATISFIED
```

`PRIV-C01` is satisfied as an approved definition condition. `PRIV-C02` remains partial until actual admitted fixtures have attributable manifest entries and provenance evidence. `PRIV-C03` is satisfied as a binding continuing-control definition.

## Explicit Non-Effects

```text
DEC_PRIV_001_RETAINED=true
implementation_entry_gate=BLOCKED
implementation_authorized=false
implementation_performed=false
tests_executed=false
implementation_evidence=false
operating_evidence=false
effectiveness_evidence=false
risk_accepted=false
formal_GDPR_compliance_determined=false
privacy_compliance_certified=false
DPIA_determination_created=false
legal_anonymization_determination_created=false
legal_opinion_created=false
```

No source code, executable fixture, test, CI, Ready transition, merge, release or publication is authorized or evidenced.
