package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kwame-Owusu/sidl/internal"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-2)

	case tea.KeyMsg:
		switch m.mode {
		case ModeHome:
			if msg.String() == "l" {
				m.mode = ModeList
			}
			if msg.String() == "e" { // "e" to enter explain mode
				m.mode = ModeExplainInput
				m.input.Focus()
			}
			if msg.String() == "q" {
				return m, tea.Quit
			}

		case ModeList:
			if msg.String() == "h" { // back to home
				m.mode = ModeHome
			}
			if msg.String() == "q" {
				return m, tea.Quit
			}

		case ModeExplainInput:
			if msg.String() == "enter" {
				// Run explain on m.input.Value()
				sid := m.input.Value()
				m.explanation = runExplain(m.sids, sid) // returns formatted string
			}
			if msg.String() == "esc" {
				m.mode = ModeHome
			}
		}
	}

	// update list or input based on mode
	if m.mode == ModeList {
		m.list, cmd = m.list.Update(msg)
	}
	if m.mode == ModeExplainInput {
		m.input, cmd = m.input.Update(msg)
	}

	return m, cmd
}

func runExplain(sids map[string]internal.Field, sid string) string {
	upperSid := strings.ToUpper(sid)
	if len(upperSid) < 2 {
		return "Invalid SID"
	}
	prefix := upperSid[:2]

	info, ok := sids[prefix]
	if !ok || info.Name == "" {
		return "Unknown SID"
	}

	name := "<none>"
	if info.Name != "" {
		name = info.Name
	}

	return fmt.Sprintf("Prefix: %s\nName: %s\nDescription: %s", prefix, name, info.Description)
}
