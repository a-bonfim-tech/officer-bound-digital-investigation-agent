# DEC-TECH-001 — Go 1.26.5 Technology Selection

| Field | Value |
|---|---|
| Decision ID | `DEC-TECH-001` |
| Decision Type | `TECHNOLOGY_SELECTION` |
| Decision Date | `2026-08-11` |
| Recorded Date | `2026-08-11` |
| Status | `Effective` |
| Classification | Attributable accountable-human governance decision |
| Project | Officer-Bound Digital Investigation Agent (OBDIA) |
| Architecture Authority | `ADR-0001 — Synthetic Officer-Bound Reference Slice v0.2.0` |
| Assessment HEAD | `c207f6c82c4b8ab0f98aff7e6e83d5db19fc2e32` |
| Accountable Human | André Luiz Vieira Bonfim |
| Decision Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Selected Language | Go `1.26.5` |
| Selected Toolchain | `go1.26.5` |
| Package Management | Go modules |

## Decision

André Luiz Vieira Bonfim, acting as Accountable Human, selects Go `1.26.5` and toolchain `go1.26.5` as the binding implementation technology for the bounded synthetic reference slice governed by ADR-0001. The selection is limited to that slice and does not establish an organization-wide Bonfim Labs technology standard.

```text
decision_id=DEC-TECH-001
accountable_human=André Luiz Vieira Bonfim
decision_date=2026-08-11
selected_language=Go
selected_language_version=1.26.5
selected_toolchain=go1.26.5
package_management=Go_modules
runtime_dependency_policy=STANDARD_LIBRARY_FIRST
initial_third_party_runtime_dependencies=0
decision_scope=ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE
technology_selected=true
technology_selection_binding=true
implementation_authorized=false
risk_accepted=false
```

The governed assessment compared Go `1.26.5`, TypeScript `7.0.2` with Node.js `24.18.0` LTS, and Python `3.14.6`. Its internal weighted results were `8.865/10`, `8.370/10`, and `7.855/10`, respectively. Go was selected for its combined explicit error handling, static typing, deterministic implementation ergonomics, minimal runtime dependency surface, integrated testing and fuzzing, module-integrity mechanisms, maintainability, reproducibility, CI integration and supply-chain manageability. The assessment did not find TypeScript/Node.js or Python technically unsuitable.

## Selected Baseline Profile

```text
LANGUAGE=Go
LANGUAGE_VERSION=1.26.5
TOOLCHAIN=go1.26.5
MODULE_MODE=Go_modules
FORMATTER=gofmt
BASELINE_STATIC_ANALYSIS=go_vet
TEST_FRAMEWORK=testing
FUZZ_TESTING=testing.F
RUNTIME_DEPENDENCY_POLICY=STANDARD_LIBRARY_FIRST
INITIAL_THIRD_PARTY_RUNTIME_DEPENDENCIES=0
```

No third-party runtime dependency is authorized by this decision. Any future runtime dependency requires governed necessity, provenance, integrity and reproducibility evidence.

## Deterministic Serialization and Security Requirements

Selection of Go does not mean that standard `encoding/json` output is canonical. A future, separately authorized implementation must define and test an explicit deterministic representation and preserve the ADR-0001 authorization invariant.

```text
typed_integrity_objects=true
unknown_fields_rejected=true
field_presence_semantics_explicit=true
unicode_policy_explicit=true
canonicalization_explicit=true
floats_in_integrity_material_prohibited_unless_separately_justified=true
golden_byte_vectors_required=true
hash_after_validation_and_canonicalization=true

authorization_default=DENY
positive_authorization_required=true
pre_authorization_connector_execution_possible=false
autonomous_authority_created=false
MODEL_CONTENT_CAN_AUTHORIZE=false
officer_mismatch=DENY
case_mismatch=DENY
scope_mismatch=DENY
connector_mismatch=DENY

revoked_grant_permanently_ineligible=true
revoked_grant_can_be_reused=false
revoked_grant_can_be_revalidated_to_ALLOW=false
future_execution_requires_new_distinct_valid_grant=true
future_execution_requires_fresh_policy_evaluation=true
future_execution_requires_fresh_ALLOW=true
```

## Supply-Chain and Reproducibility Boundary

The decision selects the language baseline but does not complete the exact external-tool profile.

```text
govulncheck=REQUIRED_IN_FUTURE_EVIDENCE_PROFILE
CodeQL_Go=CONDITIONAL_ON_ENTITLEMENT_VERIFICATION
CycloneDX_GoMod=REQUIRED_BEFORE_SBOM_CLAIM
secret_scanner=REQUIRED_BUT_EXACT_TOOL_VERSION_MUST_BE_GOVERNED
Staticcheck=OPTIONAL_PENDING_JUSTIFICATION
```

Before the Implementation Entry Gate may rely on a fully reproducible profile, governed evidence remains required for the exact Go toolchain pin; integrity of `govulncheck`, SBOM and secret-scanning tools; immutable GitHub Actions references; CodeQL entitlement if used; the CI operating-system/runtime matrix; dependency locking; tool installation; and the one-command verification contract. No unverified version or floating CI reference is binding through this decision.

## Condition and Gate Effect

Valid retention of this decision satisfies only the separate technology-decision prerequisite and `IMP-C01`.

```text
technology_selection_authorized=true
technology_selected=true
selected_language=Go
selected_language_version=1.26.5
selected_toolchain=go1.26.5
IMP_C01=SATISFIED
IMP_C04=PARTIALLY_SATISFIED
SEC_C03=PARTIALLY_SATISFIED
RES_C01=PARTIALLY_SATISFIED
human_implementation_owner_designated=true
bounded_experiment_authorized=true
bounded_experiment_scope=SYNTHETIC_REFERENCE_SLICE_ONLY
implementation_entry_gate=BLOCKED
implementation_authorized=false
risk_accepted=false
```

All twelve material threat classes remain implementation-blocking. The eleven ADR conditions retain their substance, owners and review dates. ADR-0001 remains the technology-neutral architectural authority; this record binds only the future implementation technology for the bounded reference slice.

## Explicit Non-Effects

This decision does not authorize source code, dependency installation, CI modification, acceptance-criteria test execution, risk acceptance, implementation, PR Ready transition, merge, release, publication, production deployment or baseline freeze.

```text
ADR_ARCHITECTURE_SCOPE_UNCHANGED=true
IMPLEMENTATION_TECHNOLOGY_PROFILE_SELECTED=true
SOURCE_CODE_MODIFICATION_AUTHORIZED=false
DEPENDENCY_INSTALLATION_AUTHORIZED=false
CI_MODIFICATION_AUTHORIZED=false
AC_TEST_EXECUTION_AUTHORIZED=false
IMPLEMENTATION_AUTHORIZED=false
RISK_ACCEPTANCE_AUTHORIZED=false
READY_AUTHORIZED=false
MERGE_AUTHORIZED=false
PUBLICATION_AUTHORIZED=false
BASELINE_FREEZE_AUTHORIZED=false
```
