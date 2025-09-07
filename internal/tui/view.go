package tui

import (
	"fmt"
)

func (m model) View() string {
	return fmt.Sprintf("Hello from Bubble Tea!\n\nPress Ctrl+C to quit.")
}
