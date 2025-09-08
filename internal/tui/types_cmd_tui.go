package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TypesModel struct {
	content string
	width   int
}

func NewTypesModel(content string) TypesModel {
	return TypesModel{content: content, width: 80}
}

func (m TypesModel) Init() tea.Cmd { return nil }

func (m TypesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m TypesModel) View() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("69")).
		Padding(1, 2).
		Width(m.width - 4)

	return style.Render(m.content) + "\n\nPress q to quit."
}
