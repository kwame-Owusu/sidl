package tui

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
