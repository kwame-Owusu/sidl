package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type TypesModel struct {
	table table.Model
	err   error
}

func NewTypesModel(name string) TypesModel {
	columns := []table.Column{
		{Title: "No", Width: 15},
		{Title: "Type", Width: 60},
	}

	rows := []table.Row{
		{"-", name},
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	t.SetStyles(defaultTableStyles())

	return TypesModel{table: t}
}

func (m TypesModel) Init() tea.Cmd { return nil }

func (m TypesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m TypesModel) View() string {
	return "\n" + m.table.View() + "\n\nPress q to quit."
}

// styles
func defaultTableStyles() table.Styles {
	s := table.DefaultStyles()
	return s
}
