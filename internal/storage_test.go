package internal

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFile(t *testing.T) {
	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Test valid JSON file
	t.Run("valid file", func(t *testing.T) {
		file := filepath.Join(tmpDir, "test.json")
		content := `{"field1": {"name": "test", "description": "hello world"}}`
		os.WriteFile(file, []byte(content), 0644)

		result, err := LoadFile(file)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(result) != 1 {
			t.Errorf("expected 1 field, got %d", len(result))
		}
	})

	// Test invalid extension
	t.Run("invalid extension", func(t *testing.T) {
		_, err := LoadFile("test.txt")
		if err == nil {
			t.Error("expected error for non-json file")
		}
		if !strings.Contains(err.Error(), "invalid filename") {
			t.Errorf("wrong error message: %v", err)
		}
	})

	// Test file not found
	t.Run("file not found", func(t *testing.T) {
		_, err := LoadFile("missing.json")
		if err == nil {
			t.Error("expected error for missing file")
		}
		if !strings.Contains(err.Error(), "error reading file") {
			t.Errorf("wrong error message: %v", err)
		}
	})

	// Test invalid JSON
	t.Run("invalid json", func(t *testing.T) {
		file := filepath.Join(tmpDir, "bad.json")
		os.WriteFile(file, []byte(`{invalid json`), 0644)

		_, err := LoadFile(file)
		if err == nil {
			t.Error("expected error for invalid json")
		}
		if !strings.Contains(err.Error(), "error unmarshaling") {
			t.Errorf("wrong error message: %v", err)
		}
	})
}
