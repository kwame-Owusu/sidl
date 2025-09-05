package tui

import (
	"encoding/json"
	"os"
)

func loadSIDs() ([]SIDEntry, error) {
	var raw map[string]struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	file, err := os.Open("sids.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&raw); err != nil {
		return nil, err
	}

	var list []SIDEntry
	for k, v := range raw {
		list = append(list, SIDEntry{
			Prefix:      k,
			Name:        v.Name,
			Description: v.Description,
		})
	}
	return list, nil
}
