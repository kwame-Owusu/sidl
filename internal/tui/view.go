package tui

import "fmt"

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPress q to quit.", m.err)
	}
	return m.list.View()
}
