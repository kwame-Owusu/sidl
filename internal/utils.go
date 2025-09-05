package internal

import (
	"fmt"
	"regexp"
)

var sidHexRe = regexp.MustCompile(`^[A-Z]{2}[0-9A-Fa-f]{32}$`)

func IsValidSid(sid string, file map[string]Field) (bool, error) {
	if len(sid) < 2 {
		return false, fmt.Errorf("SID too short")
	}

	identifier := sid[0:2]
	if _, ok := file[identifier]; !ok {
		return false, fmt.Errorf("Invalid twilio identifier: %s", identifier)
	}

	if !sidHexRe.MatchString(sid) {
		return false, fmt.Errorf("SID must be 34 characters: 2 letters + 32 hex chars")
	}

	return true, nil
}
