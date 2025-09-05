package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sids",
	Long:  `List all available sids in the sids.json file`,
	Run:   listSids,
}

func listSids(cmd *cobra.Command, args []string) {
	fmt.Printf("%-4s %-10s %-14s\n", "KEY", "NAME", "DESCRIPTION")
	fmt.Println(strings.Repeat("-", 90))

	for key, val := range sids {
		fmt.Printf("%-4s %-10s %-14s ", key, val.Name, val.Description)
	}
}
