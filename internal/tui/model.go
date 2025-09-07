package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type SidInfo struct {
	Name        *string `json:"name"`
	Description string  `json:"description"`
}

type Model struct {
	list list.Model
	err  error
}

func NewModel() Model {
	// Start with an empty list
	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.Title = "SIDs from sids.json"
	return Model{list: l}
}

func (m Model) Init() tea.Cmd {
	return loadSidsCmd()
}
