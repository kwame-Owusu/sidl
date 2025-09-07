package tui

import (
	"fmt"
	"sort"
)

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPress q to quit.", m.err)
	}
	if len(m.sids) == 0 {
		return "No SIDs found.\nPress q to quit."
	}

	// Sort keys for consistent output
	keys := make([]string, 0, len(m.sids))
	for k := range m.sids {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	out := "SIDs from sids.json:\n\n"
	for i, k := range keys {
		info := m.sids[k]
		name := "<none>"
		if info.Name != nil {
			name = *info.Name
		}
		out += fmt.Sprintf("%d. %s → %s\n", i+1, k, name)
	}
	out += "\nPress q to quit."
	return out
}
