package cmd

import (
	"fmt"

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
		fmt.Printf("- %s\n", sidType)
	}
}
