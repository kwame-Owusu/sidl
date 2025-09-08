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
		switch m.mode {
		case ModeHome:
			// Home mode keybindings
			if msg.String() == "l" {
				// Switch to list mode
				m.mode = ModeList
				return m, loadSidsCmd()
			}
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		case ModeList:
			// List mode keybindings
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		}

	case sidsLoadedMsg:
		if m.mode == ModeList {
			items := toListItems(msg)
			m.list.SetItems(items)
		}

	case errMsg:
		m.err = msg
	}

	// Let the list handle messages if we're in list mode
	if m.mode == ModeList {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}
