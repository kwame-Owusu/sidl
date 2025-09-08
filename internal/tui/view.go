package tui

import "fmt"

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPress q to quit.", m.err)
	}

	switch m.mode {
	case ModeHome:
		return "Welcome to sidl TUI 🎉\n\n" +
			"Available commands:\n" +
			"  list   → View all SID types\n\n" +
			"Press q to quit."
	case ModeList:
		return m.list.View()
	default:
		return "Unknown mode"
	}
}
