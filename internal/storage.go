package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func LoadFile(filename string) (map[string]Field, error) {
	var data map[string]Field

	if !strings.HasSuffix(filename, ".json") {
		return nil, fmt.Errorf("invalid filename, file has to end in .json")
	}

	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling file: %s", err)
	}
	return data, nil
}
