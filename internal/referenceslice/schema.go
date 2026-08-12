package referenceslice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"
)

// DecodeGovernedObject performs the schema-dispatch stage after strict raw
// validation. It is deliberately limited to the frozen OBDIA compound schema.
func DecodeGovernedObject(raw []byte, maxBytes int) (any, *ValidationError) {
	doc, failure := ValidateRawJSON(raw, RawInputPolicy{MaxBytes: maxBytes})
	if failure != nil {
		return nil, failure
	}
	var header struct {
		SchemaVersion string     `json:"schema_version"`
		ObjectType    ObjectType `json:"object_type"`
	}
	if err := json.Unmarshal(doc.Bytes(), &header); err != nil {
		return nil, validationFailure(ErrorValidation, "invalid object header")
	}
	if failure := ValidateSchemaVersion(header.SchemaVersion); failure != nil {
		return nil, failure
	}
	target := governedTarget(header.ObjectType)
	if target == nil {
		return nil, validationFailure(ErrorValidation, "unknown object type")
	}
	decoder := json.NewDecoder(bytes.NewReader(doc.Bytes()))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return nil, validationFailure(ErrorValidation, "schema decode rejected")
	}
	if _, err := decoder.Token(); err != io.EOF {
		return nil, validationFailure(ErrorValidation, "trailing JSON")
	}
	if failure := ValidateGovernedObject(target); failure != nil {
		return nil, failure
	}
	return target, nil
}

func governedTarget(objectType ObjectType) any {
	switch objectType {
	case ObjectTypeSyntheticOfficerIdentity:
		return &SyntheticOfficerIdentity{}
	case ObjectTypeSyntheticCaseContext:
		return &SyntheticCaseContext{}
	case ObjectTypeAuthorizationGrant:
		return &AuthorizationGrant{}
	case ObjectTypeScopedDelegation:
		return &ScopedDelegation{}
	case ObjectTypeRequestedAction:
		return &RequestedAction{}
	case ObjectTypeAuthorizationContext:
		return &AuthorizationContext{}
	case ObjectTypePolicyDecision:
		return &PolicyDecision{}
	case ObjectTypeConnectorCapability:
		return &ConnectorCapability{}
	case ObjectTypeConnectorRequest:
		return &ConnectorRequest{}
	case ObjectTypeConnectorResult:
		return &ConnectorResult{}
	case ObjectTypeEvidenceEnvelope:
		return &EvidenceEnvelope{}
	case ObjectTypeAuditEvent:
		return &AuditEvent{}
	case ObjectTypeProvenanceRecord:
		return &ProvenanceRecord{}
	case ObjectTypeRequestReservationRecord:
		return &RequestReservationRecord{}
	case ObjectTypeReplayRecord:
		return &ReplayRecord{}
	case ObjectTypeFixtureProvenance:
		return &FixtureProvenance{}
	default:
		return nil
	}
}

func ValidateGovernedObject(object any) *ValidationError {
	if object == nil {
		return validationFailure(ErrorValidation, "nil object")
	}
	v := reflect.ValueOf(object)
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return validationFailure(ErrorValidation, "nil object")
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return validationFailure(ErrorValidation, "object required")
	}
	if failure := validatePrimitiveFields(v); failure != nil {
		return failure
	}
	switch x := object.(type) {
	case *SyntheticOfficerIdentity:
		if x.ObjectType != ObjectTypeSyntheticOfficerIdentity || !x.Synthetic || (x.Status != OfficerActive && x.Status != OfficerDisabled) {
			return invalidSemantic()
		}
	case *SyntheticCaseContext:
		if x.ObjectType != ObjectTypeSyntheticCaseContext || x.Classification != "SYNTHETIC" || !x.Synthetic {
			return invalidSemantic()
		}
	case *AuthorizationGrant:
		if x.ObjectType != ObjectTypeAuthorizationGrant || len(x.Scopes) == 0 || !UniqueNonEmptyASCII(x.Scopes) || !grantStateValid(x.State) {
			return invalidSemantic()
		}
		if !timeOrder(x.IssuedAt, x.NotBefore, x.ExpiresAt) {
			return invalidSemantic()
		}
	case *ScopedDelegation:
		if x.ObjectType != ObjectTypeScopedDelegation || len(x.PermittedActions) == 0 || len(x.PermittedConnectors) == 0 || !UniqueNonEmptyASCII(x.PermittedActions) || !UniqueConnectorIDs(x.PermittedConnectors) || !strictBefore(x.NotBefore, x.ExpiresAt) {
			return invalidSemantic()
		}
	case *RequestedAction:
		if x.ObjectType != ObjectTypeRequestedAction || len(x.Parameters) > 32 {
			return invalidSemantic()
		}
	case *AuthorizationContext:
		if x.ObjectType != ObjectTypeAuthorizationContext || x.Action.ObjectType != ObjectTypeRequestedAction {
			return invalidSemantic()
		}
		if failure := ValidateGovernedObject(&x.Action); failure != nil {
			return failure
		}
	case *PolicyDecision:
		if x.ObjectType != ObjectTypePolicyDecision || (x.Decision != PolicyAllow && x.Decision != PolicyDeny) || !notAfter(x.EvaluatedAt, x.DecisionValidUntil) {
			return invalidSemantic()
		}
	case *ConnectorCapability:
		if x.ObjectType != ObjectTypeConnectorCapability {
			return invalidSemantic()
		}
	case *ConnectorRequest:
		if x.ObjectType != ObjectTypeConnectorRequest {
			return invalidSemantic()
		}
	case *ConnectorResult:
		if x.ObjectType != ObjectTypeConnectorResult || len(x.ResultData) > 256 || (x.Status != ConnectorResultSuccess && x.Status != ConnectorResultFailure) {
			return invalidSemantic()
		}
	case *EvidenceEnvelope:
		if x.ObjectType != ObjectTypeEvidenceEnvelope {
			return invalidSemantic()
		}
	case *AuditEvent:
		if x.ObjectType != ObjectTypeAuditEvent || !auditValid(x) {
			return invalidSemantic()
		}
	case *ProvenanceRecord:
		if x.ObjectType != ObjectTypeProvenanceRecord || len(x.SourceDigests) == 0 || !UniqueDigests(x.SourceDigests) {
			return invalidSemantic()
		}
	case *RequestReservationRecord:
		if x.ObjectType != ObjectTypeRequestReservationRecord {
			return invalidSemantic()
		}
	case *ReplayRecord:
		if x.ObjectType != ObjectTypeReplayRecord || !replayStateValid(x.State) {
			return invalidSemantic()
		}
	case *FixtureProvenance:
		if x.ObjectType != ObjectTypeFixtureProvenance || x.Classification != "SYNTHETIC_FIXTURE" || !x.Synthetic || x.DerivedFromRealPersonData || x.DerivedFromLiveCaseData || x.DerivedFromProductionExport || x.DerivedFromLiveConnector || x.ContainsRealCredentials || !fixtureStateValid(x) {
			return invalidSemantic()
		}
	default:
		return validationFailure(ErrorValidation, "unsupported Go object")
	}
	return nil
}

func invalidSemantic() *ValidationError {
	return validationFailure(ErrorValidation, "governed semantic validation rejected")
}

func validatePrimitiveFields(v reflect.Value) *ValidationError {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f, sf := v.Field(i), t.Field(i)
		if sf.Name == "SchemaVersion" {
			if ValidateSchemaVersion(f.String()) != nil {
				return validationFailure(ErrorUnsupportedSchema, "unsupported schema")
			}
			continue
		}
		if f.Kind() == reflect.Pointer {
			continue
		}
		if f.Kind() == reflect.Slice {
			if f.IsNil() {
				return invalidSemantic()
			}
			for j := 0; j < f.Len(); j++ {
				if failure := validateNamed(f.Index(j)); failure != nil {
					return failure
				}
			}
			continue
		}
		if failure := validateNamed(f); failure != nil {
			return failure
		}
	}
	return nil
}

func validateNamed(v reflect.Value) *ValidationError {
	if v.Kind() == reflect.Struct {
		return validatePrimitiveFields(v)
	}
	if v.Kind() != reflect.String {
		return nil
	}
	s := v.String()
	switch v.Type().Name() {
	case "OfficerID":
		return ValidateOfficerID(OfficerID(s))
	case "CaseID":
		return ValidateCaseID(CaseID(s))
	case "GrantID":
		return ValidateGrantID(GrantID(s))
	case "DelegationID":
		return ValidateDelegationID(DelegationID(s))
	case "RequestID":
		return ValidateRequestID(RequestID(s))
	case "DecisionID":
		return ValidateDecisionID(DecisionID(s))
	case "ConnectorID":
		return ValidateConnectorID(ConnectorID(s))
	case "OperationID":
		return ValidateOperationID(OperationID(s))
	case "EvidenceID":
		return ValidateEvidenceID(EvidenceID(s))
	case "AuditEventID":
		return ValidateAuditEventID(AuditEventID(s))
	case "Digest":
		return ValidateDigest(Digest(s))
	case "Timestamp":
		return ValidateTimestamp(Timestamp(s))
	case "NonEmptyASCII":
		return ValidateNonEmptyASCII(NonEmptyASCII(s))
	}
	return nil
}

func parseTime(v Timestamp) (int64, bool) {
	t, e := timeParse(string(v))
	if e != nil {
		return 0, false
	}
	return t.UnixMilli(), true
}
func strictBefore(a, b Timestamp) bool {
	x, ok := parseTime(a)
	if !ok {
		return false
	}
	y, ok := parseTime(b)
	return ok && x < y
}
func notAfter(a, b Timestamp) bool {
	x, ok := parseTime(a)
	if !ok {
		return false
	}
	y, ok := parseTime(b)
	return ok && x <= y
}
func timeOrder(issued, notBefore, expires Timestamp) bool {
	return notAfter(issued, notBefore) && strictBefore(notBefore, expires)
}
func timeParse(v string) (time.Time, error) { return time.Parse(timestampLayout, v) }

func grantStateValid(v AuthorizationGrantState) bool {
	return v == GrantActive || v == GrantExpired || v == GrantRevoked || v == GrantNotYetValid
}
func replayStateValid(v ReplayState) bool {
	return v == ReplayReserved || v == ReplayDenied || v == ReplayConnectorInvoked || v == ReplayResultRejected || v == ReplayCompleted || v == ReplayIndeterminateOutcome
}
func fixtureStateValid(x *FixtureProvenance) bool {
	if x.CreationMethod != FixtureManualInvention && x.CreationMethod != FixtureSyntheticGenerator {
		return false
	}
	if x.ReviewStatus != FixtureReviewPending && x.ReviewStatus != FixtureReviewApproved && x.ReviewStatus != FixtureReviewRejected {
		return false
	}
	if x.ReviewStatus != FixtureReviewPending {
		return x.Reviewer != nil && x.ReviewTimestamp != nil
	}
	return x.Reviewer == nil && x.ReviewTimestamp == nil
}
func auditValid(x *AuditEvent) bool {
	success := x.EventType == AuditAuthorizationAllowed || x.EventType == AuditConnectorInvoked || x.EventType == AuditOperationCompleted || x.EventType == AuditEvidenceCreated
	if success {
		return x.ErrorCode == nil
	}
	if x.ErrorCode == nil {
		return false
	}
	expected, ok := AuditEventTypeForError(*x.ErrorCode)
	return ok && expected == x.EventType
}

func containsASCII(values []NonEmptyASCII, target NonEmptyASCII) bool {
	for _, v := range values {
		if v == target {
			return true
		}
	}
	return false
}
func containsConnector(values []ConnectorID, target ConnectorID) bool {
	for _, v := range values {
		if v == target {
			return true
		}
	}
	return false
}

func minimizedError(err error) string {
	if err == nil {
		return ""
	}
	return strings.TrimSpace(fmt.Sprint(err))
}
