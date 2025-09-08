package tui

import "fmt"

func (m Model) View() string {
	switch m.mode {
	case ModeHome:
		return "Welcome to sidl TUI\n\n" +
			"  l → View all SID types\n" +
			"  e → Explain a SID\n\n" +
			"Press q to quit."
	case ModeList:
		hint := "\n\nPress h to get back to home"
		return m.list.View() + hint
	case ModeExplainInput:
		// show all explanations above input
		explStr := ""
		for _, e := range m.explanations {
			explStr += e + "\n\n"
		}
		hint := "\n\nPress h to get back to home"
		return fmt.Sprintf("%sPaste a SID and press Enter:\n%s", explStr, m.input.View()+hint)
	default:
		return "Unknown mode"
	}
}
