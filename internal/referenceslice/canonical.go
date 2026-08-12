package referenceslice

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"math"
	"sort"
	"strconv"
	"unicode/utf16"
)

// Canonicalize implements the frozen, integer-only JCS subset.
func Canonicalize(object any) ([]byte, *ValidationError) {
	if failure := ValidateGovernedObject(object); failure != nil {
		return nil, failure
	}
	raw, err := json.Marshal(object)
	if err != nil {
		return nil, validationFailure(ErrorIntegrityFailure, "marshal failed")
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	var value any
	if err := decoder.Decode(&value); err != nil {
		return nil, validationFailure(ErrorIntegrityFailure, "decode failed")
	}
	var out bytes.Buffer
	if failure := writeCanonical(&out, value); failure != nil {
		return nil, failure
	}
	return out.Bytes(), nil
}

func writeCanonical(out *bytes.Buffer, v any) *ValidationError {
	switch x := v.(type) {
	case nil:
		return validationFailure(ErrorIntegrityFailure, "null prohibited")
	case bool:
		if x {
			out.WriteString("true")
		} else {
			out.WriteString("false")
		}
	case string:
		b, _ := json.Marshal(x)
		out.Write(b)
	case json.Number:
		s := x.String()
		if bytes.ContainsAny([]byte(s), ".eE") || s == "-0" {
			return validationFailure(ErrorIntegrityFailure, "non-integer prohibited")
		}
		n, e := strconv.ParseInt(s, 10, 64)
		if e != nil || n < -9007199254740991 || n > 9007199254740991 {
			return validationFailure(ErrorIntegrityFailure, "integer outside profile")
		}
		out.WriteString(strconv.FormatInt(n, 10))
	case []any:
		out.WriteByte('[')
		for i, item := range x {
			if i > 0 {
				out.WriteByte(',')
			}
			if f := writeCanonical(out, item); f != nil {
				return f
			}
		}
		out.WriteByte(']')
	case map[string]any:
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return utf16Less(keys[i], keys[j]) })
		out.WriteByte('{')
		for i, k := range keys {
			if i > 0 {
				out.WriteByte(',')
			}
			b, _ := json.Marshal(k)
			out.Write(b)
			out.WriteByte(':')
			if f := writeCanonical(out, x[k]); f != nil {
				return f
			}
		}
		out.WriteByte('}')
	default:
		return validationFailure(ErrorIntegrityFailure, "unsupported canonical value")
	}
	return nil
}

func utf16Less(a, b string) bool {
	aa := utf16.Encode([]rune(a))
	bb := utf16.Encode([]rune(b))
	for i := 0; i < len(aa) && i < len(bb); i++ {
		if aa[i] != bb[i] {
			return aa[i] < bb[i]
		}
	}
	return len(aa) < len(bb)
}

func DomainSeparatedPreimage(objectType ObjectType, schemaVersion string, canonical []byte) ([]byte, *ValidationError) {
	if len(objectType) > math.MaxUint16 || len(schemaVersion) > math.MaxUint16 {
		return nil, validationFailure(ErrorIntegrityFailure, "domain field too long")
	}
	var out bytes.Buffer
	out.WriteString("OBDIA")
	out.WriteByte(0)
	// #nosec G115 -- both lengths are rejected above when they exceed MaxUint16.
	_ = binary.Write(&out, binary.BigEndian, uint16(len(objectType)))
	out.WriteString(string(objectType))
	// #nosec G115 -- both lengths are rejected above when they exceed MaxUint16.
	_ = binary.Write(&out, binary.BigEndian, uint16(len(schemaVersion)))
	out.WriteString(schemaVersion)
	_ = binary.Write(&out, binary.BigEndian, uint64(len(canonical)))
	out.Write(canonical)
	return out.Bytes(), nil
}

func IntegrityDigest(object any, objectType ObjectType) (Digest, []byte, *ValidationError) {
	canonical, f := Canonicalize(object)
	if f != nil {
		return "", nil, f
	}
	pre, f := DomainSeparatedPreimage(objectType, SchemaVersion, canonical)
	if f != nil {
		return "", nil, f
	}
	sum := sha256.Sum256(pre)
	return Digest(hex.EncodeToString(sum[:])), canonical, nil
}

func VerifyIntegrity(object any, objectType ObjectType, expected Digest) bool {
	actual, _, failure := IntegrityDigest(object, objectType)
	return failure == nil && actual == expected
}
