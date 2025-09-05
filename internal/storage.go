package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func LoadFile(filename string) (os.File, error) {
	if !strings.HasSuffix(filename, ".json") {
		return nil, fmt.Errorf("Invalid filename, file has to end in .json")
	}

	var file os.File
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Error reading file, %s", err)
	}
	err = json.Unmarshal(data, &file)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling file, %s", err)
	}

	return file, nil
}
