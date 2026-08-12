# Security Policy

## Scope

The Officer-Bound Digital Investigation Agent (OBDIA) repository is an independent cybersecurity research and engineering project. The current release scope is a **bounded, synthetic, local/non-production reference slice**.

Security reports are welcome for vulnerabilities in the repository's source code, validation logic, authorization boundaries, evidence/integrity mechanisms, CI configuration and supply-chain controls.

The following are outside the current operational scope and should not be treated as deployed product surfaces:

- real investigations or real-person targeting;
- production law-enforcement infrastructure;
- live investigative connectors;
- production credentials or secrets;
- live Dark Web/Tor operations;
- uncontrolled external scraping or investigative APIs;
- real wallets, private keys or blockchain transactions;
- production deployment.

## Security invariant

The bounded implementation is governed by this invariant:

```text
NO TOOL OR CONNECTOR EXECUTION BEFORE A VALID POSITIVE AUTHORIZATION DECISION
```

A report is particularly valuable if it demonstrates a way to bypass or weaken any of the following properties:

- default-deny authorization;
- officer, case, scope, delegation or connector binding;
- global request-ID uniqueness;
- replay protection;
- deterministic authorization validity windows;
- exact `ALLOW` gating before connector invocation;
- fail-closed handling of malformed, stale, revoked, unknown or indeterminate states;
- canonicalization or integrity verification;
- evidence/audit/provenance integrity;
- synthetic-data and local-mock containment;
- CI or supply-chain integrity.

## Reporting a vulnerability

Please **do not disclose suspected vulnerabilities, exploit details, credentials or sensitive data in a public issue**.

Preferred reporting path:

1. Use GitHub's private vulnerability-reporting / security-advisory interface for this repository when it is available.
2. If private vulnerability reporting is not available, contact the repository owner through GitHub and request a private channel before sharing technical exploit details.
3. Include the affected commit or release, impacted files/components, reproduction conditions, expected versus observed behavior, and a minimal proof of concept that uses synthetic data only.

Do not include real personal data, real case information, production credentials, stolen data, or interaction with uncontrolled criminal infrastructure in a report.

## Coordinated disclosure

Reports will be evaluated against the repository's governed security and publication boundaries. Valid findings should be remediated and regression-tested before public technical disclosure when premature disclosure would create unnecessary abuse risk.

A security report does not grant authorization to test third-party systems, law-enforcement systems, real persons, external services or infrastructure not owned by the reporter.

## Security evidence

The bounded reference-slice evidence manifest is retained at:

[`docs/evidence/reference-slice/evidence-manifest.json`](docs/evidence/reference-slice/evidence-manifest.json)

It records the scoped results of formatting, `go vet`, tests, acceptance criteria, fuzzing, vulnerability analysis, static security analysis, secret scanning, SBOM generation and CI. These results are bounded evidence; they are not a claim of production security, independent certification or formal compliance.

## Supported versions

Security fixes are expected to target the current `main` branch and the latest published `v0.x` release unless a release note explicitly states otherwise.
