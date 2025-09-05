package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type SIDEntry struct {
	Prefix, Name, Description string
}

type Model struct {
	sids   []SIDEntry
	cursor int
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.sids)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := titleStyle.Render("Select a Twilio SID:") + "\n\n"

	for i, sid := range m.sids {
		line := fmt.Sprintf("%s - %s", sid.Prefix, sid.Name)
		if i == m.cursor {
			s += cursorStyle.Render("> ") +
				selectedStyle.Render(line) + "\n"
			if sid.Description != "" {
				s += descStyle.Render("    "+sid.Description) + "\n\n"
			}
		} else {
			s += "  " + line + "\n"
		}
	}
	s += "\nPress ↑/↓ to navigate • Press q to quit."
	return s
}
