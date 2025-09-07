// This is a Bubble Tea command, not a Cobra command.
package tui

import (
	"encoding/json"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type sidsLoadedMsg map[string]SidInfo
type errMsg error

func loadSidsCmd() tea.Cmd {
	return func() tea.Msg {
		f, err := os.ReadFile("sids.json")
		if err != nil {
			return errMsg(err)
		}

		var sids map[string]SidInfo
		if err := json.Unmarshal(f, &sids); err != nil {
			return errMsg(err)
		}
		return sidsLoadedMsg(sids)
	}
}
