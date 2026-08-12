# OBDIA Exact Go Toolchain and Supply-Chain Profile

| Field | Value |
|---|---|
| Profile ID | `GO-TOOLCHAIN-PROFILE-001` |
| Version | `1.0.0` |
| Status | `Approved` |
| Decision Authority | `DEC-TOOLCHAIN-001` |
| Decision Date | `2026-08-11` |
| Accountable Human | André Luiz Vieira Bonfim |
| Governed Baseline | `e7884149c80d7e8ac56c746150630e5d39f4d60f` |
| Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Classification | Category B internal governed profile; Evidence Level D decision supported by current primary technical evidence |

## Purpose and Boundary

This profile freezes the exact language, development tools, scanners, SBOM generator, CI references, integrity controls and reproducibility contract for later Implementation Entry Gate evaluation. It creates no implementation, tool-installation, dependency-installation, CI, test-execution or risk-acceptance authority.

```text
GO_LANGUAGE_VERSION=1.26.5
GO_TOOLCHAIN_VERSION=go1.26.5
RUNTIME_DEPENDENCY_POLICY=STANDARD_LIBRARY_FIRST
GO_RUNTIME_DEPENDENCIES=0
INITIAL_THIRD_PARTY_RUNTIME_DEPENDENCIES=0
GOENV=off
GOTOOLCHAIN=local
```

Automatic toolchain switching and drift are prohibited.

## Go Distribution Integrity

| Platform | Artifact | SHA-256 |
|---|---|---|
| `darwin/arm64` | `go1.26.5.darwin-arm64.tar.gz` | `efb87ff28af9a188d0536ef5d42e63dd52ba8263cd7344a993cc48dd11dedb6a` |
| `linux/amd64` | `go1.26.5.linux-amd64.tar.gz` | `5c2c3b16caefa1d968a94c1daca04a7ca301a496d9b086e17ad77bb81393f053` |

Only the official Go distribution source `https://go.dev/dl/` is approved. The archive must be verified before extraction. The selected toolchain must report the exact expected version and platform before any other verification step.

## Go Module Policy

```text
GO_MODULE_POLICY=NO_REQUIRE_NO_REPLACE_NO_VENDOR_NO_PRIVATE_MODULES
GOPROXY=https://proxy.golang.org
GOSUMDB=sum.golang.org
GOPRIVATE=
GONOSUMDB=
GONOPROXY=
GOVCS=*:off
```

The initial profile fails closed on any unexpected `require`, any `replace`, a vendor directory, a private module, an unreviewed transitive dependency, `GOSUMDB=off`, an unjustified pseudo-version, a missing required checksum or a module-state change produced by verification. Bootstrap downloads are separated from offline verification and runtime execution.

## Standard Toolchain

```text
FORMATTER=gofmt
BASELINE_STATIC_ANALYSIS=go_vet
OPTIONAL_STATIC_ANALYSIS=DEFER_STATICCHECK
TEST_FRAMEWORK=testing
FUZZ_FRAMEWORK=testing.F
```

Staticcheck is not part of the frozen baseline. Adoption requires a separately justified profile change.

## Dependency Vulnerability Scanner

```text
DEPENDENCY_SCANNER=govulncheck
DEPENDENCY_SCANNER_VERSION=v1.6.0
DEPENDENCY_SCANNER_MODULE=golang.org/x/vuln
DEPENDENCY_SCANNER_SOURCE_COMMIT=19b0bb6a272792b9afa8a6983c3e9b9a1816947f
DEPENDENCY_SCANNER_INTEGRITY=h1:FeMO9Rm/HwyduOztbvKcOw+zvDEPr4I4aQNSfevFcKY=
DEPENDENCY_SCANNER_COMMAND=govulncheck ./...
```

`govulncheck` complements SAST, testing, fuzzing, secret scanning and human review. Vulnerability-database retrieval is an allowlisted bootstrap or evidence-refresh operation, not a runtime dependency.

## SAST

```text
SAST=gosec
SAST_VERSION=2.28.0
SAST_SOURCE=https://github.com/securego/gosec
SAST_SOURCE_COMMIT=9e75c0576c9878035d4221392108d458abe10fc3
SAST_CHECKSUMS_MANIFEST_SHA256=2b494032b5ea44c5451a49a9e8c178781937eb0236cdc15087cc43e222c3a450
SAST_INSTALLATION_METHOD=PINNED_OFFICIAL_RELEASE_BINARY+SHA256_VERIFICATION+SIGSTORE_BUNDLE_VERIFICATION
```

| Platform | Artifact | SHA-256 |
|---|---|---|
| `darwin/arm64` | `gosec_2.28.0_darwin_arm64.tar.gz` | `6c4993a0ab5e3007d66c87cbcb4e3948f8000971f8eeaf3ac269cbc87a603ba4` |
| `linux/amd64` | `gosec_2.28.0_linux_amd64.tar.gz` | `d7882e505b1ff345d458bf0e893eec8019bc849f861ad73a212869540dd505ff` |

Approved future command semantics:

```text
gosec -fmt=json -out=.evidence/sast/gosec.json -track-suppressions -nosec-require-rules -nosec-require-justification -tests ./...
```

`-no-fail` is prohibited. All default rules remain enabled. Global exclusions default to none. Inline suppression requires an exact rule ID and justification and remains tracked. AI autofix and all networked AI features are prohibited. Source analysis operates offline after verified bootstrap.

CodeQL is not a required dependency of this frozen profile. Its entitlement was not authoritatively determinable and code scanning was disabled at decision time; no GitHub Code Security enablement is authorized.

## Secret Scanner

```text
SECRET_SCANNER=TruffleHog
SECRET_SCANNER_VERSION=3.96.0
SECRET_SCANNER_SOURCE=https://github.com/trufflesecurity/trufflehog
SECRET_SCANNER_SOURCE_COMMIT=6f3c981e7b77f235fd2702dd74af25fc4b72bf11
SECRET_SCANNER_COMMAND=trufflehog git file://. --no-update --no-verification --results=unverified,unknown --fail
```

| Platform | Artifact | SHA-256 |
|---|---|---|
| `darwin/arm64` | `trufflehog_3.96.0_darwin_arm64.tar.gz` | `87478306b95ca2420cfb844b7582383ac60b922e262350a0088e797f328d2e62` |
| `linux/amd64` | `trufflehog_3.96.0_linux_amd64.tar.gz` | `7105f1cd6577f058a9e39d0578f1a99c8a1e481e4d3512cd8a09acfe22a0fdc0` |

Self-update and networked credential verification are prohibited. Baseline CI receives no real credentials or production secrets.

## SBOM

```text
SBOM_TOOL=cyclonedx-gomod
SBOM_TOOL_VERSION=1.10.0
SBOM_SOURCE=https://github.com/CycloneDX/cyclonedx-gomod
SBOM_SOURCE_COMMIT=ba940a6cad4202a4a6d2f9aeef33463c0011ff5f
SBOM_FORMAT=CycloneDX_JSON
CYCLONEDX_VERSION=1.6
```

| Platform | Artifact | SHA-256 |
|---|---|---|
| `darwin/arm64` | `cyclonedx-gomod_1.10.0_darwin_arm64.tar.gz` | `2a8e887efb07ed4e36d7ead879a15e055b1e50d60e5d45e407f9f2a27db4dc46` |
| `linux/amd64` | `cyclonedx-gomod_1.10.0_linux_amd64.tar.gz` | `5cce8ae99a5181be6a610ea5ed9ca9d596937cc04dc1a8f6f6b5e462d8c9900e` |

Future governed verification must generate and validate SBOM evidence. No current SBOM or execution claim is made.

## CI and Workflow Security

```text
CI_PLATFORM=GitHub_Actions_ubuntu-24.04
PER_PR_GOOS=linux
PER_PR_GOARCH=amd64
PER_PR_GO_VERSION=1.26.5
ACTION_CHECKOUT_SHA=3d3c42e5aac5ba805825da76410c181273ba90b1
ACTION_SETUP_GO_SHA=NOT_USED
OTHER_ACTION_SHAS=NONE
GITHUB_TOKEN_PERMISSIONS=contents:read
CACHE_POLICY=DISABLED_INITIALLY
```

Floating Action references, `pull_request_target` execution of untrusted PR code, `curl | sh`, unverified binaries, write-all tokens, production credentials, scanner self-update, cached executable restoration and untrusted GitHub expression interpolation into shell are prohibited.

The hosted runner label is a moving component. Future evidence must record runner image identity and installed-image metadata. The approved Go archive and external tools must not be silently replaced by runner-provided versions.

## CGO, Build and Network Policy

```text
CGO_REQUIRED=false
CGO_POLICY=CGO_ENABLED=0
BUILD_REPRODUCIBILITY_FLAGS=-trimpath -buildvcs=false
NETWORK_POLICY=ALLOWLISTED_BOOTSTRAP;OFFLINE_VERIFICATION_AND_TESTS;RUNTIME_DENY
```

Future builds must also fix `GOOS`, `GOARCH`, build tags, `ldflags`, toolchain version and repository commit. Variable timestamps, host-specific paths and undocumented build metadata are prohibited. This is a reproducibility policy, not current bit-for-bit reproducibility evidence.

## Secret-Free Baseline

```text
real_credentials_allowed=false
production_secrets_allowed=false
.env=PROHIBITED
.env.example=SYNTHETIC_PLACEHOLDERS_ONLY
test_values=SYNTHETIC_ONLY
Git_history_secret_scan=REQUIRED
CI_secrets=NONE_FOR_BASELINE
```

## One-Command Verification Contract

```text
ONE_COMMAND_VERIFICATION=GOENV=off GOTOOLCHAIN=local go run ./cmd/verify all
```

The future implementation of this interface must orchestrate toolchain and module-policy verification, `gofmt`, `go vet`, unit/integration/security tests, AC-01 through AC-20, golden-byte vectors, bounded fuzz verification, `govulncheck`, `gosec`, TruffleHog, SBOM generation and validation, governance checks, evidence-manifest generation and no-diff verification. No command, test, source code, dependency or CI workflow is created by this profile.

## Governance Effect and Non-Effects

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
implementation_entry_gate=BLOCKED
implementation_authorized=false
source_code_authorized=false
dependency_installation_authorized=false
tool_installation_authorized=false
CI_modification_authorized=false
test_execution_authorized=false
risk_accepted=false
```
