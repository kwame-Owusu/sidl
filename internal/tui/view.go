package tui

import "fmt"

func (m Model) View() string {
	switch m.mode {
	case ModeHome:
		return "Welcome to sidl\n\n" +
			"Available commands:\n" +
			"  l  → View all SID types\n" +
			"  e  → Explain a SID\n\n" +
			"Press q to quit."
	case ModeList:
		hint := "\n\nPress h to get back to home"
		return m.list.View() + hint
	case ModeExplainInput:
		if m.explanation != "" {
			return fmt.Sprintf("%s\n\nPress esc to go back.", m.explanation)
		}
		return fmt.Sprintf("Paste a SID and press Enter:\n\n%s", m.input.View())
	default:
		return "Unknown mode"
	}
}
