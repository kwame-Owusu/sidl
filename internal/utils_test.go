package internal

import (
	"testing"
)

func TestIsValidSid(t *testing.T) {
	file := map[string]Field{
		"AC": {}, "CA": {}, "SM": {}, "PN": {}, "CF": {}, "RE": {},
		"QU": {}, "AP": {}, "TK": {}, "WS": {}, "WW": {}, "WK": {},
		"WT": {}, "RM": {}, "PA": {}, "CH": {}, "US": {}, "IS": {},
		"ET": {},
	}

	// ---------------------------
	// Positive test cases
	// ---------------------------
	tests := map[string][]string{
		"accounts":      {"AC1234567890abcdef1234567890abcdef", "AC9876543210fedcba9876543210fedcba", "ACabcdef1234567890abcdef1234567890"},
		"calls":         {"CA1234567890abcdef1234567890abcdef", "CA9876543210fedcba9876543210fedcba", "CAabcdef1234567890abcdef1234567890"},
		"messages":      {"SM1234567890abcdef1234567890abcdef", "SM9876543210fedcba9876543210fedcba", "SMabcdef1234567890abcdef1234567890"},
		"phone_numbers": {"PN1234567890abcdef1234567890abcdef", "PN9876543210fedcba9876543210fedcba", "PNabcdef1234567890abcdef1234567890"},
		"conferences":   {"CF1234567890abcdef1234567890abcdef", "CF9876543210fedcba9876543210fedcba", "CFabcdef1234567890abcdef1234567890"},
		"recordings":    {"RE1234567890abcdef1234567890abcdef", "RE9876543210fedcba9876543210fedcba", "REabcdef1234567890abcdef1234567890"},
		"queues":        {"QU1234567890abcdef1234567890abcdef", "QU9876543210fedcba9876543210fedcba", "QUabcdef1234567890abcdef1234567890"},
		"applications":  {"AP1234567890abcdef1234567890abcdef", "AP9876543210fedcba9876543210fedcba", "APabcdef1234567890abcdef1234567890"},
		"trunks":        {"TK1234567890abcdef1234567890abcdef", "TK9876543210fedcba9876543210fedcba", "TKabcdef1234567890abcdef1234567890"},
		"workspaces":    {"WS1234567890abcdef1234567890abcdef", "WS9876543210fedcba9876543210fedcba", "WSabcdef1234567890abcdef1234567890"},
		"workflows":     {"WW1234567890abcdef1234567890abcdef", "WW9876543210fedcba9876543210fedcba", "WWabcdef1234567890abcdef1234567890"},
		"workers":       {"WK1234567890abcdef1234567890abcdef", "WK9876543210fedcba9876543210fedcba", "WKabcdef1234567890abcdef1234567890"},
		"tasks":         {"WT1234567890abcdef1234567890abcdef", "WT9876543210fedcba9876543210fedcba", "WTabcdef1234567890abcdef1234567890"},
		"rooms":         {"RM1234567890abcdef1234567890abcdef", "RM9876543210fedcba9876543210fedcba", "RMabcdef1234567890abcdef1234567890"},
		"participants":  {"PA1234567890abcdef1234567890abcdef", "PA9876543210fedcba9876543210fedcba", "PAabcdef1234567890abcdef1234567890"},
		"conversations": {"CH1234567890abcdef1234567890abcdef", "CH9876543210fedcba9876543210fedcba", "CHabcdef1234567890abcdef1234567890"},
		"users":         {"US1234567890abcdef1234567890abcdef", "US9876543210fedcba9876543210fedcba", "USabcdef1234567890abcdef1234567890"},
		"services":      {"IS1234567890abcdef1234567890abcdef", "IS9876543210fedcba9876543210fedcba", "ISabcdef1234567890abcdef1234567890"},
		"documents":     {"ET1234567890abcdef1234567890abcdef", "ET9876543210fedcba9876543210fedcba", "ETAbcdef1234567890abcdef1234567890"},
	}

	for category, sids := range tests {
		for _, sid := range sids {
			valid, err := IsValidSid(sid, file)
			if !valid || err != nil {
				t.Errorf("Expected valid SID for category %s: %s, got err=%v", category, sid, err)
			}
		}
	}

	// ---------------------------
	// Negative test cases
	// ---------------------------
	invalidSids := []string{
		"ZZ1234567890abcdef1234567890abcdef",   // invalid prefix
		"AC123",                                // too short
		"AC1234567890abcdef1234567890abcdef00", // too long
		"AC1234567890abcdef1234567890abcdeg",   // invalid hex char "g"
		"",                                     // empty string
		"A",                                    // single char
	}

	for _, sid := range invalidSids {
		valid, err := IsValidSid(sid, file)
		if valid || err == nil {
			t.Errorf("Expected invalid SID: %s, but got valid", sid)
		}
	}
}
