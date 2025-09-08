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
			switch msg.String() {
			case "l":
				m.mode = ModeList
			case "e":
				m.mode = ModeExplainInput
				m.input.SetValue("")
				m.input.Focus()
				return m, nil
			case "q":
				return m, tea.Quit
			}

		case ModeList:
			switch msg.String() {
			case "h":
				m.mode = ModeHome
			case "q":
				return m, tea.Quit
			default:
				// Forward all other messages to the list
				var cmd tea.Cmd
				m.list, cmd = m.list.Update(msg)
				return m, cmd
			}

		case ModeExplainInput:
			switch msg.String() {
			case "enter":
				sid := m.input.Value()
				expl := runExplain(m.sids, sid)
				m.explanations = append(m.explanations, expl)
				m.input.SetValue("") // clear input for next SID
				m.input.Focus()
			case "h":
				m.mode = ModeHome
			}
		}
	}

	// Let the list/input handle the messages
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
	if !ok {
		return "Unknown SID"
	}

	name := "<none>"
	if info.Name != "" {
		name = info.Name
	}

	return fmt.Sprintf("Prefix: %s\nName: %s\nDescription: %s", prefix, name, info.Description)
}
