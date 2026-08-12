package referenceslice

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }
type FixedClock struct{ Value time.Time }

func (c FixedClock) Now() time.Time       { return c.Value }
func canonicalTime(t time.Time) Timestamp { return Timestamp(t.UTC().Format(timestampLayout)) }

type ReservationStore struct {
	mu        sync.Mutex
	global    map[RequestID]RequestReservationRecord
	replay    map[CaseID]map[RequestID]ReplayRecord
	Available bool
}

func NewReservationStore() *ReservationStore {
	return &ReservationStore{global: map[RequestID]RequestReservationRecord{}, replay: map[CaseID]map[RequestID]ReplayRecord{}, Available: true}
}
func (s *ReservationStore) Reserve(request RequestID, caseID CaseID, digest Digest, at Timestamp) *ValidationError {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.Available {
		return validationFailure(ErrorStorageFailure, "reservation unavailable")
	}
	if _, ok := s.global[request]; ok {
		return validationFailure(ErrorReplayDetected, "request globally used")
	}
	s.global[request] = RequestReservationRecord{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeRequestReservationRecord, RequestID: request, CaseID: caseID, CanonicalInputHash: digest, ReservedAt: at}
	if s.replay[caseID] == nil {
		s.replay[caseID] = map[RequestID]ReplayRecord{}
	}
	s.replay[caseID][request] = ReplayRecord{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeReplayRecord, CaseID: caseID, RequestID: request, CanonicalInputHash: digest, ReservedAt: at, State: ReplayReserved}
	return nil
}
func (s *ReservationStore) Transition(caseID CaseID, request RequestID, state ReplayState) *ValidationError {
	s.mu.Lock()
	defer s.mu.Unlock()
	records := s.replay[caseID]
	record, ok := records[request]
	if !ok {
		return validationFailure(ErrorStorageFailure, "case replay absent")
	}
	if record.State != ReplayReserved && record.State != ReplayConnectorInvoked {
		return validationFailure(ErrorInternalInvariantFailure, "terminal replay state")
	}
	record.State = state
	records[request] = record
	return nil
}
func (s *ReservationStore) State(caseID CaseID, request RequestID) (ReplayState, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, ok := s.replay[caseID][request]
	return r.State, ok
}

type MockConnector interface {
	Invoke(ConnectorRequest) ConnectorResult
}
type LocalMock struct {
	mu     sync.Mutex
	Calls  int
	Result ConnectorResult
}

func (m *LocalMock) Invoke(req ConnectorRequest) ConnectorResult {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Calls++
	r := m.Result
	r.RequestID = req.RequestID
	r.ConnectorID = req.ConnectorID
	r.OperationID = req.OperationID
	return r
}
func (m *LocalMock) InvocationCount() int { m.mu.Lock(); defer m.mu.Unlock(); return m.Calls }

type FlowInput struct {
	Officer       SyntheticOfficerIdentity
	Case          SyntheticCaseContext
	Grant         *AuthorizationGrant
	Delegation    *ScopedDelegation
	Context       AuthorizationContext
	Capability    *ConnectorCapability
	Request       ConnectorRequest
	Connector     MockConnector
	PolicyVersion NonEmptyASCII
	DecisionID    DecisionID
	EvidenceID    EvidenceID
	AuditEventID  AuditEventID
}
type FlowOutcome struct {
	Decision PolicyDecisionValue
	Error    *ValidationError
	Policy   *PolicyDecision
	Result   *ConnectorResult
	Evidence *EvidenceEnvelope
	Audit    []AuditEvent
}
type Engine struct {
	Clock Clock
	Store *ReservationStore
}

func (e *Engine) Execute(in FlowInput) FlowOutcome {
	now := canonicalTime(e.Clock.Now())
	deny := func(code ErrorCode, detail string) FlowOutcome {
		f := validationFailure(code, detail)
		if e.Store != nil && in.Grant != nil {
			_ = e.Store.Transition(in.Context.CaseID, in.Context.RequestID, ReplayDenied)
		}
		return FlowOutcome{Decision: PolicyDeny, Error: f, Audit: []AuditEvent{denialAudit(in, now, code)}}
	}
	if e.Store == nil {
		return deny(ErrorStorageFailure, "store absent")
	}
	if in.Grant == nil || in.Delegation == nil || in.Capability == nil || in.Connector == nil {
		return deny(ErrorAuthorizationDenied, "authorization material absent")
	}
	if in.Context.EvaluatedAt != now {
		return deny(ErrorStaleDecision, "stale authorization context")
	}
	for _, object := range []any{&in.Officer, &in.Case, in.Grant, in.Delegation, &in.Context, in.Capability, &in.Request} {
		if f := ValidateGovernedObject(object); f != nil {
			return deny(f.Code, f.Detail)
		}
	}
	digest, _, f := IntegrityDigest(&in.Context, ObjectTypeAuthorizationContext)
	if f != nil {
		return deny(f.Code, f.Detail)
	}
	if f := e.Store.Reserve(in.Context.RequestID, in.Context.CaseID, digest, now); f != nil {
		return deny(f.Code, f.Detail)
	}
	grant := in.Grant
	delegation := in.Delegation
	if grant.State == GrantRevoked {
		return deny(ErrorRevokedAuthorization, "revoked")
	}
	if grant.State == GrantExpired {
		return deny(ErrorExpiredAuthorization, "expired")
	}
	if grant.State == GrantNotYetValid {
		return deny(ErrorNotYetValidAuthorization, "not yet valid")
	}
	n := e.Clock.Now().UTC()
	nb, _ := time.Parse(timestampLayout, string(grant.NotBefore))
	ex, _ := time.Parse(timestampLayout, string(grant.ExpiresAt))
	dnb, _ := time.Parse(timestampLayout, string(delegation.NotBefore))
	dex, _ := time.Parse(timestampLayout, string(delegation.ExpiresAt))
	if n.Before(nb) || n.Before(dnb) {
		return deny(ErrorNotYetValidAuthorization, "not yet valid")
	}
	if !n.Before(ex) || !n.Before(dex) {
		return deny(ErrorExpiredAuthorization, "expired")
	}
	if in.Context.OfficerID != in.Officer.OfficerID || grant.OfficerID != in.Context.OfficerID || delegation.OfficerID != in.Context.OfficerID {
		return deny(ErrorIdentityMismatch, "officer mismatch")
	}
	if in.Context.CaseID != in.Case.CaseID || grant.CaseID != in.Context.CaseID || delegation.CaseID != in.Context.CaseID {
		return deny(ErrorCaseMismatch, "case mismatch")
	}
	if delegation.AuthorizationGrantID != grant.AuthorizationGrantID || in.Context.AuthorizationGrantID != grant.AuthorizationGrantID || in.Context.DelegationID != delegation.DelegationID {
		return deny(ErrorAuthorizationDenied, "binding mismatch")
	}
	if !containsASCII(grant.Scopes, in.Context.Action.Action) || !containsASCII(delegation.PermittedActions, in.Context.Action.Action) {
		return deny(ErrorScopeMismatch, "scope mismatch")
	}
	if !containsConnector(delegation.PermittedConnectors, in.Context.ConnectorID) || in.Capability.ConnectorID != in.Context.ConnectorID {
		return deny(ErrorConnectorNotAllowed, "connector not allowed")
	}
	if in.Request.RequestID != in.Context.RequestID || in.Request.ConnectorID != in.Context.ConnectorID || in.Request.CanonicalInputHash != digest {
		return deny(ErrorValidation, "request binding mismatch")
	}
	validUntil := ex
	if dex.Before(validUntil) {
		validUntil = dex
	}
	policy := PolicyDecision{SchemaVersion: SchemaVersion, ObjectType: ObjectTypePolicyDecision, DecisionID: in.DecisionID, Decision: PolicyAllow, OfficerID: in.Context.OfficerID, CaseID: in.Context.CaseID, AuthorizationGrantID: grant.AuthorizationGrantID, DelegationID: delegation.DelegationID, RequestID: in.Context.RequestID, ConnectorID: in.Context.ConnectorID, ConnectorCapabilityVersion: in.Capability.CapabilityVersion, PolicyVersion: in.PolicyVersion, CanonicalInputHash: digest, EvaluatedAt: now, DecisionValidUntil: canonicalTime(validUntil)}
	if f := ValidateGovernedObject(&policy); f != nil {
		return deny(f.Code, f.Detail)
	}
	_ = e.Store.Transition(in.Context.CaseID, in.Context.RequestID, ReplayConnectorInvoked)
	result := in.Connector.Invoke(in.Request)
	if f := ValidateGovernedObject(&result); f != nil || result.RequestID != in.Request.RequestID || result.ConnectorID != in.Request.ConnectorID || result.OperationID != in.Request.OperationID {
		_ = e.Store.Transition(in.Context.CaseID, in.Context.RequestID, ReplayResultRejected)
		return deny(ErrorMalformedConnectorResult, "result rejected")
	}
	for _, value := range result.ResultData {
		lower := strings.ToLower(string(value))
		if strings.Contains(lower, "authorize") || strings.Contains(lower, "officer_id") || strings.Contains(lower, "case_id") || strings.Contains(lower, "expand scope") {
			_ = e.Store.Transition(in.Context.CaseID, in.Context.RequestID, ReplayResultRejected)
			return deny(ErrorMalformedConnectorResult, "authority-like connector result")
		}
	}
	resultDigest := digestStrings(result.ResultData)
	provenance := digestStrings([]NonEmptyASCII{NonEmptyASCII(policy.DecisionID), NonEmptyASCII(in.Request.OperationID)})
	evidence := EvidenceEnvelope{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeEvidenceEnvelope, EvidenceID: in.EvidenceID, OfficerID: in.Context.OfficerID, CaseID: in.Context.CaseID, AuthorizationGrantID: grant.AuthorizationGrantID, DelegationID: delegation.DelegationID, RequestID: in.Context.RequestID, PolicyDecisionID: policy.DecisionID, ConnectorID: in.Context.ConnectorID, OperationID: in.Request.OperationID, ResultDigest: resultDigest, CreatedAt: now, ProvenanceDigest: provenance}
	if f := ValidateGovernedObject(&evidence); f != nil {
		return deny(f.Code, f.Detail)
	}
	_ = e.Store.Transition(in.Context.CaseID, in.Context.RequestID, ReplayCompleted)
	return FlowOutcome{Decision: PolicyAllow, Policy: &policy, Result: &result, Evidence: &evidence, Audit: successAudits(in, now)}
}

func digestStrings(values []NonEmptyASCII) Digest {
	h := sha256.New()
	for _, v := range values {
		h.Write([]byte(v))
		h.Write([]byte{0})
	}
	return Digest(hex.EncodeToString(h.Sum(nil)))
}
func denialAudit(in FlowInput, at Timestamp, code ErrorCode) AuditEvent {
	event, _ := AuditEventTypeForError(code)
	return AuditEvent{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeAuditEvent, AuditEventID: in.AuditEventID, EventType: event, CaseID: in.Context.CaseID, RequestID: in.Context.RequestID, OccurredAt: at, ErrorCode: &code}
}
func successAudits(in FlowInput, at Timestamp) []AuditEvent {
	events := []AuditEventType{AuditAuthorizationAllowed, AuditConnectorInvoked, AuditOperationCompleted, AuditEvidenceCreated}
	out := make([]AuditEvent, 0, len(events))
	for _, event := range events {
		out = append(out, AuditEvent{SchemaVersion: SchemaVersion, ObjectType: ObjectTypeAuditEvent, AuditEventID: in.AuditEventID, EventType: event, CaseID: in.Context.CaseID, RequestID: in.Context.RequestID, OccurredAt: at})
	}
	return out
}
