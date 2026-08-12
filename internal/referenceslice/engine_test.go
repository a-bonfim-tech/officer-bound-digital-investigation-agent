package referenceslice

import (
	"bytes"
	"encoding/hex"
	"testing"
	"time"
)

const testHex = "00000000000000000000000000000001"

func ts(v string) Timestamp { return Timestamp(v) }

func validFlow() (FlowInput, *LocalMock) {
	now := ts("2026-08-12T10:00:00.000Z")
	off := OfficerID("off_" + testHex)
	caseID := CaseID("case_" + testHex)
	grant := GrantID("grant_" + testHex)
	dlg := DelegationID("dlg_" + testHex)
	req := RequestID("req_" + testHex)
	conn := ConnectorID("conn_" + testHex)
	op := OperationID("op_" + testHex)
	action := RequestedAction{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeRequestedAction, Action: "READ_SYNTHETIC", Target: "synthetic-target", Parameters: []NonEmptyASCII{}}
	ctx := AuthorizationContext{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeAuthorizationContext, OfficerID: off, CaseID: caseID, AuthorizationGrantID: grant, DelegationID: dlg, RequestID: req, ConnectorID: conn, Action: action, EvaluatedAt: now}
	digest, _, _ := IntegrityDigest(&ctx, ObjectTypeAuthorizationContext)
	mock := &LocalMock{Result: ConnectorResult{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeConnectorResult, Status: ConnectorResultSuccess, ResultData: []NonEmptyASCII{"synthetic-result"}}}
	return FlowInput{Officer: SyntheticOfficerIdentity{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeSyntheticOfficerIdentity, OfficerID: off, InstitutionID: "synthetic-lab", Synthetic: true, Status: OfficerActive}, Case: SyntheticCaseContext{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeSyntheticCaseContext, CaseID: caseID, Purpose: "bounded-test", Classification: "SYNTHETIC", Synthetic: true}, Grant: &AuthorizationGrant{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeAuthorizationGrant, AuthorizationGrantID: grant, OfficerID: off, CaseID: caseID, Scopes: []NonEmptyASCII{"READ_SYNTHETIC"}, IssuedAt: ts("2026-08-12T09:00:00.000Z"), NotBefore: ts("2026-08-12T09:00:00.000Z"), ExpiresAt: ts("2026-08-12T11:00:00.000Z"), State: GrantActive}, Delegation: &ScopedDelegation{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeScopedDelegation, DelegationID: dlg, AuthorizationGrantID: grant, OfficerID: off, CaseID: caseID, PermittedActions: []NonEmptyASCII{"READ_SYNTHETIC"}, PermittedConnectors: []ConnectorID{conn}, NotBefore: ts("2026-08-12T09:00:00.000Z"), ExpiresAt: ts("2026-08-12T11:00:00.000Z")}, Context: ctx, Capability: &ConnectorCapability{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeConnectorCapability, ConnectorID: conn, CapabilityVersion: "1", Operation: "READ_SYNTHETIC", RequestSchema: "2.0.0", ResultSchema: "2.0.0"}, Request: ConnectorRequest{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeConnectorRequest, RequestID: req, DecisionID: DecisionID("dec_" + testHex), ConnectorID: conn, OperationID: op, CanonicalInputHash: digest}, Connector: mock, PolicyVersion: "1", DecisionID: DecisionID("dec_" + testHex), EvidenceID: EvidenceID("evid_" + testHex), AuditEventID: AuditEventID("evt_" + testHex)}, mock
}

func engine() *Engine {
	now, _ := time.Parse(timestampLayout, "2026-08-12T10:00:00.000Z")
	return &Engine{Clock: FixedClock{Value: now}, Store: NewReservationStore()}
}

func TestAC01AC14AC15AC16ValidFlowEvidenceAndAudit(t *testing.T) {
	in, m := validFlow()
	out := engine().Execute(in)
	if out.Decision != PolicyAllow || out.Error != nil || out.Evidence == nil || m.InvocationCount() != 1 {
		t.Fatalf("valid flow denied: %+v", out.Error)
	}
	if len(out.Audit) == 0 {
		t.Fatal("audit absent")
	}
	for _, event := range out.Audit {
		if event.CaseID != in.Case.CaseID || event.RequestID != in.Request.RequestID {
			t.Fatal("audit does not bind governed case and request context")
		}
	}
	digest, _, f := IntegrityDigest(out.Evidence, ObjectTypeEvidenceEnvelope)
	if f != nil || !VerifyIntegrity(out.Evidence, ObjectTypeEvidenceEnvelope, digest) {
		t.Fatal("evidence integrity failed")
	}
	out.Evidence.ResultDigest = Digest("0" + string(out.Evidence.ResultDigest)[1:])
	if VerifyIntegrity(out.Evidence, ObjectTypeEvidenceEnvelope, digest) {
		t.Fatal("tampering not detected")
	}
}

func TestDenialMatrixZeroConnector(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*FlowInput)
		code   ErrorCode
	}{{"AC02_missing_authorization", func(i *FlowInput) { i.Grant = nil }, ErrorAuthorizationDenied}, {"AC03_expired_authorization", func(i *FlowInput) { i.Grant.State = GrantExpired }, ErrorExpiredAuthorization}, {"AC17c_revoked_authorization", func(i *FlowInput) { i.Grant.State = GrantRevoked }, ErrorRevokedAuthorization}, {"AC04_officer_mismatch", func(i *FlowInput) { i.Context.OfficerID = OfficerID("off_00000000000000000000000000000002") }, ErrorIdentityMismatch}, {"AC05_case_mismatch", func(i *FlowInput) { i.Context.CaseID = CaseID("case_00000000000000000000000000000002") }, ErrorCaseMismatch}, {"AC06_scope_mismatch", func(i *FlowInput) { i.Context.Action.Action = "DENIED_ACTION" }, ErrorScopeMismatch}, {"AC07_connector_not_allowlisted", func(i *FlowInput) { i.Context.ConnectorID = ConnectorID("conn_00000000000000000000000000000002") }, ErrorConnectorNotAllowed}, {"AC08_unknown_policy", func(i *FlowInput) { i.PolicyVersion = "unknown" }, ErrorAuthorizationDenied}, {"AC09_malformed_context", func(i *FlowInput) { i.Context.ObjectType = ObjectTypePolicyDecision }, ErrorValidation}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, m := validFlow()
			tt.mutate(&in)
			out := engine().Execute(in)
			if out.Decision != PolicyDeny || out.Error == nil || out.Error.Code != tt.code || m.InvocationCount() != 0 || out.Evidence != nil {
				t.Fatalf("bad denial: %+v calls=%d", out, m.InvocationCount())
			}
		})
	}
}

func TestAC13ReplayGlobalAcrossCases(t *testing.T) {
	e := engine()
	in, m := validFlow()
	if out := e.Execute(in); out.Decision != PolicyAllow {
		t.Fatal("first denied")
	}
	in2, m2 := validFlow()
	in2.Case.CaseID = CaseID("case_00000000000000000000000000000002")
	in2.Context.CaseID = in2.Case.CaseID
	in2.Grant.CaseID = in2.Case.CaseID
	in2.Delegation.CaseID = in2.Case.CaseID
	digest, _, _ := IntegrityDigest(&in2.Context, ObjectTypeAuthorizationContext)
	in2.Request.CanonicalInputHash = digest
	out := e.Execute(in2)
	if out.Error == nil || out.Error.Code != ErrorReplayDetected || m2.InvocationCount() != 0 || m.InvocationCount() != 1 {
		t.Fatal("cross-case replay not denied")
	}
}

func TestStorageFailureFailsClosed(t *testing.T) {
	in, m := validFlow()
	e := engine()
	e.Store.Available = false
	out := e.Execute(in)
	if out.Error == nil || out.Error.Code != ErrorStorageFailure || m.InvocationCount() != 0 {
		t.Fatal("storage did not fail closed")
	}
}

func TestMalformedConnectorResultRejected(t *testing.T) {
	in, m := validFlow()
	m.Result.ObjectType = ObjectTypePolicyDecision
	out := engine().Execute(in)
	if out.Error == nil || out.Error.Code != ErrorMalformedConnectorResult || out.Evidence != nil || m.InvocationCount() != 1 {
		t.Fatal("malformed result accepted")
	}
}

func TestCanonicalGoldenAndPreimage(t *testing.T) {
	off := &SyntheticOfficerIdentity{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeSyntheticOfficerIdentity, OfficerID: OfficerID("off_" + testHex), InstitutionID: "lab", Synthetic: true, Status: OfficerActive}
	canonical, f := Canonicalize(off)
	if f != nil {
		t.Fatal(f)
	}
	want := `{"institution_id":"lab","object_type":"SyntheticOfficerIdentity","officer_id":"off_00000000000000000000000000000001","schema_version":"2.0.0","status":"ACTIVE","synthetic":true}`
	if string(canonical) != want {
		t.Fatalf("canonical mismatch\n%s", canonical)
	}
	pre, f := DomainSeparatedPreimage(ObjectTypeSyntheticOfficerIdentity, SchemaVersion, canonical)
	if f != nil {
		t.Fatal(f)
	}
	if !bytes.Equal(pre[:6], []byte{'O', 'B', 'D', 'I', 'A', 0}) {
		t.Fatal("domain prefix")
	}
	if hex.EncodeToString(pre) == "" {
		t.Fatal("preimage absent")
	}
}

func TestStrictDecodeUnknownAndSchema(t *testing.T) {
	valid := `{"schema_version":"2.0.0","object_type":"SyntheticOfficerIdentity","officer_id":"off_00000000000000000000000000000001","institution_id":"lab","synthetic":true,"status":"ACTIVE"}`
	if _, f := DecodeGovernedObject([]byte(valid), 4096); f != nil {
		t.Fatal(f)
	}
	unknown := valid[:len(valid)-1] + `,"unknown":true}`
	if _, f := DecodeGovernedObject([]byte(unknown), 4096); f == nil {
		t.Fatal("unknown field accepted")
	}
	duplicate := `{"schema_version":"2.0.0","schema_version":"2.0.0","object_type":"SyntheticOfficerIdentity"}`
	if _, f := DecodeGovernedObject([]byte(duplicate), 4096); f == nil {
		t.Fatal("duplicate accepted")
	}
}

func FuzzStrictRawJSON(f *testing.F) {
	f.Add([]byte(`{"a":1}`))
	f.Fuzz(func(t *testing.T, raw []byte) { _, _ = ValidateRawJSON(raw, RawInputPolicy{MaxBytes: 4096}) })
}
