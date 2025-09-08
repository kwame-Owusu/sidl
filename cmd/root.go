package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kwame-Owusu/sidl/internal"
	"github.com/kwame-Owusu/sidl/internal/tui"
	"github.com/spf13/cobra"
	"os"
)

var sids map[string]internal.Field

var rootCmd = &cobra.Command{
	Use:   "sidl",
	Short: "A CLI tool for twilio sids",
	Long:  `A CLI tool to get detailed information about twilio sids`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.NewModel(tui.ModeHome))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version of sidl",
	Long:  `Get version of sidl`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(explainCmd)
	rootCmd.AddCommand(explainPrefixCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(typesCmd)
	rootCmd.AddCommand(tuiCmd)
	sids = internal.LoadSIDs()
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		// Exit with error code 1 if something goes wrong
		// os.Exit(1)  // optional
	}
}
