package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sids",
	Long:  `List all available information about a sid`,
	Run:   listSids,
}

func listSids(cmd *cobra.Command, args []string) {
	fmt.Printf("%-4s %-20s %-80s\n", "KEY", "NAME", "DESCRIPTION")
	fmt.Println(strings.Repeat("-", 105))
	for key, val := range sids {
		fmt.Printf("%-4s %-20s %-80s\n", key, val.Name, val.Description)
	}
}
