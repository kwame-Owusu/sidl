package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ExplainModel struct {
	prefix      string
	name        string
	description string
}

func NewExplainModel(name, description, prefix string) ExplainModel {
	return ExplainModel{
		prefix:      prefix,
		name:        fmt.Sprintf("%v", name),
		description: description,
	}
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
	return m, nil
}

func (m ExplainModel) View() string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("69"))
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("255"))

	block := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s %s", titleStyle.Render("Prefix:"), valueStyle.Render(m.prefix)),
		fmt.Sprintf("%s %s", titleStyle.Render("Name:"), valueStyle.Render(m.name)),
		fmt.Sprintf("%s %s", titleStyle.Render("Description:"), valueStyle.Render(m.description)),
	)

	card := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Render(block)

	return card + "\n\nPress q to quit."
}
