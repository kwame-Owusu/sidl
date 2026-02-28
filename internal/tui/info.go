package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func RenderInfo(prefix, name, description string) string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("69"))

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("255"))

	block := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s %s", titleStyle.Render("Prefix:"), valueStyle.Render(prefix)),
		fmt.Sprintf("%s %s", titleStyle.Render("Name:"), valueStyle.Render(name)),
		fmt.Sprintf("%s %s", titleStyle.Render("Description:"), valueStyle.Render(description)),
	)

	card := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Render(block)

	return card
}
