package cmd

import (
	"fmt"
	"strings"

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
	upperSid := strings.ToUpper(sid)
	prefix := strings.ToUpper(sid[0:2]) // Make sure prefix is uppercase too

	isValidSid, err := internal.IsValidSid(upperSid, sids)
	if err != nil {
		fmt.Printf("Error occurred while checking if sid is valid: %s\n", err)
		return
	}

	if !isValidSid {
		fmt.Printf("Unknown SID prefix: %s\n", prefix)
		return
	}

	// Use the uppercase prefix to lookup in the map
	if description, ok := sids[prefix]; ok {
		fmt.Printf("Name: %s\n", description.Name)
		fmt.Printf("Description: %s\n", description.Description)
	} else {
		fmt.Printf("SID prefix %s not found in sids.json\n", prefix)
	}
}
