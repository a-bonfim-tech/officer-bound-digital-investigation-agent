package referenceslice

import (
	"bytes"
	"encoding/json"
	"io"
	"unicode/utf16"
	"unicode/utf8"
)

type RawDocument struct {
	raw []byte
}

func (d RawDocument) Bytes() []byte {
	return append([]byte(nil), d.raw...)
}

type RawInputPolicy struct {
	MaxBytes int
}

func ValidateRawJSON(raw []byte, policy RawInputPolicy) (RawDocument, *ValidationError) {
	if policy.MaxBytes <= 0 {
		return RawDocument{}, validationFailure(ErrorInternalInvariantFailure, "invalid raw input policy")
	}
	if len(raw) > policy.MaxBytes {
		return RawDocument{}, validationFailure(ErrorOversizedInput, "raw input exceeds byte limit")
	}
	if !utf8.Valid(raw) {
		return RawDocument{}, validationFailure(ErrorValidation, "raw input is not valid UTF-8")
	}
	if err := validateJSONStringSafety(raw); err != nil {
		return RawDocument{}, err
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	if err := consumeJSONValue(decoder); err != nil {
		return RawDocument{}, validationFailure(ErrorValidation, "invalid JSON document")
	}
	if _, err := decoder.Token(); err != io.EOF {
		return RawDocument{}, validationFailure(ErrorValidation, "JSON document must contain exactly one value")
	}
	return RawDocument{raw: append([]byte(nil), raw...)}, nil
}

func consumeJSONValue(decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		return err
	}
	delimiter, isDelimiter := token.(json.Delim)
	if !isDelimiter {
		return nil
	}
	switch delimiter {
	case '{':
		seen := make(map[string]struct{})
		for decoder.More() {
			keyToken, err := decoder.Token()
			if err != nil {
				return err
			}
			key, ok := keyToken.(string)
			if !ok {
				return errInvalidJSON
			}
			if _, exists := seen[key]; exists {
				return errDuplicateMember
			}
			seen[key] = struct{}{}
			if err := consumeJSONValue(decoder); err != nil {
				return err
			}
		}
		closing, err := decoder.Token()
		if err != nil || closing != json.Delim('}') {
			if err != nil {
				return err
			}
			return errInvalidJSON
		}
		return nil
	case '[':
		for decoder.More() {
			if err := consumeJSONValue(decoder); err != nil {
				return err
			}
		}
		closing, err := decoder.Token()
		if err != nil || closing != json.Delim(']') {
			if err != nil {
				return err
			}
			return errInvalidJSON
		}
		return nil
	default:
		return errInvalidJSON
	}
}

type strictJSONError string

func (e strictJSONError) Error() string { return string(e) }

const (
	errDuplicateMember strictJSONError = "duplicate object member"
	errInvalidJSON     strictJSONError = "invalid JSON"
)

func validateJSONStringSafety(raw []byte) *ValidationError {
	for index := 0; index < len(raw); index++ {
		if raw[index] != '"' {
			continue
		}
		next, err := scanJSONString(raw, index+1)
		if err != nil {
			return err
		}
		index = next - 1
	}
	return nil
}

func scanJSONString(raw []byte, index int) (int, *ValidationError) {
	for index < len(raw) {
		current := raw[index]
		switch {
		case current == '"':
			return index + 1, nil
		case current == '\\':
			next, scalar, hasScalar, err := scanJSONEscape(raw, index)
			if err != nil {
				return 0, err
			}
			if hasScalar && forbiddenJSONScalar(scalar) {
				return 0, validationFailure(ErrorValidation, "forbidden JSON string scalar")
			}
			index = next
		case current < 0x20:
			return 0, validationFailure(ErrorValidation, "forbidden JSON string control")
		default:
			scalar, width := utf8.DecodeRune(raw[index:])
			if scalar == utf8.RuneError && width == 1 {
				return 0, validationFailure(ErrorValidation, "invalid JSON string encoding")
			}
			if forbiddenJSONScalar(scalar) {
				return 0, validationFailure(ErrorValidation, "forbidden JSON string scalar")
			}
			index += width
		}
	}
	return 0, validationFailure(ErrorValidation, "unterminated JSON string")
}

func scanJSONEscape(raw []byte, index int) (int, rune, bool, *ValidationError) {
	if index+1 >= len(raw) {
		return 0, 0, false, validationFailure(ErrorValidation, "invalid JSON escape")
	}
	escaped := raw[index+1]
	switch escaped {
	case '"', '\\', '/':
		return index + 2, 0, false, nil
	case 'b', 'f', 'n', 'r', 't':
		return 0, 0, false, validationFailure(ErrorValidation, "forbidden escaped JSON control")
	case 'u':
		first, ok := decodeHexQuad(raw, index+2)
		if !ok {
			return 0, 0, false, validationFailure(ErrorValidation, "invalid JSON Unicode escape")
		}
		next := index + 6
		if first >= 0xd800 && first <= 0xdbff {
			if next+6 > len(raw) || raw[next] != '\\' || raw[next+1] != 'u' {
				return 0, 0, false, validationFailure(ErrorValidation, "unpaired high surrogate")
			}
			second, ok := decodeHexQuad(raw, next+2)
			if !ok || second < 0xdc00 || second > 0xdfff {
				return 0, 0, false, validationFailure(ErrorValidation, "unpaired high surrogate")
			}
			return next + 6, utf16.DecodeRune(rune(first), rune(second)), true, nil
		}
		if first >= 0xdc00 && first <= 0xdfff {
			return 0, 0, false, validationFailure(ErrorValidation, "unpaired low surrogate")
		}
		return next, rune(first), true, nil
	default:
		return 0, 0, false, validationFailure(ErrorValidation, "invalid JSON escape")
	}
}

func decodeHexQuad(raw []byte, index int) (uint16, bool) {
	if index+4 > len(raw) {
		return 0, false
	}
	var value uint16
	for offset := 0; offset < 4; offset++ {
		digit, ok := hexValue(raw[index+offset])
		if !ok {
			return 0, false
		}
		value = value*16 + uint16(digit)
	}
	return value, true
}

func hexValue(value byte) (byte, bool) {
	switch {
	case value >= '0' && value <= '9':
		return value - '0', true
	case value >= 'a' && value <= 'f':
		return value - 'a' + 10, true
	case value >= 'A' && value <= 'F':
		return value - 'A' + 10, true
	default:
		return 0, false
	}
}

func forbiddenJSONScalar(value rune) bool {
	return value <= 0x1f ||
		(value >= 0x7f && value <= 0x9f) ||
		value == 0x061c ||
		value == 0x200e || value == 0x200f ||
		(value >= 0x202a && value <= 0x202e) ||
		(value >= 0x2066 && value <= 0x2069)
}
