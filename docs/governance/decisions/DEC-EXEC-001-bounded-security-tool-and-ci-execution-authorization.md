# DEC-EXEC-001 — Bounded Security Tool and CI Execution Authorization

| Field | Value |
|---|---|
| Decision ID | `DEC-EXEC-001` |
| Decision Type | `BOUNDED_FROZEN_SECURITY_TOOL_AND_CI_EXECUTION_AUTHORIZATION` |
| Decision Date | `2026-08-12` |
| Status | `Effective` |
| Accountable Human | André Luiz Vieira Bonfim |
| Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Authority Source | Accountable-Human master continuous repository execution authorization dated `2026-08-12` |

## Decision

The Accountable Human authorizes bootstrap and execution of only the exact security tools and CI profile already frozen by `DEC-TOOLCHAIN-001`, solely for the bounded synthetic OBDIA reference slice and controlled evidence generation.

```text
EXACT_FROZEN_SECURITY_TOOL_BOOTSTRAP_AUTHORIZED=true
EXACT_FROZEN_SECURITY_TOOL_EXECUTION_AUTHORIZED=true
ALLOWLISTED_SECURITY_TOOL_BOOTSTRAP_NETWORK_AUTHORIZED=true
BOUNDED_SECURITY_EVIDENCE_REFRESH_NETWORK_AUTHORIZED=true
BOUNDED_CI_CREATION_AUTHORIZED=true
BOUNDED_CI_MODIFICATION_AUTHORIZED=true
BOUNDED_CI_EXECUTION_AUTHORIZED=true
```

## Frozen Identities

```text
govulncheck_version=v1.6.0
gosec_version=2.28.0
TruffleHog_version=3.96.0
cyclonedx_gomod_version=1.10.0
CI_platform=GitHub_Actions_ubuntu-24.04
ACTION_CHECKOUT_SHA=3d3c42e5aac5ba805825da76410c181273ba90b1
ACTION_SETUP_GO_SHA=NOT_USED
GITHUB_TOKEN_PERMISSIONS=contents:read
```

All acquisition must use the exact sources, versions and integrity anchors retained by `DEC-TOOLCHAIN-001`. Substitution, floating versions, self-update and unverified artifacts are prohibited.

## Boundaries

```text
environment=LOCAL_ISOLATED_NON_PRODUCTION
data=SYNTHETIC_ONLY
runtime_network=DENY
bootstrap_network=ALLOWLISTED_EXACT_ARTIFACTS_ONLY
security_evidence_refresh_network=BOUNDED_ALLOWLISTED_ONLY
new_runtime_dependencies=false
new_test_dependencies=false
production_credentials=false
live_connectors=false
risk_accepted=false
Ready_authorized=false
merge_authorized=false
release_authorized=false
publication_authorized=false
```

CI must use `ubuntu-24.04`, `permissions: contents: read`, no production secrets, no `pull_request_target` execution of untrusted code, no floating third-party actions and no setup-go substitution. The official Go `1.26.5` archive must be verified before use. Application runtime network remains denied.

## Evidence Boundary

This decision is execution authority, not evidence that tools or CI were installed, executed or passed. Findings must not be suppressed to obtain PASS. Actual commands, versions, integrity results, exit codes and minimized output hashes must be retained separately.

## Non-Effects

This decision does not authorize real data, live or production infrastructure, production credentials, external investigative operations, new dependencies, tool substitution, risk acceptance, PR Ready, merge, release or publication.
