package referenceslice

import "time"

type ValidationError struct {
	Code   ErrorCode
	Detail string
}

func (e *ValidationError) Error() string {
	if e == nil {
		return ""
	}
	return string(e.Code)
}

func validationFailure(code ErrorCode, detail string) *ValidationError {
	return &ValidationError{Code: code, Detail: detail}
}

func ValidateSchemaVersion(value string) *ValidationError {
	if value != SchemaVersion {
		return validationFailure(ErrorUnsupportedSchema, "unsupported schema version")
	}
	return nil
}

func ValidateOfficerID(value OfficerID) *ValidationError {
	return validateIdentifier(string(value), "off_")
}

func ValidateCaseID(value CaseID) *ValidationError {
	return validateIdentifier(string(value), "case_")
}

func ValidateGrantID(value GrantID) *ValidationError {
	return validateIdentifier(string(value), "grant_")
}

func ValidateDelegationID(value DelegationID) *ValidationError {
	return validateIdentifier(string(value), "dlg_")
}

func ValidateRequestID(value RequestID) *ValidationError {
	return validateIdentifier(string(value), "req_")
}

func ValidateDecisionID(value DecisionID) *ValidationError {
	return validateIdentifier(string(value), "dec_")
}

func ValidateConnectorID(value ConnectorID) *ValidationError {
	return validateIdentifier(string(value), "conn_")
}

func ValidateOperationID(value OperationID) *ValidationError {
	return validateIdentifier(string(value), "op_")
}

func ValidateEvidenceID(value EvidenceID) *ValidationError {
	return validateIdentifier(string(value), "evid_")
}

func ValidateAuditEventID(value AuditEventID) *ValidationError {
	return validateIdentifier(string(value), "evt_")
}

func validateIdentifier(value, prefix string) *ValidationError {
	if len(value) != len(prefix)+32 || value[:min(len(value), len(prefix))] != prefix {
		return validationFailure(ErrorValidation, "invalid identifier")
	}
	if !isLowerHex(value[len(prefix):]) {
		return validationFailure(ErrorValidation, "invalid identifier")
	}
	return nil
}

func ValidateDigest(value Digest) *ValidationError {
	if len(value) != 64 || !isLowerHex(string(value)) {
		return validationFailure(ErrorValidation, "invalid digest")
	}
	return nil
}

func isLowerHex(value string) bool {
	for i := 0; i < len(value); i++ {
		if (value[i] < '0' || value[i] > '9') && (value[i] < 'a' || value[i] > 'f') {
			return false
		}
	}
	return true
}

func ValidateNonEmptyASCII(value NonEmptyASCII) *ValidationError {
	if len(value) < 1 || len(value) > 128 {
		return validationFailure(ErrorValidation, "invalid ASCII value length")
	}
	for i := 0; i < len(value); i++ {
		if value[i] < 0x20 || value[i] > 0x7e {
			return validationFailure(ErrorValidation, "invalid ASCII value")
		}
	}
	return nil
}

const timestampLayout = "2006-01-02T15:04:05.000Z"

func ValidateTimestamp(value Timestamp) *ValidationError {
	if len(value) != len(timestampLayout) {
		return validationFailure(ErrorValidation, "invalid timestamp")
	}
	parsed, err := time.Parse(timestampLayout, string(value))
	if err != nil || parsed.Location() != time.UTC || parsed.Format(timestampLayout) != string(value) {
		return validationFailure(ErrorValidation, "invalid timestamp")
	}
	return nil
}

func UniqueNonEmptyASCII(values []NonEmptyASCII) bool {
	seen := make(map[NonEmptyASCII]struct{}, len(values))
	for _, value := range values {
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func UniqueConnectorIDs(values []ConnectorID) bool {
	seen := make(map[ConnectorID]struct{}, len(values))
	for _, value := range values {
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func UniqueDigests(values []Digest) bool {
	seen := make(map[Digest]struct{}, len(values))
	for _, value := range values {
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}
