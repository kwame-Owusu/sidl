package tui

import (
	"fmt"
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

type sidListItem struct {
	key  string
	name string
	desc string
}

func (i sidListItem) Title() string       { return fmt.Sprintf("[%s] %s", i.key, i.name) }
func (i sidListItem) Description() string { return i.desc }
func (i sidListItem) FilterValue() string { return i.name }

type SidInfo struct {
	Name        *string `json:"name"`
	Description string  `json:"description"`
}

type Model struct {
	mode         mode
	list         list.Model
	err          error
	input        textinput.Model // for SID input
	explanations []string        // store result
	sids         map[string]internal.Field
}

func NewModel(startMode mode) Model {
	sids := internal.LoadSIDs()

	// convert sids to items
	items := make([]list.Item, 0, len(sids))
	for key, s := range sids {
		name := "<none>"
		if s.Name != "" {
			name = s.Name
		}
		items = append(items, sidListItem{
			key:  key,
			name: name,
			desc: s.Description,
		})
	}

	l := list.New(items, list.NewDefaultDelegate(), 60, 20)
	l.Title = "SIDs from sids.json"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

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
