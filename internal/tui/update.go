package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-2)

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	if m.mode == ModeList {
		// Handle list mode
		switch msg := msg.(type) {
		case sidsLoadedMsg:
			items := toListItems(msg)
			m.list.SetItems(items)
		case errMsg:
			m.err = msg
		}
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}
