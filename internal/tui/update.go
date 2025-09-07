package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"sort"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Resize list to match terminal size
		m.list.SetSize(msg.Width, msg.Height-2)

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case sidsLoadedMsg:
		// Convert map to []SidItem
		keys := make([]string, 0, len(msg))
		for k := range msg {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		items := make([]list.Item, 0, len(keys))
		for _, k := range keys {
			items = append(items, SidItem{Key: k, Info: msg[k]})
		}
		m.list.SetItems(items)

	case errMsg:
		m.err = msg
	}

	// Let the list handle any messages too
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
