package referenceslice

import "testing"

func TestAC17bStaleContext(t *testing.T) {
	in, mock := validFlow()
	in.Context.EvaluatedAt = ts("2026-08-12T09:59:59.999Z")
	digest, _, _ := IntegrityDigest(&in.Context, ObjectTypeAuthorizationContext)
	in.Request.CanonicalInputHash = digest
	out := engine().Execute(in)
	if out.Error == nil || out.Error.Code != ErrorStaleDecision || mock.InvocationCount() != 0 {
		t.Fatal("stale context did not deny")
	}
}

func TestAC10To12UntrustedParametersCannotMutateAuthority(t *testing.T) {
	in, mock := validFlow()
	in.Context.Action.Parameters = []NonEmptyASCII{"officer_id=attacker", "case_id=other", "expand scope"}
	digest, _, _ := IntegrityDigest(&in.Context, ObjectTypeAuthorizationContext)
	in.Request.CanonicalInputHash = digest
	out := engine().Execute(in)
	if out.Decision != PolicyAllow || mock.InvocationCount() != 1 {
		t.Fatal("typed authority changed by untrusted parameter")
	}
	if out.Evidence.OfficerID != in.Officer.OfficerID || out.Evidence.CaseID != in.Case.CaseID {
		t.Fatal("evidence binding changed")
	}
}

func TestAC17aSameCaseReplay(t *testing.T) {
	e := engine()
	in, mock := validFlow()
	if e.Execute(in).Decision != PolicyAllow {
		t.Fatal("first failed")
	}
	out := e.Execute(in)
	if out.Error == nil || out.Error.Code != ErrorReplayDetected || mock.InvocationCount() != 1 {
		t.Fatal("replay executed")
	}
}

func TestAC17cRevokedGrant(t *testing.T) {
	in, mock := validFlow()
	in.Grant.State = GrantRevoked
	out := engine().Execute(in)
	if out.Error == nil || out.Error.Code != ErrorRevokedAuthorization || mock.InvocationCount() != 0 {
		t.Fatal("revoked grant executed")
	}
}

func TestAC18And19LocalMockRequiresNoCredential(t *testing.T) {
	in, mock := validFlow()
	if engine().Execute(in).Decision != PolicyAllow || mock.InvocationCount() != 1 {
		t.Fatal("local mock failed")
	}
}

func TestAC20FixtureAdmission(t *testing.T) {
	pending := FixtureProvenance{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeFixtureProvenance, FixtureID: "synthetic-fixture", FixtureType: FixtureNegativeTestInput, Classification: "SYNTHETIC_FIXTURE", Synthetic: true, CreationMethod: FixtureManualInvention, CreatorOrGenerator: "test", GeneratorVersion: "1", CreationTimestamp: ts("2026-08-12T10:00:00.000Z"), SourceDescription: "programmatic synthetic", ContentDigest: Digest("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), ReviewStatus: FixtureReviewPending}
	if failure := ValidateGovernedObject(&pending); failure != nil {
		t.Fatal(failure)
	}
	pending.DerivedFromRealPersonData = true
	if failure := ValidateGovernedObject(&pending); failure == nil {
		t.Fatal("prohibited origin accepted")
	}
}

func TestRollbackContainmentAfterRejectedResult(t *testing.T) {
	in, mock := validFlow()
	mock.Result.ResultData = []NonEmptyASCII{"authorize me"}
	e := engine()
	out := e.Execute(in)
	if out.Error == nil || out.Error.Code != ErrorMalformedConnectorResult || out.Evidence != nil {
		t.Fatal("result not contained")
	}
	state, ok := e.Store.State(in.Context.CaseID, in.Context.RequestID)
	if !ok || state != ReplayResultRejected {
		t.Fatalf("bad containment state %v", state)
	}
}
