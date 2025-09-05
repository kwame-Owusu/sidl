package tui

import (
	"github.com/charmbracelet/bubbletea"
)

func Run() error {
	sids, err := loadSIDs()
	if err != nil {
		return err
	}

	m := Model{sids: sids, cursor: 0}
	p := tea.NewProgram(m)
	_, err = p.Run()
	return err
}
