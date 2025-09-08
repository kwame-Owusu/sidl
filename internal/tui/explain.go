package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ExplainModel struct {
	table table.Model
	err   error
}

func NewExplainModel(name, description string, prefix string) ExplainModel {
	columns := []table.Column{
		{Title: "Field", Width: 15},
		{Title: "Value", Width: 60},
	}
	rows := []table.Row{
		{"Prefix", prefix},
		{"Name", fmt.Sprintf("%v", name)},
		{"Description", description},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	t.SetStyles(defaultTableStyles())

	return ExplainModel{table: t}
}

func (m ExplainModel) Init() tea.Cmd { return nil }

func (m ExplainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m ExplainModel) View() string {
	return m.table.View() + "\nPress q to quit."
}

// styles
func defaultTableStyles() table.Styles {
	s := table.DefaultStyles()

	// Header: bold with a bottom border
	s.Header = lipgloss.NewStyle().
		Bold(true).
		Border(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(lipgloss.Color("240")).
		Padding(0, 1)

	// Regular cells: a little horizontal padding
	s.Cell = lipgloss.NewStyle().Padding(0, 1)

	// Selected row: inverted colors
	s.Selected = lipgloss.NewStyle().
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Padding(0, 1)

	return s
}
