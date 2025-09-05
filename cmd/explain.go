package cmd

import (
	"fmt"

	"github.com/kwame-Owusu/sidl/internal"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain [description]",
	Short: "Get info about a sid",
	Long:  "Get the detailed information about a SID",
	Args:  cobra.MinimumNArgs(1),
	Run:   explainSid,
}

func explainSid(cmd *cobra.Command, args []string) {
	sid := args[0]
	isValidSid, err := internal.IsValidSid(sid, sids)
	if err != nil {
		fmt.Printf("Error occured while checking if sid is valid, %s\n", err)
	}
	if !isValidSid {
		argIdentifier := args[0][0:2]
		fmt.Printf("Unknown SID prefix: %s\n", argIdentifier)
	}

	if description, ok := sids[sid[0:2]]; ok {
		fmt.Printf("description: %s\n", description.Description)
	}
}
