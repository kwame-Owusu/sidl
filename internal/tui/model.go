package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type SidInfo struct {
	Name        *string `json:"name"`
	Description string  `json:"description"`
}

type Model struct {
	sids map[string]SidInfo
	err  error
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return loadSidsCmd()
}
