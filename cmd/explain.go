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
		fmt.Printf("Error occured while checking if sid is valid: %s", err)
	}
	if !isValidSid {
		fmt.Printf("Not a valid SID, please refer to twilio docs or check sidly list, %s", args[0])
	}

	if description, ok := sids[sid]; ok {
		fmt.Printf("description: %s", description.Description)
	}
}
