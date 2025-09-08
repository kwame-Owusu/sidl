package cmd

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kwame-Owusu/sidl/internal"
	"github.com/kwame-Owusu/sidl/internal/tui"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain [description]",
	Short: "Get info about a sid",
	Long:  "Get the detailed information about a SID",
	Args:  cobra.MinimumNArgs(1),
	Run:   explainSid,
}

var explainPrefixCmd = &cobra.Command{
	Use:   "explain-prefix [description]",
	Short: "Get info about a prefix",
	Long:  "Get the detailed information about a SID based on its prefix",
	Args:  cobra.MinimumNArgs(1),
	Run:   explainPrefix,
}

func explainSid(cmd *cobra.Command, args []string) {
	sid := args[0]
	upperSid := strings.ToUpper(sid)
	prefix := strings.ToUpper(sid[0:2])

	isValidSid, err := internal.IsValidSid(upperSid, sids)
	if err != nil {
		fmt.Printf("Error occurred while checking if sid is valid: %s\n", err)
		return
	}
	if !isValidSid {
		fmt.Printf("Unknown SID prefix: %s\n", prefix)
		return
	}

	if description, ok := sids[prefix]; ok {
		// Run TUI
		p := tea.NewProgram(
			tui.NewExplainModel(description.Name, description.Description, prefix),
		)
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
		}
	} else {
		fmt.Printf("SID prefix %s not found in sids.json\n", prefix)
	}
}

func explainPrefix(cmd *cobra.Command, args []string) {
	sid := args[0]
	if len(sid) > 2 || len(sid) < 2 {
		fmt.Printf("Prefix command only needs the sid prefix, which is exactly 2 characters")
		return
	}
	prefix := strings.ToUpper(sid[0:2])

	if description, ok := sids[prefix]; ok {
		fmt.Printf("Name: %s\n", description.Name)
		fmt.Printf("Description: %s\n", description.Description)
	} else {
		fmt.Printf("SID prefix %s not found in sids.json\n", prefix)
	}
}
