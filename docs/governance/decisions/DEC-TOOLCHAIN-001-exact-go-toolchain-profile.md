# DEC-TOOLCHAIN-001 — Exact Reproducible Go Toolchain Profile

| Field | Value |
|---|---|
| Decision ID | `DEC-TOOLCHAIN-001` |
| Decision Type | `EXACT_REPRODUCIBLE_TOOLCHAIN_AND_SUPPLY_CHAIN_PROFILE` |
| Decision Date | `2026-08-11` |
| Recorded Date | `2026-08-11` |
| Status | `Effective` |
| Accountable Human | André Luiz Vieira Bonfim |
| Project | Officer-Bound Digital Investigation Agent (OBDIA) |
| Governed Baseline | `e7884149c80d7e8ac56c746150630e5d39f4d60f` |
| Decision Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Human-Readable Profile | `docs/governance/toolchain/GO_TOOLCHAIN_AND_SUPPLY_CHAIN_PROFILE.md` |
| Machine-Readable Manifest | `toolchain/go-toolchain-integrity.json` |

## Decision

André Luiz Vieira Bonfim, acting as Accountable Human, prospectively approves the exact Go toolchain and supply-chain profile retained with this record for later Implementation Entry Gate evaluation of the bounded synthetic reference slice. The machine-readable manifest and human-readable profile are integral evidence of this decision.

```text
ACCOUNTABLE_HUMAN_DECISION=APPROVE_EXACT_REPRODUCIBLE_GO_TOOLCHAIN_PROFILE
toolchain_profile_approved=true
GO_LANGUAGE_VERSION=1.26.5
GO_TOOLCHAIN_VERSION=go1.26.5
GO_RUNTIME_DEPENDENCIES=0
FORMATTER=gofmt
BASELINE_STATIC_ANALYSIS=go_vet
OPTIONAL_STATIC_ANALYSIS=DEFER_STATICCHECK
TEST_FRAMEWORK=testing
FUZZ_FRAMEWORK=testing.F
DEPENDENCY_SCANNER=govulncheck
DEPENDENCY_SCANNER_VERSION=v1.6.0
SAST=gosec
SAST_VERSION=2.28.0
SECRET_SCANNER=TruffleHog
SECRET_SCANNER_VERSION=3.96.0
SBOM_TOOL=cyclonedx-gomod
SBOM_TOOL_VERSION=1.10.0
CI_PLATFORM=GitHub_Actions_ubuntu-24.04
ACTION_CHECKOUT_SHA=3d3c42e5aac5ba805825da76410c181273ba90b1
GITHUB_TOKEN_PERMISSIONS=contents:read
CGO_POLICY=CGO_ENABLED=0
BUILD_REPRODUCIBILITY_FLAGS=-trimpath -buildvcs=false
NETWORK_POLICY=ALLOWLISTED_BOOTSTRAP;OFFLINE_VERIFICATION_AND_TESTS;RUNTIME_DENY
ONE_COMMAND_VERIFICATION=GOENV=off GOTOOLCHAIN=local go run ./cmd/verify all
```

## Approved Integrity Anchors

```text
GO_INTEGRITY_DARWIN_ARM64=efb87ff28af9a188d0536ef5d42e63dd52ba8263cd7344a993cc48dd11dedb6a
GO_INTEGRITY_LINUX_AMD64=5c2c3b16caefa1d968a94c1daca04a7ca301a496d9b086e17ad77bb81393f053
DEPENDENCY_SCANNER_INTEGRITY=h1:FeMO9Rm/HwyduOztbvKcOw+zvDEPr4I4aQNSfevFcKY=
SAST_SOURCE_COMMIT=9e75c0576c9878035d4221392108d458abe10fc3
SAST_INTEGRITY_DARWIN_ARM64=6c4993a0ab5e3007d66c87cbcb4e3948f8000971f8eeaf3ac269cbc87a603ba4
SAST_INTEGRITY_LINUX_AMD64=d7882e505b1ff345d458bf0e893eec8019bc849f861ad73a212869540dd505ff
SAST_CHECKSUMS_MANIFEST_SHA256=2b494032b5ea44c5451a49a9e8c178781937eb0236cdc15087cc43e222c3a450
SECRET_SCANNER_INTEGRITY_DARWIN_ARM64=87478306b95ca2420cfb844b7582383ac60b922e262350a0088e797f328d2e62
SECRET_SCANNER_INTEGRITY_LINUX_AMD64=7105f1cd6577f058a9e39d0578f1a99c8a1e481e4d3512cd8a09acfe22a0fdc0
SBOM_TOOL_INTEGRITY_DARWIN_ARM64=2a8e887efb07ed4e36d7ead879a15e055b1e50d60e5d45e407f9f2a27db4dc46
SBOM_TOOL_INTEGRITY_LINUX_AMD64=5cce8ae99a5181be6a610ea5ed9ca9d596937cc04dc1a8f6f6b5e462d8c9900e
```

## Module, SAST and CI Boundary

```text
GO_MODULE_POLICY=NO_REQUIRE_NO_REPLACE_NO_VENDOR_NO_PRIVATE_MODULES;GOPROXY=https://proxy.golang.org;GOSUMDB=sum.golang.org;GOVCS=*:off
GOPRIVATE=
GONOSUMDB=
GONOPROXY=
SAST_NETWORK_POLICY=ALLOWLISTED_BOOTSTRAP;OFFLINE_SOURCE_ANALYSIS;NO_AUTO_UPDATE;NO_AI_FEATURES
SAST_CI_PERMISSIONS=contents:read
SAST_CI_ACTION_SHA=NOT_APPLICABLE
ACTION_SETUP_GO_SHA=NOT_USED
OTHER_ACTION_SHAS=NONE
CACHE_POLICY=DISABLED_INITIALLY
```

CodeQL is not a dependency of the frozen profile. Its entitlement was not authoritatively determined and code scanning was disabled when the decision was issued. No GitHub Code Security or code-scanning enablement is authorized.

## Condition Disposition

The completed technology-time research and resolved exact profile satisfy the governance requirements to define and freeze the applicable runtime, tools, scanners, integrity, CI, network and reproducibility profile.

```text
IMP_C01=SATISFIED
IMP_C04=SATISFIED
SEC_C03=SATISFIED
RES_C01=SATISFIED
SEC_C01=OPEN
SEC_C02=OPEN
PRIV_C01=OPEN
PRIV_C02=OPEN
PRIV_C03=OPEN
IMP_C02=OPEN
IMP_C03=OPEN
```

## Explicit Non-Effects

This decision does not authorize or evidence implementation, source code, dependency or tool installation, CI creation or modification, test or scanner execution, AC-01 through AC-20 execution, threat mitigation, effectiveness, risk acceptance, PR Ready, merge, release, publication, production deployment or baseline freeze.

```text
implementation_entry_gate=BLOCKED
implementation_authorized=false
source_code_authorized=false
dependency_installation_authorized=false
tool_installation_authorized=false
CI_modification_authorized=false
test_execution_authorized=false
risk_accepted=false
ready_authorized=false
merge_authorized=false
publication_authorized=false
```

The twelve material threats remain implementation-blocking. `SEC-C01`, `SEC-C02`, `PRIV-C01` through `PRIV-C03`, `IMP-C02` and `IMP-C03` remain open and independently governed.
