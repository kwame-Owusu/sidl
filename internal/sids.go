package internal

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed sids.json
var sidsJSON []byte

var SIDs map[string]Field

func LoadSIDs() map[string]Field {
	if SIDs != nil {
		return SIDs
	}
	err := json.Unmarshal(sidsJSON, &SIDs)
	if err != nil {
		log.Fatalf("failed to parse embedded sids.json: %v", err)
	}
	return SIDs
}
