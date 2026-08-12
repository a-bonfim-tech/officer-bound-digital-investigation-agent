package referenceslice

import (
	"bytes"
	"testing"
)

func TestValidateRawJSON(t *testing.T) {
	directBidi := append([]byte(`{"value":"`), []byte("\u202e")...)
	directBidi = append(directBidi, []byte(`"}`)...)
	tests := []struct {
		name string
		raw  []byte
		max  int
		code ErrorCode
	}{
		{"simple object", []byte(`{"a":1}`), 64, ""},
		{"nested object", []byte(`{"a":{"b":[1,true,null]}}`), 64, ""},
		{"exact byte limit", []byte(`{"a":1}`), len([]byte(`{"a":1}`)), ""},
		{"over byte limit", []byte(`{"a":1}`), len([]byte(`{"a":1}`))-1, ErrorOversizedInput},
		{"zero byte limit", []byte(`{}`), 0, ErrorInternalInvariantFailure},
		{"invalid UTF-8", []byte{0xff}, 64, ErrorValidation},
		{"root duplicate", []byte(`{"a":1,"a":2}`), 64, ErrorValidation},
		{"nested duplicate", []byte(`{"a":{"b":1,"b":2}}`), 64, ErrorValidation},
		{"decoded duplicate", []byte(`{"a":1,"\u0061":2}`), 64, ErrorValidation},
		{"multiple values", []byte(`{} {}`), 64, ErrorValidation},
		{"trailing content", []byte(`{}x`), 64, ErrorValidation},
		{"trailing whitespace", []byte("{} \n\t"), 64, ""},
		{"malformed object", []byte(`{"a":}`), 64, ErrorValidation},
		{"malformed array", []byte(`[1,]`), 64, ErrorValidation},
		{"unpaired high surrogate", []byte(`{"a":"\uD800"}`), 64, ErrorValidation},
		{"unpaired low surrogate", []byte(`{"a":"\uDC00"}`), 64, ErrorValidation},
		{"valid surrogate pair", []byte(`{"a":"\uD83D\uDE00"}`), 64, ""},
		{"escaped NUL", []byte(`{"a":"\u0000"}`), 64, ErrorValidation},
		{"escaped newline", []byte(`{"a":"\n"}`), 64, ErrorValidation},
		{"direct bidi", directBidi, 64, ErrorValidation},
		{"escaped bidi", []byte(`{"a":"\u202e"}`), 64, ErrorValidation},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			document, err := ValidateRawJSON(test.raw, RawInputPolicy{MaxBytes: test.max})
			if test.code == "" {
				if err != nil {
					t.Fatalf("unexpected error: %#v", err)
				}
				if !bytes.Equal(document.Bytes(), test.raw) {
					t.Fatal("raw bytes were not preserved")
				}
				return
			}
			if err == nil || err.Code != test.code {
				t.Fatalf("got error %#v, want code %q", err, test.code)
			}
			if document.Bytes() != nil {
				t.Fatal("failed validation returned usable document")
			}
		})
	}
}

func TestRawDocumentBytesDefensiveCopy(t *testing.T) {
	raw := []byte(`{"a":1}`)
	document, err := ValidateRawJSON(raw, RawInputPolicy{MaxBytes: len(raw)})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	raw[0] = '['
	first := document.Bytes()
	first[0] = '['
	second := document.Bytes()
	if string(second) != `{"a":1}` {
		t.Fatalf("document bytes mutated: %q", second)
	}
}

func TestValidationErrorStringIsMinimized(t *testing.T) {
	err := &ValidationError{Code: ErrorValidation, Detail: `payload={"secret":"value"}`}
	if got := err.Error(); got != string(ErrorValidation) {
		t.Fatalf("public error leaked detail: %q", got)
	}
}
