package referenceslice

// Named domain primitives do not establish validity. In particular, Go zero
// values never establish authority; untrusted input requires later validation.
type (
	OfficerID     string
	CaseID        string
	GrantID       string
	DelegationID  string
	RequestID     string
	DecisionID    string
	ConnectorID   string
	OperationID   string
	EvidenceID    string
	AuditEventID  string
	Digest        string
	Timestamp     string
	NonEmptyASCII string
)

type SyntheticOfficerIdentity struct {
	SchemaVersion string        `json:"schema_version"`
	ObjectType    ObjectType    `json:"object_type"`
	OfficerID     OfficerID     `json:"officer_id"`
	InstitutionID NonEmptyASCII `json:"institution_id"`
	Synthetic     bool          `json:"synthetic"`
	Status        OfficerStatus `json:"status"`
}

type SyntheticCaseContext struct {
	SchemaVersion  string        `json:"schema_version"`
	ObjectType     ObjectType    `json:"object_type"`
	CaseID         CaseID        `json:"case_id"`
	Purpose        NonEmptyASCII `json:"purpose"`
	Classification string        `json:"classification"`
	Synthetic      bool          `json:"synthetic"`
}

type AuthorizationGrant struct {
	SchemaVersion        string                  `json:"schema_version"`
	ObjectType           ObjectType              `json:"object_type"`
	AuthorizationGrantID GrantID                 `json:"authorization_grant_id"`
	OfficerID            OfficerID               `json:"officer_id"`
	CaseID               CaseID                  `json:"case_id"`
	Scopes               []NonEmptyASCII         `json:"scopes"`
	IssuedAt             Timestamp               `json:"issued_at"`
	NotBefore            Timestamp               `json:"not_before"`
	ExpiresAt            Timestamp               `json:"expires_at"`
	State                AuthorizationGrantState `json:"state"`
}

type ScopedDelegation struct {
	SchemaVersion        string          `json:"schema_version"`
	ObjectType           ObjectType      `json:"object_type"`
	DelegationID         DelegationID    `json:"delegation_id"`
	AuthorizationGrantID GrantID         `json:"authorization_grant_id"`
	OfficerID            OfficerID       `json:"officer_id"`
	CaseID               CaseID          `json:"case_id"`
	PermittedActions     []NonEmptyASCII `json:"permitted_actions"`
	PermittedConnectors  []ConnectorID   `json:"permitted_connectors"`
	NotBefore            Timestamp       `json:"not_before"`
	ExpiresAt            Timestamp       `json:"expires_at"`
}

type RequestedAction struct {
	SchemaVersion string          `json:"schema_version"`
	ObjectType    ObjectType      `json:"object_type"`
	Action        NonEmptyASCII   `json:"action"`
	Target        NonEmptyASCII   `json:"target"`
	Parameters    []NonEmptyASCII `json:"parameters"`
}

type AuthorizationContext struct {
	SchemaVersion        string          `json:"schema_version"`
	ObjectType           ObjectType      `json:"object_type"`
	OfficerID            OfficerID       `json:"officer_id"`
	CaseID               CaseID          `json:"case_id"`
	AuthorizationGrantID GrantID         `json:"authorization_grant_id"`
	DelegationID         DelegationID    `json:"delegation_id"`
	RequestID            RequestID       `json:"request_id"`
	ConnectorID          ConnectorID     `json:"connector_id"`
	Action               RequestedAction `json:"action"`
	EvaluatedAt          Timestamp       `json:"evaluated_at"`
}

type PolicyDecision struct {
	SchemaVersion              string              `json:"schema_version"`
	ObjectType                 ObjectType          `json:"object_type"`
	DecisionID                 DecisionID          `json:"decision_id"`
	Decision                   PolicyDecisionValue `json:"decision"`
	OfficerID                  OfficerID           `json:"officer_id"`
	CaseID                     CaseID              `json:"case_id"`
	AuthorizationGrantID       GrantID             `json:"authorization_grant_id"`
	DelegationID               DelegationID        `json:"delegation_id"`
	RequestID                  RequestID           `json:"request_id"`
	ConnectorID                ConnectorID         `json:"connector_id"`
	ConnectorCapabilityVersion NonEmptyASCII       `json:"connector_capability_version"`
	PolicyVersion              NonEmptyASCII       `json:"policy_version"`
	CanonicalInputHash         Digest              `json:"canonical_input_hash"`
	EvaluatedAt                Timestamp           `json:"evaluated_at"`
	DecisionValidUntil         Timestamp           `json:"decision_valid_until"`
}

type ConnectorCapability struct {
	SchemaVersion     string        `json:"schema_version"`
	ObjectType        ObjectType    `json:"object_type"`
	ConnectorID       ConnectorID   `json:"connector_id"`
	CapabilityVersion NonEmptyASCII `json:"capability_version"`
	Operation         NonEmptyASCII `json:"operation"`
	RequestSchema     NonEmptyASCII `json:"request_schema"`
	ResultSchema      NonEmptyASCII `json:"result_schema"`
}

type ConnectorRequest struct {
	SchemaVersion      string      `json:"schema_version"`
	ObjectType         ObjectType  `json:"object_type"`
	RequestID          RequestID   `json:"request_id"`
	DecisionID         DecisionID  `json:"decision_id"`
	ConnectorID        ConnectorID `json:"connector_id"`
	OperationID        OperationID `json:"operation_id"`
	CanonicalInputHash Digest      `json:"canonical_input_hash"`
}

type ConnectorResult struct {
	SchemaVersion string                `json:"schema_version"`
	ObjectType    ObjectType            `json:"object_type"`
	ConnectorID   ConnectorID           `json:"connector_id"`
	RequestID     RequestID             `json:"request_id"`
	OperationID   OperationID           `json:"operation_id"`
	Status        ConnectorResultStatus `json:"status"`
	ResultData    []NonEmptyASCII       `json:"result_data"`
}

type EvidenceEnvelope struct {
	SchemaVersion        string       `json:"schema_version"`
	ObjectType           ObjectType   `json:"object_type"`
	EvidenceID           EvidenceID   `json:"evidence_id"`
	OfficerID            OfficerID    `json:"officer_id"`
	CaseID               CaseID       `json:"case_id"`
	AuthorizationGrantID GrantID      `json:"authorization_grant_id"`
	DelegationID         DelegationID `json:"delegation_id"`
	RequestID            RequestID    `json:"request_id"`
	PolicyDecisionID     DecisionID   `json:"policy_decision_id"`
	ConnectorID          ConnectorID  `json:"connector_id"`
	OperationID          OperationID  `json:"operation_id"`
	ResultDigest         Digest       `json:"result_digest"`
	CreatedAt            Timestamp    `json:"created_at"`
	ProvenanceDigest     Digest       `json:"provenance_digest"`
}

type AuditEvent struct {
	SchemaVersion string          `json:"schema_version"`
	ObjectType    ObjectType      `json:"object_type"`
	AuditEventID  AuditEventID    `json:"audit_event_id"`
	EventType     AuditEventType  `json:"event_type"`
	CaseID        CaseID          `json:"case_id"`
	RequestID     RequestID       `json:"request_id"`
	OccurredAt    Timestamp       `json:"occurred_at"`
	ErrorCode     *ErrorCode      `json:"error_code,omitempty"`
}

type ProvenanceRecord struct {
	SchemaVersion   string        `json:"schema_version"`
	ObjectType      ObjectType    `json:"object_type"`
	Producer        NonEmptyASCII `json:"producer"`
	ProducerVersion NonEmptyASCII `json:"producer_version"`
	SourceDigests   []Digest      `json:"source_digests"`
	PriorLinkDigest Digest        `json:"prior_link_digest"`
	CreatedAt       Timestamp     `json:"created_at"`
}

type RequestReservationRecord struct {
	SchemaVersion      string     `json:"schema_version"`
	ObjectType         ObjectType `json:"object_type"`
	RequestID          RequestID  `json:"request_id"`
	CaseID             CaseID     `json:"case_id"`
	CanonicalInputHash Digest     `json:"canonical_input_hash"`
	ReservedAt         Timestamp  `json:"reserved_at"`
}

type ReplayRecord struct {
	SchemaVersion      string      `json:"schema_version"`
	ObjectType         ObjectType  `json:"object_type"`
	CaseID             CaseID      `json:"case_id"`
	RequestID          RequestID   `json:"request_id"`
	CanonicalInputHash Digest      `json:"canonical_input_hash"`
	ReservedAt         Timestamp   `json:"reserved_at"`
	State              ReplayState `json:"state"`
}

type FixtureProvenance struct {
	SchemaVersion               string                `json:"schema_version"`
	ObjectType                  ObjectType            `json:"object_type"`
	FixtureID                   NonEmptyASCII         `json:"fixture_id"`
	FixtureType                 FixtureType           `json:"fixture_type"`
	Classification             string                `json:"classification"`
	Synthetic                  bool                  `json:"synthetic"`
	CreationMethod             FixtureCreationMethod `json:"creation_method"`
	CreatorOrGenerator         NonEmptyASCII         `json:"creator_or_generator"`
	GeneratorVersion           NonEmptyASCII         `json:"generator_version"`
	CreationTimestamp          Timestamp             `json:"creation_timestamp"`
	SourceDescription          NonEmptyASCII         `json:"source_description"`
	DerivedFromRealPersonData  bool                  `json:"derived_from_real_person_data"`
	DerivedFromLiveCaseData    bool                  `json:"derived_from_live_case_data"`
	DerivedFromProductionExport bool                 `json:"derived_from_production_export"`
	DerivedFromLiveConnector   bool                  `json:"derived_from_live_connector"`
	ContainsRealCredentials    bool                  `json:"contains_real_credentials"`
	ContentDigest              Digest                `json:"content_digest"`
	ReviewStatus               FixtureReviewStatus   `json:"review_status"`
	Reviewer                   *NonEmptyASCII        `json:"reviewer,omitempty"`
	ReviewTimestamp            *Timestamp            `json:"review_timestamp,omitempty"`
}
