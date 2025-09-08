package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"sort"
)

// SidItem represents one SID entry in the list.
type SidItem struct {
	Key  string
	Info SidInfo
}

func (i SidItem) Title() string {
	// Show the SID prefix (e.g., "MG")
	return i.Key
}

func (i SidItem) Description() string {
	// Show the name (or <none> if null)
	if i.Info.Name != nil {
		return *i.Info.Name
	}
	return "<none>"
}

func (i SidItem) FilterValue() string {
	// Used when filtering list items (optional)
	return i.Key
}

func toListItems(m map[string]SidInfo) []list.Item {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	items := make([]list.Item, 0, len(keys))
	for _, k := range keys {
		items = append(items, SidItem{Key: k, Info: m[k]})
	}
	return items
}
