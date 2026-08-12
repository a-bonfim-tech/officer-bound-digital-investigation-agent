package referenceslice

import (
	"strings"
	"testing"
)

func TestValidateSchemaVersion(t *testing.T) {
	if err := ValidateSchemaVersion("2.0.0"); err != nil {
		t.Fatalf("expected governed schema version to pass: %v", err)
	}
	if err := ValidateSchemaVersion("1.0.0"); err == nil || err.Code != ErrorUnsupportedSchema {
		t.Fatalf("expected unsupported schema error, got %#v", err)
	}
}

func TestIdentifierValidation(t *testing.T) {
	hex := strings.Repeat("a", 32)
	tests := []struct {
		name     string
		validate func() *ValidationError
	}{
		{"officer", func() *ValidationError { return ValidateOfficerID(OfficerID("off_" + hex)) }},
		{"case", func() *ValidationError { return ValidateCaseID(CaseID("case_" + hex)) }},
		{"grant", func() *ValidationError { return ValidateGrantID(GrantID("grant_" + hex)) }},
		{"delegation", func() *ValidationError { return ValidateDelegationID(DelegationID("dlg_" + hex)) }},
		{"request", func() *ValidationError { return ValidateRequestID(RequestID("req_" + hex)) }},
		{"decision", func() *ValidationError { return ValidateDecisionID(DecisionID("dec_" + hex)) }},
		{"connector", func() *ValidationError { return ValidateConnectorID(ConnectorID("conn_" + hex)) }},
		{"operation", func() *ValidationError { return ValidateOperationID(OperationID("op_" + hex)) }},
		{"evidence", func() *ValidationError { return ValidateEvidenceID(EvidenceID("evid_" + hex)) }},
		{"audit event", func() *ValidationError { return ValidateAuditEventID(AuditEventID("evt_" + hex)) }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.validate(); err != nil {
				t.Fatalf("valid identifier rejected: %v", err)
			}
		})
	}

	invalid := []OfficerID{
		OfficerID("case_" + hex),
		OfficerID("off_" + strings.Repeat("A", 32)),
		OfficerID("off_" + strings.Repeat("a", 31)),
		OfficerID("off_" + strings.Repeat("a", 31) + "g"),
	}
	for _, value := range invalid {
		if err := ValidateOfficerID(value); err == nil {
			t.Fatalf("invalid identifier accepted: %q", value)
		}
	}
}

func TestDigestValidation(t *testing.T) {
	if err := ValidateDigest(Digest(strings.Repeat("a", 64))); err != nil {
		t.Fatalf("valid digest rejected: %v", err)
	}
	if err := ValidateDigest(Digest(strings.Repeat("A", 64))); err == nil {
		t.Fatal("uppercase digest accepted")
	}
}

func TestNonEmptyASCIIValidation(t *testing.T) {
	tests := []struct {
		name  string
		value NonEmptyASCII
		valid bool
	}{
		{"ordinary", "value", true},
		{"spaces preserved", " value ", true},
		{"empty", "", false},
		{"too long", NonEmptyASCII(strings.Repeat("a", 129)), false},
		{"control", "a\nb", false},
		{"non ASCII", "é", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateNonEmptyASCII(test.value)
			if (err == nil) != test.valid {
				t.Fatalf("valid=%v, error=%v", test.valid, err)
			}
		})
	}
}

func TestTimestampValidation(t *testing.T) {
	tests := []struct {
		value Timestamp
		valid bool
	}{
		{"2026-08-12T13:14:15.123Z", true},
		{"2026-08-12T13:14:15.123+00:00", false},
		{"2026-08-12T13:14:15Z", false},
		{"2026-08-12T13:14:15.1234Z", false},
		{"2026-02-30T13:14:15.123Z", false},
		{"2026-08-12T13:14:60.123Z", false},
	}
	for _, test := range tests {
		err := ValidateTimestamp(test.value)
		if (err == nil) != test.valid {
			t.Fatalf("value=%q valid=%v error=%v", test.value, test.valid, err)
		}
	}
}

func TestClosedVocabularyCounts(t *testing.T) {
	if got := len(AllObjectTypes()); got != 16 {
		t.Fatalf("object type count=%d", got)
	}
	if got := len(AllErrorCodes()); got != 17 {
		t.Fatalf("error taxonomy count=%d", got)
	}
	if got := len(AllAuditEventTypes()); got != 13 {
		t.Fatalf("audit event type count=%d", got)
	}
}

func TestAuditEventTypeForError(t *testing.T) {
	expected := map[ErrorCode]AuditEventType{
		ErrorValidation: AuditValidationRejected,
		ErrorUnsupportedSchema: AuditValidationRejected,
		ErrorOversizedInput: AuditValidationRejected,
		ErrorAuthorizationDenied: AuditAuthorizationDenied,
		ErrorExpiredAuthorization: AuditAuthorizationDenied,
		ErrorNotYetValidAuthorization: AuditAuthorizationDenied,
		ErrorIdentityMismatch: AuditAuthorizationDenied,
		ErrorCaseMismatch: AuditAuthorizationDenied,
		ErrorScopeMismatch: AuditAuthorizationDenied,
		ErrorConnectorNotAllowed: AuditAuthorizationDenied,
		ErrorRevokedAuthorization: AuditRevokedGrantRejected,
		ErrorReplayDetected: AuditReplayRejected,
		ErrorStaleDecision: AuditStaleContextRejected,
		ErrorMalformedConnectorResult: AuditConnectorResultRejected,
		ErrorIntegrityFailure: AuditIntegrityFailure,
		ErrorStorageFailure: AuditStorageFailure,
		ErrorInternalInvariantFailure: AuditInternalInvariantFailure,
	}
	if len(expected) != len(AllErrorCodes()) {
		t.Fatalf("mapping count=%d", len(expected))
	}
	for code, want := range expected {
		got, ok := AuditEventTypeForError(code)
		if !ok || got != want {
			t.Fatalf("code=%q got=(%q,%v) want=%q", code, got, ok, want)
		}
	}
	if got, ok := AuditEventTypeForError("UNKNOWN"); ok || got != "" {
		t.Fatalf("unknown code mapped to (%q,%v)", got, ok)
	}
}

func TestByteExactUniquenessHelpers(t *testing.T) {
	if !UniqueNonEmptyASCII([]NonEmptyASCII{"Value", "value"}) || UniqueNonEmptyASCII([]NonEmptyASCII{"same", "same"}) {
		t.Fatal("NonEmptyASCII uniqueness is not byte exact")
	}
	if !UniqueConnectorIDs([]ConnectorID{"conn_a", "conn_A"}) || UniqueConnectorIDs([]ConnectorID{"conn_a", "conn_a"}) {
		t.Fatal("ConnectorID uniqueness is not byte exact")
	}
	if !UniqueDigests([]Digest{"aa", "AA"}) || UniqueDigests([]Digest{"aa", "aa"}) {
		t.Fatal("Digest uniqueness is not byte exact")
	}
}
