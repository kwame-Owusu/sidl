package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kwame-Owusu/sidl/internal/tui"
	"github.com/spf13/cobra"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sids",
	Long:  `List all available information about a sid`,
	Run:   listSids,
}

func listSids(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(tui.NewModel(tui.ModeList))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
