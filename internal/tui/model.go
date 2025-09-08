package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	ModeHome mode = iota
	ModeList
)

type SidInfo struct {
	Name        *string `json:"name"`
	Description string  `json:"description"`
}

type Model struct {
	mode mode
	list list.Model
	err  error
}

func NewModel(startMode mode) Model {
	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.Title = "SIDs from sids.json"
	return Model{mode: startMode, list: l}
}

func (m Model) Init() tea.Cmd {
	if m.mode == ModeList {
		return loadSidsCmd()
	}
	return nil
}
