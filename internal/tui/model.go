package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func NewModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
