# DEC-TOOLCHAIN-002 — Exact Go 1.26.5 Local Bootstrap Execution Authorization

| Field | Value |
|---|---|
| Decision ID | `DEC-TOOLCHAIN-002` |
| Decision Type | `EXACT_GO_1_26_5_LOCAL_BOOTSTRAP_EXECUTION_AUTHORIZATION` |
| Decision Date | `2026-08-12` |
| Status | `Effective` |
| Accountable Human | André Luiz Vieira Bonfim |
| Scope | `ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE` |
| Decision Basis | `CORE_VALIDATION_EXECUTION_BLOCKED_TOOLCHAIN_UNAVAILABLE` |

## Retained Human Disposition

The Accountable Human approves an exact, integrity-verified, user-scoped bootstrap of Go `1.26.5` for Darwin/arm64 solely to enable bounded local implementation validation. Retention creates repository authority but does not execute the bootstrap.

```text
decision_id=DEC-TOOLCHAIN-002
decision_type=EXACT_GO_1_26_5_LOCAL_BOOTSTRAP_EXECUTION_AUTHORIZATION
decision_date=2026-08-12
status=Effective
accountable_human=André Luiz Vieira Bonfim
scope=ADR_0001_BOUNDED_SYNTHETIC_REFERENCE_SLICE
decision_basis=CORE_VALIDATION_EXECUTION_BLOCKED_TOOLCHAIN_UNAVAILABLE
DEC_TOOLCHAIN_002_created=true
DEC_TOOLCHAIN_002_RETAINED=true
DEC_TOOLCHAIN_002_HUMAN_DISPOSITION=APPROVED
REPOSITORY_GO_1_26_5_BOOTSTRAP_AUTHORIZED=true
Go_1_26_5_bootstrap_authorized=true
bootstrap_network_authorized=true
bootstrap_installation_authorized=true
bootstrap_scope=EXACT_GO_1_26_5_DARWIN_ARM64_ONLY
bootstrap_purpose=BOUNDED_LOCAL_IMPLEMENTATION_VALIDATION_ONLY
bootstrap_authority_retained=true
bootstrap_execution_ready=false
```

## Host and Authorized Artifact

```text
reviewed_host_kernel=Darwin
reviewed_host_arch=arm64
reviewed_host_os_version=26.6.1
authorized_platform_match=true
BOOTSTRAP_PLATFORM=darwin/arm64
GO_BOOTSTRAP_VERSION=1.26.5
GO_BOOTSTRAP_TOOLCHAIN=go1.26.5
GO_BOOTSTRAP_GOOS=darwin
GO_BOOTSTRAP_GOARCH=arm64
GO_BOOTSTRAP_ARTIFACT=go1.26.5.darwin-arm64.tar.gz
GO_BOOTSTRAP_OFFICIAL_URL=https://go.dev/dl/go1.26.5.darwin-arm64.tar.gz
GO_BOOTSTRAP_EXPECTED_SHA256=efb87ff28af9a188d0536ef5d42e63dd52ba8263cd7344a993cc48dd11dedb6a
ARTIFACT_SOURCE_AUTHORITY=OFFICIAL_GO_DISTRIBUTION
```

No version substitution, package-manager artifact, third-party mirror, source build or latest-version resolution is permitted.

## Network and Installation Boundary

Bootstrap network authority is limited to HTTPS retrieval of the exact artifact from `go.dev`, including only official delivery redirects for that artifact. It does not authorize application runtime network, module downloads, GOPROXY/GOSUMDB traffic, `go get`, `go install`, auto-switching, security-tool download, CI acquisition, live connectors or investigative network use.

```text
runtime_network=DENY
BOOTSTRAP_NETWORK_AUTHORITY_IS_NOT_RUNTIME_NETWORK_AUTHORITY=true
OBDIA_TOOLCHAIN_ROOT=$HOME/.local/share/obdia/toolchains
OBDIA_GO_BOOTSTRAP_ARTIFACT_DIR=$HOME/.local/share/obdia/toolchains/artifacts
OBDIA_GO_INSTALL_PARENT=$HOME/.local/share/obdia/toolchains/installs/go1.26.5-darwin-arm64
OBDIA_GO_GOROOT=$HOME/.local/share/obdia/toolchains/installs/go1.26.5-darwin-arm64/go
installation_scope=USER_SCOPED_ONLY
privileged_execution_authorized=false
system_installation_authorized=false
permanent_PATH_mutation_authorized=false
shell_profile_mutation_authorized=false
```

No installation into `/usr/local/go`, `/opt/homebrew`, `/usr/bin` or `/usr/local/bin` is authorized.

## Fail-Closed Bootstrap Procedure

If the installation parent already exists, execution must stop without deletion, overwrite or content merging.

```text
BOOTSTRAP_ORDER=DOWNLOAD_THEN_SHA256_VERIFY_THEN_ARCHIVE_MEMBER_SAFETY_INSPECTION_THEN_EXTRACT_THEN_VERSION_VERIFY_THEN_GOROOT_VERIFY
VERIFY_AFTER_EXTRACTION_PROHIBITED=true
archive_top_level_directory=go
archive_absolute_member_detected=false
archive_parent_traversal_detected=false
archive_member_escape_detected=false
```

An observed digest different from `efb87ff28af9a188d0536ef5d42e63dd52ba8263cd7344a993cc48dd11dedb6a` requires `STOP`, `DO_NOT_EXTRACT` and `DO_NOT_EXECUTE_DOWNLOADED_CONTENT`.

Future execution must invoke the absolute `go` and `gofmt` binaries beneath the governed GOROOT and verify:

```text
go_version=go version go1.26.5 darwin/arm64
GOVERSION=go1.26.5
GOOS=darwin
GOARCH=arm64
GOROOT=$HOME/.local/share/obdia/toolchains/installs/go1.26.5-darwin-arm64/go
GOTOOLCHAIN=local
GOENV=off
GOPROXY=off
GOSUMDB=off
GOVCS=*:off
CGO_ENABLED=0
GOFLAGS=-mod=readonly
toolchain_auto_switching_authorized=false
toolchain_auto_download_authorized=false
module_download_authorized=false
```

## Existing Implementation Topology and Required Forward Sync

```text
IMPLEMENTATION_BRANCH_ALREADY_INITIALIZED=true
implementation_branch=impl/bounded-synthetic-reference-slice-2026-08
IMPLEMENTATION_BRANCH_INITIAL_BASELINE=a5dcf5684c8d87ba19011509c6bfd216af4df02d
implementation_pre_governance_sync_HEAD=814fb98bac3f8afbc23fbe0121584951e7929efd
implementation_PR=49
implementation_PR_state=OPEN_DRAFT
DEC_TOOLCHAIN_002_RETENTION_WILL_ADVANCE_GOVERNANCE_BRANCH=true
governance_advance_is_expected_post_branch_initialization=true
baseline_reconciliation_required_for_historical_branch_initialization=false
IMPLEMENTATION_BRANCH_GOVERNANCE_SYNC_REQUIRED=true
IMPLEMENTATION_BRANCH_GOVERNANCE_SYNC_COMPLETE=false
post_initialization_governance_sync_required=true
future_sync_method=SIGNED_MERGE_COMMIT_FROM_CURRENT_GOVERNANCE_BRANCH
rebase_authorized=false
force_push_authorized=false
reset_to_governance_HEAD_authorized=false
governance_decision_cherry_pick_authorized=false
history_rewrite_authorized=false
BOOTSTRAP_EXECUTION_GATE=PENDING_IMPLEMENTATION_BRANCH_GOVERNANCE_SYNC
bootstrap_execution_ready=false
```

The future signed merge must preserve `a11a6c8117bf13f7a98b078a4cd6bae57fd2d68c` and `814fb98bac3f8afbc23fbe0121584951e7929efd` as ancestors. Bootstrap execution additionally requires PR #49 to remain open/Draft and the implementation worktree to be clean.

## Preserved Authority and Non-Effects

```text
DEC_IMPL_001_RETAINED=true
DEC_IMPL_002_RETAINED=true
bounded_implementation_authorized=true
controlled_evidence_generation_authorized=true
local_test_execution_authorized=true
implementation_authorization_scope=SYNTHETIC_REFERENCE_SLICE_ONLY
authorization_scope_expanded=false
general_or_unbounded_implementation_authorized=false
production_implementation_authorized=false
investigative_use_authorized=false
initial_third_party_runtime_dependencies=0
third_party_test_dependencies=0
runtime_dependency_policy=STANDARD_LIBRARY_FIRST
dependency_download_authorized=false
security_tool_execution_authorized=false
CI_creation_authorized=false
CI_modification_authorized=false
CI_execution_authorized=false
download_performed=false
archive_downloaded=false
archive_sha256_verified=false
archive_extracted=false
installation_performed=false
toolchain_executed=false
Go_1_26_5_bootstrap_performed=false
network_bootstrap_performed=false
gofmt_executed=false
go_vet_executed=false
go_test_executed=false
testing_F_executed=false
tests_executed=false
risk_accepted=false
```

`DEC-TOOLCHAIN-002` authorizes a future bounded bootstrap only. It is not bootstrap, test, operating, effectiveness, risk-acceptance, production-readiness, CI, release or publication evidence.
