package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kwame-Owusu/sidl/internal"
)

type mode int

const (
	ModeHome mode = iota
	ModeList
	ModeExplainInput
)

type SidInfo struct {
	Name        *string `json:"name"`
	Description string  `json:"description"`
}

type Model struct {
	mode        mode
	list        list.Model
	err         error
	input       textinput.Model // for SID input
	explanation string          // store result
	sids        map[string]internal.Field
}

func NewModel(startMode mode, sids map[string]internal.Field) Model {
	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.Title = "SIDs from sids.json"

	ti := textinput.New()
	ti.Placeholder = "Paste SID here..."
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 40

	return Model{
		mode:  startMode,
		list:  l,
		input: ti,
		sids:  sids,
	}
}

func (m Model) Init() tea.Cmd {
	if m.mode == ModeList {
		return loadSidsCmd()
	}
	return nil
}
