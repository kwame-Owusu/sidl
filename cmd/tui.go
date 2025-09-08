package cmd

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/kwame-Owusu/sidl/internal/tui"
	"github.com/spf13/cobra"
	"os"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Run sidl in interactive TUI mode",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.NewModel(tui.ModeList))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}
