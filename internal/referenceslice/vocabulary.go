package referenceslice

const SchemaVersion = "2.0.0"

type ObjectType string

const (
	ObjectTypeSyntheticOfficerIdentity  ObjectType = "SyntheticOfficerIdentity"
	ObjectTypeSyntheticCaseContext      ObjectType = "SyntheticCaseContext"
	ObjectTypeAuthorizationGrant        ObjectType = "AuthorizationGrant"
	ObjectTypeScopedDelegation          ObjectType = "ScopedDelegation"
	ObjectTypeRequestedAction           ObjectType = "RequestedAction"
	ObjectTypeAuthorizationContext      ObjectType = "AuthorizationContext"
	ObjectTypePolicyDecision            ObjectType = "PolicyDecision"
	ObjectTypeConnectorCapability       ObjectType = "ConnectorCapability"
	ObjectTypeConnectorRequest          ObjectType = "ConnectorRequest"
	ObjectTypeConnectorResult           ObjectType = "ConnectorResult"
	ObjectTypeEvidenceEnvelope          ObjectType = "EvidenceEnvelope"
	ObjectTypeAuditEvent                ObjectType = "AuditEvent"
	ObjectTypeProvenanceRecord          ObjectType = "ProvenanceRecord"
	ObjectTypeRequestReservationRecord  ObjectType = "RequestReservationRecord"
	ObjectTypeReplayRecord              ObjectType = "ReplayRecord"
	ObjectTypeFixtureProvenance         ObjectType = "FixtureProvenance"
)

var objectTypes = [...]ObjectType{
	ObjectTypeSyntheticOfficerIdentity,
	ObjectTypeSyntheticCaseContext,
	ObjectTypeAuthorizationGrant,
	ObjectTypeScopedDelegation,
	ObjectTypeRequestedAction,
	ObjectTypeAuthorizationContext,
	ObjectTypePolicyDecision,
	ObjectTypeConnectorCapability,
	ObjectTypeConnectorRequest,
	ObjectTypeConnectorResult,
	ObjectTypeEvidenceEnvelope,
	ObjectTypeAuditEvent,
	ObjectTypeProvenanceRecord,
	ObjectTypeRequestReservationRecord,
	ObjectTypeReplayRecord,
	ObjectTypeFixtureProvenance,
}

func AllObjectTypes() []ObjectType {
	return append([]ObjectType(nil), objectTypes[:]...)
}

type ErrorCode string

const (
	ErrorValidation                 ErrorCode = "VALIDATION_ERROR"
	ErrorAuthorizationDenied        ErrorCode = "AUTHORIZATION_DENIED"
	ErrorExpiredAuthorization       ErrorCode = "EXPIRED_AUTHORIZATION"
	ErrorNotYetValidAuthorization   ErrorCode = "NOT_YET_VALID_AUTHORIZATION"
	ErrorRevokedAuthorization       ErrorCode = "REVOKED_AUTHORIZATION"
	ErrorIdentityMismatch           ErrorCode = "IDENTITY_MISMATCH"
	ErrorCaseMismatch               ErrorCode = "CASE_MISMATCH"
	ErrorScopeMismatch              ErrorCode = "SCOPE_MISMATCH"
	ErrorConnectorNotAllowed        ErrorCode = "CONNECTOR_NOT_ALLOWED"
	ErrorReplayDetected             ErrorCode = "REPLAY_DETECTED"
	ErrorStaleDecision              ErrorCode = "STALE_DECISION"
	ErrorMalformedConnectorResult   ErrorCode = "MALFORMED_CONNECTOR_RESULT"
	ErrorIntegrityFailure           ErrorCode = "INTEGRITY_FAILURE"
	ErrorStorageFailure             ErrorCode = "STORAGE_FAILURE"
	ErrorInternalInvariantFailure   ErrorCode = "INTERNAL_INVARIANT_FAILURE"
	ErrorUnsupportedSchema          ErrorCode = "UNSUPPORTED_SCHEMA"
	ErrorOversizedInput             ErrorCode = "OVERSIZED_INPUT"
)

var errorCodes = [...]ErrorCode{
	ErrorValidation,
	ErrorAuthorizationDenied,
	ErrorExpiredAuthorization,
	ErrorNotYetValidAuthorization,
	ErrorRevokedAuthorization,
	ErrorIdentityMismatch,
	ErrorCaseMismatch,
	ErrorScopeMismatch,
	ErrorConnectorNotAllowed,
	ErrorReplayDetected,
	ErrorStaleDecision,
	ErrorMalformedConnectorResult,
	ErrorIntegrityFailure,
	ErrorStorageFailure,
	ErrorInternalInvariantFailure,
	ErrorUnsupportedSchema,
	ErrorOversizedInput,
}

func AllErrorCodes() []ErrorCode {
	return append([]ErrorCode(nil), errorCodes[:]...)
}

type AuditEventType string

const (
	AuditValidationRejected       AuditEventType = "VALIDATION_REJECTED"
	AuditAuthorizationDenied      AuditEventType = "AUTHORIZATION_DENIED"
	AuditAuthorizationAllowed     AuditEventType = "AUTHORIZATION_ALLOWED"
	AuditReplayRejected           AuditEventType = "REPLAY_REJECTED"
	AuditStaleContextRejected     AuditEventType = "STALE_CONTEXT_REJECTED"
	AuditRevokedGrantRejected     AuditEventType = "REVOKED_GRANT_REJECTED"
	AuditConnectorInvoked         AuditEventType = "CONNECTOR_INVOKED"
	AuditConnectorResultRejected  AuditEventType = "CONNECTOR_RESULT_REJECTED"
	AuditOperationCompleted       AuditEventType = "OPERATION_COMPLETED"
	AuditEvidenceCreated          AuditEventType = "EVIDENCE_CREATED"
	AuditIntegrityFailure         AuditEventType = "INTEGRITY_FAILURE"
	AuditStorageFailure           AuditEventType = "STORAGE_FAILURE"
	AuditInternalInvariantFailure AuditEventType = "INTERNAL_INVARIANT_FAILURE"
)

var auditEventTypes = [...]AuditEventType{
	AuditValidationRejected,
	AuditAuthorizationDenied,
	AuditAuthorizationAllowed,
	AuditReplayRejected,
	AuditStaleContextRejected,
	AuditRevokedGrantRejected,
	AuditConnectorInvoked,
	AuditConnectorResultRejected,
	AuditOperationCompleted,
	AuditEvidenceCreated,
	AuditIntegrityFailure,
	AuditStorageFailure,
	AuditInternalInvariantFailure,
}

func AllAuditEventTypes() []AuditEventType {
	return append([]AuditEventType(nil), auditEventTypes[:]...)
}

func AuditEventTypeForError(code ErrorCode) (AuditEventType, bool) {
	switch code {
	case ErrorValidation, ErrorUnsupportedSchema, ErrorOversizedInput:
		return AuditValidationRejected, true
	case ErrorAuthorizationDenied, ErrorExpiredAuthorization, ErrorNotYetValidAuthorization,
		ErrorIdentityMismatch, ErrorCaseMismatch, ErrorScopeMismatch, ErrorConnectorNotAllowed:
		return AuditAuthorizationDenied, true
	case ErrorRevokedAuthorization:
		return AuditRevokedGrantRejected, true
	case ErrorReplayDetected:
		return AuditReplayRejected, true
	case ErrorStaleDecision:
		return AuditStaleContextRejected, true
	case ErrorMalformedConnectorResult:
		return AuditConnectorResultRejected, true
	case ErrorIntegrityFailure:
		return AuditIntegrityFailure, true
	case ErrorStorageFailure:
		return AuditStorageFailure, true
	case ErrorInternalInvariantFailure:
		return AuditInternalInvariantFailure, true
	default:
		return "", false
	}
}

type OfficerStatus string

const (
	OfficerActive   OfficerStatus = "ACTIVE"
	OfficerDisabled OfficerStatus = "DISABLED"
)

type AuthorizationGrantState string

const (
	GrantActive      AuthorizationGrantState = "ACTIVE"
	GrantExpired     AuthorizationGrantState = "EXPIRED"
	GrantRevoked     AuthorizationGrantState = "REVOKED"
	GrantNotYetValid AuthorizationGrantState = "NOT_YET_VALID"
)

type PolicyDecisionValue string

const (
	PolicyAllow PolicyDecisionValue = "ALLOW"
	PolicyDeny  PolicyDecisionValue = "DENY"
)

type ConnectorResultStatus string

const (
	ConnectorResultSuccess ConnectorResultStatus = "SUCCESS"
	ConnectorResultFailure ConnectorResultStatus = "FAILURE"
)

type ReplayState string

const (
	ReplayReserved             ReplayState = "RESERVED"
	ReplayDenied               ReplayState = "DENIED"
	ReplayConnectorInvoked     ReplayState = "CONNECTOR_INVOKED"
	ReplayResultRejected       ReplayState = "RESULT_REJECTED"
	ReplayCompleted            ReplayState = "COMPLETED"
	ReplayIndeterminateOutcome ReplayState = "INDETERMINATE_OUTCOME"
)

type FixtureType string

const (
	FixtureSyntheticDomainFixture       FixtureType = "SYNTHETIC_DOMAIN_FIXTURE"
	FixtureCanonicalSerializationVector FixtureType = "CANONICAL_SERIALIZATION_VECTOR"
	FixtureNegativeTestInput             FixtureType = "NEGATIVE_TEST_INPUT"
	FixtureMockConnectorResponse         FixtureType = "MOCK_CONNECTOR_RESPONSE"
	FixtureAuthorizationFixture          FixtureType = "AUTHORIZATION_FIXTURE"
	FixtureCaseFixture                   FixtureType = "CASE_FIXTURE"
)

type FixtureCreationMethod string

const (
	FixtureManualInvention    FixtureCreationMethod = "MANUAL_INVENTION"
	FixtureSyntheticGenerator FixtureCreationMethod = "SYNTHETIC_GENERATOR"
)

type FixtureReviewStatus string

const (
	FixtureReviewPending  FixtureReviewStatus = "PENDING"
	FixtureReviewApproved FixtureReviewStatus = "APPROVED"
	FixtureReviewRejected FixtureReviewStatus = "REJECTED"
)
