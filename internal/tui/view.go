package tui

import "fmt"

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPress q to quit.", m.err)
	}

	switch m.mode {
	case ModeHome:
		return "Welcome to sidl \n\n" +
			"Available commands:\n" +
			"  l  → View all SID types\n\n" +
			"Press q to quit."
	case ModeList:
		hint := "\n\nPress h to get back to home"
		return m.list.View() + hint
	default:
		return "Unknown mode"
	}
}
