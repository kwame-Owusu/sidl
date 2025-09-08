package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kwame-Owusu/sidl/internal/tui"

	"github.com/spf13/cobra"
)

var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "Get the types of sids",
	Long:  `Get the types of sids present in the sids.json`,
	Run:   listTypes,
}

func listTypes(cmd *cobra.Command, args []string) {
	sidTypes := []string{}
	for _, val := range sids {
		name := val.Name
		if name == "" {
			continue
		}
		sidTypes = append(sidTypes, name)
	}
	for _, sidType := range sidTypes {
		p := tea.NewProgram(
			tui.NewTypesModel(sidType),
		)
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
		}
	}
}
